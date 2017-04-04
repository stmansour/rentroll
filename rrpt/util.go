package rrpt

import (
	"bytes"
	"fmt"
	"gotable"
	"io"
	"rentroll/rlib"
	"time"
)

// RReportTableErrorSectionCSS holds css for errors placed in section3 of gotable
var RReportTableErrorSectionCSS = []*gotable.CSSProperty{
	{Name: "color", Value: "red"},
	{Name: "font-family", Value: "monospace"},
}

const (
	// NoRecordsFoundMsg message to show when there are no results found from db
	NoRecordsFoundMsg = "no records found"
)

// SingleTableReportHandler : single table report handler, used to get report from a table in a required output format
type SingleTableReportHandler struct {
	Found        bool
	ReportNames  []string
	TableHandler func(*ReporterInfo) gotable.Table
}

// MultiTableReportHandler : multi table report handler, used to get report from multiple tables in a required output format
type MultiTableReportHandler struct {
	Found        bool
	ReportNames  []string
	TableHandler func(*ReporterInfo) []gotable.Table
}

// ReporterInfo is for routines that want to table-ize their reporting using
// the CSV library's simple report routines.
type ReporterInfo struct {
	ReportNo              int       // index number of the report
	OutputFormat          int       // text, html, maybe more in the future
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
	Handler               func(*ReporterInfo) string
	Xbiz                  *rlib.XBusiness // may not be set in all cases
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
func TableReportHeader(tbl *gotable.Table, rn, funcname string, ri *ReporterInfo) error {
	tbl.SetTitle(ri.Xbiz.P.Designation + " " + rn)

	var s string
	if ri.RptHeaderD1 && ri.RptHeaderD2 {
		s = ri.D1.Format(rlib.RRDATEREPORTFMT) + " - " + ri.D2.Format(rlib.RRDATEREPORTFMT)
	} else if ri.RptHeaderD1 {
		s = ri.D1.Format(rlib.RRDATEREPORTFMT)
	} else if ri.RptHeaderD2 {
		s = ri.D2.Format(rlib.RRDATEREPORTFMT)
	}
	tbl.SetSection1(s)

	var s1 string
	bu, err := rlib.GetBusinessUnitByDesignation(ri.Xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s", funcname, err.Error())
		tbl.SetSection3(e.Error())
		return e
	}
	if bu.CoCode == 0 {
		s1 = bu.Name + "\n\n"
	} else {
		c, err := rlib.GetCompany(int64(bu.CoCode))
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
func TableReportHeaderBlock(tbl *gotable.Table, rn, funcname string, ri *ReporterInfo) error {
	if ri.Xbiz == nil {
		ri.Xbiz = new(rlib.XBusiness)
		rlib.GetXBusiness(ri.Bid, ri.Xbiz)
	}
	return TableReportHeader(tbl, rn, funcname, ri)
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
func ReportHeader(rn, funcname string, ri *ReporterInfo) (string, error) {
	s := ri.Xbiz.P.Designation + "\n"
	bu, err := rlib.GetBusinessUnitByDesignation(ri.Xbiz.P.Designation)
	if err != nil {
		e := fmt.Errorf("%s: error getting BusinessUnit - %s", funcname, err.Error())
		return s, e
	}
	if bu.CoCode == 0 {
		s += bu.Name + "\n\n"
	} else {
		c, err := rlib.GetCompany(int64(bu.CoCode))
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
	if ri.RptHeaderD1 && ri.RptHeaderD2 {
		s += ri.D1.Format(rlib.RRDATEREPORTFMT) + " - " + ri.D2.Format(rlib.RRDATEREPORTFMT) + "\n"
	} else if ri.RptHeaderD1 {
		s += ri.D1.Format(rlib.RRDATEREPORTFMT) + "\n"
	} else if ri.RptHeaderD2 {
		s += ri.D2.Format(rlib.RRDATEREPORTFMT) + "\n"
	}
	s += "\n"
	return s, nil
}

// ReportHeaderBlock is a wrapper for Report header. It ensures that ri.Xbiz is valid
//		and will append any error messages to the title.
//
// @params
//		  t = table containing the report
//	funcname = name of calling routine in case of error
//        ri = reporter info struct, please ensure RptHeaderD1 and RptHeaderD2 are set correctly
//
// @return
//		string = title string
func ReportHeaderBlock(rn, funcname string, ri *ReporterInfo) string {
	if ri.Xbiz == nil {
		ri.Xbiz = new(rlib.XBusiness)
		rlib.GetXBusiness(ri.Bid, ri.Xbiz)
	}
	s, err := ReportHeader(rn, funcname, ri)
	if err != nil {
		s += "\n" + err.Error() + "\n"
	}
	return s
}

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

// MultiTableTextPrint writes text output from each table to w io.Writer
func MultiTableTextPrint(m []gotable.Table, w io.Writer) {
	funcname := "MultiTableTextPrint"

	for i := 0; i < len(m); i++ {
		temp := bytes.Buffer{}
		err := m[i].TextprintTable(&temp)
		if err != nil {
			s := fmt.Sprintf("Error at %s in t.TextprintTable: %s\n", funcname, err.Error())
			fmt.Print(s)
			fmt.Fprintf(w, "%s\n", s)
		}
		temp.WriteByte('\n')
		w.Write(temp.Bytes())
	}
}

// MultiTableCSVPrint writes csv output from each table to w io.Writer
func MultiTableCSVPrint(m []gotable.Table, w io.Writer) {
	funcname := "MultiTableCSVPrint"

	// TODO: how to handle multiple csv writer
	// add one line between reports??

	for i := 0; i < len(m); i++ {
		temp := bytes.Buffer{}
		err := m[i].CSVprintTable(&temp)
		if err != nil {
			s := fmt.Sprintf("Error at %s in t.CSVprintTable: %s\n", funcname, err.Error())
			fmt.Print(s)
			fmt.Fprintf(w, "%s\n", s)
		}
		temp.WriteByte('\n')
		w.Write(temp.Bytes())
	}
}

// MultiTableHTMLPrint writes html output from each table to w io.Writer
func MultiTableHTMLPrint(m []gotable.Table, w io.Writer) {
	funcname := "MultiTableHTMLPrint"

	for i := 0; i < len(m); i++ {

		// set custom template for reports
		if i == 0 {
			// set first table layout template
			m[i].SetHTMLTemplate("./html/firsttable.html")
		} else if i == len(m)-1 {
			// set last table layout template
			m[i].SetHTMLTemplate("./html/lasttable.html")
		} else {
			// set middle table layout template
			m[i].SetHTMLTemplate("./html/middletable.html")
		}

		temp := bytes.Buffer{}
		err := m[i].HTMLprintTable(&temp)
		if err != nil {
			s := fmt.Sprintf("Error at %s in t.HTMLprintTable: %s\n", funcname, err.Error())
			fmt.Print(s)
			fmt.Fprintf(w, "%s\n", s)
		}
		w.Write(temp.Bytes())
	}
}

// MultiTablePDFPrint writes pdf output from each table to w io.Writer
func MultiTablePDFPrint(m []gotable.Table, w io.Writer) {
	funcname := "MultiTablePDFPrint"

	// TODO: how to handle multiple pdf writer
	// zip multiple files on one files??

	for i := 0; i < len(m); i++ {
		temp := bytes.Buffer{}
		err := m[i].PDFprintTable(&temp)
		if err != nil {
			s := fmt.Sprintf("Error at %s in t.PDFprintTable: %s\n", funcname, err.Error())
			fmt.Print(s)
			fmt.Fprintf(w, "%s\n", s)
		}
		w.Write(temp.Bytes())
	}
}
