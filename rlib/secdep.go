package rlib

import (
	"fmt"
	"strings"
	"time"
)

// SecDepAccts returns a slice of LiabilitySecurityDeposit accounts
func SecDepAccts(bid int64) ([]int64, error) {
	m := []int64{}
	q := fmt.Sprintf("SELECT LID FROM GLAccount WHERE BID=%d AND AcctType=%q", bid, LiabilitySecDep)
	rows, err := RRdb.Dbrr.Query(q)
	if err != nil {
		return m, err
	}
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return m, err
		}
		m = append(m, id)
	}
	err = rows.Err()
	if err != nil {
		return m, err
	}
	return m, nil
}

// GetSecDepBalance returns the amount of security deposit charge and the
// amount that was assessed for the supplied Rental Agreement and RID
// Params
//	bid  - business id
//  rid  - the rentable for which the deposit was assessed
//  raid - the Rental Agreement associated with the assessment
//  d1   - start time; do not consider assessments prior to this date
//  d2   - stop time; do not considder assessments on or after this date
func GetSecDepBalance(bid, raid, rid int64, d1, d2 *time.Time) (float64, error) {
	var xbiz XBusiness
	var m []int64
	amt := float64(0)
	InitBizInternals(bid, &xbiz)
	sda, err := SecDepAccts(bid)
	if err != nil {
		return amt, err
	}
	if len(sda) == 0 {
		return amt, fmt.Errorf("There are no accounts of type %s", LiabilitySecDep)
	}
	sa := []string{}
	for i := 0; i < len(sda); i++ {
		sa = append(sa, fmt.Sprintf("CreditLID=%d", sda[i]))
	}
	q := "SELECT ARID FROM AR WHERE (" + strings.Join(sa, " OR ") + ")"
	rows, err := RRdb.Dbrr.Query(q)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return amt, err
		}
		m = append(m, id)
	}
	err = rows.Err()
	if err != nil {
		return amt, err
	}
	// Console("q = %q\n", q)
	// Console("ARIDs: %v\n", m)
	if len(m) == 0 {
		return amt, fmt.Errorf("There are no account rules that credit a %s account", LiabilitySecDep)
	}
	sa = []string{}
	for i := 0; i < len(m); i++ {
		sa = append(sa, fmt.Sprintf("ARID=%d", m[i]))
	}
	q = fmt.Sprintf("SELECT SUM(Amount) AS Amt FROM Assessments WHERE BID=%d AND RID=%d and RAID=%d AND %q<=Start AND Stop<%q AND (%s) GROUP BY RID",
		bid, rid, raid, d1.Format(RRDATEFMTSQL), d2.Format(RRDATEFMTSQL), strings.Join(sa, " OR "))
	// Console("q: %s\n", q)
	rows, err = RRdb.Dbrr.Query(q)
	for rows.Next() {
		var x float64
		err := rows.Scan(&x)
		if err != nil {
			return amt, err
		}
		amt += x
	}
	err = rows.Err()
	if err != nil {
		return amt, err
	}
	return amt, nil
}
