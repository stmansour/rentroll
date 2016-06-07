package main

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"
)

// JFMTSPACE et al control the formatting of the journal report
const (
	JFMTSPACE   = 2  // space between cols
	JFMTINDENT  = 3  // left indent
	JFMTDESCR   = 60 // description width
	JFMTDATE    = 8  // date width
	JFMTRA      = 8  // rental agreement
	JFMTRN      = 15 // rentable name
	JFMTGLNO    = 8  // gl no
	JFMTAMOUNT  = 12 // balance width
	JFMTDECIMAL = 2  // number of decimal places
	JLINELEN    = 5*JFMTSPACE + JFMTDESCR + JFMTDATE + JFMTRA + JFMTRN + JFMTGLNO + JFMTAMOUNT
)

var jfmt struct {
	Indent              string
	Descr               string
	DescrLJ             string
	Dt                  string
	RentalAgr           string
	RentableName        string
	GLNo                string
	Amount              string
	Sp                  string
	Hdr                 string
	AmountHdrStr        string
	DatedJournalEntryRJ string
	DatedJournalEntryLJ string
	JournalHeading      string
}

type jprintctx struct {
	ReportStart time.Time
	ReportStop  time.Time
}

func initJFmt() {
	s := fmt.Sprintf("%%%ds", JFMTINDENT)
	jfmt.Indent = fmt.Sprintf(s, "")
	s = fmt.Sprintf("%%%ds", JFMTSPACE)
	jfmt.Sp = fmt.Sprintf(s, " ")
	jfmt.Descr = fmt.Sprintf("%%%ds", JFMTDESCR)
	jfmt.DescrLJ = fmt.Sprintf("%%-%ds", JFMTDESCR)
	jfmt.Dt = fmt.Sprintf("%%%ds", JFMTDATE)
	jfmt.RentalAgr = fmt.Sprintf("%%%ds", JFMTRA)
	jfmt.RentableName = fmt.Sprintf("%%%ds", JFMTRN)
	jfmt.GLNo = fmt.Sprintf("%%%ds", JFMTGLNO)
	jfmt.Amount = fmt.Sprintf("%%%d.%df", JFMTAMOUNT, JFMTDECIMAL)
	jfmt.AmountHdrStr = fmt.Sprintf("%%%ds", JFMTAMOUNT)

	// Descr, Date, Rental Agreement,  Rentable name,  GL No, Debit / (Credit)
	//                                             Descr               |  Date             | Rental Agreement          | Rentable name|              GL No  |              Debit / (Credit)
	jfmt.DatedJournalEntryRJ = jfmt.Indent + jfmt.Descr + jfmt.Sp + jfmt.Dt + jfmt.Sp + jfmt.RentalAgr + jfmt.Sp + jfmt.RentableName + jfmt.Sp + jfmt.GLNo + jfmt.Sp + jfmt.Amount + "\n"
	jfmt.DatedJournalEntryLJ = jfmt.Indent + jfmt.DescrLJ + jfmt.Sp + jfmt.Dt + jfmt.Sp + jfmt.RentalAgr + jfmt.Sp + jfmt.RentableName + jfmt.Sp + jfmt.GLNo + jfmt.Sp + jfmt.Amount + "\n"
	jfmt.JournalHeading = jfmt.Indent + jfmt.DescrLJ + "\n"
	jfmt.Hdr = jfmt.Indent + jfmt.DescrLJ + jfmt.Sp + jfmt.Dt + jfmt.Sp + jfmt.RentalAgr + jfmt.Sp + jfmt.RentableName + jfmt.Sp + jfmt.GLNo + jfmt.Sp + jfmt.AmountHdrStr + "\n"

}

