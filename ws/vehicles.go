package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// VehicleFeesResp is the response struct containing all vehicle fees
type VehicleFeesResp struct {
	Status  string             `json:"status"`
	Total   int64              `json:"total"`
	Records []rlib.BizPropsFee `json:"records"`
}

// SvcVehicleFeesHandler is used to get the vehicle fees associated with the BID
//
// wsdoc {
//  @Title  Vehicle Fees
//  @URL /v1/vehiclefees/:BID
//  @Method  GET
//  @Synopsis Get the vehicle fees associated with a BID
//  @Description  Returns all the vehicle fees for a BID
//  @Input
//  @Response VehicleFeesResp
// wsdoc }
// URL:
//       0    1       2   3
//      /v1/uservehicles/BID/TCID
//-----------------------------------------------------------------------------
func SvcVehicleFeesHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcVehicleFeesHandler"
	var (
		err         error
		g           VehicleFeesResp
		bizPropName = "general"
	)
	fmt.Printf("Entered in %s\n", funcname)

	g.Records, err = rlib.GetBizPropVehicleFees(r.Context(), d.BID, bizPropName)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// success mark
	g.Status = "success"
	g.Total = int64(len(g.Records))

	// success response
	SvcWriteResponse(d.BID, &g, w)
}
