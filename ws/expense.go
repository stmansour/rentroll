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
	"time"
)

// ExpenseGridNull is like ExpenseGrid but allows for a NULL on Rentable
type ExpenseGridNull struct {
	Recid        int64 `json:"recid"`
	EXPID        int64
	BID          int64
	BUD          rlib.XJSONBud
	RID          int64
	RAID         int64
	Amount       float64
	Dt           rlib.JSONDate
	ARID         int64
	ARName       string
	RentableName rlib.NullString
	FLAGS        uint64
	Comment      string
	LastModTime  rlib.JSONDateTime
	LastModBy    int64
	CreateTS     rlib.JSONDateTime
	CreateBy     int64
}

// ExpenseGrid contains the data from Expense that is targeted to the UI Grid that displays
// a list of Expense structs
type ExpenseGrid struct {
	Recid         int64 `json:"recid"`
	EXPID         int64
	BID           int64
	BUD           rlib.XJSONBud
	RID           int64
	RAID          int64
	Amount        float64
	Dt            rlib.JSONDate
	ARID          int64
	ARName        string
	RName         string
	FLAGS         uint64
	Comment       string
	LastModTime   rlib.JSONDateTime
	LastModBy     int64
	LastModByUser string
	CreateTS      rlib.JSONDateTime
	CreateBy      int64
	CreateByUser  string
}

// ExpenseSearchResponse is a response string to the search request for Expense records
type ExpenseSearchResponse struct {
	Status  string        `json:"status"`
	Total   int64         `json:"total"`
	Records []ExpenseGrid `json:"records"`
}

// ExpenseSaveForm is a struct to handle direct inputs from the form
type ExpenseSaveForm struct {
	Recid   int64 `json:"recid"`
	EXPID   int64
	BID     int64
	BUD     rlib.XJSONBud
	RID     int64
	RAID    int64
	Dt      rlib.JSONDate
	Amount  float64
	ARID    int64
	Comment string
	FLAGS   uint64
}

// SaveExpenseInput is the input data format for a Save command
type SaveExpenseInput struct {
	Recid    int64           `json:"recid"`
	Status   string          `json:"status"`
	FormName string          `json:"name"`
	Record   ExpenseSaveForm `json:"record"`
}

// ExpenseGetResponse is the response to a GetExpense request
type ExpenseGetResponse struct {
	Status string      `json:"status"`
	Record ExpenseGrid `json:"record"`
}

// DeleteExpenseForm used to delete record from database
type DeleteExpenseForm struct {
	ID int64
}

