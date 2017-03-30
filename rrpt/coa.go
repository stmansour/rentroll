package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// ReportCOA returns a string representation of the chart of accts
func ReportCOA(p rlib.GLAccount, tbl *gotable.Table, totalErrs *int) {

	Pldgr := ""
	lm := rlib.GetLatestLedgerMarkerByLID(p.BID, p.LID)
	if lm.LMID == 0 {
		*totalErrs++
		fmt.Printf("ReportChartOfAcctsToText: error getting latest LedgerMarker for L%08d\n", p.LID)
		return
	}

	s := ""
	if rlib.GLCASH <= p.Type && p.Type <= rlib.GLLAST {
		s = fmt.Sprintf("%d", p.Type)
	}

	var sp string
	switch p.RAAssociated {
	case 0:
		sp = "unknown"
	case 1:
		sp = "Unassociated"
	case 2:
		sp = "Associated"
	default:
		sp = fmt.Sprintf("??? invalid: %d", p.RAAssociated)
	}
	if p.PLID > 0 {
		Pldgr = rlib.RRdb.BizTypes[p.BID].GLAccounts[p.PLID].Name
	}
	const (
		BID  = 0
		LID  = iota
		PLID = iota
		LMID = iota
		Type = iota
		GLNo = iota
		Name = iota
		PGL  = iota
		QBAT = iota
		Bal  = iota
		RAA  = iota
		RAR  = iota
		Desc = iota
	)

	tbl.AddRow()
	tbl.Puts(-1, BID, fmt.Sprintf("B%08d", p.BID))
	tbl.Puts(-1, LID, p.IDtoString())
	if p.PLID > 0 {
		tbl.Puts(-1, PLID, fmt.Sprintf("L%08d", p.PLID))
	}
	tbl.Puts(-1, LMID, lm.IDtoString())
	tbl.Puts(-1, Type, s)
	tbl.Puts(-1, GLNo, p.GLNumber)
	tbl.Puts(-1, Name, p.Name)
	tbl.Puts(-1, PGL, Pldgr)
	tbl.Puts(-1, QBAT, p.AcctType)
	tbl.Putf(-1, Bal, lm.Balance)
	tbl.Puts(-1, RAA, sp)
	tbl.Puti(-1, RAR, p.RARequired)
	tbl.Puts(-1, Desc, p.Description)
}

// RRreportChartOfAccountsTable generates a table of all rlib.GLAccount accounts
func RRreportChartOfAccountsTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportChartOfAccountsTable"

	// init and prepare some values before table init
	totalErrs := 0
	ri.RptHeaderD1 = false
	ri.RptHeaderD2 = false

	rlib.InitBusinessFields(ri.Bid)
	rlib.RRdb.BizTypes[ri.Bid].GLAccounts = rlib.GetGLAccountMap(ri.Bid)

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

	// table initialization
	tbl := getRRTable()

	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("LID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PLID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("LMID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Type", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("GLNumber", 8, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Name", 40, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Parent", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Quick Books Account Type", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Rental Agreement Associated", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rental Agreement Required", 5, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Description", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// prepare table's title, sections
	err := TableReportHeaderBlock(&tbl, "Chart of Accounts", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
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
func RRreportChartOfAccounts(ri *ReporterInfo) string {
	tbl := RRreportChartOfAccountsTable(ri)
	return ReportToString(&tbl, ri)
}
