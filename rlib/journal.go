package rlib

import (
	"context"
	"fmt"
	"time"
)

//=================================================================================================
func sumAllocations(m *[]AcctRule) (float64, float64) {
	sum := float64(0.0)
	debits := float64(0.0)
	for i := 0; i < len(*m); i++ {
		if (*m)[i].Action == "c" {
			sum -= (*m)[i].Amount
		} else {
			sum += (*m)[i].Amount
			debits += (*m)[i].Amount
		}
	}
	return sum, debits
}

// builds the account rule based on an ARID
func buildRule(ctx context.Context, id int64) (string, error) {
	rule, err := GetAR(ctx, id)
	if err != nil {
		Ulog("buildRule: Error from getAR(%d):  %s\n", id, err.Error())
		return "", err
	}

	d, err := GetLedger(ctx, rule.DebitLID)
	if err != nil {
		return "", err
	}

	c, err := GetLedger(ctx, rule.CreditLID)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("d %s _, c %s _", d.GLNumber, c.GLNumber)
	return s, err
}

// GetAssessmentAccountRule looks at the supplied Assessment.  If the .AcctRule is present
// then it is returned. If it is not present, then the ARID is used and an AcctRule is built
// from the ARID.
func GetAssessmentAccountRule(ctx context.Context, a *Assessment) (string, error) {
	if len(a.AcctRule) > 0 {
		return a.AcctRule, nil
	}
	return buildRule(ctx, a.ARID)
}

// GetReceiptAccountRule looks at the supplied Receipt.  If the .AcctRule is present
// then it is returned. If it is not present, then the ARID is used and an AcctRule is built
// from the ARID.
func GetReceiptAccountRule(ctx context.Context, a *Receipt) (string, error) {
	if len(a.AcctRuleApply) > 0 {
		return a.AcctRuleApply, nil
	}
	return buildRule(ctx, a.ARID)
}

func getRuleText(ctx context.Context, id int64) (string, error) {
	rule, err := GetAR(ctx, id)
	if err != nil {
		Ulog("getRuleText: Error from getAR(%d):  %s\n", id, err.Error())
		return "", err
	}
	return rule.Name, err
}

// GetAssessmentAccountRuleText returns the text to use in reports or in a UI that describes
// the assessment
func GetAssessmentAccountRuleText(ctx context.Context, a *Assessment) (string, error) {
	if len(a.AcctRule) > 0 {
		return a.AcctRule, nil
	}
	return getRuleText(ctx, a.ARID)
}

// GetReceiptAccountRuleText returns the text to use in reports or in a UI that describes
// the Receipt
func GetReceiptAccountRuleText(ctx context.Context, a *Receipt) (string, error) {
	if len(a.AcctRuleApply) > 0 {
		return a.AcctRuleApply, nil
	}
	return getRuleText(ctx, a.ARID)
}

