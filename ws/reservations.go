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
	Recid            int64             `json:"recid"`
	RLID             int64             // rentable lease status id (reservation id)
	RID              int64             // specific rentable reserved
	RTID             int64             // the rentable type
	ConfirmationCode string            // reservation ConfirmationCode
	DtStart          rlib.JSONDateTime // res start time
	DtStop           rlib.JSONDateTime // res stop time
	FirstName        string            // res name
	LastName         string            // res name
	Email            string            // email on reservation
	Phone            string            // phone on reservation
	RentableName     string            // rentable name
	Name             string            // Rentable Type Name
}

// SearchReservationResponse is the response data for a Rental Agreement Search
type SearchReservationResponse struct {
	Status  string        `json:"status"`
	Total   int64         `json:"total"`
	Records []Reservation `json:"records"`
}

//-------------------------------------------------------------------
//                         **** SAVE ****
//-------------------------------------------------------------------

// ResDet is the structure that fully describes a reservation
type ResDet struct {
	Recid            int64             `json:"recid"`
	BID              int64             `json:"rdBID"`
	DtStart          rlib.JSONDateTime `json:"DtStart"`
	DtStop           rlib.JSONDateTime `json:"DtStop"`
	RLID             int64             `json:"RLID"`
	RTRID            int64             `json:"RTRID"`
	RTID             int64             `json:"rdRTID"`
	RID              int64             `json:"RID"`
	LeaseStatus      int64             `json:"LeaseStatus"`
	Nights           int64             // computed field `json:"Nights"`
	RentableName     string            `json:"RentableName"`
	FirstName        string            `json:"FirstName"`
	LastName         string            `json:"LastName"`
	Email            string            `json:"Email"`
	Phone            string            `json:"Phone"`
	Street           string            `json:"Street"`
	City             string            `json:"City"`
	Country          string            `json:"Country"`
	State            string            `json:"State"`
	PostalCode       string            `json:"PostalCode"`
	CCName           string            `json:"CCName"`
	CCType           string            `json:"CCType"`
	CCNumber         string            `json:"CCNumber"`
	CCExpMonth       string            `json:"CCExpMonth"`
	CCExpYear        string            `json:"CCExpYear"`
	ConfirmationCode string            `json:"ConfirmationCode"`
	Comment          string            `json:"Comment"`
}

// SaveReservation is sent to save one of open time slots as a reservation
type SaveReservation struct {
	Cmd    string `json:"cmd"`
	Record ResDet `json:"record"`
}

//-------------------------------------------------------------------
//                         **** GET ****
//-------------------------------------------------------------------

// GetReservation is the struct returned on a request for a reservation.
type GetReservation struct {
	Status string `json:"status"`
	Record ResDet `json:"record"`
}

//-----------------------------------------------------------------------------
//##########################################################################################################################################################
//-----------------------------------------------------------------------------

// SvcReservationDispatch dispatches a request for a reservation.
//       0    1          2    3 (optional)
// 		/v1/available/BID/RLID
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
			searchReservations(w, r, d)
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

// reservationRowScan scans a result from sql row and dump it in a Reservation struct
func reservationRowScan(rows *sql.Rows, q Reservation) (Reservation, error) {
	err := rows.Scan(
		&q.RLID,
		&q.RID,
		&q.DtStart,
		&q.DtStop,
		&q.FirstName,
		&q.LastName,
		&q.Email,
		&q.Phone,
		&q.ConfirmationCode,
		&q.RentableName,
		&q.RTID,
		&q.Name,
	)
	return q, err
}

// reservationGridFields holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
var reservationGridFieldsMap = map[string][]string{
	"RLID":             {"RentableLeaseStatus.RLID"},
	"RID":              {"RentableLeaseStatus.RID"},
	"DtStart":          {"RentableLeaseStatus.DtStart"},
	"DtStop":           {"RentableLeaseStatus.DtStop"},
	"FirstName":        {"RentableLeaseStatus.FirstName"},
	"LastName":         {"RentableLeaseStatus.LastName"},
	"Email":            {"RentableLeaseStatus.Email"},
	"Phone":            {"RentableLeaseStatus.Phone"},
	"ConfirmationCode": {"RentableLeaseStatus.ConfirmationCode"},
	"RentableName":     {"Rentable.RentableName"},
	"RTID":             {"RentableTypeRef.RTID"},
	"Name":             {"RentableType.Name"},
	//	"RentableType":         {"RT.Name"},
}

// which fields needs to be fetched for SQL query for rentables
var reservationQuerySelectFields = []string{
	"RentableLeaseStatus.RLID",
	"RentableLeaseStatus.RID",
	"RentableLeaseStatus.DtStart",
	"RentableLeaseStatus.DtStop",
	"RentableLeaseStatus.FirstName",
	"RentableLeaseStatus.LastName",
	"RentableLeaseStatus.Email",
	"RentableLeaseStatus.Phone",
	"RentableLeaseStatus.ConfirmationCode",
	"Rentable.RentableName",
	"RentableTypeRef.RTID",
	"RentableTypes.Name",
}

