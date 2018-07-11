package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"rentroll/rtags"
	"time"
)

// RAFlowDetailRequest is a struct to hold info for Flow which is going to be validate
type RAFlowDetailRequest struct {
	FlowID    int64
	UserRefNo string
}

// ValidateRAFlowResponse is struct to hold ErrorList for Flow
type ValidateRAFlowResponse struct {
	Total     int                `json:"total"`
	ErrorType string             `json:"errortype"`
	Errors    RAFlowFieldsErrors `json:"errors"`
}

// DatesFieldsError is struct to hold Errorlist for Dates section
type DatesFieldsError struct {
	Total  int                 `json:"total"`
	Errors map[string][]string `json:"errors"`
}

// PeopleFieldsError is struct to hold Errorlist for People section
type PeopleFieldsError struct {
	TMPTCID int64
	Total   int                 `json:"total"`
	Errors  map[string][]string `json:"errors"`
}

// PetFieldsError is struct to hold Errorlist for Pet section
type PetFieldsError struct {
	TMPPETID   int64
	Total      int                 `json:"total"`
	Errors     map[string][]string `json:"errors"`
	FeesErrors []RAFeesError       `json:"fees"`
}

// VehicleFieldsError is struct to hold Errorlist for Vehicle section
type VehicleFieldsError struct {
	TMPVID     int64
	Total      int                 `json:"total"`
	Errors     map[string][]string `json:"errors"`
	FeesErrors []RAFeesError       `json:"fees"`
}

// RentablesFieldsError is to hold Errorlist for Rentables section
type RentablesFieldsError struct {
	RID        int64
	Total      int                 `json:"total"`
	Errors     map[string][]string `json:"errors"`
	FeesErrors []RAFeesError       `json:"fees"`
}

// RAFeesError is struct to hold Errolist for Fees of vehicles
type RAFeesError struct {
	TMPASMID int64
	Total    int                 `json:"total"`
	Errors   map[string][]string `json:"errors"`
}

// ParentChildFieldsError is to hold Errorlist for Parent/Child section
type ParentChildFieldsError struct {
	PRID   int64               // parent rentable ID
	CRID   int64               // child rentable ID
	Total  int                 `json:"total"`
	Errors map[string][]string `json:"errors"`
}

// TiePeopleFieldsError is to hold Errorlist for TiePeople section
type TiePeopleFieldsError struct {
	TMPTCID int64
	Total   int                 `json:"total"`
	Errors  map[string][]string `json:"errors"`
}

// TieFieldsError is to hold Errorlist for Tie section
type TieFieldsError struct {
	TiePeople []TiePeopleFieldsError `json:"people"`
}

// RAFlowFieldsErrors is to hold Errorlist for each section of RAFlow
type RAFlowFieldsErrors struct {
	Dates       DatesFieldsError         `json:"dates"`
	People      []PeopleFieldsError      `json:"people"`
	Pets        []PetFieldsError         `json:"pets"`
	Vehicle     []VehicleFieldsError     `json:"vehicle"`
	Rentables   []RentablesFieldsError   `json:"rentables"`
	ParentChild []ParentChildFieldsError `json:"parentchild"`
	Tie         TieFieldsError           `json:"tie"`
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
		err                error
		foo                RAFlowDetailRequest
		raFlowData         RAFlowJSONData
		raFlowFieldsErrors RAFlowFieldsErrors
		g                  ValidateRAFlowResponse
	)

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("Only POST method is allowed")
		return
	}

	// unmarshal data into request data struct
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		return
	}

	// Init RAFlowFields error list
	raFlowFieldsErrors = RAFlowFieldsErrors{
		Dates: DatesFieldsError{
			Errors: map[string][]string{},
		},
		People:      []PeopleFieldsError{},
		Pets:        []PetFieldsError{},
		Vehicle:     []VehicleFieldsError{},
		Rentables:   []RentablesFieldsError{},
		ParentChild: []ParentChildFieldsError{},
		Tie: TieFieldsError{
			TiePeople: []TiePeopleFieldsError{},
		},
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
	g = basicValidateRAFlow(raFlowData, raFlowFieldsErrors)

	// If RAFlow structure have more than 1 basic validation error than it return with the list of basic validation errors
	if g.Total > 0 {
		SvcWriteResponse(d.BID, &g, w)
		return
	}

	// --------------------------------------------
	// Perform Bizlogic check validation on RAFlow
	// --------------------------------------------
	g = validateRAFlowBizLogic(r.Context(), &raFlowData, raFlowFieldsErrors)

	// If RAFlow structure have more than 1 biz logic check validation error than it return with the list of biz logic validation errors
	if g.Total > 0 {
		SvcWriteResponse(d.BID, &g, w)
		return
	}

}

