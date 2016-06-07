package rlib

import (
	"fmt"
	"strconv"
	"strings"
)

//  CSV file format:
//                                                                                   |<----- repeat Rentable name, as many as needed ... -->|
//        0      1        2        3        4           5        6            7            8
//  TemplateName,BID,Renter,Payor,RentalStart,RentalStop,Renewal,SpecialProvisions,RentableName, ...
// 		"RAT001","REH","866-123-4567","866-123-4567","2004-01-01","2015-11-08",1,"",101
// 		"RAT001","REH","866-123-4567","866-123-4567","2004-01-01","2017-07-04",1,"",107
// 		"RAT001","REH","homerj@springfield.com","866-123-4567","2015-11-21","2016-11-21",1,"",101,102

// BuildPeopleList takes a semi-colon separated list of email addresses and phone numbers
// and returns an array of transactant records for each.  If any of the addresses in the list
// cannot be resolved to a transactant, then processing stops immediately and an error is returned.
func BuildPeopleList(s string) ([]Transactant, error) {
	var m []Transactant
	var noerr error
	s2 := strings.TrimSpace(s) // either the email address or the phone number
	s1 := strings.Split(s2, ";")
	for i := 0; i < len(s1); i++ {
		s = strings.TrimSpace(s1[i]) // either the email address or the phone number
		t, err := GetTransactantByPhoneOrEmail(s)
		if err != nil && !IsSQLNoResultsError(err) {
			rerr := fmt.Errorf("CreateRentalAgreement: error retrieving transactant by phone or email: %v", err)
			Ulog("%s", rerr.Error())
			return m, rerr
		}
		if t.PID == 0 {
			rerr := fmt.Errorf("CreateRentalAgreement: could not find transactant with contact information %s\n", s)
			Ulog("%s", rerr.Error())
			return m, rerr
		}
		m = append(m, t)
	}
	return m, noerr
}

// CreateRentalAgreement creates database records for the rental agreement defined in sa[]
func CreateRentalAgreement(sa []string) {
	var ra RentalAgreement
	var payor AgreementPayor
	var m []AgreementRentable

	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "templatename" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Make sure the business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetRentalAgreementTemplateByRefNum(des)
		if len(b1.RentalTemplateNumber) == 0 {
			Ulog("CreateRentalAgreement: business with designation %s does net exist\n", sa[0])
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
			fmt.Printf("CreateRentalAgreement: could not find business named %s\n", cmpdes)
			return
		}
		ra.BID = b2.BID
	}

	//-------------------------------------------------------------------
	//  Determine the primary renter
	//-------------------------------------------------------------------
	renters, err := BuildPeopleList(sa[2])
	if err != nil { // save the full list
		return
	}

	//-------------------------------------------------------------------
	//  Determine the payor
	//-------------------------------------------------------------------
	payors, err := BuildPeopleList(sa[3])
	if err != nil { // save the full list
		return
	}
	if len(payors) > 0 {
		payor.PID = payors[0].PID // store the primary payor now, we'll update the agreement payors later
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	DtStart, err := StringToDate(sa[4])
	if err != nil {
		fmt.Printf("CreateRentalAgreement: invalid start date:  %s\n", sa[4])
		return
	}
	ra.RentalStart = DtStart

	DtStop, err := StringToDate(sa[5])
	if err != nil {
		fmt.Printf("CreateRentalAgreement: invalid stop date:  %s\n", sa[5])
		return
	}
	ra.RentalStop = DtStop

	// Until we update with new params...
	ra.PossessionStart = ra.RentalStart
	ra.PossessionStop = ra.RentalStop

	s := strings.TrimSpace(sa[6])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("CreatePeopleFromCSV: Renewal value is invalid: %s\n", s)
			return
		}
		ra.Renewal = int64(i)
	}

	ra.SpecialProvisions = sa[7]

	// the rest of the arguments are rentables that are associated with
	// this rental agreement
	for i := 8; i < len(sa); i++ {
		s = strings.TrimSpace(sa[i])
		r, _ := GetRentableByName(s, ra.BID)

		if len(r.Name) > 0 {
			var ar AgreementRentable
			ar.RID = r.RID
			ar.DtStart = DtStart
			ar.DtStop = DtStop
			m = append(m, ar)
		}
	}

	// First write the rental agreement record, then write the agreementrentables and agreement payors
	RAID, err := InsertRentalAgreement(&ra)
	if nil != err {
		fmt.Printf("CreateRentalAgreement: error inserting RentalAgreement = %v\n", err)
	}
	for i := 0; i < len(m); i++ {
		m[i].RAID = RAID
		InsertAgreementRentable(&m[i])
	}

	payor.RAID = RAID
	payor.DtStart = DtStart
	payor.DtStop = DtStop

	var at AgreementRenter
	at.DtStart = DtStart
	at.DtStop = DtStop
	at.RAID = RAID

	//==================================================
	// Now handle payors and renters...
	//==================================================
	for i := 0; i < len(renters); i++ {
		at.RENTERID = renters[i].RENTERID
		InsertAgreementRenter(&at)
	}
	for i := 0; i < len(payors); i++ {
		payor.PID = payors[i].PID
		InsertAgreementPayor(&payor)
	}
}

// LoadRentalAgreementCSV loads a csv file with rental specialty types and processes each one
func LoadRentalAgreementCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalAgreement(t[i])
	}
}
