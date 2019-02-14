package rlib

import (
	"context"
	"time"
)

// RSLeaseStatus is a slice of the string meaning of each LeaseStatus
// -- 0 = Not Leased, 1 = Leased, 2 = Reserved
var RSLeaseStatus = []string{
	"Not Leased", // 0
	"Leased",     // 1
	"Reserved",   // 2
}

// RSUseStatus is a slice of the string meaning of each UseStatus
// 0 = Ready, 1=InService, 2=Administrative, 3=Employee, 4=OwnerOccupied, 5=OfflineRennovation, 6=OfflineMaintenance, 7=Inactive(no longer a valid rentable)
var RSUseStatus = []string{
	"Ready",               // 0
	"In Service",          // 1
	"Administrative",      // 2
	"Employee",            // 3
	"Owner Occupied",      // 4
	"Offline Renovation",  // 5
	"Offline Maintenance", // 6
	"Inactive",            // 7
}

// RSMakeReadyStatus is a slice of the string meaning of each MakeReadyStatus
var RSMakeReadyStatus = []string{
	"Unknown",
	"In Progress Housekeeping",
	"In Progress Maintenance",
	"Pending Inspection",
	"Ready",
}

// RStatInfo encapsulates a RentableUseStatus structure along with t
// the associated amount for DtStart - DtStop
type RStatInfo struct {
	Amount float64
	RS     RentableUseStatus
}

// RStat returns the status for the supplied rentable during the Periods
// in gaps.  If the status is defined by RentableUseStatus entries those
// values will be used.
//
// If there is no database status available for the rentable during the
// requested period, a RentableUseStatus struct will be returned with the
// following values:
//   RSID:        0 (which further reinforces that it's not in the database)
//   LeaseStatus: VacantNotRented or VacantRented, depending on whether
//                there is a RentableUseStatus record for the rentable.
//   UseStatus:   InService
//
// INPUTS
//   ctx       - context.Context (sess, sql.Tx etc...)
//   BID       - which business
//   RID       - the rentable id
//   gaps      - slice of Periods of interest
//
// RETURNS
//   []RentableUseStatus
//             - slice of RStatus structs that defines the state
//               of the rentable during each Period in gaps
//   error     - any error encountered
//-----------------------------------------------------------------------------
func RStat(ctx context.Context, bid, rid int64, gaps []Period) ([]RStatInfo, error) {
	var (
		m   []RStatInfo
		err error
	)

	for i := 0; i < len(gaps); i++ {
		//-------------------------------------------------------------
		// Check for any special rentable status during the gap.
		// Reflect what's happening if we find anything
		//-------------------------------------------------------------
		rsa, err := GetRentableUseStatusByRange(ctx, rid, &gaps[i].D1, &gaps[i].D2)
		if err != nil {
			return m, err
		}

		if len(rsa) > 0 {
			for j := 0; j < len(rsa); j++ {
				var rs RStatInfo
				rsa[j].DtStart = gaps[i].D1
				rsa[j].DtStop = gaps[i].D2
				rs.RS = rsa[j]
				m = append(m, rs)
			}
		} else {
			var rs = RStatInfo{
				RS: RentableUseStatus{
					BID:       bid,
					RID:       rid,
					DtStart:   gaps[i].D1,
					DtStop:    gaps[i].D2,
					UseStatus: USESTATUSready,
				},
			}
			// //----------------------------------------------------------------
			// // If there is a RentalAgreement in the future, modify the status
			// //----------------------------------------------------------------
			// r, err := GetRentableLeaseStatusOnOrAfter(ctx, rid, &gaps[i].D1)
			// if err != nil {
			// 	return m, err
			// }
			//
			// if r.RLID > 0 {
			// 	rs.RS.LeaseStatus = LEASESTATUSnotleased
			// }
			m = append(m, rs)
		}
	}
	return m, err
}

// VacancyGSR returns the GSR amount for the rentable during the supplied
// period.  This value is used as the Period GSR value in the the RentRoll
// view / report.
//
// INPUT
//  xbiz  - the business
//  rid   - which rentable
//  d1-d2 - the time period
//
// RETURN
//  the amount of GSR for the period
//-----------------------------------------------------------------------------
func VacancyGSR(ctx context.Context, xbiz *XBusiness, rid int64, d1, d2 *time.Time) (float64, error) {
	var err error
	amt := float64(0)
	// Console("*** Calling VacancyDetect: %s - %s, rid = %d\n", d1.Format(RRDATEFMTSQL), d2.Format(RRDATEFMTSQL), rid)
	m, err := VacancyDetect(ctx, xbiz, d1, d2, rid)
	if err != nil {
		return amt, err
	}

	for i := 0; i < len(m); i++ {
		amt += m[i].Amount
	}

	return amt, err
}

