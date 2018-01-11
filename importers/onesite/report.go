package onesite

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/importers/core"
	"rentroll/rlib"
	"rentroll/rrpt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// getSummaryReportSection1 used to get summary for table's section1
func getSummaryReportSection1(importTime time.Time, csvFile string) string {
	// get date
	importYear, importMonth, importDate := importTime.Date()
	importDt := fmt.Sprintf("%d/%d/%d", importMonth, importDate, importYear)

	// get local timezone
	tz, _ := importTime.Zone()

	// format time in Kitchen
	kitchenFormat := importTime.Format(time.Kitchen)

	importLocalTime := kitchenFormat + " " + tz

	var reportHeader string
	reportHeader += "Date: " + importDt + "\n"
	reportHeader += "Time: " + importLocalTime + "\n"
	reportHeader += "Import File: " + csvFile + "\n"
	reportHeader += "\n"
	return reportHeader
}

// generateSummaryReport used to generate summary report from argued struct
func generateSummaryReport(
	ctx context.Context,
	summaryCount map[int]map[string]int,
	BID int64,
	currentTime time.Time,
	csvFile string,
) string {

	var tbl gotable.Table
	tbl.Init()
	tbl.SetTitle("Accord RentRoll Onesite Importer\n")
	tbl.SetSection1(getSummaryReportSection1(currentTime, csvFile))
	tbl.SetSection2("Summary")

	tbl.AddColumn("Data Type", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Total Possible", 10, gotable.CELLINT, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Total Imported", 10, gotable.CELLINT, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Issues", 10, gotable.CELLINT, gotable.COLJUSTIFYLEFT)

	// evaluate import count
	err := core.GetImportedCount(ctx, summaryCount, BID)
	if err != nil {
		rlib.Ulog("generateSummaryReport: error = %s", err.Error())
		tbl.SetSection3(err.Error())
	}

	// sort indices
	summaryCountIndexes := []int{}
	for index := range summaryCount {
		summaryCountIndexes = append(summaryCountIndexes, index)
	}
	sort.Ints(summaryCountIndexes)

	for _, dbType := range summaryCountIndexes {

		// get each db type map
		countMap := summaryCount[dbType]

		// add row
		tbl.AddRow()
		tbl.Puts(-1, 0, core.DBTypeMap[dbType])
		tbl.Puti(-1, 1, int64(countMap["possible"]))
		tbl.Puti(-1, 2, int64(countMap["imported"]))
		tbl.Puti(-1, 3, int64(countMap["issues"]))
	}

	s, err := tbl.SprintTable()
	if err != nil {
		rlib.Ulog("generateSummaryReport: error = %s", err.Error())
	}
	return s
}

// generateDetailedReport gives detailed report with (rowNumber, unit, db type, reason)
func generateDetailedReport(
	csvErrors map[int][]string,
	unitMap map[int]string,
	summaryCount map[int]map[string]int,
) (string, bool) {

	// return detailed report, tell program should it generate csv report?
	// in case of no errors, but has some warnings then csv report needs to be generated

	csvReportGenerate := true

	var tbl gotable.Table
	tbl.Init()
	tbl.SetTitle("DETAILED REPORT BY UNIT")

	tbl.AddColumn("Input Line", 6, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Unit Name", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	// tbl.AddColumn("RentRoll DB Type", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 100, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	csvErrorIndexes := []int{}
	for rowIndex := range csvErrors {
		csvErrorIndexes = append(csvErrorIndexes, rowIndex)
	}
	sort.Ints(csvErrorIndexes)

	for _, rowIndex := range csvErrorIndexes {

		// get error from index
		reportError := csvErrors[rowIndex]

		// check that rowIndex is -1
		// -1 means no data found in csv
		if rowIndex == -1 {
			tbl.AddRow()
			tbl.Puts(-1, 0, "")
			tbl.Puts(-1, 1, "")
			// tbl.Puts(-1, 2, "") //rentroll db type
			tbl.Puts(-1, 2, reportError[0])

			// append detailed section
			s, err := tbl.SprintTable()
			if err != nil {
				rlib.Ulog("generateDetailedReport: error = %s", err)
			}

			// return
			csvReportGenerate = false
			return s, csvReportGenerate
		}

		// get unit from map
		unit, _ := unitMap[rowIndex]

		// used to separate errors, warnings
		rowErrors, rowWarnings := []string{}, []string{}

		for _, reason := range reportError {
			if strings.HasPrefix(reason, "E:") {

				// if any error captured then do not generate csv report
				csvReportGenerate = false

				// red color
				reason = strings.Replace(reason, "E:", "", -1)

				// if error not appended already then
				if !core.StringInSlice(reason, rowErrors) {
					rowErrors = append(rowErrors, reason)
				}
			}
			if strings.HasPrefix(reason, "W:") {
				// orange color
				reason = strings.Replace(reason, "W:", "", -1)

				// if warning not appended already then
				if !core.StringInSlice(reason, rowWarnings) {
					rowWarnings = append(rowWarnings, reason)
				}
			}
		}

		// first put errors
		for _, errorText := range rowErrors {
			errorText := strings.Split(errorText, ">:")
			dbType, reason := errorText[0], errorText[1]
			dbType = strings.Replace(dbType, "<", "", -1)
			dbTypeInt, _ := strconv.Atoi(dbType)

			// count issues in summary report
			summaryCount[dbTypeInt]["issues"]++

			// put in tabl
			tbl.AddRow()
			tbl.Puts(-1, 0, strconv.Itoa(rowIndex))
			tbl.Puts(-1, 1, unit)
			// tbl.Puts(-1, 2, core.DBTypeMap[dbTypeInt])
			tbl.Puts(-1, 2, reason)
		}

		// then warnings
		for _, warningText := range rowWarnings {
			warningText := strings.Split(warningText, ">:")
			dbType, reason := warningText[0], warningText[1]
			dbType = strings.Replace(dbType, "<", "", -1)
			dbTypeInt, _ := strconv.Atoi(dbType)

			// prefixed with "Warning: "
			reason = "Warning: " + reason

			// count issues in summary report
			summaryCount[dbTypeInt]["issues"]++

			tbl.AddRow()
			tbl.Puts(-1, 0, strconv.Itoa(rowIndex))
			tbl.Puts(-1, 1, unit)
			// tbl.Puts(-1, 2, core.DBTypeMap[dbTypeInt])
			tbl.Puts(-1, 2, reason)
		}
	}

	// append detailed section
	s, err := tbl.SprintTable()
	if err != nil {
		rlib.Ulog("generateDetailedReport: error = %s", err)
	}

	// return
	return s, csvReportGenerate
}

// generateRCSVReport return report for all type of csv defined here from rcsv
func generateRCSVReport(
	ctx context.Context,
	business *rlib.Business,
	summaryCount map[int]map[string]int,
	csvFile string,
) string {

	var r = []rrpt.ReporterInfo{
		{ReportNo: 5, OutputFormat: gotable.TABLEOUTTEXT, Handler: rrpt.RRreportRentableTypes, Bid: business.BID},
		{ReportNo: 6, OutputFormat: gotable.TABLEOUTTEXT, Handler: rrpt.RRreportRentables, Bid: business.BID},
		{ReportNo: 7, OutputFormat: gotable.TABLEOUTTEXT, Handler: rrpt.RRreportPeople, Bid: business.BID},
		{ReportNo: 9, OutputFormat: gotable.TABLEOUTTEXT, Handler: rrpt.RRreportRentalAgreements, Bid: business.BID},
		{ReportNo: 14, OutputFormat: gotable.TABLEOUTTEXT, Handler: rrpt.RRreportCustomAttributes, Bid: business.BID},
		{ReportNo: 15, OutputFormat: gotable.TABLEOUTTEXT, Handler: rrpt.RRreportCustomAttributeRefs, Bid: business.BID},
	}

	var rcsvReport string

	title := fmt.Sprintf("RECORDS FOR BUSINESS UNIT DESIGNATION: %s", business.Name)
	rcsvReport += strings.Repeat("=", len(title))
	rcsvReport += "\n" + title + "\n"
	rcsvReport += strings.Repeat("=", len(title))
	rcsvReport += "\n\n"

	for i := 0; i < len(r); i++ {
		rcsvReport += r[i].Handler(ctx, &r[i])
		rcsvReport += strings.Repeat("=", len(title))
		rcsvReport += "\n"
	}

	return rcsvReport
}

// successReport generates success report
func successReport(
	ctx context.Context,
	business *rlib.Business,
	summaryCount map[int]map[string]int,
	csvFile string,
	debugMode int,
	currentTime time.Time,
) string {

	var report string

	// append summary report
	report += generateSummaryReport(ctx, summaryCount, business.BID, currentTime, csvFile)
	report += "\n"

	// csv report for all types if testmode is on
	if debugMode == 1 {
		report += generateRCSVReport(ctx, business, summaryCount, csvFile)
	}

	// return
	return report
}

// errorReporting used to report the errors for onesite csv
func errorReporting(
	ctx context.Context,
	business *rlib.Business,
	csvErrors map[int][]string,
	unitMap map[int]string,
	summaryCount map[int]map[string]int,
	csvFile string,
	debugMode int,
	currentTime time.Time,
) (string, bool) {

	var errReport string

	// first generate detailed report because summary count also be used in it
	// but append it after summary report
	detailedReport, csvReportGenerate := generateDetailedReport(csvErrors, unitMap, summaryCount)
	detailedReport += "\n"

	// append summary report
	errReport += generateSummaryReport(ctx, summaryCount, business.BID, currentTime, csvFile)
	errReport += "\n"

	// append detailedReport
	errReport += detailedReport

	// if true then generate csv report
	// specia case: when there are only warnings but no errors
	if csvReportGenerate && debugMode == 1 {
		errReport += generateRCSVReport(ctx, business, summaryCount, csvFile)
	}

	// return
	return errReport, csvReportGenerate
}
