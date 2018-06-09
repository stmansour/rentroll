package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
)

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
	if !(d.wsSearchReq.Cmd == "getAllFlows" || d.wsSearchReq.Cmd == "init") {
		if d.ID, err = SvcExtractIDFromURI(r.RequestURI, "FlowID", 3, w); err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	}

	rlib.Console("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "getAllFlows":
		getAllFlowsByUser(w, r, d)
		break
	case "get":
		getFlow(w, r, d)
		break
	case "init":
		initiateFlow(w, r, d)
		break
	case "delete":
		deleteFlow(w, r, d)
		break
	case "save":
		saveFlow(w, r, d)
		break
	case "migrate":
		migrateFlowDataToDB(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// FlowResponse is the response of returning updated flow with status
type FlowResponse struct {
	Record rlib.Flow `json:"record"`
	Status string    `json:"status"`
}

// initiateFlow inserts a new flow with default data and returns it
func initiateFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "initiateFlow"

	var (
		err error
		f   struct {
			FlowType string
		}
		g FlowResponse
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &f); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// initiate flow for given string
	var flowID int64
	switch f.FlowType {
	case rlib.RAFlow:
		flowID, err = insertInitialRAFlow(r.Context(), d.BID, d.sess.UID)
		break
	default:
		err = fmt.Errorf("unrecognized flowType: %s", f.FlowType)
	}

	// if error then return from here
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get flow data in return it back
	flow, err := rlib.GetFlow(r.Context(), flowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Record = flow
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// getFlow returns flow associated with given flowID
func getFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getFlow"
	var (
		err error
		f   struct {
			FlowID int64
		}
		g FlowResponse
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &f); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	flow, err := rlib.GetFlow(r.Context(), f.FlowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Record = flow
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// getAllFlowsByUser returns all flows for the current user and given flow
func getAllFlowsByUser(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getAllFlowsByUser"

	var (
		err error
		f   struct {
			FlowType string
		}
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &f); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get all flowIDs
	// recs, err := rlib.GetFlowIDsByUser(r.Context())
	m, err := rlib.GetFlowMetaDataInRange(r.Context(), &d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// initiate flow for given string
	switch f.FlowType {
	case rlib.RAFlow:
		// ResponseData response data for ra flow
		var g struct {
			Status  string               `json:"status"`
			Records []GridRAFlowResponse `json:"records"`
		}
		g.Records = []GridRAFlowResponse{}
		for i := 0; i < len(m); i++ {
			var t = GridRAFlowResponse{
				Recid:     int64(i),
				BID:       d.BID,
				FlowID:    m[i].FlowID,
				UserRefNo: m[i].UserRefNo,
				BUD:       string(rlib.GetBUDFromBIDList(d.BID)),
			}
			g.Records = append(g.Records, t)
		}
		g.Status = "success"
		SvcWriteResponse(d.BID, &g, w)
		return
	default:
		err = fmt.Errorf("unrecognized flow type: %s", f.FlowType)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// deleteFlow delete the flow from database with associated all flow parts
// for a given flowID
func deleteFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteFlow"
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

// saveFlow will save (update) a flowpart instance for a given flowpartID (d.ID)
func saveFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveFlow"
	var (
		err     error
		flowReq struct {
			FlowType    string
			FlowID      int64
			Data        json.RawMessage
			FlowPartKey string
		}
		g FlowResponse
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &flowReq); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// check that such flow instance does exist or not
	existFlow, _ := rlib.GetFlow(r.Context(), flowReq.FlowID)
	if existFlow.FlowID == 0 {
		err = fmt.Errorf("given flow with ID (%d) doesn't exist", flowReq.FlowID)
		SvcErrorReturn(w, err, funcname)
		return
	}

	// handle data for update based on flow and part type
	var jsBtData, modMetaInfo []byte
	switch flowReq.FlowType {
	case rlib.RAFlow:
		partType, ok := rlib.RAFlowPartsMap[flowReq.FlowPartKey]
		if !ok {
			err = fmt.Errorf("Unable to find part with key: %s for flowID: %d, flowType: %s, Error: %s", flowReq.FlowPartKey, flowReq.FlowID, flowReq.FlowType, err.Error())
			SvcErrorReturn(w, err, funcname)
			return
		}

		modMetaInfo, jsBtData, err = getUpdateRAFlowPartJSONData(d.BID, flowReq.Data, int(partType), &existFlow)
		if err != nil {
			err1 := fmt.Errorf("Data is not in valid format for flowID: %d, flowType: %s, Error: %s", flowReq.FlowID, flowReq.FlowType, err.Error())
			SvcErrorReturn(w, err1, funcname)
			return
		}
	default:
		err = fmt.Errorf("unrecognized flow type: %s", flowReq.FlowType)
		SvcErrorReturn(w, err, funcname)
		return
	}

	// update data with given json data key
	err = rlib.UpdateFlowData(r.Context(), flowReq.FlowPartKey, jsBtData, &existFlow)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// update data for modified meta data
	err = rlib.UpdateFlowData(r.Context(), "meta", modMetaInfo, &existFlow)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get flow data in return it back
	flow, err := rlib.GetFlow(r.Context(), existFlow.FlowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Record = flow
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// migrateFlowDataToDB saves the data from temp data stored in flowPart with flowID into actual
// database instance for the given flow type
func migrateFlowDataToDB(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveFlow"
	var (
		err error
		f   struct {
			FlowType string
			FlowID   int64
		}
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &f); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	switch f.FlowType {
	case rlib.RAFlow:
		_, err = saveRentalAgreementFlow(r.Context(), f.FlowID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		break
	default:
		err = fmt.Errorf("unrecognized flow type: %s", f.FlowType)
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}
