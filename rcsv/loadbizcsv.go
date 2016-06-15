package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// 0           1    2                    3
// Designation,Name,DefaultRentalPeriod,ParkingPermitInUse
// REH,,4,0
// BBBB,Big Bob's Barrel Barn,4,0

// SetAccrual sets the DefaultRentalPeriod attribute of the rlib.Business structure based on the provided string s
func SetAccrual(s string, b *rlib.Business) {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil || !rlib.IsValidAccrual(int64(i)) {
			fmt.Printf("Invalid Accrual value: %s\n", s)
		} else {
			b.DefaultRentalPeriod = int64(i)
		}
	}
}

// CreatePhonebookLinkedBusiness creates a new rlib.Business in the
// RentRoll database from the company in Phonebook with the supplied designation
func CreatePhonebookLinkedBusiness(sa []string, lineno int) {
	funcname := "CreatePhonebookLinkedBusiness"
	var b rlib.Business
	des := strings.TrimSpace(sa[0])
	found := true
	var err error

	if strings.ToLower(des) == "designation" {
		return // this is just the header line
	}

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 4
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	// fmt.Printf("searching for rlib.Business:  %s\n", des)
	//-------------------------------------------------------------------
	// Check to see if this rlib.Business is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) > 0 {
			rlib.Ulog("%s: line %d - rlib.Business Unit with designation %s already exists\n", funcname, lineno, des)
			return
		}
		found = false
	}

	//-------------------------------------------------------------------
	// It does not exist, see if we can find it in Phonebook...
	//-------------------------------------------------------------------
	if !found && len(des) > 0 {
		bu, err := rlib.GetBusinessUnitByDesignation(des)
		if nil != err {
			if !rlib.IsSQLNoResultsError(err) { // if the error is something other than "no match" then report and return
				rlib.Ulog("%s: line %d - Could not load rlib.Business Unit with Designation %s from Accord Directory: error = %v\n", funcname, lineno, des, err)
				return
			}
		} else {
			found = true
		}

		b.Name = bu.Name    // Phonebook rlib.Business Unit name
		b.Designation = des // rlib.Business unit designator

		// Accrual
		SetAccrual(sa[2], &b)

		// ParkingPermitInUse
		if len(sa[3]) > 0 {
			x, err := rlib.YesNoToInt(sa[3])
			if err != nil {
				fmt.Printf("SetParking: %s\n", err.Error())
				return
			}
			b.ParkingPermitInUse = int64(x)
		}
	}

	//-------------------------------------------------------------------
	// If we did not find it in Phonebook, we still need to create it,
	// so use the fields we have...
	//-------------------------------------------------------------------
	if !found {
		b.Name = sa[1]
		b.Designation = des
		SetAccrual(sa[2], &b)
		if len(sa[3]) > 0 {
			x, err := rlib.YesNoToInt(sa[3])
			if err != nil {
				fmt.Printf("SetParking: %s\n", err.Error())
				return
			}
			b.ParkingPermitInUse = int64(x)
		}
	}
	_, err = rlib.InsertBusiness(&b)
	if err != nil {
		rlib.Ulog("CreatePhonebookLinkedBusiness: error inserting rlib.Business = %v\n", err)
	}
}

// LoadBusinessCSV loads the values from the supplied csv file and creates rlib.Business records
// as needed.
func LoadBusinessCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreatePhonebookLinkedBusiness(t[i], i+1)
	}
}
