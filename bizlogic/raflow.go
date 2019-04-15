package bizlogic

import (
	"context"
	"fmt"
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
	Total           int                   `json:"total"`
	Errors          RAFlowFieldsErrors    `json:"errors"`
	NonFieldsErrors RAFlowNonFieldsErrors `json:"nonFieldsErrors"`
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
	TMPPETID  int64
	Total     int                 `json:"total"`
	Errors    map[string][]string `json:"errors"`
	FeesError FeesError           `json:"fees"`
}

// FeesError is struct to hold total and error list of fees
type FeesError struct {
	Total      int           `json:"total"`
	FeesErrors []RAFeesError `json:"errors"`
}

// VehicleFieldsError is struct to hold Errorlist for Vehicle section
type VehicleFieldsError struct {
	TMPVID    int64
	Total     int                 `json:"total"`
	Errors    map[string][]string `json:"errors"`
	FeesError FeesError           `json:"fees"`
}

// RentablesFieldsError is to hold Errorlist for Rentables section
type RentablesFieldsError struct {
	RID       int64
	Total     int                 `json:"total"`
	Errors    map[string][]string `json:"errors"`
	FeesError FeesError           `json:"fees"`
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
	TiePeople TiePeopleError `json:"people"`
}

// PeopleError is to hold list of people error and total count
type PeopleError struct {
	Total        int                 `json:"total"`
	PeopleErrors []PeopleFieldsError `json:"errors"`
}

// PetsError is to hold list of pet error and total count
type PetsError struct {
	Total     int              `json:"total"`
	PetErrors []PetFieldsError `json:"errors"`
}

// VehiclesError is to hold list of vehicle error and total count
type VehiclesError struct {
	Total         int                  `json:"total"`
	VehicleErrors []VehicleFieldsError `json:"errors"`
}

// RentablesError is to hold list of rentable error and total count
type RentablesError struct {
	Total          int                    `json:"total"`
	RentableErrors []RentablesFieldsError `json:"errors"`
}

// ParentChildrenError is to hold list of parent child error and total count
type ParentChildrenError struct {
	Total             int                      `json:"total"`
	ParentChildErrors []ParentChildFieldsError `json:"errors"`
}

// TiePeopleError is to hold list of tie people error and total error count
type TiePeopleError struct {
	Total           int                    `json:"total"`
	TiePeopleErrors []TiePeopleFieldsError `json:"errors"`
}

// RAFlowFieldsErrors is to hold Errorlist for each section of RAFlow
type RAFlowFieldsErrors struct {
	Dates       DatesFieldsError    `json:"dates"`
	People      PeopleError         `json:"people"`
	Pets        PetsError           `json:"pets"`
	Vehicle     VehiclesError       `json:"vehicles"`
	Rentables   RentablesError      `json:"rentables"`
	ParentChild ParentChildrenError `json:"parentchild"`
	Tie         TieFieldsError      `json:"tie"`
}

// RAFlowNonFieldsErrors is to hold non fields error
type RAFlowNonFieldsErrors struct {
	Dates       []string `json:"dates"`
	People      []string `json:"people"`
	Pets        []string `json:"pets"`
	Vehicle     []string `json:"vehicles"`
	Rentables   []string `json:"rentables"`
	ParentChild []string `json:"parentchild"`
	Tie         []string `json:"tie"`
}

