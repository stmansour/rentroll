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
	UserRefNo string // It is a User Reference Number that refers to Flow
	RAID      int64  // It is RAID of Rental Agreement
	Version   string // "raid" or "refno"
	Action    int64  // If '-1' then Do nothing
	Mode      string // It represents that which button submitted the form
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
	TerminationReason  int64             // 0 - NoTerminationReasonField, great than 0 - Reason for Termination
	TerminationDate    rlib.JSONDateTime // the date on which we terminate the agreement
	TerminationStarted rlib.JSONDateTime // the date on which we made the change to the RA to terminate it (need not be the same as TerminationDate)
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
	rlib.Console("Entered %s\n", funcname)

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

		// COMMIT TRANSACTION
		if tx != nil {
			err = tx.Commit()
		}
	}()

	// HTTP METHOD CHECK
	if r.Method != "POST" {
		err = fmt.Errorf("only POST method is allowed")
		// SvcErrorReturn(w, err, funcname)   // err is returned by defer function
		return
	}

	// SEE IF WE CAN UNMARSHAL THE DATA
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		// SvcErrorReturn(w, err, funcname)   // err is returned by defer function
		return
	}

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err = rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		// SvcErrorReturn(w, err, funcname)   // err is returned by defer function
		return
	}

	var flow rlib.Flow

	rlib.Console("%s: foo = %#v\n", funcname, foo)
	rlib.Console("%s: d = %#v\n", funcname, d)

	switch foo.Version {
	case "raid":
		rlib.Console("%s: handleRAIDVersion\n", funcname)
		if flow, err = handleRAIDVersion(ctx, d, foo, raFlowData); err != nil {
			rlib.Console("error returned. Message length = %d, message: <<%s>>\n\n", len(err.Error()), err.Error())
			// SvcErrorReturn(w, err, funcname)   // err is returned by defer function
			return
		}
		SvcWriteFlowResponse(ctx, d.BID, flow, w)
		return

	case "refno":
		rlib.Console("%s: handleRefNoVersion\n", funcname)
		var resp FlowResponse
		var raflowRespData RAFlowResponse

		raflowRespData, err = handleRefNoVersion(ctx, d, foo, raFlowData)
		if err != nil {
			// SvcErrorReturn(w, err, funcname)   // err is returned by defer function
			return
		}
		resp.Record = raflowRespData
		resp.Status = "success"
		SvcWriteResponse(d.BID, &resp, w)
	}

}

