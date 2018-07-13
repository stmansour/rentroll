package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"rentroll/rlib"
	"time"
)

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
	RAID         int64 // 0 = it's new, >0 = existing one
	LastTMPPETID int64
	LastTMPVID   int64
	LastTMPTCID  int64
	LastTMPASMID int64
	HavePets     bool
	HaveVehicles bool
	RAFLAGS      int64
}

// RADatesFlowData contains data in the dates part of RA flow
type RADatesFlowData struct {
	BID             int64         `validate:"number,min=1"`
	AgreementStart  rlib.JSONDate `validate:"date"` // TermStart
	AgreementStop   rlib.JSONDate `validate:"date"` // TermStop
	RentStart       rlib.JSONDate `validate:"date"`
	RentStop        rlib.JSONDate `validate:"date"`
	PossessionStart rlib.JSONDate `validate:"date"`
	PossessionStop  rlib.JSONDate `validate:"date"`
}

// RAPeopleFlowData contains data in the background-info part of RA flow
type RAPeopleFlowData struct {
	TMPTCID int64 `validate:"number,min=1"`
	BID     int64 `validate:"number,min=1"`
	TCID    int64 `validate:"number,min=1"`

	// Role
	IsRenter    bool `validate:"-"`
	IsOccupant  bool `validate:"-"`
	IsGuarantor bool `validate:"-"`

	// ---------- Basic Info -----------
	FirstName      string `validate:"string,min=1,max=100"`
	MiddleName     string `validate:"string,min=1,max=100"`
	LastName       string `validate:"string,min=1,max=100"`
	PreferredName  string `validate:"string,min=1,max=100,omitempty"`
	IsCompany      bool   `validate:"number,min=1,max=1"`
	CompanyName    string `validate:"string,min=1,max=100"`
	PrimaryEmail   string `validate:"email,omitempty"`
	SecondaryEmail string `validate:"email,omitempty"`
	WorkPhone      string `validate:"number,min=1,max=100,omitempty"`
	CellPhone      string `validate:"number,min=1,max=100,omitempty"`
	Address        string `validate:"string,min=1,max=100"`
	Address2       string `validate:"string,min=0,max=100,omitempty"`
	City           string `validate:"string,min=1,max=100"`
	State          string `validate:"string,min=1,max=25"`
	PostalCode     string `validate:"number,min=1,max=100"`
	Country        string `validate:"string,min=1,max=100"`
	Website        string `validate:"string,min=1,max=100,omitempty"`
	Comment        string `validate:"string,min=1,max=2048,omitempty"`

	// ---------- Prospect -----------
	CompanyAddress    string `validate:"string,min=1,max=100"`
	CompanyCity       string `validate:"string,min=1,max=100"`
	CompanyState      string `validate:"string,min=1,max=100"`
	CompanyPostalCode string `validate:"number,min=1,max=100"`
	CompanyEmail      string `validate:"email"`
	CompanyPhone      string `validate:"number,min=1,max=100"`
	Occupation        string `validate:"string,min=1,max=100"`

	// Current Address information
	CurrentAddress           string `validate:"string,min=1,max=100"`
	CurrentLandLordName      string `validate:"string,min=1,max=100"`
	CurrentLandLordPhoneNo   string `validate:"number,min=1"`
	CurrentLengthOfResidency string `validate:"string,min=1,max=100"`
	CurrentReasonForMoving   int64  `validate:"number,min=1"` // Reason for moving

	// Prior Address information
	PriorAddress           string `validate:"string,min=1,max=100"`
	PriorLandLordName      string `validate:"string,min=1,max=100"`
	PriorLandLordPhoneNo   string `validate:"number,min=1"`
	PriorLengthOfResidency string `validate:"string,min=1,max=100"`
	PriorReasonForMoving   int64  `validate:"number,min=1"` // Reason for moving

	// Have you ever been
	Evicted          bool   `validate:"-"` // Evicted
	EvictedDes       string `validate:"string,min=1,max=2048,omitempty"`
	Convicted        bool   `validate:"-"` // Arrested or convicted of a Convicted
	ConvictedDes     string `validate:"string,min=1,max=2048,omitempty"`
	Bankruptcy       bool   `validate:"-"` // Declared Bankruptcy
	BankruptcyDes    string `validate:"string,min=1,max=2048,omitempty"`
	OtherPreferences string `validate:"string,min=1,max=1024"`
	//FollowUpDate             rlib.JSONDate
	//CommissionableThirdParty string
	SpecialNeeds string `validate:"string,min=1,max=1024"` // In an effort to accommodate you, please advise us of any special needs

	// ---------- Payor -----------
	CreditLimit         float64 `validate:"number:float,min=0.10"`
	TaxpayorID          string  `validate:"string,min=1,max=25"`
	GrossIncome         float64 `validate:"number:float,min=0.10"`
	SSN                 string  `validate:"string,min=1,max=128"`
	DriversLicense      string  `validate:"string,min=1,max=128"`
	ThirdPartySource    int64   `validate:"number,min=1,omitempty"`
	EligibleFuturePayor bool    `validate:"-"`

	// ---------- User -----------
	Points      int64         `validate:"number,min=1,omitempty"`
	DateofBirth rlib.JSONDate `validate:"date"`
	// Emergency contact information
	EmergencyContactName      string `validate:"string,min=1,max=100"`
	EmergencyContactAddress   string `validate:"string,min=1,max=100"`
	EmergencyContactTelephone string `validate:"number,min=1,max=100"`
	EmergencyContactEmail     string `validate:"email"`
	AlternateAddress          string `validate:"string,min=1,max=100"`
	EligibleFutureUser        bool   `validate:"number,min=1"`
	Industry                  string `validate:"string,min=1,max=100"`
	SourceSLSID               int64  `validate:"number,min=1"`
}

