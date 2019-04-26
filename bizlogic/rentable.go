package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"strings"
)

// InsertRentable first validates that inserting the rentable does
// not violate any business rules. If there are no violations
// it will insert the rentable.
//
// INPUTS
//  r - the rentable to insert
//
// RETURNS
//  a slice of BizErrors encountered
//-----------------------------------------------------------------------------
func InsertRentable(ctx context.Context, r *rlib.Rentable) []BizError {
	var be []BizError
	//-------------------------------------------------------------
	// Check 1:  does a Rentable with the same name already exist?
	//-------------------------------------------------------------
	r1, err := rlib.GetRentableByName(ctx, r.RentableName, r.BID)
	if err != nil {
		s := err.Error()
		if !strings.Contains(s, "no rows") {
			return AddErrToBizErrlist(err, be)
		}
	}
	if r1.RID > 0 {
		s := fmt.Sprintf(BizErrors[RentableNameExists].Message, r.RentableName, r.BID)
		b := BizError{Errno: RentableNameExists, Message: s}
		return append(be, b)
	}
	_, err = rlib.InsertRentable(ctx, r)
	if err != nil {
		return AddErrToBizErrlist(err, be)
	}
	return nil
}

// ValidateRentableLeaseStatus checks for validity of a given rentable status , add by lina on 01/24/2019
// while insert and update in db
func ValidateRentableLeaseStatus(ctx context.Context, rl *rlib.RentableLeaseStatus) []BizError {
	var errlist []BizError
	// rlib.Console("VRLS: 0\n")
	// 1. First check BID is valid or not
	if !(rlib.BIDExists(rl.BID)) {
		s := fmt.Sprintf(BizErrors[UnknownBID].Message, rl.BID)
		b := BizError{Errno: UnknownBID, Message: s}
		errlist = append(errlist, b)
	}

	// rlib.Console("VRLS: 1\n")
	// check for RID as well
	if rl.RID < 1 {
		s := fmt.Sprintf(BizErrors[UnknownRID].Message, rl.RID)
		b := BizError{Errno: UnknownRID, Message: s}
		errlist = append(errlist, b)
	}

	// rlib.Console("VRLS: 2\n")
	// 2. check UseStatus is valid or not
	//if !(0 <= rs.UseStatus && rs.UseStatus < int64(len(rlib.RSUseStatus))) {
	//	s := fmt.Sprintf(BizErrors[InvalidRentableUseStatus].Message, rs.UseStatus)
	//	b := BizError{Errno: InvalidRentableUseStatus, Message: s}
	//	errlist = append(errlist, b)
	//}

	// rlib.Console("VRLS: 3\n")
	// 3. check LeaseStatus is valid or not
	if !(0 <= rl.LeaseStatus && rl.LeaseStatus < int64(len(rlib.RSLeaseStatus))) {
		s := fmt.Sprintf(BizErrors[InvalidRentableLeaseStatus].Message, rl.LeaseStatus)
		b := BizError{Errno: InvalidRentableLeaseStatus, Message: s}
		errlist = append(errlist, b)
	}

	// rlib.Console("VRLS: 4\n")
	// rlib.Console("VRLS: 4\n")
	// 4. Stopdate should not be before startDate
	//please check bizlogic/init.go line 66 to check if const InvalidRentableLeaseStatusDates need to update the number
	if rl.DtStop.Before(rl.DtStart) {
		s := fmt.Sprintf(BizErrors[InvalidRentableLeaseStatusDates].Message,
			rl.RLID, rl.DtStart.Format(rlib.RRDATEFMT4), rl.DtStop.Format(rlib.RRDATEFMT4))
		b := BizError{Errno: InvalidRentableLeaseStatusDates, Message: s}
		errlist = append(errlist, b)
	}

	// rlib.Console("VRLS: 5.  checking range: %s\n", rlib.ConsoleDRange(&rl.DtStart, &rl.DtStop))
	// 5. check that DtStart and DtStop don't overlap/fall in with other object.
	// associated with the same RID
	//
	// sman: 2/27/2019 - we cannot use this check as is.  You can make multiple
	//    in the UI or in the web interface -- individually a given change can
	//    appear to be an overlap violation, but if all the changes are taken as
	//    a whole then there are no overlap violations.  But the code below does
	//    not take the changes as a whole.  So I'm commenting it out for now.
	// overLappingRLQuery := `
	// SELECT
	// 	RLID
	// FROM RentableLeaseStatus
	// WHERE
	// 	RLID <> {{.RLID}} AND
	// 	DtStart < "{{.stopDate}}" AND
	// 	"{{.startDate}}" < DtStop AND
	// 	RID = {{.RID}} AND
	// 	BID = {{.BID}}
	// LIMIT 1`
	//
	// qc := rlib.QueryClause{
	// 	"BID":       strconv.FormatInt(rl.BID, 10),
	// 	"RID":       strconv.FormatInt(rl.RID, 10),
	// 	"RLID":      strconv.FormatInt(rl.RLID, 10),
	// 	"startDate": rl.DtStart.Format(rlib.RRDATEFMTSQL),
	// 	"stopDate":  rl.DtStop.Format(rlib.RRDATEFMTSQL),
	// }
	//
	// qry := rlib.RenderSQLQuery(overLappingRLQuery, qc)
	//
	// rlib.Console("Rentable Lease Status overlap qry = %s\n", qry)
	// row := rlib.RRdb.Dbrr.QueryRow(qry)
	//
	// var overLappingRLID int64
	// err := row.Scan(&overLappingRLID)
	// rlib.SkipSQLNoRowsError(&err)
	// if err != nil {
	// 	b := BizError{Errno: -1, Message: err.Error()}
	// 	errlist = append(errlist, b)
	// 	return errlist // we're done if this happens
	// }
	// // rlib.Console("VRLS: 6\n")
	//
	// if overLappingRLID > 0 {
	// 	// rlib.Console("VRLS: 6.1  -  RentableLeaseStatusDatesOverlap %d\n", RentableLeaseStatusDatesOverlap)
	// 	s := fmt.Sprintf(BizErrors[RentableLeaseStatusDatesOverlap].Message,
	// 		rl.RLID,
	// 		rlib.ConsoleDate(&rl.DtStart),
	// 		rlib.ConsoleDate(&rl.DtStop),
	// 		overLappingRLID)
	// 	// rlib.Console("VRLS: 6.2\n")
	// 	b := BizError{Errno: RentableLeaseStatusDatesOverlap, Message: s}
	// 	errlist = append(errlist, b)
	// }

	// rlib.Console("VRLS: 7\n")
	return errlist
}

