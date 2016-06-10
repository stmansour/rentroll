package rlib

import (
	"fmt"
	"strings"
)

//  CSV file format:
//
//        0    1
//        BUD, RentalTemplateNumber
//        REX, RAT001.doc
//        REX, RAT002.doc

// CreateRentalAgreementTemplate creates a database record for the values supplied in sa[]
func CreateRentalAgreementTemplate(sa []string, lineno int) {
	funcname := "CreateRentalAgreementTemplate"

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 2
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	des := strings.ToLower(strings.TrimSpace(sa[0])) // this should be BUD
	if strings.ToLower(des) == "bud" {
		return // this is just the column heading
	}
	//-------------------------------------------------------------------
	// Make sure the Business is in the database
	//-------------------------------------------------------------------
	var a RentalAgreementTemplate // start the struct we'll be saving
	if len(des) > 0 {             // make sure it's not empty
		b1, _ := GetBusinessByDesignation(des) // see if we can find the biz
		if len(b1.Designation) == 0 {
			Ulog("%s: line %d, Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		a.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Check to see if this assessment type is already in the database
	//-------------------------------------------------------------------
	des = strings.TrimSpace(sa[1]) // this should be the RentalTemplateNumber
	if len(des) > 0 {
		a1, err := GetRentalAgreementByRentalTemplateNumber(des)
		if err != nil {
			s := err.Error()
			if !strings.Contains(s, "no rows") {
				Ulog("%s: line %d -  GetRentalAgreementByRentalTemplateNumber returned error %v\n", funcname, lineno, err)
			}
		}
		if len(a1.RentalTemplateNumber) > 0 {
			Ulog("%s: line %d - RentalAgreementTemplate with RentalTemplateNumber %s already exists\n", funcname, lineno, des)
			return
		}
	}

	a.RentalTemplateNumber = des

	InsertRentalAgreementTemplate(&a)
}

// LoadRentalAgreementTemplatesCSV loads a csv file with assessment types and processes each one
func LoadRentalAgreementTemplatesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalAgreementTemplate(t[i], i+1)
	}
}
