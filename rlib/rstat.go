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
	"Ready",                // 0
	"In Service",           // 1
	"Administrative",       // 2
	"Employee",             // 3
	"Owner Occupied",       // 4
	"Offline Major Repair", // 5
	"Offline housekeeping", // 6
	"Inactive",             // 7
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
//     res - if true, all the records beginning at d1 will be set to RESERVED
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
//     under all the circumstances considered.
//
// INPUTS
//     ctx - db context
//     rls - the new reservation structure
//     res - if true, all the records beginning at d1 will be set to RESERVED.
//           An example of this is when setting the time frame for a reservation.
//-----------------------------------------------------------------------------
func SetRentableLeaseStatus(ctx context.Context, rls *RentableLeaseStatus, res bool) error {
	// funcname := "SetRentableLeaseStatus"
	// Console("\nEntered %s.  range = %s, LeaseStatus = %d\n", funcname, ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.LeaseStatus)

	var err error
	var a []RentableLeaseStatus
	d1 := rls.DtStart
	d2 := rls.DtStop
	b, err := GetRentableLeaseStatusByRange(ctx, rls.RID, &d1, &d2)
	if err != nil {
		return err
	}

	// Console("%s: Range = %s    found %d records\n", funcname, ConsoleDRange(&d1, &d2), len(b))

	//--------------------------------------------------------------------------
	// Remove any status records that are fully encompassed by rls.
	//--------------------------------------------------------------------------
	for i := 0; i < len(b); i++ {
		// Console("i = %d, RLID = %d\n", i, b[i].RLID)
		if (d1.Before(b[i].DtStart) || d1.Equal(b[i].DtStart)) &&
			(d2.After(b[i].DtStop) || d2.Equal(b[i].DtStop)) {
			// Console("%s: deleting RLID = %d ------------------------------------\n", funcname, b[i].RLID)
			if err = DeleteRentableLeaseStatus(ctx, b[i].RLID); err != nil {
				return err
			}
		} else {
			// Console("Appending RLID=%d to a[]\n", b[i].RLID)
			b = append(a, b[i])
		}
	}

	//-------------------------------------------------------------------
	// We're left with 0 or 1 or 2 items in b.  The overlap cases are
	// handled by this loop.  When it finishes, rls is is inserted.
	//-------------------------------------------------------------------
	for i := 0; i < len(b); i++ {
		before := b[i].DtStart.Before(d1)
		after := b[i].DtStop.After(d2)
		if before && after {
			//-----------------------------------------------------------
			// Case 1:             b[i]  @@@@@@@@@@@@@@@@@@@@@@@@@@@
			//                     rls          ############
			//  Result:                  @@@@@@@############@@@@@@@@
			//-----------------------------------------------------------
			n := b[i]
			n.DtStop = d1
			if _, err = InsertRentableLeaseStatus(ctx, &n); err != nil {
				return err
			}
			b[i].DtStart = d2
			if err = UpdateRentableLeaseStatus(ctx, &b[i]); err != nil {
				return err
			}
		} else if before {
			//-----------------------------------------------------------
			//  Case 2:            b[i]  @@@@@@@@@@@@
			//                     rls          ############
			//  Result:                  @@@@@@@############
			//-----------------------------------------------------------
			b[i].DtStop = d1
			if err = UpdateRentableLeaseStatus(ctx, &b[i]); err != nil {
				return err
			}
		} else if after {
			//-----------------------------------------------------------
			//  Case 2:            b[i]         @@@@@@@@@@@@
			//                     rls   ############
			//  Result:                  ############@@@@@@@
			//-----------------------------------------------------------
			b[i].DtStart = d2
			if err = UpdateRentableLeaseStatus(ctx, &b[i]); err != nil {
				return err
			}
		}
	}

	// Console("%s: Inserting %s LeaseStatus = %d\n", funcname, ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.LeaseStatus)
	_, err = InsertRentableLeaseStatus(ctx, rls)
	return err
}

