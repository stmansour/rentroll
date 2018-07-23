package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"time"
)

// WriteHandlerContext contains context information for RA Write Handlers
type WriteHandlerContext struct {
	ra     rlib.RentalAgreement
	raOrig rlib.RentalAgreement
	raf    rlib.RAFlowJSONData
	xbiz   rlib.XBusiness
}

// RAWriteHandler a handler function for part of the work of migrating
// RAFlow data into the permanent tables for a complete RentalAgreement
type RAWriteHandler struct {
	Name    string
	Handler func(context.Context, *WriteHandlerContext) error
}

// UpdateHandlers is the collection of routines to call to write a flow
// for an existing Rental Agreement back to the database as a RentalAgreement.
var ehandlers = []RAWriteHandler{
	{"RentalAgreement", F2RAUpdateExistingRA},
	{"Transactants", F2RAUpdatePeople},
	{"Pets", F2RAUpdatePets},
	{"Vehicles", F2RAUpdateVehicles},
	{"Rentables", FlowSaveRentables},
	{"Fees", Fees2RA},
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
	var x WriteHandlerContext
	var nraid int64

	//-------------------------------------------
	// Read the flow data into a data structure
	//-------------------------------------------
	flow, err := rlib.GetFlow(ctx, flowid)
	if err != nil {
		rlib.Console("\n\nERROR IN GetFlow: %s\n\n\n", err.Error())
		return nraid, err
	}
	err = json.Unmarshal(flow.Data, &x.raf)
	if err != nil {
		rlib.Console("\n\nERROR IN Unmarshal: %s\n\n\n", err.Error())
		return nraid, err
	}

	//----------------------------------------------------------------------------
	// If this is an update of an existing RAID, check to see if any changes
	// were made. Otherwise treat it as a new RAID
	//----------------------------------------------------------------------------
	if x.raf.Meta.RAID > 0 {
		//------------------------------
		// TODO: check for any changes
		//------------------------------
		changes := true
		if changes {
			nraid, err = FlowSaveRA(ctx, &x)
			return nraid, err
		}

		//----------------------------------------------------
		// if there were no changes, just delete the flow...
		//----------------------------------------------------
		//return nraid, DeleteFlow(ctx, flowid)
	}

	return nraid, nil
}

// FlowSaveRA saves a new Rental Agreement from the supplied flow. This
//     function assumes that a check has already been made to verify that
//     the RentalAgreement is either new or, if it is replacing an existing
//     rental agreement, that the data has actually been changed.
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
func FlowSaveRA(ctx context.Context, x *WriteHandlerContext) (int64, error) {
	rlib.Console("Entered FlowSaveRA\n")
	var err error
	var nraid int64

	if err = rlib.InitBizInternals(x.raf.Dates.BID, &x.xbiz); err != nil {
		return nraid, err
	}

	if x.raf.Meta.RAID > 0 {
		//------------------------------------------------------------
		// Get the rental agreement that will be superceded by the
		// one we're creating here. Update its stop dates accordingly
		//------------------------------------------------------------
		x.raOrig, err = rlib.GetRentalAgreement(ctx, x.raf.Meta.RAID)
		if err != nil {
			return nraid, err
		}
		chgs := 0
		AStart := time.Time(x.raf.Dates.AgreementStart)
		RStart := time.Time(x.raf.Dates.RentStart)
		PStart := time.Time(x.raf.Dates.PossessionStart)
		if x.raOrig.AgreementStop.After(AStart) {
			x.raOrig.AgreementStop = AStart
			chgs++
		}
		if x.raOrig.RentStop.After(RStart) {
			x.raOrig.RentStop = RStart
			chgs++
		}
		if x.raOrig.PossessionStop.After(PStart) {
			x.raOrig.PossessionStop = PStart
			chgs++
		}
		//------------------------------------------------------------------
		// If there are changes, then we stop the old Rental Agreement and
		// create a new one linked to x.raOrig
		//------------------------------------------------------------------
		if chgs > 0 {
			x.raOrig.FLAGS &= ^uint64(0x7) // clear the status
			x.raOrig.FLAGS |= 5            // set the state to Terminated
			x.raOrig.LeaseTerminationReason =
				rlib.RRdb.BizTypes[x.raOrig.BID].Msgs.S[rlib.MSGRAUPDATED].SLSID // "Rental Agreement was updated"
			sess, ok := rlib.SessionFromContext(ctx)
			if !ok {
				return nraid, rlib.ErrSessionRequired
			}
			x.raOrig.TerminatorUID = sess.UID
			x.raOrig.TerminationDate = time.Now()

			err = rlib.UpdateRentalAgreement(ctx, &x.raOrig)
			if err != nil {
				return nraid, err
			}
		}

		//------------------------------------------------------------
		// Now start the new RAID.  Link it to x.raOrig
		//------------------------------------------------------------
		x.ra.AgreementStart = time.Time(x.raf.Dates.AgreementStart)
		x.ra.AgreementStop = time.Time(x.raf.Dates.AgreementStop)
		x.ra.RentStart = time.Time(x.raf.Dates.RentStart)
		x.ra.RentStop = time.Time(x.raf.Dates.RentStop)
		x.ra.PossessionStart = time.Time(x.raf.Dates.PossessionStart)
		x.ra.PossessionStop = time.Time(x.raf.Dates.PossessionStop)
		x.ra.PRAID = x.raOrig.RAID
		x.ra.ORIGIN = x.raOrig.ORIGIN
		x.ra.BID = x.raOrig.BID
		x.ra.CSAgent = x.raOrig.CSAgent
		if x.raOrig.ORIGIN == 0 {
			x.ra.ORIGIN = x.raOrig.RAID
		}
		x.ra.RATID = x.raOrig.RATID
		x.ra.RentCycleEpoch = x.raOrig.RentCycleEpoch

		nraid, err = rlib.InsertRentalAgreement(ctx, &x.ra)
		if err != nil {
			return nraid, err
		}

		//---------------------------------------------------------------
		// Now spin through the series of handlers that move the data
		// into the permanent tables...
		//---------------------------------------------------------------
		rlib.Console("len(ehandlers) = %d\n", len(ehandlers))
		for i := 0; i < len(ehandlers); i++ {
			rlib.Console("FlowSaveRA: running handler %s\n", ehandlers[i].Name)
			if err = ehandlers[i].Handler(ctx, x); err != nil {
				return nraid, err
			}
		}
	} else {
		//-------------------------------------
		// This is a new Rental Agreement...
		//-------------------------------------
	}

	return nraid, nil
}

