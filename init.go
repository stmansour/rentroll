package main

import (
	"fmt"
	"os"
	"rentroll/rlib"
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
		"RRCommaf":     RRCommaf,
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
	ctx.Report = App.Report
	ctx.Cmd = CmdRUNBOOKS
	ctx.OutputFormat = FMTTEXT
	return ctx
}
