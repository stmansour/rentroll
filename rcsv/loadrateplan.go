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
func CreateRatePlans(sa []string, lineno int) (string, int) {
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

	rs, x := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if x > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[BUD]))
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rs += fmt.Sprintf("%s: line %d, rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
			return rs, CsvErrorSensitivity
		}
		rp.BID = b1.BID
	}
	rp.Name = strings.TrimSpace(sa[1])
	if len(rp.Name) == 0 {
		rs += fmt.Sprintf("%s: line %d - No Name found for the RatePlan\n", funcname, lineno)
		return rs, CsvErrorSensitivity
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
				rs += fmt.Sprintf("%s: line %d - Unrecognized export flag: %s\n", funcname, lineno, ssa[i])
				return rs, CsvErrorSensitivity
			}
		}
	}

	//rs += fmt.Sprintf("FLAGS = 0x%x\n", FLAGS)

	rpid, err := rlib.InsertRatePlan(&rp)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - Error inserting RatePlan.  err = %s\n", funcname, lineno, err.Error())
		return rs, CsvErrorSensitivity
	}

	// Now add the FLAGS as a custom attribute to the RatePlan
	var c rlib.CustomAttribute     // This is the custom attribute
	var cr rlib.CustomAttributeRef // This is the reference that binds it to an object
	c.Name = "FLAGS"
	c.Type = rlib.CUSTUINT
	c.Value = fmt.Sprintf("%d", FLAGS)
	cid, err := rlib.InsertCustomAttribute(&c)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - Could not insert CustomAttribute. err = %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}
	cr.ElementType = rlib.ELEMRATEPLAN
	cr.ID = rpid
	cr.CID = cid
	err = rlib.InsertCustomAttributeRef(&cr)
	if err != nil {
		rs += fmt.Sprintf("%s: line %d - Could not insert CustomAttributeRef. err = %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}
	return rs, 0
}

// LoadRatePlansCSV loads a csv file with note types
func LoadRatePlansCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateRatePlans(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
