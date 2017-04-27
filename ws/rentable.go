package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// Decoding the form data from w2ui gets tricky when certain value types are returned.
// For example, dropdown menu selections are returned as a JSON struct value
//     "AssignmentTime": { "ID": "Pre-Assign", "Text": "Pre-Assign"}
// The approach to getting this sort of information back into the appropriate struct
// is to:
//		1. Use MigrateStructVals to get pretty much everything except
//         dropdown menu selections.
//		2. Handle the dropdown menu selections separately using rlib.W2uiHTMLSelect
//         for unmarshaling

// RentableForm is a structure specifically for the UI. It will be
// automatically populated from an rlib.Rentable struct
type RentableForm struct {
	Recid        int64 `json:"recid"` // this is to support the w2ui form
	RID          int64
	RentableName string
	LastModTime  rlib.JSONTime
	LastModBy    int64
}

// RentableOther is a struct to handle the UI list box selections
type RentableOther struct {
	BID            rlib.W2uiHTMLSelect
	AssignmentTime rlib.W2uiHTMLSelect
}

// PrRentableOther is a structure specifically for the UI. It will be
// automatically populated from an rlib.Rentable struct
type PrRentableOther struct {
	Recid                int64 `json:"recid"` // this is to support the w2ui form
	RID                  int64
	BID                  rlib.XJSONBud
	RTID                 int64
	RentableName         string
	RentableType         string
	RentableStatus       string
	RAID                 int64
	RentalAgreementStart rlib.JSONTime
	RentalAgreementStop  rlib.JSONTime
	AssignmentTime       rlib.XJSONAssignmentTime
	LastModTime          rlib.JSONTime
	LastModBy            int64
}

// SearchRentablesResponse is a response string to the search request for rentables
type SearchRentablesResponse struct {
	Status  string            `json:"status"`
	Total   int64             `json:"total"`
	Records []PrRentableOther `json:"records"`
}

// GetRentableResponse is the response to a GetRentable request
type GetRentableResponse struct {
	Status string          `json:"status"`
	Record RentableDetails `json:"record"`
}

// RentableTypedownResponse is the data structure for the response to a search for people
type RentableTypedownResponse struct {
	Status  string                  `json:"status"`
	Total   int64                   `json:"total"`
	Records []rlib.RentableTypeDown `json:"records"`
}

// RentableDetails holds the details about other detailed associated data with specific rentable
type RentableDetails struct {
	BID            rlib.XJSONBud // business
	RID            int64
	RentableName   string        // Rentable Name
	RAID           int64         // Rental Agreement ID for this period
	RARDtStart     rlib.JSONTime // RentalAgreementStart Date
	RARDtStop      rlib.JSONTime // RentalAgreementStop Date
	RTID           int64         // Rentable type id
	RTRefDtStart   rlib.JSONTime // Rentable Type Reference Stop Date
	RTRefDtStop    rlib.JSONTime // Rentable Type Reference Start Date
	RentableType   string        // Rentable Type Name
	RentableStatus string        // rentable status
	RSDtStart      rlib.JSONTime // rentable status start date
	RSDtStop       rlib.JSONTime // rentable status stop date
	CurrentDate    rlib.JSONTime
}

