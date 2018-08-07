package rlib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// RAFlow etc.. all are list of all flows exist in the system
const (
	RAFlow string = "RA"
)

// RAFlowPartType is type of rental agreement flow part
type RAFlowPartType int

// DatesRAFlowPart etc. all are constants for rental agreement flow part
const (
	DatesRAFlowPart RAFlowPartType = 1 + iota // must start from 1
	PeopleRAFlowPart
	PetsRAFlowPart
	VehiclesRAFlowPart
	RentablesRAFlowPart
	ParentChildRAFlowPart
	TieRAFlowPart
)

// IsValid checks the validity of RAFlowPartType raftp
func (raftp RAFlowPartType) IsValid() bool {
	if raftp < DatesRAFlowPart || raftp > RentablesRAFlowPart {
		return false
	}

	return true
}

// String representation of RAFlowPartType
func (raftp RAFlowPartType) String() string {
	names := [...]string{
		"Agreement Dates",
		"People with background info",
		"Pets",
		"Vehicles",
		"Rentables with fees",
		"Parent/Child",
		"Tie",
	}

	// if not valid then return unknown
	if !(raftp.IsValid()) {
		return "Unknown RA FlowPart"
	}

	return names[raftp-1]
}

// RAFlowPartsMap parts of a rental agreement flow
var RAFlowPartsMap = Str2Int64Map{
	"dates":       int64(DatesRAFlowPart),
	"people":      int64(PeopleRAFlowPart),
	"pets":        int64(PetsRAFlowPart),
	"vehicles":    int64(VehiclesRAFlowPart),
	"rentables":   int64(RentablesRAFlowPart),
	"parentchild": int64(ParentChildRAFlowPart),
	"tie":         int64(TieRAFlowPart),
}

// RAFlowJSONData holds the struct for all the parts being involed in rental agreement flow
type RAFlowJSONData struct {
	Dates       RADatesFlowData         `json:"dates"`
	People      []RAPeopleFlowData      `json:"people"`
	Pets        []RAPetsFlowData        `json:"pets"`
	Vehicles    []RAVehiclesFlowData    `json:"vehicles"`
	Rentables   []RARentablesFlowData   `json:"rentables"`
	ParentChild []RAParentChildFlowData `json:"parentchild"`
	Tie         RATieFlowData           `json:"tie"`
	Meta        RAFlowMetaInfo          `json:"meta"`
}

// RAFlowMetaInfo holds meta info about a rental agreement flow data
type RAFlowMetaInfo struct {
	RAID                   int64 // 0 = it's new, >0 = existing one
	LastTMPPETID           int64
	LastTMPVID             int64
	LastTMPTCID            int64
	LastTMPASMID           int64
	HavePets               bool
	HaveVehicles           bool
	RAFLAGS                uint64
	ApplicationReadyUID    int64
	ApplicationReadyName   string
	ApplicationReadyDate   JSONDateTime
	Approver1              int64
	Approver1Name          string
	DecisionDate1          JSONDateTime
	DeclineReason1         int64
	Approver2              int64
	Approver2Name          string
	DecisionDate2          JSONDateTime
	DeclineReason2         int64
	MoveInUID              int64
	MoveInName             string
	MoveInDate             JSONDateTime
	ActiveUID              int64
	ActiveName             string
	ActiveDate             JSONDateTime
	TerminatorUID          int64
	TerminatorName         string
	TerminationDate        JSONDateTime
	LeaseTerminationReason int64
	DocumentDate           JSONDateTime
	NoticeToMoveUID        int64
	NoticeToMoveName       string
	NoticeToMoveDate       JSONDateTime
	NoticeToMoveReported   JSONDateTime
}

// RADatesFlowData contains data in the dates part of RA flow
type RADatesFlowData struct {
	BID             int64    `validate:"number,min=0"`
	AgreementStart  JSONDate `validate:"date"` // TermStart
	AgreementStop   JSONDate `validate:"date"` // TermStop
	RentStart       JSONDate `validate:"date"`
	RentStop        JSONDate `validate:"date"`
	PossessionStart JSONDate `validate:"date"`
	PossessionStop  JSONDate `validate:"date"`
	CSAgent         int64    `validate:"number,min=0"` // TODO(Steve/Sudip/Akshay): Bind webservice call. TODO(Akshay): Move it to the Meta structure
}

