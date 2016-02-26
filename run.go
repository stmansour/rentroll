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
//   |........|..............................................|............|
//      Date        Description                                 Balance
// 2 1   8    1     46                                       1    12.2
func printLedgerHeader(d1, d2 *time.Time, ra *RentalAgreement, x *XPerson, xu *XUnit) {
	fmt.Printf("=======================================================================\n")
	fmt.Printf("   Unit: %-13s\n", xu.R.Name)
	fmt.Printf("   Period %s - %s\t Payor: %s\n", d1.Format(RRDATEFMT), d2.Format(RRDATEFMT), x.trn.LastName)

	var ut UnitType
	GetUnitType(xu.U.UTID, &ut)
	fmt.Printf("   Unit Type: %s - %s %4d sqft  Mkt Rate: %5.2f\n", ut.Name, ut.Style, ut.SqFt, ut.MarketRate)

	fmt.Printf("   --------------------------------------------------------------------\n")
	fmt.Printf("   %-8s %-46s %12s\n", "Date", "Description", "Balance")
	fmt.Printf("   --------------------------------------------------------------------\n")
}
func printLedgerFooter() {
	fmt.Printf("=======================================================================\n")
}
func printDatedLedgerEntryRJ(d time.Time, s string, a float32) {
	fmt.Printf("   %8s %46s %12.2f\n", d.Format(RRDATEFMT), s, a)
}
func printLedgerEntryRJ(s string, a float32) {
	fmt.Printf("%11s %46s %12.2f\n", " ", s, a)
}
func printLedgerEntryLJ(s string, a float32) {
	fmt.Printf("%11s %-46s %12.2f\n", " ", s, a)
}
func printLedgerStringLJ(s string) {
	fmt.Printf("%11s %-46s\n", " ", s)
}

