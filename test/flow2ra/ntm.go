package main

import (
	"context"
	"encoding/json"
	"fmt"
	"rentroll/rlib"
	"rentroll/ws"
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
func setToNoticeToMove(ctx context.Context, s *rlib.Session, raid int64) error {
	var raf rlib.RAFlowJSONData

	rlib.Console("\tsetToNoticeToMove on RAID = %d\n", raid)
	//--------------------------------------
	// Make a Flow out of one of the RAs
	//--------------------------------------
	ra, err := rlib.GetRentalAgreement(ctx, raid)
	if err != nil {
		fmt.Printf("Could not read RentalAgreement: %s\n", err.Error())
		return err
	}

	//--------------------------------------
	// Insert new flow
	//--------------------------------------
	flowID, err := ws.GetRA2FlowCore(ctx, &ra, s.UID)
	if err != nil {
		rlib.Console("DoExistingRA - CB.err\n")
		fmt.Printf("Could not get Flow for RAID = %d: %s\n", ra.RAID, err.Error())
		return err
	}

	//-------------------------------------------
	// Read the flow data into a data structure
	//-------------------------------------------
	flow, err := rlib.GetFlow(ctx, flowID)
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

	rlib.Console("raf.Meta:  RAFLAGS=%d\n", raf.Meta.RAFLAGS)

	//--------------------------------------------
	// update meta info
	//--------------------------------------------
	var d []byte
	if d, err = json.Marshal(&raf.Meta); err != nil {
		return err
	}

	//----------------------------------------------------------------------
	// now call Flow2RA and make sure that it is updated to Notice-To-Move
	//----------------------------------------------------------------------
	tx, tctx, err := rlib.NewTransactionWithContext(ctx)
	if err != nil {
		fmt.Printf("Could not create transaction context: %s\n", err.Error())
		return err
	}

	// flow data with the updates to the meta fields
	if err = rlib.UpdateFlowPartData(ctx, "meta", d, &flow); err != nil {
		return err
	}

	var f2raid int64
	f2raid, err = ws.Flow2RA(tctx, flow.FlowID)
	if err != nil {
		tx.Rollback()
		rlib.Console("Flow2RA error\n")
		fmt.Printf("Could not write Flow back to db: %s\n", err.Error())
		return err
	}

	rlib.Console("Flow2RA update raid = %d\n", f2raid)

	if err = rlib.DeleteFlow(ctx, flowID); err != nil {
		fmt.Printf("Error deleting flow: %s\n", err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		fmt.Printf("Error committing transaction: %s\n", err.Error())
		return err
	}

	return nil
}
