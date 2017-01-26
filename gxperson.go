package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
)

// UI Data for Transactant, User, Payor, Prospect, Applicant
type gxperson struct {
	Recid                     int64 `json:"recid"` // this is to support the w2ui form
	TCID                      int64
	BID                       rlib.XJSONBud
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
	Website                   string
	Occupation                string
	ApplicationFee            float64       // if non-zero this Prospect is an applicant
	DesiredUsageStartDate     rlib.JSONTime // predicted rent start date
	RentableTypePreference    int64         // RentableType
	FLAGS                     uint64        // 0 = Approved/NotApproved,
	Approver                  int64         // UID from Directory
	DeclineReasonSLSID        int64         // SLSid of reason
	OtherPreferences          string        // arbitrary text
	FollowUpDate              rlib.JSONTime // automatically fill out this date to sysdate + 24hrs
	CSAgent                   int64         // Accord Directory UserID - for the CSAgent
	OutcomeSLSID              int64         // id of string from a list of outcomes. Melissa to provide reasons
	FloatingDeposit           float64       // d $(GLCASH) _, c $(GLGENRCV) _; assign to a shell of a Rental Agreement
	RAID                      int64         // created to hold On Account amount of Floating Deposit
	Points                    int64
	DateofBirth               rlib.JSONTime
	EmergencyContactName      string
	EmergencyContactAddress   string
	EmergencyContactTelephone string
	EmergencyEmail            string
	AlternateAddress          string
	EligibleFutureUser        rlib.XJSONYesNo
	Industry                  string
	SourceSLSID               int64
	CreditLimit               float64
	TaxpayorID                string
	AccountRep                int64
	EligibleFuturePayor       rlib.XJSONYesNo
	LastModTime               rlib.JSONTime
	LastModBy                 int64
}

// Accepts data from form submit.  Note that "list" data values are handled separately
// in gxpersonOther.  See note in grentable.go above gxrentableForm for further details.
type gxpersonForm struct {
	Recid                     int64 `json:"recid"` // this is to support the w2ui form
	TCID                      int64
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
	Website                   string
	Occupation                string
	ApplicationFee            float64       // if non-zero this Prospect is an applicant
	DesiredUsageStartDate     rlib.JSONTime // predicted rent start date
	RentableTypePreference    int64         // RentableType
	FLAGS                     uint64        // 0 = Approved/NotApproved,
	Approver                  int64         // UID from Directory
	DeclineReasonSLSID        int64         // SLSid of reason
	OtherPreferences          string        // arbitrary text
	FollowUpDate              rlib.JSONTime // automatically fill out this date to sysdate + 24hrs
	CSAgent                   int64         // Accord Directory UserID - for the CSAgent
	OutcomeSLSID              int64         // id of string from a list of outcomes. Melissa to provide reasons
	FloatingDeposit           float64       // d $(GLCASH) _, c $(GLGENRCV) _; assign to a shell of a Rental Agreement
	RAID                      int64         // created to hold On Account amount of Floating Deposit
	Points                    int64
	DateofBirth               rlib.JSONTime
	EmergencyContactName      string
	EmergencyContactAddress   string
	EmergencyContactTelephone string
	EmergencyEmail            string
	AlternateAddress          string
	Industry                  string
	SourceSLSID               int64
	CreditLimit               float64
	TaxpayorID                string
	AccountRep                int64
	LastModTime               rlib.JSONTime
	LastModBy                 int64
}

type gxpersonOther struct {
	BID                 rlib.W2uiHTMLSelect
	EligibleFutureUser  rlib.W2uiHTMLSelect
	EligibleFuturePayor rlib.W2uiHTMLSelect
}

