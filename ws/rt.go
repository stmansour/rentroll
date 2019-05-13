package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"sort"
	"strconv"
	"strings"
)

// RentableTypeTD is struct to list down individual rentable type
type RentableTypeTD struct {
	RTID  int64  `json:"id"`
	FLAGS uint64 `json:"FLAGS"`
	Name  string `json:"text"`
}

// RentableTypesTDResponse - (TD: TypeDown) is the response to a GetRentable request
type RentableTypesTDResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []RentableTypeTD `json:"records"`
}

// RentableTypeGridRecord struct to show record in rentabletype grid
type RentableTypeGridRecord struct {
	Recid           int64 `json:"recid"`
	RTID            int64
	BID             int64
	BUD             rlib.XJSONBud
	Style           string
	Name            string
	RentCycle       int64
	Proration       int64
	GSRPC           int64
	ManageToBudget  bool
	ReserveAfter    bool
	flags           int64 // keep it in lowercase, don't send to client side
	IsActive        bool
	IsChildRentable bool
	ARID            int64
	LastModTime     rlib.JSONDateTime
	LastModBy       int64
	CreateTS        rlib.JSONDateTime
	CreateBy        int64
}

// RentableTypeSearchResponse is a response string to the search request for rentable types records
type RentableTypeSearchResponse struct {
	Status  string                   `json:"status"`
	Total   int64                    `json:"total"`
	Records []RentableTypeGridRecord `json:"records"`
}

// RentableTypeGetResponse is the response to a GetRentableType request
type RentableTypeGetResponse struct {
	Status string                 `json:"status"`
	Record RentableTypeGridRecord `json:"record"`
}

// RIDRequest has requested RID field
type RIDRequest struct {
	ID int64
}

// RentableTypeFormSave is the input data format for a Save command
type RentableTypeFormSave struct {
	Status   string                 `json:"status"`
	Recid    int64                  `json:"recid"`
	FormName string                 `json:"name"`
	Record   RentableTypeGridRecord `json:"record"`
}

