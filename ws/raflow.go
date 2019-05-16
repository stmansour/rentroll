package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// GridRAFlowRecord is a struct to hold info for rental agreement for the grid response
type GridRAFlowRecord struct {
	Recid          int64           `json:"recid"` // this is to support the w2ui form
	RAID           rlib.NullInt64  // internal unique id
	BID            int64           // Business (so that we can process by Business)
	AgreementStart rlib.NullDate   // start date for rental agreement contract
	AgreementStop  rlib.NullDate   // stop date for rental agreement contract
	Payors         rlib.NullString // payors list attached with this RA within same time
	FlowID         rlib.NullInt64  // FlowID
	UserRefNo      rlib.NullString // FlowID - reference number
	FLAGS          rlib.NullInt64  // FLAGS
}

// GridRAFlowResponse is the response data for a Rental Agreement Search
type GridRAFlowResponse struct {
	Status  string             `json:"status"`
	Total   int64              `json:"total"`
	Records []GridRAFlowRecord `json:"records"`
}

// RAFlowResponse is a struct to hold info for flow information and relative basic validation check
type RAFlowResponse struct {
	Flow            rlib.Flow
	ValidationCheck bizlogic.ValidateRAFlowResponse
	DataFulfilled   rlib.RADataFulfilled
}

// RAFlowGridFieldsMap holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
var RAFlowGridFieldsMap = map[string][]string{
	"BID":            {"RA_CUM_FLOW.BID"},
	"RAID":           {"RA_CUM_FLOW.RAID"},
	"AgreementStart": {"RA_CUM_FLOW.AgreementStart"},
	"AgreementStop":  {"RA_CUM_FLOW.AgreementStop"},
	"Payors":         {"RA_CUM_FLOW.Payors"},
	"FlowID":         {"RA_CUM_FLOW.FlowID"},
	"UserRefNo":      {"RA_CUM_FLOW.UserRefNo"},
	"FLAGS":          {"RA_CUM_FLOW.FLAGS"},
}

// RAFlowQuerySelectFields defines which fields needs to be fetched for SQL query for rental agreements
var RAFlowQuerySelectFields = []string{
	"RA_CUM_FLOW.BID",
	"RA_CUM_FLOW.RAID",
	"RA_CUM_FLOW.AgreementStart",
	"RA_CUM_FLOW.AgreementStop",
	"RA_CUM_FLOW.Payors",
	"RA_CUM_FLOW.FlowID",
	"RA_CUM_FLOW.UserRefNo",
	"RA_CUM_FLOW.FLAGS",
}

