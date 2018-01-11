package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// RRreportSpecialties generates a report of all RentalSpecialties
func RRreportSpecialties(ctx context.Context, ri *ReporterInfo) string {
	var (
		xbiz rlib.XBusiness
		err  error
	)
	err = rlib.GetXBusiness(ctx, ri.Bid, &xbiz) // get its info
	if err != nil {
		return err.Error()
	}

	s := fmt.Sprintf("%-11s  %-9s  %-30s  %10s  %-15s\n", "RSPID", "BID", "Name", "Fee", "Description")

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
func RRreportSpecialtyAssigns(ctx context.Context, ri *ReporterInfo) string {
	var (
		xbiz rlib.XBusiness
		err  error
	)
	err = rlib.GetXBusiness(ctx, ri.Bid, &xbiz) // get its info
	if err != nil {
		return err.Error()
	}

	s := fmt.Sprintf("%9s  %9s  %-30s  %10s  %10s\n", "BID", "RID", "Specialty Name", "DtStart", "DtStop")
	rows, err := rlib.RRdb.Prepstmt.GetAllRentableSpecialtyRefs.Query(ri.Bid)
	if err != nil {
		return err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var a rlib.RentableSpecialtyRef
		err = rows.Scan(&a.BID, &a.RID, &a.RSPID, &a.DtStart, &a.DtStop, &a.LastModTime, &a.LastModBy)
		if err != nil {
			return err.Error()
		}

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
	err = rows.Err()
	if err != nil {
		return err.Error()
	}
	return s
}
