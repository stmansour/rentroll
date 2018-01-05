package rrpt

import (
	"context"
	"gotable"
	"rentroll/rlib"
)

// RRreportDepositMethodsTable generates a table for all rlib.DepositMethod
func RRreportDepositMethodsTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportDepositMethodsTable"
	var (
		err error
	)

	// table init
	tbl := getRRTable()

	tbl.AddColumn("DPMID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// prepare table's title, sections
	err = TableReportHeaderBlock(ctx, &tbl, "Deposit Methods", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	m, err := rlib.GetAllDepositMethods(ctx, ri.Bid)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		tbl.Puts(-1, 0, rlib.IDtoString("DPM", m[i].DPMID))
		tbl.Puts(-1, 1, rlib.IDtoString("B", m[i].BID))
		tbl.Puts(-1, 2, m[i].Method)
	}

	tbl.TightenColumns()
	return tbl
}

// RRreportDepositMethods generates a report of all rlib.DepositMethod
func RRreportDepositMethods(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportDepositMethodsTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
