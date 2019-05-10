package ws

import (
	"context"
	"encoding/json"
	"rentroll/rlib"
	"time"
)

// GetRA2FlowCore does all the heavy lifting to create a Flow from a
// RentalAgreement
//
// INPUTS:
//     ctx       database context for transactions
//     ra        the rental agreement to move into a flow
//     uid       uid of the person creating this flow.  Typically it
//               will be the uid in the session.
//     EditFlag  true -> get a version to edit,  false -> view existing
//
// RETURNS:
//     the new flowID
//     any error encountered
//     service data
//-------------------------------------------------------------------------
func GetRA2FlowCore(ctx context.Context, ra *rlib.RentalAgreement, d *ServiceData, EditFlag bool) (FlowID int64, err error) {

	// convert permanent ra to flow data and get it
	var raf rlib.RAFlowJSONData
	if raf, err = rlib.ConvertRA2Flow(ctx, ra, EditFlag); err != nil {
		return
	}
	//---------------------------------------------------------------------
	//  Don't change any dates if the entire agreement ended in the past
	//  or starts in the future.
	//---------------------------------------------------------------------
	now := time.Now()
	//    FUTURE: "now" is before the agreement starts      PAST:  now is after the agreement stopped
	if !(now.Before(time.Time(raf.Dates.AgreementStart)) || now.After(time.Time(raf.Dates.AgreementStop))) {
		// CHANGE THE START DATES TO TODAY
		raf.Dates.AgreementStart = rlib.JSONDate(rlib.GetTodayUTCRoundingDate())
		raf.Dates.RentStart = rlib.JSONDate(rlib.GetTodayUTCRoundingDate())
		raf.Dates.PossessionStart = rlib.JSONDate(rlib.GetTodayUTCRoundingDate())

		// ----- RENT DATES CHANGED CHECK ----- //
		newRStart := (time.Time)(raf.Dates.RentStart)
		newRStop := (time.Time)(raf.Dates.RentStop)
		if !ra.RentStart.Equal(newRStart) { // SINCE WE CHANGED ONLY START DATE
			err = rlib.RentDateChangeRAFlowUpdates(ctx, raf.Meta.BID, newRStart, newRStop, &raf)
			if err != nil {
				return
			}
		}
	}

	//--------------------------------------------------
	// Change the state to application being completed
	// as we're creating new flow
	//--------------------------------------------------
	action := int64(rlib.RAActionApplicationBeingCompleted)
	state := raf.Meta.RAFLAGS & uint64(0xF)

	//--------------------------------------------------
	// reset meta info
	//--------------------------------------------------
	ActionResetMetaData(action, state, &raf.Meta)

	//--------------------------------------------------
	// set data in meta based on Action
	//--------------------------------------------------
	if err = SetActionMetaData(ctx, d, action, &raf.Meta); err != nil {
		return
	}

	//-------------------------------------------------------------------------
	// Save the flow to the db
	//-------------------------------------------------------------------------
	var raflowJSONData []byte
	if raflowJSONData, err = json.Marshal(&raf); err != nil {
		return
	}

	//-------------------------------------------------------------------------
	// Fill out the datastructure and save it to the db as a flow...
	//-------------------------------------------------------------------------

	a := rlib.Flow{
		BID:       ra.BID,
		FlowID:    0, // it's new flowID,
		UserRefNo: rlib.GenerateUserRefNo(),
		FlowType:  rlib.RAFlow,
		ID:        ra.RAID,
		Data:      raflowJSONData,
		CreateBy:  d.sess.UID,
		LastModBy: d.sess.UID,
	}

	// insert new flow
	if FlowID, err = rlib.InsertFlow(ctx, &a); err != nil {
		return
	}
	return
}

// addRAPtoFlow adds a new person to raf.People.  The renter/occupant flags
// are only set if the corresponding input bool value is set.
//
// INPUTS
//     tcid  = the tcid of the transactant to load
//      rid  - the rentable that they are tied to
//      raf  - pointer to the flow struct to update
//      chk  - check to see if the tcid exists in raf.People before adding.
//             This is not always necessary, but only the caller knows.
// isRenter  - true if we need to set the RAPerson isRenter bool to true.
//             It should be true for Payors.
// isOccupant- true if we need to set the RAPerson isOccupant bool to true.
//             It should be true for Users.
//
// RETURNS
//     any error encountered
//     raf is updated
//-----------------------------------------------------------------------------
func addRAPtoFlow(ctx context.Context, tcid, rid int64, raf *rlib.RAFlowJSONData, chk, isRenter, isOccupant bool) error {
	// Is this user already present?
	if chk {
		for l := 0; l < len(raf.People); l++ {
			if raf.People[l].TCID == tcid {
				if isRenter {
					raf.People[l].IsRenter = true
				}
				if isOccupant {
					raf.People[l].IsOccupant = true
				}
				return nil
			}
		}
	}

	rap, err := createRAFlowPerson(ctx, tcid, raf, isOccupant) // adds person AND associated pets and vehicles
	if err != nil {
		return err
	}

	if isRenter {
		rap.IsRenter = true
	}

	if isOccupant {
		rap.IsOccupant = true

		// only tie occupants to rentable
		var t rlib.RATiePeopleData
		t.TMPTCID = rap.TMPTCID
		if rid > 0 {
			t.PRID = rid
		}
		raf.Tie.People = append(raf.Tie.People, t)
	}

	// finally append in people list
	raf.People = append(raf.People, rap)
	return nil
}