// basicValidateRAFlow validate RAFlow's fields section wise
//-------------------------------------------------------------------------
func basicValidateRAFlow(raFlowData RAFlowJSONData, raFlowFieldsErrors RAFlowFieldsErrors) ValidateRAFlowResponse {

	var (
		datesFieldsErrors       DatesFieldsError
		peopleFieldsErrors      PeopleFieldsError
		petFieldsErrors         PetFieldsError
		vehicleFieldsErrors     VehicleFieldsError
		rentablesFieldsErrors   RentablesFieldsError
		raFeesErrors            RAFeesError
		parentChildFieldsErrors ParentChildFieldsError
		tieFieldsErrors         TieFieldsError
		tiePeopleFieldsErrors   TiePeopleFieldsError
		g                       ValidateRAFlowResponse
	)

	//----------------------------------------------
	// validate RADatesFlowData structure
	// ----------------------------------------------
	// NOTE: Validation not require for the date type fields.
	// Because it handles while Unmarshalling string into rlib.JSONDate

	// call validation function
	errs := rtags.ValidateStructFromTagRules(raFlowData.Dates)

	// Modify error count for the response
	datesFieldsErrors.Total = len(errs)
	datesFieldsErrors.Errors = errs

	// Modify Total Error
	g.Total += datesFieldsErrors.Total

	// Assign dates fields error to
	raFlowFieldsErrors.Dates = datesFieldsErrors

	//----------------------------------------------
	// validate RAPeopleFlowData structure
	// ----------------------------------------------
	for _, people := range raFlowData.People {
		// call validation function
		errs := rtags.ValidateStructFromTagRules(people)

		// Modify error count for the response
		peopleFieldsErrors.Total = len(errs)
		peopleFieldsErrors.TMPTCID = people.TMPTCID
		peopleFieldsErrors.Errors = errs

		// Modify Total Error
		g.Total += peopleFieldsErrors.Total

		// Skip the row if it doesn't have error for the any fields
		if len(errs) == 0 {
			continue
		}

		raFlowFieldsErrors.People = append(raFlowFieldsErrors.People, peopleFieldsErrors)
	}

	// ----------------------------------------------
	// validate RAPetFlowData structure
	// ----------------------------------------------
	for _, pet := range raFlowData.Pets {

		// init raFeesErrors
		raFeesErrors := RAFeesError{
			Errors: map[string][]string{},
		}

		// call validation function
		errs := rtags.ValidateStructFromTagRules(pet)

		// Modify error count for the response
		petFieldsErrors.Total = len(errs)
		petFieldsErrors.TMPPETID = pet.TMPPETID
		petFieldsErrors.Errors = errs
		petFieldsErrors.FeesErrors = make([]RAFeesError, 0)

		fmt.Printf("Petfields error: %d\n", petFieldsErrors.Total)

		// ----------------------------------------------
		// validate RAPetFlowData.Fees structure
		// ----------------------------------------------
		for _, fee := range pet.Fees {
			// call validation function
			errs := rtags.ValidateStructFromTagRules(fee)

			raFeesErrors.Total = len(errs)
			raFeesErrors.TMPASMID = fee.TMPASMID
			raFeesErrors.Errors = errs

			// Modify pets error count
			petFieldsErrors.Total += raFeesErrors.Total

			// Skip the row if it doesn't have error for the any fields
			if len(errs) == 0 {
				continue
			}

			petFieldsErrors.FeesErrors = append(petFieldsErrors.FeesErrors, raFeesErrors)
		}

		// Modify total error
		g.Total += petFieldsErrors.Total

		// If there is no error in pet than skip that pet's error being added.
		if petFieldsErrors.Total == 0 {
			continue
		}

		raFlowFieldsErrors.Pets = append(raFlowFieldsErrors.Pets, petFieldsErrors)
	}

	// ----------------------------------------------
	// validate RAVehicleFlowData structure
	// ----------------------------------------------
	for _, vehicle := range raFlowData.Vehicles {

		// init raFeesErrors
		raFeesErrors := RAFeesError{
			Errors: map[string][]string{},
		}

		// call validation function
		errs := rtags.ValidateStructFromTagRules(vehicle)

		// Modify error count for the response
		vehicleFieldsErrors.Total = len(errs)
		vehicleFieldsErrors.TMPVID = vehicle.TMPVID
		vehicleFieldsErrors.Errors = errs
		vehicleFieldsErrors.FeesErrors = make([]RAFeesError, 0)

		// ----------------------------------------------
		// validate RAVehicleFlowData.Fees structure
		// ----------------------------------------------
		for _, fee := range vehicle.Fees {

			// call validation function
			errs := rtags.ValidateStructFromTagRules(fee)

			raFeesErrors.Total = len(errs)
			raFeesErrors.TMPASMID = fee.TMPASMID
			raFeesErrors.Errors = errs

			// Modify vehicle error count
			vehicleFieldsErrors.Total += raFeesErrors.Total

			// Skip the row if it doesn't have error for the any fields
			if len(errs) == 0 {
				continue
			}

			vehicleFieldsErrors.FeesErrors = append(vehicleFieldsErrors.FeesErrors, raFeesErrors)
		}

		// Modify Total Error
		g.Total += vehicleFieldsErrors.Total

		// If there is no error in vehicle than skip that vehicle's error being added.
		if vehicleFieldsErrors.Total == 0 {
			continue
		}

		raFlowFieldsErrors.Vehicle = append(raFlowFieldsErrors.Vehicle, vehicleFieldsErrors)
	}

	// ----------------------------------------------
	// validate RARentablesFlowData structure
	// ----------------------------------------------
	for _, rentable := range raFlowData.Rentables {
		// init raFeesErrors
		raFeesErrors = RAFeesError{
			Errors: map[string][]string{},
		}

		// call validation function
		errs := rtags.ValidateStructFromTagRules(rentable)

		// Modify error count for the response
		rentablesFieldsErrors.Total = len(errs)
		rentablesFieldsErrors.RID = rentable.RID
		rentablesFieldsErrors.Errors = errs
		rentablesFieldsErrors.FeesErrors = make([]RAFeesError, 0)

		// Modify Total Error
		g.Total += rentablesFieldsErrors.Total

		// ----------------------------------------------
		// validate RAVehicleFlowData.Fees structure
		// ----------------------------------------------
		for _, fee := range rentable.Fees {

			// call validation function
			errs := rtags.ValidateStructFromTagRules(fee)

			raFeesErrors.Total = len(errs)
			raFeesErrors.TMPASMID = fee.TMPASMID
			raFeesErrors.Errors = errs

			rentablesFieldsErrors.Total += raFeesErrors.Total

			// Skip the row if it doesn't have error for the any fields
			if len(errs) == 0 {
				continue
			}

			rentablesFieldsErrors.FeesErrors = append(rentablesFieldsErrors.FeesErrors, raFeesErrors)
		}

		// Modify Total Error
		g.Total += raFeesErrors.Total

		// If there is no error in vehicle than skip that rentable's error being added.
		if rentablesFieldsErrors.Total == 0 {
			continue
		}

		raFlowFieldsErrors.Rentables = append(raFlowFieldsErrors.Rentables, rentablesFieldsErrors)
	}

	// ----------------------------------------------
	// validate RAParentChildFlowData structure
	// ----------------------------------------------
	for _, parentChild := range raFlowData.ParentChild {
		// call validation function
		errs := rtags.ValidateStructFromTagRules(parentChild)

		// Skip the row if it doesn't have error for the any fields
		if len(errs) == 0 {
			continue
		}

		// Modify error count for the response
		parentChildFieldsErrors.Total = len(errs)
		parentChildFieldsErrors.PRID = parentChild.PRID
		parentChildFieldsErrors.Errors = errs

		// Modify Total Error
		g.Total += rentablesFieldsErrors.Total

		raFlowFieldsErrors.ParentChild = append(raFlowFieldsErrors.ParentChild, parentChildFieldsErrors)
	}

	// ----------------------------------------------
	// validate RATieFlowData.People structure
	// ----------------------------------------------
	for _, people := range raFlowData.Tie.People {
		// call validation function
		errs = rtags.ValidateStructFromTagRules(people)

		// Modify error count for the response
		tiePeopleFieldsErrors.Total = len(errs)
		tiePeopleFieldsErrors.TMPTCID = people.TMPTCID
		tiePeopleFieldsErrors.Errors = errs

		// Modify Total Error
		g.Total += tiePeopleFieldsErrors.Total

		tieFieldsErrors.TiePeople = append(tieFieldsErrors.TiePeople, tiePeopleFieldsErrors)
	}

	// Assign all(people/pet/vehicles) tie related error
	raFlowFieldsErrors.Tie = tieFieldsErrors

	//---------------------------------------
	// set the response
	//---------------------------------------
	g.Errors = raFlowFieldsErrors
	g.ErrorType = "basic"

	return g
}

