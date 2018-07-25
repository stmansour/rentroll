package ws

import (
	"rentroll/bizlogic"
	"rentroll/rlib"
)

// GridRAFlowResponse is a struct to hold info for rental agreement for the grid response
type GridRAFlowResponse struct {
	Recid     int64 `json:"recid"`
	BID       int64
	BUD       string
	FlowID    int64
	UserRefNo string
}

// RAFlowResponse is a struct to hold info for flow information and relative basic validation check
type RAFlowResponse struct {
	Flow          rlib.Flow
	BasicCheck    bizlogic.ValidateRAFlowResponse
	DataFulfilled rlib.RADataFulfilled
}
