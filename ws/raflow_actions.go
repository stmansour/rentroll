package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"time"
)

// RAActionDataRequest is a struct to hold info about actions taken on Rental Agreement
type RAActionDataRequest struct {
	FlowID int64  // Flow ID of Rental Agreement
	Action int64  // If '-1' then Do nothing
	Mode   string // It represents that which button submitted the form
}

// RAApprover1Data is a struct to hold data about actions of Approver1
type RAApprover1Data struct {
	Decision1      int64 // 0 - NoApprovalDecision1Field, 1 - RA Approved, 2 - RA Declined
	DeclineReason1 int64 // 0 - NoDeclineReason1Field, great than 0 - Reason for Decline
}

// RAApprover2Data is a struct to hold data about actions of Approver2
type RAApprover2Data struct {
	Decision2      int64 // 0 - NoApprovalDecision2Field, 1 - RA Approved, 2 - RA Declined
	DeclineReason2 int64 // 0 - NoDeclineReason2Field, great than 0 - Reason for Decline
}

// RAMoveInData is a struct to hold data about Move In
type RAMoveInData struct {
	DocumentDate rlib.JSONDateTime // date when rental agreement was signed
}

// RATerminationData is a struct to hold data about termination of RentalAgreement
type RATerminationData struct {
	TerminationReason int64 // 0 - NoTerminationReasonField, great than 0 - Reason for Termination
}

// RANoticeToMoveData is a struct to hold data about Notice to Move
type RANoticeToMoveData struct {
	NoticeToMoveDate rlib.JSONDateTime // date RA was given Notice-To-Move
}

