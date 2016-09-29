package main

import (
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
)

// RptGSR is the http handler routine for the Trial Balance report.
func RptGSR(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, tmpl *string) {
	*tmpl = "rptgsr.html"
	if xbiz.P.BID > 0 {
		tbl, err := rrpt.GSRReport(xbiz, &ui.D2)
		if err == nil {
			ui.ReportContent = tbl.SprintTable(rlib.TABLEOUTTEXT)
		}
	}
}
