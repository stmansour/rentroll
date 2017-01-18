package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// 0     1     2                     3         4      5       6,       7
// BUD,  RAID, Name,                 Type, Breed,    Color, Weight, DtStart, DtStop
// REX,  8,    Santa's Little Helper,Dog,  Greyhound,gray,  34.5,  2014-01-01,

// CreateRentalAgreementPetsFromCSV reads an assessment type string array and creates a database record for a pet
func CreateRentalAgreementPetsFromCSV(sa []string, lineno int) (int, error) {
	funcname := "CreateRentalAgreementPetsFromCSV"
	var pet rlib.RentalAgreementPet
	var errmsg string

	const (
		BUD    = 0
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
		{"BUD", BUD},
		{"RAID", RAID},
		{"Name", Name},
		{"Type", Type},
		{"Breed", Breed},
		{"Color", Color},
		{"Weight", Weight},
		{"DtStart", Dt1},
		{"DtStop", Dt2},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	//-------------------------------------------------------------------
	// BUD
	//-------------------------------------------------------------------
	cmpdes := strings.TrimSpace(sa[BUD])
	if len(cmpdes) > 0 {
		b2 := rlib.GetBusinessByDesignation(cmpdes)
		if b2.BID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find rlib.Business named %s\n", funcname, lineno, cmpdes)
		}
		pet.BID = b2.BID
	}

	//-------------------------------------------------------------------
	// Find Rental Agreement
	//-------------------------------------------------------------------
	pet.RAID = CSVLoaderGetRAID(sa[RAID])
	_, err = rlib.GetRentalAgreement(pet.RAID)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error loading Rental Agreement %s, err = %v\n", funcname, lineno, sa[0], err)
	}

	pet.Name = strings.TrimSpace(sa[Name])
	pet.Type = strings.TrimSpace(sa[Type])
	pet.Breed = strings.TrimSpace(sa[Breed])
	pet.Color = strings.TrimSpace(sa[Color])

	//-------------------------------------------------------------------
	// Get the Weight
	//-------------------------------------------------------------------
	pet.Weight, errmsg = rlib.FloatFromString(sa[Weight], "Weight is invalid")
	if len(errmsg) > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Weight is invalid: %s  (%s)\n", funcname, lineno, sa[5], errmsg)
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	DtStart, err := rlib.StringToDate(sa[Dt1])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[Dt1])
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
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[Dt2])
	}
	pet.DtStop = DtStop

	_, err = rlib.InsertRentalAgreementPet(&pet)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not save pet, err = %v\n", funcname, lineno, err)
	}
	return 0, nil
}

// LoadPetsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadPetsCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateRentalAgreementPetsFromCSV)
}
