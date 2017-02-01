package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// CVS record format:
// 0    1           2
// BUD, Name,       AccountNo

// CreateDepositoriesFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateDepositoriesFromCSV(sa []string, lineno int) (int, error) {
	funcname := "CreateDepositoriesFromCSV"
	var err error
	var d rlib.Depository

	const (
		BUD       = 0
		Name      = iota
		AccountNo = iota
	)
	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"AccountNo", AccountNo},
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
	if len(sa[BUD]) > 0 {
		b1 := rlib.GetBusinessByDesignation(sa[BUD])
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - rlib.Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		d.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Name
	//-------------------------------------------------------------------
	d.Name = strings.TrimSpace(sa[Name])
	if len(d.Name) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - no name for Depository. Please supply a name", funcname, lineno)
	}

	//-------------------------------------------------------------------
	// AccountNo
	//-------------------------------------------------------------------
	d.AccountNo = strings.TrimSpace(sa[AccountNo])
	if len(d.AccountNo) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - no AccountNo for Depository. Please supply AccountNo", funcname, lineno)
	}
	dup := rlib.GetDepositoryByAccount(d.BID, d.AccountNo)
	if dup.DEPID != 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  depository with account number %s already exists", funcname, lineno, d.AccountNo)
	}

	_, err = rlib.InsertDepository(&d)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error inserting depository: %v", funcname, lineno, err)
	}
	return 0, nil
}

// LoadDepositoryCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadDepositoryCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateDepositoriesFromCSV)
}
