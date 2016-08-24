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
		rlib.ReadLedgerEntries(rows, &l)
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

	l, ok = ledgerCache[s]
	if ok {
		if l.BID == bid {
			return l
		}
	}
	l = rlib.GetLedgerByGLNo(bid, s)
	if 0 == l.LID {
		rlib.Ulog("GetCachedLedgerByGL: error getting ledger %s from business %d. \n", s, bid)
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
		// fGenRcv := false
		// fSecDep := false
		// idx := 0
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
			l.RAID = j.RAID
			l.RID = j.JA[i].RID
			rlib.InsertLedgerEntry(&l)
		}
	}
}

// UpdateSubLedgerMarkers is being added to keep track of totals per Rental
// Agreement at each LedgerMarker. This was necessary in order to determine
// exactly what each RentalAgreement did with respect to a specific ledger
// account.  The RAID is saved in the LedgerEntry. However, if we don't save
// a total in a LedgerMarker, then we would need to go back to the beginning
// of time and search all LedgerEntries // for those that matched a particular
// Rental Agreement.  Instead, we will simply add a LedgerMarker for each
// Rental Agreement that affected a particular account with the total equal to
// its previous balance (if it exists) plus the activity during this period.
//
// If no LedgerMarker is found on or before d1, then one will be created.
//
// A new LedgerMarker will be created at d2 with the new balance.
//
// INPUTS
//		bid   - business id
//		plid  - parent ledger id
//		raid  - which RentalAgreement
//		d1,d2 - time range to look for ledger activity that needs to be
//              incorporated into the LedgerMarker.
//-----------------------------------------------------------------------------
func UpdateSubLedgerMarkers(bid int64, d1, d2 *time.Time) {
	funcname := "UpdateSubLedgerMarkers"
	// For each Rental Agreement
	rows, err := rlib.RRdb.Prepstmt.GetRentalAgreementByBusiness.Query(bid)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var ra rlib.RentalAgreement
		err = rlib.ReadRentalAgreements(rows, &ra)
		if err != nil {
			rlib.Ulog("%s: error reading RentalAgreement: %s\n", funcname, err.Error())
			return
		}

		// fmt.Printf("%s\n", rlib.Tline(80))
		// fmt.Printf("Processing Rental Agreement RA%08d\n", ra.RAID)

		// get all the ledger activity between d1 and d2 involving the current RentalAgreement
		m, err := rlib.GetAllLedgerEntriesForRAID(d1, d2, ra.RAID)
		if err != nil {
			rlib.Ulog("%s: GetLedgerEntriesForRAID returned error: %s\n", funcname, err.Error())
			return
		}

		// fmt.Printf("LedgerEntries for RAID = %d between %s - %s:  %d\n", ra.RAID, d1.Format(rlib.RRDATEFMT4), d2.Format(rlib.RRDATEFMT4), len(m))

		LIDprocessed := make(map[int64]int, 0)

		// Spin through all the transactions for this RAID...
		for i := 0; i < len(m); i++ {
			_, processed := LIDprocessed[m[i].LID] // check this ledger for previous processing
			if processed {                         // did we process it?
				continue // yes: move on to the next one
			}
			if m[i].Amount == float64(0) {
				continue // sometimes an entry slips in with a 0 amount, ignore it
			}

			// find the previous LedgerMarker for the GLAccount.  Create one if none exist...
			lm := rlib.LoadRALedgerMarker(bid, m[i].LID, m[i].RAID, d1)

			// fmt.Printf("%s\n", rlib.Tline(20))
			// fmt.Printf("Processing L%08d\n", m[i].LID)
			// fmt.Printf("LedgerMarker: LM%08d - %10s  Balance: %8.2f\n", lm.LMID, lm.Dt.Format(rlib.RRDATEFMT4), lm.Balance)

			// Spin through the rest of the transactions involving m[i].LID and compute the total
			tot := m[i].Amount
			for j := i + 1; j < len(m); j++ {
				if m[j].LID == m[i].LID {
					tot += m[j].Amount
					// fmt.Printf("\tLE%08d  -  %8.2f\n", m[j].LEID, m[j].Amount)
				}
			}
			LIDprocessed[m[i].LID] = 1 // mark that we've processed this ledger

			// Create a new ledger marker on d2 with the updated total...
			var lm2 rlib.LedgerMarker
			lm2.BID = lm.BID
			lm2.LID = lm.LID
			lm2.RAID = lm.RAID
			lm2.Dt = *d2
			lm2.Balance = lm.Balance + tot
			err = rlib.InsertLedgerMarker(&lm2) // lm2.LMID is updated if no error
			if err != nil {
				rlib.Ulog("%s: InsertLedgerMarker error: %s\n", funcname, err.Error())
				return
			}
			// fmt.Printf("LedgerMarker: RAID = %d, Balance = %8.2f\n", lm2.RAID, lm2.Balance)
		}
	}
	rlib.Errcheck(rows.Err())
}

func closeLedgerPeriod(xbiz *rlib.XBusiness, li *rlib.GLAccount, lm *rlib.LedgerMarker, d1, d2 *time.Time, state int64) {
	rows, err := rlib.RRdb.Prepstmt.GetLedgerEntriesInRangeByLID.Query(li.BID, li.LID, d1, d2)
	rlib.Errcheck(err)
	bal := lm.Balance
	defer rows.Close()
	for rows.Next() {
		var l rlib.LedgerEntry
		rlib.ReadLedgerEntries(rows, &l)
		bal += l.Amount
	}
	rlib.Errcheck(rows.Err())

	var nlm rlib.LedgerMarker
	nlm = *lm
	nlm.Balance = bal
	nlm.Dt = *d2
	nlm.State = state
	// fmt.Printf("nlm - %s - %s   GLNo: %s, Balance: %6.2f\n",
	// 	nlm.DtStart.Format(rlib.RRDATEFMT), nlm.Dt.Format(rlib.RRDATEFMT), nlm.GLNumber, nlm.Balance)
	rlib.InsertLedgerMarker(&nlm)
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
		lm := rlib.GetLatestLedgerMarkerByGLNo(xbiz.P.BID, t[i].GLNumber)
		if lm.LMID == 0 {
			fmt.Printf("%s: Could not get GLAccount %d (%s) in busines %d\n", funcname, t[i].LID, t[i].GLNumber, xbiz.P.BID)
			continue
		}
		// fmt.Printf("lm = %#v\n", lm)
		closeLedgerPeriod(xbiz, &t[i], &lm, d1, d2, rlib.MARKERSTATEOPEN)
	}
	rlib.Errcheck(rows.Err())

	//==============================================================================
	// Now we need to update the ledger markers for RAIDs and RIDs
	//==============================================================================
	UpdateSubLedgerMarkers(xbiz.P.BID, d1, d2)
}
