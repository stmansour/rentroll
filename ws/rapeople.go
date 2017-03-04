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

// //
// type RentalAgreementPayor struct {
// 	RAID    int64
// 	BID     int64     // Business
// 	TCID    int64     // the payor's transactant id
// 	DtStart time.Time // start date/time for this Payor
// 	DtStop  time.Time // stop date/time
// 	FLAGS   uint64    // 1<<0 is the bit that indicates this payor is a 'guarantor'
// }

// RAPeopleFormSave is the structure of data we will receive from a UI form save
type RAPeopleFormSave struct {
	RAID    int64
	TCID    int64         // the payor's transactant id
	DtStart rlib.JSONTime // start date/time for this Payor
	DtStop  rlib.JSONTime // stop date/time
	FLAGS   uint64        // 1<<0 is the bit that indicates this payor is a 'guarantor'
}

// RAPeopleOtherSave is the structure of data we will receive from a UI form save
type RAPeopleOtherSave struct {
	BID rlib.W2uiHTMLSelect // Business
}

// SaveRAPeopleInput is the input data format for a Save command
type SaveRAPeopleInput struct {
	Status   string           `json:"status"`
	Recid    int64            `json:"recid"`
	FormName string           `json:"name"`
	Record   RAPeopleFormSave `json:"record"`
}

// SaveRAPeopleOther is the input data format for the "other" data on the Save command
type SaveRAPeopleOther struct {
	Status string            `json:"status"`
	Recid  int64             `json:"recid"`
	Name   string            `json:"name"`
	Record RAPeopleOtherSave `json:"record"`
}

// RAPeopleResponse is the struct containing the JSON return values for this web service
type RAPeopleResponse struct {
	Status  string     `json:"status"`
	Total   int64      `json:"total"`
	Records []RAPeople `json:"records"`
}