// ProrateAssessment - determines the proration factor for this assessment
//
// Parameters:
//		a		 pointer to the assessment
//      d        date or the recurrence date of the assessment being analyzed
//  	d1, d2:  the time period we're being asked to analyze
//
// Returns:
//         	pf:  prorate factor = rentDur/asmtDur
//		   num:	 pf numerator, amount of rentcycle actually used expressed in units of prorateCycle
//         den:  pf denominator, the rent cycle, expressed in units of prorateCycle
//       start:	 trimmed start date (latest of RentalAgreement.PossessionStart and d1)
//        stop:	 trmmed stop date (soonest of RentalAgreement.PossessionStop and d2)
//=================================================================================================
func ProrateAssessment(ctx context.Context, xbiz *XBusiness, a *Assessment, d, d1, d2 *time.Time) (float64, int64, int64, time.Time, time.Time, error) {
	const funcname = "ProrateAssessment"

	var (
		pf          float64
		num, den    int64
		start, stop time.Time
		r           Rentable
		err         error
	)

	useStatus := int64(USESTATUSinService) // if RID==0, then it's for an application fee or similar.  Assume rentable is online.

	if a.RID > 0 {
		r, err = GetRentable(ctx, a.RID)
		if err != nil {
			return pf, num, den, start, stop, err
		}

		useStatus, err = GetRentableStateForDate(ctx, r.RID, d)
		if err != nil {
			return pf, num, den, start, stop, err
		}
	}

	// Console("GetRentableStateForDate( %d, %s ) = %d\n", r.RID, d.Format(RRDATEINPFMT), useStatus)
	switch useStatus {
	case USESTATUSemployee:
		fallthrough
	case USESTATUSinService:
		// Console("%s: at case USESTATUSinService.\n", funcname)
		ra, err := GetRentalAgreement(ctx, a.RAID)
		if err != nil {
			Ulog("ProrateAssessment: error getting rental agreement RAID=%d, err = %s\n", a.RAID, err.Error())
		} else {
			switch a.RentCycle {
			case CYCLEDAILY:
				// Console("%s: CYCLEDAILY: ra.RAID = %d, ra.RentStart = %s, ra.RentStop = %s\n", funcname, ra.RAID, ra.RentStart.Format(RRDATEFMT4), ra.RentStop.Format(RRDATEFMT4))
				pf, num, den, start, stop = CalcProrationInfo(&ra.RentStart, &ra.RentStop, d, d, a.RentCycle, a.ProrationCycle)
			case CYCLENORECUR:
				fallthrough
			case CYCLEMONTHLY:
				// Console("%s: CYCLEMONTHLY: ra.RAID = %d, ra.RentStart = %s, ra.RentStop = %s\n", funcname, ra.RAID, ra.RentStart.Format(RRDATEFMT4), ra.RentStop.Format(RRDATEFMT4))
				pf, num, den, start, stop = CalcProrationInfo(&ra.RentStart, &ra.RentStop, d1, d2, a.RentCycle, a.ProrationCycle)
			default:
				LogAndPrint("Accrual rate %d not implemented\n", a.RentCycle)
			}
		}
		// Console("Assessment = %d, Rentable = %d, RA = %d, pf = %3.2f\n", a.ASMID, r.RID, ra.RAID, pf)

	case USESTATUSadmin:
		fallthrough
	case USESTATUSownerOccupied:
		fallthrough
	case USESTATUSofflineRenovation:
		fallthrough
	case USESTATUSofflineMaintenance:
		ta, err := GetAllRentableAssessments(ctx, r.RID, d1, d2)
		if err != nil {
			return pf, num, den, start, stop, nil
		}

		if len(ta) > 0 {
			rentcycle, proration, _, err := GetRentCycleAndProration(ctx, &r, d1, xbiz)
			if err != nil {
				// TODO(Steve): dont we return error?
				Ulog("%s: error getting rent cycle for rentable %d. err = %s\n", funcname, r.RID, err.Error())
			}

			pf, num, den, start, stop = CalcProrationInfo(&(ta[0].Start), &(ta[0].Stop), d1, d2, rentcycle, proration)
			if len(ta) > 1 {
				Ulog("%s: %d Assessments affect Rentable %d (%s) with OFFLINE useStatus during %s - %s\n",
					funcname, len(ta), r.RID, r.RentableName, d1.Format(RRDATEINPFMT), d2.Format(RRDATEINPFMT))
			}
		}
	default:
		Ulog("%s: Rentable %d on %s has unknown useStatus: %d\n", funcname, r.RID, d.Format(RRDATEINPFMT), useStatus)
	}

	return pf, num, den, start, stop, err
}

