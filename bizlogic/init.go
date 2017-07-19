package bizlogic

import (
	"rentroll/rlib"
	"strconv"
	"strings"
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
	RentableTypeUnknown   = 0
	RentableStatusUnknown = 1
	InvalidField          = 2
	EditReversal          = 3
)

// InitBizLogic loads the error messages needed for validation errors
func InitBizLogic() {
	t := rlib.LoadCSV("bizerr.csv")
	for i := 0; i < len(t); i++ {
		n := strings.TrimSpace(t[i][0])
		if len(n) < 0 {
			continue
		}
		j, err := strconv.Atoi(n)
		if err != nil {
			rlib.LogAndPrint("InitAssessmentValidation: bizerr.csv - line %d, Invalid number %s\n", n)
		}
		b := BizError{Errno: j, Message: t[i][1]}
		BizErrors = append(BizErrors, b)
	}
}
