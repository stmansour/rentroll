package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

//    0           1                             2        3,               4                       5            6                7             8         9          10
// Designation, Name,                         GLNumber,  Parent GLNumber   Account TYpe,        Balance,   GL Account Status,    Associated,  Type,     Date,      AllowPost
// REH,         Bank Account FRB 2332352,     10001,      10000           bank,                     0,         active,          Yes,         10,  "2016-03-01",  Yes
// REH,         General Accounts Receivable,  11001,      11000           Accounts Receivable,      0,         active,          Yes,         11,  "2016-03-01",  Yes
// REH,         Friday Lunch Fund,            11099,      11000           Accounts Receivable,     0.00,       active,          No,

// StringToDate tries to convert the supplied string to a time.Time value. It will use the two
// formats called out in dbtypes.go:  rlib.RRDATEFMT, rlib.RRDATEINPFMT, rlib.RRDATEINPFMT2
func StringToDate(s string) (time.Time, error) {
	// try the ansi std date format first
	s = strings.TrimSpace(s)
	Dt, err := time.Parse(rlib.RRDATEINPFMT, s)
	if err != nil {
		Dt, err = time.Parse(rlib.RRDATEFMT2, s) // try excel default version
		if err != nil {
			Dt, err = time.Parse(rlib.RRDATEFMT, s) // try 0 filled version
			if nil != err {
				Dt, err = time.Parse(rlib.RRDATEFMT3, s) // try 4 digit year version
			}
		}
	}
	return Dt, err
}

