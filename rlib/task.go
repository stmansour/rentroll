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
	tl.PTLID = 0 // explicitly set this to be the parent
	tl.TLDID = tld.TLDID
	tl.Name = tld.Name
	tl.Cycle = tld.Cycle
	tl.FLAGS = tld.FLAGS
	if err = NextTLInstanceDates(pivot, &tld, &tl); err != nil {
		return tlid, err
	}
	tl.EmailList = tld.EmailList
	tl.DurWait = tld.DurWait
	if time.Duration(0) == tl.DurWait {
		tl.DurWait = 24 * time.Hour
	}
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
		if err = NextTaskInstanceDates(pivot, &tld, &t); err != nil {
			return tlid, err
		}
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

// NextTLInstanceDates computes the next instance dates after the pivot
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
func NextTLInstanceDates(pivot *time.Time, tld *TaskListDefinition, tl *TaskList) error {
	switch tld.Cycle {
	case CYCLENORECUR:
	case CYCLESECONDLY:
	case CYCLEMINUTELY:
	case CYCLEHOURLY:
	case CYCLEDAILY:
	case CYCLEWEEKLY:
		dtEpoch := time.Date(tld.Epoch.Year(), tld.Epoch.Month(), tld.Epoch.Day(), 0, 0, 0, 0, time.UTC)
		dtPivot := time.Date(pivot.Year(), pivot.Month(), pivot.Day(), 0, 0, 0, 0, time.UTC)
		offset := int(dtPivot.Sub(dtEpoch).Hours() / (7 * 24)) // number of weeks difference
		edow := int(tld.Epoch.Weekday())                       // on which day does the definition start
		edowDue := int(tld.EpochDue.Weekday())                 // what day is due
		edowPreDue := int(tld.EpochPreDue.Weekday())           // what day is pre-due
		newepoch := tld.Epoch.AddDate(0, 0, 7*offset)          // snap to nearest week to pivot
		dt := newepoch.AddDate(0, 0, edowDue-edow)
		tl.DtDue = time.Date(dt.Year(), dt.Month(), dt.Day(), tld.EpochDue.Hour(), tld.EpochDue.Minute(), 0, 0, time.UTC)
		dt = newepoch.AddDate(0, 0, edowPreDue-edow)
		tl.DtPreDue = time.Date(dt.Year(), dt.Month(), dt.Day(), tld.EpochPreDue.Hour(), tld.EpochPreDue.Minute(), 0, 0, time.UTC)

		// Console("\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n")
		// Console("Epoch = %s, nepoch = %s\n", tld.Epoch.Format(RRJSUTCDATETIME), newepoch.Format(RRJSUTCDATETIME))
		// Console("edow = %d, edowDue = %d, edowPreDue = %d\n", edow, edowDue, edowPreDue)
		// Console("tl.DtDue = %s + %d days = %s\n", newepoch.Format(RRJSUTCDATETIME), edowDue-edow, tl.DtDue.Format(RRJSUTCDATETIME))
		// Console("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n\n")

	case CYCLEMONTHLY:
		tl.DtPreDue = computeMonthlyDate(pivot, &tld.EpochPreDue)
		tl.DtDue = computeMonthlyDate(pivot, &tld.EpochDue)
		if tl.DtPreDue.After(tl.DtDue) {
			// t.DtPreDue = time.Date(t.DtDue.Year(), t.DtDue.Month(), t.DtPreDue.Day(), t.DtPreDue.Hour(), t.DtPreDue.Minute(), 0, 0, time.UTC)
			// Days Between Dates (DaysBetweenDates) https://play.golang.org/p/nkBPjPumg6-
			offset := int(tl.DtDue.Sub(tld.EpochDue).Hours() / 24)
			Console("offset = %d days\n", offset)
			tl.DtPreDue = tld.EpochPreDue.AddDate(0, 0, offset)
		}
	case CYCLEQUARTERLY:
	case CYCLEYEARLY:
	default:
		return fmt.Errorf("Unrecognized recur cycle: %d", tld.Cycle)
	}
	return nil
}

// NextTaskInstanceDates computes the next instance dates after the pivot
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
func NextTaskInstanceDates(pivot *time.Time, tld *TaskListDefinition, t *Task) error {
	switch tld.Cycle {
	case CYCLENORECUR:
	case CYCLESECONDLY:
	case CYCLEMINUTELY:
	case CYCLEHOURLY:
	case CYCLEDAILY:
	case CYCLEWEEKLY:
		edow := int(tld.Epoch.Weekday()) // on which day does the definition start
		pdow := int(pivot.Weekday())     // what's our pivot day
		if pdow != edow {
			t.DtPreDue = t.DtPreDue.AddDate(0, 0, edow-pdow)
			t.DtDue = t.DtDue.AddDate(0, 0, edow-pdow)
		}
	case CYCLEMONTHLY:
		t.DtPreDue = computeMonthlyDate(pivot, &t.DtPreDue)
		t.DtDue = computeMonthlyDate(pivot, &t.DtDue)
	case CYCLEQUARTERLY:
	case CYCLEYEARLY:
	default:
		return fmt.Errorf("Unrecognized recur cycle: %d", tld.Cycle)
	}
	return nil
}

func computeMonthlyDate(pivot, epoch *time.Time) time.Time {
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
	return time.Date(year, month, day, epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), time.UTC)

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
