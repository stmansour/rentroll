package rrpt

import (
	"fmt"
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

// // UILedgerTextReport prints a report of data that will be used to format a ledger UI.
// // This routine is primarily for testing
// func UILedgerTextReport(ui *RRuiSupport) {
// 	fmt.Printf("LEDGER MARKERS\n%s\nBalances as of:  %s\n\n", ui.B.Name, ui.DtStop)
// 	fmt.Printf("%-9s  %50s  %10s  %12s\n", "LID", "Name", "GLNumber", "Balance")
// 	lineLen := 9 + 50 + 10 + 12 + (2 * 3)
// 	for i := 0; i < len(ui.LDG.XL); i++ {
// 		fmt.Printf("L%08d  %50.50s  %10s  %12s\n", ui.LDG.XL[i].G.LID, ui.LDG.XL[i].G.Name, ui.LDG.XL[i].G.GLNumber, humanize.FormatFloat("#,###.##", ui.LDG.XL[i].LM.Balance))
// 	}
// 	s := ""
// 	for i := 0; i < lineLen; i++ {
// 		s += "-"
// 	}
// 	fmt.Println(s)
// 	fmt.Printf("%9s  %50s  %10s  %12s\n", " ", " ", " ", humanize.FormatFloat("#,###.##", LMSum(&ui.LDG.XL)))
// }

// RentableCountByRentableTypeReport returns a string report containing the count of Rentables for each RentableType
// in the specified time range
func RentableCountByRentableTypeReport(f int, xbiz *rlib.XBusiness, d1, d2 *time.Time) string {
	t := RentableCountByRentableTypeReportTbl(xbiz, d1, d2)
	return t.GetTitle() + t.SprintTable(f)
}

// RentableCountByRentableTypeReportTbl returns an rlib.Table containing the count of Rentables for each RentableType
// in the specified time range
func RentableCountByRentableTypeReportTbl(xbiz *rlib.XBusiness, d1, d2 *time.Time) rlib.Table {
	var t rlib.Table
	t.Init()

	// RentableCountByRentableTypeReport returns a structure containing the count of Rentables for each RentableType
	// in the specified time range
	m, err := GetRentableCountByRentableType(xbiz, d1, d2)
	if err != nil {
		t.SetTitle(fmt.Sprintf("RentableCountByRentableTypeReport: GetRentableCountByRentableType returned error: %s\n", err.Error()))
	}

	t.AddColumn("No. Rentables", 9, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Rentable Type Name", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Style", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Custom Attributes", 50, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.SetTitle(fmt.Sprintf("RENTABLE COUNTS BY RENTABLE TYPE\n%s\n%s to %s\n\n", xbiz.P.Name, d1.Format("January 2, 2006"), d2.Format("January 2, 2006")))

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
func RptStatementForRA(xbiz *rlib.XBusiness, d1, d2 *time.Time, ra *rlib.RentalAgreement) rlib.Table {
	rlib.LoadXRentalAgreement(ra.RAID, ra, d1, d2)
	payors := ra.GetPayorNameList(d1, d2)

	var t rlib.Table
	t.Init()
	s := fmt.Sprintf("Statement  -  Rental Agreement %s   Payor(s): %s\n", ra.IDtoString(), strings.Join(payors, ", "))
	s += fmt.Sprintf("Period  %s - %s\n\n", d1.Format("Jan 2, 2006"), d2.AddDate(0, 0, -1).Format("Jan 2, 2006"))
	t.SetTitle(s)

	// s := fmt.Sprintf("\n%-10s  %-11s  %-40s  %12s  %12s  %12s\n", "Date", "ID", "Description", "Charge", "Payment", "Balance")
	// fmt.Print(s)
	// s1 := rlib.Tline(len(s) - 2)
	// fmt.Println(s1)
	t.AddColumn("Date", 8, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("ID", 11, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Description", 40, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Charge", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Payment", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Balance", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)

	m := GetStatementData(xbiz, ra.RAID, d1, d2)
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
			// fmt.Printf("%10s  ASM%08d  %-40.40s  %12s  %12s  %12s\n",
			// 	m[i].dt.Format(rlib.RRDATEINPFMT), m[i].id, rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[m[i].asmtlid].Name, rlib.RRCommaf(amt), " ", rlib.RRCommaf(b))
			t.Puts(-1, 1, rlib.IDtoString("ASM", m[i].id))
			t.Puts(-1, 2, rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[m[i].asmtlid].Name)
			t.Putf(-1, 3, amt)
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].amt)
			d += amt
			b += amt
			// fmt.Printf("%10s  R%010d  %-40.40s  %12s  %12s  %12s\n", m[i].dt.Format(rlib.RRDATEINPFMT), m[i].id, "Payment received", " ", rlib.RRCommaf(m[i].amt), rlib.RRCommaf(b))
			t.Puts(-1, 1, rlib.IDtoString("ASM", m[i].id))
			t.Puts(-1, 2, rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[m[i].asmtlid].Name)
			t.Putf(-1, 4, amt)
		case 3: // opening balance
			// fmt.Printf("%10s  %-11s  %-40.40s  %12s  %12s  %12s\n", d1.Format(rlib.RRDATEINPFMT), " ", "Opening Balance", " ", " ", rlib.RRCommaf(b))
			t.Puts(-1, 2, "Opening Balance")
		}
		t.Putd(-1, 0, m[i].dt)
		t.Putf(-1, 5, b)

	}
	// fmt.Println(s1)
	t.AddLineAfter(t.RowCount() - 1)

	// fmt.Printf("%-10s  %-11s  %-40s  %12s  %12s  %12s\n", "Totals", " ", " ", rlib.RRCommaf(c), rlib.RRCommaf(d), rlib.RRCommaf(c+d+m[0].amt))
	t.AddRow()
	t.Putf(-1, 3, c)
	t.Putf(-1, 4, d)
	t.Putf(-1, 5, c+d+m[0].amt)

	return t
}

// RptStatementTextReport is a text version of a Statement for a RentalAgreement
func RptStatementTextReport(xbiz *rlib.XBusiness, d1, d2 *time.Time) string {
	s := ""
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementsByRange.Query(xbiz.P.BID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()

	// Spin through all the RentalAgreements that are active in this timeframe
	for rows.Next() {
		var ra rlib.RentalAgreement
		rlib.ReadRentalAgreements(rows, &ra)
		t := RptStatementForRA(xbiz, d1, d2, &ra)
		s += t.String() + "\n"
	}
	return s
}
