package rcsv

import (
	"fmt"
	"rentroll/rlib"
)

// 0            1        2
// ElementType  ID       CID
//  5 ,         123,     456

// CreateCustomAttributeRefs reads a rlib.CustomAttributeRefs string array and creates a database record
func CreateCustomAttributeRefs(sa []string, lineno int) (string, int) {
	funcname := "Createrlib.CustomAttributeRefs"
	var errmsg string
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

	rs, x := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if x > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	c.ElementType, errmsg = rlib.IntFromString(sa[0], "ElementType is invalid")
	if len(errmsg) > 0 {
		return rs, CsvErrorSensitivity
	}
	if c.ElementType < rlib.ELEMRENTABLETYPE || c.ElementType > rlib.ELEMLAST {
		rs += fmt.Sprintf("ElementType value must be a number from %d to %d\n", rlib.ELEMRENTABLETYPE, rlib.ELEMLAST)
		return rs, CsvErrorSensitivity
	}

	c.ID, errmsg = rlib.IntFromString(sa[1], "ID value cannot be converted to an integer")
	if len(errmsg) > 0 {
		return rs, CsvErrorSensitivity
	}
	c.CID, errmsg = rlib.IntFromString(sa[2], "CID value cannot be converted to an integer")
	if len(errmsg) > 0 {
		return rs, CsvErrorSensitivity
	}

	switch c.ElementType {
	case rlib.ELEMRENTABLETYPE:
		var rt rlib.RentableType
		err := rlib.GetRentableType(c.ID, &rt)
		if err != nil {
			rs += fmt.Sprintf("%s: line %d - Could not load rlib.RentableType with id %d:  error = %v\n", funcname, lineno, c.ID, err)
			return rs, CsvErrorSensitivity
		}
	}

	ref := rlib.GetCustomAttributeRef(c.ElementType, c.ID, c.CID)
	if ref.ElementType == c.ElementType && ref.CID == c.CID && ref.ID == c.ID {
		rs += fmt.Sprintf("%s: line %d - This reference already exists. No changes were made.\n", funcname, lineno)
		return rs, CsvErrorSensitivity
	}

	err := rlib.InsertCustomAttributeRef(&c)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - Could not insert CustomAttributeRef. err = %v\n", funcname, lineno, err)
	}
	return rs, 0
}

// LoadCustomAttributeRefsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadCustomAttributeRefsCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateCustomAttributeRefs(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
