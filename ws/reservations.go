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

//-------------------------------------------------------------------
//                        **** SEARCH ****
//-------------------------------------------------------------------

// Reservation defines the timerange, the type of rentable, and the specific
// rentable being reserved.
type Reservation struct {
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

// // ReservationMatch provides a list of RIDs of rentables of type RTID that are
// // available during the requested time frame (DtStart - DtStop)
// type ReservationMatch struct {
// 	BID     int64
// 	DtStart rlib.JSONDateTime
// 	DtStop  rlib.JSONDateTime
// 	RTID    int64
// 	RID     []int64
// }

// ReservationResponse is the response data for a Rental Agreement Search
type ReservationResponse struct {
	Status  string        `json:"status"`
	Total   int64         `json:"total"`
	Records []Reservation `json:"records"`
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

//-------------------------------------------------------------------
//                         **** SAVE ****
//-------------------------------------------------------------------

// ResDet is the structure that fully describes a reservation
type ResDet struct {
	Recid        int64             `json:"recid"`
	BID          int64             `json:"rdBID"`
	DtStart      rlib.JSONDateTime `json:"DtStart"`
	DtStop       rlib.JSONDateTime `json:"DtStop"`
	RLID         int64             `json:"RLID"`
	RTRID        int64             `json:"RTRID"`
	RTID         int64             `json:"rdRTID"`
	RID          int64             `json:"RID"`
	LeaseStatus  int64             `json:"LeaseStatus"`
	Nights       int64             // computed field `json:"Nights"`
	RentableName string            `json:"RentableName"`
	FirstName    string            `json:"FirstName"`
	LastName     string            `json:"LastName"`
	Email        string            `json:"Email"`
	Phone        string            `json:"Phone"`
	Street       string            `json:"Street"`
	City         string            `json:"City"`
	Country      string            `json:"Country"`
	State        string            `json:"State"`
	PostalCode   string            `json:"PostalCode"`
	CCName       string            `json:"CCName"`
	CCType       string            `json:"CCType"`
	CCNumber     string            `json:"CCNumber"`
	CCExpMonth   string            `json:"CCExpMonth"`
	CCExpYear    string            `json:"CCExpYear"`
	Comment      string            `json:"Comment"`
}

// SaveReservation is sent to save one of open time slots as a reservation
type SaveReservation struct {
	Cmd    string `json:"cmd"`
	Record ResDet `json:"record"`
}

//-------------------------------------------------------------------
//                         **** GET ****
//-------------------------------------------------------------------

// GetReservation may be called to change a reservation time

//-----------------------------------------------------------------------------
//##########################################################################################################################################################
//-----------------------------------------------------------------------------

// SvcReservationDispatch dispatches a request for a reservation.
//       0    1          2    3 (optional)
// 		/v1/reservation/BID/RLID
//
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcReservationDispatch(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcReservationDispatch"
	var err error

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d, DtStart/Stop = %s\n", d.wsSearchReq.Cmd, d.BID, rlib.ConsoleDRange(&d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop))

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID < 1 {
			SvcReservationSearch(w, r, d)
			return
		}
		getReservation(w, r, d)
	case "save":
		saveReservation(w, r, d)
	case "delete":
		deleteReservation(w, r, d)
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
	}
}

// resRowScan scans a result from sql row and dump it in a Reservation struct
func resRowScan(rows *sql.Rows, q Reservation) (Reservation, error) {
	err := rows.Scan(&q.DtStart, &q.DtStop, &q.RLID, &q.RTRID, &q.RTID, &q.RID, &q.BID, &q.RentableName)
	return q, err
}

// SvcReservationSearch searches for available rentables
// wsdoc {
//  @Title  SearchReservation
//	@URL /v1/reservation/:BUI/[RLID]
//  @Method  POST
//	@Synopsis Returns a list of RIDs
//  @Description  Finds the rentables that are available between DtStart and DtStop.
//	@Input WebGridSearchRequest
//  @Response Reservation
// wsdoc }
//------------------------------------------------------------------------------
func SvcReservationSearch(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcStatementPayors"
	const limitClause int = 100
	var err error
	var g ReservationResponse
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
	var res Reservation
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
	dtStart := time.Time(res.DtStart)
	dtStop := time.Time(res.DtStop)
	srch := fmt.Sprintf(`RentableTypeRef.BID=%d AND
        RentableLeaseStatus.DtStart <= %q AND RentableLeaseStatus.DtStop >= %q AND RentableLeaseStatus.LeaseStatus = 0 AND
		RentableTypeRef.DtStart <= %q AND RentableTypeRef.DtStop >= %q AND RentableTypeRef.RTID = %d AND
		RentableUseStatus.DtStart <= %q AND RentableUseStatus.DtStop >= %q AND RentableUseStatus.UseStatus = 0`,
		res.BID,
		dtStart.Format(rlib.RRDATEFMTSQL),
		dtStop.Format(rlib.RRDATEFMTSQL),
		dtStart.Format(rlib.RRDATEFMTSQL),
		dtStop.Format(rlib.RRDATEFMTSQL),
		res.RTID,
		dtStart.Format(rlib.RRDATEFMTSQL),
		dtStop.Format(rlib.RRDATEFMTSQL),
	)
	order := "RentableLeaseStatus.DtStart ASC" // default ORDER

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
SELECT {{.SelectClause}}
FROM RentableTypeRef
LEFT JOIN RentableLeaseStatus on RentableLeaseStatus.RID = RentableTypeRef.RID
LEFT JOIN RentableUseStatus on RentableUseStatus.RID = RentableTypeRef.RID
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
		var t Reservation
		t.Recid = j

		// get record of res
		t, err = resRowScan(rows, t)
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
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)

}

