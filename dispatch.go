package main

import (
	"context"
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
func RunCommandLine(ctx context.Context, dCtx *DispatchCtx) {
	err := rlib.InitBizInternals(dCtx.xbiz.P.BID, &dCtx.xbiz)
	if err != nil {
		fmt.Printf("Error while InitBizInternals: %s \n", err.Error())
		os.Exit(1)
	}
	rcsv.InitRCSV(&dCtx.DtStart, &dCtx.DtStop, &dCtx.xbiz)
	var ri = rrpt.ReporterInfo{
		OutputFormat: gotable.TABLEOUTTEXT,
		Bid:          dCtx.xbiz.P.BID,
		D1:           dCtx.DtStart,
		D2:           dCtx.DtStop,
		Xbiz:         &dCtx.xbiz,
		BlankLineAfterRptName: true,
	}

	switch dCtx.Report {
	case 1: // JOURNAL
		// JournalReportText(&dCtx.xbiz, &dCtx.DtStart, &dCtx.DtStop)
		tbl := rrpt.JournalReport(ctx, &ri)
		fmt.Print(tbl)

	case 2: // LEDGER
		// LedgerReportText(&dCtx.xbiz, &dCtx.DtStart, &dCtx.DtStop)
		m, _ := rrpt.LedgerReportTable(ctx, &ri)
		for i := 0; i < len(m); i++ {
			fmt.Print(m[i])
			fmt.Printf("\n\n")
		}
	case 3: // INTERNAL ACCT RULE TEST
		intTest(ctx, &dCtx.xbiz, &dCtx.DtStart, &dCtx.DtStop)
	case 4: // RENTROLL REPORT
		rrpt.RRTextReport(ctx, &ri)
	case 6: // available
	case 7: // RENTABLE COUNT BY TYPE
		t := rrpt.RentableCountByRentableTypeReportTable(ctx, &ri)
		fmt.Print(t.String())
	case 8: // STATEMENT
		fmt.Print(rrpt.RptStatementTextReport(ctx, &ri))
	case 9: // Invoice
		// dCtx.Report format:  9,IN0001  or  9,1   -- both say that we want Invoice 1 to be printed
		sa := strings.Split(dCtx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing InvoiceNo on report option.  Example:  -r 9,IN000001\n")
			os.Exit(1)
		}
		invoiceno := rcsv.CSVLoaderGetInvoiceNo(sa[1])
		rrpt.InvoiceTextReport(ctx, invoiceno)
	case 10: // LEDGER ACTIVITY
		m, _ := rrpt.LedgerActivityReportTable(ctx, &ri)
		for i := 0; i < len(m); i++ {
			fmt.Print(m[i])
			fmt.Printf("\n\n")
		}
	case 11: // RENTABLE GSR
		rrpt.GSRTextReport(ctx, &ri)
	case 12: // LEDGERBALANCE ON DATE
		// dCtx.Report format:  12,LID,RAID,date
		sa := strings.Split(dCtx.Args, ",")
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
		rrpt.LdgAcctBalOnDateTextReport(ctx, &dCtx.xbiz, lid, raid, &dt)
	case 13: // RA LEDGER DETAILS OVER RANGE
		// dCtx.Report format: 13,LID,RAID
		// date range is from -j , -k
		sa := strings.Split(dCtx.Args, ",")
		if len(sa) < 3 {
			fmt.Printf("Missing one or more parameters.  Example:  -r 13,L004,RA003\n")
			os.Exit(1)
		}
		lid := rcsv.CSVLoaderGetLedgerNo(sa[1])
		raid := rcsv.CSVLoaderGetRAID(sa[2])
		rrpt.RAAccountActivityRangeDetail(ctx, &dCtx.xbiz, lid, raid, &dCtx.DtStart, &dCtx.DtStop)
	case 14: // DELINQUENCY REPORT
		// dCtx.Report format:  14,date
		sa := strings.Split(dCtx.Args, ",")
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
		rrpt.DelinquencyTextReport(ctx, &ri)
	case 15: // Process Vacancy...
		rlib.GenVacancyJournals(ctx, &dCtx.xbiz, &dCtx.DtStart, &dCtx.DtStop)
	case 16: // Process LedgerMarkers Only
		rlib.GenerateLedgerMarkers(ctx, &dCtx.xbiz, &dCtx.DtStop)
	case 17: // LEDGER BALANCE REPORT
		rrpt.PrintLedgerBalanceReport(ctx, &ri)
	case 18: // Process Journal Entries only
		rlib.GenerateJournalRecords(ctx, &dCtx.xbiz, &dCtx.DtStart, &dCtx.DtStop, App.SkipVacCheck)
	case 19: // process Ledgers
		rlib.GenerateLedgerEntries(ctx, &dCtx.xbiz, &dCtx.DtStart, &dCtx.DtStop)
	case 20: // List market rates for rentable over time period
		// dCtx.Report format:  20,RID
		sa := strings.Split(dCtx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing parameter(s).  Example:  -r 20,R004\n")
			os.Exit(1)
		}
		rid := rcsv.CSVLoaderGetRID(sa[1])
		rrpt.RentableMarketRates(ctx, &dCtx.xbiz, rid, &dCtx.DtStart, &dCtx.DtStop)
	// case 21: // backup file list
	// 	fmt.Print(CreateDBBackupFileList())
	case 22: // delete business
		ri := rrpt.ReporterInfo{
			Xbiz:         &dCtx.xbiz,
			OutputFormat: gotable.TABLEOUTTEXT,
		}
		rrpt.RRreportBusiness(ctx, &ri)
		fmt.Printf("Deleting business: %d\n", dCtx.xbiz.P.BID)
		rlib.DeleteBusinessFromDB(ctx, dCtx.xbiz.P.BID)
	case 23: // payor statement internal view
		// dCtx.Report format:  23,TCID
		sa := strings.Split(dCtx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing one or more parameters.  Example:  -r 23,35\n")
			os.Exit(1)
		}
		tcid, ok := rlib.StringToInt64(sa[1])
		if !ok {
			fmt.Printf("Bad number: %s\n", sa[1])
		}
		tbl := rrpt.PayorStatement(ctx, dCtx.xbiz.P.BID, tcid, &dCtx.DtStart, &dCtx.DtStop, true)
		s, err := tbl.SprintTable()
		if err != nil {
			rlib.LogAndPrintError("RunCommandLine", err)
			os.Exit(1)
		}
		fmt.Print(s)
	case 24:
		// Assessments
		rlib.Console("Print Assessment report. BID = %d, dtStart = %s, dtStop = %s\n", ri.Bid, ri.D1.Format(rlib.RRDATEREPORTFMT), ri.D2.Format(rlib.RRDATEREPORTFMT))
		fmt.Println(rrpt.RRreportAssessments(ctx, &ri))
	case 25:
		// print task list.   format  -r 25,TLID
		sa := strings.Split(dCtx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing one or more parameters.  Example:  -r 23,35\n")
			os.Exit(1)
		}
		tlid, ok := rlib.StringToInt64(sa[1])
		if !ok {
			fmt.Printf("Bad number: %s\n", sa[1])
			os.Exit(1)
		}
		ri.ID = tlid
		rrpt.TaskListTextReport(ctx, &ri)

	default:
		err := rlib.GenerateJournalRecords(ctx, &dCtx.xbiz, &dCtx.DtStart, &dCtx.DtStop, App.SkipVacCheck)
		if err != nil {
			rlib.DebugPrint("Error from GenerateJournalRecords: %s\n", err.Error())
			rlib.LogAndPrintError("RunCommandLine", err)
			return
		}

		_, err = rlib.GenerateLedgerEntries(ctx, &dCtx.xbiz, &dCtx.DtStart, &dCtx.DtStop)
		if err != nil {
			rlib.DebugPrint("Error from GenerateLedgerEntries: %s\n", err.Error())
			rlib.LogAndPrintError("RunCommandLine", err)
			return
		}
	}
}
