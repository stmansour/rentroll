package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//  CSV file format:
//
// 0    1       2        3            4
// BUD, RPName,          RPRID        Specialty,   Amount,  Specialty2, Amount2, ...
// REX, FAA-P,  RPR0001, GM,          Lake View,   85%,     Fireplace,      90%
// REX, FAA-P,  1,       Flat Studio, Lake View,   100%,    Fireplace,
// REX, FAA-P,  1,       SBL,    	  Lake View,   10.25,   Fireplace,
// REX, FAA-P,  1,       KDS,    	  Lake View,   75%,     Fireplace,
// REX, FAA-T,  2,       GM,          Lake View,   90%,     Fireplace,
// REX, FAA-T,  2,       Flat Studio, Lake View,   90%,     Fireplace,
// REX, FAA-T,  2,       SBL,    	  Lake View,   11.50,   Fireplace,
// REX, FAA-T,  2,       KDS,    	  Lake View,   87%,     Fireplace,

// CreateRatePlanRefSPRate reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreateRatePlanRefSPRate(sa []string, lineno int) {
	funcname := "CreateRatePlanRefSPRate"
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
			rlib.Ulog("%s: line %d, Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
	}

	// knowing the Business we can get all the specialties and rentable types. The easy way is just to load an XBiz
	var xbiz rlib.XBusiness
	rlib.GetXBusiness(b.BID, &xbiz)

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

	var a rlib.RatePlanRefSPRate
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
	// Rentable Type
	//-------------------------------------------------------------------
	rtname := strings.TrimSpace(sa[3])
	found := false
	for k, v := range xbiz.RT { // Make sure it's something we recognize...
		if v.Name == rtname || v.Style == rtname {
			found = true
			a.RTID = k // mark the RTID
			break
		}
	}
	if !found {
		fmt.Printf("%s: line %d - could not find Specialty with name = %s\n", funcname, lineno, rtname)
		return
	}

	for i := 4; i < len(sa); i += 2 {
		p := a // start a new structure.  We just need to fill out the RSPID, Amount, and FLAG

		//-------------------------------------------------------------------
		// Specialty
		//-------------------------------------------------------------------
		name := strings.TrimSpace(sa[i])
		if len(name) == 0 { // if the specialty name is blank...
			continue // ... then ignore
		}
		// Make sure it's something we recognize...
		found = false
		for k, v := range xbiz.US {
			if v.Name == name {
				found = true
				p.RSPID = k
				break
			}
		}
		if !found {
			fmt.Printf("%s: line %d - could not find Specialty with name = %s\n", funcname, lineno, name)
			return
		}

		//-------------------------------------------------------------------
		// Amount
		// Entered as a string. If the string contains a % then the amount
		// is a percentage and we set the % flag. Otherwise, it is an
		// absolute amount
		//-------------------------------------------------------------------
		amt := strings.TrimSpace(sa[i+1])
		p.Val, ok = rlib.FloatFromString(amt, "bad amount")
		if !ok {
			return
		}
		if strings.Contains(amt, "%") {
			p.FLAGS |= rlib.FlSPRpct // mark it as a percentage
		}

		//-------------------------------------------------------------------
		// Insert the record
		//-------------------------------------------------------------------
		err = rlib.InsertRatePlanRefSPRate(&p)
		if nil != err {
			fmt.Printf("%s: lineno %d  - error inserting RatePlanRefSPRate = %v\n", funcname, lineno, err)
		}
	}
}

// LoadRatePlanRefSPRatesCSV loads a csv file with RatePlan rates for specific rentable types
func LoadRatePlanRefSPRatesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRatePlanRefSPRate(t[i], i+1)
	}
}
