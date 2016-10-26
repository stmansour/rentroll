package main

import (
	"rentroll/rlib"
	"time"
)

// XLedger has ledger fields plus LedgerMarker fields for supplied time range
type XLedger struct {
	G  rlib.GLAccount
	LM rlib.LedgerMarker
}

// UILedger contains ledger info for a user interface
type UILedger struct {
	Balance float64 // sum of all LM Balances
	BID     int64
	XL      []XLedger // all ledgers in this business
}

// RRuiSupport is a structure of data that will be passed to all html pages.
// It is the responsibility of the page function to populate the data needed by
// the page. The recommendation is to populate only the data needed.
type RRuiSupport struct {
	DtStart       string          // start of period of interest
	D1            time.Time       // time.Time value for DtStart
	DtStop        string          // end of period of interest
	D2            time.Time       // time.Time value for DtStop
	B             rlib.Business   // business associated with this report
	BL            []rlib.Business // array of all businesses, for initializing dropdown selections
	LDG           UILedger        // ledgers associated with this report
	ReportContent string          // text report content
	PgHnd         []RRPageHandler // the list of reports and handlers
	PageTitle     string          // set page title via software
}

//========================================================================================================

// LMSum takes an array of LedgerMarkers, sums the Balance value of each, and returns the sum.
// The summing skips shadow RA balance accounts
func LMSum(m *[]XLedger) float64 {
	bal := float64(0)
	for _, v := range *m {
		bal += v.LM.Balance
	}
	return bal
}

// UIInitBizList is used to fill out the array of businesses that can be used in UI forms
func UIInitBizList(ui *RRuiSupport) {
	var err error
	ui.BL, err = rlib.GetAllBusinesses()
	if err != nil {
		rlib.Ulog("UIInitBizList: err = %s\n", err.Error())
	}
	// DEBUGGING
	// for i := 0; i < len(ui.BL); i++ {
	// 	fmt.Printf("ui.BL[%d] = %#v\n", i, ui.BL[i])
	// }
}

// UIInitUISupport sets the ui context structure value for page handlers equal to App.PageHandlers
func UIInitUISupport(ui *RRuiSupport) {
	ui.PgHnd = App.PageHandlers
}

// // BuildXLedgerList initializes all ledger information for use in the UI. It loads all defined GLAccounts
// // and the LedgerMarkers for a specific period
// func BuildXLedgerList(ui *RRuiSupport, bid int64, dt time.Time) {
// 	m := rlib.GetAllLedgerMarkersOnOrBefore(bid, &dt) // map of ledger markers indexed by LID
// 	n := rlib.GetLedgerList(bid)                      // list of all ledgers
// 	k := 0
// 	for i := 0; i < len(n); i++ {
// 		var x XLedger
// 		x.G = n[i]
// 		x.LM = m[n[i].LID]
// 		if n[i].LID == 0 {
// 			fmt.Printf("found LID == 0\n")
// 			debug.PrintStack()
// 		}
// 		ui.LDG.XL = append(ui.LDG.XL, x)
// 		k++
// 	}

// 	ui.LDG.Balance = LMSum(&ui.LDG.XL)
// 	ui.LDG.BID = bid
// }
