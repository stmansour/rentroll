// This code loads a csv file with Business definitions.  Its format is:
//
// 		Company Designation, Company Name, Default Occupancy Type, Parking Permit In Use
//
// The code will try to load a Phonebook company with the designation supplied. It will pull
// default information from the Phonebook entry. For now, the information only includes the
// company name.
//
// Examples:
// 		REH,,4,0
// 		BBBB,Big Bob's Barrel Barn,4,0
package main

import (
	"context"
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"gotable"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir          *sql.DB                    // phonebook db
	dbrr           *sql.DB                    //rentroll db
	AcctDep        string                     // account depository
	ARFile         string                     // account rules
	AsmtFile       string                     // Assessments
	AssignFile     string                     // assign custom attributes
	BizFile        string                     // name of csv file with new biz info
	BldgFile       string                     // Buildings for this Business
	BUD            string                     // business unit designator
	CoaFile        string                     // chart of accounts
	CustomFile     string                     // custom attributes
	DBDir          string                     // phonebook database
	DBRR           string                     //rentroll database
	DBUser         string                     // user for all databases
	DepositFile    string                     // Deposits
	DepositoryFile string                     // Depository
	DMFile         string                     // Deposit Methods
	InvoiceFile    string                     // Invoice
	NoteTypeFile   string                     // note types
	PetFile        string                     // assign pets
	PmtTypeFile    string                     // payment types
	PmtTypes       map[int64]rlib.PaymentType // all payment types in this db
	PplFile        string                     // people for this Business
	RaFile         string                     // rental agreement cvs file
	RatFile        string                     // rentalAgreementTemplates
	RcptFile       string                     // receipts of payments
	Report         string                     // Report: 1 = Journal, 2 = Ledger, 3 = Businesses, 4 = Rentable types
	RFile          string                     // rentables csv file
	RPFile         string                     // RatePlans
	RPRefFile      string                     // RatePlanRef
	RPRRTRateFile  string                     // RatePlanRefRTRate
	RPRSPRateFile  string                     // RatePlanRefSPRate
	RSpFile        string                     // Rentable specialties
	RspRefsFile    string                     // assign specialties to rentables
	RTFile         string                     // Rentable types csv file
	SLFile         string                     // StringLists
	SrcFile        string                     // Sources
	VehicleFile    string                     // vehicles that belong to people
	DtStart        time.Time                  // range start time
	DtStop         time.Time                  // range stop time
	Xbiz           rlib.XBusiness             // xbusiness associated with -G  (BUD)
	NoAuth         bool                       // if true then skip authentication
}