// validateRAFlowBizLogic is to check RAFlow's business logic
func validateRAFlowBizLogic(ctx context.Context, a *RAFlowJSONData, raFlowFieldsErrors RAFlowFieldsErrors) ValidateRAFlowResponse {
	const funcname = "ValidateRAFlowBizLogic"
	fmt.Printf("Entered %s\n", funcname)

	var (
		datesFieldsErrors     DatesFieldsError
		peopleFieldsErrors    []PeopleFieldsError
		petFieldsErrors       []PetFieldsError
		vehicleFieldsErrors   []VehicleFieldsError
		rentablesFieldsErrors []RentablesFieldsError
		g                     ValidateRAFlowResponse
	)

	// -----------------------------------------------
	// -------- Bizlogic check on date section -------
	// -----------------------------------------------
	datesFieldsErrors = validateDatesBizLogic(a.Dates)
	// Modify global error count
	g.Total += datesFieldsErrors.Total
	// Update date section error
	raFlowFieldsErrors.Dates = datesFieldsErrors

	// -----------------------------------------------
	// ------ Bizlogic check on people section -------
	// -----------------------------------------------
	peopleFieldsErrors = validatePeopleBizLogic(a.People)
	// Modify global error count
	//g.Total += peopleFieldsErrors.Total
	// Update people section error
	raFlowFieldsErrors.People = peopleFieldsErrors

	// -----------------------------------------------
	// ------- Bizlogic check on pet section ---------
	// -----------------------------------------------
	petFieldsErrors = validatePetBizLogic(a)
	raFlowFieldsErrors.Pets = petFieldsErrors

	// -----------------------------------------------
	// ------ Bizlogic check on vehicle section ------
	// -----------------------------------------------
	vehicleFieldsErrors = validateVehicleBizLogic(a)
	raFlowFieldsErrors.Vehicle = vehicleFieldsErrors

	// -----------------------------------------------
	// ---- Bizlogic check on rentables section ------
	// -----------------------------------------------
	rentablesFieldsErrors = validateRentableBizLogic(a.Rentables)
	raFlowFieldsErrors.Rentables = rentablesFieldsErrors

	// Set the response
	g.Errors = raFlowFieldsErrors
	g.ErrorType = "biz"
	return g
}