// SetRentableUseStatusAbbr changes the use status from d1 to d2 to the supplied
// status, us.
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
//     under all the circumstances considered.
//
// INPUTS
//     ctx - db context
//     rus - the new use status structure
//-----------------------------------------------------------------------------
func SetRentableUseStatus(ctx context.Context, rus *RentableUseStatus) error {
	funcname := "SetRentableUseStatus"
	Console("\nEntered %s.  range = %s, UseStatus = %d\n", funcname, ConsoleDRange(&rus.DtStart, &rus.DtStop), rus.UseStatus)

	var err error
	var b []RentableUseStatus
	d1 := rus.DtStart
	d2 := rus.DtStop
	a, err := GetRentableUseStatusByRange(ctx, rus.RID, &d1, &d2)
	if err != nil {
		return err
	}

	Console("%s: Range = %s    found %d records\n", funcname, ConsoleDRange(&d1, &d2), len(a))

	//--------------------------------------------------------------------------
	// Remove any status records that are fully encompassed by rus.
	//--------------------------------------------------------------------------
	for i := 0; i < len(a); i++ {
		Console("i = %d, RSID = %d\n", i, a[i].RSID)
		if (d1.Before(a[i].DtStart) || d1.Equal(a[i].DtStart)) &&
			(d2.After(a[i].DtStop) || d2.Equal(a[i].DtStop)) {
			Console("%s: deleting RSID = %d ------------------------------------\n", funcname, a[i].RSID)
			if err = DeleteRentableUseStatus(ctx, a[i].RSID); err != nil {
				return err
			}
		} else {
			Console("Appending RSID=%d to a[]\n", a[i].RSID)
			b = append(b, a[i])
		}
	}

	//-------------------------------------------------------------------
	// We're left with 0 or 1 or 2 items in b.  The overlap cases are
	// handled by this loop.  When it finishes, rus is is inserted.
	//-------------------------------------------------------------------
	if len(b) == 0 {
		_, err = InsertRentableUseStatus(ctx, rus)
		return err
	}

	//------------------------------------------------------------------------
	// CASE 1  -  after simplification, there is overlap on only one record
	//------------------------------------------------------------------------
	if len(b) == 1 {
		match := b[0].UseStatus == rus.UseStatus
		before := b[0].DtStart.Before(d1)
		after := b[0].DtStop.After(d2)
		if match {
			//-----------------------------------------------
			// CASE 1a -  rus is contained by b[0] and statuses are equal
			//-----------------------------------------------
			//     b[0]: @@@@@@@@@@@@@@@@@@@@@
			//      rus:      @@@@@@@@@@@@
			//   Result: @@@@@@@@@@@@@@@@@@@@@
			//-----------------------------------------------
			Console("%s: Case 1a\n", funcname)
			if !before {
				b[0].DtStart = d1
			}
			if !after {
				b[0].DtStop = d2
			}
			return UpdateRentableUseStatus(ctx, &b[0])
		}

		if before && after {
			//-----------------------------------------------
			// CASE 1b -  rus contains b[0], match == false
			//-----------------------------------------------
			//     b[0]: @@@@@@@@@@@@@@@@@@@@@
			//      rus:      ############
			//   Result: @@@@@############@@@@
			//-----------------------------------------------
			Console("%s: Case 1b\n", funcname)
			n := b[0]
			n.DtStart = d2
			if _, err = InsertRentableUseStatus(ctx, &n); err != nil {
				return err
			}
			b[0].DtStop = d1
			if err = UpdateRentableUseStatus(ctx, &b[0]); err != nil {
				return err
			}
		}
		if !before {
			//-----------------------------------------------
			// CASE 1c -  rus prior to b[0], match == false
			//-----------------------------------------------
			//      rus: @@@@@@@@@@@@
			//     b[0]:       ##########
			//   Result: @@@@@@@@@@@@####
			//-----------------------------------------------
			Console("%s: Case 1c\n", funcname)
			b[0].DtStart = d2
			if err = UpdateRentableUseStatus(ctx, &b[0]); err != nil {
				return err
			}
		}
		if !after {
			//-----------------------------------------------
			// CASE 1d -  rus prior to b[0], match == false
			//-----------------------------------------------
			//      rus:     @@@@@@@@@@@@
			//     b[0]: ##########
			//   Result: ####@@@@@@@@@@@@
			//-----------------------------------------------
			Console("%s: Case 1d\n", funcname)
			b[0].DtStop = d1
			if err = UpdateRentableUseStatus(ctx, &b[0]); err != nil {
				return err
			}
		}
		Console("%s: Inserting %s UseStatus = %d\n", funcname, ConsoleDRange(&rus.DtStart, &rus.DtStop), rus.UseStatus)
		_, err = InsertRentableUseStatus(ctx, rus)
		return err
	}

	//------------------------------------------------------------------------
	// CASE 2  -  after simplification, there is overlap with two records
	//------------------------------------------------------------------------
	if len(b) == 2 {
		match0 := b[0].UseStatus == rus.UseStatus
		match1 := b[1].UseStatus == rus.UseStatus
		before := b[0].DtStart.Before(d1)
		after := b[1].DtStop.After(d2)
		Console("%s: Case 2 and match0 = %t, match1 = %t\n", funcname, match0, match1)
		if match0 && match1 {
			// Case 2a
			// all are the same, merge them all into b[0], delete b[1]
			//  b[0:1]   ********* ************
			//  rus            *******
			//  Result   **********************
			Console("%s: Case 2a All match\n", funcname)
			if !before {
				b[0].DtStart = d1
			}
			b[0].DtStop = b[1].DtStop
			if !after {
				b[0].DtStop = d2
			}
			if err = UpdateRentableUseStatus(ctx, &b[0]); err != nil {
				return err
			}
			return DeleteRentableUseStatus(ctx, b[1].RSID)
		}

		if !match0 && !match1 {
			// Case 2b
			// neither match. Update both b[0] and b[1], add new rus
			//  b[0:1]   @@@@@@@@@@************
			//  rus            #######
			//  Result   @@@@@@#######*********
			Console("%s: Case 2b Both do not match\n", funcname)
			if d1.After(b[0].DtStart) {
				b[0].DtStop = d1
				if err = UpdateRentableUseStatus(ctx, &b[0]); err != nil {
					return err
				}
			}
			if d2.Before(b[1].DtStop) {
				b[1].DtStart = d2
			}
			if err = UpdateRentableUseStatus(ctx, &b[1]); err != nil {
				return err
			}
			_, err = InsertRentableUseStatus(ctx, rus)
			return err
		}

		if match0 && !match1 {
			// Case 2c
			// merge rus and b[0], update b[1]
			//  b[0:1]   @@@@@@@@@@************
			//  rus            @@@@@@@
			//  Result   @@@@@@@@@@@@@*********
			Console("%s: Case 2c b[0] matches\n", funcname)
			b[0].DtStop = d2
			if err = UpdateRentableUseStatus(ctx, &b[0]); err != nil {
				return err
			}
			b[1].DtStart = d2
			return UpdateRentableUseStatus(ctx, &b[1])
		}

		if !match0 && match1 {
			// Case 2d
			// merge rus and b[1], update b[0]
			//  b[0:1]   @@@@@@@@@@************
			//  rus            *******
			//  Result   @@@@@@****************
			Console("%s: Case 2d b[0] matches\n", funcname)
			b[1].DtStart = d1
			if err = UpdateRentableUseStatus(ctx, &b[1]); err != nil {
				return err
			}
			b[0].DtStop = d1
			return UpdateRentableUseStatus(ctx, &b[0])
		}

		Console("%s: UNHANDLED CASE???\n", funcname)
	}

	return nil

}

