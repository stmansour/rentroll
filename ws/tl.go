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

// TaskList is the structure describing a task list definition
type TaskList struct {
	Recid       int64 `json:"recid"`
	TLID        int64
	TLDID       int64
	BID         int64
	Name        string
	Cycle       int64
	DtDone      rlib.JSONDate
	DtDue       rlib.JSONDate
	DtPreDue    rlib.JSONDate
	DtPreDone   rlib.JSONDate
	FLAGS       uint64
	DoneUID     int64
	PreDoneUID  int64
	Comment     string
	CreateTS    rlib.JSONDate // when was this record created
	CreateBy    int64         // employee UID (from phonebook) that created it
	LastModTime rlib.JSONDate // when was this record last written
	LastModBy   int64         // employee UID (from phonebook) that modified it
}

// SaveTaskList defines the fields supplied when Saving a TaskList
type SaveTaskList struct {
	Recid      int64 `json:"recid"`
	TLID       int64
	TLDID      int64
	BID        int64
	Name       string
	Cycle      int64
	Pivot      rlib.JSONDate
	DtDone     rlib.JSONDate
	DtDue      rlib.JSONDate
	DtPreDue   rlib.JSONDate
	DtPreDone  rlib.JSONDate
	FLAGS      int64
	DoneUID    int64
	PreDoneUID int64
	Comment    string
}

// SearchTLResponse holds the task list definition list
type SearchTLResponse struct {
	Status  string     `json:"status"`
	Total   int64      `json:"total"`
	Records []TaskList `json:"records"`
}

// TaskListInput is the input data format for a Save command
type TaskListInput struct {
	Recid    int64    `json:"recid"`
	Status   string   `json:"status"`
	FormName string   `json:"name"`
	Record   TaskList `json:"record"`
}

// GetTLResponse is the response to a GetTaskList request
type GetTLResponse struct {
	Status string   `json:"status"`
	Record TaskList `json:"record"`
}

// SaveTaskListInput is the input data format for a Save command
type SaveTaskListInput struct {
	Recid    int64        `json:"recid"`
	Status   string       `json:"status"`
	FormName string       `json:"name"`
	Record   SaveTaskList `json:"record"`
}

// which fields needs to be fetched for SQL query for assessment grid
var tlFieldsMap = map[string][]string{
	"TLID":        {"TaskList.TLID"},
	"BID":         {"TaskList.BID"},
	"Name":        {"TaskList.Name"},
	"Cycle":       {"TaskList.Cycle"},
	"DtDone":      {"TaskList.DtDone"},
	"DtDue":       {"TaskList.DtDue"},
	"DtPreDue":    {"TaskList.DtPreDue"},
	"DtPreDone":   {"TaskList.DtPreDone"},
	"FLAGS":       {"TaskList.FLAGS"},
	"DoneUID":     {"TaskList.DoneUID"},
	"PreDoneUID":  {"TaskList.PreDoneUID"},
	"Comment":     {"TaskList.Comment"},
	"CreateTS":    {"TaskList.CreateTS"},
	"CreateBy":    {"TaskList.CreateBy"},
	"LastModTime": {"TaskList.LastModTime"},
	"LastModBy":   {"TaskList.LastModBy"},
}

// which fields needs to be fetched for SQL query for assessment grid
var tlQuerySelectFields = []string{
	"TaskList.TLID",
	"TaskList.BID",
	"TaskList.Name",
	"TaskList.Cycle",
	"TaskList.DtDone",
	"TaskList.DtDue",
	"TaskList.DtPreDue",
	"TaskList.DtPreDone",
	"TaskList.FLAGS",
	"TaskList.DoneUID",
	"TaskList.PreDoneUID",
	"TaskList.Comment",
	"TaskList.CreateTS",
	"TaskList.CreateBy",
	"TaskList.LastModTime",
	"TaskList.LastModBy",
}

// TaskListRowScan scans a result from sql row and dump it in a
// TaskList struct
//
// RETURNS
//  TaskList
//-----------------------------------------------------------------------------
func TaskListRowScan(rows *sql.Rows) (TaskList, error) {
	var q TaskList
	err := rows.Scan(&q.TLID, &q.BID, &q.Name, &q.Cycle, &q.DtDone, &q.DtDue,
		&q.DtPreDue, &q.DtPreDone, &q.FLAGS, &q.DoneUID, &q.PreDoneUID, &q.Comment,
		&q.CreateTS, &q.CreateBy, &q.LastModTime, &q.LastModBy)
	return q, err
}

