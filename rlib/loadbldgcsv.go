package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

// RentableSpecialty is the structure for attributes of a rentable specialty

//  CSV file format:
// BldgNo,Address,Address2,City,State,PostalCode,Country
// 1,"2001 Creaking Oak Drive","","Springfield","MO","65803","USA"
// Designation, Name,            Fee,   Description

// CreateBuilding reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateBuilding(sa []string) {
	var b Building
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Make sure the business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			Ulog("CreateBuilding: business with designation %s does net exist\n", des)
			return
		}
		b.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// parse out the building number
	//-------------------------------------------------------------------
	if len(sa[1]) > 0 {
		i, err := strconv.Atoi(sa[1])
		if err != nil || i < 0 {
			fmt.Printf("CreateBuilding: invalid building number: %s\n", sa[1])
		} else {
			b.BLDGID = int64(i)
		}
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
	_, err := InsertBuildingWithID(&b)
	if nil != err {
		fmt.Printf("CreateBuilding: error inserting Building = %v\n", err)
	}
}

// LoadBuildingCSV loads a csv file with rental specialty types and processes each one
func LoadBuildingCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateBuilding(t[i])
	}
}
