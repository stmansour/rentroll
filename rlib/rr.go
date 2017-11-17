package rlib

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

// NotRentedString is the default string used
// to for the description of an unrented Rentable.
var NotRentedString = string("Unrented")

// This collection of functions implements the raw data-gathering
// needed to produce a RentRoll report or interface.  These routines
// are designed to be used as shown in the pseudo code below:
//
// func myRentrollReportInterface(BID, d1,d2, iftype) {
//
//     m := GetRentRollStaticInfoMap(BID,d1,d2)    // get basic rentable info
//     m = GetRentRollVariableInfoMap(BID,d1,d2,m) // Gaps, IncomeOffsets for entire collection of Rentables
//     m = GetRentRollGenTotals(m)                 // build subtotals and Grand Total
//
//     if iftype = UIView {
//         BuildViewInterface(m, d1,d2)
//     } else if iftype = Report {
//         BuildReport(m, d1,d2)
//     }
// }

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
	RentCycleGSR    float64
	PeriodGSR       float64
	IncomeOffsets   float64
	BeginReceivable float64
	DeltaReceivable float64
	EndReceivable   float64
	BeginSecDep     float64
	DeltaSecDep     float64
	EndSecDep       float64
	FLAGS           uint64
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
	"ASMID":           {"PaymentInfo.ASMID"},
	"AmountDue":       {"PaymentInfo.AmountDue"},
	"Description":     {"PaymentInfo.Description"},
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
	"PaymentInfo.ASMID",
	"PaymentInfo.AmountDue",
	"PaymentInfo.PaymentsApplied",
	"PaymentInfo.Description",
}