// validateDatesBizLogic Perform business logic check on date section
// ---------------------------------------------
// 1. Start dates must be prior to End/Stop date
// ---------------------------------------------
func validateDatesBizLogic(dates RADatesFlowData) DatesFieldsError {

	var (
		datesFieldsErrors DatesFieldsError
		err               error
	)

	// -----------------------------------------------
	// -------- Agreements Date check ----------------
	// -----------------------------------------------
	agreementStartDate := time.Time(dates.AgreementStart)
	agreementStopDate := time.Time(dates.AgreementStop)
	// Start date must be prior to End/Stop date
	if !agreementStartDate.Before(agreementStopDate) {

		// define and assign error
		err = fmt.Errorf("agreement start date must be prior to agreement stop date")
		datesFieldsErrors.Errors["AgreementStart"] = append(datesFieldsErrors.Errors["AgreementStart"], err.Error())

		// Modify date section error count
		datesFieldsErrors.Total++
	}

	// -----------------------------------------------
	// -------- Rent Date check ---------------------
	// -----------------------------------------------
	rentStartDate := time.Time(dates.RentStart)
	rentStopDate := time.Time(dates.RentStop)
	// Start date must be prior to End/Stop date
	if !rentStartDate.Before(rentStopDate) {

		// define and assign error
		err = fmt.Errorf("rent start date must be prior to rent stop date")
		datesFieldsErrors.Errors["RentStart"] = append(datesFieldsErrors.Errors["RentStart"], err.Error())

		// Modify date section error count
		datesFieldsErrors.Total++
	}

	// -----------------------------------------------
	// --------- Possession Date check ---------------
	// -----------------------------------------------
	possessionStartDate := time.Time(dates.PossessionStart)
	possessionStopDate := time.Time(dates.PossessionStop)
	// Start date must be prior to End/Stop date
	if !possessionStartDate.Before(possessionStopDate) {

		// define and assign error
		err = fmt.Errorf("possessions start date must be prior to possessions stop date")
		datesFieldsErrors.Errors["PossessionStart"] = append(datesFieldsErrors.Errors["PossessionStart"], err.Error())

		// Modify date section error count
		datesFieldsErrors.Total++
	}

	return datesFieldsErrors
}

