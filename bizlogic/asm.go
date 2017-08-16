package bizlogic

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// UpdateAssessment updates the supplied assessment, reversing existing assessments
// if necessary
//
// INPUTS
//    a = the assessment to insert
//  exp = if it is a recurring assessment and the start date is in the past, should
//        past entries be created?  true = yes
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func UpdateAssessment(anew *rlib.Assessment, mode int, dt *time.Time, exp int) []BizError {
	var err error
	var errlist []BizError

	rlib.Console("Entered bizlogic.UpdateAssessment:  anew.ASMID = %d, mode = %d, dt = %s, exp = %d\n", anew.ASMID, mode, dt.Format(rlib.RRDATEREPORTFMT), exp)
	rlib.Console("anew.FLAGS = %X\n", anew.FLAGS)

	if anew.FLAGS&0x4 != 0 {
		errlist = append(errlist, BizErrors[EditReversal])
		return errlist
	}
	//-------------------------------
	// Load existing assessment...
	//-------------------------------
	aold, err := rlib.GetAssessment(anew.ASMID)
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
		errlist = ReverseAssessment(&aold, mode, dt) // reverse the assessment itself
		if errlist != nil {
			return errlist
		}
		errlist = InsertAssessment(anew, exp) // Finally, insert the new assessment...
		if err != nil {
			return errlist
		}
	}

	err = rlib.UpdateAssessment(anew) // reversal not needed, just update the assessment
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
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessment(aold *rlib.Assessment, mode int, dt *time.Time) []BizError {
	funcname := "bizlogic.ReverseAssessment"
	var errlist []BizError
	rlib.Console("Entered ReverseAssessment\n")
	if aold.PASMID == 0 {
		mode = 2 // force behavior on the epoch
	}
	switch mode {
	case 0:
		errlist = ReverseAssessmentInstance(aold, dt)
	case 1:
		errlist = ReverseAssessmentsGoingForward(aold, &aold.Start, dt)
	case 2:
		var epoch, inst rlib.Assessment
		var err error

		//---------------------------------------------------------
		// set the epoch
		//---------------------------------------------------------
		if aold.PASMID != 0 {
			epoch, err = rlib.GetAssessment(aold.PASMID)
			if err != nil {
				return bizErrSys(&err)
			}
		} else {
			epoch = *aold
		}

		//---------------------------------------------------------
		// If it is not recurring then reverse it and we're done
		//---------------------------------------------------------
		if epoch.RentCycle == rlib.RECURNONE {
			return ReverseAssessmentInstance(&epoch, dt)
		}

		//---------------------------------------------------------
		// Get the first instance and modify forward...
		//---------------------------------------------------------
		inst, err = rlib.GetAssessmentFirstInstance(epoch.ASMID)
		if err != nil {
			return bizErrSys(&err)
		}
		errlist = ReverseAssessmentsGoingForward(&inst, &inst.Start, dt) // reverse from start of recurring instances forward
		if len(errlist) > 0 {
			return errlist
		}
		epoch.FLAGS |= 0x4 // mark that this is void
		err = rlib.UpdateAssessment(&epoch)
		if err != nil {
			return bizErrSys(&err)
		}

	default:
		err := fmt.Errorf("%s:  unsupported mode: %d", funcname, mode)
		rlib.LogAndPrintError(funcname, err)
		return bizErrSys(&err)
	}
	return errlist
}