func handleRAIDVersion(ctx context.Context, d *ServiceData, foo RAActionDataRequest, raFlowData rlib.RAFlowJSONData) (rlib.Flow, error) {
	var flow rlib.Flow
	var State uint64
	var err error
	RAID := foo.RAID
	Action := foo.Action

	// GET RENTAL AGREEMENT
	var ra rlib.RentalAgreement
	if RAID > 0 {
		ra, err = rlib.GetRentalAgreement(ctx, RAID)
		if err != nil {
			return flow, err
		}
		if ra.RAID == 0 {
			err = fmt.Errorf("rental Agreement not found with given RAID: %d", RAID)
			return flow, err
		}
		if ra.FLAGS&0xF == rlib.RASTATETerminated {
			err = fmt.Errorf("Rental Agreement has been Terminated, its state can no longer be modified")
			return flow, err
		}
	}

	switch Action {
	case
		rlib.RAActionApplicationBeingCompleted,
		rlib.RAActionSetToFirstApproval,
		rlib.RAActionSetToSecondApproval,
		rlib.RAActionSetToMoveIn:
		// If flow is present then get that flow
		flow, err = rlib.GetFlowForRAID(ctx, "RA", RAID)
		if err != nil {
			return flow, err
		}

		if flow.FlowID > 0 {
			// flow is already available, therefor send blank flow
			// to indicate warning
			blankData := []byte("{}")
			flow = rlib.Flow{FlowID: -1, Data: blankData}
			return flow, nil
		}

		// IF NOT FOUND THEN TRY TO CREATE NEW ONE FROM RAID
		// GET RENTAL AGREEMENT
		// var ra rlib.RentalAgreement
		// ra, err = rlib.GetRentalAgreement(ctx, RAID)
		// if err != nil {
		// 	return flow, err
		// }
		// if ra.RAID == 0 {
		// 	err = fmt.Errorf("rental Agreement not found with given RAID: %d", RAID)
		// 	return flow, err
		// }
		// if ra.FLAGS&0xF == rlib.RASTATETerminated {
		// 	return flow, fmt.Errorf("Rental Agreement has been Terminated, its state can no longer be modified")
		// }

		// GET THE NEW FLOW ID CREATED USING PERMANENT DATA
		var flowID int64
		EditFlag := false // we're only creating this to update the state and meta info, so we don't need to filter fees
		flowID, err = GetRA2FlowCore(ctx, &ra, d, EditFlag)
		if err != nil {
			return flow, err
		}

		// GET GENERATED FLOW USING NEW ID
		flow, err = rlib.GetFlow(ctx, flowID)
		if err != nil {
			return flow, err
		}

		ApplicationReadyName, _ := rlib.GetDirectoryPerson(ctx, ra.ApplicationReadyUID)
		MoveInName, _ := rlib.GetDirectoryPerson(ctx, ra.MoveInUID)
		ActiveName, _ := rlib.GetDirectoryPerson(ctx, ra.ActiveUID)
		Approver1Name, _ := rlib.GetDirectoryPerson(ctx, ra.Approver1)
		Approver2Name, _ := rlib.GetDirectoryPerson(ctx, ra.Approver2)
		TerminatorName, _ := rlib.GetDirectoryPerson(ctx, ra.TerminatorUID)
		NoticeToMoveName, _ := rlib.GetDirectoryPerson(ctx, ra.NoticeToMoveUID)

		// get meta in modRAFlowMeta, we're going to modify it
		modRAFlowMeta := rlib.RAFlowMetaInfo{
			BID:                    ra.BID,
			RAID:                   ra.RAID,
			RAFLAGS:                ra.FLAGS,
			ApplicationReadyUID:    ra.ApplicationReadyUID,
			ApplicationReadyName:   ApplicationReadyName.DisplayName(),
			ApplicationReadyDate:   rlib.JSONDateTime(ra.ApplicationReadyDate),
			Approver1:              ra.Approver1,
			Approver1Name:          Approver1Name.DisplayName(),
			DecisionDate1:          rlib.JSONDateTime(ra.DecisionDate1),
			DeclineReason1:         ra.DeclineReason1,
			Approver2:              ra.Approver2,
			Approver2Name:          Approver2Name.DisplayName(),
			DecisionDate2:          rlib.JSONDateTime(ra.DecisionDate2),
			DeclineReason2:         ra.DeclineReason2,
			MoveInUID:              ra.MoveInUID,
			MoveInName:             MoveInName.DisplayName(),
			MoveInDate:             rlib.JSONDateTime(ra.MoveInDate),
			ActiveUID:              ra.ActiveUID,
			ActiveName:             ActiveName.DisplayName(),
			ActiveDate:             rlib.JSONDateTime(ra.ActiveDate),
			TerminatorUID:          ra.TerminatorUID,
			TerminatorName:         TerminatorName.DisplayName(),
			TerminationDate:        rlib.JSONDateTime(ra.TerminationDate),
			TerminationStarted:     rlib.JSONDateTime(ra.TerminationStarted),
			LeaseTerminationReason: ra.LeaseTerminationReason,
			DocumentDate:           rlib.JSONDateTime(ra.DocumentDate),
			NoticeToMoveUID:        ra.NoticeToMoveUID,
			NoticeToMoveName:       NoticeToMoveName.DisplayName(),
			NoticeToMoveDate:       rlib.JSONDateTime(ra.NoticeToMoveDate),
			NoticeToMoveReported:   rlib.JSONDateTime(ra.NoticeToMoveReported),
		}

		State = ra.FLAGS & uint64(0xF)

		ActionResetMetaData(Action, State, &modRAFlowMeta) // RESET META INFO IF NEEDED

		// MODIFY META DATA
		err = SetActionMetaData(ctx, d, Action, &modRAFlowMeta)
		if err != nil {
			return flow, err
		}

		// UPDATE FLOW
		var modMetaData []byte
		modMetaData, err = json.Marshal(&modRAFlowMeta)
		if err != nil {
			return flow, err
		}

		err = rlib.UpdateFlowPartData(ctx, "meta", modMetaData, &flow)
		if err != nil {
			return flow, err
		}

		// get the updated flow
		flow, err = rlib.GetFlow(ctx, flow.FlowID)
		if err != nil {
			return flow, err
		}

	case
		rlib.RAActionCompleteMoveIn,
		rlib.RAActionReceivedNoticeToMove,
		rlib.RAActionTerminate:

		ApplicationReadyName, _ := rlib.GetDirectoryPerson(ctx, ra.ApplicationReadyUID)
		MoveInName, _ := rlib.GetDirectoryPerson(ctx, ra.MoveInUID)
		ActiveName, _ := rlib.GetDirectoryPerson(ctx, ra.ActiveUID)
		Approver1Name, _ := rlib.GetDirectoryPerson(ctx, ra.Approver1)
		Approver2Name, _ := rlib.GetDirectoryPerson(ctx, ra.Approver2)
		TerminatorName, _ := rlib.GetDirectoryPerson(ctx, ra.TerminatorUID)
		NoticeToMoveName, _ := rlib.GetDirectoryPerson(ctx, ra.NoticeToMoveUID)

		// get meta in modRAFlowMeta, we're going to modify it
		modRAFlowMeta := rlib.RAFlowMetaInfo{
			BID:                    ra.BID,
			RAID:                   ra.RAID,
			RAFLAGS:                ra.FLAGS,
			ApplicationReadyUID:    ra.ApplicationReadyUID,
			ApplicationReadyName:   ApplicationReadyName.DisplayName(),
			ApplicationReadyDate:   rlib.JSONDateTime(ra.ApplicationReadyDate),
			Approver1:              ra.Approver1,
			Approver1Name:          Approver1Name.DisplayName(),
			DecisionDate1:          rlib.JSONDateTime(ra.DecisionDate1),
			DeclineReason1:         ra.DeclineReason1,
			Approver2:              ra.Approver2,
			Approver2Name:          Approver2Name.DisplayName(),
			DecisionDate2:          rlib.JSONDateTime(ra.DecisionDate2),
			DeclineReason2:         ra.DeclineReason2,
			MoveInUID:              ra.MoveInUID,
			MoveInName:             MoveInName.DisplayName(),
			MoveInDate:             rlib.JSONDateTime(ra.MoveInDate),
			ActiveUID:              ra.ActiveUID,
			ActiveName:             ActiveName.DisplayName(),
			ActiveDate:             rlib.JSONDateTime(ra.ActiveDate),
			TerminatorUID:          ra.TerminatorUID,
			TerminatorName:         TerminatorName.DisplayName(),
			TerminationDate:        rlib.JSONDateTime(ra.TerminationDate),
			TerminationStarted:     rlib.JSONDateTime(ra.TerminationStarted),
			LeaseTerminationReason: ra.LeaseTerminationReason,
			DocumentDate:           rlib.JSONDateTime(ra.DocumentDate),
			NoticeToMoveUID:        ra.NoticeToMoveUID,
			NoticeToMoveName:       NoticeToMoveName.DisplayName(),
			NoticeToMoveDate:       rlib.JSONDateTime(ra.NoticeToMoveDate),
			NoticeToMoveReported:   rlib.JSONDateTime(ra.NoticeToMoveReported),
		}

		// take current state of Rental Agreement from FLAG
		State = ra.FLAGS & uint64(0xF)

		// RESET META INFO IF NEEDED
		ActionResetMetaData(Action, State, &modRAFlowMeta)

		// MODIFY META DATA
		err = SetActionMetaData(ctx, d, Action, &modRAFlowMeta)
		if err != nil {
			return flow, err
		}

		ra.RAID = modRAFlowMeta.RAID
		ra.FLAGS = modRAFlowMeta.RAFLAGS
		ra.ApplicationReadyUID = modRAFlowMeta.ApplicationReadyUID
		ra.ApplicationReadyDate = time.Time(modRAFlowMeta.ApplicationReadyDate)
		ra.Approver1 = modRAFlowMeta.Approver1
		ra.DecisionDate1 = time.Time(modRAFlowMeta.DecisionDate1)
		ra.DeclineReason1 = modRAFlowMeta.DeclineReason1
		ra.Approver2 = modRAFlowMeta.Approver2
		ra.DecisionDate2 = time.Time(modRAFlowMeta.DecisionDate2)
		ra.DeclineReason2 = modRAFlowMeta.DeclineReason2
		ra.MoveInUID = modRAFlowMeta.MoveInUID
		ra.MoveInDate = time.Time(modRAFlowMeta.MoveInDate)
		ra.ActiveUID = modRAFlowMeta.ActiveUID
		ra.ActiveDate = time.Time(modRAFlowMeta.ActiveDate)
		ra.TerminatorUID = modRAFlowMeta.TerminatorUID
		ra.TerminationDate = time.Time(modRAFlowMeta.TerminationDate)
		ra.TerminationStarted = time.Time(modRAFlowMeta.TerminationStarted)
		ra.LeaseTerminationReason = modRAFlowMeta.LeaseTerminationReason
		ra.DocumentDate = time.Time(modRAFlowMeta.DocumentDate)
		ra.NoticeToMoveUID = modRAFlowMeta.NoticeToMoveUID
		ra.NoticeToMoveDate = time.Time(modRAFlowMeta.NoticeToMoveDate)
		ra.NoticeToMoveReported = time.Time(modRAFlowMeta.NoticeToMoveReported)

		rlib.Console("ACTION: %d.  ra.TerminationDate = %s, ra.TerminationStarted = %s\n", Action, rlib.ConDt(&ra.TerminationDate), rlib.ConDt(&ra.TerminationStarted))

		//---------------------------------------------------------------------
		// If this is a termination, clean up the loose ends
		//---------------------------------------------------------------------
		if Action == rlib.RAActionTerminate {
			if err = TerminateRentalAgreement(ctx, &ra); err != nil {
				return flow, err
			}
		} else {
			// UPDATE RA in REAL TABLE
			err = rlib.UpdateRentalAgreement(ctx, &ra)
			if err != nil {
				return flow, err
			}

		}

		// EditFlag should be set to true only when we're creating a Flow that
		// becomes a RefNo (an amended RentalAgreement)
		EditFlag := false // this is the behavior as it was prior to the EditFlag being added.

		// Create flow to viewing in UI
		var raf rlib.RAFlowJSONData
		raf, err = rlib.ConvertRA2Flow(ctx, &ra, EditFlag)
		if err != nil {
			return flow, err
		}

		var raflowJSONData []byte
		raflowJSONData, err = json.Marshal(&raf)
		if err != nil {
			return flow, err
		}

		flow = rlib.Flow{
			BID:       ra.BID,
			FlowID:    0, // we're not creating any flow, just to see RA content
			UserRefNo: "",
			FlowType:  rlib.RAFlow,
			ID:        ra.RAID,
			Data:      raflowJSONData,
			CreateBy:  0,
			LastModBy: 0,
		}
	}

	// RETURN FLOW
	return flow, nil
}

