package main

import (
	"fmt"
	"gotable"
	"html/template"
	"net/http"
	"net/url"
	"rentroll/rlib"
	"rentroll/rrpt"
	"rentroll/ws"
	"strings"
	"time"
)

// SendWebSvcPage creates an HTML page where the content is simply
// the contents of ui.ReportContent formatted to be shown as-is in
// a mono-spaced font.
func SendWebSvcPage(w http.ResponseWriter, r *http.Request, ui *RRuiSupport) {
	funcname := "SendWebSvcPage"
	tmpl := "v1rpt.html"
	t, err := template.New(tmpl).Funcs(RRfuncMap).ParseFiles("./html/" + tmpl)
	if nil != err {
		s := fmt.Sprintf("%s: error loading template: %v\n", funcname, err)
		ui.ReportContent += s
		fmt.Print(s)
	}
	err = t.Execute(w, &ui)

	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func v1ReportHandler(reportname string, xbiz *rlib.XBusiness, ui *RRuiSupport, w http.ResponseWriter) {
	funcname := "v1ReportHandler"
	fmt.Printf("%s: reportname=%s, BID=%d,  d1 = %s, d2 = %s\n", funcname, reportname, xbiz.P.BID, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))

	var ri = rrpt.ReporterInfo{OutputFormat: gotable.TABLEOUTHTML, Bid: xbiz.P.BID, D1: ui.D1, D2: ui.D2, Xbiz: xbiz, BlankLineAfterRptName: true}
	rlib.InitBizInternals(ri.Bid, xbiz)

	// handler for reports which has single table
	var wsr = []rrpt.SingleTableReportHandler{
		{ReportNames: []string{"asmrpt", "assessments"}, TableHandler: rrpt.RRAssessmentsTable},
		{ReportNames: []string{"b", "business"}, TableHandler: rrpt.RRreportBusinessTable},
		{ReportNames: []string{"coa", "chart of accounts"}, TableHandler: rrpt.RRreportChartOfAccountsTable},
		{ReportNames: []string{"c", "custom attributes"}, TableHandler: rrpt.RRreportCustomAttributesTable},
		{ReportNames: []string{"cr", "custom attribute refs"}, TableHandler: rrpt.RRreportCustomAttributeRefsTable},
		{ReportNames: []string{"delinq", "delinquency"}, TableHandler: rrpt.DelinquencyReportTable},
		{ReportNames: []string{"dpm", "deposit methods"}, TableHandler: rrpt.RRreportDepositMethodsTable},
		{ReportNames: []string{"dep", "depositories"}, TableHandler: rrpt.RRreportDepositoryTable},
		{ReportNames: []string{"gsr"}, TableHandler: rrpt.GSRReportTable},
		{ReportNames: []string{"j"}, TableHandler: rrpt.JournalReportTable},
		{ReportNames: []string{"pmt", "payment types"}, TableHandler: rrpt.RRreportPaymentTypesTable},
		{ReportNames: []string{"r", "rentables"}, TableHandler: rrpt.RRreportRentablesTable},
		{ReportNames: []string{"ra", "rental agreements"}, TableHandler: rrpt.RRreportRentalAgreementsTable},
		{ReportNames: []string{"rat", "rental agreement templates"}, TableHandler: rrpt.RRreportRentalAgreementTemplatesTable},
		{ReportNames: []string{"rcpt", "receipts"}, TableHandler: rrpt.RRReceiptsTable},
		{ReportNames: []string{"rr", "rentroll"}, TableHandler: rrpt.RentRollReportTable},
		{ReportNames: []string{"rt", "rentable types"}, TableHandler: rrpt.RRreportRentableTypesTable},
		{ReportNames: []string{"rcbt", "rentable type counts"}, TableHandler: rrpt.RentableCountByRentableTypeReportTable},
		{ReportNames: []string{"sl", "string lists"}, TableHandler: rrpt.RRreportStringListsTable},
		{ReportNames: []string{"t", "people"}, TableHandler: rrpt.RRreportPeopleTable},
		{ReportNames: []string{"tb", "trial balance"}, TableHandler: rrpt.LedgerBalanceReportTable},
	}

	// handler for reports which has more than one table
	var wmr = []rrpt.MultiTableReportHandler{
		{ReportNames: []string{"l", "ledger"}, TableHandler: rrpt.LedgerReportTable},
		{ReportNames: []string{"la", "ledger activity"}, TableHandler: rrpt.LedgerActivityReportTable},
		{ReportNames: []string{"statements"}, TableHandler: rrpt.RptStatementReportTable},
	}

	// find reportname from list of report handler
	// first find it from single table handler
	var tsh rrpt.SingleTableReportHandler
	for j := 0; j < len(wsr); j++ {
		for _, rn := range wsr[j].ReportNames {
			if rn == strings.ToLower(reportname) {
				tsh = wsr[j]
				tsh.Found = true
				// if found then stop looking over report names list
				break
			}
		}
		// if found then stop looping over handlers
		if tsh.Found {
			break
		}
	}

	// if found then handle service for request
	if tsh.Found {
		tbl := tsh.TableHandler(&ri)

		switch ui.ReportOutputFormat {
		case gotable.TABLEOUTTEXT:
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Disposition", "attachment; filename="+reportname+".text")
			err := tbl.TextprintTable(w)
			if err != nil {
				s := fmt.Sprintf("Error in TextprintTable: %s\n", err.Error())
				fmt.Print(s)
				fmt.Fprintf(w, "%s\n", s)
			}
			return
		case gotable.TABLEOUTHTML:
			err := tbl.HTMLprintTable(w)
			if err != nil {
				s := fmt.Sprintf("Error in HTMLprintTable: %s\n", err.Error())
				fmt.Print(s)
				fmt.Fprintf(w, "%s\n", s)
			}
			return
		case gotable.TABLEOUTCSV:
			w.Header().Set("Content-Type", "text/csv")
			w.Header().Set("Content-Disposition", "attachment; filename="+reportname+".csv")
			err := tbl.CSVprintTable(w)
			if err != nil {
				s := fmt.Sprintf("Error in CSVprintTable: %s\n", err.Error())
				fmt.Print(s)
				fmt.Fprintf(w, "%s\n", s)
			}
			return
		case gotable.TABLEOUTPDF:
			w.Header().Set("Content-Type", "application/pdf")
			w.Header().Set("Content-Disposition", "attachment; filename="+reportname+".pdf")
			err := tbl.PDFprintTable(w)
			if err != nil {
				s := fmt.Sprintf("Error in PDFprintTable: %s\n", err.Error())
				fmt.Print(s)
				fmt.Fprintf(w, "%s\n", s)
			}
			return
		default:
			fmt.Fprintf(w, "%s", "Unsupported format output of report")
			return
		}
	}

	// if not found from single, then find it from multi table handler
	var tmh rrpt.MultiTableReportHandler
	for j := 0; j < len(wmr); j++ {
		for _, rn := range wmr[j].ReportNames {
			if rn == strings.ToLower(reportname) {
				tmh = wmr[j]
				tmh.Found = true
				// if found then stop looking over report name list
				break
			}
		}
		// if found then stop looking over other handler
		if tmh.Found {
			break
		}
	}

	// if found then handle service for request
	if tmh.Found {
		m := tmh.TableHandler(&ri)

		switch ui.ReportOutputFormat {
		case gotable.TABLEOUTTEXT:
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Disposition", "attachment; filename="+reportname+".text")
			rrpt.MultiTableTextPrint(m, w)
			return
		case gotable.TABLEOUTHTML:
			rrpt.MultiTableHTMLPrint(m, w)
			return
		case gotable.TABLEOUTCSV:
			w.Header().Set("Content-Type", "text/csv")
			w.Header().Set("Content-Disposition", "attachment; filename="+reportname+".csv")
			rrpt.MultiTableCSVPrint(m, w)
			return
		case gotable.TABLEOUTPDF:
			w.Header().Set("Content-Type", "application/pdf")
			w.Header().Set("Content-Disposition", "attachment; filename="+reportname+".pdf")
			rrpt.MultiTablePDFPrint(m, w)
			return
		default:
			fmt.Fprintf(w, "%s", "Unsupported format output of report")
			return
		}
	}

	// unknown report handler
	fmt.Fprintf(w, "Unknown report type: %s", reportname)
	return
}

