package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// 0            1        2
// ElementType  ID       CID
//  5 ,         123,     456

// CreateCustomAttributeRefs reads a rlib.CustomAttributeRefs string array and creates a database record
func CreateCustomAttributeRefs(sa []string, lineno int) {
	funcname := "Createrlib.CustomAttributeRefs"
	var ok bool
	var c rlib.CustomAttributeRef

	if strings.ToLower(sa[0]) == "elementtype" {
		return // it's the header line
	}

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 3
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	c.ElementType, ok = rlib.IntFromString(sa[0], "ElementType is invalid")
	if !ok {
		return
	}
	if c.ElementType < rlib.ELEMRENTABLETYPE || c.ElementType > rlib.ELEMLAST {
		fmt.Printf("ElementType value must be a number from %d to %d\n", rlib.ELEMRENTABLETYPE, rlib.ELEMLAST)
		return
	}

	c.ID, ok = rlib.IntFromString(sa[1], "ID value cannot be converted to an integer")
	if !ok {
		return
	}
	c.CID, ok = rlib.IntFromString(sa[2], "CID value cannot be converted to an integer")
	if !ok {
		return
	}

	switch c.ElementType {
	case rlib.ELEMRENTABLETYPE:
		var rt rlib.RentableType
		err := rlib.GetRentableType(c.ID, &rt)
		if err != nil {
			fmt.Printf("%s: line %d - Could not load rlib.RentableType with id %d:  error = %v\n", funcname, lineno, c.ID, err)
			return
		}
	}

	err := rlib.InsertCustomAttributeRef(&c)
	if err != nil {
		fmt.Printf("%s: line %d - Could not insert rlib.CustomAttributeRef. err = %v\n", funcname, lineno, err)
	}
}

// LoadCustomAttributeRefsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadCustomAttributeRefsCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateCustomAttributeRefs(t[i], i+1)
	}
}
