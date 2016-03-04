package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

//==========================================
//    PRID = property id
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

// RentalAgreement binds a teRAID INT NOT NULL
type RentalAgreement struct {
	RAID              int
	RATID             int
	PRID              int
	RID               int
	UNITID            int
	PID               int
	PrimaryTenant     int
	RentalStart       time.Time
	RentalStop        time.Time
	Renewal           int
	SpecialProvisions string
	LastModTime       time.Time
	LastModBy         int
}

// Transactant is the basic structure of information
// about a person who is a prospect, applicant, tenant, or payor
type Transactant struct {
	TCID           int
	TID            int
	PID            int
	PRSPID         int
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
	LastModBy      int
}

// Prospect contains info over and above
type Prospect struct {
	PRSPID         int
	TCID           int
	ApplicationFee float32 // if non-zero this prospect is an applicant
}

// Tenant contains all info common to a person
type Tenant struct {
	TID                        int
	TCID                       int
	Points                     int
	CarMake                    string
	CarModel                   string
	CarColor                   string
	CarYear                    int
	LicensePlateState          string
	LicensePlateNumber         string
	ParkingPermitNumber        string
	AccountRep                 int
	DateofBirth                string
	EmergencyContactName       string
	EmergencyContactAddress    string
	EmergencyContactTelephone  string
	EmergencyAddressEmail      string
	AlternateAddress           string
	ElibigleForFutureOccupancy int
	Industry                   string
	Source                     string
	InvoicingCustomerNumber    string
}

// Payor is attributes of the person financially responsible
// for the rent.
type Payor struct {
	PID                   int
	TCID                  int
	CreditLimit           float32
	EmployerName          string
	EmployerStreetAddress string
	EmployerCity          string
	EmployerState         string
	EmployerZipcode       string
	Occupation            string
	LastModTime           time.Time
	LastModBy             int
}

// XPerson of all person related attributes
type XPerson struct {
	trn Transactant
	tnt Tenant
	psp Prospect
	pay Payor
}

// assessment types
const (
	RENT                      = 1
	SECURITYDEPOSIT           = 2
	SECURITYDEPOSITASSESSMENT = 58

	CREDIT = 0
	DEBIT  = 1

	RTHOUSING = 1
	RTCARPORT = 2
	RTCAR     = 3

	REPORTJUSTIFYLEFT  = 0
	REPORTJUSTIFYRIGHT = 1
)

// AssessmentType describes the different types of assessments
type AssessmentType struct {
	ASMTID      int
	Name        string
	Type        int // 0 = credit, 1 = debit
	LastModTime time.Time
	LastModBy   int
}

// Assessment is a charge associated with a rentable
type Assessment struct {
	ASMID           int
	RID             int
	UNITID          int
	ASMTID          int
	RAID            int
	Amount          float32
	Start           time.Time
	Stop            time.Time
	Frequency       int
	ProrationMethod int
	LastModTime     time.Time
	LastModBy       int
}

// Property is the set of attributes describing a rental or hotel property
type Property struct {
	PRID                 int
	Address              string
	Address2             string
	City                 string
	State                string
	PostalCode           string
	Country              string
	Phone                string
	Name                 string
	DefaultOccupancyType int       // may not be default for every unit in the building: 0=unset, 1=short term, 2=longterm
	ParkingPermitInUse   int       // yes/no  0 = no, 1 = yes
	LastModTime          time.Time // when was this record last written
	LastModBy            int       // employee UID (from phonebook) that modified it
}

// PaymentType describes how a payment was made
type PaymentType struct {
	PMTID       int
	Name        string
	Description string
	LastModTime time.Time
	LastModBy   int
}

// Receipt saves the information associated with a payment made by a tenant to cover one or more assessments
type Receipt struct {
	RCPTID                   int
	PID                      int
	RAID                     int
	PMTID                    int
	Dt                       time.Time
	Amount                   float32
	ApplyToGeneralReceivable float32
	ApplyToSecurityDeposit   float32
}

