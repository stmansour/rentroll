package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strings"
	"time"
)

// RTIDCount is for counting rentables of a particular type
type RTIDCount struct {
	RT    rlib.RentableType // ID of the types we're counting
	Count int64             // the count
}

// StmtEntry describes an entry on a statement
type StmtEntry struct {
	T   int              // 1 = assessment, 2 = Receipt, 3 = Initial Balance
	ID  int64            // ASMID if t==1, RCPTID if t==2, n/a if t==3
	A   *rlib.Assessment // for type==1, the pointer to the assessment
	R   *rlib.Receipt    // for type ==2, the pointer to the receipt
	Amt float64
	Dt  time.Time
}

// GetRentableCountByRentableType returns a structure containing the count of Rentables for each RentableType
// in the specified time range
func GetRentableCountByRentableType(ctx context.Context, xbiz *rlib.XBusiness, d1, d2 *time.Time) ([]RTIDCount, error) {
	var (
		count int64
		m     []RTIDCount
		err   error
		i     = 0
	)

	for _, v := range xbiz.RT {
		s := fmt.Sprintf("SELECT COUNT(*) FROM RentableTypeRef WHERE RTID=%d AND DtStop>\"%s\" AND DtStart<\"%s\"",
			v.RTID, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
		err = rlib.RRdb.Dbrr.QueryRow(s).Scan(&count)
		if err != nil {
			fmt.Printf("GetRentableCountByRentableType: query=\"%s\"    err = %s\n", s, err.Error())
			// TODO(Steve): shouldn't we returned from here?
			// return m, err
		}
		var rc RTIDCount
		rc.Count = count
		rc.RT = v
		var cerr error
		rc.RT.CA, cerr = rlib.GetAllCustomAttributes(ctx, rlib.ELEMRENTABLETYPE, v.RTID)
		// it's not really an error if we don't find any custom attributes
		// and that thing is already handled during read operation after QueryRow operation
		if cerr != nil {
			err = cerr
			break
		}

		m = append(m, rc)
		i++
	}
	return m, err
}

// yes this is a total hack until I can get rid of some old infrastructure, it only applies to old tests though
func temporaryGetAcctsRcv(ctx context.Context, bid int64) (int64, error) {
	var (
		b   rlib.GLAccount
		err error
	)

	q := "SELECT " + rlib.RRdb.DBFields["GLAccount"] + " FROM GLAccount WHERE Name LIKE \"%receivable%\""
	row := rlib.RRdb.Dbrr.QueryRow(q)
	err = rlib.ReadGLAccount(row, &b)
	if err != nil {
		return 0, err
	}

	if b.LID == 0 {
		b, err = rlib.GetLedgerByGLNo(ctx, bid, "12000")
		if err != nil {
			return 0, err
		}
		if b.LID == 0 {
			// TODO(Steve): Don't know WHY such a code exists...
			/*fmt.Printf("Could not find Accounts Receivable\n")
			debug.PrintStack()
			os.Exit(1) // yes this is terrible.*/
			return 0, err
		}
	}
	return b.LID, err
}

// GetStatementData returns an array of StatementEntries for building a statement
func GetStatementData(ctx context.Context, bid int64, raid int64, d1, d2 *time.Time) ([]StmtEntry, error) {
	const funcname = "GetStatementData"
	var (
		err error
		m   []StmtEntry
	)
	lid, err := temporaryGetAcctsRcv(ctx, bid)
	if err != nil {
		return m, err
	}

	bal, err := rlib.GetRAAccountBalance(ctx, bid, lid, raid, d1)
	if err != nil {
		return m, err
	}

	var initBal = StmtEntry{Amt: bal, T: 3, Dt: *d1}
	m = append(m, initBal)
	n, err := rlib.GetLedgerEntriesForRAID(ctx, d1, d2, raid, lid)
	if err != nil {
		return m, err
	}
	for i := 0; i < len(n); i++ {
		var se StmtEntry
		se.Amt = n[i].Amount
		se.Dt = n[i].Dt
		j, err := rlib.GetJournal(ctx, n[i].JID)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return m, err
		}
		se.T = int(j.Type)
		se.ID = j.ID
		if se.T == rlib.JOURNALTYPEASMID {
			// read the assessment to find out what it was for...
			a, err := rlib.GetAssessment(ctx, se.ID)
			if err != nil {
				rlib.LogAndPrint("rrpt.GetStatementData: error getting asmid %d: %s\n", se.ID, err.Error())
				return m, err
			}
			se.A = &a
		} else if se.T == rlib.JOURNALTYPERCPTID {
			r, err := rlib.GetReceipt(ctx, se.ID)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return m, err
			}

			ja, err := rlib.GetJournalAllocation(ctx, n[i].JAID)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return m, err
			}

			a, err := rlib.GetAssessment(ctx, ja.ASMID)
			if err != nil {
				rlib.LogAndPrint("rrpt.GetStatementData: error getting asmid %d: %s\n", ja.ASMID, err.Error())
				return m, err
			}
			se.A = &a
			se.R = &r
		}

		m = append(m, se)
	}
	return m, err
}