// journalAssessment processes the assessment, creates a Journal entry, and returns its id
// Parameters:
//		xbiz - the business struct
//		rid - Rentable ID
//		d - date of this assessment
//		a - the assessment
//		d1-d2 - defines the timerange being covered in this period
//=================================================================================================
func journalAssessment(ctx context.Context, xbiz *XBusiness, d time.Time, a *Assessment, d1, d2 *time.Time) (Journal, error) {
	// funcname := "journalAssessment"
	// Console("Entered %s\n", funcname)
	// Console("%s: d = %s, d1 = %s, d2 = %s\n", funcname, d.Format(RRDATEREPORTFMT), d1.Format(RRDATEREPORTFMT), d2.Format(RRDATEREPORTFMT))
	var j Journal

	pf, num, den, start, stop, err := ProrateAssessment(ctx, xbiz, a, &d, d1, d2)
	if err != nil {
		return j, err
	}

	//--------------------------------------------------------------------------------------
	// This is a safeguard against issues encountered in Feb 2018 where rent assessments
	// continued after the RentalAgreement RentStop date.
	//--------------------------------------------------------------------------------------
	if pf < float64(0) {
		pf = float64(0)
	}

	// Console("%s: a.ASMTID = %d, d = %s, d1 = %s, d2 = %s\n", funcname, a.ASMID, d.Format(RRDATEFMT4), d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4))
	// Console("%s: pf = %f, num = %d, den = %d, start = %s, stop = %s\n", funcname, pf, num, den, start.Format(RRDATEFMT4), stop.Format(RRDATEFMT4))

	j = Journal{BID: a.BID, Dt: d, Type: JNLTYPEASMT, ID: a.ASMID}

	asmRules, err := GetAssessmentAccountRule(ctx, a)
	if err != nil {
		return j, err
	}

	m, err := ParseAcctRule(ctx, xbiz, a.RID, d1, d2, asmRules, a.Amount, pf) // a rule such as "d 11001 1000.0, c 40001 1100.0, d 41004 100.00"
	if err != nil {
		return j, err
	}

	// Console("%s:  m = %#v\n", funcname, m)
	// for i := 0; i < len(m); i++ {
	// 	Console("m[%d].Amount = %f,  .Action = %s   .Expr = %s\n", i, m[i].Amount, m[i].Action, m[i].Expr)
	// }

	_, j.Amount = sumAllocations(&m)
	j.Amount = RoundToCent(j.Amount)

	// Console("j.Amount = %f\n", j.Amount)

	//------------------------------------------------------------------------------------------------------
	// for non-recurring assessments (the only kind that we should be processing here) the amount may have
	// been prorated as it was a newly created recurring assessment for a RentalAgreement that was either
	// just beginning or just ending. If so, we'll update the assessment amount here the calculated
	// j.Amount != a.Amount
	//------------------------------------------------------------------------------------------------------
	if pf < 1.0 {
		a.Amount = j.Amount // update to the prorated amount
		a.Start = start     // adjust to the dates used in the proration
		a.Stop = stop       // adjust to the dates used in the proration
		a.Comment = fmt.Sprintf("Prorated: %d %s out of %d", num, ProrationUnits(a.ProrationCycle), den)
		if err := UpdateAssessment(ctx, a); err != nil {
			err = fmt.Errorf("Error updating prorated assessment amount: %s", err.Error())
			return j, err
		}
	}

	//-------------------------------------------------------------------------------------------
	// In the event that we need to prorate, pull together the pieces and determine the
	// fractional amounts so that all the entries can net to 0.00.  Essentially, this means
	// handling the $0.01 off problem when dealing with fractional numbers.  The way we'll
	// handle this is to apply the extra cent to the largest number
	//-------------------------------------------------------------------------------------------
	if pf < 1.0 {
		// new method using ProcessSum
		var asum []SumFloat
		for i := 0; i < len(m); i++ {
			var b SumFloat
			if m[i].Action == "c" {
				b.Val = -m[i].Amount
			} else {
				b.Val = m[i].Amount
			}
			b.Amount = RoundToCent(b.Val)
			b.Remainder = b.Amount - b.Val
			asum = append(asum, b)
		}
		ProcessSumFloats(asum)
		for i := 0; i < len(asum); i++ {
			if m[i].Action == "c" {
				m[i].Amount = -asum[i].Amount // the adjusted value after ProcessSumFloats
			} else {
				m[i].Amount = asum[i].Amount // the adjusted value after ProcessSumFloats
			}
		}

	}

	// Console("INSERTING JOURNAL: Date = %s, Type = %d, amount = %f\n", j.Dt, j.Type, j.Amount)

	jid, err := InsertJournal(ctx, &j)
	if err != nil {
		LogAndPrintError("error inserting Journal entry: %v\n", err)
		return j, err
	}

	s := ""
	for i := 0; i < len(m); i++ {
		s += fmt.Sprintf("%s %s %.2f", m[i].Action, m[i].AcctExpr, RoundToCent(m[i].Amount))
		if i+1 < len(m) {
			s += ", "
		}
	}
	if jid > 0 {
		var ja JournalAllocation
		ja.JID = jid
		ja.RID = a.RID
		ja.ASMID = a.ASMID
		ja.Amount = RoundToCent(j.Amount)
		ja.AcctRule = s
		ja.BID = a.BID
		ja.RAID = a.RAID

		// Console("INSERTING JOURNAL-ALLOCATION: ja.JID = %d, ja.ASMID = %d, ja.RAID = %d\n", ja.JID, ja.ASMID, ja.RAID)
		if _, err = InsertJournalAllocationEntry(ctx, &ja); err != nil {
			LogAndPrintError("journalAssessment", err)
			return j, err
		}
		j.JA = append(j.JA, ja)
	}

	return j, err
}

