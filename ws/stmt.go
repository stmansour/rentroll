package ws

import (
	"database/sql"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// StmtGrid is a structure specifically for the Web Services interface to build a
// Statements grid.
type StmtGrid struct {
	Recid           int64           `json:"recid"` // this is to support the w2ui form
	RAID            int64           // internal unique id
	BID             int64           // Business (so that we can process by Business)
	BUD             rlib.XJSONBud   // which business
	AgreementStart  rlib.JSONDate   // start date for rental agreement contract
	AgreementStop   rlib.JSONDate   // stop date for rental agreement contract
	PossessionStart rlib.JSONDate   // start date for Occupancy
	PossessionStop  rlib.JSONDate   // stop date for Occupancy
	RentStart       rlib.JSONDate   // start date for Rent
	RentStop        rlib.JSONDate   // stop date for Rent
	Payors          rlib.NullString // payors list attached with this RA within same time
}

// StmtSearchResponse is the response data for a Rental Agreement Search
type StmtSearchResponse struct {
	Status  string     `json:"status"`
	Total   int64      `json:"total"`
	Records []StmtGrid `json:"records"`
}

// stmtGridFieldsMap holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
var stmtGridFieldsMap = map[string][]string{
	"RAID":            {"RentalAgreement.RAID"},
	"AgreementStart":  {"RentalAgreement.AgreementStart"},
	"AgreementStop":   {"RentalAgreement.AgreementStop"},
	"PossessionStart": {"RentalAgreement.PossessionStart"},
	"PossessionStop":  {"RentalAgreement.PossessionStop"},
	"RentStart":       {"RentalAgreement.RentStart"},
	"RentStop":        {"RentalAgreement.RentStop"},
	"Payors":          {"Transactant.FirstName", "Transactant.LastName", "Transactant.CompanyName"},
}

// which fields needs to be fetched for SQL query for rental agreements
var stmtQuerySelectFields = []string{
	"RentalAgreement.RAID",
	"RentalAgreement.AgreementStart",
	"RentalAgreement.AgreementStop",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"GROUP_CONCAT(DISTINCT CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END SEPARATOR ', ') AS Payors",
}

// stmtRowScan scans a result from sql row and dump it in a StmtGrid struct
func stmtRowScan(rows *sql.Rows, q StmtGrid) (StmtGrid, error) {
	err := rows.Scan(&q.RAID, &q.AgreementStart, &q.AgreementStop, &q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop, &q.Payors)
	return q, err
}

// SvcStatement is the response data for a Stmt Grid search
func SvcStatement(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcStatement"
		err      error
		g        StmtSearchResponse
	)

	rlib.Console("Entered %s\n", funcname)

	const (
		limitClause int = 100
	)

	srch := fmt.Sprintf("RentalAgreement.BID=%d AND (RentalAgreement.AgreementStop>%q OR RentalAgreement.PossessionStop>%q OR RentalAgreement.RentStop>%q) AND (RentalAgreement.AgreementStart<%q OR RentalAgreement.PossessionStart<%q OR RentalAgreement.RentStart<%q)",
		d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEINPFMT), d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEINPFMT), d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEINPFMT),
		d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEINPFMT), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEINPFMT), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEINPFMT)) // default WHERE clause
	order := "RentalAgreement.RAID ASC" // default ORDER
	whereClause, orderClause := GetSearchAndSortSQL(d, stmtGridFieldsMap)
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
		"SelectClause": strings.Join(stmtQuerySelectFields, ","),
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
		var q StmtGrid
		q.Recid = i
		q.BID = d.BID
		q.BUD = rlib.GetBUDFromBIDList(q.BID)

		// get records info in struct q
		q, err = stmtRowScan(rows, q)
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
