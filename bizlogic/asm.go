package bizlogic

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// UpdateAssessment updates the supplied assessment, reversing existing assessments
// if necessary
//
// INPUTS
//    a = the assessment to insert
//  exp = if it is a recurring assessment and the start date is in the past, should
//        past entries be created?  true = yes
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func UpdateAssessment(a *rlib.Assessment) []BizError {

	// Assessments need to be backed out if any of the following change:
	//   ARID
	//   RAID
	//   RID
	//   Amount
	//   Recur Cycle
	//   Proration Cycle
	//   Start Date
	//   Stop Date if it moves backwards in time

	return nil
}

// ReverseAssessment reverses an existing assessment
//
// INPUTS
//    aold = the assessment to reverse
//    mode = how to handle recurring assessments:
//           0: just reverse this instance
//           1: reverse this and future instances
//           2: reverse all instances
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessment(aold *rlib.Assessment, mode int) []BizError {
	var errlist []BizError
	switch mode {
	case 0:
		errlist = ReverseAssessmentInstance(aold)
	case 1:
		errlist = ReverseAssessmentsGoingForward(aold, &aold.Start)
	case 2:
		var a rlib.Assessment
		var err error
		if aold.PASMID > 0 {
			a, err = rlib.GetAssessment(aold.PASMID)
			if err != nil {
				return bizErrSys(err)
			}
		}
		if a.RentCycle == rlib.RECURNONE {
			return ReverseAssessmentInstance(a)
		}

		errlist = ReverseAssessmentsGoingForward(a, &a.Start) // reverse from start of recurring instances forward
		if len(errlist) > 0 {
			return errlist
		}
		a.FLAGS |= 0x4 // mark this assessment as void
		err = rlib.UpdateAssessment(&a)
		if err != nil {
			return bizErrSys(&err)
		}

	default:
		err := fmt.Errorf("ReverseAssessment:  unsupported mode: %d", mode)
		return bizErrSys(&err)
	}
	return errlist
}

// ReverseAssessmentsGoingForward reverses an existing assessment
//
// INPUTS
//    aold = the first in a series of assessments to reverse
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessmentsGoingForward(aold *rlib.Assessment, dt *time.Time) []BizError {
	var errlist []BizError

	d2 := time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)
	m := rlib.GetAssessmentInstancesByParent(aold.PASMID, dt, &d2)
	for i := 0; i < len(m); i++ {
		errlist = ReverseAssessmentInstance(&m[i])
		if len(errlist) > 0 {
			return errlist
		}
	}
	return nil
}

// ReverseAssessmentInstance reverses a single instance of an assessment
//
// INPUTS
//    aold = the assessment to reverse
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessmentInstance(aold *rlib.Assessment) []BizError {
	anew := *aold
	anew.ASMID = 0
	anew.Amount = -anew.Amount
	anew.RPASMID = aold.ASMID
	anew.FLAGS |= 1 << 2 // set bit 2 to mark that this assessment is void
	anew.Comment = fmt.Sprintf("Reversal of %s", aold.IDtoString())

	errlist := InsertAssessment(&anew, true)
	if len(errlist) > 0 {
		return errlist
	}

	aold.Comment = fmt.Sprintf("Reversed by %s", anew.IDtoString())
	aold.FLAGS |= 1 << 2 // set bit 2 to mark that this assessment is void
	err := rlib.UpdateAssessment(aold)
	if err != nil {
		return bizErrSys(&err)
	}
	return nil
}

