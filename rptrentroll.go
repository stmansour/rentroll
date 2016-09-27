package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// RptRentRoll is the HTTP handler for the RentRoll report request
func RptRentRoll(w http.ResponseWriter, r *http.Request) {
	funcname := "RptRentRoll"
	// fmt.Printf("Entered %s\n", funcname)
	var ui RRuiSupport
	var err error
	var xbiz rlib.XBusiness

	now := time.Now()

	action := r.FormValue("action")
	// fmt.Printf("RptRentRoll: action = %s\n", action)
	if action != "RentRoll" {
		dispatchHandler(w, r)
		return
	}

	// Init to reasonable starting values
	ui.DtStart = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).Format(rlib.RRDATEINPFMT)
	ui.DtStop = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0).Format(rlib.RRDATEINPFMT)

	// override with input if available
	s := r.FormValue("DtStart")
	if s != "" {
		ui.DtStart = s
	}
	s = r.FormValue("DtStop")
	if s != "" {
		ui.DtStop = s
	}
	BID := r.FormValue("BID")

	UIInitBizList(&ui)
	dtStart, err := time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(ui.DtStart))
	dtStop, err := time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(ui.DtStop))
	if len(BID) > 0 {
		i, err := strconv.Atoi(strings.TrimSpace(BID))
		if err != nil {
			rlib.Ulog("Error fetching business with BID = %d\n", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rlib.GetXBusiness(int64(i), &xbiz)
		if xbiz.P.BID == 0 {
			rlib.Ulog("Error fetching business with BID = %d\n", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Build up the data we need in the UI Context...
		ui.B = xbiz.P
		rlib.InitBusinessFields(ui.B.BID)
		rlib.GetDefaultLedgers(xbiz.P.BID) // Gather its chart of accounts
		rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts = rlib.GetGLAccountMap(xbiz.P.BID)
		UIInitBizList(&ui)

		tbl, err := rrpt.RentRollReport(&xbiz, &dtStart, &dtStop)
		ui.ReportContent = tbl.Title + tbl.SprintRowText(len(tbl.Row)-1) + tbl.SprintLineText() + tbl.SprintTable(rlib.TABLEOUTTEXT)
	}

	w.Header().Set("Content-Type", "text/html")

	t, err := template.New("rptrentroll.html").Funcs(RRfuncMap).ParseFiles("./html/rptrentroll.html")
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}

	err = t.Execute(w, &ui)
	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
