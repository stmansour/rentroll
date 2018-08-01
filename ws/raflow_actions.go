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
		raFlowData rlib.RAFlowJSONData
		foo        RAActionDataRequest
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

	// FLAG: SHOULD WE MIGRATE DATA TO RA
	migrateData := false

	// set location for time as UTC
	var location *time.Location
	location, err = time.LoadLocation("UTC")
	if err != nil {
		return
	}

	// get current time in UTC
	var today time.Time
	today = time.Now().In(location)

	// HTTP METHOD CHECK
	if r.Method != "POST" {
		err = fmt.Errorf("Only POST method is allowed")
		return
	}

	// SEE IF WE CAN UNMARSHAL THE DATA
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

	// MODE OF OPERATION, IS IT ACTION BASED OR STATE BASED UPDATE
	MODE := foo.Mode

	// GET THE CURRENT STATE FROM THE LAST 4 BITS
	state := raFlowData.Meta.RAFLAGS & uint64(0xF)

	// WE DON'T WANT TO LOSE ALL BITS EXCEPT LAST 4 BITS (WHICH REPRESENTS THE RA STATE)
	clearedState := raFlowData.Meta.RAFLAGS & ^uint64(0xF)

	// get meta in modRAFlowMeta, we're going to modify it
	modRAFlowMeta := raFlowData.Meta

	switch MODE {
	case "Action":
		if foo.Action < int64(state) {
			for i := foo.Action; i <= int64(state); i++ {
				switch i {
				case 0: // Application Being Completed
					modRAFlowMeta.ApplicationReadyUID = 0
					modRAFlowMeta.ApplicationReadyName = ""
					modRAFlowMeta.ApplicationReadyDate = rlib.JSONDateTime(time.Time{})

				case 1: // Pending First Approval
					modRAFlowMeta.Approver1 = 0
					modRAFlowMeta.Approver1Name = ""
					modRAFlowMeta.DeclineReason1 = 0
					modRAFlowMeta.DecisionDate1 = rlib.JSONDateTime(time.Time{})

					// reset 4th bit of RAFLAG to 0
					modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS & ^uint64(1<<4)

				case 2: // Pending Second Approval
					modRAFlowMeta.Approver2 = 0
					modRAFlowMeta.Approver2Name = ""
					modRAFlowMeta.DeclineReason2 = 0
					modRAFlowMeta.DecisionDate2 = rlib.JSONDateTime(time.Time{})

					// reset 5th bit of RAFLAG to 0
					modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS & ^uint64(1<<5)

				case 3: // Move-In / Execute Modification
					modRAFlowMeta.MoveInUID = 0
					modRAFlowMeta.MoveInName = ""
					modRAFlowMeta.DocumentDate = rlib.JSONDateTime(time.Time{})

				case 4: // Active

				case 5: //Notice To Move
					modRAFlowMeta.NoticeToMoveUID = 0
					modRAFlowMeta.NoticeToMoveName = ""
					modRAFlowMeta.NoticeToMoveDate = rlib.JSONDateTime(time.Time{})
					modRAFlowMeta.NoticeToMoveReported = rlib.JSONDateTime(time.Time{})

				case 6: // Terminated
					modRAFlowMeta.TerminatorUID = 0
					modRAFlowMeta.TerminatorName = ""
					modRAFlowMeta.LeaseTerminationReason = 0
					modRAFlowMeta.TerminationDate = rlib.JSONDateTime(time.Time{})
				}
			}
		}

		// take latest RAFLAGS value at this point(in case flag bits are reset)
		clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xF)
		switch foo.Action {
		case 0: // Application Being Completed
			modRAFlowMeta.RAFLAGS = (clearedState | 0)

		case 1: // Set To First Approval
			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			modRAFlowMeta.ApplicationReadyUID = d.sess.UID
			modRAFlowMeta.ApplicationReadyName = fullName
			modRAFlowMeta.ApplicationReadyDate = rlib.JSONDateTime(today)

			modRAFlowMeta.RAFLAGS = (clearedState | 1)

		case 2: // Set To Second Approval
			modRAFlowMeta.RAFLAGS = (clearedState | 2)

		case 3: // Set To Move-In
			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			modRAFlowMeta.MoveInUID = d.sess.UID
			modRAFlowMeta.MoveInName = fullName
			modRAFlowMeta.MoveInDate = rlib.JSONDateTime(today)

			modRAFlowMeta.RAFLAGS = (clearedState | 3)

		case 4: // Complete Move-In
			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}

			modRAFlowMeta.ActiveUID = d.sess.UID
			modRAFlowMeta.ActiveName = fullName
			modRAFlowMeta.ActiveDate = rlib.JSONDateTime(today)
			modRAFlowMeta.RAFLAGS = (clearedState | 4)

			// migrate data to real table
			migrateData = true

		case 5: // Notice-To-Move
			var data RANoticeToMoveData
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				return
			}

			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				return
			}

			modRAFlowMeta.NoticeToMoveUID = d.sess.UID
			modRAFlowMeta.NoticeToMoveName = fullName
			modRAFlowMeta.NoticeToMoveDate = rlib.JSONDateTime(data.NoticeToMoveDate)
			modRAFlowMeta.NoticeToMoveReported = rlib.JSONDateTime(today)

			modRAFlowMeta.RAFLAGS = (clearedState | 5)

			// if existing RA then migrate data to real table
			if modRAFlowMeta.RAID > 0 {
				migrateData = true
			}

		case 6: // Terminate
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

				modRAFlowMeta.TerminatorUID = d.sess.UID
				modRAFlowMeta.TerminatorName = fullName
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(today)
				modRAFlowMeta.LeaseTerminationReason = data.TerminationReason

				modRAFlowMeta.RAFLAGS = (clearedState | 6)

				// if existing RA then migrate data to real table
				if modRAFlowMeta.RAID > 0 {
					migrateData = true
				}

			} else { // NO TERMINATION REASON
				err = fmt.Errorf("termination reason is not provided")
				return
			}
		}

	case "State":
		switch state {
		case 1: // Pending First Approval
			var data RAApprover1Data
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				return
			}

			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				return
			}

			if data.Decision1 == 1 { // Approved
				// set 4th bit of Flag as 1
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS | uint64(1<<4)

				// FLAGS
				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 2)
			} else if data.Decision1 == 2 && data.DeclineReason1 > 0 { // Declined
				// set 4th bit of Flag as 0
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS & ^uint64(1<<4)
				modRAFlowMeta.DeclineReason1 = data.DeclineReason1

				modRAFlowMeta.TerminatorUID = d.sess.UID
				modRAFlowMeta.TerminatorName = fullName
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(today)

				// IF BIZCACHE NOT INITIALIZED THEN
				if rlib.RRdb.BizTypes[d.BID] == nil {
					var xbiz rlib.XBusiness
					err = rlib.InitBizInternals(d.BID, &xbiz)
					if err != nil {
						return
					}
				}
				// APPLICATION DECLINED SLSID IN LEASE TERMINATION REASON
				modRAFlowMeta.LeaseTerminationReason = rlib.RRdb.BizTypes[d.BID].Msgs.S[rlib.MSGAPPDECLINED].SLSID

				// FLAGS
				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 6)
			} else {
				err = fmt.Errorf("approver1 data is not valid")
				return
			}

			modRAFlowMeta.Approver1 = d.sess.UID
			modRAFlowMeta.Approver1Name = fullName
			modRAFlowMeta.DecisionDate1 = rlib.JSONDateTime(today)

		case 2: // Pending Second Approval
			var data RAApprover2Data
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				return
			}

			var fullName string
			fullName, err = getUserFullName(ctx, d.sess.UID)
			if err != nil {
				return
			}

			if data.Decision2 == 1 { // Approved
				modRAFlowMeta.MoveInUID = d.sess.UID
				modRAFlowMeta.MoveInName = fullName
				modRAFlowMeta.MoveInDate = rlib.JSONDateTime(today)

				// set 5th bit of Flag as 1
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS | uint64(1<<5)

				// FLAGS
				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 3)
			} else if data.Decision2 == 2 && data.DeclineReason2 > 0 { // Declined
				// set 5th bit of Flag as 0
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS & ^uint64(1<<5)
				modRAFlowMeta.DeclineReason2 = data.DeclineReason2

				modRAFlowMeta.TerminatorUID = d.sess.UID
				modRAFlowMeta.TerminatorName = fullName
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(today)

				// IF BIZCACHE NOT INITIALIZED THEN
				if rlib.RRdb.BizTypes[d.BID] == nil {
					var xbiz rlib.XBusiness
					err = rlib.InitBizInternals(d.BID, &xbiz)
					if err != nil {
						return
					}
				}
				// APPLICATION DECLINED SLSID IN LEASE TERMINATION REASON
				modRAFlowMeta.LeaseTerminationReason = rlib.RRdb.BizTypes[d.BID].Msgs.S[rlib.MSGAPPDECLINED].SLSID

				// FLAGS
				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 6)
			} else {
				err = fmt.Errorf("approver2 data is not valid")
				return
			}

			modRAFlowMeta.Approver2 = d.sess.UID
			modRAFlowMeta.Approver2Name = fullName
			modRAFlowMeta.DecisionDate2 = rlib.JSONDateTime(today)

		case 3: // Move-In / Execute Modification
			var data RAMoveInData
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				return
			}
			modRAFlowMeta.DocumentDate = rlib.JSONDateTime(data.DocumentDate)
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

	if migrateData {
		// migrate data to real table via hook
		var newRAID int64
		newRAID, err = Flow2RA(ctx, flow.FlowID)
		if err != nil {
			return
		}

		if modRAFlowMeta.RAID == 0 {
			modRAFlowMeta.RAID = newRAID
			modMetaData, err = json.Marshal(&modRAFlowMeta)
			if err != nil {
				return
			}

			flow.ID = newRAID
			err = rlib.UpdateFlow(ctx, &flow)
			if err != nil {
				return
			}

			err = rlib.UpdateFlowData(ctx, "meta", modMetaData, &flow)
			if err != nil {
				return
			}
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

func getUserFullName(ctx context.Context, UID int64) (string, error) {
	person, err := rlib.GetDirectoryPerson(ctx, UID)
	if err != nil {
		return "", err
	}
	return person.DisplayName(), nil
}
