package rlib

import (
	"strconv"
	"strings"
	"time"
)

// CreateUnitType reads an Unit type string array and creates a database record for the Unit type
func CreateUnitType(sa []string) {
	if 7 != len(sa) {
		Ulog("CreateUnitType: csv file line \"%s\" does not have 7 elements. Ignored.\n", sa)
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
		Ulog("CreateUnitType: Business with designation %s does not exist or could not be loaded\n", des)
		return
	}

	//-------------------------------------------------------------------
	// Check to see if this Unit type is already in the database
	//-------------------------------------------------------------------
	var a UnitType
	a.BID = b.BID
	a.Style = strings.TrimSpace(sa[1])

	if len(a.Style) > 0 {
		ut, _ := GetUnitTypeByStyle(a.Style, b.BID)
		if ut.UTID > 0 {
			Ulog("CreateUnitType: UnitType named %s already exists\n", des)
			return
		}
	}

	//-------------------------------------------------------------------
	// Load the values based on csv input
	//-------------------------------------------------------------------

	n, err := strconv.Atoi(sa[2]) // frequency
	if err != nil || n < OCCTYPEUNSET || n > OCCTYPEYEARLY {
		Ulog("CreateUnitType: Invalid rental frequency: %s\n", sa[2])
		return
	}
	a.Frequency = int64(n)

	n, err = strconv.Atoi(sa[3]) // Proration
	if err != nil || n < OCCTYPEUNSET || n > OCCTYPEYEARLY {
		Ulog("CreateUnitType: Invalid rental proration frequency: %s\n", sa[3])
		return
	}
	a.Proration = int64(n)

	n, err = strconv.Atoi(sa[4]) // report
	if err != nil || n < 0 || n > 1 {
		Ulog("CreateUnitType: Invalid report flag: %s\n", sa[4])
		return
	}
	a.Report = int64(n)

	n, err = strconv.Atoi(sa[5]) // manage to budget
	if err != nil || n < 0 || n > 1 {
		Ulog("CreateUnitType: Invalid manage to budget flag: %s\n", sa[5])
		return
	}
	a.ManageToBudget = int64(n)

	rtid, err := InsertUnitType(&a)
	if rtid > 0 {
		var x float64
		var err error
		var m UnitMarketRate
		m.UTID = rtid
		m.DtStart = time.Now()
		m.DtStop = time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)
		if x, err = strconv.ParseFloat(sa[6], 64); err != nil {
			Ulog("CreateUnitType: Invalid floating point number: %s\n", sa[6])
			return
		}
		m.MarketRate = x
		InsertUnitMarketRates(&m)
	}
}

// LoadUnitTypesCSV loads a csv file with Unit types and processes each one
func LoadUnitTypesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateUnitType(t[i])
	}
}
