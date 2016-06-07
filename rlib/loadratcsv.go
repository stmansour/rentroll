package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

//  CSV file format:
//
//        0            1
//        TemplateName,RentalAgreementType
//        "RAT001",    2
//        "RAT002",    2

// CreateRentalAgreementTemplate creates a database record for the values supplied in sa[]
func CreateRentalAgreementTemplate(sa []string, lineno int) {
	funcname := "CreateRentalAgreementTemplate"
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if strings.ToLower(des) == "templatename" {
		return // this is just the column heading
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 2
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Check to see if this assessment type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		a1, err := GetRentalAgreementTemplateByRefNum(des)
		if err != nil {
			s := err.Error()
			if !strings.Contains(s, "no rows") {
				Ulog("%s: line %d -  GetRentalAgreementTemplateByRefNum returned error %v\n", funcname, lineno, err)
			}
		}
		if len(a1.RentalTemplateNumber) > 0 {
			Ulog("%s: line %d - RentalAgreementTemplate with RentalTemplateNumber %s already exists\n", funcname, lineno, des)
			return
		}
	}

	var a RentalAgreementTemplate
	a.RentalTemplateNumber = strings.TrimSpace(sa[0])
	s := strings.TrimSpace(sa[1])
	if len(s) > 0 {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Printf("%s: line %d - RentalAgreementType value is invalid: %s\n", funcname, lineno, s)
			return
		}
		a.RentalAgreementType = int64(i)
	}
	InsertRentalAgreementTemplate(&a)
}

// LoadRentalAgreementTemplatesCSV loads a csv file with assessment types and processes each one
func LoadRentalAgreementTemplatesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalAgreementTemplate(t[i], i+1)
	}
}
