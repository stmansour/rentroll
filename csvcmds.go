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
		m := rcsv.LoadAssessmentsCSV(path)
		ui.ReportContent += rcsv.ErrlistToString(&m)
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
		m := rcsv.LoadReceiptsCSV(path)
		ui.ReportContent = rcsv.ErrlistToString(&m)
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

	rlib.GenerateRecurInstances(xbiz, &ui.D1, &ui.D2) // generate and process assessment instances in this range
	rlib.ProcessReceiptRange(xbiz, &ui.D1, &ui.D2)    // process receipts in this range
	ui.ReportContent += fmt.Sprintf("\nJournal\nBusiness:  %s  (%s)\nPeriod:  %s - %s\n\n", xbiz.P.Name, xbiz.P.Designation, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))
	t := rrpt.JournalReport(xbiz, &ui.D1, &ui.D2)
	ui.ReportContent += t.SprintTable(rlib.RPTTEXT)
}

// CmdGenVac is the HTTP handler for generating Vacancy Journal records
func CmdGenVac(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	if xbiz.P.BID > 0 {
		nr := rlib.GenVacancyJournals(xbiz, &ui.D1, &ui.D2)
		ui.ReportContent = fmt.Sprintf("Processed range %s - %s.  Vacancy records added: %d\n", ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4), nr)
	}
}

// CmdGenLdg is the HTTP handler for generating Vacancy Journal records
func CmdGenLdg(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	// fmt.Printf("CmdGenLdg: BID=%d, d1 = %s, d2 = %s\n", xbiz.P.BID, ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4))
	if xbiz.P.BID > 0 {
		nr := rlib.GenerateLedgerRecords(xbiz, &ui.D1, &ui.D2)
		ui.ReportContent = fmt.Sprintf("Processed range %s - %s.  Ledger records added: %d\n", ui.D1.Format(rlib.RRDATEFMT4), ui.D2.Format(rlib.RRDATEFMT4), nr)
		RptLedgerActivity(w, r, xbiz, ui)
	}
}

type csvLoaderT struct {
	prefix  string              // id prefix
	handler func(string) string // actual csv loader
}

var loaders = []csvLoaderT{
//{prefix: "ASM", handler: rcsv.LoadAssessmentsCSV},
// {prefix: "B", handler: rcsv.LoadBusinessCSV},
// {prefix: "C", handler: rcsv.LoadCustomAttributesCSV},
// {prefix: "CR", handler: rcsv.LoadCustomAttributeRefsCSV},
// {prefix: "COA", handler: rcsv.LoadChartOfAccountsCSV},
// {prefix: "DPM", handler: rcsv.LoadDepositMethodsCSV},
// {prefix: "DEP", handler: rcsv.LoadDepositoryCSV},
// {prefix: "RT", handler: rcsv.LoadRentableTypesCSV},
// {prefix: "SL", handler: rcsv.LoadStringTablesCSV},
// //{prefix: "T", handler: rcsv.LoadPeopleCSV},
// {prefix: "PMT", handler: rcsv.LoadPaymentTypesCSV},
// {prefix: "R", handler: rcsv.LoadRentablesCSV},
// {prefix: "RA", handler: rcsv.LoadRentalAgreementCSV},
// {prefix: "RAT", handler: rcsv.LoadRentalAgreementTemplatesCSV},
// {prefix: "RCPT", handler: rcsv.LoadReceiptsCSV},
// {prefix: "RSP", handler: rcsv.LoadRentalSpecialtiesCSV},
}

func csvloadReporter(prefix string, xbiz *rlib.XBusiness, ui *RRuiSupport) string {
	// fmt.Printf("csvloadReporter: prefix=%s, BID=%d\n", prefix, xbiz.P.BID)
	switch prefix {
	case "ASM", "Assessments":
		return rcsv.RRreportAssessments(rlib.RPTTEXT, xbiz.P.BID)
	case "B", "Business":
		return rcsv.RRreportBusiness(rlib.RPTTEXT)
	case "COA", "Chart Of Accounts":
		rlib.InitBizInternals(xbiz.P.BID, xbiz)
		return rcsv.RRreportChartOfAccounts(rlib.RPTTEXT, xbiz.P.BID)
	case "C", "Custom Attributes":
		return rcsv.RRreportCustomAttributes(rlib.RPTTEXT)
	case "CR", "Custom Attribute Refs":
		return rcsv.RRreportCustomAttributeRefs(rlib.RPTTEXT)
	case "DPM", "Deposit Methods":
		return rcsv.RRreportDepositMethods(rlib.RPTTEXT, xbiz.P.BID)
	case "DEP", "Depositories":
		return rcsv.RRreportDepository(rlib.RPTTEXT, xbiz.P.BID)
	case "PMT", "Payment Types":
		return rcsv.RRreportPaymentTypes(rlib.RPTTEXT, xbiz.P.BID)
	case "R", "Rentables":
		return rcsv.RRreportRentables(rlib.RPTTEXT, xbiz.P.BID)
	case "RA", "Rental Agreements":
		return rcsv.RRreportRentalAgreements(rlib.RPTTEXT, xbiz.P.BID)
	case "RAT", "Rental Agreement Templates":
		return rcsv.RRreportRentalAgreementTemplates(rlib.RPTTEXT, xbiz.P.BID)
	case "RCPT", "Receipts":
		return rcsv.RRreportReceipts(rlib.RPTTEXT, xbiz.P.BID)
	case "RT", "Rentable Types":
		return rcsv.RRreportRentableTypes(rlib.RPTTEXT, xbiz.P.BID)
	case "Rentable Count By Type":
		return rrpt.RentableCountByRentableTypeReport(rlib.RPTTEXT, xbiz, &ui.D1, &ui.D2)
	case "SL", "String Lists":
		return rcsv.RRreportStringLists(rlib.RPTTEXT, xbiz.P.BID)
	case "Statements":
		return rrpt.RptStatementTextReport(xbiz, &ui.D1, &ui.D2)
	case "T", "People":
		return rcsv.RRreportPeople(rlib.RPTTEXT, xbiz.P.BID)
	}
	return "unhandled loader type: " + prefix
}

// CmdSimpleReport returns the report output in ui.ReportContent for the provid
func CmdSimpleReport(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	action := r.FormValue("action")
	// fmt.Printf("Simple report:  xbiz.P.BID = %d, action = %s\n", xbiz.P.BID, action)
	if xbiz.P.BID > 0 {
		ui.ReportContent = csvloadReporter(action, xbiz, ui)
	}
}

// CmdCSVLoad is the HTTP handler for loading csv viles
func CmdCSVLoad(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	sp := r.FormValue("sourcepage")
	if sp != "csvload" {
		return
	}

	action := r.FormValue("action")
	if len(action) > 0 {
		for i := 0; i < len(loaders); i++ {
			if action == loaders[i].prefix {
				fname := ""
				path := ""
				if GetUploadedFile(w, r, &fname, &path, ui) != 0 {
					ui.ReportContent = "Failed to upload file"
					return
				}
				if len(fname) > 0 {
					rcsv.InitRCSV(&ui.D1, &ui.D2, xbiz)
					ui.ReportContent = loaders[i].handler(path)
					ui.ReportContent += csvloadReporter(loaders[i].prefix, xbiz, ui)
				}
			}
		}
	}
}