func websvcReportHandler(reportname string, xbiz *rlib.XBusiness, ui *RRuiSupport) string {
	funcname := "websvcReportHandler"
	fmt.Printf("%s: reportname=%s, BID=%d,  d1 = %s, d2 = %s\n", funcname, reportname, xbiz.P.BID, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))

	var ri = rrpt.ReporterInfo{OutputFormat: gotable.TABLEOUTTEXT, Bid: xbiz.P.BID, D1: ui.D1, D2: ui.D2, Xbiz: xbiz, RptHeader: true, BlankLineAfterRptName: true}
	rlib.InitBizInternals(ri.Bid, xbiz)

	// handler for reports which has single table
	var wsr = []rrpt.SingleTableReportHandler{
		{ReportNames: []string{"asmrpt", "assessments"}, TableHandler: rrpt.RRAssessmentsTable},
		{ReportNames: []string{"b", "business"}, TableHandler: rrpt.RRreportBusinessTable},
		{ReportNames: []string{"coa", "chart of accounts"}, TableHandler: rrpt.RRreportChartOfAccountsTable},
		{ReportNames: []string{"c", "custom attributes"}, TableHandler: rrpt.RRreportCustomAttributesTable},
		{ReportNames: []string{"cr", "custom attribute refs"}, TableHandler: rrpt.RRreportCustomAttributeRefsTable},
		{ReportNames: []string{"delinq", "delinquency"}, TableHandler: rrpt.DelinquencyReportTable},
		{ReportNames: []string{"dpm", "deposit methods"}, TableHandler: rrpt.RRreportDepositMethodsTable},
		{ReportNames: []string{"dep", "depositories"}, TableHandler: rrpt.RRreportDepositoryTable},
		{ReportNames: []string{"gsr"}, TableHandler: rrpt.GSRReportTable},
		{ReportNames: []string{"j"}, TableHandler: rrpt.JournalReportTable},
		{ReportNames: []string{"pmt", "payment types"}, TableHandler: rrpt.RRreportPaymentTypesTable},
		{ReportNames: []string{"r", "rentables"}, TableHandler: rrpt.RRreportRentablesTable},
		{ReportNames: []string{"ra", "rental agreements"}, TableHandler: rrpt.RRreportRentalAgreementsTable},
		{ReportNames: []string{"rat", "rental agreement templates"}, TableHandler: rrpt.RRreportRentalAgreementTemplatesTable},
		{ReportNames: []string{"rcpt", "receipts"}, TableHandler: rrpt.RRReceiptsTable},
		{ReportNames: []string{"rr", "rentroll"}, TableHandler: rrpt.RentRollReportTable},
		{ReportNames: []string{"rt", "rentable types"}, TableHandler: rrpt.RRreportRentableTypesTable},
		{ReportNames: []string{"rcbt", "rentable type counts"}, TableHandler: rrpt.RentableCountByRentableTypeReportTable},
		{ReportNames: []string{"sl", "string lists"}, TableHandler: rrpt.RRreportStringListsTable},
		{ReportNames: []string{"t", "people"}, TableHandler: rrpt.RRreportPeopleTable},
		{ReportNames: []string{"tb", "trial balance"}, TableHandler: rrpt.LedgerBalanceReportTable},
	}

	// handler for reports which has more than one table
	var wmr = []rrpt.MultiTableReportHandler{
		{ReportNames: []string{"l", "ledger"}, TableHandler: rrpt.LedgerReportTable},
		{ReportNames: []string{"la", "ledger activity"}, TableHandler: rrpt.LedgerActivityReportTable},
		{ReportNames: []string{"statements"}, TableHandler: rrpt.RptStatementReportTable},
	}

	// find reportname from list of report handler
	// first find it from single table handler
	var tsh rrpt.SingleTableReportHandler
	for j := 0; j < len(wsr); j++ {
		for _, rn := range wsr[j].ReportNames {
			if rn == strings.ToLower(reportname) {
				tsh = wsr[j]
				tsh.Found = true
				// if found then stop looking over report names list
				break
			}
		}
		// if found then stop looping over handlers
		if tsh.Found {
			break
		}
	}

	// if found then handle service for request
	if tsh.Found {
		tbl := tsh.TableHandler(&ri)
		return rrpt.ReportToString(&tbl, &ri)
	}

	// if not found from single, then find it from multi table handler
	var tmh rrpt.MultiTableReportHandler
	for j := 0; j < len(wmr); j++ {
		for _, rn := range wmr[j].ReportNames {
			if rn == strings.ToLower(reportname) {
				tmh = wmr[j]
				tmh.Found = true
				// if found then stop looking over report name list
				break
			}
		}
		// if found then stop looking over other handler
		if tmh.Found {
			break
		}
	}

	// if found then handle service for request
	if tmh.Found {
		m := tmh.TableHandler(&ri)
		var s string
		// Spin through all the RentalAgreements that are active in this timeframe
		for _, tbl := range m {
			s += rrpt.ReportToString(&tbl, &ri) + "\n"
		}
		return s
	}

	// unknown report handler
	return fmt.Sprintf("Unknown report type: %s", reportname)
}

