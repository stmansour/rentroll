package main

import (
	"fmt"
	"gotable"
	"html/template"
	"net/http"
	"net/url"
	"rentroll/rcsv"
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

func v1HTMLPrint(w http.ResponseWriter, t *gotable.Table, e error) {
	if e != nil {
		if rlib.IsSQLNoResultsError(e) {
			fmt.Fprintf(w, "Nothing to report")
		} else {
			fmt.Fprintf(w, "Error: %s", e)
		}
		return
	}
	err := t.HTMLprintTable(w)
	if err != nil {
		s := fmt.Sprintf("Error in t.HTMLprintTable: %s\n", err.Error())
		fmt.Print(s)
		fmt.Fprintf(w, "%s\n", s)
	}
}

func v1ReportHandler(reportname string, xbiz *rlib.XBusiness, ui *RRuiSupport, w http.ResponseWriter) {
	fmt.Printf("v1ReportHandler: reportname=%s, BID=%d,  d1 = %s, d2 = %s\n", reportname, xbiz.P.BID, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))
	var ri = rrpt.ReporterInfo{OutputFormat: gotable.TABLEOUTHTML, Bid: xbiz.P.BID, D1: ui.D1, D2: ui.D2, Xbiz: xbiz, RptHeader: true, BlankLineAfterRptName: true}
	rlib.InitBizInternals(ri.Bid, xbiz)

	var t gotable.Table
	var err error

	switch strings.ToLower(reportname) {
	case "asmrpt", "assessments":
		t, err = rcsv.RRAssessmentsTable(&ri)
	case "b", "business":
		t = rcsv.RRreportBusinessTable(&ri)
	case "coa", "chart of accounts":
		rlib.InitBizInternals(ri.Bid, xbiz)
		t = rcsv.RRreportChartOfAccountsTable(&ri)
	// case "c", "custom attributes":
	// 	return rcsv.RRreportCustomAttributes(&ri)
	// case "cr", "custom attribute refs":
	// 	return rcsv.RRreportCustomAttributeRefs(&ri)
	// case "delinq":
	// 	rlib.InitBizInternals(ri.Bid, xbiz)
	// 	t, err := rrpt.DelinquencyReport(&ri)
	// 	if err != nil {
	// 		return err.Error()
	// 	}
	// 	s, err := t.SprintTable()
	// 	if err != nil {
	// 		s += err.Error()
	// 	}
	// 	return t.GetTitle() + s
	// case "dpm", "deposit methods":
	// 	return rcsv.RRreportDepositMethods(&ri)
	// case "dep", "depositories":
	// 	return rcsv.RRreportDepository(&ri)
	// case "gsr":
	// 	ri.D1 = ui.D2 // we want to look at the end of the range.  Set both D1 and D2 to the end of the range
	// 	t, err := rrpt.GSRReport(&ri)
	// 	if err != nil {
	// 		return err.Error()
	// 	}
	// 	s, err := t.SprintTable()
	// 	if err != nil {
	// 		s += err.Error()
	// 	}
	// 	return t.GetTitle() + s
	case "j":
		fmt.Printf("Handling report: j\n")
		rlib.InitBizInternals(ri.Bid, xbiz)
		ri.RptHeaderD1 = true
		ri.RptHeaderD2 = true
		t = rrpt.JournalReport(&ri)
		// err := t.HTMLprintTable(w)
		// if err != nil {
		// 	s := fmt.Sprintf("Error in t.HTMLprintTable: %s\n", err.Error())
		// 	fmt.Print(s)
		// 	fmt.Fprintf(w, "%s\n", s)
		// }
		// return rrpt.ReportToString(&t, &ri)
	// case "l", "la":
	// if xbiz.P.BID > 0 {
	// 	var m []gotable.Table
	// 	var rn string
	// 	if prefix == "l" {
	// 		rn = "Ledgers"
	// 	} else {
	// 		rn = "Ledger Activity"
	// 	}
	// 	ri.RptHeaderD1 = true
	// 	ri.RptHeaderD2 = true
	// 	s, err := rrpt.ReportHeader(rn, "websvcReportHandler", &ri)
	// 	if err != nil {
	// 		s += "\n" + err.Error()
	// 	}
	// 	switch prefix {
	// 	case "l": // all ledgers
	// 		m = rrpt.LedgerReport(&ri)
	// 	case "la": // ledger activity
	// 		m = rrpt.LedgerActivityReport(&ri)
	// 	}
	// 	for i := 0; i < len(m); i++ {
	// 		s1, err := m[i].SprintTable()
	// 		if err != nil {
	// 			s1 += err.Error()
	// 		}
	// 		s += m[i].GetTitle() + s1 + "\n\n"
	// 	}
	// 	// return s
	// }

	// case "pmt", "payment types":
	// 	return rcsv.RRreportPaymentTypes(&ri)
	// case "r", "rentables":
	// 	return rcsv.RRreportRentables(&ri)
	// case "ra", "rental agreements":
	// 	return rcsv.RRreportRentalAgreements(&ri)
	// case "rat", "rental agreement templates":
	// 	return rcsv.RRreportRentalAgreementTemplates(&ri)
	case "rcpt", "receipts":
		t = rcsv.RRReceiptsTable(&ri)
	// case "rr":
	// 	rlib.InitBizInternals(ri.Bid, xbiz)
	// 	return rrpt.RentRollReportString(&ri)
	// case "rt", "rentable types":
	// 	return rcsv.RRreportRentableTypes(&ri)
	// case "rcbt", "rentable Count By Type":
	// 	return rrpt.RentableCountByRentableTypeReport(&ri)
	// case "sl", "string lists":
	// 	return rcsv.RRreportStringLists(&ri)
	// case "statements":
	// 	return rrpt.RptStatementTextReport(&ri)
	// case "t", "people": // t = transactant
	// 	return rcsv.RRreportPeople(&ri)
	// case "tb":
	// 	return rrpt.PrintLedgerBalanceReportString(&ri)
	default:
		fmt.Fprintf(w, "Unknown report type: %s", reportname)
		return
	}
	// folderPath, _ := osext.ExecutableFolder()
	// t.SetHTMLTemplate(folderPath)
	// t.SetHTMLTemplateCSS(folderPath)
	v1HTMLPrint(w, &t, err)
}

