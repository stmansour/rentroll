package main

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"
)

// RentalList contains information about every rental agreement and payor that is responsible
// for a portion of the rental amount over the period being processed.
type RentalList struct {
	xp     *XPerson         // who is paying
	ra     *RentalAgreement // rental agreement for some portion or all of the rental period
	pf     float32          // prorate factor
	din    int              // days in use
	period int              // days in this period
	amount float32          // amount charged to this payor
}

// RRDATEFMT is a shorthand date format used for text output
// Use these values:	Mon Jan 2 15:04:05 MST 2006
// const RRDATEFMT = "02-Jan-2006 3:04PM MST"
// const RRDATEFMT = "01/02/06 3:04PM MST"
const RRDATEFMT = "01/02/06"

// GetRecurrences is a shorthand for assessment variables to get a list
// of dates on which charges must be assessed for a particular interval of time (d1 - d2)
func (a *Assessment) GetRecurrences(d1, d2 *time.Time) []time.Time {
	return rlib.GetRecurrences(d1, d2, &a.Start, &a.Stop, a.Frequency)
}

//          1         2         3         4         5         6         7         8
// 12345678901234567890123456789012345678901234567890123456789012345678901234567890
//     |........|....................|............|............|............|
//        Date        Reference          Charge       Credit      Balance
//   4 1   8    1         20         1    12.2    1    12.2    1    12.2
//
//      Date     Reference               Charge       Credit       Balance
//      11/30/15 Opening Balance                                      170.00
//      12/01/15 Rent                       955.00
//      12/01/15 Late Payment Fee           126.00
//               Assessments Subtotal      1081.00
//      12/07/15 Receipt #143                           1000.00
//      12/08/15 Receipt #144                            251.00
//               Recipts Subtotal                       1251.00
//               Closing Balance                                        0.00

//          1         2         3         4         5         6         7         8
// 12345678901234567890123456789012345678901234567890123456789012345678901234567890
//    ........|..............................................|............|
//      Date        Description                                 Balance
// 3     8    1     46                                       1    12.2

// FMTSPACE and the rest are formatting constants
const (
	FMTSPACE   = 1  // space between cols
	FMTINDENT  = 3  // left indent
	FMTDATE    = 8  // date width
	FMTDESCR   = 60 // description width
	FMTBALANCE = 12 // balance width
	FMTDECIMAL = 2  // number of decimal places
	LINELEN    = 2*FMTSPACE + FMTDATE + FMTDESCR + FMTBALANCE
)

