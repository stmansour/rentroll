package rrpt

import (
	"fmt"
	"rentroll/rcsv"
	"rentroll/rlib"
	"strings"
	"time"
)

type jprintctx struct {
	ReportStart time.Time
	ReportStop  time.Time
}

func setTitle(tbl *rlib.Table, xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	s := "JOURNAL\n"
	s += fmt.Sprintf("Business: %-13s\n", xbiz.P.Name)
	s += fmt.Sprintf("Period:   %s - %s\n\n", d1.Format(rlib.RRDATEFMT), d2.AddDate(0, 0, -1).Format(rlib.RRDATEFMT))
	tbl.SetTitle(s)
}

func processAcctRuleAmount(tbl *rlib.Table, xbiz *rlib.XBusiness, rid int64, d time.Time, rule string, raid int64, r *rlib.Rentable, amt float64) {
	funcname := "processAcctRuleAmount"
	m := rlib.ParseAcctRule(xbiz, rid, &d, &d, rule, amt, float64(1))
	for i := 0; i < len(m); i++ {
		amt := m[i].Amount
		if m[i].Action == "c" {
			amt = -amt
		}

		// ---------------------------------------------------------
		// This code essentially skips amounts that calculate
		// to something less than .0001 cents (a rounding error)
		x := amt
		if x < 0 {
			x = -x
		}
		if x < 0.0001 {
			continue
		}
		// ---------------------------------------------------------

		l := rlib.GetLedgerByGLNo(r.BID, m[i].Account)
		if 0 == l.LID {
			fmt.Printf("%s: Could not get GLAccount named %s in Business %d\n", funcname, m[i].Account, r.BID)
			fmt.Printf("%s: rule = \"%s\"\n", funcname, rule)
			continue
		}
		// printDatedJournalEntryRJ(l.Name, d, fmt.Sprintf("%d", raid), r.Name, m[i].Account, amt)
		tbl.AddRow()
		tbl.Puts(-1, 1, l.Name)
		tbl.Putd(-1, 2, d)
		tbl.Puts(-1, 3, rlib.IDtoString("RA", raid))
		tbl.Puts(-1, 4, r.Name)
		tbl.Puts(-1, 5, m[i].Account)
		tbl.Putf(-1, 6, amt)
	}
}

func textPrintJournalAssessment(tbl *rlib.Table, jctx *jprintctx, xbiz *rlib.XBusiness, j *rlib.Journal, a *rlib.Assessment, r *rlib.Rentable, rentDuration, assessmentDuration int64) {
	s := rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[a.ATypeLID].Name

	//-------------------------------------------------------------------------------------
	// For reporting, we want to show any proration that needs to take place. To determine
	// whether or not there is any proration:
	// 1. Check to see if the Accrual period for the Rentable in question is greater than
	//    the ProrationCycle.
	//        *  NO: there is no proration, we don't need to report anything, pf = 1
	// 2. What percent of the accrual period was the Rentable "occupied" during the range of interest
	//        *  create a time range equal to the report period [reportDtStart - reportDtStop]
	//        *  if this range is > Accrual Period, trim the range accordingly
	//        *  if this range is > "occupiedrange", trim the range acordingly
	//        *  if the resulting range == Accrual Period then we don't need to report anything, pf = 1
	// 3. Report the prorate factor numerator and denominator:
	//           pf = (resulting range duration)/AccrualPeriod (both in units of the ProrationCycle)
	//-------------------------------------------------------------------------------------
	_, pro, rtid, err := rlib.GetRentCycleAndProration(r, &a.Start, xbiz)
	if err != nil {
		rlib.Ulog("textPrintJournalAssessment: error getting RentCycle and Proration: err = %s\n", err.Error())
		return
	}
	if a.RentCycle > pro && pro != 0 && a.ProrationCycle != 0 { // if accrual > proration then we *may* need to show prorate info
		d1 := jctx.ReportStart // start with the report range
		d2 := jctx.ReportStop  // start with the report range
		if j.Dt.After(d1) {    // if this assessment is later move the start time
			d1 = j.Dt
		}
		tmp := d1.Add(rlib.CycleDuration(a.RentCycle, d1)) // start + accrual duration
		if tmp.Before(d2) {                                // if this occurs prior to the range end...
			d2 = tmp // snap the range end
		}
		ra, err := rlib.GetRentalAgreement(a.RAID) // need rental agreement to find Possession time
		rlib.Errlog(err)
		if ra.RAID > 0 { // if we found the rental agreement
			if ra.RentStart.After(d1) { // if possession started after d1
				d1 = ra.RentStart // snap the begin time
			}
			if ra.RentStop.Before(d2) { // if possession ended prior to d2
				d2 = ra.RentStop // snap the end time
			}
		}

		units := rlib.CycleDuration(pro, d1) // duration of the unit for proration
		numerator := d2.Sub(d1)
		denominator := rlib.GetProrationRange(d1, d2, a.RentCycle, pro)

		if numerator != denominator {
			s += fmt.Sprintf(" (%d/%d %s)", numerator/units, denominator/units, rlib.ProrationUnits(pro))
		}
	}

	s += fmt.Sprintf("  %s", r.Name) + " [" + xbiz.RT[rtid].Style
	if a.RentCycle > rlib.CYCLENORECUR {
		s += ", " + rlib.RentalPeriodToString(a.RentCycle)
	}
	s += "] " + j.Comment

	// printJournalSubtitle(s)
	tbl.AddRow()
	tbl.Puts(-1, 0, j.IDtoString())
	tbl.Puts(-1, 1, s)

	for i := 0; i < len(j.JA); i++ {
		processAcctRuleAmount(tbl, xbiz, r.RID, j.Dt, j.JA[i].AcctRule, j.RAID, r, j.JA[i].Amount)
	}

	// printJournalSubtitle("")
	tbl.AddRow() // nothing in this line, it's blank
}

