package rlib

import "time"

// DeleteJournalAllocations deletes the allocation records associated with the supplied jid
func DeleteJournalAllocations(jid int64) {
	_, err := RRdb.Prepstmt.DeleteJournalAllocations.Exec(jid)
	if err != nil {
		Ulog("Error deleting Journal allocations for JID = %d, error: %v\n", jid, err)
	}
}

// DeleteJournalEntry deletes the Journal record with the supplied jid
func DeleteJournalEntry(jid int64) {
	_, err := RRdb.Prepstmt.DeleteJournalEntry.Exec(jid)
	if err != nil {
		Ulog("Error deleting Journal entry for JID = %d, error: %v\n", jid, err)
	}
}

// DeleteJournalMarker deletes the JournalMarker record for the supplied jmid
func DeleteJournalMarker(jmid int64) {
	_, err := RRdb.Prepstmt.DeleteJournalMarker.Exec(jmid)
	if err != nil {
		Ulog("Error deleting Journal marker for JID = %d, error: %v\n", jmid, err)
	}
}

// DeleteLedgerEntry deletes the Ledger record with the supplied lid
func DeleteLedgerEntry(lid int64) error {
	_, err := RRdb.Prepstmt.DeleteLedgerEntry.Exec(lid)
	if err != nil {
		Ulog("Error deleting Ledger entry for LEID = %d, error: %v\n", lid, err)
	}
	return err
}

// DeleteLedger deletes the Ledger record with the supplied lid
func DeleteLedger(lid int64) error {
	_, err := RRdb.Prepstmt.DeleteLedger.Exec(lid)
	if err != nil {
		Ulog("Error deleting Ledger for LID = %d, error: %v\n", lid, err)
	}
	return err
}

// DeleteLedgerMarker deletes the LedgerMarker record with the supplied lmid
func DeleteLedgerMarker(lmid int64) error {
	_, err := RRdb.Prepstmt.DeleteLedgerMarker.Exec(lmid)
	if err != nil {
		Ulog("Error deleting Ledger marker for LEID = %d, error: %v\n", lmid, err)
	}
	return err
}

// DeleteReceipt deletes the Receipt record with the supplied rcptid
func DeleteReceipt(rcptid int64) error {
	_, err := RRdb.Prepstmt.DeleteReceipt.Exec(rcptid)
	if err != nil {
		Ulog("Error deleting Receipt for RCPTID = %d, error: %v\n", rcptid, err)
	}
	return err
}

// DeleteReceiptAllocations deletes ReceiptAllocation records with the supplied rcptid
func DeleteReceiptAllocations(rcptid int64) error {
	_, err := RRdb.Prepstmt.DeleteReceiptAllocations.Exec(rcptid)
	if err != nil {
		Ulog("Error deleting ReceiptAllocation for RCPTID = %d, error: %v\n", rcptid, err)
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

// DeleteRentableRTID deletes RentableRTID records with the supplied rid, dtstart and dtstop
func DeleteRentableRTID(rid int64, dtstart, dtstop *time.Time) error {
	_, err := RRdb.Prepstmt.DeleteRentableRTID.Exec(rid, dtstart, dtstop)
	if err != nil {
		Ulog("Error deleting RentableRTID with rid=%d, dtstart=%s, dtstop=%s, error: %v\n",
			rid, dtstart.Format(RRDATEINPFMT), dtstop.Format(RRDATEINPFMT), err)
	}
	return err
}

// DeleteRentableStatus deletes RentableStatus records with the supplied rid, dtstart and dtstop
func DeleteRentableStatus(rid int64, dtstart, dtstop *time.Time) error {
	_, err := RRdb.Prepstmt.DeleteRentableStatus.Exec(rid, dtstart, dtstop)
	if err != nil {
		Ulog("Error deleting RentableStatus with rid=%d, dtstart=%s, dtstop=%s, error: %v\n",
			rid, dtstart.Format(RRDATEINPFMT), dtstop.Format(RRDATEINPFMT), err)
	}
	return err
}

// DeleteAgreementPet deletes the pet with the specified petid from the database
func DeleteAgreementPet(petid int64) error {
	_, err := RRdb.Prepstmt.DeleteAgreementPet.Exec(petid)
	if err != nil {
		Ulog("Error deleting petid=%d error: %v\n", petid, err)
	}
	return err
}

// DeleteAllAgreementPets deletes the pet with the specified petid from the database
func DeleteAllAgreementPets(raid int64) error {
	_, err := RRdb.Prepstmt.DeleteAllAgreementPets.Exec(raid)
	if err != nil {
		Ulog("Error deleting pets for rental agreement=%d error: %v\n", raid, err)
	}
	return err
}
