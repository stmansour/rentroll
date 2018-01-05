package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"io"
	"net/url"
	"rentroll/rlib"
	"strings"
	"time"
)

// SetPDFOption sets option to pdf properties,
// if already exists then overwrites with provided value otherwise append new one
func SetPDFOption(
	pdfProps []*gotable.PDFProperty,
	optionName string,
	optionValue string,
) []*gotable.PDFProperty {

	var (
		found  bool
		newOpt = &gotable.PDFProperty{Option: optionName, Value: optionValue}
	)

	for index, opt := range pdfProps {
		if opt.Option == optionName {
			temp := append(pdfProps[:index], newOpt)
			pdfProps = append(temp, pdfProps[index+1:]...)
			found = true
			break
		}
	}

	// if not found in pdf props then make new and append it
	if !found {
		pdfProps = append(pdfProps, newOpt)
	}

	return pdfProps
}

// GetReportPDFProps returns the default pdf properties
// will be used by "wkhtmltopdf" program
func GetReportPDFProps() []*gotable.PDFProperty {
	// report generated time, in footer left area
	rptGenTime := time.Now()
	zone, _ := rptGenTime.Zone()
	rptGenTimeREP := rptGenTime.Format(gotable.DATETIMEFMT) + " " + zone

	// rptPDFProps are the pdf properties values for pdf report for rentroll software
	// pdf properties
	return []*gotable.PDFProperty{
		// disable smart shrinking
		// {Option: "--disable-smart-shrinking"},
		// custom dpi setting
		{Option: "--dpi", Value: "512"},
		// top margin
		{Option: "--margin-top", Value: "15"},
		// header font size
		{Option: "--header-font-size", Value: "7"},
		// header font
		{Option: "--header-font-name", Value: "opensans"},
		// header spacing
		{Option: "--header-spacing", Value: "3"},
		// bottom margin
		{Option: "--margin-bottom", Value: "15"},
		// footer spacing
		{Option: "--footer-spacing", Value: "5"},
		// footer font
		{Option: "--footer-font-name", Value: "opensans"},
		// footer font size
		{Option: "--footer-font-size", Value: "7"},
		// footer left content
		{Option: "--footer-left", Value: rptGenTimeREP},
		// footer right content
		{Option: "--footer-right", Value: "Page [page] of [toPage]"},
		// // page size
		// {Option: "--page-size", Value: "Letter"},
		// // orientation
		// {Option: "--orientation", Value: "Landscape"},
		// page width, defaults to US Letter with LandScape
		{Option: "--page-width", Value: "11in"},
		// page height, defaults to US Letter with LandScape
		{Option: "--page-height", Value: "8.5in"},
	}
}

// RReportTableErrorSectionCSS holds css for errors placed in section3 of gotable
var RReportTableErrorSectionCSS = []*gotable.CSSProperty{
	{Name: "color", Value: "red"},
	{Name: "font-family", Value: "monospace"},
}

const (
	// NoRecordsFoundMsg message to show when there are no results found from db
	NoRecordsFoundMsg = "no records found"
)

type singleGoTableHandler func(context.Context, *ReporterInfo) gotable.Table

// SingleTableReportHandler : single table report handler, used to get report from a table in a required output format
type SingleTableReportHandler struct {
	Found        bool
	ReportNames  []string
	TableHandler singleGoTableHandler
}

type multiGoTableHandler func(context.Context, *ReporterInfo) ([]gotable.Table, error)

// MultiTableReportHandler : multi table report handler, used to get report from multiple tables in a required output format
type MultiTableReportHandler struct {
	ReportTitle  string
	Found        bool
	ReportNames  []string
	TableHandler multiGoTableHandler
}

// ReporterInfo is for routines that want to table-ize their reporting using
// the CSV library's simple report routines.
type ReporterInfo struct {
	ReportNo              int       // index number of the report
	OutputFormat          int       // text, html, maybe more in the future
	EDI                   int       // end date inclusive -- 0 = no, 1 = yes
	Bid                   int64     // associated business
	Raid                  int64     // associated Rental Agreement if needed
	D1                    time.Time // associated date if needed
	D2                    time.Time // associated date if needed
	NeedsBID              bool      // true if BID is needed for this report
	NeedsRAID             bool      // true if RAID is needed for this report
	NeedsDt               bool      // true if a Date is needed for this report
	RptHeaderD1           bool      // true if the report's header should contain D1
	RptHeaderD2           bool      // true if the dates should show as a range D1 - D2
	BlankLineAfterRptName bool      // true if a blank line should be added after the Report Name
	Handler               func(context.Context, *ReporterInfo) string
	Xbiz                  *rlib.XBusiness // may not be set in all cases
	QueryParams           *url.Values
}

