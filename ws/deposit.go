package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
)

// DepositGrid contains the data from Deposit that is targeted to the UI Grid that displays
// a list of Deposit structs
type DepositGrid struct {
	Recid       int64 `json:"recid"`
	DID         int64
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

// DepositSearchResponse is a response string to the search request for Deposit records
type DepositSearchResponse struct {
	Status  string        `json:"status"`
	Total   int64         `json:"total"`
	Records []DepositGrid `json:"records"`
}

// DepositSaveForm contains the data from Deposit FORM that is targeted to the UI Form that displays
// a list of Deposit structs
type DepositSaveForm struct {
	Recid     int64 `json:"recid"`
	LID       int64
	DID       int64
	BID       int64
	BUD       rlib.XJSONBud
	Name      string
	AccountNo string
	LdgrName  string
	GLNumber  string
}

// DepositGridSave is the input data format for a Save command
type DepositGridSave struct {
	Status   string          `json:"status"`
	Recid    int64           `json:"recid"`
	FormName string          `json:"name"`
	Record   DepositSaveForm `json:"record"`
}

// DepositGetResponse is the response to a GetDeposit request
type DepositGetResponse struct {
	Status string      `json:"status"`
	Record DepositGrid `json:"record"`
}

// SvcHandlerDeposit formats a complete data record for an assessment for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the DID as follows:
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerDeposit(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "SvcHandlerDeposit"
		err      error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  DID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerDepositories(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				err = fmt.Errorf("DepositID is required but was not specified")
				SvcGridErrorReturn(w, err, funcname)
				return
			}
			getDeposit(w, r, d)
		}
		break
	case "save":
		saveDeposit(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

// depositGridRowScan scans a result from sql row and dump it in a PrARGrid struct
func depositGridRowScan(rows *sql.Rows, q DepositGrid) (DepositGrid, error) {
	err := rows.Scan(&q.DID, &q.LID, &q.Name, &q.AccountNo, &q.LdgrName, &q.GLNumber, &q.LastModTime, &q.LastModBy, &q.CreateTS, &q.CreateBy)
	return q, err
}

var depositSearchFieldMap = selectQueryFieldMap{
	"DID":         {"Deposit.DID"},
	"LID":         {"Deposit.LID"},
	"Name":        {"Deposit.Name"},
	"AccountNo":   {"Deposit.AccountNo"},
	"LdgrName":    {"GLAccount.Name"},
	"GLNumber":    {"GLAccount.GLNumber"},
	"LastModTime": {"Deposit.LastModTime"},
	"LastModBy":   {"Deposit.LastModBy"},
	"CreateTS":    {"Deposit.CreateTS"},
	"CreateBy":    {"Deposit.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var depositSearchSelectQueryFields = selectQueryFields{
	"Deposit.DID",
	"Deposit.LID",
	"Deposit.Name",
	"Deposit.AccountNo",
	"GLAccount.Name as LdgrName",
	"GLAccount.GLNumber",
	"Deposit.LastModTime",
	"Deposit.LastModBy",
	"Deposit.CreateTS",
	"Deposit.CreateBy",
}

// GetDeposit returns the requested assessment
// wsdoc {
//  @Title  Save Deposit
//	@URL /v1/dep/:BUI/:DID
//  @Method  GET
//	@Synopsis Update the information on a Deposit with the supplied data
//  @Description  This service updates Deposit :DID with the information supplied. All fields must be supplied.
//	@Input DepositGridSave
//  @Response SvcStatusResponse
// wsdoc }
func saveDeposit(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "saveDeposit"
		foo      DepositGridSave
		err      error
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

	var a rlib.Deposit
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, foo.Record.BUD)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	if a.DID == 0 && d.ID == 0 {
		// This is a new AR
		rlib.Console(">>>> NEW DEPOSIT IS BEING ADDED\n")
		_, err = rlib.InsertDeposit(&a)
	} else {
		// update existing record
		rlib.Console("Updating existing Deposit: %d\n", a.DID)
		err = rlib.UpdateDeposit(&a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving depository (DID=%d\n: %s", funcname, a.DID, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponse(w)
}

// GetDeposit returns the requested assessment
// wsdoc {
//  @Title  Get Deposit
//	@URL /v1/dep/:BUI/:DID
//  @Method  GET
//	@Synopsis Get information on a Deposit
//  @Description  Return all fields for assessment :DID
//	@Input WebGridSearchRequest
//  @Response DepositGetResponse
// wsdoc }
func getDeposit(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "getDeposit"
		g        DepositGetResponse
		whr      = fmt.Sprintf("Deposit.DID=%d", d.ID)
	)

	fmt.Printf("entered %s\n", funcname)

	depQuery := `
	SELECT
		{{.SelectClause}}
	FROM Deposit
	LEFT JOIN GLAccount on GLAccount.LID=Deposit.LID
	WHERE {{.WhereClause}};`

	qc := queryClauses{
		"SelectClause": strings.Join(depositSearchSelectQueryFields, ","),
		"WhereClause":  whr,
	}

	qry := renderSQLQuery(depQuery, qc)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var q DepositGrid
		q.BID = d.BID
		q.BUD = getBUDFromBIDList(q.BID)

		q, err = depositGridRowScan(rows, q)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		q.Recid = q.DID
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
