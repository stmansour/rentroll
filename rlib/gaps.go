package rlib

import "time"

// The input to this routine is a list of start-end times. These time
// periods are searched to determine if all of time period t1-t2 is covered.
// It returns an array of gaps in t1-t2 that are not covered by the
// time ranges from the input list.

// FindGaps looks for gaps between D1 and D2 that are not covered
// by the Periods in a.
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
	// Console("Entered FindGaps\n")
	var gaps []Period
	Start := *D1
	GreatestD2 := *D1
	n := TimeListAggregate(a)
	for i := 0; i < len(n); i++ {
		if n[i].D1.After(Start) {
			// Console("Start:  %s prior to %s.  Adding Gap\n", Start.Format(RRDATEFMTSQL), n[i].D1.Format(RRDATEFMTSQL))
			var x Period
			x.D1 = Start
			x.D2 = n[i].D1
			gaps = append(gaps, x)
		}
		Start = n[i].D2
		if Start.After(GreatestD2) {
			GreatestD2 = Start
		}
		// Console("updating start time to %s\n", Start.Format(RRDATEFMTSQL))
	}
	if GreatestD2.Before(*D2) {
		var x Period
		x.D1 = GreatestD2
		x.D2 = *D2
		gaps = append(gaps, x)
	}
	return gaps
}

// TimeListAggregate merges x into timelist m and returns the
// aggregated list.
// INPUT
//   m = list of existing times
//
// RETURNS
//   m     = the new list with x aggregated
//-----------------------------------------------------------------------------
func TimeListAggregate(m []Period) []Period {
	var n []Period    // the newly aggregated list
	var aggrCount int // count of aggregations

	l := len(m)
	for i := 0; i < l; i++ {
		m[i].Checked = false
	}
	for i := 0; i < l; i++ {
		overlap := false
		for j := i + 1; j < l; j++ {
			if j == i || m[j].Checked {
				continue
			}
			if PeriodOverlap(&m[i], &m[j]) {
				x := AggrPeriods(&m[i], &m[j])
				m[j].Checked = true
				n = append(n, x)
				// Console("Aggregated: %d and %d to %s - %s\n", i, j, x.D1.Format("2006-01-02"), x.D2.Format("2006-01-02"))
				aggrCount++ // another aggregation
				overlap = true
			}
		}
		if !overlap && !m[i].Checked {
			// Console("overlap = false for i = %d.  appending m[%d]\n", i, i)
			n = append(n, m[i])
		}
	}
	if aggrCount > 0 {
		// Console("Aggs = %d, recursing\n", aggrCount)
		for k := 0; k < len(n); k++ {
			// Console("k[%d] = %s - %s\n", k, n[k].D1.Format(RRDATEFMTSQL), n[k].D2.Format(RRDATEFMTSQL))
		}
		n = TimeListAggregate(n)
	}
	return n
}

// AggrPeriods merges two overlapping periods into a single period
// INPUT
//   p1,p2 = periods to aggregate
//
// RETURNS
//   p     = the aggregated period
//-----------------------------------------------------------------------------
func AggrPeriods(p1, p2 *Period) Period {
	p := *p1
	if p2.D1.Before(p.D1) {
		p.D1 = p2.D1
	}
	if p2.D2.After(p.D2) {
		p.D2 = p2.D2
	}
	return p
}
