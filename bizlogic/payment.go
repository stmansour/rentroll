package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"time"
)

// Receipts are allocated as follows:
// 1. Produce a list of Payors who have 1 or more Receipts that are not
//    fully allocated
// 2. Select a Payor to handle
// 3. Produce a list of all unpaid assessments for which the Payor is
//    responsible
// 4. Get the total amount of unallocated funds that can be applied to
//    these assessments
// 5. Apply the total amount towards paying off one of the assessments:
// 6. For each assessment:
//        * apply all or part of it to pay an assessment
//        * update the assessment if it is fully or partially paid off
//        * update the receipts used as either fully or partially allocated
// 7. Repeat step 5 until either all funds have been allocated or until
//    there are no more assessments
// 8. Any left over funds will be left available in the Unallocated Funds
//    account to use when more assessments are made.
//
// The routines in this file help perform some of these tasks.

// ROUNDINGERR may not be necessary, but if division becomes a part of the
// floating point calculations below, it is likely that subtracting things
// that should result in zero will not be exactly zero.  So, if the error is
// less than $0.001 we just call it a rounding error and assume that it's 0.
const ROUNDINGERR = float64(0.000999)

// GetAllUnpaidAssessmentsForPayor determines all the Rental Agreements for
// which the supplied Transactant is Payor, then returns a list
// of all unpaid assessments associated with any of those Rental Agreements.
//-----------------------------------------------------------------------------
func GetAllUnpaidAssessmentsForPayor(ctx context.Context, bid, tcid int64, dt *time.Time) ([]rlib.Assessment, error) {
	var a []rlib.Assessment
	var err error

	m, err := rlib.GetRentalAgreementsByPayor(ctx, bid, tcid) // Determine which Rental Agreements the Payor is responsible for...
	if err != nil {
		return a, err
	}
	// rlib.Console("*****************\n\n   GetAllUnpaidAssessmentsForPayor: date = %s, len(m) = %d\n\n****************\n", dt.Format(rlib.RRDATEFMTSQL), len(m))
	for i := 0; i < len(m); i++ {
		// rlib.Console("%d. RAID = %d\n", i, m[i].RAID)
	}
	for i := 0; i < len(m); i++ { // build the list of unpaid assessments
		n, err := rlib.GetUnpaidAssessmentsByRAID(ctx, m[i].RAID) // the list is presorted by Start date ascending
		if err != nil {
			return a, err
		}
		// rlib.Console("Unpaid assessment count for RAID %d: %d\n", m[i].RAID, len(n))
		a = append(a, n...)
	}
	return a, nil
}

// RemainingReceiptFunds returns the amount of funds left to be allocated on
// the supplied receipt
//-----------------------------------------------------------------------------
func RemainingReceiptFunds(ctx context.Context, r *rlib.Receipt) float64 {
	funcname := "RemainingReceiptFunds"
	var xbiz1 rlib.XBusiness
	var dt time.Time
	switch r.FLAGS & 3 {
	case 0:
		return r.Amount
	case 1:
		// compute the amount in
		// TODO(Steve): should we ignore error?
		m, _ := rlib.ParseAcctRule(ctx, &xbiz1, 0, &dt, &dt, r.AcctRuleApply, 0, 1.0)
		tot := r.Amount
		for i := 0; i < len(m); i++ {
			// rlib.Console("%d. %.2f  %s\n", i, m[i].Amount, m[i].Action)
			if "d" == m[i].Action {
				tot -= m[i].Amount
			}
		}
		return tot
	case 2:
		return float64(0)
	default:
		err := fmt.Errorf("unhandled flag bits 0-1 of FLAGS: %d", r.FLAGS&3)
		rlib.LogAndPrintError(funcname, err)
	}
	return float64(0)
}

