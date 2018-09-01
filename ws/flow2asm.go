package ws

// Fees2RA copies fees into permanent table Assessments.
import (
	"context"
	"fmt"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"time"
)

// Fees2RA handles all the updates necessary to move the fees defined in a flow
// into the permanent tables.
//
// INPUTS
//     ctx    - db context for transactions
//     x - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func Fees2RA(ctx context.Context, x *WriteHandlerContext) error {
	var err error

	//--------------------------------------------------
	// When was the last period closed?  Set the context
	// variable, x, so that all other routines have it.
	// Ensure that we have a valid lastClost.dt
	//--------------------------------------------------
	rlib.Console("Fees2RA - getting lastClose date\n")
	x.lastClose, err = rlib.GetLastClosePeriod(ctx, x.xbiz.P.BID)
	if err != nil {
		return err
	}
	if x.lastClose.CPID == 0 {
		x.lastClose.Dt = rlib.TIME0 // use TIME0 if not set
	}
	x.lastClose.OpenPeriodDt = x.lastClose.Dt.AddDate(0, 0, 1) // for our purposes, use the day after close
	rlib.Console("lastClose = %s\n", x.lastClose.Dt.Format(rlib.RRDATEREPORTFMT))

	//--------------------------------------------------
	// Add Rentable fees to new RA first...
	//--------------------------------------------------
	rlib.Console("Fees2RA: Rentables fees\n")
	for i := 0; i < len(x.raf.Rentables); i++ {
		for j := 0; j < len(x.raf.Rentables[i].Fees); j++ {
			if err = F2RASaveFee(ctx, x, &x.raf.Rentables[i].Fees[j], rlib.ELEMRENTABLE, x.raf.Rentables[i].RID, 0); err != nil {
				return err
			}
		}
	}
	//--------------------------------------------------
	// Add Pet fees to new RA first...
	//--------------------------------------------------
	rlib.Console("Fees2RA: Pet fees\n")
	for i := 0; i < len(x.raf.Pets); i++ {
		for j := 0; j < len(x.raf.Pets[i].Fees); j++ {
			if err = F2RASaveFee(ctx, x, &x.raf.Pets[i].Fees[j], rlib.ELEMPET, x.raf.Pets[i].PETID, x.raf.Pets[i].TMPTCID); err != nil {
				return err
			}
		}
	}
	//--------------------------------------------------
	// Add Vehicle fees to new RA first...
	//--------------------------------------------------
	rlib.Console("Fees2RA: Vehicle fees\n")
	for i := 0; i < len(x.raf.Vehicles); i++ {
		for j := 0; j < len(x.raf.Vehicles[i].Fees); j++ {
			if err = F2RASaveFee(ctx, x, &x.raf.Vehicles[i].Fees[j], rlib.ELEMVEHICLE, x.raf.Vehicles[i].VID, x.raf.Vehicles[i].TMPTCID); err != nil {
				return err
			}
		}
	}

	if err = F2RAHandleOldAssessments(ctx, x); err != nil {
		return err
	}

	//--------------------------------------------------------------------------
	// Now clean up any assessments that are associated with the old RAID but
	// that have not been updated as part of any fee in the new RAID.
	//--------------------------------------------------------------------------
	// if err = CleanUpRemainingAssessments(ctx, x); err != nil {
	// 	return err
	// }

	return nil
}

