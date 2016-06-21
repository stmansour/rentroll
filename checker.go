package main

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// CheckIssue is a generic type for describing a discrepancy to which users should be alerted
type CheckIssue struct {
	err int64 // 0 = unknown, 1 = Assessment & Contract Rent amount mismatch, 2 = Contract Rent could not be found, 3 = database error
	a   rlib.Assessment
	rar rlib.RentalAgreementRentable
	dt  time.Time
	msg string
}

// AssessmentChecker examines assessments made for rent and compares them to the
// contract rent amount associated with the rental agreement for each rentable.
// If it finds a discrepancy, it adds a CheckIssue struct to a list.  After processing
// all assessments for the supplied period, it returns the list of CheckIssues.
func AssessmentChecker(xbiz *rlib.XBusiness, d1, d2 *time.Time) []CheckIssue {
	funcname := "AssessmentChecker"
	var m []CheckIssue
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(xbiz.P.BID, d2, d1)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a rlib.Assessment
		ap := &a
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.RecurCycle, &a.ProrationCycle, &a.AcctRule, &a.Comment, &a.LastModTime, &a.LastModBy))
		if rlib.AssessmentIsType(&a, "rent", xbiz) { // process it only if it is a Rent assessment
			dl := ap.GetRecurrences(d1, d2) // get the recurrences that fall in the specified range
			for i := 0; i < len(dl); i++ {  // process each occurrence
				var c CheckIssue // open an issue to use, but don't store it unless there's actually an issue
				c.dt = dl[i]
				rar, err := rlib.FindAgreementByRentable(a.RID, &c.dt, d2) // what is the ContractRent on dl[i]?
				if err != nil {                                            // if we encounter a db error...
					if !rlib.IsSQLNoResultsError(err) { // and it's something other "couldn't find anything..."
						c.msg = fmt.Sprintf("%s: could not load Contract rent for rentable %d between %s and %s, err = %s\n",
							funcname, a.RID, d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT), err.Error())
						rlib.Ulog(c.msg) // log the error
						c.err = 3        // ... that indicates we cannot find the contract rent...
						c.a = a
						m = append(m, c) // add it to the list
						continue         // and keep moving
					}
					// we could not find a contract rent
					c.err = 2        // ... that indicates we cannot find the contract rent...
					c.a = a          // ... as called out in this assessment
					m = append(m, c) // add it to the list
					continue         // and keep moving
				}
				if a.Amount != rar.ContractRent {
					c.err = 1        // ... rent amount mismatch ...
					c.a = a          // what's called out in this assessment
					c.rar = rar      // the contract rent
					m = append(m, c) // add it to the list
				}
			}
		}
	}
	rlib.Errcheck(rows.Err())
	return m
}

// AssessmentErrToString provides a string explanation fo the supplied CheckIssue error
func AssessmentErrToString(err int64) string {
	var s string
	switch err {
	case 1:
		s = "Assessment amount and contract rent are different"
	case 2:
		s = "Contract rent could not be found"
	default:
		s = fmt.Sprintf("unrecognized error: %d", err)
	}
	return s
}

// AssessmentCheckReportText runs a check over all Assessments in period d1-d2 that are Rent assessments.
// It validates that the Assessment amount matches the ContractRent in the associated RentalAgreementRentable
// record. If there is a discrepancy, it is listed in the report.
func AssessmentCheckReportText(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	foundIssues := true
	m := AssessmentChecker(xbiz, d1, d2)
	if m == nil {
		foundIssues = false
	}
	if foundIssues && len(m) == 0 {
		foundIssues = false
	}

	fmt.Printf("ASSESSMENT CHECK\n")
	fmt.Printf("For Assessments between %s and %s\n", d1.Format(rlib.RRDATEINPFMT), d2.Format(rlib.RRDATEINPFMT))
	fmt.Printf("------------------------------------------------------------------------------\n")
	if !foundIssues {
		fmt.Printf("No issues to report. All Assessments match with Contract Rent\n")
		return
	}
	for i := 0; i < len(m); i++ {
		s := ""
		if len(m[i].msg) > 0 {
			s = m[i].msg
		} else {
			s = AssessmentErrToString(m[i].err)
		}
		fmt.Printf("%s - %s\n", m[i].dt.Format(rlib.RRDATEINPFMT), s)
		switch m[i].err {
		case 1:
			fmt.Printf("%13sAssessment A%08d amount = %6.2f, contract rent = %6.2f\n", " ", m[i].a.ASMID, m[i].a.Amount, m[i].rar.ContractRent)
		}
	}
}
