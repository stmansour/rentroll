package rcsv

import (
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
func CreateRentableSpecialtyRefsCSV(sa []string, lineno int) (string, int) {
	funcname := "CreateRentableSpecialtyRefsCSV"
	var a rlib.RentableSpecialtyRef
	var r rlib.Rentable
	var err error

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

	rs, y := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if y > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	des := strings.ToLower(strings.TrimSpace(sa[BUD]))

	var b rlib.Business
	if len(des) > 0 {
		b = rlib.GetBusinessByDesignation(des)
		if b.BID < 1 {
			rs += fmt.Sprintf("CreateRentalSpecialtyType: rlib.Business named %s not found\n", sa[0])
			return rs, CsvErrorSensitivity
		}
	}
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Find and set the rlib.Rentable
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[RID])
	if len(s) > 0 {
		// fmt.Printf("Searching: rentable name = %s, BID = %d\n", s, b.BID)
		r, err = rlib.GetRentableByName(s, b.BID)
		if err != nil {
			rs += fmt.Sprintf("%s: line %d - Error loading rlib.Rentable named: %s in Business %d.  Error = %v\n", funcname, lineno, s, b.BID, err)
			return rs, CsvErrorSensitivity
		}
	}
	a.RID = r.RID

	//-------------------------------------------------------------------
	// Make sure we can find the RentableSpecialty
	//-------------------------------------------------------------------
	name := strings.TrimSpace(sa[RentableSpecialty])
	rsp := rlib.GetRentableSpecialtyTypeByName(r.BID, name)
	if rsp.RSPID == 0 {
		rs += fmt.Sprintf("%s: line %d - could not find a rlib.RentableSpecialty named %s in rlib.Business %d\n", funcname, lineno, name, r.BID)
		return rs, CsvErrorSensitivity
	}
	a.RSPID = rsp.RSPID

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	a.DtStart, a.DtStop, err = readTwoDates(sa[DtStart], sa[DtStop], funcname, lineno)
	if err != nil {
		rs += fmt.Sprintf("%s", err.Error())
		return rs, CsvErrorSensitivity
	}

	err = rlib.InsertRentableSpecialtyRef(&a)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - error inserting assessment: %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}
	return rs, 0
}

// LoadRentableSpecialtyRefsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadRentableSpecialtyRefsCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateRentableSpecialtyRefsCSV(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
