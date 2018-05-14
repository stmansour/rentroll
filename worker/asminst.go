package worker

import (
	"context"
	"rentroll/rlib"
	"time"
	"tws"
)

// CreateAssessmentInstances is a worker that is called by TWS periodically to
// check for recurring assessments that have instances needing to be created.
// When their instance date arrives, this routine will generate the new instance.
// After generating all instances whose time has arrived it will reschedule itself
// to be called again the next day.
//-----------------------------------------------------------------------------
func CreateAssessmentInstances(item *tws.Item) {
	tws.ItemWorking(item)
	now := time.Now()
	ctx := context.Background()
	CreateAsmInstCore(ctx, &now)

	// reschedule for midnight tomorrow...
	resched := now.AddDate(0, 0, 1)
	tws.RescheduleItem(item, resched)
}

// CreateAsmInstCore provides a more testable calling routine for creating
// assessment instances
//-----------------------------------------------------------------------------
func CreateAsmInstCore(ctx context.Context, now *time.Time) {
	expire := now.Add(10 * time.Minute)
	s := rlib.SessionNew("BotToken-"+WorkerAsmtDes, WorkerAsmtDes, WorkerAsmtDes, WorkerAsmt, "", -1, &expire)
	ctx = rlib.SetSessionContextKey(ctx, s)

	// add any new recurring instances for this day...
	m, err := rlib.GetAllBusinesses(ctx)
	if err != nil {
		rlib.Ulog("Error with rlib.GetAllBusinesses: %s\n", err.Error())
	} else {
		// rlib.Console("got businesses: len(m) = %d\n", len(m))
		d1, d2 := rlib.GetMonthPeriodForDate(now)
		for i := 0; i < len(m); i++ {
			// rlib.Console("PROCESS JOURNAL ENTRIES FOR BIZ: %s - %s,  %s, %s\n", m[i].Designation, m[i].Name, d1.Format(rlib.RRDATEREPORTFMT), d2.Format(rlib.RRDATEREPORTFMT))
			var xbiz rlib.XBusiness
			err = rlib.GetXBusiness(ctx, m[i].BID, &xbiz)
			if err != nil {
				rlib.Ulog("Error with rlib.GetXBusiness: %s\n", err.Error())
			}

			err = rlib.GenerateRecurInstances(ctx, &xbiz, &d1, &d2)
			if err != nil {
				rlib.Ulog("Error with rlib.GenerateRecurInstances: %s\n", err.Error())
			}
		}
	}
}
