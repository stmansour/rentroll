package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
)

// RTFLAGS rentable type FLAGS
var RTFLAGS = rlib.Str2Int64Map{
	"IsActive":        0, // 0 = active, 1 = inactive
	"IsChildRentable": 1, // 0 = NO > can't be child, 1 = Yes > can be child
	"ManageToBudget":  2, // 0 = No , 1 = Yes
}

// IsValidRTFlag checks whether FLAGS value is valid or not
func IsValidRTFlag(FLAGS uint64) bool {

	maxFLAGVal := 0
	for _, v := range RTFLAGS {
		maxFLAGVal += 1 << uint(v)
	}

	// NOTE: if no flag is set then 0 can be the case here
	if FLAGS < 0 || FLAGS > uint64(maxFLAGVal) {
		return false
	}

	return true
}

// ValidateRentableType does the business logic checks for inserting
// and updating a Rentable Type
func ValidateRentableType(ctx context.Context, rt *rlib.RentableType) []BizError {
	var errlist []BizError
	//--------------------------------------------------------
	// First, try to load another rentable type with the same
	// name or style...
	//--------------------------------------------------------
	if len(rt.Name) == 0 {
		errlist = AddBizErrToList(errlist, MissingName)
	}
	if len(rt.Style) == 0 {
		errlist = AddBizErrToList(errlist, MissingStyleName)
	}
	if !IsValidRTFlag(rt.FLAGS) {
		rlib.Console("*** ERROR *** invalid FLAGS: %d for rt.RTID = %d\n", rt.FLAGS, rt.RTID)
		s := fmt.Sprintf(BizErrors[InvalidRTFlag].Message, rt.FLAGS, rt.RTID)
		b := BizError{Errno: InvalidRTFlag, Message: s}
		errlist = append(errlist, b)
	}

	dup, err := rlib.GetRentableTypeByName(ctx, rt.Name, rt.BID)
	if err == nil && dup.RTID != rt.RTID && dup.RTID > 0 {
		errlist = AddBizErrToList(errlist, DuplicateName)
	}

	dup, err = rlib.GetRentableTypeByStyle(ctx, rt.Style, rt.BID)
	if err == nil && dup.RTID != rt.RTID && dup.RTID > 0 {
		errlist = AddBizErrToList(errlist, DuplicateStyleName)
	}
	return errlist
}

// ValidateRentableMarketRate checks for validity of a given rentable marketRate
// while insert and update
func ValidateRentableMarketRate(ctx context.Context, rmr *rlib.RentableMarketRate) []BizError {
	var errlist []BizError
	// NOTE: we should probably check everything here
	// like belonged biz exists or not, RTID exists or not etc...

	// 1. check first MarketRate is valid or not
	if rmr.MarketRate <= 0 {
		s := fmt.Sprintf(BizErrors[InvalidRentableMarketRateAmount].Message, rmr.RMRID, rmr.MarketRate)
		b := BizError{Errno: InvalidRentableMarketRateAmount, Message: s}
		errlist = append(errlist, b)
	}

	// 2. Stopdate should not be before startDate
	if rmr.DtStop.Before(rmr.DtStart) {
		s := fmt.Sprintf(BizErrors[InvalidRentableMRDates].Message,
			rmr.RMRID, rmr.DtStop.Format(rlib.RRDATEFMT4), rmr.DtStart.Format(rlib.RRDATEFMT4))
		b := BizError{Errno: InvalidRentableMRDates, Message: s}
		errlist = append(errlist, b)
	}

	// 3. check that DtStart and DtStop don't overlap/fall in with other MarketRate object
	// associated with the same RTID
	rt := rlib.RentableType{RTID: rmr.RTID}
	err := rlib.GetRentableMarketRates(ctx, &rt)
	if err != nil {
		return bizErrSys(&err)
	}

	for _, mr := range rt.MR {
		// if same market rate object then continue
		if mr.RMRID == rmr.RMRID {
			continue
		}
		// start date should not sit between other market rate's time span
		if rlib.DateRangeOverlap(&rmr.DtStart, &rmr.DtStop, &mr.DtStart, &mr.DtStop) {
			s := fmt.Sprintf(BizErrors[RentableMRDatesOverlap].Message, rmr.RMRID, mr.RMRID)
			b := BizError{Errno: RentableMRDatesOverlap, Message: s}
			errlist = append(errlist, b)
		}
	}
	return errlist
}
