package rlib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// RentableSpecialty is the structure for attributes of a rentable specialty

// CSV file format:
//   0           1     2         3       4      5   6   7     8       9
// Designation,Style,Name,AssignmentTime,DefaultOcc,Occ,Status,DtStart,DtStop
// REX,        GM,   101,    1,        2,       2,    0,0,    ,"1/1/14",
// REX,        FS,   102,    1,        2,       2,    0
// REX,        SBL,  103,    1,        2,       2,    0
// REX,        KDS,  104,    1,        2,       2,    0
// REX,        GM,   105,    1,        2,       2,    0
// REX,        FS,   106,    1,        2,       2,    0

// CreateRentables reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRentables(sa []string, lineno int) {
	funcname := "CreateRentables"
	var err error
	var r Rentable
	var status int64
	var DtStart, DtStop time.Time

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
			Ulog("%s: line %d - business with designation %s does not exist\n", funcname, lineno, des)
			return
		}
		r.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Set the rentable type
	//-------------------------------------------------------------------
	style := strings.TrimSpace(sa[1])
	if len(style) > 0 {
		rs, _ := GetRentableTypeByStyle(style, r.BID)
		if rs.RTID == 0 {
			Ulog("%s: line %d - rentable type %s does not exist in business %s\n", funcname, lineno, style, des)
			return
		}
		r.RTID = rs.RTID
	}

	//-------------------------------------------------------------------
	// The name must be unique. Make sure we don't have any other rentable
	// with this name...
	//-------------------------------------------------------------------
	r.Name = strings.TrimSpace(sa[2])
	r1, err := GetRentableByName(r.Name, r.BID)
	if err != nil {
		s := err.Error()
		if !strings.Contains(s, "no rows") {
			fmt.Printf("%s: lineno %d - error with GetRentableByName: %s\n", funcname, lineno, err.Error())
			return
		}
	}
	if r1.RID > 0 {
		fmt.Printf("%s: lineno %d - Rentable with name \"%s\" already exists. Skipping. \n", funcname, lineno, r.Name)
		return
	}

	//-------------------------------------------------------------------
	// parse out the AssignmentTime value
	// Unknown = 0, Pre-assign = 1, assign at occupy commencement = 2
	//-------------------------------------------------------------------
	if len(sa[3]) > 0 {
		i, err := strconv.Atoi(sa[3])
		if err != nil || i < 0 || i > 2 {
			fmt.Printf("%s: lineno %d - invalid AssignmentTime number: %s\n", funcname, lineno, sa[3])
			return
		}
		r.AssignmentTime = int64(i)
	}

	//-------------------------------------------------------------------
	// parse out the DefaultOccupancy value
	//   any accrual frequency is valid
	//-------------------------------------------------------------------
	if len(sa[4]) > 0 {
		i, err := strconv.Atoi(sa[4])
		if err != nil || !IsValidAccrual(int64(i)) {
			fmt.Printf("%s: lineno %d - invalid DefaultOccupancy value: %s\n", funcname, lineno, sa[4])
			return
		}
		r.RentalPeriodDefault = int64(i)
	}

	//-------------------------------------------------------------------
	// parse out the Occupancy value
	// any accrual frequency is valid
	//-------------------------------------------------------------------
	if len(sa[5]) > 0 {
		i, err := strconv.Atoi(sa[5])
		if err != nil || !IsValidAccrual(int64(i)) {
			fmt.Printf("%s: lineno %d - invalid Occupancy value: %s\n", funcname, lineno, sa[5])
			return
		}
		r.RentalPeriod = int64(i)
	}

	//-------------------------------------------------------------------
	// parse out the Status value
	// 0 = normal, 1 = online, 2 = administrative unit, 3 = owner occupied, 4 = offline
	//-------------------------------------------------------------------
	if len(sa[6]) > 0 {
		i, err := strconv.Atoi(sa[6])
		if err != nil || i < RENTABLESTATUSONLINE || i > RENTABLESTATUSLAST {
			fmt.Printf("%s: lineno %d - invalid Status value: %s.  Must be in the range %d to %d\n",
				funcname, lineno, sa[6], RENTABLESTATUSONLINE, RENTABLESTATUSLAST)
			return
		}
		status = int64(i)

		DtStart, err = StringToDate(sa[7]) // required field
		if err != nil {
			fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[7])
			return
		}

		end := "1/1/9999"
		if len(sa) > 8 { //optional field, if not present assume year 9999
			if len(strings.TrimSpace(sa[8])) > 0 {
				end = sa[8]
			}
		}
		DtStop, err = StringToDate(end)
		if err != nil {
			fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[8])
			return
		}
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	rid, err := InsertRentable(&r)
	if nil != err {
		fmt.Printf("%s: lineno %d - error inserting Rentable = %v\n", funcname, lineno, err)
	}
	if rid > 0 {
		var rs RentableStatus
		rs.RID = rid
		rs.DtStart = DtStart
		rs.DtStop = DtStop
		rs.Status = status
		err := InsertRentableStatus(&rs)
		if err != nil {
			fmt.Printf("%s: lineno %d - error saving RentableStatus: %s\n", funcname, lineno, err.Error())
		}
	}
}

// LoadRentablesCSV loads a csv file with rental specialty types and processes each one
func LoadRentablesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentables(t[i], i+1)
	}
}
