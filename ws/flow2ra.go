package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"time"
)

// RAWriteHandler a handler function for part of the work of migrating
// RAFlow data into the permanent tables for a complete RentalAgreement
type RAWriteHandler struct {
	Name    string
	Handler func(context.Context, *rlib.F2RAWriteHandlerContext) error
}

// UpdateHandlers is the collection of routines to call to write a flow
// for an existing Rental Agreement back to the database as a RentalAgreement.
var ehandlers = []RAWriteHandler{
	{"Transactants", F2RAUpdatePeople},
	{"Pets", F2RAUpdatePets},
	{"Vehicles", F2RAUpdateVehicles},
	{"Rentables", FlowSaveRentables},
	{"Fees", bizlogic.Fees2RA},
}

// Flow2RA moves data from the Flow tabl into the permanent tables.
//
// INPUTS
//     ctx    - db context for transactions
//     flowid - the flow data to migrate into the permanent tables
//
// RETURNS
//     RAID of the newly created RentalAgreement
//     Any errors encountered
//-----------------------------------------------------------------------------
func Flow2RA(ctx context.Context, flowid int64) (int64, error) {
	rlib.Console("Entered Flow2RA\n")
	var x rlib.F2RAWriteHandlerContext
	var nraid int64

	//-------------------------------------------
	// Read the flow data into a data structure
	//-------------------------------------------
	flow, err := rlib.GetFlow(ctx, flowid)
	if err != nil {
		rlib.Console("\n\nERROR IN GetFlow: %s\n\n\n", err.Error())
		return nraid, err
	}
	err = json.Unmarshal(flow.Data, &x.Raf)
	if err != nil {
		rlib.Console("\n\nERROR IN Unmarshal: %s\n\n\n", err.Error())
		return nraid, err
	}

	//----------------------------------------------------------------------------
	// The flow datastruct is not updated with EDI correction.  Do those corrections
	// here...
	//----------------------------------------------------------------------------
	rlib.EDIHandleIncomingJSONDateRange(x.Raf.Meta.BID, &x.Raf.Dates.AgreementStart, &x.Raf.Dates.AgreementStop)
	rlib.EDIHandleIncomingJSONDateRange(x.Raf.Meta.BID, &x.Raf.Dates.RentStart, &x.Raf.Dates.RentStop)
	rlib.EDIHandleIncomingJSONDateRange(x.Raf.Meta.BID, &x.Raf.Dates.PossessionStart, &x.Raf.Dates.PossessionStop)

	//----------------------------------------------------------------------------
	// If this is an update of an existing RAID, check to see if any changes
	// were made. Otherwise treat it as a new RAID
	//----------------------------------------------------------------------------
	x.IsNewOriginRaid = x.Raf.Meta.RAID == 0
	rlib.Console("isNewOriginRaid = %t\n", x.IsNewOriginRaid)

	//---------------------------------------------
	//  UPDATE AN EXISTING RENTAL AGREEMENT...
	//---------------------------------------------
	if !x.IsNewOriginRaid { // update existing
		changes, err := rlib.RAFlowDataDiff(ctx, x.Raf.Meta.RAID)
		if err != nil {
			rlib.Console("\n\nERROR IN FlowDataDIFF: %s\n\n\n", err.Error())
			return nraid, err
		}
		rlib.Console("\tData changes found = %t\n", changes)

		//-----------------------------------------------------------------------
		// If there were changes to the data, create an amended Rental Agreement
		//-----------------------------------------------------------------------
		if changes {
			x.NewRAID, err = FlowSaveRA(ctx, &x) // this will update x.NewRAID with the new raid
			if err != nil {
				rlib.Console("\n\nERROR IN FlowSaveRA: %s\n\n\n", err.Error())
				return x.NewRAID, err
			}
		} else {
			err = fmt.Errorf("there are no data changes")
			return x.NewRAID, err
		}

		//------------------------------------------------------------
		// if there are meta data changes, then updated existing RAID
		//------------------------------------------------------------
		// rlib.Console("Just before call to FlowSaveMetaDataChanges: nraid = %d, x.NewRAID = %d\n", nraid, x.NewRAID)
		nraid, err = FlowSaveMetaDataChanges(ctx, &x)
		if err != nil {
			rlib.Console("\n\nERROR IN FlowSaveMetaDataChanges: %s\n\n\n", err.Error())
			return nraid, err
		}
		// rlib.Console("\tMetaData data updated on RAID=%d\n", nraid)

	} else {
		//---------------------------------------------
		//  CREATE A NEW RENTAL AGREEMENT...
		//---------------------------------------------
		nraid, err = FlowSaveRA(ctx, &x)
		if err != nil {
			rlib.Console("\n\nERROR IN FlowSaveRA: %s\n\n\n", err.Error())
			return nraid, err
		}
		// rlib.Console("New ORIGIN = %d\n", nraid)
		x.NewRAID = nraid
		nraid, err = FlowSaveMetaDataChanges(ctx, &x)
		if err != nil {
			rlib.Console("\n\nERROR IN FlowSaveMetaDataChanges: %s\n\n\n", err.Error())
			return nraid, err
		}
		rlib.Console("\tMetaData data updated on RAID=%d\n", nraid)
	}

	// REMOVE FLOW IF MIGRATION DONE SUCCESSFULLY
	// Delete only if state is active or above active
	var state = x.Ra.FLAGS & uint64(0xF)
	if state >= 4 && state <= 6 {
		err = rlib.DeleteFlow(ctx, flowid)
		if err != nil {
			return nraid, err
		}
	}

	// rlib.Console("\tx.OldRAID = %d, x.NewRAID = %d\n", x.OldRAID, x.NewRAID)
	return x.NewRAID, nil
}

