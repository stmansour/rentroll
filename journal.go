package main

import (
	"fmt"
	"math"
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

// calcProrationInfo is currently designed to work for nonrecurring, daily and monthly recurring rentals.
// Other frequencies will need to be added.
//
// Parameters:
//  	AgrStart,AgrStop: define the time range of either the rentalAgreement or the Assessment
//						  covering the Rentable
//  	d1, d2:           the time period we're being asked to analyze
//  	accrual:          recurring frequency of rent
//  	prorateMethod:    the recurrence frequency to use for partial coverage
//
// Returns:
//			asmtDur:  pf denominator, total number of days in period
//         	rentDur:  pf numerator, total number of days applicable to this rental agreement
//         	pf:       prorate factor = rentDur/asmtDur
//=================================================================================================
func calcProrationInfo(DtStart, DtStop, d1, d2 *time.Time, rentCycle, prorate int64) float64 {
	//-------------------------------------------------------------------
	// over what range of time does this rental apply between d1 & d2
	//-------------------------------------------------------------------
	start := *d1
	if DtStart.After(start) {
		start = *DtStart
	}
	stop := *DtStop // .Add(24 * 60 * time.Minute) -- removing this as all ranges must be NON-INCLUSIVE
	if stop.After(*d2) {
		stop = *d2
	}

	pf := float64(1.0)                          // assume full period
	cycleTime := d2.Sub(*d1)                    // denominator
	thisPeriod := stop.Sub(start)               // numerator
	if cycleTime != thisPeriod && prorate > 0 { // if cycle time and period differ AND it's NOT a one-time charge
		pf = float64(thisPeriod) / float64(cycleTime)
	}
	return pf
}

// journalAssessment processes the assessment, creates a Journal entry, and returns its id
//=================================================================================================
func journalAssessment(xbiz *rlib.XBusiness, rid int64, d time.Time, a *rlib.Assessment, d1, d2 *time.Time) error {
	pf := float64(0)

	r := rlib.GetRentable(rid)
	status := rlib.GetRentableStateForDate(r.RID, &d)
	switch status {
	case rlib.RENTABLESTATUSONLINE:
		ra, _ := rlib.GetRentalAgreement(a.RAID)
		switch a.RentCycle {
		case rlib.ACCRUALDAILY:
			pf = calcProrationInfo(&ra.PossessionStart, &ra.PossessionStop, &d, &d, a.RentCycle, a.ProrationCycle)
		case rlib.ACCRUALNORECUR:
			fallthrough
		case rlib.ACCRUALMONTHLY:
			pf = calcProrationInfo(&ra.PossessionStart, &ra.PossessionStop, d1, d2, a.RentCycle, a.ProrationCycle)
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
				rlib.Ulog("journalAssessment: error getting rent cycle for rentable %d. err = %s\n", r.RID, err.Error())
			}
			pf = calcProrationInfo(&(ta[0].Start), &(ta[0].Stop), d1, d2, rentcycle, proration)
			if len(ta) > 1 {
				rlib.Ulog("journalAssessment: %d Assessments affect Rentable %d (%s) in period %s - %s\n",
					len(ta), r.RID, r.Name, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
			}
		}
	default:
		rlib.Ulog("journalAssessment: Rentable %d is in an unknown status: %d\n", r.RID, status)
	}

	var j = rlib.Journal{BID: a.BID, Dt: d, Type: rlib.JNLTYPEASMT, ID: a.ASMID, RAID: a.RAID}

	m := rlib.ParseAcctRule(xbiz, rid, &d, &d, a.AcctRule, a.Amount, pf) // a rule such as "d 11001 1000.0, c 40001 1100.0, d 41004 100.00"
	_, j.Amount = sumAllocations(&m)
	j.Amount = rlib.RoundToCent(j.Amount)
	// fmt.Printf("After ParseAcctRule - j.Amount = %8.2f\n", j.Amount)

	//-------------------------------------------------------------------------------------------
	// In the event that we need to prorate, pull together the pieces and determine the
	// fractional amounts so that all the entries can net to 0.00.  Essentially, this means
	// handling the $0.01 off problem when dealing with fractional numbers.  The way we'll
	// handle this is to apply the extra cent to the largest number
	//-------------------------------------------------------------------------------------------
	if pf < 1.0 {
		sum := float64(0.0)
		debits := float64(0)
		k := 0 // index of the largest number
		for i := 0; i < len(m); i++ {
			m[i].Amount = rlib.RoundToCent(m[i].Amount)
			if m[i].Amount > m[k].Amount {
				k = i
			}
			if m[i].Action == "c" {
				sum -= m[i].Amount
			} else {
				sum += m[i].Amount
				debits += m[i].Amount
			}
		}
		if sum != float64(0) {
			m[k].Amount += sum // first try adding the penny
			x, xd := sumAllocations(&m)
			j.Amount = rlib.RoundToCent(xd)
			if x != float64(0) { // if that doesn't work...
				m[k].Amount -= sum + sum // subtract the penny:  remove the one we added above, then remove another, i.e.: sum + sum
				y, yd := sumAllocations(&m)
				j.Amount = rlib.RoundToCent(yd)
				// if there's some strange number that causes issues, use the one closest to 0
				if math.Abs(float64(y)) > math.Abs(float64(x)) { // if y is farther from 0 than x, go back to the value for x
					m[k].Amount += sum + sum
					j.Amount = xd
				}
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
			ja.RID = rid
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
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(xbiz.P.BID, d2, d1)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a rlib.Assessment
		ap := &a
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.ASMTID, &a.RAID, &a.Amount,
			&a.Start, &a.Stop, &a.RentCycle, &a.ProrationCycle, &a.AcctRule, &a.Comment,
			&a.LastModTime, &a.LastModBy))
		// fmt.Printf("Assessment: ASMID = %d, Amount = %8.2f\n", a.ASMID, a.Amount)
		if a.RentCycle >= rlib.RECURSECONDLY && a.RentCycle <= rlib.RECURHOURLY {
			// TBD
			fmt.Printf("Unhandled assessment recurrence type: %d\n", a.RentCycle)
		} else {
			dl := ap.GetRecurrences(d1, d2)
			// fmt.Printf("type = %d, %s - %s    len(dl) = %d\n", a.ASMTID, a.Start.Format(rlib.RRDATEFMT), a.Stop.Format(rlib.RRDATEFMT), len(dl))
			for i := 0; i < len(dl); i++ {
				journalAssessment(xbiz, a.RID, dl[i], &a, d1, d2)
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
		rntagr, _ := rlib.GetRentalAgreement(r[i].RAID)
		var j rlib.Journal
		j.BID = rntagr.BID
		j.Amount = rlib.RoundToCent(r[i].Amount)
		j.Dt = r[i].Dt
		j.Type = rlib.JNLTYPERCPT
		j.ID = r[i].RCPTID
		j.RAID = r[i].RAID
		jid, err := rlib.InsertJournalEntry(&j)
		if err != nil {
			rlib.Ulog("Error inserting Journal entry: %v\n", err)
		}
		if jid > 0 {
			// now add the Journal allocation records...
			for j := 0; j < len(r[i].RA); j++ {
				var ja rlib.JournalAllocation
				ja.JID = jid
				ja.Amount = rlib.RoundToCent(r[i].RA[j].Amount)
				ja.ASMID = r[i].RA[j].ASMID
				ja.AcctRule = r[i].RA[j].AcctRule
				a, _ := rlib.GetAssessment(ja.ASMID)
				ja.RID = a.RID
				rlib.InsertJournalAllocationEntry(&ja)
			}
		}
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
