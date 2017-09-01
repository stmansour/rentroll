package rrpt

import (
	"fmt"
	"gotable"
	"os"
	"rentroll/rlib"
	"runtime/debug"
	"strings"
	"time"
)

// RTIDCount is for counting rentables of a particular type
type RTIDCount struct {
	RT    rlib.RentableType // ID of the types we're counting
	Count int64             // the count
}

// StmtEntry describes an entry on a statement
type StmtEntry struct {
	T   int              // 1 = assessment, 2 = Receipt, 3 = Initial Balance
	ID  int64            // ASMID if t==1, RCPTID if t==2, n/a if t==3
	A   *rlib.Assessment // for type==1, the pointer to the assessment
	R   *rlib.Receipt    // for type ==2, the pointer to the receipt
	Amt float64
	Dt  time.Time
}

// GetRentableCountByRentableType returns a structure containing the count of Rentables for each RentableType
// in the specified time range
func GetRentableCountByRentableType(xbiz *rlib.XBusiness, d1, d2 *time.Time) ([]RTIDCount, error) {
	var count int64
	var m []RTIDCount
	var err error
	i := 0
	for _, v := range xbiz.RT {
		s := fmt.Sprintf("SELECT COUNT(*) FROM RentableTypeRef WHERE RTID=%d AND DtStop>\"%s\" AND DtStart<\"%s\"",
			v.RTID, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
		err = rlib.RRdb.Dbrr.QueryRow(s).Scan(&count)
		if err != nil {
			fmt.Printf("GetRentableCountByRentableType: query=\"%s\"    err = %s\n", s, err.Error())
		}
		var rc RTIDCount
		rc.Count = count
		rc.RT = v
		var cerr error
		rc.RT.CA, cerr = rlib.GetAllCustomAttributes(rlib.ELEMRENTABLETYPE, v.RTID)
		if cerr != nil {
			if !rlib.IsSQLNoResultsError(cerr) { // it's not really an error if we don't find any custom attributes
				err = cerr
				break
			}
		}

		m = append(m, rc)
		i++
	}
	return m, err
}

// yes this is a total hack until I can get rid of some old infrastructure, it only applies to old tests though
func temporaryGetAcctsRcv(bid int64) int64 {
	var b rlib.GLAccount
	q := "SELECT " + rlib.RRdb.DBFields["GLAccount"] + " FROM GLAccount WHERE Name LIKE \"%receivable%\""
	row := rlib.RRdb.Dbrr.QueryRow(q)
	rlib.ReadGLAccount(row, &b)
	if b.LID == 0 {
		b = rlib.GetLedgerByGLNo(bid, "12000")
		if b.LID == 0 {
			fmt.Printf("Could not find Accounts Receivable\n")
			debug.PrintStack()
			os.Exit(1) // yes this is terrible.
		}
	}
	return b.LID
}

// GetStatementData returns an array of StatementEntries for building a statement
func GetStatementData(bid int64, raid int64, d1, d2 *time.Time) []StmtEntry {
	var m []StmtEntry
	lid := temporaryGetAcctsRcv(bid)
	bal := rlib.GetRAAccountBalance(bid, lid, raid, d1)
	var initBal = StmtEntry{Amt: bal, T: 3, Dt: *d1}
	m = append(m, initBal)
	n, err := rlib.GetLedgerEntriesForRAID(d1, d2, raid, lid)
	if err != nil {
		return m
	}
	for i := 0; i < len(n); i++ {
		var se StmtEntry
		se.Amt = n[i].Amount
		se.Dt = n[i].Dt
		j := rlib.GetJournal(n[i].JID)
		se.T = int(j.Type)
		se.ID = j.ID
		if se.T == rlib.JOURNALTYPEASMID {
			// read the assessment to find out what it was for...
			a, err := rlib.GetAssessment(se.ID)
			if err != nil {
				rlib.LogAndPrint("rrpt.GetStatementData: error getting asmid %d: %s\n", se.ID, err.Error())
				return m
			}
			se.A = &a
		} else if se.T == rlib.JOURNALTYPERCPTID {
			r := rlib.GetReceipt(se.ID)
			ja := rlib.GetJournalAllocation(n[i].JAID)
			a, err := rlib.GetAssessment(ja.ASMID)
			if err != nil {
				rlib.LogAndPrint("rrpt.GetStatementData: error getting asmid %d: %s\n", ja.ASMID, err.Error())
				return m
			}
			se.A = &a
			se.R = &r
		}

		m = append(m, se)
	}
	return m
}

// RptStatementForRA generates a text Statement for the supplied rental agreement ra.
func RptStatementForRA(ri *ReporterInfo, ra *rlib.RentalAgreement) gotable.Table {
	funcname := "RptStatementForRA"

	// init and prepare some values before table init
	rlib.LoadXRentalAgreement(ra.RAID, ra, &ri.D1, &ri.D2)
	payors := ra.GetPayorNameList(&ri.D1, &ri.D2)

	// table init
	tbl := getRRTable()

	tbl.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 45, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Charge", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Payment", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)

	s := fmt.Sprintf("Statement  -  Rental Agreement %s\nPayor(s): %s\n", ra.IDtoString(), strings.Join(payors, ", "))
	err := TableReportHeaderBlock(&tbl, s, funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	m := GetStatementData(ri.Xbiz.P.BID, ra.RAID, &ri.D1, &ri.D2)
	var b = rlib.RoundToCent(m[0].Amt) // element 0 is always the account balance
	var c = float64(0)                 // credit
	var d = float64(0)                 // debit
	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		descr := ""
		if m[i].T == 1 || m[i].T == 2 {
			if m[i].A.ARID > 0 {
				descr = rlib.RRdb.BizTypes[ri.Xbiz.P.BID].AR[m[i].A.ARID].Name
			} else {
				descr = rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[m[i].A.ATypeLID].Name
			}
		}
		switch m[i].T {
		case 1: // assessments
			amt := rlib.RoundToCent(m[i].Amt)
			c += amt
			b += amt
			tbl.Puts(-1, 1, rlib.IDtoString("ASM", m[i].ID))
			tbl.Puts(-1, 2, descr)
			tbl.Putf(-1, 3, amt)
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].Amt)
			d += amt
			b += amt
			if m[i].A.ASMID > 0 {
				descr = fmt.Sprintf("%s (%s)", descr, m[i].A.IDtoString())
			}
			tbl.Puts(-1, 1, rlib.IDtoString("RCPT", m[i].ID))
			tbl.Puts(-1, 2, descr)
			tbl.Putf(-1, 4, amt)
		case 3: // opening balance
			tbl.Puts(-1, 2, "Opening Balance")
		}
		tbl.Putd(-1, 0, m[i].Dt)
		tbl.Putf(-1, 5, b)
	}
	tbl.AddLineAfter(tbl.RowCount() - 1)
	tbl.AddRow()
	tbl.Putf(-1, 3, c)
	tbl.Putf(-1, 4, d)
	tbl.Putf(-1, 5, c+d+m[0].Amt)

	return tbl
}

