package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"time"
)

//-------------------------------------------------------------------
//                        **** SEARCH ****
//-------------------------------------------------------------------

// SearchTaskDescriptor is the definition of a task. It is used to make instance
// which become Tasks
type SearchTaskDescriptor struct {
	Recid          int64 `json:"recid"`
	TDID           int64
	BID            int64
	TLDID          int64
	Name           string `json:"TDName"`
	Worker         string
	EpochDue       rlib.JSONDateTime
	EpochPreDue    rlib.JSONDateTime
	ChkEpochDue    bool
	ChkEpochPreDue bool
	FLAGS          int64
	Comment        string    `json:"TDComment"`
	LastModTime    time.Time // when was this record last written
	LastModBy      int64     // employee UID (from phonebook) that modified it
	CreateTS       time.Time // when was this record created
	CreateBy       int64     // employee UID (from phonebook) that created it
}

// SearchTDResponse holds the task list definition list
type SearchTDResponse struct {
	Status  string                 `json:"status"`
	Total   int                    `json:"total"`
	Records []SearchTaskDescriptor `json:"records"`
}

//-------------------------------------------------------------------
//                         **** SAVE ****
//-------------------------------------------------------------------

// SaveTaskDescriptor is the definition of a task to be saved.
type SaveTaskDescriptor struct {
	Recid          int64 `json:"recid"`
	TDID           int64
	BID            int64
	TLDID          int64
	Name           string `json:"TDName"`
	Worker         string
	EpochDue       rlib.JSONDateTime
	EpochPreDue    rlib.JSONDateTime
	ChkEpochDue    bool
	ChkEpochPreDue bool
	Comment        string `json:"TDComment"`
}

// SaveTaskDescriptorInput is the input data format for a Save command
type SaveTaskDescriptorInput struct {
	Recid    int64              `json:"recid"`
	Status   string             `json:"status"`
	FormName string             `json:"name"`
	Record   SaveTaskDescriptor `json:"record"`
}

//-------------------------------------------------------------------
//                         **** GET ****
//-------------------------------------------------------------------

