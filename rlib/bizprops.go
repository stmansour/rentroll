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
	ARRentCycle      int64
	ARProrationCycle int64
	Start            JSONDate
	Stop             JSONDate
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
	// Console("Entered GetDataFromBusinessPropertyName. Getting biz prop named %s\n", name)
	bizPropJSON = BizProps{}
	var bizProp BusinessProperties
	bizProp, err = GetBusinessPropertiesByName(ctx, name, BID)
	if err != nil {
		// Console("GetDataFromBusinessPropertyName: error getting property named %s: %s\n", name, err.Error())
		return
	}
	// Console("GetDataFromBusinessPropertyName: got bizProp, len(bizProp.Data) = %d\n", len(bizProp.Data))
	if len(bizProp.Data) > 0 {
		err = json.Unmarshal(bizProp.Data, &bizPropJSON)
	}
	return
}

// GetBizPropPetFees returns pet fees with detailed data
// defined in BizPropsFee
func GetBizPropPetFees(ctx context.Context, BID int64, bizPropName string, rStart, rStop time.Time) (fees []BizPropsFee, err error) {
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
		var ar AR
		ar, err = GetARByName(ctx, BID, n)
		if err != nil {
			return
		}

		// migrate values from ar to pf
		pf := BizPropsFee{
			BID:              ar.BID,
			ARID:             ar.ARID,
			ARName:           ar.Name,
			Amount:           ar.DefaultAmount,
			ARRentCycle:      ar.DefaultRentCycle,
			ARProrationCycle: ar.DefaultProrationCycle,
		}

		oneTimeCharge := (ar.FLAGS & (1 << ARIsNonRecurCharge)) != 0
		if oneTimeCharge {
			pf.Start = JSONDate(rStart)
			pf.Stop = JSONDate(rStart)
		} else {
			pf.Start = JSONDate(rStart)
			pf.Stop = JSONDate(rStop)
		}

		// append in the list
		fees = append(fees, pf)
	}

	// return finally
	return
}

// GetBizPropVehicleFees returns vehicle fees with detailed data
// defined in BizPropsFee
func GetBizPropVehicleFees(ctx context.Context, BID int64, bizPropName string, rStart, rStop time.Time) (fees []BizPropsFee, err error) {
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
		var ar AR
		ar, err = GetARByName(ctx, BID, n)
		if err != nil {
			return
		}

		// migrate values from ar to vf
		vf := BizPropsFee{
			BID:              ar.BID,
			ARID:             ar.ARID,
			ARName:           ar.Name,
			Amount:           ar.DefaultAmount,
			ARRentCycle:      ar.DefaultRentCycle,
			ARProrationCycle: ar.DefaultProrationCycle,
		}

		oneTimeCharge := (ar.FLAGS & (1 << ARIsNonRecurCharge)) != 0
		if oneTimeCharge {
			vf.Start = JSONDate(rStart)
			vf.Stop = JSONDate(rStart)
		} else {
			vf.Start = JSONDate(rStart)
			vf.Stop = JSONDate(rStop)
		}

		// append in the list
		fees = append(fees, vf)
	}

	// return finally
	return
}

// GetEpochListByBizPropName returns epochs configured for a business
func GetEpochListByBizPropName(ctx context.Context, BID int64, bizPropName string) (epochs BizPropsEpochs, err error) {
	const funcname = "GetEpochListByBizPropName"
	var bizPropJSON BizProps
	// Console("Entered %s\n", funcname)

	// initialize epochs
	epochs = BizPropsEpochs{}

	// get business properties data
	bizPropJSON, err = GetDataFromBusinessPropertyName(ctx, bizPropName, BID)
	if err != nil {
		return
	}
	// Console("Setting biz props\n")
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
	fmt.Printf("Entered %s\n", funcname)

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
