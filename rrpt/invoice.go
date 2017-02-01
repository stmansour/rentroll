package rrpt

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// InvoiceTextReport generates a text invoice for the supplied InvoiceNo
func InvoiceTextReport(id int64) error {
	var noerr error
	funcname := "InvoiceTextReport"
	inv, err := rlib.GetInvoice(id)
	if err != nil {
		e := fmt.Errorf("%s: error getting invoice - %s", funcname, err.Error())
		return e
	}

	var biz rlib.Business
	rlib.GetBusiness(inv.BID, &biz)

	bu, err := rlib.GetBusinessUnitByDesignation(biz.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s", funcname, err.Error())
		return e
	}
	c, err := rlib.GetCompany(int64(bu.CoCode))
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
		rlib.GetTransactant(inv.P[i].PID, &p.Trn)
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
		a, err := rlib.GetAssessment(inv.A[i].ASMID)
		if err != nil {
			fmt.Printf("Error getting Assessment %d:  %s\n", inv.A[i].ASMID, err.Error())
		}
		r := rlib.GetRentable(a.RID)
		fmt.Printf("%-10s  %-12s  %-15s  %-40.40s  %12s  %20s\n", a.Start.Format(rlib.RRDATEFMT3), a.IDtoString(),
			r.Name, rlib.RRdb.BizTypes[biz.BID].GLAccounts[a.ATypeLID].Name, rlib.RRCommaf(a.Amount), a.Comment)
		tot += a.Amount
	}
	fmt.Printf("%s\n", sep)
	fmt.Printf("%-10s  %12s  %15s  %-40s  %12s\n", "Total", " ", " ", " ", rlib.RRCommaf(tot))

	return noerr
}
