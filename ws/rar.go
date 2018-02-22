package ws

//
// This command returns the rentables associated with the supplied RAID.  If no dates are supplied
// then the current date is assumed.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// RAR is the web service structure for RentalAgreementRentables
type RAR struct {
	Recid        int64         `json:"recid"` // set to RARID
	RAID         int64         // associated rental agreement
	BID          int64         // Business
	RID          int64         // the Rentable
	RentableName string        // name of RID
	ContractRent float64       // the rent
	RARDtStart   rlib.JSONDate // start date/time for this Rentable
	RARDtStop    rlib.JSONDate // stop date/time
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
	RARDtStart   rlib.JSONDate // start date/time for this Payor
	RARDtStop    rlib.JSONDate // stop date/time
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

// UpdateRARentableInput is the input from the RARentables grid in the Rental Agreement dialog.
type UpdateRARentableInput struct {
	Cmd     string               `json:"cmd"`
	Limit   int64                `json:"limit"`
	Offset  int64                `json:"offset"`
	Changes []RARentableFormSave `json:"changes"`
}

// SvcRARentables is the dispatcher for RARentable commands
// URL:
//       0   1   2   3
// 		/v1/rar/BID/RAID
func SvcRARentables(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRARentable"
	var (
		err error
	)
	fmt.Printf("entered %s\n", funcname)
	var rcmd RARPostCmd

	now := time.Now()
	d.Dt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC) // default to current date
	rcmd.Cmd = "get"

	if r.Method == "POST" {
		if err := json.Unmarshal([]byte(d.data), &rcmd); err != nil {
			e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		if strings.Contains(d.data, `"changes":`) {
			rcmd.Cmd = "update"
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
			e := fmt.Errorf("invalid date:  %s", f[0])
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	//------------------------------------------------------
	//    Handle the command
	//------------------------------------------------------
	fmt.Printf("\nCOMMAND:  %s\n\n", rcmd.Cmd)
	d.RAID = d.ID
	switch rcmd.Cmd {
	case "get":
		GetRARentables(w, r, d)
	case "save":
		saveRARentable(w, r, d)
	case "update":
		SvcUpdateRARentable(w, r, d)
	case "delete":
		deleteRARentable(w, r, d)
	default:
		err = fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
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
//	@Input SaveRARentableInput
//  @Response SvcStatusResponse
// wsdoc }
func saveRARentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveRARentable"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var foo SaveRARentableInput
	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	var a rlib.RentalAgreementRentable
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	a.RARID = foo.Record.Recid
	a.BID, err = getBIDfromBUI(foo.Record.BUI)
	if err != nil {
		e := fmt.Errorf("%s: Could not determine Business from BUI:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	fmt.Printf("saveRARentable: a = RARID = %d, RAID = %d, BID = %d, RID = %d, ContractRent = %8.2f, DtStart = %s, DtStop = %s\n",
		a.RARID, a.RAID, a.BID, a.RID, a.ContractRent, a.RARDtStart.Format(rlib.RRDATEFMT3), a.RARDtStop.Format(rlib.RRDATEFMT3))

	m, _ := rlib.GetRentalAgreementRentables(r.Context(), d.RAID, &a.RARDtStart, &a.RARDtStop)
	for i := 0; i < len(m); i++ {
		if a.RID == m[i].RID {
			e := fmt.Errorf("That Rentable already exists in RentalAgreement %s and overlaps the time range %s - %s",
				rlib.IDtoString("RA", d.RAID), a.RARDtStart.Format(rlib.RRDATEFMT4), a.RARDtStop.Format(rlib.RRDATEFMT4))
			SvcErrorReturn(w, e, funcname)
			return
		}
	}
	fmt.Printf(">>>> NEW RARentable IS BEING ADDED\n")
	_, err = rlib.InsertRentalAgreementRentable(r.Context(), &a)
	if err != nil {
		e := fmt.Errorf("Error saving Rentable (RAID=%d): %s", d.RAID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	//-----------------------------------------------------
	// Create a Rentable Ledger marker
	//-----------------------------------------------------
	var lm = rlib.LedgerMarker{
		BID:     a.BID,
		RAID:    d.RAID,
		RID:     a.RID,
		Dt:      a.RARDtStart,
		Balance: float64(0),
		State:   rlib.LMINITIAL,
	}
	_, err = rlib.InsertLedgerMarker(r.Context(), &lm)
	if err != nil {
		e := fmt.Errorf("Error saving Rentable (RAID=%d): %s", d.RAID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.RARID) // send the new id back with the status message
}

// SvcUpdateRARentable is called when a Rentable is updated from the RentableUserGrid in the RentalAgreementDialog
func SvcUpdateRARentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcUpdateRARentable"

	fmt.Printf("Entered: %s\n", funcname)
	var foo UpdateRARentableInput

	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// This will only contain updates.  Spin through each recid and update
	// From the grid, we only allow the following changes:  RARDtStart, RARDtStop
	for i := 0; i < len(foo.Changes); i++ {
		changes := 0
		rec, err := rlib.GetRentalAgreementRentable(r.Context(), foo.Changes[i].Recid)
		if err != nil {
			e := fmt.Errorf("%s: Error getting RentalAgreementRentable:  %s", funcname, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		// The only updates allowed are to the dates and the amount.  We check those directly...
		dt := time.Time(foo.Changes[i].RARDtStart)
		if dt.Year() > 1969 {
			rec.RARDtStart = dt
			changes++
		}
		dt = time.Time(foo.Changes[i].RARDtStop)
		if dt.Year() > 1969 {
			rec.RARDtStop = dt
			changes++
		}
		if foo.Changes[i].ContractRent > float64(0) {
			rec.ContractRent = foo.Changes[i].ContractRent
			changes++
		}
		if changes > 0 {
			if err := rlib.UpdateRentalAgreementRentable(r.Context(), &rec); err != nil {
				e := fmt.Errorf("%s: Error updating RentalAgreementRentable:  %s", funcname, err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
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
	const funcname = "deleteRARentable"

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var del DeleteRARentable
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	for i := 0; i < len(del.Selected); i++ {
		if err := rlib.DeleteRentalAgreementRentable(r.Context(), del.Selected[i]); err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
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
	const funcname = "GetRARentables"
	fmt.Printf("Entered %s\n", funcname)
	var m []rlib.RentalAgreementRentable
	var rar RARList
	if d.ID > 0 {
		m, _ = rlib.GetRentalAgreementRentables(r.Context(), d.ID, &d.Dt, &d.Dt)
		fmt.Printf("d.ID = %d, d.DT = %s, len(m) = %d\n", d.ID, d.Dt.Format(rlib.RRDATEFMT3), len(m))
		for i := 0; i < len(m); i++ {
			var xr RAR
			rentable, err := rlib.GetRentable(r.Context(), m[i].RID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			rlib.MigrateStructVals(&m[i], &xr)
			xr.RentableName = rentable.RentableName
			xr.Recid = m[i].RARID
			rar.Records = append(rar.Records, xr)
		}
	}
	rar.Status = "success"
	rar.Total = int64(len(m))
	SvcWriteResponse(d.BID, &rar, w)
}
