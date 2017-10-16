package rrpt

import (
	"database/sql"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"sort"
	"strconv"
	"strings"
	"time"
)

// RentableSection etc all constants to represent to which section it belongs
const (
	RentableSection   = 1
	NoRentableSection = 2
)

// ------- Rentable Section components -------

// RentableSectionFieldsMap holds the map of field (to be shown on grid)
// for the first section (rentables Part)
var RentableSectionFieldsMap = rlib.SelectQueryFieldMap{
	"RID":             {"Rentable.RID"},                  // Rentable ID
	"RentableName":    {"Rentable.RentableName"},         // Rentable Name
	"RTID":            {"RentableTypes.RTID"},            // RentableTypes ID
	"RentableType":    {"RentableTypes.Name"},            // RentableTypes Name
	"RentCycle":       {"RentableTypes.RentCycle"},       // Rent Cycle
	"Status":          {"RentableStatus.UseStatus"},      // Rentable Status
	"MarketRate":      {"RentableMarketRate.MarketRate"}, // Rentable Market Rate
	"RAID":            {"RentalAgreement.RAID"},          // RentalAgreement ID
	"AgreementStart":  {"RentalAgreement.AgreementStart"},
	"AgreementStop":   {"RentalAgreement.AgreementStop"},
	"PossessionStart": {"RentalAgreement.PossessionStart"},
	"PossessionStop":  {"RentalAgreement.PossessionStop"},
	"RentStart":       {"RentalAgreement.RentStart"},
	"RentStop":        {"RentalAgreement.RentStop"},
	"Payors":          {"Payor.FirstName", "Payor.LastName", "Payor.CompanyName"},
	"Users":           {"User.FirstName", "User.LastName", "User.CompanyName"},
	"ASMID":           {"Assessments.ASMID"},
	"AmountDue":       {"Assessments.Amount"},
	"Description":     {"AR.Name"},
	// "RCPAID":          {"ReceiptAllocation.RCPAID"},
	"PaymentsApplied": {"ReceiptAllocation.Amount"}, // confused, is it true?
}

// RentableSectionFields holds the selectClause for the RentableSectionQuery
var RentableSectionFields = rlib.SelectQueryFields{
	"Rentable.RID",
	"Rentable.RentableName",
	"RentableTypes.RTID",
	"RentableTypes.Name AS RentableType",
	"RentableTypes.RentCycle",
	"RentableStatus.UseStatus AS Status",
	"RentableMarketRate.MarketRate",
	"RentalAgreement.RAID",
	"RentalAgreement.AgreementStart",
	"RentalAgreement.AgreementStop",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"GROUP_CONCAT(DISTINCT CASE WHEN Payor.IsCompany > 0 THEN Payor.CompanyName ELSE CONCAT(Payor.FirstName, ' ', Payor.LastName) END ORDER BY Payor.LastName ASC, Payor.FirstName ASC, Payor.CompanyName ASC SEPARATOR ', ') AS Payors",
	"GROUP_CONCAT(DISTINCT CASE WHEN User.IsCompany > 0 THEN User.CompanyName ELSE CONCAT(User.FirstName, ' ', User.LastName) END ORDER BY User.LastName ASC, User.FirstName ASC, User.CompanyName ASC SEPARATOR ', ' ) AS Users",
	"Assessments.ASMID",
	"Assessments.Amount AS AmountDue",
	"AR.Name AS Description",
	// "ReceiptAllocation.RCPAID",
	"SUM(ReceiptAllocation.Amount) AS PaymentsApplied",
}

