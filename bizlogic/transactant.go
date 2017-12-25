package bizlogic

import (
	"context"
	"rentroll/rlib"
	"time"
)

// FinalizeTransactant performs bizlogic checks on the transactant
// and initializes the LedgerMarker for any payments they make.
func FinalizeTransactant(ctx context.Context, a *rlib.Transactant) []BizError {
	var errlist []BizError
	var lm = rlib.LedgerMarker{
		LID:     0,
		BID:     a.BID,
		RAID:    0,
		RID:     0,
		TCID:    a.TCID,
		Dt:      time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
		Balance: 0.0,
		State:   rlib.LMINITIAL,
	}
	_, err := rlib.InsertLedgerMarker(ctx, &lm)
	if err != nil {
		errlist = AddErrToBizErrlist(err, errlist)
	}
	return errlist
}
