package rrpt

import (
	"database/sql"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// ------- Rentables Query components -------

// RentablesFieldsMap holds the map of field (to be shown on grid)
// to actual database fields, multiple db fields means combine those
// for the RentablesQuery
var RentablesFieldsMap = rlib.SelectQueryFieldMap{
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

// RentablesSelectFields holds the selectClause for the RentablesQuery
var RentablesSelectFields = rlib.SelectQueryFields{
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
	"GROUP_CONCAT(DISTINCT CASE WHEN Payor.IsCompany > 0 THEN Payor.CompanyName ELSE CONCAT(Payor.FirstName, ' ', Payor.LastName) END ORDER BY Payor.LastName ASC, Payor.FirstName ASC, Payor.CompanyName ASC SEPARATOR ', ') AS Payors",
	"GROUP_CONCAT(DISTINCT CASE WHEN User.IsCompany > 0 THEN User.CompanyName ELSE CONCAT(User.FirstName, ' ', User.LastName) END ORDER BY User.LastName ASC, User.FirstName ASC, User.CompanyName ASC SEPARATOR ', ' ) AS Users",
}

// RentablesQuery pulls out all rentables records for given date range
// for the rentroll report
var RentablesQuery = `
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

// RentablesQueryClause -- query clause for the RentablesQuery
var RentablesQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(RentablesSelectFields, ","),
	"WhereClause":  "Rentable.BID=%d",
	"OrderClause":  "Rentable.RentableName ASC",
	"DtStart":      "",
	"DtStop":       "",
}

// ------- Rentables Assessments Query components -------

// RentablesAsmtFields holds the fields need to be fetched by
// rentables assessment query
var RentablesAsmtFields = rlib.SelectQueryFields{
	"AR.Name as Description",
	"Assessments.RAID",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"Assessments.Amount as AmountDue",
	"SUM(ReceiptAllocation.Amount) as PaymentsApplied",
}

// RentablesAsmtQuery - query execution plan for rentable assessments
var RentablesAsmtQuery = `
    SELECT {{.SelectClause}}
    FROM Rentable
    LEFT JOIN Assessments ON (Assessments.RID=Rentable.RID AND (Assessments.FLAGS & 4)=0 AND "{{.DtStart}}" <= Start AND Stop < "{{.DtStop}}" AND (RentCycle=0 OR (RentCycle>0 AND PASMID!=0)))
    LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.ASMID=Assessments.ASMID AND "{{.DtStart}}" <= ReceiptAllocation.Dt AND ReceiptAllocation.Dt < "{{.DtStop}}")
    LEFT JOIN RentalAgreement on (RentalAgreement.RAID=Assessments.RAID)
    LEFT JOIN AR ON AR.ARID=Assessments.ARID
    WHERE {{.WhereClause}}
    GROUP BY Rentable.RID, Assessments.ASMID
    ORDER BY {{.OrderClause}};`

// RentablesAsmtQueryClause - query clause for rentable assessments
var RentablesAsmtQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(RentablesAsmtFields, ","),
	"WhereClause":  "Rentable.BID=%d", // needs to be replace %d with some BID in query execution plan
	"OrderClause":  "Assessments.RAID ASC, Assessments.Amount DESC",
	"DtStart":      "",
	"DtStop":       "",
}

// ------- Rentables No Assessments Query components -------

// RentablesNoAsmtFields holds the list of fields need to be fetched by
// rentables noasseesment query
var RentablesNoAsmtFields = rlib.SelectQueryFields{
	"AR.Name as Description",
	"ReceiptAllocation.RAID as RAID",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"ReceiptAllocation.Amount as PaymentsApplied",
}

// RentablesNoAsmtQuery - the query execution plan for
// rentables noassessments
var RentablesNoAsmtQuery = `
    SELECT {{.SelectClause}} FROM ReceiptAllocation
    LEFT JOIN RentalAgreementRentables ON (ReceiptAllocation.RAID=RentalAgreementRentables.RAID AND "{{.DtStart}}" <= RentalAgreementRentables.RARDtStop AND RentalAgreementRentables.RARDtStart < "{{.DtStop}}")
    LEFT JOIN Rentable ON (Rentable.RID=RentalAgreementRentables.RID AND Rentable.RID > 0)
    LEFT JOIN Receipt ON (RentalAgreementRentables.RAID=Receipt.RAID AND Receipt.FLAGS & 4=0 AND "{{.DtStart}}" <= Receipt.Dt AND Receipt.Dt < "{{.DtStop}}")
    LEFT JOIN RentalAgreement ON (ReceiptAllocation.RAID=RentalAgreement.RAID)
    INNER JOIN AR ON (AR.ARID = Receipt.ARID AND AR.FLAGS & 5 = 5)
    WHERE {{.WhereClause}}
    ORDER BY {{.OrderClause}};`

// RentablesNoAsmtQueryClause - query clauses for rentable NO Assessment query
var RentablesNoAsmtQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(RentablesNoAsmtFields, ","),
	"WhereClause":  "ReceiptAllocation.BID=%d AND ReceiptAllocation.RAID > 0",
	"OrderClause":  "ReceiptAllocation.Amount DESC",
	"DtStart":      "",
	"DtStop":       "",
}

// ------- No Rentables Assessments Query components -------

// NoRIDAsmtQuerySelectFields holds the list of fields need to be fetched by
// no rentable assessment query
var NoRIDAsmtQuerySelectFields = rlib.SelectQueryFields{
	"Assessments.BID",
	"Assessments.ASMID",
	"AR.Name",
	"Assessments.Amount",
	"SUM(ReceiptAllocation.Amount) as PaymentsApplied",
	"RentalAgreement.RAID",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END AS Payors",
}

// NoRIDAsmtQueryFieldMap fieldsMap for the no rentable assessment query
var NoRIDAsmtQueryFieldMap = rlib.SelectQueryFieldMap{
	"Description":     {"ARID.Name"},
	"PossessionStart": {"RentalAgreement.PossessionStart"},
	"PossessionStop":  {"RentalAgreement.PossessionStop"},
	"RentStart":       {"RentalAgreement.RentStart"},
	"RentStop":        {"RentalAgreement.RentStop"},
	"LastModTime":     {"PaymentType.LastModTime"},
	"LastModBy":       {"PaymentType.LastModBy"},
	"CreateTS":        {"PaymentType.CreateTS"},
	"CreateBy":        {"PaymentType.CreateBy"},
}

// NoRIDAsmtQuery - the query execution plan for no rentable assessment
var NoRIDAsmtQuery = `
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

// NoRIDAsmtQueryClause query clauses for no rentable assessments query
var NoRIDAsmtQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(NoRIDAsmtQuerySelectFields, ","),
	"WhereClause":  "Assessments.BID=%d AND Assessments.FLAGS&4=0 AND Assessments.RID=0 AND %q < Assessments.Stop AND Assessments.Start < %q",
	"OrderClause":  "Assessments.RAID ASC, Assessments.Start ASC",
	"DtStart":      "",
	"DtStop":       "",
}

