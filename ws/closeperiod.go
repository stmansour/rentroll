package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
)

//-------------------------------------------------------------------
//                        **** SEARCH ****
//-------------------------------------------------------------------

// no search area here because there is no main grid

//-------------------------------------------------------------------
//                         **** SAVE ****
//-------------------------------------------------------------------

// The "button" in the UI in this case is
// pressed to close a period

// SaveClosePeriod is the response to a GetTask request
type SaveClosePeriod struct {
	Cmd    string   `json:"status"`
	Record FormTask `json:"record"`
}

//-------------------------------------------------------------------
//                         **** GET ****
//-------------------------------------------------------------------

// FormClosePeriod holds the data needed for the Close Period screen
type FormClosePeriod struct {
	BID              int64
	TLID             int64             // the TaskList to which this task belongs
	TLName           string            // Name of the tasklist
	LastDtDone       rlib.JSONDateTime // Date of last completed TaskList instance
	LastDtClose      rlib.JSONDateTime // Datetime of last close
	LastLedgerMarker rlib.JSONDateTime // date/time of last LedgerMarker
	CloseTarget      rlib.JSONDateTime // due date of first period that has not been closed
	DtDone           rlib.JSONDateTime // done date of first period that has not been closed
}

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
	rlib.Console("Request: %s:  BID = %d,  TDID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

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
//	@Input ClosePeriod
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func saveClosePeriod(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	// funcname := "saveClosePeriod"
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

	//------------------------------------
	//  Get the close period TaskList...
	//------------------------------------
	if xbiz.P.ClosePeriodTLID > 0 {
		tl, err = rlib.GetLatestCompletedTaskList(r.Context(), xbiz.P.ClosePeriodTLID)
		if err != nil {
			rlib.Console("D\n")
			e := fmt.Errorf("%s: Error getting close period tasklist %d: %s", funcname, xbiz.P.ClosePeriodTLID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		rlib.Console("E - last completed tasklist: TLID = %d\n", tl.TLID)
		g.Record.TLName = tl.Name
		g.Record.LastDtDone = rlib.JSONDateTime(tl.DtDone)

	}

	//-----------------------------------
	// Get the last period closed...
	//-----------------------------------
	lcp, err := rlib.GetLastClosePeriod(r.Context(), d.BID)
	if err != nil {
		e := fmt.Errorf("%s: Error getting LastClosePeriod: %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	if lcp.CPID > 0 {
		g.Record.LastDtClose = rlib.JSONDateTime(lcp.Dt)
		rlib.Console("F - got LastClosePeriod:  CPID = %d\n", lcp.CPID)
	}

	//-----------------------------------
	// Calculate next close target...
	//-----------------------------------
	target := rlib.NextInstance(&tl.DtDue, tl.Cycle)
	g.Record.CloseTarget = rlib.JSONDateTime(target)

	//-------------------------------------------
	// Find the TaskList for this target...
	//-------------------------------------------

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
