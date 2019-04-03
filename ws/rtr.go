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

// rsUseGridRowScan scans a result from sql row and dump it in a struct for rentableTypeRefGridRec
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
		order = `RentableTypeRef.DtStart DESC`
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
		q.BID = d.BID
		q.BUD = string(rlib.GetBUDFromBIDList(q.BID))

		q, err = rtrGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		q.Recid = q.RTRID

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

	var biz rlib.Business
	if err = rlib.GetBusiness(r.Context(), d.BID, &biz); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	var EDIadjust = (biz.FLAGS & 1) != 0
	var bizErrs []bizlogic.BizError
	for _, rs := range foo.Changes {
		var a rlib.RentableTypeRef
		rlib.MigrateStructVals(&rs, &a) // the variables that don't need special handling

		if EDIadjust {
			rlib.EDIHandleIncomingDateRange(a.BID, &a.DtStart, &a.DtStop)
		}

		errs := bizlogic.ValidateRentableTypeRef(r.Context(), &a)
		if len(errs) > 0 {
			bizErrs = append(bizErrs, errs...)
			continue
		}

		// // if RTRID = 0 then insert new record
		// if a.RTRID == 0 {
		// 	_, err = rlib.InsertRentableTypeRef(r.Context(), &a)
		// 	if err != nil {
		// 		e := fmt.Errorf("Error while inserting rentable type ref:  %s", err.Error())
		// 		SvcErrorReturn(w, e, funcname)
		// 		return
		// 	}
		// } else { // else update existing one
		// 	err = rlib.UpdateRentableTypeRef(r.Context(), &a)
		// 	if err != nil {
		// 		e := fmt.Errorf("Error with updating rentable type ref (%d), RID=%d : %s", a.RTRID, a.RID, err.Error())
		// 		SvcErrorReturn(w, e, funcname)
		// 		return
		// 	}
		// }
		if err = rlib.SetRentableTypeRef(r.Context(), &a); err != nil {
			e := fmt.Errorf("Error from SetRentableTypeReference (%d), RID=%d : %s", a.RTRID, a.RID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
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
