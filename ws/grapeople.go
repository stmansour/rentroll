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
	Recid        int64  `json:"recid"` // this is to support the w2ui form
	TCID         int64  // associated rental agreement
	BID          int64  // Business
	FirstName    string // person name
	MiddleName   string // person name
	LastName     string // person name
	RID          int64  // Rentable ID
	RentableName string // rentable name
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
//	@URL /v1/rapeople/:BID/:RAID ? type=:PTYPE & dt=:DATE
//  @Method  GET
//	@Synopsis Return Rental Agreement payors or users
//  @Description  Return all Transactants of type :PTYPE (payor or user) on the supplied :DATE
//	@Input WebRequest
//  @Response RAPeopleResponse
// wsdoc }
//
// URL:
//       0    1       2    3
// 		/v1/rapeep/BID/RAID?type={payor|user}&dt=2017-02-01
//-----------------------------------------------------------------------------
func SvcRAPeople(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered SvcRAPeople\n")
	s := r.URL.String()                 // ex: /v1/rar/CCC/10?dt=2017-02-01
	fmt.Printf("s = %s\n", s)           // x
	s1 := strings.Split(s, "?")         // ex: /v1/rar/CCC/10?dt=2017-02-01
	fmt.Printf("s1 = %#v\n", s1)        // x
	ss := strings.Split(s1[0][1:], "/") // ex: []string{"v1", "rar", "CCC", "10"}
	fmt.Printf("ss = %#v\n", ss)

	//------------------------------------------------------
	// Handle URL path values
	//------------------------------------------------------
	raid, err := rlib.IntFromString(ss[3], "bad RAID value")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}

	//------------------------------------------------------
	// Handle URL parameters
	//------------------------------------------------------
	dt := time.Now()                   // default to current date
	ptype := "payor"                   //default to payor
	if len(s1) > 1 && len(s1[1]) > 0 { // override with whatever was provided
		parms := strings.Split(s1[1], "&") // parms is an array of indivdual parameters and their values
		for i := 0; i < len(parms); i++ {
			param := strings.Split(parms[i], "=") // an individual parameter and its value
			if len(param) < 2 {
				continue
			}
			fmt.Printf("param[i] value = %s\n", param[1])
			switch param[0] {
			case "dt":
				dt, err = rlib.StringToDate(param[1])
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
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	var gxp RAPeopleResponse
	if ptype == "payor" {
		m := rlib.GetRentalAgreementPayors(raid, &dt, &dt)
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
		m := rlib.GetRentalAgreementRentables(raid, &dt, &dt)
		k := 1                        // recid counter
		for j := 0; j < len(m); j++ { // for each rentable in the Rental Agreement
			rentable := rlib.GetRentable(m[j].RID)         // get the rentable
			n := rlib.GetRentableUsers(m[j].RID, &dt, &dt) // get the users associated with that rentable
			for i := 0; i < len(n); i++ {                  // add an entry for each user associated with this rentable
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
