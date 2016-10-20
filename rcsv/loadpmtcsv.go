package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// 0    1      2
// BUD, Name,  Description
// REH,"Check","Personal check from rlib.Payor"
// REH,"VISA", "Credit card charge"
// REH,"AMEX", "American Express credit card"
// REH,"Cash", "Cash"

// CreatePaymentTypeFromCSV reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreatePaymentTypeFromCSV(sa []string, lineno int) (string, int) {
	funcname := "CreatePaymentTypeFromCSV"
	var pt rlib.PaymentType
	const (
		BUD         = 0
		Name        = iota
		Description = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"Description", Description},
	}

	rs, x := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if x > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	des := strings.ToLower(strings.TrimSpace(sa[0]))

	//-------------------------------------------------------------------
	// Check to see if this rental specialty type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b := rlib.GetBusinessByDesignation(des)
		if b.BID < 1 {
			rs += fmt.Sprintf("%s: line %d - Business named %s not found\n", funcname, lineno, des)
			return rs, CsvErrorSensitivity
		}
		pt.BID = b.BID
	}

	pt.Name = strings.TrimSpace(sa[1])
	pt.Description = strings.TrimSpace(sa[2])

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	err := rlib.InsertPaymentType(&pt)
	if nil != err {
		rs += fmt.Sprintf("%s: line %d - error inserting PaymentType = %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}

	return rs, 0
}

// LoadPaymentTypesCSV loads a csv file with rental specialty types and processes each one
func LoadPaymentTypesCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreatePaymentTypeFromCSV(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