// webServiceHandler dispatches all the web service requests
// This service handles requests of the form:
//    http://x.y.z/wsvc/<uid>/<BID>?[params]
// where params can be:
//	  r=<reportname>
//    dtstart=<date>
//    dtstop=<date>
//
func webServiceHandler(w http.ResponseWriter, r *http.Request) {
	funcname := "webServiceHandler"
	fmt.Printf("Entered %s\n", funcname)
	var ui RRuiSupport
	var d ws.ServiceData
	var xbiz rlib.XBusiness
	var reportname string
	var err error

	if r.Method == "GET" {
		fmt.Printf("r.RequestURL = %s\n", r.URL.String())
		sa := strings.Split(r.URL.Path, "/") // ["", "wsvc", "<UID>", "<BID>"]
		fmt.Printf("sa = %#v\n", sa)
		d.UID, err = rlib.IntFromString(sa[2], "bad request integer value")
		if err != nil {
			ui.ReportContent = fmt.Sprintf("Error parsing request URI: %s", err.Error())
			SendWebSvcPage(w, r, &ui)
			return
		}
		d.BID, err = rlib.IntFromString(sa[3], "bad request integer value")
		if err != nil {
			ui.ReportContent = fmt.Sprintf("Error parsing request URI: %s", err.Error())
			SendWebSvcPage(w, r, &ui)
			return
		}

		if d.BID > 0 {
			rlib.GetXBusiness(d.BID, &xbiz)
		}

		m, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			ui.ReportContent = fmt.Sprintf("Error parsing request raw query: %s", err.Error())
			SendWebSvcPage(w, r, &ui)
			return
		}
		fmt.Printf("m = %#v\n", m)
		// What report
		x, ok := m["r"] // ?r=<reportname>
		if !ok {
			ui.ReportContent = "no report name found"
			SendWebSvcPage(w, r, &ui)
			return
		}
		reportname = x[0]

		tnow := time.Now()
		ui.D1 = time.Date(tnow.Year(), tnow.Month(), 1, 0, 0, 0, 0, time.UTC)
		nm := int(ui.D1.Month()) + 1
		yr := ui.D1.Year()
		if nm > int(time.December) {
			nm = int(time.January)
			yr++
		}
		ui.D2 = time.Date(yr, time.Month(nm), 1, 0, 0, 0, 0, time.UTC)
		x, ok = m["dtstart"]
		var d1 time.Time
		if ok && len(x[0]) > 0 {
			d1, err = rlib.StringToDate(x[0])
			if err == nil {
				ui.D1 = d1
				// ui.ReportContent = fmt.Sprintf("Error with dtstart value:  %s", err.Error())
				// SendWebSvcPage(w, r, &ui)
				// return
			}
		}
		x, ok = m["dtstop"]
		if ok && len(x[0]) > 0 {
			d1, err = rlib.StringToDate(x[0])
			if err == nil {
				ui.D2 = d1
				// ui.ReportContent = fmt.Sprintf("Error with dtstop value:  %s", err.Error())
				// SendWebSvcPage(w, r, &ui)
				// return
			}
		}
		var rof int // report output format indicator
		x, ok = m["rof"]
		if ok && len(x[0]) > 0 {
			if rof, ok = rlib.StringToInt(x[0]); !ok {
				rof = gotable.TABLEOUTHTML
			}
		} else {
			rof = gotable.TABLEOUTHTML
		}
		ui.ReportOutputFormat = rof
	}

	v1ReportHandler(reportname, &xbiz, &ui, w)
	// ui.ReportContent = websvcReportHandler(reportname, &xbiz, &ui)
	// SendWebSvcPage(w, r, &ui)
}
