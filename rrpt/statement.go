package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
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
	t       int   // 1 = assessment, 2 = Receipt, 3 = Initial Balance
	id      int64 // ASMID if t==1, RCPTID if t==2, n/a if t==3
	asmtlid int64 // valid only for t==1, the assessments ATypeLID
	amt     float64
	dt      time.Time
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

// GetStatementData returns an array of StatementEntries for building a statement
func GetStatementData(xbiz *rlib.XBusiness, raid int64, d1, d2 *time.Time) []StmtEntry {
	var m []StmtEntry
	bal := rlib.GetRAAccountBalance(xbiz.P.BID, rlib.RRdb.BizTypes[xbiz.P.BID].DefaultAccts[rlib.GLGENRCV].LID, raid, d1)
	var initBal = StmtEntry{amt: bal, t: 3, dt: *d1}
	m = append(m, initBal)
	n, err := rlib.GetLedgerEntriesForRAID(d1, d2, raid, rlib.RRdb.BizTypes[xbiz.P.BID].DefaultAccts[rlib.GLGENRCV].LID)
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
				return m
			}
			se.asmtlid = a.ATypeLID
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
	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)

	tbl.AddColumn("Date", 8, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 40, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Charge", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Payment", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)

	s := fmt.Sprintf("Statement  -  Rental Agreement %s\nPayor(s): %s\n", ra.IDtoString(), strings.Join(payors, ", "))
	err := TableReportHeaderBlock(&tbl, s, funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)

		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	m := GetStatementData(ri.Xbiz, ra.RAID, &ri.D1, &ri.D2)
	var b = rlib.RoundToCent(m[0].amt) // element 0 is always the account balance
	var c = float64(0)                 // credit
	var d = float64(0)                 // debit
	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		switch m[i].t {
		case 1: // assessments
			amt := rlib.RoundToCent(m[i].amt)
			c += amt
			b += amt
			tbl.Puts(-1, 1, rlib.IDtoString("ASM", m[i].id))
			tbl.Puts(-1, 2, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[m[i].asmtlid].Name)
			tbl.Putf(-1, 3, amt)
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].amt)
			d += amt
			b += amt
			tbl.Puts(-1, 1, rlib.IDtoString("RCPT", m[i].id))
			tbl.Puts(-1, 2, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[m[i].asmtlid].Name)
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

		// first table custom template
		if len(m) == 0 {
			tbl.SetHTMLTemplate("./html/firsttable.html")
		} else {
			// middle table custom template
			tbl.SetHTMLTemplate("./html/middletable.html")
		}

		m = append(m, tbl)
	}
	// last table custom template
	m[len(m)-1].SetHTMLTemplate("./html/lasttable.html")
	return m
}

// RptStatementTextReport is a text version of a Statement for a RentalAgreement
func RptStatementTextReport(ri *ReporterInfo) string {
	m := RptStatementReportTable(ri)
	var s string
	// Spin through all the RentalAgreements that are active in this timeframe
	for _, t := range m {
		s += ReportToString(&t, ri) + "\n"
	}
	return s
}
