package ws

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
	"bginfo":    int64(rlib.BackGroundInfoRAFlowPart),
	"rentables": int64(rlib.RentablesRAFlowPart),
	"feesterms": int64(rlib.FeesTermsRAFlowPart),
}

/*// RAFlowJSONData holds the struct for all the parts being involed in rental agreement flow
type RAFlowJSONData struct {
	RADatesFlowData          `json:"dates"`
	RAPeopleFlowData         `json:"people"`
	RAPetsFlowData           `json:"pets"`
	RAVehiclesFlowData       `json:"vehicles"`
	RABackgroundInfoFlowData `json:"bginfo"`
	RARentablesFlowData      `json:"rentables"`
	RAFeesTermsFlowData      `json:"feesterms"`
}*/

// RADatesFlowData contains data in the dates part of RA flow
type RADatesFlowData struct {
	AgreementStart  rlib.JSONDate `json:"AgreementStart"` // TermStart
	AgreementStop   rlib.JSONDate `json:"AgreementStop"`  // TermStop
	RentStart       rlib.JSONDate `json:"RentStart"`
	RentStop        rlib.JSONDate `json:"RentStop"`
	PossessionStart rlib.JSONDate `json:"PossessionStart"`
	PossessionStop  rlib.JSONDate `json:"PossessionStop"`
}

// RAPeopleFlowData contains data in the people part of RA flow
type RAPeopleFlowData struct {
	Payors     []rlib.TransactantTypeDown `json:"Payors"`
	Users      []rlib.TransactantTypeDown `json:"Users"`
	Guarantors []rlib.TransactantTypeDown `json:"Guarantors"`
}

// RAPetsFlowData contains data in the pets part of RA flow
type RAPetsFlowData struct {
	PETID                int64         `json:"PETID"`
	BID                  int64         `json:"BID"`
	Name                 string        `json:"Name"`
	Type                 string        `json:"Type"`
	Breed                string        `json:"Breed"`
	Coloe                string        `json:"Coloe"`
	Weight               int           `json:"Weight"`
	DtStart              rlib.JSONDate `json:"DtStart"`
	DtStop               rlib.JSONDate `json:"DtStop"`
	NonRefundablePetFee  float64       `json:"NonRefundablePetFee"`
	RefundablePetDeposit float64       `json:"RefundablePetDeposit"`
	RecurringPetFee      float64       `json:"RecurringPetFee"`
}

// RAVehiclesFlowData contains data in the vehicles part of RA flow
type RAVehiclesFlowData struct {
	VID                 int64         `json:"VID"`
	BID                 int64         `json:"BID"`
	TCID                int64         `json:"TCID"`
	VIN                 string        `json:"VIN"`
	Type                string        `json:"Type"`
	Make                string        `json:"Make"`
	Model               string        `json:"Model"`
	Color               string        `json:"Color"`
	LicensePlateState   string        `json:"LicensePlateState"`
	LicensePlateNumber  string        `json:"LicensePlateNumber"`
	ParkingPermitNumber string        `json:"ParkingPermitNumber"`
	DtStart             rlib.JSONDate `json:"DtStart"`
	DtStop              rlib.JSONDate `json:"DtStop"`
}

// RABackgroundInfoFlowData contains data in the background-info part of RA flow
type RABackgroundInfoFlowData struct {
	Applicant string `json:"Applicant"`
}

// RARentablesFlowData contains data in the rentables part of RA flow
type RARentablesFlowData struct {
	RID          int64   `json:"RID"`
	BID          int64   `json:"BID"`
	RTID         int64   `json:"RTID"`
	RentableName string  `json:"RentableName"`
	ContractRent float64 `json:"ContractRent"`
	ProrateAmt   float64 `json:"ProrateAmt"`
	TaxableAmt   float64 `json:"TaxableAmt"`
	SalesTax     float64 `json:"SalesTax"`
	TransOCC     float64 `json:"TransOCC"`
}

// RAFeesTermsFlowData contains data in the fees-terms part of RA flow
type RAFeesTermsFlowData struct {
	RID          int64   `json:"RID"`
	BID          int64   `json:"BID"`
	RTID         int64   `json:"RTID"`
	RentableName string  `json:"RentableName"`
	FeeName      string  `json:"FeeName"`
	Amount       float64 `json:"Amount"`
	Cycle        float64 `json:"Cycle"`
	SigningAmt   float64 `json:"SigningAmt"`
	ProrateAmt   float64 `json:"ProrateAmt"`
	TaxableAmt   float64 `json:"TaxableAmt"`
	SalesTax     float64 `json:"SalesTax"`
	TransOCC     float64 `json:"TransOCC"`
}