// RemainingReceiptFundsOnDate returns the amount of funds remaining in a
// receipt on the supplied date
//--------------------------------------------------------------------------
func RemainingReceiptFundsOnDate(ctx context.Context, a *rlib.Receipt, dt *time.Time) float64 {
	// TODO(Steve): should we ignore error?
	m, _ := rlib.GetReceiptAllocationsThroughDate(ctx, a.RCPTID, dt)
	amt := a.Amount
	for i := 0; i < len(m); i++ {
		amt -= m[i].Amount
	}
	return amt
}

// AssessmentUnpaidPortion computes and returns the unpaid portion of an
// assessment.
//--------------------------------------------------------------------------
func AssessmentUnpaidPortion(ctx context.Context, a *rlib.Assessment) float64 {
	funcname := "AssessmentUnpaidPortion"
	switch a.FLAGS & 3 {
	case 0:
		return a.Amount
	case 1:
		// TODO(Steve): should we ignore error?
		ra, _ := rlib.GetReceiptAllocationsByASMID(ctx, a.BID, a.ASMID)
		bal := a.Amount
		for i := 0; i < len(ra); i++ {
			bal -= ra[i].Amount
		}
		return bal
	case 2:
		return float64(0)
	default:
		err := fmt.Errorf("unhandled flag bits 0-1 of FLAGS: %d", a.FLAGS&3)
		rlib.LogAndPrintError(funcname, err)
	}
	return float64(0)
}

