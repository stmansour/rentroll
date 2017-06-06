package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// AssessmentSendForm is the outbound structure specifically for the UI. It will be
// automatically populated from an rlib.Assessment struct.
type AssessmentSendForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	ASMID          int64 // unique id for this assessment
	BID            rlib.XJSONBud
	PASMID         int64
	RID            int64
	Rentable       string
	RAID           int64
	Amount         float64
	Start          rlib.JSONTime
	Stop           rlib.JSONTime
	RentCycle      rlib.XJSONCycleFreq
	ProrationCycle rlib.XJSONCycleFreq
	InvoiceNo      int64
	ARID           int64
	Comment        string
	LastModTime    rlib.JSONTime
	LastModBy      int64
}

// AssessmentSaveForm is a structure specifically for the return value from w2ui.
// Data does not always come back in the same format it was sent. For example,
// values from dropdown lists come back in the form of a rlib.W2uiHTMLSelect struct.
// So, we break up the ingest into 2 parts. First, we read back the fields that look
// just like the xxxSendForm -- this is what is in xxxSaveForm. Then we readback
// the data that has changed, which is in the xxxSaveOther struct.  All this data
// is merged into the appropriate database structure using MigrateStructData.
type AssessmentSaveForm struct {
	Recid       int64 `json:"recid"` // this is to support the w2ui form
	ASMID       int64
	PASMID      int64
	RID         int64
	RAID        int64
	Amount      float64
	Start       rlib.JSONTime
	Stop        rlib.JSONTime
	InvoiceNo   int64
	Comment     string
	LastModTime rlib.JSONTime
	LastModBy   int64
}

// AssessmentSaveOther is a struct to handle the UI list box selections
type AssessmentSaveOther struct {
	BID            rlib.W2uiHTMLSelect
	RentCycle      rlib.W2uiHTMLSelect
	ProrationCycle rlib.W2uiHTMLSelect
	ARID           rlib.W2uiHTMLSelect
}

// AssessmentGrid is a structure specifically for the UI Grid.
type AssessmentGrid struct {
	Recid    int64  `json:"recid"` // this is to support the w2ui form
	ASMID    int64  // unique id for this assessment
	BID      int64  // which business
	PASMID   int64  // parent Assessment, if this is non-zero it means this assessment is an instance of the recurring assessment with id PASMID. When non-zero DO NOT process as a recurring assessment, it is an instance
	RID      int64  // the Rentable
	Rentable string // the RentableName
	// ATypeLID  int64         // what type of assessment
	RAID      int64           // associated Rental Agreement
	RentCycle int64           // Rent Cycle
	Amount    float64         // how much
	Start     rlib.JSONTime   // start time
	Stop      rlib.JSONTime   // stop time, may be the same as start time or later
	InvoiceNo int64           // A uniqueID for the invoice number
	ARID      int64           // which account rule
	AcctRule  rlib.NullString // expression showing how to account for the amount
}

// SearchAssessmentsResponse is a response string to the search request for assessments
type SearchAssessmentsResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []AssessmentGrid `json:"records"`
}

// SaveAssessmentInput is the input data format for a Save command
type SaveAssessmentInput struct {
	Status   string             `json:"status"`
	Recid    int64              `json:"recid"`
	FormName string             `json:"name"`
	Record   AssessmentSaveForm `json:"record"`
}

// SaveAssessmentOther is the input data format for the "other" data on the Save command
type SaveAssessmentOther struct {
	Status string              `json:"status"`
	Recid  int64               `json:"recid"`
	Name   string              `json:"name"`
	ARID   int64               `json:"ARID"`
	Record AssessmentSaveOther `json:"record"`
}

// GetAssessmentResponse is the response to a GetAssessment request
type GetAssessmentResponse struct {
	Status string             `json:"status"`
	Record AssessmentSendForm `json:"record"`
}

// assessmentGridRowScan scans a result from sql row and dump it in a AssessmentGrid struct
func assessmentGridRowScan(rows *sql.Rows, q AssessmentGrid) (AssessmentGrid, error) {
	err := rows.Scan(&q.ASMID, &q.BID, &q.PASMID, &q.RID, &q.Rentable, &q.RAID, &q.RentCycle, &q.Amount, &q.Start, &q.Stop, &q.InvoiceNo, &q.ARID, &q.AcctRule)
	return q, err
}

