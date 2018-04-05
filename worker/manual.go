package worker

import (
	// "context"
	// "rentroll/rlib"
	"rentroll/rlib"
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
	now := time.Now().In(rlib.RRdb.Zone)

	resched := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC).In(rlib.RRdb.Zone)
	tws.RescheduleItem(item, resched)

}
