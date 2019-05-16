package main

import (
	"context"
	"fmt"
	"math/rand"
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
				fn := GenerateRandomFirstName()
				ln := GenerateRandomLastName()
				var ls = rlib.RentableLeaseStatus{ // we start with this info, then check and modify as needed
					BID:              1,
					RID:              r.RID,
					LeaseStatus:      rlib.LEASESTATUSreserved,
					FirstName:        fn,
					LastName:         ln,
					Email:            GenerateRandomEmail(ln, fn),
					Phone:            GenerateRandomPhoneNumber(),
					Address:          GenerateRandomAddress(),
					City:             GenerateRandomCity(),
					State:            GenerateRandomState(),
					PostalCode:       fmt.Sprintf("%05d", rand.Intn(100000)),
					Country:          "USA",
					ConfirmationCode: rlib.GenerateUserRefNo(),
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
						break // nothing to conflict with, this time is ok
					}
					overlap := false
					for i := 0; i < len(m); i++ {
						if rlib.DateRangeOverlap(&ls.DtStart, &ls.DtStop, &m[i].DtStart, &m[i].DtStop) {
							overlap = true // cant use this
							break          // breaks out of inner loop
						}
					}
					available = !overlap // if no overlaps were found, the outer loope terminates
				}

				// rlib.Console("Found date range that works: %s\n", rlib.ConsoleDRange(&ls.DtStart, &ls.DtStop))
				m = append(m, ls) // we'll use this time
				if err := rlib.SetRentableLeaseStatus(ctx, &ls); err != nil {
					return err
				}
				//--------------------------------------------------
				// Add some reasonable UseStatus values as well...
				// 7 - 5 hours prior  Housekeeping
				// 4 - 2 hours prior  Ready
				// (selected range)   InService
				//--------------------------------------------------
				dd := ls.DtStart.Unix()
				hrs := 300 + IG.Rand.Intn(180)
				d1 := time.Unix(dd-int64(hrs), int64(0)) // Housekeeping
				mns := 120
				d2 := time.Unix(dd-int64(mns), int64(0)) // ready
				var us = rlib.RentableUseStatus{
					BID:       1,
					RID:       r.RID,
					UseStatus: rlib.USESTATUShousekeeping,
					DtStart:   d1,
					DtStop:    d2,
				}
				if err := rlib.SetRentableUseStatus(ctx, &us); err != nil { // sets Housekeeping status
					return err
				}
				us.UseStatus = rlib.USESTATUSready
				us.DtStart = d2
				us.DtStop = ls.DtStart
				if err := rlib.SetRentableUseStatus(ctx, &us); err != nil { // sets ready status
					return err
				}
				us.UseStatus = rlib.USESTATUSinService
				us.DtStart = ls.DtStart
				us.DtStop = ls.DtStop
				if err := rlib.SetRentableUseStatus(ctx, &us); err != nil { // sets in use status
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
