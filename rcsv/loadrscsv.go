package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

//  CSV file format:
// 0            1                2      3
// Designation, Name,            Fee,   Description
// REH,         "Lake View",     50.0,  "Overlooks the lake"
// REH,         "Courtyard View",50.0,  "Rear windows view the courtyard"
// REH,         "Top Floor",     100.0, "Penthouse"
// REH,         "Fireplace",     20.0,  "Wood burning, gas fireplace"

// CreateRentalSpecialty reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRentalSpecialty(sa []string, lineno int) (int, error) {
	funcname := "CreateRentalSpecialty"
	const (
		BUD         = 0
		Name        = iota
		Fee         = iota
		Description = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"Fee", Fee},
		{"Description", Description},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	des := strings.ToLower(strings.TrimSpace(sa[BUD]))

	//-------------------------------------------------------------------
	// Check to see if this rental specialty type is already in the database
	//-------------------------------------------------------------------
	var b rlib.Business

	if len(des) > 0 {
		b = rlib.GetBusinessByDesignation(des)
		if b.BID < 1 {
			return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  - rlib.Business named %s not found\n", funcname, lineno, des)
		}
	}

	var a rlib.RentableSpecialty
	var x float64

	a.Name = strings.TrimSpace(sa[Name])
	if x, err = strconv.ParseFloat(strings.TrimSpace(sa[Fee]), 64); err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  - Invalid floating point number: %s\n", funcname, lineno, sa[Fee])
	}
	a.Fee = x
	a.Description = strings.TrimSpace(sa[Description])
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Make sure we don't already have an exact rlib.Business,name match
	//-------------------------------------------------------------------
	rsp := rlib.GetRentableSpecialtyTypeByName(a.BID, a.Name)
	if rsp.RSPID > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  - rlib.Business %s already has a rlib.RentableSpecialty named %s\n", funcname, lineno, des, a.Name)
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	err = rlib.InsertRentableSpecialty(&a)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  - error inserting RentalSpecialty = %v\n", funcname, lineno, err)
	}
	return 0, nil
}

// LoadRentalSpecialtiesCSV loads a csv file with rental specialty types and processes each one
func LoadRentalSpecialtiesCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateRentalSpecialty)
}
