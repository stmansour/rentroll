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
	anew.FLAGS |= 0x4 // set bit 2 to mark that this expense is void
	anew.Comment = fmt.Sprintf("Reversal of %s", aold.IDtoShortString())

	err := rlib.InsertExpense(&anew)
	if err != nil {
		return bizErrSys(&err)
	}
	var xbiz rlib.XBusiness
	rlib.ProcessNewExpense(&anew, &xbiz)

	aold.Comment = fmt.Sprintf("Reversed by %s", anew.IDtoShortString())
	aold.FLAGS |= 0x4 // set bit 2 to mark that this expense is void
	err = rlib.UpdateExpense(aold)
	if err != nil {
		return bizErrSys(&err)
	}

	return errlist
}

// UpdateExpense updates the supplied expense, reversing existing expenses
// if necessary
//
// INPUTS
//    a = the expense to update
//  exp = if it is a recurring expense and the start date is in the past, should
//        past entries be created?  true = yes
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func UpdateExpense(anew *rlib.Expense, dt *time.Time) []BizError {
	var err error
	var errlist []BizError

	rlib.Console("Entered bizlogic.UpdateExpense:  anew.EXPID = %d, dt = %s\n", anew.EXPID, dt.Format(rlib.RRDATEREPORTFMT))
	rlib.Console("anew.FLAGS = %X\n", anew.FLAGS)

	if anew.FLAGS&0x4 != 0 {
		errlist = append(errlist, BizErrors[EditReversal])
		return errlist
	}
	//-------------------------------
	// Load existing expense...
	//-------------------------------
	aold, err := rlib.GetExpense(anew.EXPID)
	if err != nil {
		return bizErrSys(&err)
	}
	if aold.EXPID == 0 {
		err = fmt.Errorf("Expense %d not found", anew.EXPID)
		return bizErrSys(&err)
	}

	//---------------------------------------------------------------------------------
	// we need to reverse the old expense if any of the following fields have changed:
	//   ARID
	//   Amount
	//   Dt
	//---------------------------------------------------------------------------------
	if aold.ARID != anew.ARID || aold.Amount != anew.Amount || (!aold.Dt.Equal(anew.Dt)) {
		errlist = ReverseExpense(&aold, dt) // reverse the expense itself
		if errlist != nil {
			return errlist
		}
		anew.EXPID = 0 // need to insert a new record with the updated info
		err := rlib.InsertExpense(anew)
		if err != nil {
			return bizErrSys(&err)
		}
	} else {
		err = rlib.UpdateExpense(anew) // reversal not needed, just update the expense
		if err != nil {
			return bizErrSys(&err)
		}
	}
	return nil
}