// RemoveJournalEntries clears out the records in the supplied range provided the range is not closed by a JournalMarker
//=================================================================================================
func RemoveJournalEntries(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time) error {
	// Remove the Journal entries and the JournalAllocation entries
	rows, err := RRdb.Prepstmt.GetAllJournalsInRange.Query(xbiz.P.BID, d1, d2)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var j Journal
		ReadJournals(rows, &j)
		DeleteJournalAllocations(ctx, j.JID)
		DeleteJournal(ctx, j.JID)
	}

	// only delete the marker if it is in this time range and if it is not the origin marker
	jm, err := GetLastJournalMarker(ctx)
	if err != nil {
		return err
	}

	if jm.State == LMOPEN && (jm.DtStart.After(*d1) || jm.DtStart.Equal(*d1)) && (jm.DtStop.Before(*d2) || jm.DtStop.Equal(*d2)) {
		DeleteJournalMarker(ctx, jm.JMID)
	}

	RemoveLedgerEntries(ctx, xbiz, d1, d2)
	return err
}

// ProcessNewAssessmentInstance creates a Journal entry for the supplied non-recurring assessment
//=================================================================================================
func ProcessNewAssessmentInstance(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time, a *Assessment) (Journal, error) {
	funcname := "ProcessNewAssessmentInstance"
	var j Journal
	var err error
	// Console("Entered %s:  d1 = %s, d2 = %s, Assessment date: %s\n", funcname, d1.Format(RRDATEREPORTFMT), d2.Format(RRDATEREPORTFMT), a.Start.Format(RRDATEREPORTFMT))
	if a.PASMID == 0 && a.RentCycle != RECURNONE { // if this assessment is not a single instance recurrence, then return an error
		err = fmt.Errorf("%s: Function only accepts non-recurring instances, RentCycle = %d", funcname, a.RentCycle)
		LogAndPrintError(funcname, err)
		return j, err
	}
	if a.ASMID == 0 && a.RentCycle != RECURNONE {
		_, err = InsertAssessment(ctx, a)
		if nil != err {
			LogAndPrintError(funcname, err)
			return j, err
		}
	}

	// Console("%s: Calling journalAssessment for ASMID = %d\n", funcname, a.ASMID)
	j, err = journalAssessment(ctx, xbiz, a.Start, a, d1, d2)
	return j, err
}

