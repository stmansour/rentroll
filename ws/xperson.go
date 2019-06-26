package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// RPersonForm contains attributes of Transactant, User, Payor, Prospect, Applicant
// RPersonForm is the expected return data format for updating a person.
//  Note that "list" data values are handled separately
//	in RPersonOther.  See note in grentable.go above grentableForm for further details.
type RPersonForm struct {
	Recid int64 `json:"recid"` // this is to support the w2ui form
	TCID  int64
	BID   int64
	BUD   rlib.XJSONBud
	NLID  int64

	// --------------- Transactant --------------
	FirstName      string
	MiddleName     string
	LastName       string
	PreferredName  string
	IsCompany      bool   // 1 => the entity is a company, 0 = not a company
	CompanyName    string // sometimes the entity will be a company
	PrimaryEmail   string
	SecondaryEmail string
	WorkPhone      string
	CellPhone      string
	Address        string
	Address2       string
	City           string
	State          string
	PostalCode     string
	Country        string
	Website        string
	Comment        string

	// --------------- User ---------------
	Points                    int64
	DateofBirth               rlib.JSONDate
	EmergencyContactName      string
	EmergencyContactAddress   string
	EmergencyContactTelephone string
	EmergencyContactEmail     string
	AlternateEmailAddress     string
	EligibleFutureUser        bool
	Industry                  int64
	SourceSLSID               int64

	// --------------- Payor -------------------
	CreditLimit         float64
	TaxpayorID          string
	GrossIncome         float64
	DriversLicense      string
	EligibleFuturePayor bool

	// -------------- Prospect ----------------
	CompanyAddress           string
	CompanyCity              string
	CompanyState             string
	CompanyPostalCode        string
	CompanyEmail             string
	CompanyPhone             string
	Occupation               string
	CurrentAddress           string
	CurrentLandLordName      string
	CurrentLandLordPhoneNo   string
	CurrentReasonForMoving   int64
	CurrentLengthOfResidency string
	PriorAddress             string
	PriorLandLordName        string
	PriorLandLordPhoneNo     string
	PriorReasonForMoving     int64
	PriorLengthOfResidency   string
	EvictedDes               string
	ConvictedDes             string
	BankruptcyDes            string
	OtherPreferences         string // arbitrary text
	// FollowUpDate             rlib.JSONDate // automatically fill out this date to sysdate + 24hrs
	// CommissionableThirdParty string
	SpecialNeeds     string // special needs for potential renters who are disabled
	Evicted          bool
	Convicted        bool
	Bankruptcy       bool
	FLAGS            uint64 // 0 = Approved/NotApproved,
	ThirdPartySource string
	CreateTS         rlib.JSONDateTime
	CreateBy         int64
	LastModTime      rlib.JSONDateTime
	LastModBy        int64
}

// RPersonOther contains the data from selections boxes in the UI. These come back
// in structure form rather than as a single string value.
type RPersonOther struct {
	State               string
	CompanyState        string
	EligibleFutureUser  bool
	EligibleFuturePayor bool
}

// GetTransactantResponse is the response data to requests to get a transactant
type GetTransactantResponse struct {
	Status string      `json:"status"`
	Record RPersonForm `json:"record"`
}

// SearchTransactantsResponse is the data structure for the response to a search for people
type SearchTransactantsResponse struct {
	Status  string             `json:"status"`
	Total   int64              `json:"total"`
	Records []rlib.Transactant `json:"records"`
}

// TransactantsTypedownResponse is the data structure for the response to a search for people
type TransactantsTypedownResponse struct {
	Status  string                     `json:"status"`
	Total   int64                      `json:"total"`
	Records []rlib.TransactantTypeDown `json:"records"`
}

// TransactantsTypedownDetailsResponse is the data structure for the response to a search for people
type TransactantsTypedownDetailsResponse struct {
	Status  string                            `json:"status"`
	Total   int64                             `json:"total"`
	Records []rlib.TransactantDetailsTypeDown `json:"records"`
}

// DeletePersonForm holds ARID to delete it
type DeletePersonForm struct {
	TCID int64
}

