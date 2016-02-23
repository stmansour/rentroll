package main

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

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
//      Date     Reference            Charge       Credit       Balance
//      11/30/15 Opening Balance                          170.00
//      12/01/15 Rent                       955.00                   Monthly
//      12/01/15 Late Payment Fee           126.00                      None
//               Assessments Subtotal      1081.00
//      12/07/15 Receipt #143                       1000.00
//      12/08/15 Receipt #144                        251.00
//               Recipts Subtotal                       1251.00
//               Closing Balance                            0.00
func printLedgerHeader(d1, d2 *time.Time, ra *RentalAgreement, x *XPerson, xu *XUnit) {
	fmt.Printf("========================================================================\n")
	fmt.Printf("%-13sLedger for %s - %s", xu.R.Name, d1.Format(RRDATEFMT), d2.Format(RRDATEFMT))
	fmt.Printf("\t Payor: %s\n", x.trn.LastName)
	m := rlib.GetRecurrences(d1, d2, &ra.RentalStart, &ra.RentalStop, ra.Frequency)
	for i := 0; i < len(m); i++ {
		fmt.Printf("%13sRentalAgreement scheduled rent: %8.2f\n", " ", ra.ScheduledRent)
		fmt.Printf("%13sRentable market value rent: %8.2f\n", " ", xu.R.ScheduledRent)
	}
	n := GetUnitSpecialties(xu.U.UNITID)
	t := GetUnitSpecialtyTypes(&n)
	for i := 0; i < len(n); i++ {
		fmt.Printf("%13s%d - %s  %8.2f\n", " ", n[i], t[n[i]].Name, t[n[i]].Fee)
	}

	fmt.Printf("    --------------------------------------------------------------------\n")
	fmt.Printf("    %-8s %-20s %12s %12s %12s\n", "Date", "Reference", "Charge", "Credit", "Balance")
	fmt.Printf("    --------------------------------------------------------------------\n")
}
func printAsmt(a *Assessment, d *time.Time) {
	//s := fmt.Sprintf("%s (%s)", App.asmt2str[a.ASMTID], rlib.RecurIntToString(a.Frequency))
	fmt.Printf("    %8s %-20s %12.2f\n", d.Format(RRDATEFMT), App.asmt2str[a.ASMTID], a.Amount)
}
func printAsmtSubtotal(m float32) {
	fmt.Printf("%13s%-20s %12.2f\n", " ", "Assessments Subtotal", m)
}
func printRcpt(r *Receipt) {
	s := fmt.Sprintf("Receipt #%d", r.RCPTID)
	fmt.Printf("    %8s %-20s %12s %12.2f\n", r.Dt.Format(RRDATEFMT), s, " ", r.Amount)
}
func printRcptSubtotal(m float32) {
	fmt.Printf("%13s%-20s %13s%12.2f\n", " ", "Recipts Subtotal", " ", m)
}
func printLedger(L *Ledger) {
	fmt.Printf("%12s %-20s%39.2f\n", L.Dt.Format(RRDATEFMT), "Opening Balance", L.Balance)
}
func printFinalBal(m float32) {
	fmt.Printf("%13s%-20s%39.2f\n", " ", "Closing Balance", m)
}

// return a slice of assessments for the unit associated with this
// occupancy agreement.
func unitReceipts(oa *RentalAgreement, d1, d2 *time.Time) float32 {
	s := fmt.Sprintf("SELECT RCPTID, PID, Amount, Dt, ApplyToGeneralReceivable, ApplyToSecurityDeposit FROM receipt WHERE RAID=%d and Dt >= '%s' and Dt < '%s'",
		oa.RAID, d1.Format(time.RFC3339), d2.Format(time.RFC3339))
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	var tot = float32(0.0)
	var r Receipt
	for rows.Next() {
		rlib.Errcheck(rows.Scan(&r.RCPTID, &r.PID, &r.Amount, &r.Dt, &r.ApplyToGeneralReceivable, &r.ApplyToSecurityDeposit))
		tot += r.Amount
		printRcpt(&r)
	}
	return tot
}

