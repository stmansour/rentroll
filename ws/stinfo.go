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
	Recid           int64 `json:"recid"`
	RAID            int64
	BID             int64
	Balance         float64
	Payors          string
	AgreementStart  rlib.JSONDate
	AgreementStop   rlib.JSONDate
	PossessionStart rlib.JSONDate
	PossessionStop  rlib.JSONDate
	RentStart       rlib.JSONDate
	RentStop        rlib.JSONDate
	PayorUnalloc    string
}

// PayorHistory is a struct of data listing RA payors and their time ranges
type PayorHistory struct {
	RAPID           int64
	RAID            int64
	TCID            int64
	DtStart         time.Time
	DtStop          time.Time
	IsCompany       bool
	FirstName       string
	LastName        string
	CompanyName     string
	AgreementStart  time.Time
	AgreementStop   time.Time
	PossessionStart time.Time
	PossessionStop  time.Time
	RentStart       time.Time
	RentStop        time.Time
}

// StatementInfoGetResponse is the response to a GetStatementInfo request
type StatementInfoGetResponse struct {
	Status string                  `json:"status"`
	Record StatementInfoGridRecord `json:"record"`
}

// SvcGetStatementInfo returns the requested StatementInfo record
// wsdoc {
//  @Title  Get Statement Info
//	@URL /v1/stmtinfo/:BUI/:RAID
//  @Method  POST
//	@Synopsis Get information about a Rental Agreement Statement
//  @Description  Return information about a Rental Agreement Statement
//	@Input WebGridSearchRequest
//  @Response StatementInfoGetResponse
// wsdoc }
func SvcGetStatementInfo(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcGetStatementInfo"
	var (
		g StatementInfoGetResponse
	)

	rlib.Console("entered %s\n", funcname)

	d1 := time.Now()
	d2 := d1.AddDate(0, 1, 0)

	q := `SELECT RAPID,RentalAgreementPayors.RAID,RentalAgreementPayors.TCID,DtStart,DtStop,
      Transactant.IsCompany,Transactant.FirstName,Transactant.LastName,Transactant.CompanyName,
      RentalAgreement.AgreementStart, RentalAgreement.AgreementStop,
      RentalAgreement.PossessionStart, RentalAgreement.PossessionStop,
      RentalAgreement.RentStart, RentalAgreement.RentStop
      FROM RentalAgreementPayors
      LEFT JOIN Transactant ON RentalAgreementPayors.TCID=Transactant.TCID
      LEFT JOIN RentalAgreement On RentalAgreementPayors.RAID=RentalAgreement.RAID
      WHERE RentalAgreementPayors.RAID=%d AND %q<DtStop and %q>=DtStart`
	qry := fmt.Sprintf(q, d.ID, d1.Format(rlib.RRDATEFMTSQL), d2.Format(rlib.RRDATEFMTSQL))

	rlib.Console("query = %s\n", qry)
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	sa := []string{}
	var p PayorHistory
	i := 0
	for rows.Next() {
		err := rows.Scan(&p.RAPID, &p.RAID, &p.TCID, &p.DtStart, &p.DtStop, &p.IsCompany, &p.FirstName, &p.LastName, &p.CompanyName,
			&p.AgreementStart, &p.AgreementStop, &p.PossessionStart, &p.PossessionStop, &p.RentStart, &p.RentStop)
		i++
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		name := ""
		if p.IsCompany {
			name = p.CompanyName
		} else {
			name = p.FirstName + " " + p.LastName
		}
		sa = append(sa, name)
	}
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	g.Record.BID = d.BID
	g.Record.RAID = d.ID
	if i == 0 {
		//------------------------------------------------------------------------
		// If nothing was read, then the date range is not valid.  Load the info
		// for that rental agreement and fill out the dates. As for the payors,
		// get the payors over the lifetime of the agreement
		//------------------------------------------------------------------------
		ra, err := rlib.GetRentalAgreement(r.Context(), d.ID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		sap, _ := ra.GetPayorNameList(r.Context(), &ra.AgreementStart, &ra.AgreementStop)
		g.Record.Payors = strings.Join(sap, ",")
		g.Record.AgreementStart = rlib.JSONDate(ra.AgreementStart)
		g.Record.AgreementStop = rlib.JSONDate(ra.AgreementStop)
		g.Record.PossessionStart = rlib.JSONDate(ra.PossessionStart)
		g.Record.PossessionStop = rlib.JSONDate(ra.PossessionStop)
		g.Record.RentStart = rlib.JSONDate(ra.RentStart)
		g.Record.RentStop = rlib.JSONDate(ra.RentStop)
	} else {
		g.Record.Payors = strings.Join(sa, ",")
		g.Record.AgreementStart = rlib.JSONDate(p.AgreementStart)
		g.Record.AgreementStop = rlib.JSONDate(p.AgreementStop)
		g.Record.PossessionStart = rlib.JSONDate(p.PossessionStart)
		g.Record.PossessionStop = rlib.JSONDate(p.PossessionStop)
		g.Record.RentStart = rlib.JSONDate(p.RentStart)
		g.Record.RentStop = rlib.JSONDate(p.RentStop)
	}

	// now := time.Now()
	// g.Record.Balance, err = rlib.GetRAIDBalance(d.ID, &now)
	// if err != nil {
	// 	SvcErrorReturn(w, err, funcname)
	// 	return
	// }
	// //---------------------
	// // Payor balances
	// //---------------------
	// payors := rlib.GetRentalAgreementPayorsInRange(g.Record.RAID, &d1, &d2)
	// pa := ""
	// for i := 0; i < len(payors); i++ {
	// 	rlist := rlib.GetUnallocatedReceiptsByPayor(d.BID, payors[i].TCID)
	// 	var t rlib.Transactant
	// 	err := rlib.GetTransactant(payors[i].TCID, &t)
	// 	if err != nil {
	// 		SvcErrorReturn(w, err, funcname)
	// 		return
	// 	}
	// 	tot := float64(0)
	// 	for j := 0; j < len(rlist); j++ {
	// 		tot += bizlogic.RemainingReceiptFunds(&rlist[j])
	// 	}
	// 	pa += fmt.Sprintf("%s: $ %s<br>", t.GetUserName(), rlib.RRCommaf(tot))
	// }
	// g.Record.PayorUnalloc = pa

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
