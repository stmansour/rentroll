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

//-------------------------------------------------------------------
//
//                        **** SEARCH ****
//
//-------------------------------------------------------------------

// TaskListDefs is the structure describing a task list definition
type TaskListDefs struct {
	Recid          int64 `json:"recid"`
	TLDID          int64
	BID            int64
	Name           string
	Cycle          int64
	Epoch          rlib.JSONDateTime
	EpochDue       rlib.JSONDateTime
	EpochPreDue    rlib.JSONDateTime
	ChkEpochDue    bool
	ChkEpochPreDue bool
	FLAGS          int64
	Comment        string
	CreateTS       rlib.JSONDate
	CreateBy       int64
	LastModTime    rlib.JSONDate
	LastModBy      int64
}

// SearchTLDResponse holds the task list definition list
type SearchTLDResponse struct {
	Status  string         `json:"status"`
	Total   int64          `json:"total"`
	Records []TaskListDefs `json:"records"`
}

// which fields needs to be fetched for SQL query for assessment grid
var tldFieldsMap = map[string][]string{
	"TLDID":       {"TaskListDefinition.TLDID"},
	"BID":         {"TaskListDefinition.BID"},
	"Name":        {"TaskListDefinition.Name"},
	"Cycle":       {"TaskListDefinition.Cycle"},
	"Epoch":       {"TaskListDefinition.Epoch"},
	"EpochDue":    {"TaskListDefinition.EpochDue"},
	"EpochPreDue": {"TaskListDefinition.EpochPreDue"},
	"FLAGS":       {"TaskListDefinition.FLAGS"},
	"CreateTS":    {"TaskListDefinition.CreateTS"},
	"CreateBy":    {"TaskListDefinition.CreateBy"},
	"LastModTime": {"TaskListDefinition.LastModTime"},
	"LastModBy":   {"TaskListDefinition.LastModBy"},
}

// which fields needs to be fetched for SQL query for assessment grid
var tldQuerySelectFields = []string{
	"TaskListDefinition.TLDID",
	"TaskListDefinition.BID",
	"TaskListDefinition.Name",
	"TaskListDefinition.Cycle",
	"TaskListDefinition.Epoch",
	"TaskListDefinition.EpochDue",
	"TaskListDefinition.EpochPreDue",
	"TaskListDefinition.FLAGS",
	"TaskListDefinition.CreateTS",
	"TaskListDefinition.CreateBy",
	"TaskListDefinition.LastModTime",
	"TaskListDefinition.LastModBy",
}

//-------------------------------------------------------------------
//
//                         **** SAVE ****
//
//-------------------------------------------------------------------

// SaveTaskListDef defines the fields supplied when Saving a TaskListDefinition
type SaveTaskListDef struct {
	Recid          int64 `json:"recid"`
	TLDID          int64
	BID            int64
	Name           string
	Cycle          int64
	Epoch          rlib.JSONDateTime
	EpochDue       rlib.JSONDateTime
	EpochPreDue    rlib.JSONDateTime
	ChkEpochDue    bool
	ChkEpochPreDue bool
	FLAGS          int64
	Comment        string
}

// SaveTaskListDefinitionInput is the input data format for a Save command
type SaveTaskListDefinitionInput struct {
	Recid    int64           `json:"recid"`
	Status   string          `json:"status"`
	FormName string          `json:"name"`
	Record   SaveTaskListDef `json:"record"`
}

//-------------------------------------------------------------------
//
//                         **** GET ****
//
//-------------------------------------------------------------------

// GetTLDResponse is the response to a GetTaskListDefinition request
type GetTLDResponse struct {
	Status string       `json:"status"`
	Record TaskListDefs `json:"record"`
}

// ############################################################################

// TaskListDefsRowScan scans a result from sql row and dump it in a
// TaskListDefs struct
//
// RETURNS
//  TaskListDefs
//-----------------------------------------------------------------------------
func TaskListDefsRowScan(rows *sql.Rows) (TaskListDefs, error) {
	var q TaskListDefs
	err := rows.Scan(&q.TLDID, &q.BID, &q.Name, &q.Cycle, &q.Epoch, &q.EpochDue,
		&q.EpochPreDue, &q.FLAGS, &q.CreateTS, &q.CreateBy, &q.LastModTime,
		&q.LastModBy)
	return q, err
}

