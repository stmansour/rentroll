package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// JSONTime is a type for which we can control the JSON encoding format of the date
type JSONTime time.Time

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
	ApplicationFee            float64  // if non-zero this Prospect is an applicant
	DesiredUsageStartDate     JSONTime // predicted rent start date
	RentableTypePreference    int64    // RentableType
	FLAGS                     uint64   // 0 = Approved/NotApproved,
	Approver                  int64    // UID from Directory
	DeclineReasonSLSID        int64    // SLSid of reason
	OtherPreferences          string   // arbitrary text
	FollowUpDate              JSONTime // automatically fill out this date to sysdate + 24hrs
	CSAgent                   int64    // Accord Directory UserID - for the CSAgent
	OutcomeSLSID              int64    // id of string from a list of outcomes. Melissa to provide reasons
	FloatingDeposit           float64  // d $(GLCASH) _, c $(GLGENRCV) _; assign to a shell of a Rental Agreement
	RAID                      int64    // created to hold On Account amount of Floating Deposit
	LastModTime               JSONTime
	LastModBy                 int64
	Points                    int64
	DateofBirth               JSONTime
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

// MarshalJSON for time.Time values in JSON.  We need to override the default
// format that Go's JSON Marshal routines use for time format.
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("1/2/2006"))
	return []byte(stamp), nil
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
	var errstr string

	path := "/gsvc/"                // this is the part of the URL that got us into this handler
	uri := r.RequestURI[len(path):] // this pulls off the specific request
	sa := strings.Split(uri, "/")
	if len(sa) < 3 {
		e := fmt.Errorf("Error in URI, expecting /gsv/xperson/BID/TCID but found: %s", uri)
		SvcGridErrorReturn(w, e)
		return
	}
	d.BID, errstr = rlib.IntFromString(sa[1], "not an integer number")
	if len(errstr) > 0 {
		e := fmt.Errorf("Error in URI, expecting /gsv/xperson/BID/TCID:  BID is incorrect: %s", errstr)
		SvcGridErrorReturn(w, e)
		return
	}
	d.TCID, errstr = rlib.IntFromString(sa[2], "not an integer number")
	if len(errstr) > 0 {
		e := fmt.Errorf("Error in URI, expecting /gsv/xperson/BID/TCID:  TCID is incorrect: %s", errstr)
		SvcGridErrorReturn(w, e)
		return
	}

	fmt.Printf("BID = %d,  TCID = %d\n", d.BID, d.TCID)

	switch d.greq.Cmd {
	case "get":
		getXPerson(w, r, d)
		break
	case "save":
		saveXPerson(w, r, d)
		break
	}
}

func saveXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveXPerson"
	target := `"record":`
	fmt.Printf("SvcXPerson save\n")
	fmt.Printf("record data = %s\n", d.data)
	i := strings.Index(d.data, target)
	fmt.Printf("record is at index = %d\n", i)
	if i < 0 {
		e := fmt.Errorf("saveXPerson: cannot find %s in form json", target)
		SvcGridErrorReturn(w, e)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	fmt.Printf("data to unmarshal is:  %s\n", s)

	var gxp gxperson
	err := json.Unmarshal([]byte(s), &gxp)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	var xp rlib.XPerson
	rlib.MigrateStructVals(&gxp, &xp.Trn)
	rlib.MigrateStructVals(&gxp, &xp.Usr)
	rlib.MigrateStructVals(&gxp, &xp.Psp)
	rlib.MigrateStructVals(&gxp, &xp.Pay)

	err = rlib.UpdateTransactant(&xp.Trn)
	if err != nil {
		e := fmt.Errorf("%s: UpdateTransactant error:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	err = rlib.UpdateUser(&xp.Usr)
	if err != nil {
		e := fmt.Errorf("%s: UpdateUser error:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	err = rlib.UpdateProspect(&xp.Psp)
	if err != nil {
		e := fmt.Errorf("%s: UpdateProspect error:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	err = rlib.UpdatePayor(&xp.Pay)
	if err != nil {
		e := fmt.Errorf("%s: UpdatePayor err.Pay %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	var g struct {
		Status string `json:"status"`
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}

func getXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered getXPerson\n")
	var g struct {
		Status string   `json:"status"`
		Record gxperson `json:"record"`
	}
	var xp rlib.XPerson
	fmt.Printf("GetXPerson( TCID = %d )\n", d.TCID)
	rlib.GetXPerson(d.TCID, &xp)
	fmt.Printf("Begin migration to form struct\n")
	if xp.Pay.TCID > 0 {
		rlib.MigrateStructVals(&xp.Pay, &g.Record)
	}
	if xp.Psp.TCID > 0 {
		rlib.MigrateStructVals(&xp.Psp, &g.Record)
	}
	if xp.Usr.TCID > 0 {
		rlib.MigrateStructVals(&xp.Usr, &g.Record)
	}
	if xp.Trn.TCID > 0 {
		rlib.MigrateStructVals(&xp.Trn, &g.Record)
	}
	fmt.Printf("End migration\n")
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
