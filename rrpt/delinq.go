package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strings"
	"time"
)

// DelinquencyTextReport generates a text-based Delinqency report for the business in xbiz and timeframe d1 to d2.
func DelinquencyTextReport(ri *ReporterInfo) {
	tbl := DelinquencyReport(ri)
	fmt.Print(tbl)
}

// DelinquencyReportTable generates a table object for Delinqency report for the business in xbiz and timeframe d1 to d2.
func DelinquencyReportTable(ri *ReporterInfo) gotable.Table {
	funcname := "DelinquencyReportTable"

	// prepare and init some values
	ri.RptHeaderD1 = false
	ri.RptHeaderD2 = true

	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	totalErrs := 0

	// table init
	tbl := getRRTable()

	tbl.AddColumn("Rentable", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                              // column for the Rentable name
	tbl.AddColumn("Rentable Type", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                        // RentableType name
	tbl.AddColumn("Rentable Agreement", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                   // RentableType name
	tbl.AddColumn("Rentable Payors", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                      // Users of this rentable
	tbl.AddColumn("Rentable Users", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                       // Users of this rentable
	tbl.AddColumn("As of "+ri.D2.Format(rlib.RRDATEFMT3), 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT) // the Rental Agreement id
	tbl.AddColumn("30 Days Prior", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)                        // the possession start date
	tbl.AddColumn("60 Days Prior", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)                        // the possession start date
	tbl.AddColumn("90 Days Prior", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)                        // the rental start date
	tbl.AddColumn("Collection Notes", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                     // the possession start date

	// prepare table's title, sections
	err := TableReportHeaderBlock(&tbl, "Delinquency Report", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// loop through the Rentables...
	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Xbiz.P.BID)
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		// set errors in section3 and return
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	const (
		RentableName = 0
		RType        = iota
		RAgr         = iota
		RPayors      = iota
		RUsers       = iota
		D0           = iota
		D30          = iota
		D60          = iota
		D90          = iota
		CNotes       = iota
	)

	lid := rlib.RRdb.BizTypes[ri.Xbiz.P.BID].DefaultAccts[rlib.GLGENRCV].LID

	for rows.Next() {
		var r rlib.Rentable
		rlib.Errcheck(rlib.ReadRentables(rows, &r))
		rtid := rlib.GetRTIDForDate(r.RID, &ri.D2)
		//------------------------------------------------------------------------------
		// Get the RentalAgreement IDs for this rentable over the time range d1,ri.D2.
		// Note that this could result in multiple rental agreements.
		//------------------------------------------------------------------------------
		rra := rlib.GetAgreementsForRentable(r.RID, &d1, &ri.D2) // get all rental agreements for this period
		for i := 0; i < len(rra); i++ {                          //for each rental agreement id
			ra, err := rlib.GetRentalAgreement(rra[i].RAID) // load the agreement
			if err != nil {
				totalErrs++
				fmt.Printf("Error loading rental agreement %d: err = %s\n", rra[i].RAID, err.Error())
				continue
			}
			na := r.GetUserNameList(&ra.PossessionStart, &ra.PossessionStop) // get the list of user names for this time period
			usernames := strings.Join(na, ",")                               // concatenate with a comma separator
			pa := ra.GetPayorNameList(&ra.RentStart, &ra.RentStart)          // get the payors for this time period
			payornames := strings.Join(pa, ", ")                             // concatenate with comma
			d30 := ri.D2.AddDate(0, 0, -30)
			d60 := ri.D2.AddDate(0, 0, -60)
			d90 := ri.D2.AddDate(0, 0, -90)
			d2Bal := rlib.GetRentableAccountBalance(ri.Xbiz.P.BID, lid, r.RID, &ri.D2)
			d30Bal := rlib.GetRentableAccountBalance(ri.Xbiz.P.BID, lid, r.RID, &d30)
			d60Bal := rlib.GetRentableAccountBalance(ri.Xbiz.P.BID, lid, r.RID, &d60)
			d90Bal := rlib.GetRentableAccountBalance(ri.Xbiz.P.BID, lid, r.RID, &d90)

			tbl.AddRow()
			tbl.Puts(-1, RentableName, r.RentableName)
			tbl.Puts(-1, RType, ri.Xbiz.RT[rtid].Style)
			tbl.Puts(-1, RAgr, ra.IDtoString())
			tbl.Puts(-1, RPayors, payornames)
			tbl.Puts(-1, RUsers, usernames)
			tbl.Putf(-1, D0, d2Bal)
			tbl.Putf(-1, D30, d30Bal)
			tbl.Putf(-1, D60, d60Bal)
			tbl.Putf(-1, D90, d90Bal)
		}
	}
	rlib.Errcheck(rows.Err())
	if totalErrs > 0 {
		errMsg := fmt.Sprintf("Encountered %d errors while creating this report. See log.", totalErrs)
		tbl.SetSection3(errMsg)
	}
	return tbl
}

// DelinquencyReport generates a text-based Delinqency report for the business in xbiz and timeframe d1 to d2.
func DelinquencyReport(ri *ReporterInfo) string {
	tbl := DelinquencyReportTable(ri)
	return ReportToString(&tbl, ri)
}
