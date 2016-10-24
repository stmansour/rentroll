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
func CreateNoteTypes(sa []string, lineno int) (string, int) {
	funcname := "CreateNoteTypes"
	var nt rlib.NoteType
	const (
		BUD  = 0
		Name = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
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
	des := strings.ToLower(strings.TrimSpace(sa[BUD]))
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d, rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
			return rs, CsvErrorSensitivity
		}
		nt.BID = b1.BID
	}
	nt.Name = strings.TrimSpace(sa[1])
	if len(nt.Name) == 0 {
		fmt.Printf("%s: line %d - No Name found for the NoteType\n", funcname, lineno)
		return rs, CsvErrorSensitivity
	}
	_, err := rlib.InsertNoteType(&nt)
	if err != nil {
		fmt.Printf("%s: line %d - Error inserting NoteType.  err = %s\n", funcname, lineno, err.Error())
		return rs, CsvErrorSensitivity
	}
	return rs, 0
}

// LoadNoteTypesCSV loads a csv file with note types
func LoadNoteTypesCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateNoteTypes(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
