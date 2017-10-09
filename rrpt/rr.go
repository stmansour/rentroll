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

// NoRIDNoAsmQueryFieldMap holds the fieldmap for no rentable no assessment query
var NoRIDNoAsmQueryFieldMap = rlib.SelectQueryFieldMap{
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
}

type rentableRow struct {
	BID             int64           // Business (so that we can process by Business)
	RID             int64           // The rentable
	RTID            int64           // The rentable type
	RARID           rlib.NullInt64  // rental agreement rentable id
	RentableName    rlib.NullString // Name of the rentable
	RentableType    rlib.NullString // Name of the rentable type
	RentCycle       rlib.NullInt64  // Rent Cycle
	RentCycleStr    string          //String representation of Rent Cycle
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
}

// rentableRowScan scans a result from sql row and dump it in a rentableRow struct
func rentableRowScan(rows *sql.Rows, q *rentableRow) error {
	return rows.Scan(&q.BID, &q.RID, &q.RentableName, &q.RTID, &q.RentableType, &q.RentCycle, &q.GSR, &q.RARID, &q.RAID,
		&q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop, &q.Status, &q.Payors, &q.Users)
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

// RRReportTable returns the new rentroll report for the given date range.
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
func RRReportTable(ri *ReporterInfo) gotable.Table {
	const funcname = "RRReportTable"
	var (
		err     error
		startDt = ri.D1
		stopDt  = ri.D2
		tbl     = getRRTable() // gotable init for this report
		// totalErrs        = 0
		customAttrRTSqft = "Square Feet"
		grandTTL         = rentableRow{}
	)
	rlib.Console("Entered in %s", funcname)

	// use section3 for errors and apply red color
	cssListSection3 := []*gotable.CSSProperty{
		{Name: "color", Value: "red"},
		{Name: "font-family", Value: "monospace"},
	}
	tbl.SetSection3CSS(cssListSection3)

	// init some values
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true
	ri.BlankLineAfterRptName = true
	grandTTL.Description.Scan("Grand Total")

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

	//==========================
	// RENTABLES QUERY EXECUTION
	//==========================
	rentablesRows, err := GetRRReportPartSQLRows("rentables", ri.Bid,
		startDt, stopDt,
		"", "", -1, -1)

	if err != nil {
		if rlib.IsSQLNoResultsError(err) {
			tbl.SetSection3(NoRecordsFoundMsg)
		} else {
			tbl.SetSection3(err.Error())
		}
		return tbl
	}
	defer rentablesRows.Close()

	// ===========================
	// LOOP THROUGH RENTABLES ROWS
	// ===========================
	count := 0
	for rentablesRows.Next() {
		q := rentableRow{}
		if err = rentableRowScan(rentablesRows, &q); err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		if len(ri.Xbiz.RT[q.RTID].CA) > 0 { // if there are custom attributes
			c, ok := ri.Xbiz.RT[q.RTID].CA[customAttrRTSqft] // see if Square Feet is among them
			if ok {                                          // if it is...
				sqft, err := rlib.IntFromString(c.Value, "invalid customAttrRTSqft attribute")
				q.Sqft.Scan(sqft)
				if err != nil {
					tbl.SetSection3(err.Error())
					return tbl
				}
			}
		}
		if q.RentStart.Time.Year() > 1970 {
			q.RentPeriod = fmt.Sprintf("%s\n - %s", q.RentStart.Time.Format(rlib.RRDATEFMT3), q.RentStop.Time.Format(rlib.RRDATEFMT3))
		}
		if q.PossessionStart.Time.Year() > 1970 {
			q.UsePeriod = FmtRRDatePeriod(&q.PossessionStart.Time, &q.PossessionStop.Time)
		}
		for freqStr, freqNo := range rlib.CycleFreqMap {
			if q.RentCycle.Int64 == freqNo {
				q.RentCycleStr = freqStr
			}
		}

		//------------------------------------------------------------
		// There may be multiple rows for the ASSESSMENTS query and
		// the NO-ASSESSMENTS query. Hold each row rentableRow in slice
		// Also, compute subtotals as we go
		//------------------------------------------------------------
		d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
		subList := []rentableRow{}
		sub := rentableRow{}
		sub.AmountDue.Valid = true
		sub.PaymentsApplied.Valid = true
		sub.PeriodGSR.Valid = true
		sub.IncomeOffsets.Valid = true

		//========================
		//  ASSESSMENTS QUERY...
		//========================
		// here we have to apply different whereClause
		// for the rentables Assessment Query as we're looking
		// for ALL assessments for specific rentable
		rentablesAsmtAdditionalWhere := fmt.Sprintf("Rentable.RID=%d", q.RID)
		rentablesAsmtRows, err := GetRRReportPartSQLRows("rentablesAsmt", ri.Bid,
			startDt, stopDt,
			rentablesAsmtAdditionalWhere, "", -1, -1)

		if err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		defer rentablesAsmtRows.Close()

		//============================================================
		//   LOOP THROUGH ASSESSMENTS AND RECEIPTS FOR THIS RENTABLE
		//============================================================
		childCount := 0
		for rentablesAsmtRows.Next() {
			var nq = rentableRow{RID: q.RID, BID: q.BID}
			if childCount == 0 {
				nq = q
			}
			err = rentablesAsmtRows.Scan(&nq.Description, &nq.RAID, &nq.PossessionStart, &nq.PossessionStop, &nq.RentStart, &nq.RentStop, &nq.AmountDue, &nq.PaymentsApplied)
			if err != nil {
				tbl.SetSection3(err.Error())
				return tbl
			}
			setRRDatePeriodString(&tbl, &nq) // adds dates as needed
			if nq.RAID.Valid || nq.Description.Valid || nq.AmountDue.Valid || nq.PaymentsApplied.Valid {
				addToSubList(&subList, &childCount, &nq)
				updateSubTotals(&sub, &nq)
			}
		}

		//============================
		//  NO-ASSESSMENTS QUERY...
		//============================
		// we need to change whereClause for the rentables no Assessment query
		// as we're looking for ALL payments associated with specific rentable
		// but has no any assessments
		rentablesNoAsmtAdditionalWhere := fmt.Sprintf("RentalAgreementRentables.RID=%d", q.RID)
		rentablesNoAsmtRows, err := GetRRReportPartSQLRows("rentablesNoAsmt", ri.Bid,
			startDt, stopDt,
			rentablesNoAsmtAdditionalWhere, "",
			-1, -1)

		if err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		defer rentablesNoAsmtRows.Close()

		//====================================================
		//   LOOP THROUGH NO-ASSESSMENTS FOR THIS RENTABLE
		//====================================================
		for rentablesNoAsmtRows.Next() {
			var nq = rentableRow{RID: q.RID, BID: q.BID}
			if childCount == 0 {
				nq = q
			}
			err = rentablesNoAsmtRows.Scan(&nq.Description, &nq.RAID, &nq.PossessionStart, &nq.PossessionStop, &nq.RentStart, &nq.RentStop, &nq.PaymentsApplied)
			if err != nil {
				tbl.SetSection3(err.Error())
				return tbl
			}
			setRRDatePeriodString(&tbl, &nq) // adds dates as needed
			if nq.Description.Valid || nq.RAID.Valid || nq.PaymentsApplied.Valid {
				addToSubList(&subList, &childCount, &nq)
				updateSubTotals(&sub, &nq)
			}
		}

		//----------------------------------------------------------------------
		// Handle the case where both the Assesments and No-Assessment lists
		// had no matches... just add what we know...
		//----------------------------------------------------------------------
		if len(subList) == 0 {
			addToSubList(&subList, &childCount, &q)
		} else {
			//====================================================
			//   CHECK FOR GAPS IN COVERAGE
			//====================================================
			handleGaps(&subList, &ri.D1, &ri.D2)

		}

		// now add all child rows respected with rentableRow in gotable
		for _, row := range subList {
			rrTableAddRow(&tbl, row)
		}

		//----------------------------------------
		// Add the Rentable receivables totals...
		//----------------------------------------
		sub.Description.String = "Subtotal"
		sub.Description.Valid = true
		sub.BeginningRcv.Float64, sub.EndingRcv.Float64, err = rlib.GetBeginEndRARBalance(ri.Bid, q.RID, q.RAID.Int64, &startDt, &stopDt)
		sub.BeginningRcv.Valid = true
		sub.ChangeInRcv.Float64 = sub.EndingRcv.Float64 - sub.BeginningRcv.Float64
		sub.ChangeInRcv.Valid = true
		sub.EndingRcv.Valid = true

		//----------------------------------------
		// Add the Security Deposit totals...
		//----------------------------------------
		sub.BeginningSecDep.Float64, err = rlib.GetSecDepBalance(q.BID, q.RAID.Int64, q.RID, &d1, &startDt)
		if err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		sub.BeginningSecDep.Valid = true
		sub.ChangeInSecDep.Float64, err = rlib.GetSecDepBalance(q.BID, q.RAID.Int64, q.RID, &startDt, &stopDt)
		if err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		sub.ChangeInSecDep.Valid = true
		sub.EndingSecDep.Float64 = sub.BeginningSecDep.Float64 + sub.ChangeInSecDep.Float64
		sub.EndingSecDep.Valid = true

		// =====================
		// SUBTOTAL LINE
		// =====================
		tbl.AddLineAfter(len(tbl.Row) - 1)
		rrTableAddRow(&tbl, sub)
		childCount++
		tbl.AddLineAfter(len(tbl.Row) - 1) // SHOULD WE ADD LINE AFTER SUBTOTAL ROW??

		// add subTotal amounts to grand total record
		updateGrandTotals(&grandTTL, &sub)

		// =====================
		// BLANK LINE
		// =====================
		rrTableAddRow(&tbl, rentableRow{})
		childCount++

		// update the count only after adding the record
		count++
	}

	err = rentablesRows.Err()
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}
	rlib.Console("Added %d Rentable rows\n", count)

	//====================================
	//  NO Rentables ASSESSMENTS QUERY...
	//====================================
	noRIDAsmtRows, err := GetRRReportPartSQLRows("noRIDAsmt", ri.Bid,
		startDt, stopDt,
		"", "", -1, -1)

	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer noRIDAsmtRows.Close()

	// ==============================
	// LOOP THROUGH NO RID ASMT ROWS
	// ==============================
	count = 0
	for noRIDAsmtRows.Next() {
		q := rentableRow{}
		err := noRIDAsmtRows.Scan(&q.BID, &q.ASMID, &q.Description, &q.AmountDue, &q.PaymentsApplied, &q.RAID, &q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop, &q.Payors)
		if err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		setRRDatePeriodString(&tbl, &q)
		rrTableAddRow(&tbl, q)
		count++

		// add subTotal amounts to grand total record
		updateGrandTotals(&grandTTL, &q)
	}
	rlib.Console("Added noRID Asmt rows: %d", count)

	//=======================================
	//  NO Rentables No ASSESSMENTS QUERY...
	//=======================================
	noRIDNoAsmtRows, err := GetRRReportPartSQLRows("noRIDNoAsmt", ri.Bid,
		startDt, stopDt,
		"", "", -1, -1)

	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}
	defer noRIDNoAsmtRows.Close()

	// =================================
	// LOOP THROUGH NO RID NO ASMT ROWS
	// =================================
	count = 0
	for noRIDNoAsmtRows.Next() {
		q := rentableRow{}
		err := noRIDNoAsmtRows.Scan(&q.BID, &q.RAID, &q.PaymentsApplied, &q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop, &q.Description, &q.Payors)
		if err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		setRRDatePeriodString(&tbl, &q)
		rrTableAddRow(&tbl, q)
		count++

		// add subTotal amounts to grand total record
		updateGrandTotals(&grandTTL, &q)
	}

	// at last add grand total row to the table
	tbl.AddLineAfter(len(tbl.Row) - 1)
	rrTableAddRow(&tbl, grandTTL)

	return tbl
}

