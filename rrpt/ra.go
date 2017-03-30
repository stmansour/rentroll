package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strings"
	"time"
)

// RRreportRentalAgreementsTable generates a table object for All rental agreements related with Business
func RRreportRentalAgreementsTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportRentalAgreementsTable"

	// init and prepare some values before table init
	totalErrs := 0

	// table init
	tbl := getRRTable()

	tbl.AddColumn("RAID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Payor", 60, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("User", 60, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Agreement Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Agreement Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Possession Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Possession Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Cycle Epoch", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Renewal", 2, gotable.CELLINT, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Unspecified Adults", 6, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Unspecified Children", 6, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Special Provisions", 40, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Notes", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err := TableReportHeaderBlock(&tbl, "Rental Agreement", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreements.Query(ri.Bid)
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	var raid int64
	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC)

	for rows.Next() {
		var p rlib.RentalAgreement
		rlib.Errcheck(rows.Scan(&raid))

		p, err = rlib.GetXRentalAgreement(raid, &d1, &d2)
		if err != nil {
			totalErrs++
			rlib.Ulog("RRreportRentalAgreements: rlib.GetXRentalAgreement returned err = %v\n", err)
			continue
		}

		note := ""
		if p.NLID > 0 {
			nl := rlib.GetNoteList(p.NLID)
			if len(nl.N) > 0 {
				note = nl.N[0].Comment
			}
		}
		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())
		tbl.Puts(-1, 1, strings.Join(p.GetPayorNameList(&p.AgreementStart, &p.AgreementStop), ", "))
		tbl.Puts(-1, 2, strings.Join(p.GetUserNameList(&p.AgreementStart, &p.AgreementStop), ", "))
		tbl.Putd(-1, 3, p.AgreementStart)
		tbl.Putd(-1, 4, p.AgreementStop)
		tbl.Putd(-1, 5, p.PossessionStart)
		tbl.Putd(-1, 6, p.PossessionStop)
		tbl.Putd(-1, 7, p.RentStart)
		tbl.Putd(-1, 8, p.RentStop)
		tbl.Putd(-1, 9, p.RentCycleEpoch)
		tbl.Puti(-1, 10, p.Renewal)
		tbl.Puti(-1, 11, p.UnspecifiedAdults)
		tbl.Puti(-1, 12, p.UnspecifiedChildren)
		tbl.Puts(-1, 13, p.SpecialProvisions)
		tbl.Puts(-1, 14, note)
	}
	rlib.Errcheck(rows.Err())
	tbl.TightenColumns()
	if totalErrs > 0 {
		errMsg := fmt.Sprintf("Encountered %d errors while creating this report. See log.", totalErrs)
		tbl.SetSection3(errMsg)
	}
	return tbl
}

// RRreportRentalAgreements generates a report of all Businesses defined in the database.
func RRreportRentalAgreements(ri *ReporterInfo) string {
	tbl := RRreportRentalAgreementsTable(ri)
	return ReportToString(&tbl, ri)
}
