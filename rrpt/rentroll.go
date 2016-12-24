package rrpt

import (
	"fmt"
	"rentroll/rcsv"
	"rentroll/rlib"
	"strings"
	"time"
)

// OtherIncomeGLAccountName and the rest will need to become configurable parameters for this report!!
const (
	OtherIncomeGLAccountName  = string("Other Income")
	IncomeOffsetGLAccountName = string("Income Offsets")
)

// ComputeGSRandGSRRate returns the GSR and GSR rate for Rentable over time period dtStart - dtStop
func ComputeGSRandGSRRate(p *rlib.Rentable, dtStart, dtStop *time.Time, xbiz *rlib.XBusiness) (float64, float64) {
	// Compute the GSR for this period.
	x, _, _, _ := rlib.CalculateLoadedGSR(p, dtStart, dtStop, xbiz)

	// Compute the GSR Rate
	gsrRate := float64(0)                                           //initialize
	prt := rlib.SelectRentableTypeRefForDate(&p.RT, dtStart)        // The RentableType at the start of the period
	n2 := rlib.CycleDuration(xbiz.RT[prt.RTID].RentCycle, *dtStart) // rent cycle duration
	n1 := dtStop.Sub(*dtStart)                                      // duration of this particular period
	if n1 < n2 {                                                    // if < 1 rent cycle, we'll need to extrapolate
		gsrRate = float64(n2) / float64(n1) * x //  (x: GSR this period)/(n1: this period) = (y: extrapolated GSR)/(n2: rent cycle)
	} else {
		dt := dtStart.Add(n2)
		gsrRate, _, _, _ = rlib.CalculateLoadedGSR(p, dtStart, &dt, xbiz)
	}
	return x, gsrRate
}

// RentRollTextReport prints a text-based RentRoll report for the business in xbiz and timeframe d1 to d2 to stdout
func RentRollTextReport(ri *rcsv.CSVReporterInfo) {
	fmt.Print(RentRollReportString(ri))
}

// RentRollReportString returns a string containin a text-based RentRoll report for the business in xbiz and timeframe d1 to d2.
func RentRollReportString(ri *rcsv.CSVReporterInfo) string {
	tbl, err := RentRollReport(ri)
	if err == nil {
		return tbl.Title + tbl.SprintRowText(len(tbl.Row)-1) + tbl.SprintLineText() + tbl.SprintTable(rlib.TABLEOUTTEXT)
	}
	return err.Error()
}

