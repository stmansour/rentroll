package main

import (
	"fmt"
	"os"
	"rentroll/rlib"
	"strings"
	"time"
)

func initRentRoll() {
	initLists()
	initJFmt()
	initTFmt()
	rlib.RpnInit()
	// loadDefaultCashAccts()
}

func initLists() {
	App.AsmtTypes = rlib.GetAssessmentTypes()
	App.PmtTypes = rlib.GetPaymentTypes()
}

// // Basically this is turned off for now. We'll get to default cash accounts at some point.
// func loadDefaultCashAccts() {
// 	s := "SELECT BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
// 	rows, err := App.dbrr.Query(s)
// 	rlib.Errcheck(err)
// 	defer rows.Close()
// 	for rows.Next() {
// 		var xprop rlib.XBusiness
// 		rlib.Errcheck(rows.Scan(&xprop.P.BID, &xprop.P.Designation,
// 			&xprop.P.Name, &xprop.P.DefaultOccupancyType,
// 			&xprop.P.ParkingPermitInUse, &xprop.P.LastModTime, &xprop.P.LastModBy))
// 	}
// }

func createStartupCtx() DispatchCtx {
	var ctx DispatchCtx
	var err error
	ctx.DtStart, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(App.sStart))
	if err != nil {
		fmt.Printf("Invalid start date:  %s\n", App.sStart)
		os.Exit(1)
	}
	ctx.DtStop, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(App.sStop))
	if err != nil {
		fmt.Printf("Invalid start date:  %s\n", App.sStop)
		os.Exit(1)
	}
	ctx.Report = App.Report
	ctx.Cmd = CmdRUNBOOKS
	return ctx
}
