package rrpt

import (
	"fmt"
	"rentroll/rcsv"
	"rentroll/rlib"
)

// LedgerBalanceReport builds a table of trial balance information
func LedgerBalanceReport(ri *rcsv.CSVReporterInfo) rlib.Table {
	bid := ri.Xbiz.P.BID
	var tbl rlib.Table
	tbl.Init()
	tbl.AddColumn("LID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("GLNumber", 8, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 35, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Summary Balance", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)
	tbl.AddColumn("Balance", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)

	for i := int64(0); i < int64(len(rlib.RRdb.BizTypes[bid].GLAccounts)); i++ {
		acct, ok := rlib.RRdb.BizTypes[bid].GLAccounts[i]
		if ok {
			tbl.AddRow()
			tbl.Puts(-1, 0, acct.IDtoString())
			tbl.Puts(-1, 1, acct.GLNumber)
			tbl.Puts(-1, 2, acct.Name)
			if rlib.RRdb.BizTypes[bid].GLAccounts[i].AllowPost != 0 {
				tbl.Putf(-1, 4, rlib.GetAccountBalance(bid, acct.LID, &ri.D2))
			} else {
				tbl.Putf(-1, 3, rlib.GetAccountBalance(bid, acct.LID, &ri.D2))
			}
		}
	}
	tbl.Sort(0, len(tbl.Row)-1, 1)
	tbl.AddLineAfter(len(tbl.Row) - 1)                          // a line after the last row in the table
	tbl.InsertSumRow(len(tbl.Row), 0, len(tbl.Row)-1, []int{4}) // insert @ len essentially adds a row.  Only want to sum column 4
	return tbl
}

//PrintLedgerBalanceReport prints a report of data that will be used to format a ledger UI.
// This routine is primarily for testing
func PrintLedgerBalanceReport(ri *rcsv.CSVReporterInfo) {
	fmt.Print(PrintLedgerBalanceReportString(ri))
}

//PrintLedgerBalanceReportString returns a string showing the balance of all ledgers as of ri.D2
func PrintLedgerBalanceReportString(ri *rcsv.CSVReporterInfo) string {
	s := fmt.Sprintf("LEDGER MARKERS\n%s\nBalances as of:  %s\n\n", ri.Xbiz.P.Name, ri.D2.Format("January 2, 2006"))
	tbl := LedgerBalanceReport(ri)
	tbl.TightenColumns()
	return s + tbl.SprintTable(rlib.TABLEOUTTEXT)
}