func printLineOf(s string) {
	fmt.Println(strings.Repeat(" ", FMTINDENT) + strings.Repeat(s, LINELEN/len(s)))
}
func printReportDoubleLine() {
	printLineOf("=")
}
func printReportLine() {
	printLineOf("-")
}
func printReportThinLine() {
	printLineOf(" -")
}
func printLedgerHeader(xprop *XBusiness, d1, d2 *time.Time, ra *RentalAgreement, x *XPerson, xu *XUnit) {
	printReportDoubleLine()
	fmt.Printf("   Rentable:  %-13s\tType: %s\n", xu.R.Name, xprop.RT[xu.R.RTID].Name)
	fmt.Printf("   %s - %s\n", d1.Format(RRDATEFMT), d2.AddDate(0, 0, -1).Format(RRDATEFMT))

	if xu.R.RTID == RTRESIDENCE {
		fmt.Printf("   Unit Type: %s - %s %4d sqft\n", xprop.UT[xu.U.UTID].Name, xprop.UT[xu.U.UTID].Style, xprop.UT[xu.U.UTID].SqFt)
	}

	printReportLine()
	//                Indt  Date  spc descr  spc bal
	s := fmt.Sprintf("%%%ds%%-%ds%%%ds%%-%ds%%%ds%%%ds\n", FMTINDENT, FMTDATE, FMTSPACE, FMTDESCR, FMTSPACE, FMTBALANCE)
	// fmt.Print(s)
	fmt.Printf(s, "", "Date", "", "Description", "", "Balance")
	printReportLine()
}
func printDatedLedgerEntryRJ(d time.Time, st string, a float32) {
	// fmt.Printf("   %8s %46s %12.2f\n", d.Format(RRDATEFMT), s, a)
	//                Indt  Date  spc descr spc bal  dec
	s := fmt.Sprintf("%%%ds%%-%ds%%%ds%%%ds%%%ds%%%d.%df\n", FMTINDENT, FMTDATE, FMTSPACE, FMTDESCR, FMTSPACE, FMTBALANCE, FMTDECIMAL)
	// fmt.Print(s)
	fmt.Printf(s, "", d.Format(RRDATEFMT), "", st, "", a)
}
func printDatedLedgerEntryLJ(d time.Time, st string, a float32) {
	// fmt.Printf("   %8s %-46s %12.2f\n", d.Format(RRDATEFMT), st, a)
	//                Indt  Date  spc descr  spc bal  dec
	s := fmt.Sprintf("%%%ds%%-%ds%%%ds%%-%ds%%%ds%%%d.%df\n", FMTINDENT, FMTDATE, FMTSPACE, FMTDESCR, FMTSPACE, FMTBALANCE, FMTDECIMAL)
	fmt.Printf(s, "", d.Format(RRDATEFMT), "", st, "", a)
}
func printReportEntryRJ(st string, a float32) {
	// fmt.Printf("%11s %46s %12.2f\n", " ", st, a)
	//                Indt  Date  spc descr  spc bal  dec
	s := fmt.Sprintf("%%%ds%%%ds%%%ds%%%ds%%%ds%%%d.%df\n", FMTINDENT, FMTDATE, FMTSPACE, FMTDESCR, FMTSPACE, FMTBALANCE, FMTDECIMAL)
	fmt.Printf(s, "", " ", "", st, "", a)
}
func printReportEntryLJ(st string, a float32) {
	// fmt.Printf("%11s %-46s %12.2f\n", " ", s, a)
	//                Indt  Date  spc descr  spc bal  dec
	s := fmt.Sprintf("%%%ds%%-%ds%%%ds%%-%ds%%%ds%%%d.%df\n", FMTINDENT, FMTDATE, FMTSPACE, FMTDESCR, FMTSPACE, FMTBALANCE, FMTDECIMAL)
	fmt.Printf(s, "", " ", "", st, "", a)
}
func printReportStringLJ(st string) {
	// fmt.Printf("%11s %-46s\n", " ", s)
	//                Indt  Date  spc descr  spc bal  dec
	s := fmt.Sprintf("%%%ds%%-%ds%%%ds%%-%ds\n", FMTINDENT, FMTDATE, FMTSPACE, FMTDESCR)
	fmt.Printf(s, "", " ", "", st)
}

// j = justification attribute:  0 = left justify, 1 = right justify
func printAssessment(d time.Time, a *Assessment, j int) {
	t := "credit"
	if App.AsmtTypes[a.ASMTID].Type == 0 {
		t = "debit"
	}
	s := fmt.Sprintf("%s (%s)", App.AsmtTypes[a.ASMTID].Name, t)
	switch j {
	case 0:
		printDatedLedgerEntryLJ(d, s, a.Amount)
	case 1:
		printDatedLedgerEntryRJ(d, s, a.Amount)
	}
}

// return a slice of assessments for the unit associated with this
// occupancy agreement.
// func unitReceipts(ra *RentalAgreement, d1, d2 *time.Time) float32 {
// 	rows, err := App.prepstmt.getUnitReceipts.Query(ra.RAID, d1, d2)
// 	rlib.Errcheck(err)
// 	defer rows.Close()
// 	var tot = float32(0.0)
// 	var r Receipt
// 	for rows.Next() {
// 		rlib.Errcheck(rows.Scan(&r.RCPTID, &r.PID, &r.PMTID, &r.Amount, &r.Dt, &r.ApplyToGeneralReceivable, &r.ApplyToSecurityDeposit))
// 		tot += r.Amount
// 		s := fmt.Sprintf("Receipt %d  (%s)", r.RCPTID, App.PmtTypes[r.PMTID].Name)
// 		printDatedLedgerEntryRJ(r.Dt, s, r.Amount)
// 	}
// 	return tot
// }

