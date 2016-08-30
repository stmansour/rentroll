package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// 0           1    2                3,                    4,
// Bud,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC
// REH,,4,0
// BBBB,Big Bob's Barrel Barn,4,0

// GetAccrual sets the DefaultRentCycle attribute of the rlib.Business structure based on the provided string s
func GetAccrual(s string) (int64, bool) {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err == nil && rlib.IsValidAccrual(int64(i)) {
			return int64(i), true
		}
	}
	return int64(0), false
}

// CreatePhonebookLinkedBusiness creates a new rlib.Business in the
// RentRoll database from the company in Phonebook with the supplied designation
func CreatePhonebookLinkedBusiness(sa []string, lineno int) int {
	funcname := "CreatePhonebookLinkedBusiness"
	var b rlib.Business
	des := strings.TrimSpace(sa[0])
	found := true
	var err error
	var ok bool

	const (
		BUD                   = 0
		Name                  = iota
		DefaultRentCycle      = iota
		DefaultProrationCycle = iota
		DefaultGSRPC          = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"DefaultRentCycle", DefaultRentCycle},
		{"DefaultProrationCycle", DefaultProrationCycle},
		{"DefaultGSRPC", DefaultGSRPC},
	}

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return 1
	}
	if lineno == 1 {
		return 0
	}

	//-------------------------------------------------------------------
	// Check to see if this rlib.Business is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) > 0 {
			rlib.Ulog("%s: line %d - rlib.Business Unit with designation %s already exists\n", funcname, lineno, des)
			return 1
		}
		found = false
	}

	//-------------------------------------------------------------------
	// It does not exist, see if we can find it in Phonebook...
	//-------------------------------------------------------------------
	if !found && len(des) > 0 {
		bu, err := rlib.GetBusinessUnitByDesignation(des)
		if nil != err {
			if !rlib.IsSQLNoResultsError(err) { // if the error is something other than "no match" then report and return 1
				rlib.Ulog("%s: line %d - Could not load rlib.Business Unit with Designation %s from Accord Directory: error = %v\n", funcname, lineno, des, err)
				return 1
			}
		} else {
			found = true
		}
		b.Name = bu.Name    // Phonebook rlib.Business Unit name
		b.Designation = des // rlib.Business unit designator
	}

	//-----------------------------------------
	// DefaultRentCycle
	//-----------------------------------------
	if b.DefaultRentCycle, ok = GetAccrual(strings.TrimSpace(sa[2])); !ok {
		fmt.Printf("%s: line %d - Invalid Rent Cycle: %s\n", funcname, lineno, sa[2])
		return 1
	}

	//-----------------------------------------
	// DefaultProrationCycle
	//-----------------------------------------
	if b.DefaultProrationCycle, ok = GetAccrual(strings.TrimSpace(sa[3])); !ok {
		fmt.Printf("%s: line %d - Invalid Proration Cycle: %s\n", funcname, lineno, sa[3])
		return 1
	}

	//-----------------------------------------
	// DefaultGSRPC
	//-----------------------------------------
	if b.DefaultGSRPC, ok = GetAccrual(strings.TrimSpace(sa[4])); !ok {
		fmt.Printf("%s: line %d - Invalid GSRPC: %s\n", funcname, lineno, sa[4])
		return 1
	}

	//-------------------------------------------------------------------
	// If we did not find it in Phonebook, we still need to create it,
	// so use the fields we have...
	//-------------------------------------------------------------------
	if !found {
		b.Name = strings.TrimSpace(sa[1])
		b.Designation = des
	}
	_, err = rlib.InsertBusiness(&b)
	if err != nil {
		rlib.Ulog("CreatePhonebookLinkedBusiness: error inserting rlib.Business = %v\n", err)
	}
	return 0
}

// LoadBusinessCSV loads the values from the supplied csv file and creates rlib.Business records
// as needed.
func LoadBusinessCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if CreatePhonebookLinkedBusiness(t[i], i+1) > 0 {
			return
		}
	}
}
