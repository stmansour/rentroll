package rlib

import (
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

// ProrateAssessment - determines the proration factor for this assessment
//
// Parameters:
//		a			pointer to the assessment
//      d           date or the recurrence date of the assessment being analyzed
//  	d1, d2:     the time period we're being asked to analyze
//
// Returns:
//         	pf:     prorate factor = rentDur/asmtDur
//		   num:		pf numerator, amount of rentcycle actually used expressed in units of prorateCycle
//         den:     pf denominator, the rent cycle, expressed in units of prorateCycle
//       start:		trimmed start date (latest of RentalAgreement.PossessionStart and d1)
//        stop:		trmmed stop date (soonest of RentalAgreement.PossessionStop and d2)
//=================================================================================================
func ProrateAssessment(xbiz *XBusiness, a *Assessment, d, d1, d2 *time.Time) (float64, int64, int64, time.Time, time.Time) {
	funcname := "ProrateAssessment"
	pf := float64(0)
	var num, den int64
	var start, stop time.Time
	r := GetRentable(a.RID)
	status := GetRentableStateForDate(r.RID, d)
	switch status {
	case RENTABLESTATUSONLINE:
		ra, _ := GetRentalAgreement(a.RAID)
		switch a.RentCycle {
		case CYCLEDAILY:
			pf, num, den, start, stop = CalcProrationInfo(&ra.RentStart, &ra.RentStop, d, d, a.RentCycle, a.ProrationCycle)
		case CYCLENORECUR:
			fallthrough
		case CYCLEMONTHLY:
			pf, num, den, start, stop = CalcProrationInfo(&ra.RentStart, &ra.RentStop, d1, d2, a.RentCycle, a.ProrationCycle)
		default:
			fmt.Printf("Accrual rate %d not implemented\n", a.RentCycle)
		}
		// fmt.Printf("Assessment = %d, Rentable = %d, RA = %d, pf = %3.2f\n", a.ASMID, r.RID, ra.RAID, pf)

	case RENTABLESTATUSADMIN:
		fallthrough
	case RENTABLESTATUSEMPLOYEE:
		fallthrough
	case RENTABLESTATUSOWNEROCC:
		fallthrough
	case RENTABLESTATUSOFFLINE:
		ta := GetAllRentableAssessments(r.RID, d1, d2)
		if len(ta) > 0 {
			rentcycle, proration, _, err := GetRentCycleAndProration(&r, d1, xbiz)
			if err != nil {
				Ulog("%s: error getting rent cycle for rentable %d. err = %s\n", funcname, r.RID, err.Error())
			}
			pf, num, den, start, stop = CalcProrationInfo(&(ta[0].Start), &(ta[0].Stop), d1, d2, rentcycle, proration)
			if len(ta) > 1 {
				Ulog("%s: %d Assessments affect Rentable %d (%s) in period %s - %s\n",
					funcname, len(ta), r.RID, r.RentableName, d1.Format(RRDATEINPFMT), d2.Format(RRDATEINPFMT))
			}
		}
	default:
		Ulog("%s: Rentable %d is in an unknown status: %d\n", funcname, r.RID, status)
	}

	return pf, num, den, start, stop
}

// journalAssessment processes the assessment, creates a Journal entry, and returns its id
// Parameters:
//		xbiz - the business struct
//		rid - Rentable ID
//		d - date of this assessment
//		a - the assessment
//		d1-d2 - defines the timerange being covered in this period
//=================================================================================================
func journalAssessment(xbiz *XBusiness, d time.Time, a *Assessment, d1, d2 *time.Time) error {
	// funcname := "journalAssessment"
	pf, num, den, start, stop := ProrateAssessment(xbiz, a, &d, d1, d2)

	// fmt.Printf("ProrateAssessment: a.ASMTID = %d, d = %s, d1 = %s, d2 = %s\n", a.ASMID, d.Format(RRDATEFMT4), d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4))
	// fmt.Printf("pf = %f, num = %d, den = %d, start = %s, stop = %s\n", pf, num, den, start.Format(RRDATEFMT4), stop.Format(RRDATEFMT4))

	var j = Journal{BID: a.BID, Dt: d, Type: JNLTYPEASMT, ID: a.ASMID, RAID: a.RAID}

	// fmt.Printf("calling ParseAcctRule:\n  asmt = %d\n  rid = %d, d1 = %s, d2 = %s\n  a.Amount = %f\n", a.ASMID, a.RID, d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4), a.Amount)
	// fmt.Printf("RRdb.BizTypes[bid].DefaultAccts = %#v\n", RRdb.BizTypes[a.BID].DefaultAccts)

	m := ParseAcctRule(xbiz, a.RID, d1, d2, a.AcctRule, a.Amount, pf) // a rule such as "d 11001 1000.0, c 40001 1100.0, d 41004 100.00"

	// fmt.Printf("journalAssessment:  m = %#v\n", m)
	// for i := 0; i < len(m); i++ {
	// 	fmt.Printf("m[%d].Amount = %f,  .Action = %s   .Expr = %s\n", i, m[i].Amount, m[i].Action, m[i].Expr)
	// }

	_, j.Amount = sumAllocations(&m)
	j.Amount = RoundToCent(j.Amount)

	// fmt.Printf("j.Amount = %f\n", j.Amount)

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
		if err := UpdateAssessment(a); err != nil {
			err = fmt.Errorf("Error updating prorated assessment amount: %s", err.Error())
			return err
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

	// fmt.Printf("INSERTING JOURNAL: Date = %s, Type = %d, amount = %f\n", j.Dt, j.Type, j.Amount)

	jid, err := InsertJournalEntry(&j)
	if err != nil {
		Ulog("error inserting Journal entry: %v\n", err)
	} else {
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
			InsertJournalAllocationEntry(&ja)
		}
	}

	return err
}

// RemoveJournalEntries clears out the records in the supplied range provided the range is not closed by a JournalMarker
//=================================================================================================
func RemoveJournalEntries(xbiz *XBusiness, d1, d2 *time.Time) error {
	// Remove the Journal entries and the JournalAllocation entries
	rows, err := RRdb.Prepstmt.GetAllJournalsInRange.Query(xbiz.P.BID, d1, d2)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var j Journal
		Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID, &j.Comment, &j.LastModTime, &j.LastModBy))
		DeleteJournalAllocations(j.JID)
		DeleteJournalEntry(j.JID)
	}

	// only delete the marker if it is in this time range and if it is not the origin marker
	jm := GetLastJournalMarker()
	if jm.State == MARKERSTATEOPEN && (jm.DtStart.After(*d1) || jm.DtStart.Equal(*d1)) && (jm.DtStop.Before(*d2) || jm.DtStop.Equal(*d2)) {
		DeleteJournalMarker(jm.JMID)
	}

	RemoveLedgerEntries(xbiz, d1, d2)
	return err
}