// RptStatementForRA generates a text Statement for the supplied rental agreement ra.
func RptStatementForRA(ctx context.Context, ri *ReporterInfo, ra *rlib.RentalAgreement) gotable.Table {
	const funcname = "RptStatementForRA"
	var (
		err error
	)

	// table init
	tbl := getRRTable()

	tbl.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 45, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Charge", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Payment", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)

	// init and prepare some values before table init
	err = rlib.LoadXRentalAgreement(ctx, ra.RAID, ra, &ri.D1, &ri.D2)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	payors, err := ra.GetPayorNameList(ctx, &ri.D1, &ri.D2)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	// now start to format title after getting enough info for it
	s := fmt.Sprintf("Statement  -  Rental Agreement %s\nPayor(s): %s\n", ra.IDtoString(), strings.Join(payors, ", "))
	err = TableReportHeaderBlock(ctx, &tbl, s, funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	m, err := GetStatementData(ctx, ri.Xbiz.P.BID, ra.RAID, &ri.D1, &ri.D2)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	var b = rlib.RoundToCent(m[0].Amt) // element 0 is always the account balance
	var c = float64(0)                 // credit
	var d = float64(0)                 // debit
	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		descr := ""
		if m[i].T == 1 || m[i].T == 2 {
			if m[i].A.ARID > 0 {
				descr = rlib.RRdb.BizTypes[ri.Xbiz.P.BID].AR[m[i].A.ARID].Name
			} /* else {descr = rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[m[i].A.ATypeLID].Name }*/
		}
		switch m[i].T {
		case 1: // assessments
			amt := rlib.RoundToCent(m[i].Amt)
			c += amt
			b += amt
			tbl.Puts(-1, 1, rlib.IDtoString("ASM", m[i].ID))
			tbl.Puts(-1, 2, descr)
			tbl.Putf(-1, 3, amt)
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].Amt)
			d += amt
			b += amt
			if m[i].A.ASMID > 0 {
				descr = fmt.Sprintf("%s (%s)", descr, m[i].A.IDtoString())
			}
			tbl.Puts(-1, 1, rlib.IDtoString("RCPT", m[i].ID))
			tbl.Puts(-1, 2, descr)
			tbl.Putf(-1, 4, amt)
		case 3: // opening balance
			tbl.Puts(-1, 2, "Opening Balance")
		}
		tbl.Putd(-1, 0, m[i].Dt)
		tbl.Putf(-1, 5, b)
	}
	tbl.AddLineAfter(tbl.RowCount() - 1)
	tbl.AddRow()
	tbl.Putf(-1, 3, c)
	tbl.Putf(-1, 4, d)
	tbl.Putf(-1, 5, c+d+m[0].Amt)

	return tbl
}

// RptStatementReportTable is a returns list of table object for all Statement for a RentalAgreement
func RptStatementReportTable(ctx context.Context, ri *ReporterInfo) ([]gotable.Table, error) {
	const funcname = "RptStatementReportTable"
	var (
		err error
		m   []gotable.Table
	)

	// init some values
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementsByRange.Query(ri.Xbiz.P.BID, ri.D1, ri.D2)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return m, err
	}
	defer rows.Close()

	// Spin through all the RentalAgreements that are active in this timeframe
	for rows.Next() {
		var ra rlib.RentalAgreement
		err = rlib.ReadRentalAgreements(rows, &ra)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return m, err
		}
		tbl := RptStatementForRA(ctx, ri, &ra)
		m = append(m, tbl)
	}
	return m, err
}

// RptStatementTextReport is a text version of a Statement for a RentalAgreement
func RptStatementTextReport(ctx context.Context, ri *ReporterInfo) string {
	m, err := RptStatementReportTable(ctx, ri)
	if err != nil {
		return err.Error()
	}

	var s string
	// Spin through all the RentalAgreements that are active in this timeframe
	for _, tbl := range m {
		s += ReportToString(&tbl, ri) + "\n"
	}
	return s
}
