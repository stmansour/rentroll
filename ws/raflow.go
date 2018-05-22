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

// rental agreement flow part types
var raFlowPartTypes = rlib.Str2Int64Map{
	"dates":     int64(rlib.DatesRAFlowPart),
	"people":    int64(rlib.PeopleRAFlowPart),
	"pets":      int64(rlib.PetsRAFlowPart),
	"vehicles":  int64(rlib.VehiclesRAFlowPart),
	"rentables": int64(rlib.RentablesRAFlowPart),
	"feesterms": int64(rlib.FeesTermsRAFlowPart),
}

// RAFlowJSONData holds the struct for all the parts being involed in rental agreement flow
type RAFlowJSONData struct {
	RADatesFlowData     `json:"dates"`
	RAPeopleFlowData    `json:"people"`
	RAPetsFlowData      `json:"pets"`
	RAVehiclesFlowData  `json:"vehicles"`
	RARentablesFlowData `json:"rentables"`
	RAFeesTermsFlowData `json:"feesterms"`
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

// RAPetsFlowData contains data in the pets part of RA flow
type RAPetsFlowData struct {
	// Recid                int           `json:"recid"` // this is for the grid widget
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
	// Recid               int           `json:"recid"` // this is for the grid widget
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

// RAPeopleFlowData contains data in the background-info part of RA flow
type RAPeopleFlowData struct {
	// Recid int64 `json:"recid"` // this is for the grid widget
	BID  int64
	TCID int64

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
	Evicted    bool // Evicted
	Convicted  bool // Arrested or convicted of a Convicted
	Bankruptcy bool // Declared Bankruptcy

	// Emergency contact information
	EmergencyContactName    string
	EmergencyContactPhone   string
	EmergencyContactAddress string

	// RA Application information
	Comment string // In an effort to accommodate you, please advise us of any special needs
}

// RARentablesFlowData contains data in the rentables part of RA flow
type RARentablesFlowData struct {
	// Recid        int     `json:"recid"` // this is for the grid widget
	BID          int64
	RID          int64
	RTID         int64
	RentableName string
	ContractRent float64
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
	Amount          float64
	RentCycle       int64
	Epoch           int64
	RentPeriodStart rlib.JSONDate
	RentPeriodStop  rlib.JSONDate
	UsePeriodStart  rlib.JSONDate
	UsePeriodStop   rlib.JSONDate
	ContractRent    float64
	ProrateAmt      float64
	SalesTaxAmt     float64
	SalesTax        float64
	TransOccAmt     float64
	TransOcc        float64
}

// RAFeesTermsFlowData contains data in the fees-terms part of RA flow
type RAFeesTermsFlowData struct {
	// Recid        int     `json:"recid"` // this is for the grid widget
	BID          int64
	RID          int64
	RTID         int64
	RentableName string
	FeeName      string
	Amount       float64
	Cycle        float64
	SigningAmt   float64
	ProrateAmt   float64
	TaxableAmt   float64
	SalesTax     float64
	TransOcc     float64
}

// getUpdateRAFlowPartJSONData returns json data in bytes
// coming from client with checking of flow and part type to update
func getUpdateRAFlowPartJSONData(BID int64, data json.RawMessage, partType int) ([]byte, error) {

	// TODO: Add validation on field level, it must be done.

	// JSON Marshal with address
	// REF: https://stackoverflow.com/questions/21390979/custom-marshaljson-never-gets-called-in-go

	// is it blank string or null json data
	isBlankJSONData := bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))

	switch rlib.RAFlowPartType(partType) {
	case rlib.DatesRAFlowPart:
		a := RADatesFlowData{}

		// if the struct provided with some data then checks for it
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				// if it's an error then return with nil data
				return []byte(nil), err
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
		// return json marshalled for struct
		return json.Marshal(&a)

	case rlib.PeopleRAFlowPart:
		a := []RAPeopleFlowData{}

		// if the struct provided with some data then checks for it
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				// if it's an error then return with nil data
				return []byte(nil), err
			}
		}
		// return json marshalled for struct
		return json.Marshal(&a)

	case rlib.PetsRAFlowPart:
		a := []RAPetsFlowData{}

		// if the struct provided with some data then checks for it
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				// if it's an error then return with nil data
				return []byte(nil), err
			}
		}
		// return json marshalled for struct
		return json.Marshal(&a)

	case rlib.VehiclesRAFlowPart:
		a := []RAVehiclesFlowData{}

		// if the struct provided with some data then checks for it
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				// if it's an error then return with nil data
				return []byte(nil), err
			}
		}
		// return json marshalled for struct
		return json.Marshal(&a)

	case rlib.RentablesRAFlowPart:
		a := []RARentablesFlowData{}

		// if the struct provided with some data then checks for it
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				// if it's an error then return with nil data
				return []byte(nil), err
			}
		}
		// return json marshalled for struct
		return json.Marshal(&a)

	case rlib.FeesTermsRAFlowPart:
		a := []RAFeesTermsFlowData{}

		// if the struct provided with some data then checks for it
		// json validation
		if !(isBlankJSONData) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				// if it's an error then return with nil data
				return []byte(nil), err
			}
		}
		// return json marshalled for struct
		return json.Marshal(&a)

	default:
		// not valid option then return with nil data
		return []byte(nil), fmt.Errorf("unrecognized part type in RA flow: %d", partType)
	}
}

