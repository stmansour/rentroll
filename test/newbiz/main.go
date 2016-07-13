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
)
import _ "github.com/go-sql-driver/mysql"

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
	RaFile         string                     //rental agreement cvs file
	CoaFile        string                     //chart of accounts
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
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	bizPtr := flag.String("b", "", "add Business via csv file")
	bldgPtr := flag.String("D", "", "add Buildings to a Business via csv file")
	depositoryPtr := flag.String("d", "", "add Depositories to a Business via csv file")
	depositPtr := flag.String("y", "", "add Deposits via csv file")
	rtPtr := flag.String("R", "", "add Rentable types via csv file")
	rPtr := flag.String("r", "", "add rentables via csv file")
	rspPtr := flag.String("s", "", "add Rentable specialties via csv file")
	pPtr := flag.String("p", "", "add people via csv file")
	ratPtr := flag.String("T", "", "add rental agreement templates via csv file")
	raPtr := flag.String("C", "", "add rental agreements via csv file")
	coaPtr := flag.String("c", "", "add chart of accounts via csv file")
	asmtPtr := flag.String("A", "", "add Assessments via csv file")
	pmtPtr := flag.String("P", "", "add payment types via csv file")
	rcptPtr := flag.String("e", "", "add receipts via csv file")
	custPtr := flag.String("u", "", "add custom attributes via csv file")
	asgnPtr := flag.String("U", "", "assign custom attributes via csv file")
	petPtr := flag.String("E", "", "assign pets to a Rental Agreement via csv file")
	rsrefsPtr := flag.String("F", "", "assign rentable specialties to rentables via csv file")
	ntPtr := flag.String("O", "", "add NoteTypes via csv file")
	lptr := flag.String("L", "", "Report: 1-jnl, 2-ldg, 3-biz, 4-asmtypes, 5-rtypes, 6-rentables, 7-people, 8-rat, 9-ra, 10-coa, 11-asm, 12-payment types, 13-receipts, 14-CustAttr, 15-CustAttrRef, 16-Pets, 17-NoteTypes, 18-Depositories, 19-Deposits")

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

	if len(App.BizFile) > 0 {
		rcsv.LoadBusinessCSV(App.BizFile)
	}
	if len(App.PmtTypeFile) > 0 {
		rcsv.LoadPaymentTypesCSV(App.PmtTypeFile)
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
	if len(App.AsmtFile) > 0 {
		rcsv.LoadAssessmentsCSV(App.AsmtFile)
	}
	if len(App.RcptFile) > 0 {
		App.PmtTypes = rlib.GetPaymentTypes()
		rcsv.LoadReceiptsCSV(App.RcptFile, &App.PmtTypes)
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
		default:
			fmt.Printf("unimplemented report type: %s\n", App.Report)
		}
	}
}
