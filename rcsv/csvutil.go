package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// CSVColumn defines a column of the CSV file
type CSVColumn struct {
	Name  string
	Index int
}

// CsvErrLoose et al, are constants used to control whether an error on a single line causes
// the entire CSV process to terminate or continue.   If LOOSE, then it will skip the error line
// and continue to process the remaining lines.  If STRICT, then the entire CSV loading process
// will terminate if any error is encountered
const (
	CsvErrLoose  = 0
	CsvErrStrict = 1
)

// CsvErrorSensitivity is the error return value used by all the loadXYZcsv.go routines. We
// initialize to LOOSE as it is best for testing and should be OK for normal use as well.
var CsvErrorSensitivity = int(CsvErrLoose)

// ValidateCSVColumns verifies the column titles with the supplied, expected titles.
// Returns:
//   0 = everything is OK
//   1 = at least 1 column is wrong, error message already printed
func ValidateCSVColumns(csvCols []CSVColumn, sa []string, funcname string, lineno int) (string, int) {
	rs := ""
	required := len(csvCols)
	if len(sa) < required {
		rs += fmt.Sprintf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		l := len(sa)
		for i := 0; i < len(csvCols); i++ {
			if i < l {
				s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
				if s != strings.ToLower(csvCols[i].Name) {
					rs += fmt.Sprintf("%s: line %d - Error at column heading %d, expected %s, found %s\n", funcname, lineno, i, csvCols[i].Name, sa[i])
					return rs, 1
				}
			}
		}
		return rs, 1
	}

	if lineno == 1 {
		for i := 0; i < len(csvCols); i++ {
			s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
			if s != strings.ToLower(csvCols[i].Name) {
				rs += fmt.Sprintf("%s: line %d - Error at column heading %d, expected %s, found %s\n", funcname, lineno, i, csvCols[i].Name, sa[i])
				return rs, 1
			}
		}
	}
	return rs, 0
}

// CSVLoaderTransactantList takes a comma separated list of email addresses and phone numbers
// and returns an array of transactants for each.  If any of the addresses in the list
// cannot be resolved to a rlib.Transactant, then processing stops immediately and an error is returned.
func CSVLoaderTransactantList(BID int64, s string) ([]rlib.Transactant, error) {
	funcname := "CSVLoaderTransactantList"
	var m []rlib.Transactant
	var noerr error
	if "" == s {
		return m, nil
	}
	s2 := strings.TrimSpace(s) // either the email address or the phone number
	ss := strings.Split(s2, ",")
	for i := 0; i < len(ss); i++ {
		var a rlib.Transactant
		s = strings.TrimSpace(ss[i])                          // either the email address or the phone number
		n, ok := readNumAndStatusFromExpr(s, "^TC0*(.*)", "") // "" suppresses error messages
		if len(ok) == 0 {
			rlib.GetTransactant(n, &a)
		} else {
			a = rlib.GetTransactantByPhoneOrEmail(BID, s)
		}
		if 0 == a.TCID {
			rerr := fmt.Errorf("%s:  error retrieving Transactant with TCID, phone, or email: %s", funcname, s)
			rlib.Ulog("%s", rerr.Error())
			return m, rerr
		}
		m = append(m, a)
	}
	return m, noerr
}
