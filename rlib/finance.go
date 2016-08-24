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
	case CYCLENORECUR:
		s = "non-recurring"
	case CYCLESECONDLY:
		s = "secondly"
	case CYCLEMINUTELY:
		s = "minutely"
	case CYCLEHOURLY:
		s = "hourly"
	case CYCLEDAILY:
		s = "daily"
	case CYCLEWEEKLY:
		s = "weekly"
	case CYCLEMONTHLY:
		s = "monthly"
	case CYCLEQUARTERLY:
		s = "quarterly"
	case CYCLEYEARLY:
		s = "yearly"
	}
	return s
}

// ProrationUnits returns a string for the supplied accrual duration value suitable for use as units
func ProrationUnits(a int64) string {
	s := ""
	switch a {
	case CYCLENORECUR:
		s = "!!nonrecur!!"
	case CYCLESECONDLY:
		s = "seconds"
	case CYCLEMINUTELY:
		s = "minutes"
	case CYCLEHOURLY:
		s = "hours"
	case CYCLEDAILY:
		s = "days"
	case CYCLEWEEKLY:
		s = "weeks"
	case CYCLEMONTHLY:
		s = "months"
	case CYCLEQUARTERLY:
		s = "quarters"
	case CYCLEYEARLY:
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
func CycleDuration(cycle int64, epoch time.Time) time.Duration {
	var cycleDur time.Duration
	month := epoch.Month()
	year := epoch.Year()
	day := epoch.Day()
	base := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	switch cycle { // if the prorate method is less than a day, select a different duration
	case CYCLESECONDLY:
		cycleDur = time.Second // use seconds
	case CYCLEMINUTELY:
		cycleDur = time.Minute //use minutes
	case CYCLEHOURLY:
		cycleDur = time.Hour //use hours
	case CYCLEDAILY:
		cycleDur = time.Hour * 24 // assume that proration will be by day -- even if the accrual is by weeks, months, quarters, or years
	case CYCLEWEEKLY:
		cycleDur = time.Hour * 24 * 7 // weeks
	case CYCLEMONTHLY:
		target := base.AddDate(0, 1, 0)
		cycleDur = target.Sub(base) // months
	case CYCLEQUARTERLY:
		target := base.AddDate(0, 3, 0)
		cycleDur = target.Sub(base) // months
	case CYCLEYEARLY:
		target := base.AddDate(1, 0, 0)
		cycleDur = target.Sub(base) // months
	}
	return cycleDur
}

// GetProrationRange returns the duration appropriate for the provided anchor dates, Accrual Rate, and Proration Rate
func GetProrationRange(d1, d2 time.Time, RentCycle, Prorate int64) time.Duration {
	var timerange time.Duration
	accrueDur := CycleDuration(RentCycle, d1)

	// we use d1 as the anchor point
	switch RentCycle {
	case CYCLESECONDLY:
		fallthrough
	case CYCLEMINUTELY:
		fallthrough
	case CYCLEHOURLY:
		fallthrough
	case CYCLEDAILY:
		fallthrough
	case CYCLEWEEKLY:
		timerange = accrueDur
	case CYCLEMONTHLY:
		timerange = d1.AddDate(0, 1, 0).Sub(d1)
	case CYCLEQUARTERLY:
		timerange = d1.AddDate(0, 3, 0).Sub(d1)
	case CYCLEYEARLY:
		timerange = d1.AddDate(1, 0, 0).Sub(d1)
	}

	return timerange
}

// IsManageToBudget returns true if the supplied assessment type is managed to budget.
// Otherwise, it returns false.
func IsManageToBudget(xbiz *XBusiness, a *Assessment) bool {
	return RRdb.BizTypes[xbiz.P.BID].GLAccounts[a.ATypeLID].ManageToBudget == 1
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

// GetLIDFromGLAccountName returns the LID based on the supplied GLAccount name. It returns
// 0 if no account matched the supplied name
func GetLIDFromGLAccountName(bid int64, s string) int64 {
	for k, v := range RRdb.BizTypes[bid].GLAccounts {
		if v.Name == s {
			return k
		}
	}
	return int64(0)
}

// GetGLAccountChildAccts returns an array of LIDs whose parent are the suppliedbased on the supplied GLAccount name. If
// there are no child accounts, the list will be empty
func GetGLAccountChildAccts(bid, lid int64) []int64 {
	var m []int64
	for _, v := range RRdb.BizTypes[bid].GLAccounts {
		if v.PLID == lid {
			m = append(m, v.LID)
		}
	}
	return m
}

// GetRAAccountActivity returns the summed Amount balance for activity
// in GLAccount lid associated with RentalAgreement raid
func GetRAAccountActivity(bid, lid, raid int64, d1, d2 *time.Time) (float64, error) {
	var bal = float64(0)
	m, err := GetLedgerEntriesForRAID(d1, d2, raid, lid)
	if err != nil {
		return bal, err
	}
	for i := 0; i < len(m); i++ {
		bal += m[i].Amount
	}
	return bal, err
}

// GetRentableAccountActivity returns the summed Amount balance for activity
// in GLAccount lid associated with Rentable rid
func GetRentableAccountActivity(bid, lid, rid int64, d1, d2 *time.Time) (float64, error) {
	var bal = float64(0)
	m, err := GetLedgerEntriesForRentable(d1, d2, rid, lid)
	if err != nil {
		return bal, err
	}
	for i := 0; i < len(m); i++ {
		bal += m[i].Amount
	}
	return bal, err
}

// GetRAAccountBalance returns the balance of the account with LID lid on date dt. If raid is 0 then all
// transactions are considered. Otherwise, only transactions involving this RAID are considered.
func GetRAAccountBalance(bid, lid, raid int64, dt *time.Time) float64 {
	// fmt.Printf("GetRAAccountBalance: bid = %d, lid = %d, dt = %s ", bid, lid, dt.Format(RRDATEFMT4))
	bal := float64(0)
	//--------------------------------------------------------------------------------
	// First, check and see if this is a Parent to any other GLAccounts. If so, then
	// compute their totals
	//--------------------------------------------------------------------------------
	m := GetGLAccountChildAccts(bid, lid)
	for i := 0; i < len(m); i++ {
		bal += GetRAAccountBalance(bid, m[i], raid, dt)
		// fmt.Printf("L%08d child %d = L%08d  ==> bal = %d\n", lid, i, m[i], bal)
	}

	//--------------------------------------------------------------------------------
	// Compute the total for this account
	//--------------------------------------------------------------------------------
	lm := GetRALedgerMarkerOnOrBefore(bid, lid, raid, dt) // find nearest ledgermarker, use it as a basis
	if lm.LMID > 0 {
		bal += lm.Balance // we initialize the balance to this amount
		// fmt.Printf("LedgerMarker( bid=%d, lid=%d, raid=%d ) --> LM%08d,  dt = %10s, balance = %8.2f ==>  bal = %8.2f\n", bid, lid, raid, lm.LMID, lm.Dt.Format(RRDATEFMT4), lm.Balance, bal)
	}

	// Get the sum of the activeity between requested date and LedgerMarker
	activity, _ := GetRAAccountActivity(bid, lid, raid, &lm.Dt, dt)
	bal += activity

	// fmt.Printf("====>  balance = %.2f\n", bal)
	return bal
}

// GetRentableAccountBalance returns the balance of the account with LID lid on date dt. If rid is 0 then all
// transactions are considered. Otherwise, only transactions involving this RID are considered.
func GetRentableAccountBalance(bid, lid, rid int64, dt *time.Time) float64 {
	bal := float64(0)
	m := GetGLAccountChildAccts(bid, lid) // if parent acct, get info to compute aggregate balance
	for i := 0; i < len(m); i++ {
		bal += GetRentableAccountBalance(bid, m[i], rid, dt) // recurse
	}
	// Compute the total for this account. Start by getting any initial balance
	lm := GetRentableLedgerMarkerOnOrBefore(bid, lid, rid, dt) // find nearest ledgermarker, use it as a basis
	if lm.LMID > 0 {
		bal += lm.Balance // we initialize the balance to this amount
	}
	// Get the sum of the activeity between requested date and LedgerMarker
	activity, _ := GetRentableAccountActivity(bid, lid, rid, &lm.Dt, dt)
	bal += activity
	return bal
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
	if rrt.RentCycle > CYCLENORECUR { // if there's an override for RentCycle...
		rc = rrt.RentCycle // ...set it
	} else {
		rc = xbiz.RT[rtid].RentCycle
	}
	if rrt.ProrationCycle > CYCLENORECUR { // if there's an override for Propration...
		pro = rrt.ProrationCycle // ...set it
	} else {
		pro = xbiz.RT[rtid].Proration
	}

	// we need to load the RentableType for RentCycle or Proration or both...
	return rc, pro, rtid, err
}

// Prorate computes basic info to perform rent proration:
// examples:
//   DTSTART      DTSTOP       D1           D2         CYCLE  PRORATION   ASMTDUR  RENTDUR    PF      ANALYZE: START - STOP
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

	if CYCLENORECUR == prorateMethod {
		pf = 1.0
	} else {
		pf = float64(rentDur) / float64(asmtDur)
	}

	return asmtDur, rentDur, pf
}
