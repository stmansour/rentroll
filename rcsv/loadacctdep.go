package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// 0   1   2
// Bud,LID,DEPID
// REX,1,1
// REX,2,1

// CreateAccountDepository creates AccountDepository database records from the supplied CSV file lines
func CreateAccountDepository(sa []string, lineno int) (int, error) {
	funcname := "CreateAccountDepository"
	var b rlib.AccountDepository
	var err error

	const (
		BUD   = 0
		LID   = iota
		DEPID = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"LID", LID},
		{"DEPID", DEPID},
	}

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
	des := strings.TrimSpace(sa[BUD])
	if len(des) > 0 { // make sure it's not empty
		b1 := rlib.GetBusinessByDesignation(des) // see if we can find the biz
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		b.BID = b1.BID
	}

	//-----------------------------------------
	// Get the account
	//-----------------------------------------
	b.LID, err = rlib.IntFromString(sa[LID], "Invalid LID")
	if err != nil || b.LID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid GLAccountID: %s", funcname, lineno, sa[LID])
	}

	//-----------------------------------------
	// Get the Depository
	//-----------------------------------------
	b.DEPID, err = rlib.IntFromString(sa[DEPID], "Invalid DEPID")
	if err != nil || b.DEPID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid Depository ID: %s", funcname, lineno, sa[DEPID])
	}

	//----------------------------------------------------------------
	// Validate that the referenced account and depository exist...
	//----------------------------------------------------------------
	gl := rlib.GetLedger(b.LID)
	if gl.LID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No GL Account with ID = %d", funcname, lineno, b.LID)
	}

	dep, err := rlib.GetDepository(b.DEPID)
	if err != nil || dep.DEPID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No Depository with DEPID  = %d", funcname, lineno, b.DEPID)
	}

	_, err = rlib.InsertAccountDepository(&b)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: error inserting AccountDepository = %v", funcname, err)
	}

	return 0, nil
}

// LoadAccountDepositoryCSV loads the values from the supplied csv file and creates AccountDepository records.
func LoadAccountDepositoryCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateAccountDepository)
}
