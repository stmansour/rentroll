package rlib

import (
	"fmt"
	"strings"
	"time"
)

// RDateFmt is an array of date / time formats that RentRoll accepts for datetime input
var RDateFmt = []string{
	RRDATETIMEINPFMT,
	RRDATEFMT,
	RRDATEFMT2,
	RRDATEFMT3,
	RRDATEINPFMT,
}

// RentalPeriodToString takes an accrual recurrence value and returns its name as a string
func RentalPeriodToString(a int64) string {
	s := ""
	switch a {
	case ACCRUALNORECUR:
		s = "non-recurring"
	case ACCRUALSECONDLY:
		s = "secondly"
	case ACCRUALMINUTELY:
		s = "minutely"
	case ACCRUALHOURLY:
		s = "hourly"
	case ACCRUALDAILY:
		s = "daily"
	case ACCRUALWEEKLY:
		s = "weekly"
	case ACCRUALMONTHLY:
		s = "monthly"
	case ACCRUALQUARTERLY:
		s = "quarterly"
	case ACCRUALYEARLY:
		s = "yearly"
	}
	return s
}

// ProrationUnits returns a string for the supplied accrual duration value suitable for use as units
func ProrationUnits(a int64) string {
	s := ""
	switch a {
	case ACCRUALNORECUR:
		s = "!!nonrecur!!"
	case ACCRUALSECONDLY:
		s = "seconds"
	case ACCRUALMINUTELY:
		s = "minutes"
	case ACCRUALHOURLY:
		s = "hours"
	case ACCRUALDAILY:
		s = "days"
	case ACCRUALWEEKLY:
		s = "weeks"
	case ACCRUALMONTHLY:
		s = "months"
	case ACCRUALQUARTERLY:
		s = "quarters"
	case ACCRUALYEARLY:
		s = "years"
	}
	return s
}

// StringToDate tries to convert the supplied string to a time.Time value. It will use the two
// formats called out in dbtypes.go:  RRDATEFMT, RRDATEINPFMT, RRDATEINPFMT2
func s2d(s string) time.Time {
	var t time.Time
	var err error
	// try the ansi std date format first
	s = strings.TrimSpace(s)
	for i := 0; i < len(RDateFmt); i++ {
		t, err = time.Parse(RDateFmt[i], s)
		if err == nil {
			return t
		}
	}
	return t
}

// CycleDuration returns the prorateDuration in microseconds and the units as a string
func CycleDuration(prorateMethod int64, dt time.Time) time.Duration {
	var prorateDur time.Duration
	month := dt.Month()
	year := dt.Year()
	day := dt.Day()
	base := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	switch prorateMethod { // if the prorate method is less than a day, select a different duration
	case ACCRUALSECONDLY:
		prorateDur = time.Second // use seconds
	case ACCRUALMINUTELY:
		prorateDur = time.Minute //use minutes
	case ACCRUALHOURLY:
		prorateDur = time.Hour //use hours
	case ACCRUALDAILY:
		prorateDur = time.Hour * 24 // assume that proration will be by day -- even if the accrual is by weeks, months, quarters, or years
	case ACCRUALWEEKLY:
		prorateDur = time.Hour * 24 * 7 // weeks
	case ACCRUALMONTHLY:
		target := base.AddDate(0, 1, 0)
		prorateDur = target.Sub(base) // months
	case ACCRUALQUARTERLY:
		target := base.AddDate(0, 3, 0)
		prorateDur = target.Sub(base) // months
	case ACCRUALYEARLY:
		target := base.AddDate(1, 0, 0)
		prorateDur = target.Sub(base) // months
	}
	return prorateDur
}

// GetProrationRange returns the duration appropriate for the provided anchor dates, Accrual Rate, and Proration Rate
func GetProrationRange(d1, d2 time.Time, RentCycle, Prorate int64) time.Duration {
	var timerange time.Duration
	accrueDur := CycleDuration(RentCycle, d1)

	// we use d1 as the anchor point
	switch RentCycle {
	case ACCRUALSECONDLY:
		fallthrough
	case ACCRUALMINUTELY:
		fallthrough
	case ACCRUALHOURLY:
		fallthrough
	case ACCRUALDAILY:
		fallthrough
	case ACCRUALWEEKLY:
		timerange = accrueDur
	case ACCRUALMONTHLY:
		timerange = d1.AddDate(0, 1, 0).Sub(d1)
	case ACCRUALQUARTERLY:
		timerange = d1.AddDate(0, 3, 0).Sub(d1)
	case ACCRUALYEARLY:
		timerange = d1.AddDate(1, 0, 0).Sub(d1)
	}

	return timerange
}

