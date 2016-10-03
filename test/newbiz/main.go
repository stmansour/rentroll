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
	"database/sql"
	"flag"
	"fmt"
	"os"
	"rentroll/rcsv"
	"rentroll/rlib"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir          *sql.DB                    // phonebook db
	dbrr           *sql.DB                    //rentroll db
	DBDir          string                     // phonebook database
	DBRR           string                     //rentroll database
	DBUser         string                     // user for all databases
	PmtTypes       map[int64]rlib.PaymentType // all payment types in this db
	Report         string                     // Report: 1 = Journal, 2 = Ledger, 3 = Businesses, 4 = Rentable types
	BizFile        string                     // name of csv file with new biz info
	RTFile         string                     // Rentable types csv file
	RFile          string                     // rentables csv file
	RSpFile        string                     // Rentable specialties
	BldgFile       string                     // Buildings for this Business
	PplFile        string                     // people for this Business
	RatFile        string                     // rentalAgreementTemplates
	RaFile         string                     // rental agreement cvs file
	CoaFile        string                     // chart of accounts
	AsmtFile       string                     // Assessments
	PmtTypeFile    string                     // payment types
	RcptFile       string                     // receipts of payments
	CustomFile     string                     // custom attributes
	AssignFile     string                     // assign custom attributes
	PetFile        string                     // assign pets
	RspRefsFile    string                     // assign specialties to rentables
	NoteTypeFile   string                     // note types
	DepositoryFile string                     // Depository
	DepositFile    string                     // Deposits
	InvoiceFile    string                     // Invoice
	DMFile         string                     // Deposit Methods
	SrcFile        string                     // Sources
	SLFile         string                     // StringLists
	RPFile         string                     // RatePlans
	RPRefFile      string                     // RatePlanRef
	RPRRTRateFile  string                     // RatePlanRefRTRate
	RPRSPRateFile  string                     // RatePlanRefSPRate
	BUD            string                     // business unit designator
	DtStart        time.Time                  // range start time
	DtStop         time.Time                  // range stop time
	Xbiz           rlib.XBusiness             // xbusiness associated with -G  (BUD)
}

