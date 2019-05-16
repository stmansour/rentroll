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
//-----------------------------------------------------------------------------
func SetRentableLeaseStatusAbbr(ctx context.Context, bid, rid, us int64, d1, d2 *time.Time) error {
	var b = RentableLeaseStatus{
		RID:         rid,
		BID:         bid,
		DtStart:     *d1,
		DtStop:      *d2,
		Comment:     "",
		LeaseStatus: us,
	}

	return SetRentableLeaseStatus(ctx, &b)
}

// SetRentableLeaseStatus implements the proper insertion of a use status
//     under all the circumstances considered.
//
// INPUTS
//     ctx - db context
//     rls - the new use status structure
//-----------------------------------------------------------------------------
func SetRentableLeaseStatus(ctx context.Context, rls *RentableLeaseStatus) error {
	funcname := "SetRentableLeaseStatus"
	Console("\nEntered %s.  range = %s, LeaseStatus = %d\n", funcname, ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.LeaseStatus)

	var err error
	var b []RentableLeaseStatus
	d1 := rls.DtStart
	d2 := rls.DtStop
	a, err := GetRentableLeaseStatusByRange(ctx, rls.RID, &d1, &d2)
	if err != nil {
		return err
	}

	Console("%s: Range = %s    found %d records\n", funcname, ConsoleDRange(&d1, &d2), len(a))

	//--------------------------------------------------------------------------
	// Remove any status records that are fully encompassed by rls.
	//--------------------------------------------------------------------------
	for i := 0; i < len(a); i++ {
		Console("i = %d, RLID = %d\n", i, a[i].RLID)
		if (d1.Before(a[i].DtStart) || d1.Equal(a[i].DtStart)) &&
			(d2.After(a[i].DtStop) || d2.Equal(a[i].DtStop)) {
			Console("%s: deleting RLID = %d ------------------------------------\n", funcname, a[i].RLID)
			if err = DeleteRentableLeaseStatus(ctx, a[i].RLID); err != nil {
				return err
			}
		} else {
			Console("Appending RLID=%d to a[]\n", a[i].RLID)
			b = append(b, a[i])
		}
	}

	//-------------------------------------------------------------------
	// We're left with 0 or 1 or 2 items in b.  The overlap cases are
	// handled by this loop.  When it finishes, rls is is inserted.
	//-------------------------------------------------------------------
	if len(b) == 0 {
		_, err = InsertRentableLeaseStatus(ctx, rls)
		return err
	}

	//------------------------------------------------------------------------
	// CASE 1  -  after simplification, there is overlap on only one record
	//------------------------------------------------------------------------
	if len(b) == 1 {
		// match := b[0].LeaseStatus == rls.LeaseStatus
		match := RentableLeaseStatusCompare(&b[0], rls)
		before := b[0].DtStart.Before(d1)
		after := b[0].DtStop.After(d2)
		if match {
			//-----------------------------------------------
			// CASE 1a -  rls is contained by b[0] and statuses are equal
			//-----------------------------------------------
			//     b[0]: @@@@@@@@@@@@@@@@@@@@@
			//      rls:      @@@@@@@@@@@@
			//   Result: @@@@@@@@@@@@@@@@@@@@@
			//-----------------------------------------------
			Console("%s: Case 1a\n", funcname)
			if !before {
				b[0].DtStart = d1
			}
			if !after {
				b[0].DtStop = d2
			}
			return UpdateRentableLeaseStatus(ctx, &b[0])
		}

		if before && after {
			//-----------------------------------------------
			// CASE 1b -  rls contains b[0], match == false
			//-----------------------------------------------
			//     b[0]: @@@@@@@@@@@@@@@@@@@@@
			//      rls:      ############
			//   Result: @@@@@############@@@@
			//-----------------------------------------------
			Console("%s: Case 1b\n", funcname)
			n := b[0]
			n.DtStart = d2
			if _, err = InsertRentableLeaseStatus(ctx, &n); err != nil {
				return err
			}
			b[0].DtStop = d1
			if err = UpdateRentableLeaseStatus(ctx, &b[0]); err != nil {
				return err
			}
		}
		if !before {
			//-----------------------------------------------
			// CASE 1c -  rls prior to b[0], match == false
			//-----------------------------------------------
			//      rls: @@@@@@@@@@@@
			//     b[0]:       ##########
			//   Result: @@@@@@@@@@@@####
			//-----------------------------------------------
			Console("%s: Case 1c\n", funcname)
			b[0].DtStart = d2
			if err = UpdateRentableLeaseStatus(ctx, &b[0]); err != nil {
				return err
			}
		}
		if !after {
			//-----------------------------------------------
			// CASE 1d -  rls prior to b[0], match == false
			//-----------------------------------------------
			//      rls:     @@@@@@@@@@@@
			//     b[0]: ##########
			//   Result: ####@@@@@@@@@@@@
			//-----------------------------------------------
			Console("%s: Case 1d\n", funcname)
			b[0].DtStop = d1
			if err = UpdateRentableLeaseStatus(ctx, &b[0]); err != nil {
				return err
			}
		}
		Console("%s: Inserting %s LeaseStatus = %d\n", funcname, ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.LeaseStatus)
		_, err = InsertRentableLeaseStatus(ctx, rls)
		return err
	}

	//------------------------------------------------------------------------
	// CASE 2  -  after simplification, there is overlap with two records
	//------------------------------------------------------------------------
	if len(b) == 2 {
		// match0 := b[0].LeaseStatus == rls.LeaseStatus
		// match1 := b[1].LeaseStatus == rls.LeaseStatus
		match0 := RentableLeaseStatusCompare(&b[0], rls)
		match1 := RentableLeaseStatusCompare(&b[1], rls)
		before := b[0].DtStart.Before(d1)
		after := b[1].DtStop.After(d2)
		Console("%s: Case 2 and match0 = %t, match1 = %t\n", funcname, match0, match1)
		if match0 && match1 {
			// Case 2a
			// all are the same, merge them all into b[0], delete b[1]
			//  b[0:1]   ********* ************
			//  rls            *******
			//  Result   **********************
			Console("%s: Case 2a All match\n", funcname)
			if !before {
				b[0].DtStart = d1
			}
			b[0].DtStop = b[1].DtStop
			if !after {
				b[0].DtStop = d2
			}
			if err = UpdateRentableLeaseStatus(ctx, &b[0]); err != nil {
				return err
			}
			return DeleteRentableLeaseStatus(ctx, b[1].RLID)
		}

		if !match0 && !match1 {
			// Case 2b
			// neither match. Update both b[0] and b[1], add new rls
			//  b[0:1]   @@@@@@@@@@************
			//  rls            #######
			//  Result   @@@@@@#######*********
			Console("%s: Case 2b Both do not match\n", funcname)
			if d1.After(b[0].DtStart) {
				b[0].DtStop = d1
				if err = UpdateRentableLeaseStatus(ctx, &b[0]); err != nil {
					return err
				}
			}
			if d2.Before(b[1].DtStop) {
				b[1].DtStart = d2
			}
			if err = UpdateRentableLeaseStatus(ctx, &b[1]); err != nil {
				return err
			}
			_, err = InsertRentableLeaseStatus(ctx, rls)
			return err
		}

		if match0 && !match1 {
			// Case 2c
			// merge rls and b[0], update b[1]
			//  b[0:1]   @@@@@@@@@@************
			//  rls            @@@@@@@
			//  Result   @@@@@@@@@@@@@*********
			Console("%s: Case 2c b[0] matches\n", funcname)
			b[0].DtStop = d2
			if err = UpdateRentableLeaseStatus(ctx, &b[0]); err != nil {
				return err
			}
			b[1].DtStart = d2
			return UpdateRentableLeaseStatus(ctx, &b[1])
		}

		if !match0 && match1 {
			// Case 2d
			// merge rls and b[1], update b[0]
			//  b[0:1]   @@@@@@@@@@************
			//  rls            *******
			//  Result   @@@@@@****************
			Console("%s: Case 2d b[0] matches\n", funcname)
			b[1].DtStart = d1
			if err = UpdateRentableLeaseStatus(ctx, &b[1]); err != nil {
				return err
			}
			b[0].DtStop = d1
			return UpdateRentableLeaseStatus(ctx, &b[0])
		}

		Console("%s: UNHANDLED CASE???\n", funcname)
	}

	return nil

}

