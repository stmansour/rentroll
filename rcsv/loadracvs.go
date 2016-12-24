package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

//  CSV file format:
//  0      1               2             3             4                                           5               6        7                  8                                                              9
//  BUD,   RATemplateName, AgreementStart,  AgreementStop,   Payor,dtStart,dtStop;...              UserSpec,       Renewal, SpecialProvisions, "RentableName1,ContractRent2;RentableName2,ContractName2;...", Notes
//  BUD,   RATemplateName, AgreementStart,  AgreementStop,   PayorSpec,                            Usr1,d1,d2;..., Renewal, SpecialProvisions, "RentableName1,ContractRent2;RentableName2,ContractName2;...", Notes
// 	REH,   "RAT001",       "2004-01-01", "2015-11-08", "866-123-4567,dtStart,dtStop;bill@x.com...",UserSpec,       1,       "",                “U101,2500.00;U102,2350.00”,
// 	REH,   "RAT001",       "2004-01-01", "2017-07-04", "866-123-4567,dtStart,dtStop;bill@x.com",   UserSpec,       1,       "",                “U101,2500.00;U102,2350.00”,
// 	REH,   "RAT001",       "2015-11-21", "2016-11-21", "866-123-4567,,;bill@x.com,,",              UserSpec,       1,       "",                “U101,2500.00;U102,2350.00”,

