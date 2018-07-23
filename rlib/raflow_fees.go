package rlib

import (
	"context"
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

/*
// GetRAFlowInitialPetFees returns the list of initial fees based on
// RentStart/Stop been set in raflow data for a pet
//
// INPUTS
//             ctx  = db transaction context
//             BID  = Business ID
//          rStart  = rent start date
//           rStop  = rent stop date
//            meta  = RAFlowMetaInfo data
//
// RETURNS
//     RAVehiclesFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func GetRAFlowInitialPetFees(ctx context.Context, BID int64, rStart, rStop JSONDate, meta *RAFlowMetaInfo) (fees []RAFeesData, err error) {
	const funcname = "GetRAFlowInitialPetFees"
	fmt.Printf("Entered in %s\n", funcname)

	// INITIALIZE FEES
	fees = []RAFeesData{}

	// GET BUSINESS EPOCHS FOR "GENERAL" BUSINESS PROPERTIES
	var epochs BizPropsEpochs
	epochs, err = GetEpochsFromGeneralBizProps(ctx, BID)
	if err != nil {
		return
	}

	// CONVERT JSON DATE TO TIME
	d1, d2 := (time.Time)(rStart), (time.Time)(rStop)

	// GET MONTHLY EPOCH
	epochMonthly := time.Date(d1.Year(), d1.Month(), epochs.Monthly.Day(),
		d1.Hour(), d1.Minute(), d1.Second(), d1.Nanosecond(), d1.Location())

	return
}*/
