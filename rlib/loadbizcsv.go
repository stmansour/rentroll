package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

// 0           1    2                    3
// Designation,Name,DefaultOccupancyType,ParkingPermitInUse
// REH,,4,0
// BBBB,Big Bob's Barrel Barn,4,0

// SetOccType sets the DefaultOccupancyType attribute of the Business structure based on the provided string s
func SetOccType(s string, b *Business) {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil || i < OCCTYPEUNSET || i > OCCTYPEYEARLY {
			fmt.Printf("Invalid OccupancyType value: %s\n", s)
		} else {
			b.DefaultOccupancyType = int64(i)
		}
	}
}

// CreatePhonebookLinkedBusiness creates a new business in the
// RentRoll database from the company in Phonebook with the supplied designation
func CreatePhonebookLinkedBusiness(sa []string) {
	des := strings.TrimSpace(sa[0])
	found := false
	var err error

	if strings.ToLower(des) == "designation" {
		return // this is just the header line
	}

	//-------------------------------------------------------------------
	// Check to see if this business is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) > 0 {
			Ulog("CreatePhonebookLinkedBusiness: business with designation %s already exists\n", des)
			return
		}
	}

	//-------------------------------------------------------------------
	// It does not exist, see if we can find it in Phonebook...
	//-------------------------------------------------------------------
	var b Business
	if len(des) > 0 {
		c, err := GetCompanyByDesignation(des)
		if err != nil {
			e := fmt.Sprintf("%v", err)
			if !strings.Contains(e, "no rows in result") {
				Ulog("GetCompanyByDesignation: error = %s\n", e)
			}
		}
		if c.CoCode > 0 {
			found = true
			b.Name = c.CommonName
			b.Designation = des
			if len(b.Name) == 0 {
				b.Name = des
			}
			SetOccType(sa[2], &b)
			if len(sa[3]) > 0 {
				x, err := yesnoToInt(sa[3])
				if err != nil {
					fmt.Printf("SetParking: %s\n", err.Error())
					return
				}
				b.ParkingPermitInUse = int64(x)
			}
		}
	}

	//-------------------------------------------------------------------
	// If we did not find it in Phonebook, we still need to create it,
	// so use the fields we have...
	//-------------------------------------------------------------------
	if !found {
		b.Name = sa[1]
		b.Designation = des
		SetOccType(sa[2], &b)
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
		CreatePhonebookLinkedBusiness(t[i])
	}
}
