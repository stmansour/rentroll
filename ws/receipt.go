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

// ReceiptSendForm is a structure specifically for the UI. It will be
// automatically populated from an rlib.Receipt struct
type ReceiptSendForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	RCPTID         int64
	PRCPTID        int64 // Parent RCPTID, points to RCPT being amended/corrected by this receipt
	BID            int64
	DID            int64
	BUD            rlib.XJSONBud
	RAID           int64
	PMTID          int64
	Payor          string // name of the payor
	TCID           int64  // TCID of payor
	Dt             rlib.JSONDate
	DocNo          string // check number, money order number, etc.; documents the payment
	Amount         float64
	ARID           int64
	Comment        string
	OtherPayorName string // if not '', the name of a payor who paid this receipt and who may not be in our system
	LastModTime    rlib.JSONDateTime
	LastModBy      int64
	LastModByUser  string
	CreateTS       rlib.JSONDateTime
	CreateByUser   string
	CreateBy       int64
	FLAGS          uint64
	RentableName   string // FOR RECEIPT-ONLY CLIENT - to be removed when we no longer need that client
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
	BID            int64
	DID            int64
	BUD            rlib.XJSONBud
	RAID           int64
	ARID           int64
	PRCPTID        int64 // Parent RCPTID, points to RCPT being amended/corrected by this receipt
	PMTID          int64
	Dt             rlib.JSONDate
	DocNo          string // check number, money order number, etc.; documents the payment
	Amount         float64
	Payor          string // name of the payor
	TCID           int64  // TCID of payor
	Comment        string
	OtherPayorName string // if not '', the name of a payor who paid this receipt and who may not be in our system
	FLAGS          uint64
	// AcctRule       string
}

// PrReceiptGrid is a structure specifically for the UI Grid.
type PrReceiptGrid struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	RCPTID         int64
	BID            int64
	DID            int64
	TCID           int64 // TCID of payor
	PMTID          int64
	PmtTypeName    string
	Dt             rlib.JSONDate
	DocNo          string // check number, money order number, etc.; documents the payment
	Amount         float64
	Payor          rlib.NullString // name of the payor
	ARID           int64           // which account rule
	AcctRule       rlib.NullString // expression showing how to account for the amount
	FLAGS          uint64
	OtherPayorName string // if not '', the name of a payor who paid this receipt and who may not be in our system
	Comment        string
	RentableName   string // FOR RECEIPT-ONLY CLIENT - to be removed when we no longer need that client
}

// SaveReceiptInput is the input data format for a Save command
type SaveReceiptInput struct {
	Status   string          `json:"status"`
	Recid    int64           `json:"recid"`
	FormName string          `json:"name"`
	Record   ReceiptSaveForm `json:"record"`
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
	err := rows.Scan(&q.RCPTID, &q.BID, &q.TCID, &q.PMTID, &q.PmtTypeName, &q.Dt, &q.DocNo, &q.Amount, &q.Payor, &q.ARID, &q.AcctRule, &q.FLAGS, &q.DID, &q.OtherPayorName, &q.Comment)
	return q, err
}

// which fields needs to be fetched for SQL query for receipts grid
var receiptsFieldsMap = map[string][]string{
	"RCPTID":         {"Receipt.RCPTID"},
	"BID":            {"Receipt.BID"},
	"TCID":           {"Receipt.TCID"},
	"PMTID":          {"Receipt.PMTID"},
	"PmtTypeName":    {"PaymentType.Name"},
	"Dt":             {"Receipt.Dt"},
	"DocNo":          {"Receipt.DocNo"},
	"Amount":         {"Receipt.Amount"},
	"Payor":          {"Transactant.FirstName", "Transactant.LastName", "Transactant.CompanyName"},
	"ARID":           {"Receipt.ARID"},
	"AcctRule":       {"AR.Name"},
	"FLAGS":          {"Receipt.FLAGS"},
	"DID":            {"Receipt.DID"},
	"OtherPayorName": {"Receipt.OtherPayorName"},
	"Comment":        {"Receipt.Comment"},
}

