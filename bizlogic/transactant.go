package bizlogic

import (
	"context"
	"fmt"
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

// DeleteTransactant deletes a transactant as long as the transactant is
// not referenced by any Rental Agreement as a Payor or User
//--------------------------------------------------------------------------
func DeleteTransactant(ctx context.Context, BID, TCID int64) error {
	var count int

	//---------------------------------------------------------
	// If the transactant is referenced as a payor or a user
	//---------------------------------------------------------
	q := fmt.Sprintf("SELECT COUNT(RAPID) from RentalAgreementPayors WHERE BID=%d and TCID=%d", BID, TCID)
	row := rlib.RRdb.Dbrr.QueryRow(q)
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		err := fmt.Errorf("Transactant %d cannot be deleted as it is referenced as a Payor by %d Rental Agreement(s)", TCID, count)
		return err
	}
	q = fmt.Sprintf("SELECT COUNT(RUID) from RentalAgreementPayors WHERE BID=%d and TCID=%d", BID, TCID)
	row = rlib.RRdb.Dbrr.QueryRow(q)
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		err := fmt.Errorf("Transactant %d cannot be deleted as it is referenced as a User of%d Rentable(s)", TCID, count)
		return err
	}

	// delete Prospect
	if err := rlib.DeleteProspect(ctx, TCID); err != nil {
		return err
	}
	// delete Payor
	if err := rlib.DeletePayor(ctx, TCID); err != nil {
		return err
	}
	// delete User
	if err := rlib.DeleteUser(ctx, TCID); err != nil {
		return err
	}

	return rlib.DeleteTransactant(ctx, TCID)

}
