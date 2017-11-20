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

/*// RRSearchResponse is the response data for a Rental Agreement Search
type RRSearchResponse struct {
	Status        string                 `json:"status"`
	Total         int64                  `json:"total"`
	Records       []rrpt.RentRollViewRow `json:"records"`
	Summary       []rrpt.RentRollViewRow `json:"summary"`
	TotalMainRows int64                  `json:"total_main_rows"`
}*/

// RRSearchResponse is the response data for a Rental Agreement Search
type RRSearchResponse struct {
	Status        string                    `json:"status"`
	Total         int64                     `json:"total"`
	Records       []rlib.RentRollStaticInfo `json:"records"`
	Summary       []rlib.RentRollStaticInfo `json:"summary"`
	TotalMainRows int64                     `json:"total_main_rows"`
}

// SvcRR is the response data for a RR Grid search - The Rent Roll View
func SvcRR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRR"
	var (
		err     error
		g       RRSearchResponse
		reqData RRSearchRequestData
		limit   = d.wsSearchReq.Limit
	)
	rlib.Console("Entered %s\n", funcname)

	if limit == 0 {
		limit = 20
	}

	if err := json.Unmarshal([]byte(d.data), &reqData); err != nil {
		rlib.Console("Error while unmarshalling json from reqData: %s", err.Error())
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	//#############################################################################
	//
	//   JUST TESTING NEW CODE HERE...  THE NEXT FEW LINES SHOULD GO AWAY SOON
	//
	//#############################################################################

	rlib.Console("\n>>>>>>START OF NEW CODE TESTING\n")
	/*m, n, err1 := rlib.GetRentRollStaticInfoMap(
		d.BID, d.wsSearchReq.SearchDtStart, d.wsSearchReq.SearchDtStop)
	if err1 != nil {
		fmt.Printf("Error in : %s\n", err1.Error())
	}

	err1 = rlib.GetRentRollVariableInfoMap(
		d.BID, d.wsSearchReq.SearchDtStart, d.wsSearchReq.SearchDtStop, &m, &n)
	if err1 != nil {
		fmt.Printf("Error in GetRentRollVariableInfoMap: %s\n", err1.Error())
	}

	grandTTL, totalRows, mainRows, err1 := rlib.GetRentRollGenTotals(
		d.BID, d.wsSearchReq.SearchDtStart, d.wsSearchReq.SearchDtStop, &m, &n)
	if err1 != nil {
		fmt.Printf("Error in GetRentRollGenTotals: %s\n", err1.Error())
	}
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>> totalRows: %d, totalMainRows: %d\n\n\n", totalRows, mainRows)

	// iterate map in sorting order, debugging get easier
	var rids rlib.Int64Range
	for rid := range m {
		rids = append(rids, rid)
	}
	sort.Sort(rids)
	for i := 0; i < len(rids); i++ {
		fmt.Printf("Rentable %d\n", rids[i])
		for _, v := range m[rids[i]] {
			fmt.Printf("\tRID: %2d, SqFt: %7d, RAID: %2d, Use: %s - %s, %s, CycleGSR: %7.2f\n "+
				"PeriodGSR: %7.2f, IncomeOffsets: %7.2f, AmountDue: %7.2f, PaymentsApplied: %7.2f\n"+
				"BeginningRcv: %7.2f, DeltaReceivable: %7.2f, EndReceivable: %7.2f\n"+
				"BeginSecDep: %7.2f, DeltaSecDep: %7.2f, EndSecDep: %7.2f,\n"+
				"FLAGS: %d \n\n",
				v.RID.Int64, v.SqFt.Int64, v.RAID.Int64,
				v.PossessionStart.Time.Format(rlib.RRDATEFMTSQL),
				v.PossessionStop.Time.Format(rlib.RRDATEFMTSQL),
				v.Description.String,
				v.RentCycleGSR,
				v.PeriodGSR,
				v.IncomeOffsets,
				v.AmountDue.Float64,
				v.PaymentsApplied.Float64,
				v.BeginReceivable, v.DeltaReceivable, v.EndReceivable,
				v.BeginSecDep, v.DeltaSecDep, v.EndSecDep,
				v.FLAGS)
		}
	}

	// iterate map in sorting order, debugging get easier
	var raids rlib.Int64Range
	for raid := range n {
		raids = append(raids, raid)
	}
	sort.Sort(raids)
	for i := 0; i < len(raids); i++ {
		fmt.Printf("Rental Agreement: %d\n", raids[i])
		for _, v := range n[raids[i]] {
			fmt.Printf("\tRID: %2d, SqFt: %7d, RAID: %2d, Use: %s - %s, %s, CycleGSR: %7.2f\n "+
				"PeriodGSR: %7.2f, IncomeOffsets: %7.2f, AmountDue: %7.2f, PaymentsApplied: %7.2f\n"+
				"BeginningRcv: %7.2f, DeltaReceivable: %7.2f, EndReceivable: %7.2f\n"+
				"BeginSecDep: %7.2f, DeltaSecDep: %7.2f, EndSecDep: %7.2f,\n"+
				"FLAGS: %d \n\n",
				v.RID.Int64, v.SqFt.Int64, v.RAID.Int64,
				v.PossessionStart.Time.Format(rlib.RRDATEFMTSQL),
				v.PossessionStop.Time.Format(rlib.RRDATEFMTSQL),
				v.Description.String,
				v.RentCycleGSR,
				v.PeriodGSR,
				v.IncomeOffsets,
				v.AmountDue.Float64,
				v.PaymentsApplied.Float64,
				v.BeginReceivable, v.DeltaReceivable, v.EndReceivable,
				v.BeginSecDep, v.DeltaSecDep, v.EndSecDep,
				v.FLAGS)
		}
	}

	fmt.Printf("GrandTotalRow..........\n")
	fmt.Printf("PeriodGSR: %7.2f, IncomeOffsets: %7.2f, AmountDue: %7.2f, PaymentsApplied: %7.2f\n"+
		"BeginningRcv: %7.2f, DeltaReceivable: %7.2f, EndReceivable: %7.2f\n"+
		"BeginSecDep: %7.2f, DeltaSecDep: %7.2f, EndSecDep: %7.2f,\n"+
		"FLAGS: %d \n\n",
		grandTTL.PeriodGSR,
		grandTTL.IncomeOffsets,
		grandTTL.AmountDue.Float64,
		grandTTL.PaymentsApplied.Float64,
		grandTTL.BeginReceivable, grandTTL.DeltaReceivable, grandTTL.EndReceivable,
		grandTTL.BeginSecDep, grandTTL.DeltaSecDep, grandTTL.EndSecDep,
		grandTTL.FLAGS)*/

	tInit := time.Now()
	rRows, count, mainCount, err := rlib.GetRentRollRows(d.BID, d.wsSearchReq.SearchDtStart, d.wsSearchReq.SearchDtStop)
	if err != nil {
		rlib.Console("%s: Error from GetRentRollRows routine: %s", funcname, err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	diff := time.Since(tInit)
	rlib.Console("\nGetRentRollRows Time diff is %s\n\n", diff.String())

	g.Total = count
	g.TotalMainRows = mainCount

	// assign recid and append to g.Records
	recordCounter := int64(d.wsSearchReq.Offset) // count recid from offset
	summaryCounter := int64(0)
	for _, row := range rRows {
		if row.FLAGS&rlib.RentRollGrandTotalRow > 0 {
			row.Recid = summaryCounter // reset recid for summary
			g.Summary = append(g.Summary, row)
			summaryCounter++
		} else {
			row.Recid = recordCounter
			g.Records = append(g.Records, row)
			recordCounter++
		}
	}
	rlib.Console(">>>>>>>END OF NEW CODE TESTING\n\n")

	//#############################################################################

	/*//===========================================================
	// TOTAL RECORDS COUNT
	//===========================================================
	rrViewRowsCount, rrViewMainRowsCount, err := getRRTotal(d.BID, startDt, stopDt)
	if err != nil {
		rlib.Console("Error from getRRTotal routine: %s", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("rrViewRowsCount = %d, rrViewMainRowsCount = %d\n", rrViewRowsCount, rrViewMainRowsCount)

	g.Total = rrViewRowsCount
	g.Total += (rrViewMainRowsCount * 2) // for each rentroll view row, we've subtotal and blank row

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
	}*/
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// getRRTotal returns the total number of rows for rentroll view
func getRRTotal(
	BID int64,
	d1, d2 time.Time,
) (rrViewRowsCount, rrViewMainRowsCount /*, totalRentables*/ int64, err error) {

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
	rrViewMainRowsQC["SelectClause"] = "DISTINCT Rentable_CUM_RA.RID, Rentable_CUM_RA.RAID"
	rrViewMainRowsQC["GroupClause"] = "Rentable_CUM_RA.RID, Rentable_CUM_RA.RAID"
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

	/*// ----------------
	// TOTAL RENTABLES
	// ----------------
	totalRentablesQuery := fmt.Sprintf("SELECT DISTINCT Rentable.RID FROM Rentable WHERE Rentable.BID=%d", BID)
	totalRentables, err = rlib.GetQueryCount(totalRentablesQuery)
	if err != nil {
		rlib.Console("Error while calculation of totalRentables: %s\n", err.Error())
		return
	}
	rlib.Console("%s:: totalRentables = %d\n", funcname, totalRentables)*/

	return
}
