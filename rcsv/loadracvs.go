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

// BuildPayorList takes a semi-colon separated list of email addresses and phone numbers
// and returns an array of rlib.RentalAgreementPayor records for each.  If any of the addresses in the list
// cannot be resolved to a rlib.Transactant, then processing stops immediately and an error is returned.
// Each value is time sensitive (has an associated time range). If the dates are not specified, then the
// default values of dfltStart and dfltStop -- which are the start/stop time of the rental agreement --
// are used instead. This is common because the payors will usually be the same for the entire rental
// agreement lifetime.
func BuildPayorList(s string, dfltStart, dfltStop string, funcname string, lineno int) ([]rlib.RentalAgreementPayor, error) {
	var m []rlib.RentalAgreementPayor
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
		// PAYOR (rlib.Transactant)
		s = strings.TrimSpace(ss[0]) // either the email address or the phone number or TransactantID (TC0003234)

		//t := rlib.GetTransactantByPhoneOrEmail(s)
		n, _ := CSVLoaderTransactantList(s)
		if len(n) == 0 {
			rerr := fmt.Errorf("%s:  lineno %d - could not find rlib.Transactant with contact information %s\n", funcname, lineno, s)
			rlib.Ulog("%s", rerr.Error())
			return m, rerr
		}

		var payor rlib.RentalAgreementPayor
		payor.TCID = n[0].TCID

		// Now grab the dates
		if len(strings.TrimSpace(ss[1])) == 0 {
			ss[1] = dfltStart
		}
		if len(strings.TrimSpace(ss[2])) == 0 {
			ss[2] = dfltStop
		}
		payor.DtStart, payor.DtStop, _ = readTwoDates(ss[1], ss[2], funcname, lineno)

		m = append(m, payor)
	}
	return m, noerr
}

// BuildUserList parses a UserSpec and returns an array of RentableUser structs
func BuildUserList(sa, dfltStart, dfltStop string, funcname string, lineno int) ([]rlib.RentableUser, error) {
	s2 := strings.TrimSpace(sa) // TCID, email address, or the phone number
	s1 := strings.Split(s2, ";")
	var m []rlib.RentableUser
	var noerr error
	for i := 0; i < len(s1); i++ {
		ss := strings.Split(s1[i], ",")
		if len(ss) != 3 {
			err := fmt.Errorf("%s: lineno %d - invalid Status syntax. Each semi-colon separated field must have 3 values. Found %d in \"%s\"\n",
				funcname, lineno, len(ss), ss)
			return m, err
		}
		s := strings.TrimSpace(ss[0]) // TCID, email address, or the phone number
		n, err := CSVLoaderTransactantList(s)
		if err != nil {
			err := fmt.Errorf("%s: lineno %d - invalid person identifier: %s. Error = %s\n", funcname, lineno, s, err.Error())
			return m, err
		}
		var p rlib.RentableUser
		p.TCID = n[0].TCID

		if len(strings.TrimSpace(ss[1])) == 0 {
			ss[1] = dfltStart
		}
		if len(strings.TrimSpace(ss[2])) == 0 {
			ss[2] = dfltStop
		}
		p.DtStart, p.DtStop, _ = readTwoDates(ss[1], ss[2], funcname, lineno)
		m = append(m, p)
	}
	return m, noerr
}

