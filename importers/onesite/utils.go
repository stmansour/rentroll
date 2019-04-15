package onesite

import (
	"context"
	"fmt"
	"rentroll/importers/core"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// IsValidRentableUseType checks that passed string contains valid rentable use
// type in then rentroll system
//
// INPUTS:
//   s  = string found in onesite csv
//
// RETURNS:
//    bool = whether we found the string or not
//  string = rentRollStatus key string
//  string = value string
//------------------------------------------------------------------------------
func IsValidRentableUseType(s string) (bool, string, string) {
	found := false
	var tempRS, rentRollStatus string
	//--------------------------------------------------------
	// first find that passed string contains any status key
	//--------------------------------------------------------
	a := strings.ToLower(s)
	for k, v := range RentableUseTypeCSV {
		if strings.Contains(a, k) {
			tempRS = v
			rentRollStatus = k
			found = true
			break
		}
	}
	return found, rentRollStatus, tempRS
}

// csvRecordsToSkip function that should check an error
// which contains such a thing that needs to be discard
// such as. already exists, already done. etc. . . .
func csvRecordsToSkip(err error) bool {
	for _, dup := range csvRecordsSkipList {
		if strings.Contains(err.Error(), dup) {
			return true
		}
	}
	return false
}

// TO PARSE LINE, ERROR TEXT FROM RCSV ERRORS ONLY
func parseLineAndErrorFromRCSV(rcsvErr error, dbType int) (int, string, bool) {
	/*
		This parsing is only works with below pattern
		========================================
		{FunctionName}: line {LineNumber} - errorReason
		========================================
		if other pattern supplied for error then it fails
	*/

	errText := rcsvErr.Error()
	// split with separator `:` breaks into [0]{FuncName} and [1]rest of the text
	// split at most 2 substrings only
	s := strings.SplitN(errText, ":", 2)
	// we need only text without {FuncName}
	errText = s[1]
	// split with separator `-` breaks into [0] line no string and [1] actual reason for error which we want to show to user
	// split at most 2 substrings only
	s = strings.SplitN(errText, "-", 2)

	// parse error reason =================
	// now we only need the exact reason
	errText = strings.TrimSpace(s[1])
	// remove new line broker
	errText = strings.Replace(errText, "\n", "", -1)
	// consider this as Errors so need to prepand <E:>
	errText = "E:<" + core.DBTypeMapStrings[dbType] + ">:" + errText

	// parse line number =================
	// get line number string
	lineNoStr := s[0]
	// remove `line` text from lineNoStr string
	lineNoStr = strings.Replace(lineNoStr, "line", "", -1)
	// remove space from lineNoStr string
	lineNoStr = strings.TrimSpace(lineNoStr)
	// now it should contain number in string
	lineNo, err := strconv.Atoi(lineNoStr)
	if err != nil {
		// CRITICAL
		rlib.Ulog("INTERNAL ERRORS: RCSV Error is not in format of `{FunctionName}: line {LineNumber} - errorReason` for error: %s", errText)
		return lineNo, errText, false
	}
	//return
	return lineNo, errText, true
}

// ValidateUserSuppliedValues validates all user supplied values
// return error list and also business unit
func ValidateUserSuppliedValues(ctx context.Context, userValues map[string]string) ([]error, *rlib.Business) {
	var errorList []error
	var accrualRateOptText = `| 0: one time only | 1: secondly | 2: minutely | 3: hourly | 4: daily | 5: weekly | 6: monthly | 7: quarterly | 8: yearly |`

	// --------------------- BUD validation ------------------------
	BUD := userValues["BUD"]
	business, err := rlib.GetBusinessByDesignation(ctx, BUD)
	if err != nil {
		errorList = append(errorList,
			fmt.Errorf("Supplied Business Unit Designation does not exists"))
	}
	// resource not found then consider it's as an error
	if business.BID == 0 {
		errorList = append(errorList,
			fmt.Errorf("Supplied Business Unit Designation does not exists"))
	}

	// --------------------- RentCycle validation ------------------------
	RentCycle, err := strconv.Atoi(userValues["RentCycle"])
	if err != nil || RentCycle < 0 || RentCycle > 8 {
		errorList = append(errorList,
			fmt.Errorf("Please, choose Frequency value from this\n%s", accrualRateOptText))
	}

	// --------------------- Proration validation ------------------------
	Proration, err := strconv.Atoi(userValues["Proration"])
	if err != nil || Proration < 0 || Proration > 8 {
		errorList = append(errorList,
			fmt.Errorf("Please, choose Proration value from this\n%s", accrualRateOptText))
	}

	// --------------------- GSRPC validation ------------------------
	GSRPC, err := strconv.Atoi(userValues["GSRPC"])
	if err != nil || GSRPC < 0 || GSRPC > 8 {
		errorList = append(errorList,
			fmt.Errorf("Please, choose GSRPC value from this\n%s", accrualRateOptText))
	}

	// finally return error list
	return errorList, &business
}

// take int of csv index, current time in string format
func getPeopleNoteString(rowIndex int, currentTime string) string {
	return onesiteNotesPrefix + currentTime + "$" + strconv.Itoa(rowIndex)
}
