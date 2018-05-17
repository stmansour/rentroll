package worker

import (
	"context"
	"rentroll/rlib"
	"time"
	"tws"
)

// This package is responsible for creating TaskList instances for
// recurring task lists when their epoch date arrives

// TLInstanceBot is the worker responsible for creating new instances of
// check for recurring assessments that have instances needing to be created.
// It will create them when their epoch arrives.
//-----------------------------------------------------------------------------
func TLInstanceBot(item *tws.Item) {
	rlib.Ulog("TLInstanceBot\n") // log the fact that we're running

	checkInterval := 24 * time.Hour // this may come from a config file in the future
	tws.ItemWorking(item)
	now := time.Now()
	expire := now.Add(checkInterval)
	s := rlib.SessionNew("BotToken-"+TLReportBotDes, TLReportBotDes, TLReportBotDes, TLReportBot, "", -1, &expire)
	ctx := context.Background()
	ctx = rlib.SetSessionContextKey(ctx, s)
	TLInstanceBotCore(ctx, &now)

	//---------------------------------------------
	// schedule this check again in a few mins...
	//---------------------------------------------
	resched := now.Add(checkInterval)
	tws.RescheduleItem(item, resched)
}

// TLInstanceBotCore provides a more testable calling routine for processing
// Task List Definitions.  This routine checks all active task lists and
// creates a new instance recurring task instances whose epoch time has
// arrived.
//
// INPUTS:
//    ctx  - context which may include a database transaction in progress
//    now  - use this time as the epoch
//
// RETURNS:
//
//-----------------------------------------------------------------------------
func TLInstanceBotCore(ctx context.Context, now *time.Time) error {
	// var err error
	// var rows *sql.Rows

	return nil
}
