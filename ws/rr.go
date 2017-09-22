package ws

import (
	"database/sql"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// LEFT JOIN Assessments ON (Assessments.RID=Rentable.RID AND Assessments.Start>="{{.DtStart}}" AND Assessments.Stop<"{{.DtStop}}" AND (Assessments.RentCycle=0 OR (Assessments.RentCycle>0 AND Assessments.PASMID!=0)))
// LEFT JOIN AR ON (Assessments.ARID=AR.ARID)
// "ASMID":           {"Assessments.ASMID"},
// "Description":     {"AR.Name"},
// "Assessments.ASMID",
// "AR.Name as Description",

// RRGrid is a structure specifically for the Web Services interface to build a
// Statements grid.
type RRGrid struct {
	Recid           int64           `json:"recid"` // this is to support the w2ui form
	BID             int64           // Business (so that we can process by Business)
	RID             int64           // The rentable
	RTID            int64           // The rentable type
	RARID           rlib.NullInt64  // rental agreement rentable id
	RentableName    rlib.NullString // Name of the rentable
	RentableType    rlib.NullString // Name of the rentable type
	RentCycle       rlib.NullInt64  // Rent Cycle
	Status          rlib.NullInt64  // Rentable status
	RAID            rlib.NullInt64  // Rental Agreement
	ASMID           rlib.NullInt64  // Assessment
	AgreementPeriod string          // text representation of Rental Agreement time period
	AgreementStart  rlib.NullDate   // start date for RA
	AgreementStop   rlib.NullDate   // stop date for RA
	UsePeriod       string          // text representation of Occupancy(or use) time period
	PossessionStart rlib.NullDate   // start date for Occupancy
	PossessionStop  rlib.NullDate   // stop date for Occupancy
	RentPeriod      string          // text representation of Rent time period
	RentStart       rlib.NullDate   // start date for Rent
	RentStop        rlib.NullDate   // stop date for Rent
	Payors          rlib.NullString // payors list attached with this RA within same time
	Users           rlib.NullString // users associated with the rentable
	Sqft            rlib.NullInt64  // rentable sq ft
	Description     rlib.NullString
	GSR             rlib.NullFloat64
	PeriodGSR       rlib.NullFloat64
	IncomeOffsets   rlib.NullFloat64
	AmountDue       rlib.NullFloat64
	PaymentsApplied rlib.NullFloat64
	BeginningRcv    rlib.NullFloat64
	ChangeInRcv     rlib.NullFloat64
	EndingRcv       rlib.NullFloat64
	BeginningSecDep rlib.NullFloat64
	ChangeInSecDep  rlib.NullFloat64
	EndingSecDep    rlib.NullFloat64
	IsSubTotalRow   bool
	IsBlankRow      bool
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
	"BID":             {"Rentable.BID"},                   // Rentable
	"RID":             {"Rentable.RID"},                   // Rentable
	"RentableName":    {"Rentable.RentableName"},          // Rentable
	"RTID":            {"RentableTypeRef.RTID"},           // RentableTypeRef
	"RentableType":    {"RentableTypes.Name"},             // RentableTypes
	"RentCycle":       {"RentableTypes.RentCycle"},        // Rent Cycle
	"RARID":           {"RentalAgreementRentables.RARID"}, // RentalAgreementRentables
	"RAID":            {"RentableMarketRate.MarketRate"},  // GSR
	"MarketRate":      {"RentalAgreementRentables.RAID"},  // RentalAgreementRentables
	"Status":          {"RentableStatus.Status"},          // unit status
	"Payors":          {"Payor.FirstName", "Payor.LastName", "Payor.CompanyName"},
	"Users":           {"User.FirstName", "User.LastName", "User.CompanyName"},
	"PossessionStart": {"RentalAgreement.PossessionStart"},
	"PossessionStop":  {"RentalAgreement.PossessionStop"},
	"RentStart":       {"RentalAgreement.RentStart"},
	"RentStop":        {"RentalAgreement.RentStop"},
}

// which fields needs to be fetched for SQL query
var rrQuerySelectFields = []string{
	"Rentable.BID",
	"Rentable.RID",
	"Rentable.RentableName",
	"RentableTypeRef.RTID",
	"RentableTypes.Name as RentableType",
	"RentableTypes.RentCycle",
	"RentableMarketRate.MarketRate as MarketRate",
	"RentalAgreementRentables.RARID",
	"RentalAgreementRentables.RAID as RAID",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"RentableStatus.Status",
	"GROUP_CONCAT(DISTINCT CASE WHEN Payor.IsCompany > 0 THEN Payor.CompanyName ELSE CONCAT(Payor.FirstName, ' ', Payor.LastName) END SEPARATOR ', ') AS Payors",
	"GROUP_CONCAT(DISTINCT CASE WHEN User.IsCompany > 0 THEN User.CompanyName ELSE CONCAT(User.FirstName, ' ', User.LastName) END SEPARATOR ', ') AS Users",
}

var rentableAsmRcptFields = []string{
	"AR.Name as Description",
	"Assessments.Amount as AmountDue",
	"SUM(ReceiptAllocation.Amount) as PaymentsApplied",
}

// rrRowScan scans a result from sql row and dump it in a RRGrid struct
func rrRowScan(rows *sql.Rows, q RRGrid) (RRGrid, error) {
	err := rows.Scan(&q.BID, &q.RID, &q.RentableName, &q.RTID, &q.RentableType, &q.RentCycle, &q.GSR, &q.RARID, &q.RAID,
		&q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop, &q.Status, &q.Payors, &q.Users)
	return q, err
}

// SvcRR is the response data for a RR Grid search - The Rent Roll View
func SvcRR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcRR"
		err      error
		g        RRSearchResponse
		xbiz     rlib.XBusiness
		custom   = "Square Feet"
	)
	limitClause := d.wsSearchReq.Limit
	if limitClause == 0 {
		limitClause = 100
	}

	rlib.Console("Entered %s\n", funcname)
	rlib.InitBizInternals(d.BID, &xbiz)

	srch := fmt.Sprintf("Rentable.BID=%d", d.BID)                       // default WHERE clause
	order := "Rentable.RentableName ASC "                               // default ORDER
	whereClause, orderClause := GetSearchAndSortSQL(d, rrGridFieldsMap) // establish the order to use in the query
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	rentalAgrQuery := `
	SELECT
		{{.SelectClause}}
	FROM Rentable
	INNER JOIN RentableTypeRef ON RentableTypeRef.RID=Rentable.RID
	INNER JOIN RentableTypes ON RentableTypes.RTID=RentableTypeRef.RTID
	INNER JOIN RentableMarketRate ON (RentableMarketRate.RTID=RentableTypeRef.RTID AND RentableMarketRate.DtStart<"{{.DtStop}}" AND RentableMarketRate.DtStop>"{{.DtStart}}")
	INNER JOIN RentableStatus ON (RentableStatus.RID=Rentable.RID AND RentableStatus.DtStart<"{{.DtStop}}" AND RentableStatus.DtStop>"{{.DtStart}}")
	LEFT JOIN RentalAgreementRentables ON (RentalAgreementRentables.RID=Rentable.RID AND RentalAgreementRentables.RARDtStart<"{{.DtStop}}" AND RentalAgreementRentables.RARDtStop>"{{.DtStart}}")
	LEFT JOIN RentalAgreement ON (RentalAgreement.RAID=RentalAgreementRentables.RAID)
	LEFT JOIN RentalAgreementPayors ON (RentalAgreementRentables.RAID=RentalAgreementPayors.RAID AND RentalAgreementPayors.DtStart<"{{.DtStop}}" AND RentalAgreementPayors.DtStop>"{{.DtStart}}")
	LEFT JOIN Transactant as Payor ON (Payor.TCID=RentalAgreementPayors.TCID AND Payor.BID=Rentable.BID)
	LEFT JOIN RentableUsers ON (RentableUsers.RID=Rentable.RID AND RentableUsers.DtStart<"{{.DtStop}}" AND RentableUsers.DtStop>"{{.DtStart}}")
	LEFT JOIN Transactant as User ON (RentableUsers.TCID=User.TCID AND User.BID=Rentable.BID)
	WHERE {{.WhereClause}}
	GROUP BY Rentable.RID
	ORDER BY {{.OrderClause}}`

	asmRcptQuery := `
	SELECT {{.SelectClause}}
	FROM Rentable
	LEFT JOIN Assessments ON (Assessments.RID=Rentable.RID AND (Assessments.FLAGS & 4)=0 AND "{{.DtStart}}" <= Start AND Stop < "{{.DtStop}}" AND (RentCycle=0 OR (RentCycle>0 AND PASMID!=0)))
	LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.ASMID=Assessments.ASMID AND "{{.DtStart}}" <= ReceiptAllocation.Dt AND ReceiptAllocation.Dt < "{{.DtStop}}")
	LEFT JOIN Receipt ON ( Receipt.RCPTID=ReceiptAllocation.RCPTID AND (Receipt.FLAGS & 4)=0)
	LEFT JOIN AR ON AR.ARID=Assessments.ARID
	WHERE {{.WhereClause}}
	GROUP BY Assessments.ASMID
	ORDER BY {{.OrderClause}};`

	asmRcptQC := queryClauses{
		"SelectClause": strings.Join(rentableAsmRcptFields, ","),
		"OrderClause":  "Assessments.Amount DESC",
		"WhereClause":  "", // later we'll evaluate it
		"DtStart":      d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL),
		"DtStop":       d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL),
	}

	// will be substituted as query clauses
	qc := queryClauses{
		"SelectClause": strings.Join(rrQuerySelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
		"DtStart":      d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL),
		"DtStop":       d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL),
	}

	/*// get TOTAL COUNT First
	countQuery := renderSQLQuery(rentalAgrQuery, qc)
	g.Total, err = GetQueryCount(countQuery, qc)
	if err != nil {
		rlib.Console("Error from GetQueryCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("g.Total = %d\n", g.Total)*/

	// FETCH the records WITH LIMIT AND OFFSET
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`
	rentalAgrQueryWithLimit := rentalAgrQuery + limitAndOffsetClause // build query with limit and offset clause
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
		q.BID = d.BID

		// get records info in struct q
		q, err = rrRowScan(rows, q)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		// fill out more...
		if len(xbiz.RT[q.RTID].CA) > 0 { // if there are custom attributes
			c, ok := xbiz.RT[q.RTID].CA[custom] // see if Square Feet is among them
			if ok {                             // if it is...
				sqft, err := rlib.IntFromString(c.Value, "invalid sqft of custom attribute")
				// q.Sqft.Valid = true
				q.Sqft.Scan(sqft)
				if err != nil {
					SvcGridErrorReturn(w, err, funcname)
					return
				}
			}
		}
		if q.RentStart.Time.Year() > 1970 {
			q.RentPeriod = fmt.Sprintf("%s<br> - %s", q.RentStart.Time.Format(rlib.RRDATEFMT3), q.RentStop.Time.Format(rlib.RRDATEFMT3))
		}
		if q.PossessionStart.Time.Year() > 1970 {
			q.UsePeriod = q.PossessionStart.Time.Format(rlib.RRDATEFMT3) + "<br> - " + q.PossessionStop.Time.Format(rlib.RRDATEFMT3)
		}

		//------------------------------------------------------------
		// now get assessment, receipt related info
		//------------------------------------------------------------
		asmRcptQC["WhereClause"] = fmt.Sprintf("Rentable.BID=%d AND Rentable.RID=%d", q.BID, q.RID)
		arQry := renderSQLQuery(asmRcptQuery, asmRcptQC) // get formatted query with substitution of select, where, order clause
		rlib.Console("Rentable : Assessment + Receipt AMOUNT db query = %s\n", arQry)

		//------------------------------------------------------------
		// There may be multiple rows, hold each row RRGrid in slice
		// Also, compute sobtotals as we go
		//------------------------------------------------------------
		d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

		var rentableResult = []RRGrid{}
		var sub RRGrid
		sub.IsSubTotalRow = true
		sub.AmountDue.Valid = true
		sub.PaymentsApplied.Valid = true
		sub.PeriodGSR.Valid = true
		sub.IncomeOffsets.Valid = true
		// execute the query
		arRows, err := rlib.RRdb.Dbrr.Query(arQry)
		arCount := 0
		if err == nil {
			for arRows.Next() {
				if arCount > 0 { // if more than one rows per rentable then create new RRGrid struct
					var nq = RRGrid{RID: q.RID}
					err = arRows.Scan(&nq.Description, &nq.AmountDue, &nq.PaymentsApplied)
					if err != nil {
						SvcGridErrorReturn(w, err, funcname)
						return
					}
					nq.Recid = g.Total
					rentableResult = append(rentableResult, nq)
					updateSubTotals(&sub, &nq)
				} else {
					err = arRows.Scan(&q.Description, &q.AmountDue, &q.PaymentsApplied)
					if err != nil {
						SvcGridErrorReturn(w, err, funcname)
						return
					}
					q.Recid = g.Total
					rentableResult = append(rentableResult, q)
					updateSubTotals(&sub, &q)
				}
				arCount++
				g.Total++ // grid rows count
			}

			//----------------------------------------
			// Add the Rentable receivables totals...
			//----------------------------------------
			sub.Description.String = "Subtotal"
			sub.Description.Valid = true
			sub.BeginningRcv.Float64, sub.EndingRcv.Float64, err = rlib.GetBeginEndRARBalance(q.RID, q.RAID.Int64, &d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
			sub.ChangeInRcv.Float64 = sub.EndingRcv.Float64 - sub.BeginningRcv.Float64
			rlib.Console("raid=%d, rid=%d, %.2f - %.2f\n", q.RAID.Int64, q.RID, sub.BeginningRcv.Float64, sub.EndingRcv.Float64)
			rlib.Console("CHANGE = %.2f\n", sub.ChangeInRcv.Float64)
			sub.BeginningRcv.Valid = true
			sub.EndingRcv.Valid = true
			sub.ChangeInRcv.Valid = true

			//----------------------------------------
			// Add the Security Deposit totals...
			//----------------------------------------
			sub.BeginningSecDep.Float64, err = rlib.GetSecDepBalance(q.BID, q.RAID.Int64, q.RID, &d1, &d.wsSearchReq.SearchDtStart)
			if err != nil {
				SvcGridErrorReturn(w, err, funcname)
				return
			}
			sub.BeginningSecDep.Valid = true
			sub.ChangeInSecDep.Float64, err = rlib.GetSecDepBalance(q.BID, q.RAID.Int64, q.RID, &d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
			if err != nil {
				SvcGridErrorReturn(w, err, funcname)
				return
			}
			sub.ChangeInSecDep.Valid = true
			sub.EndingSecDep.Float64 = sub.BeginningSecDep.Float64 + sub.ChangeInSecDep.Float64
			sub.EndingSecDep.Valid = true

			sub.Recid = g.Total
			rentableResult = append(rentableResult, sub)
			// arCount++
			g.Total++ // grid rows count
			// add new blank row for grid
			rentableResult = append(rentableResult, RRGrid{IsBlankRow: true, Recid: g.Total})
			// arCount++
			g.Total++ // grid rows count
		}
		arRows.Close()

		g.Records = append(g.Records, rentableResult...)
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

// updateSubTotals does all subtotal calculations for the subtotal line
func updateSubTotals(sub, q *RRGrid) {
	sub.AmountDue.Float64 += q.AmountDue.Float64
	sub.PaymentsApplied.Float64 += q.PaymentsApplied.Float64
	sub.PeriodGSR.Float64 += q.PeriodGSR.Float64
	sub.IncomeOffsets.Float64 += q.IncomeOffsets.Float64
	rlib.Console("\t q.Description = %s, q.AmountDue = %.2f, q.PaymentsApplied = %.2f\n", q.Description, q.AmountDue.Float64, q.PaymentsApplied.Float64)
	rlib.Console("\t sub.AmountDue = %.2f, sub.PaymentsApplied = %.2f\n", sub.AmountDue.Float64, sub.PaymentsApplied.Float64)
}
