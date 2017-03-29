package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// GSRTextReport generates a list of GSR values for all rentables on the specified date
func GSRTextReport(ri *ReporterInfo) {
	tbl := GSRReportTable(ri)
	fmt.Print(tbl)
}

// GSRReportTable generates a list of GSR values for all rentables on the specified date
func GSRReportTable(ri *ReporterInfo) gotable.Table {
	funcname := "GSRReportTable"

	// init and prepare some values before table init
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = false

	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)

	tbl.AddColumn("Rentable", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)        // column for the Rentable name
	tbl.AddColumn("Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)           // Rentable name
	tbl.AddColumn("Rentable Type", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)  // Rentable Type
	tbl.AddColumn("Rentable Style", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT) // Rentable Style
	tbl.AddColumn("GSR", 8, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)             // 4  GSR
	tbl.AddColumn("Rent Cycle", 13, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)     // 5  Rent Cycle
	tbl.AddColumn("Prorate Cycle", 13, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)  // 6  Proration Cycle

	err := TableReportHeaderBlock(&tbl, "Gross Scheduled Rent", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)

		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Xbiz.P.BID)
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var r rlib.Rentable
		var rc, pc int64
		rlib.ReadRentables(rows, &r)                         // get the next rentable from the database
		rtr := rlib.GetRentableTypeRefForDate(r.RID, &ri.D1) // what type is it on this date?
		rc = ri.Xbiz.RT[rtr.RTID].RentCycle
		pc = ri.Xbiz.RT[rtr.RTID].Proration
		if rtr.OverrideRentCycle != 0 {
			rc = rtr.OverrideRentCycle
		}
		if rtr.OverrideProrationCycle != 0 {
			pc = rtr.OverrideProrationCycle
		}
		dt1 := ri.D1.Add(rlib.CycleDuration(rc, ri.D1))                      // 1 full cycle
		amt, _, _, err := rlib.CalculateLoadedGSR(&r, &ri.D1, &dt1, ri.Xbiz) // calculate its GSR
		if err != nil {
			fmt.Printf("%s: Rentable %d, error calculating GSR: %s\n", funcname, r.RID, err.Error())
		}
		tbl.AddRow()
		tbl.Puts(-1, 0, r.IDtoString())
		tbl.Puts(-1, 1, r.RentableName)
		tbl.Puts(-1, 2, ri.Xbiz.RT[rtr.RTID].Name)
		tbl.Puts(-1, 3, ri.Xbiz.RT[rtr.RTID].Style)
		tbl.Putf(-1, 4, amt)
		tbl.Puts(-1, 5, rlib.RentalPeriodToString(rc))
		tbl.Puts(-1, 6, rlib.RentalPeriodToString(pc))
	}
	rlib.Errcheck(rows.Err())
	return tbl
}

// GSRReport generates a text-based report based on GSRReportTable table object
func GSRReport(ri *ReporterInfo) string {
	tbl := GSRReportTable(ri)
	return ReportToString(&tbl, ri)
}
