package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
	"time"
)

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
	RptHeader             bool      // true if header should be printed as part of the report
	BlankLineAfterRptName bool      // true if a blank line should be added after the Report Name
	Handler               func(*ReporterInfo) string
	Xbiz                  *rlib.XBusiness // may not be set in all cases
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
	if ri.RptHeader {
		s, err := t.SprintTable()
		if nil != err {
			rlib.Ulog("ReportToString: error = %s", err)
		}
		return t.GetTitle() + s
	}
	s, err := t.SprintTable()
	if nil != err {
		rlib.Ulog("ReportToString: error = %s", err)
	}
	return s
}
