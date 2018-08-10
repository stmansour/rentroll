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

// SvcRAFlowVehiclesHandler handles operation on vehicles of raflow json data
//           0    1     2   3
// uri /v1/raflow-vehicles/BID/flowID
// The server command can be:
//      new
//-----------------------------------------------------------------------------------
func SvcRAFlowVehiclesHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRAFlowVehiclesHandler"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "new":
		CreateNewRAFlowVehicle(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// RAFlowNewVehicleRequest struct for new vehicle request
type RAFlowNewVehicleRequest struct {
	FlowID int64
}

// CreateNewRAFlowVehicle creates new entry in vehicles list of raflow json data
// wsdoc {
//  @Title create brand new vehicle entry in vehicle list
//  @URL /v1/raflow-vehicles/:BUI/:FlowID
//  @Method POST
//  @Synopsis create new entry of vehicle in RAFlow json data
//  @Description create new entry of vehicle in RAFlow json data with all available fees
//  @Input RAFlowNewVehicleRequest
//  @Response FlowResponse
// wsdoc }
//-----------------------------------------------------------------------------------
func CreateNewRAFlowVehicle(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "CreateNewRAFlowVehicle"
	var (
		foo        RAFlowNewVehicleRequest
		raFlowData = rlib.RAFlowJSONData{}
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

	// --------------------------------------------------------
	// APPEND FEES FOR VEHICLES
	// --------------------------------------------------------
	var newRAFlowVehicle rlib.RAVehiclesFlowData

	// get new vehicle with some initial data
	newRAFlowVehicle, err = rlib.NewRAFlowVehicle(ctx, d.BID,
		raFlowData.Dates.RentStart, raFlowData.Dates.RentStop,
		raFlowData.Dates.PossessionStart, raFlowData.Dates.PossessionStop,
		&raFlowData.Meta)
	if err != nil {
		return
	}

	// append in vehicles list
	raFlowData.Vehicles = append(raFlowData.Vehicles, newRAFlowVehicle)

	// Update HaveVehicles Flag in meta information of flow
	raFlowData.Meta.HaveVehicles = len(raFlowData.Vehicles) > 0

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
		err = rlib.UpdateRAFlowWithInitState(ctx, &flow)
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
