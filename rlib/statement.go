package rlib

import (
	"context"
	"fmt"
	"sort"
	"time"
)

// RAStmtEntry describes an entry on a statement
type RAStmtEntry struct {
	T       int                // 1 = assessment, 2 = Receipt
	A       *Assessment        // for type==1, the pointer to the assessment
	R       *ReceiptAllocation // for type ==2, the pointer to the receipt
	RNT     *Rentable          // the associated rentable, if known
	Amt     float64            // amount of the receipt or assessment
	Reverse bool               // is this a reversal?
	Dt      time.Time          // date/time of this assessment or receipt
	TCID    int64              // IF THIS IS FOR A PAYOR STATEMENT, the TCID of the Payor, otherwise 0
}

// ReceiptListEntry shows the receipts for all RentalAgreements in
// a payors statement
type ReceiptListEntry struct {
	R           Receipt
	Allocated   float64
	Unallocated float64
}

// RAStmtEntries is needed to sort the array
type RAStmtEntries []RAStmtEntry

// Len returns the size of the array
func (slice RAStmtEntries) Len() int {
	return len(slice)
}

// Less returns true if element i comes before element j
func (slice RAStmtEntries) Less(i, j int) bool {
	//return slice[i].Name < slice[j].Name;
	return slice[i].Dt.Before(slice[j].Dt)
}

// Swap swaps the two entries
func (slice RAStmtEntries) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// RAAcctBal contains details about the account balance for a Rental Agreement
type RAAcctBal struct {
	DtStart    time.Time     // Period Start
	DtStop     time.Time     // Period Stop -- up to but not including
	LmStart    LedgerMarker  // this is the starting point for the calculations
	Gap        RAStmtEntries // these entries cover the gap between the LmStart and Period DtStart
	OpeningBal float64       // balance at the open of period DtStart
	Stmt       RAStmtEntries // these are the actual statement entries
	ClosingBal float64       // balance at close of period
	RAID       int64         // which RentalAgreement is this for
}

// GetNameFromTransactantCache implements a simple cache of Transactants
func GetNameFromTransactantCache(ctx context.Context, tcid int64, payorcache map[int64]Transactant) string {
	p, ok := payorcache[tcid]
	if ok {
		return p.GetUserName()
	}
	var tr Transactant
	err := GetTransactant(ctx, tcid, &tr)
	if err != nil {
		return ""
	}

	payorcache[tr.TCID] = tr
	return tr.GetUserName()
}

// GetRAIDBalance returns the balance of the account for the supplied
// rental agreement on the date requested.
//
// Parameters
//     raid  = Rental Agreement ID
//       dt  = date for which balance is needed
//
// Returns
//  balance  = RAID account balance if err == nil
//      err  = any error that occurred or nil if no errors
//
//=============================================================================
func GetRAIDBalance(ctx context.Context, raid int64, dt *time.Time) (float64, error) {
	bal := float64(0)
	lm, err := GetRALedgerMarkerOnOrBefore(ctx, raid, dt)
	if lm.LMID == 0 {
		err := fmt.Errorf("*** ERROR ***  could not find ledger marker for RAID %d on or before %s", raid, dt.Format(RRDATEFMTSQL))
		return bal, err
	}
	if err != nil {
		err := fmt.Errorf("*** ERROR ***  could not find ledger marker for RAID %d on or before %s", raid, dt.Format(RRDATEFMTSQL))
		return bal, err
	}

	var rs RAStmtEntries
	bal = lm.Balance                                       // initialize
	b, err := GetRAIDAcctRange(ctx, raid, &lm.Dt, dt, &rs) // update with total for this range
	if err != nil {
		return bal, err
	}
	bal += b
	return bal, err
}