// RAFlowListQuery to fetch all rental agreements and flow w or w/o inter-relation
var RAFlowListQuery = `
SELECT
    {{.SelectClause}}
FROM (
    /***********************************
    Rental Agreements (with or without flow)
    UNION Flow (without RA) Collection
    - - - - - - - - - - - - - - - - - */
    (
        /*
            COLLECT ALL RENTAL AGREEMENTS WITH ASSOCIATED FLOW ENTRIES
            GROUPED BY RAID AND THEN ORDER BY RAID ASC
        */
        SELECT
            RentalAgreement.BID AS BID,
            RentalAgreement.RAID AS RAID,
            RentalAgreement.AgreementStart AS AgreementStart,
            RentalAgreement.AgreementStop AS AgreementStop,
            GROUP_CONCAT(DISTINCT CASE WHEN Payor.IsCompany > 0 THEN Payor.CompanyName ELSE CONCAT(Payor.FirstName, ' ', Payor.LastName) END ORDER BY 1 SEPARATOR ', ') AS Payors,
            RentalAgreement.FLAGS AS FLAGS,
            Flow.FlowID AS FlowID,
            Flow.UserRefNo AS UserRefNo
        FROM RentalAgreement
        LEFT JOIN RentalAgreementPayors ON RentalAgreementPayors.RAID=RentalAgreement.RAID
        LEFT JOIN Transactant AS Payor ON Payor.TCID=RentalAgreementPayors.TCID
        LEFT JOIN Flow ON Flow.ID=RentalAgreement.RAID
        WHERE
            RentalAgreement.BID={{.BID}}
            AND RentalAgreement.AgreementStart < "{{.Stop}}"
            AND RentalAgreement.AgreementStop > "{{.Start}}"
--          {{.IncludeCancelled}}
        GROUP BY RentalAgreement.RAID
        ORDER BY Payors ASC, AgreementStart ASC
    )
    UNION ALL
    (
        /*
            COLLECT ALL FLOW ENTRIES WHICH ARE NOT ASSOCIATED WITH ANY RA
            IN ORDER BY FlowID ASC
        */
        SELECT
            Flow.BID AS BID,
            NULL AS RAID,
            NULL AS AgreementStart,
            NULL AS AgreementStop,
            NULL AS Payors,
            NULL AS FLAGS,
            Flow.FlowID AS FlowID,
            Flow.UserRefNo AS UserRefNo
        FROM Flow
        WHERE Flow.BID={{.BID}} AND Flow.ID=0 AND "{{.Start}}" <= Flow.CreateTS AND Flow.CreateTS < "{{.Stop}}"
        ORDER BY Flow.FlowID ASC
    )
    /*- - - - - - - - - - - - - - - - -
    Rental Agreements (with or without flow)
    UNION Flow (without RA) Collection
    ***********************************/
) RA_CUM_FLOW
WHERE {{.WhereClause}}
/* ORDER BY RAID, FlowID (if null then it would be last otherwise) */
ORDER BY {{.OrderClause}}
`

// RAFlowQueryClause contains query clauses for raflow query
var RAFlowQueryClause = rlib.QueryClause{
	"BID":              "",
	"Start":            "",
	"Stop":             "",
	"IncludeCancelled": "AND (RentalAgreement.FLAGS & 64) = 0", // I've commented it out above because we need to see it in the RentalAgreements grid, just not in the Rentroll report or view
	"SelectClause":     strings.Join(RAFlowQuerySelectFields, ","),
	"WhereClause":      "",
	"OrderClause":      "RA_CUM_FLOW.Payors ASC, RA_CUM_FLOW.AgreementStart ASC",
}

