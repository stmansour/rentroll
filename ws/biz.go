package ws

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rcsv"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

//-------------------------------------------------------------------
//
//                        **** SEARCH ****
//
//-------------------------------------------------------------------

// Businesses is the structure describing a task list definition
type Businesses struct {
	Recid                 int64 `json:"recid"`
	BID                   int64
	BUD                   string            // reference to designation in Phonebook db
	Name                  string            //
	DefaultRentCycle      int64             // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	DefaultProrationCycle int64             // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	DefaultGSRPC          int64             // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	ClosePeriodTLID       int64             // Business used for closing a period
	ResDepARID            int64             // AR to use when creating a deposit for a (hotel) rentable
	ResForfeitARID        int64             // AR to use when creating a deposit for a (hotel) rentable
	ResRefundARID         int64             // AR to use when creating a deposit for a (hotel) rentable
	CPTLName              string            // Name of the TaskList
	FLAGS                 uint64            // the flags -- xlated to bools
	EDIenabled            bool              // true if EDI is enabled
	AllowBackdatedRA      bool              // true if backdating Rental Agreements is allowed
	Disabled              bool              // true if this business has been disabled
	LastModTime           rlib.JSONDateTime // when was this record last written
	LastModBy             int64             // employee UID (from phonebook) that modified it
	CreateTS              rlib.JSONDateTime // when was this record created
	CreateBy              int64             // employee UID (from phonebook) that created it
}

// SearchBizResponse holds the task list definition list
type SearchBizResponse struct {
	Status  string       `json:"status"`
	Total   int64        `json:"total"`
	Records []Businesses `json:"records"`
}

// which fields needs to be fetched for SQL query for assessment grid
var bizFieldsMap = map[string][]string{
	"BID":                   {"Business.BID"},
	"BUD":                   {"Business.BUD"},
	"Name":                  {"Business.Name"},
	"DefaultRentCycle":      {"Business.DefaultRentCycle"},
	"DefaultProrationCycle": {"Business.DefaultProrationCycle"},
	"DefaultGSRPC":          {"Business.DefaultGSRPC"},
	"ClosePeriodTLID":       {"Business.ClosePeriodTLID"},
	"EDIenabled":            {"Business.EDIenabled"},
	"LastModTime":           {"Business.LastModTime"},
	"LastModBy":             {"Business.LastModBy"},
	"CreateTS":              {"Business.CreateTS"},
	"CreateBy":              {"Business.CreateBy"},
}

// which fields needs to be fetched for SQL query for assessment grid
var bizQuerySelectFields = []string{
	"Business.BID",
	"Business.BUD",
	"Business.Name",
	"Business.DefaultRentCycle",
	"Business.DefaultProrationCycle",
	"Business.DefaultGSRPC",
	"Business.ClosePeriodTLID",
	"Business.FLAGS",
	"Business.LastModTime",
	"Business.LastModBy",
	"Business.CreateTS",
	"Business.CreateBy",
}

//-------------------------------------------------------------------
//
//                         **** SAVE ****
//
//-------------------------------------------------------------------

// SaveBusinessDef defines the fields supplied when Saving a Business
type SaveBusinessDef struct {
	Recid                 int64 `json:"recid"`
	BID                   int64
	BUD                   string            // reference to designation in Phonebook db
	Name                  string            //
	DefaultRentCycle      int64             // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	DefaultProrationCycle int64             // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	DefaultGSRPC          int64             // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	ClosePeriodTLID       int64             // Business used for closing a period
	ResDepARID            int64             // AR to use when creating a deposit for a (hotel) rentable
	ResForfeitARID        int64             // AR to use when creating a deposit for a (hotel) rentable
	ResRefundARID         int64             // AR to use when creating a deposit for a (hotel) rentable
	CPTLName              string            // Name of the TaskList
	FLAGS                 uint64            // the flags -- xlated to bools
	EDIenabled            bool              // true if EDI is enabled
	AllowBackdatedRA      bool              // true if backdating Rental Agreements is allowed
	Disabled              bool              // true if this business has been disabled
	LastModTime           rlib.JSONDateTime // when was this record last written
	LastModBy             int64             // employee UID (from phonebook) that modified it
	CreateTS              rlib.JSONDateTime // when was this record created
	CreateBy              int64             // employee UID (from phonebook) that created it
}

// SaveBusinessInput is the input data format for a Save command
type SaveBusinessInput struct {
	Recid    int64           `json:"recid"`
	Status   string          `json:"status"`
	FormName string          `json:"name"`
	Record   SaveBusinessDef `json:"record"`
}

//-------------------------------------------------------------------
//
//                         **** GET ****
//
//-------------------------------------------------------------------

// GetBizResponse is the response to a GetBusiness request
type GetBizResponse struct {
	Status string          `json:"status"`
	Record SaveBusinessDef `json:"record"`
}

// ############################################################################

