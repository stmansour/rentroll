package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

//-------------------------------------------------------------------
//                        **** SEARCH ****
//-------------------------------------------------------------------

// RA2Flow is a structure specifically for the Web Services interface. It will be
// automatically populated from an rlib.RentalAgreement struct. Records of this type
// are returned by the search handler
type RA2Flow struct {
	Recid          int64           `json:"recid"` // this is to support the w2ui form
	RAID           int64           // internal unique id
	BID            int64           // Business (so that we can process by Business)
	AgreementStart rlib.JSONDate   // start date for rental agreement contract
	AgreementStop  rlib.JSONDate   // stop date for rental agreement contract
	Payors         rlib.NullString // payors list attached with this RA within same time
}

// RA2FlowSearchResponse is the response data for a Rental Agreement Search
type RA2FlowSearchResponse struct {
	Status  string    `json:"status"`
	Total   int64     `json:"total"`
	Records []RA2Flow `json:"records"`
}

//-------------------------------------------------------------------
//                         **** SAVE ****
//-------------------------------------------------------------------

// RA2FlowForm is used to save a Rental Agreement.  It holds those values
type RA2FlowForm struct {
	Recid                  int64 `json:"recid"` // this is to support the w2ui form
	BID                    int64
	BUD                    rlib.XJSONBud
	RAID                   int64             // internal unique id
	RATID                  int64             // reference to Occupancy Master Agreement
	NLID                   int64             // Note ID
	AgreementStart         rlib.JSONDate     // start date for rental agreement contract
	AgreementStop          rlib.JSONDate     // stop date for rental agreement contract
	PossessionStart        rlib.JSONDate     // start date for Occupancy
	PossessionStop         rlib.JSONDate     // stop date for Occupancy
	RentStart              rlib.JSONDate     // start date for Rent
	RentStop               rlib.JSONDate     // stop date for Rent
	RentCycleEpoch         rlib.JSONDate     // Date on which rent cycle recurs. Start date for the recurring rent assessment
	UnspecifiedAdults      int64             // adults who are not accounted for in RA2FlowPayor or RentableUser structs.  Used mostly by hotels
	UnspecifiedChildren    int64             // children who are not accounted for in RA2FlowPayor or RentableUser structs.  Used mostly by hotels.
	SpecialProvisions      string            // free-form text
	LeaseType              int64             // Full Service Gross, Gross, ModifiedGross, Tripple Net
	ExpenseAdjustmentType  int64             // Base Year, No Base Year, Pass Through
	ExpensesStop           float64           // cap on the amount of oexpenses that can be passed through to the tenant
	ExpenseStopCalculation string            // note on how to determine the expense stop
	BaseYearEnd            rlib.JSONDate     // last day of the base year
	ExpenseAdjustment      rlib.JSONDate     // the next date on which an expense adjustment is due
	EstimatedCharges       float64           // a periodic fee charged to the tenant to reimburse LL for anticipated expenses
	RateChange             float64           // predetermined amount of rent increase, expressed as a percentage
	NextRateChange         rlib.JSONDate     // he next date on which a RateChange will occur
	PermittedUses          string            // indicates primary use of the space, ex: doctor's office, or warehouse/distribution, etc.
	ExclusiveUses          string            // those uses to which the tenant has the exclusive rights within a complex, ex: Trader Joe's may have the exclusive right to sell groceries
	ExtensionOption        string            // the right to extend the term of lease by giving notice to LL, ex: 2 options to extend for 5 years each
	ExtensionOptionNotice  rlib.JSONDate     // the last date by which a Tenant can give notice of their intention to exercise the right to an extension option period
	ExpansionOption        string            // the right to expand to certanin spaces that are typically contiguous to their primary space
	ExpansionOptionNotice  rlib.JSONDate     // the last date by which a Tenant can give notice of their intention to exercise the right to an Expansion Option
	RightOfFirstRefusal    string            // Tenant may have the right to purchase their premises if LL chooses to sell
	Renewal                rlib.XJSONRenewal // month to month automatic extension, or lease extension
}

