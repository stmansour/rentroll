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

// SvcUserPets is used to get the pets associated with the
// TCID supplied.
//
// wsdoc {
//  @Title  People Pets
//  @URL /v1/userpets/:BUI/:TCID
//  @Method  GET
//  @Synopsis Get the pets associated with a TCID (people)
//  @Description  Returns all the pets for the supplied TCID
//  @Input
//  @Response RAPets
// wsdoc }
// URL:
//       0    1       2   3
//      /v1/userpets/BID/TCID
//-----------------------------------------------------------------------------
func SvcUserPets(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcUserPets"
	var (
		err error
	)

	fmt.Printf("entered %s\n", funcname)
	s := r.URL.String()                 // ex: /v1/userpets/CCC/10
	s1 := strings.Split(s, "?")         // ex: /v1/userpets/CCC/10
	ss := strings.Split(s1[0][1:], "/") // ex: []string{"v1", "userpets", "CCC", "10"}
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

	// for now we just show all pets, so no need to parse date

	//------------------------------------------------------
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	var resp RAPets
	var m []rlib.RentalAgreementPet
	if TCID > 0 {
		m, err = rlib.GetPetsByTransactant(r.Context(), TCID)
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

// PetFee struct
type PetFee struct {
	BID           int64
	ARID          int64
	Name          string
	DefaultAmount float64
}

// PetFeesResp is the response struct containing all pet fees
type PetFeesResp struct {
	Status  string   `json:"status"`
	Total   int64    `json:"total"`
	Records []PetFee `json:"records"`
}

// SvcPetFeesHandler is used to get the pet fees associated with the BID
//
// wsdoc {
//  @Title  Pet Fees
//  @URL /v1/petfees/:BID
//  @Method  GET
//  @Synopsis Get the pet fees associated with a BID
//  @Description  Returns all the pet fees for a BID
//  @Input
//  @Response RAPets
// wsdoc }
// URL:
//       0    1       2   3
//      /v1/userpets/BID/TCID
//-----------------------------------------------------------------------------
func SvcPetFeesHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcPetFeesHandler"
	var (
		err         error
		bizPropName = "general"
		bizPropJSON = rlib.BizProps{}
		g           PetFeesResp
	)
	fmt.Printf("Entered in %s\n", funcname)

	// get business properties
	var bizProp rlib.BusinessProperties
	bizProp, err = rlib.GetBusinessPropertiesByName(r.Context(), bizPropName, d.BID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get json doc from Data column
	if err = json.Unmarshal(bizProp.Data, &bizPropJSON); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get pet Fees
	for _, n := range bizPropJSON.PetFees {
		var pf PetFee
		var ar rlib.AR
		ar, err = rlib.GetARByName(r.Context(), d.BID, n)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		// migrate values from ar to pf
		rlib.MigrateStructVals(&ar, &pf)

		// append in response list
		g.Records = append(g.Records, pf)
	}

	g.Status = "success"
	g.Total = int64(len(bizPropJSON.PetFees))

	// success response
	SvcWriteResponse(d.BID, &g, w)
}
