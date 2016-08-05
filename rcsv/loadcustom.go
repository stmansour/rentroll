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
func CreateCustomAttributes(sa []string, lineno int) {
	funcname := "CreateCustomAttributes"
	var ok bool
	var c rlib.CustomAttribute

	if strings.ToLower(sa[0]) == "name" {
		return // it's the header line
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 4
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	c.Type, ok = rlib.IntFromString(sa[1], "Type is invalid")
	if !ok {
		return
	}
	if c.Type < rlib.CUSTSTRING || c.Type > rlib.CUSTLAST {
		fmt.Printf("Type value must be a number from %d to %d\n", rlib.CUSTSTRING, rlib.CUSTLAST)
		return
	}

	c.Name = strings.TrimSpace(sa[0])
	c.Value = strings.TrimSpace(sa[2])
	c.Units = strings.TrimSpace(sa[3])
	switch c.Type {
	case rlib.CUSTINT:
		_, ok = rlib.IntFromString(c.Value, "Value cannot be converted to an integer")
		if !ok {
			return
		}
	case rlib.CUSTUINT:
		_, ok = rlib.IntFromString(c.Value, "Value cannot be converted to an unsigned integer")
		if !ok {
			return
		}
	case rlib.CUSTFLOAT:
		_, ok = rlib.FloatFromString(c.Value, "Value cannot be converted to an float")
		if !ok {
			return
		}
	}

	_, err := rlib.InsertCustomAttribute(&c)
	if err != nil {
		fmt.Printf("%s: line %d - Could not insert CustomAttribute. err = %v\n", funcname, lineno, err)
	}
}

// LoadCustomAttributesCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadCustomAttributesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateCustomAttributes(t[i], i+1)
	}
}
