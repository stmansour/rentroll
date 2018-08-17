package rlib

import "time"

// GetRoundingDate returns the date with rounding
func GetRoundingDate(t time.Time) (n time.Time) {
	n = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return
}

// GetTodayUTCRoundingDate returns Today's date with rounding in UTC timezone
func GetTodayUTCRoundingDate() (t time.Time) {
	now := time.Now().UTC()
	t = GetRoundingDate(now)
	return
}

// GetEpochFromBaseDate returns the epoch date based on cycle,
// start date and pre-configured base epochs
//
// The required unit(s) should be extracted from pre-configured base epochs
// to calculate proper Epoch date based on cycle, start date
// For ex.,
//     1. If cycle is MINUTELY then the unit(s) to consider from pre-configured
//        base epoch(minutely) are Second, NenoSecond
//     2. If cycle if MONTHLY then unit(s) to consider from pre-configured
//        base epoch(monthly) are Day, Hour, Minute, Second, NenoSecond
//
// Time Location: It always keeps time location from base date(b)
//
// INPUTS
//               b  = preconfigured base date
//              d1  = start date
//              d2  = stop date
//           cycle  = integer presentable number
//
// RETURNS
//     ok    - If epoch is possible in given date range then true else false
//             * In case of false, epoch still has calculated value, to be
//               happened on next cycle. It helps to determine for calle routine
//               what should it does on epoch whether it falls in range or not.
//     epoch - proper epoch for given date range
//-----------------------------------------------------------------------------
func GetEpochFromBaseDate(b, d1, d2 time.Time, cycle int64) (ok bool, epoch time.Time) {

	// TODO(Sudip): What if base date and d1 falls at same unit value

	// DECIDE BASED ON CYCLE
	switch cycle {
	case RECURNONE:
		epoch = d1

	case RECURSECONDLY:
		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), d1.Minute(), d1.Second(), b.Nanosecond(), b.Location())

		// IF EPOCH IS PASSED (PRIOR TO START), THEN MAKE IT HAPPEN ON NEXT SECOND
		if epoch.Before(d1) {
			epoch = epoch.Add(time.Second)
		}

	case RECURMINUTELY:
		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), d1.Minute(), b.Second(), b.Nanosecond(), b.Location())

		// IF EPOCH IS PASSED (PRIOR TO START), THEN MAKE IT HAPPEN ON NEXT MINUTE
		if epoch.Before(d1) {
			epoch = epoch.Add(time.Minute)
		}

	case RECURHOURLY:
		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())

		// IF EPOCH IS PASSED (PRIOR TO START), THEN MAKE IT HAPPEN ON NEXT HOUR
		if epoch.Before(d1) {
			epoch = epoch.Add(time.Hour)
		}

	case RECURDAILY:
		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())

		// IF EPOCH IS PASSED (PRIOR TO START), THEN MAKE IT HAPPEN ON NEXT DAY
		if epoch.Before(d1) {
			epoch = epoch.AddDate(0, 0, 1)
		}

	case RECURWEEKLY:
		epoch = GetWeeklyEpochDate(b, d1)

	case RECURMONTHLY:
		epoch = GetMonthlyEpochDate(b, d1)

	case RECURQUARTERLY: // TODO(Sudip): FIX QUARTER CYCLE
		epoch = time.Date(d1.Year(), d1.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())

	case RECURYEARLY:
		epoch = time.Date(d1.Year(), b.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())

		// IF EPOCH IS PASSED (PRIOR TO START), THEN MAKE IT HAPPEN ON NEXT YEAR
		if epoch.Before(d1) {
			epoch = epoch.AddDate(1, 0, 0)
		}
	}

	// IF EPOCH FALLS BEFORE THE *END DATE* (i.e, IN GIVEN DATE RANGE)
	// THEN MARK "OK" FLAG AS TRUE
	if epoch.Before(d2) {
		ok = true
	} else {
		Console("\n\n\nGetEpochFromBaseDate: base = %s, d1 = %s, d2 = %s\n", b.Format(RRDATETIMERPTFMT), d1.Format(RRDATETIMERPTFMT), d2.Format(RRDATETIMERPTFMT))
		Console("                      epoch = %s, ok = %t\n\n\n\n", epoch.Format(RRDATETIMERPTFMT), ok)
	}

	return
}