// SvcHandlerBusiness handles requests to read/write/update or
// make-inactive a specific Business.  It routes the request to
// an appropriate handler
//
// The server command can be:
//      get     - read it
//      save    - Insert or Update
//      delete  - make it inactive
//-----------------------------------------------------------------------------
func SvcHandlerBusiness(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerBusiness"
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d\n", d.wsSearchReq.Cmd, d.BID)

	if d.BID < 0 || (d.BID == 0 && d.wsSearchReq.Cmd == "get") {
		SvcSearchHandlerBusinesses(w, r, d)
		return
	}

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.BID < 0 {
			err = fmt.Errorf("BID = %d. BusinessID is required but was not specified", d.BID)
			SvcErrorReturn(w, err, funcname)
			return
		}
		getBusiness(w, r, d)
	case "save":
		saveBusiness(w, r, d)
	case "delete":
		deleteBusiness(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// BusinessesRowScan scans a result from sql row and dump it in a
// Businesses struct
//
// RETURNS
//  Businesses
//-----------------------------------------------------------------------------
func BusinessesRowScan(rows *sql.Rows) (Businesses, error) {
	var q Businesses
	err := rows.Scan(
		&q.BID,
		&q.BUD,
		&q.Name,
		&q.DefaultRentCycle,
		&q.DefaultProrationCycle,
		&q.DefaultGSRPC,
		&q.ClosePeriodTLID,
		&q.FLAGS,
		&q.LastModTime,
		&q.LastModBy,
		&q.CreateTS,
		&q.CreateBy,
	)
	return q, err
}

// SvcSearchHandlerBusinesses generates a report of all Businesses defined
// business d.BID
// wsdoc {
//  @Title  Search Businesses
//	@URL /v1/business/:BUI
//  @Method  POST
//	@Synopsis Search Businesses
//  @Description  Search all Businesses and return those that match the Search Logic.
//	@Input wsSearchReq
//  @Response SearchBizResponse
// wsdoc }
//-----------------------------------------------------------------------------
func SvcSearchHandlerBusinesses(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcSearchHandlerBusinesses"
	var g SearchBizResponse
	var err error
	rlib.Console("Entered %s\n", funcname)
	rlib.BuildBusinessDesignationMap(r.Context()) // update BUDList cache

	whr := ""
	order := `Business.Name ASC` // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, bizFieldsMap)
	if len(whereClause) > 0 {
		whr += "WHERE " + whereClause
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	query := `
	SELECT {{.SelectClause}}
	FROM Business {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(bizQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	countQuery := rlib.RenderSQLQuery(query, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("g.Total = %d\n", g.Total)

	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`

	// build query with limit and offset clause
	// if query ends with ';' then remove it
	queryWithLimit := query + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(queryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		q, err := BusinessesRowScan(rows)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		q.Recid = i

		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}

	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// deleteBusiness makes the secified Business inactive
// wsdoc {
//  @Title  Delete Business
//	@URL /v1/business/:BUI/BID
//  @Method  POST
//	@Synopsis Make a Business inactive
//  @Desc  This service makes a Business inactive. We do not deliete
//  @Desc  Business
//	@Input DeletePmtForm
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func deleteBusiness(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteBusiness"
	var del DeletePmtForm

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	var biz rlib.Business
	err := rlib.GetBusiness(r.Context(), d.ID, &biz)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	biz.FLAGS |= 0x1 // bit 0 set means it is inactive
	err = rlib.UpdateBusiness(r.Context(), &biz)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// loadStrings does a csv import of the files containing the initial string
// lists for a new business
//
//-----------------------------------------------------------------------------
func loadStrings(ctx context.Context, a *rlib.Business) error {
	rlib.Console("Entered loadStrings\n")
	fname := "sample/strings.csv"
	t := rlib.LoadCSV(fname)
	rlib.Console("loaded csv file: %s, found %d strings\n", fname, len(t))
	for i := 0; i < len(t); i++ {
		if len(t[i][0]) == 0 {
			continue
		}
		if t[i][0][0] == '#' { // if it's a comment line, don't process it, just move on
			continue
		}
		//------------------------------------------------------------------
		// Column 1 is the BUD.  No matter what BUD is in the file, we must
		// set the BUD to the that of the supplied business...
		//------------------------------------------------------------------
		if i > 0 { // skip header row
			t[i][0] = a.Designation
		}
		_, err := rcsv.CreateStringList(ctx, t[i], i+1)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveBusiness returns the requested assessment
// wsdoc {
//  @Title  Save Business
//	@URL /v1/business/:BUI/BID
//  @Method  GET
//	@Synopsis Update the information on a Business with the supplied data
//  @Description This service updates Business :BID with the
//  @Description information supplied.
//	@Input SaveBusinessInput
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func saveBusiness(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveBusiness"
	var foo SaveBusinessInput
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	//---------------------------------------------------------------------
	// Create a Business struct based on the supplied info...
	//---------------------------------------------------------------------
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		rlib.Console("%s: err = %s\ndata = %s\n", funcname, err.Error(), d.data)
		return
	}
	var a rlib.Business
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	a.Name = foo.Record.Name
	a.BID = d.BID
	a.Designation = foo.Record.BUD
	a.FLAGS &= ^int64(7) // clear bits
	if foo.Record.EDIenabled {
		a.FLAGS |= (1 << 0)
	}
	if foo.Record.AllowBackdatedRA {
		a.FLAGS |= (1 << 1)
	}
	if foo.Record.Disabled {
		a.FLAGS |= (1 << 2)
	}

	//----------------------------------------------------------------
	// Not much business logic to check here.
	// 1. Ensure that there is a name.
	// 2. If it is an insert, make sure there's no duplicate name
	//----------------------------------------------------------------
	if len(a.Name) == 0 {
		e := fmt.Errorf("%s: Required field, Name, is blank", funcname)
		SvcErrorReturn(w, e, funcname)
		return
	}
	var adup rlib.Business
	adup, err = rlib.GetBusinessByDesignation(r.Context(), a.Name)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.Name == adup.Name && a.BID != adup.BID {
		e := fmt.Errorf("%s: A Business with the name %s already exists", funcname, a.Name)
		SvcErrorReturn(w, e, funcname)
		return
	}

	// rlib.Console("Save Business: a = %#v\n", a)

	//-------------------------------------------------------
	// Bizlogic checks done. Insert or update as needed...
	//-------------------------------------------------------

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	if a.BID == 0 {
		rlib.Console("Inserting new Business\n")
		_, err = rlib.InsertBusiness(ctx, &a) // This is a new record
		if err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}
		//--------------------------------------
		// ADD STRINGLISTS for the business
		//--------------------------------------
		rlib.Console("Inserting Business stringlists\n")
		if err = loadStrings(ctx, &a); err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}
		//--------------------------------------
		// ADD BUSINESS PROPERTIES
		//--------------------------------------
		var bp = rlib.BizProps{
			PetFees:     []string{},
			VehicleFees: []string{},
		}

		epoch := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
		bp.Epochs.Daily = epoch
		bp.Epochs.Weekly = epoch
		bp.Epochs.Monthly = epoch
		bp.Epochs.Quarterly = epoch
		bp.Epochs.Yearly = epoch
		var data []byte
		if data, err = json.Marshal(&bp); err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}
		var props = rlib.BusinessProperties{
			BID:   a.BID,
			Name:  "general",
			FLAGS: 0,
			Data:  data,
		}
		if _, err = rlib.InsertBusinessProperties(ctx, &props); err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}
	} else {
		// rlib.Console("Updating existing Business: %d\n", a.BID) // update existing record
		if err = rlib.UpdateBusiness(ctx, &a); err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}
		//-----------------------------------------
		// NOW UPDATE THE BUSINESS PROPERTIES...
		//-----------------------------------------
		var bp rlib.BusinessProperties
		if bp, err = rlib.GetBusinessPropertiesByName(ctx, "general", a.BID); err != nil {
		}
		var bizprops rlib.BizProps

		if err = json.Unmarshal(bp.Data, &bizprops); err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}

		bizprops.ResDepARID = foo.Record.ResDepARID
		bizprops.ResForfeitARID = foo.Record.ResForfeitARID
		bizprops.ResRefundARID = foo.Record.ResRefundARID
		//		rlib.Console("bizprops:  ResDepARID = %d, ResForfeitARID = %d, ResRefundARID = %d\n", bizprops.ResDepARID, bizprops.ResForfeitARID, bizprops.ResRefundARID)
		if err = rlib.UpdateBusinessProperties(ctx, &bp, &bizprops); err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.BID)
}