// SetRentableTypeRef implements the proper insertion of a type reference
//     under all the circumstances considered.
//
// INPUTS
//     ctx - db context
//     rtr - the new use status structure
//-----------------------------------------------------------------------------
func SetRentableTypeRef(ctx context.Context, rtr *RentableTypeRef) error {
	// funcname := "SetRentableTypeRef"
	// Console("\nEntered %s.  range = %s, RTID = %d\n", funcname, ConsoleDRange(&rtr.DtStart, &rtr.DtStop), rtr.RTID)

	var err error
	var a []RentableTypeRef
	d1 := rtr.DtStart
	d2 := rtr.DtStop
	b, err := GetRentableTypeRefsByRange(ctx, rtr.RID, &d1, &d2)
	if err != nil {
		return err
	}

	// Console("%s: Range = %s    found %d records\n", funcname, ConsoleDRange(&d1, &d2), len(b))

	//--------------------------------------------------------------------------
	// Remove any status records that are fully encompassed by rtr.
	//--------------------------------------------------------------------------
	for i := 0; i < len(b); i++ {
		// Console("i = %d, RTRID = %d\n", i, b[i].RTRID)
		if (d1.Before(b[i].DtStart) || d1.Equal(b[i].DtStart)) &&
			(d2.After(b[i].DtStop) || d2.Equal(b[i].DtStop)) {
			// Console("%s: deleting RTRID = %d ------------------------------------\n", funcname, b[i].RTRID)
			if err = DeleteRentableTypeRef(ctx, b[i].RTRID); err != nil {
				return err
			}
		} else {
			// Console("Appending RTRID=%d to a[]\n", b[i].RTRID)
			b = append(a, b[i])
		}
	}

	//-------------------------------------------------------------------
	// We're left with 0 or 1 or 2 items in b.  The overlap cases are
	// handled by this loop.  When it finishes, rtr is is inserted.
	//-------------------------------------------------------------------
	for i := 0; i < len(b); i++ {
		before := b[i].DtStart.Before(d1)
		after := b[i].DtStop.After(d2)
		if before && after {
			//-----------------------------------------------------------
			// Case 1:             b[i]  @@@@@@@@@@@@@@@@@@@@@@@@@@@
			//                     rtr          ############
			//  Result:                  @@@@@@@############@@@@@@@@
			//-----------------------------------------------------------
			n := b[i]
			n.DtStop = d1
			if _, err = InsertRentableTypeRef(ctx, &n); err != nil {
				return err
			}
			b[i].DtStart = d2
			if err = UpdateRentableTypeRef(ctx, &b[i]); err != nil {
				return err
			}
		} else if before {
			//-----------------------------------------------------------
			//  Case 2:            b[i]  @@@@@@@@@@@@
			//                     rtr          ############
			//  Result:                  @@@@@@@############
			//-----------------------------------------------------------
			b[i].DtStop = d1
			if err = UpdateRentableTypeRef(ctx, &b[i]); err != nil {
				return err
			}
		} else if after {
			//-----------------------------------------------------------
			//  Case 2:            b[i]         @@@@@@@@@@@@
			//                     rtr   ############
			//  Result:                  ############@@@@@@@
			//-----------------------------------------------------------
			b[i].DtStart = d2
			if err = UpdateRentableTypeRef(ctx, &b[i]); err != nil {
				return err
			}
		}
	}

	// Console("%s: Inserting %s RTID = %d\n", funcname, ConsoleDRange(&rtr.DtStart, &rtr.DtStop), rtr.RTID)
	_, err = InsertRentableTypeRef(ctx, rtr)
	return err
}
