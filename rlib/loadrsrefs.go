package rlib

import (
	"fmt"
	"strings"
)

// CSV FIELDS FOR THIS MODULE
//    0    1    2                3         4
//    BUD, RID, RentalSpecialty, DtStart,  DtStStop
//    REX, 1,   Lake View,       1/1/2014,
//    REX, 1,   Fireplace,       1/1/2014,

// type RentableSpecialtyRef struct {
// 	RID         int64     // the Rentable to which this record belongs
// 	RSPID       int64     // the rentable specialty type associated with the rentable
// 	DtStart     time.Time // timerange start
// 	DtStop      time.Time // timerange stop
// 	LastModTime time.Time
// 	LastModBy   int64
// }

// CreateRentableSpecialtyRefsCSV reads an assessment type string array and creates a database record for the assessment type
func CreateRentableSpecialtyRefsCSV(sa []string, lineno int) {
	funcname := "CreateRentableSpecialtyRefsCSV"
	var a RentableSpecialtyRef
	var r Rentable
	var err error
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "bud" {
		return // this is just the column heading
	}

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 5
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	var b Business
	if len(des) > 0 {
		b, _ = GetBusinessByDesignation(des)
		if b.BID < 1 {
			Ulog("CreateRentalSpecialtyType: Business named %s not found\n", sa[0])
			return
		}
	}
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Find and set the Rentable
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[1])
	if len(s) > 0 {
		// fmt.Printf("Searching: rentable name = %s, BID = %d\n", s, b.BID)
		r, err = GetRentableByName(s, b.BID)
		if err != nil {
			fmt.Printf("%s: line %d - Error loading Rentable named: %s.  Error = %v\n", funcname, lineno, s, err)
			return
		}
	}
	a.RID = r.RID

	//-------------------------------------------------------------------
	// Make sure we can find the RentableSpecialty
	//-------------------------------------------------------------------
	name := strings.TrimSpace(sa[2])
	rsp := GetRentableSpecialtyTypeByName(r.BID, name)
	if rsp.RSPID == 0 {
		fmt.Printf("%s: line %d - could not find a RentableSpecialtyType named %s in Business %d\n", funcname, lineno, name, r.BID)
		return
	}
	a.RSPID = rsp.RSPID

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	a.DtStart, a.DtStop, err = readTwoDates(sa[2:], funcname, lineno)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	err = InsertRentableSpecialtyRef(&a)
	if err != nil {
		fmt.Printf("%s: line %d - error inserting assessment: %v\n", funcname, lineno, err)
	}
}

// LoadRentableSpecialtyRefsCSV loads a csv file with a chart of accounts and creates Ledger markers for each
func LoadRentableSpecialtyRefsCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentableSpecialtyRefsCSV(t[i], i+1)
	}
}
