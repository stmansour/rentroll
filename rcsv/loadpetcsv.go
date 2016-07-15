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
func CreateRentalAgreementPetsFromCSV(sa []string, lineno int) {
	funcname := "CreateRentalAgreementPetsFromCSV"
	var pet rlib.RentalAgreementPet
	var ok bool

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	if len(sa) < 7 {
		fmt.Printf("%s: line %d - found %d values, there must be at least 7\n", funcname, lineno, len(sa))
		return
	}

	if strings.ToLower(strings.TrimSpace(sa[0])) == "raid" {
		return // this is the header line
	}
	//-------------------------------------------------------------------
	// Find Rental Agreement
	//-------------------------------------------------------------------
	pet.RAID = CSVLoaderGetRAID(sa[0])
	_, err := rlib.GetRentalAgreement(pet.RAID)
	if nil != err {
		fmt.Printf("%s: line %d - error loading Rental Agreement %s, err = %v\n", funcname, lineno, sa[0], err)
		return
	}

	pet.Name = strings.TrimSpace(sa[1])
	pet.Type = strings.TrimSpace(sa[2])
	pet.Breed = strings.TrimSpace(sa[3])
	pet.Color = strings.TrimSpace(sa[4])

	//-------------------------------------------------------------------
	// Get the Weight
	//-------------------------------------------------------------------
	pet.Weight, ok = rlib.FloatFromString(sa[5], "Weight is invalid")
	if !ok {
		fmt.Printf("%s: line %d - Weight is invalid: %s\n", funcname, lineno, sa[5])
		return
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	DtStart, err := StringToDate(sa[6])
	if err != nil {
		fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[6])
		return
	}
	pet.DtStart = DtStart

	end := "9999-01-01"
	if len(sa) > 7 {
		if len(sa[7]) > 0 {
			end = sa[7]
		}
	}
	DtStop, err := StringToDate(end)
	if err != nil {
		fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[7])
		return
	}
	pet.DtStop = DtStop

	_, err = rlib.InsertRentalAgreementPet(&pet)
	if nil != err {
		fmt.Printf("%s: line %d - Could not save pet, err = %v\n", funcname, lineno, err)
	}
}

// LoadPetsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadPetsCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalAgreementPetsFromCSV(t[i], i+1)
	}
}
