package rcsv

import (
	"fmt"
	"regexp"
	"rentroll/rlib"
	"strings"
)

func readNumAndStatusFromExpr(sa, expr, errmsg string) (int64, bool) {
	s := strings.TrimSpace(sa)
	re, _ := regexp.Compile(expr)
	m := re.FindStringSubmatch(s) // returns this pattern:  ["xxx0000001" "2"]
	if len(m) > 0 {               // if the prefix was "DEP", m will have 2 elements, our number should be the second element
		s = m[1]
	}
	s2 := ""
	if "" != errmsg {
		s2 = errmsg + " is invalid"
	}
	return rlib.IntFromString(s, s2)
}

func readNumFromExpr(sa, expr, errmsg string) int64 {
	n, _ := readNumAndStatusFromExpr(sa, expr, errmsg)
	return n
}

// CSVLoaderGetASMID parses a string of the form ASM000000321 and returns the ASMID , in this case 321.
func CSVLoaderGetASMID(sa string) int64 {
	return readNumFromExpr(sa, "^ASM0*(.*)", "ASMID")
}

// CSVLoaderGetDEPID parses a string of the form DEP000000321 and returns the DEPID , in this case 321.
func CSVLoaderGetDEPID(sa string) int64 {
	return readNumFromExpr(sa, "^DEP0*(.*)", "DEPID")
}

// CSVLoaderGetDPMID parses a string of the form DPM000000321 and returns the DPMID , in this case 321.
func CSVLoaderGetDPMID(sa string) int64 {
	return readNumFromExpr(sa, "^DPM0*(.*)", "DPMID")
}

// CSVLoaderGetInvoiceNo parses a string of the form IN000000321 and returns the InvoiceNo , in this case 321.
func CSVLoaderGetInvoiceNo(sa string) int64 {
	return readNumFromExpr(sa, "^IN0*(.*)", "InvoiceNo")
}

// CSVLoaderGetLedgerNo parses a string of the form L000000321 and returns the Ledger , in this case 321.
func CSVLoaderGetLedgerNo(sa string) int64 {
	return readNumFromExpr(sa, "^L0*(.*)", "Ledger")
}

// CSVLoaderGetRAID parses a string of the form RA000000321 and returns the RAID , in this case 321.
func CSVLoaderGetRAID(sa string) int64 {
	return readNumFromExpr(sa, "^RA0*(.*)", "Rental Agreement")
}

// CSVLoaderGetRID parses a string of the form R000000321 and returns the RID , in this case 321.
func CSVLoaderGetRID(sa string) int64 {
	return readNumFromExpr(sa, "^R0*(.*)", "Rentable")
}

// CSVLoaderGetRPID parses a string of the form RP000000321 and returns the RPID , in this case 321.
func CSVLoaderGetRPID(sa string) int64 {
	return readNumFromExpr(sa, "^RP0*(.*)", "Rate Plan")
}

// CSVLoaderGetRPRID parses a string of the form RPR000000321 and returns the RPRID , in this case 321.
func CSVLoaderGetRPRID(sa string) int64 {
	return readNumFromExpr(sa, "^RPR0*(.*)", "Rate Plan Ref")
}

// CSVLoaderGetRCPTID parses a string of the form RCPT000000321 and returns the RCPTID , in this case 321.
func CSVLoaderGetRCPTID(sa string) int64 {
	return readNumFromExpr(sa, "^RCPT0*(.*)", "RCPTID")
}

// CSVLoaderGetTCID parses a string of the form TC000000321 and returns the TCID , in this case 321.
func CSVLoaderGetTCID(sa string) int64 {
	return readNumFromExpr(sa, "^TC0*(.*)", "TCID")
}

// CSVLoaderTransactantList takes a comma separated list of email addresses and phone numbers
// and returns an array of transactants for each.  If any of the addresses in the list
// cannot be resolved to a rlib.Transactant, then processing stops immediately and an error is returned.
func CSVLoaderTransactantList(s string) ([]rlib.Transactant, error) {
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
		if ok {
			rlib.GetTransactant(n, &a)
		} else {
			a = rlib.GetTransactantByPhoneOrEmail(s)
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