func unitSecurityDeposit(ra *RentalAgreement, d1, d2 *time.Time) float32 {
	var t = float32(0.0)
	m := GetSecurityDepositAssessments(ra.UNITID)
	for i := 0; i < len(m); i++ {
		printAssessment(m[i].Start, &m[i], 1)
		switch App.AsmtTypes[m[i].ASMTID].Type {
		case CREDIT:
			t += m[i].Amount
		case DEBIT:
			t -= m[i].Amount
		}
	}
	return t
}

// return a slice of assessments for the unit associated with this
// occupancy agreement.
func unitAssessments(ra *RentalAgreement, d1, d2 *time.Time) float32 {
	rows, err := App.prepstmt.getUnitAssessments.Query(ra.UNITID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	var tot = float32(0.0)
	var a Assessment
	ap := &a
	for rows.Next() {
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod, &a.LastModTime, &a.LastModBy))
		if a.Frequency >= rlib.RECURSECONDLY && a.Frequency <= rlib.RECURHOURLY && a.ASMTID != SECURITYDEPOSIT {
			// TBD
			fmt.Printf("Unhandled assessment recurrence type: %d\n", a.Frequency)
		} else {
			dl := ap.GetRecurrences(d1, d2)
			for i := 0; i < len(dl); i++ {
				printAssessment(dl[i], ap, REPORTJUSTIFYRIGHT)
				tot += ap.Amount
			}
		}
	}
	return tot
}

func getProrationFactor(xprop *XBusiness, xu *XUnit, ra *RentalAgreement, d1, d2 *time.Time) (float32, int, int) {
	// The beginning is the greater of ra.Start... and d1
	t1 := *d1
	if ra.RentalStart.After(t1) {
		t1 = ra.RentalStart
	}
	// The end is the lesser of ra.Stop... and d2
	// remember, stop date is inclusive, so we add 1 day before the subtraction
	t2 := ra.RentalStop.Add(24 * 60 * time.Minute)
	if d2.Before(t2) {
		t2 = *d2
	}

	// fmt.Printf("t1 = %s, t2 = %s, d1 = %s, d2 = %s\n", t1.Format(RRDATEFMT), t2.Format(RRDATEFMT), d1.Format(RRDATEFMT), d2.Format(RRDATEFMT))

	spanDays := int(d2.Sub(*d1).Hours() / 24)
	numDays := int(t2.Sub(t1).Hours() / 24)
	pf := float32(numDays) / float32(spanDays)
	// fmt.Printf("t1=%s, t2=%s, span = %d days,  usage = %d days, pf = %6.4f\n", t1.Format(RRDATEFMT), t2.Format(RRDATEFMT), spanDays, numDays, pf)
	return pf, numDays, spanDays
}

