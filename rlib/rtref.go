package rlib

import "context"

// SetRentableTypeRef implements the proper insertion of a type reference
//     under all the circumstances considered.
//
// INPUTS
//     ctx - db context
//     rus - the new use status structure
//-----------------------------------------------------------------------------
func SetRentableTypeRef(ctx context.Context, rus *RentableTypeRef) error {

	funcname := "SetRentableTypeRef"

	Console("\nEntered %s.  range = %s, RTID = %d\n", funcname, ConsoleDRange(&rus.DtStart, &rus.DtStop), rus.RTID)

	var err error
	var b []RentableTypeRef
	d1 := rus.DtStart
	d2 := rus.DtStop
	a, err := GetRentableTypeRefsByRange(ctx, rus.RID, &d1, &d2)
	if err != nil {
		return err
	}

	Console("%s: Range = %s    found %d records\n", funcname, ConsoleDRange(&d1, &d2), len(a))

	//--------------------------------------------------------------------------
	// Remove any status records that are fully encompassed by rus.
	//--------------------------------------------------------------------------
	for i := 0; i < len(a); i++ {
		Console("i = %d, RTRID = %d\n", i, a[i].RTRID)
		if (d1.Before(a[i].DtStart) || d1.Equal(a[i].DtStart)) &&
			(d2.After(a[i].DtStop) || d2.Equal(a[i].DtStop)) {
			Console("%s: deleting RTRID = %d ------------------------------------\n", funcname, a[i].RTRID)
			if err = DeleteRentableTypeRef(ctx, a[i].RTRID); err != nil {
				return err
			}
		} else {
			Console("Appending RTRID=%d to a[]\n", a[i].RTRID)
			b = append(b, a[i])
		}
	}

	//-------------------------------------------------------------------
	// We're left with 0 or 1 or 2 items in b.  The overlap cases are
	// handled by this loop.  When it finishes, rus is is inserted.
	//-------------------------------------------------------------------
	if len(b) == 0 {
		_, err = InsertRentableTypeRef(ctx, rus)
		return err
	}

	//------------------------------------------------------------------------
	// CASE 1  -  after simplification, there is overlap on only one record
	//------------------------------------------------------------------------
	if len(b) == 1 {
		match := b[0].RTID == rus.RTID
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
			return UpdateRentableTypeRef(ctx, &b[0])
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
			if _, err = InsertRentableTypeRef(ctx, &n); err != nil {
				return err
			}
			b[0].DtStop = d1
			if err = UpdateRentableTypeRef(ctx, &b[0]); err != nil {
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
			if err = UpdateRentableTypeRef(ctx, &b[0]); err != nil {
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
			if err = UpdateRentableTypeRef(ctx, &b[0]); err != nil {
				return err
			}
		}
		Console("%s: Inserting %s RTID = %d\n", funcname, ConsoleDRange(&rus.DtStart, &rus.DtStop), rus.RTID)
		_, err = InsertRentableTypeRef(ctx, rus)
		return err
	}

	//------------------------------------------------------------------------
	// CASE 2  -  after simplification, there is overlap with two records
	//------------------------------------------------------------------------
	if len(b) == 2 {
		match0 := b[0].RTID == rus.RTID
		match1 := b[1].RTID == rus.RTID
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
			if err = UpdateRentableTypeRef(ctx, &b[0]); err != nil {
				return err
			}
			return DeleteRentableTypeRef(ctx, b[1].RTRID)
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
				if err = UpdateRentableTypeRef(ctx, &b[0]); err != nil {
					return err
				}
			}
			if d2.Before(b[1].DtStop) {
				b[1].DtStart = d2
			}
			if err = UpdateRentableTypeRef(ctx, &b[1]); err != nil {
				return err
			}
			_, err = InsertRentableTypeRef(ctx, rus)
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
			if err = UpdateRentableTypeRef(ctx, &b[0]); err != nil {
				return err
			}
			b[1].DtStart = d2
			return UpdateRentableTypeRef(ctx, &b[1])
		}

		if !match0 && match1 {
			// Case 2d
			// merge rus and b[1], update b[0]
			//  b[0:1]   @@@@@@@@@@************
			//  rus            *******
			//  Result   @@@@@@****************
			Console("%s: Case 2d b[0] matches\n", funcname)
			b[1].DtStart = d1
			if err = UpdateRentableTypeRef(ctx, &b[1]); err != nil {
				return err
			}
			b[0].DtStop = d1
			return UpdateRentableTypeRef(ctx, &b[0])
		}

		Console("%s: UNHANDLED CASE???\n", funcname)
	}

	return nil
}