func readCommandLineArgs() {
	asmtPtr := flag.String("A", "", "add Assessments via csv file")
	arPtr := flag.String("ar", "", "add AccountRules via csv file")
	rpptr := flag.String("a", "", "add RatePlans via csv file")
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	bizPtr := flag.String("b", "", "add Business via csv file")
	raPtr := flag.String("C", "", "add rental agreements via csv file")
	coaPtr := flag.String("c", "", "add chart of accounts via csv file")
	bldgPtr := flag.String("D", "", "add Buildings to a Business via csv file")
	depositoryPtr := flag.String("d", "", "add Depositories to a Business via csv file")
	petPtr := flag.String("E", "", "assign pets to a Rental Agreement via csv file")
	rcptPtr := flag.String("e", "", "add receipts via csv file")
	rsrefsPtr := flag.String("F", "", "assign rentable specialties to rentables via csv file")
	rprptr := flag.String("f", "", "add RatePlanRefs via csv file")
	pDates := flag.String("g", "", "Date Range.  Example: 1/1/16,2/1/16")
	pBUD := flag.String("G", "", "BUD - business unit designator")
	pAD := flag.String("H", "", "add Account Depositories via csv file")
	invPtr := flag.String("i", "", "add Invoices via csv file")
	lptr := flag.String("L", "", "Report: 1-jnl, 2-ldg, 3-biz, 4-asmtypes, 5-rtypes, 6-rentables, 7-people, 8-rat, 9-ra, 10-coa, 11-asm, 12-payment types, 13-receipt list, 14-CustAttr, 15-CustAttrRef, 16-Pets, 17-NoteTypes, 18-Depositories, 19-Deposits, 20-Invoices, 21-Specialties, 22-Specialty Assignments, 23-Deposit Methods, 24-Sources, 25-StringList, 26-RatePlan, 27-RatePlanRef,BUD,RatePlanName, 28-BUD, 29-AcctRules, 30-Receipt, 31-HotelReceipt")
	slPtr := flag.String("l", "", "add StringLists via csv file")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	dmPtr := flag.String("m", "", "add DepositMethods via csv file")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	rprrt := flag.String("n", "", "add RatePlanRefRTRate via csv file")
	ntPtr := flag.String("O", "", "add NoteTypes via csv file")
	pmtPtr := flag.String("P", "", "add payment types via csv file")
	pPtr := flag.String("p", "", "add people via csv file")
	rtPtr := flag.String("R", "", "add Rentable types via csv file")
	rPtr := flag.String("r", "", "add rentables via csv file")
	src := flag.String("S", "", "add Sources via csv file")
	rspPtr := flag.String("s", "", "add Rentable specialties via csv file")
	ratPtr := flag.String("T", "", "add rental agreement templates via csv file")
	rprsp := flag.String("t", "", "add RatePlanRef Specialty Rates via csv file")
	asgnPtr := flag.String("U", "", "assign custom attributes via csv file")
	custPtr := flag.String("u", "", "add custom attributes via csv file")
	vehiclePtr := flag.String("V", "", "add people vehicles via csv file")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	depositPtr := flag.String("y", "", "add Deposits via csv file")
	noconPtr := flag.Bool("nocon", false, "if specified, inhibit Console output")
	noauth := flag.Bool("noauth", false, "if specified, inhibit authentication")

	flag.Parse()
	if *verPtr {
		fmt.Printf("Version:    %s\nBuild Time: %s\n", rlib.GetVersionNo(), rlib.GetBuildTime())
		os.Exit(0)
	}
	if *noconPtr {
		rlib.DisableConsole()
	} else {
		rlib.EnableConsole()
	}

	App.AcctDep = *pAD
	App.ARFile = *arPtr
	App.AsmtFile = *asmtPtr
	App.AssignFile = *asgnPtr
	App.BizFile = *bizPtr
	App.BldgFile = *bldgPtr
	App.CoaFile = *coaPtr
	App.CustomFile = *custPtr
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.DepositFile = *depositPtr
	App.DepositoryFile = *depositoryPtr
	App.DMFile = *dmPtr
	App.InvoiceFile = *invPtr
	App.NoteTypeFile = *ntPtr
	App.PetFile = *petPtr
	App.PmtTypeFile = *pmtPtr
	App.PplFile = *pPtr
	App.RaFile = *raPtr
	App.RatFile = *ratPtr
	App.RcptFile = *rcptPtr
	App.Report = *lptr
	App.RFile = *rPtr
	App.RPFile = *rpptr
	App.RPRefFile = *rprptr
	App.RPRRTRateFile = *rprrt
	App.RPRSPRateFile = *rprsp
	App.RSpFile = *rspPtr
	App.RspRefsFile = *rsrefsPtr
	App.RTFile = *rtPtr
	App.SLFile = *slPtr
	App.SrcFile = *src
	App.VehicleFile = *vehiclePtr
	App.NoAuth = *noauth

	var err error
	s := *pDates
	if len(s) > 0 {
		ss := strings.Split(s, ",")
		App.DtStart, err = rlib.StringToDate(ss[0])
		if err != nil {
			fmt.Printf("Invalid start date:  %s\n", ss[0])
			os.Exit(1)
		}
		App.DtStop, err = rlib.StringToDate(ss[1])
		if err != nil {
			fmt.Printf("Invalid stop date:  %s\n", ss[1])
			os.Exit(1)
		}
	} else if len(App.AsmtFile) > 0 || len(App.RcptFile) > 0 {
		fmt.Printf("To load Assessments or Receipts you must specify a time range.\n")
		os.Exit(1)
	}

	App.BUD = strings.TrimSpace(*pBUD)
}