// GetDisplayD2 returns the appropriate display date for the end date of a period.
//     Some people like to see the last date as up-to-and-including.  But from a
//     code perspective, we always set the end date to a value which
//     is up-to-but-not-including.  If ri.EDI == 0, then we display the date as
//     up-to-and-including.  If EDI == 1, we display the up-to-but-not-including
//     value.
//
// INPUT
//   ri - Report context
//
// RETURNS - the display date
//-----------------------------------------------------------------------------
func GetDisplayD2(ri *ReporterInfo) time.Time {
	// Adjust d2 depending on the mode...  0 = no adjustment needed, 1 = inclusive last date, so subtract 1 day
	// rlib.Console("*** GetDisplayD2:  ri.EDI = %d\n", ri.EDI)
	d2 := ri.D2
	if ri.EDI == 1 {
		d2 = d2.AddDate(0, 0, -1)
	}
	return d2
}

// TableReportHeader returns a title block of text for a report. The format is:
//
// 			Title:     <BUD> <Report Name>
//			Section1:  <date or dateRange>
//			Section2:  <Business name and address>
//
// @params
//	tbl  	      = table containing the report
//	rn	      = Report Name
//	funcname = name of calling routine in case of error
//	ri	      = reporter info struct, please ensure RptHeaderD1 and RptHeaderD2 are set correctly
//
// @return
//		string = title string
//         err = any problem that occurred
func TableReportHeader(ctx context.Context, tbl *gotable.Table, rn, funcname string, ri *ReporterInfo) error {
	tbl.SetTitle(ri.Xbiz.P.Designation + " " + rn)

	var s string
	d2 := GetDisplayD2(ri)
	// rlib.Console("*** Table Report Header:  d2 = %s\n", d2.Format(rlib.RRDATEREPORTFMT))
	if ri.RptHeaderD1 && ri.RptHeaderD2 {
		s = ri.D1.Format(rlib.RRDATEREPORTFMT) + " - " + d2.Format(rlib.RRDATEREPORTFMT)
	} else if ri.RptHeaderD1 {
		s = ri.D1.Format(rlib.RRDATEREPORTFMT)
	} else if ri.RptHeaderD2 {
		s = d2.Format(rlib.RRDATEREPORTFMT)
	}
	tbl.SetSection1(s)

	var s1 string
	bu, err := rlib.GetBusinessUnitByDesignation(ctx, ri.Xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s", funcname, err.Error())
		tbl.SetSection3(e.Error())
		return e
	}
	if bu.CoCode == 0 {
		s1 = bu.Name + "\n\n"
	} else {
		c, err := rlib.GetCompany(ctx, int64(bu.CoCode))
		if err != nil {
			e := fmt.Errorf("%s: error getting Company - %s\nBusinessUnit = %s, bu = %#v", funcname, err.Error(), ri.Xbiz.P.Designation, bu)
			tbl.SetSection3(e.Error())
			return e
		}
		s1 += fmt.Sprintf("%s\n", c.LegalName)
		s1 += fmt.Sprintf("%s\n", c.Address)
		if len(c.Address2) > 0 {
			s1 += fmt.Sprintf("%s\n", c.Address2)
		}
		s1 += fmt.Sprintf("%s, %s %s %s\n\n", c.City, c.State, c.PostalCode, c.Country)
	}
	// TODO: handle blank line thing for html???
	if ri.BlankLineAfterRptName {
		s1 += "\n"
	}
	tbl.SetSection2(s1)

	return nil
}

// TableReportHeaderBlock is a wrapper for Report header. It ensures that ri.Xbiz is valid
//		and will append any error messages to the title.
//
// @params
//		  tbl = table containing the report
//	funcname = name of calling routine in case of error
//        ri = reporter info struct, please ensure RptHeaderD1 and RptHeaderD2 are set correctly
//
// @return
//		string = title string
func TableReportHeaderBlock(ctx context.Context, tbl *gotable.Table, rn, funcname string, ri *ReporterInfo) error {
	if ri.Xbiz == nil {
		ri.Xbiz = new(rlib.XBusiness)
		if err := rlib.GetXBusiness(ctx, ri.Bid, ri.Xbiz); err != nil {
			return err
		}
	}
	return TableReportHeader(ctx, tbl, rn, funcname, ri)
}