// FlowSaveMetaDataChanges saves any change to the meta data in the flow with
//     the existing RAID.  In this case, we must always set the state of any
//     old RAID to Terminated because it is being replaced by an amended
//     agreement.
//
// INPUTS
//     ctx - db context for transactions
//     x - all the contextual info we need for performing this operation
//         Note: this routine adds ra and raChainOrig to x
//
// RETURNS
//     RAID of the Rental Agreement in which meta-data was changed.
//     Any errors encountered
//-----------------------------------------------------------------------------
func FlowSaveMetaDataChanges(ctx context.Context, x *rlib.F2RAWriteHandlerContext) (int64, error) {
	var err error
	rlib.Console("Entered FlowSaveMetaDataChanges\n")
	raid := x.NewRAID // update this one if changes were found and a new amendment was written.
	if raid == 0 {
		raid = x.Raf.Meta.RAID // update this one if no changes were found
	}
	x.Ra, err = rlib.GetRentalAgreement(ctx, raid)
	if err != nil {
		rlib.Ulog("Could not read rental agreement %d, err: %s\n", raid, err.Error())
		return raid, err
	}
	x.NewRAID = raid // we'll update this one

	//----------------------------------------------------
	// compare the meta data and update if necessary
	//----------------------------------------------------
	changes := 0
	bterminated := x.Ra.FLAGS&0xf == rlib.RASTATETerminated
	if x.Ra.FLAGS != x.Raf.Meta.RAFLAGS {
		//---------------------------------------------------------------------
		// If the FLAGs have changed, check to see if state of the permanent
		// table copy is in the Terminated state. If it is, do not change it
		// or the reason...
		//---------------------------------------------------------------------
		if bterminated {
			// DO NOTHING IF RA IS ALREADY TERMINATED
			// flags := x.Ra.FLAGS
			// if (x.Ra.FLAGS & ^uint64(0xf)) != (x.Raf.Meta.RAFLAGS & ^uint64(0xf)) { // flags other than
			// 	state := x.Ra.FLAGS & 0xf
			// 	x.Ra.FLAGS = x.Raf.Meta.RAFLAGS
			// 	x.Ra.FLAGS &= ^unit64(0xf)
			// 	x.Ra.FLAGS |= state
			// 	changes++
			// }
		} else {
			x.Ra.FLAGS = x.Raf.Meta.RAFLAGS
			changes++
		}
	}
	if x.Ra.ApplicationReadyUID != x.Raf.Meta.ApplicationReadyUID {
		changes++
		x.Ra.ApplicationReadyUID = x.Raf.Meta.ApplicationReadyUID
	}
	if !x.Ra.ApplicationReadyDate.Equal(time.Time(x.Raf.Meta.ApplicationReadyDate)) {
		changes++
		x.Ra.ApplicationReadyDate = time.Time(x.Raf.Meta.ApplicationReadyDate)
	}
	if x.Ra.Approver1 != x.Raf.Meta.Approver1 {
		changes++
		x.Ra.Approver1 = x.Raf.Meta.Approver1
	}
	if !x.Ra.DecisionDate1.Equal(time.Time(x.Raf.Meta.DecisionDate1)) {
		changes++
		x.Ra.DecisionDate1 = time.Time(x.Raf.Meta.DecisionDate1)
	}
	if x.Ra.DeclineReason1 != x.Raf.Meta.DeclineReason1 {
		changes++
		x.Ra.DeclineReason1 = x.Raf.Meta.DeclineReason1
	}
	if x.Ra.Approver2 != x.Raf.Meta.Approver2 {
		changes++
		x.Ra.Approver2 = x.Raf.Meta.Approver2
	}
	if !x.Ra.DecisionDate2.Equal(time.Time(x.Raf.Meta.DecisionDate2)) {
		changes++
		x.Ra.DecisionDate2 = time.Time(x.Raf.Meta.DecisionDate2)
	}
	if x.Ra.DeclineReason2 != x.Raf.Meta.DeclineReason2 {
		changes++
		x.Ra.DeclineReason2 = x.Raf.Meta.DeclineReason2
	}
	if x.Ra.MoveInUID != x.Raf.Meta.MoveInUID {
		changes++
		x.Ra.MoveInUID = x.Raf.Meta.MoveInUID
	}
	if !x.Ra.MoveInDate.Equal(time.Time(x.Raf.Meta.MoveInDate)) {
		changes++
		x.Ra.MoveInDate = time.Time(x.Raf.Meta.MoveInDate)
	}
	if x.Ra.ActiveUID != x.Raf.Meta.ActiveUID {
		changes++
		x.Ra.ActiveUID = x.Raf.Meta.ActiveUID
	}
	if !x.Ra.ActiveDate.Equal(time.Time(x.Raf.Meta.ActiveDate)) {
		changes++
		x.Ra.ActiveDate = time.Time(x.Raf.Meta.ActiveDate)
	}
	if x.Ra.TerminatorUID != x.Raf.Meta.TerminatorUID {
		changes++
		x.Ra.TerminatorUID = x.Raf.Meta.TerminatorUID
	}
	if !x.Ra.TerminationDate.Equal(time.Time(x.Raf.Meta.TerminationDate)) {
		changes++
		x.Ra.TerminationDate = time.Time(x.Raf.Meta.TerminationDate)
	}
	if x.Ra.LeaseTerminationReason != x.Raf.Meta.LeaseTerminationReason {
		changes++
		x.Ra.LeaseTerminationReason = x.Raf.Meta.LeaseTerminationReason
	}
	if !x.Ra.DocumentDate.Equal(time.Time(x.Raf.Meta.DocumentDate)) {
		changes++
		x.Ra.DocumentDate = time.Time(x.Raf.Meta.DocumentDate)
	}
	if x.Ra.NoticeToMoveUID != x.Raf.Meta.NoticeToMoveUID {
		changes++
		x.Ra.NoticeToMoveUID = x.Raf.Meta.NoticeToMoveUID
	}
	if !x.Ra.NoticeToMoveDate.Equal(time.Time(x.Raf.Meta.NoticeToMoveDate)) {
		changes++
		x.Ra.NoticeToMoveDate = time.Time(x.Raf.Meta.NoticeToMoveDate)
	}
	if !x.Ra.NoticeToMoveReported.Equal(time.Time(x.Raf.Meta.NoticeToMoveReported)) {
		changes++
		x.Ra.NoticeToMoveReported = time.Time(x.Raf.Meta.NoticeToMoveReported)
	}

	//---------------------------------------------------------
	// If there were any changes, update the Rental Agreement
	//---------------------------------------------------------
	rlib.Console("Metadata change count = %d\n", changes)
	if changes > 0 {
		rlib.Console("Updating RAID = %d\n", x.Ra.RAID)
		err = rlib.UpdateRentalAgreement(ctx, &x.Ra)
		if err != nil {
			return x.Raf.Meta.RAID, err
		}
	}

	return x.NewRAID, err
}

