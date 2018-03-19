package ws

import "time"

// RADatesFlowData contains data in the dates part of RA flow
type RADatesFlowData struct {
	AgreementStart  time.Time `json:"AgreementStart"` // TermStart
	AgreementStop   time.Time `json:"AgreementStop"`  // TermStop
	RentStart       time.Time `json:"RentStart"`
	RentStop        time.Time `json:"RentStop"`
	PossessionStart time.Time `json:"PossessionStart"`
	PossessionStop  time.Time `json:"PossessionStop"`
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