// TerminateRentalAgreement is called when an active RA is set to the terminated
// state. It will stop assessments, update User / payor references to end on the
// termination date, and it will update the LeaseStatus for the time after the
// Termination Date.  This is a simpler scenario than what happens in Flow2RA().
// It will terminate on the date set in ra.TerminationDate
//
// INPUTS
//    ctx  - db context
//    ra   - rental agreement struct for terminated RA
//
// RETURNS
//    any errors encountered
//------------------------------------------------------------------------------
func TerminateRentalAgreement(ctx context.Context, ra *rlib.RentalAgreement) error {
	rlib.Console("Entered TerminateRentalAgreement\n")

	var err error
	origStartDate := rlib.Earliest(&ra.RentStart, &ra.PossessionStart)
	origStopDate := rlib.Latest(&ra.RentStop, &ra.PossessionStop)
	var tx *sql.Tx
	var ok bool
	var sess *rlib.Session

	//-----------------------------------------------
	// We terminate everything on the TerminationDate
	//-----------------------------------------------
	ra.AgreementStop = ra.TerminationDate
	ra.RentStop = ra.TerminationDate
	ra.PossessionStop = ra.TerminationDate

	if err = rlib.UpdateRentalAgreement(ctx, ra); err != nil {
		return err
	}
	if sess, ok = rlib.SessionFromContext(ctx); !ok {
		return rlib.ErrSessionRequired
	}
	if tx, ok = rlib.DBTxFromContext(ctx); !ok {
		return fmt.Errorf("Could not get transaction info from context")
	}

	//--------------------------------------------------------------------------
	//  Stop all recurring assessments on the termination date
	//--------------------------------------------------------------------------
	qry := fmt.Sprintf("UPDATE Assessments SET Stop = %q,LastModBy=%d WHERE RAID=%d AND PASMID=0 AND RentCycle>0 AND Stop>%q;",
		ra.RentStop.Format(rlib.RRDATETIMESQL), sess.UID, ra.RAID, ra.RentStop.Format(rlib.RRDATETIMESQL))
	if _, err = tx.Exec(qry); err != nil {
		return err
	}

	//--------------------------------------------------------------------------
	//  Update Payors with changed end date
	//--------------------------------------------------------------------------
	qry = fmt.Sprintf("UPDATE RentalAgreementPayors SET DtStop=%q,LastModBy=%d WHERE RAID=%d AND DtStop>%q;",
		ra.RentStop.Format(rlib.RRDATETIMESQL), sess.UID, ra.RAID, ra.RentStop.Format(rlib.RRDATETIMESQL))
	if _, err = tx.Exec(qry); err != nil {
		return err
	}

	//--------------------------------------------------------------------------
	//  Update Users with changed end date
	//--------------------------------------------------------------------------
	var s string
	if s, err = GetQueryRentableList(ctx, ra.RAID); err != nil {
		return err
	}
	qry = fmt.Sprintf("UPDATE RentableUsers SET DtStop=%q,LastModBy=%d WHERE DtStart>=%q AND DtStop<=%q AND RID IN %s",
		ra.RentStop,
		sess.UID,
		origStartDate.Format(rlib.RRDATETIMESQL),
		origStopDate.Format(rlib.RRDATETIMESQL),
		s)
	if _, err = tx.Exec(qry); err != nil {
		return err
	}

	//--------------------------------------------------------------------------
	//  Look at the Lease Status following the stop date. Adjust if necessary.
	//--------------------------------------------------------------------------
	SetLeaseStatusPostStop(ctx, ra.BID, ra.RAID, &ra.TerminationDate, &origStopDate)

	return nil
}

