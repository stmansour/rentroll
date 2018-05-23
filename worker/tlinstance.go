package worker

import (
	"context"
	"database/sql"
	"fmt"
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
// Task Lists.  It determines whether the time has arrived for a new recurring
// task list instance to begin.  If so, it creates a new task list instance.
//
//
// see test program: https://play.golang.org/p/M3gQBkqHGyh
// INPUTS:
//    ctx  - context which may include a database transaction in progress
//    now  - use this time as the epoch
//
// RETURNS:
//    any error encountered
//-----------------------------------------------------------------------------
func TLInstanceBotCore(ctx context.Context, now *time.Time) error {
	var err error
	var rows *sql.Rows
	eot, err := rlib.StringToDate("3000-01-01 00:00:00 UTC")
	if err != nil {
		rlib.Ulog("error converting date to date: %s\n", err.Error())
		return err
	}
	rows, err = rlib.RRdb.Prepstmt.GetAllParentTaskLists.Query()
	if err != nil {
		rlib.Ulog("error getting rows cursor: %s\n", err.Error())
		return err
	}

	//-----------------------------------------
	// Collect the rows of interest...
	//-----------------------------------------
	var m []rlib.TaskList
	for i := 0; rows.Next(); i++ {
		var tl rlib.TaskList
		if err = rlib.ReadTaskLists(rows, &tl); err != nil {
			return err
		}
		m = append(m, tl)
	}
	if err = rows.Err(); err != nil {
		return err
	}

	//------------------------------------------------------
	// Process active parents to see if any new instances
	// need to be created now...
	//------------------------------------------------------
	for i := 0; i < len(m); i++ {

		//--------------------------------
		// skip if no due dates...
		//--------------------------------
		if m[i].Cycle == rlib.RECURNONE || m[i].DtPreDue.Year() < 1999 || m[i].DtDue.Year() < 1999 {
			continue
		}
		tld, err := rlib.GetTaskListDefinition(ctx, m[i].TLDID)
		if err != nil {
			return err
		}

		if tld.Epoch.Year() < 1999 {
			return fmt.Errorf("no epoch for TLDID = %d", m[i].TLDID)
		}
		dtNext := now.AddDate(0, 0, 1) // initialize something for dtNext
		switch m[i].Cycle {
		case rlib.RECURDAILY: // daily
			dtNext = now.AddDate(0, 0, 1)
		case rlib.RECURWEEKLY: // weekly
			dtNext = now.AddDate(0, 0, 7)
		case rlib.RECURMONTHLY: // monthly
			dtNext = now.AddDate(0, 1, 0)
		case rlib.RECURQUARTERLY: // quarterly
			dtNext = now.AddDate(0, 3, 0)
		case rlib.RECURYEARLY: // yearly
			dtNext = now.AddDate(1, 0, 0)
		}

		//----------------------------------------------
		// Look for any instances that occur this day
		//----------------------------------------------
		newepoch := getRecurrences(now, &dtNext, &tld.Epoch, &eot, m[i].Cycle)

		//----------------------------------------------------
		// Check for an existing instance on this date/time.
		// We need to make this function idempotent. That is,
		// look for an instance at newepoch where PTLID == m[i].TLID,
		// and if we find one, don't create another.
		//----------------------------------------------------
		dup, err := rlib.GetTaskListInstanceInRange(ctx, m[i].TLID, now, &dtNext)
		if err != nil {
			return err
		}
		if dup.TLID != 0 {
			rlib.LogAndPrint("*** FOUND EXISTING TaskList INSTANCE. TLID = %d, will not create duplicate\n", dup.TLID)
			continue
		}

		//----------------------------------------------
		// now create the new instance...
		//----------------------------------------------
		_, err = rlib.CreateTaskListInstance(ctx, m[i].TLDID, m[i].TLID, &newepoch)
		if err != nil {
			return err
		}
	}

	return nil
}

// getRecurrences returns a list of instance dates where an event time
// (aStart - aStop) overlaps with an interval time (start - stop).  The
// recurrence frequency maps to those that can happen for an assessment.
//
// INPUTS
//     start, stop = time range for the recurrences to be generated. For
//                   example, the start / end time of a period to be
//                   considered.
//   aStart, aStop = imposed bounds, for example the start/stop time of
//                   an Assessment
//       cycleFreq = recurrence frequency
//
// RETURNS
//   an array of instances
//-----------------------------------------------------------------------------
func getRecurrences(start, stop, aStart, aStop *time.Time, cycleFreq int64) time.Time {
	var x time.Time
	//-------------------------------------------
	// now calculate all the recurrences
	//-------------------------------------------
	switch cycleFreq {
	case rlib.RECURNONE: // no recurrence
		if rlib.DateInRange(aStart, start, stop) {
			return *aStart
		}
	case rlib.RECURDAILY: // daily
		d := start.Day()
		if start.Day() < aStart.Day() {
			d = aStart.Day()
		}
		dt := time.Date(start.Year(), start.Month(), d, aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genRegularRecurSeq(&dt, aStop, start, stop, 24*time.Hour)
	case rlib.RECURWEEKLY: // weekly
		dt := time.Date(start.Year(), start.Month(), start.Day(), aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genRegularRecurSeq(&dt, aStop, start, stop, 7*24*time.Hour)
	case rlib.RECURMONTHLY: // monthly
		// dt := time.Date(start.Year(), start.Month(), aStart.Day(), aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genMonthlyRecurSeq(aStart, aStop, start, stop, 1)
	case rlib.RECURQUARTERLY: // quarterly
		dt := time.Date(start.Year(), start.Month(), aStart.Day(), aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genMonthlyRecurSeq(&dt, aStop, start, stop, 3)
	case rlib.RECURYEARLY: // yearly
		dt := time.Date(start.Year(), aStart.Month(), aStart.Day(), aStart.Hour(), aStart.Minute(), aStart.Second(), 0, time.UTC)
		return genYearlyRecurSeq(&dt, start, stop, 1)
	}
	return x
}

// genRegularRecurSeq generates instance dates for the supplied recurring
// frequency.
//
// INPUTS:
//     a1 - a2 = time range of the assessment
//     R1 - R2 = time range for the run calculation
//     freq = chunk of time over which to quantize the assessment
//
// RETURNS:
//     first instance in timeframe R1 - R2
//-----------------------------------------------------------------------------
func genRegularRecurSeq(a1, a2, R1, R2 *time.Time, freq time.Duration) time.Time {
	//============================================
	// Set up first time range for first run...
	//============================================
	d1 := *R1
	d2 := d1.Add(freq)

	//============================================
	// Now iterate in chunks (of size: freq) and
	// save the recurrence dates...
	//============================================
	var m time.Time
	for i := 0; i < 10000; i++ {
		// fmt.Printf("iter: %d\n", i)
		//----------------------------------------
		// don't go outside the requested range
		//----------------------------------------
		if d1.After(*R2) || d1.Equal(*R2) {
			break
		}
		if d2.After(*R2) {
			d2 = *R2
		}
		//----------------------------------------
		//  does a1 - a2 overlap d1-d2
		//----------------------------------------
		// fmt.Printf("d1: %s   d2: %s\n", d1.Format(time.RFC3339), d2.Format(time.RFC3339))
		if !(a1.After(d2) || a1.Equal(d2) || a2.Before(d1) || a2.Equal(d1)) {
			d := time.Date(d1.Year(), d1.Month(), d1.Day(), a1.Hour(), a1.Minute(), a1.Second(), 0, time.UTC)
			// fmt.Printf("STORE d = %s\n", d.Format(time.RFC3339))
			return d
		}
		//----------------------------------------
		//  set the next interval to check...
		//----------------------------------------
		d1 = d1.Add(freq)
		d2 = d2.Add(freq)
	}
	return m
}

// genMonthlyRecurSeq
//
// INPUTS
//     a1,a2 -  epoch
//     R1,R2 -  timerange of interest
//
// RETURNS
//     first instance occurring in time range R1,R2
//-----------------------------------------------------------------------------
func genMonthlyRecurSeq(a1, a2, R1, R2 *time.Time, nMonths int64) time.Time {
	//============================================
	// Set up first time range for first run...
	//============================================
	d1 := *R1
	d2 := d1 // just to define it, it will be set correctly below

	//============================================
	// Now iterate in chunks (of size: nMonths) and
	// save the recurrence dates...
	//============================================
	var m time.Time
	for i := 0; i < 10000; i++ {
		mo, y := rlib.IncMonths(d1.Month(), nMonths)
		d2 = time.Date(d1.Year()+int(y), mo, d1.Day(), d1.Hour(), d1.Minute(), d1.Second(), 0, time.UTC)

		//----------------------------------------
		// don't go outside the requested range
		//----------------------------------------
		if d1.After(*R2) || d1.Equal(*R2) {
			break
		}
		if d2.After(*R2) {
			d2 = *R2
		}
		//-------------------------------------------------------------------------------------
		//  the recurrence date will be the epoch date applied to d1.  Then see if this date
		//  is in the range d1 - d2
		//-------------------------------------------------------------------------------------
		d := time.Date(d1.Year(), d1.Month(), a1.Day(), d1.Hour(), d1.Minute(), d1.Second(), 0, time.UTC)

		// for the date comparison, we are only interested whether the dates fall in line, to the times...
		dt1 := time.Date(d1.Year(), d1.Month(), d1.Day(), 0, 0, 0, 0, time.UTC)
		dt2 := time.Date(d2.Year(), d2.Month(), d2.Day(), 0, 0, 0, 0, time.UTC)
		at1 := time.Date(a1.Year(), a1.Month(), a1.Day(), 0, 0, 0, 0, time.UTC)
		at2 := time.Date(a2.Year(), a2.Month(), a2.Day(), 0, 0, 0, 0, time.UTC)

		//-------------------------------------------------------------------------------------
		// make sure it's in the interval range AND in its active timeframe, and if it is
		// then add it to the list
		//-------------------------------------------------------------------------------------
		// fmt.Printf("genMonthlyRecurSeq: d = %s, d1,d2 = %s , %s    a1,a2 = %s , %s\n", d.Format(RRDATEREPORTFMT), d1.Format(RRDATEREPORTFMT), d2.Format(RRDATEREPORTFMT), a1.Format(RRDATEREPORTFMT), a2.Format(RRDATEREPORTFMT))
		if rlib.DateInRange(&d, &dt1, &dt2) {
			if a1.Equal(d) || rlib.DateInRange(&d, &at1, &at2) {
				dt := time.Date(d.Year(), d.Month(), d.Day(), a1.Hour(), a1.Minute(), 0, 0, time.UTC)
				return dt
			}
		}
		d1 = d2 // on to the next interval
	}

	return m
}

// genMonthlyRecurSeq generates the first yearly instance between
// dates R1-R2
//
// INPUTS:
//     a1,a2 -  epoch
//     R1,R2 -  timerange of interest
//
// RETURNS:
//     the first instance in R1-R2
//-----------------------------------------------------------------------------
func genYearlyRecurSeq(d, start, stop *time.Time, n int64) time.Time {
	var x time.Time
	for i := start.Year(); i <= stop.Year(); i++ {
		dt := time.Date(i, d.Month(), d.Day(), d.Hour(), d.Minute(), 0, 0, time.UTC)
		if rlib.DateInRange(&dt, start, stop) {
			return dt
		}
	}
	return x
}
