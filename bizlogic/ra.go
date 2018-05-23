package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"time"
)

// UpdateRentalAgreement performs business checks to ensure that a
// RentalAgreement update is done along with cleanup as necessary.
//
// INPUTS
//  ctx   for db transactions
//  ra    the rental agreement being updated
//-----------------------------------------------------------------------------
func UpdateRentalAgreement(ctx context.Context, ra *rlib.RentalAgreement) []BizError {
	raOrig, err := rlib.GetRentalAgreement(ctx, ra.RAID)
	if err != nil {
		return bizErrSys(&err)
	}

	//---------------------------------------------------------------------------
	// 1. If the stop date is being changed then adjust all recurring
	// Assessments so that nothing extends past the rental agreement stop date.
	//---------------------------------------------------------------------------
	rlib.Console("bizlogic: Update Rental Agreement. raOrig.RentStop = %s, ra.RentStop = %s\n", raOrig.RentStop.Format(rlib.RRDATEREPORTFMT), ra.RentStop.Format(rlib.RRDATEREPORTFMT))

	if !raOrig.RentStop.Equal(ra.RentStop) {
		// rlib.Console("Dates do not match, calling adjust Assessments\n")
		if be := adjustAssessments(ctx, ra); be != nil {
			return be
		}
	}

	//------------------------------------------------------------------------
	// 2. If any of the start dates are prior to this RA's initial LedgerMarker
	// then we need to move the initial LedgerMarker's date back.
	//------------------------------------------------------------------------
	rlib.Console("bizlogic: 2.  Ledger Markers")
	lm, err := rlib.GetInitialLedgerMarkerByRAID(ctx, ra.RAID)
	if lm.LMID == 0 || err != nil {
		e := fmt.Errorf("Could not find initial LedgerMarker for RAID = %d", ra.RAID)
		return bizErrSys(&e)
	}
	// rlib.Console("Found initial LedgerMarker for RAID %d\n", a.RAID)
	if lm.Dt.After(ra.AgreementStart) || lm.Dt.After(ra.PossessionStart) || lm.Dt.After(ra.RentStart) {
		// find the earliest date...
		dt, _ := rlib.GetRentalAgreementEarliestDate(ctx, ra)
		// rlib.Console("Moving initial LedgerMarker to: %s\n", dt.Format(rlib.RRDATEREPORTFMT))
		if dt.Before(lm.Dt) {
			lm.Dt = dt // update the ledger marker date to the earliest date
			err = rlib.UpdateLedgerMarker(ctx, &lm)
			if err != nil {
				e := fmt.Errorf("Error saving Rental Agreement RAID = %d: %s", ra.RAID, err.Error())
				return bizErrSys(&e)
			}
		}
	}

	//-------------------------------------------------------------------------------
	// OK, if we made it this far, there are no biz rules broken. Go ahead and save
	// the Rental Agreement...
	//-------------------------------------------------------------------------------
	err = rlib.UpdateRentalAgreement(ctx, ra)
	return nil
}

// adjustAssessments handles any assessments out of the RentalAgreement time range
//-----------------------------------------------------------------------------
func adjustAssessments(ctx context.Context, ra *rlib.RentalAgreement) []BizError {
	// rlib.Console("Entered adjustAssessments\n")
	now := time.Now()
	m, err := rlib.GetEpochAssessmentsByRentalAgreement(ctx, ra.RAID)
	if err != nil {
		return bizErrSys(&err)
	}
	for i := 0; i < len(m); i++ {
		// rlib.Console("Found ASMID = %d\n", m[i].ASMID)
		//-----------------------------------------------------------------------
		// If non recurring then reverse if it is past new RentStop
		// If recurring then snap stop dates past new RentStop to RentStop
		//-----------------------------------------------------------------------
		if m[i].RentCycle == rlib.RECURNONE {
			if m[i].Start.After(ra.RentStop) {
				// rlib.Console("Reversing ASMID = %d\n", m[i].ASMID)
				if be := ReverseAssessment(ctx, &m[i], 0, &now); be != nil {
					return be
				}
			}
		} else if m[i].Stop.After(ra.RentStop) {
			m[i].Stop = ra.RentStop
			// rlib.Console("Set stop date for ASMID = %d to %s\n", m[i].ASMID, ra.RentStop.Format(rlib.RRDATEREPORTFMT))
			if err = rlib.UpdateAssessment(ctx, &m[i]); err != nil {
				return bizErrSys(&err)
			}
		}
	}
	return nil
}