// SvcSetRAState sets the state of Rental Agreement and updates meta info of RA
func SvcSetRAState(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSetRAState"
	var (
		g              FlowResponse
		raFlowResponse RAFlowResponse
		raFlowData     rlib.RAFlowJSONData
		foo            RAActionDataRequest
		err            error
		tx             *sql.Tx
		ctx            context.Context
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

	// set location for time as UTC
	var location *time.Location
	location, err = time.LoadLocation("UTC")
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get current time in UTC
	var today time.Time
	today = time.Now().In(location)

	// HTTP METHOD CHECK
	if r.Method != "POST" {
		err = fmt.Errorf("Only POST method is allowed")
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
	state := raFlowData.Meta.RAFLAGS & uint64(0xf)

	clearedState := raFlowData.Meta.RAFLAGS & ^uint64(0xf)

	switch MODE {
	case "Action":
		if foo.Action < int64(state) {
			for i := foo.Action; i <= int64(state); i++ {
				switch i {
				case 0: // Application Being Completed
				case 1: // Pending First Approval
					modRAFlowMeta.Approver1 = 0
					modRAFlowMeta.Approver1Name = ""
					modRAFlowMeta.DeclineReason1 = 0
					modRAFlowMeta.DecisionDate1 = rlib.JSONDateTime(time.Time{})

				case 2: // Pending Second Approval
					modRAFlowMeta.Approver2 = 0
					modRAFlowMeta.Approver2Name = ""
					modRAFlowMeta.DeclineReason2 = 0
					modRAFlowMeta.DecisionDate2 = rlib.JSONDateTime(time.Time{})

				case 3: // Move-In / Execute Modification
					modRAFlowMeta.DocumentDate = rlib.JSONDateTime(time.Time{})
				case 4: // Active
				case 5: // Terminated
					modRAFlowMeta.TerminatorUID = 0
					modRAFlowMeta.TerminatorName = ""
					modRAFlowMeta.LeaseTerminationReason = 0
					modRAFlowMeta.TerminationDate = rlib.JSONDateTime(time.Time{})

				case 6: //Notice To Move
					modRAFlowMeta.NoticeToMoveUID = 0
					modRAFlowMeta.NoticeToMoveName = ""
					modRAFlowMeta.NoticeToMoveDate = rlib.JSONDateTime(time.Time{})
					modRAFlowMeta.NoticeToMoveReported = rlib.JSONDateTime(time.Time{})
				}
			}
		}
		switch foo.Action {
		case 0: // Application Being Completed
			modRAFlowMeta.RAFLAGS = (clearedState | 0)

		case 1: // Set To First Approval
			modRAFlowMeta.RAFLAGS = (clearedState | 1)

		case 2: // Set To Second Approval
			modRAFlowMeta.RAFLAGS = (clearedState | 2)

		case 3: // Set To Move-In
			modRAFlowMeta.RAFLAGS = (clearedState | 3)

		case 4: // Complete Move-In

			// migrate data to real table via hook
			_, err = Flow2RA(ctx, foo.FlowID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			modRAFlowMeta.RAFLAGS = (clearedState | 4)
		case 5: // Terminate
			var data RATerminationData
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			if data.TerminationReason > 0 {
				var fullName string
				fullName, err = getUserFullName(ctx, d.sess.UID)
				if err != nil {
					SvcErrorReturn(w, err, funcname)
					return
				}

				/*RAID := flow.ID
				if RAID > 0 {
					// get the Rental Agreement from database
					var ra rlib.RentalAgreement
					ra, err = rlib.GetRentalAgreement(ctx, RAID)
					if err != nil {
						SvcErrorReturn(w, err, funcname)
						return
					}

					// modify the data
					ra.TerminatorUID = d.sess.UID
					ra.TerminationDate = time.Time(today)
					ra.LeaseTerminationReason = data.TerminationReason

					// update modified data in database
					err = rlib.UpdateRentalAgreement(ctx, &ra)
					if err != nil {
						SvcErrorReturn(w, err, funcname)
						return
					}
				}*/

				modRAFlowMeta.TerminatorUID = d.sess.UID
				modRAFlowMeta.TerminatorName = fullName
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(today)
				modRAFlowMeta.LeaseTerminationReason = data.TerminationReason

				modRAFlowMeta.RAFLAGS = (clearedState | 5)
			} else {
				// return err
				err = fmt.Errorf("termination reason not present")
				SvcErrorReturn(w, err, funcname)
				return
			}

		case 6: // Notice-To-Move
			var data RANoticeToMoveData
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			/*RAID := flow.ID
			if RAID > 0 {
				// get the Rental Agreement from database
				var ra rlib.RentalAgreement
				ra, err = rlib.GetRentalAgreement(ctx, RAID)
				if err != nil {
					SvcErrorReturn(w, err, funcname)
					return
				}

				// modify the data
				ra.NoticeToMoveUID = d.sess.UID
				ra.NoticeToMoveDate = time.Time(data.NoticeToMoveDate)
				ra.NoticeToMoveReported = time.Time(today)

				// update modified data in database
				err = rlib.UpdateRentalAgreement(ctx, &ra)
				if err != nil {
					SvcErrorReturn(w, err, funcname)
					return
				}
			}*/

			modRAFlowMeta.NoticeToMoveUID = d.sess.UID
			modRAFlowMeta.NoticeToMoveName = fullName
			modRAFlowMeta.NoticeToMoveDate = rlib.JSONDateTime(data.NoticeToMoveDate)
			modRAFlowMeta.NoticeToMoveReported = rlib.JSONDateTime(today)

			modRAFlowMeta.RAFLAGS = (clearedState | 6)

		default:
		}
	case "State":
		switch state {
		case 1: // Pending First Approval
			var data RAApprover1Data
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			if data.Decision1 == 1 { // Approved
				// set 4th bit of Flag as 1
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS | uint64(1<<4)

				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 2)
			} else if data.Decision1 == 2 && data.DeclineReason1 > 0 { // Declined
				// set 4th bit of Flag as 0
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS & ^uint64(1<<4)
				modRAFlowMeta.DeclineReason1 = data.DeclineReason1

				modRAFlowMeta.TerminatorUID = d.sess.UID
				modRAFlowMeta.TerminatorName = fullName
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(today)

				//TODO(Jay): Change it with SLSID for "Application Declined"
				modRAFlowMeta.LeaseTerminationReason = 119

				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 5)
			} else {
				// return err
				err = fmt.Errorf("approver1 data invalid")
				SvcErrorReturn(w, err, funcname)
				return
			}

			modRAFlowMeta.Approver1 = d.sess.UID
			modRAFlowMeta.Approver1Name = fullName
			modRAFlowMeta.DecisionDate1 = rlib.JSONDateTime(today)

		case 2: // Pending Second Approval
			var data RAApprover2Data
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			if data.Decision2 == 1 { // Approved
				// set 5th bit of Flag as 1
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS | uint64(1<<5)

				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 3)
			} else if data.Decision2 == 2 && data.DeclineReason2 > 0 { // Declined
				// set 5th bit of Flag as 0
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS & ^uint64(1<<5)
				modRAFlowMeta.DeclineReason2 = data.DeclineReason2

				modRAFlowMeta.TerminatorUID = d.sess.UID
				modRAFlowMeta.TerminatorName = fullName
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(today)

				//TODO(Jay): Change it with SLSID for "Application Declined"
				modRAFlowMeta.LeaseTerminationReason = 119

				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 5)
			} else {
				// return err
				err = fmt.Errorf("approver2 data invalid")
				SvcErrorReturn(w, err, funcname)
				return
			}

			modRAFlowMeta.Approver2 = d.sess.UID
			modRAFlowMeta.Approver2Name = fullName
			modRAFlowMeta.DecisionDate2 = rlib.JSONDateTime(today)

		case 3: // Move-In / Execute Modification
			var data RAMoveInData
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			modRAFlowMeta.DocumentDate = rlib.JSONDateTime(data.DocumentDate)

		default:
		}
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

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// Perform basic validation on flow data
	bizlogic.ValidateRAFlowBasic(r.Context(), &raFlowData, &raFlowResponse.BasicCheck)

	// Check DataFulfilled
	bizlogic.DataFulfilledRAFlow(r.Context(), &raFlowData, &raFlowResponse.DataFulfilled)

	raFlowResponse.Flow = flow
	// set the response
	g.Record = raFlowResponse
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

func getUserFullName(ctx context.Context, UID int64) (string, error) {
	person, err := rlib.GetDirectoryPerson(ctx, UID)
	if err != nil {
		return "", err
	}
	return person.DisplayName(), nil
}
