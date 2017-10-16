package rlib

import "time"

// The input to this routine is a list of start-end times. These time
// periods are searched to determine if all of time period t1-t2 is covered.
// It returns an array of gaps in t1-t2 that are not covered by the
// time ranges from the input list.

// https://play.golang.org/p/n9PVOVGFyI

// Period describes a span of time by specifying a start
// and end time.
type Period struct {
	D1, D2 time.Time
}

// FindGaps looks for gaps between D1 and D2 that are not covered
// by the Periods in a. It creates a map of dates from D1 to D2. The
// value at each index is boolean - true means that it is NOT covered
// by any of the periods in a, and false means that it is.
//
// INPUT
// D1,D2 = time period of interest
//     a = slice of Periods to examine for gaps between D1 and D2
//         For the primary usecase, this is the list of RentalAgreements
//         start & stop times
//
// RETURN
//     a slice of Periods that are gaps between D1 & D2
//-----------------------------------------------------------------------------
func FindGaps(D1, D2 *time.Time, a []Period) []Period {
	//---------------------------------
	// First create the map of dates
	//---------------------------------
	var b = map[string]bool{} // map["yyyy-mm-dd"]bool, ex: b["2017-10-01"]=true  means that b["2017-10-1"] is NOT covered by a
	for i := *D1; i.Before(*D2); i = i.AddDate(0, 0, 1) {
		b[i.Format(RRDATEFMTSQL)] = true // assume it is NOT covered, set to false if it is covered
	}

	//--------------------------------------------
	// Mark all dates covered by a Period in a[]
	// (for each rental agreement period)
	//--------------------------------------------
	for i := 0; i < len(a); i++ { // for each rental agreement period
		for j := a[i].D1; j.Before(a[i].D2); j = j.AddDate(0, 0, 1) { // for each day during the period
			b[j.Format(RRDATEFMTSQL)] = false // mark this date as covered
		}
	}

	//------------------------------------------------------------
	// b now looks like this (representative example):
	//
	//     b["2017-10-01"] = true
	//     b["2017-10-02"] = true
	//     b["2017-10-03"] = true
	//     ...
	//     b["2017-10-07"] = true
	//     b["2017-10-08"] = false      -- start of first gap
	//     b["2017-10-09"] = false
	//     ...
	//     b["2017-10-14"] = false      -- end of first gap
	//     b["2017-10-15"] = true
	// etc.
	//
	// Now create an array of Periods that describe the gaps...
	//------------------------------------------------------------
	var c []Period
	start := ""
	stop := ""
	for i := 0; i < len(b); i++ {
		s := D1.AddDate(0, 0, i).Format(RRDATEFMTSQL)
		if b[s] { // is this date covered by a rental agreement?
			if len(start) == 0 { // is this the start of a new Period?
				start = s // yes: mark the start time
				stop = s  //      stop & start are the same initially
			} else {
				stop = s // no: start was already set, so set the updated stop date
			}
		} else if !b[s] && len(start) > 0 { // if this date is NOT covered AND we have already started a Period...
			var p Period // ... then it is time to SET the period that we have
			p.D1, _ = StringToDate(start)
			p.D2, _ = StringToDate(stop)
			p.D2 = p.D2.AddDate(0, 0, 1) // stop date is includes all of D2
			Console("*** FIND GAPS:  D1 = %s,  D2 = %s\n", p.D1.Format(RRDATEFMTSQL), p.D2.Format(RRDATEFMTSQL))
			c = append(c, p)
			start = ""
			stop = ""
		}
		// Console("%d. b[%s] = %t\n", i, s, b[s])
	}
	if len(start) > 0 { // if we have started a range, we need to add it now
		var p Period
		p.D1, _ = StringToDate(start)
		p.D2, _ = StringToDate(stop)
		p.D2 = p.D2.AddDate(0, 0, 1) // stop date is includes all of D2
		Console("*** FIND GAPS:  D1 = %s,  D2 = %s\n", p.D1.Format(RRDATEFMTSQL), p.D2.Format(RRDATEFMTSQL))
		c = append(c, p)
	}
	return c
}

// AggregatePeriods looks for overlapping periods between D1 and D2 and
// aggregates them.
//
// INPUT
// D1,D2 = time period of interest
//     a = slice of Periods to examine for overlaps
//
// RETURN
//     a slice of Periods where overlapping Periods between D1 and D2 have
//     been aggregated.
//-----------------------------------------------------------------------------
func AggregatePeriods(D1, D2 *time.Time, a []Period) []Period {
	//------------------------------------------
	// collect the periods that overlap D1,D2
	//------------------------------------------
	var o = []Period{}
	for i := 0; i < len(a); i++ {
		if DateRangeOverlap(D1, D2, &a[i].D1, &a[i].D2) {
			o = append(o, a[i])
		}
	}

	if len(o) == 0 {
		return o
	}

	//------------------------------------------
	// combine overlapping sections
	//------------------------------------------
	var oo = []Period{o[0]}
	for i := 1; i < len(o); i++ {
		found := false
		for j := 0; j < len(oo); j++ {
			if DateRangeOverlap(&o[i].D1, &o[i].D2, &oo[j].D1, &oo[j].D2) {
				found = true
				if o[i].D1.Before(oo[j].D1) {
					oo[j].D1 = o[i].D1
				}
				if o[i].D2.After(oo[j].D2) {
					oo[j].D2 = o[i].D2
				}
			}
		}
		if !found {
			oo = append(oo, o[i])
		}
	}
	return oo
}
