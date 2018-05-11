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

// RentalAgr is a structure specifically for the Web Services interface. It will be
// automatically populated from an rlib.RentalAgreement struct. Records of this type
// are returned by the search handler
type RentalAgr struct {
	Recid                  int64 `json:"recid"` // this is to support the w2ui form
	RAID                   int64 // internal unique id
	RATID                  int64 // reference to Occupancy Master Agreement
	BID                    int64 // Business (so that we can process by Business)
	BUD                    rlib.XJSONBud
	NLID                   int64             // Note ID
	AgreementStart         rlib.JSONDate     // start date for rental agreement contract
	AgreementStop          rlib.JSONDate     // stop date for rental agreement contract
	PossessionStart        rlib.JSONDate     // start date for Occupancy
	PossessionStop         rlib.JSONDate     // stop date for Occupancy
	RentStart              rlib.JSONDate     // start date for Rent
	RentStop               rlib.JSONDate     // stop date for Rent
	RentCycleEpoch         rlib.JSONDate     // Date on which rent cycle recurs. Start date for the recurring rent assessment
	UnspecifiedAdults      int64             // adults who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels
	UnspecifiedChildren    int64             // children who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels.
	Renewal                int64             // 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
	SpecialProvisions      string            // free-form text
	LeaseType              int64             // Full Service Gross, Gross, ModifiedGross, Tripple Net
	ExpenseAdjustmentType  int64             // Base Year, No Base Year, Pass Through
	ExpensesStop           float64           // cap on the amount of oexpenses that can be passed through to the tenant
	ExpenseStopCalculation string            // note on how to determine the expense stop
	BaseYearEnd            rlib.JSONDate     // last day of the base year
	ExpenseAdjustment      rlib.JSONDate     // the next date on which an expense adjustment is due
	EstimatedCharges       float64           // a periodic fee charged to the tenant to reimburse LL for anticipated expenses
	RateChange             float64           // predetermined amount of rent increase, expressed as a percentage
	NextRateChange         rlib.JSONDate     // he next date on which a RateChange will occur
	PermittedUses          string            // indicates primary use of the space, ex: doctor's office, or warehouse/distribution, etc.
	ExclusiveUses          string            // those uses to which the tenant has the exclusive rights within a complex, ex: Trader Joe's may have the exclusive right to sell groceries
	ExtensionOption        string            // the right to extend the term of lease by giving notice to LL, ex: 2 options to extend for 5 years each
	ExtensionOptionNotice  rlib.JSONDate     // the last date by which a Tenant can give notice of their intention to exercise the right to an extension option period
	ExpansionOption        string            // the right to expand to certanin spaces that are typically contiguous to their primary space
	ExpansionOptionNotice  rlib.JSONDate     // the last date by which a Tenant can give notice of their intention to exercise the right to an Expansion Option
	RightOfFirstRefusal    string            // Tenant may have the right to purchase their premises if LL chooses to sell
	LastModTime            rlib.JSONDateTime // when was this record last written
	LastModBy              int64             // employee UID (from phonebook) that modified it
	CreateTS               rlib.JSONDateTime // when was this record last written
	CreateBy               int64             // employee UID (from phonebook) that modified it
	Payors                 rlib.NullString   // payors list attached with this RA within same time
}

