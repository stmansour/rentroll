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

// RentableLeaseStatusGridResponse to a response of grid , add on 01/23/2019
type RentableLeaseStatusGridResponse struct {
	Status  string                       `json:"status"`
	Total   int64                        `json:"total"`
	Records []RentableLeaseStatusGridRec `json:"records"`
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

// rsLeaseGridRowScan scans a result from sql row and dump it in a struct for rentableLeaseStatusGridRec, add on 01/23/2019
func rsLeaseGridRowScan(rows *sql.Rows, q RentableLeaseStatusGridRec) (RentableLeaseStatusGridRec, error) {
	err := rows.Scan(&q.RLID, &q.RID, &q.LeaseStatus, &q.DtStart, &q.DtStop, &q.Comment, &q.CreateBy, &q.LastModBy)
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

// SvcHandlerRentableLeaseStatus returns the list of Lease status for the rentableLeaseStatus add by lina on 01/23/2019
func SvcHandlerRentableLeaseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableLeaseStatus"
	var (
		err error
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  RID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

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
	case "delete":
		deleteRentableLeaseStatus(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// svcSearchHandlerRentableLeaseStatus handles market rate grid request/response
//
// add by lina on 01/23/2019
// Revised by sman throughout February 2019
//------------------------------------------------------------------------------
func svcSearchHandlerRentableLeaseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "svcSearchHandlerRentableLeaseStatus"

	var (
		g     RentableLeaseStatusGridResponse
		err   error
		order string
		whr   = fmt.Sprintf("RentableLeaseStatus.RID=%d", d.ID)
	)
	rlib.Console("Entered %s\n", funcname)

	// if date range was supplied, add it to the where clause...
	if d.wsSearchReq.SearchDtStart.Year() >= 2000 {
		whr += fmt.Sprintf(" AND DtStart < %q AND DtStop > %q",
			d.wsSearchReq.SearchDtStop.Format(rlib.RRDATETIMESQL),
			d.wsSearchReq.SearchDtStart.Format(rlib.RRDATETIMESQL),
		)
	}

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rentableLeaseStatusSearchFieldMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	} else {
		order = "DtStart DESC"
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
		rlib.Console("%s: Error from rlib.GetQueryCount: %s\n", funcname, err.Error())
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
	queryWithLimit := statusQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(queryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		rlib.Console("%s: Error from DB Query: %s\n", funcname, err.Error())
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

		//----------------
		// Handle EDI...
		//----------------
		// rlib.Console("%s: EDIHandleOutgoingJSONDateRange\n", funcname)
		rlib.EDIHandleOutgoingJSONDateRange(q.BID, &q.DtStart, &q.DtStop)

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
//------------------------------------------------------------------------------
func saveRentableLeaseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "saveRentableLeaseStatus"
		err      error
		foo      RentableLeaseStatusGridSave
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// first check that given such rentable exists or not
	if _, err = rlib.GetRentable(r.Context(), foo.RID); err != nil {
		e := fmt.Errorf("Error while getting Rentable: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// if there are no changes then nothing to do
	if len(foo.Changes) == 0 {
		SvcWriteSuccessResponse(d.BID, w)
		//
		// e := fmt.Errorf("No Rentable Status(s) provided for Rentable")
		// SvcErrorReturn(w, e, funcname)
		return
	}

	var biz rlib.Business
	if err = rlib.GetBusiness(r.Context(), d.BID, &biz); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	var EDIadjust = (biz.FLAGS & 1) != 0
	var bizErrs []bizlogic.BizError
	for _, rl := range foo.Changes {
		var a rlib.RentableLeaseStatus
		rlib.MigrateStructVals(&rl, &a) // the variables that don't need special handlingsman

		if EDIadjust {
			rlib.EDIHandleIncomingDateRange(a.BID, &a.DtStart, &a.DtStop)
		}

		rlib.Console("%s: Release Lease Status date range: %s\n", funcname, rlib.ConsoleDRange(&a.DtStart, &a.DtStop))

		errs := bizlogic.ValidateRentableLeaseStatus(r.Context(), &a)
		if len(errs) > 0 {
			bizErrs = append(bizErrs, errs...)
			continue
		}

		// Insert / update doesn't really matter here. SetRentableLeaseStatus
		// will determine what it needs to update, delete, and insert in order
		// to fill the span as
		if err = rlib.SetRentableLeaseStatus(r.Context(), &a); err != nil {
			e := fmt.Errorf("Error from SetRentableLeaseStatus (%d), RID=%d : %s", a.RLID, a.RID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	// if any rentable status has problem in bizlogic then return list
	if len(bizErrs) > 0 {
		SvcErrListReturn(w, bizErrs, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// RentableLeaseStatusGridRecDelete is a struct used in delete request for rentable status
type RentableLeaseStatusGridRecDelete struct {
	Cmd      string  `json:"cmd"`
	RLIDList []int64 `json:"RLIDList"`
	RID      int64   `json:"RID"`
}

// deleteRentableLeaseStatus used to delete rentable status records
// associated with rentable
//------------------------------------------------------------------------------
func deleteRentableLeaseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var funcname = "deleteRentableLeaseStatus"
	var err error
	var foo RentableLeaseStatusGridRecDelete

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// rlib.Console("foo.RLIDList length: %d\n", len(foo.RLIDList))

	// ------------------
	// START TRANSACTION
	// ------------------
	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	for _, rlid := range foo.RLIDList {
		err = rlib.DeleteRentableLeaseStatus(ctx, rlid)
		if err != nil {
			tx.Rollback()
			e := fmt.Errorf("Error with deleting Rentable Status(%d) for Rentable(%d): %s",
				rlid, foo.RID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	// ------------------
	// COMMIT TRANSACTION
	// ------------------
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}
