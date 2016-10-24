package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//    0  1                 2     3         4      5       6,       7
// RAID, Name,             Type, Breed,    Color, Weight, DtStart, DtStop
// 8,Santa's Little Helper,Dog,  Greyhound,gray,  34.5,  2014-01-01,

// CreateRentalAgreementPetsFromCSV reads an assessment type string array and creates a database record for a pet
func CreateRentalAgreementPetsFromCSV(sa []string, lineno int) (string, int) {
	funcname := "CreateRentalAgreementPetsFromCSV"
	var pet rlib.RentalAgreementPet
	var ok bool

	const (
		RAID   = iota
		Name   = iota
		Type   = iota
		Breed  = iota
		Color  = iota
		Weight = iota
		Dt1    = iota
		Dt2    = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"RAID", RAID},
		{"Name", Name},
		{"Type", Type},
		{"Breed", Breed},
		{"Color", Color},
		{"Weight", Weight},
		{"DtStart", Dt1},
		{"DtStop", Dt2},
	}

	rs, x := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if x > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	//-------------------------------------------------------------------
	// Find Rental Agreement
	//-------------------------------------------------------------------
	pet.RAID = CSVLoaderGetRAID(sa[RAID])
	_, err := rlib.GetRentalAgreement(pet.RAID)
	if nil != err {
		rs += fmt.Sprintf("%s: line %d - error loading Rental Agreement %s, err = %v\n", funcname, lineno, sa[0], err)
		return rs, CsvErrorSensitivity
	}

	pet.Name = strings.TrimSpace(sa[Name])
	pet.Type = strings.TrimSpace(sa[Type])
	pet.Breed = strings.TrimSpace(sa[Breed])
	pet.Color = strings.TrimSpace(sa[Color])

	//-------------------------------------------------------------------
	// Get the Weight
	//-------------------------------------------------------------------
	pet.Weight, ok = rlib.FloatFromString(sa[Weight], "Weight is invalid")
	if !ok {
		rs += fmt.Sprintf("%s: line %d - Weight is invalid: %s\n", funcname, lineno, sa[5])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	DtStart, err := rlib.StringToDate(sa[Dt1])
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[Dt1])
		return rs, CsvErrorSensitivity
	}
	pet.DtStart = DtStart

	end := "9999-01-01"
	if len(sa) > Dt2 {
		if len(sa[Dt2]) > 0 {
			end = sa[Dt2]
		}
	}
	DtStop, err := rlib.StringToDate(end)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[Dt2])
		return rs, CsvErrorSensitivity
	}
	pet.DtStop = DtStop

	_, err = rlib.InsertRentalAgreementPet(&pet)
	if nil != err {
		rs += fmt.Sprintf("%s: line %d - Could not save pet, err = %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}
	return rs, 0
}

// LoadPetsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadPetsCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateRentalAgreementPetsFromCSV(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