func printJLineOf(s string) {
	fmt.Println(strings.Repeat(" ", JFMTINDENT) + strings.Repeat(s, JLINELEN/len(s)))
}
func printJReportDoubleLine() {
	printJLineOf("=")
}
func printJReportLine() {
	printJLineOf("-")
}
func printJReportThinLine() {
	printJLineOf(" -")
}
func printJournalSubtitle(label string) {
	fmt.Printf(jfmt.JournalHeading, label)
}
func printDatedJournalEntryRJ(label string, d time.Time, ra string, rn string, glno string, a float64) {
	fmt.Printf(jfmt.DatedJournalEntryRJ, label, d.Format(rlib.RRDATEFMT), ra, rn, glno, a)
}
func printDatedJournalEntryLJ(label string, d time.Time, ra string, rn string, glno string, a float64) {
	fmt.Printf(jfmt.DatedJournalEntryLJ, label, d.Format(rlib.RRDATEFMT), ra, rn, glno, a)
}

//
func printJournalHeader(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	// fmt.Printf("         1         2         3         4         5         6         7         8\n")
	// fmt.Printf("12345678901234567890123456789012345678901234567890123456789012345678901234567890\n")
	printJReportDoubleLine()
	fmt.Printf("   Business:  %-13s\n", xbiz.P.Name)
	fmt.Printf("   %s - %s\n", d1.Format(rlib.RRDATEFMT), d2.AddDate(0, 0, -1).Format(rlib.RRDATEFMT))
	printJReportLine()
	fmt.Printf(jfmt.Hdr, "Description", "Date", "RntAgr", "Rentable", "GL No", "Amount")
	printJReportLine()
}

func processAcctRuleAmount(xbiz *rlib.XBusiness, rid int64, d time.Time, rule string, raid int64, r *rlib.Rentable, amt float64) {
	funcname := "processAcctRuleAmount"
	m := rlib.ParseAcctRule(xbiz, rid, &d, &d, rule, amt, float64(1))
	for i := 0; i < len(m); i++ {
		amt := m[i].Amount
		if m[i].Action == "c" {
			amt = -amt
		}
		l, err := rlib.GetLedgerByGLNo(r.BID, m[i].Account)
		if err != nil {
			fmt.Printf("%s: Could not get ledger for account named %s in business %d\n", funcname, m[i].Account, r.BID)
			fmt.Printf("%s: rule = \"%s\"\n", funcname, rule)
			fmt.Printf("%s: Error = %v\n", funcname, err)
			continue
		}

		printDatedJournalEntryRJ(l.Name, d, fmt.Sprintf("%d", raid), r.Name, m[i].Account, amt)
	}
}

