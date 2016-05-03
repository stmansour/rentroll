package rlib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//    0           1                             2            3                    4             5                 6           7
// Designation, Name,                         GLNumber,  GLAccountType,         Balance,   GLAccountStatus,    Associated,  Default
// REH,         Bank Account FRB 2332352,     10001,     bank,                     0,         active,          Associated,   10
// REH,         General Accounts Receivable,  11001,     Accounts Receivable,      0,         active,          Associated,   11
// REH,         Friday Lunch Fund,            11099,     Accounts Receivable,     0.00,       active,          Unssociated,

// type LedgerMarker struct {
// 		LMID     int64
// 		BID      int64
// 		PID      int64 // only valid if Type == 1
// 		GLNumber string
// 		Status   int64 // Whether a GL Account is currently unknown=0, inactive=1, active=2
// 		State    int64 //0 = unknown, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
// 		DtStart  time.Time
// 		DtStop   time.Time
// 		Balance  float64
// 		Type     int64 // flag: 0 = not a default account, 1 = Payor Account, 10-default cash, 11-GENRCV, 12-GrossSchedRENT, 13-LTL, 14-VAC
// 		Name     string
//		AcctType string  Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, Cost of Goods Sold, Other Income, Other Expense
//		LastModTime
//		LastModBy
// }

// CreateLedgerMarkers reads an assessment type string array and creates a database record for the assessment type
func CreateLedgerMarkers(sa []string) {
	inserting := true // this may be changed, depends on the value for sa[7]
	n := time.Now()
	d1 := time.Date(n.Year(), n.Month(), 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)
	var lm LedgerMarker
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Make sure the business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			Ulog("CreateLedgerMarkers: business with designation %s does net exist\n", des)
			return
		}
		lm.BID = b1.BID
	}
	lm.DtStart = d1
	lm.DtStop = d2
	lm.State = 3 // Initial marker, no prior records

	//----------------------------------------------------------------------
	// We'll either be updating an existing account or inserting a new one
	// If updating existing, preload lm with existing info...
	//----------------------------------------------------------------------
	s := strings.TrimSpace(sa[7])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil || i < DFLTCASH || i > DFLTLAST {
			fmt.Printf("Invalid Default value for account %s: %s.  Value must blank or between %d and %d\n", sa[2], s, DFLTCASH, DFLTLAST)
			return
		}
		lm1, err := GetLatestLedgerMarkerByType(lm.BID, int64(i))
		if nil != err {
			if IsSQLNoResultsError(err) {
				Ulog("CreateLedgerMarkers: No default ledger %d exists\n", i)
				return
			}
		}
		lm = lm1          // we're just going to update the existing information
		inserting = false // looks like this is an update
	}

	//-------------------------------------------------------------------
	// Set the name
	//-------------------------------------------------------------------
	lm.Name = strings.TrimSpace(sa[1])

	//-------------------------------------------------------------------
	// Make sure the account number is unique
	//-------------------------------------------------------------------
	g := strings.TrimSpace(sa[2])
	if len(g) > 0 {
		// if we're inserting a record then it must not already exist
		if inserting {
			_, err := GetLatestLedgerMarkerByGLNo(lm.BID, g)
			if nil == err {
				fmt.Printf("CreateLedgerMarkers: Account already exists: %s\n", g)
				return
			}
			// was there an error in finding an account with this GLNo?
			if !IsSQLNoResultsError(err) {
				Ulog("CreateLedgerMarkers: GL Account %s already exists\n", g)
				return
			}
		}
		lm.GLNumber = g
	}

	// Set the account type
	lm.AcctType = strings.TrimSpace(sa[3])

	// Set balance
	x, err := strconv.ParseFloat(strings.TrimSpace(sa[4]), 64)
	if err != nil {
		Ulog("CreateLedgerMarkers: Invalid balance: %s\n", sa[4])
		return
	}
	lm.Balance = x

	// Set account status
	s = strings.ToLower(strings.TrimSpace(sa[5]))
	if "active" == s {
		lm.Status = ACCTSTATUSACTIVE
	} else if "inactive" == s {
		lm.Status = ACCTSTATUSINACTIVE
	} else {
		fmt.Printf("CreateLedgerMarkers: Invalid account status: %s\n", sa[5])
		return
	}

	// Set associated
	s = strings.ToLower(strings.TrimSpace(sa[6]))
	if "associated" == s {
		lm.Status = RAASSOCIATED
	} else if "unassociated" == s {
		lm.Status = RAUNASSOCIATED
	} else {
		fmt.Printf("CreateLedgerMarkers: Invalid associated/unassociated value: %s\n", sa[6])
		return
	}

	if inserting {
		err = InsertLedgerMarker(&lm)
	} else {
		err = UpdateLedgerMarker(&lm)
	}
	if nil != err {
		fmt.Printf("CreateLedgerMarkers: Could not save ledger marker, err = %v\n", err)
	}
}

// LoadChartOfAccountsCSV loads a csv file with a chart of accounts and creates ledger markers for each
func LoadChartOfAccountsCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateLedgerMarkers(t[i])
	}
}
