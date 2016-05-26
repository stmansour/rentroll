package rlib

import (
	"database/sql"
	"time"
)

// NO and all the rest are constants that are used with the RentRoll database
const (
	NO  = int64(0) // std negative value
	YES = int64(1)

	RPTTEXT = 0
	RPTHTML = 1

	ELEMPERSON       = 1 // people
	ELEMCOMPANY      = 2 // companies
	ELEMCLASS        = 3 // classes
	ELEMSVC          = 4 // the executable service
	ELEMRENTABLETYPE = 5 // RentableType element
	ELEMLAST         = 5 // keep in sync with last one added

	CUSTSTRING = 0
	CUSTINT    = 1
	CUSTFLOAT  = 2
	CUSTLAST   = 2 // this should be maintained as matching the highest index value in the group

	RENT                      = 1
	SECURITYDEPOSIT           = 2
	SECURITYDEPOSITASSESSMENT = 58

	LMPAYORACCT        = 1 // ledger set up for a payor
	ACCTSTATUSINACTIVE = 1
	ACCTSTATUSACTIVE   = 2
	RAASSOCIATED       = 1
	RAUNASSOCIATED     = 2

	DFLTCASH       = 10
	DFLTGENRCV     = 11
	DFLTGSRENT     = 12
	DFLTLTL        = 13
	DFLTVAC        = 14
	DFLTSECDEPRCV  = 15
	DFLTSECDEPASMT = 16
	DFLTOWNREQUITY = 17
	DFLTLAST       = 17 // set this to the last default account index

	ACCRUALUNSET     = 0
	ACCRUALSECONDLY  = 1
	ACCRUALMINUTELY  = 2
	ACCRUALHOURLY    = 3
	ACCRUALDAILY     = 4
	ACCRUALWEEKLY    = 5
	ACCRUALMONTHLY   = 6
	ACCRUALQUARTERLY = 7
	ACCRUALYEARLY    = 8

	RENTABLESTATEONLINE   = 0
	RENTABLESTATEADMIN    = 1
	RENTABLESTATEEMPLOYEE = 2
	RENTABLESTATEOWNEROCC = 3
	RENTABLESTATEOFFLINE  = 4
	RENTABLESTATELAST     = 4

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

// RRDATEFMT2 is the default date format that excel outputs.  Use this for csv imports
const RRDATEFMT2 = "1/2/06"

// RRDATEFMT3 is another date format that excel outputs.  Use this for csv imports
const RRDATEFMT3 = "1/2/2006"

// RRDATEINPFMT is the shorthand for database-style dates
const RRDATEINPFMT = "2006-01-02"

//==========================================
// ASMID = Assessment id
// ASMTID = assessment type id
// AVAILID = availability id
// BID = business id
// BLDGID = building id
// CID = custom attribute id
// DISBID = disbursement id
// JAID = journal allocation id
// JID = journal id
// JMID = journal marker id
// LID = ledger entry id
// LMID = ledger marker id
// OFSID = offset id
// PID = payor id
// PMTID = payment type id
// PRSPID = Prospect id
// RAID = rental agreement / occupancy agreement
// RATID = occupancy agreement template id
// RCPTID = receipt id
// RID = rentable id
// RSPID = unit specialty id
// RTID = rentable type id
// TCID = transactant id
// TID = tenant id
//==========================================

// CustomAttribute is a struct containing user-defined custom attributes for objects
type CustomAttribute struct {
	CID         int64     // unique id
	Type        int64     // what type of value: 0 = string, 1 = int64, 2 = float64
	Name        string    // what its called
	Value       string    // string value -- will be xlated on load / store
	LastModTime time.Time // timestamp of last changed
	LastModBy   int64     // who changed it last
	fval        float64   // the float value once converted
	ival        int64     // the int value once converted
}

// CustomAttributeRef is a reference to a Custom Attribute. A query of the form:
//		SELECT CID FROM CustomAttributeRef
type CustomAttributeRef struct {
	ElementType int64 // what type of element:  1=person, 2=company, 3=business-unit, 4 = executable service, 5=RentableType
	ID          int64 // the UID of the element type. That is, if ElementType == 5, the ID is the RTID (rentable type id)
	CID         int64 // uid of the custom attribute
}

// RentalAgreementTemplate is a template used to set up new rental agreements
type RentalAgreementTemplate struct {
	RATID               int64
	ReferenceNumber     string // a string associated with each rental type agreement
	RentalAgreementType int64  // 0=unset, 1=leasehold, 2=month-to-month, 3=hotel
	LastModTime         time.Time
	LastModBy           int64
}

// RentalAgreement binds one or more payors to one or more rentables
type RentalAgreement struct {
	RAID              int64       // internal unique id
	RATID             int64       // reference to Occupancy Master Agreement
	BID               int64       // business (so that we can process by business)
	PrimaryTenant     int64       // Tenant ID of primary tenant
	RentalStart       time.Time   // start date for rental
	RentalStop        time.Time   // stop date for rental
	OccStart          time.Time   // start date for Occupancy
	OccStop           time.Time   // stop date for Occupancy
	Renewal           int64       // 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
	SpecialProvisions string      // free-form text
	LastModTime       time.Time   //	-- when was this record last written
	LastModBy         int64       // employee UID (from phonebook) that modified it
	R                 []XRentable // all the rentables
	P                 []XPerson   // all the payors
	T                 []XPerson   // all the tenants
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

// AgreementTenant describes a Tenant associated with a rental agreement
type AgreementTenant struct {
	RAID    int64
	TID     int64
	DtStart time.Time // start date/time for this Tenant
	DtStop  time.Time // stop date/time (when this person stopped being a tenant)
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
	CompanyName    string // sometimes the entity will be a company
	IsCompany      int    // 1 => the entity is a company, 0 = not a company
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
	LastModTime    time.Time
	LastModBy      int64
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
	DateofBirth                time.Time
	EmergencyContactName       string
	EmergencyContactAddress    string
	EmergencyContactTelephone  string
	EmergencyEmail             string
	AlternateAddress           string
	ElibigleForFutureOccupancy int64
	Industry                   string
	Source                     string
	InvoicingCustomerNumber    string
	LastModTime                time.Time
	LastModBy                  int64
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
	EmployerPostalCode    string
	EmployerEmail         string
	EmployerPhone         string
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
	ASMTID       int64
	OccupancyRqd int64
	Name         string
	Description  string
	LastModTime  time.Time
	LastModBy    int64
}

// Assessment is a charge associated with a rentable
type Assessment struct {
	ASMID           int64     // unique id for this assessment
	BID             int64     // what business
	RID             int64     // the rentable
	ASMTID          int64     // what type of assessment
	RAID            int64     // associated Rental Agreement
	Amount          float64   // how much
	Start           time.Time // start time
	Stop            time.Time // stop time, may be the same as start time or later
	Accrual         int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
	ProrationMethod int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
	AcctRule        string    // expression showing how to account for the amount
	Comment         string
	LastModTime     time.Time
	LastModBy       int64
}

// Business is the set of attributes describing a rental or hotel business
type Business struct {
	BID                int64
	Designation        string // reference to designation in Phonebook db
	Name               string
	DefaultAccrual     int64     // may not be default for every rentable: 0=unset, 1=short term, 2=longterm
	ParkingPermitInUse int64     // yes/no  0 = no, 1 = yes
	LastModTime        time.Time // when was this record last written
	LastModBy          int64     // employee UID (from phonebook) that modified it
}

// Building defines the location of a building that is part of a business
type Building struct {
	BLDGID      int64
	BID         int64
	Address     string
	Address2    string
	City        string
	State       string
	PostalCode  string
	Country     string
	LastModTime time.Time
	LastModBy   int
}

// PaymentType describes how a payment was made
type PaymentType struct {
	PMTID       int64
	BID         int64
	Name        string
	Description string
	LastModTime time.Time
	LastModBy   int64
}

// Receipt saves the information associated with a payment made by a tenant to cover one or more assessments
type Receipt struct {
	RCPTID      int64
	BID         int64
	RAID        int64
	PMTID       int64
	Dt          time.Time
	Amount      float64
	AcctRule    string
	Comment     string
	LastModTime time.Time
	LastModBy   int64
	RA          []ReceiptAllocation
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
	RTID           int64     // rentable type id
	BID            int64     // business
	Name           string    // name for this rental
	Assignment     int64     // can we pre-assign or assign only at commencement
	Report         int64     // 1 = apply to rentroll, 0 = skip
	DefaultOccType int64     // 0 =unset, 1 = short term, 2=longterm
	OccType        int64     // 0 =unset, 1 = short term, 2=longterm
	State          int64     // 0 = online, 1 = administrative unit, 2 = owner occupied, 3 = offline
	LastModTime    time.Time // time of last update to the db record
	LastModBy      int64     // who made the update (Phonebook UID)
}

// RentableSpecialty is the structure for attributes of a rentable specialty
type RentableSpecialty struct {
	RSPID       int64
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
	Accrual        int64
	Proration      int64
	Report         int64 // does this type of rentable show up in reporting
	ManageToBudget int64
	MR             []RentableMarketRate
	CA             []CustomAttribute
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

// XBusiness combines the Business struct and a map of the business's rentable types
type XBusiness struct {
	P  Business
	RT map[int64]RentableType      // what types of things are rented here
	US map[int64]RentableSpecialty // index = RSPID, val = RentableSpecialty
}

// XRentable is the structure that includes both the Rentable and Unit attributes
type XRentable struct {
	R       Rentable  // the rentable
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

// LedgerEntry is the structure for Ledger entry attributes
type LedgerEntry struct {
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
	LMID         int64
	BID          int64
	RAID         int64  // only valid if Type == 1
	GLNumber     string // acct system name
	Status       int64  // Whether a GL Account is currently unknown=0, inactive=1, active=2
	State        int64  // 0 = unknown, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
	DtStart      time.Time
	DtStop       time.Time
	Balance      float64
	Type         int64     // flag: 0 = not a default account, 1 = Rental Agreement Account, 10-default cash, 11-GENRCV, 12-GrossSchedRENT, 13-LTL, 14-VAC, ...
	Name         string    // descriptive name for the ledger
	AcctType     string    // Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, Cost of Goods Sold, Other Income, Other Expense
	RAAssociated int64     // 1 = Unassociated with RentalAgreement, 2 = Associated with Rental Agreement, 0 = unknown
	LastModTime  time.Time // auto updated
	LastModBy    int64     // user making the mod
}

// RRprepSQL is a collection of prepared sql statements for the RentRoll db
type RRprepSQL struct {
	DeleteJournalAllocations           *sql.Stmt
	DeleteJournalEntry                 *sql.Stmt
	DeleteJournalMarker                *sql.Stmt
	DeleteLedgerEntry                  *sql.Stmt
	DeleteLedgerMarker                 *sql.Stmt
	DeleteReceipt                      *sql.Stmt
	DeleteReceiptAllocations           *sql.Stmt
	FindTransactantByPhoneOrEmail      *sql.Stmt
	FindAgreementByRentable            *sql.Stmt
	GetAgreementPayors                 *sql.Stmt
	GetAgreementRentables              *sql.Stmt
	GetAgreementTenants                *sql.Stmt
	GetAgreementsForRentable           *sql.Stmt
	GetAllAssessmentsByBusiness        *sql.Stmt
	GetAllBusinessRentableTypes        *sql.Stmt
	GetAllBusinesses                   *sql.Stmt
	GetAllBusinessSpecialtyTypes       *sql.Stmt
	GetAllJournalsInRange              *sql.Stmt
	GetAllLedgerMarkersInRange         *sql.Stmt
	GetAllLedgersInRange               *sql.Stmt
	GetAllRentableAssessments          *sql.Stmt
	GetAllRentablesByBusiness          *sql.Stmt
	GetAllRentalAgreementTemplates     *sql.Stmt
	GetAllRentalAgreements             *sql.Stmt
	GetAllTransactants                 *sql.Stmt
	GetAssessment                      *sql.Stmt
	GetAssessmentType                  *sql.Stmt
	GetAssessmentTypeByName            *sql.Stmt
	GetBuilding                        *sql.Stmt
	GetBusiness                        *sql.Stmt
	GetBusinessByDesignation           *sql.Stmt
	GetDefaultLedgerMarkers            *sql.Stmt
	GetJournal                         *sql.Stmt
	GetJournalAllocation               *sql.Stmt
	GetJournalAllocations              *sql.Stmt
	GetJournalByRange                  *sql.Stmt
	GetJournalMarker                   *sql.Stmt
	GetJournalMarkers                  *sql.Stmt
	GetLatestLedgerMarkerByGLNo        *sql.Stmt
	GetLatestLedgerMarkerByType        *sql.Stmt
	GetLedger                          *sql.Stmt
	GetLedgerInRangeByGLNo             *sql.Stmt
	GetLedgerMarkerByGLNoDateRange     *sql.Stmt
	GetLedgerMarkerInitList            *sql.Stmt
	GetLedgerMarkers                   *sql.Stmt
	GetPaymentTypesByBusiness          *sql.Stmt
	GetPayor                           *sql.Stmt
	GetProspect                        *sql.Stmt
	GetReceipt                         *sql.Stmt
	GetReceiptAllocations              *sql.Stmt
	GetReceiptsInDateRange             *sql.Stmt
	GetRentable                        *sql.Stmt
	GetRentableByName                  *sql.Stmt
	GetRentableMarketRates             *sql.Stmt
	GetRentableSpecialties             *sql.Stmt
	GetRentableSpecialty               *sql.Stmt
	GetRentableType                    *sql.Stmt
	GetRentableTypeByStyle             *sql.Stmt
	GetRentalAgreement                 *sql.Stmt
	GetRentalAgreementByBusiness       *sql.Stmt
	GetRentalAgreementTemplate         *sql.Stmt
	GetRentalAgreementTemplateByRefNum *sql.Stmt
	GetSecurityDepositAssessment       *sql.Stmt
	GetSpecialtyByName                 *sql.Stmt
	GetTenant                          *sql.Stmt
	GetTransactant                     *sql.Stmt
	GetUnitAssessments                 *sql.Stmt
	InsertAgreementPayor               *sql.Stmt
	InsertAgreementRentable            *sql.Stmt
	InsertAgreementTenant              *sql.Stmt
	InsertAssessment                   *sql.Stmt
	InsertAssessmentType               *sql.Stmt
	InsertBuilding                     *sql.Stmt
	InsertBuildingWithID               *sql.Stmt
	InsertBusiness                     *sql.Stmt
	InsertJournal                      *sql.Stmt
	InsertJournalAllocation            *sql.Stmt
	InsertJournalMarker                *sql.Stmt
	InsertLedger                       *sql.Stmt
	InsertLedgerAllocation             *sql.Stmt
	InsertLedgerMarker                 *sql.Stmt
	InsertPayor                        *sql.Stmt
	InsertPaymentType                  *sql.Stmt
	InsertProspect                     *sql.Stmt
	InsertReceipt                      *sql.Stmt
	InsertReceiptAllocation            *sql.Stmt
	InsertRentable                     *sql.Stmt
	InsertRentableMarketRates          *sql.Stmt
	InsertRentableSpecialtyType        *sql.Stmt
	InsertRentableType                 *sql.Stmt
	InsertRentalAgreement              *sql.Stmt
	InsertRentalAgreementTemplate      *sql.Stmt
	InsertTenant                       *sql.Stmt
	InsertTransactant                  *sql.Stmt
	UpdateLedgerMarker                 *sql.Stmt
	UpdateTransactant                  *sql.Stmt
	InsertCustomAttribute              *sql.Stmt
	GetCustomAttribute                 *sql.Stmt
	DeleteCustomAttribute              *sql.Stmt
	InsertCustomAttributeRef           *sql.Stmt
	GetCustomAttributeRefs             *sql.Stmt
	DeleteCustomAttributeRef           *sql.Stmt
}

// PBprepSQL is the structure of prepared sql statements for the Phonebook db
type PBprepSQL struct {
	GetCompanyByDesignation      *sql.Stmt
	GetCompany                   *sql.Stmt
	GetBusinessUnitByDesignation *sql.Stmt
}

// BusinessTypes is a struct holding a collection of Types associated
type BusinessTypes struct {
	BID          int64
	AsmtTypes    map[int64]*AssessmentType
	PmtTypes     map[int64]*PaymentType
	DefaultAccts map[int64]*LedgerMarker // index by the predifined contants DFAC*, value = GL No of that account
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
	buildPBPreparedStatements()
}

// InitBusinessFields initialize the lists in rlib's internal data structures
func InitBusinessFields(bid int64) {
	if nil == RRdb.BizTypes[bid] {
		bt := BusinessTypes{
			BID:          bid,
			AsmtTypes:    make(map[int64]*AssessmentType),
			PmtTypes:     make(map[int64]*PaymentType),
			DefaultAccts: make(map[int64]*LedgerMarker),
		}
		RRdb.BizTypes[bid] = &bt
	}
}
