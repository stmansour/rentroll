package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// CSV file format:
//   0  1     2               3                       4                                 5
//                            "usr1;usr2;..usrN"      "S1,Strt1,Stp1;S2,Strt2,Stp2...", “A2,1/10/16,6/1/16;B2,6/1/16,”
// BUD, Name, AssignmentTime, RentableUsers,          RentableStatus,                   RentableTypeRef
// REX, 101,  1,              "bill@x.com;sue@x.com"  "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 102,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 103,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 104,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 105,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 106,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"

// readTwoDates assumes that a date string is in ss[1] and ss[2].  It will parse and return the dates
// along with any error it finds.
func readTwoDates(s1, s2 string, funcname string, lineno int) (time.Time, time.Time, error) {
	var DtStart, DtStop time.Time
	var err error
	DtStart, err = rlib.StringToDate(s1) // required field
	if err != nil {
		err = fmt.Errorf("%s: line %d - invalid start date:  %s\n", funcname, lineno, s1)
		return DtStart, DtStop, err
	}

	end := "1/1/9999"
	if len(s2) > 0 { //optional field -- MAYBE, if not present assume year 9999
		if len(strings.TrimSpace(s2)) > 0 {
			end = s2
		}
	}
	DtStop, err = rlib.StringToDate(end)
	if err != nil {
		err = fmt.Errorf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, s2)
	}
	return DtStart, DtStop, err
}

