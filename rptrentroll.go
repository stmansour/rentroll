package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strconv"
	"strings"
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
		InitBizInternals(int64(i), xbiz)
		UIInitBizList(ui)
	}
}

// RptRentRoll is the HTTP handler for the RentRoll report request
func RptRentRoll(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, tmpl *string) {
	// funcname := "RptRentRoll"
	*tmpl = "rptrentroll.html"
	if xbiz.P.BID > 0 {
		tbl, err := rrpt.RentRollReport(xbiz, &ui.D1, &ui.D2)
		if err == nil {
			ui.ReportContent = tbl.Title + tbl.SprintRowText(len(tbl.Row)-1) + tbl.SprintLineText() + tbl.SprintTable(rlib.TABLEOUTTEXT)
		} else {
			ui.ReportContent = fmt.Sprintf("Error generating RentRoll report:  %s\n", err)
		}
	}
}