// RentRollReport generates a text-based RentRoll report for the business in ri.Xbiz and timeframe d1 to d2.
func RentRollReport(ri *rcsv.CSVReporterInfo) (rlib.Table, error) {
	funcname := "RentRollReport"
	var d1, d2 *time.Time
	var tbl rlib.Table
	d1 = &ri.D1
	d2 = &ri.D2

	custom := "Square Feet"
	var noerr error
	bu, err := rlib.GetBusinessUnitByDesignation(ri.Xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s\n", funcname, err.Error())
		return tbl, e
	}
	c, err := rlib.GetCompany(int64(bu.CoCode))
	if err != nil {
		e := fmt.Errorf("%s: GetCompany(%d) returned error getting Company: %s\n", funcname, bu.CoCode, err.Error())
		return tbl, e
	}

	tbl.Init() //sets column spacing and date format to default

	s := fmt.Sprintf("%s\n", strings.ToUpper(c.LegalName))
	s += fmt.Sprintf("Rentroll report for period beginning %s up to and including %s\n\n", d1.Format(rlib.RRDATEFMT3), d2.AddDate(0, 0, -1).Format(rlib.RRDATEFMT4))
	tbl.SetTitle(s)

	totalsRSet := tbl.CreateRowset()                                                      // a rowset to sum for totals
	tbl.AddColumn("Rentable", 20, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)                   // column for the Rentable name
	tbl.AddColumn("Rentable Type", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)              // RentableType name
	tbl.AddColumn(custom, 5, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)                          // the Custom Attribute "Square Feet"
	tbl.AddColumn("Rentable Users", 30, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)             // Users of this rentable
	tbl.AddColumn("Rentable Payors", 30, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)            // Users of this rentable
	tbl.AddColumn("Rental Agreement", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)           // the Rental Agreement id
	tbl.AddColumn("Use Start", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)                    // the possession start date
	tbl.AddColumn("Use Stop", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)                     // the possession start date
	tbl.AddColumn("Rental Start", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)                 // the rental start date
	tbl.AddColumn("Rental Stop", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)                  // the rental start date
	tbl.AddColumn("Rental Agreement Start", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)       // the possession start date
	tbl.AddColumn("Rental Agreement Stop", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)        // the possession start date
	tbl.AddColumn("Rent Cycle", 12, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)                 // the rent cycle
	tbl.AddColumn("GSR Rate", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)                   // gross scheduled rent
	tbl.AddColumn("GSR This Period", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)            // gross scheduled rent
	tbl.AddColumn(IncomeOffsetGLAccountName, 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)    // GL Account
	tbl.AddColumn("Contract Rent", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)              // contract rent amounts
	tbl.AddColumn(OtherIncomeGLAccountName, 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)     // GL Account
	tbl.AddColumn("Payments Received", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)          // contract rent amounts
	tbl.AddColumn("Beginning Receivable", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Change In Receivable", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Ending Receivable", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)          // account for the associated RentalAgreement
	tbl.AddColumn("Beginning Security Deposit", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Change In Security Deposit", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Ending Security Deposit", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)    // account for the associated RentalAgreement

	const (
		RName        = 0
		RType        = iota
		RTSqFt       = iota
		RUsers       = iota
		RPayors      = iota
		RAgr         = iota
		UseStart     = iota
		UseStop      = iota
		RentStart    = iota
		RentStop     = iota
		RAgrStart    = iota
		RAgrStop     = iota
		RCycle       = iota
		GSRRate      = iota
		GSRAmt       = iota
		IncOff       = iota
		ContractRent = iota
		OtherInc     = iota
		PmtRcvd      = iota
		BeginRcv     = iota
		ChgRcv       = iota
		EndRcv       = iota
		BeginSecDep  = iota
		ChgSecDep    = iota
		EndSecDep    = iota
	)

	// loop through the Rentables...
	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Xbiz.P.BID)
	rlib.Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var p rlib.Rentable
		rlib.Errcheck(rows.Scan(&p.RID, &p.BID, &p.Name, &p.AssignmentTime, &p.LastModTime, &p.LastModBy)) // read the rentable
		p.RT = rlib.GetRentableTypeRefsByRange(p.RID, d1, d2)                                              // its RentableType is time sensitive

		rtid := p.RT[0].RTID // select its value at the beginning of this period
		sqft := int64(0)     // assume no custom attribute
		usernames := ""      // this will be the list of renters
		payornames := ""     // this will be the list of Payors
		rentCycle := ""

		if len(ri.Xbiz.RT[rtid].CA) > 0 { // if there are custom attributes
			c, ok := ri.Xbiz.RT[rtid].CA[custom] // see if Square Feet is among them
			if ok {                              // if it is...
				sqft, _ = rlib.IntFromString(c.Value, "invalid sqft of custom attribute")
			}
		}

		rentableTblRowStart := len(tbl.Row) // starting row for this rentable

		//------------------------------------------------------------------------------
		// Get the RentalAgreement IDs for this rentable over the time range d1,d2.
		// Note that this could result in multiple rental agreements.
		//------------------------------------------------------------------------------
		rra := rlib.GetAgreementsForRentable(p.RID, d1, d2) // get all rental agreements for this period
		// if len(rra) == 0 {                                  // if there are none...
		// 	tbl.AddRow() // puts to row -1 will go to the newly added row
		// 	tbl.Puts(-1, RName, "vacant")
		// 	tbl.Puts(-1, RType, ri.Xbiz.RT[rtid].Style)
		// 	tbl.Puti(-1, RTSqFt, sqft)
		// }

		for i := 0; i < len(rra); i++ { //for each rental agreement id
			ra, err := rlib.GetRentalAgreement(rra[i].RAID) // load the agreement
			if err != nil {
				fmt.Printf("Error loading rental agreement %d: err = %s\n", rra[i].RAID, err.Error())
				continue
			}
			na := p.GetUserNameList(&ra.PossessionStart, &ra.PossessionStop) // get the list of user names for this time period
			usernames = strings.Join(na, ",")                                // concatenate with a comma separator
			pa := ra.GetPayorNameList(&ra.RentStart, &ra.RentStart)          // get the payors for this time period
			payornames = strings.Join(pa, ", ")                              // concatenate with comma

			//-------------------------------------------------------------------------------------------------------
			// Get the rent cycle.  If there's an override in the RentableTypeRef, use the override. Otherwise the
			// rent cycle comes from the RentableType.
			//-------------------------------------------------------------------------------------------------------
			rcl := rlib.GetRentCycleRefList(&p, d1, d2, ri.Xbiz) // this sets r.RT to the RentableTypeRef list for d1-d2
			cycleval := rcl[len(rcl)-1].RentCycle                // save for proration use below
			prorateval := rcl[len(rcl)-1].ProrationCycle         // save for proration use below
			rentCycle = rlib.RentalPeriodToString(cycleval)      // use the rentCycle for the last day of the month

			//-------------------------------------------------------------------------------------------------------
			// Adjust the period as needed.  The request is to cover d1 - d2.  We start by setting dtstart and dtstop
			// to this range. If the renter moves in after d1, then adjust dtstart accordingly.  If the renter moves
			// out prior to d2 then adjust dtstop accordingly
			//-------------------------------------------------------------------------------------------------------
			dtstart := *d1
			if ra.RentStart.After(dtstart) {
				dtstart = ra.RentStart
			}
			dtstop := *d2
			if ra.RentStop.Before(dtstop) {
				dtstop = ra.RentStop
			}
			gsr, gsrRate := ComputeGSRandGSRRate(&p, &dtstart, &dtstop, ri.Xbiz)

			//-------------------------------------------------------------------------------------------------------
			// Get the contract rent
			// Remember that we're looping through all the rental all the rental agreements for Rentable p during the
			// period d1 - d2.  We just need to look at the RentalAgreementRentable for ra.RAID during d1-d2 and
			// adjust the start or stop if the rental agreement started after d1 or ended before d2.
			//-------------------------------------------------------------------------------------------------------
			rar, err := rlib.FindAgreementByRentable(p.RID, &dtstart, &dtstop)
			if err != nil {
				fmt.Printf("Error getting RentalAgreementRentable for RID = %d, period = %s - %s: err = %s\n",
					p.RID, dtstart.Format(rlib.RRDATEFMT3), dtstop.Format(rlib.RRDATEFMT3), err.Error())
				continue
			}

			//-------------------------------------------------------------------------------------------------------
			// Make any proration necessary to the gsr based on the date range d1-d2
			//-------------------------------------------------------------------------------------------------------
			pf, _, _, dt1, _ := rlib.CalcProrationInfo(&dtstart, &dtstop, d1, d2, cycleval, prorateval)
			numCycles := dtstop.Sub(dtstart) / rlib.CycleDuration(cycleval, dt1)
			contractRentVal := float64(0)
			if dtstop.After(dtstart) {
				contractRentVal = pf * rar.ContractRent
				if numCycles > 1 {
					contractRentVal += float64(numCycles-1) * rar.ContractRent
				}
			}

			//-------------------------------------------------------------------------------------------------------
			// Determine the LID of "Income Offsets" and "Other Income" accounts...
			//-------------------------------------------------------------------------------------------------------
			icos := float64(0)
			incOffsetAcct := rlib.GetLIDFromGLAccountName(ri.Xbiz.P.BID, IncomeOffsetGLAccountName)
			if incOffsetAcct == 0 {
				rlib.Ulog("RentRollTextReport: WARNING. IncomeOffsetGLAccountName = %q was not found in the GLAccounts\n", IncomeOffsetGLAccountName)
			}
			if incOffsetAcct > 0 {
				icosd1 := rlib.GetRAAccountBalance(ri.Xbiz.P.BID, incOffsetAcct, ra.RAID, &dtstart)
				icosd2 := rlib.GetRAAccountBalance(ri.Xbiz.P.BID, incOffsetAcct, ra.RAID, &dtstop)
				icos = icosd2 - icosd1
			}
			oic := float64(0)
			otherIncomeAcct := rlib.GetLIDFromGLAccountName(ri.Xbiz.P.BID, OtherIncomeGLAccountName)
			if otherIncomeAcct == 0 {
				rlib.Ulog("RentRollTextReport: WARNING. OtherIncomeGLAccountName = %q was not found in the GLAccounts\n", OtherIncomeGLAccountName)
			}
			if otherIncomeAcct > 0 {
				oicd1 := rlib.GetRAAccountBalance(ri.Xbiz.P.BID, otherIncomeAcct, ra.RAID, &dtstart)
				oicd2 := rlib.GetRAAccountBalance(ri.Xbiz.P.BID, otherIncomeAcct, ra.RAID, &dtstop)
				oic = oicd1 - oicd2 // I know this looks backwards. But in the report we want this number to show up as positive (normally), so we want -1 * (oicd2-oicd1)
			}

			//-------------------------------------------------------------------------------------------------------
			// Payments received... or more precisely that portion of a Receipt that went to pay an Assessment on
			// on this Rentable during this period d1 - d2.  We expand the search range to the entire report range
			//-------------------------------------------------------------------------------------------------------
			rcpts := rlib.GetReceiptsInRAIDDateRange(p.BID, ra.RAID, d1, d2) // receipts for ra.RAID during d1-d2, ReceiptAllocations are also loaded
			totpmt := float64(0)
			for j := 0; j < len(rcpts); j++ {
				for k := 0; k < len(rcpts[j].RA); k++ { // for each ReceiptAllocation read the Assessment
					a, err := rlib.GetAssessment(rcpts[j].RA[k].ASMID) // if Rentable == p.RID, we found the PaymentReceived value
					if err != nil {
						fmt.Printf("%s: Error calculating GSR for Rentable %d: err = %s\n", funcname, p.RID, err.Error())
						continue
					}
					if a.RID == p.RID {
						totpmt += rcpts[j].RA[k].Amount
					}
				}
			}

			//-------------------------------------------------------------------------------------------------------
			// Compute account balances...   begin, delta, and end for  RAbalance and Security Deposit
			//-------------------------------------------------------------------------------------------------------
			raStartBal := rlib.GetRAAccountBalance(ri.Xbiz.P.BID, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].DefaultAccts[rlib.GLGENRCV].LID, ra.RAID, d1)
			raEndBal := rlib.GetRAAccountBalance(ri.Xbiz.P.BID, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].DefaultAccts[rlib.GLGENRCV].LID, ra.RAID, d2)
			secdepStartBal := rlib.GetRAAccountBalance(ri.Xbiz.P.BID, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].DefaultAccts[rlib.GLSECDEP].LID, ra.RAID, d1)
			secdepEndBal := rlib.GetRAAccountBalance(ri.Xbiz.P.BID, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].DefaultAccts[rlib.GLSECDEP].LID, ra.RAID, d2)

			tbl.AddRow()
			tbl.Puts(-1, RName, p.Name)
			tbl.Puts(-1, RType, ri.Xbiz.RT[rtid].Style)
			tbl.Puti(-1, RTSqFt, sqft)
			tbl.Puts(-1, RUsers, usernames)
			tbl.Puts(-1, RPayors, payornames)
			tbl.Puts(-1, RAgr, ra.IDtoString())
			tbl.Putd(-1, UseStart, ra.PossessionStart)
			tbl.Putd(-1, UseStop, ra.PossessionStop)
			tbl.Putd(-1, RentStart, ra.RentStart)
			tbl.Putd(-1, RentStop, ra.RentStop)
			tbl.Putd(-1, RAgrStart, ra.AgreementStart)
			tbl.Putd(-1, RAgrStop, ra.AgreementStop)
			tbl.Puts(-1, RCycle, rentCycle)
			tbl.Putf(-1, GSRRate, gsrRate)
			tbl.Putf(-1, GSRAmt, gsr)
			tbl.Putf(-1, IncOff, icos)
			tbl.Putf(-1, ContractRent, contractRentVal)
			tbl.Putf(-1, OtherInc, oic)
			tbl.Putf(-1, PmtRcvd, totpmt)
			tbl.Putf(-1, BeginRcv, raStartBal)
			tbl.Putf(-1, ChgRcv, raEndBal-raStartBal)
			tbl.Putf(-1, EndRcv, raEndBal)
			tbl.Putf(-1, BeginSecDep, -secdepStartBal)
			tbl.Putf(-1, ChgSecDep, secdepStartBal-secdepEndBal)
			tbl.Putf(-1, EndSecDep, -secdepEndBal)
			// fmt.Printf("secdepEndBal = %8.2f, secdepStartBal = %8.2f,  diff = %8.2f\n", secdepEndBal, secdepStartBal, secdepEndBal-secdepStartBal)
		}

		//-------------------------------------------------------------------------------------------------------
		// All rental agreements have been process.  Look for vacancies
		//-------------------------------------------------------------------------------------------------------
		v := rlib.VacancyDetect(ri.Xbiz, d1, d2, &p)
		for i := 0; i < len(v); i++ {
			gsr, gsrRate := ComputeGSRandGSRRate(&p, &v[i].DtStart, &v[i].DtStop, ri.Xbiz)

			icos := float64(0)
			incOffsetAcct := rlib.GetLIDFromGLAccountName(p.BID, IncomeOffsetGLAccountName)
			if incOffsetAcct == 0 {
				rlib.Ulog("RentRollTextReport: WARNING. IncomeOffsetGLAccountName = %q was not found in the GLAccounts\n", IncomeOffsetGLAccountName)
			} else {
				icosd1 := rlib.GetRentableAccountBalance(ri.Xbiz.P.BID, incOffsetAcct, p.RID, d1)
				icosd2 := rlib.GetRentableAccountBalance(ri.Xbiz.P.BID, incOffsetAcct, p.RID, d2)
				icos = icosd2 - icosd1
			}

			m := rlib.GetRentableStatusByRange(p.RID, d1, d2)
			lastRStat := m[len(m)-1].Status
			tbl.AddRow()
			tbl.Puts(-1, RName, p.Name)
			tbl.Puts(-1, RType, ri.Xbiz.RT[rtid].Style)
			tbl.Puti(-1, RTSqFt, sqft)
			tbl.Puts(-1, RUsers, rlib.RentableStatusToString(lastRStat))
			tbl.Puts(-1, RPayors, "vacant")
			tbl.Puts(-1, RAgr, "n/a")
			// tbl.Putd(-1, UseStart, ra.PossessionStart)
			// tbl.Putd(-1, UseStop, ra.PossessionStop)
			tbl.Putd(-1, RentStart, v[i].DtStart)
			tbl.Putd(-1, RentStop, v[i].DtStop)
			// tbl.Putd(-1, RAgrStart, ra.AgreementStart)
			// tbl.Putd(-1, RAgrStop, ra.AgreementStop)
			tbl.Putf(-1, GSRRate, gsrRate)
			tbl.Putf(-1, GSRAmt, gsr)
			tbl.Putf(-1, IncOff, icos)
			// tbl.Putf(-1, ContractRent, contractRentVal)
			// tbl.Putf(-1, OtherInc, oic)
			// tbl.Putf(-1, PmtRcvd, oic)
			// tbl.Putf(-1, BeginRcv, raStartBal)
			// tbl.Putf(-1, ChgRcv, raEndBal-raStartBal)
			// tbl.Putf(-1, EndRcv, raEndBal)
			// tbl.Putf(-1, BeginSecDep, secdepStartBal)
			// tbl.Putf(-1, ChgSecDep, secdepEndBal-raStartBal)
			// tbl.Putf(-1, EndSecDep, secdepEndBal)
		}

		rentableTblRowStop := len(tbl.Row) - 1                       // ending row for this rentable
		tbl.Sort(rentableTblRowStart, rentableTblRowStop, RentStart) // chronologically sort these rows

		// if there are multiple rows for this rentable, add a line and give a subtotal
		if rentableTblRowStop-rentableTblRowStart > 0 {
			tbl.AddLineAfter(rentableTblRowStop)
			tbl.InsertSumRow(rentableTblRowStop+1, rentableTblRowStart, rentableTblRowStop,
				[]int{GSRAmt, IncOff, ContractRent, OtherInc, PmtRcvd, BeginRcv, ChgRcv, EndRcv, BeginSecDep, ChgSecDep, EndSecDep})
			tbl.AppendToRowset(totalsRSet, rentableTblRowStop+1) // in this case, add the sum row to the totalsRSet
		} else {
			tbl.AppendToRowset(totalsRSet, rentableTblRowStart) // add this row to the totalsRSET
		}
		tbl.AddRow() // Can't look ahead with rows.Next, so always add a blank line, remove the last one after loop ends.  See note on DeleteRow below.
	}
	rlib.Errcheck(rows.Err())
	if len(tbl.Row) > 0 {
		tbl.DeleteRow(len(tbl.Row) - 1)    // removes the last blank line. Can't check rows.Next twice, so no other way I can see to do this
		tbl.AddLineAfter(len(tbl.Row) - 1) // a line after the last row in the table
		tbl.InsertSumRowsetCols(totalsRSet, len(tbl.Row),
			[]int{GSRAmt, IncOff, ContractRent, OtherInc, PmtRcvd, BeginRcv, ChgRcv, EndRcv, BeginSecDep, ChgSecDep, EndSecDep})

		tbl.TightenColumns()
	}

	return tbl, noerr
}
