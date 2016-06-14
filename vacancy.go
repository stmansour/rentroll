package main

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// VacancyMarker is a structure of data defining an increment in time during which a Rentable is vacant
type VacancyMarker struct {
	DtStart time.Time // a period start time
	DtStop  time.Time // end of period
	Amount  float64   // unit market rate during this period
	comment string    // comment to include with Journal
	state   int64     // Rentable state
}

// VacancyDetect scans the time range specified and looks for pro[rate] periods of time when the
// supplied Rentable is not accounted for. For every period that it is not rented
// a VacancyMarker will be added to an array marking the vacant time period. The return value
// is the list of vacancy markers.
//========================================================================================================
func VacancyDetect(xbiz *rlib.XBusiness, d1, d2 *time.Time, r *rlib.Rentable) []VacancyMarker {
	var m []VacancyMarker
	var state int64

	//==================================================================
	// Whether it's vacant or not depends on its state. For example,
	// if it is OwnerOccupied, no rent is collected and it is not
	// considered vacant. So, the first thing to do is cache the
	// Rentable state over the period
	//==================================================================
	rsa := rlib.GetRentableStatusByRange(r.RID, d1, d2)

	//=====================================================================================
	// In the loop below, we don't want to read the database every iteration
	// for the RTID associated with the rentable as it would result in excessive database
	// reatds. In most cases, the RTID will not change, especially over short periods like
	// a month. So, read all the RTIDs for the period that the loop will process first.
	// then select from them as needed.
	//=====================================================================================
	rta := rlib.GetRentableTypeRefsByRange(r.RID, d1, d2) // get the list
	if len(rta) == 0 {
		rlib.Ulog("VacancyDetect:  No valid RTID for rentable R%08d during period %s to %s\n",
			r.RID, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
		return m // this is bad! No RTID for the supplied time range
	}
	rtidMulti := len(rta) > 1 // flag to indicate we need to look for a change in rtid in every pass
	rtid := rta[0].RTID       // initialize to the first RTID

	// We may not need to do anything if this rentable is not being managed to budget.  We didn't
	// check it earlier because the code to load the rentable type is here. If there's an issue,
	// just move the code to grabe the RTIDs to the caller and pass the array into this func.
	if xbiz.RT[rtid].ManageToBudget == 0 { // if this rentable is not managing to budget...
		return m // return an empty list now and it will essentially be ignored.
	}

	period := rlib.CycleDuration(xbiz.RT[rtid].Proration, *d1)
	t := rlib.GetAgreementsForRentable(r.RID, d1, d2) // t is an array of RentalAgreementRentables

	//========================================================
	// Mark vacancy for each time interval between d1 & d2
	//========================================================
	var dtNext time.Time
	k := 0 // number of members of m
	for dt := *d1; dt.Before(*d2); dt = dtNext {
		dtNext = dt.Add(period)
		vacant := true // assume it's vacant and reset if we find it's rented

		// fmt.Printf("VacancyDetect:  period %s - %s\n", dt.Format(rlib.RRDATEINPFMT), dtNext.Format(rlib.RRDATEINPFMT))

		rs := rlib.SelectRentableStatusForPeriod(&rsa, dt, dtNext)
		state = rlib.RENTABLESTATUSONLINE // if there is no state info, we'll assume online
		if len(rs) > 0 {
			state = rs[0].Status // If this turns out to be a problem, maybe we'll choose the state with the greatest percentage of time
		}

		switch state {
		case rlib.RENTABLESTATUSONLINE:
			// fmt.Printf("\tonline... ")
			for i := 0; i < len(t); i++ {
				if rlib.DateRangeOverlap(&t[i].DtStart, &t[i].DtStop, &dt, &dtNext) {
					// fmt.Printf("covered, RAID = %d\n", t[i].RAID)
					vacant = false // not vacant
				}
			}
		case rlib.RENTABLESTATUSADMIN:
			fallthrough
		case rlib.RENTABLESTATUSEMPLOYEE:
			fallthrough
		case rlib.RENTABLESTATUSOWNEROCC:
			fallthrough
		case rlib.RENTABLESTATUSOFFLINE:
			// fmt.Printf("\t{admin|employee|ownerocc|offline}... ")
		}
		if !vacant {
			continue
		}

		// update rtid only if its type changes during this report period...
		if rtidMulti {
			rt := rlib.SelectRentableTypeRefForDate(&rta, &dt)
			rtid = rt.RTID
			if rtid == 0 {
				rlib.Ulog("VacancyDetect:  No valid RTID for rentable R%08d during period %s to %s\n",
					r.RID, dt.Format(rlib.RRDATEINPFMT), dtNext.Format(rlib.RRDATEINPFMT))
				return m // this is bad! No RTID for the supplied time range
			}
		}

		rsa, err := rlib.GetRentableSpecialtyTypesForRentableByRange(r, &dt, &dtNext) // this gets an array of rentable specialties that overlap this time period
		if err != nil {
			rlib.Ulog("VacancyDetect:  Error retrieving rentable specialties for rentable R%08d during period %s to %s\n",
				r.RID, dt.Format(rlib.RRDATEINPFMT), dtNext.Format(rlib.RRDATEINPFMT))
			return m // this is bad! No RTID for the supplied time range

		}
		rentThisPeriod := rlib.CalculateGSR(dt, dtNext, xbiz.RT[rtid], rsa)

		//------------------------------------------------
		// optimization to compress consecutive days...
		//------------------------------------------------
		if k > 0 { // If the last entry's DtStop is the same time this one's DtStart...
			if m[k-1].DtStop.Equal(dt) && m[k-1].state == state { // and the umr is at the same rate...
				// m[k-1].Amount += umr * pf // add another increment to the amount
				m[k-1].DtStop = dtNext          // then we'll just adjust the end of that range to include this range too.
				m[k-1].Amount += rentThisPeriod // add the rent for this time increment
				m[k-1].comment = fmt.Sprintf("(%s - %s)", m[k-1].DtStart.Format("Jan 2"), m[k-1].DtStop.Format("Jan 2"))
				continue // Range extended.  Next!
			}
		}

		var v VacancyMarker       // ok, this is either the first entry or
		v.DtStart = dt            // it is disjoint from the last range
		v.DtStop = dtNext         // fill it out and
		v.state = state           // note the cause of the vacancy
		v.Amount = rentThisPeriod // save the rate so we don't need to look it up later
		// v.Amount = umr * pf // save the rate so we don't need to look it up later
		v.comment = fmt.Sprintf("(%s - %s)", v.DtStart.Format("Jan 2"), v.DtStop.Format("Jan 2"))
		m = append(m, v) // add the new VacancyMarker to the list
		k++
	}
	return m
}

// ProcessRentable looks for any time period for which the Rentable has
// no rent assessment during the supplied time range. If it finds any
// vacant time periods, it generates vacancy Journal entries for that period.
// The approach here is to look for all rental agreements that apply to this
// Rentable over the supplied time period. Spin through each of them and
// count the number of days that are covered.  Compare this to the total
// number of days in the period. Then generate a vacancy entry for the time
// period for which no rent was assessed.
//============================================================================================
func ProcessRentable(xbiz *rlib.XBusiness, d1, d2 *time.Time, r *rlib.Rentable) {
	m := VacancyDetect(xbiz, d1, d2, r)
	for i := 0; i < len(m); i++ {
		// the umr rate is in cost/accrualDuration. The duration of the VacancyMarkers
		// are in integral multiples of rangeIncDur.  We need to prorate the amount of
		// each entry accordingly
		var j rlib.Journal
		j.BID = xbiz.P.BID
		j.Amount = rlib.RoundToCent(m[i].Amount)
		j.Dt = m[i].DtStop.AddDate(0, 0, -1) // associated date is last day of period
		j.Type = rlib.JNLTYPEUNAS            // this is an unassociated entry
		j.RAID = 0                           // we really mean it, it is unassociated
		j.ID = r.RID                         // mark the associated Rentable
		j.Comment = m[i].comment             // this will note consecutive days for vacancy
		jid, err := rlib.InsertJournalEntry(&j)
		rlib.Errlog(err)
		if jid > 0 {
			var ja rlib.JournalAllocation
			ja.JID = jid
			ja.Amount = j.Amount
			ja.ASMID = 0 // it's unassociated
			ja.AcctRule = "c ${DFLTGSRENT} _,d ${DFLTVAC} _"
			ja.RID = r.RID
			// fmt.Printf("VACANCY: inserting journalAllocation entry: %#v\n", ja)
			rlib.InsertJournalAllocationEntry(&ja)
		}
	}
}

// GenVacancyJournals creates Journal entries that cover vacancy for
// every Rentable where the Rentable type is being managed to budget
//========================================================================================================
func GenVacancyJournals(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(xbiz.P.BID)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r rlib.Rentable
		rlib.Errcheck(rows.Scan(&r.RID, &r.BID, &r.Name, &r.AssignmentTime, &r.LastModTime, &r.LastModBy))
		ProcessRentable(xbiz, d1, d2, &r)
	}
	rlib.Errcheck(rows.Err())

}