// UnitReport generates a report for the supplied unit and rental agreement.
// There will always be at least one entry in b, that is: b[0]
func UnitReport(xprop *XBusiness, xu *XUnit, b *[]*RentalList, d1, d2 *time.Time) {
	printLedgerHeader(xprop, d1, d2, (*b)[0].ra, (*b)[0].xp, xu)
	budgetedRent := xprop.UT[xu.U.UTID].MarketRate
	printReportEntryRJ("Budgeted Rent", -budgetedRent) // here's what is budgeted
	for i := 0; i < len(xu.S); i++ {
		s := fmt.Sprintf("Specialty: %s", xprop.US[xu.S[i]].Name)
		printReportEntryRJ(s, -xprop.US[xu.S[i]].Fee)
		budgetedRent += xprop.US[xu.S[i]].Fee
	}
	printReportThinLine()
	printReportEntryRJ("Total Budgeted Rent", -budgetedRent)

	totDays := 0
	totPF := float32(0.0)

	for i := 0; i < len(*b); i++ {
		(*b)[i].pf, (*b)[i].din, (*b)[i].period = getProrationFactor(xprop, xu, (*b)[i].ra, d1, d2)
		s := fmt.Sprintf("%d of %d days -> %s %s:  pf = %6.4f", (*b)[i].din, (*b)[i].period, (*b)[i].xp.trn.FirstName, (*b)[i].xp.trn.LastName, (*b)[i].pf)
		printReportStringLJ(s)
		totDays += (*b)[i].din
		totPF += (*b)[i].pf
	}

	if totDays != (*b)[0].period {
		var a RentalList
		a.period = (*b)[0].period
		a.din = a.period - totDays
		a.pf = 1 - totPF
		(*b) = append(*b, &a) // does caller see this newly added value - my assumption is: no
		s := fmt.Sprintf("%2d of %2d days -> vacant:  pf = %6.4f", a.din, a.period, a.pf)
		printReportStringLJ(s)
	}

	// unitSecurityDeposit(ra*RentalAgreement, d1, d2*time.Time)

	printReportEntryLJ("NOTE: Amount to collect from receipts", budgetedRent*totPF)
	printReportEntryLJ("NOTE: loss to vacancy", budgetedRent*(1-totPF))

	// (*b) now holds the list of

	// //===================================================================
	// // OPENING BALANCES...
	// //===================================================================
	// var L Ledger
	// rlib.Errcheck(App.prepstmt.getLedger.QueryRow(xu.R.LID).Scan(&L.LID, &L.AccountNo, &L.Dt, &L.Balance))
	// printDatedLedgerEntryRJ(L.Dt, "Opening General Receivables", L.Balance)
	// printDatedLedgerEntryRJ(L.Dt, "Opening Security Deposit Collected", L.Deposit)

	// //===================================================================
	// // BUDGETED RECEIPTS...
	// //===================================================================
	// // printReportStringLJ(" ")
	// // printReportStringLJ("Budgeted Receipts")
	// var rcptTot = float32(0.0) // receipts
	// var asmtTot = float32(0.0) // assessments

	// // Rent associated with this rentable...  For each recurrence we charge for the rent AND specialties
	// // n := GetUnitSpecialties(xu.U.UNITID)
	// // t := GetUnitSpecialtyTypes(&n)
	// // m := rlib.GetRecurrences(d1, d2, &ra.RentalStart, &ra.RentalStop, ra.Frequency)

	// // for i := 0; i < len(m); i++ {
	// // 	// printDatedLedgerEntryRJ(m[i], "Unit Type Scheduled Rent", ra.ScheduledRent)
	// // 	for a := 0; a < len(n); a++ {
	// // 		s := fmt.Sprintf("Specialty: %s", t[n[a]].Name)
	// // 		printDatedLedgerEntryRJ(m[i], s, t[n[a]].Fee)
	// // 	}
	// // }
	// // unitSecurityDeposit(ra, d1, d2)

	// // totalBudgetedReceipts = br + sdTot

	// //===================================================================
	// // INCOME OFFSETS...
	// //===================================================================
	// // printReportStringLJ("Income Offsets")
	// // process the active agreements
	// // we can reject if RentalStart is > d2 or RentalStop <= d1.  Otherwise, process it
	// if !(ra.RentalStart.After(*d2) || (ra.RentalStop.Before(*d1) || ra.RentalStop.Equal(*d1))) {
	// 	// What was budgeted for this unit:
	// 	asmtTot += unitAssessments(ra, d1, d2)
	// 	printReportEntryRJ("Total Assessments", -asmtTot)
	// 	// printReportStringLJ(" ")
	// 	// printReportStringLJ("Payments Received")
	// 	rcptTot += unitReceipts(ra, d1, d2)
	// 	printReportEntryRJ("Receipts subtotal", rcptTot)
	// }
	printReportDoubleLine()
}

// RentableReport shows the transactions for the supplied rentable over the time period d1-d2
func RentableReport(xprop *XBusiness, xu *XUnit, b *[]*RentalList, d1, d2 *time.Time) {
	printLedgerHeader(xprop, d1, d2, (*b)[0].ra, (*b)[0].xp, xu)
	budgetedRent := xprop.RT[xu.R.RTID].MarketRate
	printReportEntryRJ("Budgeted Rent", -budgetedRent) // here's what is budgeted

	// Get the assessments for this rentable
	m := GetAllRentableAssessments(xu.R.RID, d1, d2)
	for i := 0; i < len(m); i++ {
		a := &m[i]
		if a.Frequency >= rlib.RECURSECONDLY && a.Frequency <= rlib.RECURHOURLY && a.ASMTID != SECURITYDEPOSIT {
			// TBD
			fmt.Printf("Unhandled assessment recurrence type: %d\n", a.Frequency)
		} else {
			dl := a.GetRecurrences(d1, d2)
			for j := 0; j < len(dl); j++ {
				printAssessment(dl[j], a, REPORTJUSTIFYRIGHT)
			}
		}
	}
}

