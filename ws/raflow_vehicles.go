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

// NewRAFlowVehicle create new vehicle entry for the raflow and returns strcture
// with fees configured it in bizprops
func NewRAFlowVehicle(ctx context.Context, BID int64, meta *RAFlowMetaInfo) (vehicle RAVehiclesFlowData, err error) {
	const funcname = "NewRAFlowVehicle"
	var (
		today = time.Now()
	)
	fmt.Printf("Entered in %s\n", funcname)

	// initialize
	// assign new TMPVID & mark in meta info
	meta.LastTMPVID++
	vehicle = RAVehiclesFlowData{
		Fees:    []RAFeesData{},
		TMPVID:  meta.LastTMPVID,
		DtStart: rlib.JSONDate(today),
		DtStop:  rlib.JSONDate(today.AddDate(1, 0, 0)),
	}

	// get vehicle fees data and feed into fees
	var vehicleFees []rlib.BizPropsVehicleFee
	vehicleFees, err = rlib.GetVehicleFeesFromGeneralBizProps(ctx, BID)
	if err != nil {
		return
	}

	// loop over fees
	for _, fee := range vehicleFees {
		meta.LastTMPASMID++ // new asm id temp
		vf := RAFeesData{
			ARID:           fee.ARID,
			ARName:         fee.ARName,
			ContractAmount: fee.Amount,
			TMPASMID:       meta.LastTMPASMID,
		}

		// append fee for this vehicle
		vehicle.Fees = append(vehicle.Fees, vf)
	}

	return
}

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
		g             FlowResponse
		foo           RAFlowNewVehicleRequest
		raFlowData    RAFlowJSONData
		err           error
		tx            *sql.Tx
		ctx           context.Context
		modRAFlowMeta RAFlowMetaInfo
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
	// APPEND FEES FOR VEHICLES
	// --------------------------------------------------------
	var newRAFlowVehicle RAVehiclesFlowData
	newRAFlowVehicle, err = NewRAFlowVehicle(r.Context(), d.BID, &modRAFlowMeta)
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
	// UPDATE JSON DOC WITH NEW META DATA
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

	// set the response
	g.Record = flow
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
