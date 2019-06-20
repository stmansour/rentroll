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

// Available defines the timerange, the type of rentable, and the specific
// rentable being reserved.
type Available struct {
	Recid        int64 `json:"recid"`
	BID          int64
	DtStart      rlib.JSONDateTime
	DtStop       rlib.JSONDateTime
	RLID         int64
	RTRID        int64
	RTID         int64
	RID          int64
	RentableName string
}

// AvailableResponse is the response data for a Rental Agreement Search
type AvailableResponse struct {
	Status  string      `json:"status"`
	Total   int64       `json:"total"`
	Records []Available `json:"records"`
}

// fields list needs to be fetched for grid
var resGridFieldsMap = map[string][]string{
	"DtStart":      {"RentableLeaseStatus.DtStart"},
	"DtStop":       {"RentableLeaseStatus.DtStop"},
	"RLID":         {"RentableLeaseStatus.RLID"},
	"RTRID":        {"RentableTypeRef.RTRID"},
	"RTID":         {"RentableTypeRef.RTID"},
	"RID":          {"RentableTypeRef.RID"},
	"BID":          {"RentableTypeRef.BID"},
	"RentableName": {"Rentable.RentableName"},
}

var resSelectFields = []string{
	"RentableLeaseStatus.DtStart",
	"RentableLeaseStatus.DtStop",
	"RentableLeaseStatus.RLID",
	"RentableTypeRef.RTRID",
	"RentableTypeRef.RTID",
	"RentableTypeRef.RID",
	"RentableTypeRef.BID",
	"Rentable.RentableName",
}

// availableRowScan scans a result from sql row and dump it in a Available struct
func availableRowScan(rows *sql.Rows, q Available) (Available, error) {
	err := rows.Scan(&q.DtStart, &q.DtStop, &q.RLID, &q.RTRID, &q.RTID, &q.RID, &q.BID, &q.RentableName)
	return q, err
}

// SvcAvailable searches for available rentables
// wsdoc {
//  @Title  SearchAvailable
//	@URL /v1/available/:BUI/[RLID]
//  @Method  POST
//	@Synopsis Returns a list of RIDs
//  @Description  Finds the rentables that are available between DtStart and DtStop.
//	@Input WebGridSearchRequest
//  @Response Available
// wsdoc }
//------------------------------------------------------------------------------
func SvcAvailable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcAvailable"
	const limitClause int = 100
	var err error
	var g AvailableResponse
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	//---------------------------------------
	// Unmarshal the reservation info...
	//---------------------------------------
	target := `"record":`
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("%s: cannot find %s in form json", funcname, target)
		SvcErrorReturn(w, e, funcname)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]

	// rentable Form Record
	var res Available
	err = json.Unmarshal([]byte(s), &res)
	rlib.Errcheck(err)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	//---------------------------------------
	// Now we can build the query...
	//---------------------------------------
	var sr string
	if res.RTID == 0 { // search for any rentable type
		m, err := rlib.GetBusinessRentableTypes(r.Context(), d.BID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
		}
		var sa []string
		for _, v := range m {
			rlib.Console("RTID = %d, FLAGS = %x\n", v.RTID, v.FLAGS)
			if v.FLAGS&0x8 == 0 {
				sa = append(sa, fmt.Sprintf("%d", v.RTID))
			}
		}
		if len(sa) == 0 {
			SvcErrorReturn(w, fmt.Errorf("no appropriate rentable types found"), funcname)
			return
		}
		sr = strings.Join(sa, ",")
	} else {
		sr = fmt.Sprintf("%d", res.RTID)
	}
	dtStart := time.Time(res.DtStart)
	dtStop := time.Time(res.DtStop)
	srch := fmt.Sprintf(`RentableTypeRef.BID=%d AND
        RentableLeaseStatus.DtStart <= %q AND RentableLeaseStatus.DtStop >= %q AND RentableLeaseStatus.LeaseStatus = 0 AND
		RentableTypeRef.DtStart <= %q AND RentableTypeRef.DtStop >= %q AND
		RentableTypeRef.RTID IN (%s) AND
		RentableUseType.DtStart <= %q AND RentableUseType.DtStop >= %q AND RentableUseType.UseType = 100`,
		res.BID,
		dtStart.Format(rlib.RRDATEFMTSQL),
		dtStop.Format(rlib.RRDATEFMTSQL),
		dtStart.Format(rlib.RRDATEFMTSQL),
		dtStop.Format(rlib.RRDATEFMTSQL),
		sr,
		dtStop.Format(rlib.RRDATEFMTSQL),
		dtStart.Format(rlib.RRDATEFMTSQL),
	)
	order := "RentableLeaseStatus.DtStart ASC,Rentable.RentableName ASC" // default ORDER

	//--------------------------------------------------
	// get WHERE clause and ORDER clause for sql query
	//--------------------------------------------------
	whereClause, orderClause := GetSearchAndSortSQL(d, resGridFieldsMap)
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	//--------------------------------------------------
	// Transactant Query Text Template
	//--------------------------------------------------
	mainQuery := `
SELECT DISTINCT {{.SelectClause}}
FROM RentableTypeRef
LEFT JOIN RentableLeaseStatus on RentableLeaseStatus.RID = RentableTypeRef.RID
LEFT JOIN RentableUseType on RentableUseType.RID = RentableTypeRef.RID
LEFT JOIN Rentable on Rentable.RID = RentableTypeRef.RID
WHERE {{.WhereClause}}
ORDER BY {{.OrderClause}}
` // don't add ';', later some parts will be added in query

	// select clause
	// RentableTypeRef.RTRID,
	// RentableTypeRef.RTID,
	// RentableTypeRef.RID,
	// RentableTypeRef.BID

	// where clause

	// will be substituted as query clauses
	qc := rlib.QueryClause{
		"SelectClause": strings.Join(resSelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
	}

	// GET TOTAL COUNTS of query
	countQuery := rlib.RenderSQLQuery(mainQuery, qc)
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
	resQueryWithLimit := mainQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(resQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	j := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var t Available
		t.Recid = j

		// get record of res
		t, err = availableRowScan(rows, t)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		g.Records = append(g.Records, t)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		j++ // update the index no matter what
	}
	// error check
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// write response
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)

}
