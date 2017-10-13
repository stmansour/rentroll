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
	"time"
)

// RRSearchResponse is the response data for a Rental Agreement Search
type RRSearchResponse struct {
	Status        string                   `json:"status"`
	Total         int64                    `json:"total"`
	Records       []rrpt.RentRollReportRow `json:"records"`
	TotalMainRows int64                    `json:"total_main_rows"`
}

// RRRequeestData - struct for request data for parent-child fashioned rentroll report view
type RRRequeestData struct {
	RentableSectionOffset   int `json:"rentableSectionOffset"`
	NoRentableSectionOffset int `json:"noRentableSectionOffset"`
}

// SvcRR is the response data for a RR Grid search - The Rent Roll View
func SvcRR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRR"
	var (
		err     error
		reqData RRRequeestData
		g       RRSearchResponse
		limit   = d.wsSearchReq.Limit
		startDt = d.wsSearchReq.SearchDtStart
		stopDt  = d.wsSearchReq.SearchDtStop
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

	//===========================================================
	// TOTAL RECORDS COUNT
	//===========================================================
	rentableSectionCount, noRentableSectionCount, totalRentables, err := getRRTotal(d.BID, startDt, stopDt)

	if err != nil {
		rlib.Console("Error from getRRTotal routine: %s", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("rentableSectionCount = %d, noRentableSectionCount = %d\n", rentableSectionCount, noRentableSectionCount)

	g.Total = rentableSectionCount
	g.Total += (totalRentables * 2) // for each rentable, we've subtotal and blank row
	g.Total += noRentableSectionCount
	g.Total++ // we'll have grand total row

	g.TotalMainRows = totalRentables + noRentableSectionCount
	g.TotalMainRows++ // we'll have grand total row

	// ===========================
	// WhereClauses, OrderClauses
	// for 2 different parts
	// ===========================
	rentableSectionWhereClause, rentableSectionOrderClause := GetSearchAndSortSQL(d, rrpt.RentableSectionFieldsMap)
	_, noRentableSectionOrderClause := GetSearchAndSortSQL(d, rrpt.NoRentableSectionFieldsMap)

	// NOW GET THE ROWS FOR RENTROLL ROUTINE
	rows, err := rrpt.RRReportRows(
		d.BID, startDt, stopDt, // BID, startDate, stopDate
		limit, // limit
		rentableSectionWhereClause, rentableSectionOrderClause, reqData.RentableSectionOffset, // rentables Part
		"", noRentableSectionOrderClause, reqData.NoRentableSectionOffset, // "No Rentable Assessment" part
	)

	// assign recid and append to g.Records
	rowCounter := int64(0)
	for _, row := range rows {
		row.Recid = rowCounter
		g.Records = append(g.Records, row)
		rowCounter++
	}
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// getRRTotal returns the total of individual total of sections covered by rentroll report
func getRRTotal(
	BID int64,
	d1, d2 time.Time,
) (rentableSectionCount, noRentableSectionCount, totalRentables int64, err error) {

	const funcname = "getRRTotal"
	rlib.Console("Entered %s\n", funcname)

	// ISSUE: search functionality contains different search scenario, in that case we need to
	// handle case that if field is presents in query FieldMap,.....
	// right now just ignore additional where clause

	// ------------------------
	// Get All Rentable Section Total
	// ------------------------
	rentableSectionQC := rlib.GetQueryClauseCopy(rrpt.RentableSectionQueryClause)
	rentableSectionQC["WhereClause"] = fmt.Sprintf(rentableSectionQC["WhereClause"], BID)

	rentableSectionCountQuery := rlib.RenderSQLQuery(rrpt.RentableSectionQuery, rentableSectionQC)
	rentableSectionCount, err = rlib.GetQueryCount(rentableSectionCountQuery, rentableSectionQC)
	if err != nil {
		rlib.Console("Error from rentableSectionCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("%s:: rentableSectionCount = %d\n", funcname, rentableSectionCount)

	// ------------------------------
	// Get NO Rentable Section COUNT
	// ------------------------------
	noRentableSectionQC := rlib.GetQueryClauseCopy(rrpt.NoRentableSectionQueryClause)
	noRentableSectionQC["WhereClause"] = fmt.Sprintf(noRentableSectionQC["WhereClause"], BID)

	noRentableSectionCountQuery := rlib.RenderSQLQuery(rrpt.NoRentableSectionQuery, noRentableSectionQC)
	// rlib.Console("noRentableSectionCountQuery db query = %s\n", noRentableSectionCountQuery)
	noRentableSectionCount, err = rlib.GetQueryCount(noRentableSectionCountQuery, noRentableSectionQC)
	if err != nil {
		rlib.Console("Error from noRentableSectionCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("%s:: noRentableSectionCount = %d\n", funcname, noRentableSectionCount)

	// ------------------------------
	// Main Rows count
	// ------------------------------
	qc := make(rlib.QueryClause)
	totalRentables, err = rlib.GetQueryCount(fmt.Sprintf("SELECT Rentable.RID FROM Rentable WHERE Rentable.BID=%d", BID), qc)
	if err != nil {
		rlib.Console("Error while calculation of totalRentables: %s\n", err.Error())
		return
	}
	rlib.Console("%s:: totalRentables = %d\n", funcname, totalRentables)

	return
}