// ValidateRAFlowParts It checks for basic and biz rules for raflow data
func ValidateRAFlowParts(ctx context.Context, raFlowFieldsErrors *RAFlowFieldsErrors, raFlowNonFieldsErrors *RAFlowNonFieldsErrors, a *rlib.RAFlowJSONData, RAID int64) error {

	var err error

	//----------------------------------------------
	// validate RADatesFlowData structure
	// ----------------------------------------------
	err = validateDates(ctx, raFlowFieldsErrors, raFlowNonFieldsErrors, a, RAID)
	if err != nil {
		return err
	}

	// ----------------------------------------------
	// validate RAPeopleFlowData structure
	// ----------------------------------------------
	err = validatePeople(ctx, raFlowFieldsErrors, raFlowNonFieldsErrors, a, RAID)
	if err != nil {
		return err
	}

	// ----------------------------------------------
	// validate RAPetFlowData structure
	// ----------------------------------------------
	err = validatePets(ctx, raFlowFieldsErrors, raFlowNonFieldsErrors, a)
	if err != nil {
		return err
	}

	// ----------------------------------------------
	// validate RAVehicleFlowData structure
	// ----------------------------------------------
	err = validateVehicles(ctx, raFlowFieldsErrors, raFlowNonFieldsErrors, a)
	if err != nil {
		return err
	}

	// ----------------------------------------------
	// validate RARentablesFlowData structure
	// ----------------------------------------------
	err = validateRentables(ctx, raFlowFieldsErrors, raFlowNonFieldsErrors, a)
	if err != nil {
		return err
	}

	// ----------------------------------------------
	// validate RAParentChildFlowData structure
	// ----------------------------------------------
	err = validateParentChild(ctx, raFlowFieldsErrors, raFlowNonFieldsErrors, a)
	if err != nil {
		return err
	}

	// ----------------------------------------------
	// validate RATieFlowData.People structure
	// ----------------------------------------------
	err = validateTiePeople(ctx, raFlowFieldsErrors, raFlowNonFieldsErrors, a)
	if err != nil {
		return err
	}

	return nil
}