// SvcTransactantTypeDown handles typedown requests for Transactants.  It returns
// FirstName, LastName, and TCID
// wsdoc {
//  @Title  Get Transactants Typedown
//	@URL /v1/transactanttd/:BUI?request={"search":"The search string","max":"Maximum number of return items"}
//	@Method GET
//	@Synopsis Fast Search for Transactants matching typed characters
//  @Desc Returns TCID, FirstName, Middlename, and LastName of Transactants that
//  @Desc match supplied chars at the beginning of FirstName or LastName
//  @Input WebTypeDownRequest
//  @Response TransactantsTypedownResponse
// wsdoc }
func SvcTransactantTypeDown(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcTransactantTypeDown"
	var (
		g   TransactantsTypedownResponse
		err error
	)
	rlib.Console("Entered %s\n", funcname)

	rlib.Console("handle typedown: GetTransactantsTypeDown( bid=%d, search=%s, limit=%d\n", d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	g.Records, err = rlib.GetTransactantTypeDown(r.Context(), d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	rlib.Console("GetTransactantTypeDown returned %d matches\n", len(g.Records))
	g.Total = int64(len(g.Records))
	if err != nil {
		e := fmt.Errorf("Error getting typedown matches: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	for i := 0; i < len(g.Records); i++ {
		g.Records[i].Recid = int64(i)
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// SvcTransactantDetailsTypeDown handles typedown requests for Transactants.  It returns
// FirstName, LastName, and TCID
// wsdoc {
//  @Title  Get Transactants Details Typedown
//	@URL /v1/transactantdettd/:BUI?request={"search":"The search string","max":"Maximum number of return items"}
//	@Method GET
//	@Synopsis Fast Search for Transactants matching typed characters
//  @Desc Returns TCID, FirstName, Middlename, and LastName of Transactants that
//  @Desc match supplied chars at the beginning of FirstName or LastName
//  @Input WebTypeDownRequest
//  @Response TransactantsTypedownResponse
// wsdoc }
func SvcTransactantDetailsTypeDown(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcTransactantDetailsTypeDown"
	var g TransactantsTypedownDetailsResponse
	var dispName string
	n := d.wsTypeDownReq.Search

	rlib.Console("Entered %s - name - %s\n", funcname, n)
	q := fmt.Sprintf(`
SELECT
	A.BID,
    A.FirstName,
    A.MiddleName,
    A.LastName,
    A.CompanyName,
    A.DispName,
    A.IsCompany,
    A.PrimaryEmail,
    A.SecondaryEmail,
    A.WorkPhone,
    A.CellPhone,
    A.Address,
    A.Address2,
    A.City,
    A.State,
    A.PostalCode
FROM
	(
		( SELECT TCID,
			BID,
			FirstName,
			MiddleName,
			LastName,
			CompanyName,
            CONCAT( FirstName, ' ', MiddleName, ' ', LastName) AS DispName,
			IsCompany,
			PrimaryEmail,
			SecondaryEmail,
			WorkPhone,
			CellPhone,
			Address,
			Address2,
			City,
			State,
			PostalCode
		FROM Transactant
		WHERE BID=%d AND IsCompany = 0 AND
			(
				FirstName LIKE '%s%%' OR
				MiddleName LIKE '%s%%' OR
				LastName LIKE '%s%%'
			)
		)
	UNION ALL
		( SELECT TCID,
			BID,
			FirstName,
			MiddleName,
			LastName,
			CompanyName,
            CompanyName AS DispName,
			IsCompany,
			PrimaryEmail,
			SecondaryEmail,
			WorkPhone,
			CellPhone,
			Address,
			Address2,
			City,
			State,
			PostalCode
		FROM Transactant
		WHERE BID=%d AND IsCompany > 0 AND
			(
				CompanyName LIKE '%s%%'
			)
		)
    ) A
WHERE
    A.BID=1
ORDER BY A.DispName ASC
LIMIT 200;`,
		d.BID, n, n, n, d.BID, n)
	rlib.Console("%s: q = %q\n", funcname, q)

	rows, err := rlib.RRdb.Dbrr.Query(q)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var a rlib.TransactantDetailsTypeDown
		err = rows.Scan(
			&a.TCID,
			&a.FirstName,
			&a.MiddleName,
			&a.LastName,
			&a.CompanyName,
			&dispName,
			&a.IsCompany,
			&a.PrimaryEmail,
			&a.SecondaryEmail,
			&a.WorkPhone,
			&a.CellPhone,
			&a.Address,
			&a.Address2,
			&a.City,
			&a.State,
			&a.PostalCode)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		a.Recid = int64(i)
		g.Records = append(g.Records, a)
	}

	g.Status = "success"
	g.Total = int64(len(g.Records))
	SvcWriteResponse(d.BID, &g, w)
}

// fields list needs to be fetched for grid
var transactantGridFieldsMap = map[string][]string{
	"TCID":           {"Transactant.TCID"},
	"BID":            {"Transactant.BID"},
	"NLID":           {"Transactant.NLID"},
	"FirstName":      {"Transactant.FirstName"},
	"MiddleName":     {"Transactant.MiddleName"},
	"LastName":       {"Transactant.LastName"},
	"PreferredName":  {"Transactant.PreferredName"},
	"CompanyName":    {"Transactant.CompanyName"},
	"IsCompany":      {"Transactant.IsCompany"},
	"PrimaryEmail":   {"Transactant.PrimaryEmail"},
	"SecondaryEmail": {"Transactant.SecondaryEmail"},
	"WorkPhone":      {"Transactant.WorkPhone"},
	"CellPhone":      {"Transactant.CellPhone"},
	"Address":        {"Transactant.Address"},
	"Address2":       {"Transactant.Address2"},
	"City":           {"Transactant.City"},
	"State":          {"Transactant.State"},
	"PostalCode":     {"Transactant.PostalCode"},
	"Country":        {"Transactant.Country"},
	"Website":        {"Transactant.Website"},
	"LastModTime":    {"Transactant.LastModTime"},
	"LastModBy":      {"Transactant.LastModBy"},
	"CreateTS":       {"Transactant.CreateTS"},
	"CreateBy":       {"Transactant.CreateBy"},
}

var transactantSelectFields = []string{
	"Transactant.TCID",
	"Transactant.BID",
	"Transactant.NLID",
	"Transactant.FirstName",
	"Transactant.MiddleName",
	"Transactant.LastName",
	"Transactant.PreferredName",
	"Transactant.CompanyName",
	"Transactant.IsCompany",
	"Transactant.PrimaryEmail",
	"Transactant.SecondaryEmail",
	"Transactant.WorkPhone",
	"Transactant.CellPhone",
	"Transactant.Address",
	"Transactant.Address2",
	"Transactant.City",
	"Transactant.State",
	"Transactant.PostalCode",
	"Transactant.Country",
	"Transactant.Website",
	"Transactant.LastModTime",
	"Transactant.LastModBy",
	"Transactant.CreateTS",
	"Transactant.CreateBy",
}

// transactantRowScan scans a result from sql row and dump it in a rlib.Transactant struct
func transactantRowScan(rows *sql.Rows, t rlib.Transactant) (rlib.Transactant, error) {
	err := rows.Scan(&t.TCID, &t.BID, &t.NLID, &t.FirstName, &t.MiddleName, &t.LastName, &t.PreferredName, &t.CompanyName, &t.IsCompany, &t.PrimaryEmail, &t.SecondaryEmail, &t.WorkPhone, &t.CellPhone, &t.Address, &t.Address2, &t.City, &t.State, &t.PostalCode, &t.Country, &t.Website, &t.LastModTime, &t.LastModBy, &t.CreateTS, &t.CreateBy)
	return t, err
}

// SvcSearchHandlerTransactants handles the search query for Transactants from the Transactant Grid.
// wsdoc {
//  @Title  Search Transactants
//	@URL /v1/transactants/:BUI
//	@Method POST
//	@Synopsis Search transactants
//  @Description Returns a list of Transactants matching the search criteria
//  @Input WebGridSearchRequest
//  @Response SearchTransactantsResponse
// wsdoc }
func SvcSearchHandlerTransactants(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerTransactants"
	var (
		err error
		g   SearchTransactantsResponse
	)
	rlib.Console("Entered %s\n", funcname)

	const (
		limitClause int = 100
	)

	srch := fmt.Sprintf("Transactant.BID=%d", d.BID)               // default WHERE clause
	order := "Transactant.LastName ASC, Transactant.FirstName ASC" // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, transactantGridFieldsMap)
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	// Transactant Query Text Template
	transactantsQuery := `
	SELECT
		{{.SelectClause}}
	FROM Transactant
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}` // don't add ';', later some parts will be added in query

	// will be substituted as query clauses
	qc := rlib.QueryClause{
		"SelectClause": strings.Join(transactantSelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
	}

	// GET TOTAL COUNTS of query
	countQuery := rlib.RenderSQLQuery(transactantsQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery) // total number of rows that match the criteria
	if err != nil {
		rlib.Console("Error from rlib.GetQueryCount: %s\n", err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("g.Total = %d\n", g.Total)

	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`

	// build query with limit and offset clause
	// if query ends with ';' then remove it
	transactantsQueryWithLimit := transactantsQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(transactantsQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var t rlib.Transactant
		t.Recid = i

		// get record of transactant
		t, err = transactantRowScan(rows, t)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		g.Records = append(g.Records, t)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}
	// error check
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// write response
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}

// SvcFormHandlerXPerson formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the TCID as follows:
//       0    1       2    3
// 		/v1/xperson/BID/TCID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcFormHandlerXPerson"
	var (
		err error
	)
	rlib.Console("Entered %s\n", funcname)

	if d.TCID, err = SvcExtractIDFromURI(r.RequestURI, "TCID", 3, w); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Request: %s:  BID = %d,  TCID = %d\n", d.wsSearchReq.Cmd, d.BID, d.TCID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getXPerson(w, r, d)
		break
	case "save":
		saveXPerson(w, r, d)
		break
	case "delete":
		deleteXPerson(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

func xpUpdatePerson(w http.ResponseWriter, r *http.Request, xp *rlib.XPerson) bool {
	funcname := "xpUpdatePerson"
	// rlib.Console("Entered %s\n", funcname)
	err := rlib.UpdateTransactant(r.Context(), &xp.Trn)
	if err != nil {
		e := fmt.Errorf("%s: UpdateTransactant error:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return true
	}

	err = rlib.UpdateUser(r.Context(), &xp.Usr)
	if err != nil {
		e := fmt.Errorf("%s: UpdateUser error:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return true
	}

	err = rlib.UpdateProspect(r.Context(), &xp.Psp)
	if err != nil {
		e := fmt.Errorf("%s: UpdateProspect error:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return true
	}

	// rlib.Console("Calling UpdatePayor, xp.Pay.CreditLimit = %6.2f\n", xp.Pay.CreditLimit)
	err = rlib.UpdatePayor(r.Context(), &xp.Pay)
	if err != nil {
		e := fmt.Errorf("%s: UpdatePayor err.Pay %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return true
	}
	// rlib.Console("successfully updated payor\n")
	return false

}

func xpInsertPerson(w http.ResponseWriter, r *http.Request, xp *rlib.XPerson) bool {
	funcname := "xpInsertPerson"
	tcid, err := rlib.InsertTransactant(r.Context(), &xp.Trn)
	if err != nil {
		e := fmt.Errorf("%s: Insert Transactant error:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return true
	}

	errlist := bizlogic.FinalizeTransactant(r.Context(), &xp.Trn)
	if len(errlist) > 0 {
		SvcErrListReturn(w, errlist, funcname)
		return true
	}

	// update tcid in user, prospect, payor struct
	xp.Usr.TCID = tcid
	xp.Pay.TCID = tcid
	xp.Psp.TCID = tcid

	_, err = rlib.InsertUser(r.Context(), &xp.Usr)
	if err != nil {
		e := fmt.Errorf("%s: Insert User error:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return true
	}

	_, err = rlib.InsertProspect(r.Context(), &xp.Psp)
	if err != nil {
		e := fmt.Errorf("%s: Insert Prospect error:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return true
	}

	_, err = rlib.InsertPayor(r.Context(), &xp.Pay)
	if err != nil {
		e := fmt.Errorf("%s: Insert Payor error:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return true
	}
	return false
}

// saveXPerson handles the Save action from the Transactant Form
// wsdoc {
//  @Title  Save Transactant
//	@URL /v1/person/:BUI/:TCID
//	@Method POST
//	@Synopsis Save a Transactant
//  @Description
//  @Input WebGridSearchRequest
//  @Response SearchTransactantsResponse
// wsdoc }
func saveXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveXPerson"
	var (
		err error
	)

	target := `"record":`
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)
	i := strings.Index(d.data, target)
	rlib.Console("record is at index = %d\n", i)
	if i < 0 {
		e := fmt.Errorf("saveXPerson: cannot find %s in form json", target)
		SvcErrorReturn(w, e, funcname)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	rlib.Console("data to unmarshal is:  %s\n", s)

	//===============================================================
	//------------------------------
	// Handle all the non-list data
	//------------------------------
	var gxp RPersonForm
	var xp rlib.XPerson

	err = json.Unmarshal([]byte(s), &gxp)
	if err != nil {
		rlib.Console("Data unmarshal error: %s\n", err.Error())
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	rlib.Console("saveXPersonL Start migration\n")
	rlib.MigrateStructVals(&gxp, &xp.Trn)
	rlib.MigrateStructVals(&gxp, &xp.Usr)
	rlib.MigrateStructVals(&gxp, &xp.Psp)
	rlib.MigrateStructVals(&gxp, &xp.Pay)
	rlib.Console("saveXPersonL Finished migration\n")
	rlib.Console("CreditLimit = %6.2f, TaxpayorID = %s, GrossIncome = %6.2f, DriversLicense = %s\n", xp.Pay.CreditLimit, xp.Pay.TaxpayorID, xp.Pay.GrossIncome, xp.Pay.DriversLicense)
	//---------------------------
	// Handle all the list data
	//---------------------------
	var gxpo RPersonOther
	err = json.Unmarshal([]byte(s), &gxpo)
	if err != nil {
		rlib.Console("Data unmarshal error: %s\n", err.Error())
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	xp.Trn.State = gxpo.State
	xp.Usr.EligibleFutureUser = gxpo.EligibleFutureUser
	xp.Psp.CompanyState = gxpo.CompanyState
	xp.Pay.EligibleFuturePayor = gxpo.EligibleFuturePayor

	// Manage Prospect FLAGS
	xp.Psp.FLAGS &= ^uint64(0x4 | 0x2 | 0x1) // 1<<0 and 1<< 1 and 1<<2:  these are the three flags that can be set.  Assume we turn them off. Set all bits to 0.
	if gxp.Evicted {
		// mask: 1 << 0
		xp.Psp.FLAGS |= 0x1
	}

	if gxp.Convicted {
		// mask: 1 << 1
		xp.Psp.FLAGS |= 0x2
	}

	if gxp.Bankruptcy {
		// mask: 1 << 2
		xp.Psp.FLAGS |= 0x4
	}

	//===============================================================
	// save or update
	//===============================================================
	if xp.Trn.TCID == 0 {
		// this is new transactant record
		fmt.Println(">>> Inserting New Transactant Record")
		if xpInsertPerson(w, r, &xp) {
			return
		}
	} else {
		rlib.Console("Updating Transactant record with TCID: %d\n", xp.Trn.TCID)
		if xpUpdatePerson(w, r, &xp) {
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// getXPerson handles the request for an XPerson from the Transactant Form
// wsdoc {
//  @Title  Get Transactant
//	@URL /v1/person/:BUI/:TCID
//	@Method POST
//	@Synopsis Read a Transactant
//  @Description
//  @Input WebGridSearchRequest
//  @Response GetTransactantResponse
// wsdoc }
func getXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		g            GetTransactantResponse
		xp           rlib.XPerson
		prospectFlag uint64
	)
	_ = rlib.GetXPerson(r.Context(), d.TCID, &xp)
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
	g.Record.BID = d.BID
	g.Record.BUD = rlib.GetBUDFromBIDList(d.BID)

	// Manage "Have you ever been"(Prospect) section FLAGS
	prospectFlag = xp.Psp.FLAGS
	g.Record.Evicted = prospectFlag&0x1 != 0
	g.Record.Convicted = prospectFlag&0x2 != 0
	g.Record.Bankruptcy = prospectFlag&0x4 != 0

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// deleteXPerson request to delete Person with TCID from database
// wsdoc {
//  @Title  Delete Transactant, User, Prospect, Payor
//	@URL /v1/person/:BUI/:TCID
//  @Method  DELETE
//	@Synopsis Delete record for a Person
//  @Description  Delete record from database :TCID
//	@Input DeletePersonForm
//  @Response SvcWriteSuccessResponse
// wsdoc }
func deleteXPerson(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteXPerson"
	var (
		del DeletePersonForm
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	if err := bizlogic.DeleteTransactant(r.Context(), d.BID, del.TCID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponse(d.BID, w)
}
