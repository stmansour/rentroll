// The purpose of this test is
package main

import (
	"context"
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"gotable"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	//"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir        *sql.DB        // phonebook db
	dbrr         *sql.DB        //rentroll db
	DBDir        string         // phonebook database
	DBRR         string         //rentroll database
	DBUser       string         // user for all databases
	PortRR       int            // rentroll port
	Bud          string         // Biz Unit Descriptor
	Xbiz         rlib.XBusiness // lots of info about this biz
	NoAuth       bool
	SkipVacCheck bool   // until the code is modified to process on each command entered, if set to false, this inibits batch processing to do vacancy calc.
	CSVLoad      string // if loading csv, this string will have index,filename
	sStart       string // start time
	sStop        string // stop time
	Report       string // if testing engine, which report/action to perform
}

// DispatchCtx is a type of struct needed for the Dispatch function. It defines
// everything needed to run a particular command. It is the responsibility of the
// caller to fill out all the needed dCtx information. Not all information is needed
// for all commands.
type DispatchCtx struct {
	Cmd          int                 // cmd to execute
	DtStart      time.Time           // period start time
	DtStop       time.Time           // period end
	xbiz         rlib.XBusiness      // BUD from cmd line
	OutputFormat int                 // how shall the output be formatted
	Report       int64               // which report to generate - this is used in batch mode operation
	Args         string              // full command line report string
	CSVLoadStr   string              // if loading CSV file, this will have index,filename
	w            http.ResponseWriter // for web responses
	r            *http.Request       // for web request introspection
}

// ReporterInfo is for routines that want to table-ize their reporting using
// the CSV library's simple report routines.
type ReporterInfo struct {
	ReportNo              int             // index number of the report
	OutputFormat          int             // text, html, maybe more in the future
	EDI                   int             // end date inclusive -- 0 = no, 1 = yes
	Bid                   int64           // associated business
	Raid                  int64           // associated Rental Agreement if needed
	D1                    time.Time       // associated date if needed
	D2                    time.Time       // associated date if needed
	ID                    int64           // specific id if a single entity is being printed
	NeedsBID              bool            // true if BID is needed for this report
	NeedsRAID             bool            // true if RAID is needed for this report
	NeedsDt               bool            // true if a Date is needed for this report
	NeedsID               bool            // true if the report requires an id
	RptHeaderD1           bool            // true if the report's header should contain D1
	RptHeaderD2           bool            // true if the dates should show as a range D1 - D2
	BlankLineAfterRptName bool            // true if a blank line should be added after the Report Name
	Style                 int             // some printouts may have multiple styles. This is the selector
	Xbiz                  *rlib.XBusiness // may not be set in all cases
	Handler               func(context.Context, *ReporterInfo) string
	QueryParams           *url.Values
}

// RTIDCount is for counting rentables of a particular type
type RTIDCount struct {
	RT    rlib.RentableType // ID of the types we're counting
	Count int64             // the count
}

// COLJUSTIFYLEFT et. al. are the constants used in the Table class
const (
	COLJUSTIFYLEFT  = 1
	COLJUSTIFYRIGHT = 2

	CELLINT      = 1
	CELLFLOAT    = 2
	CELLSTRING   = 3
	CELLDATE     = 4
	CELLDATETIME = 5

	TABLEOUTTEXT = 1
	TABLEOUTHTML = 2
	TABLEOUTPDF  = 3
	TABLEOUTCSV  = 4

	CSSFONTSIZE = 14
	NEWLINE     = "\n"
)

// RReportTableErrorSectionCSS holds css for errors placed in section3 of gotable
var RReportTableErrorSectionCSS = []*gotable.CSSProperty{
	{Name: "color", Value: "red"},
	{Name: "font-family", Value: "monospace"},
}

