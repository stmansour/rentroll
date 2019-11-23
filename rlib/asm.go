package rlib

import (
	"context"
	"fmt"
	"time"
)

// AppendComment adds the supplied string to a.Comment, separating it from
// any existing comment already in the field.
func (a *Assessment) AppendComment(s string) {
	if len(a.Comment) > 0 {
		a.Comment += " | "
	}
	a.Comment += s
}

// GetRecurrences is a shorthand for assessment variables to get a list
// of dates on which charges must be assessed for a particular interval of time (d1 - d2)
func (a *Assessment) GetRecurrences(d1, d2 *time.Time) []time.Time {
	return GetRecurrences(d1, d2, &a.Start, &a.Stop, a.RentCycle)
}

// ExpandAssessment processes an assessment. It adds instances of recurring
// assessments for the time period d1-d2 if they do not already exist. Then
// creates a journal entry for the assessment.
//
// NOTE: this routine does not check the current system date to limit the
//       expansion of recurring assessments.  The caller is responsible
//       for setting the correct date range.
//
// INPUTS
//            ctx = db context
//              a - the assessment of interest
//           xbiz - business info
//          d1,d2 - time range in the case where we need to add recurring instances
//  updateLedgers - flag indicating whether ledgers should be updated
//             lc - last close info. if assessments are added and the start
//                  date is prior to lc.Dt, then snap the start date to lc.OpenPeriodDt.
//
//                  There are edge cases where a recurring assessment
//                  definition was initially in a closed period and had to be snapped
//                  to lc. We still want to expand the instances as though they were
//                  starting from the original a.Start (before it was snapped).
//                  The lc snap still applies, but a comment is appended that
//                  indicates the original intended start date and notes that it
//                  was snapped to do that date being in a closed perion.
//
// RETURNS:
//    any error encountered
//-----------------------------------------------------------------------------
func ExpandAssessment(ctx context.Context, a *Assessment, xbiz *XBusiness, d1, d2 *time.Time, updateLedgers bool, lc1 *ClosePeriod) error {
	funcname := "ExpandAssessment"
	var j Journal
	var err error
	lc := *lc1 // local copy
	ExpAsmDtStartWasSet := !lc.ExpandAsmDtStart.Equal(TIME0)

	// Console("ENTERED %s: 1. a.ASMID = %d, Biz = %s (%d), d1 - d2 = %s - %s, RentCycle = %d\n", funcname, a.ASMID, xbiz.P.Designation, xbiz.P.BID, d1.Format(RRDATEREPORTFMT), d2.Format(RRDATEREPORTFMT), a.RentCycle)
	if lc.ExpandAsmDtStart.Year() < 1970 && a.RentCycle > RECURNONE {
		// Console("ExpandAsmDtStart = %s -- considering this to be unset.\n", lc.ExpandAsmDtStart.Format(RRDATEFMT3))
		// Console("Setting ExpandAsmDtStart to d1 (= %s) as lc value is unusable\n", d1.Format(RRDATEFMT3))
		lc.ExpandAsmDtStart = *d1
		ExpAsmDtStartWasSet = false
	}

	if a.RentCycle == RECURNONE {
		// Console("%s: calling ProcessNewAssessmentInstance(ctx, xbiz, %s, %s, a.ASMID = %d)\n", funcname, d1.Format(RRDATEFMT3), d2.Format(RRDATEFMT3), a.ASMID)
		j, err = ProcessNewAssessmentInstance(ctx, xbiz, d1, d2, a)
		if err != nil {
			LogAndPrintError(funcname, err)
			// Console("%s exiting. e0 Error = %s\n", funcname, err.Error())
			return err
		}

		if updateLedgers {
			_, err = GenerateLedgerEntriesFromJournal(ctx, xbiz, &j, d1, d2)
			if err != nil {
				LogAndPrintError(funcname, err)
				// Console("%s exiting. e1 Error = %s\n", funcname, err.Error())
				return err
			}
		}
	} else if a.RentCycle >= RECURSECONDLY && a.RentCycle <= RECURHOURLY {
		// TBD
		LogAndPrint("Unhandled assessment recurrence type: %d\n", a.RentCycle)
		// Console("%s exiting. e2 Error = %s\n", funcname, err.Error())
		return fmt.Errorf("Unhandled assessment recurrence type: %d", a.RentCycle)
	} else if a.PASMID == 0 { // it's recurring, but make sure it's a definition and not an instance - instances maintain their recur mode, but they have a PASMID > 0
		dtLimitStart := *d1     // bounds for recurrence expansion
		dtLimitStop := *d2      // bounds for recurrence expansion
		idempotentCheck := true // by default, we always do this check

		//-----------------------------------------------------------------------------
		// If we're not dealing with closed periods or back-dated rental agreements,
		// then don't start before the assessment's start date.
		//-----------------------------------------------------------------------------
		if lc.ExpandAsmDtStart.Equal(TIME0) && dtLimitStart.Before(a.Start) {
			dtLimitStart = a.Start
		}
		if dtLimitStop.After(a.Stop) {
			dtLimitStop = a.Stop
		}

		//-----------------------------------------------------------------------------
		// If we are dealing with closed periods or back-dated rental agreements,
		// then we may need to set the start date of the assessment before the
		// start date of the Recurring Assessment Definition, which may have been
		// snapped to a later date because of a closed period
		//-----------------------------------------------------------------------------
		if !lc.ExpandAsmDtStart.Equal(TIME0) && lc.ExpandAsmDtStart.Before(dtLimitStart) {
			dtLimitStart = lc.ExpandAsmDtStart
		}

		// ensure that we don't expand too far out...
		if !lc.ExpandAsmDtStop.Equal(TIME0) && lc.ExpandAsmDtStop.Before(dtLimitStop) {
			dtLimitStop = lc.ExpandAsmDtStop
		}

		//-----------------------------------------------------------------------------
		// Do the expansion list BEFORE applying any of the other limit dates.
		// The expansion gives us the date of instances needed -- regardless of
		// the closed periods. The limit dates will force those instances onto
		// acceptable dates.
		//-----------------------------------------------------------------------------
		// Console("%s: dtLimitStart = %s, dtLimitStop = %s\n", funcname, ConDt(&dtLimitStart), ConDt(&dtLimitStop))
		// Console("%s: GetRecurrences(%s, %s, %d) =\n", funcname, ConsoleDRange(&dtLimitStart, &dtLimitStop), ConsoleDRange(&dtLimitStart, &dtLimitStop), a.RentCycle)
		dl := GetRecurrences(&dtLimitStart, &dtLimitStop, &dtLimitStart, &dtLimitStop, a.RentCycle)
		// Console("%s: GetRecurrences returned %d dates\n", funcname, len(dl))
		for i := 0; i < len(dl); i++ {
			// Console("%s\n", dl[i].Format(RRDATEFMT3))
		}

		//-------------------------------------------------------------------------
		// If the expansion start date is prior to d1 (but not TIME0), then use the
		// expansion start date. It means that the caller is informing us that
		// the assessment was started in a closed period and that the asmt
		// Start date has been snapped to the nearest open date. But the expansion
		// needs to start on the date provided in lc.ExpandAsmDtStart
		//-------------------------------------------------------------------------
		if ExpAsmDtStartWasSet && lc.ExpandAsmDtStart.Before(*d1) {
			dtLimitStart = lc.ExpandAsmDtStart
		} else if a.Start.After(dtLimitStart) {
			dtLimitStart = a.Start
		}

		// Console("%s: Ensure stop limit (%s) is not past lc.ExpandAsmDtStop (%s)\n", funcname, dtLimitStop.Format(RRDATEFMT3), lc.ExpandAsmDtStop.Format(RRDATEFMT3))
		if dtLimitStop.After(lc.ExpandAsmDtStop) {
			dtLimitStop = lc.ExpandAsmDtStop
		}
		// Console("%s: Ensure stop limit (%s) is not past a.Stop (%s)\n", funcname, dtLimitStop.Format(RRDATEFMT3), a.Stop.Format(RRDATEFMT3))
		if a.Stop.Before(dtLimitStop) {
			dtLimitStop = a.Stop
		}
		// Console("%s: stop limit is %s\n", funcname, dtLimitStop.Format(RRDATEFMT3))

		//-------------------------------------------------------------------
		// Remove hours/mins differences
		//-------------------------------------------------------------------
		dtLimitStop = time.Date(dtLimitStop.Year(), dtLimitStop.Month(), dtLimitStop.Day(), 0, 0, 0, 0, time.UTC)
		dtLimitStart = time.Date(dtLimitStart.Year(), dtLimitStart.Month(), dtLimitStart.Day(), 0, 0, 0, 0, time.UTC)

		//-------------------------------------------------------------------
		// DEBUG: code to get the difference between dl and dl1
		//-------------------------------------------------------------------
		// dl1 := a.GetRecurrences(&dtLimitStart, &dtLimitStop)
		// Console("\n\n***<<<<< OLD LIMITS:  Start = %s, Stop = %s\n", a.Start.Format(RRDATEFMT4), a.Stop.Format(RRDATEFMT4))
		// Console("***>>>>> NEW LIMITS:  Start = %s, Stop = %s\n", dtLimitStart.Format(RRDATEFMT4), dtLimitStop.Format(RRDATEFMT4))
		// Console("%s: 2.4   len(dl) = %d,  len(dl1) = %d\n", funcname, len(dl), len(dl1))
		// dbug := true
		// if dbug && len(dl) != len(dl1) {
		// Console("**** YIPES  -- different expansions!\n")
		// Console("       dl[]         dl1[]\n")
		// 	ra, err := GetRentalAgreement(ctx, a.RAID)
		// 	if err != nil {
		// Console("error getting RAID = %s\n", err.Error())
		// 	}
		// 	l := len(dl)
		// 	l1 := len(dl1)
		// 	lp := l
		// 	if l1 > l {
		// 		lp = l1
		// 	}
		// 	for k := 0; k < lp; k++ {
		// 		s1 := fmt.Sprintf("%d. ", k)
		// 		if k < l {
		// 			s1 += fmt.Sprintf("%12s\t", dl[k].Format(RRDATEFMT3))
		// 		} else {
		// 			s1 += fmt.Sprintf("%12s\t", " ")
		// 		}
		// 		if k < l1 {
		// 			s1 += fmt.Sprintf("%12s", dl1[k].Format(RRDATEFMT3))
		// 		}
		// Console("%s\n", s1)
		// 	}
		// Console("%s: 2.5   RAID = %d, RentStart = %s, RentStop = %s\n\n", funcname, ra.RAID, ra.RentStart.Format(RRDATEFMT3), ra.RentStop.Format(RRDATEFMT3))
		// }

		rangeDuration := d2.Sub(dtLimitStart)
		// Console("%s: 3   -  len(dl) = %d\n", funcname, len(dl))
		for i := 0; i < len(dl); i++ {
			a1 := *a
			// Console("ExpandAssessment: 3.1  a1.Amount = %.2f\n", a1.Amount)
			a1.ASMID = 0        // ensure this is a new assessment
			a1.PASMID = a.ASMID // parent assessment
			a1.FLAGS = 0        // ensure that we don't cary forward any flags

			//-------------------------------------------------------------------
			// Now we must respect the last close date, lc.Dt.  If the date of this
			// instance is prior to lc.Dt then snap it to lc.OpenPeriodDt and add an explanation
			//-------------------------------------------------------------------
			dtStart := dl[i]
			// Console("<<<< %d >>>> setting dtstart of new instance: dtStart = %s,  lc.Dt = %s\n", i, dtStart.Format(RRDATEFMT3), lc.Dt.Format(RRDATEFMT3))
			if dtStart.Before(lc.Dt) {
				// Console("<<<< %d >>>> dtStart is before lc.Dt (%s), ==> so setting to lc.OpenPeriodDt: %s\n", i, lc.Dt.Format(RRDATEFMT3), lc.OpenPeriodDt.Format(RRDATEFMT3))
				a1.AppendComment(fmt.Sprintf("Snapping %s to open period: %s", dtStart.Format(RRDATEFMT3), lc.OpenPeriodDt.Format(RRDATEFMT3)))
				idempotentCheck = false
				// Console("     idempotentCheck = %t\n", idempotentCheck)
				dtStart = lc.OpenPeriodDt
			}
			a1.Start = dtStart
			a1.Stop = dtStart
			// Console("%s: a1.Start, a1.Stop = %s,  a1.Comment = %s\n", funcname, a1.Start.Format(RRDATEFMT3), a1.Comment)

			// // // a1.Stop = dl[i].Add(CycleDuration(a.ProrationCycle, a.Start)) // add enough time so that the recurrence calculator sees this instance

			// Console("****>>>>>>  a1.Start = %s\n", a1.Start.Format(RRDATEFMT4))
			// Console("****>>>>>>  a1.Stop  = %s\n", a1.Stop.Format(RRDATEFMT4))
			// Console("****>>>>>>  CycleDuration( %d, %s ) --->  %d\n", a.ProrationCycle, a.Start.Format(RRDATEFMT4), CycleDuration(a.ProrationCycle, a.Start))

			//--------------------------------------------------------------------------------
			// Before inserting this, validate that the RentalAgreement for this assessment
			// is still in effect.  This is because in the early versions of the Roller
			// server code, there were no checks to ensure that recurring assessments stopped
			// when their associated RentalAgreements stopped.
			//--------------------------------------------------------------------------------
			if a.RAID > 0 {
				// Console("%s: 3.2  a.RAID = %d\n", funcname, a.RAID)
				ra, err := GetRentalAgreement(ctx, a.RAID)
				if err != nil {
					LogAndPrintError(funcname, err)
					// Console("%s exiting. e3 Error = %s\n", funcname, err.Error())
					return err
				}
				// Console("%s: 3.3  ra.RentStop = %s\n", funcname, ra.RentStop)
				//----------------------------------------------------------------------
				// Check the boundaries: don't create instances outside the boundaries
				// of the recurring definition.
				// NOTE: THE ONLY EXCEPTION is if we're expanding something that
				// would otherwise expand into closed periods. We will add these
				// instances at the requested dates, but we will make notes in
				// the comments explaining what's happening
				//----------------------------------------------------------------------
				// Console("%s: 3.31 ExpAsmDtStartWasSet = %t, lc: Dt = %s, ExpandAsmStart/Stop = %s\n", funcname, ExpAsmDtStartWasSet, lc.Dt.Format(RRDATEFMT3), ConsoleDRange(&lc.ExpandAsmDtStart, &lc.ExpandAsmDtStop))
				if (a1.Start.After(ra.RentStop) || a1.Start.Equal(ra.RentStop)) && !ExpAsmDtStartWasSet {
					// Console("%s: 3.4  Do not add the new assessment\n", funcname)
					err = fmt.Errorf("%s:  Cannot add new assessment instance on %s after RentalAgreement (%s) stop date %s", funcname, a1.Start.Format(RRDATEREPORTFMT), ra.IDtoShortString(), ra.RentStop.Format(RRDATEREPORTFMT))
					LogAndPrintError(funcname, err)
					// Console("%s: 3.5  exiting. e4 Error = %s\n", funcname, err.Error())
					return err
				}
			}

			//--------------------------------------------------------------------------------
			// The generation of recurring assessment instances needs to be idempotent.
			// Check to ensure that this instance does not already exist before generating it
			//--------------------------------------------------------------------------------
			var a2 Assessment
			if idempotentCheck {
				// Console("%s: 3.6  a1.Start = %s, a1.PASMID = %d\n", funcname, a1.Start.Format(RRDATEFMT3), a1.PASMID)
				a2, err = GetAssessmentInstance(ctx, &a1.Start, a1.PASMID) // if this returns an existing instance (ASMID != 0) then it's already been processed...
				if err != nil {
					// Console("%s: 3.61  Error in GetAssessmentInstance: %s\n", funcname, err.Error())
				}
				// Console("%s: 3.7  a2.ASMID = %d\n", funcname, a2.ASMID)
			}
			if a2.ASMID == 0 { // ... otherwise, process this instance
				//--------------------------------------------------------------------------------
				// Rent is assessed on the following cycle: a.RentCycle
				// and prorated on the following cycle: a.ProrationCycle
				//--------------------------------------------------------------------------------
				rentCycleDur := CycleDuration(a.RentCycle, dl[i])
				diff := rangeDuration - rentCycleDur
				if diff < 0 {
					diff = -diff
				}
				dtb := *d1 // beginning
				dte := *d2 // end
				// Console("%s: 4.2  dtb = %s, dte = %s, diff = %d\n", funcname, dtb.Format(RRDATEFMT4), dte.Format(RRDATEFMT4), diff)
				// Console("%s: 4.21  a.ASMID = %d, a.PASMID = %d, a.Start = %s, a.Stop = %s\n", funcname, a.ASMID, a.PASMID, a.Start.Format(RRDATEFMT3), a.Stop.Format(RRDATEFMT3))

				//-------------------------------------------------------------------
				// TODO: see if we an remove this...
				// This bit of code is really old and I don't like it.  Looks like
				// there was some concern about the the date range being correct.
				//-------------------------------------------------------------------
				if diff > rentCycleDur/9 { // if this is true then...
					dtb = dl[i]                                    // use the instance start point and...
					dte = dtb.Add(CycleDuration(a.RentCycle, dtb)) // add one full cycle duration
				}

				// Console("%s: 4.22 - dtb = %s, dte = %s\n", funcname, dtb.Format(RRDATEFMT3), dte.Format(RRDATEFMT3))
				//-------------------------------------------------------------------
				// If anything has pushed dte beyond the stop date of the parent,
				// now is the time to snap it back.  This also causes the instance
				// to be prorated...
				//-------------------------------------------------------------------
				if dte.After(a.Stop) {
					dte = a.Stop
					// Console("%s: 4.23 - calling SimpleProrateAmount( amount %8.2f ,rentcycle %d, prorate %d, %s, %s, %s )\n", funcname, a.Amount, a.RentCycle, a.ProrationCycle, dtb.Format(RRDATEFMT4), dte.Format(RRDATEFMT4), a.Start.Format(RRDATEFMT4))
					amt, num, den := SimpleProrateAmount(a.Amount, a.RentCycle, a.ProrationCycle, &dtb, &dte, &a.Start)
					a1.Amount = amt
					a1.AppendComment(ProrateComment(num, den, a.ProrationCycle))
				}

				// Console("%s: 4.3   dtb = %s, dte = %s\n", funcname, dtb.Format(RRDATEFMT4), dte.Format(RRDATEFMT4))
				// Console("%s: 4.31  a1.ASMID = %d, a1.PASMID = %d, a1.Start = %s, a1.Stop = %s\n", funcname, a1.ASMID, a1.PASMID, a1.Start.Format(RRDATEFMT4), a1.Stop.Format(RRDATEFMT4))

				// Prorate this assessment if necessary...

				//-----------------------------------------------------------
				// Insert the assessment
				//-----------------------------------------------------------
				// Console("%s: 4.0, a1.Amount = %.2f\n", funcname, a1.Amount)
				_, err = InsertAssessment(ctx, &a1)
				Errlog(err)
				// Console("%s: 4.1, inserted a1.ASMID = %d, a1.Amount = %.2f\n", funcname, a1.ASMID, a1.Amount)

				//-----------------------------------------------------------
				// Update the Journal
				//-----------------------------------------------------------
				var j Journal
				j, err = ProcessNewAssessmentInstance(ctx, xbiz, &dtb, &dte, &a1)
				if err != nil {
					LogAndPrintError(funcname, err)
					return err
				}
				if updateLedgers {
					_, err = GenerateLedgerEntriesFromJournal(ctx, xbiz, &j, d1, d2)
					if err != nil {
						LogAndPrintError(funcname, err)
						return err
					}
				}
			} else if a.RentCycle >= RECURSECONDLY && a.RentCycle <= RECURHOURLY {
				LogAndPrintError(funcname, fmt.Errorf("Unhandled RentCycle frequency: %d", a.RentCycle))
			}
			// Console("%s: 5\n", funcname)
		}
	}
	// Console("%s: 6  exiting. end of func\n", funcname)
	return err
}
