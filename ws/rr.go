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
	RentableOffset    int `json:"rentableOffset"`
	NoRIDAsmtOffset   int `json:"noRIDAsmtOffset"`
	NoRIDNoAsmtOffset int `json:"noRIDNoAsmtOffset"`
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
	rentablesCount, rentablesAsmtCount, rentablesNoAsmtCount, noRIDAsmtCount, noRIDNoAsmtCount, err :=
		getRRTotal(d.BID, startDt, stopDt)

	if err != nil {
		rlib.Console("Error from getRRTotal routine: %s", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("rentablesCount = %d, rentablesAsmtCount = %d, rentablesNoAsmtCount = %d, noRIDAsmtCount = %d, noRIDNoAsmtCount = %d\n", rentablesCount, rentablesAsmtCount, rentablesNoAsmtCount, noRIDAsmtCount, noRIDNoAsmtCount)
	g.Total = (rentablesCount * 3)                 // for each RENTABLES row we'll add subTotal row and one blank row (another two rows)
	if (rentablesAsmtCount - rentablesCount) > 0 { // in case if any rentables get multiple Assessments
		g.Total += (rentablesAsmtCount - rentablesCount)
	}
	if (rentablesNoAsmtCount - rentablesCount) > 0 { // in case if any rentables get multiple result rows
		g.Total += (rentablesNoAsmtCount - rentablesCount)
	}
	g.Total += noRIDAsmtCount                                              // addition of count of assessments which are associated with any rentables
	g.Total += noRIDNoAsmtCount                                            // addition of count of rows which aren't associated with any asmt/rentables
	g.TotalMainRows = (rentablesCount + noRIDAsmtCount + noRIDNoAsmtCount) // main rows count
	g.Total++                                                              // grand Total row will be added
	g.TotalMainRows++                                                      // grand Total row will be added

	// ===========================
	// WhereClauses, OrderClauses
	// for 3 main different parts
	// ===========================
	rentablesWhereClause, rentablesOrderClause := GetSearchAndSortSQL(d, rrpt.RentablesFieldsMap)
	_, noRIDAsmtOrderClause := GetSearchAndSortSQL(d, rrpt.NoRIDAsmtQueryFieldMap)
	_, noRIDNoAsmtOrderClause := GetSearchAndSortSQL(d, rrpt.NoRIDNoAsmtQueryFieldMap)

	// NOW GET THE ROWS FOR RENTROLL ROUTINE
	rows, err := rrpt.RRReportRows(
		d.BID, startDt, stopDt, // BID, startDate, stopDate
		limit,                                                              // limit
		rentablesWhereClause, rentablesOrderClause, reqData.RentableOffset, // rentables Part
		"", noRIDAsmtOrderClause, reqData.NoRIDAsmtOffset, // "No Rentable Assessment" part
		"", noRIDNoAsmtOrderClause, reqData.NoRIDNoAsmtOffset) // "No Rentable No Assessment" part

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
func getRRTotal(BID int64, d1, d2 time.Time) (
	rentablesCount, rentablesAsmtCount,
	rentablesNoAsmtCount, noRIDAsmtCount,
	noRIDNoAsmtCount int64, err error) {
	const funcname = "getRRTotal"
	rlib.Console("Entered %s\n", funcname)

	// ISSUE: search functionality contains different search scenario, in that case we need to
	// handle case that if field is presents in query FieldMap,.....
	// right now just ignore additional where clause

	// ------------------------
	// Get All Rentables Total
	// ------------------------
	rentablesQC := rlib.GetQueryClauseCopy(rrpt.RentablesQueryClause)
	rentablesQC["WhereClause"] = fmt.Sprintf(rentablesQC["WhereClause"], BID)
	rentablesQC["DtStart"] = d1.Format(rlib.RRDATEFMTSQL)
	rentablesQC["DtStop"] = d2.Format(rlib.RRDATEFMTSQL)

	rentablesCountQuery := rlib.RenderSQLQuery(rrpt.RentablesQuery, rentablesQC)
	rentablesCount, err = rlib.GetQueryCount(rentablesCountQuery, rentablesQC)
	if err != nil {
		rlib.Console("Error from rentablesCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("rentablesCount = %d\n", rentablesCount)

	// ---------------------------------------------------
	// Get All Assessments Total associated with Rentables
	// ---------------------------------------------------
	rentablesAsmtQC := rlib.GetQueryClauseCopy(rrpt.RentablesAsmtQueryClause)
	rentablesAsmtQC["WhereClause"] = fmt.Sprintf(rentablesAsmtQC["WhereClause"], BID)
	rentablesAsmtQC["DtStart"] = rentablesQC["DtStart"]
	rentablesAsmtQC["DtStop"] = rentablesQC["DtStop"]

	rentablesAsmtCountQuery := rlib.RenderSQLQuery(rrpt.RentablesAsmtQuery, rentablesAsmtQC)
	// rlib.Console("rentablesAsmtCountQuery db query = %s\n", rentablesAsmtCountQuery)
	rentablesAsmtCount, err = rlib.GetQueryCount(rentablesAsmtCountQuery, rentablesAsmtQC)
	if err != nil {
		rlib.Console("Error from rentablesAsmtCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("rentablesAsmtCount = %d\n", rentablesAsmtCount)

	// ----------------------------------------------------------------------
	// Get All Payments associated with Rentables but not with any assessment
	// ----------------------------------------------------------------------
	rentablesNoAsmtQC := rlib.GetQueryClauseCopy(rrpt.RentablesNoAsmtQueryClause)
	rentablesNoAsmtQC["WhereClause"] = fmt.Sprintf(rentablesNoAsmtQC["WhereClause"], BID)
	rentablesNoAsmtQC["DtStart"] = rentablesQC["DtStart"]
	rentablesNoAsmtQC["DtStop"] = rentablesQC["DtStop"]

	rentablesNoAsmtCountQuery := rlib.RenderSQLQuery(rrpt.RentablesNoAsmtQuery, rentablesNoAsmtQC)
	// rlib.Console("rentablesNoAsmtCountQuery db query = %s\n", rentablesNoAsmtCountQuery)
	rentablesNoAsmtCount, err = rlib.GetQueryCount(rentablesNoAsmtCountQuery, rentablesNoAsmtQC)
	if err != nil {
		rlib.Console("Error from rentablesNoAsmtCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("rentablesNoAsmtCount = %d\n", rentablesNoAsmtCount)

	// ---------------------------------------------------------------------
	// Get All Assessments Total which are NOT associated with ANY Rentables
	// ---------------------------------------------------------------------
	noRIDAsmtQC := rlib.GetQueryClauseCopy(rrpt.NoRIDAsmtQueryClause)
	noRIDAsmtQC["WhereClause"] = fmt.Sprintf(noRIDAsmtQC["WhereClause"], BID, rentablesQC["DtStart"], rentablesQC["DtStop"])
	noRIDAsmtQC["DtStart"] = rentablesQC["DtStart"]
	noRIDAsmtQC["DtStop"] = rentablesQC["DtStop"]

	noRIDAsmtCountQuery := rlib.RenderSQLQuery(rrpt.NoRIDAsmtQuery, noRIDAsmtQC)
	// rlib.Console("noRIDAsmtCountQuery db query = %s\n", noRIDAsmtCountQuery)
	noRIDAsmtCount, err = rlib.GetQueryCount(noRIDAsmtCountQuery, noRIDAsmtQC)
	if err != nil {
		rlib.Console("Error from noRIDAsmtCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("noRIDAsmtCount = %d\n", noRIDAsmtCount)

	// ------------------------------------------------------------------------
	// Get All Payments which are not associated with any Assessement/Rentables
	// ------------------------------------------------------------------------
	noRIDNoAsmtQC := rlib.GetQueryClauseCopy(rrpt.NoRIDNoAsmtQueryClause)
	noRIDNoAsmtQC["WhereClause"] = fmt.Sprintf(noRIDNoAsmtQC["WhereClause"], BID, rentablesQC["DtStart"], rentablesQC["DtStop"])
	noRIDNoAsmtQC["DtStart"] = rentablesQC["DtStart"]
	noRIDNoAsmtQC["DtStop"] = rentablesQC["DtStop"]

	noRIDNoAsmtCountQuery := rlib.RenderSQLQuery(rrpt.NoRIDNoAsmtQuery, noRIDNoAsmtQC)
	// rlib.Console("noRIDNoAsmtCountQuery db query = %s\n", noRIDNoAsmtCountQuery)
	noRIDNoAsmtCount, err = rlib.GetQueryCount(noRIDNoAsmtCountQuery, noRIDNoAsmtQC)
	if err != nil {
		rlib.Console("Error from noRIDNoAsmtCountQuery: %s\n", err.Error())
		return
	}
	rlib.Console("noRIDNoAsmtCount = %d\n", noRIDNoAsmtCount)

	return
}
