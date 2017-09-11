package main

import (
	"fmt"
	"gotable"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strings"
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
	case 23: // payor statement internal view
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
		tbl := rrpt.PayorStatement(ctx.xbiz.P.BID, tcid, &ctx.DtStart, &ctx.DtStop, true)
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