func handleRefNoVersion(ctx context.Context, d *ServiceData, foo RAActionDataRequest, raFlowData rlib.RAFlowJSONData) (RAFlowResponse, error) {

	UserRefNo := foo.UserRefNo
	Action := foo.Action
	Mode := foo.Mode

	var flow rlib.Flow
	var err error

	var raflowRespData RAFlowResponse

	migrateData := false

	//--------------------------------------------------------------------------
	// In noauth mode, it still have tester session
	//--------------------------------------------------------------------------
	UID := d.sess.UID

	// GET FLOW BY REFNO
	flow, err = rlib.GetFlowByUserRefNo(ctx, d.BID, UserRefNo)
	if err != nil {
		return raflowRespData, err
	}
	if flow.FlowID <= 0 {
		err = fmt.Errorf("rental Agreement Flow not found with given Customer Reference No.: %s", UserRefNo)
		return raflowRespData, err
	}

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		return raflowRespData, err
	}

	raflowRespData.Flow = flow

	ValidateRAFlowAndAssignValidatedRAFlow(ctx, &raFlowData, flow, &raflowRespData)

	// CHECK DATA FULFILLED
	bizlogic.DataFulfilledRAFlow(ctx, &raFlowData, &raflowRespData.DataFulfilled)

	if raflowRespData.ValidationCheck.Total > 0 ||
		!(raflowRespData.DataFulfilled.Dates && raflowRespData.DataFulfilled.People &&
			raflowRespData.DataFulfilled.Pets && raflowRespData.DataFulfilled.Vehicles &&
			raflowRespData.DataFulfilled.Rentables && raflowRespData.DataFulfilled.ParentChild &&
			raflowRespData.DataFulfilled.Tie) {
		return raflowRespData, nil
	}

	// get meta in modRAFlowMeta, we're going to modify it
	modRAFlowMeta := raFlowData.Meta

	// GET THE CURRENT STATE FROM THE LAST 4 BITS
	State := raFlowData.Meta.RAFLAGS & uint64(0xF)

	switch Mode {
	case "Action":
		switch Action {
		case
			rlib.RAActionApplicationBeingCompleted,
			rlib.RAActionSetToFirstApproval,
			rlib.RAActionSetToSecondApproval,
			rlib.RAActionSetToMoveIn,
			rlib.RAActionCompleteMoveIn:

			// RESET META INFO IF NEEDED
			ActionResetMetaData(Action, State, &modRAFlowMeta)

			// MODIFY META DATA
			err = SetActionMetaData(ctx, d, Action, &modRAFlowMeta)
			if err != nil {
				return raflowRespData, err
			}

			if Action == rlib.RAActionCompleteMoveIn {
				migrateData = true
			}

		default:
			err = fmt.Errorf("invalid Action Taken")
			return raflowRespData, err
		}
	case "State":
		// set location for time as UTC
		var location *time.Location
		location, err = time.LoadLocation("UTC")
		if err != nil {
			return raflowRespData, err
		}

		// get current time in UTC
		var today time.Time
		// today = time.Now().In(location)
		today = rlib.Now().In(location)

		// take latest RAFLAGS value at this point(in case flag bits are reset)
		// clearedState := modRAFlowMeta.RAFLAGS & ^uint64(0xF)  // sm: this was an ineffectual assignment
		var clearedState uint64

		switch State {
		case rlib.RASTATEPendingApproval1:
			var data RAApprover1Data
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				return raflowRespData, err
			}

			var fullName string
			fullName, err = getUserFullName(ctx, UID)
			if err != nil {
				return raflowRespData, err
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

				modRAFlowMeta.TerminatorUID = UID
				modRAFlowMeta.TerminatorName = fullName
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(today)
				modRAFlowMeta.TerminationStarted = rlib.JSONDateTime(today)

				// IF BIZCACHE NOT INITIALIZED THEN
				if rlib.RRdb.BizTypes[d.BID] == nil {
					var xbiz rlib.XBusiness
					err = rlib.InitBizInternals(d.BID, &xbiz)
					if err != nil {
						return raflowRespData, err
					}
				}
				// APPLICATION DECLINED SLSID IN LEASE TERMINATION REASON
				modRAFlowMeta.LeaseTerminationReason = rlib.RRdb.BizTypes[d.BID].Msgs.S[rlib.MSGAPPDECLINED].SLSID

				// FLAGS
				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 6)
			} else {
				err = fmt.Errorf("approver1 data is not valid")
				return raflowRespData, err
			}

			modRAFlowMeta.Approver1 = UID
			modRAFlowMeta.Approver1Name = fullName
			modRAFlowMeta.DecisionDate1 = rlib.JSONDateTime(today)

		case rlib.RASTATEPendingApproval2:
			var data RAApprover2Data
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				return raflowRespData, err
			}

			var fullName string
			fullName, err = getUserFullName(ctx, UID)
			if err != nil {
				return raflowRespData, err
			}

			if data.Decision2 == 1 { // Approved
				modRAFlowMeta.MoveInUID = UID
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

				modRAFlowMeta.TerminatorUID = UID
				modRAFlowMeta.TerminatorName = fullName
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(today)
				modRAFlowMeta.TerminationStarted = rlib.JSONDateTime(today)

				// IF BIZCACHE NOT INITIALIZED THEN
				if rlib.RRdb.BizTypes[d.BID] == nil {
					var xbiz rlib.XBusiness
					err = rlib.InitBizInternals(d.BID, &xbiz)
					if err != nil {
						return raflowRespData, err
					}
				}
				// APPLICATION DECLINED SLSID IN LEASE TERMINATION REASON
				modRAFlowMeta.LeaseTerminationReason = rlib.RRdb.BizTypes[d.BID].Msgs.S[rlib.MSGAPPDECLINED].SLSID

				// FLAGS
				clearedState = modRAFlowMeta.RAFLAGS & ^uint64(0xf)
				modRAFlowMeta.RAFLAGS = (clearedState | 6)
			} else {
				err = fmt.Errorf("approver2 data is not valid")
				return raflowRespData, err
			}

			modRAFlowMeta.Approver2 = UID
			modRAFlowMeta.Approver2Name = fullName
			modRAFlowMeta.DecisionDate2 = rlib.JSONDateTime(today)

		case rlib.RASTATEMoveIn:
			var data RAMoveInData
			if err = json.Unmarshal([]byte(d.data), &data); err != nil {
				return raflowRespData, err
			}

			var fullName string
			fullName, err = getUserFullName(ctx, UID)
			if err != nil {
				return raflowRespData, err
			}

			modRAFlowMeta.MoveInUID = UID
			modRAFlowMeta.MoveInName = fullName
			modRAFlowMeta.MoveInDate = rlib.JSONDateTime(today)

			modRAFlowMeta.DocumentDate = rlib.JSONDateTime(data.DocumentDate)
		}
	}

	// UPDATE META AND FLOW
	var modMetaData []byte
	modMetaData, err = json.Marshal(&modRAFlowMeta)
	if err != nil {
		return raflowRespData, err
	}

	err = rlib.UpdateFlowPartData(ctx, "meta", modMetaData, &flow)
	if err != nil {
		return raflowRespData, err
	}

	// GET UPDATED FLOW
	flow, err = rlib.GetFlowByUserRefNo(ctx, flow.BID, flow.UserRefNo)
	if err != nil {
		return raflowRespData, err
	}
	raflowRespData.Flow = flow

	if migrateData {
		// migrate data to real table via hook
		var newRAID int64
		newRAID, err = Flow2RA(ctx, flow.FlowID)
		if err != nil {
			return raflowRespData, err
		}

		// GET RENTAL AGREEMENT
		var ra rlib.RentalAgreement
		ra, err = rlib.GetRentalAgreement(ctx, newRAID)
		if err != nil {
			return raflowRespData, err
		}
		if ra.RAID == 0 {
			err = fmt.Errorf("Rental Agreement not found with given RAID: %d", newRAID)
			return raflowRespData, err
		}

		// EditFlag should be set to true only when we're creating a Flow that
		// becomes a RefNo (an amended RentalAgreement)
		EditFlag := false // assume we're asking for the view version

		// convert permanent ra to flow data and get it
		var raf rlib.RAFlowJSONData
		raf, err = rlib.ConvertRA2Flow(ctx, &ra, EditFlag)
		if err != nil {
			return raflowRespData, err
		}

		//-------------------------------------------------------------------------
		// Save the flow to the db
		//-------------------------------------------------------------------------
		var raflowJSONData []byte
		raflowJSONData, err = json.Marshal(&raf)
		if err != nil {
			return raflowRespData, err
		}

		// After Migration, flow will be deleted
		// Hence we create a Flow structure and return it just for display purpose
		flow = rlib.Flow{
			BID:       ra.BID,
			FlowID:    0, // we're not creating any flow, just to see RA content
			UserRefNo: "",
			ID:        newRAID,
			FlowType:  rlib.RAFlow,
			Data:      raflowJSONData,
			CreateBy:  0,
			LastModBy: 0,
		}
		raflowRespData.Flow = flow
	}
	return raflowRespData, nil
}