// GetRAIDStatementInfo is written in a way that will work for cash based
// systems or accrual based systems. It looks at all the transactions
// involving the RAID provided and computes a total.  The total is computed
// up-to-and-including d2.
//
//     |<----Gap--->|
//  ---+------------+-----------------------------------------------------
//     |            |
//   LmStart        d1
//
// Parameters
//     raid  = Rental Agreement ID
//     d1,d2 = date range for which balance is computed
//
// Returns
//  RAAcctBal which has the following members:
//     LmStart
//         is the starting balance for calculations - LmStart.Balance is
//         the opening balance for LmStart.Dt, which is the nearest date on
//         or before d1.
//
//     Gap
//         is the list of Assessments and ReceiptAllocations that occurred
//         from LmStart.Dt up to (but not including) d1
//
//     OpeningBal
//         is the opening balance on d1.  The sum of LmStart.Balance and all
//         entries in Gap.
//
//     Stmt
//         is the list of Assessments and ReceiptAllocations that occurred
//         during the period d1 up to (but not including) d2
//
//     RL  is not filled in by this report.
//
//=============================================================================
func GetRAIDStatementInfo(ctx context.Context, raid int64, d1, d2 *time.Time) (RAAcctBal, error) {
	var (
		err error
		m   RAAcctBal
	)

	m.DtStart = *d1
	m.DtStop = *d2
	m.RAID = raid

	//----------------------------------------------------------------
	//  First, find the ledger marker for this RentalAgreement...
	//----------------------------------------------------------------
	m.LmStart, err = GetRALedgerMarkerOnOrBefore(ctx, raid, d1)
	if err != nil {
		return m, err
	}
	if m.LmStart.LMID == 0 { // if there's no marker on or prior to d1
		m.LmStart, err = GetRALedgerMarkerOnOrAfter(ctx, raid, d1) // see where the first marker happens
		if err != nil {
			return m, err
		}
		if m.LmStart.LMID == 0 {
			err = fmt.Errorf("*** ERROR ***  could not find ledger marker for RAID %d", raid)
			return m, err
		}
		if m.LmStart.Dt.After(*d2) { // if no find, then there is no ledger information for the supplied date range
			return m, nil // not really an error, but there's no data for this time range
		}
	}

	m.OpeningBal = m.LmStart.Balance                                  // initialize
	ob, err := GetRAIDAcctRange(ctx, raid, &m.LmStart.Dt, d1, &m.Gap) // update with total for this range
	if err != nil {
		return m, err
	}
	m.OpeningBal += ob
	sort.Sort(m.Gap)

	//----------------------------------------------------------------
	// Now get the actual Statement data and balance...
	//----------------------------------------------------------------
	m.ClosingBal = m.OpeningBal                             // initialize
	cb, err := GetRAIDAcctRange(ctx, raid, d1, d2, &m.Stmt) // update with total for the statement range
	if err != nil {
		return m, err
	}
	sort.Sort(m.Stmt)
	m.ClosingBal += cb

	return m, err
}

// GetRAIDAcctRange gets the assessment and receipt allocation entries for the
// supplied time range and returns the balance of these entries.
//=============================================================================
func GetRAIDAcctRange(ctx context.Context, raid int64, d1, d2 *time.Time, p *RAStmtEntries) (float64, error) {
	bal := float64(0)
	//----------------------------------------------------------------
	// Total all assessments in the supplied range that involve RAID.
	//----------------------------------------------------------------
	rows, err := RRdb.Prepstmt.GetAssessmentsByRAIDRange.Query(raid, d1, d2)
	if err != nil {
		return bal, err
	}
	defer rows.Close()

	for rows.Next() {
		var a Assessment
		err := ReadAssessments(rows, &a)
		if err != nil {
			return bal, err
		}

		var rnt Rentable
		err = GetRentableByID(ctx, a.RID, &rnt)
		if err != nil {
			return bal, err
		}

		se := RAStmtEntry{
			T:       1,
			A:       &a,
			Amt:     a.Amount,
			Dt:      a.Start,
			RNT:     &rnt,
			Reverse: a.FLAGS&0x4 != 0, // bit 2 is the reversal flag
		}
		(*p) = append((*p), se)
		if !se.Reverse {
			bal += se.Amt // if it is a reversal, do
		}
		// Console("ASMID  = %3d,  se.Amt = %8.2f,  bal = %8.2f,  Reverse = %t\n", a.ASMID, se.Amt, bal, se.Reverse)
	}

	//----------------------------------------------------------------
	// Total all receipts in the supplied range that involve RAID.
	//----------------------------------------------------------------
	rows, err = RRdb.Prepstmt.GetASMReceiptAllocationsInRAIDDateRange.Query(raid, d1, d2)
	if err != nil {
		return bal, err
	}
	defer rows.Close()

	for rows.Next() {
		var ra ReceiptAllocation
		err := ReadReceiptAllocations(rows, &ra)
		if err != nil {
			return bal, err
		}

		a, err := GetAssessment(ctx, ra.ASMID)
		if err != nil {
			return bal, err
		}

		var rnt Rentable
		err = GetRentableByID(ctx, a.RID, &rnt)
		if err != nil {
			return bal, err
		}

		se := RAStmtEntry{
			T:       2,
			R:       &ra,
			A:       &a,
			RNT:     &rnt,
			Amt:     ra.Amount,
			Dt:      ra.Dt,
			Reverse: ra.FLAGS&0x4 != 0, // bit 2 is the reversal flag
		}
		(*p) = append((*p), se)
		if !se.Reverse {
			bal -= se.Amt
		}
		// Console("RCPTID = %3d,  se.Amt = %8.2f,  bal = %8.2f,  Reverse = %t (RCPAID=%d),  ASMID = %3d\n", se.R.RCPTID, se.Amt, bal, se.Reverse, ra.RCPAID, se.A.ASMID)
	}
	return bal, err
}
