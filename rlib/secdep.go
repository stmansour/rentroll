package rlib

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// SecDepBalCacheCtx is the context used for the RAR Balance cache.
var SecDepBalCacheCtx = SimpleCacheCtx{
	Expiry: time.Duration(time.Minute * 5),
}
var secdepcache = map[string]*BalanceCacheEntry{} // initialize an empty cache

// SecDepBalCacheController is a go routine that will control access to
// balcache when multiple routines are trying to write to it.
//-----------------------------------------------------------------------------
func SecDepBalCacheController() {
	SecDepBalCacheCtx.SemAck = make(chan int)
	SecDepBalCacheCtx.Sem = make(chan int)
	for {
		select {
		case <-SecDepBalCacheCtx.Sem:
			SecDepBalCacheCtx.SemAck <- 1 // Let the caller know they have it
			<-SecDepBalCacheCtx.SemAck    // wait until caller is finished
		}
	}
}

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
		t := time.Now().Add(SecDepBalCacheCtx.Expiry) // it's life gets extended
		SecDepBalCacheCtx.Sem <- 1                    // request write access
		<-SecDepBalCacheCtx.SemAck                    // pause until we get access
		secdepcache[k].expire = &t                    // <<<<<<<<<<<<<<<  do the cache update
		SecDepBalCacheCtx.SemAck <- 1                 // tell the controller we're done
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
	t := time.Now().Add(SecDepBalCacheCtx.Expiry) // it gets this much time
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
	SecDepBalCacheCtx.Sem <- 1    // request write access
	<-SecDepBalCacheCtx.SemAck    // pause until we get access
	secdepcache[k] = &b           // <<<<<<<<<<<<<<<    do the cache update
	SecDepBalCacheCtx.SemAck <- 1 // tell the controller we're done
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
	for k, v := range secdepcache {
		if v == nil {
			continue
		}
		if force || now.After(*v.expire) {
			SecDepBalCacheCtx.Sem <- 1    // request write access
			<-SecDepBalCacheCtx.SemAck    // pause until we get access
			secdepcache[k] = nil          // <<<<<<<<<<<<<<<  do the cache update
			SecDepBalCacheCtx.SemAck <- 1 // tell the controller we're done
		}
		if nil != secdepcache[k] {
			i++
		}
	}
	// Console("SecDepBALCACHE After cleaning: %d\n", i)
}

// PrintSecDepBalCache is a debug routine to show the contents of the cache
//-----------------------------------------------------------------------------
func PrintSecDepBalCache() {
	i := 0
	for _, v := range secdepcache {
		if v != nil {
			i++
			// Console("bid: %d, rid: %d, raid: %d,   expire: %v\n", v.bid, v.rid, v.raid, v.expire)
		}
	}
	// Console("secdepcache size = %d  at  %v\n", i, time.Now())
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
func GetSecDepBalance(ctx context.Context, bid, raid, rid int64, d1, d2 *time.Time) (float64, error) {
	//-------------------------------
	// first, check the cache...
	//-------------------------------
	b := getSecDepCachedBalanceEntry(bid, rid, raid, d1, d2)
	if b != nil {
		return b.begin, nil
	}

	amt := float64(0)
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
	sa := []string{}
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
	if err != nil {
		return amt, err
	}
	storeSecDepBalanceInfoToCache(bid, rid, raid, d1, d2, amt /* this is the val we'll use */, amt /*placeholder*/) // cache this value, maybe we'll hit it again
	return amt, nil

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