func textPrintJournalReceipt(tbl *rlib.Table, xbiz *rlib.XBusiness, jctx *jprintctx, j *rlib.Journal, rcpt *rlib.Receipt, cashAcctNo string) {
	funcname := "textPrintJournalReceipt"
	rntagr, _ := rlib.GetXRentalAgreement(rcpt.RAID, &jctx.ReportStart, &jctx.ReportStop)
	// sa := getPayorLastNames(&rntagr, &jctx.ReportStart, &jctx.ReportStop)
	sa := rntagr.GetPayorLastNames(&jctx.ReportStart, &jctx.ReportStop)
	ps := strings.Join(sa, ",")

	s := fmt.Sprintf("Payment - %s   #%s  %.2f", ps, rcpt.DocNo, rcpt.Amount)
	// printJournalSubtitle(s)
	tbl.AddRow()
	tbl.Puts(-1, 0, j.IDtoString())
	tbl.Puts(-1, 1, s)

	// PROCESS EVERY RECEIPT ALLOCATION
	for i := 0; i < len(rcpt.RA); i++ {
		a, _ := rlib.GetAssessment(rcpt.RA[i].ASMID)
		r := rlib.GetRentable(a.RID)
		m := rlib.ParseAcctRule(xbiz, r.RID, &jctx.ReportStart, &jctx.ReportStop, rcpt.RA[i].AcctRule, rcpt.RA[i].Amount, 1.0)
		// printJournalSubtitle("\t" + rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[a.ATypeLID].Name)
		tbl.AddRow()
		tbl.Puts(-1, 1, rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[a.ATypeLID].Name)
		for k := 0; k < len(m); k++ {
			l := rlib.GetLedgerByGLNo(j.BID, m[k].Account)
			if 0 == l.LID {
				fmt.Printf("%s: Could not get GLAccount named %s in Business %d\n", funcname, m[i].Account, r.BID)
				fmt.Printf("%s: rule = \"%s\"\n", funcname, rcpt.RA[i].AcctRule)
				continue
			}
			amt := m[k].Amount
			if m[k].Action == "c" {
				amt = -amt
			}
			// s := fmt.Sprintf("%d", a.RAID)
			// printDatedJournalEntryRJ(l.Name, rcpt.Dt, s, r.Name, m[k].Account, amt)
			tbl.AddRow()
			tbl.Puts(-1, 1, l.Name)
			tbl.Putd(-1, 2, rcpt.Dt)
			tbl.Puts(-1, 3, rlib.IDtoString("RA", a.RAID))
			tbl.Puts(-1, 4, r.Name)
			tbl.Puts(-1, 5, m[k].Account)
			tbl.Putf(-1, 6, amt)
		}
	}
	// printJournalSubtitle("")
	tbl.AddRow() // nothing in this line, it's blank
}

func textPrintJournalUnassociated(tbl *rlib.Table, xbiz *rlib.XBusiness, jctx *jprintctx, j *rlib.Journal) {
	var r rlib.Rentable
	rlib.GetRentableByID(j.ID, &r) // j.ID is RID when it is unassociated (RAID == 0)
	//printJournalSubtitle(s)
	tbl.AddRow()
	tbl.Puts(-1, 0, j.IDtoString())
	tbl.Puts(-1, 1, fmt.Sprintf("Unassociated: %s %s", r.Name, j.Comment))
	for i := 0; i < len(j.JA); i++ {
		processAcctRuleAmount(tbl, xbiz, j.JA[i].RID, j.Dt, j.JA[i].AcctRule, 0, &r, j.JA[i].Amount)

	}
}

