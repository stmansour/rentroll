package rcsv

import (
	"regexp"
	"rentroll/rlib"
	"strings"
)

func readNumAndStatusFromExpr(sa, expr, errmsg string) (int64, string) {
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
	rs := ""
	v, err := rlib.IntFromString(s, s2)
	if err != nil {
		rs = err.Error()
	}
	return v, rs
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
