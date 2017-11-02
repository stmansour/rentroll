package bizlogic

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// EnsureReceiptFundsToDepositoryAccount looks at the AR used
// for the Receipt. If the debit is not
// already the Account associated with the Depository for this
// deposit, then credit that account and debit the Depository
// account with the receipt amount.
//
// @params
//  r - the receipt to be deposited in the Depository
//  a - the assesment that r is for
//  d - the Depository where the funds will be deposited.
//--------------------------------------------------------------
func EnsureReceiptFundsToDepositoryAccount(r *rlib.Receipt, a *rlib.Assessment, d *rlib.Depository) error {
	var xbiz rlib.XBusiness
	var err error
	funcname := "EnsureReceiptFundsToDepositoryAccount"
	rlib.InitBizInternals(r.BID, &xbiz)
	ar := rlib.RRdb.BizTypes[r.BID].AR[r.ARID]
	if ar.DebitLID != d.LID {
		//----------------------------------
		// debit  d.LID r.Amount
		// credit ar.DebitLID r.Amount
		//----------------------------------
		var jnl = rlib.Journal{
			BID:    r.BID,
			Amount: r.Amount,
			Dt:     time.Now(),
			Type:   rlib.JNLTYPEXFER,
			ID:     r.RCPTID,
		}
		_, err = rlib.InsertJournal(&jnl)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		var ja = rlib.JournalAllocation{
			JID: jnl.JID,
			AcctRule: fmt.Sprintf("d %s %.4f, c %s %.4f",
				rlib.RRdb.BizTypes[r.BID].GLAccounts[d.LID].GLNumber, r.Amount,
				rlib.RRdb.BizTypes[r.BID].GLAccounts[ar.DebitLID].GLNumber, r.Amount),
			Amount: r.Amount,
			BID:    r.BID,
			RAID:   r.RAID,
			ASMID:  a.ASMID,
			TCID:   r.TCID,
			RCPTID: r.RCPTID,
		}
		rlib.InsertJournalAllocationEntry(&ja)
		jnl.JA = append(jnl.JA, ja)

		//-------------------------------------------------------------------------
		// Update ledgers based on journal entry
		//-------------------------------------------------------------------------
		var l = rlib.LedgerEntry{
			BID:    jnl.BID,
			JID:    jnl.JID,
			RID:    ja.RID,
			JAID:   ja.JAID,
			RAID:   ja.RAID,
			TCID:   ja.TCID,
			Dt:     jnl.Dt,
			LID:    d.LID,
			Amount: r.Amount,
		}
		_, err = rlib.InsertLedgerEntry(&l)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		l.LID = ar.DebitLID
		l.Amount = -r.Amount
		_, err = rlib.InsertLedgerEntry(&l)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}
	}
	return nil
}

