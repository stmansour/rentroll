package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"phonebook/lib"
	"rentroll/rlib"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

// assessment types
const (
	RENT                      = 1
	SECURITYDEPOSIT           = 2
	SECURITYDEPOSITASSESSMENT = 58

	LMPAYORACCT    = 1 // ledger set up for a payor
	DFLTCASH       = 10
	DFLTGENRCV     = 11
	DFLTGSRENT     = 12
	DFLTLTL        = 13
	DFLTVAC        = 14
	DFLTSECDEPRCV  = 15
	DFLTSECDEPASMT = 16

	CREDIT = 0
	DEBIT  = 1

	RTRESIDENCE = 1
	RTCARPORT   = 2
	RTCAR       = 3

	REPORTJUSTIFYLEFT  = 0
	REPORTJUSTIFYRIGHT = 1

	JNLTYPEUNAS = 0 // record is unassociated with any assessment or receipt
	JNLTYPEASMT = 1 // record is the result of an assessment
	JNLTYPERCPT = 2 // record is the result of a receipt

	MARKERSTATEOPEN   = 0 // Journal/Ledger Marker state
	MARKERSTATECLOSED = 1
	MARKERSTATELOCKED = 2
	MARKERSTATEORIGIN = 3

	JOURNALTYPEASMID  = 1
	JOURNALTYPERCPTID = 2
)

// RRDATEFMT is a shorthand date format used for text output
// Use these values:	Mon Jan 2 15:04:05 MST 2006
// const RRDATEFMT = "02-Jan-2006 3:04PM MST"
// const RRDATEFMT = "01/02/06 3:04PM MST"
const RRDATEFMT = "01/02/06"

//==========================================
//    BID = business id
//    UTID = unit type id
//   USPID = unit specialty id
//   OFSID = offset id
//  ASMTID = assessment type id
//   PMTID = payment type id
// AVAILID = availability id
//  BLDGID = building id
//  UNITID = unit id
//    TCID = transactant id
//     TID = tenant id
//     PID = payor id
//   RATID = rental agreement template id
//    RAID = occupancy agreement
//  RCPTID = receipt id
//  DISBID = disbursement id
//     LID = ledger id
//==========================================

// RentalAgreement binds one or more payors to one or more rentables
type RentalAgreement struct {
	RAID              int64     // internal unique id
	RATID             int64     // reference to Occupancy Master Agreement
	BID               int64     // business (so that we can process by business)
	PrimaryTenant     int64     // Tenant ID of primary tenant
	RentalStart       time.Time // start date for rental
	RentalStop        time.Time // stop date for rental
	Renewal           int64     // month to month automatic renewal, lease extension options, none.
	SpecialProvisions string    // free-form text
	LastModTime       time.Time //	-- when was this record last written
	LastModBy         int64     // employee UID (from phonebook) that modified it
	R                 []XUnit   // everything about the rentable
	P                 []XPerson // everything about the payor
	// RID               int64
	// UNITID            int64
	// PID               int64
	// LID               int64
}

// AgreementRentable describes a rentable associated with a rental agreement
type AgreementRentable struct {
	RAID    int64     // associated rental agreement
	RID     int64     // the rentable
	UNITID  int64     // unit (if applicable, 0 otherwise)
	DtStart time.Time // start date/time for this rentable
	DtStop  time.Time // stop date/time
}

// AgreementPayor describes a payor associated with a rental agreement
type AgreementPayor struct {
	RAID    int64
	PID     int64
	DtStart time.Time // start date/time for this payor
	DtStop  time.Time // stop date/time
}

// Transactant is the basic structure of information
// about a person who is a prospect, applicant, tenant, or payor
type Transactant struct {
	TCID           int64
	TID            int64
	PID            int64
	PRSPID         int64
	FirstName      string
	MiddleName     string
	LastName       string
	PrimaryEmail   string
	SecondaryEmail string
	WorkPhone      string
	CellPhone      string
	Address        string
	Address2       string
	City           string
	State          string
	PostalCode     string
	Country        string
	LastModTime    time.Time
	LastModBy      int64
}

// Prospect contains info over and above
type Prospect struct {
	PRSPID         int64
	TCID           int64
	ApplicationFee float64 // if non-zero this prospect is an applicant
}

