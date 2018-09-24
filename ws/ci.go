package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// CloseInfo is the service version of rlib.CloseInfo
type CloseInfo struct {
	BID       int64             // Business ID
	LastClose rlib.JSONDateTime // last closed period
	CPID      int64             // id of last close
	BKDTRACP  bool              // backdate rental agreements in closed period allowed?
}

// SvcGetCloseInfo attempts to save the period. All checks must pass.
// wsdoc {
//  @Title  GetCloseInfo
//	@URL /v1/closeinfo/:BUI
//  @Method  GET
//	@Synopsis Close Period information for the supplied business
//  @Description Gets the date of the last closed period, and the
//  @Description business flag that indicates whether or not a Rental Agreement
//  @Description can be backdated into a closed period.
//	@Input FormClosePeriod
//  @Response SvcStatusResponse
// wsdoc }
//-----------------------------------------------------------------------------
func SvcGetCloseInfo(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcGetCloseInfo"
	ci, err := rlib.GetCloseInfo(r.Context(), d.BID)
	if err != nil {
		e := fmt.Errorf("%s: Error with GetCloseInfo BID=%d: %s", funcname, d.BID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	var wsci CloseInfo
	rlib.MigrateStructVals(&ci, &wsci)
	SvcWriteResponse(d.BID, &wsci, w)
}
