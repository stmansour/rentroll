package rlib

import (
	"time"
)

// DaysInMonth returns the number of days in a given month of a year
func DaysInMonth(year int, m time.Month) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// // MonthLastDate returns the last date of a given month of a year
// func MonthLastDate(year int, m time.Month) time.Time {
//  monthDays := DaysInMonth(year, m)
//  return time.Date(year, m+1, monthDays, 0, 0, 0, 0, time.UTC)
// }

// GetMRAmtInDateRange accepts group of rmrid, startDate and stopDate
// It calculates final marketRate value from different market rate objects
// in given date range with startDt, stopDt values
func GetMRAmtInDateRange(rmridList []int64, startDt, stopDt time.Time) (amt float64, err error) {
	const funcname = "GetMRAmtInDateRange"
	var (
		breakLoop = false
		ied       time.Time // interval end date
		mr        RentableMarketRate
	)
	Console("Entered in %s\n", funcname)

	Console("list of rmrid:= %#v\n", rmridList)

	// loop over each market rate ID
	for _, rmrid := range rmridList {

		// get marketRate object from RMRID
		mr, err = GetRentableMarketRateInstance(rmrid)
		if err != nil {
			return
		}
		Console("Amount: %f, RMRID: %d, RTID: %d, startDt: %q, stopDt: %q, DtStart: %q, DtStop: %q\n",
			mr.MarketRate, mr.RMRID, mr.RTID, startDt, stopDt, mr.DtStart, mr.DtStop)

		// if market rate time range not overlaps then continue...
		if !DateRangeOverlap(&mr.DtStart, &mr.DtStop, &startDt, &stopDt) {
			continue
		}

		isd := mr.DtStart // interval start date
		if mr.DtStart.Before(startDt) {
			isd = startDt
		}

		for {
			yDiff, mDiff := 0, 1 // year would be same, month would be next (MOSTLY)
			if isd.Month() == time.December {
				yDiff = 1 // d is being updated
				mDiff = 0
			}

			// jump to next month Start
			Console("\n\nCurrent Date: %q\n", isd)
			ied = isd.AddDate(yDiff, mDiff, 1)
			Console("Next Date: %q\n\n\n", ied)

			if ied.After(stopDt) { // if ied after stopDt then stop
				ied = stopDt
				breakLoop = true
				Console("Yes, break the loop with ied: %q\n", ied)
			}

			diff := ied.Sub(isd)
			daysDiff := int(diff.Hours() / 24)
			perDayMR := mr.MarketRate / float64(DaysInMonth(isd.Year(), isd.Month()))
			amt += float64(daysDiff) * perDayMR // number of days * per day mr
			Console("daysDiff: %d, d.Diff:%d\n", daysDiff, DaysInMonth(isd.Year(), isd.Month()))
			Console("Per day marketRate: %f", perDayMR)

			isd = ied // update d with ied (next month day)
			if breakLoop {
				break
			}
		}
	}

	Console("MarketRateValue: %f\n\n", amt)

	return
}
