package rrpt

import (
	"context"
	"gotable"
	"rentroll/rlib"
)

// RRreportBusinessTable generates a Table of all Businesses defined in the database.
func RRreportBusinessTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	// initialize table
	tbl := getRRTable()

	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BUD", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Default Rent Cycle", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Default Proration Cycle", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Default GSRPC Cycle", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	tbl.SetTitle("Business Units\n\n")

	rows, err := rlib.RRdb.Prepstmt.GetAllBusinesses.Query()
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var p rlib.Business
		err = rlib.ReadBusinesses(rows, &p)
		if err != nil {
			// set errors in section3 and return
			tbl.SetSection3(NoRecordsFoundMsg)
			return tbl
		}
		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())
		tbl.Puts(-1, 1, p.Designation)
		tbl.Puts(-1, 2, p.Name)
		tbl.Puts(-1, 3, rlib.RentalPeriodToString(p.DefaultRentCycle))
		tbl.Puts(-1, 4, rlib.RentalPeriodToString(p.DefaultProrationCycle))
		tbl.Puts(-1, 5, rlib.RentalPeriodToString(p.DefaultGSRPC))
	}
	err = rows.Err()
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	return tbl
}

// RRreportBusiness generates a String Report of all Businesses defined in the database.
func RRreportBusiness(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportBusinessTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
