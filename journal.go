package main

import (
	"fmt"
	"math"
	"rentroll/rlib"
	"time"
)

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

// calcProrationInfo returns:
//  AgrStart,AgrStop - define the time range of either the rentalAgreement or the Assessment covering the rentable
//  d1, d2 - the time period we're being asked to analyze
//  prorateMethod - the recurrence frequency to use for partial coverage
//
// Returns:
//			asmtDur:  pf denominator, total number of days in period
//         	rentDur:  pf numerator, total number of days applicable to this rental agreement
//         	pf:       prorate factor = rentDur/asmtDur
//func calcProrationInfo(ra *rlib.RentalAgreement, d1, d2 *time.Time, prorateMethod int64) (int64, int64, float64) {
func calcProrationInfo(DtStart, DtStop, d1, d2 *time.Time, prorateMethod int64) (int64, int64, float64) {
	//-------------------------------------------------------------------
	// over what range of time does this rental apply between d1 & d2
	//-------------------------------------------------------------------
	start := *d1
	if DtStart.After(start) {
		start = *DtStart
	}
	stop := DtStop.Add(24 * 60 * time.Minute)
	if stop.After(*d2) {
		stop = *d2
	}

	//-------------------------------------------------------------------------------------------
	// this code needs to be generalized based on the recurrence period and the proration period
	//-------------------------------------------------------------------------------------------
	asmtDur := int64(d2.Sub(*d1).Hours() / 24)
	rentDur := int64(stop.Sub(start).Hours() / 24)
	pf := float64(1.0)
	if rentDur != asmtDur && prorateMethod > 0 {
		pf = float64(rentDur) / float64(asmtDur)
	} else {
		rentDur = asmtDur
	}
	return asmtDur, rentDur, pf
}

// journalAssessment processes the assessment, creates a journal entry, and returns its id
func journalAssessment(xbiz *rlib.XBusiness, rid int64, d time.Time, a *rlib.Assessment, d1, d2 *time.Time) error {
	pf := float64(0)

	// ra, _ := rlib.GetRentalAgreement(a.RAID)
	// _, _, pf = calcProrationInfo(&ra.RentalStart, &ra.RentalStop, d1, d2, a.ProrationMethod)

	r := rlib.GetRentable(rid)
	switch {
	case r.State == rlib.RENTABLESTATEONLINE:
		ra, _ := rlib.GetRentalAgreement(a.RAID)
		_, _, pf = calcProrationInfo(&ra.RentalStart, &ra.RentalStop, d1, d2, a.ProrationMethod)

	case r.State == rlib.RENTABLESTATEADMIN ||
		r.State == rlib.RENTABLESTATEEMPLOYEE ||
		r.State == rlib.RENTABLESTATEOWNEROCC ||
		r.State == rlib.RENTABLESTATEOFFLINE:
		ta := rlib.GetAllRentableAssessments(r.RID, d1, d2)
		if len(ta) > 0 {
			_, _, pf = calcProrationInfo(&(ta[0].Start), &(ta[0].Stop), d1, d2, xbiz.RT[r.RTID].Proration)
			if len(ta) > 1 {
				rlib.Ulog("journalAssessment: %d assessments affect rentable %d (%s) in period %s - %s\n", len(ta), r.RID, r.Name, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
			}
		}
	default:
		rlib.Ulog("journalAssessment: rentable %d is in an unknown state: %d\n", r.RID, r.State)
	}

	var j = rlib.Journal{BID: a.BID, Dt: d, Type: rlib.JNLTYPEASMT, ID: a.ASMID, RAID: a.RAID}

	m := rlib.ParseAcctRule(xbiz, rid, d1, d2, a.AcctRule, a.Amount, pf) // a rule such as "d 11001 1000.0, c 40001 1100.0, d 41004 100.00"
	_, j.Amount = sumAllocations(&m)
	j.Amount = rlib.RoundToCent(j.Amount)

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
				m[k].Amount -= sum + sum // subtract the penny
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

	// fmt.Printf("INSERTING JOURNAL: %#v\n", j)
	jid, err := rlib.InsertJournalEntry(&j)
	if err != nil {
		rlib.Ulog("error inserting journal entry: %v\n", err)
	} else {
		//now rewrite the AcctRule...
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

// RemoveJournalEntries clears out the records in the supplied range provided the range is not closed by a journalmarker
func RemoveJournalEntries(xbiz *rlib.XBusiness, d1, d2 *time.Time) error {
	// Remove the journal entries and the journalallocation entries
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

// GenerateJournalRecords creates journal records for assessments and receipts over the supplied time range.
func GenerateJournalRecords(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	err := RemoveJournalEntries(xbiz, d1, d2)
	if err != nil {
		rlib.Ulog("Could not remove existing Journal entries from %s to %s. err = %v\n", d1.Format(rlib.RRDATEFMT), d2.Format(rlib.RRDATEFMT), err)
		return
	}

	//===========================================================
	//  PROCESS ASSESSMSENTS
	//===========================================================
	// fmt.Printf("GetAllAssessmentsByBusiness - d2 = %s   d1 = %s\n", d2.Format(rlib.RRDATEINPFMT), d1.Format(rlib.RRDATEINPFMT))
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(xbiz.P.BID, d2, d1)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a rlib.Assessment
		ap := &a
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.ASMTID, &a.RAID, &a.Amount,
			&a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod, &a.AcctRule, &a.Comment,
			&a.LastModTime, &a.LastModBy))
		// fmt.Printf("Assessment: %#v\n", a)
		if a.Frequency >= rlib.RECURSECONDLY && a.Frequency <= rlib.RECURHOURLY {
			// TBD
			fmt.Printf("Unhandled assessment recurrence type: %d\n", a.Frequency)
		} else {
			dl := ap.GetRecurrences(d1, d2)
			// fmt.Printf("type = %d, %s - %s    len(dl) = %d\n", a.ASMTID, a.Start.Format(rlib.RRDATEFMT), a.Stop.Format(rlib.RRDATEFMT), len(dl))
			for i := 0; i < len(dl); i++ {
				journalAssessment(xbiz, a.RID, dl[i], &a, d1, d2)
			}
		}
	}
	rlib.Errcheck(rows.Err())

	//===========================================================
	//  COMPUTE VACANCY
	//===========================================================
	GenVacancyJournals(xbiz, d1, d2)

	//===========================================================
	//  PROCESS RECEIPTS
	//===========================================================
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
			rlib.Ulog("Error inserting journal entry: %v\n", err)
		}
		if jid > 0 {
			// now add the journal allocation records...
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

	//===========================================================
	//  ADD JOURNAL MARKER
	//===========================================================
	var jm rlib.JournalMarker
	jm.BID = xbiz.P.BID
	jm.State = rlib.MARKERSTATEOPEN
	jm.DtStart = *d1
	jm.DtStop = (*d2).AddDate(0, 0, -1)
	rlib.InsertJournalMarker(&jm)
}
