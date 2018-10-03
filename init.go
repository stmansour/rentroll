package main

import (
	"fmt"
	"gotable"
	"os"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"rentroll/ws"
	"strings"
)

func initRentRoll() {
	rlib.RpnInit()
	bizlogic.InitBizLogic()
	ws.InitReports()
}

func createStartupCtx() DispatchCtx {
	var (
		dCtx DispatchCtx
		err  error
	)

	dCtx.DtStart, err = rlib.StringToDate(App.sStart)
	if err != nil {
		fmt.Printf("Invalid start date:  %s\n", App.sStart)
		os.Exit(1)
	}

	dCtx.DtStop, err = rlib.StringToDate(App.sStop)
	if err != nil {
		fmt.Printf("Invalid stop date:  %s\n", App.sStop)
		os.Exit(1)
	}

	des := strings.ToLower(strings.TrimSpace(App.Bud)) // this should be BUD
	if len(des) == 0 {                                 // make sure it's not empty
		fmt.Printf("No BUD specified. A BUD is required for batch mode operation\n")
		os.Exit(1)
	}
	dCtx.xbiz.P, err = rlib.GetBizByDesignation(des) // see if we can find the biz
	if err != nil /*len(dCtx.xbiz.P.Designation) == 0*/ {
		rlib.Ulog("Business Unit with designation %s does not exist: error: %s\n", des, err.Error())
		os.Exit(1)
	}
	rlib.GetXBiz(dCtx.xbiz.P.BID, &dCtx.xbiz)

	// if dateMode is on then change the stopDate value for search op
	rlib.EDIHandleIncomingDateRange(dCtx.xbiz.P.BID, &dCtx.DtStart, &dCtx.DtStop)

	// App.Report is a string, of the format:
	//   n[,s1[,s2[...]]]
	// Example:
	//   1           -- this would be for a Journal text report
	//   9,IN0001    -- this would be for a text report of Invoice 1
	//
	// The only required value is n, the report number
	sa := strings.Split(App.Report, ",") // comma separated list
	if len(App.Report) > 0 {
		dCtx.Report, _ = rlib.IntFromString(sa[0], "invalid report number")
	}
	dCtx.Args = App.Report
	dCtx.CSVLoadStr = strings.TrimSpace(App.CSVLoad)
	// fmt.Printf("dCtx.CSVLoadStr = %s\n", dCtx.CSVLoadStr)
	dCtx.Cmd = 1
	dCtx.OutputFormat = gotable.TABLEOUTTEXT
	return dCtx
}
