package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"sort"
)

// RRreportPaymentTypesTable generates a table object of all rlib.PaymentType for BID
func RRreportPaymentTypesTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRreportPaymentTypesTable"
	var (
		err error
	)

	// table init
	tbl := getRRTable()

	tbl.AddColumn("PMTID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	// set table title, sections
	err = TableReportHeaderBlock(ctx, &tbl, "Payment Types", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	m, err := rlib.GetPaymentTypesByBusiness(ctx, ri.Bid)
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
		v := m[i]
		tbl.AddRow()
		tbl.Puts(-1, 0, v.IDtoString())
		tbl.Puts(-1, 1, fmt.Sprintf("B%08d", v.BID))
		tbl.Puts(-1, 2, v.Name)
		tbl.Puts(-1, 3, v.Description)
	}
	tbl.TightenColumns()
	return tbl
}

// RRreportPaymentTypes generates a report of all rlib.GLAccount accounts
func RRreportPaymentTypes(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRreportPaymentTypesTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
