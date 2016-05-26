package main

import (
	"fmt"
	"os"
	"rentroll/rlib"
	"strings"
	"text/template"
	"time"
)

func initRentRoll() {
	initLists()
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

func initLists() {
	App.AsmtTypes = rlib.GetAssessmentTypes()
	App.PmtTypes = rlib.GetPaymentTypes()
}

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
	ctx.OutputFormat = FMTTEXT
	return ctx
}