// ProcessNewAssessmentInstance creates a Journal entry for the supplied non-recurring assessment
//=================================================================================================
func ProcessNewAssessmentInstance(xbiz *XBusiness, d1, d2 *time.Time, a *Assessment) error {
	funcname := "ProcessNewAssessmentInstance"
	var noerr error
	if a.PASMID == 0 && a.RentCycle != RECURNONE { // if this assessment is not a single instance recurrence, then return an error
		return fmt.Errorf("%s: Function only accepts non-recurring instances", funcname)
	}
	// if a.ASMID != 0 && a.RentCycle { // if this assessment has already been written to db, then return error
	// 	return fmt.Errorf("%s: Function only accepts unsaved instances. Found ASMID = %d", funcname, a.ASMID)
	// }
	if a.ASMID == 0 && a.RentCycle != RECURNONE {
		ASMID, err := InsertAssessment(a)
		if nil != err {
			return err
		}
		a.ASMID = ASMID
	}

	// fmt.Println("c1")

	journalAssessment(xbiz, a.Start, a, d1, d2)

	// fmt.Println("d1")

	return noerr
}

// ProcessNewReceipt creates a Journal entry for the supplied receipt
//=================================================================================================
func ProcessNewReceipt(xbiz *XBusiness, d1, d2 *time.Time, r *Receipt) error {
	rntagr, _ := GetRentalAgreement(r.RAID)
	var j Journal
	j.BID = rntagr.BID
	j.Amount = RoundToCent(r.Amount)
	j.Dt = r.Dt
	j.Type = JNLTYPERCPT
	j.ID = r.RCPTID
	j.RAID = r.RAID
	jid, err := InsertJournalEntry(&j)
	if err != nil {
		Ulog("Error inserting Journal entry: %v\n", err)
		return err
	}
	if jid > 0 {
		// now add the Journal allocation records...
		for j := 0; j < len(r.RA); j++ {
			var ja JournalAllocation
			ja.JID = jid
			ja.Amount = RoundToCent(r.RA[j].Amount)
			ja.ASMID = r.RA[j].ASMID
			ja.AcctRule = r.RA[j].AcctRule
			a, _ := GetAssessment(ja.ASMID)
			ja.RID = a.RID
			ja.BID = a.BID
			InsertJournalAllocationEntry(&ja)
		}
	}
	return err
}

