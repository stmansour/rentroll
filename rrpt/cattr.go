package rrpt

import (
	"context"
	"gotable"
	"rentroll/rlib"
)

// RRreportCustomAttributesTable generates a table object for custom attributes
func RRreportCustomAttributesTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportCustomAttributesTable"
	var (
		err error
	)

	// table initialization
	tbl := getRRTable()

	tbl.AddColumn("CID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Value Type", 6, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Value", 15, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Units", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// prepare table's title, section1, section2
	err = TableReportHeaderBlock(ctx, &tbl, "Custom Attributes", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllCustomAttributes.Query()
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var a rlib.CustomAttribute
		err = rlib.ReadCustomAttributes(rows, &a)
		if err != nil {
			// set errors in section3 and return
			tbl.SetSection3(err.Error())
			return tbl
		}
		tbl.AddRow()
		tbl.Puts(-1, 0, a.IDtoString())
		tbl.Puts(-1, 1, rlib.IDtoString("B", a.BID))
		tbl.Puts(-1, 2, a.TypeToString())
		tbl.Puts(-1, 3, a.Name)
		tbl.Puts(-1, 4, a.Value)
		tbl.Puts(-1, 5, a.Units)
	}
	err = rows.Err()
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	tbl.TightenColumns()
	return tbl
}

// RRreportCustomAttributes generates a report of all rlib.GLAccount accounts
func RRreportCustomAttributes(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportCustomAttributesTable(ctx, ri)
	return ReportToString(&tbl, ri)
}

// RRreportCustomAttributeRefsTable generates a table object of custom attrib references
func RRreportCustomAttributeRefsTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportCustomAttributeRefsTable"
	var (
		err error
	)

	// table initialization
	tbl := getRRTable()

	tbl.AddColumn("Element Type", 4, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ID", 4, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("CID", 4, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)

	// prepare table's title, sections
	err = TableReportHeaderBlock(ctx, &tbl, "Custom Attributes References", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllCustomAttributeRefs.Query()
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var a rlib.CustomAttributeRef
		err = rlib.ReadCustomAttributeRefs(rows, &a)
		if err != nil {
			// set errors in section3 and return
			tbl.SetSection3(err.Error())
			return tbl
		}
		tbl.AddRow()
		tbl.Puti(-1, 0, a.ElementType)
		tbl.Puts(-1, 1, rlib.IDtoString("B", a.BID))
		tbl.Puti(-1, 2, a.ID)
		tbl.Puti(-1, 3, a.CID)
	}
	err = rows.Err()
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	tbl.TightenColumns()
	return tbl
}

// RRreportCustomAttributeRefs generates a report of all rlib.GLAccount accounts
func RRreportCustomAttributeRefs(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportCustomAttributeRefsTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
