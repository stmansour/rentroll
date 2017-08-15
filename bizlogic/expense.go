package bizlogic

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// ReverseExpense reverse an expense. If the Expense has already been reversed
// it returns immediately.
//-----------------------------------------------------------------------------
func ReverseExpense(aold *rlib.Expense, dt *time.Time) []BizError {
	var errlist []BizError
	if aold.FLAGS&0x4 != 0 {
		return nil // it's already reversed
	}

	anew := *aold
	anew.EXPID = 0
	anew.Amount = -anew.Amount
	anew.RPEXPID = aold.EXPID
	anew.FLAGS |= 0x4 // set bit 2 to mark that this assessment is void
	anew.Comment = fmt.Sprintf("Reversal of %s", aold.IDtoShortString())

	err := rlib.InsertExpense(&anew)
	if err != nil {
		return bizErrSys(&err)
	}
	var xbiz rlib.XBusiness
	rlib.ProcessNewExpense(&anew, &xbiz)

	aold.Comment = fmt.Sprintf("Reversed by %s", anew.IDtoShortString())
	aold.FLAGS |= 0x4 // set bit 2 to mark that this assessment is void
	err = rlib.UpdateExpense(aold)
	if err != nil {
		return bizErrSys(&err)
	}

	return errlist
}
