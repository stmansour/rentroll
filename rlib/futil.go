package rlib

// This file is a random collection of utility routines...

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Stripchars returns a string with the characters from chars removed
func Stripchars(str, chars string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chars, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func yesnoToInt(si string) (int64, error) {
	s := strings.ToUpper(si)
	switch {
	case s == "Y" || s == "YES" || s == "1":
		return YES, nil
	case s == "N" || s == "NO" || s == "0":
		return NO, nil
	default:
		err := fmt.Errorf("Unrecognized yes/no string: %s.", si)
		return NO, err
	}
}

// IsValidAccrual returns true if a is a valid accrual value, false otherwise
func IsValidAccrual(a int64) bool {
	return !(a < ACCRUALNORECUR || a > ACCRUALYEARLY)
}

// AccrualDuration converts an accrual frequency into the time duration it represents
func AccrualDuration(a int64) time.Duration {
	var d = time.Duration(0)
	switch a {
	case ACCRUALNORECUR:
	case ACCRUALSECONDLY:
		d = time.Second
	case ACCRUALMINUTELY:
		d = time.Minute
	case ACCRUALHOURLY:
		d = time.Hour
	case ACCRUALDAILY:
		d = time.Hour * 24
	case ACCRUALWEEKLY:
		d = time.Hour * 24 * 7
	case ACCRUALMONTHLY:
		d = time.Hour * 24 * 30 // yea, I know this isn't exactly right
	case ACCRUALQUARTERLY:
		d = time.Hour * 24 * 90 // yea, I know this isn't exactly right
	case ACCRUALYEARLY:
		d = time.Hour * 24 * 365 // yea, I know this isn't exactly right
	default:
		Ulog("AccrualDuration: invalid accrual value: %d\n", a)
	}
	return d
}

// IsSQLNoResultsError returns true if the error provided is a sql err indicating no rows in the solution set.
func IsSQLNoResultsError(err error) bool {
	s := fmt.Sprintf("%v", err)
	return strings.Contains(s, "no rows in result")
}

// IntFromString converts the supplied string to an int64 value. If there
// is a problem in the conversion, it generates an error message.
func IntFromString(sa string, errmsg string) (int64, bool) {
	var n = int64(0)
	s := strings.TrimSpace(sa)
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("CreateAssessmentsFromCSV: %s: %s\n", errmsg, s)
			return n, false
		}
		n = int64(i)
	}
	return n, true
}

// FloatFromString converts the supplied string to an int64 value. If there
// is a problem in the conversion, it generates an error message.
func FloatFromString(sa string, errmsg string) (float64, bool) {
	var f = float64(0)
	s := strings.TrimSpace(sa)
	if len(s) > 0 {
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			Ulog("CreateAssessmentsFromCSV: I%s: %s\n", errmsg, sa)
			return f, false
		}
		f = x
	}
	return f, true
}

// LoadCSV loads a comma-separated-value file into an array of strings and returns the array of strings
func LoadCSV(fname string) [][]string {
	t := [][]string{}
	f, err := os.Open(fname)
	if nil == err {
		defer f.Close()
		reader := csv.NewReader(f)
		reader.FieldsPerRecord = -1
		rawCSVdata, err := reader.ReadAll()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, sa := range rawCSVdata {
			t = append(t, sa)
		}
	} else {
		Ulog("LoadCSV: could not open CSV file. err = %v\n", err)
	}
	return t
}
