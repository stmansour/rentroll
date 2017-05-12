package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRReceiptsTable generates a gotable Table object
// contains of all rlib.Receipt related with business
func RRReceiptsTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRReceiptsTable"

	// init and prepare some values before table init
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true
	ri.BlankLineAfterRptName = true

	// table init
	tbl := getRRTable()

	tbl.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RCPTID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Parent RCPTID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PMTID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Doc No", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Amount", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Account Rule", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Comment", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err := TableReportHeaderBlock(&tbl, "Receipts", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	m := rlib.GetReceipts(ri.Bid, &ri.D1, &ri.D2)
	for _, a := range m {
		tbl.AddRow()
		tbl.Putd(-1, 0, a.Dt)
		tbl.Puts(-1, 1, a.IDtoString())
		tbl.Puts(-1, 2, rlib.IDtoString("RCPT", a.PRCPTID))
		tbl.Puts(-1, 3, rlib.IDtoString("PMT", a.PMTID))
		tbl.Puts(-1, 4, a.DocNo)
		tbl.Putf(-1, 5, a.Amount)
		tbl.Puts(-1, 6, rlib.GetReceiptAccountRuleText(&a))
		tbl.Puts(-1, 7, a.Comment)
	}
	tbl.TightenColumns()
	return tbl
}

// RRreportReceipts generates a text report based on RRReceiptsTable
func RRreportReceipts(ri *ReporterInfo) string {
	// ri.D1 = time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	// ri.D2 = time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
	tbl := RRReceiptsTable(ri)
	return ReportToString(&tbl, ri)
}