// return a slice of assessments for the unit associated with this
// occupancy agreement.
func unitReceipts(ra *RentalAgreement, d1, d2 *time.Time) float32 {
	// s := fmt.Sprintf("SELECT RCPTID,PID,Amount,Dt,ApplyToGeneralReceivable,ApplyToSecurityDeposit FROM receipt WHERE RAID=%d and Dt >= '%s' and Dt < '%s'",
	// 	ra.RAID, d1.Format(time.RFC3339), d2.Format(time.RFC3339))
	// rows, err := App.dbrr.Query(s)
	rows, err := App.prepstmt.getUnitReceipts.Query(ra.RAID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	var tot = float32(0.0)
	var r Receipt
	for rows.Next() {
		rlib.Errcheck(rows.Scan(&r.RCPTID, &r.PID, &r.Amount, &r.Dt, &r.ApplyToGeneralReceivable, &r.ApplyToSecurityDeposit))
		tot += r.Amount
		printDatedLedgerEntryRJ(r.Dt, fmt.Sprintf("Receipt %d", r.RCPTID), r.Amount)
	}
	return tot
}

// return a slice of assessments for the unit associated with this
// occupancy agreement.
func unitAssessments(ra *RentalAgreement, d1, d2 *time.Time) float32 {
	// s := fmt.Sprintf("SELECT ASMID,UNITID,ASMTID,Amount,Start,Stop,Frequency FROM assessments WHERE UNITID=%d and Stop >= '%s' and Start < '%s'",
	// 	ra.UNITID, d1.Format(time.RFC3339), d2.Format(time.RFC3339))
	// rows, err := App.dbrr.Query(s)
	rows, err := App.prepstmt.getUnitAssessments.Query(ra.UNITID, d1, d2)
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
				printDatedLedgerEntryRJ(dl[i], App.asmt2str[ap.ASMTID], ap.Amount)
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

	var ra RentalAgreement
	var xp XPerson

	// Reorganize this to loop by Unit

	for rows.Next() {
		rlib.Errcheck(rows.Scan(&ra.RAID, &ra.OATID, &ra.PRID, &ra.UNITID, &ra.PID,
			&ra.PrimaryTenant, &ra.RentalStart, &ra.RentalStop, &ra.Renewal, &ra.ProrationMethod,
			&ra.ScheduledRent, &ra.Frequency, &ra.SecurityDepositAmount, &ra.SpecialProvisions,
			&ra.LastModTime, &ra.LastModBy))

		// get the Unit...
		var xu XUnit
		GetUnit(ra.UNITID, &xu.U)
		GetXUnit(xu.U.RID, &xu)

		// who is paying
		GetPayor(xu.R.PID, &xp.pay)
		xp.psp.PRSPID = 0
		xp.tnt.TID = 0
		xp.trn.TCID = 0
		GetXPerson(xp.pay.TCID, &xp)

		printLedgerHeader(d1, d2, &ra, &xp, &xu)

		//===================================================================
		// OPENING BALANCES...
		//===================================================================
		var L Ledger
		rlib.Errcheck(App.prepstmt.getLedger.QueryRow(xu.R.LID).Scan(&L.LID, &L.AccountNo, &L.Dt, &L.Balance, &L.Deposit))
		printDatedLedgerEntryRJ(L.Dt, "Opening General Receivables", L.Balance)
		printDatedLedgerEntryRJ(L.Dt, "Opening Security Deposit Collected", L.Deposit)

		//===================================================================
		// BUDGETED RECEIPTS...
		//===================================================================
		printLedgerStringLJ(" ")
		printLedgerStringLJ("Budgeted Receipts")
		var rcptTot = float32(0.0)               // receipts
		var sdTot = float32(0.0)                 // security deposit total
		var asmtTot = float32(0.0)               // assessments
		var br = float32(0.0)                    // budgeted rent
		var totalBudgetedReceipts = float32(0.0) // total budgeted receipts

		// var ut UnitType
		// GetUnitType(xu.U.UTID, &ut)
		// printDatedLedgerEntryRJ(m[i], "Budgeted Rent", ut.MarketRate)

		// Rent associated with this rentable...  For each recurrence we charge for the rent AND specialties
		n := GetUnitSpecialties(xu.U.UNITID)
		t := GetUnitSpecialtyTypes(&n)
		m := rlib.GetRecurrences(d1, d2, &ra.RentalStart, &ra.RentalStop, ra.Frequency)

		for i := 0; i < len(m); i++ {
			printDatedLedgerEntryRJ(m[i], "Unit Type Scheduled Rent", -ra.ScheduledRent)
			br += ra.ScheduledRent
			for j := 0; j < len(n); j++ {
				s := fmt.Sprintf("Specialty: %s", t[n[j]].Name)
				printDatedLedgerEntryRJ(m[i], s, -t[n[j]].Fee)
				br += t[n[j]].Fee
			}
		}
		printLedgerEntryRJ("Budgeted Rent", -br)
		printLedgerStringLJ(" ")

		printLedgerEntryRJ("Security Deposit Assessment", -sdTot)
		totalBudgetedReceipts = br + sdTot
		printLedgerEntryRJ("Total budgeted receipts", -totalBudgetedReceipts)
		printLedgerStringLJ(" ")

		//===================================================================
		// INCOME OFFSETS...
		//===================================================================
		printLedgerStringLJ("Income Offsets")
		// process the active agreements
		// we can reject if RentalStart is > d2 or RentalStop <= d1.  Otherwise, process it
		if !(ra.RentalStart.After(*d2) || (ra.RentalStop.Before(*d1) || ra.RentalStop.Equal(*d1))) {
			// What was budgeted for this unit:
			asmtTot += unitAssessments(&ra, d1, d2)
			printLedgerEntryRJ("Total Assessments", -asmtTot)

			printLedgerEntryLJ("Total Due Current Period", -(totalBudgetedReceipts + asmtTot))

			printLedgerStringLJ(" ")
			printLedgerStringLJ("Payments Received")

			rcptTot += unitReceipts(&ra, d1, d2)
			printLedgerEntryRJ("Receipts subtotal", rcptTot)

			printLedgerEntryRJ("Final Balance", rcptTot-totalBudgetedReceipts-asmtTot)
		}

	}
	printLedgerFooter()
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