// ValidateRentableUseStatus checks for validity of a given rentable status
// while insert and update in db
func ValidateRentableUseStatus(ctx context.Context, rs *rlib.RentableUseStatus) []BizError {
	var errlist []BizError

	// 1. First check BID is valid or not
	if !(rlib.BIDExists(rs.BID)) {
		s := fmt.Sprintf(BizErrors[UnknownBID].Message, rs.BID)
		b := BizError{Errno: UnknownBID, Message: s}
		errlist = append(errlist, b)
	}

	// check for RID as well
	if rs.RID < 1 {
		s := fmt.Sprintf(BizErrors[UnknownRID].Message, rs.RID)
		b := BizError{Errno: UnknownRID, Message: s}
		errlist = append(errlist, b)
	}

	// 2. check UseStatus is valid or not
	if !(0 <= rs.UseStatus && rs.UseStatus < int64(len(rlib.RSUseStatus))) {
		s := fmt.Sprintf(BizErrors[InvalidRentableUseStatus].Message, rs.UseStatus)
		b := BizError{Errno: InvalidRentableUseStatus, Message: s}
		errlist = append(errlist, b)
	}

	// 3. check LeaseStatus is valid or not
	// if !(0 <= rs.LeaseStatus && rs.LeaseStatus < int64(len(rlib.RSLeaseStatus))) {
	// 	s := fmt.Sprintf(BizErrors[InvalidRentableLeaseStatus].Message, rs.LeaseStatus)
	// 	b := BizError{Errno: InvalidRentableLeaseStatus, Message: s}
	// 	errlist = append(errlist, b)
	// }

	// 4. Stopdate should not be before startDate
	if rs.DtStop.Before(rs.DtStart) {
		s := fmt.Sprintf(BizErrors[InvalidRentableUseStatusDates].Message,
			rs.RSID, rs.DtStop.Format(rlib.RRDATEFMT4), rs.DtStart.Format(rlib.RRDATEFMT4))
		b := BizError{Errno: InvalidRentableUseStatusDates, Message: s}
		errlist = append(errlist, b)
	}

	// 5. check that DtStart and DtStop don't overlap/fall in with other object
	// associated with the same RID
	// sman: 2/27/2019 - we cannot use this check as is.  You can make multiple
	//    in the UI or in the web interface -- individually a given change can
	//    appear to be an overlap violation, but if all the changes are taken as
	//    a whole then there are no overlap violations.  But the code below does
	//    not take the changes as a whole.  So I'm commenting it out for now.
	// overLappingRSQuery := `
	// SELECT
	// 	RSID
	// FROM RentableUseStatus
	// WHERE
	// 	RSID <> {{.RSID}} AND
	// 	DtStart < "{{.stopDate}}" AND
	// 	"{{.startDate}}" < DtStop AND
	// 	RID = {{.RID}} AND
	// 	BID = {{.BID}}
	// LIMIT 1`
	//
	// qc := rlib.QueryClause{
	// 	"BID":       strconv.FormatInt(rs.BID, 10),
	// 	"RID":       strconv.FormatInt(rs.RID, 10),
	// 	"RSID":      strconv.FormatInt(rs.RSID, 10),
	// 	"startDate": rs.DtStart.Format(rlib.RRDATEFMTSQL),
	// 	"stopDate":  rs.DtStop.Format(rlib.RRDATEFMTSQL),
	// }
	//
	// qry := rlib.RenderSQLQuery(overLappingRSQuery, qc)
	// row := rlib.RRdb.Dbrr.QueryRow(qry)
	//
	// var overLappingRSID int64
	// err := row.Scan(&overLappingRSID)
	// rlib.SkipSQLNoRowsError(&err)
	// if err != nil {
	// 	panic(err.Error()) // BOOM!
	// }
	// if overLappingRSID > 0 {
	// 	s := fmt.Sprintf(BizErrors[RentableUseStatusDatesOverlap].Message, rs.RSID, overLappingRSID)
	// 	b := BizError{Errno: RentableUseStatusDatesOverlap, Message: s}
	// 	errlist = append(errlist, b)
	// }
	return errlist
}

