package rcsv

import (
	"fmt"
	"rentroll/rlib"
)

// 0            1        2
// ElementType  ID       CID
//  5 ,         123,     456

// CreateCustomAttributeRefs reads a rlib.CustomAttributeRefs string array and creates a database record
func CreateCustomAttributeRefs(sa []string, lineno int) (int, error) {
	funcname := "Createrlib.CustomAttributeRefs"
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

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	c.ElementType, err = rlib.IntFromString(sa[0], "ElementType is invalid")
	if err != nil {
		return CsvErrorSensitivity, err
	}
	if c.ElementType < rlib.ELEMRENTABLETYPE || c.ElementType > rlib.ELEMLAST {
		return CsvErrorSensitivity, fmt.Errorf("ElementType value must be a number from %d to %d\n", rlib.ELEMRENTABLETYPE, rlib.ELEMLAST)
	}

	c.ID, err = rlib.IntFromString(sa[1], "ID value cannot be converted to an integer")
	if err != nil {
		return CsvErrorSensitivity, err
	}
	c.CID, err = rlib.IntFromString(sa[2], "CID value cannot be converted to an integer")
	if err != nil {
		return CsvErrorSensitivity, err
	}

	switch c.ElementType {
	case rlib.ELEMRENTABLETYPE:
		var rt rlib.RentableType
		err := rlib.GetRentableType(c.ID, &rt)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not load rlib.RentableType with id %d:  error = %v\n", funcname, lineno, c.ID, err)
		}
	}

	ref := rlib.GetCustomAttributeRef(c.ElementType, c.ID, c.CID)
	if ref.ElementType == c.ElementType && ref.CID == c.CID && ref.ID == c.ID {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - This reference already exists. No changes were made.\n", funcname, lineno)
	}

	err = rlib.InsertCustomAttributeRef(&c)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not insert CustomAttributeRef. err = %v\n", funcname, lineno, err)
	}
	return 0, nil
}

// LoadCustomAttributeRefsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadCustomAttributeRefsCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateCustomAttributeRefs)
}
