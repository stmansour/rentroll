package rrpt

import (
	"gotable"
	"rentroll/rlib"
)

// AccountDepositoryTable returns a table containing a report of all
// AccountDepositories in the supplied business
func AccountDepositoryTable(bid int64) gotable.Table {
	m := rlib.GetAllAccountDepositories(bid)
	var tbl gotable.Table
	tbl.Init()
	var b rlib.Business
	rlib.GetBusiness(bid, &b)
	rlib.GetDefaultLedgers(b.BID) // Gather its chart of accounts
	rlib.RRdb.BizTypes[b.BID].GLAccounts = rlib.GetGLAccountMap(b.BID)

	const (
		ADID           = 0
		Business       = iota
		LID            = iota
		AccountName    = iota
		DEPID          = iota
		DepositoryName = iota
	)

	tbl.AddColumn("ADID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BUD", 5, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("LID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("GL Account Name", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("DEPID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Depository Name", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		tbl.Puts(-1, ADID, m[i].IDtoString())
		tbl.Puts(-1, Business, b.Designation)
		tbl.Puts(-1, LID, rlib.IDtoString("L", m[i].LID))
		tbl.Puts(-1, AccountName, rlib.RRdb.BizTypes[b.BID].GLAccounts[m[i].LID].Name)
		tbl.Puts(-1, DEPID, rlib.IDtoString("DEP", m[i].DEPID))

		var dep rlib.Depository
		dep, err := rlib.GetDepository(m[i].DEPID)
		if err != nil {
			rlib.Ulog("Error getting AccountDepository %d: %s\n", m[i].DEPID, err.Error())
		} else {
			tbl.Puts(-1, DepositoryName, dep.Name)
		}
	}
	tbl.TightenColumns()
	return tbl
}

// AccountDepository returns a text report for AccountDepository records
// ri contains the BID needed by this report
func AccountDepository(ri *ReporterInfo) string {
	t := AccountDepositoryTable(ri.Bid)
	s, err := t.SprintTable()
	if err != nil {
		rlib.Ulog("AccountDepository: error = %s", err.Error())
	}
	return s
}
