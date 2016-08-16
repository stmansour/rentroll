package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"sort"
	"strings"
	"time"
)

// ReportBusinessToText returns a string representation of the supplied rlib.Business suitable for a text report
func ReportBusinessToText(p *rlib.Business) string {
	return fmt.Sprintf("%4d %6s    %s\n", p.BID, p.Designation, p.Name)
}

// ReportBusinessToHTML returns a string representation of the supplied rlib.Business suitable for HTML display
func ReportBusinessToHTML(p *rlib.Business) string {
	return fmt.Sprintf("<tr><td>%d</td><td%s></td><td>%s</td></tr>", p.BID, p.Designation, p.Name)
}

// RRreportBusiness generates a report of all Businesses defined in the database.
func RRreportBusiness(t int) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllBusinesses.Query()
	rlib.Errcheck(err)
	defer rows.Close()
	s := ""
	for rows.Next() {
		var p rlib.Business
		rlib.Errcheck(rows.Scan(&p.BID, &p.Designation, &p.Name, &p.DefaultRentalPeriod, &p.ParkingPermitInUse, &p.LastModTime, &p.LastModBy))
		switch t {
		case rlib.RPTTEXT:
			s += ReportBusinessToText(&p)
		case rlib.RPTHTML:
			s += ReportBusinessToHTML(&p)
		default:
			fmt.Printf("RRreportBusiness: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// ReportRentableTypeToText returns a string representation of the supplied rlib.RentableType suitable for a text report
func ReportRentableTypeToText(p *rlib.RentableType) string {
	s := fmt.Sprintf("%-10s  %-20s  %-25s  %-9s  %-12s", p.IDtoString(), p.Style, p.Name,
		rlib.RentalPeriodToString(p.RentCycle), rlib.RentalPeriodToString(p.Proration))
	for i := 0; i < len(p.MR); i++ {
		s += fmt.Sprintf("  |  %10s - %10s $ %8.2f", p.MR[i].DtStart.Format(rlib.RRDATEFMT4), p.MR[i].DtStop.Format(rlib.RRDATEFMT4), p.MR[i].MarketRate)
	}
	s += "\n"
	return s
}

// ReportRentableTypeToHTML returns a string representation of the supplied rlib.RentableType suitable for HTML display
func ReportRentableTypeToHTML(p rlib.RentableType) string {
	return fmt.Sprintf("<tr><td>%d</td><td%s></td><td%s></td></tr>", p.RTID, p.Style, p.Name)
}

// RRreportRentableTypes generates a report of all assessment types defined in the database.
func RRreportRentableTypes(t int, bid int64) string {
	m := rlib.GetBusinessRentableTypes(bid)
	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	s := fmt.Sprintf("%-10s  %-20s  %-25s  %-9s  %-12s  %s\n", "RTID", "Style", "Name", "RentCycle", "ProrateCycle", "Dt1 - Dt2 : Rate")
	s += fmt.Sprintf("%-10s  %-20s  %-25s  %-9s  %-12s  %s\n", "----", "-----", "----", "---------", "------------", "----------------")

	// To perform the opertion you want
	for _, k := range keys {
		i := int64(k)
		p := m[i]
		switch t {
		case rlib.RPTTEXT:
			s += ReportRentableTypeToText(&p)
		case rlib.RPTHTML:
			s += ReportRentableTypeToHTML(m[i])
		default:
			fmt.Printf("RRreportRentableTypes: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportRentableToText returns a string representation of the supplied rlib.Rentable suitable for a text report
func ReportRentableToText(p *rlib.Rentable) string {
	return fmt.Sprintf("%4d  %s\n",
		p.RID, p.Name)
}

// ReportRentableToHTML returns a string representation of the supplied rlib.Rentable suitable for a text report
func ReportRentableToHTML(p *rlib.Rentable) string {
	return fmt.Sprintf("<tr><td>%d</td><td>%s</td></tr>",
		p.RID, p.Name)
}

// RRreportRentables generates a report of all Businesses defined in the database.
func RRreportRentables(t int, bid int64) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllRentablesByBusiness.Query(bid)
	rlib.Errcheck(err)
	defer rows.Close()
	s := fmt.Sprintf(" RID  Name\n")
	for rows.Next() {
		var p rlib.Rentable
		rlib.Errcheck(rows.Scan(&p.RID, &p.BID, &p.Name, &p.AssignmentTime, &p.LastModTime, &p.LastModBy))
		switch t {
		case rlib.RPTTEXT:
			s += ReportRentableToText(&p)
		case rlib.RPTHTML:
			s += ReportRentableToHTML(&p)
		default:
			fmt.Printf("RRreportRentables: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// ReportXPersonToText returns a string representation of the supplied People suitable for a text report
func ReportXPersonToText(p *rlib.XPerson) string {
	return fmt.Sprintf("TC%08d  %4d  %12s  %-25s  %-13s, %-12s %-2s  %-25s\n",
		p.Trn.TCID, p.Trn.IsCompany, p.Trn.CellPhone, p.Trn.PrimaryEmail, p.Trn.LastName, p.Trn.FirstName, p.Trn.MiddleName, p.Trn.CompanyName)
}

// ReportXPersonToHTML returns a string representation of the supplied People suitable for a text report
func ReportXPersonToHTML(p *rlib.XPerson) string {
	return fmt.Sprintf("<tr><td>TC%08d</td><td>%s</td><td>%s</td><td>%s, %s %s</td></tr>",
		p.Trn.TCID, p.Trn.CellPhone, p.Trn.PrimaryEmail, p.Trn.LastName, p.Trn.FirstName, p.Trn.MiddleName)
}

// RRreportPeople generates a report of all Businesses defined in the database.
func RRreportPeople(t int) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllTransactants.Query()
	rlib.Errcheck(err)
	defer rows.Close()
	s := fmt.Sprintf("%10s  %4s  %-12s  %-25s  %-30s  %-25s\n", "TCID", "ISCO", "CELL PHONE", "PRIMARY EMAIL", "NAME", "COMPANY")
	for rows.Next() {
		var p rlib.XPerson
		rlib.ReadTransactants(rows, &p.Trn)
		rlib.GetXPerson(p.Trn.TCID, &p)
		switch t {
		case rlib.RPTTEXT:
			s += ReportXPersonToText(&p)
		case rlib.RPTHTML:
			s += ReportXPersonToHTML(&p)
		default:
			fmt.Printf("RRreportPeople: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// ReportRentalAgreementTemplateToText returns a string representation of the supplied People suitable for a text report
func ReportRentalAgreementTemplateToText(p *rlib.RentalAgreementTemplate) string {
	return fmt.Sprintf("%5d  B%08d   %s\n", p.RATID, p.BID, p.RentalTemplateNumber)
}

// ReportRentalAgreementTemplateToHTML returns a string representation of the supplied People suitable for a text report
func ReportRentalAgreementTemplateToHTML(p *rlib.RentalAgreementTemplate) string {
	return fmt.Sprintf("<tr><td>%5d</td><td>%5d</td><td>%s</td></tr>", p.RATID, p.BID, p.RentalTemplateNumber)
}

// RRreportRentalAgreementTemplates generates a report of all Businesses defined in the database.
func RRreportRentalAgreementTemplates(t int) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreementTemplates.Query()
	rlib.Errcheck(err)
	defer rows.Close()
	s := fmt.Sprintf("RATID  BID         RentalTemplateNumber\n")
	for rows.Next() {
		var p rlib.RentalAgreementTemplate
		rlib.Errcheck(rows.Scan(&p.RATID, &p.BID, &p.RentalTemplateNumber, &p.LastModTime, &p.LastModBy))
		switch t {
		case rlib.RPTTEXT:
			s += ReportRentalAgreementTemplateToText(&p)
		case rlib.RPTHTML:
			s += ReportRentalAgreementTemplateToHTML(&p)
		default:
			fmt.Printf("RRreportRentalAgreementTemplates: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// ReportRentalAgreementToText returns a string representation of the supplied People suitable for a text report
func ReportRentalAgreementToText(p *rlib.RentalAgreement, d1, d2 *time.Time) string {
	payors := strings.Join(p.GetPayorNameList(d1, d2), ", ")

	n := make(map[string]int, 0) // this is going to be our array of unique names
	for i := 0; i < len(p.R); i++ {
		m := p.R[i].R.GetUserNameList(d1, d2)
		for j := 0; j < len(m); j++ {
			n[m[j]] = 1 // the name is added to the list only one time, as a Key value
		}
	}

	var sa []string
	for k := range n {
		sa = append(sa, k)
	}
	sort.Strings(sa)
	users := strings.Join(sa, ", ")
	return fmt.Sprintf("%5d  %-40.40s  %-40.40s  %10s  %10s  %10s  %10s\n", p.RAID, payors, users,
		p.RentStart.Format(rlib.RRDATEFMT4), p.RentStop.Format(rlib.RRDATEFMT4),
		p.AgreementStart.Format(rlib.RRDATEFMT4), p.AgreementStop.Format(rlib.RRDATEFMT4))
}

// RRreportRentalAgreements generates a report of all Businesses defined in the database.
func RRreportRentalAgreements(t int, bid int64) string {
	rows, err := rlib.RRdb.Prepstmt.GetAllRentalAgreements.Query(bid)
	rlib.Errcheck(err)
	defer rows.Close()
	s := fmt.Sprintf("%5s  %-40s  %-40s  %10s  %10s  %10s  %10s\n", "RAID", "Payor", "User", "RentStart", "RentStop", "AgrmtStart", "AgrmtStop")
	var raid int64
	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var p rlib.RentalAgreement
		rlib.Errcheck(rows.Scan(&raid))
		p, err = rlib.GetXRentalAgreement(raid, &d1, &d2)
		if err != nil {
			rlib.Ulog("RRreportRentalAgreements: rlib.GetXRentalAgreement returned err = %v\n", err)
			continue
		}
		switch t {
		case rlib.RPTTEXT:
			s += ReportRentalAgreementToText(&p, &d1, &d2)
		case rlib.RPTHTML:
			fallthrough
		default:
			fmt.Printf("RRreportRentalAgreements: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// ReportChartOfAcctsToText returns a string representation of the chart of accts
func ReportChartOfAcctsToText(p rlib.GLAccount) string {
	s := ""
	collective := ""
	lm := rlib.GetLatestLedgerMarkerByLID(p.BID, p.LID)
	if lm.LMID == 0 {
		s = fmt.Sprintf("ReportChartOfAcctsToText: error getting latest LedgerMarker for L%08d\n", p.LID)
		return s
	}

	if rlib.GLCASH <= p.Type && p.Type <= rlib.GLLAST {
		s = fmt.Sprintf("%d", p.Type)
	} else {
		switch p.Type {
		case 1:
			s = "RArcv   "
		case 2:
			s = "RAsecdep"
		}
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
		collective = rlib.RRdb.BizTypes[p.BID].GLAccounts[p.PLID].Name
	}
	return fmt.Sprintf("%5d  %5d  %-8s  %12s  %-35.35s  %20.20s  %12.2f  %12s  %5d\n",
		p.LID, lm.LMID, s, p.GLNumber, p.Name, collective, lm.Balance, sp, p.RARequired)
}

// RRreportChartOfAccounts generates a report of all rlib.GLAccount accounts
func RRreportChartOfAccounts(t int, bid int64) string {
	rlib.InitBusinessFields(bid)
	rlib.RRdb.BizTypes[bid].GLAccounts = rlib.GetGLAccountMap(bid)

	// we need to sort the GLAccounts map so that our test output comparison will be the same every time
	// We'll sort by GLNumber.  First make an array of all the LIDs
	var a []int64
	for k := range rlib.RRdb.BizTypes[bid].GLAccounts {
		a = append(a, k)
	}
	// now sort based on GLNumber...
	m := rlib.RRdb.BizTypes[bid].GLAccounts // for notational convenience
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if m[a[i]].GLNumber > m[a[j]].GLNumber {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	s := fmt.Sprintf("%5s  %5s  %-8s  %12s  %-40s  %-35s  %-12s  %-12s  %-5s\n", "LID", "LMID", "TYPE", "GLNUMBER", "NAME", "COLLECTIVE", "BALANCE", "RAASSOC", "RARQD")
	for i := 0; i < len(a); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += ReportChartOfAcctsToText(m[a[i]])
		case rlib.RPTHTML:
			fmt.Printf("unimplemented\n")
		default:
			fmt.Printf("RRreportChartOfAccounts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportAssessmentToText returns a string representation of the chart of accts
func ReportAssessmentToText(p *rlib.Assessment) string {
	ra := "unassociated"
	if p.RAID > 0 {
		ra = fmt.Sprintf("RA%08d", p.RAID)
	}
	return fmt.Sprintf("%s  ASM%08d  %10s  R%08d  %4d  %4d  %9.2f  %-35s  %s\n",
		p.IDtoString(), p.PASMID, ra, p.RID, p.RentCycle, p.ProrationCycle,
		p.Amount, rlib.RRdb.BizTypes[p.BID].GLAccounts[p.ATypeLID].Name, p.AcctRule)
}

// ReportAssessmentToHTML returns a string representation of the chart of accts
func ReportAssessmentToHTML(p *rlib.Assessment) string {
	ra := "unassociated"
	if p.RAID > 0 {
		ra = fmt.Sprintf("RA%08d", p.RAID)
	}
	return fmt.Sprintf("<tr><td>ASM%08d</td><td>%12s</td><td>RA%08d</td><td%d</td><td>%8.2f</d></tr\n",
		p.ASMID, ra, p.RID, p.RentCycle, p.Amount)
}

// RRreportAssessments generates a report of all rlib.GLAccount accounts
func RRreportAssessments(t int, bid int64) string {
	d1 := time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
	rlib.InitBusinessFields(bid)
	rlib.RRdb.BizTypes[bid].GLAccounts = rlib.GetGLAccountMap(bid)
	rows, err := rlib.RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(bid, d2, d1)
	rlib.Errcheck(err)
	defer rows.Close()
	s := fmt.Sprintf("%11s  %11s  %10s  %9s  %4s  %4s  %9s  %-35s  %s\n", "ASMID", "PASMID", "RAID", "RID", "RCYC", "PCYC", "AMOUNT", "TYPE", "ACCOUNT RULE")
	for rows.Next() {
		var a rlib.Assessment
		rlib.ReadAssessments(rows, &a)
		switch t {
		case rlib.RPTTEXT:
			s += ReportAssessmentToText(&a)
		case rlib.RPTHTML:
			s += ReportAssessmentToHTML(&a)
		default:
			fmt.Printf("RRreportAssessments: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// ReportPaymentTypesToText returns a string representation of the rlib.PaymentType struct
func ReportPaymentTypesToText(p *rlib.PaymentType) string {
	return fmt.Sprintf("PT%08d     B%08d   %s\n",
		p.PMTID, p.BID, p.Name)
}

// ReportPaymentTypesToHTML returns a string representation of the rlib.PaymentType struct
func ReportPaymentTypesToHTML(p *rlib.PaymentType) string {
	return fmt.Sprintf("<tr><td>PT%08d</td><td>B%08d</td><td>%s</td></tr>\n",
		p.PMTID, p.BID, p.Name)
}

// RRreportPaymentTypes generates a report of all rlib.GLAccount accounts
func RRreportPaymentTypes(t int, bid int64) string {
	m := rlib.GetPaymentTypesByBusiness(bid)

	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	s := fmt.Sprintf("      PTID           BID   Name\n")
	for _, k := range keys {
		i := int64(k)
		v := m[i]
		switch t {
		case rlib.RPTTEXT:
			s += ReportPaymentTypesToText(&v)
		case rlib.RPTHTML:
			s += ReportPaymentTypesToHTML(&v)
		default:
			fmt.Printf("RRreportChartOfAccounts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportReceiptToText returns a string representation of the chart of accts
func ReportReceiptToText(p *rlib.Receipt) string {
	return fmt.Sprintf("RCPT%08d   %8.2f  %s\n",
		p.RCPTID, p.Amount, p.AcctRule)
}

// ReportReceiptToHTML returns a string representation of the chart of accts
func ReportReceiptToHTML(p *rlib.Receipt) string {
	return fmt.Sprintf("<tr><td>RCPT%08d</td><td>%8.2f</td><td>%s</d></tr\n",
		p.RCPTID, p.Amount, p.AcctRule)
}

// RRreportReceipts generates a report of all rlib.GLAccount accounts
func RRreportReceipts(t int, bid int64) string {
	d1 := time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
	m := rlib.GetReceipts(bid, &d1, &d2)
	s := fmt.Sprintf("      RCPTID     Amount  AcctRule\n")

	for _, a := range m {
		switch t {
		case rlib.RPTTEXT:
			s += ReportReceiptToText(&a)
		case rlib.RPTHTML:
			s += ReportReceiptToHTML(&a)
		default:
			fmt.Printf("RRreportReceipts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportCustomAttributeToText returns a string representation of the chart of accts
func ReportCustomAttributeToText(p *rlib.CustomAttribute) string {
	return fmt.Sprintf("%8d  %9d  %-25s  %25s %-10s\n",
		p.CID, p.Type, p.Name, p.Value, p.Units)
}

// RRreportCustomAttributes generates a report of all rlib.GLAccount accounts
func RRreportCustomAttributes(t int) string {
	rows, err := rlib.RRdb.Dbrr.Query("SELECT CID,Type,Name,Value,Units FROM CustomAttr")
	rlib.Errcheck(err)
	defer rows.Close()
	s := fmt.Sprintf("%-8s  %-9s  %-25s  %25s %-10s\n", "CID", "VALUETYPE", "Name", "Value", "Units")

	for rows.Next() {
		var a rlib.CustomAttribute
		rlib.Errcheck(rows.Scan(&a.CID, &a.Type, &a.Name, &a.Value, &a.Units))

		switch t {
		case rlib.RPTTEXT:
			s += ReportCustomAttributeToText(&a)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportReceipts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// ReportCustomAttributeRefToText returns a string representation of the chart of accts
func ReportCustomAttributeRefToText(p *rlib.CustomAttributeRef) string {
	return fmt.Sprintf("%6d  %8d  %8d\n",
		p.ElementType, p.ID, p.CID)
}

// RRreportCustomAttributeRefs generates a report of all rlib.GLAccount accounts
func RRreportCustomAttributeRefs(t int) string {
	rows, err := rlib.RRdb.Dbrr.Query("SELECT ElementType,ID,CID FROM CustomAttrRef")
	rlib.Errcheck(err)
	defer rows.Close()
	s := fmt.Sprintf("ELEMID        ID       CID\n")
	for rows.Next() {
		var a rlib.CustomAttributeRef
		rlib.Errcheck(rows.Scan(&a.ElementType, &a.ID, &a.CID))

		switch t {
		case rlib.RPTTEXT:
			s += ReportCustomAttributeRefToText(&a)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportReceipts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
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
func RRreportRentalAgreementPets(t int, raid int64) string {
	m := rlib.GetAllRentalAgreementPets(raid)
	s := fmt.Sprintf("%-11s  %-10s  %-25s  %-15s  %-15s  %-15s  %-9s  %-10s  %-10s\n", "PETID", "RAID", "Name", "Type", "Breed", "Color", "Weight", "DtStart", "DtStop")
	for i := 0; i < len(m); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += ReportRentalAgreementPetToText(&m[i])
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportRentalAgreementPets: unrecognized print format: %d\n", t)
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
func RRreportNoteTypes(t int, bid int64) string {
	m := rlib.GetAllNoteTypes(bid)
	s := fmt.Sprintf("%-10s  %-9s  %-50s\n", "NTID", "BID", "Name")
	for i := 0; i < len(m); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += ReportNoteTypeToText(&m[i])
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportNoteTypes: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// RRreportDepository generates a report of all rlib.GLAccount accounts
func RRreportDepository(t int, bid int64) string {
	m := rlib.GetAllDepositories(bid)
	s := fmt.Sprintf("%-11s  %-9s  %-12s %-50s\n", "DEPID", "BID", "AccountNo", "Name")
	for i := 0; i < len(m); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("DEP%08d  B%08d  %-12s %-50s\n", m[i].DEPID, m[i].BID, m[i].AccountNo, m[i].Name)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportNoteTypes: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// RRreportDeposits generates a report of all rlib.GLAccount accounts
func RRreportDeposits(t int, bid int64) string {
	d1, _ := rlib.StringToDate("1/1/1970")
	d2, _ := rlib.StringToDate("12/31/9999")
	m := rlib.GetAllDepositsInRange(bid, &d1, &d2)
	s := fmt.Sprintf("%-10s  %-11s  %-9s  %-8s %s\n", "Date", "DEPID", "BID", "Amount", "Receipts")
	for i := 0; i < len(m); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("%10s  DEP%08d  B%08d  %8.2f  ",
				m[i].Dt.Format(rlib.RRDATEINPFMT), m[i].DEPID, m[i].BID, m[i].Amount)
			for j := 0; j < len(m[i].DP); j++ {
				s += fmt.Sprintf("RCPT%08d ", m[i].DP[j].RCPTID)
			}
			s += "\n"
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportNoteTypes: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// RRreportInvoices generates a report of all rlib.GLAccount accounts
func RRreportInvoices(t int, bid int64) string {
	d1, _ := rlib.StringToDate("1/1/1970")
	d2, _ := rlib.StringToDate("12/31/9999")
	m := rlib.GetAllInvoicesInRange(bid, &d1, &d2)

	s := fmt.Sprintf("%-10s  %10s  %-9s  %-10s  %-8s  %-15s\n", "Date", "InvoiceNo", "BID", "Due Date", "Amount", "DeliveredBy")
	for i := 0; i < len(m); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("%10s  IN%08d  B%08d  %10s  %8.2f  %-15s\n",
				m[i].Dt.Format(rlib.RRDATEINPFMT), m[i].InvoiceNo, m[i].BID, m[i].DtDue.Format(rlib.RRDATEINPFMT), m[i].Amount, m[i].DeliveredBy)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportNoteTypes: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// RRreportSpecialties generates a report of all RentalSpecialties
func RRreportSpecialties(t int, bid int64) string {
	s := fmt.Sprintf("%-11s  %-9s  %-30s  %10s  %-15s\n", "RSPID", "BID", "Name", "Fee", "Description")
	var xbiz rlib.XBusiness
	rlib.GetXBusiness(bid, &xbiz) // get its info

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
		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("%11s  B%08d  %-30s  %10s  %s\n",
				v.IDtoString(), v.BID, v.Name, rlib.RRCommaf(v.Fee), v.Description)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSpecialties: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// RRreportSpecialtyAssigns generates a report of all RentalSpecialty Assignments accounts
func RRreportSpecialtyAssigns(t int, bid int64) string {
	var xbiz rlib.XBusiness
	rlib.GetXBusiness(bid, &xbiz) // get its info

	s := fmt.Sprintf("%9s  %9s  %-30s  %10s  %10s\n", "BID", "RID", "Specialty Name", "DtStart", "DtStop")
	rows, err := rlib.RRdb.Prepstmt.GetAllRentableSpecialtyRefs.Query(bid)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a rlib.RentableSpecialtyRef
		rlib.Errcheck(rows.Scan(&a.BID, &a.RID, &a.RSPID, &a.DtStart, &a.DtStop, &a.LastModTime, &a.LastModBy))

		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("B%08d  R%08d  %-30s  %10s  %10s\n",
				a.BID, a.RID, xbiz.US[a.RSPID].Name, a.DtStart.Format(rlib.RRDATEFMT3), a.DtStop.Format(rlib.RRDATEFMT3))
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSpecialtyAssigns: unrecognized print format: %d\n", t)
			return ""
		}
	}
	rlib.Errcheck(rows.Err())
	return s
}

// RRreportDepositMethods generates a report of all rlib.GLAccount accounts
func RRreportDepositMethods(t int, bid int64) string {
	m := rlib.GetAllDepositMethods(bid)

	s := fmt.Sprintf("%8s  %-10s  %s\n", "DPMID", "BID", "Name")
	for i := 0; i < len(m); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("%8d  B%08d  %s\n", m[i].DPMID, m[i].BID, m[i].Name)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportDepositMethods: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// RRreportSources generates a report of all rlib.GLAccount accounts
func RRreportSources(t int, bid int64) string {
	m, _ := rlib.GetAllDemandSources(bid)

	s := fmt.Sprintf("%-9s  %-9s  %-35s  %-35s\n", "DSID", "BID", "Name", "Industry")
	for i := 0; i < len(m); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("S%08d  B%08d  %-35s  %-35s\n", m[i].DSID, m[i].BID, m[i].Name, m[i].Industry)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSources: unrecognized print format: %d\n", t)
			return ""
		}
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

// RRreportStringLists generates a report of all StringLists for the supplied business (bid)
func RRreportStringLists(t int, bid int64) string {
	var (
		s        string
		cat, val string
	)
	m := rlib.GetAllStringLists(bid)

	for i := 0; i < len(m); i++ {
		s += fmt.Sprintf("String List Name:  %s\n", m[i].Name)
		s += fmt.Sprintf("%-11s  %-25s  %-50s\n", "SLSID", "Category", "Value")
		s += fmt.Sprintf("%-11s  %-25s  %-50s\n", "-----", "--------", "-----")
		for j := 0; j < len(m[i].S); j++ {
			cat, val = getCategory(m[i].S[j].Value)
			switch t {
			case rlib.RPTTEXT:
				s += fmt.Sprintf("SLS%08d  %-25s  %-50s\n", m[i].S[j].SLSID, cat, val)
			case rlib.RPTHTML:
				fmt.Printf("UNIMPLEMENTED\n")
			default:
				fmt.Printf("RRreportSources: unrecognized print format: %d\n", t)
				return ""
			}
		}
		s += "\n"
	}
	return s
}

// RRreportRatePlans generates a report of all RateLists for the supplied business (bid)
func RRreportRatePlans(t int, bid int64) string {
	m := rlib.GetAllRatePlans(bid)

	s := fmt.Sprintf("%-10s  %-9s  %-50s\n", "RPID", "BID", "Name")
	s += fmt.Sprintf("%-10s  %-9s  %-50s\n", "----", "---", "----")

	for i := 0; i < len(m); i++ {
		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("RP%08d  B%08d  %-50s\n", m[i].RPID, m[i].BID, m[i].Name)
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportRatePlans: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// RRreportRatePlanRefs generates a report for RatePlans in business bid and valid on today's date
func RRreportRatePlanRefs(t int, bid int64, d1, d2 *time.Time) string {
	funcname := "RRreportRatePlanRefs"
	var rp rlib.RatePlan
	var xbiz rlib.XBusiness
	rlib.GetXBusiness(bid, &xbiz)

	m := rlib.GetAllRatePlanRefsInRange(d1, d2)
	if len(m) == 0 {
		fmt.Printf("%s: could not load RatePlanRefs for timerange %s - %s\n", funcname, d1.Format(rlib.RRDATEFMT4), d2.Format(rlib.RRDATEFMT4))
		return ""
	}

	s := fmt.Sprintf("%-15s  %-11s  %-10s  %-10s  %-8s  %-6s  %-9s  %-9s  %-20s\n", "RatePlan", "RPRID", "DtStart", "DtStop", "MaxNoFee", "FeeAge", "Fee", "CancelFee", "PromoCode")
	s += fmt.Sprintf("%-15s  %-11s  %-10s  %-10s  %-8s  %-6s  %-9s  %-9s  %-20s\n", "--------", "-----", "----------", "----------", "--------", "------", "---------", "---------", "---------")

	for i := 0; i < len(m); i++ {
		p := m[i]
		rlib.GetRatePlan(p.RPID, &rp)
		rlib.GetRatePlanRefFull(p.RPRID, &p)
		switch t {
		case rlib.RPTTEXT:
			s += fmt.Sprintf("%-15.15s  RPR%08d  %10s  %10s  %8d  %6d  %9.2f  %9.2f  %s\n",
				rp.Name, p.RPRID, p.DtStart.Format(rlib.RRDATEFMT4), p.DtStop.Format(rlib.RRDATEFMT4),
				p.MaxNoFeeUsers, p.FeeAppliesAge, p.AdditionalUserFee, p.CancellationFee, p.PromoCode)
			s += RRreportRatePlanRefRTRates(t, &p, &xbiz)
			s += "\n"
		case rlib.RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportRatePlans: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// RRreportRatePlanRefRTRates generates a report of rates for all RentableTypes using RatePlanRef rpr.RPRID
func RRreportRatePlanRefRTRates(t int, rpr *rlib.RatePlanRef, xbiz *rlib.XBusiness) string {
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
