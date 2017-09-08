package ws

//

import (
	"database/sql"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// StatementPayor is a structure to fill the statement detail grid
type StatementPayor struct {
	Recid       int64         `json:"recid"` // this is to support the w2ui form
	RAID        int64         // internal unique id
	BID         int64         // Business (so that we can process by Business)
	BUD         rlib.XJSONBud // which business
	FirstName   string
	MiddleName  string
	LastName    string
	CompanyName string
	FLAGS       uint64 // Rcpt / Asmt flags
}

// StmtPayorResponse is the response data for a Rental Agreement Search
type StmtPayorResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []StatementPayor `json:"records"`
}

// fields list needs to be fetched for grid
var payorGridFieldsMap = map[string][]string{
	"TCID":          {"Transactant.TCID"},
	"FirstName":     {"Transactant.FirstName"},
	"MiddleName":    {"Transactant.MiddleName"},
	"LastName":      {"Transactant.LastName"},
	"PreferredName": {"Transactant.PreferredName"},
	"CompanyName":   {"Transactant.CompanyName"},
	"IsCompany":     {"Transactant.IsCompany"},
}

var payorSelectFields = []string{
	"Transactant.TCID",
	"Transactant.BID",
	"Transactant.FirstName",
	"Transactant.MiddleName",
	"Transactant.LastName",
	"Transactant.PreferredName",
	"Transactant.CompanyName",
	"Transactant.IsCompany",
}

// SvcPayorStmtDispatch formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the TCID as follows:
//       0    1       2    3
// 		/v1/xperson/BID/TCID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcPayorStmtDispatch(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcPayorStmtDispatch"
	var err error
	rlib.Console("Entered %s\n", funcname)
	if len(d.pathElements) > 3 {
		if d.TCID, err = SvcExtractIDFromURI(r.RequestURI, "TCID", 3, w); err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
	}
	rlib.Console("Request: %s:  BID = %d,  TCID = %d\n", d.wsSearchReq.Cmd, d.BID, d.TCID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID == -1 {
			SvcStatementPayor(w, r, d)
			return
		}
		getPayorStmt(w, r, d)
	case "save":
		savePayorStmt(w, r, d)
	case "delete":
		deletePayorStmt(w, r, d)
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
	}
}

// payorRowScan scans a result from sql row and dump it in a rlib.Transactant struct
func payorRowScan(rows *sql.Rows, t rlib.Transactant) (rlib.Transactant, error) {
	err := rows.Scan(&t.TCID, &t.BID, &t.FirstName, &t.MiddleName, &t.LastName, &t.PreferredName, &t.CompanyName, &t.IsCompany)
	return t, err
}