var expenseMethodSearchFieldMap = rlib.SelectQueryFieldMap{
	"EXPID":        {"Expense.EXPID"},
	"BID":          {"Expense.BID"},
	"RID":          {"Expense.RID"},
	"RAID":         {"Expense.RAID"},
	"ARID":         {"Expense.ARID"},
	"ARName":       {"AR.Name"},
	"Amount":       {"Expense.Amount"},
	"Dt":           {"Expense.Dt"},
	"RentableName": {"Rentable.RentableName"},
	"FLAGS":        {"Expense.FLAGS"},
	"Comment":      {"Expense.Comment"},
	"LastModTime":  {"Expense.LastModTime"},
	"LastModBy":    {"Expense.LastModBy"},
	"CreateTS":     {"Expense.CreateTS"},
	"CreateBy":     {"Expense.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var expenseMethodSearchSelectQueryFields = rlib.SelectQueryFields{
	"Expense.EXPID",
	"Expense.BID",
	"Expense.RID",
	"Expense.RAID",
	"Expense.ARID",
	"AR.Name",
	"Expense.Amount",
	"Expense.Dt",
	"Rentable.RentableName",
	"Expense.FLAGS",
	"Expense.Comment",
	"Expense.LastModTime",
	"Expense.LastModBy",
	"Expense.CreateTS",
	"Expense.CreateBy",
}

// pmtRowScan scans a result from sql row and dump it in a ExpenseGrid struct
func expenseRowScan(rows *sql.Rows) (ExpenseGrid, error) {
	var a ExpenseGridNull
	err := rows.Scan(&a.EXPID, &a.BID, &a.RID, &a.RAID, &a.ARID, &a.ARName, &a.Amount, &a.Dt, &a.RentableName, &a.FLAGS, &a.Comment, &a.LastModTime, &a.LastModBy, &a.CreateTS, &a.CreateBy)
	var b ExpenseGrid
	rlib.MigrateStructVals(&a, &b)
	if a.RentableName.Valid {
		b.RName = a.RentableName.String
	}
	return b, err
}

// SvcHandlerExpense dispatches the web request to the appropriate handler:
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerExpense(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerExpense"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  EXPID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerExpenses(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				err = fmt.Errorf("ExpenseID is required but was not specified")
				SvcErrorReturn(w, err, funcname)
				return
			}
			getExpense(w, r, d)
		}
	case "save":
		saveExpense(w, r, d)
	case "delete":
		deleteExpense(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// SvcSearchHandlerExpenses generates a report of all Expenses defined business d.BID
// wsdoc {
//  @Title  Search Expense
//	@URL /v1/expenses/:BUI
//  @Method  POST
//	@Synopsis Search Expenses
//  @Descr  Search all Expense and return those that match the Search Logic.
//  @Descr  The search criteria includes start and stop dates of interest.
//	@Input WebGridSearchRequest
//  @Response ExpenseSearchResponse
// wsdoc }
func SvcSearchHandlerExpenses(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerExpenses"
	var (
		g     ExpenseSearchResponse
		err   error
		order = "EXPID ASC" // default ORDER
		whr   = fmt.Sprintf("Expense.BID=%d AND %q <= Dt AND Dt < %q", d.BID,
			d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL),
			d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	)

	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	_, orderClause := GetSearchAndSortSQL(d, expenseMethodSearchFieldMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	theQuery := `
	SELECT
		{{.SelectClause}}
	FROM Expense
	LEFT JOIN AR ON Expense.ARID = AR.ARID
	LEFT JOIN Rentable on Expense.RID = Rentable.RID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(expenseMethodSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(theQuery, qc)
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
	theQueryWithLimit := theQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(theQueryWithLimit, qc)
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

		q, err := expenseRowScan(rows)
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

// deleteExpense deletes a payment type from the database
// wsdoc {
//  @Title  Delete Expense
//	@URL /v1/expense/:BUI/:RAID
//  @Method  POST
//	@Synopsis Reverses an Expense
//  @Desc  This service reverses a Expense.
//	@Input DeleteExpenseForm
//  @Response SvcStatusResponse
// wsdoc }
func deleteExpense(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteExpense"
	var (
		del DeleteExpenseForm
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	a, err := rlib.GetExpense(r.Context(), del.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	now := time.Now() // mark Assessment reversed at this time
	errlist := bizlogic.ReverseExpense(ctx, &a, &now)
	if len(errlist) > 0 {
		tx.Rollback()
		SvcErrListReturn(w, errlist, funcname)
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// GetExpense returns the requested assessment
// wsdoc {
//  @Title  Save Expense
//	@URL /v1/expense/:BUI/:EXPID
//  @Method  GET
//	@Synopsis Update the information on a Expense with the supplied data
//  @Description  This service updates Expense :EXPID with the information supplied. All fields must be supplied.
//	@Input SaveExpenseInput
//  @Response SvcStatusResponse
// wsdoc }
func saveExpense(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveExpense"
	var (
		foo SaveExpenseInput
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

	var a rlib.Expense
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[string(foo.Record.BUD)]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, foo.Record.BUD)
		rlib.Ulog("%s", e.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	if a.EXPID == 0 && d.ID == 0 {
		_, err = rlib.InsertExpense(r.Context(), &a)
		var xbiz rlib.XBusiness
		err = rlib.ProcessNewExpense(r.Context(), &a, &xbiz)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	} else {
		fmt.Printf("Updating existing Expense: %d\n", a.EXPID)
		now := time.Now() // in case reversal is necessary
		errlist := bizlogic.UpdateExpense(r.Context(), &a, &now)
		if len(errlist) > 0 {
			SvcErrListReturn(w, errlist, funcname)
			return
		}
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving Expense %d: %s", funcname, a.EXPID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// GetExpense returns the requested assessment
// wsdoc {
//  @Title  Get Payment Type
//	@URL /v1/expense/:BUI/:EXPID
//  @Method  GET
//	@Synopsis Get information on a Expense
//  @Description  Return all fields for assessment :EXPID
//	@Input WebGridSearchRequest
//  @Response ExpenseGetResponse
// wsdoc }
func getExpense(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getExpense"
	var (
		g   ExpenseGetResponse
		a   rlib.Expense
		err error
		gg  ExpenseGrid
	)

	rlib.Console("entered %s.  Expense ID = %d\n", funcname, d.ID)
	a, err = rlib.GetExpense(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.RID > 0 {
		rentable, _ := rlib.GetRentable(r.Context(), a.RID)
		gg.RName = rentable.RentableName
	}

	if a.EXPID > 0 {
		rlib.MigrateStructVals(&a, &gg)
		gg.BUD = rlib.GetBUDFromBIDList(gg.BID)
		gg.CreateByUser = rlib.GetNameForUID(r.Context(), a.CreateBy)
		gg.LastModByUser = rlib.GetNameForUID(r.Context(), a.LastModBy)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