func websvcReportHandler(prefix string, xbiz *rlib.XBusiness, ui *RRuiSupport) string {
	fmt.Printf("websvcReportHandler: prefix=%s, BID=%d,  d1 = %s, d2 = %s\n", prefix, xbiz.P.BID, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))
	var ri = rrpt.ReporterInfo{OutputFormat: gotable.TABLEOUTTEXT, Bid: xbiz.P.BID, D1: ui.D1, D2: ui.D2, Xbiz: xbiz, RptHeader: true, BlankLineAfterRptName: true}
	rlib.InitBizInternals(ri.Bid, xbiz)

	switch strings.ToLower(prefix) {
	case "asmrpt", "assessments":
		return rcsv.RRreportAssessments(&ri)
	case "b", "business":
		t := rcsv.RRreportBusinessTable(&ri)
		s, err := t.SprintTable()
		if err != nil {
			s += err.Error()
		}
		return t.GetTitle() + s
	case "coa", "chart of accounts":
		rlib.InitBizInternals(ri.Bid, xbiz)
		return rcsv.RRreportChartOfAccounts(&ri)
	case "c", "custom attributes":
		return rcsv.RRreportCustomAttributes(&ri)
	case "cr", "custom attribute refs":
		return rcsv.RRreportCustomAttributeRefs(&ri)
	case "delinq":
		rlib.InitBizInternals(ri.Bid, xbiz)
		t, err := rrpt.DelinquencyReport(&ri)
		if err != nil {
			return err.Error()
		}
		s, err := t.SprintTable()
		if err != nil {
			s += err.Error()
		}
		return t.GetTitle() + s
	case "dpm", "deposit methods":
		return rcsv.RRreportDepositMethods(&ri)
	case "dep", "depositories":
		return rcsv.RRreportDepository(&ri)
	case "gsr":
		ri.D1 = ui.D2 // we want to look at the end of the range.  Set both D1 and D2 to the end of the range
		t, err := rrpt.GSRReport(&ri)
		if err != nil {
			return err.Error()
		}
		s, err := t.SprintTable()
		if err != nil {
			s += err.Error()
		}
		return t.GetTitle() + s
	case "j":
		rlib.InitBizInternals(ri.Bid, xbiz)
		ri.RptHeaderD1 = true
		ri.RptHeaderD2 = true
		t := rrpt.JournalReport(&ri)
		return rrpt.ReportToString(&t, &ri)
	case "l", "la":
		if xbiz.P.BID > 0 {
			var m []gotable.Table
			var rn string
			if prefix == "l" {
				rn = "Ledgers"
			} else {
				rn = "Ledger Activity"
			}
			ri.RptHeaderD1 = true
			ri.RptHeaderD2 = true
			s, err := rrpt.ReportHeader(rn, "websvcReportHandler", &ri)
			if err != nil {
				s += "\n" + err.Error()
			}
			switch prefix {
			case "l": // all ledgers
				m = rrpt.LedgerReport(&ri)
			case "la": // ledger activity
				m = rrpt.LedgerActivityReport(&ri)
			}
			for i := 0; i < len(m); i++ {
				s1, err := m[i].SprintTable()
				if err != nil {
					s1 += err.Error()
				}
				s += m[i].GetTitle() + s1 + "\n\n"
			}
			return s
		}

	case "pmt", "payment types":
		return rcsv.RRreportPaymentTypes(&ri)
	case "r", "rentables":
		return rcsv.RRreportRentables(&ri)
	case "ra", "rental agreements":
		return rcsv.RRreportRentalAgreements(&ri)
	case "rat", "rental agreement templates":
		return rcsv.RRreportRentalAgreementTemplates(&ri)
	case "rcpt", "receipts":
		return rcsv.RRreportReceipts(&ri)
	case "rr":
		rlib.InitBizInternals(ri.Bid, xbiz)
		return rrpt.RentRollReportString(&ri)
	case "rt", "rentable types":
		return rcsv.RRreportRentableTypes(&ri)
	case "rcbt", "rentable Count By Type":
		return rrpt.RentableCountByRentableTypeReport(&ri)
	case "sl", "string lists":
		return rcsv.RRreportStringLists(&ri)
	case "statements":
		return rrpt.RptStatementTextReport(&ri)
	case "t", "people": // t = transactant
		return rcsv.RRreportPeople(&ri)
	case "tb":
		return rrpt.PrintLedgerBalanceReportString(&ri)
	}
	return "unhandled loader type: " + prefix
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
	}
	switch strings.ToLower(reportname) {
	case "asmrpt", "assessments", "b", "business", "coa", "chart of accounts", "j", "rcpt", "receipts":
		v1ReportHandler(reportname, &xbiz, &ui, w)
	default:
		ui.ReportContent = websvcReportHandler(reportname, &xbiz, &ui)
		SendWebSvcPage(w, r, &ui)
	}
	// ui.ReportContent = websvcReportHandler(reportname, &xbiz, &ui)
	// SendWebSvcPage(w, r, &ui)
}
