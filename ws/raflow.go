package ws

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"sort"
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
}

// RADatesFlowData contains data in the dates part of RA flow
type RADatesFlowData struct {
	BID             int64
	AgreementStart  rlib.JSONDate // TermStart
	AgreementStop   rlib.JSONDate // TermStop
	RentStart       rlib.JSONDate
	RentStop        rlib.JSONDate
	PossessionStart rlib.JSONDate
	PossessionStop  rlib.JSONDate
}

// RAPeopleFlowData contains data in the background-info part of RA flow
type RAPeopleFlowData struct {
	TMPTCID int64
	BID     int64
	TCID    int64

	// Transactant fields
	// Role
	IsRenter    bool
	IsOccupant  bool
	IsGuarantor bool

	// Applicant information
	FirstName      string
	MiddleName     string
	LastName       string
	PreferredName  string
	IsCompany      bool
	CompanyName    string
	SSN            string
	DriversLicense string
	PrimaryEmail   string
	SecondaryEmail string
	WorkPhone      string
	CellPhone      string
	Address        string
	Address2       string
	City           string
	State          string
	PostalCode     string
	Country        string
	Website        string

	// User fields
	DateofBirth rlib.JSONDate

	// Emergency contact information
	EmergencyContactName    string
	EmergencyContactAddress string
	EmergencyContactPhone   string
	EmergencyContactEmail   string
	AlternateAddress        string
	EligibleFutureUser      bool
	Industry                string
	SourceSLSID             int64

	// Payor fields
	CreditLimit         float64
	TaxpayorID          string
	ThirdPartySource    int64
	GrossIncome         float64
	EligibleFuturePayor bool

	// Prospect fields
	CompanyAddress    string
	CompanyCity       string
	CompanyState      string
	CompanyPostalCode string
	CompanyEmail      string
	CompanyPhone      string
	Occupation        string

	// Current Address information
	CurrentAddress           string
	CurrentLandLordName      string
	CurrentLengthOfResidency string
	CurrentLandLordPhoneNo   string
	CurrentReasonForMoving   int64 // Reason for moving

	// Prior Address information
	PriorAddress           string
	PriorLandLordName      string
	PriorLengthOfResidency string
	PriorLandLordPhoneNo   string
	PriorReasonForMoving   int64 // Reason for moving

	// Have you ever been
	Evicted                  bool // Evicted
	EvictedDes               string
	Convicted                bool // Arrested or convicted of a Convicted
	ConvictedDes             string
	Bankruptcy               bool // Declared Bankruptcy
	BankruptcyDes            string
	DesiredUsageStartDate    rlib.JSONDate
	RentableTypePreference   int64
	FLAGS                    int64
	Approver                 int64
	DeclineReasonSLSID       int64
	OtherPreferences         string
	FollowUpDate             rlib.JSONDate
	CSAgent                  int64
	OutcomeSLSID             int64
	CommissionableThirdParty string

	// RA Application information
	Comment string // In an effort to accommodate you, please advise us of any special needs
}

// RAPetsFlowData contains data in the pets part of RA flow
type RAPetsFlowData struct {
	TMPPETID             int64
	BID                  int64
	PETID                int64
	TMPTCID              int64
	Name                 string
	Type                 string
	Breed                string
	Color                string
	Weight               int
	DtStart              rlib.JSONDate
	DtStop               rlib.JSONDate
	NonRefundablePetFee  float64
	RefundablePetDeposit float64
	RecurringPetFee      float64
}

// RAVehiclesFlowData contains data in the vehicles part of RA flow
type RAVehiclesFlowData struct {
	TMPVID              int64
	BID                 int64
	VID                 int64
	TMPTCID             int64
	VIN                 string
	VehicleType         string
	VehicleMake         string
	VehicleModel        string
	VehicleColor        string
	VehicleYear         int64
	LicensePlateState   string
	LicensePlateNumber  string
	ParkingPermitNumber string
	ParkingPermitFee    float64
	DtStart             rlib.JSONDate
	DtStop              rlib.JSONDate
}

// RARentablesFlowData contains data in the rentables part of RA flow
type RARentablesFlowData struct {
	BID          int64
	RID          int64
	RTID         int64
	RTFLAGS      uint64
	RentableName string
	RentCycle    int64
	AtSigningAmt float64
	ProrateAmt   float64
	TaxableAmt   float64
	SalesTax     float64
	TransOcc     float64
	Fees         []RARentableFeesData
}

