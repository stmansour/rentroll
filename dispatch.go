package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"text/template"
)

const ()

// RunBooks runs a series of commands to handle command line run requests
func RunBooks(ctx *DispatchCtx) {
	s := "SELECT BID,DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy from business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() { // For every business
		var xbiz rlib.XBusiness
		rlib.Errcheck(rows.Scan(&xbiz.P.BID, &xbiz.P.Designation, &xbiz.P.Name, &xbiz.P.DefaultRentalPeriod, &xbiz.P.ParkingPermitInUse, &xbiz.P.LastModTime, &xbiz.P.LastModBy))
		rlib.GetXBusiness(xbiz.P.BID, &xbiz) // get its info
		rlib.InitBusinessFields(xbiz.P.BID)
		rlib.GetDefaultLedgers(xbiz.P.BID) // Gather its chart of accounts

		// and generate the requested report...
		switch ctx.Report {
		case 1:
			JournalReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 2:
			LedgerReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 3:
			intTest(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 4:
			fmt.Printf("biz csv = %s\n", App.bizfile)
		default:
			GenerateJournalRecords(&xbiz, &ctx.DtStart, &ctx.DtStop)
			GenerateLedgerRecords(&xbiz, &ctx.DtStart, &ctx.DtStop)
		}
	}
}

// Dispatch generates the supplied report for all properties
func Dispatch(ctx *DispatchCtx) {
	switch ctx.Cmd {
	case CmdRUNBOOKS:
		RunBooks(ctx)
	default:
		fmt.Printf("Unrecognized command: %d\n", ctx.Cmd)
	}
}

func dispatchHandler(w http.ResponseWriter, r *http.Request) {
	var ui RRuiSupport
	funcname := "dispatchHandler"
	w.Header().Set("Content-Type", "text/html")

	action := r.FormValue("action")
	url := r.FormValue("url")
	if len(action) > 0 {
		switch action {
		case "Ledger Balance":
			http.Redirect(w, r, url, http.StatusFound)
			return
		default:
			fmt.Printf("%s: Unrecognized action: %s\n", funcname, action)
		}
	}

	t, err := template.New("dispatch.html").Funcs(RRfuncMap).ParseFiles("./html/dispatch.html")
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}
	err = t.Execute(w, &ui)
	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
