package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"rentroll/rlib"
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
		foo        RAFlowNewPetRequest
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

	// --------------------------------------------------------
	// APPEND FEES FOR PETS
	// --------------------------------------------------------
	var newRAFlowPet rlib.RAPetsFlowData

	// get new entry for pet with some initial data
	newRAFlowPet, err = rlib.NewRAFlowPet(ctx, d.BID,
		raFlowData.Dates.RentStart, raFlowData.Dates.RentStop,
		raFlowData.Dates.PossessionStart, raFlowData.Dates.PossessionStop,
		&raFlowData.Meta)
	if err != nil {
		return
	}

	// append in pets list
	raFlowData.Pets = append(raFlowData.Pets, newRAFlowPet)

	// Update HavePets Flag in meta information of flow
	raFlowData.Meta.HavePets = len(raFlowData.Pets) > 0

	// LOOK FOR DATA CHANGES
	var originData rlib.RAFlowJSONData
	err = json.Unmarshal(flow.Data, &originData)
	if err != nil {
		return
	}

	// IF THERE ARE DATA CHANGES THEN ONLY UPDATE THE FLOW
	if !reflect.DeepEqual(originData, raFlowData) {
		// GET JSON DATA FROM THE STRUCT
		var modFlowData []byte
		modFlowData, err = json.Marshal(&raFlowData)
		if err != nil {
			return
		}

		// ASSIGN JSON MARSHALLED MODIFIED DATA
		flow.Data = modFlowData

		// NOW UPDATE THE WHOLE FLOW
		err = rlib.UpdateFlowWithInitState(ctx, &flow)
		if err != nil {
			return
		}

		// get the modified flow
		flow, err = rlib.GetFlow(ctx, flow.FlowID)
		if err != nil {
			return
		}
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