// ------- No Rentables No Assessments Query components -------

// NoRIDNoAsmtQuerySelectFields holds the list of fields need to be fetched by
// no rentable no assessment query
var NoRIDNoAsmtQuerySelectFields = rlib.SelectQueryFields{
	"ReceiptAllocation.BID",
	"ReceiptAllocation.RAID",
	"ReceiptAllocation.Amount",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"AR.Name",
	"CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END AS Payors",
}

// NoRIDNoAsmtQueryFieldMap holds the fieldmap for no rentable no assessment query
var NoRIDNoAsmtQueryFieldMap = rlib.SelectQueryFieldMap{
	"Description":     {"ARID.Name"},
	"PossessionStart": {"RentalAgreement.PossessionStart"},
	"PossessionStop":  {"RentalAgreement.PossessionStop"},
	"RentStart":       {"RentalAgreement.RentStart"},
	"RentStop":        {"RentalAgreement.RentStop"},
	"LastModTime":     {"PaymentType.LastModTime"},
	"LastModBy":       {"PaymentType.LastModBy"},
	"CreateTS":        {"PaymentType.CreateTS"},
	"CreateBy":        {"PaymentType.CreateBy"},
}

// NoRIDNoAsmtQuery - the query execution plan for no rentables no assessments part
var NoRIDNoAsmtQuery = `
    SELECT {{.SelectClause}}
    FROM ReceiptAllocation
    LEFT JOIN RentalAgreementRentables ON RentalAgreementRentables.RAID=ReceiptAllocation.RAID
    LEFT JOIN RentalAgreement ON RentalAgreement.RAID=ReceiptAllocation.RAID
    INNER JOIN Receipt ON (Receipt.RCPTID = ReceiptAllocation.RCPTID)
    INNER JOIN AR ON (AR.ARID = Receipt.ARID AND AR.FLAGS = 5)
    INNER JOIN Transactant ON (Transactant.TCID = Receipt.TCID)
    WHERE {{.WhereClause}}
    ORDER BY {{.OrderClause}}`

// NoRIDNoAsmtQueryClause - the query clause for no rentable no assessment query
var NoRIDNoAsmtQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(NoRIDNoAsmtQuerySelectFields, ","),
	"WhereClause":  "ReceiptAllocation.BID=%d AND ReceiptAllocation.ASMID=0 AND ReceiptAllocation.RAID>0 AND %q <= ReceiptAllocation.Dt AND ReceiptAllocation.Dt < %q AND RentalAgreementRentables.RID is NULL",
	"OrderClause":  "ReceiptAllocation.Dt ASC",
	"DtStart":      "",
	"DtStop":       "",
}

