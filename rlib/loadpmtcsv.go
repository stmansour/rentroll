package rlib

import (
	"fmt"
	"strings"
)

// 0            1    2
// Designation, Name,Description
// REH,"Check","Personal check from payor"
// REH,"VISA","Credit card charge"
// REH,"AMEX", "American Express credit card"
// REH,"Cash","Cash"

// CreatePaymentTypeFromCSV reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreatePaymentTypeFromCSV(sa []string) {
	var pt PaymentType
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Check to see if this rental specialty type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b, _ := GetBusinessByDesignation(des)
		if b.BID < 1 {
			Ulog("CreatePaymentTypeFromCSVType: Business named %s not found\n", des)
			return
		}
		pt.BID = b.BID
	}

	pt.Name = strings.TrimSpace(sa[1])
	pt.Description = strings.TrimSpace(sa[2])

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	err := InsertPaymentType(&pt)
	if nil != err {
		fmt.Printf("CreatePaymentTypeFromCSV: error inserting PaymentType = %v\n", err)
	}
}

// LoadPaymentTypesCSV loads a csv file with rental specialty types and processes each one
func LoadPaymentTypesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreatePaymentTypeFromCSV(t[i])
	}
}
