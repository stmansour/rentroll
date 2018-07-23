package rlib

import (
	"context"
	"fmt"
	"time"
)

// GetRAFlowFeeCycles returns rent and proration cycle
// for a given ARID, RID.
//
// It will look for accoun rules default cycles first.
// If RID, rentStart, rentStop are provided then
// it will consider rentableType's cycle
//
// INPUTS
//             ctx  = db transaction context
//             BID  = Business ID
//            ARID  = account rule ID
//             RID  = Rentable ID
//       rentStart  = Rent Start Date
//
// RETURNS
//     RentCycle numeric value
//     ProrationCycle numeric value
//     any error encountered
//-----------------------------------------------------------------------------
func GetRAFlowFeeCycles(ctx context.Context, ARID, RID int64, rentStart time.Time) (RentCycle, ProrationCycle int64, err error) {

	// GET ACCOUNT RULE FIRST
	var ar AR
	ar, err = GetAR(ctx, ARID)
	if err != nil {
		return
	}

	// TAKE THOSE FROM AR BY DEFAULT
	RentCycle = ar.DefaultRentCycle
	ProrationCycle = ar.DefaultProrationCycle

	// IF NO RID THEN JUST RETURN
	if RID == 0 {
		return
	}

	// IF RENTABLE IS AVAILABLE THEN
	var rrt RentableTypeRef
	rrt, err = GetRentableTypeRefForDate(ctx, RID, &rentStart)
	if err != nil {
		return
	}

	// IF NO RENTABLE TYPE REF THEN JUST RETURN
	if rrt.RID == 0 {
		return
	}

	// GET RENTABLE TYPE
	var rt RentableType
	err = GetRentableType(ctx, rrt.RTID, &rt)
	if err != nil {
		return
	}

	// TAKE CYCLES FROM RENTABLE TYPE, IF RTID > 0
	if rt.RTID > 0 {
		if rrt.OverrideRentCycle > RECURNONE { // if there's an override for RentCycle...
			RentCycle = rrt.OverrideRentCycle // ...set it
		} else {
			RentCycle = rt.RentCycle
		}
		if rrt.OverrideProrationCycle > RECURNONE { // if there's an override for Propration...
			ProrationCycle = rrt.OverrideProrationCycle // ...set it
		} else {
			ProrationCycle = rt.Proration
		}
	}

	return
}

// GetRAFlowInitialPetFees returns the list of initial fees based on
// RentStart/Stop been set in raflow data for a pet
//
// INPUTS
//             ctx  = db transaction context
//             BID  = Business ID
//             RID  = Rentable ID
//          rStart  = rent start date
//           rStop  = rent stop date
//            meta  = RAFlowMetaInfo data
//
// RETURNS
//     RAVehiclesFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func GetRAFlowInitialPetFees(ctx context.Context, BID, RID int64, rStart, rStop JSONDate, meta *RAFlowMetaInfo) (fees []RAFeesData, err error) {
	const funcname = "GetRAFlowInitialPetFees"
	var (
		bizPropName = "general"
		d1          = (time.Time)(rStart)
		d2          = (time.Time)(rStop)
	)
	fmt.Printf("Entered in %s\n", funcname)

	// INITIALIZE FEES
	fees = []RAFeesData{}

	// GET PET FEES FROM BUSINESS PROPERTIES
	var petFees []BizPropsPetFee
	petFees, err = GetBizPropPetFees(ctx, BID, bizPropName)
	if err != nil {
		return
	}

	// FOR EACH FEE CONFIGURED IN BIZPROP
	for _, petFee := range petFees {

		// GET RENT, PRORATION CYCLE
		var RentCycle, ProrationCycle int64
		RentCycle, ProrationCycle, err = GetRAFlowFeeCycles(ctx, petFee.ARID, RID, d1)
		if err != nil {
			return
		}

		// GET EPOCH
		var epoch time.Time
		_, epoch, err = GetEpochByBizPropName(ctx, BID, bizPropName, d1, d2, RentCycle)
		if err != nil {
			tot, np, tp := SimpleProrateAmount(petFee.Amount, RentCycle, ProrationCycle, &d1, &d2, &epoch)

			meta.LastTMPASMID++

			// ADD FEE IN LIST
			raFee := RAFeesData{
				TMPASMID:        meta.LastTMPASMID,
				ASMID:           0,
				ARID:            petFee.ARID,
				ARName:          petFee.ARName,
				ContractAmount:  tot,
				RentCycle:       RentCycle,
				Start:           rStart,
				Stop:            rStop,
				AtSigningPreTax: 0.00,
				SalesTax:        0.00,
				TransOccTax:     0.00,
				Comment:         fmt.Sprintf("prorated for %d of %d %s", np, tp, ProrationUnits(ProrationCycle)),
			}
			fees = append(fees, raFee)
		}
	}

	return
}