// FlowSaveRA saves a new Rental Agreement from the supplied flow. This
//     function assumes that a check has already been made to verify that
//     the RentalAgreement is either new or, if it is replacing an existing
//     rental agreement, that the data has actually been changed.
//
// INPUTS
//     ctx - db context for transactions
//     x - all the contextual info we need for performing this operation
//         Note: this routine adds ra and raChainOrig to x
//
// RETURNS
//     RAID of newly created Rental Agreement or updated Rental Agreement
//         if only meta-data was changed.
//     Any errors encountered
//-----------------------------------------------------------------------------
func FlowSaveRA(ctx context.Context, x *rlib.F2RAWriteHandlerContext) (int64, error) {
	rlib.Console("Entered FlowSaveRA\n")
	// rlib.Console("x.Ra.RentStop = %s, x.Ra.PossessionStop = %s\n", x.Ra.RentStop, x.Ra.PossessionStop)
	// rlib.Console("x.Raf.Dates.RentStop, x.Raf.Dates.PossessionStop = %s\n", rlib.ConsoleJSONDRange(&x.Raf.Dates.RentStop, &x.Raf.Dates.PossessionStop))
	var err error
	var nraid int64
	var raOrig rlib.RentalAgreement

	if err = rlib.InitBizInternals(x.Raf.Meta.BID, &x.Xbiz); err != nil {
		return nraid, err
	}

	rlib.Console("A x.Raf.Dates.Term %s\n", rlib.ConsoleJSONDRange(&x.Raf.Dates.AgreementStart, &x.Raf.Dates.AgreementStop))
	rlib.Console("x.Raf.Meta.RAID = %d\n", x.Raf.Meta.RAID)
	if x.Raf.Meta.RAID > 0 {
		rlib.Console("A1\n")
		//------------------------------------------------------------
		// Get the rental agreement chain that will be updated by the
		// one we're creating here. Update its stop dates accordingly
		//------------------------------------------------------------
		x.OldRAID = x.Raf.Meta.RAID
		raOrig, err = rlib.GetRentalAgreement(ctx, x.OldRAID)
		if err != nil {
			return nraid, err
		}
		if raOrig.ORIGIN == int64(0) {
			x.RaChainOrig = append(x.RaChainOrig, raOrig)
		} else {
			x.RaChainOrig, err = rlib.GetRentalAgreementChain(ctx, raOrig.ORIGIN)
			if err != nil {
				return nraid, err
			}
		}
		for i := 0; i < len(x.RaChainOrig); i++ {
			rlib.Console("RaChainOrigin[%d] = RAID-%d\n", i, x.RaChainOrig[i].RAID)
		}
		x.RaChainOrigUnchanged = make([]rlib.RentalAgreement, len(x.RaChainOrig))
		copy(x.RaChainOrigUnchanged, x.RaChainOrig)

		//--------------------------------------------------
		// Now press forward using the adjusted dates...
		//--------------------------------------------------
		AStart := time.Time(x.Raf.Dates.AgreementStart)
		RStart := time.Time(x.Raf.Dates.RentStart)
		PStart := time.Time(x.Raf.Dates.PossessionStart)

		AStop := time.Time(x.Raf.Dates.AgreementStop)
		RStop := time.Time(x.Raf.Dates.RentStop)
		PStop := time.Time(x.Raf.Dates.PossessionStop)
		now := rlib.Now() // can only override system time during testing
		rlib.Console("AStart/AStop = %s\n", rlib.ConsoleDRange(&AStart, &AStop))
		rlib.Console("rlib.Now() = %s\n", rlib.ConDt(&now))

		chgs := 0
		x.RaOrigIndex = -1 // mark that there is nothing to link to at the moment

		//------------------------------------------------------------------
		//  Fix up the dates of the affected Rental Agreements. We only
		//  want to change the stop dates of ACTIVE Rental Agreements,
		//  and only if its stop dates are AFTER the start date of the
		//  amended RA
		//------------------------------------------------------------------
		for i := 0; i < len(x.RaChainOrig); i++ {
			rlib.Console("RAChain[%d] = RAID %d. Date mods...\n", i, x.RaChainOrig[i].RAID)
			state := x.RaChainOrig[i].FLAGS & 0xf
			if !(state == rlib.RASTATEActive || state == rlib.RASTATENoticeToMove) {
				rlib.Console("state indicates not active, skipping\n")
				continue // we're only interested in active RAs
			}

			x.RaOrigIndex = i // keep track of the RA currently active.

			//------------------------------------------------------------------
			//  The not-so-obvious change is if the new RentalAgreement has a
			//  timespan that is disjoint from the existing rental agreement AND
			//  the existing rental agreement stop time is in the future (i.e.,
			//  after the system "now" date/time).  We need to make sure that we
			//  properly adjust RentableLeaseStatus if the old and new RAs are
			//  disjoint.  We also need to adjust if the old RA had a stop date
			//  in the future and we need to trim its stop date back to "now".
			//
			//  scenario 1:    ###### new ######|<---gap--->|**** old ****
			//  scenario 2:    ###### old ######|<---gap--->|**** new ****
			//
			//  modifier:                  now
			//                              |
			//                 ###### old ######
			//------------------------------------------------------------------
			if x.RaChainOrig[i].AgreementStop.After(AStart) && AStop.After(x.RaChainOrig[i].AgreementStart) {
				rlib.Console("\t(a.) setting AgreementStop to %s\n", AStart.Format(rlib.RRDATEFMT3))
				x.RaChainOrig[i].AgreementStop = AStart
				chgs++
			}
			if x.RaChainOrig[i].AgreementStop.After(now) && !AStart.Equal(x.RaChainOrig[i].AgreementStop) { // not adjacent check
				rlib.Console("\t(b.) setting AgreementStop to %s\n", now.Format(rlib.RRDATEFMT3))
				x.RaChainOrig[i].AgreementStop = now
				chgs++
			}

			if x.RaChainOrig[i].RentStop.After(RStart) && RStop.After(x.RaChainOrig[i].RentStart) {
				rlib.Console("\t(a.) setting RentStop to %s\n", RStart.Format(rlib.RRDATEFMT3))
				x.RaChainOrig[i].RentStop = RStart
				chgs++
			}
			if x.RaChainOrig[i].RentStop.After(now) && !PStart.Equal(x.RaChainOrig[i].RentStop) { // not adjacent check
				rlib.Console("\t(b.) setting RentStop to %s\n", now.Format(rlib.RRDATEFMT3))
				x.RaChainOrig[i].RentStop = now
				chgs++
			}

			if x.RaChainOrig[i].PossessionStop.After(PStart) && PStop.After(x.RaChainOrig[i].PossessionStart) {
				rlib.Console("\t(a.) setting PossessionStop to %s\n", PStart.Format(rlib.RRDATEFMT3))
				x.RaChainOrig[i].PossessionStop = PStart
				chgs++
			}
			if x.RaChainOrig[i].PossessionStop.After(now) && !PStart.Equal(x.RaChainOrig[i].PossessionStop) { // not adjacent check
				rlib.Console("\t(b.) setting PossessionStop to %s\n", now.Format(rlib.RRDATEFMT3))
				x.RaChainOrig[i].PossessionStop = now
				chgs++
			}

			p1 := rlib.Earliest(&x.RaChainOrig[i].PossessionStart, &x.RaChainOrig[i].RentStart)
			p2 := rlib.Latest(&x.RaChainOrig[i].PossessionStop, &x.RaChainOrig[i].RentStop)
			disjoint := p1.After(PStop) || p2.Before(PStart)

			//---------------------------------------------------------------------------------
			// If the periods are disjoint, then we need to update the RentableLeaseStatus
			// of all Rentables in the RentalAgreement
			//---------------------------------------------------------------------------------
			rlib.Console("disjoint = %t, PossessionStop = %s\n", disjoint, rlib.ConDt(&x.RaChainOrig[i].PossessionStop))
			if disjoint && x.RaChainOrigUnchanged[i].PossessionStop.After(now) {
				rlib.Console("\n\n***** ADJUSTING LEASE STATUS *****\n\n")
				d1 := x.RaChainOrig[i].PossessionStart
				if d1.Before(now) {
					d1 = now
				}
				// Find all rentables in the old rental agreement
				rarl, err := rlib.GetAllRentalAgreementRentables(ctx, x.Raf.Meta.RAID)
				if err != nil {
					return nraid, err
				}
				for _, v := range rarl {
					var l = rlib.RentableLeaseStatus{
						BID:         x.RaChainOrig[i].BID,
						DtStart:     d1,
						DtStop:      rlib.ENDOFTIME,
						LeaseStatus: rlib.LEASESTATUSnotleased,
						RID:         v.RID,
					}
					q := fmt.Sprintf("select %s from RentableLeaseStatus where RID=%d AND BID=%d AND DtStart >= %q ORDER BY DtStart LIMIT 1;",
						rlib.RRdb.DBFields["RentableLeaseStatus"], v.RID, v.BID, x.RaChainOrigUnchanged[i].PossessionStop.Format(rlib.RRDATEFMTSQL))
					row := rlib.RRdb.Dbrr.QueryRow(q)
					var lsNext rlib.RentableLeaseStatus
					if err = rlib.ReadRentableLeaseStatus(row, &lsNext); err != nil {
						return nraid, err
					}
					rlib.Console("Found RentableLeaseStatus: %d  %s\n", lsNext.LeaseStatus, rlib.ConsoleDRange(&lsNext.DtStart, &lsNext.DtStop))

					if lsNext.RLID > 0 {
						futurelimit := now.AddDate(0, rlib.FUTURERESLIMIT, 0)
						if x.RaChainOrigUnchanged[i].PossessionStop.Equal(lsNext.DtStart) && lsNext.DtStop.After(futurelimit) {
							//---------------------------------------------------------------
							// If the next lease status is adjacent to the old RA's DtStop,
							// then just extend.
							//---------------------------------------------------------------
							l = lsNext
							l.DtStart = d1
							l.LeaseStatus = rlib.LEASESTATUSnotleased
						} else {
							//---------------------------------------------------------------
							// If it is disjoint, then terminate the
							// new one at the beginning of the record just found...
							//---------------------------------------------------------------
							l.DtStop = lsNext.DtStart
						}
						if err = rlib.SetRentableLeaseStatus(ctx, &l, false /*purge the 3rd arg asap*/); err != nil {
							return nraid, err
						}
					}
				}
			}

			// //----------------------------------------------------------------
			// // DEPRECATE THIS CODE AS SOON AS THE ABOVE CODE PASSES ALL TESTS
			// //----------------------------------------------------------------
			// if x.RaChainOrig[i].AgreementStop.After(AStart) {
			// 	// rlib.Console("\tsetting AgreementStop to %s\n", AStart.Format(rlib.RRDATEFMT3))
			// 	x.RaChainOrig[i].AgreementStop = AStart
			// 	chgs++
			// }
			// if x.RaChainOrig[i].RentStop.After(RStart) {
			// 	// rlib.Console("\tsetting RentStop to %s\n", RStart.Format(rlib.RRDATEFMT3))
			// 	x.RaChainOrig[i].RentStop = RStart
			// 	chgs++
			// }
			// if x.RaChainOrig[i].PossessionStop.After(PStart) {
			// 	// rlib.Console("\tsetting PossessionStop to %s\n", PStart.Format(rlib.RRDATEFMT3))
			// 	x.RaChainOrig[i].PossessionStop = PStart
			// 	chgs++
			// }

			//------------------------------------------------------------------
			// If there are changes, then we stop the old Rental Agreement and
			// create a new one linked to x.RaChainOrig[i].
			//------------------------------------------------------------------
			if chgs > 0 {
				x.RaChainOrig[i].FLAGS &= ^uint64(0xf)           // clear the status
				x.RaChainOrig[i].FLAGS |= rlib.RAActionTerminate // set the state to Terminated
				x.RaChainOrig[i].LeaseTerminationReason =
					rlib.RRdb.BizTypes[x.RaChainOrig[i].BID].Msgs.S[rlib.MSGRAUPDATED].SLSID // "Rental Agreement was updated"

				if err = setRATerminator(ctx, &x.RaChainOrig[i]); err != nil {
					return nraid, err
				}

				rlib.Console("Updating RAID %d.  AgreementStart = %s, AgreementStop = %s\n", x.RaChainOrig[i].RAID, x.RaChainOrig[i].AgreementStart.Format(rlib.RRDATEFMT3), x.RaChainOrig[i].AgreementStop.Format(rlib.RRDATEFMT3))
				err = rlib.UpdateRentalAgreement(ctx, &x.RaChainOrig[i])
				if err != nil {
					return nraid, err
				}
			}
		}

		//---------------------------------------------------------------------
		// if x.RaOriginIndex has not yet been set, set it to the last RA in
		// the chain chronologically.  The chain is ordered by AgreementStart
		// so, we'll link it to the last one in the chain...
		//---------------------------------------------------------------------
		if x.RaOrigIndex < 0 {
			x.RaOrigIndex = len(x.RaChainOrig) - 1
		}

		//------------------------------------------------------------
		// Now start the new RAID.  Link it to x.RaChainOrig[i]
		//------------------------------------------------------------
		initRA(ctx, x)
		// rlib.Console("B0: After call to initRA: x.Ra.RentStop = %s, x.Ra.PossessionStop = %s\n", x.Ra.RentStop, x.Ra.PossessionStop)

		i := x.RaOrigIndex // makes it easier to read the following lines
		rlib.Console("After updates x.RaOrigIndex = %d  -> RAID = %d.  x.RaChainOrig[i]: AgreementStart = %s, AgreementStop = %s\n", i, x.RaChainOrig[i].RAID, x.RaChainOrig[i].AgreementStart.Format(rlib.RRDATEFMT3), x.RaChainOrig[i].AgreementStop.Format(rlib.RRDATEFMT3))
		rlib.Console("x.RaChainOrigUnchanged[i]: AgreementStart = %s, AgreementStop = %s\n", x.RaChainOrigUnchanged[i].AgreementStart.Format(rlib.RRDATEFMT3), x.RaChainOrigUnchanged[i].AgreementStop.Format(rlib.RRDATEFMT3))
		x.Ra.PRAID = x.RaChainOrig[i].RAID
		x.Ra.ORIGIN = x.RaChainOrig[i].ORIGIN
		x.Ra.BID = x.RaChainOrig[i].BID
		if x.RaChainOrig[i].ORIGIN == 0 {
			x.Ra.ORIGIN = x.RaChainOrig[i].RAID
		}
		x.Ra.RATID = x.RaChainOrig[i].RATID
		x.Ra.RentCycleEpoch = x.RaChainOrig[i].RentCycleEpoch

	} else {
		rlib.Console("B1\n")
		//-------------------------------------
		// This is a new Rental Agreement...
		//-------------------------------------
		initRA(ctx, x)
		// rlib.Console("B1: After call to initRA: x.Ra.RentStop = %s, x.Ra.PossessionStop = %s\n", x.Ra.RentStop, x.Ra.PossessionStop)
	}

	nraid, err = rlib.InsertRentalAgreement(ctx, &x.Ra)
	if err != nil {
		return nraid, err
	}
	x.NewRAID = nraid
	//-----------------------------------------------------
	// Create a RentalAgreement Ledger marker
	//-----------------------------------------------------
	var lm = rlib.LedgerMarker{
		BID:     x.Ra.BID,
		RAID:    x.NewRAID,
		RID:     0,
		Dt:      x.Ra.AgreementStart,
		Balance: float64(0),
		State:   rlib.LMINITIAL,
	}
	_, err = rlib.InsertLedgerMarker(ctx, &lm)
	if err != nil {
		return nraid, err
	}

	//---------------------------------------------------------------
	// Now spin through the series of handlers that move the data
	// into the permanent tables...
	//---------------------------------------------------------------
	for i := 0; i < len(ehandlers); i++ {
		rlib.Console("FlowSaveRA: running handler %s\n", ehandlers[i].Name)
		// rlib.Console("before: x.Ra.RentStop = %s, x.Ra.PossessionStop = %s\n", x.Ra.RentStop, x.Ra.PossessionStop)
		if err = ehandlers[i].Handler(ctx, x); err != nil {
			rlib.Console("error returned from handler %s: %s\n", ehandlers[i].Name, err.Error())
			return nraid, err
		}
	}

	//-------------------------------------------------------------------------
	// Final Step:  The raid that had no parent gets x.Ra as its parent. We
	// must use x.RaChainOrig because its agreement dates may have been
	// modified and we don't want to lose that.
	//-------------------------------------------------------------------------
	// rlib.Console("FINAL STEP\n")
	for i := 0; i < len(x.RaChainOrig); i++ {
		// rlib.Console("Checking RAID %d\n", x.RaChainOrig[i].RAID)
		if x.RaChainOrig[i].PRAID == 0 {
			//---------------------------------------------------------------
			// One last check before updating... if this RAID's State is not
			// Terminated, then we need to terminate it and set the reason
			//---------------------------------------------------------------
			if x.RaChainOrig[i].FLAGS&0xf != rlib.RASTATETerminated {
				x.RaChainOrig[i].FLAGS &= ^uint64(0xf)
				x.RaChainOrig[i].FLAGS |= rlib.RASTATETerminated
				x.RaChainOrig[i].LeaseTerminationReason = rlib.RRdb.BizTypes[x.Ra.BID].Msgs.S[rlib.MSGRAUPDATED].SLSID // "Rental Agreement was updated"
				if err = setRATerminator(ctx, &x.RaChainOrig[i]); err != nil {
					return nraid, err
				}
			}
			// rlib.Console("UPDATING x.RaChainOrig[i]\n")
			if err = rlib.UpdateRentalAgreement(ctx, &x.RaChainOrig[i]); err != nil {
				return nraid, err
			}
		}
	}
	// rlib.Console("DONE\n")

	return nraid, nil
}

