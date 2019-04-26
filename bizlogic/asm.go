package bizlogic

import (
	"context"
	"database/sql"
	"fmt"
	"rentroll/rlib"
	"time"
)

// UpdateAssessment updates the supplied assessment, reversing existing assessments
// if necessary
//
// Aug 27, 2018 - if a recurring definition is updated and its stop date is set
// back in time, we now reverse all instance of it on or after the new stop date.
//
// INPUTS
//    ctx  = database context
//    anew = the assessment to update
//    mode = how to handle recurring assessments:
//           0: just reverse this instance
//           1: reverse this and future instances
//           2: reverse all instances
//    dt   = date of modification
//    lc   = date of last close
//    exp  = if it is a recurring assessment and the start date is in the past,
//           should past entries be created?  1 = yes
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func UpdateAssessment(ctx context.Context, anew *rlib.Assessment, mode int, dt *time.Time, lc *rlib.ClosePeriod, exp int) []BizError {
	var err error
	var errlist []BizError

	// rlib.Console("Entered bizlogic.UpdateAssessment:  anew.ASMID = %d, mode = %d, dt = %s, exp = %d\n", anew.ASMID, mode, dt.Format(rlib.RRDATEREPORTFMT), exp)
	// rlib.Console("anew.FLAGS = %X\n", anew.FLAGS)

	errlist = ValidateAssessment(ctx, anew) // make sure it passes simple validation first
	if len(errlist) > 0 {
		return errlist
	}

	//------------------------------------------------
	// make sure we're not editing a reversed entry
	//------------------------------------------------
	if anew.FLAGS&0x4 != 0 {
		errlist = append(errlist, BizErrors[EditReversal])
		return errlist
	}
	//-------------------------------
	// Load existing assessment...
	//-------------------------------
	aold, err := rlib.GetAssessment(ctx, anew.ASMID)
	if err != nil {
		return bizErrSys(&err)
	}

	if aold.ASMID == 0 {
		err = fmt.Errorf("Assessment %d not found", anew.ASMID)
		return bizErrSys(&err)
	}

	//---------------------------------------------------------------------------------
	// we need to reverse the old assessment if any of the following fields have changed:
	//   ARID
	//   RAID
	//   RID
	//   Amount
	//   Recur Cycle
	//   Proration Cycle
	//   Start Date
	//   Stop Date if it moves backwards in time
	//---------------------------------------------------------------------------------
	reverse := aold.ARID != anew.ARID ||
		aold.RAID != anew.RAID ||
		aold.RID != anew.RID ||
		aold.Amount != anew.Amount ||
		aold.RentCycle != anew.RentCycle ||
		aold.ProrationCycle != anew.ProrationCycle ||
		(!aold.Start.Equal(anew.Start)) ||
		(!aold.Stop.Equal(anew.Stop))
	if reverse {
		//---------------------------------------------------------------------------
		// Reverse any instances which will be out of the new updated date range
		// This must be done before calling ReverseAssessment, which modifies
		// aold.  ReverseAssessmentsAfterStop needs to work with aold before
		// any changes are made to it.
		//---------------------------------------------------------------------------
		if anew.Stop.Before(aold.Stop) { // is new stop date earlier in time than the old one?
			errlist = ReverseAssessmentsAfterStop(ctx, &aold, &anew.Stop, dt, lc)
			if len(errlist) > 0 {
				return errlist
			}
		}

		//---------------------------------------------------------------------------
		// Now reverse aold...
		//---------------------------------------------------------------------------
		errlist = ReverseAssessment(ctx, &aold, mode, dt, lc) // reverse the assessment itself
		if errlist != nil {
			return errlist
		}

		//---------------------------------------------------------------------------
		// This is going to be a new assessment that replaces an assessment which
		// has just been reversed. So it is NOT reversed, and it is NOT paid in
		// any part. So, we need to reset the flags.  Bits 0:1 define payment and
		// bit 2 defines reversal.  So clear bits 0:2...
		//---------------------------------------------------------------------------
		anew.FLAGS &= ^uint64(7) // clears the first 3 bits
		// rlib.Console("ANEW.FLAGS = %d\n", anew.FLAGS)

		errlist = InsertAssessment(ctx, anew, exp, lc) // Finally, insert the new assessment...
		if errlist != nil {
			return errlist
		}
		return nil
	}

	err = rlib.UpdateAssessment(ctx, anew) // reversal not needed, just update the assessment
	if err != nil {
		return bizErrSys(&err)
	}
	return nil
}

