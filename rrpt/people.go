package rrpt

import (
	"context"
	"gotable"
	"rentroll/rlib"
)

// RRreportPeopleTable generates a table object of all transactants relavant to BID in the database.
func RRreportPeopleTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportPeopleTable"
	var (
		err error
	)

	// table init
	tbl := getRRTable()

	tbl.AddColumn("TCID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("First Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Middle Name", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Last Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Company", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Is Company", 3, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Cell Phone", 17, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Primary Email", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set title, sections
	err = TableReportHeaderBlock(ctx, &tbl, "People", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllTransactantsForBID.Query(ri.Bid)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var p rlib.XPerson
		err = rlib.ReadTransactants(rows, &p.Trn)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			tbl.SetSection3(err.Error())
			return tbl
		}

		err = rlib.GetXPerson(ctx, p.Trn.TCID, &p)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			tbl.SetSection3(err.Error())
			return tbl
		}

		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())
		tbl.Puts(-1, 1, p.Trn.FirstName)
		tbl.Puts(-1, 2, p.Trn.MiddleName)
		tbl.Puts(-1, 3, p.Trn.LastName)
		tbl.Puts(-1, 4, p.Trn.CompanyName)
		tbl.Puts(-1, 5, rlib.BoolToYesNoString(p.Trn.IsCompany))
		tbl.Puts(-1, 6, p.Trn.CellPhone)
		tbl.Puts(-1, 7, p.Trn.PrimaryEmail)
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

// RRreportPeople generates a report of all Businesses defined in the database.
func RRreportPeople(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportPeopleTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
