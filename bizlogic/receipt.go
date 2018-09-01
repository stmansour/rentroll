package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"time"
)

// UpdateReceipt accepts an updated rlib.Receipt structure.  It will load the
// existing receipt, compare the fields, and take appropriate action.
//
// In editing a receipt, the following fields require edits to multiple tables:
//   * date
//   * Amount
//   * Account Rule
//
// Editing the date requires updates to
//   * Receipt
//   * ReceiptAllocation
//   * Journal
//   * JournalAllocation
//   * two LedgerEntries
//
// INPUTS
//   rnew = updated receipt
//     dt = date to use for entry reversal if necessary. This date is not used
//          if rnew can be updated without having to reverse it's prior entry.
//
// RETURNS
//    err = any error that was encountered.
//-------------------------------------------------------------------------------
func UpdateReceipt(ctx context.Context, rnew *rlib.Receipt, dt *time.Time) error {
	// funcname := "bizlogic.UpdateReceipt"

	if rnew.FLAGS&0x4 != 0 {
		return fmt.Errorf("This item cannot be edited, it has been reversed") // it's already reversed
	}

	errlist := ValidateReceipt(ctx, rnew)
	if errlist != nil {
		return BizErrorListToError(errlist)
	}

	//-------------------------------
	// Load existing receipt...
	//-------------------------------
	rold, err := rlib.GetReceipt(ctx, rnew.RCPTID)
	if err != nil {
		return err
	}
	if rold.RCPTID == 0 { // if resource not found then raise the error
		return fmt.Errorf("Receipt %d not found", rnew.RCPTID)
	}

	//---------------------------------------------------------------------------------
	// we need to reverse the old receipt if any of the following fields have changed:
	//    * Dt
	//    * Amount
	//    * AccountRule
	//    * RAID
	//---------------------------------------------------------------------------------
	reverse := (!rold.Dt.Equal(rnew.Dt)) || rold.Amount != rnew.Amount || rold.ARID != rnew.ARID || rold.RAID != rnew.RAID
	if reverse {
		err := ReverseReceipt(ctx, &rold, dt) // reverse the receipt itself
		if err != nil {
			return err
		}
		err = InsertReceipt(ctx, rnew) // Insert the new receipt...
		if err != nil {
			return err
		}
		if rnew.DID > 0 { // update DepositPart if necessary
			var dp = rlib.DepositPart{
				DID:    rnew.DID,
				BID:    rnew.BID,
				RCPTID: rnew.RCPTID,
			}
			_, err = rlib.InsertDepositPart(ctx, &dp)
			if err != nil {
				return err
			}
		}
		// the deposit total may have changed...
		if rold.Amount != rnew.Amount && rnew.DID > 0 {
			dep, err := rlib.GetDeposit(ctx, rnew.DID)
			if err != nil {
				return err
			}
			dep.Amount = dep.Amount - rold.Amount + rnew.Amount
			return rlib.UpdateDeposit(ctx, &dep)
		}
		return nil
	}

	return rlib.UpdateReceipt(ctx, rnew) // reversal not needed, just update the receipt
}

