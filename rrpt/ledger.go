package rrpt

import (
	"fmt"
	"rentroll/rcsv"
	"rentroll/rlib"
	"sort"
	"time"
)

func printLedgerHeader(tbl *rlib.Table, xbiz *rlib.XBusiness, l *rlib.GLAccount, d1, d2 *time.Time) {
	// printTReportDoubleLine()
	// tbl.AddLineBefore(0)
	s := "LEDGER\n"
	s += fmt.Sprintf("Business: %-13s\n", xbiz.P.Name)
	s += fmt.Sprintf("Account:  %s - %s\n", l.GLNumber, l.Name)
	s += fmt.Sprintf("Period:   %s - %s\n", d1.Format(rlib.RRDATEFMT), d2.AddDate(0, 0, -1).Format(rlib.RRDATEFMT))
	tbl.SetTitle(s)
}

// returns the payment/accessment reason, Rentable name
func getLedgerEntryDescription(l *rlib.LedgerEntry) (string, string, string) {
	j := rlib.GetJournal(l.JID)
	sra := fmt.Sprintf("%9d", j.RAID)
	switch j.Type {
	case rlib.JNLTYPEUNAS:
		r := rlib.GetRentable(j.ID) // j.ID is set to RID when the type is unassociated
		return "Unassociated", r.Name, sra
	case rlib.JNLTYPERCPT:
		ja, _ := rlib.GetJournalAllocation(l.JAID)
		a, _ := rlib.GetAssessment(ja.ASMID)
		r := rlib.GetRentable(a.RID)
		rcpt := rlib.GetReceipt(j.ID) // ID is the receipt id
		p := fmt.Sprintf("Payment #%s - ", rcpt.DocNo)
		return p + rlib.RRdb.BizTypes[l.BID].GLAccounts[a.ATypeLID].Name, r.Name, sra
	case rlib.JNLTYPEASMT:
		a, _ := rlib.GetAssessment(j.ID)
		r := rlib.GetRentable(a.RID)
		return "Assessment - " + rlib.RRdb.BizTypes[l.BID].GLAccounts[a.ATypeLID].Name, r.Name, sra

	default:
		fmt.Printf("getLedgerEntryDescription: unrecognized type: %d\n", j.Type)
	}
	return "x", "x", "x"
}

func reportTextProcessLedgerMarker(tbl *rlib.Table, xbiz *rlib.XBusiness, lm *rlib.LedgerMarker, d1, d2 *time.Time) {
	l := rlib.GetLedger(lm.LID)
	if 0 == l.LID {
		return
	}
	bal := lm.Balance
	printLedgerHeader(tbl, xbiz, &l, d1, d2)
	// printLedgerDescrAndBal("Opening Balance", *d1, lm.Balance)
	tbl.AddRow()
	tbl.Puts(-1, 0, "Opening Balance")
	tbl.Putf(-1, 6, lm.Balance)

	// rows, err := rlib.RRdb.Prepstmt.GetLedgerEntriesInRangeByGLNo.Query(l.BID, l.GLNumber, d1, d2)
	rows, err := rlib.RRdb.Prepstmt.GetLedgerEntriesInRangeByLID.Query(l.BID, l.LID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var l rlib.LedgerEntry
		rlib.ReadLedgerEntries(rows, &l)
		bal += l.Amount
		descr, rn, sra := getLedgerEntryDescription(&l)
		// printDatedLedgerEntryRJ(descr, l.Dt, l.JID, sra, rn, l.Amount, bal)
		tbl.AddRow()
		tbl.Puts(-1, 0, descr)
		tbl.Putd(-1, 1, l.Dt)
		tbl.Puts(-1, 2, rlib.IDtoString("J", l.JID))
		tbl.Puts(-1, 3, sra)
		tbl.Puts(-1, 4, rn)
		tbl.Putf(-1, 5, l.Amount)
		tbl.Putf(-1, 6, bal)
	}
	rlib.Errcheck(rows.Err())
	// printTReportLine()
	tbl.AddLineAfter(tbl.Rows() - 1)
	// printLedgerDescrAndBal("Closing Balance", d2.AddDate(0, 0, -1), bal)
	tbl.AddRow()
	tbl.Puts(-1, 0, "Closing Balance")
	tbl.Putd(-1, 1, d2.AddDate(0, 0, -1))
	tbl.Putf(-1, 6, bal)
	// fmt.Printf("\n\n")
}

func initTableColumns(tbl *rlib.Table) {
	tbl.AddColumn("Description", 55, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)      // 0
	tbl.AddColumn("Date", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)               // 1
	tbl.AddColumn("Journal ID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)        // 2
	tbl.AddColumn("Rental Agreement", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT) // 3
	tbl.AddColumn("Rentable Name", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)    // 4
	tbl.AddColumn("Amount", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)           // 5
	tbl.AddColumn("Balance", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)          // 6
}

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

// LedgerActivityReport generates a Table Ledger for active accounts during the supplied time range
func LedgerActivityReport(ri *rcsv.CSVReporterInfo) []rlib.Table {
	var m []rlib.Table
	// get the ids of the distinct ledgers that have been updated during &ri.D1-&ri.D2
	// that is, only 1 occurrence of each LID
	var t int64arr
	rows, err := rlib.RRdb.Dbrr.Query("SELECT DISTINCT LID FROM LedgerEntry ORDER BY Dt,RAID ASC")
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var lid int64
		rlib.Errcheck(rows.Scan(&lid))
		t = append(t, lid)
	}
	rlib.Errcheck(rows.Err())
	sort.Sort(t)
	// fmt.Printf("Sorted t:  %v\n", t)

	for i := 0; i < len(t); i++ {
		lm := rlib.GetLedgerMarkerOnOrBefore(ri.Xbiz.P.BID, t[i], &ri.D1)
		if lm.LMID < 1 {
			fmt.Printf("LedgerActivityReport: GLAccount %d -- no LedgerMarker on or before: %s\n", t[i], ri.D1.Format(rlib.RRDATEFMT))
		} else {
			var tbl rlib.Table
			tbl.Init()
			initTableColumns(&tbl)
			reportTextProcessLedgerMarker(&tbl, ri.Xbiz, &lm, &ri.D1, &ri.D2)
			m = append(m, tbl)
		}
	}
	return m
}

// LedgerReport generates a Table Ledger for the supplied Business and time range
func LedgerReport(ri *rcsv.CSVReporterInfo) []rlib.Table {
	var m []rlib.Table
	t := rlib.GetLedgerList(ri.Xbiz.P.BID) // this list contains the list of all GLAccount numbers
	for i := 0; i < len(t); i++ {
		lm := rlib.GetLedgerMarkerOnOrBefore(ri.Xbiz.P.BID, t[i].LID, &ri.D1)
		if lm.LMID < 1 {
			fmt.Printf("LedgerReport: GLNumber %s -- no LedgerMarker on or before: %s\n", t[i].GLNumber, ri.D1.Format(rlib.RRDATEFMT))
		} else {
			var tbl rlib.Table
			tbl.Init()
			initTableColumns(&tbl)
			reportTextProcessLedgerMarker(&tbl, ri.Xbiz, &lm, &ri.D1, &ri.D2)
			m = append(m, tbl)
		}
	}
	return m
}
