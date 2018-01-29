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

// ValidateRentableStatus checks for validity of a given rentable status
// while insert and update in db
func ValidateRentableStatus(ctx context.Context, rs *rlib.RentableStatus) []BizError {
	var errlist []BizError

	// 1. First check BID is valid or not
	if !(rlib.IsBIDExists(rs.BID)) {
		s := fmt.Sprintf(BizErrors[UnknownBID].Message, rs.BID)
		b := BizError{Errno: UnknownBID, Message: s}
		errlist = append(errlist, b)
	}

	// 2. check UseStatus is valid or not
	if !(0 <= rs.UseStatus && rs.UseStatus < int64(len(rlib.RSUseStatus))) {
		s := fmt.Sprintf(BizErrors[InvalidRentableUseStatus].Message, rs.UseStatus)
		b := BizError{Errno: InvalidRentableUseStatus, Message: s}
		errlist = append(errlist, b)
	}

	// 3. check LeaseStatus is valid or not
	if !(0 <= rs.LeaseStatus && rs.LeaseStatus < int64(len(rlib.RSLeaseStatus))) {
		s := fmt.Sprintf(BizErrors[InvalidRentableLeaseStatus].Message, rs.LeaseStatus)
		b := BizError{Errno: InvalidRentableLeaseStatus, Message: s}
		errlist = append(errlist, b)
	}

	// 4. Stopdate should not be before startDate
	if rs.DtStop.Before(rs.DtStart) {
		s := fmt.Sprintf(BizErrors[InvalidRentableStatusDates].Message,
			rs.RSID, rs.DtStop.Format(rlib.RRDATEFMT4), rs.DtStart.Format(rlib.RRDATEFMT4))
		b := BizError{Errno: InvalidRentableStatusDates, Message: s}
		errlist = append(errlist, b)
	}

	// 5. check that DtStart and DtStop don't overlap/fall in with other object
	// associated with the same RID
	// TODO(Sudip): ignore error as of now
	rsList, _ := rlib.GetAllRentableStatus(ctx, rs.RID)

	for _, rsRow := range rsList {
		// if same object then continue
		if rs.RSID == rsRow.RSID {
			continue
		}
		// start date should not sit between other market rate's time span
		if rlib.DateRangeOverlap(&rs.DtStart, &rs.DtStop, &rsRow.DtStart, &rsRow.DtStop) {
			s := fmt.Sprintf(BizErrors[RentableStatusDatesOverlap].Message, rs.RSID, rsRow.RSID)
			b := BizError{Errno: RentableStatusDatesOverlap, Message: s}
			errlist = append(errlist, b)
		}
	}
	return errlist
}

// ValidateRentableTypeRef checks for validity of a given rentable type ref
// while insert and update in db
func ValidateRentableTypeRef(ctx context.Context, rtr *rlib.RentableTypeRef) []BizError {
	var errlist []BizError

	// 1. First check BID is valid or not
	if !(rlib.IsBIDExists(rtr.BID)) {
		s := fmt.Sprintf(BizErrors[UnknownBID].Message, rtr.BID)
		b := BizError{Errno: UnknownBID, Message: s}
		errlist = append(errlist, b)
	}

	// 2. Stopdate should not be before startDate
	if rtr.DtStop.Before(rtr.DtStart) {
		s := fmt.Sprintf(BizErrors[InvalidRentableTypeRefDates].Message,
			rtr.RTRID, rtr.DtStop.Format(rlib.RRDATEFMT4), rtr.DtStart.Format(rlib.RRDATEFMT4))
		b := BizError{Errno: InvalidRentableTypeRefDates, Message: s}
		errlist = append(errlist, b)
	}

	// TODO(Sudip): make query to get any rentable type ref that bring validation false here

	/*// 3. check that DtStart and DtStop don't overlap/fall in with other object
	// associated with the same RID
	rsList := rlib.GetAllRentableStatus(ctx, rtr.RID)

	for _, rsRow := range rsList {
		// if same object then continue
		if rtr.RSID == rsRow.RSID {
			continue
		}
		// start date should not sit between other market rate's time span
		if rlib.DateRangeOverlap(&rtr.DtStart, &rtr.DtStop, &rsRow.DtStart, &rsRow.DtStop) {
			s := fmt.Sprintf(BizErrors[RentableStatusDatesOverlap].Message, rtr.RMRID, rsRow.RMRID)
			b := BizError{Errno: RentableStatusDatesOverlap, Message: s}
			errlist = append(errlist, b)
		}
	}*/
	return errlist
}
