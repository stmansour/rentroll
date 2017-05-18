package bizlogic

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// Receipts are allocated as follows:
// 1. Produce a list of Payors who have 1 or more Receipts that are not fully allocated
// 2. Select a Payor to handle
// 3. Produce a list of all unpaid assessments for which the Payor is responsible
// 4. Get the total amount of unallocated funds that can be applied to these assessments
// 5. Apply the total amount towards paying off one of the assessments:
// 6. For each assessment:
//        * apply all or part of it to pay an assessment
//        * update the assessment if it is fully or partially paid off
//        * update the receipts used as either fully or partially allocated
// 7. Repeat step 5 until either all funds have been allocated or until there are
//    no more assessments
// 8. Any left over funds will be left available in the Unallocated Funds account to
//    use when more assessments are made.
//
// The routines in this file help perform some of these tasks.

// ROUNDINGERR may not be necessary, but if division becomes a part of the floating
// point calculations below, it is likely that subtracting things that should result
// in zero will not be exactly zero.  So, if the error is less than $0.01 we just call
// it a rounding error and assume that it's 0.
const ROUNDINGERR = float64(0.00999)

// GetAllUnpaidAssessmentsForPayor determines all the Rental Agreements for which
// the supplied Transactant is Payor at time dt, then returns a list of all unpaid
// assessments associated with these Rental Agreements.
func GetAllUnpaidAssessmentsForPayor(bid, tcid int64, dt *time.Time) []rlib.Assessment {
	var a []rlib.Assessment
	m := rlib.GetRentalAgreementsByPayor(bid, tcid, dt) // Determine which Rental Agreements the Payor is responsible for...
	for i := 0; i < len(m); i++ {                       // build the list of unpaid assessments
		n := rlib.GetUnpaidAssessmentsByRAID(m[i].RAID) // the list is presorted by Start date ascending
		a = append(a, n...)
	}
	return a
}

