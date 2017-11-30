package worker

import (
	"rentroll/rlib"
	"time"
	"tws"
)

// CleanAcctSliceCache is a worker that that cleans the RAR Balance cache.
//-----------------------------------------------------------------------------
func CleanAcctSliceCache(item *tws.Item) {
	tws.ItemWorking(item) // inform the tws system that we're working

	rlib.CleanGLAcctSliceCache(false) // false means don't remove everything, remove only if expire time is < Now()

	// reschedule after the caches default time to live...
	resched := time.Now().Add(rlib.GLAcctCacheCtx.Expiry)
	tws.RescheduleItem(item, resched)
}