// FlowSaveRentables adds/updates rentables from the flow data.  This means
// that we update or add the RentalAgreementRentables list.  Update means
// that we set the stop date for the existing RentalAgreementRentables RAID.
// Then we add the Rentables in x.raf.Rentables[] into a
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
func FlowSaveRentables(ctx context.Context, x *WriteHandlerContext) error {
	rlib.Console("Entered FlowSaveRentables\n")
	//----------------------------------------------------------------
	// Update the stop date on any existing RentalAgreementRentables
	//----------------------------------------------------------------
	if x.raf.Meta.RAID > 0 {
		rarl, err := rlib.GetAllRentalAgreementRentables(ctx, x.raf.Meta.RAID)
		if err != nil {
			return err
		}
		for _, v := range rarl {
			v.RARDtStop = time.Time(x.raf.Dates.AgreementStart)
			if err = rlib.UpdateRentalAgreementRentable(ctx, &v); err != nil {
				return err
			}
			//----------------------------------------------------------------
			// Fix up the users
			//----------------------------------------------------------------
			rul, err := rlib.GetRentableUsersInRange(ctx, v.RID, &x.raOrig.PossessionStart, &x.ra.PossessionStop)
			if err != nil {
				return err
			}
			for _, ru := range rul {
				ru.DtStop = x.raOrig.PossessionStop
				if err = rlib.UpdateRentableUser(ctx, &ru); err != nil {
					return err
				}
			}
		}
		//----------------------------------------------------------------
		// Fix up the payors
		//----------------------------------------------------------------
		t, err := rlib.GetRentalAgreementPayorsByRAID(ctx, x.raf.Meta.RAID)
		if err != nil {
			return err
		}
		for _, v := range t {
			if v.DtStop.After(x.ra.RentStart) {
				v.DtStop = x.ra.RentStart
				if err = rlib.UpdateRentalAgreementPayor(ctx, &v); err != nil {
					return err
				}
			}
		}
	}

	//----------------------------------------------------------------
	// Add a RentalAgreementRentable entry for each Rentable
	//----------------------------------------------------------------
	for _, v := range x.raf.Rentables {
		var rar = rlib.RentalAgreementRentable{
			RAID:         x.ra.RAID,
			BID:          x.ra.BID,
			RID:          v.RID,
			CLID:         0, // TODO:
			ContractRent: 0, // TODO:
			RARDtStart:   time.Time(x.raf.Dates.PossessionStart),
			RARDtStop:    time.Time(x.raf.Dates.PossessionStop),
		}
		_, err := rlib.InsertRentalAgreementRentable(ctx, &rar)
		if err != nil {
			return err
		}

		//----------------------------------------------------------------
		// Add the users
		//----------------------------------------------------------------
		for _, v1 := range x.raf.People {
			var a = rlib.RentableUser{
				RID:     v.RID,
				BID:     x.ra.BID,
				TCID:    v1.TCID,
				DtStart: x.ra.PossessionStart,
				DtStop:  x.ra.PossessionStop,
			}
			if _, err := rlib.InsertRentableUser(ctx, &a); err != nil {
				return err
			}
		}
	}
	//----------------------------------------------------------------
	// Add the payers
	//----------------------------------------------------------------
	for _, v := range x.raf.People {
		var a = rlib.RentalAgreementPayor{
			RAID:    x.ra.RAID,
			BID:     x.ra.BID,
			TCID:    v.TCID,
			DtStart: x.ra.PossessionStart,
			DtStop:  x.ra.PossessionStop,
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
func F2RAUpdatePets(ctx context.Context, x *WriteHandlerContext) error {
	rlib.Console("Entered F2RAUpdatePets\n")
	var err error
	for i := 0; i < len(x.raf.Pets); i++ {
		var pet rlib.RentalAgreementPet
		if x.raf.Pets[i].PETID > 0 {
			pet, err = rlib.GetRentalAgreementPet(ctx, x.raf.Pets[i].PETID)
			if err != nil {
				return err
			}
			rlib.MigrateStructVals(&x.raf.Pets[i], &pet)
			if err = rlib.UpdateRentalAgreementPet(ctx, &pet); err != nil {
				return err
			}
			continue // all done, move on to the next pet
		}
		rlib.MigrateStructVals(&x.raf.Pets[i], &pet)
		pet.RAID = x.raf.Meta.RAID
		_, err = rlib.InsertRentalAgreementPet(ctx, &pet)
		if err != nil {
			return err
		}
	}
	return nil
}

// F2RAUpdateVehicles updates all pets
//
// INPUTS
//     ctx    - db context for transactions
//     x - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func F2RAUpdateVehicles(ctx context.Context, x *WriteHandlerContext) error {
	rlib.Console("Entered F2RAUpdateVehicles\n")
	for i := 0; i < len(x.raf.Vehicles); i++ {
		tcid, err := findVehiclePointPerson(x, x.raf.Vehicles[i].TMPTCID, x.raf.Vehicles[i].TMPVID)
		if err != nil {
			return err
		}
		//-------------------------------
		// handle existing vehicles...
		//-------------------------------
		if x.raf.Vehicles[i].VID > 0 {
			vehicles, err := rlib.GetVehiclesByTransactant(ctx, tcid)
			if err != nil {
				return err
			}
			for j := 0; j < len(vehicles); j++ {
				rlib.MigrateStructVals(&x.raf.Vehicles[i], &vehicles[j])
				vehicles[j].TCID = tcid
				// rlib.Console("Just before UpdateVehicle: vehicles[j] = %#v\n", vehicles[j])
				if err = rlib.UpdateVehicle(ctx, &vehicles[j]); err != nil {
					return err
				}
			}
			continue // all done, move on to the next vehicle
		}
		//-------------------------------
		// handle new vehicles...
		//-------------------------------
		var vehicle rlib.Vehicle
		rlib.MigrateStructVals(&x.raf.Vehicles[i], &vehicle)
		vehicle.TCID = tcid
		_, err = rlib.InsertVehicle(ctx, &vehicle)
		if err != nil {
			return err
		}
	}
	return nil
}

// findVehiclePointPerson returns the TCID of the person associated with
// vehicle TMPVID
//--------------------------------------------------------------------------------
func findVehiclePointPerson(x *WriteHandlerContext, t, tmpvid int64) (int64, error) {
	tcid := int64(0)
	// find the point person
	for j := 0; j < len(x.raf.People); j++ {
		if t == x.raf.People[j].TMPTCID {
			tcid = x.raf.People[j].TCID
			break
		}
	}
	if 0 == tcid {
		return tcid, fmt.Errorf("No TCID found for Vehicle VID=%d", tmpvid)
	}
	return tcid, nil
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
func F2RAUpdatePeople(ctx context.Context, x *WriteHandlerContext) error {
	var err error
	rlib.Console("Entered F2RAUpdatePeople\n")

	//-------------------------------------------------------------------
	// Spin through all the people and update or create as needed
	//-------------------------------------------------------------------
	for i := 0; i < len(x.raf.People); i++ {
		var xp rlib.XPerson
		tcid := x.raf.People[i].TCID
		rlib.Console("Found persond: TMPTCID = %d, TCID = %d\n", x.raf.People[i].TMPTCID, tcid)
		if tcid > 0 {
			//---------------------------
			// Update existing...
			//---------------------------
			if err = rlib.GetXPerson(ctx, tcid, &xp); err != nil {
				return err
			}
			rlib.MigrateStructVals(&x.raf.People[i], &xp.Trn)
			rlib.MigrateStructVals(&x.raf.People[i], &xp.Usr)
			rlib.MigrateStructVals(&x.raf.People[i], &xp.Psp)
			rlib.MigrateStructVals(&x.raf.People[i], &xp.Pay)
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
			rlib.MigrateStructVals(&x.raf.People[i], &xp.Trn)
			rlib.MigrateStructVals(&x.raf.People[i], &xp.Usr)
			rlib.MigrateStructVals(&x.raf.People[i], &xp.Psp)
			rlib.MigrateStructVals(&x.raf.People[i], &xp.Pay)
			tcid, err := rlib.InsertTransactant(ctx, &xp.Trn)
			if nil != err {
				return err
			}
			if tcid == 0 {
				return fmt.Errorf("Insert returned a 0 id")
			}
			x.raf.People[i].TCID = tcid
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

// F2RAUpdateExistingRA creates a new rental agreement and links it to its
// parent
//
// INPUTS
//     ctx    - db context for transactions
//     x - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func F2RAUpdateExistingRA(ctx context.Context, x *WriteHandlerContext) error {
	// var err error
	rlib.Console("Entered F2RAUpdateExistingRA\n")
	return nil
}
