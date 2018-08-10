package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"sort"
	"time"
)

// RARentableDetailsRequest is struct for request for rentable fees
type RARentableDetailsRequest struct {
	RID    int64
	FlowID int64
}

// SvcRAFlowRentableHandler handles operation on rentable of raflow json data
//           0    1     2   3
// uri /v1/raflow-rentable/BID/flowID
// The server command can be:
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcRAFlowRentableHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRAFlowRentableHandler"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "save":
		SaveRAFlowRentableDetails(w, r, d)
		break
	case "delete":
		DeleteRAFlowRentable(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// SaveRAFlowRentableDetails save Rentable and generates a list of rentable fees with auto populate AR fees
// It modifies raflow json doc by writing Fees data to raflow "rentables" component data
// wsdoc {
//  @Title Saves rentable with list of Rentable fees with auto populate AR fees
//  @URL /v1/raflow-rentable/:BUI/
//  @Method  GET
//  @Synopsis Save Rentable with Fees list
//  @Description Save rentable with all fees with auto populate AR fees
//  @Input RARentableDetailsRequest
//  @Response FlowResponse
// wsdoc }
func SaveRAFlowRentableDetails(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SaveRAFlowRentableDetails"
	var (
		rfd         rlib.RARentablesFlowData
		raFlowData  rlib.RAFlowJSONData
		foo         RARentableDetailsRequest
		feesRecords = []rlib.RAFeesData{}
		today       = time.Now()
		err         error
		tx          *sql.Tx
		ctx         context.Context
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
			raFlowData.Meta.LastTMPASMID++
			rec := rlib.RAFeesData{
				TMPASMID:       raFlowData.Meta.LastTMPASMID,
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

		raFlowData.Meta.LastTMPASMID++
		rec := rlib.RAFeesData{
			TMPASMID:       raFlowData.Meta.LastTMPASMID,
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

	// ----------------------------------------------
	// SYNC RECORDS IN OTHER SECTIONS
	// ----------------------------------------------
	// SYNC TIE RECORDS ON CHANGE OF PEOPLE
	rlib.SyncTieRecords(&raFlowData)

	// SYNC PARENT CHILD RECORDS ON CHANGE OF PEOPLE
	rlib.SyncParentChildRecords(&raFlowData)

	// LOOK FOR DATA CHANGES
	var originData rlib.RAFlowJSONData
	err = json.Unmarshal(flow.Data, &originData)
	if err != nil {
		return
	}

	// IF THERE ARE DATA CHANGES THEN ONLY UPDATE THE FLOW
	if !reflect.DeepEqual(originData, raFlowData) {
		// GET JSON DATA FROM THE STRUCT
		var modFlowData []byte
		modFlowData, err = json.Marshal(&raFlowData)
		if err != nil {
			return
		}

		// ASSIGN JSON MARSHALLED MODIFIED DATA
		flow.Data = modFlowData

		// NOW UPDATE THE WHOLE FLOW
		err = rlib.UpdateFlowWithInitState(ctx, &flow)
		if err != nil {
			return
		}

		// get the modified flow
		flow, err = rlib.GetFlow(ctx, flow.FlowID)
		if err != nil {
			return
		}
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err = tx.Commit(); err != nil {
		return
	}

	// -------------------
	// WRITE FLOW RESPONSE
	// -------------------
	SvcWriteFlowResponse(ctx, d.BID, flow, w)
	return
}

// RAFlowDeleteRentableRequest is struct for request to remove person from json data
type RAFlowDeleteRentableRequest struct {
	RID    int64
	FlowID int64
}

// DeleteRAFlowRentable remove rentable and syncs the records in parent/child, tie sections
// wsdoc {
//  @Title Remove Rentable entry
//  @URL /v1/raflow-rentable/:BUI/:FlowID
//  @Method POST
//  @Synopsis Remove Rentable from RAFlow json data
//  @Description Remove Rentable from RAFlow json data
//  @Input RAFlowDeleteRentableRequest
//  @Response FlowResponse
// wsdoc }
func DeleteRAFlowRentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "DeleteRAFlowRentable"
	var (
		raFlowData rlib.RAFlowJSONData
		foo        RAFlowDeleteRentableRequest
		err        error
		tx         *sql.Tx
		ctx        context.Context
	)
	fmt.Printf("Entered %s\n", funcname)

	// ===============================================
	// defer function to handle transactaion rollback
	// ===============================================
	defer func() {
		if err != nil {
			if tx != nil {
				tx.Rollback()
			}
			SvcErrorReturn(w, err, funcname)
			return
		}
	}()

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("Only POST method is allowed")
		return
	}

	// unmarshal data into request data struct
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		return
	}

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err = rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		return
	}

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

	// ----------------------------------------------
	// REMOVE ASSOCIATED PETS
	// ----------------------------------------------
	for i := range raFlowData.Rentables {
		if raFlowData.Rentables[i].RID == foo.RID {
			// remove this pet from the list
			raFlowData.Rentables = append(raFlowData.Rentables[:i], raFlowData.Rentables[i+1:]...)

			break
		}
	}

	// ----------------------------------------------
	// SYNC RECORDS IN OTHER SECTIONS
	// ----------------------------------------------
	// SYNC TIE RECORDS ON CHANGE OF PEOPLE
	rlib.SyncTieRecords(&raFlowData)

	// SYNC PARENT CHILD RECORDS ON CHANGE OF PEOPLE
	rlib.SyncParentChildRecords(&raFlowData)

	// LOOK FOR DATA CHANGES
	var originData rlib.RAFlowJSONData
	err = json.Unmarshal(flow.Data, &originData)
	if err != nil {
		return
	}

	// IF THERE ARE DATA CHANGES THEN ONLY UPDATE THE FLOW
	if !reflect.DeepEqual(originData, raFlowData) {
		// GET JSON DATA FROM THE STRUCT
		var modFlowData []byte
		modFlowData, err = json.Marshal(&raFlowData)
		if err != nil {
			return
		}

		// ASSIGN JSON MARSHALLED MODIFIED DATA
		flow.Data = modFlowData

		// NOW UPDATE THE WHOLE FLOW
		err = rlib.UpdateFlowWithInitState(ctx, &flow)
		if err != nil {
			return
		}

		// get the modified flow
		flow, err = rlib.GetFlow(ctx, flow.FlowID)
		if err != nil {
			return
		}
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err = tx.Commit(); err != nil {
		return
	}

	// -------------------
	// WRITE FLOW RESPONSE
	// -------------------
	SvcWriteFlowResponse(ctx, d.BID, flow, w)
	return
}
