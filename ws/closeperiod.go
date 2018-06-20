package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

//-------------------------------------------------------------------
//                        **** SEARCH ****
//-------------------------------------------------------------------

// no search area here because there is no main grid

// FormClosePeriod holds the data needed for the Close Period screen
type FormClosePeriod struct {
	BID              int64
	TLID             int64             // the TaskList to which this task belongs
	TLName           string            // Name of the tasklist
	LastDtDone       rlib.JSONDateTime // Date of last completed TaskList instance
	LastDtClose      rlib.JSONDateTime // Datetime of last close
	LastLedgerMarker rlib.JSONDateTime // date/time of last LedgerMarker
	CloseTarget      rlib.JSONDateTime // due date of first period that has not been closed
	TLIDTarget       int64             // need to complete this tasklist in order to close target period
	TLNameTarget     string            // name associated with TLIDTarget
	DtDueTarget      rlib.JSONDateTime // due date/time of TLIDTarget
	DtDoneTarget     rlib.JSONDateTime // done date of TLIDTarget
	DtDone           rlib.JSONDateTime // done date of first period that has not been closed
}

//-------------------------------------------------------------------
//                         **** SAVE ****
//-------------------------------------------------------------------

// The "button" in the UI in this case is
// pressed to close a period

// SaveClosePeriod is the response to a GetTask request
type SaveClosePeriod struct {
	Cmd    string          `json:"cmd"`
	Record FormClosePeriod `json:"record"`
}

//-------------------------------------------------------------------
//                         **** GET ****
//-------------------------------------------------------------------

// GetClosePeriodResponse is the response to a GetClosePeriod request
type GetClosePeriodResponse struct {
	Status string          `json:"status"`
	Record FormClosePeriod `json:"record"`
}

//-----------------------------------------------------------------------------
//#############################################################################
//-----------------------------------------------------------------------------

