package main

func deleteJournalAllocations(jid int64) {
	_, err := App.prepstmt.deleteJournalAllocations.Exec(jid)
	if err != nil {
		ulog("Error deleting journal allocations for JID = %d, error: %v\n", jid, err)
	}
}

func deleteJournalEntry(jid int64) {
	_, err := App.prepstmt.deleteJournalEntry.Exec(jid)
	if err != nil {
		ulog("Error deleting journal entry for JID = %d, error: %v\n", jid, err)
	}
}

func deleteJournalMarker(jmid int64) {
	_, err := App.prepstmt.deleteJournalMarker.Exec(jmid)
	if err != nil {
		ulog("Error deleting journal marker for JID = %d, error: %v\n", jmid, err)
	}
}

func deleteLedgerEntry(lid int64) error {
	_, err := App.prepstmt.deleteLedgerEntry.Exec(lid)
	if err != nil {
		ulog("Error deleting ledger entry for LID = %d, error: %v\n", lid, err)
	}
	return err
}

func deleteLedgerMarker(lmid int64) error {
	_, err := App.prepstmt.deleteLedgerMarker.Exec(lmid)
	if err != nil {
		ulog("Error deleting ledger marker for LID = %d, error: %v\n", lmid, err)
	}
	return err
}