// RAPeopleFlowData contains data in the background-info part of RA flow
type RAPeopleFlowData struct {
	TMPTCID int64 `validate:"number,min=1"`
	BID     int64 `validate:"number,min=0"`
	TCID    int64 `validate:"number,min=0"`

	// Role
	IsRenter    bool `validate:"-"`
	IsOccupant  bool `validate:"-"`
	IsGuarantor bool `validate:"-"`

	// ---------- Basic Info -----------
	FirstName      string `validate:"string,min=1,max=100,omitempty"` // It handles in business logic if isCompany flag is false
	MiddleName     string `validate:"string,min=1,max=100,omitempty"`
	LastName       string `validate:"string,min=1,max=100,omitempty"` // It handles in business logic if isCompany flag is false
	PreferredName  string `validate:"string,min=1,max=100,omitempty"`
	IsCompany      bool   `validate:"-"`
	CompanyName    string `validate:"string,min=1,max=100,omitempty"` // It is required when IsCompany flag is true. It'll be checked in bizlogic validation.
	PrimaryEmail   string `validate:"email"`
	SecondaryEmail string `validate:"email,omitempty"`
	WorkPhone      string `validate:"string,min=1,max=100,omitempty"` // Either Workphone or CellPhone is compulsory. It'll be checked in bizlogic validation
	CellPhone      string `validate:"string,min=1,max=100,omitempty"` // Either Workphone or CellPhone is compulsory. It'll be checked in bizlogic validation
	Address        string `validate:"string,min=1,max=100"`
	Address2       string `validate:"string,min=0,max=100,omitempty"`
	City           string `validate:"string,min=1,max=100"`
	State          string `validate:"string,min=1,max=25"`
	PostalCode     string `validate:"string,min=1,max=100"`
	Country        string `validate:"string,min=1,max=100"`
	Website        string `validate:"string,min=1,max=100,omitempty"`
	Comment        string `validate:"string,min=1,max=2048,omitempty"`

	// ---------- Prospect -----------
	CompanyAddress    string `validate:"string,min=1,max=100,omitempty"`
	CompanyCity       string `validate:"string,min=1,max=100,omitempty"`
	CompanyState      string `validate:"string,min=1,max=100,omitempty"`
	CompanyPostalCode string `validate:"string,min=1,max=100,omitempty"`
	CompanyEmail      string `validate:"email,omitempty"`
	CompanyPhone      string `validate:"string,min=1,max=100,omitempty"`
	Occupation        string `validate:"string,min=1,max=100"`

	// Current Address information
	CurrentAddress           string `validate:"string,min=1,max=100,omitempty"`
	CurrentLandLordName      string `validate:"string,min=1,max=100,omitempty"`
	CurrentLandLordPhoneNo   string `validate:"string,min=1,max=100,omitempty"`
	CurrentLengthOfResidency string `validate:"string,min=1,max=100,omitempty"`
	CurrentReasonForMoving   int64  `validate:"number,min=1,omitempty"` // Reason for moving

	// Prior Address information
	PriorAddress           string `validate:"string,min=1,max=100,omitempty"`
	PriorLandLordName      string `validate:"string,min=1,max=100,omitempty"`
	PriorLandLordPhoneNo   string `validate:"string,min=1,max=100,omitempty"`
	PriorLengthOfResidency string `validate:"string,min=1,max=100,omitempty"`
	PriorReasonForMoving   int64  `validate:"number,min=1,omitempty"` // Reason for moving

	// Have you ever been
	Evicted          bool   `validate:"-"` // Evicted
	EvictedDes       string `validate:"string,min=1,max=2048,omitempty"`
	Convicted        bool   `validate:"-"` // Arrested or convicted of a Convicted
	ConvictedDes     string `validate:"string,min=1,max=2048,omitempty"`
	Bankruptcy       bool   `validate:"-"` // Declared Bankruptcy
	BankruptcyDes    string `validate:"string,min=1,max=2048,omitempty"`
	OtherPreferences string `validate:"string,min=1,max=1024,omitempty"`
	//FollowUpDate             JSONDate
	//CommissionableThirdParty string
	SpecialNeeds string `validate:"string,min=1,max=1024,omitempty"` // In an effort to accommodate you, please advise us of any special needs
	// It'll be none. If there is no special needs
	ThirdPartySource string `validate:"string,min=1,max=100,omitempty"`

	// ---------- Payor -----------
	CreditLimit         float64 `validate:"number:float,min=0.00,omitempty"`
	TaxpayorID          string  `validate:"string,min=1,max=25,omitempty"`   // It requires when transanctant is renter or gurantor. It handles via business logic
	GrossIncome         float64 `validate:"number:float,min=0.00,omitempty"` // When role is set to renter or guarantor than it is compulsory. It'll be check via bizlogic.
	DriversLicense      string  `validate:"string,min=1,max=128"`
	EligibleFuturePayor bool    `validate:"-"`

	// ---------- User -----------
	Points      int64    `validate:"number,min=1,omitempty"`
	DateofBirth JSONDate `validate:"-"`
	// Emergency contact information
	EmergencyContactName      string `validate:"string,min=1,max=100"`
	EmergencyContactAddress   string `validate:"string,min=1,max=100"`
	EmergencyContactTelephone string `validate:"string,min=1,max=100"`
	EmergencyContactEmail     string `validate:"email"`
	AlternateEmailAddress     string `validate:"string,min=1,max=100,omitempty"`
	EligibleFutureUser        bool   `validate:"-"`
	Industry                  int64  `validate:"number,min=0,omitempty"`
	SourceSLSID               int64  `validate:"number,min=1"` // It is compulsory when role is set to renter or user. It'll be check via bizlogic.
}

// RAPetsFlowData contains data in the pets part of RA flow
type RAPetsFlowData struct {
	TMPPETID int64        `validate:"number,min=1"`
	BID      int64        `validate:"number,min=0"`
	PETID    int64        `validate:"number,min=0"`
	TMPTCID  int64        `validate:"number,min=0"`
	Name     string       `validate:"string,min=1,max=100"`
	Type     string       `validate:"string,min=1,max=100"`
	Breed    string       `validate:"string,min=1,max=100"`
	Color    string       `validate:"string,min=1,max=100"`
	Weight   float64      `validate:"number:float,min=0.0"`
	DtStart  JSONDate     `validate:"date"`
	DtStop   JSONDate     `validate:"date"`
	Fees     []RAFeesData `validate:"-"`
}

// RAVehiclesFlowData contains data in the vehicles part of RA flow
type RAVehiclesFlowData struct {
	TMPVID              int64        `validate:"number,min=1"`
	BID                 int64        `validate:"number,min=0"`
	VID                 int64        `validate:"number,min=0"`
	TMPTCID             int64        `validate:"number,min=0"`
	VIN                 string       `validate:"string,min=1,omitempty"`
	VehicleType         string       `validate:"string,min=1,max=80"`
	VehicleMake         string       `validate:"string,min=1,max=80"`
	VehicleModel        string       `validate:"string,min=1,max=80"`
	VehicleColor        string       `validate:"string,min=1,max=80"`
	VehicleYear         int64        `validate:"number,min=1900,max=2150"`
	LicensePlateState   string       `validate:"string,min=1,max=80"`
	LicensePlateNumber  string       `validate:"string,min=1,max=80"`
	ParkingPermitNumber string       `validate:"string,min=1,max=80,omitempty"`
	DtStart             JSONDate     `validate:"date"`
	DtStop              JSONDate     `validate:"date"`
	Fees                []RAFeesData `validate:"-"`
}

