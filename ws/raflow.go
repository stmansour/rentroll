package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"rentroll/rlib"
	"time"
)

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
	Transactant string `json:"Transactant"`
	Payor       bool   `json:"Payor"`
	User        bool   `json:"User"`
	Guarantor   bool   `json:"Guarantor"`
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

	switch rlib.RAFlowPartType(partType) {
	case rlib.DatesRAFlowPart:
		var a RADatesFlowData
		err := json.Unmarshal(data, &a)
		if err != nil {
			return []byte(nil), err
		}
		return json.Marshal(&a)
	case rlib.PeopleRAFlowPart:
		var a RAPeopleFlowData
		err := json.Unmarshal(data, &a)
		if err != nil {
			return []byte(nil), err
		}
		return json.Marshal(&a)
	case rlib.PetsRAFlowPart:
		var a []RAPetsFlowData
		err := json.Unmarshal(data, &a)
		if err != nil {
			return []byte(nil), err
		}
		return json.Marshal(&a)
	case rlib.VehiclesRAFlowPart:
		var a []RAVehiclesFlowData
		err := json.Unmarshal(data, &a)
		if err != nil {
			return []byte(nil), err
		}
		return json.Marshal(&a)
	case rlib.BackGroundInfoRAFlowPart:
		var a RABackgroundInfoFlowData
		err := json.Unmarshal(data, &a)
		if err != nil {
			return []byte(nil), err
		}
		return json.Marshal(&a)
	case rlib.RentablesRAFlowPart:
		var a []RARentablesFlowData
		err := json.Unmarshal(data, &a)
		if err != nil {
			return []byte(nil), err
		}
		return json.Marshal(&a)
	case rlib.FeesTermsRAFlowPart:
		var a []RAFeesTermsFlowData
		err := json.Unmarshal(data, &a)
		if err != nil {
			return []byte(nil), err
		}
		return json.Marshal(&a)
	default:
		return []byte(nil), fmt.Errorf("unrecognized part type in RA flow: %d", partType)
	}
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
	ids, err := rlib.GetFlowIDsByUser(ctx, "RA")
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
