package rlib

import (
	"strings"
	"time"
)

var recurUnits = []struct {
	Name  string
	Adj   string
	Index int64
}{
	{"second", "secondly", RECURSECONDLY},
	{"minute", "minutely", RECURMINUTELY},
	{"hour", "hourly", RECURHOURLY},
	{"day", "daily", RECURDAILY},
	{"week", "weekly", RECURWEEKLY},
	{"month", "monthly", RECURMONTHLY},
	{"quarter", "quarterly", RECURQUARTERLY},
	{"year", "yearly", RECURYEARLY},
}

// RentalPeriodToString takes an accrual recurrence value and returns its
// name as a string
//=============================================================================
func RentalPeriodToString(a int64) string {
	if a == int64(0) {
		return "non-recurring"
	}
	for i := 0; i < len(recurUnits); i++ {
		if a == recurUnits[i].Index {
			return recurUnits[i].Adj
		}
	}
	return ""
}

// ProrationUnits returns a string for the supplied accrual duration value
// suitable for use as units
//=============================================================================
func ProrationUnits(a int64) string {
	if a == int64(0) {
		return "non-recurring"
	}
	for i := 0; i < len(recurUnits); i++ {
		if a == recurUnits[i].Index {
			return recurUnits[i].Name + "s"
		}
	}
	return ""
}

// RecurUnitsStringToIndex returns the index for the recur type that matches
// the supplied string.  Since it matches only the core of the units name,
// it can be suffixed with either "s" or "ly"
//=============================================================================
func RecurUnitsStringToIndex(a string) int64 {
	for i := 0; i < len(recurUnits); i++ {
		if strings.Contains(a, recurUnits[i].Name) || strings.Contains(a, recurUnits[i].Adj) {
			return recurUnits[i].Index
		}
	}
	return RECURNONE
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

// InstanceDateCoveringDate calculates the start date (the instance date)
// of the recurring period in which the supplied target date falls.
//
// Example why this is needed.  Suppose that a Rental Agreement is amended
// on August 28, 2018.  There is a rentable in this Rental Agreement that
// has a monthly assessment. This business defines that monthly assessments are
// due on the 1st of the month. So, we need to prorate the assessment for the
// Aug 1 thru Aug 28 for the original RA.  So we will need to change the
// assessment instance that covers Aug 28, 2018. To do that, we call this
// function. Typically, we already have the recurring assessment definition
// or some instance of it, and we also have the new RentalAgreement which
// contains the target date we need.
//
//     var asm Assessment
//     var ra RentalAgreement
//     ...
//     instanceStart := InstanceDateCoveringDate(&asm.Start, &ra.RentStart, &asm.RentCycle)
//
// Now, instanceStart is set to the assessment instance date we will need to
// modify.
//
// GoPlayground: https://play.golang.org/p/hftEyrCVe7w
//
// INPUTS
//  epoch - an actual instance date/time, can be the epoch or any instance
//  t     - target date/time -- find the instance start date that covers t
//  cycle - recurrence cycle
//
// RETURNS
//  next instance date/time
//-----------------------------------------------------------------------------
func InstanceDateCoveringDate(epoch, t *time.Time, cycle int64) time.Time {
	var x time.Time
	switch cycle {
	case RECURSECONDLY:
		x = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), epoch.Nanosecond(), epoch.Location())
		if x.After(*t) {
			x = x.Add(-1 * time.Second)
		}
		return x
	case RECURMINUTELY:
		x = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
		if x.After(*t) {
			x = x.Add(-1 * time.Minute)
		}
		return x
	case RECURHOURLY:
		x = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
		if x.After(*t) {
			x = x.Add(-1 * time.Hour)
		}
		return x
	case RECURDAILY:
		x = time.Date(t.Year(), t.Month(), t.Day(), epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
		if x.After(*t) {
			x = x.Add(-24 * time.Hour)
		}
		return x
	case RECURWEEKLY:
		a := epoch.Weekday() // this is the day of the week we want x to fall on
		x = time.Date(t.Year(), t.Month(), t.Day(), epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
		for i := 0; i < 7 && x.Weekday() != a; i++ { // i just guarantees that we won't loop forever
			x = x.Add(-24 * time.Hour)
		}
		return x
	}

	d1 := epoch.Day()
	day := d1
	if d1 > 28 {
		d1 = 28
	}

	//qoff := int(epoch.Month()-1) % 3 // offset within the quarter By Lina
	qoff := int(epoch.Month()) % 3 // offset within the quarter
	dt := time.Date(t.Year(), t.Month(), d1, epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
	if cycle == RECURMONTHLY {
		if dt.After(*t) {
			dt = dt.AddDate(0, -1, 0)
		}
	}
	if cycle == RECURQUARTERLY {
		for i := 0; i < 3 && qoff != (int(dt.Month())%3); i++ { // i guarantees that we won't loop forever
			dt = dt.AddDate(0, -1, 0)
		}
		if dt.After(*t) { //By Lina
			dt = dt.AddDate(0, -3, 0)
		}
	}
	if cycle == RECURYEARLY { //By Lina

		if epoch.Month() != time.February || epoch.Day() != 29 {
			dt = time.Date(t.Year(), epoch.Month(), epoch.Day(), epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
		} else { ////handle the situations like "2/29/2020"
			dayofFeb := LastDOM(epoch.Month(), t.Year())
			dt = time.Date(t.Year(), epoch.Month(), dayofFeb, epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
		}
		if dt.After(*t) {
			dt = dt.AddDate(-1, 0, 0)
		}
	}
	if day >= 28 { // snap to the last day of this month...
		day = LastDOM(dt.Month(), dt.Year())
		dt = time.Date(dt.Year(), dt.Month(), day, dt.Hour(), dt.Minute(), dt.Second(), dt.Nanosecond(), dt.Location())
	}
	return dt
}

// IsInstanceDate returns true if the supplied date is an instance date of the
// supplied epoch and cycle frequency.
//
// INPUTS:
//    epoch - any instance date of the recurring series
//    d     - the date in question
//    cycle - repeating cycle for the the recurring series
//    proration - repeating cycle for calculating proration
//
// RETURN:
//    true means d is an instance of the recurring series
//    false means it is not
//------------------------------------------------------------------------------
func IsInstanceDate(epoch, d *time.Time, cycle, proration int64) bool {
	// Console("Entered IsInstanceDate:  epoch = %s, d = %s, cycle = %d\n", epoch.Format(RRDATEREPORTFMT), d.Format(RRDATEREPORTFMT), cycle)
	ok, d2 := GetEpochFromBaseDate(*epoch, *d, ENDOFTIME, cycle)
	if !ok { // should never happen
		Ulog("IsInstanceDate received !ok unexpectedly\n")
		// Console("IsInstanceDate received !ok unexpectedly\n")
	}
	// Console("d2 determined as %s\n", d2.Format(RRDATEREPORTFMT))
	dur := CycleDuration(proration, *epoch)
	sep := d2.Sub(*d)
	if sep < 0 {
		sep = -sep
	}
	// Console("sep = %d, return value = %t\n", sep, sep < dur)
	return sep < dur
}
