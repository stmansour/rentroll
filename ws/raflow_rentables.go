package ws

import (
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
		g              FlowResponse
		raFlowResponse RAFlowResponse
		rfd            rlib.RARentablesFlowData
		raFlowData     rlib.RAFlowJSONData
		foo            RARentableFeesDataRequest
		feesRecords    = []rlib.RAFeesData{}
		today          = time.Now()
		err            error
		tx             *sql.Tx
		ctx            context.Context
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

	// HTTP METHOD CHECK
	if r.Method != "POST" {
		err := fmt.Errorf("Only POST method is allowed")
		SvcErrorReturn(w, err, funcname)
		return
	}

	// SEE IF WE CAN UNMARSHAL THE DATA
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err = rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		return
	}

	//-------------------------------------------------------
	// FLOW EXISTENCE CHECK
	//-------------------------------------------------------
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

	// get meta in modRAFlowMeta, we're going to modify it
	modRAFlowMeta := raFlowData.Meta

	//-------------------------------------------------------
	// FIND RENTABLE AND RENTABLETYPE FROM REQUEST RID
	//-------------------------------------------------------
	var rentable rlib.Rentable
	rentable, err = rlib.GetRentable(ctx, foo.RID)
	if err != nil {
		return
	}

	// get rentableType
	var rtid int64
	rtid, err = rlib.GetRTIDForDate(ctx, foo.RID, &today)
	if err != nil {
		return
	}

	var rt rlib.RentableType
	err = rlib.GetRentableType(ctx, rtid, &rt)
	if err != nil {
		return
	}

	//-------------------------------------------------------
	// GET ACCOUNT RULE ASSOCIATED WITH FOUND RENTABLE TYPE
	// AND APPEND IT'S FEES IN RECORD LIST
	//-------------------------------------------------------
	// now get account rule based on this rentabletype
	var ar rlib.AR
	ar, _ = rlib.GetAR(ctx, rt.ARID)

	if ar.ARID > 0 {
		// make sure the IsRentASM is marked true
		if ar.FLAGS&0x10 != 0 {
			modRAFlowMeta.LastTMPASMID++
			rec := rlib.RAFeesData{
				TMPASMID:       modRAFlowMeta.LastTMPASMID,
				ARID:           ar.ARID,
				ARName:         ar.Name,
				ContractAmount: ar.DefaultAmount,
				Start:          rlib.JSONDate(today),
				Stop:           rlib.JSONDate(today.AddDate(1, 0, 0)),
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

	//-------------------------------------------------------
	// GET ALL AUTO POPULATED ACCOUNT RULES
	// APPEND FEES IN THE LIST EXCEPT RENTASM ONE AS WE
	// FOUND THAT PREVIOUSLY
	//-------------------------------------------------------
	// get all auto populated to new RA marked account rules by integer representation
	var m []rlib.AR
	arFLAGVal := 1 << uint64(bizlogic.ARFLAGS["AutoPopulateToNewRA"])
	m, err = rlib.GetARsByFLAGS(ctx, d.BID, uint64(arFLAGVal))
	if err != nil {
		return
	}

	// append feesRecords in ascending order
	for _, ar := range m {
		if ar.FLAGS&(1<<uint64(bizlogic.ARFLAGS["IsRentASM"])) != 0 { /*|| // if it's rent asm then continue
			ar.FLAGS&(1<<uint64(bizlogic.ARFLAGS["PETIDReq"])) != 0 || // if it's pet related AR
			ar.FLAGS&(1<<uint64(bizlogic.ARFLAGS["VIDReq"])) != 0 { */ // if it's vehicle related AR
			continue
		}

		modRAFlowMeta.LastTMPASMID++
		rec := rlib.RAFeesData{
			TMPASMID:       modRAFlowMeta.LastTMPASMID,
			ARID:           ar.ARID,
			ARName:         ar.Name,
			ContractAmount: ar.DefaultAmount,
			Start:          rlib.JSONDate(today),
			Stop:           rlib.JSONDate(today.AddDate(1, 0, 0)),
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

	//-------------------------------------------------------
	// NOW SORT THE FEES LIST BASED ON ARNAME
	// AND INSERT IT IN RENTABLE DATA
	//-------------------------------------------------------
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

	//-------------------------------------------------------
	// MODIFY RENTABLE JSON DATA IN RAFLOW
	//-------------------------------------------------------
	var modRData []byte
	modRData, err = json.Marshal(&raFlowData.Rentables)
	if err != nil {
		return
	}

	// update flow with this modified rentable part
	err = rlib.UpdateFlowData(ctx, "rentables", modRData, &flow)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//-------------------------------------------------------
	// MODIFY META DATA TOO
	//-------------------------------------------------------
	if raFlowData.Meta.LastTMPASMID < modRAFlowMeta.LastTMPASMID {
		var modMetaData []byte
		modMetaData, err = json.Marshal(&modRAFlowMeta)
		if err != nil {
			return
		}

		err = rlib.UpdateFlowData(ctx, "meta", modMetaData, &flow)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
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

	// get unmarshalled raflow data into struct
	err = json.Unmarshal(flow.Data, &raFlowData)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// Perform basic validation on flow data
	bizlogic.ValidateRAFlowBasic(r.Context(), &raFlowData, &raFlowResponse.BasicCheck)

	raFlowResponse.Flow = flow
	// set the response
	g.Record = raFlowResponse
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
