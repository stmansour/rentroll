package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// This command returns people associated with a Rental Agreement.
// Current date is assumed unless a date is provided to override.
// type defaults to "payor" unless it is provided.  If provided it must be
// one of {payor|user}

// RAPeople defines a person for the web service interface
type RAPeople struct {
	Recid        int64         `json:"recid"` // this is to support the w2ui form
	TCID         int64         // associated rental agreement
	BID          int64         // Business
	FirstName    string        // person name
	MiddleName   string        // person name
	LastName     string        // person name
	RID          int64         // Rentable ID
	RentableName string        // rentable name
	DtStart      rlib.JSONTime // start date/time for this Rentable
	DtStop       rlib.JSONTime // stop date/time
}

// RAPeopleResponse is the struct containing the JSON return values for this web service
type RAPeopleResponse struct {
	Status  string     `json:"status"`
	Total   int64      `json:"total"`
	Records []RAPeople `json:"records"`
}

var pTypeList = []string{"payor", "user"}

// SvcRAPeople is used to get the Payor(s) or the User(s) associated with the
// RAID supplied.
//
// wsdoc {
//  @Title  Rental Agreement People
//	@URL /v1/rapeople/:BUI/:RAID ? dt=:DATE & type=:PRSTYPE
//  @Method  GET
//	@Synopsis Get Rental Agreement payors or users
//  @Description  Get the Transactants of type :PRSTYPE who are associated with the
//  @Description  Rental Agreement :RAID on the supplied :DATE.
//  @Description  Note that :PRSTYPE is optional. If it is not present, :Payor is assumed.
//	@Input none
//  @Response RAPeopleResponse
// wsdoc }
//
// URL:
//       0    1       2    3
// 		/v1/rapeople/BID/RAID?type={payor|user}&dt=2017-02-01
//-----------------------------------------------------------------------------
func SvcRAPeople(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered SvcRAPeople\n")
	s := r.URL.String()                 // ex: /v1/rar/CCC/10?dt=2017-02-01
	fmt.Printf("s = %s\n", s)           // x
	s1 := strings.Split(s, "?")         // ex: /v1/rar/CCC/10?dt=2017-02-01
	fmt.Printf("s1 = %#v\n", s1)        // x
	ss := strings.Split(s1[0][1:], "/") // ex: []string{"v1", "rar", "CCC", "10"}
	fmt.Printf("ss = %#v\n", ss)

	var err error
	//------------------------------------------------------
	// Handle URL path values
	//------------------------------------------------------
	d.RAID, err = rlib.IntFromString(ss[3], "bad RAID value")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}

	//------------------------------------------------------
	// Handle URL parameters
	//------------------------------------------------------
	d.Dt = time.Now()                  // default to current date
	ptype := "payor"                   // default to payor
	if len(s1) > 1 && len(s1[1]) > 0 { // override with whatever was provided
		parms := strings.Split(s1[1], "&") // parms is an array of indivdual parameters and their values
		for i := 0; i < len(parms); i++ {
			param := strings.Split(parms[i], "=") // an individual parameter and its value
			if len(param) < 2 {
				continue
			}
			fmt.Printf("param[i] value = %s\n", param[1])
			switch param[0] {
			case "d.Dt":
				d.Dt, err = rlib.StringToDate(param[1])
				if err != nil {
					SvcGridErrorReturn(w, fmt.Errorf("invalid date:  %s", param[1]))
					return
				}
			case "type":
				found := false
				for j := 0; j < len(pTypeList); j++ {
					if pTypeList[j] == param[1] {
						ptype = pTypeList[j]
						found = true
						break
					}
				}
				if !found {
					SvcGridErrorReturn(w, fmt.Errorf("invalid person type:  %s", param[1]))
					return
				}
			}
		}
	}

	//------------------------------------------------------
	// Handle the command
	//------------------------------------------------------
	fmt.Printf("\n>>>>>>>>>>>>>>>>>  COMMAND:  %s   <<<<<<<<<<<<<<<<<<<<<\n\n", d.wsSearchReq.Cmd)
	switch d.wsSearchReq.Cmd {
	case "get":
		SvcGetRAPeople(ptype, w, r, d)
	case "save":
		SvcGridErrorReturn(w, fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd))
	case "delete":
		SvcGridErrorReturn(w, fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd))
	default:
		SvcGridErrorReturn(w, fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd))
	}

}

// SvcGetRAPeople is used to get either the Payor(s) or User(s) associated
// with a Rental Agreement.
//
// wsdoc {
//  @Title  Rental Agreement People
//	@URL /v1/rapeople/:BUI/:RAID ? dt=:DATE & type=:PRSTYPE
//  @Method  GET
//	@Synopsis Get Rental Agreement payors or users
//  @Description  Get the Transactants of type :PRSTYPE who are associated with the
//  @Description  Rental Agreement :RAID on the supplied :DATE.
//  @Description  Note that :PRSTYPE is optional. If it is not present, :Payor is assumed.
//	@Input none
//  @Response RAPeopleResponse
// wsdoc }
//
// URL:
//       0    1       2    3
// 		/v1/rapeople/BID/RAID?type={payor|user}&dt=2017-02-01
//-----------------------------------------------------------------------------
func SvcGetRAPeople(ptype string, w http.ResponseWriter, r *http.Request, d *ServiceData) {
	//------------------------------------------------------
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	var gxp RAPeopleResponse
	if ptype == "payor" {
		m := rlib.GetRentalAgreementPayors(d.RAID, &d.Dt, &d.Dt)
		for i := 0; i < len(m); i++ {
			var p rlib.Transactant
			rlib.GetTransactant(m[i].TCID, &p)
			var xr RAPeople
			rlib.MigrateStructVals(&p, &xr)
			xr.Recid = int64(i + 1) // must set AFTER MigrateStructVals in case src contains recid
			gxp.Records = append(gxp.Records, xr)
		}
	} else if ptype == "user" {
		// first get the rentables associated with the Rental Agreement...
		m := rlib.GetRentalAgreementRentables(d.RAID, &d.Dt, &d.Dt)
		k := 1                        // recid counter
		for j := 0; j < len(m); j++ { // for each rentable in the Rental Agreement
			rentable := rlib.GetRentable(m[j].RID)             // get the rentable
			n := rlib.GetRentableUsers(m[j].RID, &d.Dt, &d.Dt) // get the users associated with that rentable
			for i := 0; i < len(n); i++ {                      // add an entry for each user associated with this rentable
				var p rlib.Transactant
				rlib.GetTransactant(n[i].TCID, &p)
				var xr RAPeople
				rlib.MigrateStructVals(&rentable, &xr)
				rlib.MigrateStructVals(&p, &xr)
				xr.Recid = int64(k) // must set AFTER MigrateStructVals in case src contains recid
				k++
				gxp.Records = append(gxp.Records, xr)
			}
		}
	} else {
		rlib.LogAndPrintError("SvcRAPeople", fmt.Errorf("Unrecognized person req: %s", ptype))
	}

	//------------------------------------------------------
	// marshal gxp and send it!
	//------------------------------------------------------
	gxp.Status = "success"
	gxp.Total = int64(len(gxp.Records))
	fmt.Printf("gxp = %#v\n", gxp)
	b, err := json.Marshal(&gxp)
	if err != nil {
		SvcGridErrorReturn(w, fmt.Errorf("cannot marshal gxp:  %s", err.Error()))
		return
	}
	//fmt.Printf("len(b) = %d\n", len(b))
	//fmt.Printf("b = %s\n", string(b))
	w.Write(b)
}
