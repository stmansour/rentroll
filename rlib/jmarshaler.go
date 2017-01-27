package rlib

import (
	"fmt"
	"time"
)

// JSONTime is a wrapper around time.Time. We need it
// in order to be able to control the formatting used
// on the date values sent to the w2ui controls.  Without
// this wrapper, the default time format used by the
// JSON encoder / decoder does not work with the w2ui
// controls
type JSONTime time.Time

var earliestDate = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)

// MarshalJSON overrides the default time.Time handler and sends
// date strings of the form YYYY-MM-DD. Any date prior to Jan 1, 1900
// is snapped to Jan 1, 1900.
//--------------------------------------------------------------------
func (t *JSONTime) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t)
	if ts.Before(earliestDate) {
		ts = earliestDate
	}
	// val := fmt.Sprintf("\"%s\"", ts.Format("2006-01-02"))
	val := fmt.Sprintf("\"%s\"", ts.Format(RRDATEFMT3))
	return []byte(val), nil
}

// UnmarshalJSON overrides the default time.Time handler and reads in
// date strings of the form YYYY-MM-DD.  Any date prior to Jan 1, 1900
// is snapped to Jan 1, 1900.
//--------------------------------------------------------------------
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = Stripchars(s, "\"")
	// x, err := time.Parse("2006-01-02", s)
	x, err := StringToDate(s)
	if err != nil {
		return err
	}
	if x.Before(earliestDate) {
		x = earliestDate
	}
	*t = JSONTime(x)
	return nil
}

// MarshalJSON deals with XJSONYesNo
//--------------------------------------------------------------------
func (t *XJSONYesNo) MarshalJSON() ([]byte, error) {
	return []byte("\"" + string(*t) + "\""), nil
}

// UnmarshalJSON deals with XJSONYesNo
//--------------------------------------------------------------------
func (t *XJSONYesNo) UnmarshalJSON(b []byte) error {
	*t = XJSONYesNo(Stripchars(string(b), "\""))
	return nil
}