// ProcessJournalEntry processes an assessment. It adds instances of recurring assessments for
// the time period d1-d2 if they do not already exist. Then creates a journal entry for the assessment.
func ProcessJournalEntry(a *Assessment, xbiz *XBusiness, d1, d2 *time.Time) {
	// fmt.Printf("ProcessJournalEntry: 1.   a = %#v\n", a)
	if a.RentCycle == RECURNONE {
		ProcessNewAssessmentInstance(xbiz, d1, d2, a)
	} else if a.RentCycle >= RECURSECONDLY && a.RentCycle <= RECURHOURLY {
		// TBD
		fmt.Printf("Unhandled assessment recurrence type: %d\n", a.RentCycle)
	} else {
		// fmt.Printf("ProcessJournalEntry: 2\n")
		dl := a.GetRecurrences(d1, d2)
		rangeDuration := d2.Sub(*d1)
		// fmt.Printf("ProcessJournalEntry: 3... len(dl) = %d\n", len(dl))
		for i := 0; i < len(dl); i++ {
			a1 := *a
			a1.Start = dl[i]                                              // use the instance date
			a1.Stop = dl[i].Add(CycleDuration(a.ProrationCycle, a.Start)) // add enough time so that the recurrence calculator sees this instance
			a1.ASMID = 0                                                  // ensure this is a new assessment
			a1.PASMID = a.ASMID                                           // parent assessment

			// The generation of recurring assessment instances needs to be idempotent.
			// Check to ensure that this instance does not already exist before generating it
			a2, _ := GetAssessmentInstance(&a1.Start, a1.PASMID) // if this returns an existing instance (ASMID != 0) then it's already been processed...
			if a2.ASMID == 0 {                                   // ... otherwise, process this instance
				_, err := InsertAssessment(&a1)
				Errlog(err)
				// fmt.Printf("ProcessJournalEntry: 4, inserted a1.ASMID = %d\n", a1.ASMID)

				// Rent is assessed on the following cycle: a.RentCycle
				// and prorated on the following cycle: a.ProrationCycle
				rentCycleDur := CycleDuration(a.RentCycle, dl[i])
				diff := rangeDuration - rentCycleDur
				if diff < 0 {
					diff = -diff
				}
				dtb := *d1
				dte := *d2
				if diff > rentCycleDur/9 { // if this is true then
					dtb = dl[i] // add one full cycle diration
					dte = dtb.Add(CycleDuration(a.RentCycle, dtb))
				}
				ProcessNewAssessmentInstance(xbiz, &dtb, &dte, &a1)
			}
		}
	}
}

// GenerateRecurInstances creates Assessment instance records for recurring Assessments and then
// creates the corresponding journal instances for the new assessment instances
//=================================================================================================
func GenerateRecurInstances(xbiz *XBusiness, d1, d2 *time.Time) {
	// fmt.Printf("GetRecurringAssessmentsByBusiness - d2 = %s   d1 = %s\n", d2.Format(RRDATEINPFMT), d1.Format(RRDATEINPFMT))
	rows, err := RRdb.Prepstmt.GetRecurringAssessmentsByBusiness.Query(xbiz.P.BID, d2, d1) // only get recurring instances without a parent
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a Assessment
		ReadAssessments(rows, &a)
		ProcessJournalEntry(&a, xbiz, d1, d2)
	}
	Errcheck(rows.Err())
}

// ProcessReceiptRange creates Journal records for Receipts in the supplied date range
//=================================================================================================
func ProcessReceiptRange(xbiz *XBusiness, d1, d2 *time.Time) {
	r := GetReceipts(xbiz.P.BID, d1, d2)
	for i := 0; i < len(r); i++ {
		j := GetJournalByReceiptID(r[i].RCPTID)
		if j.JID == 0 {
			ProcessNewReceipt(xbiz, d1, d2, &r[i])
		}
	}
}

// CreateJournalMarker creates a Journal Marker record for the supplied date range
//=================================================================================================
func CreateJournalMarker(xbiz *XBusiness, d1, d2 *time.Time) {
	var jm JournalMarker
	jm.BID = xbiz.P.BID
	jm.State = MARKERSTATEOPEN
	jm.DtStart = *d1
	jm.DtStop = *d2
	InsertJournalMarker(&jm)
}

// GenerateJournalRecords creates Journal records for Assessments and receipts over the supplied time range.
//=================================================================================================
func GenerateJournalRecords(xbiz *XBusiness, d1, d2 *time.Time, skipVac bool) {
	// err := RemoveJournalEntries(xbiz, d1, d2)
	// if err != nil {
	// 	Ulog("Could not remove existing Journal entries from %s to %s. err = %v\n", d1.Format(RRDATEFMT), d2.Format(RRDATEFMT), err)
	// 	return
	// }
	GenerateRecurInstances(xbiz, d1, d2)
	if !skipVac {
		GenVacancyJournals(xbiz, d1, d2)
	}
	ProcessReceiptRange(xbiz, d1, d2)
	CreateJournalMarker(xbiz, d1, d2)
}
