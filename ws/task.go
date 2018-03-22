package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// SearchTask is the definition of a task. It is used to make instance
// which become Tasks
type SearchTask struct {
	Recid       int64 `json:"recid"`
	TID         int64
	BID         int64
	TLID        int64     // the TaskList to which this task belongs
	Name        string    // Task text
	Worker      string    // Name of the associated work function
	DtDue       time.Time // Task Due Date
	DtPreDue    time.Time // Pre Completion due date
	DtDone      time.Time // Task completion Date
	DtPreDone   time.Time // Task Pre Completion Date
	FLAGS       int64
	DoneUID     int64     // user who marked task as done
	PreDoneUID  int64     // user who marked task as predone
	Comment     string    // any user comments
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// SearchTaskResponse holds the task list definition list
type SearchTaskResponse struct {
	Status  string       `json:"status"`
	Total   int          `json:"total"`
	Records []SearchTask `json:"records"`
}

// SaveTaskInput is the input data format for a Save command
type SaveTaskInput struct {
	Recid    int64      `json:"recid"`
	Status   string     `json:"status"`
	FormName string     `json:"name"`
	Record   SearchTask `json:"record"`
}

// GetTaskResponse is the response to a GetTask request
type GetTaskResponse struct {
	Status string     `json:"status"`
	Record SearchTask `json:"record"`
}

// SvcSearchTaskHandler returns the Tasks associated with the supplied
// TLID. This search handler was not implemented like many of the other
// handlers because the only use case we are supporting for Tasks
// is to search for those that belong to a particular Task.
// wsdoc {
//  @Title  Search Tasks
//	@URL /v1/tasks/:BUI/TLID
//  @Method  POST
//	@Synopsis Search Tasks
//  @Description  Search all Tasks associated with the supplied TDID.
//  @Description  This call ignores any limit and simply returns all TDs.
//	@Input wsSearchReq
//  @Response SearchTaskResponse
// wsdoc }
//-----------------------------------------------------------------------------
func SvcSearchTaskHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcSearchTaskHandler"
	rlib.Console("Entered %s.  d.ID = %d\n", funcname, d.ID)

	tds, err := rlib.GetTasks(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("TaskCount = %d\n", len(tds))
	var g SearchTaskResponse
	for i := 0; i < len(tds); i++ {
		var t SearchTask
		rlib.MigrateStructVals(&tds[i], &t)
		t.Recid = int64(i)
		g.Records = append(g.Records, t)
	}
	g.Status = "success"
	g.Total = len(g.Records)
	SvcWriteResponse(d.BID, &g, w)
}

// SvcHandlerTask handles requests to read/write/update or
// make-inactive a specific Task.  It routes the request to
// an appropriate handler
//
// The server command can be:
//      get     - read it
//      save    - Insert or Update
//      delete  - make it inactive
//-----------------------------------------------------------------------------
func SvcHandlerTask(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerTask"
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  TDID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID < 0 {
			err = fmt.Errorf("TaskID is required but was not specified")
			SvcErrorReturn(w, err, funcname)
			return
		}
		getTask(w, r, d)
	case "save":
		saveTask(w, r, d)
	case "delete":
		deleteTask(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// deleteTask makes the secified Task inactive
// wsdoc {
//  @Title  Delete Task
//	@URL /v1/task/:BUI/TID
//  @Method  POST
//	@Synopsis Delete Task Descriptor TID
//  @Desc  This service deletes the Task with the supplied TDID.
//	@Input DeletePmtForm
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func deleteTask(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteTask"
	var del DeletePmtForm

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	err := rlib.DeleteTask(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// GetTask returns the requested assessment
// wsdoc {
//  @Title  Save Task
//	@URL /v1/task/:BUI/TID
//  @Method  GET
//	@Synopsis Update Task information
//  @Description This service updates Task TID with the
//  @Description information supplied.
//	@Input SaveTaskInput
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func saveTask(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveTask"
	var foo SaveTaskInput
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	//---------------------------------------------------------------------
	// Create a Task struct based on the supplied info...
	//---------------------------------------------------------------------
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	var a rlib.Task
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	a.Name = foo.Record.Name
	a.BID = d.BID

	//-------------------------------------------------------
	// Bizlogic checks...
	//-------------------------------------------------------
	// e := bizlogic.ValidateTask(r.Context(), &a)
	// if len(e) > 0 {
	// 	SvcErrorReturn(w, bizlogic.BizErrorListToError(e), funcname)
	// 	return
	// }

	//-------------------------------------------------------
	// Insert or update as needed...
	//-------------------------------------------------------
	rlib.Console("a.TID = %d, d.ID = %d\n", a.TID, d.ID)
	if a.TID == 0 && d.ID == 0 {
		rlib.Console("Inserting new Task\n")
		err = rlib.InsertTask(r.Context(), &a) // This is a new record
	} else {
		rlib.Console("Updating existing Task: %d\n", a.TID) // update existing record
		err = rlib.UpdateTask(r.Context(), &a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving Task %s (%d): %s", funcname, a.Name, a.TID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.TID)
}

// GetTask returns the requested Task
// wsdoc {
//  @Title  Get Task
//	@URL /v1/task/:BUI/TID
//  @Method  GET
//	@Synopsis Get information on a Task
//  @Description  Return all fields for assessment :TID
//	@Input WebGridSearchRequest
//  @Response GetTaskResponse
// wsdoc }
//-----------------------------------------------------------------------------
func getTask(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getTask"
	var g GetTaskResponse
	var a rlib.Task
	var err error

	rlib.Console("entered %s, getting TID = %d\n", funcname, d.ID)
	a, err = rlib.GetTask(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.TID > 0 {
		var gg SearchTask
		rlib.MigrateStructVals(&a, &gg)
		gg.Recid = gg.TID
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
