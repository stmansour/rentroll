package rlib

import (
	"encoding/csv"
	"fmt"
	"os"
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
