package bizlogic

// Fees2RA copies fees into permanent table Assessments.
import (
	"context"
	"fmt"
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
func Fees2RA(ctx context.Context, x *rlib.F2RAWriteHandlerContext) error {
	var err error

	//--------------------------------------------------
	// When was the last period closed?  Set the context
	// variable, x, so that all other routines have it.
	// Ensure that we have a valid lastClost.dt
	//--------------------------------------------------
	rlib.Console("Fees2RA - getting LastClose date\n")
	x.LastClose, err = rlib.GetLastClosePeriod(ctx, x.Xbiz.P.BID)
	if err != nil {
		return err
	}
	if x.LastClose.CPID == 0 {
		x.LastClose.Dt = rlib.TIME0               // use TIME0 if not set
		x.LastClose.ExpandAsmDtStart = rlib.TIME0 //
	}
	x.LastClose.ExpandAsmDtStop = x.Ra.RentStop                // do not expand past this date
	x.LastClose.OpenPeriodDt = x.LastClose.Dt.AddDate(0, 0, 1) // for our purposes, use the day after close
	rlib.Console("x.LastClose.Dt = %s, x.LastClose.OpenPeriodDt = %s, x.LastClose.ExpandAsmDtStart = %s\n",
		x.LastClose.Dt.Format(rlib.RRDATEREPORTFMT), x.LastClose.OpenPeriodDt.Format(rlib.RRDATEREPORTFMT), x.LastClose.ExpandAsmDtStop.Format(rlib.RRDATEREPORTFMT))

	//--------------------------------------------------
	// Add Rentable fees to new RA first...
	//--------------------------------------------------
	rlib.Console("Fees2RA: Rentables fees\n")
	for i := 0; i < len(x.Raf.Rentables); i++ {
		for j := 0; j < len(x.Raf.Rentables[i].Fees); j++ {
			rlib.Console("\tRentables[%d].Fees[%d]:  AR = %s, Amount = %8.2f\n", i, j, x.Raf.Rentables[i].Fees[j].ARName, x.Raf.Rentables[i].Fees[j].ContractAmount)
			if err = F2RASaveFee(ctx, x, &x.Raf.Rentables[i].Fees[j], rlib.ELEMRENTABLE, x.Raf.Rentables[i].RID, 0); err != nil {
				return err
			}
		}
	}
	//--------------------------------------------------
	// Add Pet fees to new RA first...
	//--------------------------------------------------
	rlib.Console("Fees2RA: Pet fees\n")
	for i := 0; i < len(x.Raf.Pets); i++ {
		for j := 0; j < len(x.Raf.Pets[i].Fees); j++ {
			rlib.Console("\tPets[%d].Fees[%d]:  AR = %s, Amount = %8.2f\n", i, j, x.Raf.Pets[i].Fees[j].ARName, x.Raf.Pets[i].Fees[j].ContractAmount)
			if err = F2RASaveFee(ctx, x, &x.Raf.Pets[i].Fees[j], rlib.ELEMPET, x.Raf.Pets[i].PETID, x.Raf.Pets[i].TMPTCID); err != nil {
				return err
			}
		}
	}
	//--------------------------------------------------
	// Add Vehicle fees to new RA first...
	//--------------------------------------------------
	rlib.Console("Fees2RA: Vehicle fees\n")
	for i := 0; i < len(x.Raf.Vehicles); i++ {
		for j := 0; j < len(x.Raf.Vehicles[i].Fees); j++ {
			rlib.Console("\tVehicles[%d].Fees[%d]:  AR = %s, Amount = %8.2f\n", i, j, x.Raf.Vehicles[i].Fees[j].ARName, x.Raf.Vehicles[i].Fees[j].ContractAmount)
			if err = F2RASaveFee(ctx, x, &x.Raf.Vehicles[i].Fees[j], rlib.ELEMVEHICLE, x.Raf.Vehicles[i].VID, x.Raf.Vehicles[i].TMPTCID); err != nil {
				return err
			}
		}
	}

	if err = F2RAHandleOldAssessments(ctx, x); err != nil {
		return err
	}

	return nil
}

// F2RASaveFee handles all the updates necessary to move the supplied fee into
// the permanent tables.  Remember that the new Rental Agreement may be back-
// dated which means that recurring assessments may need new instances to be
// generated so that we're up-to-date
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
func F2RASaveFee(ctx context.Context, x *rlib.F2RAWriteHandlerContext, fee *rlib.RAFeesData, eltype, id, tmptcid int64) error {
	rlib.Console("Entered F2RASaveFee, x.LastClose.ExpandAsmDtStart = %s\n", x.LastClose.ExpandAsmDtStart.Format(rlib.RRDATEREPORTFMT))
	//-------------------------------------------------------------------
	// Create a new assessment from this day forward...
	//-------------------------------------------------------------------
	var b rlib.Assessment
	b.Comment = fee.Comment
	dt := time.Time(x.Raf.Dates.AgreementStart)
	if fee.ASMID > 0 {
		b.AppendComment(fmt.Sprintf("Derived from RAID %d, ASMID %d", x.Raf.Meta.RAID, fee.ASMID))
	}
	Start := time.Time(fee.Start) // the start time will be either the fee start
	if Start.Before(dt) {         // or the start of the new rental agreement
		Start = dt // whichever is later
	}
	b.Stop = time.Time(fee.Stop)
	b.BID = x.Raf.Meta.BID

	//-------------------------------------------------------------------
	// Set the Element Type and ID if necessary
	//-------------------------------------------------------------------
	b.AssocElemType = eltype
	b.AssocElemID = id

	//-------------------------------------------------------------------
	// find the RID associated with this pet
	//-------------------------------------------------------------------
	rlib.Console("F2RASaveFee A\n")
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

	rlib.Console("F2RASaveFee B\n")
	//-------------------------------
	// Handle EDI on date range...
	//-------------------------------
	d1 := time.Time(fee.Start)
	d2 := time.Time(fee.Stop)
	rlib.EDIHandleIncomingDateRange(b.BID, &d1, &d2)

	// rlib.Console("bid = %d, fee ARID = %d\n", b.BID, fee.ARID)
	b.Amount = fee.ContractAmount
	b.AcctRule = ""
	b.RentCycle = fee.RentCycle
	b.RAID = x.Ra.RAID
	b.Start = d1
	b.Stop = d2
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

	rlib.Console("F2RASaveFee C\n")
	// rlib.Console("\n\n$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$\n")
	// rlib.Console("InsertAssessment: Amount = %8.2f, RentCycle = %d, Start = %s, Stop = %s\n", b.Amount, b.RentCycle, b.Start.Format(rlib.RRDATEFMT3), b.Stop.Format(rlib.RRDATEFMT3))
	// rlib.Console("x.LastClose.CPID = %d, x.LastClose.Dt = %s\n", x.LastClose.CPID, x.LastClose.Dt.Format(rlib.RRDATEREPORTFMT))
	// rlib.Console("x.LastClose.ExpandAsmDtStart = %s, x.LastClose.ExpandAsmDtStart = %s\n", x.LastClose.ExpandAsmDtStart.Format(rlib.RRDATEREPORTFMT), x.LastClose.ExpandAsmDtStart.Format(rlib.RRDATEREPORTFMT))
	// rlib.Console("x.LastClose.ExpandAsmDtStop  = %s, x.LastClose.ExpandAsmDtStop  = %s\n", x.LastClose.ExpandAsmDtStop.Format(rlib.RRDATEREPORTFMT), x.LastClose.ExpandAsmDtStop.Format(rlib.RRDATEREPORTFMT))
	// rlib.Console("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$\n\n")

	if errlist := InsertAssessment(ctx, &b, 1 /*expand*/, &x.LastClose); len(errlist) > 0 {
		return BizErrorListToError(errlist)
	}
	rlib.Console("just inserted ASMID = %d\n", b.ASMID)
	rlib.Console("**********\n\n")

	rlib.Console("F2RASaveFee D\n")

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
func GetRIDForTMPTCID(ctx context.Context, x *rlib.F2RAWriteHandlerContext, tmptcid int64) int64 {
	for i := 0; i < len(x.Raf.Tie.People); i++ {
		if x.Raf.Tie.People[i].TMPTCID == tmptcid {
			return x.Raf.Tie.People[i].PRID
		}
	}
	return -1
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
func F2RAHandleOldAssessments(ctx context.Context, x *rlib.F2RAWriteHandlerContext) error {
	var err error
	var n []rlib.Assessment
	var skipASMID int64
	rlib.Console("Entered F2RAHandleOldAssessments\n")

	//=========================================================================
	//  FOR EVERY RENTAL AGREEMENT THAT IS IMPACTED BY THIS UPDATE...
	//=========================================================================
	for i := 0; i < len(x.RaChainOrig); i++ {

		ra := x.RaChainOrig[i]
		// rlib.Console("Setting ExpandAsmDtStop to RentStop of RAID %d: %s\n", ra.RAID, ra.RentStop.Format(rlib.RRDATEFMT3))
		x.LastClose.ExpandAsmDtStop = ra.RentStop // do not expand past this date
		raUnchanged := x.RaChainOrigUnchanged[i]

		// rlib.Console("A3: ra.RAID = %d\n", ra.RAID)
		//-------------------------------------------------------------------------
		//  Only process if there's time overlap.  In this case we need to compare
		//  the time range of the old RA before any changes were made, so we need
		//  to use raUnchanged
		//-------------------------------------------------------------------------
		if !rlib.DateRangeOverlap(&x.Ra.RentStart, &x.Ra.RentStop, &raUnchanged.RentStart, &raUnchanged.RentStop) {
			// rlib.Console("A3.1 no overlap: %s - %s, %s - %s\n", x.Ra.RentStart.Format(rlib.RRDATEREPORTFMT), x.Ra.RentStop.Format(rlib.RRDATEREPORTFMT), ra.RentStart.Format(rlib.RRDATEREPORTFMT), ra.RentStop.Format(rlib.RRDATEREPORTFMT))
			continue
		}

		// rlib.Console("A4 - overlaps the amended RA\n")
		//-----------------------------------------------------------------------
		// Need to process this one. Start with its recurring asm definitions...
		//-----------------------------------------------------------------------
		n, err = rlib.GetRecurringAssessmentDefsByRAID(ctx, ra.RAID, &x.Ra.RentStart, &x.Ra.RentStop)
		if err != nil {
			return err
		}
		rlib.Console("A5 - found %d recurring Assessments for RAID %d\n", len(n), ra.RAID)

		//=========================================================================
		//  FOR EVERY RECURRING ASSESSMENT DEFINITION IN THIS RENTAL AGREEMENT...
		//=========================================================================
		for _, v := range n {
			rlib.Console("A6 - ASMID=%d\n", v.ASMID)
			if v.FLAGS&(1<<2) != 0 {
				continue // skip it if it has already been Reversed
			}
			//---------------------------------------------------------------
			//  The date we use for the change depends on whether or not the
			//  financial period at the start date of the assessment has
			//  been closed.
			//---------------------------------------------------------------
			dt := v.Start // assume it will be on the assessment start date
			if v.Start.Before(x.LastClose.Dt) {
				// rlib.Console("A6.1 - v.Start is prior to the last close period. Snapping dt to: %s\n", dt.Format(rlib.RRDATERECEIPTFMT))
				dt = x.LastClose.Dt
			}

			// rlib.Console("A6.2 - dt for changes = %s\n", dt.Format(rlib.RRDATEREPORTFMT))

			rlib.Console("A6.3 - v Start/Stop = %s\n", rlib.ConsoleDRange(&v.Start, &v.Stop))
			//---------------------------------------------------------------
			//  The assessment will be totally replaced if the new RA start
			//  date is prior to the Assessment start date.
			//---------------------------------------------------------------
			if !v.Start.Before(x.Ra.RentStart) {
				//---------------------------------------------
				// Reverse the whole thing; all instances...
				//---------------------------------------------
				rlib.Console("A7 -- REVERSE THE ASSESSMENT!! amended RA starts prior to ASM Start: %s\n", v.Start.Format(rlib.RRDATEREPORTFMT))
				//---------------------------------------------------------------------
				// This call reverses the supplied instances and all future instances
				//---------------------------------------------------------------------
				be := ReverseAssessment(ctx, &v, 2 /*from dt onward*/, &dt, &x.LastClose)
				if len(be) > 0 {
					// rlib.Console("A7 error\n")
					return BizErrorListToError(be)
				}
			} else {
				rlib.Console("A8 -- REVERSE from this time forward\n")
				//-------------------------------------------------------------
				// Reverse the instances that occur in periods on or after
				// x.Ra.RentStart.  We know the epoch as we have the  recurring
				// definition. So, we want to create the epoch day based on
				// v.Start and the start date of the switchover -- x.Ra.RentStart
				//-------------------------------------------------------------
				target := rlib.InstanceDateCoveringDate(&v.Start, &x.Ra.RentStart, v.RentCycle)
				t2 := target.AddDate(0, 0, 1)
				// rlib.Console("A8.1  v.ASMID=%d, target=%s, t2=%s\n", v.ASMID, target.Format(rlib.RRDATEREPORTFMT), t2.Format(rlib.RRDATEREPORTFMT))
				// n, err = rlib.GetAssessmentInstancesByRAIDRange(ctx, ra.RAID, &target, &t2)
				n, err = rlib.GetInstancesByDateRange(ctx, v.ASMID, &target, &t2)
				// rlib.Console("A8.12\n")
				if err != nil {
					// rlib.Console("A8.15 GetInstancesByDateRange returns err: %\n", err.Error())
					return err
				}
				// rlib.Console("A8.2 GetInstancesByDateRange returns %d matches\n", len(n))

				if len(n) == 0 {
					// rlib.Console("A8.5 -- cannot find instance date near x.Ra.RentStart!!\n")
				} else {
					// rlib.Console("A9 - reversing assessments from = %s forward, starting with ASMID = %d\n", n[0].Start.Format(rlib.RRDATEREPORTFMT), n[0].ASMID)
					//---------------------------------------------------------------------
					// This call reverses the supplied instances and all future instances
					//---------------------------------------------------------------------
					errlist := ReverseAssessment(ctx, &n[0], 1 /*this point forward*/, &dt, &x.LastClose)
					if len(errlist) > 0 {
						// rlib.Console("A9 error\n")
						return BizErrorListToError(errlist)
					}
					//-------------------------------------------------------------
					// PRORATE ASSESSMENT IF NEEDED
					// If the switchover (x.Ra.RentStart) date is NOT an instance
					// date of the the assessment (epoch = v.Start), then a prorated
					// payment to cover for the partial period is needed from the
					// instance epoch to the new rental agreement start
					//-------------------------------------------------------------
					// NOTE: may need to rethink this

					isinst := rlib.IsInstanceDate(&target, &x.Ra.RentStart, v.RentCycle, v.ProrationCycle)
					if !isinst {
						// rlib.Console("A9.1 - new RA rentstart (%s) was found NOT to be an instance date of ASMID = %d\n", x.Ra.RentStart.Format(rlib.RRDATEREPORTFMT), v.ASMID)
						//------------------------------------------------------
						// In this case we need to create a prorated assessment
						// that covers from target to x.Ra.RentStart
						//-----------------------------------------------------
						asm := n[0]
						amt, count, totcount := rlib.SimpleProrateAmount(v.Amount, v.RentCycle, v.ProrationCycle, &target, &x.Ra.RentStart, &target)
						thru := x.Ra.RentStart.Add(-rlib.CycleDuration(v.ProrationCycle, v.Start))
						// asm.AppendComment(fmt.Sprintf("prorated for %d of %d %s (covers %s thru %s)", count, totcount, rlib.ProrationUnits(v.ProrationCycle), target.Format(rlib.RRDATEFMT3), thru.Format(rlib.RRDATEFMT3)))
						asm.AppendComment(rlib.ProrateComment(count, totcount, v.ProrationCycle) + fmt.Sprintf(" (covers %s thru %s)", target.Format(rlib.RRDATEFMT3), thru.Format(rlib.RRDATEFMT3)))

						asm.Amount = amt
						asm.RentCycle = rlib.RECURNONE      // not part of a series
						asm.ProrationCycle = rlib.RECURNONE // no proration here
						asm.FLAGS = 0
						asm.Stop = asm.Start
						// rlib.Console("\n\n**********\ncalling InsertAssessment")
						if errlist := InsertAssessment(ctx, &asm, 0 /*no expanding*/, &x.LastClose); len(errlist) > 0 {
							return BizErrorListToError(errlist)
						}
						skipASMID = asm.ASMID
						// rlib.Console("A9.2 - just inserted asm = %d, skipASMID set\n", skipASMID)
						// rlib.Console("**********\n\n\n")
					} else {
						// rlib.Console("A9.3 - new RA rentstart (%s) was found to be an instance date of ASMID = %d\n", x.Ra.RentStart.Format(rlib.RRDATEREPORTFMT), v.ASMID)
						// rlib.Console("     - so will not add a prorated rent assessment\n")
					}
				}
				//-------------------------------------------------------------
				// Set the stop date for v to x.Ra.RentStart.  Since we've
				// already reversed only those instances that needed to be
				// reversed, we do not call the bizlogic version of this routine.
				// This is one of the rare exceptions where we just want to change
				// the end date and nothing else.
				//-------------------------------------------------------------
				v.Stop = x.Ra.RentStart
				if err = rlib.UpdateAssessment(ctx, &v); err != nil {
					return err
				}
			}
		}

		rlib.Console("A10 - HANDLE INSTANCES\n")
		//-----------------------------------------------------------------------
		// REVERSE ALL REMAINING INSTANCES IMPACTED BY THE NEW RENTAL AGREEMENT
		//-----------------------------------------------------------------------
		n, err = rlib.GetAssessmentInstancesByRAIDRange(ctx, ra.RAID, &x.Ra.RentStart, &rlib.ENDOFTIME)
		if err != nil {
			return err
		}
		rlib.Console("A11 -  Found %d instances for RAID %d in the range %s\n", len(n), ra.RAID, rlib.ConsoleDRange(&x.Ra.RentStart, &rlib.ENDOFTIME))
		for _, v := range n {
			if v.ASMID == skipASMID {
				continue // this one is OK, we just added it
			}
			rlib.Console("A12 -  ASMID = %d\n", v.ASMID)
			if v.FLAGS&(1<<2) != 0 {
				rlib.Console("A12.1 - reversed, skipping\n")
				continue // skip reversed assessments
			}
			if v.Start.Before(x.Ra.RentStart) {
				continue // v.Stop was >= x.Ra.RentStart so the overlap period matched the query, not a problem, just skip it
			}
			//---------------------------------------------------------------
			//  The date we use for the change depends on whether or not the
			//  financial period at the start date of the assessment has
			//  been closed.
			//---------------------------------------------------------------
			dt := v.Start // assume it will be on the assessment start date
			if v.Start.Before(x.LastClose.Dt) {
				dt = x.LastClose.Dt
			}
			rlib.Console("A12.2 - Reversal dates will be as of: %s\n", dt.Format(rlib.RRDATEREPORTFMT))
			//----------------------------
			// Now process the instance
			//----------------------------
			if !v.Start.Before(x.Ra.RentStart) {
				// Reverse the whole thing
				rlib.Console("A13 - Reversing ASMID = %d\n", v.ASMID)
				be := ReverseAssessment(ctx, &v, 0 /*this instance*/, &dt, &x.LastClose)
				if len(be) > 0 {
					rlib.Console("A13 error\n")
					PrintBizErrorList(be)
					return BizErrorListToError(be)
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
	//--------------------------------------------------------------------
	// There is one more edge case to check.  If the old RA ends in the
	// same rental period as the new RA begins then we need to handle
	// proration correctly.
	//--------------------------------------------------------------------
	// rlib.InstanceDateCoveringDate(epoch, t, cycle)
	// rlib.Console("x.Ra.Meta.RAID = %d\n", x.Raf.Meta.RAID)
	// for i := 0; i < len(x.RaChainOrig); i++ {
	// 	rlib.Console("raChain[%d] = RAID = %d, start/stop = %s\n", i, x.RaChainOrig[i].RAID, rlib.ConsoleDRange(&x.RaChainOrig[i].RentStart, &x.RaChainOrig[i].RentStop))
	//
	// }

	return nil
}