// RentalAgrForm is used to save a Rental Agreement.  It holds those values
type RentalAgrForm struct {
	Recid                  int64 `json:"recid"` // this is to support the w2ui form
	BID                    int64
	BUD                    rlib.XJSONBud
	RAID                   int64             // internal unique id
	RATID                  int64             // reference to Occupancy Master Agreement
	NLID                   int64             // Note ID
	AgreementStart         rlib.JSONDate     // start date for rental agreement contract
	AgreementStop          rlib.JSONDate     // stop date for rental agreement contract
	PossessionStart        rlib.JSONDate     // start date for Occupancy
	PossessionStop         rlib.JSONDate     // stop date for Occupancy
	RentStart              rlib.JSONDate     // start date for Rent
	RentStop               rlib.JSONDate     // stop date for Rent
	RentCycleEpoch         rlib.JSONDate     // Date on which rent cycle recurs. Start date for the recurring rent assessment
	UnspecifiedAdults      int64             // adults who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels
	UnspecifiedChildren    int64             // children who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels.
	SpecialProvisions      string            // free-form text
	LeaseType              int64             // Full Service Gross, Gross, ModifiedGross, Tripple Net
	ExpenseAdjustmentType  int64             // Base Year, No Base Year, Pass Through
	ExpensesStop           float64           // cap on the amount of oexpenses that can be passed through to the tenant
	ExpenseStopCalculation string            // note on how to determine the expense stop
	BaseYearEnd            rlib.JSONDate     // last day of the base year
	ExpenseAdjustment      rlib.JSONDate     // the next date on which an expense adjustment is due
	EstimatedCharges       float64           // a periodic fee charged to the tenant to reimburse LL for anticipated expenses
	RateChange             float64           // predetermined amount of rent increase, expressed as a percentage
	NextRateChange         rlib.JSONDate     // he next date on which a RateChange will occur
	PermittedUses          string            // indicates primary use of the space, ex: doctor's office, or warehouse/distribution, etc.
	ExclusiveUses          string            // those uses to which the tenant has the exclusive rights within a complex, ex: Trader Joe's may have the exclusive right to sell groceries
	ExtensionOption        string            // the right to extend the term of lease by giving notice to LL, ex: 2 options to extend for 5 years each
	ExtensionOptionNotice  rlib.JSONDate     // the last date by which a Tenant can give notice of their intention to exercise the right to an extension option period
	ExpansionOption        string            // the right to expand to certanin spaces that are typically contiguous to their primary space
	ExpansionOptionNotice  rlib.JSONDate     // the last date by which a Tenant can give notice of their intention to exercise the right to an Expansion Option
	RightOfFirstRefusal    string            // Tenant may have the right to purchase their premises if LL chooses to sell
	Renewal                rlib.XJSONRenewal // month to month automatic extension, or lease extension
}

// RentalAgrSearchResponse is the response data for a Rental Agreement Search
type RentalAgrSearchResponse struct {
	Status  string      `json:"status"`
	Total   int64       `json:"total"`
	Records []RentalAgr `json:"records"`
}

// GetRentalAgreementResponse is the response data for GetRentalAgreement
type GetRentalAgreementResponse struct {
	Status string    `json:"status"`
	Record RentalAgr `json:"record"`
}

// DeleteRentalAgreementForm used while deleteRA request
type DeleteRentalAgreementForm struct {
	RAID int64
}

