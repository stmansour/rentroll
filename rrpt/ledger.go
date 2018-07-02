package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"sort"
	"time"
)

func printLedgerHeader(tbl *gotable.Table, xbiz *rlib.XBusiness, l *rlib.GLAccount, d1, d2 *time.Time) {
	// printTReportDoubleLine()
	// tbl.AddLineBefore(0)
	// s := "LEDGER\n"
	// s += fmt.Sprintf("Business: %-13s\n", xbiz.P.Name)
	// s += fmt.Sprintf("Account:  %s - %s\n", l.GLNumber, l.Name)
	// s += fmt.Sprintf("Period:   %s - %s\n", d1.Format(rlib.RRDATEFMT), d2.AddDate(0, 0, -1).Format(rlib.RRDATEFMT))
	// tbl.SetTitle(s)
	title := fmt.Sprintf("%s (GL Account: %s)\n", l.Name, l.GLNumber)
	tbl.SetTitle(title)
	// tbl.SetTitle(fmt.Sprintf("%s\n", l.Name))
	// tbl.SetSection1(fmt.Sprintf("GL Account: %s\n", l.GLNumber))
}

// returns the payment/accessment reason, Rentable name
func getLedgerEntryDescription(ctx context.Context, l *rlib.LedgerEntry) (string, string, string) {
	sra := ""
	j, err := rlib.GetJournal(ctx, l.JID)
	if err != nil {
		return "x", "x", "x"
	}

	if l.RAID > 0 {
		sra = fmt.Sprintf("%9d", l.RAID)
	}
	switch j.Type {
	case rlib.JNLTYPEUNAS:
		r, err := rlib.GetRentable(ctx, j.ID) // j.ID is set to RID when the type is unassociated
		if err != nil {
			return "x", "x", "x"
		}
		return "Unassociated", r.RentableName, sra
	case rlib.JNLTYPERCPT:
		/*
			ja, err := rlib.GetJournalAllocation(ctx, l.JAID)
			if err != nil {
				return "x", "x", "x"
			}
				a, err := rlib.GetAssessment(ctx, ja.ASMID)
				if err != nil {
					return "x", "x", "x"
				}
		*/
		r, err := rlib.GetRentable(ctx, l.RID)
		if err != nil {
			return "x", "x", "x"
		}
		rcpt, err := rlib.GetReceipt(ctx, j.ID) // ID is the receipt id
		if err != nil {
			return "x", "x", "x"
		}
		p := fmt.Sprintf("Payment #%s - ", rcpt.DocNo)
		if rcpt.ARID > 0 {
			debit := rlib.RRdb.BizTypes[l.BID].AR[rcpt.ARID].DebitLID
			p += fmt.Sprintf("deposited to %s (%s)", rlib.RRdb.BizTypes[l.BID].GLAccounts[debit].GLNumber, rlib.RRdb.BizTypes[l.BID].GLAccounts[debit].Name)
		} /* else {
			p += rlib.RRdb.BizTypes[l.BID].GLAccounts[a.ATypeLID].Name
		}*/
		return p, r.RentableName, sra
	case rlib.JNLTYPEASMT:
		reason := ""
		a, err := rlib.GetAssessment(ctx, j.ID)
		if err != nil {
			return "x", "x", "x"
		}
		r, err := rlib.GetRentable(ctx, a.RID)
		if err != nil {
			return "x", "x", "x"
		}
		if a.ARID > 0 {
			ar, err := rlib.GetAR(ctx, a.ARID)
			if err != nil {
				return "x", "x", "x"
			}
			reason = ar.Name

		} else {
			reason = rlib.RRdb.BizTypes[l.BID].GLAccounts[a.ATypeLID].Name
		}
		return "Assessment - " + reason, r.RentableName, sra
	case rlib.JNLTYPEEXP:
		reason := ""
		a, err := rlib.GetExpense(ctx, j.ID)
		if err != nil {
			return "x", "x", "x"
		}
		r, err := rlib.GetRentable(ctx, a.RID)
		if err != nil {
			return "x", "x", "x"
		}
		if a.ARID > 0 {
			ar, err := rlib.GetAR(ctx, a.ARID)
			if err != nil {
				return "x", "x", "x"
			}
			reason = ar.Name
		}
		return "Expense - " + reason, r.RentableName, sra
	case rlib.JNLTYPEXFER:
		return "Transfer", "", sra

	default:
		fmt.Printf("getLedgerEntryDescription: unrecognized type: %d\n", j.Type)
	}
	return "x", "x", "x"
}

