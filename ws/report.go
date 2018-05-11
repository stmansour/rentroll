package ws

import (
	"context"
	"fmt"
	"gotable"
	"html/template"
	"net/http"
	"net/url"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strings"
	"time"
)

// ReportServiceHandler is the choke-point for all report requests.  It was
// changed from being the original /wsvc entry point in Dec 2017.  Now, all
// report requests are part of the /v1 entry point.
//
// This service handles requests of the form:
//    http://x.y.z/v1/report/BUD?[params]
// where params can be:
//	  r=<reportname>
//    dtstart=<date>
//    dtstop=<date>
//    edi={0|1}         {0 = default = end date is non-inclusive,
//                       1 = end date is inclusive}
//-----------------------------------------------------------------------------
func ReportServiceHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "ReportServiceHandler"
	var (
		ui         rrpt.ReportContext
		xbiz       rlib.XBusiness
		reportname string
		err        error
		edi        = int(0) // 0 = end date is not inclusive, 1 = end date is inclusive
	)
	rlib.Console("Entered %s\n", funcname)

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := fmt.Errorf("Only GET method is allowed")
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("r.RequestURL = %s\n", r.URL.String())
	// sa := strings.Split(r.URL.Path, "/") // ["", "wsvc", "<BID>"]
	// rlib.Console("sa = %#v\n", sa)
	// d.BID, err = rlib.IntFromString(sa[2], "bad request integer value")
	// if err != nil {
	// 	ui.ReportContent = fmt.Sprintf("Error parsing request URI: %s", err.Error())
	// 	SendWebSvcPage(w, r, &ui)
	// 	return
	// }

	err = rlib.GetXBusiness(r.Context(), d.BID, &xbiz)
	if err != nil {
		ui.ReportContent = fmt.Sprintf("Error while fetching business: %s", err.Error())
		SendWebSvcPage(w, r, &ui)
		return
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

	ui.ID = d.ID // pass this along, regardless of its value. Reports that detail a specific entity will use it

	tnow := time.Now()
	ui.D1 = time.Date(tnow.Year(), tnow.Month(), 1, 0, 0, 0, 0, time.UTC)
	nm := int(ui.D1.Month()) + 1
	yr := ui.D1.Year()
	if nm > int(time.December) {
		nm = int(time.January)
		yr++
	}
	ui.D2 = time.Date(yr, time.Month(nm), 1, 0, 0, 0, 0, time.UTC)

	//---------------------------------
	// end date inclusive flag
	//---------------------------------
	edifl, ok := m["edi"]
	if ok && edifl[0] == "1" {
		edi = 1
	}
	ui.EDI = edi // set the context flag

	// rlib.Console("edi flag = %d\n", edi)

	//---------------------------------
	// dtstart
	//---------------------------------
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
	//---------------------------------
	// dtstop - must always be the up-to-but-not-including date.
	//
	//---------------------------------
	x, ok = m["dtstop"]
	var d2 time.Time
	if ok && len(x[0]) > 0 {
		d2, err = rlib.StringToDate(x[0])

		// if dateMode is on then change the stopDate value for search op
		rlib.HandleFrontEndDates(d.BID, &d1, &d2)

		if err == nil {
			ui.D2 = d2
			// ui.ReportContent = fmt.Sprintf("Error with dtstop value:  %s", err.Error())
			// SendWebSvcPage(w, r, &ui)
			// return
		}
	}

	//---------------------------------
	// report output format
	//---------------------------------
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

	v1ReportHandler(r.Context(), reportname, &xbiz, &ui, w, &m)
	// ui.ReportContent = websvcReportHandler(reportname, &xbiz, &ui)
	// SendWebSvcPage(w, r, &ui)
}

