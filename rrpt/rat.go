package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// RRreportRentalAgreementTemplatesTable generates a table object for all rental agreement templates
func RRreportRentalAgreementTemplatesTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportRentalAgreementTemplatesTable"

	// table init
	tbl := getRRTable()

	tbl.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RA Template ID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RA Template Name", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err := TableReportHeaderBlock(ctx, &tbl, "Rental Agreement Templates", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementTemplates.Query()
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var p rlib.RentalAgreementTemplate
		err = rlib.ReadRentalAgreementTemplates(rows, &p)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			tbl.SetSection3(err.Error())
			return tbl
		}
		tbl.AddRow()
		tbl.Puts(-1, 0, fmt.Sprintf("B%08d", p.BID))
		tbl.Puts(-1, 1, p.IDtoString())
		tbl.Puts(-1, 2, p.RATemplateName)
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

// RRreportRentalAgreementTemplates generates a report for all rental agreement templates
func RRreportRentalAgreementTemplates(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportRentalAgreementTemplatesTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
