package main

import (
	"fmt"
	"gotable"
	"os"
	"rentroll/rlib"
	"rentroll/ws"
	"strings"
	"text/template"
)

func initRentRoll() {
	// initJFmt()
	// initTFmt()
	rlib.RpnInit()

	RRfuncMap = template.FuncMap{
		"DateToString": rlib.DateToString,
		"GetVersionNo": ws.GetVersionNo,
		"getBuildTime": ws.GetBuildTime,
		"RRCommaf":     rlib.RRCommaf,
		"LMSum":        LMSum,
	}

}

func createStartupCtx() DispatchCtx {
	var ctx DispatchCtx
	var err error
	ctx.DtStart, err = rlib.StringToDate(App.sStart)
	if err != nil {
		fmt.Printf("Invalid start date:  %s\n", App.sStart)
		os.Exit(1)
	}
	ctx.DtStop, err = rlib.StringToDate(App.sStop)
	if err != nil {
		fmt.Printf("Invalid start date:  %s\n", App.sStop)
		os.Exit(1)
	}

	des := strings.ToLower(strings.TrimSpace(App.Bud)) // this should be BUD
	if len(des) == 0 {                                 // make sure it's not empty
		fmt.Printf("No BUD specified. A BUD is required for batch mode operation\n")
		os.Exit(1)
	}
	ctx.xbiz.P = rlib.GetBusinessByDesignation(des) // see if we can find the biz
	if len(ctx.xbiz.P.Designation) == 0 {
		rlib.Ulog("Business Unit with designation %s does not exist\n", des)
		os.Exit(1)
	}
	rlib.GetXBusiness(ctx.xbiz.P.BID, &ctx.xbiz)

	// App.Report is a string, of the format:
	//   n[,s1[,s2[...]]]
	// Example:
	//   1           -- this would be for a Journal text report
	//   9,IN0001    -- this would be for a text report of Invoice 1
	//
	// The only required value is n, the report number
	sa := strings.Split(App.Report, ",") // comma separated list
	if len(App.Report) > 0 {
		ctx.Report, _ = rlib.IntFromString(sa[0], "invalid report number")
	}
	ctx.Args = App.Report
	ctx.CSVLoadStr = strings.TrimSpace(App.CSVLoad)
	// fmt.Printf("ctx.CSVLoadStr = %s\n", ctx.CSVLoadStr)
	ctx.Cmd = 1
	ctx.OutputFormat = gotable.TABLEOUTTEXT
	return ctx
}
