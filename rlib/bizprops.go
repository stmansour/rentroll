package rlib

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// BizPropsVehicleFee struct holds detailed info about vehicle fee
// configured in business properties
type BizPropsVehicleFee struct {
	BID    int64
	ARID   int64
	ARName string
	Amount float64
}

// BizPropsPetFee struct holds detailed info about pet fee
// configured in business properties
type BizPropsPetFee struct {
	BID    int64
	ARID   int64
	ARName string
	Amount float64
}

// BizPropsEpochs defines the default trigger due times for recurring
// assessments.
type BizPropsEpochs struct {
	Daily     time.Time // default hour:minute:second rent becomes due on daily rentals
	Weekly    time.Time // default dayOfTheWeek:hour:minute:second rent becomes due on weekly rentals
	Monthly   time.Time // default dayOfTheMonth:hour:minute:second rent becomes due on Monthly rentals
	Quarterly time.Time // default month:dayOfMonth:hour:minute:second rent becomes due on Quarterly rentals
	Yearly    time.Time // default month:dayOfMonth:hour:minute:second rent becomes due on Quarterly rentals
}

// GetDataFromBusinessPropertyName returns instance of BizProps with
// JSON parsing from business properties data for requested name
func GetDataFromBusinessPropertyName(ctx context.Context, name string, BID int64) (bizPropJSON BizProps, err error) {

	// initialize
	bizPropJSON = BizProps{}

	// get business properties
	var bizProp BusinessProperties
	bizProp, err = GetBusinessPropertiesByName(ctx, name, BID)
	if err != nil {
		return
	}

	// get json doc from Data column
	err = json.Unmarshal(bizProp.Data, &bizPropJSON)
	return
}

// GetPetFeesFromGeneralBizProps returns pet fees with detailed data
// defined in BizPropsPetFee
func GetPetFeesFromGeneralBizProps(ctx context.Context, BID int64) (fees []BizPropsPetFee, err error) {
	const funcname = "GetPetFeesFromGeneralBizProps"
	var (
		bizPropName = "general"
		bizPropJSON BizProps
	)
	fmt.Printf("Entered in %s\n", funcname)

	// initialize pet fees
	fees = []BizPropsPetFee{}

	// get business properties data
	bizPropJSON, err = GetDataFromBusinessPropertyName(ctx, bizPropName, BID)
	if err != nil {
		return
	}

	// get pet Fees
	for _, n := range bizPropJSON.PetFees {
		var pf BizPropsPetFee
		var ar AR
		ar, err = GetARByName(ctx, BID, n)
		if err != nil {
			return
		}

		// migrate values from ar to pf
		pf.BID = ar.BID
		pf.ARID = ar.ARID
		pf.ARName = ar.Name
		pf.Amount = ar.DefaultAmount

		// append in the list
		fees = append(fees, pf)
	}

	// return finally
	return
}

// GetVehicleFeesFromGeneralBizProps returns vehicle fees with detailed data
// defined in BizPropsVehicleFee
func GetVehicleFeesFromGeneralBizProps(ctx context.Context, BID int64) (fees []BizPropsVehicleFee, err error) {
	const funcname = "GetVehicleFeesFromGeneralBizProps"
	var (
		bizPropName = "general"
		bizPropJSON BizProps
	)
	fmt.Printf("Entered in %s\n", funcname)

	// initialize vehicle fees
	fees = []BizPropsVehicleFee{}

	// get business properties data
	bizPropJSON, err = GetDataFromBusinessPropertyName(ctx, bizPropName, BID)
	if err != nil {
		return
	}

	// get pet Fees
	for _, n := range bizPropJSON.VehicleFees {
		var vf BizPropsVehicleFee
		var ar AR
		ar, err = GetARByName(ctx, BID, n)
		if err != nil {
			return
		}

		// migrate values from ar to vf
		vf.BID = ar.BID
		vf.ARID = ar.ARID
		vf.ARName = ar.Name
		vf.Amount = ar.DefaultAmount

		// append in the list
		fees = append(fees, vf)
	}

	// return finally
	return
}

// GetEpochsFromGeneralBizProps returns epochs configured for a business
func GetEpochsFromGeneralBizProps(ctx context.Context, BID int64) (epochs BizPropsEpochs, err error) {
	const funcname = "GetVehicleFeesFromGeneralBizProps"
	var (
		bizPropName = "general"
		bizPropJSON BizProps
	)
	fmt.Printf("Entered in %s\n", funcname)

	// initialize epochs
	epochs = BizPropsEpochs{}

	// get business properties data
	bizPropJSON, err = GetDataFromBusinessPropertyName(ctx, bizPropName, BID)
	if err != nil {
		return
	}
	epochs = bizPropJSON.Epochs

	return
}

// GetEpochForCycleInDateRange returns the epoch date based on cycle,
// start date, end date, and pre-configured base date
//
// The required unit(s) should be extracted from pre-configured base date
// to calculate proper Epoch date based on cycle, start/end dates
// For ex.,
//     1. If cycle is SECONDLY then the unit(s) to consider from pre-configured
//        base date are Second, NenoSecond
//     2. If cycle if MONTHLY then unit(s) to consider from pre-configured base date
//        are Month, Day, Hour, Minute, Second, NenoSecond
//
// It always keeps time location from start date(d1)
//
// INPUTS
//               b  = preconfigured base date
//              d1  = start date
//              d2  = stop date
//           cycle  = integer presentable number
//
// RETURNS
//     epoch - proper epoch for given date range
//     any error encountered
//-----------------------------------------------------------------------------
func GetEpochForCycleInDateRange(b, d1, d2 time.Time, cycle int64) (epoch time.Time, err error) {
	if d2.Before(d1) {
		err = fmt.Errorf("end date: %q should not be prior to Start date: %q", d2, d1)
		return
	}

	// TODO(Sudip): What if base date and d1 falls at same unit value

	// KEEP ASSIGN START DATE TIME LOCATION IN "EPOCH"
	loc := d1.Location()

	// DECIDE BASED ON CYCLE
	switch cycle {
	case RECURNONE:
		epoch = d1
	case RECURSECONDLY:
		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), d1.Minute(), b.Second(), b.Nanosecond(), loc)
	case RECURMINUTELY:
		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), b.Minute(), b.Second(), b.Nanosecond(), loc)
	case RECURHOURLY:
		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), loc)
	case RECURDAILY:
		epoch = time.Date(d1.Year(), d1.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), loc)
	case RECURWEEKLY:
		epoch = time.Date(d1.Year(), d1.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), loc)
		epoch = epoch.AddDate(0, 0, 7) // ADD MORE 7 DAYS
	case RECURMONTHLY:
		epoch = time.Date(d1.Year(), b.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), loc)
	case RECURQUARTERLY:
		epoch = time.Date(d1.Year(), b.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), loc)
		epoch = epoch.AddDate(0, 3, 0) // ADD MORE 3 MONTHS
	case RECURYEARLY:
		epoch = time.Date(b.Year(), b.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), loc)
	default:
		err = fmt.Errorf("cycle:%d, is out of range", cycle)
		return
	}

	// IF "EPOCH" FALLS AFTER END DATE
	if epoch.After(d2) {
		err = fmt.Errorf("epoch:%q falls after end date: %q", epoch, d2)
		return
	}

	return
}
