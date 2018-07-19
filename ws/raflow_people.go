package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// RAPersonDetailsRequest is struct for request for person details
type RAPersonDetailsRequest struct {
	TCID   int64
	FlowID int64
}

// RAFlowRemovePersonRequest is struct for request to remove person from json data
type RAFlowRemovePersonRequest struct {
	TMPTCID int64
	FlowID  int64
}

// SvcRAFlowPersonHandler handles operation on person of raflow json data
//           0    1     2   3
// uri /v1/raflow-person/BID/flowID
// The server command can be:
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcRAFlowPersonHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRAFlowPersonHandler"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "save":
		SaveRAFlowPersonDetails(w, r, d)
		break
	case "delete":
		DeleteRAFlowPerson(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// SaveRAFlowPersonDetails saves person details with list of pets and vehicles
// It modifies raflow json doc by writing fetched pets and vehicles data
// wsdoc {
//  @Title Save Person details with list of Pets & Vehicles
//  @URL /v1/raflow-persondetails/:BUI/
//  @Method  GET
//  @Synopsis Save Person Details for RAFlow
//  @Description Save details about person with pets and vehicles
//  @Input RAPersonDetailsRequest
//  @Response FlowResponse
// wsdoc }
func SaveRAFlowPersonDetails(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SaveRAFlowPersonDetails"
	var (
		raFlowData    rlib.RAFlowJSONData
		foo           RAPersonDetailsRequest
		modRAFlowMeta rlib.RAFlowMetaInfo // we might need to update meta info
		g             FlowResponse
		err           error
		tx            *sql.Tx
		ctx           context.Context
		prospectFlag  uint64
	)
	fmt.Printf("Entered %s\n", funcname)

	// ===============================================
	// defer function to handle transactaion rollback
	// ===============================================
	defer func() {
		if err != nil {
			// if tx is not nil then roll back
			if tx != nil {
				tx.Rollback()
			}
			SvcErrorReturn(w, err, funcname)
			return
		}
	}()

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("only POST method is allowed")
		return
	}

	// unmarshal data into request data struct
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		return
	}

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err = rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		return
	}

	// get flow and it must exist
	var flow rlib.Flow
	flow, err = rlib.GetFlow(ctx, foo.FlowID)
	if err != nil {
		return
	}

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		return
	}

	// get flow meta data in modRAFlowMeta, which is going to modified if required
	modRAFlowMeta = raFlowData.Meta

	// ----------------------------------------------
	// get person details with given TCID
	// ----------------------------------------------
	personTMPTCID := int64(0)

	// this is for accept Transactant, so find it by TCID
	tcidExistInJSONData := false
	for i := range raFlowData.People {
		if raFlowData.People[i].TCID == foo.TCID {
			tcidExistInJSONData = true
			personTMPTCID = raFlowData.People[i].TMPTCID
			break
		}
	}

	if !tcidExistInJSONData {
		newRAFlowPerson := rlib.RAPeopleFlowData{}
		var xp rlib.XPerson
		err = rlib.GetXPerson(ctx, foo.TCID, &xp)
		if err != nil {
			return
		}

		// migrate field values to Person details
		if xp.Pay.TCID > 0 {
			rlib.MigrateStructVals(&xp.Pay, &newRAFlowPerson)
		}
		if xp.Psp.TCID > 0 {
			rlib.MigrateStructVals(&xp.Psp, &newRAFlowPerson)
		}
		if xp.Usr.TCID > 0 {
			rlib.MigrateStructVals(&xp.Usr, &newRAFlowPerson)
		}
		if xp.Trn.TCID > 0 {
			rlib.MigrateStructVals(&xp.Trn, &newRAFlowPerson)
		}
		newRAFlowPerson.BID = d.BID

		// check for additional flags IsRenter, IsOccupant
		newRAFlowPerson.IsOccupant = true
		if len(raFlowData.People) == 0 { // this is first transactant
			newRAFlowPerson.IsRenter = true
		}

		// custom tmp tcid
		modRAFlowMeta.LastTMPTCID++
		newRAFlowPerson.TMPTCID = modRAFlowMeta.LastTMPTCID
		personTMPTCID = newRAFlowPerson.TMPTCID

		// Manage "Have you ever been"(Prospect) section FLAGS
		prospectFlag = xp.Psp.FLAGS
		newRAFlowPerson.Evicted = prospectFlag&0x1 != 0    // 1 << 0
		newRAFlowPerson.Convicted = prospectFlag&0x2 != 0  // 1 << 1
		newRAFlowPerson.Bankruptcy = prospectFlag&0x4 != 0 // 1 << 2

		// append in json list
		raFlowData.People = append(raFlowData.People, newRAFlowPerson)

		var modPeopleData []byte
		modPeopleData, err = json.Marshal(&raFlowData.People)
		if err != nil {
			return
		}

		// update flow with this modified people part
		err = rlib.UpdateFlowData(ctx, "people", modPeopleData, &flow)
		if err != nil {
			return
		}
	}

	// -------------------------------------------
	// find pets list associated with current TCID
	// -------------------------------------------

	// get the list of pets
	var petList []rlib.RentalAgreementPet
	petList, err = rlib.GetPetsByTransactant(ctx, foo.TCID)
	if err != nil {
		return
	}

	// find this RID in flow data rentable list
	shouldModifyPetsData := false
	for i := range petList {
		exist := false
		for k := range raFlowData.Pets {
			if petList[i].PETID == raFlowData.Pets[k].PETID {
				exist = true
				break
			}
		}

		// if does not exist then append in the raflow data
		if !exist {
			// create new pet info
			var newRAFlowPet rlib.RAPetsFlowData
			newRAFlowPet, err = rlib.NewRAFlowPet(ctx, d.BID,
				raFlowData.Dates.PossessionStart, raFlowData.Dates.PossessionStop, &modRAFlowMeta)
			if err != nil {
				return
			}

			// migrate data
			rlib.MigrateStructVals(&petList[i], &newRAFlowPet)

			// TMPTCID
			newRAFlowPet.TMPTCID = personTMPTCID

			// append in pets list
			raFlowData.Pets = append(raFlowData.Pets, newRAFlowPet)

			// should modify the content in raflow json?
			shouldModifyPetsData = true
		}
	}

	if shouldModifyPetsData {
		var modPetsData []byte
		modPetsData, err = json.Marshal(&raFlowData.Pets)
		if err != nil {
			return
		}

		// update flow with this modified pets part
		err = rlib.UpdateFlowData(ctx, "pets", modPetsData, &flow)
		if err != nil {
			return
		}
	}

	// -----------------------------------------------
	// find vehicles list associated with current TCID
	// -----------------------------------------------

	// get the list of pets
	var vehicleList []rlib.Vehicle
	vehicleList, err = rlib.GetVehiclesByTransactant(ctx, foo.TCID)
	if err != nil {
		return
	}

	// loop over list and append it in raflow data
	shouldModifyVehiclesData := false
	for i := range vehicleList {
		exist := false
		for k := range raFlowData.Vehicles {
			if vehicleList[i].VID == raFlowData.Vehicles[k].VID {
				exist = true
				break
			}
		}

		// if does not exist then append in the raflow data
		if !exist {
			// create new pet info
			var newRAFlowVehicle rlib.RAVehiclesFlowData
			newRAFlowVehicle, err = rlib.NewRAFlowVehicle(ctx, d.BID,
				raFlowData.Dates.PossessionStart, raFlowData.Dates.PossessionStop, &modRAFlowMeta)
			if err != nil {
				return
			}

			// migrate existing values
			rlib.MigrateStructVals(&vehicleList[i], &newRAFlowVehicle)

			// contact person
			newRAFlowVehicle.TMPTCID = personTMPTCID

			// append in vehicles list of json data
			raFlowData.Vehicles = append(raFlowData.Vehicles, newRAFlowVehicle)

			// should modify content for raflow json
			shouldModifyVehiclesData = true
		}
	}

	if shouldModifyVehiclesData {
		// get marshalled data
		var modVData []byte
		modVData, err = json.Marshal(&raFlowData.Vehicles)
		if err != nil {
			return
		}

		// update flow with this modified vehicles part
		err = rlib.UpdateFlowData(ctx, "vehicles", modVData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// update meta info if required
	// ----------------------------------------------
	if raFlowData.Meta.LastTMPASMID < modRAFlowMeta.LastTMPASMID {

		// Update HavePets Flag in meta information of flow
		modRAFlowMeta.HavePets = len(raFlowData.Pets) > 0
		modRAFlowMeta.HaveVehicles = len(raFlowData.Vehicles) > 0

		// get marshalled data
		var modMetaData []byte
		modMetaData, err = json.Marshal(&modRAFlowMeta)
		if err != nil {
			return
		}

		// update flow with this modified meta part
		err = rlib.UpdateFlowData(ctx, "meta", modMetaData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// return response
	// ----------------------------------------------

	// get the modified flow
	flow, err = rlib.GetFlow(ctx, flow.FlowID)
	if err != nil {
		return
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err = tx.Commit(); err != nil {
		return
	}

	// set the response
	g.Record = flow
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// DeleteRAFlowPerson remove person from raflow data as well as removes
// associated pets and vehicles data too
// wsdoc {
//  @Title Remvoe Person with list of associated Pets & Vehicles
//  @URL /v1/raflow-person/:BUI/:FlowID
//  @Method POST
//  @Synopsis Remove Person from RAFlow json data
//  @Description Remove details about person with associated pets and vehicles
//  @Input RAFlowRemovePersonRequest
//  @Response FlowResponse
// wsdoc }
func DeleteRAFlowPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "DeleteRAFlowPerson"
	var (
		raFlowData rlib.RAFlowJSONData
		foo        RAFlowRemovePersonRequest
		g          FlowResponse
		err        error
		tx         *sql.Tx
		ctx        context.Context
	)
	fmt.Printf("Entered %s\n", funcname)

	// ===============================================
	// defer function to handle transactaion rollback
	// ===============================================
	defer func() {
		if err != nil {
			if tx != nil {
				tx.Rollback()
			}
			SvcErrorReturn(w, err, funcname)
			return
		}
	}()

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("Only POST method is allowed")
		return
	}

	// unmarshal data into request data struct
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		return
	}

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err = rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		return
	}

	// get flow and it must exist
	var flow rlib.Flow
	flow, err = rlib.GetFlow(ctx, foo.FlowID)
	if err != nil {
		return
	}

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		return
	}

	// ----------------------------------------------
	// get person details with given TMPTCID
	// ----------------------------------------------
	personTMPTCID := int64(0)

	// this is for accept Transactant, so find it by TMPTCID
	tcidExistInJSONData := false
	for i := range raFlowData.People {
		if raFlowData.People[i].TMPTCID == foo.TMPTCID {
			tcidExistInJSONData = true
			personTMPTCID = raFlowData.People[i].TMPTCID

			// remove the element then
			raFlowData.People = append(raFlowData.People[:i], raFlowData.People[i+1:]...)

			break
		}
	}

	if tcidExistInJSONData {
		var modPeopleData []byte
		modPeopleData, err = json.Marshal(&raFlowData.People)
		if err != nil {
			return
		}

		// update flow with this modified people part
		err = rlib.UpdateFlowData(ctx, "people", modPeopleData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// remove associated pets
	// ----------------------------------------------
	shouldModifyPetsData := false
	for i := range raFlowData.Pets {
		if raFlowData.Pets[i].TMPTCID == personTMPTCID {
			shouldModifyPetsData = true
			// remove this pet from the list
			raFlowData.Pets = append(raFlowData.Pets[:i], raFlowData.Pets[i+1:]...)
		}
	}

	if shouldModifyPetsData {
		var modPetsData []byte
		modPetsData, err = json.Marshal(&raFlowData.Pets)
		if err != nil {
			return
		}

		// update flow with this modified pets part
		err = rlib.UpdateFlowData(ctx, "pets", modPetsData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// remove associated vehicles
	// ----------------------------------------------
	shouldModifyVehiclesData := false
	for i := range raFlowData.Vehicles {
		if raFlowData.Vehicles[i].TMPTCID == personTMPTCID {
			shouldModifyVehiclesData = true
			// remove this pet from the list
			raFlowData.Vehicles = append(raFlowData.Vehicles[:i], raFlowData.Vehicles[i+1:]...)
		}
	}

	if shouldModifyVehiclesData {
		// get marshalled data
		var modVData []byte
		modVData, err = json.Marshal(&raFlowData.Vehicles)
		if err != nil {
			return
		}

		// update flow with this modified vehicles part
		err = rlib.UpdateFlowData(ctx, "vehicles", modVData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// return response
	// ----------------------------------------------

	// get the modified flow
	flow, err = rlib.GetFlow(ctx, flow.FlowID)
	if err != nil {
		return
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err = tx.Commit(); err != nil {
		return
	}

	// set the response
	g.Record = flow
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
