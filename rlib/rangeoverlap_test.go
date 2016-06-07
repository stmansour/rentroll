package rlib

import "testing"

// Edge case range overlap tests

type testData struct {
	d1, d2, d3, d4 string
	expect         bool
}

func TestOverlap(t *testing.T) {
	var m = []testData{
		{"2016-07-01", "2016-07-01", "2014-01-01", "2016-07-05", true},
		{"2016-07-05", "2016-07-05", "2014-01-01", "2016-07-05", false},
		{"2016-07-05", "2016-07-05", "2016-07-05", "2016-07-05", true},  // all the same point
		{"2016-07-05", "2016-07-05", "2016-01-05", "2016-12-05", true},  // point, range
		{"2016-06-05", "2016-08-05", "2016-07-05", "2016-07-05", true},  // range, point
		{"2016-07-05", "2016-07-05", "2016-07-01", "2016-07-05", false}, // point,range -- point matches range end
		{"2016-07-05", "2016-07-05", "2016-07-05", "2016-07-10", true},  // point,range -- point matches range start
		{"2016-07-01", "2016-07-05", "2016-07-05", "2016-07-05", false}, // range,point -- point matches range end
		{"2016-07-05", "2016-07-09", "2016-07-05", "2016-07-05", true},  // range,point -- point matches range start
	}

	for i := 0; i < len(m); i++ {
		a1, _ := StringToDate(m[i].d1)
		a2, _ := StringToDate(m[i].d2)
		s1, _ := StringToDate(m[i].d3)
		s2, _ := StringToDate(m[i].d4)
		r := DateRangeOverlap(&a1, &a2, &s1, &s2)
		t.Logf("DateRangeOverlap( %s, %s, %s, %s )  expect %v, got %v\n", m[i].d1, m[i].d2, m[i].d3, m[i].d4, m[i].expect, r)
		if r != m[i].expect {
			t.Errorf("DateRangeOverlap( %s, %s, %s, %s )  expect %v, got %v\n", m[i].d1, m[i].d2, m[i].d3, m[i].d4, m[i].expect, r)
		}
	}
}
