package bizlogic

import (
	"fmt"
	"rentroll/rlib"
)

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
			}
			if rlist[i].DID == 0 {
				rlist[i].DID = a.DID
				err = rlib.UpdateReceipt(&rlist[i])
				if err != nil {
					e = AddErrToBizErrlist(err, e)
				}
			}
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
			current[curDepParts[i].RCPTID]++ // mark that we've acutally processed this entry
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
			current[newRcpts[i]]++ // mark that we've acutally processed this entry
		}
	}
	return e
}
