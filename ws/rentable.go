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

// PrRentableOther is a structure specifically for the UI. It will be
// automatically populated from an rlib.Rentable struct
type PrRentableOther struct {
	Recid                int64 `json:"recid"` // this is to support the w2ui form
	BID                  rlib.XJSONBud
	RID                  int64
	RentableName         string
	RTRID                rlib.NullInt64
	RentableType         rlib.NullString
	RTID                 rlib.NullInt64
	RSID                 rlib.NullInt64
	UseStatus            rlib.NullInt64
	LeaseStatus          rlib.NullInt64
	RARID                rlib.NullInt64
	RAID                 rlib.NullInt64
	RentalAgreementStart rlib.NullDate
	RentalAgreementStop  rlib.NullDate
	LastModTime          rlib.JSONDateTime
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
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	BID            int64
	BUD            rlib.XJSONBud
	RID            int64
	RentableName   string
	AssignmentTime int64
	Comment        string // for notes such as Alarm codes and other things
	LastModTime    rlib.JSONDateTime
	LastModBy      int64
	CreateTS       rlib.JSONDateTime
	CreateBy       int64
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
	const funcname = "SvcRentableTypeDown"
	var (
		g   RentableTypedownResponse
		err error
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("handle typedown: GetRentablesTypeDown( bid=%d, search=%s, limit=%d\n", d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	g.Records, err = rlib.GetRentableTypeDown(r.Context(), d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	if err != nil {
		e := fmt.Errorf("Error getting typedown matches: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	rlib.Console("GetRentableTypeDown returned %d matches\n", len(g.Records))
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// rentablesGridFields holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
var rentablesGridFieldsMap = map[string][]string{
	"RID":                  {"R.RID"},
	"RentableName":         {"R.RentableName"},
	"RTRID":                {"RTR.RTRID"},
	"RTID":                 {"RT.RTID"},
	"RentableType":         {"RT.Name"},
	"RSID":                 {"RS.RSID"},
	"UseStatus":            {"RS.UseStatus"},
	"LeaseStatus":          {"RS.LeaseStatus"},
	"RARID":                {"RAR.RARID"},
	"RAID":                 {"RAR.RAID"},
	"RentalAgreementStart": {"RAR.RARDtStart"},
	"RentalAgreementStop":  {"RAR.RARDtStop"},
}

// which fields needs to be fetched for SQL query for rentables
var rentablesQuerySelectFields = []string{
	"R.RID",
	"R.RentableName",
	"RTR.RTRID",
	"RT.Name as RentableType",
	"RT.RTID",
	"RS.RSID",
	"RS.UseStatus",
	"RS.LeaseStatus",
	"RAR.RARID",
	"RAR.RAID",
	"RAR.RARDtStart as RentalAgreementStart",
	"RAR.RARDtStop as RentalAgreementStop",
}

// rentablesRowScan scans a result from sql row and dump it in a PrRentableOther struct
func rentablesRowScan(rows *sql.Rows, q PrRentableOther) (PrRentableOther, error) {
	err := rows.Scan(&q.RID, &q.RentableName, &q.RTRID, &q.RentableType, &q.RTID, &q.RSID, &q.UseStatus, &q.LeaseStatus, &q.RARID, &q.RAID, &q.RentalAgreementStart, &q.RentalAgreementStop)
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
	const funcname = "SvcSearchHandlerRentables"
	var (
		err         error
		g           SearchRentablesResponse
		currentTime = time.Now()
	)
	rlib.Console("Entered %s\n", funcname)

	// fetch records from the database under this limit
	const (
		limitClause int = 100
	)

	// Show All Renbles no matter in what state they are,
	srch := fmt.Sprintf(`R.BID=%d`, d.BID)

	// show active rentable first by RenalAgreement Dates
	order := "R.RID ASC, RAR.RARID DESC, RTR.RTRID DESC, RS.RSID DESC" // default ORDER

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
	FROM Rentable AS R
	LEFT JOIN (
        SELECT RID, RTID, RTRID
        FROM RentableTypeRef
        WHERE DtStop > "{{.searchStart}}" AND DtStart <= "{{.searchStop}}" AND BID={{.BID}}
        GROUP BY RTRID
        ORDER BY RTRID DESC
    ) AS RTR ON R.RID=RTR.RID
    LEFT JOIN (
        SELECT RTID, Name
        FROM RentableTypes
        GROUP BY RTID
        ORDER BY RTID DESC
    ) RT ON RTR.RTID=RT.RTID
    LEFT JOIN (
        SELECT UseStatus, LeaseStatus, RID, RSID
        FROM RentableStatus
        WHERE DtStop > "{{.searchStart}}" AND DtStart <= "{{.searchStop}}" AND BID={{.BID}}
        GROUP BY RSID
        ORDER BY RSID DESC
    ) AS RS ON RS.RID=R.RID
    LEFT JOIN (
        SELECT RARID, RID, RAID, RARDtStart, RARDtStop
        FROM RentalAgreementRentables
        WHERE RARDtStop > "{{.searchStart}}" AND RARDtStart <= "{{.searchStop}}" AND BID={{.BID}}
        GROUP BY RARID
        ORDER BY RARID DESC
    ) AS RAR ON RAR.RID=R.RID
    WHERE {{.WhereClause}}
    GROUP BY R.RID
    ORDER BY {{.OrderClause}}` // don't add ';', later some parts will be added in query

	// will be substituted as query clauses
	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rentablesQuerySelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
		"currentTime":  currentTime.Format(rlib.RRDATEFMTSQL),                 // show associated instance(s) active as of current time
		"searchStart":  d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), // selected range start
		"searchStop":   d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL),  // selected range stop
		"BID":          strconv.FormatInt(d.BID, 10),
	}

	// GET TOTAL COUNT OF RESULTS
	countQuery := rlib.RenderSQLQuery(rentablesQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
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
	rentablesQueryWithLimit := rentablesQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(rentablesQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
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
			SvcErrorReturn(w, err, funcname)
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
		SvcErrorReturn(w, err, funcname)
		return
	}

	// write response
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
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
	const funcname = "SvcFormHandlerRentable"
	var (
		err error
	)
	rlib.Console("Entered %s\n", funcname)

	if d.RID, err = SvcExtractIDFromURI(r.RequestURI, "RID", 3, w); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Request: %s:  BID = %d,  RID = %d\n", d.wsSearchReq.Cmd, d.BID, d.RID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getRentable(w, r, d)
		break
	case "save":
		saveRentable(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// func dumpRTRList(m []rlib.RentableTypeRef) {
// 	for i := 0; i < len(m); i++ {
// 		rlib.Console("m[%d] range = %s - %s\n", i, m[i].DtStart.Format(rlib.RRDATEINPFMT), m[i].DtStop.Format(rlib.RRDATEINPFMT))
// 	}
// 	rlib.Console("----------------------------\n")
// }

/*// AdjustRTRTimeList determines what edits and/or inserts are needed to
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
func AdjustRTRTimeList(ctx context.Context, rtr *rlib.RentableTypeRef, r *rlib.Rentable) ([]rlib.RentableTypeRef, []rlib.RentableTypeRef) {
	const funcname = "AdjustRTRTimeList"
	var m []rlib.RentableTypeRef
	R, _ := rlib.GetRentableTypeRefs(ctx, r.RID)
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
					// rlib.Console("AdjustRTRTimeList:  add period rtr start to R[%d] start, rtr.DtStop moved forward:  %s\n", i, rtr.DtStart.Format(rlib.RRDATEINPFMT))
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
					// rlib.Console("AdjustRTRTimeList:  different types append:  i = %d\n", i)
					// dumpRTRList(m)
				} else {
					// rlib.Console("AdjustRTRTimeList: rtr is covered.  rtAdded set to true:  i = %d\n", i)
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
func AdjustRSTimeList(ctx context.Context, rs *rlib.RentableStatus, r *rlib.Rentable) ([]rlib.RentableStatus, []rlib.RentableStatus) {
	const funcname = "AdjustRSTimeList"
	var m []rlib.RentableStatus
	R, _ := rlib.GetAllRentableStatus(ctx, r.RID)
	l := len(R)
	rsAdded := false // flag to mark whether rs still needs to be added after loop
	for i := 0; i < l; i++ {
		if !rsAdded && rlib.DateRangeOverlap(&rs.DtStart, &rs.DtStop, &R[i].DtStart, &R[i].DtStop) {
			if rs.UseStatus == R[i].UseStatus { // same rentable status?
				if rs.DtStart.After(R[i].DtStart) { // adjust start time to the earliest
					rs.DtStart = R[i].DtStart
				}
				if rs.DtStop.Before(R[i].DtStop) { // adjust stop time to the latest
					rs.DtStop = R[i].DtStop
				}
			} else { // different types
				if rs.DtStart.Equal(R[i].DtStart) && rs.DtStop.Equal(R[i].DtStop) { // same date range, just a status change
					rs1 := R[i]
					rs1.UseStatus = rs.UseStatus
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
}*/

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
	const funcname = "saveRentable"
	var (
		err error
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	target := `"record":`
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("saveRentable: cannot find %s in form json", target)
		SvcErrorReturn(w, e, funcname)
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
		SvcErrorReturn(w, e, funcname)
		return
	}

	var (
		ok       bool
		rentable rlib.Rentable
		/*rs          rlib.RentableStatus
		rtr         rlib.RentableTypeRef
		currentTime = time.Now()*/
	)

	// checks for valid values
	requestedBID, ok := rlib.RRdb.BUDlist[string(rfRecord.BUD)]
	if !ok {
		e := fmt.Errorf("Invalid Business ID found. BUD: %s", rfRecord.BUD)
		SvcErrorReturn(w, e, funcname)
		return
	}

	/*// check whether rentable type is provided or not
	if !(rfRecord.RTID > 0) {
		e := fmt.Errorf("Rentable Type must be provided")
		SvcErrorReturn(w, e, funcname)
		return
	}*/

	// // StopDate should not be before Today's date
	// if !(rlib.IsDateBefore((time.Time)(rfRecord.RTRefDtStart), (time.Time)(rfRecord.RTRefDtStop))) {
	// 	e := fmt.Errorf("RentableTypeRef Stop Date should not be before Start Date")
	// 	SvcErrorReturn(w, e, funcname)
	// 	return
	// }
	// // StopDate should not be before Today's date
	// if !(rlib.IsDateBefore((time.Time)(rfRecord.RSDtStart), (time.Time)(rfRecord.RSDtStop))) {
	// 	e := fmt.Errorf("RentableStatus Stop Date should not be before Start Date")
	// 	SvcErrorReturn(w, e, funcname)
	// 	return
	// }

	if rfRecord.RID > 0 {
		rlib.Console("Updating Rentable with RID: %d ...\n", rfRecord.RID)
		// get Rentable from RID
		rentable, err = rlib.GetRentable(r.Context(), rfRecord.RID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		if !(rentable.RID > 0) {
			e := fmt.Errorf("No such Rentable exists, RID: %d", rfRecord.RID)
			SvcErrorReturn(w, e, funcname)
			return
		}

		// TODO: if business value is changed then shouldn't we keep
		// the record of tie-up of this rentable with previous business?

		rentable.BID = requestedBID
		rentable.RentableName = rfRecord.RentableName
		rentable.AssignmentTime = rfRecord.AssignmentTime
		rentable.Comment = rfRecord.Comment
		// Now just update the Rentable Record
		err = rlib.UpdateRentable(r.Context(), &rentable)
		if err != nil {
			e := fmt.Errorf("Error updating rentable: %s", err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		rlib.Console("Rentable record has been updated with RID: %d\n", rentable.RID)

		/*// ---------------- UPDATE RENTABLE TYPE REFERENCE ------------------------

		// get rental type ref object associated with this rentable
		rtr, err = rlib.GetRentableTypeRef(r.Context(), rfRecord.RTRID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		// Create an updated version of rtr with the info submitted on this call
		rtr1 := rtr
		rtr1.DtStart = (time.Time)(rfRecord.RTRefDtStart)
		rtr1.DtStop = (time.Time)(rfRecord.RTRefDtStop)
		rtr1.RTID = rfRecord.RTID

		// if anything changed, remake the list of RTRs
		if !rtr1.DtStart.Equal(rtr.DtStart) || !rtr1.DtStop.Equal(rtr.DtStop) || rtr1.RTID != rtr.RTID {
			m, n := AdjustRTRTimeList(r.Context(), &rtr1, &rt) // returns current list and new list
			for i := 0; i < len(m); i++ {                      // delete the current list
				err = rlib.DeleteRentableTypeRef(r.Context(), m[i].RTRID)
				if err != nil {
					SvcErrorReturn(w, err, funcname)
					return
				}
			}
			for i := 0; i < len(n); i++ { // insert the new list
				_, err = rlib.InsertRentableTypeRef(r.Context(), &n[i])
				if err != nil {
					SvcErrorReturn(w, err, funcname)
					return
				}
			}
		}

		// ---------------- UPDATE RENTABLE STATUS ------------------------

		// get rental status record associated with this rentable
		rs, err := rlib.GetRentableStatus(r.Context(), rfRecord.RSID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		// Create an updated version of rtr with the info submitted on this call
		rs1 := rs
		rs1.DtStart = (time.Time)(rfRecord.RSDtStart)
		rs1.DtStop = (time.Time)(rfRecord.RSDtStop)
		rs1.UseStatus = rlib.RentableStatusToNumber(rfRecord.RentableStatus)

		// if anything changed, remake the list of RTRs
		if !rs1.DtStart.Equal(rs.DtStart) || !rs1.DtStop.Equal(rs.DtStop) || rs1.UseStatus != rs.UseStatus {
			m, n := AdjustRSTimeList(r.Context(), &rs1, &rt) // returns current list and new list
			for i := 0; i < len(m); i++ {                    // delete the current list
				err = rlib.DeleteRentableStatus(r.Context(), m[i].RSID)
				if err != nil {
					SvcErrorReturn(w, err, funcname)
					return
				}
			}
			for i := 0; i < len(n); i++ { // insert the new list
				_, err = rlib.InsertRentableStatus(r.Context(), &n[i])
				if err != nil {
					SvcErrorReturn(w, err, funcname)
					return
				}
			}
		}*/
	} else {
		fmt.Println("Inserting new Rentable Record...")

		// --------------------- INSERT RENTABLE RECORD -------------------------
		rentable.BID = requestedBID
		rentable.RentableName = rfRecord.RentableName
		rentable.AssignmentTime = rfRecord.AssignmentTime
		rentable.Comment = rfRecord.Comment
		rid, err := rlib.InsertRentable(r.Context(), &rentable)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		if !(rid > 0) {
			e := fmt.Errorf("Unable to insert new Rentable record")
			SvcErrorReturn(w, e, funcname)
			return
		}
		// assign RID for this rentable
		rentable.RID = rid
		rlib.Console("New Rentable record has been saved with RID: %d\n", rentable.RID)

		/*// ------------------------- INSERT RENTABLE STATUS ---------------------------

		// insert rentable status for this Rentable
		rs.RID = rt.RID
		rs.BID = rt.BID
		rs.UseStatus = rlib.RentableStatusToNumber(rfRecord.RentableStatus)
		rs.DtStart = currentTime
		rs.DtStop = (time.Time)(rfRecord.RSDtStop)
		_, err = rlib.InsertRentableStatus(r.Context(), &rs)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		rlib.Console("RentableStatus has been saved for Rentable(%d), RSID: %d\n", rt.RID, rs.RSID)

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

		_, err = rlib.InsertRentableTypeRef(r.Context(), &rtr)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		rlib.Console("RentableTypeRef has been saved for Rentable(%d), RTRID: %d\n", rt.RID, rtr.RTRID)*/
	}

	SvcWriteSuccessResponseWithID(d.BID, w, rentable.RID)
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
	const funcname = "getRentable"
	var (
		g GetRentableResponse
	)
	rlib.Console("entered %s\n", funcname)

	rentable, err := rlib.GetRentable(r.Context(), d.RID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	var gg RentableDetails
	rlib.MigrateStructVals(&rentable, &gg) // the variables that don't need special handling
	gg.BUD = getBUDFromBIDList(gg.BID)
	g.Record = gg

	// write response
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// RentableStatusGridResponse to a response of grid
type RentableStatusGridResponse struct {
	Status  string                  `json:"status"`
	Total   int64                   `json:"total"`
	Records []RentableStatusGridRec `json:"records"`
}

// RentableStatusGridRec to a row record of the grid
type RentableStatusGridRec struct {
	Recid                 int64 `json:"recid"`
	RSID                  int64
	BID                   int64
	BUD                   string
	RID                   int64
	UseStatus             int64
	LeaseStatus           int64
	DtStart               rlib.JSONDate
	DtStop                rlib.JSONDate
	DtNoticeToVacate      rlib.JSONDate
	DtNoticeToVacateIsSet bool
	CreateBy              int64
	LastModBy             int64
	Comment               string
}

// rsGridRowScan scans a result from sql row and dump it in a struct for rentableStatusGridRec
func rsGridRowScan(rows *sql.Rows, q RentableStatusGridRec) (RentableStatusGridRec, error) {
	err := rows.Scan(&q.RSID, &q.RID, &q.UseStatus, &q.LeaseStatus, &q.DtStart, &q.DtStop, &q.DtNoticeToVacate, &q.CreateBy, &q.LastModBy)
	if err == nil {
		// Year 2000 date in UTC
		Y2KDt := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
		if (time.Time)(q.DtNoticeToVacate).After(Y2KDt) {
			q.DtNoticeToVacateIsSet = true
		}
	}
	return q, err
}

var rentableStatusSearchFieldMap = rlib.SelectQueryFieldMap{
	"RSID":             {"RentableStatus.RSID"},
	"RID":              {"RentableStatus.RID"},
	"UseStatus":        {"RentableStatus.UseStatus"},
	"LeaseStatus":      {"RentableStatus.LeaseStatus"},
	"DtStart":          {"RentableStatus.DtStart"},
	"DtStop":           {"RentableStatus.DtStop"},
	"DtNoticeToVacate": {"RentableStatus.DtNoticeToVacate"},
	"CreateBy":         {"RentableStatus.CreateBy"},
	"LastModBy":        {"RentableStatus.LastModBy"},
}

// which fields needs to be fetch to satisfy the struct
var rentableStatusSearchSelectQueryFields = rlib.SelectQueryFields{
	"RentableStatus.RSID",
	"RentableStatus.RID",
	"RentableStatus.UseStatus",
	"RentableStatus.LeaseStatus",
	"RentableStatus.DtStart",
	"RentableStatus.DtStop",
	"RentableStatus.DtNoticeToVacate",
	"RentableStatus.CreateBy",
	"RentableStatus.LastModBy",
}

// SvcHandlerRentableStatus returns the list of status for the rentable
func SvcHandlerRentableStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableStatus"
	var (
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  RID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	// This operation requires Rentable ID
	if d.ID < 0 {
		err = fmt.Errorf("ID for Rentable is not specified")
		SvcErrorReturn(w, err, funcname)
		return
	}

	switch d.wsSearchReq.Cmd {
	case "get":
		svcSearchHandlerRentableStatus(w, r, d) // it is a query for the grid.
		break
	case "save":
		saveRentableStatus(w, r, d)
		break
	case "delete":
		deleteRentableStatus(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// svcSearchHandlerRentableStatus handles market rate grid request/response
func svcSearchHandlerRentableStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	const funcname = "svcSearchHandlerRentableStatus"

	var (
		g     RentableStatusGridResponse
		err   error
		order = `RentableStatus.RSID ASC`
		whr   = fmt.Sprintf("RentableStatus.RID=%d", d.ID)
	)
	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rentableStatusSearchFieldMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	statusQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentableStatus
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rentableStatusSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(statusQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		fmt.Printf("%s: Error from rlib.GetQueryCount: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
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
	queryWithLimit := statusQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(queryWithLimit, qc)
	fmt.Printf("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q RentableStatusGridRec
		q.Recid = i
		q.BID = d.BID
		q.BUD = string(getBUDFromBIDList(q.BID))

		q, err = rsGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}

	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}

// RentableStatusGridSave is the input data format for a Save command
type RentableStatusGridSave struct {
	Cmd      string                  `json:"cmd"`
	Selected []int64                 `json:"selected"`
	Limit    int64                   `json:"limit"`
	Offset   int64                   `json:"offset"`
	Changes  []RentableStatusGridRec `json:"changes"`
	RID      int64                   `json:"RID"`
}

// saveRentableStatus save/update rentable status associated with Rentable
func saveRentableStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "saveRentableStatus"
		err      error
		foo      RentableStatusGridSave
	)
	fmt.Printf("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	fmt.Printf("foo Changes: %v\n", foo.Changes)

	// first check that given such rentable exists or not
	if _, err = rlib.GetRentable(r.Context(), foo.RID); err != nil {
		e := fmt.Errorf("Error while getting Rentable: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// if there are no changes then nothing to do
	if len(foo.Changes) == 0 {
		e := fmt.Errorf("No Rentable Status(s) provided for Rentable")
		SvcErrorReturn(w, e, funcname)
		return
	}

	var bizErrs []bizlogic.BizError
	for _, rs := range foo.Changes {
		var a rlib.RentableStatus
		rlib.MigrateStructVals(&rs, &a) // the variables that don't need special handling

		errs := bizlogic.ValidateRentableStatus(r.Context(), &a)
		if len(errs) > 0 {
			bizErrs = append(bizErrs, errs...)
			continue
		}

		// if RSID = 0 then insert new record
		if a.RSID == 0 {
			_, err = rlib.InsertRentableStatus(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error while inserting rentable status:  %s", err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		} else { // else update existing one
			err = rlib.UpdateRentableStatus(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error with updating rentable status (%d), RID=%d : %s", a.RSID, a.RID, err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		}
	}

	// if any rentable status has problem in bizlogic then return list
	if len(bizErrs) > 0 {
		SvcErrListReturn(w, bizErrs, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// RentableStatusGridRecDelete is a struct used in delete request for rentable status
type RentableStatusGridRecDelete struct {
	Cmd      string  `json:"cmd"`
	RSIDList []int64 `json:"RSIDList"`
	RID      int64   `json:"RID"`
}

// deleteRentableStatus used to delete rentable status records associated with rentable
func deleteRentableStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "deleteRentableStatus"
		err      error
		foo      RentableStatusGridRecDelete
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// TODO(Sudip): better should delete batch under atomic transaction
	for _, rsid := range foo.RSIDList {
		err = rlib.DeleteRentableStatus(r.Context(), rsid)
		if err != nil {
			e := fmt.Errorf("Error with deleting Rentable Status(%d) for Rentable(%d): %s",
				rsid, foo.RID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
}

// RentableTypeRefGridResponse to a response of grid
type RentableTypeRefGridResponse struct {
	Status  string                   `json:"status"`
	Total   int64                    `json:"total"`
	Records []RentableTypeRefGridRec `json:"records"`
}

// RentableTypeRefGridRec to a row record of the grid
type RentableTypeRefGridRec struct {
	Recid                  int64 `json:"recid"`
	RTRID                  int64
	RTID                   int64
	BID                    int64
	BUD                    string
	RID                    int64
	OverrideRentCycle      int64
	OverrideProrationCycle int64
	DtStart                rlib.JSONDate
	DtStop                 rlib.JSONDate
	CreateBy               int64
	LastModBy              int64
}

// rsGridRowScan scans a result from sql row and dump it in a struct for rentableTypeRefGridRec
func rtrGridRowScan(rows *sql.Rows, q RentableTypeRefGridRec) (RentableTypeRefGridRec, error) {
	err := rows.Scan(&q.RTRID, &q.RTID, &q.RID, &q.OverrideRentCycle, &q.OverrideProrationCycle, &q.DtStart, &q.DtStop, &q.CreateBy, &q.LastModBy)
	return q, err
}

var rentableTypeRefSearchFieldMap = rlib.SelectQueryFieldMap{
	"RTRID":                  {"RentableTypeRef.RTRID"},
	"RTID":                   {"RentableTypeRef.RTID"},
	"RID":                    {"RentableTypeRef.RID"},
	"OverrideRentCycle":      {"RentableTypeRef.OverrideRentCycle"},
	"OverrideProrationCycle": {"RentableTypeRef.OverrideProrationCycle"},
	"DtStart":                {"RentableTypeRef.DtStart"},
	"DtStop":                 {"RentableTypeRef.DtStop"},
	"CreateBy":               {"RentableTypeRef.CreateBy"},
	"LastModBy":              {"RentableTypeRef.LastModBy"},
}

// which fields needs to be fetch to satisfy the struct
var rentableTypeRefSearchSelectQueryFields = rlib.SelectQueryFields{
	"RentableTypeRef.RTRID",
	"RentableTypeRef.RTID",
	"RentableTypeRef.RID",
	"RentableTypeRef.OverrideRentCycle",
	"RentableTypeRef.OverrideProrationCycle",
	"RentableTypeRef.DtStart",
	"RentableTypeRef.DtStop",
	"RentableTypeRef.CreateBy",
	"RentableTypeRef.LastModBy",
}

// SvcHandlerRentableTypeRef returns the list of rentable type references
func SvcHandlerRentableTypeRef(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableTypeRef"
	var (
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  RID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	// This operation requires Rentable ID
	if d.ID < 0 {
		err = fmt.Errorf("ID for Rentable is not specified")
		SvcErrorReturn(w, err, funcname)
		return
	}

	switch d.wsSearchReq.Cmd {
	case "get":
		svcSearchHandlerRentableTypeRef(w, r, d) // it is a query for the grid.
		break
	case "save":
		saveRentableTypeRef(w, r, d)
		break
	case "delete":
		deleteRentableTypeRef(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// svcSearchHandlerRentableTypeRef handles market rate grid request/response
func svcSearchHandlerRentableTypeRef(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "svcSearchHandlerRentableTypeRef"

	var (
		g     RentableTypeRefGridResponse
		err   error
		order = `RentableTypeRef.RTRID ASC`
		whr   = fmt.Sprintf("RentableTypeRef.RID=%d", d.ID)
	)
	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rentableTypeRefSearchFieldMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	statusQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentableTypeRef
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rentableTypeRefSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(statusQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		fmt.Printf("%s: Error from rlib.GetQueryCount: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
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
	queryWithLimit := statusQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(queryWithLimit, qc)
	fmt.Printf("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q RentableTypeRefGridRec
		q.Recid = i
		q.BID = d.BID
		q.BUD = string(getBUDFromBIDList(q.BID))

		q, err = rtrGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}

	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}

// RentableTypeRefGridSave is the input data format for a Save command for rentable type ref instances
type RentableTypeRefGridSave struct {
	Cmd      string                   `json:"cmd"`
	Selected []int64                  `json:"selected"`
	Limit    int64                    `json:"limit"`
	Offset   int64                    `json:"offset"`
	Changes  []RentableTypeRefGridRec `json:"changes"`
	RID      int64                    `json:"RID"`
}

// saveRentableTypeRef save/update rentable type ref associated with Rentable
func saveRentableTypeRef(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveRentableTypeRef"
	var (
		err error
		foo RentableTypeRefGridSave
	)
	fmt.Printf("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	fmt.Printf("foo Changes: %v\n", foo.Changes)

	// first check that given such rentable exists or not
	if _, err = rlib.GetRentable(r.Context(), foo.RID); err != nil {
		e := fmt.Errorf("Error while getting Rentable: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// if there are no changes then nothing to do
	if len(foo.Changes) == 0 {
		e := fmt.Errorf("No Rentable Type Ref(s) provided for Rentable")
		SvcErrorReturn(w, e, funcname)
		return
	}

	var bizErrs []bizlogic.BizError
	for _, rs := range foo.Changes {
		var a rlib.RentableTypeRef
		rlib.MigrateStructVals(&rs, &a) // the variables that don't need special handling

		errs := bizlogic.ValidateRentableTypeRef(r.Context(), &a)
		if len(errs) > 0 {
			bizErrs = append(bizErrs, errs...)
			continue
		}

		// if RTRID = 0 then insert new record
		if a.RTRID == 0 {
			_, err = rlib.InsertRentableTypeRef(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error while inserting rentable type ref:  %s", err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		} else { // else update existing one
			err = rlib.UpdateRentableTypeRef(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error with updating rentable type ref (%d), RID=%d : %s", a.RTRID, a.RID, err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		}
	}

	// if any rentable type ref has problem in bizlogic then return list
	if len(bizErrs) > 0 {
		SvcErrListReturn(w, bizErrs, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// RentableTypeRefGridRecDelete is a struct used in delete request for rentable type ref
type RentableTypeRefGridRecDelete struct {
	Cmd       string  `json:"cmd"`
	RTRIDList []int64 `json:"RTRIDList"`
	RID       int64   `json:"RID"`
}

// deleteRentableTypeRef used to delete rentable type ref records associated with rentable
func deleteRentableTypeRef(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteRentableTypeRef"
	var (
		err error
		foo RentableTypeRefGridRecDelete
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// TODO(Sudip): better should delete batch under atomic transaction

	for _, rsid := range foo.RTRIDList {
		err = rlib.DeleteRentableTypeRef(r.Context(), rsid)
		if err != nil {
			e := fmt.Errorf("Error with deleting Rentable Status(%d) for Rentable(%d): %s",
				rsid, foo.RID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
}
