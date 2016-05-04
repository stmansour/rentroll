package rlib

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func yesnoToInt(s string) int64 {
	s = strings.ToUpper(s)
	switch {
	case s == "Y" || s == "YES" || s == "1":
		return YES
	case s == "N" || s == "NO" || s == "0":
		return NO
	default:
		fmt.Printf("Unrecognized yes/no string: %s. Returning default = No\n", s)
		return NO
	}
}

// IsSQLNoResultsError returns true if the error provided is a sql err indicating no rows in the solution set.
func IsSQLNoResultsError(err error) bool {
	s := fmt.Sprintf("%v", err)
	return strings.Contains(s, "no rows in result")
}

// IntFromString converts the supplied string to an int64 value. If there
// is a problem in the conversion, it generates an error message.
func IntFromString(sa string, errmsg string) int64 {
	var n = int64(0)
	s := strings.TrimSpace(sa)
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("CreateAssessmentsFromCSV: %s: %s\n", errmsg, s)
			return n
		}
		n = int64(i)
	}
	return n
}

// FloatFromString converts the supplied string to an int64 value. If there
// is a problem in the conversion, it generates an error message.
func FloatFromString(sa string, errmsg string) float64 {
	var f = float64(0)
	s := strings.TrimSpace(sa)
	if len(s) > 0 {
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			Ulog("CreateAssessmentsFromCSV: I%s: %s\n", errmsg, sa)
			return f
		}
		f = x
	}
	return f
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
	}
	return t
}
