package rlib

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

// RARStmtEntry descriGetRARBalancebes an entry on a statement
type RARStmtEntry struct {
	T       int                // 1 = assessment, 2 = Receipt
	A       *Assessment        // for type==1, the pointer to the assessment
	R       *ReceiptAllocation // for type ==2, the pointer to the receipt
	Amt     float64            // amount of the receipt or assessment
	Reverse bool               // is this a reversal?
	Dt      time.Time          // date/time of this assessment or receipt
	TCID    int64              // IF THIS IS FOR A PAYOR STATEMENT, the TCID of the Payor, otherwise 0
}

// BalanceCacheEntry is the data type for balance cache entries.
type BalanceCacheEntry struct {
	bid    int64
	rid    int64
	raid   int64
	d1     *time.Time
	d2     *time.Time
	begin  float64
	end    float64
	expire *time.Time
}

// RARBalCacheCtx is the context used for the RAR Balance cache.
var RARBalCacheCtx = SimpleCacheCtx{
	Expiry: time.Duration(time.Minute * 5),
}
var balcache = map[string]*BalanceCacheEntry{} // initialize an empty cache

// RARBalCacheController is a go routine that will control access to
// balcache when multiple routines are trying to write to it.
//-----------------------------------------------------------------------------
func RARBalCacheController() {
	RARBalCacheCtx.SemAck = make(chan int)
	RARBalCacheCtx.Sem = make(chan int)
	for {
		select {
		case <-RARBalCacheCtx.Sem:
			RARBalCacheCtx.SemAck <- 1 // Let the caller know they have it
			<-RARBalCacheCtx.SemAck    // wait until caller is finished
		}
	}
}

// getCacheKey returns the string used as a key in the map for
// the supplied input variables.
//
// INPUTS
//  bid  - biz id
//  rid  - retntable id
//  raid - rental agreement id
//  d1   - time range start
//  d2   - time range stop
//
// RETURNS
//  a key string
//-----------------------------------------------------------------------------
func getCacheKey(bid, rid, raid int64, d1, d2 *time.Time) string {
	return fmt.Sprintf("%d %d %d %s %s", bid, rid, raid, d1.Format(RRDATEFMTSQL), d2.Format(RRDATEFMTSQL))
}

// getCachedRARBalanceEntry retrieves the balance from the cache if
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
func getCachedRARBalanceEntry(bid, rid, raid int64, d1, d2 *time.Time) *BalanceCacheEntry {
	k := getCacheKey(bid, rid, raid, d1, d2)
	// Console("getCachedRARBalanceEntry: key = %s\n", k)
	b, ok := balcache[k]
	if ok {
		if b == nil {
			return nil
		}
		t := time.Now().Add(RARBalCacheCtx.Expiry) // it's life gets extended
		RARBalCacheCtx.Sem <- 1                    // request write access
		<-RARBalCacheCtx.SemAck                    // pause until we get access
		balcache[k].expire = &t                    // <<<<<<<<<<<<<<<  do the cache update
		RARBalCacheCtx.SemAck <- 1                 // tell the controller we're done

		// Console("Expire updated: %s,  key = %s\n", t.Format(RRDATETIMEINPFMT), k)
		return b
	}
	return nil
}

// storeRARBalanceInfoToCache stores the supplied cache entry. Since this
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
func storeRARBalanceInfoToCache(bid, rid, raid int64, d1, d2 *time.Time, begin, end float64) {
	t := time.Now().Add(RARBalCacheCtx.Expiry) // it gets this much time
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

	RARBalCacheCtx.Sem <- 1    // request write access
	<-RARBalCacheCtx.SemAck    // pause until we get access
	balcache[k] = &b           // <<<<<<<<<<<<<<<    do the cache update
	RARBalCacheCtx.SemAck <- 1 // tell the controller we're done
}

