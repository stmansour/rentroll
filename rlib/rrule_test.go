package rlib

import (
	"testing"
	"time"
)

const longForm = "Jan 2, 2006 3:04pm (MST)"

// Parse a time value from a string in the standard Unix format.
//   	t, err := time.Parse(time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")
//		t, _ := time.Parse(longForm,        "Feb 3, 2013 at 7:54pm (PST)"

//  struct with initializers for dummy assessments
type rrtest struct {
	amt    float32
	start  string
	stop   string
	freq   int
	expect int // expected number of occurrences
}

func runRecurTest(t *testing.T, name string, rrcase []rrtest, dt1, dt2 string) {
	//=====================================================
	//  Time range we'll be testing:  Feb 3 - Feb 4
	//=====================================================
	dtstart, err := time.Parse(longForm, dt1)
	Errcheck(err)
	dtstop, err := time.Parse(longForm, dt2)
	Errcheck(err)

	for i := 0; i < len(rrcase); i++ {
		t1, err := time.Parse(longForm, rrcase[i].start)
		Errcheck(err)
		t2, err := time.Parse(longForm, rrcase[i].stop)
		Errcheck(err)
		m := GetRecurrences(&dtstart, &dtstop, &t1, &t2, rrcase[i].freq)
		for j := 0; j < len(m); j++ {
			t.Logf("%s[%d]:  %2d. %s\n", name, i, j, m[j].Format(longForm))
		}
		if len(m) != rrcase[i].expect {
			t.Errorf("%s[%d]:  len(m)=%d,  Expected: %d", name, i, len(m), rrcase[i].expect)
		}
	}
}

func TestRecurNone(t *testing.T) {
	var rrcase = []rrtest{
		{100.00, "Feb 3, 2016  7:54pm (UTC)", "Feb 3, 2016 7:54pm (UTC)", RECURNONE, 1},
		{100.00, "Feb 3, 2016 12:00am (UTC)", "Feb 3, 2016 12:00am (UTC)", RECURNONE, 1},
		{100.00, "Feb 4, 2016 12:00am (UTC)", "Feb 4, 2016 12:00am (UTC)", RECURNONE, 0},
		{100.00, "Feb 3, 2016 11:59pm (UTC)", "Feb 3, 2016 11:59pm (UTC)", RECURNONE, 1},
		{100.00, "Feb 2, 2016  7:54pm (UTC)", "Feb 2, 2016 7:54pm (UTC)", RECURNONE, 0},
		{100.00, "Feb 4, 2016  7:54pm (UTC)", "Feb 4, 2016 7:54pm (UTC)", RECURNONE, 0},
	}
	runRecurTest(t, "Non-Recurring", rrcase[0:], "Feb 3, 2016 12:00am (UTC)", "Feb 4, 2016 12:00am (UTC)")
}

//                                d1                       d2
//        ---------------------------------------------------------------------------
//        |                       |                        |                        |
//        |    (6)[a1-a2]         |(1)[a1--a2]    (2)[a1-a2]       (7)[a1-a2]       |
//        |                       |                        |                        |
//        |                    (3)[a1---a2]       (4)[a1------a2]                   |
//        |                       |                        |                        |
//        |             (5)[a1---------a2]                 |                        |
//        |                       |                        |                        |
//        |             (8)[a1--------------------------------a2]                   |
//        |                       |                        |                        |
//        |              (9)[a1-a2]                    (10)[a1-a2]                  |
//        |                       |                        |                        |
//       TUE---------------------WED----------------------THU----------------------FRI
//      FEB 1                   FEB 2                    FEB 3                    FEB 4
//
//  d1 = Feb 2 12:00am
//  d2 = Feb 3 12:00am

func TestRecurDaily(t *testing.T) {
	var rrcase = []rrtest{
		{100.00, "Feb 2, 2016  3:00am (UTC)", "Feb 2, 2016 11:00am (UTC)", RECURDAILY, 1}, // 1
		{100.00, "Feb 2, 2016  9:00pm (UTC)", "Feb 3, 2016 12:00am (UTC)", RECURDAILY, 1}, // 2
		{100.00, "Feb 2, 2016 12:00am (UTC)", "Feb 2, 2016  4:30am (UTC)", RECURDAILY, 1}, // 3
		{100.00, "Feb 2, 2016  9:00pm (UTC)", "Feb 3, 2016  3:00am (UTC)", RECURDAILY, 1}, // 4
		{100.00, "Feb 1, 2016 11:59pm (UTC)", "Feb 2, 2016  3:00am (UTC)", RECURDAILY, 1}, // 5
		{100.00, "Feb 1, 2016 10:00am (UTC)", "Feb 1, 2016  2:00pm (UTC)", RECURDAILY, 0}, // 6
		{100.00, "Feb 3, 2016 11:00am (UTC)", "Feb 3, 2016  7:00pm (UTC)", RECURDAILY, 0}, // 7
		{100.00, "Feb 1, 2016 11:59pm (UTC)", "Feb 3, 2016  3:00am (UTC)", RECURDAILY, 1}, // 8
		{100.00, "Feb 1, 2016 10:00pm (UTC)", "Feb 2, 2016 12:00am (UTC)", RECURDAILY, 0}, // 9
		{100.00, "Feb 3, 2016 12:00am (UTC)", "Feb 3, 2016  2:00am (UTC)", RECURDAILY, 0}, // 10
		{100.00, "Feb 5, 2016 11:59pm (UTC)", "Feb 3, 2016 11:59pm (UTC)", RECURDAILY, 0}, // 11 (bogus dates)
	}
	runRecurTest(t, "DAILY", rrcase[0:], "Feb 2, 2016 12:00am (UTC)", "Feb 3, 2016 12:00am (UTC)")
}

