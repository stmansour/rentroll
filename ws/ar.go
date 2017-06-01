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

// ARSendForm is a structure specifically for the UI. It will be
// automatically populated from an rlib.AR struct
type ARSendForm struct {
	Recid            int64 `json:"recid"` // this is to support the w2ui form
	ARID             int64
	BID              rlib.XJSONBud
	Name             string
	ARType           int64
	DebitLID         int64
	DebitLedgerName  string
	CreditLID        int64
	CreditLedgerName string
	Description      string
	DtStart          rlib.JSONTime
	DtStop           rlib.JSONTime
	raRequired       int
	PriorToRAStart   bool // is it ok to charge prior to RA start
	PriorToRAStop    bool // is it ok to charge after RA stop
	LastModTime      rlib.JSONTime
	LastModBy        int64
}

// ARSaveForm is a structure specifically for the return value from w2ui.
// Data does not always come back in the same format it was sent. For example,
// values from dropdown lists come back in the form of a rlib.W2uiHTMLSelect struct.
// So, we break up the ingest into 2 parts. First, we read back the fields that look
// just like the xxxSendForm -- this is what is in xxxSaveForm. Then we readback
// the data that has changed, which is in the xxxSaveOther struct.  All this data
// is merged into the appropriate database structure using MigrateStructData.
type ARSaveForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	ARID           int64
	Name           string
	Description    string
	DtStart        rlib.JSONTime
	DtStop         rlib.JSONTime
	PriorToRAStart bool // is it ok to charge prior to RA start
	PriorToRAStop  bool // is it ok to charge after RA stop
	LastModTime    rlib.JSONTime
	LastModBy      int64
}

// ARSaveOther is a struct to handle the UI list box selections
type ARSaveOther struct {
	BID       rlib.W2uiHTMLSelect
	CreditLID rlib.W2uiHTMLSelect
	DebitLID  rlib.W2uiHTMLSelect
	ARType    rlib.W2uiHTMLSelect
}

// PrARGrid is a structure specifically for the UI Grid.
type PrARGrid struct {
	Recid            int64 `json:"recid"` // this is to support the w2ui form
	ARID             int64
	BID              int64
	Name             string
	ARType           int64
	DebitLID         int64
	DebitLedgerName  string
	CreditLID        int64
	CreditLedgerName string
	Description      string
	DtStart          rlib.JSONTime
	DtStop           rlib.JSONTime
}

// SaveARInput is the input data format for a Save command
type SaveARInput struct {
	Status   string     `json:"status"`
	Recid    int64      `json:"recid"`
	FormName string     `json:"name"`
	Record   ARSaveForm `json:"record"`
}

// SaveAROther is the input data format for the "other" data on the Save command
type SaveAROther struct {
	Status string      `json:"status"`
	Recid  int64       `json:"recid"`
	Name   string      `json:"name"`
	Record ARSaveOther `json:"record"`
}

// SearchARsResponse is a response string to the search request for receipts
type SearchARsResponse struct {
	Status  string     `json:"status"`
	Total   int64      `json:"total"`
	Records []PrARGrid `json:"records"`
}

// GetARResponse is the response to a GetAR request
type GetARResponse struct {
	Status string     `json:"status"`
	Record ARSendForm `json:"record"`
}

// DeleteARForm holds ARID to delete it
type DeleteARForm struct {
	ARID int64
}

// arGridRowScan scans a result from sql row and dump it in a PrARGrid struct
func arGridRowScan(rows *sql.Rows, q PrARGrid) (PrARGrid, error) {
	err := rows.Scan(&q.ARID, &q.BID, &q.Name, &q.ARType, &q.DebitLID, &q.DebitLedgerName, &q.CreditLID, &q.CreditLedgerName, &q.Description, &q.DtStart, &q.DtStop)
	return q, err
}

// which fields needs to be fetched for SQL query for receipts grid
var arFieldsMap = selectQueryFieldMap{
	"ARID":             {"AR.ARID"},
	"BID":              {"AR.BID"},
	"Name":             {"AR.Name"},
	"ARType":           {"AR.ARType"},
	"DebitLID":         {"AR.DebitLID"},
	"DebitLedgerName":  {"debitQuery.Name"},
	"CreditLID":        {"AR.CreditLID"},
	"CreditLedgerName": {"creditQuery.Name"},
	"Description":      {"AR.Description"},
	"DtStart":          {"AR.DtStart"},
	"DtStop":           {"AR.DtStop"},
}

