package rcsv

import (
	"context"
	"fmt"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"strings"
)

// CVS record format:
// 0    1         2            3      4
// BUD, Date,    DepositoryID, DepositMethodID, ReceiptSpec
// REX, 5/21/16, DEP001,       DPM01, "RCPT00001,2"

// CreateDepositsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateDepositsFromCSV(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateDepositsFromCSV"
	var (
		err error
		d   rlib.Deposit
	)

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

	//rlib.Console("%s: processing input: %#v\n", funcname, sa)

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	bud := strings.ToLower(strings.TrimSpace(sa[BUD]))
	if len(bud) > 0 {
		b1, err := rlib.GetBusinessByDesignation(ctx, bud)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Business with designation %s does not exist", funcname, lineno, sa[BUD])
		}
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Business with designation %s does not exist", funcname, lineno, sa[BUD])
		}
		d.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Date
	//-------------------------------------------------------------------
	d.Dt, err = rlib.StringToDate(sa[Date])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid start date:  %s", funcname, lineno, sa[Date])
	}

	//-------------------------------------------------------------------
	// Depository
	//-------------------------------------------------------------------
	d.DEPID = CSVLoaderGetDEPID(sa[DepositoryID])
	if d.DEPID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Skipping because Depository %s was not found", funcname, lineno, sa[DepositoryID])
	}

	//-------------------------------------------------------------------
	// Deposit Method
	//-------------------------------------------------------------------
	d.DPMID = CSVLoaderGetDPMID(sa[DepositMethodID])
	if d.DEPID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Skipping because Deposit Method %s was not found", funcname, lineno, sa[DepositMethodID])
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
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - no receipts found. You must supply at least one receipt", funcname, lineno)
	}
	for i := 0; i < len(ssa); i++ {
		//rlib.Console("%d. %s\n", i, ssa[i])
		id := CSVLoaderGetRCPTID(ssa[i])
		if 0 == id {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid receipt number: %s", funcname, lineno, ssa[i])
		}
		rcpts = append(rcpts, id)

		// load each receipt so that we can total the amount and see if it matches Amount
		rc, err := rlib.GetReceipt(ctx, id)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error while getting receipt number: %s, error: %s", funcname, lineno, ssa[i], err.Error())
		}
		tot += rc.Amount
		mm = append(mm, rc) // may need this later
	}
	d.Amount = tot

	//-------------------------------------------------------------------
	// We have all we need. Write the records...
	//-------------------------------------------------------------------
	//rlib.Console("CreateDepositsFromCSV:  deposit.Amount = %8.2f\n", d.Amount)

	// id, err := rlib.InsertDeposit(&d)
	// if err != nil {
	// 	return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error inserting deposit: %v", funcname, lineno, err)
	// }
	// for i := 0; i < len(rcpts); i++ {
	// 	//rlib.Console("Receipt Parts: %d. %d\n", i, rcpts[i])

	// 	var a rlib.DepositPart
	// 	a.DID = id
	// 	a.BID = d.BID
	// 	a.RCPTID = rcpts[i]
	// 	_, err = rlib.InsertDepositPart(&a)
	// 	if nil != err {
	// 		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error inserting deposit part: %v", funcname, lineno, err)
	// 	}
	// }
	errlist := bizlogic.SaveDeposit(ctx, &d, rcpts)
	if len(errlist) > 0 {
		srr := ""
		for i := 0; i < len(errlist); i++ {
			srr += errlist[i].Message + "\n"
		}
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error saving deposit: %s", funcname, lineno, srr)
	}

	return 0, nil
}

// LoadDepositCSV loads a csv file with deposits and creates Deposit records
func LoadDepositCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateDepositsFromCSV)
}