// WKHTMLTOPDFCMD command : html > pdf
const (
	WKHTMLTOPDFCMD = "wkhtmltopdf"
	TEMPSTORE      = "."
	DATETIMEFMT    = "_2 Jan 2006 3:04 PM" // Actual Format: _2 Jan 2006 3:04 PM UTC
)

var pdfProps = []*gotable.PDFProperty{
	//
	{Option: "--no-collate"},
	// top margin
	{Option: "-T", Value: "15"},
	// header center content
	{Option: "--header-center", Value: "Smoke Test Report Table"},
	// header font size
	{Option: "--header-font-size", Value: "7"},
	// header font
	{Option: "--header-font-name", Value: "opensans"},
	// header spacing
	{Option: "--header-spacing", Value: "3"},
	// bottom margin
	{Option: "-B", Value: "15"},
	// footer spacing
	{Option: "--footer-spacing", Value: "5"},
	// footer font
	{Option: "--footer-font-name", Value: "opensans"},
	// footer font size
	{Option: "--footer-font-size", Value: "7"},
	// footer left content
	{Option: "--footer-left", Value: time.Now().Format(DATETIMEFMT)},
	// footer right content
	{Option: "--footer-right", Value: "Page [page] of [toPage]"},
	// page size
	{Option: "--page-size", Value: "Letter"},
	// orientation
	{Option: "--orientation", Value: "Landscape"},
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)") //
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")
	noauth := flag.Bool("noauth", false, "if specified, inhibit authentication")

	flag.Parse()

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.PortRR = *portPtr
	App.Bud = *pBud
	App.NoAuth = *noauth
}

// getRRTable returns a table with some basic initialization
// to be used in all reports of rentroll software
func getRRTable() gotable.Table {
	var tbl gotable.Table
	tbl.Init()

	// after table is ready then set css only
	// section3 will be used as error section
	// so apply css here
	tbl.SetSection3CSS(RReportTableErrorSectionCSS)
	tbl.SetNoRowsCSS(RReportTableErrorSectionCSS)
	tbl.SetNoHeadersCSS(RReportTableErrorSectionCSS)

	return tbl
}