// RentableSectionQuery pulls out all rentable section records for given date range
// for the rentroll report
// Uses @DtStart and @DtStop mysql variables, so it needs to be set before
// executing this query
var RentableSectionQuery = `
SELECT DISTINCT
    {{.SelectClause}}
FROM
    Rentable
        LEFT JOIN
    RentalAgreementRentables ON (RentalAgreementRentables.RID = Rentable.RID
        AND @DtStart <= RentalAgreementRentables.RARDtStop
        AND @DtStop > RentalAgreementRentables.RARDtStart)
        LEFT JOIN
    RentalAgreement ON (RentalAgreement.RAID = RentalAgreementRentables.RAID
        AND @DtStart <= RentalAgreement.AgreementStop
        AND @DtStop > RentalAgreement.AgreementStart)
        LEFT JOIN
    RentalAgreementPayors ON (RentalAgreement.RAID = RentalAgreementPayors.RAID
        AND @DtStart <= RentalAgreementPayors.DtStop
        AND @DtStop > RentalAgreementPayors.DtStart)
        LEFT JOIN
    Transactant AS Payor ON (Payor.TCID = RentalAgreementPayors.TCID
        AND Payor.BID = Rentable.BID)
        LEFT JOIN
    RentableUsers ON (RentableUsers.RID = Rentable.RID
        AND @DtStart <= RentableUsers.DtStop
        AND @DtStop > RentableUsers.DtStart)
        LEFT JOIN
    Transactant AS User ON (RentableUsers.TCID = User.TCID
        AND User.BID = Rentable.BID)
        LEFT JOIN
    RentableTypeRef ON RentableTypeRef.RID = Rentable.RID
        LEFT JOIN
    RentableTypes ON RentableTypes.RTID = RentableTypeRef.RTID
        LEFT JOIN
    RentableStatus ON (RentableStatus.RID = Rentable.RID
        AND @DtStart <= RentableStatus.DtStop
        AND @DtStop > RentableStatus.DtStart)
        LEFT JOIN
    RentableMarketRate ON (RentableMarketRate.RTID = RentableTypeRef.RTID
        AND @DtStart <= RentableMarketRate.DtStop
        AND @DtStop > RentableMarketRate.DtStart)
        LEFT JOIN
    Assessments ON (Assessments.RAID = RentalAgreement.RAID
        AND Assessments.RID = Rentable.RID
        AND (Assessments.FLAGS & 4) = 0
        AND @DtStart <= Assessments.Start
        AND @DtStop > Assessments.Stop
        AND (Assessments.RentCycle = 0
        OR (Assessments.RentCycle > 0
        AND Assessments.PASMID != 0)))
        LEFT JOIN
    AR ON AR.ARID = Assessments.ARID
        LEFT JOIN
    ReceiptAllocation ON (ReceiptAllocation.RAID = RentalAgreement.RAID
        AND @DtStart <= ReceiptAllocation.Dt
        AND ReceiptAllocation.Dt < @DtStop)
WHERE
    {{.WhereClause}}
GROUP BY Rentable.RID, RentalAgreement.RAID, Assessments.Amount DESC, ReceiptAllocation.RCPAID
ORDER BY {{.OrderClause}};`

// RentableSectionQueryClause -- query clause for the RentableSectionQuery
var RentableSectionQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(RentableSectionFields, ","),
	"WhereClause":  "Rentable.BID=%d",
	"OrderClause":  "Rentable.RentableName, RentalAgreement.AgreementStart, RentalAgreement.AgreementStop",
}

// RentableSectionRowScan scans a result from sql row and dump it in a RentRollReportRow struct
func RentableSectionRowScan(rows *sql.Rows, q *RentRollReportRow) error {
	return rows.Scan(&q.RID, &q.RentableName,
		&q.RTID, &q.RentableType, &q.RentCycle, &q.Status, &q.GSR,
		&q.RAID, &q.AgreementStart, &q.AgreementStop,
		&q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop,
		&q.Payors, &q.Users, &q.ASMID, &q.AmountDue, &q.Description,
		&q.PaymentsApplied)
}

// ------- NO Rentable Section components -------

// NoRentableSectionFieldsMap holds the map of field (to be shown on grid)
// for the second section (No Rentables Part)
var NoRentableSectionFieldsMap = rlib.SelectQueryFieldMap{
	"RAID":            {"RentalAgreement.RAID"}, // RentalAgreement ID
	"AgreementStart":  {"RentalAgreement.AgreementStart"},
	"AgreementStop":   {"RentalAgreement.AgreementStop"},
	"PossessionStart": {"RentalAgreement.PossessionStart"},
	"PossessionStop":  {"RentalAgreement.PossessionStop"},
	"RentStart":       {"RentalAgreement.RentStart"},
	"RentStop":        {"RentalAgreement.RentStop"},
	"Payors":          {"Payor.FirstName", "Payor.LastName", "Payor.CompanyName"},
	"ASMID":           {"Assessments.ASMID"},
	"AmountDue":       {"Assessments.Amount"},
	"Description":     {"AR.Name"},
	// "RCPAID":          {"ReceiptAllocation.RCPAID"},
	"PaymentsApplied": {"ReceiptAllocation.Amount"}, // confused, is it true?
}