// getUpdateRAFlowPartJSONData returns json data in bytes
// coming from client with checking of flow and part type to update
func getUpdateRAFlowPartJSONData(data json.RawMessage, partType int) ([]byte, error) {

	// TODO: Add validation on field level, it must be done.

	// JSON Marshal with address
	// REF: https://stackoverflow.com/questions/21390979/custom-marshaljson-never-gets-called-in-go

	currentDateTime := time.Now()
	nextYearDateTime := currentDateTime.AddDate(1, 0, 0)

	switch rlib.RAFlowPartType(partType) {
	case rlib.DatesRAFlowPart:
		a := RADatesFlowData{
			RentStart:       rlib.JSONDate(currentDateTime),
			RentStop:        rlib.JSONDate(nextYearDateTime),
			AgreementStart:  rlib.JSONDate(currentDateTime),
			AgreementStop:   rlib.JSONDate(nextYearDateTime),
			PossessionStart: rlib.JSONDate(currentDateTime),
			PossessionStop:  rlib.JSONDate(nextYearDateTime),
		}
		if !(bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				return []byte(nil), err
			}
		}
		return json.Marshal(&a)
	case rlib.PeopleRAFlowPart:
		a := RAPeopleFlowData{
			Payors:     []rlib.TransactantTypeDown{},
			Users:      []rlib.TransactantTypeDown{},
			Guarantors: []rlib.TransactantTypeDown{},
		}
		if !(bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				return []byte(nil), err
			}
		}
		return json.Marshal(&a)
	case rlib.PetsRAFlowPart:
		a := []RAPetsFlowData{}
		if !(bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				return []byte(nil), err
			}
		}
		return json.Marshal(&a)
	case rlib.VehiclesRAFlowPart:
		a := []RAVehiclesFlowData{}
		if !(bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				return []byte(nil), err
			}
		}
		return json.Marshal(&a)
	case rlib.BackGroundInfoRAFlowPart:
		var a RABackgroundInfoFlowData
		if !(bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				return []byte(nil), err
			}
		}
		return json.Marshal(&a)
	case rlib.RentablesRAFlowPart:
		a := []RARentablesFlowData{}
		if !(bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				return []byte(nil), err
			}
		}
		return json.Marshal(&a)
	case rlib.FeesTermsRAFlowPart:
		a := []RAFeesTermsFlowData{}
		if !(bytes.Equal([]byte(data), []byte(``)) || bytes.Equal([]byte(data), []byte(`null`))) {
			err := json.Unmarshal(data, &a)
			if err != nil {
				return []byte(nil), err
			}
		}
		return json.Marshal(&a)
	default:
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
	flowID = rlib.GetFlowID(UID)

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
		rlib.DatesRAFlowPart:          rlib.FlowPart{},
		rlib.PeopleRAFlowPart:         rlib.FlowPart{},
		rlib.PetsRAFlowPart:           rlib.FlowPart{},
		rlib.VehiclesRAFlowPart:       rlib.FlowPart{},
		rlib.BackGroundInfoRAFlowPart: rlib.FlowPart{},
		rlib.RentablesRAFlowPart:      rlib.FlowPart{},
		rlib.FeesTermsRAFlowPart:      rlib.FlowPart{},
	}

	// insert in order to ease
	var keys rlib.Int64Range
	for k := range initRAFlowMap {
		keys = append(keys, int64(k))
	}
	sort.Sort(keys)

	// assign part type
	for _, v := range keys {
		partTypeID := rlib.RAFlowPartType(v)
		// fmt.Printf("partTypeID: %s: %d\n", partTypeID, partTypeID)

		// get blank flow part
		a := initRAFlowMap[rlib.RAFlowPartType(partTypeID)]

		// assign pre-defined init flow data
		a = initRAFlowPart

		// modify part type
		a.PartType = int(partTypeID)

		// get json strctured data from go struct
		a.Data, _ = getUpdateRAFlowPartJSONData(a.Data, a.PartType)

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
	fmt.Println("list of flowIds", ids)
	fmt.Println(flowID)

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