// Tenant contains all info common to a person
type Tenant struct {
	TID                        int64
	TCID                       int64
	Points                     int64
	CarMake                    string
	CarModel                   string
	CarColor                   string
	CarYear                    int64
	LicensePlateState          string
	LicensePlateNumber         string
	ParkingPermitNumber        string
	AccountRep                 int64
	DateofBirth                string
	EmergencyContactName       string
	EmergencyContactAddress    string
	EmergencyContactTelephone  string
	EmergencyAddressEmail      string
	AlternateAddress           string
	ElibigleForFutureOccupancy int64
	Industry                   string
	Source                     string
	InvoicingCustomerNumber    string
}

// Payor is attributes of the person financially responsible
// for the rent.
type Payor struct {
	PID                   int64
	TCID                  int64
	CreditLimit           float64
	EmployerName          string
	EmployerStreetAddress string
	EmployerCity          string
	EmployerState         string
	EmployerZipcode       string
	Occupation            string
	LastModTime           time.Time
	LastModBy             int64
}

// XPerson of all person related attributes
type XPerson struct {
	trn Transactant
	tnt Tenant
	psp Prospect
	pay Payor
}

// AssessmentType describes the different types of assessments
type AssessmentType struct {
	ASMTID      int64
	Name        string
	Type        int64 // 0 = credit, 1 = debit
	LastModTime time.Time
	LastModBy   int64
}

// Assessment is a charge associated with a rentable
type Assessment struct {
	ASMID           int64
	BID             int64
	RID             int64
	UNITID          int64
	ASMTID          int64
	RAID            int64
	Amount          float64
	Start           time.Time
	Stop            time.Time
	Frequency       int64
	ProrationMethod int64
	AcctRule        string
	Comment         string
	LastModTime     time.Time
	LastModBy       int64
}

// Business is the set of attributes describing a rental or hotel business
type Business struct {
	BID                  int64
	Address              string
	Address2             string
	City                 string
	State                string
	PostalCode           string
	Country              string
	Phone                string
	Name                 string
	DefaultOccupancyType int64     // may not be default for every unit in the building: 0=unset, 1=short term, 2=longterm
	ParkingPermitInUse   int64     // yes/no  0 = no, 1 = yes
	LastModTime          time.Time // when was this record last written
	LastModBy            int64     // employee UID (from phonebook) that modified it
}

// PaymentType describes how a payment was made
type PaymentType struct {
	PMTID       int64
	Name        string
	Description string
	LastModTime time.Time
	LastModBy   int64
}

// Receipt saves the information associated with a payment made by a tenant to cover one or more assessments
type Receipt struct {
	RCPTID   int64
	BID      int64
	RAID     int64
	PMTID    int64
	Dt       time.Time
	Amount   float64
	AcctRule string
	Comment  string
	RA       []ReceiptAllocation
}

// ReceiptAllocation defines an allocation of a receipt amount.
type ReceiptAllocation struct {
	RCPTID   int64
	Amount   float64
	ASMID    int64
	AcctRule string
}

// Rentable is the basic struct for  entities to rent
type Rentable struct {
	RID            int64     // unique id for this rentable
	LID            int64     // the ledger
	RTID           int64     // rentable type id
	BID            int64     // business
	UNITID         int64     // associated unit (if applicable, 0 otherwise)
	Name           string    // name for this rental
	Assignment     int64     // can we pre-assign or assign only at commencement
	Report         int64     // 1 = apply to rentroll, 0 = skip
	DefaultOccType int64     // 0 =unset, 1 = short term, 2=longterm
	OccType        int64     // 0 =unset, 1 = short term, 2=longterm
	LastModTime    time.Time // time of last update to the db record
	LastModBy      int64     // who made the update (Phonebook UID)
}

// Unit is the structure for unit attributes
type Unit struct {
	UNITID      int64     // unique id for this unit -- it is unique across all properties and buildings
	BLDGID      int64     // which building
	UTID        int64     // which unit type
	RID         int64     // which ledger keeps track of what's owed on this unit
	AVAILID     int64     // how is the unit made available
	LastModTime time.Time //	-- when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
}

// UnitSpecialtyType is the structure for attributes of a unit specialty
type UnitSpecialtyType struct {
	USPID       int64
	BID         int64
	Name        string
	Fee         float64
	Description string
}