// RARentablesFlowData contains data in the rentables part of RA flow
type RARentablesFlowData struct {
	BID             int64 `validate:"number,min=0"`
	RID             int64 `validate:"number,min=0"`
	RTID            int64 `validate:"number,min=0"`
	RTFLAGS         uint64
	RentableName    string  `validate:"string,min=1,max=100"`
	RentCycle       int64   `validate:"number,min=0"`
	AtSigningPreTax float64 `validate:"number:float,min=0.00"`
	SalesTax        float64 `validate:"number:float,min=0.00"`
	// SalesTaxAmt    float64 // FUTURE RELEASE
	TransOccTax float64 `validate:"number:float,min=0.00"`
	// TransOccAmt    float64 // FUTURE RELEASE
	Fees []RAFeesData `validate:"-"`
}

// RAFeesData struct used for pet, vehicles, rentable fees
type RAFeesData struct {
	TMPASMID        int64    `validate:"number,min=1"` // unique ID to manage fees uniquely across all fees in raflow json data
	ASMID           int64    `validate:"number,min=0"` // the permanent table assessment id if it is an existing RAID
	ARID            int64    `validate:"number,min=1"`
	ARName          string   `validate:"string,min=1,max=100"`
	ContractAmount  float64  `validate:"number:float,min=0.00"`
	RentCycle       int64    `validate:"number,min=0"`
	Start           JSONDate `validate:"date"`
	Stop            JSONDate `validate:"date"`
	AtSigningPreTax float64  `validate:"number:float,min=0.00"`
	SalesTax        float64  `validate:"number:float,min=0.00"`
	// SalesTaxAmt     float64       // FUTURE RELEASE
	TransOccTax float64 `validate:"number:float,min=0.00"`
	// TransOccAmt float64 // FUTURE RELEASE
	Comment string `validate:"string,min=1,max=256,omitempty"`
}

// RAParentChildFlowData contains data in the Parent/Child part of RA flow
type RAParentChildFlowData struct {
	BID  int64 `validate:"number,min=0"`
	PRID int64 `validate:"number,min=0"` // parent rentable ID
	CRID int64 `validate:"number,min=0"` // child rentable ID
}

// RATieFlowData contains data in the tie part of RA flow
type RATieFlowData struct {
	People []RATiePeopleData `json:"people"`
}

// RATiePeopleData holds data from tie section for a payor to a rentable
type RATiePeopleData struct {
	BID     int64 `validate:"number,min=0"`
	PRID    int64 `validate:"number,min=0"`
	TMPTCID int64 `validate:"number,min=0"` // user's temp json record reference id
}

