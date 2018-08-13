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
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

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
//-------------------------------------------------------------------------
func ValidateRAFlow2(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "ValidateRAFlow"
	fmt.Printf("Entered %s\n", funcname)

	var (
		err                error
		foo                RAFlowDetailRequest
		raFlowData         rlib.RAFlowJSONData
		raFlowFieldsErrors bizlogic.RAFlowFieldsErrors
		g                  bizlogic.ValidateRAFlowResponse
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

	raFlowFieldsErrors = bizlogic.RAFlowFieldsErrors{
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

	// ---------------------------------------
	// Perform basic validation on RAFlow
	// ---------------------------------------
	bizlogic.ValidateRAFlowBasic(r.Context(), &raFlowFieldsErrors, &raFlowData, &g)

	// If RAFlow structure have more than 1 basic validation error than it return with the list of basic validation errors
	//if g.Total > 0 {
	//	SvcWriteResponse(d.BID, &g, w)
	//	return
	//}

	// --------------------------------------------
	// Perform Bizlogic check validation on RAFlow
	// --------------------------------------------
	bizlogic.ValidateRAFlowBizLogic(r.Context(), &raFlowFieldsErrors, &raFlowData, &g, flow.ID)

	g.Errors = raFlowFieldsErrors

	// If RAFlow structure have more than 1 biz logic check validation error than it return with the list of biz logic validation errors
	if g.Total > 0 {
		SvcWriteResponse(d.BID, &g, w)
		return
	}

	SvcWriteResponse(d.BID, &g, w)
}

// ValidateRAFlow validate RAFlow's fields section wise
func ValidateRAFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "ValidateRAFlow"
	fmt.Printf("Entered %s\n", funcname)

	var (
		err                   error
		foo                   RAFlowDetailRequest
		raFlowData            rlib.RAFlowJSONData
		raFlowFieldsErrors    bizlogic.RAFlowFieldsErrors
		raFlowNonFieldsErrors bizlogic.RAFlowNonFieldsErrors
		g                     bizlogic.ValidateRAFlowResponse
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

	// init raFlowFieldsErrors
	initRAFlowFieldsErrors(&raFlowFieldsErrors)

	initRAFlowNonFieldsErrors(&raFlowNonFieldsErrors)

	bizlogic.ValidateRAFlowParts(r.Context(), &raFlowFieldsErrors, &raFlowNonFieldsErrors, &raFlowData, flow.ID)

	totalFieldsError := raFlowFieldsErrors.Dates.Total + raFlowFieldsErrors.People.Total + raFlowFieldsErrors.Pets.Total + raFlowFieldsErrors.Vehicle.Total + raFlowFieldsErrors.Rentables.Total + raFlowFieldsErrors.ParentChild.Total + raFlowFieldsErrors.Tie.TiePeople.Total
	totalNonFieldsError := len(raFlowNonFieldsErrors.Dates) + len(raFlowNonFieldsErrors.People) + len(raFlowNonFieldsErrors.Pets) + len(raFlowNonFieldsErrors.Rentables) + len(raFlowNonFieldsErrors.Vehicle) + len(raFlowNonFieldsErrors.ParentChild) + len(raFlowNonFieldsErrors.Tie)
	g.Total += totalFieldsError + totalNonFieldsError
	g.Errors = raFlowFieldsErrors
	g.NonFieldsErrors = raFlowNonFieldsErrors
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
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
