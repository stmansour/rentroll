package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

//                            Ledger NAME or GLAccountID or LID works
// 0   1                      2             3
// Bud,Name,                  ARType,       Debit,              Credit,                    Allocated,RAIDRequired,SubARSpec             ,Description
// REX,Rent,                  Assessment,   2,                  8,                         No,       No,         ,                      ,Rent assessment
// REX,AutoGenFloatSECDEPAsmt,SubAssessment,RentRollReceivables,BankAcct,                  No,       No,         ,                      ,auto create assessment
// REX,RCVFloatSecDep,        Receipt,      Undeposited Funds,  Floating Security Deposits,Yes,      Yes,        ,AutoGenFloatSECDEPAsmt,take payment and apply to auto gen'd asmt
// REX,FNB,                   Receipt,      Receivables,        7,                         Yes,      Yes,        ,                      ,payments that are deposited in First National Bank

// CreateAR creates AR database records from the supplied CSV file lines
func CreateAR(sa []string, lineno int) (int, error) {
	funcname := "CreateAR"
	var b rlib.AR
	var err error

	const (
		BUD          = 0
		Name         = iota
		ARType       = iota
		DebitLID     = iota
		CreditLID    = iota
		Allocated    = iota
		ShowOnRA     = iota
		RAIDRequired = iota
		SubARSpec    = iota
		Description  = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"ARType", ARType},
		{"Debit", DebitLID},
		{"Credit", CreditLID},
		{"Allocated", Allocated},
		{"ShowOnRA", ShowOnRA},
		{"RAIDRequired", RAIDRequired},
		{"SubARSpec", SubARSpec},
		{"Description", Description},
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
	des := strings.TrimSpace(sa[BUD])
	if len(des) > 0 { // make sure it's not empty
		b1 := rlib.GetBusinessByDesignation(des) // see if we can find the biz
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		b.BID = b1.BID
	}

	//-----------------------------------------
	// Get the name
	//-----------------------------------------
	b.Name = sa[Name]
	b2, err := rlib.GetARByName(b.BID, b.Name)
	if err != nil && !rlib.IsSQLNoResultsError(err) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error attempting to read existing records with name = %s: %s", funcname, lineno, b.Name, err.Error())
	}
	if b2.Name == b.Name {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - An AR rule with name = %s already exists. Ignoring this line", funcname, lineno, b.Name)
	}

	//-----------------------------------------
	// Get the type
	//-----------------------------------------
	s := strings.TrimSpace(strings.ToLower(sa[ARType]))
	switch s {
	case "assessment":
		b.ARType = 0
	case "receipt":
		b.ARType = 1
	case "expense":
		b.ARType = 2
	case "sub-assessment":
		b.ARType = 3
	default:
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - ARType must be either Assessment or Receipt.  Found: %s", funcname, lineno, s)
	}

	//----------------------------------------------------------------
	// Get the Debit account
	//----------------------------------------------------------------
	var gl rlib.GLAccount
	// rlib.Console("sa[DebitLID] = %s\n", sa[DebitLID])
	b.DebitLID, err = rlib.IntFromString(sa[DebitLID], "Invalid DebitLID") // first see if it is a LID
	if err == nil && b.DebitLID > 0 {                                      // try the LID first
		gl = rlib.GetLedger(b.DebitLID)
	}
	if gl.LID == 0 && len(sa[DebitLID]) > 0 {
		gl = rlib.GetLedgerByName(b.BID, sa[DebitLID]) // if it's not a number, try the Name
		if gl.LID == 0 {
			gl = rlib.GetLedgerByGLNo(b.BID, sa[DebitLID]) // if not a name, then try GLNumber
		}
	}
	// rlib.Console("DEBIT LID = %s\n", gl.Name)
	b.DebitLID = gl.LID
	if b.DebitLID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not find GLAccount for = %s", funcname, lineno, sa[DebitLID])
	}

	//----------------------------------------------------------------
	// Get the Credit account
	//----------------------------------------------------------------
	var glc rlib.GLAccount
	// rlib.Console("sa[CreditLID] = %s\n", sa[CreditLID])
	b.CreditLID, err = rlib.IntFromString(sa[CreditLID], "Invalid CreditLID")
	if err == nil || b.CreditLID > 0 {
		glc = rlib.GetLedger(b.CreditLID)
	}
	if glc.LID == 0 && len(sa[CreditLID]) > 0 {
		glc = rlib.GetLedgerByName(b.BID, sa[CreditLID])
		if glc.LID == 0 {
			glc = rlib.GetLedgerByGLNo(b.BID, sa[CreditLID]) // if not a name, then try GLNumber
		}
	}
	// rlib.Console("CREDIT LID = %s\n", glc.Name)
	b.CreditLID = glc.LID
	if b.CreditLID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid GLAccountID: %s", funcname, lineno, sa[CreditLID])
	}

	//----------------------------------------------------------------
	// Set flags
	//----------------------------------------------------------------
	var alloc, show, rarqd int64
	var yn string
	yn = strings.TrimSpace(sa[Allocated])
	if len(yn) == 0 {
		yn = "no"
	}
	alloc, err = rlib.YesNoToInt(yn)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid Allocated column value: %s", funcname, lineno, sa[Allocated])
	}
	if alloc > 0 {
		b.FLAGS |= 1 << 0
	}
	yn = strings.TrimSpace(sa[ShowOnRA])
	if len(yn) == 0 {
		yn = "no"
	}
	show, err = rlib.YesNoToInt(yn)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid ShowOnRA column value: %s", funcname, lineno, sa[ShowOnRA])
	}
	if show > 0 {
		b.FLAGS |= 1 << 1
	}
	yn = strings.TrimSpace(sa[RAIDRequired])
	if len(yn) == 0 {
		yn = "no"
	}
	rarqd, err = rlib.YesNoToInt(yn)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid RAIDRequired column value: %s", funcname, lineno, sa[RAIDRequired])
	}
	if rarqd > 0 {
		b.FLAGS |= 1 << 2
	}

	//----------------------------------------------------------------
	// Get SubARs - add the array of subars to b, then update the
	// subars with this ARID after we've saved b below
	//----------------------------------------------------------------
	sars := strings.TrimSpace(sa[SubARSpec])
	if len(sars) > 0 {
		sarsa := strings.Split(sars, ",")
		for i := 0; i < len(sarsa); i++ {
			var x rlib.SubAR
			x.BID = b.BID
			subar, err := rlib.GetARByName(b.BID, sarsa[i])
			if err != nil {
				return CsvErrorSensitivity, fmt.Errorf("%s: line %d - could not get Account Rule named: %s, err: %s", funcname, lineno, sarsa[i], err.Error())
			}
			x.SubARID = subar.ARID
			b.SubARs = append(b.SubARs, x) // we will need to update these structs with b.ARID after we save it below
		}
		if len(b.SubARs) > 0 {
			b.FLAGS |= (1 << 3) // bit 3 indicates that there are subars on this rule
		}
	}

	//----------------------------------------------------------------
	// Get the Description
	//----------------------------------------------------------------
	b.Description = sa[Description]

	_, err = rlib.InsertAR(&b)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: error inserting AR = %v", funcname, err)
	}

	for i := 0; i < len(b.SubARs); i++ {
		b.SubARs[i].ARID = b.ARID
		err = rlib.InsertSubAR(&b.SubARs[i])
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error saving SubAR[%d]: %s", funcname, lineno, i, err.Error())
		}
	}

	return 0, nil
}

// LoadARCSV loads the values from the supplied csv file and creates AR records.
func LoadARCSV(fname string) []error {
	return LoadRentRollCSV(fname, CreateAR)
}
