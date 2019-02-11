package main

import (
	"context"
	"rentroll/rlib"
	"time"
)

// HotelBookings creates randomized hotel bookings for the rentables where
// the RentableType FLAGS bit 3 is set to 0.
//
// INPUTS:
//    ctx    - database context for txn
//    dbconf - the config info for creating this db
//
// RETURNS:
//    err - any errors encountered
//------------------------------------------------------------------------------
func HotelBookings(ctx context.Context, dbConf *GenDBConf) error {
	rlib.Console("Entered HotelBookings\n")

	q := "SELECT " + rlib.RRdb.DBFields["Rentable"] + " from Rentable WHERE BID=1;"
	rows, err := rlib.RRdb.Dbrr.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()

	totaldays := int(dbConf.HotelReserveDtStop.Sub(dbConf.HotelReserveDtStart) / (time.Hour * 24))
	bookdays := int(float64(totaldays) * dbConf.HotelReservePct)
	rlib.Console("Days in booking period: %d\n", totaldays)
	rlib.Console("Target days to book each hotel room: = %d\n", bookdays)

	for rows.Next() {
		var r rlib.Rentable
		if err := rlib.ReadRentables(rows, &r); err != nil {
			return err
		}
		var m []rlib.RentableTypeRef
		m, err = rlib.GetRentableTypeRefsByRange(ctx, r.RID, &dbConf.HotelReserveDtStart, &dbConf.HotelReserveDtStop)
		if err != nil {
			return err
		}

		for j := 0; j < len(m); j++ {
			if dbConf.xbiz.RT[m[j].RTID].FLAGS&(1<<3) != 0 {
				continue
			}

			// rlib.Console("Booking %s (RID=%d)\n", r.RentableName, r.RID)

			//------------------------------------------------------
			// book it for approximately the target number of days
			//------------------------------------------------------
			var m []rlib.RentableLeaseStatus
			for booked := 0; booked < bookdays; {
				bk := 1 + IG.Rand.Intn(7) // length of stay
				// rlib.Console("Reserve for %d days,  total = %d\n", bk, booked)

				//------------------------------------------------------
				// Add RentableLeaseStatus - find some available time.
				// start with a guess on the time range
				//------------------------------------------------------
				var ls = rlib.RentableLeaseStatus{ // we start with this info, then check and modify as needed
					BID:         1,
					RID:         r.RID,
					LeaseStatus: rlib.LEASESTATUSreserved,
				}

				//--------------------------------------------------------------
				// Choose a random time range and make sure it does not overlap
				// anything else
				//--------------------------------------------------------------
				for available := false; !available; {
					// rlib.Console("*\n")
					x := IG.Rand.Intn(totaldays - bk)
					ls.DtStart = dbConf.HotelReserveDtStart.AddDate(0, 0, x)
					ls.DtStop = dbConf.HotelReserveDtStart.AddDate(0, 0, x+bk)
					// rlib.Console("Attempt to reserve: %s\n", rlib.ConsoleDRange(&ls.DtStart, &ls.DtStop))
					// rlib.Console("len(m) = %d\n", len(m))
					if len(m) == 0 {
						available = true
					}
					for i := 0; i < len(m); i++ {
						if !rlib.DateRangeOverlap(&ls.DtStart, &ls.DtStop, &m[i].DtStart, &m[i].DtStop) {
							available = true
							break
						}
					}
				}
				// rlib.Console("Found date range that works: %s\n", rlib.ConsoleDRange(&ls.DtStart, &ls.DtStop))
				m = append(m, ls) // we'll use this time
				if err := rlib.SetRentableLeaseStatusAbbr(ctx, ls.BID, ls.RID, ls.LeaseStatus, &ls.DtStart, &ls.DtStop, false); err != nil {
					return err
				}
				// rlib.Console("Scheduled: RID: %d,  %s\n", ls.RID, rlib.ConsoleDRange(&ls.DtStart, &ls.DtStop))
				booked += bk
			}
		}

		if false {
			var n []rlib.RentableLeaseStatus
			n, err := rlib.GetRentableLeaseStatusByRange(ctx, r.RID, &dbConf.HotelReserveDtStart, &dbConf.HotelReserveDtStop)
			if err != nil {
				return err
			}
			for i := 0; i < len(n); i++ {
				rlib.Console("%d. %s\n", i, rlib.ConsoleDRange(&n[i].DtStart, &n[i].DtStop))
			}
		}
	}
	rlib.Console("Normal Exit: HotelBookings\n")
	return nil
}
