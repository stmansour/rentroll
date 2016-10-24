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
func CreateRatePlanRef(sa []string, lineno int) (string, int) {
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

	rs, x := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if x > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	des := strings.ToLower(strings.TrimSpace(sa[BUD]))

	//-------------------------------------------------------------------
	// BUD
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b = rlib.GetBusinessByDesignation(des)
		if len(b.Designation) == 0 {
			rs += fmt.Sprintf("%s: line %d, rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
			return rs, CsvErrorSensitivity
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
			rs += fmt.Sprintf("%s: line %d - RatePlan named %s not found\n", funcname, lineno, rpname)
			return rs, CsvErrorSensitivity
		}
	}

	var a rlib.RatePlanRef
	var err error
	var ok bool

	//-------------------------------------------------------------------
	// DtStart
	//-------------------------------------------------------------------
	dt := sa[DtStart]
	a.DtStart, err = rlib.StringToDate(dt)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[DtStart])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// DtStop
	//-------------------------------------------------------------------
	dt = sa[DtStop]
	a.DtStop, err = rlib.StringToDate(dt)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[DtStop])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Fee Applies Age
	//-------------------------------------------------------------------
	a.FeeAppliesAge, ok = rlib.IntFromString(sa[FeeAppliesAge], "Invalid FeeAppliesAge")
	if !ok {
		rs += fmt.Sprintf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[FeeAppliesAge])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Max No Fee Users
	//-------------------------------------------------------------------
	a.MaxNoFeeUsers, ok = rlib.IntFromString(sa[MaxNoFeeUsers], "Invalid MaxNoFeeUsers")
	if !ok {
		rs += fmt.Sprintf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[MaxNoFeeUsers])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// AdditionalUserFee
	//-------------------------------------------------------------------
	a.AdditionalUserFee, ok = rlib.FloatFromString(sa[AdditionalUserFee], "Invalid Additional User Fee")
	if !ok {
		rs += fmt.Sprintf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[AdditionalUserFee])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// CancellationFee
	//-------------------------------------------------------------------
	a.CancellationFee, ok = rlib.FloatFromString(sa[CancellationFee], "Invalid Cancellation Fee")
	if !ok {
		rs += fmt.Sprintf("%s: lineno %d  -  Invalid number: %s\n", funcname, lineno, sa[CancellationFee])
		return rs, CsvErrorSensitivity
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
				rs += fmt.Sprintf("%s: line %d - Unrecognized export flag: %s\n", funcname, lineno, ssa[i])
				return rs, CsvErrorSensitivity
			}
		}
	}

	//-------------------------------------------------------------------
	// Insert the record
	//-------------------------------------------------------------------
	a.RPID = rp.RPID
	_, err = rlib.InsertRatePlanRef(&a)
	if nil != err {
		rs += fmt.Sprintf("%s: lineno %d  - error inserting RatePlanRef = %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}
	return rs, 0
}

// LoadRatePlanRefsCSV loads a csv file with rental specialty types and processes each one
func LoadRatePlanRefsCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateRatePlanRef(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