// RAPetsFlowData contains data in the pets part of RA flow
type RAPetsFlowData struct {
	TMPPETID int64         `validate:"number,min=1"`
	BID      int64         `validate:"number,min=1"`
	PETID    int64         `validate:"number,min=1"`
	TMPTCID  int64         `validate:"number,min=1"`
	Name     string        `validate:"string,min=1,max=100"`
	Type     string        `validate:"string,min=1,max=100"`
	Breed    string        `validate:"string,min=1,max=100"`
	Color    string        `validate:"string,min=1,max=100"`
	Weight   float64       `validate:"number:float,min=0.0"`
	DtStart  rlib.JSONDate `validate:"date"`
	DtStop   rlib.JSONDate `validate:"date"`
	Fees     []RAFeesData  `validate:"-"`
}

// RAVehiclesFlowData contains data in the vehicles part of RA flow
type RAVehiclesFlowData struct {
	TMPVID              int64         `validate:"number,min=1"`
	BID                 int64         `validate:"number,min=1"`
	VID                 int64         `validate:"number,min=1"`
	TMPTCID             int64         `validate:"number,min=1"`
	VIN                 string        `validate:"string,min=1"`
	VehicleType         string        `validate:"string,min=1,max=80"`
	VehicleMake         string        `validate:"string,min=1,max=80"`
	VehicleModel        string        `validate:"string,min=1,max=80"`
	VehicleColor        string        `validate:"string,min=1,max=80"`
	VehicleYear         int64         `validate:"number,min=1"`
	LicensePlateState   string        `validate:"string,min=1,max=80"`
	LicensePlateNumber  string        `validate:"string,min=1,max=80"`
	ParkingPermitNumber string        `validate:"string,min=1,max=80"`
	DtStart             rlib.JSONDate `validate:"date"`
	DtStop              rlib.JSONDate `validate:"date"`
	Fees                []RAFeesData  `validate:"-"`
}

// RARentablesFlowData contains data in the rentables part of RA flow
type RARentablesFlowData struct {
	BID             int64 `validate:"number,min=1"`
	RID             int64 `validate:"number,min=1"`
	RTID            int64 `validate:"number,min=1"`
	RTFLAGS         uint64
	RentableName    string  `validate:"string,min=1,max=100"`
	RentCycle       int64   `validate:"number,min=1"`
	AtSigningPreTax float64 `validate:"number:float,min=0.00"`
	SalesTax        float64 `validate:"number:float,min=0.00"`
	// SalesTaxAmt    float64 // FUTURE RELEASE
	TransOccTax float64 `validate:"number:float,min=0.00"`
	// TransOccAmt    float64 // FUTURE RELEASE
	Fees []RAFeesData `validate:"-"`
}