// SvcRentableTypeDown handles typedown requests for Rentables.  It returns
// FirstName, LastName, and TCID
// wsdoc {
//  @Title  Get Rentables Typedown
//	@URL /v1/Rentabletd/:BUI?request={"search":"The search string","max":"Maximum number of return items"}
//	@Method GET
//	@Synopsis Fast Search for Rentables matching typed characters
//  @Desc Returns TCID, FirstName, Middlename, and LastName of Rentables that
//  @Desc match supplied chars at the beginning of FirstName or LastName
//  @Input WebTypeDownRequest
//  @Response RentablesTypedownResponse
// wsdoc }
func SvcRentableTypeDown(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var g RentableTypedownResponse
	var err error
	fmt.Printf("handle typedown: GetRentablesTypeDown( bid=%d, search=%s, limit=%d\n", d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	g.Records, err = rlib.GetRentableTypeDown(d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	if err != nil {
		SvcGridErrorReturn(w, fmt.Errorf("Error getting typedown matches: %s", err.Error()))
		return
	}
	fmt.Printf("GetRentableTypeDown returned %d matches\n", len(g.Records))
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(&g, w)
}

// rentablesGridFields holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
var rentablesGridFieldsMap = map[string][]string{
	"RID":                  {"Rentable.RID"},
	"RentableName":         {"Rentable.RentableName"},
	"RentableType":         {"RentableTypes.Name"},
	"RTID":                 {"RentableTypes.RTID"},
	"RentableStatus":       {"RentableStatus.Status"},
	"RAID":                 {"RentalAgreementRentables.RAID"},
	"RentalAgreementStart": {"RentalAgreementRentables.RARDtStart"},
	"RentalAgreementStop":  {"RentalAgreementRentables.RARDtStop"},
}

// which fields needs to be fetched for SQL query for rentables
var rentablesQuerySelectFields = []string{
	"Rentable.RID",
	"Rentable.RentableName",
	"RentableTypes.Name as RentableType",
	"RentableTypes.RTID",
	"RentableStatus.Status as RentableStatus",
	"RentalAgreementRentables.RAID",
	"RentalAgreementRentables.RARDtStart as RentalAgreementStart",
	"RentalAgreementRentables.RARDtStop as RentalAgreementStop",
}

// rentablesRowScan scans a result from sql row and dump it in a PrRentableOther struct
func rentablesRowScan(rows *sql.Rows, q PrRentableOther) PrRentableOther {
	var rStatus int64

	rlib.Errcheck(rows.Scan(&q.RID, &q.RentableName, &q.RentableType, &q.RTID, &rStatus, &q.RAID, &q.RentalAgreementStart, &q.RentalAgreementStop))

	// convert status int to string, human readable
	q.RentableStatus = rlib.RentableStatusToString(rStatus)

	return q
}

// SvcSearchHandlerRentables generates a report of all Rentables defined business d.BID
// wsdoc {
//  @Title  Search Rentables
//	@URL /v1/rentables/:BUI
//  @Method  POST
//	@Synopsis Search Rentables
//  @Description  Search all Rentables and return those that match the Search Logic
//	@Input WebGridSearchRequest
//  @Response SearchRentablesResponse
// wsdoc }
func SvcSearchHandlerRentables(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	fmt.Printf("Entered SvcSearchHandlerRentables\n")

	var (
		err error
		g   SearchRentablesResponse
		t   = time.Now()
	)

	// default search (where clause) and sort (order by clause)
	srch := fmt.Sprintf("Rentable.BID=%d AND RentalAgreementRentables.RARDtStop>%q", d.BID, t.Format(rlib.RRDATEINPFMT)) // default WHERE clause
	order := "Rentable.RentableName ASC"                                                                                 // default ORDER

	// check that RentableStatus is there in search fields
	// if exists then modify it
	var rStatusSearch []GenSearch
	for i := 0; i < len(d.wsSearchReq.Search); i++ {
		if d.wsSearchReq.Search[i].Field == "RentableStatus" {
			for index, status := range rlib.RentableStatusString {
				if strings.Contains(status, strings.ToLower(d.wsSearchReq.Search[i].Value)) && strings.TrimSpace(d.wsSearchReq.Search[i].Value) != "" {
					rStatusSearch = append(rStatusSearch, GenSearch{
						Operator: "is", Field: "RentableStatus", Value: rlib.IntToString(index), Type: "int",
					})
				}
			}
			// remove original rentable status from search
			d.wsSearchReq.Search = append(d.wsSearchReq.Search[:i], d.wsSearchReq.Search[i+1:]...)
		}
	}

	// append modified status search fields
	d.wsSearchReq.Search = append(d.wsSearchReq.Search, rStatusSearch...)

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rentablesGridFieldsMap)
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	// Rentables Query Text Template
	rentablesQuery := `
	SELECT
		{{.SelectClause}}
	FROM Rentable
	INNER JOIN RentableTypeRef ON Rentable.RID=RentableTypeRef.RID
	INNER JOIN RentableTypes ON RentableTypeRef.RTID=RentableTypes.RTID
	INNER JOIN RentableStatus ON RentableStatus.RID=Rentable.RID
	INNER JOIN RentalAgreementRentables ON RentalAgreementRentables.RID=Rentable.RID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}};
	`

	// will be substituted as query clauses
	qc := queryClauses{
		"SelectClause": strings.Join(rentablesQuerySelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
	}

	// get formatted query with substitution of select, where, order clause
	q := renderSQLQuery(rentablesQuery, qc)
	fmt.Printf("db query = %s\n", q)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	// get records by iteration
	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q PrRentableOther
		q.Recid = i
		q.BID = rlib.XJSONBud(fmt.Sprintf("%d", d.BID))

		// get records in q struct
		q = rentablesRowScan(rows, q)

		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}
	// error check
	rlib.Errcheck(rows.Err())

	// get total count of results
	g.Total, err = GetQueryCount(q, qc)
	if err != nil {
		fmt.Printf("Error from GetQueryCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}
	fmt.Printf("g.Total = %d\n", g.Total)

	// write response
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// SvcFormHandlerRentable formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the RID as follows:
//           0    1         2   3
// uri 		/v1/rentable/BUD/RID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerRentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcFormHandlerRentable\n")

	var err error
	if d.RID, err = SvcExtractIDFromURI(r.RequestURI, "RID", 3, w); err != nil {
		return
	}

	fmt.Printf("Request: %s:  BID = %d,  RID = %d\n", d.wsSearchReq.Cmd, d.BID, d.RID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getRentable(w, r, d)
		break
	case "save":
		saveRentable(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err)
		return
	}
}

// saveRentable returns the requested rentable
// wsdoc {
//  @Title  Save Rentable
//	@URL /v1/rentable/:BUI/:RID
//  @Method  GET
//	@Synopsis Update the information on a Rentable with the supplied data
//  @Description  This service updates Rentable :RID with the information supplied. All fields must be supplied.
//	@Input WebGridSearchRequest
//  @Response SvcStatusResponse
// wsdoc }
func saveRentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	// funcname := "saveRentable"
	target := `"record":`
	// fmt.Printf("SvcFormHandlerRentable save\n")
	// fmt.Printf("record data = %s\n", d.data)
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("saveRentable: cannot find %s in form json", target)
		SvcGridErrorReturn(w, e)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	var foo RentableForm
	err := json.Unmarshal([]byte(s), &foo)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	// migrate the variables that transfer without needing special handling...
	var a rlib.Rentable
	rlib.MigrateStructVals(&foo, &a)

	// now get the stuff that requires special handling...
	var bar RentableOther
	err = json.Unmarshal([]byte(s), &bar)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.BID.ID]
	if !ok {
		e := fmt.Errorf("Could not map BID value: %s", bar.BID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	a.AssignmentTime, ok = rlib.AssignmentTimeMap[bar.AssignmentTime.ID]
	if !ok {
		e := fmt.Errorf("Could not map AssignmentTime value: %s", bar.AssignmentTime.ID)
		SvcGridErrorReturn(w, e)
		return
	}

	// Now just update the database
	err = rlib.UpdateRentable(&a)
	if err != nil {
		e := fmt.Errorf("Error updating rentable: %s", err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	SvcWriteSuccessResponse(w)
}

// which fields needs to be fetched for SQL query for rentable details
var rentableFormSelectFields = []string{
	"Rentable.RID",
	"Rentable.RentableName",
	"RentalAgreementRentables.RAID",
	"RentalAgreementRentables.RARDtStart",
	"RentalAgreementRentables.RARDtStop",
	"RentableTypeRef.RTID",
	"RentableTypeRef.DtStart as RTRefDtStart",
	"RentableTypeRef.DtStop as RTRefDtStop",
	"RentableTypes.Name",
	"RentableStatus.Status as RentableStatus",
	"RentableStatus.DtStart as RSDtStart",
	"RentableStatus.DtStop as RSDtStop",
}

// getRentable returns the requested rentable
// wsdoc {
//  @Title  Get Rentable
//	@URL /v1/rentable/:BUI/:RID
//  @Method  GET
//	@Synopsis Get information on a Rentable
//  @Description  Return all fields for rentable :RID
//	@Input WebGridSearchRequest
//  @Response GetRentableResponse
// wsdoc }
func getRentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered getRentable\n")
	var (
		g GetRentableResponse
		t = time.Now()
	)

	rentableQuery := `
	SELECT
		{{.SelectClause}}
	FROM Rentable
	INNER JOIN RentableTypeRef ON Rentable.RID = RentableTypeRef.RID
	INNER JOIN RentableTypes ON RentableTypeRef.RTID=RentableTypes.RTID
	INNER JOIN RentableStatus ON RentableStatus.RID=Rentable.RID
	INNER JOIN RentalAgreementRentables ON RentalAgreementRentables.RID=Rentable.RID
	WHERE {{.WhereClause}};
	`

	// will be substituted as query clauses
	qc := queryClauses{
		"SelectClause": strings.Join(rentableFormSelectFields, ","),
		"WhereClause":  fmt.Sprintf("Rentable.BID=%d AND Rentable.RID=%d AND RentalAgreementRentables.RARDtStop>%q", d.BID, d.RID, t.Format(rlib.RRDATEINPFMT)),
	}

	// get formatted query with substitution of select, where, order clause
	q := renderSQLQuery(rentableQuery, qc)
	fmt.Printf("db query = %s\n", q)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var gg RentableDetails
		gg.CurrentDate = rlib.JSONTime(t)

		for bud, bid := range rlib.RRdb.BUDlist {
			if bid == d.BID {
				gg.BID = rlib.XJSONBud(bud)
				break
			}
		}

		var rStatus int64
		rows.Scan(&gg.RID, &gg.RentableName, &gg.RAID, &gg.RARDtStart, &gg.RARDtStop, &gg.RTID, &gg.RTRefDtStart, &gg.RTRefDtStop, &gg.RentableType, &rStatus, &gg.RSDtStart, &gg.RSDtStop)

		// convert status int to string, human readable
		gg.RentableStatus = rlib.RentableStatusToString(rStatus)

		g.Record = gg
	}
	// error check
	rlib.Errcheck(rows.Err())

	// write response
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
