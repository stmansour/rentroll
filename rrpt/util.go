package rrpt

import (
	"rentroll/rlib"
	"time"
)

// ReporterInfo is for routines that want to table-ize their reporting using
// the CSV library's simple report routines.
type ReporterInfo struct {
	ReportNo     int       // index number of the report
	OutputFormat int       // text, html, maybe more in the future
	Bid          int64     // associated business
	Raid         int64     // associated Rental Agreement if needed
	D1           time.Time // associated date if needed
	D2           time.Time // associated date if needed
	NeedsBID     bool      // true if BID is needed for this report
	NeedsRAID    bool      // true if RAID is needed for this report
	NeedsDt      bool      // true if a Date is needed for this report
	RptHeaderD1  bool      // true if the report's header should contain D1
	RptHeaderD2  bool      // true if the dates should show as a range D1 - D2
	Handler      func(*ReporterInfo) string
	Xbiz         *rlib.XBusiness // may not be set in all cases
}
