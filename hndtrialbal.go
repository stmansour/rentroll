package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"text/template"
	"time"
)

func hndTrialBalance(w http.ResponseWriter, r *http.Request) {
	funcname := "trialBalanceHandler"
	var ui RRuiSupport
	var err error
	var biz rlib.Business

	fmt.Printf("Entered %s\n", funcname)

	D1 := r.FormValue("DtStart")
	D2 := r.FormValue("DtStop")
	des := r.FormValue("Business")

	ui.DtStart, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(D1))
	if err != nil {
		fmt.Printf("%s: Invalid start date:  %s\n", funcname, D1)
	}
	ui.DtStop, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(D2))
	if err != nil {
		fmt.Printf("%s: Invalid start date:  %s\n", funcname, D2)
	}

	if len(des) > 0 {
		biz, _ := rlib.GetBusinessByDesignation(des)
		if len(biz.Designation) == 0 {
			rlib.Ulog("%s: Business with designation %s does net exist\n", funcname, des)
			return
		}
	}

	BuildXLedgerList(&ui, biz.BID, ui.DtStart, ui.DtStop)

	w.Header().Set("Content-Type", "text/html")

	t, err := template.New("trialbal.html").Funcs(RRfuncMap).ParseFiles("./html/trialbal.html")
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}
	fmt.Printf("%s - calling t.execute\n", funcname)
	fmt.Printf("ui = %#v\n", ui)
	err = t.Execute(w, &ui)
	fmt.Printf("returned from t.Execute.\n")
	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
