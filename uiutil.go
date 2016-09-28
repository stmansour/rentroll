package main

import (
	"fmt"
	"net/http"
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

// RTIDCount is for counting rentables of a particular type
type RTIDCount struct {
	RT    rlib.RentableType // ID of the types we're counting
	Count int64             // the count
}

// StatementEntry is a struct containing references to an Assessment or a Receipt that is
// part of a billing statement associated with a RentalAgreement
type StatementEntry struct {
	t   int              // type: 1 = assessment, 2 = Receipt, 3 = Initial Balance
	a   *rlib.Assessment // for type==1, the pointer to the assessment
	r   *rlib.Receipt    // for type ==2, the pointer to the receipt
	bal float64          // opening balance
}

// StmtEntry describes an entry on a statement
type StmtEntry struct {
	t       int   // 1 = assessment, 2 = Receipt, 3 = Initial Balance
	id      int64 // ASMID if t==1, RCPTID if t==2, n/a if t==3
	asmtlid int64 // valid only for t==1, the assessments ATypeLID
	amt     float64
	dt      time.Time
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
}

// RRPageHandler is a structure of page names and handlers
type RRPageHandler struct {
	ReportName   string                                   // report name
	FormPageName string                                   // name of form to collect information for this report
	URL          string                                   // url for this handler
	Handler      func(http.ResponseWriter, *http.Request) // the actual handler function
	// ReportPageName string
	// ReportHandler  string
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

// GetRentableCountByRentableType returns a structure containing the count of Rentables for each RentableType
// in the specified time range
func GetRentableCountByRentableType(xbiz *rlib.XBusiness, d1, d2 *time.Time) ([]RTIDCount, error) {
	var count int64
	var m []RTIDCount
	var err error
	i := 0
	for _, v := range xbiz.RT {
		s := fmt.Sprintf("SELECT COUNT(*) FROM RentableTypeRef WHERE RTID=%d AND DtStop>\"%s\" AND DtStart<\"%s\"",
			v.RTID, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
		err = rlib.RRdb.Dbrr.QueryRow(s).Scan(&count)
		if err != nil {
			fmt.Printf("GetRentableCountByRentableType: query=\"%s\"    err = %s\n", s, err.Error())
		}
		var rc RTIDCount
		rc.Count = count
		rc.RT = v
		var cerr error
		rc.RT.CA, cerr = rlib.GetAllCustomAttributes(rlib.ELEMRENTABLETYPE, v.RTID)
		if cerr != nil {
			if !rlib.IsSQLNoResultsError(cerr) { // it's not really an error if we don't find any custom attributes
				err = cerr
				break
			}
		}

		m = append(m, rc)
		i++
	}
	return m, err
}

// GetStatementData returns an array of StatementEntries for building a statement
func GetStatementData(xbiz *rlib.XBusiness, raid int64, d1, d2 *time.Time) []StmtEntry {
	var m []StmtEntry
	bal := rlib.GetRAAccountBalance(xbiz.P.BID, rlib.RRdb.BizTypes[xbiz.P.BID].DefaultAccts[rlib.GLGENRCV].LID, raid, d1)
	var initBal = StmtEntry{amt: bal, t: 3, dt: *d1}
	m = append(m, initBal)
	n, err := rlib.GetLedgerEntriesForRAID(d1, d2, raid, rlib.RRdb.BizTypes[xbiz.P.BID].DefaultAccts[rlib.GLGENRCV].LID)
	if err != nil {
		return m
	}
	for i := 0; i < len(n); i++ {
		var se StmtEntry
		se.amt = n[i].Amount
		se.dt = n[i].Dt
		j, err := rlib.GetJournal(n[i].JID)
		if err != nil {
			return m
		}
		se.t = int(j.Type)
		se.id = j.ID
		if se.t == rlib.JOURNALTYPEASMID {
			// read the assessment to find out what it was for...
			a, err := rlib.GetAssessment(se.id)
			if err != nil {
				return m
			}
			se.asmtlid = a.ATypeLID
		}
		m = append(m, se)
	}
	return m
}
