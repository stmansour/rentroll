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
func CreateDepositoriesFromCSV(sa []string, lineno int) (string, int) {
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
	if len(sa[BUD]) > 0 {
		b1 := rlib.GetBusinessByDesignation(sa[BUD])
		if len(b1.Designation) == 0 {
			rs += fmt.Sprintf("%s: line %d - rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
			return rs, CsvErrorSensitivity
		}
		d.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Name
	//-------------------------------------------------------------------
	d.Name = strings.TrimSpace(sa[Name])
	if len(d.Name) == 0 {
		rs += fmt.Sprintf("%s: line %d - no name for Depository. Please supply a name\n", funcname, lineno)
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// AccountNo
	//-------------------------------------------------------------------
	d.AccountNo = strings.TrimSpace(sa[AccountNo])
	if len(d.AccountNo) == 0 {
		rs += fmt.Sprintf("%s: line %d - no AccountNo for Depository. Please supply AccountNo\n", funcname, lineno)
		return rs, CsvErrorSensitivity
	}

	_, err = rlib.InsertDepository(&d)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d -  error inserting depository: %v\n", funcname, lineno, err)
	}
	return rs, 0
}

// LoadDepositoryCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadDepositoryCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateDepositoriesFromCSV(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
