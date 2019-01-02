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
					BID:         bid,
					RID:         rid,
					DtStart:     gaps[i].D1,
					DtStop:      gaps[i].D2,
					UseStatus:   USESTATUSready,
					LeaseStatus: LEASESTATUSnotleased,
				},
			}
			//----------------------------------------------------------------
			// If there is a RentalAgreement in the future, modify the status
			//----------------------------------------------------------------
			r, err := GetRentableUseStatusOnOrAfter(ctx, rid, &gaps[i].D1)
			if err != nil {
				return m, err
			}

			if r.RSID > 0 {
				rs.RS.LeaseStatus = LEASESTATUSnotleased
			}
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
func (t *RentableUseStatus) LeaseStatusStringer() string {
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

// SetRentableLeaseStatus changes the use status from d1 to d2 to the supplied
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
//     us  - new lease status
//     d1  - start time for status us
//     d2  - stop time for status us
//-----------------------------------------------------------------------------
func SetRentableLeaseStatus(ctx context.Context, bid, rid, us int64, d1, d2 *time.Time) error {
	a, err := GetRentableLeaseStatusByRange(ctx, rid, d1, d2)
	if err != nil {
		return err
	}

	if len(a) == 0 {
		var b = RentableLeaseStatus{
			RID:         rid,
			BID:         bid,
			DtStart:     *d1,
			DtStop:      *d2,
			Comment:     "",
			LeaseStatus: us,
		}
		_, err = InsertRentableLeaseStatus(ctx, &b)
		return err
	}

	if len(a) == 1 {
		b := a[0]
		dtStartPrior := b.DtStart.Before(*d1)
		dtStartEqual := b.DtStart.Equal(*d1)
		dtStartAfter := b.DtStart.After(*d1)
		dtStopPrior := b.DtStop.Before(*d2)
		dtStopEqual := b.DtStop.Equal(*d2)
		dtStopAfter := b.DtStop.After(*d2)

		b.DtStart = *d1
		b.DtStop = *d2
		b.LeaseStatus = us

		if (dtStartEqual || dtStartAfter) && (dtStopPrior || dtStopEqual) { // case 2
			return UpdateRentableLeaseStatus(ctx, &b)
		}

		if a[0].DtStop.Before(*d1) || a[0].DtStop.Equal(*d1) || a[0].DtStart.After(*d2) || a[0].DtStart.Equal(*d2) { // case 1 or case 3
			_, err = InsertRentableLeaseStatus(ctx, &b)
			return err
		}

		if dtStartPrior && dtStopPrior { // case 4
			a[0].DtStop = *d1
			if err = UpdateRentableLeaseStatus(ctx, &a[0]); err != nil {
				return err
			}
			_, err = InsertRentableLeaseStatus(ctx, &b)
			return err
		}

		if dtStopAfter && dtStartAfter { // case 5
			a[0].DtStart = *d2
			if err = UpdateRentableLeaseStatus(ctx, &a[0]); err != nil {
				return err
			}
			_, err = InsertRentableLeaseStatus(ctx, &b)
			return err
		}

		if dtStartPrior && dtStopAfter { // case 6
			tmp := a[0].DtStop
			a[0].DtStop = *d1
			if err = UpdateRentableLeaseStatus(ctx, &a[0]); err != nil {
				return err
			}
			a[0].DtStop = tmp
			a[0].DtStart = *d2
			_, err = InsertRentableLeaseStatus(ctx, &a[0])
			if err != nil {
				return err
			}
			_, err = InsertRentableLeaseStatus(ctx, &b)
			return err
		}
	}
	return nil
}

// SetRentableUseStatus changes the use status from d1 to d2 to the supplied
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
//     us  - new lease status
//     d1  - start time for status us
//     d2  - stop time for status us
//-----------------------------------------------------------------------------
func SetRentableUseStatus(ctx context.Context, bid, rid, us int64, d1, d2 *time.Time) error {
	a, err := GetRentableUseStatusByRange(ctx, rid, d1, d2)
	if err != nil {
		return err
	}

	if len(a) == 0 {
		var b = RentableUseStatus{
			RID:       rid,
			BID:       bid,
			DtStart:   *d1,
			DtStop:    *d2,
			Comment:   "",
			UseStatus: us,
		}
		_, err = InsertRentableUseStatus(ctx, &b)
		return err
	}

	if len(a) == 1 {
		b := a[0]
		dtStartPrior := b.DtStart.Before(*d1)
		dtStartEqual := b.DtStart.Equal(*d1)
		dtStartAfter := b.DtStart.After(*d1)
		dtStopPrior := b.DtStop.Before(*d2)
		dtStopEqual := b.DtStop.Equal(*d2)
		dtStopAfter := b.DtStop.After(*d2)

		b.DtStart = *d1
		b.DtStop = *d2
		b.UseStatus = us

		if (dtStartEqual || dtStartAfter) && (dtStopPrior || dtStopEqual) { // case 2
			return UpdateRentableUseStatus(ctx, &b)
		}

		if a[0].DtStop.Before(*d1) || a[0].DtStop.Equal(*d1) || a[0].DtStart.After(*d2) || a[0].DtStart.Equal(*d2) { // case 1 or case 3
			_, err = InsertRentableUseStatus(ctx, &b)
			return err
		}

		if dtStartPrior && dtStopPrior { // case 4
			a[0].DtStop = *d1
			if err = UpdateRentableUseStatus(ctx, &a[0]); err != nil {
				return err
			}
			_, err = InsertRentableUseStatus(ctx, &b)
			return err
		}

		if dtStopAfter && dtStartAfter { // case 5
			a[0].DtStart = *d2
			if err = UpdateRentableUseStatus(ctx, &a[0]); err != nil {
				return err
			}
			_, err = InsertRentableUseStatus(ctx, &b)
			return err
		}

		if dtStartPrior && dtStopAfter { // case 6
			tmp := a[0].DtStop
			a[0].DtStop = *d1
			if err = UpdateRentableUseStatus(ctx, &a[0]); err != nil {
				return err
			}
			a[0].DtStop = tmp
			a[0].DtStart = *d2
			_, err = InsertRentableUseStatus(ctx, &a[0])
			if err != nil {
				return err
			}
			_, err = InsertRentableUseStatus(ctx, &b)
			return err
		}
	}
	return nil
}
