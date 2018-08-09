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
	var petBizFees []BizPropsFee
	petBizFees, err = GetBizPropPetFees(ctx, BID, bizPropName)
	if err != nil {
		return
	}

	// PREPARE THE PET BASE FEES FROM BIZ FEES
	var petFees = []RAFeesData{}
	for i := range petBizFees {
		raFee := RAFeesData{ContractAmount: petBizFees[i].Amount}
		MigrateStructVals(&petBizFees[i], &raFee)
		raFee.RentCycle = petBizFees[i].ARRentCycle
		raFee.ProrationCycle = petBizFees[i].ARProrationCycle
		petFees = append(petFees, raFee)
	}

	// GET CALCULATED FEES FROM THIS BIZ CONIGURED FEES
	fees, err = GetCalculatedFeesFromBaseFees(ctx, BID, bizPropName, rStart, rStop, petFees, false)
	if err != nil {
		return
	}

	// UPDATE LASTASMID FOR EACH FEE
	for i := range fees {
		meta.LastTMPASMID++
		fees[i].TMPASMID = meta.LastTMPASMID
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
	var vehicleBizFees []BizPropsFee
	vehicleBizFees, err = GetBizPropVehicleFees(ctx, BID, bizPropName)
	if err != nil {
		return
	}

	// PREPARE THE VEHICLE BASE FEES FROM BIZ FEES
	var vehicleFees = []RAFeesData{}
	for i := range vehicleBizFees {
		raFee := RAFeesData{ContractAmount: vehicleBizFees[i].Amount}
		MigrateStructVals(&vehicleBizFees[i], &raFee)
		raFee.RentCycle = vehicleBizFees[i].ARRentCycle
		raFee.ProrationCycle = vehicleBizFees[i].ARProrationCycle
		vehicleFees = append(vehicleFees, raFee)
	}

	// GET CALCULATED FEES FROM THIS BIZ CONIGURED FEES
	fees, err = GetCalculatedFeesFromBaseFees(ctx, BID, bizPropName, rStart, rStop, vehicleFees, false)
	if err != nil {
		return
	}

	// UPDATE LASTASMID FOR EACH FEE
	for i := range fees {
		meta.LastTMPASMID++
		fees[i].TMPASMID = meta.LastTMPASMID
	}

	return
}

// GetCalculatedFeesFromBaseFees get the actual calculate fees based on
// given base fees list
func GetCalculatedFeesFromBaseFees(ctx context.Context, BID int64, bizPropName string,
	rStart, rStop time.Time,
	baseFees []RAFeesData,
	removeOneTimeCharge bool,
) (fees []RAFeesData, err error) {

	const funcname = "GetCalculatedFeesFromBaseFees"
	fmt.Printf("Entered in %s, \n", funcname)
	// fmt.Printf("%s: removeOneTimeCharge: %t, rStart: %s, rStop: %s\n", funcname,
	// 	removeOneTimeCharge, rStart.Format(RRDATEFMT3), rStop.Format(RRDATEFMT3))

	// INITIALIZE FEES
	fees = []RAFeesData{}

	// FOR EACH FEE FROM BASE FEES
	for _, baseFee := range baseFees {

		// GET AR FROM ARID IN FEES
		var ar AR
		ar, err = GetAR(ctx, baseFee.ARID)
		if err != nil {
			return
		}

		// GET RENT, PRORATION CYCLE
		RentCycle := baseFee.RentCycle
		ProrationCycle := baseFee.ProrationCycle

		// ========================================================================
		// GET EPOCH BASED ON RENTCYCLE FOR THIS BASE FEE
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
		oneTimeCharge := (ar.FLAGS & (1 << 6)) != 0
		rentAsmCharge := (ar.FLAGS & (1 << 4)) != 0

		if oneTimeCharge && !removeOneTimeCharge {
			// ADD FEE IN LIST
			raFee := RAFeesData{
				TMPASMID:        0,
				ASMID:           0,
				ARID:            baseFee.ARID,
				ARName:          baseFee.ARName,
				ContractAmount:  baseFee.ContractAmount,
				RentCycle:       RECURNONE,
				Start:           JSONDate(rStart),
				Stop:            JSONDate(rStart),
				AtSigningPreTax: 0.00,
				SalesTax:        0.00,
				TransOccTax:     0.00,
				Comment:         "",
			}
			fees = append(fees, raFee)
		} else if rentAsmCharge { // IT MUST BE RENT ASM ONE

			// CHECK FOR PRORATED AMOUNT REQUIRED
			needProratedRent := rStart.Day() != epoch.Day()

			// START DAY IS NOT SAME AS EPOCH THEN CALCULATE PRORATED AMOUNT
			if needProratedRent {
				td2 := time.Date(rStart.Year(), rStart.Month(), epoch.Day(), rStart.Hour(), rStart.Minute(), rStart.Second(), rStart.Nanosecond(), rStart.Location())
				td2 = NextPeriod(&td2, RentCycle)

				tot, np, tp := SimpleProrateAmount(baseFee.ContractAmount, RentCycle, ProrationCycle, &rStart, &td2, &epoch)
				cmt := ""
				if tot < baseFee.ContractAmount {
					cmt = fmt.Sprintf("prorated for %d of %d %s", np, tp, ProrationUnits(ProrationCycle))
				}

				// ADD FEE IN LIST
				raFee := RAFeesData{
					TMPASMID:        0,
					ASMID:           0,
					ARID:            baseFee.ARID,
					ARName:          baseFee.ARName,
					ContractAmount:  tot,
					RentCycle:       RECURNONE,
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
			raFee := RAFeesData{
				TMPASMID:        0,
				ASMID:           0,
				ARID:            baseFee.ARID,
				ARName:          baseFee.ARName,
				ContractAmount:  baseFee.ContractAmount,
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
