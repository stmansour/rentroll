package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// 0    1
// BUD, Name
// REX,Payment
// REX,Deposit

// CreateNoteTypes reads a CustomAttributes string array and creates a database record
func CreateNoteTypes(sa []string, lineno int) {
	funcname := "CreateNoteTypes"
	var nt rlib.NoteType
	required := 2
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "bud" {
		return // this is just the column heading
	}
	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d, rlib.Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		nt.BID = b1.BID
	}
	nt.Name = strings.TrimSpace(sa[1])
	if len(nt.Name) == 0 {
		fmt.Printf("%s: line %d - No Name found for the NoteType\n", funcname, lineno)
		return
	}
	_, err := rlib.InsertNoteType(&nt)
	if err != nil {
		fmt.Printf("%s: line %d - Error inserting NoteType.  err = %s\n", funcname, lineno, err.Error())
	}
}

// LoadNoteTypesCSV loads a csv file with note types
func LoadNoteTypesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateNoteTypes(t[i], i+1)
	}
}
