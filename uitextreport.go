package main

import (
	"fmt"
	"rentroll/rlib"
	"time"

	"github.com/dustin/go-humanize"
)

// UILedgerTextReport prints a report of data that will be used to format a ledger UI.
// This routine is primarily for testing
func UILedgerTextReport(ui *RRuiSupport) {
	fmt.Printf("LEDGER MARKERS\n%s\nOpening Balances:  %s\n\n", ui.B.Name, ui.DtStop.Format("January 2, 2006"))
	fmt.Printf("%40s  %10s  %12s\n", "Name", "GLNumber", "Balance")
	for i := 0; i < len(ui.LDG.XL); i++ {
		fmt.Printf("%40s  %10s  %12s\n", ui.LDG.XL[i].G.Name, ui.LDG.XL[i].G.GLNumber, humanize.FormatFloat("#,###.##", ui.LDG.XL[i].LM.Balance))
	}
	s := ""
	for i := 0; i < 66; i++ {
		s += "-"
	}
	fmt.Println(s)
	fmt.Printf("%40s  %10s  %12s\n", " ", " ", humanize.FormatFloat("#,###.##", LMSum(&ui.LDG.XL)))
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
	s := fmt.Sprintf("%13s  %-18s  %-6s  %s\n", "No. Rentables", "Name", "Style", "Custom Attributes")
	fmt.Print(s)
	w := len(s) - 1 // subtract 1 for the newline character
	s = ""
	for i := 0; i < w; i++ {
		s += "-"
	}
	fmt.Println(s)

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
		fmt.Printf("%13d  %-18s  %-6s", m[j].Count, m[j].RT.Name, m[j].RT.Style)
		for k := 0; k < len(m[j].RT.CA); k++ {
			fmt.Printf("   %s: %s", m[j].RT.CA[k].Name, m[j].RT.CA[k].Value)
		}
		fmt.Printf("\n")
	}
}

// UIStatementForRA generates a text Statement for the supplied rental agreement ra.
func UIStatementForRA(xbiz *rlib.XBusiness, d1, d2 *time.Time, ra *rlib.RentalAgreement) {
	rlib.LoadXRentalAgreement(ra.RAID, ra, d1, d2)
	fmt.Printf("Statement  -  Rental Agreement %s   Payor(s): %s\n", ra.GetName(), ra.PayorsToString())
	fmt.Printf("Period  %s - %s\n", d1.Format("Jan 2, 2006"), d2.AddDate(0, 0, -1).Format("Jan 2, 2006"))
	s := fmt.Sprintf("\n%-10s  %-11s  %-30s  %12s  %12s  %12s\n", "Date", "ID", "Description", "Charge", "Payment", "Balance")
	fmt.Print(s)
	k := len(s) - 2
	s = ""
	for i := 0; i < k; i++ {
		s += "-"
	}
	fmt.Println(s)
	m := GetStatementData(xbiz, ra.RAID, d1, d2)

	var b = rlib.RoundToCent(m[0].bal) // element 0 is always the account balance
	var c = float64(0)                 // credit
	var d = float64(0)                 // debit
	for i := 0; i < len(m); i++ {
		switch m[i].t {
		case 1: // assessments
			dl := m[i].a.GetRecurrences(d1, d2)
			if len(dl) == 0 {
				fmt.Printf("Recurrence date for Assessment ASM%08d not found for period %s - %s\n", m[i].a.ASMID, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
				continue
			}
			for j := 0; j < len(dl); j++ {
				pf := ProrateAssessment(xbiz, m[i].a, &dl[j], d1, d2)
				amt := rlib.RoundToCent(pf * m[i].a.Amount)
				c += amt
				b += amt
				fmt.Printf("%10s  ASM%08d  %-30s  %12s  %12s  %12s\n",
					dl[0].Format(rlib.RRDATEINPFMT), m[i].a.ASMID, rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[m[i].a.ATypeLID].Name, rlib.RRCommaf(amt), " ", rlib.RRCommaf(b))
			}
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].r.Amount)
			d += amt
			b -= amt
			fmt.Printf("%10s  R%010d  %-30s  %12s  %12s  %12s\n", m[i].r.Dt.Format(rlib.RRDATEINPFMT), m[i].r.RCPTID, "Payment received", " ", rlib.RRCommaf(m[i].r.Amount), rlib.RRCommaf(b))
		case 3: // opening balance
			fmt.Printf("%10s  %-11s  %-30s  %12s  %12s  %12s\n", d1.Format(rlib.RRDATEINPFMT), " ", "Opening Balance", " ", " ", rlib.RRCommaf(b))
		}

	}
	fmt.Println(s)
	fmt.Printf("%-10s  %-11s  %-30s  %12s  %12s  %12s\n", "Totals", " ", " ", rlib.RRCommaf(c), rlib.RRCommaf(d), rlib.RRCommaf(c-d+m[0].bal))
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
		rlib.ReadRentalAgreement(rows, &ra)
		UIStatementForRA(xbiz, d1, d2, &ra)
	}
}
