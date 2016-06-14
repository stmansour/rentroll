package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

//  CSV file format:
//  0                     1      2             3             4                                            5        6                  7             ... 8, 9, ... as many as needed
//  RentalTemplateNumber, BUD,   RentalStart,  RentalStop,   Payor,                                       Renewal, SpecialProvisions, RentableName, ...
// 	"RAT001",             REH,   "2004-01-01", "2015-11-08", "866-123-4567,dtStart,dtStop;bill@x.com...", 1,       "",                101
// 	"RAT001",             REH,   "2004-01-01", "2017-07-04", "866-123-4567,dtStart,dtStop;bill@x.com",    1,       "",                107
// 	"RAT001",             REH,   "2015-11-21", "2016-11-21", "866-123-4567,,;bill@x.com,,",               1,       "",                101,102

// BuildPeopleList takes a semi-colon separated list of email addresses and phone numbers
// and returns an array of RentalAgreementPayor records for each.  If any of the addresses in the list
// cannot be resolved to a Transactant, then processing stops immediately and an error is returned.
// Each value is time sensitive (has an associated time range). If the dates are not specified, then the
// default values of dfltStart and dfltStop -- which are the start/stop time of the rental agreement --
// are used instead. This is common because the payors will usually be the same for the entire rental
// agreement lifetime.
func BuildPeopleList(s string, dfltStart, dfltStop string, funcname string, lineno int) ([]RentalAgreementPayor, error) {
	var m []RentalAgreementPayor
	var noerr error
	s2 := strings.TrimSpace(s) // either the email address or the phone number
	s1 := strings.Split(s2, ";")
	for i := 0; i < len(s1); i++ {
		ss := strings.Split(s1[i], ",")
		if len(ss) != 3 {
			err := fmt.Errorf("%s: lineno %d - invalid Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
				funcname, lineno, len(ss), ss)
			return m, err
		}
		var payor RentalAgreementPayor

		// PAYOR (Transactant)
		s = strings.TrimSpace(ss[0]) // either the email address or the phone number
		t, err := GetTransactantByPhoneOrEmail(s)
		if err != nil && !IsSQLNoResultsError(err) {
			rerr := fmt.Errorf("%s:  lineno %d - error retrieving Transactant by phone or email: %v", funcname, lineno, err)
			Ulog("%s", rerr.Error())
			return m, rerr
		}
		if t.PID == 0 {
			rerr := fmt.Errorf("%s:  lineno %d - could not find Transactant with contact information %s\n", funcname, lineno, s)
			Ulog("%s", rerr.Error())
			return m, rerr
		}
		payor.PID = t.PID

		// Now grab the dates
		if len(strings.TrimSpace(ss[1])) == 0 {
			ss[1] = dfltStart
		}
		if len(strings.TrimSpace(ss[2])) == 0 {
			ss[2] = dfltStop
		}
		payor.DtStart, payor.DtStop, err = readTwoDates(ss, funcname, lineno)

		m = append(m, payor)
	}
	return m, noerr
}

// CreateRentalAgreement creates database records for the rental agreement defined in sa[]
func CreateRentalAgreement(sa []string, lineno int) {
	funcname := "CreateRentalAgreement"
	var ra RentalAgreement
	var m []RentalAgreementRentable

	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "rentaltemplatenumber" {
		return // this is just the column heading
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 8
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Make sure the Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetRentalAgreementByRentalTemplateNumber(des)
		if len(b1.RentalTemplateNumber) == 0 {
			Ulog("%s: line %d - Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		ra.RATID = b1.RATID
	}

	//-------------------------------------------------------------------
	// See if the biz exists, if so, set the BID
	//-------------------------------------------------------------------
	cmpdes := strings.TrimSpace(sa[1])
	if len(cmpdes) > 0 {
		b2, _ := GetBusinessByDesignation(cmpdes)
		if b2.BID == 0 {
			fmt.Printf("%s: line %d - could not find Business named %s\n", funcname, lineno, cmpdes)
			return
		}
		ra.BID = b2.BID
	}

	//-------------------------------------------------------------------
	// RentalStartDate
	//-------------------------------------------------------------------
	dfltStart := sa[2]
	DtStart, err := StringToDate(dfltStart)
	if err != nil {
		fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[2])
		return
	}
	ra.RentalStart = DtStart

	//-------------------------------------------------------------------
	// RentalStopDate
	//-------------------------------------------------------------------
	dfltStop := sa[3]
	DtStop, err := StringToDate(dfltStop)
	if err != nil {
		fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[3])
		return
	}
	ra.RentalStop = DtStop

	// Until we update with new params...
	ra.PossessionStart = ra.RentalStart
	ra.PossessionStop = ra.RentalStop

	//-------------------------------------------------------------------
	//  The Payors
	//-------------------------------------------------------------------
	payors, err := BuildPeopleList(sa[4], dfltStart, dfltStop, funcname, lineno)
	if err != nil { // save the full list
		fmt.Printf("%s", err.Error())
		return
	}

	//-------------------------------------------------------------------
	// Renewal
	//-------------------------------------------------------------------

	s := strings.TrimSpace(sa[5])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("%s: line %d - Renewal value is invalid: %s\n", funcname, lineno, s)
			return
		}
		ra.Renewal = int64(i)
	}

	//-------------------------------------------------------------------
	// Special Provisions
	//-------------------------------------------------------------------
	ra.SpecialProvisions = sa[6]

	//-------------------------------------------------------------------
	// Rentables  -- all remaining columns are rentables
	//-------------------------------------------------------------------
	for i := 7; i < len(sa); i++ {
		s = strings.TrimSpace(sa[i])
		r, _ := GetRentableByName(s, ra.BID)

		if len(r.Name) > 0 {
			var ar RentalAgreementRentable
			ar.RID = r.RID
			ar.DtStart = DtStart
			ar.DtStop = DtStop
			m = append(m, ar)
		}
	}

	//------------------------------------
	// Write the rental agreement record
	//-----------------------------------
	RAID, err := InsertRentalAgreement(&ra)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting RentalAgreement = %v\n", funcname, lineno, err)
	}

	//------------------------------
	// Add the rentables
	//------------------------------
	for i := 0; i < len(m); i++ {
		m[i].RAID = RAID
		InsertRentalAgreementRentable(&m[i])
	}

	//------------------------------
	// Add the payors
	//------------------------------
	for i := 0; i < len(payors); i++ {
		payors[i].RAID = RAID
		InsertRentalAgreementPayor(&payors[i])
	}
}

// LoadRentalAgreementCSV loads a csv file with rental specialty types and processes each one
func LoadRentalAgreementCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalAgreement(t[i], i+1)
	}
}