// ReverseAssessment reverses an existing assessment
//
// INPUTS
//    aold = the assessment to reverse
//    mode = how to handle recurring assessments:
//           0: just reverse this instance
//           1: reverse this and future instances
//           2: reverse all instances
//    dt   = time to mark when the reversal was made
//    lc   = last closing date, date when the last period was closed.
//           Reversing entries will be made on the later date of dt or lc.
//           The routine honors the last close date.
//
// RETURNS
//    a slice of BizErrors
//-----------------------------------------------------------------------------
func ReverseAssessment(ctx context.Context, aold *rlib.Assessment, mode int, dt *time.Time, lc *rlib.ClosePeriod) []BizError {
	funcname := "bizlogic.ReverseAssessment"
	var errlist []BizError
	rlib.Console("#####>>>>>>>>> Entered ReverseAssessment. ASMID = %d, mode = %d,  dt = %s\n", aold.ASMID, mode, dt.Format(rlib.RRDATEFMTSQL))
	// debug.PrintStack()

	if aold.PASMID == 0 && aold.RentCycle > 0 {
		mode = 2 // force behavior on the epoch
	}
	rlib.Console("ReverseAssessment: processing forward with mode = %d,  dt = %s\n", mode, dt.Format(rlib.RRDATEFMTSQL))
	switch mode {
	case 0:
		errlist = ReverseAssessmentInstance(ctx, aold, dt, lc)
	case 1:
		errlist = ReverseAssessmentsGoingForward(ctx, aold, &aold.Start, dt, lc)
	case 2:
		var epoch, inst rlib.Assessment
		var err error

		//---------------------------------------------------------
		// set the epoch
		//---------------------------------------------------------
		if aold.PASMID != 0 {
			epoch, err = rlib.GetAssessment(ctx, aold.PASMID)
			if err != nil {
				rlib.Console("EXITING ReverseAssessment.  PT 1\n")
				return bizErrSys(&err)
			}
		} else {
			epoch = *aold
		}

		//---------------------------------------------------------
		// If it is not recurring then reverse it and we're done
		//---------------------------------------------------------
		if epoch.RentCycle == rlib.RECURNONE {
			rlib.Console("EXITING ReverseAssessment.  PT 2\n")
			return ReverseAssessmentInstance(ctx, &epoch, dt, lc)
		}

		//---------------------------------------------------------
		// Get the first instance and modify forward...
		//---------------------------------------------------------
		inst, err = rlib.GetAssessmentFirstInstance(ctx, epoch.ASMID)
		if err != nil {
			rlib.Console("EXITING ReverseAssessment.  PT 3\n")
			return bizErrSys(&err)
		}
		if inst.ASMID > 0 { // only need to do the following if any instances have been created
			rlib.Console("*** -- GetAssesmentFirstInstance returns: ASMID = %d\n", inst.ASMID)
			rlib.Console("*** -- Calling ReverseAssessmentGoingFwd(ctx, asmid=%d)\n", inst.ASMID)
			errlist = ReverseAssessmentsGoingForward(ctx, &inst, &inst.Start, dt, lc) // reverse from start of recurring instances forward
			if len(errlist) > 0 {
				rlib.Console("EXITING ReverseAssessment.  PT 4\n")
				return errlist
			}
		}
		epoch.FLAGS |= 0x4 // mark that this is void
		err = rlib.UpdateAssessment(ctx, &epoch)
		if err != nil {
			rlib.Console("EXITING ReverseAssessment.  PT 5\n")
			return bizErrSys(&err)
		}

	default:
		err := fmt.Errorf("%s:  unsupported mode: %d", funcname, mode)
		rlib.LogAndPrintError(funcname, err)
		rlib.Console("EXITING ReverseAssessment.  PT 6\n")
		return bizErrSys(&err)
	}
	rlib.Console("EXITING ReverseAssessment\n")
	return errlist
}

// ReverseAssessmentsGoingForward reverses an existing assessment.  If aold
// points to an instance, the recurring ASMID is taken to be the PASMID. If
// aold points to a recurring sequence, then the ASMID is used.
//
// INPUTS
//    ctx     = context needed for db transactions
//    aold    = the first in a series of assessments to reverse
//    dtStart = reverse instances from this point in time forward
//    dt      = time to mark when the reversal was made
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessmentsGoingForward(ctx context.Context, aold *rlib.Assessment, dtStart, dt *time.Time, lc *rlib.ClosePeriod) []BizError {
	rlib.Console("ENTERED: ReverseAssessmentsGoingForward\n")
	var errlist []BizError
	var RecurringAsmDefID = aold.PASMID // assume that this points to the recurring definition
	if int64(0) == RecurringAsmDefID {  // if it does not...
		RecurringAsmDefID = aold.ASMID // ...then change to the recurring dev
	}

	if rlib.TIME0.After(*dtStart) {
		rlib.Console("dtStart = %s -- prior to rlib.TIME0\n", dtStart.Format(rlib.RRDATEREPORTFMT))
		// debug.PrintStack()
	}

	d2 := rlib.ENDOFTIME
	// rlib.Console("aold.PASMID = %d, dtStart = %s, dt = %s\n", aold.PASMID, dtStart.Format(rlib.RRDATEREPORTFMT), dt.Format(rlib.RRDATEREPORTFMT))

	m, err := rlib.GetAssessmentInstancesByParent(ctx, RecurringAsmDefID, dtStart, &d2)
	if err != nil {
		return bizErrSys(&err)
	}

	// rlib.Console("Number of instances to reverse: %d\n", len(m))
	for i := 0; i < len(m); i++ {
		errlist = ReverseAssessmentInstance(ctx, &m[i], dt, lc)
		if len(errlist) > 0 {
			return errlist
		}
	}

	//---------------------------------------------------------------------------
	// Since all future instances are being reversed, we need to stop generating
	// new instances.  So, we need to set the Parent stop date to dtStart.
	//---------------------------------------------------------------------------
	asm, err := rlib.GetAssessment(ctx, aold.PASMID)
	if err != nil {
		return bizErrSys(&err)
	}
	asm.Stop = *dtStart
	if err = rlib.UpdateAssessment(ctx, &asm); err != nil {
		return bizErrSys(&err)
	}

	return errlist
}