// F2RAHandleOldAssessments handles all the assessments associated with any
// RAID in the RAID chain for the ORIGIN RAID that are affected by the new
// amendment.
//
//
// INPUTS
//     ctx  - db context for transactions
//     x    - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func F2RAHandleOldAssessments(ctx context.Context, x *WriteHandlerContext) error {
	var err error
	var n []rlib.Assessment
	var skipASMID int64
	rlib.Console("Entered F2RAHandleOldAssessments\n")

	for _, ra := range x.raChainOrig {
		rlib.Console("%d ", ra.RAID)
	}
	rlib.Console("\n")

	rlib.Console("A2 - Processing RA Chain\n")
	for i := 0; i < len(x.raChainOrig); i++ {
		ra := x.raChainOrig[i]
		raUnchanged := x.raChainOrigUnchanged[i]

		rlib.Console("A3: ra.RAID = %d\n", ra.RAID)
		//-------------------------------------------------------------------------
		//  Only process if there's time overlap.  In this case we need to compare
		//  the time range of the old RA before any changes were made, so we need
		//  to use raUnchanged
		//-------------------------------------------------------------------------
		if !rlib.DateRangeOverlap(&x.ra.RentStart, &x.ra.RentStop, &raUnchanged.RentStart, &raUnchanged.RentStop) {
			rlib.Console("A3.1 no overlap: %s - %s, %s - %s\n", x.ra.RentStart.Format(rlib.RRDATEREPORTFMT), x.ra.RentStop.Format(rlib.RRDATEREPORTFMT), ra.RentStart.Format(rlib.RRDATEREPORTFMT), ra.RentStop.Format(rlib.RRDATEREPORTFMT))
			continue
		}

		rlib.Console("A4 - overlaps the amended RA\n")
		//-----------------------------------------------------------------------
		// Need to process this one. Start with its recurring asm definitions...
		//-----------------------------------------------------------------------
		n, err = rlib.GetRecurringAssessmentDefsByRAID(ctx, ra.RAID, &x.ra.RentStart, &x.ra.RentStop)
		if err != nil {
			return err
		}
		rlib.Console("A5 - found %d recurring Assessments for RAID %d\n", len(n), ra.RAID)
		for _, v := range n {
			rlib.Console("A6 - ASMID=%d\n", v.ASMID)
			if v.FLAGS&(1<<2) != 0 {
				continue // skip it if it has already been
			}
			//---------------------------------------------------------------
			//  The date we use for the change depends on whether or not the
			//  financial period at the start date of the assessment has
			//  been closed.
			//---------------------------------------------------------------
			dt := v.Start // assume it will be on the assessment start date
			if v.Start.Before(x.lastClose.Dt) {
				rlib.Console("A6.1 - v.Start is prior to the last close period. Snapping dt to: %s\n", dt.Format(rlib.RRDATERECEIPTFMT))
				dt = x.lastClose.Dt
			}

			rlib.Console("A6.2 - dt for changes = %s\n", dt.Format(rlib.RRDATEREPORTFMT))

			//---------------------------------------------------------------
			//  The assessment will be totally replaced if the new RA start
			//  date is prior to the Assessment start date.
			//---------------------------------------------------------------
			if !v.Start.Before(x.ra.RentStart) {
				//---------------------------------------------
				// Reverse the whole thing; all instances...
				//---------------------------------------------
				rlib.Console("A7 -- REVERSE THE ASSESSMENT!! amended RA starts prior to ASM Start: %s\n", v.Start.Format(rlib.RRDATEREPORTFMT))
				be := bizlogic.ReverseAssessment(ctx, &v, 2 /*from dt onward*/, &dt, &x.lastClose)
				if len(be) > 0 {
					rlib.Console("A7 error\n")
					return bizlogic.BizErrorListToError(be)
				}
			} else {
				rlib.Console("A8 -- REVERSE from this time forward\n")
				//-------------------------------------------------------------
				// Reverse the instances that occur in periods on or after
				// x.ra.RentStart.  We know the epoch as we have the  recurring
				// definition. So, we want to create the epoch day based on
				// v.Start and the start date of the switchover -- x.ra.RentStart
				//-------------------------------------------------------------
				target := rlib.InstanceDateCoveringDate(&v.Start, &x.ra.RentStart, v.RentCycle)
				t2 := target.AddDate(0, 0, 1)
				n, err = rlib.GetAssessmentInstancesByRAIDRange(ctx, ra.RAID, &target, &t2)
				if err != nil {
					return err
				}
				if len(n) == 0 {
					rlib.Console("A8.1 -- cannot find instance date near x.ra.RentStart!!\n")
				} else {
					rlib.Console("A9 - reversing assessments from = %s forward, starting with ASMID = %d\n", n[0].Start.Format(rlib.RRDATEREPORTFMT), n[0].ASMID)
					errlist := bizlogic.ReverseAssessment(ctx, &n[0], 1 /*this point forward*/, &dt, &x.lastClose)
					if len(errlist) > 0 {
						rlib.Console("A9 error\n")
						return bizlogic.BizErrorListToError(errlist)
					}
					//-------------------------------------------------------------
					// PRORATE ASSESSMENT IF NEEDED
					// Create a prorated assessment for the transition period only
					// if the switchover (x.ra.RentStart) date is NOT an instance
					// date of the the assessment (epoch = v.Start)
					//-------------------------------------------------------------
					isinst := rlib.IsInstanceDate(&target, &x.ra.RentStart, v.RentCycle, v.ProrationCycle)
					if !isinst {
						rlib.Console("A9.1 - new RA rentstart (%s) was found NOT to be an instance date of ASMID = %d\n", x.ra.RentStart.Format(rlib.RRDATEREPORTFMT), v.ASMID)
						//------------------------------------------------------
						// In this case we need to create a prorated assessment
						// that covers from target to x.ra.RentStart
						//-----------------------------------------------------
						asm := n[0]
						amt, count, totcount := rlib.SimpleProrateAmount(v.Amount, v.RentCycle, v.ProrationCycle, &target, &x.ra.RentStart, &target)
						asm.AppendComment(fmt.Sprintf("prorated for %d of %d %s", count, totcount, rlib.ProrationUnits(v.ProrationCycle)))
						asm.Amount = amt
						asm.RentCycle = rlib.RECURNONE      // not part of a series
						asm.ProrationCycle = rlib.RECURNONE // no proration here
						asm.FLAGS = 0
						asm.Stop = asm.Start
						if errlist := bizlogic.InsertAssessment(ctx, &asm, 0 /*no expanding*/, &x.lastClose); len(errlist) > 0 {
							return bizlogic.BizErrorListToError(errlist)
						}
						skipASMID = asm.ASMID
						rlib.Console("A9.2 - just inserted asm = %d, skipASMID set\n", skipASMID)
					} else {
						rlib.Console("A9.3 - new RA rentstart (%s) was found to be an instance date of ASMID = %d\n", x.ra.RentStart.Format(rlib.RRDATEREPORTFMT), v.ASMID)
						rlib.Console("     - so will not add a prorated rent assessment\n")
					}
				}
				//-------------------------------------------------------------
				// Set the stop date for v to x.ra.RentStart.  Since we've
				// already reversed only those instances that needed to be
				// reversed, we do not call the bizlogic version of this routine.
				// This is one of the rare exceptions where we just want to change
				// the end date and nothing else.
				//-------------------------------------------------------------
				v.Stop = x.ra.RentStart
				if err = rlib.UpdateAssessment(ctx, &v); err != nil {
					return err
				}
			}
		}

		rlib.Console("A10 - HANDLE INSTANCES\n")
		//-----------------------------------------------------------------------
		// Now handle instances...
		//-----------------------------------------------------------------------
		n, err = rlib.GetAssessmentInstancesByRAIDRange(ctx, ra.RAID, &x.ra.RentStart, &rlib.ENDOFTIME)
		if err != nil {
			return err
		}
		rlib.Console("A11 -  Found %d instances in the range %s - %s\n", len(n), x.ra.RentStart.Format(rlib.RRDATEREPORTFMT), rlib.ENDOFTIME.Format(rlib.RRDATEREPORTFMT))
		for _, v := range n {
			if v.ASMID == skipASMID {
				continue // this one is OK, we just added it
			}
			rlib.Console("A12 -  ASMID = %d\n", v.ASMID)
			if v.FLAGS&(1<<2) != 0 {
				rlib.Console("A12.1 - reversed, skipping\n")
				continue // skip reversed assessments
			}
			if v.Start.Before(x.ra.RentStart) {
				continue // v.Stop was >= x.ra.RentStart so the overlap period matched the query, not a problem, just skip it
			}
			//---------------------------------------------------------------
			//  The date we use for the change depends on whether or not the
			//  financial period at the start date of the assessment has
			//  been closed.
			//---------------------------------------------------------------
			dt := v.Start // assume it will be on the assessment start date
			if v.Start.Before(x.lastClose.Dt) {
				dt = x.lastClose.Dt
			}
			rlib.Console("A12.2 - Reversal dates will be as of: %s\n", dt.Format(rlib.RRDATEREPORTFMT))
			//----------------------------
			// Now process the instance
			//----------------------------
			if !v.Start.Before(x.ra.RentStart) {
				// Reverse the whole thing
				rlib.Console("A13 - Reversing ASMID = %d\n", v.ASMID)
				be := bizlogic.ReverseAssessment(ctx, &v, 0 /*this instance*/, &dt, &x.lastClose)
				if len(be) > 0 {
					rlib.Console("A13 error\n")
					bizlogic.PrintBizErrorList(be)
					return bizlogic.BizErrorListToError(be)
				}
			} else {
				// This should not happen. Checking for it just to make sure that
				// the code is working as expected
				rlib.Console("\n\n**** ERROR ****    **** ERROR ****    **** ERROR ****    \n\n")
				rlib.Console("\nLook for this line of code in F2RAHandleOldAssessments()\n")
				rlib.Console("Assessment ASMID = %d, Start date is out of expected range:  %s\n", v.ASMID, v.Start.Format(rlib.RRDATEREPORTFMT))
				rlib.Console("\n\n**** ERROR ****    **** ERROR ****    **** ERROR ****    \n\n")
			}
		}
	}
	rlib.Console("A14\n")

	return nil
}