// SelectRentableStatusForPeriod returns a subset of Rentable states that overlap the supplied range.
func SelectRentableStatusForPeriod(rsa *[]RentableStatus, dt1, dt2 time.Time) []RentableStatus {
	var m []RentableStatus
	for i := 0; i < len(*rsa); i++ {
		if DateRangeOverlap(&(*rsa)[i].DtStart, &(*rsa)[i].DtStop, &dt1, &dt2) {
			var rs RentableStatus
			rs = (*rsa)[i]
			m = append(m, rs)
		}
	}
	return m
}

// GetRentableStateForDate returns the status of the Rentable on the supplied date
func GetRentableStateForDate(rid int64, dt *time.Time) int64 {
	status := int64(RENTABLESTATUSUNKNOWN)
	d2 := dt.Add(24 * time.Hour)
	m := GetRentableStatusByRange(rid, dt, &d2)
	if len(m) > 0 {
		status = m[0].Status
	}
	return status
}

// GetRentCycleAndProration returns the RentCycle (and Proration) to use for the supplied rentable and date.
// If the override RentCycle is set for this time period, it is returned. Otherwise, the RentCycle for this
// Rentable's RentableType is returned
// Returns:
//		RentCycle
//		Proration
//		rtid for the supplied date
//		error
func GetRentCycleAndProration(r *Rentable, dt *time.Time, xbiz *XBusiness) (int64, int64, int64, error) {
	var err error
	var rc, pro, rtid int64

	rrt := GetRentableTypeRefForDate(r.RID, dt)
	if rrt.RID == 0 {
		return rc, pro, rtid, fmt.Errorf("No RentableTypeRef for %s", dt.Format(RRDATEINPFMT))
	}
	rtid = GetRTIDForDate(r.RID, dt)
	if rrt.RentCycle > ACCRUALNORECUR { // if there's an override for RentCycle...
		rc = rrt.RentCycle // ...set it
	} else {
		rc = xbiz.RT[rtid].RentCycle
	}
	if rrt.Proration > ACCRUALNORECUR { // if there's an override for Propration...
		pro = rrt.Proration // ...set it
	} else {
		pro = xbiz.RT[rtid].Proration
	}

	// we need to load the RentableType for RentCycle or Proration or both...
	return rc, pro, rtid, err
}

// Prorate computes basic info to perform rent proration:
// examples:
//   DTSTART      DTSTOP       D1           D2         ACCRUAL  PRORATION   ASMTDUR  RENTDUR    PF      ANALYZE: START - STOP
//   2004-01-01   2015-11-08   2015-11-01   2015-12-01    6          4        30        8      0.2667   2015-11-01 - 2015-11-09
//   2015-11-21   2016-11-21   2015-11-01   2015-12-01    6          4        30        10     0.3333   2015-11-21 - 2015-12-01
//   2015-11-21   2016-11-21   2015-11-01   2015-12-01    0          0        30        30     1.0000   2015-11-21 - 2015-12-01
//
// Parameters:
//  	Start,Stop:     rental agreement period covering the Rentable
//  	d1, d2:         time period the Rentable was rented
//  	accrual:        rent cycle
//  	prorateMethod:  method (usually the recur frequency) used to calculate proration
//
// Returns:
//      asmtdur = rent cycle
//      rentdur = duration actually rented
//      pf      = proration factor, multiply rent/proratcycle * (prorate cycles) to get the prorated rent.
// ----------------------------------------------------------------------------------------------------------
func Prorate(RAStart, RAStop, asmtStart, asmtStop time.Time, accrual, prorateMethod int64) (int64, int64, float64) {
	var asmtDur int64
	var rentDur int64
	var pf float64

	prorateDur := CycleDuration(prorateMethod, asmtStart)
	//-------------------------------------------------------------------
	// Scope the Rental Agreement period down to this assessment period.
	// Overlap the Rental Agreement period (RAStart to RAStop) with the
	// assessment period (asmtStart - asmtStop)
	//-------------------------------------------------------------------
	start := asmtStart
	if RAStart.After(start) {
		start = RAStart
	}
	stop := RAStop.Add(prorateDur)
	if stop.After(asmtStop) {
		stop = asmtStop
	}

	// fmt.Printf("scoped period:  %s - %s\n", start.Format(RRDATEINPFMT), stop.Format(RRDATEINPFMT))
	asmtDur = int64(asmtStop.Sub(asmtStart) / prorateDur)
	rentDur = int64(stop.Sub(start) / prorateDur)

	// fmt.Printf("rentDur = %d %s\n", rentDur, units)
	// fmt.Printf("asmtDur = %d %s\n", asmtDur, units)

	if ACCRUALNORECUR == prorateMethod {
		pf = 1.0
	} else {
		pf = float64(rentDur) / float64(asmtDur)
	}

	return asmtDur, rentDur, pf
}
