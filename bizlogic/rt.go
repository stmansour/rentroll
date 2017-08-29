package bizlogic

import (
	"fmt"
	"rentroll/rlib"
)

// ValidateRentableType does the business logic checks for inserting
// and updating a Rentable Type
func ValidateRentableType(rt *rlib.RentableType) []BizError {
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
	dup, err := rlib.GetRentableTypeByName(rt.Name, rt.BID)
	if err == nil && dup.RTID != rt.RTID && dup.RTID > 0 {
		errlist = AddBizErrToList(errlist, DuplicateName)
	}

	dup, err = rlib.GetRentableTypeByStyle(rt.Style, rt.BID)
	if err == nil && dup.RTID != rt.RTID && dup.RTID > 0 {
		errlist = AddBizErrToList(errlist, DuplicateStyleName)
	}
	return errlist
}

// ValidateRentableMarketRate checks for validity of a given rentable marketRate
// while insert and update
func ValidateRentableMarketRate(rmr *rlib.RentableMarketRate) []BizError {
	var errlist []BizError
	// NOTE: we should probably check everything here
	// like belonged biz exists or not, RTID exists or not etc...

	// 1. check first MarketRate is valid or not
	if rmr.MarketRate <= 0 {
		errlist = AddBizErrToList(errlist, InvalidRentableMarketRate)
	}

	// 2. Stopdate should not be before startDate
	if rmr.DtStop.Before(rmr.DtStart) {
		s := fmt.Sprintf(BizErrors[InvalidRentableMRDates].Message, rmr.RMRID, -1, "EndDate is before StartDate")
		b := BizError{Errno: InvalidRentableMRDates, Message: s}
		errlist = append(errlist, b)
	}

	// 3. check that DtStart and DtStop don't overlap/fall in with other MarketRate object
	// associated with the same RTID
	rt := rlib.RentableType{RTID: rmr.RTID}
	rlib.GetRentableMarketRates(&rt)

	for _, mr := range rt.MR {
		// if same market rate object then continue
		if mr.RMRID == rmr.RMRID {
			continue
		}
		// start date should not sit between other market rate's time span
		if rmr.DtStart.After(mr.DtStart) && rmr.DtStart.Before(mr.DtStop) {
			s := fmt.Sprintf(BizErrors[InvalidRentableMRDates].Message, rmr.RMRID, mr.RMRID, "StartDate")
			b := BizError{Errno: InvalidRentableMRDates, Message: s}
			errlist = append(errlist, b)
		}
		// end date should not sit between other market rate's time span
		if rmr.DtStop.After(mr.DtStart) && rmr.DtStop.Before(mr.DtStop) {
			s := fmt.Sprintf(BizErrors[InvalidRentableMRDates].Message, rmr.RMRID, mr.RMRID, "EndDate")
			b := BizError{Errno: InvalidRentableMRDates, Message: s}
			errlist = append(errlist, b)
		}
		// start/stop should not cover other market rate's time span
		if rmr.DtStart.Before(mr.DtStart) && rmr.DtStop.After(mr.DtStop) {
			s := fmt.Sprintf(BizErrors[InvalidRentableMRDates].Message, rmr.RMRID, mr.RMRID, "Covering")
			b := BizError{Errno: InvalidRentableMRDates, Message: s}
			errlist = append(errlist, b)
		}
	}
	return errlist
}
