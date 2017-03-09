package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// SvcRAPayor is the dispatcher for RAPayor commands
// URL:
//       0   1       2    3
// 		/v1/rapayor/BID/RAID?dt=2017-02-01
func SvcRAPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var err error
	fmt.Printf("\tentered SvcRAPayor\n")
	s := r.URL.String()                 // ex: /v1/rar/CCC/10?dt=2017-02-01
	fmt.Printf("\ts = %s\n", s)         // x
	s1 := strings.Split(s, "?")         // ex: /v1/rar/CCC/10?dt=2017-02-01
	fmt.Printf("\ts1 = %#v\n", s1)      // x
	ss := strings.Split(s1[0][1:], "/") // ex: []string{"v1", "rar", "CCC", "10"}
	fmt.Printf("\tss = %#v\n", ss)

	//------------------------------------------------------
	// Handle URL path values
	//------------------------------------------------------
	d.RAID, err = rlib.IntFromString(ss[3], "bad ID value")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}

	//------------------------------------------------------
	// Handle URL parameters
	//------------------------------------------------------
	now := time.Now()
	d.Dt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC) // default to current date
	if len(s1) > 1 && len(s1[1]) > 0 {                                         // override with whatever was provided
		parms := strings.Split(s1[1], "&") // parms is an array of indivdual parameters and their values
		for i := 0; i < len(parms); i++ {
			param := strings.Split(parms[i], "=") // an individual parameter and its value
			if len(param) < 2 {
				continue
			}
			fmt.Printf("param[i] value = %s\n", param[1])
			switch param[0] {
			case "cmd":
				d.wsSearchReq.Cmd = strings.TrimSpace(param[1])
			case "dt":
				d.Dt, err = rlib.StringToDate(param[1])
				if err != nil {
					SvcGridErrorReturn(w, fmt.Errorf("invalid date:  %s", param[1]))
					return
				}
			}
		}
	}

	//------------------------------------------------------
	//    Handle the command
	//------------------------------------------------------
	fmt.Printf("\nCOMMAND:  %s\n\n", d.wsSearchReq.Cmd)
	switch d.wsSearchReq.Cmd {
	case "get":
		SvcGetRAPeople("rapayor", w, r, d)
	case "save":
		saveRAPayor(w, r, d)
		return
	case "delete":
		deleteRAPayor(w, r, d)
		return
	default:
		SvcGridErrorReturn(w, fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd))
	}
}

// deleteRAPayor deletes a payor from a rental agreement
// wsdoc {
//  @Title  Delete RAPayor
//	@URL /v1/rapayor/:BUI/:RAID
//  @Method  GET
//	@Synopsis Delete a Rental Agreement Payor
//  @Desc  This service deletes a RAPayor. If this is the only payor
//  @Desc  then an error is returned
//	@Input DeleteRAPeople
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
		if e := rlib.DeleteRentalAgreementPayorByRBT(d.RAID, d.BID, del.TCID); e != nil {
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
		err = rlib.UpdateRentalAgreementPayorByRBT(&a)
	}
	if err != nil {
		e := fmt.Errorf("%s: Error saving RAPayor (RAID=%d\n: %s", funcname, d.RAID, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	SvcWriteSuccessResponse(w)
}