func setRATerminator(ctx context.Context, ra *rlib.RentalAgreement) error {
	//--------------------------------------------------------------------------
	// In noauth mode, we still have tester session, get it from the context
	//--------------------------------------------------------------------------
	sess, ok := rlib.SessionFromContext(ctx)
	if !ok {
		return rlib.ErrSessionRequired
	}

	ra.TerminatorUID = sess.UID
	// ra.TerminationDate = time.Now()
	ra.TerminationDate = rlib.Now()
	return nil
}

// initRA initializes a rental agreement structure with information from flow
// data. upon completion, x.Ra will be filled out with basic information that
// can be pulled from x.Raf
//
// INPUTS
//     ctx - db context for transactions
//     x - all the contextual info we need for performing this operation
//         Note: this routine adds ra and raOrig to x
//
// RETURNS
//     nothing at this time
//-----------------------------------------------------------------------------
func initRA(ctx context.Context, x *rlib.F2RAWriteHandlerContext) {
	//-------------------------------------
	// Adjust dates for EDI...
	//-------------------------------------
	AStart := time.Time(x.Raf.Dates.AgreementStart)
	RStart := time.Time(x.Raf.Dates.RentStart)
	PStart := time.Time(x.Raf.Dates.PossessionStart)
	AStop := time.Time(x.Raf.Dates.AgreementStop)
	RStop := time.Time(x.Raf.Dates.RentStop)
	PStop := time.Time(x.Raf.Dates.PossessionStop)

	rlib.Console("initRA: Agreement start/stop = %s\n", rlib.ConsoleJSONDRange(&x.Raf.Dates.AgreementStart, &x.Raf.Dates.AgreementStop))

	// This is handled much earlier now
	// rlib.EDIHandleIncomingDateRange(x.Raf.Meta.BID, &AStart, &AStop)
	// rlib.EDIHandleIncomingDateRange(x.Raf.Meta.BID, &RStart, &RStop)
	// rlib.EDIHandleIncomingDateRange(x.Raf.Meta.BID, &PStart, &PStop)

	x.Ra.PRAID = int64(0)
	x.Ra.ORIGIN = int64(0)
	x.Ra.BID = x.Raf.Meta.BID
	x.Ra.AgreementStart = AStart
	x.Ra.AgreementStop = AStop
	x.Ra.RentStart = RStart
	x.Ra.RentStop = RStop
	x.Ra.PossessionStart = PStart
	x.Ra.PossessionStop = PStop
	x.Ra.CSAgent = x.Raf.Dates.CSAgent
	x.Ra.FLAGS = x.Raf.Meta.RAFLAGS
	x.Ra.Approver1 = x.Raf.Meta.Approver1
	x.Ra.DeclineReason1 = x.Raf.Meta.DeclineReason1
	x.Ra.DecisionDate1 = time.Time(x.Raf.Meta.DecisionDate1)
	x.Ra.Approver2 = x.Raf.Meta.Approver2
	x.Ra.DeclineReason2 = x.Raf.Meta.DeclineReason2
	x.Ra.DecisionDate2 = time.Time(x.Raf.Meta.DecisionDate2)
	x.Ra.CSAgent = x.Raf.Dates.CSAgent
	x.Ra.NoticeToMoveDate = time.Time(x.Raf.Meta.NoticeToMoveDate)
	x.Ra.NoticeToMoveReported = time.Time(x.Raf.Meta.NoticeToMoveReported)
	x.Ra.TerminatorUID = x.Raf.Meta.TerminatorUID
	x.Ra.TerminationDate = time.Time(x.Raf.Meta.TerminationDate)
	// x.Ra.FollowUpDate = time.Time(x.Raf.Meta.FollowUpDate)
	// x.Ra.Outcome = x.Raf.Meta.Outcome
	// x.Ra.NoticeToMoveUID = x.Raf.Meta.NoticeToMoveUID
	// x.Ra.OtherPreferences = x.Raf.Meta.OtherPreferences
}