// rentalAgrGridFieldsMap holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
var rentalAgrGridFieldsMap = map[string][]string{
	"RAID":                   {"RentalAgreement.RAID"},
	"RATID":                  {"RentalAgreement.RATID"},
	"NLID":                   {"RentalAgreement.NLID"},
	"AgreementStart":         {"RentalAgreement.AgreementStart"},
	"AgreementStop":          {"RentalAgreement.AgreementStop"},
	"PossessionStart":        {"RentalAgreement.PossessionStart"},
	"PossessionStop":         {"RentalAgreement.PossessionStop"},
	"RentStart":              {"RentalAgreement.RentStart"},
	"RentStop":               {"RentalAgreement.RentStop"},
	"RentCycleEpoch":         {"RentalAgreement.RentCycleEpoch"},
	"UnspecifiedAdults":      {"RentalAgreement.UnspecifiedAdults"},
	"UnspecifiedChildren":    {"RentalAgreement.UnspecifiedChildren"},
	"Renewal":                {"RentalAgreement.Renewal"},
	"SpecialProvisions":      {"RentalAgreement.SpecialProvisions"},
	"LeaseType":              {"RentalAgreement.LeaseType"},
	"ExpenseAdjustmentType":  {"RentalAgreement.ExpenseAdjustmentType"},
	"ExpensesStop":           {"RentalAgreement.ExpensesStop"},
	"ExpenseStopCalculation": {"RentalAgreement.ExpenseStopCalculation"},
	"BaseYearEnd":            {"RentalAgreement.BaseYearEnd"},
	"ExpenseAdjustment":      {"RentalAgreement.ExpenseAdjustment"},
	"EstimatedCharges":       {"RentalAgreement.EstimatedCharges"},
	"RateChange":             {"RentalAgreement.RateChange"},
	"NextRateChange":         {"RentalAgreement.NextRateChange"},
	"PermittedUses":          {"RentalAgreement.PermittedUses"},
	"ExclusiveUses":          {"RentalAgreement.ExclusiveUses"},
	"ExtensionOption":        {"RentalAgreement.ExtensionOption"},
	"ExtensionOptionNotice":  {"RentalAgreement.ExtensionOptionNotice"},
	"ExpansionOption":        {"RentalAgreement.ExpansionOption"},
	"ExpansionOptionNotice":  {"RentalAgreement.ExpansionOptionNotice"},
	"RightOfFirstRefusal":    {"RentalAgreement.RightOfFirstRefusal"},
	"LastModTime":            {"RentalAgreement.LastModTime"},
	"LastModBy":              {"RentalAgreement.LastModBy"},
	"CreateTS":               {"RentalAgreement.CreateTS"},
	"CreateBy":               {"RentalAgreement.CreateBy"},
	"Payors":                 {"Transactant.FirstName", "Transactant.LastName", "Transactant.CompanyName"},
}

// which fields needs to be fetched for SQL query for rental agreements
var rentalAgrQuerySelectFields = []string{
	"RentalAgreement.RAID",
	"RentalAgreement.RATID",
	"RentalAgreement.NLID",
	"RentalAgreement.AgreementStart",
	"RentalAgreement.AgreementStop",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"RentalAgreement.RentCycleEpoch",
	"RentalAgreement.UnspecifiedAdults",
	"RentalAgreement.UnspecifiedChildren",
	"RentalAgreement.Renewal",
	"RentalAgreement.SpecialProvisions",
	"RentalAgreement.LeaseType",
	"RentalAgreement.ExpenseAdjustmentType",
	"RentalAgreement.ExpensesStop",
	"RentalAgreement.ExpenseStopCalculation",
	"RentalAgreement.BaseYearEnd",
	"RentalAgreement.ExpenseAdjustment",
	"RentalAgreement.EstimatedCharges",
	"RentalAgreement.RateChange",
	"RentalAgreement.NextRateChange",
	"RentalAgreement.PermittedUses",
	"RentalAgreement.ExclusiveUses",
	"RentalAgreement.ExtensionOption",
	"RentalAgreement.ExtensionOptionNotice",
	"RentalAgreement.ExpansionOption",
	"RentalAgreement.ExpansionOptionNotice",
	"RentalAgreement.RightOfFirstRefusal",
	"RentalAgreement.LastModTime",
	"RentalAgreement.LastModBy",
	"RentalAgreement.CreateTS",
	"RentalAgreement.CreateBy",
	"GROUP_CONCAT(DISTINCT CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END ORDER BY Transactant.TCID ASC SEPARATOR ', ') AS Payors",
}

// RentalAgreementTypedown is the struct of data needed for typedown when searching for a RentalAgreement
type RentalAgreementTypedown struct {
	Recid       int64 `json:"recid"`
	TCID        int64
	FirstName   string
	MiddleName  string
	LastName    string
	CompanyName string
	IsCompany   bool
	RAID        int64
}

// RentalAgreementTypedownResponse is the data structure for the response to a search for people
type RentalAgreementTypedownResponse struct {
	Status  string                    `json:"status"`
	Total   int64                     `json:"total"`
	Records []RentalAgreementTypedown `json:"records"`
}

