package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// ReceiptSendForm is a structure specifically for the UI. It will be
// automatically populated from an rlib.Receipt struct
type ReceiptSendForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	RCPTID         int64
	PRCPTID        int64 // Parent RCPTID, points to RCPT being amended/corrected by this receipt
	BID            rlib.XJSONBud
	PMTID          int64
	Dt             rlib.JSONTime
	DocNo          string // check number, money order number, etc.; documents the payment
	Amount         float64
	AcctRule       string
	Comment        string
	OtherPayorName string // if not '', the name of a payor who paid this receipt and who may not be in our system
	LastModTime    rlib.JSONTime
	LastModBy      int64
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
	AcctRule       string
	Comment        string
	OtherPayorName string // if not '', the name of a payor who paid this receipt and who may not be in our system
	LastModTime    rlib.JSONTime
	LastModBy      int64
}

// ReceiptSaveOther is a struct to handle the UI list box selections
type ReceiptSaveOther struct {
	BID rlib.W2uiHTMLSelect
}

// PrReceiptGrid is a structure specifically for the UI Grid.
type PrReceiptGrid struct {
	Recid  int64 `json:"recid"` // this is to support the w2ui form
	RCPTID int64
	BID    rlib.XJSONBud
	//RAID   int64
	PMTID  int64
	Dt     rlib.JSONTime
	DocNo  string // check number, money order number, etc.; documents the payment
	Amount float64
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
	funcname := "SvcSearchHandlerReceipts"
	fmt.Printf("Entered %s\n", funcname)
	var (
		err error
		g   SearchReceiptsResponse
	)
	order := "Dt ASC"                                                          // default ORDER
	q := fmt.Sprintf("SELECT %s FROM Receipt ", rlib.RRdb.DBFields["Receipt"]) // the fields we want
	qw := fmt.Sprintf("BID=%d AND Dt >= %q and Dt < %q", d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	q += "WHERE " + qw + " ORDER BY "
	if len(d.wsSearchReq.Sort) > 0 {
		for i := 0; i < len(d.wsSearchReq.Sort); i++ {
			if i > 0 {
				q += ","
			}
			q += d.wsSearchReq.Sort[i].Field + " " + d.wsSearchReq.Sort[i].Direction
		}
	} else {
		q += order
	}

	// now set up the offset and limit
	q += fmt.Sprintf(" LIMIT %d OFFSET %d", d.wsSearchReq.Limit, d.wsSearchReq.Offset)
	fmt.Printf("db query = %s\n", q)

	g.Total, err = GetRowCount("Receipt", qw)
	if err != nil {
		fmt.Printf("Error from GetRowCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}
	rows, err := rlib.RRdb.Dbrr.Query(q)
	if err != nil {
		fmt.Printf("Error from DB Query: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var p rlib.Receipt
		var q PrReceiptGrid
		rlib.ReadReceipts(rows, &p)
		rlib.MigrateStructVals(&p, &q)
		q.Recid = p.RCPTID
		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}
	fmt.Printf("g.Total = %d\n", g.Total)
	rlib.Errcheck(rows.Err())
	w.Header().Set("Content-Type", "application/json")
	g.Status = "success"
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
	fmt.Printf("Entered SvcFormHandlerReceipt\n")
	var err error
	if d.RCPTID, err = SvcExtractIDFromURI(r.RequestURI, "RCPTID", 3, w); err != nil {
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
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err)
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
	funcname := "saveReceipt"
	fmt.Printf("SvcFormHandlerReceipt save\n")
	fmt.Printf("record data = %s\n", d.data)

	var foo SaveReceiptInput
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	var a rlib.Receipt
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	fmt.Printf("saveReceipt - first migrate: a = %#v\n", a)

	var bar SaveReceiptOther
	if err := json.Unmarshal(data, &bar); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.Record.BID.ID]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, bar.Record.BID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	fmt.Printf("saveReceipt - second migrate: a = %#v\n", a)

	var err error
	if a.RCPTID == 0 && d.RCPTID == 0 {
		// This is a new Receipt
		fmt.Printf(">>>> NEW RECEIPT IS BEING ADDED\n")
		_, err = rlib.InsertReceipt(&a)
	} else {
		// update existing record
		err = rlib.UpdateReceipt(&a)
	}
	if err != nil {
		e := fmt.Errorf("%s: Error saving receipt (RCPTID=%d\n: %s", funcname, d.RCPTID, err.Error())
		SvcGridErrorReturn(w, e)
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
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
