package rlib

import (
	"context"
	"time"
)

// SetRentableMarketRateAbbr changes the market rate from d1 to d2 to the supplied
// status, x. It adds and modifies existing records as needed.
//
// INPUTS
//     ctx - db context
//     bid - which business
//     rtid - which rentable type
//     x   - new market rate
//     d1  - start time for status x
//     d2  - stop time for status x
//-----------------------------------------------------------------------------
func SetRentableMarketRateAbbr(ctx context.Context, bid, rtid int64, x float64, d1, d2 *time.Time) error {
	var b = RentableMarketRate{
		RTID:       rtid,
		BID:        bid,
		DtStart:    *d1,
		DtStop:     *d2,
		MarketRate: x,
	}

	return SetRentableMarketRate(ctx, &b)
}

// SetRentableMarketRate implements the proper insertion of a market rate
//     under all the circumstances considered.
//
// INPUTS
//     ctx - db context
//     rls - the new market rate structure
//-----------------------------------------------------------------------------
func SetRentableMarketRate(ctx context.Context, rls *RentableMarketRate) error {
	// funcname := "SetRentableMarketRate"
	// Console("\nEntered %s.  range = %s, MarketRate = %d\n", funcname, ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.MarketRate)

	var err error
	var b []RentableMarketRate
	d1 := rls.DtStart
	d2 := rls.DtStop
	a, err := GetRentableMarketRateByRange(ctx, rls.RTID, &d1, &d2)
	if err != nil {
		return err
	}

	// Console("%s: Range = %s    found %d records\n", funcname, ConsoleDRange(&d1, &d2), len(a))

	//--------------------------------------------------------------------------
	// Remove any status records that are fully encompassed by rls.
	//--------------------------------------------------------------------------
	for i := 0; i < len(a); i++ {
		// Console("i = %d, RMRID = %d\n", i, a[i].RMRID)
		if (d1.Before(a[i].DtStart) || d1.Equal(a[i].DtStart)) &&
			(d2.After(a[i].DtStop) || d2.Equal(a[i].DtStop)) {
			// Console("%s: deleting RMRID = %d ------------------------------------\n", funcname, a[i].RMRID)
			if err = DeleteRentableMarketRate(ctx, a[i].RMRID); err != nil {
				return err
			}
		} else {
			// Console("Appending RMRID=%d to a[]\n", a[i].RMRID)
			b = append(b, a[i])
		}
	}

	//-------------------------------------------------------------------
	// We're left with 0 or 1 or 2 items in b.  The overlap cases are
	// handled by this loop.  When it finishes, rls is is inserted.
	//-------------------------------------------------------------------
	if len(b) == 0 {
		_, err = InsertRentableMarketRate(ctx, rls)
		return err
	}

	//------------------------------------------------------------------------
	// CASE 1  -  after simplification, there is overlap on only one record
	//------------------------------------------------------------------------
	if len(b) == 1 {
		match := b[0].MarketRate == rls.MarketRate
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
			// Console("%s: Case 1a\n", funcname)
			if !before {
				b[0].DtStart = d1
			}
			if !after {
				b[0].DtStop = d2
			}
			return UpdateRentableMarketRate(ctx, &b[0])
		}

		if before && after {
			//-----------------------------------------------
			// CASE 1b -  rls contains b[0], match == false
			//-----------------------------------------------
			//     b[0]: @@@@@@@@@@@@@@@@@@@@@
			//      rls:      ############
			//   Result: @@@@@############@@@@
			//-----------------------------------------------
			// Console("%s: Case 1b\n", funcname)
			n := b[0]
			n.DtStart = d2
			if _, err = InsertRentableMarketRate(ctx, &n); err != nil {
				return err
			}
			b[0].DtStop = d1
			if err = UpdateRentableMarketRate(ctx, &b[0]); err != nil {
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
			// Console("%s: Case 1c\n", funcname)
			b[0].DtStart = d2
			if err = UpdateRentableMarketRate(ctx, &b[0]); err != nil {
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
			// Console("%s: Case 1d\n", funcname)
			b[0].DtStop = d1
			if err = UpdateRentableMarketRate(ctx, &b[0]); err != nil {
				return err
			}
		}
		// Console("%s: Inserting %s MarketRate = %d\n", funcname, ConsoleDRange(&rls.DtStart, &rls.DtStop), rls.MarketRate)
		_, err = InsertRentableMarketRate(ctx, rls)
		return err
	}

	//------------------------------------------------------------------------
	// CASE 2  -  after simplification, there is overlap with two records
	//------------------------------------------------------------------------
	if len(b) == 2 {
		match0 := b[0].MarketRate == rls.MarketRate
		match1 := b[1].MarketRate == rls.MarketRate
		before := b[0].DtStart.Before(d1)
		after := b[1].DtStop.After(d2)
		// Console("%s: Case 2 and match0 = %t, match1 = %t\n", funcname, match0, match1)
		if match0 && match1 {
			// Case 2a
			// all are the same, merge them all into b[0], delete b[1]
			//  b[0:1]   ********* ************
			//  rls            *******
			//  Result   **********************
			// Console("%s: Case 2a All match\n", funcname)
			if !before {
				b[0].DtStart = d1
			}
			b[0].DtStop = b[1].DtStop
			if !after {
				b[0].DtStop = d2
			}
			if err = UpdateRentableMarketRate(ctx, &b[0]); err != nil {
				return err
			}
			return DeleteRentableMarketRate(ctx, b[1].RMRID)
		}

		if !match0 && !match1 {
			// Case 2b
			// neither match. Update both b[0] and b[1], add new rls
			//  b[0:1]   @@@@@@@@@@************
			//  rls            #######
			//  Result   @@@@@@#######*********
			// Console("%s: Case 2b Both do not match\n", funcname)
			if d1.After(b[0].DtStart) {
				b[0].DtStop = d1
				if err = UpdateRentableMarketRate(ctx, &b[0]); err != nil {
					return err
				}
			}
			if d2.Before(b[1].DtStop) {
				b[1].DtStart = d2
			}
			if err = UpdateRentableMarketRate(ctx, &b[1]); err != nil {
				return err
			}
			_, err = InsertRentableMarketRate(ctx, rls)
			return err
		}

		if match0 && !match1 {
			// Case 2c
			// merge rls and b[0], update b[1]
			//  b[0:1]   @@@@@@@@@@************
			//  rls            @@@@@@@
			//  Result   @@@@@@@@@@@@@*********
			// Console("%s: Case 2c b[0] matches\n", funcname)
			b[0].DtStop = d2
			if err = UpdateRentableMarketRate(ctx, &b[0]); err != nil {
				return err
			}
			b[1].DtStart = d2
			return UpdateRentableMarketRate(ctx, &b[1])
		}

		if !match0 && match1 {
			// Case 2d
			// merge rls and b[1], update b[0]
			//  b[0:1]   @@@@@@@@@@************
			//  rls            *******
			//  Result   @@@@@@****************
			// Console("%s: Case 2d b[0] matches\n", funcname)
			b[1].DtStart = d1
			if err = UpdateRentableMarketRate(ctx, &b[1]); err != nil {
				return err
			}
			b[0].DtStop = d1
			return UpdateRentableMarketRate(ctx, &b[0])
		}
		// Console("%s: UNHANDLED CASE???\n", funcname)
	}
	return nil
}