// ValidateRentableTypeRef checks for validity of a given rentable type ref
// while insert and update in db
func ValidateRentableTypeRef(ctx context.Context, rtr *rlib.RentableTypeRef) []BizError {
	rlib.Console("Entered ValidateRentableTypeRef\n")
	var errlist []BizError

	rlib.Console("Entered ValidateRentableTypeRef\n")

	// 1. First check BID is valid or not
	if !(rlib.BIDExists(rtr.BID)) {
		rlib.Console("ValidateRentableTypeRef: BIDExists error]\n")
		s := fmt.Sprintf(BizErrors[UnknownBID].Message, rtr.BID)
		b := BizError{Errno: UnknownBID, Message: s}
		errlist = append(errlist, b)
	}

	// check for RID as well
	if rtr.RID < 1 {
		rlib.Console("ValidateRentableTypeRef: RID < 1 error].  rtr.RID = %d\n", rtr.RID)
		s := fmt.Sprintf(BizErrors[UnknownRID].Message, rtr.RID)
		b := BizError{Errno: UnknownRID, Message: s}
		errlist = append(errlist, b)
	}

	// check for RTID as well
	if rtr.RTID < 1 {
		rlib.Console("ValidateRentableTypeRef: RTID < 1 error].  rtr.RTID = %d\n", rtr.RTID)
		s := fmt.Sprintf(BizErrors[UnknownRTID].Message, rtr.RTID)
		b := BizError{Errno: UnknownRTID, Message: s}
		errlist = append(errlist, b)
	}

	// 2. Stopdate should not be before startDate
	if rtr.DtStop.Before(rtr.DtStart) {
		rlib.Console("ValidateRentableTypeRef: stop before start error]\n")
		s := fmt.Sprintf(BizErrors[InvalidRentableTypeRefDates].Message,
			rtr.RTRID, rtr.DtStop.Format(rlib.RRDATEFMT4), rtr.DtStart.Format(rlib.RRDATEFMT4))
		b := BizError{Errno: InvalidRentableTypeRefDates, Message: s}
		errlist = append(errlist, b)
	}

	// sman: 3/4/2019 - we cannot use this check as is.  You can make multiple
	//    in the UI or in the web interface -- individually a given change can
	//    appear to be an overlap violation, but if all the changes are taken as
	//    a whole then there are no overlap violations.  But the code below does
	//    not take the changes as a whole.  So I'm commenting it out for now.
	// // 3. Check that any other instance doesn't overlap with given date range
	// overLappingRTRQuery := `
	// SELECT
	// 	RTRID
	// FROM RentableTypeRef
	// WHERE
	// 	RTRID <> {{.RTRID}} AND
	// 	DtStart < "{{.stopDate}}" AND
	// 	"{{.startDate}}" < DtStop AND
	// 	RID = {{.RID}} AND
	// 	BID = {{.BID}}
	// LIMIT 1`
	//
	// qc := rlib.QueryClause{
	// 	"BID":       strconv.FormatInt(rtr.BID, 10),
	// 	"RID":       strconv.FormatInt(rtr.RID, 10),
	// 	"RTRID":     strconv.FormatInt(rtr.RTRID, 10),
	// 	"startDate": rtr.DtStart.Format(rlib.RRDATEFMTSQL),
	// 	"stopDate":  rtr.DtStop.Format(rlib.RRDATEFMTSQL),
	// }
	//
	// qry := rlib.RenderSQLQuery(overLappingRTRQuery, qc)
	// row := rlib.RRdb.Dbrr.QueryRow(qry)
	//
	// var overLappingRTRID int64
	// err := row.Scan(&overLappingRTRID)
	// rlib.SkipSQLNoRowsError(&err)
	// if err != nil {
	// 	panic(err.Error()) // BOOM!
	// }
	// if overLappingRTRID > 0 {
	// 	rlib.Console("ValidateRentableTypeRef: overlapping RTRID error]\n")
	// 	s := fmt.Sprintf(BizErrors[RentableTypeRefDatesOverlap].Message, rtr.RTRID, overLappingRTRID)
	// 	b := BizError{Errno: RentableTypeRefDatesOverlap, Message: s}
	// 	errlist = append(errlist, b)
	// }

	/*// 3. check that DtStart and DtStop don't overlap/fall in with other object
	// associated with the same RID
	rsList := rlib.GetAllRentableUseStatus(ctx, rtr.RID)

	for _, rsRow := range rsList {
		// if same object then continue
		if rtr.RSID == rsRow.RSID {
			continue
		}
		// start date should not sit between other market rate's time span
		if rlib.DateRangeOverlap(&rtr.DtStart, &rtr.DtStop, &rsRow.DtStart, &rsRow.DtStop) {
			s := fmt.Sprintf(BizErrors[RentableUseStatusDatesOverlap].Message, rtr.RMRID, rsRow.RMRID)
			b := BizError{Errno: RentableUseStatusDatesOverlap, Message: s}
			errlist = append(errlist, b)
		}
	}*/
	return errlist
}

