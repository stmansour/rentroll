package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"sort"
	"strings"
	"time"
)

// RRreportBusiness generates a report of all Businesses defined in the database.
func RRreportBusiness(ri *CSVReporterInfo) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllBusinesses.Query()
	rlib.Errcheck(err)
	defer rows.Close()

	var tbl rlib.Table
	tbl.Init()
	tbl.AddColumn("BID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("BUD", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Name", 20, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Default Rent Cycle", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Default Proration Cycle", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Default GSRPC Cycle", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for rows.Next() {
		var p rlib.Business
		rlib.ReadBusinesses(rows, &p)
		tbl.AddRow()
		tbl.Puts(-1, 0, p.IDtoString())
		tbl.Puts(-1, 1, p.Designation)
		tbl.Puts(-1, 2, p.Name)
		tbl.Puts(-1, 3, rlib.RentalPeriodToString(p.DefaultRentCycle))
		tbl.Puts(-1, 4, rlib.RentalPeriodToString(p.DefaultProrationCycle))
		tbl.Puts(-1, 5, rlib.RentalPeriodToString(p.DefaultGSRPC))
	}
	rlib.Errcheck(rows.Err())
	return tbl.SprintTable(ri.OutputFormat)
}

// ReportCOA returns a string representation of the chart of accts
func ReportCOA(p rlib.GLAccount, t *rlib.Table) {
	Pldgr := ""
	lm := rlib.GetLatestLedgerMarkerByLID(p.BID, p.LID)
	if lm.LMID == 0 {
		fmt.Printf("ReportChartOfAcctsToText: error getting latest LedgerMarker for L%08d\n", p.LID)
		return
	}

	s := ""
	if rlib.GLCASH <= p.Type && p.Type <= rlib.GLLAST {
		s = fmt.Sprintf("%d", p.Type)
	}

	sp := ""
	switch p.RAAssociated {
	case 0:
		sp = "unknown"
	case 1:
		sp = "Unassociated"
	case 2:
		sp = "Associated"
	default:
		sp = fmt.Sprintf("??? invalid: %d", p.RAAssociated)
	}
	if p.PLID > 0 {
		Pldgr = rlib.RRdb.BizTypes[p.BID].GLAccounts[p.PLID].Name
	}
	const (
		BID  = 0
		LID  = iota
		PLID = iota
		LMID = iota
		Type = iota
		GLNo = iota
		Name = iota
		PGL  = iota
		QBAT = iota
		Bal  = iota
		RAA  = iota
		RAR  = iota
		Desc = iota
	)

	t.AddRow()
	t.Puts(-1, BID, fmt.Sprintf("B%08d", p.BID))
	t.Puts(-1, LID, p.IDtoString())
	if p.PLID > 0 {
		t.Puts(-1, PLID, fmt.Sprintf("L%08d", p.PLID))
	}
	t.Puts(-1, LMID, lm.IDtoString())
	t.Puts(-1, Type, s)
	t.Puts(-1, GLNo, p.GLNumber)
	t.Puts(-1, Name, p.Name)
	t.Puts(-1, PGL, Pldgr)
	t.Puts(-1, QBAT, p.AcctType)
	t.Putf(-1, Bal, lm.Balance)
	t.Puts(-1, RAA, sp)
	t.Puti(-1, RAR, p.RARequired)
	t.Puts(-1, Desc, p.Description)
}