// GetTDResponse is the response to a GetTaskDescriptor request
type GetTDResponse struct {
	Status string               `json:"status"`
	Record SearchTaskDescriptor `json:"record"`
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------

// SvcSearchTDHandler returns the TaskDescriptors associated with the supplied
// TLDID. This search handler was not implemented like many of the other
// handlers because the only use case we are supporting for TaskDescriptors
// is to search for those that belong to a particular TaskDescriptor.
// wsdoc {
//  @Title  Search TaskDescriptors
//	@URL /v1/tds/:BUI/TLDID
//  @Method  POST
//	@Synopsis Search TaskDescriptors
//  @Description  Search all TaskDescriptors associated with the supplied TLDID.
//  @Description  This call ignores any limit and simply returns all TDs.
//	@Input wsSearchReq
//  @Response SearchTDResponse
// wsdoc }
//-----------------------------------------------------------------------------
func SvcSearchTDHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcSearchTDHandler"
	rlib.Console("Entered %s.  d.ID = %d\n", funcname, d.ID)

	tds, err := rlib.GetTaskListDescriptors(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("TaskCount = %d\n", len(tds))
	var g SearchTDResponse
	for i := 0; i < len(tds); i++ {
		var t SearchTaskDescriptor
		rlib.MigrateStructVals(&tds[i], &t)
		t.Recid = int64(i)
		g.Records = append(g.Records, t)
	}
	g.Status = "success"
	g.Total = len(g.Records)
	SvcWriteResponse(d.BID, &g, w)
}

// SvcHandlerTaskDescriptor handles requests to read/write/update or
// make-inactive a specific TaskDescriptor.  It routes the request to
// an appropriate handler
//
// The server command can be:
//      get     - read it
//      save    - Insert or Update
//      delete  - make it inactive
//-----------------------------------------------------------------------------
func SvcHandlerTaskDescriptor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerTaskDescriptor"
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  TLDID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID < 0 {
			err = fmt.Errorf("TaskDescriptorID is required but was not specified")
			SvcErrorReturn(w, err, funcname)
			return
		}
		getTaskDescriptor(w, r, d)
	case "save":
		saveTaskDescriptor(w, r, d)
	case "delete":
		deleteTaskDescriptor(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// deleteTaskDescriptor makes the secified TaskDescriptor inactive
// wsdoc {
//  @Title  Delete TaskDescriptor
//	@URL /v1/deposit/:BUI/TDID
//  @Method  POST
//	@Synopsis Delete Task Descriptor TDID
//  @Desc  This service deletes the TaskDescriptor with the supplied TLDID.
//	@Input DeletePmtForm
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func deleteTaskDescriptor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteTaskDescriptor"
	var del DeletePmtForm

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	err := rlib.DeleteTaskDescriptor(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// GetTaskDescriptor returns the requested assessment
// wsdoc {
//  @Title  Save TaskDescriptor
//	@URL /v1/deposit/:BUI/TDID
//  @Method  GET
//	@Synopsis Update TaskDescriptor information
//  @Description This service updates TaskDescriptor TDID with the
//  @Description information supplied.
//	@Input SaveTaskDescriptorInput
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func saveTaskDescriptor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveTaskDescriptor"
	var foo SaveTaskDescriptorInput
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	//---------------------------------------------------------------------
	// Create a TaskDescriptor struct based on the supplied info...
	//---------------------------------------------------------------------
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	rlib.Console("*** AFTER UNMARSHAL *** foo.Record.Comment = %s\n", foo.Record.Comment)
	var a rlib.TaskDescriptor
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	a.Name = foo.Record.Name
	a.BID = d.BID

	//-------------------------------------------------------
	// Bizlogic checks...
	//-------------------------------------------------------
	e := bizlogic.ValidateTaskDescriptor(r.Context(), &a)
	if len(e) > 0 {
		SvcErrorReturn(w, bizlogic.BizErrorListToError(e), funcname)
		return
	}

	//-------------------------------------------------------
	// Insert or update as needed...
	//-------------------------------------------------------
	rlib.Console("a.TDID = %d, d.ID = %d\n", a.TDID, d.ID)
	if a.TDID == 0 && d.ID == 0 {
		rlib.Console("Inserting new TaskDescriptor\n")
		err = rlib.InsertTaskDescriptor(r.Context(), &a) // This is a new record
	} else {
		rlib.Console("Updating existing TaskDescriptor: %d\n", a.TDID) // update existing record
		err = rlib.UpdateTaskDescriptor(r.Context(), &a)
	}

	if err != nil {
		e := fmt.Errorf("%s: Error saving TaskDescriptor : %s (%d)", funcname, a.Name, a.TDID)
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.TDID)
}

// GetTaskDescriptor returns the requested TaskDescriptor
// wsdoc {
//  @Title  Get TaskDescriptor
//	@URL /v1/deposit/:BUI/TDID
//  @Method  GET
//	@Synopsis Get information on a TaskDescriptor
//  @Description  Return all fields for assessment :TDID
//	@Input WebGridSearchRequest
//  @Response GetTDResponse
// wsdoc }
//-----------------------------------------------------------------------------
func getTaskDescriptor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getTaskDescriptor"
	var g GetTDResponse
	var a rlib.TaskDescriptor
	var err error

	rlib.Console("entered %s\n", funcname)
	a, err = rlib.GetTaskDescriptor(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.TDID > 0 {
		var gg SearchTaskDescriptor
		rlib.MigrateStructVals(&a, &gg)
		gg.Recid = gg.TDID
		gg.ChkEpochDue = a.EpochDue.Year() > 1970
		gg.ChkEpochPreDue = a.EpochPreDue.Year() > 1970
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