// LeaseStatusStringer returns the string associated with the LeaseStatus
// in struct t.
//-----------------------------------------------------------------------------
func (t *RentableLeaseStatus) LeaseStatusStringer() string {
	return LeaseStatusString(t.LeaseStatus)
}

// LeaseStatusString returns the string associated with LeaseStatus ls
//-----------------------------------------------------------------------------
func LeaseStatusString(ls int64) string {
	i := int(ls)
	if i > len(RSLeaseStatus) {
		i = 0
	}
	return RSLeaseStatus[i]
}

// UseStatusStringer returns the string associated with the UseStatus
// in struct t.
//-----------------------------------------------------------------------------
func (t *RentableUseStatus) UseStatusStringer() string {
	return UseStatusString(t.UseStatus)
}

// UseStatusString returns the string associated with UseStatus us
//-----------------------------------------------------------------------------
func UseStatusString(us int64) string {
	i := int(us)
	if i > len(RSUseStatus) {
		i = 0
	}
	return RSUseStatus[i]
}

// MakeReadyStatusStringer returns the string associated with the MakeReadyStatus
// in struct t.
//-----------------------------------------------------------------------------
func (t *Rentable) MakeReadyStatusStringer() string {
	return MakeReadyStatusString(t.MRStatus)
}

// MakeReadyStatusString returns the string associated with MakeReadyStatus mr
//-----------------------------------------------------------------------------
func MakeReadyStatusString(mr int64) string {
	i := int(mr)
	if i > len(RSMakeReadyStatus) {
		i = 0
	}
	return RSMakeReadyStatus[i]
}

// SetRentableLeaseStatusAbbr changes the use status from d1 to d2 to the supplied
// status, us. It adds and modifies existing records as needed.
//
// INPUTS
//     ctx - db context
//     bid - which business
//     rid - which rentable
//     us  - new lease status
//     d1  - start time for status us
//     d2  - stop time for status us
//     res - if true, all the record beginning at d1 will be set to RESERVED
//-----------------------------------------------------------------------------
func SetRentableLeaseStatusAbbr(ctx context.Context, bid, rid, us int64, d1, d2 *time.Time, res bool) error {
	var b = RentableLeaseStatus{
		RID:         rid,
		BID:         bid,
		DtStart:     *d1,
		DtStop:      *d2,
		Comment:     "",
		LeaseStatus: us,
	}

	return SetRentableLeaseStatus(ctx, &b, res)
}