// SvcRentableTypesTD generates a report of all Rentables defined business d.BID
// wsdoc {
//  @Title  Rentable Type List
//  @URL /v1/rtlist/:BUI
//  @Method  GET
//  @Synopsis Get Rentable Types
//  @Description Get all rentable types list for a requested business
//  @Input WebGridSearchRequest
//  @Response RentableTypesTypeDownResponse
// wsdoc }
func SvcRentableTypesTD(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRentableTypesTD"
	var (
		g RentableTypesTDResponse
	)
	rlib.Console("Entered in %s, service handler for SvcRentableTypesList\n", funcname)

	// get rentable types for a business
	m, err := rlib.GetBusinessRentableTypes(r.Context(), d.BID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
	}

	// sort keys
	var keys rlib.Int64Range
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	// append records in ascending order
	var rentableTypesList []RentableTypeTD
	for _, rtid := range keys {
		rentableTypesList = append(rentableTypesList, RentableTypeTD{RTID: m[rtid].RTID, FLAGS: m[rtid].FLAGS, Name: m[rtid].Name})
	}
	g.Records = rentableTypesList
	rlib.Console("GetBusinessRentableTypes returned %d records\n", len(g.Records))
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// SvcHandlerRentableType formats a complete data record for an assessment for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the RTID as follows:
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerRentableType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableType"
	var (
		err error
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  RTID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerRentableTypes(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				err = fmt.Errorf("ID for RentableType is required but was not specified")
				SvcErrorReturn(w, err, funcname)
				return
			}
			getRentableType(w, r, d)
		}
		break
	case "save":
		saveRentableType(w, r, d)
		break
	case "deactivate":
		deactivateRentableType(w, r, d)
		break
	case "reactivate":
		reactivateRentableType(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// rtGridRowScan scans a result from sql row and dump it in a struct for rentableGrid
func rentableTypeGridRowScan(rows *sql.Rows, q RentableTypeGridRecord) (RentableTypeGridRecord, error) {
	err := rows.Scan(&q.RTID, &q.Style, &q.Name, &q.RentCycle, &q.Proration, &q.GSRPC, &q.flags,
		&q.ARID, &q.LastModTime, &q.LastModBy, &q.CreateTS, &q.CreateBy)

	return q, err
}

var rtSearchFieldMap = rlib.SelectQueryFieldMap{
	"RTID":        {"RentableTypes.RTID"},
	"Style":       {"RentableTypes.Style"},
	"Name":        {"RentableTypes.Name"},
	"RentCycle":   {"RentableTypes.RentCycle"},
	"Proration":   {"RentableTypes.Proration"},
	"GSRPC":       {"RentableTypes.GSRPC"},
	"FLAGS":       {"RentableTypes.FLAGS"},
	"ARID":        {"RentableTypes.ARID"},
	"LastModTime": {"RentableTypes.LastModTime"},
	"LastModBy":   {"RentableTypes.LastModBy"},
	"CreateTS":    {"RentableTypes.CreateTS"},
	"CreateBy":    {"RentableTypes.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var rtSearchSelectQueryFields = rlib.SelectQueryFields{
	"RentableTypes.RTID",
	"RentableTypes.Style",
	"RentableTypes.Name",
	"RentableTypes.RentCycle",
	"RentableTypes.Proration",
	"RentableTypes.GSRPC",
	"RentableTypes.FLAGS",
	"RentableTypes.ARID",
	"RentableTypes.LastModTime",
	"RentableTypes.LastModBy",
	"RentableTypes.CreateTS",
	"RentableTypes.CreateBy",
}

// SvcSearchHandlerRentableTypes generates a report of all RentableTypes defined business d.BID
// wsdoc {
//  @Title  Search RentableType
//	@URL /v1/rt/:BUI
//  @Method  POST
//	@Synopsis Search RentableType
//  @Descr  Search all RentableTypes and return those that match the Search Logic.
//  @Descr  The search criteria includes start and stop dates of interest.
//	@Input WebGridSearchRequest
//  @Response RentableTypeSearchResponse
// wsdoc }
func SvcSearchHandlerRentableTypes(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerRentableTypes"
	var (
		g     RentableTypeSearchResponse
		err   error
		order = `RentableTypes.RTID ASC` // default ORDER in sql result
		whr   = fmt.Sprintf(`RentableTypes.BID=%d`, d.BID)
	)
	rlib.Console("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rtSearchFieldMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	rentableTypeSearchQuery := `
	SELECT {{.SelectClause}}
	FROM RentableTypes
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rtSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(rentableTypeSearchQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		rlib.Console("%s: Error from rlib.GetQueryCount: %s\n", funcname, err.Error())
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
	rentableTypeQueryWithLimit := rentableTypeSearchQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(rentableTypeQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		rlib.Console("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q RentableTypeGridRecord
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		q, err = rentableTypeGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		q.IsActive = q.flags&(1<<0) == 0        // Inactive when bit is set
		q.IsChildRentable = q.flags&(1<<1) != 0 // can be a child if set
		q.ManageToBudget = q.flags&(1<<2) != 0  // manage to budget when set
		q.ReserveAfter = q.flags&(1<<3) != 0    // reserve after when set

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
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}

// getRentableType returns the requested RentableType record
// wsdoc {
//  @Title  Get RentableType
//	@URL /v1/rt/:BUI/:RTID
//  @Method  GET
//	@Synopsis Get information on a RentableType
//  @Description  Return all fields for RentableType :RTID
//	@Input WebGridSearchRequest
//  @Response RentableTypeGetResponse
// wsdoc }
func getRentableType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getRentableType"
	var (
		g   RentableTypeGetResponse
		whr = fmt.Sprintf("RentableTypes.RTID=%d", d.ID)
	)

	rlib.Console("entered %s\n", funcname)

	rentableTypeQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentableTypes
	LEFT JOIN RentableMarketRate on RentableTypes.RTID=RentableMarketRate.RTID
	WHERE {{.WhereClause}};`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rtSearchSelectQueryFields, ","),
		"WhereClause":  whr,
	}

	qry := rlib.RenderSQLQuery(rentableTypeQuery, qc)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		rlib.Console("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var q RentableTypeGridRecord
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		q, err = rentableTypeGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		rlib.Console("\n\n******\nRTID = %d, FLAGS = %d\n******\n\n", q.RTID, q.flags)

		q.IsActive = q.flags&(1<<0) == 0        // Inactive when bit is set
		q.IsChildRentable = q.flags&(1<<1) != 0 // can be a child if set
		q.ManageToBudget = q.flags&(1<<2) != 0  // manage to budget when set
		q.ReserveAfter = q.flags&(1<<3) != 0    // reserve after when set

		q.Recid = q.RTID
		g.Record = q
	}
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}

// deactivateRentableType updates requested RentableType to inactive state
// wsdoc {
//  @Title  Delete RentableType
//	@URL /v1/rt/:BUI/:RTID
//  @Method  POST
//	@Synopsis Delete a RentableType
//  @Desc  This service inactivates a RentableType.
//	@Input RIDRequest
//  @Response SvcStatusResponse
// wsdoc }
func deactivateRentableType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deactivateRentableType"
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	var foo RIDRequest
	if err := json.Unmarshal([]byte(d.data), &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	rt := rlib.RentableType{RTID: foo.ID}
	if err := rlib.UpdateRentableTypeToInactive(r.Context(), &rt); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// reactivateRentableType re-activates a RentableType from the database
// wsdoc {
//  @Title  Reactivate RentableType
//	@URL /v1/rt/:BUI/:RTID
//  @Method  POST
//	@Synopsis Reactivate a RentableType (deleted previously)
//  @Desc  This service reactivates a RentableType.
//	@Input RIDRequest
//  @Response SvcStatusResponse
// wsdoc }
func reactivateRentableType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "reactivateRentableType"
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	var foo RIDRequest
	if err := json.Unmarshal([]byte(d.data), &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	rt := rlib.RentableType{RTID: foo.ID}
	if err := rlib.UpdateRentableTypeToActive(r.Context(), &rt); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// saveRentableType save the requested rentabletype with given data
// wsdoc {
//  @Title  Save RentableType
//	@URL /v1/rt/:BUI/:RTID
//  @Method  GET
//	@Synopsis Update the information on a RentableType with the supplied data
//  @Description  This service updates RentableType :RTID with the information supplied. All fields must be supplied.
//	@Input RentableTypeFormSave
//  @Response SvcStatusResponse
// wsdoc }
func saveRentableType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveRentableType"
	var (
		foo RentableTypeFormSave
		err error
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	var a rlib.RentableType
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	rlib.Console("RentableType Record: %#v\n", a)

	//
	// var ok bool
	// a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	// if !ok {
	// 	e := fmt.Errorf("%s: Could not map BID value: %s", funcname, foo.Record.BUD)
	// 	rlib.Ulog("%s", e.Error())
	// 	SvcErrorReturn(w, e, funcname)
	// 	return
	// }
	rlib.Console("foo.Record.IsChildRentable = %v\n", foo.Record.IsChildRentable)
	rlib.Console("foo.Record.ManageToBudget = %v\n", foo.Record.ManageToBudget)

	// NOTE: There is a separate API to deactive/reactive rentabletypes

	// this check is needed because to maintain FLAGS value for a rentabletype
	if !foo.Record.IsActive { // this would be hidden field on client side
		a.FLAGS |= (1 << 0) // 1 means inactive, if not set then only feed "1"
	}
	if foo.Record.IsChildRentable { // 1 << 1 -- first bit
		a.FLAGS |= (1 << 1)
	}
	if foo.Record.ManageToBudget { // 1<<
		a.FLAGS |= (1 << 2) // set it to 1
	}
	if foo.Record.ReserveAfter { // 1<<
		a.FLAGS |= (1 << 3) // set it to 1
	}

	errlist := bizlogic.ValidateRentableType(r.Context(), &a)
	if len(errlist) > 0 {
		SvcErrListReturn(w, errlist, funcname)
		return
	}

	if a.RTID == 0 && d.ID == 0 {
		// This is a new AR
		rlib.Console(">>>> NEW RentableType IS BEING ADDED\n")
		a.RTID, err = rlib.InsertRentableType(r.Context(), &a)
		if err != nil {
			e := fmt.Errorf("%s: unable to save RentableType record: %s", funcname, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	} else {
		// update existing record
		rlib.Console("Updating existing RentableType: %d, FLAGS = %x\n", a.RTID, a.FLAGS)
		err = rlib.UpdateRentableType(r.Context(), &a)
		if err != nil {
			e := fmt.Errorf("%s: unable to update RentableType (RTID=%d\n: %s", funcname, a.RTID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.RTID)
}
