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

// customAttrRTSqft for rentabletypes
const customAttrRTSqft = "Square Feet" // custom attribute for all rentables

// RentRollViewSectionFieldsMap holds the map of field (alias)
// to actual database field with table reference
// It could refer multiple fields
// It would be helpful in search operation with field values within db from API
var RentRollViewFieldsMap = rlib.SelectQueryFieldMap{
	"RID":             {"Rentable_CUM_RA.RID"},
	"RentableName":    {"Rentable_CUM_RA.RentableName"},
	"RTID":            {"RentableTypes.RTID"},
	"RentableType":    {"RentableTypes.Name"},
	"RAID":            {"Rentable_CUM_RA.RAID"},
	"AgreementStart":  {"Rentable_CUM_RA.AgreementStart"},
	"AgreementStop":   {"Rentable_CUM_RA.AgreementStop"},
	"PossessionStart": {"Rentable_CUM_RA.PossessionStart"},
	"PossessionStop":  {"Rentable_CUM_RA.PossessionStop"},
	"RentStart":       {"Rentable_CUM_RA.RentStart"},
	"RentStop":        {"Rentable_CUM_RA.RentStop"},
	"Payors":          {"Payor.FirstName", "Payor.LastName", "Payor.CompanyName"},
	"Users":           {"User.FirstName", "User.LastName", "User.CompanyName"},
	"RentCycle":       {"RentableTypes.RentCycle"},
	"Status":          {"RentableStatus.UseStatus"},
	"GSR":             {"RentableMarketRate.MarketRate"}, // "MarketRate": {"RentableMarketRate.MarketRate"},
	"ASMID":           {"Assessments.ASMID"},
	"AmountDue":       {"Assessments.Amount"},
	"Description":     {"AR.Name"},
}

// RentRollViewFields holds the list of fields need to be fetched
// from database for the RentRollView Query
// Field should be refer by actual db table with (.)
var RentRollViewFields = rlib.SelectQueryFields{
	"Rentable_CUM_RA.RID",
	"Rentable_CUM_RA.RentableName",
	"RentableTypes.RTID",
	"RentableTypes.Name AS RentableType",
	"(CASE WHEN Assessments.ASMID > 0 THEN ASMAR.Name ELSE RCPTAR.Name END) AS Description",
	"GROUP_CONCAT(DISTINCT CASE WHEN User.IsCompany > 0 THEN User.CompanyName ELSE CONCAT(User.FirstName, ' ', User.LastName) END ORDER BY User.LastName ASC, User.FirstName ASC, User.CompanyName ASC SEPARATOR ', ' ) AS Users",
	"GROUP_CONCAT(DISTINCT CASE WHEN Payor.IsCompany > 0 THEN Payor.CompanyName ELSE CONCAT(Payor.FirstName, ' ', Payor.LastName) END ORDER BY Payor.LastName ASC, Payor.FirstName ASC, Payor.CompanyName ASC SEPARATOR ', ') AS Payors",
	"Rentable_CUM_RA.RARID",
	"Rentable_CUM_RA.RAID",
	"Rentable_CUM_RA.AgreementStart",
	"Rentable_CUM_RA.AgreementStop",
	"Rentable_CUM_RA.PossessionStart",
	"Rentable_CUM_RA.PossessionStop",
	"Rentable_CUM_RA.RentStart",
	"Rentable_CUM_RA.RentStop",
	"RentableTypes.RentCycle",
	"RentableStatus.UseStatus AS Status",
	"RentableMarketRate.MarketRate AS GSR",
	"Assessments.ASMID",
	"Assessments.Amount AS AmountDue",
	"(CASE WHEN Assessments.ASMID > 0 THEN SUM(DISTINCT ReceiptAllocation.Amount) ELSE ReceiptAllocation.Amount END) AS PaymentsApplied",
}