func textPrintJournalEntry(tbl *rlib.Table, xbiz *rlib.XBusiness, jctx *jprintctx, j *rlib.Journal, rentDuration, assessmentDuration int64) {
	switch j.Type {
	case rlib.JNLTYPEUNAS:
		textPrintJournalUnassociated(tbl, xbiz, jctx, j)
	case rlib.JNLTYPERCPT:
		rcpt := rlib.GetReceipt(j.ID)
		textPrintJournalReceipt(tbl, xbiz, jctx, j, &rcpt, rlib.RRdb.BizTypes[xbiz.P.BID].DefaultAccts[rlib.GLCASH].GLNumber /*"10001"*/)
	case rlib.JNLTYPEASMT:
		a, _ := rlib.GetAssessment(j.ID)
		r := rlib.GetRentable(a.RID)
		textPrintJournalAssessment(tbl, jctx, xbiz, j, &a, &r, rentDuration, assessmentDuration)
	default:
		fmt.Printf("printJournalEntry: unrecognized type: %d\n", j.Type)
	}
}

func textReportJournalEntry(tbl *rlib.Table, xbiz *rlib.XBusiness, j *rlib.Journal, jctx *jprintctx) {
	//-------------------------------------------------------------------------------------
	// over what range of time does this rental apply between jctx.ReportStart & jctx.ReportStop?
	// the rental possession dates may be different than the report range...
	//--------------------------------------------------------------------------------------
	start := jctx.ReportStart // start with the report range
	stop := jctx.ReportStop   // start with the report range
	if j.RAID > 0 {           // is there an associated rental agreement?
		ra, _ := rlib.GetRentalAgreement(j.RAID) // if so, get it
		if ra.RentStart.After(start) {           // if posession of rental starts later...
			start = ra.RentStart // ...then make an adjustment
		}
		stop = ra.AgreementStop // .Add(24 * 60 * time.Minute) -- removing this as all ranges should be NON-INCLUSIVE
		if stop.After(jctx.ReportStop) {
			stop = jctx.ReportStop
		}
	}

	//-------------------------------------------------------------------------------------------
	// this code needs to be generalized based on the recurrence period and the proration period
	//-------------------------------------------------------------------------------------------
	fullAccrualPeriod := int64(jctx.ReportStop.Sub(jctx.ReportStart).Hours() / 24)
	thisAccrualPeriod := int64(stop.Sub(start).Hours() / 24)

	// fmt.Printf("start = %s, stop = %s, fullAccrualPeriod, thisAccrualPeriod =  %d, %d\n", start.Format(rlib.RRDATEINPFMT), stop.Format(rlib.RRDATEINPFMT), fullAccrualPeriod, thisAccrualPeriod)
	textPrintJournalEntry(tbl, xbiz, jctx, j, thisAccrualPeriod, fullAccrualPeriod)

}

// JournalReport returns a Journal report in an rlib.Table for the supplied Business and time range
func JournalReport(ri *rcsv.CSVReporterInfo) rlib.Table {
	var tbl rlib.Table
	tbl.Init()
	tbl.AddColumn("Journal ID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)  // 0
	tbl.AddColumn("Description", 70, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT) // 1
	tbl.AddColumn("Date", 8, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)           // 2
	tbl.AddColumn("RntAgr", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)      // 3
	tbl.AddColumn("Rentable", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)    // 4
	tbl.AddColumn("GLAccount", 8, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)    // 5
	tbl.AddColumn("Amount", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)      // 6

	jctx := jprintctx{ri.D1, ri.D2}
	setTitle(&tbl, ri.Xbiz, &ri.D1, &ri.D2)
	rows, err := rlib.RRdb.Prepstmt.GetAllJournalsInRange.Query(ri.Xbiz.P.BID, &ri.D1, &ri.D2)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var j rlib.Journal
		rlib.Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID, &j.Comment, &j.LastModTime, &j.LastModBy))
		// fmt.Printf("JournalReportText: JID = %d\n", j.JID)
		rlib.GetJournalAllocations(j.JID, &j)
		textReportJournalEntry(&tbl, ri.Xbiz, &j, &jctx)
	}
	rlib.Errcheck(rows.Err())
	return tbl
}
