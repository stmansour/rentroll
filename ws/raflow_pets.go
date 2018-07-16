package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// NewRAFlowPet create new pet entry for the raflow and returns strcture
// with fees configured it in bizprops
func NewRAFlowPet(ctx context.Context, BID int64, meta *rlib.RAFlowMetaInfo) (pet rlib.RAPetsFlowData, err error) {
	const funcname = "NewRAFlowPet"
	var (
		today = time.Now()
	)
	fmt.Printf("Entered in %s\n", funcname)

	// initialize
	// assign new TMPPETID & mark in meta info
	meta.LastTMPPETID++
	pet = rlib.RAPetsFlowData{
		Fees:     []rlib.RAFeesData{},
		TMPPETID: meta.LastTMPPETID,
		DtStart:  rlib.JSONDate(today),
		DtStop:   rlib.JSONDate(today.AddDate(1, 0, 0)),
	}

	// get pet fees data and feed into fees
	var petFees []rlib.BizPropsPetFee
	petFees, err = rlib.GetPetFeesFromGeneralBizProps(ctx, BID)
	if err != nil {
		return
	}

	// loop over fees
	for _, fee := range petFees {
		meta.LastTMPASMID++ // new asm id temp
		pf := rlib.RAFeesData{
			ARID:           fee.ARID,
			ARName:         fee.ARName,
			ContractAmount: fee.Amount,
			TMPASMID:       meta.LastTMPASMID,
		}

		// append fee for this pet
		pet.Fees = append(pet.Fees, pf)
	}

	return
}

// SvcRAFlowPetsHandler handles operation on pets of raflow json data
//           0    1     2   3
// uri /v1/raflow-pets/BID/flowID
// The server command can be:
//      new
//-----------------------------------------------------------------------------------
func SvcRAFlowPetsHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRAFlowPetsHandler"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  FlowID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "new":
		CreateNewRAFlowPet(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// RAFlowNewPetRequest struct for new pet request
type RAFlowNewPetRequest struct {
	FlowID int64
}

// CreateNewRAFlowPet creates new entry in pets list of raflow json data
// wsdoc {
//  @Title create brand new pet entry in pet list
//  @URL /v1/raflow-pets/:BUI/:FlowID
//  @Method POST
//  @Synopsis create new entry of pet in RAFlow json data
//  @Description create new entry of pet in RAFlow json data with all available fees
//  @Input RAFlowNewPetRequest
//  @Response FlowResponse
// wsdoc }
//-----------------------------------------------------------------------------------
func CreateNewRAFlowPet(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "CreateNewRAFlowPet"
	var (
		g             FlowResponse
		foo           RAFlowNewPetRequest
		raFlowData    rlib.RAFlowJSONData
		err           error
		tx            *sql.Tx
		ctx           context.Context
		modRAFlowMeta rlib.RAFlowMetaInfo
	)
	fmt.Printf("Entered in %s\n", funcname)

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
	modRAFlowMeta = raFlowData.Meta

	// --------------------------------------------------------
	// APPEND FEES FOR PETS
	// --------------------------------------------------------
	var newRAFlowPet rlib.RAPetsFlowData
	newRAFlowPet, err = NewRAFlowPet(r.Context(), d.BID, &modRAFlowMeta)
	if err != nil {
		return
	}

	// append in pets list
	raFlowData.Pets = append(raFlowData.Pets, newRAFlowPet)

	// --------------------------------------------------------
	// UPDATE JSON DOC WITH NEW PET DATA (BLANK)
	// --------------------------------------------------------
	var modPetsData []byte
	modPetsData, err = json.Marshal(&raFlowData.Pets)
	if err != nil {
		return
	}

	// update flow with this modified pets part
	err = rlib.UpdateFlowData(ctx, "pets", modPetsData, &flow)
	if err != nil {
		return
	}

	// --------------------------------------------------------
	// UPDATE JSON DOC WITH NEW META DATA
	// --------------------------------------------------------
	if raFlowData.Meta.LastTMPASMID < modRAFlowMeta.LastTMPASMID {

		// Update HavePets Flag in meta information of flow
		modRAFlowMeta.HavePets = len(raFlowData.Pets) > 0

		// get marshalled data
		var modMetaData []byte
		modMetaData, err = json.Marshal(&modRAFlowMeta)
		if err != nil {
			return
		}

		// update flow with this modified meta part
		err = rlib.UpdateFlowData(ctx, "meta", modMetaData, &flow)
		if err != nil {
			return
		}
	}

	// ----------------------------------------------
	// RETURN RESPONSE
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

	// set the response
	g.Record = flow
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
