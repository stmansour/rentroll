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

// RRGrid is a structure specifically for the Web Services interface to build a
// Statements grid.
type RRGrid struct {
	Recid            int64           `json:"recid"` // this is to support the w2ui form
	BID              int64           // Business (so that we can process by Business)
	RID              int64           // The rentable
	RTID             int64           // The rentable type
	RARID            rlib.NullInt64  // rental agreement rentable id
	RentableName     rlib.NullString // Name of the rentable
	RentableType     rlib.NullString // Name of the rentable type
	RentCycle        rlib.NullInt64  // Rent Cycle
	Status           rlib.NullInt64  // Rentable status
	RAID             rlib.NullInt64  // Rental Agreement
	ASMID            rlib.NullInt64  // Assessment
	AgreementPeriod  string          // text representation of Rental Agreement time period
	AgreementStart   rlib.NullDate   // start date for RA
	AgreementStop    rlib.NullDate   // stop date for RA
	UsePeriod        string          // text representation of Occupancy(or use) time period
	PossessionStart  rlib.NullDate   // start date for Occupancy
	PossessionStop   rlib.NullDate   // stop date for Occupancy
	RentPeriod       string          // text representation of Rent time period
	RentStart        rlib.NullDate   // start date for Rent
	RentStop         rlib.NullDate   // stop date for Rent
	Payors           rlib.NullString // payors list attached with this RA within same time
	Users            rlib.NullString // users associated with the rentable
	Sqft             rlib.NullInt64  // rentable sq ft
	Description      rlib.NullString
	GSR              rlib.NullFloat64
	PeriodGSR        rlib.NullFloat64
	IncomeOffsets    rlib.NullFloat64
	AmountDue        rlib.NullFloat64
	PaymentsApplied  rlib.NullFloat64
	BeginningRcv     rlib.NullFloat64
	ChangeInRcv      rlib.NullFloat64
	EndingRcv        rlib.NullFloat64
	BeginningSecDep  rlib.NullFloat64
	ChangeInSecDep   rlib.NullFloat64
	EndingSecDep     rlib.NullFloat64
	IsMainRow        bool
	IsRentableRow    bool
	IsSubTotalRow    bool
	IsBlankRow       bool
	IsNoRIDAsmtRow   bool
	IsNoRIDNoAsmtRow bool
}

// RRSearchResponse is the response data for a Rental Agreement Search
type RRSearchResponse struct {
	Status        string   `json:"status"`
	Total         int64    `json:"total"`
	Records       []RRGrid `json:"records"`
	TotalMainRows int64    `json:"total_main_rows"`
}

// RRRequeestData - struct for request data for parent-child fashioned rentroll report view
type RRRequeestData struct {
	RentableOffset    int `json:"rentableOffset"`
	NoRIDAsmtOffset   int `json:"noRIDAsmtOffset"`
	NoRIDNoAsmtOffset int `json:"noRIDNoAsmtOffset"`
}

// rrRentablesRowScan scans a result from sql row and dump it in a RRGrid struct
func rrRentablesRowScan(rows *sql.Rows, q *RRGrid) error {
	return rows.Scan(&q.BID, &q.RID, &q.RentableName, &q.RTID, &q.RentableType, &q.RentCycle, &q.GSR, &q.RARID, &q.RAID,
		&q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop, &q.Status, &q.Payors, &q.Users)
}

// ------- Rentables Query components -------

