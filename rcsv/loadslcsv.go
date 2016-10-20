package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

var (
	bud string          // business, if changed then write the stringlist
	a   rlib.StringList // the string list we build up
)

//  CSV file format:
//
//        0    1              2
//        BUD, Name,          Value
//        REX, ApplDenyReason,Bad Credit
//        REX, ApplDenyReason,Criminal Record
//        REX, ApplDenyReason,Bad references
//		  REX, MoveOutReason,New Job
//		  REX, MoveOutReason,Can't afford it

func writeStringList() error {
	var err error
	if len(a.Name) > 0 {
		var t rlib.StringList
		rlib.GetStringListByName(a.BID, a.Name, &t) // do we already have a stringlist by this name?
		if t.SLID > 0 {                             // t.SLID will be nonzero if so
			err = rlib.DeleteStringList(t.SLID) // delete what's there if it exists
			if err != nil {
				return err
			}
		}
	}
	_, err = rlib.InsertStringList(&a) // update the db with this list
	if err != nil {
		return err
	}
	var b rlib.StringList
	a = b // reset the list so we can build up the new one
	return err
}

// CreateStringList creates a database record for the values supplied in sa[]
func CreateStringList(sa []string, lineno int) (string, int) {
	funcname := "CreateStringList"
	rs := ""
	const (
		BUD   = 0
		Name  = iota
		Value = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"Value", Value},
	}

	rs, x := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if x > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[0])) // this should be BUD
	if len(des) > 0 {                                // make sure it's not empty
		b1 := rlib.GetBusinessByDesignation(des) // see if we can find the biz
		if len(b1.Designation) == 0 {
			rs += fmt.Sprintf("%s: line %d, Business with designation %s does not exist\n", funcname, lineno, sa[0])
			return rs, CsvErrorSensitivity
		}
		// if business is changed, write the stringlist
		if len(bud) > 0 && des != bud {
			writeStringList()
		}
		a.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Name
	//-------------------------------------------------------------------
	name := strings.TrimSpace(sa[1])
	if len(name) > 0 {
		if len(a.Name) > 0 && a.Name != name { // did the name of the StringList change from the last time?
			writeStringList() // yes: write what we have and start a new one
			bud = des         // only the Name changed, not the business. Restore the bud value
		}
		a.Name = name
	}

	//-------------------------------------------------------------------
	// Value
	//-------------------------------------------------------------------
	var sls rlib.SLString
	sls.Value = strings.TrimSpace(sa[2])
	a.S = append(a.S, sls)
	return rs, 0
}

// LoadStringTablesCSV loads a csv file with assessment types and processes each one
func LoadStringTablesCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateStringList(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	if len(a.S) > 0 {
		err := writeStringList()
		if err != nil {
			rs += fmt.Sprintf("Error writing string list: %s\n", err.Error())
		}
	}
	return rs
}
