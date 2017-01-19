package onesite

import (
	"rentroll/rlib"
	"strconv"
	"strings"
)

// IsValidRentableStatus checks that passed string contains valid rentable status
// acoording to rentroll system
func IsValidRentableStatus(s string) (bool, string, string) {
	found := false
	var tempRS, rentRollStatus string
	// first find that passed string contains any status key
	a := strings.ToLower(s)
	for k, v := range RentableStatusCSV {
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
func parseLineAndErrorFromRCSV(rcsvErr error) (int, string, bool) {
	/*
		This parsing is only works with below pattern
		========================================
		{FunctionName}: line {LineNumber} - errorReason
		========================================
		if other pattern supplied for error then it fails
	*/

	errText := rcsvErr.Error()
	// split with separator `:` breaks into [0]{FuncName} and [1]rest of the text
	s := strings.Split(errText, ":")
	// we need only text without {FuncName}
	errText = s[1]
	// split with separator `-` breaks into [0] line no string and [1] actual reason for error which we want to show to user
	s = strings.Split(errText, "-")

	// parse error reason =================
	// now we only need the exact reason
	errText = strings.TrimSpace(s[1])
	// remove new line broker
	errText = strings.Replace(errText, "\n", "", -1)

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
		rlib.Ulog("rcsv loaders should do something about returning error format")
		return lineNo, errText, false
	}
	//return
	return lineNo, errText, true
}