// RAFeesData struct used for pet, vehicles, rentable fees
type RAFeesData struct {
	TMPASMID        int64         `validate:"number,min=1"` // unique ID to manage fees uniquely across all fees in raflow json data
	ASMID           int64         `validate:"number,min=0"` // the permanent table assessment id if it is an existing RAID
	ARID            int64         `validate:"number,min=1"`
	ARName          string        `validate:"string,min=1,max=100"`
	ContractAmount  float64       `validate:"number:float,min=0.00"`
	RentCycle       int64         `validate:"number,min=1"`
	Start           rlib.JSONDate `validate:"date"`
	Stop            rlib.JSONDate `validate:"date"`
	AtSigningPreTax float64       `validate:"number:float,min=0.00"`
	SalesTax        float64       `validate:"number:float,min=0.00"`
	// SalesTaxAmt    float64 // FUTURE RELEASE
	TransOccTax float64 `validate:"number:float,min=0.00"`
	// TransOccAmt    float64 // FUTURE RELEASE
}

// RAParentChildFlowData contains data in the Parent/Child part of RA flow
type RAParentChildFlowData struct {
	BID  int64 `validate:"number,min=1"`
	PRID int64 `validate:"number,min=1"` // parent rentable ID
	CRID int64 `validate:"number,min=1"` // child rentable ID
}

// RATieFlowData contains data in the tie part of RA flow
type RATieFlowData struct {
	People []RATiePeopleData `json:"people"`
}

// RATiePetsData holds data from tie section for a pet to a rentable
type RATiePetsData struct {
	BID      int64
	PRID     int64
	TMPPETID int64 // reference to pet record ID stored temporarily
}

// RATieVehiclesData holds data from tie section for a vehicle to a rentable
type RATieVehiclesData struct {
	BID    int64
	PRID   int64
	TMPVID int64 // reference to vehicle record ID in json
}

// RATiePeopleData holds data from tie section for a payor to a rentable
type RATiePeopleData struct {
	BID     int64 `validate:"number,min=1"`
	PRID    int64 `validate:"number,min=1"`
	TMPTCID int64 `validate:"number,min=1"` // user's temp json record reference id
}