// RentRollViewQuery - the actual database raw query
var RentRollViewQuery = `
SELECT
    {{.SelectClause}}
FROM
    (SELECT
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
        RentalAgreement
    RIGHT JOIN RentalAgreementRentables ON (RentalAgreementRentables.RAID = RentalAgreement.RAID
        AND @DtStart <= RentalAgreementRentables.RARDtStop
        AND @DtStop > RentalAgreementRentables.RARDtStart)
    RIGHT JOIN Rentable ON Rentable.RID = RentalAgreementRentables.RID UNION SELECT
        RentalAgreement.BID,
            RentalAgreement.RAID,
            RentalAgreement.AgreementStart,
            RentalAgreement.AgreementStop,
            RentalAgreement.PossessionStart,
            RentalAgreement.PossessionStop,
            RentalAgreement.RentStart,
            RentalAgreement.RentStop,
            NULL,
            NULL,
            RentalAgreementRentables.RARID
    FROM
        RentalAgreement
    LEFT JOIN RentalAgreementRentables ON (RentalAgreementRentables.RAID = RentalAgreement.RAID
        AND @DtStart <= RentalAgreementRentables.RARDtStop
        AND @DtStop > RentalAgreementRentables.RARDtStart)
    WHERE
        RentalAgreementRentables.RAID IS NULL) AS Rentable_CUM_RA
        LEFT JOIN
    RentalAgreementPayors ON (Rentable_CUM_RA.RAID = RentalAgreementPayors.RAID
        AND Rentable_CUM_RA.BID = RentalAgreementPayors.BID
        AND @DtStart <= RentalAgreementPayors.DtStop
        AND @DtStop > RentalAgreementPayors.DtStart)
        LEFT JOIN
    Transactant AS Payor ON (Payor.TCID = RentalAgreementPayors.TCID
        AND Payor.BID = Rentable_CUM_RA.BID)
        LEFT JOIN
    RentableTypeRef ON (RentableTypeRef.RID = Rentable_CUM_RA.RID
        AND RentableTypeRef.BID=Rentable_CUM_RA.BID)
        LEFT JOIN
    RentableTypes ON (RentableTypes.RTID = RentableTypeRef.RTID
        AND RentableTypes.BID = RentableTypeRef.BID)
        LEFT JOIN
    RentableMarketRate ON (RentableMarketRate.RTID = RentableTypes.RTID
        AND RentableMarketRate.BID = RentableTypes.BID
        AND @DtStart <= RentableMarketRate.DtStop
        AND @DtStop > RentableMarketRate.DtStart)
        LEFT JOIN
    RentableStatus ON (RentableStatus.RID = Rentable_CUM_RA.RID
        AND RentableStatus.BID = Rentable_CUM_RA.BID
        AND @DtStart <= RentableStatus.DtStop
        AND @DtStop > RentableStatus.DtStart)
        LEFT JOIN
    RentableUsers ON (RentableUsers.RID = Rentable_CUM_RA.RID
        AND RentableUsers.RID = Rentable_CUM_RA.RID
        AND @DtStart <= RentableUsers.DtStop
        AND @DtStop > RentableUsers.DtStart)
        LEFT JOIN
    Transactant AS User ON (RentableUsers.TCID = User.TCID
        AND User.BID = Rentable_CUM_RA.BID)
        LEFT JOIN
    Assessments ON (Assessments.RAID = Rentable_CUM_RA.RAID
        AND Assessments.BID = Rentable_CUM_RA.BID
        AND (Assessments.FLAGS & 4) = 0
        AND @DtStart <= Assessments.Stop
        AND @DtStop > Assessments.Start
        AND (Assessments.RentCycle = 0
        OR (Assessments.RentCycle > 0
        AND Assessments.PASMID != 0))
        AND Assessments.RID = CASE
        WHEN Rentable_CUM_RA.RID > 0 THEN Rentable_CUM_RA.RID
        ELSE 0
    END)
        LEFT JOIN
    AR as ASMAR ON (ASMAR.ARID=Assessments.ARID
        AND ASMAR.BID = Assessments.BID)
        LEFT JOIN
    Receipt ON (Receipt.RAID = Rentable_CUM_RA.RAID
        AND Receipt.BID = Rentable_CUM_RA.BID
        AND (Receipt.FLAGS & 4) = 0
        AND @DtStart <= Receipt.Dt
        AND Receipt.Dt < @DtStop)
        LEFT JOIN
    AR as RCPTAR ON (RCPTAR.ARID=Receipt.ARID
        AND RCPTAR.BID = Receipt.BID)
        LEFT JOIN
    ReceiptAllocation ON (ReceiptAllocation.RCPTID = Receipt.RCPTID
        AND ReceiptAllocation.BID = Receipt.BID
        AND ReceiptAllocation.RAID = Rentable_CUM_RA.RAID
        AND (CASE WHEN ReceiptAllocation.ASMID > 0 THEN ReceiptAllocation.ASMID = Assessments.ASMID ELSE 1 END)
        AND @DtStart <= ReceiptAllocation.Dt
        AND ReceiptAllocation.Dt < @DtStop)
WHERE {{.WhereClause}}
GROUP BY {{.GroupClause}}
ORDER BY {{.OrderClause}};`

// RentRollViewQueryClause - the query clause for RentRoll View
// helpful when user wants custom sorting, searching within API
var RentRollViewQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(RentRollViewFields, ","),
	"WhereClause":  "Rentable_CUM_RA.BID = @BID AND @DtStart <= Rentable_CUM_RA.AgreementStop AND @DtStop > Rentable_CUM_RA.AgreementStart",
	"GroupClause":  "Rentable_CUM_RA.RID, Rentable_CUM_RA.RAID, (CASE WHEN Assessments.ASMID > 0 THEN Assessments.ASMID ELSE ReceiptAllocation.RCPAID END)",
	"OrderClause":  "-Rentable_CUM_RA.RID DESC, -Rentable_CUM_RA.RAID DESC, (CASE WHEN Assessments.ASMID > 0 THEN Assessments.Amount ELSE ReceiptAllocation.Amount END) DESC",
}

