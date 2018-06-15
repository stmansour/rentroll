package rlib

import (
	"context"
	"fmt"
	"math"
	"time"
)

// LiabilitySecDep is the string used to identify a security deposit
// liability account. It is used in queries.
const (
	LiabilitySecDep    = "Liability Security Deposit"
	AccountsReceivable = "Accounts Receivable"
)

// RDateFmt is an array of date / time formats that RentRoll accepts for datetime input
var RDateFmt = []string{
	RRDATETIMEINPFMT,
	RRDATEFMT,
	RRDATEFMT2,
	RRDATEFMT3,
	RRDATEINPFMT,
}

// QBAcctType - this was previously a text field. But because we dropped the
// notion of "default accounts", it was important to formalize the list of account
// types.
var QBAcctType []string

// Assets        = Liabilities + Owner Equity
// ∂ OwnerEquity = (t1)Income(t1) - Expenses(t1)
// Contract Rent = Gross Scheduled Rent - Income Offsets
// Receipt       = Contract Rent - Expenses

// QBAcctInfo indicates how numbers should be processed in the account Rules
var QBAcctInfo = []struct {
	Name   string
	Negate bool // Indicates whether or not negate an assessment amount if showing it as a debit
}{
	{"Asset", false},            // Asset         D +   C -
	{AccountsReceivable, false}, // Asset         D +   C -
	{"Cash", false},             // Asset         D +   C -
	{"Expense", false},          // Expense Acct  D +   C -
	{"Liabilities", true},       // Liabilities   D -   C +
	{LiabilitySecDep, true},     // Liabilities   D -   C +
	{"Income", true},            // Income Acct   D -   C +
	{"Income Offset", true},     // Income Acct   D -   C +
	{"Other Income", true},      // Income Acct   D -   C +
}

// AccountTypeNegateFlag returns the Negate flag associated with the supplied account type.
// Basically, if you're DEBITing the account the flag tells whether or not you should negate it.
func AccountTypeNegateFlag(s string) bool {
	for i := 0; i < len(QBAcctInfo); i++ {
		if s == QBAcctInfo[i].Name {
			return QBAcctInfo[i].Negate
		}
	}
	return false
}

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