// GetAllRAFlows returns all existing Rental Agreements and all Flows
// are being edited for rental agreement.
func GetAllRAFlows(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "GetAllRAFlows"
	var (
		qc    = rlib.GetQueryClauseCopy(RAFlowQueryClause)
		srch  = fmt.Sprintf("RA_CUM_FLOW.BID=%d", d.BID)
		order = qc["OrderClause"]
		resp  = GridRAFlowResponse{
			Records: []GridRAFlowRecord{},
		}
		err error
	)
	fmt.Printf("Entered in %s\n", funcname)

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, RAFlowGridFieldsMap)
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	// assign modified queryclauses
	qc["WhereClause"] = srch
	qc["OrderClause"] = order
	qc["BID"] = strconv.FormatInt(d.BID, 10)
	qc["Start"] = d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL)
	qc["Stop"] = d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL)

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(RAFlowListQuery, qc)
	resp.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		rlib.Console("Error from rlib.GetQueryCount: %s\n", err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("g.Total = %d\n", resp.Total)

	//----------------------------------------------------------------------
	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	//----------------------------------------------------------------------
	limitAndOffsetClause := ` LIMIT {{.LimitClause}} OFFSET {{.OffsetClause}};`

	//----------------------------------------------------------------------
	// build query with limit and offset clause
	//----------------------------------------------------------------------
	RAFlowQueryWithLimit := RAFlowListQuery + limitAndOffsetClause

	//----------------------------------------------------------------------
	// Add limit and offset value
	//----------------------------------------------------------------------
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	//----------------------------------------------------------------------
	// get formatted query with substitution of select, where, order clause
	//----------------------------------------------------------------------
	qry := rlib.RenderSQLQuery(RAFlowQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	//----------------------------------------------------------------------
	// execute the query
	//----------------------------------------------------------------------
	var rows *sql.Rows
	rows, err = rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q GridRAFlowRecord
		q.Recid = i
		if err = rows.Scan(&q.BID, &q.RAID, &q.AgreementStart, &q.AgreementStop, &q.Payors, &q.FlowID, &q.UserRefNo, &q.FLAGS); err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		//----------------
		// Handle EDI...
		//----------------
		rlib.EDIHandleNDOutgoingDateRange(q.BID, &q.AgreementStart, &q.AgreementStop)

		resp.Records = append(resp.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}
	if err = rows.Err(); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	resp.Status = "success"
	SvcWriteResponse(d.BID, &resp, w)
}

// GetRAFlowRequest struct to get data from request
type GetRAFlowRequest struct {
	UserRefNo string
	RAID      int64
	Version   string // "raid" or "refno"
}

// GetRAFlow returns all existing Rental Agreements and all Flows
// are being edited for rental agreement.
func GetRAFlow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "GetRAFlow"
	var (
		flow rlib.Flow
		err  error
		req  GetRAFlowRequest
		tx   *sql.Tx
		ctx  context.Context
	)
	rlib.Console("Entered %s, \n", funcname)

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

		// COMMIT TRANSACTION
		if tx != nil {
			err = tx.Commit()
		}
	}()

	// ------- unmarshal the request data  ---------------
	if err := json.Unmarshal([]byte(d.data), &req); err != nil {
		return
	}

	// rlib.Console("req.Version = %s\n", req.Version)

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err = rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		return
	}

	// EditFlag should be set to true only when we're creating a Flow that
	// becomes a RefNo (an amended RentalAgreement)
	EditFlag := false // assume we're asking for the view version

	// BASED ON MODE DO OPERATION
	switch req.Version {
	case "raid":
		// GET RENTAL AGREEMENT
		var ra rlib.RentalAgreement
		ra, err = rlib.GetRentalAgreement(ctx, req.RAID)
		if err != nil {
			return
		}
		if ra.RAID == 0 {
			err = fmt.Errorf("Rental Agreement not found with given RAID: %d", req.RAID)
			return
		}

		// convert permanent ra to flow data and get it
		var raf rlib.RAFlowJSONData
		raf, err = rlib.ConvertRA2Flow(ctx, &ra, EditFlag)
		if err != nil {
			return
		}

		//-------------------------------------------------------------------------
		// Save the flow to the db
		//-------------------------------------------------------------------------
		var raflowJSONData []byte
		raflowJSONData, err = json.Marshal(&raf)
		if err != nil {
			return
		}

		//-------------------------------------------------------------------------
		// Fill out the datastructure and save it to the db as a flow...
		//-------------------------------------------------------------------------
		flow = rlib.Flow{
			BID:       ra.BID,
			FlowID:    0, // we're not creating any flow, just to see RA content
			UserRefNo: "",
			FlowType:  rlib.RAFlow,
			ID:        ra.RAID,
			Data:      raflowJSONData,
			CreateBy:  0,
			LastModBy: 0,
		}

		// -------------------
		// WRITE FLOW RESPONSE
		// -------------------
		SvcWriteFlowResponse(ctx, d.BID, flow, w)
		return

	case "refno":
		// CREATE ONE ONLY WHEN REF.NO IS BLANK
		// rlib.Console("req.UserRefNo = %s\n", req.UserRefNo)
		if req.UserRefNo == "" {

			// CHECK IF ANY FLOW EXIST WITH GIVEN RAID
			flow, err = rlib.GetFlowForRAID(ctx, "RA", req.RAID)
			if err != nil {
				return
			}
			if flow.FlowID > 0 {
				err = fmt.Errorf("flow already exists with given refno: %s", flow.UserRefNo)
				return
			}

			// rlib.Console("Generating new flow %s\n", req.UserRefNo)
			// IF NOT FOUND THEN TRY TO CREATE NEW ONE FROM RAID
			// GET RENTAL AGREEMENT
			var ra rlib.RentalAgreement
			ra, err = rlib.GetRentalAgreement(ctx, req.RAID)
			if err != nil {
				return
			}
			if ra.RAID == 0 {
				err = fmt.Errorf("rental Agreement not found with given RAID: %d", req.RAID)
				return
			}

			// GET THE NEW FLOW ID CREATED USING PERMANENT DATA
			var flowID int64
			EditFlag = true
			flowID, err = GetRA2FlowCore(ctx, &ra, d, EditFlag)
			if err != nil {
				return
			}

			// GET GENERATED FLOW USING NEW ID
			flow, err = rlib.GetFlow(ctx, flowID)
			if err != nil {
				return
			}

			// -------------------
			// WRITE FLOW RESPONSE
			// -------------------
			SvcWriteFlowResponse(ctx, d.BID, flow, w)
			return
		}

		// rlib.Console("Using existing flow\n")
		// GIVEN REF.NO SHOULD BE VALID
		if len(req.UserRefNo) != rlib.UserRefNoLength {
			err = fmt.Errorf("given refno is not valid: %s ", req.UserRefNo)
			return
		}

		// IF REF NO IS PROVIDED THEN TRY TO FIND
		// GET THE FLOW BY REFERENCE NO IF RAID == 0
		flow, err = rlib.GetFlowByUserRefNo(ctx, d.BID, req.UserRefNo)
		if err != nil {
			return
		}

		// IF FLOW FOUND WITH REF.NO THEN RETURN THE RESPONSE
		if flow.FlowID == 0 {
			err = fmt.Errorf("flow not found with given refno: %s ", req.UserRefNo)
			return
		}
		// -------------------
		// WRITE FLOW RESPONSE
		// -------------------
		SvcWriteFlowResponse(ctx, d.BID, flow, w)
		return

	default:
		err = fmt.Errorf("Invalid version: (%s) to get raflow for RAID: %d", req.Version, req.RAID)
		return
	}
}