// SaveDeposit validates that all the information in the deposit
// meets the business criteria for a deposit and Inserts or Updates
// the deposit if the criteria are met.
//
// @params
//	a - the depost struct
//  rcpts - an array of RCPTIDs that are being targeted for the deposit.
//
// @returns
//	errlist - an array of errors
//-----------------------------------------------------------------------
func SaveDeposit(a *rlib.Deposit, newRcpts []int64) []BizError {
	rlib.Console("SaveDeposit: 0\n")
	var e []BizError
	var rlist []rlib.Receipt
	tot := float64(0)
	//------------------------------------------------------------
	// First, validate that all newRcpts are eligible for inclusion
	// in this receipt
	//------------------------------------------------------------
	for i := 0; i < len(newRcpts); i++ {
		r := rlib.GetReceipt(newRcpts[i])
		tot += r.Amount
		if r.DID != 0 && r.DID != a.DID {
			s := fmt.Sprintf(BizErrors[ReceiptAlreadyDeposited].Message, rlib.IDtoShortString("RCPT", r.RCPTID), rlib.IDtoShortString("D", r.DID))
			b := BizError{Errno: ReceiptAlreadyDeposited, Message: s}
			e = append(e, b)
			continue
		}
		if r.BID != a.BID {
			s := fmt.Sprintf(BizErrors[ReceiptBizMismatch].Message, rlib.IDtoShortString("RCPT", r.RCPTID))
			b := BizError{Errno: ReceiptBizMismatch, Message: s}
			e = append(e, b)
			continue
		}
		rlist = append(rlist, r)
	}
	//------------------------------------------------------------
	// next, validate that the total of all newRcpts matches Amount
	//------------------------------------------------------------
	if tot != a.Amount {
		e = AddBizErrToList(e, DepositTotalMismatch)
		return e
	}

	dep, err := rlib.GetDepository(a.DEPID) // get the depository for this deposit
	if err != nil {
		var be []BizError
		return AddErrToBizErrlist(err, be)
	}

	//------------------------------------------------------------
	// Save the deposit
	//------------------------------------------------------------
	if a.DID == 0 {
		_, err := rlib.InsertDeposit(a)
		if err != nil {
			e = AddErrToBizErrlist(err, e)
		}
		for i := 0; i < len(newRcpts); i++ {
			var dp = rlib.DepositPart{
				DID:    a.DID,
				BID:    a.BID,
				RCPTID: newRcpts[i],
			}
			err = rlib.InsertDepositPart(&dp)
			if err != nil {
				e = AddErrToBizErrlist(err, e)
				continue
			}
			if rlist[i].DID == 0 {
				rlist[i].DID = a.DID
				err = rlib.UpdateReceipt(&rlist[i])
				if err != nil {
					e = AddErrToBizErrlist(err, e)
					continue
				}
			}
			//-----------------------------------------------------------------------
			// We need to make sure at this point that the funds for the receipt
			// are in the account associated with the depository.
			//-----------------------------------------------------------------------
			rlib.Console("DEPOSIT - 1\n")
			ja := rlib.GetJournalAllocationByASMandRCPTID(newRcpts[i])
			l := len(ja) // if l == 0 it is a normal receipt -- no associated auto-generated Assessment
			rlib.Console("DEPOSIT - 2\n")
			if l > 1 {
				rlib.Console("DEPOSIT - 3\n")
				err = fmt.Errorf("Multiple JournalAllocations for auto-generated Assessment. RCPTID = %d", newRcpts[i])
			} else if l == 1 {
				rlib.Console("DEPOSIT - 4\n")
				a, err := rlib.GetAssessment(ja[0].ASMID)
				if err != nil {
					rlib.Console("DEPOSIT - 5\n")
					e = AddErrToBizErrlist(err, e)
					continue
				}
				rlib.Console("DEPOSIT - 6\n")
				err = EnsureReceiptFundsToDepositoryAccount(&rlist[i], &a, &dep)
			}
			rlib.Console("DEPOSIT - 7\n")
		}
	} else {
		err := rlib.UpdateDeposit(a)
		if err != nil {
			e = AddErrToBizErrlist(err, e)
		}
		//---------------------------------------------------------------------------
		// If any receipts have been removed from the previous version.  To do
		// this we will compare the list of current Deposit's RCPTIDs to the
		// list of newly proposed RCPTIDs.  We will compare the two lists and
		// produce 2 new lists: addlist and removelist.  Then we will add and
		// link the addlist, and unlink the removelist.  The new Receipts are
		// already provided in newRcpts.
		//---------------------------------------------------------------------------
		curDepParts, err := rlib.GetDepositParts(a.DID)
		if err != nil {
			e = AddErrToBizErrlist(err, e)
			return e
		}

		current := map[int64]int{}
		for i := 0; i < len(curDepParts); i++ {
			current[curDepParts[i].RCPTID] = 0 // mark each receipt as initialized to 0
		}

		var addlist []int64
		for i := 0; i < len(newRcpts); i++ {
			_, ok := current[newRcpts[i]]
			if !ok {
				addlist = append(addlist, newRcpts[i])
			}
		}

		var newlist = map[int64]int{}
		for i := 0; i < len(newRcpts); i++ {
			newlist[newRcpts[i]] = 0
		}

		var removelist []int64
		for i := 0; i < len(curDepParts); i++ {
			_, ok := newlist[curDepParts[i].RCPTID]
			if !ok {
				removelist = append(removelist, curDepParts[i].RCPTID)
			}
		}

		//--------------------------------------------------------
		// Remove the deposit link in the removelist receipts...
		//--------------------------------------------------------
		for i := 0; i < len(removelist); i++ {
			r := rlib.GetReceipt(removelist[i])
			if r.RCPTID == 0 {
				err := fmt.Errorf("could not load receipt %d", removelist[i])
				e = AddErrToBizErrlist(err, e)
			}
			r.DID = 0
			err := rlib.UpdateReceipt(&r)
			if err != nil {
				e = AddErrToBizErrlist(err, e)
			}
			//---------------------------------------
			// Now remove the Deposit Part record...
			//---------------------------------------
			for j := 0; j < len(curDepParts); j++ {
				if curDepParts[j].RCPTID == removelist[i] {
					err = rlib.DeleteDepositPart(curDepParts[j].DPID)
					if err != nil {
						e = AddErrToBizErrlist(err, e)
					}
					break
				}
			}
			current[curDepParts[i].RCPTID]++ // mark that we've actually processed this entry
		}
		//--------------------------------------------------------
		// Add the deposit link in the addlist receipts...
		//--------------------------------------------------------
		for i := 0; i < len(addlist); i++ {
			r := rlib.GetReceipt(addlist[i])
			if r.RCPTID == 0 {
				err := fmt.Errorf("could not load receipt %d", addlist[i])
				e = AddErrToBizErrlist(err, e)
			}
			r.DID = a.DID
			err := rlib.UpdateReceipt(&r)
			if err != nil {
				e = AddErrToBizErrlist(err, e)
			}
			//-----------------------------------------
			// Add a deposit part for this receipt...
			//-----------------------------------------
			var dp = rlib.DepositPart{
				DID:    a.DID,
				BID:    a.BID,
				RCPTID: r.RCPTID,
			}
			err = rlib.InsertDepositPart(&dp)
			if err != nil {
				e = AddErrToBizErrlist(err, e)
			}
			current[newRcpts[i]]++ // mark that we've actually processed this entry
		}
	}
	return e
}
