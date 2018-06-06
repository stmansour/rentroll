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
//  PTLID  - Parent (first in chain) for this recurring series.
//  pivot  - date on or after which the instance will be created
//
// RETURNS
//  error  - any error encountered
//
//-----------------------------------------------------------------------------
func CreateTaskListInstance(ctx context.Context, TLDID, PTLID int64, pivot *time.Time) (TaskList, error) {
	var tl TaskList
	tld, err := GetTaskListDefinition(ctx, TLDID)
	if err != nil {
		return tl, err
	}

	//------------------------------------------------------
	//  First, compute the dates needed for the TaskList
	//------------------------------------------------------
	tl.BID = tld.BID
	tl.PTLID = PTLID
	tl.TLDID = tld.TLDID
	tl.Name = tld.Name
	tl.Cycle = tld.Cycle
	tl.FLAGS = tld.FLAGS
	if err = NextTLInstanceDates(pivot, &tld, &tl); err != nil {
		return tl, err
	}
	tl.EmailList = tld.EmailList
	tl.DurWait = tld.DurWait
	if time.Duration(0) == tl.DurWait {
		tl.DurWait = 24 * time.Hour
	}
	if err != nil {
		return tl, err
	}

	//----------------------------------------------------
	// Create the new TaskList
	//----------------------------------------------------
	err = InsertTaskList(ctx, &tl)
	if err != nil {
		return tl, err
	}

	//----------------------------------------------------
	// Get the associated tasks...
	//----------------------------------------------------
	tds, err := GetTaskListDescriptors(ctx, tld.TLDID)
	if err != nil {
		return tl, err
	}

	for i := 0; i < len(tds); i++ {
		var t Task
		if err = NextTaskInstanceDates(pivot, &tld, &tds[i], &t); err != nil {
			return tl, err
		}
		t.Name = tds[i].Name
		t.Worker = tds[i].Worker
		t.FLAGS = tds[i].FLAGS
		t.TLID = tl.TLID
		t.BID = tld.BID

		if err = InsertTask(ctx, &t); err != nil {
			return tl, err
		}
	}

	return tl, nil
}

// NextTLInstanceDates computes the next instance dates after the pivot
// based on the supplied frequency
//
// INPUTS
//  pivot  - date on or after which the instance will be created
//  tld    - pointer to Task List Definition struct
//  tl     - ptr to Task List
//
// RETURNS
//  error  - any error encountered
//-----------------------------------------------------------------------------
func NextTLInstanceDates(pivot *time.Time, tld *TaskListDefinition, tl *TaskList) error {
	switch tld.Cycle {
	case RECURNONE:
		tl.DtDue = tld.EpochDue
		tl.DtPreDue = tld.EpochPreDue
	case RECURSECONDLY:
	case RECURMINUTELY:
	case RECURHOURLY:
	case RECURDAILY:
		deltap := time.Duration(0)
		if tld.EpochPreDue.Year() > 1999 && tld.Epoch.Year() > 1999 {
			deltap = tld.EpochPreDue.Sub(tld.Epoch)
		}
		tl.DtPreDue = pivot.Add(deltap)

		deltad := time.Duration(0)
		if tld.EpochDue.Year() > 1999 && tld.Epoch.Year() > 1999 {
			deltad = tld.EpochDue.Sub(tld.Epoch)
		}
		tl.DtDue = pivot.Add(deltad)

		// Console("\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n")
		// Console("PIVOT = %s\n", pivot.Format(RRJSUTCDATETIME))
		// Console("DtDue = %s\n", tl.DtDue.Format(RRJSUTCDATETIME))
		// Console("after adjustment -   DtDue = %s\n", tl.DtDue.Format(RRJSUTCDATETIME))
		// Console("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n\n")

	case RECURWEEKLY:
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
		if tl.DtDue.Before(dtPivot) {
			tl.DtDue = tl.DtDue.AddDate(0, 0, 7)
			tl.DtPreDue = tl.DtPreDue.AddDate(0, 0, 7)
		}

		// Console("\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n")
		// Console("Epoch = %s, nepoch = %s\n", tld.Epoch.Format(RRJSUTCDATETIME), newepoch.Format(RRJSUTCDATETIME))
		// Console("edow = %d, edowDue = %d, edowPreDue = %d\n", edow, edowDue, edowPreDue)
		// Console("tl.DtDue = %s + %d days = %s\n", newepoch.Format(RRJSUTCDATETIME), edowDue-edow, tl.DtDue.Format(RRJSUTCDATETIME))
		// Console("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n\n")

	case RECURMONTHLY:
		tl.DtPreDue = computeMonthlyDate(pivot, &tld.EpochPreDue)
		tl.DtDue = computeMonthlyDate(pivot, &tld.EpochDue)
		if tl.DtPreDue.After(tl.DtDue) {
			offset := int(tl.DtDue.Sub(tld.EpochDue).Hours() / 24)
			tl.DtPreDue = tld.EpochPreDue.AddDate(0, 0, offset)
		}
	case RECURQUARTERLY:
	case RECURYEARLY:
		tlepoch := time.Date(pivot.Year(), tld.Epoch.Month(), tld.Epoch.Day(), tld.Epoch.Hour(), tld.Epoch.Minute(), 0, 0, time.UTC)
		deltap := time.Duration(0)
		if tld.EpochPreDue.Year() > 1999 && tld.Epoch.Year() > 1999 {
			deltap = tld.EpochPreDue.Sub(tld.Epoch)
		}
		tl.DtPreDue = tlepoch.Add(deltap)

		deltad := time.Duration(0)
		if tld.EpochDue.Year() > 1999 && tld.Epoch.Year() > 1999 {
			deltad = tld.EpochDue.Sub(tld.Epoch)
		}
		tl.DtDue = tlepoch.Add(deltad)

	default:
		return fmt.Errorf("Unrecognized recur cycle: %d", tld.Cycle)
	}
	return nil
}

