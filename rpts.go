package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
)

// RptDelinq is the HTTP handler for the RentRoll report request
func RptDelinq(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, tmpl *string) {
	*tmpl = "rptdelinq.html"
	if xbiz.P.BID > 0 {
		tbl, err := rrpt.DelinquencyReport(xbiz, &ui.D2)
		if err == nil {
			ui.ReportContent = tbl.String()
		} else {
			ui.ReportContent = fmt.Sprintf("Error generating Delinquency report:  %s\n", err)
		}
	}
}

// RptJournal is the HTTP handler for the Journal report request
func RptJournal(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, tmpl *string) {
	*tmpl = "rptjournal.html"
	if xbiz.P.BID > 0 {
		tbl := rrpt.JournalReport(xbiz, &ui.D1, &ui.D2)
		ui.ReportContent = tbl.Title + tbl.SprintTable(rlib.TABLEOUTTEXT)
	}
}