// addToSubList is a convenience function that adds a new rentableRow struct to the
// supplied grid struct and updates the
//
// INPUTS
//           g = pointer to a slice of rentableRow structs to which p will be added
//  childCount = pointer to a counter to increment when a record is added
//-----------------------------------------------------------------------------
func addToSubList(g *[]rentableRow, childCount *int, p *rentableRow) {
	(*childCount)++
	*g = append(*g, *p)
}

// updateSubTotals does all subtotal calculations for the subtotal line
//-----------------------------------------------------------------------------
func updateSubTotals(sub, q *rentableRow) {
	sub.AmountDue.Float64 += q.AmountDue.Float64
	sub.PaymentsApplied.Float64 += q.PaymentsApplied.Float64
	sub.PeriodGSR.Float64 += q.PeriodGSR.Float64
	sub.IncomeOffsets.Float64 += q.IncomeOffsets.Float64
	// rlib.Console("\t q.Description = %s, q.AmountDue = %.2f, q.PaymentsApplied = %.2f\n", q.Description, q.AmountDue.Float64, q.PaymentsApplied.Float64)
	// rlib.Console("\t sub.AmountDue = %.2f, sub.PaymentsApplied = %.2f\n", sub.AmountDue.Float64, sub.PaymentsApplied.Float64)
}

