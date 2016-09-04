package rcsv

import (
	"fmt"
	"rentroll/rlib"
)

// 0            1        2
// ElementType  ID       CID
//  5 ,         123,     456

// CreateCustomAttributeRefs reads a rlib.CustomAttributeRefs string array and creates a database record
func CreateCustomAttributeRefs(sa []string, lineno int) int {
	funcname := "Createrlib.CustomAttributeRefs"
	var ok bool
	var c rlib.CustomAttributeRef

	const (
		ElementType = 0
		ID          = iota
		CID         = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"ElementType", ElementType},
		{"ID", ID},
		{"CID", CID},
	}

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return 1
	}
	if lineno == 1 {
		return 0
	}

	c.ElementType, ok = rlib.IntFromString(sa[0], "ElementType is invalid")
	if !ok {
		return CsvErrorSensitivity
	}
	if c.ElementType < rlib.ELEMRENTABLETYPE || c.ElementType > rlib.ELEMLAST {
		fmt.Printf("ElementType value must be a number from %d to %d\n", rlib.ELEMRENTABLETYPE, rlib.ELEMLAST)
		return CsvErrorSensitivity
	}

	c.ID, ok = rlib.IntFromString(sa[1], "ID value cannot be converted to an integer")
	if !ok {
		return CsvErrorSensitivity
	}
	c.CID, ok = rlib.IntFromString(sa[2], "CID value cannot be converted to an integer")
	if !ok {
		return CsvErrorSensitivity
	}

	switch c.ElementType {
	case rlib.ELEMRENTABLETYPE:
		var rt rlib.RentableType
		err := rlib.GetRentableType(c.ID, &rt)
		if err != nil {
			fmt.Printf("%s: line %d - Could not load rlib.RentableType with id %d:  error = %v\n", funcname, lineno, c.ID, err)
			return CsvErrorSensitivity
		}
	}

	err := rlib.InsertCustomAttributeRef(&c)
	if err != nil {
		fmt.Printf("%s: line %d - Could not insert CustomAttributeRef. err = %v\n", funcname, lineno, err)
	}
	return 0
}

// LoadCustomAttributeRefsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadCustomAttributeRefsCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if CreateCustomAttributeRefs(t[i], i+1) > 0 {
			return
		}
	}
}