// GetWeeklyEpochDate adjusts epoch date for weekly cycle
func GetWeeklyEpochDate(base, d1 time.Time) (epoch time.Time) {
	d1wd := int(d1.Weekday())  // START DATE WEEKDAY
	bwd := int(base.Weekday()) // BASE EPOCH DATE WEEKDAY

	// CALCULATE EPOCH
	epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), base.Hour(), base.Minute(), base.Second(), base.Nanosecond(), base.Location())

	// IF IT IS SAME DAY
	if bwd == d1wd {
		// IF START IS BEFORE EPOCH THEN NOTHING TO DO
		// BUT, IF START IS AFTER EPOCH ON THE SAME DAY i.e, EPOHC IS PASSED THEN
		if d1.After(epoch) {
			epoch = epoch.AddDate(0, 0, 7-d1wd+bwd)
		}
	} else if d1wd < bwd { // IF START WEEKDAY IS BEFORE
		epoch = epoch.AddDate(0, 0, bwd-d1wd)
	} else { // IF START WEEKDAY IS FALLS AFTER
		epoch = epoch.AddDate(0, 0, 7-d1wd+bwd)
	}
	return
}

// GetMonthlyEpochDate adjusts epoch date for monthly cycle
// It will take care especially month ending dates where epoch is set on days
// greater than 28.
func GetMonthlyEpochDate(base, d1 time.Time) (epoch time.Time) {

	// ending days for each month starts from 1st index
	var endingDays = [...]int{0, // 0 means nothing, don't be fool to access this :)
		31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, // starts from Jan
	}

	// DAY, MONTH, YEAR
	day := base.Day()
	month := d1.Month() % 12
	year := d1.Year() + int(d1.Month())/12

	// CALCULATE EPOCH
	epoch = time.Date(year, month, day, base.Hour(), base.Minute(), base.Second(), base.Nanosecond(), base.Location())

	// IF EPOCH IS PASSED (PRIOR TO START), THEN MAKE IT HAPPEN ON NEXT MONTH
	if epoch.Before(d1) {
		month++
	}

	// IF EPOCH DAY IS LESS OR EQUAL 28 DAYS THEN
	if day > 28 {
		// SPECIAL CASE FOR FEBRUARY WITH LEAP YEAR HANDLING
		if month == time.February {
			if isLeap(epoch.Year()) {
				day = 29
			} else {
				day = 28
			}
		} else {
			// LET'S ASSUME END DAY IS EPOCH DAY
			if day > endingDays[month] {
				day = endingDays[month]
			}
		}
	}

	// absolute epoch
	epoch = time.Date(year, month, day, epoch.Hour(), epoch.Minute(), epoch.Second(), epoch.Nanosecond(), epoch.Location())
	return
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// // SnapToInstance date takes an arbitrary date b and returns the closest instance
// // date on or before the d.
// //
// // INPUTS
// //   b  the date near the instance you want
// //  d1  any instance start date
// //-------------------------------------------------------------------------------
// func SnapToInstanceDate(b, d1 time.Time, cycle int64) time.Time {
//
// 	// DECIDE BASED ON CYCLE
// 	switch cycle {
// 	case RECURNONE:
// 		return d;
//
// 	case RECURSECONDLY:
// 		return time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), d1.Minute(), d1.Second(), b.Nanosecond(), b.Location())
//
// 	case RECURMINUTELY:
// 		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), d1.Minute(), b.Second(), b.Nanosecond(), b.Location())
//
// 	case RECURHOURLY:
// 		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), d1.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())
//
// 		// IF EPOCH IS PASSED (PRIOR TO START), THEN MAKE IT HAPPEN ON NEXT HOUR
// 		if epoch.Before(d1) {
// 			epoch = epoch.Add(time.Hour)
// 		}
//
// 	case RECURDAILY:
// 		epoch = time.Date(d1.Year(), d1.Month(), d1.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())
//
// 	case RECURWEEKLY:
// 		epoch = GetWeeklyEpochDate(b, d1)
//
// 	case RECURMONTHLY:
// 		return time.Date(b.Year(), b.Month(), d1.Day(), d1.Hour(), d1.Minute(), d1.Second(), d1.Nanosecond(), d1.Location())
//
// 	case RECURQUARTERLY:
// 		epoch = time.Date(d1.Year(), d1.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())
//
// 	case RECURYEARLY:
// 		epoch = time.Date(d1.Year(), b.Month(), b.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())
//
// 		// IF EPOCH IS PASSED (PRIOR TO START), THEN MAKE IT HAPPEN ON NEXT YEAR
// 		if epoch.Before(d1) {
// 			epoch = epoch.AddDate(1, 0, 0)
// 		}
// 	}
//
// 	// IF EPOCH FALLS BEFORE THE *END DATE* (i.e, IN GIVEN DATE RANGE)
// 	// THEN MARK "OK" FLAG AS TRUE
// 	if epoch.Before(d2) {
// 		ok = true
// 	}
//
// 	return
// }
