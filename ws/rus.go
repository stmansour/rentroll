package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// RentableUseStatusGridResponse to a response of grid
type RentableUseStatusGridResponse struct {
	Status  string                     `json:"status"`
	Total   int64                      `json:"total"`
	Records []RentableUseStatusGridRec `json:"records"`
}

// RentableUseStatusGridRec to a row record of the grid
type RentableUseStatusGridRec struct {
	Recid       int64 `json:"recid"`
	RSID        int64
	BID         int64
	BUD         string
	RID         int64
	UseStatus   int64
	LeaseStatus int64
	DtStart     rlib.JSONDate
	DtStop      rlib.JSONDate
	Comment     string
	CreateBy    int64
	LastModBy   int64
	// DtNoticeToVacateIsSet bool
}

// SvcHandlerRentableUseStatus returns the list of status for the rentable
func SvcHandlerRentableUseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableUseStatus"
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
		svcSearchHandlerRentableUseStatus(w, r, d) // it is a query for the grid.
		break
	case "save":
		saveRentableUseStatus(w, r, d)
		break
	case "delete":
		deleteRentableUseStatus(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// svcSearchHandlerRentableUseStatus handles market rate grid request/response
func svcSearchHandlerRentableUseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	const funcname = "svcSearchHandlerRentableUseStatus"

	var (
		g     RentableUseStatusGridResponse
		err   error
		order = `RentableUseStatus.DtStart ASC`
		whr   = fmt.Sprintf("RentableUseStatus.RID=%d", d.ID)
	)
	rlib.Console("Entered %s\n", funcname)

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
	FROM RentableUseStatus
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
		var q RentableUseStatusGridRec
		q.Recid = i
		q.BID = d.BID
		q.BUD = string(rlib.GetBUDFromBIDList(q.BID))

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

// RentableUseStatusGridRecDelete is a struct used in delete request for rentable status
type RentableUseStatusGridRecDelete struct {
	Cmd      string  `json:"cmd"`
	RSIDList []int64 `json:"RSIDList"`
	RID      int64   `json:"RID"`
}

// deleteRentableUseStatus used to delete rentable status records associated with rentable
func deleteRentableUseStatus(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "deleteRentableUseStatus"
		err      error
		foo      RentableUseStatusGridRecDelete
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
		err = rlib.DeleteRentableUseStatus(r.Context(), rsid)
		if err != nil {
			e := fmt.Errorf("Error with deleting Rentable Status(%d) for Rentable(%d): %s",
				rsid, foo.RID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
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
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	rlib.Console("foo Changes: %v\n", foo.Changes)

	// first check that given such rentable exists or not
	if _, err = rlib.GetRentable(r.Context(), foo.RID); err != nil {
		e := fmt.Errorf("Error while getting Rentable: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// if there are no changes then nothing to do
	if len(foo.Changes) == 0 {
		SvcWriteSuccessResponse(d.BID, w)
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
	for _, rs := range foo.Changes {

		var a rlib.RentableUseStatus
		rlib.MigrateStructVals(&rs, &a) // the variables that don't need special handling

		if EDIadjust {
			rlib.EDIHandleIncomingDateRange(a.BID, &a.DtStart, &a.DtStop)
		}

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