// RentRollViewRow represents the individual row record in rentroll view/report
type RentRollViewRow struct {
	Recid             int64 `json:"recid"` // support for w2ui web service
	BID               int64
	RID               rlib.NullInt64
	RentableName      rlib.NullString
	RTID              rlib.NullInt64
	RentableType      rlib.NullString
	Sqft              rlib.NullInt64
	RARID             rlib.NullInt64 // rental agreement rentable id
	RAID              int64
	RAIDStr           string // string representation of RAID
	AgreementStart    rlib.JSONDate
	AgreementStop     rlib.JSONDate
	AgreementPeriod   string // string presentation of Agreement period
	PossessionStart   rlib.JSONDate
	PossessionStop    rlib.JSONDate
	UsePeriod         string // string presentation of usage period
	RentStart         rlib.JSONDate
	RentStop          rlib.JSONDate
	RentPeriod        string // string presentation of rent period
	Status            rlib.NullInt64
	RentCycle         rlib.NullInt64
	RentCycleStr      string // String representation of Rent Cycle
	Payors            string
	Users             rlib.NullString
	GSR               rlib.NullFloat64
	PeriodGSR         rlib.NullFloat64
	IncomeOffsets     rlib.NullFloat64
	ASMID             rlib.NullInt64
	AmountDue         rlib.NullFloat64
	PaymentsApplied   rlib.NullFloat64
	Description       rlib.NullString  // Account rule Name - referred from Assessments/Receipt
	BeginningRcv      rlib.NullFloat64 // Receivable amount at beginning period
	ChangeInRcv       rlib.NullFloat64 // Change in receivable
	EndingRcv         rlib.NullFloat64 // Ending receivable
	BeginningSecDep   rlib.NullFloat64 // Beginning security deposit
	ChangeInSecDep    rlib.NullFloat64 // Change in security deposit
	EndingSecDep      rlib.NullFloat64 // Ending security deposit
	IsMainRow         bool             // is main row
	IsRentableMainRow bool             // is rentable main row which holds all static data
	IsSubTotalRow     bool             // is subtotal row
	IsBlankRow        bool             // is blank row
	IsRentRollViewRow bool             // is rentroll normal row fetched from database
}

// rentrollViewRowScan scans a result from sql row and dump it in a RentRollViewRow struct
func rentrollViewRowScan(rows *sql.Rows, q *RentRollViewRow) error {
	return rows.Scan(&q.RID, &q.RentableName, &q.RTID, &q.RentableType,
		&q.Description, &q.Users, &q.Payors, &q.RARID, &q.RAID,
		&q.AgreementStart, &q.AgreementStop, &q.PossessionStart, &q.PossessionStop,
		&q.RentStart, &q.RentStop, &q.RentCycle, &q.Status, &q.GSR,
		&q.ASMID, &q.AmountDue, &q.PaymentsApplied)
}

