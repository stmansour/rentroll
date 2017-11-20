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
)

// RRSearchRequestData is the struct for request data for rentroll webview
type RRSearchRequestData struct {
	RowsOffset int `json:"rows_offset"` // rentroll report rows offset
}

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

	/*// TODO: still where, order clause
	// ===========================
	// WhereClauses, OrderClauses
	// ===========================
	rrWhere, rrOrder := GetSearchAndSortSQL(d, rrpt.RentRollViewFieldsMap)*/

	// get rentroll rows
	rRows, count, mainCount, err := rlib.GetRentRollRows(d.BID, d.wsSearchReq.SearchDtStart, d.wsSearchReq.SearchDtStop)
	if err != nil {
		rlib.Console("%s: Error from GetRentRollRows routine: %s", funcname, err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	// assign total to grid struct
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

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}
