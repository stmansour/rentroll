package ws

//
// This command returns the rentables associated with the supplied RAID.  If no dates are supplied
// then the current date is assumed.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// RAR is the web service structure for RentalAgreementRentables
type RAR struct {
	Recid        int64         `json:"recid"` // this is to support the w2ui form
	RAID         int64         // associated rental agreement
	BID          int64         // Business
	RID          int64         // the Rentable
	RentableName string        // name of RID
	ContractRent float64       // the rent
	RARDtStart   rlib.JSONTime // start date/time for this Rentable
	RARDtStop    rlib.JSONTime // stop date/time
}

// RARList is the struct containing the JSON return values for this web service
type RARList struct {
	Status  string `json:"status"`
	Total   int64  `json:"total"`
	Records []RAR  `json:"records"`
}

// RARentableFormSave is the structure of data we will receive from a UI form save
type RARentableFormSave struct {
	Recid        int64         `json:"recid"` // this is to support the w2ui form
	RAID         int64         // which RAID
	RID          int64         // the rentable id
	BUI          string        // in this case we could get an BID or a BUD
	RentableName string        // name of RID
	ContractRent float64       // the rent
	RARDtStart   rlib.JSONTime // start date/time for this Payor
	RARDtStop    rlib.JSONTime // stop date/time
}

// SaveRARentableInput is the input data format for a Save command
type SaveRARentableInput struct {
	Cmd      string             `json:"cmd"`
	Recid    int64              `json:"recid"`
	FormName string             `json:"name"`
	Record   RARentableFormSave `json:"record"`
}

// DeleteRARentable is the command structure returned when a Payor is
// deleted from the PayorList grid in the RentalAgreement Details dialog
type DeleteRARentable struct {
	Cmd      string  `json:"cmd"`
	Selected []int64 `json:"selected"`
	Limit    int     `json:"limit"`
	Offset   int     `json:"offset"`
	RID      int64   `json:"RID"`
}

// SvcRARentables is the dispatcher for RARentable commands
// URL:
//       0   1   2   3
// 		/v1/rar/BID/RAID
func SvcRARentables(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var err error
	funcname := "SvcRARentable"
	fmt.Printf("\tentered %s\n", funcname)
	var rcmd RARPostCmd

	now := time.Now()
	d.Dt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC) // default to current date
	rcmd.Cmd = "get"

	if r.Method == "POST" {
		if err := json.Unmarshal([]byte(d.data), &rcmd); err != nil {
			e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
			SvcGridErrorReturn(w, e)
			return
		}
	}

	//----------------------------------------------------
	// pick up any HTTP GET params of interest
	//----------------------------------------------------
	q := r.URL.Query()
	if f := q["cmd"]; len(f) > 0 {
		rcmd.Cmd = f[0]
	}
	if f := q["dt"]; len(f) > 0 {
		d.Dt, err = rlib.StringToDate(f[0])
		if err != nil {
			SvcGridErrorReturn(w, fmt.Errorf("invalid date:  %s", f[0]))
			return
		}
	}

	//------------------------------------------------------
	//    Handle the command
	//------------------------------------------------------
	fmt.Printf("\nCOMMAND:  %s\n\n", d.wsSearchReq.Cmd)
	d.RAID = d.ID
	switch rcmd.Cmd {
	case "get":
		GetRARentables(w, r, d)
	case "save":
		saveRARentable(w, r, d)
		return
	case "delete":
		deleteRARentable(w, r, d)
		return
	default:
		SvcGridErrorReturn(w, fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd))
	}
}

