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
