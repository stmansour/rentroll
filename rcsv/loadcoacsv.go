package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

//   0   1                             2          3,               4             5                    6          7                  8             9         10       11         12          13              14
// BUD,  Name,                         GLNumber,  Parent GLNumber, Collective    Account TYpe,        Balance,   Account Status, Associated,  Type,     Date,     AllowPosting, RARequired, ManageToBudget, Description
// REH,  Bank Account FRB 2332352,     10001,     10000,           Cash,         bank,                0,         active,            Yes,         10,  "2016-03-01",  Yes,       0,          0,              Bla bla bla
// REH,  General Accounts Receivable,  11001,     11000,           Cash,         Accounts Receivable, 0,         active,            Yes,         11,  "2016-03-01",  Yes,       0,          0,              Bla bla bla
// REH,  Friday Lunch Fund,            11099,     11000,           Cash,         Accounts Receivable, 0.00,      active,            No,

// CreateLedgerMarkers reads an assessment type string array and creates a database record for the assessment type
func CreateLedgerMarkers(sa []string, lineno int) (int, error) {
	funcname := "CreateLedgerMarkers"
	inserting := true // this may be changed, depends on the value for sa[7]
	var lm rlib.LedgerMarker
	var l rlib.GLAccount

	const (
		BUD            = 0
		Name           = iota
		GLNumber       = iota
		ParentGLNumber = iota
		Collective     = iota
		AccountType    = iota
		Balance        = iota
		AccountStatus  = iota
		Associated     = iota
		Type           = iota
		Date           = iota
		AllowPosting   = iota
		cRARequired    = iota
		ManageToBudget = iota
		Description    = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"GLNumber", GLNumber},
		{"ParentGLNumber", ParentGLNumber},
		{"Collective", Collective},
		{"AccountType", AccountType},
		{"Balance", Balance},
		{"AccountStatus", AccountStatus},
		{"Associated", Associated},
		{"Type", Type},
		{"Date", Date},
		{"AllowPosting", AllowPosting},
		{"RARequired", cRARequired},
		{"ManageToBudget", ManageToBudget},
		{"Description", Description},
	}

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
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
		}
		lm.BID = b1.BID
		l.BID = b1.BID
	}

	lm.State = 3 // Initial marker, no prior records

	// fmt.Printf("LOADCSV - BEGIN: %v\n", sa)

	//----------------------------------------------------------------------
	// TYPE
	// We'll either be updating an existing account or inserting a new one
	// If updating existing, preload lm with existing info...
	//----------------------------------------------------------------------
	s := strings.TrimSpace(sa[9])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)

		// fmt.Printf("0.1  -  s = %s,  i = %d\n", s, i)

		if err != nil || !(i == 0 || (rlib.GLCASH <= i && i <= rlib.GLLAST)) {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid Default value for account %s: %s.  Value must blank, 0, or between %d and %d\n",
				funcname, lineno, sa[2], s, rlib.GLCASH, rlib.GLLAST)
		}
		l1 := rlib.GetLedgerByType(l.BID, int64(i))
		if l1.LID == 0 {
			return CsvErrorSensitivity, nil
		}
		// fmt.Println("0.2")
		l = l1            // update existing
		inserting = false // looks like this is an update

		lm1 := rlib.GetLatestLedgerMarkerByType(l.BID, l.Type)

		// fmt.Printf("0.25:  lm1 = %#v\n", lm1)
		if lm1.LMID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No default rlib.LedgerMarker for business %d, type = %d\n", funcname, lineno, l.BID, l.Type)
		}
		// fmt.Println("0.3")
		lm = lm1 // we're just going to update the existing information
	}
	// fmt.Println("A")

	//----------------------------------------------------------------------
	// NAME
	//----------------------------------------------------------------------
	l.Name = strings.TrimSpace(sa[1])

	// fmt.Println("B")
	//----------------------------------------------------------------------
	// GLNUMBER
	// Make sure the account number is unique
	//----------------------------------------------------------------------
	g := strings.TrimSpace(sa[2])
	if len(g) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - You must suppy a GL Number for this entry\n", funcname, lineno)
	}
	if len(g) > 0 {
		// if we're inserting a record then it must not already exist
		if inserting {
			ldg := rlib.GetLedgerByGLNo(lm.BID, g)
			if ldg.LID > 0 {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Account already exists: %s\n", funcname, lineno, g)
			}
			// // was there an error in finding an account with this GLNo?
			// if !rlib.IsSQLNoResultsError(err) {
			// 	return CsvErrorSensitivity, fmt.Errorf("%s: line %d, GL Account %s already exists\n", funcname, lineno, g)
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
	g = strings.TrimSpace(sa[3])
	if len(g) > 0 {
		parent := rlib.GetLedgerByGLNo(l.BID, g)
		if parent.LID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error getting GLAccount: %s\n", funcname, lineno, g)
		}
		l.PLID = parent.LID
	}
	// fmt.Println("D")

	//----------------------------------------------------------------------
	// Collective
	//----------------------------------------------------------------------
	// strings.TrimSpace(sa[4])

	//----------------------------------------------------------------------
	// ACCOUNT TYPE
	//----------------------------------------------------------------------
	l.AcctType = strings.TrimSpace(sa[5])

	//----------------------------------------------------------------------
	// OPENING BALANCE
	//----------------------------------------------------------------------
	lm.Balance = float64(0) // assume a 0 starting balance
	g = strings.TrimSpace(sa[6])
	if len(g) > 0 {
		x, err := strconv.ParseFloat(g, 64)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid balance: %s\n", funcname, lineno, sa[6])
		}
		lm.Balance = x
	}
	// fmt.Println("E")

	//----------------------------------------------------------------------
	// GLACCOUNT STATUS
	//----------------------------------------------------------------------
	s = strings.ToLower(strings.TrimSpace(sa[7]))
	if "active" == s {
		l.Status = rlib.ACCTSTATUSACTIVE
	} else if "inactive" == s {
		l.Status = rlib.ACCTSTATUSINACTIVE
	} else {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid account status: %s\n", funcname, lineno, sa[7])
	}

	// fmt.Println("F")
	//----------------------------------------------------------------------
	// ASSOCIATED
	//----------------------------------------------------------------------
	s = strings.ToLower(strings.TrimSpace(sa[8]))
	if len(s) == 0 || "associated" == s || s == "y" || s == "yes" || s == "1" {
		l.RAAssociated = rlib.RAASSOCIATED
	} else if "unassociated" == s || s == "n" || s == "no" || s == "0" {
		l.RAAssociated = rlib.RAUNASSOCIATED
	} else {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid associated/unassociated value: %s\n", funcname, lineno, sa[8])
	}

	// fmt.Println("G")
	//----------------------------------------------------------------------
	// TYPE
	//----------------------------------------------------------------------
	s = strings.TrimSpace(sa[9])
	if len(s) > 0 {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - IsCompany value is invalid: %s\n", funcname, lineno, s)
		}
		if i < 0 || (2 <= i && i <= 9) || i > rlib.GLLAST {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Type value is invalid: %s\n", funcname, lineno, s)
		}
		l.Type = int64(i)
	}

	//----------------------------------------------------------------------
	// DATE for opening balance
	//----------------------------------------------------------------------
	DtStop, err := rlib.StringToDate(sa[10])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[10])
	}
	lm.Dt = DtStop

	//----------------------------------------------------------------------
	// ALLOW POST
	//----------------------------------------------------------------------
	l.AllowPost, err = rlib.YesNoToInt(sa[11])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid value for AllowPost:  %s\n", funcname, lineno, sa[11])
	}

	//----------------------------------------------------------------------
	// RAREQUIRED
	//----------------------------------------------------------------------
	RARequired, errmsg := rlib.IntFromString(sa[12], fmt.Sprintf("Invalid number for RARequired. Must be a number between %d and %d", rlib.RARQDINRANGE, rlib.RARQDLAST))
	if len(errmsg) > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s\n", errmsg)
	}
	if RARequired < rlib.RARQDINRANGE || RARequired > rlib.RARQDLAST {
		return CsvErrorSensitivity, fmt.Errorf("Invalid number for RARequired. Must be a number between %d and %d\n", rlib.RARQDINRANGE, rlib.RARQDLAST)
	}
	l.RARequired = RARequired

	//----------------------------------------------------------------------
	// MANAGE TO BUDGET
	//----------------------------------------------------------------------
	l.ManageToBudget, err = rlib.YesNoToInt(sa[13])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid yes/no value: %s\n", funcname, lineno, sa[13])
	}

	//----------------------------------------------------------------------
	// DESCRIPTION
	//----------------------------------------------------------------------
	if len(sa[14]) > 1024 {
		b := []byte(sa[14])
		l.Description = string(b[:1024])
	} else {
		l.Description = sa[14]
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
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not save rlib.GLAccount marker, err = %v\n", funcname, lineno, err)
	}

	// Now update the markers
	if inserting {
		err = rlib.InsertLedgerMarker(&lm)
	} else {
		err = rlib.UpdateLedgerMarker(&lm)
	}
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not save rlib.GLAccount marker, err = %v\n", funcname, lineno, err)
	}

	return 0, nil
}

// LoadChartOfAccountsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadChartOfAccountsCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateLedgerMarkers)
}
