package ws

import (
	"database/sql"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// RRGrid is a structure specifically for the Web Services interface to build a
// Statements grid.
type RRGrid struct {
	Recid           int64           `json:"recid"` // this is to support the w2ui form
	BID             int64           // Business (so that we can process by Business)
	RID             int64           // The rentable
	RTID            int64           // The rentable type
	RentableName    string          // Name of the rentable
	RTName          string          // Name of the rentable type
	RentCycle       int64           // Rent Cycle
	RARID           rlib.NullInt64  // Rental Agreement Rentable id where
	RAID            rlib.NullInt64  // Rental Agreement
	AgreementPeriod string          // text representation of Rental Agreement time period
	AgreementStart  rlib.JSONDate   // start date for RA
	AgreementStop   rlib.JSONDate   // stop date for RA
	UsePeriod       string          // text representation of Occupancy(or use) time period
	PossessionStart rlib.JSONDate   // start date for Occupancy
	PossessionStop  rlib.JSONDate   // stop date for Occupancy
	RentPeriod      string          // text representation of Rent time period
	RentStart       rlib.JSONDate   // start date for Rent
	RentStop        rlib.JSONDate   // stop date for Rent
	Payors          rlib.NullString // payors list attached with this RA within same time
	Users           rlib.NullString // users associated with the rentable
	Sqft            int64
	Description     string
	GSR             float64
	PeriodGSR       string
	IncomeOffsets   float64
	AmountDue       float64
	PaymentsApplied float64
	BeginningRcv    float64
	ChangeInRcv     float64
	EndingRcv       float64
	BeginningSecDep float64
	ChangeInSecDep  float64
	EndingSecDep    float64
}

// RRSearchResponse is the response data for a Rental Agreement Search
type RRSearchResponse struct {
	Status  string   `json:"status"`
	Total   int64    `json:"total"`
	Records []RRGrid `json:"records"`
}

// rrGridFieldsMap holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
var rrGridFieldsMap = map[string][]string{
	"BID":          {"Rentable.BID"},                   // Rentable
	"RID":          {"Rentable.RID"},                   // Rentable
	"RentableName": {"Rentable.RentableName"},          // Rentable
	"RTID":         {"RentableTypeRef.RTID"},           // RentableTypeRef
	"RTName":       {"RentableTypes.Name"},             // RentableTypes
	"RentCycle":    {"RentableTypes.RentCycle"},        // RentableTypes
	"RARID":        {"RentalAgreementRentables.RARID"}, // RentalAgreementRentables
	"RAID":         {"RentalAgreementRentables.RAID"},  // RentalAgreementRentables
	// "PossessionStart": {"RentalAgreement.PossessionStart"},
	// "PossessionStop":  {"RentalAgreement.PossessionStop"},
	// "RentStart":       {"RentalAgreement.RentStart"},
	// "RentStop":        {"RentalAgreement.RentStop"},
	//"Payors":          {"Transactant.FirstName", "Transactant.LastName", "Transactant.CompanyName"},
}

// which fields needs to be fetched for SQL query
var rrQuerySelectFields = []string{
	"Rentable.BID",
	"Rentable.RID",
	"Rentable.RentableName",
	"RentableTypeRef.RTID",
	"RentableTypes.Name",
	"RentableTypes.RentCycle",
	"RentalAgreementRentables.RARID",
	"RentalAgreementRentables.RAID",
	// "RentalAgreement.PossessionStart",
	// "RentalAgreement.PossessionStop",
	// "RentalAgreement.RentStart",
	// "RentalAgreement.RentStop",
	//"GROUP_CONCAT(DISTINCT CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END SEPARATOR ', ') AS Payors",
}

// rrRowScan scans a result from sql row and dump it in a RRGrid struct
func rrRowScan(rows *sql.Rows, q RRGrid) (RRGrid, error) {
	err := rows.Scan(&q.BID, &q.RID, &q.RentableName, &q.RTID, &q.RTName, &q.RentCycle, &q.RARID, &q.RAID /*, &q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop, &q.Payors*/)
	return q, err
}

// SvcRR is the response data for a RR Grid search - The Rent Roll View
func SvcRR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcSearchHandlerRentalAgr"
		err      error
		g        RRSearchResponse
	)

	rlib.Console("Entered %s\n", funcname)

	const (
		limitClause int = 100
	)

	srch := fmt.Sprintf("Rentable.BID=%d", d.BID)                       // default WHERE clause
	order := "Rentable.RentableName ASC "                               // default ORDER
	whereClause, orderClause := GetSearchAndSortSQL(d, rrGridFieldsMap) // establish the order to use in the query
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	// Rental Agreement Query Text Template
	rentalAgrQuery := `
	SELECT {{.SelectClause}}
	FROM Rentable
	LEFT JOIN RentableTypeRef ON Rentable.RID=RentableTypeRef.RID
	LEFT JOIN RentableTypes ON RentableTypeRef.RTID=RentableTypes.RTID
	LEFT JOIN RentalAgreementRentables ON RentalAgreementRentables.RID=Rentable.RID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}` // don't add ';', later some parts will be added in query

	// will be substituted as query clauses
	qc := queryClauses{
		"SelectClause": strings.Join(rrQuerySelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := renderSQLQuery(rentalAgrQuery, qc)
	g.Total, err = GetQueryCount(countQuery, qc)
	if err != nil {
		rlib.Console("Error from GetQueryCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("g.Total = %d\n", g.Total)

	limitAndOffsetClause := `LIMIT {{.LimitClause}} OFFSET {{.OffsetClause}};` // FETCH the records WITH LIMIT AND OFFSET
	rentalAgrQueryWithLimit := rentalAgrQuery + limitAndOffsetClause           // build query with limit and offset clause
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)
	qry := renderSQLQuery(rentalAgrQueryWithLimit, qc) // get formatted query with substitution of select, where, order clause
	rlib.Console("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q RRGrid
		q.Recid = i
		q.BID = d.BID

		// get records info in struct q
		q, err = rrRowScan(rows, q)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
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
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}
