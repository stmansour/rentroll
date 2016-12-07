package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// this is a structure specifically for the UI. It will be
// automatically populated from an rlib.XPerson struct
type gxperson struct {
	Recid                     int64 `json:"recid"` // this is to support the w2ui form
	TCID                      int64
	BID                       int64
	NLID                      int64
	FirstName                 string
	MiddleName                string
	LastName                  string
	PreferredName             string
	CompanyName               string // sometimes the entity will be a company
	IsCompany                 int    // 1 => the entity is a company, 0 = not a company
	PrimaryEmail              string
	SecondaryEmail            string
	WorkPhone                 string
	CellPhone                 string
	Address                   string
	Address2                  string
	City                      string
	State                     string
	PostalCode                string
	Country                   string
	EmployerName              string
	EmployerStreetAddress     string
	EmployerCity              string
	EmployerState             string
	EmployerPostalCode        string
	EmployerEmail             string
	EmployerPhone             string
	Occupation                string
	ApplicationFee            float64   // if non-zero this Prospect is an applicant
	DesiredUsageStartDate     time.Time // predicted rent start date
	RentableTypePreference    int64     // RentableType
	FLAGS                     uint64    // 0 = Approved/NotApproved,
	Approver                  int64     // UID from Directory
	DeclineReasonSLSID        int64     // SLSid of reason
	OtherPreferences          string    // arbitrary text
	FollowUpDate              time.Time // automatically fill out this date to sysdate + 24hrs
	CSAgent                   int64     // Accord Directory UserID - for the CSAgent
	OutcomeSLSID              int64     // id of string from a list of outcomes. Melissa to provide reasons
	FloatingDeposit           float64   // d $(GLCASH) _, c $(GLGENRCV) _; assign to a shell of a Rental Agreement
	RAID                      int64     // created to hold On Account amount of Floating Deposit
	LastModTime               time.Time
	LastModBy                 int64
	Points                    int64
	DateofBirth               time.Time
	EmergencyContactName      string
	EmergencyContactAddress   string
	EmergencyContactTelephone string
	EmergencyEmail            string
	AlternateAddress          string
	EligibleFutureUser        int64
	Industry                  string
	SourceSLSID               int64
	CreditLimit               float64
	TaxpayorID                string
	AccountRep                int64
	EligibleFuturePayor       int64
}

// SvcXPerson formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the TCID as follows:
// 		/gsvc/xperson/BID/TCID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcXPerson\n")

	var g struct {
		Status string   `json:"status"`
		Record gxperson `json:"record"`
	}

	path := "/gsvc/"                // this is the part of the URL that got us into this handler
	uri := r.RequestURI[len(path):] // this pulls off the specific request
	sa := strings.Split(uri, "/")
	if len(sa) < 3 {
		e := fmt.Errorf("Error in URI, expecting /gsv/xperson/BID/TCID but found: %s", uri)
		SvcGridErrorReturn(w, e)
		return
	}
	bid, errstr := rlib.IntFromString(sa[1], "not an integer number")
	if len(errstr) > 0 {
		e := fmt.Errorf("Error in URI, expecting /gsv/xperson/BID/TCID:  BID is incorrect: %s", errstr)
		SvcGridErrorReturn(w, e)
		return
	}
	tcid, errstr := rlib.IntFromString(sa[2], "not an integer number")
	if len(errstr) > 0 {
		e := fmt.Errorf("Error in URI, expecting /gsv/xperson/BID/TCID:  TCID is incorrect: %s", errstr)
		SvcGridErrorReturn(w, e)
		return
	}

	fmt.Printf("bid = %d,  tcid = %d\n", bid, tcid)

	switch d.greq.Cmd {
	case "get":
		var xp rlib.XPerson
		rlib.GetXPerson(tcid, &xp)
		if xp.Pay.TCID > 0 {
			rlib.MigrateStructVals(&xp.Pay, &g.Record)
		}
		if xp.Psp.TCID > 0 {
			rlib.MigrateStructVals(&xp.Psp, &g.Record)
		}
		if xp.Tnt.TCID > 0 {
			rlib.MigrateStructVals(&xp.Tnt, &g.Record)
		}
		if xp.Trn.TCID > 0 {
			rlib.MigrateStructVals(&xp.Trn, &g.Record)
		}
		g.Status = "success"
		b, err := json.Marshal(g)
		if err != nil {
			e := fmt.Errorf("SvcXPerson: Error marshaling json data: %s", err.Error())
			SvcGridErrorReturn(w, e)
			return
		}
		// fmt.Printf("first 100 chars of response: %100.100s\n", string(b))
		fmt.Printf("Response Data:  %s\n", string(b))

		w.Write(b)
		break
	case "save":
		break
	}

}
