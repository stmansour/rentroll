package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// RRreportRentalAgreementsTable generates a table object for All rental agreements related with Business
func RRreportRentalAgreementsTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportRentalAgreementsTable"
	var (
		err error
	)

	// init and prepare some values before table init
	totalErrs := 0

	// table init
	tbl := getRRTable()

	tbl.AddColumn("RAID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Payor", 60, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("User", 60, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Agreement Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Agreement Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Possession Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Possession Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Cycle Epoch", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Renewal", 2, gotable.CELLINT, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Unspecified Adults", 6, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Unspecified Children", 6, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Special Provisions", 40, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Notes", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err = TableReportHeaderBlock(ctx, &tbl, "Rental Agreement", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreements.Query(ri.Bid)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	var raid int64
	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC)

	for rows.Next() {
		var p rlib.RentalAgreement
		err = rows.Scan(&raid)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			totalErrs++
			continue
		}

		p, err = rlib.GetXRentalAgreement(ctx, raid, &d1, &d2)
		if err != nil {
			totalErrs++
			rlib.Ulog("%s: rlib.GetXRentalAgreement returned err = %v\n", funcname, err)
			continue
		}

		note := ""
		if p.NLID > 0 {
			nl, err := rlib.GetNoteList(ctx, p.NLID)
			if err != nil {
				totalErrs++
				rlib.Ulog("%s: rlib.GetNoteList returned err = %s\n", funcname, err.Error())
				continue
			}
			if len(nl.N) > 0 {
				note = nl.N[0].Comment
			}
		}

		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())

		payors, err := p.GetPayorNameList(ctx, &p.AgreementStart, &p.AgreementStop)
		if err != nil {
			totalErrs++
			rlib.Ulog("%s: p.GetPayorNameList returned err = %s\n", funcname, err.Error())
			continue
		}
		tbl.Puts(-1, 1, strings.Join(payors, ","))

		users, err := p.GetUserNameList(ctx, &p.AgreementStart, &p.AgreementStop)
		if err != nil {
			totalErrs++
			rlib.Ulog("%s: p.GetUserNameList returned err = %s\n", funcname, err.Error())
			continue
		}
		tbl.Puts(-1, 2, strings.Join(users, ", "))
		tbl.Putd(-1, 3, p.AgreementStart)
		tbl.Putd(-1, 4, p.AgreementStop)
		tbl.Putd(-1, 5, p.PossessionStart)
		tbl.Putd(-1, 6, p.PossessionStop)
		tbl.Putd(-1, 7, p.RentStart)
		tbl.Putd(-1, 8, p.RentStop)
		tbl.Putd(-1, 9, p.RentCycleEpoch)
		tbl.Puti(-1, 10, p.Renewal)
		tbl.Puti(-1, 11, p.UnspecifiedAdults)
		tbl.Puti(-1, 12, p.UnspecifiedChildren)
		tbl.Puts(-1, 13, p.SpecialProvisions)
		tbl.Puts(-1, 14, note)
	}
	err = rows.Err()
	if err != nil {
		totalErrs++
		rlib.Ulog("%s: rows.Err() returned err = %s\n", funcname, err.Error())
	}

	tbl.TightenColumns()
	if totalErrs > 0 {
		errMsg := fmt.Sprintf("Encountered %d errors while creating this report. See log.", totalErrs)
		tbl.SetSection3(errMsg)
	}
	return tbl
}

// RRreportRentalAgreements generates a report of all Businesses defined in the database.
func RRreportRentalAgreements(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportRentalAgreementsTable(ctx, ri)
	return ReportToString(&tbl, ri)
}

