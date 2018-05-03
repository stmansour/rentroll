package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
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
	tbl.AddColumn("Task", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PreDueDate", 20, gotable.CELLDATETIME, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PreDoneDate", 20, gotable.CELLDATETIME, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PreApprovedBy", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("DueDate", 20, gotable.CELLDATETIME, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ApprovedBy", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

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
	// tbl.SetSection1(s)

	m, err := rlib.GetTasks(ctx, ri.ID)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		tbl.Puts(-1, Status, "tbd")
		tbl.Puts(-1, eTask, m[i].Name)
		tbl.Putdt(-1, PreDueDate, m[i].DtPreDue)
		tbl.Puts(-1, PreApprovedBy, rlib.GetNameForUID(ctx, m[i].PreDoneUID))
		tbl.Putdt(-1, DueDate, m[i].DtDue)
		tbl.Puts(-1, ApprovedBy, rlib.GetNameForUID(ctx, m[i].DoneUID))
	}
	return tbl
}

// TaskListReport generates a text-based report based on TaskListReportTable table object
func TaskListReport(ctx context.Context, ri *ReporterInfo) string {
	tbl := TaskListReportTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
