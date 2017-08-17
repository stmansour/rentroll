package bizlogic

import "rentroll/rlib"

// ValidateAcctRule ensures that the data in the supplied
// account rule is valid. It returns descriptive errors for data
// that is not valid.
//
// @Params
//		a = the AccountRule to check
//
// @Returns
//       errlist - list of violations.  If len(errlist) == 0 then
//                 no errors were found
func ValidateAcctRule(a *rlib.AR) []BizError {
	var e []BizError

	if a.CreditLID == 0 {
		rlib.Console("*** ERROR ***  a.CreditLid = %d\n", a.CreditLID)
		e = AddBizErrToList(e, BadCreditAccount)
	} else {
		l := rlib.GetLedger(a.CreditLID)
		if l.LID == 0 || l.BID != a.BID {
			rlib.Console("*** ERROR ***  could not load a.CreditLid = %d\n", a.CreditLID)
			e = AddBizErrToList(e, BadCreditAccount)
		}
	}
	rlib.Console("after credit acct check: len(e) = %d\n", len(e))
	if a.DebitLID == 0 {
		rlib.Console("*** ERROR ***  a.DebitLid = %d\n", a.DebitLID)
		e = AddBizErrToList(e, BadDebitAccount)
	} else {
		l := rlib.GetLedger(a.DebitLID)
		if l.LID == 0 || l.BID != a.BID {
			rlib.Console("*** ERROR ***  could not load a.DebitLid = %d\n", a.DebitLID)
			e = AddBizErrToList(e, BadDebitAccount)
		}
	}
	rlib.Console("after debit acct check: len(e) = %d\n", len(e))
	return e
}
