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

// AssessmentSendForm is the outbound structure specifically for the UI. It will be
// automatically populated from an rlib.Assessment struct.
type AssessmentSendForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	ASMID          int64 // unique id for this assessment
	BID            int64
	BUD            rlib.XJSONBud
	PASMID         int64
	RID            int64
	Rentable       rlib.NullString
	RAID           int64
	Amount         float64
	Start          rlib.JSONDate
	Stop           rlib.JSONDate
	RentCycle      int64
	ProrationCycle int64
	InvoiceNo      int64
	ARID           int64
	Comment        string
	LastModTime    rlib.JSONDateTime
	LastModBy      int64
	LastModByUser  string
	CreateTS       rlib.JSONDateTime
	CreateBy       int64
	CreateByUser   string
	ExpandPastInst int // if this is a new  Assessment and its epoch date is in the past, do we create instances in the past after saving the recurring Assessment?
	FLAGS          uint64
	Mode           int // initializes edit mode: 0 = this instance only, 1 = this and future, 2 = all
}

// AssessmentSaveForm is a structure specifically for the return value from w2ui.
// Data does not always come back in the same format it was sent. For example,
// values from dropdown lists come back in the form of a rlib.W2uiHTMLSelect struct.
// So, we break up the ingest into 2 parts. First, we read back the fields that look
// just like the xxxSendForm -- this is what is in xxxSaveForm. Then we readback
// the data that has changed, which is in the xxxSaveOther struct.  All this data
// is merged into the appropriate database structure using MigrateStructData.
type AssessmentSaveForm struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	BID            int64
	BUD            rlib.XJSONBud
	RentCycle      int64
	ProrationCycle int64
	ARID           int64
	ASMID          int64
	PASMID         int64
	RID            int64
	RAID           int64
	Amount         float64
	Start          rlib.JSONDate
	Stop           rlib.JSONDate
	InvoiceNo      int64
	Comment        string
	ReverseMode    int // if this a Reversal (delete), then 0 = this instance only, 1 = this and future instances, 2 = all instances
	ExpandPastInst int // if this is a new  Assessment and its epoch date is in the past, do we create instances in the past after saving the recurring Assessment?
	FLAGS          uint64
	Mode           int // 0 = this instance only, 1 = this and future, 2 = all
}

