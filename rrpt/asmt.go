package rrpt

import (
	"context"
	"gotable"
	"rentroll/rlib"
)

// RRAssessmentsTable generates a gotable table object
// for report of all rlib.Assessment records related with business
func RRAssessmentsTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRAssessmentsTable"
	var (
		err error
		bid = ri.Bid
		d1  = ri.D1
		d2  = ri.D2
	)

	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true
	ri.BlankLineAfterRptName = true

	// initialize table for assessments report
	tbl := getRRTable()

	tbl.AddColumn("ASMID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RAID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rentable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Cycle", 13, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Proration Cycle", 13, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Amount", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("AsmType", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("AR Name", 80, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// prepare table's title, section1, section2, section3 if there are any error
	err = TableReportHeaderBlock(ctx, &tbl, "Assessments", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	rlib.InitBusinessFields(bid)
	rlib.RRdb.BizTypes[bid].GLAccounts, err = rlib.GetGLAccountMap(ctx, bid)
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(bid, d2, d1)
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer rows.Close()

	// fit records in table row one by one
	for rows.Next() {
		var a rlib.Assessment
		err = rlib.ReadAssessments(rows, &a)
		if err != nil {
			// set errors in section3 and return
			tbl.SetSection3(err.Error())
			return tbl
		}

		r, err := rlib.GetRentable(ctx, a.RID)
		if err != nil {
			// set errors in section3 and return
			tbl.SetSection3(err.Error())
			return tbl
		}

		tbl.AddRow()
		tbl.Puts(-1, 0, a.IDtoString())
		tbl.Puts(-1, 1, rlib.IDtoString("RA", a.RAID))
		tbl.Puts(-1, 2, r.RentableName)
		tbl.Puts(-1, 3, rlib.RentalPeriodToString(a.RentCycle))
		tbl.Puts(-1, 4, rlib.RentalPeriodToString(a.ProrationCycle))
		tbl.Putf(-1, 5, a.Amount)
		tbl.Puts(-1, 6, rlib.RRdb.BizTypes[a.BID].GLAccounts[a.ATypeLID].Name)
		ar, err := rlib.GetAssessmentAccountRuleText(ctx, &a)
		if err != nil {
			// set errors in section3 and return
			tbl.SetSection3(err.Error())
			return tbl
		}
		tbl.Puts(-1, 7, ar)
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

// RRreportAssessments generates a text report of all rlib.Assessments records
func RRreportAssessments(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRAssessmentsTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
