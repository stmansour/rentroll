package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
)

// ReceiptForm is a structure specifically for the UI. It will be
// automatically populated from an rlib.Receipt struct
type ReceiptForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	RCPTID         int64
	PRCPTID        int64 // Parent RCPTID, points to RCPT being amended/corrected by this receipt
	BID            rlib.XJSONBud
	RAID           int64
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

// ReceiptOther is a struct to handle the UI list box selections
type ReceiptOther struct {
	BID rlib.W2uiHTMLSelect
}

// PrReceiptGrid is a structure specifically for the UI Grid.
type PrReceiptGrid struct {
	Recid  int64 `json:"recid"` // this is to support the w2ui form
	RCPTID int64
	BID    rlib.XJSONBud
	RAID   int64
	PMTID  int64
	Dt     rlib.JSONTime
	DocNo  string // check number, money order number, etc.; documents the payment
	Amount float64
}

// SearchReceiptsResponse is a response string to the search request for receipts
type SearchReceiptsResponse struct {
	Status  string          `json:"status"`
	Total   int64           `json:"total"`
	Records []PrReceiptGrid `json:"records"`
}

// GetReceiptResponse is the response to a GetReceipt request
type GetReceiptResponse struct {
	Status string      `json:"status"`
	Record ReceiptForm `json:"record"`
}

// SvcSearchHandlerReceipts generates a report of all Receipts defined business d.BID
// wsdoc {
//  @Title  Search Receipts
//	@URL /v1/receipts/:BUI
//  @Method  POST
//	@Synopsis Search Receipts
//  @Description  Search all Receipts and return those that match the Search Logic
//	@Input WebRequest
//  @Response SearchReceiptsResponse
// wsdoc }
func SvcSearchHandlerReceipts(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	fmt.Printf("Entered SvcSearchHandlerReceipts\n")

	var p rlib.Receipt
	var err error
	var g SearchReceiptsResponse

	// TODO: Add dates to default search -- this month
	srch := fmt.Sprintf("BID=%d", d.BID) // default WHERE clause
	order := "Dt ASC"                    // default ORDER
	q, qw := gridBuildQuery("Receipt", srch, order, d, &p)

	// set g.Total to the total number of rows of this data...
	g.Total, err = GetRowCount("Receipt", qw)
	if err != nil {
		fmt.Printf("Error from GetRowCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}

	fmt.Printf("db query = %s\n", q)

	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	i := int64(d.webreq.Offset)
	count := 0
	for rows.Next() {
		var p rlib.Receipt
		var q PrReceiptGrid
		rlib.ReadReceipts(rows, &p)
		rlib.MigrateStructVals(&p, &q)
		q.Recid = i
		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.webreq.Limit {
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
//           0    1         2   3
// uri 		/v1/receipt/BUD/RCPTID
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

	fmt.Printf("Request: %s:  BID = %d,  RCPTID = %d\n", d.webreq.Cmd, d.BID, d.RCPTID)

	switch d.webreq.Cmd {
	case "get":
		getReceipt(w, r, d)
		break
	case "save":
		saveReceipt(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.webreq.Cmd)
		SvcGridErrorReturn(w, err)
		return
	}
}

// GetReceipt returns the requested receipt
// wsdoc {
//  @Title  Save Receipt
//	@URL /v1/receipt/:BUI/:RCPTID
//  @Method  GET
//	@Synopsis Update the information on a Receipt with the supplied data
//  @Description  This service updates Receipt :RCPTID with the information supplied. All fields must be supplied.
//	@Input WebRequest
//  @Response SvcStatusResponse
// wsdoc }
func saveReceipt(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveReceipt"
	target := `"record":`
	fmt.Printf("SvcFormHandlerReceipt save\n")
	fmt.Printf("record data = %s\n", d.data)
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("%s: cannot find %s in form json", funcname, target)
		SvcGridErrorReturn(w, e)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	var foo ReceiptForm
	err := json.Unmarshal([]byte(s), &foo)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	// migrate the variables that transfer without needing special handling...
	var a rlib.Receipt
	rlib.MigrateStructVals(&foo, &a)

	// now get the stuff that requires special handling...
	var bar ReceiptOther
	err = json.Unmarshal([]byte(s), &bar)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.BID.ID]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, bar.BID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	// Now just update the database
	err = rlib.UpdateReceipt(&a)
	if err != nil {
		e := fmt.Errorf("%s: Error updating receipt: %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	SvcWriteSuccessResponse(w)
}

// GetReceipt returns the requested receipt
// wsdoc {
//  @Title  Get Receipt
//	@URL /v1/receipt/:BUI/:RCPTID
//  @Method  GET
//	@Synopsis Get information on a Receipt
//  @Description  Return all fields for receipt :RCPTID
//	@Input WebRequest
//  @Response GetReceiptResponse
// wsdoc }
func getReceipt(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered getReceipt\n")
	var g GetReceiptResponse
	a := rlib.GetReceiptNoAllocations(d.RCPTID)
	if a.RCPTID > 0 {
		var gg ReceiptForm
		rlib.MigrateStructVals(&a, &gg)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
