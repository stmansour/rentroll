package rlib

import (
	"extres"
	"fmt"
	"log"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
)

// RRCommaf returns a floating point number formated with commas for every 3 orders of magnitude
// and 2 points after the decimal
func RRCommaf(x float64) string {
	return humanize.FormatFloat("#,###.##", x)
}

// Ulog is Phonebooks's standard logger
func Ulog(format string, a ...interface{}) {
	p := fmt.Sprintf(format, a...)
	log.Print(p)
	// debug.PrintStack()
}

// Stripchars returns a string with the characters from chars removed
func Stripchars(str, chars string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(chars, r) {
			return r
		}
		return -1
	}, str)
}

// Errcheck - saves a bunch of typing, prints error if it exists
//            and provides a traceback as well
// Note that the error is printed only if the environment is NOT production.
func Errcheck(err error) {
	if err != nil {
		if IsSQLNoResultsError(err) {
			return
		}
		if extres.APPENVPROD != AppConfig.Env {
			fmt.Printf("error = %v\n", err)
		}
		debug.PrintStack()
		log.Fatal(err)
	}
}

// LogAndPrint logs and prints the message
func LogAndPrint(format string, a ...interface{}) {
	p := fmt.Sprintf(format, a...)
	log.Print(p)
	fmt.Print(p)
	// debug.PrintStack()
}

// LogAndPrintError encapsulates logging and printing an error.
// Note that the error is printed only if the environment is NOT production.
func LogAndPrintError(funcname string, err error) {
	errmsg := fmt.Sprintf("%s: err = %v\n", funcname, err)
	Ulog(errmsg)
	if extres.APPENVPROD != AppConfig.Env {
		fmt.Println(errmsg)
	}
}

// CheckLogAndPrintError does LogAndPrintError only if the error is not nil
func CheckLogAndPrintError(funcname string, err error) {
	if err != nil {
		LogAndPrintError(funcname, err)
	}
}

// Errlog - logs the error, but does not stop or quit
func Errlog(err error) {
	if err != nil {
		Ulog("error = %v\n", err)
	}
}

// Tline returns a string of dashes that is the specified length
func Tline(n int) string {
	p := make([]byte, n)
	for i := 0; i < n; i++ {
		p[i] = '-'
	}
	return string(p)
}

// Mkstr returns a string of n of the supplied character that is the specified length
func Mkstr(n int, c byte) string {
	p := make([]byte, n)
	for i := 0; i < n; i++ {
		p[i] = c
	}
	return string(p)
}

// RoundToCent rounds the supplied amount to the nearest cent.
func RoundToCent(x float64) float64 {
	var xtra = float64(0.5)
	if x < float64(0) {
		xtra = -xtra
	}
	return float64(int64(x*float64(100)+xtra)) / float64(100)
}

// YesNoToInt takes multiple forms of "Yes" and converts to integer 1, multiple forms of "No" to integer 0
func YesNoToInt(si string) (int64, error) {
	s := strings.ToUpper(strings.TrimSpace(si))
	switch {
	case s == "Y" || s == "YES" || s == "1" || s == "T" || s == "TRUE":
		return YES, nil
	case s == "N" || s == "NO" || s == "0" || s == "F" || s == "FALSE":
		return NO, nil
	default:
		err := fmt.Errorf("Unrecognized yes/no string: %s", si)
		return NO, err
	}
}

// YesNoToString returns an appropriate string representation of the value i assummed to be YES or NO
func YesNoToString(i int64) string {
	switch i {
	case YES:
		return "Yes"
	case NO:
		return "No"
	default:
		return fmt.Sprintf("??? %d", i)
	}
}

// YesNoToBool takes multiple forms of "Yes" and converts to bool "true", multiple forms of "No" to bool "false"
func YesNoToBool(si string) (bool, error) {
	s := strings.ToUpper(strings.TrimSpace(si))
	switch {
	case s == "Y" || s == "YES" || s == "1" || s == "T" || s == "TRUE":
		return true, nil
	case s == "N" || s == "NO" || s == "0" || s == "F" || s == "FALSE":
		return false, nil
	default:
		err := fmt.Errorf("Unrecognized yes/no string: %s", si)
		return false, err
	}
}

// BoolToYesNoString returns an appropriate string representation of the value i bool
func BoolToYesNoString(i bool) string {
	if i {
		return "Yes"
	}
	return "No"
}

// DateToString rounds the supplied amount to the nearest cent.
func DateToString(t time.Time) string {
	return t.Format("01/02/2006")
}

