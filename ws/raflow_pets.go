package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// SvcRAFlowPetsHandler handles operation on pets of raflow json data
//           0    1     2   3
// uri /v1/raflow-pets/BID/flowID
// The server command can be:
//      new
//-----------------------------------------------------------------------------------
func SvcRAFlowPetsHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRAFlowPetsHandler"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "new":
		CreateNewRAFlowPet(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// RAFlowNewPetRequest struct for new pet request
type RAFlowNewPetRequest struct {
	FlowID int64
}

// CreateNewRAFlowPet creates new entry in pets list of raflow json data
// wsdoc {
//  @Title create brand new pet entry in pet list
//  @URL /v1/raflow-pets/:BUI/:FlowID
//  @Method POST
//  @Synopsis create new entry of pet in RAFlow json data
//  @Description create new entry of pet in RAFlow json data with all available fees
//  @Input RAFlowNewPetRequest
//  @Response FlowResponse
// wsdoc }
//-----------------------------------------------------------------------------------
func CreateNewRAFlowPet(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "CreateNewRAFlowPet"
	var (
		foo           RAFlowNewPetRequest
		raFlowData    rlib.RAFlowJSONData
		err           error
		tx            *sql.Tx
		ctx           context.Context
		modRAFlowMeta rlib.RAFlowMetaInfo
	)
	fmt.Printf("Entered in %s\n", funcname)

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
	modRAFlowMeta = raFlowData.Meta

	// --------------------------------------------------------
	// APPEND FEES FOR PETS
	// --------------------------------------------------------
	var newRAFlowPet rlib.RAPetsFlowData

	// There is no associated contact person for the brand new entry of this pet
	RID := int64(0)

	// get new entry for pet with some initial data
	newRAFlowPet, err = rlib.NewRAFlowPet(ctx, d.BID, RID,
		raFlowData.Dates.RentStart, raFlowData.Dates.RentStop,
		raFlowData.Dates.PossessionStart, raFlowData.Dates.PossessionStop,
		&modRAFlowMeta)
	if err != nil {
		return
	}

	// append in pets list
	raFlowData.Pets = append(raFlowData.Pets, newRAFlowPet)

	// --------------------------------------------------------
	// UPDATE JSON DOC WITH NEW PET DATA (BLANK)
	// --------------------------------------------------------
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

	// --------------------------------------------------------
	// UPDATE JSON DOC WITH NEW META DATA IF APPLICABLE
	// --------------------------------------------------------
	if raFlowData.Meta.LastTMPASMID < modRAFlowMeta.LastTMPASMID {

		// Update HavePets Flag in meta information of flow
		modRAFlowMeta.HavePets = len(raFlowData.Pets) > 0

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
	// RETURN RESPONSE
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

	// -------------------
	// WRITE FLOW RESPONSE
	// -------------------
	SvcWriteFlowResponse(ctx, d.BID, flow, w)
	return
}

// SvcPetFeesHandler is used to get the pet fees based on provided command
//
// URL:
//       0    1       2   3
//      /v1/petfees/BID/TCID
// The server command can be:
//      recalculate
//-----------------------------------------------------------------------------
func SvcPetFeesHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcPetFeesHandler"
	var (
		err error
	)
	fmt.Printf("Entered in %s\n", funcname)

	switch d.wsSearchReq.Cmd {
	case "recalculate":
		RecalculatePetFees(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// RecalculatePetFeeRequest struct to handle
type RecalculatePetFeeRequest struct {
	FlowID   int64
	TMPPETID int64
	RID      int64
}

// RecalculatePetFees re-calculate pet fees and make changes in flow json if required
// wsdoc {
//  @Title  Recalculate Pet Fees
//  @URL /v1/petfees/:BID/:FlowID
//  @Method  POST
//  @Synopsis recalculate pet fees
//  @Description returns flow doc with modification in pet fees
//  @Input
//  @Response FlowResponse
// wsdoc }
func RecalculatePetFees(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "RecalculatePetFees"
	var (
		req        RecalculatePetFeeRequest
		raFlowData rlib.RAFlowJSONData
		err        error
		tx         *sql.Tx
		ctx        context.Context
	)
	fmt.Printf("Entered in %s\n", funcname)

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
	if err = json.Unmarshal([]byte(d.data), &req); err != nil {
		return
	}

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err = rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		return
	}

	//-------------------------------------------------------
	// GET THE FLOW FROM FlowID AND UNMARSHALL CONTENT
	//-------------------------------------------------------
	var flow rlib.Flow
	flow, err = rlib.GetFlow(ctx, req.FlowID)
	if err != nil {
		return
	}

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		return
	}

	// START AND STOP IN time.Time
	rStart := (time.Time)(raFlowData.Dates.RentStart)
	rStop := (time.Time)(raFlowData.Dates.RentStop)
	modRAFlowMeta := raFlowData.Meta

	//----------------------------
	// GET THE CYCLES BASED ON RID
	//----------------------------
	var RC, PC int64 // RC = Rent Cycle, PC = Proration Cycle
	err = rlib.GetRAFlowFeeCycles(ctx, req.RID, rStart, &RC, &PC)
	if err != nil {
		return
	}

	// CHECK IF ABOVE RC AND PC ARE SAME OR NOT, IF NOT THEN GET INITIAL FEES AGAIN
	petIndex := -1
	for i := range raFlowData.Pets {
		if req.TMPPETID == raFlowData.Pets[i].TMPPETID {
			petIndex = i
			break
		}
	}

	// LOOK AT THE FIRST FEE INSTANCE IN FEE
	cyclesChanges := true
	if len(raFlowData.Pets[petIndex].Fees) > 0 {
		// AS ONLY RENTCYCLE IS AVAILABLE IN FLOW FEES DATA
		if raFlowData.Pets[petIndex].Fees[0].RentCycle == RC {
			cyclesChanges = false
		}
	}

	// IF CYCLES ARE CHANGED THEN GET FEES AND RE-ASSIGN
	if cyclesChanges {
		raFlowData.Pets[petIndex].Fees, err = rlib.GetRAFlowInitialPetFees(
			ctx, d.BID, req.RID, rStart, rStop, &modRAFlowMeta)

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

	// --------------------------------------------------------
	// UPDATE JSON DOC WITH NEW META DATA IF APPLICABLE
	// --------------------------------------------------------
	if raFlowData.Meta.LastTMPASMID < modRAFlowMeta.LastTMPASMID {

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
	// RETURN RESPONSE
	// ----------------------------------------------

	// get the modified flow
	flow, err = rlib.GetFlow(ctx, req.FlowID)
	if err != nil {
		return
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err = tx.Commit(); err != nil {
		return
	}

	// -------------------
	// WRITE FLOW RESPONSE
	// -------------------
	SvcWriteFlowResponse(ctx, d.BID, flow, w)
	return
}
