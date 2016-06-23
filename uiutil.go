package main

import (
	"rentroll/rlib"
	"time"

	"github.com/dustin/go-humanize"
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
	DtStart time.Time // start of period of interest
	DtStop  time.Time // end of period of interest
	B       rlib.Business
	LDG     UILedger
}

// RRCommaf returns a floating point number formated with commas for every 3 orders of magnitude
// and 2 points after the decimal
func RRCommaf(x float64) string {
	return humanize.FormatFloat("#,###.##", x)
}

// LMSum takes an array of LedgerMarkers, sums the Balance value of each, and returns the sum
func LMSum(m *[]XLedger) float64 {
	bal := float64(0)
	for _, v := range *m {
		bal += v.LM.Balance
	}
	return bal
}

// BuildXLedgerList initializes all ledger information for use in the UI. It loads all defined GLAccounts
// and the LedgerMarkers for a specific period
func BuildXLedgerList(ui *RRuiSupport, bid int64, d1, d2 time.Time) {
	m := rlib.GetAllLedgerMarkersInRange(bid, &d1, &d2) // map of ledger markers indexed by LID
	n := rlib.GetLedgerList(bid)                        // list of all ledgers
	k := 0
	for i := 0; i < len(n); i++ {
		// if n[i].Type == 1 {
		// 	continue
		// }
		// fmt.Printf("n[%d] = %d  %s  %s  type=%d\n", i, n[i].BID, n[i].Name, n[i].GLNumber, n[i].Type)
		var x XLedger
		x.G = n[i]
		x.LM = m[n[i].LID]
		ui.LDG.XL = append(ui.LDG.XL, x)
		k++
	}

	ui.LDG.Balance = LMSum(&ui.LDG.XL)
	ui.LDG.BID = bid
}
