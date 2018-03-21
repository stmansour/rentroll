package rlib

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// PayorStatementInfo summarizes information about a particular
// RentalAgreement in a payor statement
type PayorStatementInfo struct {
	RAB []RAAcctBal
	RL  []ReceiptListEntry // list of receipts for the RentalAgreement(s) associated with a payor
}

// PayorsStatement returns a slice of RAStmtEntry structs that describe
// a Statement for the payors whose TCIDs are listed in payors
//
// Parameters
//  ctx    - context for db txn
//  bid    - which business
//	payors - a slice containing all the payors to be included in the Statement
//  d1,d2  - start and stop of the time period for the statement
//
// Returns  an array of RAAcctBal records
//          any error that occurred or nil if no errors
//---------------------------------------------------------------------------
func PayorsStatement(ctx context.Context, bid int64, payors []int64, d1, d2 *time.Time) (PayorStatementInfo, error) {
	var err error
	var psi PayorStatementInfo

	// Console("PayorsStatement:  bid = %d, d1 = %s, d2 = %s\n", bid, d1.Format(RRDATEREPORTFMT), d2.Format(RRDATEREPORTFMT))
	// for i := 0; i < len(payors); i++ {
	// Console("payors[%d] = %d\n", i, payors[i])
	// }

	//-------------------------------------------------------------
	// Build the list of Rental Agreements this report will cover
	//-------------------------------------------------------------
	var ram = map[int64]int{}
	for i := 0; i < len(payors); i++ {
		m, err := GetRentalAgreementsByPayorRange(ctx, bid, payors[i], d1, d2)
		if err != nil {
			return psi, err
		}

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

	for i := 0; i < len(ras); i++ {
		rasi, err := GetRAIDStatementInfo(ctx, ras[i], d1, d2)
		if err != nil {
			return psi, err
		}
		// Console("got rasi for %d.  len(rasi.Stmt) = %d\n", ras[i], len(rasi.Stmt))
		psi.RAB = append(psi.RAB, rasi)
	}
	psi.RL, err = ReceiptSummary(ctx, ras, d1, d2)

	return psi, err
}

// ReceiptSummary returns a slice of receipts made by all Payors
// responsible for the supplied list of RentalAgreements
//
// INPUTS
//  ctx      - context for db txn
//  raidlist - slice of RAIDs for which we must find the payors
//  d1,d2  - start and stop of the time period for the statement
//
// Returns  an array of RAAcctBal records
//          any error that occurred or nil if no errors
//---------------------------------------------------------------------------
func ReceiptSummary(ctx context.Context, raidlist []int64, d1, d2 *time.Time) ([]ReceiptListEntry, error) {
	var (
		bid int64
		rl  = []ReceiptListEntry{}
		err error
	)

	pm := map[int64]int{}
	for i := 0; i < len(raidlist); i++ {
		m, err := GetRentalAgreementPayorsInRange(ctx, raidlist[i], d1, d2)
		if err != nil {
			return rl, err
		}
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

	q := `SELECT
		{{.SelectClause}}
	FROM Receipt
	WHERE
		((FLAGS & 4)=0) AND BID={{.BID}} AND "{{.d1}}" <= Dt AND Dt < "{{.d2}}" AND ({{.Payors}})
	ORDER BY Dt ASC;`

	qc := QueryClause{
		"SelectClause": RRdb.DBFields["Receipt"],
		"BID":          strconv.FormatInt(bid, 10),
		"d1":           d1.Format(RRDATEFMTSQL),
		"d2":           d2.Format(RRDATEFMTSQL),
		"Payors":       strings.Join(pl, " OR "),
	}

	qry := RenderSQLQuery(q, qc)

	Console("PAYOR STATEMENT Receipt Query:  %s\n", qry)
	rows, err := RRdb.Dbrr.Query(qry)
	if err != nil {
		return rl, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Receipt
		err := ReadReceipts(rows, &r)
		if err != nil {
			return rl, err
		}

		_, alloc, unalloc, err := GetReceiptAllocationAmountsOnDate(ctx, r.RCPTID, d2)
		if err != nil {
			return rl, err
		}

		x := ReceiptListEntry{
			R:           r,
			Allocated:   alloc,
			Unallocated: unalloc,
		}
		rl = append(rl, x)
	}

	return rl, err
}