// SendWebSvcPage creates an HTML page where the content is simply
// the contents of ui.ReportContent formatted to be shown as-is in
// a mono-spaced font.
func SendWebSvcPage(w http.ResponseWriter, r *http.Request, ui *rrpt.ReportContext) {
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

func v1ReportHandler(ctx context.Context, reportname string, xbiz *rlib.XBusiness, ui *rrpt.ReportContext, w http.ResponseWriter, qp *url.Values) {
	const funcname = "v1ReportHandler"
	rlib.Console("%s: reportname=%s, BID=%d,  d1 = %s, d2 = %s\n", funcname, reportname, xbiz.P.BID, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))

	var (
		err error
		ri  = rrpt.ReporterInfo{
			BlankLineAfterRptName: true,
			OutputFormat:          gotable.TABLEOUTHTML,
			Bid:                   xbiz.P.BID,
			D1:                    ui.D1,
			D2:                    ui.D2,
			Xbiz:                  xbiz,
			EDI:                   ui.EDI, // End-Date-includes
			QueryParams:           qp,
			ID:                    ui.ID,
		}
	)

	// init business internals first
	err = rlib.InitBizInternals(ri.Bid, xbiz)
	if err != nil {
		s := fmt.Sprintf("Error in InitBizInternals: %s\n", err.Error())
		fmt.Print(s)
		fmt.Fprintf(w, "%s\n", s)
		return
	}

	// handler for reports which has single table
	var wsr = []rrpt.SingleTableReportHandler{
		{ReportNames: []string{"RPTar", "account rules"}, TableHandler: rrpt.RRARTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTasmrpt", "assessments"}, TableHandler: rrpt.RRAssessmentsTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTb", "business"}, TableHandler: rrpt.RRreportBusinessTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTc", "custom attributes"}, TableHandler: rrpt.RRreportCustomAttributesTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTcoa", "chart of accounts"}, TableHandler: rrpt.RRreportChartOfAccountsTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTcr", "custom attribute refs"}, TableHandler: rrpt.RRreportCustomAttributeRefsTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTdelinq", "delinquency"}, TableHandler: rrpt.DelinquencyReportTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTdep", "depositories"}, TableHandler: rrpt.RRreportDepositoryTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTdpm", "deposit methods"}, TableHandler: rrpt.RRreportDepositMethodsTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTgsr", "gsr"}, TableHandler: rrpt.GSRReportTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTj", "journals"}, TableHandler: rrpt.JournalReportTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTpayorstmt", "payor statements"}, TableHandler: rrpt.RRPayorStatement, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTpeople", "people"}, TableHandler: rrpt.RRreportPeopleTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTpmt", "payment types"}, TableHandler: rrpt.RRreportPaymentTypesTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTr", "rentables"}, TableHandler: rrpt.RRreportRentablesTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTra", "rental agreements"}, TableHandler: rrpt.RRreportRentalAgreementsTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTrastmt", "rental agreement statements"}, TableHandler: rrpt.RRRentalAgreementStatements, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTrat", "rental agreement templates"}, TableHandler: rrpt.RRreportRentalAgreementTemplatesTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTrcbt", "rentable type counts"}, TableHandler: rrpt.RentableCountByRentableTypeReportTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTrcpt", "receipt"}, TableHandler: rrpt.RRRcptOnlyReceiptTable, PDFprops: rrpt.ReceiptPDFProps, HTMLTemplate: "receipt.html", NeedsCustomPDFDimension: false, NeedsPDFTitle: false},
		{ReportNames: []string{"RPTrcpthotel", ""}, TableHandler: rrpt.RRRcptHotelReceiptTable, PDFprops: rrpt.ReceiptPDFProps, HTMLTemplate: "rcpthotel.html", NeedsCustomPDFDimension: false, NeedsPDFTitle: false},
		{ReportNames: []string{"RPTrcptlist", "receipts"}, TableHandler: rrpt.RRReceiptsTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTrr", "rentroll"}, TableHandler: rrpt.RRReportTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTrt", "rentable types"}, TableHandler: rrpt.RRreportRentableTypesTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTsl", "string lists"}, TableHandler: rrpt.RRreportStringListsTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTt", "people"}, TableHandler: rrpt.RRreportPeopleTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTtb", "trial balance"}, TableHandler: rrpt.LedgerBalanceReportTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
		{ReportNames: []string{"RPTtl", "task list"}, TableHandler: rrpt.TaskListReportTable, PDFprops: nil, HTMLTemplate: "", NeedsCustomPDFDimension: true, NeedsPDFTitle: true},
	}

	// handler for reports which has more than one table
	var wmr = []rrpt.MultiTableReportHandler{
		{ReportTitle: "Ledger", ReportNames: []string{"RPTl", "ledger"}, TableHandler: rrpt.LedgerReportTable, PDFprops: nil, NeedsCustomPDFDimension: true},
		{ReportTitle: "Ledger Activity", ReportNames: []string{"RPTla", "ledger activity"}, TableHandler: rrpt.LedgerActivityReportTable, PDFprops: nil, NeedsCustomPDFDimension: true},
		{ReportTitle: "Report Statements", ReportNames: []string{"RPTstatements", "report statements"}, TableHandler: rrpt.RptStatementReportTable, PDFprops: nil, NeedsCustomPDFDimension: true},
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
		tbl := tsh.TableHandler(ctx, &ri)

		// format downloadable report name
		attachmentName := ri.Xbiz.P.Designation + "-" + strings.Title(tsh.ReportNames[1])
		if !ri.D1.IsZero() {
			fromDate := rrpt.GetAttachmentDate(ri.D1)
			attachmentName += "-From:" + fromDate
		}
		if !ri.D2.IsZero() {
			d2 := ri.D2

			// if EDI mode enabled then we should subtract one day from the report name
			rlib.HandleStopDateEDI(ri.Bid, &d2)

			toDate := rrpt.GetAttachmentDate(d2)
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

			// custom template if available
			tfname := tsh.HTMLTemplate
			if len(tfname) > 0 {
				bud := rlib.GetBUDFromBIDList(ri.Bid)
				tfname = "webclient/html/rpt-templates/" + strings.ToUpper(string(bud)) + "/" + tfname
				err := tbl.SetHTMLTemplate(tfname)
				if err != nil {
					s := fmt.Sprintf("Error in CSVprintTable: %s\n", err.Error())
					fmt.Print(s)
					fmt.Fprintf(w, "%s\n", s)
				}
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

			// custom template if available
			tfname := tsh.HTMLTemplate
			// rlib.Console("report.go:  tfname = %s\n", tfname)
			if len(tfname) > 0 {
				var err error
				bud := rlib.GetBUDFromBIDList(ri.Bid)
				tfname = "webclient/html/rpt-templates/" + strings.ToUpper(string(bud)) + "/" + tfname

				// cwd, err := os.Getwd()
				// rlib.Console("report.go:  cwd = %s\n", cwd)
				// rlib.Console("report.go:  setHTMLTemplate to %s\n", tfname)
				err = tbl.SetHTMLTemplate(tfname)
				if err != nil {
					s := fmt.Sprintf("Error in CSVprintTable: %s\n", err.Error())
					fmt.Print(s)
					fmt.Fprintf(w, "%s\n", s)
				}
			}

			// pdf props title
			var pdfProps []*gotable.PDFProperty
			if tsh.PDFprops != nil {
				pdfProps = tsh.PDFprops
			} else {
				pdfProps = rrpt.GetReportPDFProps()
			}

			// get page size and orientation, set title
			if tsh.NeedsPDFTitle {
				// NOTE: There is no support of custom title right now
				pdfProps = rrpt.SetPDFOption(pdfProps, "--header-center", tbl.Title)
			}

			// if custom dimension needy, then get those from client side
			if tsh.NeedsCustomPDFDimension {
				// pdf page width from the UI
				pdfPageWidth := rlib.Float64ToString(ui.PDFPageWidth) + ui.PDFPageSizeUnit
				pdfProps = rrpt.SetPDFOption(pdfProps, "--page-width", pdfPageWidth)

				// pdf page height from the UI
				pdfPageHeight := rlib.Float64ToString(ui.PDFPageHeight) + ui.PDFPageSizeUnit
				pdfProps = rrpt.SetPDFOption(pdfProps, "--page-height", pdfPageHeight)
			}

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
		m, err := tmh.TableHandler(ctx, &ri)
		if err != nil {
			s := fmt.Sprintf("Error in Multi Table Handler for \"%s\" : %s\n", tmh.ReportTitle, err.Error())
			fmt.Print(s)
			fmt.Fprintf(w, "%s\n", s)
			return
		}

		// format downloadable report name
		attachmentName := ri.Xbiz.P.Designation + "-" + strings.Title(tmh.ReportNames[1])
		if !ri.D1.IsZero() {
			fromDate := rrpt.GetAttachmentDate(ri.D1)
			attachmentName += "-From" + fromDate
		}
		if !ri.D2.IsZero() {
			d2 := ri.D2

			// if EDI mode enabled then we should subtract one day from the report name
			rlib.HandleStopDateEDI(ri.Bid, &d2)

			toDate := rrpt.GetAttachmentDate(d2)
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

			// now, prepare pdf title for this multi table report
			var pdfTitle string
			pdfTitle += tmh.ReportTitle

			if !ri.D1.IsZero() {
				fromDate := rrpt.GetAttachmentDate(ri.D1)
				pdfTitle += " From " + fromDate
			}
			if !ri.D2.IsZero() {
				d2 := ri.D2

				// if EDI mode enabled then we should subtract one day from the report name
				rlib.HandleStopDateEDI(ri.Bid, &d2)

				toDate := rrpt.GetAttachmentDate(d2)
				pdfTitle += " To " + toDate
			}

			// get pdf props slice
			var pdfProps []*gotable.PDFProperty
			if tsh.PDFprops != nil {
				pdfProps = tsh.PDFprops
			} else {
				pdfProps = rrpt.GetReportPDFProps()
			}

			// report title
			pdfProps = rrpt.SetPDFOption(pdfProps, "--header-center", pdfTitle)

			// if custom dimension is set then get values from UI
			if tmh.NeedsCustomPDFDimension {
				pw := rlib.Float64ToString(ui.PDFPageWidth) + ui.PDFPageSizeUnit
				pdfProps = rrpt.SetPDFOption(pdfProps, "--page-width", pw)
				ph := rlib.Float64ToString(ui.PDFPageHeight) + ui.PDFPageSizeUnit
				pdfProps = rrpt.SetPDFOption(pdfProps, "--page-height", ph)
			}

			// now finally hit to print the report in io.Writer of w
			err := gotable.MultiTablePDFPrint(m, w, pdfProps)
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

	// unknown report handler
	fmt.Fprintf(w, "Unknown report type: %s", reportname)
	return
}