// getUpdateRAFlowPartJSONData returns json data in bytes
// coming from client with checking of flow and part type to update
func getUpdateRAFlowPartJSONData(BID int64, data json.RawMessage, partType int, flow *rlib.Flow) ([]byte, []byte, error) {

	var (
		modFlowPartData = []byte(nil)
		modMetaData     = []byte(nil)
		err             error
		raFlowData      RAFlowJSONData
	)

	// get the whole raflow data from Flow type data
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		// if it's an error then return with nil data
		return modMetaData, modFlowPartData, err
	}

	// JSON Marshal with address
	// REF: https://stackoverflow.com/questions/21390979/custom-marshaljson-never-gets-called-in-go

	// is it blank string or null json data
	isBlankJSONData := bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))

	switch rlib.RAFlowPartType(partType) {
	case rlib.DatesRAFlowPart:
		a := RADatesFlowData{}

		// if the struct provided with some data then check it for
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				// if it's an error then return with nil data
				return modMetaData, modFlowPartData, err
			}
		} else {
			// it's null/blank data then initialize with default data
			currentDateTime := time.Now()
			nextYearDateTime := currentDateTime.AddDate(1, 0, 0)

			a.BID = BID
			a.RentStart = rlib.JSONDate(currentDateTime)
			a.RentStop = rlib.JSONDate(nextYearDateTime)
			a.AgreementStart = rlib.JSONDate(currentDateTime)
			a.AgreementStop = rlib.JSONDate(nextYearDateTime)
			a.PossessionStart = rlib.JSONDate(currentDateTime)
			a.PossessionStop = rlib.JSONDate(nextYearDateTime)
		}

		// json marshalled for struct
		modFlowPartData, err = json.Marshal(&a)

	case rlib.PeopleRAFlowPart:
		a := []RAPeopleFlowData{}

		// if the struct provided with some data then check it for
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)

			// auto assign TMPTCID
			for i := range a {
				if a[i].TMPTCID == 0 { // if zero then assign new from last saved ID
					raFlowData.Meta.LastTMPTCID++
					a[i].TMPTCID = raFlowData.Meta.LastTMPTCID
				}
			}

			if err != nil {
				// if it's an error then return with nil data
				return modMetaData, modFlowPartData, err
			}
		}

		// json marshalled for struct
		modFlowPartData, err = json.Marshal(&a)

	case rlib.PetsRAFlowPart:
		a := []RAPetsFlowData{}

		// if the struct provided with some data then check it for
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)

			// auto assign TMPPETID
			for i := range a {
				// If Fees not initialized then
				if len(a[i].Fees) == 0 {
					a[i].Fees = []RAFeesData{}
				}

				if a[i].TMPPETID == 0 { // if zero then assign new from last saved ID
					raFlowData.Meta.LastTMPPETID++
					a[i].TMPPETID = raFlowData.Meta.LastTMPPETID

					// manage TMPASMID in Fees
					for j := range a[i].Fees {
						if a[i].Fees[j].TMPASMID == 0 {
							raFlowData.Meta.LastTMPASMID++
							a[i].Fees[j].TMPASMID = raFlowData.Meta.LastTMPASMID
						}
					}
				}
			}

			// Update HavePets flag in meta information
			raFlowData.Meta.HavePets = len(a) > 0

			if err != nil {
				// if it's an error then return with nil data
				return modMetaData, modFlowPartData, err
			}
		}

		// json marshalled for struct
		modFlowPartData, err = json.Marshal(&a)

	case rlib.VehiclesRAFlowPart:
		a := []RAVehiclesFlowData{}

		// if the struct provided with some data then check it for
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)

			// auto assign TMPVID
			for i := range a {
				// If Fees not initialized then
				if len(a[i].Fees) == 0 {
					a[i].Fees = []RAFeesData{}
				}

				if a[i].TMPVID == 0 { // if zero then assign new from last saved ID
					raFlowData.Meta.LastTMPVID++
					a[i].TMPVID = raFlowData.Meta.LastTMPVID

					// manage TMPASMID in fees
					for j := range a[i].Fees {
						if a[i].Fees[j].TMPASMID == 0 {
							raFlowData.Meta.LastTMPASMID++
							a[i].Fees[j].TMPASMID = raFlowData.Meta.LastTMPASMID
						}
					}
				}
			}

			// Update HaveVehicles flag in meta information
			raFlowData.Meta.HaveVehicles = len(a) > 0

			if err != nil {
				// if it's an error then return with nil data
				return modMetaData, modFlowPartData, err
			}
		}

		// json marshalled for struct
		modFlowPartData, err = json.Marshal(&a)

	case rlib.RentablesRAFlowPart:
		a := []RARentablesFlowData{}

		// if the struct provided with some data then check it for
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)

			for i := range a {

				// If Fees not initialized then
				if len(a[i].Fees) == 0 {
					a[i].Fees = []RAFeesData{}
				}

				// manage TMPASMID in fees
				for j := range a[i].Fees {
					if a[i].Fees[j].TMPASMID == 0 {
						raFlowData.Meta.LastTMPASMID++
						a[i].Fees[j].TMPASMID = raFlowData.Meta.LastTMPASMID
					}
				}

			}

			if err != nil {
				// if it's an error then return with nil data
				return modMetaData, modFlowPartData, err
			}
		}

		// json marshalled for struct
		modFlowPartData, err = json.Marshal(&a)

	case rlib.ParentChildRAFlowPart:
		a := []RAParentChildFlowData{}

		// if the struct provided with some data then check it for
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				// if it's an error then return with nil data
				return modMetaData, modFlowPartData, err
			}
		}

		// json marshalled for struct
		modFlowPartData, err = json.Marshal(&a)

	case rlib.TieRAFlowPart:
		a := RATieFlowData{}

		// if the struct provided with some data then check it for
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)

			// check for each sliced data field
			// if it's blank then initialize it
			if len(a.People) == 0 {
				a.People = []RATiePeopleData{}
			}

			if err != nil {
				// if it's an error then return with nil data
				return modMetaData, modFlowPartData, err
			}
		}

		// json marshalled for struct
		modFlowPartData, err = json.Marshal(&a)

	default:
		err = fmt.Errorf("unrecognized part type in RA flow: %d", partType)
	}

	// if error occured in above switch cases execution
	// while marshaling content in json then only
	if err != nil {
		return modMetaData, modFlowPartData, err
	}

	// now marshal json data back to raflow
	modMetaData, err = json.Marshal(&raFlowData.Meta)
	if err != nil {
		// if it's an error then return with nil data
		return modMetaData, modFlowPartData, err
	}

	// finally return with modified data
	return modMetaData, modFlowPartData, err
}