// // F2RASaveFee handles all the updates necessary to move the
// // supplied fee into the permanent tables.
// //
// // INPUTS
// //     ctx  - db context for transactions
// //     x    - all the contextual info we need for performing this operation
// //     elt  - element type if is this is bound to a pet or vehicle
// //     id   - RID if elt == rlib.ELEMRENTABLE, or tmpid of the element
// //            (TMPPETID, TMPVID), valid if elt > 0
// //     tcid - tmptcid of the transactant responsible, valid if elt > 0
// //
// // RETURNS
// //     Any errors encountered
// //-----------------------------------------------------------------------------
// func F2RASaveFee(ctx context.Context, x *WriteHandlerContext, fee *rlib.RAFeesData, eltype, id, tmptcid int64) error {
//
// 	// SKIPPING ALL THIS FOR NOW...   I THINK F2RAHandleOldAssessments SHOULD
// 	// HANDLE EVERYTHING...
// 	// VERIFY and REMOVE THIS CODE IF SO.
//
// 	// rlib.Console("F2RASaveFee processing fee = %d, ASMID = %d\n", fee.TMPASMID, fee.ASMID)
// 	// if 0 < fee.ASMID {
// 	// 	return F2RAUpdateExistingAssessment(ctx, x, fee, eltype, id, tmptcid)
// 	// }
// 	return F2RASaveNewFee(ctx, x, fee, eltype, id, tmptcid)
// }

