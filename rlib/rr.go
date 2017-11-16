package rlib

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

// RentRollStaticInfo is a struct to hold the all static data
// those are received from database per row.
//
// TBD, example/test = washing machine breaks during rental period
// then offset issue ??
type RentRollStaticInfo struct {
	BID             int64
	RID             NullInt64
	RentableName    NullString
	RTID            NullInt64
	RentableType    NullString
	RentCycle       NullInt64
	Status          NullInt64
	Users           NullString
	RARID           NullInt64
	RAID            NullInt64
	AgreementStart  NullDate
	AgreementStop   NullDate
	PossessionStart NullDate
	PossessionStop  NullDate
	RentStart       NullDate
	RentStop        NullDate
	Payors          NullString
	ASMID           NullInt64
	AmountDue       NullFloat64
	PaymentsApplied NullFloat64
	Description     NullString
}

// RentRollVariableData is a struct to hold variable data for a rentroll view
// per rentable
type RentRollVariableData struct {
	RID           int64
	RAID          NullInt64
	Gap           Period
	IncomeOffsets NullFloat64 // should we agreegate this for a rentable?
}

// RentRollStaticInfoRowScan scans a result from sql row and dump it in a RentRollStaticInfo struct
func RentRollStaticInfoRowScan(rows *sql.Rows, q *RentRollStaticInfo) error {
	return rows.Scan(&q.RID, &q.RentableName, &q.RTID, &q.RentableType,
		&q.RentCycle, &q.Status, &q.Users, &q.RARID, &q.RAID,
		&q.AgreementStart, &q.AgreementStop, &q.PossessionStart, &q.PossessionStop,
		&q.RentStart, &q.RentStop, &q.Payors,
		&q.ASMID, &q.AmountDue, &q.PaymentsApplied, &q.Description)
}

// RentRollStaticInfoFieldsMap holds the map of field (alias)
// to actual database field with table reference
// It could refer multiple fields
// It would be helpful in search operation with field values within db from API
var RentRollStaticInfoFieldsMap = SelectQueryFieldMap{
	"RID":             {"Rentable_CUM_RA.RID"},
	"RentableName":    {"Rentable_CUM_RA.RentableName"},
	"RTID":            {"RentableTypes.RTID"},
	"RentableType":    {"RentableTypes.Name"},
	"RentCycle":       {"RentableTypes.RentCycle"},
	"Status":          {"RentableStatus.UseStatus"},
	"Users":           {"User.FirstName", "User.LastName", "User.CompanyName"},
	"RAID":            {"Rentable_CUM_RA.RAID"},
	"AgreementStart":  {"Rentable_CUM_RA.AgreementStart"},
	"AgreementStop":   {"Rentable_CUM_RA.AgreementStop"},
	"PossessionStart": {"Rentable_CUM_RA.PossessionStart"},
	"PossessionStop":  {"Rentable_CUM_RA.PossessionStop"},
	"RentStart":       {"Rentable_CUM_RA.RentStart"},
	"RentStop":        {"Rentable_CUM_RA.RentStop"},
	"Payors":          {"Payor.FirstName", "Payor.LastName", "Payor.CompanyName"},
	"ASMID":           {"Assessments.ASMID"},
	"AmountDue":       {"Assessments.Amount"},
	"Description":     {"AR.Name"},
}

