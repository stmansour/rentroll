package rcsv

import (
	"context"
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
func CreatePaymentTypeFromCSV(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreatePaymentTypeFromCSV"
	var (
		pt, dup rlib.PaymentType
		err     error
	)

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

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	des := strings.ToLower(strings.TrimSpace(sa[BUD]))

	//-------------------------------------------------------------------
	// Check to see if this rental specialty type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b, err := rlib.GetBusinessByDesignation(ctx, des)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, error while getting business by designation(%s): %s", funcname, lineno, des, err.Error())
		}
		if b.BID < 1 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Business named %s not found", funcname, lineno, des)
		}
		pt.BID = b.BID
	}

	pt.Name = strings.TrimSpace(sa[1])
	pt.Description = strings.TrimSpace(sa[2])

	// TODO(Steve): ignore error?
	_ = rlib.GetPaymentTypeByName(ctx, pt.BID, pt.Name, &dup)
	if dup.PMTID > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Skipping because payment type named %s already exists", funcname, lineno, pt.Name)
	}

	//-------------------------------------------------------------------
	// OK, just insert the record and we're done
	//-------------------------------------------------------------------
	_, err = rlib.InsertPaymentType(ctx, &pt)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting PaymentType = %v", funcname, lineno, err)
	}

	return 0, nil
}

// LoadPaymentTypesCSV loads a csv file with rental specialty types and processes each one
func LoadPaymentTypesCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreatePaymentTypeFromCSV)
}
