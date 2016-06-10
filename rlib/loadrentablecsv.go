package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

// RentableSpecialty is the structure for attributes of a Rentable specialty

// CSV file format:
//   0  1     2               3                                 4
//                            "S1,Strt1,Stp1;S2,Strt2,Stp2...", “A2,1/10/16,6/1/16;B2,6/1/16,”
// BUD, Name, AssignmentTime, RentableStatus,                   RentableTypeRef
// REX, 101,  1,              "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 102,  1,              "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 103,  1,              "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 104,  1,              "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 105,  1,              "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 106,  1,              "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"

// CreateRentables reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRentables(sa []string, lineno int) {
	funcname := "CreateRentables"
	var err error
	var r Rentable

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

	//-------------------------------------------------------------------
	// Make sure the Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			Ulog("%s: line %d - Business with bud %s does not exist\n", funcname, lineno, des)
			return
		}
		r.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// The name must be unique. Make sure we don't have any other Rentable
	// with this name...
	//-------------------------------------------------------------------
	r.Name = strings.TrimSpace(sa[1])
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
		i, err := strconv.Atoi(sa[2])
		if err != nil || i < 0 || i > 2 {
			fmt.Printf("%s: lineno %d - invalid AssignmentTime number: %s\n", funcname, lineno, sa[2])
			return
		}
		r.AssignmentTime = int64(i)
	}

	// //-------------------------------------------------------------------
	// // parse out the DefaultOccupancy value
	// //   any accrual frequency is valid
	// //-------------------------------------------------------------------
	// if len(sa[3]) > 0 {
	// 	i, err := strconv.Atoi(sa[3])
	// 	if err != nil || !IsValidAccrual(int64(i)) {
	// 		fmt.Printf("%s: lineno %d - invalid DefaultOccupancy value: %s\n", funcname, lineno, sa[3])
	// 		return
	// 	}
	// 	r.RentalPeriodDefault = int64(i)
	// }

	// //-------------------------------------------------------------------
	// // parse out the Occupancy value
	// // any accrual frequency is valid
	// //-------------------------------------------------------------------
	// if len(sa[4]) > 0 {
	// 	i, err := strconv.Atoi(sa[4])
	// 	if err != nil || !IsValidAccrual(int64(i)) {
	// 		fmt.Printf("%s: lineno %d - invalid Occupancy value: %s\n", funcname, lineno, sa[4])
	// 		return
	// 	}
	// 	r.RentCycle = int64(i)
	// }

	//-----------------------------------------------------------------------------------
	// PARSE THE STATUS 3-TUPLEs
	// "S1,Strt1,Stp1;S2,Strt2,Stp2 ..."
	//-----------------------------------------------------------------------------------
	if 0 == len(strings.TrimSpace(sa[3])) {
		fmt.Printf("%s: lineno %d - RentableStatus value is required.\n",
			funcname, lineno)
		return
	}
	var m []RentableStatus          // keep every RentableStatus we find in an array
	st := strings.Split(sa[3], ";") // split it on Status 3-tuple separator (;)
	for i := 0; i < len(st); i++ {  //spin through the 3-tuples
		ss := strings.Split(st[i], ",")
		if len(ss) != 3 {
			fmt.Printf("%s: lineno %d - invalid Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
				funcname, lineno, len(ss), ss)
			return
		}

		var rs RentableStatus // struct for the data in this 3-tuple
		ix, err := strconv.Atoi(ss[0])
		if err != nil || ix < RENTABLESTATUSONLINE || ix > RENTABLESTATUSLAST {
			fmt.Printf("%s: lineno %d - invalid Status value: %s.  Must be in the range %d to %d\n",
				funcname, lineno, ss[0], RENTABLESTATUSONLINE, RENTABLESTATUSLAST)
			return
		}
		rs.Status = int64(ix)

		rs.DtStart, err = StringToDate(ss[1]) // required field
		if err != nil {
			fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, ss[1])
			return
		}

		end := "1/1/9999"
		if len(ss) > 2 { //optional field -- MAYBE, if not present assume year 9999
			if i+1 != len(st) {
				fmt.Printf("%s: line %d - unspecified stop date is only allowed on the last RentableStatus in the list\n", funcname, lineno)
				return
			}
			if len(strings.TrimSpace(ss[2])) > 0 {
				end = ss[2]
			}
		}
		rs.DtStop, err = StringToDate(end)
		if err != nil {
			fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, ss[2])
			return
		}
		m = append(m, rs) // add this struct to the list
	}
	if len(m) == 0 {
		fmt.Printf("%s: lineno %d - RentableStatus value is required.\n",
			funcname, lineno)
		return
	}

	//-----------------------------------------------------------------------------------
	// PARSE THE RTID 3-TUPLEs
	// "RTname1,startDate1,stopDate1;RTname2,startDate2,stopDate2;..."
	//-----------------------------------------------------------------------------------
	if 0 == len(strings.TrimSpace(sa[4])) {
		fmt.Printf("%s: lineno %d - Rentable RTID Ref value is required.\n",
			funcname, lineno)
		return
	}
	var n []RentableRTID
	st = strings.Split(sa[4], ";") // split on RTID 3-tuple seperator (;)
	for i := 0; i < len(st); i++ { // spin through the 3-tuples
		ss := strings.Split(st[i], ",") // separate the 3 parts
		if len(ss) != 3 {
			fmt.Printf("%s: lineno %d - invalid RTID syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
				funcname, lineno, len(ss), ss)
			return
		}

		var rt RentableRTID                                                     // struct for the data in this 3-tuple
		rstruct, err := GetRentableTypeByStyle(strings.TrimSpace(ss[0]), r.BID) // find the RentableType being referenced
		if err != nil {
			fmt.Printf("%s: lineno %d - Could not load rentable type with style name: %s  -- error = %s\n",
				funcname, lineno, ss[0], err.Error())
			return
		}
		rt.RTID = rstruct.RTID

		rt.DtStart, err = StringToDate(ss[1]) // required field
		if err != nil {
			fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, ss[1])
			return
		}

		end := "1/1/9999"
		if len(ss) > 2 { //optional field, if not present assume year 9999
			if i+1 != len(st) {
				fmt.Printf("%s: line %d - unspecified stop date is only allowed on the last Rentable RTID Ref in the list\n", funcname, lineno)
				return
			}
			if len(strings.TrimSpace(ss[2])) > 0 {
				end = ss[2]
			}
		}
		rt.DtStop, err = StringToDate(end)
		if err != nil {
			fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, ss[2])
			return
		}
		n = append(n, rt) // add this struct to the list
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and its sub-records and we're done
	//-------------------------------------------------------------------
	rid, err := InsertRentable(&r)
	if nil != err {
		fmt.Printf("%s: lineno %d - error inserting Rentable = %v\n", funcname, lineno, err)
	}
	if rid > 0 {
		for i := 0; i < len(m); i++ {
			m[i].RID = rid
			err := InsertRentableStatus(&m[i])
			if err != nil {
				fmt.Printf("%s: lineno %d - error saving RentableStatus: %s\n", funcname, lineno, err.Error())
			}
		}
		for i := 0; i < len(n); i++ {
			n[i].RID = rid
			err := InsertRentableRTID(&n[i])
			if err != nil {
				fmt.Printf("%s: lineno %d - error saving RentableStatus: %s\n", funcname, lineno, err.Error())
			}
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
