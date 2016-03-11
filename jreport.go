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
	JFMTDESCR   = 40 // description width
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

func processAcctRuleAmount(d time.Time, rule string, raid int, x float32, r *Rentable) {
	if len(rule) > 0 {
		sa := strings.Split(rule, ",")
		for i := 0; i < len(sa); i++ {
			t := strings.TrimSpace(sa[i])
			ta := strings.Split(t, " ")
			action := strings.ToLower(strings.TrimSpace(ta[0]))
			acct := strings.TrimSpace(ta[1])
			amt := x
			if action == "c" {
				amt = -amt
			}
			l := GetLedgerByGLNo(acct)
			printDatedJournalEntryRJ(l.Name, d, fmt.Sprintf("%d", raid), r.Name, acct, amt)
		}
	}
}

func printJournalAssessment(d time.Time, a *Assessment, pf float32, rentDuration, assessmentDuration int) {
	r := GetRentable(a.RID)
	// xp := GetXPersonByPID(r.PID)
	s := fmt.Sprintf("A%08d  %s", a.ASMID, App.AsmtTypes[a.ASMTID].Name)
	if rentDuration != assessmentDuration {
		s = fmt.Sprintf("%s (%d/%d days)", s, rentDuration, assessmentDuration)
	}
	printJournalSubtitle(s)
	processAcctRuleAmount(d, a.AcctRule, a.RAID, a.Amount*pf, &r)
	printJournalSubtitle("")
}

// JournalReport do a journal for the supplied dates
func JournalReport(xprop *XBusiness, d1, d2 *time.Time) {
	//===========================================================
	//  PROCESS ASSESSMSENTS
	//===========================================================
	printJournalHeader(xprop, d1, d2)
	rows, err := App.prepstmt.getAllAssessmentsByBusiness.Query(xprop.P.BID, d2, d1)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a Assessment
		ap := &a
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod, &a.AcctRule, &a.LastModTime, &a.LastModBy))
		if a.Frequency >= rlib.RECURSECONDLY && a.Frequency <= rlib.RECURHOURLY {
			// TBD
			fmt.Printf("Unhandled assessment recurrence type: %d\n", a.Frequency)
		} else {
			dl := ap.GetRecurrences(d1, d2)
			// fmt.Printf("type = %d,  len(dl) = %d\n", a.ASMTID, len(dl))
			for i := 0; i < len(dl); i++ {
				//-------------------------------------------------------------------
				// over what range of time does this rental apply between d1 & d2
				//-------------------------------------------------------------------
				ra, _ := GetRentalAgreement(a.RAID)
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
				pf := float32(1.0)
				if rentDuration != assessmentDuration && a.ProrationMethod > 0 {
					pf = float32(rentDuration) / float32(assessmentDuration)
				} else {
					rentDuration = assessmentDuration
				}
				printJournalAssessment(dl[i], ap, pf, rentDuration, assessmentDuration)
			}
		}
	}
	rlib.Errcheck(rows.Err())

	//===========================================================
	//  PROCESS RECEIPTS
	//===========================================================
	r := GetReceipts(xprop.P.BID, d1, d2)
	for i := 0; i < len(r); i++ {
		rnt := GetRentable(r[i].RAID)
		s := fmt.Sprintf("P%08d  Payment", r[i].RCPTID)
		printJournalSubtitle(s)
		for j := 0; j < len(r[i].RA); j++ {
			// pull the assessment that gave rise to this portion...
			a, err := GetAssessment(r[i].RA[j].ASMID)
			if nil == err {
				//-------------------------------------------------------------------------------------
				// look at the rule for this assessment. Whatever account was debited needs to be
				// credited for this amount and cash should be debited this amount
				//-------------------------------------------------------------------------------------
				if len(a.AcctRule) > 0 {
					sa := strings.Split(a.AcctRule, ",")
					for k := 0; k < len(sa); k++ {
						t := strings.TrimSpace(sa[k])
						ta := strings.Split(t, " ")
						action := strings.ToLower(strings.TrimSpace(ta[0]))
						acct := strings.TrimSpace(ta[1])
						if action == "d" {
							rule := fmt.Sprintf("c %s, d 10001", acct)
							processAcctRuleAmount(r[i].Dt, rule, r[i].RAID, r[i].RA[j].Amount, &rnt)
						}
					}
				}
			} else {
				fmt.Printf("err = %v loading assessment %d\n", err, r[i].RA[j].ASMID)
			}
		}
		printJournalSubtitle("")
	}
}
