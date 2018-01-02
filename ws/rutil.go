package ws

import (
	"html/template"
	"rentroll/rlib"
	"time"
)

// RRfuncMap is a map of functions passed to each html page that can be referenced
// as needed to produce the page
var RRfuncMap map[string]interface{}

// ReportContext is a structure of data that will be passed to all html pages.
// It is the responsibility of the page function to populate the data needed by
// the page. The recommendation is to populate only the data needed.
type ReportContext struct {
	Language           string          // what language
	Template           string          // which template
	DtStart            string          // start of period of interest
	D1                 time.Time       // time.Time value for DtStart
	DtStop             string          // end of period of interest
	D2                 time.Time       // time.Time value for DtStop
	B                  rlib.Business   // business associated with this report
	BL                 []rlib.Business // array of all businesses, for initializing dropdown selections
	ReportContent      string          // text report content
	PageTitle          string          // set page title via software
	ReportOutputFormat int             // indicates text, html, or pdf
	PDFPageWidth       float64         // page width
	PDFPageHeight      float64         // page height
	PDFPageSizeUnit    string          // page size unit, default is inch ("in")
	EDI                int             // end date inclusive - 0 = end date is not inclusive, 1 = end date is inclusive
	// LDG                UILedger        // ledgers associated with this report
}

// InitReports initializes the reports subsystem. Historically it
// did more than it does today. Currently, it initializes the map
// of functions that can be used by a startup web page such as home.html
// or rhome.html
//-----------------------------------------------------------------------------
func InitReports() {
	RRfuncMap = template.FuncMap{
		"DateToString": rlib.DateToString,
		"GetVersionNo": GetVersionNo,
		"getBuildTime": GetBuildTime,
		"RRCommaf":     rlib.RRCommaf,
	}
}

//========================================================================================================

// LMSum takes an array of LedgerMarkers, sums the Balance value of each, and returns the sum.
// The summing skips shadow RA balance accounts
// func LMSum(m *[]XLedger) float64 {
// 	bal := float64(0)
// 	for _, v := range *m {
// 		bal += v.LM.Balance
// 	}
// 	return bal
// }

// // UIInitBizList is used to fill out the array of businesses that can be used in UI forms
// func UIInitBizList(ui *ReportContext) {
// 	var err error
// 	ui.BL, err = rlib.GetAllBusinesses()
// 	if err != nil {
// 		rlib.Ulog("UIInitBizList: err = %s\n", err.Error())
// 	}
// 	// DEBUGGING
// 	// for i := 0; i < len(ui.BL); i++ {
// 	// 	fmt.Printf("ui.BL[%d] = %#v\n", i, ui.BL[i])
// 	// }
// }
// // XLedger has ledger fields plus LedgerMarker fields for supplied time range
// type XLedger struct {
// 	G  rlib.GLAccount
// 	LM rlib.LedgerMarker
// }
// // UILedger contains ledger info for a user interface
// type UILedger struct {
// 	Balance float64 // sum of all LM Balances
// 	BID     int64
// 	XL      []XLedger // all ledgers in this business
// }