// type ReporterInfo struct {
// 	ReportNo              int       // index number of the report
// 	OutputFormat          int       // text, html, maybe more in the future
// 	Bid                   int64     // associated business
// 	Raid                  int64     // associated Rental Agreement if needed
// 	D1                    time.Time // associated date if needed
// 	D2                    time.Time // associated date if needed
// 	NeedsBID              bool      // true if BID is needed for this report
// 	NeedsRAID             bool      // true if RAID is needed for this report
// 	NeedsDt               bool      // true if a Date is needed for this report
// 	RptHeaderD1           bool      // true if the report's header should contain D1
// 	RptHeaderD2           bool      // true if the dates should show as a range D1 - D2
// 	BlankLineAfterRptName bool      // true if a blank line should be added after the Report Name
// 	Handler               func(*ReporterInfo) string
// 	Xbiz                  *rlib.XBusiness // may not be set in all cases
// }

// // StatementReportByRAID returns a gotable with the statement report for the supplied RAID.
// func StatementReportByRAID(ri *ReporterInfo, RAID int64) (gotable.Table, error) {
// 	tbl := getRRTable()

// 	m, err := rlib.GetRAIDStatementInfo(RAID, &ri.D1, &ri.D2)
// 	if err != nil {
// 		return tbl, err
// 	}

// 	//--------------------------------------------
// 	// Set the opening balance.
// 	//--------------------------------------------
// 	var b, c, d float64
// 	var a = StatementDetail{
// 		BID:     sd.BID,
// 		BUD:     rlib.XJSONBud(bud),
// 		RAID:    sd.ID,
// 		Dt:      rlib.JSONDate(m.DtStart),
// 		Descr:   "Opening Balance",
// 		Balance: m.OpeningBal,
// 	}

// 	b = m.OpeningBal
// 	for i := 0; i < len(m.Stmt); i++ {
// 		var a = StatementDetail{
// 			BID:  sd.BID,
// 			BUD:  rlib.XJSONBud(bud),
// 			RAID: sd.ID,
// 			Dt:   rlib.JSONDate(m.Stmt[i].Dt),
// 		}

