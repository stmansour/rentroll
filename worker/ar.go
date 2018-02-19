package worker

import (
	"rentroll/rlib"
	"time"
	"tws"
)

// CleanARSliceCache is a worker that that cleans the Account Rule cache.
//-----------------------------------------------------------------------------
func CleanARSliceCache(item *tws.Item) {
	tws.ItemWorking(item) // inform the tws system that we're working

	rlib.CleanARSliceCache(false) // false means don't remove everything, remove only if expire time is < Now()

	// reschedule after the caches default time to live...
	resched := time.Now().Add(rlib.ARCacheCtx.Expiry)
	tws.RescheduleItem(item, resched)
}
