package bizlogic

import (
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
func UpdateReceipt(rnew *rlib.Receipt, dt *time.Time) error {
	if rnew.FLAGS&0x4 != 0 {
		return fmt.Errorf("This item cannot be edited, it has been reversed") // it's already reversed
	}

	errlist := ValidateReceipt(rnew)
	if errlist != nil {
		return BizErrorListToError(errlist)
	}

	//-------------------------------
	// Load existing receipt...
	//-------------------------------
	rold := rlib.GetReceipt(rnew.RCPTID)
	if rold.RCPTID == 0 {
		return fmt.Errorf("Receipt %d not found", rnew.RCPTID)
	}

	//---------------------------------------------------------------------------------
	// we need to reverse the old receipt if any of the following fields have changed:
	//    * Dt
	//    * Amount
	//    * AccountRule
	//---------------------------------------------------------------------------------
	reverse := (!rold.Dt.Equal(rnew.Dt)) || rold.Amount != rnew.Amount || rold.ARID != rnew.ARID
	if reverse {
		err := ReverseReceipt(&rold, dt) // reverse the receipt itself
		if err != nil {
			return err
		}
		return InsertReceipt(rnew) // Finally, insert the new receipt...
	}

	return rlib.UpdateReceipt(rnew) // reversal not needed, just update the receipt
}

// ReverseReceipt reverses the supplied receipt. It links the
// reversal back to the supplied receipt
// RETURNS
//    any error that occurred, or nil if no error
//-------------------------------------------------------------------------------
func ReverseReceipt(r *rlib.Receipt, dt *time.Time) error {
	if r.FLAGS&0x04 != 0 {
		return nil // it's already reversed
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
	if err := InsertReceipt(&rr); err != nil {
		return err
	}

	//----------------------------------------------------------------------
	// The old receipt may have been allocated. Load its ReceiptAllocations
	// and reverse any allocation that was applied towards an Assessment
	//----------------------------------------------------------------------
	if len(r.RA) == 0 { // if RA slice is empty, it could be because they were not loaded
		rlib.GetReceiptAllocations(r.RCPTID, r) // try to load them just to make sure
	}

	for i := 0; i < len(r.RA); i++ {
		if r.RA[i].ASMID == 0 {
			continue // skip entries that are not allocations to an assessment
		}
		ra := r.RA[i]          // copy it and make the reversal changes
		ra.RCPTID = rr.RCPTID  // the reversal receipt id
		ra.Amount = -ra.Amount // reverse the allocation
		ra.Dt = *dt            // date of reversal

		// build a new AcctRule
		var xbiz1 rlib.XBusiness // not actually used, but needed for the call to ParseAcctRule
		n := rlib.ParseAcctRule(&xbiz1, 0, dt, dt, ra.AcctRule, 0, 1.0)
		acctrule := ""
		for k := 0; k < len(n); k++ {
			acctrule += fmt.Sprintf("ASM(%d) %s %s %.2f", ra.ASMID, n[k].Action, n[k].Account, ra.Amount)
			if k+1 < len(n) {
				acctrule += ","
			}
		}
		ra.AcctRule = acctrule
		_, err := rlib.InsertReceiptAllocation(&ra)
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
	err := rlib.UpdateReceipt(r)
	if err != nil {
		return err
	}

	//------------------------------------------------------
	// reverse any payments allocated from this receipt...
	//------------------------------------------------------
	if (r.FLAGS & 0x3) > 0 {
		ReverseAllocation(r, rr.RCPTID, dt)
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
func ReverseAllocation(r *rlib.Receipt, revRCPTID int64, dt *time.Time) error {
	funcname := "bizlogic.ReverseAllocation"
	var err error

	//------------------------------------------------------
	// Spin through all journal entries that reference
	// r.RCPTID. If it represents a payment allocation then
	// reverse it.
	//------------------------------------------------------
	m := rlib.GetJournalsByReceiptID(r.RCPTID)
	for i := 0; i < len(m); i++ {
		//-----------------------------------------------------------
		// Reverse all the JournalAllocation entries in which
		// the funds of r have been applied to a receipt.
		//-----------------------------------------------------------
		rlib.GetJournalAllocations(&m[i]) // load all its allocations
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
			_, err = rlib.InsertJournal(&jnl)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return err
			}

			//-------------------------------------------------------------------------
			// Next, add the JournalAllocation reversal
			//-------------------------------------------------------------------------
			var xbiz1 rlib.XBusiness // not actually used
			n := rlib.ParseAcctRule(&xbiz1, 0, dt, dt, m[i].JA[j].AcctRule, 0, 1.0)
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
			rlib.InsertJournalAllocationEntry(&ja)
			jnl.JA = append(jnl.JA, ja)

			//-------------------------------------------------------------------------
			// Next, reverse the ledger entries...
			//-------------------------------------------------------------------------
			le := rlib.GetLedgerEntriesByJAID(r.BID, m[i].JA[j].JAID)
			for k := 0; k < len(le); k++ {
				nle := le[k]
				nle.JAID = ja.JAID       // our newly created reversing Journal Allocation
				nle.JID = ja.JID         // which is tied to the reversing Journal entry
				nle.Amount = -nle.Amount // this reverses the amount
				_, err = rlib.InsertLedgerEntry(&nle)
				if err != nil {
					rlib.LogAndPrintError(funcname, err)
					return err
				}
			}

			//-------------------------------------------------------------------------
			// Finally, update the assessment that was allocated payment from this receipt...
			//-------------------------------------------------------------------------
			a, err := rlib.GetAssessment(m[i].JA[j].ASMID)
			if err != nil {
				return err
			}
			unpaid := AssessmentUnpaidPortion(&a) // how much of this assessment is still unpaid?
			paid := a.Amount - unpaid             // how much remains to be paid
			remaining := paid - m[i].Amount       // how much remains after removing this allocation

			newflags := uint64(0) // assume nothing has been paid on the assessment after this reversal
			if remaining > 0 {    // if any portion has still been paid...
				newflags = uint64(1) // ... then mark as partially paid
			}
			b := uint64(0x3)    // the bits of interest
			b = ^b              // flip the bits
			a.FLAGS &= b        // clear those bits in FLAGS
			a.FLAGS |= newflags // set new status
			err = rlib.UpdateAssessment(&a)
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
	return rlib.UpdateReceipt(r)
}

// InsertReceipt adds a new receipt and updates the journal and ledgers
//-------------------------------------------------------------------------------
func InsertReceipt(a *rlib.Receipt) error {
	funcname := "bizlogic.InsertReceipt"
	errlist := ValidateReceipt(a)
	if errlist != nil {
		return BizErrorListToError(errlist)
	}
	_, err := rlib.InsertReceipt(a)
	if err != nil {
		return err
	}

	var xbiz rlib.XBusiness
	rlib.InitBizInternals(a.BID, &xbiz)
	ar := rlib.RRdb.BizTypes[a.BID].AR[a.ARID]               // get the AR for this receipt...
	ard := rlib.RRdb.BizTypes[a.BID].GLAccounts[ar.DebitLID] // get GL Account Info for debits and credits
	arc := rlib.RRdb.BizTypes[a.BID].GLAccounts[ar.CreditLID]

	//------------------------------------------------
	// create the receipt allocation
	//------------------------------------------------
	var ra rlib.ReceiptAllocation
	ra.RCPTID = a.RCPTID
	ra.Amount = a.Amount
	ra.AcctRule = fmt.Sprintf("d %s _, c %s _", ard.GLNumber, arc.GLNumber)
	ra.BID = a.BID
	ra.Dt = a.Dt
	_, err = rlib.InsertReceiptAllocation(&ra)
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
	jnl, err := rlib.ProcessNewReceipt(&xbiz, &d1, &d2, a)
	if err != nil {
		e := fmt.Errorf("%s:  Error in rlib.ProcessNewReceipt: %s", funcname, err.Error())
		rlib.Ulog("%s", e.Error())
		return e
	}

	//------------------------------------------------
	// Add it to the Ledgers
	//------------------------------------------------
	rlib.GetJournalAllocations(&jnl)
	rlib.InitLedgerCache()
	rlib.GenerateLedgerEntriesFromJournal(&xbiz, &jnl, &d1, &d2)
	return nil
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
func ValidateReceipt(r *rlib.Receipt) []BizError {
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
