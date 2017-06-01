package rlib

import "time"

// DeleteAccountDepository deletes AccountDepository records with the supplied id
func DeleteAccountDepository(id int64) error {
	_, err := RRdb.Prepstmt.DeleteAccountDepository.Exec(id)
	if err != nil {
		Ulog("Error deleting AccountDepository for id = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteAR deletes AR records with the supplied id
func DeleteAR(id int64) error {
	_, err := RRdb.Prepstmt.DeleteAR.Exec(id)
	if err != nil {
		Ulog("Error deleting AR for id = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteCustomAttribute deletes CustomAttribute records with the supplied id
func DeleteCustomAttribute(id int64) error {
	_, err := RRdb.Prepstmt.DeleteCustomAttribute.Exec(id)
	if err != nil {
		Ulog("Error deleting CustomAttribute for id = %d, error: %v\n", id, err)
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

// DeleteDemandSource deletes the DemandSource with the specified id from the database
func DeleteDemandSource(id int64) error {
	_, err := RRdb.Prepstmt.DeleteDemandSource.Exec(id)
	if err != nil {
		Ulog("Error deleting DemandSource for SourceSLSID=%d error: %v\n", id, err)
	}
	return err
}

// DeleteDeposit deletes the Deposit associated with the supplied id
// For convenience, this routine calls DeleteDepositParts. The DepositParts are
// tightly bound to the Deposit. If a Deposit is deleted, the parts should be deleted as well.
func DeleteDeposit(id int64) {
	_, err := RRdb.Prepstmt.DeleteDeposit.Exec(id)
	if err != nil {
		Ulog("Error deleting Deposit for DID = %d, error: %v\n", id, err)
	}
}

// DeleteDepository deletes the Depository associated with the supplied id
func DeleteDepository(id int64) error {
	_, err := RRdb.Prepstmt.DeleteDepository.Exec(id)
	if err != nil {
		Ulog("Error deleting Depository where DEPID = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteDepositMethod deletes ALL the DepositMethod associated with the supplied id
func DeleteDepositMethod(id int64) {
	_, err := RRdb.Prepstmt.DeleteDepositMethod.Exec(id)
	if err != nil {
		Ulog("Error deleting DepositMethod where DPMID = %d, error: %v\n", id, err)
	}
}

// DeleteDepositParts deletes ALL the DepositParts associated with the supplied id
func DeleteDepositParts(id int64) {
	_, err := RRdb.Prepstmt.DeleteDepositParts.Exec(id)
	if err != nil {
		Ulog("Error deleting DepositParts where DID = %d, error: %v\n", id, err)
	}
}

// DeleteInvoice deletes the Invoice associated with the supplied id
// For convenience, this routine calls DeleteInvoiceAssessments. The InvoiceAssessments are
// tightly bound to the Invoice. If a Invoice is deleted, the parts should be deleted as well.
// It also updates
func DeleteInvoice(id int64) error {
	_, err := RRdb.Prepstmt.DeleteInvoice.Exec(id)
	if err != nil {
		Ulog("Error deleting Invoice for InvoiceNo = %d, error: %v\n", id, err)
		return err
	}
	return DeleteInvoiceAssessments(id)
}

// DeleteInvoiceAssessments deletes ALL the InvoiceAssessments associated with the supplied InvoiceNo
func DeleteInvoiceAssessments(id int64) error {
	_, err := RRdb.Prepstmt.DeleteInvoiceAssessments.Exec(id)
	if err != nil {
		Ulog("Error deleting InvoiceAssessments where InvoiceNo = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteJournalAllocation deletes the allocation record with the supplied jid
func DeleteJournalAllocation(id int64) {
	_, err := RRdb.Prepstmt.DeleteJournalAllocation.Exec(id)
	if err != nil {
		Ulog("Error deleting Journal allocation for JAID = %d, error: %v\n", id, err)
	}
}

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

// DeleteLedgerEntry deletes the LedgerEntry record with the supplied lid
func DeleteLedgerEntry(lid int64) error {
	_, err := RRdb.Prepstmt.DeleteLedgerEntry.Exec(lid)
	if err != nil {
		Ulog("Error deleting LedgerEntry for LEID = %d, error: %v\n", lid, err)
	}
	return err
}

// DeleteLedger deletes the GLAccount record with the supplied lid
func DeleteLedger(lid int64) error {
	_, err := RRdb.Prepstmt.DeleteLedger.Exec(lid)
	if err != nil {
		Ulog("Error deleting GLAccount for LID = %d, error: %v\n", lid, err)
	}
	return err
}

// DeleteLedgerMarker deletes the LedgerMarker record with the supplied lmid
func DeleteLedgerMarker(lmid int64) error {
	_, err := RRdb.Prepstmt.DeleteLedgerMarker.Exec(lmid)
	if err != nil {
		Ulog("Error deleting LedgerMarker for LEID = %d, error: %v\n", lmid, err)
	}
	return err
}

// DeleteNote deletes the Note with the supplied id and all its children
// PLEASE USE DeleteNoteAndChildNotes IF POSSIBLE
func DeleteNote(nid int64) error {
	var n Note
	GetNote(nid, &n)
	return DeleteNoteAndChildNotes(&n)
}

// DeleteNoteAndChildNotes deletes supplied Note and all its child notes
func DeleteNoteAndChildNotes(p *Note) error {
	for i := 0; i < len(p.CN); i++ {
		err := DeleteNoteAndChildNotes(&p.CN[i])
		if err != nil {
			Ulog("Error deleting Note for NID = %d, error: %v\n", p.CN[i].NID, err)
		}
	}
	err := DeleteNoteInternal(p.NID)
	return err
}

// DeleteNoteInternal deletes the Note record with the supplied nid. Does not look at child notes.
func DeleteNoteInternal(nid int64) error {
	_, err := RRdb.Prepstmt.DeleteNote.Exec(nid)
	if err != nil {
		Ulog("Error deleting Note for NID = %d, error: %v\n", nid, err)
	}
	return err
}

// DeleteNoteList deletes the Note record with the supplied nid
func DeleteNoteList(nl *NoteList) error {
	for i := 0; i < len(nl.N); i++ {
		err := DeleteNoteAndChildNotes(&nl.N[i])
		if err != nil {
			Ulog("Error deleting Note for NID = %d, error: %v\n", nl.N[i].NID, err)
		}
	}
	_, err := RRdb.Prepstmt.DeleteNoteList.Exec(nl.NLID)
	if err != nil {
		Ulog("Error deleting Note for NID = %d, error: %v\n", nl.NLID, err)
	}
	return err
}

// DeleteNoteType deletes the NoteType record with the supplied nid
func DeleteNoteType(nid int64) error {
	_, err := RRdb.Prepstmt.DeleteNoteType.Exec(nid)
	if err != nil {
		Ulog("Error deleting NoteType for NID = %d, error: %v\n", nid, err)
	}
	return err
}

// DeleteRatePlan deletes RatePlan records with the supplied id
func DeleteRatePlan(id int64) error {
	_, err := RRdb.Prepstmt.DeleteRatePlan.Exec(id)
	if err != nil {
		Ulog("Error deleting RatePlan for id = %d, error: %v\n", id, err)
	}
	return err
}

// DeletePaymentType deletes PaymentType records with the supplied id
func DeletePaymentType(id int64) error {
	_, err := RRdb.Prepstmt.DeletePaymentType.Exec(id)
	if err != nil {
		Ulog("Error deleting PaymentType for id = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteRatePlanRef deletes RatePlanRef records with the supplied cid
func DeleteRatePlanRef(id int64) error {
	_, err := RRdb.Prepstmt.DeleteRatePlanRef.Exec(id)
	if err != nil {
		Ulog("Error deleting id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteRatePlanRefRTRate deletes RatePlanRefRTRate records with the supplied cid
func DeleteRatePlanRefRTRate(rtrid, rtid int64) error {
	_, err := RRdb.Prepstmt.DeleteRatePlanRefRTRate.Exec(rtrid, rtid)
	if err != nil {
		Ulog("Error deleting rtrid=%d rtid=%d error: %v\n", rtrid, rtid, err)
	}
	return err
}

// DeleteRatePlanRefSPRate deletes RatePlanRefSPRate records with the supplied cid
func DeleteRatePlanRefSPRate(rtrid, rspid int64) error {
	_, err := RRdb.Prepstmt.DeleteRatePlanRefSPRate.Exec(rtrid, rspid)
	if err != nil {
		Ulog("Error deleting rtrid=%d rspid=%d error: %v\n", rtrid, rspid, err)
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

// DeleteReceiptAllocation deletes the ReceiptAllocation record with the supplied id
func DeleteReceiptAllocation(id int64) error {
	_, err := RRdb.Prepstmt.DeleteReceiptAllocation.Exec(id)
	if err != nil {
		Ulog("Error deleting ReceiptAllocation for RCPAID = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteReceiptAllocations deletes ReceiptAllocation records with the supplied rcptid
func DeleteReceiptAllocations(rcptid int64) error {
	_, err := RRdb.Prepstmt.DeleteReceiptAllocations.Exec(rcptid)
	if err != nil {
		Ulog("Error deleting ReceiptAllocations for RCPTID = %d, error: %v\n", rcptid, err)
	}
	return err
}

// DeleteRentableTypeRef deletes RentableTypeRef records with the supplied rid, dtstart and dtstop
func DeleteRentableTypeRef(rid int64, dtstart, dtstop *time.Time) error {
	_, err := RRdb.Prepstmt.DeleteRentableTypeRef.Exec(rid, dtstart, dtstop)
	if err != nil {
		Ulog("Error deleting RentableTypeRef with rid=%d, dtstart=%s, dtstop=%s, error: %v\n",
			rid, dtstart.Format(RRDATEINPFMT), dtstop.Format(RRDATEINPFMT), err)
	}
	return err
}

// DeleteRentableSpecialtyRef deletes RentableSpecialtyRef records with the supplied rid, dtstart and dtstop
func DeleteRentableSpecialtyRef(rid int64, dtstart, dtstop *time.Time) error {
	_, err := RRdb.Prepstmt.DeleteRentableSpecialtyRef.Exec(rid, dtstart, dtstop)
	if err != nil {
		Ulog("Error deleting RentableSpecialtyRef with rid=%d, dtstart=%s, dtstop=%s, error: %v\n",
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

// DeleteRentalAgreementPayor deletes the Payor with the specified id from the database
func DeleteRentalAgreementPayor(id int64) error {
	_, err := RRdb.Prepstmt.DeleteRentalAgreementPayor.Exec(id)
	if err != nil {
		Ulog("Error deleting RAPID=%d error: %v\n", id, err)
	}
	return err
}

// DeleteRentalAgreementPayorByRBT deletes the payor from the RentalAgreement
func DeleteRentalAgreementPayorByRBT(raid, bid, tcid int64) error {
	_, err := RRdb.Prepstmt.DeleteRentalAgreementPayorByRBT.Exec(raid, bid, tcid)
	if err != nil {
		Ulog("Error deleting raid=%d, bid=%d, tcid=%d error: %s\n", raid, bid, tcid, err.Error())
	}
	return err
}

// DeleteRentableUserByRBT deletes the payor from the RentalAgreement
func DeleteRentableUserByRBT(rid, bid, tcid int64) error {
	_, err := RRdb.Prepstmt.DeleteRentableUserByRBT.Exec(rid, bid, tcid)
	if err != nil {
		Ulog("Error deleting rid=%d, bid=%d, tcid=%d error: %s\n", rid, bid, tcid, err.Error())
	}
	return err
}

// DeleteRentalAgreementPet deletes the pet with the specified petid from the database
func DeleteRentalAgreementPet(petid int64) error {
	_, err := RRdb.Prepstmt.DeleteRentalAgreementPet.Exec(petid)
	if err != nil {
		Ulog("Error deleting petid=%d error: %v\n", petid, err)
	}
	return err
}

// DeleteRentalAgreementRentable deletes the rentable with the specified id from the database
func DeleteRentalAgreementRentable(id int64) error {
	_, err := RRdb.Prepstmt.DeleteRentalAgreementRentable.Exec(id)
	if err != nil {
		Ulog("Error deleting id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteAllRentalAgreementPets deletes all pets associated with the specified raid
func DeleteAllRentalAgreementPets(id int64) error {
	_, err := RRdb.Prepstmt.DeleteAllRentalAgreementPets.Exec(id)
	if err != nil {
		Ulog("Error deleting pets for rental agreement=%d error: %v\n", id, err)
	}
	return err
}

// DeleteRentableUser deletes the User with the specified id from the database
func DeleteRentableUser(id int64) error {
	_, err := RRdb.Prepstmt.DeleteRentableUser.Exec(id)
	if err != nil {
		Ulog("Error deleting RUID=%d error: %v\n", id, err)
	}
	return err
}

// DeleteStringList deletes the StringList with the specified id from the database
func DeleteStringList(id int64) error {
	err := DeleteSLStrings(id)
	if err != nil {
		if !IsSQLNoResultsError(err) {
			return err
		}
	}
	_, err = RRdb.Prepstmt.DeleteStringList.Exec(id)
	if err != nil {
		Ulog("Error deleting id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteSLString deletes the SLString with the specified id from the database
func DeleteSLString(id int64) error {
	_, err := RRdb.Prepstmt.DeleteSLString.Exec(id)
	if err != nil {
		Ulog("Error deleting SLString id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteSLStrings deletes all SLString with the specified SLID from the database
func DeleteSLStrings(id int64) error {
	var err error
	if id > 0 {
		_, err = RRdb.Prepstmt.DeleteSLStrings.Exec(id)
		if err != nil {
			if !IsSQLNoResultsError(err) {
				Ulog("Error deleting id=%d error: %v\n", id, err)
			}
		}
	}
	return err
}

// DeleteTransactant deletes the Transactant with the specified id from the database
func DeleteTransactant(id int64) error {
	_, err := RRdb.Prepstmt.DeleteTransactant.Exec(id)
	if err != nil {
		Ulog("Error deleting Transactant id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteUser deletes the User with the specified id from the database
func DeleteUser(id int64) error {
	_, err := RRdb.Prepstmt.DeleteUser.Exec(id)
	if err != nil {
		Ulog("Error deleting User id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteProspect deletes the Prospect with the specified id from the database
func DeleteProspect(id int64) error {
	_, err := RRdb.Prepstmt.DeleteProspect.Exec(id)
	if err != nil {
		Ulog("Error deleting Prospect id=%d error: %v\n", id, err)
	}
	return err
}

// DeletePayor deletes the Payor with the specified id from the database
func DeletePayor(id int64) error {
	_, err := RRdb.Prepstmt.DeletePayor.Exec(id)
	if err != nil {
		Ulog("Error deleting Payor id=%d error: %v\n", id, err)
	}
	return err
}
