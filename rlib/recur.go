package rlib

import "time"

// RentalPeriodToString takes an accrual recurrence value and returns its
// name as a string
//=============================================================================
func RentalPeriodToString(a int64) string {
	s := ""
	switch a {
	case RECURNONE:
		s = "non-recurring"
	case RECURSECONDLY:
		s = "secondly"
	case RECURMINUTELY:
		s = "minutely"
	case RECURHOURLY:
		s = "hourly"
	case RECURDAILY:
		s = "daily"
	case RECURWEEKLY:
		s = "weekly"
	case RECURMONTHLY:
		s = "monthly"
	case RECURQUARTERLY:
		s = "quarterly"
	case RECURYEARLY:
		s = "yearly"
	}
	return s
}

// ProrationUnits returns a string for the supplied accrual duration value
// suitable for use as units
//=============================================================================
func ProrationUnits(a int64) string {
	s := ""
	switch a {
	case RECURNONE:
		s = "!!nonrecur!!"
	case RECURSECONDLY:
		s = "seconds"
	case RECURMINUTELY:
		s = "minutes"
	case RECURHOURLY:
		s = "hours"
	case RECURDAILY:
		s = "days"
	case RECURWEEKLY:
		s = "weeks"
	case RECURMONTHLY:
		s = "months"
	case RECURQUARTERLY:
		s = "quarters"
	case RECURYEARLY:
		s = "years"
	}
	return s
}

// CycleDuration returns the prorateDuration in microseconds and the units as
// a string
//=============================================================================
func CycleDuration(cycle int64, epoch time.Time) time.Duration {
	var cycleDur time.Duration
	month := epoch.Month()
	year := epoch.Year()
	day := epoch.Day()
	base := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	switch cycle { // if the prorate method is less than a day, select a different duration
	case RECURSECONDLY:
		cycleDur = time.Second // use seconds
	case RECURMINUTELY:
		cycleDur = time.Minute //use minutes
	case RECURHOURLY:
		cycleDur = time.Hour //use hours
	case RECURDAILY:
		cycleDur = time.Hour * 24 // assume that proration will be by day -- even if the accrual is by weeks, months, quarters, or years
	case RECURWEEKLY:
		cycleDur = time.Hour * 24 * 7 // weeks
	case RECURMONTHLY:
		target := base.AddDate(0, 1, 0)
		cycleDur = target.Sub(base) // months
	case RECURQUARTERLY:
		target := base.AddDate(0, 3, 0)
		cycleDur = target.Sub(base) // months
	case RECURYEARLY:
		target := base.AddDate(1, 0, 0)
		cycleDur = target.Sub(base) // months
	}
	return cycleDur
}

// GetProrationRange returns the duration appropriate for the provided anchor
// dates, Accrual Rate, and Proration Rate
//=============================================================================
func GetProrationRange(d1, d2 time.Time, RentCycle, Prorate int64) time.Duration {
	var timerange time.Duration
	accrueDur := CycleDuration(RentCycle, d1)

	// we use d1 as the anchor point
	switch RentCycle {
	case RECURSECONDLY:
		fallthrough
	case RECURMINUTELY:
		fallthrough
	case RECURHOURLY:
		fallthrough
	case RECURDAILY:
		fallthrough
	case RECURWEEKLY:
		timerange = accrueDur
	case RECURMONTHLY:
		timerange = d1.AddDate(0, 1, 0).Sub(d1)
	case RECURQUARTERLY:
		timerange = d1.AddDate(0, 3, 0).Sub(d1)
	case RECURYEARLY:
		timerange = d1.AddDate(1, 0, 0).Sub(d1)
	}

	return timerange
}

// GetPreviousInstance calculates the previous instance based on the supplied instance
// datetime and the recur cycle.
//
// INPUTS
//  d     - current instance date/time
//  cycle - recurrence cycle
//
// RETURNS
//  previous instance date/time
//-----------------------------------------------------------------------------
func GetPreviousInstance(d time.Time, cycle int64) time.Time {
	months := 0
	switch cycle {
	case RECURSECONDLY:
		fallthrough
	case RECURMINUTELY:
		fallthrough
	case RECURHOURLY:
		fallthrough
	case RECURDAILY:
		fallthrough
	case RECURWEEKLY:
		dur := CycleDuration(cycle, d)
		return d.Add(-dur)
	case RECURMONTHLY:
		months = 1
	case RECURQUARTERLY:
		months = 3
	case RECURYEARLY:
		months = 12
	}
	day := d.Day()
	d1 := day
	if d1 > 28 {
		d1 = 28
	}
	dt := time.Date(d.Year(), d.Month(), d1, d.Hour(), d.Minute(), d.Second(), d.Nanosecond(), time.UTC)
	prev := dt.AddDate(0, -months, 0)
	if day > 28 { // snap to the last day of this month...
		day = LastDOM(prev.Month(), prev.Year())
		prev = time.Date(prev.Year(), prev.Month(), day, prev.Hour(), prev.Minute(), prev.Second(), prev.Nanosecond(), time.UTC)
	}
	return prev
}

// GetNextInstance calculates the next instance based on the supplied instance
// datetime and the recur cycle.
//
// INPUTS
//  d     - current instance date/time
//  cycle - recurrence cycle
//
// RETURNS
//  next instance date/time
//-----------------------------------------------------------------------------
func GetNextInstance(d time.Time, cycle int64) time.Time {
	months := 0
	switch cycle {
	case RECURSECONDLY:
		fallthrough
	case RECURMINUTELY:
		fallthrough
	case RECURHOURLY:
		fallthrough
	case RECURDAILY:
		fallthrough
	case RECURWEEKLY:
		dur := CycleDuration(cycle, d)
		return d.Add(dur)
	case RECURMONTHLY:
		months = 1
	case RECURQUARTERLY:
		months = 3
	case RECURYEARLY:
		months = 12
	}
	day := d.Day()
	d1 := day
	if d1 > 28 {
		d1 = 28
	}
	dt := time.Date(d.Year(), d.Month(), d1, d.Hour(), d.Minute(), d.Second(), d.Nanosecond(), time.UTC)
	next := dt.AddDate(0, months, 0)
	if day > 28 { // snap to the last day of this month...
		day = LastDOM(next.Month(), next.Year())
		next = time.Date(next.Year(), next.Month(), day, next.Hour(), next.Minute(), next.Second(), next.Nanosecond(), time.UTC)
	}
	return next
}

// TODO: see about replacing NextPeriod with GetNextInstance

// NextPeriod computes the next period start given the current period start
// and the recur cycle
//
// INPUTS:
//  t     - curren start time
//  cycle - 0 = norecur, 1 = secondly, ... 7 = Yearly
//
// RETURNS:
//  next instance start time.
//---------------------------------------------------------------------------
func NextPeriod(t *time.Time, cycle int64) time.Time {
	var ret time.Time
	switch cycle { // if the prorate method is less than a day, select a different duration
	case RECURNONE:
		ret = *t
	case RECURSECONDLY:
		ret = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()+1, t.Nanosecond(), t.Location())
	case RECURMINUTELY:
		ret = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+1, t.Second(), t.Nanosecond(), t.Location())
	case RECURHOURLY:
		ret = time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+1, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	case RECURDAILY:
		ret = t.AddDate(0, 0, 1)
	case RECURWEEKLY:
		ret = t.AddDate(0, 0, 7)
	case RECURMONTHLY:
		ret = t.AddDate(0, 1, 0)
	case RECURQUARTERLY:
		ret = t.AddDate(0, 3, 0)
	case RECURYEARLY:
		ret = t.AddDate(1, 0, 0)
	}
	return ret
}
