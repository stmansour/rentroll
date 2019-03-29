package rcsv

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// CSV file format:
//   0  1     2               3                       4                                 5
//                            "usr1;usr2;..usrN"      "S1,Strt1,Stp1;S2,Strt2,Stp2...", “A2,1/10/16,6/1/16;B2,6/1/16,”
// BUD, Name, AssignmentTime, RentableUsers,          RentableUseStatus,                   RentableTypeRef
// REX, 101,  1,              "bill@x.com;sue@x.com"  "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 102,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 103,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 104,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 105,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"
// REX, 106,  1,                                      "1,1/1/14,6/15/16;2,6/15/16,",    "A2,1/1/14,6/1/16;B2,6/1/16,"

// readTwoDates assumes that a date string is in ss[1] and ss[2].  It will parse and return the dates
// along with any error it finds.
func readTwoDates(s1, s2 string, funcname string, lineno int, col string) (time.Time, time.Time, error) {
	var DtStart, DtStop time.Time
	var err error
	DtStart, err = rlib.StringToDate(s1) // required field
	if err != nil {
		err = fmt.Errorf("%s: line %d   column: %s - invalid start date:  %s", funcname, lineno, col, s1)
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
		err = fmt.Errorf("%s: line %d   column: %s - invalid stop date:  %s", funcname, lineno, col, s2)
	}
	return DtStart, DtStop, err
}