// CreateRentalAgreement creates database records for the rental agreement defined in sa[]
func CreateRentalAgreement(sa []string, lineno int) int {
	funcname := "CreateRentalAgreement"
	var ra rlib.RentalAgreement
	var m []rlib.RentalAgreementRentable

	const (
		BUD               = 0
		RATemplateName    = iota
		AgreementStart    = iota
		AgreementStop     = iota
		PossessionStart   = iota
		PossessionStop    = iota
		RentStart         = iota
		RentStop          = iota
		RentCycleEpoch    = iota
		PayorSpec         = iota
		UserSpec          = iota
		Renewal           = iota
		SpecialProvisions = iota
		RentableSpec      = iota
		Notes             = iota
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
		{"Renewal", Renewal},
		{"SpecialProvisions", SpecialProvisions},
		{"RentableSpec", RentableSpec},
		{"Notes", Notes},
	}

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return 1
	}
	if lineno == 1 {
		return 0
	}

	//-------------------------------------------------------------------
	// BUD
	//-------------------------------------------------------------------
	cmpdes := strings.TrimSpace(sa[BUD])
	if len(cmpdes) > 0 {
		b2 := rlib.GetBusinessByDesignation(cmpdes)
		if b2.BID == 0 {
			fmt.Printf("%s: line %d - could not find rlib.Business named %s\n", funcname, lineno, cmpdes)
			return CsvErrorSensitivity
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
			rlib.Ulog("%s: line %d - rlib.Business with designation %s does not exist\n", funcname, lineno, sa[RATemplateName])
			return CsvErrorSensitivity
		}
		ra.RATID = b1.RATID
	}

	//-------------------------------------------------------------------
	// AgreementStartDate
	//-------------------------------------------------------------------
	dfltStart := sa[AgreementStart]
	DtStart, err := rlib.StringToDate(dfltStart)
	if err != nil {
		fmt.Printf("%s: line %d - invalid agreement start date:  %s\n", funcname, lineno, sa[AgreementStart])
		return CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// AgreementStopDate
	//-------------------------------------------------------------------
	dfltStop := sa[AgreementStop]
	DtStop, err := rlib.StringToDate(dfltStop)
	if err != nil {
		fmt.Printf("%s: line %d - invalid agreement stop date:  %s\n", funcname, lineno, sa[AgreementStop])
		return CsvErrorSensitivity
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
			fmt.Printf("%s: line %d - invalid possession start date:  %s\n", funcname, lineno, sa[PossessionStart])
			return CsvErrorSensitivity
		}
	}
	if len(sa[PossessionStop]) > 0 {
		ra.PossessionStop, err = rlib.StringToDate(sa[PossessionStop])
		if err != nil {
			fmt.Printf("%s: line %d - invalid possession stop date:  %s\n", funcname, lineno, sa[PossessionStop])
			return CsvErrorSensitivity
		}
	}
	if len(sa[RentStart]) > 0 {
		ra.RentStart, err = rlib.StringToDate(sa[RentStart])
		if err != nil {
			fmt.Printf("%s: line %d - invalid Rent start date:  %s\n", funcname, lineno, sa[RentStart])
			return CsvErrorSensitivity
		}
	}
	if len(sa[RentStop]) > 0 {
		ra.RentStop, err = rlib.StringToDate(sa[RentStop])
		if err != nil {
			fmt.Printf("%s: line %d - invalid Rent stop date:  %s\n", funcname, lineno, sa[RentStop])
			return CsvErrorSensitivity
		}
	}
	if len(sa[RentCycleEpoch]) > 0 {
		ra.RentCycleEpoch, err = rlib.StringToDate(sa[RentCycleEpoch])
		if err != nil {
			fmt.Printf("%s: line %d - invalid Rent cycle epoch date:  %s\n", funcname, lineno, sa[RentCycleEpoch])
			return CsvErrorSensitivity
		}
	}

	//-------------------------------------------------------------------
	//  The Payors
	//-------------------------------------------------------------------
	payors, err := BuildPayorList(sa[PayorSpec], dfltStart, dfltStop, funcname, lineno)
	if err != nil { // save the full list
		fmt.Printf("%s", err.Error())
		return CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	//  The Users
	//-------------------------------------------------------------------
	users, err := BuildUserList(sa[UserSpec], dfltStart, dfltStop, funcname, lineno)
	if err != nil { // save the full list
		fmt.Printf("%s", err.Error())
		return CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Renewal
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[Renewal])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("%s: line %d - Renewal value is invalid: %s\n", funcname, lineno, s)
			return CsvErrorSensitivity
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
			fmt.Printf("%s: line %d - Badly formatted string: %s . Format for each semicolon delimited part must be RentableName,ContractRent\n", funcname, lineno, ss)
			return CsvErrorSensitivity

		}
		var rar rlib.RentalAgreementRentable
		rnt, err := rlib.GetRentableByName(sss[0], ra.BID)
		if err != nil {
			fmt.Printf("%s: line %d - Could not load rentable named: %s  err = %s\n", funcname, lineno, sss[0], err.Error())
			return CsvErrorSensitivity
		}
		x, err := strconv.ParseFloat(strings.TrimSpace(sss[1]), 64)
		if err != nil {
			rlib.Ulog("%s: line %d - Invalid amount:  %s\n", funcname, lineno, sss[1])
			return CsvErrorSensitivity
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
			fmt.Printf("%s: line %d - error creating NoteList = %s\n", funcname, lineno, err.Error())
			return CsvErrorSensitivity
		}
		var n rlib.Note
		n.Comment = note
		n.NTID = 1 // first comment type
		n.NLID = nl.NLID
		_, err = rlib.InsertNote(&n)
		if err != nil {
			fmt.Printf("%s: line %d - error creating NoteList = %s\n", funcname, lineno, err.Error())
			return CsvErrorSensitivity
		}
		ra.NLID = nl.NLID
	}

	//------------------------------------
	// Write the rental agreement record
	//-----------------------------------
	RAID, err := rlib.InsertRentalAgreement(&ra)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.RentalAgreement = %v\n", funcname, lineno, err)
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
	return 0
}

// LoadRentalAgreementCSV loads a csv file with rental specialty types and processes each one
func LoadRentalAgreementCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if t[i][0] == "#" {
			continue
		}
		if CreateRentalAgreement(t[i], i+1) > 0 {
			return
		}
	}
}
