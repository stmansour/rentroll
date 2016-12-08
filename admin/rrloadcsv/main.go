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
	"rentroll/rrpt"
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
	VehicleFile    string                     // vehicles that belong to people
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
	pDates := flag.String("g", "", "Date Range.  Example: 1/1/16,2/1/16")
	pBUD := flag.String("G", "", "BUD - business unit designator")
	invPtr := flag.String("i", "", "add Invoices via csv file")
	lptr := flag.String("L", "", "Report: 1-jnl, 2-ldg, 3-biz, 4-asmtypes, 5-rtypes, 6-rentables, 7-people, 8-rat, 9-ra, 10-coa, 11-asm, 12-payment types, 13-receipts, 14-CustAttr, 15-CustAttrRef, 16-Pets, 17-NoteTypes, 18-Depositories, 19-Deposits, 20-Invoices, 21-Specialties, 22-Specialty Assignments, 23-Deposit Methods, 24-Sources, 25-StringList, 26-RatePlan, 27-RatePlanRef,BUD,RatePlanName, 28-BUD")
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
	App.VehicleFile = *vehiclePtr
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

type csvimporter struct {
	Name    string
	CmdOpt  string
	Handler func(string) string
}

func rrDoLoad(fname string, handler func(string) []error) {
	m := handler(fname)
	fmt.Print(rcsv.ErrlistToString(&m))
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
		rrDoLoad(App.BizFile, rcsv.LoadBusinessCSV)
	}
	if len(App.SLFile) > 0 {
		rrDoLoad(App.SLFile, rcsv.LoadStringTablesCSV)
	}
	if len(App.PmtTypeFile) > 0 {
		rrDoLoad(App.PmtTypeFile, rcsv.LoadPaymentTypesCSV)
	}
	if len(App.DMFile) > 0 {
		rrDoLoad(App.DMFile, rcsv.LoadDepositMethodsCSV)
	}
	if len(App.SrcFile) > 0 {
		rrDoLoad(App.SrcFile, rcsv.LoadSourcesCSV)
	}
	if len(App.RTFile) > 0 {
		rrDoLoad(App.RTFile, rcsv.LoadRentableTypesCSV)
	}
	if len(App.CustomFile) > 0 {
		rrDoLoad(App.CustomFile, rcsv.LoadCustomAttributesCSV)
	}
	if len(App.DepositoryFile) > 0 {
		rrDoLoad(App.DepositoryFile, rcsv.LoadDepositoryCSV)
	}
	if len(App.RSpFile) > 0 {
		rrDoLoad(App.RSpFile, rcsv.LoadRentalSpecialtiesCSV)
	}
	if len(App.BldgFile) > 0 {
		rrDoLoad(App.BldgFile, rcsv.LoadBuildingCSV)
	}
	if len(App.PplFile) > 0 {
		rrDoLoad(App.PplFile, rcsv.LoadPeopleCSV)
	}
	if len(App.VehicleFile) > 0 {
		rrDoLoad(App.VehicleFile, rcsv.LoadVehicleCSV)
	}
	if len(App.RFile) > 0 {
		rrDoLoad(App.RFile, rcsv.LoadRentablesCSV)
	}
	if len(App.RspRefsFile) > 0 {
		rrDoLoad(App.RspRefsFile, rcsv.LoadRentableSpecialtyRefsCSV)
	}
	if len(App.RatFile) > 0 {
		rrDoLoad(App.RatFile, rcsv.LoadRentalAgreementTemplatesCSV)
	}
	if len(App.RaFile) > 0 {
		rrDoLoad(App.RaFile, rcsv.LoadRentalAgreementCSV)
	}
	if len(App.PetFile) > 0 {
		rrDoLoad(App.PetFile, rcsv.LoadPetsCSV)
	}
	if len(App.CoaFile) > 0 {
		rrDoLoad(App.CoaFile, rcsv.LoadChartOfAccountsCSV)
	}
	if len(App.RPFile) > 0 {
		rrDoLoad(App.RPFile, rcsv.LoadRatePlansCSV)
	}
	if len(App.RPRefFile) > 0 {
		rrDoLoad(App.RPRefFile, rcsv.LoadRatePlanRefsCSV)
	}
	if len(App.RPRRTRateFile) > 0 {
		rrDoLoad(App.RPRRTRateFile, rcsv.LoadRatePlanRefRTRatesCSV)
	}
	if len(App.RPRSPRateFile) > 0 {
		rrDoLoad(App.RPRSPRateFile, rcsv.LoadRatePlanRefSPRatesCSV)
	}
	if len(App.AsmtFile) > 0 {
		rrDoLoad(App.AsmtFile, rcsv.LoadAssessmentsCSV)
	}
	if len(App.RcptFile) > 0 {
		App.PmtTypes = rlib.GetPaymentTypes()
		rrDoLoad(App.RcptFile, rcsv.LoadReceiptsCSV)
	}
	if len(App.DepositFile) > 0 {
		rrDoLoad(App.DepositFile, rcsv.LoadDepositCSV)
	}
	if len(App.AssignFile) > 0 {
		rrDoLoad(App.AssignFile, rcsv.LoadCustomAttributeRefsCSV)
	}
	if len(App.NoteTypeFile) > 0 {
		rrDoLoad(App.NoteTypeFile, rcsv.LoadNoteTypesCSV)
	}
	if len(App.InvoiceFile) > 0 {
		rrDoLoad(App.InvoiceFile, rcsv.LoadInvoicesCSV)
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
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportPeople(rlib.RPTTEXT, bid))
		case 8:
			bid := loaderGetBiz(sa[1])
			fmt.Printf("%s\n", rcsv.RRreportRentalAgreementTemplates(rlib.RPTTEXT, bid))
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
		case 28:
			bizErrCheck(sa)
			bid := loaderGetBiz(sa[1])
			tbl := rrpt.VehicleReport(bid)
			fmt.Printf("%s\n", tbl.SprintTable(rlib.TABLEOUTTEXT))
		default:
			fmt.Printf("unimplemented report type: %s\n", App.Report)
		}
	}
}
