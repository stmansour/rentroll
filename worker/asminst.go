package worker

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"time"
	"tws"
)

// CreateAssessmentInstances is a worker that is called by TWS periodically to
// check for recurring assessments that have instances needing to be created.
// When their instance date arrives, this routine will generate the new instance.
// After generating all instances whose time has arrived it will reschedule itself
// to be called again the next day.
func CreateAssessmentInstances(item *tws.Item) {
	tws.ItemWorking(item)

	// create background context
	ctx := context.Background()

	// create session and set in the context
	// TODO(Steve): we should have some pre-defined sessions or something else
	// and based on that we can define permissions on those pre-defined sessions
	// so that they can access only required information
	// For ex. for command line applications we've to do this!
	expire := time.Now().Add(10 * time.Minute)
	s := rlib.SessionNew("workerToken", "tws-worker", "tws-session", -99, "", -99, &expire)
	ctx = rlib.SetSessionContextKey(ctx, s)

	// add any new recurring instances for this day...
	m, err := rlib.GetAllBusinesses(ctx)
	if err != nil {
		rlib.Ulog("Error with rlib.GetAllBusinesses: %s\n", err.Error())
	} else {
		now := time.Now()
		d1, d2 := rlib.GetMonthPeriodForDate(&now)
		for i := 0; i < len(m); i++ {
			fmt.Printf("PROCESS JOURNAL ENTRIES FOR BIZ: %s - %s\n", m[i].Designation, m[i].Name)
			fmt.Printf("call rlib.GenerateRecurInstances(xbiz, %s, %s)\n", d1.Format(rlib.RRDATEREPORTFMT), d2.Format(rlib.RRDATEREPORTFMT))
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

	// reschedule for midnight tomorrow...
	now := time.Now().In(rlib.RRdb.Zone)
	resched := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC).In(rlib.RRdb.Zone)
	tws.RescheduleItem(item, resched)
}
