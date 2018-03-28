package ws

import (
	"encoding/json"
	"fmt"
	"rentroll/rlib"
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
}

// RAPetsFlowData contains data in the pets part of RA flow
type RAPetsFlowData struct {
}

// RAVehiclesFlowData contains data in the vehicles part of RA flow
type RAVehiclesFlowData struct {
}

// RABackgroundInfoFlowData contains data in the background-info part of RA flow
type RABackgroundInfoFlowData struct {
}

// RARentablesFlowData contains data in the rentables part of RA flow
type RARentablesFlowData struct {
}

// RAFeesTermsFlowData contains data in the fees-terms part of RA flow
type RAFeesTermsFlowData struct {
}

// isValidUpdateRAFlowPartJSONData handle data coming from client with checking
// of flow and part type to update
func isValidUpdateRAFlowPartJSONData(data json.RawMessage, partType int) bool {
	var (
		err error
		a   interface{}
	)

	// TODO: Add validation on field level, it must be done.

	switch rlib.RAFlowPartType(partType) {
	case rlib.DatesRAFlowPart:
		a = RADatesFlowData{}
	case rlib.PeopleRAFlowPart:
		a = RAPeopleFlowData{}
	case rlib.PetsRAFlowPart:
		a = RAPetsFlowData{}
	case rlib.VehiclesRAFlowPart:
		a = RAVehiclesFlowData{}
	case rlib.BackGroundInfoRAFlowPart:
		a = RABackgroundInfoFlowData{}
	case rlib.RentablesRAFlowPart:
		a = RARentablesFlowData{}
	case rlib.FeesTermsRAFlowPart:
		a = RAFeesTermsFlowData{}
	default:
		err = fmt.Errorf("unrecognized part type in RA flow: %d", partType)
	}

	// now try to load json data into the struct
	err = json.Unmarshal(data, &a)

	return err == nil
}
