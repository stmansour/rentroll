package main

import (
	"context"
	"encoding/json"
	"rentroll/rlib"
	"time"
)

// setToNoticeToMove handles all the updates necessary to move the fees defined in a flow
// into the permanent tables.
//
// INPUTS
//     ctx    - db context for transactions
//     x - all the contextual info we need for performing this operation
//
// RETURNS
//     Any errors encountered
//-----------------------------------------------------------------------------
func setToNoticeToMove(ctx context.Context, flowid int64) error {
	var raf rlib.RAFlowJSONData
	//-------------------------------------------
	// Read the flow data into a data structure
	//-------------------------------------------
	flow, err := rlib.GetFlow(ctx, flowid)
	if err != nil {
		return err
	}
	err = json.Unmarshal(flow.Data, &raf)
	if err != nil {
		return err
	}

	dt := time.Time(raf.Dates.AgreementStart).AddDate(0, 2, 8) // 2 months and 8 days after the start of the agreement
	dtr := time.Time(raf.Dates.AgreementStart).AddDate(0, 1, 8)

	raf.Meta.RAFLAGS &= ^uint64(0xf)             // clear the first 4 bits
	raf.Meta.RAFLAGS |= rlib.RASTATENoticeToMove // make it NoticeToMove
	raf.Meta.NoticeToMoveUID = 287
	raf.Meta.NoticeToMoveDate = rlib.JSONDateTime(dt)
	raf.Meta.NoticeToMoveReported = rlib.JSONDateTime(dtr)

	//--------------------------------------------
	// update meta info
	//--------------------------------------------
	var d []byte
	if d, err = json.Marshal(&raf.Meta); err != nil {
		return err
	}
	if err = rlib.UpdateFlowData(ctx, "meta", d, &flow); err != nil {
		return err
	}

	return nil
}
