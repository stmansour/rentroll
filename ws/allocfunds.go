package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// UnallocatedReceiptsPayors individual records to be shown on grid
type UnallocatedReceiptsPayors struct {
	Recid int `json:"recid"`
	TCID  int64
	Name  string
	BID   int64
}

// SearchAllocFundsResponse used to send down the list of records on allocated funds grid
type SearchAllocFundsResponse struct {
	Status  string                      `json:"status"`
	Total   int64                       `json:"total"`
	Records []UnallocatedReceiptsPayors `json:"records"`
}

// SvcSearchHandlerAllocFunds generates a list of all payors who has unallocated receipts
// wsdoc {
//  @Title  Search Payors with unallocated receipts
//  @URL /v1/allocfunds/:BUI
//  @Method  GET, POST
//  @Synopsis Return a list of payors with Unallocated receipts
//  @Description This service returns a list of Payors with unallocated receipts
//  @Input WebGridSearchRequest
//  @Response SearchAllocFundsResponse
// wsdoc }
func SvcSearchHandlerAllocFunds(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "SvcSearchHandlerAllocFunds"
		err      error
		g        SearchAllocFundsResponse
	)

	fmt.Printf("Entered %s\n", funcname)

	rows, err := rlib.RRdb.Prepstmt.GetUnallocatedReceipts.Query(d.BID)
	rlib.Errcheck(err)
	defer rows.Close()

	i := d.wsSearchReq.Offset
	count := 0
	for rows.Next() {
		var q UnallocatedReceiptsPayors

		// get receipts record
		var r rlib.Receipt
		rlib.ReadReceipts(rows, &r)

		// get Transactant records
		var t rlib.Transactant
		err := rlib.GetTransactant(r.TCID, &t)
		rlib.Errcheck(err)

		q.Recid = i
		q.TCID = r.TCID
		q.BID = r.BID
		q.Name = t.FirstName + " " + t.LastName

		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}
	rlib.Errcheck(rows.Err())

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}
