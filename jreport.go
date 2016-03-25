package main

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"
)

// JFMTSPACE et al control the formatting of the journal report
const (
	JFMTSPACE   = 1  // space between cols
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
func printDatedJournalEntryRJ(label string, d time.Time, ra string, rn string, glno string, a float32) {
	fmt.Printf(jfmt.DatedJournalEntryRJ, label, d.Format(RRDATEFMT), ra, rn, glno, a)
}
func printDatedJournalEntryLJ(label string, d time.Time, ra string, rn string, glno string, a float32) {
	fmt.Printf(jfmt.DatedJournalEntryLJ, label, d.Format(RRDATEFMT), ra, rn, glno, a)
}

//
func printJournalHeader(xprop *XBusiness, d1, d2 *time.Time /*, ra *RentalAgreement, x *XPerson, xu *XUnit*/) {
	// fmt.Printf("         1         2         3         4         5         6         7         8\n")
	// fmt.Printf("12345678901234567890123456789012345678901234567890123456789012345678901234567890\n")
	printJReportDoubleLine()
	fmt.Printf("   Business:  %-13s\n", xprop.P.Name)
	fmt.Printf("   %s - %s\n", d1.Format(RRDATEFMT), d2.AddDate(0, 0, -1).Format(RRDATEFMT))
	printJReportLine()
	fmt.Printf(jfmt.Hdr, "Description", "Date", "RntAgr", "Rentable", "GL No", "Amount")
	printJReportLine()
}

func processAcctRuleAmount(d time.Time, rule string, raid int, r *Rentable) {
	m := parseAcctRule(rule, float32(1))
	for i := 0; i < len(m); i++ {
		amt := m[i].Amount
		if m[i].Action == "c" {
			amt = -amt
		}
		l := GetLedgerMarkerByGLNo(m[i].Account)
		printDatedJournalEntryRJ(l.Name, d, fmt.Sprintf("%d", raid), r.Name, m[i].Account, amt)
	}
}

func textPrintJournalAssessment(j *Journal, a *Assessment, r *Rentable, rentDuration, assessmentDuration int) {
	s := fmt.Sprintf("J%08d  %s", j.JID, App.AsmtTypes[a.ASMTID].Name)
	if rentDuration != assessmentDuration && a.ProrationMethod > 0 {
		s = fmt.Sprintf("%s (%d/%d days)", s, rentDuration, assessmentDuration)
	}
	printJournalSubtitle(s)

	for i := 0; i < len(j.JA); i++ {
		processAcctRuleAmount(j.Dt, j.JA[i].AcctRule, j.RAID, r)
	}
	printJournalSubtitle("")
}

func textPrintJournalReceipt(j *Journal, rcpt *Receipt, cashAcctNo string) {
	rntagr, _ := GetRentalAgreement(rcpt.RAID)
	xp := GetXPersonByPID(rntagr.PID)
	s := fmt.Sprintf("J%08d  Payment - %s  %.2f", j.JID, xp.trn.LastName, rcpt.Amount)
	printJournalSubtitle(s)

	for i := 0; i < len(rcpt.RA); i++ {
		a, _ := GetAssessment(rcpt.RA[i].ASMID)
		r := GetRentable(a.RID)
		m := parseAcctRule(rcpt.RA[i].AcctRule, 1.0)
		printJournalSubtitle("\t" + App.AsmtTypes[a.ASMTID].Name)
		for k := 0; k < len(m); k++ {
			l := GetLedgerMarkerByGLNo(m[k].Account)
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

func textPrintJournalEntry(j *Journal, rentDuration, assessmentDuration int) {
	switch j.Type {
	case JNLTYPERCPT:
		rcpt := GetReceipt(j.ID)
		textPrintJournalReceipt(j, &rcpt, App.DefaultCash[rcpt.BID].GLNumber)
	case JNLTYPEASMT:
		a, _ := GetAssessment(j.ID)
		r := GetRentable(a.RID)
		textPrintJournalAssessment(j, &a, &r, rentDuration, assessmentDuration)
	default:
		fmt.Printf("printJournalEntry: unrecognized type: %d\n", j.Type)
	}
}

func textReportJournalEntry(j *Journal, d1, d2 *time.Time) {
	//-------------------------------------------------------------------
	// over what range of time does this rental apply between d1 & d2
	//-------------------------------------------------------------------
	ra, _ := GetRentalAgreement(j.RAID)
	start := *d1
	if ra.RentalStart.After(start) {
		start = ra.RentalStart
	}
	stop := ra.RentalStop.Add(24 * 60 * time.Minute)
	if stop.After(*d2) {
		stop = *d2
	}
	//-------------------------------------------------------------------------------------------
	// this code needs to be generalized based on the recurrence period and the proration period
	//-------------------------------------------------------------------------------------------
	assessmentDuration := int(d2.Sub(*d1).Hours() / 24)
	rentDuration := int(stop.Sub(start).Hours() / 24)

	textPrintJournalEntry(j, rentDuration, assessmentDuration)

}

// JournalReportText generates a textual journal report for the supplied business and time range
func JournalReportText(xprop *XBusiness, d1, d2 *time.Time) {
	printJournalHeader(xprop, d1, d2)
	rows, err := App.prepstmt.getAllJournalsInRange.Query(xprop.P.BID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var j Journal
		rlib.Errcheck(rows.Scan(&j.JID, &j.BID, &j.RAID, &j.Dt, &j.Amount, &j.Type, &j.ID))
		GetJournalAllocations(j.JID, &j)
		textReportJournalEntry(&j, d1, d2)
	}
	rlib.Errcheck(rows.Err())
}
