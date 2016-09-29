package main

import (
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
)

// RptTrialBalance is the http handler routine for the Trial Balance report.
func RptTrialBalance(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, tmpl *string) {
	var err error
	*tmpl = "rpttrialbal.html"
	if xbiz.P.BID > 0 {
		tbl := rrpt.LedgerBalanceReport(xbiz, &ui.D2)
		if err == nil {
			ui.ReportContent = tbl.SprintTable(rlib.TABLEOUTTEXT)
		}
	}
}
