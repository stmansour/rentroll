package rcsv

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"strings"
)

// CSV FIELDS FOR THIS MODULE
//    0    1    2                3         4
//    BUD, RID, RentalSpecialty, DtStart,  DtStStop
//    REX, 1,   Lake View,       1/1/2014,
//    REX, 1,   Fireplace,       1/1/2014,

// type rlib.RentableSpecialtyRef struct {
// 	RID         int64     // the rlib.Rentable to which this record belongs
// 	RSPID       int64     // the rentable specialty type associated with the rentable
// 	DtStart     time.Time // timerange start
// 	DtStop      time.Time // timerange stop
// 	LastModTime time.Time
// 	LastModBy   int64
// }

// CreateRentableSpecialtyRefsCSV reads an assessment type string array and creates a database record for the assessment type
func CreateRentableSpecialtyRefsCSV(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateRentableSpecialtyRefsCSV"
	var (
		err error
		a   rlib.RentableSpecialtyRef
		r   rlib.Rentable
	)

	const (
		BUD               = 0
		RID               = iota
		RentableSpecialty = iota
		DtStart           = iota
		DtStop            = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"RID", RID},
		{"RentableSpecialty", RentableSpecialty},
		{"DtStart", DtStart},
		{"DtStop", DtStop},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	des := strings.ToLower(strings.TrimSpace(sa[BUD]))

	var b rlib.Business
	if len(des) > 0 {
		b, err = rlib.GetBusinessByDesignation(ctx, des)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, error while getting business by designation(%s): %s", funcname, lineno, des, err.Error())
		}
		if b.BID < 1 {
			return CsvErrorSensitivity, fmt.Errorf("CreateRentalSpecialtyType: rlib.Business named %s not found", sa[0])
		}
	}
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Find and set the rlib.Rentable
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[RID])
	if len(s) > 0 {
		// fmt.Printf("Searching: rentable name = %s, BID = %d\n", s, b.BID)
		r, err = rlib.GetRentableByName(ctx, s, b.BID)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error loading rlib.Rentable named: %s in Business %d.  Error = %v", funcname, lineno, s, b.BID, err)
		}
	}
	a.RID = r.RID

	//-------------------------------------------------------------------
	// Make sure we can find the RentableSpecialty
	//-------------------------------------------------------------------
	name := strings.TrimSpace(sa[RentableSpecialty])
	rsp, err := rlib.GetRentableSpecialtyTypeByName(ctx, r.BID, name)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error getting rlib.RentableSpecialty named %s in rlib.Business %d: %s", funcname, lineno, name, r.BID, err.Error())
	}
	if rsp.RSPID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find a rlib.RentableSpecialty named %s in rlib.Business %d", funcname, lineno, name, r.BID)
	}
	a.RSPID = rsp.RSPID

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	a.DtStart, a.DtStop, err = readTwoDates(sa[DtStart], sa[DtStop], funcname, lineno, "DtStart/DtStop")
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s", err.Error())
	}

	_, err = rlib.InsertRentableSpecialtyRef(ctx, &a)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting assessment: %v", funcname, lineno, err)
	}
	return 0, nil
}

// LoadRentableSpecialtyRefsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadRentableSpecialtyRefsCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateRentableSpecialtyRefsCSV)
}
