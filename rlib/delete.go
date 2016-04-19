package rlib

// DeleteJournalAllocations deletes the allocation records associated with the supplied jid
func DeleteJournalAllocations(jid int64) {
	_, err := RRdb.Prepstmt.DeleteJournalAllocations.Exec(jid)
	if err != nil {
		Ulog("Error deleting journal allocations for JID = %d, error: %v\n", jid, err)
	}
}

// DeleteJournalEntry deletes the journal record with the supplied jid
func DeleteJournalEntry(jid int64) {
	_, err := RRdb.Prepstmt.DeleteJournalEntry.Exec(jid)
	if err != nil {
		Ulog("Error deleting journal entry for JID = %d, error: %v\n", jid, err)
	}
}

// DeleteJournalMarker deletes the journalmarker record for the supplied jmid
func DeleteJournalMarker(jmid int64) {
	_, err := RRdb.Prepstmt.DeleteJournalMarker.Exec(jmid)
	if err != nil {
		Ulog("Error deleting journal marker for JID = %d, error: %v\n", jmid, err)
	}
}

// DeleteLedgerEntry deletes the ledger record with the supplied lid
func DeleteLedgerEntry(lid int64) error {
	_, err := RRdb.Prepstmt.DeleteLedgerEntry.Exec(lid)
	if err != nil {
		Ulog("Error deleting ledger entry for LID = %d, error: %v\n", lid, err)
	}
	return err
}

// DeleteLedgerMarker deletes the ledgermarker record with the supplied lmid
func DeleteLedgerMarker(lmid int64) error {
	_, err := RRdb.Prepstmt.DeleteLedgerMarker.Exec(lmid)
	if err != nil {
		Ulog("Error deleting ledger marker for LID = %d, error: %v\n", lmid, err)
	}
	return err
}