// ReverseAssessmentsAfterStop reverses all assessment instances of aold
// that occur after dtStop.  The date associated with the reversal is dt
//
// This function should be called before moving the dtStop date of a
// recurring assessment back in time.
//
// INPUTS
//    ctx     = context needed for db transactions
//    aold    = the recurring assessment
//    dtStop  = new date on which the recurring instance will be stopped.
//    dt      = time to mark when the reversal was made
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessmentsAfterStop(ctx context.Context, aold *rlib.Assessment, dtStop, dt *time.Time, lc *rlib.ClosePeriod) []BizError {
	var errlist []BizError

	rlib.Console("ENTERED: ReverseAssessmentsAfterStop\n")

	d2 := rlib.ENDOFTIME
	// rlib.Console("aold.ASMID = %d, dtStop = %s, dt = %s\n", aold.ASMID, dtStop.Format(rlib.RRDATEREPORTFMT), dt.Format(rlib.RRDATEREPORTFMT))

	m, err := rlib.GetAssessmentInstancesByParent(ctx, aold.ASMID, dtStop, &d2)
	if err != nil {
		return bizErrSys(&err)
	}

	now := rlib.Now()

	// rlib.Console("Number of instances to reverse: %d\n", len(m))
	for i := 0; i < len(m); i++ {
		if m[i].Start.After(now) {
			// rlib.Console("ReverseAssessmentsAfterStop: ATTEMPT TO REVERSE FUTURE ASSESSMENT: m[i] - ASMID = %d Start = %s\n", m[i].ASMID, m[i].Start.Format(rlib.RRDATEFMT3))
			continue
		}
		errlist = ReverseAssessmentInstance(ctx, &m[i], dt, lc)
		if len(errlist) > 0 {
			return errlist
		}
	}
	return errlist
}

// ReverseAssessmentInstance reverses a single instance of an assessment.
// If the assessment has already been reversed, we return immediately.
//
// INPUTS
//    aold = the assessment to reverse
//      dt = the time to mark for the reversal (when it was made)
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessmentInstance(ctx context.Context, aold *rlib.Assessment, dt *time.Time, lc *rlib.ClosePeriod) []BizError {
	rlib.Console("Entered ReverseAssessmentInstance:  ASMID = %d, PASMID = %d, date = %s\n", aold.ASMID, aold.PASMID, aold.Start.Format(rlib.RRDATEFMT3))
	if aold.FLAGS&0x4 != 0 {
		// rlib.Console("ReverseAssessmentInstance:  ASMID = %d has already been reversed!\n", aold.ASMID)
		// debug.PrintStack()
		return nil // it's already reversed
	}

	// I think we must allow future assessments to be reversed.  Not sure why this
	// code was put in, but I am commenting it out now. I'll validate the tests
	// all still work.
	// 4/26/2019 sman:
	// now := rlib.Now()
	// if aold.Start.After(now) {
	// 	// rlib.Console("ReverseAssessmentInstance: ATTEMPT TO REVERSE FUTURE ASSESSMENT: aold - ASMID = %d Start = %s\n", aold.ASMID, aold.Start.Format(rlib.RRDATEFMT3))
	// 	return nil
	// }

	anew := *aold
	anew.Comment = ""
	anew.ASMID = 0
	anew.Amount = -anew.Amount
	anew.RPASMID = aold.ASMID
	anew.FLAGS |= 0x4 // set bit 2 to mark that this assessment is void
	anew.AppendComment(fmt.Sprintf("Reversal of %s", aold.IDtoString()))

	// rlib.Console("RAI: anew = %#v\n", anew)

	//-----------------------------------------------------------
	// If the Start/stop dates are prior to lc, snap them to LC
	//-----------------------------------------------------------
	// rlib.Console("LC Check: anew.Start = %s,  lc.Dt = %s\n", anew.Start.Format(rlib.RRDATEREPORTFMT), lc.Dt.Format(rlib.RRDATEREPORTFMT))
	if anew.Start.Before(lc.Dt) {
		anew.AppendComment(fmt.Sprintf("Snapping %s open period %s", anew.Start.Format(rlib.RRDATEFMT3), lc.OpenPeriodDt.Format(rlib.RRDATEFMT3)))
		anew.Start = lc.OpenPeriodDt
		anew.Stop = lc.OpenPeriodDt
	}

	errlist := InsertAssessment(ctx, &anew, 1, lc)
	if len(errlist) > 0 {
		// rlib.Console("RAI: err 1\n")
		return errlist
	}

	aold.AppendComment(fmt.Sprintf("Reversed by %s", anew.IDtoString()))
	aold.FLAGS |= 0x4 // set bit 2 to mark that this assessment is void
	err := rlib.UpdateAssessment(ctx, aold)
	if err != nil {
		// rlib.Console("RAI: err 2\n")
		return bizErrSys(&err)
	}

	if aold.AGRCPTID == 0 {
		err = DeallocateAppliedFunds(ctx, aold, anew.ASMID, dt)
		if err != nil {
			// rlib.Console("RAI: err 3\n")
			return bizErrSys(&err)
		}
	} else {
		//---------------------------------------------------------
		// handle auto-generated assessments a little different...
		// See if there was a funds transfer to a bank account...
		//---------------------------------------------------------
		be := ReverseAutoGenAsmt(ctx, aold)
		if len(be) > 0 {
			// rlib.Console("RAI: err 4\n")
			return be
		}
	}
	// rlib.Console("Exiting ReverseAssessmentInstance\n")
	return nil
}

