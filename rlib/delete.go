package rlib

import (
	"context"
	"extres"
	"time"
)

// DeleteAR deletes AR records with the supplied id
func DeleteAR(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteAR)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteAR.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting AR for id = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteAssessment deletes Assessment record with the supplied id
func DeleteAssessment(ctx context.Context, asmid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{asmid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteAssessment)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteAssessment.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Assessment for id = %d, error: %v\n", asmid, err)
	}
	return err
}

// DeleteCustomAttribute deletes CustomAttribute records with the supplied id
func DeleteCustomAttribute(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteCustomAttribute)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteCustomAttribute.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting CustomAttribute for id = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteCustomAttributeRef deletes CustomAttributeRef records with the supplied cid
func DeleteCustomAttributeRef(ctx context.Context, elemid, id, cid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{elemid, id, cid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteCustomAttributeRef)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteCustomAttributeRef.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting elemid=%d, id=%d, cid=%d, error: %v\n", elemid, id, cid, err)
	}
	return err
}

// DeleteDemandSource deletes the DemandSource with the specified id from the database
func DeleteDemandSource(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteDemandSource)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteDemandSource.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting DemandSource for SourceSLSID=%d error: %v\n", id, err)
	}
	return err
}

// DeleteDeposit deletes the Deposit associated with the supplied id
// For convenience, this routine calls DeleteDepositPart. The DepositParts are
// tightly bound to the Deposit. If a Deposit is deleted, the parts should be deleted as well.
func DeleteDeposit(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteDeposit)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteDeposit.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Deposit for DID = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteDepository deletes the Depository associated with the supplied id
func DeleteDepository(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteDepository)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteDepository.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Depository where DEPID = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteDepositMethod deletes ALL the DepositMethod associated with the supplied id
func DeleteDepositMethod(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteDepositMethod)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteDepositMethod.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting DepositMethod where DPMID = %d, error: %v\n", id, err)
	}

	return err
}

// DeleteDepositPart deletes ALL the DepositParts associated with the supplied id
func DeleteDepositPart(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteDepositPart)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteDepositPart.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting DepositParts where DID = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteExpense deletes the Expense associated with the supplied id
func DeleteExpense(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteExpense)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteExpense.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Invoice for InvoiceNo = %d, error: %v\n", id, err)
		return err
	}
	return nil
}

// DeleteInvoice deletes the Invoice associated with the supplied id
// For convenience, this routine calls DeleteInvoiceAssessments. The InvoiceAssessments are
// tightly bound to the Invoice. If a Invoice is deleted, the parts should be deleted as well.
// It also updates
func DeleteInvoice(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteInvoice)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteInvoice.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Invoice for InvoiceNo = %d, error: %v\n", id, err)
		return err
	}
	return DeleteInvoiceAssessments(ctx, id)
}

// DeleteInvoiceAssessments deletes ALL the InvoiceAssessments associated with the supplied InvoiceNo
func DeleteInvoiceAssessments(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteInvoiceAssessments)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteInvoiceAssessments.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting InvoiceAssessments where InvoiceNo = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteJournalAllocation deletes the allocation record with the supplied jid
func DeleteJournalAllocation(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteJournalAllocation)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteJournalAllocation.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Journal allocation for JAID = %d, error: %v\n", id, err)
	}

	return err
}

// DeleteJournalAllocations deletes the allocation records associated with the supplied jid
func DeleteJournalAllocations(ctx context.Context, jid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{jid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteJournalAllocations)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteJournalAllocations.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Journal allocations for JID = %d, error: %v\n", jid, err)
	}

	return err
}

// DeleteJournal deletes the Journal record with the supplied jid
func DeleteJournal(ctx context.Context, jid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{jid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteJournal)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteJournal.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Journal entry for JID = %d, error: %v\n", jid, err)
	}

	return err
}

// DeleteJournalMarker deletes the JournalMarker record for the supplied jmid
func DeleteJournalMarker(ctx context.Context, jmid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{jmid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteJournalMarker)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteJournalMarker.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Journal marker for JID = %d, error: %v\n", jmid, err)
	}

	return err
}

// DeleteLedgerEntry deletes the LedgerEntry record with the supplied id
func DeleteLedgerEntry(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteLedgerEntry)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteLedgerEntry.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting LedgerEntry for LEID = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteLedger deletes the GLAccount record with the supplied lid
func DeleteLedger(ctx context.Context, lid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{lid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteLedger)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteLedger.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting GLAccount for LID = %d, error: %v\n", lid, err)
	}
	return err
}

