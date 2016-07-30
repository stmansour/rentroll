package rrpt

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"
)

// RentRollTextReport generates a text-based RentRoll report for the supplied business and timeframe
func RentRollTextReport(xbiz *rlib.XBusiness, d1, d2 *time.Time) error {
	funcname := "RentRollTextReport"
	custom := "Square Feet"
	var noerr error
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
	tbl.AddColumn("Rentable", "s", 10, 0)         // column for the Rentable name
	tbl.AddColumn("Rentable Type", "s", 15, 0)    // RentableType name
	tbl.AddColumn(custom, "s", 5, 1)              // the Custom Attribute "Square Feet"
	tbl.AddColumn("Rentable Users", "s", 30, 0)   // Users of this rentable
	tbl.AddColumn("Rentable Payors", "s", 30, 0)  // Users of this rentable
	tbl.AddColumn("Rental Agreement", "s", 10, 0) // the Rental Agreement id
	tbl.AddColumn("Use Start", "s", 12, 0)        // the possession start date
	tbl.AddColumn("Use Stop", "s", 12, 0)         // the possession start date
	tbl.AddColumn("Rental Start", "s", 12, 0)     // the rental start date
	tbl.AddColumn("Rental Stop", "s", 12, 0)      // the rental start date
	tbl.AddColumn("Rent Cycle", "s", 12, 0)       // the rent cycle
	tbl.AddColumn("GSR Rate", "s", 10, 1)         // gross scheduled rent
	tbl.AddColumn("GSR This Period", "s", 10, 1)  // gross scheduled rent
	tbl.AddColumn("Contract Rent", "s", 10, 1)    // contract rent amounts
	tbl.AddColumn("Payment Received", "s", 10, 1) // contract rent amounts

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
		rentCycle := ""      // how rent accrues
		gsrstr := ""         // Gross Scheduled Rent
		gsrRateStr := ""     // GSR Rate this period
		contractRent := ""   // contract rent
		pmtRcvd := ""        // payment

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
			tbl.Printf(p.Name, xbiz.RT[rtid].Name, sqft, usernames, payornames, raid, possStart, possStop,
				rentStart, rentStop, rentCycle, gsrRateStr, gsrstr, contractRent, pmtRcvd)
		}
		for i := 0; i < len(rra); i++ { //for each rental agreement id
			ra, err := rlib.GetRentalAgreement(rra[i].RAID) // load the agreement
			if err != nil {
				fmt.Printf("Error loading rental agreement %d: err = %s\n", rra[i].RAID, err.Error())
				continue
			}
			na := ra.GetUserNameList(d1, d2)                       // get the list of user names for this time period
			raid = ra.IDtoString()                                 // standard id format
			usernames = strings.Join(na, ",")                      // concatenate with a comma separator
			pa := ra.GetPayorNameList(d1, d2)                      // get the payors for this time period
			payornames = strings.Join(pa, ", ")                    // concatenate with comma
			rentStart = ra.AgreementStart.Format(rlib.RRDATEFMT4)  // rental start
			rentStop = ra.AgreementStop.Format(rlib.RRDATEFMT4)    // rental stop
			possStart = ra.PossessionStart.Format(rlib.RRDATEFMT4) // possession start
			possStop = ra.PossessionStop.Format(rlib.RRDATEFMT4)   // possession stop

			//-------------------------------------------------------------------------------------------------------
			// Get the rent cycle.  If there's an override in the RentableTypeRef, use the override. Otherwise the
			// rent cycle comes from the RentableType.
			//-------------------------------------------------------------------------------------------------------
			rcl := rlib.GetRentCycleRefList(&p, d1, d2, xbiz) // this sets r.RT to the RentableTypeRef list for d1-d2
			cycleval := rcl[len(rcl)-1].RentCycle             // save for proration use below
			prorateval := rcl[len(rcl)-1].ProrationCycle      // save for proration use below
			rentCycle = rlib.RentalPeriodToString(cycleval)   // use the rentCycle for the last day of the month

			//-------------------------------------------------------------------------------------------------------
			// Calculate the Gross Scheduled Rent for this Rentable.  We have most of what we need, but we do need
			// to fetch the RentableSpecialtyTypes
			//-------------------------------------------------------------------------------------------------------
			gsr, err := rlib.CalculateLoadedGSR(&p, d1, d2, xbiz)
			if err != nil {
				fmt.Printf("%s: Error calculating GSR for Rentable %d: err = %s\n", funcname, p.RID, err.Error())
				continue
			}
			gsrstr = rlib.RRCommaf(gsr)

			//-------------------------------------------------------------------------------------------------------
			// Calculate the Gross Scheduled Rent Rate for this period. The rate will be the amount divided by the
			// number of periods...
			//-------------------------------------------------------------------------------------------------------
			periods := rlib.GetRentableCycles(&p, d1, d2, xbiz)
			gsrRate := gsr / float64(periods)
			gsrRateStr = rlib.RRCommaf(gsrRate)

			//-------------------------------------------------------------------------------------------------------
			// Get the contract rent
			// Remember that we're looping through all the rental all the rental agreements for Rentable p during the
			// period d1 - d2.  We just need to look at the RentalAgreementRentable for ra.RAID during d1-d2 and
			// adjust the start or stop if the rental agreement started after d1 or ended before d2.
			//-------------------------------------------------------------------------------------------------------
			dtstart := *d1
			if ra.AgreementStart.After(dtstart) {
				dtstart = ra.AgreementStart
			}
			dtstop := *d2
			if ra.AgreementStop.Before(dtstop) {
				dtstop = ra.AgreementStop
			}
			rar, err := rlib.FindAgreementByRentable(p.RID, &dtstart, &dtstop)
			if err != nil {
				fmt.Printf("Error getting RentalAgreementRentable for RID = %d, period = %s - %s: err = %s\n",
					p.RID, dtstart.Format(rlib.RRDATEFMT3), dtstop.Format(rlib.RRDATEFMT3), err.Error())
				continue
			}
			//-------------------------------------------------------------------------------------------------------
			// Make any proration necessary to the gsr based on the date range d1-d2
			//-------------------------------------------------------------------------------------------------------
			pf, _, _, _, _ := rlib.CalcProrationInfo(&ra.PossessionStart, &ra.PossessionStop, d1, d2, cycleval, prorateval)
			contractRentVal := pf * rar.ContractRent
			contractRent = rlib.RRCommaf(contractRentVal)

			//-------------------------------------------------------------------------------------------------------
			// Payments received... or more precisely that portion of a Receipt that went to pay an Assessment on
			// on this Rentable during this period d1 - d2
			//-------------------------------------------------------------------------------------------------------

			// get all the receipts for ra.RAID that occurred during d1-d2
			rcpts := rlib.GetReceiptsInRAIDDateRange(p.BID, ra.RAID, d1, d2) // this has the ReceiptAllocations already loaded
			for j := 0; j < len(rcpts); j++ {
				fmt.Printf("rcpts[%d] -> %s\n", j, rcpts[j].IDtoString())
				// for each ReceiptAllocation read the Assessment
				for k := 0; k < len(rcpts[j].RA); k++ {
					// if the Assessment's Rentable is p.RID then we have found the PaymentReceived value
					a, err := rlib.GetAssessment(rcpts[j].RA[k].ASMID)
					if err != nil {
						fmt.Printf("%s: Error calculating GSR for Rentable %d: err = %s\n", funcname, p.RID, err.Error())
						continue
					}
					if a.RID == p.RID {
						pmtRcvd := rlib.RRCommaf(rcpts[j].RA[k].Amount)
						tbl.Printf(p.Name, xbiz.RT[rtid].Name, sqft, usernames, payornames, raid, possStart, possStop,
							rentStart, rentStop, rentCycle, gsrRateStr, gsrstr, contractRent, pmtRcvd)

					}
				}

			}
		}
		fmt.Printf("\n")
	}
	rlib.Errcheck(rows.Err())
	tbl.PrintLine()

	return noerr
}
