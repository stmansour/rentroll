package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// RentalAgreementTypedown is the struct of data needed for typedown when searching for a RentalAgreement
type RentalAgreementTypedown struct {
	Recid       int64 `json:"recid"`
	TCID        int64
	FirstName   string
	MiddleName  string
	LastName    string
	CompanyName string
	IsCompany   bool
	RAID        int64
}

// RentalAgreementTypedownResponse is the data structure for the response to a search for people
type RentalAgreementTypedownResponse struct {
	Status  string                    `json:"status"`
	Total   int64                     `json:"total"`
	Records []RentalAgreementTypedown `json:"records"`
}

// GetRentalAgreementTypeDown returns the values needed for typedown controls:
// input:   bid - business
//            s - string or substring to search for
//        limit - return no more than this many matches
// return a slice of TransactantTypeDowns and an error.
func GetRentalAgreementTypeDown(bid int64, s string, limit int) ([]RentalAgreementTypedown, error) {
	var m []RentalAgreementTypedown
	s = "%" + s + "%"
	rows, err := rlib.RRdb.Prepstmt.GetRentalAgreementTypeDown.Query(bid, s, s, s, limit)
	if err != nil {
		return m, err
	}
	defer rows.Close()
	for rows.Next() {
		var t RentalAgreementTypedown
		err = rows.Scan(&t.TCID, &t.FirstName, &t.MiddleName, &t.LastName, &t.CompanyName, &t.IsCompany, &t.RAID)
		if err != nil {
			return m, err
		}
		m = append(m, t)
	}
	return m, nil
}

// SvcRentalAgreementTypeDown handles typedown requests for RentalAgreements.  It returns
// the RAID for the associated payor
// wsdoc {
//  @Title  Get Transactants Typedown
//	@URL /v1/ratd/:BUI?request={"search":"The search string","max":"Maximum number of return items"}
//	@Method GET
//	@Synopsis Fast Search for Transactants matching typed characters
//  @Desc Returns TCID, FirstName, Middlename, and LastName of Transactants that
//  @Desc match supplied chars at the beginning of FirstName or LastName
//  @Input WebTypeDownRequest
//  @Response TransactantsTypedownResponse
// wsdoc }
func SvcRentalAgreementTypeDown(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRentalAgreementTypeDown"
	var (
		g   RentalAgreementTypedownResponse
		err error
	)
	rlib.Console("Entered %s\n", funcname)

	rlib.Console("handle typedown: GetRentalAgreementTypeDown( bid=%d, search=%s, limit=%d\n", d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	g.Records, err = GetRentalAgreementTypeDown(d.BID, d.wsTypeDownReq.Search, d.wsTypeDownReq.Max)
	rlib.Console("GetRentalAgreementTypeDown returned %d matches\n", len(g.Records))
	g.Total = int64(len(g.Records))
	if err != nil {
		e := fmt.Errorf("Error getting typedown matches: %s", err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	for i := 0; i < len(g.Records); i++ {
		g.Records[i].Recid = int64(i)
	}
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
