package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"time"
)

// ReceiptPDFProps holds the override properties needed for the receipt report
var ReceiptPDFProps = []*gotable.PDFProperty{
	//
	{Option: "--no-collate"},
	// top margin
	{Option: "-T", Value: "15"},
	// header center content
	// {Option: "--header-center", Value: "Header Center"},
	// header font size
	{Option: "--header-font-size", Value: "7"},
	// header font
	{Option: "--header-font-name", Value: "opensans"},
	// header spacing
	{Option: "--header-spacing", Value: "3"},
	// bottom margin
	{Option: "-B", Value: "15"},
	// footer spacing
	{Option: "--footer-spacing", Value: "5"},
	// footer font
	{Option: "--footer-font-name", Value: "opensans"},
	// footer font size
	{Option: "--footer-font-size", Value: "7"},
	// footer left content
	{Option: "--footer-left", Value: time.Now().Format(gotable.DATETIMEFMT)},
	// footer right content
	{Option: "--footer-right", Value: "Page [page] of [toPage]"},
	// page size
	{Option: "--page-size", Value: "A5"},
	// orientation
	{Option: "--orientation", Value: "Portrait"},
	// sizing
	{Option: "--dpi", Value: "1600"},
}

// RRRcptOnlyReceiptTable generates a receipt suitable for printing
//
// INPUT
//  ctx    - context containing session, existing db transactions, etc.
//  ri     - report information
//
// RETURNS
//  the gotable
//-----------------------------------------------------------------------------
func RRRcptOnlyReceiptTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	ri.Style = 0 // ensure permanent resident receipt
	return rcptOnlyReceiptTable(ctx, ri)
}

// RRRcptHotelReceiptTable generates the hotel receipt
func RRRcptHotelReceiptTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	ri.Style = 1 // ensure hotel receipt
	return rcptOnlyReceiptTable(ctx, ri)
}

// rcptOnlyReceiptTable
//  style  - 0 - Permanent resident hotel
//           1 - hotel receipt
func rcptOnlyReceiptTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "RRRcptOnlyReceiptTable"
	var err error

	tbl := getRRTable()
	tbl.AddColumn("", 155, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("", 75, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	switch ri.Style {
	case 0:
		tbl.SetTitle("Receipt of Delivery")
	case 1:
		tbl.SetTitle("Receipt for Payment")
	default:
		tbl.SetTitle("Receipt")
	}

	// get records from db
	m, err := rlib.GetReceipt(ctx, ri.ID)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	//----------------------------------------------
	// Load the receipt types for this business...
	//----------------------------------------------
	sPmtType := fmt.Sprintf("PMTID = %d ??", m.PMTID)

	n, _ := rlib.GetPaymentTypesByBusiness(ctx, m.BID) // get the payment types for this business
	sPmtType = n[m.PMTID].Name

	rname, _ := rlib.ROCExtractRentableName(m.Comment)

	tbl.AddRow()
	tbl.AddRow()
	tbl.Puts(-1, 0, "Receipt Number")
	tbl.Puts(-1, 1, rlib.IDtoShortString("RCPT", m.RCPTID))

	tbl.AddRow()
	tbl.Puts(-1, 0, "Date")
	tbl.Puts(-1, 1, m.Dt.Format(rlib.RRDATERECEIPTFMT))

	// if the receipt has been reversed, set up some extra info
	if m.FLAGS&rlib.RCPTREVERSED != 0 {
		tbl.AddRow()
		tbl.Puts(-1, 0, "")
		if m.PRCPTID > 0 {
			tbl.Puts(-1, 1, fmt.Sprintf("*** THIS RECEIPT IS VOID *** reverses %s", rlib.IDtoShortString("RCPT", m.PRCPTID)))
		} else {
			tbl.Puts(-1, 1, "*** THIS RECEIPT IS VOID ***")
		}
	}

	tbl.AddRow()
	tbl.Puts(-1, 0, "Received From")
	tbl.Puts(-1, 1, m.OtherPayorName)

	tbl.AddRow()
	tbl.Puts(-1, 0, "Resident Address")
	tbl.Puts(-1, 1, rname)

	tbl.AddRow()
	tbl.Puts(-1, 0, "Received By")
	tbl.Puts(-1, 1, rlib.GetNameForUID(ctx, m.CreateBy))

	tbl.AddRow()
	tbl.Puts(-1, 0, "Amount")
	tbl.Puts(-1, 1, "$"+rlib.RRCommaf(m.Amount))

	tbl.AddRow()
	tbl.Puts(-1, 0, "Form of Payment")
	tbl.Puts(-1, 1, sPmtType)

	tbl.AddRow()
	tbl.Puts(-1, 0, "Reference")
	tbl.Puts(-1, 1, m.DocNo)

	if ri.Style == 0 {
		tbl.AddRow()
		tbl.AddRow()
		tbl.AddRow()
		tbl.Puts(-1, 0, "Email address")
		tbl.Puts(-1, 1, "____________________________________________________")

		tbl.AddRow()
		tbl.AddRow()
		tbl.AddLineAfter(len(tbl.Row) - 1)
	}

	tbl.TightenColumns()

	return tbl
}

// RRRcptOnlyReceipt generates a report
func RRRcptOnlyReceipt(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRRcptOnlyReceiptTable(ctx, ri)
	return ReportToString(&tbl, ri)
}

// RRRcptOnlyHotelReceipt generates a report
func RRRcptOnlyHotelReceipt(ctx context.Context, ri *ReporterInfo) string {
	tbl := RRRcptHotelReceiptTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
