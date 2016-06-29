package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func hndTrialBalance(w http.ResponseWriter, r *http.Request) {
	funcname := "trialBalanceHandler"
	var ui RRuiSupport
	var err error
	var biz rlib.Business

	D1 := r.FormValue("DtStart")
	D2 := r.FormValue("DtStop")
	BID := r.FormValue("BID")

	ui.DtStart, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(D1))
	if err != nil {
		fmt.Printf("%s: Invalid start date:  %s\n", funcname, D1)
	}
	ui.DtStop, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(D2))
	if err != nil {
		fmt.Printf("%s: Invalid start date:  %s\n", funcname, D2)
	}

	if len(BID) > 0 {
		i, err := strconv.Atoi(strings.TrimSpace(BID))
		if err != nil {
			rlib.Ulog("Error fetching business with BID = %d\n", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rlib.GetBusiness(int64(i), &biz)
		if biz.BID == 0 {
			rlib.Ulog("Error fetching business with BID = %d\n", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	BuildXLedgerList(&ui, biz.BID, ui.DtStart, ui.DtStop)
	ui.B = biz

	w.Header().Set("Content-Type", "text/html")

	t, err := template.New("trialbal.html").Funcs(RRfuncMap).ParseFiles("./html/trialbal.html")
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}

	err = t.Execute(w, &ui)
	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