// updateGrandTotals does grand total from subTotal Rows
//-----------------------------------------------------------------------------
func updateGrandTotals(grandTotal, subTotal *rentableRow) {
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
	if number > 0 {
		return strconv.FormatInt(number, 10)
	}
	if blank {
		return ""
	}
	return strconv.FormatInt(number, 10)
}

// float64ToStr returns the string represenation of float64 type number
// if blank is set to true, then it will returns blank string otherwise returns 0.00
func float64ToStr(number float64, blank bool) string {
	if number > 0 {
		return strconv.FormatFloat(number, 'f', 2, 64)
	}
	if blank {
		return ""
	}
	return strconv.FormatFloat(number, 'f', 2, 64)
}

// column numbers for gotable report
const (
	rrRName       = 0
	rrRType       = iota
	rrSqFt        = iota
	rrDescr       = iota
	rrUsers       = iota
	rrPayors      = iota
	rrRAgr        = iota
	rrUsePeriod   = iota
	rrRentPeriod  = iota
	rrRAgrStart   = iota
	rrRAgrStop    = iota
	rrRentCycle   = iota
	rrGSRRate     = iota
	rrGSRAmt      = iota
	rrIncOff      = iota
	rrAmtDue      = iota
	rrPmtRcvd     = iota
	rrBeginRcv    = iota
	rrChgRcv      = iota
	rrEndRcv      = iota
	rrBeginSecDep = iota
	rrChgSecDep   = iota
	rrEndSecDep   = iota
)

