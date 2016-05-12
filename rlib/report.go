package rlib

import (
	"fmt"
	"sort"
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
	return fmt.Sprintf("%4d - %s\n", p.ASMTID, p.Name)
}

// ReportAssessmentTypeToHTML returns a string representation of the supplied AssessmentType suitable for HTML display
func ReportAssessmentTypeToHTML(p AssessmentType) string {
	return fmt.Sprintf("<tr><td>%d</td><td%s></td><td>%s</td></tr>", p.ASMTID, p.Name, p.Description)
}

// RRreportAssessmentTypes generates a report of all assessment types defined in the database.
func RRreportAssessmentTypes(t int) string {
	s := ""
	m := GetAssessmentTypes()

	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		i := int64(k)
		// 	fmt.Println("Key:", k, "Value:", m[k])
		// }

		// for _, v := range m {
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
	return fmt.Sprintf("%4d - %s\n", p.RID, p.Name)
}

// ReportRentableToHTML returns a string representation of the supplied Rentable suitable for a text report
func ReportRentableToHTML(p *Rentable) string {
	return fmt.Sprintf("<tr><td>%d</td><td>%s</td></tr>", p.RID, p.Name)
}

// RRreportRentables generates a report of all businesses defined in the database.
func RRreportRentables(t int, bid int64) string {
	rows, err := RRdb.Prepstmt.GetAllRentablesByBusiness.Query(bid)
	Errcheck(err)
	defer rows.Close()
	s := ""
	for rows.Next() {
		var p Rentable
		Errcheck(rows.Scan(&p.RID, &p.RTID, &p.BID, &p.Name, &p.Assignment, &p.Report, &p.DefaultOccType, &p.OccType, &p.LastModTime, &p.LastModBy))
		switch t {
		case RPTTEXT:
			s += ReportRentableToText(&p)
		case RPTHTML:
			s += ReportRentableToHTML(&p)
		default:
			fmt.Printf("RRreportBusiness: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}

// ReportXPersonToText returns a string representation of the supplied People suitable for a text report
func ReportXPersonToText(p *XPerson) string {
	return fmt.Sprintf("%5d %5d %5d  %s, %s %s\n", p.Trn.TCID, p.Tnt.TID, p.Pay.PID, p.Trn.LastName, p.Trn.FirstName, p.Trn.MiddleName)
}

// ReportXPersonToHTML returns a string representation of the supplied People suitable for a text report
func ReportXPersonToHTML(p *XPerson) string {
	return fmt.Sprintf("<tr><td>%5d</td><td>%5d</td><td>%5d</td><td>%s, %s %s</td></tr>", p.Trn.TCID, p.Tnt.TID, p.Pay.PID, p.Trn.LastName, p.Trn.FirstName, p.Trn.MiddleName)
}

// RRreportPeople generates a report of all businesses defined in the database.
func RRreportPeople(t int) string {
	rows, err := RRdb.Prepstmt.GetAllTransactants.Query()
	Errcheck(err)
	defer rows.Close()
	fmt.Printf(" TCID   TID   PID  Name\n")
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
			fmt.Printf("RRreportBusiness: unrecognized print format: %d\n", t)
			return ""
		}
	}
	Errcheck(rows.Err())
	return s
}
