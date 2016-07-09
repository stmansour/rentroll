package main

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// RemoveLedgerEntries clears out the records in the supplied range provided the range is not closed by a LedgerMarker
func RemoveLedgerEntries(xbiz *rlib.XBusiness, d1, d2 *time.Time) error {
	// Remove the LedgerEntries and the ledgerallocation entries
	rows, err := rlib.RRdb.Prepstmt.GetAllLedgerEntriesInRange.Query(xbiz.P.BID, d1, d2)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var l rlib.LedgerEntry
		rlib.Errcheck(rows.Scan(&l.LEID, &l.BID, &l.JID, &l.JAID, &l.LID, &l.RAID,
			&l.Dt, &l.Amount, &l.Comment, &l.LastModTime, &l.LastModBy))
		rlib.DeleteLedgerEntry(l.LEID)
	}
	return err
}

// ledgerCache is a mapping of glNames to ledger structs
var ledgerCache map[string]rlib.GLAccount

// initLedgerCache starts a new ledger cache
func initLedgerCache() {
	ledgerCache = make(map[string]rlib.GLAccount)
}

// GetCachedLedgerByGL checks the cache with index string s. If there is an entry there and the BID matches the
// requested BID we return the ledger struct immediately. Otherwise, the ledger is loaded from the database and
// stored in the cache at index s.  If no ledger is found with GLNumber s, then a ledger with LID = 0 is returned.
func GetCachedLedgerByGL(bid int64, s string) rlib.GLAccount {
	var l rlib.GLAccount
	var ok bool
	var err error

	l, ok = ledgerCache[s]
	if ok {
		if l.BID == bid {
			return l
		}
	}
	l, err = rlib.GetLedgerByGLNo(bid, s)
	if err != nil {
		rlib.Ulog("GetCachedLedgerByGL: error getting ledger %s from business %d, err = %s\n", s, bid, err.Error())
		l.LID = 0
	} else {
		ledgerCache[s] = l
	}
	return l
}

// GenerateLedgerEntriesFromJournal creates all the LedgerEntries necessary to describe the Journal entry provided
func GenerateLedgerEntriesFromJournal(xbiz *rlib.XBusiness, j *rlib.Journal, d1, d2 *time.Time) {
	for i := 0; i < len(j.JA); i++ {
		m := rlib.ParseAcctRule(xbiz, j.JA[i].RID, d1, d2, j.JA[i].AcctRule, j.JA[i].Amount, 1.0)
		for k := 0; k < len(m); k++ {
			var l rlib.LedgerEntry
			l.BID = xbiz.P.BID
			l.JID = j.JID
			l.JAID = j.JA[i].JAID
			l.Dt = j.Dt
			l.Amount = rlib.RoundToCent(m[k].Amount)
			if m[k].Action == "c" {
				l.Amount = -l.Amount
			}
			ledger := GetCachedLedgerByGL(l.BID, m[k].Account)
			l.LID = ledger.LID
			// l.GLNo = ledger.GLNumber
			l.RAID = j.RAID
			rlib.InsertLedgerEntry(&l)
		}
	}
}

func closeLedgerPeriod(xbiz *rlib.XBusiness, li *rlib.GLAccount, lm *rlib.LedgerMarker, d1, d2 *time.Time, state int64) {
	// rows, err := rlib.RRdb.Prepstmt.GetLedgerEntriesInRangeByGLNo.Query(li.BID, li.GLNumber, d1, d2)
	rows, err := rlib.RRdb.Prepstmt.GetLedgerEntriesInRangeByLID.Query(li.BID, li.LID, d1, d2)
	rlib.Errcheck(err)
	bal := lm.Balance
	defer rows.Close()
	for rows.Next() {
		var l rlib.LedgerEntry
		rlib.Errcheck(rows.Scan(&l.LEID, &l.BID, &l.JID, &l.JAID, &l.LID, &l.RAID, &l.Dt,
			&l.Amount, &l.Comment, &l.LastModTime, &l.LastModBy))
		bal += l.Amount
	}
	rlib.Errcheck(rows.Err())

	var nlm rlib.LedgerMarker
	nlm = *lm
	nlm.Balance = bal
	nlm.DtStart = *d1
	nlm.DtStop = d2.AddDate(0, 0, -1) // TODO: subtracting 1 day may not be correct
	nlm.State = state
	// fmt.Printf("nlm - %s - %s   GLNo: %s, Balance: %6.2f\n",
	// 	nlm.DtStart.Format(rlib.RRDATEFMT), nlm.DtStop.Format(rlib.RRDATEFMT), nlm.GLNumber, nlm.Balance)
	rlib.InsertLedgerMarker(&nlm)
}