// RARentableFeesData struct
type RARentableFeesData struct {
	BID             int64
	RID             int64
	ARID            int64
	ARName          string
	ContractAmount  float64
	RentCycle       int64
	Epoch           int64
	RentPeriodStart rlib.JSONDate
	RentPeriodStop  rlib.JSONDate
	UsePeriodStart  rlib.JSONDate
	UsePeriodStop   rlib.JSONDate
	AtSigningAmt    float64
	ProrateAmt      float64
	SalesTaxAmt     float64
	SalesTax        float64
	TransOccAmt     float64
	TransOcc        float64
}

// RAParentChildFlowData contains data in the Parent/Child part of RA flow
type RAParentChildFlowData struct {
	BID  int64
	PRID int64 // parent rentable ID
	CRID int64 // child rentable ID
}

// RATieFlowData contains data in the tie part of RA flow
type RATieFlowData struct {
	Pets     []RATiePetsData     `json:"pets"`
	Vehicles []RATieVehiclesData `json:"vehicles"`
	People   []RATiePeopleData   `json:"people"`
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
	BID     int64
	PRID    int64
	TMPTCID int64 // user's temp json record reference id
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

	// TODO: Add validation on field level, it must be done.

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
				if a[i].TMPPETID == 0 { // if zero then assign new from last saved ID
					raFlowData.Meta.LastTMPPETID++
					a[i].TMPPETID = raFlowData.Meta.LastTMPPETID
				}
			}

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
				if a[i].TMPVID == 0 { // if zero then assign new from last saved ID
					raFlowData.Meta.LastTMPVID++
					a[i].TMPVID = raFlowData.Meta.LastTMPVID
				}
			}

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

			// check for each rentable data's Fees field
			// if it's blank then initialize it
			for i := range a {
				if len(a[i].Fees) == 0 {
					a[i].Fees = []RARentableFeesData{}
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
			if len(a.Pets) == 0 {
				a.Pets = []RATiePetsData{}
			}
			if len(a.Vehicles) == 0 {
				a.Vehicles = []RATieVehiclesData{}
			}
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
			Pets:     []RATiePetsData{},
			Vehicles: []RATieVehiclesData{},
			People:   []RATiePeopleData{},
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

// RARentableFeesDataRequest is struct for request for rentable fees
type RARentableFeesDataRequest struct {
	RID    int64
	FlowID int64
}

// SvcGetRAFlowRentableFeesData generates a list of rentable fees with auto populate AR fees
// It modifies raflow json doc by writing Fees data to raflow "rentables" component data
// wsdoc {
//  @Title Get list of Rentable fees with auto populate AR fees
//  @URL /v1/raflow-rentable-fees/:BUI/
//  @Method  GET
//  @Synopsis Get Rentable Fees list
//  @Description Get all rentable fees with auto populate AR fees
//  @Input RARentableFeesDataRequest
//  @Response FlowResponse
// wsdoc }
func SvcGetRAFlowRentableFeesData(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcGetRAFlowRentableFeesData"
	var (
		g           FlowResponse
		rfd         RARentablesFlowData
		raFlowData  RAFlowJSONData
		foo         RARentableFeesDataRequest
		feesRecords = []RARentableFeesData{}
		today       = time.Now()
	)
	fmt.Printf("Entered %s\n", funcname)

	if r.Method != "POST" {
		err := fmt.Errorf("Only POST method is allowed")
		SvcErrorReturn(w, err, funcname)
		return
	}

	if err := json.Unmarshal([]byte(d.data), &foo); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get flow and it must exist
	flow, err := rlib.GetFlow(r.Context(), foo.FlowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get rentable
	rentable, err := rlib.GetRentable(r.Context(), foo.RID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get rentableType
	rtid, err := rlib.GetRTIDForDate(r.Context(), foo.RID, &today)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	var rt rlib.RentableType
	err = rlib.GetRentableType(r.Context(), rtid, &rt)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// now get account rule based on this rentabletype
	ar, _ := rlib.GetAR(r.Context(), rt.ARID)
	if ar.ARID > 0 {
		// make sure the IsRentASM is marked true
		if ar.FLAGS&0x10 != 0 {
			rec := RARentableFeesData{
				BID:             ar.BID,
				ARID:            ar.ARID,
				RID:             foo.RID,
				ARName:          ar.Name,
				ContractAmount:  ar.DefaultAmount,
				RentPeriodStart: rlib.JSONDate(today),
				RentPeriodStop:  rlib.JSONDate(today.AddDate(1, 0, 0)),
				UsePeriodStart:  rlib.JSONDate(today),
				UsePeriodStop:   rlib.JSONDate(today.AddDate(1, 0, 0)),
			}

			// If it have is non recur charge true
			if ar.FLAGS&0x40 != 0 {
				rec.RentCycle = 0 // norecur: index 0 in app.cycleFreq
			} else {
				rec.RentCycle = rt.RentCycle
			}

			feesRecords = append(feesRecords, rec)
		}
	}

	// get all auto populated to new RA marked account rules by integer representation
	arFLAGVal := 1 << uint64(bizlogic.ARFLAGS["AutoPopulateToNewRA"])
	m, err := rlib.GetARsByFLAGS(r.Context(), d.BID, uint64(arFLAGVal))
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// append feesRecords in ascending order
	for _, ar := range m {
		if ar.FLAGS&0x10 != 0 { // if it's rent asm then continue
			continue
		}

		rec := RARentableFeesData{
			BID:             ar.BID,
			ARID:            ar.ARID,
			RID:             foo.RID,
			ARName:          ar.Name,
			ContractAmount:  ar.DefaultAmount,
			RentPeriodStart: rlib.JSONDate(today),
			RentPeriodStop:  rlib.JSONDate(today.AddDate(1, 0, 0)),
			UsePeriodStart:  rlib.JSONDate(today),
			UsePeriodStop:   rlib.JSONDate(today.AddDate(1, 0, 0)),
		}

		// If it have is non recur charge  flag true
		if ar.FLAGS&0x40 != 0 {
			rec.RentCycle = 0 // norecur: index 0 in app.cycleFreq
		} else {
			rec.RentCycle = rt.RentCycle
		}

		/*if ar.FLAGS&0x20 != 0 { // same will be applied to Security Deposit ASM
			rec.Amount = ar.DefaultAmount
		}*/

		// now append rec in feesRecords
		feesRecords = append(feesRecords, rec)
	}

	// sort based on name, needs version 1.8 later of golang
	sort.Slice(feesRecords, func(i, j int) bool { return feesRecords[i].ARName < feesRecords[j].ARName })

	// assign calculated data in rentable data
	rfd.BID = d.BID
	rfd.RID = rentable.RID
	rfd.RentableName = rentable.RentableName
	rfd.RTID = rt.RTID
	rfd.RTFLAGS = rt.FLAGS
	rfd.RentCycle = rt.RentCycle
	rfd.Fees = feesRecords

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// find this RID in flow data rentable list
	var rIndex = -1
	for i := range raFlowData.Rentables {
		if raFlowData.Rentables[i].RID == rfd.RID {
			rIndex = i
		}
	}

	// if record not found then push it in the list
	if rIndex < 0 {
		raFlowData.Rentables = append(raFlowData.Rentables, rfd)
	} else {
		raFlowData.Rentables[rIndex] = rfd
	}

	modRData, err := json.Marshal(&raFlowData.Rentables)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// update flow with this modified rentable part
	err = rlib.UpdateFlowData(r.Context(), "rentables", modRData, &flow)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get the modified flow
	flow, err = rlib.GetFlow(r.Context(), flow.FlowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// set the response
	g.Record = flow
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// RAPersonDetailsRequest is struct for request for person details
type RAPersonDetailsRequest struct {
	TCID   int64
	FlowID int64
}

// RAFlowRemovePersonRequest is struct for request to remove person from json data
type RAFlowRemovePersonRequest struct {
	TMPTCID int64
	FlowID  int64
}

// SvcGetRAFlowPersonHandler handles operation on person of raflow json data
//           0    1     2   3
// uri /v1/raflow-person/BID/flowID
// The server command can be:
//      get
//      delete
//-----------------------------------------------------------------------------------
func SvcGetRAFlowPersonHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcGetRAFlowPersonHandler"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "save":
		SaveRAFlowPersonDetails(w, r, d)
		break
	case "delete":
		DeleteRAFlowPerson(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// SaveRAFlowPersonDetails saves person details with list of pets and vehicles
// It modifies raflow json doc by writing fetched pets and vehicles data
// wsdoc {
//  @Title Save Person details with list of Pets & Vehicles
//  @URL /v1/raflow-persondetails/:BUI/
//  @Method  GET
//  @Synopsis Save Person Details for RAFlow
//  @Description Save details about person with pets and vehicles
//  @Input RAPersonDetailsRequest
//  @Response FlowResponse
// wsdoc }
func SaveRAFlowPersonDetails(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SaveRAFlowPersonDetails"
	var (
		raFlowData           RAFlowJSONData
		foo                  RAPersonDetailsRequest
		newRAFlowMeta        RAFlowMetaInfo // we might need to update meta info
		shouldModifyMetaData bool
		g                    FlowResponse
		err                  error
		tx                   *sql.Tx
		ctx                  context.Context
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

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("Only POST method is allowed")
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

	// ----------------------------------------------
	// get person details with given TCID
	// ----------------------------------------------
	personTMPTCID := int64(0)

	// this is for accept Transactant, so find it by TCID
	tcidExistInJSONData := false
	for i := range raFlowData.People {
		if raFlowData.People[i].TCID == foo.TCID {
			tcidExistInJSONData = true
			personTMPTCID = raFlowData.People[i].TMPTCID
			break
		}
	}

	if !tcidExistInJSONData {
		newRAFlowPerson := RAPeopleFlowData{}
		var xp rlib.XPerson
		err = rlib.GetXPerson(ctx, foo.TCID, &xp)
		if err != nil {
			return
		}

		// migrate field values to Person details
		if xp.Pay.TCID > 0 {
			rlib.MigrateStructVals(&xp.Pay, &newRAFlowPerson)
		}
		if xp.Psp.TCID > 0 {
			rlib.MigrateStructVals(&xp.Psp, &newRAFlowPerson)
		}
		if xp.Usr.TCID > 0 {
			rlib.MigrateStructVals(&xp.Usr, &newRAFlowPerson)
		}
		if xp.Trn.TCID > 0 {
			rlib.MigrateStructVals(&xp.Trn, &newRAFlowPerson)
		}
		newRAFlowPerson.BID = d.BID

		// check for additional flags IsRenter, IsOccupant
		newRAFlowPerson.IsOccupant = true
		if len(raFlowData.People) == 0 { // this is first transactant
			newRAFlowPerson.IsRenter = true
		}

		// custom tmp tcid
		newRAFlowPerson.TMPTCID = raFlowData.Meta.LastTMPTCID + 1
		newRAFlowMeta.LastTMPTCID = newRAFlowPerson.TMPTCID
		personTMPTCID = newRAFlowPerson.TMPTCID

		// we need to update meta at the end, as new TMPTCID assigned
		shouldModifyMetaData = true

		// append in json list
		raFlowData.People = append(raFlowData.People, newRAFlowPerson)

		var modPeopleData []byte
		modPeopleData, err = json.Marshal(&raFlowData.People)
		if err != nil {
			return
		}

		// update flow with this modified people part
		err = rlib.UpdateFlowData(ctx, "people", modPeopleData, &flow)
		if err != nil {
			return
		}
	}

	// -------------------------------------------
	// find pets list associated with current TCID
	// -------------------------------------------

	// get the list of pets
	var petList []rlib.RentalAgreementPet
	petList, err = rlib.GetPetsByTransactant(ctx, foo.TCID)
	if err != nil {
		return
	}

	// find this RID in flow data rentable list
	shouldModifyPetsData := false
	for i := range petList {
		exist := false
		for k := range raFlowData.Pets {
			if petList[i].PETID == raFlowData.Pets[k].PETID {
				exist = true
				break
			}
		}

		// if does not exist then append in the raflow data
		if !exist {
			var newRAFlowPet RAPetsFlowData
			rlib.MigrateStructVals(&petList[i], &newRAFlowPet)

			// assign new TMPPETID & mark in meta info
			newRAFlowPet.TMPPETID = raFlowData.Meta.LastTMPPETID + 1
			newRAFlowMeta.LastTMPPETID = newRAFlowPet.TMPPETID

			// reference of TMPTCID for the person
			newRAFlowPet.TMPTCID = personTMPTCID

			// append in pets list of json data
			raFlowData.Pets = append(raFlowData.Pets, newRAFlowPet)
			shouldModifyPetsData = true
			shouldModifyMetaData = true // as new TMPPETID
		}
	}

	if shouldModifyPetsData {
		var modPetsData []byte
		modPetsData, err = json.Marshal(&raFlowData.Pets)
		if err != nil {
			return
		}

		// update flow with this modified pets part
		err = rlib.UpdateFlowData(ctx, "pets", modPetsData, &flow)
		if err != nil {
			return
		}
	}

	// -----------------------------------------------
	// find vehicles list associated with current TCID
	// -----------------------------------------------

	// get the list of pets
	var vehicleList []rlib.Vehicle
	vehicleList, err = rlib.GetVehiclesByTransactant(ctx, foo.TCID)
	if err != nil {
		return
	}

	// loop over list and append it in raflow data
	shouldModifyVehiclesData := false
	for i := range vehicleList {
		exist := false
		for k := range raFlowData.Vehicles {
			if vehicleList[i].VID == raFlowData.Vehicles[k].VID {
				exist = true
				break
			}
		}

		// if does not exist then append in the raflow data
		if !exist {
			var newRAFlowVehicle RAVehiclesFlowData
			rlib.MigrateStructVals(&vehicleList[i], &newRAFlowVehicle)

			// assign new TMPVID
			newRAFlowVehicle.TMPVID = raFlowData.Meta.LastTMPVID + 1
			newRAFlowMeta.LastTMPVID = newRAFlowVehicle.TMPVID

			// reference of TMPTCID to this vehicle
			newRAFlowVehicle.TMPTCID = personTMPTCID

			// append in vehicles list of json data
			raFlowData.Vehicles = append(raFlowData.Vehicles, newRAFlowVehicle)
			shouldModifyVehiclesData = true
			shouldModifyMetaData = true // as new TMPVID assigned
		}
	}

	if shouldModifyVehiclesData {
		// get marshalled data
		var modVData []byte
		modVData, err = json.Marshal(&raFlowData.Vehicles)
		if err != nil {
			return
		}

		// update flow with this modified vehicles part
		err = rlib.UpdateFlowData(ctx, "vehicles", modVData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// update meta info if required
	// ----------------------------------------------
	if shouldModifyMetaData {
		// get marshalled data
		var modMetaData []byte
		modMetaData, err = json.Marshal(&newRAFlowMeta)
		if err != nil {
			return
		}

		// update flow with this modified meta part
		err = rlib.UpdateFlowData(ctx, "meta", modMetaData, &flow)
		if err != nil {
			return
		}
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

// DeleteRAFlowPerson remove person from raflow data as well as removes
// associated pets and vehicles data too
// wsdoc {
//  @Title Remvoe Person with list of associated Pets & Vehicles
//  @URL /v1/raflow-person/:BUI/:FlowID
//  @Method POST
//  @Synopsis Remove Person from RAFlow json data
//  @Description Remove details about person with associated pets and vehicles
//  @Input RAFlowRemovePersonRequest
//  @Response FlowResponse
// wsdoc }
func DeleteRAFlowPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "DeleteRAFlowPerson"
	var (
		raFlowData RAFlowJSONData
		foo        RAFlowRemovePersonRequest
		g          FlowResponse
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
			if tx != nil {
				tx.Rollback()
			}
			SvcErrorReturn(w, err, funcname)
			return
		}
	}()

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("Only POST method is allowed")
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

	// ----------------------------------------------
	// get person details with given TMPTCID
	// ----------------------------------------------
	personTMPTCID := int64(0)

	// this is for accept Transactant, so find it by TMPTCID
	tcidExistInJSONData := false
	for i := range raFlowData.People {
		if raFlowData.People[i].TMPTCID == foo.TMPTCID {
			tcidExistInJSONData = true
			personTMPTCID = raFlowData.People[i].TMPTCID

			// remove the element then
			raFlowData.People = append(raFlowData.People[:i], raFlowData.People[i+1:]...)

			break
		}
	}

	if tcidExistInJSONData {
		var modPeopleData []byte
		modPeopleData, err = json.Marshal(&raFlowData.People)
		if err != nil {
			return
		}

		// update flow with this modified people part
		err = rlib.UpdateFlowData(ctx, "people", modPeopleData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// remove associated pets
	// ----------------------------------------------
	shouldModifyPetsData := false
	for i := range raFlowData.Pets {
		if raFlowData.Pets[i].TMPTCID == personTMPTCID {
			shouldModifyPetsData = true
			// remove this pet from the list
			raFlowData.Pets = append(raFlowData.Pets[:i], raFlowData.Pets[i+1:]...)
		}
	}

	if shouldModifyPetsData {
		var modPetsData []byte
		modPetsData, err = json.Marshal(&raFlowData.Pets)
		if err != nil {
			return
		}

		// update flow with this modified pets part
		err = rlib.UpdateFlowData(ctx, "pets", modPetsData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// remove associated vehicles
	// ----------------------------------------------
	shouldModifyVehiclesData := false
	for i := range raFlowData.Vehicles {
		if raFlowData.Vehicles[i].TMPTCID == personTMPTCID {
			shouldModifyVehiclesData = true
			// remove this pet from the list
			raFlowData.Vehicles = append(raFlowData.Vehicles[:i], raFlowData.Vehicles[i+1:]...)
		}
	}

	if shouldModifyVehiclesData {
		// get marshalled data
		var modVData []byte
		modVData, err = json.Marshal(&raFlowData.Vehicles)
		if err != nil {
			return
		}

		// update flow with this modified vehicles part
		err = rlib.UpdateFlowData(ctx, "vehicles", modVData, &flow)
		if err != nil {
			return
		}
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