// ReportHeader returns a title block of text for a report.
// @params
//		  rn = Report Name
//	funcname = name of calling routine in case of error
//        ri = reporter info struct, please ensure RptHeaderD1 and RptHeaderD2 are set correctly
//
// @return
//		string = title string
//         err = any problem that occurred
func ReportHeader(ctx context.Context, rn, funcname string, ri *ReporterInfo) (string, error) {
	s := ri.Xbiz.P.Designation + "\n"
	bu, err := rlib.GetBusinessUnitByDesignation(ctx, ri.Xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s", funcname, err.Error())
		return s, e
	}
	if bu.CoCode == 0 {
		s += bu.Name + "\n\n"
	} else {
		c, err := rlib.GetCompany(ctx, int64(bu.CoCode))
		if err != nil {
			e := fmt.Errorf("%s: error getting Company - %s\nBusinessUnit = %s, bu = %#v", funcname, err.Error(), ri.Xbiz.P.Designation, bu)
			return s, e
		}
		s += fmt.Sprintf("%s\n", c.LegalName)
		s += fmt.Sprintf("%s\n", c.Address)
		if len(c.Address2) > 0 {
			s += fmt.Sprintf("%s\n", c.Address2)
		}
		s += fmt.Sprintf("%s, %s %s %s\n\n", c.City, c.State, c.PostalCode, c.Country)
	}
	s += rn
	if ri.BlankLineAfterRptName {
		s += "\n"
	}

	d2 := GetDisplayD2(ri)
	rlib.Console("*** Report Header:  d2 = %s\n", d2.Format(rlib.RRDATEREPORTFMT))

	if ri.RptHeaderD1 && ri.RptHeaderD2 {
		s += ri.D1.Format(rlib.RRDATEREPORTFMT) + " - " + d2.Format(rlib.RRDATEREPORTFMT) + "\n"
	} else if ri.RptHeaderD1 {
		s += ri.D1.Format(rlib.RRDATEREPORTFMT) + "\n"
	} else if ri.RptHeaderD2 {
		s += d2.Format(rlib.RRDATEREPORTFMT) + "\n"
	}
	s += "\n"
	return s, nil
}

/*// ReportHeaderBlock is a wrapper for Report header. It ensures that ri.Xbiz is valid
//		and will append any error messages to the title.
//
// @params
//		  t = table containing the report
//	funcname = name of calling routine in case of error
//        ri = reporter info struct, please ensure RptHeaderD1 and RptHeaderD2 are set correctly
//
// @return
//		string = title string
func ReportHeaderBlock(ctx context.Context, rn, funcname string, ri *ReporterInfo) string {
	if ri.Xbiz == nil {
		ri.Xbiz = new(rlib.XBusiness)
		err = rlib.GetXBusiness(ctx, ri.Bid, ri.Xbiz)
		if err != nil {
			return err.Error() // should return error from here!!!
		}
	}
	s, err := ReportHeader(rn, funcname, ri)
	if err != nil {
		s += "\n" + err.Error() + "\n"
	}
	return s
}*/

// ReportToString returns a string version of the report. It uses information in
// 		ri for the output format and whether or not to include the title.
//
// @params
//		  t = table containing the report
//        ri = reporter info struct, please ensure RptHeaderD1 and RptHeaderD2 are set correctly
//
// @return
//		string version of the report
func ReportToString(t *gotable.Table, ri *ReporterInfo) string {
	s, err := t.SprintTable()
	if nil != err {
		s += err.Error()
		rlib.Ulog("ReportToString: error = %s", err)
	}
	return s
}

// getRRTable returns a table with some basic initialization
// to be used in all reports of rentroll software
func getRRTable() gotable.Table {
	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)
	tbl.SetNoRowsCSS(RReportTableErrorSectionCSS)
	tbl.SetNoHeadersCSS(RReportTableErrorSectionCSS)

	return tbl
}

// MultiTablePDFPrint writes pdf output from each table to w io.Writer
func MultiTablePDFPrint(m []gotable.Table, w io.Writer, pdfTitle string, pdfPageWidth float64, pdfPageHeight float64, pdfPageSizeUnit string) {

	// pdf props title
	pdfProps := GetReportPDFProps()

	pdfProps = SetPDFOption(pdfProps, "--header-center", pdfTitle)
	pw := rlib.Float64ToString(pdfPageWidth) + pdfPageSizeUnit
	pdfProps = SetPDFOption(pdfProps, "--page-width", pw)
	ph := rlib.Float64ToString(pdfPageHeight) + pdfPageSizeUnit
	pdfProps = SetPDFOption(pdfProps, "--page-height", ph)

	gotable.MultiTablePDFPrint(m, w, pdfProps)

}

// GetAttachmentDate used to get date for attachements served over web
func GetAttachmentDate(t time.Time) string {
	y, m, d := t.Date()
	year := fmt.Sprintf("%04d", y)
	month := strings.ToUpper(m.String()[:3])
	date := fmt.Sprintf("%02d", d)
	formatDate := year + "-" + month + "-" + date
	return formatDate
}
