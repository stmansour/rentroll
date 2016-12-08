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
func CreateRatePlanRef(sa []string, lineno int) (int, error) {
	funcname := "CreateRatePlanRef"
	var b rlib.Business

	const (
		BUD               = 0
		RPName            = iota
		DtStart           = iota
		DtStop            = iota
		FeeAppliesAge     = iota
		MaxNoFeeUsers     = iota
		AdditionalUserFee = iota
		CancellationFee   = iota
		PromoCode         = iota
		Flags             = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"RPName", RPName},
		{"DtStart", DtStart},
		{"DtStop", DtStop},
		{"FeeAppliesAge", FeeAppliesAge},
		{"MaxNoFeeUsers", MaxNoFeeUsers},
		{"AdditionalUserFee", AdditionalUserFee},
		{"CancellationFee", CancellationFee},
		{"PromoCode", PromoCode},
		{"Flags", Flags},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	des := strings.ToLower(strings.TrimSpace(sa[BUD]))

	//-------------------------------------------------------------------
	// BUD
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b = rlib.GetBusinessByDesignation(des)
		if len(b.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
		}
	}

	//-------------------------------------------------------------------
	// RatePlan Name
	//-------------------------------------------------------------------
	var rp rlib.RatePlan
	rpname := strings.ToLower(strings.TrimSpace(sa[RPName]))
	if len(rpname) > 0 {
		rlib.GetRatePlanByName(b.BID, rpname, &rp)
		if rp.RPID < 1 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - RatePlan named %s not found\n", funcname, lineno, rpname)
		}
	}

	var a rlib.RatePlanRef
	var errmsg string

	//-------------------------------------------------------------------
	// DtStart
	//-------------------------------------------------------------------
	dt := sa[DtStart]
	a.DtStart, err = rlib.StringToDate(dt)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[DtStart])
	}

	//-------------------------------------------------------------------
	// DtStop
	//-------------------------------------------------------------------
	dt = sa[DtStop]
	a.DtStop, err = rlib.StringToDate(dt)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[DtStop])
	}

	//-------------------------------------------------------------------
	// Fee Applies Age
	//-------------------------------------------------------------------
	a.FeeAppliesAge, errmsg = rlib.IntFromString(sa[FeeAppliesAge], "Invalid FeeAppliesAge")
	if len(errmsg) > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[FeeAppliesAge])
	}

	//-------------------------------------------------------------------
	// Max No Fee Users
	//-------------------------------------------------------------------
	a.MaxNoFeeUsers, errmsg = rlib.IntFromString(sa[MaxNoFeeUsers], "Invalid MaxNoFeeUsers")
	if len(errmsg) > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[MaxNoFeeUsers])
	}

	//-------------------------------------------------------------------
	// AdditionalUserFee
	//-------------------------------------------------------------------
	a.AdditionalUserFee, errmsg = rlib.FloatFromString(sa[AdditionalUserFee], "Invalid Additional User Fee")
	if len(errmsg) > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[AdditionalUserFee])
	}

	//-------------------------------------------------------------------
	// CancellationFee
	//-------------------------------------------------------------------
	a.CancellationFee, errmsg = rlib.FloatFromString(sa[CancellationFee], "Invalid Cancellation Fee")
	if len(errmsg) > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[CancellationFee])
	}

	//-------------------------------------------------------------------
	// PromoCode
	//-------------------------------------------------------------------
	a.PromoCode = strings.TrimSpace(sa[PromoCode])

	//-------------------------------------------------------------------
	// FLAGS
	//-------------------------------------------------------------------
	ss := strings.TrimSpace(sa[Flags])
	if len(ss) > 0 {
		ssa := strings.Split(ss, ",")
		for i := 0; i < len(ssa); i++ {
			switch strings.ToLower(ssa[i]) {
			case "hide":
				a.FLAGS |= rlib.FlRTRRefHide // do not show this rate plan to users
			default:
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Unrecognized export flag: %s\n", funcname, lineno, ssa[i])
			}
		}
	}

	//-------------------------------------------------------------------
	// Insert the record
	//-------------------------------------------------------------------
	a.RPID = rp.RPID
	_, err = rlib.InsertRatePlanRef(&a)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: lineno %d  - error inserting RatePlanRef = %v\n", funcname, lineno, err)
	}
	return 0, nil
}

// LoadRatePlanRefsCSV loads a csv file with rental specialty types and processes each one
func LoadRatePlanRefsCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateRatePlanRef)
}
