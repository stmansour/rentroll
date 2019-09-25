package ws

import (
	"net/http"
	"rentroll/rlib"
)

// SvcCheckIn performs the tasks needed to check in a user based on the
// supplied RLID.
//
// wsdoc {
//  @Title  CheckIn
//	@URL /v1/checkin/:BUI/[RLID]
//  @Method  POST
//	@Synopsis Performs check in functions
//  @Description  Ensures that there is a rent assessment and makes it active.
//	@Input WebGridSearchRequest
//  @Response Reservation
// wsdoc }
//------------------------------------------------------------------------------
func SvcCheckIn(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcCheckIn"
	// var res ResDet
	var err error
	var a rlib.RentableLeaseStatus
	rlib.Console("entered %s, getting BID = %d, RLID = %d\n", funcname, d.BID, d.ID)

	if a, err = rlib.GetRentableLeaseStatus(r.Context(), d.ID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Successfully read RLID %d\n", a.RLID)

	SvcWriteSuccessResponse(d.BID, w)
}
