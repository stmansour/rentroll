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

// RentableCountByRentableTypeReport returns a string report containing the count of Rentables for each RentableType
// in the specified time range
func RentableCountByRentableTypeReport(ri *ReporterInfo) string {
	t := RentableCountByRentableTypeReportTbl(ri)
	return ReportToString(&t, ri)
}

// RentableCountByRentableTypeReportTbl returns an gotable.Table containing the count of Rentables for each RentableType
// in the specified time range
func RentableCountByRentableTypeReportTbl(ri *ReporterInfo) gotable.Table {
	funcname := "RentableCountByRentableTypeReportTbl"
	var t gotable.Table
	t.Init()
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true
	err := TableReportHeaderBlock(&t, "Rentable Counts By Rentable Type", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
	}
	// RentableCountByRentableTypeReport returns a structure containing the count of Rentables for each RentableType
	// in the specified time range
	m, err := GetRentableCountByRentableType(ri.Xbiz, &ri.D1, &ri.D2)
	if err != nil {
		t.SetTitle(t.GetTitle() + "\n" + fmt.Sprintf("%s: GetRentableCountByRentableType returned error: %s\n", funcname, err.Error()))
	}

	t.AddColumn("No. Rentables", 9, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Rentable Type Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Style", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Custom Attributes", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// need to sort these into a predictable order... they are messing up the tests as they
	// seem to come back in random orders on different runs...
	var keys []int
	for i := 0; i < len(m); i++ {
		keys = append(keys, i)
	}

	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if m[keys[i]].RT.Name > m[keys[j]].RT.Name {
				k := keys[i]
				keys[i] = keys[j]
				keys[j] = k
			}
		}
	}

	for i := 0; i < len(keys); i++ {
		j := int64(keys[i])
		// fmt.Printf("%13d  %-20.20s  %-6s", m[j].Count, m[j].RT.Name, m[j].RT.Style)
		t.AddRow()
		t.Puti(-1, 0, m[j].Count)
		t.Puts(-1, 1, m[j].RT.Name)
		t.Puts(-1, 2, m[j].RT.Style)
		s := ""
		for k, v := range m[j].RT.CA {
			if len(s) > 0 {
				s += ", "
			}
			s += fmt.Sprintf("%s: %s %s", k, v.Value, v.Units)
		}
		t.Puts(-1, 3, s)
	}
	t.TightenColumns()
	return t
}

// RptStatementForRA generates a text Statement for the supplied rental agreement ra.
func RptStatementForRA(ri *ReporterInfo, ra *rlib.RentalAgreement) gotable.Table {
	rlib.LoadXRentalAgreement(ra.RAID, ra, &ri.D1, &ri.D2)
	payors := ra.GetPayorNameList(&ri.D1, &ri.D2)

	var t gotable.Table
	t.Init()
	s := fmt.Sprintf("Statement  -  Rental Agreement %s\nPayor(s): %s\n", ra.IDtoString(), strings.Join(payors, ", "))
	t.SetTitle(ReportHeaderBlock(s, "RptStatementForRA", ri))
	t.AddColumn("Date", 8, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("ID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Description", 40, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Charge", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Payment", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)

	m := GetStatementData(ri.Xbiz, ra.RAID, &ri.D1, &ri.D2)
	var b = rlib.RoundToCent(m[0].amt) // element 0 is always the account balance
	var c = float64(0)                 // credit
	var d = float64(0)                 // debit
	for i := 0; i < len(m); i++ {
		t.AddRow()
		switch m[i].t {
		case 1: // assessments
			amt := rlib.RoundToCent(m[i].amt)
			c += amt
			b += amt
			t.Puts(-1, 1, rlib.IDtoString("ASM", m[i].id))
			t.Puts(-1, 2, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[m[i].asmtlid].Name)
			t.Putf(-1, 3, amt)
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].amt)
			d += amt
			b += amt
			t.Puts(-1, 1, rlib.IDtoString("RCPT", m[i].id))
			t.Puts(-1, 2, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[m[i].asmtlid].Name)
			t.Putf(-1, 4, amt)
		case 3: // opening balance
			t.Puts(-1, 2, "Opening Balance")
		}
		t.Putd(-1, 0, m[i].dt)
		t.Putf(-1, 5, b)

	}
	t.AddLineAfter(t.RowCount() - 1)
	t.AddRow()
	t.Putf(-1, 3, c)
	t.Putf(-1, 4, d)
	t.Putf(-1, 5, c+d+m[0].amt)

	return t
}

// RptStatementTextReport is a text version of a Statement for a RentalAgreement
func RptStatementTextReport(ri *ReporterInfo) string {
	s := ""
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementsByRange.Query(ri.Xbiz.P.BID, ri.D1, ri.D2)
	rlib.Errcheck(err)
	defer rows.Close()
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true
	// Spin through all the RentalAgreements that are active in this timeframe
	for rows.Next() {
		var ra rlib.RentalAgreement
		rlib.ReadRentalAgreements(rows, &ra)
		t := RptStatementForRA(ri, &ra)
		s += ReportToString(&t, ri) + "\n"
	}
	return s
}
