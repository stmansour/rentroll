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
	"GROUP_CONCAT(DISTINCT CASE WHEN Payor.IsCompany > 0 THEN Payor.CompanyName ELSE CONCAT(Payor.FirstName, ' ', Payor.LastName) END SEPARATOR ', ') AS Payors",
	"GROUP_CONCAT(DISTINCT CASE WHEN User.IsCompany > 0 THEN User.CompanyName ELSE CONCAT(User.FirstName, ' ', User.LastName) END SEPARATOR ', ') AS Users",
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
	"BID":          "",
}

// ------- Rentables Assessments Query components -------

// RentablesAsmtFields holds the fields need to be fetched by
// rentables assessment query
var RentablesAsmtFields = rlib.SelectQueryFields{
	"Assessments.RAID",
	"AR.Name as Description",
	"Assessments.Amount as AmountDue",
	"SUM(ReceiptAllocation.Amount) as PaymentsApplied",
}

// RentablesAsmtQuery - query execution plan for rentable assessments
var RentablesAsmtQuery = `
    SELECT {{.SelectClause}}
    FROM Rentable
    LEFT JOIN Assessments ON (Assessments.RID=Rentable.RID AND (Assessments.FLAGS & 4)=0 AND "{{.DtStart}}" <= Start AND Stop < "{{.DtStop}}" AND (RentCycle=0 OR (RentCycle>0 AND PASMID!=0)))
    LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.ASMID=Assessments.ASMID AND "{{.DtStart}}" <= ReceiptAllocation.Dt AND ReceiptAllocation.Dt < "{{.DtStop}}")
    LEFT JOIN AR ON AR.ARID=Assessments.ARID
    WHERE {{.WhereClause}}
    GROUP BY Rentable.RID, Assessments.ASMID
    ORDER BY {{.OrderClause}};`

// RentablesAsmtQueryClause - query clause for rentable assessments
var RentablesAsmtQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(RentablesAsmtFields, ","),
	"WhereClause":  "Rentable.BID=%d", // needs to be replace %d with some BID in query execution plan
	"OrderClause":  "Rentable.RID ASC, Assessments.Amount DESC",
	"DtStart":      "",
	"DtStop":       "",
	"BID":          "",
}

// ------- Rentables No Assessments Query components -------

// RentablesNoAsmtFields holds the list of fields need to be fetched by
// rentables noasseesment query
var RentablesNoAsmtFields = rlib.SelectQueryFields{
	"AR.Name as Description",
	"ReceiptAllocation.RAID as RAID",
	"ReceiptAllocation.Amount as PaymentsApplied",
}

