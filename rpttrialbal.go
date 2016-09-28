package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
	"text/template"
)

// RptTrialBalance is the http handler routine for the Trial Balance report.
func RptTrialBalance(w http.ResponseWriter, r *http.Request) {
	funcname := "trialBalanceHandler"
	// fmt.Printf("Entered %s\n", funcname)
	var ui RRuiSupport
	var err error
	var xbiz rlib.XBusiness

	getBizStartStop(w, r, &ui, &xbiz)
	if xbiz.P.BID > 0 {
		tbl := rrpt.LedgerBalanceReport(&xbiz, &ui.D2)
		if err == nil {
			ui.ReportContent = tbl.SprintTable(rlib.TABLEOUTTEXT)
		}
	}

	w.Header().Set("Content-Type", "text/html")

	t, err := template.New("rpttrialbal.html").Funcs(RRfuncMap).ParseFiles("./html/rpttrialbal.html")
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}

	err = t.Execute(w, &ui)
	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