// RentableType is the set of attributes describing the different types of rentable items
type RentableType struct {
	RTID           int64
	BID            int64
	Name           string
	Frequency      int64
	Proration      int64
	Report         int64
	ManageToBudget int64
	MR             []RentableMarketRate
	MRCurrent      float64 // the current market rate (historical values are in MR)
	LastModTime    time.Time
	LastModBy      int64
}

// RentableMarketRate describes the market rate rent for a rentable type over a time period
type RentableMarketRate struct {
	RTID       int64
	MarketRate float64
	DtStart    time.Time
	DtStop     time.Time
}

// UnitType is the set of attributes describing the different types of housing within a business
type UnitType struct {
	UTID        int64
	BID         int64
	Style       string
	Name        string
	SqFt        int64
	Frequency   int64
	Proration   int64
	MR          []UnitMarketRate
	MRCurrent   float64 // the current market rate (historical values are in MR)
	LastModTime time.Time
	LastModBy   int64
}

// UnitMarketRate describes the market rate rent for a unit type over a time period
type UnitMarketRate struct {
	UTID       int64
	MarketRate float64
	DtStart    time.Time
	DtStop     time.Time
}

// XType combines RentableType and UnitType
type XType struct {
	RT RentableType
	UT UnitType
}

// XBusiness combines the Business struct and a map of the business's unit types
type XBusiness struct {
	P  Business
	RT map[int64]RentableType      // what types of things are rented here
	UT map[int64]UnitType          // info about the units
	US map[int64]UnitSpecialtyType // index = USPID, val = UnitSpecialtyType
}

// XUnit is the structure that includes both the Rentable and Unit attributes
type XUnit struct {
	R       Rentable  // the rentable
	U       Unit      // unit (if applicable)
	S       []int64   // list of specialties associated with the rentable
	DtStart time.Time // Start date/time for this rentable (associated with the Rental Agreement, but may have different dates)
	DtStop  time.Time // Stop time for this rentable
}

// Journal is the set of attributes describing a journal entry
type Journal struct {
	JID         int64               // unique id for this journal entry
	BID         int64               // unique id of business
	RAID        int64               // unique id of Rental Agreement
	Dt          time.Time           // when this entry was made
	Amount      float64             // the amount
	Type        int64               // 1 means this is an assessment, 2 means it is a payment
	ID          int64               // if Type == 1 then it is the ASMID that caused this entry, of Type ==2 then it is the RCPTID
	Comment     string              // for notes like "prior period adjustment"
	LastModTime time.Time           // auto updated
	LastModBy   int64               // user making the mod
	JA          []JournalAllocation // an array of journal allocations, breaks the payment or assessment down, total of all the allocations equals the "Amount" above
}

// JournalAllocation describes how the associated journal amount is allocated
type JournalAllocation struct {
	JAID     int64   // unique id for this allocation
	JID      int64   // associated journal entry
	RID      int64   // associated rentable
	Amount   float64 // amount of this allocation
	ASMID    int64   // associated AssessmentID -- source of the charge/payment
	AcctRule string  // describes how this amount distributed across the accounts
}

// JournalMarker describes a period of time where the journal entries have been locked down
type JournalMarker struct {
	JMID    int64
	BID     int64
	State   int64
	DtStart time.Time
	DtStop  time.Time
}

// Ledger is the structure for Ledger attributes
type Ledger struct {
	LID         int64
	BID         int64
	JID         int64
	JAID        int64
	GLNumber    string
	Dt          time.Time
	Amount      float64
	Comment     string    // for notes like "prior period adjustment"
	LastModTime time.Time // auto updated
	LastModBy   int64     // user making the mod
}

// LedgerMarker describes a period of time period described. The Balance can be
// used going forward from DtStop
type LedgerMarker struct {
	LMID     int64
	BID      int64
	PID      int64 // only valid if Type == 1
	GLNumber string
	Status   int64
	State    int64
	DtStart  time.Time
	DtStop   time.Time
	Balance  float64
	Type     int64
	Name     string
}

