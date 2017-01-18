package rlib

import (
	"time"
)

// ProcessRentable looks for any time period for which the Rentable has
// no rent assessment during the supplied time range. If it finds any
// vacant time periods, it generates vacancy Journal entries for that period.
// The approach here is to look for all rental agreements that apply to this
// Rentable over the supplied time period. Spin through each of them and
// count the number of days that are covered.  Compare this to the total
// number of days in the period. Then generate a vacancy entry for the time
// period for which no rent was assessed.
// The return value is the number of vacancy records added
//============================================================================================
func ProcessRentable(xbiz *XBusiness, d1, d2 *time.Time, r *Rentable) int {
	nr := 0
	m := VacancyDetect(xbiz, d1, d2, r)
	// fmt.Printf("ProcessRentable: r = %s (%d), period=(%s - %s) len(m) = %d\n", r.Name, r.RID, d1.Format("Jan 2"), d2.Format("Jan 2"), len(m))
	for i := 0; i < len(m); i++ {
		// the umr rate is in cost/accrualDuration. The duration of the VacancyMarkers
		// are in integral multiples of rangeIncDur.  We need to prorate the amount of
		// each entry accordingly
		var j Journal
		j.BID = xbiz.P.BID
		j.Amount = RoundToCent(m[i].Amount)
		// TODO: fix the next line
		j.Dt = m[i].DtStop.AddDate(0, 0, -1) // associated date is period end - 1 proration cycle (or 1 sec if no proration)
		j.Type = JNLTYPEUNAS                 // this is an unassociated entry
		j.RAID = 0                           // we really mean it, it is unassociated
		j.ID = r.RID                         // mark the associated Rentable
		j.Comment = m[i].Comment             // this will note consecutive days for vacancy
		// fmt.Printf("ProcessRentable: insert journal entry: %s - %s, %8.2f\n", j.Dt.Format(RRDATEINPFMT), j.Comment, j.Amount)

		// TODO: this check must be more thorough. Examine the RentCycle, query over the time of the rent cycle
		//       to see if other records were generated and handle.
		//       By convention, the period for Vacancy detection is:  supplied range, date/time of journal entry =
		//       rentcycle - 1 prorationcycle (or one second whichever is larger)
		// These entries must be idempotent. Make sure it does not already exist.
		jv := GetJournalVacancy(r.RID, &j.Dt, &m[i].DtStop)
		if jv.JID != 0 { // if the JID >0 ..
			continue // then this entry was already generated, keep going
		}

		jid, err := InsertJournalEntry(&j)
		Errlog(err)
		nr++
		if jid > 0 {
			var ja JournalAllocation
			ja.JID = jid
			ja.Amount = j.Amount
			ja.ASMID = 0 // it's unassociated
			ja.AcctRule = "c ${GLGSRENT} _,d ${GLVAC} _"
			ja.RID = r.RID
			ja.BID = r.BID
			// fmt.Printf("VACANCY: inserting journalAllocation entry: %#v\n", ja)
			InsertJournalAllocationEntry(&ja)
			j.JA = append(j.JA, ja)
		}
		initLedgerCache()
		GenerateLedgerEntriesFromJournal(xbiz, &j, d1, d2)
	}
	return nr
}

// GenVacancyJournals creates Journal entries that cover vacancy for
// every Rentable where the Rentable type is being managed to budget
//===============================================================================================
func GenVacancyJournals(xbiz *XBusiness, d1, d2 *time.Time) int {
	nr := 0
	rows, err := RRdb.Prepstmt.GetAllRentablesByBusiness.Query(xbiz.P.BID)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r Rentable
		Errcheck(rows.Scan(&r.RID, &r.BID, &r.Name, &r.AssignmentTime, &r.LastModTime, &r.LastModBy))
		nr += ProcessRentable(xbiz, d1, d2, &r)
	}
	Errcheck(rows.Err())
	return nr
}
