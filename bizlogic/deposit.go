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
func SaveDeposit(a *rlib.Deposit, rcpts []int64) []BizError {
	var e []BizError
	var rlist []rlib.Receipt
	//------------------------------------------------------------
	// First, validate that all rcpts are eligible for inclusion
	// in this receipt
	//------------------------------------------------------------
	for i := 0; i < len(rcpts); i++ {
		r := rlib.GetReceipt(rcpts[i])
		if r.DID != 0 && r.DID != a.DID {
			s := fmt.Sprintf(BizErrors[ReceiptAlreadyDeposited].Message, rlib.IDtoShortString("RCPT", r.RCPTID), rlib.IDtoShortString("D", r.DID))
			b := BizError{Errno: ReceiptAlreadyDeposited, Message: s}
			e = append(e, b)
			continue
		}
		rlist = append(rlist, r)
	}
	if a.DID == 0 {
		_, err := rlib.InsertDeposit(a)
		if err != nil {
			e = AddErrToBizErrlist(err, e)
		}
		for i := 0; i < len(rcpts); i++ {
			var dp = rlib.DepositPart{
				DID:    a.DID,
				BID:    a.BID,
				RCPTID: rcpts[i],
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
	}
	return e
}
