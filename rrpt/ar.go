package rrpt

import (
	"context"
	"gotable"
	"rentroll/rlib"
)

// RRARTable generates a gotable table object
// for report of all rlib.AR rules records related with business
func RRARTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRARTable"
	var (
		err error
		bid = ri.Bid
		// d1  = ri.D1
		// d2  = ri.D2
	)

	// rlib.Console("Entered %s\n", funcname)

	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true
	ri.BlankLineAfterRptName = true

	// initialize table for assessments report
	tbl := getRRTable()

	const (
		BUD          = 0
		Name         = iota
		ARType       = iota
		DebitLID     = iota
		CreditLID    = iota
		Allocated    = iota
		ShowOnRA     = iota
		RAIDRequired = iota
		SubARSpec    = iota
		Description  = iota
	)

	tbl.AddColumn("BUD", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ARType", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Debit", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Credit", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Allocated", 5, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("ShowOnRA", 5, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("RAIDRequired", 5, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("SubARSpec", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// prepare table's title, section1, section2, section3 if there are any error
	err = TableReportHeaderBlock(ctx, &tbl, "Account Rules", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// rlib.Console("%s: bid = %d\n", funcname, bid)
	rlib.Console("InitBusinessFields\n")
	var b rlib.XBusiness
	err = rlib.InitBizInternals(bid, &b)
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}
	rlib.Console("After InitBizInternals, len(rlib.RRdb.BizTypes[%d].AR) = %d\n", bid, len(rlib.RRdb.BizTypes[bid].AR))

	// rlib.Console("GetAllARs\n")
	m, err := rlib.GetAllARs(ctx, bid)
	if err != nil {
		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	// rlib.Console("Number of records: %d\n", len(m))
	// fit records in table row one by one
	for i := 0; i < len(m); i++ {
		//-------------------
		// init some indeces
		//-------------------
		art := m[i].ARType
		if art < 0 {
			art = 0
		}
		if art > rlib.ARSUBASSESSMENT {
			art = rlib.ARASSESSMENT
		}
		alloc := int64(0)
		if m[i].FLAGS&0x1 > 0 {
			alloc = 1
		}
		show := int64(0)
		if m[i].FLAGS&0x2 > 0 {
			show = 1
		}
		raidrqd := int64(0)
		if m[i].FLAGS&0x2 > 0 {
			raidrqd = 1
		}
		subar := ""
		for j := 0; j < len(m[i].SubARs); j++ {
			subarid := m[i].SubARs[j].ARID
			if subarid > 0 {
				if len(subar) > 0 {
					subar += ", "
				}
				subar += rlib.RRdb.BizTypes[bid].AR[subarid].Name
			}
		}

		//-----------------------
		// Now emit the columns
		//-----------------------
		tbl.AddRow()
		tbl.Puts(-1, BUD, b.P.Designation)
		tbl.Puts(-1, Name, m[i].Name)
		tbl.Puts(-1, ARType, rlib.ARTypesList[art])
		tbl.Puts(-1, DebitLID, rlib.RRdb.BizTypes[bid].GLAccounts[m[i].DebitLID].Name)
		tbl.Puts(-1, CreditLID, rlib.RRdb.BizTypes[bid].GLAccounts[m[i].CreditLID].Name)
		tbl.Puts(-1, Allocated, rlib.YesNoToString(alloc))
		tbl.Puts(-1, ShowOnRA, rlib.YesNoToString(show))
		tbl.Puts(-1, RAIDRequired, rlib.YesNoToString(raidrqd))
		tbl.Puts(-1, SubARSpec, subar)
		tbl.Puts(-1, Description, m[i].Description)
	}

	tbl.TightenColumns()
	return tbl
}

// RRreportAR generates a text report of all rlib.AR records
func RRreportAR(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRARTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
