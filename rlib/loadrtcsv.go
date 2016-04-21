package rlib

import (
	"strconv"
	"strings"
	"time"
)

// CreateRentableType reads an Rentable type string array and creates a database record for the Rentable type
func CreateRentableType(sa []string) {
	if 7 != len(sa) {
		Ulog("CreateRentableType: csv file line \"%s\" does not have 7 elements. Ignored.\n", sa)
		return
	}
	des := strings.TrimSpace(sa[0])
	if des == "Designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Make sure the business exists...
	//-------------------------------------------------------------------
	b, err := GetBusinessByDesignation(des)
	if err != nil || b.BID == 0 {
		Ulog("CreateRentableType: Business with designation %s does not exist or could not be loaded\n", des)
		return
	}

	//-------------------------------------------------------------------
	// Check to see if this Rentable type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		a1, _ := GetRentableTypeByName(des, b.BID)
		if len(a1.Name) > 0 {
			Ulog("CreateRentableType: RentableType named %s already exists\n", des)
			return
		}
	}

	//-------------------------------------------------------------------
	// Load the values based on csv input
	//-------------------------------------------------------------------
	var a RentableType
	a.BID = b.BID
	a.Name = strings.TrimSpace(sa[1])

	n, err := strconv.Atoi(sa[2]) // frequency
	if err != nil || n < OCCTYPEUNSET || n > OCCTYPEYEARLY {
		Ulog("CreateRentableType: Invalid rental frequency: %s\n", sa[2])
		return
	}
	a.Frequency = int64(n)

	n, err = strconv.Atoi(sa[3]) // Proration
	if err != nil || n < OCCTYPEUNSET || n > OCCTYPEYEARLY {
		Ulog("CreateRentableType: Invalid rental proration frequency: %s\n", sa[3])
		return
	}
	a.Proration = int64(n)

	n, err = strconv.Atoi(sa[4]) // report
	if err != nil || n < 0 || n > 1 {
		Ulog("CreateRentableType: Invalid report flag: %s\n", sa[4])
		return
	}
	a.Report = int64(n)

	n, err = strconv.Atoi(sa[5]) // manage to budget
	if err != nil || n < 0 || n > 1 {
		Ulog("CreateRentableType: Invalid manage to budget flag: %s\n", sa[5])
		return
	}
	a.ManageToBudget = int64(n)

	rtid, err := InsertRentableType(&a)
	if rtid > 0 {
		var x float64
		var err error
		var m RentableMarketRate
		m.RTID = rtid
		m.DtStart = time.Now()
		m.DtStop = time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
		if x, err = strconv.ParseFloat(sa[6], 64); err != nil {
			Ulog("CreateRentableType: Invalid floating point number: %s\n", sa[6])
			return
		}
		m.MarketRate = x
		InsertRentableMarketRates(&m)
	}
}

// LoadRentableTypesCSV loads a csv file with Rentable types and processes each one
func LoadRentableTypesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentableType(t[i])
	}
}
