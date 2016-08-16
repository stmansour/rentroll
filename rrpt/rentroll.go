package rrpt

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"
)

// OtherIncomeGLAccountName and the rest will need to become configurable parameters for this report!!
const (
	OtherIncomeGLAccountName  = string("Other Income")
	IncomeOffsetGLAccountName = string("Income Offsets")
)

// RentRollTextReport generates a text-based RentRoll report for the business in xbiz and timeframe d1 to d2.
func RentRollTextReport(xbiz *rlib.XBusiness, d1, d2 *time.Time) error {
	funcname := "RentRollTextReport"
	custom := "Square Feet"
	var noerr error
	fmt.Printf("xbiz.P.Designation = %s\n", xbiz.P.Designation)
	bu, err := rlib.GetBusinessUnitByDesignation(xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s\n", funcname, err.Error())
		return e
	}
	c, err := rlib.GetCompany(int64(bu.CoCode))
	if err != nil {
		e := fmt.Errorf("%s: error getting Company - %s\n", funcname, err.Error())
		return e
	}
	fmt.Printf("%s\n", strings.ToUpper(c.LegalName))
	fmt.Printf("Rentroll report for period beginning %s and ending %s\n\n", d1.Format(rlib.RRDATEFMT3), d2.Format(rlib.RRDATEFMT3))

	var table TextReport
	tbl := &table
	tbl.Spacing = 2
	tbl.AddColumn("Rentable", "s", 10, 0)                   // column for the Rentable name
	tbl.AddColumn("Rentable Type", "s", 15, 0)              // RentableType name
	tbl.AddColumn(custom, "s", 5, 1)                        // the Custom Attribute "Square Feet"
	tbl.AddColumn("Rentable Users", "s", 30, 0)             // Users of this rentable
	tbl.AddColumn("Rentable Payors", "s", 30, 0)            // Users of this rentable
	tbl.AddColumn("Rental Agreement", "s", 10, 0)           // the Rental Agreement id
	tbl.AddColumn("Use Start", "s", 10, 0)                  // the possession start date
	tbl.AddColumn("Use Stop", "s", 10, 0)                   // the possession start date
	tbl.AddColumn("Rental Start", "s", 10, 0)               // the rental start date
	tbl.AddColumn("Rental Stop", "s", 10, 0)                // the rental start date
	tbl.AddColumn("Rental Agreement Start", "s", 10, 0)     // the possession start date
	tbl.AddColumn("Rental Agreement Stop", "s", 10, 0)      // the possession start date
	tbl.AddColumn("Rent Cycle", "s", 12, 0)                 // the rent cycle
	tbl.AddColumn("GSR Rate", "s", 10, 1)                   // gross scheduled rent
	tbl.AddColumn("GSR This Period", "s", 10, 1)            // gross scheduled rent
	tbl.AddColumn(IncomeOffsetGLAccountName, "s", 10, 1)    // GL Account
	tbl.AddColumn("Contract Rent", "s", 10, 1)              // contract rent amounts
	tbl.AddColumn(OtherIncomeGLAccountName, "s", 10, 1)     // GL Account
	tbl.AddColumn("Payments Received", "s", 10, 1)          // contract rent amounts
	tbl.AddColumn("Beginning Receivable", "s", 10, 1)       // account for the associated RentalAgreement
	tbl.AddColumn("Change In Receivable", "s", 10, 1)       // account for the associated RentalAgreement
	tbl.AddColumn("Ending Receivable", "s", 10, 1)          // account for the associated RentalAgreement
	tbl.AddColumn("Beginning Security Deposit", "s", 10, 1) // account for the associated RentalAgreement
	tbl.AddColumn("Change In Security Deposit", "s", 10, 1) // account for the associated RentalAgreement
	tbl.AddColumn("Ending Security Deposit", "s", 10, 1)    // account for the associated RentalAgreement

	tbl.PrintColHdr()
	tbl.PrintLine()
	// loop through the Rentables...
	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(xbiz.P.BID)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var p rlib.Rentable
		rlib.Errcheck(rows.Scan(&p.RID, &p.BID, &p.Name, &p.AssignmentTime, &p.LastModTime, &p.LastModBy)) // read the rentable
		p.RT = rlib.GetRentableTypeRefsByRange(p.RID, d1, d2)                                              // its RentableType is time sensitive

		rtid := p.RT[0].RTID // select its value at the beginning of this period
		sqft := ""           // assume no custom attribute
		usernames := ""      // this will be the list of renters
		payornames := ""     // this will be the list of Payors
		raid := ""           // assume it's vacant
		possStart := ""      // possession start date
		possStop := ""       // possession stop date
		rentStart := ""      // rental start date
		rentStop := ""       // rental stop date
		agreementStart := "" // rental start date
		agreementStop := ""  // rental stop date
		rentCycle := ""      // how rent accrues
		gsrstr := ""         // Gross Scheduled Rent
		gsrRateStr := ""     // GSR Rate this period
		incomeOffsets := ""  // Gl Account balance
		contractRent := ""   // contract rent
		otherIncome := ""    // Gl Account balance
		pmtRcvd := ""        // payment
		beginRcv := ""
		chgRcv := ""
		endRcv := ""
		beginSecDep := ""
		chgSecDep := ""
		endSecDep := ""

		if len(xbiz.RT[rtid].CA) > 0 { // if there are custom attributes
			c, ok := xbiz.RT[rtid].CA[custom] // see if Square Feet is among them
			if ok {                           // if it is...
				sqft = c.Value // update the sqft value to the custom attribute
			}
		}

		//------------------------------------------------------------------------------
		// Get the RentalAgreement IDs for this rentable over the time range d1,d2.
		// Note that this could result in multiple rental agreements.
		//------------------------------------------------------------------------------
		rra := rlib.GetAgreementsForRentable(p.RID, d1, d2) // get all rental agreements for this period
		if len(rra) == 0 {                                  // if there are none...
			usernames = "vacant"
			tbl.Printf(p.Name, xbiz.RT[rtid].Style, sqft, usernames, payornames, raid, possStart, possStop,
				rentStart, rentStop, agreementStart, agreementStop, rentCycle, gsrRateStr, gsrstr, incomeOffsets, contractRent, otherIncome, pmtRcvd,
				beginRcv, chgRcv, endRcv, beginSecDep, chgSecDep, endSecDep)
		}

		for i := 0; i < len(rra); i++ { //for each rental agreement id
			ra, err := rlib.GetRentalAgreement(rra[i].RAID) // load the agreement
			if err != nil {
				fmt.Printf("Error loading rental agreement %d: err = %s\n", rra[i].RAID, err.Error())
				continue
			}
			na := p.GetUserNameList(d1, d2)                            // get the list of user names for this time period
			raid = ra.IDtoString()                                     // standard id format
			usernames = strings.Join(na, ",")                          // concatenate with a comma separator
			pa := ra.GetPayorNameList(d1, d2)                          // get the payors for this time period
			payornames = strings.Join(pa, ", ")                        // concatenate with comma
			rentStart = ra.RentStart.Format(rlib.RRDATEFMT4)           // rental start
			rentStop = ra.RentStop.Format(rlib.RRDATEFMT4)             // rental stop
			possStart = ra.PossessionStart.Format(rlib.RRDATEFMT4)     // possession start
			possStop = ra.PossessionStop.Format(rlib.RRDATEFMT4)       // possession stop
			agreementStart = ra.AgreementStart.Format(rlib.RRDATEFMT4) // agreement start
			agreementStop = ra.AgreementStop.Format(rlib.RRDATEFMT4)   // agreement stop

			//-------------------------------------------------------------------------------------------------------
			// Get the rent cycle.  If there's an override in the RentableTypeRef, use the override. Otherwise the
			// rent cycle comes from the RentableType.
			//-------------------------------------------------------------------------------------------------------
			rcl := rlib.GetRentCycleRefList(&p, d1, d2, xbiz) // this sets r.RT to the RentableTypeRef list for d1-d2
			cycleval := rcl[len(rcl)-1].RentCycle             // save for proration use below
			prorateval := rcl[len(rcl)-1].ProrationCycle      // save for proration use below
			rentCycle = rlib.RentalPeriodToString(cycleval)   // use the rentCycle for the last day of the month

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
			//-------------------------------------------------------------------------------------------------------
			// Calculate the Gross Scheduled Rent for this Rentable.  We have most of what we need, but we do need
			// to fetch the RentableSpecialtyTypes
			//-------------------------------------------------------------------------------------------------------
			gsr, err := rlib.CalculateLoadedGSR(&p, &dtstart, &dtstop, xbiz)
			if err != nil {
				fmt.Printf("%s: Error calculating GSR for Rentable %d: err = %s\n", funcname, p.RID, err.Error())
				continue
			}
			// /*DEBUG*/ fmt.Printf("CalculateLoadedGSR for %s (RTID: %d) RentCycle = %d,  %s - %s:  %6.2f\n", p.Name, rtid, xbiz.RT[rtid].RentCycle, ra.RentStart.Format(rlib.RRDATEFMT4), ra.RentStop.Format(rlib.RRDATEFMT4), gsr)
			gsrstr = rlib.RRCommaf(gsr)

			//-------------------------------------------------------------------------------------------------------
			// Calculate the Gross Scheduled Rent Rate for this period. The rate will be the amount divided by the
			// number of periods...
			//-------------------------------------------------------------------------------------------------------
			periods := rlib.GetRentableCycles(&p, &dtstart, &dtstop, xbiz)
			gsrRate := gsr / float64(periods)
			gsrRateStr = rlib.RRCommaf(gsrRate)

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
			// /*DEBUG*/pf, num, den, dt1, dt2 := rlib.CalcProrationInfo(&dtstart, &dtstop, d1, d2, cycleval, prorateval)
			pf, _, _, dt1, _ := rlib.CalcProrationInfo(&dtstart, &dtstop, d1, d2, cycleval, prorateval)
			numCycles := dtstop.Sub(dtstart) / rlib.CycleDuration(cycleval, dt1)
			contractRentVal := pf * rar.ContractRent
			if numCycles > 1 {
				contractRentVal += float64(numCycles-1) * rar.ContractRent
			}
			// /*DEBUG*/fmt.Printf("Rent start-stop: %s - %s,   d1-d2: %s - %s,   numCycles = %d\n",
			// 	dtstart.Format(rlib.RRDATEFMT4), dtstop.Format(rlib.RRDATEFMT4),
			// 	d1.Format(rlib.RRDATEFMT4), d2.Format(rlib.RRDATEFMT4), numCycles)
			// /*DEBUG*/fmt.Printf("Num/Den = %d/%d, cycleval = %d, prorateval = %d,  rar.ContractRent = %6.2f,  pf = %1.4f, contractRentVal = %6.2f,  dt1-dt2: %s - %s\n",
			// 	num, den, cycleval, prorateval, rar.ContractRent, pf, contractRentVal, dt1.Format(rlib.RRDATEFMT4), dt2.Format(rlib.RRDATEFMT4))
			contractRent = rlib.RRCommaf(contractRentVal)

			// ISSUE:  the following needs to be made general purpose

			//-------------------------------------------------------------------------------------------------------
			// Determine the LID of "Income Offsets" and "Other Income" accounts...
			//-------------------------------------------------------------------------------------------------------
			incOffsetAcct := rlib.GetLIDFromGLAccountName(xbiz.P.BID, IncomeOffsetGLAccountName)
			otherIncomeAcct := rlib.GetLIDFromGLAccountName(xbiz.P.BID, OtherIncomeGLAccountName)
			icos := float64(0)
			oic := float64(0)

			// /*DEBUG*/ fmt.Printf("incOffsetAcct = %d, otherIncomeAcct = %d\n", incOffsetAcct, otherIncomeAcct)
			if incOffsetAcct == 0 {
				rlib.Ulog("RentRollTextReport: WARNING. IncomeOffsetGLAccountName = %q was not found in the GLAccounts\n", IncomeOffsetGLAccountName)
			}
			if otherIncomeAcct == 0 {
				rlib.Ulog("RentRollTextReport: WARNING. OtherIncomeGLAccountName = %q was not found in the GLAccounts\n", OtherIncomeGLAccountName)
			}

			if incOffsetAcct > 0 {
				icos = rlib.GetAccountBalanceForDate(xbiz.P.BID, incOffsetAcct, ra.RAID, &dtstop)
			}

			if otherIncomeAcct > 0 {
				oic = rlib.GetAccountBalanceForDate(xbiz.P.BID, otherIncomeAcct, ra.RAID, &dtstop)
			}

			incomeOffsets = rlib.RRCommaf(icos)
			otherIncome = rlib.RRCommaf(oic)

			//-------------------------------------------------------------------------------------------------------
			// Payments received... or more precisely that portion of a Receipt that went to pay an Assessment on
			// on this Rentable during this period d1 - d2
			//-------------------------------------------------------------------------------------------------------

			// get all the receipts for ra.RAID that occurred during d1-d2
			rcpts := rlib.GetReceiptsInRAIDDateRange(p.BID, ra.RAID, d1, d2) // this has the ReceiptAllocations already loaded
			totpmt := float64(0)
			for j := 0; j < len(rcpts); j++ {
				// /*DEBUG*/ fmt.Printf("rcpts[%d] -> %s\n", j, rcpts[j].IDtoString())
				// for each ReceiptAllocation read the Assessment
				for k := 0; k < len(rcpts[j].RA); k++ {
					// if the Assessment's Rentable is p.RID then we have found the PaymentReceived value
					a, err := rlib.GetAssessment(rcpts[j].RA[k].ASMID)
					if err != nil {
						fmt.Printf("%s: Error calculating GSR for Rentable %d: err = %s\n", funcname, p.RID, err.Error())
						continue
					}
					if a.RID == p.RID {
						totpmt += rcpts[j].RA[k].Amount
					}
				}
			}
			pmtRcvd = rlib.RRCommaf(totpmt)

			//-------------------------------------------------------------------------------------------------------
			// Compute account balances...   begin, delta, and end for  RAbalance and Security Deposit
			//-------------------------------------------------------------------------------------------------------
			raLdg, err := rlib.GetRABalanceLedger(xbiz.P.BID, ra.RAID)
			rlib.Errcheck(err)
			secdepLdg, err := rlib.GetSecDepBalanceLedger(xbiz.P.BID, ra.RAID)
			rlib.Errcheck(err)
			raStartBal := rlib.GetAccountBalanceForDate(xbiz.P.BID, raLdg.LID, ra.RAID, &dtstart)
			raEndBal := rlib.GetAccountBalanceForDate(xbiz.P.BID, raLdg.LID, ra.RAID, &dtstop)
			secdepStartBal := rlib.GetAccountBalanceForDate(xbiz.P.BID, secdepLdg.LID, ra.RAID, &dtstart)
			secdepEndBal := rlib.GetAccountBalanceForDate(xbiz.P.BID, secdepLdg.LID, ra.RAID, &dtstop)
			beginRcv = rlib.RRCommaf(raStartBal)
			endRcv = rlib.RRCommaf(raEndBal)
			chgRcv = rlib.RRCommaf(raEndBal - raStartBal)
			beginSecDep = rlib.RRCommaf(secdepStartBal)
			endSecDep = rlib.RRCommaf(secdepEndBal)
			chgSecDep = rlib.RRCommaf(secdepEndBal - secdepStartBal)

			tbl.Printf(p.Name, xbiz.RT[rtid].Style, sqft, usernames, payornames, raid, possStart, possStop,
				rentStart, rentStop, agreementStart, agreementStop, rentCycle, gsrRateStr, gsrstr, incomeOffsets,
				contractRent, otherIncome, pmtRcvd, beginRcv, chgRcv, endRcv, beginSecDep, chgSecDep, endSecDep)
		}
		fmt.Printf("\n")
	}
	rlib.Errcheck(rows.Err())
	tbl.PrintLine()

	return noerr
}
