package main

import (
	"fmt"
	"net/http"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strings"
	"text/template"
)

// InitBizInternals initializes several internal structures with information about the business.
// xbiz.P.BID must be set to the BID of interest before calling this function
func InitBizInternals(bid int64, xbiz *rlib.XBusiness) {
	rlib.GetXBusiness(bid, xbiz) // get its info
	rlib.InitBusinessFields(bid)
	rlib.GetDefaultLedgers(bid) // Gather its chart of accounts
	rlib.RRdb.BizTypes[bid].GLAccounts = rlib.GetGLAccountMap(bid)
	rlib.GetAllNoteTypes(bid)
	rlib.LoadRentableTypeCustomaAttributes(xbiz)
}

// RunBooks runs a series of commands to handle command line run requests
func RunBooks(ctx *DispatchCtx) {
	var err error

	InitBizInternals(ctx.xbiz.P.BID, &ctx.xbiz)
	rcsv.InitRCSV(&ctx.DtStart, &ctx.DtStop, &ctx.xbiz)
	// fmt.Printf("RunBooks: Rcsv.Xbiz = %#v\n", rcsv.Rcsv.Xbiz)

	// first handle requests for a CSV Load...
	// if len(ctx.CSVLoadStr) > 0 {
	// 	// fmt.Printf("CSVLoadStr = %s\n", ctx.CSVLoadStr)
	// 	ss := strings.Split(ctx.CSVLoadStr, ",") // index,fname
	// 	if len(ss) < 2 {
	// 		fmt.Printf("Invalid CSVLoader Request:  %s.  Need index,filename\n", ctx.CSVLoadStr)
	// 		os.Exit(1)
	// 	}
	// 	i, err := strconv.Atoi(ss[0])
	// 	if err != nil {
	// 		fmt.Printf("Invalid CSVLoaderIndex: %s.  err: %s\n", ss[0], err.Error())
	// 		os.Exit(1)
	// 	}
	// 	// fmt.Printf("calling LoadCSV( %d , %s )\n", i, ss[1])
	// 	rcsv.LoadCSV(i, ss[1])
	// 	return
	// }

	// and generate the requested report...
	switch ctx.Report {
	case 1: // JOURNAL
		// JournalReportText(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		tbl := rrpt.JournalReport(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		fmt.Print(tbl)

	case 2: // LEDGER
		// LedgerReportText(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		m := rrpt.LedgerReport(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		for i := 0; i < len(m); i++ {
			fmt.Print(m[i])
			fmt.Printf("\n\n")
		}
	case 3: // INTERNAL ACCT RULE TEST
		intTest(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	case 4: // RENTROLL REPORT
		err = rrpt.RentRollTextReport(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		if err != nil {
			fmt.Printf("RentRoll text report error: %s\n", err.Error())
		}
	case 5: // ASSESSMENT CHECK REPORT
		AssessmentCheckReportText(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	case 6: // available
	case 7: // RENTABLE COUNT BY TYPE
		UIRentableCountByRentableTypeReport(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	case 8: // STATEMENT
		UIStatementTextReport(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
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
		// LedgerActivityReport(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		m := rrpt.LedgerActivityReport(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
		for i := 0; i < len(m); i++ {
			fmt.Print(m[i])
			fmt.Printf("\n\n")
		}
	case 11: // RENTABLE GSR
		rrpt.GSRTextReport(&ctx.xbiz, &ctx.DtStart)
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
		err = rrpt.DelinquencyTextReport(&ctx.xbiz, &dt)
		if err != nil {
			fmt.Printf("Delinquency text report error: %s\n", err.Error())
		}
	case 15: // Process Vacancy...
		rlib.GenVacancyJournals(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	case 16: // Process LedgerMarkers Only
		rlib.GenerateLedgerMarkers(&ctx.xbiz, &ctx.DtStop)
	case 17: // LEDGER BALANCE REPORT
		rrpt.PrintLedgerBalanceReport(&ctx.xbiz, &ctx.DtStop)
	case 18: // Process Journal Entries only
		rlib.GenerateJournalRecords(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop, App.SkipVacCheck)
	case 19: // process Ledgers
		rlib.GenerateLedgerRecords(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	case 20: // List market rates for rentable over time period
		// ctx.Report format:  20,RID
		sa := strings.Split(ctx.Args, ",")
		if len(sa) < 2 {
			fmt.Printf("Missing parameter(s).  Example:  -r 20,R004\n")
			os.Exit(1)
		}
		rid := rcsv.CSVLoaderGetRID(sa[1])
		rrpt.RentableMarketRates(&ctx.xbiz, rid, &ctx.DtStart, &ctx.DtStop)
	default:
		rlib.GenerateJournalRecords(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop, App.SkipVacCheck)
		rlib.GenerateLedgerRecords(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	}
}

// Dispatch generates the supplied report for all properties
func Dispatch(ctx *DispatchCtx) {
	switch ctx.Cmd {
	case CmdRUNBOOKS:
		RunBooks(ctx)
	default:
		fmt.Printf("Unrecognized command: %d\n", ctx.Cmd)
	}
}

func dispatchHandler(w http.ResponseWriter, r *http.Request) {
	funcname := "dispatchHandler"
	tmpl := "dispatch.html"
	var ui RRuiSupport
	var xbiz rlib.XBusiness

	getBizStartStop(w, r, &ui, &xbiz)
	ui.PgHnd = App.PageHandlers
	action := r.FormValue("action")

	// fmt.Printf("action = %s\n", action)

	if action == "Assessments" {
		CmdCsvAssess(w, r, &xbiz, &ui, &tmpl)
	} else if len(action) > 0 && action != "Home" {
		w.Header().Set("Content-Type", "text/html")
		for i := 0; i < len(App.PageHandlers); i++ {
			if action == App.PageHandlers[i].ReportName {
				// fmt.Printf("dispatchHandler:  calling handler for %s\n", App.PageHandlers[i].ReportName)
				App.PageHandlers[i].Handler(w, r, &xbiz, &ui, &tmpl)
				break
			}
		}
	}
	t, err := template.New(tmpl).Funcs(RRfuncMap).ParseFiles("./html/" + tmpl)
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}
	err = t.Execute(w, &ui)

	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
