package rlib

import (
	"testing"
	"time"
)

/*// Business Properties Epochs
var bizEpochs = BizPropsEpochs{
	Daily:     time.Date(2018, 07, 21, 6, 0, 0, 0, time.UTC),  // DAILY     - AT 6 AM
	Weekly:    time.Date(2018, 07, 21, 15, 0, 0, 0, time.UTC), // WEEKLY    - AT EVERY SATURDAY - 3 PM
	Monthly:   time.Date(2018, 07, 01, 3, 0, 0, 0, time.UTC),  // MONTHLY   - AT EVERY 1st DAY OF A MONTH AT 3 AM
	Quarterly: time.Date(2018, 07, 21, 0, 0, 0, 0, time.UTC),  // QUARTERLY - IGNORE THIS ONE
	Yearly:    time.Date(2018, 06, 01, 0, 0, 0, 0, time.UTC),  // YEARLY    - ON 1st JUNE AT 12 AM
}*/

// epochCase struct with required inputs and expected output
type epochCase struct {
	base, start, stop time.Time // base, start, stop date -- inputs
	epoch             time.Time // Expected epoch
	ok                bool      // if epoch falls in the range
}

// TestWeeklyEpochDate weekly epoch date test
func TestWeeklyEpochDate(t *testing.T) {

	// cycle to consider
	cycle := int64(RECURWEEKLY)

	// various different cases
	var cases = []epochCase{
		{
			base:  time.Date(2018, 07, 21, 15, 0, 0, 0, time.UTC),  // WEEKLY    - AT EVERY SATURDAY - 3 PM
			start: time.Date(2018, 07, 21, 14, 24, 0, 0, time.UTC), // 21 Jul 2018 2:24 PM, SATURDAY
			stop:  time.Date(2018, 07, 29, 0, 0, 0, 0, time.UTC),   // 29 Jul 2018 12 AM, SUNDAY
			epoch: time.Date(2018, 07, 21, 15, 0, 0, 0, time.UTC),  // ON THE SAME DAY AT 3 PM
			ok:    true,
		},
		{
			base:  time.Date(2018, 07, 21, 15, 0, 0, 0, time.UTC),   // WEEKLY    - AT EVERY SATURDAY - 3 PM
			start: time.Date(2018, 07, 21, 16, 10, 30, 0, time.UTC), // 21 Jul 2018 4:10:30 PM, SATURDAY
			stop:  time.Date(2018, 07, 29, 0, 0, 0, 0, time.UTC),    // 29 Jul 2018 12 AM, SUNDAY
			epoch: time.Date(2018, 07, 28, 15, 0, 0, 0, time.UTC),   // ON NEXT WEEK SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    true,
		},
		{
			base:  time.Date(2018, 07, 21, 15, 0, 0, 0, time.UTC),  // WEEKLY    - AT EVERY SATURDAY - 3 PM
			start: time.Date(2018, 07, 19, 10, 0, 20, 0, time.UTC), // 21 Jul 2018 10:00:20 AM, THURSDAY
			stop:  time.Date(2018, 07, 29, 0, 0, 0, 0, time.UTC),   // 29 Jul 2018 12 AM, SUNDAY
			epoch: time.Date(2018, 07, 21, 15, 0, 0, 0, time.UTC),  // ON NEXT WEEK SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    true,
		},
		{
			base:  time.Date(2018, 07, 21, 15, 0, 0, 0, time.UTC),   // WEEKLY    - AT EVERY SATURDAY - 3 PM
			start: time.Date(2018, 07, 22, 21, 45, 00, 0, time.UTC), // 21 Jul 2018 10:45 PM, SUNDAY
			stop:  time.Date(2018, 07, 30, 0, 0, 0, 0, time.UTC),    // 29 Jul 2018 12 AM, SUNDAY
			epoch: time.Date(2018, 07, 28, 15, 0, 0, 0, time.UTC),   // ON NEXT WEEK SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    true,
		},
	}

	for i := 0; i < len(cases); i++ {

		// GET THE INPUTS FROM STRUCT
		ok, epoch := GetEpochFromBaseDate(cases[i].base, cases[i].start, cases[i].stop, cycle)

		// LOG MESSAGE
		t.Logf("info: case[%d] : GetEpochFromBaseDate( %s, %s, %s, %s ) expect ( %t, %s ), got ( %t, %s )\n",
			i,
			cases[i].base.Format(RRDATETIMEINPFMT), cases[i].start.Format(RRDATETIMEINPFMT),
			cases[i].stop.Format(RRDATETIMEINPFMT), RentalPeriodToString(cycle),
			cases[i].ok, cases[i].epoch,
			ok, epoch)

		// IF DO NOT EQUAL IN ALL EXPECTED VAR
		if !(cases[i].epoch.Equal(epoch) && cases[i].ok == ok) {
			t.Errorf("error: case[%d] : GetEpochFromBaseDate( %s, %s, %s, %s ) expect ( %t, %s ), got ( %t, %s )\n",
				i,
				cases[i].base.Format(RRDATETIMEINPFMT), cases[i].start.Format(RRDATETIMEINPFMT),
				cases[i].stop.Format(RRDATETIMEINPFMT), RentalPeriodToString(cycle),
				cases[i].ok, cases[i].epoch,
				ok, epoch)
		}
	}
}