// Round floating point numbers to the specified number of decimal places.
// Without this kind of routine, adding large list of numbers in something
// like the Rentroll view's Grand Total can result in being off by $0.01.
// Adding the Round function to the results of the rounding calculation fixes
// this kind of problem.
//
// INPUTS
//  val     - value to be rounded
//  roundOn - the pivot point for rounding
//  placess - number of decimal places to round to
//
// OUTPUT
//  newVal  - the rounded number
//-----------------------------------------------------------------------------
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// SimpleProrateAmount prorates the supplied amount for the supplied dates.
//
// INPUTS:
//  amt          - unprorated amount for the RentCycle
//  RentCycle    - 0 = norecur, 1 = secondly, ... 7 = Yearly
//  ProrateCycle - cycle over which RentCycle is divided to prorate usage
//  d1           - start time
//  d2           - stop time
//  epoch        - used to determine the start of the rent period
//
// RETURNS:
//  the prorated amount
//  the number of periods during this cycle
//  the total number of possible periods this cycle
//---------------------------------------------------------------------------
func SimpleProrateAmount(amt float64, RentCycle, Prorate int64, d1, d2, epoch *time.Time) (float64, int64, int64) {
	// Console("Entered SimpleProrateAmount: amt = %.2f, RentCycle = %d, Prorate = %d, d1 = %s, d2 = %s\n", amt, RentCycle, Prorate, d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4))
	var thisepoch time.Time
	if RECURNONE == Prorate || RECURNONE == RentCycle {
		return amt, int64(1), int64(1)
	}
	switch RentCycle {
	case RECURSECONDLY:
		fallthrough
	case RECURMINUTELY:
		fallthrough
	case RECURHOURLY:
		fallthrough
	case RECURDAILY:
		thisepoch = *epoch
	case RECURWEEKLY:
		thisepoch = *epoch
	case RECURMONTHLY:
		thisepoch = time.Date(d1.Year(), d1.Month(), epoch.Day(), epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
	case RECURQUARTERLY:
		thisepoch = time.Date(d1.Year(), d1.Month(), epoch.Day(), epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
	case RECURYEARLY:
		thisepoch = time.Date(d1.Year(), d1.Month(), epoch.Day(), epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
	}

	cycdur := CycleDuration(RentCycle, thisepoch)
	proratedur := CycleDuration(Prorate, thisepoch)

	dur := d2.Sub(*d1)
	numPeriods := int64(dur) / int64(proratedur)
	totalPeriods := int64(cycdur) / int64(proratedur)
	// Console("numPeriods = %d, totalPeriods = %d\n", numPeriods, totalPeriods)
	rounded := Round(amt*float64(numPeriods)/float64(totalPeriods), .5, 2)
	return rounded, numPeriods, totalPeriods
}

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

// SelectRentableStatusForPeriod returns a subset of Rentable states that
// overlap the supplied range.
//=============================================================================
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

// GetRentableStateForDate returns the status of the Rentable on the supplied
// date
//=============================================================================
func GetRentableStateForDate(ctx context.Context, rid int64, dt *time.Time) (int64, error) {
	status := int64(RENTABLESTATUSUNKNOWN)
	d2 := dt.Add(24 * time.Hour)

	m, err := GetRentableStatusByRange(ctx, rid, dt, &d2)
	if err != nil {
		return status, err
	}

	if len(m) > 0 {
		status = m[0].UseStatus
	}
	return status, err
}

// GetLIDFromGLAccountName returns the LID based on the supplied GLAccount
// name. It returns
// 0 if no account matched the supplied name
//=============================================================================
func GetLIDFromGLAccountName(bid int64, s string) int64 {
	for k, v := range RRdb.BizTypes[bid].GLAccounts {
		if v.Name == s {
			return k
		}
	}
	return int64(0)
}

// GetGLAccountChildAccts returns an array of LIDs whose parent are the
// suppliedbased on the supplied GLAccount name. If there are no child
// accounts, the list will be empty
//=============================================================================
func GetGLAccountChildAccts(ctx context.Context, bid, lid int64) ([]int64, error) {
	var m []int64
	for _, v := range RRdb.BizTypes[bid].GLAccounts {
		if v.PLID == lid {
			m = append(m, v.LID)
		}
	}

	// TODO(): returning things from the memory cache,
	// need authorization of user on top of this call/here.
	return m, nil
}

// GetAccountActivity returns the summed Amount balance for activity
// in GLAccount lid associated with RentalAgreement raid
//=============================================================================
func GetAccountActivity(ctx context.Context, bid, lid int64, d1, d2 *time.Time) (float64, error) {
	var bal = float64(0)
	m, err := GetLedgerEntriesInRange(ctx, d1, d2, bid, lid)
	if err != nil {
		return bal, err
	}
	for i := 0; i < len(m); i++ {
		bal += m[i].Amount
	}
	return bal, err
}

// GetRAAccountActivity returns the summed Amount balance for activity
// in GLAccount lid associated with RentalAgreement raid
//=============================================================================
func GetRAAccountActivity(ctx context.Context, bid, lid, raid int64, d1, d2 *time.Time) (float64, error) {
	var bal = float64(0)
	m, err := GetLedgerEntriesForRAID(ctx, d1, d2, raid, lid)
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
//=============================================================================
func GetRentableAccountActivity(ctx context.Context, bid, lid, rid int64, d1, d2 *time.Time) (float64, error) {
	var bal = float64(0)
	m, err := GetLedgerEntriesForRentable(ctx, d1, d2, rid, lid)
	if err != nil {
		return bal, err
	}
	for i := 0; i < len(m); i++ {
		bal += m[i].Amount
	}
	return bal, err
}

// GetAccountTypeBalance totals the leaf node accounts for the supplied
// account type
//   a = AccountType for which to retrieve balance
// bid = which business
//  dt = balance on this date
//
// RETURNS:
//   float64 balance
//   error or nil
//=============================================================================
func GetAccountTypeBalance(ctx context.Context, a string, bid int64, dt *time.Time) (float64, error) {
	bal := float64(0)
	found := false
	for i := 0; i < len(QBAcctType); i++ { // make sure we have a valid
		found := QBAcctType[i] == a
		if found {
			break
		}
	}
	if !found {
		return bal, fmt.Errorf("Account Type %s is unknown", a)
	}
	_, ok := RRdb.BizTypes[bid]
	if !ok {
		return bal, fmt.Errorf("No business found for BID = %d", bid)
	}

	rows, err := RRdb.Prepstmt.GetLedgerList.Query(bid)
	if err != nil {
		return bal, err
	}
	defer rows.Close()

	for rows.Next() {
		var r GLAccount
		ReadGLAccounts(rows, &r)
		if r.AcctType == a && r.AllowPost {
			b, err := GetAccountBalance(ctx, bid, r.LID, dt)
			if err != nil {
				return bal, err
			}
			bal += b
		}
	}

	return bal, rows.Err()
}

// GetRAAccountBalance returns the balance of the account with LID lid on date
// dt. If raid is 0 then all transactions are considered. Otherwise, only
// transactions involving this RAID are considered.
//=============================================================================
func GetRAAccountBalance(ctx context.Context, bid, lid, raid int64, dt *time.Time) (float64, error) {
	// fmt.Printf("GetRAAccountBalance: bid = %d, lid = %d, raid = %d, dt = %s ", bid, lid, raid, dt.Format(RRDATEFMT4))
	bal := float64(0)
	//--------------------------------------------------------------------------------
	// First, check and see if this is a Parent to any other GLAccounts. If so, then
	// compute their totals
	//--------------------------------------------------------------------------------
	m, err := GetGLAccountChildAccts(ctx, bid, lid)
	if err != nil {
		return bal, err
	}

	for i := 0; i < len(m); i++ {

		b, err := GetRAAccountBalance(ctx, bid, m[i], raid, dt)
		if err != nil {
			return bal, err
		}
		bal += b
		// fmt.Printf("L%08d child %d = L%08d  ==> bal = %8.2f\n", lid, i, m[i], bal)
	}

	//--------------------------------------------------------------------------------
	// Compute the total for this account. If this ledger does not allow posts, don't
	// consider its Balance.
	//--------------------------------------------------------------------------------
	// TODO(Steve): if no ledger marker found then should we raise the error?
	lm, err := GetRALedgerMarkerOnOrBeforeDeprecated(ctx, bid, lid, raid, dt) // find nearest ledgermarker, use it as a basis
	if err != nil {
		return bal, err
	}

	// fmt.Printf("GetRALedgerMarkerOnOrBeforeDeprecated(bid,lid,raid,dt) = lm.LMID = %d, lm.Dt = %s\n", lm.LMID, lm.Dt.Format(RRDATEFMT4))
	if lm.LMID > 0 && RRdb.BizTypes[bid].GLAccounts[lid].AllowPost {
		bal += lm.Balance // we initialize the balance to this amount
		// fmt.Printf("LedgerMarker( bid=%d, lid=%d, raid=%d ) --> LM%08d,  dt = %10s, lm.Balance = %8.2f ==>  bal = %8.2f\n", bid, lid, raid, lm.LMID, lm.Dt.Format(RRDATEFMT4), lm.Balance, bal)
	}

	// Get the sum of the activity between requested date and LedgerMarker
	var activity float64

	// TODO(Steve): should we really ignore errors from here?
	if raid != 0 {
		activity, _ = GetRAAccountActivity(ctx, bid, lid, raid, &lm.Dt, dt)
		// fmt.Printf("GetRAAccountActivity(bid, lid, raid, &lm.Dt, dt) = %8.2f\n", activity)
	} else {
		activity, _ = GetAccountActivity(ctx, bid, lid, &lm.Dt, dt)
		// fmt.Printf("GetAccountActivity(bid, lid, &lm.Dt, dt) = %8.2f\n", activity)
	}
	bal += activity

	// fmt.Printf("====>  balance = %.2f\n", bal)
	return bal, nil
}

// GetAccountBalance returns the balance of the account with LID lid on date dt.
// It's just a wrapper around GetRAAccountBalance with raid set to 0.  This returns
// the account balance we're after, but with a more obvious function name to call.
//=============================================================================
func GetAccountBalance(ctx context.Context, bid, lid int64, dt *time.Time) (float64, error) {
	return GetRAAccountBalance(ctx, bid, lid, 0, dt)
}

// GetRentableAccountBalance returns the balance of the account with LID lid
// on date dt. If rid is 0 then all transactions are considered. Otherwise,
// only transactions involving this RID are considered.
//=============================================================================
func GetRentableAccountBalance(ctx context.Context, bid, lid, rid int64, dt *time.Time) (float64, error) {
	// fmt.Printf("GetRAAccountBalance: bid = %d, lid = %d, rid = %d, dt = %s\n", bid, lid, rid, dt.Format(RRDATEFMT4))
	bal := float64(0)
	m, err := GetGLAccountChildAccts(ctx, bid, lid) // if parent acct, get info to compute aggregate balance
	if err != nil {
		return bal, err
	}
	for i := 0; i < len(m); i++ {
		b, err := GetRentableAccountBalance(ctx, bid, m[i], rid, dt) // recurse
		if err != nil {
			return bal, err
		}
		bal += b
		// fmt.Printf("L%08d child %d = L%08d  ==> bal = %8.2f\n", lid, i, m[i], bal)
	}
	// Compute the total for this account. Start by getting any initial balance
	lm, err := GetRentableLedgerMarkerOnOrBefore(ctx, bid, lid, rid, dt) // find nearest ledgermarker, use it as a basis
	if err != nil {
		return bal, err
	}
	// fmt.Printf("GetRentableLedgerMarkerOnOrBefore( bid, lid, rid, dt = %s) --> LM%08d, \n", dt.Format(RRDATEFMT4), lm.LMID)
	if lm.LMID > 0 {
		bal += lm.Balance // we initialize the balance to this amount
		// fmt.Printf("LedgerMarkerOnOrBefore( bid=%d, lid=%d, rid=%d,  dt = %10s ) --> LM%08d, lm.Balance = %8.2f ==>  bal = %8.2f\n", bid, lid, rid, dt.Format(RRDATEFMT4), lm.LMID, lm.Balance, bal)
	}
	// Get the sum of the activity between requested date and LedgerMarker
	var activity float64

	// TODO(Steve): should really ignore the errors from here?
	if rid != 0 {
		activity, _ = GetRentableAccountActivity(ctx, bid, lid, rid, &lm.Dt, dt)
		// fmt.Printf("GetRentableAccountActivity(bid=%d, lid=%d, rid=%d, &lm.Dt = %s, dt = %s) = %8.2f\n", bid, lid, rid, lm.Dt.Format(RRDATEFMT4), dt.Format(RRDATEFMT4), activity)
	} else {
		activity, _ = GetAccountActivity(ctx, bid, lid, &lm.Dt, dt)
		// fmt.Printf("GetAccountActivity(bid=%d, lid=%d, &lm.Dt = %s, dt = %s) = %8.2f\n", bid, lid, lm.Dt.Format(RRDATEFMT4), dt.Format(RRDATEFMT4), activity)
	}

	bal += activity
	// fmt.Printf("====>  balance = %.2f\n", bal)
	return bal, err
}

// GetRentCycleAndProration returns the RentCycle (and Proration) to use for
// the supplied rentable and date. If the override RentCycle is set for this
// time period, it is returned. Otherwise, the RentCycle for this Rentable's
// RentableType is returned
// Returns:
//		RentCycle
//		Proration
//		rtid for the supplied date
//		error
//=============================================================================
func GetRentCycleAndProration(ctx context.Context, r *Rentable, dt *time.Time, xbiz *XBusiness) (int64, int64, int64, error) {
	var err error
	var rc, pro, rtid int64

	rrt, err := GetRentableTypeRefForDate(ctx, r.RID, dt)
	if err != nil {
		return rc, pro, rtid, err
	}

	if rrt.RID == 0 {
		return rc, pro, rtid, fmt.Errorf("No RentableTypeRef for %s", dt.Format(RRDATEINPFMT))
	}
	rtid, err = GetRTIDForDate(ctx, r.RID, dt)
	if err != nil {
		return rc, pro, rtid, err
	}
	if rrt.OverrideRentCycle > RECURNONE { // if there's an override for RentCycle...
		rc = rrt.OverrideRentCycle // ...set it
	} else {
		rc = xbiz.RT[rtid].RentCycle
	}
	if rrt.OverrideProrationCycle > RECURNONE { // if there's an override for Propration...
		pro = rrt.OverrideProrationCycle // ...set it
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
// func Prorate(RAStart, RAStop, asmtStart, asmtStop time.Time, accrual, prorateMethod int64) (int64, int64, float64) {
// 	var asmtDur int64
// 	var rentDur int64
// 	var pf float64

// 	prorateDur := CycleDuration(prorateMethod, asmtStart)
// 	//-------------------------------------------------------------------
// 	// Scope the Rental Agreement period down to this assessment period.
// 	// Overlap the Rental Agreement period (RAStart to RAStop) with the
// 	// assessment period (asmtStart - asmtStop)
// 	//-------------------------------------------------------------------
// 	start := asmtStart
// 	if RAStart.After(start) {
// 		start = RAStart
// 	}
// 	stop := RAStop.Add(prorateDur)
// 	if stop.After(asmtStop) {
// 		stop = asmtStop
// 	}

// 	// fmt.Printf("scoped period:  %s - %s\n", start.Format(RRDATEINPFMT), stop.Format(RRDATEINPFMT))
// 	asmtDur = int64(asmtStop.Sub(asmtStart) / prorateDur)
// 	rentDur = int64(stop.Sub(start) / prorateDur)

// 	// fmt.Printf("rentDur = %d %s\n", rentDur, units)
// 	// fmt.Printf("asmtDur = %d %s\n", asmtDur, units)

// 	if RECURNONE == prorateMethod {
// 		pf = 1.0
// 	} else {
// 		pf = float64(rentDur) / float64(asmtDur)
// 	}

// 	return asmtDur, rentDur, pf
// }