// ProcessNewReceipt creates a Journal entry for the supplied receipt
//-----------------------------------------------------------------------------
func ProcessNewReceipt(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time, r *Receipt) (Journal, error) {
	var j Journal
	j.BID = xbiz.P.BID
	j.Amount = RoundToCent(r.Amount)
	j.Dt = r.Dt
	j.Type = JNLTYPERCPT
	j.ID = r.RCPTID
	// j.RAID = r.RAID
	jid, err := InsertJournal(ctx, &j)
	if err != nil {
		Ulog("Error inserting Journal entry: %v\n", err)
		return j, err
	}
	if jid > 0 {
		// now add the Journal allocation records...
		for i := 0; i < len(r.RA); i++ {
			// Console("r.RA[%d] id = %d\n", i, r.RA[i].RCPAID)
			// rntagr, _ := GetRentalAgreement(r.RA[i].RAID) // what Rental Agreements did this payment affect and the amounts for each
			var ja JournalAllocation
			ja.JID = jid
			ja.TCID = r.TCID
			ja.Amount = RoundToCent(r.RA[i].Amount)
			ja.BID = j.BID
			ja.ASMID = r.RA[i].ASMID
			ja.AcctRule = r.RA[i].AcctRule
			if ja.ASMID > 0 { // there may not be an assessment associated, it could be unallocated funds
				// TODO(Steve): should we ignore error?
				a, _ := GetAssessment(ctx, ja.ASMID) // but if there is an associated assessment, then mark the RID and RAID
				ja.RID = a.RID
				ja.RAID = r.RA[i].RAID
			}
			ja.TCID = r.TCID
			if _, err = InsertJournalAllocationEntry(ctx, &ja); err != nil {
				LogAndPrintError("ProcessNewReceipt", err)
				return j, err
			}
			j.JA = append(j.JA, ja)
		}
	}
	return j, nil
}

// ProcessNewExpense adds a new expense instance.
//-----------------------------------------------------------------------------
func ProcessNewExpense(ctx context.Context, a *Expense, xbiz *XBusiness) error {
	InitBizInternals(a.BID, xbiz)
	var j = Journal{
		BID:    xbiz.P.BID,
		Amount: a.Amount,
		Dt:     a.Dt,
		Type:   JNLTYPEEXP,
		ID:     a.EXPID,
	}
	_, err := InsertJournal(ctx, &j)
	if err != nil {
		Ulog("Error inserting Journal Expense entry: %v\n", err)
		return err
	}
	var ja = JournalAllocation{
		JID:    j.JID,
		BID:    j.BID,
		RID:    a.RID,
		RAID:   a.RAID,
		Amount: a.Amount,
		EXPID:  a.EXPID,
	}
	clid := RRdb.BizTypes[a.BID].AR[a.ARID].CreditLID
	dlid := RRdb.BizTypes[a.BID].AR[a.ARID].DebitLID
	ja.AcctRule = fmt.Sprintf("d %s %.2f, c %s %.2f",
		RRdb.BizTypes[a.BID].GLAccounts[dlid].GLNumber, a.Amount,
		RRdb.BizTypes[a.BID].GLAccounts[clid].GLNumber, a.Amount)
	if _, err = InsertJournalAllocationEntry(ctx, &ja); err != nil {
		LogAndPrintError("ProcessNewReceipt", err)
		return err
	}
	j.JA = append(j.JA, ja)
	d1 := time.Date(a.Dt.Year(), a.Dt.Month(), 1, 0, 0, 0, 0, time.UTC)
	d2 := d1.AddDate(0, 1, 0)
	InitLedgerCache()
	GenerateLedgerEntriesFromJournal(ctx, xbiz, &j, &d1, &d2)
	return nil
}