// SvcSearchHandlerTransactants handles the search query for Transactants from the Transactant Grid.
func SvcSearchHandlerTransactants(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcSearchHandlerTransactants")
	var p rlib.Transactant
	var err error
	var g struct {
		Status  string             `json:"status"`
		Total   int64              `json:"total"`
		Records []rlib.Transactant `json:"records"`
	}

	srch := fmt.Sprintf("BID=%d", d.BID)   // default WHERE clause
	order := "LastName ASC, FirstName ASC" // default ORDER
	q, qw := gridBuildQuery("Transactant", srch, order, d, &p)
	fmt.Printf("db query = %s\n", q)

	g.Total, err = GetRowCount("Transactant", qw) // total number of rows that match the criteria
	if err != nil {
		fmt.Printf("Error from GetRowCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}

	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	i := int64(d.webreq.Offset)
	count := 0
	for rows.Next() {
		var p rlib.Transactant
		rlib.ReadTransactants(rows, &p)
		p.Recid = i
		g.Records = append(g.Records, p)
		count++ // update the count only after adding the record
		if count >= d.webreq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}
	fmt.Printf("Loaded %d transactants\n", len(g.Records))
	fmt.Printf("g.Total = %d\n", g.Total)
	rlib.Errcheck(rows.Err())
	w.Header().Set("Content-Type", "application/json")
	g.Status = "success"
	SvcWriteResponse(&g, w)
}

// SvcFormHandlerXPerson formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the TCID as follows:
//       0    1       2    3
// 		/gsvc/xperson/BID/TCID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcFormHandlerXPerson\n")
	var err error

	if d.TCID, err = SvcExtractIDFromURI(r.RequestURI, "TCID", 3, w); err != nil {
		return
	}

	fmt.Printf("Request: %s:  BID = %d,  TCID = %d\n", d.webreq.Cmd, d.BID, d.TCID)

	switch d.webreq.Cmd {
	case "get":
		getXPerson(w, r, d)
		break
	case "save":
		saveXPerson(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s\n", d.webreq.Cmd)
		SvcGridErrorReturn(w, err)
		return
	}
}

// saveXPerson handles the Save action from the Transactant Form
func saveXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveXPerson"
	target := `"record":`
	fmt.Printf("SvcFormHandlerXPerson save\n")
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

	//===============================================================
	//------------------------------
	// Handle all the non-list data
	//------------------------------
	var gxp gxpersonForm
	var xp rlib.XPerson

	err := json.Unmarshal([]byte(s), &gxp)
	if err != nil {
		fmt.Printf("Data unmarshal error: %s\n", err.Error())
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	rlib.MigrateStructVals(&gxp, &xp.Trn)
	rlib.MigrateStructVals(&gxp, &xp.Usr)
	rlib.MigrateStructVals(&gxp, &xp.Psp)
	rlib.MigrateStructVals(&gxp, &xp.Pay)

	//---------------------------
	// Handle all the list data
	//---------------------------
	var gxpo gxpersonOther
	err = json.Unmarshal([]byte(s), &gxpo)
	if err != nil {
		fmt.Printf("Data unmarshal error: %s\n", err.Error())
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	var ok bool
	xp.Trn.BID, ok = rlib.RRdb.BUDlist[gxpo.BID.ID]
	if !ok {
		e := fmt.Errorf("Could not map BID value: %s\n", gxpo.BID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	xp.Usr.BID = xp.Trn.BID
	xp.Pay.BID = xp.Trn.BID
	xp.Psp.BID = xp.Trn.BID

	xp.Usr.EligibleFutureUser, ok = rlib.YesNoMap[gxpo.EligibleFutureUser.ID]
	if !ok {
		e := fmt.Errorf("Could not map EligibleFutureUser value: %s\n", gxpo.EligibleFutureUser.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	xp.Pay.EligibleFuturePayor, ok = rlib.YesNoMap[gxpo.EligibleFuturePayor.ID]
	if !ok {
		e := fmt.Errorf("Could not map EligibleFuturePayor value: %s\n", gxpo.EligibleFuturePayor.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	//===============================================================

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
	SvcWriteSuccessResponse(w)
}

// getXPerson handles the request for an XPerson from the Transactant Form
func getXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var g struct {
		Status string   `json:"status"`
		Record gxperson `json:"record"`
	}
	var xp rlib.XPerson
	rlib.GetXPerson(d.TCID, &xp)
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
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
