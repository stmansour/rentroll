package rrpt

import (
	"fmt"
	"rentroll/rcsv"
	"rentroll/rlib"
)

// GSRTextReport generates a list of GSR values for all rentables on the specified date
func GSRTextReport(ri *rcsv.CSVReporterInfo) error {
	tbl, err := GSRReport(ri)
	fmt.Print(tbl)
	return err
}

// GSRReport generates a list of GSR values for all rentables on the specified date
func GSRReport(ri *rcsv.CSVReporterInfo) (rlib.Table, error) {
	funcname := "GSRTextReport"
	var tbl rlib.Table
	bu, err := rlib.GetBusinessUnitByDesignation(ri.Xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s\n", funcname, err.Error())
		return tbl, e
	}
	c, err := rlib.GetCompany(int64(bu.CoCode))
	if err != nil {
		e := fmt.Errorf("%s: error getting Company - %s\n", funcname, err.Error())
		return tbl, e
	}

	tbl.Init() //sets column spacing and date format to default

	s := fmt.Sprintf("%s\n", c.LegalName)
	s += fmt.Sprintf("%s\n", c.Address)
	if len(c.Address2) > 0 {
		s += fmt.Sprintf("%s\n", c.Address2)
	}
	s += fmt.Sprintf("%s, %s %s %s\n\n", c.City, c.State, c.PostalCode, c.Country)
	s += fmt.Sprintf("Gross Scheduled Rent for all rentables for one full cycle as of %s\n\n", ri.D1.Format(rlib.RRDATEFMT4))
	tbl.SetTitle(s)

	// fmt.Printf("%-9s  %-15s  %-15s  %-15s  %-8s  %-13s  %-13s\n", "Rentable", "Name", "Rentable Type", "Rentable Style", "GSR", "Rent Cycle", "Prorate Cycle")
	// fmt.Printf("%-9s  %-15s  %-15s  %-15s  %-8s  %-13s  %-13s\n", "---------", "---------------", "---------------", "---------------", "--------", "-------------", "-------------")

	tbl.AddColumn("Rentable", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)        // column for the Rentable name
	tbl.AddColumn("Name", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)           // Rentable name
	tbl.AddColumn("Rentable Type", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)  // Rentable Type
	tbl.AddColumn("Rentable Style", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT) // Rentable Style
	tbl.AddColumn("GSR", 8, rlib.CELLFLOAT, rlib.COLJUSTIFYLEFT)              // 4  GSR
	tbl.AddColumn("Rent Cycle", 13, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)     // 5  Rent Cycle
	tbl.AddColumn("Prorate Cycle", 13, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)  // 6  Proration Cycle

	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Xbiz.P.BID)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r rlib.Rentable
		var rc, pc int64
		rlib.ReadRentables(rows, &r)                         // get the next rentable from the database
		rtr := rlib.GetRentableTypeRefForDate(r.RID, &ri.D1) // what type is it on this date?
		rc = ri.Xbiz.RT[rtr.RTID].RentCycle
		pc = ri.Xbiz.RT[rtr.RTID].Proration
		if rtr.OverrideRentCycle != 0 {
			rc = rtr.OverrideRentCycle
		}
		if rtr.OverrideProrationCycle != 0 {
			pc = rtr.OverrideProrationCycle
		}
		dt1 := ri.D1.Add(rlib.CycleDuration(rc, ri.D1))                      // 1 full cycle
		amt, _, _, err := rlib.CalculateLoadedGSR(&r, &ri.D1, &dt1, ri.Xbiz) // calculate its GSR
		if err != nil {
			fmt.Printf("%s: Rentable %d, error calculating GSR: %s\n", funcname, r.RID, err.Error())
		}
		// fmt.Printf("%9s  %-15s  %-15s  %-15s  %8s  %-13s  %-13s\n",
		// 	r.IDtoString(), r.Name, ri.Xbiz.RT[rtr.RTID].Name, ri.Xbiz.RT[rtr.RTID].Style,
		// rlib.RRCommaf(amt), rlib.RentalPeriodToString(rc), rlib.RentalPeriodToString(pc))
		tbl.AddRow()
		tbl.Puts(-1, 0, r.IDtoString())
		tbl.Puts(-1, 1, r.Name)
		tbl.Puts(-1, 2, ri.Xbiz.RT[rtr.RTID].Name)
		tbl.Puts(-1, 3, ri.Xbiz.RT[rtr.RTID].Style)
		tbl.Putf(-1, 4, amt)
		tbl.Puts(-1, 5, rlib.RentalPeriodToString(rc))
		tbl.Puts(-1, 6, rlib.RentalPeriodToString(pc))
	}
	rlib.Errcheck(rows.Err())
	return tbl, err
}
