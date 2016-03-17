package main

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"
)

// journalAssessment processes the assessment, creates a journal entry, and returns its id
func journalAssessment(d time.Time, a *Assessment, d1, d2 *time.Time) (int, float32) {
	// r := GetRentable(a.RID)
	// xp := GetXPersonByPID(r.PID)
	// s := fmt.Sprintf("A%08d  %s", a.ASMID, App.AsmtTypes[a.ASMTID].Name)
	// if rentDuration != assessmentDuration {
	// 	s = fmt.Sprintf("%s (%d/%d days)", s, rentDuration, assessmentDuration)
	// }
	// printJournalSubtitle(s)
	// processAcctRuleAmount(d, a.AcctRule, a.RAID, a.Amount*pf, &r)
	// printJournalSubtitle("")
	//-------------------------------------------------------------------
	// over what range of time does this rental apply between d1 & d2
	//-------------------------------------------------------------------
	ra, _ := GetRentalAgreement(a.RAID)
	start := *d1
	if ra.RentalStart.After(start) {
		start = ra.RentalStart
	}
	stop := ra.RentalStop.Add(24 * 60 * time.Minute)
	if stop.After(*d2) {
		stop = *d2
	}
	//-------------------------------------------------------------------------------------------
	// this code needs to be generalized based on the recurrence period and the proration period
	//-------------------------------------------------------------------------------------------
	assessmentDuration := int(d2.Sub(*d1).Hours() / 24)
	rentDuration := int(stop.Sub(start).Hours() / 24)
	pf := float32(1.0)
	if rentDuration != assessmentDuration && a.ProrationMethod > 0 {
		pf = float32(rentDuration) / float32(assessmentDuration)
	} else {
		rentDuration = assessmentDuration
	}

	var j Journal
	j.BID = a.BID
	j.Amount = a.Amount * pf
	j.Dt = d
	j.Type = JNLTYPEASMT
	j.ID = a.ASMID
	j.RAID = a.RAID

	// fmt.Printf("Amount = %6.2f\n", j.Amount)

	jid, err := InsertJournalEntry(&j)
	if err != nil {
		ulog("error inserting journal entry: %v\n", err)
	}

	return jid, pf
}

// journalAllocation assumes that all fields except AcctRule are filled in
func journalAllocation(ja *JournalAllocation, rule string) {
	if len(rule) > 0 {
		sa := strings.Split(rule, ",")
		for k := 0; k < len(sa); k++ {
			t := strings.TrimSpace(sa[k])
			ta := strings.Split(t, " ")
			action := strings.ToLower(strings.TrimSpace(ta[0]))
			acct := strings.TrimSpace(ta[1])
			if action == "d" {
				ja.AcctRule = fmt.Sprintf("c %s, d 10001", acct)
			}
		}
	}
	InsertJournalAllocationEntry(ja)
}

// RemoveJournalEntries clears out the records in the supplied range provided the range is not closed by a journalmarker
func RemoveJournalEntries(xprop *XBusiness, d1, d2 *time.Time) error {
	// Remove the journal entries and the journalallocation entries
	rows, err := App.prepstmt.getAllJournalsInRange.Query(xprop.P.BID, d1, d2)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var j Journal
		rlib.Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID))
		deleteJournalAllocations(j.JID)
		deleteJournalEntry(j.JID)
	}

	// only delete the marker if it is in this time range and if it is not the origin marker
	jm := GetLastJournalMarker()
	if jm.State == MARKERSTATEOPEN {
		deleteJournalMarker(jm.JMID)
	}

	return err
}

// GenerateJournalRecords creates journal records for assessments and receipts over the supplied time range.
func GenerateJournalRecords(xprop *XBusiness, d1, d2 *time.Time) {
	err := RemoveJournalEntries(xprop, d1, d2)
	if err != nil {
		ulog("Could not remove existin Journal Entries from %s to %s\n", d1.Format(RRDATEFMT), d2.Format(RRDATEFMT))
		return
	}

	//===========================================================
	//  PROCESS ASSESSMSENTS
	//===========================================================
	rows, err := App.prepstmt.getAllAssessmentsByBusiness.Query(xprop.P.BID, d2, d1)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a Assessment
		ap := &a
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod, &a.AcctRule, &a.LastModTime, &a.LastModBy))
		if a.Frequency >= rlib.RECURSECONDLY && a.Frequency <= rlib.RECURHOURLY {
			// TBD
			fmt.Printf("Unhandled assessment recurrence type: %d\n", a.Frequency)
		} else {
			dl := ap.GetRecurrences(d1, d2)
			// fmt.Printf("type = %d, %s - %s    len(dl) = %d\n", a.ASMTID, a.Start.Format(RRDATEFMT), a.Stop.Format(RRDATEFMT), len(dl))
			for i := 0; i < len(dl); i++ {
				var ja JournalAllocation
				jid, pf := journalAssessment(dl[i], &a, d1, d2)
				if jid > 0 {
					ja.JID = jid
					ja.ASMID = a.ASMID
					ja.Amount = a.Amount * pf
					ja.AcctRule = a.AcctRule
					InsertJournalAllocationEntry(&ja)
				}
			}
		}
	}
	rlib.Errcheck(rows.Err())

	//===========================================================
	//  PROCESS RECEIPTS
	//===========================================================
	r := GetReceipts(xprop.P.BID, d1, d2)
	for i := 0; i < len(r); i++ {
		rntagr, _ := GetRentalAgreement(r[i].RAID)
		var j Journal
		j.BID = rntagr.BID
		j.Amount = r[i].Amount
		j.Dt = r[i].Dt
		j.Type = JNLTYPERCPT
		j.ID = r[i].RCPTID
		j.RAID = r[i].RAID
		jid, err := InsertJournalEntry(&j)
		if err != nil {
			ulog("Error inserting journal entry: %v\n", err)
		}
		if jid > 0 {
			// now add the journal allocation records...
			for j := 0; j < len(r[i].RA); j++ {
				var ja JournalAllocation
				ja.JID = jid
				ja.Amount = r[i].RA[j].Amount
				ja.ASMID = r[i].RA[j].ASMID
				ja.AcctRule = ""

				a, err := GetAssessment(ja.ASMID)
				if err != nil {
					ulog("Error reading assessment %d:  %s\n", ja.ASMID, err)
				} else {
					journalAllocation(&ja, a.AcctRule)
				}
			}
		}
	}

	//===========================================================
	//  ADD JOURNAL MARKER
	//===========================================================
	var jm JournalMarker
	jm.BID = xprop.P.BID
	jm.State = MARKERSTATEOPEN
	jm.DtStart = *d1
	jm.DtStop = *d2
	InsertJournalMarker(&jm)
}