// NoRentableSectionFields - holds the list of fields need to be selected for No Rentable section
var NoRentableSectionFields = rlib.SelectQueryFields{
	"RentalAgreement.RAID",
	"RentalAgreement.AgreementStart",
	"RentalAgreement.AgreementStop",
	"RentalAgreement.PossessionStart",
	"RentalAgreement.PossessionStop",
	"RentalAgreement.RentStart",
	"RentalAgreement.RentStop",
	"GROUP_CONCAT(DISTINCT CASE WHEN Payor.IsCompany > 0 THEN Payor.CompanyName ELSE CONCAT(Payor.FirstName, ' ', Payor.LastName) END ORDER BY Payor.LastName ASC, Payor.FirstName ASC, Payor.CompanyName ASC SEPARATOR ', ') AS Payors",
	"Assessments.ASMID",
	"Assessments.Amount AS AmountDue",
	"AR.Name AS Description",
	// "ReceiptAllocation.RCPAID",
	"SUM(ReceiptAllocation.Amount) AS PaymentsApplied",
	// "GROUP_CONCAT(DISTINCT ReceiptAllocation.RCPAID SEPARATOR ', ') AS RCPAIDList",
}

// NoRentableSectionQuery - query execution plan for noRentable section
var NoRentableSectionQuery = `
SELECT DISTINCT
    {{.SelectClause}}
FROM
    ReceiptAllocation
        INNER JOIN
    Receipt ON (Receipt.RCPTID = ReceiptAllocation.RCPTID
        AND @DtStart <= Receipt.Dt
        AND Receipt.Dt < @DtStop)
        LEFT JOIN
    Transactant AS Payor ON (Payor.TCID = Receipt.TCID)
        INNER JOIN
    RentalAgreement ON (RentalAgreement.RAID = ReceiptAllocation.RAID
        AND RentalAgreement.RAID > 0)
        LEFT JOIN
    RentalAgreementRentables ON (RentalAgreementRentables.RAID = RentalAgreement.RAID)
        LEFT JOIN
    Assessments ON (Assessments.RAID = RentalAgreement.RAID
        AND (Assessments.FLAGS & 4) = 0
        AND Assessments.RID = 0
        AND @DtStart <= Assessments.Stop
        AND @DtStop > Assessments.Start
        AND (Assessments.RentCycle = 0
        OR (Assessments.RentCycle > 0
        AND Assessments.PASMID != 0)))
        LEFT JOIN
    AR ON (AR.ARID = Assessments.ARID
        OR (AR.ARID = Receipt.ARID AND AR.FLAGS = 5))
WHERE
    @DtStart <= ReceiptAllocation.Dt
        AND ReceiptAllocation.Dt < @DtStop AND Receipt.FLAGS & 4 = 0
GROUP BY RentalAgreement.RAID, Assessments.ASMID
ORDER BY {{.OrderClause}};`

// NoRentableSectionQueryClause -- query clause for the NoRentableSectionQuery
var NoRentableSectionQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(NoRentableSectionFields, ","),
	"WhereClause":  "ReceiptAllocation.BID=%d AND Receipt.FLAGS&4=0 AND @DtStart <= ReceiptAllocation.Dt AND ReceiptAllocation.Dt < @DtStop AND (RentalAgreementRentables.RID=0 OR RentalAgreementRentables.RID IS NULL)",
	"OrderClause":  "RentalAgreement.RAID, Assessments.Amount DESC",
}

// NoRentableSectionRowScan scans a result from sql row and dump it in a RentRollReportRow struct
func NoRentableSectionRowScan(rows *sql.Rows, q *RentRollReportRow) error {
	return rows.Scan(&q.RAID, &q.AgreementStart, &q.AgreementStop,
		&q.PossessionStart, &q.PossessionStop, &q.RentStart, &q.RentStop,
		&q.Payors, &q.ASMID, &q.AmountDue, &q.Description,
		&q.PaymentsApplied)
}

