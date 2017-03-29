package rrpt

import (
	"gotable"
	"rentroll/rlib"
	"strings"
)

func getCategory(s string) (string, string) {
	cat := ""
	val := ""
	loc := strings.Index(s, "^")
	if loc > 0 {
		cat = strings.TrimSpace(s[:loc])
		if len(s) > loc+1 {
			val = strings.TrimSpace(s[loc+1:])
		}
	} else {
		val = s
	}
	return cat, val
}

// RRreportStringListsTable generates a table object of all StringLists for the supplied business (ri.Bid)
func RRreportStringListsTable(ri *ReporterInfo) gotable.Table {
	funcname := "RRreportStringListsTable"

	// init and prepare some values before table init
	var (
		cat, val string
	)

	// table init
	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)

	tbl.AddColumn("SLSID", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Category", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Value", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	err := TableReportHeaderBlock(&tbl, "String Lists", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)

		// set errors in section3 and return
		tbl.SetSection3(err.Error())
		return tbl
	}

	m := rlib.GetAllStringLists(ri.Bid)

	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		tbl.Puts(-1, 0, m[i].Name)
		for j := 0; j < len(m[i].S); j++ {
			cat, val = getCategory(m[i].S[j].Value)
			tbl.AddRow()
			tbl.Puts(-1, 0, rlib.IDtoString("SLS", m[i].S[j].SLSID))
			tbl.Puts(-1, 1, cat)
			tbl.Puts(-1, 2, val)
		}
	}
	return tbl
}

// RRreportStringLists generates a report of all StringLists for the supplied business (ri.Bid)
func RRreportStringLists(ri *ReporterInfo) string {
	tbl := RRreportStringListsTable(ri)
	return ReportToString(&tbl, ri)
}
