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

// RentableUseTypeGridResponse to a response of grid
type RentableUseTypeGridResponse struct {
	Status  string                   `json:"status"`
	Total   int64                    `json:"total"`
	Records []RentableUseTypeGridRec `json:"records"`
}

// RentableUseTypeGridRec to a row record of the grid
type RentableUseTypeGridRec struct {
	Recid     int64 `json:"recid"`
	UTID      int64
	BID       int64
	BUD       string
	RID       int64
	UseType   int64
	DtStart   rlib.JSONDate
	DtStop    rlib.JSONDate
	Comment   string
	CreateBy  int64
	LastModBy int64
}

// rutGridRowScan scans a result from sql row and dump it in a struct
// for rentableStatusGridRec
//-----------------------------------------------------------------------------
func rutGridRowScan(rows *sql.Rows, q RentableUseTypeGridRec) (RentableUseTypeGridRec, error) {
	err := rows.Scan(&q.UTID, &q.RID, &q.UseType /*&q.LeaseStatus,*/, &q.DtStart, &q.DtStop, &q.Comment, &q.CreateBy, &q.LastModBy)
	return q, err
}

// SvcHandlerRentableUseType returns the list of status for the rentable
//-----------------------------------------------------------------------------
func SvcHandlerRentableUseType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableUseType"
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
		svcSearchHandlerRentableUseType(w, r, d) // it is a query for the grid.
		break
	case "save":
		saveRentableUseType(w, r, d)
		break
	case "delete":
		deleteRentableUseType(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// which fields needs to be fetch to satisfy the struct
var rentableUseTypeSearchSelectQueryFields = rlib.SelectQueryFields{
	"RentableUseType.UTID",
	"RentableUseType.RID",
	"RentableUseType.UseType",
	"RentableUseType.DtStart",
	"RentableUseType.DtStop",
	"RentableUseType.Comment",
	"RentableUseType.CreateBy",
	"RentableUseType.LastModBy",
}

var rentableUseTypeSearchFieldMap = rlib.SelectQueryFieldMap{
	"UTID":      {"RentableUseType.UTID"},
	"RID":       {"RentableUseType.RID"},
	"UseType":   {"RentableUseType.UseType"},
	"DtStart":   {"RentableUseType.DtStart"},
	"DtStop":    {"RentableUseType.DtStop"},
	"Comment":   {"RentableUseType.Comment"},
	"CreateBy":  {"RentableUseType.CreateBy"},
	"LastModBy": {"RentableUseType.LastModBy"},
}

// svcSearchHandlerRentableUseType handles market rate grid request/response
//-----------------------------------------------------------------------------
func svcSearchHandlerRentableUseType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "svcSearchHandlerRentableUseType"
	var g RentableUseTypeGridResponse
	var err error
	var order = `RentableUseType.DtStart DESC`

	whr := fmt.Sprintf("RentableUseType.RID=%d", d.ID)
	rlib.Console("Entered %s\n", funcname)

	//------------------------------------------------------
	// get where clause and order clause for sql query
	//------------------------------------------------------
	whereClause, orderClause := GetSearchAndSortSQL(d, rentableUseTypeSearchFieldMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	statusQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentableUseType
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rentableUseTypeSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	//------------------------------------------------------
	// get TOTAL COUNT First
	//------------------------------------------------------
	countQuery := rlib.RenderSQLQuery(statusQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		rlib.Console("%s: Error from rlib.GetQueryCount: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("g.Total = %d\n", g.Total)

	//------------------------------------------------------
	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	//------------------------------------------------------
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`

	//------------------------------------------------------
	// build query with limit and offset clause
	// if query ends with ';' then remove it
	//------------------------------------------------------
	queryWithLimit := statusQuery + limitAndOffsetClause

	//------------------------------------------------------
	// Add limit and offset value
	//------------------------------------------------------
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	//------------------------------------------------------
	// get formatted query with substitution of select,
	// where, order clause
	//------------------------------------------------------
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
		var q RentableUseTypeGridRec
		q.Recid = i
		q.BID = d.BID
		q.BUD = string(rlib.GetBUDFromBIDList(q.BID))

		q, err = rutGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
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

// RentableUseTypeGridRecDelete is a struct used in delete request for rentable status
type RentableUseTypeGridRecDelete struct {
	Cmd      string  `json:"cmd"`
	UTIDList []int64 `json:"UTIDList"`
	RID      int64   `json:"RID"`
}

// deleteRentableUseType used to delete rentable status records associated
// with rentable
//-----------------------------------------------------------------------------
func deleteRentableUseType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var funcname = "deleteRentableUseType"
	var err error
	var foo RentableUseTypeGridRecDelete

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// ------------------
	// START TRANSACTION
	// ------------------
	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	for _, rsid := range foo.UTIDList {
		rlib.Console("Deleting UTID = %d\n", rsid)
		err = rlib.DeleteRentableUseType(ctx, rsid)
		if err != nil {
			tx.Rollback()
			e := fmt.Errorf("Error with deleting Rentable Status(%d) for Rentable(%d): %s", rsid, foo.RID, err.Error())
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
	rlib.Console("%s: DB Transaction successfully committed\n", funcname)
	SvcWriteSuccessResponse(d.BID, w)
}

// RentableUseTypeGridSave is the input data format for a Save command
type RentableUseTypeGridSave struct {
	Cmd      string                   `json:"cmd"`
	Selected []int64                  `json:"selected"`
	Limit    int64                    `json:"limit"`
	Offset   int64                    `json:"offset"`
	Changes  []RentableUseTypeGridRec `json:"changes"`
	RID      int64                    `json:"RID"`
}

// saveRentableUseType save/update rentable status associated with Rentable
//-----------------------------------------------------------------------------
func saveRentableUseType(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var funcname = "saveRentableUseType"
	var err error
	var foo RentableUseTypeGridSave
	var rnt rlib.Rentable
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// for i := 0; i < len(foo.Changes); i++ {
	// 	rlib.Console("Changes[%d]: UTID=%d  BID=%d  RID=%d  UseType=%d  Dates: %s\n",
	// 		i,
	// 		foo.Changes[i].UTID,
	// 		foo.Changes[i].BID,
	// 		foo.Changes[i].RID,
	// 		foo.Changes[i].UseType,
	// 		rlib.ConsoleJSONDRange(&foo.Changes[i].DtStart, &foo.Changes[i].DtStop))
	// }

	//------------------------------------------------------
	// first check that given such rentable exists or not
	//------------------------------------------------------
	if rnt, err = rlib.GetRentable(r.Context(), foo.Changes[0].RID); err != nil {
		s := fmt.Sprintf("Error while getting Rentable: %s", err.Error())
		e := fmt.Errorf(s)
		// rlib.Console(s)
		SvcErrorReturn(w, e, funcname)
		return
	}
	if rnt.RID != foo.Changes[0].RID {
		s := fmt.Sprintf("Rentable with RID=%d does not exist", foo.Changes[0].RID)
		e := fmt.Errorf(s)
		// rlib.Console(s)
		SvcErrorReturn(w, e, funcname)
		return
	}
	// rlib.Console("successfully loaded rentable: RID = %d\n", foo.Changes[0].RID)

	//------------------------------------------------------
	// if there are no changes then nothing to do
	//------------------------------------------------------
	if len(foo.Changes) == 0 {
		rlib.Console("No changes detected. No action taken\n")
		SvcWriteSuccessResponse(d.BID, w)
		return
	}

	var biz rlib.Business
	if err = rlib.GetBusiness(r.Context(), d.BID, &biz); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	var EDIadjust = (biz.FLAGS & 1) != 0 // sman: these are datetime values
	var bizErrs []bizlogic.BizError
	for _, rs := range foo.Changes {

		var a rlib.RentableUseType
		rlib.MigrateStructVals(&rs, &a) // the variables that don't need special handling

		if EDIadjust {
			rlib.EDIHandleIncomingDateRange(a.BID, &a.DtStart, &a.DtStop)
		}

		if err = rlib.SetRentableUseType(r.Context(), &a); err != nil {
			e := fmt.Errorf("Error from SetRentableUseType (%d), RID=%d : %s", a.UTID, a.RID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	//------------------------------------------------------
	// if any rentable status has problem in bizlogic then return list
	//------------------------------------------------------
	if len(bizErrs) > 0 {
		SvcErrListReturn(w, bizErrs, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}