// validatePeopleBizLogic Perform business logic check on people section
// ----------------------------------------------------------------------
// 1. If isCompany flag is true then CompanyName is required
// 2. If isCompany flag is false than FirstName and LastName are required
// 3. If only one person exist in the list, then it should have isRenter role marked as true.
// ----------------------------------------------------------------------
func validatePeopleBizLogic(people []RAPeopleFlowData) []PeopleFieldsError {
	var (
		peopleFieldsError  PeopleFieldsError
		peopleFieldsErrors []PeopleFieldsError
		err                error
	)

	// ----------- Check rule no. 3 ----------------
	// TODO(Akshay): Here is glitch. It will create two entry with same TMPTCID if it have error for rule no. 1/2 and 3.
	if len(people) == 1 {
		if !people[0].IsRenter {
			err = fmt.Errorf("person should be renter")
			peopleFieldsError.TMPTCID = people[0].TMPTCID
			peopleFieldsError.Errors["IsRenter"] = append(peopleFieldsError.Errors["IsRenter"], err.Error())
			peopleFieldsError.Total++
			peopleFieldsErrors = append(peopleFieldsErrors, peopleFieldsError)
		}
	}

	err = fmt.Errorf("should not be blank")
	for _, p := range people {

		peopleFieldsError.TMPTCID = p.TMPTCID

		// ----------- Check rule no. 1 and 2 ----------------
		if p.IsCompany {
			if len(p.CompanyName) == 0 {
				peopleFieldsError.Errors["CompanyName"] = append(peopleFieldsError.Errors["CompanyName"], err.Error())
				peopleFieldsError.Total++
			}
		} else {

			if len(p.FirstName) == 0 {
				peopleFieldsError.Errors["FirstName"] = append(peopleFieldsError.Errors["FirstName"], err.Error())
				peopleFieldsError.Total++
			}

			if len(p.LastName) == 0 {
				peopleFieldsError.Errors["LastName"] = append(peopleFieldsError.Errors["LastName"], err.Error())
				peopleFieldsError.Total++
			}

		}

		peopleFieldsErrors = append(peopleFieldsErrors, peopleFieldsError)
	}

	return peopleFieldsErrors
}

