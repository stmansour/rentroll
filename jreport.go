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
	JFMTDESCR   = 50 // description width
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
	fmt.Printf(jfmt.DatedJournalEntryRJ, label, d.Format(RRDATEFMT), ra, rn, glno, a)
}
func printDatedJournalEntryLJ(label string, d time.Time, ra string, rn string, glno string, a float64) {
	fmt.Printf(jfmt.DatedJournalEntryLJ, label, d.Format(RRDATEFMT), ra, rn, glno, a)
}

//
func printJournalHeader(xbiz *XBusiness, d1, d2 *time.Time /*, ra *RentalAgreement, x *XPerson, xu *XUnit*/) {
	// fmt.Printf("         1         2         3         4         5         6         7         8\n")
	// fmt.Printf("12345678901234567890123456789012345678901234567890123456789012345678901234567890\n")
	printJReportDoubleLine()
	fmt.Printf("   Business:  %-13s\n", xbiz.P.Name)
	fmt.Printf("   %s - %s\n", d1.Format(RRDATEFMT), d2.AddDate(0, 0, -1).Format(RRDATEFMT))
	printJReportLine()
	fmt.Printf(jfmt.Hdr, "Description", "Date", "RntAgr", "Rentable", "GL No", "Amount")
	printJReportLine()
}

func processAcctRuleAmount(xbiz *XBusiness, rid int64, d time.Time, rule string, raid int64, r *Rentable, amt float64) {
	m := parseAcctRule(xbiz, rid, &d, &d, rule, amt, float64(1))
	for i := 0; i < len(m); i++ {
		amt := m[i].Amount
		if m[i].Action == "c" {
			amt = -amt
		}
		l := GetLedgerMarkerByGLNo(r.BID, m[i].Account)
		printDatedJournalEntryRJ(l.Name, d, fmt.Sprintf("%d", raid), r.Name, m[i].Account, amt)
	}
}

func textPrintJournalAssessment(xbiz *XBusiness, j *Journal, a *Assessment, r *Rentable, rentDuration, assessmentDuration int64) {
	s := fmt.Sprintf("J%08d  %s", j.JID, App.AsmtTypes[a.ASMTID].Name)
	if rentDuration != assessmentDuration && a.ProrationMethod > 0 {
		s = fmt.Sprintf("%s (%d/%d days)", s, rentDuration, assessmentDuration)
	}
	printJournalSubtitle(s)

	for i := 0; i < len(j.JA); i++ {
		processAcctRuleAmount(xbiz, r.RID, j.Dt, j.JA[i].AcctRule, j.RAID, r, j.JA[i].Amount)
	}
	printJournalSubtitle("")
}

// getPayorLastNames returns an array of strings that contains the last names
// of every payor responsible for this rental agreement
func getPayorLastNames(ra *RentalAgreement, d1, d2 *time.Time) []string {
	var sa []string
	for i := 0; i < len(ra.P); i++ {
		if d1.Before(ra.RentalStop) && d2.After(ra.RentalStart) {
			sa = append(sa, ra.P[i].trn.LastName)
		}
	}
	return sa
}

func textPrintJournalReceipt(xbiz *XBusiness, d1, d2 *time.Time, j *Journal, rcpt *Receipt, cashAcctNo string) {
	rntagr, _ := GetXRentalAgreement(rcpt.RAID, d1, d2)
	sa := getPayorLastNames(&rntagr, d1, d2)
	ps := strings.Join(sa, ",")

	s := fmt.Sprintf("J%08d  Payment - %s  %.2f", j.JID, ps, rcpt.Amount)
	printJournalSubtitle(s)

	// PROCESS EVERY RECEIPT ALLOCATION
	for i := 0; i < len(rcpt.RA); i++ {
		a, _ := GetAssessment(rcpt.RA[i].ASMID)
		r := GetRentable(a.RID)
		m := parseAcctRule(xbiz, r.RID, d1, d2, rcpt.RA[i].AcctRule, rcpt.RA[i].Amount, 1.0)
		printJournalSubtitle("\t" + App.AsmtTypes[a.ASMTID].Name)
		for k := 0; k < len(m); k++ {
			l := GetLedgerMarkerByGLNo(j.BID, m[k].Account)
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

func textPrintJournalUnassociated(xbiz *XBusiness, d1, d2 *time.Time, j *Journal) {
	var r Rentable
	GetRentableByID(j.ID, &r) // j.ID is RID when it is unassociated (RAID == 0)
	printJournalSubtitle(fmt.Sprintf("J%08d  Unassociated: %s", j.JID, r.Name))
	for i := 0; i < len(j.JA); i++ {
		processAcctRuleAmount(xbiz, j.JA[i].RID, j.Dt, j.JA[i].AcctRule, 0, &r, j.JA[i].Amount)

	}
}

func textPrintJournalEntry(xbiz *XBusiness, d1, d2 *time.Time, j *Journal, rentDuration, assessmentDuration int64) {
	switch j.Type {
	case JNLTYPEUNAS:
		textPrintJournalUnassociated(xbiz, d1, d2, j)
	case JNLTYPERCPT:
		rcpt := GetReceipt(j.ID)
		textPrintJournalReceipt(xbiz, d1, d2, j, &rcpt, App.BizTypes[xbiz.P.BID].DefaultAccts[DFLTCASH].GLNumber /*"10001"*/)
	case JNLTYPEASMT:
		a, _ := GetAssessment(j.ID)
		r := GetRentable(a.RID)
		textPrintJournalAssessment(xbiz, j, &a, &r, rentDuration, assessmentDuration)
	default:
		fmt.Printf("printJournalEntry: unrecognized type: %d\n", j.Type)
	}
}

func textReportJournalEntry(xbiz *XBusiness, j *Journal, d1, d2 *time.Time) {
	//-------------------------------------------------------------------
	// over what range of time does this rental apply between d1 & d2
	//-------------------------------------------------------------------
	start := *d1
	stop := *d2
	if j.RAID > 0 {
		ra, _ := GetRentalAgreement(j.RAID)
		if ra.RentalStart.After(start) {
			start = ra.RentalStart
		}
		stop = ra.RentalStop.Add(24 * 60 * time.Minute)
		if stop.After(*d2) {
			stop = *d2
		}
	}
	//-------------------------------------------------------------------------------------------
	// this code needs to be generalized based on the recurrence period and the proration period
	//-------------------------------------------------------------------------------------------
	assessmentDuration := int64(d2.Sub(*d1).Hours() / 24)
	rentDuration := int64(stop.Sub(start).Hours() / 24)

	textPrintJournalEntry(xbiz, d1, d2, j, rentDuration, assessmentDuration)

}

// JournalReportText generates a textual journal report for the supplied business and time range
func JournalReportText(xbiz *XBusiness, d1, d2 *time.Time) {
	printJournalHeader(xbiz, d1, d2)
	rows, err := App.prepstmt.getAllJournalsInRange.Query(xbiz.P.BID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var j Journal
		rlib.Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID))
		GetJournalAllocations(j.JID, &j)
		textReportJournalEntry(xbiz, &j, d1, d2)
	}
	rlib.Errcheck(rows.Err())
}
