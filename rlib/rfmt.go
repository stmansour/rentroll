package rlib

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
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
		if APPENVPROD != AppConfig.Env {
			fmt.Printf("error = %v\n", err)
		}
		debug.PrintStack()
		log.Fatal(err)
	}
}

// LogAndPrintError encapsulates logging and printing an error.
// Note that the error is printed only if the environment is NOT production.
func LogAndPrintError(funcname string, err error) {
	errmsg := fmt.Sprintf("%s: err = %v\n", funcname, err)
	Ulog(errmsg)
	if APPENVPROD != AppConfig.Env {
		fmt.Println(errmsg)
	}
}

// CheckLogAndPrintError does LogAndPrintError only if the error is not nil
func CheckLogAndPrintError(funcname string, err error) {
	if err != nil {
		LogAndPrintError(funcname, err)
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

// Errlog - logs the error, but does not stop or quit
func Errlog(err error) {
	if err != nil {
		Ulog("error = %v\n", err)
	}
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
	case s == "Y" || s == "YES" || s == "1":
		return YES, nil
	case s == "N" || s == "NO" || s == "0":
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

// DateToString rounds the supplied amount to the nearest cent.
func DateToString(t time.Time) string {
	return t.Format("01/02/2006")
}

// RentableStatusString is an array of strings representing
var RentableStatusString = []string{
	"unknown",
	"online",
	"admin",
	"employee",
	"owner occupied",
	"offline",
}

// RentableStatusToString returns a string representation for the status value
func RentableStatusToString(n int64) string {
	if n > int64(len(RentableStatusString)-1) || n < 0 {
		n = 0
	}
	return RentableStatusString[n]
}

var acceptedDateFmts = []string{
	RRDATEINPFMT,
	RRDATEFMT2,
	RRDATEFMT,
	RRDATEFMT3,
	RRDATETIMEINPFMT,
}

// StringToDate tries to convert the supplied string to a time.Time value. It will use the two
// formats called out in dbtypes.go:  RRDATEFMT, RRDATEINPFMT, RRDATEINPFMT2
func StringToDate(s string) (time.Time, error) {
	// try the ansi std date format first
	var Dt time.Time
	var err error
	s = strings.TrimSpace(s)
	for i := 0; i < len(acceptedDateFmts); i++ {
		Dt, err = time.Parse(acceptedDateFmts[i], s)
		if nil == err {
			break
		}
	}
	return Dt, err
}

// IncMonths enables arithmetic operations on months. Returns
// two values =  years & months.
func IncMonths(m time.Month, n int64) (time.Month, int64) {
	y := int64(0)
	mo := int64(m) + n - int64(1)
	y += mo / int64(12)
	mo = mo % int64(12)
	m = time.Month(mo + 1)
	return m, y
}