// RentablesNoAsmtQuery - the query execution plan for
// rentables noassessments
var RentablesNoAsmtQuery = `
    SELECT {{.SelectClause}} FROM ReceiptAllocation
    LEFT JOIN RentalAgreementRentables ON (ReceiptAllocation.RAID=RentalAgreementRentables.RAID AND "{{.DtStart}}" <= RentalAgreementRentables.RARDtStop AND RentalAgreementRentables.RARDtStart < "{{.DtStop}}")
    LEFT JOIN Rentable ON (Rentable.RID=RentalAgreementRentables.RID AND Rentable.RID > 0)
    LEFT JOIN Receipt ON (RentalAgreementRentables.RAID=Receipt.RAID AND Receipt.FLAGS & 4=0 AND "{{.DtStart}}" <= Receipt.Dt AND Receipt.Dt < "{{.DtStop}}")
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
	"BID":          "",
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
	"CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END AS Payors",
}

// NoRIDAsmtQueryFieldMap fieldsMap for the no rentable assessment query
var NoRIDAsmtQueryFieldMap = rlib.SelectQueryFieldMap{
	"Description": {"ARID.Name"},
	"LastModTime": {"PaymentType.LastModTime"},
	"LastModBy":   {"PaymentType.LastModBy"},
	"CreateTS":    {"PaymentType.CreateTS"},
	"CreateBy":    {"PaymentType.CreateBy"},
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
	"BID":          "",
}

// ------- No Rentables No Assessments Query components -------

// NoRIDNoAsmtQuerySelectFields holds the list of fields need to be fetched by
// no rentable no assessment query
var NoRIDNoAsmtQuerySelectFields = rlib.SelectQueryFields{
	"ReceiptAllocation.BID",
	"ReceiptAllocation.RAID",
	"ReceiptAllocation.Amount",
	"AR.Name",
	"CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END AS Payors",
}

// NoRIDNoAsmQueryFieldMap holds the fieldmap for no rentable no assessment query
var NoRIDNoAsmQueryFieldMap = rlib.SelectQueryFieldMap{
	"Description": {"ARID.Name"},
	"LastModTime": {"PaymentType.LastModTime"},
	"LastModBy":   {"PaymentType.LastModBy"},
	"CreateTS":    {"PaymentType.CreateTS"},
	"CreateBy":    {"PaymentType.CreateBy"},
}

// NoRIDNoAsmtQuery - the query execution plan for no rentables no assessments part
var NoRIDNoAsmtQuery = `
    SELECT {{.SelectClause}}
    FROM ReceiptAllocation
    LEFT JOIN RentalAgreementRentables ON RentalAgreementRentables.RAID=ReceiptAllocation.RAID
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
	"BID":          "",
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
	fmt.Printf("Entered in : %s\n", funcname)

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
	SqFtStr         string          // string representation of Sq ft
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
//        For ex. I don't know!!
func RRReportTable(ri *ReporterInfo) gotable.Table {
	const funcname = "RRReportTable"
	var (
		err     error
		startDt = ri.D1
		stopDt  = ri.D2
		tbl     = getRRTable() // gotable init for this report
		// totalErrs        = 0
		customAttrRTSqft = "Square Feet"
	)
	fmt.Printf("Entered in %s", funcname)

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

	// set table title, sections
	err = TableReportHeaderBlock(&tbl, "Rentroll", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// Add columns to the table
	tbl.AddColumn("Rentable", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                   // column for the Rentable name
	tbl.AddColumn("Rentable Type", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)              // RentableType name
	tbl.AddColumn("SqFt", 5, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)                          // the Custom Attribute "Square Feet"
	tbl.AddColumn("Description", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                // the Custom Attribute "Square Feet"
	tbl.AddColumn("Users", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                      // Users of this rentable
	tbl.AddColumn("Payors", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                     // Users of this rentable
	tbl.AddColumn("Rental Agreement", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)           // the Rental Agreement id
	tbl.AddColumn("Use Period", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                 // the use period
	tbl.AddColumn("Rent Period", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)                  // the rent period
	tbl.AddColumn("Rent Cycle", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                 // the rent cycle
	tbl.AddColumn("GSR Rate", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)                   // gross scheduled rent
	tbl.AddColumn("Period GSR", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)                 // gross scheduled rent
	tbl.AddColumn("Income Offsets", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)             // GL Account
	tbl.AddColumn("Payments Applied", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)           // contract rent amounts
	tbl.AddColumn("Beginning Receivable", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Change In Receivable", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Ending Receivable", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)          // account for the associated RentalAgreement
	tbl.AddColumn("Beginning Security Deposit", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Change In Security Deposit", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Ending Security Deposit", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)    // account for the associated RentalAgreement

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
			q.UsePeriod = q.PossessionStart.Time.Format(rlib.RRDATEFMT3) + "\n - " + q.PossessionStop.Time.Format(rlib.RRDATEFMT3)
		}
		for freqStr, freqNo := range rlib.CycleFreqMap {
			if q.RentCycle.Int64 == freqNo {
				q.RentCycleStr = freqStr
			}
		}
		if q.Sqft.Int64 > 0 {
			q.SqFtStr = strconv.FormatInt(q.Sqft.Int64, 10)
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
			err = rentablesAsmtRows.Scan(&nq.RAID, &nq.Description, &nq.AmountDue, &nq.PaymentsApplied)
			if err != nil {
				tbl.SetSection3(err.Error())
				return tbl
			}
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
			err = rentablesNoAsmtRows.Scan(&nq.Description, &nq.RAID, &nq.PaymentsApplied)
			if err != nil {
				tbl.SetSection3(err.Error())
				return tbl
			}
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

		// add line Before the subtotal
		tbl.AddLineAfter(len(tbl.Row) - 1)
		rrTableAddRow(&tbl, sub)
		childCount++
		// SHOULD WE ADD LINE AFTER SUBTOTAL ROW??

		// Add one blank ROW
		tbl.AddLineAfter(len(tbl.Row) - 1)
		/*rrTableAddRow(&tbl, rentableRow{})
		  childCount++*/

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
		err := noRIDAsmtRows.Scan(&q.BID, &q.ASMID, &q.Description, &q.AmountDue, &q.PaymentsApplied, &q.RAID, &q.Payors)
		if err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		rrTableAddRow(&tbl, q)
		count++
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
		err := noRIDNoAsmtRows.Scan(&q.BID, &q.RAID, &q.PaymentsApplied, &q.Description, &q.Payors)
		if err != nil {
			tbl.SetSection3(err.Error())
			return tbl
		}
		rrTableAddRow(&tbl, q)
		count++
	}

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

// rrTableAddRow adds row in gotable struct with information
// given by rentableRow struct
func rrTableAddRow(tbl *gotable.Table, q rentableRow) {

	// column numbers for gotable report
	const (
		RName     = 0
		RType     = iota
		SqFt      = iota
		Descr     = iota
		Users     = iota
		Payors    = iota
		RAgr      = iota
		UsePeriod = iota
		// UseStart     = iota
		// UseStop      = iota
		RentPeriod = iota
		// RentStart    = iota
		// RentStop     = iota
		RAgrStart = iota
		RAgrStop  = iota
		RentCycle = iota
		GSRRate   = iota
		GSRAmt    = iota
		IncOff    = iota
		AmtDue    = iota
		// ContractRent = iota
		// OtherInc     = iota
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
	tbl.Puts(-1, SqFt, q.SqFtStr) // even it has been defined as CELLINT, still you can put string content, STRANGE!!!
	tbl.Puts(-1, Descr, q.Description.String)
	tbl.Puts(-1, Users, q.Users.String)
	tbl.Puts(-1, Payors, q.Payors.String)
	tbl.Puts(-1, RAgr, fmt.Sprintf("RA-%d", q.RAID.Int64))
	tbl.Puts(-1, UsePeriod, q.UsePeriod)
	tbl.Puts(-1, RentPeriod, q.RentPeriod)
	tbl.Puts(-1, RentCycle, q.RentCycleStr)
	tbl.Putf(-1, GSRRate, q.GSR.Float64)
	tbl.Putf(-1, GSRAmt, q.PeriodGSR.Float64)
	tbl.Putf(-1, IncOff, q.IncomeOffsets.Float64)
	tbl.Putf(-1, AmtDue, q.AmountDue.Float64)
	tbl.Putf(-1, PmtRcvd, q.PaymentsApplied.Float64)
	tbl.Putf(-1, BeginRcv, q.BeginningRcv.Float64)
	tbl.Putf(-1, ChgRcv, q.ChangeInRcv.Float64)
	tbl.Putf(-1, EndRcv, q.EndingRcv.Float64)
	tbl.Putf(-1, BeginSecDep, q.BeginningSecDep.Float64)
	tbl.Putf(-1, ChgSecDep, q.ChangeInSecDep.Float64)
	tbl.Putf(-1, EndSecDep, q.EndingSecDep.Float64)
}