// which fields needs to be fetched for SQL query for receipts grid
var receiptsQuerySelectFields = []string{
	"Receipt.RCPTID",
	"Receipt.BID",
	"Receipt.TCID",
	"Receipt.PMTID",
	"PaymentType.Name as PmtTypeName",
	"Receipt.Dt",
	"Receipt.DocNo",
	"Receipt.Amount",
	"CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END AS Payor",
	"Receipt.ARID",
	"AR.Name as AcctRule",
	"Receipt.FLAGS",
	"Receipt.DID",
	"Receipt.OtherPayorName",
	"Receipt.Comment",
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
	const funcname = "SvcSearchHandlerReceipts"
	var (
		err error
		g   SearchReceiptsResponse
	)
	rlib.Console("Entered %s\n", funcname)

	whr := `Receipt.BID=%d AND Receipt.Dt >= %q and Receipt.Dt < %q`
	whr = fmt.Sprintf(whr, d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	order := "Receipt.Dt ASC, Receipt.RCPTID ASC" // default ORDER

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
	LEFT JOIN PaymentType ON Receipt.PMTID=PaymentType.PMTID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(receiptsQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(receiptsQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		rlib.Console("Error from rlib.GetQueryCount: %s\n", err.Error())
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
	receiptsQueryWithLimit := receiptsQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(receiptsQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		rlib.Console("Error from DB Query: %s\n", err.Error())
		SvcErrorReturn(w, err, funcname)
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
			SvcErrorReturn(w, err, funcname)
			return
		}

		//---------------------------------------------------------------
		// RECEIPT-ONLY CLIENT UPDATE...
		// extract the RentableName from the comment if it is present...
		//---------------------------------------------------------------
		if d.wsSearchReq.Client == rlib.RECEIPTONLYCLIENT {
			q.RentableName, q.Comment = rlib.ROCExtractRentableName(q.Comment)
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
	const funcname = "SvcFormHandlerReceipt"
	var (
		err error
	)
	rlib.Console("Entered %s\n", funcname)

	if d.RCPTID, err = SvcExtractIDFromURI(r.RequestURI, "RCPTID", 3, w); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Request: %s:  BID = %d,  RCPTID = %d\n", d.wsSearchReq.Cmd, d.BID, d.RCPTID)

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
		SvcErrorReturn(w, err, funcname)
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
	const funcname = "saveReceipt"
	var (
		err error
		foo SaveReceiptInput
		a   rlib.Receipt
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	//-------------------------------------------------
	//  First, parse out the main form data into a...
	//-------------------------------------------------
	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	// rlib.Console("saveReceipt - first migrate: a = %#v\n", a)

	//------------------------------------------
	//  Update or Insert as appropriate...
	//------------------------------------------
	if a.RCPTID == 0 && d.RCPTID == 0 {
		//-------------------------------------------------------------------
		// there is one special case: if the client is the receipt-only
		//-------------------------------------------------------------------
		if d.wsSearchReq.Client == rlib.RECEIPTONLYCLIENT {
			//----------------------------------------------------------------
			// There is no field for the RentableName in a receipt. But we
			// need one for the Receipt-Only client. We will encode it onto
			// the comment field with double-braces.  We will remove the
			// double braces and RentableName when the Read-Only client reads
			// back these receipts.  They will be transferred to the client
			// in a RentableName field.
			//----------------------------------------------------------------
			if len(d.wsSearchReq.RentableName) > 0 {
				a.Comment += rlib.ROCPRE + d.wsSearchReq.RentableName + rlib.ROCPOST
			}
			_, err = rlib.InsertReceipt(r.Context(), &a)
		} else {
			err = bizlogic.InsertReceipt(r.Context(), &a)
		}
		if err != nil {
			e := fmt.Errorf("%s:  Error in rlib.ProcessNewReceipt: %s", funcname, err.Error())
			rlib.Ulog("%s", e.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	} else {
		//-------------------------------------------------------------------
		// there is one special case: if the client is the receipt-only
		// client for Isola Bella, we skip the business checks because only
		// the receipts are being saved, nothing else.  This will go away
		// in the future when we're able to keep the payors up-to-date for
		// Isola Bella.
		//-------------------------------------------------------------------
		rlib.Console("%s: d.RCPTID = %d, a.RCPTID = %d\n", d.RCPTID)
		if d.wsSearchReq.Client == rlib.RECEIPTONLYCLIENT {
			rcpt, err := rlib.GetReceipt(r.Context(), d.RCPTID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			if receiptOnlyUIUpdateAndReverse(&a, &rcpt, w, r, d) {
				return
			}
			SvcWriteSuccessResponseWithID(d.BID, w, a.RCPTID)
			return
		}
		now := time.Now() // this is the time we're making the change if a reversal needs to be done
		err = bizlogic.UpdateReceipt(r.Context(), &a, &now)
		if err != nil {
			e := fmt.Errorf("%s: Error saving receipt (RCPTID=%d)\n: %s", funcname, d.RCPTID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.RCPTID)
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
	const funcname = "getReceipt"
	rlib.Console("entered %s\n", funcname)
	var g GetReceiptResponse
	a, _ := rlib.GetReceiptNoAllocations(r.Context(), d.RCPTID)
	if a.RCPTID > 0 {
		var gg ReceiptSendForm
		gg.BID = d.BID
		gg.BUD = rlib.GetBUDFromBIDList(d.BID)

		// migrate receipt values in resp struct
		rlib.MigrateStructVals(&a, &gg)

		if a.TCID > 0 {
			var t rlib.Transactant
			_ = rlib.GetTransactant(r.Context(), a.TCID, &t)
			if t.TCID > 0 {
				tcid := strconv.FormatInt(t.TCID, 10)
				gg.Payor = t.GetFullTransactantName() + " (TCID: " + tcid + ")"
			}
		}

		// RECEIPT-ONLY CLIENT - Remove when this client is no longer needed
		if d.wsSearchReq.Client == rlib.RECEIPTONLYCLIENT {
			gg.RentableName, gg.Comment = rlib.ROCExtractRentableName(gg.Comment)
		}

		gg.CreateByUser = rlib.GetNameForUID(r.Context(), a.CreateBy)
		gg.LastModByUser = rlib.GetNameForUID(r.Context(), a.LastModBy)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// deleteReceipt reverses the requested receipt and other linked records
// wsdoc {
//  @Title  Reverse Receipt
//	@URL /v1/receipt/:BUI/:RCPTID
//  @Method  POST
//	@Synopsis Reverse a Receipt
//  @Description  *** WARNING ***  Only use this command if you really know what you're doing.
//  @Description  Delete Receipt records for requested RCPTID. It also deletes associated
//  @Description  Journal, JournalAllocation, and ReceiptAllocation records.
//	@Input DeleteRcptForm
//  @Response SvcWriteSuccessResponse
// wsdoc }
func deleteReceipt(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteReceipt"
	var (
		del DeleteRcptForm
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rcpt, err := rlib.GetReceipt(r.Context(), del.RCPTID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//---------------------------------------------------------------
	// RECEIPT-ONLY CLIENT UPDATE...
	// This reversal is much simpler than the one in biz logic
	// Here, we simply set the flags and make a new receipt to
	// negate the one being reversed.
	//---------------------------------------------------------------
	if d.wsSearchReq.Client == rlib.RECEIPTONLYCLIENT {
		if receiptOnlyUIReverse(&rcpt, w, r, d) {
			return
		}
	} else {
		//-------------------------------------------------------
		// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
		//-------------------------------------------------------
		tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		// reverse receipt in atomic transaction mode
		dt := time.Now()
		err = bizlogic.ReverseReceipt(ctx, &rcpt, &dt)
		if err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}

		// ------------------
		// COMMIT TRANSACTION
		// ------------------
		if err := tx.Commit(); err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// receiptOnlyUIReverse
//
// INPUTS
//  rcpt - receipt to be reversed
//  w      - http handle
//  r      - http request
//  d      - service data
//
// RETURNS
//  bool   - false = no errors occurred
//         - true = error occurred and was sent to client
//---------------------------------------------------------------------
func receiptOnlyUIReverse(rcpt *rlib.Receipt, w http.ResponseWriter, r *http.Request, d *ServiceData) bool {
	const funcname = "receiptOnlyUIReverse"
	var (
		err error
	)

	if rcpt.FLAGS&0x04 != 0 {
		err := fmt.Errorf("Error: receipt %s has already been reversed", rlib.IDtoShortString("RCPT", rcpt.RCPTID))
		SvcErrorReturn(w, err, funcname)
		return true
	}
	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return true
	}

	//------------------------------------------------------
	// Build the new receipt
	//------------------------------------------------------
	rname, _ := rlib.ROCExtractRentableName(rcpt.Comment)
	rr := *rcpt
	rr.RCPTID = int64(0)
	rr.Amount = -rr.Amount
	rr.Comment = fmt.Sprintf("Reversal of receipt %s", rcpt.IDtoShortString()) + rlib.ROCPRE + rname + rlib.ROCPOST
	rr.PRCPTID = rcpt.RCPTID  // link to parent
	rr.FLAGS |= rlib.RCPTvoid // mark that it is voided
	_, err = rlib.InsertReceipt(ctx, &rr)
	if err != nil {
		tx.Rollback() // TBD: abort transaction
		SvcErrorReturn(w, err, funcname)
		return true
	}
	//------------------------------------------------------
	// update the flags on the original receipt
	//------------------------------------------------------
	rcpt.FLAGS |= rlib.RCPTvoid
	if err = rlib.UpdateReceipt(ctx, rcpt); err != nil {
		SvcErrorReturn(w, err, funcname)
		tx.Rollback() // TBD: abort transaction
		return true
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err := tx.Commit(); err != nil {
		tx.Rollback() // TBD: abort transaction
		SvcErrorReturn(w, err, funcname)
		return true
	}
	return false
}

// receiptOnlyUIUpdateAndReverse
//
// INPUTS
//  uprcpt - the updated receipt structure from the UI
//  rcpt   - the original version of the receipt prior to updating
//  w      - http handle
//  r      - http request
//  d      - service data
//---------------------------------------------------------------------
func receiptOnlyUIUpdateAndReverse(uprcpt, rcpt *rlib.Receipt, w http.ResponseWriter, r *http.Request, d *ServiceData) bool {
	funcname := "receiptOnlyUIUpdateAndReverse"
	var err error
	rcptAddr, _ := rlib.ROCExtractRentableName(rcpt.Comment)
	uprcptAddr := d.wsSearchReq.RentableName

	//------------------------------------------
	// check to see if a reversal is needed...
	//------------------------------------------
	bReverse := uprcpt.Amount != rcpt.Amount
	bReverse = bReverse || uprcpt.Dt.Year() != rcpt.Dt.Year() || uprcpt.Dt.Month() != rcpt.Dt.Month() || uprcpt.Dt.Day() != rcpt.Dt.Day()
	bReverse = bReverse || uprcptAddr != rcptAddr
	bReverse = bReverse || uprcpt.OtherPayorName != rcpt.OtherPayorName
	if bReverse {
		if receiptOnlyUIReverse(rcpt, w, r, d) {
			return true
		}
	}

	//-------------------------------------------------------------------
	// Regardless of what happened with the reversal, we need to
	// properly encode the comment field of the new or updated receipt.
	//-------------------------------------------------------------------
	if len(d.wsSearchReq.RentableName) > 0 {
		uprcpt.Comment += rlib.ROCPRE + d.wsSearchReq.RentableName + rlib.ROCPOST
	}

	//--------------------------------------------------------
	// If we did not reverse uprcpt, the simply update it.
	// If we reversed it, we'll need to create a new receipt
	//--------------------------------------------------------
	if bReverse {
		_, err = rlib.InsertReceipt(r.Context(), uprcpt)
	} else {
		err = rlib.UpdateReceipt(r.Context(), uprcpt)
	}
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return true
	}
	return false
}
