package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// TODO:
// Remove Associated
// Remove Collective
// Allow Posting
// RA Required

//   0   1                             2          3,               4                    5          6                 7           8
// BUD,  Name,                         GLNumber,  Parent GLNumber, Account Type,        Balance,   Account Status,   Date,       Description
// REH,  Bank Account FRB 2332352,     10001,     10000,           bank,                0,         active,         "2016-03-01", Bla bla bla
// REH,  General Accounts Receivable,  11001,     11000,           Accounts Receivable, 0,         active,         "2016-03-01", Bla bla bla
// REH,  Friday Lunch Fund,            11099,     11000,           Accounts Receivable, 0.00,      active,

// CreateLedgerMarkers reads an assessment type string array and creates a database record for the assessment type
func CreateLedgerMarkers(sa []string, lineno int) (int, error) {
	funcname := "CreateLedgerMarkers"
	inserting := true // this may be changed, depends on the value for sa[7]
	var lm rlib.LedgerMarker
	var l rlib.GLAccount
	var parent rlib.GLAccount

	const (
		BUD            = 0
		Name           = iota
		GLNumber       = iota
		ParentGLNumber = iota
		AccountType    = iota
		Balance        = iota
		AccountStatus  = iota
		Date           = iota
		Description    = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"GLNumber", GLNumber},
		{"ParentGLNumber", ParentGLNumber},
		{"AccountType", AccountType},
		{"Balance", Balance},
		{"AccountStatus", AccountStatus},
		{"Date", Date},
		{"Description", Description},
	}

	l.AllowPost = 1 // default is to allow, server will modify as needed

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		// fmt.Printf("Looking for BUD:  %s\n", des)
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, rlib.Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		lm.BID = b1.BID
		l.BID = b1.BID
	}

	lm.State = rlib.LMINITIAL // Initial marker, no prior records

	//----------------------------------------------------------------------
	// NAME
	//----------------------------------------------------------------------
	l.Name = strings.TrimSpace(sa[Name])

	// fmt.Println("B")
	//----------------------------------------------------------------------
	// GLNUMBER
	// Make sure the account number is unique
	//----------------------------------------------------------------------
	g := strings.TrimSpace(sa[GLNumber])
	if len(g) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - You must supply a GL Number for this entry", funcname, lineno)
	}
	if len(g) > 0 {
		// if we're inserting a record then it must not already exist
		if inserting {
			ldg := rlib.GetLedgerByGLNo(lm.BID, g)
			if ldg.LID > 0 {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Account already exists: %s", funcname, lineno, g)
			}
			// // was there an error in finding an account with this GLNo?
			// if !rlib.IsSQLNoResultsError(err) {
			// 	return CsvErrorSensitivity, fmt.Errorf("%s: line %d, GL Account %s already exists", funcname, lineno, g)
			// 	return rs,CsvErrorSensitivity
			// }
		}
		l.GLNumber = g
	}

	// fmt.Println("C")
	//----------------------------------------------------------------------
	// PARENT GLNUMBER
	//----------------------------------------------------------------------
	l.PLID = int64(0) // assume no parent
	g = strings.TrimSpace(sa[ParentGLNumber])
	if len(g) > 0 {
		parent = rlib.GetLedgerByGLNo(l.BID, g)
		if parent.LID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error getting GLAccount: %s", funcname, lineno, g)
		}
		l.PLID = parent.LID
	}
	// fmt.Println("D")

	//----------------------------------------------------------------------
	// ACCOUNT TYPE
	//----------------------------------------------------------------------
	l.AcctType = strings.TrimSpace(sa[AccountType])

	//----------------------------------------------------------------------
	// OPENING BALANCE
	//----------------------------------------------------------------------
	lm.Balance = float64(0) // assume a 0 starting balance
	g = strings.TrimSpace(sa[Balance])
	if len(g) > 0 {
		x, err := strconv.ParseFloat(g, 64)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid balance: %s", funcname, lineno, sa[Balance])
		}
		lm.Balance = x
	}
	// fmt.Println("E")

	//----------------------------------------------------------------------
	// GLACCOUNT STATUS
	//----------------------------------------------------------------------
	s := strings.ToLower(strings.TrimSpace(sa[AccountStatus]))
	if "active" == s {
		l.Status = rlib.ACCTSTATUSACTIVE
	} else if "inactive" == s {
		l.Status = rlib.ACCTSTATUSINACTIVE
	} else {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid account status: %s", funcname, lineno, sa[AccountStatus])
	}

	// fmt.Println("F")

	//----------------------------------------------------------------------
	// DATE for opening balance
	//----------------------------------------------------------------------
	_, err = rlib.StringToDate(sa[Date])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid stop date:  %s", funcname, lineno, sa[Date])
	}
	lm.Dt = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC) // always force the initial ledger marker to "the beginning of time"

	// //----------------------------------------------------------------------
	// // ALLOW POST
	// //----------------------------------------------------------------------
	// l.AllowPost, err = rlib.YesNoToInt(sa[AllowPosting])
	// if err != nil {
	// 	return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid value for AllowPost:  %s", funcname, lineno, sa[AllowPosting])
	// }

	//----------------------------------------------------------------------
	// DESCRIPTION
	//----------------------------------------------------------------------
	if len(sa[Description]) > 1024 {
		b := []byte(sa[Description])
		l.Description = string(b[:1024])
	} else {
		l.Description = sa[Description]
	}

	//=======================================================================================

	// fmt.Printf("LOADCSV - SAVE:  Inserting = %v\n", inserting)
	// fmt.Printf("                 l = %#v\n", l)

	// Insert / Update the rlib.GLAccount first, we may need the LID
	if inserting {
		var lid int64
		lid, err = rlib.InsertLedger(&l)
		lm.LID = lid
	} else {
		err = rlib.UpdateLedger(&l)
		lm.LID = l.LID
	}
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not save rlib.GLAccount marker, err = %v", funcname, lineno, err)
	}

	// Now update the markers
	if inserting {
		err = rlib.InsertLedgerMarker(&lm)
	} else {
		err = rlib.UpdateLedgerMarker(&lm)
	}
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not save rlib.GLAccount marker, err = %v", funcname, lineno, err)
	}

	//--------------------------------------------------------------
	// If this entry called out a Parent, make sure the Parent's
	// AllowPost attribute is set properly
	//--------------------------------------------------------------
	if l.PLID > 0 && l.PLID == parent.LID {
		if parent.AllowPost == 1 {
			parent.AllowPost = 0
			rlib.UpdateLedger(&parent)
		}
	}

	return 0, nil
}

// LoadChartOfAccountsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadChartOfAccountsCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateLedgerMarkers)
}