// validateDates
// BizCheck
// ---------------------------------------------
// 1. Start dates must be prior to End/Stop date
// 2. If RAID > 0 then the all Start Dates on the Dates/Agent flow part must be >= Start dates on the RAID
// ---------------------------------------------
func validateDates(ctx context.Context, raFlowFieldsErrors *RAFlowFieldsErrors, raFlowNonFieldsErrors *RAFlowNonFieldsErrors, a *rlib.RAFlowJSONData, RAID int64) error {

	var (
		err error
	)

	// NOTE: Validation not require for the date type fields.
	// Because it handles while Unmarshalling string into rlib.JSONDate

	dates := a.Dates
	// call validation function
	errs := rtags.ValidateStructFromTagRules(dates)
	// Modify error count for the response and initialize error object
	datesFieldsErrors := DatesFieldsError{
		Total:  len(errs),
		Errors: errs,
	}

	// -----------------------------------------------
	// -------- Agreements Date check ----------------
	// -----------------------------------------------
	agreementStartDate := time.Time(dates.AgreementStart)
	agreementStopDate := time.Time(dates.AgreementStop)
	// Start date must be prior to End/Stop date
	if agreementStartDate.After(agreementStopDate) {

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
	if rentStartDate.After(rentStopDate) {

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
	if possessionStartDate.After(possessionStopDate) {

		// define and assign error
		err = fmt.Errorf("possessions start date must be prior to possessions stop date")
		datesFieldsErrors.Errors["PossessionStart"] = append(datesFieldsErrors.Errors["PossessionStart"], err.Error())

		// Modify date section error count
		datesFieldsErrors.Total++
	}

	// --------------------------------------------------
	// 2. If RAID > 0 then the all Start Dates on the Dates/Agent flow part must be >= Start dates on the RAID
	// --------------------------------------------------
	if RAID > 0 {

		ra, err := rlib.GetRentalAgreement(ctx, RAID)
		if err != nil {
			return err
		}

		raAgreementStartDate := time.Time(ra.AgreementStart)
		raRentStartDate := time.Time(ra.RentStart)
		raPossessionStartDate := time.Time(ra.PossessionStart)

		if !(agreementStartDate.Equal(raAgreementStartDate) || agreementStartDate.After(raAgreementStartDate)) {
			// define and assign error
			err = fmt.Errorf("agreement start date must be after or equal to RAID: %d agreement start date", RAID)
			datesFieldsErrors.Errors["AgreementStart"] = append(datesFieldsErrors.Errors["AgreementStart"], err.Error())

			// Modify date section error count
			datesFieldsErrors.Total++
		}

		if !(rentStartDate.Equal(raRentStartDate) || rentStartDate.After(raRentStartDate)) {

			// define and assign error
			err = fmt.Errorf("rent start date must be after or equal to RAID: %d rent start date", RAID)
			datesFieldsErrors.Errors["RentStart"] = append(datesFieldsErrors.Errors["RentStart"], err.Error())

			// Modify date section error count
			datesFieldsErrors.Total++
		}

		if !(possessionStartDate.Equal(raPossessionStartDate) || possessionStartDate.After(raPossessionStartDate)) {

			// define and assign error
			err = fmt.Errorf("possession start date must be after or equal to RAID: %d possession start date", RAID)
			datesFieldsErrors.Errors["PossessionStart"] = append(datesFieldsErrors.Errors["PossessionStart"], err.Error())

			// Modify date section error count
			datesFieldsErrors.Total++
		}

	}

	raFlowFieldsErrors.Dates = datesFieldsErrors

	return nil

}

// validatePeople
// BizCheck
// ----------------------------------------------------------------------
// 1. If isCompany flag is true then CompanyName is required
// 2. If isCompany flag is false than FirstName and LastName are required
// 3. If only one person exist in the list, then it should have isRenter role marked as true.
// 4. If role is set to Renter or guarantor than it must have mentioned GrossIncome
// 5. Either Workphone or CellPhone is compulsory when a transanctant isn't occupant
// 6. EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyEmail are required when IsCompany flag is false.
// 7. SourceSLSID must be greater than 0 when role is set to Renter
// 8. When it is brand new RA Application(RAID==0) it require "current" address related information
// 9. TaxpayorID is only require when role is set to Renter or Guarantor
// 10. Occupantion is only require when role is set to Renter or Gurantor
// 11. Primary email is only require when role is set to Renter or Gurantor
// ----------------------------------------------------------------------
func validatePeople(ctx context.Context, raFlowFieldsErrors *RAFlowFieldsErrors, raFlowNonFieldsErrors *RAFlowNonFieldsErrors, a *rlib.RAFlowJSONData, RAID int64) error {

	var (
		err error
	)

	people := a.People

	for _, p := range people {
		// call validation function
		errs := rtags.ValidateStructFromTagRules(p)

		// Modify error count for the response
		peopleFieldsError := PeopleFieldsError{
			Total:   len(errs),
			TMPTCID: p.TMPTCID,
			Errors:  errs,
		}

		//-----------------------
		// Business logic check
		//-----------------------
		err = fmt.Errorf("must not be blank")
		// ----------- Check rule no. 1  ----------------
		// If isCompany flag is true then CompanyName is required
		if p.IsCompany && len(p.CompanyName) == 0 {
			peopleFieldsError.Errors["CompanyName"] = append(peopleFieldsError.Errors["CompanyName"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 2  ----------------
		// If isCompany flag is false than FirstName and LastName are required
		if !p.IsCompany && len(p.FirstName) == 0 {
			peopleFieldsError.Errors["FirstName"] = append(peopleFieldsError.Errors["FirstName"], err.Error())
			peopleFieldsError.Total++
		}

		if !p.IsCompany && len(p.LastName) == 0 {
			peopleFieldsError.Errors["LastName"] = append(peopleFieldsError.Errors["LastName"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 6  ----------------
		// EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyEmail are required when IsCompany flag is false.
		if !p.IsCompany && p.EmergencyContactName == "" {
			peopleFieldsError.Errors["EmergencyContactName"] = append(peopleFieldsError.Errors["EmergencyContactName"], err.Error())
			peopleFieldsError.Total++
		}

		if !p.IsCompany && p.EmergencyContactAddress == "" {
			peopleFieldsError.Errors["EmergencyContactAddress"] = append(peopleFieldsError.Errors["EmergencyContactAddress"], err.Error())
			peopleFieldsError.Total++
		}

		if !p.IsCompany && p.EmergencyContactTelephone == "" {
			peopleFieldsError.Errors["EmergencyContactTelephone"] = append(peopleFieldsError.Errors["EmergencyContactTelephone"], err.Error())
			peopleFieldsError.Total++
		}

		if !p.IsCompany && p.EmergencyContactEmail == "" {
			peopleFieldsError.Errors["EmergencyContactEmail"] = append(peopleFieldsError.Errors["EmergencyContactEmail"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 8  ----------------
		// When it is brand new RA Application(RAID==0) it require "current" address related information
		if p.CurrentAddress == "" && RAID == 0 {
			peopleFieldsError.Errors["CurrentAddress"] = append(peopleFieldsError.Errors["CurrentAddress"], err.Error())
			peopleFieldsError.Total++
		}

		if p.CurrentLandLordName == "" && RAID == 0 {
			peopleFieldsError.Errors["CurrentLandLordName"] = append(peopleFieldsError.Errors["CurrentLandLordName"], err.Error())
			peopleFieldsError.Total++
		}

		if p.CurrentLandLordPhoneNo == "" && RAID == 0 {
			peopleFieldsError.Errors["CurrentLandLordPhoneNo"] = append(peopleFieldsError.Errors["CurrentLandLordPhoneNo"], err.Error())
			peopleFieldsError.Total++
		}

		if p.CurrentLengthOfResidency == "" && RAID == 0 {
			peopleFieldsError.Errors["CurrentLengthOfResidency"] = append(peopleFieldsError.Errors["CurrentLengthOfResidency"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 4  ----------------
		// If role is set to Renter or guarantor than it must have mentioned GrossIncome
		err = fmt.Errorf("gross income must be greater than 0.00")
		if (p.IsRenter || p.IsGuarantor) && !(p.GrossIncome > 0.00) {
			peopleFieldsError.Errors["GrossIncome"] = append(peopleFieldsError.Errors["GrossIncome"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 5  ----------------
		// Either Workphone or CellPhone is compulsory when a transanctant isn't occupant
		err = fmt.Errorf("provide workphone or cellphone number")
		if p.WorkPhone == "" && p.CellPhone == "" && (p.IsRenter || p.IsGuarantor) {
			peopleFieldsError.Errors["WorkPhone"] = append(peopleFieldsError.Errors["WorkPhone"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 7  ----------------
		// SourceSLSID must be greater than 0 when role is set to Renter, User
		err = fmt.Errorf("provide SourceSLSID")
		if p.IsRenter && !(p.SourceSLSID > 0) {
			peopleFieldsError.Errors["SourceSLSID"] = append(peopleFieldsError.Errors["SourceSLSID"], err.Error())
			peopleFieldsError.Total++
		}

		err = fmt.Errorf("must provide reason")
		if p.CurrentReasonForMoving == 0 && RAID == 0 {
			peopleFieldsError.Errors["CurrentReasonForMoving"] = append(peopleFieldsError.Errors["CurrentReasonForMoving"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 9  ----------------
		// 9.TaxpayorID is only require when role is set to Renter or Guarantor
		err = fmt.Errorf("no taxpayer ID available")
		if (p.IsRenter || p.IsGuarantor) && p.TaxpayorID == "" {
			peopleFieldsError.Errors["TaxpayorID"] = append(peopleFieldsError.Errors["TaxpayorID"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 10  ----------------
		// 10. Occupantion is only require when role is set to Renter or Guarantor
		err = fmt.Errorf("must not be blank")
		if (p.IsRenter || p.IsGuarantor) && p.Occupation == "" {
			peopleFieldsError.Errors["Occupation"] = append(peopleFieldsError.Errors["Occupation"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 11  ----------------
		// 11. Primary email is only require when role is set to Renter or Guarantor
		if (p.IsRenter || p.IsGuarantor) && p.PrimaryEmail == "" {
			peopleFieldsError.Errors["PrimaryEmail"] = append(peopleFieldsError.Errors["PrimaryEmail"], err.Error())
			peopleFieldsError.Total++
		}

		// ----------- Check rule no. 12  ----------------
		// 12. Driving Lic is only require when role is set to Renter or Guarantor
		if (p.IsRenter || p.IsGuarantor) && p.DriversLicense == "" {
			peopleFieldsError.Errors["DriversLicense"] = append(peopleFieldsError.Errors["DriversLicense"], err.Error())
			peopleFieldsError.Total++
		}

		// Skip the row if it doesn't have error for the any fields
		if len(peopleFieldsError.Errors) > 0 {
			raFlowFieldsErrors.People.PeopleErrors = append(raFlowFieldsErrors.People.PeopleErrors, peopleFieldsError)
		}

		// Modify Total Error
		raFlowFieldsErrors.People.Total += peopleFieldsError.Total
	}

	// ----------- Check rule no. 3 ----------------
	// If only one person exist in the list, then it should have isRenter role marked as true
	if len(people) == 1 && !people[0].IsRenter {
		err = fmt.Errorf("person must be renter")

		if len(raFlowFieldsErrors.People.PeopleErrors) == 1 {
			raFlowFieldsErrors.People.PeopleErrors[0].Errors["IsRenter"] = append(raFlowFieldsErrors.People.PeopleErrors[0].Errors["IsRenter"], err.Error())
			raFlowFieldsErrors.People.PeopleErrors[0].Total++
		} else {
			var peopleFieldsError = PeopleFieldsError{
				Errors: make(map[string][]string, 0),
			}

			peopleFieldsError.TMPTCID = people[0].TMPTCID
			peopleFieldsError.Errors["IsRenter"] = append(peopleFieldsError.Errors["IsRenter"], err.Error())
			peopleFieldsError.Total++
			raFlowFieldsErrors.People.PeopleErrors = append(raFlowFieldsErrors.People.PeopleErrors, peopleFieldsError)
		}

		raFlowFieldsErrors.People.Total++
	}

	return nil
}

// validatePets is to check Pets basic and business logic
// BizCheck
// ----------------------------------------------------------------------
// 1. Every pet must be associated with a transactant
// ----------------------------------------------------------------------
func validatePets(ctx context.Context, raFlowFieldsErrors *RAFlowFieldsErrors, raFlowNonFieldsErrors *RAFlowNonFieldsErrors, a *rlib.RAFlowJSONData) error {
	// ----------------------------------------------
	// validate RAPetFlowData structure
	// ----------------------------------------------
	for _, pet := range a.Pets {

		// call validation function
		errs := rtags.ValidateStructFromTagRules(pet)

		// Modify error count for the response
		petFieldsErrors := PetFieldsError{
			Total:    len(errs),
			TMPPETID: pet.TMPPETID,
			Errors:   errs,
			FeesError: FeesError{
				Total:      0,
				FeesErrors: make([]RAFeesError, 0),
			},
		}

		if !isAssociatedWithPerson(pet.TMPTCID, a.People) {
			//Error
			err := fmt.Errorf("pet must be associated with a person")
			// list error
			petFieldsErrors.Errors["TMPTCID"] = append(petFieldsErrors.Errors["TMPTCID"], err.Error())
			// Modify error count
			petFieldsErrors.Total++
		}

		// ----------------------------------------------
		// validate RAPetFlowData.Fees structure
		// ----------------------------------------------
		err := validateFees(ctx, &petFieldsErrors.FeesError, pet.Fees)
		if err != nil {
			return err
		}

		petFieldsErrors.Total += petFieldsErrors.FeesError.Total

		// If there is no error in pet than skip that pet's error being added.
		if petFieldsErrors.Total > 0 {
			raFlowFieldsErrors.Pets.PetErrors = append(raFlowFieldsErrors.Pets.PetErrors, petFieldsErrors)
			// Modify total error
			raFlowFieldsErrors.Pets.Total += petFieldsErrors.Total
		}
	}

	return nil
}

// validateVehicles
// BizCheck
// ----------------------------------------------------------------------
// 1. Every vehicle must be associated with a transactant
// ----------------------------------------------------------------------
func validateVehicles(ctx context.Context, raFlowFieldsErrors *RAFlowFieldsErrors, raFlowNonFieldsErrors *RAFlowNonFieldsErrors, a *rlib.RAFlowJSONData) error {
	var (
		err error
	)
	for _, vehicle := range a.Vehicles {

		// call validation function
		errs := rtags.ValidateStructFromTagRules(vehicle)

		// Modify error count for the response
		vehicleFieldsError := VehicleFieldsError{
			Total:  len(errs),
			TMPVID: vehicle.TMPVID,
			Errors: errs,
			FeesError: FeesError{
				Total:      0,
				FeesErrors: make([]RAFeesError, 0),
			},
		}

		// ------------- Check for rule no 1 ---------------
		if !isAssociatedWithPerson(vehicle.TMPTCID, a.People) {
			//Error
			err = fmt.Errorf("vehicle must be associated with a person")

			// Modify error count
			vehicleFieldsError.Total++

			// list error
			vehicleFieldsError.Errors["TMPTCID"] = append(vehicleFieldsError.Errors["TMPTCID"], err.Error())
		}

		// ----------------------------------------------
		// validate RAVehicleFlowData.Fees structure
		// ----------------------------------------------
		err := validateFees(ctx, &vehicleFieldsError.FeesError, vehicle.Fees)
		if err != nil {
			return err
		}
		vehicleFieldsError.Total += vehicleFieldsError.FeesError.Total

		// If there is no error in vehicle than skip that vehicle's error being added.
		if vehicleFieldsError.Total > 0 {
			raFlowFieldsErrors.Vehicle.VehicleErrors = append(raFlowFieldsErrors.Vehicle.VehicleErrors, vehicleFieldsError)
			// Modify Total Error
			raFlowFieldsErrors.Vehicle.Total += vehicleFieldsError.Total
		}
	}

	return nil
}

// validateRentableBizLogic Perform business logic check on rentable section
// ----------------------------------------------------------------------
// 1. There must be one parent rentables available. (Parent rentables decide based on RTFlags)
// ----------------------------------------------------------------------
func validateRentables(ctx context.Context, raFlowFieldsErrors *RAFlowFieldsErrors, raFlowNonFieldsErrors *RAFlowNonFieldsErrors, a *rlib.RAFlowJSONData) error {
	var (
		err error
	)

	rentables := a.Rentables
	parentRentableCount := 0

	for _, rentable := range rentables {

		// call validation function
		errs := rtags.ValidateStructFromTagRules(rentable)

		// Modify error count for the response
		rentablesFieldsError := RentablesFieldsError{
			Total:  len(errs),
			RID:    rentable.RID,
			Errors: errs,
			FeesError: FeesError{
				Total:      0,
				FeesErrors: make([]RAFeesError, 0),
			},
		}

		// Check if rentable is parent. If yes than increment parentRentableCount
		// And use this count to check there is parent rentable exists or not.
		if rentable.RTFLAGS&(1<<1) == 0 {
			parentRentableCount++
		}

		// ----------------------------------------------
		// validate RARentableFlowData.Fees structure
		// ----------------------------------------------
		err := validateFees(ctx, &rentablesFieldsError.FeesError, rentable.Fees)
		if err != nil {
			return err
		}
		rentablesFieldsError.Total += rentablesFieldsError.FeesError.Total

		// If there is no error in vehicle than skip that rentable's error being added.
		if rentablesFieldsError.Total > 0 {
			raFlowFieldsErrors.Rentables.RentableErrors = append(raFlowFieldsErrors.Rentables.RentableErrors, rentablesFieldsError)
			// Modify Total Error
			raFlowFieldsErrors.Rentables.Total += rentablesFieldsError.Total
		}
	}

	// There must be one parent rentable
	if !(parentRentableCount > 0) {
		err = fmt.Errorf("must have at least one parent rentable")
		raFlowNonFieldsErrors.Rentables = append(raFlowNonFieldsErrors.Rentables, err.Error())
	}

	return nil
}

// validateFees
// BizCheck
// ----------------------------------------------------------------------
// 1. Start date must be prior or equal to Stop date
// 2. Check fee must be exist in the database
// ----------------------------------------------------------------------
func validateFees(ctx context.Context, feesError *FeesError, fees []rlib.RAFeesData) error {
	for _, fee := range fees {
		// call validation function
		errs := rtags.ValidateStructFromTagRules(fee)

		raFeesErrors := RAFeesError{
			Total:    len(errs),
			TMPASMID: fee.TMPASMID,
			Errors:   errs,
		}

		// -----------------------------------------------
		// --------- Check for rule no 1 ---------------
		// -----------------------------------------------
		startDate := time.Time(fee.Start)
		stopDate := time.Time(fee.Stop)
		// Start date must be prior to End/Stop date
		if !(startDate.Equal(stopDate) || startDate.Before(stopDate)) {
			// define and assign error
			err := fmt.Errorf("start date must be prior to stop date")
			raFeesErrors.Errors["Start"] = append(raFeesErrors.Errors["Start"], err.Error())
			// Modify fees section error count
			raFeesErrors.Total++
		}

		// -----------------------------------------------
		// --------- Check for rule no 2 -----------------
		// -----------------------------------------------
		// 2. Check fee must be exist in the database
		ar, err := rlib.GetAR(ctx, fee.ARID)
		if err != nil {
			return err
		}

		if !(ar.ARID > 0) {
			err = fmt.Errorf("fee associated account rule doesn't exist")
			raFeesErrors.Errors["ARID"] = append(raFeesErrors.Errors["ARID"], err.Error())
			// Modify fees section error count
			raFeesErrors.Total++
		}

		// Skip the row if it doesn't have error for the any fields
		if len(raFeesErrors.Errors) > 0 {
			feesError.FeesErrors = append(feesError.FeesErrors, raFeesErrors)
			feesError.Total += len(raFeesErrors.Errors)
		}
	}

	return nil
}

// validateParentChild
// BizCheck
// ----------------------------------------------------------------------
// 1. If there are any entries are in the list then id of parent/child rentable must be greater than 0. Also check does it exist in database?
// ----------------------------------------------------------------------
func validateParentChild(ctx context.Context, raFlowFieldsErrors *RAFlowFieldsErrors, raFlowNonFieldsErrors *RAFlowNonFieldsErrors, a *rlib.RAFlowJSONData) error {
	for _, pc := range a.ParentChild {
		// call validation function
		errs := rtags.ValidateStructFromTagRules(pc)

		// Modify error count for the response
		parentChildFieldsError := ParentChildFieldsError{
			Total:  len(errs),
			PRID:   pc.PRID,
			CRID:   pc.CRID,
			Errors: errs,
		}

		// Check PRID exists in database which refer to RID in rentable table
		r, err := rlib.GetRentable(ctx, pc.PRID)
		if err != nil {
			return err
		}
		// Not exist than RID will be 0
		if !(r.RID > 0 && pc.PRID > 0) {
			err = fmt.Errorf("parent rentable must exists")
			parentChildFieldsError.Errors["PRID"] = append(parentChildFieldsError.Errors["PRID"], err.Error())
			parentChildFieldsError.Total++
		}

		// Check CRID exists in database which refer to RID in rentable table
		r, err = rlib.GetRentable(ctx, pc.CRID)
		if err != nil {
			return err
		}
		// Not exist than RID will be 0
		if !(r.RID > 0 && pc.CRID > 0) {
			err = fmt.Errorf("child rentable must exists")
			parentChildFieldsError.Errors["CRID"] = append(parentChildFieldsError.Errors["CRID"], err.Error())
			parentChildFieldsError.Total++
		}

		if parentChildFieldsError.Total > 0 {
			raFlowFieldsErrors.ParentChild.ParentChildErrors = append(raFlowFieldsErrors.ParentChild.ParentChildErrors, parentChildFieldsError)
			// Modify Total Error
			raFlowFieldsErrors.ParentChild.Total += parentChildFieldsError.Total
		}
	}

	return nil
}

// validateTiePeople
// BizCheck
// ----------------------------------------------------------------------
// 1. PRID must be greater than 0. It should exists in database
// 2. Person must be occupant.
// ----------------------------------------------------------------------
func validateTiePeople(ctx context.Context, raFlowFieldsErrors *RAFlowFieldsErrors, raFlowNonFieldsErrors *RAFlowNonFieldsErrors, a *rlib.RAFlowJSONData) error {

	occupantCount := 0

	for _, p := range a.Tie.People {
		// call validation function
		errs := rtags.ValidateStructFromTagRules(p)

		// Modify error count for the response
		tiePeopleFieldsError := TiePeopleFieldsError{
			Total:   len(errs),
			TMPTCID: p.TMPTCID,
			Errors:  errs,
		}

		// ---------- Check rule no 1 ---------------
		// 1. PRID must be greater than 0. It should exists in database
		// Check PRID exists in database which refer to RID in rentable table
		r, err := rlib.GetRentable(ctx, p.PRID)
		if err != nil {
			return err
		}
		// Not exist than RID will be 0
		if !(r.RID > 0 && p.PRID > 0) {
			err = fmt.Errorf("parent rentable must be tied")
			tiePeopleFieldsError.Errors["PRID"] = append(tiePeopleFieldsError.Errors["PRID"], err.Error())
			tiePeopleFieldsError.Total++
		}

		// ---------- Check rule no 2 ---------------
		// 2. Person must be occupant.
		if !isPersonOccupant(p.TMPTCID, a.People) {
			// Person is not occupant
			err = fmt.Errorf("person must be an occupant")
			tiePeopleFieldsError.Errors["IsOccupant"] = append(tiePeopleFieldsError.Errors["IsOccupant"], err.Error())
			tiePeopleFieldsError.Total++
		} else {
			// Person is occupant
			occupantCount++
		}

		if tiePeopleFieldsError.Total > 0 {
			raFlowFieldsErrors.Tie.TiePeople.TiePeopleErrors = append(raFlowFieldsErrors.Tie.TiePeople.TiePeopleErrors, tiePeopleFieldsError)
			// Modify Total Error
			raFlowFieldsErrors.Tie.TiePeople.Total += tiePeopleFieldsError.Total
		}
	}

	if !(occupantCount > 0) {
		err := fmt.Errorf("must have at least one occupant")
		raFlowNonFieldsErrors.Tie = append(raFlowNonFieldsErrors.Tie, err.Error())
	}

	return nil
}

// isPersonOccupant ensures that person TMPTCID has occupant status
func isPersonOccupant(TMPTCID int64, people []rlib.RAPeopleFlowData) bool {
	for _, p := range people {
		if p.TMPTCID == TMPTCID && p.IsOccupant {
			return true
		}
		continue
	}
	return false
}

// isAssociatedWithPerson Check Pets/Vehicles is associated with Person or not
func isAssociatedWithPerson(TMPTCID int64, people []rlib.RAPeopleFlowData) bool {
	for _, p := range people {
		if p.TMPTCID == TMPTCID {
			return true
		}
		continue
	}
	return false
}

// DataFulfilledRAFlow Check flow data is fulfilled or not.
func DataFulfilledRAFlow(ctx context.Context, a *rlib.RAFlowJSONData, d *rlib.RADataFulfilled) {

	// --------------------------
	// Check for date section
	// --------------------------
	dates := a.Dates
	if !(time.Time(dates.RentStart).IsZero() || time.Time(dates.RentStop).IsZero() || time.Time(dates.AgreementStart).IsZero() || time.Time(dates.AgreementStop).IsZero() || time.Time(dates.PossessionStart).IsZero() || time.Time(dates.PossessionStop).IsZero()) {
		d.Dates = true
	}

	// --------------------------
	// Check for people section
	// --------------------------
	renterCount := 0
	for _, people := range a.People {
		if people.IsRenter {
			renterCount++
		}
	}
	if renterCount > 0 {
		d.People = true
	}

	// --------------------------
	// Check for pet section
	// --------------------------
	if !a.Meta.HavePets {
		d.Pets = true
	} else {
		if len(a.Pets) > 0 {
			d.Pets = true
		} else {
			d.Pets = false
		}
	}

	// --------------------------
	// Check for vehicle section
	// --------------------------
	if !a.Meta.HaveVehicles {
		d.Vehicles = true
	} else {
		if len(a.Vehicles) > 0 {
			d.Vehicles = true
		} else {
			d.Vehicles = false
		}
	}

	// ---------------------------
	// Check for rentables section
	// ---------------------------
	// There must be at least one parent rentable
	for _, rentable := range a.Rentables {
		if rentable.RTFLAGS&(1<<1) == 0 {
			d.Rentables = true
			break
		}
	}

	// -----------------------------
	// Check for parentchild section
	// -----------------------------
	// ==============================================================//
	//  ****************** VALIDATION SCENARIOS *********************//
	// ==============================================================//
	// 1.   If there are no child rentables then it is fine
	// Ex:  People only want to stay at apartment. They might
	//      no require child rentables like washing machine,
	//      car parking space etc...
	//
	// 2.   There must be at least one parent rentables in rentables
	//      section. People come to stay at rooms/apartments, so it
	//      doesn't make sense of not having any parent rentables.
	//
	// 3.   If any child rentables listed in rentables section then
	//      it must be associated with parent rentables.
	// Ex:  Washing machine (a child rentable) must be associated to
	//      an apartment(a parent rentable) where the people are
	//      living.
	// ==============================================================//
	for _, pc := range a.ParentChild {
		if !(pc.CRID > 0 && pc.PRID > 0) {
			d.ParentChild = false
			break
		} else {
			d.ParentChild = true
		}
	}

	// If there are no child rentables
	if len(a.ParentChild) == 0 {
		d.ParentChild = true
	}

	// -----------------------------
	// Check for tie section
	// -----------------------------
	// There must be at least one person
	if len(a.Tie.People) > 0 {
		d.Tie = true
	}

}