// RentRollStaticInfoQuery gives the static data for rentroll rows
//-----------------------------------------------------------------------------
var RentRollStaticInfoQuery = `
SELECT
    {{.SelectClause}}
FROM
    (
        (
        /*
         *  Collect All Rentables no matter whether they got any rental agreement
         *  or not.
         */
        SELECT
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
        FROM Rentable
            LEFT JOIN RentalAgreementRentables ON (RentalAgreementRentables.BID = Rentable.BID
                AND RentalAgreementRentables.RID = Rentable.RID
                AND @DtStart <= RentalAgreementRentables.RARDtStop
                AND @DtStop > RentalAgreementRentables.RARDtStart)
            LEFT JOIN RentalAgreement ON (RentalAgreement.BID = RentalAgreementRentables.BID
                AND RentalAgreement.RAID = RentalAgreementRentables.RAID
                AND @DtStart <= RentalAgreement.AgreementStop
                AND @DtStop > RentalAgreement.AgreementStart)
        WHERE
            Rentable.BID = @BID
        )
        UNION
        (
        /*
         *  Collect All Rental Agreements which aren't associated with any
         *  rentables.
         */
        SELECT
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
        FROM RentalAgreement
            LEFT JOIN RentalAgreementRentables ON (RentalAgreementRentables.BID = RentalAgreement.BID
                AND RentalAgreementRentables.RAID = RentalAgreement.RAID
                AND @DtStart <= RentalAgreementRentables.RARDtStop
                AND @DtStop > RentalAgreementRentables.RARDtStart
            )
        WHERE RentalAgreement.BID = @BID
            AND RentalAgreementRentables.RAID IS NULL
            AND @DtStart <= RentalAgreement.AgreementStop
            AND @DtStop > RentalAgreement.AgreementStart
        )
    ) AS Rentable_CUM_RA
        /*
         *  Get Payors info through RentalAgreementPayors and Transactant
         */
        LEFT JOIN RentalAgreementPayors ON (Rentable_CUM_RA.RAID = RentalAgreementPayors.RAID
            AND Rentable_CUM_RA.BID = RentalAgreementPayors.BID
            AND @DtStart <= RentalAgreementPayors.DtStop
            AND @DtStop > RentalAgreementPayors.DtStart)
        LEFT JOIN Transactant AS Payor ON (Payor.TCID = RentalAgreementPayors.TCID
            AND Payor.BID = Rentable_CUM_RA.BID)
        /*
         *  RentableTypes join to get RentableType
         */
        LEFT JOIN RentableTypeRef ON (RentableTypeRef.RID = Rentable_CUM_RA.RID
            AND RentableTypeRef.BID = Rentable_CUM_RA.BID
            AND @DtStart <= RentableTypeRef.DtStop
            AND @DtStop > RentableTypeRef.DtStart
            AND RentableTypeRef.DtStart >= Rentable_CUM_RA.AgreementStart
            AND RentableTypeRef.DtStop <= Rentable_CUM_RA.AgreementStop)
        LEFT JOIN RentableTypes ON (RentableTypes.RTID = RentableTypeRef.RTID
            AND RentableTypes.BID = RentableTypeRef.BID)
        /*
         *  RentableStatus join to get the status
         */
        LEFT JOIN RentableStatus ON (RentableStatus.RID = Rentable_CUM_RA.RID
            AND RentableStatus.BID = Rentable_CUM_RA.BID
            AND @DtStart <= RentableStatus.DtStop
            AND @DtStop > RentableStatus.DtStart
            AND RentableStatus.DtStart >= Rentable_CUM_RA.AgreementStart
            AND RentableStatus.DtStop <= Rentable_CUM_RA.AgreementStop)
        /*
         *  get Users list through RentableUsers with Transactant join
         */
        LEFT JOIN RentableUsers ON (RentableUsers.RID = Rentable_CUM_RA.RID
            AND RentableUsers.RID = Rentable_CUM_RA.RID
            AND @DtStart <= RentableUsers.DtStop
            AND @DtStop > RentableUsers.DtStart
            AND RentableUsers.DtStart >= Rentable_CUM_RA.AgreementStart
            AND RentableUsers.DtStop <= Rentable_CUM_RA.AgreementStop)
        LEFT JOIN Transactant AS User ON (RentableUsers.TCID = User.TCID
            AND User.BID = Rentable_CUM_RA.BID)
        LEFT JOIN (
            /***********************************
            Assessments UNION Receipt Collection
            - - - - - - - - - - - - - - - - - */
            SELECT
                AsmRcptCollection.AmountDue AS AmountDue,
                AsmRcptCollection.ASMID,
                AsmRcptCollection.PaymentsApplied,
                AsmRcptCollection.RCPAID,
                AsmRcptCollection.RAID,
                AsmRcptCollection.RID,
                (CASE
                    WHEN AsmRcptCollection.ASMID > 0 THEN ASMARName
                    ELSE RCPTARName
                END) AS Description
            FROM
                ((
                    /*
                    Collect All Assessments with ReceiptAllocation info
                    which fall in the given report dates.
                    */
                    SELECT
                        Assessments.Amount AS AmountDue,
                        Assessments.ASMID AS ASMID,
                        SUM(DISTINCT ReceiptAllocation.Amount) as PaymentsApplied,
                        GROUP_CONCAT(DISTINCT ReceiptAllocation.RCPAID) AS RCPAID,
                        Assessments.RAID AS RAID,
                        Assessments.RID AS RID,
                        ASMAR.Name AS ASMARName,
                        NULL AS RCPTARName
                    FROM
                        Assessments
                        LEFT JOIN ReceiptAllocation ON (ReceiptAllocation.BID=Assessments.BID
                            AND ReceiptAllocation.RAID = Assessments.RAID
                            AND ReceiptAllocation.ASMID = Assessments.ASMID
                            AND @DtStart <= ReceiptAllocation.Dt
                            AND ReceiptAllocation.Dt < @DtStop)
                        LEFT JOIN Receipt ON (Receipt.BID=ReceiptAllocation.BID
                            -- AND Receipt.RAID = ReceiptAllocation.RAID // Receipt might have not updated with RAID
                            AND Receipt.RCPTID=ReceiptAllocation.RCPTID
                            AND (Receipt.FLAGS & 4) = 0
                            AND @DtStart <= Receipt.Dt
                            AND Receipt.Dt < @DtStop)
                        LEFT JOIN AR AS ASMAR ON (ASMAR.BID = Assessments.BID
                            AND ASMAR.ARID = Assessments.ARID)
                    WHERE Assessments.BID=@BID
                        AND (Assessments.RentCycle = 0 OR (Assessments.RentCycle > 0 AND Assessments.PASMID != 0))
                        AND (Assessments.FLAGS & 4) = 0
                        AND @DtStart <= Assessments.Stop
                        AND @DtStop > Assessments.Start
                ) UNION (
                    /*
                    Collect All Receipt/ReceiptAllocation of which associated assessments
                    those don't fall in the given report dates.
                    */
                    SELECT
                        NULL AS AmountDue,
                        NULL AS ASMID,
                        ReceiptAllocation.Amount AS PaymentsApplied,
                        ReceiptAllocation.RCPAID AS RCPAID,
                        ReceiptAllocation.RAID AS RAID,
                        NULL AS RID,
                        NULL AS ASMARName,
                        RCPTAR.Name AS RCPTARName
                    FROM
                        Receipt
                        INNER JOIN ReceiptAllocation ON (Receipt.BID=ReceiptAllocation.BID
                            -- AND ReceiptAllocation.RAID = Receipt.RAID // Receipt might have not updated with RAID
                            AND Receipt.RCPTID=ReceiptAllocation.RCPTID
                            AND ReceiptAllocation.ASMID > 0)
                        LEFT JOIN Assessments ON (Assessments.BID=ReceiptAllocation.BID
                            AND Assessments.RAID = ReceiptAllocation.RAID
                            AND Assessments.ASMID=ReceiptAllocation.ASMID
                            AND (Assessments.RentCycle = 0 OR (Assessments.RentCycle > 0 AND Assessments.PASMID != 0))
                            AND (Assessments.FLAGS & 4) = 0
                            AND @DtStart <= Assessments.Stop
                            AND @DtStop > Assessments.Start)
                        LEFT JOIN AR AS RCPTAR ON (RCPTAR.BID = Receipt.BID
                            AND RCPTAR.ARID = Receipt.ARID)
                    WHERE Receipt.BID=@BID
                        AND Assessments.ASMID IS NULL
                        AND (Receipt.FLAGS & 4) = 0
                        AND @DtStart <= Receipt.Dt
                        AND Receipt.Dt < @DtStop
                )) AS AsmRcptCollection
            -- Avoid any rows in which both Assessment and Receipt parts are Null
            WHERE COALESCE(AsmRcptCollection.ASMID, AsmRcptCollection.PaymentsApplied) IS NOT NULL

            /* - - - - - - - - - - - - - - - - -
            Assessments UNION Receipt Collection
            ************************************/
            ) PaymentInfo ON (PaymentInfo.RAID = Rentable_CUM_RA.RAID
                AND (CASE WHEN PaymentInfo.RID > 0 THEN PaymentInfo.RID=Rentable_CUM_RA.RID ELSE 1 END)
            )
/* GROUP BY RID, RAID, ASMID, RCPAID (In case ASMID=0)*/
GROUP BY {{.GroupClause}}
/* ORDER BY RID (if null then it would be last otherwise), RAID, AmountDue if ASMID >0 else PaymentsApplied */
ORDER BY {{.OrderClause}};
`

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
	"GroupClause":  "Rentable_CUM_RA.RID , Rentable_CUM_RA.RAID , (CASE WHEN PaymentInfo.ASMID > 0 THEN PaymentInfo.ASMID ELSE PaymentInfo.RCPAID END)",
	"OrderClause":  "- Rentable_CUM_RA.RID DESC , - Rentable_CUM_RA.RAID DESC , (CASE WHEN PaymentInfo.ASMID > 0 THEN PaymentInfo.AmountDue ELSE PaymentInfo.PaymentsApplied END) DESC",
}

