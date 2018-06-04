package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"sort"
)

// RRreportRentableTypesTable generates a table object of all Rentable Types defined in the database, for all businesses.
func RRreportRentableTypesTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportRentableTypesTable"

	// table init
	tbl := getRRTable()

	tbl.AddColumn("RTID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                    // 0
	tbl.AddColumn("Style", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                   // 1
	tbl.AddColumn("Name", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                    // 2
	tbl.AddColumn("Rent Cycle", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)               // 3
	tbl.AddColumn("Proration Cycle", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)          // 4
	tbl.AddColumn("GSRPC", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                    // 5
	tbl.AddColumn("Available?", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)              // 6
	tbl.AddColumn("Manage To Budget", 3, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)         // 7
	tbl.AddColumn("Is Child Type?", 3, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)           // 8
	tbl.AddColumn("Dt1 - Dt2 : Market Rate", 96, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT) // 9

	// set table title, sections
	err := TableReportHeaderBlock(ctx, &tbl, "Rentable Types", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	m, err := rlib.GetBusinessRentableTypes(ctx, ri.Bid)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		i := int64(k)
		p := m[i]

		// just before adding row to gotable, handle end date inclusion if applicable
		// for market rates
		rlib.HandleInterfaceEDI(&p, ri.Bid)

		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())
		tbl.Puts(-1, 1, p.Style)
		tbl.Puts(-1, 2, p.Name)
		tbl.Puts(-1, 3, rlib.RentalPeriodToString(p.RentCycle))
		tbl.Puts(-1, 4, rlib.RentalPeriodToString(p.Proration))
		tbl.Puts(-1, 5, rlib.RentalPeriodToString(p.GSRPC))

		// available or not
		availabilityStr := "Yes"
		if p.FLAGS&0x1 != 0 {
			availabilityStr = "No (Out of Service)"
		}
		tbl.Puts(-1, 6, availabilityStr) // availability

		// manage to budget
		manageToBudgetStr := "No"
		if p.FLAGS&0x4 != 0 {
			manageToBudgetStr = "Yes"
		}
		tbl.Puts(-1, 7, manageToBudgetStr)

		// child type?
		isChildTypeStr := "No"
		if p.FLAGS&0x2 != 0 {
			isChildTypeStr = "Yes"
		}
		tbl.Puts(-1, 8, isChildTypeStr) // is it child type

		// date ranges
		s := ""
		for i := 0; i < len(p.MR); i++ {
			s += fmt.Sprintf("%8s - %8s: $%8.2f", p.MR[i].DtStart.Format(rlib.RRDATEFMT4),
				p.MR[i].DtStop.Format(rlib.RRDATEFMT4), p.MR[i].MarketRate)
			if i+1 < len(p.MR) {
				s += ",  "
			}
		}
		tbl.Puts(-1, 9, s)
	}
	tbl.TightenColumns()
	return tbl
}

// RRreportRentableTypes generates a report of all Rentable Types defined in the database, for all businesses.
func RRreportRentableTypes(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportRentableTypesTable(ctx, ri)
	return ReportToString(&tbl, ri)
}

// RentableCountByRentableTypeReportTable returns an gotable.Table containing the count of Rentables for each RentableType
// in the specified time range
func RentableCountByRentableTypeReportTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RentableCountByRentableTypeReportTable"

	// init and prepare some values before table init
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true

	// table init
	tbl := getRRTable()

	tbl.AddColumn("No. Rentables", 9, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Rentable Type Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Style", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Custom Attributes", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err := TableReportHeaderBlock(ctx, &tbl, "Rentable Counts By Rentable Type", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	// RentableCountByRentableTypeReport returns a structure containing the count of Rentables for each RentableType
	// in the specified time range
	m, err := GetRentableCountByRentableType(ctx, ri.Xbiz, &ri.D1, &ri.D2)
	if err != nil {
		errMsg := fmt.Sprintf("%s: GetRentableCountByRentableType returned error: %s\n", funcname, err.Error())
		tbl.SetSection3(errMsg)
		return tbl
	}

	// need to sort these into a predictable order... they are messing up the tests as they
	// seem to come back in random orders on different runs...
	var keys []int
	for i := 0; i < len(m); i++ {
		keys = append(keys, i)
	}

	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if m[keys[i]].RT.Name > m[keys[j]].RT.Name {
				k := keys[i]
				keys[i] = keys[j]
				keys[j] = k
			}
		}
	}

	for i := 0; i < len(keys); i++ {
		j := int64(keys[i])
		// fmt.Printf("%13d  %-20.20s  %-6s", m[j].Count, m[j].RT.Name, m[j].RT.Style)
		tbl.AddRow()
		tbl.Puti(-1, 0, m[j].Count)
		tbl.Puts(-1, 1, m[j].RT.Name)
		tbl.Puts(-1, 2, m[j].RT.Style)
		s := ""
		for k, v := range m[j].RT.CA {
			if len(s) > 0 {
				s += ", "
			}
			s += fmt.Sprintf("%s: %s %s", k, v.Value, v.Units)
		}
		tbl.Puts(-1, 3, s)
	}
	tbl.TightenColumns()
	return tbl
}

// RentableCountByRentableTypeReport returns a string report containing the count of Rentables for each RentableType
// in the specified time range
func RentableCountByRentableTypeReport(ctx context.Context, ri *ReporterInfo) string {
	tbl := RentableCountByRentableTypeReportTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
