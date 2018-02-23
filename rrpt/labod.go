package rrpt

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"time"
)

// LdgAcctBalOnDateTextReport generates the balance of a particular ledger for the supplied date filtered by the RentalAgreement ID
func LdgAcctBalOnDateTextReport(ctx context.Context, xbiz *rlib.XBusiness, lid, raid int64, dt *time.Time) {
	bal, err := rlib.GetRAAccountBalance(ctx, xbiz.P.BID, lid, raid, dt)
	if err != nil {
		fmt.Printf("Error while getting balance for RA%08d - %s\n", raid, err.Error())
	}
	fmt.Printf("Account Balance of Ledger L%08d (%s) for RA%08d as of %s:  %s\n",
		lid, rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[lid].Name, raid, dt.Format(rlib.RRDATEFMT4), rlib.RRCommaf(bal))
}

// RAAccountActivityRangeDetail generates a report of the ledger entries that affect the RentalAgreements ledger during d1-d2
func RAAccountActivityRangeDetail(ctx context.Context, xbiz *rlib.XBusiness, lid, raid int64, d1, d2 *time.Time) {
	var bal = float64(0)
	m, err := rlib.GetLedgerEntriesForRAID(ctx, d1, d2, raid, lid)
	if err != nil {
		fmt.Printf("RAAccountActivityRangeDetail: GetLedgerEntriesForRAID returned error: %s\n", err.Error())
		return
	}
	fmt.Printf("ACCOUNT ACTIVITY\n")
	fmt.Printf("Rental Agreement: RA%08d\n", raid)
	fmt.Printf("Account:  %s  (L%08d)\n\n", rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[lid].Name, lid)

	fmt.Printf("%10s  %8s  %10s  %9s  %10s  %s\n", "Date", "AMOUNT", "LEID", "JID", "JAID", "Comment")
	fmt.Printf("%10s  %8s  %10s  %9s  %10s  %s\n",
		rlib.Tline(10), rlib.Tline(8), rlib.Tline(10), rlib.Tline(9), rlib.Tline(10), rlib.Tline(10))
	for i := 0; i < len(m); i++ {
		fmt.Printf("%10s  %8.2f  LE%08d  J%08d  JA%08d  %s\n",
			m[i].Dt.Format(rlib.RRDATEFMT4), m[i].Amount, m[i].LEID, m[i].JID, m[i].JAID, m[i].Comment)
		bal += m[i].Amount
	}
	s := rlib.Tline(10 + 8 + 10 + 9 + 10 + 10 + (2 * 7))
	fmt.Printf("%s\n", s)
	fmt.Printf("%10s  %8.2f\n", "Total", bal)

	fmt.Println()
	bal1, err := rlib.GetRAAccountBalance(ctx, xbiz.P.BID, lid, raid, d1)
	if err != nil {
		fmt.Printf("RAAccountActivityRangeDetail: GetRAAccountBalance returned error: %s\n", err.Error())
		return
	}

	bal2, err := rlib.GetRAAccountBalance(ctx, xbiz.P.BID, lid, raid, d2)
	if err != nil {
		fmt.Printf("RAAccountActivityRangeDetail: GetRAAccountBalance returned error: %s\n", err.Error())
		return
	}

	// handle end date inclusion on "d2"
	if rlib.EDIEnabledForBID(xbiz.P.BID) {
		*d2 = d2.AddDate(0, 0, -1)
	}

	fmt.Printf("Account Balance on %10s  -  %10s\n", d1.Format(rlib.RRDATEFMT4), rlib.RRCommaf(bal1))
	fmt.Printf("Account Balance on %10s  -  %10s\n", d2.Format(rlib.RRDATEFMT4), rlib.RRCommaf(bal2))
	fmt.Printf("Change ---> %8.2f\n", bal2-bal1)
}
