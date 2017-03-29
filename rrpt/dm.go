package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRreportDepositMethodsTable generates a table for all rlib.DepositMethod
func RRreportDepositMethodsTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportDepositMethodsTable"

	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)

	tbl.AddColumn("DPMID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	err := TableReportHeaderBlock(&tbl, "Deposit Methods", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)

		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	m := rlib.GetAllDepositMethods(ri.Bid)
	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		tbl.Puts(-1, 0, rlib.IDtoString("DPM", m[i].DPMID))
		tbl.Puts(-1, 1, rlib.IDtoString("B", m[i].BID))
		tbl.Puts(-1, 2, m[i].Name)
	}
	tbl.TightenColumns()
	return tbl
}

// RRreportDepositMethods generates a report of all rlib.DepositMethod
func RRreportDepositMethods(ri *ReporterInfo) string {
	tbl := RRreportDepositMethodsTable(ri)
	return ReportToString(&tbl, ri)
}