// insertInitialRAFlow writes a bunch of flow's sections record for a particular RA
// This should be run under atomic transaction mode as per DB design of flow
// This is very special case that we're not returning primary key generated from database
// instead we're generating in form of string which we return if tx will be succeed.
func insertInitialRAFlow(ctx context.Context, BID, UID int64) (string, error) {

	var (
		flowID string
		err    error
		ok     bool
	)

	// ------------
	// SPECIAL CASE
	// ------------
	var (
		newTx bool
		tx    *sql.Tx
	)

	if tx, ok = rlib.DBTxFromContext(ctx); !ok { // if transaction is NOT supplied
		newTx = true
		tx, err = rlib.RRdb.Dbrr.Begin()
		if err != nil {
			return flowID, err
		}
		ctx = rlib.SetDBTxContextKey(ctx, tx)
	}

	// getFlowID first
	flowID = rlib.GetFlowID()

	// initRAFlowPart
	initRAFlowPart := rlib.FlowPart{
		BID:       BID,
		Flow:      rlib.RAFlow,
		FlowID:    flowID,
		PartType:  0,
		Data:      json.RawMessage([]byte("null")), // JSON "null" primitive type
		CreateBy:  UID,
		LastModBy: UID,
	}

	// Rental agreement flow parts map init
	// maybe we can just override the above pre-defined initFlowPart struct
	initRAFlowMap := map[rlib.RAFlowPartType]rlib.FlowPart{
		rlib.DatesRAFlowPart:     initRAFlowPart,
		rlib.PeopleRAFlowPart:    initRAFlowPart,
		rlib.PetsRAFlowPart:      initRAFlowPart,
		rlib.VehiclesRAFlowPart:  initRAFlowPart,
		rlib.RentablesRAFlowPart: initRAFlowPart,
		rlib.FeesTermsRAFlowPart: initRAFlowPart,
	}

	// insert in order to ease
	var keys rlib.Int64Range
	for k := range initRAFlowMap {
		keys = append(keys, int64(k))
	}
	sort.Sort(keys)

	// assign part type
	for _, partTypeIDi64 := range keys {

		// get blank flow part
		a := initRAFlowMap[rlib.RAFlowPartType(partTypeIDi64)]

		// modify part type
		a.PartType = int(partTypeIDi64)

		// get json strctured data from go struct and feed it back into a Data field
		a.Data, _ = getUpdateRAFlowPartJSONData(BID, a.Data, a.PartType)

		// insert each flowpart of RA flow
		_, err = rlib.InsertFlowPart(ctx, &a)
		if err != nil {
			rlib.Ulog("Error while inserting FlowPart BULK-WRITE: %s\n", err.Error())
		}
	}

	if newTx { // if new transaction then commit it
		// if error then rollback
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			rlib.Ulog("Error while Committing transaction | inserting FlowPart BULK-WRITE: %s\n", err.Error())
			// err = insertError(err, "InitialRAFlow", nil)
			return flowID, err
		}
	}

	return flowID, err
}

// RARentableFeesDataRequest is struct for request for rentable fees
type RARentableFeesDataRequest struct {
	RID int64
}

// RARentableFeesDataListResponse for listing down all RARentableFeesData
// in the grid
type RARentableFeesDataListResponse struct {
	Status  string               `json:"status"`
	Total   int64                `json:"total"`
	Records []RARentableFeesData `json:"records"`
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
		g       RARentableFeesDataListResponse
		foo     RARentableFeesDataRequest
		records []RARentableFeesData
		today   = time.Now()
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
				Amount:          ar.DefaultAmount,
				ContractRent:    ar.DefaultAmount,
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

			records = append(records, rec)
		}
	}

	// get all auto populated to new RA marked account rules by integer representation
	arFLAGVal := 1 << uint64(bizlogic.ARFLAGS["AutoPopulateToNewRA"])
	m, err := rlib.GetARsByFLAGS(r.Context(), d.BID, uint64(arFLAGVal))
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// append records in ascending order
	for _, ar := range m {
		if ar.FLAGS&0x10 != 0 { // if it's rent asm then continue
			continue
		}

		rec := RARentableFeesData{
			BID:             ar.BID,
			ARID:            ar.ARID,
			RID:             foo.RID,
			ARName:          ar.Name,
			Amount:          ar.DefaultAmount,
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

		// now append rec in records
		records = append(records, rec)
	}

	// sort based on name, needs version 1.8 later of golang
	sort.Slice(records, func(i, j int) bool { return records[i].ARName < records[j].ARName })

	g.Records = records
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// saveRentalAgreementFlow saves data for the given flowID to real multi variant database instances
// from the temporary data stored in FlowPart table
func saveRentalAgreementFlow(ctx context.Context, flowID string) error {
	var (
		RAID int64
		err  error
	)

	// first check that such a given flowID does exist or not
	var found bool
	ids, err := rlib.GetFlowIDsByUser(ctx, rlib.RAFlow)
	if err != nil {
		return err
	}

	for _, id := range ids {
		if id == flowID {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Such flowID: %s does not exist", flowID)
	}

	// -------------- SAVING PARTS --------------------

	// ==================
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
	fmt.Printf("Newly created rental agreement with RAID: %d\n", RAID)

	return nil
}

// GridRAFlowResponse is a struct to hold info for rental agreement for the grid response
type GridRAFlowResponse struct {
	Recid  int64 `json:"recid"`
	BID    int64
	BUD    string
	FlowID string
}
