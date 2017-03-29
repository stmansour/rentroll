package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// RRreportRentalAgreementTemplatesTable generates a table object for all rental agreement templates
func RRreportRentalAgreementTemplatesTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportRentalAgreementTemplatesTable"

	// table init
	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)

	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RA Template ID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RA Template Name", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	err := TableReportHeaderBlock(&tbl, "Rental Agreement Templates", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)

		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementTemplates.Query()
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var p rlib.RentalAgreementTemplate
		rlib.ReadRentalAgreementTemplates(rows, &p)
		tbl.AddRow()
		tbl.Puts(-1, 0, fmt.Sprintf("B%08d", p.BID))
		tbl.Puts(-1, 1, p.IDtoString())
		tbl.Puts(-1, 2, p.RATemplateName)
	}
	rlib.Errcheck(rows.Err())
	tbl.TightenColumns()
	return tbl
}

// RRreportRentalAgreementTemplates generates a report for all rental agreement templates
func RRreportRentalAgreementTemplates(ri *ReporterInfo) string {
	tbl := RRreportRentalAgreementTemplatesTable(ri)
	return ReportToString(&tbl, ri)
}