// UpdateRAFlowJSON updates json data based on requested
// flowPart (string)
func UpdateRAFlowJSON(ctx context.Context, BID int64, dataToUpdate json.RawMessage, flowPart string, flow *Flow) (err error) {
	const funcname = "UpdateRAFlowJSON"
	var (
		raFlowData          RAFlowJSONData
		modFlowPartData     = []byte(nil)
		possessDatesChanged bool
		rentDatesChanged    bool
	)
	fmt.Printf("Entered in %s\n", funcname)

	// CHECK REQUESTED FLOW PART IS VALID
	RAFlowPart, OK := RAFlowPartsMap[flowPart]
	if !OK {
		err = fmt.Errorf("RAFlow part key: %s with flowID: %d is not valid, Error: %s",
			flowPart, flow.FlowID, err.Error())
		return
	}

	// ----- GET THE RAFLOW DATA FROM FLOW ------ //
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		return
	}
	meta := raFlowData.Meta

	// JSON MARSHAL WITH ADDRESS
	// REF: https://stackoverflow.com/questions/21390979/custom-marshaljson-never-gets-called-in-go

	// BYTES DATA BLANK OR NULL VALUE CHECK
	isBlankJSONData := bytes.Equal([]byte(dataToUpdate), []byte(``)) ||
		bytes.Equal([]byte(dataToUpdate), []byte(`null`)) ||
		bytes.Equal([]byte(dataToUpdate), []byte(nil))

	//-------------------------------------------------------
	// FLOW PARTS SWITCH CASES
	//-------------------------------------------------------
	switch RAFlowPartType(RAFlowPart) {
	case DatesRAFlowPart:
		a := RADatesFlowData{}

		// IF DATA IS PROVIDED THEN
		if !(isBlankJSONData) {
			err = json.Unmarshal(dataToUpdate, &a)
			if err != nil {
				return
			}

			// ----- POSSESSION DATES CHANGED CHECK ----- //
			if !((time.Time)(raFlowData.Dates.PossessionStart).Equal((time.Time)(a.PossessionStart)) &&
				(time.Time)(raFlowData.Dates.PossessionStop).Equal((time.Time)(a.PossessionStop))) {
				possessDatesChanged = true
			}

			// ----- RENT DATES CHANGED CHECK ----- //
			if !((time.Time)(raFlowData.Dates.RentStart).Equal((time.Time)(a.RentStart)) &&
				(time.Time)(raFlowData.Dates.RentStop).Equal((time.Time)(a.RentStop))) {
				rentDatesChanged = true
			}

		} else {
			// IF DATA IS BLANK OR NULL THEN INITIAZE WITH IT SOME DEFAULTS
			currentDateTime := time.Now()
			nextYearDateTime := currentDateTime.AddDate(1, 0, 0)

			a.BID = BID
			a.RentStart = JSONDate(currentDateTime)
			a.RentStop = JSONDate(nextYearDateTime)
			a.AgreementStart = JSONDate(currentDateTime)
			a.AgreementStop = JSONDate(nextYearDateTime)
			a.PossessionStart = JSONDate(currentDateTime)
			a.PossessionStop = JSONDate(nextYearDateTime)
		}

		// MODIFIED PART DATA
		modFlowPartData, err = json.Marshal(&a)
		if err != nil {
			return
		}

	case PeopleRAFlowPart:
		a := []RAPeopleFlowData{}

		// IF DATA IS PROVIDED THEN
		if !(isBlankJSONData) {
			err = json.Unmarshal(dataToUpdate, &a)
			if err != nil {
				return
			}

			// IF NOT TMPTCID THEN ASSIGN IT
			for i := range a {
				if a[i].TMPTCID == 0 {
					meta.LastTMPTCID++
					a[i].TMPTCID = meta.LastTMPTCID
				}

				// if Special needs are none, then it should indicate none
				if a[i].SpecialNeeds == "" {
					a[i].SpecialNeeds = "None"
				}
			}

		}

		// MODIFIED PART DATA
		modFlowPartData, err = json.Marshal(&a)
		if err != nil {
			return
		}

	case PetsRAFlowPart:
		a := []RAPetsFlowData{}

		// IF DATA IS PROVIDED THEN
		if !(isBlankJSONData) {
			err = json.Unmarshal(dataToUpdate, &a)
			if err != nil {
				return
			}

			// IF NOT TMPPETID THEN ASSIGN IT
			for i := range a {

				// IF NOT FEES LIST THEN
				if len(a[i].Fees) == 0 {
					a[i].Fees = []RAFeesData{}
				}

				if a[i].TMPPETID == 0 {
					meta.LastTMPPETID++
					a[i].TMPPETID = meta.LastTMPPETID

					// IF NOT TMPASMID IN EACH FEE THEN
					for j := range a[i].Fees {
						if a[i].Fees[j].TMPASMID == 0 {
							meta.LastTMPASMID++
							a[i].Fees[j].TMPASMID = meta.LastTMPASMID
						}
					}
				}
			}

			// HAVEPETS  - BASED ON PET LIST LENGTH
			meta.HavePets = len(a) > 0
		}

		// MODIFIED PART DATA
		modFlowPartData, err = json.Marshal(&a)
		if err != nil {
			return
		}

	case VehiclesRAFlowPart:
		a := []RAVehiclesFlowData{}

		// IF DATA IS PROVIDED THEN
		if !(isBlankJSONData) {
			err = json.Unmarshal(dataToUpdate, &a)
			if err != nil {
				return
			}

			// IF NOT TMPPETID THEN
			for i := range a {

				// IF NOT FEES ASSOCIATED
				if len(a[i].Fees) == 0 {
					a[i].Fees = []RAFeesData{}
				}

				if a[i].TMPVID == 0 {
					meta.LastTMPVID++
					a[i].TMPVID = meta.LastTMPVID

					// IF NOT TMPASMID IN EACH FEE
					for j := range a[i].Fees {
						if a[i].Fees[j].TMPASMID == 0 {
							meta.LastTMPASMID++
							a[i].Fees[j].TMPASMID = meta.LastTMPASMID
						}
					}
				}
			}

			// HAVE VEHICLES - BASED ON VEHICLE LIST LENGTH
			meta.HaveVehicles = len(a) > 0
		}

		// MODIFIED PART DATA
		modFlowPartData, err = json.Marshal(&a)
		if err != nil {
			return
		}

	case RentablesRAFlowPart:
		a := []RARentablesFlowData{}

		// IF DATA IS PROVIDED THEN
		if !(isBlankJSONData) {
			err = json.Unmarshal(dataToUpdate, &a)
			if err != nil {
				return
			}

			// FEES TMPASMID
			for i := range a {

				// IF NOT FEES ASSOCIATED THEN
				if len(a[i].Fees) == 0 {
					a[i].Fees = []RAFeesData{}
				}

				// IF NOT TMPASMID IN EACH FEE THEN
				for j := range a[i].Fees {
					if a[i].Fees[j].TMPASMID == 0 {
						meta.LastTMPASMID++
						a[i].Fees[j].TMPASMID = meta.LastTMPASMID
					}
				}

			}
		}

		// MODIFIED PART DATA
		modFlowPartData, err = json.Marshal(&a)
		if err != nil {
			return
		}

	case ParentChildRAFlowPart:
		a := []RAParentChildFlowData{}

		// IF DATA IS PROVIDED THEN
		if !(isBlankJSONData) {
			err = json.Unmarshal(dataToUpdate, &a)
			if err != nil {
				return
			}
		}

		// MODIFIED PART DATA
		modFlowPartData, err = json.Marshal(&a)
		if err != nil {
			return
		}

	case TieRAFlowPart:
		a := RATieFlowData{}

		// IF DATA IS PROVIDED THEN
		if !(isBlankJSONData) {
			err = json.Unmarshal(dataToUpdate, &a)
			if err != nil {
				return
			}

			// IF NOT PEOPLE IN TIE THEN
			if len(a.People) == 0 {
				a.People = []RATiePeopleData{}
			}
		}

		// MODIFIED PART DATA
		modFlowPartData, err = json.Marshal(&a)
		if err != nil {
			return
		}

	default:
		err = fmt.Errorf("unrecognized part type in RA flow: %s", flowPart)
		return
	}

	// NOW UPDATE THE JSON FOR GIVEN FLOW PART
	err = UpdateFlowData(ctx, flowPart, modFlowPartData, flow)
	if err != nil {
		return
	}

	// IF MODIFIED META IS NOT EQUAL TO FLOW META THEN ONLY UPDATE META JSON
	if !reflect.DeepEqual(meta, raFlowData.Meta) {
		var modMetaInfo []byte
		modMetaInfo, err = json.Marshal(&meta)
		if err != nil {
			return
		}
		err = UpdateFlowData(ctx, "meta", modMetaInfo, flow)
		if err != nil {
			return
		}
	}

	// IF POSSESSION DATES WERE CHANGED THEN TAKE NECESSARY ACTION
	if possessDatesChanged {
		err = possessDateChangeRAFlowUpdates(ctx, flow.FlowID)
		if err != nil {
			return
		}
	}

	// IF RENT DATES WERE CHANGED THEN TAKE NECESSARY ACTION
	if rentDatesChanged {
		err = rentDateChangeRAFlowUpdates(ctx, flow.FlowID)
		if err != nil {
			return
		}
	}

	return
}