// DeleteLedgerMarker deletes the LedgerMarker record with the supplied lmid
func DeleteLedgerMarker(ctx context.Context, lmid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{lmid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteLedgerMarker)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteLedgerMarker.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting LedgerMarker for LEID = %d, error: %v\n", lmid, err)
	}
	return err
}

// DeleteNote deletes the Note with the supplied id and all its children
// PLEASE USE DeleteNoteAndChildNotes IF POSSIBLE
func DeleteNote(ctx context.Context, nid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var n Note
	err = GetNote(ctx, nid, &n)
	if err != nil {
		return err
	}
	return DeleteNoteAndChildNotes(ctx, &n)
}

// DeleteNoteAndChildNotes deletes supplied Note and all its child notes
func DeleteNoteAndChildNotes(ctx context.Context, p *Note) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	for i := 0; i < len(p.CN); i++ {
		err = DeleteNoteAndChildNotes(ctx, &p.CN[i])
		if err != nil {
			Ulog("Error deleting Note for NID = %d, error: %v\n", p.CN[i].NID, err)
		}
	}
	err = DeleteNoteInternal(ctx, p.NID)
	return err
}

// DeleteNoteInternal deletes the Note record with the supplied nid. Does not look at child notes.
func DeleteNoteInternal(ctx context.Context, nid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{nid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteNote)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteNote.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Note for NID = %d, error: %v\n", nid, err)
	}
	return err
}

// DeleteNoteList deletes the Note record with the supplied nid
func DeleteNoteList(ctx context.Context, nl *NoteList) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	for i := 0; i < len(nl.N); i++ {
		err = DeleteNoteAndChildNotes(ctx, &nl.N[i])
		if err != nil {
			Ulog("Error deleting Note for NID = %d, error: %v\n", nl.N[i].NID, err)
		}
	}
	fields := []interface{}{nl.NLID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteNoteList)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteNoteList.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Note for NID = %d, error: %v\n", nl.NLID, err)
	}
	return err
}

// DeleteNoteType deletes the NoteType record with the supplied nid
func DeleteNoteType(ctx context.Context, nid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{nid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteNoteType)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteNoteType.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting NoteType for NID = %d, error: %v\n", nid, err)
	}
	return err
}

// DeleteRatePlan deletes RatePlan records with the supplied id
func DeleteRatePlan(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRatePlan)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRatePlan.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RatePlan for id = %d, error: %v\n", id, err)
	}
	return err
}

// DeletePaymentType deletes PaymentType records with the supplied id
func DeletePaymentType(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeletePaymentType)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeletePaymentType.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting PaymentType for id = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteRatePlanRef deletes RatePlanRef records with the supplied cid
func DeleteRatePlanRef(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRatePlanRef)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRatePlanRef.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteRatePlanRefRTRate deletes RatePlanRefRTRate records with the supplied cid
func DeleteRatePlanRefRTRate(ctx context.Context, rtrid, rtid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rtrid, rtid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRatePlanRefRTRate)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRatePlanRefRTRate.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting rtrid=%d rtid=%d error: %v\n", rtrid, rtid, err)
	}
	return err
}

// DeleteRatePlanRefSPRate deletes RatePlanRefSPRate records with the supplied cid
func DeleteRatePlanRefSPRate(ctx context.Context, rtrid, rspid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rtrid, rspid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRatePlanRefSPRate)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRatePlanRefSPRate.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting rtrid=%d rspid=%d error: %v\n", rtrid, rspid, err)
	}
	return err
}

// DeleteReceipt deletes the Receipt record with the supplied rcptid
func DeleteReceipt(ctx context.Context, rcptid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rcptid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteReceipt)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteReceipt.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Receipt for RCPTID = %d, error: %v\n", rcptid, err)
	}
	return err
}

// DeleteReceiptAllocation deletes the ReceiptAllocation record with the supplied id
func DeleteReceiptAllocation(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteReceiptAllocation)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteReceiptAllocation.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting ReceiptAllocation for RCPAID = %d, error: %v\n", id, err)
	}
	return err
}

// DeleteReceiptAllocations deletes ReceiptAllocation records with the supplied rcptid
func DeleteReceiptAllocations(ctx context.Context, rcptid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rcptid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteReceiptAllocations)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteReceiptAllocations.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting ReceiptAllocations for RCPTID = %d, error: %v\n", rcptid, err)
	}
	return err
}

// DeleteRentableTypeRefWithRTID deletes RentableTypeRef records with the supplied RTID
func DeleteRentableTypeRefWithRTID(ctx context.Context, rtid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rtid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentableTypeRefWithRTID)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentableTypeRefWithRTID.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RentableTypeRef with rtid=%d\n", rtid, err)
	}
	return err
}