// SetRentableLeaseStatus implements the proper insertion of a reservation
//     under all the circumstances considered (these are described in detail in
//     https://docs.google.com/presentation/d/1v3eEvATppP501MVM6vjv4VoQBDgZo-gq_wPPqUhblV4/edit#slide=id.g4f52e75848_0_0
//     in the slide Inserting A Reservation [d1-d2] Into LeaseStatus).
//
// INPUTS
//     ctx - db context
//     rls - the new reservation structure
//     res - if true, all the record beginning at d1 will be set to RESERVED
//-----------------------------------------------------------------------------
func SetRentableLeaseStatus(ctx context.Context, rls *RentableLeaseStatus, res bool) error {
	// funcname := "SetRentableLeaseStatus"
	// Console("Entered %s.  range = %s, LeaseStatus = %d\n", funcname, ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.LeaseStatus)

	var err error
	d1 := rls.DtStart
	d2 := rls.DtStop
	a, err := GetRentableLeaseStatusByRange(ctx, rls.RID, &d1, &d2)
	if err != nil {
		return err
	}

	// Console("%s: Range = %s    found %d records\n", funcname, ConsoleDRange(&d1, &d2), len(a))

	//--------------------------------------------------------------------------
	// If a has more than 2 entries, we can simply delete all but the first and
	// last entries.
	//--------------------------------------------------------------------------
	la := len(a) // save original length
	if la > 2 {
		for i := 1; i <= la-2; i++ {
			// Console("%s: deleting RLID = %d\n", funcname, a[i].RLID)
			if err = DeleteRentableLeaseStatus(ctx, a[i].RLID); err != nil {
				return err
			}
		}
	}

	// Deal with some edge cases...
	// Console("%s: la = %d\n", funcname, la)
	switch la {
	case 0:
		// Console("%s: case 0\n", funcname)
		// nothing to do
		break
	case 1:
		// Console("%s: case 1, RLID=%d, %s LeaseStatus = %d\n", funcname, a[0].RLID, ConsoleDRange(&a[0].DtStart, &a[0].DtStop), a[0].LeaseStatus)
		before := a[0].DtStart.Before(d1)
		after := a[0].DtStop.After(d2)
		// Console("before = %t, after = %t\n", before, after)
		//--------------------------------------------------------------------------
		// If the LeaseStatus extends before AND after d1,d2 then split into two
		// statuses -- one before and one after d1,d2
		//--------------------------------------------------------------------------
		if before && after {
			b1 := a[0]     // earlier split
			b1.DtStop = d1 // update existing and stop it at the reservation start
			// Console("A: Updating RLID=%d, %s  LeaseStatus = %d\n", b1.RLID, ConsoleDRange(&b1.DtStart, &b1.DtStop), b1.LeaseStatus)
			err = UpdateRentableLeaseStatus(ctx, &b1)
			if err != nil {
				return err
			}

			b2 := a[0]      // later split
			b2.DtStart = d2 // start it at the reservation end
			b2.RLID = 0     // this is a new record
			if res {
				b2.LeaseStatus = LEASESTATUSreserved
			}

			// Console("B: Inserting %s  LeaseStatus = %d\n", ConsoleDRange(&b2.DtStart, &b2.DtStop), b2.LeaseStatus)
			_, err = InsertRentableLeaseStatus(ctx, &b2)
			if err != nil {
				return err
			}
		} else if before {
			// set its stop date to d1
			b1 := a[0]
			b1.DtStop = d1
			// Console("C: Updating RLID=%d, %s  LeaseStatus = %d\n", b1.RLID, ConsoleDRange(&b1.DtStart, &b1.DtStop), b1.LeaseStatus)
			err = UpdateRentableLeaseStatus(ctx, &b1)
			if err != nil {
				return err
			}
		} else if after {
			// set its stop date to d1
			b1 := a[0]
			b1.DtStart = d2
			if res {
				b1.LeaseStatus = LEASESTATUSreserved
			}
			// Console("D: Updating RLID=%d, %s  LeaseStatus = %d\n", b1.RLID, ConsoleDRange(&b1.DtStart, &b1.DtStop), b1.LeaseStatus)
			err = UpdateRentableLeaseStatus(ctx, &b1)
			if err != nil {
				return err
			}
		} else {
			// if we hit this point, then the record that was found is totally
			// overwritten by the new record. So, we will just delete this one.
			// Console("E: Deleting RLID=%d, %s  LeaseStatus = %d\n", a[0].RLID, ConsoleDRange(&a[0].DtStart, &a[0].DtStop), a[0].LeaseStatus)
			if err = DeleteRentableLeaseStatus(ctx, a[0].RLID); err != nil {
				return err
			}
		}
	default: // there were at least 2, trim first and last...
		// Console("%s: case (default)\n", funcname)
		// if la == 2 {
		// 	for i := 0; i < 2; i++ {
		// 		// Console("[%d] - %s\n", i, ConsoleDRange(&a[i].DtStart, &a[i].DtStop))
		// 	}
		// }
		if la == 2 &&
			(a[0].DtStart.Equal(d1) || a[0].DtStart.After(d1)) &&
			(a[0].DtStop.Equal(d2) || a[0].DtStop.Before(d2)) {
			if err = DeleteRentableLeaseStatus(ctx, a[0].RLID); err != nil {
				return err
			}
		} else {
			if a[0].DtStart.Before(d1) {
				a[0].DtStop = d1
				// Console("F: Updating RLID=%d, %s  LeaseStatus = %d\n", a[0].RLID, ConsoleDRange(&a[0].DtStart, &a[0].DtStop), a[0].LeaseStatus)
				err = UpdateRentableLeaseStatus(ctx, &a[0])
				if err != nil {
					return err
				}
			}
			if a[la-1].DtStop.After(d2) {
				a[la-1].DtStart = d2
				if res {
					a[la-1].LeaseStatus = LEASESTATUSreserved
				}
				// Console("G: Updating RLID=%d, %s  LeaseStatus = %d\n", a[la-1].RLID, ConsoleDRange(&a[la-1].DtStart, &a[la-1].DtStop), a[la-1].LeaseStatus)
				err = UpdateRentableLeaseStatus(ctx, &a[la-1])
				if err != nil {
					return err
				}
			}
		}
	}

	// Console("%s: Inserting %s LeaseStatus = %d\n", funcname, ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.LeaseStatus)

	_, err = InsertRentableLeaseStatus(ctx, rls)
	return err
}