// GetRentalAgreementTypeDown returns the values needed for typedown controls:
// input:   bid - business
//            s - string or substring to search for
//        limit - return no more than this many matches
// return a slice of TransactantTypeDowns and an error.
func GetRentalAgreementTypeDown(bid int64, s string, limit int) ([]RentalAgreementTypedown, error) {
	var m []RentalAgreementTypedown
	s = "%" + s + "%"
	rows, err := rlib.RRdb.Prepstmt.GetRentalAgreementTypeDown.Query(bid, s, s, s, limit)
	if err != nil {
		return m, err
	}
	defer rows.Close()
	for rows.Next() {
		var t RentalAgreementTypedown
		err = rows.Scan(&t.TCID, &t.FirstName, &t.MiddleName, &t.LastName, &t.CompanyName, &t.IsCompany, &t.RAID)
		if err != nil {
			return m, err
		}
		m = append(m, t)
	}
	return m, nil
}

// SvcRentalAgreementTypeDown handles typedown requests for RentalAgreements.  It returns
// the RAID for the associated payor
// wsdoc {
//  @Title  Get Transactants Typedown
//	@URL /v1/ratd/:BUI?request={"search":"The search string","max":"Maximum number of return items"}
//	@Method GET
//	@Synopsis Fast Search for Transactants matching typed characters
//  @Desc Returns TCID, FirstName, Middlename, and LastName of Transactants that
//  @Desc match supplied chars at the beginning of FirstName or LastName
//  @Input WebTypeDownRequest
//  @Response TransactantsTypedownResponse
// wsdoc }
func SvcRentalAgreementTypeDown(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRentalAgreementTypeDown"
	var (
		g   RentalAgreementTypedownResponse
		err error
	)
	rlib.Console("Entered %s\n", funcname)

	rlib.Console("handle typedown: GetRentalAgreementTypeDown( bid=%d, search=%s, limit=%d\n", d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	g.Records, err = GetRentalAgreementTypeDown(d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	rlib.Console("GetRentalAgreementTypeDown returned %d matches\n", len(g.Records))
	g.Total = int64(len(g.Records))
	if err != nil {
		e := fmt.Errorf("Error getting typedown matches: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	for i := 0; i < len(g.Records); i++ {
		g.Records[i].Recid = int64(i)
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// rentalAgrRowScan scans a result from sql row and dump it in a RentalAgr struct
func rentalAgrRowScan(rows *sql.Rows, q RentalAgr) (RentalAgr, error) {
	err := rows.Scan(&q.RAID, &q.RATID, &q.NLID, &q.AgreementStart, &q.AgreementStop, &q.PossessionStart, &q.PossessionStop,
		&q.RentStart, &q.RentStop, &q.RentCycleEpoch, &q.UnspecifiedAdults, &q.UnspecifiedChildren, &q.Renewal, &q.SpecialProvisions,
		&q.LeaseType, &q.ExpenseAdjustmentType, &q.ExpensesStop, &q.ExpenseStopCalculation, &q.BaseYearEnd, &q.ExpenseAdjustment,
		&q.EstimatedCharges, &q.RateChange, &q.NextRateChange, &q.PermittedUses, &q.ExclusiveUses, &q.ExtensionOption,
		&q.ExtensionOptionNotice, &q.ExpansionOption, &q.ExpansionOptionNotice, &q.RightOfFirstRefusal,
		&q.LastModTime, &q.LastModBy, &q.CreateTS, &q.CreateBy, &q.Payors)
	return q, err
}

// SvcSearchHandlerRentalAgr generates a report of all RentalAgreements defined business d.BID
// wsdoc {
//  @Title  Search Rental Agreements
//	@URL /v1/rentalagrs/:BUI
//  @Method  GET, POST
//	@Synopsis Return Rental Agreements that match the criteria provided.
//  @Description
//	@Input WebGridSearchRequest
//  @Response RentalAgrSearchResponse
// wsdoc }
func SvcSearchHandlerRentalAgr(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerRentalAgr"
	var (
		err error
		g   RentalAgrSearchResponse
	)

	rlib.Console("Entered %s\n", funcname)

	const (
		limitClause int = 100
	)

	// srch := fmt.Sprintf("RentalAgreement.BID=%d AND (RentalAgreement.AgreementStop>%q OR RentalAgreement.PossessionStop>%q OR RentalAgreement.RentStop>%q)",
	// 	d.BID, t.Format(rlib.RRDATEINPFMT), t.Format(rlib.RRDATEINPFMT), t.Format(rlib.RRDATEINPFMT)) // default WHERE clause
	srch := fmt.Sprintf("RentalAgreement.BID=%d AND RentalAgreement.AgreementStop>%q AND RentalAgreement.AgreementStart<%q",
		d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL)) // default WHERE clause
	order := "RentalAgreement.RAID ASC" // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, rentalAgrGridFieldsMap)
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	// Rental Agreement Query Text Template
	rentalAgrQuery := `
	SELECT
		{{.SelectClause}}
	FROM RentalAgreement
	LEFT JOIN RentalAgreementPayors ON RentalAgreementPayors.RAID=RentalAgreement.RAID
	LEFT JOIN Transactant ON Transactant.TCID=RentalAgreementPayors.TCID
	WHERE {{.WhereClause}}
	GROUP BY RentalAgreement.RAID
	ORDER BY {{.OrderClause}}` // don't add ';', later some parts will be added in query

	// will be substituted as query clauses
	qc := rlib.QueryClause{
		"SelectClause": strings.Join(rentalAgrQuerySelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(rentalAgrQuery, qc)
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
	rentalAgrQueryWithLimit := rentalAgrQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(rentalAgrQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q RentalAgr
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		// get records info in struct q
		q, err = rentalAgrRowScan(rows, q)
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

// SvcFormHandlerRentalAgreement formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the RAID as follows:
//       0    1          2    3
// 		/v1/RentalAgrs/BID/RAID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerRentalAgreement(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcFormHandlerRentalAgreement"
	var (
		err error
	)
	rlib.Console("Entered %s\n", funcname)

	if d.RAID, err = SvcExtractIDFromURI(r.RequestURI, "RAID", 3, w); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Requester UID = %d, BID = %d,  RAID = %d\n", d.UID, d.BID, d.RAID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getRentalAgreement(w, r, d)
		break
	case "save":
		saveRentalAgreement(w, r, d)
		break
	case "delete":
		deleteRentalAgreement(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// wsdoc {
//  @Title  Save Rental Agreement
//	@URL /v1/rentalagr/:BUI/:RAID
//  @Method  POST
//	@Synopsis Save (create or update) a Rental Agreement
//  @Description This service returns the single-valued attributes of a Rental Agreement.
//	@Input WebGridSearchRequest
//  @Response SvcStatusResponse
// wsdoc }
func saveRentalAgreement(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveRentalAgreement"
	var (
		err error
	)

	target := `"record":`
	rlib.Console("SvcFormHandlerRentalAgreement save\n")
	rlib.Console("record data = %s\n", d.data)
	i := strings.Index(d.data, target)
	rlib.Console("record is at index = %d\n", i)
	if i < 0 {
		e := fmt.Errorf("saveRentalAgreement: cannot find %s in form json", target)
		SvcErrorReturn(w, e, funcname)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	rlib.Console("data to unmarshal is:  %s\n", s)

	//===============================================================
	//------------------------------
	// Handle all the non-list data
	//------------------------------
	var foo RentalAgrForm

	err = json.Unmarshal([]byte(s), &foo)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	// migrate the variables that transfer without needing special handling...
	var a rlib.RentalAgreement
	rlib.MigrateStructVals(&foo, &a)

	rlib.Console("B1\n")

	var ok bool
	a.Renewal, ok = rlib.RenewalMap[string(foo.Renewal)]
	if !ok {
		e := fmt.Errorf("could not map %s to a Renewal value", foo.Renewal)
		rlib.LogAndPrintError(funcname, e)
		SvcErrorReturn(w, e, funcname)
		return
	}

	//===============================================================

	tx, ctx, err := rlib.NewTransactionWithContext(r.Context())
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	// Now just update the database
	if a.RAID > 0 {
		be := bizlogic.UpdateRentalAgreement(ctx, &a)
		if be != nil {
			tx.Rollback()
			err = bizlogic.BizErrorListToError(be)
			e := fmt.Errorf("Error saving Rental Agreement RAID = %d: %s", a.RAID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	} else {
		_, err = rlib.InsertRentalAgreement(ctx, &a)
		if err != nil {
			tx.Rollback()
			SvcErrorReturn(w, err, funcname)
			return
		}
		var lm rlib.LedgerMarker
		lm.Dt = a.AgreementStart
		lm.RAID = a.RAID
		lm.State = rlib.LMINITIAL
		_, err = rlib.InsertLedgerMarker(ctx, &lm)
		if err != nil {
			tx.Rollback()
			e := fmt.Errorf("Error saving Rental Agreement RAID = %d: %s", a.RAID, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWriteSuccessResponseWithID(d.BID, w, a.RAID)
}

// https://play.golang.org/p/gfOhByMroo

// wsdoc {
//  @Title  Get Rental Agreement
//	@URL /v1/rentalagr/:BUI/:RAID
//	@Method POST or GET
//	@Synopsis Get a Rental Agreement
//  @Description This service returns the single-valued attributes of a Rental Agreement.
//  @Input WebGridSearchRequest
//  @Response GetRentalAgreementResponse
// wsdoc }
func getRentalAgreement(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getRentalAgreement"
	var (
		err error
		g   GetRentalAgreementResponse
	)

	rlib.Console("Entered %s\n", funcname)

	a, err := rlib.GetRentalAgreement(r.Context(), d.RAID)
	if err != nil {
		e := fmt.Errorf("getRentalAgreement: cannot read RentalAgreement RAID = %d, err = %s", d.RAID, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	if a.RAID > 0 {
		var gg RentalAgr
		rlib.MigrateStructVals(&a, &gg)
		gg.BUD = rlib.GetBUDFromBIDList(gg.BID)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// wsdoc {
//  @Title  Delete Rental Agreement
//	@URL /v1/rentalagr/:BUI/:RAID
//	@Method POST
//	@Synopsis Delete a Rental Agreement
//  @Description This service delete the requested Rental Agreement with RAID and deletes associated pets, users, payors, references to rentables.
//  @Input DeleteRentalAgreementForm
//  @Response SvcStatusResponse
// wsdoc }
func deleteRentalAgreement(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteRentalAgreement"
	var (
		err error
		del DeleteRentalAgreementForm
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	delRAID := del.RAID

	// first get rentalAgreement
	ra, err := rlib.GetRentalAgreement(r.Context(), delRAID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// remove all pets associated with this rental Agreement
	// TODO(Sudip): better should pass transaction here for batch delete
	if err = rlib.DeleteAllRentalAgreementPets(r.Context(), delRAID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// remove all payors associated with this rental Agreement
	// TODO(Sudip): better should pass transaction here for batch delete
	if err = rlib.DeleteAllRentalAgreementPayors(r.Context(), delRAID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// remove all rentable users associated with this rental Agreement
	// TODO(Sudip): better should start transaction here for batch delete
	rarList, _ := rlib.GetRentalAgreementRentables(r.Context(), delRAID, &ra.AgreementStart, &ra.AgreementStop)
	for _, rar := range rarList {
		rUsers, _ := rlib.GetRentableUsersInRange(r.Context(), rar.RID, &rar.RARDtStart, &rar.RARDtStop)
		for _, ru := range rUsers {
			if err := rlib.DeleteRentableUser(r.Context(), ru.RUID); err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
		}
	}

	// remove all references to rentables associated with this rental Agreement
	// TODO(Sudip): better should start transaction here for batch delete
	if err = rlib.DeleteAllRentalAgreementRentables(r.Context(), delRAID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// finally delete this rental Agreement
	if err = rlib.DeleteRentalAgreement(r.Context(), delRAID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}
