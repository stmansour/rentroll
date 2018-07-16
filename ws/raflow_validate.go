package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
)

// RAFlowDetailRequest is a struct to hold info for Flow which is going to be validate
type RAFlowDetailRequest struct {
	FlowID    int64
	UserRefNo string
}

// SvcValidateRAFlow is used to check/validate RAFlow's struct
//------------------------------------------------------------------------------
func SvcValidateRAFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcValidateRAFlow"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		ValidateRAFlow(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// ValidateRAFlow validate RAFlow's fields section wise
//-------------------------------------------------------------------------
func ValidateRAFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "ValidateRAFlow"
	fmt.Printf("Entered %s\n", funcname)

	var (
		err        error
		foo        RAFlowDetailRequest
		raFlowData rlib.RAFlowJSONData
		g          bizlogic.ValidateRAFlowResponse
	)

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("only POST method is allowed")
		return
	}

	// unmarshal data into request data struct
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		return
	}

	// Get flow information from the table to validate fields value
	flow, err := rlib.GetFlow(r.Context(), foo.FlowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// When flowId doesn't exists in database return and give error that flowId doesn't exists
	if flow.FlowID == 0 {
		err = fmt.Errorf("flowID %d - doesn't exists", foo.FlowID)
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// ---------------------------------------
	// Perform basic validation on RAFlow
	// ---------------------------------------
	// TODO(Akshay): Enable basic validation check
	bizlogic.ValidateRAFlowBasic(r.Context(), &raFlowData, &g)

	// If RAFlow structure have more than 1 basic validation error than it return with the list of basic validation errors
	if g.Total > 0 {
		SvcWriteResponse(d.BID, &g, w)
		return
	}

	// --------------------------------------------
	// Perform Bizlogic check validation on RAFlow
	// --------------------------------------------
	bizlogic.ValidateRAFlowBizLogic(r.Context(), &raFlowData, &g)

	// If RAFlow structure have more than 1 biz logic check validation error than it return with the list of biz logic validation errors
	if g.Total > 0 {
		SvcWriteResponse(d.BID, &g, w)
		return
	}

	SvcWriteResponse(d.BID, &g, w)
}
