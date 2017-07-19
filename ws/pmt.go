package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// PaymentTypeGrid contains the data from PaymentType that is targeted to the UI Grid that displays
// a list of PaymentType structs
type PaymentTypeGrid struct {
	Recid       int64 `json:"recid"`
	PMTID       int64
	BID         int64
	BUD         rlib.XJSONBud
	Name        string
	Description string
	LastModTime time.Time
	LastModBy   int64
}

// PaymentTypeSearchResponse is a response string to the search request for PaymentType records
type PaymentTypeSearchResponse struct {
	Status  string            `json:"status"`
	Total   int64             `json:"total"`
	Records []PaymentTypeGrid `json:"records"`
}

// PaymentTypeSaveForm is a struct to handle direct inputs from the form
type PaymentTypeSaveForm struct {
	Recid       int64 `json:"recid"`
	PMTID       int64
	BID         int64
	Name        string
	Description string
	LastModTime time.Time
	LastModBy   int64
}

// PaymentTypeSaveOther is a struct to handle the UI list box selections
type PaymentTypeSaveOther struct {
	BUD rlib.W2uiHTMLSelect
}

// SavePaymentTypeInput is the input data format for a Save command
type SavePaymentTypeInput struct {
	Status   string              `json:"status"`
	Recid    int64               `json:"recid"`
	FormName string              `json:"name"`
	Record   PaymentTypeSaveForm `json:"record"`
}

// SavePaymentTypeOther is the input data format for the "other" data on the Save command
type SavePaymentTypeOther struct {
	Status string               `json:"status"`
	Recid  int64                `json:"recid"`
	Name   string               `json:"name"`
	Record PaymentTypeSaveOther `json:"record"`
}

// PaymentTypeGetResponse is the response to a GetPaymentType request
type PaymentTypeGetResponse struct {
	Status string          `json:"status"`
	Record PaymentTypeGrid `json:"record"`
}

// DeletePmtForm used to delete record from database
type DeletePmtForm struct {
	ID int64
}

// pmtRowScan scans a result from sql row and dump it in a PaymentTypeGrid struct
func pmtRowScan(rows *sql.Rows, q PaymentTypeGrid) (PaymentTypeGrid, error) {
	err := rows.Scan(&q.PMTID, &q.Name, &q.Description, &q.LastModTime, &q.LastModBy)
	return q, err
}

// SvcHandlerPaymentType formats a complete data record for an assessment for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the PMTID as follows:
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerPaymentType(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "SvcHandlerPaymentType"
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  PMTID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerPaymentTypes(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				err = fmt.Errorf("PaymentTypeID is required but was not specified")
				SvcGridErrorReturn(w, err, funcname)
				return
			}
			getPaymentType(w, r, d)
		}
		break
	case "save":
		savePaymentType(w, r, d)
		break
	case "delete":
		deletePaymentType(w, r, d)
		break
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

var pmtSearchFieldMap = selectQueryFieldMap{
	"PMTID":       {"PaymentType.PMTID"},
	"Name":        {"PaymentType.Name"},
	"Description": {"PaymentType.Description"},
	"LastModTime": {"PaymentType.LastModTime"},
	"LastModBy":   {"PaymentType.LastModBy"},
}

// which fields needs to be fetch to satisfy the struct
var pmtSearchSelectQueryFields = selectQueryFields{
	"PaymentType.PMTID",
	"PaymentType.Name",
	"PaymentType.Description",
	"PaymentType.LastModTime",
	"PaymentType.LastModBy",
}

