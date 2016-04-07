package main

import (
	"rentroll/rlib"
	"time"
)

// ProcessRentable looks for any time period for which the rentable has
// no rent assessment during the supplied time range. If it finds any
// vacant time periods, it generates vacancy journal entries for that period.
// The approach here is to look for all rental agreements that apply to this
// rentable over the supplied time period. Spin through each of them and
// count the number of days that are covered.  Compare this to the total
// number of days in the period. Then generate a vacancy entry for the time
// period for which no rent was assessed.
func ProcessRentable(xbiz *XBusiness, d1, d2 *time.Time, r *Rentable) {
	//--------------------------------------------------------------------------------------
	// Find all rental agreements that cover r during time period d1-d2
	//--------------------------------------------------------------------------------------
	t := GetAgreementsForRentable(r.RID, d1, d2)
	var n = int64(0) // total number of days covered by all rental agreements for this rentable during d1-d2
	var m int64      // total number of days in the period d1-d2
	var k int64      // number of days covered by an agreement.

	// fmt.Printf("RAs for %s:  %d\n", r.Name, len(t))

	for i := 0; i < len(t); i++ {
		ra, _ := GetRentalAgreement(t[i].RAID)
		if t[i].RID == r.RID {
			m, k, _ = calcProrationInfo(&ra, d1, d2, xbiz.RT[r.RTID].Proration)
			n += k
		}
	}

	//--------------------------------------------------------------------------------------
	// if no rental agreements for this rentable, then n will be 0
	// otherwise, if the total number of days (m) is not covered by n, then we have vacancy
	//--------------------------------------------------------------------------------------
	if n == 0 || n != m {
		pf := float64(1)
		umr := GetRentableMarketRate(xbiz, r, d1, d2) // this call does the right thing whether or not the rentable is a unit
		if n != 0 {
			pf = float64(m-n) / float64(m)
		}
		// fmt.Printf("Rentable %s (rid=%d, unitid=%d). Period = %d days, covered for %d days, vacant for %d days. MarketRate = %6.2f. prorated: %6.2f\n", r.Name, r.RID, r.UNITID, m, n, m-n, umr, umr*pf)
		var j Journal
		j.BID = xbiz.P.BID
		j.Amount = rlib.RoundToCent(umr * pf)
		j.Dt = d2.AddDate(0, 0, -1) // associated date is last day of period
		if d1.After(j.Dt) {
			j.Dt = *d1
		}
		j.Type = JNLTYPEUNAS
		j.ID = r.RID
		j.RAID = 0 // this one is unassociated
		jid, err := InsertJournalEntry(&j)
		if err != nil {
			ulog("Error inserting journal entry: %v\n", err)
		}
		if jid > 0 {
			var ja JournalAllocation
			ja.JID = jid
			ja.Amount = rlib.RoundToCent(j.Amount)
			ja.ASMID = 0 // it's unassociated
			ja.AcctRule = "c ${DFLTGSRENT} _,d ${DFLTVAC} _"
			ja.RID = r.RID
			InsertJournalAllocationEntry(&ja)
		}
	}
}

// GenVacancyJournals creates journal entries that cover vacancy for
// every rentable where the rentable type is being managed to budget
func GenVacancyJournals(xbiz *XBusiness, d1, d2 *time.Time) {
	rows, err := App.prepstmt.getAllRentablesByBusiness.Query(xbiz.P.BID)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r Rentable
		rlib.Errcheck(rows.Scan(&r.RID, &r.LID, &r.RTID, &r.BID, &r.UNITID, &r.Name, &r.Assignment, &r.Report, &r.LastModTime, &r.LastModBy))
		if xbiz.RT[r.RTID].ManageToBudget > 0 {
			ProcessRentable(xbiz, d1, d2, &r)
		}
	}
	rlib.Errcheck(rows.Err())

}