// NextTaskInstanceDates computes the next instance dates after the pivot
// based on the supplied frequency
//
// INPUTS
//  pivot  - date on or after which the instance will be created
//  tzoff  - user's client timezone offset in minutes.  For web browsers
//           this is determined by:
//               offset = new Date().getTimezoneOffset();
//  tld    - pointer to Task List Definition struct
//  td     - ptr to Task definition
//  t      - the task
//
// RETURNS
//  error  - any error encountered
//
//-----------------------------------------------------------------------------
func NextTaskInstanceDates(pivot *time.Time, tld *TaskListDefinition, td *TaskDescriptor, t *Task) error {
	switch tld.Cycle {
	case RECURNONE:
		t.DtDue = td.EpochDue
		t.DtPreDue = tld.EpochPreDue
	case RECURSECONDLY:
	case RECURMINUTELY:
	case RECURHOURLY:
	case RECURDAILY:
		t.DtDue = time.Date(pivot.Year(), pivot.Month(), pivot.Day(), td.EpochDue.Hour(), td.EpochDue.Minute(), 0, 0, time.UTC)
		t.DtPreDue = time.Date(pivot.Year(), pivot.Month(), pivot.Day(), td.EpochPreDue.Hour(), td.EpochPreDue.Minute(), 0, 0, time.UTC)

	case RECURWEEKLY:
		dtEpoch := time.Date(tld.Epoch.Year(), tld.Epoch.Month(), tld.Epoch.Day(), 0, 0, 0, 0, time.UTC)
		dtPivot := time.Date(pivot.Year(), pivot.Month(), pivot.Day(), 0, 0, 0, 0, time.UTC)
		offset := int(dtPivot.Sub(dtEpoch).Hours() / (7 * 24)) // number of weeks difference
		edow := int(tld.Epoch.Weekday())                       // on which day does the definition start
		edowDue := int(td.EpochDue.Weekday())                  // what day is due
		edowPreDue := int(td.EpochPreDue.Weekday())            // what day is pre-due
		newepoch := tld.Epoch.AddDate(0, 0, 7*offset)          // snap to nearest week to pivot
		dt := newepoch.AddDate(0, 0, edowDue-edow)
		t.DtDue = time.Date(dt.Year(), dt.Month(), dt.Day(), td.EpochDue.Hour(), td.EpochDue.Minute(), 0, 0, time.UTC)
		dt = newepoch.AddDate(0, 0, edowPreDue-edow)
		t.DtPreDue = time.Date(dt.Year(), dt.Month(), dt.Day(), td.EpochPreDue.Hour(), td.EpochPreDue.Minute(), 0, 0, time.UTC)
		if t.DtDue.Before(dtPivot) {
			t.DtDue = t.DtDue.AddDate(0, 0, 7)
			t.DtPreDue = t.DtPreDue.AddDate(0, 0, 7)
		}

	case RECURMONTHLY:
		t.DtPreDue = computeMonthlyDate(pivot, &td.EpochPreDue)
		t.DtDue = computeMonthlyDate(pivot, &td.EpochDue)
		if t.DtPreDue.After(t.DtDue) {
			// t.DtPreDue = time.Date(t.DtDue.Year(), t.DtDue.Month(), t.DtPreDue.Day(), t.DtPreDue.Hour(), t.DtPreDue.Minute(), 0, 0, time.UTC)
			// Days Between Dates (DaysBetweenDates) https://play.golang.org/p/nkBPjPumg6-
			offset := int(t.DtDue.Sub(td.EpochDue).Hours() / 24)
			Console("offset = %d days\n", offset)
			t.DtPreDue = tld.EpochPreDue.AddDate(0, 0, offset)
		}
	case RECURQUARTERLY:
	case RECURYEARLY:
		tlepoch := time.Date(pivot.Year(), tld.Epoch.Month(), tld.Epoch.Day(), tld.Epoch.Hour(), tld.Epoch.Minute(), 0, 0, time.UTC)
		deltap := time.Duration(0)
		if td.EpochPreDue.Year() > 1999 && tld.Epoch.Year() > 1999 {
			deltap = td.EpochPreDue.Sub(tld.Epoch)
		}
		t.DtPreDue = tlepoch.Add(deltap)

		deltad := time.Duration(0)
		if td.EpochDue.Year() > 1999 && tld.Epoch.Year() > 1999 {
			deltad = td.EpochDue.Sub(tld.Epoch)
		}
		t.DtDue = tlepoch.Add(deltad)
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
