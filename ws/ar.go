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

// ARSendForm is a structure specifically for the UI. It will be
// automatically populated from an rlib.AR struct
type ARSendForm struct {
	Recid               int64 `json:"recid"` // this is to support the w2ui form
	ARID                int64
	BUD                 rlib.XJSONBud
	BID                 int64
	Name                string
	ARType              int64
	DebitLID            int64
	DebitLedgerName     string
	CreditLID           int64
	CreditLedgerName    string
	Description         string
	DtStart             rlib.JSONDate
	DtStop              rlib.JSONDate
	FLAGS               uint64
	AutoPopulateToNewRA bool
	raRequired          int
	PriorToRAStart      bool    // is it ok to charge prior to RA start
	PriorToRAStop       bool    // is it ok to charge after RA stop
	ApplyRcvAccts       bool    // if true, mark the receipt as fully paid based on RcvAccts
	RAIDrqd             bool    // if true, it will require receipts to supply a RAID
	IsRentAR            bool    // if true, then it represents Rent AR
	DefaultAmount       float64 // default amount for this account rule
	LastModTime         rlib.JSONDateTime
	LastModBy           int64
	CreateTS            rlib.JSONDateTime
	CreateBy            int64
}

// ARSaveForm is a structure specifically for the return value from w2ui.
// Data does not always come back in the same format it was sent. For example,
// values from dropdown lists come back in the form of a rlib.W2uiHTMLSelect struct.
// So, we break up the ingest into 2 parts. First, we read back the fields that look
// just like the xxxSendForm -- this is what is in xxxSaveForm. Then we readback
// the data that has changed, which is in the xxxSaveOther struct.  All this data
// is merged into the appropriate database structure using MigrateStructData.
type ARSaveForm struct {
	Recid               int64 `json:"recid"` // this is to support the w2ui form
	ARID                int64
	BID                 int64
	BUD                 rlib.XJSONBud
	CreditLID           int64
	DebitLID            int64
	ARType              int64
	Name                string
	Description         string
	DtStart             rlib.JSONDate
	DtStop              rlib.JSONDate
	PriorToRAStart      bool // is it ok to charge prior to RA start
	PriorToRAStop       bool // is it ok to charge after RA stop
	ApplyRcvAccts       bool
	RAIDrqd             bool
	DefaultAmount       float64
	AutoPopulateToNewRA bool
	IsRentAR            bool
}

// PrARGrid is a structure specifically for the UI Grid.
type PrARGrid struct {
	Recid            int64 `json:"recid"` // this is to support the w2ui form
	ARID             int64
	BID              int64
	BUD              rlib.XJSONBud
	Name             string
	ARType           int64
	DebitLID         int64
	DebitLedgerName  string
	CreditLID        int64
	CreditLedgerName string
	Description      string
	DtStart          rlib.JSONDate
	DtStop           rlib.JSONDate
}

// SaveARInput is the input data format for a Save command
type SaveARInput struct {
	Status   string     `json:"status"`
	Recid    int64      `json:"recid"`
	FormName string     `json:"name"`
	Record   ARSaveForm `json:"record"`
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
	err := rows.Scan(&q.ARID, &q.Name, &q.ARType, &q.DebitLID, &q.DebitLedgerName, &q.CreditLID, &q.CreditLedgerName, &q.Description, &q.DtStart, &q.DtStop)
	return q, err
}