// CreateLedgerMarkers reads an assessment type string array and creates a database record for the assessment type
func CreateLedgerMarkers(sa []string, lineno int) {
	funcname := "CreateLedgerMarkers"
	inserting := true // this may be changed, depends on the value for sa[7]
	var lm rlib.LedgerMarker
	var l rlib.GLAccount
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "bud" {
		return // this is just the column heading
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 11
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d, rlib.Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		lm.BID = b1.BID
		l.BID = b1.BID
	}

	lm.State = 3 // Initial marker, no prior records

	//----------------------------------------------------------------------
	// TYPE
	// We'll either be updating an existing account or inserting a new one
	// If updating existing, preload lm with existing info...
	//----------------------------------------------------------------------
	s := strings.TrimSpace(sa[8])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil || !(i == 0 || (rlib.DFLTCASH <= i && i <= rlib.DFLTLAST)) {
			fmt.Printf("%s: line %d - Invalid Default value for account %s: %s.  Value must blank, 0, or between %d and %d\n",
				funcname, lineno, sa[2], s, rlib.DFLTCASH, rlib.DFLTLAST)
			return
		}
		l1, err := rlib.GetLedgerByType(l.BID, int64(i))
		if nil != err {
			if rlib.IsSQLNoResultsError(err) {
				rlib.Ulog("%s: line %d - No default rlib.GLAccount %d exists\n", funcname, lineno, i)
				return
			}
		}
		l = l1            // update existing
		inserting = false // looks like this is an update
		lm1, err := rlib.GetLatestLedgerMarkerByType(l.BID, l.Type)
		if nil != err {
			if rlib.IsSQLNoResultsError(err) {
				rlib.Ulog("%s: line %d - No default rlib.LedgerMarker %d exists\n", funcname, lineno, i)
				return
			}
		}
		lm = lm1 // we're just going to update the existing information
	}

	//----------------------------------------------------------------------
	// NAME
	//----------------------------------------------------------------------
	l.Name = strings.TrimSpace(sa[1])

	//----------------------------------------------------------------------
	// GLNUMBER
	// Make sure the account number is unique
	//----------------------------------------------------------------------
	g := strings.TrimSpace(sa[2])
	if len(g) == 0 {
		fmt.Printf("%s: line %d - You must suppy a GL Number for this entry\n", funcname, lineno)
		return
	}
	if len(g) > 0 {
		// if we're inserting a record then it must not already exist
		if inserting {
			_, err := rlib.GetLedgerByGLNo(lm.BID, g)
			if nil == err {
				fmt.Printf("%s: line %d - Account already exists: %s\n", funcname, lineno, g)
				return
			}
			// was there an error in finding an account with this GLNo?
			if !rlib.IsSQLNoResultsError(err) {
				rlib.Ulog("%s: line %d, GL Account %s already exists\n", funcname, lineno, g)
				return
			}
		}
		l.GLNumber = g
	}

	//----------------------------------------------------------------------
	// PARENT GLNUMBER
	//----------------------------------------------------------------------
	l.PLID = int64(0) // assume no parent
	g = strings.TrimSpace(sa[3])
	if len(g) > 0 {
		parent, err := rlib.GetLedgerByGLNo(l.BID, g)
		if nil != err {
			fmt.Printf("%s: line %d - Error getting GLAccount: %s,  error = %s\n", funcname, lineno, g, err.Error())
			return
		}
		l.PLID = parent.LID
	}

	//----------------------------------------------------------------------
	// ACCOUNT TYPE
	//----------------------------------------------------------------------
	l.AcctType = strings.TrimSpace(sa[4])

	//----------------------------------------------------------------------
	// OPENING BALANCE
	//----------------------------------------------------------------------
	lm.Balance = float64(0) // assume a 0 starting balance
	g = strings.TrimSpace(sa[5])
	if len(g) > 0 {
		x, err := strconv.ParseFloat(g, 64)
		if err != nil {
			rlib.Ulog("%s: line %d - Invalid balance: %s\n", funcname, lineno, sa[5])
			return
		}
		lm.Balance = x
	}

	//----------------------------------------------------------------------
	// GLACCOUNT STATUS
	//----------------------------------------------------------------------
	s = strings.ToLower(strings.TrimSpace(sa[6]))
	if "active" == s {
		l.Status = rlib.ACCTSTATUSACTIVE
	} else if "inactive" == s {
		l.Status = rlib.ACCTSTATUSINACTIVE
	} else {
		fmt.Printf("%s: line %d - Invalid account status: %s\n", funcname, lineno, sa[6])
		return
	}

	//----------------------------------------------------------------------
	// ASSOCIATED
	//----------------------------------------------------------------------
	s = strings.ToLower(strings.TrimSpace(sa[7]))
	if "associated" == s || s == "y" || s == "yes" || s == "1" {
		l.RAAssociated = rlib.RAASSOCIATED
	} else if "unassociated" == s || s == "n" || s == "no" || s == "0" {
		l.RAAssociated = rlib.RAUNASSOCIATED
	} else {
		fmt.Printf("%s: line %d - Invalid associated/unassociated value: %s\n", funcname, lineno, sa[7])
		return
	}

	//----------------------------------------------------------------------
	// TYPE
	//----------------------------------------------------------------------
	s = strings.TrimSpace(sa[8])
	if len(s) > 0 {

		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Printf("%s: line %d - IsCompany value is invalid: %s\n", funcname, lineno, s)
			return
		}
		if i < 0 || (2 <= i && i <= 9) || i > rlib.DFLTLAST {
			fmt.Printf("%s: line %d - Type value is invalid: %s\n", funcname, lineno, s)
			return
		}
		l.Type = int64(i)
	}

	//----------------------------------------------------------------------
	// DATE for opening balance
	//----------------------------------------------------------------------
	DtStop, err := StringToDate(sa[9])
	if err != nil {
		fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[9])
		return
	}
	lm.DtStop = DtStop
	lm.DtStart = DtStop.AddDate(0, -1, 0)

	//----------------------------------------------------------------------
	// ALLOW POST
	//----------------------------------------------------------------------
	l.AllowPost, err = rlib.YesNoToInt(sa[10])
	if err != nil {
		fmt.Printf("%s: line %d - invalid value for AllowPost:  %s\n", funcname, lineno, sa[10])
		return
	}

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
		fmt.Printf("%s: line %d - Could not save rlib.GLAccount marker, err = %v\n", funcname, lineno, err)
	}

	// Now update the markers
	if inserting {
		err = rlib.InsertLedgerMarker(&lm)
	} else {
		err = rlib.UpdateLedgerMarker(&lm)
	}
	if nil != err {
		fmt.Printf("%s: line %d - Could not save rlib.GLAccount marker, err = %v\n", funcname, lineno, err)
	}
}

// LoadChartOfAccountsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadChartOfAccountsCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateLedgerMarkers(t[i], i+1)
	}
}