// formatRentRollViewQuery returns the formatted query
// with given limit, offset if applicable.
func formatRentRollViewQuery(
	BID int64, d1, d2 time.Time,
	additionalWhere, orderBy string,
	limit, offset int,
) string {
	const funcname = "formatRentRollViewQuery"

	var (
		qry   = RentRollViewQuery
		qc    = rlib.GetQueryClauseCopy(RentRollViewQueryClause)
		where = qc["WhereClause"]
		order = qc["OrderClause"]
	)
	rlib.Console("Entered in : %s\n", funcname)

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
	return rlib.RenderSQLQuery(qry, qc)

	// tInit := time.Now()
	// qExec, err := rlib.RRdb.Dbrr.Query(dbQry)
	// diff := time.Since(tInit)
	// rlib.Console("\nQQQQQQuery Time diff for %s is %s\n\n", rrPart, diff.String())
	// return qExec, err
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
func setRRDatePeriodString(r *RentRollViewRow, viewRows *[]RentRollViewRow) {
	if len(*viewRows) < 1 {
		return
	}
	/*if len(*viewRows) > 0 {
		return
	}*/

	lastRow := (*viewRows)[len(*viewRows)-1]
	if lastRow.RAID == r.RAID {
		r.AgreementPeriod = ""
		r.RentPeriod = ""
		r.UsePeriod = ""
		r.Payors = ""
		r.Users.String = ""
		r.Users.Valid = false
		r.RAIDStr = ""
	}
}

// formatRentableChildRow formats new Renable Section Row
// into Child Row pattern
func formatRentableChildRow(r *RentRollViewRow) {
	// const funcname = "formatRentableChildRow"

	// set some values to blank
	r.RentableName.String = ""
	r.RentableType.String = ""
	r.Sqft.Int64 = 0
	r.Sqft.Valid = false
	r.IsRentableMainRow = false
	r.IsMainRow = false
	r.GSR.Float64 = 0
	r.GSR.Valid = false
	r.RentCycleStr = ""
}

// GetRentRollViewRows - returns the list of rows for given date range with BID
func GetRentRollViewRows(BID int64,
	startDt, stopDt time.Time,
	pageRowsLimit int,
	whr, odr string, offset int,
) ([]RentRollViewRow, error) {

	const funcname = "GetRentRollViewRows"
	var (
		err                 error
		d1Str               = startDt.Format(rlib.RRDATEFMTSQL)
		d2Str               = stopDt.Format(rlib.RRDATEFMTSQL)
		xbiz                rlib.XBusiness
		rrViewRows          = []RentRollViewRow{}               //
		rentableRowsMap     = make(map[int64][]RentRollViewRow) // per rentable it will hold sublist of rows
		rentableRowsMapKeys = []int64{}
		noRentableRows      = []RentRollViewRow{}
		// grandTTL            = RentRollViewRow{IsMainRow: true} // grand total row
	)
	rlib.Console("Entered in %s\n", funcname)

	// initialize some structures and some required things
	rlib.InitBizInternals(BID, &xbiz)

	// if there is no limit then reset limit, offset value
	if pageRowsLimit <= 0 {
		offset = -1
		pageRowsLimit = -1
	}

	// get formatted query
	fmtQuery := formatRentRollViewQuery(BID, startDt, stopDt, whr, odr, pageRowsLimit, offset)

	// Now, start the database transaction
	tx, err := rlib.RRdb.Dbrr.Begin()
	if err != nil {
		return rrViewRows, err
	}

	// set some mysql variables through `tx`
	if _, err = tx.Exec("SET @BID:=?", BID); err != nil {
		tx.Rollback()
		return rrViewRows, err
	}
	if _, err = tx.Exec("SET @DtStart:=?", d1Str); err != nil {
		tx.Rollback()
		return rrViewRows, err
	}
	if _, err = tx.Exec("SET @DtStop:=?", d2Str); err != nil {
		tx.Rollback()
		return rrViewRows, err
	}

	// Execute query in current transaction for Rentable section
	rrRows, err := tx.Query(fmtQuery)
	if err != nil {
		tx.Rollback()
		return rrViewRows, err
	}
	defer rrRows.Close()

	// ======================
	// LOOP THROUGH ALL ROWS
	// ======================
	count := 0
	for rrRows.Next() {
		// just assume that it is MainRow, if later encountered that it is child row
		// then "formatRentableChildRow" function would take care of it. :)
		q := RentRollViewRow{BID: BID, IsRentRollViewRow: true}

		// scan the database row
		if err = rentrollViewRowScan(rrRows, &q); err != nil {
			return rrViewRows, err
		}

		// format rental agreement
		if q.RAID > 0 {
			raidStr := int64ToStr(q.RAID, true)
			q.RAIDStr = "RA-" + raidStr
		}

		if (time.Time)(q.RentStart).Year() > 1970 {
			q.RentPeriod = fmtRRDatePeriod((*time.Time)(&q.RentStart), (*time.Time)(&q.RentStop))
		}
		if (time.Time)(q.PossessionStart).Year() > 1970 {
			q.UsePeriod = fmtRRDatePeriod((*time.Time)(&q.PossessionStart), (*time.Time)(&q.PossessionStop))
		}
		if (time.Time)(q.AgreementStart).Year() > 1970 {
			q.AgreementPeriod = fmtRRDatePeriod((*time.Time)(&q.AgreementStart), (*time.Time)(&q.AgreementStop))
		}

		// get current row RID
		rowRID := q.RID.Int64

		// only applicable if row has Rentable
		if rowRID > 0 && q.RID.Valid {
			if len(xbiz.RT[q.RTID.Int64].CA) > 0 { // if there are custom attributes
				c, ok := xbiz.RT[q.RTID.Int64].CA[customAttrRTSqft] // see if Square Feet is among them
				if ok {                                             // if it is...
					sqft, err := rlib.IntFromString(c.Value, "invalid customAttrRTSqft attribute")
					q.Sqft.Scan(sqft)
					if err != nil {
						rlib.Console("%s: Error while scanning custom attribute sqft: %s\n", funcname, err.Error())
						// return rrViewRows, err
					}
				}
			}

			// Rent Cycle formatting
			for freqStr, freqNo := range rlib.CycleFreqMap {
				if q.RentCycle.Int64 == freqNo {
					q.RentCycleStr = freqStr
				}
			}

			// if key found from map, then it is child row, otherwise it is new rentable
			if _, ok := rentableRowsMap[rowRID]; !ok {
				// IT IS *NEW* rentable row
				// store key in the mapKeys list
				rentableRowsMapKeys = append(rentableRowsMapKeys, rowRID)
			}

			// append new rentable row / formatted child row in map sublist
			rentableRowsMap[rowRID] = append(rentableRowsMap[rowRID], q)
		} else {
			// append no rentable rows and keep it inside separate slice
			noRentableRows = append(noRentableRows, q)
		}

		// update the rentableSectionCount only after adding the record
		count++
	}

	// check for any errors from row results
	err = rrRows.Err()
	if err != nil {
		tx.Rollback()
		return rrViewRows, err
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return rrViewRows, err
	}
	rlib.Console("Added %d rows\n", count)

	// =========================
	// APPEND ALL RENTABLE ROWS
	// =========================

	// sort the map keys first
	sort.Slice(rentableRowsMapKeys, func(i, j int) bool {
		return rentableRowsMapKeys[i] < rentableRowsMapKeys[j]
	})

	// loop through all rentables with map
	for _, RID := range rentableRowsMapKeys {
		// get the sublist from map
		rentableSubList := rentableRowsMap[RID]

		// first handle rentable gaps
		handleRentableGaps(&xbiz, RID, &rentableSubList, startDt, stopDt)

		// sort the list of all rows per rentable
		sort.Slice(rentableSubList, func(i, j int) bool {
			if (time.Time)(rentableSubList[i].PossessionStart).Equal(
				(time.Time)(rentableSubList[j].PossessionStart)) {
				if rentableSubList[i].AmountDue.Float64 == rentableSubList[j].AmountDue.Float64 {
					return rentableSubList[i].PaymentsApplied.Float64 > rentableSubList[j].PaymentsApplied.Float64 // descending order
				}
				return rentableSubList[i].AmountDue.Float64 > rentableSubList[j].AmountDue.Float64 // descending order
			}
			return (time.Time)(rentableSubList[i].PossessionStart).Before(
				(time.Time)(rentableSubList[j].PossessionStart))
		})

		// add subtotal row and format child rows for this rentable
		addSubTotalRowANDFormatChildRows(BID, &rentableSubList, startDt, stopDt, &rrViewRows)

		// now add blankRow
		rentableSubList = append(rentableSubList, RentRollViewRow{IsBlankRow: true})

		// now add this rentableRowsList to original result row list
		rrViewRows = append(rrViewRows, rentableSubList...)
	}

	// ============================
	// APPEND ALL NO RENTABLE ROWS
	// ============================
	for _, noRentableRow := range noRentableRows {
		noRentableRow.IsMainRow = true
		setRRDatePeriodString(&noRentableRow, &rrViewRows)
		rrViewRows = append(rrViewRows, noRentableRow)
	}

	// ================
	// GRAND TOTAL ROW
	// ================
	if len(rrViewRows) > 0 {
		gt, err := getGrandTotal(BID, startDt, stopDt)
		if err != nil {
			rlib.Console("getGrandTotal: Error = %s\n", err.Error())
		} else {
			rrViewRows = append(rrViewRows, gt)
		}
	}

	return rrViewRows, err
}

// addSubTotalRowANDFormatChildRows - adds subtotal row for sublist belongs to a rentable
// also formats child rows for a rentable
func addSubTotalRowANDFormatChildRows(
	BID int64,
	subRows *[]RentRollViewRow,
	startDt, stopDt time.Time,
	viewRows *[]RentRollViewRow,
) {
	const funcname = "addSubTotalRowANDFormatChildRows"
	rlib.Console("Entered in %s\n", funcname)

	var (
		err         error
		d70         = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
		subTotalRow = RentRollViewRow{BID: BID, IsSubTotalRow: true}
	)

	// mark some flag as valid (true)
	subTotalRow.AmountDue.Valid = true
	subTotalRow.PaymentsApplied.Valid = true
	subTotalRow.PeriodGSR.Valid = true
	subTotalRow.IncomeOffsets.Valid = true
	subTotalRow.BeginningRcv.Valid = true
	subTotalRow.ChangeInRcv.Valid = true
	subTotalRow.EndingRcv.Valid = true
	subTotalRow.BeginningSecDep.Valid = true
	subTotalRow.ChangeInSecDep.Valid = true
	subTotalRow.EndingSecDep.Valid = true

	// loop through all rows belongs to a rentable
	for i, rentableRow := range *subRows {
		subTotalRow.AmountDue.Float64 += rentableRow.AmountDue.Float64
		subTotalRow.PaymentsApplied.Float64 += rentableRow.PaymentsApplied.Float64
		subTotalRow.PeriodGSR.Float64 += rentableRow.PeriodGSR.Float64
		subTotalRow.IncomeOffsets.Float64 += rentableRow.IncomeOffsets.Float64

		if i > 0 {
			// format child row
			formatRentableChildRow(&(*subRows)[i])

			// format RA period dates
			setRRDatePeriodString(&(*subRows)[i], viewRows)
		} else {
			rentableRow.IsMainRow = true
			rentableRow.IsRentableMainRow = true
			(*subRows)[0] = rentableRow
		}
	}

	// Description
	subTotalRow.Description.Scan("Subtotal")

	// BeginningRcv, EndingRcv
	if subTotalRow.RAID == 0 && subTotalRow.RID.Int64 == 0 {
		rlib.Console("GetBeginEndRARBalance: BID=%d, RID=%d, RAID=%d, start/stop = %s - %s\n", subTotalRow.BID, subTotalRow.RID,
			subTotalRow.RAID, startDt.Format(rlib.RRDATEFMTSQL), stopDt.Format(rlib.RRDATEFMTSQL))
	}
	subTotalRow.BeginningRcv.Float64, subTotalRow.EndingRcv.Float64, err =
		rlib.GetBeginEndRARBalance(subTotalRow.BID, subTotalRow.RID.Int64,
			subTotalRow.RAID, &startDt, &stopDt)
	if err != nil {
		rlib.Console("%s: Error while calculating BeginningRcv, EndingRcv:: %s", funcname, err.Error())
	}

	// ChangeInRcv
	subTotalRow.ChangeInRcv.Float64 = (subTotalRow.EndingRcv.Float64 - subTotalRow.BeginningRcv.Float64)

	// BeginningSecDep
	subTotalRow.BeginningSecDep.Float64, err = rlib.GetSecDepBalance(
		subTotalRow.BID, subTotalRow.RAID, subTotalRow.RID.Int64, &d70, &startDt)
	if err != nil {
		rlib.Console("%s: Error while calculating BeginningSecDep:: %s", funcname, err.Error())
	}

	// Change in SecDep
	subTotalRow.ChangeInSecDep.Float64, err = rlib.GetSecDepBalance(
		subTotalRow.BID, subTotalRow.RAID, subTotalRow.RID.Int64, &startDt, &stopDt)
	if err != nil {
		rlib.Console("%s: Error while calculating BeginningSecDep:: %s", funcname, err.Error())
	}

	// EndingSecDep
	subTotalRow.EndingSecDep.Float64 = (subTotalRow.BeginningSecDep.Float64 + subTotalRow.ChangeInSecDep.Float64)

	// append to subRows List
	(*subRows) = append((*subRows), subTotalRow)
}

// handleRentableGaps identifies periods during which the Rentable is not
// covered by a RentalAgreement. It updates the list with entries
// describing the gaps
//----------------------------------------------------------------------
func handleRentableGaps(xbiz *rlib.XBusiness, rid int64, sl *[]RentRollViewRow, d1, d2 time.Time) {
	var a = []rlib.Period{}
	for i := 0; i < len(*sl); i++ {
		var p = rlib.Period{
			D1: (time.Time)((*sl)[i].PossessionStart),
			D2: (time.Time)((*sl)[i].PossessionStop),
		}
		a = append(a, p)
		rlib.Console("SEARCH FOR GAPS: added %s - %s\n", p.D1.Format(rlib.RRDATEFMTSQL), p.D2.Format(rlib.RRDATEFMTSQL))
	}
	b := rlib.FindGaps(&d1, &d2, a)
	for i := 0; i < len(b); i++ {
		rlib.Console("Gap[%d]: %s - %s\n", i, b[i].D1.Format(rlib.RRDATEFMTSQL), b[i].D2.Format(rlib.RRDATEFMTSQL))
	}
	rsa := rlib.RStat(xbiz.P.BID, rid, b)
	for i := 0; i < len(rsa); i++ {
		if rsa[i].RS.DtStart.Before(rlib.TIME0) {
			rsa[i].RS.DtStart = d1
		}
		rlib.Console("rsa[%d]: %s - %s, LeaseStatus=%d, UseStatus=%d\n", i, rsa[i].RS.DtStart.Format(rlib.RRDATEFMTSQL), rsa[i].RS.DtStop.Format(rlib.RRDATEFMTSQL), rsa[i].RS.LeaseStatus, rsa[i].RS.UseStatus)
		var r RentRollViewRow

		// get and feed rentable info
		rentable := rlib.GetRentable(rsa[i].RS.RID)
		r.RID.Scan(rentable.RID)
		r.RentableName.Scan(rentable.RentableName)

		// get rentabletype info and feed it
		m := rlib.GetRentableTypeRefsByRange(r.RID.Int64, &d1, &d2)
		if len(m) > 0 {
			var rt rlib.RentableType
			rlib.GetRentableType(m[0].RTID, &rt)
			r.RTID.Scan(rt.RTID)
			r.RentableType.Scan(rt.Name)
			r.RentCycle.Scan(rt.RentCycle) // ? is it required
			// Rent Cycle formatting
			for freqStr, freqNo := range rlib.CycleFreqMap {
				if r.RentCycle.Int64 == freqNo {
					r.RentCycleStr = freqStr
				}
			}

			// sqft
			if len(rt.CA) > 0 { // if there are custom attributes
				c, ok := rt.CA[customAttrRTSqft] // see if Square Feet is among them
				if ok {                          // if it is...
					sqft, err := rlib.IntFromString(c.Value, "invalid customAttrRTSqft attribute")
					r.Sqft.Scan(sqft)
					if err != nil {
						rlib.Console("%s: Error while scanning sqft: %s\n", err.Error())
					}
				}
			}
			// market Rate GSR
			r.GSR.Scan(rlib.GetRentableMarketRate(xbiz, &rentable, &d1, &d2))
		}

		r.PossessionStart = rlib.JSONDate(rsa[i].RS.DtStart)
		r.PossessionStop = rlib.JSONDate(rsa[i].RS.DtStop)
		r.Description.Scan("Vacant")
		r.Users.Scan(rsa[i].RS.UseStatusStringer())
		r.Payors = rsa[i].RS.LeaseStatusStringer()
		r.UsePeriod = fmtRRDatePeriod(&rsa[i].RS.DtStart, &rsa[i].RS.DtStop)
		r.PeriodGSR.Scan(rlib.VacancyGSR(xbiz, rid, &rsa[i].RS.DtStart, &rsa[i].RS.DtStop))
		r.IncomeOffsets.Scan(r.PeriodGSR.Float64) // TBD: Validate.  For now, just assume that they're equal
		(*sl) = append((*sl), r)
	}
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
	tbl.AddColumn("Amount Due", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)                 // Amount due
	tbl.AddColumn("Payments Applied", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)           // contract rent amounts
	tbl.AddColumn("Beginning Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Change In Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Ending Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)          // account for the associated RentalAgreement
	tbl.AddColumn("Beginning Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Change In Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Ending Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)    // account for the associated RentalAgreement

	// NOW GET THE ROWS FOR RENTROLL ROUTINE
	rows, err := GetRentRollViewRows(
		ri.Bid, ri.D1, ri.D2, // BID, startDate, stopDate
		-1,         // limit
		"", "", -1, // where, order, offset
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

// rrTableAddRow adds row in gotable struct with information
// given by RentRollViewRow struct
func rrTableAddRow(tbl *gotable.Table, q RentRollViewRow) {

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
	tbl.Puts(-1, Payors, q.Payors)
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
		tbl.Puts(-1, GSRRate, float64ToStr(q.GSR.Float64, true))
		tbl.Puts(-1, GSRAmt, float64ToStr(q.PeriodGSR.Float64, true))
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

// ==================================
// RENTROLL GRAND TOTAL CALCULATIONS
// ==================================

// GrandTotalFields holds the list of fields needs to be
// fetched by query
var GrandTotalFields = rlib.SelectQueryFields{
	"Rentable_CUM_RA.RID",
	"Rentable_CUM_RA.RAID",
	"RentableMarketRate.MarketRate as GSR",
	"SUM(DISTINCT Assessments.Amount) as AmountDue",
	"SUM(DISTINCT ReceiptAllocation.Amount) AS PaymentsApplied",
}

// GrandTotalQuery - the query execution plan for to calculate
// grand total for rentable section
var GrandTotalQuery = `
SELECT
    {{.SelectClause}}
FROM
    (SELECT
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
        RentalAgreement
    RIGHT JOIN RentalAgreementRentables ON (RentalAgreementRentables.RAID = RentalAgreement.RAID
        AND @DtStart <= RentalAgreementRentables.RARDtStop
        AND @DtStop > RentalAgreementRentables.RARDtStart)
    RIGHT JOIN Rentable ON Rentable.RID = RentalAgreementRentables.RID UNION SELECT
        RentalAgreement.BID,
            RentalAgreement.RAID,
            RentalAgreement.AgreementStart,
            RentalAgreement.AgreementStop,
            RentalAgreement.PossessionStart,
            RentalAgreement.PossessionStop,
            RentalAgreement.RentStart,
            RentalAgreement.RentStop,
            NULL,
            NULL,
            RentalAgreementRentables.RARID
    FROM
        RentalAgreement
    LEFT JOIN RentalAgreementRentables ON (RentalAgreementRentables.RAID = RentalAgreement.RAID
        AND @DtStart <= RentalAgreementRentables.RARDtStop
        AND @DtStop > RentalAgreementRentables.RARDtStart)
    WHERE
        RentalAgreementRentables.RAID IS NULL) AS Rentable_CUM_RA
        LEFT JOIN
    RentableTypeRef ON (RentableTypeRef.RID = Rentable_CUM_RA.RID
        AND RentableTypeRef.BID=Rentable_CUM_RA.BID)
        LEFT JOIN
    RentableTypes ON (RentableTypes.RTID = RentableTypeRef.RTID
        AND RentableTypes.BID = RentableTypeRef.BID)
        LEFT JOIN
    RentableMarketRate ON (RentableMarketRate.RTID = RentableTypes.RTID
        AND RentableMarketRate.BID = RentableTypes.BID
        AND @DtStart <= RentableMarketRate.DtStop
        AND @DtStop > RentableMarketRate.DtStart)
        LEFT JOIN
    Assessments ON (Assessments.RAID = Rentable_CUM_RA.RAID
        AND Assessments.BID = Rentable_CUM_RA.BID
        AND @DtStart <= Assessments.Stop
        AND @DtStop > Assessments.Start
        AND (Assessments.RentCycle = 0
        OR (Assessments.RentCycle > 0
        AND Assessments.PASMID != 0))
        AND (CASE WHEN Rentable_CUM_RA.RID > 0 THEN Assessments.RID = Rentable_CUM_RA.RID ELSE 1 END)
        AND (Assessments.FLAGS & 4) = 0)
        LEFT JOIN
    Receipt ON (Receipt.RAID = Rentable_CUM_RA.RAID
        AND Receipt.BID = Rentable_CUM_RA.BID
        AND (Receipt.FLAGS & 4) = 0
        AND @DtStart <= Receipt.Dt
        AND Receipt.Dt < @DtStop)
        LEFT JOIN
    ReceiptAllocation ON (ReceiptAllocation.RCPTID = Receipt.RCPTID
        AND ReceiptAllocation.BID = Receipt.BID
        AND ReceiptAllocation.RAID = Rentable_CUM_RA.RAID
        AND (CASE WHEN ReceiptAllocation.ASMID > 0 THEN ReceiptAllocation.ASMID = Assessments.ASMID ELSE 1 END)
        AND @DtStart <= ReceiptAllocation.Dt
        AND ReceiptAllocation.Dt < @DtStop)
WHERE
    {{.WhereClause}}
GROUP BY {{.GroupClause}}
ORDER BY {{.OrderClause}};`

// GrandTotalQueryClause - the query clause for rentables Grand total query
var GrandTotalQueryClause = rlib.QueryClause{
	"SelectClause": strings.Join(GrandTotalFields, ", "),
	"WhereClause":  "Rentable_CUM_RA.BID = @BID AND @DtStart <= Rentable_CUM_RA.AgreementStop AND @DtStop > Rentable_CUM_RA.AgreementStart",
	"GroupClause":  "Rentable_CUM_RA.RID , Rentable_CUM_RA.RAID",
	"OrderClause":  "- Rentable_CUM_RA.RID DESC , - Rentable_CUM_RA.RAID DESC",
}

// getGrandTotal - calculates the grand total for rentroll report
// for all rentables covered by start and stop date range
func getGrandTotal(BID int64, startDt, stopDt time.Time) (grandTTL RentRollViewRow, err error) {
	const funcname = "getGrandTotal"
	var (
		d1Str = startDt.Format(rlib.RRDATEFMTSQL)
		d2Str = stopDt.Format(rlib.RRDATEFMTSQL)
		qc    = rlib.GetQueryClauseCopy(GrandTotalQueryClause)
	)
	rlib.Console("Entered in %s\n", funcname)

	// mark some fields as true for grand total row
	grandTTL.Description.Scan("Grand Total")
	grandTTL.IsMainRow = true
	grandTTL.IncomeOffsets.Valid = true // do we need this?
	grandTTL.PeriodGSR.Valid = true     // do we need this?
	grandTTL.AmountDue.Valid = true
	grandTTL.PaymentsApplied.Valid = true

	// get formatted query for rentable grand total
	fmtQuery := rlib.RenderSQLQuery(GrandTotalQuery, qc)

	// start transaction for rentable Grand Total
	tx, err := rlib.RRdb.Dbrr.Begin()
	if err != nil {
		return
	}

	// NOW, set mysql variables for date values
	if _, err = tx.Exec("SET @BID:=?", BID); err != nil {
		tx.Rollback()
		return
	}
	if _, err = tx.Exec("SET @DtStart:=?", d1Str); err != nil {
		tx.Rollback()
		return
	}
	if _, err = tx.Exec("SET @DtStop:=?", d2Str); err != nil {
		tx.Rollback()
		return
	}

	// Now hit the query
	rows, err := tx.Query(fmtQuery)
	if err != nil {
		tx.Rollback()
		return
	}
	defer rows.Close()

	// iterate through all rows
	for rows.Next() {
		var RID rlib.NullInt64
		var RAID int64
		var MR, AMT, PMT rlib.NullFloat64
		rows.Scan(&RID, &RAID, &MR, &AMT, &PMT)
		rlib.Console("\nRID: %d, RAID: %d, MR: %f, AMT: %f, PMT: %f\n", RID.Int64, RAID, MR.Float64, AMT.Float64, PMT.Float64)

		// ===== some basic flat amount calculations ADDITION
		grandTTL.AmountDue.Float64 += AMT.Float64
		grandTTL.PaymentsApplied.Float64 += PMT.Float64
		// grandTTL.PeriodGSR.Float64 += row.PeriodGSR.Float64
		// grandTTL.IncomeOffsets.Float64 += row.IncomeOffsets.Float64
	}

	// check for any errors from row results
	err = rows.Err()
	if err != nil {
		tx.Rollback()
		return
	}

	// commit rentable Section Transaction, finally
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}

	return
}
