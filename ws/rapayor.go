package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// RAPayor defines a person for the web service interface
type RAPayor struct {
	Recid        int64         `json:"recid"` // this the RAPID
	TCID         int64         // associated rental agreement
	BID          int64         // Business
	FirstName    string        // person name
	MiddleName   string        // person name
	LastName     string        // person name
	RID          int64         // Rentable ID
	RentableName string        // rentable name
	DtStart      rlib.JSONDate // start date/time for this Rentable
	DtStop       rlib.JSONDate // stop date/time
	IsCompany    bool
	CompanyName  string
}

// RAPayorFormSave defines a person for the web service interface
type RAPayorFormSave struct {
	Recid        int64 `json:"recid"` // this the RAPID
	BID          int64
	TCID         int64         // associated rental agreement
	RAID         int64         // which rental agreement
	FirstName    string        // person name
	MiddleName   string        // person name
	LastName     string        // person name
	RID          int64         // Rentable ID
	RentableName string        // rentable name
	DtStart      rlib.JSONDate // start date/time for this Rentable
	DtStop       rlib.JSONDate // stop date/time
}

// RAPayorResponse is the struct containing the JSON return values for this web service
type RAPayorResponse struct {
	Status  string    `json:"status"`
	Total   int64     `json:"total"`
	Records []RAPayor `json:"records"`
}

// DeleteRAPayor is the command structure returned when a Payor is
// deleted from the PayorList grid in the RentalAgreement Details dialog
type DeleteRAPayor struct {
	Cmd      string  `json:"cmd"`
	Selected []int64 `json:"selected"` // this will contain the RAPIDs of the payors to delete
	Limit    int     `json:"limit"`
	Offset   int     `json:"offset"`
	DtStart  rlib.JSONDate
	DtStop   rlib.JSONDate
}

// RARPostCmd is the input data format for a Save command
type RARPostCmd struct {
	Cmd      string `json:"cmd"`
	Recid    int64  `json:"recid"`
	FormName string `json:"name"`
}

// SaveRAPayorInput is the input data format for a Save command
type SaveRAPayorInput struct {
	Status   string          `json:"status"`
	Recid    int64           `json:"recid"`
	FormName string          `json:"name"`
	Record   RAPayorFormSave `json:"record"`
}

// UpdateRAPayorInput is the input from the RUsers grid.
type UpdateRAPayorInput struct {
	Cmd     string            `json:"cmd"`
	Limit   int64             `json:"limit"`
	Offset  int64             `json:"offset"`
	Changes []RAPayorFormSave `json:"changes"`
}

// SvcRAPayor is the dispatcher for RAPayor commands
// URL:
//       0   1       2    3
// 		/v1/rapayor/BID/RAID?dt=2017-02-01
func SvcRAPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRAPayor"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)

	now := time.Now()
	d.Dt = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC) // default to current date
	d.RAID = d.ID

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
			err = fmt.Errorf("invalid date:  %s", f[0])
			SvcErrorReturn(w, err, funcname)
			return
		}
	}

	//------------------------------------------------------
	//    Handle the command
	//------------------------------------------------------
	fmt.Printf("\nCOMMAND:  %s\n\n", d.wsSearchReq.Cmd)
	fmt.Printf("\tRAID = %d\n", d.RAID)

	switch d.wsSearchReq.Cmd {
	case "get":
		SvcGetRAPayor(w, r, d)
	case "save":
		saveRAPayor(w, r, d)
		return
	case "delete":
		deleteRAPayor(w, r, d)
		return
	default:
		err = fmt.Errorf("unhandled command:  %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
	}
}