// collection of prepared sql statements
type prepSQL struct {
	rentalAgreementByBusiness      *sql.Stmt
	getRentalAgreement             *sql.Stmt
	getUnit                        *sql.Stmt
	getLedger                      *sql.Stmt
	getTransactant                 *sql.Stmt
	getTenant                      *sql.Stmt
	getRentable                    *sql.Stmt
	getProspect                    *sql.Stmt
	getPayor                       *sql.Stmt
	getUnitSpecialties             *sql.Stmt
	getUnitSpecialtyType           *sql.Stmt
	getRentableType                *sql.Stmt
	getUnitType                    *sql.Stmt
	getXType                       *sql.Stmt
	getUnitReceipts                *sql.Stmt
	getUnitAssessments             *sql.Stmt
	getAllRentableAssessments      *sql.Stmt
	getAssessment                  *sql.Stmt
	getAssessmentType              *sql.Stmt
	getSecurityDepositAssessment   *sql.Stmt
	getAllRentablesByBusiness      *sql.Stmt
	getAllBusinessRentableTypes    *sql.Stmt
	getRentableMarketRates         *sql.Stmt
	getAllBusinessUnitTypes        *sql.Stmt
	getUnitMarketRates             *sql.Stmt
	getBusiness                    *sql.Stmt
	getAllBusinessSpecialtyTypes   *sql.Stmt
	getAllAssessmentsByBusiness    *sql.Stmt
	getReceipt                     *sql.Stmt
	getReceiptsInDateRange         *sql.Stmt
	getReceiptAllocations          *sql.Stmt
	getDefaultLedgerMarkers        *sql.Stmt
	getAllJournalsInRange          *sql.Stmt
	getJournalAllocations          *sql.Stmt
	getJournalByRange              *sql.Stmt
	getJournalMarker               *sql.Stmt
	getJournalMarkers              *sql.Stmt
	getJournal                     *sql.Stmt
	getJournalAllocation           *sql.Stmt
	insertJournalMarker            *sql.Stmt
	insertJournal                  *sql.Stmt
	insertJournalAllocation        *sql.Stmt
	deleteJournalAllocations       *sql.Stmt
	deleteJournalEntry             *sql.Stmt
	deleteJournalMarker            *sql.Stmt
	getAllLedgersInRange           *sql.Stmt
	getLedgerMarkers               *sql.Stmt
	getLedgerMarkerByGLNo          *sql.Stmt
	getLedgerInRangeByGLNo         *sql.Stmt
	getLedgerMarkerInitList        *sql.Stmt
	insertLedgerMarker             *sql.Stmt
	insertLedger                   *sql.Stmt
	insertLedgerAllocation         *sql.Stmt
	deleteLedgerEntry              *sql.Stmt
	deleteLedgerMarker             *sql.Stmt
	getAllLedgerMarkersInRange     *sql.Stmt
	getAgreementRentables          *sql.Stmt
	getAgreementPayors             *sql.Stmt
	getAgreementsForRentable       *sql.Stmt
	getLatestLedgerMarkerByGLNo    *sql.Stmt
	getLedgerMarkerByGLNoDateRange *sql.Stmt
	// getUnitRentalAgreements      *sql.Stmt
}

// BusinessTypes is a struct holding a collection of Types associated
type BusinessTypes struct {
	BID          int64
	AsmtTypes    map[int64]*AssessmentType
	PmtTypes     map[int64]*PaymentType
	DefaultAccts map[int64]*LedgerMarker // index by DFAC..., value = GL No of that account
}

// App is the global data structure for this app
var App struct {
	dbdir     *sql.DB // phonebook db
	dbrr      *sql.DB //rentroll db
	DBDir     string  // phonebook database
	DBRR      string  //rentroll database
	DBUser    string  // user for all databases
	prepstmt  prepSQL //prepared sql for rentroll
	Report    int64   // if testing engine, which report/action to perform
	AsmtTypes map[int64]AssessmentType
	PmtTypes  map[int64]PaymentType
	BizTypes  map[int64]*BusinessTypes
}

// This is Phonebooks's standard logger
func ulog(format string, a ...interface{}) {
	p := fmt.Sprintf(format, a...)
	log.Print(p)
}

