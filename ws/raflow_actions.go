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

// RAActionDataRequest is a struct to hold info about actions taken on Rental Agreement
type RAActionDataRequest struct {
	FlowID               int64         // Flow ID of Rental Agreement
	Action               int64         // If '-1' then Do nothing
	Decision1            int64         // 0 - NoApprovalDecision1Field, 1 - RA Approved, 2 - RA Declined
	DeclineReason1       int64         // 0 - NoDeclineReason1Field, great than 0 - Reason for Decline
	Decision2            int64         // 0 - NoApprovalDecision2Field, 1 - RA Approved, 2 - RA Declined
	DeclineReason2       int64         // 0 - NoDeclineReason2Field, great than 0 - Reason for Decline
	TerminationReason    int64         // 0 - NoTerminationReasonField, great than 0 - Reason for Termination
	DocumentDate         rlib.JSONDate // date when rental agreement was signed
	NoticeToMoveDate     rlib.JSONDate // date RA was given Notice-To-Move
	NoticeToMoveReported rlib.JSONDate // date RA was set to Terminated because of moving out
	Mode                 string        // It represents that which button submitted the form
}

// SvcSetRAState sets the state of Rental Agreement and updates meta info of RA
func SvcSetRAState(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSetRAState"
	var (
		g          FlowResponse
		raFlowData rlib.RAFlowJSONData
		foo        RAActionDataRequest
		today      = time.Now()
		err        error
		tx         *sql.Tx
		ctx        context.Context
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

	MODE := foo.Mode
	state := raFlowData.Meta.RAFLAGS & ^(0xf)

	switch MODE {
	case "Action":
		switch foo.Action {
		case 0: // Application Being Completed
			modRAFlowMeta.RAFLAGS = (state | 0)

		case 1: // Set To First Approval
			modRAFlowMeta.RAFLAGS = (state | 1)

		case 2: // Set To Second Approval
			modRAFlowMeta.RAFLAGS = (state | 2)

		case 3: // Set To Move-In
			modRAFlowMeta.RAFLAGS = (state | 3)

		case 4: // Complete Move-In
			modRAFlowMeta.RAFLAGS = (state | 4)
		case 5: // Terminate
			if foo.TerminationReason > 0 {
				modRAFlowMeta.TerminatorUID = d.sess.UID
				modRAFlowMeta.TerminationDate = rlib.JSONDate(today)
				modRAFlowMeta.LeaseTerminationReason = foo.TerminationReason

				modRAFlowMeta.RAFLAGS = (state | 5)
			} else {
				// return err
				err := fmt.Errorf("Termination Reason not present")
				SvcErrorReturn(w, err, funcname)
				return
			}

		case 6: // Notice-To-Move
			modRAFlowMeta.NoticeToMoveDate = foo.NoticeToMoveDate
			modRAFlowMeta.NoticeToMoveReported = foo.NoticeToMoveReported

			modRAFlowMeta.RAFLAGS = (state | 6)

		default:
		}
	}

	/*switch foo.Action {
	case 0: // Application Being Completed
		modRAFlowMeta.RAFLAGS = (state | 0)

	case 1: // Set To First Approval
		modRAFlowMeta.RAFLAGS = (state | 1)

	case 2: // Set To Second Approval
		modRAFlowMeta.RAFLAGS = (state | 2)

	case 3: // Set To Move-In
		modRAFlowMeta.RAFLAGS = (state | 3)

	case 4: // Complete Move-In
		modRAFlowMeta.RAFLAGS = (state | 4)

	case 5: // Terminate
		modRAFlowMeta.RAFLAGS = (state | 5)

	case 6: // Recieved Notice-To-Move
		modRAFlowMeta.RAFLAGS = (state | 6)

	default:
	}*/

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
