package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// RAActionDataRequest is a struct to hold info about actions taken on Rental Agreement
type RAActionDataRequest struct {
	FlowID   int64 // Flow ID of Rental Agreement
	Decision int64 // 0 - NoApprovalField, 1 - RA Approved, 2 - RA Declined
	Reason   int64 // 0 - NoReasonField, great than 0 - Reason for Decline or Terminate
	Action   int64 // If '-1' then Do nothing

}

// SvcSetRAState sets the state of Rental Agreement and updates meta info of RA
func SvcSetRAState(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSetRAState"
	var (
		g          FlowResponse
		raFlowData rlib.RAFlowJSONData
		foo        RAActionDataRequest
		// today      = time.Now()
		err error
		tx  *sql.Tx
		ctx context.Context
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

	// HTTP METHOD CHECK
	if r.Method != "POST" {
		err := fmt.Errorf("Only POST method is allowed")
		SvcErrorReturn(w, err, funcname)
		return
	}

	// SEE IF WE CAN UNMARSHAL THE DATA
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		SvcErrorReturn(w, err, funcname)
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
	// FLOW EXISTENCE CHECK
	//-------------------------------------------------------
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

	// get meta in modRAFlowMeta, we're going to modify it
	modRAFlowMeta := raFlowData.Meta

	state := raFlowData.Meta.RAFLAGS & ^(0xf)

	switch foo.Action {
	case 0: // Edit Rental Agreement Information
		modRAFlowMeta.RAFLAGS = (state | 0)

	case 1: // Authorize First Approval
		modRAFlowMeta.RAFLAGS = (state | 2)

	case 2: // Authorize Second Approval
		modRAFlowMeta.RAFLAGS = (state | 3)

	case 3: // Complete Move In
		modRAFlowMeta.RAFLAGS = (state | 4)

	case 4: // Terminate
		modRAFlowMeta.RAFLAGS = (state | 5)

	case 5: // Recieved Notice To Move
		modRAFlowMeta.RAFLAGS = (state | 6)

	default:
	}

	//-------------------------------------------------------
	// MODIFY META DATA TOO
	//-------------------------------------------------------
	var modMetaData []byte
	modMetaData, err = json.Marshal(&modRAFlowMeta)
	if err != nil {
		return
	}

	err = rlib.UpdateFlowData(ctx, "meta", modMetaData, &flow)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
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
