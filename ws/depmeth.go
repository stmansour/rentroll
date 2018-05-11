package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// DepositMethodGrid contains the data from DepositMethod that is targeted to the UI Grid that displays
// a list of DepositMethod structs
type DepositMethodGrid struct {
	Recid       int64 `json:"recid"`
	DPMID       int64
	BID         int64
	BUD         rlib.XJSONBud
	Name        string
	LastModTime rlib.JSONDateTime
	LastModBy   int64
	CreateTS    rlib.JSONDateTime
	CreateBy    int64
}

// DepositMethodSearchResponse is a response string to the search request for DepositMethod records
type DepositMethodSearchResponse struct {
	Status  string              `json:"status"`
	Total   int64               `json:"total"`
	Records []DepositMethodGrid `json:"records"`
}

// DepositMethodSaveForm is a struct to handle direct inputs from the form
type DepositMethodSaveForm struct {
	Recid       int64 `json:"recid"`
	DPMID       int64
	BID         int64
	BUD         rlib.XJSONBud
	Name        string
	Description string
}

// SaveDepositMethodInput is the input data format for a Save command
type SaveDepositMethodInput struct {
	Recid    int64                 `json:"recid"`
	Status   string                `json:"status"`
	FormName string                `json:"name"`
	Record   DepositMethodSaveForm `json:"record"`
}

// DepositMethodGetResponse is the response to a GetDepositMethod request
type DepositMethodGetResponse struct {
	Status string            `json:"status"`
	Record DepositMethodGrid `json:"record"`
}

var depositMethodSearchFieldMap = rlib.SelectQueryFieldMap{
	"DPMID":       {"DepositMethod.DPMID"},
	"BID":         {"DepositMethod.BID"},
	"Method":      {"DepositMethod.Method"},
	"LastModTime": {"DepositMethod.LastModTime"},
	"LastModBy":   {"DepositMethod.LastModBy"},
	"CreateTS":    {"DepositMethod.CreateTS"},
	"CreateBy":    {"DepositMethod.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var depositMethodSearchSelectQueryFields = rlib.SelectQueryFields{
	"DepositMethod.DPMID",
	"DepositMethod.BID",
	"DepositMethod.Method",
	"DepositMethod.LastModTime",
	"DepositMethod.LastModBy",
	"DepositMethod.CreateTS",
	"DepositMethod.CreateBy",
}

// pmtRowScan scans a result from sql row and dump it in a DepositMethodGrid struct
func depositRowScan(rows *sql.Rows) (DepositMethodGrid, error) {
	var a DepositMethodGrid
	err := rows.Scan(&a.DPMID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	return a, err
}

// SvcHandlerDepositMethod formats a complete data record for an assessment for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the DPMID as follows:
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerDepositMethod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerDepositMethod"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  DPMID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerDepositMethods(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				err = fmt.Errorf("DepositMethodID is required but was not specified")
				SvcErrorReturn(w, err, funcname)
				return
			}
			getDepositMethod(w, r, d)
		}
	case "save":
		saveDepositMethod(w, r, d)
	case "delete":
		deleteDepositMethod(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// SvcSearchHandlerDepositMethods generates a report of all DepositMethods defined business d.BID
// wsdoc {
//  @Title  Search DepositMethod
//	@URL /v1/deposits/:BUI
//  @Method  POST
//	@Synopsis Search DepositMethods
//  @Descr  Search all DepositMethod and return those that match the Search Logic.
//  @Descr  The search criteria includes start and stop dates of interest.
//	@Input WebGridSearchRequest
//  @Response DepositMethodSearchResponse
// wsdoc }
func SvcSearchHandlerDepositMethods(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerDepositMethods"
	var (
		g     DepositMethodSearchResponse
		err   error
		order = "DPMID ASC" // default ORDER
		whr   = fmt.Sprintf("BID=%d", d.BID)
	)

	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	_, orderClause := GetSearchAndSortSQL(d, depositMethodSearchFieldMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	pmtQuery := `
	SELECT
		{{.SelectClause}}
	FROM DepositMethod
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(depositMethodSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(pmtQuery, qc)
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
	pmtQueryWithLimit := pmtQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(pmtQueryWithLimit, qc)
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

		q, err := depositRowScan(rows)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		q.Recid = i
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

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

// deleteDepositMethod deletes a deposit method from the database
// wsdoc {
//  @Title  Delete Deposit Method
//	@URL /v1/deposit/:BUI/DMID
//  @Method  POST
//	@Synopsis Delete a Deposit Method
//  @Desc  This service deletes a DepositMethod.
//	@Input DeleteDepMethForm
//  @Response SvcStatusResponse
// wsdoc }
func deleteDepositMethod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteDepository"
	var (
		del DeletePmtForm
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	err = rlib.DeleteDepositMethod(r.Context(), del.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// GetDepositMethod returns the requested assessment
// wsdoc {
//  @Title  Save DepositMethod
//	@URL /v1/deposit/:BUI/:DPMID
//  @Method  GET
//	@Synopsis Update the information on a DepositMethod with the supplied data
//  @Description  This service updates DepositMethod :DPMID with the information supplied. All fields must be supplied.
//	@Input SaveDepositMethodInput
//  @Response SvcStatusResponse
// wsdoc }
func saveDepositMethod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveDepositMethod"
	var (
		foo SaveDepositMethodInput
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	data := []byte(d.data)

	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	var a rlib.DepositMethod
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	var ok bool
	a.Method = foo.Record.Name
	a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, foo.Record.BUD)
		rlib.Ulog("%s", e.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	if len(a.Method) == 0 {
		e := fmt.Errorf("%s: Required field, Name, is blank", funcname)
		SvcErrorReturn(w, e, funcname)
		return
	}

	var adup rlib.DepositMethod
	adup, err = rlib.GetDepositMethodByName(r.Context(), a.BID, a.Method)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.Method == adup.Method && a.DPMID != adup.DPMID {
		e := fmt.Errorf("%s: A DepositMethod with the name %s already exists", funcname, a.Method)
		SvcErrorReturn(w, e, funcname)
		return
	}

	if a.DPMID == 0 && d.ID == 0 {
		// This is a new AR
		_, err = rlib.InsertDepositMethod(r.Context(), &a)
	} else {
		// update existing record
		fmt.Printf("Updating existing Payment Type: %d\n", a.DPMID)
		err = rlib.UpdateDepositMethod(r.Context(), &a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving Payment Type : %s", funcname, a.Method)
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// GetDepositMethod returns the requested DepositMethod
// wsdoc {
//  @Title  Get Payment Type
//	@URL /v1/deposit/:BUI/:DPMID
//  @Method  GET
//	@Synopsis Get information on a DepositMethod
//  @Description  Return all fields for assessment :DPMID
//	@Input WebGridSearchRequest
//  @Response DepositMethodGetResponse
// wsdoc }
func getDepositMethod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getDepositMethod"
	var (
		g   DepositMethodGetResponse
		a   rlib.DepositMethod
		err error
	)

	fmt.Printf("entered %s\n", funcname)
	a, err = rlib.GetDepositMethod(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.DPMID > 0 {
		var gg DepositMethodGrid
		rlib.MigrateStructVals(&a, &gg)
		gg.Name = a.Method
		gg.BUD = rlib.GetBUDFromBIDList(gg.BID)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
