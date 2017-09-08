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
	var (
		funcname = "SvcGetPayorStmInfo"
		g        PayorStmtInfoGetResponse
		t        rlib.Transactant
	)

	rlib.Console("entered %s\n", funcname)

	err := rlib.GetTransactant(d.ID, &t)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.MigrateStructVals(&t, &g.Record)
	g.Record.PayorIsCompany = t.IsCompany == 1
	rlib.Console("g.Record = %#v\n", g.Record)

	// d1 := time.Now()
	// d2 := d1.AddDate(0, 1, 0)

	// q := `SELECT RAPID,RentalAgreementPayors.RAID,RentalAgreementPayors.TCID,DtStart,DtStop,
	//      Transactant.IsCompany,Transactant.FirstName,Transactant.LastName,Transactant.CompanyName,
	//      RentalAgreement.AgreementStart, RentalAgreement.AgreementStop,
	//      RentalAgreement.PossessionStart, RentalAgreement.PossessionStop,
	//      RentalAgreement.RentStart, RentalAgreement.RentStop
	//      FROM RentalAgreementPayors
	//      LEFT JOIN Transactant ON RentalAgreementPayors.TCID=Transactant.TCID
	//      LEFT JOIN RentalAgreement On RentalAgreementPayors.RAID=RentalAgreement.RAID
	//      WHERE RentalAgreementPayors.TCID=%d AND %q<DtStop and %q>=DtStart`
	// qry := fmt.Sprintf(q, d.ID, d1.Format(rlib.RRDATEFMTSQL), d2.Format(rlib.RRDATEFMTSQL))

	// rlib.Console("query = %s\n", qry)
	// rows, err := rlib.RRdb.Dbrr.Query(qry)
	// if err != nil {
	// 	SvcGridErrorReturn(w, err, funcname)
	// 	return
	// }
	// defer rows.Close()

	// sa := []string{}
	// var p PayorHistory
	// i := 0
	// for rows.Next() {
	// 	err := rows.Scan(&p.RAPID, &p.RAID, &p.TCID, &p.DtStart, &p.DtStop, &p.IsCompany, &p.FirstName, &p.LastName, &p.CompanyName,
	// 		&p.AgreementStart, &p.AgreementStop, &p.PossessionStart, &p.PossessionStop, &p.RentStart, &p.RentStop)
	// 	i++
	// 	if err != nil {
	// 		SvcGridErrorReturn(w, err, funcname)
	// 		return
	// 	}
	// 	name := ""
	// 	if p.IsCompany > 0 {
	// 		name = p.CompanyName
	// 	} else {
	// 		name = p.FirstName + " " + p.LastName
	// 	}
	// 	sa = append(sa, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	SvcGridErrorReturn(w, err, funcname)
	// 	return
	// }
	// g.Record.BID = d.BID
	// g.Record.RAID = d.ID
	// if i == 0 {
	// 	//------------------------------------------------------------------------
	// 	// If nothing was read, then the date range is not valid.  Load the info
	// 	// for that rental agreement and fill out the dates. As for the payors,
	// 	// get the payors over the lifetime of the agreement
	// 	//------------------------------------------------------------------------
	// 	ra, err := rlib.GetRentalAgreement(d.ID)
	// 	if err != nil {
	// 		SvcGridErrorReturn(w, err, funcname)
	// 		return
	// 	}
	// 	sap := ra.GetPayorNameList(&ra.AgreementStart, &ra.AgreementStop)
	// 	g.Record.Payors = strings.Join(sap, ",")
	// 	g.Record.AgreementStart = rlib.JSONDate(ra.AgreementStart)
	// 	g.Record.AgreementStop = rlib.JSONDate(ra.AgreementStop)
	// 	g.Record.PossessionStart = rlib.JSONDate(ra.PossessionStart)
	// 	g.Record.PossessionStop = rlib.JSONDate(ra.PossessionStop)
	// 	g.Record.RentStart = rlib.JSONDate(ra.RentStart)
	// 	g.Record.RentStop = rlib.JSONDate(ra.RentStop)
	// } else {
	// 	g.Record.Payors = strings.Join(sa, ",")
	// 	g.Record.AgreementStart = rlib.JSONDate(p.AgreementStart)
	// 	g.Record.AgreementStop = rlib.JSONDate(p.AgreementStop)
	// 	g.Record.PossessionStart = rlib.JSONDate(p.PossessionStart)
	// 	g.Record.PossessionStop = rlib.JSONDate(p.PossessionStop)
	// 	g.Record.RentStart = rlib.JSONDate(p.RentStart)
	// 	g.Record.RentStop = rlib.JSONDate(p.RentStop)
	// }
	// now := time.Now()
	// g.Record.Balance, err = rlib.GetRAIDBalance(d.ID, &now)
	// if err != nil {
	// 	SvcGridErrorReturn(w, err, funcname)
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
	// 		SvcGridErrorReturn(w, err, funcname)
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
	SvcWriteResponse(&g, w)
}