// PayAssessment handles paying an assessment, or as much as possible of
// the assessment
//
// @params
//  a      - the assessment being paid
//  rcpt   - the receipt from which funds will be taken to pay the
//           assessment
//  needed - the amount needed to fully pay the assessment.  Its value on
//           return is set to the amount still needed to pay off the
//           assessment.
//  amt    - pointer to the amount to apply toward the assessment.
//           This amount may be less than what is needed -- the remaining
//           funds on a receipt may not always be enough to cover the
//           assessment.  If the supplied receipt has enough funds to cover
//           *amt, then upon return *amt will be 0.00.  If there were not
//           enough funds, *amt will contain the amount still needed to be
//           paid by another receipt.
//  dt     - timestamp to mark on the allocation for this payment
//--------------------------------------------------------------------------
func PayAssessment(ctx context.Context, a *rlib.Assessment, rcpt *rlib.Receipt, needed *float64, amt *float64, dt *time.Time) error {
	funcname := "PayAssessment"

	amtToUse := *amt
	if *amt > *needed {
		amtToUse = *needed
	}

	amtAvailableInRcpt := RemainingReceiptFunds(ctx, rcpt)
	if amtAvailableInRcpt < *amt { // if the amount of funds in this receipt is less than what the user asked to pay...
		amtToUse = amtAvailableInRcpt // ... update the amount to use
	}
	(*amt) -= amtToUse // update the amt for the caller

	//-----------------------
	// pay the assessment
	//-----------------------
	var ra rlib.ReceiptAllocation
	ra.Amount = amtToUse    // this is what can be applied to pay off the assessment
	ra.ASMID = a.ASMID      // this assessment
	ra.BID = a.BID          // this business
	ra.RCPTID = rcpt.RCPTID // bind this allocation the the receipt
	ra.Dt = *dt             // the date is the one supplied to this routine, may be different than the Receipt's date
	ra.RAID = a.RAID        // this Rental Agreement
	// rlib.Console("a.BID = %d, a.ARID = %d\n", a.BID, a.ARID)
	car := rlib.RRdb.BizTypes[a.BID].AR[a.ARID]    // this is the assessment's Account Rule
	dar := rlib.RRdb.BizTypes[a.BID].AR[rcpt.ARID] // debit -- this is the receipt's Account Rule, credit account

	//---------------------------------------------------------------------------------
	// Do not allow the ReceiptAllocation date to be prior to the assessment's date...
	//---------------------------------------------------------------------------------
	if a.Start.After(ra.Dt) {
		ra.Dt = a.Start
	}

	// rlib.Console("Pay Assessment: Assessment Rule:  Debit %s, Credit %s\n", rlib.RRdb.BizTypes[a.BID].GLAccounts[car.DebitLID].Name, rlib.RRdb.BizTypes[a.BID].GLAccounts[car.CreditLID].Name)
	// rlib.Console("Pay Assessment:    Receipt Rule:  Debit %s, Credit %s\n", rlib.RRdb.BizTypes[a.BID].GLAccounts[dar.DebitLID].Name, rlib.RRdb.BizTypes[a.BID].GLAccounts[dar.CreditLID].Name)

	dacct := rlib.RRdb.BizTypes[a.BID].GLAccounts[dar.CreditLID] // we debit what was credited in the Receipt's AcctRuleReceive
	cacct := rlib.RRdb.BizTypes[a.BID].GLAccounts[car.DebitLID]  // we credit what was debited in the Assessments ARID

	ra.AcctRule = fmt.Sprintf("ASM(%d) d %s %.2f,c %s %.2f", a.ASMID, dacct.GLNumber, amtToUse, cacct.GLNumber, amtToUse)
	_, err := rlib.InsertReceiptAllocation(ctx, &ra)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}
	rcpt.RA = append(rcpt.RA, ra)
	// rlib.Console("wrote ReceiptAllocation %d\n", ra.RCPAID)

	//--------------------------------------------------------
	// mark the assessment as either fully or partially paid
	//--------------------------------------------------------
	d := uint64(0x3)
	d = ^d
	a.FLAGS &= d // zero-out bits 0-1
	if *needed-amtToUse < ROUNDINGERR {
		a.FLAGS |= 2 // 2 = paid in full
		// rlib.Console("Fully paid assessment %d\n", a.ASMID)
	} else {
		a.FLAGS |= 1 // 1 = partially paid
		// rlib.Console("Partially paid assessment %d\n", a.ASMID)
	}
	err = rlib.UpdateAssessment(ctx, a)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}
	(*needed) -= ra.Amount
	// rlib.Console("Amount still owed on assessment %d:  %.2f\n", a.ASMID, *needed)

	//------------------------------------------------------------------
	// update the receipt as partially or fully allocated as needed...
	//------------------------------------------------------------------
	rcpt.FLAGS &= 0x7ffffffc // zero-out bits 0-1
	if amtAvailableInRcpt-amtToUse > ROUNDINGERR {
		// rlib.Console("SET RECEIPT FLAGS TO: 1 - some funds remain\n")
		rcpt.FLAGS |= 1 // there are still some funds left */
	} else {
		// rlib.Console("SET RECEIPT FLAGS TO: 2 - funds fully allocated\n")
		rcpt.FLAGS |= 2 // this receipt is now fully allocated
	}
	//------------------------------------------------------------------
	// update the Apply portion of the Receipts account rulse
	//------------------------------------------------------------------
	if len(rcpt.AcctRuleApply) > 0 {
		rcpt.AcctRuleApply += "," + ra.AcctRule
	} else {
		rcpt.AcctRuleApply = ra.AcctRule
	}
	err = rlib.UpdateReceipt(ctx, rcpt)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}
	// rlib.Console("Funds remaining in RCPTID %d = %.2f\n", rcpt.RCPTID, RemainingReceiptFunds(rcpt))

	//-------------------------------------------------------------------------
	// Find the journal entry for this Receipt and add a journal allocation
	// based on the allocation we just did for the receipt
	//-------------------------------------------------------------------------

	// New
	var jnl = rlib.Journal{
		BID:    a.BID,
		Amount: amtToUse,
		Dt:     *dt,
		Type:   rlib.JNLTYPERCPT,
		ID:     rcpt.RCPTID,
	}
	_, err = rlib.InsertJournal(ctx, &jnl)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}

	var ja = rlib.JournalAllocation{
		JID:      jnl.JID,
		AcctRule: ra.AcctRule,
		Amount:   amtToUse,
		BID:      jnl.BID,
		RAID:     a.RAID,
		RID:      a.RID,
		ASMID:    a.ASMID,
		TCID:     rcpt.TCID,
		RCPTID:   rcpt.RCPTID,
	}
	_, err = rlib.InsertJournalAllocationEntry(ctx, &ja)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}
	jnl.JA = append(jnl.JA, ja)

	//-------------------------------------------------------------------------
	// Update ledgers based on journal entry
	//-------------------------------------------------------------------------
	var l = rlib.LedgerEntry{
		BID:    jnl.BID,
		JID:    jnl.JID,
		RID:    ja.RID,
		JAID:   ja.JAID,
		RAID:   ja.RAID,
		TCID:   ja.TCID,
		Dt:     jnl.Dt,
		LID:    dacct.LID,
		Amount: amtToUse,
	}
	_, err = rlib.InsertLedgerEntry(ctx, &l)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}

	l.LID = cacct.LID
	l.Amount = -amtToUse
	_, err = rlib.InsertLedgerEntry(ctx, &l)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}

	return nil
}