// deleteRAPayor deletes a payor from a rental agreement
// wsdoc {
//  @Title  Delete RAPayor
//	@URL /v1/rapayor/:BUI/:RAID
//  @Method  POST
//	@Synopsis Delete a Rental Agreement Payor
//  @Desc  This service deletes a RAPayor. If this is the only payor
//  @Desc  then an error is returned
//	@Input DeleteRAPayor
//  @Response SvcStatusResponse
// wsdoc }
func deleteRAPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteRAPayor"

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	var del DeleteRAPayor
	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// dtStart := time.Time(del.DtStart)
	// dtStop := time.Time(del.DtStop)

	for i := 0; i < len(del.Selected); i++ {
		// // first validate that the selected ids are part of the supplied raid
		// rlib.Console("Looking for payors in range:  RAID=%d, dtStart = %s, dtStop = %s\n", d.RAID, dtStart.Format(rlib.RRDATEREPORTFMT), dtStop.Format(rlib.RRDATEREPORTFMT))
		// m, _ := rlib.GetRentalAgreementPayorsInRange(r.Context(), d.RAID, &dtStart, &dtStop)
		// for j := 0; j < len(m); j++ {
		// 	if m[j].RAPID == del.Selected[i] {
		// 		if err := rlib.DeleteRentalAgreementPayor(r.Context(), del.Selected[i]); err != nil {
		// 			SvcErrorReturn(w, err, funcname)
		// 			return
		// 		}
		// 		SvcWriteSuccessResponse(d.BID, w)
		// 		return
		// 	}
		// }
		if err := rlib.DeleteRentalAgreementPayor(r.Context(), del.Selected[i]); err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		SvcWriteSuccessResponse(d.BID, w)
		return
	}
	e := fmt.Errorf("%s: Payor is was not listed as a payor for Rental Agreement %d during that time period", funcname, d.RAID)
	SvcErrorReturn(w, e, funcname)
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
//	@Input SaveRAPayorInput
//  @Response SvcStatusResponse
// wsdoc }
func saveRAPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveRAPayor"
	var (
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	// First determine if it is a new record, or a change...
	if strings.Contains(d.data, `"changes":`) {
		SvcUpdateRAPayor(w, r, d)
		return
	}
	var foo SaveRAPayorInput
	data := []byte(d.data)
	if err := json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	var a rlib.RentalAgreementPayor
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling

	fmt.Printf("saveRAPayor - first migrate: a = RAID = %d, BID = %d, TCID = %d, DtStart = %s, DtStop = %s\n",
		a.RAID, a.BID, a.TCID, a.DtStart.Format(rlib.RRDATEFMT3), a.DtStop.Format(rlib.RRDATEFMT3))

	// Try to read an existing record...
	m, _ := rlib.GetRentalAgreementPayorsInRange(r.Context(), a.RAID, &a.DtStart, &a.DtStop)
	fmt.Printf("found %d payors for RAID %d during period %s - %s\n", len(m), a.RAID, a.DtStart.Format(rlib.RRDATEFMT4), a.DtStop.Format(rlib.RRDATEFMT4))
	for i := 0; i < len(m); i++ {
		fmt.Printf("m[i].TCID = %d, a.TCID = %d, %s - %s,  %s - %s\n",
			m[i].TCID, a.TCID,
			a.DtStart.Format(rlib.RRDATEFMT4), a.DtStop.Format(rlib.RRDATEFMT4),
			m[i].DtStart.Format(rlib.RRDATEFMT4), m[i].DtStop.Format(rlib.RRDATEFMT4))
		fmt.Printf("DateRangeOverlap = %t\n", rlib.DateRangeOverlap(&a.DtStart, &a.DtStop, &m[i].DtStart, &m[i].DtStop))
		if m[i].TCID == a.TCID && rlib.DateRangeOverlap(&a.DtStart, &a.DtStop, &m[i].DtStart, &m[i].DtStop) {
			e := fmt.Errorf("There is already an overlapping record for %s %s from %s to %s",
				foo.Record.FirstName, foo.Record.LastName,
				time.Time(foo.Record.DtStart).Format(rlib.RRDATEFMT4),
				time.Time(foo.Record.DtStop).Format(rlib.RRDATEFMT4))
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	// This is a new RAPayor
	_, err = rlib.InsertRentalAgreementPayor(r.Context(), &a)
	if err != nil {
		e := fmt.Errorf("%s: Error saving RAPayor (RAID=%d\n: %s", funcname, d.RAID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.RAPID)
}

// SvcUpdateRAPayor is called when a Rentable User is updated from the RentableUserGrid
func SvcUpdateRAPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcUpdateRAPayor"
	var (
		err error
	)

	fmt.Printf("Entered: %s\n", funcname)
	var foo UpdateRAPayorInput
	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// This will only contain updates.  Spin through each recid and update
	// From the grid, we only allow the following changes:  DtStart, DtStop
	for i := 0; i < len(foo.Changes); i++ {
		changes := 0
		rapayor, err := rlib.GetRentalAgreementPayor(r.Context(), foo.Changes[i].Recid)
		if err != nil {
			e := fmt.Errorf("%s: Error getting RentalAgreementPayor:  %s", funcname, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		fmt.Printf("Found rapayor: %#v\n", rapayor)
		dt := time.Time(foo.Changes[i].DtStart)
		if dt.Year() > 1969 {
			rapayor.DtStart = dt
			changes++
		}
		dt = time.Time(foo.Changes[i].DtStop)
		if dt.Year() > 1969 {
			rapayor.DtStop = dt
			changes++
		}
		if changes > 0 {
			if err := rlib.UpdateRentalAgreementPayor(r.Context(), &rapayor); err != nil {
				e := fmt.Errorf("%s: Error updating RentalAgreementPayor:  %s", funcname, err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// SvcGetRAPayor is used to get either the Payor(s) or User(s) associated
// with a Rental Agreement.
//
// wsdoc {
//  @Title  Rental Agreement Payor
//	@URL /v1/rapayor/:BUI/:RAID ? dt=:DATE & type=:PRSTYPE
//  @Method  GET
//	@Synopsis Get Rental Agreement payors or users
//  @Description  Get the Transactants of type :PRSTYPE who are associated with the
//  @Description  Rental Agreement :RAID on the supplied :DATE.
//  @Description  Note that :PRSTYPE is optional. If it is not present, :Payor is assumed.
//	@Input none
//  @Response RAPayorResponse
// wsdoc }
//
// URL:
//       0    1       2    3
// 		/v1/RAPayor/BID/RAID?type={payor|user}&dt=2017-02-01
//      /v1/rapayor/REX/5
//-----------------------------------------------------------------------------
func SvcGetRAPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	//------------------------------------------------------
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	const funcname = "SvcGetRAPayor"
	var (
		gxp RAPayorResponse
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	d1 := d.wsSearchReq.SearchDtStart
	d2 := d.wsSearchReq.SearchDtStop

	//--------------------------------------------------------------------------
	// This is a hack... at some point the date being sent up was not correct.
	// Since we can display all the Payors we need to, if the daterange looks
	// incorrect it gets set to all-time
	//--------------------------------------------------------------------------
	if d1.Year() < 1970 {
		d1 = rlib.TIME0
		d2 = rlib.ENDOFTIME
	}

	m, _ := rlib.GetRentalAgreementPayorsInRange(r.Context(), d.RAID, &d1, &d2)
	for i := 0; i < len(m); i++ {
		var p rlib.Transactant
		err = rlib.GetTransactant(r.Context(), m[i].TCID, &p)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		var xr RAPayor
		fmt.Printf("before migrate: m[i].DtStart = %s, m[i].DtStop = %s\n", m[i].DtStart.Format(rlib.RRDATEFMT3), m[i].DtStop.Format(rlib.RRDATEFMT3))
		rlib.MigrateStructVals(&p, &xr)
		rlib.MigrateStructVals(&m[i], &xr)
		xr1 := time.Time(xr.DtStart)
		xr2 := time.Time(xr.DtStop)
		xr.Recid = int64(i)
		fmt.Printf("after migrate: xr.DtStart = %s, xr.DtStop = %s\n", xr1.Format(rlib.RRDATEFMT3), xr2.Format(rlib.RRDATEFMT3))
		xr.Recid = m[i].RAPID // must set AFTER MigrateStructVals in case src contains recid
		gxp.Records = append(gxp.Records, xr)
	}
	//------------------------------------------------------
	// marshal gxp and send it!
	//------------------------------------------------------
	gxp.Status = "success"
	gxp.Total = int64(len(gxp.Records))
	SvcWriteResponse(d.BID, &gxp, w)
}