// Rentable is the basic struct for  entities to rent
type Rentable struct {
	RID            int    // unique id for this rentable
	LID            int    // the ledger
	RTID           int    // rentable type id
	PRID           int    // property
	PID            int    // payor
	RAID           int    // occupancy agreement
	UNITID         int    // associated unit (if applicable, 0 otherwise)
	Name           string // name for this rental
	Assignment     int    // can we pre-assign or assign only at commencement
	Report         int    // 1 = apply to rentroll, 0 = skip
	DefaultOccType int    // unset, short term, longterm
	OccType        int    // unset, short term, longterm
	LastModTime    time.Time
	LastModBy      int
}

// Unit is the structure for unit attributes
type Unit struct {
	UNITID      int       // unique id for this unit -- it is unique across all properties and buildings
	BLDGID      int       // which building
	UTID        int       // which unit type
	RID         int       // which ledger keeps track of what's owed on this unit
	AVAILID     int       // how is the unit made available
	LastModTime time.Time //	-- when was this record last written
	LastModBy   int       // employee UID (from phonebook) that modified it
}

// UnitSpecialtyType is the structure for attributes of a unit specialty
type UnitSpecialtyType struct {
	USPID       int
	PRID        int
	Name        string
	Fee         float32
	Description string
}

// RentableType is the set of attributes describing the different types of rentable items
type RentableType struct {
	RTID        int
	PRID        int
	Name        string
	MarketRate  float32
	Frequency   int
	Proration   int
	LastModTime time.Time
	LastModBy   int
}

// UnitType is the set of attributes describing the different types of housing within a property
type UnitType struct {
	UTID        int
	PRID        int
	Style       string
	Name        string
	SqFt        int
	MarketRate  float32
	Frequency   int
	Proration   int
	LastModTime time.Time
	LastModBy   int
}

// XProperty combines the Property struct and a map of the property's unit types
type XProperty struct {
	P  Property
	RT map[int]RentableType      // what types of things are rented here
	UT map[int]UnitType          // info about the units
	US map[int]UnitSpecialtyType // index = USPID, val = UnitSpecialtyType
}

// XUnit is the structure that includes both the Rentable and Unit attributes
type XUnit struct {
	R Rentable
	U Unit
	S []int
}

// Ledger is the structure for Ledger attributes
type Ledger struct {
	LID      int       // unique id for this Ledger
	GLNumber string    // if not '' then it's a link a QB account
	Dt       time.Time // balance date and time
	Balance  float32   // balance amount
	// Deposit   float32   // deposit balance
}

// collection of prepared sql statements
type prepSQL struct {
	occAgrByProperty             *sql.Stmt
	getUnit                      *sql.Stmt
	getLedger                    *sql.Stmt
	getTransactant               *sql.Stmt
	getTenant                    *sql.Stmt
	getRentable                  *sql.Stmt
	getProspect                  *sql.Stmt
	getPayor                     *sql.Stmt
	getRentalAgreement           *sql.Stmt
	getUnitSpecialties           *sql.Stmt
	getUnitSpecialtyType         *sql.Stmt
	getRentableType              *sql.Stmt
	getUnitType                  *sql.Stmt
	getUnitReceipts              *sql.Stmt
	getUnitAssessments           *sql.Stmt
	getAllRentableAssessments    *sql.Stmt
	getAssessmentType            *sql.Stmt
	getSecurityDepositAssessment *sql.Stmt
	getUnitRentalAgreements      *sql.Stmt
	getAllRentablesByProperty    *sql.Stmt
	getAllPropertyRentableTypes  *sql.Stmt
	getAllPropertyUnitTypes      *sql.Stmt
	getProperty                  *sql.Stmt
	getAllPropertySpecialtyTypes *sql.Stmt
}

// App is the global data structure for this app
var App struct {
	dbdir     *sql.DB
	dbrr      *sql.DB
	DBDir     string
	DBRR      string
	DBUser    string
	prepstmt  prepSQL
	AsmtTypes map[int]AssessmentType
	PmtTypes  map[int]PaymentType
	Report    int
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	rptPtr := flag.Int("r", 1, "report: 1 = journal, 2 = rentable")
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

func main() {
	readCommandLineArgs()

	var err error
	s := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", App.DBUser, App.DBDir)
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
	buildPreparedStatements()
	initLists()

	//  func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	start := time.Date(2015, time.November, 1, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC)
	switch App.Report {
	case 1:
		JournalReport(start, stop)
	case 2:
		RentRollAll(start, stop)
	default:
		fmt.Printf("Unknown report type: %d\n", App.Report)
	}
}
