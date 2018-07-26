package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
)

// This command returns pets associated with a Rental Agreement.
// Current date is assumed unless a date is provided to override.

// RAPets is the struct containing the JSON return values for this web service
type RAPets struct {
	Status  string                    `json:"status"`
	Total   int64                     `json:"total"`
	Records []rlib.RentalAgreementPet `json:"records"`
}

// SvcRAPets is used to get the pets associated with the
// RAID supplied.
//
// wsdoc {
//  @Title  Rental Agreement Pets
//	@URL /v1/rapets/:BUI/:RAID ? dt=:DATE
//  @Method  GET
//	@Synopsis Get the pets associated with a Rental Agreement
//  @Description  Returns all the pets for the supplied Rental Agreement as of :DATE
//	@Input
//  @Response RAPets
// wsdoc }
// URL:
//       0    1       2    3
// 		/v1/rapets/BID/RAID?dt=2017-02-01
//-----------------------------------------------------------------------------
func SvcRAPets(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRAPets"
	var (
		err error
	)

	fmt.Printf("entered %s\n", funcname)
	s := r.URL.String()                 // ex: /v1/rapets/CCC/10?dt=2017-02-01
	fmt.Printf("s = %s\n", s)           // x
	s1 := strings.Split(s, "?")         // ex: /v1/rapets/CCC/10?dt=2017-02-01
	fmt.Printf("s1 = %#v\n", s1)        // x
	ss := strings.Split(s1[0][1:], "/") // ex: []string{"v1", "rapets", "CCC", "10"}
	fmt.Printf("ss = %#v\n", ss)

	//------------------------------------------------------
	// Handle URL path values
	//------------------------------------------------------
	raid, err := rlib.IntFromString(ss[3], "bad RAID value")
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// for now we just show all pets, so no need to parse date

	//------------------------------------------------------
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	var gxp RAPets
	var m []rlib.RentalAgreementPet
	if raid > 0 {
		m, err = rlib.GetAllRentalAgreementPets(r.Context(), raid)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	}

	gxp.Records = m

	//------------------------------------------------------
	// marshal gxp and send it!
	//------------------------------------------------------
	gxp.Status = "success"
	gxp.Total = int64(len(gxp.Records))
	fmt.Printf("gxp = %#v\n", gxp)
	b, err := json.Marshal(&gxp)
	if err != nil {
		err = fmt.Errorf("cannot marshal gxp:  %s", err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("b = %s\n", string(b))
	w.Write(b)
}

// SvcPetFeesHandler is used to get the pet fees based on provided command
//
// URL:
//       0    1       2   3
//      /v1/petfees/BID/TCID
// The server command can be:
//      recalculate
//-----------------------------------------------------------------------------
func SvcPetFeesHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcPetFeesHandler"
	var (
		err error
	)
	fmt.Printf("Entered in %s\n", funcname)

	switch d.wsSearchReq.Cmd {
	case "recalculate":
		RecalculatePetFees(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// RecalculatePetFeeRequest struct to handle
type RecalculatePetFeeRequest struct {
	FlowID int64
	RID    int64
}

// RecalculatePetFees re-calculate pet fees and make changes in flow json if required
// wsdoc {
//  @Title  Recalculate Pet Fees
//  @URL /v1/petfees/:BID/:FlowID
//  @Method  POST
//  @Synopsis recalculate pet fees
//  @Description returns flow doc with modification in pet fees
//  @Input
//  @Response FlowResponse
// wsdoc }
func RecalculatePetFees(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "RecalculatePetFees"
	/*var (
		req RecalculatePetFeeRequest
		g   FlowResponse
	)*/
	fmt.Printf("Entered in %s\n", funcname)
}
