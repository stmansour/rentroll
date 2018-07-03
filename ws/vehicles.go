package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
)

// This command returns vehicles associated with a Rental Agreement.
// Current date is assumed unless a date is provided to override.

// VehiclesResp is the struct containing the JSON return values for this web service
type VehiclesResp struct {
	Status  string         `json:"status"`
	Total   int64          `json:"total"`
	Records []rlib.Vehicle `json:"records"`
}

// SvcUserVehicles is used to get the vehicles associated with the
// TCID supplied.
//
// wsdoc {
//  @Title  People Vehicles
//  @URL /v1/uservehicles/:BUI/:TCID
//  @Method  GET
//  @Synopsis Get the vehicles associated with a TCID (people)
//  @Description  Returns all the vehicles for the supplied TCID
//  @Input
//  @Response VehiclesResp
// wsdoc }
// URL:
//       0    1           2   3
//      /v1/uservehicles/BID/TCID
//-----------------------------------------------------------------------------
func SvcUserVehicles(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcUserVehicles"
	var (
		err error
	)

	fmt.Printf("entered %s\n", funcname)
	s := r.URL.String()                 // ex: /v1/uservehicles/CCC/10
	s1 := strings.Split(s, "?")         // ex: /v1/uservehicles/CCC/10
	ss := strings.Split(s1[0][1:], "/") // ex: []string{"v1", "uservehicles", "CCC", "10"}
	fmt.Printf("ss = %#v\n", ss)

	//------------------------------------------------------
	// Handle URL path values
	//------------------------------------------------------
	var TCID int64
	TCID, err = rlib.IntFromString(ss[3], "bad TCID value")
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//------------------------------------------------------
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	var resp VehiclesResp
	var m []rlib.Vehicle
	if TCID > 0 {
		m, err = rlib.GetVehiclesByTransactant(r.Context(), TCID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	}

	resp.Records = m

	//------------------------------------------------------
	// marshal resp and send it!
	//------------------------------------------------------
	resp.Status = "success"
	resp.Total = int64(len(resp.Records))
	fmt.Printf("resp = %#v\n", resp)

	b, err := json.Marshal(&resp)
	if err != nil {
		err = fmt.Errorf("cannot marshal resp:  %s", err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("b = %s\n", string(b))
	SvcWriteResponse(d.BID, &resp, w)
}

// VehicleFeesResp is the response struct containing all vehicle fees
type VehicleFeesResp struct {
	Status  string                    `json:"status"`
	Total   int64                     `json:"total"`
	Records []rlib.BizPropsVehicleFee `json:"records"`
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
		err error
		g   VehicleFeesResp
	)
	fmt.Printf("Entered in %s\n", funcname)

	g.Records, err = rlib.GetVehicleFeesFromGeneralBizProps(r.Context(), d.BID)
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