// ReverseAutoGenAsmt - Removes handles Journal, JournalAllocation, and
// GLAccount adjustments for auto-generated assessments for things like
// floating deposits and application fees.
//
// INPUTS
//    aold = the assessment to reverse
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAutoGenAsmt(ctx context.Context, aold *rlib.Assessment) []BizError {
	const funcname = "ReverseAutoGenAsmt"
	var (
		err error
	)

	jx, err := rlib.GetJournalByTypeAndID(ctx, rlib.JNLTYPEXFER, aold.AGRCPTID)
	if err != nil {
		return bizErrSys(&err)
	}

	if jx.JID > 0 {
		err = rlib.GetJournalAllocations(ctx, &jx)
		if err != nil {
			return bizErrSys(&err)
		}

		if len(jx.JA) > 0 {
			m := rlib.ParseSimpleAcctRule(jx.JA[0].AcctRule)

			//--------------------
			// journal
			//--------------------
			jnl, err := rlib.GetJournal(ctx, jx.JA[0].JID)
			if err != nil {
				return bizErrSys(&err)
			}

			jnl.Comment = fmt.Sprintf("Reversal of J-%d", jnl.JID)
			jnl.JID = 0
			jnl.Amount = -jnl.Amount

			_, err = rlib.InsertJournal(ctx, &jnl) // this will update jnl.JID
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return bizErrSys(&err)
			}

			//--------------------
			// journal allocation
			//--------------------
			ja := jx.JA[0] // copy of the original we're reversing
			ja.JID = jnl.JID
			ja.AcctRule = fmt.Sprintf("%s %s %.4f, %s %s %.4f",
				m[0].Action, m[0].Account, -m[0].Amount,
				m[1].Action, m[1].Account, -m[1].Amount)
			ja.Amount = -ja.Amount
			_, err = rlib.InsertJournalAllocationEntry(ctx, &ja)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return bizErrSys(&err)
			}

			//-------------
			// ledgers
			//-------------
			n, err := rlib.GetLedgerEntriesByJAID(ctx, aold.BID, jx.JA[0].JAID)
			if err != nil {
				return bizErrSys(&err)
			}

			for i := 0; i < len(n); i++ {
				le := n[i]
				le.LEID = 0
				le.Comment = fmt.Sprintf("reversal of LE-%d", n[i].LEID)
				le.Amount = -le.Amount
				le.JAID = ja.JAID
				le.JID = jx.JA[0].JID
				_, err = rlib.InsertLedgerEntry(ctx, &le)
				if err != nil {
					rlib.LogAndPrintError(funcname, err)
					return bizErrSys(&err)
				}
			}
		}
	}
	return nil
}

