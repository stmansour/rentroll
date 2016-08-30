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
func CreateRatePlans(sa []string, lineno int) {
	funcname := "CreateRatePlans"
	var rp rlib.RatePlan
	var FLAGS uint64
	required := 3
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "bud" {
		return // this is just the column heading
	}
	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d, rlib.Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		rp.BID = b1.BID
	}
	rp.Name = strings.TrimSpace(sa[1])
	if len(rp.Name) == 0 {
		fmt.Printf("%s: line %d - No Name found for the RatePlan\n", funcname, lineno)
		return
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
				fmt.Printf("%s: line %d - Unrecognized export flag: %s\n", funcname, lineno, ssa[i])
				return
			}
		}
	}

	//rlib.Ulog("FLAGS = 0x%x\n", FLAGS)

	rpid, err := rlib.InsertRatePlan(&rp)
	if err != nil {
		fmt.Printf("%s: line %d - Error inserting RatePlan.  err = %s\n", funcname, lineno, err.Error())
	}

	// Now add the FLAGS as a custom attribute to the RatePlan
	var c rlib.CustomAttribute     // This is the custom attribute
	var cr rlib.CustomAttributeRef // This is the reference that binds it to an object
	c.Name = "FLAGS"
	c.Type = rlib.CUSTUINT
	c.Value = fmt.Sprintf("%d", FLAGS)
	cid, err := rlib.InsertCustomAttribute(&c)
	if err != nil {
		fmt.Printf("%s: line %d - Could not insert CustomAttribute. err = %v\n", funcname, lineno, err)
	}
	cr.ElementType = rlib.ELEMRATEPLAN
	cr.ID = rpid
	cr.CID = cid
	err = rlib.InsertCustomAttributeRef(&cr)
	if err != nil {
		fmt.Printf("%s: line %d - Could not insert CustomAttributeRef. err = %v\n", funcname, lineno, err)
	}

}

// LoadRatePlansCSV loads a csv file with note types
func LoadRatePlansCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRatePlans(t[i], i+1)
	}
}
