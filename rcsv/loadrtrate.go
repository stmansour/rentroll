package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//  CSV file format:
//                             RT name or style  string with or without %
// 0    1             2        3                 4
// BUD, RPName,       RPRID    RentableType,     Amount
// REX, FAA-P, RPR0001, GM,               85%
// REX, FAA-P,   1,     Flat Studio,      1400
// REX, FAA-P,   1,     SBL,    			 1500
// REX, FAA-P,   1,     KDS,    			 75%
// REX, FAA-T,      1,     GM,               90%
// REX, FAA-T,      1,     Flat Studio,      90%
// REX, FAA-T,      1,     SBL,    			 1500
// REX, FAA-T,      1,     KDS,    			 80%

// CreateRatePlanRefRTRate reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRatePlanRefRTRate(sa []string, lineno int) {
	funcname := "CreateRatePlanRefRTRate"
	var b rlib.Business

	required := 5
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}
	//-------------------------------------------------------------------
	// BUD
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "bud" {
		return // this is just the column heading
	}
	if len(des) > 0 {
		b, _ = rlib.GetBusinessByDesignation(des)
		if len(b.Designation) == 0 {
			rlib.Ulog("%s: line %d, rlib.Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
	}

	//-------------------------------------------------------------------
	// RatePlan Name
	//-------------------------------------------------------------------
	var rp rlib.RatePlan
	rpname := strings.ToLower(strings.TrimSpace(sa[1]))
	if len(rpname) > 0 {
		rlib.GetRatePlanByName(b.BID, rpname, &rp)
		if rp.RPID < 1 {
			rlib.Ulog("%s: line %d - RatePlan named %s not found\n", funcname, lineno, rpname)
			return
		}
	}

	var a rlib.RatePlanRefRTRate
	var err error
	var ok bool

	//-------------------------------------------------------------------
	// RPRef
	//-------------------------------------------------------------------
	a.RPRID = CSVLoaderGetRPRID(strings.TrimSpace(sa[2]))
	if 0 == a.RPRID {
		rlib.Ulog("%s: line %d - Bad value for RatePlanRef ID: %s\n", funcname, lineno, sa[2])
		return
	}

	//-------------------------------------------------------------------
	// RT Style
	// identifies the RentableType
	//-------------------------------------------------------------------
	name := strings.TrimSpace(sa[3])
	rt, err := rlib.GetRentableTypeByStyle(name, b.BID)
	if err != nil {
		fmt.Printf("%s: line %d - could not load RentableType with Style = %s,  err:  %s\n", funcname, lineno, sa[3], err.Error())
		return
	}
	a.RTID = rt.RTID

	//-------------------------------------------------------------------
	// Amount
	// Entered as a string. If the string contains a % then the amount
	// is a percentage and we set the % flag. Otherwise, it is an absolute amount
	//-------------------------------------------------------------------
	amt := strings.TrimSpace(sa[4])
	a.Val, ok = rlib.FloatFromString(amt, "bad amount")
	if !ok {
		return
	}
	if strings.Contains(amt, "%") {
		a.FLAGS |= rlib.FlRTRpct
	}

	//-------------------------------------------------------------------
	// Insert the record
	//-------------------------------------------------------------------
	err = rlib.InsertRatePlanRefRTRate(&a)
	if nil != err {
		fmt.Printf("%s: lineno %d  - error inserting RatePlanRefRTRate = %v\n", funcname, lineno, err)
	}
}

// LoadRatePlanRefRTRatesCSV loads a csv file with RatePlan rates for specific rentable types
func LoadRatePlanRefRTRatesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRatePlanRefRTRate(t[i], i+1)
	}
}