// SvcSearchHandlerTaskList generates a report of all TaskList defined
// business d.BID
// wsdoc {
//  @Title  Search TaskList
//	@URL /v1/tls/:BUI
//  @Method  POST
//	@Synopsis Search TaskList
//  @Description  Search all TaskList and return those that match the Search Logic.
//	@Input wsSearchReq
//  @Response SearchTLResponse
// wsdoc }
//-----------------------------------------------------------------------------
func SvcSearchHandlerTaskList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcSearchHandlerTaskList"
	var g SearchTLResponse
	var err error
	rlib.Console("Entered %s\n", funcname)

	whr := `TaskList.BID = %d AND TaskList.FLAGS & 1 = 0` // only get the Active TaskLists
	whr = fmt.Sprintf(whr, d.BID)
	order := `TaskList.Name ASC` // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, tlFieldsMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	query := `
	SELECT {{.SelectClause}}
	FROM TaskList
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(tlQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	countQuery := rlib.RenderSQLQuery(query, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
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
	queryWithLimit := query + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(queryWithLimit, qc)
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
		q, err := TaskListRowScan(rows)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		q.Recid = i

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
	SvcWriteResponse(d.BID, &g, w)
}

// SvcHandlerTaskList handles requests to read/write/update or
// make-inactive a specific TaskList.  It routes the request to
// an appropriate handler
//
// The server command can be:
//      get     - read it
//      save    - Insert or Update
//      delete  - make it inactive
//-----------------------------------------------------------------------------
func SvcHandlerTaskList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerTaskList"
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  TLID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID < 0 {
			err = fmt.Errorf("TaskListID is required but was not specified")
			SvcErrorReturn(w, err, funcname)
			return
		}
		getTaskList(w, r, d)
	case "save":
		saveTaskList(w, r, d)
	case "delete":
		deleteTaskList(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// deleteTaskList makes the secified TaskList inactive
// wsdoc {
//  @Title  Delete TaskList
//	@URL /v1/tl/:BUI/TLID
//  @Method  POST
//	@Synopsis Make a TaskList inactive
//  @Desc  This service makes a TaskList inactive. We do not deliete
//  @Desc  TaskLists
//	@Input DeletePmtForm
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func deleteTaskList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteTaskList"
	var del DeletePmtForm

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	if err := rlib.DeleteTaskList(r.Context(), d.ID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// GetTaskList returns the requested assessment
// wsdoc {
//  @Title  Save TaskList
//	@URL /v1/tl/:BUI/TLID
//  @Method  GET
//	@Synopsis Update the information on a TaskList with the supplied data
//  @Description This service updates TaskList TLID with the info
//  @Description information supplied.
//	@Input SaveTaskListInput
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func saveTaskList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveTaskList"
	var foo SaveTaskListInput
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	//---------------------------------------------------------------------
	// Create a TaskList struct based on the supplied info...
	//---------------------------------------------------------------------
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	var a rlib.TaskList
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	a.Name = foo.Record.Name
	a.BID = d.BID

	//----------------------------------------------------------------
	// Not much business logic to check here.
	// 1. Ensure that there is a name.
	//----------------------------------------------------------------
	if len(a.Name) == 0 {
		e := fmt.Errorf("%s: Required field, Name, is blank", funcname)
		SvcErrorReturn(w, e, funcname)
		return
	}

	//-------------------------------------------------------
	// Bizlogic checks done. Insert or update as needed...
	//-------------------------------------------------------
	if a.TLID == 0 && d.ID == 0 {
		if foo.Record.TLDID == 0 {
			e := fmt.Errorf("%s: Could not create TaskList because definition id (%d) does not exist", funcname, foo.Record.TLDID)
			SvcErrorReturn(w, e, funcname)
			return
		}
		tld, err := rlib.GetTaskListDefinition(r.Context(), foo.Record.TLDID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		if tld.TLDID == 0 {
			e := fmt.Errorf("%s: Could not create TaskList because definition id (%d) does not exist", funcname, foo.Record.TLDID)
			SvcErrorReturn(w, e, funcname)
			return
		}
		pivot := time.Time(foo.Record.Pivot)
		tlid, err := rlib.CreateTaskListInstance(r.Context(), foo.Record.TLDID, &pivot)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		tl, err := rlib.GetTaskList(r.Context(), tlid)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		tl.Comment = foo.Record.Comment
		err = rlib.UpdateTaskList(r.Context(), &tl)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	} else {
		err = rlib.UpdateTaskList(r.Context(), &a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving TaskList : %s (%d)", funcname, a.Name, a.TLID)
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.TLID)
}

// GetTaskList returns the requested TaskList
// wsdoc {
//  @Title  Get TaskList
//	@URL /v1/tl/:BUI/:TLID
//  @Method  GET
//	@Synopsis Get information on a TaskList
//  @Description  Return all fields for assessment :TLID
//	@Input WebGridSearchRequest
//  @Response GetTLResponse
// wsdoc }
//-----------------------------------------------------------------------------
func getTaskList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getTaskList"
	var g GetTLResponse
	var a rlib.TaskList
	var err error

	rlib.Console("entered %s\n", funcname)
	a, err = rlib.GetTaskList(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.TLID > 0 {
		var gg TaskList
		rlib.MigrateStructVals(&a, &gg)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
