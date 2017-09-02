package main

import (
	"fmt"
	"gotable"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strings"
	"time"
)

// RRPHreport et al are categorizations of commands
const (
	RRPHrpt = 0
	RRPHcmd = iota
	RRPHcsv = iota
	RRPHadm = iota
	RRPHnon = iota // suppress this button
)

// RunCommandLine runs a series of commands to handle command line run requests
func RunCommandLine(ctx *DispatchCtx) {
	rlib.InitBizInternals(ctx.xbiz.P.BID, &ctx.xbiz)
	rcsv.InitRCSV(&ctx.DtStart, &ctx.DtStop, &ctx.xbiz)
	var ri = rrpt.ReporterInfo{OutputFormat: gotable.TABLEOUTTEXT, Bid: ctx.xbiz.P.BID, D1: ctx.DtStart, D2: ctx.DtStop, Xbiz: &ctx.xbiz, BlankLineAfterRptName: true}

	switch ctx.Report {
	case 1: // JOURNAL
		// JournalReportText(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		tbl := rrpt.JournalReport(&ri)
		fmt.Print(tbl)

	case 2: // LEDGER
		// LedgerReportText(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		m := rrpt.LedgerReportTable(&ri)
		for i := 0; i < len(m); i++ {
			fmt.Print(m[i])
			fmt.Printf("\n\n")
		}
	case 3: // INTERNAL ACCT RULE TEST
		intTest(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	case 4: // RENTROLL REPORT
		rrpt.RentRollTextReport(&ri)
	case 6: // available
	case 7: // RENTABLE COUNT BY TYPE
		t := rrpt.RentableCountByRentableTypeReportTable(&ri)
		fmt.Print(t.String())
	case 8: // STATEMENT
		fmt.Print(rrpt.RptStatementTextReport(&ri))
	case 9: // Invoice
		// ctx.Report format:  9,IN0001  or  9,1   -- both say that we want Invoice 1 to be printed
		sa := strings.Split(ctx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing InvoiceNo on report option.  Example:  -r 9,IN000001\n")
			os.Exit(1)
		}
		invoiceno := rcsv.CSVLoaderGetInvoiceNo(sa[1])
		rrpt.InvoiceTextReport(invoiceno)
	case 10: // LEDGER ACTIVITY
		m := rrpt.LedgerActivityReportTable(&ri)
		for i := 0; i < len(m); i++ {
			fmt.Print(m[i])
			fmt.Printf("\n\n")
		}
	case 11: // RENTABLE GSR
		rrpt.GSRTextReport(&ri)
	case 12: // LEDGERBALANCE ON DATE
		// ctx.Report format:  12,LID,RAID,date
		sa := strings.Split(ctx.Args, ",")
		if len(sa) < 4 {
			fmt.Printf("Missing one or more parameters.  Example:  -r 12,L004,RA003,2016-07-04\n")
			os.Exit(1)
		}
		lid := rcsv.CSVLoaderGetLedgerNo(sa[1])
		raid := rcsv.CSVLoaderGetRAID(sa[2])
		dt, err := rlib.StringToDate(sa[3])
		if err != nil {
			fmt.Printf("Bad date string: %s\n", sa[3])
			os.Exit(1)
		}
		rrpt.LdgAcctBalOnDateTextReport(&ctx.xbiz, lid, raid, &dt)
	case 13: // RA LEDGER DETAILS OVER RANGE
		// ctx.Report format: 13,LID,RAID
		// date range is from -j , -k
		sa := strings.Split(ctx.Args, ",")
		if len(sa) < 3 {
			fmt.Printf("Missing one or more parameters.  Example:  -r 13,L004,RA003\n")
			os.Exit(1)
		}
		lid := rcsv.CSVLoaderGetLedgerNo(sa[1])
		raid := rcsv.CSVLoaderGetRAID(sa[2])
		rrpt.RAAccountActivityRangeDetail(&ctx.xbiz, lid, raid, &ctx.DtStart, &ctx.DtStop)
	case 14: // DELINQUENCY REPORT
		// ctx.Report format:  14,date
		sa := strings.Split(ctx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing one or more parameters.  Example:  -r 14,2016-05-25\n")
			os.Exit(1)
		}
		dt, err := rlib.StringToDate(sa[1])
		if err != nil {
			fmt.Printf("Bad date string: %s\n", sa[1])
			os.Exit(1)
		}
		ri.D2 = dt
		rrpt.DelinquencyTextReport(&ri)
	case 15: // Process Vacancy...
		rlib.GenVacancyJournals(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	case 16: // Process LedgerMarkers Only
		rlib.GenerateLedgerMarkers(&ctx.xbiz, &ctx.DtStop)
	case 17: // LEDGER BALANCE REPORT
		rrpt.PrintLedgerBalanceReport(&ri)
	case 18: // Process Journal Entries only
		rlib.GenerateJournalRecords(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop, App.SkipVacCheck)
	case 19: // process Ledgers
		rlib.GenerateLedgerEntries(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	case 20: // List market rates for rentable over time period
		// ctx.Report format:  20,RID
		sa := strings.Split(ctx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing parameter(s).  Example:  -r 20,R004\n")
			os.Exit(1)
		}
		rid := rcsv.CSVLoaderGetRID(sa[1])
		rrpt.RentableMarketRates(&ctx.xbiz, rid, &ctx.DtStart, &ctx.DtStop)
	case 21: // backup file list
		fmt.Print(CreateDBBackupFileList())
	case 22: // delete business
		ri := rrpt.ReporterInfo{Xbiz: &ctx.xbiz, OutputFormat: gotable.TABLEOUTTEXT}
		rrpt.RRreportBusiness(&ri)
		fmt.Printf("Deleting business: %d\n", ctx.xbiz.P.BID)
		rlib.DeleteBusinessFromDB(ctx.xbiz.P.BID)
	case 23: // payor statement
		// ctx.Report format:  23,TCID
		sa := strings.Split(ctx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing one or more parameters.  Example:  -r 23,35\n")
			os.Exit(1)
		}
		tcid, ok := rlib.StringToInt64(sa[1])
		if !ok {
			fmt.Printf("Bad number: %s\n", sa[1])
		}
		tbl := PayorStatement(ctx.xbiz.P.BID, tcid, &ctx.DtStart, &ctx.DtStop)
		s, err := tbl.SprintTable()
		if err != nil {
			rlib.LogAndPrintError("RunCommandLine", err)
			return
		}
		fmt.Print(s)

	default:
		rlib.GenerateJournalRecords(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop, App.SkipVacCheck)
		rlib.GenerateLedgerEntries(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	}
}

// PayorStatement builds a statement for a Payor for a time period
func PayorStatement(bid, tcid int64, d1, d2 *time.Time) gotable.Table {
	var t gotable.Table
	var xbiz rlib.XBusiness

	const (
		Date           = 0
		Payor          = iota
		Description    = iota
		RAID           = iota
		ASMID          = iota
		RCPTID         = iota
		Rentable       = iota
		UnappliedFunds = iota
		AppliedFunds   = iota
		Assessment     = iota
		Balance        = iota
	)

	//
	// UGH!
	//=======================================================================
	rlib.InitBizInternals(bid, &xbiz)
	rlib.Console("bid = %d\n", bid)
	_, ok := rlib.RRdb.BizTypes[bid]
	if !ok {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d]", bid)
		t.SetSection3(e.Error())
		return t
	}
	if len(rlib.RRdb.BizTypes[bid].GLAccounts) == 0 {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d].GLAccounts", bid)
		t.SetSection3(e.Error())
		return t
	}
	//=======================================================================
	// UGH!

	t.Init()
	t.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Payor(s)", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Description", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("RAID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("ASMID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("RCPTID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Rentable", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Unapplied Funds", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Applied Funds", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Assessment", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)

	t.SetTitle("Payor Statement\n\n")
	payors := []int64{tcid}
	m, err := rlib.PayorsStatement(bid, payors, d1, d2)
	if err != nil {
		t.SetSection3("Error from PayorsStatement: " + err.Error())
		return t
	}

	payorcache := map[int64]rlib.Transactant{}

	//------------------------
	// Generate the table
	//------------------------
	for i := 0; i < len(m); i++ { // for each RA

		raidstr := rlib.IDtoShortString("RA", m[i].RAID)
		ra, err := rlib.GetRentalAgreement(m[i].RAID)
		if err != nil {
			rlib.LogAndPrintError("PayorStatement", err)
			continue
		}
		rentableName := ra.GetTheRentableName(d1, d2)
		t.AddRow()
		t.Puts(-1, Description, fmt.Sprintf("*** RENTAL AGREEMENT %d ***", m[i].RAID))
		t.AddRow()
		t.Puts(-1, Description, "Opening balance")
		t.Putd(-1, Date, m[i].DtStart)
		t.Putf(-1, Balance, m[i].OpeningBal)

		//------------------------
		// init running totals
		//------------------------
		bal := m[i].OpeningBal
		asmts := float64(0)
		applied := asmts
		// unapplied := asmts

		for j := 0; j < len(m[i].Stmt); j++ { // for each line in the statement
			t.AddRow()
			t.Putd(-1, Date, m[i].Stmt[j].Dt)
			t.Puts(-1, RAID, raidstr)
			t.Puts(-1, Rentable, rentableName)
			if m[i].Stmt[j].TCID > 0 {
				t.Puts(-1, Payor, rlib.IDtoShortString("TC", m[i].Stmt[j].TCID))
			}

			descr := ""
			if m[i].Stmt[j].Reverse {
				descr += "REVERSAL: "
			}
			amt := m[i].Stmt[j].Amt

			switch m[i].Stmt[j].T {
			case 1: // assessments
				t.Putf(-1, Assessment, amt)
				if m[i].Stmt[j].A.ARID > 0 { // The description will be the name of the Account Rule...
					descr += rlib.RRdb.BizTypes[bid].AR[m[i].Stmt[j].A.ARID].Name
				} else {
					descr += rlib.RRdb.BizTypes[bid].GLAccounts[m[i].Stmt[j].A.ATypeLID].Name
				}
				if m[i].Stmt[j].A.ASMID > 0 {
					t.Puts(-1, ASMID, rlib.IDtoShortString("ASM", m[i].Stmt[j].A.ASMID))
				}
				if m[i].Stmt[j].A.RAID > 0 { // Payor(s) = all payors associated with RentalAgreement
					pyrs := rlib.GetRentalAgreementPayorsInRange(m[i].Stmt[j].A.RAID, d1, d2)
					sa := []string{}
					for k := 0; k < len(pyrs); k++ {
						sa = append(sa, getNameFromCache(pyrs[k].TCID, payorcache))
					}
					t.Puts(-1, Payor, strings.Join(sa, ","))
				}
				if !m[i].Stmt[j].Reverse { // update running totals if not a reversal
					asmts += amt
					bal += amt
				} else {
					descr += " (" + m[i].Stmt[j].A.Comment + ")"
				}
			case 2: // receipts
				t.Putf(-1, AppliedFunds, amt)
				rcptid := m[i].Stmt[j].R.RCPTID
				descr += "Receipt allocation"
				if rcptid > 0 {
					t.Puts(-1, RCPTID, rlib.IDtoShortString("RCPT", rcptid))
					rcpt := rlib.GetReceipt(rcptid)
					if rcpt.RCPTID > 0 {
						name := getNameFromCache(rcpt.TCID, payorcache)
						t.Puts(-1, Payor, name)
					}
				}
				if m[i].Stmt[j].A.ASMID > 0 {
					t.Puts(-1, ASMID, rlib.IDtoShortString("ASM", m[i].Stmt[j].A.ASMID))
				}
				if !m[i].Stmt[j].Reverse {
					applied += amt
					bal -= amt
				} else {
					rcpt := rlib.GetReceipt(m[i].Stmt[j].R.RCPTID)
					if rcpt.RCPTID > 0 && len(rcpt.Comment) > 0 {
						descr += " (" + rcpt.Comment + ")"
					}
				}
			}
			t.Putf(-1, Balance, bal)
			t.Puts(-1, Description, descr)
		}
		t.AddLineAfter(len(t.Row) - 1)
		t.AddRow()
		t.Putd(-1, Date, m[i].DtStop)
		t.Puts(-1, Description, "Closing balance")
		t.Putf(-1, AppliedFunds, applied)
		t.Putf(-1, Assessment, asmts)
		t.Putf(-1, Balance, m[i].ClosingBal)
		t.AddRow()
	}
	return t
}

func getNameFromCache(tcid int64, payorcache map[int64]rlib.Transactant) string {
	p, ok := payorcache[tcid]
	if ok {
		return p.GetUserName()
	} else {
		var tr rlib.Transactant
		err := rlib.GetTransactant(tcid, &tr)
		if err == nil {
			payorcache[tr.TCID] = tr
			return tr.GetUserName()
		}
	}
	return ""
}