// CreateRentables reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRentables(sa []string, lineno int) int {
	funcname := "CreateRentables"
	var err error
	var r rlib.Rentable

	const (
		BUD             = 0
		Name            = iota
		AssignmentTime  = iota
		RUserSpec       = iota
		RentableStatus  = iota
		RentableTypeRef = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"AssignmentTime", AssignmentTime},
		{"RUserSpec", RUserSpec},
		{"RentableStatus", RentableStatus},
		{"RentableTypeRef", RentableTypeRef},
	}

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return 1
	}
	if lineno == 1 {
		return 0
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d - Business with bud %s does not exist\n", funcname, lineno, des)
			return CsvErrorSensitivity
		}
		r.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// The name must be unique. Make sure we don't have any other rlib.Rentable
	// with this name...
	//-------------------------------------------------------------------
	r.Name = strings.TrimSpace(sa[1])
	r1, err := rlib.GetRentableByName(r.Name, r.BID)
	if err != nil {
		s := err.Error()
		if !strings.Contains(s, "no rows") {
			fmt.Printf("%s: lineno %d - error with rlib.GetRentableByName: %s\n", funcname, lineno, err.Error())
			return CsvErrorSensitivity
		}
	}
	if r1.RID > 0 {
		fmt.Printf("%s: lineno %d - Rentable with name \"%s\" already exists. Skipping. \n", funcname, lineno, r.Name)
		return CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// parse out the AssignmentTime value
	// Unknown = 0, Pre-assign = 1, assign at occupy commencement = 2
	//-------------------------------------------------------------------
	if len(sa[2]) > 0 {
		i, err := strconv.Atoi(sa[2])
		if err != nil || i < 0 || i > 2 {
			fmt.Printf("%s: lineno %d - invalid AssignmentTime number: %s\n", funcname, lineno, sa[2])
			return CsvErrorSensitivity
		}
		r.AssignmentTime = int64(i)
	}

	//-----------------------------------------------------------------------------------
	// USER 3-TUPLEs
	// "user1,dtstart1,dtstop1;user2,dtstart2,dtstop2;..."
	// example:
	// "ednak@springfield.com,1/1/2013,11/9/2015;homerj@springfield.com,11/20/2015,;marge@springfield.com,11/20/2015,"
	//-----------------------------------------------------------------------------------
	var rul []rlib.RentableUser // keep every rlib.RentableUser we find in an array
	if 0 < len(strings.TrimSpace(sa[3])) {
		st := strings.Split(sa[3], ";") // split it on Status 3-tuple separator (;)
		for i := 0; i < len(st); i++ {  //spin through the 3-tuples
			ss := strings.Split(st[i], ",")
			if len(ss) != 3 {
				fmt.Printf("%s: lineno %d - invalid Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
					funcname, lineno, len(ss), ss)
				return CsvErrorSensitivity
			}

			var ru rlib.RentableUser // struct for the data in this 3-tuple
			name := strings.TrimSpace(ss[0])
			// t := rlib.GetTransactantByPhoneOrEmail(name)
			n, err := CSVLoaderTransactantList(name)
			if n[0].TCID == 0 {
				rerr := fmt.Sprintf("%s: line %d - could not find Transactant with contact information %s\n", funcname, lineno, name)
				fmt.Printf("%s", rerr)
				return CsvErrorSensitivity
			}
			ru.TCID = n[0].TCID

			ru.DtStart, ru.DtStop, err = readTwoDates(ss[1], ss[2], funcname, lineno)
			if err != nil {
				fmt.Printf("%s", err.Error())
				return CsvErrorSensitivity
			}
			rul = append(rul, ru) // add this struct to the list
		}
	}

	//-----------------------------------------------------------------------------------
	// STATUS 3-TUPLEs
	// "S1,Strt1,Stp1;S2,Strt2,Stp2 ..."
	//-----------------------------------------------------------------------------------
	if 0 == len(strings.TrimSpace(sa[4])) {
		fmt.Printf("%s: lineno %d - rlib.RentableStatus value is required.\n",
			funcname, lineno)
		return CsvErrorSensitivity
	}
	var m []rlib.RentableStatus     // keep every rlib.RentableStatus we find in an array
	st := strings.Split(sa[4], ";") // split it on Status 3-tuple separator (;)
	for i := 0; i < len(st); i++ {  //spin through the 3-tuples
		ss := strings.Split(st[i], ",")
		if len(ss) != 3 {
			fmt.Printf("%s: lineno %d - invalid Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
				funcname, lineno, len(ss), ss)
			return CsvErrorSensitivity
		}

		var rs rlib.RentableStatus // struct for the data in this 3-tuple
		ix, err := strconv.Atoi(ss[0])
		if err != nil || ix < rlib.RENTABLESTATUSONLINE || ix > rlib.RENTABLESTATUSLAST {
			fmt.Printf("%s: lineno %d - invalid Status value: %s.  Must be in the range %d to %d\n",
				funcname, lineno, ss[0], rlib.RENTABLESTATUSONLINE, rlib.RENTABLESTATUSLAST)
			return CsvErrorSensitivity
		}
		rs.Status = int64(ix)

		rs.DtStart, rs.DtStop, err = readTwoDates(ss[1], ss[2], funcname, lineno)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return CsvErrorSensitivity
		}
		m = append(m, rs) // add this struct to the list
	}
	if len(m) == 0 {
		fmt.Printf("%s: lineno %d - rlib.RentableStatus value is required.\n",
			funcname, lineno)
		return CsvErrorSensitivity
	}

	//-----------------------------------------------------------------------------------
	// RTID 3-TUPLEs
	// "RTname1,Amount,startDate1,stopDate1;RTname2,startDate2,stopDate2;..."
	//-----------------------------------------------------------------------------------
	if 0 == len(strings.TrimSpace(sa[5])) {
		fmt.Printf("%s: lineno %d - rlib.Rentable RTID Ref value is required.\n",
			funcname, lineno)
		return CsvErrorSensitivity
	}
	var n []rlib.RentableTypeRef
	st = strings.Split(sa[5], ";") // split on RTID 3-tuple seperator (;)
	for i := 0; i < len(st); i++ { // spin through the 3-tuples
		ss := strings.Split(st[i], ",") // separate the 3 parts
		if len(ss) != 3 {
			fmt.Printf("%s: lineno %d - invalid RTID syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
				funcname, lineno, len(ss), ss)
			return CsvErrorSensitivity
		}

		var rt rlib.RentableTypeRef                                                  // struct for the data in this 3-tuple
		rstruct, err := rlib.GetRentableTypeByStyle(strings.TrimSpace(ss[0]), r.BID) // find the rlib.RentableType being referenced
		if err != nil {
			fmt.Printf("%s: lineno %d - Could not load rentable type with style name: %s  -- error = %s\n",
				funcname, lineno, ss[0], err.Error())
			return CsvErrorSensitivity
		}
		rt.RTID = rstruct.RTID

		rt.DtStart, rt.DtStop, err = readTwoDates(ss[1], ss[2], funcname, lineno)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return CsvErrorSensitivity
		}
		n = append(n, rt) // add this struct to the list
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and its sub-records and we're done
	//-------------------------------------------------------------------
	rid, err := rlib.InsertRentable(&r)
	if nil != err {
		fmt.Printf("%s: lineno %d - error inserting rlib.Rentable = %v\n", funcname, lineno, err)
	}
	if rid > 0 {
		for i := 0; i < len(rul); i++ {
			rul[i].RID = rid
			rlib.InsertRentableUser(&rul[i])
		}
		for i := 0; i < len(m); i++ {
			m[i].RID = rid
			err := rlib.InsertRentableStatus(&m[i])
			if err != nil {
				fmt.Printf("%s: lineno %d - error saving rlib.RentableStatus: %s\n", funcname, lineno, err.Error())
			}
		}
		for i := 0; i < len(n); i++ {
			n[i].RID = rid
			err := rlib.InsertRentableTypeRef(&n[i])
			if err != nil {
				fmt.Printf("%s: lineno %d - error saving rlib.RentableStatus: %s\n", funcname, lineno, err.Error())
			}
		}
	}
	return 0
}

// LoadRentablesCSV loads a csv file with rental specialty types and processes each one
func LoadRentablesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if CreateRentables(t[i], i+1) > 0 {
			return
		}
	}
}