// RentRollStaticInfoFields holds the list of fields need to be fetched
// from database for the RentRollView Query
// Field should be refer by actual db table with (.)
var RentRollStaticInfoFields = SelectQueryFields{
	"Rentable_CUM_RA.RID",
	"Rentable_CUM_RA.RentableName",
	"RentableTypes.RTID",
	"RentableTypes.Name AS RentableType",
	"RentableTypes.RentCycle",
	"RentableStatus.UseStatus AS Status",
	"GROUP_CONCAT(DISTINCT CASE WHEN User.IsCompany > 0 THEN User.CompanyName ELSE CONCAT(User.FirstName, ' ', User.LastName) END ORDER BY User.LastName ASC, User.FirstName ASC, User.CompanyName ASC SEPARATOR ', ' ) AS Users",
	"Rentable_CUM_RA.RARID",
	"Rentable_CUM_RA.RAID",
	"Rentable_CUM_RA.AgreementStart",
	"Rentable_CUM_RA.AgreementStop",
	"Rentable_CUM_RA.PossessionStart",
	"Rentable_CUM_RA.PossessionStop",
	"Rentable_CUM_RA.RentStart",
	"Rentable_CUM_RA.RentStop",
	"GROUP_CONCAT(DISTINCT CASE WHEN Payor.IsCompany > 0 THEN Payor.CompanyName ELSE CONCAT(Payor.FirstName, ' ', Payor.LastName) END ORDER BY Payor.LastName ASC, Payor.FirstName ASC, Payor.CompanyName ASC SEPARATOR ', ') AS Payors",
	"Assessments.ASMID",
	"Assessments.Amount AS AmountDue",
	"(CASE WHEN Assessments.ASMID > 0 THEN SUM(DISTINCT ReceiptAllocation.Amount) ELSE ReceiptAllocation.Amount END) AS PaymentsApplied",
	"(CASE WHEN Assessments.ASMID > 0 THEN ASMAR.Name ELSE RCPTAR.Name END) AS Description",
}

