package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//  CSV file format:
// 0    1             2         3         4              5              6                   7                8          9
// BUD, RPName,       DtStart,  DtStop,   FeeAppliesAge, MaxNoFeeUsers, AdditionalUserFee,  CancellationFee, PromoCode, Flags
// REX, A1-Transient, 1/1/2016, 7/1/2016, 12,            2,             10.0,               25.00,                    , Hide
// REX, A1-LongTerm,  1/1/2016, 7/1/2016, 12,            2,             50.0,               25.00,                    , Hide
// REX, A2-Transient, 1/1/2016, 7/1/2016, 12,            2,             15.0,               25.00,                    ,
// REX, A2-LongTerm,  1/1/2016, 7/1/2016, 12,            2,             75.0,               25.00,                    ,

// CreateRatePlanRef reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRatePlanRef(sa []string, lineno int) {
	funcname := "CreateRatePlanRef"
	var b rlib.Business

	required := 10
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

	var a rlib.RatePlanRef
	var err error
	var ok bool

	//-------------------------------------------------------------------
	// DtStart
	//-------------------------------------------------------------------
	dt := sa[2]
	a.DtStart, err = rlib.StringToDate(dt)
	if err != nil {
		fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[2])
		return
	}

	//-------------------------------------------------------------------
	// DtStop
	//-------------------------------------------------------------------
	dt = sa[3]
	a.DtStop, err = rlib.StringToDate(dt)
	if err != nil {
		fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[3])
		return
	}

	//-------------------------------------------------------------------
	// Fee Applies Age
	//-------------------------------------------------------------------
	a.FeeAppliesAge, ok = rlib.IntFromString(sa[4], "Invalid FeeAppliesAge")
	if !ok {
		fmt.Printf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[4])
		return
	}

	//-------------------------------------------------------------------
	// Max No Fee Users
	//-------------------------------------------------------------------
	a.MaxNoFeeUsers, ok = rlib.IntFromString(sa[5], "Invalid MaxNoFeeUsers")
	if !ok {
		fmt.Printf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[5])
		return
	}

	//-------------------------------------------------------------------
	// AdditionalUserFee
	//-------------------------------------------------------------------
	a.AdditionalUserFee, ok = rlib.FloatFromString(sa[6], "Invalid Additional User Fee")
	if !ok {
		fmt.Printf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[6])
		return
	}

	//-------------------------------------------------------------------
	// CancellationFee
	//-------------------------------------------------------------------
	a.CancellationFee, ok = rlib.FloatFromString(sa[7], "Invalid Additional User Fee")
	if !ok {
		fmt.Printf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[7])
		return
	}

	//-------------------------------------------------------------------
	// PromoCode
	//-------------------------------------------------------------------
	a.PromoCode = strings.TrimSpace(sa[8])

	//-------------------------------------------------------------------
	// FLAGS
	//-------------------------------------------------------------------
	ss := strings.TrimSpace(sa[9])
	if len(ss) > 0 {
		ssa := strings.Split(ss, ",")
		for i := 0; i < len(ssa); i++ {
			switch strings.ToLower(ssa[i]) {
			case "hide":
				a.FLAGS |= rlib.FlRTRRefHide // do not show this rate plan to users
			default:
				fmt.Printf("%s: line %d - Unrecognized export flag: %s\n", funcname, lineno, ssa[i])
				return
			}
		}
	}

	//-------------------------------------------------------------------
	// Insert the record
	//-------------------------------------------------------------------
	a.RPID = rp.RPID
	_, err = rlib.InsertRatePlanRef(&a)
	if nil != err {
		fmt.Printf("%s: lineno %d  - error inserting RatePlanRef = %v\n", funcname, lineno, err)
	}
}

// LoadRatePlanRefsCSV loads a csv file with rental specialty types and processes each one
func LoadRatePlanRefsCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRatePlanRef(t[i], i+1)
	}
}
