package main

import (
	"fmt"
	"html/template"
	"net/http"
	"rentroll/rlib"
)

// HomeUIHandler sends the main UI to the browser
func HomeUIHandler(w http.ResponseWriter, r *http.Request) {
	var ui RRuiSupport
	funcname := "HomeUIHandler"
	tmpl := "home.html"

	t, err := template.New(tmpl).Funcs(RRfuncMap).ParseFiles("./html/home.html")
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