// ReverseAssessmentsGoingForward reverses an existing assessment
//
// INPUTS
//    aold = the first in a series of assessments to reverse
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessmentsGoingForward(aold *rlib.Assessment, dtStart, dt *time.Time) []BizError {
	var errlist []BizError

	rlib.Console("ENTERED: ReverseAssessmentsGoingForward\n")

	d2 := time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)
	rlib.Console("aold.PASMID = %d, dtStart = %s, dt = %s\n", aold.PASMID, dtStart.Format(rlib.RRDATEREPORTFMT), dt.Format(rlib.RRDATEREPORTFMT))

	m := rlib.GetAssessmentInstancesByParent(aold.PASMID, dtStart, &d2)
	rlib.Console("Number of instances to reverse: %d\n", len(m))
	for i := 0; i < len(m); i++ {
		errlist = ReverseAssessmentInstance(&m[i], dt)
		if len(errlist) > 0 {
			return errlist
		}
	}
	return nil
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
func ReverseAssessmentInstance(aold *rlib.Assessment, dt *time.Time) []BizError {
	if aold.FLAGS&0x4 != 0 {
		return nil // it's already reversed
	}

	anew := *aold
	anew.ASMID = 0
	anew.Amount = -anew.Amount
	anew.RPASMID = aold.ASMID
	anew.FLAGS |= 0x4 // set bit 2 to mark that this assessment is void
	anew.Comment = fmt.Sprintf("Reversal of %s", aold.IDtoString())

	errlist := InsertAssessment(&anew, 1)
	if len(errlist) > 0 {
		return errlist
	}

	aold.Comment = fmt.Sprintf("Reversed by %s", anew.IDtoString())
	aold.FLAGS |= 0x4 // set bit 2 to mark that this assessment is void
	err := rlib.UpdateAssessment(aold)
	if err != nil {
		return bizErrSys(&err)
	}

	err = DeallocateAppliedFunds(aold, anew.ASMID, dt)
	if err != nil {
		return bizErrSys(&err)
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
func DeallocateAppliedFunds(a *rlib.Assessment, asmtRevID int64, dt *time.Time) error {
	funcname := "bizlogic.DeallocateAppliedFunds"
	//--------------------------------------------------------------
	// Find all JournalAllocations that reference Assessment a that
	// also have a ReceiptID.
	//--------------------------------------------------------------
	JA := rlib.GetJournalAllocationByASMID(a.ASMID)
	for i := 0; i < len(JA); i++ {
		if JA[i].RCPTID == 0 {
			continue
		}

		rcpt := rlib.GetReceipt(JA[i].RCPTID)

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
		_, err := rlib.InsertJournal(&jnl)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		//-------------------------------------------------------------------------
		// Next, add the JournalAllocation reversal
		//-------------------------------------------------------------------------
		var xbiz1 rlib.XBusiness // not actually used
		n := rlib.ParseAcctRule(&xbiz1, 0, dt, dt, JA[i].AcctRule, 0, 1.0)
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
		rlib.InsertJournalAllocationEntry(&ja)
		jnl.JA = append(jnl.JA, ja)

		//-------------------------------------------------------------------------
		// Next, reverse the ledger entries...
		//-------------------------------------------------------------------------
		le := rlib.GetLedgerEntriesByJAID(rcpt.BID, JA[i].JAID)
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
		// Next, reverse the receiptAllocation for this assessment...
		//-------------------------------------------------------------------------
		m := rlib.GetReceiptAllocationsByASMID(rcpt.BID, a.ASMID)
		for k := 0; k < len(m); k++ {
			m[k].FLAGS |= 0x4 // set bit 2 to indicate that this is a voided entry
			vra := m[k]
			vra.Amount = -vra.Amount
			vra.AcctRule = acctrule
			vra.Dt = *dt
			vra.RAID = ja.RAID
			_, err = rlib.InsertReceiptAllocation(&vra)
			if err != nil {
				return err
			}
			err := rlib.UpdateReceiptAllocation(&m[k]) // update its flags to indicate it is voided
			if err != nil {
				return err
			}
		}

		//-------------------------------------------------------------------------
		// Next, mark the flag on the receipt indicating some or all of its funds
		// are now available. This journal allocation (JA[i]) is being deallocated
		// so those funds are now available from the receipt...
		//-------------------------------------------------------------------------
		rlib.GetReceiptAllocations(rcpt.RCPTID, &rcpt)
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

		nar := rlib.ParseAcctRule(&xbiz1, 0, dt, dt, rar, 0, 1.0)
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
		rlib.UpdateReceipt(&rcpt)

		//-------------------------------------------------------------------------
		// Finally, update the assessment that was allocated payment from this receipt...
		//-------------------------------------------------------------------------
		unpaid := AssessmentUnpaidPortion(a) // how much of this assessment is still unpaid?
		paid := a.Amount - unpaid            // how much remains to be paid
		remaining := paid - JA[i].Amount     // how much remains after removing this allocation

		newflags := uint64(0) // assume nothing has been paid on the assessment after this reversal
		if remaining > 0 {    // if any portion has still been paid...
			newflags = uint64(1) // ... then mark as partially paid
		}
		a.FLAGS &= ^(uint64(0x3)) // clear the bits of interest
		a.FLAGS |= newflags | 0x4 // set new status and mark as voided
		err = rlib.UpdateAssessment(a)
		if err != nil {
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
//        past entries be created?  true = yes
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func InsertAssessment(a *rlib.Assessment, exp int) []BizError {
	funcname := "bizlogic.InsertAssessment"
	rlib.Console("Entered %s\n", funcname)
	var errlist []BizError
	errlist = ValidateAssessment(a) // Make sure there are no bizlogic errors before saving
	if len(errlist) > 0 {
		return errlist
	}

	rlib.Console("A.  a.BID = %d, a.ARID = %d\n", a.BID, a.ARID)
	var xbiz rlib.XBusiness
	rlib.InitBizInternals(a.BID, &xbiz)

	//-------------------------------------------------------------------------
	// If the AcctRule sends money to an offset account, mark it as an offset.
	//-------------------------------------------------------------------------
	clid := rlib.RRdb.BizTypes[a.BID].AR[a.ARID].CreditLID // this is the assessment's Account Rule credit ledger
	dlid := rlib.RRdb.BizTypes[a.BID].AR[a.ARID].DebitLID  // this is the assessment's Account Rule debit ledger

	// rlib.Console("Pay Assessment: Assessment Rule:  Debit %s, Credit %s\n", rlib.RRdb.BizTypes[a.BID].GLAccounts[car.DebitLID].Name, rlib.RRdb.BizTypes[a.BID].GLAccounts[car.CreditLID].Name)
	// rlib.Console("Pay Assessment:    Receipt Rule:  Debit %s, Credit %s\n", rlib.RRdb.BizTypes[a.BID].GLAccounts[dar.DebitLID].Name, rlib.RRdb.BizTypes[a.BID].GLAccounts[dar.CreditLID].Name)

	if rlib.RRdb.BizTypes[a.BID].GLAccounts[dlid].FLAGS&0x1 > 0 || rlib.RRdb.BizTypes[a.BID].GLAccounts[clid].FLAGS&0x1 > 0 {
		a.FLAGS &= 0x8ffffffffffffffc //zero bits 0:1
		a.FLAGS |= 0x3                // indicate that this is an OFFSET and should not be processd during payment allocation
	}

	rlib.Console("B\n")
	_, err := rlib.InsertAssessment(a) // No bizlogic errors, save it
	if err != nil {
		return bizErrSys(&err)
	}

	rlib.Console("C\n")
	//------------------------------------------------
	// Add the journal and ledger entries
	//------------------------------------------------
	rlib.GetXBusiness(a.BID, &xbiz)
	d1, d2 := rlib.GetMonthPeriodForDate(&a.Start) // TODO: probably needs to be more generalized
	rlib.InitLedgerCache()
	if a.RentCycle == rlib.RECURNONE { // for nonrecurring, use existng struct: a
		rlib.ProcessJournalEntry(a, &xbiz, &d1, &d2, true)
	} else if exp != 0 && a.PASMID == 0 { // only expand if we're asked and if we're not an instance
		now := rlib.DateAtTimeZero(time.Now())
		dt := rlib.DateAtTimeZero(a.Start)
		if !dt.After(now) {
			createInstancesToDate(a, &xbiz)
		}
	}
	rlib.Console("D\n")
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
func ValidateAssessment(a *rlib.Assessment) []BizError {
	var e []BizError
	if a.RID > 0 {
		//--------------------------------------------------------------------------
		//  Check for assessment timeframe prior to or after Rentable's type being defined
		//--------------------------------------------------------------------------
		rtl := rlib.GetRentableTypeRefs(a.RID) // these are returned in chronological order
		l := len(rtl)
		if l == 0 {
			e = append(e, BizErrors[RentableTypeUnknown])
		} else {
			if a.Stop.Before(rtl[0].DtStart) || a.Start.After(rtl[l-1].DtStop) {
				e = append(e, BizErrors[RentableTypeUnknown])
			}
		}

		//--------------------------------------------------------------------------
		//  Check for assessment timeframe prior to or after Rentable's status being defined
		//--------------------------------------------------------------------------
		rsl := rlib.GetRentableStatusByRange(a.RID, &a.Start, &a.Stop)
		l = len(rsl)
		if l == 0 {
			// rlib.Console("ValidateAssessment: l=0\n")
			e = append(e, BizErrors[RentableStatusUnknown])
		} else {
			// rlib.Console("ValidateAssessment: a.Start-Stop = %s - %s\n", a.Start.Format(rlib.RRDATEINPFMT), a.Stop.Format(rlib.RRDATEINPFMT))
			// rlib.Console("ValidateAssessment: rtl = %s - %s\n", rtl[0].DtStart.Format(rlib.RRDATEINPFMT), rtl[l-1].DtStop.Format(rlib.RRDATEINPFMT))
			if a.Stop.Before(rtl[0].DtStart) || a.Start.After(rtl[l-1].DtStop) {
				e = append(e, BizErrors[RentableStatusUnknown])
			}
		}

	}
	return e
}

// createInstancesToDate creates all instances of a recurring Assessments up to the
// supplied date
//
// INPUTS
//    a = the recurring assessment
// xbiz = Business information
//
// RETURNS
//
//-------------------------------------------------------------------------------------
func createInstancesToDate(a *rlib.Assessment, xbiz *rlib.XBusiness) {
	now := time.Now()
	as := time.Date(a.Start.Year(), a.Start.Month(), a.Start.Day(), 0, 0, 0, 0, time.UTC)
	m := rlib.GetRecurrences(&a.Start, &a.Stop, &as, &now, a.RentCycle) // get all from the begining up to now
	for i := 0; i < len(m); i++ {
		dt1, dt2 := rlib.GetMonthPeriodForDate(&m[i])
		rlib.ProcessJournalEntry(a, xbiz, &dt1, &dt2, true) // this generates the assessment instances
	}
}