// ValidateRAFlowAndAssignValidatedRAFlow validate raflow and assign it to the response
func ValidateRAFlowAndAssignValidatedRAFlow(ctx context.Context, raFlowData *rlib.RAFlowJSONData, flow rlib.Flow, raflowRespData *RAFlowResponse) {
	var (
		raFlowFieldsErrors    bizlogic.RAFlowFieldsErrors
		raFlowNonFieldsErrors bizlogic.RAFlowNonFieldsErrors
	)
	// init raFlowFieldsErrors
	initRAFlowFieldsErrors(&raFlowFieldsErrors)
	initRAFlowNonFieldsErrors(&raFlowNonFieldsErrors)
	bizlogic.ValidateRAFlowParts(ctx, &raFlowFieldsErrors, &raFlowNonFieldsErrors, raFlowData, flow.ID)
	totalFieldsError := raFlowFieldsErrors.Dates.Total + raFlowFieldsErrors.People.Total + raFlowFieldsErrors.Pets.Total + raFlowFieldsErrors.Vehicle.Total + raFlowFieldsErrors.Rentables.Total + raFlowFieldsErrors.ParentChild.Total + raFlowFieldsErrors.Tie.TiePeople.Total
	totalNonFieldsError := len(raFlowNonFieldsErrors.Dates) + len(raFlowNonFieldsErrors.People) + len(raFlowNonFieldsErrors.Pets) + len(raFlowNonFieldsErrors.Rentables) + len(raFlowNonFieldsErrors.Vehicle) + len(raFlowNonFieldsErrors.ParentChild) + len(raFlowNonFieldsErrors.Tie)
	raflowRespData.ValidationCheck.Errors = raFlowFieldsErrors
	raflowRespData.ValidationCheck.NonFieldsErrors = raFlowNonFieldsErrors
	raflowRespData.ValidationCheck.Total += totalFieldsError + totalNonFieldsError
}
