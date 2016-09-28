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

func getBizStartStop(w http.ResponseWriter, r *http.Request, ui *RRuiSupport, xbiz *rlib.XBusiness) {
	now := time.Now()
	var err error

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

	UIInitBizList(ui)
	ui.D1, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(ui.DtStart))
	if err != nil {
		rlib.Ulog("getBizStartStop: could not parse %s, err = %v\n", ui.DtStart, err)
	}
	ui.D2, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(ui.DtStop))
	if err != nil {
		rlib.Ulog("getBizStartStop: could not parse %s, err = %v\n", ui.DtStop, err)
	}
	if len(BID) > 0 {
		i, err := strconv.Atoi(strings.TrimSpace(BID))
		if err != nil {
			rlib.Ulog("Error fetching business with BID = %d\n", i)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rlib.GetXBusiness(int64(i), xbiz)
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
		UIInitBizList(ui)
	}
}

// RptRentRoll is the HTTP handler for the RentRoll report request
func RptRentRoll(w http.ResponseWriter, r *http.Request) {
	funcname := "RptRentRoll"
	// fmt.Printf("Entered %s\n", funcname)
	var ui RRuiSupport
	var xbiz rlib.XBusiness

	action := r.FormValue("action")
	// fmt.Printf("RptRentRoll: action = %s\n", action)
	if action != "RentRoll" {
		dispatchHandler(w, r)
		return
	}

	getBizStartStop(w, r, &ui, &xbiz)
	if xbiz.P.BID > 0 {
		tbl, err := rrpt.RentRollReport(&xbiz, &ui.D1, &ui.D2)
		if err == nil {
			ui.ReportContent = tbl.Title + tbl.SprintRowText(len(tbl.Row)-1) + tbl.SprintLineText() + tbl.SprintTable(rlib.TABLEOUTTEXT)
		} else {
			ui.ReportContent = fmt.Sprintf("Error generating RentRoll report:  %s\n", err)
		}
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
