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
	RTID int64  `json:"id"`
	Name string `json:"text"`
}

// RentableTypesTDResponse - (TD: TypeDown) is the response to a GetRentable request
type RentableTypesTDResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []RentableTypeTD `json:"records"`
}

// RentableTypeGridRecord struct to show record in rentabletype grid
type RentableTypeGridRecord struct {
	Recid          int64 `json:"recid"`
	RTID           int64
	BID            int64
	BUD            rlib.XJSONBud
	Style          string
	Name           string
	RentCycle      int64
	Proration      int64
	GSRPC          int64
	ManageToBudget int64
	FLAGS          int64
	ARID           int64
	LastModTime    rlib.JSONDateTime
	LastModBy      int64
	CreateTS       rlib.JSONDateTime
	CreateBy       int64
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

// DeleteRentableTypeForm used to inactive Rentable Type
type DeleteRentableTypeForm struct {
	ID int64
}

// ReactivateRentableTypeForm used to reactivate Rentable Type
type ReactivateRentableTypeForm struct {
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
	fmt.Printf("Entered in %s, service handler for SvcRentableTypesList\n", funcname)

	// get rentable types for a business
	m, _ := rlib.GetBusinessRentableTypes(r.Context(), d.BID)

	// sort keys
	var keys rlib.Int64Range
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	// append records in ascending order
	var rentableTypesList []RentableTypeTD
	for _, rtid := range keys {
		rentableTypesList = append(rentableTypesList, RentableTypeTD{RTID: m[rtid].RTID, Name: m[rtid].Name})
	}
	g.Records = rentableTypesList
	fmt.Printf("GetBusinessRentableTypes returned %d records\n", len(g.Records))
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

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  RTID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

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
	case "delete":
		deleteRentableType(w, r, d)
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
	err := rows.Scan(&q.RTID, &q.Style, &q.Name, &q.RentCycle, &q.Proration, &q.GSRPC, &q.ManageToBudget, &q.FLAGS,
		&q.ARID, &q.LastModTime, &q.LastModBy, &q.CreateTS, &q.CreateBy)
	return q, err
}

var rtSearchFieldMap = rlib.SelectQueryFieldMap{
	"RTID":           {"RentableTypes.RTID"},
	"Style":          {"RentableTypes.Style"},
	"Name":           {"RentableTypes.Name"},
	"RentCycle":      {"RentableTypes.RentCycle"},
	"Proration":      {"RentableTypes.Proration"},
	"GSRPC":          {"RentableTypes.GSRPC"},
	"ManageToBudget": {"RentableTypes.ManageToBudget"},
	"FLAGS":          {"RentableTypes.FLAGS"},
	"ARID":           {"RentableTypes.ARID"},
	"LastModTime":    {"RentableTypes.LastModTime"},
	"LastModBy":      {"RentableTypes.LastModBy"},
	"CreateTS":       {"RentableTypes.CreateTS"},
	"CreateBy":       {"RentableTypes.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var rtSearchSelectQueryFields = rlib.SelectQueryFields{
	"RentableTypes.RTID",
	"RentableTypes.Style",
	"RentableTypes.Name",
	"RentableTypes.RentCycle",
	"RentableTypes.Proration",
	"RentableTypes.GSRPC",
	"RentableTypes.ManageToBudget",
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
		whr   = fmt.Sprintf(`RentableTypes.BID=%d
				AND (RentableMarketRate.DtStart <= %q OR RentableMarketRate.DtStart IS NULL)
				AND (RentableMarketRate.DtStop >%q OR RentableMarketRate.DtStop IS NULL)`,
			d.BID, d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL))
	)
	fmt.Printf("Entered %s\n", funcname)

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
	LEFT JOIN RentableMarketRate on RentableTypes.RTID=RentableMarketRate.RTID
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
		fmt.Printf("%s: Error from rlib.GetQueryCount: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("g.Total = %d\n", g.Total)

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
	fmt.Printf("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
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

	fmt.Printf("entered %s\n", funcname)

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
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
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

// deleteRentableType deletes a RentableType from the database
// wsdoc {
//  @Title  Delete RentableType
//	@URL /v1/rt/:BUI/:RTID
//  @Method  POST
//	@Synopsis Delete a RentableType
//  @Desc  This service deletes a RentableType.
//	@Input WebGridDelete
//  @Response SvcStatusResponse
// wsdoc }
func deleteRentableType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteRentableType"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var del DeleteRentableTypeForm
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	if err := rlib.DeleteRentableType(r.Context(), del.ID); err != nil {
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
//	@Input ReactivateRentableTypeForm
//  @Response SvcStatusResponse
// wsdoc }
func reactivateRentableType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "reactivateRentableType"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var reActF ReactivateRentableTypeForm
	if err := json.Unmarshal([]byte(d.data), &reActF); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	rt := rlib.RentableType{RTID: reActF.ID}
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

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	var a rlib.RentableType
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	fmt.Printf("RentableType Record: %#v\n", a)

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, foo.Record.BUD)
		rlib.Ulog("%s", e.Error())
		SvcErrorReturn(w, e, funcname)
		return
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
		fmt.Printf("Updating existing RentableType: %d\n", a.RTID)
		err = rlib.UpdateRentableType(r.Context(), &a)
		if err != nil {
			e := fmt.Errorf("%s: unable to update RentableType (RTID=%d\n: %s", funcname, a.RTID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.RTID)
}

// RentableMarketRateGridResponse holds the struct for grids response
type RentableMarketRateGridResponse struct {
	Status  string                      `json:"status"`
	Total   int64                       `json:"total"`
	Records []RentableMarketRateGridRec `json:"records"`
}

// RentableMarketRateGridRec holds a struct for single record of grid
type RentableMarketRateGridRec struct {
	Recid      int64 `json:"recid"`
	BID        int64
	BUD        rlib.XJSONBud
	RMRID      int64
	RTID       int64
	MarketRate float64
	DtStart    rlib.JSONDate
	DtStop     rlib.JSONDate
}

// rmrGridRowScan scans a result from sql row and dump it in a struct for rentableGrid
func rmrGridRowScan(rows *sql.Rows, q RentableMarketRateGridRec) (RentableMarketRateGridRec, error) {
	err := rows.Scan(&q.RTID, &q.RMRID, &q.MarketRate, &q.DtStart, &q.DtStop)
	return q, err
}

var rmrSearchFieldMap = rlib.SelectQueryFieldMap{
	"RTID":       {"RentableMarketRate.RTID"},
	"RMRID":      {"RentableMarketRate.RMRID"},
	"MarketRate": {"RentableMarketRate.MarketRate"},
	"DtStart":    {"RentableMarketRate.DtStart"},
	"DtStop":     {"RentableMarketRate.DtStop"},
}

// which fields needs to be fetch to satisfy the struct
var rmrSearchSelectQueryFields = rlib.SelectQueryFields{
	"RentableMarketRate.RTID",
	"RentableMarketRate.RMRID",
	"RentableMarketRate.MarketRate",
	"RentableMarketRate.DtStart",
	"RentableMarketRate.DtStop",
}

// SvcHandlerRentableMarketRates returns the list of market rates for given rentable type
func SvcHandlerRentableMarketRates(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableMarketRates"
	var (
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  RTID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	// This operation requires RentableType ID
	if d.ID < 0 {
		err = fmt.Errorf("ID for RentableType is required but was not specified")
		SvcErrorReturn(w, err, funcname)
		return
	}

	switch d.wsSearchReq.Cmd {
	case "get":
		svcSearchHandlerRentableMarketRates(w, r, d) // it is a query for the grid.
		break
	case "save":
		saveRentableTypeMarketRates(w, r, d)
		break
	case "delete":
		deleteRentableTypeMarketRates(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// svcSearchHandlerRentableMarketRates handles market rate grid request/response
func svcSearchHandlerRentableMarketRates(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "svcSearchHandlerRentableMarketRates"
	var (
		g     RentableMarketRateGridResponse
		err   error
		order = `RentableMarketRate.RMRID ASC`
		whr   = fmt.Sprintf("RentableMarketRate.RTID=%d", d.ID)
	)
	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rmrSearchFieldMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	mrQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentableMarketRate
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rmrSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(mrQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		fmt.Printf("%s: Error from rlib.GetQueryCount: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("g.Total = %d\n", g.Total)

	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`

	// build query with limit and offset clause
	// if query ends with ';' then remove it
	queryWithLimit := mrQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(queryWithLimit, qc)
	fmt.Printf("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q RentableMarketRateGridRec
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		q, err = rmrGridRowScan(rows, q)
		if err != nil {
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

	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}

// MarketRateGridSave is the input data format for a Save command
type MarketRateGridSave struct {
	Cmd      string                      `json:"cmd"`
	Selected []int64                     `json:"selected"`
	Limit    int64                       `json:"limit"`
	Offset   int64                       `json:"offset"`
	Changes  []RentableMarketRateGridRec `json:"changes"`
}

// saveRentableTypeMarketRates save/update market rates associated with RentableType
func saveRentableTypeMarketRates(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveRentableTypeMarketRates"
	var (
		err error
		foo MarketRateGridSave
	)
	fmt.Printf("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	fmt.Printf("foo Changes: %v", foo.Changes)

	if len(foo.Changes) == 0 {
		e := fmt.Errorf("No MarketRate(s) provided for RentableType")
		SvcErrorReturn(w, e, funcname)
		return
	}

	// first check that, associated RentableType has allowed ManageToBudget field
	// if not then return with error
	var rt rlib.RentableType
	rtid := foo.Changes[0].RTID
	if err = rlib.GetRentableType(r.Context(), rtid, &rt); err != nil {
		e := fmt.Errorf("Error while getting RentableType: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	if rt.ManageToBudget == 0 {
		e := fmt.Errorf("ManageToBudget is not enabled at this moment")
		SvcErrorReturn(w, e, funcname)
		return
	}

	var bizErrs []bizlogic.BizError
	for _, mr := range foo.Changes {
		var a rlib.RentableMarketRate
		rlib.MigrateStructVals(&mr, &a) // the variables that don't need special handling

		errs := bizlogic.ValidateRentableMarketRate(r.Context(), &a)
		if len(errs) > 0 {
			bizErrs = append(bizErrs, errs...)
			continue
		}

		// insert new marketRate
		if a.RMRID == 0 {
			_, err = rlib.InsertRentableMarketRates(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error while inserting market rate:  %s", err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		} else { // else update existing one
			err = rlib.UpdateRentableMarketRateInstance(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error with updating market rate instance (%d), RTID=%d : %s", a.RMRID, a.RTID, err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		}
	}

	// if any marketRate has problem in bizlogic then return list
	if len(bizErrs) > 0 {
		SvcErrListReturn(w, bizErrs, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// MarketRateGridDelete is a struct used in delete request for market rates
type MarketRateGridDelete struct {
	Cmd       string  `json:"cmd"`
	RMRIDList []int64 `json:"RMRIDList"`
}

// deleteRentableTypeMarketRates used to delete multiple market rate records associated with rentable type
func deleteRentableTypeMarketRates(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteRentableTypeMarketRates"
	var (
		err error
		foo MarketRateGridDelete
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	for _, rmrid := range foo.RMRIDList {
		err = rlib.DeleteRentableMarketRateInstance(r.Context(), rmrid)
		if err != nil {
			e := fmt.Errorf("Error with deleting MarketRate with ID %d:  %s", rmrid, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
}
