package rlib

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// BizPropsFee struct holds detailed info about any fee
// configured in business properties
type BizPropsFee struct {
	BID              int64
	ARID             int64
	ARName           string
	Amount           float64
	ARFLAGS          uint64
	ARRentCycle      int64
	ARProrationCycle int64
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

// GetBizPropPetFees returns pet fees with detailed data
// defined in BizPropsFee
func GetBizPropPetFees(ctx context.Context, BID int64, bizPropName string) (fees []BizPropsFee, err error) {
	const funcname = "GetBizPropPetFees"
	var (
		bizPropJSON BizProps
	)
	fmt.Printf("Entered in %s\n", funcname)

	// initialize pet fees
	fees = []BizPropsFee{}

	// get business properties data
	bizPropJSON, err = GetDataFromBusinessPropertyName(ctx, bizPropName, BID)
	if err != nil {
		return
	}

	// get pet Fees
	for _, n := range bizPropJSON.PetFees {
		var pf BizPropsFee
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
		pf.ARFLAGS = ar.FLAGS
		pf.ARRentCycle = ar.DefaultRentCycle
		pf.ARProrationCycle = ar.DefaultProrationCycle

		// append in the list
		fees = append(fees, pf)
	}

	// return finally
	return
}

// GetBizPropVehicleFees returns vehicle fees with detailed data
// defined in BizPropsFee
func GetBizPropVehicleFees(ctx context.Context, BID int64, bizPropName string) (fees []BizPropsFee, err error) {
	const funcname = "GetBizPropVehicleFees"
	var (
		bizPropJSON BizProps
	)
	fmt.Printf("Entered in %s\n", funcname)

	// initialize vehicle fees
	fees = []BizPropsFee{}

	// get business properties data
	bizPropJSON, err = GetDataFromBusinessPropertyName(ctx, bizPropName, BID)
	if err != nil {
		return
	}

	// get pet Fees
	for _, n := range bizPropJSON.VehicleFees {
		var vf BizPropsFee
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
		vf.ARFLAGS = ar.FLAGS
		vf.ARRentCycle = ar.DefaultRentCycle
		vf.ARProrationCycle = ar.DefaultProrationCycle

		// append in the list
		fees = append(fees, vf)
	}

	// return finally
	return
}

// GetEpochListByBizPropName returns epochs configured for a business
func GetEpochListByBizPropName(ctx context.Context, BID int64, bizPropName string) (epochs BizPropsEpochs, err error) {
	const funcname = "GetEpochListByBizPropName"
	var (
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

// GetEpochByBizPropName returns epochs configured for a business
// GetEpochFromBaseDate returns the epoch date based on cycle,
// start date and pre-configured base epochs
//
// The required unit(s) should be extracted from pre-configured base epochs
// to calculate proper Epoch date based on cycle, start date
// For ex.,
//     1. If cycle is MINUTELY then the unit(s) to consider from pre-configured
//        base epoch(minutely) are Second, NenoSecond
//     2. If cycle if MONTHLY then unit(s) to consider from pre-configured
//        base epoch(monthly) are Day, Hour, Minute, Second, NenoSecond
//
// Time Location: It always keeps time location from base date(b)
//
// INPUTS
//             ctx  = context.Context instance
//             BID  = Business ID
//     bizPropName  = business property name
//              d1  = start date
//              d2  = stop date
//           cycle  = integer presentable number
//
// RETURNS
//     ok    - If epoch is possible in given date range then true else false
//             * In case of false, epoch still has calculated value, to be
//               happened on next cycle. It helps to determine for calle routine
//               what should it does on epoch whether it falls in range or not.
//     epoch - absolute epoch for given date range
//     err   - any error encounterd
//-----------------------------------------------------------------------------
func GetEpochByBizPropName(ctx context.Context, BID int64, bizPropName string, d1, d2 time.Time, cycle int64) (ok bool, epoch time.Time, err error) {
	const funcname = "GetEpochByBizPropName"
	fmt.Printf("Entered in %s\n", funcname)

	// initialize epochs
	var epochs BizPropsEpochs
	epochs, err = GetEpochListByBizPropName(ctx, BID, bizPropName)
	if err != nil {
		return
	}

	// TAKE BASE DATE FROM THE BIZPROP EPOCHS
	var baseEpoch time.Time
	switch cycle {
	case RECURDAILY:
		baseEpoch = epochs.Daily
	case RECURWEEKLY:
		baseEpoch = epochs.Weekly
	case RECURMONTHLY:
		baseEpoch = epochs.Monthly
	case RECURQUARTERLY:
		baseEpoch = epochs.Quarterly
	case RECURYEARLY:
		baseEpoch = epochs.Yearly
	}

	// GET AN ABSOLUTE EPOCH FROM BASE DATE
	ok, epoch = GetEpochFromBaseDate(baseEpoch, d1, d2, cycle)

	return
}