// GetBusiness returns the requested Business
// wsdoc {
//  @Title  Get Business
//	@URL /v1/business/:BUI/:BID
//  @Method  GET
//	@Synopsis Get information on a Business
//  @Description  Return all fields for assessment :BID
//	@Input WebGridSearchRequest
//  @Response GetBizResponse
// wsdoc }
//-----------------------------------------------------------------------------
func getBusiness(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getBusiness"
	var g GetBizResponse
	var a rlib.Business
	var err error

	rlib.Console("entered %s\n", funcname)
	err = rlib.GetBusiness(r.Context(), d.BID, &a)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	var bp rlib.BizProps
	if bp, err = rlib.GetDataFromBusinessPropertyName(r.Context(), "general", a.BID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	if a.BID > 0 {
		var gg SaveBusinessDef
		rlib.MigrateStructVals(&a, &gg)
		gg.ResDepARID = bp.ResDepARID
		gg.ResForfeitARID = bp.ResForfeitARID
		gg.ResRefundARID = bp.ResRefundARID

		// rlib.Console("\n\n\n***    a.ResDepARID = %d\n\n\n", gg.ResDepARID)
		gg.BUD = a.Designation
		gg.EDIenabled = a.FLAGS&(1<<0) != 0
		gg.AllowBackdatedRA = a.FLAGS&(1<<1) != 0
		gg.Disabled = a.FLAGS&(1<<2) != 0
		if a.ClosePeriodTLID > 0 {
			tl, err := rlib.GetTaskList(r.Context(), a.ClosePeriodTLID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			gg.CPTLName = tl.Name
		}
		g.Record = gg
	}

	// rlib.Console("Biz info for BID %d:  ClosePeriodTLID = %d\n", a.BID, a.ClosePeriodTLID)

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
