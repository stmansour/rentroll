package rlib

import (
	"fmt"
	"sort"
	"time"
)

// ReportBusinessToText returns a string representation of the supplied business suitable for a text report
func ReportBusinessToText(p *Business) string {
	return fmt.Sprintf("%4d %6s    %s\n", p.BID, p.Designation, p.Name)
}

// ReportBusinessToHTML returns a string representation of the supplied business suitable for HTML display
func ReportBusinessToHTML(p *Business) string {
	return fmt.Sprintf("<tr><td>%d</td><td%s></td><td>%s</td></tr>", p.BID, p.Designation, p.Name)
}

// RRreportBusiness generates a report of all businesses defined in the database.
func RRreportBusiness(t int) string {
	rows, err := RRdb.Prepstmt.GetAllBusinesses.Query()
	Errcheck(err)
	defer rows.Close()
	s := ""
	for rows.Next() {
		var p Business
		Errcheck(rows.Scan(&p.BID, &p.Designation, &p.Name, &p.DefaultOccupancyType, &p.ParkingPermitInUse, &p.LastModTime, &p.LastModBy))
		switch t {
		case RPTTEXT:
			s += ReportBusinessToText(&p)
		case RPTHTML:
			s += ReportBusinessToHTML(&p)
		default:
			fmt.Printf("RRreportBusiness: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}

// ReportAssessmentTypeToText returns a string representation of the supplied AssessmentType suitable for a text report
func ReportAssessmentTypeToText(p AssessmentType) string {
	return fmt.Sprintf("%4d  %6d   %s\n", p.ASMTID, p.OccupancyRqd, p.Name)
}

// ReportAssessmentTypeToHTML returns a string representation of the supplied AssessmentType suitable for HTML display
func ReportAssessmentTypeToHTML(p AssessmentType) string {
	return fmt.Sprintf("<tr><td>%d</td><td%s></td><td>%s</td></tr>", p.ASMTID, p.Name, p.Description)
}

// RRreportAssessmentTypes generates a report of all assessment types defined in the database.
func RRreportAssessmentTypes(t int) string {
	s := ""
	m := GetAssessmentTypes()

	fmt.Printf("Name  OccRqd   Description\n")
	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		i := int64(k)
		switch t {
		case RPTTEXT:
			s += ReportAssessmentTypeToText(m[i])
		case RPTHTML:
			s += ReportAssessmentTypeToHTML(m[i])
		default:
			fmt.Printf("RRreportAssessmentTypes: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportRentableTypeToText returns a string representation of the supplied RentableType suitable for a text report
func ReportRentableTypeToText(p RentableType) string {
	return fmt.Sprintf("%4d - %s  -  %s\n", p.RTID, p.Style, p.Name)
}

// ReportRentableTypeToHTML returns a string representation of the supplied RentableType suitable for HTML display
func ReportRentableTypeToHTML(p RentableType) string {
	return fmt.Sprintf("<tr><td>%d</td><td%s></td><td%s></td></tr>", p.RTID, p.Style, p.Name)
}

// RRreportRentableTypes generates a report of all assessment types defined in the database.
func RRreportRentableTypes(t int, bid int64) string {
	s := ""
	m := GetBusinessRentableTypes(bid)

	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	fmt.Printf("RTID   Style      Name\n")

	// To perform the opertion you want
	for _, k := range keys {
		i := int64(k)
		// 	fmt.Println("Key:", k, "Value:", m[k])
		// }

		// for _, v := range m {
		switch t {
		case RPTTEXT:
			s += ReportRentableTypeToText(m[i])
		case RPTHTML:
			s += ReportRentableTypeToHTML(m[i])
		default:
			fmt.Printf("RRreportRentableTypes: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportRentableToText returns a string representation of the supplied Rentable suitable for a text report
func ReportRentableToText(p *Rentable) string {
	return fmt.Sprintf("%4d  %5d  %s\n",
		p.RID, p.State, p.Name)
}

// ReportRentableToHTML returns a string representation of the supplied Rentable suitable for a text report
func ReportRentableToHTML(p *Rentable) string {
	return fmt.Sprintf("<tr><td>%d</td><td>%d</td><td>%s</td></tr>",
		p.RID, p.State, p.Name)
}

// RRreportRentables generates a report of all businesses defined in the database.
func RRreportRentables(t int, bid int64) string {
	rows, err := RRdb.Prepstmt.GetAllRentablesByBusiness.Query(bid)
	Errcheck(err)
	defer rows.Close()
	s := ""
	fmt.Printf(" RID  State  Name\n")
	for rows.Next() {
		var p Rentable
		Errcheck(rows.Scan(&p.RID, &p.RTID, &p.BID, &p.Name, &p.Assignment, &p.Report, &p.DefaultOccType, &p.OccType, &p.State, &p.LastModTime, &p.LastModBy))
		switch t {
		case RPTTEXT:
			s += ReportRentableToText(&p)
		case RPTHTML:
			s += ReportRentableToHTML(&p)
		default:
			fmt.Printf("RRreportRentables: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}

// ReportXPersonToText returns a string representation of the supplied People suitable for a text report
func ReportXPersonToText(p *XPerson) string {
	return fmt.Sprintf("%5d %5d %5d  %10s  %-25s  %s, %s %s\n",
		p.Trn.TCID, p.Tnt.TID, p.Pay.PID, p.Trn.CellPhone, p.Trn.PrimaryEmail, p.Trn.LastName, p.Trn.FirstName, p.Trn.MiddleName)
}

// ReportXPersonToHTML returns a string representation of the supplied People suitable for a text report
func ReportXPersonToHTML(p *XPerson) string {
	return fmt.Sprintf("<tr><td>%5d</td><td>%5d</td><td>%5d</td><td>%s</td><td>%s</td><td>%s, %s %s</td></tr>",
		p.Trn.TCID, p.Tnt.TID, p.Pay.PID, p.Trn.CellPhone, p.Trn.PrimaryEmail, p.Trn.LastName, p.Trn.FirstName, p.Trn.MiddleName)
}

// RRreportPeople generates a report of all businesses defined in the database.
func RRreportPeople(t int) string {
	rows, err := RRdb.Prepstmt.GetAllTransactants.Query()
	Errcheck(err)
	defer rows.Close()
	fmt.Printf(" TCID   TID   PID  Cell Phone    Primary Email              Name\n")
	s := ""
	for rows.Next() {
		var p XPerson
		Errcheck(rows.Scan(&p.Trn.TCID, &p.Trn.TID, &p.Trn.PID, &p.Trn.PRSPID, &p.Trn.FirstName, &p.Trn.MiddleName, &p.Trn.LastName, &p.Trn.PrimaryEmail, &p.Trn.SecondaryEmail, &p.Trn.WorkPhone, &p.Trn.CellPhone, &p.Trn.Address, &p.Trn.Address2, &p.Trn.City, &p.Trn.State, &p.Trn.PostalCode, &p.Trn.Country, &p.Trn.LastModTime, &p.Trn.LastModBy))
		GetXPerson(p.Trn.TCID, &p)
		switch t {
		case RPTTEXT:
			s += ReportXPersonToText(&p)
		case RPTHTML:
			s += ReportXPersonToHTML(&p)
		default:
			fmt.Printf("RRreportPeople: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}

// ReportRentalAgreementTemplateToText returns a string representation of the supplied People suitable for a text report
func ReportRentalAgreementTemplateToText(p *RentalAgreementTemplate) string {
	return fmt.Sprintf("%5d  %6d   %s\n", p.RATID, p.RentalAgreementType, p.ReferenceNumber)
}

// ReportRentalAgreementTemplateToHTML returns a string representation of the supplied People suitable for a text report
func ReportRentalAgreementTemplateToHTML(p *RentalAgreementTemplate) string {
	return fmt.Sprintf("<tr><td>%5d</td><td>%5d</td><td>%s</td></tr>", p.RATID, p.RentalAgreementType, p.ReferenceNumber)
}

// RRreportRentalAgreementTemplates generates a report of all businesses defined in the database.
func RRreportRentalAgreementTemplates(t int) string {
	rows, err := RRdb.Prepstmt.GetAllRentalAgreementTemplates.Query()
	Errcheck(err)
	defer rows.Close()
	fmt.Printf("RATID  RAType  TemplateName\n")
	s := ""
	for rows.Next() {
		var p RentalAgreementTemplate
		Errcheck(rows.Scan(&p.RATID, &p.ReferenceNumber, &p.RentalAgreementType, &p.LastModTime, &p.LastModBy))
		switch t {
		case RPTTEXT:
			s += ReportRentalAgreementTemplateToText(&p)
		case RPTHTML:
			s += ReportRentalAgreementTemplateToHTML(&p)
		default:
			fmt.Printf("RRreportRentalAgreementTemplates: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}

// ReportRentalAgreementToText returns a string representation of the supplied People suitable for a text report
func ReportRentalAgreementToText(p *RentalAgreement) string {
	return fmt.Sprintf("%5d   %6d  %6d (%s %s)\n",
		p.RAID, p.PrimaryTenant, p.P[0].Trn.PID, p.P[0].Trn.FirstName, p.P[0].Trn.LastName)
}

// ReportRentalAgreementToHTML returns a string representation of the supplied People suitable for a text report
func ReportRentalAgreementToHTML(p *RentalAgreement) string {
	return fmt.Sprintf("<tr><td>%5d</td><td>%6d</td><td%6d (%s %s)</td>\n",
		p.RAID, p.PrimaryTenant, p.P[0].Trn.PID, p.P[0].Trn.FirstName, p.P[0].Trn.LastName)
}

// RRreportRentalAgreements generates a report of all businesses defined in the database.
func RRreportRentalAgreements(t int, bid int64) string {
	rows, err := RRdb.Prepstmt.GetAllRentalAgreements.Query(bid)
	Errcheck(err)
	defer rows.Close()
	fmt.Printf("RAID     Payor       Tenant\n")
	s := ""
	var raid int64
	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var p RentalAgreement

		Errcheck(rows.Scan(&raid))
		p, err = GetXRentalAgreement(raid, &d1, &d2)
		if err != nil {
			Ulog("RRreportRentalAgreements: GetXRentalAgreement returned err = %v\n", err)
			continue
		}
		switch t {
		case RPTTEXT:
			s += ReportRentalAgreementToText(&p)
		case RPTHTML:
			s += ReportRentalAgreementToHTML(&p)
		default:
			fmt.Printf("RRreportRentalAgreements: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}

// ReportChartOfAcctsToText returns a string representation of the chart of accts
func ReportChartOfAcctsToText(p *LedgerMarker) string {
	s := ""
	if DFLTCASH <= p.Type && p.Type <= DFLTLAST {
		s = fmt.Sprintf("%4d", p.Type)
	}
	return fmt.Sprintf("%5d  %4s  %12s   %12.2f   %s\n",
		p.LMID, s, p.GLNumber, p.Balance, p.Name)
}

// ReportChartOfAcctsToHTML returns a string representation of the chart of accts
func ReportChartOfAcctsToHTML(p *LedgerMarker) string {
	s := ""
	if DFLTCASH <= p.Type && p.Type <= DFLTLAST {
		s = fmt.Sprintf("%d", p.Type)
	}
	return fmt.Sprintf("<tr><td>%5d</td><td>%4s</td><td>%s</td><td>%12.2f</td><td>%s</td></tr>\n",
		p.LMID, s, p.GLNumber, p.Balance, p.Name)
}

// RRreportChartOfAccounts generates a report of all ledger accounts
func RRreportChartOfAccounts(t int, bid int64) string {
	m := GetLedgerMarkerInitList(bid)
	//                               123456789012
	fmt.Printf("  LID   Type  GLAccountNo         Amount   Name\n")
	s := ""
	for i := 0; i < len(m); i++ {
		switch t {
		case RPTTEXT:
			s += ReportChartOfAcctsToText(&m[i])
		case RPTHTML:
			s += ReportChartOfAcctsToHTML(&m[i])
		default:
			fmt.Printf("RRreportChartOfAccounts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportAssessmentToText returns a string representation of the chart of accts
func ReportAssessmentToText(p *Assessment) string {
	ra := "unassociated"
	if p.RAID > 0 {
		ra = fmt.Sprintf("RA%08d", p.RAID)
	}
	return fmt.Sprintf("ASM%08d  %12s  R%08d     %2d  %9.2f\n",
		p.ASMID, ra, p.RID, p.Frequency, p.Amount)
}

// ReportAssessmentToHTML returns a string representation of the chart of accts
func ReportAssessmentToHTML(p *Assessment) string {
	ra := "unassociated"
	if p.RAID > 0 {
		ra = fmt.Sprintf("RA%08d", p.RAID)
	}
	return fmt.Sprintf("<tr><td>ASM%08d</td><td>%12s</td><td>RA%08d</td><td%d</td><td>%8.2f</d></tr\n",
		p.ASMID, ra, p.RID, p.Frequency, p.Amount)
}

// RRreportAssessments generates a report of all ledger accounts
func RRreportAssessments(t int, bid int64) string {
	d1 := time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
	rows, err := RRdb.Prepstmt.GetAllAssessmentsByBusiness.Query(bid, d2, d1)
	Errcheck(err)
	defer rows.Close()
	fmt.Printf("      ASMID          RAID        RID   Freq     Amount\n")
	s := ""
	for rows.Next() {
		var a Assessment
		Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.ASMTID, &a.RAID, &a.Amount,
			&a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod, &a.AcctRule, &a.Comment,
			&a.LastModTime, &a.LastModBy))
		switch t {
		case RPTTEXT:
			s += ReportAssessmentToText(&a)
		case RPTHTML:
			s += ReportAssessmentToHTML(&a)
		default:
			fmt.Printf("RRreportAssessments: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}

// ReportPaymentTypesToText returns a string representation of the PaymentType struct
func ReportPaymentTypesToText(p *PaymentType) string {
	return fmt.Sprintf("PT%08d     B%08d   %s\n",
		p.PMTID, p.BID, p.Name)
}

// ReportPaymentTypesToHTML returns a string representation of the PaymentType struct
func ReportPaymentTypesToHTML(p *PaymentType) string {
	return fmt.Sprintf("<tr><td>PT%08d</td><td>B%08d</td><td>%s</td></tr>\n",
		p.PMTID, p.BID, p.Name)
}

// RRreportPaymentTypes generates a report of all ledger accounts
func RRreportPaymentTypes(t int, bid int64) string {
	m := GetPaymentTypesByBusiness(bid)

	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	fmt.Printf("      PTID           BID   Name\n")
	s := ""
	for _, k := range keys {
		i := int64(k)
		v := m[i]
		switch t {
		case RPTTEXT:
			s += ReportPaymentTypesToText(&v)
		case RPTHTML:
			s += ReportPaymentTypesToHTML(&v)
		default:
			fmt.Printf("RRreportChartOfAccounts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportReceiptToText returns a string representation of the chart of accts
func ReportReceiptToText(p *Receipt) string {
	return fmt.Sprintf("RCPT%08d   %8.2f  %s\n",
		p.RCPTID, p.Amount, p.AcctRule)
}

// ReportReceiptToHTML returns a string representation of the chart of accts
func ReportReceiptToHTML(p *Receipt) string {
	return fmt.Sprintf("<tr><td>RCPT%08d</td><td>%8.2f</td><td>%s</d></tr\n",
		p.RCPTID, p.Amount, p.AcctRule)
}

// RRreportReceipts generates a report of all ledger accounts
func RRreportReceipts(t int, bid int64) string {
	d1 := time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC)
	m := GetReceipts(bid, &d1, &d2)
	fmt.Printf("      RCPTID     Amount  AcctRule\n")
	s := ""
	for _, a := range m {
		switch t {
		case RPTTEXT:
			s += ReportReceiptToText(&a)
		case RPTHTML:
			s += ReportReceiptToHTML(&a)
		default:
			fmt.Printf("RRreportReceipts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	return s
}

// ReportCustomAttributeToText returns a string representation of the chart of accts
func ReportCustomAttributeToText(p *CustomAttribute) string {
	return fmt.Sprintf("%8d   %4d   %s, %s\n",
		p.CID, p.Type, p.Name, p.Value)
}

// RRreportCustomAttributes generates a report of all ledger accounts
func RRreportCustomAttributes(t int) string {
	rows, err := RRdb.dbrr.Query("SELECT CID,Type,Name,Value FROM customattr")
	Errcheck(err)
	defer rows.Close()
	fmt.Printf("     CID   TYPE   Name, Value\n")
	s := ""
	for rows.Next() {
		var a CustomAttribute
		Errcheck(rows.Scan(&a.CID, &a.Type, &a.Name, &a.Value))

		switch t {
		case RPTTEXT:
			s += ReportCustomAttributeToText(&a)
		case RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportReceipts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}

// ReportCustomAttributeRefToText returns a string representation of the chart of accts
func ReportCustomAttributeRefToText(p *CustomAttributeRef) string {
	return fmt.Sprintf("%6d  %8d  %8d\n",
		p.ElementType, p.ID, p.CID)
}

// RRreportCustomAttributeRefs generates a report of all ledger accounts
func RRreportCustomAttributeRefs(t int) string {
	rows, err := RRdb.dbrr.Query("SELECT ElementType,ID,CID FROM customattrref")
	Errcheck(err)
	defer rows.Close()
	fmt.Printf("ELEMID        ID       CID\n")
	s := ""
	for rows.Next() {
		var a CustomAttributeRef
		Errcheck(rows.Scan(&a.ElementType, &a.ID, &a.CID))

		switch t {
		case RPTTEXT:
			s += ReportCustomAttributeRefToText(&a)
		case RPTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportReceipts: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}
