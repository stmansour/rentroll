package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// This command returns people associated with a Rental Agreement.
// Current date is assumed unless a date is provided to override.
// type defaults to "payor" unless it is provided.  If provided it must be
// one of {payor|user}

// RAPeopleFormSave is the structure of data we will receive from a UI form save
type RAPeopleFormSave struct {
	RAID    int64
	BID     int64
	TCID    int64         // the payor's transactant id
	RID     int64         // same struct type used for adding Users.  RID will be populated here, not RAID
	DtStart rlib.JSONDate // start date/time for this Payor
	DtStop  rlib.JSONDate // stop date/time
	FLAGS   uint64        // 1<<0 is the bit that indicates this payor is a 'guarantor'
}

// SaveRAPeopleInput is the input data format for a Save command
type SaveRAPeopleInput struct {
	Status   string           `json:"status"`
	Recid    int64            `json:"recid"`
	FormName string           `json:"name"`
	Record   RAPeopleFormSave `json:"record"`
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

// RAPeople defines a person for the web service interface
type RAPeople struct {
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

// RAPeopleResponse is the struct containing the JSON return values for this web service
type RAPeopleResponse struct {
	Status  string     `json:"status"`
	Total   int64      `json:"total"`
	Records []RAPeople `json:"records"`
}

type raPeopleContext struct {
	pType string // are we working on a payor or a user
}

var pTypeList = []string{"payor", "user"}

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
	const funcname = "SvcGetRAPeople"
	//------------------------------------------------------
	// Get the transactants... either payors or users...
	//------------------------------------------------------
	var (
		gxp RAPeopleResponse
		err error
	)
	if ptype == "rapayor" {
		m, _ := rlib.GetRentalAgreementPayorsInRange(r.Context(), d.RAID, &d.Dt, &d.Dt)
		for i := 0; i < len(m); i++ {
			var p rlib.Transactant
			err = rlib.GetTransactant(r.Context(), m[i].TCID, &p)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
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
	} else if ptype == "ruser" {
		// first get the rentables associated with the Rental Agreement...
		m, _ := rlib.GetRentalAgreementRentables(r.Context(), d.RAID, &d.Dt, &d.Dt)
		fmt.Printf("GetRentalAgreementRentables for RAID = %d, date = %s,  return count = %d\n", d.RAID, d.Dt.Format(rlib.RRDATEFMT3), len(m))
		k := 1                        // recid counter
		for j := 0; j < len(m); j++ { // for each rentable in the Rental Agreement
			rentable, err := rlib.GetRentable(r.Context(), m[j].RID) // get the rentable
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			n, err := rlib.GetRentableUsersInRange(r.Context(), m[j].RID, &d.Dt, &d.Dt) // get the users associated with that rentable
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			fmt.Printf("Rentable: %d, date = %s, rentable user count: %d\n", m[j].RID, d.Dt.Format(rlib.RRDATEFMT3), len(n))
			for i := 0; i < len(n); i++ { // add an entry for each user associated with this rentable
				var p rlib.Transactant
				err = rlib.GetTransactant(r.Context(), n[i].TCID, &p)
				if err != nil {
					SvcErrorReturn(w, err, funcname)
					return
				}
				var xr RAPeople
				rlib.MigrateStructVals(&n[i], &xr)
				rlib.MigrateStructVals(&rentable, &xr)
				rlib.MigrateStructVals(&p, &xr)
				xr.Recid = n[i].RUID
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
	SvcWriteResponse(d.BID, &gxp, w)
}
