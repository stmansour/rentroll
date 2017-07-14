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
//    a = the assessment to insert
//
// RETURNS
//    a slice of BizErrors
//-------------------------------------------------------------------------------------
func ReverseAssessment(aold *rlib.Assessment) []BizError {
	anew := *aold
	anew.Amount = -anew.Amount
	anew.RPASMID = aold.ASMID
	anew.Comment = fmt.Sprintf("Reversal of %s", aold.IDtoString())

	errlist := InsertAssessment(&anew, true)
	if len(errlist) > 0 {
		return errlist
	}

	aold.Comment = fmt.Sprintf("Reversed by %s", anew.IDtoString())
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
	} else if exp {
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

// bizErrSys just encapsulates returning an error in a []BizError.  The Errno
// is set to 0.
//
// INPUTS
//  err = pointer to an error
//
// RETURNS
//  a slize of BizError containing the error message
//-------------------------------------------------------------------------------------
func bizErrSys(err *error) []BizError {
	var errlist []BizError
	berr := BizError{
		Errno:   0, // system error
		Message: "Error inserting assessment = " + (*err).Error(),
	}
	errlist = append(errlist, berr)
	return errlist
}
