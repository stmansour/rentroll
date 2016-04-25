package rlib

import (
	"strconv"
	"strings"
	"time"
)

func getBusinessBID(des string) int64 {
	//-------------------------------------------------------------------
	// Make sure the business exists...
	//-------------------------------------------------------------------
	b, err := GetBusinessByDesignation(des)
	if err != nil || b.BID == 0 {
		Ulog("getBusinessBID: Business with designation %s does not exist or could not be loaded\n", des)
		return 0
	}
	return b.BID
}

// CreateRentableType reads an Rentable type string array and creates a database record for the Rentable type
//  [0]        [1]      [2]   			[3]       [4]       [5]    [6]            [7]
// Designation,Style,	Name, 			Frequency,Proration,Report,ManageToBudget,MarketRate
// REH,        "GM",	"Geezer Miser", 6,		  4,	 	1,		1,			  1100.00
// REH,        "FS",	"Flat Studio",  6,		  4,	 	1,		1,			  1500.00
// REH,        "SBL",	"SB Loft",     	6,		  4,	 	1,		1,			  1750.00
// REH,        "KDS",	"KD Suite",    	6,		  4,	 	1,		1,			  2000.00
// REH,        "VEH",	Vehicle,       	3,		  0,	 	1,		1,			  10.0
// REH,        "CPT",	Carport,       	6,		  4,	 	1,		1,			  35.0
func CreateRentableType(sa []string) {
	if 8 != len(sa) {
		Ulog("CreateRentableType: csv file line \"%s\" does not have 7 elements. Ignored.\n", sa)
		return
	}
	des := strings.TrimSpace(sa[0])
	if des == "Designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Check to see if this Rentable type is already in the database
	//-------------------------------------------------------------------
	var a RentableType
	bid := getBusinessBID(des)
	if bid == 0 {
		return
	}

	a.BID = bid
	a.Style = strings.TrimSpace(sa[1])
	if len(a.Name) > 0 {
		rt, _ := GetRentableTypeByStyle(a.Style, bid)
		if rt.RTID > 0 {
			Ulog("getBusinessBID: RentableType named %s already exists\n", a.Style)
			return
		}
	}

	a.Name = strings.TrimSpace(sa[2])

	//-------------------------------------------------------------------
	// Load the values based on csv input
	//-------------------------------------------------------------------
	n, err := strconv.Atoi(strings.TrimSpace(sa[3])) // frequency
	if err != nil || n < OCCTYPEUNSET || n > OCCTYPEYEARLY {
		Ulog("CreateRentableType: Invalid rental frequency: %s\n", sa[3])
		return
	}
	a.Frequency = int64(n)

	n, err = strconv.Atoi(strings.TrimSpace(sa[4])) // Proration
	if err != nil || n < OCCTYPEUNSET || n > OCCTYPEYEARLY {
		Ulog("CreateRentableType: Invalid rental proration frequency: %s\n", sa[4])
		return
	}
	a.Proration = int64(n)

	n, err = strconv.Atoi(strings.TrimSpace(sa[5])) // report
	if err != nil || n < 0 || n > 1 {
		Ulog("CreateRentableType: Invalid report flag: %s\n", sa[5])
		return
	}
	a.Report = int64(n)

	n, err = strconv.Atoi(strings.TrimSpace(sa[6])) // manage to budget
	if err != nil || n < 0 || n > 1 {
		Ulog("CreateRentableType: Invalid manage to budget flag: %s\n", sa[6])
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
		if x, err = strconv.ParseFloat(strings.TrimSpace(sa[7]), 64); err != nil {
			Ulog("CreateRentableType: Invalid floating point number: %s\n", sa[7])
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
