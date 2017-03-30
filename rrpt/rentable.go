package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRreportRentablesTable generates a table of all rentables for BUD defined in the database.
func RRreportRentablesTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportRentablesTable"

	// table init
	tbl := getRRTable()

	tbl.AddColumn("RID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Assignment Time", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err := TableReportHeaderBlock(&tbl, "Rentables", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Bid)
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var p rlib.Rentable
		var s string
		rlib.ReadRentables(rows, &p)
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
	rlib.Errcheck(rows.Err())
	tbl.TightenColumns()
	return tbl
}

// RRreportRentables generates a report of all Businesses defined in the database.
func RRreportRentables(ri *ReporterInfo) string {
	tbl := RRreportRentablesTable(ri)
	return ReportToString(&tbl, ri)
}