// SvcStatementPayor is the response data for a Stmt Grid search
// wsdoc {
//  @Title  Statement Detail
//	@URL /v1/payorstmt/:BUI/:TCID
//  @Method  POST
//	@Synopsis Returns account details for the supplied TCID in the date range
//  @Description  Returns the assessments and receipts for the time range
//	@Input WebGridSearchRequest
//  @Response StmtPayorResponse
// wsdoc }
func SvcStatementPayor(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcStatementPayors"
		err      error
		g        SearchTransactantsResponse
	)
	rlib.Console("Entered %s\n", funcname)

	const (
		limitClause int = 100
	)

	srch := fmt.Sprintf("Transactant.BID=%d AND %q < RentalAgreementPayors.DtStop AND RentalAgreementPayors.DtStart < %q", d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	order := "Transactant.LastName ASC, Transactant.FirstName ASC, Transactant.CompanyName ASC" // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, payorGridFieldsMap)
	if len(whereClause) > 0 {
		srch += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	// Transactant Query Text Template
	payorsQuery := `
	SELECT
		{{.SelectClause}}
	FROM Transactant
	INNER JOIN RentalAgreementPayors ON RentalAgreementPayors.TCID=Transactant.TCID
	WHERE {{.WhereClause}}
	GROUP BY Transactant.TCID ORDER BY {{.OrderClause}}` // don't add ';', later some parts will be added in query

	// will be substituted as query clauses
	qc := queryClauses{
		"SelectClause": strings.Join(payorSelectFields, ","),
		"WhereClause":  srch,
		"OrderClause":  order,
	}

	// GET TOTAL COUNTS of query
	countQuery := renderSQLQuery(payorsQuery, qc)
	g.Total, err = GetQueryCount(countQuery, qc) // total number of rows that match the criteria
	if err != nil {
		rlib.Console("Error from GetQueryCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("g.Total = %d\n", g.Total)

	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`

	// build query with limit and offset clause
	// if query ends with ';' then remove it
	payorsQueryWithLimit := payorsQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(limitClause)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := renderSQLQuery(payorsQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var t rlib.Transactant
		t.Recid = i

		// get record of payor
		t, err = payorRowScan(rows, t)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		g.Records = append(g.Records, t)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}
	// error check
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	// write response
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

type payorStmtEntry struct {
	Recid           int64 `json:"recid"`
	Date            rlib.JSONDate
	Reverse         bool
	Payor           string
	TCID            int64
	RAID            string
	ASMID           string
	RCPTID          string
	RentableName    string
	Description     string
	UnappliedAmount float64
	AppliedAmount   float64
	Assessment      float64
	Balance         float64
}

// PayorStmtDetailResponse is the response data for a detailed PayorStatement targeted for a grid
type PayorStmtDetailResponse struct {
	Status  string           `json:"status"`
	Total   int64            `json:"total"`
	Records []payorStmtEntry `json:"records"`
}

// getPayorStmt is the response data for a grid Payor Statement
// wsdoc {
//  @Title  Statement Detail
//	@URL /v1/payorstmt/:BUI/:TCID
//  @Method  POST
//	@Synopsis Returns account details for the supplied TCID in the date range
//  @Description  Returns the assessments and receipts for the time range
//	@Input WebGridSearchRequest
//  @Response StmtPayorResponse
// wsdoc }
func getPayorStmt(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "getPayorStmt"
	internal := true
	var psdr PayorStmtDetailResponse
	var xbiz rlib.XBusiness

	// UGH!
	//=======================================================================
	rlib.InitBizInternals(d.BID, &xbiz)
	_, ok := rlib.RRdb.BizTypes[d.BID]
	if !ok {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d]", d.BID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	if len(rlib.RRdb.BizTypes[d.BID].GLAccounts) == 0 {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d].GLAccounts", d.BID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	//=======================================================================
	// UGH!

	payors := []int64{d.ID}
	m, err := rlib.PayorsStatement(d.BID, payors, &d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	payorcache := map[int64]rlib.Transactant{}
	lenmRL := len(m.RL)

	//------------------------------------------------------
	// Generate the Receipt Summary
	//------------------------------------------------------
	// Identify section
	{
		var pe payorStmtEntry
		pe.Description = "*** RECEIPT SUMMARY ***"
		psdr.Records = append(psdr.Records, pe)
	}
	if len(m.RL) == 0 {
		var pe payorStmtEntry
		pe.Description = "No receipts this period"
		psdr.Records = append(psdr.Records, pe)
	} else {
		for i := 0; i < len(m.RL); i++ {
			var pe payorStmtEntry
			if m.RL[i].R.TCID != d.ID {
				continue
			} else {
				pe.Date = rlib.JSONDate(m.RL[i].R.Dt)
				pe.Payor = rlib.GetNameFromTransactantCache(m.RL[i].R.TCID, payorcache)
				pe.RCPTID = rlib.IDtoShortString("RCPT", m.RL[i].R.RCPTID)
				pe.Description = "Receipt " + m.RL[i].R.DocNo
				pe.UnappliedAmount = m.RL[i].Unallocated
				pe.AppliedAmount = m.RL[i].Allocated
				pe.Balance = m.RL[i].R.Amount
			}
			psdr.Records = append(psdr.Records, pe)
		}
	}

	//------------------------------------------------------
	// Unapplied Funds...
	//------------------------------------------------------
	if internal {
		// Identify section
		{
			var pe1 payorStmtEntry
			psdr.Records = append(psdr.Records, pe1)
			var pe payorStmtEntry
			pe.Description = "*** UNAPPLIED FUNDS ***"
			psdr.Records = append(psdr.Records, pe)
		}
		if len(m.RL) == 0 {
			var pe payorStmtEntry
			pe.Description = "No allocations this period"
			psdr.Records = append(psdr.Records, pe)
		} else {
			for i := 0; i < lenmRL; i++ {
				if m.RL[i].R.TCID == d.ID {
					continue
				}
				var pe payorStmtEntry
				pe.Date = rlib.JSONDate(m.RL[i].R.Dt)
				pe.Payor = rlib.GetNameFromTransactantCache(m.RL[i].R.TCID, payorcache)

				//----------------------------------------------------
				// If the payor only has one RAID and it is THIS one
				// then we can list the details of the receipt
				//----------------------------------------------------
				l1 := rlib.GetRentalAgreementsByPayorRange(d.BID, m.RL[i].R.TCID, &d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
				if len(l1) == 1 {
					pe.RAID = rlib.IDtoShortString("RA", l1[0].RAID)
					pe.RCPTID = rlib.IDtoShortString("RCPT", m.RL[i].R.RCPTID)
					pe.Description = "Receipt " + m.RL[i].R.DocNo
					pe.UnappliedAmount = m.RL[i].Unallocated
					pe.AppliedAmount = m.RL[i].Allocated
					pe.Balance = m.RL[i].R.Amount

				} else {
					pe.Description = "TBD"
				}
				psdr.Records = append(psdr.Records, pe)
			}
		}
	}

	//------------------------------------------------------
	// Generate the per-Rental-Agreement information...
	//------------------------------------------------------
	for i := 0; i < len(m.RAB); i++ { // for each RA
		ra, err := rlib.GetRentalAgreement(m.RAB[i].RAID)
		if err != nil {
			rlib.LogAndPrintError("PayorStatement", err)
			continue
		}

		// Identify report section
		{
			var pe1 payorStmtEntry
			psdr.Records = append(psdr.Records, pe1)
			var pe payorStmtEntry
			pe.Description = fmt.Sprintf("*** RENTAL AGREEMENT %d ***", m.RAB[i].RAID)
			psdr.Records = append(psdr.Records, pe)
		}

		// Opening Balance
		{
			var pe payorStmtEntry
			pe.Description = "Opening balance"
			pe.Date = rlib.JSONDate(m.RAB[i].DtStart)
			pe.Balance = m.RAB[i].OpeningBal
			pe.RentableName = ra.GetTheRentableName(&d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
			psdr.Records = append(psdr.Records, pe)
		}

		//------------------------
		// init running totals
		//------------------------
		bal := m.RAB[i].OpeningBal
		asmts := float64(0)
		applied := asmts
		// unapplied := asmts

		for j := 0; j < len(m.RAB[i].Stmt); j++ { // for each line in the statement
			var pe payorStmtEntry
			pe.Date = rlib.JSONDate(m.RAB[i].Stmt[j].Dt)
			pe.RAID = rlib.IDtoShortString("RA", m.RAB[i].RAID)

			if m.RAB[i].Stmt[j].TCID > 0 {
				pe.TCID = m.RAB[i].Stmt[j].TCID
			}

			descr := ""
			if m.RAB[i].Stmt[j].Reverse {
				descr += "REVERSAL: "
			}
			amt := m.RAB[i].Stmt[j].Amt

			switch m.RAB[i].Stmt[j].T {
			case 1: // assessments
				pe.Assessment = amt
				if m.RAB[i].Stmt[j].A.ARID > 0 { // The description will be the name of the Account Rule...
					descr += rlib.RRdb.BizTypes[d.BID].AR[m.RAB[i].Stmt[j].A.ARID].Name
				} else {
					descr += rlib.RRdb.BizTypes[d.BID].GLAccounts[m.RAB[i].Stmt[j].A.ATypeLID].Name
				}
				if m.RAB[i].Stmt[j].RNT.RID > 0 {
					pe.RentableName = m.RAB[i].Stmt[j].RNT.RentableName
				}
				if m.RAB[i].Stmt[j].A.ASMID > 0 {
					pe.ASMID = rlib.IDtoShortString("ASM", m.RAB[i].Stmt[j].A.ASMID)
				}
				if m.RAB[i].Stmt[j].A.RAID > 0 { // Payor(s) = all payors associated with RentalAgreement
					pyrs := rlib.GetRentalAgreementPayorsInRange(m.RAB[i].Stmt[j].A.RAID, &d.wsSearchReq.SearchDtStart, &d.wsSearchReq.SearchDtStop)
					sa := []string{}
					for k := 0; k < len(pyrs); k++ {
						sa = append(sa, rlib.GetNameFromTransactantCache(pyrs[k].TCID, payorcache))
					}
					pe.Payor = strings.Join(sa, ",")
				}
				if !m.RAB[i].Stmt[j].Reverse { // update running totals if not a reversal
					asmts += amt
					bal += amt
				} else {
					descr += " (" + m.RAB[i].Stmt[j].A.Comment + ")"
				}
			case 2: // receipts
				pe.AppliedAmount = amt
				rcptid := m.RAB[i].Stmt[j].R.RCPTID
				pe.RCPTID = rlib.IDtoShortString("RCPT", rcptid)
				descr += "Receipt allocation"
				if rcptid > 0 {
					pe.RCPTID = rlib.IDtoShortString("RCPT", rcptid)
					rcpt := rlib.GetReceipt(rcptid)
					if rcpt.RCPTID > 0 {
						pe.Payor = rlib.GetNameFromTransactantCache(rcpt.TCID, payorcache)
					}
				}
				if m.RAB[i].Stmt[j].A.ASMID > 0 {
					pe.ASMID = rlib.IDtoShortString("ASM", m.RAB[i].Stmt[j].A.ASMID)
				}
				if !m.RAB[i].Stmt[j].Reverse {
					applied += amt
					bal -= amt
				} else {
					rcpt := rlib.GetReceipt(m.RAB[i].Stmt[j].R.RCPTID)
					if rcpt.RCPTID > 0 && len(rcpt.Comment) > 0 {
						descr += " (" + rcpt.Comment + ")"
					}
				}
			}
			pe.Balance = bal
			pe.Description = descr
			psdr.Records = append(psdr.Records, pe)
		}

		var epe = payorStmtEntry{
			Date:          rlib.JSONDate(m.RAB[i].DtStop),
			Description:   "Closing balance",
			AppliedAmount: applied,
			Assessment:    asmts,
			Balance:       m.RAB[i].ClosingBal,
		}
		psdr.Records = append(psdr.Records, epe)
	}
	// write response
	psdr.Status = "success"
	for i := 0; i < len(psdr.Records); i++ {
		psdr.Records[i].Recid = int64(i)
	}
	psdr.Total = int64(len(psdr.Records))
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&psdr, w)
}

func savePayorStmt(w http.ResponseWriter, r *http.Request, d *ServiceData) {

}

func deletePayorStmt(w http.ResponseWriter, r *http.Request, d *ServiceData) {

}
