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

// DepositoryGrid contains the data from Depository that is targeted to the UI Grid that displays
// a list of Depository structs
type DepositoryGrid struct {
	Recid       int64 `json:"recid"`
	DEPID       int64
	BID         int64
	BUD         rlib.XJSONBud
	LID         int64
	Name        string
	AccountNo   string
	LdgrName    string
	GLNumber    string
	LastModTime rlib.JSONDateTime
	LastModBy   int64
	CreateTS    rlib.JSONDateTime
	CreateBy    int64
}

// DepositorySearchResponse is a response string to the search request for Depository records
type DepositorySearchResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []DepositoryGrid `json:"records"`
}

// DepositorySaveForm contains the data from Depository FORM that is targeted to the UI Form that displays
// a list of Depository structs
type DepositorySaveForm struct {
	Recid     int64 `json:"recid"`
	LID       int64
	DEPID     int64
	BID       int64
	BUD       rlib.XJSONBud
	Name      string
	AccountNo string
	LdgrName  string
	GLNumber  string
}

// DepositoryGridSave is the input data format for a Save command
type DepositoryGridSave struct {
	Status   string             `json:"status"`
	Recid    int64              `json:"recid"`
	FormName string             `json:"name"`
	Record   DepositorySaveForm `json:"record"`
}

// DepositoryGetResponse is the response to a GetDepository request
type DepositoryGetResponse struct {
	Status string         `json:"status"`
	Record DepositoryGrid `json:"record"`
}

// DeleteDepForm used to delete form
type DeleteDepForm struct {
	ID int64
}