// LoadRABalanceLedger returns a balance ledger for the supplied RentalAgreement, creating it if necessary.
func LoadRABalanceLedger(ra *rlib.RentalAgreement, d1, d2 *time.Time, bid int64) (rlib.GLAccount, error) {
	l, err := rlib.GetRABalanceLedger(bid, ra.RAID)
	if err != nil {
		if rlib.IsSQLNoResultsError(err) {
			var l rlib.GLAccount
			l.BID = bid
			l.Type = rlib.RABALANCEACCOUNT
			l.RAAssociated = 2
			l.RAID = ra.RAID
			l.Status = 2
			l.Name = fmt.Sprintf("RA%08d Balance", ra.RAID)
			l.LID, err = rlib.InsertLedger(&l)
			// fmt.Printf("LoadRABalanceLedger: CREATING LedgerBalance account: for RAID = %d;  LID = %d\n", ra.RAID, l.LID)
			return l, err
		}
		rlib.Ulog("LoadRABalanceLedger: error getting RABalanceLedger for BID=%d, RAID=%d, err = %s\n", bid, ra.RAID, err.Error())
	}
	return l, err
}

// GetAccountBalanceForRA returns the summed Amount balance for activity
// in GLAccount lid associated with RentalAgreement raid
func GetAccountBalanceForRA(bid, lid, raid int64, d1, d2 *time.Time) (float64, error) {
	var bal = float64(0)
	m, err := rlib.GetLedgerEntriesForRAID(d1, d2, raid, lid)
	if err != nil {
		return bal, err
	}
	for i := 0; i < len(m); i++ {
		bal += m[i].Amount
	}
	return bal, err
}