// AutoAllocatePayorReceipts applies the amount of the supplied receipt to
// allocate payments to all unpaid based assessments for which the payor is
// responsible.
//
// @params:
//	tcid = TCID of payor
//  dt   = date to be used for allocations
//
// @returns
//  any error encountered
//--------------------------------------------------------------------------
func AutoAllocatePayorReceipts(ctx context.Context, tcid int64, dt *time.Time) error {
	// funcname := "AutoAllocatePayorReceipts"
	// rlib.Console("Entered %s\n", funcname)
	var t rlib.Transactant
	if err := rlib.GetTransactant(ctx, tcid, &t); err != nil {
		rlib.Console("error getting GetTransactant(%d): %s\n", tcid, err.Error())
		return err
	}

	m, err := GetAllUnpaidAssessmentsForPayor(ctx, t.BID, tcid, dt)
	if err != nil {
		return err
	}

	n, err := rlib.GetUnallocatedReceiptsByPayor(ctx, t.BID, tcid)
	if err != nil {
		return err
	}

	// rlib.Console("Unpaid assessments for payor = %s %s: %d\n", t.FirstName, t.LastName, len(m))
	// rlib.Console("The receipts to be allocated to pay the assessments: %d\n\n", len(n))

	//-----------------------------------------------------------------------
	// Begin paying off bills.  For each assessment, go through the receipts,
	// starting with the oldest, use all its funds, then move on to the next
	//-----------------------------------------------------------------------
	for i := 0; i < len(m); i++ {
		// rlib.Console("ASMID = %d, Amount = %.2f, AR = %d\n", m[i].ASMID, m[i].Amount, m[i].ARID)
		for j := 0; j < len(n); j++ {
			if n[j].FLAGS&3 == 2 { // if there are no funds left in this receipt...
				continue // move on to the next receipt
			}
			// First, determine the amount needed for payment...
			needed := AssessmentUnpaidPortion(ctx, &m[i])
			amt := RemainingReceiptFunds(ctx, &n[j])
			// rlib.Console("Needed for ASMID %d :  %.2f\n", m[i].ASMID, needed)
			// rlib.Console("Funds remaining in receipt %d:  %.2f\n", n[j].RCPTID, amt)
			a, err := rlib.GetAssessment(ctx, m[i].ASMID)
			if err != nil {
				return err
			}
			paymentDate := a.Start
			err = PayAssessment(ctx, &m[i], &n[j], &needed, &amt, &paymentDate)
			if err != nil {
				return err
			}
			// rlib.Console(">>>>> Paid assesment %d using receipt %d\n", m[i].ASMID, n[j].RCPTID)
			if needed < ROUNDINGERR { // if we've paid off the assessment...
				break // ... then move on to the next assessment
			}
		}
		// rlib.Console("Completed ASMID = %d\n", m[i].ASMID)
		// rlib.Console("-------------------------------------\n\n")
	}

	return nil
}
