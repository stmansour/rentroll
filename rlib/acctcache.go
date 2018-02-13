package rlib

import (
	"fmt"
	"time"
)

// GLAcctSliceCacheEntry is the data type for balance cache entries.
type GLAcctSliceCacheEntry struct {
	bid      int64
	accttype string
	m        []int64
	expire   *time.Time
}

// GLAcctCacheCtx is the context used for the RAR Balance cache.
var GLAcctCacheCtx = SimpleCacheCtx{
	Expiry: time.Duration(time.Minute * 5),
}
var glacctcache = map[string]*GLAcctSliceCacheEntry{} // initialize an empty cache

// GLAcctCacheController is a go routine that will control access to
// glacctcache when multiple routines are trying to write to it.
//-----------------------------------------------------------------------------
func GLAcctCacheController() {
	GLAcctCacheCtx.SemAck = make(chan int)
	GLAcctCacheCtx.Sem = make(chan int)
	for {
		select {
		case <-GLAcctCacheCtx.Sem:
			GLAcctCacheCtx.SemAck <- 1 // Let the caller know they have it
			<-GLAcctCacheCtx.SemAck    // wait until caller is finished
		}
	}
}

// getGLAcctCacheKey returns the string used as a key in the map for
// the supplied input variables.
//
// INPUTS
//  bid  - biz id
//  accttype - what type of account
//
// RETURNS
//  a key string
//-----------------------------------------------------------------------------
func getGLAcctCacheKey(bid int64, accttype string) string {
	return fmt.Sprintf("%d %s", bid, accttype)
}

// getCachedGLAcctSliceEntry retrieves the balance from the cache if
// it exists. If it is found, it's expire time is extended.
//
// INPUTS
//  bid  - biz id
//  accttype - what type of account
//
// RETURNS
//  pointer to the GLAcctSliceCacheEntry if it exists otherwise it
//  returns nil.
//-----------------------------------------------------------------------------
func getCachedGLAcctSliceEntry(bid int64, accttype string) *GLAcctSliceCacheEntry {
	k := getGLAcctCacheKey(bid, accttype)
	// Console("getCachedGLAcctSliceEntry: key = %s\n", k)
	b, ok := glacctcache[k]
	if ok {
		if b == nil {
			return nil
		}
		t := time.Now().Add(GLAcctCacheCtx.Expiry) // it's life gets extended
		GLAcctCacheCtx.Sem <- 1                    // request write access
		<-GLAcctCacheCtx.SemAck                    // pause until we get access
		glacctcache[k].expire = &t                 // <<<<<<<<<<<<<<<  do the cache update
		GLAcctCacheCtx.SemAck <- 1                 // tell the controller we're done

		// Console("Expire updated: %s,  key = %s\n", t.Format(RRDATETIMEINPFMT), k)
		return b
	}
	return nil
}

// storeGLAcctSliceToCache stores the supplied cache entry. Since this
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
func storeGLAcctSliceToCache(bid int64, accttype string, m []int64) {
	t := time.Now().Add(GLAcctCacheCtx.Expiry) // it gets this much time
	b := GLAcctSliceCacheEntry{
		bid:      bid,
		accttype: accttype,
		m:        m,
		expire:   &t,
	}
	k := getGLAcctCacheKey(bid, accttype)
	// Console("Expire: %s,  key: %s\n", t.Format(RRDATETIMEINPFMT), k)

	GLAcctCacheCtx.Sem <- 1    // request write access
	<-GLAcctCacheCtx.SemAck    // pause until we get access
	glacctcache[k] = &b        // <<<<<<<<<<<<<<<    do the cache update
	GLAcctCacheCtx.SemAck <- 1 // tell the controller we're done
}

// CleanGLAcctSliceCache examines all the cache values and essentially
// removes the ones that have timed out.  If the force flag is true
// then all entries are removed from the cache
//
// INPUTS
//  force - a boolean where true means remove all entries from the cache
//
// RETURNS
//  nothing
//-----------------------------------------------------------------------------
func CleanGLAcctSliceCache(force bool) {

	now := time.Now()
	// Console("Entered Clean, force = %t,  now = %v\n", force, now)
	i := 0
	for k, v := range glacctcache {
		if v == nil {
			continue
		}
		if force || now.After(*v.expire) {
			GLAcctCacheCtx.Sem <- 1    // request write access
			<-GLAcctCacheCtx.SemAck    // pause until we get access
			glacctcache[k] = nil       // <<<<<<<<<<<<<<<  do the cache update
			GLAcctCacheCtx.SemAck <- 1 // tell the controller we're done
		}
		if nil != glacctcache[k] {
			i++
		}
	}
	// Console("glacctcache After cleaning: %d\n", i)
}

// GLAcctCacheSize is a debug routine to show the contents of the cache
//-----------------------------------------------------------------------------
func GLAcctCacheSize() int {
	i := 0
	for _, v := range glacctcache {
		if v != nil {
			i++
		}
	}
	return i
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

	b := getCachedGLAcctSliceEntry(bid, accttype)
	if b != nil {
		return b.m, nil
	}

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

	storeGLAcctSliceToCache(bid, accttype, m)

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
