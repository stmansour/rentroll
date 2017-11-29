package rlib

import (
	"fmt"
	"strings"
	"time"
)

// SecDepBalCacheExpiry is the time that entries in the balance cache live
var SecDepBalCacheExpiry = time.Duration(time.Minute * 1) // a cached value is good for 1 mins

var secdepcache = map[string]*BalanceCacheEntry{} // initialize an empty cache

// getSecDepCachedBalanceEntry retrieves the balance from the cache if
// it exists. If it is found, it's expire time is extended.
//
// INPUTS
//  bid  - biz id
//  rid  - retntable id
//  raid - rental agreement id
//  d1   - time range start
//  d2   - time range stop
//
// RETURNS
//  pointer to the BalanceCacheEntry if it exists otherwise it
//  returns nil.
//-----------------------------------------------------------------------------
func getSecDepCachedBalanceEntry(bid, rid, raid int64, d1, d2 *time.Time) *BalanceCacheEntry {
	k := getCacheKey(bid, rid, raid, d1, d2)
	// Console("getSecDepCachedBalanceEntry: key = %s\n", k)
	b, ok := secdepcache[k]
	if ok {
		if b == nil {
			return nil
		}
		t := time.Now().Add(SecDepBalCacheExpiry) // it's life gets extended
		secdepcache[k].expire = &t
		// Console("Expire updated: %s,  key = %s\n", t.Format(RRDATETIMEINPFMT), k)
		return b
	}
	return nil
}

// storeSecDepBalanceInfoToCache stores the supplied cache entry. Since this
// routine is private, it does not check the cache for an existing
// entry at this key. It assumes the caller understands how to use it.
//
// INPUTS
//  bid  - biz id
//  rid  - retntable id
//  raid - rental agreement id
//  d1   - time range start
//  d2   - time range stop
//
// RETURNS
//  nothing
//-----------------------------------------------------------------------------
func storeSecDepBalanceInfoToCache(bid, rid, raid int64, d1, d2 *time.Time, begin, end float64) {
	t := time.Now().Add(SecDepBalCacheExpiry) // it gets this much time
	b := BalanceCacheEntry{
		bid:    bid,
		rid:    rid,
		raid:   raid,
		d1:     d1,
		d2:     d2,
		begin:  begin,
		end:    end,
		expire: &t,
	}
	k := getCacheKey(bid, rid, raid, d1, d2)
	// Console("Expire: %s,  key: %s\n", t.Format(RRDATETIMEINPFMT), k)
	secdepcache[k] = &b
}

// CleanSecDepBalanceInfoCache examines all the cache values and essentially
// removes the ones that have timed out.  If the force flag is true
// then all entries are removed from the cache
//
// INPUTS
//  force - a boolean where true means remove all entries from the cache
//
// RETURNS
//  nothing
//-----------------------------------------------------------------------------
func CleanSecDepBalanceInfoCache(force bool) {

	now := time.Now()
	// Console("Entered Clean, force = %t,  now = %v\n", force, now)
	i := 0
	for k, v := range balcache {
		if v == nil {
			continue
		}
		if force || now.After(*v.expire) {
			balcache[k] = nil
		}
		if nil != balcache[k] {
			i++
		}
	}
	// Console("BALCACHE After cleaning: %d\n", i)
}

// AcctSlice returns a slice of GLAccounts that match the supplied accttype
//
// PARAMS
//      bid - which business
// accttype - account type of interest
//
// RETURNS
// []int64 - slice of LIDs that are of type Security Deposit Liability
//   error - any error encountered
//-----------------------------------------------------------------------------
func AcctSlice(bid int64, accttype string) ([]int64, error) {
	m := []int64{}
	q := fmt.Sprintf("SELECT LID FROM GLAccount WHERE BID=%d AND AcctType=%q", bid, accttype)
	rows, err := RRdb.Dbrr.Query(q)
	if err != nil {
		return m, err
	}
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return m, err
		}
		m = append(m, id)
	}
	err = rows.Err()
	return m, err
}

// AcctRulesSlice returns a slice of ARIDs for the Rules that work with Security
// Deposit accounts.
//
// PARAMS
//     bid - which business
//
// RETURNS
// []int64 - slice of ARIDs that affect Security Deposit Liability accounts
//   error - any error encountered
//-----------------------------------------------------------------------------
func AcctRulesSlice(sda []int64) ([]int64, error) {
	var m []int64
	//-----------------------------------------------------------
	// What Account Rules involve the Security Deposit accounts
	//-----------------------------------------------------------
	sa := []string{}
	for i := 0; i < len(sda); i++ {
		sa = append(sa, fmt.Sprintf("CreditLID=%d OR DebitLID=%d", sda[i], sda[i]))
	}
	q := "SELECT ARID FROM AR WHERE (" + strings.Join(sa, " OR ") + ")"
	// Console("GetSecDepBalance: q = %s\n", q)

	rows, err := RRdb.Dbrr.Query(q)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return m, err
		}
		m = append(m, id)
	}
	err = rows.Err()
	return m, err
}

