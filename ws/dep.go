package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// DepositoryGrid contains the data from Depository that is targeted to the UI Grid that displays
// a list of Depository structs
type DepositoryGrid struct {
	Recid       int64 `json:"recid"`
	DEPID       int64
	BID         int64
	LID         int64
	Name        string
	AccountNo   string
	LdgrName    string
	GLNumber    string
	LastModTime time.Time
	LastModBy   int64
}

// DepositorySearchResponse is a response string to the search request for Depository records
type DepositorySearchResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []DepositoryGrid `json:"records"`
}

// DepositoryGridSave is the input data format for a Save command
type DepositoryGridSave struct {
	Status   string           `json:"status"`
	Recid    int64            `json:"recid"`
	FormName string           `json:"name"`
	Record   DepositoryGrid   `json:"record"`
	Changes  []DepositoryGrid `json:"changes"`
}

// DepositoryGetResponse is the response to a GetDepository request
type DepositoryGetResponse struct {
	Status string         `json:"status"`
	Record DepositoryGrid `json:"record"`
}

// SvcHandlerDepository formats a complete data record for an assessment for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the DEPID as follows:
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcHandlerDepository(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcHandlerDepository\n")
	fmt.Printf("Request: %s:  BID = %d,  DEPID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcSearchHandlerDepositories(w, r, d) // it is a query for the grid.
		} else {
			if d.ID < 0 {
				SvcGridErrorReturn(w, fmt.Errorf("DepositoryID is required but was not specified"))
				return
			}
			getDepository(w, r, d)
		}
		break
	case "save":
		saveDepository(w, r, d)
		break
	case "delete":
		deleteDepository(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err)
		return
	}
}

// SvcSearchHandlerDepositories generates a report of all Depositories defined business d.BID
// wsdoc {
//  @Title  Search Depositories
//	@URL /v1/dep/:BUI
//  @Method  POST
//	@Synopsis Search Depositories
//  @Descr  Search all Depository and return those that match the Search Logic.
//  @Descr  The search criteria includes start and stop dates of interest.
//	@Input WebGridSearchRequest
//  @Response DepositorySearchResponse
// wsdoc }
func SvcSearchHandlerDepositories(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcSearchHandlerDepositories"
	fmt.Printf("Entered %s\n", funcname)
	var (
		g   DepositorySearchResponse
		err error
	)

	order := "DEPID ASC"                                                             // default ORDER
	q := fmt.Sprintf("SELECT %s FROM Depository ", rlib.RRdb.DBFields["Depository"]) // the fields we want
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

	g.Total, err = GetRowCount("Depository", qw)
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
		var p rlib.Depository
		var q DepositoryGrid
		rlib.ReadDepositories(rows, &p)
		rlib.MigrateStructVals(&p, &q)
		q.Recid = p.DEPID
		ldg := rlib.GetLedger(p.LID)
		if ldg.LID == p.LID { // if it didn't read the ledger def, ldg.LID will == 0
			q.LdgrName = ldg.Name
			q.GLNumber = ldg.GLNumber
		}
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

// deleteDepository deletes a payment type from the database
// wsdoc {
//  @Title  Delete Depository
//	@URL /v1/dep/:BUI/:RAID
//  @Method  POST
//	@Synopsis Delete a Payment Type
//  @Desc  This service deletes a Depository.
//	@Input WebGridDelete
//  @Response SvcStatusResponse
// wsdoc }
func deleteDepository(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "deleteDepository"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)
	var del WebGridDelete
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	for i := 0; i < len(del.Selected); i++ {
		if err := rlib.DeleteDepository(del.Selected[i]); err != nil {
			SvcGridErrorReturn(w, err)
			return
		}
	}
	SvcWriteSuccessResponse(w)
}

// GetDepository returns the requested assessment
// wsdoc {
//  @Title  Save Depository
//	@URL /v1/dep/:BUI/:DEPID
//  @Method  GET
//	@Synopsis Update the information on a Depository with the supplied data
//  @Description  This service updates Depository :DEPID with the information supplied. All fields must be supplied.
//	@Input DepositoryGridSave
//  @Response SvcStatusResponse
// wsdoc }
func saveDepository(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveDepository"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var foo DepositoryGridSave
	data := []byte(d.data)
	err := json.Unmarshal(data, &foo)

	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	if len(foo.Changes) == 0 { // This is a new record
		var a rlib.Depository
		rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
		fmt.Printf("a = %#v\n", a)
		fmt.Printf(">>>> NEW PAYMENT TYPE IS BEING ADDED\n")
		_, err = rlib.InsertDepository(&a)
		if err != nil {
			e := fmt.Errorf("%s: Error saving assessment (DEPID=%d\n: %s", funcname, a.DEPID, err.Error())
			SvcGridErrorReturn(w, e)
			return
		}
	} else { // update existing or add new record(s)
		fmt.Printf("prior to JSONchangeParseUtil:  d.BID = %d\n", d.BID)
		if err = JSONchangeParseUtil(d.data, depositoryUpdate, d); err != nil {
			SvcGridErrorReturn(w, err)
			return
		}
	}
	SvcWriteSuccessResponse(w)
}

// depositoryUpdate unmarshals the supplied string. If Recid > 0 it updates the
// Depository record using Recid as the DEPID.  If Recid == 0, then it inserts a
// new Depository record.
func depositoryUpdate(s string, d *ServiceData) error {
	var err error
	b := []byte(s)
	var rec DepositoryGrid
	if err = json.Unmarshal(b, &rec); err != nil { // first parse to determine the record ID we need to load
		return err
	}
	if rec.Recid > 0 { // is this an update?
		pt, err := rlib.GetDepository(rec.Recid) // now load that record...
		if err != nil {
			return err
		}
		if err = json.Unmarshal(b, &pt); err != nil { // merge in the changes...
			return err
		}
		return rlib.UpdateDepository(&pt) // and save the result
	}
	// no, it is a new table entry that has not been saved...
	var a rlib.Depository
	if err := json.Unmarshal(b, &a); err != nil { // merge in the changes...
		return err
	}
	a.BID = d.BID
	fmt.Printf("a = %#v\n", a)
	fmt.Printf(">>>> NEW DEPOSITORY IS BEING ADDED\n")
	_, err = rlib.InsertDepository(&a)
	return err
}

// GetDepository returns the requested assessment
// wsdoc {
//  @Title  Get Depository
//	@URL /v1/dep/:BUI/:DEPID
//  @Method  GET
//	@Synopsis Get information on a Depository
//  @Description  Return all fields for assessment :DEPID
//	@Input WebGridSearchRequest
//  @Response DepositoryGetResponse
// wsdoc }
func getDepository(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "getDepository"
	fmt.Printf("entered %s\n", funcname)
	var g DepositoryGetResponse
	a, err := rlib.GetDepository(d.ID)
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	if a.DEPID > 0 {
		var gg DepositoryGrid
		rlib.MigrateStructVals(&a, &gg)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