// possessDateChangeRAFlowUpdates updates raflow json with required
// modification if possession dates are changed
func possessDateChangeRAFlowUpdates(ctx context.Context, flowID int64) (err error) {

	// ----- GET THE UPDATED FLOW ----- //
	var flow Flow
	flow, err = GetFlow(ctx, flowID)
	if err != nil {
		return
	}
	if flow.FlowID == 0 {
		err = fmt.Errorf("given flow with ID (%d) doesn't exist", flowID)
		return
	}

	// ----- GET THE RAFLOW DATA FROM FLOW ------ //
	var raFlowData RAFlowJSONData
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		return
	}

	// ----- UPDATE PETS DTSTART/DTSTOP WITH NEW DATES ----- //
	pets := raFlowData.Pets
	for i := range pets {
		pets[i].DtStart = raFlowData.Dates.PossessionStart
		pets[i].DtStop = raFlowData.Dates.PossessionStop
	}

	var modPetsData []byte
	modPetsData, err = json.Marshal(&pets)
	if err != nil {
		return
	}
	err = UpdateFlowData(ctx, "pets", modPetsData, &flow)
	if err != nil {
		return
	}

	// ----- UPDATE PETS DTSTART/DTSTOP WITH NEW DATES ----- //
	vehicles := raFlowData.Vehicles
	for i := range vehicles {
		vehicles[i].DtStart = raFlowData.Dates.PossessionStart
		vehicles[i].DtStop = raFlowData.Dates.PossessionStop
	}

	var modVehiclesData []byte
	modVehiclesData, err = json.Marshal(&vehicles)
	if err != nil {
		return
	}
	err = UpdateFlowData(ctx, "vehicles", modVehiclesData, &flow)
	if err != nil {
		return
	}

	return
}

// rentDateChangeRAFlowUpdates updates raflow json with required
// modification if rent dates are changed
func rentDateChangeRAFlowUpdates(ctx context.Context, flowID int64) (err error) {

	// ----- GET THE UPDATED FLOW ----- //
	var flow Flow
	flow, err = GetFlow(ctx, flowID)
	if err != nil {
		return
	}
	if flow.FlowID == 0 {
		err = fmt.Errorf("given flow with ID (%d) doesn't exist", flowID)
		return
	}

	// ----- GET THE RAFLOW DATA FROM FLOW ------ //
	var raFlowData RAFlowJSONData
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		return
	}

	// IMPLEMENTATION ERROR //
	// TODO(Sudip): re-calculate fees for pets, vehicles, rentables
	return /*fmt.Errorf("implementation error")*/
}

// InsertInitialRAFlow writes a bunch of flow's sections record for a particular RA
func InsertInitialRAFlow(ctx context.Context, BID, UID int64) (int64, error) {

	var (
		flowID int64
		err    error
	)

	// current date and next year date
	currentDateTime := time.Now()
	nextYearDateTime := currentDateTime.AddDate(1, 0, 0)

	// rental agreement flow data
	initialRAFlow := RAFlowJSONData{
		Dates: RADatesFlowData{
			BID:             BID,
			RentStart:       JSONDate(currentDateTime),
			RentStop:        JSONDate(nextYearDateTime),
			AgreementStart:  JSONDate(currentDateTime),
			AgreementStop:   JSONDate(nextYearDateTime),
			PossessionStart: JSONDate(currentDateTime),
			PossessionStop:  JSONDate(nextYearDateTime),
			CSAgent:         UID, // CS Agent value to the UID of the logged in user
		},
		People:      []RAPeopleFlowData{},
		Pets:        []RAPetsFlowData{},
		Vehicles:    []RAVehiclesFlowData{},
		Rentables:   []RARentablesFlowData{},
		ParentChild: []RAParentChildFlowData{},
		Tie: RATieFlowData{
			People: []RATiePeopleData{},
		},
	}

	// get json marshelled byte data for above struct
	raflowJSONData, err := json.Marshal(&initialRAFlow)
	if err != nil {
		Ulog("Error while marshalling json data of initialRAFlow: %s\n", err.Error())
		return flowID, err
	}

	// initial Flow struct
	Console("New Flow\n")
	a := Flow{
		BID:       BID,
		FlowID:    0, // it's new flowID,
		UserRefNo: GenerateUserRefNo(),
		FlowType:  RAFlow,
		Data:      raflowJSONData,
		CreateBy:  UID,
		LastModBy: UID,
	}

	Console("New flow UserRefNo = %s\n", a.UserRefNo)

	// insert new flow
	flowID, err = InsertFlow(ctx, &a)
	if err != nil {
		Ulog("Error while inserting Flow: %s\n", err.Error())
		return flowID, err
	}

	return flowID, err
}

// RAFlowDataDiff returns true/false if there is any data
// difference between raflow json data and permanent table
// data
func RAFlowDataDiff(ctx context.Context, RAID int64) (isDiff bool, err error) {
	const funcname = "RAFlowDataDiff"
	var (
		ra                      RentalAgreement
		flow                    Flow
		flowData, permanentData RAFlowJSONData
	)

	// --------------------------------------------------------
	// GET DATA FROM TEMP FLOW table
	// --------------------------------------------------------
	flow, err = GetFlowForRAID(ctx, "RA", RAID)
	if err != nil {
		return
	}

	err = json.Unmarshal(flow.Data, &flowData)
	if err != nil {
		return
	}

	// --------------------------------------------------------
	// GET PERMANENT RA and GET FLOWDATA FOR IT
	// --------------------------------------------------------
	ra, err = GetRentalAgreement(ctx, RAID)
	if err != nil {
		return
	}

	// convert permanent ra to flow data and get it
	permanentData, err = ConvertRA2Flow(ctx, &ra)
	if err != nil {
		return
	}

	// --------------------------------------------------------
	// WE DON'T NEED meta DATA TO DIFF BETWEEN BOTH JSON DATA
	// --------------------------------------------------------
	permanentData.Meta, flowData.Meta = RAFlowMetaInfo{}, RAFlowMetaInfo{}

	// NOW TAKE THE DIFF USING REFLECT
	sameData := reflect.DeepEqual(permanentData, flowData) // returns true, if both are equal
	isDiff = !sameData

	return
}