// F2RASaveFee handles all the updates necessary to move the
// supplied fee into the permanent tables.
//
// INPUTS
//     ctx  - db context for transactions
//     x    - all the contextual info we need for performing this operation
//     elt  - element type if is this is bound to a pet or vehicle
//     id   - RID if elt == rlib.ELEMRENTABLE, or tmpid of the element
//            (TMPPETID, TMPVID), valid if elt > 0
//     tcid - tmptcid of the transactant responsible, valid if elt > 0
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func F2RASaveFee(ctx context.Context, x *WriteHandlerContext, fee *rlib.RAFeesData, eltype, id, tmptcid int64) error {
	// rlib.Console("Entered F2RASaveFee\n")
	//-------------------------------------------------------------------
	// Create a new assessment from this day forward...
	//-------------------------------------------------------------------
	var b rlib.Assessment
	dt := time.Time(x.raf.Dates.AgreementStart)
	if fee.ASMID > 0 {
		b.AppendComment(fmt.Sprintf("Continuation of ASMID %d from RAID %d", fee.ASMID, x.raf.Meta.RAID))
	}
	Start := time.Time(fee.Start) // the start time will be either the fee start
	if Start.Before(dt) {         // or the start of the new rental agreement
		Start = dt // whichever is later
	}
	b.Stop = time.Time(fee.Stop)
	b.BID = x.raf.Meta.BID

	//-------------------------------------------------------------------
	// Set the Element Type and ID if necessary
	//-------------------------------------------------------------------
	b.AssocElemType = eltype
	b.AssocElemID = id

	//-------------------------------------------------------------------
	// find the RID associated with this pet
	//-------------------------------------------------------------------
	switch eltype {
	case rlib.ELEMRENTABLE:
		b.RID = id
	case rlib.ELEMPET:
		if b.RID = GetRIDForTMPTCID(ctx, x, tmptcid); b.RID <= 0 {
			return fmt.Errorf("No RID associated with TMPTCID = %d", tmptcid)
		}
		// rlib.Console("GetRIDForTMPTCID( TMPTCID=%d) ===> %d\n", tmptcid, b.RID)
		// rlib.Console("    ID for this pet is %d\n", b.AssocElemID)
	case rlib.ELEMVEHICLE:
		if b.RID = GetRIDForTMPTCID(ctx, x, tmptcid); b.RID <= 0 {
			return fmt.Errorf("No RID associated with TMPTCID = %d", tmptcid)
		}
		// rlib.Console("GetRIDForTMPTCID( TMPTCID=%d) ===> %d\n", tmptcid, b.RID)
		// rlib.Console("    ID for this vehicle is %d\n", b.AssocElemID)
	}

	// rlib.Console("bid = %d, fee ARID = %d\n", b.BID, fee.ARID)
	b.Amount = fee.ContractAmount
	b.AcctRule = ""
	b.RentCycle = fee.RentCycle
	b.RAID = x.ra.RAID
	b.Start = time.Time(fee.Start)
	b.Stop = time.Time(fee.Stop)
	b.RentCycle = fee.RentCycle
	b.ProrationCycle = rlib.RRdb.BizTypes[b.BID].AR[fee.ARID].DefaultProrationCycle
	b.InvoiceNo = 0
	b.ARID = fee.ARID
	switch eltype {
	case rlib.ELEMRENTABLE:
		// nothing to do at this time
	case rlib.ELEMPET:
		b.FLAGS |= 1 << 3  // PETID required
		b.AssocElemID = id // must be the PETID
		b.AssocElemType = eltype
	case rlib.ELEMVEHICLE:
		b.FLAGS |= 1 << 4  // VID required
		b.AssocElemID = id // must be the PETID
		b.AssocElemType = eltype
	}
	if errlist := bizlogic.InsertAssessment(ctx, &b, 1 /*expand*/, &x.lastClose); len(errlist) > 0 {
		return bizlogic.BizErrorListToError(errlist)
	}

	return nil
}

