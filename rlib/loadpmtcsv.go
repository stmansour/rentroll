package rlib

import (
	"fmt"
	"strings"
)

// type Receipt struct {
// 	RCPTID   int64
// 	BID      int64
// 	RAID     int64
// 	PMTID    int64
// 	Dt       time.Time
// 	Amount   float64
// 	AcctRule string
// 	Comment  string
// 	RA       []ReceiptAllocation
// }

// type ReceiptAllocation struct {
// 	RCPTID   int64
// 	Amount   float64
// 	ASMID    int64
// 	AcctRule string
// }

// 0            1    2
// Designation, Name,Description
// REH,"Check","Personal check from payor"
// REH,"VISA","Credit card charge"
// REH,"AMEX", "American Express credit card"
// REH,"Cash","Cash"

// CreatePaymentTypeFromCSV reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreatePaymentTypeFromCSV(sa []string, lineno int) {
	funcname := "CreatePaymentTypeFromCSV"
	var pt PaymentType
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 3
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Check to see if this rental specialty type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b, _ := GetBusinessByDesignation(des)
		if b.BID < 1 {
			Ulog("%s: line %d - Business named %s not found\n", funcname, lineno, des)
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
		fmt.Printf("%s: line %d - error inserting PaymentType = %v\n", funcname, lineno, err)
	}
}

// LoadPaymentTypesCSV loads a csv file with rental specialty types and processes each one
func LoadPaymentTypesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreatePaymentTypeFromCSV(t[i], i+1)
	}
}