// validatePetBizLogic Perform business logic check on pet section
// ----------------------------------------------------------------------
// 1. Every pet must be associated with a transactant
// 2. Pets are optional. Means if HavePets is set to false in meta
// information than it should not have any pets.
// 3. DtStart must be prior to DtStop
// ----------------------------------------------------------------------
func validatePetBizLogic(a *RAFlowJSONData) []PetFieldsError {
	var (
		petFieldsError  PetFieldsError
		petFieldsErrors []PetFieldsError
		err             error
	)

	// If meta doesn't set HavePets to true than RAFlow shouldn't have any pets
	//if a.Meta.HavePets && len(a.Pets) != 0{
	//	//err = fmt.Errorf("should be")
	//	//petFieldsError.Total++
	//	//petFieldsError.Errors["pet"] = append(petFieldsError.Errors["pet"], )
	//
	//}

	// ------------- Check for rule no 1 ---------------
	for _, pet := range a.Pets {
		if !isAssociatedWithPerson(pet.TMPTCID, a.People) {
			//Error
			err = fmt.Errorf("pet must be associated with a person")

			// Modify error count
			petFieldsError.Total++

			// Get pet tmp id
			petFieldsError.TMPPETID = pet.TMPPETID

			// list error
			petFieldsError.Errors["TMPPETID"] = append(petFieldsError.Errors["TMPPETID"], err.Error())
		}

		// -----------------------------------------------
		// --------- Check for rule no 3 ---------------
		// -----------------------------------------------
		startDate := time.Time(pet.DtStart)
		stopDate := time.Time(pet.DtStop)
		// Start date must be prior to End/Stop date
		if !startDate.Before(stopDate) {

			// define and assign error
			err = fmt.Errorf("start date must be prior to stop date")
			petFieldsError.Errors["DtStart"] = append(petFieldsError.Errors["DtStart"], err.Error())

			// Modify vehicle section error count
			petFieldsError.Total++
		}

		// ---------------------------------------------------
		// --------- Biz logic check for fees section --------
		// ---------------------------------------------------
		petFieldsError.FeesErrors = validateFeesBizLogic(pet.Fees)
		petFieldsErrors = append(petFieldsErrors, petFieldsError)
	}

	return petFieldsErrors
}

// validateVehicleBizLogic Perform business logic check on vehicle section
// ----------------------------------------------------------------------
// 1. Every vehicle must be associated with a transactant
// 2. Vehicle are optional. Means if HaveVehicles is set to false in meta
// information than it should not have any vehicles.
// 3. DtStart must be prior to DtStop
// ----------------------------------------------------------------------
func validateVehicleBizLogic(a *RAFlowJSONData) []VehicleFieldsError {
	var (
		vehicleFieldsError  VehicleFieldsError
		vehicleFieldsErrors []VehicleFieldsError
		err                 error
	)

	for _, vehicle := range a.Vehicles {
		// Get vehicle tmp id
		vehicleFieldsError.TMPVID = vehicle.TMPVID

		// ------------- Check for rule no 1 ---------------
		if !isAssociatedWithPerson(vehicle.TMPTCID, a.People) {
			//Error
			err = fmt.Errorf("vehicle must be associated with a person")

			// Modify error count
			vehicleFieldsError.Total++

			// list error
			vehicleFieldsError.Errors["TMPVID"] = append(vehicleFieldsError.Errors["TMPVID"], err.Error())
		}

		// -----------------------------------------------
		// --------- Check for rule no 3 ---------------
		// -----------------------------------------------
		startDate := time.Time(vehicle.DtStart)
		stopDate := time.Time(vehicle.DtStop)
		// Start date must be prior to End/Stop date
		if !startDate.Before(stopDate) {

			// define and assign error
			err = fmt.Errorf("start date must be prior to stop date")
			vehicleFieldsError.Errors["DtStart"] = append(vehicleFieldsError.Errors["DtStart"], err.Error())

			// Modify vehicle section error count
			vehicleFieldsError.Total++
		}

		// ---------------------------------------------------
		// --------- Biz logic check for fees section --------
		// ---------------------------------------------------
		vehicleFieldsError.FeesErrors = validateFeesBizLogic(vehicle.Fees)

		vehicleFieldsErrors = append(vehicleFieldsErrors, vehicleFieldsError)
	}

	return vehicleFieldsErrors
}

