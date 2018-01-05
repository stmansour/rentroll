package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// ReportCOA returns a string representation of the chart of accts
func ReportCOA(p rlib.GLAccount, tbl *gotable.Table, totalErrs *int) {

	Pldgr := ""

	if p.PLID > 0 {
		Pldgr = rlib.RRdb.BizTypes[p.BID].GLAccounts[p.PLID].Name
	}

	const (
		GLNo = iota
		Name = iota
		PGL  = iota
		AT   = iota
		Desc = iota
	)
	tbl.AddRow()
	tbl.Puts(-1, GLNo, p.GLNumber)
	tbl.Puts(-1, Name, p.Name)
	tbl.Puts(-1, PGL, Pldgr)
	tbl.Puts(-1, AT, p.AcctType)
	tbl.Puts(-1, Desc, p.Description)
}

// RRreportChartOfAccountsTable generates a table of all rlib.GLAccount accounts
func RRreportChartOfAccountsTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportChartOfAccountsTable"
	var (
		err       error
		totalErrs = 0
	)

	// init and prepare some values before table init
	ri.RptHeaderD1 = false
	ri.RptHeaderD2 = false

	// table initialization
	tbl := getRRTable()

	tbl.AddColumn("GLNumber", 8, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Name", 40, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Parent", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Account Type", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	rlib.InitBusinessFields(ri.Bid)
	rlib.RRdb.BizTypes[ri.Bid].GLAccounts, err = rlib.GetGLAccountMap(ctx, ri.Bid)
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	var a []int64                                          // Sort the map so test output will be the same every time. Sort by GLNumber.
	for k := range rlib.RRdb.BizTypes[ri.Bid].GLAccounts { // First make an array of all the LIDs
		a = append(a, k)
	}
	// now sort based on GLNumber, then by LID...
	m := rlib.RRdb.BizTypes[ri.Bid].GLAccounts // for notational convenience
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			isGreater := m[a[i]].GLNumber > m[a[j]].GLNumber
			isEqual := m[a[i]].GLNumber == m[a[j]].GLNumber
			if isGreater || (isEqual && m[a[i]].LID > m[a[j]].LID) {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	// prepare table's title, sections
	err = TableReportHeaderBlock(ctx, &tbl, "Chart of Accounts", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	for i := 0; i < len(a); i++ {
		ReportCOA(m[a[i]], &tbl, &totalErrs)
	}
	if totalErrs > 0 {
		errMsg := fmt.Sprintf("Encountered %d errors while creating this report. See log.", totalErrs)
		tbl.SetSection3(errMsg)
	}
	return tbl
}

// RRreportChartOfAccounts generates a report of all rlib.GLAccount accounts
func RRreportChartOfAccounts(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportChartOfAccountsTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
