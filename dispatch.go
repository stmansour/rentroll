package main

import (
	"fmt"
	"net/http"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// RRPHreport et al are categorizations of commands
const (
	RRPHrpt = 0
	RRPHcmd = iota
	RRPHcsv = iota
	RRPHadm = iota
)

// RRPageHandler is a structure of page names and handlers
type RRPageHandler struct {
	Name     string                                                                  // report name
	PageName string                                                                  // name of form to collect information for this report
	Handler  func(http.ResponseWriter, *http.Request, *rlib.XBusiness, *RRuiSupport) // the actual handler function
	Category int                                                                     // report, command, admin
}

func initPageHandlers() {
	var m = []RRPageHandler{
		// ReportName FormPageName URL
		{"Delinquency", "rptdelinq.html", RptDelinq, RRPHrpt},
		{"GSR", "rptgsr.html", RptGSR, RRPHrpt},
		{"Journal", "rptjournal.html", RptJournal, RRPHrpt},
		{"Ledger", "rptledger.html", RptLedger, RRPHrpt},
		{"Ledger Activity", "rptledgeract.html", RptLedgerActivity, RRPHrpt},
		{"RentRoll", "rptrentroll.html", RptRentRoll, RRPHrpt},
		{"Trial Balance", "rpttrialbal.html", RptTrialBalance, RRPHrpt},
		{"Assessments", "csvassess.html", CmdCsvAssess, RRPHcsv},
		{"Receipts", "csvrcpt.html", CmdCsvRcpt, RRPHcsv},
		{"Generate Journals", "cmdgenjnl.html", CmdGenJnl, RRPHcmd},
		// {"Generate Ledgers", "cmdgenldg.html", CmdGenLdg, RRPHcmd},
		{"Backup", "admbkup.html", AdmBkup, RRPHadm},
		{"Restore", "admrestore.html", AdmRestore, RRPHadm},

		// {ReportName: "Business", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Chart of Accounts", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Custom Attributes", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Delinquency", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Deposits", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Deposit Methods", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Depositories ", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Invoice", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Invoice Report", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Journal", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Ledger Activity", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Market Rate for Rentable", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Note Types", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "People", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Pets", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Payment Types", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Rental Agreements", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Rental Agreement Account Balance", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Rentable Count by Rentable Type", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Rentables", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "RatePlans", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "RatePlanRef", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Rentable Specialty Assignments", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Rentable Types", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Rentable Specialties", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Specialty Assignments", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Sources", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Statements", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Rental Agreement Templates", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Custom Attribute Assignments", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
	}
	for i := 0; i < len(m); i++ {
		App.PageHandlers = append(App.PageHandlers, m[i])
	}

}

// GetUIContext initializes the structures used by the UI based on some common form elements.
func GetUIContext(w http.ResponseWriter, r *http.Request, ui *RRuiSupport, xbiz *rlib.XBusiness) {
	now := time.Now()
	var err error

	// Init to reasonable starting values
	ui.DtStart = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).Format(rlib.RRDATEINPFMT)
	ui.DtStop = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0).Format(rlib.RRDATEINPFMT)

	// override with input if available
	s := r.FormValue("DtStart")
	if s != "" {
		ui.DtStart = s
	}
	s = r.FormValue("DtStop")
	if s != "" {
		ui.DtStop = s
	}
	BID := r.FormValue("BID")

	UIInitBizList(ui)
	ui.D1, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(ui.DtStart))
	if err != nil {
		rlib.Ulog("GetUIContext: could not parse %s, err = %v\n", ui.DtStart, err)
	}
	ui.D2, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(ui.DtStop))
	if err != nil {
		rlib.Ulog("GetUIContext: could not parse %s, err = %v\n", ui.DtStop, err)
	}
	if len(BID) > 0 {
		i, err := strconv.Atoi(strings.TrimSpace(BID))
		if err != nil {
			rlib.Ulog("Error fetching business with BID = %d\n", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rlib.InitBizInternals(int64(i), xbiz)
		UIInitBizList(ui)
	}
}

func dispatchHandler(w http.ResponseWriter, r *http.Request) {
	funcname := "dispatchHandler"
	tmpl := "dispatch.html"
	var ui RRuiSupport
	var xbiz rlib.XBusiness

	GetUIContext(w, r, &ui, &xbiz)
	ui.PgHnd = App.PageHandlers
	action := r.FormValue("action")

	if len(action) > 0 {
		w.Header().Set("Content-Type", "text/html")
		for i := 0; i < len(App.PageHandlers); i++ {
			if action == App.PageHandlers[i].Name && nil != App.PageHandlers[i].Handler {
				App.PageHandlers[i].Handler(w, r, &xbiz, &ui)
				tmpl = App.PageHandlers[i].PageName
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

// RunCommandLine runs a series of commands to handle command line run requests
func RunCommandLine(ctx *DispatchCtx) {
	var err error

	rlib.InitBizInternals(ctx.xbiz.P.BID, &ctx.xbiz)
	rcsv.InitRCSV(&ctx.DtStart, &ctx.DtStop, &ctx.xbiz)

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
		RunCommandLine(ctx)
	default:
		fmt.Printf("Unrecognized command: %d\n", ctx.Cmd)
	}
}
