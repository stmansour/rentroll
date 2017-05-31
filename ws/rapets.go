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
	var (
		funcname = "SvcRAPets"
		err      error
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
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	// for now we just show all pets, so no need to parse date

	//------------------------------------------------------
	// Handle URL parameters
	//------------------------------------------------------
	// dt := time.Now()                   // default to current date
	// if len(s1) > 1 && len(s1[1]) > 0 { // override with whatever was provided
	// 	parms := strings.Split(s1[1], "&") // parms is an array of indivdual parameters and their values
	// 	for i := 0; i < len(parms); i++ {
	// 		param := strings.Split(parms[i], "=") // an individual parameter and its value
	// 		if len(param) < 2 {
	// 			continue
	// 		}
	// 		fmt.Printf("param[i] value = %s\n", param[1])
	// 		switch param[0] {
	// 		case "dt":
	// 			dt, err = rlib.StringToDate(param[1])
	// 			if err != nil {
	// 				SvcGridErrorReturn(w, fmt.Errorf("invalid date:  %s", param[1]))
	// 				return
	// 			}
	// 		}
	// 	}
	// }

	//------------------------------------------------------
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	var gxp RAPets
	var m []rlib.RentalAgreementPet
	if raid > 0 {
		m = rlib.GetAllRentalAgreementPets(raid)
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
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("b = %s\n", string(b))
	w.Write(b)
}
