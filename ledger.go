package main

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// RemoveLedgerEntries clears out the records in the supplied range provided the range is not closed by a ledgermarker
func RemoveLedgerEntries(xbiz *XBusiness, d1, d2 *time.Time) error {
	// Remove the ledger entries and the ledgerallocation entries
	rows, err := App.prepstmt.getAllLedgersInRange.Query(xbiz.P.BID, d1, d2)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var l Ledger
		rlib.Errcheck(rows.Scan(&l.LID, &l.BID, &l.JID, &l.JAID, &l.GLNumber, &l.Dt, &l.Amount))
		deleteLedgerEntry(l.LID)

		// only delete the marker if it is in this time range and if it is not the origin marker
		lm := GetLastLedgerMarker(xbiz.P.BID)
		if lm.State == MARKERSTATEOPEN && (lm.DtStart.After(*d1) || lm.DtStart.Equal(*d1)) && (lm.DtStop.Before(*d2) || lm.DtStop.Equal(*d2)) {
			deleteLedgerMarker(lm.LMID)
		}
	}
	return err
}

// GenerateLedgerMarker creates all the ledger markers for the supplied time period
func GenerateLedgerMarker(xbiz *XBusiness, d1, d2 *time.Time, bal float64) {
	lm := GetLastLedgerMarker(xbiz.P.BID)
	diff := lm.DtStop.Sub(*d1)
	if diff < 0 {
		diff = -diff
	}
	if diff > 24*time.Hour {
		s := fmt.Sprintf("Gap between last ledger marker (stop = %s) and next ledger marker (start = %s) is > than 1 day\n",
			lm.DtStop.Format(RRDATEFMT), d1.Format(RRDATEFMT))
		fmt.Println(s)
		ulog(s)
	}
	var l LedgerMarker
	l.BID = xbiz.P.BID
	l.Balance = bal
	l.DtStart = *d1
	l.DtStop = *d2
	l.GLNumber = lm.GLNumber
	l.Name = lm.Name
	l.State = lm.State
	lm.State = lm.Status

	InsertLedgerMarker(&l)
}

// GenerateLedgerEntriesFromJournal creates all the ledger records necessary to describe the journal entry provided
func GenerateLedgerEntriesFromJournal(xbiz *XBusiness, j *Journal, d1, d2 *time.Time) {
	lm := GetLastLedgerMarker(xbiz.P.BID)
	if lm.DtStop.Equal(d1.AddDate(0, 0, -1)) {
		// pfmt.Printf("Generating next month's ledgers\n")
	} else {
		fmt.Printf("Generating these ledgers will destroy other periods of ledger records\n")
	}
	bal := lm.Balance

	for i := 0; i < len(j.JA); i++ {
		m := parseAcctRule(xbiz, j.JA[i].RID, d1, d2, j.JA[i].AcctRule, j.JA[i].Amount, 1.0)
		for k := 0; k < len(m); k++ {
			var l Ledger
			l.BID = xbiz.P.BID
			l.JID = j.JID
			l.JAID = j.JA[i].JAID
			l.Dt = j.Dt
			l.Amount = rlib.RoundToCent(m[k].Amount)
			if m[k].Action == "c" {
				l.Amount = -l.Amount
			}
			l.GLNumber = m[k].Account
			InsertLedgerEntry(&l)

			bal += l.Amount
		}
	}
	GenerateLedgerMarker(xbiz, d1, d2, bal)
}

// GenerateLedgerRecords creates ledgers records based on the journal records over the supplied time range.
func GenerateLedgerRecords(xbiz *XBusiness, d1, d2 *time.Time) {
	err := RemoveLedgerEntries(xbiz, d1, d2)
	if err != nil {
		ulog("Could not remove existing Ledger entries from %s to %s. err = %v\n", d1.Format(RRDATEFMT), d2.Format(RRDATEFMT), err)
		return
	}
	//==============================================================================
	// Loop through the journal records for this time period, update all ledgers...
	//==============================================================================
	rows, err := App.prepstmt.getAllJournalsInRange.Query(xbiz.P.BID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	// fmt.Printf("Loading Journal Entries from %s to %s.\n", d1.Format(RRDATEFMT), d2.Format(RRDATEFMT))
	for rows.Next() {
		var j Journal
		rlib.Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID))
		GetJournalAllocations(j.JID, &j)
		GenerateLedgerEntriesFromJournal(xbiz, &j, d1, d2)
	}
	rlib.Errcheck(rows.Err())
}