// which fields needs to be fetched for SQL query for receipts grid
var arQuerySelectFields = selectQueryFields{
	"AR.ARID",
	"AR.BID",
	"AR.Name",
	"AR.ARType",
	"AR.DebitLID",
	"debitQuery.Name as DebitLedgerName",
	"AR.CreditLID",
	"creditQuery.Name as CreditLedgerName",
	"AR.Description",
	"AR.DtStart",
	"AR.DtStop",
}

// SvcSearchHandlerARs generates a report of all ARs defined business d.BID
// wsdoc {
//  @Title  Search Account Rules
//	@URL /v1/ars/:BUI
//  @Method  POST
//	@Synopsis Search Account Rules
//  @Description  Search all ARs and return those that match the Search Logic.
//  @Desc By default, the search is made for receipts from "today" to 31 days prior.
//	@Input WebGridSearchRequest
//  @Response SearchARsResponse
// wsdoc }
func SvcSearchHandlerARs(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcSearchHandlerARs"
	fmt.Printf("Entered %s\n", funcname)

	switch d.wsSearchReq.Cmd {
	case "get":
		getARGrid(w, r, d)
		break
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

// getARGrid returns a list of ARs for w2ui grid
// wsdoc {
//  @Title  list ARs
//	@URL /v1/ars/:BUI
//  @Method  GET
//	@Synopsis Get Account Rules
//  @Description  Get all ARs associated with BID
//  @Desc By default, the search is made for receipts from "today" to 31 days prior.
//	@Input WebGridSearchRequest
//  @Response SearchARsResponse
// wsdoc }
func getARGrid(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "getARGrid"
	var (
		err error
		g   SearchARsResponse
	)

	order := `AR.Name ASC` // default ORDER
	whr := fmt.Sprintf("AR.BID=%d", d.BID)

	// get where clause and order clause for sql query
	_, orderClause := GetSearchAndSortSQL(d, arFieldsMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	arQuery := `
	SELECT
		{{.SelectClause}}
	FROM AR
	INNER JOIN GLAccount as debitQuery on AR.DebitLID=debitQuery.LID
	INNER JOIN GLAccount as creditQuery on AR.CreditLID=creditQuery.LID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := queryClauses{
		"SelectClause": strings.Join(arQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := renderSQLQuery(arQuery, qc)
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
	arQueryWithLimit := arQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := renderSQLQuery(arQueryWithLimit, qc)
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
		var q PrARGrid
		q.Recid = i

		q, err = arGridRowScan(rows, q)
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

	// error check
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// SvcFormHandlerAR formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the ARID as follows:
//           0    1     2   3
// uri 		/v1/receipt/BUI/ARID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerAR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcFormHandlerAR"
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)
	if d.ARID, err = SvcExtractIDFromURI(r.RequestURI, "ARID", 3, w); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	fmt.Printf("Request: %s:  BID = %d,  ID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ARID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getARForm(w, r, d)
		break
	case "save":
		saveARForm(w, r, d)
		break
	case "delete":
		deleteARForm(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

// saveARForm returns the requested receipt
// wsdoc {
//  @Title  Save AR
//	@URL /v1/ars/:BUI/:ARID
//  @Method  GET
//	@Synopsis Save a AR
//  @Desc  This service saves a AR.  If :ARID exists, it will
//  @Desc  be updated with the information supplied. All fields must
//  @Desc  be supplied. If ARID is 0, then a new receipt is created.
//	@Input SaveARInput
//  @Response SvcStatusResponse
// wsdoc }
func saveARForm(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "saveARForm"
		foo      SaveARInput
		bar      SaveAROther
		a        rlib.AR
		err      error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err := json.Unmarshal(data, &foo); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	if err := json.Unmarshal(data, &bar); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	// migrate foo.Record data to a struct's fields
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	fmt.Printf("saveAR - first migrate: a = %#v\n", a)

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.Record.BID.ID]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, bar.Record.BID.ID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	a.CreditLID, ok = rlib.StringToInt64(bar.Record.CreditLID.ID) // CreditLID has drop list
	if !ok {
		e := fmt.Errorf("%s: invalid CreditLID value: %s", funcname, bar.Record.CreditLID.ID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	a.DebitLID, ok = rlib.StringToInt64(bar.Record.DebitLID.ID) // DebitLID has drop list
	if !ok {
		e := fmt.Errorf("%s: invalid DebitLID value: %s", funcname, bar.Record.DebitLID.ID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	a.ARType, ok = rlib.StringToInt64(bar.Record.ARType.ID) // ArType has drop list
	if !ok {
		e := fmt.Errorf("%s: Invalid ARType value: %s", funcname, bar.Record.ARType.ID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	fmt.Printf("saveAR - second migrate: a = %#v\n", a)

	// get PriorToRAStart and PriorToRAStop values and accordingly get RARequired field value
	formBoolMap := [2]bool{foo.Record.PriorToRAStart, foo.Record.PriorToRAStop}
	for raReq, boolMap := range raRequiredMap {
		if boolMap == formBoolMap {
			a.RARequired = int64(raReq)
			break
		}
	}

	// save or update
	if a.ARID == 0 && d.ARID == 0 {
		// This is a new AR
		fmt.Printf(">>>> NEW RECEIPT IS BEING ADDED\n")
		_, err = rlib.InsertAR(&a)
	} else {
		// update existing record
		fmt.Printf("Updating existing AR: %d\n", a.ARID)
		err = rlib.UpdateAR(&a)
	}
	if err != nil {
		e := fmt.Errorf("Error saving receipt (ARID=%d\n: %s", d.ARID, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(w, a.ARID)
}

// which fields needs to be fetched for SQL query for receipts grid
var getARQuerySelectFields = selectQueryFields{
	"AR.ARID",
	"AR.Name",
	"AR.ARType",
	"AR.DebitLID",
	"debitQuery.Name as DebitLedgerName",
	"AR.CreditLID",
	"creditQuery.Name as CreditLedgerName",
	"AR.Description",
	"AR.DtStart",
	"AR.DtStop",
	"AR.RARequired",
}

// for what RARequired value, prior and after value are
var raRequiredMap = map[int][2]bool{
	0: {false, false}, // during
	1: {true, false},  // prior or during
	2: {false, true},  // after or during
	3: {true, true},   // after or during or prior
}

// getARForm returns the requested ars
// wsdoc {
//  @Title  Get AR
//	@URL /v1/ars/:BUI/:ARID
//  @Method  GET
//	@Synopsis Get information on a AR
//  @Description  Return all fields for ars :ARID
//	@Input WebGridSearchRequest
//  @Response GetARResponse
// wsdoc }
func getARForm(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "getARForm"
		g        GetARResponse
		err      error
	)
	fmt.Printf("entered %s\n", funcname)

	arQuery := `
	SELECT
		{{.SelectClause}}
	FROM AR
	INNER JOIN GLAccount as debitQuery on AR.DebitLID=debitQuery.LID
	INNER JOIN GLAccount as creditQuery on AR.CreditLID=creditQuery.LID
	WHERE {{.WhereClause}};`

	qc := queryClauses{
		"SelectClause": strings.Join(getARQuerySelectFields, ","),
		"WhereClause":  fmt.Sprintf("AR.BID=%d AND AR.ARID=%d", d.BID, d.ARID),
	}

	// get formatted query with substitution of select, where, order clause
	q := renderSQLQuery(arQuery, qc)
	fmt.Printf("db query = %s\n", q)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(q)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var gg ARSendForm

		gg.BID = getBUDFromBIDList(d.BID)

		err = rows.Scan(&gg.ARID, &gg.Name, &gg.ARType, &gg.DebitLID, &gg.DebitLedgerName, &gg.CreditLID, &gg.CreditLedgerName, &gg.Description, &gg.DtStart, &gg.DtStop, &gg.raRequired)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		// according to RARequired map, fill out PriorToRAStart, PriorToRAStop values
		raReqMappedVal := raRequiredMap[gg.raRequired]
		gg.PriorToRAStart = raReqMappedVal[0]
		gg.PriorToRAStop = raReqMappedVal[1]
		g.Record = gg
	}

	// error check
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// deleteAR request delete AR from database
// wsdoc {
//  @Title  Get AR
//	@URL /v1/ars/:BUI/:ARID
//  @Method  DELETE
//	@Synopsis Delete record for a AR
//  @Description  Delete record from database ars :ARID
//	@Input WebGridSearchRequest
//  @Response SvcWriteSuccessResponse
// wsdoc }
func deleteARForm(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "deleteARForm"
		del      DeleteARForm
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	if err := rlib.DeleteAR(del.ARID); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(w)
}
