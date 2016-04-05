package main

import (
	"fmt"
	"rentroll/rlib"
	"strings"
	"time"
)

// TFMTSPACE et al control the formatting of the journal report
const (
	TFMTSPACE   = 2  // space between cols
	TFMTINDENT  = 3  // left indent
	TFMTDESCR   = 45 // description width
	TFMTDATE    = 8  // date width
	TFMTRA      = 10 // rental agreement
	TFMTJID     = 9  // journal id
	TFMTRN      = 15 // rentable name
	TFMTAMOUNT  = 12 // balance width
	TFMTDECIMAL = 2  // number of decimal places
	TLINELEN    = 6*TFMTSPACE + TFMTDESCR + TFMTDATE + TFMTJID + TFMTRA + TFMTRN + TFMTAMOUNT + TFMTAMOUNT
)

var tfmt struct {
	Indent             string // left indent
	Descr              string // Description
	DescrLJ            string
	Dt                 string // date
	JID                string // transaction id
	RentalAgr          string // rental agreement
	RentableName       string // rentable name
	Amount             string // amount
	Balance            string // balance
	Sp                 string
	Hdr                string
	AmountHdrStr       string
	JIDHdrStr          string
	DatedLedgerEntryRJ string
	DatedLedgerEntryLJ string
	LedgerHeading      string
	DescrAndBal        string
}

func initTFmt() {
	s := fmt.Sprintf("%%%ds", TFMTINDENT)
	tfmt.Indent = fmt.Sprintf(s, "")
	s = fmt.Sprintf("%%%ds", TFMTSPACE)
	tfmt.Sp = fmt.Sprintf(s, " ")
	tfmt.Descr = fmt.Sprintf("%%%ds", TFMTDESCR)                   // Description
	tfmt.DescrLJ = fmt.Sprintf("%%-%ds", TFMTDESCR)                // Description LJ
	tfmt.Dt = fmt.Sprintf("%%%ds", TFMTDATE)                       // date
	tfmt.JID = fmt.Sprintf("%%%dd", TFMTJID)                       // Journal id
	tfmt.JIDHdrStr = fmt.Sprintf("%%%ds", TFMTJID)                 // amount for header
	tfmt.RentalAgr = fmt.Sprintf("%%%ds", TFMTRA)                  // rental agreement
	tfmt.RentableName = fmt.Sprintf("%%%ds", TFMTRN)               // rentable name
	tfmt.Amount = fmt.Sprintf("%%%d.%df", TFMTAMOUNT, TFMTDECIMAL) // digits
	tfmt.AmountHdrStr = fmt.Sprintf("%%%ds", TFMTAMOUNT)           // amount

	// Descr, Date, JID, Rental Agreement,  Rentable name,  Amount, Balance
	tfmt.DatedLedgerEntryRJ = tfmt.Indent + tfmt.Descr + tfmt.Sp + tfmt.Dt + tfmt.Sp + tfmt.JID +
		tfmt.Sp + tfmt.RentalAgr + tfmt.Sp + tfmt.RentableName + tfmt.Sp + tfmt.Amount +
		tfmt.Sp + tfmt.Amount + "\n"
	tfmt.DatedLedgerEntryLJ = tfmt.Indent + tfmt.DescrLJ + tfmt.Sp + tfmt.Dt + tfmt.Sp + tfmt.JID +
		tfmt.Sp + tfmt.RentalAgr + tfmt.Sp + tfmt.RentableName + tfmt.Sp + tfmt.Amount +
		tfmt.Sp + tfmt.Amount + "\n"
	tfmt.LedgerHeading = tfmt.Indent + tfmt.DescrLJ + "\n"
	tfmt.Hdr = tfmt.Indent + tfmt.DescrLJ + tfmt.Sp + tfmt.Dt + tfmt.Sp + tfmt.JIDHdrStr + tfmt.Sp +
		tfmt.RentalAgr + tfmt.Sp + tfmt.RentableName + tfmt.Sp + tfmt.AmountHdrStr + tfmt.Sp +
		tfmt.AmountHdrStr + "\n"
	tfmt.DescrAndBal = tfmt.Indent + tfmt.DescrLJ + tfmt.Sp + tfmt.Dt + tfmt.Sp +
		fmt.Sprintf(fmt.Sprintf("%%%ds", TFMTJID+TFMTSPACE+TFMTRA+TFMTSPACE+TFMTRN+TFMTSPACE+TFMTAMOUNT+TFMTSPACE), " ") +
		tfmt.Amount + "\n"
}

