package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

// RentableSpecialty is the structure for attributes of a rentable specialty

// CSV file format:
//   0           1     2         3      4       5      7
// Designation,Style,Name,Assignment,Report,DefaultOcc,Occ
//	REH,1,GM, "101",1,1,2,2
//	REH,1,FS, "102",1,1,2,2
//	REH,1,SBL,"103",1,1,2,2
//	REH,1,KDS,"104",1,1,2,2
//	REH,1,GM, "105",1,1,2,2
//	REH,1,FS, "106",1,1,2,2
//	REH,1,SBL,"107",1,1,2,2

// CreateRentables reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRentables(sa []string) {
	var err error
	var r Rentable
	var rt RentableType
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
			Ulog("CreateRentables: business with designation %s does not exist\n", des)
			return
		}
		r.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Make sure the RentableType is in the database
	//-------------------------------------------------------------------
	if len(sa[1]) > 0 {
		i, err := strconv.Atoi(strings.TrimSpace(sa[1]))
		if err != nil {
			fmt.Printf("Could not find RentableType %s\n", sa[1])
		}
		err = GetRentableType(int64(i), &rt)
		if err != nil {
			Ulog("CreateRentables: could not load rentable type %s, err = %s\n", sa[1], err.Error())
			return
		}
	}

	//-------------------------------------------------------------------
	// Set the rentable type
	//-------------------------------------------------------------------
	style := strings.TrimSpace(sa[2])
	if len(style) > 0 {
		rs, _ := GetRentableTypeByStyle(style, r.BID)
		if rs.RTID == 0 {
			Ulog("CreateRentables: rentable type %s does not exist in business %s\n", style, des)
			return
		}
		r.RTID = rs.RTID
	}

	//-------------------------------------------------------------------
	// The name must be unique. Make sure we don't have any other rentable
	// with this name...
	//-------------------------------------------------------------------
	r.Name = strings.TrimSpace(sa[3])
	r1, err := GetRentableByName(r.Name, r.BID)
	if err != nil {
		s := err.Error()
		if !strings.Contains(s, "no rows") {
			fmt.Printf("CreateRentables: error with GetRentableByName: %s\n", err.Error())
			return
		}
	}
	if r1.RID > 0 {
		fmt.Printf("Rentable with name \"%s\" already exists. Skipping. \n", r.Name)
		return
	}

	//-------------------------------------------------------------------
	// parse out the Assignment value
	// Unknown = 0, Pre-assign = 1, assign at occupy commencement = 2
	//-------------------------------------------------------------------
	if len(sa[4]) > 0 {
		i, err := strconv.Atoi(sa[4])
		if err != nil || i < 0 || i > 2 {
			fmt.Printf("CreateRentables: invalid Assignment number: %s\n", sa[4])
			return
		}
		r.Assignment = int64(i)
	}

	//-------------------------------------------------------------------
	// parse out the Report value
	// 1 = apply to rentroll, 0 = skip on rentroll
	//-------------------------------------------------------------------
	if len(sa[5]) > 0 {
		i, err := strconv.Atoi(sa[5])
		if err != nil || i < 0 || i > 1 {
			fmt.Printf("CreateRentables: invalid Report number: %s\n", sa[5])
			return
		}
		r.Report = int64(i)
	}

	//-------------------------------------------------------------------
	// parse out the DefaultOccupancy value
	// 0 =unset, 1 = short term, 2=longterm
	//-------------------------------------------------------------------
	if len(sa[6]) > 0 {
		i, err := strconv.Atoi(sa[6])
		if err != nil || i < 0 || i > 2 {
			fmt.Printf("CreateRentables: invalid DefaultOccupancy number: %s\n", sa[6])
			return
		}
		r.DefaultOccType = int64(i)
	}

	//-------------------------------------------------------------------
	// parse out the Occupancy value
	// 0 =unset, 1 = short term, 2=longterm
	//-------------------------------------------------------------------
	if len(sa[7]) > 0 {
		i, err := strconv.Atoi(sa[7])
		if err != nil || i < 0 || i > 2 {
			fmt.Printf("CreateRentables: invalid Occupancy number: %s\n", sa[7])
			return
		}
		r.OccType = int64(i)
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	_, err = InsertRentable(&r)
	if nil != err {
		fmt.Printf("CreateRentables: error inserting Rentable = %v\n", err)
	}
}

// LoadRentablesCSV loads a csv file with rental specialty types and processes each one
func LoadRentablesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentables(t[i])
	}
}