// DeleteRAPeople is the command structure returned when a Payor is
// deleted from the PayorList grid in the RentalAgreement Details dialog
type DeleteRAPeople struct {
	Cmd      string `json:"cmd"`
	Selected []int  `json:"selected"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	TCID     int64  `json:"TCID"`
}

type raPeopleContext struct {
	pType string // are we working on a payor or a user
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
// URL can be user or payor:
//       0    1       2    3
// 		/v1/rapeople/BID/RAID?type={payor|user}&dt=2017-02-01
// 		/v1/payor/BID/RAID?dt=2017-02-01
// 		/v1/user/BID/RAID?&dt=2017-02-01
//-----------------------------------------------------------------------------
func SvcRAPeople(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered SvcRAPeople\n")
	s := r.URL.String()                 // ex: /v1/rar/CCC/10?dt=2017-02-01
	fmt.Printf("s = %s\n", s)           // x
	s1 := strings.Split(s, "?")         // ex: /v1/rar/CCC/10?dt=2017-02-01
	fmt.Printf("s1 = %#v\n", s1)        // x
	ss := strings.Split(s1[0][1:], "/") // ex: []string{"v1", "rar", "CCC", "10"}
	fmt.Printf("ss = %#v\n", ss)
	ctx := raPeopleContext{pType: ss[1]} // pType will be user or payor

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
						ctx.pType = pTypeList[j]
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
	//    Handle the command
	//------------------------------------------------------
	fmt.Printf("\n>>>>>>>>>>>>>>>>>  COMMAND:  %s   <<<<<<<<<<<<<<<<<<<<<\n\n", d.wsSearchReq.Cmd)
	switch d.wsSearchReq.Cmd {
	case "get":
		SvcGetRAPeople(ctx.pType, w, r, d)
	case "save":
		if ctx.pType == "rapayor" {
			saveRAPayor(w, r, d)
			return
		}
		if ctx.pType == "rauser" {

		}
		SvcGridErrorReturn(w, fmt.Errorf("unhandled command for %s:  %s", ctx.pType, d.wsSearchReq.Cmd))
	case "delete":
		deleteRAPayor(w, r, d)
	default:
		SvcGridErrorReturn(w, fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd))
	}

}

// deleteRAPayor saves or adds a new payor to the RentalAgreementsPayor
// wsdoc {
//  @Title  Delete RAPayor
//	@URL /v1/rapayor/:BUI/:RAID
//  @Method  GET
//	@Synopsis Save RAPayor
//  @Desc  This service saves a RAPayor.  If :RAID exists, it will
//  @Desc  be updated with the information supplied. All fields must
//  @Desc  be supplied. If RAID is 0, then a new RAPayor is created.
//	@Input RAPeopleOtherSave
//	@Input SaveRAPeopleInput
//  @Response SvcStatusResponse
// wsdoc }
func deleteRAPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "deleteRAPayor"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)
	var del DeleteRAPeople
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	fmt.Printf("Delete:  RAID = %d, BID = %d, TCID = %d\n", d.RAID, d.BID, del.TCID)

	m := rlib.GetRentalAgreementPayors(d.RAID, &d.Dt, &d.Dt)
	if len(m) == 0 {
		e := fmt.Errorf("%s: There are no payors for this Rental Agreement", funcname)
		SvcGridErrorReturn(w, e)
		return
	}
	if len(m) == 1 {
		e := fmt.Errorf("%s: Cannot delete the only payor from a Rental Agreement.  Add another payor, then delete", funcname)
		SvcGridErrorReturn(w, e)
		return
	}
	for i := 0; i < len(m); i++ {
		if m[i].TCID != del.TCID {
			continue
		}
		if e := rlib.DeleteRentalAgreementPayor(d.RAID, d.BID, del.TCID); e != nil {
			SvcGridErrorReturn(w, e)
			return
		}
		SvcWriteSuccessResponse(w)
		return
	}
	e := fmt.Errorf("Payor with TCID %d is not a payor for Rental Agreement %s", del.TCID, rlib.IDtoString("RA", d.RAID))
	SvcGridErrorReturn(w, e)
}

// saveRAPayor saves or adds a new payor to the RentalAgreementsPayor
// wsdoc {
//  @Title  Save RAPayor
//	@URL /v1/rapayor/:BUI/:RAID
//  @Method  GET
//	@Synopsis Save RAPayor
//  @Desc  This service saves a RAPayor.  If :RAID exists, it will
//  @Desc  be updated with the information supplied. All fields must
//  @Desc  be supplied. If RAID is 0, then a new RAPayor is created.
//	@Input RAPeopleOtherSave
//	@Input SaveRAPeopleInput
//  @Response SvcStatusResponse
// wsdoc }
func saveRAPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveRAPayor"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var foo SaveRAPeopleInput
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	var a rlib.RentalAgreementPayor
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	fmt.Printf("saveRAPayor - first migrate: a = RAID = %d, BID = %d, TCID = %d, DtStart = %s, DtStop = %s\n",
		a.RAID, a.BID, a.TCID, a.DtStart.Format(rlib.RRDATEFMT3), a.DtStop.Format(rlib.RRDATEFMT3))

	var bar SaveRAPeopleOther
	if err := json.Unmarshal(data, &bar); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.Record.BID.ID]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, bar.Record.BID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	fmt.Printf("saveRAPayor - second migrate: a = RAID = %d, BID = %d, TCID = %d, DtStart = %s, DtStop = %s\n",
		a.RAID, a.BID, a.TCID, a.DtStart.Format(rlib.RRDATEFMT3), a.DtStop.Format(rlib.RRDATEFMT3))

	var err error
	// Try to read an existing record...
	_, err = rlib.GetRentalAgreementPayor(a.RAID, a.BID, a.TCID)
	if err != nil && !strings.Contains(err.Error(), "no rows") {
		fmt.Printf("Error reading RentalAgreementPayors: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}

	if err != nil {
		// This is a new RAPayor
		fmt.Printf(">>>> NEW RAPayor IS BEING ADDED\n")
		_, err = rlib.InsertRentalAgreementPayor(&a)
	} else {
		// update existing record
		fmt.Printf(">>>> Updating existing RAPayor\n")
		err = rlib.UpdateRentalAgreementPayor(&a)
	}
	if err != nil {
		e := fmt.Errorf("%s: Error saving RAPayor (RAID=%d\n: %s", funcname, d.RAID, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	SvcWriteSuccessResponse(w)
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
//      /v1/rapayor/REX/5
//-----------------------------------------------------------------------------
func SvcGetRAPeople(ptype string, w http.ResponseWriter, r *http.Request, d *ServiceData) {
	//------------------------------------------------------
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	var gxp RAPeopleResponse
	if ptype == "rapayor" {
		m := rlib.GetRentalAgreementPayors(d.RAID, &d.Dt, &d.Dt)
		for i := 0; i < len(m); i++ {
			var p rlib.Transactant
			rlib.GetTransactant(m[i].TCID, &p)
			var xr RAPeople
			fmt.Printf("before migrate: m[i].DtStart = %s, m[i].DtStop = %s\n", m[i].DtStart.Format(rlib.RRDATEFMT3), m[i].DtStop.Format(rlib.RRDATEFMT3))
			rlib.MigrateStructVals(&p, &xr)
			rlib.MigrateStructVals(&m[i], &xr)
			xr1 := time.Time(xr.DtStart)
			xr2 := time.Time(xr.DtStop)
			fmt.Printf("after migrate: xr.DtStart = %s, xr.DtStop = %s\n", xr1.Format(rlib.RRDATEFMT3), xr2.Format(rlib.RRDATEFMT3))
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
