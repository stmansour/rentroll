package rlib

import (
	"fmt"
	"time"
)

//========================================================================================================
// This module calculates rent totals:
//		CalculateGSR - calculates Gross Scheduled Rent based on increments of GSRPC and MarketRate information
//		CalculateCR - calculates contract rent based on increments of GSRPC
//
// Each Rentable is a particular Rentable type.
// Each Rentable type has a "GSRPC" - the increment of time on which rent is calculated.
// This is different than the rent cycle. The rental rate may change during the rent cycle.  When this happens
// the actual rent is calculated in "GSRPC" increments over the rent cycle.  The rate at the
// beginning of the increment will be applied for the entire increment, even if the rental rate changes during
// a given increment. Any change to the rental rates will be applied on a per-increment basis.
//
// Example:
//     RentableType "A" has a MarketRate of $1000 per month, which is in effect from 2013-Jan-01 through 2016-Jul-15
//     On July 15 the MarketRate changes to $1200 per month.
// Problem:
//     What is the GrossScheduledRent for July 2016?
// Method:
//     The GSRPC for RentableType "A" is 1 day. So, rent is calculated by the day for each of the 31 days in July.
//         For July  1 through July 14 the rent will be: (1000 $ per month) / (31 days/month) * (14 days) =  451.61
//         for July 15 through July 31 the rent will be: (1200 $ per month) / (31 days/month) * (17 days) =  658.06
//         Total GSRPC for July...........................................................................: 1109.68
//
// It doesn't matter how many times the MarketRate rent changes during a monthly period, the rent will be calculated
// based on the Market Rate in effect at the beginning of each increment of time.
//
// This method can be used to calculate prorated rent as well.  If the rent does not change during the month, the
// calculation is straightforward:  (Contract rent amount per month) / (number of days in the month) * (days occupied).
//
// The functions below implement this method of rent calculation.
//========================================================================================================

// FindApplicableMarketRate returns the market rate in effect at the datetime provided
// Params:
//   dt = the datetime for which we want the rate
//    m = array of MarketRate structs
//========================================================================================================
func FindApplicableMarketRate(dt, start, stop time.Time, mr []RentableMarketRate) float64 {
	// fmt.Printf("FindApplicableMarketRate:  dt = %s, start = %s, stop = %s, len(mr) = %d\n",
	// 	dt.Format(RRDATEINPFMT), start.Format(RRDATEINPFMT), stop.Format(RRDATEINPFMT), len(mr))
	var rate = float64(0)
	for i := 0; i < len(mr); i++ {
		if DateInRange(&dt, &mr[i].DtStart, &mr[i].DtStop) {
			rate = mr[i].MarketRate
			break
		}
	}
	return rate
}

// GetProrationCycle returns either the override proration or the Rentable's RentableType proration for
// the supplied date.
// Params:
//    dt = The time/date for which we need information
//   rta = an array of RentableTypeRefs for the supplied date/time
//
// Returns:
//   RentCycle
//   ProrationCycle
//   GSRPC
//   err
//========================================================================================================
func GetProrationCycle(dt *time.Time, r *Rentable, rta *[]RentableTypeRef, xbiz *XBusiness) (int64, int64, int64, error) {
	rentCycle := int64(-1)
	prorationCycle := int64(-1)
	var err error
	rt := SelectRentableTypeRefForDate(rta, dt)

	// fmt.Printf("GetProrationCycle: rt = (%s - %s) rentcycle=%d, prorate=%d, rtid=%d\n",
	// 	rt.DtStart.Format("1/2/06"), rt.DtStop.Format("1/2/06"), rt.RentCycle, rt.ProrationCycle, rt.RTID)

	if rt.ProrationCycle > ACCRUALNORECUR { // if there's an override
		prorationCycle = rt.ProrationCycle //use the override
	}
	if rt.RentCycle > ACCRUALNORECUR { // if there's an override...
		rentCycle = rt.RentCycle // ...use it
	}

	// determine the rentable type for time dt
	if rt.RTID == 0 {
		err = fmt.Errorf("GetProrationCycle:  No valid RTID for rentable R%08d during period %s\n", r.RID, dt.Format(RRDATEINPFMT))
		return 0, 0, 0, err // this is bad! No RTID for the supplied time range
	}
	if prorationCycle < 0 { // if there was no override..
		prorationCycle = xbiz.RT[rt.RTID].Proration
	}
	if rentCycle < 0 {
		rentCycle = xbiz.RT[rt.RTID].RentCycle
	}
	return rentCycle, prorationCycle, xbiz.RT[rt.RTID].GSRPC, err
}