// RentableLeaseStatusCompare a deep comparison of the two lease status structs
//     on all fields EXCEPT dates and credit card info
//
// INPUTS
//     a - pointer to one of the lease status structs
//     b = a pointer to the other
//
// RETURNS
//     true if the fields match, false otherwise
//-----------------------------------------------------------------------------
func RentableLeaseStatusCompare(a, b *RentableLeaseStatus) bool {
	// if a.RID != b.RID {
	// 	Console("*** MISCOMPARE on RID ***\n")
	// }
	// if a.BID != b.BID {
	// 	Console("*** MISCOMPARE on BID ***\n")
	// }
	// if a.Comment != b.Comment {
	// 	Console("*** MISCOMPARE on Comment ***\n")
	// }
	// if a.FirstName != b.FirstName {
	// 	Console("*** MISCOMPARE on FirstName ***\n")
	// }
	// if a.LastName != b.LastName {
	// 	Console("*** MISCOMPARE on LastName ***\n")
	// }
	// if a.Email != b.Email {
	// 	Console("*** MISCOMPARE on Email ***\n")
	// }
	// if a.Phone != b.Phone {
	// 	Console("*** MISCOMPARE on Phone ***\n")
	// }
	// if a.Address != b.Address {
	// 	Console("*** MISCOMPARE on Address ***\n")
	// }
	// if a.Address2 != b.Address2 {
	// 	Console("*** MISCOMPARE on Address2 ***\n")
	// }
	// if a.City != b.City {
	// 	Console("*** MISCOMPARE on City ***\n")
	// }
	// if a.State != b.State {
	// 	Console("*** MISCOMPARE on State ***\n")
	// }
	// if a.PostalCode != b.PostalCode {
	// 	Console("*** MISCOMPARE on PostalCode ***\n")
	// }
	// if a.Country != b.Country {
	// 	Console("*** MISCOMPARE on Country ***\n")
	// }
	// if a.ConfirmationCode != b.ConfirmationCode {
	// 	Console("*** MISCOMPARE on ConfirmationCode ***\n")
	// }

	/*
		// if a.CCName != b.CCName {
		// 	Console("*** MISCOMPARE on CCName ***\n")
		// }
		// if a.CCType != b.CCType {
		// 	Console("*** MISCOMPARE on CCType ***\n")
		// }
		// if a.CCNumber != b.CCNumber {
		// 	Console("*** MISCOMPARE on CCNumber ***\n")
		// }
		// if a.CCExpMonth != b.CCExpMonth {
		// 	Console("*** MISCOMPARE on CCExpMonth ***\n")
		// }
	*/

	if a.RID != b.RID ||
		a.BID != b.BID ||
		a.Comment != b.Comment ||
		a.FirstName != b.FirstName ||
		a.LastName != b.LastName ||
		a.Email != b.Email ||
		a.Phone != b.Phone ||
		a.Address != b.Address ||
		a.Address2 != b.Address2 ||
		a.City != b.City ||
		a.State != b.State ||
		a.PostalCode != b.PostalCode ||
		a.Country != b.Country ||
		// a.CCName != b.CCName ||
		// a.CCType != b.CCType ||
		// a.CCNumber != b.CCNumber ||
		// a.CCExpMonth != b.CCExpMonth ||
		a.ConfirmationCode != b.ConfirmationCode ||
		a.LeaseStatus != b.LeaseStatus {
		return false
	}
	return true
}