// searchReservations
// wsdoc {
//  @Title  SaveReservation
//	@URL /v1/reservation/:BUI/
//  @Method  POST
//	@Synopsis Returns a list of reservations matching the supplied criteria
//  @Description  Saves the ReleaseStatus. If RLID is 0, a new status is created.
//  @Description  If RLID is > 0 it is simply updated
//	@Input WebGridSearchRequest
//  @Response Reservation
// wsdoc }
//------------------------------------------------------------------------------
func searchReservations(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "searchReservations"
	var err error
	var g SearchReservationResponse
	var currentTime = time.Now()

	rlib.Console("Entered %s\n", funcname)

	limitClause := 100
	srch := fmt.Sprintf(`RentableLeaseStatus.BID = %d
		AND RentableTypeRef.DtStart < %q
		AND RentableTypeRef.DtStop > %q
		AND RentableTypes.FLAGS & 8 = 0
		AND RentableLeaseStatus.LeaseStatus=2
		AND RentableLeaseStatus.DtStart < %q
		AND RentableLeaseStatus.DtStop > %q`,
		d.BID,
		d.wsSearchReq.SearchDtStop.Format(rlib.RRDATETIMESQL),
		d.wsSearchReq.SearchDtStart.Format(rlib.RRDATETIMESQL),
		d.wsSearchReq.SearchDtStop.Format(rlib.RRDATETIMESQL),
		d.wsSearchReq.SearchDtStart.Format(rlib.RRDATETIMESQL))

	order := "RentableLeaseStatus.DtStart ASC,RentableLeaseStatus.LastName ASC,RentableLeaseStatus.Email ASC" // default ORDER is by start date

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, reservationGridFieldsMap)
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	// Rentables Query Text Template
	reservationsQuery := `
SELECT
	{{.SelectClause}}
FROM
    RentableTypeRef
        LEFT JOIN
    RentableTypes ON (RentableTypeRef.RTID = RentableTypes.RTID)
		LEFT JOIN
	RentableLeaseStatus ON (RentableLeaseStatus.RID = RentableTypeRef.RID)
		LEFT JOIN
	Rentable ON (Rentable.RID = RentableTypeRef.RID)
WHERE
	{{.WhereClause}}
ORDER BY
	{{.OrderClause}}` // don't add ';', later some parts will be added in query

	// will be substituted as query clauses
	qc := rlib.QueryClause{
		"SelectClause": strings.Join(reservationQuerySelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
		"currentTime":  currentTime.Format(rlib.RRDATEFMTSQL),                 // show associated instance(s) active as of current time
		"searchStart":  d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), // selected range start
		"searchStop":   d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL),  // selected range stop
		"BID":          strconv.FormatInt(d.BID, 10),
	}

	// GET TOTAL COUNT OF RESULTS
	countQuery := rlib.RenderSQLQuery(reservationsQuery, qc)
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
	reservationsQueryWithLimit := reservationsQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(reservationsQueryWithLimit, qc)
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
		var q Reservation
		q.Recid = i
		// q.BID = rlib.XJSONBud(fmt.Sprintf("%d", d.BID))

		// get records in q struct
		q, err = reservationRowScan(rows, q)
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

// saveReservation
// wsdoc {
//  @Title  SaveReservation
//	@URL /v1/available/:BUI/[RLID]
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
		RLID:             res.RLID,
		RID:              res.RID,
		BID:              d.BID, // DEBUG: res.BID is 0 in ws func test 41b ?????
		DtStart:          time.Time(res.DtStart),
		DtStop:           time.Time(res.DtStop),
		LeaseStatus:      res.LeaseStatus,
		Comment:          res.Comment,
		FirstName:        res.FirstName,
		LastName:         res.LastName,
		Email:            res.Email,
		Phone:            res.Phone,
		Address:          res.Street,
		City:             res.City,
		State:            res.State,
		PostalCode:       res.PostalCode,
		Country:          res.Country,
		CCName:           res.CCName,
		CCType:           res.CCType,
		CCNumber:         res.CCNumber,
		CCExpMonth:       res.CCExpMonth,
		ConfirmationCode: rlib.GenerateUserRefNo(),
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

// getReservation
// wsdoc {
//  @Title  Get Reservation
//	@URL /v1/reservation/:BUI/[RLID]
//  @Method  POST
//	@Synopsis Returns a reservation associated with the supplied RLID
//  @Description  Saves the ReleaseStatus. If RLID is 0, a new status is created.
//  @Description  If RLID is > 0 it is simply updated
//	@Input WebGridSearchRequest
//  @Response Reservation
// wsdoc }
//------------------------------------------------------------------------------
func getReservation(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getReservation"
	var g GetReservation
	var a rlib.RentableLeaseStatus
	var err error

	rlib.Console("entered %s, getting RLID = %d\n", funcname, d.ID)
	a, err = rlib.GetRentableLeaseStatus(r.Context(), d.ID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	b, err := rlib.GetRentableTypeRefForDate(r.Context(), a.RID, &a.DtStart)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.RLID > 0 {
		var gg ResDet
		rlib.MigrateStructVals(&a, &gg)
		gg.RTID = b.RTID
		gg.Recid = gg.RLID
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
