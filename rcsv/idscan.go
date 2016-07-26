package rcsv

import (
	"fmt"
	"regexp"
	"rentroll/rlib"
	"strings"
)

func readNumFromExpr(sa, expr, errmsg string) int64 {
	s := strings.TrimSpace(sa)
	re, _ := regexp.Compile(expr)
	m := re.FindStringSubmatch(s) // returns this pattern:  ["xxx0000001" "2"]
	if len(m) > 0 {               // if the prefix was "DEP", m will have 2 elements, our number should be the second element
		s = m[1]
	}
	s2 := errmsg + " is invalid"
	id, _ := rlib.IntFromString(s, s2)
	return id
}

// CSVLoaderGetDEPID parses a string of the form DEP000000321 and returns the DEPID , in this case 321.
func CSVLoaderGetDEPID(sa string) int64 {
	return readNumFromExpr(sa, "^DEP0*(.*)", "DEPID")
}

// CSVLoaderGetDPMID parses a string of the form DPM000000321 and returns the DPMID , in this case 321.
func CSVLoaderGetDPMID(sa string) int64 {
	return readNumFromExpr(sa, "^DPM0*(.*)", "DPMID")
}

// CSVLoaderGetRCPTID parses a string of the form RCPT000000321 and returns the RCPTID , in this case 321.
func CSVLoaderGetRCPTID(sa string) int64 {
	return readNumFromExpr(sa, "^RCPT0*(.*)", "RCPTID")
}

// CSVLoaderGetRAID parses a string of the form RA000000321 and returns the RAID , in this case 321.
func CSVLoaderGetRAID(sa string) int64 {
	return readNumFromExpr(sa, "^RA0*(.*)", "Rental Agreement")
}

// CSVLoaderGetASMID parses a string of the form ASM000000321 and returns the ASMID , in this case 321.
func CSVLoaderGetASMID(sa string) int64 {
	return readNumFromExpr(sa, "^ASM0*(.*)", "ASMID")
}

// CSVLoaderGetInvoiceNo parses a string of the form IN000000321 and returns the InvoiceNo , in this case 321.
func CSVLoaderGetInvoiceNo(sa string) int64 {
	return readNumFromExpr(sa, "^IN0*(.*)", "InvoiceNo")
}

// CSVLoaderTransactantList takes a comma separated list of email addresses and phone numbers
// and returns an array of transactants for each.  If any of the addresses in the list
// cannot be resolved to a rlib.Transactant, then processing stops immediately and an error is returned.
func CSVLoaderTransactantList(s string) ([]rlib.Transactant, error) {
	funcname := "CSVLoaderTransactantList"
	var m []rlib.Transactant
	var noerr error
	s2 := strings.TrimSpace(s) // either the email address or the phone number
	ss := strings.Split(s2, ",")
	for i := 0; i < len(ss); i++ {
		s = strings.TrimSpace(ss[i]) // either the email address or the phone number
		t, err := rlib.GetTransactantByPhoneOrEmail(s)
		if err != nil && !rlib.IsSQLNoResultsError(err) {
			rerr := fmt.Errorf("%s:  error retrieving Transactant by phone or email: %v", funcname, err)
			rlib.Ulog("%s", rerr.Error())
			return m, rerr
		}
		if t.PID == 0 {
			rerr := fmt.Errorf("%s:  could not find Transactant with contact information %s\n", funcname, s)
			rlib.Ulog("%s", rerr.Error())
			return m, rerr
		}

		m = append(m, t)
	}
	return m, noerr
}
