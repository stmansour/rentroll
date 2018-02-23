package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// RRreportRatePlans generates a report of all RateLists for the supplied business (ri.Bid)
func RRreportRatePlans(ctx context.Context, ri *ReporterInfo) string {
	m, err := rlib.GetAllRatePlans(ctx, ri.Bid)
	if err != nil {
		return err.Error()
	}

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
func RRreportRatePlanRefs(ctx context.Context, ri *ReporterInfo) string {
	const funcname = "RRreportRatePlanRefs"
	var (
		err  error
		rp   rlib.RatePlan
		xbiz rlib.XBusiness
	)
	err = rlib.GetXBusiness(ctx, ri.Bid, &xbiz)
	if err != nil {
		return err.Error()
	}

	m, err := rlib.GetAllRatePlanRefsInRange(ctx, &ri.D1, &ri.D2)
	if err != nil {
		return err.Error()
	}

	if len(m) == 0 {
		fmt.Printf("%s: could not load RatePlanRefs for timerange %s - %s\n", funcname, ri.D1.Format(rlib.RRDATEFMT4), ri.D2.Format(rlib.RRDATEFMT4))
		return ""
	}

	s := fmt.Sprintf("%-15s  %-11s  %-10s  %-10s  %-8s  %-6s  %-9s  %-9s  %-20s\n", "RatePlan", "RPRID", "DtStart", "DtStop", "MaxNoFee", "FeeAge", "Fee", "CancelFee", "PromoCode")
	s += fmt.Sprintf("%-15s  %-11s  %-10s  %-10s  %-8s  %-6s  %-9s  %-9s  %-20s\n", "--------", "-----", "----------", "----------", "--------", "------", "---------", "---------", "---------")

	for i := 0; i < len(m); i++ {
		p := m[i]
		err = rlib.GetRatePlan(ctx, p.RPID, &rp)
		if err != nil {
			return err.Error()
		}
		err = rlib.GetRatePlanRefFull(ctx, p.RPRID, &p)
		if err != nil {
			return err.Error()
		}

		// just before printing out report, modify end date mode if applicable
		rlib.HandleInterfaceEDI(&p, ri.Bid)

		switch ri.OutputFormat {
		case gotable.TABLEOUTTEXT:
			s += fmt.Sprintf("%-15.15s  RPR%08d  %10s  %10s  %8d  %6d  %9.2f  %9.2f  %s\n",
				rp.Name, p.RPRID, p.DtStart.Format(rlib.RRDATEFMT4), p.DtStop.Format(rlib.RRDATEFMT4),
				p.MaxNoFeeUsers, p.FeeAppliesAge, p.AdditionalUserFee, p.CancellationFee, p.PromoCode)
			s += RRreportRatePlanRefRTRates(ctx, &p, &xbiz)
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
func RRreportRatePlanRefRTRates(ctx context.Context, rpr *rlib.RatePlanRef, xbiz *rlib.XBusiness) string {
	var (
	// err error
	)

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
		m, err := rlib.GetAllRatePlanRefSPRates(ctx, p.RPRID, p.RTID) // almost certainly not in the order we want them
		if err != nil {
			return err.Error()
		}

		for j := 0; j < len(m)-1; j++ { // we order them just to be sure
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