func getUserFullName(ctx context.Context, UID int64) (string, error) {
	person, err := rlib.GetDirectoryPerson(ctx, UID)
	if err != nil {
		return "", err
	}
	return person.DisplayName(), nil
}

// ActionResetMetaData resets the info in meta based on the action upto the current state
func ActionResetMetaData(Action int64, State uint64, modRAFlowMeta *rlib.RAFlowMetaInfo) {
	if Action <= int64(State) {
		for i := Action; i <= int64(State); i++ {
			switch i {
			case rlib.RAActionApplicationBeingCompleted:
				modRAFlowMeta.ApplicationReadyUID = 0
				modRAFlowMeta.ApplicationReadyName = ""
				modRAFlowMeta.ApplicationReadyDate = rlib.JSONDateTime(time.Time{})

			case rlib.RAActionSetToFirstApproval:
				modRAFlowMeta.Approver1 = 0
				modRAFlowMeta.Approver1Name = ""
				modRAFlowMeta.DeclineReason1 = 0
				modRAFlowMeta.DecisionDate1 = rlib.JSONDateTime(time.Time{})

				// reset 4th bit of RAFLAG to 0
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS & ^uint64(1<<4)

			case rlib.RAActionSetToSecondApproval:
				modRAFlowMeta.Approver2 = 0
				modRAFlowMeta.Approver2Name = ""
				modRAFlowMeta.DeclineReason2 = 0
				modRAFlowMeta.DecisionDate2 = rlib.JSONDateTime(time.Time{})

				// reset 5th bit of RAFLAG to 0
				modRAFlowMeta.RAFLAGS = modRAFlowMeta.RAFLAGS & ^uint64(1<<5)

			case rlib.RAActionSetToMoveIn:
				modRAFlowMeta.MoveInUID = 0
				modRAFlowMeta.MoveInName = ""
				modRAFlowMeta.MoveInDate = rlib.JSONDateTime(time.Time{})
				modRAFlowMeta.DocumentDate = rlib.JSONDateTime(time.Time{})

			case rlib.RAActionCompleteMoveIn:
				modRAFlowMeta.ActiveUID = 0
				modRAFlowMeta.ActiveName = ""
				modRAFlowMeta.ActiveDate = rlib.JSONDateTime(time.Time{})

			case rlib.RAActionReceivedNoticeToMove:
				modRAFlowMeta.NoticeToMoveUID = 0
				modRAFlowMeta.NoticeToMoveName = ""
				modRAFlowMeta.NoticeToMoveDate = rlib.JSONDateTime(time.Time{})
				modRAFlowMeta.NoticeToMoveReported = rlib.JSONDateTime(time.Time{})

			case rlib.RAActionTerminate:
				modRAFlowMeta.TerminatorUID = 0
				modRAFlowMeta.TerminatorName = ""
				modRAFlowMeta.LeaseTerminationReason = 0
				modRAFlowMeta.TerminationDate = rlib.JSONDateTime(time.Time{})
				modRAFlowMeta.TerminationStarted = rlib.JSONDateTime(time.Time{})
			}
		}
	}
}