func textPrintJournalAssessment(jctx *jprintctx, xbiz *rlib.XBusiness, j *rlib.Journal, a *rlib.Assessment, r *rlib.Rentable, rentDuration, assessmentDuration int64) {
	s := fmt.Sprintf("J%08d  %s", j.JID, App.AsmtTypes[a.ASMTID].Name)

	//-------------------------------------------------------------------------------------
	// For reporting, we want to show any proration that needs to take place. To determine
	// whether or not there is any proration:
	// 1. Check to see if the Accrual period for the rentable in question is greater than
	//    the ProrationMethod.
	//        *  NO: there is no proration, we don't need to report anything, pf = 1
	// 2. What percent of the accrual period was the rentable "occupied" during the range of interest
	//        *  create a time range equal to the report period [reportDtStart - reportDtStop]
	//        *  if this range is > Accrual Period, trim the range accordingly
	//        *  if this range is > "occupiedrange", trim the range acordingly
	//        *  if the resulting range == Accrual Period then we don't need to report anything, pf = 1
	// 3. Report the prorate factor numerator and denominator:
	//           pf = (resulting range duration)/AccrualPeriod (both in units of the prorationMethod)
	//-------------------------------------------------------------------------------------
	pro := xbiz.RT[r.RTID].Proration

	// fmt.Printf("A0  pro = %d, r.RentalPeriod = %d\n", pro, r.RentalPeriod)
	if r.RentalPeriod > pro && pro != 0 && a.ProrationMethod != 0 { // if accrual > proration then we *may* need to show prorate info
		d1 := jctx.ReportStart // start with the report range
		d2 := jctx.ReportStop  // start with the report range

		// fmt.Printf("A1  d1 = %s, d2 = %s\n", d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))

		if j.Dt.After(d1) { // if this assessment is later move the start time
			d1 = j.Dt
		}
		tmp := d1.Add(rlib.ProrateDuration(r.RentalPeriod, d1)) // start + accrual duration
		if tmp.Before(d2) {                                     // if this occurs prior to the range end...
			d2 = tmp // snap the range end
		}

		// fmt.Printf("B  d1 = %s, d2 = %s\n", d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))

		ra, err := rlib.GetRentalAgreement(a.RAID) // need rental agreement to find Possession time
		rlib.Errlog(err)
		if ra.RAID > 0 { // if we found the rental agreement
			if ra.PossessionStart.After(d1) { // if possession started after d1
				d1 = ra.PossessionStart // snap the begin time
			}
			if ra.PossessionStop.Before(d2) { // if possession ended prior to d2
				d2 = ra.PossessionStop // snap the end time
			}
		}

		// fmt.Printf("C  d1 = %s, d2 = %s\n", d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))

		units := rlib.ProrateDuration(pro, d1) // duration of the unit for proration
		numerator := d2.Sub(d1)
		denominator := rlib.GetProrationRange(d1, d2, r.RentalPeriod, xbiz.RT[r.RTID].Proration)

		// fmt.Printf("   units = %v,  numerator = %v, denominator = %v\n", units, numerator, denominator)

		if numerator != denominator {
			s += fmt.Sprintf(" (%d/%d %s)", numerator/units, denominator/units, rlib.ProrationUnits(pro))
		}

		// s = fmt.Sprintf("%s (%d/%d days)", s, rentDuration, assessmentDuration)
	}

	s += fmt.Sprintf("  %s", r.Name) + " [" + xbiz.RT[r.RTID].Style
	if a.RentalPeriod > rlib.ACCRUALNORECUR {
		s += ", " + rlib.RentalPeriodToString(r.RentalPeriod)
	}
	s += "] " + j.Comment

	printJournalSubtitle(s)

	for i := 0; i < len(j.JA); i++ {
		processAcctRuleAmount(xbiz, r.RID, j.Dt, j.JA[i].AcctRule, j.RAID, r, j.JA[i].Amount)
	}
	printJournalSubtitle("")
}

// getPayorLastNames returns an array of strings that contains the last names
// of every payor responsible for this rental agreement
func getPayorLastNames(ra *rlib.RentalAgreement, d1, d2 *time.Time) []string {
	var sa []string
	for i := 0; i < len(ra.P); i++ {
		if d1.Before(ra.RentalStop) && d2.After(ra.RentalStart) {
			sa = append(sa, ra.P[i].Trn.LastName)
		}
	}
	return sa
}

func textPrintJournalReceipt(xbiz *rlib.XBusiness, jctx *jprintctx, j *rlib.Journal, rcpt *rlib.Receipt, cashAcctNo string) {
	funcname := "textPrintJournalReceipt"
	rntagr, _ := rlib.GetXRentalAgreement(rcpt.RAID, &jctx.ReportStart, &jctx.ReportStop)
	sa := getPayorLastNames(&rntagr, &jctx.ReportStart, &jctx.ReportStop)
	ps := strings.Join(sa, ",")

	s := fmt.Sprintf("J%08d  Payment - %s  %.2f", j.JID, ps, rcpt.Amount)
	printJournalSubtitle(s)

	// PROCESS EVERY RECEIPT ALLOCATION
	for i := 0; i < len(rcpt.RA); i++ {
		a, _ := rlib.GetAssessment(rcpt.RA[i].ASMID)
		r := rlib.GetRentable(a.RID)
		m := rlib.ParseAcctRule(xbiz, r.RID, &jctx.ReportStart, &jctx.ReportStop, rcpt.RA[i].AcctRule, rcpt.RA[i].Amount, 1.0)
		printJournalSubtitle("\t" + App.AsmtTypes[a.ASMTID].Name)
		for k := 0; k < len(m); k++ {
			l, err := rlib.GetLedgerByGLNo(j.BID, m[k].Account)
			if err != nil {
				fmt.Printf("%s: Could not get ledger for account named %s in business %d\n", funcname, m[i].Account, r.BID)
				fmt.Printf("%s: rule = \"%s\"\n", funcname, rcpt.RA[i].AcctRule)
				fmt.Printf("%s: Error = %v\n", funcname, err)
				continue
			}
			amt := m[k].Amount
			if m[k].Action == "c" {
				amt = -amt
			}
			s := fmt.Sprintf("%d", a.RAID)
			printDatedJournalEntryRJ(l.Name, rcpt.Dt, s, r.Name, m[k].Account, amt)
		}
	}
	printJournalSubtitle("")
}

