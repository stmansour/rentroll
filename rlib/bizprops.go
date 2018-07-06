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

// GetPetFeesFromGeneralBizProps returns pet fees with detailed data
// defined in BizPropsPetFee
func GetPetFeesFromGeneralBizProps(ctx context.Context, BID int64) (fees []BizPropsPetFee, err error) {
	const funcname = "GetPetFeesFromGeneralBizProps"
	var (
		bizPropName = "general"
		bizPropJSON = BizProps{}
	)
	fees = []BizPropsPetFee{}
	fmt.Printf("Entered in %s\n", funcname)

	// get business properties
	var bizProp BusinessProperties
	bizProp, err = GetBusinessPropertiesByName(ctx, bizPropName, BID)
	if err != nil {
		return
	}

	// get json doc from Data column
	if err = json.Unmarshal(bizProp.Data, &bizPropJSON); err != nil {
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
		bizPropJSON = BizProps{}
	)
	fees = []BizPropsVehicleFee{}
	fmt.Printf("Entered in %s\n", funcname)

	// get business properties
	var bizProp BusinessProperties
	bizProp, err = GetBusinessPropertiesByName(ctx, bizPropName, BID)
	if err != nil {
		return
	}

	// get json doc from Data column
	if err = json.Unmarshal(bizProp.Data, &bizPropJSON); err != nil {
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