// SetActionMetaData sets the info in meta based on the given action
func SetActionMetaData(ctx context.Context, d *ServiceData, Action int64, modRAFlowMeta *rlib.RAFlowMetaInfo) error {
	var err error

	// set location for time as UTC
	var location *time.Location
	location, err = time.LoadLocation("UTC")
	if err != nil {
		return err
	}

	// get current time in UTC
	var today time.Time
	// today = time.Now().In(location)
	today = rlib.Now().In(location)

	//--------------------------------------------------------------------------
	// In noauth mode, it still have tester session
	//--------------------------------------------------------------------------
	UID := d.sess.UID

	// take latest RAFLAGS value at this point(in case flag bits are reset)
	clearedState := modRAFlowMeta.RAFLAGS & ^uint64(0xF)
	switch Action {
	case rlib.RAActionApplicationBeingCompleted:
		modRAFlowMeta.RAFLAGS = (clearedState | 0)

	case rlib.RAActionSetToFirstApproval:
		var fullName string
		fullName, err = getUserFullName(ctx, UID)
		if err != nil {
			return err
		}

		modRAFlowMeta.ApplicationReadyUID = UID
		modRAFlowMeta.ApplicationReadyName = fullName
		modRAFlowMeta.ApplicationReadyDate = rlib.JSONDateTime(today)

		modRAFlowMeta.RAFLAGS = (clearedState | 1)

	case rlib.RAActionSetToSecondApproval:
		modRAFlowMeta.RAFLAGS = (clearedState | 2)

	case rlib.RAActionSetToMoveIn:
		var fullName string
		fullName, err = getUserFullName(ctx, UID)
		if err != nil {
			return err
		}

		modRAFlowMeta.MoveInUID = UID
		modRAFlowMeta.MoveInName = fullName
		modRAFlowMeta.MoveInDate = rlib.JSONDateTime(today)

		modRAFlowMeta.RAFLAGS = (clearedState | 3)

	case rlib.RAActionCompleteMoveIn:
		var fullName string
		fullName, err = getUserFullName(ctx, UID)
		if err != nil {
			return err
		}

		modRAFlowMeta.ActiveUID = UID
		modRAFlowMeta.ActiveName = fullName
		modRAFlowMeta.ActiveDate = rlib.JSONDateTime(today)
		modRAFlowMeta.RAFLAGS = (clearedState | 4)

	case rlib.RAActionReceivedNoticeToMove:
		var data RANoticeToMoveData
		if err = json.Unmarshal([]byte(d.data), &data); err != nil {
			return err
		}

		var fullName string
		fullName, err = getUserFullName(ctx, UID)
		if err != nil {
			return err
		}

		modRAFlowMeta.NoticeToMoveUID = UID
		modRAFlowMeta.NoticeToMoveName = fullName
		modRAFlowMeta.NoticeToMoveDate = rlib.JSONDateTime(data.NoticeToMoveDate)
		modRAFlowMeta.NoticeToMoveReported = rlib.JSONDateTime(today)

		modRAFlowMeta.RAFLAGS = (clearedState | 5)

	case rlib.RAActionTerminate:
		rlib.Console("Entered case: RAActionTerminate")
		var data RATerminationData
		if err = json.Unmarshal([]byte(d.data), &data); err != nil {
			return err
		}

		if data.TerminationReason > 0 {
			var fullName string
			fullName, err = getUserFullName(ctx, UID)
			if err != nil {
				return err
			}

			modRAFlowMeta.TerminatorUID = UID
			modRAFlowMeta.TerminatorName = fullName
			modRAFlowMeta.TerminationDate = data.TerminationDate
			modRAFlowMeta.TerminationStarted = data.TerminationStarted
			modRAFlowMeta.LeaseTerminationReason = data.TerminationReason

			modRAFlowMeta.RAFLAGS = (clearedState | 6)

			rlib.Console("Termination Date: %s\n", rlib.ConsoleJSONDate(&data.TerminationDate))
			rlib.Console("Termination Started: %s\n", rlib.ConsoleJSONDate(&data.TerminationStarted))

		} else { // NO TERMINATION REASON
			err = fmt.Errorf("termination reason is not provided")
			return err
		}
	}
	return nil
}
