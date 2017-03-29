package rcsv

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
	"rentroll/rrpt"
	"sort"
	"strings"
	"time"
)

// RRreportRentableTypesTable generates a table object of all Rentable Types defined in the database, for all businesses.
func RRreportRentableTypesTable(ri *rrpt.ReporterInfo) gotable.Table {
	funcname := "RRreportRentableTypesTable"

	m := rlib.GetBusinessRentableTypes(ri.Bid)
	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	var t gotable.Table
	t.Init()
	err := rrpt.TableReportHeaderBlock(&t, "Rentable Types", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
	}

	t.AddColumn("RTID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                    // 0
	t.AddColumn("Style", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                   // 1
	t.AddColumn("Name", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                    // 2
	t.AddColumn("Rent Cycle", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)               // 3
	t.AddColumn("Proration Cycle", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)          // 4
	t.AddColumn("GSRPC", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                    // 5
	t.AddColumn("Manage To Budget", 3, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)         // 6
	t.AddColumn("Dt1 - Dt2 : Market Rate", 96, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT) // 7

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
			s += fmt.Sprintf("%8s - %8s: $%8.2f", p.MR[i].DtStart.Format(rlib.RRDATEFMT2),
				p.MR[i].DtStop.Format(rlib.RRDATEFMT2), p.MR[i].MarketRate)
			if i+1 < len(p.MR) {
				s += ",  "
			}
		}
		t.Puts(-1, 7, s)
	}
	t.TightenColumns()
	return t
}

// RRreportRentableTypes generates a report of all Rentable Types defined in the database, for all businesses.
func RRreportRentableTypes(ri *rrpt.ReporterInfo) string {
	t := RRreportRentableTypesTable(ri)
	return rrpt.ReportToString(&t, ri)
}

// RRreportRentablesTable generates a table of all rentables for BUD defined in the database.
func RRreportRentablesTable(ri *rrpt.ReporterInfo) gotable.Table {
	funcname := "RRreportRentablesTable"

	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(ri.Bid)
	rlib.Errcheck(err)
	defer rows.Close()

	var t gotable.Table
	t.Init()
	err = rrpt.TableReportHeaderBlock(&t, "Rentables", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
	}

	t.AddColumn("RID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Name", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Assignment Time", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

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
		t.Puts(-1, 1, p.RentableName)
		t.Puts(-1, 2, s)
	}
	rlib.Errcheck(rows.Err())
	t.TightenColumns()
	return t
}

// RRreportRentables generates a report of all Businesses defined in the database.
func RRreportRentables(ri *rrpt.ReporterInfo) string {
	t := RRreportRentablesTable(ri)
	return rrpt.ReportToString(&t, ri)
}

// RRreportRentalAgreementTemplatesTable generates a table object for all rental agreement templates
func RRreportRentalAgreementTemplatesTable(ri *rrpt.ReporterInfo) gotable.Table {
	funcname := "RRreportRentalAgreementTemplatesTable"

	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementTemplates.Query()
	rlib.Errcheck(err)
	defer rows.Close()
	var t gotable.Table

	t.Init()
	err = rrpt.TableReportHeaderBlock(&t, "Rental Agreement Templates", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
	}

	t.AddColumn("BID", 9, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("RA Template ID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("RA Template Name", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
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
	return t
}

// RRreportRentalAgreementTemplates generates a report for all rental agreement templates
func RRreportRentalAgreementTemplates(ri *rrpt.ReporterInfo) string {
	t := RRreportRentalAgreementTemplatesTable(ri)
	return rrpt.ReportToString(&t, ri)
}

// RRreportRentalAgreementsTable generates a table object for All rental agreements related with Business
func RRreportRentalAgreementsTable(ri *rrpt.ReporterInfo) gotable.Table {
	funcname := "RRreportRentalAgreementsTable"

	var err error
	totalErrs := 0
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreements.Query(ri.Bid)
	rlib.Errcheck(err)
	defer rows.Close()

	var t gotable.Table
	t.Init()
	err = rrpt.TableReportHeaderBlock(&t, "Rental Agreement", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
	}

	t.AddColumn("RAID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Payor", 60, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("User", 60, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Agreement Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Agreement Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Possession Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Possession Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Rent Start", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Rent Stop", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Rent Cycle Epoch", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Renewal", 2, gotable.CELLINT, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Unspecified Adults", 6, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Unspecified Children", 6, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Special Provisions", 40, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Notes", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	var raid int64
	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var p rlib.RentalAgreement
		rlib.Errcheck(rows.Scan(&raid))
		p, err = rlib.GetXRentalAgreement(raid, &d1, &d2)
		if err != nil {
			totalErrs++
			rlib.Ulog("RRreportRentalAgreements: rlib.GetXRentalAgreement returned err = %v\n", err)
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
	if totalErrs > 0 {
		errMsg := fmt.Sprintf("Encountered %d errors while creating this report. See log.", totalErrs)
		t.SetSection3(errMsg)

		// use section3 for errors and apply red color
		cssListSection3 := []*gotable.CSSProperty{
			{Name: "color", Value: "red"},
			{Name: "font-family", Value: "monospace"},
		}
		t.SetSection3CSS(cssListSection3)
	}
	return t
}

// RRreportRentalAgreements generates a report of all Businesses defined in the database.
func RRreportRentalAgreements(ri *rrpt.ReporterInfo) string {
	t := RRreportRentalAgreementsTable(ri)
	return rrpt.ReportToString(&t, ri)
}

// RRreportReceipts generates a report of all rlib.GLAccount accounts
func RRreportReceipts(ri *rrpt.ReporterInfo) string {
	// ri.D1 = time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	// ri.D2 = time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
	t := RRReceiptsTable(ri)
	return rrpt.ReportToString(&t, ri)
}

// RRReceiptsTable generates a report of all rlib.GLAccount accounts
func RRReceiptsTable(ri *rrpt.ReporterInfo) gotable.Table {
	m := rlib.GetReceipts(ri.Bid, &ri.D1, &ri.D2)
	var t gotable.Table
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true
	ri.BlankLineAfterRptName = true
	rrpt.TableReportHeaderBlock(&t, "Receipts", "RRReceiptsTable", ri)
	t.Init()
	t.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("RCPTID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Parent RCPTID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("PMTID", 11, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Doc No", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Account Rule", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Comment", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	for _, a := range m {
		t.AddRow()
		t.Putd(-1, 0, a.Dt)
		t.Puts(-1, 1, a.IDtoString())
		t.Puts(-1, 2, rlib.IDtoString("RCPT", a.PRCPTID))
		t.Puts(-1, 3, rlib.IDtoString("PMT", a.PMTID))
		t.Puts(-1, 4, a.DocNo)
		t.Putf(-1, 5, a.Amount)
		t.Puts(-1, 6, a.AcctRule)
		t.Puts(-1, 7, a.Comment)
	}
	t.TightenColumns()
	return t
}

// RRreportInvoices generates a report of all rlib.GLAccount accounts
func RRreportInvoices(ri *rrpt.ReporterInfo) string {
	var t gotable.Table
	t.Init()
	t.SetTitle(rrpt.ReportHeaderBlock("Invoices", "RRreportInvoices", ri))
	t.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("InvoiceNo", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("BID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Due Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("DeliveredBy", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

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
	return rrpt.ReportToString(&t, ri)
}

// RRreportDeposits generates a report of all rlib.Deposit
func RRreportDeposits(ri *rrpt.ReporterInfo) string {
	funcname := "RRreportDeposits"
	m := rlib.GetAllDepositsInRange(ri.Bid, &Rcsv.DtStart, &Rcsv.DtStop)
	var t gotable.Table
	t.Init()

	err := rrpt.TableReportHeaderBlock(&t, "Deposit", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
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
func RRreportStringListsTable(ri *rrpt.ReporterInfo) gotable.Table {
	funcname := "RRreportStringListsTable"

	var (
		cat, val string
	)
	m := rlib.GetAllStringLists(ri.Bid)

	var t gotable.Table
	t.Init()
	err := rrpt.TableReportHeaderBlock(&t, "String Lists", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
	}

	t.AddColumn("SLSID", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Category", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Value", 50, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

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
	return t
}

// RRreportStringLists generates a report of all StringLists for the supplied business (ri.Bid)
func RRreportStringLists(ri *rrpt.ReporterInfo) string {
	t := RRreportStringListsTable(ri)
	return rrpt.ReportToString(&t, ri)
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
func RRreportRentalAgreementPets(ri *rrpt.ReporterInfo) string {
	m := rlib.GetAllRentalAgreementPets(ri.Raid)
	s := fmt.Sprintf("%-11s  %-10s  %-25s  %-15s  %-15s  %-15s  %-9s  %-10s  %-10s\n", "PETID", "RAID", "Name", "Type", "Breed", "Color", "Weight", "DtStart", "DtStop")
	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case gotable.TABLEOUTTEXT:
			s += ReportRentalAgreementPetToText(&m[i])
		case gotable.TABLEOUTHTML:
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
func RRreportNoteTypes(ri *rrpt.ReporterInfo) string {
	m := rlib.GetAllNoteTypes(ri.Bid)
	s := fmt.Sprintf("%-10s  %-9s  %-50s\n", "NTID", "BID", "Name")
	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case gotable.TABLEOUTTEXT:
			s += ReportNoteTypeToText(&m[i])
		case gotable.TABLEOUTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportNoteTypes: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportSpecialties generates a report of all RentalSpecialties
func RRreportSpecialties(ri *rrpt.ReporterInfo) string {
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
		case gotable.TABLEOUTTEXT:
			s += fmt.Sprintf("%11s  B%08d  %-30s  %10s  %s\n",
				v.IDtoString(), v.BID, v.Name, rlib.RRCommaf(v.Fee), v.Description)
		case gotable.TABLEOUTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSpecialties: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportSpecialtyAssigns generates a report of all RentalSpecialty Assignments accounts
func RRreportSpecialtyAssigns(ri *rrpt.ReporterInfo) string {
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
		case gotable.TABLEOUTTEXT:
			s += fmt.Sprintf("B%08d  R%08d  %-30s  %10s  %10s\n",
				a.BID, a.RID, xbiz.US[a.RSPID].Name, a.DtStart.Format(rlib.RRDATEFMT3), a.DtStop.Format(rlib.RRDATEFMT3))
		case gotable.TABLEOUTHTML:
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
func RRreportSources(ri *rrpt.ReporterInfo) string {
	m, _ := rlib.GetAllDemandSources(ri.Bid)

	s := fmt.Sprintf("%-9s  %-9s  %-35s  %-35s\n", "SourceSLSID", "BID", "Name", "Industry")
	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case gotable.TABLEOUTTEXT:
			s += fmt.Sprintf("S%08d  B%08d  %-35s  %-35s\n", m[i].SourceSLSID, m[i].BID, m[i].Name, m[i].Industry)
		case gotable.TABLEOUTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSources: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportRatePlans generates a report of all RateLists for the supplied business (ri.Bid)
func RRreportRatePlans(ri *rrpt.ReporterInfo) string {
	m := rlib.GetAllRatePlans(ri.Bid)

	s := fmt.Sprintf("%-10s  %-9s  %-50s\n", "RPID", "BID", "Name")
	s += fmt.Sprintf("%-10s  %-9s  %-50s\n", "----", "---", "----")

	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case gotable.TABLEOUTTEXT:
			s += fmt.Sprintf("RP%08d  B%08d  %-50s\n", m[i].RPID, m[i].BID, m[i].Name)
		case gotable.TABLEOUTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportRatePlans: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}

// RRreportRatePlanRefs generates a report for RatePlans in business ri.Bid and valid on today's date
func RRreportRatePlanRefs(ri *rrpt.ReporterInfo) string {
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
		case gotable.TABLEOUTTEXT:
			s += fmt.Sprintf("%-15.15s  RPR%08d  %10s  %10s  %8d  %6d  %9.2f  %9.2f  %s\n",
				rp.Name, p.RPRID, p.DtStart.Format(rlib.RRDATEFMT4), p.DtStop.Format(rlib.RRDATEFMT4),
				p.MaxNoFeeUsers, p.FeeAppliesAge, p.AdditionalUserFee, p.CancellationFee, p.PromoCode)
			s += RRreportRatePlanRefRTRates(&p, &xbiz)
			s += "\n"
		case gotable.TABLEOUTHTML:
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
