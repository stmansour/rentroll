package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
	"text/template"
)

// RptDelinq is the HTTP handler for the RentRoll report request
func RptDelinq(w http.ResponseWriter, r *http.Request) {
	funcname := "RptDelinq"
	// fmt.Printf("Entered %s\n", funcname)
	var ui RRuiSupport
	var err error
	var xbiz rlib.XBusiness

	action := r.FormValue("action")
	// fmt.Printf("RptDelinq: action = %s\n", action)
	if action != "Delinquency" {
		dispatchHandler(w, r)
		return
	}

	getBizStartStop(w, r, &ui, &xbiz)
	if xbiz.P.BID > 0 {
		tbl, err := rrpt.DelinquencyReport(&xbiz, &ui.D2)
		if err == nil {
			ui.ReportContent = tbl.String()
		} else {
			ui.ReportContent = fmt.Sprintf("Error generating Delinquency report:  %s\n", err)
		}
	}

	w.Header().Set("Content-Type", "text/html")

	t, err := template.New("rptdelinq.html").Funcs(RRfuncMap).ParseFiles("./html/rptdelinq.html")
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}

	err = t.Execute(w, &ui)
	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
