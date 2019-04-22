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

// RAFlowDetailRequest is a struct to hold info for Flow which is going to be validate
type RAFlowDetailRequest struct {
	FlowID    int64
	UserRefNo string
}

// SvcValidateRAFlow is used to check/validate RAFlow's struct
// wsdoc {
//  @Title Validate RAFlow's basic and biz logic check
//  @URL /v1/validate-raflow/:BUI
//  @Method  POST
//  @Synopsis Validate RAFlow
//  @Description Perform basic validation and businness logic check validation on RAFlow
//  @Input RAFlowDetailRequest
//  @Response bizlogic.ValidateRAFlowResponse
// wsdoc }
func SvcValidateRAFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcValidateRAFlow"
	var (
		err error
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		ValidateRAFlow(w, r, d)
		break
	default:
		err = fmt.Errorf("unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// ValidateRAFlow validate RAFlow's fields section wise
func ValidateRAFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "ValidateRAFlow"
	rlib.Console("Entered %s\n", funcname)

	var (
		err                   error
		foo                   RAFlowDetailRequest
		raFlowData            rlib.RAFlowJSONData
		raFlowFieldsErrors    bizlogic.RAFlowFieldsErrors
		raFlowNonFieldsErrors bizlogic.RAFlowNonFieldsErrors
		raflowRespData        RAFlowResponse
		resp                  FlowResponse
		g                     bizlogic.ValidateRAFlowResponse
		ctx                   context.Context
		tx                    *sql.Tx
	)

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

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("only POST method is allowed")
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

	// Get flow information from the table to validate fields value
	var flow rlib.Flow
	flow, err = rlib.GetFlow(ctx, foo.FlowID)
	if err != nil {
		return
	}

	// When flowId doesn't exists in database return and give error that flowId doesn't exists
	if flow.FlowID == 0 {
		err = fmt.Errorf("flowID %d - doesn't exists", foo.FlowID)
		return
	}

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		return
	}

	rlib.EDIHandleIncomingJSONDateRange(raFlowData.Meta.BID, &raFlowData.Dates.AgreementStart, &raFlowData.Dates.AgreementStop)
	rlib.EDIHandleIncomingJSONDateRange(raFlowData.Meta.BID, &raFlowData.Dates.PossessionStart, &raFlowData.Dates.PossessionStop)
	rlib.EDIHandleIncomingJSONDateRange(raFlowData.Meta.BID, &raFlowData.Dates.RentStart, &raFlowData.Dates.RentStop)

	rlib.Console("DtStart, DtStop = %s\n", rlib.ConsoleJSONDRange(&raFlowData.Dates.AgreementStart, &raFlowData.Dates.AgreementStop))

	// CHECK DATA FULFILLED
	bizlogic.DataFulfilledRAFlow(ctx, &raFlowData, &raflowRespData.DataFulfilled)

	// init raFlowFieldsErrors
	initRAFlowFieldsErrors(&raFlowFieldsErrors)

	initRAFlowNonFieldsErrors(&raFlowNonFieldsErrors)

	err = bizlogic.ValidateRAFlowParts(ctx, &raFlowFieldsErrors, &raFlowNonFieldsErrors, &raFlowData, flow.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
	}

	totalFieldsError := raFlowFieldsErrors.Dates.Total + raFlowFieldsErrors.People.Total + raFlowFieldsErrors.Pets.Total + raFlowFieldsErrors.Vehicle.Total + raFlowFieldsErrors.Rentables.Total + raFlowFieldsErrors.ParentChild.Total + raFlowFieldsErrors.Tie.TiePeople.Total
	totalNonFieldsError := len(raFlowNonFieldsErrors.Dates) + len(raFlowNonFieldsErrors.People) + len(raFlowNonFieldsErrors.Pets) + len(raFlowNonFieldsErrors.Rentables) + len(raFlowNonFieldsErrors.Vehicle) + len(raFlowNonFieldsErrors.ParentChild) + len(raFlowNonFieldsErrors.Tie)
	g.Total += totalFieldsError + totalNonFieldsError
	g.Errors = raFlowFieldsErrors
	g.NonFieldsErrors = raFlowNonFieldsErrors

	rlib.Console("Total Field errors: %d, non-field errors: %d\n", totalFieldsError, totalNonFieldsError)

	if g.Total == 0 {
		// SET STATE OF THIS FLOW TO PENDING FIRST APPROVAL
		// ONLY IF STATE IS == "APPLICATION BEING COMPLETED"
		if raFlowData.Meta.RAFLAGS&(0xF) == 0 {
			// AS STATE IS == "APP BEING COMPLETED"
			// DON'T NEED TO RESET META
			// DIRECTLY SET ACTION META WITH PENDING FIRST APPROVAL
			action := int64(rlib.RAActionSetToFirstApproval)

			err = SetActionMetaData(ctx, d, action, &raFlowData.Meta)
			if err != nil {
				return
			}

			// UPDATE META PART OF THE FLOW
			var modMetaData []byte
			modMetaData, err = json.Marshal(&raFlowData.Meta)
			if err != nil {
				return
			}
			err = rlib.UpdateFlowPartData(ctx, "meta", modMetaData, &flow)
			if err != nil {
				return
			}

			// GET THE UPDATED FLOW
			flow, err = rlib.GetFlow(ctx, flow.FlowID)
			if err != nil {
				return
			}
		}
	}

	// update flow
	raflowRespData.Flow = flow

	raflowRespData.ValidationCheck = g
	resp.Status = "success"
	resp.Record = raflowRespData

	SvcWriteResponse(d.BID, &resp, w)
}

func initRAFlowFieldsErrors(raFlowFieldsErrors *bizlogic.RAFlowFieldsErrors) {
	*raFlowFieldsErrors = bizlogic.RAFlowFieldsErrors{
		Dates: bizlogic.DatesFieldsError{
			Errors: make(map[string][]string, 0),
		},
		People: bizlogic.PeopleError{
			Total:        0,
			PeopleErrors: []bizlogic.PeopleFieldsError{},
		},
		Pets: bizlogic.PetsError{
			Total:     0,
			PetErrors: []bizlogic.PetFieldsError{},
		},
		Vehicle: bizlogic.VehiclesError{
			Total:         0,
			VehicleErrors: []bizlogic.VehicleFieldsError{},
		},
		Rentables: bizlogic.RentablesError{
			Total:          0,
			RentableErrors: []bizlogic.RentablesFieldsError{},
		},
		ParentChild: bizlogic.ParentChildrenError{
			Total:             0,
			ParentChildErrors: []bizlogic.ParentChildFieldsError{},
		},
		Tie: bizlogic.TieFieldsError{
			TiePeople: bizlogic.TiePeopleError{
				Total:           0,
				TiePeopleErrors: []bizlogic.TiePeopleFieldsError{},
			},
		},
	}
}

func initRAFlowNonFieldsErrors(raFlowNonFieldsErrors *bizlogic.RAFlowNonFieldsErrors) {
	// Initialize non fields errors
	*raFlowNonFieldsErrors = bizlogic.RAFlowNonFieldsErrors{
		Dates:       make([]string, 0),
		People:      make([]string, 0),
		Pets:        make([]string, 0),
		Vehicle:     make([]string, 0),
		Rentables:   make([]string, 0),
		ParentChild: make([]string, 0),
		Tie:         make([]string, 0),
	}
}
