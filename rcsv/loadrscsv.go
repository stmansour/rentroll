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
func CreateRentalSpecialty(sa []string, lineno int) (string, int) {
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

	rs, y := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if y > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	des := strings.ToLower(strings.TrimSpace(sa[BUD]))

	//-------------------------------------------------------------------
	// Check to see if this rental specialty type is already in the database
	//-------------------------------------------------------------------
	var b rlib.Business

	if len(des) > 0 {
		b = rlib.GetBusinessByDesignation(des)
		if b.BID < 1 {
			rs += fmt.Sprintf("%s: lineno %d  - rlib.Business named %s not found\n", funcname, lineno, des)
			return rs, CsvErrorSensitivity
		}
	}

	var a rlib.RentableSpecialty
	var x float64
	var err error

	a.Name = strings.TrimSpace(sa[Name])
	if x, err = strconv.ParseFloat(strings.TrimSpace(sa[Fee]), 64); err != nil {
		rs += fmt.Sprintf("%s: lineno %d  - Invalid floating point number: %s\n", funcname, lineno, sa[Fee])
		return rs, CsvErrorSensitivity
	}
	a.Fee = x
	a.Description = strings.TrimSpace(sa[Description])
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Make sure we don't already have an exact rlib.Business,name match
	//-------------------------------------------------------------------
	rsp := rlib.GetRentableSpecialtyTypeByName(a.BID, a.Name)
	if rsp.RSPID > 0 {
		rs += fmt.Sprintf("%s: lineno %d  - rlib.Business %s already has a rlib.RentableSpecialty named %s\n", funcname, lineno, des, a.Name)
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	err = rlib.InsertRentableSpecialty(&a)
	if nil != err {
		rs += fmt.Sprintf("%s: lineno %d  - error inserting RentalSpecialty = %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}
	return rs, 0
}

// LoadRentalSpecialtiesCSV loads a csv file with rental specialty types and processes each one
func LoadRentalSpecialtiesCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateRentalSpecialty(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
