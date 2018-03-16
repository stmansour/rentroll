package bizlogic

import (
	"log"
	"rentroll/rlib"
	"strconv"
	"strings"

	"github.com/kardianos/osext"
)

// BizError is the basic structure containing an error number and a message
type BizError struct {
	Errno   int
	Message string
}

// BizErrors is the list of BizError structs that hold all known Assessment
// validation errors
var BizErrors []BizError

// RentableTypeUnknown, et al, are the error numbers for us in BizErrors
const (
	RentableTypeUnknown             = 0
	RentableStatusUnknown           = 1
	InvalidField                    = 2
	EditReversal                    = 3
	PostToSummaryAcct               = 4
	RuleUsesAcct                    = 5
	AcctHasLedgerEntries            = 6
	AcctRefInRule                   = 7
	MissingName                     = 8  // Missing required Name field
	DuplicateName                   = 9  // Duplicate Name. An item with that name already exists.
	MissingStyleName                = 10 // Style name is missing.
	DuplicateStyleName              = 11 // Duplicate Style name.  An item with that style name already exists.
	BadDebitAccount                 = 12 // The Debit account is not valid.
	BadCreditAccount                = 13 // The Credit account is not valid.
	ReceiptAlreadyDeposited         = 14 // The receipt is already a member of another deposit
	ReceiptBizMismatch              = 15 // The receipt belongs to a different business
	DepositTotalMismatch            = 16 // the total of the supplied receipts does not match
	InvalidRentableMarketRateAmount = 17 // the amount of marketRate for rentable is invalid
	InvalidRentableMRDates          = 18 // market rate dates should not overlap with other market rates
	RentableMRDatesOverlap          = 19
	RentableNameExists              = 20 // A rentable with that name already exists
	AsmDateRangeNotAllowed          = 21 // Non recur asmts must have equivalent start/stop dates
	StartDateAfterStopDate          = 22 // Stop date occurs before start date
	UnknownBID                      = 23 // Unknown BID
	UnknownBUD                      = 24 // Unknown BUD
	InvalidRentableUseStatus        = 25 // invalid rentable use status
	InvalidRentableLeaseStatus      = 26 // invalid rentable lease status
	InvalidRentableStatusDates      = 27 // invalid rentable status dates
	RentableStatusDatesOverlap      = 28 // rentable status dates overlapping
	InvalidRentableTypeRefDates     = 29 // invalid rentable type ref dates
	RentableTypeRefDatesOverlap     = 30 // rentable type ref dates overlapping
	UnknownRID                      = 31 // Unknown Rentable
	UnknownRTID                     = 32 // Unknown Rentable Type
	UnknownRAID                     = 33 // Unknown Rental Agreement
	UnknownARType                   = 34 // Unknown ARType
	UnknownTLDID                    = 35 // task list definition does not exist
	ImproperTLDID                   = 36 // task list definition does not belong to the specified business
)

// InitBizLogic loads the error messages needed for validation errors
func InitBizLogic() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fname := folderPath + "/bizerr.csv"
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		n := strings.TrimSpace(t[i][0])
		if len(n) < 0 {
			continue
		}
		j, err := strconv.Atoi(n)
		if err != nil {
			rlib.LogAndPrint("InitBizLogic: bizerr.csv - line %d, Invalid number %s\n", n)
		}
		b := BizError{Errno: j, Message: t[i][1]}
		BizErrors = append(BizErrors, b)
	}
}

// AddBizErrToList updates the supplied error list with the error message
// corresponding to errno
func AddBizErrToList(e []BizError, errno int) []BizError {
	if errno < 0 || errno+1 > len(BizErrors) {
		return e
	}
	b := BizError{Errno: errno, Message: BizErrors[errno].Message}
	e = append(e, b)
	return e
}

// AddErrToBizErrlist add a standard error to the biz errlist, and sets errno to -1
func AddErrToBizErrlist(e error, el []BizError) []BizError {
	b := BizError{Errno: -1, Message: e.Error()}
	return append(el, b)
}
