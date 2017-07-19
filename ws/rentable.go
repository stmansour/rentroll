package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
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
	LastModTime  rlib.JSONDate
	LastModBy    int64
}

// RentableOther is a struct to handle the UI list box selections
type RentableOther struct {
	BID rlib.W2uiHTMLSelect
	// AssignmentTime rlib.W2uiHTMLSelect
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
	RARID                rlib.NullInt64
	RAID                 rlib.NullInt64
	RentalAgreementStart rlib.NullTime
	RentalAgreementStop  rlib.NullTime
	// AssignmentTime       rlib.XJSONAssignmentTime
	LastModTime rlib.JSONDate
	LastModBy   int64
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
	Recid          int64         `json:"recid"` // this is to support the w2ui form
	BID            rlib.XJSONBud // business
	RID            int64
	RentableName   string         // Rentable Name
	RARID          rlib.NullInt64 // RentalAgreementRentable ID
	RAID           rlib.NullInt64 // Rental Agreement ID for this period
	RARDtStart     rlib.NullTime  // RentalAgreementStart Date
	RARDtStop      rlib.NullTime  // RentalAgreementStop Date
	RTID           int64          // Rentable type id
	RTRID          int64          // Rentable Type Reference ID
	RTRefDtStart   rlib.JSONDate  // Rentable Type Reference Stop Date
	RTRefDtStop    rlib.JSONDate  // Rentable Type Reference Start Date
	RentableType   string         // Rentable Type Name
	RSID           int64          // Rentable Status ID
	RentableStatus string         // rentable status
	RSDtStart      rlib.JSONDate  // rentable status start date
	RSDtStop       rlib.JSONDate  // rentable status stop date
	CurrentDate    rlib.JSONDate
	AssignmentTime int64 // assignment time
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
	var (
		funcname = "SvcRentableTypeDown"
		g        RentableTypedownResponse
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)

	fmt.Printf("handle typedown: GetRentablesTypeDown( bid=%d, search=%s, limit=%d\n", d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	g.Records, err = rlib.GetRentableTypeDown(d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	if err != nil {
		e := fmt.Errorf("Error getting typedown matches: %s", err.Error())
		SvcGridErrorReturn(w, e, funcname)
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
	"RARID":                {"RentalAgreementRentables.RARID"},
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
	"RentalAgreementRentables.RARID",
	"RentalAgreementRentables.RAID",
	"RentalAgreementRentables.RARDtStart as RentalAgreementStart",
	"RentalAgreementRentables.RARDtStop as RentalAgreementStop",
}

// rentablesRowScan scans a result from sql row and dump it in a PrRentableOther struct
func rentablesRowScan(rows *sql.Rows, q PrRentableOther) (PrRentableOther, error) {
	var rStatus int64

	err := rows.Scan(&q.RID, &q.RentableName, &q.RentableType, &q.RTID, &rStatus, &q.RARID, &q.RAID, &q.RentalAgreementStart, &q.RentalAgreementStop)

	// convert status int to string, human readable
	q.RentableStatus = rlib.RentableStatusToString(rStatus)

	return q, err
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

	var (
		funcname = "SvcSearchHandlerRentables"
		err      error
		g        SearchRentablesResponse
		// currentTime = time.Now()
	)
	fmt.Printf("Entered %s\n", funcname)

	// fetch records from the database under this limit
	const (
		limitClause int = 100
	)

	// default search (where clause) and sort (order by clause)
	// defaultWhere := `Rentable.BID=%d
	// 	AND (RentalAgreementRentables.RARDtStart<=%q OR RentalAgreementRentables.RARDtStart IS NULL)
	// 	AND (RentalAgreementRentables.RARDtStop>%q OR RentalAgreementRentables.RARDtStop IS NULL)
	// 	AND (RentableTypeRef.DtStart<=%q OR RentableTypeRef.DtStart IS NULL)
	// 	AND (RentableTypeRef.DtStop>%q OR RentableTypeRef.DtStop IS NULL)
	// 	AND (RentableStatus.DtStart<=%q OR RentableStatus.DtStart IS NULL)
	// 	AND (RentableStatus.DtStop>%q OR RentableStatus.DtStop IS NULL)`
	// srch := fmt.Sprintf(defaultWhere, d.BID, currentTime, currentTime, currentTime, currentTime, currentTime, currentTime)  // default WHERE clause

	// Show All Renbles no matter in what state they are,
	srch := fmt.Sprintf(`Rentable.BID=%d`, d.BID)
	// show active rentable first by RenalAgreement Dates
	// order := "RentalAgreementRentables.RARDtStop DESC, RentalAgreementRentables.RARDtStart DESC, Rentable.RentableName ASC" // default ORDER
	order := "Rentable.RID ASC,RentalAgreementRentables.RARID ASC" // default ORDER

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
	SELECT DISTINCT
		{{.SelectClause}}
	FROM Rentable
	INNER JOIN RentableTypeRef ON Rentable.RID=RentableTypeRef.RID
	INNER JOIN RentableTypes ON RentableTypeRef.RTID=RentableTypes.RTID
	INNER JOIN RentableStatus ON RentableStatus.RID=Rentable.RID
	LEFT JOIN RentalAgreementRentables ON RentalAgreementRentables.RID=Rentable.RID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}` // don't add ';', later some parts will be added in query

	// will be substituted as query clauses
	qc := queryClauses{
		"SelectClause": strings.Join(rentablesQuerySelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
	}

	// GET TOTAL COUNT OF RESULTS
	countQuery := renderSQLQuery(rentablesQuery, qc)
	g.Total, err = GetQueryCount(countQuery, qc)
	if err != nil {
		fmt.Printf("Error from GetQueryCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("g.Total = %d\n", g.Total)

	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`

	// build query with limit and offset clause
	// if query ends with ';' then remove it
	rentablesQueryWithLimit := rentablesQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := renderSQLQuery(rentablesQueryWithLimit, qc)
	fmt.Printf("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	// get records by iteration
	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q PrRentableOther
		q.Recid = i
		q.BID = rlib.XJSONBud(fmt.Sprintf("%d", d.BID))

		// get records in q struct
		q, err = rentablesRowScan(rows, q)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}
	// error check
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

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

	var (
		funcname = "SvcFormHandlerRentable"
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)

	if d.RID, err = SvcExtractIDFromURI(r.RequestURI, "RID", 3, w); err != nil {
		SvcGridErrorReturn(w, err, funcname)
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
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

// func dumpRTRList(m []rlib.RentableTypeRef) {
// 	for i := 0; i < len(m); i++ {
// 		fmt.Printf("m[%d] range = %s - %s\n", i, m[i].DtStart.Format(rlib.RRDATEINPFMT), m[i].DtStop.Format(rlib.RRDATEINPFMT))
// 	}
// 	fmt.Printf("----------------------------\n")
// }

// AdjustRTRTimeList determines what edits and/or inserts are needed to
// add the supplied rtr struct to the the existing RentableTypeRef records.
// Records are added as needed except where there is overlap. Overlaps are
// handled as illustrated in the following example (Rentable Types RT1 and
// RT2 are just examples)
//
// Example 1: Overlap similar type:
//
//            existing RTRs    new rtr         Result
//  t0 ----                    begin RT1    begin RT1
//                             |            |
//  t1 ----   begin RT1        |            |
//            |                |            |
//  t2 ----   end RT1          |            |
//                             |            |
//  t3 ----                    end RT1      end RT1
//
// Example 2: Overlap different types:
//
//            existing RTRs    new rtr         Result
//  t0 ----                    begin RT2    begin RT2    t0-t1 RT2
//                             |            end RT2
//  t1 ----   begin RT1        |            begin RT1    t1-t2 RT1
//            |                |            |
//  t2 ----   end RT1          |            end RT1
//                             |            begin RT2    t2-t3 RT2
//  t3 ----                    end RT2      end RT2
//
//
// @returns
//	1. existing array of RTRs  (these will need to be deleted)
//	2. the new set of RTRs     (these will need to be inserted)
func AdjustRTRTimeList(rtr *rlib.RentableTypeRef, r *rlib.Rentable) ([]rlib.RentableTypeRef, []rlib.RentableTypeRef) {
	var m []rlib.RentableTypeRef
	R := rlib.GetRentableTypeRefs(r.RID)
	l := len(R)
	rtrAdded := false // flag to mark whether rtr still needs to be added after loop
	for i := 0; i < l; i++ {
		if !rtrAdded && rlib.DateRangeOverlap(&rtr.DtStart, &rtr.DtStop, &R[i].DtStart, &R[i].DtStop) {
			if rtr.RTID == R[i].RTID { // same rentable type?
				if rtr.DtStart.After(R[i].DtStart) { // adjust start time to the earliest
					rtr.DtStart = R[i].DtStart
				}
				if rtr.DtStop.Before(R[i].DtStop) { // adjust stop time to the latest
					rtr.DtStop = R[i].DtStop
				}
			} else { // different types
				if rtr.DtStart.Equal(R[i].DtStart) && rtr.DtStop.Equal(R[i].DtStop) { // same date range, just a RentableType change
					rt := R[i]
					rt.RTID = rtr.RTID
					m = append(m, rt)
					rtrAdded = true
				} else if rtr.DtStart.Before(R[i].DtStart) { // if R[i] starts before rtr adjust rtr start point...
					rt := *rtr
					rt.DtStop = R[i].DtStart           // stop rt just as R[i] begins
					m = append(m, rt)                  // new version up to the point where R[i] starts
					m = append(m, R[i])                // R[i] stays as is
					rtr.DtStart = R[i].DtStop          // adjust rtr's new start point
					if rtr.DtStart.Equal(rtr.DtStop) { // are we finished?
						rtrAdded = true // we don't need to add rtr now
					}
					// fmt.Printf("AdjustRTRTimeList:  add period rtr start to R[%d] start, rtr.DtStop moved forward:  %s\n", i, rtr.DtStart.Format(rlib.RRDATEINPFMT))
					// dumpRTRList(m)
				} else if rtr.DtStart.After(R[i].DtStart) { // if rtr starts after R[i], adjust R[i] end time
					rt := R[i]              // start with a copy of rtr
					rt.DtStop = rtr.DtStart // stop this new one just as R[i] begins
					m = append(m, rt)
					if R[i].DtStop.After(rtr.DtStop) { // R[i] may last longer than rs.  If so, we need to cover the time after rs stops to when R[i] stops
						rt := R[i]
						rt.DtStart = rtr.DtStop
						m = append(m, rt)
					}
					// fmt.Printf("AdjustRTRTimeList:  different types append:  i = %d\n", i)
					// dumpRTRList(m)
				} else {
					// fmt.Printf("AdjustRTRTimeList: rtr is covered.  rtAdded set to true:  i = %d\n", i)
					rtrAdded = true
				}
			}
		} else { // the timespans do not overlap
			m = append(m, R[i]) // add this just as it is
		}
	}
	if !rtrAdded {
		m = append(m, *rtr) // add rtr to the list after all adjustments
	}
	return R, m
}

// AdjustRSTimeList - just like AdjustRTRTimeList except for RentableStatus records.  There's probably a better
// way to make both these functions into one.
func AdjustRSTimeList(rs *rlib.RentableStatus, r *rlib.Rentable) ([]rlib.RentableStatus, []rlib.RentableStatus) {
	var m []rlib.RentableStatus
	R := rlib.GetAllRentableStatus(r.RID)
	l := len(R)
	rsAdded := false // flag to mark whether rs still needs to be added after loop
	for i := 0; i < l; i++ {
		if !rsAdded && rlib.DateRangeOverlap(&rs.DtStart, &rs.DtStop, &R[i].DtStart, &R[i].DtStop) {
			if rs.Status == R[i].Status { // same rentable status?
				if rs.DtStart.After(R[i].DtStart) { // adjust start time to the earliest
					rs.DtStart = R[i].DtStart
				}
				if rs.DtStop.Before(R[i].DtStop) { // adjust stop time to the latest
					rs.DtStop = R[i].DtStop
				}
			} else { // different types
				if rs.DtStart.Equal(R[i].DtStart) && rs.DtStop.Equal(R[i].DtStop) { // same date range, just a status change
					rs1 := R[i]
					rs1.Status = rs.Status
					m = append(m, rs1)
					rsAdded = true
				} else if rs.DtStart.Before(R[i].DtStart) { // if R[i] starts before rs adjust rs start point...
					rs1 := *rs
					rs1.DtStop = R[i].DtStart        // stop rt just as R[i] begins
					m = append(m, rs1)               // new version up to the point where R[i] starts
					m = append(m, R[i])              // R[i] stays as is
					rs.DtStart = R[i].DtStop         // adjust rs's new start point
					if rs.DtStart.Equal(rs.DtStop) { // are we finished?
						rsAdded = true // we don't need to add rs now
					}
				} else if rs.DtStart.After(R[i].DtStart) { // if rs starts after R[i], adjust R[i] end time
					rs1 := R[i]             // start with a copy of rs
					rs1.DtStop = rs.DtStart // stop this new one just as R[i] begins
					m = append(m, rs1)
					if R[i].DtStop.After(rs.DtStop) { // R[i] may last longer than rs.  If so, we need to cover the time after rs stops to when R[i] stops
						rs1 := R[i]
						rs1.DtStart = rs.DtStop
						m = append(m, rs1)
					}
				} else {
					rsAdded = true
				}
			}
		} else { // the timespans do not overlap
			m = append(m, R[i]) // add this just as it is
		}
	}
	if !rsAdded {
		m = append(m, *rs) // add rs to the list after all adjustments
	}
	return R, m
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
	var (
		funcname = "saveRentable"
		err      error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	target := `"record":`
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("saveRentable: cannot find %s in form json", target)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]

	// rentable Form Record
	var rfRecord RentableDetails
	err = json.Unmarshal([]byte(s), &rfRecord)
	rlib.Errcheck(err)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	var (
		ok          bool
		rt          rlib.Rentable
		rs          rlib.RentableStatus
		rtr         rlib.RentableTypeRef
		currentTime = time.Now()
	)

	// checks for valid values
	requestedBID, ok := rlib.RRdb.BUDlist[string(rfRecord.BID)]
	if !ok {
		e := fmt.Errorf("Invalid Business ID found. BID: %s", rfRecord.BID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	// check whether rentable type is provided or not
	if !(rfRecord.RTID > 0) {
		e := fmt.Errorf("Rentable Type must be provided")
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	// // StopDate should not be before Today's date
	// if !(rlib.IsDateBefore((time.Time)(rfRecord.RTRefDtStart), (time.Time)(rfRecord.RTRefDtStop))) {
	// 	e := fmt.Errorf("RentableTypeRef Stop Date should not be before Start Date")
	// 	SvcGridErrorReturn(w, e, funcname)
	// 	return
	// }
	// // StopDate should not be before Today's date
	// if !(rlib.IsDateBefore((time.Time)(rfRecord.RSDtStart), (time.Time)(rfRecord.RSDtStop))) {
	// 	e := fmt.Errorf("RentableStatus Stop Date should not be before Start Date")
	// 	SvcGridErrorReturn(w, e, funcname)
	// 	return
	// }

	if rfRecord.RID > 0 {
		fmt.Printf("Updating Rentable with RID: %d ...\n", rfRecord.RID)
		// get Rentable from RID
		rt = rlib.GetRentable(rfRecord.RID)
		if !(rt.RID > 0) {
			e := fmt.Errorf("No such Rentable exists, RID: %d", rfRecord.RID)
			SvcGridErrorReturn(w, e, funcname)
			return
		}

		// TODO: if business value is changed then shouldn't we keep
		// the record of tie-up of this rentable with previous business?

		rt.BID = requestedBID
		rt.RentableName = rfRecord.RentableName
		rt.AssignmentTime = rfRecord.AssignmentTime
		// Now just update the Rentable Record
		err = rlib.UpdateRentable(&rt)
		if err != nil {
			e := fmt.Errorf("Error updating rentable: %s", err.Error())
			SvcGridErrorReturn(w, e, funcname)
			return
		}
		fmt.Printf("Rentable record has been updated with RID: %d\n", rt.RID)

		// ---------------- UPDATE RENTABLE TYPE REFERENCE ------------------------

		// get rental type ref object associated with this rentable
		rtr, err = rlib.GetRentableTypeRef(rfRecord.RTRID)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		// Create an updated version of rtr with the info submitted on this call
		rtr1 := rtr
		rtr1.DtStart = (time.Time)(rfRecord.RTRefDtStart)
		rtr1.DtStop = (time.Time)(rfRecord.RTRefDtStop)
		rtr1.RTID = rfRecord.RTID

		// if anything changed, remake the list of RTRs
		if !rtr1.DtStart.Equal(rtr.DtStart) || !rtr1.DtStop.Equal(rtr.DtStop) || rtr1.RTID != rtr.RTID {
			m, n := AdjustRTRTimeList(&rtr1, &rt) // returns current list and new list
			for i := 0; i < len(m); i++ {         // delete the current list
				err = rlib.DeleteRentableTypeRef(m[i].RTRID)
				if err != nil {
					SvcGridErrorReturn(w, err, funcname)
					return
				}
			}
			for i := 0; i < len(n); i++ { // insert the new list
				err = rlib.InsertRentableTypeRef(&n[i])
				if err != nil {
					SvcGridErrorReturn(w, err, funcname)
					return
				}
			}
		}

		// ---------------- UPDATE RENTABLE STATUS ------------------------

		// get rental status record associated with this rentable
		rs, err := rlib.GetRentableStatus(rfRecord.RSID)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		// Create an updated version of rtr with the info submitted on this call
		rs1 := rs
		rs1.DtStart = (time.Time)(rfRecord.RSDtStart)
		rs1.DtStop = (time.Time)(rfRecord.RSDtStop)
		rs1.Status = rlib.RentableStatusToNumber(rfRecord.RentableStatus)

		// if anything changed, remake the list of RTRs
		if !rs1.DtStart.Equal(rs.DtStart) || !rs1.DtStop.Equal(rs.DtStop) || rs1.Status != rs.Status {
			m, n := AdjustRSTimeList(&rs1, &rt) // returns current list and new list
			for i := 0; i < len(m); i++ {       // delete the current list
				err = rlib.DeleteRentableStatus(m[i].RSID)
				if err != nil {
					SvcGridErrorReturn(w, err, funcname)
					return
				}
			}
			for i := 0; i < len(n); i++ { // insert the new list
				err = rlib.InsertRentableStatus(&n[i])
				if err != nil {
					SvcGridErrorReturn(w, err, funcname)
					return
				}
			}
		}
	} else {
		fmt.Println("Inserting new Rentable Record...")
		fmt.Printf("Given RTID is %d\n", rfRecord.RTID)

		// --------------------- INSERT RENTABLE RECORD -------------------------
		rt.BID = requestedBID
		rt.RentableName = rfRecord.RentableName
		rt.AssignmentTime = rfRecord.AssignmentTime
		rid, err := rlib.InsertRentable(&rt)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		if !(rid > 0) {
			e := fmt.Errorf("Unable to insert new Rentable record")
			SvcGridErrorReturn(w, e, funcname)
			return
		}
		// assign RID for this rentable
		rt.RID = rid
		fmt.Printf("New Rentable record has been saved with RID: %d\n", rt.RID)

		// ------------------------- INSERT RENTABLE STATUS ---------------------------

		// insert rentable status for this Rentable
		rs.RID = rt.RID
		rs.BID = rt.BID
		rs.Status = rlib.RentableStatusToNumber(rfRecord.RentableStatus)
		rs.DtStart = currentTime
		rs.DtStop = (time.Time)(rfRecord.RSDtStop)
		err = rlib.InsertRentableStatus(&rs)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		fmt.Printf("RentableStatus has been saved for Rentable(%d), RSID: %d\n", rt.RID, rs.RSID)

		// ---------------------------- INSERT RENTABLE TYPE REF ---------------------

		// insert RentableTypeRef for this Rentable
		rtr.BID = rt.BID
		rtr.RID = rt.RID
		rtr.RTID = rfRecord.RTID
		rtr.DtStart = currentTime
		rtr.DtStop = (time.Time)(rfRecord.RTRefDtStop)
		// which default values should be inserted for OverrideRentCycle, OverrideProrationCycle
		// NOTE: don't worry about these two fields as of now
		// rtr.OverrideRentCycle = 0
		// rtr.OverrideProrationCycle = 0

		err = rlib.InsertRentableTypeRef(&rtr)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		fmt.Printf("RentableTypeRef has been saved for Rentable(%d), RTRID: %d\n", rt.RID, rtr.RTRID)
	}

	SvcWriteSuccessResponse(w)
}

// which fields needs to be fetched for SQL query for rentable details
var rentableFormSelectFields = []string{
	"Rentable.RID",
	"Rentable.RentableName",
	"RentalAgreementRentables.RARID",
	"RentalAgreementRentables.RAID",
	"RentalAgreementRentables.RARDtStart",
	"RentalAgreementRentables.RARDtStop",
	"RentableTypeRef.RTID",
	"RentableTypeRef.RTRID",
	"RentableTypeRef.DtStart as RTRefDtStart",
	"RentableTypeRef.DtStop as RTRefDtStop",
	"RentableTypes.Name",
	"RentableStatus.RSID",
	"RentableStatus.Status as RentableStatus",
	"RentableStatus.DtStart as RSDtStart",
	"RentableStatus.DtStop as RSDtStop",
	"Rentable.AssignmentTime",
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

	var (
		funcname = "getRentable"
		g        GetRentableResponse
		t        = time.Now()
	)
	fmt.Printf("entered %s\n", funcname)

	rentableQuery := `
	SELECT
		DISTINCT {{.SelectClause}}
	FROM Rentable
	INNER JOIN RentableTypeRef ON Rentable.RID = RentableTypeRef.RID
	INNER JOIN RentableTypes ON RentableTypeRef.RTID=RentableTypes.RTID
	INNER JOIN RentableStatus ON RentableStatus.RID=Rentable.RID
	LEFT JOIN RentalAgreementRentables ON RentalAgreementRentables.RID=Rentable.RID
	WHERE {{.WhereClause}};
	`

	// will be substituted as query clauses
	qc := queryClauses{
		"SelectClause": strings.Join(rentableFormSelectFields, ","),
		"WhereClause":  fmt.Sprintf("Rentable.BID=%d AND Rentable.RID=%d AND (RentalAgreementRentables.RARDtStart<=%q OR RentalAgreementRentables.RARDtStart IS NULL ) AND (RentalAgreementRentables.RARDtStop>%q OR RentalAgreementRentables.RARDtStop IS NULL) ", d.BID, d.RID, t.Format(rlib.RRDATEINPFMT), t.Format(rlib.RRDATEINPFMT)),
	}

	// get formatted query with substitution of select, where, order clause
	q := renderSQLQuery(rentableQuery, qc)
	fmt.Printf("db query = %s\n", q)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(q)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var gg RentableDetails
		gg.CurrentDate = rlib.JSONDate(t)

		for bud, bid := range rlib.RRdb.BUDlist {
			if bid == d.BID {
				gg.BID = rlib.XJSONBud(bud)
				break
			}
		}

		var rStatus int64
		err = rows.Scan(&gg.RID, &gg.RentableName, &gg.RARID, &gg.RAID, &gg.RARDtStart, &gg.RARDtStop, &gg.RTID, &gg.RTRID, &gg.RTRefDtStart, &gg.RTRefDtStop, &gg.RentableType, &gg.RSID, &rStatus, &gg.RSDtStart, &gg.RSDtStop, &gg.AssignmentTime)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		// convert status int to string, human readable
		gg.RentableStatus = rlib.RentableStatusToString(rStatus)

		g.Record = gg
	}
	// error check
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	// write response
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