func TestEdgesDaily(t *testing.T) {
	var rrcase = []rrtest{
		{100.00, "Feb 2, 2016  3:00am (UTC)", "Feb 2, 2016 11:00am (UTC)", RECURDAILY, 1}, // 1
		{100.00, "Feb 2, 2016  9:00pm (UTC)", "Feb 3, 2016 12:00am (UTC)", RECURDAILY, 1}, // 2
		{100.00, "Feb 2, 2016 12:00am (UTC)", "Feb 2, 2016  4:30am (UTC)", RECURDAILY, 1}, // 3
		{100.00, "Feb 2, 2016  9:00pm (UTC)", "Feb 3, 2016  3:00am (UTC)", RECURDAILY, 2}, // 4
		{100.00, "Feb 1, 2016 11:59pm (UTC)", "Feb 2, 2016  3:00am (UTC)", RECURDAILY, 2}, // 5
		{100.00, "Feb 1, 2016 10:00am (UTC)", "Feb 1, 2016  2:00pm (UTC)", RECURDAILY, 1}, // 6
		{100.00, "Feb 3, 2016 11:00am (UTC)", "Feb 3, 2016  7:00pm (UTC)", RECURDAILY, 1}, // 7
		{100.00, "Feb 1, 2016 11:59pm (UTC)", "Feb 3, 2016  3:00am (UTC)", RECURDAILY, 3}, // 8
		{100.00, "Feb 1, 2016 10:00pm (UTC)", "Feb 2, 2016 12:00am (UTC)", RECURDAILY, 1}, // 9
		{100.00, "Feb 3, 2016 12:00am (UTC)", "Feb 3, 2016  2:00am (UTC)", RECURDAILY, 1}, // 10
		{100.00, "Feb 1, 2016 11:59pm (UTC)", "Feb 2, 2016 12:00am (UTC)", RECURDAILY, 1}, // 11
		{100.00, "Feb 2, 2016 11:59pm (UTC)", "Feb 3, 2016 12:00am (UTC)", RECURDAILY, 1}, // 12
	}
	runRecurTest(t, "DAILY-edge", rrcase[0:], "Feb 1, 2016 12:00am (UTC)", "Feb 4, 2016 12:00am (UTC)")
}

func TestRecurWeekly(t *testing.T) {
	var rrcase = []rrtest{
		{100.00, "Feb  1, 2016 12:30am (UTC)", "Feb  2, 2016 12:00am (UTC)", RECURWEEKLY, 1}, // Case 1
		{100.00, "Feb  4, 2016 12:00am (UTC)", "Feb  8, 2016 12:00am (UTC)", RECURWEEKLY, 1}, // Case 2
		{100.00, "Feb  1, 2016 12:00am (UTC)", "Feb  4, 2016 12:00am (UTC)", RECURWEEKLY, 1}, // Case 3
		{100.00, "Feb  4, 2016 12:00am (UTC)", "Feb  9, 2016 12:00am (UTC)", RECURWEEKLY, 2}, // Case 4
		{100.00, "Jan 20, 2016 12:00am (UTC)", "Feb  4, 2016 11:59pm (UTC)", RECURWEEKLY, 1}, // Case 5
		{100.00, "Jan 15, 2016 12:00am (UTC)", "Jan 30, 2016 11:59pm (UTC)", RECURWEEKLY, 0}, // Case 6
		{100.00, "Mar  7, 2016 12:00am (UTC)", "Mar 10, 2016 12:00am (UTC)", RECURWEEKLY, 0}, // Case 7
		{100.00, "Jan 30, 2016 11:59pm (UTC)", "Mar  2, 2016 12:00am (UTC)", RECURWEEKLY, 5}, // Case 8
		{100.00, "Jan 30, 2016 12:00am (UTC)", "Feb  1, 2016 12:00am (UTC)", RECURWEEKLY, 0}, // Case 9
		{100.00, "Mar  1, 2016 12:00am (UTC)", "Mar  9, 2016 12:00am (UTC)", RECURWEEKLY, 0}, // Case 10
		{100.00, "Feb  1, 2016 12:00am (UTC)", "Mar  1, 2016 12:00am (UTC)", RECURWEEKLY, 5}, // Case
	}
	runRecurTest(t, "WEEKLY", rrcase[0:], "Feb 1, 2016 12:00am (UTC)", "Mar 1, 2016 12:00am (UTC)")
}