// return a slice of assessments for the unit associated with this
// occupancy agreement.
func unitAssessments(oa *RentalAgreement, d1, d2 *time.Time) float32 {
	s := fmt.Sprintf("SELECT ASMID,UNITID,ASMTID,Amount,Start,Stop,Frequency FROM assessments WHERE UNITID=%d and Stop >= '%s' and Start < '%s'",
		oa.UNITID, d1.Format(time.RFC3339), d2.Format(time.RFC3339))
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	var tot = float32(0.0)
	var a Assessment
	ap := &a
	for rows.Next() {
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.UNITID, &a.ASMTID, &a.Amount, &a.Start, &a.Stop, &a.Frequency))
		if a.Frequency == rlib.RECURHOURLY {
			// TBD
		} else {
			dl := ap.GetRecurrences(d1, d2)
			for i := 0; i < len(dl); i++ {
				printAsmt(ap, &dl[i])
				tot += ap.Amount
			}
		}
	}
	return tot
}

// RentRollByProperty calculates all charges for the specified property that occur in
// the supplied start / stop time range.
func RentRollByProperty(PRID int, d1, d2 *time.Time) {
	rows, err := App.prepstmt.occAgrByProperty.Query(PRID)
	rlib.Errcheck(err)
	defer rows.Close()

	var oa RentalAgreement
	var xp XPerson

	// Reorganize this to loop by Unit

	for rows.Next() {
		rlib.Errcheck(rows.Scan(&oa.RAID, &oa.OATID, &oa.PRID, &oa.UNITID, &oa.PID, &oa.PrimaryTenant, &oa.RentalStart, &oa.RentalStop, &oa.Renewal, &oa.ProrationMethod, &oa.ScheduledRent, &oa.Frequency, &oa.SecurityDepositAmount, &oa.SpecialProvisions, &oa.LastModTime, &oa.LastModBy))

		// get the Unit...
		var xu XUnit
		GetUnit(oa.UNITID, &xu.U)
		GetXUnit(xu.U.RID, &xu)

		// who is paying
		GetPayor(xu.R.PID, &xp.pay)
		xp.psp.PRSPID = 0
		xp.tnt.TID = 0
		xp.trn.TCID = 0
		GetXPerson(xp.pay.TCID, &xp)

		printLedgerHeader(d1, d2, &oa, &xp, &xu)

		// get Ledger info
		var L Ledger
		rlib.Errcheck(App.prepstmt.getLedger.QueryRow(xu.R.LID).Scan(&L.LID, &L.AccountNo, &L.Dt, &L.Balance, &L.Deposit))
		b := L.Balance
		printLedger(&L)

		// process the active agreements
		// we can reject if RentalStart is > d2 or RentalStop <= d1.  Otherwise, process it
		if !(oa.RentalStart.After(*d2) || (oa.RentalStop.Before(*d1) || oa.RentalStop.Equal(*d1))) {
			// What was budgeted for this unit:
			asmtTot := unitAssessments(&oa, d1, d2)
			printAsmtSubtotal(asmtTot)
			b += asmtTot
			rcptTot := unitReceipts(&oa, d1, d2)
			printRcptSubtotal(rcptTot)
			b -= rcptTot
		}
		printFinalBal(b)
	}
	fmt.Printf("------------------------------------------------------------------------\n")
	rlib.Errcheck(rows.Err())
}

// RentRollAll do a rentroll for all properties
func RentRollAll(d1, d2 time.Time) {
	s := "SELECT PRID,Address,Address2,City,State,PostalCode,Country,Phone,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from property"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var p Property
		rlib.Errcheck(rows.Scan(&p.PRID, &p.Address, &p.Address2, &p.City, &p.State, &p.PostalCode, &p.Country, &p.Phone, &p.Name, &p.DefaultOccupancyType, &p.ParkingPermitInUse, &p.LastModTime, &p.LastModBy))
		fmt.Printf("Property: %s  (%d)\n", p.Name, p.PRID)
		RentRollByProperty(p.PRID, &d1, &d2)
	}
}
