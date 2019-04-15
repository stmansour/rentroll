package rlib

import (
	"fmt"
	"strings"
	"time"
)

// ARSliceCacheEntry is the data type for balance cache entries.
type ARSliceCacheEntry struct {
	bid      int64
	accttype string
	m        []int64
	expire   *time.Time
}

// ARCacheCtx is the context used for the RAR Balance cache.
var ARCacheCtx = SimpleCacheCtx{
	Expiry: time.Duration(time.Minute * 5),
}
var arcache = map[string]*ARSliceCacheEntry{} // initialize an empty cache

// ARCacheController is a go routine that will control access to
// arcache when multiple routines are trying to write to it.
//-----------------------------------------------------------------------------
func ARCacheController() {
	ARCacheCtx.SemAck = make(chan int)
	ARCacheCtx.Sem = make(chan int)
	for {
		select {
		case <-ARCacheCtx.Sem:
			ARCacheCtx.SemAck <- 1 // Let the caller know they have it
			<-ARCacheCtx.SemAck    // wait until caller is finished
		}
	}
}

// getARCacheKey returns the string used as a key in the map for
// the supplied input variables.
//
// INPUTS
//  bid  - biz id
//  accttype - what type of account
//
// RETURNS
//  a key string
//-----------------------------------------------------------------------------
func getARCacheKey(bid int64, accttype string) string {
	return fmt.Sprintf("%d %s", bid, accttype)
}

// getCachedARSliceEntry retrieves the balance from the cache if
// it exists. If it is found, it's expire time is extended.
//
// INPUTS
//  bid  - biz id
//  accttype - what type of account
//
// RETURNS
//  pointer to the ARSliceCacheEntry if it exists otherwise it
//  returns nil.
//-----------------------------------------------------------------------------
func getCachedARSliceEntry(bid int64, accttype string) *ARSliceCacheEntry {
	k := getARCacheKey(bid, accttype)
	// Console("getCachedARSliceEntry: key = %s\n", k)
	b, ok := arcache[k]
	if ok {
		if b == nil {
			return nil
		}
		t := time.Now().Add(ARCacheCtx.Expiry) // it's life gets extended
		ARCacheCtx.Sem <- 1                    // request write access
		<-ARCacheCtx.SemAck                    // pause until we get access
		arcache[k].expire = &t                 // <<<<<<<<<<<<<<<  do the cache update
		ARCacheCtx.SemAck <- 1                 // tell the controller we're done

		// Console("Expire updated: %s,  key = %s\n", t.Format(RRDATETIMEINPFMT), k)
		return b
	}
	return nil
}

// storeARSliceToCache stores the supplied cache entry. Since this
// routine is private, it does not check the cache for an existing
// entry at this key. It assumes the caller understands how to use it.
//
// INPUTS
//  bid  - biz id
//  accttype - what type of account
//
// RETURNS
//  nothing
//-----------------------------------------------------------------------------
func storeARSliceToCache(bid int64, accttype string, m []int64) {
	t := time.Now().Add(ARCacheCtx.Expiry) // it gets this much time
	b := ARSliceCacheEntry{
		bid:      bid,
		accttype: accttype,
		m:        m,
		expire:   &t,
	}
	k := getARCacheKey(bid, accttype)
	// Console("Expire: %s,  key: %s\n", t.Format(RRDATETIMEINPFMT), k)

	ARCacheCtx.Sem <- 1    // request write access
	<-ARCacheCtx.SemAck    // pause until we get access
	arcache[k] = &b        // <<<<<<<<<<<<<<<    do the cache update
	ARCacheCtx.SemAck <- 1 // tell the controller we're done
}

// CleanARSliceCache examines all the cache values and essentially
// removes the ones that have timed out.  If the force flag is true
// then all entries are removed from the cache
//
// INPUTS
//  force - a boolean where true means remove all entries from the cache
//
// RETURNS
//  nothing
//-----------------------------------------------------------------------------
func CleanARSliceCache(force bool) {
	now := time.Now()
	// Console("Entered Clean, force = %t,  now = %v\n", force, now)
	i := 0
	for k, v := range arcache {
		if v == nil {
			continue
		}
		if force || now.After(*v.expire) {
			ARCacheCtx.Sem <- 1    // request write access
			<-ARCacheCtx.SemAck    // pause until we get access
			arcache[k] = nil       // <<<<<<<<<<<<<<<  do the cache update
			ARCacheCtx.SemAck <- 1 // tell the controller we're done
		}
		if nil != arcache[k] {
			i++
		}
	}
	// Console("arcache After cleaning: %d\n", i)
}

// ARCacheSize is a debug routine to show the contents of the cache
//-----------------------------------------------------------------------------
func ARCacheSize() int {
	i := 0
	for _, v := range arcache {
		if v != nil {
			i++
		}
	}
	return i
}

// AcctRulesSlice returns a slice of ARIDs for the Rules that work with Security
// Deposit accounts.
//
// PARAMS
//      sda - slice of accounts of type acctype
// accttype - the gl account type
//      bid - which business
//
// RETURNS
// []int64 - slice of ARIDs that affect Security Deposit Liability accounts
//   error - any error encountered
//-----------------------------------------------------------------------------
func AcctRulesSlice(sda []int64, bid int64, accttype string) ([]int64, error) {
	b := getCachedARSliceEntry(bid, accttype)
	if b != nil {
		return b.m, nil
	}

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
	if err != nil {
		return m, err
	}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return m, err
		}
		m = append(m, id)
	}
	err = rows.Err()

	storeARSliceToCache(bid, accttype, m)
	return m, err
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

	return AcctRulesSlice(sda, bid, LiabilitySecDep)
}