// CleanRARBalanceInfoCache examines all the cache values and essentially
// removes the ones that have timed out.  If the force flag is true
// then all entries are removed from the cache
//
// INPUTS
//  force - a boolean where true means remove all entries from the cache
//
// RETURNS
//  nothing
//-----------------------------------------------------------------------------
func CleanRARBalanceInfoCache(force bool) {

	now := time.Now()
	// Console("Entered Clean, force = %t,  now = %v\n", force, now)
	i := 0
	for k, v := range balcache {
		if v == nil {
			continue
		}
		if force || now.After(*v.expire) {
			RARBalCacheCtx.Sem <- 1    // request write access
			<-RARBalCacheCtx.SemAck    // pause until we get access
			balcache[k] = nil          // <<<<<<<<<<<<<<<  do the cache update
			RARBalCacheCtx.SemAck <- 1 // tell the controller we're done
		}
		if nil != balcache[k] {
			i++
		}
	}
	// Console("BALCACHE After cleaning: %d\n", i)
}

// RARBalCacheSize is a debug routine to show the contents of the cache
//-----------------------------------------------------------------------------
func RARBalCacheSize() int {
	i := 0
	for _, v := range balcache {
		if v != nil {
			i++
		}
	}
	return i
}

// GetBeginEndRARBalance gets the balance associated with a Rentable and a
// Rental Agreement at a particular point in time.
//
// INPUTS
//   rid  - RID of Rentable
//   raid - RAID of Rental Agreement
//   d1   - time for which balance is requested
//   d2   - time for which balance is requested
//
// RETURNS
//   float64 - the balance for the Rentable rid in Rental Agreement raid at
//             time dt
//   error   - any error encountered
//-----------------------------------------------------------------------------
func GetBeginEndRARBalance(ctx context.Context, bid, rid, raid int64, d1, d2 *time.Time) (float64, float64, error) {
	//----------------------------------------
	// try to get it from the cache first...
	//----------------------------------------
	b := getCachedRARBalanceEntry(bid, rid, raid, d1, d2)
	if b != nil {
		return b.begin, b.end, nil
	}

	var err error
	var begin, end float64
	begin, err = GetRARBalance(ctx, bid, rid, raid, d1)
	if err != nil {
		return b.begin, b.end, nil
	}

	end, err = GetRARBalance(ctx, bid, rid, raid, d2)
	if err != nil {
		return b.begin, b.end, nil
	}

	storeRARBalanceInfoToCache(bid, rid, raid, d1, d2, begin, end) // cache this value, maybe we'll hit it again

	return begin, end, err
}

// GetRARBalance gets the balance associated with a Rentable and a
// Rental Agreement at a particular point in time.
//
// INPUTS
//   bid     - biz id
//   rid     - RID of Rentable
//   raid    - RAID of Rental Agreement
//   dt      - time for which balance is requested
//
// RETURNS
//   float64 - the balance for the Rentable rid in Rental Agreement raid at
//             time dt
//   error   - any error encountered
//-----------------------------------------------------------------------------
func GetRARBalance(ctx context.Context, bid, rid, raid int64, dt *time.Time) (float64, error) {
	const funcname = "GetRARBalance"

	var (
		bal float64
		err error
	)

	lm, err := GetRARentableLedgerMarkerOnOrBefore(ctx, raid, rid, dt)
	if lm.LMID == 0 {
		LogAndPrint("%s: Could not find LedgerMarker for RAID=%d, RID=%d, on or before %s\n",
			funcname, raid, rid, dt.Format(RRDATEFMT3))
		return bal, err
	}
	if err != nil {
		LogAndPrint("%s: Could not find LedgerMarker for RAID=%d, RID=%d, on or before %s, error: %s\n",
			funcname, raid, rid, dt.Format(RRDATEFMT3), err.Error())
		return bal, err
	}

	//------------------------------------------------------------------
	// Get all the assessments and payments for this RAID, RID pair...
	//------------------------------------------------------------------
	b, err := GetRARAcctRange(ctx, bid, raid, rid, &lm.Dt, dt)
	if err != nil {
		return bal, err
	}
	bal = lm.Balance + b
	return bal, err
}

