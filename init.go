package main

import (
	"fmt"
	"os"
	"rentroll/rlib"
	"strings"
	"text/template"
)

func initRentRoll() {
	initJFmt()
	initTFmt()
	rlib.RpnInit()

	RRfuncMap = template.FuncMap{
		"DateToString": rlib.DateToString,
		"getVersionNo": getVersionNo,
		"getBuildTime": getBuildTime,
		"RRCommaf":     rlib.RRCommaf,
		"LMSum":        LMSum,
	}
}

func createStartupCtx() DispatchCtx {
	var ctx DispatchCtx
	var err error
	ctx.DtStart, err = rlib.StringToDate(App.sStart)
	if err != nil {
		fmt.Printf("Invalid start date:  %s\n", App.sStart)
		os.Exit(1)
	}
	ctx.DtStop, err = rlib.StringToDate(App.sStop)
	if err != nil {
		fmt.Printf("Invalid start date:  %s\n", App.sStop)
		os.Exit(1)
	}

	des := strings.ToLower(strings.TrimSpace(App.Bud)) // this should be BUD
	if len(des) == 0 {                                 // make sure it's not empty
		fmt.Printf("No BUD specified. A BUD is required for batch mode operation\n")
		os.Exit(1)
	}
	ctx.xbiz.P = rlib.GetBusinessByDesignation(des) // see if we can find the biz
	if len(ctx.xbiz.P.Designation) == 0 {
		rlib.Ulog("Business Unit with designation %s does not exist\n", des)
		os.Exit(1)
	}
	rlib.GetXBusiness(ctx.xbiz.P.BID, &ctx.xbiz)

	// App.Report is a string, of the format:
	//   n[,s1[,s2[...]]]
	// Example:
	//   1           -- this would be for a Journal text report
	//   9,IN0001    -- this would be for a text report of Invoice 1
	//
	// The only required value is n, the report number
	sa := strings.Split(App.Report, ",") // comma separated list
	if len(App.Report) > 0 {
		ctx.Report, _ = rlib.IntFromString(sa[0], "invalid report number")
	}
	ctx.Args = App.Report
	ctx.CSVLoadStr = strings.TrimSpace(App.CSVLoad)
	// fmt.Printf("ctx.CSVLoadStr = %s\n", ctx.CSVLoadStr)
	ctx.Cmd = CmdRUNBOOKS
	ctx.OutputFormat = FMTTEXT
	return ctx
}

func initPageHandlers() {
	var m = []RRPageHandler{
		// ReportName FormPageName URL

		{ReportName: "Trial Balance", FormPageName: "formtrialbal.html", URL: "/trialbalance/", Handler: RptTrialBalance},
		{ReportName: "RentRoll", FormPageName: "rptrentroll.html", URL: "/rptrentroll/", Handler: RptRentRoll},
		{ReportName: "Delinquency", FormPageName: "rptdelinq.html", URL: "/rptdelinq/", Handler: RptDelinq},
		// {ReportName: "Journal", FormPageName: "journal.html", URL: "/rptjournal/", Handler: RptJournal},
		{ReportName: "Home", FormPageName: "dispatch.html", URL: "/dispatch/", Handler: dispatchHandler},

		// {ReportName: "Assessments", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Business", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Chart of Accounts", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Custom Attributes", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Delinquency", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Deposits", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Deposit Methods", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Depositories ", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "GSR", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Invoice", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Invoice Report", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Journal", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Ledger", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Ledger Activity", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Market Rate for Rentable", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Note Types", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "People", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Pets", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Payment Types", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
		// {ReportName: "Receipts", FormPageName: "formtrialbal.html", FormHandler: "/trialbalance/", ReportPageName: "", ReportHandler: "/trialbalance/"},
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
