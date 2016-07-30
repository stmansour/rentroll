package main

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

//=================================================================================================
func sumAllocations(m *[]rlib.AcctRule) (float64, float64) {
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
func ProrateAssessment(xbiz *rlib.XBusiness, a *rlib.Assessment, d, d1, d2 *time.Time) (float64, int64, int64, time.Time, time.Time) {
	funcname := "ProrateAssessment"
	pf := float64(0)
	var num, den int64
	var start, stop time.Time
	r := rlib.GetRentable(a.RID)
	status := rlib.GetRentableStateForDate(r.RID, d)
	switch status {
	case rlib.RENTABLESTATUSONLINE:
		ra, _ := rlib.GetRentalAgreement(a.RAID)
		switch a.RentCycle {
		case rlib.CYCLEDAILY:
			pf, num, den, start, stop = rlib.CalcProrationInfo(&ra.PossessionStart, &ra.PossessionStop, d, d, a.RentCycle, a.ProrationCycle)
		case rlib.CYCLENORECUR:
			fallthrough
		case rlib.CYCLEMONTHLY:
			pf, num, den, start, stop = rlib.CalcProrationInfo(&ra.PossessionStart, &ra.PossessionStop, d1, d2, a.RentCycle, a.ProrationCycle)
		default:
			fmt.Printf("Accrual rate %d not implemented\n", a.RentCycle)
		}
		// fmt.Printf("Assessment = %d, Rentable = %d, RA = %d, pf = %3.2f\n", a.ASMID, r.RID, ra.RAID, pf)

	case rlib.RENTABLESTATUSADMIN:
		fallthrough
	case rlib.RENTABLESTATUSEMPLOYEE:
		fallthrough
	case rlib.RENTABLESTATUSOWNEROCC:
		fallthrough
	case rlib.RENTABLESTATUSOFFLINE:
		ta := rlib.GetAllRentableAssessments(r.RID, d1, d2)
		if len(ta) > 0 {
			rentcycle, proration, _, err := rlib.GetRentCycleAndProration(&r, d1, xbiz)
			if err != nil {
				rlib.Ulog("%s: error getting rent cycle for rentable %d. err = %s\n", funcname, r.RID, err.Error())
			}
			pf, num, den, start, stop = rlib.CalcProrationInfo(&(ta[0].Start), &(ta[0].Stop), d1, d2, rentcycle, proration)
			if len(ta) > 1 {
				rlib.Ulog("%s: %d Assessments affect Rentable %d (%s) in period %s - %s\n",
					funcname, len(ta), r.RID, r.Name, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
			}
		}
	default:
		rlib.Ulog("%s: Rentable %d is in an unknown status: %d\n", funcname, r.RID, status)
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
func journalAssessment(xbiz *rlib.XBusiness, d time.Time, a *rlib.Assessment, d1, d2 *time.Time) error {
	// funcname := "journalAssessment"
	pf, num, den, start, stop := ProrateAssessment(xbiz, a, &d, d1, d2)
	var j = rlib.Journal{BID: a.BID, Dt: d, Type: rlib.JNLTYPEASMT, ID: a.ASMID, RAID: a.RAID}

	// fmt.Printf("calling ParseAcctRule:\n  asmt = %#v\n  rid = %d\n", a, rid)
	m := rlib.ParseAcctRule(xbiz, a.RID, d1, d2, a.AcctRule, a.Amount, pf) // a rule such as "d 11001 1000.0, c 40001 1100.0, d 41004 100.00"
	_, j.Amount = sumAllocations(&m)
	j.Amount = rlib.RoundToCent(j.Amount)

	// fmt.Printf("After ParseAcctRule - j.Amount = %8.2f\n", j.Amount)

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
		a.Comment = fmt.Sprintf("Prorated: %d %s out of %d", num, rlib.ProrationUnits(a.ProrationCycle), den)
		if err := rlib.UpdateAssessment(a); err != nil {
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
		var asum []rlib.SumFloat
		for i := 0; i < len(m); i++ {
			var b rlib.SumFloat
			if m[i].Action == "c" {
				b.Val = -m[i].Amount
			} else {
				b.Val = m[i].Amount
			}
			b.Amount = rlib.RoundToCent(b.Val)
			b.Remainder = b.Amount - b.Val
			asum = append(asum, b)
		}
		rlib.ProcessSumFloats(asum)
		for i := 0; i < len(asum); i++ {
			if m[i].Action == "c" {
				m[i].Amount = -asum[i].Amount // the adjusted value after ProcessSumFloats
			} else {
				m[i].Amount = asum[i].Amount // the adjusted value after ProcessSumFloats
			}
		}

	}

	// fmt.Printf("INSERTING JOURNAL: Date = %s, Type = %d, amount = %f\n", j.Dt, j.Type, j.Amount)
	jid, err := rlib.InsertJournalEntry(&j)
	if err != nil {
		rlib.Ulog("error inserting Journal entry: %v\n", err)
	} else {
		//now build up the AcctRule...
		s := ""
		for i := 0; i < len(m); i++ {
			s += fmt.Sprintf("%s %s %.2f", m[i].Action, m[i].Account, rlib.RoundToCent(m[i].Amount))
			if i+1 < len(m) {
				s += ", "
			}
		}
		if jid > 0 {
			var ja rlib.JournalAllocation
			ja.JID = jid
			ja.RID = a.RID
			ja.ASMID = a.ASMID
			ja.Amount = rlib.RoundToCent(j.Amount)
			ja.AcctRule = s
			rlib.InsertJournalAllocationEntry(&ja)
		}
	}

	return err
}

// RemoveJournalEntries clears out the records in the supplied range provided the range is not closed by a JournalMarker
//=================================================================================================
func RemoveJournalEntries(xbiz *rlib.XBusiness, d1, d2 *time.Time) error {
	// Remove the Journal entries and the JournalAllocation entries
	rows, err := rlib.RRdb.Prepstmt.GetAllJournalsInRange.Query(xbiz.P.BID, d1, d2)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var j rlib.Journal
		rlib.Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID, &j.Comment, &j.LastModTime, &j.LastModBy))
		rlib.DeleteJournalAllocations(j.JID)
		rlib.DeleteJournalEntry(j.JID)
	}

	// only delete the marker if it is in this time range and if it is not the origin marker
	jm := rlib.GetLastJournalMarker()
	if jm.State == rlib.MARKERSTATEOPEN && (jm.DtStart.After(*d1) || jm.DtStart.Equal(*d1)) && (jm.DtStop.Before(*d2) || jm.DtStop.Equal(*d2)) {
		rlib.DeleteJournalMarker(jm.JMID)
	}

	RemoveLedgerEntries(xbiz, d1, d2)
	return err
}

// ProcessNewAssessmentInstance creates a Journal entry for the supplied non-recurring assessment
//=================================================================================================
func ProcessNewAssessmentInstance(xbiz *rlib.XBusiness, d1, d2 *time.Time, a *rlib.Assessment) error {
	funcname := "ProcessNewAssessmentInstance"
	var noerr error
	if a.PASMID == 0 && a.RentCycle != rlib.RECURNONE { // if this assessment is not a single instance recurrence, then return an error
		return fmt.Errorf("%s: Function only accepts non-recurring instances", funcname)
	}
	// if a.ASMID != 0 && a.RentCycle { // if this assessment has already been written to db, then return error
	// 	return fmt.Errorf("%s: Function only accepts unsaved instances. Found ASMID = %d", funcname, a.ASMID)
	// }
	if a.ASMID == 0 {
		ASMID, err := rlib.InsertAssessment(a)
		if nil != err {
			return err
		}
		a.ASMID = ASMID
	}
	journalAssessment(xbiz, a.Start, a, d1, d2)
	return noerr
}

// ProcessNewReceipt creates a Journal entry for the supplied receipt
//=================================================================================================
func ProcessNewReceipt(xbiz *rlib.XBusiness, d1, d2 *time.Time, r *rlib.Receipt) error {
	rntagr, _ := rlib.GetRentalAgreement(r.RAID)
	var j rlib.Journal
	j.BID = rntagr.BID
	j.Amount = rlib.RoundToCent(r.Amount)
	j.Dt = r.Dt
	j.Type = rlib.JNLTYPERCPT
	j.ID = r.RCPTID
	j.RAID = r.RAID
	jid, err := rlib.InsertJournalEntry(&j)
	if err != nil {
		rlib.Ulog("Error inserting Journal entry: %v\n", err)
		return err
	}
	if jid > 0 {
		// now add the Journal allocation records...
		for j := 0; j < len(r.RA); j++ {
			var ja rlib.JournalAllocation
			ja.JID = jid
			ja.Amount = rlib.RoundToCent(r.RA[j].Amount)
			ja.ASMID = r.RA[j].ASMID
			ja.AcctRule = r.RA[j].AcctRule
			a, _ := rlib.GetAssessment(ja.ASMID)
			ja.RID = a.RID
			rlib.InsertJournalAllocationEntry(&ja)
		}
	}
	return err
}

// GenerateJournalRecords creates Journal records for Assessments and receipts over the supplied time range.
//=================================================================================================
func GenerateJournalRecords(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	err := RemoveJournalEntries(xbiz, d1, d2)
	if err != nil {
		rlib.Ulog("Could not remove existing Journal entries from %s to %s. err = %v\n", d1.Format(rlib.RRDATEFMT), d2.Format(rlib.RRDATEFMT), err)
		return
	}

	//-----------------------------------------------------------
	//  PROCESS ASSESSMSENTS
	//-----------------------------------------------------------
	// fmt.Printf("GetAllAssessmentsByBusiness - d2 = %s   d1 = %s\n", d2.Format(rlib.RRDATEINPFMT), d1.Format(rlib.RRDATEINPFMT))
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(xbiz.P.BID, d2, d1) // only get instances without a parent
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a rlib.Assessment
		ap := &a
		rlib.ReadAssessments(rows, &a)
		if a.RentCycle == rlib.RECURNONE {
			// journalAssessment(xbiz, a.Start, &a, d1, d2)
			ProcessNewAssessmentInstance(xbiz, d1, d2, &a)
		} else if a.RentCycle >= rlib.RECURSECONDLY && a.RentCycle <= rlib.RECURHOURLY {
			// TBD
			fmt.Printf("Unhandled assessment recurrence type: %d\n", a.RentCycle)
		} else {
			dl := ap.GetRecurrences(d1, d2)
			// fmt.Printf("type = %d, %s - %s    len(dl) = %d\n", a.ATypeLID, a.Start.Format(rlib.RRDATEFMT), a.Stop.Format(rlib.RRDATEFMT), len(dl))
			for i := 0; i < len(dl); i++ {
				a1 := a
				a1.Start = dl[i]    // use the instance date
				a1.Stop = a.Start   // start and stop are the same
				a1.ASMID = 0        // ensure this is a new assessment
				a1.PASMID = a.ASMID // parent assessment
				ProcessNewAssessmentInstance(xbiz, d1, d2, &a1)
			}
		}
	}
	rlib.Errcheck(rows.Err())

	//-----------------------------------------------------------
	//  COMPUTE VACANCY
	//-----------------------------------------------------------
	GenVacancyJournals(xbiz, d1, d2)

	//-----------------------------------------------------------
	//  PROCESS RECEIPTS
	//-----------------------------------------------------------
	r := rlib.GetReceipts(xbiz.P.BID, d1, d2)
	for i := 0; i < len(r); i++ {
		ProcessNewReceipt(xbiz, d1, d2, &r[i])
	}

	//-----------------------------------------------------------
	//  ADD JOURNAL MARKER
	//-----------------------------------------------------------
	var jm rlib.JournalMarker
	jm.BID = xbiz.P.BID
	jm.State = rlib.MARKERSTATEOPEN
	jm.DtStart = *d1
	jm.DtStop = (*d2).AddDate(0, 0, -1)
	rlib.InsertJournalMarker(&jm)
}
