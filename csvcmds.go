package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"time"
)

// 1MB
const MAXMEMORY = 1 * 1024 * 1024

var tmpdir = string("tmp")

// GetUploadedFile is the HTTP handler for the Journal report request
func GetUploadedFile(w http.ResponseWriter, r *http.Request, fname, path *string, ui *RRuiSupport) int {
	ok, _ := PathExists(tmpdir)
	if !ok {
		err := os.Mkdir(tmpdir, 0777)
		if err != nil {
			ui.ReportContent += fmt.Sprintf("Error creating directory %s: %s\n", tmpdir, err.Error())
			return 1
		}
	}
	if err := r.ParseMultipartForm(MAXMEMORY); err != nil {
		ui.ReportContent += fmt.Sprintf("error with ParseMultipartForm: err = %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusForbidden)
		return 1
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			(*fname) = fmt.Sprintf("asm%d.txt", time.Now().UnixNano())
			file, err := fileHeader.Open()
			if err != nil {
				ui.ReportContent += fmt.Sprintf("Error with fileHeader.Open():  %s\n", err.Error())
				return 1
			}
			(*path) = fmt.Sprintf("tmp/%s", *fname)
			buf, err := ioutil.ReadAll(file)
			if err != nil {
				ui.ReportContent += fmt.Sprintf("Error with ioutil.ReadAll:  %s\n", err.Error())
				return 1
			}
			if err := ioutil.WriteFile(*path, buf, os.ModePerm); err != nil {
				ui.ReportContent += fmt.Sprintf("Error with ioutil.WriteFile:  %s\n", err.Error())
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
		ui.ReportContent += rcsv.LoadAssessmentsCSV(path)
		t := rcsv.RRAssessmentsTable(xbiz.P.BID, &ui.D1, &ui.D2)
		ui.ReportContent += fmt.Sprintf("\nAssessments\nBusiness:  %s  (%s)\nPeriod:  %s - %s\n\n", xbiz.P.Name, xbiz.P.Designation, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))
		ui.ReportContent += t.SprintTable(rlib.RPTTEXT)
		removeUploadFile(path, ui)
	}
}

// CmdCsvRcpt is the HTTP handler for the Journal report request
func CmdCsvRcpt(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
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

// CmdGenJnl is the HTTP handler for generating Journal recurring instances
func CmdGenJnl(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	sp := r.FormValue("sourcepage")
	if sp != "genjnl" {
		return
	}

	rlib.GenerateJournalRecords(xbiz, &ui.D1, &ui.D2, false)
	// rlib.GenerateLedgerRecords(&ctx.xbiz, &ctx.DtStart, &ctx.DtStop)
	ui.ReportContent += fmt.Sprintf("\nJournal\nBusiness:  %s  (%s)\nPeriod:  %s - %s\n\n", xbiz.P.Name, xbiz.P.Designation, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))
	t := rrpt.JournalReport(xbiz, &ui.D1, &ui.D2)
	ui.ReportContent += t.SprintTable(rlib.RPTTEXT)
}