// GetRA2FlowResponse is the response data for GetRA2Flow
type GetRA2FlowResponse struct {
	Status string  `json:"status"`
	Record RA2Flow `json:"record"`
}

//-------------------------------------------------------------------
//                         **** DELETE ****
//-------------------------------------------------------------------

// DeleteRA2FlowForm used while deleteRA request
type DeleteRA2FlowForm struct {
	RAID int64
}

//* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// RA2FlowGridFieldsMap holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
var RA2FlowGridFieldsMap = map[string][]string{
	"RAID":           {"RentalAgreement.RAID"},
	"AgreementStart": {"RentalAgreement.AgreementStart"},
	"AgreementStop":  {"RentalAgreement.AgreementStop"},
	"Payors":         {"Transactant.FirstName", "Transactant.LastName", "Transactant.CompanyName"},
}

// RA2FlowQuerySelectFields defines which fields needs to be fetched for SQL query for rental agreements
var RA2FlowQuerySelectFields = []string{
	"RentalAgreement.RAID",
	"RentalAgreement.AgreementStart",
	"RentalAgreement.AgreementStop",
	"GROUP_CONCAT(DISTINCT CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END ORDER BY Transactant.TCID ASC SEPARATOR ', ') AS Payors",
}

// SvcHandlerRA2Flow handles requests for working with existing Rental
// Agreements
//
// The server command can be:
//      get     - read it
//      save    - Close the period (oldest unclosed period)
//      delete  - Reopen period
//-----------------------------------------------------------------------------
func SvcHandlerRA2Flow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRA2Flow"

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  d.ID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID < 0 {
			SvcSearchHandlerRA2Flow(w, r, d)
			return
		}
		getRA2Flow(w, r, d)
	case "save":
		saveRA2Flow(w, r, d)
	case "delete":
		deleteRA2Flow(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// SvcSearchHandlerRA2Flow generates a report of all RA2Flows defined business d.BID
// wsdoc {
//  @Title  Search Rental Agreements
//	@URL /v1/rentalagrs/:BUI
//  @Method  GET, POST
//	@Synopsis Return Rental Agreements that match the criteria provided.
//  @Description
//	@Input WebGridSearchRequest
//  @Response RA2FlowSearchResponse
// wsdoc }
func SvcSearchHandlerRA2Flow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerRA2Flow"
	const limitClause int = 100
	var err error
	var g RA2FlowSearchResponse

	rlib.Console("Entered %s\n", funcname)
	srch := fmt.Sprintf("RentalAgreement.BID=%d AND RentalAgreement.AgreementStop>%q AND RentalAgreement.AgreementStart<%q",
		d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL)) // default WHERE clause
	order := "RentalAgreement.RAID ASC" // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, RA2FlowGridFieldsMap)
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	// Rental Agreement Query Text Template
	RA2FlowQuery := `
	SELECT {{.SelectClause}}
	FROM RentalAgreement
	LEFT JOIN RentalAgreementPayors ON RentalAgreementPayors.RAID=RentalAgreement.RAID
	LEFT JOIN Transactant ON Transactant.TCID=RentalAgreementPayors.TCID
	WHERE {{.WhereClause}}
	GROUP BY RentalAgreement.RAID
	ORDER BY {{.OrderClause}}` // don't add ';', later some parts will be added in query

	//----------------------------------------------------------------------
	// Substitute query clauses
	//----------------------------------------------------------------------
	qc := rlib.QueryClause{
		"SelectClause": strings.Join(RA2FlowQuerySelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(RA2FlowQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		rlib.Console("Error from rlib.GetQueryCount: %s\n", err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("g.Total = %d\n", g.Total)

	//----------------------------------------------------------------------
	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	//----------------------------------------------------------------------
	limitAndOffsetClause := ` LIMIT {{.LimitClause}} OFFSET {{.OffsetClause}};`

	//----------------------------------------------------------------------
	// build query with limit and offset clause
	//----------------------------------------------------------------------
	RA2FlowQueryWithLimit := RA2FlowQuery + limitAndOffsetClause

	//----------------------------------------------------------------------
	// Add limit and offset value
	//----------------------------------------------------------------------
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	//----------------------------------------------------------------------
	// get formatted query with substitution of select, where, order clause
	//----------------------------------------------------------------------
	qry := rlib.RenderSQLQuery(RA2FlowQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	//----------------------------------------------------------------------
	// execute the query
	//----------------------------------------------------------------------
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q RA2Flow
		q.Recid = i
		q.BID = d.BID
		if err = rows.Scan(&q.RAID, &q.AgreementStart, &q.AgreementStop, &q.Payors); err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		g.Records = append(g.Records, q)
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

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// wsdoc {
//  @Title  Save Rental Agreement
//	@URL /v1/rentalagr/:BUI/:RAID
//  @Method  POST
//	@Synopsis Save (create or update) a Rental Agreement
//  @Description This service returns the single-valued attributes of a Rental Agreement.
//	@Input WebGridSearchRequest
//  @Response SvcStatusResponse
// wsdoc }
func saveRA2Flow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveRA2Flow"
	err := fmt.Errorf("unimplemented")
	SvcErrorReturn(w, err, funcname)
}

// wsdoc {
//  @Title  Get Rental Agreement
//	@URL /v1/rentalagr/:BUI/:RAID
//	@Method POST or GET
//	@Synopsis Get a Rental Agreement
//  @Description This service returns the single-valued attributes of a Rental Agreement.
//  @Input WebGridSearchRequest
//  @Response GetRA2FlowResponse
// wsdoc }
func getRA2Flow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getRA2Flow"
	var (
		flow           rlib.Flow
		g              FlowResponse
		raFlowResponse RAFlowResponse
		raFlowData     rlib.RAFlowJSONData
	)

	if d.ID < 1 {
		SvcErrorReturn(w, fmt.Errorf("invalid RAID: %d", d.ID), funcname)
		return
	}
	ra, err := rlib.GetRentalAgreement(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
	}

	//-------------------------------------------------------------------------
	//  Check to see if a flow already exists for this RAID. If so, just
	//  use it.
	//-------------------------------------------------------------------------
	ctx := r.Context()
	flow, err = rlib.GetFlowForRAID(ctx, "RA", ra.RAID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
	}

	if flow.ID == ra.RAID {
		// get unmarshalled raflow data into struct
		err = json.Unmarshal(flow.Data, &raFlowData)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		// Perform basic validation on flow data
		bizlogic.ValidateRAFlowBasic(r.Context(), &raFlowData, &raFlowResponse.BasicCheck)
		// Check DataFulfilled
		bizlogic.DataFulfilledRAFlow(r.Context(), &raFlowData, &raFlowResponse.DataFulfilled)
		raFlowResponse.Flow = flow
		// set the response
		g.Record = raFlowResponse
		g.Status = "success"
		SvcWriteResponse(d.BID, &g, w)
		return
	}

	flowID, err := GetRA2FlowCore(ctx, &ra, d.sess.UID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
	}

	flow, err = rlib.GetFlow(ctx, flowID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
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

	// Check DataFulfilled
	bizlogic.DataFulfilledRAFlow(r.Context(), &raFlowData, &raFlowResponse.DataFulfilled)

	raFlowResponse.Flow = flow
	// set the response
	g.Record = raFlowResponse
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// GetRA2FlowCore does all the heavy lifting to create a Flow from a
// RentalAgreement
//
// INPUTS:
//     ctx    database context for transactions
//     ra     the rental agreement to move into a flow
//     uid    uid of the person creating this flow.  Typically it
//            will be the uid in the session.
//
// RETURNS:
//     the new flowID
//     any error encountered
//-------------------------------------------------------------------------
func GetRA2FlowCore(ctx context.Context, ra *rlib.RentalAgreement, uid int64) (int64, error) {
	var flowID int64

	// convert permanent ra to flow data and get it
	raf, err := rlib.ConvertRA2Flow(ctx, ra)
	if err != nil {
		return flowID, err
	}

	//-------------------------------------------------------------------------
	// Save the flow to the db
	//-------------------------------------------------------------------------
	raflowJSONData, err := json.Marshal(&raf)
	if err != nil {
		return flowID, err
	}

	//-------------------------------------------------------------------------
	// Fill out the datastructure and save it to the db as a flow...
	//-------------------------------------------------------------------------
	a := rlib.Flow{
		BID:       ra.BID,
		FlowID:    0, // it's new flowID,
		UserRefNo: rlib.GenerateUserRefNo(),
		FlowType:  rlib.RAFlow,
		ID:        ra.RAID,
		Data:      raflowJSONData,
		CreateBy:  uid,
		LastModBy: uid,
	}

	// insert new flow
	flowID, err = rlib.InsertFlow(ctx, &a)
	if err != nil {
		return flowID, err
	}
	return flowID, nil
}

// ConvertRA2Flow does all the heavy lifting to convert existing
// rental agreement data to raflow data
//
// INPUTS:
//     ctx    database context for transactions
//     ra     the rental agreement to move into a flow
//
// RETURNS:
//     the rlib.RAFlowJSONData
//     any error encountered
//-------------------------------------------------------------------------
func ConvertRA2Flow(ctx context.Context, ra *rlib.RentalAgreement) (rlib.RAFlowJSONData, error) {
	const funcname = "ConvertRA2Flow"
	fmt.Printf("Entered in %s\n", funcname)

	//-------------------------------------------------------------
	// This is the datastructure we need to fill out and save...
	//-------------------------------------------------------------
	var raf = rlib.RAFlowJSONData{
		Dates: rlib.RADatesFlowData{
			BID:             ra.BID,
			RentStart:       rlib.JSONDate(ra.RentStart),
			RentStop:        rlib.JSONDate(ra.RentStop),
			AgreementStart:  rlib.JSONDate(ra.AgreementStart),
			AgreementStop:   rlib.JSONDate(ra.AgreementStop),
			PossessionStart: rlib.JSONDate(ra.PossessionStart),
			PossessionStop:  rlib.JSONDate(ra.PossessionStop),
		},
		People:      []rlib.RAPeopleFlowData{},
		Pets:        []rlib.RAPetsFlowData{},
		Vehicles:    []rlib.RAVehiclesFlowData{},
		Rentables:   []rlib.RARentablesFlowData{},
		ParentChild: []rlib.RAParentChildFlowData{},
		Tie: rlib.RATieFlowData{
			People: []rlib.RATiePeopleData{},
		},
		Meta: rlib.RAFlowMetaInfo{
			RAID:                   ra.RAID,
			RAFLAGS:                ra.FLAGS,
			Approver1:              ra.Approver1,
			DecisionDate1:          rlib.JSONDateTime(ra.DecisionDate1),
			DeclineReason1:         ra.DeclineReason1,
			Approver2:              ra.Approver2,
			DecisionDate2:          rlib.JSONDateTime(ra.DecisionDate2),
			DeclineReason2:         ra.DeclineReason2,
			TerminatorUID:          ra.TerminatorUID,
			TerminationDate:        rlib.JSONDateTime(ra.TerminationDate),
			LeaseTerminationReason: ra.LeaseTerminationReason,
			DocumentDate:           rlib.JSONDateTime(ra.DocumentDate),
			NoticeToMoveDate:       rlib.JSONDateTime(ra.NoticeToMoveDate),
			NoticeToMoveReported:   rlib.JSONDateTime(ra.NoticeToMoveReported),
		},
	}

	//-------------------------------------------------------------------------
	// Add Users...
	//
	// Note: we need to add users before payors in order to
	// ensure that all the pets and vehicles are loaded.  This is because
	// of two behaviors of the code.  First, pets and vehicles are loaded only
	// when the person is loaded with the User (occupant) flag set.  Second,
	// a person is not added twice, and if you load a Payor first -- the pets
	// and vehicles will NOT be loaded then when you call it the second time
	// to load the same transactant as a User the code will see that the
	// transactant has already been loaded and return without doing anything
	// other than setting the Payor (renter) flag.
	//-------------------------------------------------------------------------
	n, err := rlib.GetAllRentalAgreementRentables(ctx, ra.RAID)
	if err != nil {
		return raf, nil
	}
	for j := 0; j < len(n); j++ {
		rulist, err := rlib.GetRentableUsersInRange(ctx, n[j].RID, &ra.AgreementStart, &ra.AgreementStop)
		if err != nil {
			return raf, nil
		}
		for k := 0; k < len(rulist); k++ {
			addRAPtoFlow(ctx, rulist[k].TCID, n[j].RID, &raf, true, false, true)
		}
	}

	//-------------------------------------------------------------------------
	// Add Payors...
	//-------------------------------------------------------------------------
	m, err := rlib.GetRentalAgreementPayorsInRange(ctx, ra.RAID, &ra.AgreementStart, &ra.AgreementStop)
	if err != nil {
		return raf, nil
	}
	for i := 0; i < len(m); i++ {
		if err = addRAPtoFlow(ctx, m[i].TCID, 0 /*no RID here*/, &raf, true /*check dups*/, true /*renter*/, false); err != nil {
			return raf, nil
		}
	}

	//-------------------------------------------------------------------------
	// Add Rentables
	//-------------------------------------------------------------------------
	now := time.Now()
	o, err := rlib.GetRentalAgreementRentables(ctx, ra.RAID, &ra.AgreementStart, &ra.AgreementStop)
	if err != nil {
		return raf, nil
	}
	for i := 0; i < len(o); i++ {
		rnt, err := rlib.GetRentable(ctx, o[i].RID)
		if err != nil {
			return raf, nil
		}
		rtr, err := rlib.GetRentableTypeRefForDate(ctx, o[i].RID, &now)
		if err != nil {
			return raf, nil
		}
		var rt rlib.RentableType
		if err = rlib.GetRentableType(ctx, rtr.RTID, &rt); err != nil {
			return raf, nil
		}
		var rfd = rlib.RARentablesFlowData{
			BID:          o[i].BID,
			RID:          o[i].RID,
			RTID:         rtr.RTID,
			RTFLAGS:      rt.FLAGS,
			RentableName: rnt.RentableName,
			RentCycle:    rt.RentCycle,
		}

		//---------------------------------------------------------
		// Add the assessments associated with the Rentable...
		// For this we want to load all 1-time fees and all
		// recurring fees.
		//---------------------------------------------------------
		asms, err := rlib.GetAssessmentsByRAIDRID(ctx, rfd.BID, ra.RAID, rfd.RID)
		if err != nil {
			return raf, nil
		}
		for j := 0; j < len(asms); j++ {

			//----------------------------------------------------------
			// Get the account rule for this assessment...
			//----------------------------------------------------------
			ar, err := rlib.GetAR(ctx, asms[j].ARID)
			if err != nil {
				return raf, nil
			}

			//----------------------------------------------------------
			// Handle Rentable Fees that are NOT Pet or Vehicle related
			//----------------------------------------------------------
			if ar.FLAGS&(128|256) == 0 {
				raf.Meta.LastTMPASMID++

				var fee = rlib.RAFeesData{
					TMPASMID:       raf.Meta.LastTMPASMID,
					ASMID:          asms[j].ASMID,
					ARID:           asms[j].ARID,
					ARName:         ar.Name,
					ContractAmount: asms[j].Amount,
					RentCycle:      asms[j].RentCycle,
					Start:          rlib.JSONDate(asms[j].Start),
					Stop:           rlib.JSONDate(asms[j].Stop),
					Comment:        asms[j].Comment,
				}
				rfd.Fees = append(rfd.Fees, fee)
			}

			//----------------------------------------------------------
			// Handle PET Fees
			//----------------------------------------------------------
			// TMPASMID        int64 // unique ID to manage fees uniquely across all fees in raflow json data
			// ASMID           int64 // the permanent table assessment id if it is an existing RAID
			// ARID            int64
			// ARName          string
			// ContractAmount  float64
			// RentCycle       int64
			// Start           rlib.JSONDate
			// Stop            rlib.JSONDate
			// AtSigningPreTax float64
			// SalesTax        float64
			if ar.FLAGS&(128) != 0 { // Is it a pet fee?
				petid := asms[j].AssocElemID // find the pet...
				for k := 0; k < len(raf.Pets); k++ {
					if raf.Pets[k].PETID == petid {
						raf.Meta.LastTMPASMID++
						var pf = rlib.RAFeesData{
							TMPASMID:       raf.Meta.LastTMPASMID,
							ARID:           asms[j].ARID,
							ASMID:          asms[j].ASMID,
							ARName:         ar.Name,
							RentCycle:      asms[j].RentCycle,
							Start:          rlib.JSONDate(asms[j].Start),
							Stop:           rlib.JSONDate(asms[j].Stop),
							ContractAmount: asms[j].Amount,
							Comment:        asms[j].Comment,
						}
						raf.Pets[k].Fees = append(raf.Pets[k].Fees, pf)
						break
					}
				}
			}
			//----------------------------------------------------------
			// Handle VEHICLE Fees
			//----------------------------------------------------------
			if ar.FLAGS&(256) != 0 { // Is it a vehicle fee?
				vid := asms[j].AssocElemID // find the vehicle...
				for k := 0; k < len(raf.Vehicles); k++ {
					if raf.Vehicles[k].VID == vid {
						raf.Meta.LastTMPASMID++
						var pf = rlib.RAFeesData{
							TMPASMID:       raf.Meta.LastTMPASMID,
							ARID:           asms[j].ARID,
							ASMID:          asms[j].ASMID,
							ARName:         ar.Name,
							ContractAmount: asms[j].Amount,
							RentCycle:      asms[j].RentCycle,
							Start:          rlib.JSONDate(asms[j].Start),
							Stop:           rlib.JSONDate(asms[j].Stop),
							Comment:        asms[j].Comment,
						}
						raf.Vehicles[k].Fees = append(raf.Vehicles[k].Fees, pf)
						break
					}
				}
			}
		}

		raf.Rentables = append(raf.Rentables, rfd)
	}

	return raf, nil
}

// addRAPtoFlow adds a new person to raf.People.  The renter/occupant flags
// are only set if the corresponding input bool value is set.
//
// INPUTS
//     tcid  = the tcid of the transactant to load
//      rid  - the rentable that they are tied to
//      raf  - pointer to the flow struct to update
//      chk  - check to see if the tcid exists in raf.People before adding.
//             This is not always necessary, but only the caller knows.
// isRenter  - true if we need to set the RAPerson isRenter bool to true.
//             It should be true for Payors.
// isOccupant- true if we need to set the RAPerson isOccupant bool to true.
//             It should be true for Users.
//
// RETURNS
//     any error encountered
//     raf is updated
//-----------------------------------------------------------------------------
func addRAPtoFlow(ctx context.Context, tcid, rid int64, raf *rlib.RAFlowJSONData, chk, isRenter, isOccupant bool) error {
	// Is this user already present?
	if chk {
		for l := 0; l < len(raf.People); l++ {
			if raf.People[l].TCID == tcid {
				if isRenter {
					raf.People[l].IsRenter = true
				}
				if isOccupant {
					raf.People[l].IsOccupant = true
				}
				return nil
			}
		}
	}

	rap, err := createRAFlowPerson(ctx, tcid, raf, isOccupant) // adds person AND associated pets and vehicles
	if err != nil {
		return err
	}

	if isRenter {
		rap.IsRenter = true
	}

	if isOccupant {
		rap.IsOccupant = true

		// only tie occupants to rentable
		var t rlib.RATiePeopleData
		t.BID = rap.BID
		t.TMPTCID = rap.TMPTCID
		if rid > 0 {
			t.PRID = rid
		}
		raf.Tie.People = append(raf.Tie.People, t)
	}

	// finally append in people list
	raf.People = append(raf.People, rap)
	return nil
}

// createRAFlowPerson returns a new rlib.RAPeopleFlowData based on the supplied
// tcid. It does not set the Renter or Occupant flags
//
// INPUTS
//          ctx  = db transaction context
//         tcid  = the tcid of the transactant to load
//          raf  = pointer to rlib.RAFlowJSONData
// addDependents = adds dependents (currently pets and vehicles) to the flow
//                 data in addition to the transactant data. The recommended
//                 usage of this flag is to set it to true when the person
//                 being added is a user.
//
// RETURNS
//     rlib.RAPeopleFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func createRAFlowPerson(ctx context.Context, tcid int64, raf *rlib.RAFlowJSONData, addDependents bool) (rlib.RAPeopleFlowData, error) {
	var p rlib.Transactant
	var pu rlib.User
	var pp rlib.Payor
	var pr rlib.Prospect
	var rap rlib.RAPeopleFlowData
	var err error

	raf.Meta.LastTMPTCID++
	rap.TMPTCID = raf.Meta.LastTMPTCID // set this now so it is available when creating pets and vehicles
	if err = rlib.GetTransactant(ctx, tcid, &p); err != nil {
		return rap, err
	}
	if err = rlib.GetUser(ctx, tcid, &pu); err != nil {
		return rap, err
	}
	if err = rlib.GetPayor(ctx, tcid, &pp); err != nil {
		return rap, err
	}
	if err = rlib.GetProspect(ctx, tcid, &pr); err != nil {
		return rap, err
	}
	rlib.MigrateStructVals(&p, &rap)
	rlib.MigrateStructVals(&pp, &rap)
	rlib.MigrateStructVals(&pu, &rap)
	rlib.MigrateStructVals(&pr, &rap)

	if addDependents {
		if err = addFlowPersonVehicles(ctx, tcid, rap.TMPTCID, raf); err != nil {
			return rap, err
		}
		if err = addFlowPersonPets(ctx, tcid, rap.TMPTCID, raf); err != nil {
			return rap, err
		}
	}
	return rap, nil
}

// addFlowPersonPets adds pets belonging to tcid to the supplied
// rlib.RAFlowJSONData struct
//
// INPUTS
//      ctx  = db transaction context
//     tcid  = the tcid of the transactant to load
//
// RETURNS
//     rlib.RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func addFlowPersonPets(ctx context.Context, tcid, tmptcid int64, raf *rlib.RAFlowJSONData) error {
	petList, err := rlib.GetPetsByTransactant(ctx, tcid)
	if err != nil {
		return err
	}
	for i := 0; i < len(petList); i++ {
		raf.Meta.LastTMPPETID++
		var p = rlib.RAPetsFlowData{
			TMPTCID:  tmptcid,
			TMPPETID: raf.Meta.LastTMPPETID,
		}
		rlib.MigrateStructVals(&petList[i], &p)
		raf.Pets = append(raf.Pets, p)
	}
	return nil
}

// addFlowPersonVehicles adds vehicles belonging to tcid to the supplied
// rlib.RAFlowJSONData struct
//
// INPUTS
//      ctx  = db transaction context
//     tcid  = the tcid of the transactant to load
//
// RETURNS
//     rlib.RAPetsFlowData structure
//     any error encountered
//-----------------------------------------------------------------------------
func addFlowPersonVehicles(ctx context.Context, tcid, tmptcid int64, raf *rlib.RAFlowJSONData) error {
	vehicleList, err := rlib.GetVehiclesByTransactant(ctx, tcid)
	if err != nil {
		return err
	}
	for i := 0; i < len(vehicleList); i++ {
		raf.Meta.LastTMPVID++
		var v = rlib.RAVehiclesFlowData{
			TMPTCID: tmptcid,
			TMPVID:  raf.Meta.LastTMPVID,
		}
		rlib.MigrateStructVals(&vehicleList[i], &v)
		raf.Vehicles = append(raf.Vehicles, v)
	}
	return nil
}

// wsdoc {
//  @Title  Delete Rental Agreement
//	@URL /v1/rentalagr/:BUI/:RAID
//	@Method POST
//	@Synopsis Delete a Rental Agreement
//  @Description This service delete the requested Rental Agreement with RAID and deletes associated pets, users, payors, references to rentables.
//  @Input DeleteRA2FlowForm
//  @Response SvcStatusResponse
// wsdoc }
func deleteRA2Flow(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteRA2Flow"
	err := fmt.Errorf("unimplemented")
	SvcErrorReturn(w, err, funcname)
}