func bizErrCheck(sa []string) {
	if len(sa) < 2 {
		fmt.Printf("Company Designation is required to list Rental Agreements\n")
		os.Exit(1)
	}
}

func loaderGetBiz(ctx context.Context, s string) int64 {
	bid, err := rcsv.GetBusinessBID(ctx, s)
	if /*bid == 0*/ err != nil {
		fmt.Printf("unrecognized Business designator: %s\n", s)
		os.Exit(1)
	}
	return bid
}

type csvimporter struct {
	Name    string
	CmdOpt  string
	Handler func(string) []error
}

func rrDoLoad(ctx context.Context, fname string, handler rcsv.CSVLoadHandlerFunc) {
	// fmt.Printf("calling handler for: %q\n", fname)
	m := handler(ctx, fname)
	fmt.Print(rcsv.ErrlistToString(&m))
}

func main() {
	readCommandLineArgs()
	var err error

	//----------------------------
	// Open RentRoll database
	//----------------------------
	if err = rlib.RRReadConfig(); err != nil {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	s := extres.GetSQLOpenString(rlib.AppConfig.RRDbname, &rlib.AppConfig)
	// rlib.Console("Attempting  sql.Open for database=%s, dbuser=%s\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser)
	// rlib.Console("Open string = %s\n", s)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	// rlib.Console("err = nil for sql.Open\n")
	defer App.dbrr.Close()
	// rlib.Console("Attempting to Ping App.dbrr\n")
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("App.dbrr.Ping for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	// rlib.Console("err = nil for App.dbrr.Ping()\n")

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
	rlib.InitDBHelpers(App.dbrr, App.dbdir)
	rlib.SetNoAuthFlag(App.NoAuth) // currently needed for testing
	rlib.SessionInit(10)           // must be called before calling InitBizInternals

	// create background context
	var now = time.Now()
	var ctx = context.Background()
	if !App.NoAuth {
		// rlib.Console("Creating session for %s, designator = %s\n", rlib.BotReg[rlib.CSVLoaderApp].Name, rlib.BotReg[rlib.CSVLoaderApp].Designator)
		expire := now.Add(10 * time.Minute)
		s := rlib.SessionNew("CSVLoader-app"+fmt.Sprintf("%010x", expire.Unix()),
			rlib.BotReg[rlib.CSVLoaderApp].Designator,
			rlib.BotReg[rlib.CSVLoaderApp].Designator,
			rlib.CSVLoaderApp, "", -1, &expire)
		ctx = rlib.SetSessionContextKey(ctx, s)
	}

	//----------------------------------------------------
	// initialize the CSV infrastructure
	//----------------------------------------------------
	if len(App.BUD) > 0 {
		b2, err := rlib.GetBusinessByDesignation(ctx, App.BUD)
		if err != nil {
			fmt.Printf("Could not find Business Unit named %s, Error=%s\n", App.BUD, err.Error())
			os.Exit(1)
		}
		if b2.BID == 0 {
			// If resource not found then also raise the Error
			fmt.Printf("Could not find Business Unit named %s\n", App.BUD)
			os.Exit(1)
		}

		err = rlib.GetXBusiness(ctx, b2.BID, &App.Xbiz)
		if err != nil {
			fmt.Printf("Could not load Business with BID(%d), Error=%s\n", b2.BID, err.Error())
			os.Exit(1)
		}
		if b2.BID == 0 {
			// If resource not found then also raise the Error
			fmt.Printf("Could not load Business with BID(%d)\n", b2.BID)
			os.Exit(1)
		}
	} else if len(App.AsmtFile) > 0 || len(App.RcptFile) > 0 {
		fmt.Printf("To load Assessments or Receipts you must provide a business unit\n")
		os.Exit(1)
	}
	if App.Xbiz.P.BID > 0 {
		rcsv.InitRCSV(&App.DtStart, &App.DtStop, &App.Xbiz)
		err = rlib.InitBizInternals(App.Xbiz.P.BID, &App.Xbiz)
		if err != nil /*b2.BID == 0*/ {
			fmt.Printf("error while InitBizInternals call: %s\n", err.Error())
			os.Exit(1)
		}
	}

	if len(App.RcptFile) > 0 && App.Xbiz.P.BID > 0 {
		App.PmtTypes, err = rlib.GetPaymentTypesByBusiness(ctx, App.Xbiz.P.BID)
		if err != nil {
			fmt.Printf("error while get payment types for BID(%d): %s\n", App.Xbiz.P.BID, err.Error())
			os.Exit(1)
		}
	}

	//----------------------------------------------------
	// Do all the file loading
	//----------------------------------------------------
	var h = []rcsv.CSVLoadHandler{
		{Fname: App.BizFile, Handler: rcsv.LoadBusinessCSV},
		{Fname: App.SLFile, Handler: rcsv.LoadStringTablesCSV},
		{Fname: App.PmtTypeFile, Handler: rcsv.LoadPaymentTypesCSV},
		{Fname: App.DMFile, Handler: rcsv.LoadDepositMethodsCSV},
		{Fname: App.SrcFile, Handler: rcsv.LoadSourcesCSV},
		{Fname: App.RTFile, Handler: rcsv.LoadRentableTypesCSV},
		{Fname: App.CustomFile, Handler: rcsv.LoadCustomAttributesCSV},
		{Fname: App.DepositoryFile, Handler: rcsv.LoadDepositoryCSV},
		{Fname: App.RSpFile, Handler: rcsv.LoadRentalSpecialtiesCSV},
		{Fname: App.BldgFile, Handler: rcsv.LoadBuildingCSV},
		{Fname: App.PplFile, Handler: rcsv.LoadPeopleCSV},
		{Fname: App.VehicleFile, Handler: rcsv.LoadVehicleCSV},
		{Fname: App.RFile, Handler: rcsv.LoadRentablesCSV},
		{Fname: App.RspRefsFile, Handler: rcsv.LoadRentableSpecialtyRefsCSV},
		{Fname: App.RatFile, Handler: rcsv.LoadRentalAgreementTemplatesCSV},
		{Fname: App.RaFile, Handler: rcsv.LoadRentalAgreementCSV},
		{Fname: App.PetFile, Handler: rcsv.LoadPetsCSV},
		{Fname: App.CoaFile, Handler: rcsv.LoadChartOfAccountsCSV},
		{Fname: App.ARFile, Handler: rcsv.LoadARCSV},
		{Fname: App.RPFile, Handler: rcsv.LoadRatePlansCSV},
		{Fname: App.RPRefFile, Handler: rcsv.LoadRatePlanRefsCSV},
		{Fname: App.RPRRTRateFile, Handler: rcsv.LoadRatePlanRefRTRatesCSV},
		{Fname: App.RPRSPRateFile, Handler: rcsv.LoadRatePlanRefSPRatesCSV},
		{Fname: App.AsmtFile, Handler: rcsv.LoadAssessmentsCSV},
		{Fname: App.RcptFile, Handler: rcsv.LoadReceiptsCSV},
		{Fname: App.DepositFile, Handler: rcsv.LoadDepositCSV},
		{Fname: App.AssignFile, Handler: rcsv.LoadCustomAttributeRefsCSV},
		{Fname: App.NoteTypeFile, Handler: rcsv.LoadNoteTypesCSV},
		{Fname: App.InvoiceFile, Handler: rcsv.LoadInvoicesCSV},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			rrDoLoad(ctx, h[i].Fname, h[i].Handler)
		}
	}

	//----------------------------------------------------
	// Now do all the reporting
	//----------------------------------------------------
	var r = []rrpt.ReporterInfo{
		{ReportNo: 3, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: false, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportBusiness},
		{ReportNo: 5, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportRentableTypes},
		{ReportNo: 6, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportRentables},
		{ReportNo: 7, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportPeople},
		{ReportNo: 8, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportRentalAgreementTemplates},
		{ReportNo: 9, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportRentalAgreements},
		{ReportNo: 10, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportChartOfAccounts},
		{ReportNo: 11, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: true, Handler: rrpt.RRreportAssessments},
		{ReportNo: 12, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportPaymentTypes},
		{ReportNo: 13, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: true, Handler: rrpt.RRreportReceipts},
		{ReportNo: 14, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportCustomAttributes},
		{ReportNo: 15, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportCustomAttributeRefs},
		{ReportNo: 16, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: false, NeedsRAID: true, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportPets},
		{ReportNo: 17, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportNoteTypes},
		{ReportNo: 18, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportDepository},
		{ReportNo: 19, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: true, Handler: rrpt.RRreportDeposits},
		{ReportNo: 20, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: true, Handler: rrpt.RRreportInvoices},
		{ReportNo: 21, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportSpecialties},
		{ReportNo: 22, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportSpecialtyAssigns},
		{ReportNo: 23, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportDepositMethods},
		{ReportNo: 24, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportSources},
		{ReportNo: 25, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportStringLists},
		{ReportNo: 26, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportRatePlans},
		{ReportNo: 27, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: true, Handler: rrpt.RRreportRatePlanRefs},
		{ReportNo: 28, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.VehicleReport},
		{ReportNo: 29, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: false, NeedsDt: false, Handler: rrpt.RRreportAR},
		{ReportNo: 30, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: true, NeedsDt: false, Handler: rrpt.RRRcptOnlyReceipt},
		{ReportNo: 31, OutputFormat: gotable.TABLEOUTTEXT, NeedsBID: true, NeedsRAID: false, NeedsID: true, NeedsDt: false, Handler: rrpt.RRRcptOnlyHotelReceipt},
	}

	if len(App.Report) > 0 {
		sa := strings.Split(App.Report, ",")
		rptno, err := strconv.Atoi(sa[0])
		if err != nil {
			fmt.Printf("Invalid report number: %s\n", sa[0])
			os.Exit(1)
		}
		idx := -1
		for j := 0; j < len(r); j++ {
			if r[j].ReportNo == rptno {
				idx = j
				break
			}
		}
		if idx < 0 {
			fmt.Printf("unimplemented report type: %s\n", App.Report)
			os.Exit(1)
		}
		r[idx].D1 = time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC) // init
		r[idx].D2 = time.Date(9999, time.January, 0, 0, 0, 0, 0, time.UTC) // init
		if r[idx].NeedsBID {
			bizErrCheck(sa)
			r[idx].Bid = loaderGetBiz(ctx, sa[1])
			var xbiz rlib.XBusiness
			err = rlib.GetXBusiness(ctx, r[idx].Bid, &xbiz)
			if err != nil {
				fmt.Printf("error while loading business for BID(%d): %s\n", r[idx].Bid, err.Error())
				os.Exit(1)
			}
			r[idx].Xbiz = &xbiz
		}

		if r[idx].NeedsDt {
			r[idx].D1 = App.DtStart
			r[idx].D2 = App.DtStop
		}
		if r[idx].NeedsRAID {
			r[idx].Raid = rcsv.CSVLoaderGetRAID(sa[1])
		}
		if r[idx].NeedsID {
			r[idx].ID = rcsv.CSVLoaderGetRCPTID(sa[2])
		}

		fmt.Printf("%s\n", r[idx].Handler(ctx, &r[idx]))
	}
}
