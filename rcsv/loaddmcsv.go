package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//  CSV file format:
//
//        0    1
//        BUD, Name
//        REX, Hand Delivery
//        REX, Scanned Batch
//        REX, CC Shift 4, CC NAYAX, ACH, US Mail...

// CreateDepositMethod creates a database record for the values supplied in sa[]
func CreateDepositMethod(sa []string, lineno int) {
	funcname := "CreateDepositMethod"

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 2
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	des := strings.ToLower(strings.TrimSpace(sa[0])) // this should be BUD
	if strings.ToLower(des) == "bud" {
		return // this is just the column heading
	}
	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	var a rlib.DepositMethod // start the struct we'll be saving
	if len(des) > 0 {        // make sure it's not empty
		b1, _ := rlib.GetBusinessByDesignation(des) // see if we can find the biz
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d, Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		a.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Check to see if this name type is already in the database
	//-------------------------------------------------------------------
	name := strings.TrimSpace(sa[1]) // this should be the RentalTemplateNumber
	if len(name) > 0 {
		a1, err := rlib.GetDepositMethodByName(a.BID, name)
		if err != nil {
			s := err.Error()
			if !strings.Contains(s, "no rows") {
				rlib.Ulog("%s: line %d -   returned error %v\n", funcname, lineno, err)
			}
		}
		if len(a1.Name) > 0 {
			rlib.Ulog("%s: line %d - DepositMethod with Name %s already exists\n", funcname, lineno, name)
			return
		}
	}

	a.Name = name
	rlib.InsertDepositMethod(&a)
}

// LoadDepositMethodsCSV loads a csv file with assessment types and processes each one
func LoadDepositMethodsCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateDepositMethod(t[i], i+1)
	}
}
