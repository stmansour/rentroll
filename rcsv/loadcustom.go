package rcsv

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"strings"
)

// 0    1              2       	   3		4
// BUD, Name, 	      ValueType,  Value,	Units
// REX, "Square Feet", 0-2 , 	   "1638",  "sqft"

// CreateCustomAttributes reads a CustomAttributes string array and creates a database record
func CreateCustomAttributes(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateCustomAttributes"
	var (
		err    error
		errmsg string
		c      rlib.CustomAttribute
	)

	const (
		BUD       = 0
		Name      = iota
		ValueType = iota
		Value     = iota
		Units     = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"ValueType", ValueType},
		{"Value", Value},
		{"Units", Units},
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
		b2, err := rlib.GetBusinessByDesignation(ctx, cmpdes)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find Business named %s", funcname, lineno, cmpdes)
		}
		if b2.BID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find Business named %s", funcname, lineno, cmpdes)
		}
		c.BID = b2.BID
	}

	c.Type, err = rlib.IntFromString(sa[ValueType], "Type is invalid")
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, err.Error())
	}
	if c.Type < rlib.CUSTSTRING || c.Type > rlib.CUSTLAST {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Type value must be a number from %d to %d", funcname, lineno, rlib.CUSTSTRING, rlib.CUSTLAST)
	}

	c.Name = strings.TrimSpace(sa[Name])
	c.Value = strings.TrimSpace(sa[Value])
	c.Units = strings.TrimSpace(sa[Units])
	switch c.Type {
	case rlib.CUSTINT:
		_, err = rlib.IntFromString(c.Value, "Value cannot be converted to an integer")
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, err.Error())
		}
	case rlib.CUSTUINT:
		_, err = rlib.IntFromString(c.Value, "Value cannot be converted to an unsigned integer")
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, err.Error())
		}
	case rlib.CUSTFLOAT:
		_, errmsg = rlib.FloatFromString(c.Value, "Value cannot be converted to an float")
		if len(errmsg) > 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, errmsg)
		}
	}

	dup, err := rlib.GetCustomAttributeByVals(ctx, c.Type, c.Name, c.Value, c.Units)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error checking for duplicate Custom Attributes: %s", funcname, lineno, err.Error())
	}
	if dup.CID > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s:: skipping this because a custom attribute with Type = %d, Name = %s, Value = %s, Units = %s already exists", funcname, lineno, DupCustomAttribute, c.Type, c.Name, c.Value, c.Units)
	}

	_, err = rlib.InsertCustomAttribute(ctx, &c)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not insert CustomAttribute. err = %v", funcname, lineno, err)
	}
	return 0, nil
}

// LoadCustomAttributesCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadCustomAttributesCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateCustomAttributes)
}
