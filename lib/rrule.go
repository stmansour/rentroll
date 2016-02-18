package main

import "time"

func dateInRange(dt, start, stop *time.Time) bool {
	// fmt.Printf("dt: %s\n", dt.Format(time.RFC3339))
	// fmt.Printf("start: %s   stop: %s\n", start.Format(time.RFC3339), stop.Format(time.RFC3339))
	return (dt.Equal(*start) || dt.After(*start)) && dt.Before(*stop)
}

// a1 - a2 = time range of the assessment
// R1 - R2 = time range for the run calculation
// freq = chunk of time over which to quantize the assessment
func genRegularRecurSeq(a1, a2, R1, R2 *time.Time, freq time.Duration) []time.Time {
	// fmt.Printf("\n---------------------\n")
	// fmt.Printf("a1: %s   a2: %s\n", a1.Format(time.RFC3339), a2.Format(time.RFC3339))
	// fmt.Printf("R1: %s   R2: %s\n", R1.Format(time.RFC3339), R2.Format(time.RFC3339))

	//============================================
	// Set up first time range for first run...
	//============================================
	d1 := *R1
	d2 := d1.Add(freq)

	//============================================
	// Now iterate in chunks (of size: freq) and
	// save the recurrence dates...
	//============================================
	var m []time.Time
	for i := 0; i < 10000; i++ {
		// fmt.Printf("iter: %d\n", i)
		//----------------------------------------
		// don't go outside the requested range
		//----------------------------------------
		if d1.After(*R2) || d1.Equal(*R2) {
			break
		}
		if d2.After(*R2) {
			d2 = *R2
		}
		//----------------------------------------
		//  does a1 - a2 overlap d1-d2
		//----------------------------------------
		// fmt.Printf("d1: %s   d2: %s\n", d1.Format(time.RFC3339), d2.Format(time.RFC3339))
		if !(a1.After(d2) || a1.Equal(d2) || a2.Before(d1) || a2.Equal(d1)) {
			d := time.Date(d1.Year(), d1.Month(), d1.Day(), a1.Hour(), a1.Minute(), a1.Second(), 0, time.UTC)
			// fmt.Printf("STORE d = %s\n", d.Format(time.RFC3339))
			m = append(m, d)
		}
		//----------------------------------------
		//  set the next interval to check...
		//----------------------------------------
		d1 = d1.Add(freq)
		d2 = d2.Add(freq)
	}
	return m
}

func genMonthlyRecurSeq(a1, a2, R1, R2 *time.Time, nMonths int) []time.Time {

	//============================================
	// Set up first time range for first run...
	//============================================
	d1 := *R1
	d2 := d1 // just to define it, it will be set correctly below

	//============================================
	// Now iterate in chunks (of size: nMonths) and
	// save the recurrence dates...
	//============================================
	var m []time.Time
	for i := 0; i < 10000; i++ {
		mo, y := incMonths(d1.Month(), nMonths)
		d2 = time.Date(d1.Year()+y, mo, d1.Day(), d1.Hour(), d1.Minute(), d1.Second(), 0, time.UTC)

		//----------------------------------------
		// don't go outside the requested range
		//----------------------------------------
		if d1.After(*R2) || d1.Equal(*R2) {
			break
		}
		if d2.After(*R2) {
			d2 = *R2
		}
		//-------------------------------------------------------------------------------------
		//  the recurrence date will be the epoch date applied to d1.  Then see if this date
		//  is in the range d1 - d2
		//-------------------------------------------------------------------------------------
		d := time.Date(d1.Year(), d1.Month(), a1.Day(), d1.Hour(), d1.Minute(), d1.Second(), 0, time.UTC)

		//-------------------------------------------------------------------------------------
		// make sure it's in the interval range AND in its active timeframe, and if it is
		// then add it to the list
		//-------------------------------------------------------------------------------------
		if dateInRange(&d, &d1, &d2) && dateInRange(&d, a1, a2) {
			m = append(m, d)
		}
		d1 = d2 // on to the next interval
	}

	return m
}

func genYearlyRecurSeq(d, start, stop *time.Time, n int) []time.Time {
	var m []time.Time
	dt := *d
	for dateInRange(&dt, start, stop) {
		m = append(m, dt)
		dt = time.Date(dt.Year()+n, dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), 0, time.UTC)
	}
	return m
}

func (a *assessment) getRecurrences(start, stop *time.Time) []time.Time {
	var m []time.Time

	//-------------------------------------------
	// first ensure that the data is not bad...
	//-------------------------------------------
	if a.Start.After(a.Stop) {
		return m
	}

	//-------------------------------------------
	// next, ensure that the assessment falls in the time range...
	//-------------------------------------------
	if a.Frequency > RECURNONE &&
		(a.Stop.Equal(*start) || a.Stop.Before(*start) ||
			a.Start.After(*stop) || a.Start.Equal(*stop)) {
		return m
	}

	//-------------------------------------------
	// first insure that the data is not bad...
	//-------------------------------------------

	switch a.Frequency {
	case RECURNONE: // no recurrence
		b := &a.Start
		if dateInRange(b, start, stop) {
			m = append(m, a.Start)
			return m
		}
	case RECURDAILY: // daily
		d := start.Day()
		if start.Day() < a.Start.Day() {
			d = a.Start.Day()
		}
		dt := time.Date(start.Year(), start.Month(), d, a.Start.Hour(), a.Start.Minute(), a.Start.Second(), 0, time.UTC)
		return genRegularRecurSeq(&dt, &a.Stop, start, stop, 24*time.Hour)
	case RECURWEEKLY: // weekly
		dt := time.Date(start.Year(), start.Month(), start.Day(), a.Start.Hour(), a.Start.Minute(), a.Start.Second(), 0, time.UTC)
		return genRegularRecurSeq(&dt, &a.Stop, start, stop, 7*24*time.Hour)
	case RECURMONTHLY: // monthly
		// dt := time.Date(start.Year(), start.Month(), a.Start.Day(), a.Start.Hour(), a.Start.Minute(), a.Start.Second(), 0, time.UTC)
		return genMonthlyRecurSeq(&a.Start, &a.Stop, start, stop, 1)
	case RECURQUARTERLY: // quarterly
		dt := time.Date(start.Year(), start.Month(), a.Start.Day(), a.Start.Hour(), a.Start.Minute(), a.Start.Second(), 0, time.UTC)
		return genMonthlyRecurSeq(&dt, &a.Stop, start, stop, 3)
	case RECURYEARLY: // yearly
		dt := time.Date(start.Year(), a.Start.Month(), a.Start.Day(), a.Start.Hour(), a.Start.Minute(), a.Start.Second(), 0, time.UTC)
		return genYearlyRecurSeq(&dt, start, stop, 1)
	}
	return m
}
