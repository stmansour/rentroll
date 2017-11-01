package ws

// The RentRoll view/report combines all the complexities of the rentroll
// program into a single user presentation.  Its overall organization is to
// produce a list of all rentables and show details about them and their
// financial performance over a period of time.  Note: Some rentroll reports
// show a snapshot in time rather than info over a period of time.
//
// This RentRoll view/report can span any time range desired by the user:
// a day, 12 days, a month, several months, whatever the user needs.
//
// Use Cases:
// 1. Normal case - a rentable is covered by a single rental agreement for
//    the entire period of the view/report.  The first line will have all the
//    details about the rentable, and show the largest Assessment or Receipt
//    found during the period. Subsequent lines will show other assessments
//    and receipts during the period. Most payments will be made associated
//    with some receipt.  Associated payments and receipts (or receipt
//    allocations) should be shown on the same line.
//
// 2. Some entries will be shown that have no associated rentable. Use cases
//    include Application Fees, Floating Deposits, and perhaps more in the
//    future. These entries are shown after all the Rentables.
//
// 3. A rentable may be part of multiple Rental Agreements during the time
//    period requested. Assessments and Receipts for each rental agreement
//    will be shown chronologically. Within the rows for each Rental Agreement
//    Assessments and Receipts are shown with the largest amounts on the top
//    and the lowest amounts on the bottom.  The time gaps between the end
//    of one RentalAgreement and the beginning of the next may or may not
//    have associate RentableStatus records.  If no rentable status records
//    exist, the report will show a UseStatus of InService and a LeaseStatus
//    of VacantRented if there is no RentalAgreement in the future or
//    VacantNotRented if there is.
//
//-----------------------------------------------------------------------------

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strings"
	"time"
)

// RRSearchRequestData is the struct for request data for rentroll webview
type RRSearchRequestData struct {
	RowsOffset int `json:"rows_offset"` // rentroll report rows offset
}

// RRSearchResponse is the response data for a Rental Agreement Search
type RRSearchResponse struct {
	Status        string                 `json:"status"`
	Total         int64                  `json:"total"`
	Records       []rrpt.RentRollViewRow `json:"records"`
	Summary       []rrpt.RentRollViewRow `json:"summary"`
	TotalMainRows int64                  `json:"total_main_rows"`
}

