package ws

import (
	"net/http"
	"rentroll/rlib"
)

// StatementInfoGridRecord struct to show record in rentabletype grid
type StatementInfoGridRecord struct {
	Recid   int64 `json:"recid"`
	RAID    int64
	BID     int64
	BUD     rlib.XJSONBud
	Balance float64
}

// StatementInfoGetResponse is the response to a GetStatementInfo request
type StatementInfoGetResponse struct {
	Status string                  `json:"status"`
	Record StatementInfoGridRecord `json:"record"`
}

// SvcGetStatementInfo returns the requested StatementInfo record
// wsdoc {
//  @Title  Get Statement Info
//	@URL /v1/rt/:BUI/:RAID
//  @Method  POST
//	@Synopsis Get information about a Rental Agreement Statement
//  @Description  Return information about a Rental Agreement Statement
//	@Input WebGridSearchRequest
//  @Response StatementInfoGetResponse
// wsdoc }
func SvcGetStatementInfo(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcGetStatementInfo"
		g        StatementInfoGetResponse
	)

	rlib.Console("entered %s\n", funcname)

	// d1 := time.Now()
	// d2 := d1.AddDate(0, 1, 0)

	// m := rlib.GetRentalAgreementPayorsInRange(d.ID, &d1, &d2)
	// var sa []string
	// for i := 0; i < len(m); i++ {
	// 	sa = append(sa, m[i]. )
	// }
	// if err != nil {
	// 	fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
	// 	SvcGridErrorReturn(w, err, funcname)
	// 	return
	// }

	// bal := rlib.GetRAAccountBalance(d.BID, lid, raid, d1)

	// if err != nil {
	// 	SvcGridErrorReturn(w, err, funcname)
	// 	return
	// }

	g.Status = "success"
	SvcWriteResponse(&g, w)
}