// DeallocateAppliedFunds - Removes any funds applied to this assessment.
// INPUTS
//    a         = receipt to be voided
//    asmtRevID = ASMID of the reversal assessment
//    dt        = time we want the funds to be marked as deallocated
//
// RETURNS
//    any error that occurred, or nil if no error
//-------------------------------------------------------------------------------
func DeallocateAppliedFunds(ctx context.Context, a *rlib.Assessment, asmtRevID int64, dt *time.Time) error {
	funcname := "bizlogic.DeallocateAppliedFunds"
	//--------------------------------------------------------------
	// Find all JournalAllocations that reference Assessment a that
	// also have a ReceiptID.
	//--------------------------------------------------------------
	JA, err := rlib.GetJournalAllocationByASMID(ctx, a.ASMID)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}

	for i := 0; i < len(JA); i++ {
		if JA[i].RCPTID == 0 {
			continue
		}

		rcpt, err := rlib.GetReceipt(ctx, JA[i].RCPTID)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		//--------------------------------
		// Reverse the Journal Entry...
		//--------------------------------
		var jnl = rlib.Journal{
			BID:    rcpt.BID,
			Amount: -JA[i].Amount, // reverse the amount
			Type:   rlib.JNLTYPEASMT,
			ID:     asmtRevID, // this is the rcptid of the reversal receipt
			Dt:     *dt,       // reversal date
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
		n, err := rlib.ParseAcctRule(ctx, &xbiz1, 0, dt, dt, JA[i].AcctRule, 0, 1.0)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		acctrule := ""
		// revAcctRule := ""
		for k := 0; k < len(n); k++ {
			acctrule += fmt.Sprintf("ASM(%d) %s %s %.2f", JA[i].ASMID, n[k].Action, n[k].Account, jnl.Amount)
			// revAcctRule += fmt.Sprintf("ASM(%d) %s %s %.2f", JA[i].ASMID, n[k].Action, n[k].Account, -jnl.Amount)
			if k+1 < len(n) {
				acctrule += ","
				// revAcctRule += ","
			}
		}
		var ja = rlib.JournalAllocation{
			JID:      jnl.JID,
			AcctRule: acctrule,
			Amount:   jnl.Amount,
			BID:      jnl.BID,
			RAID:     JA[i].RAID,
			RID:      JA[i].RID,
			ASMID:    JA[i].ASMID,
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
		// Next, reverse the ledger entries...
		//-------------------------------------------------------------------------
		le, err := rlib.GetLedgerEntriesByJAID(ctx, rcpt.BID, JA[i].JAID)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
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
		// Next, reverse the receiptAllocation for this assessment...
		//-------------------------------------------------------------------------
		m, err := rlib.GetReceiptAllocationsByASMID(ctx, rcpt.BID, a.ASMID)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		for k := 0; k < len(m); k++ {
			// if this allocation does not reference a.ASMID, then skip it
			// also, if it's already reversed, skip it
			if m[k].ASMID != a.ASMID || m[k].FLAGS&4 != 0 {
				continue
			}
			m[k].FLAGS |= 0x4 // set bit 2 to indicate that this is a voided entry
			vra := m[k]
			vra.Amount = -vra.Amount
			vra.AcctRule = acctrule
			vra.Dt = *dt
			vra.RAID = ja.RAID
			_, err = rlib.InsertReceiptAllocation(ctx, &vra)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return err
			}

			err = rlib.UpdateReceiptAllocation(ctx, &m[k]) // update its flags to indicate it is voided
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return err
			}
		}

		//-------------------------------------------------------------------------
		// Next, mark the flag on the receipt indicating some or all of its funds
		// are now available. This journal allocation (JA[i]) is being deallocated
		// so those funds are now available from the receipt...
		//-------------------------------------------------------------------------
		err = rlib.GetReceiptAllocations(ctx, rcpt.RCPTID, &rcpt)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		rar := ""
		for k := 0; k < len(rcpt.RA); k++ {
			if rcpt.RA[k].ASMID == 0 {
				continue // only want the entries for AcctRuleApply
			}
			rar += rcpt.RA[k].AcctRule
			if k+1 < len(rcpt.RA) {
				rar += ","
			}
		}

		nar, err := rlib.ParseAcctRule(ctx, &xbiz1, 0, dt, dt, rar, 0, 1.0)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		tot := rcpt.Amount
		for i := 0; i < len(nar); i++ {
			if "d" == nar[i].Action {
				tot -= nar[i].Amount
			}
		}
		f := uint64(0) // assume it's all available
		if tot != rcpt.Amount {
			f = 1
		}
		rcpt.FLAGS &= ^(uint64(0x3)) // remove whatever status was there before
		rcpt.FLAGS |= f              // 0 = the entire amount is available, 1 = some is still available
		rcpt.AcctRuleApply = rar

		err = rlib.UpdateReceipt(ctx, &rcpt)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		//-------------------------------------------------------------------------
		// Finally, update the assessment that was allocated payment from this receipt...
		//-------------------------------------------------------------------------
		unpaid := AssessmentUnpaidPortion(ctx, a) // how much of this assessment is still unpaid?
		paid := a.Amount - unpaid                 // how much remains to be paid
		remaining := paid - JA[i].Amount          // how much remains after removing this allocation

		newflags := uint64(0) // assume nothing has been paid on the assessment after this reversal
		if remaining > 0 {    // if any portion has still been paid...
			newflags = uint64(1) // ... then mark as partially paid
		}
		a.FLAGS &= ^(uint64(0x3)) // clear the bits of interest
		a.FLAGS |= newflags | 0x4 // set new status and mark as voided

		err = rlib.UpdateAssessment(ctx, a)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}
	}
	return nil
}

