package rlib

import (
	"context"
	"fmt"
	"time"
)

// CreateTaskListInstance creates a new task list based on the supplied
// definition and Epoch Date.
//
// INPUTS
//  ctx    - context for database transactions
//  TLDID  - Task List Definition ID
//  pivot  - date on or after which the instance will be created
//
// RETURNS
//  error  - any error encountered
//
//-----------------------------------------------------------------------------
func CreateTaskListInstance(ctx context.Context, TLDID int64, pivot *time.Time) (int64, error) {
	var tlid = int64(0)
	tld, err := GetTaskListDefinition(ctx, TLDID)
	if err != nil {
		return tlid, err
	}

	//------------------------------------------------------
	//  First, compute the dates needed for the TaskList
	//------------------------------------------------------
	var tl TaskList
	tl.BID = tld.BID
	tl.Name = tld.Name
	tl.Cycle = tld.Cycle
	tl.FLAGS = tld.FLAGS
	tl.DtDue, err = NextInstanceDate(&tld.EpochDue, pivot, tld.Cycle)
	if err != nil {
		return tlid, err
	}
	tl.DtPreDue, err = NextInstanceDate(&tld.EpochPreDue, pivot, tld.Cycle)
	if err != nil {
		return tlid, err
	}

	//----------------------------------------------------
	// Create the new TaskList
	//----------------------------------------------------
	err = InsertTaskList(ctx, &tl)
	if err != nil {
		return tlid, err
	}
	tlid = tl.TLID

	//----------------------------------------------------
	// Get the associated tasks...
	//----------------------------------------------------
	tds, err := GetTaskListDescriptors(ctx, tld.TLDID)
	if err != nil {
		return tlid, err
	}

	Console("Found tld.TLDID = %d, TaskCount = %d, name = %s\n", tld.TLDID, len(tds), tld.Name)
	for i := 0; i < len(tds); i++ {
		var t Task
		t.DtDue, err = NextInstanceDate(&tds[i].EpochDue, pivot, tld.Cycle)
		if err != nil {
			return tlid, err
		}
		t.DtPreDue, err = NextInstanceDate(&tds[i].EpochPreDue, pivot, tld.Cycle)
		if err != nil {
			return tlid, err
		}
		// Console("%2d. %s, DtDue: %s, DtPreDue: %s\n", i, tds[i].Name, t.DtDue.Format(RRDATEREPORTFMT), t.DtPreDue.Format(RRDATEREPORTFMT))
		t.Name = tds[i].Name
		t.Worker = tds[i].Worker
		t.FLAGS = tds[i].FLAGS
		t.TLID = tl.TLID
		t.BID = tld.BID

		if err = InsertTask(ctx, &t); err != nil {
			return tlid, err
		}
	}

	return tlid, nil
}

// NextInstanceDate computes the next instance date after the pivot
// based on the supplied frequency
//
// INPUTS
//  ctx    - context for database transactions
//  TLDID  - Task List Definition ID
//  pivot  - date on or after which the instance will be created
//
// RETURNS
//  error  - any error encountered
//
//-----------------------------------------------------------------------------
func NextInstanceDate(epoch, pivot *time.Time, freq int64) (time.Time, error) {
	switch freq {
	case CYCLENORECUR:
		return *epoch, nil
	case CYCLESECONDLY:
	case CYCLEMINUTELY:
	case CYCLEHOURLY:
	case CYCLEDAILY:
	case CYCLEWEEKLY:
	case CYCLEMONTHLY:
		month := pivot.Month()
		year := pivot.Year()
		epochday := epoch.Day()
		if pivot.Day() > epochday {
			if month == time.December {
				month = time.January
				year++
			} else {
				month++
			}
		}
		day := epochday
		if day >= 28 {
			day = LastDOM(month, year)
		}
		dt := time.Date(year, month, day, epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), time.UTC)
		return dt, nil
	case CYCLEQUARTERLY:
	case CYCLEYEARLY:
	default:
		return *epoch, fmt.Errorf("Unrecognized recur cycle: %d", freq)
	}
	return *epoch, nil
}

// LastDOM computes and returns the last day of the supplied month & year
//
// INPUTS
//  month  - time.January - time.December
//  year   - year
//
// RETURNS
//  error  - any error encountered
//
//-----------------------------------------------------------------------------
func LastDOM(m time.Month, y int) int {
	if m == time.December {
		m = time.January
		y++
	} else {
		m++
	}
	dt := time.Date(y, m, 0, 0, 0, 0, 0, time.UTC)
	return dt.Day()
}
