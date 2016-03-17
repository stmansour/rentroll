package main

// InsertJournalEntry writes a new journal entry to the database
func InsertJournalEntry(j *Journal) (int, error) {
	var rid = int(0)
	res, err := App.prepstmt.insertJournal.Exec(j.BID, j.RAID, j.Dt, j.Amount, j.Type, j.ID)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int(id)
		}
	}
	return rid, err
}

// InsertJournalAllocationEntry writes a new journalallocation record to the database
func InsertJournalAllocationEntry(ja *JournalAllocation) error {
	_, err := App.prepstmt.insertJournalAllocation.Exec(ja.JID, ja.Amount, ja.ASMID, ja.AcctRule)
	return err
}

// InsertJournalMarker writes a new journalmarker record to the database
func InsertJournalMarker(jm *JournalMarker) error {
	_, err := App.prepstmt.insertJournalMarker.Exec(jm.BID, jm.State, jm.DtStart, jm.DtStop)
	return err
}
