package rlib

import (
	"fmt"
	"strings"
	"time"
)

// the CSV file format:
//    0         1             2     3       4             5             6          7                8
// Designation,RentableName, ASMTID,Amount, Start,        Stop,         Frequency, ProrationMethod, AcctRule
// REH,         "101",       1,     1000.00,"2014-07-01", "2015-11-08", 6,         4,               "d ${DFLTGENRCV} _, c ${DFLTGSRENT} ${UMR}, d ${DFLTLTL} ${UMR} _ -"
// REH,         "101",       1,     1200.00,"2015-11-21", "2016-11-21", 6,         4,               "d ${DFLTGENRCV} _, c ${DFLTGSRENT} ${UMR}, d ${DFLTLTL} ${UMR} ${aval(${DFLTGENRCV})} -"

// type Assessment struct {
// 	ASMID           int64     // unique id for this assessment
// 	BID             int64     // what business
// 	RID             int64     // the rentable
// 	ASMTID          int64     // what type of assessment
// 	RAID            int64     // associated Rental Agreement
// 	Amount          float64   // how much
// 	Start           time.Time // start time
// 	Stop            time.Time // stop time, may be the same as start time or later
// 	Frequency       int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
// 	ProrationMethod int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
// 	AcctRule        string    // expression showing how to account for the amount
// 	Comment         string
// 	LastModTime     time.Time
// 	LastModBy       int64
// }

// CreateAssessmentsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateAssessmentsFromCSV(sa []string, AsmtTypes *map[int64]AssessmentType) {
	var a Assessment
	var r Rentable
	var err error
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Make sure the business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			Ulog("CreateLedgerMarkers: business with designation %s does net exist\n", sa[0])
			return
		}
		a.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Find and set the rentable
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[1])
	if len(s) > 0 {
		r, err = GetRentableByName(s, a.BID)
		if err != nil {
			fmt.Printf("CreateAssessmentsFromCSV: Error loading rentable named: %s.  Error = %v\n", s, err)
			return
		}
		a.RID = r.RID
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	DtStart, err := time.Parse(RRDATEINPFMT, strings.TrimSpace(sa[4]))
	if err != nil {
		fmt.Printf("CreateAssessmentsFromCSV: invalid start date:  %s\n", sa[4])
		return
	}
	a.Start = DtStart

	DtStop, err := time.Parse(RRDATEINPFMT, strings.TrimSpace(sa[5]))
	if err != nil {
		fmt.Printf("CreateAssessmentsFromCSV: invalid stop date:  %s\n", sa[5])
		return
	}
	a.Stop = DtStop

	//-------------------------------------------------------------------
	// Find and set the rental agreement
	//-------------------------------------------------------------------
	ra, err := FindAgreementByRentable(a.RID, &DtStart, &DtStop)
	if err != nil {
		fmt.Printf("CreateAssessmentsFromCSV: Error finding rental agreement for rentable %s.  Error = %v\n", r.Name, err)
		return
	}
	a.RAID = ra.RAID

	//-------------------------------------------------------------------
	// Determine the amount
	//-------------------------------------------------------------------
	a.Amount = FloatFromString(sa[3], "Amount is invalid")

	//-------------------------------------------------------------------
	// Assessment Type
	//-------------------------------------------------------------------
	a.ASMTID = IntFromString(sa[2], "Assessment type is invalid")
	_, ok := (*AsmtTypes)[a.ASMTID]
	if !ok {
		fmt.Printf("CreateAssessmentsFromCSV: Assessment type is invalid: %s\n", sa[2])
		return
	}

	//-------------------------------------------------------------------
	// Frequency
	//-------------------------------------------------------------------
	a.Frequency = IntFromString(sa[6], "Assessment value is invalid")
	if a.Frequency < OCCTYPESECONDLY || a.Frequency > OCCTYPEYEARLY {
		fmt.Printf("CreateAssessmentsFromCSV: Frequency must be between %d and %d.  Found %d\n", OCCTYPESECONDLY, OCCTYPEYEARLY, a.Frequency)
		return
	}

	//-------------------------------------------------------------------
	// Proration
	//-------------------------------------------------------------------
	a.ProrationMethod = IntFromString(sa[7], "Proration value is invalid")
	if a.ProrationMethod < OCCTYPESECONDLY || a.ProrationMethod > OCCTYPEYEARLY {
		fmt.Printf("CreateAssessmentsFromCSV: Proration must be between %d and %d.  Found %d\n", OCCTYPESECONDLY, OCCTYPEYEARLY, a.ProrationMethod)
		return
	}
	if a.ProrationMethod > a.Frequency {
		fmt.Printf("CreateAssessmentsFromCSV: Proration granularity (%d) must be more frequent than the Frequency (%d)\n", a.ProrationMethod, a.Frequency)
		return
	}

	//-------------------------------------------------------------------
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	a.AcctRule = sa[8]

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(a.AcctRule) == 0 || a.ASMTID == 0 ||
		a.Amount == 0 || a.RID == 0 || a.RAID == 0 || a.BID == 0 {
		fmt.Printf("Skipping this record\n")
		return
	}

	err = InsertAssessment(&a)
	if err != nil {
		fmt.Printf("CreateAssessmentsFromCSV: error inserting assessment: %v\n", err)
	}

}

// LoadAssessmentsCSV loads a csv file with a chart of accounts and creates ledger markers for each
func LoadAssessmentsCSV(fname string, AsmtTypes *map[int64]AssessmentType) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateAssessmentsFromCSV(t[i], AsmtTypes)
	}
}
