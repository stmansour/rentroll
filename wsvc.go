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
	t, err := template.New(tmpl).Funcs(RRfuncMap).ParseFiles("./webclient/html/" + tmpl)
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

func v1ReportHandler(reportname string, xbiz *rlib.XBusiness, ui *RRuiSupport, w http.ResponseWriter, qp *url.Values) {
	funcname := "v1ReportHandler"
	rlib.Console("%s: reportname=%s, BID=%d,  d1 = %s, d2 = %s\n", funcname, reportname, xbiz.P.BID, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))

	var ri = rrpt.ReporterInfo{
		OutputFormat: gotable.TABLEOUTHTML,
		Bid:          xbiz.P.BID,
		D1:           ui.D1,
		D2:           ui.D2,
		Xbiz:         xbiz,
		BlankLineAfterRptName: true,
		QueryParams:           qp}

	// init business internals first
	rlib.InitBizInternals(ri.Bid, xbiz)

	// handler for reports which has single table
	var wsr = []rrpt.SingleTableReportHandler{
		{ReportNames: []string{"RPTasmrpt", "assessments"}, TableHandler: rrpt.RRAssessmentsTable},
		{ReportNames: []string{"RPTb", "business"}, TableHandler: rrpt.RRreportBusinessTable},
		{ReportNames: []string{"RPTcoa", "chart of accounts"}, TableHandler: rrpt.RRreportChartOfAccountsTable},
		{ReportNames: []string{"RPTc", "custom attributes"}, TableHandler: rrpt.RRreportCustomAttributesTable},
		{ReportNames: []string{"RPTcr", "custom attribute refs"}, TableHandler: rrpt.RRreportCustomAttributeRefsTable},
		{ReportNames: []string{"RPTdelinq", "delinquency"}, TableHandler: rrpt.DelinquencyReportTable},
		{ReportNames: []string{"RPTdpm", "deposit methods"}, TableHandler: rrpt.RRreportDepositMethodsTable},
		{ReportNames: []string{"RPTdep", "depositories"}, TableHandler: rrpt.RRreportDepositoryTable},
		{ReportNames: []string{"RPTgsr", "gsr"}, TableHandler: rrpt.GSRReportTable},
		{ReportNames: []string{"RPTj", "journals"}, TableHandler: rrpt.JournalReportTable},
		{ReportNames: []string{"RPTpeople", "people"}, TableHandler: rrpt.RRreportPeopleTable},
		{ReportNames: []string{"RPTpmt", "payment types"}, TableHandler: rrpt.RRreportPaymentTypesTable},
		{ReportNames: []string{"RPTr", "rentables"}, TableHandler: rrpt.RRreportRentablesTable},
		{ReportNames: []string{"RPTra", "rental agreements"}, TableHandler: rrpt.RRreportRentalAgreementsTable},
		{ReportNames: []string{"RPTrat", "rental agreement templates"}, TableHandler: rrpt.RRreportRentalAgreementTemplatesTable},
		{ReportNames: []string{"RPTrcpt", "receipts"}, TableHandler: rrpt.RRReceiptsTable},
		// {ReportNames: []string{"RPTrr", "rentroll"}, TableHandler: rrpt.RentRollReportTable},
		{ReportNames: []string{"RPTrr", "rentroll"}, TableHandler: rrpt.RRReportTable},
		{ReportNames: []string{"RPTrt", "rentable types"}, TableHandler: rrpt.RRreportRentableTypesTable},
		{ReportNames: []string{"RPTrcbt", "rentable type counts"}, TableHandler: rrpt.RentableCountByRentableTypeReportTable},
		{ReportNames: []string{"RPTsl", "string lists"}, TableHandler: rrpt.RRreportStringListsTable},
		{ReportNames: []string{"RPTt", "people"}, TableHandler: rrpt.RRreportPeopleTable},
		{ReportNames: []string{"RPTtb", "trial balance"}, TableHandler: rrpt.LedgerBalanceReportTable},
		{ReportNames: []string{"RPTpayorstmt", "payor statements"}, TableHandler: rrpt.RRPayorStatement},
		{ReportNames: []string{"RPTrastmt", "rental agreement statements"}, TableHandler: rrpt.RRRentalAgreementStatements},
	}

	// handler for reports which has more than one table
	var wmr = []rrpt.MultiTableReportHandler{
		{ReportTitle: "Ledger", ReportNames: []string{"RPTl", "ledger"}, TableHandler: rrpt.LedgerReportTable},
		{ReportTitle: "Ledger Activity", ReportNames: []string{"RPTla", "ledger activity"}, TableHandler: rrpt.LedgerActivityReportTable},
		{ReportTitle: "Report Statements", ReportNames: []string{"RPTstatements", "report statements"}, TableHandler: rrpt.RptStatementReportTable},
	}

	// find reportname from list of report handler
	// first find it from single table handler
	var tsh rrpt.SingleTableReportHandler
	for j := 0; j < len(wsr); j++ {
		for _, rn := range wsr[j].ReportNames {
			if strings.ToLower(rn) == strings.ToLower(reportname) {
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

		// format downloadable report name
		attachmentName := ri.Xbiz.P.Designation + "-" + strings.Title(tsh.ReportNames[1])
		if !ri.D1.IsZero() {
			fromDate := rrpt.GetAttachmentDate(ri.D1)
			attachmentName += "-From:" + fromDate
		}
		if !ri.D2.IsZero() {
			toDate := rrpt.GetAttachmentDate(ri.D2)
			attachmentName += "To:" + toDate
		}

		switch ui.ReportOutputFormat {
		case gotable.TABLEOUTTEXT:
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Disposition", "attachment; filename="+attachmentName+".text")
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
			w.Header().Set("Content-Disposition", "attachment; filename="+attachmentName+".csv")
			err := tbl.CSVprintTable(w)
			if err != nil {
				s := fmt.Sprintf("Error in CSVprintTable: %s\n", err.Error())
				fmt.Print(s)
				fmt.Fprintf(w, "%s\n", s)
			}
			return
		case gotable.TABLEOUTPDF:
			w.Header().Set("Content-Type", "application/pdf")
			w.Header().Set("Content-Disposition", "attachment; filename="+attachmentName+".pdf")

			// pdf props title
			pdfProps := rrpt.RRpdfProps

			// get page size and orientation, set title
			pdfProps = rrpt.SetPDFOption(pdfProps, "--header-center", tbl.Title)
			pdfPageWidth := rlib.Float64ToString(ui.PDFPageWidth) + ui.PDFPageSizeUnit
			pdfProps = rrpt.SetPDFOption(pdfProps, "--page-width", pdfPageWidth)
			pdfPageHeight := rlib.Float64ToString(ui.PDFPageHeight) + ui.PDFPageSizeUnit
			pdfProps = rrpt.SetPDFOption(pdfProps, "--page-height", pdfPageHeight)

			err := tbl.PDFprintTable(w, pdfProps)
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
			if strings.ToLower(rn) == strings.ToLower(reportname) {
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

		// format downloadable report name
		attachmentName := ri.Xbiz.P.Designation + "-" + strings.Title(tmh.ReportNames[1])
		if !ri.D1.IsZero() {
			fromDate := rrpt.GetAttachmentDate(ri.D1)
			attachmentName += "-From" + fromDate
		}
		if !ri.D2.IsZero() {
			toDate := rrpt.GetAttachmentDate(ri.D2)
			attachmentName += "To" + toDate
		}

		switch ui.ReportOutputFormat {
		case gotable.TABLEOUTTEXT:
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Disposition", "attachment; filename="+attachmentName+".text")
			gotable.MultiTableTextPrint(m, w)
			return
		case gotable.TABLEOUTHTML:
			gotable.MultiTableHTMLPrint(m, w)
			return
		case gotable.TABLEOUTCSV:
			w.Header().Set("Content-Type", "text/csv")
			w.Header().Set("Content-Disposition", "attachment; filename="+attachmentName+".csv")
			gotable.MultiTableCSVPrint(m, w)
			return
		case gotable.TABLEOUTPDF:
			w.Header().Set("Content-Type", "application/pdf")
			w.Header().Set("Content-Disposition", "attachment; filename="+attachmentName+".pdf")

			var pdfTitle string
			pdfTitle += tmh.ReportTitle

			if !ri.D1.IsZero() {
				fromDate := rrpt.GetAttachmentDate(ri.D1)
				pdfTitle += " From " + fromDate
			}
			if !ri.D2.IsZero() {
				toDate := rrpt.GetAttachmentDate(ri.D2)
				pdfTitle += " To " + toDate
			}

			rrpt.MultiTablePDFPrint(m, w, pdfTitle, ui.PDFPageWidth, ui.PDFPageHeight, ui.PDFPageSizeUnit)
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
	rlib.Console("Entered %s\n", funcname)
	var ui RRuiSupport
	var d ws.ServiceData
	var xbiz rlib.XBusiness
	var reportname string
	var err error

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Only GET method is allowed")
		return
	}

	rlib.Console("r.RequestURL = %s\n", r.URL.String())
	sa := strings.Split(r.URL.Path, "/") // ["", "wsvc", "<BID>"]
	rlib.Console("sa = %#v\n", sa)
	d.BID, err = rlib.IntFromString(sa[2], "bad request integer value")
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
	rlib.Console("m = %#v\n", m)
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
	// pdf page width
	var pdfWidth float64
	x, ok = m["pw"]
	if ok && len(x[0]) > 0 {
		if pdfWidth, ok = rlib.StringToFloat64(x[0]); ok {
			ui.PDFPageWidth = pdfWidth
		}
	}
	// pdf page height
	var pdfHeight float64
	x, ok = m["ph"]
	if ok && len(x[0]) > 0 {
		if pdfHeight, ok = rlib.StringToFloat64(x[0]); ok {
			ui.PDFPageHeight = pdfHeight
		}
	}
	// pdf page size unit, take default `inch` as of now
	ui.PDFPageSizeUnit = "in"

	v1ReportHandler(reportname, &xbiz, &ui, w, &m)
	// ui.ReportContent = websvcReportHandler(reportname, &xbiz, &ui)
	// SendWebSvcPage(w, r, &ui)
}