// rrTableAddRow adds row in gotable struct with information
// given by rentableRow struct
func rrTableAddRow(tbl *gotable.Table, q rentableRow) {

	tbl.AddRow()
	tbl.Puts(-1, rrRName, q.RentableName.String)
	tbl.Puts(-1, rrRType, q.RentableType.String)
	tbl.Puts(-1, rrSqFt, int64ToStr(q.Sqft.Int64, true))
	tbl.Puts(-1, rrDescr, q.Description.String)
	tbl.Puts(-1, rrUsers, q.Users.String)
	tbl.Puts(-1, rrPayors, q.Payors.String)
	raidStr := int64ToStr(q.RAID.Int64, true)
	raStr := ""
	if len(raidStr) > 0 {
		raStr = "RA-" + raidStr
	}
	tbl.Puts(-1, rrRAgr, raStr)
	tbl.Puts(-1, rrUsePeriod, q.UsePeriod)
	tbl.Puts(-1, rrRentPeriod, q.RentPeriod)
	tbl.Puts(-1, rrRentCycle, q.RentCycleStr)
	tbl.Puts(-1, rrGSRRate, float64ToStr(q.GSR.Float64, false))
	tbl.Puts(-1, rrGSRAmt, float64ToStr(q.PeriodGSR.Float64, false))
	tbl.Puts(-1, rrIncOff, float64ToStr(q.IncomeOffsets.Float64, false))
	tbl.Puts(-1, rrAmtDue, float64ToStr(q.AmountDue.Float64, false))
	tbl.Puts(-1, rrPmtRcvd, float64ToStr(q.PaymentsApplied.Float64, false))
	tbl.Puts(-1, rrBeginRcv, float64ToStr(q.BeginningRcv.Float64, false))
	tbl.Puts(-1, rrChgRcv, float64ToStr(q.ChangeInRcv.Float64, false))
	tbl.Puts(-1, rrEndRcv, float64ToStr(q.EndingRcv.Float64, false))
	tbl.Puts(-1, rrBeginSecDep, float64ToStr(q.BeginningSecDep.Float64, false))
	tbl.Puts(-1, rrChgSecDep, float64ToStr(q.ChangeInSecDep.Float64, false))
	tbl.Puts(-1, rrEndSecDep, float64ToStr(q.EndingSecDep.Float64, false))
}

// handleGaps identifies periods during which the Rentable is not
// covered by a RentalAgreement. It updates the list with entries
// describing the gaps
//----------------------------------------------------------------------
func handleGaps(sl *[]rentableRow, d1, d2 *time.Time) {
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
		var r rentableRow
		r.PossessionStart.Scan(b[i].D1)
		r.PossessionStop.Scan(b[i].D2)
		r.Description.Scan("Vacancy")
		r.UsePeriod = FmtRRDatePeriod(&b[i].D1, &b[i].D2)
		(*sl) = append((*sl), r)
	}
}

// FmtRRDatePeriod formats a start and end time as needed byt the
// column headers in the RentRoll view/report
//
// INPUT
// d1 - start of period
// d2 - end of period
//
// RETURN
// string with formated dates
//----------------------------------------------------------------------
func FmtRRDatePeriod(d1, d2 *time.Time) string {
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
// sublist = the slice of rentableRow structs
// nq = the current entry but not yet added to sublist
//
// RETURN
// void
//----------------------------------------------------------------------
func setRRDatePeriodString(tbl *gotable.Table, nq *rentableRow) {
	showDates := true // only list dates if the rental agreement changed
	if len(tbl.Row) > 0 {
		prevRAgr := tbl.Get(len(tbl.Row)-1, rrRAgr)
		showDates = prevRAgr.Sval != "RA-"+strconv.FormatInt(nq.RAID.Int64, 10)
	}
	SetRRDateStrings(showDates, &nq.UsePeriod, &nq.RentPeriod,
		&nq.PossessionStart.Time, &nq.PossessionStop.Time, &nq.RentStart.Time, &nq.RentStop.Time)
}

// SetRRDateStrings updates the two supplied date strings if showDates is true
func SetRRDateStrings(showDates bool, s1, s2 *string, t1, t2, t3, t4 *time.Time) {
	if showDates {
		(*s1) = FmtRRDatePeriod(t1, t2)
		(*s2) = FmtRRDatePeriod(t3, t4)
	}
}
