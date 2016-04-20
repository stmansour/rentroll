package rlib

import (
	"encoding/csv"
	"fmt"
	"os"
)

// LoadCSV loads a comma-separated-value file into an array of strings and returns the array of strings
func LoadCSV(fname string) [][]string {
	t := [][]string{}
	f, err := os.Open(fname)
	Errcheck(err)
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
	return t
}
