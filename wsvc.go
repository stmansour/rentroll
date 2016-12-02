package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"rentroll/rlib"
	"time"
)

// webServiceHandler dispatches all the web service requests
func webServiceHandler(w http.ResponseWriter, r *http.Request) {
	funcname := "webServiceHandler"
	fmt.Printf("Entered %s\n", funcname)
	var ui RRuiSupport
	var xbiz rlib.XBusiness

	path := "/wsvc/"              // this is the part of the URL that got us into this handler
	u := r.RequestURI[len(path):] // this pulls off the specific request

	cmdinfo, err := url.QueryUnescape(u)
	if err != nil {
		e := fmt.Errorf("%s: Error with QueryUnescape: %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	fmt.Printf("cmdinfo = %s\n", cmdinfo)

	rlib.GetXBusiness(1, &xbiz)
	ui.D1 = time.Date(2016, time.August, 1, 0, 0, 0, 0, time.UTC)
	ui.D2 = time.Date(2016, time.September, 1, 0, 0, 0, 0, time.UTC)
	ui.ReportContent = csvloadReporter(cmdinfo, &xbiz, &ui)

	tmpl := "gsvcrpt.html"
	t, err := template.New(tmpl).Funcs(RRfuncMap).ParseFiles("./html/" + tmpl)
	if nil != err {
		s := fmt.Sprintf("%s: error loading template: %v\n", funcname, err)
		ui.ReportContent += s
		fmt.Print(s)
	}
	err = t.Execute(w, &ui)

	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