// LedgerReportsByRentable looks for every rental agreement that overlaps [d1,d2) for the supplied unit.
// It then processes each rental agreement over the specified time range.
func LedgerReportsByRentable(xprop *XBusiness, xu *XUnit, d1, d2 *time.Time) {
	rows, err := App.prepstmt.getUnitRentalAgreements.Query(xu.R.UNITID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	// billing := make([]*RentalList, 0)
	var billing []*RentalList
	for rows.Next() {
		var ra RentalAgreement
		rlib.Errcheck(rows.Scan(&ra.RAID, &ra.RATID, &ra.BID, &ra.RID, &ra.UNITID, &ra.PID, &ra.PrimaryTenant, &ra.RentalStart, &ra.RentalStop, &ra.Renewal, &ra.SpecialProvisions, &ra.LastModTime, &ra.LastModBy))
		var xp XPerson
		GetPayor(ra.PID, &xp.pay)
		xp.psp.PRSPID = 0 // force load
		xp.tnt.TID = 0    // force load
		xp.trn.TCID = 0   // force load
		GetXPerson(xp.pay.TCID, &xp)
		var b RentalList
		b.ra = &ra
		b.xp = &xp
		billing = append(billing, &b)
	}
	switch xu.R.RTID {
	case RTRESIDENCE:
		UnitReport(xprop, xu, &billing, d1, d2)
	default:
		RentableReport(xprop, xu, &billing, d1, d2)
	}
	rlib.Errcheck(rows.Err())
}

// LedgerReportsByBusiness calculates all charges for the specified business that occur in
// the supplied start / stop time range.
func LedgerReportsByBusiness(xprop *XBusiness, d1, d2 *time.Time) {
	rows, err := App.prepstmt.getAllRentablesByBusiness.Query(xprop.P.BID)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var xu XUnit
		rlib.Errcheck(rows.Scan(&xu.R.RID, &xu.R.LID, &xu.R.RTID, &xu.R.BID, &xu.R.UNITID, &xu.R.Name, &xu.R.Assignment, &xu.R.Report, &xu.R.LastModTime, &xu.R.LastModBy))
		if xu.R.UNITID > 0 {
			GetXUnit(xu.R.RID, &xu)
			LedgerReportsByRentable(xprop, &xu, d1, d2)
		} else {
			fmt.Printf("Rentable ID %d: name = %s, not a unit\n", xu.R.RID, xu.R.Name)
		}
	}
	rlib.Errcheck(rows.Err())
}

// ReportAll generates the supplied report for all properties
func ReportAll(d1, d2 time.Time, report int) {
	s := "SELECT BID,Address,Address2,City,State,PostalCode,Country,Phone,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var xprop XBusiness
		rlib.Errcheck(rows.Scan(&xprop.P.BID, &xprop.P.Address, &xprop.P.Address2, &xprop.P.City, &xprop.P.State,
			&xprop.P.PostalCode, &xprop.P.Country, &xprop.P.Phone, &xprop.P.Name, &xprop.P.DefaultOccupancyType,
			&xprop.P.ParkingPermitInUse, &xprop.P.LastModTime, &xprop.P.LastModBy))
		GetXBusiness(xprop.P.BID, &xprop)
		// fmt.Printf("Business: %s  (%d)\n", xprop.P.Name, xprop.P.BID)

		switch report {
		case 1:
			JournalReport(&xprop, &d1, &d2)
		case 2:
			LedgerReportsByBusiness(&xprop, &d1, &d2)
		default:
			fmt.Printf("Generating Journal Records for %s to %s\n", d1.Format(RRDATEFMT), d2.Format(RRDATEFMT))
			GenerateJournalRecords(&xprop, &d1, &d2)
		}
	}
}