// FlowSaveRentables adds/updates rentables from the flow data.  This means
// that we update or add the RentalAgreementRentables list.  Update means
// that we set the stop date for the existing RentalAgreementRentables RAID.
// Then we add the Rentables in x.Raf.Rentables[] into a
// RentalAgreementRentables for the new RAID
//
// INPUTS
//     ctx - db context for transactions
//     x - all the contextual info we need for performing this operation
//         Note: this routine adds ra and raOrig to x
//
// RETURNS
//     RAID of newly created Rental Agreement
//     Any errors encountered
//-----------------------------------------------------------------------------
func FlowSaveRentables(ctx context.Context, x *rlib.F2RAWriteHandlerContext) error {
	rlib.Console("Entered FlowSaveRentables: x.Ra.RentStop = %s, x.Ra.PossessionStop = %s\n", rlib.ConDt(&x.Ra.RentStop), rlib.ConDt(&x.Ra.PossessionStop))
	//----------------------------------------------------------------
	// Update the stop date on any existing RentalAgreementRentables
	//----------------------------------------------------------------
	if x.Raf.Meta.RAID > 0 {
		rarl, err := rlib.GetAllRentalAgreementRentables(ctx, x.Raf.Meta.RAID)
		if err != nil {
			return err
		}
		for _, v := range rarl {
			v.RARDtStop = time.Time(x.Raf.Dates.AgreementStart)
			if err = rlib.UpdateRentalAgreementRentable(ctx, &v); err != nil {
				return err
			}
			//----------------------------------------------------------------
			// Fix up the users
			//----------------------------------------------------------------
			rul, err := rlib.GetRentableUsersInRange(ctx, v.RID, &x.RaChainOrig[x.RaOrigIndex].PossessionStart, &x.Ra.PossessionStop)
			if err != nil {
				return err
			}
			for _, ru := range rul {
				ru.DtStop = x.RaChainOrig[x.RaOrigIndex].PossessionStop
				if err = rlib.UpdateRentableUser(ctx, &ru); err != nil {
					return err
				}
			}
		}
		//----------------------------------------------------------------
		// Fix up the payors
		//----------------------------------------------------------------
		t, err := rlib.GetRentalAgreementPayorsByRAID(ctx, x.Raf.Meta.RAID)
		if err != nil {
			return err
		}
		for _, v := range t {
			if v.DtStop.After(x.Ra.RentStart) {
				v.DtStop = x.Ra.RentStart
				if err = rlib.UpdateRentalAgreementPayor(ctx, &v); err != nil {
					return err
				}
			}
		}
	}

	//----------------------------------------------------------------
	// Set the range of time to show the rentable as leased...
	//----------------------------------------------------------------
	rlib.Console("FlowSaveRentables: x.Ra.RentStop = %s, x.Ra.PossessionStop = %s\n", rlib.ConDt(&x.Ra.RentStop), rlib.ConDt(&x.Ra.PossessionStop))
	d1 := x.Ra.RentStart
	d2 := x.Ra.RentStop

	if x.Ra.PossessionStart.Before(d1) {
		d1 = x.Ra.PossessionStart
	}
	if x.Ra.PossessionStop.Before(d2) {
		d2 = x.Ra.PossessionStop
	}
	//----------------------------------------------------------------
	// Add a RentalAgreementRentable entry for each Rentable
	//----------------------------------------------------------------
	for _, v := range x.Raf.Rentables {
		var rar = rlib.RentalAgreementRentable{
			RAID:         x.Ra.RAID,
			BID:          x.Ra.BID,
			RID:          v.RID,
			CLID:         0, // TODO:
			ContractRent: 0, // TODO:
			RARDtStart:   time.Time(x.Raf.Dates.PossessionStart),
			RARDtStop:    time.Time(x.Raf.Dates.PossessionStop),
		}
		_, err := rlib.InsertRentalAgreementRentable(ctx, &rar)
		if err != nil {
			return err
		}

		//----------------------------------------------------------------
		// Mark this rentable as leased...
		//----------------------------------------------------------------
		var rls = rlib.RentableLeaseStatus{
			DtStart:     d1,
			DtStop:      d2,
			BID:         x.Ra.BID,
			RID:         v.RID,
			LeaseStatus: rlib.LEASESTATUSleased,
		}
		rlib.Console("FlowSaveRentables: calling SetRentableLeaseStatus\n")
		rlib.Console("\n\n********************\nSetRentableLeaseStatus:  %s  lease status = %d\n", rlib.ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.LeaseStatus)
		if err = rlib.SetRentableLeaseStatus(ctx, &rls, true); err != nil {
			return err
		}
		rlib.Console("**********************\n\n")

		//----------------------------------------------------------------
		// Add the users
		//----------------------------------------------------------------
		for _, v1 := range x.Raf.People {
			var a = rlib.RentableUser{
				RID:     v.RID,
				BID:     x.Ra.BID,
				TCID:    v1.TCID,
				DtStart: x.Ra.PossessionStart,
				DtStop:  x.Ra.PossessionStop,
			}
			if _, err := rlib.InsertRentableUser(ctx, &a); err != nil {
				return err
			}
		}
	}
	//----------------------------------------------------------------
	// Add the payers
	//----------------------------------------------------------------
	for _, v := range x.Raf.People {
		var a = rlib.RentalAgreementPayor{
			RAID:    x.Ra.RAID,
			BID:     x.Ra.BID,
			TCID:    v.TCID,
			DtStart: x.Ra.PossessionStart,
			DtStop:  x.Ra.PossessionStop,
			FLAGS:   0,
		}
		if _, err := rlib.InsertRentalAgreementPayor(ctx, &a); err != nil {
			return err
		}
	}

	return nil
}

