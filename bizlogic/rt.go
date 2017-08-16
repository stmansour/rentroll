package bizlogic

import "rentroll/rlib"

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