// createRAFlowPerson returns a new rlib.RAPeopleFlowData based on the supplied
// tcid. It does not set the Renter or Occupant flags
//
// INPUTS
//          ctx  = db transaction context
//         tcid  = the tcid of the transactant to load
//          raf  = pointer to rlib.RAFlowJSONData
// addDependents = adds dependents (currently pets and vehicles) to the flow
//                 data in addition to the transactant data. The recommended
//                 usage of this flag is to set it to true when the person
//                 being added is a user.
//
// RETURNS
//     rlib.RAPeopleFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func createRAFlowPerson(ctx context.Context, tcid int64, raf *rlib.RAFlowJSONData, addDependents bool) (rlib.RAPeopleFlowData, error) {
	var p rlib.Transactant
	var pu rlib.User
	var pp rlib.Payor
	var pr rlib.Prospect
	var rap rlib.RAPeopleFlowData
	var err error

	raf.Meta.LastTMPTCID++
	rap.TMPTCID = raf.Meta.LastTMPTCID // set this now so it is available when creating pets and vehicles
	if err = rlib.GetTransactant(ctx, tcid, &p); err != nil {
		return rap, err
	}
	if err = rlib.GetUser(ctx, tcid, &pu); err != nil {
		return rap, err
	}
	if err = rlib.GetPayor(ctx, tcid, &pp); err != nil {
		return rap, err
	}
	if err = rlib.GetProspect(ctx, tcid, &pr); err != nil {
		return rap, err
	}
	rlib.MigrateStructVals(&p, &rap)
	rlib.MigrateStructVals(&pp, &rap)
	rlib.MigrateStructVals(&pu, &rap)
	rlib.MigrateStructVals(&pr, &rap)

	if addDependents {
		if err = addFlowPersonVehicles(ctx, tcid, rap.TMPTCID, raf); err != nil {
			return rap, err
		}
		if err = addFlowPersonPets(ctx, tcid, rap.TMPTCID, raf); err != nil {
			return rap, err
		}
	}
	return rap, nil
}

// addFlowPersonPets adds pets belonging to tcid to the supplied
// rlib.RAFlowJSONData struct
//
// INPUTS
//      ctx  = db transaction context
//     tcid  = the tcid of the transactant to load
//
// RETURNS
//     rlib.RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func addFlowPersonPets(ctx context.Context, tcid, tmptcid int64, raf *rlib.RAFlowJSONData) error {
	petList, err := rlib.GetPetsByTransactant(ctx, tcid)
	if err != nil {
		return err
	}
	for i := 0; i < len(petList); i++ {
		raf.Meta.LastTMPPETID++
		var p = rlib.RAPetsFlowData{
			TMPTCID:  tmptcid,
			TMPPETID: raf.Meta.LastTMPPETID,
		}
		rlib.MigrateStructVals(&petList[i], &p)
		raf.Pets = append(raf.Pets, p)
	}
	return nil
}

// addFlowPersonVehicles adds vehicles belonging to tcid to the supplied
// rlib.RAFlowJSONData struct
//
// INPUTS
//      ctx  = db transaction context
//     tcid  = the tcid of the transactant to load
//
// RETURNS
//     rlib.RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func addFlowPersonVehicles(ctx context.Context, tcid, tmptcid int64, raf *rlib.RAFlowJSONData) error {
	vehicleList, err := rlib.GetVehiclesByTransactant(ctx, tcid)
	if err != nil {
		return err
	}
	for i := 0; i < len(vehicleList); i++ {
		raf.Meta.LastTMPVID++
		var v = rlib.RAVehiclesFlowData{
			TMPTCID: tmptcid,
			TMPVID:  raf.Meta.LastTMPVID,
		}
		rlib.MigrateStructVals(&vehicleList[i], &v)
		raf.Vehicles = append(raf.Vehicles, v)
	}
	return nil
}
