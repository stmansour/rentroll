package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//                     Ledger NAME or ID works
// 0   1   2           3
// Bud,Name,ARType,    DebitLID,CreditLID,Description
// REX,Rent,Assessment,2,       8,        Rent assessment, accrual based, manage to budget
// REX,FNB, Receipt,   3,       7,        payments that are deposited in First National Bank

// CreateAR creates AR database records from the supplied CSV file lines
func CreateAR(sa []string, lineno int) (int, error) {
	funcname := "CreateAR"
	var b rlib.AR
	var err error

	const (
		BUD         = 0
		Name        = iota
		ARType      = iota
		DebitLID    = iota
		CreditLID   = iota
		Description = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"ARType", ARType},
		{"DebitLID", DebitLID},
		{"CreditLID", CreditLID},
		{"Description", Description},
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
	// Get the name
	//-----------------------------------------
	b.Name = sa[Name]
	b2, err := rlib.GetARByName(b.BID, b.Name)
	if err != nil && !rlib.IsSQLNoResultsError(err) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error attempting to read existing records with name = %s: %s", funcname, lineno, b.Name, err.Error())
	}
	if b2.Name == b.Name {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - An AR rule with name = %s already exists. Ignoring this line", funcname, lineno, b.Name)
	}

	//-----------------------------------------
	// Get the type
	//-----------------------------------------
	s := strings.TrimSpace(strings.ToLower(sa[ARType]))
	switch s {
	case "assessment":
		b.ARType = 0
	case "receipt":
		b.ARType = 1
	case "expense":
		b.ARType = 2
	default:
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - ARType must be either Assessment or Receipt.  Found: %s", funcname, lineno, s)
	}

	//----------------------------------------------------------------
	// Get the Debit account
	//----------------------------------------------------------------
	b.DebitLID, err = rlib.IntFromString(sa[DebitLID], "Invalid DebitLID") // first see if it is a LID
	if err == nil && b.DebitLID > 0 {
		gl := rlib.GetLedger(b.DebitLID)
		if gl.LID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No GL Account with ID = %d", funcname, lineno, b.DebitLID)
		}
	} else {
		l := rlib.GetLedgerByName(b.BID, sa[DebitLID])
		if l.LID > 0 {
			b.DebitLID = l.LID
		}
	}
	if b.DebitLID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not find GLAccount for = %s", funcname, lineno, sa[DebitLID])
	}

	//----------------------------------------------------------------
	// Get the Credit account
	//----------------------------------------------------------------
	b.CreditLID, err = rlib.IntFromString(sa[CreditLID], "Invalid CreditLID")
	if err == nil || b.CreditLID > 0 {
		gl := rlib.GetLedger(b.CreditLID)
		if gl.LID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No GL Account with ID = %d", funcname, lineno, b.CreditLID)
		}
	} else {
		l := rlib.GetLedgerByName(b.BID, sa[CreditLID])
		if l.LID > 0 {
			b.CreditLID = l.LID
		}
	}
	if b.CreditLID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid GLAccountID: %s", funcname, lineno, sa[CreditLID])
	}

	//----------------------------------------------------------------
	// Get the Description
	//----------------------------------------------------------------
	b.Description = sa[Description]

	_, err = rlib.InsertAR(&b)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: error inserting AR = %v", funcname, err)
	}

	return 0, nil
}

// LoadARCSV loads the values from the supplied csv file and creates AR records.
func LoadARCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateAR)
}
