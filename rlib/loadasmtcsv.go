package rlib

import (
	"fmt"
	"strings"
)

// ValidAssessmentDate determines whether the assessment type supplied can be assessed during the assessment's defined period
// given the supplied Rental Agreement period.
// Returns true if the assessment is valid, false otherwise
func ValidAssessmentDate(a *Assessment, asmt *AssessmentType, ra *RentalAgreement) bool {
	v := false // be pessimistic
	inRange := (DateInRange(&a.Start, &ra.RentalStart, &ra.RentalStop) || a.Start.Equal(ra.RentalStart)) && (DateInRange(&a.Stop, &ra.RentalStart, &ra.RentalStop) || a.Stop.Equal(ra.RentalStop))
	before := a.Start.Before(ra.RentalStart) && a.Stop.Before(ra.RentalStop)
	after := (a.Start.After(ra.RentalStart) || a.Start.Equal(ra.RentalStart)) && (a.Stop.After(ra.RentalStop) || a.Stop.Equal(ra.RentalStop))
	switch asmt.RARequired {
	case RARQDINRANGE:
		v = inRange
	case RARQDPRIOR:
		v = inRange || before
	case RARQDAFTER:
		v = inRange || after
	case RARQDANY:
		v = true
	}
	return v
}

// CSV FIELDS FOR THIS MODULE
//    0         1             2      3       4             5             6     7             8                9
// Designation,RentableName, ASMTID, Amount, Start,        Stop,         RAID, RentCycle, ProrationMethod, AcctRule
// REH,         "101",       1,      1000.00,"2014-07-01", "2015-11-08", 1,    6,            4,               "d ${DFLTGENRCV} _, c ${DFLTGSRENT} ${UMR}, d ${DFLTLTL} ${UMR} _ -"
// REH,         "101",       1,      1200.00,"2015-11-21", "2016-11-21", 2,    6,            4,               "d ${DFLTGENRCV} _, c ${DFLTGSRENT} ${UMR}, d ${DFLTLTL} ${UMR} ${aval(${DFLTGENRCV})} -"

// type Assessment struct {
// 	ASMID           int64     // unique id for this assessment
// 	BID             int64     // what Business
// 	RID             int64     // the Rentable
// 	ASMTID          int64     // what type of assessment
// 	RAID            int64     // associated Rental Agreement
// 	Amount          float64   // how much
// 	Start           time.Time // start time
// 	Stop            time.Time // stop time, may be the same as start time or later
// 	Accrual         int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
// 	ProrationMethod int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
// 	AcctRule        string    // expression showing how to account for the amount
// 	Comment         string
// 	LastModTime     time.Time
// 	LastModBy       int64
// }

// CreateAssessmentsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateAssessmentsFromCSV(sa []string, lineno int, AsmtTypes *map[int64]AssessmentType) {
	funcname := "CreateAssessmentsFromCSV"
	var a Assessment
	var r Rentable
	var err error
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 10
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Make sure the Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			Ulog("%s: line %d - Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		a.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Find and set the Rentable
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[1])
	if len(s) > 0 {
		r, err = GetRentableByName(s, a.BID)
		if err != nil {
			fmt.Printf("%s: line %d - Error loading Rentable named: %s.  Error = %v\n", funcname, lineno, s, err)
			return
		}
		a.RID = r.RID
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	DtStart, err := StringToDate(sa[4])
	if err != nil {
		fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[4])
		return
	}
	a.Start = DtStart

	DtStop, err := StringToDate(sa[5])
	if err != nil {
		fmt.Printf("%s: line %d - invalid stop date:  %s\n", funcname, lineno, sa[5])
		return
	}
	a.Stop = DtStop

	//-------------------------------------------------------------------
	// Assessment Type
	//-------------------------------------------------------------------
	a.ASMTID, _ = IntFromString(sa[2], "Assessment type is invalid")
	asmt, ok := (*AsmtTypes)[a.ASMTID]
	if !ok {
		fmt.Printf("%s: line %d - Assessment type is invalid: %s\n", funcname, lineno, sa[2])
		return
	}

	//-------------------------------------------------------------------
	// Rental Agreement ID
	//-------------------------------------------------------------------
	a.RAID, _ = IntFromString(sa[6], "Assessment type is invalid")
	if a.RAID > 0 {
		ra, err := GetRentalAgreement(a.RAID) // for the call to ValidAssessmentDate, we need the entire agreement start/stop period
		if err != nil {
			fmt.Printf("%s: line %d - error loading Rental Agreement with RAID = %s,  error = %s\n", funcname, lineno, sa[6], err.Error())
		}
		if !ValidAssessmentDate(&a, &asmt, &ra) {
			fmt.Printf("%s: line %d - Assessment occurs outside the allowable time range for the Rentable Agreement Require attribute value: %d\n",
				funcname, lineno, asmt.RARequired)
			return
		}
	}

	//-------------------------------------------------------------------
	// Determine the amount
	//-------------------------------------------------------------------
	a.Amount, _ = FloatFromString(sa[3], "Amount is invalid")

	//-------------------------------------------------------------------
	// Accrual
	//-------------------------------------------------------------------
	a.RentCycle, _ = IntFromString(sa[7], "Accrual value is invalid")
	if !IsValidAccrual(a.RentCycle) {
		fmt.Printf("%s: line %d - Accrual must be between %d and %d.  Found %s\n", funcname, lineno, ACCRUALSECONDLY, ACCRUALYEARLY, sa[7])
		return
	}

	//-------------------------------------------------------------------
	// Proration
	//-------------------------------------------------------------------
	a.ProrationMethod, _ = IntFromString(sa[8], "Proration value is invalid")
	if !IsValidAccrual(a.ProrationMethod) {
		fmt.Printf("%s: line %d - Proration must be between %d and %d.  Found %d\n", funcname, lineno, ACCRUALSECONDLY, ACCRUALYEARLY, a.ProrationMethod)
		return
	}
	if a.ProrationMethod > a.RentCycle {
		fmt.Printf("%s: line %d - Proration granularity (%d) must be more frequent than the Accrual (%d)\n", funcname, lineno, a.ProrationMethod, a.RentCycle)
		return
	}

	//-------------------------------------------------------------------
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	a.AcctRule = sa[9]

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(a.AcctRule) == 0 {
		fmt.Printf("%s: line %d - Skipping this record as there is no AcctRule\n", funcname, lineno)
		return
	}
	if a.Amount == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the Amount is 0\n", funcname, lineno)
		return
	}
	if a.RID == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the Rentable ID could not be found\n", funcname, lineno)
		return
	}
	if a.ASMTID == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the AssessmentType could not be found\n", funcname, lineno)
		return
	}
	if a.BID == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the Business could not be found\n", funcname, lineno)
		return
	}

	if a.RAID == 0 {
		fmt.Printf("%s: line %d - Skipping this record as the Rental Agreement could not be found\n", funcname, lineno)
		return
	}

	err = InsertAssessment(&a)
	if err != nil {
		fmt.Printf("%s: line %d - error inserting assessment: %v\n", funcname, lineno, err)
	}

}

// LoadAssessmentsCSV loads a csv file with a chart of accounts and creates Ledger markers for each
func LoadAssessmentsCSV(fname string, AsmtTypes *map[int64]AssessmentType) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateAssessmentsFromCSV(t[i], i+1, AsmtTypes)
	}
}
