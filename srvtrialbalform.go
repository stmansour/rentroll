package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"text/template"
)

func srvformTrialBalance(w http.ResponseWriter, r *http.Request) {
	var ui RRuiSupport
	funcname := "serveFormTrialBalance"
	w.Header().Set("Content-Type", "text/html")

	t, err := template.New("formtrialbal.html").Funcs(RRfuncMap).ParseFiles("./html/formtrialbal.html")
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}
	UIInitBizList(&ui)
	err = t.Execute(w, &ui)
	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
