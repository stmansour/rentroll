package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRAssessmentsTable generates a gotable table object
// for report of all rlib.Assessment records related with business
func RRAssessmentsTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRAssessmentsTable"

	// initialize and prepare some ReporterInfo values
	bid := ri.Bid
	d1 := ri.D1
	d2 := ri.D2

	rlib.InitBusinessFields(bid)
	rlib.RRdb.BizTypes[bid].GLAccounts = rlib.GetGLAccountMap(bid)

	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true
	ri.BlankLineAfterRptName = true

	// initialize table for assessments report
	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)

	tbl.AddColumn("ASMID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("PASMID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RAID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Cycle", 13, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Proration Cycle", 13, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Amount", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("AsmType", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Account Rule", 80, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// prepare table's title, section1, section2
	err := TableReportHeaderBlock(&tbl, "Assessments", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)

		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(bid, d2, d1)
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	// fit records in table row one by one
	for rows.Next() {
		var a rlib.Assessment
		rlib.ReadAssessments(rows, &a)
		tbl.AddRow()
		tbl.Puts(-1, 0, a.IDtoString())
		tbl.Puts(-1, 1, rlib.IDtoString("ASM", a.PASMID))
		tbl.Puts(-1, 2, rlib.IDtoString("RA", a.RAID))
		tbl.Puts(-1, 3, rlib.IDtoString("R", a.RID))
		tbl.Puts(-1, 4, rlib.RentalPeriodToString(a.RentCycle))
		tbl.Puts(-1, 5, rlib.RentalPeriodToString(a.ProrationCycle))
		tbl.Putf(-1, 6, a.Amount)
		tbl.Puts(-1, 7, rlib.RRdb.BizTypes[a.BID].GLAccounts[a.ATypeLID].Name)
		tbl.Puts(-1, 8, a.AcctRule)
	}
	rlib.Errcheck(rows.Err())
	tbl.TightenColumns()
	return tbl
}

// RRreportAssessments generates a text report of all rlib.Assessments records
func RRreportAssessments(ri *ReporterInfo) string {
	tbl := RRAssessmentsTable(ri)
	return ReportToString(&tbl, ri)
}