// which fields needs to be fetched for SQL query for assessment grid
var asmFieldsMap = map[string][]string{
	"ASMID":        {"Assessments.ASMID"},
	"BID":          {"Assessments.BID"},
	"PASMID":       {"Assessments.PASMID"},
	"RID":          {"Assessments.RID"},
	"RentableName": {"Rentable.RentableName"},
	"RAID":         {"Assessments.RAID"},
	"RentCycle":    {"Assessments.RentCycle"},
	"Amount":       {"Assessments.Amount"},
	"Start":        {"Assessments.Start"},
	"Stop":         {"Assessments.Stop"},
	"InvoiceNo":    {"Assessments.InvoiceNo"},
	"ARID":         {"Assessments.ARID"},
	"AcctRule":     {"AR.Name"},
}

// which fields needs to be fetched for SQL query for assessment grid
var asmQuerySelectFields = []string{
	"Assessments.ASMID",
	"Assessments.BID",
	"Assessments.PASMID",
	"Assessments.RID",
	"Rentable.RentableName",
	"Assessments.RAID",
	"Assessments.RentCycle",
	"Assessments.Amount",
	"Assessments.Start",
	"Assessments.Stop",
	"Assessments.InvoiceNo",
	"Assessments.ARID",
	"AR.Name",
}

