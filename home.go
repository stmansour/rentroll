package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"path/filepath"
	"rentroll/rlib"
	"rentroll/ws"
	"strings"

	"github.com/kardianos/osext"
)

// HomeUIHandler is the standard home url handler for Roller
//------------------------------------------------------------------
func HomeUIHandler(w http.ResponseWriter, r *http.Request) {
	internalHomeUIHandler(w, r, "home.html")
}

// RHomeUIHandler is the special Receipt-only interface for
// Isola Bella.
//------------------------------------------------------------------
func RHomeUIHandler(w http.ResponseWriter, r *http.Request) {
	internalHomeUIHandler(w, r, "rhome.html")
}

// HomeUIHandler sends the main UI to the browser
// The forms of the url that are acceptable:
//		/home/
//		/home/<lang>
//		/home/<lang>/<tmpl>
//
// <lang> specifies the language.  The default is en-us
// <tmpl> specifies which template to use. The default is "dflt"
//------------------------------------------------------------------
func internalHomeUIHandler(w http.ResponseWriter, r *http.Request, appPage string) {
	// var ui ReportContext
	var ui struct {
		Language      string
		Template      string
		BL            []rlib.Business
		ReportContent string
	}
	var err error
	funcname := "HomeUIHandler"
	//appPage := "home.html"
	lang := "en-us"
	tmpl := "default"

	cwd, err := osext.ExecutableFolder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := "/home/"                // this is the part of the URL that got us into this handler
	uri := r.RequestURI[len(path):] // this pulls off the specific request

	s, err := url.QueryUnescape(strings.TrimSpace(r.URL.String()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("HOME HANDLER:  RL = %s\n", s)

	f := rlib.Stripchars(r.FormValue("filename"), `"`)
	if len(f) > 0 {

		appPage = strings.TrimSpace(f)
	}

	// use   http://domain/home/{lang}/{tmpl}  to set template
	if len(uri) > 0 {
		s1 := strings.Split(uri, "?")
		sa := strings.Split(s1[0], "/")
		n := len(sa)
		if n > 0 {
			lang = sa[0]
			if n > 1 {
				tmpl = sa[1]
			}
		}
	}

	ui.Language = lang
	ui.Template = tmpl
	ui.BL, err = rlib.GetAllBiz(r.Context())
	if err != nil {
		rlib.Ulog("GetAllBiz: err = %s\n", err.Error())
	}

	clientDir := filepath.Join(cwd, "webclient")
	htmlDir := filepath.Join(clientDir, "html")
	tmplFile := filepath.Join(htmlDir, appPage)

	t, err := template.New(appPage).Funcs(ws.RRfuncMap).ParseFiles(tmplFile)
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