// SetRentableUseStatusAbbr changes the use status from d1 to d2 to the supplied
// status, us. It adds and modifies existing records as needed.  If the
// rentable is already in the supplied use state for time range d1-d2 then
// no changes are made to the database.
//
// In the diagram, dt1 = lease status DtStart, dt2 = Dt
//                     d1               d2
//   +--------+        |   +--------+   |         +--------+
//   | case 1 |        |   | case 2 |   |         | case 3 |
//   +--------+        |   +--------+   |         +--------+
//  dt2      dt2       |  dt1      dt2  |        dt1      dt2
//                     |                |
//        +---------------+         +------------------+
//        |    case 4     |         |       case 5     |
//        +---------------+         +------------------+
//       dt1           | dt2       dt1  |             dt2
//                     |                |
//        +--------------------------------------------+
//        |                  case 6                    |
//        +--------------------------------------------+
//       dt1           |                |             dt2
//                     |                |
//
// INPUTS
//     ctx - db context
//     bid - which business
//     rid - which rentable
//     us  - new use status
//     d1  - start time for status us
//     d2  - stop time for status us
//-----------------------------------------------------------------------------
func SetRentableUseStatusAbbr(ctx context.Context, bid, rid, us int64, d1, d2 *time.Time) error {
	var b = RentableUseStatus{
		RID:       rid,
		BID:       bid,
		DtStart:   *d1,
		DtStop:    *d2,
		Comment:   "",
		UseStatus: us,
	}
	return SetRentableUseStatus(ctx, &b)

}

// SetRentableUseStatus implements the proper insertion of a use status
//     under all the circumstances considered (these are described in detail in
//     https://docs.google.com/presentation/d/1v3eEvATppP501MVM6vjv4VoQBDgZo-gq_wPPqUhblV4/edit#slide=id.g4f52e75848_0_0
//     in the slide Inserting A Reservation [d1-d2] Into LeaseStatus). UseStatus
//     works exactly the same w2layout
//
// INPUTS
//     ctx - db context
//     rls - the new RentableUseStatus structure
//-----------------------------------------------------------------------------
func SetRentableUseStatus(ctx context.Context, rls *RentableUseStatus) error {
	// funcname := "SetRentableUseStatus"
	// Console("Entered %s\n", funcname)

	var err error
	d1 := rls.DtStart
	d2 := rls.DtStop
	a, err := GetRentableUseStatusByRange(ctx, rls.RID, &d1, &d2)
	if err != nil {
		return err
	}

	// Console("%s: Range = %s    found %d records\n", funcname, ConsoleDRange(&d1, &d2), len(a))

	//--------------------------------------------------------------------------
	// If a has more than 2 entries, we can simply delete all but the first and
	// last entries.
	//--------------------------------------------------------------------------
	la := len(a) // save original length
	if la > 2 {
		for i := 1; i <= la-2; i++ {
			// Console("%s: deleting RLID = %d\n", funcname, a[i].RLID)
			if err = DeleteRentableUseStatus(ctx, a[i].RSID); err != nil {
				return err
			}
		}
	}

	// Deal with some edge cases...
	switch la {
	case 0:
		// nothing to do
		break
	case 1:
		// Console("%s: case 1\n", funcname)
		before := a[0].DtStart.Before(d1)
		after := a[0].DtStop.After(d2)
		//--------------------------------------------------------------------------
		// If the LeaseStatus extends before AND after d1,d2 then split into two
		// statuses -- one before and one after d1,d2
		//--------------------------------------------------------------------------
		if before && after {
			b1 := a[0]     // earlier split
			b1.DtStop = d1 // update existing and stop it at the reservation start
			err = UpdateRentableUseStatus(ctx, &b1)
			if err != nil {
				return err
			}

			b2 := a[0]      // later split
			b2.DtStart = d2 // start it at the reservation end
			b2.RSID = 0     // this is a new record
			_, err = InsertRentableUseStatus(ctx, &b2)
			if err != nil {
				return err
			}
		} else if before {
			// set its stop date to d1
			b1 := a[0]
			b1.DtStop = d1
			err = UpdateRentableUseStatus(ctx, &b1)
			if err != nil {
				return err
			}
		} else if after {
			// set its stop date to d1
			b1 := a[0]
			b1.DtStart = d2
			err = UpdateRentableUseStatus(ctx, &b1)
			if err != nil {
				return err
			}
		}
	default: // there were at least 3, trim first and last...
		if a[0].DtStart.Before(d1) {
			a[0].DtStop = d1
			err = UpdateRentableUseStatus(ctx, &a[0])
			if err != nil {
				return err
			}
		}
		if a[la-1].DtStop.After(d2) {
			a[la-1].DtStart = d2
			err = UpdateRentableUseStatus(ctx, &a[la-1])
			if err != nil {
				return err
			}
		}
	}

	// Console("%s: Inserting RentableUseStatus = %d, %s\n", funcname, rls.LeaseStatus, ConsoleDRange(&rls.DtStart, &rls.DtStop))

	_, err = InsertRentableUseStatus(ctx, rls)
	return err
}