// GetRIDForTMPTCID finds the RID associated with the supplied tmptcid.
// This routine is called when we have a pet or a vehicle and we need to
// know what RID to associate it with. The RATiePeopleData datastruct is
// organized by tmptcid.  PRID is the Parent RID for that person.
//
// INPUTS
//     ctx     - db context for transactions
//     x       - all the contextual info we need for performing this operation
//     tmptcid - tmptcid for person we want the associated RID
//
// RETURNS
//     RID of associated rentable, or -1 if not found
//-----------------------------------------------------------------------------
func GetRIDForTMPTCID(ctx context.Context, x *WriteHandlerContext, tmptcid int64) int64 {
	for i := 0; i < len(x.raf.Tie.People); i++ {
		if x.raf.Tie.People[i].TMPTCID == tmptcid {
			return x.raf.Tie.People[i].PRID
		}
	}
	return -1
}

// CleanUpRemainingAssessments handles all the assessments associated with the
// original RAID that were not found while handling the amended RAID.
//
// 1. Get RAChain
// 2. for each RA that overlaps x.ra.RentStart/Stop
// 3.    get all assessment instances that overlap x.ra.RentStart/Stop
// 4.    reverse the assessments
//
// INPUTS
//     ctx  - db context for transactions
//     x    - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
// func CleanUpRemainingAssessments(ctx context.Context, x *WriteHandlerContext) error {
// 	rlib.Console("Entered CleanUpRemainingAssessments")
// 	if x.raf.Meta.RAID == 0 {
// 		rlib.Console("No cleanup necessary. x.raf.Meta.RAID is 0\n")
// 		return nil // nothing to do, no old RAID
// 	}
// 	//--------------------------------------------------------------------------
// 	// Get the list of any recurring assessments associated with the old rental
// 	// agreement that overlap the time range of the new rental agreement.
// 	//--------------------------------------------------------------------------
// 	m, err := rlib.GetRecurringAssessmentDefsByRAID(ctx, x.raOrig.RAID, &x.ra.RentStart, &x.ra.RentStop)
// 	if err != nil {
// 		return err
// 	}
// 	rlib.Console("Found %d recurring assessment definitions to review\n", len(m))
// 	for _, v := range m {
// 		rlib.Console("ASMID = %d\n", v.ASMID)
// 		if v.RentCycle == rlib.RECURNONE {
// 			// If it is a non-recurring assessment, reverse it.
// 			be := bizlogic.ReverseAssessment(ctx, &v, 0, &x.ra.RentStart)
// 			if len(be) > 0 {
// 				return bizlogic.BizErrorListToError(be)
// 			}
// 		} else {
// 			// If it is a recurring assessment, stop it.
// 			if err = bizlogic.UpdateAssessmentEndDate(ctx, &v, &x.ra.RentStart); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	rlib.Console("*** Completed processing recurring assessment definitions\n")
// 	//--------------------------------------------------------------------------
// 	// Anything non-recurring that happens as of the start date of the amended
// 	// agreement must be deleted (reversed).
// 	//--------------------------------------------------------------------------
// 	m, err = rlib.GetNorecurAssessmentsByRAIDRange(ctx, x.raOrig.RAID, &x.ra.RentStart, &x.ra.RentStop)
// 	if err != nil {
// 		return err
// 	}
// 	rlib.Console("Found %d non-recurring assessments to reverse\n", len(m))
// 	for _, v := range m {
// 		v.AppendComment(fmt.Sprintf("Reversing due to amended RAID %d", x.ra.RAID))
// 		bizlogic.ReverseAssessment(ctx, &v, 0, &x.ra.RentStart)
// 	}
// 	rlib.Console("*** Completed processing non-recurring\n")
// 	return nil
// }

