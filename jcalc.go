package main

import (
	"rentroll/rlib"
	"time"
)

//========================================================================================================
// this module calculates rent totals:
//		CalculateGSR - calculates Gross Scheduled Rent based on increments of GSRPC and MarketRate information
//		CalculateCR - calculates contract rent based on increments of GSRPC
//
// Each Rentable is a particular Rentable type.
// Each Rentable type has a "calculateFrequency" - the increment of time on which rent is calculated.
// This is different than the rent cycle. The rental rate may change during the rent cycle.  When this happens
// the actual rent is calculated in "calculateFrequency" increments over the rent cycle.  The rate at the
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
// But what if the rent changes mid-cycle?
//
// Example:
//     Sue rents an apartment of RentableType "A".  Her contract rent is $950/month. Her rent is scheduled to be raised
//     on July 15 to the MarketRate of $1200/month. She moves out of the apartment in July. She occupies it from July 1
//     through July 18, and she is gone on July 19.
// Problem:
//     What is Sue's rent for July?
// Method:
//     Proration on Sue's apartment is "daily".  So, her rent is:
//         July  1 - July 14, her contract rent is (950 $ per month) / (31 days per month) * 14 days = 398.39
//         July 15 - July 18, her contract rent is (1200 $ per month) / (31 days per month) * 4 days = 154.84
//         Total rent for July.......................................................................: 553.23
//
// The functions below implement this method of rent calculation.
//========================================================================================================

// FindApplicableMarketRate returns the market rate in effect at the datetime provided
// Params:
//   dt = the datetime for which we want the rate
//    m = array of MarketRate structs
//========================================================================================================
func FindApplicableMarketRate(dt, start, stop time.Time, mr []rlib.RentableMarketRate) float64 {
	// fmt.Printf("FindApplicableMarketRate:  dt = %s, start = %s, stop = %s, len(mr) = %d\n",
	// 	dt.Format(rlib.RRDATEINPFMT), start.Format(rlib.RRDATEINPFMT), stop.Format(rlib.RRDATEINPFMT), len(mr))
	var rate = float64(0)
	for i := 0; i < len(mr); i++ {
		if rlib.DateInRange(&dt, &mr[i].DtStart, &mr[i].DtStop) {
			rate = mr[i].MarketRate
			break
		}
	}
	return rate
}

// CalculateGSR calculates the gross scheduled rent as described above.
// Params:
//   d1 = start datetime of the period
//   d2 = stop datetime of the period
//   rt = array of RentableMarketRate structures that covers all rental rates during the period d1 - d2.
//        This array is the MR attribute in the RentableMarketRate struct
//========================================================================================================
func CalculateGSR(d1, d2 time.Time, rt rlib.RentableType) float64 {
	var total = float64(0)                               // init total
	prorateDur := rlib.CycleDuration(rt.Proration, d1)   // the proration cycle expressed as a duration
	inc := prorateDur                                    // increment durations for rent calculation -- FOR NOW I'VE MAPPED IT TO PRORATE CYCLE
	rentCycleDur := rlib.CycleDuration(rt.RentCycle, d1) // this is the rentcycle expressed as a duration

	for d := d1; d.Before(d2); d = d.Add(inc) { // spin through the period in the defined increments
		rate := FindApplicableMarketRate(d, d1, d2, rt.MR)  // find the rate applicable for this increment
		rent := float64(inc) * rate / float64(rentCycleDur) // how much for the period: inc
		total += rent                                       // increment the total by this amount
	}
	return total
}
