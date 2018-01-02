package rrpt

import (
	"context"
	"gotable"
	"rentroll/rlib"
)

// RRreportDeposits generates a report of all rlib.Deposit
func RRreportDeposits(ctx context.Context, ri *ReporterInfo) string {
	const funcname = "RRreportDeposits"
	var (
		err error
	)

	var t gotable.Table
	t.Init()

	// TODO(Steve): this routine moved from rcsv/report.go, should we replace Rcsv.DtStart, Rcsv.DtStop with "ri" dates
	m, err := rlib.GetAllDepositsInRange(ctx, ri.Bid, &ri.D1, &ri.D2)
	if err != nil {
		t.SetSection3(err.Error())
		return t.String()
	}

	err = TableReportHeaderBlock(ctx, &t, "Deposit", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		t.SetSection3(err.Error())
		return t.String()
	}

	t.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("DEPID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Receipts", 60, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	for i := 0; i < len(m); i++ {
		s := ""
		for j := 0; j < len(m[i].DP); j++ {
			s += rlib.IDtoString("RCPT", m[i].DP[j].RCPTID)
			if j+1 < len(m[i].DP) {
				s += ", "
			}
		}
		t.AddRow()
		t.Putd(-1, 0, m[i].Dt)
		t.Puts(-1, 1, m[i].IDtoString())
		t.Puts(-1, 2, rlib.IDtoString("B", m[i].BID))
		t.Putf(-1, 3, m[i].Amount)
		t.Puts(-1, 4, s)
	}

	t.TightenColumns()
	s, err := t.SprintTable()
	if nil != err {
		rlib.Ulog("RRreportDeposits: error = %s", err)
	}
	return s
}
