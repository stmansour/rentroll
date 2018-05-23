package rcsv

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"strings"
)

// // ValidAssessmentDate determines whether the assessment type supplied can be assessed during the assessment's defined period
// // given the supplied Rental Agreement period.
// // Returns true if the assessment is valid, false otherwise
// func ValidAssessmentDate(a *rlib.Assessment, asmt *rlib.GLAccount, ra *rlib.RentalAgreement) bool {
// 	v := false // be pessimistic
// 	t1 := rlib.DateInRange(&a.Start, &ra.AgreementStart, &ra.AgreementStop)
// 	t2 := a.Start.Equal(ra.AgreementStart)
// 	t3 := rlib.DateInRange(&a.Stop, &ra.AgreementStart, &ra.AgreementStop)
// 	t4 := a.Stop.Equal(ra.AgreementStop)

// 	// fmt.Printf("a.Start = %s, a.Stop = %s, ra.AgrStart = %s, ra.AgrStop = %s\n", a.Start.Format(rlib.RRDATEFMT4), a.Stop.Format(rlib.RRDATEFMT4), ra.AgreementStart.Format(rlib.RRDATEFMT4), ra.AgreementStop.Format(rlib.RRDATEFMT4))
// 	// fmt.Printf("t1 = %t, t2 = %t, t3 = %t, t4 = %t\n", t1, t2, t3, t4)

// 	inRange := (t1 || t2) && (t3 || t4)
// 	before := a.Start.Before(ra.AgreementStart) && a.Stop.Before(ra.AgreementStop)
// 	after := (a.Start.After(ra.AgreementStart) || a.Start.Equal(ra.AgreementStart)) && (a.Stop.After(ra.AgreementStop) || a.Stop.Equal(ra.AgreementStop))

// 	switch asmt.RARequired {
// 	case rlib.RARQDINRANGE:
// 		v = inRange
// 	case rlib.RARQDPRIOR:
// 		v = inRange || before
// 	case rlib.RARQDAFTER:
// 		v = inRange || after
// 	case rlib.RARQDANY:
// 		v = true
// 	}
// 	// fmt.Printf("inRange = %t, before = %t, after = %t, v = %t, GLAccount = %s (%d)\n", inRange, before, after, v, asmt.Name, asmt.LID)
// 	return v
// }

// CSV FIELDS FOR THIS MODULE
//    0  1             2      3       4             5             6     7             8                9          10
// BUD   ,RentableName, GLAcctID, Amount, Start,        Stop,         RAID, RentCycle, ProrationCycle, InvoiceNo, AcctRule,                                                                         AR
// REH,  "101",       1,      1000.00,"2014-07-01", "2015-11-08", 1,    6,            4,               20122,     "d ${GLGENRCV} _, c ${GLGSRENT} ${UMR}, d ${GLLTL} ${UMR} _ -",                   Rent Non-Taxable
// REH,  "101",       1,      1200.00,"2015-11-21", "2016-11-21", 2,    6,            4,               739928,    "d ${GLGENRCV} _, c ${GLGSRENT} ${UMR}, d ${GLLTL} ${UMR} ${aval(${GLGENRCV})} -", "Rent Payment Check"

// CreateAssessmentsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateAssessmentsFromCSV(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreateAssessmentsFromCSV"
	var (
		err  error
		xbiz rlib.XBusiness
		r    rlib.Rentable
		a    rlib.Assessment
		des  = strings.ToLower(strings.TrimSpace(sa[0]))
	)

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
		AR             = iota
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
		{"AR", AR},
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
	if len(des) > 0 {
		b1, err := rlib.GetBusinessByDesignation(ctx, des)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - rlib.Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - rlib.Business with designation %s does not exist", funcname, lineno, sa[0])
		}
		err = rlib.InitBizInternals(b1.BID, &xbiz) // this initializes a number of internal variables that the internals need and is efficient if they are already loaded
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error while initializing biz internals. Error: %s", funcname, lineno, err.Error())
		}
		Rcsv.Xbiz = &xbiz
		a.BID = Rcsv.Xbiz.P.BID
	}

	//-------------------------------------------------------------------
	// Find and set the rlib.Rentable
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[RentableName])
	if len(s) > 0 {
		r, err = rlib.GetRentableByName(ctx, s, a.BID)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error loading rlib.Rentable named: %s.  Error = %v", funcname, lineno, s, err)
		}
		a.RID = r.RID
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	d1, err := rlib.StringToDate(sa[DtStart])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid start date:  %s", funcname, lineno, sa[DtStart])
	}
	a.Start = d1

	d2, err := rlib.StringToDate(sa[DtStop])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - invalid stop date:  %s", funcname, lineno, sa[DtStop])
	}
	a.Stop = d2

	//-------------------------------------------------------------------
	// rlib.Assessment Type
	//-------------------------------------------------------------------
	// var gla rlib.GLAccount
	var ok bool
	if len(sa[GLAcctID]) > 0 {
		a.ATypeLID, _ = rlib.IntFromString(sa[GLAcctID], "value for Assessment type is invalid")
		// asmt, ok := (*AsmtTypes)[a.ATypeLID]
		rlib.InitBusinessFields(a.BID)
		// rlib.GetDefaultLedgers(a.BID) // Gather its chart of accounts
		// TODO(Steve): ignore error?
		rlib.RRdb.BizTypes[a.BID].GLAccounts, _ = rlib.GetGLAccountMap(ctx, a.BID)
		/*gla,*/ _, ok = rlib.RRdb.BizTypes[a.BID].GLAccounts[a.ATypeLID]
		if !ok {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Assessment type is invalid: %s", funcname, lineno, sa[2])
		}
	}

	//-------------------------------------------------------------------
	// Rental Agreement ID
	//-------------------------------------------------------------------
	a.RAID, _ = rlib.IntFromString(sa[RAID], "Rental Agreement ID is invalid")
	if a.RAID > 0 {
		ra, err := rlib.GetRentalAgreement(ctx, a.RAID) // for the call to ValidAssessmentDate, we need the entire agreement start/stop period
		if err != nil {
			fmt.Printf("%s: line %d - error loading Rental Agreement with RAID = %s,  error = %s\n", funcname, lineno, sa[6], err.Error())
		}
		if ra.RAID == 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Rental agreement %d could not be found", funcname, lineno, a.RAID)
		}
		// 	if !ValidAssessmentDate(&a, &gla, &ra) {
		// 		rs := fmt.Sprintf("%s: line %d - Assessment occurs outside the allowable time range for the Rentable Agreement Require attribute value: %d\n",
		// 			funcname, lineno, gla.RARequired)
		// 		rs += fmt.Sprintf("Rental Agreement start/stop: %s - %s \n", ra.AgreementStart.Format(rlib.RRDATEFMT3), ra.AgreementStop.Format(rlib.RRDATEFMT3))
		// 		return CsvErrorSensitivity, fmt.Errorf("%s      Assessment start/stop: %s - %s ", rs, a.Start.Format(rlib.RRDATEFMT3), a.Stop.Format(rlib.RRDATEFMT3))
		// 	}
	}

	//-------------------------------------------------------------------
	// Determine the amount
	//-------------------------------------------------------------------
	a.Amount, _ = rlib.FloatFromString(sa[Amount], "Amount is invalid")

	//-------------------------------------------------------------------
	// Accrual
	//-------------------------------------------------------------------
	a.RentCycle, _ = rlib.IntFromString(sa[RentCycle], "Accrual value is invalid")
	if !rlib.IsValidAccrual(a.RentCycle) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Accrual must be between %d and %d.  Found %d", funcname, lineno, rlib.RECURSECONDLY, rlib.RECURYEARLY, a.RentCycle)
	}

	//-------------------------------------------------------------------
	// Proration
	//-------------------------------------------------------------------
	a.ProrationCycle, _ = rlib.IntFromString(sa[ProrationCycle], "Proration value is invalid")
	if !rlib.IsValidAccrual(a.ProrationCycle) {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Proration must be between %d and %d.  Found %d", funcname, lineno, rlib.RECURSECONDLY, rlib.RECURYEARLY, a.ProrationCycle)
	}
	if a.ProrationCycle > a.RentCycle {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Proration granularity (%d) must be more frequent than the Accrual (%d)", funcname, lineno, a.ProrationCycle, a.RentCycle)
	}

	//-------------------------------------------------------------------
	// Set the InvoiceNo.  If not specified, just leave it as 0
	//-------------------------------------------------------------------
	a.InvoiceNo, err = rlib.IntFromString(sa[InvoiceNo], "InvoiceNo is invalid")
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("Bad InvoiceNo: " + sa[InvoiceNo])
	}

	//-------------------------------------------------------------------
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	a.AcctRule = sa[AcctRule]

	//-------------------------------------------------------------------
	// Set the ARID
	//-------------------------------------------------------------------
	s = strings.TrimSpace(sa[AR])
	if len(s) > 0 {
		rule, err := rlib.GetARByName(ctx, a.BID, s)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not load AR named %s: %s", funcname, lineno, s, err.Error())
		}
		a.ARID = rule.ARID
	}

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(a.AcctRule) == 0 && a.RID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Skipping this record as there is no AcctRule", funcname, lineno)
	}
	if a.Amount == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Skipping this record as the Amount is 0", funcname, lineno)
	}
	if a.RID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Skipping this record as the rlib.Rentable ID could not be found", funcname, lineno)
	}
	if a.BID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Skipping this record as the rlib.Business could not be found", funcname, lineno)
	}

	if a.RAID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Skipping this record as the Rental Agreement could not be found", funcname, lineno)
	}

	// TODO(Steve): ignore error?
	adup, _ := rlib.GetAssessmentDuplicate(ctx, &a.Start, a.Amount, a.PASMID, a.RID, a.RAID, a.ATypeLID)
	if adup.ASMID != 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - this is a duplicate of an existing assessment: %s", funcname, lineno, adup.IDtoString())
	}

	_, err = rlib.InsertAssessment(ctx, &a)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting assessment: %v", funcname, lineno, err)
	}

	// process this new assessment over the requested time range...
	err = rlib.ProcessJournalEntry(ctx, &a, Rcsv.Xbiz, &Rcsv.DtStart, &Rcsv.DtStop, false)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error while processing journal entries. Error: %s", funcname, lineno, err.Error())
	}

	return 0, nil
}

// LoadAssessmentsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadAssessmentsCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreateAssessmentsFromCSV)
}