// SvcSearchHandlerPaymentTypes generates a report of all PaymentTypes defined business d.BID
// wsdoc {
//  @Title  Search PaymentType
//	@URL /v1/pmts/:BUI
//  @Method  POST
//	@Synopsis Search PaymentTypes
//  @Descr  Search all PaymentType and return those that match the Search Logic.
//  @Descr  The search criteria includes start and stop dates of interest.
//	@Input WebGridSearchRequest
//  @Response PaymentTypeSearchResponse
// wsdoc }
func SvcSearchHandlerPaymentTypes(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcSearchHandlerPaymentTypes"
		g        PaymentTypeSearchResponse
		err      error
		order    = "PMTID ASC" // default ORDER
		whr      = fmt.Sprintf("BID=%d", d.BID)
	)

	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	_, orderClause := GetSearchAndSortSQL(d, pmtSearchFieldMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	pmtQuery := `
	SELECT
		{{.SelectClause}}
	FROM PaymentType
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := queryClauses{
		"SelectClause": strings.Join(pmtSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := renderSQLQuery(pmtQuery, qc)
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
	pmtQueryWithLimit := pmtQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := renderSQLQuery(pmtQueryWithLimit, qc)
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
		var q PaymentTypeGrid
		q.Recid = i
		q.BID = d.BID
		q.BUD = getBUDFromBIDList(q.BID)

		q, err = pmtRowScan(rows, q)
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

// deletePaymentType deletes a payment type from the database
// wsdoc {
//  @Title  Delete Payment Type
//	@URL /v1/pmt/:BUI/:RAID
//  @Method  POST
//	@Synopsis Delete a Payment Type
//  @Desc  This service deletes a PaymentType.
//	@Input DeletePmtForm
//  @Response SvcStatusResponse
// wsdoc }
func deletePaymentType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "deleteDepository"
		del      DeletePmtForm
		err      error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	if err = rlib.DeletePaymentType(del.ID); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(w)
}

// GetPaymentType returns the requested assessment
// wsdoc {
//  @Title  Save PaymentType
//	@URL /v1/pmt/:BUI/:PMTID
//  @Method  GET
//	@Synopsis Update the information on a PaymentType with the supplied data
//  @Description  This service updates PaymentType :PMTID with the information supplied. All fields must be supplied.
//	@Input SavePaymentTypeInput
//  @Response SvcStatusResponse
// wsdoc }
func savePaymentType(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "savePaymentType"
		foo      SavePaymentTypeInput
		bar      SavePaymentTypeOther
		err      error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	data := []byte(d.data)

	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	if err := json.Unmarshal(data, &bar); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	var a rlib.PaymentType
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.Record.BUD.ID]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, bar.Record.BUD.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	if len(a.Name) == 0 {
		e := fmt.Errorf("%s: Required field, Name, is blank", funcname)
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	var adup rlib.PaymentType
	rlib.GetPaymentTypeByName(a.BID, a.Name, &adup)
	if a.Name == adup.Name {
		e := fmt.Errorf("%s: A PaymentType with the name %s already exists", funcname, a.Name)
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	if a.PMTID == 0 && d.ID == 0 {
		// This is a new AR
		err = rlib.InsertPaymentType(&a)
	} else {
		// update existing record
		fmt.Printf("Updating existing Payment Type: %d\n", a.PMTID)
		err = rlib.UpdatePaymentType(&a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving Payment Type : %s", funcname, a.Name)
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponse(w)
}

// GetPaymentType returns the requested assessment
// wsdoc {
//  @Title  Get Payment Type
//	@URL /v1/pmt/:BUI/:PMTID
//  @Method  GET
//	@Synopsis Get information on a PaymentType
//  @Description  Return all fields for assessment :PMTID
//	@Input WebGridSearchRequest
//  @Response PaymentTypeGetResponse
// wsdoc }
func getPaymentType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "getPaymentType"
		g        PaymentTypeGetResponse
	)

	fmt.Printf("entered %s\n", funcname)
	var a rlib.PaymentType
	rlib.GetPaymentType(d.ID, &a)
	if a.PMTID > 0 {
		var gg PaymentTypeGrid
		rlib.MigrateStructVals(&a, &gg)

		gg.BUD = getBUDFromBIDList(gg.BID)

		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