// validateRentableBizLogic Perform business logic check on rentable section
// ----------------------------------------------------------------------
// 1. There must be one parent rentables available. (Parent rentables decide based on RTFlags)
// 2. For every rentables, there must be one entry for the Fees.
// ----------------------------------------------------------------------
func validateRentableBizLogic(rentables []RARentablesFlowData) []RentablesFieldsError {

	var (
		rentablesFieldsError  RentablesFieldsError
		rentablesFieldsErrors []RentablesFieldsError
		err                   error
	)

	parentRentableCount := 0

	for _, rentable := range rentables {

		// There must be one entry for the Fees
		// ----------- Check for rule no 2 ------------
		if len(rentable.Fees) < 1 {
			err = fmt.Errorf("should be one entry for the fees")
			rentablesFieldsError.Total++
			rentablesFieldsError.Errors["Fees"] = append(rentablesFieldsError.Errors["Fees"], err.Error())
		}

		// Check if rentable is parent. If yes than increment parentRentableCount
		if rentable.RTFLAGS&(1<<1) == 0 {
			parentRentableCount++
		}

		// ---------------------------------------------------
		// --------- Biz logic check for fees section --------
		// ---------------------------------------------------
		rentablesFieldsError.FeesErrors = validateFeesBizLogic(rentable.Fees)

		// Modify rentable error list
		rentablesFieldsErrors = append(rentablesFieldsErrors, rentablesFieldsError)
	}

	// There must be one parent rentable
	// TODO(Akshay): Add this error to rentables
	//if parentRentableCount < 1 {
	//	err = fmt.Errorf("should be at least one parent rentable")
	//}

	return rentablesFieldsErrors
}

// validateFeesBizLogic perform business logic check on fees section
// ----------------------------------------------------------------------
// 1. Start date must be prior to Stop date
// ----------------------------------------------------------------------
func validateFeesBizLogic(fees []RAFeesData) []RAFeesError {
	var (
		raFeesError  RAFeesError
		raFeesErrors []RAFeesError
		err          error
	)

	for _, fee := range fees {
		// -----------------------------------------------
		// --------- Check for rule no 1 ---------------
		// -----------------------------------------------
		startDate := time.Time(fee.Start)
		stopDate := time.Time(fee.Stop)
		// Start date must be prior to End/Stop date
		if !startDate.Before(stopDate) {
			// define and assign error
			err = fmt.Errorf("start date must be prior to stop date")
			raFeesError.TMPASMID = fee.TMPASMID
			raFeesError.Errors["Start"] = append(raFeesError.Errors["Stop"], err.Error())

			// Modify vehicle section error count
			raFeesError.Total++
		}

		raFeesErrors = append(raFeesErrors, raFeesError)
	}

	return raFeesErrors
}

// isAssociatedWithPerson Check Pets/Vehicles is associated with Person or not
func isAssociatedWithPerson(TMPTCID int64, people []RAPeopleFlowData) bool {
	for _, p := range people {
		if p.TMPTCID == TMPTCID {
			return true
		} else {
			continue
		}
	}
	return false
}