func readCommandLineArgs() {
	asmtPtr := flag.String("A", "", "add Assessments via csv file")
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
	pDates := flag.String("g", "", "Date Rage.  Example: 1/1/16,2/1/16")
	pBUD := flag.String("G", "", "BUD - business unit designator")
	invPtr := flag.String("i", "", "add Invoices via csv file")
	lptr := flag.String("L", "", "Report: 1-jnl, 2-ldg, 3-biz, 4-asmtypes, 5-rtypes, 6-rentables, 7-people, 8-rat, 9-ra, 10-coa, 11-asm, 12-payment types, 13-receipts, 14-CustAttr, 15-CustAttrRef, 16-Pets, 17-NoteTypes, 18-Depositories, 19-Deposits, 20-Invoices, 21-Specialties, 22-Specialty Assignments, 23-Deposit Methods, 24-Sources, 25-StringList, 26-RatePlan, 27-RatePlanRef,BUD,RatePlanName")
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
	src := flag.String("S", "", "add Rentable specialties via csv file")
	rspPtr := flag.String("s", "", "add Rentable specialties via csv file")
	ratPtr := flag.String("T", "", "add rental agreement templates via csv file")
	rprsp := flag.String("t", "", "add RatePlanRef Specialty Rates via csv file")
	asgnPtr := flag.String("U", "", "assign custom attributes via csv file")
	custPtr := flag.String("u", "", "add custom attributes via csv file")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	depositPtr := flag.String("y", "", "add Deposits via csv file")

	flag.Parse()
	if *verPtr {
		fmt.Printf("Version:    %s\nBuild Time: %s\n", getVersionNo(), getBuildTime())
		os.Exit(0)
	}
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.BizFile = *bizPtr
	App.AsmtFile = *asmtPtr
	App.RTFile = *rtPtr
	App.RSpFile = *rspPtr
	App.BldgFile = *bldgPtr
	App.RFile = *rPtr
	App.PplFile = *pPtr
	App.RatFile = *ratPtr
	App.RaFile = *raPtr
	App.CoaFile = *coaPtr
	App.PmtTypeFile = *pmtPtr
	App.RcptFile = *rcptPtr
	App.Report = *lptr
	App.CustomFile = *custPtr
	App.AssignFile = *asgnPtr
	App.PetFile = *petPtr
	App.RspRefsFile = *rsrefsPtr
	App.NoteTypeFile = *ntPtr
	App.DepositoryFile = *depositoryPtr
	App.DepositFile = *depositPtr
	App.InvoiceFile = *invPtr
	App.DMFile = *dmPtr
	App.SrcFile = *src
	App.SLFile = *slPtr
	App.RPFile = *rpptr
	App.RPRefFile = *rprptr
	App.RPRRTRateFile = *rprrt
	App.RPRSPRateFile = *rprsp
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

func loaderGetBiz(s string) int64 {
	bid := rcsv.GetBusinessBID(s)
	if bid == 0 {
		fmt.Printf("unrecognized Business designator: %s\n", s)
		os.Exit(1)
	}
	return bid
}

func main() {
	readCommandLineArgs()
	rlib.RRReadConfig()

	var err error

	//----------------------------
	// Open RentRoll database
	//----------------------------
	// s := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", DBUser, DBRR)
	s := rlib.RRGetSQLOpenString(App.DBRR)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBRR, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBRR, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = rlib.RRGetSQLOpenString(App.DBDir)
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

	//----------------------------------------------------
	// initialize the CSV infrastructure
	//----------------------------------------------------
	if len(App.BUD) > 0 {
		b2 := rlib.GetBusinessByDesignation(App.BUD)
		if b2.BID == 0 {
			fmt.Printf("Could not find Business Unit named %s\n", App.BUD)
			os.Exit(1)
		}
		rlib.GetXBusiness(b2.BID, &App.Xbiz)
	} else if len(App.AsmtFile) > 0 || len(App.RcptFile) > 0 {
		fmt.Printf("To load Assessments or Receipts you must provide a business unit\n")
		os.Exit(1)
	}
	if App.Xbiz.P.BID > 0 {
		rcsv.InitRCSV(&App.DtStart, &App.DtStop, &App.Xbiz)
	}

	//----------------------------------------------------
	// Now, on with the main portion of the program...
	//----------------------------------------------------
	if len(App.BizFile) > 0 {
		rcsv.LoadBusinessCSV(App.BizFile)
	}
	if len(App.SLFile) > 0 {
		rcsv.LoadStringTablesCSV(App.SLFile)
	}
	if len(App.PmtTypeFile) > 0 {
		rcsv.LoadPaymentTypesCSV(App.PmtTypeFile)
	}
	if len(App.DMFile) > 0 {
		rcsv.LoadDepositMethodsCSV(App.DMFile)
	}
	if len(App.SrcFile) > 0 {
		rcsv.LoadSourcesCSV(App.SrcFile)
	}
	if len(App.RTFile) > 0 {
		rcsv.LoadRentableTypesCSV(App.RTFile)
	}
	if len(App.CustomFile) > 0 {
		rcsv.LoadCustomAttributesCSV(App.CustomFile)
	}
	if len(App.DepositoryFile) > 0 {
		rcsv.LoadDepositoryCSV(App.DepositoryFile)
	}
	if len(App.RSpFile) > 0 {
		rcsv.LoadRentalSpecialtiesCSV(App.RSpFile)
	}
	if len(App.BldgFile) > 0 {
		rcsv.LoadBuildingCSV(App.BldgFile)
	}
	if len(App.PplFile) > 0 {
		rcsv.LoadPeopleCSV(App.PplFile)
	}
	if len(App.RFile) > 0 {
		rcsv.LoadRentablesCSV(App.RFile)
	}
	if len(App.RspRefsFile) > 0 {
		rcsv.LoadRentableSpecialtyRefsCSV(App.RspRefsFile)
	}
	if len(App.RatFile) > 0 {
		rcsv.LoadRentalAgreementTemplatesCSV(App.RatFile)
	}
	if len(App.RaFile) > 0 {
		rcsv.LoadRentalAgreementCSV(App.RaFile)
	}
	if len(App.PetFile) > 0 {
		rcsv.LoadPetsCSV(App.PetFile)
	}
	if len(App.CoaFile) > 0 {
		rcsv.LoadChartOfAccountsCSV(App.CoaFile)
	}
	if len(App.RPFile) > 0 {
		rcsv.LoadRatePlansCSV(App.RPFile)
	}
	if len(App.RPRefFile) > 0 {
		rcsv.LoadRatePlanRefsCSV(App.RPRefFile)
	}
	if len(App.RPRRTRateFile) > 0 {
		rcsv.LoadRatePlanRefRTRatesCSV(App.RPRRTRateFile)
	}
	if len(App.RPRSPRateFile) > 0 {
		rcsv.LoadRatePlanRefSPRatesCSV(App.RPRSPRateFile)
	}
	if len(App.AsmtFile) > 0 {
		s := rcsv.LoadAssessmentsCSV(App.AsmtFile)
		fmt.Print(s)
	}
	if len(App.RcptFile) > 0 {
		App.PmtTypes = rlib.GetPaymentTypes()
		s := rcsv.LoadReceiptsCSV(App.RcptFile)
		fmt.Print(s)
	}
	if len(App.DepositFile) > 0 {
		rcsv.LoadDepositCSV(App.DepositFile)
	}
	if len(App.AssignFile) > 0 {
		rcsv.LoadCustomAttributeRefsCSV(App.AssignFile)
	}
	if len(App.NoteTypeFile) > 0 {
		rcsv.LoadNoteTypesCSV(App.NoteTypeFile)
	}
	if len(App.InvoiceFile) > 0 {
		rcsv.LoadInvoicesCSV(App.InvoiceFile)
	}

	if len(App.Report) > 0 {
		sa := strings.Split(App.Report, ",")
		i := int64(0)
		if len(sa) > 0 {
			i, _ = rlib.IntFromString(sa[0], "report number")
		}
		switch i {
		case 1:
			fmt.Printf("1 - not yet implemented\n")
		case 2:
			fmt.Printf("2 - not yet implemented\n")
		case 3:
			fmt.Printf("%s\n", rcsv.RRreportBusiness(rlib.RPTTEXT))
		case 5:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportRentableTypes(rlib.RPTTEXT, bid))
		case 6:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportRentables(rlib.RPTTEXT, bid))
		case 7:
			fmt.Printf("%s\n", rcsv.RRreportPeople(rlib.RPTTEXT))
		case 8:
			fmt.Printf("%s\n", rcsv.RRreportRentalAgreementTemplates(rlib.RPTTEXT))
		case 9:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportRentalAgreements(rlib.RPTTEXT, bid))
		case 10:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportChartOfAccounts(rlib.RPTTEXT, bid))
		case 11:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportAssessments(rlib.RPTTEXT, bid))
		case 12:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportPaymentTypes(rlib.RPTTEXT, bid))
		case 13:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportReceipts(rlib.RPTTEXT, bid))
		case 14:
			fmt.Printf("%s\n", rcsv.RRreportCustomAttributes(rlib.RPTTEXT))
		case 15:
			fmt.Printf("%s\n", rcsv.RRreportCustomAttributeRefs(rlib.RPTTEXT))
		case 16:
			raid := rcsv.CSVLoaderGetRAID(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportRentalAgreementPets(rlib.RPTTEXT, raid))
		case 17:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportNoteTypes(rlib.RPTTEXT, bid))
		case 18:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportDepository(rlib.RPTTEXT, bid))
		case 19:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportDeposits(rlib.RPTTEXT, bid))
		case 20:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportInvoices(rlib.RPTTEXT, bid))
		case 21:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportSpecialties(rlib.RPTTEXT, bid))
		case 22:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportSpecialtyAssigns(rlib.RPTTEXT, bid))
		case 23:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportDepositMethods(rlib.RPTTEXT, bid))
		case 24:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportSources(rlib.RPTTEXT, bid))
		case 25:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportStringLists(rlib.RPTTEXT, bid))
		case 26:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportRatePlans(rlib.RPTTEXT, bid))
		case 27:
			bizErrCheck(sa)
			dt := time.Now()
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportRatePlanRefs(rlib.RPTTEXT, bid, &dt, &dt))
		default:
			fmt.Printf("unimplemented report type: %s\n", App.Report)
		}
	}
}