func reportTextProcessLedgerMarker(ctx context.Context, tbl *gotable.Table, xbiz *rlib.XBusiness, lm *rlib.LedgerMarker, d1, d2 *time.Time) {
	l, err := rlib.GetLedger(ctx, lm.LID)
	if err != nil {
		tbl.SetSection3(err.Error())
		return
	}
	if 0 == l.LID {
		tbl.SetSection3(fmt.Sprintf("Ledger not found for LID: %d", lm.LID))
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
	if err != nil {
		tbl.SetSection3(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var l rlib.LedgerEntry
		err = rlib.ReadLedgerEntries(rows, &l)
		if err != nil {
			tbl.SetSection3(err.Error())
			return
		}
		bal += l.Amount
		descr, rn, sra := getLedgerEntryDescription(ctx, &l)
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
	err = rows.Err()
	if err != nil {
		tbl.SetSection3(err.Error())
		return
	}
	// printTReportLine()
	tbl.AddLineAfter(tbl.RowCount() - 1)
	// printLedgerDescrAndBal("Closing Balance", d2.AddDate(0, 0, -1), bal)
	tbl.AddRow()
	tbl.Puts(-1, 0, "Closing Balance")
	tbl.Putd(-1, 1, d2.AddDate(0, 0, -1))
	tbl.Putf(-1, 6, bal)
	// fmt.Printf("\n\n")
}

func initTableColumns(tbl *gotable.Table) {
	tbl.AddColumn("Description", 55, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)      // 0
	tbl.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)               // 1
	tbl.AddColumn("Journal ID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)        // 2
	tbl.AddColumn("Rental Agreement", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT) // 3
	tbl.AddColumn("Rentable Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)    // 4
	tbl.AddColumn("Amount", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)           // 5
	tbl.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)          // 6
}

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

// LedgerActivityReportTable generates a Table Ledger for active accounts during the supplied time range
func LedgerActivityReportTable(ctx context.Context, ri *ReporterInfo) ([]gotable.Table, error) {
	const funcname = "LedgerActivityReportTable"
	var (
		err error
		m   []gotable.Table
		t   int64arr
	)

	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true

	// get the ids of the distinct ledgers that have been updated during &ri.D1-&ri.D2
	// that is, only 1 occurrence of each LID
	rows, err := rlib.RRdb.Dbrr.Query("SELECT DISTINCT LID FROM LedgerEntry ORDER BY Dt,RAID ASC")
	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var lid int64
		err = rows.Scan(&lid)
		if err != nil {
			return m, err
		}
		t = append(t, lid)
	}
	err = rows.Err()
	if err != nil {
		return m, err
	}
	sort.Sort(t)
	// fmt.Printf("Sorted t:  %v\n", t)

	for i := 0; i < len(t); i++ {
		tbl := getRRTable()
		initTableColumns(&tbl)

		// prepare table's title, sections
		err = TableReportHeaderBlock(ctx, &tbl, "LedgerActivity", funcname, ri)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			tbl.SetSection3(err.Error())
			return m, err
		}

		lm, err := rlib.GetLedgerMarkerOnOrBefore(ctx, ri.Xbiz.P.BID, t[i], &ri.D1)
		if err != nil {
			return m, err
		}

		if lm.LMID > 0 {
			reportTextProcessLedgerMarker(ctx, &tbl, ri.Xbiz, &lm, &ri.D1, &ri.D2)
		}

		m = append(m, tbl)
	}
	return m, err
}

// LedgerActivityReport returns text based report from LedgerActivityReportTable
func LedgerActivityReport(ctx context.Context, ri *ReporterInfo) string {
	m, err := LedgerActivityReportTable(ctx, ri)
	if err != nil {
		return "Error while creating ledger activity report: " + err.Error()
	}

	var s string
	// Spin through all the RentalAgreements that are active in this timeframe
	for _, tbl := range m {
		s += ReportToString(&tbl, ri) + "\n"
	}
	return s
}

// LedgerReportTable generates a Table Ledger for the supplied Business and time range
func LedgerReportTable(ctx context.Context, ri *ReporterInfo) ([]gotable.Table, error) {
	const funcname = "LedgerReportTable"
	var (
		err error
		m   []gotable.Table
	)

	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true

	t, err := rlib.GetLedgerList(ctx, ri.Xbiz.P.BID) // this list contains the list of all GLAccount numbers
	if err != nil {
		return m, err
	}

	for i := 0; i < len(t); i++ {
		tbl := getRRTable()
		initTableColumns(&tbl)

		// prepare table's title, sections
		err = TableReportHeaderBlock(ctx, &tbl, "Ledger", funcname, ri)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			tbl.SetSection3(err.Error())
			return m, err
		}

		lm, err := rlib.GetLedgerMarkerOnOrBefore(ctx, ri.Xbiz.P.BID, t[i].LID, &ri.D1)
		if err != nil {
			return m, err
		}

		if lm.LMID > 0 {
			reportTextProcessLedgerMarker(ctx, &tbl, ri.Xbiz, &lm, &ri.D1, &ri.D2)
		}

		m = append(m, tbl)
	}
	return m, err
}

// LedgerReport returns text report for LedgerReportTable
func LedgerReport(ctx context.Context, ri *ReporterInfo) string {
	m, err := LedgerReportTable(ctx, ri)
	if err != nil {
		return "Error while creating ledger report: " + err.Error()
	}
	var s string
	// Spin through all the RentalAgreements that are active in this timeframe
	for _, tbl := range m {
		s += ReportToString(&tbl, ri) + "\n"
	}
	return s
}
