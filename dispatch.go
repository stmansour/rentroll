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
	s := "SELECT BID,BUD,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy from Business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() { // For every Business
		var xbiz rlib.XBusiness
		rlib.Errcheck(rows.Scan(&xbiz.P.BID, &xbiz.P.Designation, &xbiz.P.Name, &xbiz.P.DefaultRentalPeriod, &xbiz.P.ParkingPermitInUse, &xbiz.P.LastModTime, &xbiz.P.LastModBy))
		rlib.GetXBusiness(xbiz.P.BID, &xbiz) // get its info
		rlib.InitBusinessFields(xbiz.P.BID)
		rlib.GetDefaultLedgers(xbiz.P.BID) // Gather its chart of accounts
		rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts = rlib.GetGLAccountMap(xbiz.P.BID)

		// and generate the requested report...
		switch ctx.Report {
		case 1: // JOURNAL
			JournalReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 2: // LEDGER
			LedgerReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 3: // INTERNAL ACCT RULE TEST
			intTest(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 4: // ??? available ???
			fmt.Printf("biz csv = %s\n", App.bizfile)
		case 5: // ASSESSMENT CHECK REPORT
			AssessmentCheckReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 6: // LEDGER BALANCE REPORT
			var ui RRuiSupport
			ui.B = xbiz.P
			ui.DtStart = ctx.DtStart
			ui.DtStop = ctx.DtStop
			BuildXLedgerList(&ui, xbiz.P.BID, ctx.DtStart, ctx.DtStop)
			UILedgerTextReport(&ui)
		case 7: // RENTABLE COUNT BY TYPE
			UIRentableCountByRentableTypeReport(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 8: // STATEMENT
			UIStatementTextReport(&xbiz, &ctx.DtStart, &ctx.DtStop)
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