// RRRentalAgreementStatementTable returns gotable.Table for rental agreement statements
func RRRentalAgreementStatementTable(ctx context.Context, BID, RAID int64, d1, d2 *time.Time) gotable.Table {
	const funcname = "SvcStatementDetails"
	var (
		err  error
		xbiz rlib.XBusiness
	)
	rlib.Console("Entered %s\n", funcname)

	// table init
	tbl := getRRTable()

	const (
		Date         = 0
		ID           = iota
		Rentable     = iota
		Description  = iota
		Assessment   = iota
		AppliedFunds = iota
		Balance      = iota
	)

	//
	// UGH!
	//=======================================================================
	err = rlib.InitBizInternals(BID, &xbiz)
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}

	rlib.Console("BID = %d\n", BID)
	_, ok := rlib.RRdb.BizTypes[BID]
	if !ok {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d]", BID)
		tbl.SetSection3(e.Error())
		return tbl
	}
	if len(rlib.RRdb.BizTypes[BID].GLAccounts) == 0 {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d].GLAccounts", BID)
		tbl.SetSection3(e.Error())
		return tbl
	}
	//=======================================================================
	// UGH!

	//--------------------------------------------
	// Get the statement data...
	//--------------------------------------------
	m, err := rlib.GetRAIDStatementInfo(ctx, RAID, d1, d2)
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}

	tbl.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ID", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Rentable", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Assessment", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Applied Funds", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)

	tbl.SetTitle(fmt.Sprintf("Rental Agreement %d - Statement", RAID))
	//tbl.SetSection1(fmt.Sprintf("Statement for: Rental Agreement %d", RAID))
	ra, err := rlib.GetRentalAgreement(ctx, RAID)
	if err != nil {
		tbl.SetSection3("Unable to get Rental Agreement info: " + err.Error())
		return tbl
	}
	sap, err := ra.GetPayorNameList(ctx, &ra.AgreementStart, &ra.AgreementStop)
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}
	tbl.SetSection1(fmt.Sprintf("Statement Period: %s - %s <br>\n%s", d1.Format(rlib.RRDATEFMT3), d2.Format(rlib.RRDATEFMT3), strings.Join(sap, ", ")))

	//--------------------------------------------
	// Set the opening balance.
	//--------------------------------------------
	var b, c, d float64
	b = m.OpeningBal

	tbl.AddRow()
	tbl.Putd(-1, Date, m.DtStart)
	tbl.Puts(-1, ID, "")
	tbl.Puts(-1, Rentable, "")
	tbl.Puts(-1, Description, "Opening Balance")
	tbl.Putf(-1, Assessment, c)
	tbl.Putf(-1, AppliedFunds, d)
	tbl.Putf(-1, Balance, m.ClosingBal)

	if len(m.Stmt) == 0 {
		tbl.AddRow()
		tbl.Puts(-1, Description, "No receipts this period")
	} else {
		for i := 0; i < len(m.Stmt); i++ {
			tbl.AddRow()

			//---------------------------------------------------------------------
			// There are two things we need from the Account Rules.  First
			// is the name of the rule, which is basically an explanation of the
			// charge or payment.  Second, we find out if we need to negate the
			// number in its usage.  The Negate flag indicates whether the
			// Amount of an Assessment should be negated prior to using it in the
			// context of a credit.
			//---------------------------------------------------------------------
			id := ""
			descr := ""
			if m.Stmt[i].T == 1 || m.Stmt[i].T == 2 {
				if m.Stmt[i].A.ARID > 0 {
					descr = rlib.RRdb.BizTypes[BID].AR[m.Stmt[i].A.ARID].Name
				} else {
					descr = rlib.RRdb.BizTypes[BID].GLAccounts[m.Stmt[i].A.ATypeLID].Name
				}
			}
			switch m.Stmt[i].T {
			case 1: // assessments
				amt := m.Stmt[i].Amt
				id = rlib.IDtoShortString("ASM", m.Stmt[i].A.ASMID)
				if !m.Stmt[i].Reverse {
					c += amt
					b += amt
				} else {
					descr += " (" + m.Stmt[i].A.Comment + ")"
				}

			case 2: // receipts
				amt := m.Stmt[i].Amt
				id = rlib.IDtoShortString("RCPT", m.Stmt[i].R.RCPTID)
				if m.Stmt[i].A.ASMID > 0 {
					descr += " (" + rlib.IDtoShortString("ASM", m.Stmt[i].A.ASMID) + ")"
				}
				if !m.Stmt[i].Reverse {
					d += amt
					b -= amt
				} else {
					rcpt, err := rlib.GetReceipt(ctx, m.Stmt[i].R.RCPTID)
					if err != nil {
						tbl.SetSection3(err.Error())
						return tbl
					}

					comment := ""
					if rcpt.RCPTID > 0 {
						comment += rcpt.Comment
					}
					descr += " (" + comment + ")"
				}
			}

			tbl.Putd(-1, Date, m.Stmt[i].Dt)
			tbl.Puts(-1, ID, id)
			tbl.Puts(-1, Rentable, m.Stmt[i].RNT.RentableName)
			tbl.Puts(-1, Description, descr)
			tbl.Putf(-1, Assessment, c)
			tbl.Putf(-1, AppliedFunds, d)
			tbl.Putf(-1, Balance, b)
		}
	}

	// closing balance
	tbl.AddRow()
	tbl.Putd(-1, Date, m.DtStop.AddDate(0, 0, -1))
	tbl.Puts(-1, ID, "")
	tbl.Puts(-1, Rentable, "")
	tbl.Puts(-1, Description, "Closing Balance")
	tbl.Putf(-1, Assessment, c)
	tbl.Putf(-1, AppliedFunds, d)
	tbl.Putf(-1, Balance, m.ClosingBal)

	return tbl
}

// RRRentalAgreementStatements gives string representation of table
func RRRentalAgreementStatements(ctx context.Context, ri *ReporterInfo) gotable.Table {
	// find raid from query params
	var (
		raid int64
	)
	raidStr := ri.QueryParams.Get("raid")
	raid, _ = strconv.ParseInt(raidStr, 10, 64)

	tbl := RRRentalAgreementStatementTable(ctx, ri.Bid, raid, &ri.D1, &ri.D2)
	return tbl
}