func textPrintJournalUnassociated(xbiz *rlib.XBusiness, jctx *jprintctx, j *rlib.Journal) {
	var r rlib.Rentable
	rlib.GetRentableByID(j.ID, &r) // j.ID is RID when it is unassociated (RAID == 0)
	printJournalSubtitle(fmt.Sprintf("J%08d  Unassociated: %s %s", j.JID, r.Name, j.Comment))
	for i := 0; i < len(j.JA); i++ {
		processAcctRuleAmount(xbiz, j.JA[i].RID, j.Dt, j.JA[i].AcctRule, 0, &r, j.JA[i].Amount)

	}
}

func textPrintJournalEntry(xbiz *rlib.XBusiness, jctx *jprintctx, j *rlib.Journal, rentDuration, assessmentDuration int64) {
	switch j.Type {
	case rlib.JNLTYPEUNAS:
		textPrintJournalUnassociated(xbiz, jctx, j)
	case rlib.JNLTYPERCPT:
		rcpt := rlib.GetReceipt(j.ID)
		textPrintJournalReceipt(xbiz, jctx, j, &rcpt, rlib.RRdb.BizTypes[xbiz.P.BID].DefaultAccts[rlib.DFLTCASH].GLNumber /*"10001"*/)
	case rlib.JNLTYPEASMT:
		a, _ := rlib.GetAssessment(j.ID)
		r := rlib.GetRentable(a.RID)
		textPrintJournalAssessment(jctx, xbiz, j, &a, &r, rentDuration, assessmentDuration)
	default:
		fmt.Printf("printJournalEntry: unrecognized type: %d\n", j.Type)
	}
}

func textReportJournalEntry(xbiz *rlib.XBusiness, j *rlib.Journal, jctx *jprintctx) {

	//-------------------------------------------------------------------------------------
	// over what range of time does this rental apply between jctx.ReportStart & jctx.ReportStop?
	// the rental possession dates may be different than the report range...
	//--------------------------------------------------------------------------------------
	start := jctx.ReportStart // start with the report range
	stop := jctx.ReportStop   // start with the report range
	if j.RAID > 0 {           // is there an associated rental agreement?
		ra, _ := rlib.GetRentalAgreement(j.RAID) // if so, get it
		if ra.PossessionStart.After(start) {     // if posession of rental starts later...
			start = ra.PossessionStart // ...then make an adjustment
		}
		stop = ra.RentalStop // .Add(24 * 60 * time.Minute) -- removing this as all ranges should be NON-INCLUSIVE
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
	textPrintJournalEntry(xbiz, jctx, j, thisAccrualPeriod, fullAccrualPeriod)

}

// JournalReportText generates a textual journal report for the supplied business and time range
func JournalReportText(xbiz *rlib.XBusiness, reportDtStart, reportDtStop *time.Time) {
	jctx := jprintctx{*reportDtStart, *reportDtStop}
	printJournalHeader(xbiz, reportDtStart, reportDtStop)
	rows, err := rlib.RRdb.Prepstmt.GetAllJournalsInRange.Query(xbiz.P.BID, reportDtStart, reportDtStop)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var j rlib.Journal
		rlib.Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID, &j.Comment, &j.LastModTime, &j.LastModBy))
		rlib.GetJournalAllocations(j.JID, &j)
		textReportJournalEntry(xbiz, &j, &jctx)
	}
	rlib.Errcheck(rows.Err())
}
