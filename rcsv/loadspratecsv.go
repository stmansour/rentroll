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
func CreateRatePlanRefSPRate(sa []string, lineno int) (int, error) {
	funcname := "CreateRatePlanRefSPRate"
	var b rlib.Business

	const (
		BUD          = 0
		RPName       = iota
		RPRID        = iota
		RentableType = iota
		Amount       = iota
	)

	required := 5
	if len(sa) < required {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - found %d values, there must be at least %d", funcname, lineno, len(sa), required)
	}

	//-------------------------------------------------------------------
	// BUD
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "bud" {
		return 0, nil // this is just the column heading
	}
	if len(des) > 0 {
		b = rlib.GetBusinessByDesignation(des)
		if len(b.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, Business with designation %s does not exist", funcname, lineno, sa[0])
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
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - RatePlan named %s not found", funcname, lineno, rpname)
		}
	}

	var a rlib.RatePlanRefSPRate
	var err error
	var errmsg string

	a.BID = b.BID

	//-------------------------------------------------------------------
	// RPRef
	//-------------------------------------------------------------------
	a.RPRID = CSVLoaderGetRPRID(strings.TrimSpace(sa[2]))
	if 0 == a.RPRID {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Bad value for RatePlanRef ID: %s", funcname, lineno, sa[2])
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
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find Specialty with name = %s", funcname, lineno, rtname)
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
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find Specialty with name = %s", funcname, lineno, name)
		}

		//-------------------------------------------------------------------
		// Amount
		// Entered as a string. If the string contains a % then the amount
		// is a percentage and we set the % flag. Otherwise, it is an
		// absolute amount
		//-------------------------------------------------------------------
		amt := strings.TrimSpace(sa[i+1])
		p.Val, errmsg = rlib.FloatFromString(amt, "bad amount")
		if len(errmsg) > 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d  - %s", funcname, lineno, errmsg)
		}
		if strings.Contains(amt, "%") {
			p.FLAGS |= rlib.FlSPRpct // mark it as a percentage
		}

		//-------------------------------------------------------------------
		// Insert the record
		//-------------------------------------------------------------------
		err = rlib.InsertRatePlanRefSPRate(&p)
		if nil != err {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d  - error inserting RatePlanRefSPRate = %v", funcname, lineno, err)
		}
	}
	return 0, nil
}

// LoadRatePlanRefSPRatesCSV loads a csv file with RatePlan rates for specific rentable types
func LoadRatePlanRefSPRatesCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateRatePlanRefSPRate)
}
