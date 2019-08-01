package main

import (
	"context"
	"fmt"
	"math/rand"
	"rentroll/bizlogic"
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

	//--------------------------------
	// set the time as 1:00 AM
	// This can be replaced when we save the proper epoch values for each RT
	//--------------------------------
	epoch := time.Now()
	if epoch.Hour() > 1 {
		epoch = epoch.AddDate(0, 0, 1) // use tommorrow's date
	}

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
					BID:              1,
					RID:              r.RID,
					LeaseStatus:      rlib.LEASESTATUSreserved,
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

				//-------------------------------------------------
				// create a Transactant...
				//-------------------------------------------------
				fn := GenerateRandomFirstName()
				ln := GenerateRandomLastName()
				var t = rlib.Transactant{
					BID:          ls.BID,
					FirstName:    fn,
					LastName:     ln,
					PrimaryEmail: GenerateRandomEmail(ln, fn),
					CellPhone:    GenerateRandomPhoneNumber(),
					Address:      GenerateRandomAddress(),
					City:         GenerateRandomCity(),
					State:        GenerateRandomState(),
					PostalCode:   fmt.Sprintf("%05d", rand.Intn(100000)),
					Country:      "USA",
				}
				if _, err = rlib.InsertTransactant(ctx, &t); err != nil {
					return err
				}

				//-----------------------------------------
				// CREATE COMPANIES...
				//-----------------------------------------
				if IG.Rand.Intn(100) < 30 {
					var t rlib.Transactant
					t.CompanyName = GenerateRandomCompany()
					t.Address = GenerateRandomAddress()
					t.City = GenerateRandomCity()
					t.State = GenerateRandomState()
					t.Country = "USA"
					t.PostalCode = fmt.Sprintf("%05d", rand.Intn(100000))
					t.PrimaryEmail = GenerateRandomCompanyEmail(t.CompanyName)
					t.WorkPhone = GenerateRandomPhoneNumber()
					t.IsCompany = true
					if _, err := rlib.InsertTransactant(ctx, &t); err != nil {
						return err
					}
				}

				//-------------------------------------------------
				// create a RentalAgreement...
				//-------------------------------------------------
				now := rlib.Now()
				var ra = rlib.RentalAgreement{
					BID:                 ls.BID,
					AgreementStart:      now,
					AgreementStop:       ls.DtStop,
					PossessionStart:     ls.DtStart,
					PossessionStop:      ls.DtStop,
					RentStart:           ls.DtStart,
					RentStop:            ls.DtStop,
					RentCycleEpoch:      epoch,
					FLAGS:               0,
					CSAgent:             GenerateValidUID(),
					UnspecifiedAdults:   int64(1 + IG.Rand.Intn(3)),
					UnspecifiedChildren: int64(IG.Rand.Intn(3)),
				}
				if _, err = rlib.InsertRentalAgreement(ctx, &ra); err != nil {
					return err
				}

				//-------------------------------------
				// Assign the Rentable
				//-------------------------------------
				var rar rlib.RentalAgreementRentable
				rar.BID = ls.BID
				rar.RARDtStart = ls.DtStart
				rar.RARDtStop = ls.DtStop
				rar.RAID = ra.RAID
				rar.RID = ls.RID
				if _, err = rlib.InsertRentalAgreementRentable(ctx, &rar); err != nil {
					return err
				}

				//-------------------------------------
				// Assign Payor
				//-------------------------------------
				var rap rlib.RentalAgreementPayor
				rap.BID = ls.BID
				rap.DtStart = ls.DtStart
				rap.DtStop = ls.DtStop
				rap.RAID = ra.RAID
				rap.TCID = t.TCID
				if _, err = rlib.InsertRentalAgreementPayor(ctx, &rap); err != nil {
					return err
				}

				//-------------------------------------
				// Create the deposit...
				// assume it's $50 / night
				//-------------------------------------
				var asm rlib.Assessment
				asm.ARID = dbConf.ResDepARID
				asm.Amount = float64(bk * 50)
				asm.RAID = ra.RAID
				asm.BID = ls.BID
				asm.RID = ls.RID
				asm.Start = now
				asm.Stop = asm.Start
				asm.RentCycle = rlib.RECURNONE
				asm.ProrationCycle = rlib.RECURNONE
				// if _, err = rlib.InsertAssessment(ctx, &asm); err != nil {
				// 	return err
				// }
				be := bizlogic.InsertAssessment(ctx, &asm, 1, &noClose)
				if be != nil {
					return bizlogic.BizErrorListToError(be)
				}

				//-----------------------------------------------------
				// Create a RentalAgreement Ledger marker
				//-----------------------------------------------------
				var lm = rlib.LedgerMarker{
					BID:     ra.BID,
					RAID:    ra.RAID,
					RID:     0,
					Dt:      epoch.AddDate(0, -1, 0),
					Balance: float64(0),
					State:   rlib.LMINITIAL,
				}
				if _, err = rlib.InsertLedgerMarker(ctx, &lm); err != nil {
					return err
				}

				//-------------------------------------
				// Ledger for the RAID's rentable
				//-------------------------------------
				lm.RID = ls.RID
				if _, err = rlib.InsertLedgerMarker(ctx, &lm); err != nil {
					return err
				}
				// rlib.Console("Created LedgerMarker for RAID=%d, RID=%d on %s\n", lm.RAID, lm.RID, rlib.ConDt(&lm.Dt))

				//-------------------------------------
				// Assign Users...
				//-------------------------------------
				var rau rlib.RentableUser
				rau.BID = ls.BID
				rau.RID = ls.RID
				rau.DtStart = ls.DtStart
				rau.DtStop = ls.DtStop
				rau.TCID = t.TCID
				if _, err = rlib.InsertRentableUser(ctx, &rau); err != nil {
					return err
				}

				//-------------------------------------
				// Now create the Lease Status (reservation)
				//-------------------------------------
				ls.RAID = ra.RAID
				if err = rlib.SetRentableLeaseStatus(ctx, &ls); err != nil {
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
