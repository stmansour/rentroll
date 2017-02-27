package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// SearchGLAccountsResponse is the response data to a request for GLAccounts
type SearchGLAccountsResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []rlib.GLAccount `json:"records"`
}

// SvcSearchHandlerGLAccounts generates a report of all GLAccounts for a the business unit
// called out in d.BID
// wsdoc {
//  @Title  Search General Ledger Accounts
//	@URL /v1/accounts/:BUI
//  @Method  GET, POST
//	@Synopsis Return a list of General Ledger Accounts
//  @Description This service returns a list of General Ledger accounts
//	@Input WebGridSearchRequest
//  @Response SearchGLAccountsResponse
// wsdoc }
func SvcSearchHandlerGLAccounts(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcSearchHandlerGLAccounts\n")
	var p rlib.GLAccount
	var err error
	var g SearchGLAccountsResponse

	srch := fmt.Sprintf("BID=%d", d.BID) // default WHERE clause
	order := "GLNumber ASC, Name ASC"    // default ORDER
	q, qw := gridBuildQuery("GLAccount", srch, order, d, &p)

	// set g.Total to the total number of rows of this data...
	g.Total, err = GetRowCount("GLAccount", qw)
	if err != nil {
		fmt.Printf("Error from GetRowCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}

	fmt.Printf("db query = %s\n", q)

	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	i := d.wsSearchReq.Offset
	count := 0
	for rows.Next() {
		var p rlib.GLAccount
		rlib.ReadGLAccounts(rows, &p)
		p.Recid = i
		g.Records = append(g.Records, p)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}

	rlib.Errcheck(rows.Err())
	w.Header().Set("Content-Type", "application/json")
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