// CreateRentalAgreement creates database records for the rental agreement defined in sa[]
func CreateRentalAgreement(sa []string, lineno int) (int, error) {
	funcname := "CreateRentalAgreement"

	var ra rlib.RentalAgreement
	var m []rlib.RentalAgreementRentable

	const (
		BUD                 = 0
		RATemplateName      = iota
		AgreementStart      = iota
		AgreementStop       = iota
		PossessionStart     = iota
		PossessionStop      = iota
		RentStart           = iota
		RentStop            = iota
		RentCycleEpoch      = iota
		PayorSpec           = iota
		UserSpec            = iota
		UnspecifiedAdults   = iota
		UnspecifiedChildren = iota
		Renewal             = iota
		SpecialProvisions   = iota
		RentableSpec        = iota
		Notes               = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"RATemplateName", RATemplateName},
		{"AgreementStart", AgreementStart},
		{"AgreementStop", AgreementStop},
		{"PossessionStart", PossessionStart},
		{"PossessionStop", PossessionStop},
		{"RentStart", RentStart},
		{"RentStop", RentStop},
		{"RentCycleEpoch", RentCycleEpoch},
		{"PayorSpec", PayorSpec},
		{"UserSpec", UserSpec},
		{"UnspecifiedAdults", UnspecifiedAdults},
		{"UnspecifiedChildren", UnspecifiedChildren},
		{"Renewal", Renewal},
		{"SpecialProvisions", SpecialProvisions},
		{"RentableSpec", RentableSpec},
		{"Notes", Notes},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	//-------------------------------------------------------------------
	// BUD
	//-------------------------------------------------------------------
	cmpdes := strings.TrimSpace(sa[BUD])
	if len(cmpdes) > 0 {
		b2 := rlib.GetBusinessByDesignation(cmpdes)
		if b2.BID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find rlib.Business named %s\n", funcname, lineno, cmpdes)
		}
		ra.BID = b2.BID
	}

	//-------------------------------------------------------------------
	// Make sure the RentalTemplate exists
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[RATemplateName]))
	if len(des) > 0 {
		b1 := rlib.GetRentalAgreementByRATemplateName(des)
		if len(b1.RATemplateName) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - rlib.Business with designation %s does not exist\n", funcname, lineno, sa[RATemplateName])
		}
		ra.RATID = b1.RATID
	}

	//-------------------------------------------------------------------
	// AgreementStartDate
	//-------------------------------------------------------------------
	dfltStart := sa[AgreementStart]
	DtStart, err := rlib.StringToDate(dfltStart)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid agreement start date:  %s\n", funcname, lineno, sa[AgreementStart])
	}

	//-------------------------------------------------------------------
	// AgreementStopDate
	//-------------------------------------------------------------------
	dfltStop := sa[AgreementStop]
	DtStop, err := rlib.StringToDate(dfltStop)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid agreement stop date:  %s\n", funcname, lineno, sa[AgreementStop])
	}
	ra.AgreementStop = DtStop

	// Initialize to default values
	ra.AgreementStart = DtStart
	ra.PossessionStart = DtStart
	ra.RentStart = DtStart
	ra.RentCycleEpoch = DtStart
	ra.PossessionStop = ra.AgreementStop
	ra.RentStop = ra.AgreementStop

	if len(sa[PossessionStart]) > 0 {
		ra.PossessionStart, err = rlib.StringToDate(sa[PossessionStart])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid possession start date:  %s\n", funcname, lineno, sa[PossessionStart])
		}
	}
	if len(sa[PossessionStop]) > 0 {
		ra.PossessionStop, err = rlib.StringToDate(sa[PossessionStop])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid possession stop date:  %s\n", funcname, lineno, sa[PossessionStop])
		}
	}
	if len(sa[RentStart]) > 0 {
		ra.RentStart, err = rlib.StringToDate(sa[RentStart])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Rent start date:  %s\n", funcname, lineno, sa[RentStart])
		}
	}
	if len(sa[RentStop]) > 0 {
		ra.RentStop, err = rlib.StringToDate(sa[RentStop])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Rent stop date:  %s\n", funcname, lineno, sa[RentStop])
		}
	}
	if len(sa[RentCycleEpoch]) > 0 {
		ra.RentCycleEpoch, err = rlib.StringToDate(sa[RentCycleEpoch])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Rent cycle epoch date:  %s\n", funcname, lineno, sa[RentCycleEpoch])
		}
	}

	//-------------------------------------------------------------------
	//  The Payors
	//-------------------------------------------------------------------
	payors, err := BuildPayorList(ra.BID, sa[PayorSpec], dfltStart, dfltStop, funcname, lineno)
	if err != nil { // save the full list
		return CsvErrorSensitivity, err
	}

	//-------------------------------------------------------------------
	//  The Users
	//-------------------------------------------------------------------
	users, err := BuildUserList(ra.BID, sa[UserSpec], dfltStart, dfltStop, funcname, lineno)
	if err != nil { // save the full list
		return CsvErrorSensitivity, err
	}
	//-------------------------------------------------------------------
	//  The Unspecified Adults and Children
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[UnspecifiedAdults])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - UnspecifiedAdults value is invalid: %s\n", funcname, lineno, s)
		}
		ra.UnspecifiedAdults = int64(i)
	}
	s = strings.TrimSpace(sa[UnspecifiedChildren])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - UnspecifiedChildren value is invalid: %s\n", funcname, lineno, s)
		}
		ra.UnspecifiedChildren = int64(i)
	}

	//-------------------------------------------------------------------
	// Renewal
	//-------------------------------------------------------------------
	s = strings.TrimSpace(sa[Renewal])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Renewal value is invalid: %s\n", funcname, lineno, s)
		}
		ra.Renewal = int64(i)
	}

	//-------------------------------------------------------------------
	// Special Provisions
	//-------------------------------------------------------------------
	ra.SpecialProvisions = sa[SpecialProvisions]

	//-------------------------------------------------------------------
	// Rentables  -- all remaining columns are rentables
	//-------------------------------------------------------------------
	ss := strings.Split(sa[RentableSpec], ";")
	for i := 0; i < len(ss); i++ {
		sss := strings.Split(ss[i], ",")
		if len(sss) != 2 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Badly formatted string: %s . Format for each semicolon delimited part must be RentableName,ContractRent\n", funcname, lineno, ss)

		}
		var rar rlib.RentalAgreementRentable
		rnt, err := rlib.GetRentableByName(sss[0], ra.BID)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not load rentable named: %s  err = %s\n", funcname, lineno, sss[0], err.Error())
		}
		x, err := strconv.ParseFloat(strings.TrimSpace(sss[1]), 64)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid amount:  %s\n", funcname, lineno, sss[1])
		}
		rar.RID = rnt.RID
		rar.DtStart = DtStart
		rar.DtStop = DtStop
		rar.ContractRent = x
		m = append(m, rar)
	}

	//-------------------------------------------------------------------
	// Notes
	//-------------------------------------------------------------------
	note := strings.TrimSpace(sa[Notes])
	if len(note) > 0 {
		var nl rlib.NoteList
		nl.NLID, err = rlib.InsertNoteList(&nl)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error creating NoteList = %s\n", funcname, lineno, err.Error())
		}
		var n rlib.Note
		n.Comment = note
		n.NTID = 1 // first comment type
		n.NLID = nl.NLID
		_, err = rlib.InsertNote(&n)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error creating NoteList = %s\n", funcname, lineno, err.Error())
		}
		ra.NLID = nl.NLID
	}

	//-------------------------------------------------------------------
	// look for any rental agreements already in existence that cover
	// the rentables referenced in this one...
	//-------------------------------------------------------------------
	for i := 0; i < len(m); i++ {
		rra := rlib.GetAgreementsForRentable(m[i].RID, &ra.AgreementStart, &ra.AgreementStop)
		for j := 0; j < len(rra); j++ {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Rentable %s is already included in Rental Agreement %s from %s to %s\n",
				funcname, lineno,
				rlib.IDtoString("R", rra[j].RID), rlib.IDtoString("RA", rra[j].RAID),
				rra[j].DtStart.Format(rlib.RRDATEFMT4), rra[j].DtStop.Format(rlib.RRDATEFMT4))
		}
	}

	//-----------------------------------------------
	// Validate that we have at least one payor...
	//-----------------------------------------------
	if len(payors) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - There are no valid Payors for this Rental Agreement.\n", funcname, lineno)
	}

	//------------------------------------
	// Write the rental agreement record
	//-----------------------------------
	RAID, err := rlib.InsertRentalAgreement(&ra)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting rlib.RentalAgreement = %v\n", funcname, lineno, err)
	}

	//------------------------------------------------------------
	// Add the rentables, and the users of those rentables...
	//------------------------------------------------------------
	for i := 0; i < len(m); i++ {
		m[i].RAID = RAID
		rlib.InsertRentalAgreementRentable(&m[i])
		for j := 0; j < len(users); j++ {
			users[j].RID = m[i].RID
			rlib.InsertRentableUser(&users[j])
		}
	}

	//------------------------------
	// Add the payors
	//------------------------------
	for i := 0; i < len(payors); i++ {
		payors[i].RAID = RAID
		rlib.InsertRentalAgreementPayor(&payors[i])
	}
	return 0, nil
}

// LoadRentalAgreementCSV loads a csv file with rental specialty types and processes each one
func LoadRentalAgreementCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateRentalAgreement)
}
