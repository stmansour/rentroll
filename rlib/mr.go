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
		d         = startDt
		breakLoop = false
		nmd       time.Time // next month first date
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
		Console("Amount: %f, RMRID: %d, RTID: %d, startDt: %q, stopDt: %q\n",
			mr.MarketRate, mr.RMRID, mr.RTID, startDt, stopDt)

		for {
			yDiff, mDiff := 0, 1 // year would be same, month would be next (MOSTLY)
			if d.Month() == time.December {
				yDiff = 1 // d is being updated
				mDiff = 0
			}

			// jump to next month Start
			Console("\n\nCurrent Date: %q\n", d)
			nmd = d.AddDate(yDiff, mDiff, 1)
			Console("Next Date: %q\n\n\n", nmd)

			if nmd.After(stopDt) { // if nmd after stopDt then stop
				nmd = stopDt
				breakLoop = true
				Console("Yes, break the loop with nmd: %q\n", nmd)
			}

			diff := nmd.Sub(d)
			daysDiff := int(diff.Hours() / 24)
			perDayMR := mr.MarketRate / float64(DaysInMonth(d.Year(), d.Month()))
			amt += float64(daysDiff) * perDayMR // number of days * per day mr
			Console("daysDiff: %d, d.Diff:%d\n", daysDiff, DaysInMonth(d.Year(), d.Month()))
			Console("Per day marketRate: %f", perDayMR)

			d = nmd // update d with nmd (next month day)
			if breakLoop {
				break
			}
		}
	}

	Console("MarketRateValue: %f\n\n", amt)

	return
}
