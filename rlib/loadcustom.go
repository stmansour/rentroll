package rlib

import (
	"fmt"
	"strings"
)

// 0              1       	   2
// Name 	      ValueType    Value
// "Square Feet", 0-2 , 	   "1638"

// CreateCustomAttributes reads a CustomAttributes string array and creates a database record
func CreateCustomAttributes(sa []string, lineno int) {
	funcname := "CreateCustomAttributes"
	var ok bool
	var c CustomAttribute

	if strings.ToLower(sa[0]) == "name" {
		return // it's the header line
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 3
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	c.Type, ok = IntFromString(sa[1], "Type is invalid")
	if !ok {
		return
	}
	if c.Type < CUSTSTRING || c.Type > CUSTLAST {
		fmt.Printf("Type value must be a number from %d to %d\n", CUSTSTRING, CUSTLAST)
		return
	}

	c.Name = sa[0]
	c.Value = sa[2]
	switch c.Type {
	case CUSTINT:
		_, ok = IntFromString(c.Value, "Value cannot be converted to an integer")
		if !ok {
			return
		}
	case CUSTFLOAT:
		_, ok = FloatFromString(c.Value, "Value cannot be converted to an float")
		if !ok {
			return
		}
	}

	_, err := InsertCustomAttribute(&c)
	if err != nil {
		fmt.Printf("%s: line %d - Could not insert CustomAttribute. err = %v\n", funcname, lineno, err)
	}
}

// LoadCustomAttributesCSV loads a csv file with a chart of accounts and creates ledger markers for each
func LoadCustomAttributesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateCustomAttributes(t[i], i+1)
	}
}
