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

// RentableLeaseStatusGridRec to a row record of the grid, add on 01/23/2019
type RentableLeaseStatusGridRec struct {
	Recid       int64 `json:"recid"`
	RLID        int64
	BID         int64
	BUD         string
	RID         int64
	LeaseStatus int64
	DtStart     rlib.JSONDate
	DtStop      rlib.JSONDate
	Comment     string
	CreateBy    int64
	LastModBy   int64
	// DtNoticeToVacateIsSet bool
}

// rsLeaseGridRowScan scans a result from sql row and dump it in a struct for rentableLeaseStatusGridRec, add on 01/23/2019
func rsLeaseGridRowScan(rows *sql.Rows, q RentableLeaseStatusGridRec) (RentableLeaseStatusGridRec, error) {
	err := rows.Scan(&q.RLID, &q.RID, &q.LeaseStatus, &q.DtStart, &q.DtStop, &q.Comment, &q.CreateBy, &q.LastModBy)
	return q, err
}

// rsGridRowScan scans a result from sql row and dump it in a struct for rentableStatusGridRec
func rsGridRowScan(rows *sql.Rows, q RentableUseStatusGridRec) (RentableUseStatusGridRec, error) {
	err := rows.Scan(&q.RSID, &q.RID, &q.UseStatus /*&q.LeaseStatus,*/, &q.DtStart, &q.DtStop, &q.Comment, &q.CreateBy, &q.LastModBy)
	// if err == nil {
	// 	// Year 2000 date in UTC
	// 	Y2KDt := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	// 	if (time.Time)(q.DtNoticeToVacate).After(Y2KDt) {
	// 		q.DtNoticeToVacateIsSet = true
	// 	}
	// }
	return q, err
}

var rentableLeaseStatusSearchFieldMap = rlib.SelectQueryFieldMap{ //add on 01/23/2019
	"RLID":        {"RentableLeaseStatus.RLID"},
	"RID":         {"RentableLeaseStatus.RID"},
	"LeaseStatus": {"RentableLeaseStatus.LeaseStatus"},
	"DtStart":     {"RentableLeaseStatus.DtStart"},
	"DtStop":      {"RentableLeaseStatus.DtStop"},
	"Comment":     {"RentableLeaseStatus.Comment"},
	"CreateBy":    {"RentableLeaseStatus.CreateBy"},
	"LastModBy":   {"RentableLeaseStatus.LastModBy"},
}

// which fields needs to be fetch to satisfy the struct
var rentableLeaseStatusSearchSelectQueryFields = rlib.SelectQueryFields{ //add on 01/23/2019
	"RentableLeaseStatus.RLID",
	"RentableLeaseStatus.RID",
	"RentableLeaseStatus.LeaseStatus",
	"RentableLeaseStatus.DtStart",
	"RentableLeaseStatus.DtStop",
	"RentableLeaseStatus.Comment",
	"RentableLeaseStatus.CreateBy",
	"RentableLeaseStatus.LastModBy",
}

var rentableStatusSearchFieldMap = rlib.SelectQueryFieldMap{
	"RSID":      {"RentableUseStatus.RSID"},
	"RID":       {"RentableUseStatus.RID"},
	"UseStatus": {"RentableUseStatus.UseStatus"},
	// "LeaseStatus": {"RentableUseStatus.LeaseStatus"},
	"DtStart":   {"RentableUseStatus.DtStart"},
	"DtStop":    {"RentableUseStatus.DtStop"},
	"Comment":   {"RentableUseStatus.Comment"},
	"CreateBy":  {"RentableUseStatus.CreateBy"},
	"LastModBy": {"RentableUseStatus.LastModBy"},
}

// which fields needs to be fetch to satisfy the struct
var rentableStatusSearchSelectQueryFields = rlib.SelectQueryFields{
	"RentableUseStatus.RSID",
	"RentableUseStatus.RID",
	"RentableUseStatus.UseStatus",
	// "RentableUseStatus.LeaseStatus",
	"RentableUseStatus.DtStart",
	"RentableUseStatus.DtStop",
	"RentableUseStatus.Comment",
	"RentableUseStatus.CreateBy",
	"RentableUseStatus.LastModBy",
}

