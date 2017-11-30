package worker

import (
	"rentroll/rlib"
	"time"
	"tws"
)

// CleanRARBalanceCache is a worker that that cleans the RAR Balance cache.
//-----------------------------------------------------------------------------
func CleanRARBalanceCache(item *tws.Item) {
	tws.ItemWorking(item) // inform the tws system that we're working

	rlib.CleanRARBalanceInfoCache(false) // false means don't remove everything, remove only if expire time is < Now()

	// reschedule after the caches default time to live...
	resched := time.Now().Add(rlib.RARBalCacheCtx.Expiry)
	tws.RescheduleItem(item, resched)
}
