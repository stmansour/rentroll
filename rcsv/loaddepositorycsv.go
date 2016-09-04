package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// CVS record format:
// 0    1           2
// BUD, Name,       AccountNo

// CreateDepositoriesFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateDepositoriesFromCSV(sa []string, lineno int) {
	funcname := "CreateDepositoriesFromCSV"
	var err error
	var d rlib.Depository

	bud := strings.ToLower(strings.TrimSpace(sa[0]))
	if bud == "bud" {
		return // this is just the column heading
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 3
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(bud) > 0 {
		b1 := rlib.GetBusinessByDesignation(bud)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d - rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
			return
		}
		d.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Name
	//-------------------------------------------------------------------
	d.Name = strings.TrimSpace(sa[1])
	if len(d.Name) == 0 {
		rlib.Ulog("%s: line %d - no name for Depository. Please supply a name\n", funcname, lineno)
		return
	}

	//-------------------------------------------------------------------
	// AccountNo
	//-------------------------------------------------------------------
	d.AccountNo = strings.TrimSpace(sa[2])
	if len(d.AccountNo) == 0 {
		rlib.Ulog("%s: line %d - no AccountNo for Depository. Please supply AccountNo\n", funcname, lineno)
		return
	}

	_, err = rlib.InsertDepository(&d)
	if err != nil {
		fmt.Printf("%s: line %d -  error inserting depository: %v\n", funcname, lineno, err)
	}
}

// LoadDepositoryCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadDepositoryCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateDepositoriesFromCSV(t[i], i+1)
	}
}