// 		//----------------------------------------------------------------------------
// 		// The name of the rule is basically an explanation of the charge or payment.
// 		//----------------------------------------------------------------------------
// 		descr := ""
// 		if m.Stmt[i].T == 1 || m.Stmt[i].T == 2 {
// 			if m.Stmt[i].A.ARID > 0 {
// 				descr = rlib.RRdb.BizTypes[sd.BID].AR[m.Stmt[i].A.ARID].Name
// 			} else {
// 				descr = rlib.RRdb.BizTypes[sd.BID].GLAccounts[m.Stmt[i].A.ATypeLID].Name
// 			}
// 		}
// 		switch m.Stmt[i].T {
// 		case 1: // assessments
// 			amt := m.Stmt[i].Amt
// 			a.Dt = rlib.JSONDate(m.Stmt[i].Dt)
// 			a.ID = rlib.IDtoShortString("ASM", m.Stmt[i].A.ASMID)
// 			a.Descr = descr
// 			a.Reverse = m.Stmt[i].Reverse
// 			a.AsmtAmount = amt
// 			if !a.Reverse {
// 				c += amt
// 				b += amt
// 			} else {
// 				a.Descr += " (" + m.Stmt[i].A.Comment + ")"
// 			}
// 			a.FLAGS = m.Stmt[i].A.FLAGS
// 		case 2: // receipts
// 			amt := m.Stmt[i].Amt
// 			a.RcptAmount = amt
// 			a.ID = rlib.IDtoShortString("RCPT", m.Stmt[i].R.RCPTID)
// 			a.Reverse = m.Stmt[i].Reverse
// 			if m.Stmt[i].A.ASMID > 0 {
// 				descr += " (" + rlib.IDtoShortString("ASM", m.Stmt[i].A.ASMID) + ")"
// 			}
// 			a.Descr = descr
// 			if !a.Reverse {
// 				d += amt
// 				b -= amt
// 			} else {
// 				rcpt := rlib.GetReceipt(m.Stmt[i].R.RCPTID)
// 				comment := ""
// 				if rcpt.RCPTID > 0 {
// 					comment += rcpt.Comment
// 				}
// 				a.Descr += " (" + comment + ")"
// 			}
// 			a.FLAGS = m.Stmt[i].R.FLAGS
// 		}
// 		//---------------------------------------------------
// 		// only add it if it is what was requested...
// 		//---------------------------------------------------
// 		if i >= offset {
// 			a.RentableName = m.Stmt[i].RNT.RentableName
// 			a.Balance = b
// 			a.Recid = int64(i)
// 			g.Records = append(g.Records, a)
// 			count++
// 			if count >= sd.wsSearchReq.Limit { // terminate the loop if we've hit the max count
// 				break
// 			}
// 		}
// 	}

// 	a = StatementDetail{
// 		BID:        sd.BID,
// 		BUD:        rlib.XJSONBud(bud),
// 		RAID:       sd.ID,
// 		Dt:         rlib.JSONDate(m.DtStop.AddDate(0, 0, -1)),
// 		Descr:      "Closing Balance",
// 		Balance:    m.ClosingBal,
// 		AsmtAmount: c,
// 		RcptAmount: d,
// 	}
// 	g.Records = append(g.Records, a)
// 	g.Total = int64(len(m.Stmt) + 2)
// 	g.Status = "success"

// 	w.Header().Set("Content-Type", "application/json")
// 	SvcWriteResponse(&g, w)
// }

// RptStatementReportTable is a returns list of table object for all Statement for a RentalAgreement
func RptStatementReportTable(ri *ReporterInfo) []gotable.Table {
	var m []gotable.Table
	// init some values
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementsByRange.Query(ri.Xbiz.P.BID, ri.D1, ri.D2)
	rlib.Errcheck(err)
	if rlib.IsSQLNoResultsError(err) {
		return m
	}
	defer rows.Close()

	// Spin through all the RentalAgreements that are active in this timeframe
	for rows.Next() {
		var ra rlib.RentalAgreement
		rlib.ReadRentalAgreements(rows, &ra)
		tbl := RptStatementForRA(ri, &ra)
		m = append(m, tbl)
	}
	return m
}

// RptStatementTextReport is a text version of a Statement for a RentalAgreement
func RptStatementTextReport(ri *ReporterInfo) string {
	m := RptStatementReportTable(ri)
	var s string
	// Spin through all the RentalAgreements that are active in this timeframe
	for _, tbl := range m {
		s += ReportToString(&tbl, ri) + "\n"
	}
	return s
}