// AssessmentGrid is a structure specifically for the UI Grid.
type AssessmentGrid struct {
	Recid     int64           `json:"recid"` // this is to support the w2ui form
	ASMID     int64           // unique id for this assessment
	BID       int64           // which business
	PASMID    int64           // parent Assessment, if this is non-zero it means this assessment is an instance of the recurring assessment with id PASMID. When non-zero DO NOT process as a recurring assessment, it is an instance
	RID       int64           // the Rentable
	Rentable  rlib.NullString // the RentableName
	RAID      int64           // associated Rental Agreement
	RentCycle int64           // Rent Cycle
	Amount    float64         // how much
	Start     rlib.JSONDate   // start time
	Stop      rlib.JSONDate   // stop time, may be the same as start time or later
	InvoiceNo int64           // A uniqueID for the invoice number
	ARID      int64           // which account rule
	AcctRule  rlib.NullString // expression showing how to account for the amount
	FLAGS     uint64
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

// GetAssessmentResponse is the response to a GetAssessment request
type GetAssessmentResponse struct {
	Status string             `json:"status"`
	Record AssessmentSendForm `json:"record"`
}

// DeleteAsmForm holds ASMID to delete it
type DeleteAsmForm struct {
	ASMID       int64 // which assessment
	ReverseMode int   // 0 = this instance only, 1 = this and future instances, 2 = all instances
}

// assessmentGridRowScan scans a result from sql row and dump it in a AssessmentGrid struct
func assessmentGridRowScan(rows *sql.Rows, q AssessmentGrid) (AssessmentGrid, error) {
	err := rows.Scan(&q.ASMID, &q.BID, &q.PASMID, &q.RID, &q.Rentable, &q.RAID, &q.RentCycle, &q.Amount, &q.Start, &q.Stop, &q.InvoiceNo, &q.ARID, &q.AcctRule, &q.FLAGS)
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
	"FLAGS":        {"Assessments.FLAGS"},
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
	"Assessments.FLAGS",
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
	const funcname = "SvcSearchHandlerAssessments"
	var (
		g   SearchAssessmentsResponse
		err error
	)

	rlib.Console("Entered %s\n", funcname)

	//  for Instances DtFrom <= Assessments.Start AND Assessments.Start < DtTo, for epochs does the Search Start/Stop overlap the Assessment Range
	whr := `Assessments.BID = %d AND (((Assessments.PASMID!=0 OR Assessments.RentCycle=0) AND %q <= Assessments.Start AND Assessments.Start < %q) OR (Assessments.PASMID=0 AND Assessments.Stop >= %q AND Assessments.Start < %q))`
	whr = fmt.Sprintf(whr, d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	order := `Assessments.Start ASC, Assessments.RAID ASC, Assessments.ASMID ASC` // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, asmFieldsMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	asmQuery := `
	SELECT DISTINCT
		{{.SelectClause}}
	FROM Assessments
	LEFT JOIN Rentable ON Assessments.RID=Rentable.RID
	LEFT JOIN AR ON Assessments.ARID=AR.ARID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(asmQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(asmQuery, qc)
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
	asmQueryWithLimit := asmQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(asmQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
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
	const funcname = "SvcFormHandlerAssessment"
	var (
		err error
	)

	rlib.Console("Entered %s\n", funcname)

	if d.ASMID, err = SvcExtractIDFromURI(r.RequestURI, "ASMID", 3, w); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Request: %s:  BID = %d,  ASMID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ASMID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getAssessment(w, r, d)
		break
	case "save":
		saveAssessment(w, r, d)
		break
	case "delete":
		deleteAssessment(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
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
	const funcname = "saveAssessment"
	var (
		err     error
		errlist []bizlogic.BizError
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	var foo SaveAssessmentInput
	data := []byte(d.data)

	err = json.Unmarshal(data, &foo)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("\nAfter unmarshal: foo = %#v\n", foo)
	//----------------------------------------------------------
	// Parse the standard variables from the return struct...
	//----------------------------------------------------------
	var a rlib.Assessment
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	rlib.Console("\nAfter MigrateStructVals: a = %#v\n", a)
	rlib.Console("Start = %s, Stop = %s\n\n", a.Start.Format(rlib.RRDATEINPFMT), a.Stop.Format(rlib.RRDATEINPFMT))

	// Now just update the database
	if a.ASMID == 0 && d.ASMID == 0 {
		errlist := bizlogic.InsertAssessment(r.Context(), &a, foo.Record.ExpandPastInst)
		if len(errlist) > 0 {
			SvcErrListReturn(w, errlist, funcname)
			return
		}
	} else if a.ASMID > 0 || d.ASMID > 0 {
		rlib.Console(">>>> UPDATE EXISTING ASSESSMENT  ASMID = %d\n", a.ASMID)
		now := time.Now() // mark Assessment reversed at this time
		errlist = bizlogic.UpdateAssessment(r.Context(), &a, foo.Record.Mode, &now, foo.Record.ExpandPastInst)
		if len(errlist) > 0 {
			SvcErrListReturn(w, errlist, funcname)
			return
		}
	} else {
		err = fmt.Errorf("Unknown state: note an update, and not a new record")
	}
	if err != nil {
		e := fmt.Errorf("Error saving assessment (ASMID=%d): %s", d.ASMID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
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
	"Assessments.CreateTS",
	"Assessments.CreateBy",
	"Assessments.FLAGS",
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
	const funcname = "getAssessment"
	var (
		g   GetAssessmentResponse
		err error
	)

	rlib.Console("entered %s\n", funcname)

	asmQuery := `
	SELECT
		{{.SelectClause}}
	FROM Assessments
	LEFT JOIN Rentable ON Assessments.RID=Rentable.RID
	WHERE {{.WhereClause}};`

	// will be substituted as query clauses
	qc := rlib.QueryClause{
		"SelectClause": strings.Join(asmFormSelectFields, ","),
		"WhereClause":  fmt.Sprintf("Assessments.BID=%d AND Assessments.ASMID=%d", d.BID, d.ASMID),
	}

	// get formatted query with substitution of select, where, order clause
	q := rlib.RenderSQLQuery(asmQuery, qc)
	rlib.Console("db query = %s\n", q)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(q)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var gg AssessmentSendForm
		gg.ASMID = d.ASMID
		gg.BID = d.BID
		gg.BUD = rlib.GetBUDFromBIDList(gg.BID)

		err = rows.Scan(&gg.PASMID, &gg.RID, &gg.Rentable, &gg.RAID, &gg.Amount, &gg.Start, &gg.Stop, &gg.RentCycle, &gg.ProrationCycle, &gg.InvoiceNo, &gg.ARID, &gg.Comment, &gg.LastModTime, &gg.LastModBy, &gg.CreateTS, &gg.CreateBy, &gg.FLAGS)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		gg.ExpandPastInst = 1 // assume we don't expand unless told otherwise

		gg.CreateByUser = rlib.GetNameForUID(r.Context(), gg.CreateBy)
		gg.LastModByUser = rlib.GetNameForUID(r.Context(), gg.LastModBy)
		g.Record = gg
	}
	// error check
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// write response
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// deleteAssessment deletes the requested assessment with ASMID
// wsdoc {
//  @Title  Delete Assessment
//	@URL /v1/asm/:BUI/:ASMID
//  @Method  POST
//	@Synopsis Delete an Assessment record and associate Journal entries
//  @Description  Delete Assessment for requested ASMID.  Only use this
//  @Description  command if you really know what you're doing.
//	@Input DeleteAsmForm
//  @Response SvcWriteSuccessResponse
// wsdoc }
func deleteAssessment(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteAssessment"
	var (
		del DeleteAsmForm
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	a, err := rlib.GetAssessment(r.Context(), del.ASMID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Reversal Mode = %d\n", del.ReverseMode)
	//-------------------------------------------------------
	// GET THE NEW `tx`, UPDATED CTX FROM THE REQUEST CONTEXT
	//-------------------------------------------------------
	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// reverse assessment in atomic transaction
	now := time.Now() // mark Assessment reversed at this time
	errlist := bizlogic.ReverseAssessment(ctx, &a, del.ReverseMode, &now)
	if len(errlist) > 0 {
		tx.Rollback()
		SvcErrListReturn(w, errlist, funcname)
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

	SvcWriteSuccessResponse(d.BID, w)
}
