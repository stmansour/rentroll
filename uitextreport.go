package main

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// UILedgerTextReport prints a report of data that will be used to format a ledger UI.
// This routine is primarily for testing
func UILedgerTextReport(ui *RRuiSupport) {
	fmt.Printf("LEDGER MARKERS\n%s\nBalances as of:  %s\n\n", ui.B.Name, ui.DtStop)
	fmt.Printf("%-9s  %50s  %10s  %12s\n", "LID", "Name", "GLNumber", "Balance")
	lineLen := 9 + 50 + 10 + 12 + (2 * 3)
	for i := 0; i < len(ui.LDG.XL); i++ {
		fmt.Printf("L%08d  %50.50s  %10s  %12s\n", ui.LDG.XL[i].G.LID, ui.LDG.XL[i].G.Name, ui.LDG.XL[i].G.GLNumber, humanize.FormatFloat("#,###.##", ui.LDG.XL[i].LM.Balance))
	}
	s := ""
	for i := 0; i < lineLen; i++ {
		s += "-"
	}
	fmt.Println(s)
	fmt.Printf("%9s  %50s  %10s  %12s\n", " ", " ", " ", humanize.FormatFloat("#,###.##", LMSum(&ui.LDG.XL)))
}

// UIRentableCountByRentableTypeReport returns a structure containing the count of Rentables for each RentableType
// in the specified time range
func UIRentableCountByRentableTypeReport(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	fmt.Printf("RENTABLE COUNTS BY RENTABLE TYPE\n%s\n%s to %s\n\n",
		xbiz.P.Name, d1.Format("January 2, 2006"), d2.Format("January 2, 2006"))
	m, err := GetRentableCountByRentableType(xbiz, d1, d2)
	if err != nil {
		fmt.Printf("UIRentableCountByRentableTypeReport: GetRentableCountByRentableType returned error: %s\n", err.Error())
	}

	var t rlib.Table
	t.Init()
	t.AddColumn("No. Rentables", 9, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Rentable Type Name", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Style", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Custom Attributes", 50, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

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
	fmt.Print(t.SprintTable(rlib.TABLEOUTTEXT))
}

// UIStatementForRA generates a text Statement for the supplied rental agreement ra.
func UIStatementForRA(xbiz *rlib.XBusiness, d1, d2 *time.Time, ra *rlib.RentalAgreement) {
	rlib.LoadXRentalAgreement(ra.RAID, ra, d1, d2)
	payors := ra.GetPayorNameList(d1, d2)
	fmt.Printf("Statement  -  Rental Agreement %s   Payor(s): %s\n", ra.IDtoString(), strings.Join(payors, ", "))
	fmt.Printf("Period  %s - %s\n", d1.Format("Jan 2, 2006"), d2.AddDate(0, 0, -1).Format("Jan 2, 2006"))
	s := fmt.Sprintf("\n%-10s  %-11s  %-40s  %12s  %12s  %12s\n", "Date", "ID", "Description", "Charge", "Payment", "Balance")
	fmt.Print(s)
	s1 := rlib.Tline(len(s) - 2)
	fmt.Println(s1)

	m := GetStatementData(xbiz, ra.RAID, d1, d2)
	var b = rlib.RoundToCent(m[0].amt) // element 0 is always the account balance
	var c = float64(0)                 // credit
	var d = float64(0)                 // debit
	for i := 0; i < len(m); i++ {
		switch m[i].t {
		case 1: // assessments
			amt := rlib.RoundToCent(m[i].amt)
			c += amt
			b += amt
			fmt.Printf("%10s  ASM%08d  %-40.40s  %12s  %12s  %12s\n",
				m[i].dt.Format(rlib.RRDATEINPFMT), m[i].id, rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[m[i].asmtlid].Name, rlib.RRCommaf(amt), " ", rlib.RRCommaf(b))
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].amt)
			d += amt
			b += amt
			fmt.Printf("%10s  R%010d  %-40.40s  %12s  %12s  %12s\n", m[i].dt.Format(rlib.RRDATEINPFMT), m[i].id, "Payment received", " ", rlib.RRCommaf(m[i].amt), rlib.RRCommaf(b))
		case 3: // opening balance
			fmt.Printf("%10s  %-11s  %-40.40s  %12s  %12s  %12s\n", d1.Format(rlib.RRDATEINPFMT), " ", "Opening Balance", " ", " ", rlib.RRCommaf(b))
		}

	}
	fmt.Println(s1)
	fmt.Printf("%-10s  %-11s  %-40s  %12s  %12s  %12s\n", "Totals", " ", " ", rlib.RRCommaf(c), rlib.RRCommaf(d), rlib.RRCommaf(c+d+m[0].amt))
	fmt.Printf("\n")
}

// UIStatementTextReport is a text version of a Statement for a RentalAgreement
func UIStatementTextReport(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementsByRange.Query(xbiz.P.BID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()

	// Spin through all the RentalAgreements that are active in this timeframe
	for rows.Next() {
		var ra rlib.RentalAgreement
		rlib.ReadRentalAgreements(rows, &ra)
		UIStatementForRA(xbiz, d1, d2, &ra)
	}
}
