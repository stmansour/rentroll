package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// PaymentTypeGrid contains the data from PaymentType that is targeted to the UI Grid that displays
// a list of PaymentType structs
type PaymentTypeGrid struct {
	Recid       int64 `json:"recid"`
	PMTID       int64
	BID         int64
	Name        string
	Description string
	LastModTime time.Time
	LastModBy   int64
}

// PaymentTypeSearchResponse is a response string to the search request for PaymentType records
type PaymentTypeSearchResponse struct {
	Status  string            `json:"status"`
	Total   int64             `json:"total"`
	Records []PaymentTypeGrid `json:"records"`
}

// PaymentTypeGridSave is the input data format for a Save command
type PaymentTypeGridSave struct {
	Status   string            `json:"status"`
	Recid    int64             `json:"recid"`
	FormName string            `json:"name"`
	Record   PaymentTypeGrid   `json:"record"`
	Changes  []PaymentTypeGrid `json:"changes"`
}

// PaymentTypeGetResponse is the response to a GetPaymentType request
type PaymentTypeGetResponse struct {
	Status string          `json:"status"`
	Record PaymentTypeGrid `json:"record"`
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
	fmt.Printf("Entered SvcHandlerPaymentType\n")
	fmt.Printf("Request: %s:  BID = %d,  PMTID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerPaymentTypes(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				SvcGridErrorReturn(w, fmt.Errorf("PaymentTypeID is required but was not specified"))
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
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err)
		return
	}
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
	funcname := "SvcSearchHandlerPaymentTypes"
	fmt.Printf("Entered %s\n", funcname)
	var (
		g   PaymentTypeSearchResponse
		err error
	)

	order := "PMTID ASC"                                                               // default ORDER
	q := fmt.Sprintf("SELECT %s FROM PaymentType ", rlib.RRdb.DBFields["PaymentType"]) // the fields we want
	qw := fmt.Sprintf("BID=%d", d.BID)
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
	fmt.Printf("rowcount query conditions: %s\ndb query = %s\n", qw, q)

	g.Total, err = GetRowCount("PaymentType", qw)
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
		var p rlib.PaymentType
		var q PaymentTypeGrid
		rlib.ReadPaymentTypes(rows, &p)
		rlib.MigrateStructVals(&p, &q)
		q.Recid = p.PMTID
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

// deletePaymentType deletes a payment type from the database
// wsdoc {
//  @Title  Delete Payment Type
//	@URL /v1/pmt/:BUI/:RAID
//  @Method  POST
//	@Synopsis Delete a Payment Type
//  @Desc  This service deletes a PaymentType.
//	@Input WebGridDelete
//  @Response SvcStatusResponse
// wsdoc }
func deletePaymentType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "deletePaymentType"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)
	var del WebGridDelete
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	for i := 0; i < len(del.Selected); i++ {
		if err := rlib.DeletePaymentType(del.Selected[i]); err != nil {
			SvcGridErrorReturn(w, err)
			return
		}
	}
	SvcWriteSuccessResponse(w)
}

// GetPaymentType returns the requested assessment
// wsdoc {
//  @Title  Save PaymentType
//	@URL /v1/pmt/:BUI/:PMTID
//  @Method  GET
//	@Synopsis Update the information on a PaymentType with the supplied data
//  @Description  This service updates PaymentType :PMTID with the information supplied. All fields must be supplied.
//	@Input PaymentTypeGridSave
//  @Response SvcStatusResponse
// wsdoc }
func savePaymentType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "savePaymentType"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var foo PaymentTypeGridSave
	data := []byte(d.data)
	err := json.Unmarshal(data, &foo)

	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	if len(foo.Changes) == 0 { // This is a new record
		var a rlib.PaymentType
		rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
		fmt.Printf("a = %#v\n", a)
		fmt.Printf(">>>> NEW PAYMENT TYPE IS BEING ADDED\n")
		if err = rlib.InsertPaymentType(&a); err != nil {
			e := fmt.Errorf("%s: Error saving assessment (PMTID=%d\n: %s", funcname, a.PMTID, err.Error())
			SvcGridErrorReturn(w, e)
			return
		}
	} else { // update existing or add new record(s)
		fmt.Printf("prior to JSONchangeParseUtil:  d.BID = %d\n", d.BID)
		if err = JSONchangeParseUtil(d.data, paymentTypeUpdate, d); err != nil {
			SvcGridErrorReturn(w, err)
			return
		}
	}
	SvcWriteSuccessResponse(w)
}

// paymentTypeUpdate unmarshals the supplied string. If Recid > 0 it updates the
// PaymentType record using Recid as the PMTID.  If Recid == 0, then it inserts a
// new PaymentType record.
func paymentTypeUpdate(s string, d *ServiceData) error {
	b := []byte(s)
	var rec PaymentTypeGrid
	if err := json.Unmarshal(b, &rec); err != nil { // first parse to determine the record ID we need to load
		return err
	}
	var pt rlib.PaymentType
	if rec.Recid > 0 { // is this an update?
		rlib.GetPaymentType(rec.Recid, &pt)            // now load that record...
		if err := json.Unmarshal(b, &pt); err != nil { // merge in the changes...
			return err
		}
		return rlib.UpdatePaymentType(&pt) // and save the result
	}
	// no, it is a new table entry that has not been saved...
	var a rlib.PaymentType
	if err := json.Unmarshal(b, &a); err != nil { // merge in the changes...
		return err
	}
	a.BID = d.BID
	fmt.Printf("a = %#v\n", a)
	fmt.Printf(">>>> NEW PAYMENT TYPE IS BEING ADDED\n")
	return rlib.InsertPaymentType(&a)
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
	funcname := "getPaymentType"
	fmt.Printf("entered %s\n", funcname)
	var g PaymentTypeGetResponse
	var a rlib.PaymentType
	rlib.GetPaymentType(d.ID, &a)
	if a.PMTID > 0 {
		var gg PaymentTypeGrid
		rlib.MigrateStructVals(&a, &gg)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