// RentRollStaticInfoQuery gives the static data for rentroll rows
var RentRollStaticInfoQuery = `
SELECT
    {{.SelectClause}}
FROM
    ((SELECT
        RentalAgreement.BID,
            RentalAgreement.RAID,
            RentalAgreement.AgreementStart,
            RentalAgreement.AgreementStop,
            RentalAgreement.PossessionStart,
            RentalAgreement.PossessionStop,
            RentalAgreement.RentStart,
            RentalAgreement.RentStop,
            Rentable.RID,
            Rentable.RentableName,
            RentalAgreementRentables.RARID
    FROM
        Rentable
    LEFT JOIN RentalAgreementRentables ON (RentalAgreementRentables.BID = Rentable.BID
        AND RentalAgreementRentables.RID = Rentable.RID
        AND @DtStart <= RentalAgreementRentables.RARDtStop
        AND @DtStop > RentalAgreementRentables.RARDtStart)
    LEFT JOIN RentalAgreement ON (RentalAgreement.BID = RentalAgreementRentables.BID
        AND RentalAgreement.RAID = RentalAgreementRentables.RAID
        AND @DtStart <= RentalAgreement.AgreementStop
        AND @DtStop > RentalAgreement.AgreementStart)
    WHERE
        Rentable.BID = @BID) UNION (SELECT
        RentalAgreement.BID,
            RentalAgreement.RAID,
            RentalAgreement.AgreementStart,
            RentalAgreement.AgreementStop,
            RentalAgreement.PossessionStart,
            RentalAgreement.PossessionStop,
            RentalAgreement.RentStart,
            RentalAgreement.RentStop,
            NULL AS RID,
            NULL AS RentableName,
            RentalAgreementRentables.RARID
    FROM
        RentalAgreement
    LEFT JOIN RentalAgreementRentables ON (RentalAgreementRentables.BID = RentalAgreement.BID
        AND RentalAgreementRentables.RAID = RentalAgreement.RAID
        AND @DtStart <= RentalAgreementRentables.RARDtStop
        AND @DtStop > RentalAgreementRentables.RARDtStart)
    WHERE
        RentalAgreement.BID = @BID
            AND RentalAgreementRentables.RAID IS NULL
            AND @DtStart <= RentalAgreement.AgreementStop
            AND @DtStop > RentalAgreement.AgreementStart)) AS Rentable_CUM_RA
    LEFT JOIN RentalAgreementPayors ON (Rentable_CUM_RA.BID = RentalAgreementPayors.BID
        AND Rentable_CUM_RA.RAID = RentalAgreementPayors.RAID
        AND @DtStart <= RentalAgreementPayors.DtStop
        AND @DtStop > RentalAgreementPayors.DtStart)
    LEFT JOIN Transactant AS Payor ON (Payor.BID = Rentable_CUM_RA.BID
        AND Payor.TCID = RentalAgreementPayors.TCID)
    LEFT JOIN RentableTypeRef ON (RentableTypeRef.BID = Rentable_CUM_RA.BID
        AND RentableTypeRef.RID = Rentable_CUM_RA.RID)
    LEFT JOIN RentableTypes ON (RentableTypes.BID = RentableTypeRef.BID
        AND RentableTypes.RTID = RentableTypeRef.RTID)
    LEFT JOIN RentableMarketRate ON (RentableMarketRate.BID = RentableTypes.BID
        AND RentableMarketRate.RTID = RentableTypes.RTID
        AND @DtStart <= RentableMarketRate.DtStop
        AND @DtStop > RentableMarketRate.DtStart)
    LEFT JOIN RentableStatus ON (RentableStatus.BID = Rentable_CUM_RA.BID
        AND RentableStatus.RID = Rentable_CUM_RA.RID
        AND @DtStart <= RentableStatus.DtStop
        AND @DtStop > RentableStatus.DtStart)
    LEFT JOIN RentableUsers ON (RentableUsers.BID = Rentable_CUM_RA.BID
        AND RentableUsers.RID = Rentable_CUM_RA.RID
        AND RentableUsers.DtStart >= Rentable_CUM_RA.AgreementStart
        AND RentableUsers.DtStop <= Rentable_CUM_RA.AgreementStop
        AND @DtStart <= RentableUsers.DtStop
        AND @DtStop > RentableUsers.DtStart)
    LEFT JOIN Transactant AS User ON (User.BID = Rentable_CUM_RA.BID
        AND RentableUsers.TCID = User.TCID)
    LEFT JOIN Assessments ON (Assessments.BID = Rentable_CUM_RA.BID
        AND Assessments.RAID = Rentable_CUM_RA.RAID
        AND (Assessments.RentCycle = 0 OR (Assessments.RentCycle > 0 AND Assessments.PASMID != 0))
        AND (CASE WHEN Rentable_CUM_RA.RID > 0 THEN Assessments.RID = Rentable_CUM_RA.RID ELSE 1 END)
        AND (Assessments.FLAGS & 4) = 0
        AND @DtStart <= Assessments.Stop
        AND @DtStop > Assessments.Start)
    LEFT JOIN AR AS ASMAR ON (ASMAR.BID = Assessments.BID
        AND ASMAR.ARID = Assessments.ARID)
    LEFT JOIN Receipt ON (Receipt.BID = Rentable_CUM_RA.BID
        AND Receipt.RAID = Rentable_CUM_RA.RAID
        AND (Receipt.FLAGS & 4) = 0
        AND @DtStart <= Receipt.Dt
        AND Receipt.Dt < @DtStop)
    LEFT JOIN AR AS RCPTAR ON (RCPTAR.BID = Receipt.BID
        AND RCPTAR.ARID = Receipt.ARID)
    LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.BID = Receipt.BID
        AND ReceiptAllocation.RAID = Rentable_CUM_RA.RAID
        AND ReceiptAllocation.RCPTID = Receipt.RCPTID
        AND ReceiptAllocation.ASMID > 0
        AND ReceiptAllocation.ASMID = Assessments.ASMID
        AND @DtStart <= ReceiptAllocation.Dt
        AND ReceiptAllocation.Dt < @DtStop)
GROUP BY {{.GroupClause}}
ORDER BY {{.OrderClause}};`

/*
+------+---------------------------------------------+
| NOTE | Need to take care about search operation    |
|      | As currently, we don't have the whereClause |
|      | (not required) in the viewQuery             |
+------+---------------------------------------------+
*/

// RentRollStaticInfoQueryClause - the query clause for RentRoll View
// helpful when user wants custom sorting, searching within API
var RentRollStaticInfoQueryClause = QueryClause{
	"SelectClause": strings.Join(RentRollStaticInfoFields, ","),
	"WhereClause":  "",
	"GroupClause":  "Rentable_CUM_RA.RID, Rentable_CUM_RA.RAID, Assessments.ASMID",
	"OrderClause":  "-Rentable_CUM_RA.RID DESC, -Rentable_CUM_RA.RAID DESC, Assessments.Amount DESC",
}

