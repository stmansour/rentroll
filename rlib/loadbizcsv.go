package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

// 0           1    2                    3
// Designation,Name,DefaultRentalPeriod,ParkingPermitInUse
// REH,,4,0
// BBBB,Big Bob's Barrel Barn,4,0

// SetAccrual sets the DefaultRentalPeriod attribute of the Business structure based on the provided string s
func SetAccrual(s string, b *Business) {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil || !IsValidAccrual(int64(i)) {
			fmt.Printf("Invalid Accrual value: %s\n", s)
		} else {
			b.DefaultRentalPeriod = int64(i)
		}
	}
}

// CreatePhonebookLinkedBusiness creates a new business in the
// RentRoll database from the company in Phonebook with the supplied designation
func CreatePhonebookLinkedBusiness(sa []string, lineno int) {
	funcname := "CreatePhonebookLinkedBusiness"
	var b Business
	des := strings.TrimSpace(sa[0])
	found := true
	var err error

	if strings.ToLower(des) == "designation" {
		return // this is just the header line
	}

	fmt.Printf("searching for business:  %s\n", des)
	//-------------------------------------------------------------------
	// Check to see if this business is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) > 0 {
			Ulog("%s: line %d - Business Unit with designation %s already exists\n", funcname, lineno, des)
			return
		}
		found = false
	}

	//-------------------------------------------------------------------
	// It does not exist, see if we can find it in Phonebook...
	//-------------------------------------------------------------------
	if !found && len(des) > 0 {
		bu, err := GetBusinessUnitByDesignation(des)
		if nil != err {
			if !IsSQLNoResultsError(err) { // if the error is something other than "no match" then report and return
				Ulog("%s: line %d - Could not load Business Unit with Designation %s from Accord Directory: error = %v\n", funcname, lineno, des, err)
				return
			}
		} else {
			found = true
		}

		b.Name = bu.Name    // Phonebook Business Unit name
		b.Designation = des // business unit designator

		// Accrual
		SetAccrual(sa[2], &b)

		// ParkingPermitInUse
		if len(sa[3]) > 0 {
			x, err := yesnoToInt(sa[3])
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
			x, err := yesnoToInt(sa[3])
			if err != nil {
				fmt.Printf("SetParking: %s\n", err.Error())
				return
			}
			b.ParkingPermitInUse = int64(x)
		}
	}
	_, err = InsertBusiness(&b)
	if err != nil {
		Ulog("CreatePhonebookLinkedBusiness: error inserting business = %v\n", err)
	}
}

// LoadBusinessCSV loads the values from the supplied csv file and creates Business records
// as needed.
func LoadBusinessCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreatePhonebookLinkedBusiness(t[i], i+1)
	}
}
