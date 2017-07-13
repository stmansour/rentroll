package bizlogic

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// UpdateReceipt accepts an updated rlib.Receipt structure.  It will load the
// existing receipt, compare the fields, and take appropriate action.
//
// In editing a receipt, the following fields require edits to multiple tables:
//   * date
//   * Amount
//   * Account Rule
//
// Editing the date requires updates to
//   * Receipt
//   * ReceiptAllocation
//   * Journal
//   * JournalAllocation
//   * two LedgerEntries
//-------------------------------------------------------------------------------
func UpdateReceipt(rnew *rlib.Receipt) error {
	var err error
	//-------------------------------
	// Load existing receipt...
	//-------------------------------
	// fmt.Printf("bizlogic.UpdateReceipt: A\n")
	rold := rlib.GetReceipt(rnew.RCPTID)
	if rold.RCPTID == 0 {
		return fmt.Errorf("Receipt %d not found", rnew.RCPTID)
	}

	// fmt.Printf("bizlogic.UpdateReceipt: B\n")
	//---------------------------------------------------------------------------------
	// we need to reverse the old receipt if any of the following fields have changed:
	//    * Dt
	//    * Amount
	//    * AccountRule
	//---------------------------------------------------------------------------------
	reverse := (!rold.Dt.Equal(rnew.Dt)) || rold.Amount != rnew.Amount || rold.ARID != rnew.ARID
	if reverse {
		fmt.Printf("bizlogic.UpdateReceipt: C\n")
		err = ReverseReceipt(&rold)
		if err != nil {
			return err
		}
		return InsertReceipt(rnew)
	}

	// fmt.Printf("bizlogic.UpdateReceipt: D\n")
	return rlib.UpdateReceipt(rnew)
}

// ReverseReceipt reverses the payment from the supplied receipt. It links the
// reversal back to the supplied receipt
//-------------------------------------------------------------------------------
func ReverseReceipt(r *rlib.Receipt) error {
	rr := *r
	rr.RCPTID = int64(0)
	rr.Amount = -rr.Amount
	rr.Comment = fmt.Sprintf("Reversal of receipt %s", r.IDtoString())
	rr.PRCPTID = r.RCPTID // link to parent
	rr.RA = []rlib.ReceiptAllocation{}
	return InsertReceipt(&rr)
}

// InsertReceipt adds a new receipt and updates the journal and ledgers
//-------------------------------------------------------------------------------
func InsertReceipt(a *rlib.Receipt) error {
	funcname := "bizlogic.InsertReceipt"
	_, err := rlib.InsertReceipt(a)
	if err != nil {
		return err
	}

	var xbiz rlib.XBusiness
	rlib.InitBizInternals(a.BID, &xbiz)
	ar := rlib.RRdb.BizTypes[a.BID].AR[a.ARID]               // get the AR for this receipt...
	ard := rlib.RRdb.BizTypes[a.BID].GLAccounts[ar.DebitLID] // get GL Account Info for debits and credits
	arc := rlib.RRdb.BizTypes[a.BID].GLAccounts[ar.CreditLID]

	//------------------------------------------------
	// create the receipt allocation
	//------------------------------------------------
	var ra rlib.ReceiptAllocation
	ra.RCPTID = a.RCPTID
	ra.Amount = a.Amount
	ra.AcctRule = fmt.Sprintf("d %s _, c %s _", ard.GLNumber, arc.GLNumber)
	ra.BID = a.BID
	ra.Dt = a.Dt
	rlib.InsertReceiptAllocation(&ra)
	a.RA = append(a.RA, ra)

	//------------------------------------------------
	// Add it to the Journal
	//------------------------------------------------
	d1 := time.Date(a.Dt.Year(), a.Dt.Month(), 1, 0, 0, 0, 0, rlib.RRdb.Zone)
	mon, year := rlib.IncMonths(a.Dt.Month(), int64(a.Dt.Year()))
	d2 := time.Date(int(year), mon, 1, 0, 0, 0, 0, rlib.RRdb.Zone)
	jnl, err := rlib.ProcessNewReceipt(&xbiz, &d1, &d2, a)
	if err != nil {
		e := fmt.Errorf("%s:  Error in rlib.ProcessNewReceipt: %s", funcname, err.Error())
		rlib.Ulog("%s", e.Error())
		return e
	}

	//------------------------------------------------
	// Add it to the Ledgers
	//------------------------------------------------
	rlib.GetJournalAllocations(&jnl)
	rlib.InitLedgerCache()
	rlib.GenerateLedgerEntriesFromJournal(&xbiz, &jnl, &d1, &d2)
	return nil
}
