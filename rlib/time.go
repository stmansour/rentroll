package rlib

import (
	"extres"
	"strings"
	"time"
)

var islocalhost, isNotProd, haveTime bool

// Now is a wrapper around time.Now that enables testing to work well with
// "known-good" or "gold" files. It address issues encountered when testing
// involves the expansion of recurring things:  tasks, assessments, ...
// The problem has to do with expanding past instances to bring them
// up-to-date with the current system date.  As an example, suppose we have
// a monthly recurring Assessment with DtStart = Jan 1, 2018 and
// DtStop = Jan 1, 2020.  If we do the expansion in October 2018, we will get
// 10 instances.  If we do the expansion in November 2018, we will get 11 instances.
// That's fine, but the problem has to do with the "gold" files, the known
// good output.  As time progresses, many internal values will change between
// a test run in October 2018 and the same test being run in November 2018
// if the test involves expanding recurring instances.  The number of
// assessments in a database will change. The ASMIDs of all assessments created
// after a call that expanded assessments will change each month. This means
// that the contents of "gold" files that work in October 2018 will not work
// in November 2018 -- there will be more assessments, and the ASMIDs will
// change.
//
// To address this, we use an internal function: rlib.Now() instead of
// time.Now() to get the "system" date and time.  For normal functional
// operation, rlib.Now() returns time.Now().  But for testing, you can set
// the date that you want it to return.
//
// To ensure that a user-specified date can only be returned when testing and
// that it NEVER be applied during a running session in production, this
// routine will ensure that the database host is "localhost".  Any other value
// will cause it to return time.Now()
//------------------------------------------------------------------------------
func Now() time.Time {
	if islocalhost && isNotProd && haveTime {
		return QAInfra.Now
	}
	return time.Now()
}

// InitQAInfra initializes internal flags so that they don't need to be
// computed on every call to rlib.Now()
//------------------------------------------------------------------------------
func InitQAInfra() {
	islocalhost = strings.Contains(AppConfig.Dbhost, "localhost")
	isNotProd = AppConfig.Env != extres.APPENVPROD
	haveTime = QAInfra.Now.Year() > TIME0.Year()
}

// Earliest returns the earlier of two dates.
//
// INPUTS
//   two dates to compare
//
// RETURNS
//   the earlier of the two dates
//------------------------------------------------------------------------------
func Earliest(d1, d2 *time.Time) time.Time {
	if d1.Before(*d2) {
		return *d1
	}
	return *d2
}

// Latest returns the later of two dates.
//
// INPUTS
//   two dates to compare
//
// RETURNS
//   the later of the two dates
//------------------------------------------------------------------------------
func Latest(d1, d2 *time.Time) time.Time {
	if d1.After(*d2) {
		return *d1
	}
	return *d2
}
