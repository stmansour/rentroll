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
	IsMainRow       bool
	IsSubTotalRow   bool
	IsBlankRow      bool
	IsNoRIDRow      bool
	IsRentableRow   bool
}

// RRSearchResponse is the response data for a Rental Agreement Search
type RRSearchResponse struct {
	Status        string   `json:"status"`
	Total         int64    `json:"total"`
	Records       []RRGrid `json:"records"`
	TotalMainRows int64    `json:"total_main_rows"`
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
var rrGridSelectFields = []string{
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

var rentablesAsmtFields = []string{
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

// RRRequeestData - struct for request data for parent-child fashioned rentroll report view
type RRRequeestData struct {
	RentableOffset int `json:"rentableOffset"`
	NoRIDOffset    int `json:"noRIDOffset"`
}

// getRRTotal returns the total of individual total of sections covered by rentroll report
func getRRTotal(BID int64, rentablesQuery string, rentablesQC queryClauses) (
	rentablesCount int64,
	rentablesAsmtCount int64,
	noRIDAsmtCount int64,
	err error,
) {
	var (
		funcname = "getRRTotal"
	)
	fmt.Printf("Entered %s\n", funcname)

	// ------------------------
	// Get All Rentables Total
	// ------------------------
	rentablesCountQuery := renderSQLQuery(rentablesQuery, rentablesQC)
	rentablesCount, err = GetQueryCount(rentablesCountQuery, rentablesQC)
	if err != nil {
		rlib.Console("Error from rentablesCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("rentablesCount = %d\n", rentablesCount)

	// ---------------------------------------------------
	// Get All Assessments Total associated with Rentables
	// ---------------------------------------------------
	totalRentablesAsmtQuery := `
	SELECT
		COUNT(*)
	FROM Rentable
	LEFT JOIN Assessments ON (Assessments.RID=Rentable.RID AND (Assessments.FLAGS & 4)=0 AND "{{.DtStart}}" <= Start AND Stop < "{{.DtStop}}" AND (RentCycle=0 OR (RentCycle>0 AND PASMID!=0)))
	WHERE {{.WhereClause}};`

	totalRentablesAsmtQC := queryClauses{
		"WhereClause": fmt.Sprintf("Rentable.BID=%d", BID), // later we'll evaluate it
		"DtStart":     rentablesQC["DtStart"],
		"DtStop":      rentablesQC["DtStop"],
	}

	rentablesAsmtCountQ := renderSQLQuery(totalRentablesAsmtQuery, totalRentablesAsmtQC)
	// rlib.Console("rentablesAsmtCountQ db query = %s\n", rentablesAsmtCountQ)
	err = rlib.RRdb.Dbrr.QueryRow(rentablesAsmtCountQ).Scan(&rentablesAsmtCount)
	if err != nil {
		rlib.Console("Error from rentablesAsmtCountQ: %s\n", err.Error())
		return
	}
	rlib.Console("rentablesAsmtCount = %d\n", rentablesAsmtCount)

	// ---------------------------------------------------------------------
	// Get All Assessments Total which are NOT associated with ANY Rentables
	// ---------------------------------------------------------------------
	totalAsmtQuery := `
	SELECT
		COUNT(*)
	FROM Assessments
	WHERE {{.WhereClause}};`

	totalAsmtQC := queryClauses{
		"WhereClause": fmt.Sprintf(`Assessments.BID=%d AND Assessments.RID=0 AND Assessments.FLAGS&4=0 AND %q <= Start AND Stop < %q AND (RentCycle=0 OR (RentCycle>0 AND PASMID!=0))`,
			BID, rentablesQC["DtStart"], rentablesQC["DtStop"]),
	}

	totalAsmtCountQ := renderSQLQuery(totalAsmtQuery, totalAsmtQC)
	// rlib.Console("totalAsmtCountQ db query = %s\n", totalAsmtCountQ)
	err = rlib.RRdb.Dbrr.QueryRow(totalAsmtCountQ).Scan(&noRIDAsmtCount)
	if err != nil {
		rlib.Console("Error from totalAsmtCountQ: %s\n", err.Error())
		return
	}
	rlib.Console("noRIDAsmtCount = %d\n", noRIDAsmtCount)

	return
}

// SvcRR is the response data for a RR Grid search - The Rent Roll View
func SvcRR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname    = "SvcRR"
		err         error
		reqData     RRRequeestData
		g           RRSearchResponse
		xbiz        rlib.XBusiness
		custom      = "Square Feet"
		limitClause = d.wsSearchReq.Limit
	)
	if limitClause == 0 {
		limitClause = 20
	}
	rlib.Console("Entered %s\n", funcname)

	// get the request data first
	if err = json.Unmarshal([]byte(d.data), &reqData); err != nil {
		rlib.Console("Error while unmarshalling d.data: %s\n", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	// init some business internals first
	rlib.InitBizInternals(d.BID, &xbiz)

	// ---------------------
	// Rentables Query Part
	// ---------------------
	rentablesQSrch := fmt.Sprintf("Rentable.BID=%d", d.BID)                       // default WHERE clause
	rentablesQOrder := "Rentable.RentableName ASC "                               // default rentablesQOrder
	whereClause, rentablesQOrderClause := GetSearchAndSortSQL(d, rrGridFieldsMap) // establish the rentablesQOrder to use in the query
	if len(whereClause) > 0 {
		rentablesQSrch += " AND (" + whereClause + ")"
	}
	if len(rentablesQOrderClause) > 0 {
		rentablesQOrder = rentablesQOrderClause
	}

	rentablesQuery := `
	SELECT DISTINCT
		{{.SelectClause}}
	FROM Rentable
	LEFT JOIN RentableTypeRef ON RentableTypeRef.RID=Rentable.RID
	LEFT JOIN RentableTypes ON RentableTypes.RTID=RentableTypeRef.RTID
	LEFT JOIN RentableMarketRate ON (RentableMarketRate.RTID=RentableTypeRef.RTID AND RentableMarketRate.DtStart<"{{.DtStop}}" AND RentableMarketRate.DtStop>"{{.DtStart}}")
	LEFT JOIN RentableStatus ON (RentableStatus.RID=Rentable.RID AND RentableStatus.DtStart<"{{.DtStop}}" AND RentableStatus.DtStop>"{{.DtStart}}")
	LEFT JOIN RentalAgreementRentables ON (RentalAgreementRentables.RID=Rentable.RID AND RentalAgreementRentables.RARDtStart<"{{.DtStop}}" AND RentalAgreementRentables.RARDtStop>"{{.DtStart}}")
	LEFT JOIN RentalAgreement ON (RentalAgreement.RAID=RentalAgreementRentables.RAID)
	LEFT JOIN RentalAgreementPayors ON (RentalAgreementRentables.RAID=RentalAgreementPayors.RAID AND RentalAgreementPayors.DtStart<"{{.DtStop}}" AND RentalAgreementPayors.DtStop>"{{.DtStart}}")
	LEFT JOIN Transactant as Payor ON (Payor.TCID=RentalAgreementPayors.TCID AND Payor.BID=Rentable.BID)
	LEFT JOIN RentableUsers ON (RentableUsers.RID=Rentable.RID AND RentableUsers.DtStart<"{{.DtStop}}" AND RentableUsers.DtStop>"{{.DtStart}}")
	LEFT JOIN Transactant as User ON (RentableUsers.TCID=User.TCID AND User.BID=Rentable.BID)
	WHERE {{.WhereClause}}
	GROUP BY Rentable.RID
	ORDER BY {{.OrderClause}}`

	// will be substituted as query clauses
	rentablesQC := queryClauses{
		"SelectClause": strings.Join(rrGridSelectFields, ","),
		"WhereClause":  rentablesQSrch,
		"OrderClause":  rentablesQOrder,
		"DtStart":      d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL),
		"DtStop":       d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL),
	}

	// ------------------------------------
	// Assessments for Rentables Query Part
	// ------------------------------------
	// LEFT JOIN Receipt ON ( Receipt.RCPTID=ReceiptAllocation.RCPTID AND (Receipt.FLAGS & 4)=0)
	rentablesAsmtQuery := `
	SELECT {{.SelectClause}}
	FROM Rentable
	LEFT JOIN Assessments ON (Assessments.RID=Rentable.RID AND (Assessments.FLAGS & 4)=0 AND "{{.DtStart}}" <= Start AND Stop < "{{.DtStop}}" AND (RentCycle=0 OR (RentCycle>0 AND PASMID!=0)))
	LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.ASMID=Assessments.ASMID AND "{{.DtStart}}" <= ReceiptAllocation.Dt AND ReceiptAllocation.Dt < "{{.DtStop}}")
	LEFT JOIN AR ON AR.ARID=Assessments.ARID
	WHERE {{.WhereClause}}
	GROUP BY Assessments.ASMID
	ORDER BY {{.OrderClause}};`

	rentablesAsmtQC := queryClauses{
		"SelectClause": strings.Join(rentablesAsmtFields, ","),
		"OrderClause":  "Assessments.Amount DESC",
		"WhereClause":  "", // later we'll evaluate it
		"DtStart":      d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL),
		"DtStop":       d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL),
	}

	// -------------------
	// TOTAL RECORDS COUNT
	// -------------------
	rentablesCount, rentablesAsmtCount, noRIDAsmtCount, err := getRRTotal(d.BID, rentablesQuery, rentablesQC)
	if err != nil {
		rlib.Console("Error from getRRTotal routine: %s", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("rentablesCount = %d, rentablesAsmtCount = %d, noRIDAsmtCount = %d\n", rentablesCount, rentablesAsmtCount, noRIDAsmtCount)

	// for each RENTABLES row we'll add subTotal row and one blank row (another two rows)
	g.Total = (rentablesCount * 3)
	// in case if any rentables got multiple Assessments
	if (rentablesAsmtCount - rentablesCount) > 0 {
		g.Total += (rentablesAsmtCount - rentablesCount)
	}
	// addition of count of assessments which are associated with any rentables
	g.Total += noRIDAsmtCount
	// main rows count
	g.TotalMainRows = (rentablesCount + noRIDAsmtCount)

	// ------------------------------------
	// Start to fetch rentables rows first
	// WITH LIMIT AND OFFSET
	// ------------------------------------
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`
	rentablesQueryWithLimit := rentablesQuery + limitAndOffsetClause // build query with limit and offset clause
	rentablesQC["LimitClause"] = strconv.Itoa(limitClause)
	rentablesQC["OffsetClause"] = strconv.Itoa(reqData.RentableOffset)
	rentablesQ := renderSQLQuery(rentablesQueryWithLimit, rentablesQC) // get formatted query with substitution of select, where, rentablesQOrder clause
	rlib.Console("db query = %s\n", rentablesQ)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(rentablesQ)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	recidCount := i
	count := 0

	//================================================================
	//   LOOP THROUGH RENTABLES
	//================================================================
	for rows.Next() {
		var q = RRGrid{BID: d.BID, Recid: recidCount, IsMainRow: true, IsRentableRow: true}

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
		rentablesAsmtQC["WhereClause"] = fmt.Sprintf("Rentable.BID=%d AND Rentable.RID=%d", q.BID, q.RID)
		rentablesAsmtQ := renderSQLQuery(rentablesAsmtQuery, rentablesAsmtQC) // get formatted query with substitution of select, where, order clause
		// rlib.Console("Rentable : Assessment + Receipt AMOUNT db query = %s\n", rentablesAsmtQ)

		//------------------------------------------------------------
		// There may be multiple rows, hold each row RRGrid in slice
		// Also, compute sobtotals as we go
		//------------------------------------------------------------
		d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
		rentableASMList := []RRGrid{}
		sub := RRGrid{IsSubTotalRow: true}
		sub.AmountDue.Valid = true
		sub.PaymentsApplied.Valid = true
		sub.PeriodGSR.Valid = true
		sub.IncomeOffsets.Valid = true

		// execute the query
		rentablesAsmtRows, err := rlib.RRdb.Dbrr.Query(rentablesAsmtQ)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		defer rentablesAsmtRows.Close()

		childCount := 0
		if err == nil {

			//================================================================
			//   LOOP THROUGH RENTABLES
			//================================================================

			for rentablesAsmtRows.Next() {
				if childCount > 0 { // if more than one rows per rentable then create new RRGrid struct
					var nq = RRGrid{RID: q.RID, BID: q.BID, Recid: recidCount}
					err = rentablesAsmtRows.Scan(&nq.Description, &nq.AmountDue, &nq.PaymentsApplied)
					if err != nil {
						SvcGridErrorReturn(w, err, funcname)
						return
					}
					rentableASMList = append(rentableASMList, nq)
					updateSubTotals(&sub, &nq)
				} else {
					err = rentablesAsmtRows.Scan(&q.Description, &q.AmountDue, &q.PaymentsApplied)
					if err != nil {
						SvcGridErrorReturn(w, err, funcname)
						return
					}
					rentableASMList = append(rentableASMList, q)
					updateSubTotals(&sub, &q)
				}
				childCount++
				recidCount++
			}

			if len(rentableASMList) == 0 { // that means no assessments found, then just append rentable info
				rentableASMList = append(rentableASMList, q)
			}

			// add list in g.Records field
			g.Records = append(g.Records, rentableASMList...)

			//----------------------------------------
			// Add the Rentable receivables totals...
			//----------------------------------------
			sub.Description.String = "Subtotal"
			sub.Description.Valid = true
			sub.BeginningRcv.Float64, sub.EndingRcv.Float64, err = rlib.GetBeginEndRARBalance(d.BID, q.RID, q.RAID.Int64, &d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
			sub.BeginningRcv.Valid = true
			sub.ChangeInRcv.Float64 = sub.EndingRcv.Float64 - sub.BeginningRcv.Float64
			sub.ChangeInRcv.Valid = true
			// rlib.Console("raid=%d, rid=%d, %.2f - %.2f\n", q.RAID.Int64, q.RID, sub.BeginningRcv.Float64, sub.EndingRcv.Float64)
			// rlib.Console("CHANGE = %.2f\n", sub.ChangeInRcv.Float64)
			sub.EndingRcv.Valid = true

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

			sub.Recid = recidCount
			g.Records = append(g.Records, sub)
			childCount++
			recidCount++

			// add new blank row for grid
			g.Records = append(g.Records, RRGrid{IsBlankRow: true, Recid: recidCount})
			childCount++
			recidCount++
		}

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

	//-------------------------------------------------------------------------
	// Now we need to handle the cases where there are assessments but no
	// associated Rentables...
	//-------------------------------------------------------------------------
	pageRentablesCount := int(i) - d.wsSearchReq.Offset
	rlib.Console("d.wsSearchReq.Offset: %d, i: %d, pageRentablesCount:%d\n", d.wsSearchReq.Offset, i, pageRentablesCount)
	rlib.Console("CHECK TO CALL getNoRentableRows: g.TotalMainRows = %d, g.Total = %d, Limit = %d\n", g.TotalMainRows, g.Total, d.wsSearchReq.Limit)
	if pageRentablesCount < d.wsSearchReq.Limit {
		queryOffset := int64(reqData.NoRIDOffset) // need to work out the calculation for this
		noRIDQlimit := int64(d.wsSearchReq.Limit - pageRentablesCount)
		rlib.Console("noRIDQlimit:%d\n", noRIDQlimit)
		if noRIDQlimit < 0 {
			noRIDQlimit = 0
		}
		err = getNoRentableRows(&g, recidCount, queryOffset, noRIDQlimit, d)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
	}

	fmt.Printf("Total: %d, MainTotalRows: %d\n", g.Total, g.TotalMainRows)
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
	// rlib.Console("\t q.Description = %s, q.AmountDue = %.2f, q.PaymentsApplied = %.2f\n", q.Description, q.AmountDue.Float64, q.PaymentsApplied.Float64)
	// rlib.Console("\t sub.AmountDue = %.2f, sub.PaymentsApplied = %.2f\n", sub.AmountDue.Float64, sub.PaymentsApplied.Float64)
}

// getNoRentableRows updates g with all Assessments associated with RAIDs but
// no Rentable.
//
// INPUTS
//       g - response struct
//   limit - how many more rows can be added to g
//  offset - recid starts at this amount
//       d - service data
//-----------------------------------------------------------------------------
func getNoRentableRows(g *RRSearchResponse, recidoffset, queryOffset, limit int64, d *ServiceData) error {
	funcname := "getNoRentableRows"
	rlib.Console("Entered %s\n", funcname)
	DtStartStr := d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL)
	DtStopStr := d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL)

	//--------------------------------------------------
	// which fields needs to be fetched for SQL query
	//--------------------------------------------------
	var rrNoRIDQuerySelectFields = []string{
		"Assessments.BID",
		"Assessments.ASMID",
		"Assessments.Amount",
		"SUM(ReceiptAllocation.Amount) as PaymentsApplied",
		"RentalAgreement.RAID",
		"GROUP_CONCAT(DISTINCT CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END SEPARATOR ', ') AS Payors",
	}

	//--------------------------------------------------
	// Select the appropriate assessments
	//--------------------------------------------------
	where := fmt.Sprintf("Assessments.BID=%d AND Assessments.FLAGS&4=0 AND Assessments.RID=0 AND %q < Assessments.Stop AND Assessments.Start < %q", d.BID, DtStartStr, DtStopStr)

	//--------------------------------------------------
	// How to order
	//--------------------------------------------------
	order := "Assessments.RAID ASC,Assessments.Start ASC"
	_, orderClause := GetSearchAndSortSQL(d, pmtSearchFieldMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	//--------------------------------------------------
	// The full query...
	//--------------------------------------------------
	noRIDQuery := `
	SELECT {{.SelectClause}}
	FROM Assessments
	LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.ASMID=Assessments.ASMID)
	LEFT JOIN RentalAgreement ON RentalAgreement.RAID=Assessments.RAID
	LEFT JOIN RentalAgreementPayors ON RentalAgreementPayors.RAID=RentalAgreement.RAID
	LEFT JOIN Transactant ON Transactant.TCID=RentalAgreementPayors.TCID
	WHERE {{.WhereClause}}
	GROUP BY Assessments.ASMID
	ORDER BY {{.OrderClause}}`

	qc := queryClauses{
		"SelectClause": strings.Join(rrNoRIDQuerySelectFields, ","),
		"WhereClause":  where,
		"OrderClause":  order,
		"DtStart":      DtStartStr,
		"DtStop":       DtStopStr,
	}

	// get total of this query
	countQuery := renderSQLQuery(noRIDQuery, qc)
	noRIDTotal, err := GetQueryCount(countQuery, qc)
	if err != nil {
		return err
	}
	rlib.Console("noRIDTotal records = %d,  limit = %d\n", noRIDTotal, limit)

	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`
	noRIDQueryWithLimit := noRIDQuery + limitAndOffsetClause
	qc["LimitClause"] = fmt.Sprintf("%d", limit)
	qc["OffsetClause"] = fmt.Sprintf("%d", queryOffset)

	// get formatted query with substitution of select, where, order clause
	qry := renderSQLQuery(noRIDQueryWithLimit, qc)
	rlib.Console("noRID db query = %s\n", qry)

	//--------------------------------------------------
	// perform the query and process the results
	//--------------------------------------------------
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()

	recidCount := int64(recidoffset)
	for rows.Next() {
		q := RRGrid{Recid: recidCount, IsMainRow: true, IsNoRIDRow: true}
		err := rows.Scan(&q.BID, &q.ASMID, &q.AmountDue, &q.PaymentsApplied, &q.RAID, &q.Payors)
		if err != nil {
			return err
		}
		g.Records = append(g.Records, q)
		recidCount++
		// rlib.Console("added: ASMID=%d, AmountDue=%.2f\n", q.ASMID.Int64, q.AmountDue.Float64)
	}
	return nil
}

// getNoAssessmentRows updates g with all Assessments associated with RAIDs but
// no Rentable.
//
// INPUTS
//       g - response struct
//   limit - how many more rows can be added to g
//  offset - recid starts at this amount
//       d - service data
//-----------------------------------------------------------------------------
func getNoAssessmentRows(g *RRSearchResponse, recidoffset, queryOffset, limit int64, d *ServiceData) error {
	funcname := "getNoAssessmentRows"
	rlib.Console("Entered %s\n", funcname)
	DtStartStr := d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL)
	DtStopStr := d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL)

	//--------------------------------------------------
	// which fields needs to be fetched for SQL query
	//--------------------------------------------------
	var rrNoRIDQuerySelectFields = []string{
		"Assessments.BID",
		"Assessments.ASMID",
		"Assessments.Amount",
		"SUM(ReceiptAllocation.Amount) as PaymentsApplied",
		"RentalAgreement.RAID",
		"GROUP_CONCAT(DISTINCT CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END SEPARATOR ', ') AS Payors",
	}

	//--------------------------------------------------
	// Select the appropriate assessments
	//--------------------------------------------------
	where := fmt.Sprintf("Assessments.BID=%d AND Assessments.FLAGS&4=0 AND Assessments.RID=0 AND %q < Assessments.Stop AND Assessments.Start < %q", d.BID, DtStartStr, DtStopStr)

	//--------------------------------------------------
	// How to order
	//--------------------------------------------------
	order := "Assessments.RAID ASC,Assessments.Start ASC"
	_, orderClause := GetSearchAndSortSQL(d, pmtSearchFieldMap)
	if len(orderClause) > 0 {
		order = orderClause
	}

	//--------------------------------------------------
	// The full query...
	//--------------------------------------------------
	noRIDQuery := `
	SELECT {{.SelectClause}}
	FROM Assessments
	LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.ASMID=Assessments.ASMID)
	LEFT JOIN RentalAgreement ON RentalAgreement.RAID=Assessments.RAID
	LEFT JOIN RentalAgreementPayors ON RentalAgreementPayors.RAID=RentalAgreement.RAID
	LEFT JOIN Transactant ON Transactant.TCID=RentalAgreementPayors.TCID
	WHERE {{.WhereClause}}
	GROUP BY Assessments.ASMID
	ORDER BY {{.OrderClause}}`

	qc := queryClauses{
		"SelectClause": strings.Join(rrNoRIDQuerySelectFields, ","),
		"WhereClause":  where,
		"OrderClause":  order,
		"DtStart":      DtStartStr,
		"DtStop":       DtStopStr,
	}

	// get total of this query
	countQuery := renderSQLQuery(noRIDQuery, qc)
	noRIDTotal, err := GetQueryCount(countQuery, qc)
	if err != nil {
		return err
	}
	rlib.Console("noRIDTotal records = %d,  limit = %d\n", noRIDTotal, limit)

	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`
	noRIDQueryWithLimit := noRIDQuery + limitAndOffsetClause
	qc["LimitClause"] = fmt.Sprintf("%d", limit)
	qc["OffsetClause"] = fmt.Sprintf("%d", queryOffset)

	// get formatted query with substitution of select, where, order clause
	qry := renderSQLQuery(noRIDQueryWithLimit, qc)
	rlib.Console("noRID db query = %s\n", qry)

	//--------------------------------------------------
	// perform the query and process the results
	//--------------------------------------------------
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()

	recidCount := int64(recidoffset)
	for rows.Next() {
		q := RRGrid{Recid: recidCount, IsMainRow: true, IsNoRIDRow: true}
		err := rows.Scan(&q.BID, &q.ASMID, &q.AmountDue, &q.PaymentsApplied, &q.RAID, &q.Payors)
		if err != nil {
			return err
		}
		g.Records = append(g.Records, q)
		recidCount++
		// rlib.Console("added: ASMID=%d, AmountDue=%.2f\n", q.ASMID.Int64, q.AmountDue.Float64)
	}
	return nil
}