// saveRARentable saves or adds a new payor to the RentalAgreementsPayor
// wsdoc {
//  @Title  Save RARentable
//	@URL /v1/rar/:BUI/:RAID
//  @Method  POST
//	@Synopsis Save RARentable
//  @Desc  This service saves a RARentable.  If :RAID exists, it will
//  @Desc  be updated with the information supplied. All fields must
//  @Desc  be supplied. If RAID is 0, then a new RARentable is created.
//	@Input RARentableOtherSave
//	@Input SaveRARentableInput
//  @Response SvcStatusResponse
// wsdoc }
func saveRARentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveRARentable"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var foo SaveRARentableInput
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	var a rlib.RentalAgreementRentable
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	a.RARID = foo.Record.Recid
	a.BID = getBIDfromBUI(foo.Record.BUI)

	fmt.Printf("saveRARentable: a = RARID = %d, RAID = %d, BID = %d, RID = %d, ContractRent = %8.2f, DtStart = %s, DtStop = %s\n",
		a.RARID, a.RAID, a.BID, a.RID, a.ContractRent, a.RARDtStart.Format(rlib.RRDATEFMT3), a.RARDtStop.Format(rlib.RRDATEFMT3))

	var err error
	// // Try to read an existing record...
	// if a.RARID > 0 {
	// 	_, err = rlib.GetRentalAgreementRentable(rarid)
	// 	if err != nil && !strings.Contains(err.Error(), "no rows") {
	// 		fmt.Printf("Error reading RentalAgreementPayors: %s\n", err.Error())
	// 		SvcGridErrorReturn(w, err)
	// 		return
	// 	}
	// }

	if a.RARID == 0 {
		// This is a new RARentable
		fmt.Printf(">>>> NEW RARentable IS BEING ADDED\n")
		_, err = rlib.InsertRentalAgreementRentable(&a)
	} else {
		// update existing record
		fmt.Printf(">>>> Updating existing RARentable\n")
		err = rlib.UpdateRentalAgreementRentable(&a)
	}
	if err != nil {
		e := fmt.Errorf("%s: Error saving RARentable (RAID=%d\n: %s", funcname, d.RAID, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	SvcWriteSuccessResponseWithID(w, a.RARID) // send the new id back with the status message
}

// deleteRARentable deletes a rentable from a rental agreement
// wsdoc {
//  @Title  Delete RARentable
//	@URL /v1/rar/:BUI/:RAID
//  @Method  POST
//	@Synopsis Delete a Rental Agreement Payor
//  @Desc  This service deletes a RARentable. If this is the only rentable
//  @Desc  then an error is returned
//	@Input DeleteRARentable
//  @Response SvcStatusResponse
// wsdoc }
func deleteRARentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "deleteRARentable"
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)
	var del DeleteRARentable
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	for i := 0; i < len(del.Selected); i++ {
		if err := rlib.DeleteRentalAgreementRentable(del.Selected[i]); err != nil {
			SvcGridErrorReturn(w, err)
			return
		}
	}
	SvcWriteSuccessResponse(w)
}

// GetRARentables returns the Rentables associated with the RAID supplied
//  Called with URL:
//       0    1   2   3
// 		/v1/rar/BID/RAID?dt=2017-01-03
// wsdoc {
//  @Title  Rental Agreement Rentables
//	@URL /v1/rar/:BUI/:RAID [ ? dt=:DATE ]
//	@Method GET
//	@Synopsis Get Rentables for Rental Agreement :RAID
//  @Desc Returns all the rentables associated Rental Agreement RAID as of :DATE
//  @Desc If DATE is not provided, then the current date is assumed.
//  @Input none
//  @Response RAR
// wsdoc }
func GetRARentables(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var rar RARList
	m := rlib.GetRentalAgreementRentables(d.ID, &d.Dt, &d.Dt)
	fmt.Printf("d.ID = %d, d.DT = %s, len(m) = %d\n", d.ID, d.Dt.Format(rlib.RRDATEFMT3), len(m))
	for i := 0; i < len(m); i++ {
		var xr RAR
		xr.Recid = int64(i + 1)
		r := rlib.GetRentable(m[i].RID)
		rlib.MigrateStructVals(&m[i], &xr)
		xr.RentableName = r.RentableName
		rar.Records = append(rar.Records, xr)
	}
	rar.Status = "success"
	rar.Total = int64(len(m))
	SvcWriteResponse(&rar, w)
}
