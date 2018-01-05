package rcsv

import (
	"context"
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
func CreateRentalAgreement(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateRentalAgreement"
	var (
		err error
		ra  rlib.RentalAgreement
		m   []rlib.RentalAgreementRentable
	)

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
		b2, err := rlib.GetBusinessByDesignation(ctx, cmpdes)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d, error while getting business by designation(%s): %s", funcname, lineno, cmpdes, err.Error())
		}
		if b2.BID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not find rlib.Business named %s", funcname, lineno, cmpdes)
		}
		ra.BID = b2.BID
	}

	//-------------------------------------------------------------------
	// Make sure the RentalTemplate exists
	//-------------------------------------------------------------------
	des := strings.ToLower(strings.TrimSpace(sa[RATemplateName]))
	if len(des) > 0 {
		b1, err := rlib.GetRentalAgreementByRATemplateName(ctx, des)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error while getting ra template %s: %s", funcname, lineno, sa[RATemplateName], err.Error())
		}
		if len(b1.RATemplateName) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - ra template %s does not exist", funcname, lineno, sa[RATemplateName])
		}
		ra.RATID = b1.RATID
	}

	//-------------------------------------------------------------------
	// AgreementStartDate
	//-------------------------------------------------------------------
	dfltStart := sa[AgreementStart]
	DtStart, err := rlib.StringToDate(dfltStart)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid agreement start date:  %s", funcname, lineno, sa[AgreementStart])
	}

	//-------------------------------------------------------------------
	// AgreementStopDate
	//-------------------------------------------------------------------
	dfltStop := sa[AgreementStop]
	DtStop, err := rlib.StringToDate(dfltStop)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid agreement stop date:  %s", funcname, lineno, sa[AgreementStop])
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
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid possession start date:  %s", funcname, lineno, sa[PossessionStart])
		}
	}
	if len(sa[PossessionStop]) > 0 {
		ra.PossessionStop, err = rlib.StringToDate(sa[PossessionStop])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid possession stop date:  %s", funcname, lineno, sa[PossessionStop])
		}
	}
	if len(sa[RentStart]) > 0 {
		ra.RentStart, err = rlib.StringToDate(sa[RentStart])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Rent start date:  %s", funcname, lineno, sa[RentStart])
		}
	}
	if len(sa[RentStop]) > 0 {
		ra.RentStop, err = rlib.StringToDate(sa[RentStop])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Rent stop date:  %s", funcname, lineno, sa[RentStop])
		}
	}
	if len(sa[RentCycleEpoch]) > 0 {
		ra.RentCycleEpoch, err = rlib.StringToDate(sa[RentCycleEpoch])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Rent cycle epoch date:  %s", funcname, lineno, sa[RentCycleEpoch])
		}
	}

	//-------------------------------------------------------------------
	//  The Payors
	//-------------------------------------------------------------------
	payors, err := BuildPayorList(ctx, ra.BID, sa[PayorSpec], dfltStart, dfltStop, funcname, lineno)
	if err != nil { // save the full list
		return CsvErrorSensitivity, err
	}

	//-------------------------------------------------------------------
	//  The Users
	//-------------------------------------------------------------------
	users, err := BuildUserList(ctx, ra.BID, sa[UserSpec], dfltStart, dfltStop, funcname, lineno)
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
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - UnspecifiedAdults value is invalid: %s", funcname, lineno, s)
		}
		ra.UnspecifiedAdults = int64(i)
	}
	s = strings.TrimSpace(sa[UnspecifiedChildren])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - UnspecifiedChildren value is invalid: %s", funcname, lineno, s)
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
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Renewal value is invalid: %s", funcname, lineno, s)
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
	ss := strings.Split(strings.TrimSpace(sa[RentableSpec]), ";")
	if len(ss) > 0 && len(ss[0]) > 0 {
		for i := 0; i < len(ss); i++ {
			sss := strings.Split(ss[i], ",")
			if len(sss) != 2 {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Badly formatted string: %s . Format for each semicolon delimited part must be RentableName,ContractRent", funcname, lineno, ss)

			}
			var rar rlib.RentalAgreementRentable
			rnt, err := rlib.GetRentableByName(ctx, sss[0], ra.BID)
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not load rentable named: %s  err = %s", funcname, lineno, sss[0], err.Error())
			}
			x, err := strconv.ParseFloat(strings.TrimSpace(sss[1]), 64)
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid amount:  %s", funcname, lineno, sss[1])
			}
			rar.RID = rnt.RID
			rar.RARDtStart = DtStart
			rar.RARDtStop = DtStop
			rar.ContractRent = x
			m = append(m, rar)
		}
	}

	//-------------------------------------------------------------------
	// Notes
	//-------------------------------------------------------------------
	note := strings.TrimSpace(sa[Notes])
	if len(note) > 0 {
		var nl rlib.NoteList
		nl.BID = ra.BID
		nl.NLID, err = rlib.InsertNoteList(ctx, &nl)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error creating NoteList = %s", funcname, lineno, err.Error())
		}
		var n rlib.Note
		n.Comment = note
		n.NTID = 1 // first comment type
		n.BID = nl.BID
		n.NLID = nl.NLID
		_, err = rlib.InsertNote(ctx, &n)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error creating NoteList = %s", funcname, lineno, err.Error())
		}
		ra.NLID = nl.NLID
	}

	//-------------------------------------------------------------------
	// look for any rental agreements already in existence that cover
	// the rentables referenced in this one...
	//-------------------------------------------------------------------
	for i := 0; i < len(m); i++ {
		// TODO(Steve): ignore error?
		rra, _ := rlib.GetAgreementsForRentable(ctx, m[i].RID, &ra.AgreementStart, &ra.AgreementStop)
		for j := 0; j < len(rra); j++ {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s:: Rentable %s is already included in Rental Agreement %s from %s to %s",
				funcname, lineno, RentableAlreadyRented,
				rlib.IDtoString("R", rra[j].RID), rlib.IDtoString("RA", rra[j].RAID),
				rra[j].RARDtStart.Format(rlib.RRDATEFMT4), rra[j].RARDtStop.Format(rlib.RRDATEFMT4))
		}
	}

	//-----------------------------------------------
	// Validate that we have at least one payor...
	//-----------------------------------------------
	if len(payors) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No valid payors for this rental agreement", funcname, lineno)
	}

	//------------------------------------
	// Write the rental agreement record
	//-----------------------------------
	RAID, err := rlib.InsertRentalAgreement(ctx, &ra)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting rlib.RentalAgreement = %v", funcname, lineno, err)
	}
	var lm rlib.LedgerMarker
	lm.Dt = ra.AgreementStart
	lm.RAID = ra.RAID
	lm.State = rlib.LMINITIAL
	_, err = rlib.InsertLedgerMarker(ctx, &lm)

	//------------------------------------------------------------
	// Add the rentables, and the users of those rentables...
	//------------------------------------------------------------
	for i := 0; i < len(m); i++ {
		m[i].RAID = RAID
		m[i].BID = ra.BID
		rlib.InsertRentalAgreementRentable(ctx, &m[i])
		//-----------------------------------------------------
		// Create a Rentable Ledger marker
		//-----------------------------------------------------
		var rlm = rlib.LedgerMarker{
			BID:     ra.BID,
			RAID:    RAID,
			RID:     m[i].RID,
			Dt:      m[i].RARDtStart,
			Balance: float64(0),
			State:   rlib.LMINITIAL,
		}
		_, err = rlib.InsertLedgerMarker(ctx, &rlm)
		if nil != err {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting Rentable LedgerMarker = %v", funcname, lineno, err)
		}

		//------------------------------
		// Add the users
		//------------------------------
		for j := 0; j < len(users); j++ {
			users[j].RID = m[i].RID
			users[j].BID = ra.BID
			_, err := rlib.InsertRentableUser(ctx, &users[j])
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting RentableUser = %v", funcname, lineno, err)
			}
		}
	}

	//------------------------------
	// Add the payors
	//------------------------------
	for i := 0; i < len(payors); i++ {
		payors[i].RAID = RAID
		payors[i].BID = ra.BID
		rlib.InsertRentalAgreementPayor(ctx, &payors[i])
	}
	return 0, nil
}

// LoadRentalAgreementCSV loads a csv file with rental specialty types and processes each one
func LoadRentalAgreementCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateRentalAgreement)
}
