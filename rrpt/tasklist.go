package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"time"
)

// TaskListTextReport generates a status report for the supplied TaskList
//
// INPUTS
//  ctx - context for transactions
//  ri  - Repoter Info.  ri.ID in this case is the TLID of the tasklist to be
//        printed
//
// RETURNS
//
//-----------------------------------------------------------------------------
func TaskListTextReport(ctx context.Context, ri *ReporterInfo) {
	tbl := TaskListReportTable(ctx, ri)
	fmt.Print(tbl)
}

// TaskListReportTable generates a status report for the supplied TaskList
//
// INPUTS
//  ctx - context for transactions
//  ri  - Repoter Info.  ri.ID in this case is the TLID of the tasklist to be
//        printed
//
// RETURNS
//  the gotable containing the report data
//-----------------------------------------------------------------------------
func TaskListReportTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "TaskListReportTable"
	var err error
	var now = time.Now()

	const (
		Status        = 0
		eTask         = iota
		PreDueDate    = iota
		PreDoneDate   = iota
		PreApprovedBy = iota
		DueDate       = iota
		DoneDate      = iota
		ApprovedBy    = iota
	)

	// init and prepare some values before table init
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = false

	// table init
	tbl := getRRTable()

	tbl.AddColumn("Status", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Task", 40, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PreDueDate", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PreDoneDate", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PreApprovedBy", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("DueDate", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("DoneDate", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ApprovedBy", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	tl, err := rlib.GetTaskList(ctx, ri.ID)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}
	// prepare table's title, sections
	err = TableReportHeaderBlock(ctx, &tbl, tl.Name, funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	// overwrite a few things
	tbl.SetTitle(tl.Name)
	s := fmt.Sprintf("Due %s", tl.DtDue.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT))
	tbl.SetSection1(s)

	if tl.DoneUID > 0 {
		s = fmt.Sprintf("Completed %s by %s", tl.DtDone.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT), rlib.GetNameForUID(ctx, tl.DoneUID))
		if tl.DtDone.After(tl.DtDue) && tl.DtDue.Year() > 1900 {
			s = "LATE, " + s
		}
	} else {
		s = "Not Yet Completed"
		if now.After(tl.DtDue) && tl.DtDue.Year() > 1900 {
			s = "OVERDUE, " + s
		}
	}
	tbl.SetSection2(s)

	m, err := rlib.GetTasks(ctx, ri.ID)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		st := " "
		if m[i].DtDue.Year() > 1970 { // is there a due date?
			if now.After(m[i].DtDue) {
				st = "LATE"
			}
			if m[i].DtDone.Year() > 1970 { // is there a done date?
				if !m[i].DtDone.After(m[i].DtDue) {
					st = "+"
				}
			}
		}
		tbl.Puts(-1, Status, st)
		tbl.Puts(-1, eTask, m[i].Name)

		tbl.Puts(-1, PreDueDate, m[i].DtPreDue.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT))
		tbl.Puts(-1, DueDate, m[i].DtDue.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT))

		if m[i].DtPreDone.Year() > 1970 {
			tbl.Puts(-1, PreDoneDate, m[i].DtPreDone.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT))
		}
		if m[i].PreDoneUID > 0 {
			tbl.Puts(-1, PreApprovedBy, rlib.GetNameForUID(ctx, m[i].PreDoneUID))
		}

		if m[i].DtDone.Year() > 1970 {
			tbl.Puts(-1, DoneDate, m[i].DtDone.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT))
		}
		if m[i].DoneUID > 0 {
			tbl.Puts(-1, ApprovedBy, rlib.GetNameForUID(ctx, m[i].DoneUID))
		}
	}
	return tbl
}

// TaskListReport generates a text-based report based on TaskListReportTable table object
func TaskListReport(ctx context.Context, ri *ReporterInfo) string {
	tbl := TaskListReportTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