// ConvertRA2Flow does all the heavy lifting to convert existing
// rental agreement data to raflow data
//
// INPUTS:
//     ctx    database context for transactions
//     ra     the rental agreement to move into a flow
//
// RETURNS:
//     the RAFlowJSONData
//     any error encountered
//-------------------------------------------------------------------------
func ConvertRA2Flow(ctx context.Context, ra *RentalAgreement) (RAFlowJSONData, error) {
	const funcname = "ConvertRA2Flow"

	//-------------------------------------------------------------
	// This is the datastructure we need to fill out and save...
	//-------------------------------------------------------------
	ApplicationReadyName, _ := GetDirectoryPerson(ctx, ra.ApplicationReadyUID)
	MoveInName, _ := GetDirectoryPerson(ctx, ra.MoveInUID)
	ActiveName, _ := GetDirectoryPerson(ctx, ra.ActiveUID)
	Approver1Name, _ := GetDirectoryPerson(ctx, ra.Approver1)
	Approver2Name, _ := GetDirectoryPerson(ctx, ra.Approver2)
	TerminatorName, _ := GetDirectoryPerson(ctx, ra.TerminatorUID)
	NoticeToMoveName, _ := GetDirectoryPerson(ctx, ra.NoticeToMoveUID)

	var raf = RAFlowJSONData{
		Dates: RADatesFlowData{
			BID:             ra.BID,
			RentStart:       JSONDate(ra.RentStart),
			RentStop:        JSONDate(ra.RentStop),
			AgreementStart:  JSONDate(ra.AgreementStart),
			AgreementStop:   JSONDate(ra.AgreementStop),
			PossessionStart: JSONDate(ra.PossessionStart),
			PossessionStop:  JSONDate(ra.PossessionStop),
			CSAgent:         ra.CSAgent,
		},
		People:      []RAPeopleFlowData{},
		Pets:        []RAPetsFlowData{},
		Vehicles:    []RAVehiclesFlowData{},
		Rentables:   []RARentablesFlowData{},
		ParentChild: []RAParentChildFlowData{},
		Tie: RATieFlowData{
			People: []RATiePeopleData{},
		},
		Meta: RAFlowMetaInfo{
			RAID:                   ra.RAID,
			RAFLAGS:                ra.FLAGS,
			ApplicationReadyUID:    ra.ApplicationReadyUID,
			ApplicationReadyName:   ApplicationReadyName.DisplayName(),
			ApplicationReadyDate:   JSONDateTime(ra.ApplicationReadyDate),
			Approver1:              ra.Approver1,
			Approver1Name:          Approver1Name.DisplayName(),
			DecisionDate1:          JSONDateTime(ra.DecisionDate1),
			DeclineReason1:         ra.DeclineReason1,
			Approver2:              ra.Approver2,
			Approver2Name:          Approver2Name.DisplayName(),
			DecisionDate2:          JSONDateTime(ra.DecisionDate2),
			DeclineReason2:         ra.DeclineReason2,
			MoveInUID:              ra.MoveInUID,
			MoveInName:             MoveInName.DisplayName(),
			MoveInDate:             JSONDateTime(ra.MoveInDate),
			ActiveUID:              ra.ActiveUID,
			ActiveName:             ActiveName.DisplayName(),
			ActiveDate:             JSONDateTime(ra.ActiveDate),
			TerminatorUID:          ra.TerminatorUID,
			TerminatorName:         TerminatorName.DisplayName(),
			TerminationDate:        JSONDateTime(ra.TerminationDate),
			LeaseTerminationReason: ra.LeaseTerminationReason,
			DocumentDate:           JSONDateTime(ra.DocumentDate),
			NoticeToMoveUID:        ra.NoticeToMoveUID,
			NoticeToMoveName:       NoticeToMoveName.DisplayName(),
			NoticeToMoveDate:       JSONDateTime(ra.NoticeToMoveDate),
			NoticeToMoveReported:   JSONDateTime(ra.NoticeToMoveReported),
		},
	}

	//-------------------------------------------------------------------------
	// Add Users...
	//
	// Note: we need to add users before payors in order to
	// ensure that all the pets and vehicles are loaded.  This is because
	// of two behaviors of the code.  First, pets and vehicles are loaded only
	// when the person is loaded with the User (occupant) flag set.  Second,
	// a person is not added twice, and if you load a Payor first -- the pets
	// and vehicles will NOT be loaded then when you call it the second time
	// to load the same transactant as a User the code will see that the
	// transactant has already been loaded and return without doing anything
	// other than setting the Payor (renter) flag.
	//-------------------------------------------------------------------------
	n, err := GetAllRentalAgreementRentables(ctx, ra.RAID)
	if err != nil {
		return raf, nil
	}
	for j := 0; j < len(n); j++ {
		rulist, err := GetRentableUsersInRange(ctx, n[j].RID, &ra.AgreementStart, &ra.AgreementStop)
		if err != nil {
			return raf, nil
		}
		for k := 0; k < len(rulist); k++ {
			addRAPtoFlow(ctx, rulist[k].TCID, n[j].RID, &raf, true, false, true)
		}
	}

	//-------------------------------------------------------------------------
	// Add Payors...
	//-------------------------------------------------------------------------
	m, err := GetRentalAgreementPayorsInRange(ctx, ra.RAID, &ra.AgreementStart, &ra.AgreementStop)
	if err != nil {
		return raf, nil
	}
	for i := 0; i < len(m); i++ {
		if err = addRAPtoFlow(ctx, m[i].TCID, 0 /*no RID here*/, &raf, true /*check dups*/, true /*renter*/, false); err != nil {
			return raf, nil
		}
	}

	//-------------------------------------------------------------------------
	// Add Rentables
	//-------------------------------------------------------------------------
	now := time.Now()
	o, err := GetRentalAgreementRentables(ctx, ra.RAID, &ra.AgreementStart, &ra.AgreementStop)
	if err != nil {
		return raf, nil
	}
	for i := 0; i < len(o); i++ {
		rnt, err := GetRentable(ctx, o[i].RID)
		if err != nil {
			return raf, nil
		}
		rtr, err := GetRentableTypeRefForDate(ctx, o[i].RID, &now)
		if err != nil {
			return raf, nil
		}
		var rt RentableType
		if err = GetRentableType(ctx, rtr.RTID, &rt); err != nil {
			return raf, nil
		}
		var rfd = RARentablesFlowData{
			BID:          o[i].BID,
			RID:          o[i].RID,
			RTID:         rtr.RTID,
			RTFLAGS:      rt.FLAGS,
			RentableName: rnt.RentableName,
			RentCycle:    rt.RentCycle,
			Fees:         []RAFeesData{},
		}

		//---------------------------------------------------------
		// Add the assessments associated with the Rentable...
		// For this we want to load all 1-time fees and all
		// recurring fees.
		//---------------------------------------------------------
		asms, err := GetAssessmentsByRAIDRID(ctx, rfd.BID, ra.RAID, rfd.RID)
		if err != nil {
			return raf, nil
		}
		for j := 0; j < len(asms); j++ {

			//----------------------------------------------------------
			// Get the account rule for this assessment...
			//----------------------------------------------------------
			ar, err := GetAR(ctx, asms[j].ARID)
			if err != nil {
				return raf, nil
			}

			//----------------------------------------------------------
			// Handle Rentable Fees that are NOT Pet or Vehicle related
			//----------------------------------------------------------
			if ar.FLAGS&(128|256) == 0 {
				raf.Meta.LastTMPASMID++
				var fee = RAFeesData{
					TMPASMID:       raf.Meta.LastTMPASMID,
					ASMID:          asms[j].ASMID,
					ARID:           asms[j].ARID,
					ARName:         ar.Name,
					ContractAmount: asms[j].Amount,
					RentCycle:      asms[j].RentCycle,
					Start:          JSONDate(asms[j].Start),
					Stop:           JSONDate(asms[j].Stop),
					Comment:        asms[j].Comment,
				}
				rfd.Fees = append(rfd.Fees, fee)
			}

			//----------------------------------------------------------
			// Handle PET Fees
			//----------------------------------------------------------
			// TMPASMID        int64 // unique ID to manage fees uniquely across all fees in raflow json data
			// ASMID           int64 // the permanent table assessment id if it is an existing RAID
			// ARID            int64
			// ARName          string
			// ContractAmount  float64
			// RentCycle       int64
			// Start           JSONDate
			// Stop            JSONDate
			// AtSigningPreTax float64
			// SalesTax        float64
			if ar.FLAGS&(128) != 0 { // Is it a pet fee?
				petid := asms[j].AssocElemID // find the pet...
				for k := 0; k < len(raf.Pets); k++ {
					if raf.Pets[k].PETID == petid {
						raf.Meta.LastTMPASMID++
						var pf = RAFeesData{
							TMPASMID:       raf.Meta.LastTMPASMID,
							ARID:           asms[j].ARID,
							ASMID:          asms[j].ASMID,
							ARName:         ar.Name,
							RentCycle:      asms[j].RentCycle,
							Start:          JSONDate(asms[j].Start),
							Stop:           JSONDate(asms[j].Stop),
							ContractAmount: asms[j].Amount,
							Comment:        asms[j].Comment,
						}
						raf.Pets[k].Fees = append(raf.Pets[k].Fees, pf)
						break
					}
				}
			}
			//----------------------------------------------------------
			// Handle VEHICLE Fees
			//----------------------------------------------------------
			if ar.FLAGS&(256) != 0 { // Is it a vehicle fee?
				vid := asms[j].AssocElemID // find the vehicle...
				for k := 0; k < len(raf.Vehicles); k++ {
					if raf.Vehicles[k].VID == vid {
						raf.Meta.LastTMPASMID++
						var pf = RAFeesData{
							TMPASMID:       raf.Meta.LastTMPASMID,
							ARID:           asms[j].ARID,
							ASMID:          asms[j].ASMID,
							ARName:         ar.Name,
							ContractAmount: asms[j].Amount,
							RentCycle:      asms[j].RentCycle,
							Start:          JSONDate(asms[j].Start),
							Stop:           JSONDate(asms[j].Stop),
							Comment:        asms[j].Comment,
						}
						raf.Vehicles[k].Fees = append(raf.Vehicles[k].Fees, pf)
						break
					}
				}
			}
		}

		raf.Rentables = append(raf.Rentables, rfd)
	}

	// Console("\n\n******\nExiting ConvertRA2Flow, RAFLAGS = %d\n******\n\n", raf.Meta.RAFLAGS)
	return raf, nil
}