// GetRentRollStaticInfoMap returns a map of RID -> all structs that holds static info
// for rentroll report
func GetRentRollStaticInfoMap(BID int64, startDt, stopDt time.Time,
) (map[int64][]RentRollStaticInfo, []int64, error) {

	const funcname = "GetRentRollStaticInfoMap"
	var (
		err           error
		staticInfoMap = make(map[int64][]RentRollStaticInfo)
		xbiz          XBusiness
		d1Str         = startDt.Format(RRDATEFMTSQL)
		d2Str         = stopDt.Format(RRDATEFMTSQL)
		mapKeys       = []int64{}
	)
	Console("Entered in %s\n", funcname)

	// initialize some structures and some required things
	InitBizInternals(BID, &xbiz)

	// get formatted query
	fmtQuery := formatRentRollStaticInfoQuery(BID, startDt, stopDt, "", "", -1, -1)

	// Now, start the database transaction
	tx, err := RRdb.Dbrr.Begin()
	if err != nil {
		return staticInfoMap, mapKeys, err
	}

	// set some mysql variables through `tx`
	if _, err = tx.Exec("SET @BID:=?", BID); err != nil {
		tx.Rollback()
		return staticInfoMap, mapKeys, err
	}
	if _, err = tx.Exec("SET @DtStart:=?", d1Str); err != nil {
		tx.Rollback()
		return staticInfoMap, mapKeys, err
	}
	if _, err = tx.Exec("SET @DtStop:=?", d2Str); err != nil {
		tx.Rollback()
		return staticInfoMap, mapKeys, err
	}

	// Execute query in current transaction for Rentable section
	rrRows, err := tx.Query(fmtQuery)
	if err != nil {
		tx.Rollback()
		return staticInfoMap, mapKeys, err
	}
	defer rrRows.Close()

	// ======================
	// LOOP THROUGH ALL ROWS
	// ======================
	count := 0
	for rrRows.Next() {
		// just assume that it is MainRow, if later encountered that it is child row
		// then "formatRentableChildRow" function would take care of it. :)
		q := RentRollStaticInfo{BID: BID}

		// scan the database row
		if err = RentRollStaticInfoRowScan(rrRows, &q); err != nil {
			return staticInfoMap, mapKeys, err
		}

		// if key found from map, then it is child row, otherwise it is new rentable
		if _, ok := staticInfoMap[q.RID.Int64]; !ok {
			// IT IS *NEW* rentable row
			// store key in the mapKeys list
			mapKeys = append(mapKeys, q.RID.Int64)
		}

		// append new rentable row / formatted child row in map sublist
		staticInfoMap[q.RID.Int64] = append(staticInfoMap[q.RID.Int64], q)

		// update the count only after adding the record
		count++
	}

	// check for any errors from row results
	err = rrRows.Err()
	if err != nil {
		tx.Rollback()
		return staticInfoMap, mapKeys, err
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return staticInfoMap, mapKeys, err
	}
	Console("Added %d rows\n", count)

	return staticInfoMap, mapKeys, nil
}

// formatRentRollStaticInfoQuery returns the formatted query
// with given limit, offset if applicable.
func formatRentRollStaticInfoQuery(BID int64, d1, d2 time.Time,
	additionalWhere, orderBy string, limit, offset int) string {

	const funcname = "formatRentRollStaticInfoQuery"
	var (
		qry   = RentRollStaticInfoQuery
		qc    = GetQueryClauseCopy(RentRollStaticInfoQueryClause)
		where = qc["WhereClause"]
		order = qc["OrderClause"]
	)
	Console("Entered in : %s\n", funcname)

	// if additional conditions are provided then append
	if len(additionalWhere) > 0 {
		where += " AND (" + additionalWhere + ")"
	}
	// override orders of query results if it is given
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
	return RenderSQLQuery(qry, qc)

	// tInit := time.Now()
	// qExec, err := RRdb.Dbrr.Query(dbQry)
	// diff := time.Since(tInit)
	// Console("\nQQQQQQuery Time diff for %s is %s\n\n", rrPart, diff.String())
	// return qExec, err
}
