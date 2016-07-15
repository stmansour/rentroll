package main

import (
	"fmt"
	"os"
	"rentroll/rlib"
	"strings"
	"text/template"
)

func initRentRoll() {
	initJFmt()
	initTFmt()
	rlib.RpnInit()

	RRfuncMap = template.FuncMap{
		"DateToString": rlib.DateToString,
		"getVersionNo": getVersionNo,
		"getBuildTime": getBuildTime,
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
	ctx.Cmd = CmdRUNBOOKS
	ctx.OutputFormat = FMTTEXT
	return ctx
}
