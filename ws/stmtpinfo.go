package ws

import (
	"net/http"
	"rentroll/rlib"
)

// PayorStmtInfoGridRecord struct to show record in rentabletype grid
type PayorStmtInfoGridRecord struct {
	Recid          int64 `json:"recid"`
	BID            int64
	TCID           int64
	FirstName      string
	MiddleName     string
	LastName       string
	CompanyName    string
	PayorIsCompany bool
	Address        string
	Payors         string
}

// PayorStmtInfoGetResponse is the response to a GetPayorStmtInfo request
type PayorStmtInfoGetResponse struct {
	Status string                  `json:"status"`
	Record PayorStmtInfoGridRecord `json:"record"`
}

// SvcGetPayorStmInfo returns the requested StatementInfo record
// wsdoc {
//  @Title  Get Payor Statement Info
//	@URL /v1/payorstmtinfo/:BUI/:RAID
//  @Method  POST
//	@Synopsis Get information about a Rental Agreement Statement
//  @Description  Return information about a Rental Agreement Statement
//	@Input WebGridSearchRequest
//  @Response StatementInfoGetResponse
// wsdoc }
func SvcGetPayorStmInfo(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcGetPayorStmInfo"
	var (
		g PayorStmtInfoGetResponse
		t rlib.Transactant
	)

	rlib.Console("entered %s\n", funcname)

	err := rlib.GetTransactant(r.Context(), d.ID, &t)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	rlib.MigrateStructVals(&t, &g.Record)
	g.Record.PayorIsCompany = t.IsCompany == true
	g.Record.Address = t.SingleLineAddress()
	rlib.Console("g.Record = %#v\n", g.Record)

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
