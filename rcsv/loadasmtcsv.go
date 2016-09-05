package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// ValidAssessmentDate determines whether the assessment type supplied can be assessed during the assessment's defined period
// given the supplied Rental Agreement period.
// Returns true if the assessment is valid, false otherwise
func ValidAssessmentDate(a *rlib.Assessment, asmt *rlib.GLAccount, ra *rlib.RentalAgreement) bool {
	v := false // be pessimistic
	inRange := (rlib.DateInRange(&a.Start, &ra.AgreementStart, &ra.AgreementStop) || a.Start.Equal(ra.AgreementStart)) && (rlib.DateInRange(&a.Stop, &ra.AgreementStart, &ra.AgreementStop) || a.Stop.Equal(ra.AgreementStop))
	before := a.Start.Before(ra.AgreementStart) && a.Stop.Before(ra.AgreementStop)
	after := (a.Start.After(ra.AgreementStart) || a.Start.Equal(ra.AgreementStart)) && (a.Stop.After(ra.AgreementStop) || a.Stop.Equal(ra.AgreementStop))
	switch asmt.RARequired {
	case rlib.RARQDINRANGE:
		v = inRange
	case rlib.RARQDPRIOR:
		v = inRange || before
	case rlib.RARQDAFTER:
		v = inRange || after
	case rlib.RARQDANY:
		v = true
	}
	return v
}

// CSV FIELDS FOR THIS MODULE
//    0  1             2      3       4             5             6     7             8                9          10
// BUD   ,RentableName, GLAcctID, Amount, Start,        Stop,         RAID, RentCycle, ProrationCycle, InvoiceNo, AcctRule
// REH,  "101",       1,      1000.00,"2014-07-01", "2015-11-08", 1,    6,            4,               20122,     "d ${GLGENRCV} _, c ${GLGSRENT} ${UMR}, d ${GLLTL} ${UMR} _ -"
// REH,  "101",       1,      1200.00,"2015-11-21", "2016-11-21", 2,    6,            4,               739928,    "d ${GLGENRCV} _, c ${GLGSRENT} ${UMR}, d ${GLLTL} ${UMR} ${aval(${GLGENRCV})} -"

// CreateAssessmentsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateAssessmentsFromCSV(sa []string, lineno int) int {
	funcname := "CreateAssessmentsFromCSV"
	var a rlib.Assessment
	var r rlib.Rentable
	var err error
	des := strings.ToLower(strings.TrimSpace(sa[0]))

	const (
		BUD            = 0
		RentableName   = iota
		GLAcctID       = iota
		Amount         = iota
		DtStart        = iota
		DtStop         = iota
		RAID           = iota
		RentCycle      = iota
		ProrationCycle = iota
		InvoiceNo      = iota
		AcctRule       = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"RentableName", RentableName},
		{"GLAcctID", GLAcctID},
		{"Amount", Amount},
		{"DtStart", DtStart},
		{"DtStop", DtStop},
		{"RAID", RAID},
		{"RentCycle", RentCycle},
		{"ProrationCycle", ProrationCycle},
		{"InvoiceNo", InvoiceNo},
		{"AcctRule", AcctRule},
	}

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return 1
	}
	if lineno == 1 {
		return 0
	}
	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1 := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d - rlib.Business with designation %s does not exist\n", funcname, lineno, sa[0])
			return CsvErrorSensitivity
		}
		a.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Find and set the rlib.Rentable
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[1])
	if len(s) > 0 {
		r, err = rlib.GetRentableByName(s, a.BID)
		if err != nil {
			fmt.Printf("%s: line %d - Error loading rlib.Rentable named: %s.  Error = %v\n", funcname, lineno, s, err)
			return CsvErrorSensitivity
		}
		a.RID = r.RID
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	d1, err := rlib.StringToDate(sa[4])
	if err != nil {
		fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[4])
		return CsvErrorSensitivity
	}
	a.Start = d1

	d2, err := rlib.StringToDate(sa[5])
	if err != nil {
		fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[5])
		return CsvErrorSensitivity
	}
	a.Stop = d2

	//-------------------------------------------------------------------
	// rlib.Assessment Type
	//-------------------------------------------------------------------
	a.ATypeLID, _ = rlib.IntFromString(sa[2], "rlib.Assessment type is invalid")
	// asmt, ok := (*AsmtTypes)[a.ATypeLID]
	rlib.InitBusinessFields(a.BID)
	rlib.GetDefaultLedgers(a.BID) // Gather its chart of accounts
	rlib.RRdb.BizTypes[a.BID].GLAccounts = rlib.GetGLAccountMap(a.BID)
	gla, ok := rlib.RRdb.BizTypes[a.BID].GLAccounts[a.ATypeLID]
	if !ok {
		fmt.Printf("%s: line %d - rlib.Assessment type is invalid: %s\n", funcname, lineno, sa[2])
		return CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Rental Agreement ID
	//-------------------------------------------------------------------
	a.RAID, _ = rlib.IntFromString(sa[6], "Assessment type is invalid")
	if a.RAID > 0 {
		ra, err := rlib.GetRentalAgreement(a.RAID) // for the call to ValidAssessmentDate, we need the entire agreement start/stop period
		if err != nil {
			fmt.Printf("%s: line %d - error loading Rental Agreement with RAID = %s,  error = %s\n", funcname, lineno, sa[6], err.Error())
		}
		if !ValidAssessmentDate(&a, &gla, &ra) {
			fmt.Printf("%s: line %d - Assessment occurs outside the allowable time range for the Rentable Agreement Require attribute value: %d\n",
				funcname, lineno, gla.RARequired)
			fmt.Printf("Rental Agreement start/stop: %s - %s \n", ra.AgreementStart.Format(rlib.RRDATEFMT3), ra.AgreementStop.Format(rlib.RRDATEFMT3))
			fmt.Printf("      Assessment start/stop: %s - %s \n", a.Start.Format(rlib.RRDATEFMT3), a.Stop.Format(rlib.RRDATEFMT3))
			return CsvErrorSensitivity
		}
	}

	//-------------------------------------------------------------------
	// Determine the amount
	//-------------------------------------------------------------------
	a.Amount, _ = rlib.FloatFromString(sa[3], "Amount is invalid")

	//-------------------------------------------------------------------
	// Accrual
	//-------------------------------------------------------------------
	a.RentCycle, _ = rlib.IntFromString(sa[7], "Accrual value is invalid")
	if !rlib.IsValidAccrual(a.RentCycle) {
		fmt.Printf("%s: line %d - Accrual must be between %d and %d.  Found %s\n", funcname, lineno, rlib.CYCLESECONDLY, rlib.CYCLEYEARLY, sa[7])
		return CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Proration
	//-------------------------------------------------------------------
	a.ProrationCycle, _ = rlib.IntFromString(sa[8], "Proration value is invalid")
	if !rlib.IsValidAccrual(a.ProrationCycle) {
		fmt.Printf("%s: line %d - Proration must be between %d and %d.  Found %d\n", funcname, lineno, rlib.CYCLESECONDLY, rlib.CYCLEYEARLY, a.ProrationCycle)
		return CsvErrorSensitivity
	}
	if a.ProrationCycle > a.RentCycle {
		fmt.Printf("%s: line %d - Proration granularity (%d) must be more frequent than the Accrual (%d)\n", funcname, lineno, a.ProrationCycle, a.RentCycle)
		return CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Set the InvoiceNo.  If not specified, just leave it as 0
	//-------------------------------------------------------------------
	a.InvoiceNo, ok = rlib.IntFromString(sa[9], "InvoiceNo is invalid")
	if !ok {
		return CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	a.AcctRule = sa[10]

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(a.AcctRule) == 0 {
		fmt.Printf("%s: line %d - Skipping this record as there is no AcctRule\n", funcname, lineno)
		return CsvErrorSensitivity
	}
	if a.Amount == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the Amount is 0\n", funcname, lineno)
		return CsvErrorSensitivity
	}
	if a.RID == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the rlib.Rentable ID could not be found\n", funcname, lineno)
		return CsvErrorSensitivity
	}
	if a.BID == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the rlib.Business could not be found\n", funcname, lineno)
		return CsvErrorSensitivity
	}

	if a.RAID == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the Rental Agreement could not be found\n", funcname, lineno)
		return CsvErrorSensitivity
	}

	_, err = rlib.InsertAssessment(&a)
	if err != nil {
		fmt.Printf("%s: line %d - error inserting assessment: %v\n", funcname, lineno, err)
	}

	// process this new assessment over the requested time range...
	// fmt.Printf("Rcsv.Xbiz = %#v\n", *Rcsv.Xbiz)
	rlib.ProcessJournalEntry(&a, Rcsv.Xbiz, &Rcsv.DtStart, &Rcsv.DtStop)

	return 0
}

// LoadAssessmentsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadAssessmentsCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if t[i][0] == "#" {
			continue
		}
		CreateAssessmentsFromCSV(t[i], i+1)
	}
}
