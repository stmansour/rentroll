package main

import (
	"fmt"
	"gotable"
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
)

// RptDelinq is the HTTP handler for the RentRoll report request
func RptDelinq(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	if xbiz.P.BID > 0 {
		var ri rrpt.ReporterInfo
		ri.Xbiz = xbiz
		ri.D2 = ui.D2
		tbl, err := rrpt.DelinquencyReport(&ri)
		if err == nil {
			ui.ReportContent = tbl.String()
		} else {
			ui.ReportContent = fmt.Sprintf("Error generating Delinquency report:  %s\n", err)
		}
	}
}

// RptGSR is the http handler routine for the Trial Balance report.
func RptGSR(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	if xbiz.P.BID > 0 {
		var ri rrpt.ReporterInfo
		ri.Xbiz = xbiz
		ri.D1 = ui.D2 // set both dates to the range end
		ri.D2 = ui.D2
		tbl, err := rrpt.GSRReport(&ri)
		if err == nil {
			s, err := tbl.SprintTable()
			if nil != err {
				s += err.Error()
			}
			ui.ReportContent = s
		}
	}
}

// RptJournal is the HTTP handler for the Journal report request
func RptJournal(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	if xbiz.P.BID > 0 {
		var ri = rrpt.ReporterInfo{Xbiz: xbiz, D1: ui.D1, D2: ui.D2, OutputFormat: gotable.TABLEOUTTEXT}
		ri.OutputFormat = gotable.TABLEOUTTEXT
		tbl := rrpt.JournalReport(&ri)
		ri.RptHeader = true
		ri.RptHeaderD1 = true
		ri.RptHeaderD2 = true
		// ui.ReportContent = tbl.Title + tbl.SprintTable()
		ui.ReportContent = rrpt.ReportToString(&tbl, &ri)
	}
}

// RptLedgerHandler is the HTTP handler for the Ledger report request
func RptLedgerHandler(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, sel int) {
	var ri = rrpt.ReporterInfo{Xbiz: xbiz, D1: ui.D1, D2: ui.D2}
	var m []gotable.Table
	var rn string
	if sel == 0 {
		rn = "Ledgers"
	} else {
		rn = "Ledger Activity"
	}
	s, err := rrpt.ReportHeader(rn, "RptLedgerHandler", &ri)
	if err != nil {
		s += "\n" + err.Error()
	}
	ui.ReportContent += s

	if xbiz.P.BID > 0 {
		switch sel {
		case 0: // all ledgers
			m = rrpt.LedgerReport(&ri)
		case 1: // ledger activity
			m = rrpt.LedgerActivityReport(&ri)
		}
		ui.ReportContent = ""
		for i := 0; i < len(m); i++ {
			s, err := m[i].SprintTable()
			if err != nil {
				s += err.Error()
			}
			ui.ReportContent += m[i].GetTitle() + s + "\n\n"
		}
	}
}

// RptLedger is the HTTP handler for the Ledger report request
func RptLedger(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	RptLedgerHandler(w, r, xbiz, ui, 0)
}

// RptLedgerActivity is the HTTP handler for the Ledger report request
func RptLedgerActivity(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	RptLedgerHandler(w, r, xbiz, ui, 1)
}

// RptRentRoll is the HTTP handler for the RentRoll report request
func RptRentRoll(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	var ri = rrpt.ReporterInfo{Xbiz: xbiz, D1: ui.D1, D2: ui.D2}
	if xbiz.P.BID > 0 {
		tbl, err := rrpt.RentRollReport(&ri)
		if err == nil {
			tout, err := tbl.SprintTable()
			if err != nil {
				rlib.Ulog("RptRentRoll:  error = %s", err)
				ui.ReportContent = err.Error()
			} else {
				ui.ReportContent = tout
			}
		} else {
			ui.ReportContent = fmt.Sprintf("Error generating RentRoll report:  %s\n", err)
		}
	}
}

// RptTrialBalance is the http handler routine for the Trial Balance report.
func RptTrialBalance(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	var err error
	var ri = rrpt.ReporterInfo{Xbiz: xbiz, D1: ui.D1, D2: ui.D2}
	if xbiz.P.BID > 0 {
		tbl := rrpt.LedgerBalanceReport(&ri)
		if err == nil {
			s, err := tbl.SprintTable()
			if err != nil {
				s += err.Error()
			}
			ui.ReportContent = s
		}
	}
}
