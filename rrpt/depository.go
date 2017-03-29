package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRreportDepositoryTable generates a table object for all rlib.Depository
func RRreportDepositoryTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportDepositoryTable"

	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)

	tbl.AddColumn("DEPID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("AccountNo", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	err := TableReportHeaderBlock(&tbl, "Depositories", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)

		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	m := rlib.GetAllDepositories(ri.Bid)
	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		tbl.Puts(-1, 0, rlib.IDtoString("DEP", m[i].DEPID))
		tbl.Puts(-1, 1, rlib.IDtoString("B", m[i].BID))
		tbl.Puts(-1, 2, m[i].AccountNo)
		tbl.Puts(-1, 3, m[i].Name)
	}
	tbl.TightenColumns()
	return tbl
}

// RRreportDepository generates a report of all rlib.Depository
func RRreportDepository(ri *ReporterInfo) string {
	tbl := RRreportDepositoryTable(ri)
	return ReportToString(&tbl, ri)
}
