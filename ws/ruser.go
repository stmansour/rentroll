package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// UpdateRAPeopleInput is the input from the RUsers grid.
type UpdateRAPeopleInput struct {
	Cmd     string     `json:"cmd"`
	Limit   int64      `json:"limit"`
	Offset  int64      `json:"offset"`
	Changes []RAPeople `json:"changes"`
}

// SvcRUser is the dispatcher for RUser commands
// URL:
//        0   1   2    3
// 		/v1/ruser/BID/RID?&dt=2017-02-01
func SvcRUser(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "SvcRUser"
		err      error
	)
	fmt.Printf("entered %s\n", funcname)

	now := time.Now()
	d.Dt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC) // default to current date

	//----------------------------------------------------
	// pick up any HTTP GET params of interest
	//----------------------------------------------------
	q := r.URL.Query()
	if f := q["cmd"]; len(f) > 0 {
		d.wsSearchReq.Cmd = f[0]
	}
	if f := q["dt"]; len(f) > 0 {
		d.Dt, err = rlib.StringToDate(f[0])
		if err != nil {
			e := fmt.Errorf("invalid date:  %s", f[0])
			SvcGridErrorReturn(w, e, funcname)
			return
		}
	}

	//------------------------------------------------------
	//    Handle the command
	//------------------------------------------------------
	fmt.Printf("\nCOMMAND:  %s\n\n", d.wsSearchReq.Cmd)
	switch d.wsSearchReq.Cmd {
	case "get":
		d.RAID = d.ID // in this case the ID is the RAID; we get users for all rentables under the RAID
		SvcGetRAPeople("ruser", w, r, d)
	case "save":
		d.RID = d.ID // id is rid for this command
		saveRUser(w, r, d)
		return
	case "delete":
		d.RID = d.ID // id is rid for this command
		deleteRUser(w, r, d)
		return
	default:
		err = fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
	}
}

// deleteRUser deletes a rentable user
// wsdoc {
//  @Title  Delete RAPayor
//	@URL /v1/ruser/:BUI/:RID
//  @Method  GET
//	@Synopsis Delete a Rentable User
//  @Desc  This service deletes a Rentable User.
//  @Desc  then an error is returned
//	@Input DeleteRAPeople
//  @Response SvcStatusResponse
// wsdoc }
func deleteRUser(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "deleteRUser"
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var del DeleteRAPeople
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	fmt.Printf("Delete:  RID = %d, BID = %d, TCID = %d\n", d.RID, d.BID, del.TCID)

	_, err = rlib.GetRentableUserByRBT(d.RID, d.BID, del.TCID)
	if err != nil {
		e := fmt.Errorf("Error retrieving RentableUser: %s", err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	if err := rlib.DeleteRentableUserByRBT(d.RID, d.BID, del.TCID); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(w)
	return
}

// saveRUser saves or adds a new user to the RentalAgreementsUser
// wsdoc {
//  @Title  Save RUser
//	@URL /v1/ruser/:BUI/:RID
//  @Method  POST
//	@Synopsis Save an RUser
//  @Desc  This service saves a RAUser.  If :RAID exists, it will
//  @Desc  be updated with the information supplied. All fields must
//  @Desc  be supplied. If RAID is 0, then a new RAUser is created.
//	@Input RAPeopleOtherSave
//	@Input SaveRAPeopleInput
//  @Response SvcStatusResponse
// wsdoc }
func saveRUser(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "saveRUser"
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	// First determine if it is a new record, or a change...
	if strings.Contains(d.data, `"changes":`) {
		SvcUpdateRUser(w, r, d)
		return
	}

	var foo SaveRAPeopleInput
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	var a rlib.RentableUser
	fmt.Printf("foo.Record = %#v\n", foo.Record)
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	// fmt.Printf("saveRUser - first migrate: a = RID = %d, BID = %d, TCID = %d, DtStart = %s, DtStop = %s\n",
	// 	a.RID, a.BID, a.TCID, a.DtStart.Format(rlib.RRDATEFMT3), a.DtStop.Format(rlib.RRDATEFMT3))

	var bar SaveRAPeopleOther
	if err := json.Unmarshal(data, &bar); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.Record.BID.ID]
	if !ok {
		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, bar.Record.BID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	// fmt.Printf("saveRUser - second migrate: a = RID = %d, BID = %d, TCID = %d, DtStart = %s, DtStop = %s\n",
	// 	a.RID, a.BID, a.TCID, a.DtStart.Format(rlib.RRDATEFMT3), a.DtStop.Format(rlib.RRDATEFMT3))

	// make sure we don't already have this user and that there's no overlap
	// with an existing record...
	n := rlib.GetRentableUsersInRange(a.RID, &a.DtStart, &a.DtStop)
	for i := 0; i < len(n); i++ {
		if a.TCID == n[i].TCID && rlib.DateRangeOverlap(&a.DtStart, &a.DtStop, &n[i].DtStart, &n[i].DtStop) {
			e := fmt.Errorf("There is already an overlapping record for this user from %s to %s",
				n[i].DtStart.Format(rlib.RRDATEFMT4), n[i].DtStop.Format(rlib.RRDATEFMT4))
			SvcGridErrorReturn(w, e, funcname)
			return
		}
	}

	// Try to read an existing record...
	_, err = rlib.GetRentableUserByRBT(a.RID, a.BID, a.TCID)
	if err != nil && !strings.Contains(err.Error(), "no rows") {
		fmt.Printf("Error from GetRentableUserByRBT: %s\n", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	if err == nil {
		var t rlib.Transactant
		err = rlib.GetTransactant(a.TCID, &t)
		err = fmt.Errorf("%s (%s) is already listed as a user", t.GetUserName(), t.IDtoString())
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	if err = rlib.InsertRentableUser(&a); err != nil {
		e := fmt.Errorf("%s: Error saving RUser (RID=%d): %s", funcname, d.RID, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	SvcWriteSuccessResponseWithID(w, a.RUID)
}

// SvcUpdateRUser is called when a Rentable User is updated from the RentableUserGrid
func SvcUpdateRUser(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcUpdateRUser"
	)
	fmt.Printf("Entered: %s\n", funcname)
	var foo UpdateRAPeopleInput
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	// This will only contain updates.  Spin through each recid and update
	// From the grid, we only allow the following changes:  DtStart, DtStop
	for i := 0; i < len(foo.Changes); i++ {
		changes := 0
		ruser, err := rlib.GetRentableUser(foo.Changes[i].Recid)
		if err != nil {
			e := fmt.Errorf("%s: Error getting RentableUser:  %s", funcname, err.Error())
			SvcGridErrorReturn(w, e, funcname)
			return
		}
		fmt.Printf("Found ruser: %#v\n", ruser)
		dt := time.Time(foo.Changes[i].DtStart)
		if dt.Year() > 1969 {
			ruser.DtStart = dt
			changes++
		}
		dt = time.Time(foo.Changes[i].DtStop)
		if dt.Year() > 1969 {
			ruser.DtStop = dt
			changes++
		}
		if changes > 0 {
			if err := rlib.UpdateRentableUser(&ruser); err != nil {
				e := fmt.Errorf("%s: Error updating RentableUser:  %s", funcname, err.Error())
				SvcGridErrorReturn(w, e, funcname)
				return
			}
		}
	}
	SvcWriteSuccessResponse(w)
}
