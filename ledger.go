package main

import (
	"rentroll/rlib"
	"time"
)

// RemoveLedgerEntries clears out the records in the supplied range provided the range is not closed by a ledgermarker
func RemoveLedgerEntries(xbiz *rlib.XBusiness, d1, d2 *time.Time) error {
	// Remove the ledger entries and the ledgerallocation entries
	rows, err := rlib.RRdb.Prepstmt.GetAllLedgersInRange.Query(xbiz.P.BID, d1, d2)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var l rlib.Ledger
		rlib.Errcheck(rows.Scan(&l.LID, &l.BID, &l.JID, &l.JAID, &l.GLNumber,
			&l.Dt, &l.Amount, &l.Comment, &l.LastModTime, &l.LastModBy))
		rlib.DeleteLedgerEntry(l.LID)

		// only delete the marker if it is in this time range and if it is not the origin marker
		// lm := GetLastLedgerMarker(xbiz.P.BID)
		// if lm.State == MARKERSTATEOPEN && (lm.DtStart.After(*d1) || lm.DtStart.Equal(*d1)) && (lm.DtStop.Before(*d2) || lm.DtStop.Equal(*d2)) {
		// 	deleteLedgerMarker(lm.LMID)
		// }
	}
	return err
}

// GenerateLedgerEntriesFromJournal creates all the ledger records necessary to describe the journal entry provided
func GenerateLedgerEntriesFromJournal(xbiz *rlib.XBusiness, j *rlib.Journal, d1, d2 *time.Time) {
	// lm := GetLastLedgerMarker(xbiz.P.BID)
	// if lm.DtStop.Equal(d1.AddDate(0, 0, -1)) {
	// 	// pfmt.Printf("Generating next month's ledgers\n")
	// } else {
	// 	fmt.Printf("Generating these ledgers will destroy other periods of ledger records\n")
	// }
	// bal := lm.Balance

	for i := 0; i < len(j.JA); i++ {
		m := parseAcctRule(xbiz, j.JA[i].RID, d1, d2, j.JA[i].AcctRule, j.JA[i].Amount, 1.0)
		for k := 0; k < len(m); k++ {
			var l rlib.Ledger
			l.BID = xbiz.P.BID
			l.JID = j.JID
			l.JAID = j.JA[i].JAID
			l.Dt = j.Dt
			l.Amount = rlib.RoundToCent(m[k].Amount)
			if m[k].Action == "c" {
				l.Amount = -l.Amount
			}
			l.GLNumber = m[k].Account
			rlib.InsertLedgerEntry(&l)

			// bal += l.Amount
		}
	}
}

func closeLedgerPeriod(xbiz *rlib.XBusiness, lm *rlib.LedgerMarker, d1, d2 *time.Time, state int64) {
	rows, err := rlib.RRdb.Prepstmt.GetLedgerInRangeByGLNo.Query(lm.BID, lm.GLNumber, d1, d2)
	rlib.Errcheck(err)
	bal := lm.Balance
	defer rows.Close()
	for rows.Next() {
		var l rlib.Ledger
		rlib.Errcheck(rows.Scan(&l.LID, &l.BID, &l.JID, &l.JAID, &l.GLNumber, &l.Dt,
			&l.Amount, &l.Comment, &l.LastModTime, &l.LastModBy))
		bal += l.Amount
	}
	rlib.Errcheck(rows.Err())

	var nlm rlib.LedgerMarker
	nlm = *lm
	nlm.Balance = bal
	nlm.DtStart = *d1
	nlm.DtStop = d2.AddDate(0, 0, -1)
	nlm.State = state
	// fmt.Printf("nlm - %s - %s   GLNo: %s, Balance: %6.2f\n",
	// 	nlm.DtStart.Format(rlib.RRDATEFMT), nlm.DtStop.Format(rlib.RRDATEFMT), nlm.GLNumber, nlm.Balance)
	rlib.InsertLedgerMarker(&nlm)
}

// GenerateLedgerRecords creates ledgers records based on the journal records over the supplied time range.
func GenerateLedgerRecords(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	err := RemoveLedgerEntries(xbiz, d1, d2)
	if err != nil {
		rlib.Ulog("Could not remove existing Ledger entries from %s to %s. err = %v\n", d1.Format(rlib.RRDATEFMT), d2.Format(rlib.RRDATEFMT), err)
		return
	}
	//==============================================================================
	// Loop through the journal records for this time period, update all ledgers...
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
	// Spin through all ledgers and update the ledger markers with the ending balance...
	//==============================================================================
	t := rlib.GetLedgerMarkerInitList(xbiz.P.BID) // this list contains the list of all ledger account numbers
	// fmt.Printf("len(t) =  %d\n", len(t))
	for i := 0; i < len(t); i++ {
		lm := rlib.GetLatestLedgerMarkerByGLNo(xbiz.P.BID, t[i].GLNumber)
		// fmt.Printf("lm = %#v\n", lm)
		closeLedgerPeriod(xbiz, &lm, d1, d2, rlib.MARKERSTATEOPEN)
	}
	rlib.Errcheck(rows.Err())
}
