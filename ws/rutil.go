package ws

import (
	"html/template"
	"rentroll/rlib"
	"time"
)

// RRfuncMap is a map of functions passed to each html page that can be referenced
// as needed to produce the page
var RRfuncMap map[string]interface{}

// ReportContext is a structure of data that will be passed to all html pages.
// It is the responsibility of the page function to populate the data needed by
// the page. The recommendation is to populate only the data needed.
type ReportContext struct {
	Language           string          // what language
	Template           string          // which template
	DtStart            string          // start of period of interest
	D1                 time.Time       // time.Time value for DtStart
	DtStop             string          // end of period of interest
	D2                 time.Time       // time.Time value for DtStop
	ID                 int64           // ID for reports that detail a specific entity
	B                  rlib.Business   // business associated with this report
	BL                 []rlib.Business // array of all businesses, for initializing dropdown selections
	ReportContent      string          // text report content
	PageTitle          string          // set page title via software
	ReportOutputFormat int             // indicates text, html, or pdf
	PDFPageWidth       float64         // page width
	PDFPageHeight      float64         // page height
	PDFPageSizeUnit    string          // page size unit, default is inch ("in")
	EDI                int             // end date inclusive - 0 = end date is not inclusive, 1 = end date is inclusive
	// LDG                UILedger        // ledgers associated with this report
}

// InitReports initializes the reports subsystem. Historically it
// did more than it does today. Currently, it initializes the map
// of functions that can be used by a startup web page such as home.html
// or rhome.html
//-----------------------------------------------------------------------------
func InitReports() {
	RRfuncMap = template.FuncMap{
		"DateToString": rlib.DateToString,
		"GetVersionNo": rlib.GetVersionNo,
		"getBuildTime": rlib.GetBuildTime,
		"RRCommaf":     rlib.RRCommaf,
	}
}
