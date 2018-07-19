package rlib

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