func printTLineOf(s string) {
	fmt.Println(strings.Repeat(" ", TFMTINDENT) + strings.Repeat(s, TLINELEN/len(s)))
}
func printTReportDoubleLine() {
	printTLineOf("=")
}
func printTReportLine() {
	printTLineOf("-")
}
func printTReportThinLine() {
	printTLineOf(" -")
}
func printTSubtitle(label string) {
	fmt.Printf(tfmt.LedgerHeading, label)
}
func printDatedLedgerEntryRJ(label string, d time.Time, jid int64, ra string, rn string, a, b float64) {
	fmt.Printf(tfmt.DatedLedgerEntryRJ, label, d.Format(RRDATEFMT), jid, ra, rn, a, b)
}
func printDatedLedgerEntryLJ(label string, d time.Time, jid int64, ra string, rn string, a, b float64) {
	fmt.Printf(tfmt.DatedLedgerEntryLJ, label, d.Format(RRDATEFMT), jid, ra, rn, a, b)
}
func printLedgerHeaderText(lm *LedgerMarker) {
	printTSubtitle(lm.Name)
}
func printLedgerDescrAndBal(s string, d time.Time, x float64) {
	fmt.Printf(tfmt.DescrAndBal, s, d.Format(RRDATEFMT), x)
}

//
func printLedgerHeader(xbiz *XBusiness, lm *LedgerMarker, d1, d2 *time.Time /*, ra *RentalAgreement, x *XPerson, xu *XUnit*/) {
	printTReportDoubleLine()
	fmt.Printf("   Business:  %-13s\n", xbiz.P.Name)
	printTSubtitle(lm.GLNumber + " - " + lm.Name)
	fmt.Printf("   %s - %s\n", d1.Format(RRDATEFMT), d2.AddDate(0, 0, -1).Format(RRDATEFMT))
	printTReportLine()
	fmt.Printf(tfmt.Hdr, "Description", "Date", "JournalID", "RntAgr", "Rentable", "Amount", "Balance")
	printTReportLine()
}

// returns the payment/accessment reason, rentable name
func getLedgerEntryDescription(l *Ledger) (string, string, string) {
	j, _ := GetJournal(l.JID)
	sra := fmt.Sprintf("%9d", j.RAID)
	switch j.Type {
	case JNLTYPERCPT:
		ja, _ := GetJournalAllocation(l.JAID)
		a, _ := GetAssessment(ja.ASMID)
		r := GetRentable(a.RID)
		return "Payment - " + App.AsmtTypes[a.ASMTID].Name, r.Name, sra
	case JNLTYPEASMT:
		a, _ := GetAssessment(j.ID)
		r := GetRentable(a.RID)
		return "Assessment - " + App.AsmtTypes[a.ASMTID].Name, r.Name, sra

	default:
		fmt.Printf("printJournalEntry: unrecognized type: %d\n", j.Type)
	}
	return "x", "x", "x"
}

func reportTextProcessLedgerMarker(xbiz *XBusiness, lm *LedgerMarker, d1, d2 *time.Time) {
	bal := lm.Balance
	printLedgerHeader(xbiz, lm, d1, d2)
	printLedgerDescrAndBal("Opening Balance", *d1, lm.Balance)
	rows, err := App.prepstmt.getLedgerInRangeByGLNo.Query(lm.BID, lm.GLNumber, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var l Ledger
		rlib.Errcheck(rows.Scan(&l.LID, &l.BID, &l.JID, &l.JAID, &l.GLNumber, &l.Dt, &l.Amount))
		bal += l.Amount
		descr, rn, sra := getLedgerEntryDescription(&l)
		printDatedLedgerEntryRJ(descr, l.Dt, l.JID, sra, rn, l.Amount, bal)
	}
	rlib.Errcheck(rows.Err())
	printTReportLine()
	printLedgerDescrAndBal("Closing Balance", d2.AddDate(0, 0, -1), bal)
	fmt.Printf("\n\n")
}

// LedgerReportText generates a textual journal report for the supplied business and time range
func LedgerReportText(xbiz *XBusiness, d1, d2 *time.Time) {
	// fmt.Printf("BID=%d, d1: %s, d2: %s\n", xbiz.P.BID, d1, d2)
	rows, err := App.prepstmt.getAllLedgerMarkersInRange.Query(xbiz.P.BID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var lm LedgerMarker
		rlib.Errcheck(rows.Scan(&lm.LMID, &lm.BID, &lm.PID, &lm.GLNumber, &lm.Status, &lm.State, &lm.DtStart, &lm.DtStop, &lm.Balance, &lm.Type, &lm.Name))
		reportTextProcessLedgerMarker(xbiz, &lm, d1, d2)
	}
	rlib.Errcheck(rows.Err())
}