// DeleteRentableTypeRef deletes RentableTypeRef records with the supplied rtrid
func DeleteRentableTypeRef(ctx context.Context, rtrid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rtrid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentableTypeRef)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentableTypeRef.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RentableTypeRef with rtrid=%d\n", rtrid, err)
	}
	return err
}

// DeleteRentableMarketRateInstance deletes RentableMarketRate instance with given RMRID
func DeleteRentableMarketRateInstance(ctx context.Context, rmrid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rmrid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentableMarketRateInstance)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentableMarketRateInstance.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RentableMarketRate with rmrid=%d, error: %v\n", rmrid, err)
	}
	return err
}

// DeleteRentableSpecialtyRef deletes RentableSpecialtyRef records with the supplied rid, dtstart and dtstop
func DeleteRentableSpecialtyRef(ctx context.Context, rid int64, dtstart, dtstop *time.Time) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rid, dtstart, dtstop}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentableSpecialtyRef)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentableSpecialtyRef.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RentableSpecialtyRef with rid=%d, dtstart=%s, dtstop=%s, error: %v\n",
			rid, dtstart.Format(RRDATEINPFMT), dtstop.Format(RRDATEINPFMT), err)
	}
	return err
}

// DeleteRentableStatus deletes RentableStatus records with the supplied rsid
func DeleteRentableStatus(ctx context.Context, rsid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rsid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentableStatus)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentableStatus.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RentableStatus with rsid=%d\n", rsid, err)
	}
	return err
}

// DeleteRentalAgreementPayor deletes the Payor with the specified id from the database
func DeleteRentalAgreementPayor(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentalAgreementPayor)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentalAgreementPayor.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RAPID=%d error: %v\n", id, err)
	}
	return err
}

// DeleteRentalAgreementPayorByRBT deletes the payor from the RentalAgreement
func DeleteRentalAgreementPayorByRBT(ctx context.Context, raid, bid, tcid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{raid, bid, tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentalAgreementPayorByRBT)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentalAgreementPayorByRBT.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting raid=%d, bid=%d, tcid=%d error: %s\n", raid, bid, tcid, err.Error())
	}
	return err
}

// DeleteRentableUserByRBT deletes the payor from the RentalAgreement
func DeleteRentableUserByRBT(ctx context.Context, rid, bid, tcid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{rid, bid, tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentableUserByRBT)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentableUserByRBT.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting rid=%d, bid=%d, tcid=%d error: %s\n", rid, bid, tcid, err.Error())
	}
	return err
}

// DeleteRentalAgreementPet deletes the pet with the specified petid from the database
func DeleteRentalAgreementPet(ctx context.Context, petid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{petid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentalAgreementPet)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentalAgreementPet.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting petid=%d error: %v\n", petid, err)
	}
	return err
}

// DeleteRentalAgreementRentable deletes the rentable with the specified id from the database
func DeleteRentalAgreementRentable(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentalAgreementRentable)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentalAgreementRentable.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteRentalAgreement deletes the rentable with the specified id from the database
func DeleteRentalAgreement(ctx context.Context, raid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentalAgreement)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentalAgreement.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RentalAgreement with raid=%d error: %v\n", raid, err)
	}
	return err
}

// DeleteAllRentalAgreementRentables deletes all pets associated with the specified raid
func DeleteAllRentalAgreementRentables(ctx context.Context, raid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteAllRentalAgreementRentables)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteAllRentalAgreementRentables.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Rentables for rental agreement=%d error: %v\n", raid, err)
	}
	return err
}

// DeleteAllRentalAgreementPayors deletes all pets associated with the specified raid
func DeleteAllRentalAgreementPayors(ctx context.Context, raid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteAllRentalAgreementPayors)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteAllRentalAgreementPayors.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Payors for rental agreement=%d error: %v\n", raid, err)
	}
	return err
}

// DeleteAllRentalAgreementPets deletes all pets associated with the specified raid
func DeleteAllRentalAgreementPets(ctx context.Context, raid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteAllRentalAgreementPets)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteAllRentalAgreementPets.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting pets for rental agreement=%d error: %v\n", raid, err)
	}
	return err
}

// DeleteRentableUser deletes the User with the specified id from the database
func DeleteRentableUser(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteRentableUser)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteRentableUser.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting RUID=%d error: %v\n", id, err)
	}
	return err
}

