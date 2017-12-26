package rrpt

import (
	"context"
	"gotable"
	"rentroll/rlib"
)

// RRreportRentablesTable generates a table of all rentables for BUD defined in the database.
func RRreportRentablesTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportRentablesTable"

	// table init
	tbl := getRRTable()

	tbl.AddColumn("RID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Assignment Time", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err := TableReportHeaderBlock(ctx, &tbl, "Rentables", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Bid)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var p rlib.Rentable
		var s string

		err = rlib.ReadRentables(rows, &p)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			tbl.SetSection3(err.Error())
			return tbl
		}

		switch p.AssignmentTime {
		case 0:
			s = "unknown"
		case 1:
			s = "pre-assign"
		case 2:
			s = "no pre-assign"
		}
		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())
		tbl.Puts(-1, 1, p.RentableName)
		tbl.Puts(-1, 2, s)
	}
	err = rows.Err()
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	tbl.TightenColumns()
	return tbl
}

// RRreportRentables generates a report of all Businesses defined in the database.
func RRreportRentables(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportRentablesTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
