package ws

import (
	"bytes"
	"context"
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
	RAID             int64 // 0 = it's new, >0 = existing one
	PetLastTMPID     int64
	VehicleLastTMPID int64
	PeopleLastTMPID  int64
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
	TMPID int64
	BID   int64
	TCID  int64

	// Role
	IsRenter    bool
	IsOccupant  bool
	IsGuarantor bool

	// Applicant information
	FirstName    string
	MiddleName   string
	LastName     string
	BirthDate    string
	IsCompany    bool
	CompanyName  string
	SSN          string
	DriverLicNo  string
	TelephoneNo  string
	EmailAddress string
	Employer     string
	Phone        string
	Address      string
	Address2     string
	City         string
	State        string
	PostalCode   string
	Position     string
	GrossWages   float64

	// Current Address information
	CurrentAddress           string
	CurrentLandLordName      string
	CurrentLengthOfResidency int
	CurrentLandLordPhoneNo   string
	CurrentReasonForMoving   string // Reason for moving

	// Prior Address information
	PriorAddress           string
	PriorLandLordName      string
	PriorLengthOfResidency int
	PriorLandLordPhoneNo   string
	PriorReasonForMoving   string // Reason for moving

	// Have you ever been
	Evicted       bool // Evicted
	EvictedDes    string
	Convicted     bool // Arrested or convicted of a Convicted
	ConvictedDes  string
	Bankruptcy    bool // Declared Bankruptcy
	BankruptcyDes string

	// Emergency contact information
	EmergencyContactName    string
	EmergencyContactPhone   string
	EmergencyContactAddress string

	// RA Application information
	Comment string // In an effort to accommodate you, please advise us of any special needs
}

// RAPetsFlowData contains data in the pets part of RA flow
type RAPetsFlowData struct {
	TMPID                int64
	BID                  int64
	PETID                int64
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
	TMPID               int64
	BID                 int64
	VID                 int64
	TCID                int64
	VIN                 string
	Type                string
	Make                string
	Model               string
	Color               string
	Year                string
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
	Pets     []RAPetsTieData     `json:"pets"`
	Vehicles []RAVehiclesTieData `json:"vehicles"`
	Payors   []RAPayorsTieData   `json:"payors"`
}

// RAPetsTieData holds data from tie section for a pet to a rentable
type RAPetsTieData struct {
	BID   int64
	PRID  int64
	REFID int64 // reference to pet record ID stored temporarily
}

// RAVehiclesTieData holds data from tie section for a vehicle to a rentable
type RAVehiclesTieData struct {
	BID   int64
	PRID  int64
	REFID int64 // reference to vehicle record ID in json
}

// RAPayorsTieData holds data from tie section for a payor to a rentable
type RAPayorsTieData struct {
	BID   int64
	PRID  int64
	REFID int64 // user's temp json record reference id
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

			// auto assign TMPID
			for i := range a {
				if a[i].TMPID == 0 { // if zero then assign new from last saved ID
					raFlowData.Meta.PetLastTMPID++
					a[i].TMPID = raFlowData.Meta.PetLastTMPID
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

			// auto assign TMPID
			for i := range a {
				if a[i].TMPID == 0 { // if zero then assign new from last saved ID
					raFlowData.Meta.VehicleLastTMPID++
					a[i].TMPID = raFlowData.Meta.VehicleLastTMPID
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
				a.Pets = []RAPetsTieData{}
			}
			if len(a.Vehicles) == 0 {
				a.Vehicles = []RAVehiclesTieData{}
			}
			if len(a.Payors) == 0 {
				a.Payors = []RAPayorsTieData{}
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
			Pets:     []RAPetsTieData{},
			Vehicles: []RAVehiclesTieData{},
			Payors:   []RAPayorsTieData{},
		},
	}

	// get json marshelled byte data for above struct
	raflowJSONData, err := json.Marshal(&initialRAFlow)
	if err != nil {
		rlib.Ulog("Error while marshalling json data of initialRAFlow: %s\n", err.Error())
		return flowID, err
	}

	// initial Flow struct
	a := rlib.Flow{
		BID:       BID,
		FlowID:    0, // it's new flowID,
		FlowType:  rlib.RAFlow,
		Data:      raflowJSONData,
		CreateBy:  UID,
		LastModBy: UID,
	}

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

// SvcGetRentableFeesData generates a list of rentable fees with auto populate AR fees
// wsdoc {
//  @Title Get list of Rentable fees with auto populate AR fees
//  @URL /v1/raflow-rentable-fees/:BUI/
//  @Method  GET
//  @Synopsis Get Rentable Fees list
//  @Description Get all rentable fees with auto populate AR fees
//  @Input RARentableFeesDataRequest
//  @Response RARentableFeesDataListResponse
// wsdoc }
func SvcGetRentableFeesData(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcGetRentableFeesData"
	var (
		g           FlowResponse
		rfd         RARentablesFlowData
		raflowData  RAFlowJSONData
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
	err = json.Unmarshal(flow.Data, &raflowData)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// find this RID in flow data rentable list
	var rIndex = -1
	for i := range raflowData.Rentables {
		if raflowData.Rentables[i].RID == rfd.RID {
			rIndex = i
		}
	}

	// if record not found then push it in the list
	if rIndex < 0 {
		raflowData.Rentables = append(raflowData.Rentables, rfd)
	} else {
		raflowData.Rentables[rIndex] = rfd
	}

	modRData, err := json.Marshal(&raflowData.Rentables)
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
	Recid  int64 `json:"recid"`
	BID    int64
	BUD    string
	FlowID int64
}
