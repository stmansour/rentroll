package rlib

import (
	"context"
	"fmt"
	"time"
)

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
//     list of pet fees
//     any error encountered
//-----------------------------------------------------------------------------
func GetRAFlowInitialPetFees(ctx context.Context,
	BID int64,
	rStart, rStop time.Time,
	meta *RAFlowMetaInfo,
) (fees []RAFeesData, err error) {

	const funcname = "GetRAFlowInitialPetFees"
	var (
		bizPropName = "general"
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

		// ========================================================================
		// GET EPOCH BASED ON RENTCYCLE FOR THIS PET FEE
		// ========================================================================
		// TODO(Sudip & Steve): WHEN WE INTEGRATE EPOCHS IN RENTABLE TYPES       //
		//                      WE SHOULD TAKE EPOCHS FIRST FROM IT THEN         //
		//                      FROM BIZPROPS IN CASE NOT FOUND IN RENTALBE TYPE //
		// ========================================================================
		var epoch time.Time
		_, epoch, err = GetEpochByBizPropName(ctx, BID, bizPropName, rStart, rStop, RentCycle)
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
				Start:           JSONDate(rStart),
				Stop:            JSONDate(rStart),
				AtSigningPreTax: 0.00,
				SalesTax:        0.00,
				TransOccTax:     0.00,
				Comment:         "",
			}
			fees = append(fees, raFee)

		} else if petFee.ARFLAGS&(1<<4) != 0 { // IT MUST BE RENT ASM ONE

			// CHECK FOR PRORATED AMOUNT REQUIRED
			needProratedRent := rStart.Day() != epoch.Day()

			// START DAY IS NOT SAME AS EPOCH THEN CALCULATE PRORATED AMOUNT
			if needProratedRent {
				td2 := time.Date(rStart.Year(), rStart.Month(), epoch.Day(), rStart.Hour(), rStart.Minute(), rStart.Second(), rStart.Nanosecond(), rStart.Location())
				td2 = NextPeriod(&td2, RentCycle)

				tot, np, tp := SimpleProrateAmount(petFee.Amount, RentCycle, ProrationCycle, &rStart, &td2, &epoch)
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
					Start:           JSONDate(rStart),
					Stop:            JSONDate(rStart),
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
				Stop:            JSONDate(rStop),
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
//          rStart  = rent start date
//           rStop  = rent stop date
//            meta  = RAFlowMetaInfo data
//
// RETURNS
//     list of vehicle fees
//     any error encountered
//-----------------------------------------------------------------------------
func GetRAFlowInitialVehicleFees(ctx context.Context,
	BID int64,
	rStart, rStop time.Time,
	meta *RAFlowMetaInfo,
) (fees []RAFeesData, err error) {

	const funcname = "GetRAFlowInitialVehicleFees"
	var (
		bizPropName = "general"
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

		// ========================================================================
		// GET EPOCH BASED ON RENTCYCLE FOR THIS VEHICLE FEE
		// ========================================================================
		// TODO(Sudip & Steve): WHEN WE INTEGRATE EPOCHS IN RENTABLE TYPES       //
		//                      WE SHOULD TAKE EPOCHS FIRST FROM IT THEN         //
		//                      FROM BIZPROPS IN CASE NOT FOUND IN RENTALBE TYPE //
		// ========================================================================
		var epoch time.Time
		_, epoch, err = GetEpochByBizPropName(ctx, BID, bizPropName, rStart, rStop, RentCycle)
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
				Start:           JSONDate(rStart),
				Stop:            JSONDate(rStart),
				AtSigningPreTax: 0.00,
				SalesTax:        0.00,
				TransOccTax:     0.00,
				Comment:         "",
			}
			fees = append(fees, raFee)

		} else if vehicleFee.ARFLAGS&(1<<4) != 0 { // IT MUST BE RENT ASM ONE

			// CHECK FOR PRORATED AMOUNT REQUIRED
			needProratedRent := rStart.Day() != epoch.Day()

			// START DAY IS NOT SAME AS EPOCH THEN CALCULATE PRORATED AMOUNT
			if needProratedRent {
				td2 := time.Date(rStart.Year(), rStart.Month(), epoch.Day(), rStart.Hour(), rStart.Minute(), rStart.Second(), rStart.Nanosecond(), rStart.Location())
				td2 = NextPeriod(&td2, RentCycle)

				tot, np, tp := SimpleProrateAmount(vehicleFee.Amount, RentCycle, ProrationCycle, &rStart, &td2, &epoch)
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
					Start:           JSONDate(rStart),
					Stop:            JSONDate(rStart),
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
				Stop:            JSONDate(rStop),
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
