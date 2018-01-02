package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRReceiptsTable generates a gotable Table object
// contains of all rlib.Receipt related with business
func RRReceiptsTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRReceiptsTable"
	const (
		Date    = 0
		RCPTID  = iota
		PRCPTID = iota
		PMTID   = iota
		DocNo   = iota
		Amount  = iota
		Comment = iota
		// AccountRule = iota
	)

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
	tbl.AddColumn("Comment", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	// tbl.AddColumn("Account Rule", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err := TableReportHeaderBlock(&tbl, "Tendered Payment Log", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	m := rlib.GetReceipts(ri.Bid, &ri.D1, &ri.D2)
	for _, a := range m {
		tbl.AddRow()
		tbl.Putd(-1, Date, a.Dt)
		tbl.Puts(-1, RCPTID, a.IDtoString())
		tbl.Puts(-1, PRCPTID, rlib.IDtoString("RCPT", a.PRCPTID))
		tbl.Puts(-1, PMTID, rlib.IDtoString("PMT", a.PMTID))
		tbl.Puts(-1, DocNo, a.DocNo)
		tbl.Putf(-1, Amount, a.Amount)
		tbl.Puts(-1, Comment, a.Comment)
		// tbl.Puts(-1, 6, rlib.GetReceiptAccountRuleText(&a))
	}
	tbl.TightenColumns()
	tbl.AddLineAfter(len(tbl.Row) - 1)
	tbl.InsertSumRow(len(tbl.Row), 0, len(tbl.Row)-1, []int{Amount}) // insert @ len essentially adds a row.  Only want to sum Amount column
	return tbl
}

// RRreportReceipts generates a text report based on RRReceiptsTable
func RRreportReceipts(ri *ReporterInfo) string {
	// ri.D1 = time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	// ri.D2 = time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
	tbl := RRReceiptsTable(ri)
	return ReportToString(&tbl, ri)
}
