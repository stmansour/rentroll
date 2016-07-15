package main

import (
	"fmt"
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

// RRuiSupport is a structure of data that will be passed to all html pages.
// It is the responsibility of the page function to populate the data needed by
// the page. The recommendation is to populate only the data needed.
type RRuiSupport struct {
	DtStart time.Time       // start of period of interest
	DtStop  time.Time       // end of period of interest
	B       rlib.Business   // business associated with this report
	BL      []rlib.Business // array of all businesses, for initializing dropdown selections
	LDG     UILedger        // ledgers associated with this report
}

//========================================================================================================

// LMSum takes an array of LedgerMarkers, sums the Balance value of each, and returns the sum
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

// BuildXLedgerList initializes all ledger information for use in the UI. It loads all defined GLAccounts
// and the LedgerMarkers for a specific period
func BuildXLedgerList(ui *RRuiSupport, bid int64, d1, d2 time.Time) {
	m := rlib.GetAllLedgerMarkersInRange(bid, &d1, &d2) // map of ledger markers indexed by LID
	n := rlib.GetLedgerList(bid)                        // list of all ledgers
	k := 0
	for i := 0; i < len(n); i++ {
		var x XLedger
		x.G = n[i]
		x.LM = m[n[i].LID]
		ui.LDG.XL = append(ui.LDG.XL, x)
		k++
	}

	ui.LDG.Balance = LMSum(&ui.LDG.XL)
	ui.LDG.BID = bid
}

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

// GetStatementData returns an array of StatementEntry structs that are part of the statement for the supplied
// time range.
func GetStatementData(xbiz *rlib.XBusiness, raid int64, d1, d2 *time.Time) []StatementEntry {
	var m []StatementEntry

	// OPENING BALANCE - WE MUST SET THIS ONE FIRST - it must be element [0] in the array
	l, err := rlib.GetRABalanceLedger(xbiz.P.BID, raid) // get the GLAccount for this Rental Agreement
	if err != nil {
		rlib.Ulog("Error getting GLAccount for BID=%d, RAID=%d,  err = %s\n", xbiz.P.BID, raid, err.Error())
		return m
	}

	dtStop := d1
	dtStart := dtStop.AddDate(0, 0, -1)
	lm, err := rlib.GetLedgerMarkerByLIDDateRange(xbiz.P.BID, l.LID, &dtStart, dtStop)
	var se StatementEntry
	se.bal = float64(0)
	if nil == err {
		se.bal = lm.Balance
	}
	se.t = 3
	m = append(m, se)

	// ASSESSMENTS
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByRAID.Query(raid, d2, d1)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a rlib.Assessment
		rlib.ReadAssessment(rows, &a)
		dl := a.GetRecurrences(d1, d2)
		if len(dl) > 0 {
			var se StatementEntry
			se.t = 1
			se.a = &a
			m = append(m, se)
		}
	}
	rlib.Errcheck(rows.Err())

	// RECEIPTS
	t := rlib.GetReceiptsInRAIDDateRange(xbiz.P.BID, raid, d1, d2)
	for i := 0; i < len(t); i++ {
		var se StatementEntry
		se.t = 2
		se.r = &t[i]
		m = append(m, se)
	}
	return m
}