// SecDepAccts returns a slice of LiabilitySecurityDeposit accounts
//
// PARAMS
//     bid - which business
//
// RETURNS
// []int64 - slice of LIDs that are of type Security Deposit Liability
//   error - any error encountered
//-----------------------------------------------------------------------------
func SecDepAccts(bid int64) ([]int64, error) {
	return AcctSlice(bid, LiabilitySecDep)
}

// SecDepRules returns a slice of ARIDs for the Rules that work with Security
// Deposit accounts.
//
// PARAMS
//     bid - which business
//
// RETURNS
// []int64 - slice of ARIDs that affect Security Deposit Liability accounts
//   error - any error encountered
//-----------------------------------------------------------------------------
func SecDepRules(bid int64) ([]int64, error) {
	var m []int64
	sda, err := SecDepAccts(bid)
	if err != nil {
		return m, err
	}
	if len(sda) == 0 {
		return m, fmt.Errorf("There are no accounts of type %s where BID = %d", LiabilitySecDep, bid)
	}

	return AcctRulesSlice(sda)
}

// GetSecDepBalance returns the amount of security deposit charge and the
// amount that was assessed for the supplied Rental Agreement and RID
//
// PARAMS
//	bid  - business id
//  rid  - the rentable for which the deposit was assessed
//  raid - the Rental Agreement associated with the assessment
//  d1   - start time; do not consider assessments prior to this date
//  d2   - stop time; do not considder assessments on or after this date
//
// RETURNS
// float64 - Amount of change in Security Deposit Balance between d1 and d2
//   error - any error encountered
//-----------------------------------------------------------------------------
func GetSecDepBalance(bid, raid, rid int64, d1, d2 *time.Time) (float64, error) {
	amt := float64(0)
	sa := []string{}
	m, err := SecDepRules(bid)
	if err != nil {
		return amt, fmt.Errorf("Error in SecDepRules: %s", err.Error())
	}
	if len(m) == 0 {
		return amt, fmt.Errorf("There are no account rules that credit a %s account", LiabilitySecDep)
	}
	//-----------------------------------------------------------
	// What Assessments use the account rules found above?
	//-----------------------------------------------------------
	sa = []string{}
	for i := 0; i < len(m); i++ {
		sa = append(sa, fmt.Sprintf("ARID=%d", m[i]))
	}
	q := fmt.Sprintf("SELECT SUM(Amount) AS Amt FROM Assessments WHERE BID=%d AND RID=%d and RAID=%d AND %q<=Start AND Stop<%q AND (%s)",
		bid, rid, raid, d1.Format(RRDATEFMTSQL), d2.Format(RRDATEFMTSQL), strings.Join(sa, " OR "))
	// Console("=======>>>>>>  q:  %s\n", q)
	rows, err := RRdb.Dbrr.Query(q)
	for rows.Next() {
		var x NullFloat64
		err := rows.Scan(&x)
		if err != nil {
			return amt, err
		}
		amt += x.Float64
	}
	err = rows.Err()
	return amt, err
}

// GetSecDepBalanceOnDate
// func GetSecDepBalanceOnDate(bid, raid, rid int64, d1, d2 *time.Time) (float64, error) {
// 	amt := float64(0)
// 	sa := []string{}
// 	m, err := SecDepRules(bid)
// 	if err != nil {
// 		return amt, fmt.Errorf("Error in SecDepRules: %s", err.Error())
// 	}
// 	if len(m) == 0 {
// 		return amt, fmt.Errorf("There are no account rules that credit a %s account", LiabilitySecDep)
// 	}
// 	//-----------------------------------------------------------
// 	// What Assessments use the account rules found above?
// 	//-----------------------------------------------------------
// 	sa = []string{}
// 	for i := 0; i < len(m); i++ {
// 		sa = append(sa, fmt.Sprintf("ARID=%d", m[i]))
// 	}
// 	q := fmt.Sprintf("SELECT SUM(Amount) AS Amt FROM Assessments WHERE BID=%d AND RID=%d and RAID=%d AND %q<=Start AND Stop<%q AND (%s) GROUP BY RID",
// 		bid, rid, raid, d1.Format(RRDATEFMTSQL), d2.Format(RRDATEFMTSQL), strings.Join(sa, " OR "))
// 	// Console("=======>>>>>>  q:  %s\n", q)
// 	rows, err := RRdb.Dbrr.Query(q)
// 	for rows.Next() {
// 		var x float64
// 		err := rows.Scan(&x)
// 		if err != nil {
// 			return amt, err
// 		}
// 		amt += x
// 	}
// 	err = rows.Err()
// 	return amt, err
// }