// formatReportSectionQuery returns the formatted query
// with given limit, offset if applicable for given section
// If given part doesn't exist then it will return nil with error
func formatReportSectionQuery(
	rentrollSection int, BID int64, d1, d2 time.Time,
	additionalWhere, orderBy string, limit, offset int,
) (string, error) {
	const funcname = "formatReportSectionQuery"
	var (
		qry   string
		qc    rlib.QueryClause
		where string
		order string
	)
	rlib.Console("Entered in : %s\n", funcname)

	// based on part, decide query and queryClause
	switch rentrollSection {
	case RentableSection:
		qry = RentableSectionQuery
		qc = rlib.GetQueryClauseCopy(RentableSectionQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID)
		break
	case NoRentableSection:
		qry = NoRentableSectionQuery
		qc = rlib.GetQueryClauseCopy(NoRentableSectionQueryClause)
		where = fmt.Sprintf(qc["WhereClause"], BID)
		break
	default:
		return "", fmt.Errorf("No such section (%d) exists in rentroll report", rentrollSection)
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
	return rlib.RenderSQLQuery(qry, qc), nil

	// tInit := time.Now()
	// qExec, err := rlib.RRdb.Dbrr.Query(dbQry)
	// diff := time.Since(tInit)
	// rlib.Console("\nQQQQQQuery Time diff for %s is %s\n\n", rrPart, diff.String())
	// return qExec, err
}

// RentRollReportRow represents the row that holds the data for rentroll report
// it could be used by rentroll webservice view as well as for the gotable report
type RentRollReportRow struct {
	Recid                    int64            `json:"recid"` // this is to support the w2ui form
	BID                      int64            // Business (so that we can process by Business)
	RID                      int64            // The rentable
	RTID                     int64            // The rentable type
	RAID                     rlib.NullInt64   // Rental Agreement
	RARID                    rlib.NullInt64   // rental agreement rentable id
	RAIDStr                  string           // RentalAgreement representational string
	ASMID                    rlib.NullInt64   // Assessment
	RentableName             rlib.NullString  // Name of the rentable
	RentableType             rlib.NullString  // Name of the rentable type
	Sqft                     rlib.NullInt64   // rentable square feet
	Description              rlib.NullString  // account rule name
	RentCycle                rlib.NullInt64   // Rent Cycle
	RentCycleStr             string           // String representation of Rent Cycle
	Status                   rlib.NullInt64   // Rentable status
	AgreementStart           rlib.NullDate    // start date for RA
	AgreementStop            rlib.NullDate    // stop date for RA
	AgreementPeriod          string           // text representation of Rental Agreement time period
	PossessionStart          rlib.NullDate    // start date for Occupancy
	PossessionStop           rlib.NullDate    // stop date for Occupancy
	UsePeriod                string           // text representation of Occupancy(or use) time period
	RentStart                rlib.NullDate    // start date for Rent
	RentStop                 rlib.NullDate    // stop date for Rent
	RentPeriod               string           // text representation of Rent time period
	Payors                   rlib.NullString  // payors list attached with this RA within same time
	Users                    rlib.NullString  // users associated with the rentable
	GSR                      rlib.NullFloat64 // Gross scheduled rate
	PeriodGSR                rlib.NullFloat64 // Periodic gross scheduled rate
	IncomeOffsets            rlib.NullFloat64 // Income Offset amount
	AmountDue                rlib.NullFloat64 // Amount needs to be paid by Payor(s)
	PaymentsApplied          rlib.NullFloat64 // Amount collected by Payor(s) for Assessments
	BeginningRcv             rlib.NullFloat64 // Receivable amount at beginning period
	ChangeInRcv              rlib.NullFloat64 // Change in receivable
	EndingRcv                rlib.NullFloat64 // Ending receivable
	BeginningSecDep          rlib.NullFloat64 // Beginning security deposit
	ChangeInSecDep           rlib.NullFloat64 // Change in security deposit
	EndingSecDep             rlib.NullFloat64 // Ending security deposit
	IsMainRow                bool             // is main row
	IsRentableSectionMainRow bool             // is rentable section main row which holds all static data
	IsSubTotalRow            bool             // is sustotal row
	IsBlankRow               bool             // is blank row
	IsNoRentableSectionRow   bool             // is "No Rentable" row
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

// formatSubTotalRow formats subtotal row by picking only meaningful
// fields from RentRollReportRow struct
func formatSubTotalRow(subTotalRow *RentRollReportRow, startDt, stopDt time.Time) {
	const funcname = "formatSubTotalRow"
	var (
		err error
		d70 = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	)

	// mark flag
	subTotalRow.IsSubTotalRow = true

	// Description
	subTotalRow.Description.Scan("Subtotal")

	// AmountDue
	subTotalRow.AmountDue.Scan(subTotalRow.AmountDue.Float64)

	// PaymentsApplied
	subTotalRow.PaymentsApplied.Scan(subTotalRow.PaymentsApplied.Float64)

	// PeriodGSR
	subTotalRow.PeriodGSR.Scan(subTotalRow.PeriodGSR.Float64)

	// IncomeOffsets
	subTotalRow.IncomeOffsets.Scan(subTotalRow.IncomeOffsets.Float64)

	// BeginningRcv, EndingRcv
	subTotalRow.BeginningRcv.Float64, subTotalRow.EndingRcv.Float64, err =
		rlib.GetBeginEndRARBalance(subTotalRow.BID, subTotalRow.RID,
			subTotalRow.RAID.Int64, &startDt, &stopDt)
	if err != nil {
		rlib.Console("%s: Error while calculating BeginningRcv, EndingRcv:: %s", funcname, err.Error())
	} else {
		subTotalRow.BeginningRcv.Valid = true
		subTotalRow.EndingRcv.Valid = true
	}

	// ChangeInRcv
	subTotalRow.ChangeInRcv.Scan(subTotalRow.EndingRcv.Float64 - subTotalRow.BeginningRcv.Float64)

	// BeginningSecDep
	subTotalRow.BeginningSecDep.Float64, err = rlib.GetSecDepBalance(
		subTotalRow.BID, subTotalRow.RAID.Int64, subTotalRow.RID, &d70, &startDt)
	if err != nil {
		rlib.Console("%s: Error while calculating BeginningSecDep:: %s", funcname, err.Error())
	} else {
		subTotalRow.BeginningSecDep.Valid = true
	}

	// Change in SecDep
	subTotalRow.ChangeInSecDep.Float64, err = rlib.GetSecDepBalance(
		subTotalRow.BID, subTotalRow.RAID.Int64, subTotalRow.RID, &startDt, &stopDt)
	if err != nil {
		rlib.Console("%s: Error while calculating BeginningSecDep:: %s", funcname, err.Error())
	} else {
		subTotalRow.ChangeInSecDep.Valid = true
	}

	// EndingSecDep
	subTotalRow.EndingSecDep.Scan(subTotalRow.BeginningSecDep.Float64 + subTotalRow.ChangeInSecDep.Float64)
}

// setRRDatePeriodString updates the "r" UsePeriod and RentPeriod members
// if it is either the first row in resultList or if the RentalAgreement has
// changed since the last entry in list.
//
// INPUT
// r = the current entry but not yet added to sublist
// lastRow = the last entry from the result list
//
// RETURN
// void
//----------------------------------------------------------------------
func setRRDatePeriodString(r, lastRow *RentRollReportRow) {
	if lastRow.RAID.Int64 != r.RAID.Int64 {
		r.AgreementPeriod = fmtRRDatePeriod(&r.PossessionStart.Time, &r.PossessionStop.Time)
		r.UsePeriod = fmtRRDatePeriod(&r.PossessionStart.Time, &r.PossessionStop.Time)
		r.RentPeriod = fmtRRDatePeriod(&r.RentStart.Time, &r.RentStop.Time)
	} else {
		r.AgreementPeriod = ""
		r.RentPeriod = ""
		r.UsePeriod = ""
		r.Payors.String = ""
		r.Payors.Valid = false
		r.Users.String = ""
		r.Users.Valid = false
		r.RentCycleStr = ""
		r.RAIDStr = ""
	}
}

// formatRentableSectionChildRow formats new Renable Section Row
// into Child Row pattern
func formatRentableSectionChildRow(r, lastRow *RentRollReportRow) {
	// const funcname = "formatRentableSectionChildRow"

	// set some values to blank
	r.RentableName.String = ""
	r.RentableName.Valid = false
	r.RentableType.String = ""
	r.RentableType.Valid = false
	r.Sqft.Int64 = 0
	r.Sqft.Valid = false
	r.Description.String = ""
	r.Description.Valid = false
	r.IsRentableSectionMainRow = false
	r.IsMainRow = false
	r.GSR.Float64 = 0
	r.GSR.Valid = false

	setRRDatePeriodString(r, lastRow)
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
	rentableSectionWhr, rentableSectionOdr string, rentableSectionOffset int,
	norentableSectionWhr, norentableSectionOdr string, noRentableSectionOffset int,
) ([]RentRollReportRow, error) {

	const funcname = "RRReportRows"
	var (
		err                        error
		d1Str                      = startDt.Format(rlib.RRDATEFMTSQL)
		d2Str                      = stopDt.Format(rlib.RRDATEFMTSQL)
		customAttrRTSqft           = "Square Feet" // custom attribute for all rentables
		xbiz                       rlib.XBusiness
		rptMainRowsCount           = 0                                   // report main rows count
		rentableRowsMap            = make(map[int64][]RentRollReportRow) // per rentable it will hold sublist of rows
		rentableRowsMapKeys        = []int64{}
		noRentableSectionRowsLimit = 0                                  // limit on "NO rentable Section" rows
		grandTTL                   = RentRollReportRow{IsMainRow: true} // grand total row
	)
	rlib.Console("Entered in %s\n", funcname)

	// init some structure
	reportRows := []RentRollReportRow{}
	rlib.InitBizInternals(BID, &xbiz) // init some business internals first

	//=========================================================================
	//                           RENTABLE SECTION                            //
	//=========================================================================

	// if there is no limit then it is meaningless having a value for below variables
	if pageRowsLimit <= 0 {
		rentableSectionOffset = -1
		pageRowsLimit = -1
	}

	// get formatted query - rentable section
	rentableSectionFmtQuery, err := formatReportSectionQuery(RentableSection, BID,
		startDt, stopDt,
		rentableSectionWhr, rentableSectionOdr,
		pageRowsLimit, rentableSectionOffset)
	if err != nil {
		return reportRows, err
	}

	// start transaction
	rentableSectionTx, err := rlib.RRdb.Dbrr.Begin()
	if err != nil {
		return reportRows, err
	}
	// NOW, set mysql variables for date values
	if _, err = rentableSectionTx.Exec("SET @DtStart:=?", d1Str); err != nil {
		rentableSectionTx.Rollback()
		return reportRows, err
	}
	if _, err = rentableSectionTx.Exec("SET @DtStop:=?", d2Str); err != nil {
		rentableSectionTx.Rollback()
		return reportRows, err
	}

	// Execute query in current transaction for Rentable section
	rentableSectionRows, err := rentableSectionTx.Query(rentableSectionFmtQuery)
	if err != nil {
		rentableSectionTx.Rollback()
		return reportRows, err
	}
	defer rentableSectionRows.Close()

	// ============================
	// LOOP THROUGH RENTABLES ROWS
	// ============================
	rentableSectionCount := 0

	for rentableSectionRows.Next() {
		// just assume that it is MainRow, if later encountered that it is child row
		// then "formatRentableSectionChildRow" function would take care of it. :)
		q := RentRollReportRow{IsMainRow: true, IsRentableSectionMainRow: true}
		if err = RentableSectionRowScan(rentableSectionRows, &q); err != nil {
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
		raidStr := int64ToStr(q.RAID.Int64, true)
		raStr := ""
		if len(raidStr) > 0 {
			raStr = "RA-" + raidStr
		}
		q.RAIDStr = raStr

		// get current row RID
		rowRID := q.RID

		// if key found from map, then it is child row, otherwise it is new rentable
		if _, ok := rentableRowsMap[rowRID]; ok {
			// it is child row, first format it
			formatRentableSectionChildRow(&q, &rentableRowsMap[rowRID][len(rentableRowsMap[rowRID])-1])
		} else { // new rentable row
			// store key in the mapKeys list
			rentableRowsMapKeys = append(rentableRowsMapKeys, rowRID)
			rptMainRowsCount++
		}

		// append new rentable row / formatted child row in map sublist
		rentableRowsMap[rowRID] = append(rentableRowsMap[rowRID], q)

		// update the rentableSectionCount only after adding the record
		rentableSectionCount++
	}
	// check for any errors from row results
	err = rentableSectionRows.Err()
	if err != nil {
		rentableSectionTx.Rollback()
		return reportRows, err
	}

	// commit rentable Section Transaction, finally
	if err = rentableSectionTx.Commit(); err != nil {
		rentableSectionTx.Rollback()
		return reportRows, err
	}

	rlib.Console("Added %d Rentable Section rows\n", rentableSectionCount)

	// sort the map keys first
	sort.Slice(rentableRowsMapKeys, func(i, j int) bool {
		return rentableRowsMapKeys[i] < rentableRowsMapKeys[j]
	})

	// loop through all rentables with map
	for _, RID := range rentableRowsMapKeys {
		// get the sublist from map
		rentableSubList := rentableRowsMap[RID]

		// first handle rentable gaps
		handleRentableGaps(BID, RID, &rentableSubList, startDt, stopDt)

		// sort the list of all rows per rentable
		sort.Slice(rentableSubList, func(i, j int) bool {
			return rentableSubList[i].PossessionStart.Time.Before(
				rentableSubList[j].PossessionStart.Time)
		})

		// now add subtotal row
		addSubTotals(&rentableSubList, &grandTTL, startDt, stopDt)

		// now add blankRow
		rentableSubList = append(rentableSubList, RentRollReportRow{IsBlankRow: true})

		// now add this rentableRowsList to original result row list
		reportRows = append(reportRows, rentableSubList...)
	}

	// if for given limit, rows are feed within page then return
	if isReportComplete(pageRowsLimit, rptMainRowsCount) {
		return reportRows, err
	}

	//=========================================================================
	//                        NO RENTABLE SECTION                            //
	//=========================================================================

	// if no limit then reset the values
	if pageRowsLimit <= 0 {
		norentableSectionWhr = ""
		norentableSectionOdr = ""
		noRentableSectionRowsLimit = -1
		noRentableSectionOffset = -1
	} else {
		noRentableSectionRowsLimit = pageRowsLimit - len(reportRows)
		if noRentableSectionRowsLimit < 0 {
			noRentableSectionRowsLimit = 0 // make sure it doesn't have minus value
		}
	}

	// get formatted query string
	noRentableSectionFmtQuery, err := formatReportSectionQuery(NoRentableSection, BID,
		startDt, stopDt,
		norentableSectionWhr, norentableSectionOdr,
		noRentableSectionRowsLimit, noRentableSectionOffset)
	if err != nil {
		return reportRows, err
	}

	// start transaction
	noRentableSectionTx, err := rlib.RRdb.Dbrr.Begin()
	if err != nil {
		return reportRows, err
	}
	// NOW, set mysql variables for date values
	if _, err = noRentableSectionTx.Exec("SET @DtStart:=?", d1Str); err != nil {
		noRentableSectionTx.Rollback()
		return reportRows, err
	}
	if _, err = noRentableSectionTx.Exec("SET @DtStop:=?", d2Str); err != nil {
		noRentableSectionTx.Rollback()
		return reportRows, err
	}

	// Execute query in current transaction for Rentable section
	noRentableSectionRows, err := noRentableSectionTx.Query(noRentableSectionFmtQuery)
	if err != nil {
		noRentableSectionTx.Rollback()
		return reportRows, err
	}
	defer noRentableSectionRows.Close()

	// ======================================
	// LOOP THROUGH NO RENTABLE SECTION ROWS
	// ======================================
	noRentableSectionRowsCount := 0
	for noRentableSectionRows.Next() {
		q := RentRollReportRow{IsMainRow: true, IsNoRentableSectionRow: true}

		if err = NoRentableSectionRowScan(noRentableSectionRows, &q); err != nil {
			return reportRows, err
		}

		setRRDatePeriodString(&q, &reportRows[len(reportRows)-1])

		// APPEND NO-RID-ASMT ROW IN LIST
		reportRows = append(reportRows, q)

		// add subTotal amounts to grand total record
		updateGrandTotals(&grandTTL, &q)

		noRentableSectionRowsCount++
	}
	// check for any errors from rows results
	err = noRentableSectionRows.Err()
	if err != nil {
		noRentableSectionTx.Rollback()
		return reportRows, err
	}

	// commit rentable Section Transaction, finally
	if err = noRentableSectionTx.Commit(); err != nil {
		noRentableSectionTx.Rollback()
		return reportRows, err
	}

	rlib.Console("Added noRID Asmt rows: %d", noRentableSectionRowsCount)
	rptMainRowsCount += noRentableSectionRowsCount // how many total rows have been added to list

	// ================
	// GRAND TOTAL ROW
	// ================
	if len(reportRows) > 0 {
		grandTTL.Description.Scan("Grand Total")
		reportRows = append(reportRows, grandTTL)
	}

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
		"", "", -1, // "rentables" Section
		"", "", -1, // "No Rentable Section"
	)

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

// addSubTotals does all subtotal calculations for the subtotal line
//-----------------------------------------------------------------------------
func addSubTotals(subRows *[]RentRollReportRow, g *RentRollReportRow, d1, d2 time.Time) {
	sub := RentRollReportRow{}

	for _, row := range *subRows {
		sub.AmountDue.Float64 += row.AmountDue.Float64
		sub.PaymentsApplied.Float64 += row.PaymentsApplied.Float64
		sub.PeriodGSR.Float64 += row.PeriodGSR.Float64
		sub.IncomeOffsets.Float64 += row.IncomeOffsets.Float64
	}

	formatSubTotalRow(&sub, d1, d2)

	// append to subRows List
	(*subRows) = append((*subRows), sub)

	// update grand total
	updateGrandTotals(g, &sub)

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
	tbl.Puts(-1, RAgr, q.RAIDStr)
	tbl.Puts(-1, UsePeriod, q.UsePeriod)
	tbl.Puts(-1, RentPeriod, q.RentPeriod)
	tbl.Puts(-1, RentCycle, q.RentCycleStr)
	if q.IsBlankRow {
		tbl.Puts(-1, GSRRate, float64ToStr(q.GSR.Float64, true))
		tbl.Puts(-1, GSRAmt, float64ToStr(q.PeriodGSR.Float64, true))
		tbl.Puts(-1, IncOff, float64ToStr(q.IncomeOffsets.Float64, true))
		tbl.Puts(-1, AmtDue, float64ToStr(q.AmountDue.Float64, true))
		tbl.Puts(-1, PmtRcvd, float64ToStr(q.PaymentsApplied.Float64, true))
		tbl.Puts(-1, BeginRcv, float64ToStr(q.BeginningRcv.Float64, true))
		tbl.Puts(-1, ChgRcv, float64ToStr(q.ChangeInRcv.Float64, true))
		tbl.Puts(-1, EndRcv, float64ToStr(q.EndingRcv.Float64, true))
		tbl.Puts(-1, BeginSecDep, float64ToStr(q.BeginningSecDep.Float64, true))
		tbl.Puts(-1, ChgSecDep, float64ToStr(q.ChangeInSecDep.Float64, true))
		tbl.Puts(-1, EndSecDep, float64ToStr(q.EndingSecDep.Float64, true))
	} else {
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
}

// handleRentableGaps identifies periods during which the Rentable is not
// covered by a RentalAgreement. It updates the list with entries
// describing the gaps
//----------------------------------------------------------------------
func handleRentableGaps(bid, rid int64, sl *[]RentRollReportRow, d1, d2 time.Time) {
	var a = []rlib.Period{}
	for i := 0; i < len(*sl); i++ {
		// it only make sense when possessionStart or possessionStop falls
		// into the date range
		if (*sl)[i].PossessionStart.Time.Before(d1) && (*sl)[i].PossessionStop.Time.After(d2) {
			continue
		}
		var p = rlib.Period{
			D1: (*sl)[i].PossessionStart.Time,
			D2: (*sl)[i].PossessionStop.Time,
		}
		a = append(a, p)
	}
	b := rlib.FindGaps(&d1, &d2, a)
	for i := 0; i < len(b); i++ {
		rlib.Console("Gap[%d]: %s - %s\n", i, b[i].D1.Format(rlib.RRDATEFMTSQL), b[i].D2.Format(rlib.RRDATEFMTSQL))
	}
	rsa := rlib.RStat(bid, rid, b)
	for i := 0; i < len(rsa); i++ {
		rlib.Console("rsa[%d]: %s - %s, LeaseStatus=%d, UseStatus=%d\n", i, rsa[i].DtStart.Format(rlib.RRDATEFMTSQL), rsa[i].DtStop.Format(rlib.RRDATEFMTSQL), rsa[i].LeaseStatus, rsa[i].UseStatus)
		var r RentRollReportRow
		r.PossessionStart.Scan(rsa[i].DtStart)
		r.PossessionStop.Scan(rsa[i].DtStop)
		r.Description.Scan("Vacant")
		r.Users.Scan(rsa[i].UseStatusStringer())
		r.Payors.Scan(rsa[i].LeaseStatusStringer())
		r.UsePeriod = fmtRRDatePeriod(&rsa[i].DtStart, &rsa[i].DtStop)
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
