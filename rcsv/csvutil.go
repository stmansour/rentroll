package rcsv

import (
	"context"
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

	// DupTransactant et al., are error identfiers for the CSV Loader
	DupTransactant        = "DuplicateTransactant"
	DupRentableType       = "DuplicateRentableType"
	DupCustomAttribute    = "DuplicateCustomAttribute"
	DupRentable           = "DuplicateRentable"
	RentableAlreadyRented = "RentableAlreadyRented"
)

// CsvErrorSensitivity is the error return value used by all the loadXYZcsv.go routines. We
// initialize to LOOSE as it is best for testing and should be OK for normal use as well.
var CsvErrorSensitivity = int(CsvErrLoose)

// CSVLoadHandlerFunc type of load handler function
type CSVLoadHandlerFunc func(context.Context, string) []error

// CSVLoadHandler struct is for routines that want to table-ize their loading.
type CSVLoadHandler struct {
	Fname   string
	Handler CSVLoadHandlerFunc
}

type csvHandlerFunc func(context.Context, []string, int) (int, error)

// LoadRentRollCSV performs a general purpose load.  It opens the supplied file name, and processes
// it line-by-line by calling the supplied handler function.
// Return Values
//		[]error  -  an array of errors encountered by the handler function during the load
//--------------------------------------------------------------------------------------------------
func LoadRentRollCSV(ctx context.Context, fname string, handler csvHandlerFunc) []error {
	var m []error
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if len(t[i][0]) == 0 {
			continue
		}
		if t[i][0][0] == '#' { // if it's a comment line, don't process it, just move on
			continue
		}
		s, err := handler(ctx, t[i], i+1)
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
//
// INPUTS:
// CSVColumn - array containing expected CSV column headings
// sa        - array of column headings found in input file
// funcname  - calling function name -- for error messages
// lineno    - for error messages.
//
// RETURNS:
//   bool --> false = everything is OK,  true = at least 1 column is wrong, error message already printed
//   err  --> nil if no problems
//--------------------------------------------------------------------------------------------------
func ValidateCSVColumnsErr(csvCols []CSVColumn, sa []string, funcname string, lineno int) (bool, error) {
	// rlib.Console("Entered ValidateCSVColumnsErr\n")
	required := len(csvCols)
	if len(sa) < required {
		l := len(sa)
		for i := 0; i < len(csvCols); i++ {
			if i < l {
				// s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
				sb := []byte(sa[i])
				var snew []byte
				for j := 0; j < len(sb); j++ {
					if int(sb[j]) < 128 {
						snew = append(snew, sb[j])
					}
				}
				var sorig = string(snew)
				s := rlib.Stripchars(strings.ToLower(sorig), " ")

				//--------------------------------------------------------------
				// Try to match the column name to one of the required names...
				//--------------------------------------------------------------
				s1 := strings.ToLower(csvCols[i].Name) // this needs to be expanded, it could have special characters
				// rlib.Console("getStrMatchArray(%q)\n", s1)
				m, err := getStrMatchArray(s1)
				if err != nil {
					return true, fmt.Errorf("Error = %s", err.Error())
				}
				found := false
				for j := 0; j < len(m); j++ {
					found = s == m[j]
					if found {
						break
					}
				}
				if !found {
					// rlib.Console("\n\nValidateCSVColumnsErr: heading miscompare:  %q (len = %d) with  %q (len=%d)\n", s, len(s), s1, len(s1))
					return true, fmt.Errorf("%s: line %d - Error at column heading %d, expected %q, found %q", funcname, lineno, i, csvCols[i].Name, sa[i])
				}
			}
		}
		return true, fmt.Errorf("%s: line %d - found %d values, there must be at least %d", funcname, lineno, len(sa), required)
	}

	if lineno == 1 {
		for i := 0; i < len(csvCols); i++ {
			//------------------------------------------------------------------------
			// in UTF-8 encoded files, which often come via files saved by Excel (and similar),
			// the first 3 bytes will be something like \xEF\xBB\xBF indicating byte
			// ordering. They are not ASCII chars. We need to strip out those
			// characters if they appear.
			//------------------------------------------------------------------------
			sb := []byte(strings.TrimSpace(sa[i]))
			var snew []byte
			// rlib.Console("ValidateCSVColumnsErr: M\n")
			for j := 0; j < len(sb); j++ {
				// rlib.Console("sb[%d] = %d\n", j, int(sb[j]))
				if int(sb[j]) < 128 {
					snew = append(snew, sb[j])
				}
			}
			var sorig = string(snew)
			s := rlib.Stripchars(strings.ToLower(sorig), " ")

			s1 := strings.ToLower(csvCols[i].Name)
			// rlib.Console("getStrMatchArray(%q)\n", s1)
			m, err := getStrMatchArray(s1)
			if err != nil {
				return true, fmt.Errorf("Error = %s", err.Error())
			}
			found := false
			for j := 0; j < len(m); j++ {
				found = s == m[j]
				if found {
					break
				}
			}
			if !found {
				return true, fmt.Errorf("%s: line %d - Error at column heading %d, expected %q, found %q", funcname, lineno, i, csvCols[i].Name, sorig)
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
func CSVLoaderTransactantList(ctx context.Context, BID int64, s string) ([]rlib.Transactant, error) {
	const funcname = "CSVLoaderTransactantList"

	var (
		err error
		m   []rlib.Transactant
	)

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
			err = rlib.GetTransactant(ctx, n, &a)
			if err != nil {
				rerr := fmt.Errorf("%s:  error retrieving Transactant with TCID, phone, or email: %s", funcname, s)
				return m, rerr
			}
		} else {
			a, err = rlib.GetTransactantByPhoneOrEmail(ctx, BID, s)
			if err != nil {
				rerr := fmt.Errorf("%s:  error retrieving Transactant with TCID, phone, or email: %s", funcname, s)
				return m, rerr
			}
		}
		if 0 == a.TCID {
			rerr := fmt.Errorf("%s:  error retrieving Transactant with TCID, phone, or email: %s", funcname, s)
			//fmt.Printf("%s\n", rerr.Error())
			return m, rerr
		}
		m = append(m, a)
	}
	return m, err
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
func BuildPayorList(ctx context.Context, BID int64, s string, dfltStart, dfltStop string, funcname string, lineno int) ([]rlib.RentalAgreementPayor, error) {
	var (
		m []rlib.RentalAgreementPayor
		// err error
	)
	// var noerr error
	s2 := strings.TrimSpace(s) // either the email address or the phone number
	if len(s2) == 0 {
		return m, fmt.Errorf("%s: line %d - Required Payor field is blank", funcname, lineno)
	}
	s1 := strings.Split(s2, ";")
	for i := 0; i < len(s1); i++ {
		ss := strings.Split(s1[i], ",")
		if len(ss) != 3 {
			return m, fmt.Errorf("%s: line %d - invalid Payor Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"",
				funcname, lineno, len(ss), ss)
		}
		s = strings.TrimSpace(ss[0]) // either the email address or the phone number or TransactantID (TC0003234)
		if len(s) == 0 {
			return m, fmt.Errorf("%s: line %d - Required Payor field is blank", funcname, lineno)
		}
		n, err := CSVLoaderTransactantList(ctx, BID, s)
		if err != nil {
			return m, fmt.Errorf("%s:  line %d - could not find rlib.Transactant with contact information %s", funcname, lineno, s)
		}
		if len(n) == 0 {
			return m, fmt.Errorf("%s:  line %d - could not find rlib.Transactant with contact information %s", funcname, lineno, s)
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
		payor.DtStart, payor.DtStop, _ = readTwoDates(ss[1], ss[2], funcname, lineno, "Payor")

		m = append(m, payor)
	}
	return m, nil
}

// BuildUserList parses a UserSpec and returns an array of RentableUser structs
func BuildUserList(ctx context.Context, BID int64, sa, dfltStart, dfltStop string, funcname string, lineno int) ([]rlib.RentableUser, error) {
	var (
		m []rlib.RentableUser
		// err error
	)

	s2 := strings.TrimSpace(sa) // TCID, email address, or the phone number
	if len(s2) == 0 {
		return m, fmt.Errorf("%s: line %d - Required User field is blank", funcname, lineno)
	}
	s1 := strings.Split(s2, ";")
	var noerr error
	for i := 0; i < len(s1); i++ {
		ss := strings.Split(s1[i], ",")
		if len(ss) != 3 {
			err := fmt.Errorf("%s: line %d - invalid Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"",
				funcname, lineno, len(ss), ss)
			return m, err
		}
		s := strings.TrimSpace(ss[0]) // TCID, email address, or the phone number
		if len(s) == 0 {
			return m, fmt.Errorf("%s: line %d - Required User field is blank", funcname, lineno)
		}
		n, err := CSVLoaderTransactantList(ctx, BID, s)
		if err != nil {
			return m, fmt.Errorf("%s: line %d - invalid person identifier: %s. Error = %s", funcname, lineno, s, err.Error())
		}
		var p rlib.RentableUser
		p.TCID = n[0].TCID

		if len(strings.TrimSpace(ss[1])) == 0 {
			ss[1] = dfltStart
		}
		if len(strings.TrimSpace(ss[2])) == 0 {
			ss[2] = dfltStop
		}
		p.DtStart, p.DtStop, _ = readTwoDates(ss[1], ss[2], funcname, lineno, "User")
		m = append(m, p)
	}
	return m, noerr
}

// // BID is the business id of the business unit to which the people belong
// func x(BID int64) {
// 	rows, err := rlib.RRdb.Prepstmt.GetAllTransactantsForBID.Query(BID)
// 	rlib.Errcheck(err)
// 	defer rows.Close()
// 	for rows.Next() {
// 		var tr rlib.Transactant
// 		rlib.ReadTransactants(rows, &tr)
// 		// Now dow whatever you need to do with the information in the transactant tr
// 	}
// 	rlib.Errcheck(rows.Err())
// }
