package main

import (
	"fmt"
	"net/http"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strings"
	"text/template"
)

const ()

// RunBooks runs a series of commands to handle command line run requests
func RunBooks(ctx *DispatchCtx) {
	s := "SELECT BID,BUD,Name,DefaultRentCycle,LastModTime,LastModBy from Business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() { // For every Business
		var xbiz rlib.XBusiness
		rlib.Errcheck(rows.Scan(&xbiz.P.BID, &xbiz.P.Designation, &xbiz.P.Name, &xbiz.P.DefaultRentCycle, &xbiz.P.LastModTime, &xbiz.P.LastModBy))
		rlib.GetXBusiness(xbiz.P.BID, &xbiz) // get its info
		rlib.InitBusinessFields(xbiz.P.BID)
		rlib.GetDefaultLedgers(xbiz.P.BID) // Gather its chart of accounts
		rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts = rlib.GetGLAccountMap(xbiz.P.BID)
		rlib.GetAllNoteTypes(xbiz.P.BID)
		rlib.LoadRentableTypeCustomaAttributes(&xbiz)

		// and generate the requested report...
		switch ctx.Report {
		case 1: // JOURNAL
			JournalReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 2: // LEDGER
			LedgerReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 3: // INTERNAL ACCT RULE TEST
			intTest(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 4: // RENTROLL REPORT
			err = rrpt.RentRollTextReport(&xbiz, &ctx.DtStart, &ctx.DtStop)
			if err != nil {
				fmt.Printf("RentRoll text report error: %s\n", err.Error())
			}
		case 5: // ASSESSMENT CHECK REPORT
			AssessmentCheckReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 6: // LEDGER BALANCE REPORT
			var ui RRuiSupport
			ui.B = xbiz.P
			ui.DtStart = ctx.DtStart
			ui.DtStop = ctx.DtStop
			BuildXLedgerList(&ui, xbiz.P.BID, ctx.DtStop)
			UILedgerTextReport(&ui)
		case 7: // RENTABLE COUNT BY TYPE
			UIRentableCountByRentableTypeReport(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 8: // STATEMENT
			UIStatementTextReport(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 9: // Invoice
			// ctx.Report format:  9,IN0001  or  9,1   -- both say that we want Invoice 1 to be printed
			sa := strings.Split(ctx.Args, ",")
			if len(sa) < 2 {
				fmt.Printf("Missing InvoiceNo on report option.  Example:  -r 9,IN000001\n")
				os.Exit(1)
			}
			invoiceno := rcsv.CSVLoaderGetInvoiceNo(sa[1])
			rrpt.InvoiceTextReport(invoiceno)
		case 10: // LEDGER ACTIVITY
			LedgerActivityReport(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 11: // RENTABLE GSR
			rrpt.GSRTextReport(&xbiz, &ctx.DtStart)
		case 12: // LEDGERBALANCE ON DATE
			// ctx.Report format:  12,LID,RAID,date
			sa := strings.Split(ctx.Args, ",")
			if len(sa) < 4 {
				fmt.Printf("Missing one or more parameters.  Example:  -r 12,L004,RA003,2016-07-04\n")
				os.Exit(1)
			}
			lid := rcsv.CSVLoaderGetLedgerNo(sa[1])
			raid := rcsv.CSVLoaderGetRAID(sa[2])
			dt, err := rlib.StringToDate(sa[3])
			if err != nil {
				fmt.Printf("Bad date string: %s\n", sa[3])
				os.Exit(1)
			}
			rrpt.LdgAcctBalOnDateTextReport(&xbiz, lid, raid, &dt)
		case 13: // RA LEDGER DETAILS OVER RANGE
			// ctx.Report format: 13,LID,RAID
			// date range is from -j , -k
			sa := strings.Split(ctx.Args, ",")
			if len(sa) < 3 {
				fmt.Printf("Missing one or more parameters.  Example:  -r 13,L004,RA003\n")
				os.Exit(1)
			}
			lid := rcsv.CSVLoaderGetLedgerNo(sa[1])
			raid := rcsv.CSVLoaderGetRAID(sa[2])
			rrpt.RAAccountActivityRangeDetail(&xbiz, lid, raid, &ctx.DtStart, &ctx.DtStop)
		case 14: // DELINQUENCY REPORT
			// ctx.Report format:  14,date
			sa := strings.Split(ctx.Args, ",")
			if len(sa) < 2 {
				fmt.Printf("Missing one or more parameters.  Example:  -r 14,2016-05-25\n")
				os.Exit(1)
			}
			dt, err := rlib.StringToDate(sa[1])
			if err != nil {
				fmt.Printf("Bad date string: %s\n", sa[1])
				os.Exit(1)
			}
			err = rrpt.DelinquencyTextReport(&xbiz, &dt)
			if err != nil {
				fmt.Printf("Delinquency text report error: %s\n", err.Error())
			}
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
