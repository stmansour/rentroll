package rlib

import (
	"fmt"
	"time"
)

// GetRecurrences is a shorthand for assessment variables to get a list
// of dates on which charges must be assessed for a particular interval of time (d1 - d2)
func (a *Assessment) GetRecurrences(d1, d2 *time.Time) []time.Time {
	return GetRecurrences(d1, d2, &a.Start, &a.Stop, a.RentCycle)
}

// DateInRange returns true if dt is >= start AND db < stop, otherwise it returns false
func DateInRange(dt, start, stop *time.Time) bool {
	// fmt.Printf("dt: %s\n", dt.Format(time.RFC3339))
	// fmt.Printf("start: %s   stop: %s\n", start.Format(time.RFC3339), stop.Format(time.RFC3339))
	return (dt.Equal(*start) || dt.After(*start)) && dt.Before(*stop)
}

// DateRangeOverlap returns true if time ranges a1-a2 overlaps timerange s1-s2, otherwise it returns false
func DateRangeOverlap(a1, a2, s1, s2 *time.Time) bool {
	sse := s1.Equal(*s2)
	aae := a1.Equal(*a2)
	ase := a1.Equal(*s1)

	if sse && aae && ase { // single point in time, all equal
		return true
	}

	if sse { // s = point, a = range
		return (s1.Equal(*a1) || s1.After(*a1)) && s1.Before(*a2)
	}

	if aae { // a = point, s = range
		return (a1.Equal(*s1) || a1.After(*s1)) && a1.Before(*s2)
	}

	return a1.Before(*s2) && a2.After(*s1)
}

// PeriodOverlap returns true of the two periods overlap, false otherwise.
func PeriodOverlap(p1, p2 *Period) bool {
	return DateRangeOverlap(&p1.D1, &p1.D2, &p2.D1, &p2.D2)
}

// ContainDateRange returns a time range t1,t2 as follows:
// t1,t2 are initially set to the supplied date range d1,d2
//     If the initial boundary begins after d1 AND it is before b2
//     then t1 is snapped to b1
//     If the ending boundary ends before d2 AND iti is after b1
//     then t2 is snapped to b2.
// The net effect is to return a time range
//
// INPUTS
//  d1,d2 - initial date range
//  b1,b2 - bounds for the initial date range
//
// RETURNS
//  modified date range t1,t2
//  any errors encountered or nil if no error
//-----------------------------------------------------------------------------
func ContainDateRange(d1, d2, b1, b2 *time.Time) (time.Time, time.Time, error) {
	t1 := *d1
	t2 := *d2
	if b2.Before(*b1) || b1.After(*b2) {
		err := fmt.Errorf("ContainDateRange: invalid boundaries: %s,%s", b1.Format(RRDATEFMTSQL), b2.Format(RRDATEFMTSQL))
		return t1, t2, err
	}
	if b1.After(*d1) && b1.Before(*b2) {
		t1 = *b1
	}
	if b2.Before(*d2) && b2.After(*b1) {
		t2 = *b2
	}
	return t1, t2, nil
}

// a1 - a2 = time range of the assessment
// R1 - R2 = time range for the run calculation
// freq = chunk of time over which to quantize the assessment
func genRegularRecurSeq(a1, a2, R1, R2 *time.Time, freq time.Duration) []time.Time {
	// fmt.Printf("\n---------------------\n")
	// fmt.Printf("a1: %s   a2: %s\n", a1.Format(time.RFC3339), a2.Format(time.RFC3339))
	// fmt.Printf("R1: %s   R2: %s\n", R1.Format(time.RFC3339), R2.Format(time.RFC3339))

	//============================================
	// Set up first time range for first run...
	//============================================
	d1 := *R1
	d2 := d1.Add(freq)

	//============================================
	// Now iterate in chunks (of size: freq) and
	// save the recurrence dates...
	//============================================
	var m []time.Time
	for i := 0; i < 10000; i++ {
		// fmt.Printf("iter: %d\n", i)
		//----------------------------------------
		// don't go outside the requested range
		//----------------------------------------
		if d1.After(*R2) || d1.Equal(*R2) {
			break
		}
		if d2.After(*R2) {
			d2 = *R2
		}
		//----------------------------------------
		//  does a1 - a2 overlap d1-d2
		//----------------------------------------
		// fmt.Printf("d1: %s   d2: %s\n", d1.Format(time.RFC3339), d2.Format(time.RFC3339))
		if !(a1.After(d2) || a1.Equal(d2) || a2.Before(d1) || a2.Equal(d1)) {
			d := time.Date(d1.Year(), d1.Month(), d1.Day(), a1.Hour(), a1.Minute(), a1.Second(), 0, time.UTC)
			// fmt.Printf("STORE d = %s\n", d.Format(time.RFC3339))
			m = append(m, d)
		}
		//----------------------------------------
		//  set the next interval to check...
		//----------------------------------------
		d1 = d1.Add(freq)
		d2 = d2.Add(freq)
	}
	return m
}

