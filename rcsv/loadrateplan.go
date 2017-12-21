package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//      RatePlan
// |<------|--------------->|
// 0    1        2
// BUD, Name,  Exports
// REX, X1,    GDS,Sabre
// REX, X2,

// CreateRatePlans reads a RatePlan string array and creates a database record
func CreateRatePlans(sa []string, lineno int) (int, error) {
	funcname := "CreateRatePlans"
	var rp rlib.RatePlan
	var FLAGS uint64

	const (
		BUD     = 0
		Name    = iota
		Exports = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"Exports", Exports},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[BUD]))
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, rlib.Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		rp.BID = b1.BID
	}
	rp.Name = strings.TrimSpace(sa[1])
	if len(rp.Name) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No Name found for the RatePlan", funcname, lineno)
	}
	// need to check for another RatePlan of the same name

	//-------------------------------------------------------------------
	// Exports
	//-------------------------------------------------------------------
	ss := strings.TrimSpace(sa[2])
	if len(ss) > 0 {
		ssa := strings.Split(ss, ",")
		for i := 0; i < len(ssa); i++ {
			switch ssa[i] {
			case "GDS":
				FLAGS |= rlib.FlRatePlanGDS
			case "Sabre":
				FLAGS |= rlib.FlRatePlanSabre
			default:
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Unrecognized export flag: %s", funcname, lineno, ssa[i])
			}
		}
	}

	//return CsvErrorSensitivity, fmt.Errorf("FLAGS = 0x%x", FLAGS)

	rpid, err := rlib.InsertRatePlan(&rp)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error inserting RatePlan.  err = %s", funcname, lineno, err.Error())
	}

	// Now add the FLAGS as a custom attribute to the RatePlan
	var c rlib.CustomAttribute     // This is the custom attribute
	var cr rlib.CustomAttributeRef // This is the reference that binds it to an object
	c.Name = "FLAGS"
	c.BID = rp.BID
	c.Type = rlib.CUSTUINT
	c.Value = fmt.Sprintf("%d", FLAGS)
	cid, err := rlib.InsertCustomAttribute(&c)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not insert CustomAttribute. err = %v", funcname, lineno, err)
	}
	cr.ElementType = rlib.ELEMRATEPLAN
	cr.ID = rpid
	cr.BID = rp.BID
	cr.CID = cid
	_, err = rlib.InsertCustomAttributeRef(&cr)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not insert CustomAttributeRef. err = %v", funcname, lineno, err)
	}
	return 0, nil
}

// LoadRatePlansCSV loads a csv file with note types
func LoadRatePlansCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateRatePlans)
}
