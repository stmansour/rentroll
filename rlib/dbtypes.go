package rlib

import (
	"database/sql"
	"time"
)

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

	OCCTYPEUNSET     = 0
	OCCTYPEHOURLY    = 1
	OCCTYPEDAILY     = 2
	OCCTYPEWEEKLY    = 3
	OCCTYPEMONTHLY   = 4
	OCCTYPEQUARTERLY = 5
	OCCTYPEYEARLY    = 6

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
//   USPID = unit specialty id
//   OFSID = offset id
//  ASMTID = assessment type id
//   PMTID = payment type id
// AVAILID = availability id
//  BLDGID = building id
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
}

// AgreementRentable describes a rentable associated with a rental agreement
type AgreementRentable struct {
	RAID    int64     // associated rental agreement
	RID     int64     // the rentable
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
	Trn Transactant
	Tnt Tenant
	Psp Prospect
	Pay Payor
}

// AssessmentType describes the different types of assessments
type AssessmentType struct {
	ASMTID      int64
	Name        string
	Description string
	LastModTime time.Time
	LastModBy   int64
}

// Assessment is a charge associated with a rentable
type Assessment struct {
	ASMID           int64
	BID             int64
	RID             int64
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
	Designation          string // reference to designation in Phonebook db
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
	RTID        int64     // which rentable type
	RID         int64     // which ledger keeps track of what's owed on this unit
	AVAILID     int64     // how is the unit made available
	LastModTime time.Time //	-- when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
}

// RentableSpecialty is the structure for attributes of a unit specialty
type RentableSpecialty struct {
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
	Style          string
	Name           string
	Frequency      int64
	Proration      int64
	Report         int64 // does this type of rentable show up in reporting
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

// XBusiness combines the Business struct and a map of the business's unit types
type XBusiness struct {
	P  Business
	RT map[int64]RentableType      // what types of things are rented here
	US map[int64]RentableSpecialty // index = USPID, val = RentableSpecialty
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

// RRprepSQL is a collection of prepared sql statements for the RentRoll db
type RRprepSQL struct {
	GetRentalAgreementByBusiness   *sql.Stmt
	GetRentalAgreement             *sql.Stmt
	GetLedger                      *sql.Stmt
	GetTransactant                 *sql.Stmt
	GetTenant                      *sql.Stmt
	GetRentable                    *sql.Stmt
	GetProspect                    *sql.Stmt
	GetPayor                       *sql.Stmt
	GetRentableSpecialties         *sql.Stmt
	GetRentableSpecialty           *sql.Stmt
	GetRentableType                *sql.Stmt
	GetRentableTypeByName          *sql.Stmt
	InsertRentableType             *sql.Stmt
	GetUnitReceipts                *sql.Stmt
	GetUnitAssessments             *sql.Stmt
	GetAllRentableAssessments      *sql.Stmt
	GetAssessment                  *sql.Stmt
	GetAssessmentType              *sql.Stmt
	GetSecurityDepositAssessment   *sql.Stmt
	GetAllRentablesByBusiness      *sql.Stmt
	GetAllBusinessRentableTypes    *sql.Stmt
	GetRentableMarketRates         *sql.Stmt
	InsertRentableMarketRates      *sql.Stmt
	GetBusiness                    *sql.Stmt
	GetBusinessByDesignation       *sql.Stmt
	GetAllBusinessSpecialtyTypes   *sql.Stmt
	GetAllAssessmentsByBusiness    *sql.Stmt
	GetReceipt                     *sql.Stmt
	GetReceiptsInDateRange         *sql.Stmt
	GetReceiptAllocations          *sql.Stmt
	GetDefaultLedgerMarkers        *sql.Stmt
	GetAllJournalsInRange          *sql.Stmt
	GetJournalAllocations          *sql.Stmt
	GetJournalByRange              *sql.Stmt
	GetJournalMarker               *sql.Stmt
	GetJournalMarkers              *sql.Stmt
	GetJournal                     *sql.Stmt
	GetJournalAllocation           *sql.Stmt
	InsertJournalMarker            *sql.Stmt
	InsertJournal                  *sql.Stmt
	InsertJournalAllocation        *sql.Stmt
	DeleteJournalAllocations       *sql.Stmt
	DeleteJournalEntry             *sql.Stmt
	DeleteJournalMarker            *sql.Stmt
	GetAllLedgersInRange           *sql.Stmt
	GetLedgerMarkers               *sql.Stmt
	GetLedgerMarkerByGLNo          *sql.Stmt
	GetLedgerInRangeByGLNo         *sql.Stmt
	GetLedgerMarkerInitList        *sql.Stmt
	InsertLedgerMarker             *sql.Stmt
	InsertLedger                   *sql.Stmt
	InsertLedgerAllocation         *sql.Stmt
	DeleteLedgerEntry              *sql.Stmt
	DeleteLedgerMarker             *sql.Stmt
	GetAllLedgerMarkersInRange     *sql.Stmt
	GetAgreementRentables          *sql.Stmt
	GetAgreementPayors             *sql.Stmt
	GetAgreementsForRentable       *sql.Stmt
	GetLatestLedgerMarkerByGLNo    *sql.Stmt
	GetLedgerMarkerByGLNoDateRange *sql.Stmt
	InsertBusiness                 *sql.Stmt
	InsertAssessmentType           *sql.Stmt
	GetAssessmentTypeByName        *sql.Stmt
}

// PBprepSQL is the structure of prepared sql statements for the Phonebook db
type PBprepSQL struct {
	GetCompanyByDesignation *sql.Stmt
}

// BusinessTypes is a struct holding a collection of Types associated
type BusinessTypes struct {
	BID          int64
	AsmtTypes    map[int64]*AssessmentType
	PmtTypes     map[int64]*PaymentType
	DefaultAccts map[int64]*LedgerMarker // index by DFAC..., value = GL No of that account
}

// RRdb is a struct with all variables needed by the db infrastructure
var RRdb struct {
	Prepstmt RRprepSQL
	PBsql    PBprepSQL
	dbdir    *sql.DB // phonebook db
	dbrr     *sql.DB //rentroll db
	BizTypes map[int64]*BusinessTypes
}

// InitDBHelpers initializes the db infrastructure
func InitDBHelpers(dbrr, dbdir *sql.DB) {
	RRdb.dbdir = dbdir
	RRdb.dbrr = dbrr
	RRdb.BizTypes = make(map[int64]*BusinessTypes, 0)
	buildPreparedStatements()
}
