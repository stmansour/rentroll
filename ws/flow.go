package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// FlowPartJSONData is a struct in which request data will be get
type FlowPartJSONData struct {
	FlowPartID     int64           `json:"FlowPartID"`
	BID            int64           `json:"BID"`
	Flow           string          `json:"Flow"`
	FlowID         string          `json:"FlowID"`
	PartType       int             `json:"PartType"`
	Data           json.RawMessage `json:"Data"`
	flowPartStruct interface{}
}

// GetFlowResponse response struct to get a whole flow with its all part
type GetFlowResponse struct {
	Status  string          `json:"status"`
	Records []rlib.FlowPart `json:"records"`
}

// FlowRequest holds flowtype string to get flow
type FlowRequest struct {
	Flow string `json:"Flow"`
}

// FlowIDRequest holds the flowID to get/insert/delete flow
type FlowIDRequest struct {
	FlowID string `json:"FlowID"`
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

	// if d.ID, err = SvcExtractIDFromURI(r.RequestURI, "FlowID", 3, w); err != nil {
	// 	SvcErrorReturn(w, err, funcname)
	// 	return
	// }

	rlib.Console("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getFlow(w, r, d)
		break
	case "init":
		initiateFlow(w, r, d)
		break
	case "delete":
		deleteFlow(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// initiateFlow returns all flowparts associated with given flowID
func initiateFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "initiateFlow"
	var (
		err error
		g   FlowIDRequest // it's not request, but yeah this should feed up the response
		f   FlowRequest
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &f); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// initiate flow for given string
	var flowID string
	switch f.Flow {
	case "RA":
		flowID, err = rlib.InsertInitialRAFlow(r.Context(), d.BID)
		break
	default:
		err = fmt.Errorf("unrecognized flow: %s", f.Flow)
	}

	// if error then return from here
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.FlowID = flowID
	SvcWriteResponse(d.BID, &g, w)
}

// getFlow returns all flowparts associated with given flowID
func getFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getFlow"
	var (
		err error
		g   GetFlowResponse
		f   FlowIDRequest
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &f); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Records, err = rlib.GetFlowPartsByFlowID(r.Context(), f.FlowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// deleteFlow delete the flow from database with associated all flow parts
// for a given flowID
func deleteFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteFlow"
	var (
		err error
		del FlowIDRequest
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// delete flow parts by flowID
	err = rlib.DeleteFlowPartsByFlowID(r.Context(), del.FlowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// GetFlowPartResponse response struct to get flow part
type GetFlowPartResponse struct {
	Status string        `json:"status"`
	Record rlib.FlowPart `json:"record"`
}

// SvcHandlerFlowPart handles operations on a single flow part by given flowID
// For this call, we expect the URI to contain the BID and the FlowPartID as follows:
//           0  1    2   3
// uri      /v1/asm/:BID/FlowPartID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerFlowPart(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerFlowPart"
	var (
		err error
	)

	rlib.Console("Entered %s\n", funcname)

	if d.ID, err = SvcExtractIDFromURI(r.RequestURI, "FlowPartID", 3, w); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Request: %s:  BID = %d,  FlowPartID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getFlowPart(w, r, d)
		break
	case "save":
		saveFlowPart(w, r, d)
		break
	// case "delete":
	//  deleteFlowPart(w, r, d)
	//  break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// getFlowPart will return a flowpart instance for a given flowpartID (d.ID)
func getFlowPart(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getFlowPart"
	var (
		err error
		g   GetFlowPartResponse
	)

	rlib.Console("Entered %s\n", funcname)
	// rlib.Console("record data = %s\n", d.data)

	// get flow part by its ID
	g.Record, err = rlib.GetFlowPart(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// saveFlowPart will save (update) a flowpart instance for a given flowpartID (d.ID)
func saveFlowPart(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveFlowPart"
	var (
		err        error
		fpJSONData FlowPartJSONData
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &fpJSONData); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// handle data for update based on flow and part type
	var validDataFormat bool
	switch fpJSONData.Flow {
	case "RA":
		validDataFormat = isValidUpdateRAFlowPartJSONData(fpJSONData.Data, fpJSONData.PartType)
	default:
		err = fmt.Errorf("unrecognized flow: %s", fpJSONData.Flow)
		SvcErrorReturn(w, err, funcname)
		return
	}

	// check if any error occured in previous switch case
	if !validDataFormat {
		err = fmt.Errorf("Data is not in valid format for flow: %s, partType: %d", fpJSONData.Flow, fpJSONData.PartType)
		SvcErrorReturn(w, err, funcname)
		return
	}

	// migrate request json data to flow part struct
	var fp rlib.FlowPart
	rlib.MigrateStructVals(&fpJSONData, &fp) // the variables that don't need special handling
	fp.Data = []byte(fpJSONData.Data)

	// get flow part by its ID
	err = rlib.UpdateFlowPart(r.Context(), &fp)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}
