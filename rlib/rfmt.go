package rlib

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"
	"time"
)

// RECURNONE - RECURLAST are the allowed recurrence types
const (
	RECURNONE      = 0
	RECURSECONDLY  = 1
	RECURMINUTELY  = 2
	RECURHOURLY    = 3
	RECURDAILY     = 4
	RECURWEEKLY    = 5
	RECURMONTHLY   = 6
	RECURQUARTERLY = 7
	RECURYEARLY    = 8
	RECURLAST      = RECURYEARLY
)

// Ulog is Phonebooks's standard logger
func Ulog(format string, a ...interface{}) {
	p := fmt.Sprintf(format, a...)
	log.Print(p)
	// debug.PrintStack()
}

// Errcheck - saves a bunch of typing, prints error if it exists
//            and provides a traceback as well
func Errcheck(err error) {
	if err != nil {
		fmt.Printf("error = %v\n", err)
		debug.PrintStack()
		log.Fatal(err)
	}
}

// Errlog - logs the error, but does not stop or quit
func Errlog(err error) {
	if err != nil {
		Ulog("error = %v\n", err)
	}
}

// LogAndPrintError encapsulates logging and printing an error
func LogAndPrintError(funcname string, err error) {
	errmsg := fmt.Sprintf("%s: err = %v\n", funcname, err)
	Ulog(errmsg)
	fmt.Println(errmsg)
}

// RoundToCent rounds the supplied amount to the nearest cent.
func RoundToCent(x float64) float64 {
	return float64(int64(x*float64(100)+float64(0.5))) / float64(100)
}

// DateToString rounds the supplied amount to the nearest cent.
func DateToString(t time.Time) string {
	return t.Format("01/02/2006")
}

// RecurStringToInt supply a recurrence string and the int64  representation is returned
func RecurStringToInt(s string) int64 {
	var i int64
	s = strings.ToUpper(s)
	s = strings.Replace(s, " ", "", -1)
	switch {
	case s == "NONE":
		i = RECURNONE
	case s == "HOURLY":
		i = RECURHOURLY
	case s == "DAILY":
		i = RECURDAILY
	case s == "WEEKLY":
		i = RECURWEEKLY
	case s == "MONTHLY":
		i = RECURMONTHLY
	case s == "QUARTERLY":
		i = RECURQUARTERLY
	case s == "YEARLY":
		i = RECURYEARLY
	default:
		fmt.Printf("Unknown recurrence type: %s\n", s)
		i = RECURNONE
	}
	return i
}

// MonthToInt enables arithmetic operation on months
func MonthToInt(m time.Month) int64 {
	switch m {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12
	}
	return 0 // should never happen
}

// IncMonths enables arithmetic operations on months. Returns
// two values =  years & months.
func IncMonths(m time.Month, n int64) (time.Month, int64) {
	y := int64(0)
	mo := MonthToInt(m) + n - int64(1)
	y += mo / int64(12)
	mo = mo % int64(12)
	switch mo {
	case 0:
		m = time.January
	case 1:
		m = time.February
	case 2:
		m = time.March
	case 3:
		m = time.April
	case 4:
		m = time.May
	case 5:
		m = time.June
	case 6:
		m = time.July
	case 7:
		m = time.August
	case 8:
		m = time.September
	case 9:
		m = time.October
	case 10:
		m = time.November
	case 11:
		m = time.December
	}
	return m, y
}
