package rrpt

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// GSRTextReport generates a list of GSR values for all rentables on the specified date
func GSRTextReport(xbiz *rlib.XBusiness, dt *time.Time) error {
	funcname := "GSRTextReport"
	bu, err := rlib.GetBusinessUnitByDesignation(xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s\n", funcname, err.Error())
		return e
	}
	c, err := rlib.GetCompany(int64(bu.CoCode))
	if err != nil {
		e := fmt.Errorf("%s: error getting Company - %s\n", funcname, err.Error())
		return e
	}

	fmt.Printf("%s\n", c.LegalName)
	fmt.Printf("%s\n", c.Address)
	if len(c.Address2) > 0 {
		fmt.Printf("%s\n", c.Address2)
	}
	fmt.Printf("%s, %s %s %s\n\n", c.City, c.State, c.PostalCode, c.Country)

	fmt.Printf("Gross Scheduled Rent for all rentables for one full cycle as of %s\n\n", dt.Format(rlib.RRDATEFMT4))
	fmt.Printf("%-9s  %-15s  %-15s  %-15s  %-8s  %-13s  %-13s\n", "Rentable", "Name", "Rentable Type", "Rentable Style", "GSR", "Rent Cycle", "Prorate Cycle")
	fmt.Printf("%-9s  %-15s  %-15s  %-15s  %-8s  %-13s  %-13s\n", "---------", "---------------", "---------------", "---------------", "--------", "-------------", "-------------")
	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(xbiz.P.BID)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r rlib.Rentable
		var rc, pc int64
		rlib.ReadRentables(rows, &r)                     // get the next rentable from the database
		rtr := rlib.GetRentableTypeRefForDate(r.RID, dt) // what type is it on this date?
		rc = xbiz.RT[rtr.RTID].RentCycle
		pc = xbiz.RT[rtr.RTID].Proration
		if rtr.OverrideRentCycle != 0 {
			rc = rtr.OverrideRentCycle
		}
		if rtr.OverrideProrationCycle != 0 {
			pc = rtr.OverrideProrationCycle
		}
		dt1 := dt.Add(rlib.CycleDuration(rc, *dt))                    // 1 full cycle
		amt, _, _, err := rlib.CalculateLoadedGSR(&r, dt, &dt1, xbiz) // calculate its GSR
		if err != nil {
			fmt.Printf("%s: Rentable %d, error calculating GSR: %s\n", funcname, r.RID, err.Error())
		}
		fmt.Printf("%9s  %-15s  %-15s  %-15s  %8s  %-13s  %-13s\n",
			r.IDtoString(), r.Name, xbiz.RT[rtr.RTID].Name, xbiz.RT[rtr.RTID].Style, rlib.RRCommaf(amt), rlib.RentalPeriodToString(rc), rlib.RentalPeriodToString(pc))
	}
	rlib.Errcheck(rows.Err())
	return err
}
