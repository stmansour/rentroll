package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

// CreateRentalAgreementTemplate creates a database record for the values supplied in sa[]
func CreateRentalAgreementTemplate(sa []string) {
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if strings.ToLower(des) == "templatename" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Check to see if this assessment type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		a1, err := GetRentalAgreementTemplateByRefNum(des)
		if err != nil {
			s := err.Error()
			if !strings.Contains(s, "no rows") {
				Ulog("CreateRentalAgreementTemplate:  GetRentalAgreementTemplateByRefNum returned error %v\n", err)
			}
		}
		if len(a1.RentalTemplateNumber) > 0 {
			Ulog("CreateRentalAgreementTemplate: RentalAgreementTemplate with RentalTemplateNumber %s already exists\n", des)
			return
		}
	}

	var a RentalAgreementTemplate
	a.RentalTemplateNumber = strings.TrimSpace(sa[0])
	s := strings.TrimSpace(sa[1])
	if len(s) > 0 {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Printf("CreateRentalAgreementTemplate: RentalAgreementType value is invalid: %s\n", s)
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
		CreateRentalAgreementTemplate(t[i])
	}
}