// ProcessJournalEntry processes an assessment. It adds instances of recurring
// assessments for the time period d1-d2 if they do not already exist. Then
// creates a journal entry for the assessment.
//-----------------------------------------------------------------------------
func ProcessJournalEntry(ctx context.Context, a *Assessment, xbiz *XBusiness, d1, d2 *time.Time, updateLedgers bool) error {
	funcname := "ProcessJournalEntry"
	var j Journal
	var err error
	// Console("ProcessJournalEntry: 1. a.ASMID = %d, Biz = %s (%d), d1 - d2 = %s - %s\n",
	// a.ASMID, xbiz.P.Designation, xbiz.P.BID, d1.Format(RRDATEREPORTFMT), d2.Format(RRDATEREPORTFMT))
	if a.RentCycle == RECURNONE {
		j, err = ProcessNewAssessmentInstance(ctx, xbiz, d1, d2, a)
		if err != nil {
			LogAndPrintError(funcname, err)
			return err
		}

		if updateLedgers {
			_, err = GenerateLedgerEntriesFromJournal(ctx, xbiz, &j, d1, d2)
			if err != nil {
				LogAndPrintError(funcname, err)
				return err
			}
		}
	} else if a.RentCycle >= RECURSECONDLY && a.RentCycle <= RECURHOURLY {
		// TBD
		LogAndPrint("Unhandled assessment recurrence type: %d\n", a.RentCycle)
	} else {
		// Console("ProcessJournalEntry: 2\n")
		// Console("Instances that occur between %s and %s for assessment dates(%s-%s)\n", d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4), a.Start.Format(RRDATEFMT4), a.Stop.Format(RRDATEFMT4))
		dl := a.GetRecurrences(d1, d2)
		rangeDuration := d2.Sub(*d1)
		// Console("ProcessJournalEntry: 3... len(dl) = %d\n", len(dl))
		for i := 0; i < len(dl); i++ {
			a1 := *a
			// Console("ProcessJournalEntry: 3.1  a1.Amount = %.2f\n", a1.Amount)
			a1.FLAGS = 0                                                  // ensure that we don't cary forward any flags
			a1.Start = dl[i]                                              // use the instance date
			a1.Stop = dl[i].Add(CycleDuration(a.ProrationCycle, a.Start)) // add enough time so that the recurrence calculator sees this instance
			a1.ASMID = 0                                                  // ensure this is a new assessment
			a1.PASMID = a.ASMID                                           // parent assessment
			// Console("****>>>>>>  a1.Start = %s\n", a1.Start.Format(RRDATEFMT4))
			// Console("****>>>>>>  a1.Stop  = %s\n", a1.Stop.Format(RRDATEFMT4))
			// Console("****>>>>>>  CycleDuration( %d, %s ) --->  %d\n", a.ProrationCycle, a.Start.Format(RRDATEFMT4), CycleDuration(a.ProrationCycle, a.Start))

			//--------------------------------------------------------------------------------
			// Before inserting this, validate that the RentalAgreement for this assessment
			// is still in effect.  This is because in the early versions of the Roller
			// server code, there were no checks to ensure that recurring assessments stopped
			// when their associated RentalAgreements stopped.
			//--------------------------------------------------------------------------------
			if a.RAID > 0 {
				// Console("a.RAID = %d\n", a.RAID)
				ra, err := GetRentalAgreement(ctx, a.RAID)
				if err != nil {
					LogAndPrintError(funcname, err)
					return err
				}
				// Console("ra.RentStop = %s\n", ra.RentStop)
				if a1.Start.After(ra.RentStop) || a1.Start.Equal(ra.RentStop) {
					// Console("Do not add the new assessment\n")
					err = fmt.Errorf("%s:  Cannot add new assessment instance on %s after RentalAgreement (%s) stop date %s", funcname, a1.Start.Format(RRDATEREPORTFMT), ra.IDtoShortString(), ra.RentStop.Format(RRDATEREPORTFMT))
					LogAndPrintError(funcname, err)
					return err
				}
			}

			//--------------------------------------------------------------------------------
			// The generation of recurring assessment instances needs to be idempotent.
			// Check to ensure that this instance does not already exist before generating it
			//--------------------------------------------------------------------------------
			a2, err := GetAssessmentInstance(ctx, &a1.Start, a1.PASMID) // if this returns an existing instance (ASMID != 0) then it's already been processed...
			if err != nil {
				// Console("Error in GetAssessmentInstance: %s\n", err.Error())
			}
			if a2.ASMID == 0 { // ... otherwise, process this instance
				// Console("ProcessJournalEntry: 4.0, a1.Amount = %.2f\n", a1.Amount)
				_, err = InsertAssessment(ctx, &a1)
				Errlog(err)
				// Console("ProcessJournalEntry: 4.1, inserted a1.ASMID = %d, a1.Amount = %.2f\n", a1.ASMID, a1.Amount)

				//--------------------------------------------------------------------------------
				// Rent is assessed on the following cycle: a.RentCycle
				// and prorated on the following cycle: a.ProrationCycle
				//--------------------------------------------------------------------------------
				rentCycleDur := CycleDuration(a.RentCycle, dl[i])
				diff := rangeDuration - rentCycleDur
				if diff < 0 {
					diff = -diff
				}
				dtb := *d1
				dte := *d2
				// Console("dtb = %s, dte = %s, diff = %d\n", dtb.Format(RRDATEFMT4), dte.Format(RRDATEFMT4), diff)
				if diff > rentCycleDur/9 { // if this is true then
					dtb = dl[i] // add one full cycle diration
					dte = dtb.Add(CycleDuration(a.RentCycle, dtb))
				}
				j, err := ProcessNewAssessmentInstance(ctx, xbiz, &dtb, &dte, &a1)
				if err != nil {
					LogAndPrintError(funcname, err)
					return err
				}
				if updateLedgers {
					_, err = GenerateLedgerEntriesFromJournal(ctx, xbiz, &j, d1, d2)
					if err != nil {
						LogAndPrintError(funcname, err)
						return err
					}
				}
			} else if a.RentCycle >= RECURSECONDLY && a.RentCycle <= RECURHOURLY {
				LogAndPrintError(funcname, fmt.Errorf("Unhandled RentCycle frequency: %d", a.RentCycle))
			}
			// Console("ProcessJournalEntry: 5\n")
		}
	}
	return err
}

