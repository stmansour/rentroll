package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"strings"
	"time"
)

// CheckInData is the payload data for the CheckIn Command.
type CheckInData struct {
	TZOffset int
}

// SvcCheckIn performs the tasks needed to check in a user based on the
// supplied RLID.
//
// wsdoc {
//  @Title  CheckIn
//	@URL /v1/checkin/:BUI/RLID
//  @Method  POST
//	@Synopsis Performs check in functions
//  @Description  Performs the following tasks:
//  @Description  *  Ensures that there is a rent assessment for RLID and makes it active.
//  @Description  *  Sets the LeaseStatus field to 1 (rented)
//	@Input WebGridSearchRequest
//  @Response Reservation
// wsdoc }
//------------------------------------------------------------------------------
func SvcCheckIn(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcCheckIn"
	var err error

	rlib.Console("entered %s, getting BID = %d, RLID = %d\n", funcname, d.BID, d.ID)

	target := `"record":`
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("%s: cannot find %s in form json", funcname, target)
		SvcErrorReturn(w, e, funcname)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	//---------------------------------------------------
	// Read the payload data from the client
	//---------------------------------------------------
	var chkin CheckInData
	err = json.Unmarshal([]byte(s), &chkin)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	rlib.Console("Successfully read payload:  TZOffset = %d\n", chkin.TZOffset)

	//---------------------------------------------------------------------------
	// See if there are any rent active Assessments. If so, they've already been
	// checked in.
	//---------------------------------------------------------------------------
	var dt rlib.NullDate
	var ARID, ASMID, LeaseStatus, RAID, FLAGS rlib.NullInt64
	// var DefaultAmount rlib.NullFloat64
	var rows *sql.Rows
	q := fmt.Sprintf(`
        SELECT RentableLeaseStatus.RAID,
        		RentableLeaseStatus.DtStart,
        		RentableLeaseStatus.LeaseStatus,
				Assessments.ASMID,
                AR.ARID,
                AR.FLAGS
        FROM RentableLeaseStatus
        	LEFT JOIN Assessments ON (Assessments.RAID = RentableLeaseStatus.RAID AND Assessments.FLAGS & 2 = 0)
            LEFT JOIN AR ON (AR.ARID = Assessments.ARID AND AR.FLAGS & 16 > 0)
        WHERE
        	RentableLeaseStatus.RLID=%d;`, d.ID)
	if rows, err = rlib.RRdb.Dbrr.Query(q); err != nil {
		if !rlib.IsSQLNoResultsError(err) {
			SvcErrorReturn(w, err, funcname)
			return
		}
	}

	for i := 0; rows.Next(); i++ {
		if err = rows.Scan(&RAID, &dt, &LeaseStatus, &ASMID, &ARID /* &DefaultAmount, */, &FLAGS); err != nil && !rlib.IsSQLNoResultsError(err) {
			SvcErrorReturn(w, err, funcname)
			return
		}

		//--------------------------------------------------
		// Check for the existence of a rent assessment...
		//--------------------------------------------------
		if ASMID.Valid && ASMID.Int64 > 0 && FLAGS.Valid && FLAGS.Int64&(1<<4) > 0 {
			SvcErrorReturn(w, fmt.Errorf("There is already an active rent assessment (ASMID %d) for this reservation", ASMID.Int64), funcname)
			return
		}
	}
	//--------------------------------------------------
	// Make sure that the status is "reserved"
	//--------------------------------------------------
	if LeaseStatus.Valid && LeaseStatus.Int64 != 2 {
		SvcErrorReturn(w, fmt.Errorf("The lease status for this reservation was not \"reserved\", it was %d", LeaseStatus.Int64), funcname)
		return
	}

	rlib.Console("ARID = %d, FLAGS = %d (hex: %x,  bin: %b)\n", ARID.Int64, FLAGS.Int64, FLAGS.Int64, FLAGS.Int64)

	//---------------------------------------------------
	// If the dates don't match, do not accept...
	//---------------------------------------------------
	now := rlib.Now()
	dt1 := dt.Time.Add(time.Duration(-chkin.TZOffset))
	rlib.Console("Adjusted server-now, res DtStart = %s\n", rlib.ConsoleDRange(&now, &dt1))

	if dt1.Year() != now.Year() || dt1.Month() != now.Month() || dt1.Day() != now.Day() {
		SvcErrorReturn(w, fmt.Errorf("Check-in for this reservation can only occur on %s", rlib.ConDt(&dt.Time)), funcname)
		return
	}

	//---------------------------------------------------
	// If we get to this point, we can continue:
	//    update the rentable lease status to rented
	//    create an active rent assessment
	//---------------------------------------------------
	var rls rlib.RentableLeaseStatus
	if rls, err = rlib.GetRentableLeaseStatus(r.Context(), d.ID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("rls:  RLID=%d, start/stop = %s\n", rls.RLID, rlib.ConsoleDRange(&rls.DtStart, &rls.DtStop))

	var amt float64
	var rc, pc, ar int64
	if amt, rc, pc, ar, err = getReservationLeaseAmount(r.Context(), rls.RID, &rls.DtStart, &rls.DtStop); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if rc == 0 {
		err = fmt.Errorf("Could not find info needed to make rent assessment.  Amount = %6.2f, RentCycle = %d, ProationCycle = %d, ARID = %d", amt, rc, pc, ar)
		SvcErrorReturn(w, err, funcname)
		return
	}

	var asm = rlib.Assessment{
		BID:            rls.BID,
		RID:            rls.RID,
		RAID:           rls.RAID,
		Amount:         amt,
		Start:          rls.DtStart,
		Stop:           rls.DtStop,
		RentCycle:      rc,
		ProrationCycle: pc,
		ARID:           ar,
		Comment:        fmt.Sprintf("Reservation %s (RLID=%d)", rls.ConfirmationCode, rls.RLID),
	}

	//---------------------------------------
	// start transaction
	//---------------------------------------
	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rls.LeaseStatus = rlib.LEASESTATUSleased // set lease status to rented
	if err = rlib.UpdateRentableLeaseStatus(ctx, &rls); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	//---------------------------------------
	// create a recurring rent assessment
	//---------------------------------------
	if be := bizlogic.InsertAssessment(ctx, &asm, 0, &noClose); len(be) > 0 {
		err = bizlogic.BizErrorListToError(be)
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	//---------------------------------------
	// Create the first instance assessment
	//---------------------------------------
	var ai = rlib.Assessment{
		BID:            rls.BID,
		RID:            rls.RID,
		RAID:           rls.RAID,
		PASMID:         asm.ASMID,
		Amount:         amt,
		Start:          rls.DtStart,
		Stop:           rls.DtStart,
		RentCycle:      rc,
		ProrationCycle: pc,
		ARID:           ar,
		Comment:        fmt.Sprintf("Reservation %s (RLID=%d)", rls.ConfirmationCode, rls.RLID),
	}
	if be := bizlogic.InsertAssessment(ctx, &ai, 0, &noClose); len(be) > 0 {
		err = bizlogic.BizErrorListToError(be)
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	//---------------------------------------
	// commit
	//---------------------------------------
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Inserted rent assessment.  ASMID = %d, Amount = %6.2f\n", asm.ASMID, asm.Amount)

	SvcWriteSuccessResponse(d.BID, w)
}

// getReservationLeaseAmount returns the amount for the rent assessment on
// the Lease
//
// INPUTS
//
// RETURNS
//   amt = amount of rent
//   rc =  rentcycle
//   pc =  prorationCycle
//   ar =  ARID, account rule to use for assessment
//   error = any errors encountered
//----------------------------------------------------------------------------
func getReservationLeaseAmount(ctx context.Context, RID int64, DtStart, DtStop *time.Time) (float64, int64, int64, int64, error) {
	var row *sql.Row
	var amt float64
	var rc, pc, ar int64
	var err error
	dt := DtStart.Format(rlib.RRDATEFMTSQL)
	q := fmt.Sprintf(`
		SELECT AR.DefaultAmount,
                AR.DefaultRentCycle,
                AR.DefaultProrationCycle,
				RentableTypes.ARID
        FROM RentableTypeRef
        LEFT JOIN RentableTypes ON RentableTypes.RTID = RentableTypeRef.RTID
        LEFT JOIN AR ON (AR.ARID = RentableTypes.ARID)
        WHERE
        	RentableTypeRef.RID=%d AND RentableTypeRef.DtStart <= %q AND %q < RentableTypeRef.DtStop;`, RID, dt, dt)
	rlib.Console("getResLeaseAmount. query = %s\n", q)
	row = rlib.RRdb.Dbrr.QueryRow(q)
	if err = row.Scan(&amt, &rc, &pc, &ar); err != nil {
		if !rlib.IsSQLNoResultsError(err) {
			return amt, rc, pc, ar, err
		}
	}
	return amt, rc, pc, ar, nil
}
