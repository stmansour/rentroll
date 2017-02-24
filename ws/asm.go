package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
)

// // Assessment is a charge associated with a Rentable
// type Assessment struct {
// 	ASMID          int64     // unique id for this assessment
// 	PASMID         int64     // parent Assessment, if this is non-zero it means this assessment is an instance of the recurring assessment with id PASMID. When non-zero DO NOT process as a recurring assessment, it is an instance
// 	BID            int64     // what Business
// 	RID            int64     // the Rentable
// 	ATypeLID       int64     // what type of assessment
// 	RAID           int64     // associated Rental Agreement
// 	Amount         float64   // how much
// 	Start          time.Time // start time
// 	Stop           time.Time // stop time, may be the same as start time or later
// 	RentCycle      int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, G = quarterly, 8 = yearly
// 	ProrationCycle int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
// 	InvoiceNo      int64     // A uniqueID for the invoice number
// 	AcctRule       string    // expression showing how to account for the amount
// 	Comment        string
// 	LastModTime    time.Time
// 	LastModBy      int64
// }

// AssessmentForm is a structure specifically for the UI. It will be
// automatically populated from an rlib.Assessment struct
type AssessmentForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	ASMID          int64 // unique id for this assessment
	BID            rlib.XJSONBud
	PASMID         int64               // parent Assessment, if this is non-zero it means this assessment is an instance of the recurring assessment with id PASMID. When non-zero DO NOT process as a recurring assessment, it is an instance
	RID            int64               // the Rentable
	ATypeLID       int64               // what type of assessment
	RAID           int64               // associated Rental Agreement
	Amount         float64             // how much
	Start          rlib.JSONTime       // start time
	Stop           rlib.JSONTime       // stop time, may be the same as start time or later
	RentCycle      rlib.XJSONCycleFreq // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, G = quarterly, 8 = yearly
	ProrationCycle rlib.XJSONCycleFreq // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
	InvoiceNo      int64               // A uniqueID for the invoice number
	AcctRule       string              // expression showing how to account for the amount
	Comment        string
	LastModTime    rlib.JSONTime
	LastModBy      int64
}

// AssessmentOther is a struct to handle the UI list box selections
type AssessmentOther struct {
	BID rlib.W2uiHTMLSelect
}

// AssessmentGrid is a structure specifically for the UI Grid.
type AssessmentGrid struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	ASMID          int64 // unique id for this assessment
	BID            rlib.XJSONBud
	PASMID         int64               // parent Assessment, if this is non-zero it means this assessment is an instance of the recurring assessment with id PASMID. When non-zero DO NOT process as a recurring assessment, it is an instance
	RID            int64               // the Rentable
	ATypeLID       int64               // what type of assessment
	RAID           int64               // associated Rental Agreement
	Amount         float64             // how much
	Start          rlib.JSONTime       // start time
	Stop           rlib.JSONTime       // stop time, may be the same as start time or later
	RentCycle      rlib.XJSONCycleFreq // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, G = quarterly, 8 = yearly
	ProrationCycle rlib.XJSONCycleFreq // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
	InvoiceNo      int64               // A uniqueID for the invoice number
	AcctRule       string              // expression showing how to account for the amount
}

// SearchAssessmentsResponse is a response string to the search request for assessments
type SearchAssessmentsResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []AssessmentGrid `json:"records"`
}

// GetAssessmentResponse is the response to a GetAssessment request
type GetAssessmentResponse struct {
	Status string         `json:"status"`
	Record AssessmentForm `json:"record"`
}

// SvcSearchHandlerAssessments generates a report of all Assessments defined business d.BID
// wsdoc {
//  @Title  Search Assessments
//	@URL /v1/asm/:BUI
//  @Method  POST
//	@Synopsis Search Assessments
//  @Description  Search all Assessments and return those that match the Search Logic
//	@Input WebRequest
//  @Response SearchAssessmentsResponse
// wsdoc }
func SvcSearchHandlerAssessments(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	fmt.Printf("Entered SvcSearchHandlerAssessments\n")

	var p rlib.Assessment
	var err error
	var g SearchAssessmentsResponse

	// TODO: Add dates to default search -- this month
	srch := fmt.Sprintf("BID=%d", d.BID) // default WHERE clause
	order := "Start ASC"                 // default ORDER
	q, qw := gridBuildQuery("Assessments", srch, order, d, &p)

	// set g.Total to the total number of rows of this data...
	g.Total, err = GetRowCount("Assessments", qw)
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
		var p rlib.Assessment
		var q AssessmentGrid
		rlib.ReadAssessments(rows, &p)
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

// SvcFormHandlerAssessment formats a complete data record for an assessment for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the ASMID as follows:
//           0  1    2   3
// uri 		/v1/asm/:BUI/ASMID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerAssessment(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcFormHandlerAssessment\n")

	var err error
	if d.ASMID, err = SvcExtractIDFromURI(r.RequestURI, "ASMID", 3, w); err != nil {
		return
	}

	fmt.Printf("Request: %s:  BID = %d,  ASMID = %d\n", d.webreq.Cmd, d.BID, d.ASMID)

	switch d.webreq.Cmd {
	case "get":
		getAssessment(w, r, d)
		break
	case "save":
		saveAssessment(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.webreq.Cmd)
		SvcGridErrorReturn(w, err)
		return
	}
}

// GetAssessment returns the requested assessment
// wsdoc {
//  @Title  Save Assessment
//	@URL /v1/asm/:BUI/:ASMID
//  @Method  GET
//	@Synopsis Update the information on a Assessment with the supplied data
//  @Description  This service updates Assessment :ASMID with the information supplied. All fields must be supplied.
//	@Input WebRequest
//  @Response SvcStatusResponse
// wsdoc }
func saveAssessment(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveAssessment"
	target := `"record":`
	fmt.Printf("SvcFormHandlerAssessment save\n")
	fmt.Printf("record data = %s\n", d.data)
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("%s: cannot find %s in form json", funcname, target)
		SvcGridErrorReturn(w, e)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	var foo AssessmentForm
	err := json.Unmarshal([]byte(s), &foo)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	// migrate the variables that transfer without needing special handling...
	var a rlib.Assessment
	rlib.MigrateStructVals(&foo, &a)

	// now get the stuff that requires special handling...
	var bar AssessmentOther
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
	err = rlib.UpdateAssessment(&a)
	if err != nil {
		e := fmt.Errorf("%s: Error updating assessment: %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	SvcWriteSuccessResponse(w)
}

// GetAssessment returns the requested assessment
// wsdoc {
//  @Title  Get Assessment
//	@URL /v1/asm/:BUI/:ASMID
//  @Method  GET
//	@Synopsis Get information on a Assessment
//  @Description  Return all fields for assessment :ASMID
//	@Input WebRequest
//  @Response GetAssessmentResponse
// wsdoc }
func getAssessment(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "getAssessment"
	fmt.Printf("entered getAssessment\n")
	var g GetAssessmentResponse
	a, err := rlib.GetAssessment(d.ASMID)
	if err != nil {
		e := fmt.Errorf("%s: Error reading assessments: %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	if a.ASMID > 0 {
		var gg AssessmentForm
		rlib.MigrateStructVals(&a, &gg)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
