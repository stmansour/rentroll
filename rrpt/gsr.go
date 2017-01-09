package rrpt

import (
	"fmt"
	"rentroll/rlib"
)

// GSRTextReport generates a list of GSR values for all rentables on the specified date
func GSRTextReport(ri *ReporterInfo) error {
	tbl, err := GSRReport(ri)
	fmt.Print(tbl)
	return err
}

// ReportHeader returns a title block of text for a report.
// @params
//		  rn = Report Name
//	funcname = name of calling routine in case of error
//        ri = reporter info struct, please ensure RptHeaderD1 and RptHeaderD2 are set correctly
//
// @return
//		string = title string
//         err = any problem that occurred
func ReportHeader(rn, funcname string, ri *ReporterInfo) (string, error) {
	s := ""
	bu, err := rlib.GetBusinessUnitByDesignation(ri.Xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s\n", funcname, err.Error())
		return s, e
	}
	c, err := rlib.GetCompany(int64(bu.CoCode))
	if err != nil {
		e := fmt.Errorf("%s: error getting Company - %s\n", funcname, err.Error())
		return s, e
	}
	s += fmt.Sprintf("%s\n", c.LegalName)
	s += fmt.Sprintf("%s\n", c.Address)
	if len(c.Address2) > 0 {
		s += fmt.Sprintf("%s\n", c.Address2)
	}
	s += fmt.Sprintf("%s, %s %s %s\n\n", c.City, c.State, c.PostalCode, c.Country)
	s += rn + "\n"
	if ri.RptHeaderD1 {
		s += ri.D1.Format(rlib.RRDATEFMT4)
		if ri.RptHeaderD2 {
			s += " - " + ri.D2.Format(rlib.RRDATEFMT4)
		}
		s += "\n"
	}
	s += "\n"
	return s, nil
}

// GSRReport generates a list of GSR values for all rentables on the specified date
func GSRReport(ri *ReporterInfo) (rlib.Table, error) {
	funcname := "GSRTextReport"
	var tbl rlib.Table
	tbl.Init() //sets column spacing and date format to default
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = false
	s, err := ReportHeader("Gross Scheduled Rent", funcname, ri)
	if err != nil {
		return tbl, err
	}
	tbl.SetTitle(s)
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
