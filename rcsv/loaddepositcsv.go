package rcsv

import (
	"fmt"
	"regexp"
	"rentroll/rlib"
	"strings"
)

// CVS record format:
// 0    1         2     3      4
// BUD, Date,    DEPID, Amount,Receipts
// REX, 5/21/16, DEP001,2000,  "RCPT00001,2"

// CSVLoaderGetDEPID parses a string of the for DEP000000321 and returns the DEPID , in this case 321.
func CSVLoaderGetDEPID(sa string) int64 {
	s := strings.TrimSpace(sa)
	re, _ := regexp.Compile("^DEP0*(.*)")
	m := re.FindStringSubmatch(s) // returns this pattern:  ["DEP0000001" "2"]
	if len(m) > 0 {               // if the prefix was "DEP", m will have 2 elements, our number should be the second element
		s = m[1]
	}
	id, _ := rlib.IntFromString(s, "DEPID number is invalid")
	return id
}

// CSVLoaderGetRCPTID parses a string of the for RCPT000000321 and returns the RCPTID , in this case 321.
func CSVLoaderGetRCPTID(sa string) int64 {
	s := strings.TrimSpace(sa)
	re, _ := regexp.Compile("^RCPT0*(.*)")
	m := re.FindStringSubmatch(s) // returns this pattern:  ["RCPT0000001" "2"]
	if len(m) > 0 {               // if the prefix was "DEP", m will have 2 elements, our number should be the second element
		s = m[1]
	}
	id, _ := rlib.IntFromString(s, "RCPTID number is invalid")
	return id
}

// CreateDepositsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateDepositsFromCSV(sa []string, lineno int) {
	funcname := "CreateDepositsFromCSV"
	var err error
	var d rlib.Deposit

	bud := strings.ToLower(strings.TrimSpace(sa[0]))
	if bud == "bud" {
		return // this is just the column heading
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 5
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(bud) > 0 {
		b1, _ := rlib.GetBusinessByDesignation(bud)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d - Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		d.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Date
	//-------------------------------------------------------------------
	d.Dt, err = StringToDate(sa[1])
	if err != nil {
		fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[1])
		return
	}

	//-------------------------------------------------------------------
	// Depository
	//-------------------------------------------------------------------
	d.DEPID = CSVLoaderGetDEPID(sa[2])
	if d.DEPID == 0 {
		rlib.Ulog("%s: line %d - Depository %s was not found. Skipping this item.\n", funcname, lineno, sa[1])
		return
	}

	//-------------------------------------------------------------------
	// Amount
	//-------------------------------------------------------------------
	d.Amount, _ = rlib.FloatFromString(sa[3], "deposit Amount is invalid")

	//-------------------------------------------------------------------
	// Receipts - comma separated list of RCPTIDs. Could be of the form
	// RCPT00001 or simply 1.
	//-------------------------------------------------------------------
	var rcpts []int64
	s := strings.TrimSpace(sa[4])
	ssa := strings.Split(s, ",")
	if len(ssa) == 0 {
		rlib.Ulog("%s: line %d - no receipts found. You must supply at least one receipt\n", funcname, lineno)
		return
	}
	for i := 0; i < len(ssa); i++ {
		id := CSVLoaderGetRCPTID(ssa[i])
		if 0 == id {
			rlib.Ulog("%s: line %d - invalid receipt number: %s\n", funcname, lineno, ssa[i])
			return
		}
		rcpts = append(rcpts, id)
	}

	//-------------------------------------------------------------------
	// We have all we need. Write the records...
	//-------------------------------------------------------------------
	id, err := rlib.InsertDeposit(&d)
	if err != nil {
		fmt.Printf("%s: line %d -  error inserting deposit: %v\n", funcname, lineno, err)
		return
	}
	for i := 0; i < len(rcpts); i++ {
		var a rlib.DepositPart
		a.DID = id
		a.RCPTID = rcpts[i]
		err = rlib.InsertDepositPart(&a)
		if nil != err {
			fmt.Printf("%s: line %d -  error inserting deposit part: %v\n", funcname, lineno, err)
			return
		}
	}
}

// LoadDepositCSV loads a csv file with deposits and creates Deposit records
func LoadDepositCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateDepositsFromCSV(t[i], i+1)
	}
}
