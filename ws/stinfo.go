package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// StatementInfoGridRecord struct to show record in rentabletype grid
type StatementInfoGridRecord struct {
	Recid   int64 `json:"recid"`
	RAID    int64
	BID     int64
	Balance float64
	Payors  string
}

// PayorHistory is a struct of data listing RA payors and their time ranges
type PayorHistory struct {
	RAPID       int64
	RAID        int64
	TCID        int64
	DtStart     time.Time
	DtStop      time.Time
	IsCompany   int
	FirstName   string
	LastName    string
	CompanyName string
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

	d1 := time.Now()
	d2 := d1.AddDate(0, 1, 0)

	qry := fmt.Sprintf("SELECT RAPID,RAID,RentalAgreementPayors.TCID,DtStart,DtStop,Transactant.IsCompany,Transactant.FirstName,Transactant.LastName,Transactant.CompanyName FROM RentalAgreementPayors LEFT JOIN Transactant ON RentalAgreementPayors.TCID=Transactant.TCID WHERE RentalAgreementPayors.RAID=%d AND %q<DtStop and %q>=DtStart",
		d.ID, d1.Format(rlib.RRDATEFMTSQL), d2.Format(rlib.RRDATEFMTSQL))

	rlib.Console("query = %s\n", qry)
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	sa := []string{}
	for rows.Next() {
		var p PayorHistory
		err := rows.Scan(&p.RAPID, &p.RAID, &p.TCID, &p.DtStart, &p.DtStop, &p.IsCompany, &p.FirstName, &p.LastName, &p.CompanyName)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		name := ""
		if p.IsCompany > 0 {
			name = p.CompanyName
		} else {
			name = p.FirstName + " " + p.LastName
		}
		sa = append(sa, name)
	}
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	g.Record.Payors = strings.Join(sa, ",")
	g.Record.BID = d.BID
	g.Record.RAID = d.ID

	g.Status = "success"
	SvcWriteResponse(&g, w)
}
