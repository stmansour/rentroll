package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"time"
)

// 1MB
const MAXMEMORY = 1 * 1024 * 1024

// GetUploadedFile is the HTTP handler for the Journal report request
func GetUploadedFile(w http.ResponseWriter, r *http.Request, fname, path *string, ui *RRuiSupport) int {
	if err := r.ParseMultipartForm(MAXMEMORY); err != nil {
		rlib.Ulog("err = %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	// fmt.Printf("for range = r.MultipartForm.File...\n")
	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			// fmt.Printf("fileHeader = %v\n", fileHeader)
			// fmt.Printf("fileHeader.Filename = %s\n", fileHeader.Filename)
			(*fname) = fmt.Sprintf("asm%d.txt", time.Now().UnixNano())
			file, err := fileHeader.Open()
			if err != nil {
				ui.ReportContent = fmt.Sprintf("Error with fileHeader.Open():  %s\n", err.Error())
				return 1
			}
			(*path) = fmt.Sprintf("tmp/%s", *fname)
			buf, err := ioutil.ReadAll(file)
			if err != nil {
				ui.ReportContent = fmt.Sprintf("Error with ioutil.ReadAll:  %s\n", err.Error())
				return 1
			}
			if err := ioutil.WriteFile(*path, buf, os.ModePerm); err != nil {
				ui.ReportContent = fmt.Sprintf("Error with ioutil.WriteFile:  %s\n", err.Error())
				return 1
			}
		}
	}
	return 0
}

func removeUploadFile(path string, ui *RRuiSupport) {
	err := os.Remove(path)
	if err != nil {
		ui.ReportContent += fmt.Sprintf("\nError deleting temp file: %s\n", err.Error())
	}
}

// CmdCsvAssess is the HTTP handler for the Journal report request
func CmdCsvAssess(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	// *tmpl = "csvassess.html"
	fname := ""
	path := ""

	sp := r.FormValue("sourcepage")
	if sp != "assessment" {
		return
	}
	if GetUploadedFile(w, r, &fname, &path, ui) != 0 {
		return
	}

	if len(fname) > 0 {
		rcsv.InitRCSV(&ui.D1, &ui.D2, xbiz)
		ui.ReportContent = rcsv.LoadAssessmentsCSV(path)
		t := rcsv.RRAssessmentsTable(xbiz.P.BID, &ui.D1, &ui.D2)
		ui.ReportContent += fmt.Sprintf("\nAssessments\nBusiness:  %s  (%s)\nPeriod:  %s - %s\n\n", xbiz.P.Name, xbiz.P.Designation, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))
		ui.ReportContent += t.SprintTable(rlib.RPTTEXT)
		removeUploadFile(path, ui)
	}
}

// CmdCsvRcpt is the HTTP handler for the Journal report request
func CmdCsvRcpt(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	// *tmpl = "csvassess.html"
	fname := ""
	path := ""

	sp := r.FormValue("sourcepage")
	if sp != "receipts" {
		return
	}
	if GetUploadedFile(w, r, &fname, &path, ui) != 0 {
		return
	}

	if len(fname) > 0 {
		rcsv.InitRCSV(&ui.D1, &ui.D2, xbiz)
		ui.ReportContent = rcsv.LoadReceiptsCSV(path)
		t := rcsv.RRReceiptsTable(xbiz.P.BID, &ui.D1, &ui.D2)
		ui.ReportContent += fmt.Sprintf("\nReceipts\nBusiness:  %s  (%s)\nPeriod:  %s - %s\n\n", xbiz.P.Name, xbiz.P.Designation, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))
		ui.ReportContent += t.SprintTable(rlib.RPTTEXT)
		removeUploadFile(path, ui)
	}
}
