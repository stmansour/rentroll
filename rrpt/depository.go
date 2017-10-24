package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// RRreportDepositoryTable generates a table object for all rlib.Depository
func RRreportDepositoryTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportDepositoryTable"

	// table init
	tbl := getRRTable()

	tbl.AddColumn("DEPID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("AccountNo", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("GLAccount", 45, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// prepare table title, sections
	err := TableReportHeaderBlock(&tbl, "Depositories", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// get records from db
	m := rlib.GetAllDepositories(ri.Bid)
	for i := 0; i < len(m); i++ {
		var l rlib.GLAccount
		if m[i].LID > 0 {
			l = rlib.GetLedger(m[i].LID)
		}
		tbl.AddRow()
		tbl.Puts(-1, 0, rlib.IDtoString("DEP", m[i].DEPID))
		tbl.Puts(-1, 1, rlib.IDtoString("B", m[i].BID))
		tbl.Puts(-1, 2, m[i].AccountNo)
		tbl.Puts(-1, 3, m[i].Name)
		if l.LID > 0 {
			tbl.Puts(-1, 4, fmt.Sprintf("%s (%s)", l.GLNumber, l.Name))
		}
	}
	tbl.TightenColumns()
	return tbl
}

// RRreportDepository generates a report of all rlib.Depository
func RRreportDepository(ri *ReporterInfo) string {
	tbl := RRreportDepositoryTable(ri)
	return ReportToString(&tbl, ri)
}
