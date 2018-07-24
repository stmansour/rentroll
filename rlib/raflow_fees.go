package rlib

import (
	"context"
	"fmt"
	"time"
)

// GetRAFlowFeeCycles returns rent and proration cycle
// for a given RID (if rentable available).
//
// It will look for accoun rules default cycles first.
// If RID, rentStart, rentStop are provided then
// it will consider rentableType's cycle
//
// INPUTS
//             ctx  = db transaction context
//             BID  = Business ID
//             RID  = Rentable ID
//       rentStart  = Rent Start Date
//
// RETURNS
//     RentCycle numeric value
//     ProrationCycle numeric value
//     any error encountered
//-----------------------------------------------------------------------------
func GetRAFlowFeeCycles(ctx context.Context, RID int64, rentStart time.Time, RentCycle, ProrationCycle *int64) (err error) {

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
			*RentCycle = rrt.OverrideRentCycle // ...set it
		} else {
			*RentCycle = rt.RentCycle
		}
		if rrt.OverrideProrationCycle > RECURNONE { // if there's an override for Propration...
			*ProrationCycle = rrt.OverrideProrationCycle // ...set it
		} else {
			*ProrationCycle = rt.Proration
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
//     list of pet fees
//     any error encountered
//-----------------------------------------------------------------------------
func GetRAFlowInitialPetFees(
	ctx context.Context,
	BID, RID int64,
	rStart, rStop JSONDate,
	meta *RAFlowMetaInfo,
) (fees []RAFeesData, err error) {

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
	var petFees []BizPropsFee
	petFees, err = GetBizPropPetFees(ctx, BID, bizPropName)
	if err != nil {
		return
	}

	// FOR EACH FEE CONFIGURED IN BIZPROP
	for _, petFee := range petFees {

		// GET RENT, PRORATION CYCLE
		RentCycle := petFee.ARRentCycle
		ProrationCycle := petFee.ARProrationCycle
		err = GetRAFlowFeeCycles(ctx, RID, d1, &RentCycle, &ProrationCycle)
		if err != nil {
			return
		}

		// ========================================================================
		// GET EPOCH BASED ON RENTCYCLE FOR THIS PET FEE
		// ========================================================================
		// TODO(Sudip & Steve): WHEN WE INTEGRATE EPOCHS IN RENTABLE TYPES       //
		//                      WE SHOULD TAKE EPOCHS FIRST FROM IT THEN         //
		//                      FROM BIZPROPS IN CASE NOT FOUND IN RENTALBE TYPE //
		// ========================================================================
		var epoch time.Time
		_, epoch, err = GetEpochByBizPropName(ctx, BID, bizPropName, d1, d2, RentCycle)
		if err != nil {
			return
		}

		//--------------------------------------------------------------
		// Here are the AR.FLAGS bits:
		//
		//     1<<4 -  Is Rent Assessment
		//     1<<6 -  Is non-recur charge
		//     1<<7 -  PETID required
		//--------------------------------------------------------------
		// IF IT IS NON-RECUR CHARGE THEN
		oneTimeCharge := (petFee.ARFLAGS & (1 << 6)) != 0
		if oneTimeCharge {
			// ADD FEE IN LIST
			meta.LastTMPASMID++
			raFee := RAFeesData{
				TMPASMID:        meta.LastTMPASMID,
				ASMID:           0,
				ARID:            petFee.ARID,
				ARName:          petFee.ARName,
				ContractAmount:  petFee.Amount,
				RentCycle:       RECURNONE,
				Start:           rStart,
				Stop:            rStart,
				AtSigningPreTax: 0.00,
				SalesTax:        0.00,
				TransOccTax:     0.00,
				Comment:         "",
			}
			fees = append(fees, raFee)

		} else if petFee.ARFLAGS&(1<<4) != 0 { // IT MUST BE RENT ASM ONE

			// CHECK FOR PRORATED AMOUNT REQUIRED
			needProratedRent := d1.Day() != epoch.Day()

			// START DAY IS NOT SAME AS EPOCH THEN CALCULATE PRORATED AMOUNT
			if needProratedRent {
				td2 := time.Date(d1.Year(), d1.Month(), epoch.Day(), d1.Hour(), d1.Minute(), d1.Second(), d1.Nanosecond(), d1.Location())
				td2 = NextPeriod(&td2, RentCycle)

				tot, np, tp := SimpleProrateAmount(petFee.Amount, RentCycle, ProrationCycle, &d1, &td2, &epoch)
				cmt := ""
				if tot < petFee.Amount {
					cmt = fmt.Sprintf("prorated for %d of %d %s", np, tp, ProrationUnits(ProrationCycle))
				}

				// ADD FEE IN LIST
				meta.LastTMPASMID++
				raFee := RAFeesData{
					TMPASMID:        meta.LastTMPASMID,
					ASMID:           0,
					ARID:            petFee.ARID,
					ARName:          petFee.ARName,
					ContractAmount:  tot,
					RentCycle:       RentCycle,
					Start:           rStart,
					Stop:            rStart,
					AtSigningPreTax: 0.00,
					SalesTax:        0.00,
					TransOccTax:     0.00,
					Comment:         cmt,
				}
				fees = append(fees, raFee)
			}

			// CALCULATE RECURRING ONE FROM EPOCH DATE
			// ADD FEE IN LIST
			meta.LastTMPASMID++
			raFee := RAFeesData{
				TMPASMID:        meta.LastTMPASMID,
				ASMID:           0,
				ARID:            petFee.ARID,
				ARName:          petFee.ARName,
				ContractAmount:  petFee.Amount,
				RentCycle:       RentCycle,
				Start:           JSONDate(epoch),
				Stop:            rStop,
				AtSigningPreTax: 0.00,
				SalesTax:        0.00,
				TransOccTax:     0.00,
				Comment:         "",
			}
			fees = append(fees, raFee)
		}
	}

	return
}

// GetRAFlowInitialVehicleFees returns the list of initial fees based on
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
//     list of vehicle fees
//     any error encountered
//-----------------------------------------------------------------------------
func GetRAFlowInitialVehicleFees(
	ctx context.Context,
	BID, RID int64,
	rStart, rStop JSONDate,
	meta *RAFlowMetaInfo,
) (fees []RAFeesData, err error) {

	const funcname = "GetRAFlowInitialVehicleFees"
	var (
		bizPropName = "general"
		d1          = (time.Time)(rStart)
		d2          = (time.Time)(rStop)
	)
	fmt.Printf("Entered in %s\n", funcname)

	// INITIALIZE FEES
	fees = []RAFeesData{}

	// GET VEHICLE FEES FROM BUSINESS PROPERTIES
	var vehicleFees []BizPropsFee
	vehicleFees, err = GetBizPropVehicleFees(ctx, BID, bizPropName)
	if err != nil {
		return
	}

	// FOR EACH FEE CONFIGURED IN BIZPROP
	for _, vehicleFee := range vehicleFees {

		// GET RENT, PRORATION CYCLE
		RentCycle := vehicleFee.ARRentCycle
		ProrationCycle := vehicleFee.ARProrationCycle
		err = GetRAFlowFeeCycles(ctx, RID, d1, &RentCycle, &ProrationCycle)
		if err != nil {
			return
		}

		// ========================================================================
		// GET EPOCH BASED ON RENTCYCLE FOR THIS VEHICLE FEE
		// ========================================================================
		// TODO(Sudip & Steve): WHEN WE INTEGRATE EPOCHS IN RENTABLE TYPES       //
		//                      WE SHOULD TAKE EPOCHS FIRST FROM IT THEN         //
		//                      FROM BIZPROPS IN CASE NOT FOUND IN RENTALBE TYPE //
		// ========================================================================
		var epoch time.Time
		_, epoch, err = GetEpochByBizPropName(ctx, BID, bizPropName, d1, d2, RentCycle)
		if err != nil {
			return
		}

		//--------------------------------------------------------------
		// Here are the AR.FLAGS bits:
		//
		//     1<<4 -  Is Rent Assessment
		//     1<<6 -  Is non-recur charge
		//     1<<8 -  VID required
		//--------------------------------------------------------------
		// IF IT IS NON-RECUR CHARGE THEN
		oneTimeCharge := (vehicleFee.ARFLAGS & (1 << 6)) != 0
		if oneTimeCharge {
			// ADD FEE IN LIST
			meta.LastTMPASMID++
			raFee := RAFeesData{
				TMPASMID:        meta.LastTMPASMID,
				ASMID:           0,
				ARID:            vehicleFee.ARID,
				ARName:          vehicleFee.ARName,
				ContractAmount:  vehicleFee.Amount,
				RentCycle:       RECURNONE,
				Start:           rStart,
				Stop:            rStart,
				AtSigningPreTax: 0.00,
				SalesTax:        0.00,
				TransOccTax:     0.00,
				Comment:         "",
			}
			fees = append(fees, raFee)

		} else if vehicleFee.ARFLAGS&(1<<4) != 0 { // IT MUST BE RENT ASM ONE

			// CHECK FOR PRORATED AMOUNT REQUIRED
			needProratedRent := d1.Day() != epoch.Day()

			// START DAY IS NOT SAME AS EPOCH THEN CALCULATE PRORATED AMOUNT
			if needProratedRent {
				td2 := time.Date(d1.Year(), d1.Month(), epoch.Day(), d1.Hour(), d1.Minute(), d1.Second(), d1.Nanosecond(), d1.Location())
				td2 = NextPeriod(&td2, RentCycle)

				tot, np, tp := SimpleProrateAmount(vehicleFee.Amount, RentCycle, ProrationCycle, &d1, &td2, &epoch)
				cmt := ""
				if tot < vehicleFee.Amount {
					cmt = fmt.Sprintf("prorated for %d of %d %s", np, tp, ProrationUnits(ProrationCycle))
				}

				// ADD FEE IN LIST
				meta.LastTMPASMID++
				raFee := RAFeesData{
					TMPASMID:        meta.LastTMPASMID,
					ASMID:           0,
					ARID:            vehicleFee.ARID,
					ARName:          vehicleFee.ARName,
					ContractAmount:  tot,
					RentCycle:       RentCycle,
					Start:           rStart,
					Stop:            rStart,
					AtSigningPreTax: 0.00,
					SalesTax:        0.00,
					TransOccTax:     0.00,
					Comment:         cmt,
				}
				fees = append(fees, raFee)
			}

			// CALCULATE RECURRING ONE FROM EPOCH DATE
			// ADD FEE IN LIST
			meta.LastTMPASMID++
			raFee := RAFeesData{
				TMPASMID:        meta.LastTMPASMID,
				ASMID:           0,
				ARID:            vehicleFee.ARID,
				ARName:          vehicleFee.ARName,
				ContractAmount:  vehicleFee.Amount,
				RentCycle:       RentCycle,
				Start:           JSONDate(epoch),
				Stop:            rStop,
				AtSigningPreTax: 0.00,
				SalesTax:        0.00,
				TransOccTax:     0.00,
				Comment:         "",
			}
			fees = append(fees, raFee)
		}
	}

	return
}
