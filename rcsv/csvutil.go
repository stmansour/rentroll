package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"
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

// CSVLoadHandler struct is for routines that want to table-ize their loading.
type CSVLoadHandler struct {
	Fname   string
	Handler func(string) []error
}

// CSVReporterInfo is for routines that want to table-ize their reporting using
// the CSV library's simple report routines.
type CSVReporterInfo struct {
	ReportNo     int       // index number of the report
	OutputFormat int       // text, html, maybe more in the future
	Bid          int64     // associated business
	Raid         int64     // associated Rental Agreement if needed
	D1           time.Time // associated date if needed
	D2           time.Time // associated date if needed
	NeedsBID     bool      // true if BID is needed for this report
	NeedsRAID    bool      // true if RAID is needed for this report
	NeedsDt      bool      // true if a Date is needed for this report
	Handler      func(*CSVReporterInfo) string
	Xbiz         *rlib.XBusiness // may not be set in all cases
}

// LoadRentRollCSV performs a general purpose load.  It opens the supplied file name, and processes
// it line-by-line by calling the supplied handler function.
// Return Values
//		[]error  -  an array of errors encountered by the handler function during the load
//--------------------------------------------------------------------------------------------------
func LoadRentRollCSV(fname string, handler func([]string, int) (int, error)) []error {
	var m []error
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if t[i][0] == "#" { // if it's a comment line, don't process it, just move on
			continue
		}
		s, err := handler(t[i], i+1)
		if err != nil {
			m = append(m, err)
		}
		if s > 0 { // if handler indicates that we need to stop...
			break //... then exit out of the loop
		}
	}
	return m
}

// ValidateCSVColumnsErr verifies the column titles with the supplied, expected titles.
// Returns:
//   bool --> false = everything is OK,  true = at least 1 column is wrong, error message already printed
//   err  --> nil if no problems
func ValidateCSVColumnsErr(csvCols []CSVColumn, sa []string, funcname string, lineno int) (bool, error) {
	required := len(csvCols)
	if len(sa) < required {
		l := len(sa)
		for i := 0; i < len(csvCols); i++ {
			if i < l {
				s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
				if s != strings.ToLower(csvCols[i].Name) {
					return true, fmt.Errorf("%s: line %d - Error at column heading %d, expected %s, found %s\n", funcname, lineno, i, csvCols[i].Name, sa[i])
				}
			}
		}
		return true, fmt.Errorf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
	}

	if lineno == 1 {
		for i := 0; i < len(csvCols); i++ {
			s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
			if s != strings.ToLower(csvCols[i].Name) {
				return true, fmt.Errorf("%s: line %d - Error at column heading %d, expected %s, found %s\n", funcname, lineno, i, csvCols[i].Name, sa[i])
			}
		}
	}
	return false, nil
}

// ValidateCSVColumns wrapper for ValidateCSVColumnsErr
func ValidateCSVColumns(csvCols []CSVColumn, sa []string, funcname string, lineno int) (int, error) {
	t := 0
	b, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if b {
		t = 1
	}
	return t, err
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
			//fmt.Printf("%s\n", rerr.Error())
			return m, rerr
		}
		m = append(m, a)
	}
	return m, noerr
}

// ErrlistToString converts an errorlist into a string suitable for printout
func ErrlistToString(m *[]error) string {
	rs := ""
	for i := 0; i < len(*m); i++ {
		s := (*m)[i].Error()
		if s[len(s)-1:] != "\n" {
			s += "\n"
		}
		rs += s
	}
	return rs
}

// BuildPayorList takes a semi-colon separated list of email addresses and phone numbers
// and returns an array of rlib.RentalAgreementPayor records for each.  If any of the addresses in the list
// cannot be resolved to a rlib.Transactant, then processing stops immediately and an error is returned.
// Each value is time sensitive (has an associated time range). If the dates are not specified, then the
// default values of dfltStart and dfltStop -- which are the start/stop time of the rental agreement --
// are used instead. This is common because the payors will usually be the same for the entire rental
// agreement lifetime.
func BuildPayorList(BID int64, s string, dfltStart, dfltStop string, funcname string, lineno int) ([]rlib.RentalAgreementPayor, error) {
	var m []rlib.RentalAgreementPayor
	// var noerr error
	s2 := strings.TrimSpace(s) // either the email address or the phone number
	if len(s2) == 0 {
		return m, fmt.Errorf("%s: lineno %d - Required Payor field is blank\n", funcname, lineno)
	}
	s1 := strings.Split(s2, ";")
	for i := 0; i < len(s1); i++ {
		ss := strings.Split(s1[i], ",")
		if len(ss) != 3 {
			return m, fmt.Errorf("%s: lineno %d - invalid Payor Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
				funcname, lineno, len(ss), ss)
		}
		s = strings.TrimSpace(ss[0]) // either the email address or the phone number or TransactantID (TC0003234)
		if len(s) == 0 {
			return m, fmt.Errorf("%s: lineno %d - Required Payor field is blank\n", funcname, lineno)
		}
		n, _ := CSVLoaderTransactantList(BID, s)
		if len(n) == 0 {
			return m, fmt.Errorf("%s:  lineno %d - could not find rlib.Transactant with contact information %s\n", funcname, lineno, s)
		}

		var payor rlib.RentalAgreementPayor
		payor.TCID = n[0].TCID

		// Now grab the dates
		if len(strings.TrimSpace(ss[1])) == 0 {
			ss[1] = dfltStart
		}
		if len(strings.TrimSpace(ss[2])) == 0 {
			ss[2] = dfltStop
		}
		payor.DtStart, payor.DtStop, _ = readTwoDates(ss[1], ss[2], funcname, lineno)

		m = append(m, payor)
	}
	return m, nil
}

// BuildUserList parses a UserSpec and returns an array of RentableUser structs
func BuildUserList(BID int64, sa, dfltStart, dfltStop string, funcname string, lineno int) ([]rlib.RentableUser, error) {
	var m []rlib.RentableUser
	s2 := strings.TrimSpace(sa) // TCID, email address, or the phone number
	if len(s2) == 0 {
		return m, fmt.Errorf("%s: lineno %d - Required User field is blank\n", funcname, lineno)
	}
	s1 := strings.Split(s2, ";")
	var noerr error
	for i := 0; i < len(s1); i++ {
		ss := strings.Split(s1[i], ",")
		if len(ss) != 3 {
			err := fmt.Errorf("%s: lineno %d - invalid Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
				funcname, lineno, len(ss), ss)
			return m, err
		}
		s := strings.TrimSpace(ss[0]) // TCID, email address, or the phone number
		if len(s) == 0 {
			return m, fmt.Errorf("%s: lineno %d - Required User field is blank\n", funcname, lineno)
		}
		n, err := CSVLoaderTransactantList(BID, s)
		if err != nil {
			return m, fmt.Errorf("%s: lineno %d - invalid person identifier: %s. Error = %s\n", funcname, lineno, s, err.Error())
		}
		var p rlib.RentableUser
		p.TCID = n[0].TCID

		if len(strings.TrimSpace(ss[1])) == 0 {
			ss[1] = dfltStart
		}
		if len(strings.TrimSpace(ss[2])) == 0 {
			ss[2] = dfltStop
		}
		p.DtStart, p.DtStop, _ = readTwoDates(ss[1], ss[2], funcname, lineno)
		m = append(m, p)
	}
	return m, noerr
}