// RemainingReceiptFunds returns the amount of funds left to be allocated on the supplied receipt
func RemainingReceiptFunds(xbiz *rlib.XBusiness, r *rlib.Receipt) float64 {
	funcname := "RemainingReceiptFunds"
	var dt time.Time
	switch r.FLAGS & 3 {
	case 0:
		return r.Amount
	case 1:
		// compute the amount in
		m := rlib.ParseAcctRule(xbiz, 0, &dt, &dt, r.AcctRuleApply, 0, 1.0)
		tot := r.Amount
		for i := 0; i < len(m); i++ {
			//fmt.Printf("%d. %.2f  %s\n", i, m[i].Amount, m[i].Action)
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

// AssessmentUnpaidPortion computes and returns the unpaid portion of an
// assessment.
func AssessmentUnpaidPortion(a *rlib.Assessment) float64 {
	funcname := "AssessmentUnpaidPortion"
	switch a.FLAGS & 3 {
	case 0:
		return a.Amount
	case 1:
		ra := rlib.GetReceiptAllocationsByASMID(a.BID, a.ASMID)
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

// payAssessment
// @params
//  xbiz
//  a      - the assessment being paid
//  rcpt   - the receipt from which funds will be taken to pay the assessment
//  needed - the amount needed to fully pay the assessment
//  amt    - the amount that will be taken from the receipt to apply toward the assessment.
//           this amount may be less than needed -- the remaining funds on a receipt may not
//           always be enough to cover the assessment.
//  dt     - timestamp to mark on the allocation for this payment
func payAssessment(xbiz *rlib.XBusiness, a *rlib.Assessment, rcpt *rlib.Receipt, needed *float64, amt float64, dt *time.Time) error {
	funcname := "payAssessment"

	amtToUse := amt
	if amt > *needed {
		amtToUse = *needed
	}

	//-----------------------
	// pay the assessment
	//-----------------------
	var ra rlib.ReceiptAllocation
	ra.Amount = amtToUse                           // this is what can be applied to pay off the assessment
	ra.ASMID = a.ASMID                             // this assessment
	ra.BID = a.BID                                 // this business
	ra.RCPTID = rcpt.RCPTID                        // bind this allocation the the receipt
	ra.Dt = *dt                                    // the date is the one supplied to this routine, may be different than the Receipt's date
	car := rlib.RRdb.BizTypes[a.BID].AR[a.ARID]    // this is the assessment's Account Rule
	dar := rlib.RRdb.BizTypes[a.BID].AR[rcpt.ARID] // debit -- this is the receipt's Account Rule, credit account
	dacct := rlib.RRdb.BizTypes[a.BID].GLAccounts[dar.CreditLID]
	cacct := rlib.RRdb.BizTypes[a.BID].GLAccounts[car.DebitLID]
	ra.AcctRule = fmt.Sprintf("ASM(%d) d %s %.2f,c %s %.2f", a.ASMID, dacct.GLNumber, amtToUse, cacct.GLNumber, amtToUse)
	_, err := rlib.InsertReceiptAllocation(&ra)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}
	rcpt.RA = append(rcpt.RA, ra)
	// fmt.Printf("wrote ReceiptAllocation %d\n", ra.RCPAID)

	//--------------------------------------------------------
	// mark the assessment as either fully or partially paid
	//--------------------------------------------------------
	a.FLAGS &= 0x7ffffffc // zero-out bits 0-1
	if *needed-amtToUse < ROUNDINGERR {
		a.FLAGS |= 2 // 2 = paid in full
		fmt.Printf("Fully paid assessment %d\n", a.ASMID)
	} else {
		a.FLAGS |= 1 // 1 = partially paid
		fmt.Printf("Partially paid assessment %d\n", a.ASMID)
	}
	err = rlib.UpdateAssessment(a)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}
	(*needed) -= ra.Amount
	fmt.Printf("Amount still owed on assessment %d:  %.2f\n", a.ASMID, *needed)

	//------------------------------------------------------------------
	// update the receipt as partially or fully allocated as needed...
	//------------------------------------------------------------------
	rcpt.FLAGS &= 0x7ffffffc // zero-out bits 0-1
	if amt-amtToUse > ROUNDINGERR {
		fmt.Printf("SET RECEIPT FLAGS TO: 1 - some funds remain\n")
		rcpt.FLAGS |= 1 // there are still some funds left */
	} else {
		fmt.Printf("SET RECEIPT FLAGS TO: 2 - funds fully allocated\n")
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
	err = rlib.UpdateReceipt(rcpt)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}
	fmt.Printf("Funds remaining in RCPTID %d = %.2f\n", rcpt.RCPTID, RemainingReceiptFunds(xbiz, rcpt))

	//-------------------------------------------------------------------------
	// Find the journal entry for this Receipt and add a journal allocation
	// based on the allocation we just did for the receipt
	//-------------------------------------------------------------------------
	jnl := rlib.GetJournalByReceiptID(rcpt.RCPTID)
	rlib.GetJournalAllocations(&jnl)
	var ja rlib.JournalAllocation
	ja.JID = jnl.JID
	ja.AcctRule = ra.AcctRule
	ja.Amount = amtToUse
	ja.BID = jnl.BID
	ja.RAID = a.RAID
	ja.RID = a.RID
	ja.TCID = rcpt.TCID
	rlib.InsertJournalAllocationEntry(&ja)
	jnl.JA = append(jnl.JA, ja)

	//-------------------------------------------------------------------------
	// Update ledgers based on journal entry
	//-------------------------------------------------------------------------
	var l rlib.LedgerEntry
	l.BID = jnl.BID
	l.JID = jnl.JID
	l.RID = ja.RID
	l.JAID = ja.JAID
	l.RAID = ja.RAID
	l.TCID = ja.TCID
	l.Dt = jnl.Dt
	l.LID = dacct.LID
	l.Amount = amtToUse
	_, err = rlib.InsertLedgerEntry(&l)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}

	l.LID = cacct.LID
	l.Amount = -amtToUse
	_, err = rlib.InsertLedgerEntry(&l)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}

	return nil
}

// AutoAllocatePayorReceipts applies the amount of the supplied receipt to allocate
// payments to all unpaid based assessments for which the payor is
// responsible.
// @params:
//  xbiz = XBusiness struct for the associated business
//	tcid = TCID of payor
//  dt   = date to be used for allocations
func AutoAllocatePayorReceipts(xbiz *rlib.XBusiness, tcid int64, dt *time.Time) error {
	funcname := "AutoAllocatePayorReceipts"
	fmt.Printf("Entered %s\n", funcname)
	var t rlib.Transactant
	if err := rlib.GetTransactant(tcid, &t); err != nil {
		fmt.Printf("error with GetTransactant(%d): %s\n", tcid, err.Error())
		return err
	}
	m := GetAllUnpaidAssessmentsForPayor(t.BID, tcid, dt)
	n := rlib.GetUnallocatedReceiptsByPayor(t.BID, tcid)

	fmt.Printf("Unpaid assessments for payor = %s %s: %d\n", t.FirstName, t.LastName, len(m))
	fmt.Printf("The receipts to be allocated to pay the assessments: %d\n\n", len(n))

	//-----------------------------------------------------------------------
	// Begin paying off bills.  For each assessment, go through the receipts,
	// starting with the oldest, use all its funds, then move on to the next
	//-----------------------------------------------------------------------
	for i := 0; i < len(m); i++ {
		fmt.Printf("ASMID = %d, Amount = %.2f, AR = %d\n", m[i].ASMID, m[i].Amount, m[i].ARID)
		for j := 0; j < len(n); j++ {
			if n[j].FLAGS&3 == 2 { // if there are no funds left in this receipt...
				continue // move on to the next receipt
			}
			// First, determine the amount needed for payment...
			needed := AssessmentUnpaidPortion(&m[i])
			amt := RemainingReceiptFunds(xbiz, &n[j])
			fmt.Printf("Needed for ASMID %d :  %.2f\n", m[i].ASMID, needed)
			fmt.Printf("Funds remaining in receipt %d:  %.2f\n", n[j].RCPTID, amt)
			err := payAssessment(xbiz, &m[i], &n[j], &needed, amt, dt)
			fmt.Printf("\n")
			if err != nil {
				return err
			}
			if needed < ROUNDINGERR { // if we've paid off the assessment...
				break // ... then move on to the next assessment
			}
		}
		fmt.Printf("Completed ASMID = %d\n", m[i].ASMID)
		fmt.Printf("-------------------------------------\n\n")
	}

	return nil
}
