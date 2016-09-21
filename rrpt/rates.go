package rrpt

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// RentableMarketRates prints a report of the rentable rid's rent rates between d1 and d2
func RentableMarketRates(xbiz *rlib.XBusiness, rid int64, d1, d2 *time.Time) {
	r := rlib.GetRentable(rid)
	m := rlib.GetRentableTypeRefsByRange(r.RID, d1, d2)

	fmt.Printf("RENTABLE RENT RATES\nRentable: %s  (%s)\nPeriod %s - %s\n\n", r.Name, r.IDtoString(), d1.Format(rlib.RRDATEFMT4), d2.Format(rlib.RRDATEFMT4))
	var tbl rlib.Table
	tbl.Init()
	tbl.AddColumn("Start", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Stop", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Rentable Type", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Rent Cycle", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Proration Cycle", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Gross Scheduled Rent", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)

	dt1 := *d1

	for i := 0; i < len(m); i++ { // spin through all the RentableTypes this rentable has been during d1 - d2
		// List all the rate information for type m[i].RTID
		dt2 := *d2 // this is the indexing end date, we track it to know when we're done
		for j := 0; j < len(xbiz.RT[m[i].RTID].MR); j++ {
			// Rate start time must be <= dt1 and its stop time must be > dt1
			if (xbiz.RT[m[i].RTID].MR[j].DtStart.Before(dt1) || xbiz.RT[m[i].RTID].MR[j].DtStart.Equal(dt1)) && xbiz.RT[m[i].RTID].MR[j].DtStop.After(dt1) {
				dt2 = xbiz.RT[m[i].RTID].MR[j].DtStop // start the indexing end date at the end of this MarketRate's period
				if dt2.After(*d2) {                   // if it's past our total end date...
					dt2 = *d2 // ...then snap to the end date
				}
				rcycle := xbiz.RT[m[i].RTID].RentCycle
				if m[i].OverrideRentCycle != 0 {
					rcycle = m[i].OverrideRentCycle
				}
				pcycle := xbiz.RT[m[i].RTID].Proration
				if m[i].OverrideProrationCycle != 0 {
					pcycle = m[i].OverrideProrationCycle
				}
				tbl.AddRow()
				tbl.Putd(-1, 0, dt1)
				tbl.Putd(-1, 1, dt2)
				tbl.Puts(-1, 2, xbiz.RT[m[i].RTID].Style)
				tbl.Puts(-1, 3, rlib.RentalPeriodToString(rcycle))
				tbl.Puts(-1, 4, rlib.RentalPeriodToString(pcycle))
				tbl.Putf(-1, 5, xbiz.RT[m[i].RTID].MR[j].MarketRate)

				dt1 = xbiz.RT[m[i].RTID].MR[j].DtStop // advance our start time pointer
			}
			if dt2.After(*d2) || dt2.Equal(*d2) { // if we're at the end of the requested period...
				break // ... then break out
			}
		}
		if dt2.After(*d2) || dt2.Equal(*d2) { // if we're at the end of the requested period...
			break // ... then break out
		}
	}

	tbl.TightenColumns()
	fmt.Print(tbl.SprintTable(rlib.TABLEOUTTEXT))
}
