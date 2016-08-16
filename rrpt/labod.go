package rrpt

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// LdgAcctBalOnDateTextReport generates the balance of a particular ledger for the supplied date filtered by the RentalAgreement ID
func LdgAcctBalOnDateTextReport(xbiz *rlib.XBusiness, lid, raid int64, dt *time.Time) {
	bal := rlib.GetAccountBalanceForDate(xbiz.P.BID, lid, raid, dt)
	fmt.Printf("Account Balance of Ledger L%08d (%s) for RA%08d as of %s:  %s\n",
		lid, rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[lid].Name, raid, dt.Format(rlib.RRDATEFMT4), rlib.RRCommaf(bal))
}