// AcceptedDateFmts is the array of string formats that StringToDate accepts
var AcceptedDateFmts = []string{
	RRDATEINPFMT,
	RRDATEFMT2,
	RRDATEFMT,
	RRDATEFMT3,
	RRJSUTCDATETIME,
	RRDATETIMEW2UIFMT,
	RRDATETIMEINPFMT,
	RRDATETIMEFMT,
	RRDATEREPORTFMT,
	RRDATERECEIPTFMT,
}

// StringToDate tries to convert the supplied string to a time.Time value. It will use the
// formats called out in dbtypes.go:  RRDATEFMT, RRDATEINPFMT, RRDATEINPFMT2, ...
//
// for further experimentation, try: https://play.golang.org/p/JNUnA5zbMoz
//----------------------------------------------------------------------------------
func StringToDate(s string) (time.Time, error) {
	// try the ansi std date format first
	var Dt time.Time
	var err error
	s = strings.TrimSpace(s)
	for i := 0; i < len(AcceptedDateFmts); i++ {
		Dt, err = time.Parse(AcceptedDateFmts[i], s)
		if nil == err {
			return Dt, nil
		}
	}
	return Dt, fmt.Errorf("Date could not be decoded")
}

// DateAtTimeZero returns the supplied date at time 0 UTC of its day,month,year
func DateAtTimeZero(d time.Time) time.Time {
	d1 := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
	return d1
}

// IncMonths enables arithmetic operations on months. Returns
// two values =  years & months.
// @params
//	m = current month
//  n = number of months to increment
func IncMonths(m time.Month, n int64) (time.Month, int64) {
	y := int64(0)
	mo := int64(m) + n - int64(1)
	y += mo / int64(12)
	mo = mo % int64(12)
	m = time.Month(mo + 1)
	return m, y
}

// GetMonthPeriodForDate is used to get the containing month start and end dates for the
// supplied date.  That is, if a = Jul 13, 2017, the return values will be 2017-JUL-01 and
// 2017-AUG-01.
//
// INPUTS  -  a = any datetime
//
// RETURNS    d1 = first day 00:00:00 of the month of a
//            d2 = first day 00:00:00 of the month after a
//----------------------------------------------------------------------------------
func GetMonthPeriodForDate(a *time.Time) (time.Time, time.Time) {
	d1 := time.Date(a.Year(), a.Month(), 1, 0, 0, 0, 0, RRdb.Zone)
	mon, inc := IncMonths(a.Month(), int64(1))
	d2 := time.Date(int(inc)+a.Year(), mon, 1, 0, 0, 0, 0, RRdb.Zone)
	return d1, d2
}

// StringToInt simply converts string to int and returns it with ok flag
func StringToInt(s string) (int, bool) {
	var i int
	var err error

	if i, err = strconv.Atoi(s); err != nil {
		return i, false
	}
	return i, true
}

// IntToString simply converts int to string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// StringToFloat64 simply converts string to float64 and returns it with ok flag
func StringToFloat64(s string) (float64, bool) {
	var f64 float64
	var err error

	if f64, err = strconv.ParseFloat(s, 64); err != nil {
		return f64, false
	}
	return f64, true
}

// Float64ToString simply converts float64 to string
func Float64ToString(f64 float64) string {
	return fmt.Sprintf("%f", f64)
}

// Int64Range holds list of int64 kind of, varibles
// custom sorting, use sort.Sort(Int64Range)
type Int64Range []int64

func (a Int64Range) Len() int           { return len(a) }
func (a Int64Range) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64Range) Less(i, j int) bool { return a[i] < a[j] }

// SortInt64 implements the sort function for a slice of int64.
func SortInt64(a *Int64Range) {
	sort.Sort(*a)
}

// StringToInt64 simply converts string to int64 and returns it with ok flag
func StringToInt64(s string) (int64, bool) {
	var i64 int64
	var err error

	if i64, err = strconv.ParseInt(s, 10, 64); err != nil {
		fmt.Println("StringToInt64:= ", err.Error())
		return i64, false
	}
	return i64, true
}

// Int64InSlice just check that element exists in slice or not
func Int64InSlice(i64 int64, list []int64) bool {
	for _, b := range list {
		if b == i64 {
			return true
		}
	}
	return false
}

// IsDateBefore check whether bT(beforeTime)'s time/day is before of aT(afterTime)'s time
func IsDateBefore(bT, aT time.Time) bool {
	return bT.Before(aT)
}

// DateDiff gives diff between two dates with type of int64
func DateDiff(a, b time.Time) int64 {
	return int64(a.Sub(b))
}

// DebugPrint just an alias for fmt.Printf, made for the development purpose only!
// Once development done, it would be easy to remove all of statements from the entire code
// by this method name identification easily
func DebugPrint(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
