package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// 0              1       	   2		3
// Name, 	      ValueType,    Value,	Units
// "Square Feet", 0-2 , 	   "1638",  "sqft"

// CreateCustomAttributes reads a CustomAttributes string array and creates a database record
func CreateCustomAttributes(sa []string, lineno int) int {
	funcname := "CreateCustomAttributes"
	var ok bool
	var c rlib.CustomAttribute

	const (
		Name      = 0
		ValueType = iota
		Value     = iota
		Units     = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"Name", Name},
		{"ValueType", ValueType},
		{"Value", Value},
		{"Units", Units},
	}

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return 1
	}
	if lineno == 1 {
		return 0
	}

	c.Type, ok = rlib.IntFromString(sa[1], "Type is invalid")
	if !ok {
		return CsvErrorSensitivity
	}
	if c.Type < rlib.CUSTSTRING || c.Type > rlib.CUSTLAST {
		fmt.Printf("Type value must be a number from %d to %d\n", rlib.CUSTSTRING, rlib.CUSTLAST)
		return CsvErrorSensitivity
	}

	c.Name = strings.TrimSpace(sa[0])
	c.Value = strings.TrimSpace(sa[2])
	c.Units = strings.TrimSpace(sa[3])
	switch c.Type {
	case rlib.CUSTINT:
		_, ok = rlib.IntFromString(c.Value, "Value cannot be converted to an integer")
		if !ok {
			return CsvErrorSensitivity
		}
	case rlib.CUSTUINT:
		_, ok = rlib.IntFromString(c.Value, "Value cannot be converted to an unsigned integer")
		if !ok {
			return CsvErrorSensitivity
		}
	case rlib.CUSTFLOAT:
		_, ok = rlib.FloatFromString(c.Value, "Value cannot be converted to an float")
		if !ok {
			return CsvErrorSensitivity
		}
	}

	_, err := rlib.InsertCustomAttribute(&c)
	if err != nil {
		fmt.Printf("%s: line %d - Could not insert CustomAttribute. err = %v\n", funcname, lineno, err)
	}
	return 0
}

// LoadCustomAttributesCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadCustomAttributesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if CreateCustomAttributes(t[i], i+1) > 0 {
			return
		}
	}
}