func genMonthlyRecurSeq(a1, a2, R1, R2 *time.Time, nMonths int64) []time.Time {

	//============================================
	// Set up first time range for first run...
	//============================================
	d1 := *R1
	d2 := d1 // just to define it, it will be set correctly below

	//============================================
	// Now iterate in chunks (of size: nMonths) and
	// save the recurrence dates...
	//============================================
	var m []time.Time
	for i := 0; i < 10000; i++ {
		mo, y := IncMonths(d1.Month(), nMonths)
		d2 = time.Date(d1.Year()+int(y), mo, d1.Day(), d1.Hour(), d1.Minute(), d1.Second(), 0, time.UTC)

		//----------------------------------------
		// don't go outside the requested range
		//----------------------------------------
		if d1.After(*R2) || d1.Equal(*R2) {
			break
		}
		if d2.After(*R2) {
			d2 = *R2
		}
		//-------------------------------------------------------------------------------------
		//  the recurrence date will be the epoch date applied to d1.  Then see if this date
		//  is in the range d1 - d2
		//-------------------------------------------------------------------------------------
		d := time.Date(d1.Year(), d1.Month(), a1.Day(), d1.Hour(), d1.Minute(), d1.Second(), 0, time.UTC)

		// for the date comparison, we are only interested whether the dates fall in line, to the times...
		dt1 := time.Date(d1.Year(), d1.Month(), d1.Day(), 0, 0, 0, 0, time.UTC)
		dt2 := time.Date(d2.Year(), d2.Month(), d2.Day(), 0, 0, 0, 0, time.UTC)
		at1 := time.Date(a1.Year(), a1.Month(), a1.Day(), 0, 0, 0, 0, time.UTC)
		at2 := time.Date(a2.Year(), a2.Month(), a2.Day(), 0, 0, 0, 0, time.UTC)

		//-------------------------------------------------------------------------------------
		// make sure it's in the interval range AND in its active timeframe, and if it is
		// then add it to the list
		//-------------------------------------------------------------------------------------
		// fmt.Printf("genMonthlyRecurSeq: d = %s, d1,d2 = %s , %s    a1,a2 = %s , %s\n", d.Format(RRDATEREPORTFMT), d1.Format(RRDATEREPORTFMT), d2.Format(RRDATEREPORTFMT), a1.Format(RRDATEREPORTFMT), a2.Format(RRDATEREPORTFMT))
		if DateInRange(&d, &dt1, &dt2) {
			if a1.Equal(d) || DateInRange(&d, &at1, &at2) {
				m = append(m, d)
			}
		}
		d1 = d2 // on to the next interval
	}

	return m
}

func genYearlyRecurSeq(d, start, stop *time.Time, n int64) []time.Time {
	var m []time.Time
	dt := *d
	for DateInRange(&dt, start, stop) {
		m = append(m, dt)
		dt = time.Date(dt.Year()+int(n), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), 0, time.UTC)
	}
	return m
}

// GetRecurrences returns a list of instance dates where an event time
// (aStart - aStop) overlaps with an interval time (start - stop).  The
// recurrence frequency maps to those that can happen for an assessment.
//
// INPUTS
//     start, stop = time range for the recurrences to be generated. For
//                   example, the start / end time of a period to be
//                   considered.
//   aStart, aStop = imposed bounds, for example the start/stop time of
//                   an Assessment
//       cycleFreq = recurrence frequency
//
// RETURNS
//   an array of instances
//-----------------------------------------------------------------------------
func GetRecurrences(start, stop, aStart, aStop *time.Time, cycleFreq int64) []time.Time {
	var m []time.Time
	//-------------------------------------------
	// first ensure that the data is not bad...
	//-------------------------------------------
	if aStart.After(*aStop) {
		return m
	}

	//----------------------------------------------------------------
	// next, ensure that the assessment falls in the time range...
	//----------------------------------------------------------------
	if cycleFreq > RECURNONE && (aStop.Equal(*start) || aStop.Before(*start) || aStart.After(*stop) || aStart.Equal(*stop)) {
		return m
	}

	//-------------------------------------------
	// now calculate all the recurrences
	//-------------------------------------------
	switch cycleFreq {
	case RECURNONE: // no recurrence
		if DateInRange(aStart, start, stop) {
			m = append(m, *aStart)
			return m
		}
	case RECURDAILY: // daily
		d := start.Day()
		if start.Day() < aStart.Day() {
			d = aStart.Day()
		}
		dt := time.Date(start.Year(), start.Month(), d, aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genRegularRecurSeq(&dt, aStop, start, stop, 24*time.Hour)
	case RECURWEEKLY: // weekly
		dt := time.Date(start.Year(), start.Month(), start.Day(), aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genRegularRecurSeq(&dt, aStop, start, stop, 7*24*time.Hour)
	case RECURMONTHLY: // monthly
		// dt := time.Date(start.Year(), start.Month(), aStart.Day(), aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genMonthlyRecurSeq(aStart, aStop, start, stop, 1)
	case RECURQUARTERLY: // quarterly
		dt := time.Date(start.Year(), start.Month(), aStart.Day(), aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genMonthlyRecurSeq(&dt, aStop, start, stop, 3)
	case RECURYEARLY: // yearly
		dt := time.Date(start.Year(), aStart.Month(), aStart.Day(), aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genYearlyRecurSeq(&dt, start, stop, 1)
	}
	return m
}
