package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
)

// ARFLAGS account rules FLAGS
var ARFLAGS = rlib.Str2Int64Map{
	"ApplyFundsToReceiveAccts": 0,
	"AutoPopulateToNewRA":      1,
	"RAIDRequired":             2,
	"SubARIDsOnly":             3,
	"IsRentASM":                4,
	"IsSecDepASM":              5,
	"IsNonRecurCharge":         6,
}

// ARType user defined type of account rule
type ARType int64

// AssessmentAR etc... are all constant for Account Rule Type
const (
	AssessmentAR    ARType = rlib.ARASSESSMENT
	ReceiptAR       ARType = rlib.ARRECEIPT
	ExpenseAR       ARType = rlib.AREXPENSE
	SubAssessmentAR ARType = rlib.ARSUBASSESSMENT
)

// IsValid checks the validity of ARType ar
func (ar ARType) IsValid() bool {
	if ar < AssessmentAR || ar > SubAssessmentAR {
		return false
	}

	return true
}

// String representation of ARType
func (ar ARType) String() string {
	names := [...]string{
		"Assessment",
		"Receipt",
		"Expense",
		"Sub-Assessment",
	}

	if ar < AssessmentAR || ar > SubAssessmentAR {
		return "Unknown"
	}

	return names[ar]
}

// IsValidARFlag checks whether FLAGS value is valid or not
func IsValidARFlag(FLAGS uint64) bool {

	maxFLAGVal := 0
	for _, v := range ARFLAGS {
		maxFLAGVal += 1 << uint(v)
	}

	// NOTE: if no flag is set then 0 can be the case here
	if FLAGS < 0 || FLAGS > uint64(maxFLAGVal) {
		return false
	}

	// if IsRentASM and IsSecDepASM both are set
	// both should be mutually exclusive
	if FLAGS&0x20 != 0 && FLAGS&0x10 != 0 {
		return false
	}
	return true
}

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
func ValidateAcctRule(ctx context.Context, a *rlib.AR) []BizError {
	var e []BizError

	if !ARType(a.ARType).IsValid() {
		rlib.Console("*** ERROR *** invalid ARType: %d for a.ARID = %d\n", a.ARID)
		s := fmt.Sprintf(BizErrors[UnknownARType].Message, a.ARType, a.ARID)
		b := BizError{Errno: UnknownARType, Message: s}
		e = append(e, b)
	}

	if !IsValidARFlag(a.FLAGS) {
		rlib.Console("*** ERROR *** invalid FLAGS: %d for a.ARID = %d\n", a.FLAGS, a.ARID)
		s := fmt.Sprintf(BizErrors[InvalidARFlag].Message, a.FLAGS, a.ARID)
		b := BizError{Errno: InvalidARFlag, Message: s}
		e = append(e, b)
	}

	if a.CreditLID == 0 {
		rlib.Console("*** ERROR ***  a.CreditLid = %d\n", a.CreditLID)
		e = AddBizErrToList(e, BadCreditAccount)
	} else {
		// l, err := rlib.GetLedger(ctx, a.CreditLID)
		_, err := rlib.GetLedger(ctx, a.CreditLID)
		if err != nil {
			rlib.Console("rlib.GetLedger error: %s", err.Error())
			rlib.Console("*** ERROR ***  could not load a.CreditLid = %d\n", a.CreditLID)
			e = AddBizErrToList(e, BadCreditAccount)
		}

		/*if l.LID == 0 || l.BID != a.BID {
			rlib.Console("*** ERROR ***  could not load a.CreditLid = %d\n", a.CreditLID)
			e = AddBizErrToList(e, BadCreditAccount)
		}*/
	}

	rlib.Console("after credit acct check: len(e) = %d\n", len(e))
	if a.DebitLID == 0 {
		rlib.Console("*** ERROR ***  a.DebitLid = %d\n", a.DebitLID)
		e = AddBizErrToList(e, BadDebitAccount)
	} else {
		// l, err := rlib.GetLedger(ctx, a.DebitLID)
		_, err := rlib.GetLedger(ctx, a.DebitLID)
		if err != nil {
			rlib.Console("rlib.GetLedger error: %s", err.Error())
			rlib.Console("*** ERROR ***  could not load a.DebitLid = %d\n", a.DebitLID)
			e = AddBizErrToList(e, BadDebitAccount)
		}

		/*if l.LID == 0 || l.BID != a.BID {
			rlib.Console("*** ERROR ***  could not load a.DebitLid = %d\n", a.DebitLID)
			e = AddBizErrToList(e, BadDebitAccount)
		}*/
	}
	rlib.Console("after debit acct check: len(e) = %d\n", len(e))
	return e
}