// GetRecurrences is a shorthand for assessment variables to get a list
// of dates on which charges must be assessed for a particular interval of time (d1 - d2)
func (a *Assessment) GetRecurrences(d1, d2 *time.Time) []time.Time {
	return rlib.GetRecurrences(d1, d2, &a.Start, &a.Stop, a.Frequency)
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	rptPtr := flag.Int64("r", 0, "report: 0 = generate journal records, 1 = journal, 2 = rentable")
	flag.Parse()
	if *verPtr {
		fmt.Printf("Version: %s\nBuilt:   %s\n", getVersionNo(), getBuildTime())
		os.Exit(0)
	}
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.Report = *rptPtr
}

func intTest(xbiz *XBusiness, d1, d2 *time.Time) {
	fmt.Printf("INTERNAL TEST\n")
	m := parseAcctRule(xbiz, 1, d1, d2, "d ${DFLTGENRCV} 1000.0, c 40001 ${UMR}, d 41004 ${UMR} ${aval(${DFLTGENRCV})} -", float64(1000), float64(8)/float64(30))

	for i := 0; i < len(m); i++ {
		fmt.Printf("m[%d] = %#v\n", i, m[i])
	}
	fmt.Printf("DONE\n")
}

// Dispatch generates the supplied report for all properties
func Dispatch(d1, d2 time.Time, report int64) {
	s := "SELECT BID,Address,Address2,City,State,PostalCode,Country,Phone,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var xbiz XBusiness
		rlib.Errcheck(rows.Scan(&xbiz.P.BID, &xbiz.P.Address, &xbiz.P.Address2, &xbiz.P.City, &xbiz.P.State,
			&xbiz.P.PostalCode, &xbiz.P.Country, &xbiz.P.Phone, &xbiz.P.Name, &xbiz.P.DefaultOccupancyType,
			&xbiz.P.ParkingPermitInUse, &xbiz.P.LastModTime, &xbiz.P.LastModBy))
		GetXBusiness(xbiz.P.BID, &xbiz)
		if nil == App.BizTypes[xbiz.P.BID] {
			bt := BusinessTypes{
				BID:          xbiz.P.BID,
				AsmtTypes:    make(map[int64]*AssessmentType),
				PmtTypes:     make(map[int64]*PaymentType),
				DefaultAccts: make(map[int64]*LedgerMarker),
			}
			App.BizTypes[xbiz.P.BID] = &bt
		}
		GetDefaultLedgerMarkers(xbiz.P.BID)
		// fmt.Printf("Dispatch: After call to GetDefaultLedgerMarkers: App.BizTypes[%d].DefaultAccts = %#v\n", xbiz.P.BID, App.BizTypes[xbiz.P.BID].DefaultAccts)

		switch report {
		case 1:
			JournalReportText(&xbiz, &d1, &d2)
		case 2:
			LedgerReportText(&xbiz, &d1, &d2)
		case 3:
			intTest(&xbiz, &d1, &d2)
		default:
			// fmt.Printf("Generating Journal Records for %s through %s\n", d1.Format(RRDATEFMT), d2.AddDate(0, 0, -1).Format(RRDATEFMT))
			GenerateJournalRecords(&xbiz, &d1, &d2)
			GenerateLedgerRecords(&xbiz, &d1, &d2)
		}
	}
}

func main() {
	readCommandLineArgs()

	var err error
	// s := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", App.DBUser, App.DBDir)
	s := rlib.RRGetSQLOpenString(App.DBUser, App.DBRR)

	App.dbdir, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBDir, App.DBUser, err)
	}
	defer App.dbdir.Close()
	err = App.dbdir.Ping()
	if nil != err {
		fmt.Printf("App.DBDir.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBDir, App.DBUser, err)
	}

	s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", App.DBUser, App.DBRR)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBRR, App.DBUser, err)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("App.DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBRR, App.DBUser, err)
	}
	s = lib.GetSQLOpenString(App.DBUser, App.DBDir)
	App.dbdir, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open: Error = %v\n", err)
	}
	defer App.dbdir.Close()
	err = App.dbdir.Ping()
	if nil != err {
		fmt.Printf("App.dbdir.Ping: Error = %v\n", err)
	}

	initRentRoll()

	//  func Date(year int64 , month Month, day, hour, min, sec, nsec int64 , loc *Location) Time
	start := time.Date(2015, time.November, 1, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC)
	Dispatch(start, stop, App.Report)
}
