package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//  0     1            2        3
//  BUD,  ElementType, ID,      CID
//  REX,  5 ,          123,     456

// CreateCustomAttributeRefs reads a rlib.CustomAttributeRefs string array and creates a database record
func CreateCustomAttributeRefs(sa []string, lineno int) (int, error) {
	funcname := "Createrlib.CustomAttributeRefs"
	var c rlib.CustomAttributeRef

	const (
		BUD         = 0
		ElementType = iota
		ID          = iota
		CID         = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"ElementType", ElementType},
		{"ID", ID},
		{"CID", CID},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	//-------------------------------------------------------------------
	// BUD
	//-------------------------------------------------------------------
	cmpdes := strings.TrimSpace(sa[BUD])
	if len(cmpdes) > 0 {
		b2 := rlib.GetBusinessByDesignation(cmpdes)
		if b2.BID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find Business named %s", funcname, lineno, cmpdes)
		}
		c.BID = b2.BID
	}

	c.ElementType, err = rlib.IntFromString(sa[ElementType], "ElementType is invalid")
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, err.Error())
	}
	if c.ElementType < rlib.ELEMRENTABLETYPE || c.ElementType > rlib.ELEMLAST {
		return CsvErrorSensitivity, fmt.Errorf("ElementType value must be a number from %d to %d", rlib.ELEMRENTABLETYPE, rlib.ELEMLAST)
	}

	c.ID, err = rlib.IntFromString(sa[ID], "ID value cannot be converted to an integer")
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, err.Error())
	}
	c.CID, err = rlib.IntFromString(sa[CID], "CID value cannot be converted to an integer")
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, err.Error())
	}

	switch c.ElementType {
	case rlib.ELEMRENTABLETYPE:
		var rt rlib.RentableType
		err := rlib.GetRentableType(c.ID, &rt)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not load rlib.RentableType with id %d:  error = %v", funcname, lineno, c.ID, err)
		}
	}

	ref := rlib.GetCustomAttributeRef(c.ElementType, c.ID, c.CID)
	if ref.ElementType == c.ElementType && ref.CID == c.CID && ref.ID == c.ID {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - This reference already exists, no changes made", funcname, lineno)
	}

	err = rlib.InsertCustomAttributeRef(&c)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not insert CustomAttributeRef. err = %v", funcname, lineno, err)
	}
	return 0, nil
}

// LoadCustomAttributeRefsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadCustomAttributeRefsCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateCustomAttributeRefs)
}