// CalculateGSR calculates the gross scheduled rent as described above.
// Params:
//   d1 = start datetime of the period
//   d2 = stop datetime of the period
//   rt = array of RentableMarketRate structures that covers all rental rates during the period d1 - d2.
//        This array is the MR attribute in the RentableMarketRate struct
//  rsa = array of rentable specialties that apply to the rentable we're calculating
//========================================================================================================
func CalculateGSR(d1, d2 time.Time, r *Rentable, rta *[]RentableTypeRef, rsa []RentableSpecialtyType, xbiz *XBusiness) float64 {
	var total = float64(0) // init total

	// rentCycle, prorateCycle, gsrpc, err := GetProrationCycle(&d1, r, rta, xbiz)
	rentCycle, _, gsrpc, err := GetProrationCycle(&d1, r, rta, xbiz)
	if err != nil {
		Ulog("CalculateGSR: GetProrationCycle returned error: %s\n", err.Error())
		return float64(0)
	}
	if rentCycle < 0 || gsrpc < 0 {
		Ulog("CalculateGSR: warning: one or more cycle values is unset\n")
	}
	// prorateDur := CycleDuration(prorateCycle, d1) // the proration cycle expressed as a duration
	inc := CycleDuration(gsrpc, d1)              // increment durations for rent calculation
	rentCycleDur := CycleDuration(rentCycle, d1) // this is the rentcycle expressed as a duration
	rtr := SelectRentableTypeRefForDate(rta, &d1)

	// fmt.Printf("CalculateGSR: rentCycle = %d (%v), prorateCycle = %d (%v), gsrpc = %d (%v)\n",
	// 	rentCycle, rentCycleDur, prorateCycle, prorateDur, gsrpc, inc)
	// fmt.Printf("rtr = (%s - %s) rtid = %d\n", rtr.DtStart.Format("1/2/06"), rtr.DtStop.Format("1/2/06"), rtr.RTID)

	for d := d1; d.Before(d2); d = d.Add(inc) { // spin through the period in the defined increments
		rate := FindApplicableMarketRate(d, d1, d2, xbiz.RT[rtr.RTID].MR) // find the rate applicable for this increment
		rent := float64(inc) * rate / float64(rentCycleDur)               // how much for the period: inc
		total += rent                                                     // increment the total by this amount
		for i := 0; i < len(rsa); i++ {
			total += rsa[i].Fee * float64(inc) / float64(rentCycleDur)
		}
	}
	return total
}

// CalculateLoadedGSR calculates the gross scheduled rent including any Specialties associated with the rentable.
// Params:
//   d1 = start datetime of the period
//   d2 = stop datetime of the period
//   rt = array of RentableMarketRate structures that covers all rental rates during the period d1 - d2.
//        This array is the MR attribute in the RentableMarketRate struct
//========================================================================================================
func CalculateLoadedGSR(r *Rentable, d1, d2 *time.Time, xbiz *XBusiness) (float64, error) {
	funcname := "CalculateLoadedGSR"
	var err error
	gsr := float64(0) // total rent, to update on each pass through the loop below

	// fmt.Printf("%s, r = %#v, d1 = %s, d2 = %s\n", funcname, *r, d1.Format(RRDATEINPFMT), d2.Format(RRDATEINPFMT))

	rta := GetRentableTypeRefsByRange(r.RID, d1, d2) // get the list
	if len(rta) == 0 {
		err = fmt.Errorf("%s:  No valid RTID for rentable R%08d during period %s to %s\n",
			funcname, r.RID, d1.Format(RRDATEINPFMT), d2.Format(RRDATEINPFMT))
		Ulog("%s", err.Error())
		return gsr, err // this is bad! No RTID for the supplied time range
	}

	// find the Gross Scheduled Rent Proration Cycle - GSRPC - the intervals over which the GSR is calculated
	_, _, gsrpc, err := GetProrationCycle(d1, r, &rta, xbiz)
	if err != nil {
		return gsr, err
	}
	period := CycleDuration(gsrpc, *d1)          // increment of time we'll use to determine gsr in increments between d1 & d2
	dtNext := *d1                                // initialize so that the variable is known
	for dt := *d1; dt.Before(*d2); dt = dtNext { // spin through time period d1 - d2 in increments of gsrpc and add up the GSR
		dtNext = dt.Add(period) // establish the end of the period.  We'll add up the gsr for period dt to dtNext.
		//--------------------------------------------------------------------
		// Get the RentableSpecialties applicable for this increment...
		//--------------------------------------------------------------------
		rsa, nerr := GetRentableSpecialtyTypesForRentableByRange(r, &dt, &dtNext) // this gets an array of rentable specialties that overlap this time period
		if nerr != nil {
			err = fmt.Errorf("%s:  error getting specialties for rentable R%08d during period %s to %s.  err = %s\n",
				funcname, r.RID, dt.Format(RRDATEINPFMT), dtNext.Format(RRDATEINPFMT), nerr.Error())
			Ulog("%s", err.Error())
			break
		}
		//------------------------------------------------------------------
		// Finally, calculate the GSR for this increment...
		//------------------------------------------------------------------
		rentThisPeriod := CalculateGSR(dt, dtNext, r, &rta, rsa, xbiz)
		gsr += rentThisPeriod
	}
	return gsr, err
}
