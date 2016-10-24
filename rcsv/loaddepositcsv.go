package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// CVS record format:
// 0    1         2            3      4
// BUD, Date,    DepositoryID, DepositMethodID, ReceiptSpec
// REX, 5/21/16, DEP001,       DPM01, "RCPT00001,2"

// CreateDepositsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateDepositsFromCSV(sa []string, lineno int) (string, int) {
	rs := ""
	funcname := "CreateDepositsFromCSV"
	var err error
	var d rlib.Deposit

	const (
		BUD             = 0
		Date            = iota
		DepositoryID    = iota
		DepositMethodID = iota
		ReceiptSpec     = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Date", Date},
		{"DepositoryID", DepositoryID},
		{"DepositMethodID", DepositMethodID},
		{"ReceiptSpec", ReceiptSpec},
	}

	rs, x := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if x > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	bud := strings.ToLower(strings.TrimSpace(sa[BUD]))
	if len(bud) > 0 {
		b1 := rlib.GetBusinessByDesignation(bud)
		if len(b1.Designation) == 0 {
			rs += fmt.Sprintf("%s: line %d - Business with designation %s does not exist\n", funcname, lineno, sa[BUD])
			return rs, CsvErrorSensitivity
		}
		d.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Date
	//-------------------------------------------------------------------
	d.Dt, err = rlib.StringToDate(sa[Date])
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[Date])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Depository
	//-------------------------------------------------------------------
	d.DEPID = CSVLoaderGetDEPID(sa[DepositoryID])
	if d.DEPID == 0 {
		rs += fmt.Sprintf("%s: line %d - Depository %s was not found. Skipping this item.\n", funcname, lineno, sa[DepositoryID])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Deposit Method
	//-------------------------------------------------------------------
	d.DPMID = CSVLoaderGetDPMID(sa[DepositMethodID])
	if d.DEPID == 0 {
		rs += fmt.Sprintf("%s: line %d - Deposit Method %s was not found. Skipping this item.\n", funcname, lineno, sa[DepositMethodID])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Receipts - comma separated list of RCPTIDs. Could be of the form
	// RCPT00001 or simply 1.
	//-------------------------------------------------------------------
	var rcpts []int64
	var mm []rlib.Receipt
	var tot = float64(0)

	s := strings.TrimSpace(sa[ReceiptSpec])
	ssa := strings.Split(s, ",")
	if len(ssa) == 0 {
		rs += fmt.Sprintf("%s: line %d - no receipts found. You must supply at least one receipt\n", funcname, lineno)
		return rs, CsvErrorSensitivity
	}
	for i := 0; i < len(ssa); i++ {
		id := CSVLoaderGetRCPTID(ssa[i])
		if 0 == id {
			rs += fmt.Sprintf("%s: line %d - invalid receipt number: %s\n", funcname, lineno, ssa[i])
			return rs, CsvErrorSensitivity
		}
		rcpts = append(rcpts, id)

		// load each receipt so that we can total the amount and see if it matches Amount
		rc := rlib.GetReceipt(id)
		tot += rc.Amount
		mm = append(mm, rc) // may need this later
	}
	d.Amount = tot

	//-------------------------------------------------------------------
	// We have all we need. Write the records...
	//-------------------------------------------------------------------
	id, err := rlib.InsertDeposit(&d)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d -  error inserting deposit: %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}
	for i := 0; i < len(rcpts); i++ {
		var a rlib.DepositPart
		a.DID = id
		a.RCPTID = rcpts[i]
		err = rlib.InsertDepositPart(&a)
		if nil != err {
			rs += fmt.Sprintf("%s: line %d -  error inserting deposit part: %v\n", funcname, lineno, err)
			return rs, CsvErrorSensitivity
		}
	}
	return rs, 0
}

// LoadDepositCSV loads a csv file with deposits and creates Deposit records
func LoadDepositCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateDepositsFromCSV(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
