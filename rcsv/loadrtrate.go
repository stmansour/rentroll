package rcsv

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"strings"
)

//  CSV file format:
//                        RT name or style  string with or without %
// 0    1             2   3                 4
// BUD, RPName, RPRID     RentableType,     Amount
// REX, FAA-P,  RPR0001,  GM,               85%
// REX, FAA-P,  1,        Flat Studio,      1400
// REX, FAA-P,  1,        SBL,    			1500
// REX, FAA-P,  1,        KDS,    			75%
// REX, FAA-T,  1,        GM,               90%
// REX, FAA-T,  1,        Flat Studio,      90%
// REX, FAA-T,  1,        SBL,    			1500
// REX, FAA-T,  1,        KDS,    			80%

// CreateRatePlanRefRTRate reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRatePlanRefRTRate(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateRatePlanRefRTRate"
	var (
		b   rlib.Business
		err error
	)

	const (
		BUD          = 0
		RPName       = iota
		RPRID        = iota
		RentableType = iota
		Amount       = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"RPName", RPName},
		{"RPRID", RPRID},
		{"RentableType", RentableType},
		{"Amount", Amount},
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
		b, err = rlib.GetBusinessByDesignation(ctx, des)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, error while getting business by designation(%s): %s", funcname, lineno, des, err.Error())
		}
		if len(b.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, rlib.Business with designation %s does not exist", funcname, lineno, sa[BUD])
		}
	}

	//-------------------------------------------------------------------
	// RatePlan Name
	//-------------------------------------------------------------------
	var rp rlib.RatePlan
	rpname := strings.ToLower(strings.TrimSpace(sa[RPName]))
	if len(rpname) > 0 {
		err = rlib.GetRatePlanByName(ctx, b.BID, rpname, &rp)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error getting RatePlan named %s not found: %s", funcname, lineno, rpname, err.Error())
		}
		if rp.RPID < 1 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - RatePlan named %s not found", funcname, lineno, rpname)
		}
	}

	var a rlib.RatePlanRefRTRate
	var errmsg string

	a.BID = b.BID

	//-------------------------------------------------------------------
	// RPRef
	//-------------------------------------------------------------------
	a.RPRID = CSVLoaderGetRPRID(strings.TrimSpace(sa[RPRID]))
	if 0 == a.RPRID {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Bad value for RatePlanRef ID: %s", funcname, lineno, sa[RPRID])
	}

	//-------------------------------------------------------------------
	// RT Style
	// identifies the RentableType
	//-------------------------------------------------------------------
	name := strings.TrimSpace(sa[RentableType])
	rt, err := rlib.GetRentableTypeByStyle(ctx, name, b.BID)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not load RentableType with Style = %s,  err:  %s", funcname, lineno, sa[RentableType], err.Error())
	}
	a.RTID = rt.RTID

	//-------------------------------------------------------------------
	// Amount
	// Entered as a string. If the string contains a % then the amount
	// is a percentage and we set the % flag. Otherwise, it is an absolute amount
	//-------------------------------------------------------------------
	amt := strings.TrimSpace(sa[Amount])
	a.Val, errmsg = rlib.FloatFromString(amt, "bad amount")
	if len(errmsg) > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, errmsg)
	}
	if strings.Contains(amt, "%") {
		a.FLAGS |= rlib.FlRTRpct
	}

	//-------------------------------------------------------------------
	// Insert the record
	//-------------------------------------------------------------------
	_, err = rlib.InsertRatePlanRefRTRate(ctx, &a)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d  - error inserting RatePlanRefRTRate = %v", funcname, lineno, err)
	}
	return 0, nil
}

// LoadRatePlanRefRTRatesCSV loads a csv file with RatePlan rates for specific rentable types
func LoadRatePlanRefRTRatesCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateRatePlanRefRTRate)

}
