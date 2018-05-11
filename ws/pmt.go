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

// PaymentTypeGrid contains the data from PaymentType that is targeted to the UI Grid that displays
// a list of PaymentType structs
type PaymentTypeGrid struct {
	Recid       int64 `json:"recid"`
	PMTID       int64
	BID         int64
	BUD         rlib.XJSONBud
	Name        string
	Description string
	LastModTime rlib.JSONDateTime
	LastModBy   int64
	CreateTS    rlib.JSONDateTime
	CreateBy    int64
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
	BUD         rlib.XJSONBud
	Name        string
	Description string
}

// SavePaymentTypeInput is the input data format for a Save command
type SavePaymentTypeInput struct {
	Status   string              `json:"status"`
	Recid    int64               `json:"recid"`
	FormName string              `json:"name"`
	Record   PaymentTypeSaveForm `json:"record"`
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
	err := rows.Scan(&q.PMTID, &q.Name, &q.Description, &q.LastModTime, &q.LastModBy, &q.CreateTS, &q.CreateBy)
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
	const funcname = "SvcHandlerPaymentType"
	var (
		err error
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
				SvcErrorReturn(w, err, funcname)
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
		SvcErrorReturn(w, err, funcname)
		return
	}
}

var pmtSearchFieldMap = rlib.SelectQueryFieldMap{
	"PMTID":       {"PaymentType.PMTID"},
	"Name":        {"PaymentType.Name"},
	"Description": {"PaymentType.Description"},
	"LastModTime": {"PaymentType.LastModTime"},
	"LastModBy":   {"PaymentType.LastModBy"},
	"CreateTS":    {"PaymentType.CreateTS"},
	"CreateBy":    {"PaymentType.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var pmtSearchSelectQueryFields = rlib.SelectQueryFields{
	"PaymentType.PMTID",
	"PaymentType.Name",
	"PaymentType.Description",
	"PaymentType.LastModTime",
	"PaymentType.LastModBy",
	"PaymentType.CreateTS",
	"PaymentType.CreateBy",
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
	const funcname = "SvcSearchHandlerPaymentTypes"
	var (
		g     PaymentTypeSearchResponse
		err   error
		order = "PMTID ASC" // default ORDER
		whr   = fmt.Sprintf("BID=%d", d.BID)
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

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(pmtSearchSelectQueryFields, ","),
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
		var q PaymentTypeGrid
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		q, err = pmtRowScan(rows, q)
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

	if err = rlib.DeletePaymentType(r.Context(), del.ID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
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
	const funcname = "savePaymentType"
	var (
		foo SavePaymentTypeInput
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

	var a rlib.PaymentType
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

	var adup rlib.PaymentType
	_ = rlib.GetPaymentTypeByName(r.Context(), a.BID, a.Name, &adup)
	if a.Name == adup.Name && a.PMTID != adup.PMTID {
		e := fmt.Errorf("%s: A PaymentType with the name %s already exists", funcname, a.Name)
		SvcErrorReturn(w, e, funcname)
		return
	}

	if a.PMTID == 0 && d.ID == 0 {
		// This is a new AR
		_, err = rlib.InsertPaymentType(r.Context(), &a)
	} else {
		// update existing record
		fmt.Printf("Updating existing Payment Type: %d\n", a.PMTID)
		err = rlib.UpdatePaymentType(r.Context(), &a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving Payment Type : %s", funcname, a.Name)
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
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
	const funcname = "getPaymentType"
	var (
		g PaymentTypeGetResponse
	)

	fmt.Printf("entered %s\n", funcname)
	var a rlib.PaymentType
	_ = rlib.GetPaymentType(r.Context(), d.ID, &a)
	if a.PMTID > 0 {
		var gg PaymentTypeGrid
		rlib.MigrateStructVals(&a, &gg)

		gg.BUD = rlib.GetBUDFromBIDList(gg.BID)

		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