// GenerateRABalances creates the ledgerMarkers for the Type 1 RentalAgreement accounts
func GenerateRABalances(bid int64, d1, d2 *time.Time) error {
	var state = int64(rlib.MARKERSTATEOPEN)
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementsByRange.Query(bid, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()

	// Spin through all the RentalAgreements that are active in this timeframe
	for rows.Next() {
		var ra rlib.RentalAgreement
		rlib.Errcheck(rows.Scan(&ra.RAID, &ra.RATID, &ra.BID, &ra.NID, &ra.RentalStart, &ra.RentalStop, &ra.PossessionStart, &ra.PossessionStop,
			&ra.Renewal, &ra.SpecialProvisions, &ra.LastModTime, &ra.LastModBy))

		// get or create a ledger for this rental agreement
		l, err := LoadRABalanceLedger(&ra, d1, d2, bid)
		if err != nil {
			return err
		}

		// initialize balance from the last marker if it exists
		openingBal := float64(0)
		lm, err := rlib.GetLatestLedgerMarkerByLID(bid, l.LID)
		if err != nil {
			if !rlib.IsSQLNoResultsError(err) {
				return err
			}
			state = int64(rlib.MARKERSTATEORIGIN)
		} else {
			// if the stop date of this marker is past our startdate, then we have big problems.
			if d1.Before(lm.DtStop) {
				return fmt.Errorf("GenerateRABalances: existing LedgerMarker for RAID %d has stop date %s, past current period start date %s\n",
					ra.RAID, lm.DtStop.Format(rlib.RRDATEINPFMT), d1.Format(rlib.RRDATEINPFMT))
			}
			openingBal = lm.Balance
		}

		// With the opening balance now set, we now need to add up the activity that has happened over the current period.
		// This means we total up all the activity in the GeneralReceivables account during this period.
		lid := rlib.RRdb.BizTypes[bid].DefaultAccts[rlib.GLGENRCV].LID
		delta, err := GetAccountBalanceForRA(bid, lid, ra.RAID, d1, d2)
		if err != nil {
			fmt.Printf("error returned from GetAccountBalanceForRA:  err = %s\n", err.Error()) // ****** PURGE ME *******
			return err
		}

		// Create a new LedgerMarker for GLAccount l with the updated balance:
		var nlm rlib.LedgerMarker
		nlm.LID = l.LID
		nlm.BID = bid
		nlm.DtStart = *d1
		nlm.DtStop = *d2
		nlm.Balance = openingBal + delta
		nlm.State = state
		// fmt.Printf("INSERTING LEDGER MARKER:  %s - %s   LID: %d, Balance: %6.2f\n",
		// 	nlm.DtStart.Format(rlib.RRDATEFMT), nlm.DtStop.Format(rlib.RRDATEFMT), nlm.LID, nlm.Balance)
		err = rlib.InsertLedgerMarker(&nlm)
		rlib.Errlog(err)
	}
	rlib.Errcheck(rows.Err())
	return err
}

// GenerateLedgerRecords creates ledgers records based on the Journal records over the supplied time range.
func GenerateLedgerRecords(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	funcname := "GenerateLedgerRecords"
	err := RemoveLedgerEntries(xbiz, d1, d2)
	if err != nil {
		rlib.Ulog("Could not remove existing LedgerEntries from %s to %s. err = %v\n", d1.Format(rlib.RRDATEFMT), d2.Format(rlib.RRDATEFMT), err)
		return
	}
	initLedgerCache()
	//==============================================================================
	// Loop through the Journal records for this time period, update all ledgers...
	//==============================================================================
	rows, err := rlib.RRdb.Prepstmt.GetAllJournalsInRange.Query(xbiz.P.BID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	// fmt.Printf("Loading Journal Entries from %s to %s.\n", d1.Format(rlib.RRDATEFMT), d2.Format(rlib.RRDATEFMT))
	for rows.Next() {
		var j rlib.Journal
		rlib.Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID, &j.Comment, &j.LastModTime, &j.LastModBy))
		rlib.GetJournalAllocations(j.JID, &j)
		GenerateLedgerEntriesFromJournal(xbiz, &j, d1, d2)
	}
	rlib.Errcheck(rows.Err())

	//==============================================================================
	// Now that all the ledgers have been updated, we can close the ledgers and mark
	// their state as MARKERSTATEOPEN
	// Spin through all ledgers and update the LedgerMarkers with the ending balance...
	//==============================================================================
	t := rlib.GetLedgerList(xbiz.P.BID) // this list contains the list of all GLAccount numbers
	// fmt.Printf("len(t) =  %d\n", len(t))
	for i := 0; i < len(t); i++ {
		if t[i].Type != rlib.RABALANCEACCOUNT {
			lm, err := rlib.GetLatestLedgerMarkerByGLNo(xbiz.P.BID, t[i].GLNumber)
			if err != nil {
				fmt.Printf("%s: Could not get GLAccount %d (%s) in busines %d\n", funcname, t[i].LID, t[i].GLNumber, xbiz.P.BID)
				fmt.Printf("%s: Error = %v\n", funcname, err)
				continue
			}
			// fmt.Printf("lm = %#v\n", lm)
			closeLedgerPeriod(xbiz, &t[i], &lm, d1, d2, rlib.MARKERSTATEOPEN)
		}
	}
	rlib.Errcheck(rows.Err())
	rlib.Errlog(GenerateRABalances(xbiz.P.BID, d1, d2))
}