// which fields needs to be fetched for SQL query for receipts grid
var arFieldsMap = rlib.SelectQueryFieldMap{
	"ARID":             {"AR.ARID"},
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
var arQuerySelectFields = rlib.SelectQueryFields{
	"AR.ARID",
	"AR.Name",
	"AR.ARType",
	"AR.DebitLID",
	"CONCAT(debitQuery.GLNumber,' (',debitQuery.Name,')') as DebitLedgerName",
	"AR.CreditLID",
	"CONCAT(creditQuery.GLNumber,' (',creditQuery.Name,')') as CreditLedgerName",
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
	const funcname = "SvcSearchHandlerARs"
	fmt.Printf("Entered %s\n", funcname)

	switch d.wsSearchReq.Cmd {
	case "get":
		getARGrid(w, r, d)
		break
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
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
	const funcname = "getARGrid"
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

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(arQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(arQuery, qc)
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
	arQueryWithLimit := arQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(arQueryWithLimit, qc)
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
		var q PrARGrid
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(d.BID)

		q, err = arGridRowScan(rows, q)
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

	// error check
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
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
	const funcname = "SvcFormHandlerAR"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	if d.ARID, err = SvcExtractIDFromURI(r.RequestURI, "ARID", 3, w); err != nil {
		SvcErrorReturn(w, err, funcname)
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
		SvcErrorReturn(w, err, funcname)
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
	const funcname = "saveARForm"
	var (
		foo SaveARInput
		a   rlib.AR
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err := json.Unmarshal(data, &foo); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// migrate foo.Record data to a struct's fields
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	fmt.Printf("saveAR - first migrate: a = %#v\n", a)

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, foo.Record.BUD)
		SvcErrorReturn(w, e, funcname)
		return
	}

	// get PriorToRAStart and PriorToRAStop values and accordingly get RARequired field value
	formBoolMap := [2]bool{foo.Record.PriorToRAStart, foo.Record.PriorToRAStop}
	for raReq, boolMap := range raRequiredMap {
		if boolMap == formBoolMap {
			a.RARequired = int64(raReq)
			break
		}
	}

	a.FLAGS &= ^uint64(0x4 & 0x2 & 0x1) // 1<<0 and 1<< 1 and 1<<2:  these are the three flags that can be set.  Assume we turn them off
	if foo.Record.ApplyRcvAccts {
		a.FLAGS |= 0x1
	}
	if foo.Record.AutoPopulateToNewRA {
		a.FLAGS |= 0x2
	}
	if foo.Record.RAIDrqd && a.ARType == rlib.ARRECEIPT {
		a.FLAGS |= 0x4
	}
	if foo.Record.IsRentAR { // IsRentAR - 1<<4
		a.FLAGS |= 0x10
	}
	rlib.Console("=============>>>>>>>>>> a.FLAGS = %x\n", a.FLAGS)

	// Ensure that the supplied data is valid
	e := bizlogic.ValidateAcctRule(r.Context(), &a)
	if len(e) > 0 {
		SvcErrListReturn(w, e, funcname)
		return
	}

	// save or update
	if a.ARID == 0 && d.ARID == 0 {
		// This is a new AR
		fmt.Printf(">>>> NEW RECEIPT IS BEING ADDED\n")
		_, err = rlib.InsertAR(r.Context(), &a)
	} else {
		// update existing record
		fmt.Printf("Updating existing AR: %d\n", a.ARID)
		err = rlib.UpdateAR(r.Context(), &a)
	}
	if err != nil {
		e := fmt.Errorf("Error saving receipt (ARID=%d\n: %s", d.ARID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.ARID)
}

// which fields needs to be fetched for SQL query for receipts grid
var getARQuerySelectFields = rlib.SelectQueryFields{
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
	"AR.DefaultAmount",
	"AR.FLAGS",
	"AR.LastModTime",
	"AR.LastModBy",
	"AR.CreateTS",
	"AR.CreateBy",
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
	const funcname = "getARForm"
	var (
		g   GetARResponse
		err error
	)
	fmt.Printf("entered %s\n", funcname)

	arQuery := `
	SELECT {{.SelectClause}}
	FROM AR
	INNER JOIN GLAccount as debitQuery on AR.DebitLID=debitQuery.LID
	INNER JOIN GLAccount as creditQuery on AR.CreditLID=creditQuery.LID
 	WHERE {{.WhereClause}};`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(getARQuerySelectFields, ","),
		"WhereClause":  fmt.Sprintf("AR.BID=%d AND AR.ARID=%d", d.BID, d.ARID),
	}

	// get formatted query with substitution of select, where, order clause
	q := rlib.RenderSQLQuery(arQuery, qc)
	fmt.Printf("db query = %s\n", q)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(q)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var gg ARSendForm

		gg.BID = d.BID
		gg.BUD = rlib.GetBUDFromBIDList(d.BID)

		rlib.Console("gg.BUD = %s\n", gg.BUD)

		err = rows.Scan(&gg.ARID, &gg.Name, &gg.ARType, &gg.DebitLID, &gg.DebitLedgerName,
			&gg.CreditLID, &gg.CreditLedgerName, &gg.Description, &gg.DtStart, &gg.DtStop,
			&gg.raRequired, &gg.DefaultAmount, &gg.FLAGS, &gg.LastModTime, &gg.LastModBy, &gg.CreateTS, &gg.CreateBy)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		// according to RARequired map, fill out PriorToRAStart, PriorToRAStop values
		raReqMappedVal := raRequiredMap[gg.raRequired]
		gg.PriorToRAStart = raReqMappedVal[0]
		gg.PriorToRAStop = raReqMappedVal[1]

		if gg.FLAGS&0x1 != 0 {
			gg.ApplyRcvAccts = true
		}
		if gg.FLAGS&0x2 != 0 {
			gg.AutoPopulateToNewRA = true
		}
		if gg.FLAGS&0x4 != 0 {
			gg.RAIDrqd = true
		}
		if gg.FLAGS&0x10 != 0 {
			gg.IsRentAR = true
		}

		g.Record = gg
		rlib.Console("g.Record.BUD = %s\n", g.Record.BUD)
	}

	// error check
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}

// deleteAR request delete AR from database
// wsdoc {
//  @Title  Delete AR
//	@URL /v1/ars/:BUI/:ARID
//  @Method  DELETE
//	@Synopsis Delete record for a AR
//  @Description  Delete record from database ars :ARID
//	@Input WebGridSearchRequest
//  @Response SvcWriteSuccessResponse
// wsdoc }
func deleteARForm(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteARForm"
	var (
		del DeleteARForm
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	if err := rlib.DeleteAR(r.Context(), del.ARID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// ListedAR is struct to list down individual account rule record
type ListedAR struct {
	BID  int64  `json:"BID"`
	ARID int64  `json:"ARID"` // Account Rule ID
	Name string `json:"Name"` // Account rule name
}

// ARsListResponse is the response to list down all account rules
type ARsListResponse struct {
	Status  string     `json:"status"`
	Total   int64      `json:"total"`
	Records []ListedAR `json:"records"`
}

// ARsListRequestByFLAGS is the request struct for listing down account rules by FLAGS
type ARsListRequestByFLAGS struct {
	FLAGS uint64 `json:"FLAGS"`
}

// ARsListRequestType represents for which type of request to list down ARs
type ARsListRequestType struct {
	Type string `json:"type"`
}

// SvcARsList generates a list of all ARs with respect of business id specified by d.BID
// wsdoc {
//  @Title Get list of ARs
//  @URL /v1/arslist/:BUI
//  @Method  GET
//  @Synopsis Get ARs list
//  @Description Get all Account rules list for the requested business
//  @Input WebGridSearchRequest
//  @Response ARsListResponse
// wsdoc }
func SvcARsList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcARsList"
	var (
		g   ARsListResponse
		foo ARsListRequestType
	)
	fmt.Printf("Entered %s\n", funcname)

	if r.Method != "POST" {
		err := fmt.Errorf("Only POST method is allowed")
		SvcErrorReturn(w, err, funcname)
		return
	}

	data := []byte(d.data)

	// get the type first
	err := json.Unmarshal(data, &foo)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	switch foo.Type {
	case "FLAGS":
		bar := ARsListRequestByFLAGS{}
		err = json.Unmarshal(data, &bar)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		// we should get ar by flag value directly instead of parsing flag value from
		// requested a bit only, in case client wants to fetch ars based on multiple flags bit
		// client should request with final flag value only
		if !bizlogic.IsValidARFlag(bar.FLAGS) {
			err := fmt.Errorf("FLAGS value is invalid: %d", bar.FLAGS)
			SvcErrorReturn(w, err, funcname)
			return
		}

		// get account rules by FLAGS integer representation from binary value
		m, err := rlib.GetARsByFLAGS(r.Context(), d.BID, bar.FLAGS)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		// append records in ascending order
		var arList []ListedAR
		for _, ar := range m {
			arList = append(arList, ListedAR{BID: ar.BID, ARID: ar.ARID, Name: ar.Name})
		}

		// sort based on name, needs version 1.8 later of golang
		sort.Slice(arList, func(i, j int) bool { return arList[i].Name < arList[j].Name })

		g.Records = arList
		g.Total = int64(len(g.Records))
		g.Status = "success"
		SvcWriteResponse(d.BID, &g, w)
	default:
		err := fmt.Errorf("%s: Unhandled %s command", funcname, foo.Type)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	}
}