// ReverseReceipt reverses the supplied receipt. It links the
// reversal back to the supplied receipt
// RETURNS
//    any error that occurred, or nil if no error
//-------------------------------------------------------------------------------
func ReverseReceipt(ctx context.Context, r *rlib.Receipt, dt *time.Time) error {
	var err error
	rlib.Console("Entered ReverseReceipt\n")
	if r.FLAGS&0x04 != 0 {
		return nil // it's already reversed
	}
	rlib.Console("ReverseReceipt processing RCPTID = %d\n", r.RCPTID)

	//----------------------------------------------------------------------
	// The old receipt may have been allocated. Load its ReceiptAllocations
	// and reverse any allocation that was applied towards an Assessment
	//----------------------------------------------------------------------
	if len(r.RA) == 0 { // if RA slice is empty, it could be because they were not loaded
		err = rlib.GetReceiptAllocations(ctx, r.RCPTID, r) // try to load them just to make sure
		if err != nil {
			return err
		}
	}

	//------------------------------------------------------
	// Build the new receipt
	//------------------------------------------------------
	rr := *r
	rr.RCPTID = int64(0)
	rr.Amount = -rr.Amount
	rr.Comment = fmt.Sprintf("Reversal of receipt %s", r.IDtoString())
	rr.PRCPTID = r.RCPTID     // link to parent
	rr.FLAGS |= rlib.RCPTvoid // mark that it is voided
	rr.RA = []rlib.ReceiptAllocation{}
	if err = insertReceiptInternal(ctx, &rr, r, dt); err != nil {
		return err
	}
	//-----------------------------------------------------------
	// Newly created Receipts will have allocations with 0 FLAGS
	// We need to have a flag that shows the new allocation is
	// part of a reversal...
	//-----------------------------------------------------------
	for i := 0; i < len(rr.RA); i++ {
		rlib.Console("Before FLAGS update, rr.RA[i].FLAGS = %d\n", rr.RA[i].FLAGS)
		rr.RA[i].FLAGS |= rlib.RCPTvoid
		if err = rlib.UpdateReceiptAllocation(ctx, &rr.RA[i]); err != nil {
			return err
		}
		rlib.Console("Reverse Receipt loc 1: updated ReceiptAllocation: RCPAID = %d, FLAGS = %d\n", rr.RA[i].RCPAID, rr.RA[i].FLAGS)
	}

	//---------------------------------------------------------------------------
	// If the receipt was part of a deposit, add the reversal to the deposit...
	//---------------------------------------------------------------------------
	if r.DID > 0 {
		var dp = rlib.DepositPart{
			DID:    rr.DID,
			BID:    rr.BID,
			RCPTID: rr.RCPTID,
		}
		_, err := rlib.InsertDepositPart(ctx, &dp)
		if err != nil {
			return err
		}
	}

	//--------------------------------------------------
	// Reverse any allocations that have been made...
	//--------------------------------------------------
	for i := 0; i < len(r.RA); i++ {
		rlib.Console("Second Try: Before FLAGS update, rr.RA[i].FLAGS = %d\n", r.RA[i].FLAGS)
		r.RA[i].FLAGS |= rlib.RCPTvoid
		if err := rlib.UpdateReceiptAllocation(ctx, &r.RA[i]); err != nil {
			return err
		}
		rlib.Console("Reverse Receipt loc 2: updated ReceiptAllocation: RCPAID = %d, FLAGS = %d\n", r.RA[i].RCPAID, r.RA[i].FLAGS)
		if r.RA[i].ASMID == 0 {
			//--------------------------------------------------
			// This allocation has no associated Assessment.
			// Could be vending income, or similar.
			// It's already been marked as void, just continue
			//--------------------------------------------------
			continue
		}
		ra := r.RA[i]          // copy it and make the reversal changes
		ra.RCPTID = rr.RCPTID  // the reversal receipt id
		ra.Amount = -ra.Amount // reverse the allocation
		ra.Dt = *dt            // date of reversal
		ra.RAID = r.RA[i].RAID // RAID toward which this receipt was applied

		// build a new AcctRule
		var xbiz1 rlib.XBusiness // not actually used, but needed for the call to ParseAcctRule
		// TODO(Steve): should we ignore error?
		n, _ := rlib.ParseAcctRule(ctx, &xbiz1, 0, dt, dt, ra.AcctRule, 0, 1.0)
		acctrule := ""
		for k := 0; k < len(n); k++ {
			acctrule += fmt.Sprintf("ASM(%d) %s %s %.2f", ra.ASMID, n[k].Action, n[k].Account, ra.Amount)
			if k+1 < len(n) {
				acctrule += ","
			}
		}
		ra.AcctRule = acctrule
		rlib.Console("Reverse Receipt loc 3: updated ReceiptAllocation: RCPAID = %d, FLAGS = %d\n", ra.RCPAID, ra.FLAGS)
		_, err := rlib.InsertReceiptAllocation(ctx, &ra)
		if err != nil {
			return err
		}
		rr.RA = append(rr.RA, ra)
	}

	//------------------------------------------------------
	// mark the old receipt as voided
	//------------------------------------------------------
	r.FLAGS |= rlib.RCPTvoid // mark that it is voided
	if len(r.Comment) > 0 {
		r.Comment += ", "
	}
	r.Comment += fmt.Sprintf("Reversed by receipt %s", rr.IDtoString())
	err = rlib.UpdateReceipt(ctx, r)
	if err != nil {
		return err
	}

	//------------------------------------------------------
	// reverse any payments allocated from this receipt...
	//------------------------------------------------------
	if (r.FLAGS & 0x3) > 0 {
		ReverseAllocation(ctx, r, rr.RCPTID, dt)
	}

	return err
}