// F2RAUpdateExistingAssessment handles all the updates necessary to move the
// supplied fee into the permanent tables.
//
//
// INPUTS
//     ctx  - db context for transactions
//     x    - all the contextual info we need for performing this operation
//     elt  - element type if is this is bound to a pet or vehicle
//     id   - RID if elt == rlib.ELEMRENTABLE, or tmpid of the element
//            (TMPPETID, TMPVID), valid if elt > 0
//     tcid - tmptcid of the transactant responsible, valid if elt > 0
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
// func F2RAUpdateExistingAssessment(ctx context.Context, x *WriteHandlerContext, fee *rlib.RAFeesData, eltype, id, tmptcid int64) error {
// 	rlib.Console("Entered F2RAUpdateExistingAssessment\n")
// 	if fee.ASMID == int64(0) {
// 		return fmt.Errorf("fee.ASMID must be > 0")
// 	}
// 	a, err := rlib.GetAssessment(ctx, fee.ASMID)
// 	if err != nil {
// 		return err
// 	}
//
// 	//-------------------------------------------------------------------
// 	// skip any assessments that finished prior to this Rental Agreement
// 	//-------------------------------------------------------------------
// 	dt := time.Time(x.raf.Dates.AgreementStart)
// 	stop := time.Time(fee.Stop)
// 	if stop.Before(dt) {
// 		return nil // don't need to process this one
// 	}
//
// 	//-------------------------------------------------------------------
// 	// skip any non-recurring assessment that has been paid...
// 	//-------------------------------------------------------------------
// 	if a.FLAGS&3 == 2 {
// 		return nil // don't need to process this
// 	}
//
// 	//-------------------------------------------------------------------
// 	// If it's recurring we'll just stop it on the start date of the new
// 	// rental agreement
// 	//-------------------------------------------------------------------
// 	if a.RentCycle > rlib.RECURNONE {
// 		err = bizlogic.UpdateAssessmentEndDate(ctx, &a, &dt)
// 		if err != nil {
// 			return err
// 		}
// 		// a.Stop = dt
// 		// if err = rlib.UpdateAssessment(ctx, &a); err != nil {
// 		// 	return err
// 		// }
// 	}
//
// 	err = F2RASaveNewFee(ctx, x, fee, eltype, id, tmptcid)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