// addRAPtoFlow adds a new person to raf.People.  The renter/occupant flags
// are only set if the corresponding input bool value is set.
//
// INPUTS
//     tcid  = the tcid of the transactant to load
//      rid  - the rentable that they are tied to
//      raf  - pointer to the flow struct to update
//      chk  - check to see if the tcid exists in raf.People before adding.
//             This is not always necessary, but only the caller knows.
// isRenter  - true if we need to set the RAPerson isRenter bool to true.
//             It should be true for Payors.
// isOccupant- true if we need to set the RAPerson isOccupant bool to true.
//             It should be true for Users.
//
// RETURNS
//     any error encountered
//     raf is updated
//-----------------------------------------------------------------------------
func addRAPtoFlow(ctx context.Context, tcid, rid int64, raf *RAFlowJSONData, chk, isRenter, isOccupant bool) error {
	// Is this user already present?
	if chk {
		for l := 0; l < len(raf.People); l++ {
			if raf.People[l].TCID == tcid {
				if isRenter {
					raf.People[l].IsRenter = true
				}
				if isOccupant {
					raf.People[l].IsOccupant = true
				}
				return nil
			}
		}
	}

	rap, err := createRAFlowPerson(ctx, tcid, raf, isOccupant) // adds person AND associated pets and vehicles
	if err != nil {
		return err
	}

	if isRenter {
		rap.IsRenter = true
	}

	if isOccupant {
		rap.IsOccupant = true

		// only tie occupants to rentable
		var t RATiePeopleData
		t.BID = rap.BID
		t.TMPTCID = rap.TMPTCID
		if rid > 0 {
			t.PRID = rid
		}
		raf.Tie.People = append(raf.Tie.People, t)
	}

	// finally append in people list
	raf.People = append(raf.People, rap)
	return nil
}

