package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// DepositGrid contains the data from Deposit that is targeted to the UI Grid that displays
// a list of Deposit structs
type DepositGrid struct {
	Recid         int64 `json:"recid"`
	DID           int64
	BID           int64
	BUD           rlib.XJSONBud
	DEPID         int64
	DEPName       string
	DPMID         int64
	DPMName       string
	Dt            rlib.JSONDate
	Amount        float64
	ClearedAmount float64
	FLAGS         uint64
	LastModTime   rlib.JSONDateTime
	LastModBy     int64
	CreateTS      rlib.JSONDateTime
	CreateBy      int64
}

// DepositSaveForm contains the data from Deposit FORM that is targeted to the UI Form that displays
// a list of Deposit structs
type DepositSaveForm struct {
	Recid         int64             `json:"recid"` //
	LID           int64             // GL Account for the depository
	DID           int64             // deposit id
	BID           int64             // biz
	BUD           rlib.XJSONBud     // 3-letter code
	DEPID         int64             // Depository ID
	DPMID         int64             // Deposit Method
	Dt            rlib.JSONDateTime // Date of deposit
	Amount        float64           // amount deposited
	ClearedAmount float64           // amount the bank cleared for the deposit
	FLAGS         uint64            // bit 0 = Bank Has provided Cleared Amount
}

// DepositGridSave is the input data format for a Save command
type DepositGridSave struct {
	Cmd      string          `json:"cmd"`
	Recid    int64           `json:"recid"`
	FormName string          `json:"name"`
	Receipts []int64         `json:"Receipts"`
	Record   DepositSaveForm `json:"record"`
}

// DepositDeleteForm is used to delete a depos
type DepositDeleteForm struct {
	Cmd      string `json:"cmd"`
	Recid    int64  `json:"recid"`
	FormName string `json:"formname"`
	DID      int64
}

// DepositGetResponse is the response to a GetDeposit request
type DepositGetResponse struct {
	Status string      `json:"status"`
	Record DepositGrid `json:"record"`
}