func getReservation(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getReservation"
	rlib.Console("Entered %s\n", funcname)
}

// saveReservation
// wsdoc {
//  @Title  SaveReservation
//	@URL /v1/reservation/:BUI/[RLID]
//  @Method  POST
//	@Synopsis Returns a list of RIDs
//  @Description  Saves the ReleaseStatus. If RLID is 0, a new status is created.
//  @Description  If RLID is > 0 it is simply updated
//	@Input WebGridSearchRequest
//  @Response Reservation
// wsdoc }
//------------------------------------------------------------------------------
func saveReservation(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveReservation"
	rlib.Console("Entered %s\n", funcname)
	target := `"record":`
	i := strings.Index(d.data, target)
	if i < 0 {
		e := fmt.Errorf("%s: cannot find %s in form json", funcname, target)
		SvcErrorReturn(w, e, funcname)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]

	// rlib.Console("Data to unmarshal = %s\n", s)

	// Reservation
	var res ResDet
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// rlib.Console("Successfully unmarshaled reservation: %s %s\n", res.FirstName, res.LastName)
	// rlib.Console("    res.BID: %d   d.BID = %d\n", res.BID, d.BID)
	ctx := rlib.SetSessionContextKey(r.Context(), d.sess)

	var rls = rlib.RentableLeaseStatus{
		RLID:        res.RLID,
		RID:         res.RID,
		BID:         d.BID, // DEBUG: res.BID is 0 in ws func test 41b ?????
		DtStart:     time.Time(res.DtStart),
		DtStop:      time.Time(res.DtStop),
		LeaseStatus: res.LeaseStatus,
		Comment:     res.Comment,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Phone:       res.Phone,
		Address:     res.Street,
		City:        res.City,
		State:       res.State,
		PostalCode:  res.PostalCode,
		Country:     res.Country,
		CCName:      res.CCName,
		CCType:      res.CCType,
		CCNumber:    res.CCNumber,
		CCExpMonth:  res.CCExpMonth,
	}

	err = rlib.SetRentableLeaseStatus(ctx, &rls, false)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	SvcWriteSuccessResponse(rls.BID, w)
}

func deleteReservation(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteReservation"
	rlib.Console("Entered %s\n", funcname)
}
