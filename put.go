package main

// InsertJournalEntry writes a new journal entry to the database
func InsertJournalEntry(j *Journal) (int64, error) {
	var rid = int64(0)
	res, err := App.prepstmt.insertJournal.Exec(j.BID, j.RAID, j.Dt, j.Amount, j.Type, j.ID)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	}
	return rid, err
}

// InsertJournalAllocationEntry writes a new journalallocation record to the database
func InsertJournalAllocationEntry(ja *JournalAllocation) error {
	_, err := App.prepstmt.insertJournalAllocation.Exec(ja.JID, ja.RID, ja.Amount, ja.ASMID, ja.AcctRule)
	return err
}

// InsertJournalMarker writes a new journalmarker record to the database
func InsertJournalMarker(jm *JournalMarker) error {
	_, err := App.prepstmt.insertJournalMarker.Exec(jm.BID, jm.State, jm.DtStart, jm.DtStop)
	return err
}

// InsertLedgerMarker writes a new journalmarker record to the database
func InsertLedgerMarker(l *LedgerMarker) error {
	_, err := App.prepstmt.insertLedgerMarker.Exec(l.LMID, l.BID, l.GLNumber, l.State, l.DtStart, l.DtStop, l.Balance, l.Type, l.Name)
	return err
}

// InsertLedgerEntry writes a new journal entry to the database
func InsertLedgerEntry(l *Ledger) (int64, error) {
	var rid = int64(0)
	res, err := App.prepstmt.insertLedger.Exec(l.BID, l.JID, l.JAID, l.GLNumber, l.Dt, l.Amount)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		ulog("Error inserting ledger:  %v\n", err)
	}
	return rid, err
}