// GetRARAcctRange returns the change in balance for the supplie RAID,RID
// combination over the supplied time range.
//
// INPUTS
//   raid - RAID of Rental Agreement
//   rid  - RID of Rentable
//   d1   - time for which balance is requested
//   d2   - time for which balance is requested
//
// RETURNS
//   float64 - the balance for the Rentable rid in Rental Agreement raid at
//             time dt
//   error   - any error encountered
//-----------------------------------------------------------------------------
func GetRARAcctRange(ctx context.Context, bid, raid, rid int64, d1, d2 *time.Time) (float64, error) {
	const funcname = "GetRARAcctRange"
	// Console("Entered %s\n", funcname)
	var (
		bal float64
		err error
	)

	acctRules := ""
	rcvAccts, err := AcctSlice(bid, AccountsReceivable)
	if err != nil {
		LogAndPrintError(funcname, err)
		return bal, err
	}
	if len(rcvAccts) == 0 {
		LogAndPrintError(funcname, fmt.Errorf("GetRARAcctRange: there are no accounts of type %s", AccountsReceivable))
		return bal, err
	}
	qryAccts, err := AcctRulesSlice(rcvAccts, bid, AccountsReceivable)
	if err != nil {
		LogAndPrintError(funcname, err)
		return bal, err
	}

	l := len(qryAccts)
	if 0 > l {
		acctRules = " AND ("
		for i := 0; i < l; i++ {
			acctRules += fmt.Sprintf("ARID=%d", qryAccts[i])
			if i+1 < l {
				acctRules += " OR "
			}
		}
		acctRules += ")"
	}

	q := `
	SELECT
		{{.SelectClause}}
	FROM Assessments
	WHERE
		(RentCycle=0 OR (RentCycle>0 AND PASMID>0)) AND
		RAID={{.RAID}} AND RID={{.RID}} AND
		Stop>="{{.d1}}" AND Start<"{{.d2}}" {{.AcctRules}};`

	qc := QueryClause{
		"SelectClause": RRdb.DBFields["Assessments"],
		"RAID":         strconv.FormatInt(raid, 10),
		"RID":          strconv.FormatInt(rid, 10),
		"d1":           d1.Format(RRDATEFMTSQL),
		"d2":           d2.Format(RRDATEFMTSQL),
		"AcctRules":    acctRules,
	}

	qry := RenderSQLQuery(q, qc)
	rows, err := RRdb.Dbrr.Query(qry)
	if err != nil {
		return bal, err
	}
	defer rows.Close()

	// Console("GetRARAcctRange: query = %s\n", q)

	//------------------------------------------------------------------------
	// Total all assessments in the supplied range that involve RID in RAID.
	//------------------------------------------------------------------------
	for rows.Next() {
		var a Assessment
		err := ReadAssessments(rows, &a)
		if err != nil {
			return bal, err
		}

		if 0 == a.FLAGS&0x4 { // if this is not a reversal...
			bal += a.Amount // ... then add it to the balance
		}
		// Console("\tASMID = %d, FLAGS=%x  Amount = %.2f,  bal = %.2f\n", a.ASMID, a.FLAGS, a.Amount, bal)

		//----------------------------------------------------------------
		// Total all receipts applied toward this ASMID
		//----------------------------------------------------------------
		innerRows, err := RRdb.Prepstmt.GetASMReceiptAllocationsInRARDateRange.Query(raid, a.ASMID, d1, d2)
		if err != nil {
			return bal, err
		}
		defer innerRows.Close()

		for innerRows.Next() {
			var ra ReceiptAllocation
			err = ReadReceiptAllocations(innerRows, &ra)
			if err != nil {
				return bal, err
			}

			bal -= ra.Amount
			// Console("\tRCPAID = %d, Amount = %.2f,  bal = %.2f\n", ra.RCPAID, ra.Amount, bal)
		}

		err = innerRows.Err()
		if err != nil {
			return bal, err
		}
	}
	// Console("---------->>>>> RETURNING BALANCE = %.2f\n", bal)
	return bal, rows.Err()
}
