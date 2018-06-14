package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// LedgerBalanceReportTable builds a table of trial balance information
func LedgerBalanceReportTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "LedgerBalanceReportTable"
	var (
		err error
	)

	// init and prepare some values before table init
	bid := ri.Xbiz.P.BID
	ri.RptHeaderD2 = true
	ri.BlankLineAfterRptName = true

	// table init
	tbl := getRRTable()

	tbl.AddColumn("GLNumber", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Parent Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)

	// set table title, sections
	err = TableReportHeaderBlock(ctx, &tbl, "Trial Balance", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	for _, acct := range rlib.RRdb.BizTypes[bid].GLAccounts {
		tbl.AddRow()
		tbl.Puts(-1, 0, acct.GLNumber)
		tbl.Puts(-1, 1, acct.Name)
		if acct.AllowPost {
			b, err := rlib.GetAccountBalance(ctx, bid, acct.LID, &ri.D2)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				tbl.SetSection3(err.Error())
				return tbl
			}
			tbl.Putf(-1, 3, b)
		} else {
			b, err := rlib.GetAccountBalance(ctx, bid, acct.LID, &ri.D2)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				tbl.SetSection3(err.Error())
				return tbl
			}
			tbl.Putf(-1, 2, b)
		}
	}
	tbl.Sort(0, len(tbl.Row)-1, 0)
	tbl.AddLineAfter(len(tbl.Row) - 1)                          // a line after the last row in the table
	tbl.InsertSumRow(len(tbl.Row), 0, len(tbl.Row)-1, []int{3}) // insert @ len essentially adds a row.  Only want to sum column 3
	return tbl
}

//PrintLedgerBalanceReport prints a report of data that will be used to format a ledger UI.
// This routine is primarily for testing
func PrintLedgerBalanceReport(ctx context.Context, ri *ReporterInfo) {
	fmt.Print(PrintLedgerBalanceReportString(ctx, ri))
}

//PrintLedgerBalanceReportString returns a string showing the balance of all ledgers as of ri.D2
func PrintLedgerBalanceReportString(ctx context.Context, ri *ReporterInfo) string {
	//s := fmt.Sprintf("LEDGER MARKERS\n%s\nBalances as of:  %s\n\n", ri.Xbiz.P.Name, ri.D2.Format("January 2, 2006"))
	tbl := LedgerBalanceReportTable(ctx, ri)
	tbl.TightenColumns()
	// return s + tbl.SprintTable()
	return ReportToString(&tbl, ri)
}
