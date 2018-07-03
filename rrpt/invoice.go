package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strings"
)

// InvoiceTextReport generates a text invoice for the supplied InvoiceNo
func InvoiceTextReport(ctx context.Context, id int64) error {
	const funcname = "InvoiceTextReport"
	var (
		err error
	)

	inv, err := rlib.GetInvoice(ctx, id)
	if err != nil {
		e := fmt.Errorf("%s: error getting invoice - %s", funcname, err.Error())
		return e
	}

	var biz rlib.Business
	err = rlib.GetBusiness(ctx, inv.BID, &biz)
	if err != nil {
		e := fmt.Errorf("%s: error getting business - %s", funcname, err.Error())
		return e
	}

	bu, err := rlib.GetBusinessUnitByDesignation(ctx, biz.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s", funcname, err.Error())
		return e
	}
	c, err := rlib.GetCompany(ctx, int64(bu.CoCode))
	if err != nil {
		e := fmt.Errorf("%s: error getting Company - %s", funcname, err.Error())
		return e
	}

	//---------------------------------------
	//  HEADING
	//---------------------------------------
	fmt.Printf("%s\n", strings.ToUpper(biz.Name))
	fmt.Printf("INVOICE\n")
	fmt.Printf("\n")
	fmt.Printf("%-15s %s\n", "Date:", inv.Dt.Format(rlib.RRDATEFMT3))
	fmt.Printf("%-15s %s\n\n", "Invoice Number:", inv.IDtoString())
	fmt.Printf("%-15s %s\n", "Remit To:", c.LegalName)
	fmt.Printf("%-15s %s\n", " ", c.Address)
	if len(c.Address2) > 0 {
		fmt.Printf("%-15s %s\n", " ", c.Address2)
	}
	fmt.Printf("%-15s %s, %s %s %s\n\n", " ", c.City, c.State, c.PostalCode, c.Country)

	//---------------------------------------
	//  Payors
	//---------------------------------------
	for i := 0; i < len(inv.P); i++ {
		s := " "
		if i == 0 {
			s = "Due From:"
		}
		var p rlib.XPerson
		err = rlib.GetTransactant(ctx, inv.P[i].PID, &p.Trn)
		if err != nil {
			e := fmt.Errorf("%s: Error while getting Transactant - %s", funcname, err.Error())
			return e
		}

		middle := " "
		if len(p.Trn.MiddleName) > 0 {
			middle += p.Trn.MiddleName + " "
		}
		fmt.Printf("%-15s %s%s%s \n", s, p.Trn.FirstName, middle, p.Trn.LastName)
		fmt.Printf("%-15s %s\n", " ", p.Trn.Address)
		if len(p.Trn.Address2) > 0 {
			fmt.Printf("%-15s %s\n", " ", p.Trn.Address2)
		}
		fmt.Printf("%-15s %s, %s %s %s\n\n", " ", p.Trn.City, p.Trn.State, p.Trn.PostalCode, p.Trn.Country)
	}
	fmt.Printf("%-15s %s\n\n", "Delivered By:", inv.DeliveredBy)

	fmt.Printf("%-15s %s\n", "Amount Due:", rlib.RRCommaf(inv.Amount))
	fmt.Printf("%-15s %s\n", "Date Due:", inv.DtDue.Format(rlib.RRDATEFMT3))
	fmt.Printf("\n")

	//---------------------------------------
	//  Assessments
	//---------------------------------------
	fmt.Printf("%-10s  %12s  %-15s  %-40.40s  %12s  %-20s\n", "Date", "AssessmentID", "Rentable", "Description", "Amount", "Comment")
	width := 10 + 12 + 15 + 40 + 12 + 20 + 2*5
	sep := ""
	for i := 0; i < width; i++ {
		sep += "-"
	}
	fmt.Printf("%s\n", sep)
	var tot = float64(0)
	for i := 0; i < len(inv.A); i++ {
		a, err := rlib.GetAssessment(ctx, inv.A[i].ASMID)
		if err != nil {
			fmt.Printf("Error getting Assessment %d:  %s\n", inv.A[i].ASMID, err.Error())
			return err
		}

		r, err := rlib.GetRentable(ctx, a.RID)
		if err != nil {
			return err
		}

		fmt.Printf("%-10s  %-12s  %-15s  %-40.40s  %12s  %20s\n", a.Start.Format(rlib.RRDATEFMT3), a.IDtoString(),
			r.RentableName,
			"", /*rlib.RRdb.BizTypes[biz.BID].GLAccounts[a.ATypeLID].Name*/
			rlib.RRCommaf(a.Amount), a.Comment)
		tot += a.Amount
	}
	fmt.Printf("%s\n", sep)
	fmt.Printf("%-10s  %12s  %15s  %-40s  %12s\n", "Total", " ", " ", " ", rlib.RRCommaf(tot))

	return err
}

// RRreportInvoices generates a report of all rlib.GLAccount accounts
func RRreportInvoices(ctx context.Context, ri *ReporterInfo) string {
	const funcname = "RRreportInvoices"
	var (
		err error
		tbl gotable.Table
	)
	tbl.Init()
	err = TableReportHeaderBlock(ctx, &tbl, "Invoices", funcname, ri)
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl.String()
	}

	tbl.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("InvoiceNo", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("BID", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Due Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Amount", 10, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("DeliveredBy", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	m, err := rlib.GetAllInvoicesInRange(ctx, ri.Bid, &ri.D1, &ri.D2)
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl.String()
	}

	for i := 0; i < len(m); i++ {
		tbl.AddRow()
		tbl.Putd(-1, 0, m[i].Dt)
		tbl.Puts(-1, 1, m[i].IDtoString())
		tbl.Puts(-1, 2, rlib.IDtoString("B", m[i].BID))
		tbl.Putd(-1, 3, m[i].DtDue)
		tbl.Putf(-1, 4, m[i].Amount)
		tbl.Puts(-1, 5, m[i].DeliveredBy)
	}

	tbl.TightenColumns()
	return ReportToString(&tbl, ri)
}
