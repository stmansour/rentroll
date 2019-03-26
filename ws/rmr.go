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

// RentableMarketRateGridResponse holds the struct for grids response
type RentableMarketRateGridResponse struct {
	Status  string                      `json:"status"`
	Total   int64                       `json:"total"`
	Records []RentableMarketRateGridRec `json:"records"`
}

// RentableMarketRateGridRec holds a struct for single record of grid
type RentableMarketRateGridRec struct {
	Recid      int64 `json:"recid"`
	BID        int64
	BUD        rlib.XJSONBud
	RMRID      int64
	RTID       int64
	MarketRate float64
	DtStart    rlib.JSONDate
	DtStop     rlib.JSONDate
}

// MarketRateGridSave is the input data format for a Save command
type MarketRateGridSave struct {
	Cmd      string                      `json:"cmd"`
	Selected []int64                     `json:"selected"`
	Limit    int64                       `json:"limit"`
	Offset   int64                       `json:"offset"`
	Changes  []RentableMarketRateGridRec `json:"changes"`
}

// MarketRateGridDelete is a struct used in delete request for market rates
type MarketRateGridDelete struct {
	Cmd       string  `json:"cmd"`
	RMRIDList []int64 `json:"RMRIDList"`
}

// rmrGridRowScan scans a result from sql row and dump it in a struct for rentableGrid
func rmrGridRowScan(rows *sql.Rows, q RentableMarketRateGridRec) (RentableMarketRateGridRec, error) {
	err := rows.Scan(&q.RTID, &q.RMRID, &q.MarketRate, &q.DtStart, &q.DtStop)
	return q, err
}

var rmrSearchFieldMap = rlib.SelectQueryFieldMap{
	"RTID":       {"RentableMarketRate.RTID"},
	"RMRID":      {"RentableMarketRate.RMRID"},
	"MarketRate": {"RentableMarketRate.MarketRate"},
	"DtStart":    {"RentableMarketRate.DtStart"},
	"DtStop":     {"RentableMarketRate.DtStop"},
}

// which fields needs to be fetch to satisfy the struct
var rmrSearchSelectQueryFields = rlib.SelectQueryFields{
	"RentableMarketRate.RTID",
	"RentableMarketRate.RMRID",
	"RentableMarketRate.MarketRate",
	"RentableMarketRate.DtStart",
	"RentableMarketRate.DtStop",
}

// SvcHandlerRentableMarketRates returns the list of market rates for given rentable type
func SvcHandlerRentableMarketRates(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerRentableMarketRates"
	var (
		err error
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  RTID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	// This operation requires RentableType ID
	if d.ID < 0 {
		err = fmt.Errorf("ID for RentableType is required but was not specified")
		SvcErrorReturn(w, err, funcname)
		return
	}

	switch d.wsSearchReq.Cmd {
	case "get":
		svcSearchHandlerRentableMarketRates(w, r, d) // it is a query for the grid.
		break
	case "save":
		saveRentableTypeMarketRates(w, r, d)
		break
	case "delete":
		deleteRentableTypeMarketRates(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// svcSearchHandlerRentableMarketRates handles market rate grid request/response
func svcSearchHandlerRentableMarketRates(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "svcSearchHandlerRentableMarketRates"
	var (
		g     RentableMarketRateGridResponse
		err   error
		order = `RentableMarketRate.DtStart DESC`
		whr   = fmt.Sprintf("RentableMarketRate.RTID=%d", d.ID)
	)
	fmt.Printf("Entered %s\n", funcname)

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rmrSearchFieldMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	mrQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentableMarketRate
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rmrSearchSelectQueryFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(mrQuery, qc)
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
	queryWithLimit := mrQuery + limitAndOffsetClause

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
		var q RentableMarketRateGridRec
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		q, err = rmrGridRowScan(rows, q)
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

// saveRentableTypeMarketRates save/update market rates associated with RentableType
func saveRentableTypeMarketRates(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveRentableTypeMarketRates"
	var (
		err error
		foo MarketRateGridSave
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
	fmt.Printf("foo Changes: %v", foo.Changes)

	if len(foo.Changes) == 0 {
		e := fmt.Errorf("No MarketRate(s) provided for RentableType")
		SvcErrorReturn(w, e, funcname)
		return
	}

	// first check that, associated RentableType has allowed ManageToBudget field
	// if not then return with error
	var rt rlib.RentableType
	rtid := foo.Changes[0].RTID
	if err = rlib.GetRentableType(r.Context(), rtid, &rt); err != nil {
		e := fmt.Errorf("Error while getting RentableType: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	if rt.FLAGS&0x4 == 0 {
		e := fmt.Errorf("ManageToBudget is not enabled at this moment")
		SvcErrorReturn(w, e, funcname)
		return
	}

	var biz rlib.Business
	if err = rlib.GetBusiness(r.Context(), d.BID, &biz); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	var EDIadjust = (biz.FLAGS & 1) != 0
	rlib.Console("%s: EDIadjust = %t\n", funcname, EDIadjust)
	var bizErrs []bizlogic.BizError
	for _, mr := range foo.Changes {
		var a rlib.RentableMarketRate

		rlib.MigrateStructVals(&mr, &a) // the variables that don't need special handling

		if EDIadjust {
			rlib.EDIHandleIncomingDateRange(a.BID, &a.DtStart, &a.DtStop)
			rlib.Console("%s: after EDI adjust date range = %s\n", funcname, rlib.ConsoleDRange(&a.DtStart, &a.DtStop))
		}

		errs := bizlogic.ValidateRentableMarketRate(r.Context(), &a)
		if len(errs) > 0 {
			bizErrs = append(bizErrs, errs...)
			continue
		}

		// // insert new marketRate
		// if a.RMRID == 0 {
		// 	_, err = rlib.InsertRentableMarketRate(r.Context(), &a)
		// 	if err != nil {
		// 		e := fmt.Errorf("Error while inserting market rate:  %s", err.Error())
		// 		SvcErrorReturn(w, e, funcname)
		// 		return
		// 	}
		// } else { // else update existing one
		// 	err = rlib.UpdateRentableMarketRate(r.Context(), &a)
		// 	if err != nil {
		// 		e := fmt.Errorf("Error with updating market rate instance (%d), RTID=%d : %s", a.RMRID, a.RTID, err.Error())
		// 		SvcErrorReturn(w, e, funcname)
		// 		return
		// 	}
		// }
		if err = rlib.SetRentableMarketRate(r.Context(), &a); err != nil {
			e := fmt.Errorf("Error from SetRentableMarketRate (%d), RTID=%d : %s", a.RMRID, a.RTID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}

	}

	// if any marketRate has problem in bizlogic then return list
	if len(bizErrs) > 0 {
		SvcErrListReturn(w, bizErrs, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// deleteRentableTypeMarketRates used to delete multiple market rate records associated with rentable type
func deleteRentableTypeMarketRates(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteRentableTypeMarketRates"
	var (
		err error
		foo MarketRateGridDelete
	)
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data: %s\n", d.data)

	data := []byte(d.data)
	if err = json.Unmarshal(data, &foo); err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	for _, rmrid := range foo.RMRIDList {
		err = rlib.DeleteRentableMarketRate(r.Context(), rmrid)
		if err != nil {
			e := fmt.Errorf("Error with deleting MarketRate with ID %d:  %s", rmrid, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}
	SvcWriteSuccessResponse(d.BID, w)
}
