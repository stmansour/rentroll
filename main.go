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
//   OATID = occupancy agreement template id
//    RAID = occupancy agreement
//  RCPTID = receipt id
//  DISBID = disbursement id
//     LID = ledger id
//==========================================

// RentalAgreement binds a tenant to a unit
type RentalAgreement struct {
	RAID                  int
	OATID                 int
	PRID                  int
	UNITID                int
	PID                   int
	PrimaryTenant         int
	RentalStart           time.Time
	RentalStop            time.Time
	Renewal               int
	ProrationMethod       int
	ScheduledRent         float32
	Frequency             int
	SecurityDepositAmount float32
	SpecialProvisions     string
	LastModTime           time.Time
	LastModBy             int
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

// Assessment is a charge associated with a unit
type Assessment struct {
	ASMID     int
	UNITID    int
	ASMTID    int
	Amount    float32
	Start     time.Time
	Stop      time.Time
	Frequency int
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
	DefaultOccupancyType int       // default for every unit in the building: 0=unset, 1=daily, 2=weekly, 3=monthly, 4=quarterly, 5=yearly
	ParkingPermitInUse   int       // yes/no  0 = no, 1 = yes
	LastModTime          time.Time // when was this record last written
	LastModBy            int       // employee UID (from phonebook) that modified it
}

// Receipt saves the information associated with a payment made by a tenant to cover one or more assessments
type Receipt struct {
	RCPTID                   int
	PID                      int
	RAID                     int
	Dt                       time.Time
	Amount                   float32
	ApplyToGeneralReceivable float32
	ApplyToSecurityDeposit   float32
}

// Rentable is the basic struct for  entities to rent
type Rentable struct {
	RID         int    // unique id for this rentable
	LID         int    // the ledger
	RTID        int    // rentable type id
	PRID        int    // property
	PID         int    // payor
	RAID        int    // occupancy agreement
	UNITID      int    // associated unit (if applicable, 0 otherwise)
	Name        string // name for this rental
	Assignment  int    // can we pre-assign or assign only at commencement
	Report      int    // 1 = apply to rentroll, 0 = skip
	LastModTime time.Time
	LastModBy   int
}

// Unit is the structure for unit attributes
type Unit struct {
	UNITID         int       // unique id for this unit -- it is unique across all properties and buildings
	BLDGID         int       // which building
	UTID           int       // which unit type
	RID            int       // which ledger keeps track of what's owed on this unit
	AVAILID        int       // how is the unit made available
	DefaultOccType int       // unset, short term, longterm
	OccType        int       // unset, short term, longterm
	LastModTime    time.Time //	-- when was this record last written
	LastModBy      int       // employee UID (from phonebook) that modified it
}

// UnitSpecialtyType is the structure for attributes of a unit specialty
type UnitSpecialtyType struct {
	USPID       int
	PRID        int
	Name        string
	Fee         float32
	Description string
}

// UnitType is the structure for attributes of a unit type
type UnitType struct {
	UTID        int
	PRID        int
	Style       string
	Name        string
	SqFt        int
	MarketRate  float32
	LastModTime time.Time
	LastModBy   int
}

// XUnit is the structure that includes both the Rentable and Unit attributes
type XUnit struct {
	R Rentable
	U Unit
}

// Ledger is the structure for Ledger attributes
type Ledger struct {
	LID       int       // unique id for this Ledger
	AccountNo string    // if not '' then it's a link a QB account
	Dt        time.Time // balance date and time
	Balance   float32   // balance amount
	Deposit   float32   // deposit balance
}

// collection of prepared sql statements
type prepSQL struct {
	occAgrByProperty     *sql.Stmt
	getUnit              *sql.Stmt
	getLedger            *sql.Stmt
	getTransactant       *sql.Stmt
	getTenant            *sql.Stmt
	getRentable          *sql.Stmt
	getProspect          *sql.Stmt
	getPayor             *sql.Stmt
	getRentalAgreement   *sql.Stmt
	getUnitSpecialties   *sql.Stmt
	getUnitSpecialtyType *sql.Stmt
	getUnitType          *sql.Stmt
	getUnitReceipts      *sql.Stmt
	getUnitAssessments   *sql.Stmt
}

// App is the global data structure for this app
var App struct {
	dbdir    *sql.DB
	dbrr     *sql.DB
	DBDir    string
	DBRR     string
	DBUser   string
	prepstmt prepSQL
	asmt2int map[string]int
	asmt2str map[int]string
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	flag.Parse()
	if *verPtr {
		fmt.Printf("Version: %s\nBuilt:   %s\n", getVersionNo(), getBuildTime())
		os.Exit(0)
	}
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
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
	start := time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	RentRollAll(start, stop)
}