// InsertAssessment performs bizlogic checks first, then inserts the Assessment,
// then adds the associated Journal and Ledger entries
//
// INPUTS
//    a = the assessment to insert
//  exp = if it is a recurring assessment and the start date is in the past, should
//        past entries be created?  true = yes
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func InsertAssessment(a *rlib.Assessment, exp bool) []BizError {
	var errlist []BizError
	errlist = ValidateAssessment(a) // Make sure there are no bizlogic errors before saving
	if len(errlist) > 0 {
		return errlist
	}
	_, err := rlib.InsertAssessment(a) // No bizlogic errors, save it
	if err != nil {
		return bizErrSys(&err)
	}

	//------------------------------------------------
	// Add the journal and ledger entries
	//------------------------------------------------
	var xbiz rlib.XBusiness
	rlib.GetXBusiness(a.BID, &xbiz)
	d1, d2 := rlib.GetMonthPeriodForDate(&a.Start) // TODO: probably needs to be more generalized
	rlib.InitLedgerCache()
	if a.RentCycle == rlib.RECURNONE { // for nonrecurring, use existng struct: a
		rlib.ProcessJournalEntry(a, &xbiz, &d1, &d2, true)
	} else if exp && a.PASMID == 0 { // only expand if we're asked and if we're not an instance
		now := rlib.DateAtTimeZero(time.Now())
		dt := rlib.DateAtTimeZero(a.Start)
		if !dt.After(now) {
			createInstancesToDate(a, &xbiz)
		}
	}
	return nil
}

// ValidateAssessment checks to see whether the assessment violates any
// business logic.
//
// INPUTS
//    a = the assessment to validate
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ValidateAssessment(a *rlib.Assessment) []BizError {
	var e []BizError
	if a.RID > 0 {
		//--------------------------------------------------------------------------
		//  Check for assessment timeframe prior to or after Rentable's type being defined
		//--------------------------------------------------------------------------
		rtl := rlib.GetRentableTypeRefs(a.RID) // these are returned in chronological order
		l := len(rtl)
		if l == 0 {
			e = append(e, BizErrors[RentableTypeUnknown])
		} else {
			if a.Stop.Before(rtl[0].DtStart) || a.Start.After(rtl[l-1].DtStop) {
				e = append(e, BizErrors[RentableTypeUnknown])
			}
		}

		//--------------------------------------------------------------------------
		//  Check for assessment timeframe prior to or after Rentable's status being defined
		//--------------------------------------------------------------------------
		rsl := rlib.GetRentableStatusByRange(a.RID, &a.Start, &a.Stop)
		l = len(rsl)
		if l == 0 {
			fmt.Printf("ValidateAssessment: l=0\n")
			e = append(e, BizErrors[RentableStatusUnknown])
		} else {
			fmt.Printf("ValidateAssessment: a.Start-Stop = %s - %s\n", a.Start.Format(rlib.RRDATEINPFMT), a.Stop.Format(rlib.RRDATEINPFMT))
			fmt.Printf("ValidateAssessment: rtl = %s - %s\n", rtl[0].DtStart.Format(rlib.RRDATEINPFMT), rtl[l-1].DtStop.Format(rlib.RRDATEINPFMT))
			if a.Stop.Before(rtl[0].DtStart) || a.Start.After(rtl[l-1].DtStop) {
				e = append(e, BizErrors[RentableStatusUnknown])
			}
		}

	}
	return e
}

// createInstancesToDate creates all instances of a recurring Assessments up to the
// supplied date
//
// INPUTS
//    a = the recurring assessment
// xbiz = Business information
//
// RETURNS
//
//-------------------------------------------------------------------------------------
func createInstancesToDate(a *rlib.Assessment, xbiz *rlib.XBusiness) {
	now := time.Now()
	as := time.Date(a.Start.Year(), a.Start.Month(), a.Start.Day(), 0, 0, 0, 0, time.UTC)
	m := rlib.GetRecurrences(&a.Start, &a.Stop, &as, &now, a.RentCycle) // get all from the begining up to now
	for i := 0; i < len(m); i++ {
		dt1, dt2 := rlib.GetMonthPeriodForDate(&m[i])
		rlib.ProcessJournalEntry(a, xbiz, &dt1, &dt2, true) // this generates the assessment instances
	}
}