// SvcSearchHandlerTaskListDefs generates a report of all TaskListDefs defined
// business d.BID
// wsdoc {
//  @Title  Search TaskListDefs
//	@URL /v1/tlds/:BUI
//  @Method  POST
//	@Synopsis Search TaskListDefs
//  @Description  Search all TaskListDefs and return those that match the Search Logic.
//	@Input wsSearchReq
//  @Response SearchTLDResponse
// wsdoc }
//-----------------------------------------------------------------------------
func SvcSearchHandlerTaskListDefs(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcSearchHandlerTaskListDefs"
	var g SearchTLDResponse
	var err error
	rlib.Console("Entered %s\n", funcname)

	whr := `TaskListDefinition.BID = %d AND TaskListDefinition.FLAGS & 1 = 0` // only get the Active tasklistdefinitions
	whr = fmt.Sprintf(whr, d.BID)
	order := `TaskListDefinition.Name ASC` // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, tldFieldsMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	query := `
	SELECT {{.SelectClause}}
	FROM TaskListDefinition
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(tldQuerySelectFields, ","),
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
		q, err := TaskListDefsRowScan(rows)
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

// SvcHandlerTaskListDefinition handles requests to read/write/update or
// make-inactive a specific TaskListDefinition.  It routes the request to
// an appropriate handler
//
// The server command can be:
//      get     - read it
//      save    - Insert or Update
//      delete  - make it inactive
//-----------------------------------------------------------------------------
func SvcHandlerTaskListDefinition(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerTaskListDefinition"
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  TLDID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID < 0 {
			err = fmt.Errorf("TaskListDefinitionID is required but was not specified")
			SvcErrorReturn(w, err, funcname)
			return
		}
		getTaskListDefinition(w, r, d)
	case "save":
		saveTaskListDefinition(w, r, d)
	case "delete":
		deleteTaskListDefinition(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// deleteTaskListDefinition makes the secified TaskListDefinition inactive
// wsdoc {
//  @Title  Delete TaskListDefinition
//	@URL /v1/tld/:BUI/TLDID
//  @Method  POST
//	@Synopsis Make a TaskListDefinition inactive
//  @Desc  This service makes a TaskListDefinition inactive. We do not deliete
//  @Desc  TaskListDefinitions
//	@Input DeletePmtForm
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func deleteTaskListDefinition(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteTaskListDefinition"
	var del DeletePmtForm

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	tld, err := rlib.GetTaskListDefinition(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	tld.FLAGS |= 0x1 // bit 0 set means it is inactive
	err = rlib.UpdateTaskListDefinition(r.Context(), &tld)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// GetTaskListDefinition returns the requested assessment
// wsdoc {
//  @Title  Save TaskListDefinition
//	@URL /v1/tld/:BUI/TLDID
//  @Method  GET
//	@Synopsis Update the information on a TaskListDefinition with the supplied data
//  @Description This service updates TaskListDefinition :TLDID with the
//  @Description information supplied.
//	@Input SaveTaskListDefinitionInput
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func saveTaskListDefinition(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveTaskListDefinition"
	var foo SaveTaskListDefinitionInput
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	//---------------------------------------------------------------------
	// Create a TaskListDefinition struct based on the supplied info...
	//---------------------------------------------------------------------
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	var a rlib.TaskListDefinition
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	a.Name = foo.Record.Name
	a.BID = d.BID

	//----------------------------------------------------------------
	// Not much business logic to check here.
	// 1. Ensure that there is a name.
	// 2. If it is an insert, make sure there's no duplicate name
	//----------------------------------------------------------------
	if len(a.Name) == 0 {
		e := fmt.Errorf("%s: Required field, Name, is blank", funcname)
		SvcErrorReturn(w, e, funcname)
		return
	}
	var adup rlib.TaskListDefinition
	adup, err = rlib.GetTaskListDefinitionByName(r.Context(), d.BID, a.Name)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.Name == adup.Name && a.TLDID != adup.TLDID {
		e := fmt.Errorf("%s: A TaskListDefinition with the name %s already exists", funcname, a.Name)
		SvcErrorReturn(w, e, funcname)
		return
	}

	//-------------------------------------------------------
	// Bizlogic checks done. Insert or update as needed...
	//-------------------------------------------------------
	if a.TLDID == 0 && d.ID == 0 {
		// rlib.Console("Inserting new TaskListDefinition\n")
		err = rlib.InsertTaskListDefinition(r.Context(), &a) // This is a new record
	} else {
		// rlib.Console("Updating existing TaskListDefinition: %d\n", a.TLDID) // update existing record
		err = rlib.UpdateTaskListDefinition(r.Context(), &a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving TaskListDefinition : %s (%d)", funcname, a.Name, a.TLDID)
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.TLDID)
}

// GetTaskListDefinition returns the requested TaskListDefinition
// wsdoc {
//  @Title  Get TaskListDefinition
//	@URL /v1/tld/:BUI/:TLDID
//  @Method  GET
//	@Synopsis Get information on a TaskListDefinition
//  @Description  Return all fields for assessment :TLDID
//	@Input WebGridSearchRequest
//  @Response GetTLDResponse
// wsdoc }
//-----------------------------------------------------------------------------
func getTaskListDefinition(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getTaskListDefinition"
	var g GetTLDResponse
	var a rlib.TaskListDefinition
	var err error

	rlib.Console("entered %s\n", funcname)
	a, err = rlib.GetTaskListDefinition(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.TLDID > 0 {
		var gg TaskListDefs
		rlib.MigrateStructVals(&a, &gg)
		gg.ChkEpochPreDue = a.EpochPreDue.Year() > 1900
		gg.ChkEpochDue = a.EpochDue.Year() > 1900
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
