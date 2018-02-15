package worker

import (
	"rentroll/rlib"
	"time"
	"tws"
)

// CleanSecDepBalanceCache is a worker that that cleans the Security Deposit
// Balance cache.
//-----------------------------------------------------------------------------
func CleanSecDepBalanceCache(item *tws.Item) {
	tws.ItemWorking(item) // inform the tws system that we're working

	rlib.CleanSecDepBalanceInfoCache(false) // false means don't remove everything, remove only if expire time is < Now()

	rlib.PrintSecDepBalCache()
	// reschedule after the caches default time to live...
	resched := time.Now().Add(rlib.SecDepBalCacheCtx.Expiry)
	tws.RescheduleItem(item, resched)
}
