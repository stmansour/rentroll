package rcsv

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"strings"
)

// 0    1
// BUD, Name
// REX,Payment
// REX,Deposit

// CreateNoteTypes reads a CustomAttributes string array and creates a database record
func CreateNoteTypes(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateNoteTypes"
	var (
		err error
		nt  rlib.NoteType
	)

	const (
		BUD  = 0
		Name = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
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
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, error while getting business by designation(%s): %s", funcname, lineno, sa[0], err.Error())
		}
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, rlib.Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		nt.BID = b1.BID
	}
	nt.Name = strings.TrimSpace(sa[1])
	if len(nt.Name) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No Name found for the NoteType", funcname, lineno)
	}
	_, err = rlib.InsertNoteType(ctx, &nt)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error inserting NoteType.  err = %s", funcname, lineno, err.Error())
	}
	return 0, nil
}

// LoadNoteTypesCSV loads a csv file with note types
func LoadNoteTypesCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateNoteTypes)
}
