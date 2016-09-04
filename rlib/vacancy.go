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
//============================================================================================
func ProcessRentable(xbiz *XBusiness, d1, d2 *time.Time, r *Rentable) {
	m := VacancyDetect(xbiz, d1, d2, r)
	// fmt.Printf("ProcessRentable: r = %s (%d), period=(%s - %s) len(m) = %d\n", r.Name, r.RID, d1.Format("Jan 2"), d2.Format("Jan 2"), len(m))
	for i := 0; i < len(m); i++ {
		// the umr rate is in cost/accrualDuration. The duration of the VacancyMarkers
		// are in integral multiples of rangeIncDur.  We need to prorate the amount of
		// each entry accordingly
		var j Journal
		j.BID = xbiz.P.BID
		j.Amount = RoundToCent(m[i].Amount)
		j.Dt = m[i].DtStop.AddDate(0, 0, -1) // associated date is last day of period
		j.Type = JNLTYPEUNAS                 // this is an unassociated entry
		j.RAID = 0                           // we really mean it, it is unassociated
		j.ID = r.RID                         // mark the associated Rentable
		j.Comment = m[i].Comment             // this will note consecutive days for vacancy
		// fmt.Printf("ProcessRentable: insert journal entry: %s - %s, %8.2f\n", j.Dt.Format(RRDATEINPFMT), j.Comment, j.Amount)
		jid, err := InsertJournalEntry(&j)
		Errlog(err)
		if jid > 0 {
			var ja JournalAllocation
			ja.JID = jid
			ja.Amount = j.Amount
			ja.ASMID = 0 // it's unassociated
			ja.AcctRule = "c ${GLGSRENT} _,d ${GLVAC} _"
			ja.RID = r.RID
			// fmt.Printf("VACANCY: inserting journalAllocation entry: %#v\n", ja)
			InsertJournalAllocationEntry(&ja)
		}
	}
}

// GenVacancyJournals creates Journal entries that cover vacancy for
// every Rentable where the Rentable type is being managed to budget
//===============================================================================================
func GenVacancyJournals(xbiz *XBusiness, d1, d2 *time.Time) {
	rows, err := RRdb.Prepstmt.GetAllRentablesByBusiness.Query(xbiz.P.BID)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r Rentable
		Errcheck(rows.Scan(&r.RID, &r.BID, &r.Name, &r.AssignmentTime, &r.LastModTime, &r.LastModBy))
		ProcessRentable(xbiz, d1, d2, &r)
	}
	Errcheck(rows.Err())
}