// InsertAssessment performs bizlogic checks first, then inserts the Assessment,
// then adds the associated Journal and Ledger entries
//
// INPUTS
//    a = the assessment to insert
//  exp = if it is a recurring assessment and the start date is in the past, should
//        past entries be created? 0 = NO, non-zero = YES
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func InsertAssessment(ctx context.Context, a *rlib.Assessment, exp int, lc *rlib.ClosePeriod) []BizError {
	// funcname := "bizlogic.InsertAssessment"
	// rlib.Console("*************************************************************************\n")
	// rlib.Console("Entered %s\n", funcname)
	// rlib.Console("Amount = %8.2f, RentCycle = %d, Start = %s, Stop = %s\n", a.Amount, a.RentCycle, a.Start.Format(rlib.RRDATEFMT3), a.Stop.Format(rlib.RRDATEFMT3))
	// rlib.Console("lc.Dt = %s\n", lc.Dt.Format(rlib.RRDATEREPORTFMT))
	// rlib.Console("lc.CPID= %d, lc.Dt = %s\n", lc.CPID, lc.Dt.Format(rlib.RRDATEFMT3))
	// rlib.Console("lc.ExpandAsmDtStart = %s\n", lc.ExpandAsmDtStart.Format(rlib.RRDATEFMT3))
	// rlib.Console("lc.ExpandAsmDtStop = %s\n", lc.ExpandAsmDtStop.Format(rlib.RRDATEREPORTFMT))
	// rlib.Console("*************************************************************************\n")

	var errlist []BizError
	errlist = ValidateAssessment(ctx, a) // Make sure there are no bizlogic errors before saving
	if len(errlist) > 0 {
		// rlib.Console("%s:  failed ValidateAssessment\n", funcname)
		return errlist
	}

	// rlib.Console("%s: A.  a.ASMID = %d, a.BID = %d, a.ARID = %d, a.RAID = %d, amount = %8.2f\n", funcname, a.ASMID, a.BID, a.ARID, a.RAID, a.Amount)
	var xbiz rlib.XBusiness
	err := rlib.InitBizInternals(a.BID, &xbiz)
	if err != nil {
		// rlib.Console("%s:  failed to InitBizInternals\n", funcname)
		return bizErrSys(&err)
	}

	//-------------------------------------------------------------------------
	// If the AcctRule sends money to an offset account, mark it as an offset.
	//-------------------------------------------------------------------------
	clid := rlib.RRdb.BizTypes[a.BID].AR[a.ARID].CreditLID // this is the assessment's Account Rule credit ledger
	dlid := rlib.RRdb.BizTypes[a.BID].AR[a.ARID].DebitLID  // this is the assessment's Account Rule debit ledger

	if rlib.RRdb.BizTypes[a.BID].GLAccounts[dlid].FLAGS&0x1 > 0 || rlib.RRdb.BizTypes[a.BID].GLAccounts[clid].FLAGS&0x1 > 0 {
		//a.FLAGS &= 0x8ffffffffffffffc //zero bits 0:1
		a.FLAGS |= 0x3 // indicate that this is an OFFSET and should not be processd during payment allocation
	}

	//------------------------------------------------------------
	// Save original start date before close-period-snap-check
	// because we may need to expand instances...
	//------------------------------------------------------------
	originalStart := a.Start
	originalComment := a.Comment

	//------------------------------------------------------------
	// If start date is prior to last close, snap to last close
	//------------------------------------------------------------
	if a.Start.Before(lc.Dt) {
		// rlib.Console(">>>>> Attempt to create an assessment in a closed period\n")
		// rlib.Console(">>>>> requested start date = %s\n", a.Start.Format(rlib.RRDATEFMT3))
		atype := "instance" // update assumption
		if a.RentCycle > rlib.RECURNONE {
			// rlib.Console(">>>>> \n")
			//-----------------------------------------------------------------
			// create it at the first epoch date on or after the open period
			//-----------------------------------------------------------------
			// d := rlib.GetNextInstance(*lc, a.RentCycle)
			// rlib.Console("BEFORE rlib.InstanceDateCoveringDate lc.ExpandAsmDtStart = %s\n", lc.ExpandAsmDtStart.Format(rlib.RRDATEFMT3))
			d := rlib.InstanceDateCoveringDate(&a.Start, &lc.OpenPeriodDt, a.RentCycle)
			// rlib.Console(">>>>> assesment recurs, so snapping to next instance on or after: %s\n", lc.OpenPeriodDt.Format(rlib.RRDATEFMT3))
			// rlib.Console(">>>>> next instance is: %s\n", d.Format(rlib.RRDATEFMT3))
			lc.ExpandAsmDtStart = a.Start // save this for expansion if needed
			lc.ExpandAsmDtStop = a.Stop   // don't expand beyond this date
			// rlib.Console("AFTER >>>>> ExpandAsmDtStart = %s\n", lc.ExpandAsmDtStart.Format(rlib.RRDATEFMT3))
			a.Start = d
			if a.PASMID > 0 { // this is an instance, so set the stop date as well...
				a.Stop = d
			} else { // this is a definition.  make sure the start date is before the stop date
				atype = "definition" // assume this
				if a.Start.After(a.Stop) {
					a.Stop = a.Start // move to free area
				}
			}
		} else {
			//------------------------------------------------------------------
			// This is a non-recurring assessment - move it to the open period
			//------------------------------------------------------------------
			a.Start = lc.OpenPeriodDt
			a.Stop = lc.OpenPeriodDt
		}
		a.AppendComment(fmt.Sprintf("BL- Snapped %s %s to open period %s", atype, originalStart.Format(rlib.RRDATEFMT3), lc.OpenPeriodDt.Format(rlib.RRDATEFMT3)))
		// rlib.Console(">>>>> after adjustments: Start = %s, Stop = %s\n", a.Start.Format(rlib.RRDATEFMT3), a.Stop.Format(rlib.RRDATEFMT3))
	}

	// rlib.Console("B:   insert assessment\n")
	_, err = rlib.InsertAssessment(ctx, a) // No bizlogic errors, save it
	if err != nil {
		return bizErrSys(&err)
	}

	// rlib.Console("C\n")
	//------------------------------------------------
	// Add the journal and ledger entries
	//------------------------------------------------
	err = rlib.GetXBusiness(ctx, a.BID, &xbiz)
	if err != nil {
		return bizErrSys(&err)
	}

	d1, d2 := rlib.GetMonthPeriodForDate(&a.Start) // TODO: probably needs to be more generalized
	// rlib.Console("C - 0.5:  d1,d2 = %s\n", // rlib.ConsoleDRange(&d1, &d2))
	// rlib.Console("    Note that lc.ExpandAsmDtStart,Stop = %s\n", // rlib.ConsoleDRange(&lc.ExpandAsmDtStart, &lc.ExpandAsmDtStop))

	rlib.InitLedgerCache()

	a.Comment = originalComment        // set this back to the original comment, it will be updated as necessary
	if a.RentCycle == rlib.RECURNONE { // for nonrecurring, use existng struct: a
		rlib.ExpandAssessment(ctx, a, &xbiz, &d1, &d2, true, lc) // generates assessment instances
	} else if exp != 0 && a.PASMID == 0 && 0 == (a.FLAGS&(1<<6)) { // only expand if we're asked and if we're not an instance, and not a single instanced assessment
		// rlib.Console("C1\n")
		now := rlib.DateAtTimeZero(rlib.Now())
		dt := rlib.DateAtTimeZero(a.Start)
		if !dt.After(now) {
			// rlib.Console("C2  lc.ExpandAsmDtStart = %s\n", rlib.ConDt(&lc.ExpandAsmDtStart))
			err := createInstancesToDate(ctx, a, &xbiz, lc, &originalStart)
			if err != nil {
				return bizErrSys(&err)
			}
		}
	}
	// rlib.Console("D - exiting %s\n", funcname)
	return nil
}

