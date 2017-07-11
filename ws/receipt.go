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

// ReceiptSendForm is a structure specifically for the UI. It will be
// automatically populated from an rlib.Receipt struct
type ReceiptSendForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	RCPTID         int64
	PRCPTID        int64 // Parent RCPTID, points to RCPT being amended/corrected by this receipt
	BID            rlib.XJSONBud
	PMTID          int64
	Payor          string // name of the payor
	TCID           int64  // TCID of payor
	Dt             rlib.JSONTime
	DocNo          string // check number, money order number, etc.; documents the payment
	Amount         float64
	ARID           int64
	Comment        string
	OtherPayorName string // if not '', the name of a payor who paid this receipt and who may not be in our system
	LastModTime    rlib.JSONTime
	LastModBy      int64
	//AcctRule       string
}

// ReceiptSaveForm is a structure specifically for the return value from w2ui.
// Data does not always come back in the same format it was sent. For example,
// values from dropdown lists come back in the form of a rlib.W2uiHTMLSelect struct.
// So, we break up the ingest into 2 parts. First, we read back the fields that look
// just like the xxxSendForm -- this is what is in xxxSaveForm. Then we readback
// the data that has changed, which is in the xxxSaveOther struct.  All this data
// is merged into the appropriate database structure using MigrateStructData.
type ReceiptSaveForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	RCPTID         int64
	PRCPTID        int64 // Parent RCPTID, points to RCPT being amended/corrected by this receipt
	PMTID          int64
	Dt             rlib.JSONTime
	DocNo          string // check number, money order number, etc.; documents the payment
	Amount         float64
	Payor          string // name of the payor
	TCID           int64  // TCID of payor
	Comment        string
	OtherPayorName string // if not '', the name of a payor who paid this receipt and who may not be in our system
	LastModTime    rlib.JSONTime
	LastModBy      int64
	// AcctRule       string
}

// ReceiptSaveOther is a struct to handle the UI list box selections
type ReceiptSaveOther struct {
	BID  rlib.W2uiHTMLSelect
	ARID rlib.W2uiHTMLSelect
}

// PrReceiptGrid is a structure specifically for the UI Grid.
type PrReceiptGrid struct {
	Recid    int64 `json:"recid"` // this is to support the w2ui form
	RCPTID   int64
	BID      int64
	TCID     int64 // TCID of payor
	PMTID    int64
	Dt       rlib.JSONTime
	DocNo    string // check number, money order number, etc.; documents the payment
	Amount   float64
	Payor    rlib.NullString // name of the payor
	ARID     int64           // which account rule
	AcctRule rlib.NullString // expression showing how to account for the amount
}

// SaveReceiptInput is the input data format for a Save command
type SaveReceiptInput struct {
	Status   string          `json:"status"`
	Recid    int64           `json:"recid"`
	FormName string          `json:"name"`
	Record   ReceiptSaveForm `json:"record"`
}

// SaveReceiptOther is the input data format for the "other" data on the Save command
type SaveReceiptOther struct {
	Status string           `json:"status"`
	Recid  int64            `json:"recid"`
	Name   string           `json:"name"`
	Record ReceiptSaveOther `json:"record"`
}

// SearchReceiptsResponse is a response string to the search request for receipts
type SearchReceiptsResponse struct {
	Status  string          `json:"status"`
	Total   int64           `json:"total"`
	Records []PrReceiptGrid `json:"records"`
}

// GetReceiptResponse is the response to a GetReceipt request
type GetReceiptResponse struct {
	Status string          `json:"status"`
	Record ReceiptSendForm `json:"record"`
}

// DeleteRcptForm holds RCPTID to delete it
type DeleteRcptForm struct {
	RCPTID int64
}

// receiptsGridRowScan scans a result from sql row and dump it in a PrReceiptGrid struct
func receiptsGridRowScan(rows *sql.Rows, q PrReceiptGrid) (PrReceiptGrid, error) {
	err := rows.Scan(&q.RCPTID, &q.BID, &q.TCID, &q.PMTID, &q.Dt, &q.DocNo, &q.Amount, &q.Payor, &q.ARID, &q.AcctRule)
	return q, err
}

// which fields needs to be fetched for SQL query for receipts grid
var receiptsFieldsMap = map[string][]string{
	"RCPTID":   {"Receipt.RCPTID"},
	"BID":      {"Receipt.BID"},
	"TCID":     {"Receipt.TCID"},
	"PMTID":    {"Receipt.PMTID"},
	"Dt":       {"Receipt.Dt"},
	"DocNo":    {"Receipt.DocNo"},
	"Amount":   {"Receipt.Amount"},
	"Payor":    {"Transactant.FirstName", "Transactant.LastName", "Transactant.CompanyName"},
	"ARID":     {"Receipt.ARID"},
	"AcctRule": {"AR.Name"},
}

