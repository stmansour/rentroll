package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRreportCustomAttributesTable generates a table object for custom attributes
func RRreportCustomAttributesTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportCustomAttributesTable"

	// table initialization
	tbl := getRRTable()

	tbl.AddColumn("CID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Value Type", 6, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Value", 15, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Units", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// prepare table's title, section1, section2
	err := TableReportHeaderBlock(&tbl, "Custom Attributes", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllCustomAttributes.Query()
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var a rlib.CustomAttribute
		rlib.ReadCustomAttributes(rows, &a)
		tbl.AddRow()
		tbl.Puts(-1, 0, a.IDtoString())
		tbl.Puts(-1, 1, rlib.IDtoString("B", a.BID))
		tbl.Puts(-1, 2, a.TypeToString())
		tbl.Puts(-1, 3, a.Name)
		tbl.Puts(-1, 4, a.Value)
		tbl.Puts(-1, 5, a.Units)
	}
	rlib.Errcheck(rows.Err())
	tbl.TightenColumns()
	return tbl
}

// RRreportCustomAttributes generates a report of all rlib.GLAccount accounts
func RRreportCustomAttributes(ri *ReporterInfo) string {
	tbl := RRreportCustomAttributesTable(ri)
	return ReportToString(&tbl, ri)
}

// RRreportCustomAttributeRefsTable generates a table object of custom attrib references
func RRreportCustomAttributeRefsTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportCustomAttributeRefsTable"

	// table initialization
	tbl := getRRTable()

	tbl.AddColumn("Element Type", 4, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ID", 4, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("CID", 4, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)

	// prepare table's title, sections
	err := TableReportHeaderBlock(&tbl, "Custom Attributes References", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllCustomAttributeRefs.Query()
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var a rlib.CustomAttributeRef
		rlib.ReadCustomAttributeRefs(rows, &a)
		tbl.AddRow()
		tbl.Puti(-1, 0, a.ElementType)
		tbl.Puts(-1, 1, rlib.IDtoString("B", a.BID))
		tbl.Puti(-1, 2, a.ID)
		tbl.Puti(-1, 3, a.CID)
	}
	rlib.Errcheck(rows.Err())
	tbl.TightenColumns()
	return tbl
}

// RRreportCustomAttributeRefs generates a report of all rlib.GLAccount accounts
func RRreportCustomAttributeRefs(ri *ReporterInfo) string {
	tbl := RRreportCustomAttributeRefsTable(ri)
	return ReportToString(&tbl, ri)
}
