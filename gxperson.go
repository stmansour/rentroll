package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
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
	EligibleFutureUser        int64
	Industry                  string
	SourceSLSID               int64
	CreditLimit               float64
	TaxpayorID                string
	AccountRep                int64
	EligibleFuturePayor       int64
	LastModTime               rlib.JSONTime
	LastModBy                 int64
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

	i := int64(d.greq.Offset)
	count := 0
	for rows.Next() {
		var p rlib.Transactant
		rlib.ReadTransactants(rows, &p)
		p.Recid = i
		g.Records = append(g.Records, p)
		count++ // update the count only after adding the record
		if count >= d.greq.Limit {
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
// 		/gsvc/xperson/UID/BID/TCID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcFormHandlerXPerson\n")
	var err error

	path := "/gsvc/"                // this is the part of the URL that got us into this handler
	uri := r.RequestURI[len(path):] // this pulls off the specific request
	sa := strings.Split(uri, "/")
	if len(sa) < 3 {
		e := fmt.Errorf("Error in URI, expecting /gsv/xperson/USRID/BID/TCID but found: %s", uri)
		SvcGridErrorReturn(w, e)
		return
	}
	d.UID, err = rlib.IntFromString(sa[1], "not an integer number")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	d.BID, err = rlib.IntFromString(sa[2], "not an integer number")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	d.TCID, err = rlib.IntFromString(sa[3], "not an integer number")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}

	fmt.Printf("Requester UID = %d, BID = %d,  TCID = %d\n", d.UID, d.BID, d.TCID)

	switch d.greq.Cmd {
	case "get":
		getXPerson(w, r, d)
		break
	case "save":
		saveXPerson(w, r, d)
		break
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

	var gxp gxperson
	err := json.Unmarshal([]byte(s), &gxp)
	if err != nil {
		fmt.Printf("Data unmarshal error: %s\n", err.Error())
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	fmt.Printf("Begin struct data migration\n")
	var xp rlib.XPerson
	rlib.MigrateStructVals(&gxp, &xp.Trn)
	rlib.MigrateStructVals(&gxp, &xp.Usr)
	rlib.MigrateStructVals(&gxp, &xp.Psp)
	rlib.MigrateStructVals(&gxp, &xp.Pay)
	fmt.Printf("end migration\n")

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
	// fmt.Printf("entered getXPerson\n")
	var g struct {
		Status string   `json:"status"`
		Record gxperson `json:"record"`
	}
	var xp rlib.XPerson
	// fmt.Printf("GetXPerson( TCID = %d )\n", d.TCID)
	rlib.GetXPerson(d.TCID, &xp)
	// fmt.Printf("Begin migration to form struct\n")
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
	// fmt.Printf("End migration\n")
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
