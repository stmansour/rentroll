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

// MarshalJSON overrides the default time.Time handler and sends
// date strings of the form YYYY-MM-DD.
//--------------------------------------------------------------------
func (t *JSONTime) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t)
	val := fmt.Sprintf("\"%s\"", ts.Format("2006-01-02"))
	return []byte(val), nil
}

// UnmarshalJSON overrides the default time.Time handler and reads in
// date strings of the form YYYY-MM-DD.
//--------------------------------------------------------------------
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = Stripchars(s, "\"")
	x, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*t = JSONTime(x)
	return nil
}
