package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

// RentableSpecialty is the structure for attributes of a rentable specialty

// type RentableSpecialty struct {
// 	RSPID       int64
// 	BID         int64
// 	Name        string
// 	Fee         float64
// 	Description string
// }

//  CSV file format:
// REH, "Lake View",     50.0,  "Overlooks the lake"
// REH, "Courtyard View",50.0,  "Rear windows view the courtyard"
// REH, "Top Floor",     100.0, "Penthouse"
// REH, "Fireplace",     20.0,  "Wood burning, gas fireplace"

// CreateRentalSpecialty reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRentalSpecialty(sa []string) {
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Check to see if this rental specialty type is already in the database
	//-------------------------------------------------------------------
	var b Business

	if len(des) > 0 {
		b, _ := GetBusinessByDesignation(des)
		if b.BID < 1 {
			Ulog("CreateRentalSpecialtyType: Business named %s not found\n", des)
			return
		}
	}

	var a RentableSpecialty
	var x float64
	var err error

	a.Name = strings.TrimSpace(sa[1])
	if x, err = strconv.ParseFloat(strings.TrimSpace(sa[2]), 64); err != nil {
		Ulog("CreateRentalSpecialty: Invalid floating point number: %s\n", sa[2])
		return
	}
	a.Fee = x
	a.Description = strings.TrimSpace(sa[3])
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Make sure we don't already have an exact business,name match
	//-------------------------------------------------------------------
	rsp := GetSpecialtyByName(a.BID, a.Name)
	if rsp.RSPID > 0 {
		fmt.Printf("CreateRentalSpecialty: Business %s already has a RentableSpecialty named %s\n", des, a.Name)
		return
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	err = InsertRentableSpecialty(&a)
	if nil != err {
		fmt.Printf("CreateRentalSpecialty: error inserting RentalSpecialty = %v\n", err)
	}

}

// LoadRentalSpecialtiesCSV loads a csv file with rental specialty types and processes each one
func LoadRentalSpecialtiesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalSpecialty(t[i])
	}
}
