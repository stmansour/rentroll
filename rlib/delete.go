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

// DeleteReceipt deletes the Receipt record with the supplied rcptid
func DeleteReceipt(rcptid int64) error {
	_, err := RRdb.Prepstmt.DeleteReceipt.Exec(rcptid)
	if err != nil {
		Ulog("Error deleting receipt for RCPTID = %d, error: %v\n", rcptid, err)
	}
	return err
}

// DeleteReceiptAllocations deletes ReceiptAllocation records with the supplied rcptid
func DeleteReceiptAllocations(rcptid int64) error {
	_, err := RRdb.Prepstmt.DeleteReceiptAllocations.Exec(rcptid)
	if err != nil {
		Ulog("Error deleting receiptallocation for RCPTID = %d, error: %v\n", rcptid, err)
	}
	return err
}

// DeleteCustomAttribute deletes CustomAttribute records with the supplied cid
func DeleteCustomAttribute(cid int64) error {
	_, err := RRdb.Prepstmt.DeleteCustomAttribute.Exec(cid)
	if err != nil {
		Ulog("Error deleting CustomAttribute for cid = %d, error: %v\n", cid, err)
	}
	return err
}

// DeleteCustomAttributeRef deletes CustomAttributeRef records with the supplied cid
func DeleteCustomAttributeRef(elemid, id, cid int64) error {
	_, err := RRdb.Prepstmt.DeleteCustomAttributeRef.Exec(elemid, id, cid)
	if err != nil {
		Ulog("Error deleting elemid=%d, id=%d, cid=%d, error: %v\n", elemid, id, cid, err)
	}
	return err
}