// DeleteRentable first validates that it can be deleted.  If so, it performs
// the delete. Otherwise it returns an error indicating the problem.
//
// INPUTS
//  r - the rentable to delete
//
// RETURNS
//  any error encountered
//-----------------------------------------------------------------------------
func DeleteRentable(ctx context.Context, r *rlib.Rentable) error {
	var count int

	//-----------------------------------------------------------
	// Are there any assessments associated with the rentable?
	// If so, we cannot delete it.
	//-----------------------------------------------------------
	q := fmt.Sprintf("SELECT COUNT(ASMID) FROM Assessments WHERE BID = %d and RID = %d", r.BID, r.RID)
	row := rlib.RRdb.Dbrr.QueryRow(q)
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		err := fmt.Errorf("Rentable %d cannot be deleted as it is referenced by %d Assessment(s)", r.RID, count)
		return err
	}

	//-----------------------------------------------------------
	// Is this rentable called out in any rental agreement?
	// If so, we cannot delete it.
	//-----------------------------------------------------------
	q = fmt.Sprintf("SELECT COUNT(RARID) FROM RentalAgreementRentables WHERE BID = %d and RID = %d", r.BID, r.RID)
	row = rlib.RRdb.Dbrr.QueryRow(q)
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		err := fmt.Errorf("Rentable %d cannot be deleted as it is referenced by %d Rental Agreement(s)", r.RID, count)
		return err
	}

	//-----------------------------------------------------------
	// Is this rentable reserved in the future?
	// If so, we cannot delete it.
	//-----------------------------------------------------------
	now := rlib.Now()
	q = fmt.Sprintf("SELECT COUNT(RLID) FROM RentableLeaseStatus WHERE LeaseStatus=%d and BID=%d and RID=%d and DtStart > %q", rlib.LEASESTATUSreserved, r.BID, r.RID, now.Format(rlib.RRDATEFMTSQL))
	row = rlib.RRdb.Dbrr.QueryRow(q)
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		err := fmt.Errorf("Rentable %d cannot be deleted because there are %d reservation(s) for it in the future", r.RID, count)
		return err
	}

	return rlib.DeleteRentable(ctx, r.RID)
}