// insertInitialRAFlow writes a bunch of flow's sections record for a particular RA
func insertInitialRAFlow(ctx context.Context, BID, UID int64) (int64, error) {

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
			RentStart:       rlib.JSONDate(currentDateTime),
			RentStop:        rlib.JSONDate(nextYearDateTime),
			AgreementStart:  rlib.JSONDate(currentDateTime),
			AgreementStop:   rlib.JSONDate(nextYearDateTime),
			PossessionStart: rlib.JSONDate(currentDateTime),
			PossessionStop:  rlib.JSONDate(nextYearDateTime),
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
		rlib.Ulog("Error while marshalling json data of initialRAFlow: %s\n", err.Error())
		return flowID, err
	}

	// initial Flow struct
	rlib.Console("New Flow\n")
	a := rlib.Flow{
		BID:       BID,
		FlowID:    0, // it's new flowID,
		UserRefNo: rlib.GenerateUserRefNo(),
		FlowType:  rlib.RAFlow,
		Data:      raflowJSONData,
		CreateBy:  UID,
		LastModBy: UID,
	}

	rlib.Console("New flow UserRefNo = %s\n", a.UserRefNo)

	// insert new flow
	flowID, err = rlib.InsertFlow(ctx, &a)
	if err != nil {
		rlib.Ulog("Error while inserting Flow: %s\n", err.Error())
		return flowID, err
	}

	return flowID, err
}

// saveRentalAgreementFlow saves data for the given flowID to real multi variant database instances
// from the temporary data stored in FlowPart table
func saveRentalAgreementFlow(ctx context.Context, flowID int64) (int64, error) {
	var (
		RAID int64
		err  error
	)

	// first check that such a given flowID does exist or not
	var found bool
	ids, err := rlib.GetFlowIDsByUser(ctx)
	if err != nil {
		return RAID, err
	}

	for _, id := range ids {
		if id == flowID {
			found = true
			break
		}
	}

	if !found {
		return RAID, fmt.Errorf("Such flowID: %d does not exist", flowID)
	}

	// -------------- SAVING PARTS --------------------

	/*// ==================
	// 1. Agreement Dates
	// ==================
	datesFlowPart, err := rlib.GetFlowPartByPartType(ctx, flowID, int(rlib.DatesRAFlowPart))
	if err != nil {
		return err
	}

	var dtFD RADatesFlowData
	err = json.Unmarshal(datesFlowPart.Data, &dtFD)
	if err != nil {
		return err
	}

	// now, create a rental agreement using this basic dates info
	var ra = rlib.RentalAgreement{
		RentStart:       time.Time(dtFD.RentStart),
		RentStop:        time.Time(dtFD.RentStop),
		AgreementStart:  time.Time(dtFD.AgreementStart),
		AgreementStop:   time.Time(dtFD.AgreementStop),
		PossessionStart: time.Time(dtFD.PossessionStart),
		PossessionStop:  time.Time(dtFD.PossessionStop),
	}
	RAID, err = rlib.InsertRentalAgreement(ctx, &ra)
	if err != nil {
		return err
	}
	fmt.Printf("Newly created rental agreement with RAID: %d\n", RAID)*/

	return RAID, nil
}

// GridRAFlowResponse is a struct to hold info for rental agreement for the grid response
type GridRAFlowResponse struct {
	Recid     int64 `json:"recid"`
	BID       int64
	BUD       string
	FlowID    int64
	UserRefNo string
}