// F2RAUpdatePets updates all pets. If the pet already exists then
// it just updates the pet. If the pet is
//
// INPUTS
//     ctx    - db context for transactions
//     x - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func F2RAUpdatePets(ctx context.Context, x *rlib.F2RAWriteHandlerContext) (err error) {
	rlib.Console("Entered F2RAUpdatePets\n")

	for i := 0; i < len(x.Raf.Pets); i++ {
		// get contact person
		var petTCID int64
		petTCID, err = GetTCIDForTMPTCID(x, x.Raf.Pets[i].TMPTCID)
		if err != nil {
			return err
		}

		// rlib.Console("petTCID = %d\n", petTCID)

		// PET ENTRY
		var pet rlib.Pet
		var bind rlib.TBind

		// updated tbind will be from new raid start time and for all future time.
		bind.SourceElemType = rlib.ELEMPERSON
		bind.SourceElemID = petTCID
		bind.AssocElemType = rlib.ELEMPET
		bind.AssocElemID = x.Raf.Pets[i].PETID
		bind.DtStart = x.Ra.PossessionStart
		bind.DtStop = rlib.ENDOFTIME
		bind.BID = x.Ra.BID

		// If PET exists then update it
		if x.Raf.Pets[i].PETID > 0 {
			// rlib.Console("Found existing pet %d\n", x.Raf.Pets[i].PETID)
			pet, err = rlib.GetPet(ctx, x.Raf.Pets[i].PETID)
			if err != nil {
				return err
			}
			// rlib.Console("A\n")
			//-----------------------------------------------------------------
			// update it if necessary
			//-----------------------------------------------------------------
			chgs := 0
			if pet.RAID != x.Ra.RAID {
				pet.RAID = x.Ra.RAID
				chgs++
			}
			if pet.Type != x.Raf.Pets[i].Type {
				pet.Type = x.Raf.Pets[i].Type
				chgs++
			}
			if pet.Breed != x.Raf.Pets[i].Breed {
				pet.Breed = x.Raf.Pets[i].Breed
				chgs++
			}
			if pet.Color != x.Raf.Pets[i].Color {
				pet.Color = x.Raf.Pets[i].Color
				chgs++
			}
			if pet.Weight != x.Raf.Pets[i].Weight {
				pet.Weight = x.Raf.Pets[i].Weight
				chgs++
			}
			if pet.Name != x.Raf.Pets[i].Name {
				pet.Name = x.Raf.Pets[i].Name
				chgs++
			}

			if chgs > 0 {
				if err = rlib.UpdatePet(ctx, &pet); err != nil {
					return err
				}
			}

			//----------------------------------------------------------------
			// 1. If the TCID did not change from the one in the TBind that
			//    overlaps the amend RAID start time, then no action is taken
			//    to the TBinds for this pet.
			//
			// 2. For the TBind that overlaps the amended RAID start time,
			//    update its DtStop to the start of the amended RAID.
			//
			// 3. Remove any other TBinds for this pet in the future.
			//
			// 4. Create a new TBind for this pet where DtStart = start of
			//    the amended RAID and DtStop is set to EOT (end of time).
			//----------------------------------------------------------------

			//----------------------------------------------------------
			// Get the TBinds for the pets impacted by this amendment.
			// Source type = PERSON,  Assoc type = PET
			//----------------------------------------------------------
			var tblist []rlib.TBind
			tblist, err = rlib.GetTBindAssocsByRange(ctx, pet.BID, rlib.ELEMPET, pet.PETID, rlib.ELEMPERSON, &bind.DtStart, &bind.DtStop)
			if err != nil {
				return err
			}
			// rlib.Console("B  tblist size = %d, pet.BID = %d, pet.PETID =%d\n", len(tblist), pet.BID, pet.PETID)
			//----------------------------------------------------
			// if only 1 and person hasn't changed, we're done
			//----------------------------------------------------
			if len(tblist) == 1 && tblist[0].SourceElemType == rlib.ELEMPERSON && tblist[0].SourceElemID == petTCID {
				// rlib.Console("C\n")
				return nil
			}
			// rlib.Console("D\n")
			//-------------------------------------------------------------
			// Spin through the records, update the overlapping record,
			// and remove the rest.
			//-------------------------------------------------------------
			for _, tb := range tblist {
				// rlib.Console("E tb.TBID = %d\n", tb.TBID)
				if rlib.DateInRange(&bind.DtStart, &tb.DtStart, &tb.DtStop) { // overlaps amended RAID ??
					tb.DtStop = bind.DtStart // YES: update its stop date
					// rlib.Console("F update TBind\n")
					if err = rlib.UpdateTBind(ctx, &tb); err != nil {
						return err
					}
				} else {
					// rlib.Console("F delete TBind\n")
					if err = rlib.DeleteTBind(ctx, tb.TBID); err != nil { // NO: delete it
						return err
					}
				}
			}
			// rlib.Console("G done with for loop\n")
			//-------------------------------------------------------------
			// Now add the TBind...
			//-------------------------------------------------------------
			_, err = rlib.InsertTBind(ctx, &bind)
			if err != nil {
				return err
			}
		} else {
			// rlib.Console("NEW PET\n")
			rlib.MigrateStructVals(&x.Raf.Pets[i], &pet)
			pet.TCID = petTCID
			pet.BID = x.Ra.BID
			// rlib.Console("NEW PET BID  = %d\n", pet.BID)

			// TODO: remove these soon, they are deprecated
			pet.DtStart = bind.DtStart
			pet.DtStop = bind.DtStop
			pet.RAID = x.Ra.RAID

			x.Raf.Pets[i].PETID, err = rlib.InsertPet(ctx, &pet)
			if err != nil {
				return err
			}
			bind.AssocElemID = pet.PETID
			_, err = rlib.InsertTBind(ctx, &bind)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// F2RAUpdateVehicles updates all vehicles from a flow. If the Vehicle already
// exists then it just updates the vehicle. If the Vehicle is new it creats
// it in the permanent tables.
//
// INPUTS
//     ctx - db context for transactions
//     x   - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func F2RAUpdateVehicles(ctx context.Context, x *rlib.F2RAWriteHandlerContext) (err error) {
	rlib.Console("Entered F2RAUpdateVehicles\n")

	for i := 0; i < len(x.Raf.Vehicles); i++ {
		// get contact person
		var VehicleTCID int64
		VehicleTCID, err = GetTCIDForTMPTCID(x, x.Raf.Vehicles[i].TMPTCID)
		if err != nil {
			return err
		}

		// VEHICLE ENTRY
		var v rlib.Vehicle
		var bind rlib.TBind

		// updated tbind will be from new raid start time and for all future time.
		bind.SourceElemType = rlib.ELEMPERSON
		bind.SourceElemID = VehicleTCID
		bind.AssocElemType = rlib.ELEMVEHICLE
		bind.AssocElemID = x.Raf.Vehicles[i].VID
		bind.DtStart = x.Ra.PossessionStart
		bind.DtStop = rlib.ENDOFTIME
		bind.BID = x.Ra.BID

		// If VEHICLE exists then update it
		if x.Raf.Vehicles[i].VID > 0 {
			err = rlib.GetVehicle(ctx, x.Raf.Vehicles[i].VID, &v)
			if err != nil {
				return err
			}
			//-----------------------------------------------------------------
			// update it if necessary
			//-----------------------------------------------------------------
			chgs := 0
			if v.VehicleType != x.Raf.Vehicles[i].VehicleType {
				v.VehicleType = x.Raf.Vehicles[i].VehicleType
				chgs++
			}
			if v.VehicleMake != x.Raf.Vehicles[i].VehicleMake {
				v.VehicleMake = x.Raf.Vehicles[i].VehicleMake
				chgs++
			}
			if v.VehicleModel != x.Raf.Vehicles[i].VehicleModel {
				v.VehicleModel = x.Raf.Vehicles[i].VehicleModel
				chgs++
			}
			if v.VehicleColor != x.Raf.Vehicles[i].VehicleColor {
				v.VehicleColor = x.Raf.Vehicles[i].VehicleColor
				chgs++
			}
			if v.VehicleYear != x.Raf.Vehicles[i].VehicleYear {
				v.VehicleYear = x.Raf.Vehicles[i].VehicleYear
				chgs++
			}
			if v.VIN != x.Raf.Vehicles[i].VIN {
				v.VIN = x.Raf.Vehicles[i].VIN
				chgs++
			}
			if v.LicensePlateState != x.Raf.Vehicles[i].LicensePlateState {
				v.LicensePlateState = x.Raf.Vehicles[i].LicensePlateState
				chgs++
			}
			if v.LicensePlateNumber != x.Raf.Vehicles[i].LicensePlateNumber {
				v.LicensePlateNumber = x.Raf.Vehicles[i].LicensePlateNumber
				chgs++
			}
			if v.ParkingPermitNumber != x.Raf.Vehicles[i].ParkingPermitNumber {
				v.ParkingPermitNumber = x.Raf.Vehicles[i].ParkingPermitNumber
				chgs++
			}

			if chgs > 0 {
				if err = rlib.UpdateVehicle(ctx, &v); err != nil {
					return err
				}
			}
			//----------------------------------------------------------------
			// 1. If the TCID did not change from the one in the TBind that
			//    overlaps the amend RAID start time, then no action is taken
			//    to the TBinds for this Vehicle.
			//
			// 2. For the TBind that overlaps the amended RAID start time,
			//    update its DtStop to the start of the amended RAID.
			//
			// 3. Remove any other TBinds for this Vehicle in the future.
			//
			// 4. Create a new TBind for this Vehicle where DtStart = start of
			//    the amended RAID and DtStop is set to EOT (end of time).
			//----------------------------------------------------------------

			//----------------------------------------------------------
			// Get the TBinds for the Vehicles impacted by this amendment.
			// Source type = PERSON,  Assoc type = VEHICLE
			//----------------------------------------------------------
			var tblist []rlib.TBind
			tblist, err = rlib.GetTBindAssocsByRange(ctx, v.BID, rlib.ELEMVEHICLE, v.VID, rlib.ELEMPERSON, &bind.DtStart, &bind.DtStop)
			if err != nil {
				return err
			}
			//----------------------------------------------------
			// if only 1 and person hasn't changed, we're done
			//----------------------------------------------------
			if len(tblist) == 1 && tblist[0].SourceElemType == rlib.ELEMPERSON && tblist[0].SourceElemID == VehicleTCID {
				return nil
			}
			//-------------------------------------------------------------
			// Spin through the records, update the overlapping record,
			// and remove the rest.
			//-------------------------------------------------------------
			for _, tb := range tblist {
				if rlib.DateInRange(&bind.DtStart, &tb.DtStart, &tb.DtStop) { // overlaps amended RAID ??
					tb.DtStop = bind.DtStart // YES: update its stop date
					if err = rlib.UpdateTBind(ctx, &tb); err != nil {
						return err
					}
				} else {
					if err = rlib.DeleteTBind(ctx, tb.TBID); err != nil { // NO: delete it
						return err
					}
				}
			}
			//-------------------------------------------------------------
			// Now add the TBind...
			//-------------------------------------------------------------
			_, err = rlib.InsertTBind(ctx, &bind)
			if err != nil {
				return err
			}
		} else {
			rlib.MigrateStructVals(&x.Raf.Vehicles[i], &v)
			v.BID = x.Ra.BID

			// TODO: remove these soon, they are deprecated
			v.TCID = VehicleTCID
			v.DtStart = bind.DtStart
			v.DtStop = bind.DtStop

			x.Raf.Vehicles[i].VID, err = rlib.InsertVehicle(ctx, &v)
			if err != nil {
				return err
			}
			bind.AssocElemID = v.VID
			_, err = rlib.InsertTBind(ctx, &bind)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetTCIDForTMPTCID finds the TCID associated with the supplied tmptcid.
//
// INPUTS
//     ctx     - db context for transactions
//     x       - all the contextual info we need for performing this operation
//     TMPTCID - TMPTCID for person we want the associated RID
//
// RETURNS
//     TCID of associated Transactant, or -1 if not found
//-----------------------------------------------------------------------------
func GetTCIDForTMPTCID(x *rlib.F2RAWriteHandlerContext, TMPTCID int64) (TCID int64, err error) {

	for i := 0; i < len(x.Raf.People); i++ {
		if x.Raf.People[i].TMPTCID == TMPTCID {
			TCID = x.Raf.People[i].TCID
			break
		}
	}

	if 0 == TCID {
		return TCID, fmt.Errorf("No TCID found for TMPTCID = %d", TMPTCID)
	}
	return TCID, nil
}

// F2RAUpdatePeople adds or updates all people information.
//
// INPUTS
//     ctx    - db context for transactions
//     x - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func F2RAUpdatePeople(ctx context.Context, x *rlib.F2RAWriteHandlerContext) error {
	var err error
	// rlib.Console("Entered F2RAUpdatePeople\n")

	//-------------------------------------------------------------------
	// Spin through all the people and update or create as needed
	//-------------------------------------------------------------------
	for i := 0; i < len(x.Raf.People); i++ {
		var xp rlib.XPerson
		tcid := x.Raf.People[i].TCID
		// rlib.Console("Found persond: TMPTCID = %d, TCID = %d\n", x.Raf.People[i].TMPTCID, tcid)
		if tcid > 0 {
			//---------------------------
			// Update existing...
			//---------------------------
			if err = rlib.GetXPerson(ctx, tcid, &xp); err != nil {
				return err
			}
			rlib.MigrateStructVals(&x.Raf.People[i], &xp.Trn)
			rlib.MigrateStructVals(&x.Raf.People[i], &xp.Usr)
			rlib.MigrateStructVals(&x.Raf.People[i], &xp.Psp)
			rlib.MigrateStructVals(&x.Raf.People[i], &xp.Pay)
			if err = rlib.UpdateTransactant(ctx, &xp.Trn); nil != err {
				return err
			}
			if err = rlib.UpdateUser(ctx, &xp.Usr); nil != err {
				return err
			}
			if err = rlib.UpdatePayor(ctx, &xp.Pay); nil != err {
				return err
			}
			if err = rlib.UpdateProspect(ctx, &xp.Psp); nil != err {
				return err
			}
		} else {
			//---------------------------
			// Create new Transactant...
			//---------------------------
			rlib.MigrateStructVals(&x.Raf.People[i], &xp.Trn)
			rlib.MigrateStructVals(&x.Raf.People[i], &xp.Usr)
			rlib.MigrateStructVals(&x.Raf.People[i], &xp.Psp)
			rlib.MigrateStructVals(&x.Raf.People[i], &xp.Pay)
			xp.Trn.BID = x.Raf.Meta.BID
			tcid, err := rlib.InsertTransactant(ctx, &xp.Trn)
			if nil != err {
				return err
			}
			if tcid == 0 {
				return fmt.Errorf("Insert returned a 0 id")
			}
			x.Raf.People[i].TCID = tcid
			xp.Trn.TCID = tcid
			xp.Usr.TCID = tcid
			xp.Usr.BID = xp.Trn.BID
			xp.Pay.TCID = tcid
			xp.Pay.BID = xp.Trn.BID
			xp.Psp.TCID = tcid
			xp.Psp.BID = xp.Trn.BID
			_, err = rlib.InsertUser(ctx, &xp.Usr)
			if nil != err {
				return err
			}
			_, err = rlib.InsertPayor(ctx, &xp.Pay)
			if nil != err {
				return err
			}
			_, err = rlib.InsertProspect(ctx, &xp.Psp)
			if nil != err {
				return err
			}
			errlist := bizlogic.FinalizeTransactant(ctx, &xp.Trn)
			if len(errlist) > 0 {
				return bizlogic.BizErrorListToError(errlist)
			}
		}
	}
	return nil
}
