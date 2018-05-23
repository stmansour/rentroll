package rlib

import (
	"context"
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

// DateTrim returns the later of 2 start dates and the earlier of two stop dates
//
// Parameters:
//  	DtStart,DtStop: time range 1
//  	d1, d2:         time range 2
//
// Returns:
//			dtrim1:  later of the two start dates
//         	dtrim2:  earlier of the two end dates
//=================================================================================================
func DateTrim(DtStart, DtStop, d1, d2 *time.Time) (time.Time, time.Time) {
	start := *d1
	if DtStart.After(start) {
		start = *DtStart
	}
	stop := *DtStop // .Add(24 * 60 * time.Minute) -- removing this as all ranges must be NON-INCLUSIVE
	if stop.After(*d2) {
		stop = *d2
	}
	return start, stop
}

// CalcProrationInfo is currently designed to work for nonrecurring, daily and monthly recurring rentals.
// Other frequencies will need to be added.
//
// Parameters:
//  	DtStart,DtStop: defines the time range of either the rentalAgreement or the Assessment
//						covering the Rentable
//  	d1, d2:         the time period we're being asked to analyze
//  	accrual:        recurring frequency of rent
//  	prorateMethod:  the recurrence frequency to use for partial coverage
//
// Returns:
//			asmtDur:  pf denominator, total number of days in period
//         	rentDur:  pf numerator, total number of days applicable to this rental agreement
//         	pf:       prorate factor = rentDur/asmtDur
//=================================================================================================
func CalcProrationInfo(DtStart, DtStop, d1, d2 *time.Time, rentCycle, prorate int64) (float64, int64, int64, time.Time, time.Time) {
	// Console("CalcProrationInfo\n")
	// Console("DtStart = %s, DtStop = %s\n", DtStart.Format(RRDATEINPFMT), DtStop.Format(RRDATEINPFMT))
	// Console("     d1 = %s,     d2 = %s\n", d1.Format(RRDATEINPFMT), d2.Format(RRDATEINPFMT))
	// Console("\nrentCycle = %d, prorate = %d\n", rentCycle, prorate)

	//-------------------------------------------------------------------
	// over what range of time does this rental apply between d1 & d2
	//-------------------------------------------------------------------
	start, stop := DateTrim(DtStart, DtStop, d1, d2)
	pf := float64(1.0)            // assume full period
	cycleTime := d2.Sub(*d1)      // denominator
	thisPeriod := stop.Sub(start) // numerator
	if cycleTime != thisPeriod && // if cycle time and period differ
		prorate > 0 && // AND it's NOT a one-time charge
		rentCycle > prorate { // AND the rentcycle is larger than the proratecycle
		pf = float64(thisPeriod) / float64(cycleTime)
	}
	num := int64(0)
	den := int64(0)
	// if thisPeriod > 0 && prorate > 0 && rentCycle > prorate {
	if prorate > 0 {
		div := CycleDuration(prorate, *d1)
		num = int64(thisPeriod / div)
		den = int64(cycleTime / div)
	}
	return pf, num, den, start, stop
}

// FindApplicableMarketRate returns the market rate in effect at the datetime provided
// Params:
//   dt = the datetime for which we want the rate
//    m = array of MarketRate structs
//========================================================================================================
func FindApplicableMarketRate(dt, start, stop time.Time, mr []RentableMarketRate) float64 {
	// Console("FindApplicableMarketRate:  dt = %s, start = %s, stop = %s, len(mr) = %d\n",
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
func GetProrationCycle(ctx context.Context, dt *time.Time, rid int64, rta *[]RentableTypeRef, xbiz *XBusiness) (int64, int64, int64, error) {
	rentCycle := int64(-1)
	prorationCycle := int64(-1)
	var err error
	rt, err := SelectRentableTypeRefForDate(ctx, rta, dt)
	if err != nil {
		return rentCycle, prorationCycle, 0, err
	}

	// Console("GetProrationCycle: rt = (%s - %s) rentcycle=%d, prorate=%d, rtid=%d\n",
	// 	rt.DtStart.Format("1/2/06"), rt.DtStop.Format("1/2/06"), rt.RentCycle, rt.ProrationCycle, rt.RTID)

	if rt.OverrideProrationCycle > RECURNONE { // if there's an override
		prorationCycle = rt.OverrideProrationCycle //use the override
	}
	if rt.OverrideRentCycle > RECURNONE { // if there's an override...
		rentCycle = rt.OverrideRentCycle // ...use it
	}

	// determine the rentable type for time dt
	if rt.RTID == 0 {
		err = fmt.Errorf("GetProrationCycle:  No valid RTID for rentable R%08d on date: %s", rid, dt.Format(RRDATEINPFMT))
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
//  rta = array of RentableMarketRate structures that covers all rental rates during the period d1 - d2.
//        This array is the MR attribute in the RentableMarketRate struct
//  rsa = array of rentable specialties that apply to the rentable we're calculating
//========================================================================================================
func CalculateGSR(ctx context.Context, d1, d2 time.Time, rid int64, rta *[]RentableTypeRef, rsa []RentableSpecialty, xbiz *XBusiness) (float64, error) {
	var total = float64(0) // init total
	// Console("Entered CalculateGSR: d1 = %s, d2 = %s, rid = %d\n", d1.Format(RRDATEFMTSQL), d2.Format(RRDATEFMTSQL), rid)

	// Get the first date that overlaps the rta values
	dt := d1
	for dt.Before(d2) {
		found := false
		for i := 0; i < len(*rta); i++ {
			if DateInRange(&dt, &(*rta)[i].DtStart, &(*rta)[i].DtStop) {
				found = true
				break
			}
		}
		if found {
			break
		}
		dt = dt.AddDate(0, 0, 1)
	}

	rentCycle, _, gsrpc, err := GetProrationCycle(ctx, &dt, rid, rta, xbiz)
	if err != nil {
		Ulog("CalculateGSR: GetProrationCycle returned error: %s\n", err.Error())
		return float64(0), err
	}
	if rentCycle < 0 || gsrpc < 0 {
		Ulog("CalculateGSR: warning: one or more cycle values is unset\n")
	}
	// prorateDur := CycleDuration(prorateCycle, d1) // the proration cycle expressed as a duration
	inc := CycleDuration(gsrpc, dt)              // increment durations for rent calculation
	rentCycleDur := CycleDuration(rentCycle, dt) // this is the rentcycle expressed as a duration
	rtr, err := SelectRentableTypeRefForDate(ctx, rta, &dt)
	if err != nil {
		return float64(0), err
	}

	// Console("CalculateGSR: rentCycle = %d (%v), gsrpc = %d (%v)\n", rentCycle, rentCycleDur, gsrpc, inc)
	// Console("rtr = (%s - %s) rtid = %d\n", rtr.DtStart.Format("1/2/06"), rtr.DtStop.Format("1/2/06"), rtr.RTID)

	for d := dt; d.Before(d2); d = d.Add(inc) { // spin through the period in the defined increments
		rate := FindApplicableMarketRate(d, d1, d2, xbiz.RT[rtr.RTID].MR) // find the rate applicable for this increment
		rent := float64(inc) * rate / float64(rentCycleDur)               // how much for the period: inc
		total += rent                                                     // increment the total by this amount
		for i := 0; i < len(rsa); i++ {
			total += rsa[i].Fee * float64(inc) / float64(rentCycleDur)
		}
	}
	return total, err
}

// CalculateNumberOfCycles calculates the number of rent cycles for the supplied period d1-d2 and cycle time.
// It always returns an integral number of cycles -- even if the date range implies some
// fractional level.
//
//
// Params:
//   d1 = start datetime of the period
//   d2 = stop datetime of the period
//   c  = cycle time {}
//========================================================================================================
func CalculateNumberOfCycles(d1, d2 *time.Time, c int64) int64 {
	dur := d2.Sub(*d1)                // duration of the period
	cycleDur := CycleDuration(c, *d1) // duration of the cycle-time
	// Console("period = %s - %s,  dur = %d, cycleDur = %d\n", d1.Format(RRDATEFMT3), d2.Format(RRDATEFMT3), dur, cycleDur)
	n := int64(float64(dur)/float64(cycleDur) + 0.5) // number of cycles
	return n
}

// GetRentCycleRefList returns an array of RentCycleRefs indicating the RentCycle & ProrationCycle
// for the period d1-d2.  Most often there is a single entyr in the array. However, if the
// rent cycle changed during d1-d2, the returned array will have entries for any changes.
// Params:
//   r    = the rentable of interest
//   d1   = start datetime of the period
//   d2   = stop datetime of the period
//   xbiz = biz struct containing all the RentableTypes for the business
//========================================================================================================
func GetRentCycleRefList(ctx context.Context, r *Rentable, d1, d2 *time.Time, xbiz *XBusiness) ([]RentCycleRef, error) {
	var (
		m   []RentCycleRef
		err error
	)

	r.RT, err = GetRentableTypeRefsByRange(ctx, r.RID, d1, d2)
	if err != nil {
		return m, err
	}

	for i := 0; i < len(r.RT); i++ {
		var rcr RentCycleRef
		rcr.DtStart = r.RT[i].DtStart
		rcr.DtStop = r.RT[i].DtStop
		rcr.RentCycle = r.RT[i].OverrideRentCycle
		rcr.ProrationCycle = r.RT[i].OverrideProrationCycle
		if rcr.RentCycle == 0 {
			rcr.RentCycle = xbiz.RT[r.RT[i].RTID].RentCycle
			rcr.ProrationCycle = xbiz.RT[r.RT[i].RTID].Proration
		}
		m = append(m, rcr)
	}
	return m, err
}

// GetRentableCycles calculates the number of rent cycles for the supplied rentable given the
// period d1-d2. It always returns an integral number of cycles -- even if the date range implies some
// fractional level ==> it may return 0
// Params:
//   d1 = start datetime of the period
//   d2 = stop datetime of the period
//   r  = The rentable we're interested in
//========================================================================================================
func GetRentableCycles(ctx context.Context, r *Rentable, d1, d2 *time.Time, xbiz *XBusiness) (int64, error) {

	var n int64

	rcl, err := GetRentCycleRefList(ctx, r, d1, d2, xbiz)
	if err != nil {
		return n, err
	}

	for i := 0; i < len(rcl); i++ {
		// note that DtStart may be long befor d1 and DtStop period may be set to "infinity". We need to clip it at d1,d2
		start := rcl[i].DtStart
		if d1.After(start) {
			start = *d1
		}
		stop := rcl[i].DtStop
		if d2.Before(stop) {
			stop = *d2
		}
		n += CalculateNumberOfCycles(&start, &stop, rcl[i].RentCycle)
	}
	return n, err
}

// GSRdata is a struct with a datetime and a GSR amount that describes the GSR for that timeslice.
// An array of these structs returned with CalculateLoadedGSR fully describes how the GSR is calculated.
// This method is necessary to account for changes in GSR during a time period.
type GSRdata struct {
	Dt     time.Time // datetime - in increments of GSRPC durations
	Amount float64   // amount for GSR during this period
}

// CalculateLoadedGSR calculates the gross scheduled rent including any Specialties associated with the rentable.
// Params:
//   d1 = start datetime of the period
//   d2 = stop datetime of the period
//   rt = array of RentableMarketRate structures that covers all rental rates during the period d1 - d2.
//        This array is the MR attribute in the RentableMarketRate struct
// Returns:
//   float64 - loaded GSR for d1 to d2
//   []GSRdata - an array of GSRdata structs, with the date/time and gsr Amount in increments of GSRPC from d1 to d2
//   time.Duration - the GSRPC for this rentable
//   error - any error returned by the routines looking for data values
//========================================================================================================
func CalculateLoadedGSR(ctx context.Context, rBID, rRID int64, d1, d2 *time.Time, xbiz *XBusiness) (float64, []GSRdata, time.Duration, error) {
	funcname := "CalculateLoadedGSR"
	var period = time.Duration(0)
	var m []GSRdata
	var err error
	gsr := float64(0) // total rent, to update on each pass through the loop below

	// Console("Entered %s: rRID = %d, d1 = %s, d2 = %s\n", funcname, rRID, d1.Format(RRDATEINPFMT), d2.Format(RRDATEINPFMT))

	rta, err := GetRentableTypeRefsByRange(ctx, rRID, d1, d2) // get the list
	if err != nil {
		return gsr, m, period, err
	}

	if len(rta) == 0 {
		err = fmt.Errorf("%s:  No valid RTID for rentable R%08d during period %s to %s",
			funcname, rRID, d1.Format(RRDATEINPFMT), d2.Format(RRDATEINPFMT))
		Ulog("%s", err.Error())
		return gsr, m, period, err // this is bad! No RTID for the supplied time range
	}
	dtFirst := *d1
	for dtFirst.Before(*d2) {
		found := false
		for i := 0; i < len(rta); i++ {
			// Console("rta[%d] - %s - %s\n", i, rta[i].DtStart.Format(RRDATETIMEINPFMT), rta[i].DtStop.Format(RRDATETIMEINPFMT))
			if DateInRange(&dtFirst, &rta[i].DtStart, &rta[i].DtStop) {
				// Console("FOUND the match at %d\n", i)
				found = true
				break
			}
		}
		if found {
			break
		}
		dtFirst = dtFirst.AddDate(0, 0, 1)
	}

	// Console("%s, dtFirst = %s\n", funcname, dtFirst.Format(RRDATETIMEINPFMT))

	// find the Gross Scheduled Rent Proration Cycle - GSRPC - the intervals over which the GSR is calculated
	_, _, gsrpc, err := GetProrationCycle(ctx, &dtFirst, rRID, &rta, xbiz)
	if err != nil {

		// Console("%s: error from GetProrationCycle: %s\n", funcname, err.Error())

		return gsr, m, period, err
	}

	// Console("%s: gsrpc = %v, dtFirst = %s\n", funcname, gsrpc, dtFirst.Format(RRDATEFMT4))
	if gsrpc == int64(0) {
		err = fmt.Errorf("%s: GSRPC == 0 for BID=%d, RID=%d, d1 = %s, d2 = %s", funcname, rBID, rRID, d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4))
		Ulog(err.Error())
		// Console(err.Error())
		return float64(0), m, period, nil
	}

	period = CycleDuration(gsrpc, dtFirst)           // increment of time we'll use to determine gsr in increments between dtFirst & d2
	dtNext := dtFirst                                // initialize so that the variable is known
	for dt := dtFirst; dt.Before(*d2); dt = dtNext { // spin through time period d1 - d2 in increments of gsrpc and add up the GSR
		dtNext = dt.Add(period) // establish the end of the period.  We'll add up the gsr for period dt to dtNext.
		//--------------------------------------------------------------------
		// Get the RentableSpecialties applicable for this increment...
		//--------------------------------------------------------------------
		rsa, nerr := GetRentableSpecialtyTypesForRentableByRange(ctx, rBID, rRID, &dt, &dtNext) // this gets an array of rentable specialties that overlap this time period
		if nerr != nil {
			err = fmt.Errorf("%s:  error getting specialties for rentable R%08d during period %s to %s.  err = %s",
				funcname, rRID, dt.Format(RRDATEINPFMT), dtNext.Format(RRDATEINPFMT), nerr.Error())
			Ulog("%s", err.Error())
			break
		}
		//------------------------------------------------------------------
		// Finally, calculate the GSR for this increment...
		//------------------------------------------------------------------
		rentThisPeriod, err := CalculateGSR(ctx, dt, dtNext, rRID, &rta, rsa, xbiz)
		if err != nil {
			return gsr, m, period, err
		}

		var g = GSRdata{Dt: dt, Amount: rentThisPeriod}
		m = append(m, g)
		gsr += rentThisPeriod
		// Console("%s: rentThisPeriod = %.2f,  cumulative total: %.2f\n", dt.Format(RRDATEFMTSQL), rentThisPeriod, gsr)
	}
	gsr = RoundToCent(gsr) // do this to ensure that we avoid the off-by-a-penny errors
	return gsr, m, period, err
}