// GetRRReportPartSQLRows returns the sql.Rows for the given looking part
// of the rentroll report
// If given part doesn't exist then it will return nil with error
func GetRRReportPartSQLRows(
	rrPart string,
	BID int64,
	d1, d2 time.Time,
	additionalWhere, orderBy string,
	limit, offset int,
) (*sql.Rows, error) {
	const funcname = "GetRRReportPartSQLRows"
	var (
		qry   string
		qc    rlib.QueryClause
		where string
		order string
		d1Str = d1.Format(rlib.RRDATEFMTSQL)
		d2Str = d2.Format(rlib.RRDATEFMTSQL)
	)
	rlib.Console("Entered in : %s\n", funcname)

	// based on part, decide query and queryClause
	switch rrPart {
	case "rentables":
		qry = RentablesQuery
		qc = rlib.GetQueryClauseCopy(RentablesQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID)
		break
	case "rentablesAsmt":
		qry = RentablesAsmtQuery
		qc = rlib.GetQueryClauseCopy(RentablesAsmtQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID)
		break
	case "rentablesNoAsmt":
		qry = RentablesNoAsmtQuery
		qc = rlib.GetQueryClauseCopy(RentablesNoAsmtQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID)
		break
	case "noRIDAsmt":
		qry = NoRIDAsmtQuery
		qc = rlib.GetQueryClauseCopy(NoRIDAsmtQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID, d1Str, d2Str)
		break
	case "noRIDNoAsmt":
		qry = NoRIDNoAsmtQuery
		qc = rlib.GetQueryClauseCopy(NoRIDNoAsmtQueryClause)
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
	dbQry := rlib.RenderSQLQuery(qry, qc)
	rlib.Console("db query for %s = %s\n", rrPart, dbQry)

	// return the query execution
	return rlib.RRdb.Dbrr.Query(dbQry)

	// tInit := time.Now()
	// qExec, err := rlib.RRdb.Dbrr.Query(dbQry)
	// diff := time.Since(tInit)
	// rlib.Console("\nQQQQQQuery Time diff for %s is %s\n\n", rrPart, diff.String())
	// return qExec, err
}

// RentRollReportRow represents the row that holds the data for rentroll report
// it could be used by rentroll webservice view as well as for the gotable report
type RentRollReportRow struct {
	Recid            int64            `json:"recid"` // this is to support the w2ui form
	BID              int64            // Business (so that we can process by Business)
	RID              int64            // The rentable
	RTID             int64            // The rentable type
	RAID             rlib.NullInt64   // Rental Agreement
	RARID            rlib.NullInt64   // rental agreement rentable id
	ASMID            rlib.NullInt64   // Assessment
	RentableName     rlib.NullString  // Name of the rentable
	RentableType     rlib.NullString  // Name of the rentable type
	Sqft             rlib.NullInt64   // rentable square feet
	Description      rlib.NullString  // account rule name
	RentCycle        rlib.NullInt64   // Rent Cycle
	RentCycleStr     string           // String representation of Rent Cycle
	Status           rlib.NullInt64   // Rentable status
	AgreementStart   rlib.NullDate    // start date for RA
	AgreementStop    rlib.NullDate    // stop date for RA
	AgreementPeriod  string           // text representation of Rental Agreement time period
	PossessionStart  rlib.NullDate    // start date for Occupancy
	PossessionStop   rlib.NullDate    // stop date for Occupancy
	UsePeriod        string           // text representation of Occupancy(or use) time period
	RentStart        rlib.NullDate    // start date for Rent
	RentStop         rlib.NullDate    // stop date for Rent
	RentPeriod       string           // text representation of Rent time period
	Payors           rlib.NullString  // payors list attached with this RA within same time
	Users            rlib.NullString  // users associated with the rentable
	GSR              rlib.NullFloat64 // Gross scheduled rate
	PeriodGSR        rlib.NullFloat64 // Periodic gross scheduled rate
	IncomeOffsets    rlib.NullFloat64 // Income Offset amount
	AmountDue        rlib.NullFloat64 // Amount needs to be paid by Payor(s)
	PaymentsApplied  rlib.NullFloat64 // Amount collected by Payor(s) for Assessments
	BeginningRcv     rlib.NullFloat64 // Receivable amount at beginning period
	ChangeInRcv      rlib.NullFloat64 // Change in receivable
	EndingRcv        rlib.NullFloat64 // Ending receivable
	BeginningSecDep  rlib.NullFloat64 // Beginning security deposit
	ChangeInSecDep   rlib.NullFloat64 // Change in security deposit
	EndingSecDep     rlib.NullFloat64 // Ending security deposit
	IsMainRow        bool             // is main row
	IsRentableRow    bool             // is rentable row
	IsSubTotalRow    bool             // is sustotal row
	IsBlankRow       bool             // is blank row
	IsNoRIDAsmtRow   bool             // is "No rentable assessment" row
	IsNoRIDNoAsmtRow bool             // is "No rentable No assessment" row
}

// RentRollReportRowScan scans a result from sql row and dump it in a RentRollReportRow struct
func RentRollReportRowScan(rows *sql.Rows, q *RentRollReportRow) error {
	return rows.Scan(&q.BID, &q.RID, &q.RentableName, &q.RTID, &q.RentableType,
		&q.RentCycle, &q.GSR, &q.RARID, &q.RAID,
		&q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop,
		&q.Status, &q.Payors, &q.Users)
}

// RRTextReport prints a text-based RentRoll report
// for the business in xbiz and timeframe d1 to d2 to stdout
func RRTextReport(ri *ReporterInfo) {
	fmt.Print(RRReport(ri))
}

// RRReport returns a string containin a text-based RentRoll report
// for the business in xbiz and timeframe d1 to d2.
func RRReport(ri *ReporterInfo) string {
	tbl := RRReportTable(ri)
	return ReportToString(&tbl, ri)
}

// RRReportRows returns the new rentroll report for the given date range and business id.
//
// The table rows are categorized by five types.
// 1. Rentables Row
//        Basically, it will include static and time base info.
//        If it has more than one assessment then
//        there will be separate child rows for that,
//        including only amount related info.
// 2. Rentables Row without Assessments
//        Any Payment/Receipt which are associated with rentables
//        but has no associated assessments
//        For ex. vending machine
// 3. Rentables with some special status code
//        For ex. under maintainance, vacant, etc..
// 4. All assessments which are not associated with any rentable
//        For ex. 'Application Fee' on rental agreement
// 5. All receipts which are not associated with any rentable nor with any assessment
//        For ex. Application Fees, Floating Deposits
// This routine is commonly used by both report and webservice view.
// So, for webservice view, routine needs be called with additional params
// such as limit, some offset values.
func RRReportRows(BID int64,
	startDt, stopDt time.Time,
	pageRowsLimit int,
	rentablesWC, rentablesQC string, rentablesOffset int,
	noRIDAsmtWC, noRIDAsmtQC string, noRIDAsmtOffset int,
	noRIDNoAsmtWC, noRIDNoAsmtQC string, noRIDNoAsmtOffset int,
) ([]RentRollReportRow, error) {

	const funcname = "RRReportRows"
	var (
		err                  error
		customAttrRTSqft     = "Square Feet"                      // custom attribute for all rentables
		grandTTL             = RentRollReportRow{IsMainRow: true} // grand total row
		xbiz                 rlib.XBusiness
		noRIDAsmtRowsLimit   = 0 // limit on "NO rentable assessment" rows
		noRIDNoAsmtRowsLimit = 0 // limit on "NO rentable NO assessment" rows
		rptMainRowsCount     = 0 // report main rows count
	)
	rlib.Console("Entered in %s\n", funcname)

	// init some structure
	reportRows := []RentRollReportRow{}
	rlib.InitBizInternals(BID, &xbiz) // init some business internals first

	//==========================
	// RENTABLES QUERY EXECUTION
	//==========================
	// if there is no limit then it is meaningless having a value for below variables
	if pageRowsLimit <= 0 {
		rentablesOffset = -1
		pageRowsLimit = -1
	}

	rentablesRows, err := GetRRReportPartSQLRows("rentables", BID,
		startDt, stopDt,
		rentablesWC, rentablesQC,
		pageRowsLimit, rentablesOffset)

	if err != nil {
		return reportRows, err
	}
	defer rentablesRows.Close()

	// ===========================
	// LOOP THROUGH RENTABLES ROWS
	// ===========================
	rentablesCount := 0
	for rentablesRows.Next() {
		q := RentRollReportRow{IsMainRow: true, IsRentableRow: true}
		if err = RentRollReportRowScan(rentablesRows, &q); err != nil {
			return reportRows, err
		}
		if len(xbiz.RT[q.RTID].CA) > 0 { // if there are custom attributes
			c, ok := xbiz.RT[q.RTID].CA[customAttrRTSqft] // see if Square Feet is among them
			if ok {                                       // if it is...
				sqft, err := rlib.IntFromString(c.Value, "invalid customAttrRTSqft attribute")
				q.Sqft.Scan(sqft)
				if err != nil {
					return reportRows, err
				}
			}
		}
		if q.RentStart.Time.Year() > 1970 {
			q.RentPeriod = fmt.Sprintf("%s\n - %s", q.RentStart.Time.Format(rlib.RRDATEFMT3), q.RentStop.Time.Format(rlib.RRDATEFMT3))
		}
		if q.PossessionStart.Time.Year() > 1970 {
			q.UsePeriod = fmtRRDatePeriod(&q.PossessionStart.Time, &q.PossessionStop.Time)
		}
		for freqStr, freqNo := range rlib.CycleFreqMap {
			if q.RentCycle.Int64 == freqNo {
				q.RentCycleStr = freqStr
			}
		}

		//------------------------------------------------------------
		// There may be multiple rows for the ASSESSMENTS query and
		// the NO-ASSESSMENTS query. Hold each row RentRollReportRow in slice
		// Also, compute subtotals as we go
		//------------------------------------------------------------
		d70 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
		subList := []RentRollReportRow{}
		subTotalRow := RentRollReportRow{IsSubTotalRow: true}
		subTotalRow.AmountDue.Valid = true
		subTotalRow.PaymentsApplied.Valid = true
		subTotalRow.PeriodGSR.Valid = true
		subTotalRow.IncomeOffsets.Valid = true

		//========================
		//  ASSESSMENTS QUERY...
		//========================
		// here we have to apply different whereClause
		// for the rentables Assessment Query as we're looking
		// for ALL assessments for specific rentable
		rentablesAsmtAdditionalWhere := fmt.Sprintf("Rentable.RID=%d", q.RID)
		rentablesAsmtRows, err := GetRRReportPartSQLRows("rentablesAsmt", q.BID,
			startDt, stopDt,
			rentablesAsmtAdditionalWhere, "",
			-1, -1) // we need to fetch all asmt

		if err != nil {
			return reportRows, err
		}
		defer rentablesAsmtRows.Close()

		//============================================================
		//   LOOP THROUGH ASSESSMENTS AND RECEIPTS FOR THIS RENTABLE
		//============================================================
		rentablesChildRowsCount := 0
		for rentablesAsmtRows.Next() {
			var nq = RentRollReportRow{RID: q.RID, BID: q.BID}
			if rentablesChildRowsCount == 0 {
				nq = q
			}
			if err = rentablesAsmtRows.Scan(&nq.Description, &nq.RAID,
				&nq.PossessionStart, &nq.PossessionStop, &nq.RentStart, &nq.RentStop,
				&nq.AmountDue, &nq.PaymentsApplied,
			); err != nil {
				return reportRows, err
			}
			setRRDatePeriodString(&subList, &nq) // adds dates as needed
			if nq.RAID.Valid || nq.Description.Valid || nq.AmountDue.Valid || nq.PaymentsApplied.Valid {
				addToSubList(&subList, &rentablesChildRowsCount, &nq)
				updateSubTotals(&subTotalRow, &nq)
			}
		}

		//============================
		//  NO-ASSESSMENTS QUERY...
		//============================
		// we need to change whereClause for the rentables no Assessment query
		// as we're looking for ALL payments associated with specific rentable
		// but has no any assessments
		rentablesNoAsmtAdditionalWhere := fmt.Sprintf("RentalAgreementRentables.RID=%d", q.RID)
		rentablesNoAsmtRows, err := GetRRReportPartSQLRows("rentablesNoAsmt", q.BID,
			startDt, stopDt,
			rentablesNoAsmtAdditionalWhere, "",
			-1, -1) // need to fetch all receipts

		if err != nil {
			return reportRows, err
		}
		defer rentablesNoAsmtRows.Close()

		//====================================================
		//   LOOP THROUGH NO-ASSESSMENTS FOR THIS RENTABLE
		//====================================================
		for rentablesNoAsmtRows.Next() {
			var nq = RentRollReportRow{RID: q.RID, BID: q.BID}
			if rentablesChildRowsCount == 0 {
				nq = q
			}
			err = rentablesNoAsmtRows.Scan(&nq.Description, &nq.RAID,
				&nq.PossessionStart, &nq.PossessionStop, &nq.RentStart, &nq.RentStop,
				&nq.PaymentsApplied)

			if err != nil {
				return reportRows, err
			}
			setRRDatePeriodString(&subList, &nq) // adds dates as needed
			if nq.Description.Valid || nq.RAID.Valid || nq.PaymentsApplied.Valid {
				addToSubList(&subList, &rentablesChildRowsCount, &nq)
				updateSubTotals(&subTotalRow, &nq)
			}
		}

		//----------------------------------------------------------------------
		// Handle the case where both the Assesments and No-Assessment lists
		// had no matches... just add what we know...
		//----------------------------------------------------------------------
		if len(subList) == 0 {
			addToSubList(&subList, &rentablesChildRowsCount, &q)
		} else {
			//====================================================
			//   CHECK FOR GAPS IN COVERAGE
			//====================================================
			handleRentableGaps(&subList, &startDt, &stopDt)
		}

		// NOW, ADD append all sublist row in main data struture
		// ------------------------------------------------------
		reportRows = append(reportRows, subList...)

		//----------------------------------------
		// Add the Rentable receivables totals...
		//----------------------------------------
		subTotalRow.Description.String = "Subtotal"
		subTotalRow.Description.Valid = true
		subTotalRow.BeginningRcv.Float64, subTotalRow.EndingRcv.Float64, err = rlib.GetBeginEndRARBalance(q.BID, q.RID, q.RAID.Int64, &startDt, &stopDt)
		subTotalRow.BeginningRcv.Valid = true
		subTotalRow.ChangeInRcv.Float64 = subTotalRow.EndingRcv.Float64 - subTotalRow.BeginningRcv.Float64
		subTotalRow.ChangeInRcv.Valid = true
		subTotalRow.EndingRcv.Valid = true

		//----------------------------------------
		// Add the Security Deposit totals...
		//----------------------------------------
		subTotalRow.BeginningSecDep.Float64, err = rlib.GetSecDepBalance(q.BID, q.RAID.Int64, q.RID, &d70, &startDt)
		if err != nil {
			return reportRows, err
		}
		subTotalRow.BeginningSecDep.Valid = true
		subTotalRow.ChangeInSecDep.Float64, err = rlib.GetSecDepBalance(q.BID, q.RAID.Int64, q.RID, &startDt, &stopDt)
		if err != nil {
			return reportRows, err
		}
		subTotalRow.ChangeInSecDep.Valid = true
		subTotalRow.EndingSecDep.Float64 = subTotalRow.BeginningSecDep.Float64 + subTotalRow.ChangeInSecDep.Float64
		subTotalRow.EndingSecDep.Valid = true

		// NOW ADD SUB TOTAL ROW IN LIST
		reportRows = append(reportRows, subTotalRow)
		rentablesChildRowsCount++

		// add subTotal amounts to grand total record
		updateGrandTotals(&grandTTL, &subTotalRow)

		// ALSO, ADD BLANK ROW
		reportRows = append(reportRows, RentRollReportRow{IsBlankRow: true})
		rentablesChildRowsCount++

		// update the rentablesCount only after adding the record
		rentablesCount++
	}

	err = rentablesRows.Err()
	if err != nil {
		return reportRows, err
	}
	rlib.Console("Added %d Rentable rows\n", rentablesCount)
	rptMainRowsCount += rentablesCount // how many total rows have been added to list

	// if for given limit, rows are feed within page then return
	if isReportComplete(pageRowsLimit, rptMainRowsCount) {
		return reportRows, err
	}

	//====================================
	// NO-RENTABLE ASSESSMENTS QUERY...
	//====================================

	// if no limit then reset the values
	if pageRowsLimit <= 0 {
		noRIDAsmtWC = ""
		noRIDAsmtQC = ""
		noRIDAsmtRowsLimit = -1
		noRIDAsmtOffset = -1
	} else {
		noRIDAsmtRowsLimit = pageRowsLimit - len(reportRows)
		if noRIDAsmtRowsLimit < 0 {
			noRIDAsmtRowsLimit = 0 // make sure it doesn't have minus value
		}
	}

	noRIDAsmtRows, err := GetRRReportPartSQLRows("noRIDAsmt", BID,
		startDt, stopDt,
		noRIDAsmtWC, noRIDAsmtQC,
		noRIDAsmtRowsLimit, noRIDAsmtOffset)

	if err != nil {
		return reportRows, err
	}
	defer noRIDAsmtRows.Close()

	// ==============================
	// LOOP THROUGH NO RID ASMT ROWS
	// ==============================
	noRIDAsmtRowsCount := 0
	for noRIDAsmtRows.Next() {
		q := RentRollReportRow{IsMainRow: true, IsNoRIDAsmtRow: true}
		err = noRIDAsmtRows.Scan(&q.BID, &q.ASMID, &q.Description,
			&q.AmountDue, &q.PaymentsApplied, &q.RAID,
			&q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop, &q.Payors)

		if err != nil {
			return reportRows, err
		}
		setRRDatePeriodString(&reportRows, &q)

		// APPEND NO-RID-ASMT ROW IN LIST
		reportRows = append(reportRows, q)
		noRIDAsmtRowsCount++

		// add subTotal amounts to grand total record
		updateGrandTotals(&grandTTL, &q)
	}
	rlib.Console("Added noRID Asmt rows: %d", noRIDAsmtRowsCount)
	rptMainRowsCount += noRIDAsmtRowsCount // how many total rows have been added to list

	// if for given limit, rows are feed within page then return
	if isReportComplete(pageRowsLimit, rptMainRowsCount) {
		return reportRows, err
	}

	//=======================================
	//  NO Rentables No ASSESSMENTS QUERY...
	//=======================================

	// if no limit then reset the values
	if pageRowsLimit <= 0 {
		noRIDNoAsmtWC = ""
		noRIDNoAsmtQC = ""
		noRIDNoAsmtRowsLimit = -1
		noRIDNoAsmtOffset = -1
	} else {
		noRIDNoAsmtRowsLimit = pageRowsLimit - len(reportRows)
		if noRIDNoAsmtRowsLimit < 0 {
			noRIDNoAsmtRowsLimit = 0 // make sure it doesn't have minus value
		}
	}

	noRIDNoAsmtRows, err := GetRRReportPartSQLRows("noRIDNoAsmt", BID,
		startDt, stopDt,
		noRIDNoAsmtWC, noRIDNoAsmtQC,
		noRIDNoAsmtRowsLimit, noRIDNoAsmtOffset)

	if err != nil {
		return reportRows, err
	}
	defer noRIDNoAsmtRows.Close()

	// =================================
	// LOOP THROUGH NO RID NO ASMT ROWS
	// =================================
	noRIDNoAsmtRowsCount := 0
	for noRIDNoAsmtRows.Next() {
		q := RentRollReportRow{IsMainRow: true, IsNoRIDNoAsmtRow: true}
		err = noRIDNoAsmtRows.Scan(&q.BID, &q.RAID, &q.PaymentsApplied,
			&q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop,
			&q.Description, &q.Payors)

		if err != nil {
			return reportRows, err
		}
		setRRDatePeriodString(&reportRows, &q)

		// APPEND NO-RID-NO-ASMT ROW IN LIST
		reportRows = append(reportRows, q)
		noRIDNoAsmtRowsCount++

		// add subTotal amounts to grand total record
		updateGrandTotals(&grandTTL, &q)
	}
	rlib.Console("Added noRID NoAsmt rows: %d", noRIDNoAsmtRowsCount)
	rptMainRowsCount += noRIDNoAsmtRowsCount // how many total rows have been added to list

	// ================
	// GRAND TOTAL ROW
	// ================
	grandTTL.Description.Scan("Grand Total")
	reportRows = append(reportRows, grandTTL)

	return reportRows, err
}

// isReportComplete checks whether page result rows is filled completely with given limit.
// only applicable for virtual scrolling.
func isReportComplete(pageRowsLimit int, mainRowsCount int) bool {
	if pageRowsLimit > 0 {
		if mainRowsCount >= pageRowsLimit {
			return true
		}
		return false
	}
	return false
}

// RRReportTable returns the gotable representation for rentroll report
func RRReportTable(ri *ReporterInfo) gotable.Table {
	const funcname = "RRReportTable"
	var (
		err error
		tbl = getRRTable() // gotable init for this report
	)
	rlib.Console("Entered in %s", funcname)

	// use section3 for errors and apply red color
	cssListSection3 := []*gotable.CSSProperty{
		{Name: "color", Value: "red"},
		{Name: "font-family", Value: "monospace"},
	}
	tbl.SetSection3CSS(cssListSection3)

	// set table title, sections
	err = TableReportHeaderBlock(&tbl, "Rentroll", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// Add columns to the table
	tbl.AddColumn("Rentable", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                    // column for the Rentable name
	tbl.AddColumn("Rentable Type", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)               // RentableType name
	tbl.AddColumn("SqFt", 5, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)                        // the Custom Attribute "Square Feet"
	tbl.AddColumn("Description", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                 // the Custom Attribute "Square Feet"
	tbl.AddColumn("Users", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                       // Users of this rentable
	tbl.AddColumn("Payors", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                      // Users of this rentable
	tbl.AddColumn("Rental Agreement", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)            // the Rental Agreement id
	tbl.AddColumn("Use Period", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                  // the use period
	tbl.AddColumn("Rent Period", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                 // the rent period
	tbl.AddColumn("Rent Cycle", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                  // the rent cycle
	tbl.AddColumn("GSR Rate", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)                   // gross scheduled rent
	tbl.AddColumn("Period GSR", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)                 // gross scheduled rent
	tbl.AddColumn("Income Offsets", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)             // GL Account
	tbl.AddColumn("Payments Applied", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)           // contract rent amounts
	tbl.AddColumn("Beginning Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Change In Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Ending Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)          // account for the associated RentalAgreement
	tbl.AddColumn("Beginning Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Change In Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Ending Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)    // account for the associated RentalAgreement

	// NOW GET THE ROWS FOR RENTROLL ROUTINE
	rows, err := RRReportRows(
		ri.Bid, ri.D1, ri.D2, // BID, startDate, stopDate
		-1,         // limit
		"", "", -1, // rentables Part
		"", "", -1, // "No Rentable Assessment" part
		"", "", -1) // "No Rentable No Assessment" part

	// if any error encountered then just set it to section3
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}

	for index, row := range rows {
		if row.IsSubTotalRow { // add line before subtotal Row
			// tbl.AddLineBefore(index) // AddLineBefore is not working
			tbl.AddLineAfter(index - 1)
		}
		rrTableAddRow(&tbl, row)
	}
	tbl.AddLineAfter(len(tbl.Row) - 2) // Grand Total line, Rows index start from zero

	return tbl
}

// addToSubList is a convenience function that adds a new RentRollReportRow struct to the
// supplied grid struct and updates the
//
// INPUTS
//           g = pointer to a slice of RentRollReportRow structs to which p will be added
//  childCount = pointer to a counter to increment when a record is added
//-----------------------------------------------------------------------------
func addToSubList(g *[]RentRollReportRow, childCount *int, p *RentRollReportRow) {
	(*childCount)++
	*g = append(*g, *p)
}

// updateSubTotals does all subtotal calculations for the subtotal line
//-----------------------------------------------------------------------------
func updateSubTotals(sub, q *RentRollReportRow) {
	sub.AmountDue.Float64 += q.AmountDue.Float64
	sub.PaymentsApplied.Float64 += q.PaymentsApplied.Float64
	sub.PeriodGSR.Float64 += q.PeriodGSR.Float64
	sub.IncomeOffsets.Float64 += q.IncomeOffsets.Float64
	// rlib.Console("\t q.Description = %s, q.AmountDue = %.2f, q.PaymentsApplied = %.2f\n", q.Description, q.AmountDue.Float64, q.PaymentsApplied.Float64)
	// rlib.Console("\t sub.AmountDue = %.2f, sub.PaymentsApplied = %.2f\n", sub.AmountDue.Float64, sub.PaymentsApplied.Float64)
}

// updateGrandTotals does grand total from subTotal Rows
//-----------------------------------------------------------------------------
func updateGrandTotals(grandTotal, subTotal *RentRollReportRow) {
	grandTotal.AmountDue.Float64 += subTotal.AmountDue.Float64
	grandTotal.PaymentsApplied.Float64 += subTotal.PaymentsApplied.Float64
	grandTotal.PeriodGSR.Float64 += subTotal.PeriodGSR.Float64
	grandTotal.IncomeOffsets.Float64 += subTotal.IncomeOffsets.Float64
	// rlib.Console("\t subTotal.Description = %s, subTotal.AmountDue = %.2f, subTotal.PaymentsApplied = %.2f\n", subTotal.Description, subTotal.AmountDue.Float64, subTotal.PaymentsApplied.Float64)
	// rlib.Console("\t grandTotal.AmountDue = %.2f, grandTotal.PaymentsApplied = %.2f\n", grandTotal.AmountDue.Float64, grandTotal.PaymentsApplied.Float64)
}

// int64ToStr returns the string represenation of int64 type number
// if blank is set to true, then it will returns blank string otherwise returns 0
func int64ToStr(number int64, blank bool) string {
	nStr := strconv.FormatInt(number, 10)
	if nStr == "0" {
		if blank {
			return ""
		}
	}
	return nStr
}

// float64ToStr returns the string represenation of float64 type number
// if blank is set to true, then it will returns blank string otherwise returns 0.00
func float64ToStr(number float64, blank bool) string {
	nStr := strconv.FormatFloat(number, 'f', 2, 64)
	if nStr == "0.00" {
		if blank {
			return ""
		}
	}
	return nStr
}

// rrTableAddRow adds row in gotable struct with information
// given by RentRollReportRow struct
func rrTableAddRow(tbl *gotable.Table, q RentRollReportRow) {

	// column numbers for gotable report
	const (
		RName       = 0
		RType       = iota
		SqFt        = iota
		Descr       = iota
		Users       = iota
		Payors      = iota
		RAgr        = iota
		UsePeriod   = iota
		RentPeriod  = iota
		RAgrStart   = iota
		RAgrStop    = iota
		RentCycle   = iota
		GSRRate     = iota
		GSRAmt      = iota
		IncOff      = iota
		AmtDue      = iota
		PmtRcvd     = iota
		BeginRcv    = iota
		ChgRcv      = iota
		EndRcv      = iota
		BeginSecDep = iota
		ChgSecDep   = iota
		EndSecDep   = iota
	)

	tbl.AddRow()
	tbl.Puts(-1, RName, q.RentableName.String)
	tbl.Puts(-1, RType, q.RentableType.String)
	tbl.Puts(-1, SqFt, int64ToStr(q.Sqft.Int64, true))
	tbl.Puts(-1, Descr, q.Description.String)
	tbl.Puts(-1, Users, q.Users.String)
	tbl.Puts(-1, Payors, q.Payors.String)
	raidStr := int64ToStr(q.RAID.Int64, true)
	raStr := ""
	if len(raidStr) > 0 {
		raStr = "RA-" + raidStr
	}
	tbl.Puts(-1, RAgr, raStr)
	tbl.Puts(-1, UsePeriod, q.UsePeriod)
	tbl.Puts(-1, RentPeriod, q.RentPeriod)
	tbl.Puts(-1, RentCycle, q.RentCycleStr)
	tbl.Puts(-1, GSRRate, float64ToStr(q.GSR.Float64, false))
	tbl.Puts(-1, GSRAmt, float64ToStr(q.PeriodGSR.Float64, false))
	tbl.Puts(-1, IncOff, float64ToStr(q.IncomeOffsets.Float64, false))
	tbl.Puts(-1, AmtDue, float64ToStr(q.AmountDue.Float64, false))
	tbl.Puts(-1, PmtRcvd, float64ToStr(q.PaymentsApplied.Float64, false))
	tbl.Puts(-1, BeginRcv, float64ToStr(q.BeginningRcv.Float64, false))
	tbl.Puts(-1, ChgRcv, float64ToStr(q.ChangeInRcv.Float64, false))
	tbl.Puts(-1, EndRcv, float64ToStr(q.EndingRcv.Float64, false))
	tbl.Puts(-1, BeginSecDep, float64ToStr(q.BeginningSecDep.Float64, false))
	tbl.Puts(-1, ChgSecDep, float64ToStr(q.ChangeInSecDep.Float64, false))
	tbl.Puts(-1, EndSecDep, float64ToStr(q.EndingSecDep.Float64, false))
}

// handleRentableGaps identifies periods during which the Rentable is not
// covered by a RentalAgreement. It updates the list with entries
// describing the gaps
//----------------------------------------------------------------------
func handleRentableGaps(sl *[]RentRollReportRow, d1, d2 *time.Time) {
	var a = []rlib.Period{}
	for i := 0; i < len(*sl); i++ {
		var p = rlib.Period{
			D1: (*sl)[i].PossessionStart.Time,
			D2: (*sl)[i].PossessionStop.Time,
		}
		a = append(a, p)
	}
	b := rlib.FindGaps(d1, d2, a)
	for i := 0; i < len(b); i++ {
		var r RentRollReportRow
		r.PossessionStart.Scan(b[i].D1)
		r.PossessionStop.Scan(b[i].D2)
		r.Description.Scan("Vacancy")
		r.UsePeriod = fmtRRDatePeriod(&b[i].D1, &b[i].D2)
		(*sl) = append((*sl), r)
	}
}

// fmtRRDatePeriod formats a start and end time as needed byt the
// column headers in the RentRoll view/report
//
// INPUT
// d1 - start of period
// d2 - end of period
//
// RETURN
// string with formated dates
//----------------------------------------------------------------------
func fmtRRDatePeriod(d1, d2 *time.Time) string {
	if d1.Year() > 1970 && d2.Year() > 1970 {
		return d1.Format(rlib.RRDATEFMT3) + "<br> - " + d2.Format(rlib.RRDATEFMT3)
	}
	return ""
}

// setRRDatePeriodString updates the nq UsePeriod and RentPeriod members
// if it is either the first row in subList or if the RentalAgreement has
// changed since the last entry in subList.
//
// INPUT
// sublist = the slice of RentRollReportRow structs
// nq = the current entry but not yet added to sublist
//
// RETURN
// void
//----------------------------------------------------------------------
func setRRDatePeriodString(rows *[]RentRollReportRow, nq *RentRollReportRow) {
	showDates := true // only list dates if the rental agreement changed
	if len(*rows) > 0 {
		showDates = (*rows)[len(*rows)-1].RAID.Int64 != nq.RAID.Int64
	}
	if showDates {
		nq.UsePeriod = fmtRRDatePeriod(&nq.PossessionStart.Time, &nq.PossessionStop.Time)
		nq.RentPeriod = fmtRRDatePeriod(&nq.RentStart.Time, &nq.RentStop.Time)
	}
}
