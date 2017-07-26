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

// StatementEntry is a struct containing references to an Assessment or a Receipt that is
// part of a billing statement associated with a RentalAgreement
type StatementEntry struct {
	t   int              // type: 1 = assessment, 2 = Receipt, 3 = Initial Balance
	a   *rlib.Assessment // for type==1, the pointer to the assessment
	r   *rlib.Receipt    // for type ==2, the pointer to the receipt
	bal float64          // opening balance
}

// StmtEntry describes an entry on a statement
type StmtEntry struct {
	t   int              // 1 = assessment, 2 = Receipt, 3 = Initial Balance
	id  int64            // ASMID if t==1, RCPTID if t==2, n/a if t==3
	a   *rlib.Assessment // for type==1, the pointer to the assessment
	r   *rlib.Receipt    // for type ==2, the pointer to the receipt
	amt float64
	dt  time.Time
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
func GetStatementData(xbiz *rlib.XBusiness, raid int64, d1, d2 *time.Time) []StmtEntry {
	var m []StmtEntry
	lid := temporaryGetAcctsRcv(xbiz.P.BID)
	bal := rlib.GetRAAccountBalance(xbiz.P.BID, lid, raid, d1)
	var initBal = StmtEntry{amt: bal, t: 3, dt: *d1}
	m = append(m, initBal)
	n, err := rlib.GetLedgerEntriesForRAID(d1, d2, raid, lid)
	if err != nil {
		return m
	}
	for i := 0; i < len(n); i++ {
		var se StmtEntry
		se.amt = n[i].Amount
		se.dt = n[i].Dt
		j := rlib.GetJournal(n[i].JID)
		se.t = int(j.Type)
		se.id = j.ID
		if se.t == rlib.JOURNALTYPEASMID {
			// read the assessment to find out what it was for...
			a, err := rlib.GetAssessment(se.id)
			if err != nil {
				rlib.LogAndPrint("rrpt.GetStatementData: error getting asmid %d: %s\n", se.id, err.Error())
				return m
			}
			se.a = &a
		} else if se.t == rlib.JOURNALTYPERCPTID {
			r := rlib.GetReceipt(se.id)
			ja := rlib.GetJournalAllocation(n[i].JAID)
			a, err := rlib.GetAssessment(ja.ASMID)
			if err != nil {
				rlib.LogAndPrint("rrpt.GetStatementData: error getting asmid %d: %s\n", ja.ASMID, err.Error())
				return m
			}
			se.a = &a
			se.r = &r
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

	m := GetStatementData(ri.Xbiz, ra.RAID, &ri.D1, &ri.D2)
	var b = rlib.RoundToCent(m[0].amt) // element 0 is always the account balance
	var c = float64(0)                 // credit
	var d = float64(0)                 // debit
	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		descr := ""
		if m[i].t == 1 || m[i].t == 2 {
			if m[i].a.ARID > 0 {
				descr = rlib.RRdb.BizTypes[ri.Xbiz.P.BID].AR[m[i].a.ARID].Name
			} else {
				descr = rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[m[i].a.ATypeLID].Name
			}
		}
		switch m[i].t {
		case 1: // assessments
			amt := rlib.RoundToCent(m[i].amt)
			c += amt
			b += amt
			tbl.Puts(-1, 1, rlib.IDtoString("ASM", m[i].id))
			tbl.Puts(-1, 2, descr)
			tbl.Putf(-1, 3, amt)
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].amt)
			d += amt
			b += amt
			if m[i].a.ASMID > 0 {
				descr = fmt.Sprintf("%s (%s)", descr, m[i].a.IDtoString())
			}
			tbl.Puts(-1, 1, rlib.IDtoString("RCPT", m[i].id))
			tbl.Puts(-1, 2, descr)
			tbl.Putf(-1, 4, amt)
		case 3: // opening balance
			tbl.Puts(-1, 2, "Opening Balance")
		}
		tbl.Putd(-1, 0, m[i].dt)
		tbl.Putf(-1, 5, b)
	}
	tbl.AddLineAfter(tbl.RowCount() - 1)
	tbl.AddRow()
	tbl.Putf(-1, 3, c)
	tbl.Putf(-1, 4, d)
	tbl.Putf(-1, 5, c+d+m[0].amt)

	return tbl
}

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
