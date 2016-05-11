package rlib

import "fmt"

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