// which fields needs to be fetched for SQL query for receipts grid
var receiptsQuerySelectFields = []string{
	"Receipt.RCPTID",
	"Receipt.BID",
	"Receipt.TCID",
	"Receipt.PMTID",
	"Receipt.Dt",
	"Receipt.DocNo",
	"Receipt.Amount",
	"CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END AS Payor",
	"Receipt.ARID",
	"AR.Name",
}

// SvcSearchHandlerReceipts generates a report of all Receipts defined business d.BID
// wsdoc {
//  @Title  Search Receipts
//	@URL /v1/receipts/:BUI
//  @Method  POST
//	@Synopsis Search Receipts
//  @Description  Search all Receipts and return those that match the Search Logic.
//  @Desc By default, the search is made for receipts from "today" to 31 days prior.
//	@Input WebGridSearchRequest
//  @Response SearchReceiptsResponse
// wsdoc }
func SvcSearchHandlerReceipts(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcSearchHandlerReceipts"
		err      error
		g        SearchReceiptsResponse
	)
	fmt.Printf("Entered %s\n", funcname)

	whr := `Receipt.BID=%d AND Receipt.Dt >= %q and Receipt.Dt < %q`
	whr = fmt.Sprintf(whr, d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	order := "Receipt.Dt ASC" // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, receiptsFieldsMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	receiptsQuery := `
	SELECT
		{{.SelectClause}}
	FROM Receipt
	LEFT JOIN Transactant ON Receipt.TCID=Transactant.TCID
	LEFT JOIN AR ON Receipt.ARID=AR.ARID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := queryClauses{
		"SelectClause": strings.Join(receiptsQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := renderSQLQuery(receiptsQuery, qc)
	g.Total, err = GetQueryCount(countQuery, qc)
	if err != nil {
		fmt.Printf("Error from GetQueryCount: %s\n", err.Error())
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
	receiptsQueryWithLimit := receiptsQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := renderSQLQuery(receiptsQueryWithLimit, qc)
	fmt.Printf("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("Error from DB Query: %s\n", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q PrReceiptGrid
		q.Recid = i

		q, err = receiptsGridRowScan(rows, q)
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

// SvcFormHandlerReceipt formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the RCPTID as follows:
//           0    1     2   3
// uri 		/v1/receipt/BUI/RCPTID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerReceipt(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcFormHandlerReceipt"
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)

	if d.RCPTID, err = SvcExtractIDFromURI(r.RequestURI, "RCPTID", 3, w); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	fmt.Printf("Request: %s:  BID = %d,  RCPTID = %d\n", d.wsSearchReq.Cmd, d.BID, d.RCPTID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getReceipt(w, r, d)
		break
	case "save":
		saveReceipt(w, r, d)
		break
	case "delete":
		deleteReceipt(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

// saveReceipt returns the requested receipt
// wsdoc {
//  @Title  Save Receipt
//	@URL /v1/receipt/:BUI/:RCPTID
//  @Method  GET
//	@Synopsis Save a Receipt
//  @Desc  This service saves a Receipt.  If :RCPTID exists, it will
//  @Desc  be updated with the information supplied. All fields must
//  @Desc  be supplied. If RCPTID is 0, then a new receipt is created.
//	@Input SaveReceiptInput
//  @Response SvcStatusResponse
// wsdoc }
func saveReceipt(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "saveReceipt"
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	//-------------------------------------------------
	//  First, parse out the main form data into a...
	//-------------------------------------------------
	var foo SaveReceiptInput
	var a rlib.Receipt
	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	fmt.Printf("saveReceipt - first migrate: a = %#v\n", a)

	//----------------------------------------------
	//  Next, parse out the other form data and
	//  update the fields in a...
	//----------------------------------------------
	var bar SaveReceiptOther
	if err := json.Unmarshal(data, &bar); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	fmt.Printf("bar.Record = %#v\n", bar.Record)
	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.Record.BID.ID]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, bar.Record.BID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	a.ARID, err = rlib.IntFromString(bar.Record.ARID.ID, "Invalid ARID")
	if err != nil {
		e := fmt.Errorf("%s: Bad ARID value: %s", funcname, bar.Record.ARID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	fmt.Printf("saveReceipt - second migrate: a = %#v\n", a)

	//------------------------------------------
	//  Update or Insert as appropriate...
	//------------------------------------------
	if a.RCPTID == 0 && d.RCPTID == 0 {
		// This is a new Receipt
		fmt.Printf(">>>> NEW RECEIPT IS BEING ADDED\n")
		_, err = rlib.InsertReceipt(&a)

		var xbiz rlib.XBusiness
		rlib.InitBizInternals(a.BID, &xbiz)

		// get the AR for this receipt...
		ar := rlib.RRdb.BizTypes[a.BID].AR[a.ARID]

		// get GL Account Info for
		ard := rlib.RRdb.BizTypes[a.BID].GLAccounts[ar.DebitLID]
		arc := rlib.RRdb.BizTypes[a.BID].GLAccounts[ar.CreditLID]

		// create the receipt allocation
		var ra rlib.ReceiptAllocation
		ra.RCPTID = a.RCPTID
		ra.Amount = a.Amount
		ra.AcctRule = fmt.Sprintf("d %s _, c %s _", ard.GLNumber, arc.GLNumber)
		ra.BID = a.BID
		ra.Dt = a.Dt
		rlib.InsertReceiptAllocation(&ra)

		a.RA = append(a.RA, ra)

		//------------------------------------------------
		// Add it to the Journal
		//------------------------------------------------
		d1 := time.Date(a.Dt.Year(), a.Dt.Month(), 1, 0, 0, 0, 0, rlib.RRdb.Zone)
		mon, year := rlib.IncMonths(a.Dt.Month(), int64(a.Dt.Year()))
		d2 := time.Date(int(year), mon, 1, 0, 0, 0, 0, rlib.RRdb.Zone)
		jnl, err := rlib.ProcessNewReceipt(&xbiz, &d1, &d2, &a)
		if err != nil {
			e := fmt.Errorf("%s:  Error in rlib.ProcessNewReceipt: %s", funcname, err.Error())
			rlib.Ulog("%s", e.Error())
			SvcGridErrorReturn(w, e, funcname)
			return
		}

		rlib.GetJournalAllocations(&jnl)
		fmt.Printf("Journal record created: JID = %d, len(jnl.JA) = %d\n", jnl.JID, len(jnl.JA))
		rlib.InitLedgerCache()
		rlib.GenerateLedgerEntriesFromJournal(&xbiz, &jnl, &d1, &d2)

	} else {
		// update existing record
		err = rlib.UpdateReceipt(&a)
	}
	if err != nil {
		e := fmt.Errorf("%s: Error saving receipt (RCPTID=%d\n: %s", funcname, d.RCPTID, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(w, a.RCPTID)
}

// GetReceipt returns the requested receipt
// wsdoc {
//  @Title  Get Receipt
//	@URL /v1/receipt/:BUI/:RCPTID
//  @Method  GET
//	@Synopsis Get information on a Receipt
//  @Description  Return all fields for receipt :RCPTID
//	@Input WebGridSearchRequest
//  @Response GetReceiptResponse
// wsdoc }
func getReceipt(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered getReceipt\n")
	var g GetReceiptResponse
	a := rlib.GetReceiptNoAllocations(d.RCPTID)
	if a.RCPTID > 0 {
		var gg ReceiptSendForm
		rlib.MigrateStructVals(&a, &gg)
		if a.TCID > 0 {
			var t rlib.Transactant
			rlib.GetTransactant(a.TCID, &t)
			if t.TCID > 0 {
				tcid := strconv.FormatInt(t.TCID, 10)
				// feed Payor pattern, it may change depend on the pattern
				// front-end form payor field and this pattern need to be same
				// pattern: "{{name}} (TCID: {{tcid}})"
				if t.IsCompany > 0 {
					gg.Payor = t.CompanyName + " (TCID: " + tcid + ")"
				} else {
					gg.Payor = t.FirstName + " " + t.LastName + " (TCID: " + tcid + ")"
				}
			}

		}
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}

// deleteReceipt deletes the requested receipt and other linked records
// wsdoc {
//  @Title  Delete Receipt
//	@URL /v1/receipt/:BUI/:RCPTID
//  @Method  POST
//	@Synopsis Delete a Receipt
//  @Description  *** WARNING ***  Only use this command if you really know what you're doing.
//  @Description  Delete Receipt records for requested RCPTID. It also deletes associated
//  @Description  Journal, JournalAllocation, and ReceiptAllocation records.
//	@Input DeleteRcptForm
//  @Response SvcWriteSuccessResponse
// wsdoc }
func deleteReceipt(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "deleteReceipt"
		del      DeleteRcptForm
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	j := rlib.GetJournalByReceiptID(del.RCPTID)
	rlib.GetJournalAllocations(&j)
	for k := 0; k < len(j.JA); k++ {
		m := rlib.GetLedgerEntriesByJAID(d.BID, j.JA[k].JAID)
		for i := 0; i < len(m); i++ {
			rlib.DeleteLedgerEntry(m[i].LEID)
		}
	}
	rlib.DeleteJournalAllocations(j.JID)
	rlib.DeleteJournal(j.JID)
	if err := rlib.DeleteReceiptAllocations(del.RCPTID); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	if err := rlib.DeleteReceipt(del.RCPTID); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(w)
}
