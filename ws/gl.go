package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"sort"
)

// GLAccountTypeDownRecord is struct to list down individual GLAccount record
type GLAccountTypeDownRecord struct {
	LID  int64  `json:"id"`   // Ledger account ID
	Name string `json:"text"` // Ledger account name
}

// GLAccountsTypeDownResponse is the response to list down GLAccounts
type GLAccountsTypeDownResponse struct {
	Status  string                    `json:"status"`
	Total   int64                     `json:"total"`
	Records []GLAccountTypeDownRecord `json:"records"`
}

// SvcGLAccountsList generates a list of all GLAccounts with defined business d.BID
// wsdoc {
//  @Title  GLAccount List
//  @URL /v1/gllist/:BUI
//  @Method  GET
//  @Synopsis Get GLAccounts
//  @Description Get all General Ledger Account's list for a requested business
//  @Input WebGridSearchRequest
//  @Response GLAccountsTypeDownResponse
// wsdoc }
func SvcGLAccountsList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Println("Entered service handler for SvcGLAccountsList")

	var (
		g GLAccountsTypeDownResponse
	)

	// get rentable types for a business
	m := rlib.GetGLAccountMap(d.BID)
	fmt.Printf("GetGLAccountMap returned %d records\n", len(g.Records))

	// sort keys
	var keys rlib.Int64Range
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	// append records in ascending order
	var glAccountList []GLAccountTypeDownRecord
	for _, lid := range keys {
		glAccountList = append(glAccountList, GLAccountTypeDownRecord{LID: m[lid].LID, Name: m[lid].Name})
	}

	g.Records = glAccountList
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
