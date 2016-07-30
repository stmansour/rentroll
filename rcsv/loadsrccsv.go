package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// CSV FIELDS FOR THIS MODULE
//    0    1     2
//    BUD, Name, Industry
//    REX, FAA,  Aviation
//    REX, IRS,  Excessive Rules

// CreateSourceCSV reads an assessment type string array and creates a database record for the assessment type
func CreateSourceCSV(sa []string, lineno int) {
	funcname := "CreateSourceCSV"
	var a rlib.DemandSource
	var err error

	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "bud" {
		return // this is just the column heading
	}

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 3
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Business
	//-------------------------------------------------------------------
	var b rlib.Business
	if len(des) > 0 {
		b, _ = rlib.GetBusinessByDesignation(des)
		if b.BID < 1 {
			rlib.Ulog("CreateRentalSpecialtyType: rlib.Business named %s not found\n", sa[0])
			return
		}
	}
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Name
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[1])
	if len(s) > 0 {
		var src rlib.DemandSource
		rlib.GetDemandSourceByName(b.BID, s, &src)
		if len(src.Name) > 0 {
			fmt.Printf("%s: line %d - DemandSource named %s already exists.\n", funcname, lineno, s)
			return
		}
	}
	a.Name = s

	//-------------------------------------------------------------------
	// Industry
	//-------------------------------------------------------------------
	a.Industry = strings.TrimSpace(sa[2])

	_, err = rlib.InsertDemandSource(&a)
	if err != nil {
		fmt.Printf("%s: line %d - error inserting DemandSource: %v\n", funcname, lineno, err)
	}
}

// LoadSourcesCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadSourcesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateSourceCSV(t[i], i+1)
	}
}