// GenerateRecurInstances creates Assessment instance records for recurring Assessments and then
// creates the corresponding journal instances for the new assessment instances
//=================================================================================================
func GenerateRecurInstances(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time) error {
	// Console("GetRecurringAssessmentsByBusiness - d1 = %s   d2 = %s\n", d1.Format(RRDATEINPFMT, d2.Format(RRDATEINPFMT)))
	rows, err := RRdb.Prepstmt.GetRecurringAssessmentsByBusiness.Query(xbiz.P.BID, d2, d1) // only get recurring instances without a parent
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var a Assessment
		err = ReadAssessments(rows, &a)
		if err != nil {
			return err
		}
		// Console("rlib.GenerateRecurInstances: Process ASMID = %d\n", a.ASMID)

		err = ProcessJournalEntry(ctx, &a, xbiz, d1, d2, false)
		if err != nil {
			return err
		}
	}

	return rows.Err()
}

// ProcessReceiptRange creates Journal records for Receipts in the supplied date range
//=================================================================================================
func ProcessReceiptRange(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time) error {
	r, err := GetReceipts(ctx, xbiz.P.BID, d1, d2)
	if err != nil {
		return err
	}

	for i := 0; i < len(r); i++ {
		j, err := GetJournalByReceiptID(ctx, r[i].RCPTID)
		if err != nil || j.JID == 0 { // TODO(Steve): are we sure that we want to proceed if err != nil?
			// if you want log the error then separate this condition in two if clauses
			_, err = ProcessNewReceipt(ctx, xbiz, d1, d2, &r[i])
			if err != nil {
				return err
			}
		}
	}

	return err
}

// CreateJournalMarker creates a Journal Marker record for the supplied date range
//=================================================================================================
func CreateJournalMarker(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time) error {
	const funcname = "CreateJournalMarker"
	var jm JournalMarker
	jm.BID = xbiz.P.BID
	jm.State = LMOPEN
	jm.DtStart = *d1
	jm.DtStop = *d2

	_, err := InsertJournalMarker(ctx, &jm)
	if err != nil {
		Ulog("%s: Error while inserting journal marker: %s\n", funcname, err.Error())
	}
	return err
}

// GenerateJournalRecords creates Journal records for Assessments and receipts over the supplied time range.
//=================================================================================================
func GenerateJournalRecords(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time, skipVac bool) error {
	// err := RemoveJournalEntries(xbiz, d1, d2)
	// if err != nil {
	// 	Ulog("Could not remove existing Journal entries from %s to %s. err = %v\n", d1.Format(RRDATEFMT), d2.Format(RRDATEFMT), err)
	// 	return
	// }
	var (
		err error
	)

	err = GenerateRecurInstances(ctx, xbiz, d1, d2)
	if err != nil {
		return err
	}
	if !skipVac {
		_, err = GenVacancyJournals(ctx, xbiz, d1, d2)
		if err != nil {
			return err
		}
	}
	err = ProcessReceiptRange(ctx, xbiz, d1, d2)
	if err != nil {
		return err
	}
	return CreateJournalMarker(ctx, xbiz, d1, d2)
	// return err
}