// DeleteStringList deletes the StringList with the specified id from the database
func DeleteStringList(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	err = DeleteSLStrings(ctx, id)
	if err != nil {
		return err
	}
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteStringList)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteStringList.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteSLString deletes the SLString with the specified id from the database
func DeleteSLString(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteSLString)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteSLString.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting SLString id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteSLStrings deletes all SLString with the specified SLID from the database
func DeleteSLStrings(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	if id > 0 {
		fields := []interface{}{id}
		if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
			stmt := tx.Stmt(RRdb.Prepstmt.DeleteSLStrings)
			defer stmt.Close()
			_, err = stmt.Exec(fields...)
		} else {
			_, err = RRdb.Prepstmt.DeleteSLStrings.Exec(fields...)
		}
		if err != nil {
			Ulog("Error deleting id=%d error: %v\n", id, err)
		}
	}
	return err
}

// DeleteSubAR deletes the SubAR with the specified id from the database
func DeleteSubAR(ctx context.Context, sarid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{sarid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteSubAR)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteSubAR.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting SubAR id=%d error: %v\n", sarid, err)
	}
	return err
}

// DeleteSubARs deletes all SubAR with the specified SLID from the database
func DeleteSubARs(ctx context.Context, arid int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	if arid > 0 {
		fields := []interface{}{arid}
		if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
			stmt := tx.Stmt(RRdb.Prepstmt.DeleteSubARs)
			defer stmt.Close()
			_, err = stmt.Exec(fields...)
		} else {
			_, err = RRdb.Prepstmt.DeleteSubARs.Exec(fields...)
		}
		if err != nil {
			Ulog("Error deleting ARID=%d error: %v\n", arid, err)
		}
	}
	return err
}

func delContextProblem(ctx context.Context) bool {
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return true
		}
	}
	return false
}

//*****************************************************************
// TASK, TASKLIST, TASK LIST DEFINITION, TASK LIST DESCRIPTOR
//*****************************************************************

// DeleteTask deletes the Task with the specified id from the database
func DeleteTask(ctx context.Context, id int64) error {
	var err error
	if delContextProblem(ctx) {
		return ErrSessionRequired
	}
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteTask)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteTask.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Task id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteTaskList deletes the TaskList with the specified id from the database
func DeleteTaskList(ctx context.Context, id int64) error {
	var err error
	if delContextProblem(ctx) {
		return ErrSessionRequired
	}
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteTaskList)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteTaskList.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting TaskList id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteTaskListTasks deletes the Tasks tied to the TaskList with
// the specified id from the database
func DeleteTaskListTasks(ctx context.Context, id int64) error {
	var err error
	if delContextProblem(ctx) {
		return ErrSessionRequired
	}
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteTaskListTasks)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteTaskListTasks.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Tasks id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteTaskDescriptor deletes the TaskDescriptor with the specified id from the database
func DeleteTaskDescriptor(ctx context.Context, id int64) error {
	var err error
	if delContextProblem(ctx) {
		return ErrSessionRequired
	}
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteTaskDescriptor)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteTaskDescriptor.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting TaskDescriptor id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteTaskListDefinition deletes the TaskListDefinition with the specified id from the database
func DeleteTaskListDefinition(ctx context.Context, id int64) error {
	var err error
	if delContextProblem(ctx) {
		return ErrSessionRequired
	}
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteTaskListDefinition)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteTaskListDefinition.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting TaskListDefinition id=%d error: %v\n", id, err)
	}
	return err
}

//*****************************************************************************
//  TRANSACTANT, PAYOR, USER, PROSPECT
//*****************************************************************************

// DeleteTransactant deletes the Transactant with the specified id from the database
func DeleteTransactant(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteTransactant)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteTransactant.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Transactant id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteUser deletes the User with the specified id from the database
func DeleteUser(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteUser)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteUser.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting User id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteProspect deletes the Prospect with the specified id from the database
func DeleteProspect(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteProspect)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteProspect.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Prospect id=%d error: %v\n", id, err)
	}
	return err
}

// DeletePayor deletes the Payor with the specified id from the database
func DeletePayor(ctx context.Context, id int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeletePayor)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeletePayor.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting Payor id=%d error: %v\n", id, err)
	}
	return err
}

// DeleteFlow deletes a flow with the given FlowID
func DeleteFlow(ctx context.Context, FlowID int64) error {
	var err error

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	fields := []interface{}{FlowID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.DeleteFlow)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.DeleteFlow.Exec(fields...)
	}
	if err != nil {
		Ulog("Error deleting FlowParts for FlowID = %d, error: %v\n", FlowID, err)
	}
	return err
}
