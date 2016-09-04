package rcsv

import (
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
func CreateRentalAgreementTemplate(sa []string, lineno int) int {
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

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return 1
	}
	if lineno == 1 {
		return 0
	}

	des := strings.ToLower(strings.TrimSpace(sa[0])) // this should be BUD

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	var a rlib.RentalAgreementTemplate // start the struct we'll be saving
	if len(des) > 0 {                  // make sure it's not empty
		b1 := rlib.GetBusinessByDesignation(des) // see if we can find the biz
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d, rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
			return CsvErrorSensitivity
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
			rlib.Ulog("%s: line %d - RentalAgreementTemplate with RATemplateName %s already exists\n", funcname, lineno, des)
			return CsvErrorSensitivity
		}
	}

	a.RATemplateName = des
	rlib.InsertRentalAgreementTemplate(&a)
	return 0
}

// LoadRentalAgreementTemplatesCSV loads a csv file with assessment types and processes each one
func LoadRentalAgreementTemplatesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalAgreementTemplate(t[i], i+1)
	}
}
