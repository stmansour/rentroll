package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
)

// FlowResponse is the response of returning updated flow with status
type FlowResponse struct {
	Record interface{} `json:"record"`
	Status string      `json:"status"`
}

// SvcHandlerFlow handles operations on a whole flow which affects on its
// all flow parts associated with given flowID
// For this call, we expect the URI to contain the BID and the FlowID as follows:
//           0  1    2   3
// uri      /v1/asm/:BID/FlowID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerFlow"
	var (
		err error
	)

	rlib.Console("Entered %s\n", funcname)

	// if the command is not from listed below then do check for flowID
	if !(d.wsSearchReq.Cmd == "all" || d.wsSearchReq.Cmd == "init") {
		if d.ID, err = SvcExtractIDFromURI(r.RequestURI, "FlowID", 3, w); err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	}

	rlib.Console("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "all":
		GetAllFlowsByType(w, r, d)
		break
	case "get":
		GetFlowByType(w, r, d)
		break
	case "init":
		InitiateFlow(w, r, d)
		break
	case "delete":
		DeleteFlow(w, r, d)
		break
	case "save":
		SaveFlow(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// InitiateFlow inserts a new flow with default data and returns it
func InitiateFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "InitiateFlow"

	var (
		err error
		req struct {
			FlowType string
		}
		flow rlib.Flow
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &req); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// initiate flow for given string
	var flowID int64
	switch req.FlowType {
	case rlib.RAFlow:
		flowID, err = rlib.InsertInitialRAFlow(r.Context(), d.BID, d.sess.UID)
		break
	default:
		err = fmt.Errorf("unrecognized flowType: %s", req.FlowType)
	}

	// if error then return from here
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get flow data in return it back
	flow, err = rlib.GetFlow(r.Context(), flowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// -------------------
	// WRITE FLOW RESPONSE
	// -------------------
	SvcWriteFlowResponse(r.Context(), d.BID, flow, w)
	return
}

// FlowTypeRequest struct
type FlowTypeRequest struct {
	FlowType string
}

// GetAllFlowsByType returns all flows for the current user and given flow
func GetAllFlowsByType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "GetAllFlowsByType"

	var (
		err error
		f   FlowTypeRequest
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &f); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// initiate flow for given string
	switch f.FlowType {
	case rlib.RAFlow:
		GetAllRAFlows(w, r, d)
		return
	default:
		err = fmt.Errorf("unrecognized flow type: %s", f.FlowType)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// GetFlowByType returns flow associated with given flowID
func GetFlowByType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "GetFlowByType"
	var (
		err error
		req FlowTypeRequest
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &req); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// initiate flow for given string
	switch req.FlowType {
	case rlib.RAFlow:
		GetRAFlow(w, r, d)
		return
	default:
		err = fmt.Errorf("unrecognized flow type: %s", req.FlowType)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// DeleteFlow delete the flow from database with associated all flow parts
// for a given flowID
func DeleteFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "DeleteFlow"
	var (
		err error
		del struct {
			FlowID int64
		}
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// delete flow parts by flowID
	err = rlib.DeleteFlow(r.Context(), del.FlowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
	return
}

// SaveFlowRequest struct
type SaveFlowRequest struct {
	FlowType    string
	FlowID      int64
	Data        json.RawMessage
	FlowPartKey string
}

// SaveFlow will save (update) a flowpart instance for a given flowpartID (d.ID)
func SaveFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SaveFlow"
	var (
		err     error
		flowReq SaveFlowRequest
		tx      *sql.Tx
		ctx     context.Context
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

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

	// ------- unmarshal the request data  ---------------
	if err = json.Unmarshal([]byte(d.data), &flowReq); err != nil {
		return
	}

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err = rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		return
	}

	// check that such flow instance does exist or not
	flow, _ := rlib.GetFlow(ctx, flowReq.FlowID)
	if flow.FlowID == 0 {
		err = fmt.Errorf("given flow with ID (%d) doesn't exist", flowReq.FlowID)
		return
	}

	// handle data for update based on flow and part type
	switch flowReq.FlowType {
	case rlib.RAFlow:
		err = rlib.UpdateRAFlowJSON(ctx, d.BID, flowReq.Data, flowReq.FlowPartKey, &flow)
		if err != nil {
			return
		}
	default:
		err = fmt.Errorf("unrecognized flow type: %s", flowReq.FlowType)
		return
	}

	// ----------------------------------------------
	// RETURN RESPONSE
	// ----------------------------------------------

	// get flow data in return it back
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

// SvcWriteFlowResponse writes response in w http.ResponseWrite especially for flow data
func SvcWriteFlowResponse(ctx context.Context, BID int64, flow rlib.Flow, w http.ResponseWriter) {
	const funcname = "SvcWriteFlowResponse"
	var (
		err            error
		raFlowData     rlib.RAFlowJSONData
		resp           FlowResponse
		raflowRespData = RAFlowResponse{Flow: flow}
	)
	fmt.Printf("Entered in %s\n", funcname)

	// GET UNMARSHALLED RAFLOW DATA INTO STRUCT
	err = json.Unmarshal(raflowRespData.Flow.Data, &raFlowData)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// PERFORM BASIC VALIDATION ON FLOW DATA
	bizlogic.ValidateRAFlowBasic(ctx, &raFlowData, &raflowRespData.BasicCheck)

	// CHECK DATA FULFILLED
	bizlogic.DataFulfilledRAFlow(ctx, &raFlowData, &raflowRespData.DataFulfilled)

	resp.Record = raflowRespData
	resp.Status = "success"
	SvcWriteResponse(BID, &resp, w)
}
