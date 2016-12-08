package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//  CSV file format:
//
//        0    1
//        BUD, RATemplateName
//        REX, RAT001.doc
//        REX, RAT002.doc

// CreateRentalAgreementTemplate creates a database record for the values supplied in sa[]
func CreateRentalAgreementTemplate(sa []string, lineno int) (int, error) {
	funcname := "CreateRentalAgreementTemplate"

	const (
		BUD            = 0
		RATemplateName = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"RATemplateName", RATemplateName},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	des := strings.ToLower(strings.TrimSpace(sa[0])) // this should be BUD

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	var a rlib.RentalAgreementTemplate // start the struct we'll be saving
	if len(des) > 0 {                  // make sure it's not empty
		b1 := rlib.GetBusinessByDesignation(des) // see if we can find the biz
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
		}
		a.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Check to see if this assessment type is already in the database
	//-------------------------------------------------------------------
	des = strings.TrimSpace(sa[1]) // this should be the RATemplateName
	if len(des) > 0 {
		a1 := rlib.GetRentalAgreementByRATemplateName(des)
		if len(a1.RATemplateName) > 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - RentalAgreementTemplate with RATemplateName %s already exists\n", funcname, lineno, des)
		}
	}

	a.RATemplateName = des
	rlib.InsertRentalAgreementTemplate(&a)
	return 0, nil
}

// LoadRentalAgreementTemplatesCSV loads a csv file with assessment types and processes each one
func LoadRentalAgreementTemplatesCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateRentalAgreementTemplate)
}