// SvcSearchHandlerAssessments generates a report of all Assessments defined business d.BID
// wsdoc {
//  @Title  Search Assessments
//	@URL /v1/asms/:BUI
//  @Method  POST
//	@Synopsis Search Assessments
//  @Description  Search all Assessments and return those that match the Search Logic.
//  @Descr The search criteria includes start and stop dates of interest.
//	@Input WebGridSearchRequest
//  @Response SearchAssessmentsResponse
// wsdoc }
func SvcSearchHandlerAssessments(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "SvcSearchHandlerAssessments"
		g        SearchAssessmentsResponse
		err      error
	)

	fmt.Printf("Entered %s\n", funcname)

	whr := `Assessments.BID = %d AND Assessments.Stop > %q AND Assessments.Start < %q`
	whr = fmt.Sprintf(whr, d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	order := `Start ASC, RAID ASC` // default ORDER

	// get where clause and order clause for sql query
	_, orderClause := GetSearchAndSortSQL(d, asmFieldsMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	asmQuery := `
	SELECT
		{{.SelectClause}}
	FROM Assessments
	INNER JOIN Rentable ON Assessments.RID=Rentable.RID
	LEFT JOIN AR ON Assessments.ARID=AR.ARID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := queryClauses{
		"SelectClause": strings.Join(asmQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := renderSQLQuery(asmQuery, qc)
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
	asmQueryWithLimit := asmQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := renderSQLQuery(asmQueryWithLimit, qc)
	fmt.Printf("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q AssessmentGrid
		q.Recid = i

		q, err = assessmentGridRowScan(rows, q)
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

	var (
		funcname = "SvcFormHandlerAssessment"
		err      error
	)

	fmt.Printf("Entered %s\n", funcname)

	if d.ASMID, err = SvcExtractIDFromURI(r.RequestURI, "ASMID", 3, w); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	fmt.Printf("Request: %s:  BID = %d,  ASMID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ASMID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getAssessment(w, r, d)
		break
	case "save":
		saveAssessment(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

// GetAssessment returns the requested assessment
// wsdoc {
//  @Title  Save Assessment
//	@URL /v1/asm/:BUI/:ASMID
//  @Method  POST
//	@Synopsis Update the information on a Assessment with the supplied data
//  @Description  This service updates Assessment :ASMID with the information supplied. All fields must be supplied.
//	@Input SaveAssessmentInput
//  @Response SvcStatusResponse
// wsdoc }
func saveAssessment(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "saveAssessment"
		err      error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var foo SaveAssessmentInput
	data := []byte(d.data)

	err = json.Unmarshal(data, &foo)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	//----------------------------------------------------------
	// Parse the standard variables from the return struct...
	//----------------------------------------------------------
	var a rlib.Assessment
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	//----------------------------------------------------------
	// Now get the other variables and copy them into a...
	//----------------------------------------------------------
	var bar SaveAssessmentOther
	err = json.Unmarshal(data, &bar) // and now the other variables
	fmt.Printf("\n\n#################\nafter unmarshal, bar = %#v\n#################\n", bar)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.Record.BID.ID]
	if !ok {
		e := fmt.Errorf("Could not map BID value: %s", bar.Record.BID.ID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	a.RentCycle = rlib.CycleFreqMap[bar.Record.RentCycle.ID]
	a.ProrationCycle = rlib.CycleFreqMap[bar.Record.ProrationCycle.ID]
	a.ARID, err = strconv.ParseInt(bar.Record.ARID.ID, 10, 64)
	if err != nil {
		e := fmt.Errorf("Could not convert ARID %s to an int", bar.Record.BID.ID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	fmt.Printf("after conversion: a.ARID = %d\n", a.ARID)

	// Now just update the database
	if a.ASMID == 0 && d.ASMID == 0 {
		// This is a new record
		fmt.Printf(">>>> NEW ASSESSMENT IS BEING ADDED\n")
		_, err = rlib.InsertAssessment(&a)

	} else if a.ASMID > 0 || d.ASMID > 0 {
		fmt.Printf(">>>> UPDATE EXISTING ASSESSMENT  ASMID = %d\n", a.ASMID)
		err = rlib.UpdateAssessment(&a)
	} else {
		err = fmt.Errorf("Unknown state: note an update, and not a new record")
	}
	if err != nil {
		e := fmt.Errorf("Error saving assessment (ASMID=%d\n: %s", d.ASMID, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	SvcWriteSuccessResponse(w)
}

var asmFormSelectFields = []string{
	"Assessments.PASMID",
	"Assessments.RID",
	"Rentable.RentableName",
	"Assessments.RAID",
	"Assessments.Amount",
	"Assessments.Start",
	"Assessments.Stop",
	"Assessments.RentCycle",
	"Assessments.ProrationCycle",
	"Assessments.InvoiceNo",
	"Assessments.ARID",
	"Assessments.Comment",
	"Assessments.LastModTime",
	"Assessments.LastModBy",
}

// GetAssessment returns the requested assessment
// wsdoc {
//  @Title  Get Assessment
//	@URL /v1/asm/:BUI/:ASMID
//  @Method  GET
//	@Synopsis Get information on a Assessment
//  @Description  Return all fields for assessment :ASMID
//	@Input WebGridSearchRequest
//  @Response GetAssessmentResponse
// wsdoc }
func getAssessment(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "getAssessment"
		g        GetAssessmentResponse
		err      error
	)

	fmt.Printf("entered %s\n", funcname)

	asmQuery := `
	SELECT
		{{.SelectClause}}
	FROM Assessments
	INNER JOIN Rentable ON Assessments.RID=Rentable.RID
	WHERE {{.WhereClause}};`

	// will be substituted as query clauses
	qc := queryClauses{
		"SelectClause": strings.Join(asmFormSelectFields, ","),
		"WhereClause":  fmt.Sprintf("Assessments.BID=%d AND Assessments.ASMID=%d", d.BID, d.ASMID),
	}

	// get formatted query with substitution of select, where, order clause
	q := renderSQLQuery(asmQuery, qc)
	fmt.Printf("db query = %s\n", q)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(q)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var gg AssessmentSendForm
		gg.ASMID = d.ASMID

		// get bud for BID field
		for bud, bid := range rlib.RRdb.BUDlist {
			if bid == d.BID {
				gg.BID = rlib.XJSONBud(bud)
				break
			}
		}

		var rentCycle, prorationCycle int64

		err = rows.Scan(&gg.PASMID, &gg.RID, &gg.Rentable, &gg.RAID, &gg.Amount, &gg.Start, &gg.Stop, &rentCycle, &prorationCycle, &gg.InvoiceNo, &gg.ARID, &gg.Comment, &gg.LastModTime, &gg.LastModBy)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		for freqStr, freqNo := range rlib.CycleFreqMap {
			if rentCycle == freqNo {
				gg.RentCycle = rlib.XJSONCycleFreq(freqStr)
			}
			if prorationCycle == freqNo {
				gg.ProrationCycle = rlib.XJSONCycleFreq(freqStr)
			}
		}

		g.Record = gg
	}
	// error check
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	// write response
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
