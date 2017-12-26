package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strings"
	"time"
)

// DelinquencyTextReport generates a text-based Delinqency report for the business in xbiz and timeframe d1 to d2.
func DelinquencyTextReport(ctx context.Context, ri *ReporterInfo) {
	tbl := DelinquencyReport(ctx, ri)
	fmt.Print(tbl)
}

// DelinquencyReportTable generates a table object for Delinqency report for the business in xbiz and timeframe d1 to d2.
func DelinquencyReportTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "DelinquencyReportTable"

	var (
		err       error
		totalErrs = 0
		d1        = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	)

	// prepare and init some values
	ri.RptHeaderD1 = false
	ri.RptHeaderD2 = true

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
	err = TableReportHeaderBlock(ctx, &tbl, "Delinquency Report", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// loop through the Rentables...
	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Xbiz.P.BID)
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer rows.Close()

	// lid := rlib.RRdb.BizTypes[ri.Xbiz.P.BID].DefaultAccts[rlib.GLGENRCV].LID  // this was the old way to do it, when we had default accounts
	// TODO(Steve): Shouldn't we have restriction over this method?
	// Shouldn't we pass context here?
	m := rlib.GetReceivableAccounts(ri.Xbiz.P.BID)
	if len(m) < 1 {
		return tbl
	}
	lid := m[0] // for now, just take the first one, we'll need to pass in which one as we start to support multiple

	for rows.Next() {
		var r rlib.Rentable
		err = rlib.ReadRentables(rows, &r)
		if err != nil {
			totalErrs++
			rlib.Console("Error reading rentable data: err = %s\n", err.Error())
			continue
		}

		rtid, err := rlib.GetRTIDForDate(ctx, r.RID, &ri.D2)
		if err != nil {
			totalErrs++
			rlib.Console("Error loading rentable type for Rentable %d: err = %s\n", r.RID, err.Error())
			continue
		}
		//------------------------------------------------------------------------------
		// Get the RentalAgreement IDs for this rentable over the time range d1,ri.D2.
		// Note that this could result in multiple rental agreements.
		//------------------------------------------------------------------------------
		rra, err := rlib.GetAgreementsForRentable(ctx, r.RID, &d1, &ri.D2) // get all rental agreements for this period
		if err != nil {
			totalErrs++
			rlib.Console("Error loading rental agreement for Rentable %d: err = %s\n", r.RID, err.Error())
			continue
		}
		for i := 0; i < len(rra); i++ { //for each rental agreement id
			ra, err := rlib.GetRentalAgreement(ctx, rra[i].RAID) // load the agreement
			if err != nil {
				totalErrs++
				rlib.Console("Error loading rental agreement %d: err = %s\n", rra[i].RAID, err.Error())
				continue
			}
			na, err := r.GetUserNameList(ctx, &ra.PossessionStart, &ra.PossessionStop) // get the list of user names for this time period
			if err != nil {
				totalErrs++
				rlib.Console("Error while getting user list for rental agreement %d: err = %s\n", rra[i].RAID, err.Error())
				continue
			}

			usernames := strings.Join(na, ",")                                // concatenate with a comma separator
			pa, err := ra.GetPayorNameList(ctx, &ra.RentStart, &ra.RentStart) // get the payors for this time period
			if err != nil {
				totalErrs++
				rlib.Console("Error while getting payor list for rental agreement %d: err = %s\n", rra[i].RAID, err.Error())
				continue
			}

			payornames := strings.Join(pa, ", ") // concatenate with comma
			d30 := ri.D2.AddDate(0, 0, -30)
			d60 := ri.D2.AddDate(0, 0, -60)
			d90 := ri.D2.AddDate(0, 0, -90)
			d2Bal, err := rlib.GetRentableAccountBalance(ctx, ri.Xbiz.P.BID, lid, r.RID, &ri.D2)
			if err != nil {
				totalErrs++
				rlib.Console("Error while d2Bal for Rentable %d: err = %s\n", r.RID, err.Error())
				continue
			}

			d30Bal, err := rlib.GetRentableAccountBalance(ctx, ri.Xbiz.P.BID, lid, r.RID, &d30)
			if err != nil {
				totalErrs++
				rlib.Console("Error while d30Bal for Rentable %d: err = %s\n", r.RID, err.Error())
				continue
			}

			d60Bal, err := rlib.GetRentableAccountBalance(ctx, ri.Xbiz.P.BID, lid, r.RID, &d60)
			if err != nil {
				totalErrs++
				rlib.Console("Error while d60Bal for Rentable %d: err = %s\n", r.RID, err.Error())
				continue
			}

			d90Bal, err := rlib.GetRentableAccountBalance(ctx, ri.Xbiz.P.BID, lid, r.RID, &d90)
			if err != nil {
				totalErrs++
				rlib.Console("Error while d90Bal for Rentable %d: err = %s\n", r.RID, err.Error())
				continue
			}

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

	err = rows.Err()
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}

	if totalErrs > 0 {
		errMsg := fmt.Sprintf("Encountered %d errors while creating this report. See log.", totalErrs)
		tbl.SetSection3(errMsg)
		return tbl
	}

	return tbl
}

// DelinquencyReport generates a text-based Delinqency report for the business in xbiz and timeframe d1 to d2.
func DelinquencyReport(ctx context.Context, ri *ReporterInfo) string {
	tbl := DelinquencyReportTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