// TestMonthlyEpochDate monthly epoch date test
func TestMonthlyEpochDate(t *testing.T) {

	// cycle to consider
	cycle := int64(RECURMONTHLY)

	// various different cases
	var cases = []epochCase{
		{
			base:  time.Date(2018, 7, 1, 3, 0, 0, 0, time.UTC),   // MONTHLY   - AT EVERY 1st DAY OF A MONTH AT 3 AM
			start: time.Date(2018, 6, 21, 10, 0, 0, 0, time.UTC), // 21 Jun 2018 10:00 AM
			stop:  time.Date(2018, 7, 29, 0, 0, 0, 0, time.UTC),  // 29 Jul 2018 12 AM
			epoch: time.Date(2018, 7, 1, 3, 0, 0, 0, time.UTC),   // ON NEXT MONTH SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    true,
		},
		{
			base:  time.Date(2018, 7, 1, 3, 0, 0, 0, time.UTC),     // MONTHLY   - AT EVERY 1st DAY OF A MONTH AT 3 AM
			start: time.Date(2018, 7, 21, 16, 10, 30, 0, time.UTC), // 21 Jul 2018 4:10:30 PM
			stop:  time.Date(2018, 7, 29, 0, 0, 0, 0, time.UTC),    // 29 Jul 2018 12 AM
			epoch: time.Date(2018, 8, 1, 3, 0, 0, 0, time.UTC),     // ON NEXT WEEK SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    false,                                           // AS EPOCH WILL NOT FALL IN THE GIVEN RANGE
		},
		{
			base:  time.Date(2018, 7, 1, 1, 0, 0, 0, time.UTC),    // MONTHLY   - AT EVERY 1st DAY OF A MONTH AT 3 AM
			start: time.Date(2018, 6, 19, 10, 0, 20, 0, time.UTC), // 21 Jul 2018 10:00:20 AM
			stop:  time.Date(2018, 7, 29, 0, 0, 0, 0, time.UTC),   // 29 Jul 2018 12 AM
			epoch: time.Date(2018, 7, 1, 1, 0, 0, 0, time.UTC),    // ON NEXT MONTH SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    true,
		},
		{
			base:  time.Date(2018, 1, 31, 3, 0, 0, 0, time.UTC),   // MONTHLY   - AT EVERY LAST DAY OF A MONTH AT 3 AM
			start: time.Date(2018, 6, 22, 21, 45, 0, 0, time.UTC), // 21 Jul 2018 10:45 PM
			stop:  time.Date(2018, 7, 30, 0, 0, 0, 0, time.UTC),   // 29 Jul 2018 12 AM
			epoch: time.Date(2018, 6, 30, 3, 0, 0, 0, time.UTC),   // ON NEXT MONTH SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    true,
		},
		{
			base:  time.Date(2018, 1, 31, 3, 0, 0, 0, time.UTC),   // MONTHLY   - AT EVERY LAST DAY OF A MONTH AT 3 AM
			start: time.Date(2018, 2, 22, 21, 45, 0, 0, time.UTC), // 22 Feb 2018 10:45 PM
			stop:  time.Date(2018, 4, 30, 0, 0, 0, 0, time.UTC),   // 30 Apr 2018 12 AM
			epoch: time.Date(2018, 2, 28, 3, 0, 0, 0, time.UTC),   // ON NEXT MONTH SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    true,
		},
		{
			base:  time.Date(2018, 7, 31, 3, 0, 0, 0, time.UTC), // MONTHLY   - AT EVERY LAST DAY OF A MONTH AT 3 AM
			start: time.Date(2018, 1, 31, 3, 0, 0, 0, time.UTC), // 1 Jan 2018 3 AM
			stop:  time.Date(2018, 4, 30, 0, 0, 0, 0, time.UTC), // 30 Apr 2018 12 AM
			epoch: time.Date(2018, 1, 31, 3, 0, 0, 0, time.UTC), // ON NEXT MONTH SAME DAY AT SAME TIME OF BASE EPOCH
			ok:    true,
		},
	}

	for i := 0; i < len(cases); i++ {

		// GET THE INPUTS FROM STRUCT
		ok, epoch := GetEpochFromBaseDate(cases[i].base, cases[i].start, cases[i].stop, cycle)

		// LOG MESSAGE
		t.Logf("info: case[%d] : GetEpochFromBaseDate( %s, %s, %s, %s ) expect ( %t, %s ), got ( %t, %s )\n",
			i,
			cases[i].base.Format(RRDATETIMEINPFMT), cases[i].start.Format(RRDATETIMEINPFMT),
			cases[i].stop.Format(RRDATETIMEINPFMT), RentalPeriodToString(cycle),
			cases[i].ok, cases[i].epoch,
			ok, epoch)

		// IF DO NOT EQUAL IN ALL EXPECTED VAR
		if !(cases[i].epoch.Equal(epoch) && cases[i].ok == ok) {
			t.Errorf("error: case[%d] : GetEpochFromBaseDate( %s, %s, %s, %s ) expect ( %t, %s ), got ( %t, %s )\n",
				i,
				cases[i].base.Format(RRDATETIMEINPFMT), cases[i].start.Format(RRDATETIMEINPFMT),
				cases[i].stop.Format(RRDATETIMEINPFMT), RentalPeriodToString(cycle),
				cases[i].ok, cases[i].epoch,
				ok, epoch)
		}
	}
}
