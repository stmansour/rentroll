package main

import (
	"fmt"
	"rentroll/rlib"
)

// RunBooks runs a series of commands to handle command line run requests
func RunBooks(ctx *DispatchCtx) {
	s := "SELECT BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() { // For every business
		var xbiz rlib.XBusiness
		rlib.Errcheck(rows.Scan(&xbiz.P.BID, &xbiz.P.Designation, &xbiz.P.Name, &xbiz.P.DefaultOccupancyType, &xbiz.P.ParkingPermitInUse, &xbiz.P.LastModTime, &xbiz.P.LastModBy))
		rlib.GetXBusiness(xbiz.P.BID, &xbiz) // get its info
		rlib.InitBusinessFields(xbiz.P.BID)
		rlib.GetDefaultLedgerMarkers(xbiz.P.BID) // Gather its chart of accounts

		// and generate the requested report...
		switch ctx.Report {
		case 1:
			JournalReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 2:
			LedgerReportText(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 3:
			intTest(&xbiz, &ctx.DtStart, &ctx.DtStop)
		case 4:
			fmt.Printf("biz csv = %s\n", App.bizfile)
		default:
			GenerateJournalRecords(&xbiz, &ctx.DtStart, &ctx.DtStop)
			GenerateLedgerRecords(&xbiz, &ctx.DtStart, &ctx.DtStop)
		}
	}
}

// Dispatch generates the supplied report for all properties
func Dispatch(ctx *DispatchCtx) {
	switch ctx.Cmd {
	case CmdRUNBOOKS:
		RunBooks(ctx)
	default:
		fmt.Printf("Unrecognized command: %d\n", ctx.Cmd)
	}
}