func TestRecurMonthly(t *testing.T) {
	var rrcase = []rrtest{
		{100.00, "Feb  1, 2016 12:00am (UTC)", "Mar 17, 2017 12:00am (UTC)", RECURMONTHLY, 1}, // 0
		{100.00, "Jan  4, 2016 12:00am (UTC)", "Jan  1, 2017 12:00am (UTC)", RECURMONTHLY, 2}, // 1
		{100.00, "Jan  1, 2016 12:00am (UTC)", "Jan  1, 2017 12:00am (UTC)", RECURMONTHLY, 2}, // 2
		{100.00, "Jan  4, 2016 12:00am (UTC)", "Feb  9, 2016 12:00am (UTC)", RECURMONTHLY, 2}, // 3
		{100.00, "Jan  1, 2016 12:00am (UTC)", "Feb  1, 2016 12:00am (UTC)", RECURMONTHLY, 1}, // 4
		{100.00, "Jan  1, 2015 12:00am (UTC)", "Dec 31, 2016 12:00am (UTC)", RECURMONTHLY, 2}, // 5
		{100.00, "Apr  1, 2016 12:00am (UTC)", "Jun  1, 2016 12:00am (UTC)", RECURMONTHLY, 0}, // 6
	}
	runRecurTest(t, "MONTHLY", rrcase[0:], "Jan 1, 2016 12:00am (UTC)", "Mar 1, 2016 12:00am (UTC)")
}

func TestRecurMonthly2(t *testing.T) {
	var rrcase = []rrtest{
		{100.00, "Feb  1, 2016 12:00am (UTC)", "Mar 17, 2017 12:00am (UTC)", RECURMONTHLY, 11}, // 0
		{100.00, "Jan  4, 2016 12:00am (UTC)", "Jan  1, 2017 12:00am (UTC)", RECURMONTHLY, 12}, // 1
		{100.00, "Jan  1, 2016 12:00am (UTC)", "Jan  1, 2017 12:00am (UTC)", RECURMONTHLY, 12}, // 2
		{100.00, "Jan  4, 2016 12:00am (UTC)", "Feb  9, 2017 12:00am (UTC)", RECURMONTHLY, 12}, // 3
	}
	runRecurTest(t, "MONTHLY2", rrcase[0:], "Jan 1, 2016 12:00am (UTC)", "Jan 1, 2017 12:00am (UTC)")
}

func TestRecurQuarterly(t *testing.T) {
	var rrcase = []rrtest{
		{100.00, "Feb  1, 2016 12:00am (UTC)", "Mar 17, 2017 12:00am (UTC)", RECURQUARTERLY, 4}, // 0
		{100.00, "Jan  4, 2016 12:00am (UTC)", "Jan  1, 2017 12:00am (UTC)", RECURQUARTERLY, 4}, // 1
		{100.00, "Jan  1, 2016 12:00am (UTC)", "Jan  1, 2017 12:00am (UTC)", RECURQUARTERLY, 4}, // 2
		{100.00, "Jan  4, 2016 12:00am (UTC)", "JUL  9, 2017 12:00am (UTC)", RECURQUARTERLY, 4}, // 3
	}
	runRecurTest(t, "QUARTERLY", rrcase[0:], "Jan 1, 2016 12:00am (UTC)", "Jan 1, 2017 12:00am (UTC)")
}

func TestRecurYearly(t *testing.T) {
	var rrcase = []rrtest{
		{100.00, "Feb  1, 2015 12:00am (UTC)", "Mar 16, 2017 12:00am (UTC)", RECURYEARLY, 1}, // 0
		{100.00, "Jan  4, 2013 12:00am (UTC)", "Jan  1, 2016 12:00am (UTC)", RECURYEARLY, 0}, // 1
		{100.00, "Jan  1, 2017 12:00am (UTC)", "Jan  1, 2018 12:00am (UTC)", RECURYEARLY, 0}, // 2
		{100.00, "Jan  4, 2016 12:00am (UTC)", "JUL  9, 2017 12:00am (UTC)", RECURYEARLY, 1}, // 3
		{100.00, "Jul  4, 2016 12:00am (UTC)", "JUL  9, 2017 12:00am (UTC)", RECURYEARLY, 1}, // 3
	}
	runRecurTest(t, "YEARLY", rrcase[0:], "Jan 1, 2016 12:00am (UTC)", "Jan 1, 2017 12:00am (UTC)")
}
