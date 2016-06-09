package rlib

import (
	"fmt"
	"strings"
)

// 0            1        2
// ElementType  ID       CID
//  5 ,         123,     456

// CreateCustomAttributeRefs reads a CustomAttributeRefs string array and creates a database record
func CreateCustomAttributeRefs(sa []string, lineno int) {
	funcname := "CreateCustomAttributeRefs"
	var ok bool
	var c CustomAttributeRef

	if strings.ToLower(sa[0]) == "elementtype" {
		return // it's the header line
	}

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 3
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	c.ElementType, ok = IntFromString(sa[0], "ElementType is invalid")
	if !ok {
		return
	}
	if c.ElementType < ELEMRENTABLETYPE || c.ElementType > ELEMLAST {
		fmt.Printf("ElementType value must be a number from %d to %d\n", ELEMRENTABLETYPE, ELEMLAST)
		return
	}

	c.ID, ok = IntFromString(sa[1], "ID value cannot be converted to an integer")
	if !ok {
		return
	}
	c.CID, ok = IntFromString(sa[2], "CID value cannot be converted to an integer")
	if !ok {
		return
	}

	switch c.ElementType {
	case ELEMRENTABLETYPE:
		var rt RentableType
		err := GetRentableType(c.ID, &rt)
		if err != nil {
			fmt.Printf("%s: line %d - Could not load RentableType with id %d:  error = %v\n", funcname, lineno, c.ID, err)
			return
		}
	}

	err := InsertCustomAttributeRef(&c)
	if err != nil {
		fmt.Printf("%s: line %d - Could not insert CustomAttributeRef. err = %v\n", funcname, lineno, err)
	}
}

// LoadCustomAttributeRefsCSV loads a csv file with a chart of accounts and creates Ledger markers for each
func LoadCustomAttributeRefsCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateCustomAttributeRefs(t[i], i+1)
	}
}