// CreateRentables reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRentables(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateRentables"
	var (
		err error
		r   rlib.Rentable
	)

	const (
		BUD               = 0
		Name              = iota
		AssignmentTime    = iota
		RUserSpec         = iota
		RentableUseStatus = iota
		RentableTypeRef   = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"AssignmentTime", AssignmentTime},
		{"RUserSpec", RUserSpec},
		{"RentableUseStatus", RentableUseStatus},
		{"RentableTypeRef", RentableTypeRef},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[BUD]))
	if len(des) > 0 {
		b1, err := rlib.GetBusinessByDesignation(ctx, des)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, error while getting business by designation(%s): %s", funcname, lineno, des, err.Error())
		}
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Business with bud %s does not exist", funcname, lineno, des)
		}
		r.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// The name must be unique. Make sure we don't have any other rlib.Rentable
	// with this name...
	//-------------------------------------------------------------------
	r.RentableName = strings.TrimSpace(sa[Name])
	r1, err := rlib.GetRentableByName(ctx, r.RentableName, r.BID)
	if err != nil {
		s := err.Error()
		if !strings.Contains(s, "no rows") {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error with rlib.GetRentableByName: %s", funcname, lineno, err.Error())
		}
	}
	if r1.RID > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s:: Rentable with name \"%s\" already exists. Skipping. ", funcname, lineno, DupRentable, r.RentableName)
	}

	//-------------------------------------------------------------------
	// parse out the AssignmentTime value
	// Unknown = 0, Pre-assign = 1, assign at occupy commencement = 2
	//-------------------------------------------------------------------
	if len(sa[AssignmentTime]) > 0 {
		i, err := strconv.Atoi(sa[AssignmentTime])
		if err != nil || i < 0 || i > 2 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid AssignmentTime number: %s", funcname, lineno, sa[AssignmentTime])
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
	if 0 < len(strings.TrimSpace(sa[RUserSpec])) {
		st := strings.Split(sa[RUserSpec], ";") // split it on Status 3-tuple separator (;)
		for i := 0; i < len(st); i++ {          //spin through the 3-tuples
			ss := strings.Split(st[i], ",")
			if len(ss) != 3 {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid User Specification. Each semi-colon separated field must have 3 values. Found %d in \"%s\"",
					funcname, lineno, len(ss), ss)
			}

			var ru rlib.RentableUser // struct for the data in this 3-tuple
			name := strings.TrimSpace(ss[0])
			n, err := CSVLoaderTransactantList(ctx, r.BID, name)
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error Loading transactant list: %s", funcname, lineno, err.Error())
			}
			if len(n) == 0 || n[0].TCID == 0 {
				rerr := fmt.Sprintf("%s: line %d - could not find Transactant with contact information %s\n", funcname, lineno, name)
				return CsvErrorSensitivity, fmt.Errorf("%s", rerr)
			}
			ru.TCID = n[0].TCID

			ru.DtStart, ru.DtStop, err = readTwoDates(ss[1], ss[2], funcname, lineno, "RUserSpec")
			if err != nil {
				return CsvErrorSensitivity, err
			}
			rul = append(rul, ru) // add this struct to the list
		}
	}

	//-----------------------------------------------------------------------------------
	// STATUS 3-TUPLEs
	// "S1,Strt1,Stp1;S2,Strt2,Stp2 ..."
	//-----------------------------------------------------------------------------------
	if 0 == len(strings.TrimSpace(sa[RentableUseStatus])) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - rlib.RentableUseStatus value is required",
			funcname, lineno)
	}
	var m []rlib.RentableUseStatus                  // keep every rlib.RentableUseStatus we find in an array
	st := strings.Split(sa[RentableUseStatus], ";") // split it on Status 3-tuple separator (;)
	for i := 0; i < len(st); i++ {                  //spin through the 3-tuples
		ss := strings.Split(st[i], ",")
		if len(ss) != 3 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Rentable Status. Each semi-colon separated field must have 3 values. Found %d in \"%s\"",
				funcname, lineno, len(ss), ss)
		}

		var rst rlib.RentableUseStatus // struct for the data in this 3-tuple
		ix, err := strconv.Atoi(ss[0])
		if err != nil || ix < rlib.USESTATUSready || ix > rlib.USESTATUSLAST {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Status value: %s.  Must be in the range %d to %d",
				funcname, lineno, ss[0], rlib.USESTATUSready, rlib.USESTATUSLAST)
		}
		rst.UseStatus = int64(ix)

		rst.DtStart, rst.DtStop, err = readTwoDates(ss[1], ss[2], funcname, lineno, "RentableUseStatus")
		if err != nil {
			return CsvErrorSensitivity, err
		}
		m = append(m, rst) // add this struct to the list
	}
	if len(m) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - RentableUseStatus value is required",
			funcname, lineno)
	}

	//-----------------------------------------------------------------------------------
	// RTID 3-TUPLEs
	// "RTname1,Amount,startDate1,stopDate1;RTname2,startDate2,stopDate2;..."
	//-----------------------------------------------------------------------------------
	if 0 == len(strings.TrimSpace(sa[RentableTypeRef])) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Rentable RTID Ref value is required",
			funcname, lineno)
	}
	var n []rlib.RentableTypeRef
	st = strings.Split(sa[RentableTypeRef], ";") // split on RTID 3-tuple separator (;)
	for i := 0; i < len(st); i++ {               // spin through the 3-tuples
		ss := strings.Split(st[i], ",") // separate the 3 parts
		if len(ss) != 3 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid RTID syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"",
				funcname, lineno, len(ss), ss)
		}

		var rt rlib.RentableTypeRef                                                       // struct for the data in this 3-tuple
		rstruct, err := rlib.GetRentableTypeByStyle(ctx, strings.TrimSpace(ss[0]), r.BID) // find the rlib.RentableType being referenced
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not load rentable type with style name: %s  -- error = %s",
				funcname, lineno, ss[0], err.Error())
		}
		if rstruct.RTID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not load rentable type with style name: %s",
				funcname, lineno, ss[0])
		}
		rt.RTID = rstruct.RTID

		rt.DtStart, rt.DtStop, err = readTwoDates(ss[1], ss[2], funcname, lineno, "RentableTypeRef")
		if err != nil {
			return CsvErrorSensitivity, err
		}
		n = append(n, rt) // add this struct to the list
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and its sub-records and we're done
	//-------------------------------------------------------------------
	rid, err := rlib.InsertRentable(ctx, &r)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting rlib.Rentable = %v", funcname, lineno, err)
	}
	if rid > 0 {
		for i := 0; i < len(rul); i++ {
			rul[i].RID = rid
			rul[i].BID = r.BID
			_, err := rlib.InsertRentableUser(ctx, &rul[i])
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error saving rlib.RentableUser: %s", funcname, lineno, err.Error())
			}
		}
		for i := 0; i < len(m); i++ {
			m[i].RID = rid
			m[i].BID = r.BID
			_, err := rlib.InsertRentableUseStatus(ctx, &m[i])
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error saving rlib.RentableUseStatus: %s", funcname, lineno, err.Error())
			}
			var ls = rlib.RentableLeaseStatus{
				BID:         m[i].BID,
				RID:         m[i].RID,
				LeaseStatus: rlib.LEASESTATUSnotleased,
				DtStart:     m[i].DtStart,
				DtStop:      m[i].DtStop,
			}
			if _, err = rlib.InsertRentableLeaseStatus(ctx, &ls); err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error saving rlib.RentableLeaseStatus: %s", funcname, lineno, err.Error())
			}
			var ru = rlib.RentableUseType{
				BID:     m[i].BID,
				RID:     m[i].RID,
				UseType: rlib.USETYPEstandard,
				DtStart: m[i].DtStart,
				DtStop:  m[i].DtStop,
			}
			if _, err = rlib.InsertRentableUseType(ctx, &ru); err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error saving rlib.RentableUseType: %s", funcname, lineno, err.Error())
			}
		}
		for i := 0; i < len(n); i++ {
			n[i].RID = rid
			n[i].BID = r.BID
			_, err := rlib.InsertRentableTypeRef(ctx, &n[i])
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error saving rlib.RentableTypeRef: %s", funcname, lineno, err.Error())
			}
		}
	}
	return 0, nil

}

// LoadRentablesCSV loads a csv file with rental specialty types and processes each one
func LoadRentablesCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateRentables)
}