// ValidateAssessment checks to see whether the assessment violates any
// business logic.
//
// INPUTS
//    a = the assessment to validate
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ValidateAssessment(ctx context.Context, a *rlib.Assessment) []BizError {
	var e []BizError
	var raid, bid int64
	var row *sql.Row
	// rlib.Console("ValidateAssessment: A\n")
	//----------------------------------------------
	// Validate that we have a RAID that exists...
	//----------------------------------------------
	qry := fmt.Sprintf("SELECT RAID,BID FROM RentalAgreement WHERE RAID=%d", a.RAID)
	// err := rlib.RRdb.Dbrr.QueryRow(qry).Scan(&raid, &bid)
	if tx, ok := rlib.DBTxFromContext(ctx); ok { // if transaction is supplied
		row = tx.QueryRow(qry)
	} else {
		row = rlib.RRdb.Dbrr.QueryRow(qry)
	}
	err := row.Scan(&raid, &bid)
	if err != nil {
		if rlib.IsSQLNoResultsError(err) {
			s := fmt.Sprintf(BizErrors[UnknownRAID].Message, a.RAID, a.BID)
			b := BizError{Errno: UnknownRAID, Message: s}
			e = append(e, b)
			// rlib.Console("ValidateAssessment: B\n")
		} else {
			return bizErrSys(&err)
		}
	}
	//----------------------------------------------------
	// Validate that it is part of the same Business...
	//----------------------------------------------------
	// rlib.Console("ValidateAssessment: C0\n")
	if bid != a.BID {
		s := fmt.Sprintf(BizErrors[UnknownRAID].Message, a.RAID, a.BID)
		b := BizError{Errno: UnknownRAID, Message: s}
		// rlib.Console("ValidateAssessment: C\n")
		e = append(e, b)
	}

	// rlib.Console("ValidateAssessment: D0\n")
	if a.RID > 0 {
		// rlib.Console("ValidateAssessment: D1\n")
		//--------------------------------------------------------------------------
		//  Check for assessment timeframe prior to or after Rentable's type being defined
		//--------------------------------------------------------------------------
		rtl, err := rlib.GetRentableTypeRefs(ctx, a.RID) // these are returned in chronological order
		if err != nil {
			elist := bizErrSys(&err)
			// rlib.Console("ValidateAssessment: D\n")
			e = append(e, elist[0])
		}

		// rlib.Console("ValidateAssessment: E0\n")
		l := len(rtl)
		// rlib.Console("len(rtl) = %d\n", l)
		if l == 0 {
			// rlib.Console("ValidateAssessment: E\n")
			e = append(e, BizErrors[RentableTypeUnknown])
		} else {
			if a.Stop.Before(rtl[0].DtStart) || a.Start.After(rtl[l-1].DtStop) {
				// rlib.Console("ASMID = %d\n\tStart = %s, Stop = %s\n\tRentableType[0].start = %s, RentableType[%d].stop = %s", a.ASMID, a.Start.Format(rlib.RRDATEREPORTFMT), a.Stop.Format(rlib.RRDATEREPORTFMT), rtl[0].DtStart.Format(rlib.RRDATEREPORTFMT), l-1, rtl[l-1].DtStop.Format(rlib.RRDATEREPORTFMT))
				// rlib.Console("ValidateAssessment: F\n")
				e = append(e, BizErrors[RentableTypeUnknown])
			}
		}

		//--------------------------------------------------------------------------
		//  Check for assessment timeframe prior to or after Rentable's status being defined
		//--------------------------------------------------------------------------
		// rlib.Console("ValidateAssessment: G0\n")
		// rlib.Console("ValidateAssessment: a.Start-Stop = %s - %s\n", a.Start.Format(rlib.RRDATEINPFMT), a.Stop.Format(rlib.RRDATEINPFMT))
		rsl, err := rlib.GetRentableUseStatusByRange(ctx, a.RID, &a.Start, &a.Stop)
		if err != nil {
			elist := bizErrSys(&err)
			// rlib.Console("ValidateAssessment: G\n")
			e = append(e, elist[0])
		}

		// rlib.Console("ValidateAssessment: H0\n")
		l = len(rsl)
		// rlib.Console("len(rtl) = %d, len(rsl) = %d\n", len(rtl), len(rsl))
		if l == 0 {
			// rlib.Console("ValidateAssessment: H\n")
			// rlib.Console("ValidateAssessment: l=0 --> GetRentableUseStatusByRange( a.RID=%d, %s)\n", a.RID, rlib.ConsoleDRange(&a.Start, &a.Stop))
			e = append(e, BizErrors[RentableUseStatusUnknown])
		} else {
			// rlib.Console("ValidateAssessment: H1\n")
			// rlib.Console("ValidateAssessment: rtl = %s - %s\n", rsl[0].DtStart.Format(rlib.RRDATEINPFMT), rsl[l-1].DtStop.Format(rlib.RRDATEINPFMT))
			if a.Stop.Before(rsl[0].DtStart) || a.Start.After(rsl[l-1].DtStop) {
				// rlib.Console("ValidateAssessment: I\n")
				e = append(e, BizErrors[RentableUseStatusUnknown])
			}
		}

		//--------------------------------------------------------------------------
		//  If the assessment is non-recurring - then start and stop date should
		//  be the same.
		// NOTE: removed 9/11/2018 - sman
		//--------------------------------------------------------------------------
		// if a.RentCycle == rlib.RECURNONE && !a.Start.Equal(a.Stop) {
		// 	e = append(e, BizErrors[AsmDateRangeNotAllowed])
		// }

		//--------------------------------------------------------------------------
		// We cannot have the start date occurring after the stop date
		//--------------------------------------------------------------------------
		// rlib.Console("ValidateAssessment: J0\n")
		if a.Stop.Before(a.Start) {
			rlib.Console("ValidateAssessment: J  - start after stop!   a.Start / a.Stop = %s\n", rlib.ConsoleDRange(&a.Start, &a.Stop))
			e = append(e, BizErrors[StartDateAfterStopDate])
		}
	}
	// rlib.Console("ValidateAssessment: exiting.  len(errlist) = %d\n", len(e))
	return e
}

