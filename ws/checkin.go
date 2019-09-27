package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	row := rlib.RRdb.Dbrr.QueryRow(q)
	if err = row.Scan(&RAID, &dt, &LeaseStatus, &ASMID, &ARID, &FLAGS); err != nil && !rlib.IsSQLNoResultsError(err) {
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

	SvcWriteSuccessResponse(d.BID, w)
}
