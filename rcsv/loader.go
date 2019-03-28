package rcsv

import (
	"rentroll/rlib"
	"time"
)

// Rcsv contains the shared data used by the RCS loaders
var Rcsv struct {
	DtStart time.Time
	DtStop  time.Time
	Xbiz    *rlib.XBusiness
}

// InitRCSV initializes the shared data used by they RCS loaders.
func InitRCSV(d1, d2 *time.Time, xbiz *rlib.XBusiness) {
	Rcsv.DtStart = *d1
	Rcsv.DtStop = *d2
	Rcsv.Xbiz = xbiz

	// if dateMode is on then change the stopDate value for search op
	rlib.EDIHandleIncomingDateRange(Rcsv.Xbiz.P.BID, &Rcsv.DtStart, &Rcsv.DtStop)
}
