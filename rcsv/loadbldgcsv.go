package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

//  CSV file format:
// 0           1      2       3        4    5     6          7
// Designation,BldgNo,Address,Address2,City,State,PostalCode,Country
// REH,1,"2001 Creaking Oak Drive","","Springfield","MO","65803","USA"

// CreateBuilding reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateBuilding(sa []string, lineno int) {
	funcname := "CreateBuilding"
	var b rlib.Building
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 8
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}
	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d - rlib.Business with designation %s does net exist\n", funcname, lineno, des)
			return
		}
		b.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// parse out the rlib.Building number
	//-------------------------------------------------------------------
	if len(sa[1]) > 0 {
		i, err := strconv.Atoi(sa[1])
		if err != nil || i < 0 {
			fmt.Printf("%s: line %d - invalid rlib.Building number: %s\n", funcname, lineno, sa[1])
			return
		}
		b.BLDGID = int64(i)
	}

	b.Address = strings.TrimSpace(sa[2])
	b.Address2 = strings.TrimSpace(sa[3])
	b.City = strings.TrimSpace(sa[4])
	b.State = strings.TrimSpace(sa[5])
	b.PostalCode = strings.TrimSpace(sa[6])
	b.Country = strings.TrimSpace(sa[7])

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	_, err := rlib.InsertBuildingWithID(&b)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.Building = %v\n", funcname, lineno, err)
	}
}

// LoadBuildingCSV loads a csv file with rental specialty types and processes each one
func LoadBuildingCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateBuilding(t[i], i+1)
	}
}
