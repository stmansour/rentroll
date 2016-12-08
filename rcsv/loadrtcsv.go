package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// GetBusinessBID returns the BID for the rlib.Business with the supplied designation
func GetBusinessBID(des string) int64 {
	//-------------------------------------------------------------------
	// Make sure the rlib.Business exists...
	//-------------------------------------------------------------------
	b := rlib.GetBusinessByDesignation(des)
	if b.BID == 0 {
		rlib.Ulog("GetBusinessBID: rlib.Business with designation %s does not exist or could not be loaded\n", des)
		return 0
	}
	return b.BID
}

// CreateRentableType reads an rlib.Rentable type string array and creates a database record for the rlib.Rentable type
//
//                                                                               Repeat as many 3-tuples as needed
//                                                                                   /----------^-------------\
//  [0]        [1]      [2]   			[3]            [4]     5      6              7          8       9
// Designation,Style,	Name, 			RentCycle,  Proration, GSRPC, ManageToBudget,MarketRate,DtStart,DtStop
// REH,        "GM",	"Geezer Miser", 6,		       4,      4,     1,             1100.00,   1/1/2015, 1/1/2017
// REH,        "FS",	"Flat Studio",  6,		       4,      4,     1,             1500.00,   1/1/2015, 1/1/2017
// REH,        "SBL",	"SB Loft",     	6,		       4,      4,     1,             1750.00,   1/1/2015, 1/1/2017
// REH,        "KDS",	"KD Suite",    	6,		       4,      4,     1,             2000.00,   1/1/2015, 1/1/2017
// REH,        "VEH",	Vehicle,       	3,		       0,      4,     1,             10.0,   1/1/2015, 1/1/2017
// REH,        "CPT",	Carport,       	6,		       4,      4,     1,             35.0,   1/1/2015, 1/1/2017
func CreateRentableType(sa []string, lineno int) (int, error) {
	funcname := "CreateRentableType"
	const (
		BUD            = 0
		Style          = iota
		Name           = iota
		RentCycle      = iota
		Proration      = iota
		GSRPC          = iota
		ManageToBudget = iota
		MarketRate     = iota
		DtStart        = iota
		DtStop         = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Style", Style},
		{"Name", Name},
		{"RentCycle", RentCycle},
		{"Proration", Proration},
		{"GSRPC", GSRPC},
		{"ManageToBudget", ManageToBudget},
		{"MarketRate", MarketRate},
		{"DtStart", DtStart},
		{"DtStop", DtStop},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	//-------------------------------------------------------------------
	// Check to see if this rlib.Rentable type is already in the database
	//-------------------------------------------------------------------
	des := strings.TrimSpace(sa[0])
	var a rlib.RentableType
	bid := GetBusinessBID(des)
	if bid == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d  - rlib.Business named %s not found\n", funcname, lineno, sa[0])
	}

	a.BID = bid
	a.Style = strings.TrimSpace(sa[1])
	if len(a.Style) > 0 {
		rt, err := rlib.GetRentableTypeByStyle(a.Style, bid)
		if nil != err && !rlib.IsSQLNoResultsError(err) {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - err = %v\n", funcname, lineno, err)
		}
		if rt.RTID > 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - rlib.RentableType named %s already exists\n", funcname, lineno, a.Style)
		}
	}

	a.Name = strings.TrimSpace(sa[2])

	//-------------------------------------------------------------------
	// Load the values based on csv input
	//-------------------------------------------------------------------
	n, err := strconv.Atoi(strings.TrimSpace(sa[3])) // frequency
	if err != nil || !rlib.IsValidAccrual(int64(n)) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid rental frequency: %s\n", funcname, lineno, sa[3])
	}
	a.RentCycle = int64(n)

	n, err = strconv.Atoi(strings.TrimSpace(sa[4])) // Proration
	if err != nil || !rlib.IsValidAccrual(int64(n)) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid rental proration frequency: %s\n", funcname, lineno, sa[4])
	}
	a.Proration = int64(n)
	if a.Proration > a.RentCycle {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Proration frequency (%d) must be greater than rental frequency (%d)\n", funcname, lineno, a.Proration, a.RentCycle)
	}

	n, err = strconv.Atoi(strings.TrimSpace(sa[5])) // Proration
	if err != nil || !rlib.IsValidAccrual(int64(n)) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid rental GSRPC: %s\n", funcname, lineno, sa[5])
	}
	a.GSRPC = int64(n)

	n64, err := rlib.YesNoToInt(strings.TrimSpace(sa[6])) // manage to budget
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid manage to budget flag: %s\n", funcname, lineno, sa[6])
	}
	a.ManageToBudget = int64(n64)

	rtid, err := rlib.InsertRentableType(&a)

	// rlib.Rentable Market Rates are provided in 3-tuples starting at index 7 - Amount,startdata,enddate
	if rtid > 0 {
		for i := 7; i < len(sa); i += 3 {
			if len(sa[i]) == 0 { // this will happen when programs like excel save the csv file
				continue
			}
			var x float64
			var err error
			var m rlib.RentableMarketRate
			m.RTID = rtid
			if x, err = strconv.ParseFloat(strings.TrimSpace(sa[i]), 64); err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid floating point number: %s   err = %s\n", funcname, lineno, sa[i], err.Error())
			}
			m.MarketRate = x
			DtStart, err := rlib.StringToDate(sa[i+1])
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[i+1])
			}
			m.DtStart = DtStart
			DtStop, err := rlib.StringToDate(sa[i+2])
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[i+2])
			}
			m.DtStop = DtStop
			if m.DtStart.After(m.DtStop) {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Stop date (%s) must be after Start date (%s)\n", funcname, lineno, m.DtStop, m.DtStart)
			}
			rlib.InsertRentableMarketRates(&m)
		}
	}
	return 0, nil
}

// LoadRentableTypesCSV loads a csv file with rlib.Rentable types and processes each one
func LoadRentableTypesCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateRentableType)
}
