package main

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

func (a *assessment) GetRecurrences(start, stop *time.Time) []time.Time {
	return rlib.GetRecurrences(start, stop, &a.Start, &a.Stop, a.Frequency)
}

// return a slice of assessments for the unit associated with this
// occupancy agreement.
func propertyAssessments(oa *occupancyAgreement, start, stop *time.Time) {
	// pull in the assessments associated with the unit for this occupancy agreement
	// rows, err := App.dbrr.Query()
	// rlib.Errcheck(err)
	// defer rows.Close()
	// SELECT ASMID,UNITID,ASMTID,Amount,Start,Stop,Frequency FROM assessments WHERE Stop > '2015-12-01T00:00:00Z' and Start <= '2015-12-31T23:59:59Z'
	s := fmt.Sprintf("SELECT ASMID,UNITID,ASMTID,Amount,Start,Stop,Frequency FROM assessments WHERE UNITID=%d and Stop >= '%s' and Start < '%s'",
		oa.UNITID, start.Format(time.RFC3339), stop.Format(time.RFC3339))
	// fmt.Printf("s = %s\n", s)
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	var a assessment
	ap := &a
	for rows.Next() {
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.UNITID, &a.ASMTID, &a.Amount, &a.Start, &a.Stop, &a.Frequency))
		// fmt.Printf("\ta: %2d, amount: %8.2f,  freq: %d\n", a.ASMTID, a.Amount, a.Frequency)
		ap.GetRecurrences(start, stop)
	}
}

// calculate all charges for the specified property that occur in
// the supplied start / stop time range.
func doPropertyAssessments(start, stop *time.Time, PRID int) {
	// find all
	//rows, err := App.dbrr.Query("SELECT OAID,OATID,PRID,UNITID,PID,PrimaryTenant,OccupancyStart,OccupancyStop,Renewal,ProrationMethod,SecurityDepositAmount from occupancyagreement where PRID=?", PRID)
	rows, err := App.prepstmt.occAgrByProperty.Query(PRID)
	rlib.Errcheck(err)
	defer rows.Close()

	var oa occupancyAgreement

	for rows.Next() {
		rlib.Errcheck(rows.Scan(&oa.OAID, &oa.OATID, &oa.PRID, &oa.UNITID, &oa.PID, &oa.PrimaryTenant,
			&oa.OccupancyStart, &oa.OccupancyStop, &oa.Renewal, &oa.ProrationMethod, &oa.SecurityDepositAmount))

		// process the active agreements
		fmt.Printf("Unit %d assessments:\n", oa.UNITID)
		if (oa.OccupancyStart.Equal(*start) || oa.OccupancyStop.After(*start)) && oa.OccupancyStart.Before(*stop) {
			propertyAssessments(&oa, start, stop)
		}
	}
	rlib.Errcheck(rows.Err())
}
