package ws

import (
	"context"
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

//-------------------------------------------------------------------
//                        **** SEARCH ****
//-------------------------------------------------------------------

// Reservation defines the timerange, the type of rentable, and the specific
// rentable being reserved.
type Reservation struct {
	Recid               int64             `json:"recid"`
	RLID                int64             // rentable lease status id (reservation id)
	RID                 int64             // specific rentable reserved
	RTID                int64             // the rentable type
	TCID                int64             `json:"TCID"`
	RAID                int64             `json:"RAID"`
	UnspecifiedAdults   int               `json:"UnspecifiedAdults"`
	UnspecifiedChildren int               `json:"UnspecifiedChildren"`
	ConfirmationCode    string            // reservation ConfirmationCode
	DtStart             rlib.JSONDateTime // res start time
	DtStop              rlib.JSONDateTime // res stop time
	IsCompany           bool              `json:"IsCompany"`
	CompanyName         string            `json:"CompanyName"`
	FirstName           string            // res name
	LastName            string            // res name
	Email               string            // email on reservation
	Phone               string            // phone on reservation
	RentableName        string            // rentable name
	Name                string            // Rentable Type Name
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
	Recid               int64             `json:"recid"`               //
	BID                 int64             `json:"rdBID"`               //
	TCID                int64             `json:"TCID"`                // Transactant
	RAID                int64             `json:"RAID"`                //
	DtStart             rlib.JSONDateTime `json:"DtStart"`             //
	DtStop              rlib.JSONDateTime `json:"DtStop"`              //
	RLID                int64             `json:"RLID"`                //
	RTRID               int64             `json:"RTRID"`               //
	RTID                int64             `json:"rdRTID"`              //
	RID                 int64             `json:"RID"`                 //
	Rate                float64           `json:"Rate"`                // base room rate (default amount in default AR for RTID)
	DBAmount            rlib.NullFloat64  `json:"DBAmount"`            // amount being charged for the rentable
	Amount              float64           `json:"Amount"`              // deposit on the rentable
	Deposit             float64           `json:"Deposit"`             // deposit on the rentable
	DepASMID            int64             `json:"DepASMID"`            // deposit assessment
	Discount            float64           `json:"Discount"`            // discount rate
	LeaseStatus         int64             `json:"LeaseStatus"`         //
	Nights              int64             `json:"Nights"`              //
	UnspecifiedAdults   int               `json:"UnspecifiedAdults"`   //
	UnspecifiedChildren int               `json:"UnspecifiedChildren"` //
	RentableName        string            `json:"RentableName"`        //
	IsCompany           bool              `json:"IsCompany"`           // Transactant
	CompanyName         string            `json:"CompanyName"`         // Transactant
	FirstName           string            `json:"FirstName"`           // Transactant
	MiddleName          string            `json:"MiddleName"`          // Transactant
	LastName            string            `json:"LastName"`            // Transactant
	Email               string            `json:"Email"`               // Transactant
	Phone               string            `json:"Phone"`               // Transactant
	Street              string            `json:"Street"`              // Transactant
	City                string            `json:"City"`                // Transactant
	Country             string            `json:"Country"`             // Transactant
	State               string            `json:"State"`               // Transactant
	PostalCode          string            `json:"PostalCode"`          // Transactant
	CCName              string            `json:"CCName"`
	CCType              string            `json:"CCType"`
	CCNumber            string            `json:"CCNumber"`
	CCExpMonth          string            `json:"CCExpMonth"`
	CCExpYear           string            `json:"CCExpYear"`
	ConfirmationCode    string            `json:"ConfirmationCode"`
	Comment             string            `json:"Comment"`
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
		&q.RAID,
		&q.TCID,
		&q.UnspecifiedAdults,
		&q.UnspecifiedChildren,
		&q.IsCompany,
		&q.CompanyName,
		&q.FirstName,
		&q.LastName,
		&q.DtStart,
		&q.DtStop,
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
	"RLID":                {"RentableLeaseStatus.RLID"},
	"RID":                 {"RentableLeaseStatus.RID"},
	"RAID":                {"RentableLeaseStatus.RAID"},
	"TCID":                {"RentalAgreementPayors.TCID"},
	"UnspecifiedAdults":   {"RentalAgreement.UnspecifiedAdults"},
	"UnspecifiedChildren": {"RentalAgreement.UnspecifiedChildren"},
	"IsCompany":           {"Transactant.IsCompany"},
	"CompanyName":         {"Transactant.CompanyName"},
	"FirstName":           {"Transactant.FirstName"},
	"LastName":            {"Transactant.LastName"},
	"DtStart":             {"RentableLeaseStatus.DtStart"},
	"DtStop":              {"RentableLeaseStatus.DtStop"},
	"Email":               {"Transactant.PrimaryEmail"},
	"Phone":               {"Transactant.CellPhone"},
	"ConfirmationCode":    {"RentableLeaseStatus.ConfirmationCode"},
	"RentableName":        {"Rentable.RentableName"},
	"RTID":                {"RentableTypeRef.RTID"},
	"Name":                {"RentableType.Name"},
	//	"RentableType":         {"RT.Name"},
}

// which fields needs to be fetched for SQL query for rentables
var reservationQuerySelectFields = []string{
	"RentableLeaseStatus.RLID",
	"RentableLeaseStatus.RID",
	"RentableLeaseStatus.RAID",
	"RentalAgreementPayors.TCID",
	"RentalAgreement.UnspecifiedAdults",
	"RentalAgreement.UnspecifiedChildren",
	"Transactant.IsCompany",
	"Transactant.CompanyName",
	"Transactant.FirstName",
	"Transactant.LastName",
	"RentableLeaseStatus.DtStart",
	"RentableLeaseStatus.DtStop",
	"Transactant.PrimaryEmail",
	"Transactant.CellPhone",
	"RentableLeaseStatus.ConfirmationCode",
	"Rentable.RentableName",
	"RentableTypeRef.RTID",
	"RentableTypes.Name",
}

// searchReservations
// wsdoc {
//  @Title  Search Reservations
//	@URL /v1/reservation/:BUI/
//  @Method  POST
//	@Synopsis Returns a list of reservations matching the supplied criteria
//  @Description
//  @Description
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

	order := "RentableLeaseStatus.DtStart ASC,Transactant.LastName ASC,Transactant.PrimaryEmail ASC" // default ORDER is by start date

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
		LEFT JOIN
	RentalAgreement ON (RentalAgreement.RAID = RentableLeaseStatus.RAID)
		LEFT JOIN
	RentalAgreementPayors ON (RentalAgreementPayors.RAID = RentableLeaseStatus.RAID)
		Left JOIN
	Transactant ON (Transactant.TCID = RentalAgreementPayors.TCID)
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

// getReservationStruct
//------------------------------------------------------------------------------
func getReservationStruct(id int64) (ResDet, error) {
	var err error
	var a ResDet

	q := fmt.Sprintf(`SELECT
	RentableLeaseStatus.RLID,
    RentableLeaseStatus.RID,
    RentableLeaseStatus.RAID,
    RentalAgreementPayors.TCID,
	RentalAgreement.UnspecifiedAdults,
	RentalAgreement.UnspecifiedChildren,
	AR.DefaultAmount,
    Transactant.IsCompany,
    Transactant.CompanyName,
    Transactant.FirstName,
    Transactant.LastName,
    RentableLeaseStatus.DtStart,
    RentableLeaseStatus.DtStop,
    Transactant.PrimaryEmail,
    Transactant.CellPhone,
    Transactant.Address,
    Transactant.City,
    Transactant.State,
    Transactant.PostalCode,
    RentableLeaseStatus.Comment,
    RentableLeaseStatus.ConfirmationCode,
    Rentable.RentableName,
    RentableTypeRef.RTID,
	Assessments.ASMID,
	Assessments.Amount
FROM
    RentableTypeRef
        LEFT JOIN
    RentableTypes ON (RentableTypeRef.RTID = RentableTypes.RTID)
		LEFT JOIN
	RentableLeaseStatus ON (RentableLeaseStatus.RID = RentableTypeRef.RID)
		LEFT JOIN
	Rentable ON (Rentable.RID = RentableTypeRef.RID)
		LEFT JOIN
	RentalAgreement ON (RentalAgreement.RAID = RentableLeaseStatus.RAID)
		LEFT JOIN
	RentalAgreementPayors ON (RentalAgreementPayors.RAID = RentableLeaseStatus.RAID)
        Left JOIN
	Transactant ON (Transactant.TCID = RentalAgreementPayors.TCID)
		LEFT JOIN
	AR on (RentableTypes.ARID = AR.ARID)
		LEFT JOIN
	Assessments on (Assessments.RAID = RentableLeaseStatus.RAID)
WHERE
	RentableLeaseStatus.RLID = %d;`, id)
	rlib.Console("Query = %s\n", q)
	row := rlib.RRdb.Dbrr.QueryRow(q)
	err = row.Scan(
		&a.RLID,
		&a.RID,
		&a.RAID,
		&a.TCID,
		&a.UnspecifiedAdults,
		&a.UnspecifiedChildren,
		&a.DBAmount,
		&a.IsCompany,
		&a.CompanyName,
		&a.FirstName,
		&a.LastName,
		&a.DtStart,
		&a.DtStop,
		&a.Email,
		&a.Phone,
		&a.Street,
		&a.City,
		&a.State,
		&a.PostalCode,
		&a.Comment,
		&a.ConfirmationCode,
		&a.RentableName,
		&a.RTID,
		&a.DepASMID,
		&a.Deposit,
	)
	if a.DBAmount.Valid {
		a.Amount = a.DBAmount.Float64
	}
	return a, err
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
	var err error
	var a ResDet
	rlib.Console("entered %s, getting RLID = %d\n", funcname, d.ID)

	if a, err = getReservationStruct(d.ID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	g.Status = "success"
	g.Record = a
	SvcWriteResponse(d.BID, &g, w)
}

// saveReservation
// The steps to do this are:
// 1. If this is updating an existing reservation (RLID > 0) read the current
//    version in the database first.
// 	  a. If either the start or stop time moved, then free up the rentable
//       during the old timeslot.
// 2. Write the new RentableLeaseStatus
//
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
	var oldRls rlib.RentableLeaseStatus
	var dtOrRIDchanged bool
	var err error

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

	//---------------------------------------------------
	// Read the Reservation Form data from the client
	//---------------------------------------------------
	var res, resOld ResDet
	err = json.Unmarshal([]byte(s), &res)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	rlib.Console("UnspecifiedAdults = %d, UnspecifiedChildren = %d\n", res.UnspecifiedAdults, res.UnspecifiedChildren)

	now := rlib.Now()
	dt := time.Time(res.DtStart).AddDate(0, 0, 1) // give it one day grace period, which will account for all timezone issues

	if res.RLID > 0 {
		if resOld, err = getReservationStruct(res.RLID); err != nil {
			SvcErrorReturn(w, err, funcname)
		}
	}

	if now.After(dt) {
		err = fmt.Errorf("You cannot create reservations in the past")
		SvcErrorReturn(w, err, funcname)
		return
	}

	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//-------------------------------------------
	// create / update transactant info...
	//-------------------------------------------
	var t rlib.Transactant
	rlib.MigrateStructVals(&res, &t)
	t.PrimaryEmail = res.Email
	t.CellPhone = res.Phone
	t.Address = res.Street
	t.IsCompany = res.IsCompany
	t.CompanyName = res.CompanyName
	if res.TCID > 0 {
		err = updateResTransactant(ctx, r, d, &res, &t)
	} else {

		_, err = rlib.InsertTransactant(ctx, &t)
	}
	if err != nil {
		e := fmt.Errorf("Error saving Transactant:  %s", err.Error())
		tx.Rollback()
		SvcErrorReturn(w, e, funcname)
		return
	}

	//----------------------------------------------------------
	// Create / update Rental Agreement...
	//----------------------------------------------------------
	var ra rlib.RentalAgreement
	if res.RAID > 0 {
		updateResRentalAgreement(ctx, r, d, &res, &resOld, &t, &ra)
	} else {
		insertResRentalAgreement(ctx, r, d, &res, &resOld, &t, &ra)
	}
	//----------------------------------------------------------
	// Create the Rentable Lease Status
	//----------------------------------------------------------
	rlib.Console("res.Comment = %s\n", res.Comment)
	var rls = rlib.RentableLeaseStatus{
		RLID:             res.RLID,
		RID:              res.RID,
		BID:              d.BID, // DEBUG: res.BID is 0 in ws func test 41b ?????
		RAID:             ra.RAID,
		DtStart:          time.Time(res.DtStart),
		DtStop:           time.Time(res.DtStop),
		LeaseStatus:      res.LeaseStatus,
		Comment:          res.Comment,
		ConfirmationCode: res.ConfirmationCode,
	}

	if res.RLID > 0 {
		if oldRls, err = rlib.GetRentableLeaseStatus(ctx, res.RLID); err != nil {
			e := fmt.Errorf("Error saving Transactant:  %s", err.Error())
			tx.Rollback()
			SvcErrorReturn(w, e, funcname)
			return
		}

		//-----------------------------------------------------------------------------
		// If the reservation time changes or the Rentable changes, free the old slot
		//-----------------------------------------------------------------------------
		dtOrRIDchanged = !rls.DtStart.Equal(oldRls.DtStart) || !rls.DtStop.Equal(oldRls.DtStop) || rls.RID != oldRls.RID
		if dtOrRIDchanged {
			var x = rlib.RentableLeaseStatus{
				RID:         oldRls.RID,
				BID:         oldRls.BID,
				LeaseStatus: rlib.LEASESTATUSnotleased,
				DtStart:     oldRls.DtStart,
				DtStop:      oldRls.DtStop,
			}
			if err = rlib.SetRentableLeaseStatus(ctx, &x); err != nil {
				e := fmt.Errorf("Error in SetRentableLeaseStatus:  %s", err.Error())
				tx.Rollback()
				SvcErrorReturn(w, e, funcname)
				return
			}
		}
	}

	if rls.RLID == 0 || dtOrRIDchanged {
		if res.RLID > 0 {
			rls.ConfirmationCode = oldRls.ConfirmationCode
		} else {
			rls.ConfirmationCode = rlib.GenerateUserRefNo()
		}

		err = rlib.SetRentableLeaseStatus(ctx, &rls)
		if err != nil {
			e := fmt.Errorf("Error in SetRentableLeaseStatus:  %s", err.Error())
			tx.Rollback()
			SvcErrorReturn(w, e, funcname)
			return
		}
	} else {
		if err = rlib.UpdateRentableLeaseStatus(ctx, &rls); err != nil {
			e := fmt.Errorf("Error in SetRentableLeaseStatus:  %s", err.Error())
			tx.Rollback()
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	//----------------------------------
	// All done, commit and exit
	//----------------------------------
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponseWithID(rls.BID, w, rls.RLID)
}

// updateResTransactant
// Update the transactant based on the information in the res struct
//
// INPUTS
//    ctx - database context
//    r   - the http request
//    d   - service data
//    res - the data from the Reservation Form
//    t   - transactant struct prefilled with data from res
//
// RETURNS
//    any error encountered
//------------------------------------------------------------------------------
func updateResTransactant(ctx context.Context, r *http.Request, d *ServiceData, res *ResDet, t *rlib.Transactant) error {
	var t0 rlib.Transactant
	rlib.Console("Entered updateResTransactant\n")

	//------------------------------------------------------------------
	// update existing transactant but don't destroy any fields that
	// are not available in the reservation form
	//------------------------------------------------------------------
	if err := rlib.GetTransactant(ctx, t.TCID, &t0); err != nil {
		return err
	}
	count := 0
	if t.IsCompany != t0.IsCompany {
		t0.IsCompany = t.IsCompany
		count++
	}
	if t.CompanyName != t0.CompanyName {
		t0.CompanyName = t.CompanyName
		count++
	}
	if t.FirstName != t0.FirstName {
		t0.FirstName = t.FirstName
		count++
	}
	if t.MiddleName != t0.MiddleName {
		t0.MiddleName = t.MiddleName
		count++
	}
	if t.LastName != t0.LastName {
		t0.LastName = t.LastName
		count++
	}
	if t.PrimaryEmail != t0.PrimaryEmail {
		t0.PrimaryEmail = t.PrimaryEmail
		count++
	}
	if t.CellPhone != t0.CellPhone {
		t0.CellPhone = t.CellPhone
		count++
	}
	if t.Address != t0.Address {
		t0.Address = t.Address
		count++
	}
	if t.City != t0.City {
		t0.City = t.City
		count++
	}
	if t.Country != t0.Country {
		t0.Country = t.Country
		count++
	}
	if t.State != t0.State {
		t0.State = t.State
		count++
	}
	if t.PostalCode != t0.PostalCode {
		t0.PostalCode = t.PostalCode
		count++
	}
	if count > 0 {
		return rlib.UpdateTransactant(ctx, &t0)
	}
	return nil // if nothing changed, no update was necessary
}

func initRAfromReservation(ra *rlib.RentalAgreement, res *ResDet, d *ServiceData) {
	rlib.Console("Entered initRAfromReservation\n")
	now := rlib.Now()

	x := time.Time(res.DtStart)
	dt1 := time.Date(x.Year(), x.Month(), x.Day(), 15, 0, 0, 0, rlib.RRdb.Zone) // check in at 3:00pm
	x = time.Time(res.DtStop)
	epoch := time.Date(x.Year(), x.Month(), x.Day(), 0, 0, 0, 0, rlib.RRdb.Zone) // midnight in this timezone
	dt2 := time.Date(x.Year(), x.Month(), x.Day(), 11, 0, 0, 0, rlib.RRdb.Zone)  // check in at 11:00am
	(*ra) = rlib.RentalAgreement{
		BID:                 res.BID,
		AgreementStart:      now,
		AgreementStop:       dt2,
		PossessionStart:     dt1,
		PossessionStop:      dt2,
		RentStart:           dt1,
		RentStop:            dt2,
		RentCycleEpoch:      epoch,
		FLAGS:               0,
		CSAgent:             d.sess.UID,
		UnspecifiedAdults:   int64(res.UnspecifiedAdults),
		UnspecifiedChildren: int64(res.UnspecifiedChildren),
	}
}

// insertResRentalAgreement
// Create a new rental agreement for the reservation
//
// INPUTS
//    ctx - database context
//    r   - the http request
//    d   - service data
//    res - the data from the Reservation Form
//    resOld - existing res info, if resOld.RLID == 0 then ignore
//    t   - transactant struct prefilled with data from res
//
// RETURNS
//    any error encountered
//------------------------------------------------------------------------------
func insertResRentalAgreement(ctx context.Context, r *http.Request, d *ServiceData, res, resOld *ResDet, t *rlib.Transactant, ra *rlib.RentalAgreement) error {
	var err error
	rlib.Console("Entered insertResRentalAgreement\n")
	initRAfromReservation(ra, res, d)
	rlib.Console("New RA: UnspecifiedAdults = %d, UnspecifiedChildren = %d\n", ra.UnspecifiedAdults, ra.UnspecifiedChildren)
	if _, err = rlib.InsertRentalAgreement(ctx, ra); err != nil {
		return err
	}

	//----------------------------------------------------------
	// Create RentalAgreementRentable
	//----------------------------------------------------------
	var rar rlib.RentalAgreementRentable
	rar.BID = res.BID
	rar.RARDtStart = time.Time(res.DtStart)
	rar.RARDtStop = time.Time(res.DtStop)
	rar.RAID = ra.RAID
	rar.RID = res.RID
	if _, err = rlib.InsertRentalAgreementRentable(ctx, &rar); err != nil {
		return err
	}

	//----------------------------------------------------------
	// Create Payor
	//----------------------------------------------------------
	var rap rlib.RentalAgreementPayor
	rap.BID = res.BID
	rap.DtStart = time.Time(res.DtStart)
	rap.DtStop = time.Time(res.DtStop)
	rap.RAID = ra.RAID
	rap.TCID = t.TCID
	if _, err = rlib.InsertRentalAgreementPayor(ctx, &rap); err != nil {
		return err
	}

	//----------------------------------------------------------
	// Create Deposit assessment
	//----------------------------------------------------------
	if res.Deposit > float64(0) {
		var bp rlib.BizProps
		if bp, err = rlib.GetDataFromBusinessPropertyName(ctx, "general", res.BID); err != nil {
			return err
		}
		now := rlib.Now()
		var a = rlib.Assessment{
			BID:            ra.BID,
			RID:            res.RID,
			RAID:           ra.RAID,
			Amount:         res.Deposit,
			Start:          now,
			Stop:           now,
			RentCycle:      rlib.RECURNONE,
			ProrationCycle: rlib.RECURNONE,
			ARID:           bp.ResDepARID,
		}
		rlib.Console("InsertAssessment.  a.RAID = %d\n", a.RAID)
		if be := bizlogic.InsertAssessment(ctx, &a, 0, &noClose); len(be) > 0 {
			return bizlogic.BizErrorListToError(be)
		}
	}

	return nil
}

// cancelReservation
// Void the rental agreement for this reservation, reverse its assessments,
// and free up its rentables.
//
// INPUTS
//    ctx - database context
//    r   - the http request
//    d   - service data
//    res - the data from the Reservation Form
//
// RETURNS
//    any error encountered
//------------------------------------------------------------------------------
func cancelReservation(ctx context.Context, r *http.Request, d *ServiceData, res *ResDet) error {
	funcname := "cancelReservation"
	var err error
	var ra rlib.RentalAgreement
	rlib.Console("Entered %s\n", funcname)

	//--------------------------------------------------------------------------
	// Terminate the Rental Agreement, mark the reason as Reservation Canceled.
	// This call also reverses all assessments associated with the RA
	//--------------------------------------------------------------------------
	if ra, err = rlib.GetRentalAgreement(ctx, res.RAID); err != nil {
		return err
	}
	return VoidRentalAgreement(ctx, &ra, rlib.MSGRESCANCELED)
}

// updateResRentalAgreement
// Examine the differences between the existing RA and the info supplied in
// res.  Update the existing RA if possible, otherwise delete it and create
// a new one based on what's in res
//
// INPUTS
//    ctx - database context
//    r   - the http request
//    d   - service data
//    res - the data from the Reservation Form
//    t   - transactant struct prefilled with data from res
//
// RETURNS
//    any error encountered
//------------------------------------------------------------------------------
func updateResRentalAgreement(ctx context.Context, r *http.Request, d *ServiceData, res, resOld *ResDet, t *rlib.Transactant, ra *rlib.RentalAgreement) error {
	rlib.Console("Entered updateResRentalAgreement\n")
	var err error
	if err = cancelReservation(ctx, r, d, res); err != nil {
		return err
	}
	if err = insertResRentalAgreement(ctx, r, d, res, resOld, t, ra); err != nil {
		return err
	}
	res.RAID = ra.RAID
	return nil
}

func deleteReservation(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	// const funcname = "deleteReservation"
	// rlib.Console("Entered %s\n", funcname)
	// rls, err := rlib.GetRentableLeaseStatus(r.Context(), d.ID)
}