// SvcRR is the response data for a RR Grid search - The Rent Roll View
func SvcRR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRR"
	var (
		err     error
		g       RRSearchResponse
		reqData RRSearchRequestData
		limit   = d.wsSearchReq.Limit
		startDt = d.wsSearchReq.SearchDtStart
		stopDt  = d.wsSearchReq.SearchDtStop
	)
	if limit == 0 {
		limit = 20
	}
	rlib.Console("Entered %s\n", funcname)

	if err := json.Unmarshal([]byte(d.data), &reqData); err != nil {
		rlib.Console("Error while unmarshalling json from reqData: %s", err.Error())
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	//===========================================================
	// TOTAL RECORDS COUNT
	//===========================================================
	rrViewRowsCount, rrViewMainRowsCount, totalRentables, err := getRRTotal(d.BID, startDt, stopDt)
	if err != nil {
		rlib.Console("Error from getRRTotal routine: %s", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("rrViewRowsCount = %d, rrViewMainRowsCount = %d\n", rrViewRowsCount, rrViewMainRowsCount)

	g.Total = rrViewRowsCount
	g.Total += (totalRentables * 2) // for each rentable, we've subtotal and blank row

	g.TotalMainRows = rrViewMainRowsCount

	// ===========================
	// WhereClauses, OrderClauses
	// ===========================
	rrWhere, rrOrder := GetSearchAndSortSQL(d, rrpt.RentRollViewFieldsMap)

	// NOW GET THE ROWS FOR RENTROLL ROUTINE
	rows, err := rrpt.GetRentRollViewRows(
		d.BID, startDt, stopDt, // BID, startDate, stopDate
		limit, // limit
		rrWhere, rrOrder, reqData.RowsOffset,
	)

	// assign recid and append to g.Records
	recordCounter := int64(d.wsSearchReq.Offset) // count recid from offset
	summaryCounter := int64(0)
	for _, row := range rows {
		if row.IsGrandTotalRow {
			row.Recid = summaryCounter // reset recid for summary
			g.Summary = append(g.Summary, row)
			summaryCounter++
		} else {
			row.Recid = recordCounter
			g.Records = append(g.Records, row)
			recordCounter++
		}
	}
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// getRRTotal returns the total number of rows for rentroll view
func getRRTotal(
	BID int64,
	d1, d2 time.Time,
) (rrViewRowsCount, rrViewMainRowsCount, totalRentables int64, err error) {

	const funcname = "getRRTotal"
	var (
		d1Str = d1.Format(rlib.RRDATEFMTSQL)
		d2Str = d2.Format(rlib.RRDATEFMTSQL)
	)
	rlib.Console("Entered %s\n", funcname)

	// ISSUE: search functionality contains different search scenario, in that case we need to
	// handle case that if field is presents in query FieldMap,.....
	// right now just ignore additional where clause

	// ALL report rows
	rrViewQC := rlib.GetQueryClauseCopy(rrpt.RentRollViewQueryClause)
	rrViewQC["SelectClause"] = "COUNT(*)"
	rrViewCountSubQuery := rlib.RenderSQLQuery(rrpt.RentRollViewQuery, rrViewQC)
	// if rrViewCountSubQuery ends with ';' then remove it
	rrViewCountSubQuery = strings.TrimSuffix(strings.TrimSpace(rrViewCountSubQuery), ";")

	// ALL main rows
	rrViewMainRowsQC := rlib.GetQueryClauseCopy(rrpt.RentRollViewQueryClause)
	rrViewMainRowsQC["SelectClause"] = "Rentable_CUM_RA.RID"
	rrViewMainRowsQC["GroupClause"] = "Rentable_CUM_RA.RID"
	rrViewMainRowsCountSubQuery := rlib.RenderSQLQuery(rrpt.RentRollViewQuery, rrViewMainRowsQC)
	// if rrViewMainRowsCountSubQuery ends with ';' then remove it
	rrViewMainRowsCountSubQuery = strings.TrimSuffix(strings.TrimSpace(rrViewMainRowsCountSubQuery), ";")

	// replace select clause first and get count query
	rrViewRowsTotalQueryForm := `SELECT COUNT(*) FROM ({{.query}}) as T;`

	// Now, start the database transaction
	tx, err := rlib.RRdb.Dbrr.Begin()
	if err != nil {
		return
	}

	// set some mysql variables through `tx`
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

	// Execute query in current transaction for all rows count
	rrViewRowsTotalQuery := rlib.RenderSQLQuery(rrViewRowsTotalQueryForm,
		map[string]string{"query": rrViewCountSubQuery})
	if err = tx.QueryRow(rrViewRowsTotalQuery).Scan(&rrViewRowsCount); err != nil {
		tx.Rollback()
		return
	}

	// Execute query in current transaction for main rows count
	rrViewMainRowsTotalQuery := rlib.RenderSQLQuery(rrViewRowsTotalQueryForm,
		map[string]string{"query": rrViewMainRowsCountSubQuery})
	if err = tx.QueryRow(rrViewMainRowsTotalQuery).Scan(&rrViewMainRowsCount); err != nil {
		tx.Rollback()
		return
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}
	rlib.Console("%s:: rrViewRowsCount = %d\n", funcname, rrViewRowsCount)

	// ----------------
	// TOTAL RENTABLES
	// ----------------
	totalRentablesQuery := fmt.Sprintf("SELECT DISTINCT Rentable.RID FROM Rentable WHERE Rentable.BID=%d", BID)
	totalRentables, err = rlib.GetQueryCount(totalRentablesQuery)
	if err != nil {
		rlib.Console("Error while calculation of totalRentables: %s\n", err.Error())
		return
	}
	rlib.Console("%s:: totalRentables = %d\n", funcname, totalRentables)

	return
}