// SvcHandlerDepository formats a complete data record for an assessment for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the DEPID as follows:
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerDepository(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerDepository"
	var (
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  DEPID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerDepositories(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				err = fmt.Errorf("DepositoryID is required but was not specified")
				SvcErrorReturn(w, err, funcname)
				return
			}
			getDepository(w, r, d)
		}
		break
	case "save":
		saveDepository(w, r, d)
		break
	case "delete":
		deleteDepository(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// depGridRowScan scans a result from sql row and dump it in a PrARGrid struct
func depGridRowScan(rows *sql.Rows, q DepositoryGrid) (DepositoryGrid, error) {
	err := rows.Scan(&q.DEPID, &q.LID, &q.Name, &q.AccountNo, &q.LdgrName, &q.GLNumber, &q.LastModTime, &q.LastModBy, &q.CreateTS, &q.CreateBy)
	return q, err
}

var depSearchFieldMap = rlib.SelectQueryFieldMap{
	"DEPID":       {"Depository.DEPID"},
	"LID":         {"Depository.LID"},
	"Name":        {"Depository.Name"},
	"AccountNo":   {"Depository.AccountNo"},
	"LdgrName":    {"GLAccount.Name"},
	"GLNumber":    {"GLAccount.GLNumber"},
	"LastModTime": {"Depository.LastModTime"},
	"LastModBy":   {"Depository.LastModBy"},
	"CreateTS":    {"Depository.CreateTS"},
	"CreateBy":    {"Depository.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var depSearchSelectQueryFields = rlib.SelectQueryFields{
	"Depository.DEPID",
	"Depository.LID",
	"Depository.Name",
	"Depository.AccountNo",
	"GLAccount.Name as LdgrName",
	"GLAccount.GLNumber",
	"Depository.LastModTime",
	"Depository.LastModBy",
	"Depository.CreateTS",
	"Depository.CreateBy",
}

// SvcSearchHandlerDepositories generates a report of all Depositories defined business d.BID
// wsdoc {
//  @Title  Search Depositories
//	@URL /v1/dep/:BUI
//  @Method  POST
//	@Synopsis Search Depositories
//  @Descr  Search all Depository and return those that match the Search Logic.
//  @Descr  The search criteria includes start and stop dates of interest.
//	@Input WebGridSearchRequest
//  @Response DepositorySearchResponse
// wsdoc }
func SvcSearchHandlerDepositories(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerDepositories"
	var (
		g     DepositorySearchResponse
		err   error
		order = `Depository.DEPID ASC` // default ORDER in sql result
		whr   = fmt.Sprintf("Depository.BID=%d", d.BID)
	)
	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	_, orderClause := GetSearchAndSortSQL(d, depSearchFieldMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	depSearchQuery := `
	SELECT
		{{.SelectClause}}
	FROM Depository
	LEFT JOIN GLAccount on GLAccount.LID=Depository.LID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(depSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(depSearchQuery, qc)
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
	depQueryWithLimit := depSearchQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(depQueryWithLimit, qc)
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
		var q DepositoryGrid
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		q, err = depGridRowScan(rows, q)
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

// deleteDepository deletes a payment type from the database
// wsdoc {
//  @Title  Delete Depository
//	@URL /v1/dep/:BUI/:RAID
//  @Method  POST
//	@Synopsis Delete a Payment Type
//  @Desc  This service deletes a Depository.
//	@Input WebGridDelete
//  @Response SvcStatusResponse
// wsdoc }
func deleteDepository(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteDepository"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var del DeleteDepForm
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	if err := rlib.DeleteDepository(r.Context(), del.ID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// GetDepository returns the requested assessment
// wsdoc {
//  @Title  Save Depository
//	@URL /v1/dep/:BUI/:DEPID
//  @Method  GET
//	@Synopsis Update the information on a Depository with the supplied data
//  @Description  This service updates Depository :DEPID with the information supplied. All fields must be supplied.
//	@Input DepositoryGridSave
//  @Response SvcStatusResponse
// wsdoc }
func saveDepository(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveDepository"
	var (
		foo DepositoryGridSave
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

	var a rlib.Depository
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, foo.Record.BUD)
		rlib.Ulog("%s", e.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	if len(a.Name) == 0 {
		e := fmt.Errorf("%s: Required field, Name, is blank", funcname)
		SvcErrorReturn(w, e, funcname)
		return
	}

	adup, _ := rlib.GetDepositoryByName(r.Context(), a.BID, a.Name)
	if a.Name == adup.Name && a.DEPID != adup.DEPID {
		e := fmt.Errorf("%s: A Depository with the name %s already exists", funcname, a.Name)
		SvcErrorReturn(w, e, funcname)
		return
	}

	adup, _ = rlib.GetDepositoryByLID(r.Context(), a.BID, a.LID)
	if a.LID == adup.LID && a.DEPID != adup.DEPID {
		l, _ := rlib.GetLedger(r.Context(), a.LID)
		e := fmt.Errorf("%s: A Depository for Account %s (%s) already exists", funcname, l.GLNumber, l.Name)
		SvcErrorReturn(w, e, funcname)
		return
	}

	if a.DEPID == 0 && d.ID == 0 {
		// This is a new AR
		fmt.Printf(">>>> NEW DEPOSITORY IS BEING ADDED\n")
		_, err = rlib.InsertDepository(r.Context(), &a)
	} else {
		// update existing record
		fmt.Printf("Updating existing Depository: %d\n", a.DEPID)
		err = rlib.UpdateDepository(r.Context(), &a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving depository (DEPID=%d\n: %s", funcname, a.DEPID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// GetDepository returns the requested assessment
// wsdoc {
//  @Title  Get Depository
//	@URL /v1/dep/:BUI/:DEPID
//  @Method  GET
//	@Synopsis Get information on a Depository
//  @Description  Return all fields for assessment :DEPID
//	@Input WebGridSearchRequest
//  @Response DepositoryGetResponse
// wsdoc }
func getDepository(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getDepository"
	var (
		g   DepositoryGetResponse
		whr = fmt.Sprintf("Depository.DEPID=%d", d.ID)
	)

	fmt.Printf("entered %s\n", funcname)

	depQuery := `
	SELECT
		{{.SelectClause}}
	FROM Depository
	LEFT JOIN GLAccount on GLAccount.LID=Depository.LID
	WHERE {{.WhereClause}};`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(depSearchSelectQueryFields, ","),
		"WhereClause":  whr,
	}

	qry := rlib.RenderSQLQuery(depQuery, qc)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var q DepositoryGrid
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		q, err = depGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		q.Recid = q.DEPID
		g.Record = q
	}
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
