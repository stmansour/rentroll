package rlib

import (
	"fmt"
	"time"
)

// RARStmtEntry descriGetRARBalancebes an entry on a statement
type RARStmtEntry struct {
	T       int                // 1 = assessment, 2 = Receipt
	A       *Assessment        // for type==1, the pointer to the assessment
	R       *ReceiptAllocation // for type ==2, the pointer to the receipt
	Amt     float64            // amount of the receipt or assessment
	Reverse bool               // is this a reversal?
	Dt      time.Time          // date/time of this assessment or receipt
	TCID    int64              // IF THIS IS FOR A PAYOR STATEMENT, the TCID of the Payor, otherwise 0
}

// AcctRcvAccts returns a slice of AccountsReceivable accounts
//
// PARAMS
//     bid - which business
//
// RETURNS
// []int64 - slice of LIDs that are of type Accounts Receivable
//   error - any error encountered
//-----------------------------------------------------------------------------
// func AcctRcvAccts(bid int64) ([]int64, error) {
// 	return AcctSlice(bid, AccountsReceivable)
// }

// GetBeginEndRARBalance gets the balance associated with a Rentable and a
// Rental Agreement at a particular point in time.
//
// INPUTS
//   rid  - RID of Rentable
//   raid - RAID of Rental Agreement
//   d1   - time for which balance is requested
//   d2   - time for which balance is requested
//
// RETURNS
//   float64 - the balance for the Rentable rid in Rental Agreement raid at
//             time dt
//   error   - any error encountered
//-----------------------------------------------------------------------------
func GetBeginEndRARBalance(bid, rid, raid int64, d1, d2 *time.Time) (float64, float64, error) {
	var err error
	begin := float64(0)
	end := float64(0)
	begin, err = GetRARBalance(bid, rid, raid, d1)
	end, err = GetRARBalance(bid, rid, raid, d2)
	return begin, end, err
}

// GetRARBalance gets the balance associated with a Rentable and a
// Rental Agreement at a particular point in time.
//
// INPUTS
//   bid     - biz id
//   rid     - RID of Rentable
//   raid    - RAID of Rental Agreement
//   dt      - time for which balance is requested
//
// RETURNS
//   float64 - the balance for the Rentable rid in Rental Agreement raid at
//             time dt
//   error   - any error encountered
//-----------------------------------------------------------------------------
func GetRARBalance(bid, rid, raid int64, dt *time.Time) (float64, error) {
	funcname := "GetRARBalance"
	bal := float64(0)

	lm := GetRARentableLedgerMarkerOnOrBefore(raid, rid, dt)
	if lm.LMID == 0 {
		LogAndPrint("%s: Could not find LedgerMarker for RAID=%d, RID=%d, on or before %s\n",
			funcname, raid, rid, dt.Format(RRDATEFMT3))
		return bal, nil
	}

	//------------------------------------------------------------------
	// Get all the assessments and payments for this RAID, RID pair...
	//------------------------------------------------------------------
	bal = lm.Balance + GetRARAcctRange(bid, raid, rid, &lm.Dt, dt)
	return bal, nil
}

// GetRARAcctRange returns the change in balance for the supplie RAID,RID
// combination over the supplied time range.
//
// INPUTS
//   raid - RAID of Rental Agreement
//   rid  - RID of Rentable
//   d1   - time for which balance is requested
//   d2   - time for which balance is requested
//
// RETURNS
//   float64 - the balance for the Rentable rid in Rental Agreement raid at
//             time dt
//   error   - any error encountered
//-----------------------------------------------------------------------------
func GetRARAcctRange(bid, raid, rid int64, d1, d2 *time.Time) float64 {
	funcname := "GetRARAcctRange"
	// Console("Entered %s\n", funcname)
	bal := float64(0)

	acctRules := ""
	rcvAccts, err := AcctSlice(bid, AccountsReceivable)
	if err != nil {
		LogAndPrintError(funcname, err)
		return bal
	}
	if len(rcvAccts) == 0 {
		LogAndPrintError(funcname, fmt.Errorf("GetRARAcctRange: there are no accounts of type %s", AccountsReceivable))
		return bal
	}
	qryAccts, err := AcctRulesSlice(rcvAccts)
	if nil == err {
		l := len(qryAccts)
		if 0 > l {
			acctRules = " AND ("
			for i := 0; i < l; i++ {
				acctRules += fmt.Sprintf("ARID=%d", qryAccts[i])
				if i+1 < l {
					acctRules += " OR "
				}
			}
			acctRules += ")"
		}
	} else {
		LogAndPrintError(funcname, err)
	}

	q := fmt.Sprintf("SELECT %s FROM Assessments WHERE (RentCycle=0  OR (RentCycle>0 AND PASMID>0)) AND RAID=%d AND RID=%d AND Stop>=%q AND Start<%q %s",
		RRdb.DBFields["Assessments"], raid, rid, d1.Format(RRDATEFMTSQL), d2.Format(RRDATEFMTSQL), acctRules)
	// Console("q = %s\n", q)
	rows, err := RRdb.Dbrr.Query(q)
	Errcheck(err)
	defer rows.Close()

	// Console("GetRARAcctRange: query = %s\n", q)

	//------------------------------------------------------------------------
	// Total all assessments in the supplied range that involve RID in RAID.
	//------------------------------------------------------------------------
	for rows.Next() {
		var a Assessment
		ReadAssessments(rows, &a)
		if 0 == a.FLAGS&0x4 { // if this is not a reversal...
			bal += a.Amount // ... then add it to the balance
		}
		// Console("\tASMID = %d, FLAGS=%x  Amount = %.2f,  bal = %.2f\n", a.ASMID, a.FLAGS, a.Amount, bal)

		//----------------------------------------------------------------
		// Total all receipts applied toward this ASMID
		//----------------------------------------------------------------
		innerRows, err := RRdb.Prepstmt.GetASMReceiptAllocationsInRARDateRange.Query(raid, a.ASMID, d1, d2)
		Errcheck(err)
		defer innerRows.Close()
		for innerRows.Next() {
			var ra ReceiptAllocation
			ReadReceiptAllocations(innerRows, &ra)
			bal -= ra.Amount
			// Console("\tRCPAID = %d, Amount = %.2f,  bal = %.2f\n", ra.RCPAID, ra.Amount, bal)
		}
	}
	// Console("---------->>>>> RETURNING BALANCE = %.2f\n", bal)
	return bal
}
