package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"sort"
	"strconv"
	"strings"
	"time"
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
	LastModTime    rlib.JSONDateTime
	LastModBy      int64
	RMRID          int64
	MarketRate     float64
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

// DeleteRentableTypeForm used to delete form
type DeleteRentableTypeForm struct {
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
	fmt.Println("Entered service handler for SvcRentableTypesList")

	var (
		g RentableTypesTDResponse
	)

	// get rentable types for a business
	m := rlib.GetBusinessRentableTypes(d.BID)

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
	SvcWriteResponse(&g, w)
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

	var (
		funcname = "SvcHandlerRentableType"
		err      error
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
				SvcGridErrorReturn(w, err, funcname)
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
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

// rtGridRowScan scans a result from sql row and dump it in a struct for rentableGrid
func rentableTypeGridRowScan(rows *sql.Rows, q RentableTypeGridRecord) (RentableTypeGridRecord, error) {
	err := rows.Scan(&q.RTID, &q.Style, &q.Name, &q.RentCycle, &q.Proration, &q.GSRPC, &q.ManageToBudget, &q.LastModTime, &q.LastModBy, &q.RMRID, &q.MarketRate)
	return q, err
}

var rtSearchFieldMap = selectQueryFieldMap{
	"RTID":           {"RentableTypes.RTID"},
	"Style":          {"RentableTypes.Style"},
	"Name":           {"RentableTypes.Name"},
	"RentCycle":      {"RentableTypes.RentCycle"},
	"Proration":      {"RentableTypes.Proration"},
	"GSRPC":          {"RentableTypes.GSRPC"},
	"ManageToBudget": {"RentableTypes.ManageToBudget"},
	"LastModTime":    {"RentableTypes.LastModTime"},
	"LastModBy":      {"RentableTypes.LastModBy"},
	"RMRID":          {"RentableMarketRate.RMRID"},
	"MarketRate":     {"RentableMarketRate.MarketRate"},
}

// which fields needs to be fetch to satisfy the struct
var rtSearchSelectQueryFields = selectQueryFields{
	"RentableTypes.RTID",
	"RentableTypes.Style",
	"RentableTypes.Name",
	"RentableTypes.RentCycle",
	"RentableTypes.Proration",
	"RentableTypes.GSRPC",
	"RentableTypes.ManageToBudget",
	"RentableTypes.LastModTime",
	"RentableTypes.LastModBy",
	"RentableMarketRate.RMRID",
	"RentableMarketRate.MarketRate",
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

	var (
		funcname    = "SvcSearchHandlerRentableTypes"
		g           RentableTypeSearchResponse
		err         error
		currentTime = time.Now().UTC()
		order       = `RentableTypes.RTID ASC` // default ORDER in sql result
		whr         = fmt.Sprintf("RentableTypes.BID=%d AND RentableMarketRate.DtStop>%q", d.BID, currentTime)
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
	SELECT
		{{.SelectClause}}
	FROM RentableTypes
	LEFT JOIN RentableMarketRate on RentableTypes.RTID=RentableMarketRate.RTID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := queryClauses{
		"SelectClause": strings.Join(rtSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := renderSQLQuery(rentableTypeSearchQuery, qc)
	g.Total, err = GetQueryCount(countQuery, qc)
	if err != nil {
		fmt.Printf("%s: Error from GetQueryCount: %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, err, funcname)
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
	qry := renderSQLQuery(rentableTypeQueryWithLimit, qc)
	fmt.Printf("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q RentableTypeGridRecord
		q.Recid = i
		q.BID = d.BID
		q.BUD = getBUDFromBIDList(q.BID)

		q, err = rentableTypeGridRowScan(rows, q)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
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
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
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

	var (
		funcname = "getRentableType"
		g        RentableTypeGetResponse
		whr      = fmt.Sprintf("RentableTypes.RTID=%d", d.ID)
	)

	fmt.Printf("entered %s\n", funcname)

	rentableTypeQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentableTypes
	LEFT JOIN RentableMarketRate on RentableTypes.RTID=RentableMarketRate.RTID
	WHERE {{.WhereClause}};`

	qc := queryClauses{
		"SelectClause": strings.Join(rtSearchSelectQueryFields, ","),
		"WhereClause":  whr,
	}

	qry := renderSQLQuery(rentableTypeQuery, qc)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var q RentableTypeGridRecord
		q.BID = d.BID
		q.BUD = getBUDFromBIDList(q.BID)

		q, err = rentableTypeGridRowScan(rows, q)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		q.Recid = q.RTID
		g.Record = q
	}
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	SvcWriteResponse(&g, w)
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
	var (
		funcname = "deleteRentableType"
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var del DeleteRentableTypeForm
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	// DeleteRentableType is still not implemented
	if err := rlib.DeleteRentableType(del.ID); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	// Remove rentableMarketRate related with it
	if err := rlib.DeleteRentableTypeRefWithRTID(del.ID); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(w)
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

	var (
		funcname    = "saveRentableType"
		foo         RentableTypeFormSave
		err         error
		currentTime = time.Now().UTC()
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	var a rlib.RentableType
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	var b rlib.RentableMarketRate
	rlib.MigrateStructVals(&foo.Record, &b) // the variables that don't need special handling

	fmt.Printf("RentableType Record: %#v\n", a)
	fmt.Printf("RentableMarketRate Record: %#v\n", b)

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, foo.Record.BUD)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	if a.RTID == 0 && d.ID == 0 {
		// This is a new AR
		fmt.Printf(">>>> NEW RentableType IS BEING ADDED\n")
		// override RTID if insertion done successful
		a.RTID, err = rlib.InsertRentableType(&a)
		if err != nil {
			e := fmt.Errorf("%s: unable to save RentableType record: %s", funcname, err.Error())
			SvcGridErrorReturn(w, e, funcname)
			return
		}
		b.RTID = a.RTID
		b.DtStart = currentTime
		b.DtStop = time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
		err = rlib.InsertRentableMarketRates(&b)
		if err != nil {
			e := fmt.Errorf("%s: unable to save RentableMarketRate record: %s", funcname, err.Error())
			SvcGridErrorReturn(w, e, funcname)
			return
		}
	} else {
		// update existing record
		fmt.Printf("Updating existing RentableType: %d\n", a.RTID)
		err = rlib.UpdateRentableType(&a)
		if err != nil {
			e := fmt.Errorf("%s: unable to update RentableType (RTID=%d\n: %s", funcname, a.RTID, err.Error())
			SvcGridErrorReturn(w, e, funcname)
			return
		}

		// first get RentableMarketRate instance from RMRID
		rmr, err := rlib.GetRentableMarketRateInstance(foo.Record.RMRID)
		if err != nil {
			e := fmt.Errorf("%s: unable to update RentableType (RTID=%d\n: %s", funcname, a.RTID, err.Error())
			SvcGridErrorReturn(w, e, funcname)
			return
		}

		// now check marketrate is changed or not
		if rmr.MarketRate != foo.Record.MarketRate { // if it is not same, then update rentableMarketRate with new record
			// mark old record's stop date as Today's date
			rmr.DtStop = currentTime
			err = rlib.UpdateRentableMarketRateInstance(&rmr)
			if err != nil {
				e := fmt.Errorf("%s: unable to update RentableType (RTID=%d\n: %s", funcname, a.RTID, err.Error())
				SvcGridErrorReturn(w, e, funcname)
				return
			}

			// insert new record
			newRMR := rmr
			newRMR.MarketRate = b.MarketRate
			newRMR.DtStart = currentTime
			newRMR.DtStop = time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
			err = rlib.InsertRentableMarketRates(&newRMR)
			if err != nil {
				e := fmt.Errorf("%s: unable to update RentableType (RTID=%d\n: %s", funcname, a.RTID, err.Error())
				SvcGridErrorReturn(w, e, funcname)
				return
			}
		}

	}

	SvcWriteSuccessResponse(w)
}
