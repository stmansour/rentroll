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
		foo           RAFlowNewVehicleRequest
		raFlowData    = rlib.RAFlowJSONData{}
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
	modRAFlowMeta = raFlowData.Meta

	// --------------------------------------------------------
	// APPEND FEES FOR VEHICLES
	// --------------------------------------------------------
	var newRAFlowVehicle rlib.RAVehiclesFlowData

	// There is no associated contact person for the brand new entry of this vehicle
	RID := int64(0)

	// get new vehicle with some initial data
	newRAFlowVehicle, err = rlib.NewRAFlowVehicle(ctx, d.BID, RID,
		raFlowData.Dates.RentStart, raFlowData.Dates.RentStop,
		raFlowData.Dates.PossessionStart, raFlowData.Dates.PossessionStop,
		&modRAFlowMeta)
	if err != nil {
		return
	}

	// append in vehicles list
	raFlowData.Vehicles = append(raFlowData.Vehicles, newRAFlowVehicle)

	// --------------------------------------------------------
	// UPDATE JSON DOC WITH NEW VEHICLE DATA (BLANK)
	// --------------------------------------------------------
	var modVehiclesData []byte
	modVehiclesData, err = json.Marshal(&raFlowData.Vehicles)
	if err != nil {
		return
	}

	// update flow with this modified vehicles part
	err = rlib.UpdateFlowData(ctx, "vehicles", modVehiclesData, &flow)
	if err != nil {
		return
	}

	// --------------------------------------------------------
	// UPDATE JSON DOC WITH NEW META DATA IF APPLICABLE
	// --------------------------------------------------------
	if raFlowData.Meta.LastTMPASMID < modRAFlowMeta.LastTMPASMID {

		// Update HaveVehicles Flag in meta information of flow
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

// SvcVehicleFeesHandler is used to get the vehicle fees based on provided command
// URL:
//       0    1       2   3
//      /v1/vehiclefees/BID/TCID
// The server command can be:
//      recalculate
//-----------------------------------------------------------------------------
func SvcVehicleFeesHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcVehicleFeesHandler"
	var (
		err error
	)
	fmt.Printf("Entered in %s\n", funcname)

	switch d.wsSearchReq.Cmd {
	case "recalculate":
		RecalculateVehicleFees(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// RecalculateVehicleFeeRequest struct to handle
type RecalculateVehicleFeeRequest struct {
	FlowID int64
	TMPVID int64
	RID    int64
}

// RecalculateVehicleFees re-calculate vehicle fees and make changes in flow json if required
// wsdoc {
//  @Title  Recalculate Vehicle Fees
//  @URL /v1/vehiclefees/:BID/:FlowID
//  @Method  POST
//  @Synopsis recalculate vehicle fees
//  @Description returns flow doc with modification in vehicle fees
//  @Input
//  @Response FlowResponse
// wsdoc }
func RecalculateVehicleFees(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "RecalculateVehicleFees"
	var (
		req        RecalculateVehicleFeeRequest
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
	vehicleIndex := -1
	for i := range raFlowData.Vehicles {
		if req.TMPVID == raFlowData.Vehicles[i].TMPVID {
			vehicleIndex = i
			break
		}
	}

	// LOOK AT THE FIRST FEE INSTANCE IN FEE
	cyclesChanges := true
	if len(raFlowData.Vehicles[vehicleIndex].Fees) > 0 {
		// AS ONLY RENTCYCLE IS AVAILABLE IN FLOW FEES DATA
		if raFlowData.Vehicles[vehicleIndex].Fees[0].RentCycle == RC {
			cyclesChanges = false
		}
	}

	// IF CYCLES ARE CHANGED THEN GET FEES AND RE-ASSIGN
	if cyclesChanges {
		raFlowData.Vehicles[vehicleIndex].Fees, err = rlib.GetRAFlowInitialVehicleFees(
			ctx, d.BID, req.RID, rStart, rStop, &modRAFlowMeta)

		var modVehiclesData []byte
		modVehiclesData, err = json.Marshal(&raFlowData.Vehicles)
		if err != nil {
			return
		}

		// update flow with this modified vehicles part
		err = rlib.UpdateFlowData(ctx, "vehicles", modVehiclesData, &flow)
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