// ReverseAllocation reverses any payments funded by this receipt.
// INPUTS
//    r = receipt to be voided
//    revRCPTID = RID of the reversal receipt
//
// RETURNS
//    any error that occurred, or nil if no error
//-------------------------------------------------------------------------------
func ReverseAllocation(ctx context.Context, r *rlib.Receipt, revRCPTID int64, dt *time.Time) error {
	funcname := "bizlogic.ReverseAllocation"
	var err error

	//------------------------------------------------------
	// Spin through all journal entries that reference
	// r.RCPTID. If it represents a payment allocation then
	// reverse it.
	//------------------------------------------------------
	m, err := rlib.GetJournalsByReceiptID(ctx, r.RCPTID)
	if err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {
		//-----------------------------------------------------------
		// Reverse all the JournalAllocation entries in which
		// the funds of r have been applied to a receipt.
		//-----------------------------------------------------------
		err = rlib.GetJournalAllocations(ctx, &m[i]) // load all its allocations
		if err != nil {
			continue
		}
		if len(m[i].JA) == 0 {
			continue
		}
		if m[i].JA[0].RCPTID == 0 { // may be an entry for dep to bank and put into unapplied funds
			continue
		}

		for j := 0; j < len(m[i].JA); j++ {
			//--------------------------------------------
			// First, reverse the journal entry
			//--------------------------------------------
			var jnl = rlib.Journal{
				BID:    r.BID,
				Amount: -m[i].Amount, // reverse the amount
				ID:     revRCPTID,    // this is the rcptid of the reversal receipt
				Dt:     *dt,          // reversal date
				Type:   rlib.JNLTYPERCPT,
			}
			_, err = rlib.InsertJournal(ctx, &jnl)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return err
			}

			//-------------------------------------------------------------------------
			// Next, add the JournalAllocation reversal
			//-------------------------------------------------------------------------
			var xbiz1 rlib.XBusiness // not actually used
			// TODO(Steve): should we ignore error?
			n, _ := rlib.ParseAcctRule(ctx, &xbiz1, 0, dt, dt, m[i].JA[j].AcctRule, 0, 1.0)
			acctrule := ""
			for k := 0; k < len(n); k++ {
				acctrule += fmt.Sprintf("ASM(%d) %s %s %.2f", m[i].JA[j].ASMID, n[k].Action, n[k].Account, jnl.Amount)
				if k+1 < len(n) {
					acctrule += ","
				}
			}
			var ja = rlib.JournalAllocation{
				JID:      jnl.JID,
				AcctRule: acctrule,
				Amount:   jnl.Amount,
				BID:      jnl.BID,
				RAID:     m[i].JA[j].RAID,
				RID:      m[i].JA[j].RID,
				ASMID:    m[i].JA[j].ASMID,
				TCID:     r.TCID,
				RCPTID:   revRCPTID,
			}
			_, err = rlib.InsertJournalAllocationEntry(ctx, &ja)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return err
			}
			jnl.JA = append(jnl.JA, ja)

			//-------------------------------------------------------------------------
			// Next, reverse the ledger entries...
			//-------------------------------------------------------------------------
			le, err := rlib.GetLedgerEntriesByJAID(ctx, r.BID, m[i].JA[j].JAID)
			if err != nil {
				return err
			}

			for k := 0; k < len(le); k++ {
				nle := le[k]
				nle.JAID = ja.JAID       // our newly created reversing Journal Allocation
				nle.JID = ja.JID         // which is tied to the reversing Journal entry
				nle.Amount = -nle.Amount // this reverses the amount
				_, err = rlib.InsertLedgerEntry(ctx, &nle)
				if err != nil {
					rlib.LogAndPrintError(funcname, err)
					return err
				}
			}

			//-------------------------------------------------------------------------
			// Finally, update the assessment that was allocated payment from this receipt...
			//-------------------------------------------------------------------------
			a, err := rlib.GetAssessment(ctx, m[i].JA[j].ASMID)
			if err != nil {
				return err
			}
			unpaid := AssessmentUnpaidPortion(ctx, &a) // how much of this assessment is still unpaid?
			paid := a.Amount - unpaid                  // how much remains to be paid
			remaining := paid - m[i].Amount            // how much remains after removing this allocation

			newflags := uint64(0) // assume nothing has been paid on the assessment after this reversal
			if remaining > 0 {    // if any portion has still been paid...
				newflags = uint64(1) // ... then mark as partially paid
			}
			b := uint64(0x3)    // the bits of interest
			b = ^b              // flip the bits
			a.FLAGS &= b        // clear those bits in FLAGS
			a.FLAGS |= newflags // set new status
			err = rlib.UpdateAssessment(ctx, &a)
			if err != nil {
				return err
			}
		}
	}

	//------------------------------------------------------
	// void the old receipt
	//------------------------------------------------------
	b := uint64(0x3)
	b = ^b
	r.FLAGS &= b   // remove any payment related flags that might confuse anyone
	r.FLAGS |= 0x4 // set bit 2 to indicate that it has been voided
	return rlib.UpdateReceipt(ctx, r)
}

