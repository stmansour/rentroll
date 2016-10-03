package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
)

// 1MB
const MAXMEMORY = 1 * 1024 * 1024

// CmdCsvAssess is the HTTP handler for the Journal report request
func CmdCsvAssess(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport, tmpl *string) {
	*tmpl = "csvassess.html"
	fname := ""

	sp := r.FormValue("sourcepage")
	if sp != "assessment" {
		// fmt.Printf("Not called from sourcepage == assessment, returning\n")
		return
	}
	if err := r.ParseMultipartForm(MAXMEMORY); err != nil {
		rlib.Ulog("err = %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	// fmt.Printf("for range = r.MultipartForm.File...\n")
	for _, fileHeaders := range r.MultipartForm.File {
		fmt.Printf("fileHeaders = %v\n", fileHeaders)
		for _, fileHeader := range fileHeaders {
			fmt.Printf("fileHeader = %v\n", fileHeader)
			fmt.Printf("fileHeader.Filename = %s\n", fileHeader.Filename)
			fname = "xyz.txt"
			file, _ := fileHeader.Open()
			path := fmt.Sprintf("tmp/%s", fname)
			buf, _ := ioutil.ReadAll(file)
			ioutil.WriteFile(path, buf, os.ModePerm)
		}
	}

	if len(fname) > 0 {
		rcsv.LoadAssessmentsCSV(fname)

	}
}
