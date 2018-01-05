package rcsv

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// CVS record format:
//	                GLAccount can be Account Name, GLNumber, or LID
// 0    1           2         3
// BUD, GLAccount,  Name,     AccountNo

// CreateDepositoriesFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateDepositoriesFromCSV(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateDepositoriesFromCSV"
	var (
		err error
		d   rlib.Depository
	)

	const (
		BUD       = 0
		LID       = iota
		Name      = iota
		AccountNo = iota
	)
	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"GLAccount", LID},
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
		b1, err := rlib.GetBusinessByDesignation(ctx, sa[BUD])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - rlib.Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - rlib.Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		d.BID = b1.BID
	}

	if len(sa[LID]) > 0 {
		var acct rlib.GLAccount
		i, err := strconv.Atoi(sa[LID])
		if err == nil {
			d.LID = int64(i)
		}
		// validate that this is a valid LID
		if d.LID > 0 {
			// TODO(Steve): ignore error?
			acct, _ = rlib.GetLedger(ctx, d.LID)
		}
		if acct.LID == 0 {
			// TODO(Steve): ignore error?
			gl, _ := rlib.GetLedgerByGLNo(ctx, d.BID, sa[LID])
			if gl.LID == 0 {
				// TODO(Steve): ignore error?
				gl, _ = rlib.GetLedgerByName(ctx, d.BID, sa[LID]) // see if we can find it by name
				if gl.LID == 0 {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No GL Account with Name or AccountNumber = %s", funcname, lineno, sa[LID])
				}
			}
			d.LID = gl.LID
		}
	}
	if d.LID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No GL Account with Name or AccountNumber = %s", funcname, lineno, sa[LID])
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

	// TODO(Steve): error handling in this case?
	dup, _ /*err*/ := rlib.GetDepositoryByAccount(ctx, d.BID, d.AccountNo)
	/*if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  depository with account number %s already exists", funcname, lineno, d.AccountNo)
	}*/
	if dup.DEPID != 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  depository with account number %s already exists", funcname, lineno, d.AccountNo)
	}

	_, err = rlib.InsertDepository(ctx, &d)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error inserting depository: %v", funcname, lineno, err)
	}
	return 0, nil
}

// LoadDepositoryCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadDepositoryCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateDepositoriesFromCSV)
}