// createInstancesToDate creates all instances of a recurring Assessments up to the
// supplied date
//
// INPUTS
//     a = the recurring assessment
//  xbiz = Business information
//    lc = date to use which is past the last close date
// start = start date on assessment before any snapping that occurred due to
//         checking closed period
//
// RETURNS
//
//-------------------------------------------------------------------------------------
func createInstancesToDate(ctx context.Context, a *rlib.Assessment, xbiz *rlib.XBusiness, lc *rlib.ClosePeriod, start *time.Time) error {
	// rlib.Console("\n\n*** Entered createInstancesToDate   start = %s,  lc.dt = %s\n", start.Format(rlib.RRDATEFMT3), lc.Dt.Format(rlib.RRDATEFMT3))
	// rlib.Console("a.Start = %s, a.Stop = %s\n", a.Start.Format(rlib.RRDATEFMT3), a.Stop.Format(rlib.RRDATEFMT3))

	//-------------------------------------------------------------------------
	// Create instances from original start date to now. ExpandAssessment will
	// handle expanding the instances and snapping to open period dates as necessary.
	//-------------------------------------------------------------------------
	d1 := *start
	d2 := rlib.Now()                                                          // never go further in the future than the current time
	if lc.ExpandAsmDtStop.After(rlib.TIME0) && d2.After(lc.ExpandAsmDtStop) { // snap to expansion stop date if needed
		d2 = a.Stop
	}

	// rlib.Console("\n\n**** createInstancesToDate calling ExpandAssessment ASMID = %d, d1 = %s, d2 = %s\n", a.ASMID, d1.Format(rlib.RRDATEFMT3), d2.Format(rlib.RRDATEFMT3))
	// rlib.Console("createInstancesToDate: lc.ExpandAsmDtStart,Stop = %s\n", rlib.ConsoleDRange(&lc.ExpandAsmDtStart, &lc.ExpandAsmDtStop))
	err := rlib.ExpandAssessment(ctx, a, xbiz, &d1, &d2, true, lc) // this generates the assessment instances
	if err != nil {
		// rlib.Console("Exiting createInstancesToDate err = %s\n", err.Error())
		return err
	}

	// rlib.Console("Exiting createInstancesToDate\n\n\n")
	return nil
}