// ReportTextOutput tests the rentable type report in text format
func ReportTextOutput(tbl *gotable.Table) {
	(*tbl).TightenColumns()

	fname := "RentableTypeReport_test.txt"
	f, err := os.Create(fname)
	if nil != err {
		//t.Errorf("RentaleCountByRentableTypeReport_test: Error creating file %s: %s\n", fname, err.Error())
		//fmt.Printf("RentaleCountByRentableTypeReport_test: Error creating file: %s\n", err.Error())
		log.Fatalf("%s: RentableTypeReport_test: Error creating file: %s\n", fname, err.Error())
	}

	if err := tbl.TextprintTable(f); err != nil {
		log.Fatalf("RentableTypeReport_test: Error creating TEXT output: %s\n", err.Error())
	}
	// close file after operation
	f.Close()

	// now compare what we have to the known-good output
	b, _ := ioutil.ReadFile("./testdata/RentableTypeReport_test.txt")
	sb, _ := ioutil.ReadFile("./RentableTypeReport_test.txt")

	if len(b) != len(sb) {
		// fmt.Printf("smoke_test: Expected len = %d,  found len = %d\n", len(b), len(sb))
		log.Fatalf("RentableTypeReport_test: Expected len = %d,  found len = %d\n", len(b), len(sb))
	}
	if len(sb) > 0 && len(b) > 0 {
		for i := 0; i < len(b); i++ {
			if i < len(sb) && sb[i] != b[i] {
				log.Fatalf("RentableTypeReport_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, b[i], b[i], sb[i], sb[i])
				// fmt.Printf("smoke_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, b[i], b[i], sb[i], sb[i])
				break
			}
		}
	}

	// add test case for SprintTable, it should be same as TextPrintTable output
	s, err := tbl.SprintTable()
	st := []byte(s)
	if err != nil {
		log.Fatalf("RentableTypeReport_test: Error creating TEXT output: %s\n", err.Error())
	}
	if len(st) != len(sb) {
		log.Fatalf("RentableTypeReport_test: Expected len = %d,  found len = %d\n", len(st), len(sb))
	}
	if len(sb) > 0 && len(st) > 0 {
		for i := 0; i < len(st); i++ {
			if i < len(sb) && sb[i] != st[i] {
				log.Fatalf("RentableTypeReport_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, st[i], st[i], sb[i], sb[i])
				break
			}
		}
	}
}

// ReportCSVOutput tests the rentable type report in CSV format
func ReportCSVOutput(tbl *gotable.Table) {
	fname := "RentableTypeReport_test.csv"
	f, err := os.Create(fname)
	if nil != err {
		log.Fatalf("RentableTypeReport_test: Error creating file %s: %s\n", fname, err.Error())
		// fmt.Printf("RentableTypeReport_test: Error creating file: %s\n", err.Error())
	}

	if err := tbl.CSVprintTable(f); err != nil {
		log.Fatalf("RentableTypeReport_test: Error creating CSV output: %s\n", err.Error())
	}
	// close file after operation
	f.Close()

	// now compare what we have to the known-good output
	b, _ := ioutil.ReadFile("./testdata/RentableTypeReport_test.csv")
	sb, _ := ioutil.ReadFile("./RentableTypeReport_test.csv")

	if len(b) != len(sb) {
		// fmt.Printf("smoke_test: Expected len = %d,  found len = %d\n", len(b), len(sb))
		log.Fatalf("RentableTypeReport_test: Expected len = %d,  found len = %d\n", len(b), len(sb))
	}
	if len(sb) > 0 && len(b) > 0 {
		for i := 0; i < len(b); i++ {
			if i < len(sb) && sb[i] != b[i] {
				log.Fatalf("RentableTypeReport_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, b[i], b[i], sb[i], sb[i])
				// fmt.Printf("smoke_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, b[i], b[i], sb[i], sb[i])
				break
			}
		}
	}
}

// ReportHTMLOutput tests the rentable type report in HTML format
func ReportHTMLOutput(tbl *gotable.Table) {
	fname := "RentableTypeReport_test.html"
	f, err := os.Create(fname)
	if nil != err {
		log.Fatalf("RentableTypeReport_test: Error creating file %s: %s\n", fname, err.Error())
		// fmt.Printf("smoke_test: Error creating file: %s\n", err.Error())
	}

	if err := tbl.HTMLprintTable(f); err != nil {
		log.Fatalf("RentableTypeReport_test: Error creating HTML output: %s\n", err.Error())
	}
	// close file after operation
	f.Close()

	// now compare what we have to the known-good output
	b, _ := ioutil.ReadFile("./testdata/RentableTypeReport_test.html")
	sb, _ := ioutil.ReadFile("./RentableTypeReport_test.html")

	if len(b) != len(sb) {
		// fmt.Printf("RentableTypeReport_test: Expected len = %d,  found len = %d\n", len(b), len(sb))
		log.Fatalf("RentableTypeReport_test: Expected len = %d,  found len = %d\n", len(b), len(sb))
	}
	if len(sb) > 0 && len(b) > 0 {
		for i := 0; i < len(b); i++ {
			if i < len(sb) && sb[i] != b[i] {
				log.Fatalf("RentableTypeReport_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, b[i], b[i], sb[i], sb[i])
				// fmt.Printf("RentableTypeReport_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, b[i], b[i], sb[i], sb[i])
				break
			}
		}
	}
}

// ReportPDFOutput tests the rentable type report in PDF format
func ReportPDFOutput(tbl *gotable.Table) {
	fname := "RentableTypeReport_test.pdf"
	f, err := os.Create(fname)
	if nil != err {
		log.Fatalf("RentableTypeReport_test: Error creating file %s: %s\n", fname, err.Error())
		// fmt.Printf("RentableTypeReport_test: Error creating file: %s\n", err.Error())
	}

	if err := tbl.PDFprintTable(f, pdfProps); err != nil {
		log.Fatalf("RentableTypeReport_test: Error creating PDF output: %s\n", err.Error())
	}
	// close file after operation
	f.Close()

	// now compare what we have to the known-good output
	b, _ := ioutil.ReadFile("./testdata/RentableTypeReport_test.pdf")
	sb, _ := ioutil.ReadFile("./RentableTypeReport_test.pdf")

	if len(sb) == 0 {
		log.Fatalf("RentableTypeReport_test: Expected some content in PDF output file,  found len = 0")
	}

	if len(b) != len(sb) {
		// fmt.Printf("RentableTypeReport_test: Expected len = %d,  found len = %d\n", len(b), len(sb))
		log.Fatalf("RentableTypeReport_test: Expected len = %d,  found len = %d\n", len(b), len(sb))
	}
	if len(sb) > 0 && len(b) > 0 {
		for i := 0; i < len(b); i++ {
			if i < len(sb) && sb[i] != b[i] {
				log.Fatalf("RentableTypeReport_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, b[i], b[i], sb[i], sb[i])
				// fmt.Printf("RentableTypeReport_test: micompare at character %d, expected %x (%c), found %x (%c)\n", i, b[i], b[i], sb[i], sb[i])
				break
			}
		}
	}
}

func main() {
	var err error
	readCommandLineArgs()
	rlib.RRReadConfig()

	//----------------------------
	// Open RentRoll database
	//----------------------------
	if err = rlib.RRReadConfig(); err != nil {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	s := extres.GetSQLOpenString(rlib.AppConfig.RRDbname, &rlib.AppConfig)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = extres.GetSQLOpenString(rlib.AppConfig.Dbname, &rlib.AppConfig)
	App.dbdir, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open: Error = %v\n", err)
		os.Exit(1)
	}
	err = App.dbdir.Ping()
	if nil != err {
		fmt.Printf("dbdir.Ping: Error = %v\n", err)
		os.Exit(1)
	}

	rlib.RpnInit()
	App.NoAuth = true // for testing
	rlib.InitDBHelpers(App.dbrr, App.dbdir)
	rlib.SetNoAuthFlag(App.NoAuth)
	rlib.SessionInit(10) // must be called before calling InitBizInternals

	// create background context
	ctx := context.Background()

	biz, err := rlib.GetBusinessByDesignation(ctx, App.Bud)
	if err != nil {
		fmt.Printf("Could not find Business Unit named %s, Error: %s\n", App.Bud, err.Error())
		os.Exit(1)
	}
	if biz.BID == 0 {
		fmt.Printf("Could not find Business Unit named %s\n", App.Bud)
		os.Exit(1)
	}

	err = rlib.InitBizInternals(biz.BID, &App.Xbiz)
	if err != nil {
		fmt.Printf("Error in InitBizInternals: %s\n", err.Error())
		os.Exit(1)
	}

	DoTestRentableType(ctx)
}

// DoTestRentableType tests the rentable type report
func DoTestRentableType(ctx context.Context) {

	var ri = rrpt.ReporterInfo{
		OutputFormat: gotable.TABLEOUTTEXT,
		NeedsBID:     true,  // true if BID is needed for this report
		NeedsRAID:    false, // true if RAID is needed for this report
		NeedsID:      false, // true if the report requires an id
		//D1:           tmpT1, // associated date if needed
		//D2:           tmpT2,
		//Xbiz: &App.Xbiz,
	}
	ri.Bid = 1
	tlb := rrpt.RRreportRentableTypesTable(ctx, &ri)
	ReportTextOutput(&tlb)
	ReportCSVOutput(&tlb)
	ReportHTMLOutput(&tlb)
	ReportPDFOutput(&tlb)
}
