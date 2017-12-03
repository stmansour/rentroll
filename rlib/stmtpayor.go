package rlib

import (
	"fmt"
	"strings"
	"time"
)

// PayorStatementInfo summarizes information about a particular
// RentalAgreement in a payor statement
type PayorStatementInfo struct {
	RAB []RAAcctBal
	RL  []ReceiptListEntry // list of receipts for the RentalAgreement(s) associated with a payor
}

// group payor statement
// First, collect every Rental Agreement we need to check...
// Next, get all the assessments made to any of those Rental Agreements
// during the period d1-d2
// Next, get all the receipts that each payor made during d1-d2
// sort the entire list by date

// PayorsStatement returns a slice of RAStmtEntry structs that describe
// a Statement for the payors whose TCIDs are listed in payors
//
// Parameters
//  bid    - which business
//	payors - a slice containing all the payors to be included in the Statement
//  d1,d2  - start and stop of the time period for the statement
//
// Returns  an array of RAAcctBal records
//          any error that occurred or nil if no errors
//---------------------------------------------------------------------------
func PayorsStatement(bid int64, payors []int64, d1, d2 *time.Time) (PayorStatementInfo, error) {
	// Console("PayorsStatement:  bid = %d, d1 = %s, d2 = %s\n", bid, d1.Format(RRDATEREPORTFMT), d2.Format(RRDATEREPORTFMT))
	// for i := 0; i < len(payors); i++ {
	// Console("payors[%d] = %d\n", i, payors[i])
	// }

	//-------------------------------------------------------------
	// Build the list of Rental Agreements this report will cover
	//-------------------------------------------------------------
	var ram = map[int64]int{}
	for i := 0; i < len(payors); i++ {
		m := GetRentalAgreementsByPayorRange(bid, payors[i], d1, d2)
		for j := 0; j < len(m); j++ {
			ram[m[j].RAID] = 1
			// Console("PayorsStatement: RAID found: %d\n", m[j].RAID)
		}
	}

	var ras Int64Range
	for k := range ram {
		ras = append(ras, k)
	}
	SortInt64(&ras)
	// Console("Sorted RAID array:  %v\n", ras)

	//-------------------------------------------------------------
	// Build the list of
	//-------------------------------------------------------------
	var psi PayorStatementInfo
	var err error
	for i := 0; i < len(ras); i++ {
		rasi, err := GetRAIDStatementInfo(ras[i], d1, d2)
		if err != nil {
			return psi, err
		}
		// Console("got rasi for %d.  len(rasi.Stmt) = %d\n", ras[i], len(rasi.Stmt))
		psi.RAB = append(psi.RAB, rasi)
	}
	psi.RL, err = ReceiptSummary(ras, d1, d2)

	return psi, err
}

// ReceiptSummary returns a slice of receipts made by all Payors
// responsible for the supplied list of RentalAgreements
func ReceiptSummary(raidlist []int64, d1, d2 *time.Time) ([]ReceiptListEntry, error) {
	var bid int64
	rl := []ReceiptListEntry{}
	q := "SELECT " + RRdb.DBFields["Receipt"] + " FROM Receipt WHERE 0=(FLAGS & 4) AND ("
	pm := map[int64]int{}
	for i := 0; i < len(raidlist); i++ {
		m := GetRentalAgreementPayorsInRange(raidlist[i], d1, d2)
		for j := 0; j < len(m); j++ {
			bid = m[j].BID
			pm[m[j].TCID] = 1
		}
	}
	pl := []string{}
	for k := range pm {
		pl = append(pl, fmt.Sprintf("TCID=%d", k))
	}
	if len(pl) == 0 {
		err := fmt.Errorf("No payors found")
		return rl, err
	}
	q += strings.Join(pl, " OR ")
	q += fmt.Sprintf(") AND BID=%d AND %q<=Dt AND Dt<%q ORDER BY Dt ASC", bid, d1.Format(RRDATEFMTSQL), d2.Format(RRDATEFMTSQL))
	// Console("PAYOR STATEMENT Receipt Query:  %s\n", q)
	rows, err := RRdb.Dbrr.Query(q)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var r Receipt
		ReadReceipts(rows, &r)
		_, alloc, unalloc := GetReceiptAllocationAmountsOnDate(r.RCPTID, d2)

		x := ReceiptListEntry{
			R:           r,
			Allocated:   alloc,
			Unallocated: unalloc,
		}
		rl = append(rl, x)
	}
	return rl, nil
}
