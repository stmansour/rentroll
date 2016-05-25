package rlib

import (
	"fmt"
	"strings"
)

// 0       1               2
// Type    Name 	       Value
//  0-2 ,  "Square Feet",  "1638"

// CreateCustomAttributes reads a CustomAttributes string array and creates a database record
func CreateCustomAttributes(sa []string, lineno int) {
	funcname := "CreateCustomAttributes"
	var ok bool
	var c CustomAttribute

	if strings.ToLower(sa[0]) == "type" {
		return // it's the header line
	}

	c.Type, ok = IntFromString(sa[0], "Type is invalid")
	if !ok {
		return
	}
	if c.Type < CUSTSTRING || c.Type > CUSTLAST {
		fmt.Printf("Type value must be a number from %d to %d\n", CUSTSTRING, CUSTLAST)
		return
	}

	c.Name = sa[1]
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
