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
func CreateRentalSpecialty(sa []string, lineno int) {
	funcname := "CreateRentalSpecialty"
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 4
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Check to see if this rental specialty type is already in the database
	//-------------------------------------------------------------------
	var b rlib.Business

	if len(des) > 0 {
		b, _ = rlib.GetBusinessByDesignation(des)
		if b.BID < 1 {
			rlib.Ulog("%s: lineno %d  - rlib.Business named %s not found\n", funcname, lineno, des)
			return
		}
	}

	var a rlib.RentableSpecialty
	var x float64
	var err error

	a.Name = strings.TrimSpace(sa[1])
	if x, err = strconv.ParseFloat(strings.TrimSpace(sa[2]), 64); err != nil {
		rlib.Ulog("%s: lineno %d  - Invalid floating point number: %s\n", funcname, lineno, sa[2])
		return
	}
	a.Fee = x
	a.Description = strings.TrimSpace(sa[3])
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Make sure we don't already have an exact rlib.Business,name match
	//-------------------------------------------------------------------
	rsp := rlib.GetRentableSpecialtyTypeByName(a.BID, a.Name)
	if rsp.RSPID > 0 {
		fmt.Printf("%s: lineno %d  - rlib.Business %s already has a rlib.RentableSpecialty named %s\n", funcname, lineno, des, a.Name)
		return
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	err = rlib.InsertRentableSpecialty(&a)
	if nil != err {
		fmt.Printf("%s: lineno %d  - error inserting RentalSpecialty = %v\n", funcname, lineno, err)
	}

}

// LoadRentalSpecialtiesCSV loads a csv file with rental specialty types and processes each one
func LoadRentalSpecialtiesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalSpecialty(t[i], i+1)
	}
}
