package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
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
func EnsureReceiptFundsToDepositoryAccount(ctx context.Context, r *rlib.Receipt, asmid int64, d *rlib.Depository, deposit *rlib.Deposit) error {
	const funcname = "EnsureReceiptFundsToDepositoryAccount"
	var (
		xbiz rlib.XBusiness
		err  error
	)

	err = rlib.InitBizInternals(r.BID, &xbiz)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return err
	}

	ar := rlib.RRdb.BizTypes[r.BID].AR[r.ARID] // this is the Account Rule for the Receipt

	//-------------------------------------------------------------------------
	// If we have not already debited the GLAccount associated with the
	// Depository, this is the place to do so.  CREDIT what the receipt
	// account rule debited, then DEBIT the Depository GLAccount
	//-------------------------------------------------------------------------
	if ar.DebitLID != d.LID {
		sAcctRule := fmt.Sprintf("d %s _, c %s _",
			rlib.RRdb.BizTypes[r.BID].GLAccounts[d.LID].GLNumber,
			rlib.RRdb.BizTypes[r.BID].GLAccounts[ar.DebitLID].GLNumber) // string acct rule for what we're about to do: debit bank, credit UndepFunds
		//----------------------------------
		// debit  d.LID r.Amount
		// credit ar.DebitLID r.Amount
		//----------------------------------
		var jnl = rlib.Journal{
			BID:     r.BID,
			Amount:  r.Amount,
			Dt:      deposit.Dt,
			Type:    rlib.JNLTYPEXFER,
			ID:      r.RCPTID,
			Comment: fmt.Sprintf("auto-transfer for deposit %s", deposit.IDtoShortString()),
		}
		_, err = rlib.InsertJournal(ctx, &jnl)
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
			ASMID:  asmid,
			TCID:   r.TCID,
			RCPTID: r.RCPTID,
		}
		_, err = rlib.InsertJournalAllocationEntry(ctx, &ja)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}
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
		_, err = rlib.InsertLedgerEntry(ctx, &l)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		l.LID = ar.DebitLID
		l.Amount = -r.Amount
		_, err = rlib.InsertLedgerEntry(ctx, &l)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return err
		}

		//-------------------------------------------------------------------------
		// ReceiptAllocation reflection of what just happened...
		//-------------------------------------------------------------------------
		var ra rlib.ReceiptAllocation
		ra.RCPTID = r.RCPTID
		ra.Amount = r.Amount
		ra.AcctRule = sAcctRule
		ra.BID = r.BID
		ra.Dt = r.Dt
		ra.RAID = r.RAID
		_, err = rlib.InsertReceiptAllocation(ctx, &ra)
		if err != nil {
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
func SaveDeposit(ctx context.Context, a *rlib.Deposit, newRcpts []int64) []BizError {
	// rlib.Console("SaveDeposit: 0\n")
	var e []BizError
	var rlist []rlib.Receipt
	tot := float64(0)
	//----------------------------------------------------------------
	// First, validate that all newRcpts are eligible for inclusion
	// in this receipt
	//----------------------------------------------------------------
	for i := 0; i < len(newRcpts); i++ {
		r, err := rlib.GetReceipt(ctx, newRcpts[i])
		if err != nil {
			elist := bizErrSys(&err)
			e = append(e, elist[0])
			continue
		}
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

	dep, err := rlib.GetDepository(ctx, a.DEPID) // get the depository for this deposit
	if err != nil {
		var be []BizError
		return AddErrToBizErrlist(err, be)
	}

	//------------------------------------------------------------
	// Save the deposit
	//------------------------------------------------------------
	if a.DID == 0 {
		_, err := rlib.InsertDeposit(ctx, a)
		if err != nil {
			e = AddErrToBizErrlist(err, e)
		}
		for i := 0; i < len(newRcpts); i++ {
			var dp = rlib.DepositPart{
				DID:    a.DID,
				BID:    a.BID,
				RCPTID: newRcpts[i],
			}
			_, err = rlib.InsertDepositPart(ctx, &dp)
			if err != nil {
				e = AddErrToBizErrlist(err, e)
				continue
			}
			if rlist[i].DID == 0 {
				rlist[i].DID = a.DID
				err = rlib.UpdateReceipt(ctx, &rlist[i])
				if err != nil {
					e = AddErrToBizErrlist(err, e)
					continue
				}
			}
			//-----------------------------------------------------------------------
			// We need to make sure at this point that the funds for the receipt
			// are in the account associated with the depository.  Basically, this
			// entails looking at the AccountRule used on the Receipt. If the Debit
			// account is NOT the same as the Depository's account, then
			// automatically generate a transfer
			//-----------------------------------------------------------------------
			var xbiz rlib.XBusiness
			err = rlib.InitBizInternals(a.BID, &xbiz)
			if err != nil {
				return bizErrSys(&err)
			}
			debitLID := rlib.RRdb.BizTypes[a.BID].AR[rlist[i].ARID].DebitLID // find the receipt's Account Rule Debit LID

			// rlib.Console("debitLID = %d - AcctRule: %s\n", debitLID, rlib.RRdb.BizTypes[a.BID].AR[rlist[i].ARID].Name)
			// rlib.Console("Depostitory LID = %d\n", dep.LID)

			// compare to LID for the Depository
			if debitLID != dep.LID { // if they're not the same, transfer to the appropriate account
				asmid := int64(0) // assume no auto-gen Assessment is associated with this receipt
				ja, err := rlib.GetJournalAllocationByASMandRCPTID(ctx, newRcpts[i])
				if err != nil {
					return AddErrToBizErrlist(err, e)
				}

				if len(ja) == 1 { // if there is an associated auto-gen'd assessment
					asmt, err := rlib.GetAssessment(ctx, ja[0].ASMID)
					if err != nil {
						return AddErrToBizErrlist(err, e)
					}
					asmid = asmt.ASMID // use the correct ASMID
				}
				err = EnsureReceiptFundsToDepositoryAccount(ctx, &rlist[i], asmid, &dep, a)
				if err != nil {
					return AddErrToBizErrlist(err, e)
				}
			}
		}
	} else {
		err := rlib.UpdateDeposit(ctx, a)
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
		curDepParts, err := rlib.GetDepositParts(ctx, a.DID)
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
			r, err := rlib.GetReceipt(ctx, removelist[i])
			if err != nil {
				e = AddErrToBizErrlist(err, e)
			}
			if r.RCPTID == 0 {
				err := fmt.Errorf("could not load receipt %d", removelist[i])
				e = AddErrToBizErrlist(err, e)
			}
			r.DID = 0
			err = rlib.UpdateReceipt(ctx, &r)
			if err != nil {
				e = AddErrToBizErrlist(err, e)
			}
			//---------------------------------------
			// Now remove the Deposit Part record...
			//---------------------------------------
			for j := 0; j < len(curDepParts); j++ {
				if curDepParts[j].RCPTID == removelist[i] {
					err = rlib.DeleteDepositPart(ctx, curDepParts[j].DPID)
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
			r, err := rlib.GetReceipt(ctx, addlist[i])
			if err != nil {
				e = AddErrToBizErrlist(err, e)
			}
			if r.RCPTID == 0 { // if resource not found then also raise the error
				err := fmt.Errorf("could not load receipt %d", addlist[i])
				e = AddErrToBizErrlist(err, e)
			}
			r.DID = a.DID
			err = rlib.UpdateReceipt(ctx, &r)
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
			_, err = rlib.InsertDepositPart(ctx, &dp)
			if err != nil {
				e = AddErrToBizErrlist(err, e)
			}
			current[newRcpts[i]]++ // mark that we've actually processed this entry
		}
	}
	return e
}