// rentablesFieldsMap holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
// for the rentablesQuery
var rentablesFieldsMap = map[string][]string{
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

// rentablesSelectFields holds the selectClause for the rentablesQuery
var rentablesSelectFields = []string{
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

// rentablesQuery pulls out all rentables records for given date range
// for the rentroll report
var rentablesQuery = `
	SELECT DISTINCT {{.SelectClause}}
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
	ORDER BY {{.OrderClause}};`

// rentablesQueryClause -- query clause for the rentablesQuery
var rentablesQueryClause = queryClauses{
	"SelectClause": strings.Join(rentablesSelectFields, ","),
	"WhereClause":  "Rentable.BID=%d",
	"OrderClause":  "Rentable.RentableName ASC",
	"DtStart":      "",
	"DtStop":       "",
	"BID":          "",
}

// ------- Rentables Assessments Query components -------

var rentablesAsmtFields = []string{
	"AR.Name as Description",
	"Assessments.Amount as AmountDue",
	"SUM(ReceiptAllocation.Amount) as PaymentsApplied",
}

var rentablesAsmtQuery = `
	SELECT {{.SelectClause}}
	FROM Rentable
	LEFT JOIN Assessments ON (Assessments.RID=Rentable.RID AND (Assessments.FLAGS & 4)=0 AND "{{.DtStart}}" <= Start AND Stop < "{{.DtStop}}" AND (RentCycle=0 OR (RentCycle>0 AND PASMID!=0)))
	LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.ASMID=Assessments.ASMID AND "{{.DtStart}}" <= ReceiptAllocation.Dt AND ReceiptAllocation.Dt < "{{.DtStop}}")
	LEFT JOIN AR ON AR.ARID=Assessments.ARID
	WHERE {{.WhereClause}}
	GROUP BY Rentable.RID, Assessments.ASMID
	ORDER BY {{.OrderClause}};`

var rentablesAsmtQueryClause = queryClauses{
	"SelectClause": strings.Join(rentablesAsmtFields, ","),
	"WhereClause":  "Rentable.BID=%d", // needs to be replace %d with some BID in query execution plan
	"OrderClause":  "Rentable.RID ASC, Assessments.Amount DESC",
	"DtStart":      "",
	"DtStop":       "",
	"BID":          "",
}

// ------- Rentables No Assessments Query components -------

var rentablesNoAsmtFields = []string{
	"AR.Name as Description",
	"ReceiptAllocation.RAID as RAID",
	"ReceiptAllocation.Amount as PaymentsApplied",
}

var rentablesNoAsmtQuery = `
	SELECT {{.SelectClause}} FROM ReceiptAllocation
	LEFT JOIN RentalAgreementRentables ON (ReceiptAllocation.RAID=RentalAgreementRentables.RAID AND "{{.DtStart}}" <= RentalAgreementRentables.RARDtStop AND RentalAgreementRentables.RARDtStart < "{{.DtStop}}")
	LEFT JOIN Rentable ON (Rentable.RID=RentalAgreementRentables.RID AND Rentable.RID > 0)
	LEFT JOIN Receipt ON (RentalAgreementRentables.RAID=Receipt.RAID AND Receipt.FLAGS & 4=0 AND "{{.DtStart}}" <= Receipt.Dt AND Receipt.Dt < "{{.DtStop}}")
	INNER JOIN AR ON (AR.ARID = Receipt.ARID AND AR.FLAGS & 5 = 5)
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}};`

var rentablesNoAsmtQueryClause = queryClauses{
	"SelectClause": strings.Join(rentablesNoAsmtFields, ","),
	"WhereClause":  "ReceiptAllocation.BID=%d AND ReceiptAllocation.RAID > 0",
	"OrderClause":  "ReceiptAllocation.Amount DESC",
	"DtStart":      "",
	"DtStop":       "",
	"BID":          "",
}

// ------- No Rentables Assessments Query components -------

var noRIDAsmtQuerySelectFields = []string{
	"Assessments.BID",
	"Assessments.ASMID",
	"AR.Name",
	"Assessments.Amount",
	"SUM(ReceiptAllocation.Amount) as PaymentsApplied",
	"RentalAgreement.RAID",
	"CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END AS Payors",
}

var noRIDAsmtQueryFieldMap = selectQueryFieldMap{
	"Description": {"ARID.Name"},
	"LastModTime": {"PaymentType.LastModTime"},
	"LastModBy":   {"PaymentType.LastModBy"},
	"CreateTS":    {"PaymentType.CreateTS"},
	"CreateBy":    {"PaymentType.CreateBy"},
}

var noRIDAsmtQuery = `
	SELECT {{.SelectClause}}
	FROM Assessments
	LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.ASMID=Assessments.ASMID)
	LEFT JOIN AR ON (AR.ARID=Assessments.ARID)
	LEFT JOIN RentalAgreement ON RentalAgreement.RAID=Assessments.RAID
	LEFT JOIN RentalAgreementPayors ON RentalAgreementPayors.RAID=RentalAgreement.RAID
	LEFT JOIN Transactant ON Transactant.TCID=RentalAgreementPayors.TCID
	WHERE {{.WhereClause}}
	GROUP BY Assessments.ASMID
	ORDER BY {{.OrderClause}};`

var noRIDAsmtQueryClause = queryClauses{
	"SelectClause": strings.Join(noRIDAsmtQuerySelectFields, ","),
	"WhereClause":  "Assessments.BID=%d AND Assessments.FLAGS&4=0 AND Assessments.RID=0 AND %q < Assessments.Stop AND Assessments.Start < %q",
	"OrderClause":  "Assessments.RAID ASC, Assessments.Start ASC",
	"DtStart":      "",
	"DtStop":       "",
	"BID":          "",
}

// ------- No Rentables No Assessments Query components -------

var noRIDNoAsmtQuerySelectFields = []string{
	"ReceiptAllocation.BID",
	"ReceiptAllocation.RAID",
	"ReceiptAllocation.Amount",
	"AR.Name",
	"CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END AS Payors",
}

var noRIDNoAsmQueryFieldMap = selectQueryFieldMap{
	"Description": {"ARID.Name"},
	"LastModTime": {"PaymentType.LastModTime"},
	"LastModBy":   {"PaymentType.LastModBy"},
	"CreateTS":    {"PaymentType.CreateTS"},
	"CreateBy":    {"PaymentType.CreateBy"},
}

var noRIDNoAsmtQuery = `
	SELECT {{.SelectClause}}
	FROM ReceiptAllocation
	LEFT JOIN RentalAgreementRentables ON RentalAgreementRentables.RAID=ReceiptAllocation.RAID
	INNER JOIN Receipt ON (Receipt.RCPTID = ReceiptAllocation.RCPTID)
	INNER JOIN AR ON (AR.ARID = Receipt.ARID AND AR.FLAGS = 5)
	INNER JOIN Transactant ON (Transactant.TCID = Receipt.TCID)
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

var noRIDNoAsmtQueryClause = queryClauses{
	"SelectClause": strings.Join(noRIDNoAsmtQuerySelectFields, ","),
	"WhereClause":  "ReceiptAllocation.BID=%d AND ReceiptAllocation.ASMID=0 AND ReceiptAllocation.RAID>0 AND %q <= ReceiptAllocation.Dt AND ReceiptAllocation.Dt < %q AND RentalAgreementRentables.RID is NULL",
	"OrderClause":  "ReceiptAllocation.Dt ASC",
	"DtStart":      "",
	"DtStop":       "",
	"BID":          "",
}

// GetRentRollReportPartSQLRows returns the sql.Rows for the given looking part
// of the rentroll report
// If given part doesn't exist then it will return nil with error
func GetRentRollReportPartSQLRows(
	rrPart string,
	BID int64,
	d1, d2 time.Time,
	additionalWhere, orderBy string,
	limit, offset int,
) (*sql.Rows, error) {
	const funcname = "GetRentRollReportPartSQLRows"
	var (
		qry   string
		qc    queryClauses
		where string
		order string
		d1Str = d1.Format(rlib.RRDATEFMTSQL)
		d2Str = d2.Format(rlib.RRDATEFMTSQL)
	)
	fmt.Printf("Entered in : %s\n", funcname)

	// based on part, decide query and queryClause
	switch rrPart {
	case "rentables":
		qry = rentablesQuery
		qc = getQueryClauseCopy(rentablesQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID)
		break
	case "rentablesAsmt":
		qry = rentablesAsmtQuery
		qc = getQueryClauseCopy(rentablesAsmtQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID)
		break
	case "rentablesNoAsmt":
		qry = rentablesNoAsmtQuery
		qc = getQueryClauseCopy(rentablesNoAsmtQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID)
		break
	case "noRIDAsmt":
		qry = noRIDAsmtQuery
		qc = getQueryClauseCopy(noRIDAsmtQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID, d1Str, d2Str)
		break
	case "noRIDNoAsmt":
		qry = noRIDNoAsmtQuery
		qc = getQueryClauseCopy(noRIDNoAsmtQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID, d1Str, d2Str)
		break
	default:
		return nil, fmt.Errorf("No such part (%s) exists in rentroll report", rrPart)
	}

	// if additional conditions are provided then append
	if len(additionalWhere) > 0 {
		where += " AND (" + additionalWhere + ")"
	}
	// override orders of query results if it is given
	order = qc["OrderClause"]
	if len(orderBy) > 0 {
		order = orderBy
	}

	// now feed the value in queryclause
	qc["WhereClause"] = where
	qc["OrderClause"] = order
	qc["DtStart"] = d1Str
	qc["DtStop"] = d2Str

	// if limit and offset both are present then
	// we've to add limit and offset clause
	if limit > 0 && offset >= 0 {
		// if query ends with ';' then remove it
		qry = strings.TrimSuffix(strings.TrimSpace(qry), ";")

		// now add LIMIT and OFFSET clause
		qry += ` LIMIT {{.LimitClause}} OFFSET {{.OffsetClause}};`

		// feed the values of limit and offset
		qc["LimitClause"] = strconv.Itoa(limit)
		qc["OffsetClause"] = strconv.Itoa(offset)
	}

	// get formatted query with substitution of select, where, rentablesQOrder clause
	dbQry := renderSQLQuery(qry, qc)
	rlib.Console("db query for %s = %s\n", rrPart, dbQry)

	// return the query execution
	return rlib.RRdb.Dbrr.Query(dbQry)
}

// SvcRR is the response data for a RR Grid search - The Rent Roll View
func SvcRR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcRR"
		err      error
		reqData  RRRequeestData
		g        RRSearchResponse
		xbiz     rlib.XBusiness
		custom   = "Square Feet"
		limit    = d.wsSearchReq.Limit
		startDt  = d.wsSearchReq.SearchDtStart
		stopDt   = d.wsSearchReq.SearchDtStop
	)
	if limit == 0 {
		limit = 20
	}
	rlib.Console("Entered %s\n", funcname)
	if err = json.Unmarshal([]byte(d.data), &reqData); err != nil {
		rlib.Console("Error while unmarshalling d.data: %s\n", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.InitBizInternals(d.BID, &xbiz) // init some business internals first

	//===========================================================
	// TOTAL RECORDS COUNT
	//===========================================================
	rentablesCount, rentablesAsmtCount, rentablesNoAsmtCount, noRIDAsmtCount, noRIDNoAsmtCount, err :=
		getRRTotal(d.BID, startDt, stopDt)

	if err != nil {
		rlib.Console("Error from getRRTotal routine: %s", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("rentablesCount = %d, rentablesAsmtCount = %d, rentablesNoAsmtCount = %d, noRIDAsmtCount = %d, noRIDNoAsmtCount = %d\n", rentablesCount, rentablesAsmtCount, rentablesNoAsmtCount, noRIDAsmtCount, noRIDNoAsmtCount)
	g.Total = (rentablesCount * 3)                 // for each RENTABLES row we'll add subTotal row and one blank row (another two rows)
	if (rentablesAsmtCount - rentablesCount) > 0 { // in case if any rentables get multiple Assessments
		g.Total += (rentablesAsmtCount - rentablesCount)
	}
	if (rentablesNoAsmtCount - rentablesCount) > 0 { // in case if any rentables get multiple result rows
		g.Total += (rentablesNoAsmtCount - rentablesCount)
	}
	g.Total += noRIDAsmtCount                                              // addition of count of assessments which are associated with any rentables
	g.Total += noRIDNoAsmtCount                                            // addition of count of rows which aren't associated with any asmt/rentables
	g.TotalMainRows = (rentablesCount + noRIDAsmtCount + noRIDNoAsmtCount) // main rows count

	//=============================
	// RENTABLES QUERY CALCULATION
	//=============================
	// establish the rentablesQOrder to use in the query
	rentablesQWhereClause, rentablesQOrderClause := GetSearchAndSortSQL(d, rentablesFieldsMap)
	rentablesRows, err := GetRentRollReportPartSQLRows("rentables", d.BID,
		startDt, stopDt,
		rentablesQWhereClause, rentablesQOrderClause,
		limit, reqData.RentableOffset)

	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rentablesRows.Close()

	//================================================================
	//   LOOP THROUGH RENTABLES
	//================================================================
	i := int64(d.wsSearchReq.Offset)
	recidCount := i
	count := 0
	for rentablesRows.Next() {
		//------------------------------------------------------------------
		// load record info into q and fill out what time-based we can...
		//------------------------------------------------------------------
		var q = RRGrid{IsMainRow: true, IsRentableRow: true}
		if err = rrRentablesRowScan(rentablesRows, &q); err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		if len(xbiz.RT[q.RTID].CA) > 0 { // if there are custom attributes
			c, ok := xbiz.RT[q.RTID].CA[custom] // see if Square Feet is among them
			if ok {                             // if it is...
				sqft, err := rlib.IntFromString(c.Value, "invalid sqft of custom attribute")
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
		// There may be multiple rows for the ASSESSMENTS query and
		// the NO-ASSESSMENTS query. Hold each row RRGrid in slice
		// Also, compute subtotals as we go
		//------------------------------------------------------------
		d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
		subList := []RRGrid{}
		sub := RRGrid{IsSubTotalRow: true}
		sub.AmountDue.Valid = true
		sub.PaymentsApplied.Valid = true
		sub.PeriodGSR.Valid = true
		sub.IncomeOffsets.Valid = true

		//================================================================
		//  ASSESSMENTS QUERY...
		//================================================================
		// here we have to apply different whereClause
		// for the rentables Assessment Query as we're looking
		// for ALL assessments for specific rentable
		rentablesAsmtAdditionalWhere := fmt.Sprintf("Rentable.RID=%d", q.RID)
		rentablesAsmtRows, err := GetRentRollReportPartSQLRows("rentablesAsmt", d.BID,
			startDt, stopDt,
			rentablesAsmtAdditionalWhere, "",
			-1, -1)

		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		defer rentablesAsmtRows.Close()

		//================================================================
		//   LOOP THROUGH ASSESSMENTS AND RECEIPTS FOR THIS RENTABLE
		//================================================================
		childCount := 0
		for rentablesAsmtRows.Next() {
			var nq = RRGrid{RID: q.RID, BID: q.BID}
			if childCount == 0 {
				nq = q
			}
			err = rentablesAsmtRows.Scan(&nq.Description, &nq.AmountDue, &nq.PaymentsApplied)
			if err != nil {
				SvcGridErrorReturn(w, err, funcname)
				return
			}
			if nq.Description.Valid || nq.AmountDue.Valid || nq.PaymentsApplied.Valid {
				addToSubList(&subList, &childCount, &recidCount, &nq)
				updateSubTotals(&sub, &nq)
			}
		}

		//================================================================
		//  NO-ASSESSMENTS QUERY...
		//================================================================
		// we need to change whereClause for the rentables no Assessment query
		// as we're looking for ALL payments associated with specific rentable
		// but has no any assessments
		rentablesNoAsmtAdditionalWhere := fmt.Sprintf("RentalAgreementRentables.RID=%d", q.RID)
		rentablesNoAsmtRows, err := GetRentRollReportPartSQLRows("rentablesNoAsmt", d.BID,
			startDt, stopDt,
			rentablesNoAsmtAdditionalWhere, "",
			-1, -1)

		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		defer rentablesNoAsmtRows.Close()

		//================================================================
		//   LOOP THROUGH NO-ASSESSMENTS FOR THIS RENTABLE
		//================================================================
		for rentablesNoAsmtRows.Next() {
			var nq = RRGrid{RID: q.RID, BID: q.BID, Recid: recidCount}
			if childCount == 0 {
				nq = q
			}
			err = rentablesNoAsmtRows.Scan(&nq.Description, &nq.RAID, &nq.PaymentsApplied)
			if err != nil {
				SvcGridErrorReturn(w, err, funcname)
				return
			}
			if nq.Description.Valid || nq.RAID.Valid || nq.PaymentsApplied.Valid {
				addToSubList(&subList, &childCount, &recidCount, &nq)
				updateSubTotals(&sub, &nq)
			}
		}

		//----------------------------------------------------------------------
		// Handle the case where both the Assesments and No-Assessment lists
		// had no matches... just add what we know...
		//----------------------------------------------------------------------
		if len(subList) == 0 {
			addToSubList(&subList, &childCount, &recidCount, &q)
		}

		//----------------------------------------
		// Add the Rentable receivables totals...
		//----------------------------------------
		sub.Description.String = "Subtotal"
		sub.Description.Valid = true
		sub.BeginningRcv.Float64, sub.EndingRcv.Float64, err = rlib.GetBeginEndRARBalance(d.BID, q.RID, q.RAID.Int64, &d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
		sub.BeginningRcv.Valid = true
		sub.ChangeInRcv.Float64 = sub.EndingRcv.Float64 - sub.BeginningRcv.Float64
		sub.ChangeInRcv.Valid = true
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

		addToSubList(&subList, &childCount, &recidCount, &sub)
		addToSubList(&subList, &childCount, &recidCount, &RRGrid{IsBlankRow: true}) // add new blank before the next rentable

		g.Records = append(g.Records, subList...) // update response record list

		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}

	err = rentablesRows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("Added %d Rentable rows\n", recidCount)

	//-------------------------------------------------------------------------
	// Now we need to handle the cases where there are assessments but no
	// associated Rentables...
	//-------------------------------------------------------------------------
	var pageRowsCount int
	pageRowsCount = int(i) - d.wsSearchReq.Offset
	rlib.Console("d.wsSearchReq.Offset: %d, i: %d, pageRowsCount:%d\n", d.wsSearchReq.Offset, i, pageRowsCount)
	rlib.Console("CHECK TO CALL getNoRIDAsmtRows: g.TotalMainRows = %d, g.Total = %d, Limit = %d\n", g.TotalMainRows, g.Total, d.wsSearchReq.Limit)
	if pageRowsCount < d.wsSearchReq.Limit {
		noRIDAsmtLimit := d.wsSearchReq.Limit - pageRowsCount
		rlib.Console("noRIDAsmtLimit:%d\n", noRIDAsmtLimit)
		if noRIDAsmtLimit < 0 {
			noRIDAsmtLimit = 0
		}
		newRecidCount, err := getNoRIDAsmtRows(&g, recidCount, reqData.NoRIDAsmtOffset, noRIDAsmtLimit, d)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		rlib.Console("Added %d noRID with Assessment rows\n", newRecidCount)

		// NEED to update this for rows Count, useful for next sections, for pageRowsCount
		// These are the main rows so we're calculating directly with the recidCount
		i += newRecidCount - recidCount
		pageRowsCount = int(i) - d.wsSearchReq.Offset

		// Now override original counter with new count
		recidCount = newRecidCount
	}

	//-------------------------------------------------------------------------
	// Now we need to handle the cases where there are no assessments as well as
	// no Rentables...
	//-------------------------------------------------------------------------
	rlib.Console("d.wsSearchReq.Offset: %d, i: %d, pageRowsCount:%d\n", d.wsSearchReq.Offset, i, pageRowsCount)
	rlib.Console("CHECK TO CALL getNoRIDNoAsmtRows: g.TotalMainRows = %d, g.Total = %d, Limit = %d\n", g.TotalMainRows, g.Total, d.wsSearchReq.Limit)
	if pageRowsCount < d.wsSearchReq.Limit {
		noRIDNoAsmtlimit := d.wsSearchReq.Limit - pageRowsCount
		rlib.Console("noRIDNoAsmtlimit:%d\n", noRIDNoAsmtlimit)
		if noRIDNoAsmtlimit < 0 {
			noRIDNoAsmtlimit = 0
		}

		newRecidCount, err := getNoRIDNoAsmtRows(&g, recidCount, reqData.NoRIDNoAsmtOffset, noRIDNoAsmtlimit, d)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		rlib.Console("Added %d noRID noASMID rows\n", newRecidCount)

		// NEED to update this for rows Count, useful for next sections, for pageRowsCount
		// These are the main rows so we're calculating directly with the recidCount
		i += newRecidCount - recidCount
		pageRowsCount = int(i) - d.wsSearchReq.Offset

		// Now override original counter with new count
		recidCount = newRecidCount
	}

	fmt.Printf("PageRowsCount: %d, Total: %d, MainTotalRows: %d\n", pageRowsCount, g.Total, g.TotalMainRows)
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// addToSubList is a convenience function that adds a new RRGrid struct to the
// supplied grid struct and updates the
//
// INPUTS
//           g = pointer to a slice of RRGrid structs to which p will be added
//  childCount = pointer to a counter to increment when a record is added
//  recidCount = pointer to a counter of recid values
//-----------------------------------------------------------------------------
func addToSubList(g *[]RRGrid, childCount *int, recidCount *int64, p *RRGrid) {
	p.Recid = *recidCount // first add Recid to RRGrid struct then update
	(*childCount)++
	(*recidCount)++
	*g = append(*g, *p)
}

// updateSubTotals does all subtotal calculations for the subtotal line
//-----------------------------------------------------------------------------
func updateSubTotals(sub, q *RRGrid) {
	sub.AmountDue.Float64 += q.AmountDue.Float64
	sub.PaymentsApplied.Float64 += q.PaymentsApplied.Float64
	sub.PeriodGSR.Float64 += q.PeriodGSR.Float64
	sub.IncomeOffsets.Float64 += q.IncomeOffsets.Float64
	// rlib.Console("\t q.Description = %s, q.AmountDue = %.2f, q.PaymentsApplied = %.2f\n", q.Description, q.AmountDue.Float64, q.PaymentsApplied.Float64)
	// rlib.Console("\t sub.AmountDue = %.2f, sub.PaymentsApplied = %.2f\n", sub.AmountDue.Float64, sub.PaymentsApplied.Float64)
}

// getNoRIDAsmtRows updates g with all Assessments associated with RAIDs but
// no Rentable.
//
// INPUTS
//       g - response struct
//   limit - how many more rows can be added to g
//  offset - recid starts at this amount
//       d - service data
//
// RETURN
//   recidCount - id to be used on next record added
//   error - any error encountered
//-----------------------------------------------------------------------------
func getNoRIDAsmtRows(
	g *RRSearchResponse,
	recidOffset int64,
	queryOffset, limit int,
	d *ServiceData,
) (int64, error) {

	const funcname = "getNoRIDAsmtRows"
	var (
		startDt = d.wsSearchReq.SearchDtStart
		stopDt  = d.wsSearchReq.SearchDtStop
	)
	rlib.Console("Entered %s\n", funcname)

	//--------------------------------------------------
	// How to order
	//--------------------------------------------------
	_, noRIDAsmtOrder := GetSearchAndSortSQL(d, noRIDAsmtQueryFieldMap)

	// TOTAL is already calculated in getRRTotal routine,
	// so NO need to do that here

	//--------------------------------------------------
	// perform the query and process the results
	//--------------------------------------------------
	noRIDAsmtRows, err := GetRentRollReportPartSQLRows("noRIDAsmt", d.BID,
		startDt, stopDt,
		"", noRIDAsmtOrder,
		limit, queryOffset)

	if err != nil {
		return recidOffset, err
	}
	defer noRIDAsmtRows.Close()

	for noRIDAsmtRows.Next() {
		q := RRGrid{Recid: recidOffset, IsMainRow: true, IsNoRIDAsmtRow: true}
		err := noRIDAsmtRows.Scan(&q.BID, &q.ASMID, &q.Description, &q.AmountDue, &q.PaymentsApplied, &q.RAID, &q.Payors)
		if err != nil {
			return recidOffset, err
		}
		g.Records = append(g.Records, q)
		recidOffset++
		// rlib.Console("added: ASMID=%d, AmountDue=%.2f\n", q.ASMID.Int64, q.AmountDue.Float64)
	}
	return recidOffset, nil
}

// getNoRIDNoAsmtRows updates g with all Assessments associated with RAIDs but
// no Rentable.
//
// INPUTS
//       g - response struct
//   limit - how many more rows can be added to g
//  offset - recid starts at this amount
//       d - service data
//
// RETURNS
//   int64 - recidCount, the recid to be used on the next record added
//   error - any error encountered
//-----------------------------------------------------------------------------
func getNoRIDNoAsmtRows(
	g *RRSearchResponse,
	recidOffset int64,
	queryOffset, limit int,
	d *ServiceData,
) (int64, error) {
	const funcname = "getNoRIDNoAsmtRows"
	var (
		startDt = d.wsSearchReq.SearchDtStart
		stopDt  = d.wsSearchReq.SearchDtStop
	)
	rlib.Console("Entered %s\n", funcname)

	//--------------------------------------------------
	// How to order
	//--------------------------------------------------
	_, noRIDNoAsmtOrder := GetSearchAndSortSQL(d, noRIDNoAsmQueryFieldMap)

	// TOTAL is already calculated in getRRTotal routine,
	// so NO need to do that here

	//--------------------------------------------------
	// perform the query and process the results
	//--------------------------------------------------
	noRIDNoAsmtRows, err := GetRentRollReportPartSQLRows("noRIDNoAsmt", d.BID,
		startDt, stopDt,
		"", noRIDNoAsmtOrder,
		limit, queryOffset)

	if err != nil {
		return recidOffset, err
	}
	defer noRIDNoAsmtRows.Close()

	for noRIDNoAsmtRows.Next() {
		q := RRGrid{Recid: recidOffset, IsMainRow: true, IsNoRIDNoAsmtRow: true}
		err := noRIDNoAsmtRows.Scan(&q.BID, &q.RAID, &q.PaymentsApplied, &q.Description, &q.Payors)
		if err != nil {
			return recidOffset, err
		}
		g.Records = append(g.Records, q)
		recidOffset++
	}
	return recidOffset, nil
}

// getRRTotal returns the total of individual total of sections covered by rentroll report
func getRRTotal(
	BID int64,
	d1, d2 time.Time,
) (
	rentablesCount, rentablesAsmtCount, rentablesNoAsmtCount, noRIDAsmtCount, noRIDNoAsmtCount int64,
	err error,
) {

	const funcname = "getRRTotal"
	fmt.Printf("Entered %s\n", funcname)

	// ISSUE: search functionality contains different search scenario, in that case we need to
	// handle case that if field is presents in query FieldMap,.....
	// right now just ignore additional where clause

	// ------------------------
	// Get All Rentables Total
	// ------------------------
	rentablesQC := getQueryClauseCopy(rentablesQueryClause)
	rentablesQC["WhereClause"] = fmt.Sprintf(rentablesQC["WhereClause"], BID)
	rentablesQC["DtStart"] = d1.Format(rlib.RRDATEFMTSQL)
	rentablesQC["DtStop"] = d2.Format(rlib.RRDATEFMTSQL)

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
	rentablesAsmtQC := getQueryClauseCopy(rentablesAsmtQueryClause)
	rentablesAsmtQC["WhereClause"] = fmt.Sprintf(rentablesAsmtQC["WhereClause"], BID)
	rentablesAsmtQC["DtStart"] = rentablesQC["DtStart"]
	rentablesAsmtQC["DtStop"] = rentablesQC["DtStop"]

	rentablesAsmtCountQuery := renderSQLQuery(rentablesAsmtQuery, rentablesAsmtQC)
	// rlib.Console("rentablesAsmtCountQuery db query = %s\n", rentablesAsmtCountQuery)
	rentablesAsmtCount, err = GetQueryCount(rentablesAsmtCountQuery, rentablesAsmtQC)
	if err != nil {
		rlib.Console("Error from rentablesAsmtCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("rentablesAsmtCount = %d\n", rentablesAsmtCount)

	// ----------------------------------------------------------------------
	// Get All Payments associated with Rentables but not with any assessment
	// ----------------------------------------------------------------------
	rentablesNoAsmtQC := getQueryClauseCopy(rentablesNoAsmtQueryClause)
	rentablesNoAsmtQC["WhereClause"] = fmt.Sprintf(rentablesNoAsmtQC["WhereClause"], BID)
	rentablesNoAsmtQC["DtStart"] = rentablesQC["DtStart"]
	rentablesNoAsmtQC["DtStop"] = rentablesQC["DtStop"]

	rentablesNoAsmtCountQuery := renderSQLQuery(rentablesNoAsmtQuery, rentablesNoAsmtQC)
	// rlib.Console("rentablesNoAsmtCountQuery db query = %s\n", rentablesNoAsmtCountQuery)
	rentablesNoAsmtCount, err = GetQueryCount(rentablesNoAsmtCountQuery, rentablesNoAsmtQC)
	if err != nil {
		rlib.Console("Error from rentablesNoAsmtCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("rentablesNoAsmtCount = %d\n", rentablesNoAsmtCount)

	// ---------------------------------------------------------------------
	// Get All Assessments Total which are NOT associated with ANY Rentables
	// ---------------------------------------------------------------------
	noRIDAsmtQC := getQueryClauseCopy(noRIDAsmtQueryClause)
	noRIDAsmtQC["WhereClause"] = fmt.Sprintf(noRIDAsmtQC["WhereClause"], BID, rentablesQC["DtStart"], rentablesQC["DtStop"])
	noRIDAsmtQC["DtStart"] = rentablesQC["DtStart"]
	noRIDAsmtQC["DtStop"] = rentablesQC["DtStop"]

	noRIDAsmtCountQuery := renderSQLQuery(noRIDAsmtQuery, noRIDAsmtQC)
	// rlib.Console("noRIDAsmtCountQuery db query = %s\n", noRIDAsmtCountQuery)
	noRIDAsmtCount, err = GetQueryCount(noRIDAsmtCountQuery, noRIDAsmtQC)
	if err != nil {
		rlib.Console("Error from noRIDAsmtCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("noRIDAsmtCount = %d\n", noRIDAsmtCount)

	// ------------------------------------------------------------------------
	// Get All Payments which are not associated with any Assessement/Rentables
	// ------------------------------------------------------------------------
	noRIDNoAsmtQC := getQueryClauseCopy(noRIDNoAsmtQueryClause)
	noRIDNoAsmtQC["WhereClause"] = fmt.Sprintf(noRIDNoAsmtQC["WhereClause"], BID, rentablesQC["DtStart"], rentablesQC["DtStop"])
	noRIDNoAsmtQC["DtStart"] = rentablesQC["DtStart"]
	noRIDNoAsmtQC["DtStop"] = rentablesQC["DtStop"]

	noRIDNoAsmtCountQuery := renderSQLQuery(noRIDNoAsmtQuery, noRIDNoAsmtQC)
	// rlib.Console("noRIDNoAsmtCountQuery db query = %s\n", noRIDNoAsmtCountQuery)
	noRIDNoAsmtCount, err = GetQueryCount(noRIDNoAsmtCountQuery, noRIDNoAsmtQC)
	if err != nil {
		rlib.Console("Error from noRIDNoAsmtCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("noRIDNoAsmtCount = %d\n", noRIDNoAsmtCount)

	return
}
