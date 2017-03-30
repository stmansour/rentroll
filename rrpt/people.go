package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// RRreportPeopleTable generates a table object of all transactants relavant to BID in the database.
func RRreportPeopleTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportPeopleTable"

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
	err := TableReportHeaderBlock(&tbl, "People", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllTransactantsForBID.Query(ri.Bid)
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	for rows.Next() {
		var p rlib.XPerson
		rlib.ReadTransactants(rows, &p.Trn)
		rlib.GetXPerson(p.Trn.TCID, &p)
		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())
		tbl.Puts(-1, 1, p.Trn.FirstName)
		tbl.Puts(-1, 2, p.Trn.MiddleName)
		tbl.Puts(-1, 3, p.Trn.LastName)
		tbl.Puts(-1, 4, p.Trn.CompanyName)
		tbl.Puts(-1, 5, rlib.YesNoToString(int64(p.Trn.IsCompany)))
		tbl.Puts(-1, 6, p.Trn.CellPhone)
		tbl.Puts(-1, 7, p.Trn.PrimaryEmail)
	}
	rlib.Errcheck(rows.Err())
	tbl.TightenColumns()
	return tbl
}

// RRreportPeople generates a report of all Businesses defined in the database.
func RRreportPeople(ri *ReporterInfo) string {
	tbl := RRreportPeopleTable(ri)
	return ReportToString(&tbl, ri)
}