// createRAFlowPerson returns a new RAPeopleFlowData based on the supplied
// tcid. It does not set the Renter or Occupant flags
//
// INPUTS
//          ctx  = db transaction context
//         tcid  = the tcid of the transactant to load
//          raf  = pointer to RAFlowJSONData
// addDependents = adds dependents (currently pets and vehicles) to the flow
//                 data in addition to the transactant data. The recommended
//                 usage of this flag is to set it to true when the person
//                 being added is a user.
//
// RETURNS
//     RAPeopleFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func createRAFlowPerson(ctx context.Context, tcid int64, raf *RAFlowJSONData, addDependents bool) (RAPeopleFlowData, error) {
	var p Transactant
	var pu User
	var pp Payor
	var pr Prospect
	var rap RAPeopleFlowData
	var err error

	raf.Meta.LastTMPTCID++
	rap.TMPTCID = raf.Meta.LastTMPTCID // set this now so it is available when creating pets and vehicles
	if err = GetTransactant(ctx, tcid, &p); err != nil {
		return rap, err
	}
	if err = GetUser(ctx, tcid, &pu); err != nil {
		return rap, err
	}
	if err = GetPayor(ctx, tcid, &pp); err != nil {
		return rap, err
	}
	if err = GetProspect(ctx, tcid, &pr); err != nil {
		return rap, err
	}
	MigrateStructVals(&p, &rap)
	MigrateStructVals(&pp, &rap)
	MigrateStructVals(&pu, &rap)
	MigrateStructVals(&pr, &rap)

	if addDependents {
		if err = addFlowPersonVehicles(ctx, tcid, rap.TMPTCID, raf); err != nil {
			return rap, err
		}
		if err = addFlowPersonPets(ctx, tcid, rap.TMPTCID, raf); err != nil {
			return rap, err
		}
	}
	return rap, nil
}

// addFlowPersonPets adds pets belonging to tcid to the supplied
// RAFlowJSONData struct
//
// INPUTS
//      ctx  = db transaction context
//     tcid  = the tcid of the transactant to load
//
// RETURNS
//     RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func addFlowPersonPets(ctx context.Context, tcid, tmptcid int64, raf *RAFlowJSONData) error {
	petList, err := GetPetsByTransactant(ctx, tcid)
	if err != nil {
		return err
	}
	for i := 0; i < len(petList); i++ {
		raf.Meta.LastTMPPETID++
		var p = RAPetsFlowData{
			TMPTCID:  tmptcid,
			TMPPETID: raf.Meta.LastTMPPETID,
			Fees:     []RAFeesData{},
		}
		MigrateStructVals(&petList[i], &p)
		raf.Pets = append(raf.Pets, p)
	}
	return nil
}

// addFlowPersonVehicles adds vehicles belonging to tcid to the supplied
// RAFlowJSONData struct
//
// INPUTS
//      ctx  = db transaction context
//     tcid  = the tcid of the transactant to load
//
// RETURNS
//     RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func addFlowPersonVehicles(ctx context.Context, tcid, tmptcid int64, raf *RAFlowJSONData) error {
	vehicleList, err := GetVehiclesByTransactant(ctx, tcid)
	if err != nil {
		return err
	}
	for i := 0; i < len(vehicleList); i++ {
		raf.Meta.LastTMPVID++
		var v = RAVehiclesFlowData{
			TMPTCID: tmptcid,
			TMPVID:  raf.Meta.LastTMPVID,
			Fees:    []RAFeesData{},
		}
		MigrateStructVals(&vehicleList[i], &v)
		raf.Vehicles = append(raf.Vehicles, v)
	}
	return nil
}

// NewRAFlowPet create new pet entry for the raflow and returns strcture
// with fees configured it in bizprops
//
// INPUTS
//             ctx  = db transaction context
//             BID  = Business ID
//          pStart  = possession start date
//           pStop  = possession stop date
//            meta  = RAFlowMetaInfo data
//
// RETURNS
//     RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func NewRAFlowPet(ctx context.Context, BID int64, rStart, rStop, pStart, pStop JSONDate, meta *RAFlowMetaInfo) (pet RAPetsFlowData, err error) {
	const funcname = "NewRAFlowPet"
	fmt.Printf("Entered in %s\n", funcname)

	// initialize
	// assign new TMPPETID & mark in meta info
	meta.LastTMPPETID++
	pet = RAPetsFlowData{
		TMPPETID: meta.LastTMPPETID,
		DtStart:  pStart,
		DtStop:   pStop,
		Fees:     []RAFeesData{},
	}

	// GET PET INITIAL FEES, META SHOULD BE UPDATED IN CALLER FUNCTION
	pet.Fees, err = GetRAFlowInitialPetFees(ctx, BID, (time.Time)(rStart), (time.Time)(rStop), meta)

	return
}

// NewRAFlowVehicle create new vehicle entry for the raflow and returns strcture
// with fees configured it in bizprops
//
// INPUTS
//             ctx  = db transaction context
//             BID  = Business ID
//          pStart  = possession start date
//           pStop  = possession stop date
//            meta  = RAFlowMetaInfo data
//
// RETURNS
//     RAVehiclesFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func NewRAFlowVehicle(ctx context.Context, BID int64, rStart, rStop, pStart, pStop JSONDate, meta *RAFlowMetaInfo) (vehicle RAVehiclesFlowData, err error) {
	const funcname = "NewRAFlowVehicle"
	fmt.Printf("Entered in %s\n", funcname)

	// initialize
	// assign new TMPVID & mark in meta info
	meta.LastTMPVID++
	vehicle = RAVehiclesFlowData{
		TMPVID:  meta.LastTMPVID,
		DtStart: pStart,
		DtStop:  pStop,
		Fees:    []RAFeesData{},
	}

	// GET VEHICLE INITIAL FEES, META SHOULD BE UPDATED IN CALLER FUNCTION
	vehicle.Fees, err = GetRAFlowInitialVehicleFees(ctx, BID, (time.Time)(rStart), (time.Time)(rStop), meta)

	return
}