// RRreportChartOfAccounts generates a report of all rlib.GLAccount accounts
func RRreportChartOfAccounts(ri *CSVReporterInfo) string {
	rlib.InitBusinessFields(ri.Bid)
	rlib.RRdb.BizTypes[ri.Bid].GLAccounts = rlib.GetGLAccountMap(ri.Bid)

	var a []int64                                          // Sort the map so test output will be the same every time. Sort by GLNumber.
	for k := range rlib.RRdb.BizTypes[ri.Bid].GLAccounts { // First make an array of all the LIDs
		a = append(a, k)
	}
	// now sort based on GLNumber, then by LID...
	m := rlib.RRdb.BizTypes[ri.Bid].GLAccounts // for notational convenience
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			isGreater := m[a[i]].GLNumber > m[a[j]].GLNumber
			isEqual := m[a[i]].GLNumber == m[a[j]].GLNumber
			if isGreater || (isEqual && m[a[i]].LID > m[a[j]].LID) {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	var t rlib.Table
	t.Init()
	t.AddColumn("BID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("LID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("PLID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("LMID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Type", 8, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("GLNumber", 8, rlib.CELLSTRING, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Name", 40, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Parent", 35, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Quick Books Account Type", 20, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Balance", 12, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Rental Agreement Associated", 12, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Rental Agreement Required", 5, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Description", 25, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for i := 0; i < len(a); i++ {
		ReportCOA(m[a[i]], &t)
	}
	return t.SprintTable(ri.OutputFormat)
}

// RRreportRentableTypes generates a report of all Rentable Types defined in the database, for all businesses.
func RRreportRentableTypes(ri *CSVReporterInfo) string {
	m := rlib.GetBusinessRentableTypes(ri.Bid)
	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	var t rlib.Table
	t.Init()
	t.AddColumn("RTID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)                    // 0
	t.AddColumn("Style", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)                   // 1
	t.AddColumn("Name", 25, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)                    // 2
	t.AddColumn("Rent Cycle", 8, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)               // 3
	t.AddColumn("Proration Cycle", 8, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)          // 4
	t.AddColumn("GSRPC", 8, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)                    // 5
	t.AddColumn("Manage To Budget", 3, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)         // 6
	t.AddColumn("Dt1 - Dt2 : Market Rate", 96, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT) // 7

	for _, k := range keys {
		i := int64(k)
		p := m[i]
		t.AddRow()
		t.Puts(-1, 0, p.IDtoString())
		t.Puts(-1, 1, p.Style)
		t.Puts(-1, 2, p.Name)
		t.Puts(-1, 3, rlib.RentalPeriodToString(p.RentCycle))
		t.Puts(-1, 4, rlib.RentalPeriodToString(p.Proration))
		t.Puts(-1, 5, rlib.RentalPeriodToString(p.GSRPC))
		t.Puts(-1, 6, rlib.YesNoToString(p.ManageToBudget))
		s := ""
		for i := 0; i < len(p.MR); i++ {
			s += fmt.Sprintf("|  %8s - %8s $%8.2f", p.MR[i].DtStart.Format(rlib.RRDATEFMT2),
				p.MR[i].DtStop.Format(rlib.RRDATEFMT2), p.MR[i].MarketRate)
		}
		t.Puts(-1, 7, s)
	}
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportPeople generates a report of all Businesses defined in the database.
func RRreportPeople(ri *CSVReporterInfo) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllTransactantsForBID.Query(ri.Bid)
	rlib.Errcheck(err)
	defer rows.Close()

	var t rlib.Table
	t.Init()
	t.AddColumn("TCID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("First Name", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Middle Name", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Last Name", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Company", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Is Company", 3, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Cell Phone", 17, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Primary Email", 25, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for rows.Next() {
		var p rlib.XPerson
		rlib.ReadTransactants(rows, &p.Trn)
		rlib.GetXPerson(p.Trn.TCID, &p)
		t.AddRow()
		t.Puts(-1, 0, p.IDtoString())
		t.Puts(-1, 1, p.Trn.FirstName)
		t.Puts(-1, 2, p.Trn.MiddleName)
		t.Puts(-1, 3, p.Trn.LastName)
		t.Puts(-1, 4, p.Trn.CompanyName)
		t.Puts(-1, 5, rlib.YesNoToString(int64(p.Trn.IsCompany)))
		t.Puts(-1, 6, p.Trn.CellPhone)
		t.Puts(-1, 7, p.Trn.PrimaryEmail)
	}
	rlib.Errcheck(rows.Err())
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportRentables generates a report of all Businesses defined in the database.
func RRreportRentables(ri *CSVReporterInfo) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Bid)
	rlib.Errcheck(err)
	defer rows.Close()

	var t rlib.Table
	t.Init()
	t.AddColumn("RID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Name", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Assignment Time", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for rows.Next() {
		var p rlib.Rentable
		var s string
		rlib.ReadRentables(rows, &p)
		switch p.AssignmentTime {
		case 0:
			s = "unknown"
		case 1:
			s = "pre-assign"
		case 2:
			s = "no pre-assign"
		}
		t.AddRow()
		t.Puts(-1, 0, p.IDtoString())
		t.Puts(-1, 1, p.Name)
		t.Puts(-1, 2, s)
	}
	rlib.Errcheck(rows.Err())
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportCustomAttributes generates a report of all rlib.GLAccount accounts
func RRreportCustomAttributes(ri *CSVReporterInfo) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllCustomAttributes.Query()
	rlib.Errcheck(err)
	defer rows.Close()

	var t rlib.Table
	t.Init()
	t.AddColumn("CID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Value Type", 6, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Name", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Value", 15, rlib.CELLSTRING, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Units", 15, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for rows.Next() {
		var a rlib.CustomAttribute
		rlib.ReadCustomAttributes(rows, &a)
		t.AddRow()
		t.Puts(-1, 0, a.IDtoString())
		t.Puts(-1, 1, a.TypeToString())
		t.Puts(-1, 2, a.Name)
		t.Puts(-1, 3, a.Value)
		t.Puts(-1, 4, a.Units)
	}
	rlib.Errcheck(rows.Err())
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportCustomAttributeRefs generates a report of all rlib.GLAccount accounts
func RRreportCustomAttributeRefs(ri *CSVReporterInfo) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllCustomAttributeRefs.Query()
	rlib.Errcheck(err)
	defer rows.Close()
	var t rlib.Table
	t.Init()
	t.AddColumn("Element Type", 4, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("ID", 4, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("CID", 4, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	for rows.Next() {
		var a rlib.CustomAttributeRef
		rlib.ReadCustomAttributeRefs(rows, &a)
		t.AddRow()
		t.Puti(-1, 0, a.ElementType)
		t.Puti(-1, 1, a.ID)
		t.Puti(-1, 2, a.CID)
	}
	rlib.Errcheck(rows.Err())
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportRentalAgreementTemplates generates a report of all Businesses defined in the database.
func RRreportRentalAgreementTemplates(ri *CSVReporterInfo) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementTemplates.Query()
	rlib.Errcheck(err)
	defer rows.Close()
	var t rlib.Table
	t.Init()
	t.AddColumn("BID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("RA Template ID", 11, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("RA Template Name", 25, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	for rows.Next() {
		var p rlib.RentalAgreementTemplate
		rlib.ReadRentalAgreementTemplates(rows, &p)
		t.AddRow()
		t.Puts(-1, 0, fmt.Sprintf("B%08d", p.BID))
		t.Puts(-1, 1, p.IDtoString())
		t.Puts(-1, 2, p.RATemplateName)
	}
	rlib.Errcheck(rows.Err())
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportRentalAgreements generates a report of all Businesses defined in the database.
func RRreportRentalAgreements(ri *CSVReporterInfo) string {
	rs := ""
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreements.Query(ri.Bid)
	rlib.Errcheck(err)
	defer rows.Close()
	var t rlib.Table
	t.Init()
	t.AddColumn("RAID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Payor", 60, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("User", 60, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Agreement Start", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Agreement Stop", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Possession Start", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Possession Stop", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Rent Start", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Rent Stop", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Rent Cycle Epoch", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Renewal", 2, rlib.CELLINT, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Unspecified Adults", 6, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Unspecified Children", 6, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Special Provisions", 40, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Notes", 30, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	var raid int64
	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var p rlib.RentalAgreement
		rlib.Errcheck(rows.Scan(&raid))
		p, err = rlib.GetXRentalAgreement(raid, &d1, &d2)
		if err != nil {
			rs += fmt.Sprintf("RRreportRentalAgreements: rlib.GetXRentalAgreement returned err = %v\n", err)
			continue
		}
		note := ""
		if p.NLID > 0 {
			nl := rlib.GetNoteList(p.NLID)
			if len(nl.N) > 0 {
				note = nl.N[0].Comment
			}
		}
		t.AddRow()
		t.Puts(-1, 0, p.IDtoString())
		t.Puts(-1, 1, strings.Join(p.GetPayorNameList(&p.AgreementStart, &p.AgreementStop), ", "))
		t.Puts(-1, 2, strings.Join(p.GetUserNameList(&p.AgreementStart, &p.AgreementStop), ", "))
		t.Putd(-1, 3, p.AgreementStart)
		t.Putd(-1, 4, p.AgreementStop)
		t.Putd(-1, 5, p.PossessionStart)
		t.Putd(-1, 6, p.PossessionStop)
		t.Putd(-1, 7, p.RentStart)
		t.Putd(-1, 8, p.RentStop)
		t.Putd(-1, 9, p.RentCycleEpoch)
		t.Puti(-1, 10, p.Renewal)
		t.Puti(-1, 11, p.UnspecifiedAdults)
		t.Puti(-1, 12, p.UnspecifiedChildren)
		t.Puts(-1, 13, p.SpecialProvisions)
		t.Puts(-1, 14, note)
	}
	rlib.Errcheck(rows.Err())
	t.TightenColumns()
	return rs + t.SprintTable(ri.OutputFormat)
}

// RRreportPaymentTypes generates a report of all rlib.GLAccount accounts
func RRreportPaymentTypes(ri *CSVReporterInfo) string {
	m := rlib.GetPaymentTypesByBusiness(ri.Bid)
	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	var t rlib.Table
	t.Init()
	t.AddColumn("PMTID", 11, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("BID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Name", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Description", 30, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for _, k := range keys {
		i := int64(k)
		v := m[i]
		t.AddRow()
		t.Puts(-1, 0, v.IDtoString())
		t.Puts(-1, 1, fmt.Sprintf("B%08d", v.BID))
		t.Puts(-1, 2, v.Name)
		t.Puts(-1, 3, v.Description)
	}
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportAssessments generates a report of all rlib.GLAccount accounts
func RRreportAssessments(ri *CSVReporterInfo) string {
	ri.D1 = time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	ri.D2 = time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)

	t := RRAssessmentsTable(ri)
	return t.GetTitle() + "\n" + t.SprintTable(ri.OutputFormat)
}

// RRAssessmentsTable generates a report of all rlib.GLAccount accounts
func RRAssessmentsTable(ri *CSVReporterInfo) rlib.Table {
	bid := ri.Bid
	d1 := ri.D1
	d2 := ri.D2
	rlib.InitBusinessFields(bid)
	rlib.RRdb.BizTypes[bid].GLAccounts = rlib.GetGLAccountMap(bid)
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(bid, d2, d1)
	rlib.Errcheck(err)
	defer rows.Close()

	var t rlib.Table
	t.SetTitle(fmt.Sprintf("Assessments:  BID = %s, from %s up to %s", rlib.IDtoString("B", bid), d1.Format(rlib.RRDATEFMT4), d2.Format(rlib.RRDATEFMT4)))
	t.Init()
	t.AddColumn("ASMID", 11, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("PASMID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("RAID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("RID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Rent Cycle", 13, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Proration Cycle", 13, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("AsmType", 50, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Account Rule", 80, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	for rows.Next() {
		var a rlib.Assessment
		rlib.ReadAssessments(rows, &a)
		t.AddRow()
		t.Puts(-1, 0, a.IDtoString())
		t.Puts(-1, 1, rlib.IDtoString("ASM", a.PASMID))
		t.Puts(-1, 2, rlib.IDtoString("RA", a.RAID))
		t.Puts(-1, 3, rlib.IDtoString("R", a.RID))
		t.Puts(-1, 4, rlib.RentalPeriodToString(a.RentCycle))
		t.Puts(-1, 5, rlib.RentalPeriodToString(a.ProrationCycle))
		t.Putf(-1, 6, a.Amount)
		t.Puts(-1, 7, rlib.RRdb.BizTypes[a.BID].GLAccounts[a.ATypeLID].Name)
		t.Puts(-1, 8, a.AcctRule)
	}
	rlib.Errcheck(rows.Err())
	t.TightenColumns()
	return t
}

// RRreportReceipts generates a report of all rlib.GLAccount accounts
func RRreportReceipts(ri *CSVReporterInfo) string {
	ri.D1 = time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	ri.D2 = time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
	t := RRReceiptsTable(ri)
	return t.SprintTable(ri.OutputFormat)
}

// RRReceiptsTable generates a report of all rlib.GLAccount accounts
func RRReceiptsTable(ri *CSVReporterInfo) rlib.Table {
	m := rlib.GetReceipts(ri.Bid, &ri.D1, &ri.D2)
	var t rlib.Table
	t.Init()
	t.AddColumn("Date", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("RCPTID", 12, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Parent RCPTID", 12, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("RAID", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("PMTID", 11, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Doc No", 25, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Account Rule", 50, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Comment", 50, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for _, a := range m {
		t.AddRow()
		t.Putd(-1, 0, a.Dt)
		t.Puts(-1, 1, a.IDtoString())
		t.Puts(-1, 2, rlib.IDtoString("RCPT", a.PRCPTID))
		t.Puts(-1, 3, rlib.IDtoString("RA", a.RAID))
		t.Puts(-1, 4, rlib.IDtoString("PMT", a.RAID))
		t.Puts(-1, 5, a.DocNo)
		t.Putf(-1, 6, a.Amount)
		t.Puts(-1, 7, a.AcctRule)
		t.Puts(-1, 8, a.Comment)
	}
	t.TightenColumns()
	return t
}

// RRreportInvoices generates a report of all rlib.GLAccount accounts
func RRreportInvoices(ri *CSVReporterInfo) string {
	var t rlib.Table
	t.Init()
	t.AddColumn("Date", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("InvoiceNo", 12, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("BID", 12, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Due Date", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("DeliveredBy", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	m := rlib.GetAllInvoicesInRange(ri.Bid, &Rcsv.DtStart, &Rcsv.DtStop)
	for i := 0; i < len(m); i++ {
		t.AddRow()
		t.Putd(-1, 0, m[i].Dt)
		t.Puts(-1, 1, m[i].IDtoString())
		t.Puts(-1, 2, rlib.IDtoString("B", m[i].BID))
		t.Putd(-1, 3, m[i].DtDue)
		t.Putf(-1, 4, m[i].Amount)
		t.Puts(-1, 5, m[i].DeliveredBy)
	}
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportDepository generates a report of all rlib.GLAccount accounts
func RRreportDepository(ri *CSVReporterInfo) string {
	m := rlib.GetAllDepositories(ri.Bid)
	var t rlib.Table
	t.Init()
	t.AddColumn("DEPID", 11, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("BID", 12, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("AccountNo", 12, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Name", 35, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	for i := 0; i < len(m); i++ {
		t.AddRow()
		t.Puts(-1, 0, rlib.IDtoString("DEP", m[i].DEPID))
		t.Puts(-1, 1, rlib.IDtoString("B", m[i].BID))
		t.Puts(-1, 2, m[i].AccountNo)
		t.Puts(-1, 3, m[i].Name)
	}
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportDepositMethods generates a report of all rlib.GLAccount accounts
func RRreportDepositMethods(ri *CSVReporterInfo) string {
	m := rlib.GetAllDepositMethods(ri.Bid)
	var t rlib.Table
	t.Init()
	t.AddColumn("DPMID", 11, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("BID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Name", 30, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	for i := 0; i < len(m); i++ {
		t.AddRow()
		t.Puts(-1, 0, rlib.IDtoString("DPM", m[i].DPMID))
		t.Puts(-1, 1, rlib.IDtoString("B", m[i].BID))
		t.Puts(-1, 2, m[i].Name)
	}
	t.TightenColumns()
	return t.SprintTable(ri.OutputFormat)
}

// RRreportDeposits generates a report of all rlib.GLAccount accounts
func RRreportDeposits(ri *CSVReporterInfo) string {
	m := rlib.GetAllDepositsInRange(ri.Bid, &Rcsv.DtStart, &Rcsv.DtStop)
	var t rlib.Table
	t.Init()
	t.AddColumn("Date", 10, rlib.CELLDATE, rlib.COLJUSTIFYLEFT)
	t.AddColumn("DEPID", 11, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("BID", 9, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, rlib.CELLFLOAT, rlib.COLJUSTIFYRIGHT)
	t.AddColumn("Receipts", 60, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
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
	return t.SprintTable(ri.OutputFormat)
}

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

// RRreportStringLists generates a report of all StringLists for the supplied business (ri.Bid)
func RRreportStringLists(ri *CSVReporterInfo) string {
	var (
		cat, val string
	)
	m := rlib.GetAllStringLists(ri.Bid)

	var t rlib.Table
	t.Init()
	t.AddColumn("SLSID", 20, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Category", 25, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Value", 50, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for i := 0; i < len(m); i++ {
		t.AddRow()
		t.Puts(-1, 0, m[i].Name)
		for j := 0; j < len(m[i].S); j++ {
			cat, val = getCategory(m[i].S[j].Value)
			t.AddRow()
			t.Puts(-1, 0, rlib.IDtoString("SLS", m[i].S[j].SLSID))
			t.Puts(-1, 1, cat)
			t.Puts(-1, 2, val)
		}
	}
	return t.SprintTable(ri.OutputFormat)
}

// ReportRentalAgreementPetToText returns a string representation of the chart of accts
func ReportRentalAgreementPetToText(p *rlib.RentalAgreementPet) string {
	end := ""
	if p.DtStop.Year() < rlib.YEARFOREVER {
		end = p.DtStop.Format(rlib.RRDATEINPFMT)
	}
	return fmt.Sprintf("PET%08d  RA%08d  %-25s  %-15s  %-15s  %-15s  %6.2f lb  %-10s  %-10s\n",
		p.PETID, p.RAID, p.Name, p.Type, p.Breed, p.Color, p.Weight, p.DtStart.Format(rlib.RRDATEINPFMT), end)
}

// RRreportRentalAgreementPets generates a report of all rlib.GLAccount accounts
func RRreportRentalAgreementPets(ri *CSVReporterInfo) string {
	m := rlib.GetAllRentalAgreementPets(ri.Raid)
	s := fmt.Sprintf("%-11s  %-10s  %-25s  %-15s  %-15s  %-15s  %-9s  %-10s  %-10s\n", "PETID", "RAID", "Name", "Type", "Breed", "Color", "Weight", "DtStart", "DtStop")
	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case rlib.RPTTEXT:
			s += ReportRentalAgreementPetToText(&m[i])
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportRentalAgreementPets: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// ReportNoteTypeToText returns a string representation of the chart of accts
func ReportNoteTypeToText(p *rlib.NoteType) string {
	return fmt.Sprintf("NT%08d  B%08d  %-50s\n",
		p.NTID, p.BID, p.Name)
}

// RRreportNoteTypes generates a report of all rlib.GLAccount accounts
func RRreportNoteTypes(ri *CSVReporterInfo) string {
	m := rlib.GetAllNoteTypes(ri.Bid)
	s := fmt.Sprintf("%-10s  %-9s  %-50s\n", "NTID", "BID", "Name")
	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case rlib.RPTTEXT:
			s += ReportNoteTypeToText(&m[i])
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportNoteTypes: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportSpecialties generates a report of all RentalSpecialties
func RRreportSpecialties(ri *CSVReporterInfo) string {
	s := fmt.Sprintf("%-11s  %-9s  %-30s  %10s  %-15s\n", "RSPID", "BID", "Name", "Fee", "Description")
	var xbiz rlib.XBusiness
	rlib.GetXBusiness(ri.Bid, &xbiz) // get its info

	// Order the rentableSpecialtyTypes into a known order.
	m := make([]int64, len(xbiz.US))
	i := 0
	for k := range xbiz.US {
		m[i] = k
		i++
	}
	for i := 0; i < len(m)-1; i++ {
		for j := i + 1; j < len(m); j++ {
			if xbiz.US[m[i]].Name > xbiz.US[m[j]].Name {
				m[i], m[j] = m[j], m[i]
			}
		}
	}

	// now print
	for i := 0; i < len(m); i++ {
		v := xbiz.US[m[i]]
		switch ri.OutputFormat {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("%11s  B%08d  %-30s  %10s  %s\n",
				v.IDtoString(), v.BID, v.Name, rlib.RRCommaf(v.Fee), v.Description)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSpecialties: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportSpecialtyAssigns generates a report of all RentalSpecialty Assignments accounts
func RRreportSpecialtyAssigns(ri *CSVReporterInfo) string {
	var xbiz rlib.XBusiness
	rlib.GetXBusiness(ri.Bid, &xbiz) // get its info

	s := fmt.Sprintf("%9s  %9s  %-30s  %10s  %10s\n", "BID", "RID", "Specialty Name", "DtStart", "DtStop")
	rows, err := rlib.RRdb.Prepstmt.GetAllRentableSpecialtyRefs.Query(ri.Bid)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a rlib.RentableSpecialtyRef
		rlib.Errcheck(rows.Scan(&a.BID, &a.RID, &a.RSPID, &a.DtStart, &a.DtStop, &a.LastModTime, &a.LastModBy))

		switch ri.OutputFormat {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("B%08d  R%08d  %-30s  %10s  %10s\n",
				a.BID, a.RID, xbiz.US[a.RSPID].Name, a.DtStart.Format(rlib.RRDATEFMT3), a.DtStop.Format(rlib.RRDATEFMT3))
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSpecialtyAssigns: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// RRreportSources generates a report of all rlib.GLAccount accounts
func RRreportSources(ri *CSVReporterInfo) string {
	m, _ := rlib.GetAllDemandSources(ri.Bid)

	s := fmt.Sprintf("%-9s  %-9s  %-35s  %-35s\n", "SourceSLSID", "BID", "Name", "Industry")
	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("S%08d  B%08d  %-35s  %-35s\n", m[i].SourceSLSID, m[i].BID, m[i].Name, m[i].Industry)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSources: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportRatePlans generates a report of all RateLists for the supplied business (ri.Bid)
func RRreportRatePlans(ri *CSVReporterInfo) string {
	m := rlib.GetAllRatePlans(ri.Bid)

	s := fmt.Sprintf("%-10s  %-9s  %-50s\n", "RPID", "BID", "Name")
	s += fmt.Sprintf("%-10s  %-9s  %-50s\n", "----", "---", "----")

	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("RP%08d  B%08d  %-50s\n", m[i].RPID, m[i].BID, m[i].Name)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportRatePlans: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportRatePlanRefs generates a report for RatePlans in business ri.Bid and valid on today's date
func RRreportRatePlanRefs(ri *CSVReporterInfo) string {
	funcname := "RRreportRatePlanRefs"
	var rp rlib.RatePlan
	var xbiz rlib.XBusiness
	rlib.GetXBusiness(ri.Bid, &xbiz)

	m := rlib.GetAllRatePlanRefsInRange(&ri.D1, &ri.D2)
	if len(m) == 0 {
		fmt.Printf("%s: could not load RatePlanRefs for timerange %s - %s\n", funcname, ri.D1.Format(rlib.RRDATEFMT4), ri.D2.Format(rlib.RRDATEFMT4))
		return ""
	}

	s := fmt.Sprintf("%-15s  %-11s  %-10s  %-10s  %-8s  %-6s  %-9s  %-9s  %-20s\n", "RatePlan", "RPRID", "DtStart", "DtStop", "MaxNoFee", "FeeAge", "Fee", "CancelFee", "PromoCode")
	s += fmt.Sprintf("%-15s  %-11s  %-10s  %-10s  %-8s  %-6s  %-9s  %-9s  %-20s\n", "--------", "-----", "----------", "----------", "--------", "------", "---------", "---------", "---------")

	for i := 0; i < len(m); i++ {
		p := m[i]
		rlib.GetRatePlan(p.RPID, &rp)
		rlib.GetRatePlanRefFull(p.RPRID, &p)
		switch ri.OutputFormat {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("%-15.15s  RPR%08d  %10s  %10s  %8d  %6d  %9.2f  %9.2f  %s\n",
				rp.Name, p.RPRID, p.DtStart.Format(rlib.RRDATEFMT4), p.DtStop.Format(rlib.RRDATEFMT4),
				p.MaxNoFeeUsers, p.FeeAppliesAge, p.AdditionalUserFee, p.CancellationFee, p.PromoCode)
			s += RRreportRatePlanRefRTRates(&p, &xbiz)
			s += "\n"
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportRatePlans: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportRatePlanRefRTRates generates a report of rates for all RentableTypes using RatePlanRef rpr.RPRID
func RRreportRatePlanRefRTRates(rpr *rlib.RatePlanRef, xbiz *rlib.XBusiness) string {
	var sporder []int64
	s := fmt.Sprintf("\n\t%-10s  %-10s  %-20s  %-8s    ", "RTID", "Style", "Name", "Rate")
	for _, v := range xbiz.US {
		s += fmt.Sprintf("  %-10.10s", v.Name)
		sporder = append(sporder, v.RSPID)
	}
	s += fmt.Sprintf("\n\t%-10s  %-10s  %-20s  %-8s  ", "----------", "----------", "--------------------", "----------")
	for range xbiz.US {
		s += fmt.Sprintf("  ----------")
	}
	s += "\n"

	// To perform the opertion you want
	for i := 0; i < len(rpr.RT); i++ {
		p := rpr.RT[i]
		s += fmt.Sprintf("\tRT%08d  %-10s  %-20s  ", p.RTID, xbiz.RT[p.RTID].Style, xbiz.RT[p.RTID].Name)
		if (p.FLAGS & rlib.FlRTRna) != 0 { // ...make sure it's not telling us to ignore this rentable type
			s += "n/a"
			continue // this RentableType is not affected
		}
		s1 := " "                           // assume p.Val is absolute
		if (p.FLAGS & rlib.FlRTRpct) != 0 { // if it's actually a percentage then...
			p.Val *= 100 // make the percentage more presentable
			s1 = "%"     // and add a % character
		}
		s += fmt.Sprintf("%8.2f %s  ", p.Val, s1)
		// Now add the Specialties
		m := rlib.GetAllRatePlanRefSPRates(p.RPRID, p.RTID) // almost certainly not in the order we want them
		for j := 0; j < len(m)-1; j++ {                     // we order them just to be sure
			if m[j].RSPID == sporder[j] { // if it's already in the right index for the column heading
				continue // then just continue
			}
			for k := j + 1; k < len(m); k++ { // need to find sporder[j] and put it in m[j]
				if m[k].RSPID == sporder[j] { // is this the one we're after?
					m[j], m[k] = m[k], m[j] // yes: swap m[j] and m[k]
					break                   // we're done with position j; break out of this loop and continue the j loop
				}
			}
		}
		// now m is ordered just like the column headings. Print out each amount
		for j := 0; j < len(m); j++ {
			s1 = " "
			fmt.Printf("m[%d]: RPRID=%d, RTID=%d, RSPID=%d, Val=%f\n", j, m[j].RPRID, m[j].RTID, m[j].RSPID, m[j].Val)
			v := m[j].Val
			if (m[j].FLAGS & rlib.FlSPRpct) != 0 { // if it's actually a percentage then...
				v *= 100 // make the percentage more presentable
				s1 = "%" // and add a % character
			}
			s += fmt.Sprintf("%8.2f %s  ", v, s1)
		}
		s += "\n"
	}
	return s
}