// InsertReceipt adds a new receipt and updates the journal and ledgers
//-------------------------------------------------------------------------------
func InsertReceipt(ctx context.Context, a *rlib.Receipt) error {
	return insertReceiptInternal(ctx, a, nil, nil)
}

// insertReceiptInternal adds a new receipt and updates the journal and ledgers,
// It also handles reversal receipts that happen to have SubARs that created
// associated Assessments. It will find and reverse those assessments as well.
//
// INPUT
//  a      - the receipt to insert
//  origin - if being called in a reversal, this is the original receipt (the
//           one being reversed)
//  dt     - if a reversal is made, use this date for the reversal
//-------------------------------------------------------------------------------
func insertReceiptInternal(ctx context.Context, a, origin *rlib.Receipt, dt *time.Time) error {

	funcname := "bizlogic.InsertReceipt"
	var (
		err  error
		xbiz rlib.XBusiness
	)

	//------------------------------------------------
	// Set up context around the Account Rule
	//------------------------------------------------

	err = rlib.InitBizInternals(a.BID, &xbiz)
	if err != nil {
		return err
	}

	ar := rlib.RRdb.BizTypes[a.BID].AR[a.ARID]               // get the AR for this receipt...
	ard := rlib.RRdb.BizTypes[a.BID].GLAccounts[ar.DebitLID] // get GL Account Info for debits and credits
	arc := rlib.RRdb.BizTypes[a.BID].GLAccounts[ar.CreditLID]

	//----------------------------------------------------------
	// FLAGS bit 0,1 (i.e., binary 0011) indicates how the
	// receipt is allocated: 0 => unallocated, 1 => partially,
	// 2 => fully.
	//----------------------------------------------------------
	a.FLAGS &= ^uint64(0x3) // assume account rule does NOT allocate fully on insert
	if ar.FLAGS&0x1 != 0 {  // if rule indicate fully-allocate
		a.FLAGS |= 0x2 // mark the state as fully allocated
	}

	errlist := ValidateReceipt(ctx, a)
	if errlist != nil {
		return BizErrorListToError(errlist)
	}

	_, err = rlib.InsertReceipt(ctx, a)
	if err != nil {
		return err
	}

	//------------------------------------------------------------
	// Check to see if the Account Rule has Sub-Account Rules.
	// If this happens to be a Reversal, then do not do anything
	//
	//------------------------------------------------------------
	if ar.FLAGS&(1<<3) > 0 {
		ar.SubARs, err = rlib.GetSubARs(ctx, ar.ARID)
		if err != nil {
			return err
		}

		noClose := rlib.ClosePeriod{
			Dt:           rlib.TIME0,
			OpenPeriodDt: rlib.TIME0,
		}

		for i := 0; i < len(ar.SubARs); i++ {
			sub := rlib.RRdb.BizTypes[a.BID].AR[ar.SubARs[i].SubARID]
			switch sub.ARType {
			case rlib.ARASSESSMENT:
			case rlib.ARRECEIPT:
			case rlib.AREXPENSE:
			case rlib.ARSUBASSESSMENT:
				if a.FLAGS&(1<<2) != 0 { // this receipt is a reversal. We need to reverse the associated assessment
					var agasmt rlib.Assessment
					q := fmt.Sprintf("SELECT %s FROM Assessments WHERE AGRCPTID=%d", rlib.RRdb.DBFields["Assessments"], origin.RCPTID)
					row := rlib.RRdb.Dbrr.QueryRow(q)
					err = rlib.ReadAssessment(row, &agasmt)
					if err != nil {
						return err
					}

					if agasmt.ASMID > 0 {
						be := ReverseAssessment(ctx, &agasmt, 0, dt, &noClose)
						if len(be) > 0 {
							return BizErrorListToError(be)
						}
					}
				} else {
					asmt, be := CreateSubAssessment(ctx, &sub, a) // create a new assessment that this receipt pays for
					if len(be) > 0 {
						return BizErrorListToError(be)
					}
					var ra rlib.ReceiptAllocation // creat the apply portion of this receipt
					ra.RCPTID = a.RCPTID
					ra.Amount = a.Amount
					// for this case, we follow the Assessment's debit and credit
					ra.AcctRule = fmt.Sprintf("ASM(%d) d %s _, c %s _", asmt.ASMID,
						rlib.RRdb.BizTypes[a.BID].GLAccounts[sub.DebitLID].GLNumber,
						rlib.RRdb.BizTypes[a.BID].GLAccounts[sub.CreditLID].GLNumber)
					ra.ASMID = asmt.ASMID
					ra.BID = a.BID
					ra.Dt = a.Dt
					ra.RAID = a.RAID
					_, err = rlib.InsertReceiptAllocation(ctx, &ra)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	//--------------------------------------------------------------------------
	// create the receive portion of the receipt allocation receipt allocation
	//--------------------------------------------------------------------------
	var ra rlib.ReceiptAllocation
	ra.RCPTID = a.RCPTID
	ra.Amount = a.Amount
	ra.AcctRule = fmt.Sprintf("d %s _, c %s _", ard.GLNumber, arc.GLNumber)
	ra.BID = a.BID
	ra.Dt = a.Dt
	ra.RAID = a.RAID
	_, err = rlib.InsertReceiptAllocation(ctx, &ra)
	if err != nil {
		return err
	}
	a.RA = append(a.RA, ra)

	//------------------------------------------------
	// Add it to the Journal
	//------------------------------------------------
	d1 := time.Date(a.Dt.Year(), a.Dt.Month(), 1, 0, 0, 0, 0, rlib.RRdb.Zone)
	mon, year := rlib.IncMonths(a.Dt.Month(), int64(a.Dt.Year()))
	d2 := time.Date(int(year), mon, 1, 0, 0, 0, 0, rlib.RRdb.Zone)
	jnl, err := rlib.ProcessNewReceipt(ctx, &xbiz, &d1, &d2, a)
	if err != nil {
		e := fmt.Errorf("%s:  Error in rlib.ProcessNewReceipt: %s", funcname, err.Error())
		rlib.Ulog("%s", e.Error())
		return e
	}

	//------------------------------------------------
	// Add it to the Ledgers
	//------------------------------------------------
	err = rlib.GetJournalAllocations(ctx, &jnl)
	if err != nil {
		return err
	}

	rlib.InitLedgerCache()

	// TODO(Steve): ignore int here?
	_, err = rlib.GenerateLedgerEntriesFromJournal(ctx, &xbiz, &jnl, &d1, &d2)
	if err != nil {
		return err
	}

	return err
}

// CreateSubAssessment creates an assessment that is being associated with a
// particular receipt. Example usages are for use cases such as floating
// deposits or application fees.
//
// INPUT:
//  sub  - the sub-assessment account rule
//  a    - the receipt that will be associated with the assessment.
//
// RETURNS:
//
//-----------------------------------------------------------------------------
func CreateSubAssessment(ctx context.Context, sub *rlib.AR, a *rlib.Receipt) (rlib.Assessment, []BizError) {
	var b rlib.Assessment
	// for any value not set below, the default value is correct
	b.BID = a.BID
	b.RAID = a.RAID
	b.ARID = sub.ARID
	b.AGRCPTID = a.RCPTID // this is the receipt that caused this assessment to be generated
	b.Amount = a.Amount
	b.Start = a.Dt
	b.Stop = a.Dt
	b.Comment = "Auto-generated by Account Rule (" + sub.Name + ")"
	b.FLAGS = a.FLAGS

	noClose := rlib.ClosePeriod{
		Dt:           rlib.TIME0,
		OpenPeriodDt: rlib.TIME0,
	}

	be := InsertAssessment(ctx, &b, 0, &noClose)

	//--------------------------------------------------------------------
	// The JournalAllocation record associated with this assessment must
	// now be updated with a.RCPTID to bind the two together
	//--------------------------------------------------------------------
	m, err := rlib.GetJournalAllocationByASMID(ctx, b.ASMID)
	if err != nil {
		return b, AddErrToBizErrlist(err, be)
	}

	for i := 0; i < len(m); i++ {
		m[i].RCPTID = a.RCPTID
		err := rlib.UpdateJournalAllocation(ctx, &m[i])
		if err != nil {
			be = AddErrToBizErrlist(err, be)
		}
	}
	return b, be
}

// ValidateReceipt checks to see whether the assessment violates any
// business logic or any fields are missing or bad.
//
// INPUTS
//    r = the receipt to validate
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ValidateReceipt(ctx context.Context, r *rlib.Receipt) []BizError {
	var e []BizError
	fields := []string{}
	if r.TCID == 0 {
		fields = append(fields, "Payor")
	}
	if len(fields) > 0 {
		msg := BizErrors[InvalidField].Message
		for i := 0; i < len(fields); i++ {
			msg += fmt.Sprintf("\n%s", fields[i])
		}
		e = append(e, BizError{Errno: InvalidField, Message: msg})
	}
	if len(e) > 0 {
		return e
	}
	return nil
}
