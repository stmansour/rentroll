package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
)

// CheckOutData is the payload data for the CheckOut Command.
type CheckOutData struct {
	TZOffset int
}

// SvcCheckOut performs the tasks needed to check in a user based on the
// supplied RLID.
//
// wsdoc {
//  @Title  CheckOut
//	@URL /v1/CheckOut/:BUI/RLID
//  @Method  POST
//	@Synopsis Performs check in functions
//  @Description  Performs the following tasks:
//  @Description  *  Ensures that if we check out prior to the stop date on the
//  @Description     RentalAgreement that we stop the rent assessment and we
//  @Description     terminate the rental agreement on the current date.
//	@Input WebGridSearchRequest
//  @Response Reservation
// wsdoc }
//------------------------------------------------------------------------------
func SvcCheckOut(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcCheckOut"
	var err error

	rlib.Console("entered %s, getting BID = %d, RLID = %d\n", funcname, d.BID, d.ID)

	target := `"record":`
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("%s: cannot find %s in form json", funcname, target)
		SvcErrorReturn(w, e, funcname)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	//---------------------------------------------------
	// Read the payload data from the client
	//---------------------------------------------------
	var chkout CheckOutData
	err = json.Unmarshal([]byte(s), &chkout)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	rlib.Console("Successfully read payload:  TZOffset = %d\n", chkout.TZOffset)

	SvcWriteSuccessResponse(d.BID, w)
}
