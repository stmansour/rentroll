package main

import (
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
)

// RptLedgerHandler is the HTTP handler for the Ledger report request
func RptLedgerHandler(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, sel int) {
	var m []rlib.Table
	if xbiz.P.BID > 0 {
		switch sel {
		case 0: // all ledgers
			m = rrpt.LedgerReport(xbiz, &ui.D1, &ui.D2)
		case 1: // ledger activity
			m = rrpt.LedgerActivityReport(xbiz, &ui.D1, &ui.D2)
		}
		ui.ReportContent = ""
		for i := 0; i < len(m); i++ {
			ui.ReportContent += m[i].Title + m[i].SprintTable(rlib.TABLEOUTTEXT) + "\n\n"
		}
	}
}

// RptLedger is the HTTP handler for the Ledger report request
func RptLedger(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, tmpl *string) {
	RptLedgerHandler(w, r, xbiz, ui, 0)
	*tmpl = "rptledger.html"
}

// RptLedgerActivity is the HTTP handler for the Ledger report request
func RptLedgerActivity(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, tmpl *string) {
	RptLedgerHandler(w, r, xbiz, ui, 1)
	*tmpl = "rptledgeract.html"
}
