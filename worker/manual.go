package worker

import (
	// "context"
	// "rentroll/rlib"

	"time"
	"tws"
)

// ProcessManualTask processes a new instance of a manual task
//
// INPUTS
//  item   - the tws.Item struct
//
// RETURNS
//  nothing
//-------------------------------------------------------------------
func ProcessManualTask(item *tws.Item) {
	tws.ItemWorking(item)
	// reschedule for midnight tomorrow...
	now := time.Now()

	resched := now.AddDate(0, 0, 1)
	tws.RescheduleItem(item, resched)

}