// GetRentRollStaticInfoMap returns a map of RID -> all structs that holds
// static info for rentroll report.
//
// INPUTS
//	BID      - the business
//  startDt  - report/view start time
//  stopDt   - report/view stop time
//
// RETURNS
//	1:  a map of slices of static info structs.  map key is Rentable ID (RID)
//  2:  a slice of RIDs which are present in the map
//  3:  any error encountered
//-----------------------------------------------------------------------------
func GetRentRollStaticInfoMap(BID int64, startDt, stopDt time.Time) (map[int64][]RentRollStaticInfo, []int64, error) {
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
		if err = RentRollStaticInfoRowScan(rrRows, &q); err != nil { // scan next record
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
		count++
	}

	if err = rrRows.Err(); err != nil {
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
//-----------------------------------------------------------------------------
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

// GetRentRollVariableInfoMap processes static info map, produces an updated
//      map. It updates the map with vacancy information for each component
//      as necessary.
//
// INPUTS
//	BID      - the business
//  startDt  - report/view start time
//  stopDt   - report/view stop time
//  m        - map created by GetRentRollStaticInfoMap
//
// RETURNS
//	1:  An updated map of slices of RentRollStaticInfo structs.
//  2:  Any error encountered
//-----------------------------------------------------------------------------
func GetRentRollVariableInfoMap(BID int64, startDt, stopDt time.Time,
	m *map[int64][]RentRollStaticInfo) error {

	var xbiz XBusiness
	var err error
	InitBizInternals(BID, &xbiz)

	rentrollMapGapHandler(BID, startDt, stopDt, m)
	err = rentrollMapGSRHandler(BID, startDt, stopDt, m, &xbiz)
	if err != nil {
		return err
	}

	return nil
}

// rentrollMapGapHandler examines the supplied map and adds entries as needed to
//     describe vacancies (periods where the rentable is unrented).
//
// INPUTS
//	BID      - the business
//  startDt  - report/view start time
//  stopDt   - report/view stop time
//  m        - pointer to map created by GetRentRollStaticInfoMap
//
// RETURNS
//  no return value
//-----------------------------------------------------------------------------
func rentrollMapGapHandler(BID int64, startDt, stopDt time.Time,
	m *map[int64][]RentRollStaticInfo) {
	for k, v := range *m {
		var a = []Period{}
		//--------------------------------------
		// look at all the rows for Rentable k
		//--------------------------------------
		for i := 0; i < len(v); i++ {
			var p = Period{
				D1: v[i].PossessionStart.Time,
				D2: v[i].PossessionStop.Time,
			}
			a = append(a, p)
		}
		b := FindGaps(&startDt, &stopDt, a) // look for gaps
		if len(b) == 0 {                    // did we find any?
			continue // NO: move on to the next Rentable
		}
		//--------------------------------------------------------------------
		// Found some gaps, create a slice of RentRollVariableData structs,
		// and add it to the map.
		//--------------------------------------------------------------------
		for i := 0; i < len(b); i++ {
			//----------------------------------------------------------------
			// If the gap start and end time match the report range start and
			// end time then the Rentable is unrented for the entire period.
			// So, we will use the existing row rather than adding a new row.
			//----------------------------------------------------------------
			if b[i].D1.Equal(startDt) && b[i].D2.Equal(stopDt) {
				(*m)[k][0].RID.Scan(k)
				(*m)[k][0].PossessionStart.Scan(b[i].D1) // vacancy ranges is shown in "use" column
				(*m)[k][0].PossessionStop.Scan(b[i].D2)
				(*m)[k][0].Description.Scan(NotRentedString)
				continue
			}
			var g RentRollStaticInfo
			g.BID = BID
			g.RID.Scan(k)
			g.PossessionStart.Scan(b[i].D1) // vacancy ranges is shown in "use" column
			g.PossessionStop.Scan(b[i].D2)
			g.Description.Scan(NotRentedString)
			(*m)[k] = append((*m)[k], g)
		}
	}
}

// rentrollMapGSRHandler examines the supplied map and adds GSR information.
//
// INPUTS
//	BID      - the business
//  startDt  - report/view start time
//  stopDt   - report/view stop time
//  m        - pointer to map created by GetRentRollStaticInfoMap
//  xbiz     - XBusiness for getting info about RentableType and more
//
// RETURNS
//  any error encountered or nil if no error occurred
//-----------------------------------------------------------------------------
func rentrollMapGSRHandler(BID int64, startDt, stopDt time.Time,
	m *map[int64][]RentRollStaticInfo, xbiz *XBusiness) error {
	var err error
	for k, v := range *m { // for every component
		var gsrAmt float64
		gsrAmt, _, _, err = CalculateLoadedGSR(BID, k, &startDt, &stopDt, xbiz)
		if err != nil {
			return err
		}
		gsr := GetRentableMarketRate(xbiz, k, &startDt, &stopDt)
		v[0].RentCycleGSR = gsr
		v[0].PeriodGSR = gsrAmt
	}
	return nil
}

// GetRentRollGenTotals generates the subtotal rows and grand total rows
//      of a RentRoll datastructure.
//
// INPUTS
//	BID      - the business
//  startDt  - report/view start time
//  stopDt   - report/view stop time
//  m        - pointer to map created by GetRentRollStaticInfoMap
//  xbiz     - XBusiness for getting info about RentableType and more
//
// RETURNS
//  any error encountered or nil if no error occurred
//-----------------------------------------------------------------------------
func GetRentRollGenTotals(BID int64, startDt, stopDt time.Time,
	m *map[int64][]RentRollStaticInfo) error {
	return nil
}
