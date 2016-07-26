package rlib

import "testing"

type testRCycle struct {
	d1, d2 string
	c      int64
	expect int64
}

func TestRentCycles(t *testing.T) {
	var m = []testRCycle{
		{"2016-07-01", "2016-07-01", 6, 0},  // monthly: 1 day
		{"2016-07-01", "2016-07-05", 6, 0},  // monhtly: 4 days
		{"2016-07-01", "2016-08-01", 6, 1},  // monthly: 1 month
		{"2016-07-01", "2016-08-05", 6, 1},  // monthly: 1 month 4 days
		{"2016-07-01", "2016-09-01", 6, 2},  // monthly: 2 months
		{"2016-07-01", "2016-10-01", 6, 3},  // monthly: 3 months
		{"2016-07-01", "2016-11-01", 6, 4},  // monthly: 4 months
		{"2016-07-01", "2017-01-01", 6, 6},  // monthly: 6 months
		{"2016-07-01", "2017-07-01", 6, 12}, // monthly: 1 year
		{"2016-07-01", "2016-08-01", 4, 31}, // daily: 1 month
		{"2016-06-01", "2016-07-01", 4, 30}, // daily: 1 month
	}

	for i := 0; i < len(m); i++ {
		d1, _ := StringToDate(m[i].d1)
		d2, _ := StringToDate(m[i].d2)
		r := CalculateNumberOfCycles(&d1, &d2, m[i].c)
		t.Logf("CalculateNumberOfCycles( %s, %s, %s )  expect %d, got %d\n", d1.Format(RRDATEFMT3), d2.Format(RRDATEFMT3), RentalPeriodToString(m[i].c), m[i].expect, r)
		if r != m[i].expect {
			t.Errorf("CalculateNumberOfCycles( %s, %s, %s )  expect %d, got %d\n", d1.Format(RRDATEFMT3), d2.Format(RRDATEFMT3), RentalPeriodToString(m[i].c), m[i].expect, r)
		}
	}
}