// SvcHandlerDeposit dispatches the web request to the appropriate handler
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerDeposit(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerDeposit"
	var (
		err error
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  DID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerDeposits(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				err = fmt.Errorf("DepositID is required but was not specified")
				SvcErrorReturn(w, err, funcname)
				return
			}
			getDeposit(w, r, d)
		}
		break
	case "save":
		saveDeposit(w, r, d)
	case "delete":
		deleteDeposit(w, r, d)
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// depositGridRowScan scans a result from sql row and dump it in a PrARGrid struct
func depositGridRowScan(rows *sql.Rows, a *DepositGrid) error {
	err := rows.Scan(&a.DID, &a.BID, &a.DEPID, &a.DEPName, &a.DPMID, &a.DPMName, &a.Dt, &a.Amount, &a.ClearedAmount, &a.FLAGS, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
	return err
}

// DepositSearchResponse is a response string to the search request for Deposit records
type DepositSearchResponse struct {
	Status  string        `json:"status"`
	Total   int64         `json:"total"`
	Records []DepositGrid `json:"records"`
}

var depositSearchFieldMap = rlib.SelectQueryFieldMap{
	"DID":           {"Deposit.DID"},
	"BID":           {"Deposit.BID"},
	"DEPID":         {"Deposit.DEPID"},
	"DEPName":       {"Depository.Name"},
	"DPMID":         {"Deposit.DPMID"},
	"DPMName":       {"DepositMethod.Method"},
	"Dt":            {"Deposit.Dt"},
	"Amount":        {"Deposit.Amount"},
	"ClearedAmount": {"Deposit.ClearedAmount"},
	"FLAGS":         {"Deposit.FLAGS"},
	"LastModTime":   {"Deposit.LastModTime"},
	"LastModBy":     {"Deposit.LastModBy"},
	"CreateTS":      {"Deposit.CreateTS"},
	"CreateBy":      {"Deposit.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var depositSearchSelectQueryFields = rlib.SelectQueryFields{
	"Deposit.DID",
	"Deposit.BID",
	"Deposit.DEPID",
	"Depository.Name",
	"Deposit.DPMID",
	"DepositMethod.Method",
	"Deposit.Dt",
	"Deposit.Amount",
	"Deposit.ClearedAmount",
	"Deposit.FLAGS",
	"Deposit.LastModTime",
	"Deposit.LastModBy",
	"Deposit.CreateTS",
	"Deposit.CreateBy",
}

// SvcSearchHandlerDeposits generates a report of all Deposits defined business d.BID
// wsdoc {
//  @Title  Search Deposits
//	@URL /v1/deposit/:BUI
//  @Method  POST
//	@Synopsis Search Deposits
//  @Descr  Search all PaymentType and return those that match the Search Logic.
//  @Descr  The search criteria includes start and stop dates of interest.
//	@Input WebGridSearchRequest
//  @Response DepositSearchResponse
// wsdoc }
func SvcSearchHandlerDeposits(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerDeposits"
	var (
		g     DepositSearchResponse
		err   error
		order = "Deposit.DID ASC" // default ORDER
		whr   = fmt.Sprintf("Deposit.BID=%d AND %q <= Deposit.Dt AND Deposit.Dt < %q",
			d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL),
			d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	)

	rlib.Console("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	_, orderClause := GetSearchAndSortSQL(d, depositSearchFieldMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	theQuery := `
	SELECT {{.SelectClause}}
	FROM Deposit
	LEFT JOIN Depository ON Deposit.DEPID = Depository.DEPID
	LEFT JOIN DepositMethod ON Deposit.DPMID = DepositMethod.DPMID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(depositSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(theQuery, qc)
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
	theQueryWithLimit := theQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(theQueryWithLimit, qc)
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
		var q DepositGrid
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		err = depositGridRowScan(rows, &q)
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

// SaveDeposit returns the requested assessment
// wsdoc {
//  @Title  Save Deposit
//	@URL /v1/deposit/:BUI/:DID
//  @Method  POST
//	@Synopsis Update the information on a Deposit with the supplied data
//  @Description  This service updates Deposit :DID with the information supplied. All fields must be supplied.
//	@Input DepositGridSave
//  @Response SvcStatusResponse
// wsdoc }
func saveDeposit(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveDeposit"
	var (
		foo DepositGridSave
		// err error
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

	var a rlib.Deposit
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	var ok bool

	a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	if !ok {
		//-------------------------
		// one more thing to try:
		//-------------------------
		rlib.RRdb.BUDlist, rlib.RRdb.BizCache = rlib.BuildBusinessDesignationMap()
		a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
		if !ok {
			e := fmt.Errorf("%s: Could not map BUD value: %s", funcname, foo.Record.BUD)
			rlib.Ulog("%s", e.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	if a.DID == 0 && d.ID == 0 {
		// This is a new AR
		rlib.Console(">>>> NEW DEPOSIT IS BEING ADDED\n")
		rlib.Console("Receipts[] = %#v\n", foo.Receipts)
		e := bizlogic.SaveDeposit(r.Context(), &a, foo.Receipts)
		if len(e) > 0 {
			SvcErrListReturn(w, e, funcname)
			return
		}
	} else {
		// update existing record
		rlib.Console("Updating existing Deposit: %d\n", a.DID)
		e := bizlogic.SaveDeposit(r.Context(), &a, foo.Receipts)
		if len(e) > 0 {
			SvcErrListReturn(w, e, funcname)
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// DeleteDeposit returns the requested assessment
// wsdoc {
//  @Title  Delete Deposit
//	@URL /v1/deposit/:BUI/DID
//  @Method  GET
//	@Synopsis Delete a deposit
//  @Description  Deletes deposit DID and
//	@Input WebGridSearchRequest
//  @Response DepositGetResponse
// wsdoc }
func deleteDeposit(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteDeposit"
	var (
		del DepositDeleteForm
		err error
	)

	rlib.Console("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//----------------------------------------
	// Remove the deposit parts and mark each
	// Receipt as no longer a member of DID
	//----------------------------------------
	m, err := rlib.GetDepositParts(r.Context(), del.DID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	for i := 0; i < len(m); i++ {
		rcpt, _ := rlib.GetReceipt(r.Context(), m[i].RCPTID)
		if rcpt.RCPTID > 0 {
			rcpt.DID = 0
			if err = rlib.UpdateReceipt(r.Context(), &rcpt); err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
		}
		if err = rlib.DeleteDepositPart(r.Context(), m[i].DPID); err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	}
	err = rlib.DeleteDeposit(r.Context(), del.DID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// GetDeposit returns the requested assessment
// wsdoc {
//  @Title  Get Deposit
//	@URL /v1/deposit/:BUI/:DID
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
		whr      = fmt.Sprintf("Deposit.BID=%d AND Deposit.DID=%d", d.BID, d.ID)
	)

	rlib.Console("entered %s\n", funcname)

	depQuery := `
	SELECT
		{{.SelectClause}}
	FROM Deposit
	LEFT JOIN Depository ON Deposit.DEPID = Depository.DEPID
	LEFT JOIN DepositMethod ON Deposit.DPMID = DepositMethod.DPMID
	WHERE {{.WhereClause}};`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(depositSearchSelectQueryFields, ","),
		"WhereClause":  whr,
	}

	qry := rlib.RenderSQLQuery(depQuery, qc)
	rlib.Console("query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		rlib.Console("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var q DepositGrid
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		err = depositGridRowScan(rows, &q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		q.Recid = q.DID
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
