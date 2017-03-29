package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRreportBusinessTable generates a Table of all Businesses defined in the database.
func RRreportBusinessTable(ri *ReporterInfo) gotable.Table {
	// initialize table
	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)

	tbl.SetTitle("Business Units")
	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BUD", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Default Rent Cycle", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Default Proration Cycle", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Default GSRPC Cycle", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	rows, err := rlib.RRdb.Prepstmt.GetAllBusinesses.Query()
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var p rlib.Business
		rlib.ReadBusinesses(rows, &p)
		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())
		tbl.Puts(-1, 1, p.Designation)
		tbl.Puts(-1, 2, p.Name)
		tbl.Puts(-1, 3, rlib.RentalPeriodToString(p.DefaultRentCycle))
		tbl.Puts(-1, 4, rlib.RentalPeriodToString(p.DefaultProrationCycle))
		tbl.Puts(-1, 5, rlib.RentalPeriodToString(p.DefaultGSRPC))
	}
	rlib.Errcheck(rows.Err())
	return tbl
}

// RRreportBusiness generates a String Report of all Businesses defined in the database.
func RRreportBusiness(ri *ReporterInfo) string {
	tbl := RRreportBusinessTable(ri)
	return ReportToString(&tbl, ri)
}