// SvcHandlerClosePeriod handles requests for closing a period
//
// The server command can be:
//      get     - read it
//      save    - Close the period (oldest unclosed period)
//      delete  - Reopen period
//-----------------------------------------------------------------------------
func SvcHandlerClosePeriod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerClosePeriod"

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  d.ID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getClosePeriod(w, r, d)
	case "save":
		saveClosePeriod(w, r, d)
	case "delete":
		deleteClosePeriod(w, r, d)
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// SaveClosePeriod attempts to save the period. All checks must pass.
// wsdoc {
//  @Title  Save ClosePeriod
//	@URL /v1/closeperiod/:BUI/TID
//  @Method  GET
//	@Synopsis Update ClosePeriod information
//  @Description This service attempts to close the oldest unclosed period
//  @Description after performing a myriad of tests
//	@Input FormClosePeriod
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func saveClosePeriod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveClosePeriod"
	var xbiz rlib.XBusiness
	var foo SaveClosePeriod
	data := []byte(d.data)

	err := json.Unmarshal(data, &foo)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//------------------------------------
	//  Get business info...
	//------------------------------------
	if err = rlib.GetXBusiness(r.Context(), d.BID, &xbiz); err != nil {
		rlib.Console("B\n")
		e := fmt.Errorf("%s: Error getting business %d: %s", funcname, d.BID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	//-------------------------------------------------
	// Get the close date from the
	// tasklist associated with this close...
	//-------------------------------------------------
	rlib.Console("TaskList: TLID = %d\n", foo.Record.TLIDTarget)
	tl, err := rlib.GetTaskList(r.Context(), foo.Record.TLIDTarget)
	if err != nil {
		e := fmt.Errorf("%s: Error getting TaskList %d: %s", funcname, foo.Record.TLIDTarget, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	if err = rlib.InitBizInternals(d.BID, &xbiz); err != nil {
		e := fmt.Errorf("%s: Error InitBizInternals BID = %d: %s", funcname, d.BID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	//-------------------------------------------------------------
	// TRANSACTION:
	//    1 - write the ClosePeriod entry
	//    2 - Write the LedgerMarkers on the tasklist due date...
	//-------------------------------------------------------------
	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//-------------------------------------------------------------
	//  Close Period...
	//-------------------------------------------------------------
	var cp rlib.ClosePeriod
	cp.TLID = foo.Record.TLIDTarget
	cp.BID = d.BID
	cp.Dt = tl.DtDue
	_, err = rlib.InsertClosePeriod(ctx, &cp)
	if err != nil {
		tx.Rollback()
		e := fmt.Errorf("%s: Error writing ClosePeriod: %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	//-------------------------------------------------------------
	//  Generate Ledger Markers...
	//-------------------------------------------------------------
	if err = rlib.GenerateLedgerMarkers(ctx, &xbiz, &tl.DtDue); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	//-------------------------------------------------------------
	//  Generate RAID Ledger Markers...
	//-------------------------------------------------------------

	//-------------------------------------------------------------
	// COMMIT TRANSACTION
	//-------------------------------------------------------------
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// GetClosePeriod returns the requested ClosePeriod
// wsdoc {
//  @Title  GetClosePeriod
//	@URL /v1/closeperiod/:BUI/TID
//  @Method  GET
//	@Synopsis Get information on a ClosePeriod
//  @Description  Returns information about the business CloseTaskList, the
//  @Description  last closed period, the current close period
//	@Input WebGridSearchRequest
//  @Response GetClosePeriodResponse
// wsdoc }
//-----------------------------------------------------------------------------
func getClosePeriod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getClosePeriod"
	var g GetClosePeriodResponse
	var xbiz rlib.XBusiness
	var err error
	var tl rlib.TaskList
	var tlNext rlib.TaskList
	var lcp rlib.ClosePeriod

	rlib.Console("entered %s, getting BID = %d\n", funcname, d.BID)
	rlib.Console("A\n")

	//------------------------------------
	//  Get business info...
	//------------------------------------
	if err = rlib.GetXBusiness(r.Context(), d.BID, &xbiz); err != nil {
		rlib.Console("B\n")
		e := fmt.Errorf("%s: Error getting business %d: %s", funcname, d.BID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	rlib.Console("C\n")
	g.Record.BID = d.BID
	g.Record.TLID = xbiz.P.ClosePeriodTLID
	if 0 == xbiz.P.ClosePeriodTLID {
		goto EXITNOW
	}

	//-----------------------------------
	// Get the last period closed...
	//-----------------------------------
	lcp, err = rlib.GetLastClosePeriod(r.Context(), d.BID)
	if err != nil {
		e := fmt.Errorf("%s: Error getting LastClosePeriod: %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	if lcp.CPID > 0 {
		//----------------------------------------------
		//  Get the TaskList for the last closed period
		//----------------------------------------------
		g.Record.LastDtClose = rlib.JSONDateTime(lcp.Dt)
		rlib.Console("C1 - got LastClosePeriod:  CPID = %d\n", lcp.CPID)
		tl, err = rlib.GetTaskList(r.Context(), lcp.TLID)
		if err != nil {
			rlib.Console("D\n")
			e := fmt.Errorf("%s: Error getting close period tasklist %d: %s", funcname, lcp.TLID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		rlib.Console("E - last completed tasklist: TLID = %d\n", tl.TLID)
		g.Record.TLName = tl.Name
		g.Record.LastDtDone = rlib.JSONDateTime(tl.DtDone)
	} else {
		//----------------------------------------------------------------
		// no entries yet.  So our target date needs to be the due date
		// of the task list associated with the business.
		//----------------------------------------------------------------
		tlNext, err = rlib.GetTaskList(r.Context(), g.Record.TLID)
		if err != nil {
			rlib.Console("F\n")
			e := fmt.Errorf("%s: Error getting close period tasklist %d: %s", funcname, g.Record.TLID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		g.Record.TLName = tlNext.Name
		g.Record.CloseTarget = rlib.JSONDateTime(tlNext.DtDue)
	}

	//-------------------------------------------------------------------------
	// Calculate next close target.  If we were able to find tl above it means
	// that we have closed at least one period. If not, it means no periods
	// have ever been closed.
	//-------------------------------------------------------------------------
	if tl.TLID > 0 && lcp.CPID > 0 && xbiz.P.ClosePeriodTLID > 0 {
		//---------------------------------------------------------------------
		// We have already closed a period.  Just figure out the next instance.
		//---------------------------------------------------------------------
		target := rlib.NextInstance(&tl.DtDue, tl.Cycle)
		g.Record.CloseTarget = rlib.JSONDateTime(target)
		dt1 := target.Add(-1 * time.Hour)
		dt2 := target.Add(1 * time.Hour)
		id := tl.PTLID // assume it's an instance and set the id to the parent
		if id == 0 {   // check to see if this is actually the parent instance
			id = tl.TLID // if it's the parent, then just use TLID
		}
		// rlib.Console("\n\nSEARCHING FOR INSTANCE:  PTLID = %d, dt1 = %s, dt2 = %s\n", id, dt1.Format(rlib.RRDATETIMERPTFMT), dt2.Format(rlib.RRDATETIMERPTFMT))
		tlNext, err = rlib.GetTaskListInstanceInRange(r.Context(), id, &dt1, &dt2)
		if err != nil {
			e := fmt.Errorf("%s: Error getting next TaskList instance: %s", funcname, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	//--------------------------------------------------------
	// Return basic info on next instance if we have it...
	//--------------------------------------------------------
	if tlNext.TLID > 0 { // did we find the next instance?
		g.Record.TLIDTarget = tlNext.TLID
		g.Record.TLNameTarget = tlNext.Name
		g.Record.DtDueTarget = rlib.JSONDateTime(tlNext.DtDue)
		g.Record.DtDoneTarget = rlib.JSONDateTime(tlNext.DtDone)
	}

EXITNOW:
	SvcWriteResponse(d.BID, &g, w)
}

// deleteClosePeriod reopens the ClosePeriod specified
// wsdoc {
//  @Title  Delete ClosePeriod
//	@URL /v1/closeperiod/:BUI/TID
//  @Method  POST
//	@Synopsis Reopen ClosePeriod with the supplied date
//  @Desc  This service deletes the ClosePeriod with the supplied TDID.
//	@Input DeletePmtForm
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func deleteClosePeriod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	// const funcname = "deleteClosePeriod"
}