// SvcHandlerRentableLeaseStatus returns the list of Lease status for the rentableLeaseStatus add by lina on 01/23/2019
func SvcHandlerRentableLeaseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableLeaseStatus"
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
		svcSearchHandlerRentableLeaseStatus(w, r, d) // it is a query for the grid.
		break
	case "save":
		saveRentableLeaseStatus(w, r, d)
		break
	//case "delete":
	//	deleteRentableUseStatus(w, r, d)
	//	break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// svcSearchHandlerRentableLeaseStatus handles market rate grid request/response, add by lina on 01/23/2019
func svcSearchHandlerRentableLeaseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "svcSearchHandlerRentableLeaseStatus"

	var (
		g     RentableLeaseStatusGridResponse
		err   error
		order = `RentableLeaseStatus.DtStart ASC`
		whr   = fmt.Sprintf("RentableLeaseStatus.RID=%d", d.ID)
	)
	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rentableLeaseStatusSearchFieldMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	statusQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentableLeaseStatus
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rentableLeaseStatusSearchSelectQueryFields, ","),
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
		var q RentableLeaseStatusGridRec
		q.Recid = i
		q.BID = d.BID
		q.BUD = string(rlib.GetBUDFromBIDList(q.BID))

		q, err = rsLeaseGridRowScan(rows, q)
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

// saveRentableLeaseStatus save/update rentable status associated with Rentable , add by lina 01/23/2019
func saveRentableLeaseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "saveRentableLeaseStatus"
		err      error
		foo      RentableLeaseStatusGridSave
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
	for _, rl := range foo.Changes {
		var a rlib.RentableLeaseStatus
		rlib.MigrateStructVals(&rl, &a) // the variables that don't need special handling

		errs := bizlogic.ValidateRentableLeaseStatus(r.Context(), &a)
		if len(errs) > 0 {
			bizErrs = append(bizErrs, errs...)
			continue
		}

		// if RSID = 0 then insert new record
		if a.RLID == 0 {
			_, err = rlib.InsertRentableLeaseStatus(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error while inserting rentable status:  %s", err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		} else { // else update existing one
			err = rlib.UpdateRentableLeaseStatus(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error with updating rentable status (%d), RID=%d : %s", a.RLID, a.RID, err.Error())
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

// RentableLeaseStatusGridSave is the input data format for a Save command, add on 01/23/2019
type RentableLeaseStatusGridSave struct {
	Cmd      string                       `json:"cmd"`
	Selected []int64                      `json:"selected"`
	Limit    int64                        `json:"limit"`
	Offset   int64                        `json:"offset"`
	Changes  []RentableLeaseStatusGridRec `json:"changes"`
	RID      int64                        `json:"RID"`
}

// RentableUseStatusGridSave is the input data format for a Save command
type RentableUseStatusGridSave struct {
	Cmd      string                     `json:"cmd"`
	Selected []int64                    `json:"selected"`
	Limit    int64                      `json:"limit"`
	Offset   int64                      `json:"offset"`
	Changes  []RentableUseStatusGridRec `json:"changes"`
	RID      int64                      `json:"RID"`
}

// saveRentableUseStatus save/update rentable status associated with Rentable
func saveRentableUseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "saveRentableUseStatus"
		err      error
		foo      RentableUseStatusGridSave
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
		var a rlib.RentableUseStatus
		rlib.MigrateStructVals(&rs, &a) // the variables that don't need special handling

		errs := bizlogic.ValidateRentableUseStatus(r.Context(), &a)
		if len(errs) > 0 {
			bizErrs = append(bizErrs, errs...)
			continue
		}

		// if RSID = 0 then insert new record
		if a.RSID == 0 {
			_, err = rlib.InsertRentableUseStatus(r.Context(), &a)
			if err != nil {
				e := fmt.Errorf("Error while inserting rentable status:  %s", err.Error())
				SvcErrorReturn(w, e, funcname)
				return
			}
		} else { // else update existing one
			err = rlib.UpdateRentableUseStatus(r.Context(), &a)
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
