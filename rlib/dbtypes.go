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

	RABALANCEACCOUNT = 1 // GLAccount set up for a RentalAgreement balance

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

	ACCRUALNORECUR   = 0
	ACCRUALSECONDLY  = 1
	ACCRUALMINUTELY  = 2
	ACCRUALHOURLY    = 3
	ACCRUALDAILY     = 4
	ACCRUALWEEKLY    = 5
	ACCRUALMONTHLY   = 6
	ACCRUALQUARTERLY = 7
	ACCRUALYEARLY    = 8

	YEARFOREVER = 9000 // an arbitrary year, anything >= to this year is taken to mean "unbounded", or no end date.

	RARQDINRANGE = 0 // assessment must during the Rental Agreement period
	RARQDPRIOR   = 1 // can be assessed prior to or during the Rental Agreement  period
	RARQDAFTER   = 2 // can be assessed during  or after theRental Agreement period
	RARQDANY     = 3 // can be assessed anytime: before, during, or after the Rental Agreement period
	RARQDLAST    = 3 // keep in sync with last

	RENTABLESTATUSUNKNOWN  = 0
	RENTABLESTATUSONLINE   = 1
	RENTABLESTATUSADMIN    = 2
	RENTABLESTATUSEMPLOYEE = 3
	RENTABLESTATUSOWNEROCC = 4
	RENTABLESTATUSOFFLINE  = 5
	RENTABLESTATUSLAST     = 5 // keep in sync with last

	CREDIT = 0
	DEBIT  = 1

	RTRESIDENCE = 1
	RTCARPORT   = 2
	RTCAR       = 3

	REPORTJUSTIFYLEFT  = 0
	REPORTJUSTIFYRIGHT = 1

	JNLTYPEUNAS = 0 // record is unassociated with any assessment or Receipt
	JNLTYPEASMT = 1 // record is the result of an assessment
	JNLTYPERCPT = 2 // record is the result of a Receipt

	MARKERSTATEOPEN   = 0 // Journal/LedgerMarker state
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

// RRDATETIMEINPFMT is the shorthand for database-style dates
const RRDATETIMEINPFMT = "2006-01-02 15:04:00 MST"

//==========================================
// ASMID = Assessment id
// ASMTID = assessment type id
// AVAILID = availability id
// BID = Business id
// BLDGID = Building id
// CID = custom attribute id
// DISBID = disbursement id
// JAID = Journal allocation id
// JID = Journal id
// JMID = Journal marker id
// LEID = LedgerEntry id
// LMID = LedgerMarker id
// OFSID = offset id
// PID = Payor id
// PMTID = payment type id
// PRSPID = Prospect id
// RAID = rental agreement / occupancy agreement
// RATID = occupancy agreement template id
// RCPTID = Receipt id
// USERID = User id
// RID = Rentable id
// RSPID = unit specialty id
// RTID = Rentable type id
// TCID = Transactant id
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
	ElementType int64 // what type of element:  1=person, 2=company, 3=Business-unit, 4 = executable service, 5=RentableType
	ID          int64 // the UID of the element type. That is, if ElementType == 5, the ID is the RTID (Rentable type id)
	CID         int64 // uid of the custom attribute
}

// RentalAgreementTemplate is a template used to set up new rental agreements
type RentalAgreementTemplate struct {
	RATID                int64
	BID                  int64
	RentalTemplateNumber string // RentalTemplateNumber a string associated with each rental type agreement (essentially, the doc name)
	LastModTime          time.Time
	LastModBy            int64
}

// RentalAgreement binds one or more payors to one or more rentables
type RentalAgreement struct {
	RAID              int64       // internal unique id
	RATID             int64       // reference to Occupancy Master Agreement
	BID               int64       // Business (so that we can process by Business)
	RentalStart       time.Time   // start date for rental
	RentalStop        time.Time   // stop date for rental
	PossessionStart   time.Time   // start date for Occupancy
	PossessionStop    time.Time   // stop date for Occupancy
	Renewal           int64       // 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
	SpecialProvisions string      // free-form text
	LastModTime       time.Time   //	-- when was this record last written
	LastModBy         int64       // employee UID (from phonebook) that modified it
	R                 []XRentable // all the rentables
	P                 []XPerson   // all the payors
	T                 []XPerson   // all the renters
}

// RentalAgreementRentable describes a Rentable associated with a rental agreement
type RentalAgreementRentable struct {
	RAID         int64     // associated rental agreement
	RID          int64     // the Rentable
	ContractRent float64   // the rent
	DtStart      time.Time // start date/time for this Rentable
	DtStop       time.Time // stop date/time
}

// RentalAgreementPayor describes a Payor associated with a rental agreement
type RentalAgreementPayor struct {
	RAID    int64
	PID     int64
	DtStart time.Time // start date/time for this Payor
	DtStop  time.Time // stop date/time
}

// RentableUser describes a User associated with a rental agreement
type RentableUser struct {
	RID     int64     // associated Rentable
	USERID  int64     // pointer to Transactant
	DtStart time.Time // start date/time for this User
	DtStop  time.Time // stop date/time (when this person stopped being a User)
}

// RentalAgreementPet describes a pet associated with a rental agreement. There can be as many as needed.
type RentalAgreementPet struct {
	PETID       int64
	RAID        int64
	Type        string
	Breed       string
	Color       string
	Weight      float64
	Name        string
	DtStart     time.Time
	DtStop      time.Time
	LastModTime time.Time
	LastModBy   int64
}

// Transactant is the basic structure of information
// about a person who is a Prospect, applicant, User, or Payor
type Transactant struct {
	TCID           int64
	USERID         int64
	PID            int64
	PRSPID         int64
	FirstName      string
	MiddleName     string
	LastName       string
	PreferredName  string
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
	Website        string // person's website
	Notes          string // general text
	LastModTime    time.Time
	LastModBy      int64
}

// Prospect contains info over and above
type Prospect struct {
	PRSPID                int64
	TCID                  int64
	EmployerName          string
	EmployerStreetAddress string
	EmployerCity          string
	EmployerState         string
	EmployerPostalCode    string
	EmployerEmail         string
	EmployerPhone         string
	Occupation            string
	ApplicationFee        float64 // if non-zero this Prospect is an applicant
	LastModTime           time.Time
	LastModBy             int64
}

// User contains all info common to a person
type User struct {
	USERID                    int64
	TCID                      int64
	Points                    int64
	CarMake                   string
	CarModel                  string
	CarColor                  string
	CarYear                   int64
	LicensePlateState         string
	LicensePlateNumber        string
	ParkingPermitNumber       string
	DateofBirth               time.Time
	EmergencyContactName      string
	EmergencyContactAddress   string
	EmergencyContactTelephone string
	EmergencyEmail            string
	AlternateAddress          string
	EligibleFutureUser        int64
	Industry                  string
	Source                    string
	LastModTime               time.Time
	LastModBy                 int64
}

// Payor is attributes of the person financially responsible
// for the rent.
type Payor struct {
	PID                 int64
	TCID                int64
	CreditLimit         float64
	TaxpayorID          string
	AccountRep          int64
	EligibleFuturePayor int64
	LastModTime         time.Time
	LastModBy           int64
}

// XPerson of all person related attributes
type XPerson struct {
	Trn Transactant
	Tnt User
	Psp Prospect
	Pay Payor
}

// AssessmentType describes the different types of Assessments
type AssessmentType struct {
	ASMTID         int64
	RARequired     int64 // 0 = during rental period, 1 = valid prior or during, 2 = valid during or after, 3 = valid before, during, and after
	ManageToBudget int64 //  0 = do not manage to budget; no ContractRent amount required. 1 = Manage to budget, ContractRent required.
	Name           string
	Description    string
	LastModTime    time.Time
	LastModBy      int64
}

// Assessment is a charge associated with a Rentable
type Assessment struct {
	ASMID          int64     // unique id for this assessment
	BID            int64     // what Business
	RID            int64     // the Rentable
	ASMTID         int64     // what type of assessment
	RAID           int64     // associated Rental Agreement
	Amount         float64   // how much
	Start          time.Time // start time
	Stop           time.Time // stop time, may be the same as start time or later
	RecurCycle     int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
	ProrationCycle int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
	AcctRule       string    // expression showing how to account for the amount
	Comment        string
	LastModTime    time.Time
	LastModBy      int64
}

// Business is the set of attributes describing a rental or hotel Business
type Business struct {
	BID                 int64
	Designation         string // reference to designation in Phonebook db
	Name                string
	DefaultRentalPeriod int64     // may not be default for every Rentable: 0=unset, 1=short term, 2=longterm
	ParkingPermitInUse  int64     // yes/no  0 = no, 1 = yes
	LastModTime         time.Time // when was this record last written
	LastModBy           int64     // employee UID (from phonebook) that modified it
}

// Building defines the location of a Building that is part of a Business
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

// Receipt saves the information associated with a payment made by a User to cover one or more Assessments
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

// ReceiptAllocation defines an allocation of a Receipt amount.
type ReceiptAllocation struct {
	RCPTID   int64
	Amount   float64
	ASMID    int64
	AcctRule string
}

// RentableSpecialtyType is the structure for attributes of a Rentable specialty
type RentableSpecialtyType struct {
	RSPID       int64
	BID         int64
	Name        string
	Fee         float64 // proration inherited from the rentable / rentable type.
	Description string
}

// RentableType is the set of attributes describing the different types of Rentable items
type RentableType struct {
	RTID           int64                // unique identifier for this RentableType
	BID            int64                // the business unit to which this RentableType belongs
	Style          string               // a short name
	Name           string               // longer name
	RentCycle      int64                // frequency at which rent accrues, 0 = not set or n/a, 1 = secondly, 2=minutely, 3=hourly, 4=daily, 5=weekly, 6=monthly...
	Proration      int64                // frequency for prorating rent if the full rentcycle is not used
	GSPRC          int64                // Time increments in which GSR is calculated to account for rate changes
	ManageToBudget int64                // 0=no, 1 = yes
	MR             []RentableMarketRate // array of time sensitive market rates
	CA             []CustomAttribute    // associated custom attributes
	MRCurrent      float64              // the current market rate (historical values are in MR)
	LastModTime    time.Time
	LastModBy      int64
}

// RentableMarketRate describes the market rate rent for a Rentable type over a time period
type RentableMarketRate struct {
	RTID       int64
	MarketRate float64
	DtStart    time.Time
	DtStop     time.Time
}

// Rentable is the basic struct for  entities to rent
type Rentable struct {
	RID            int64             // unique id for this Rentable
	BID            int64             // Business
	Name           string            // name for this rental
	AssignmentTime int64             // can we pre-assign or assign only at commencement
	LastModTime    time.Time         // time of last update to the db record
	LastModBy      int64             // who made the update (Phonebook UID)
	RT             []RentableTypeRef // the list of RTIDs and timestamps for this Rentable
	//-- RentalPeriodDefault int64          // 0 =unset, 1 = short term, 2=longterm
}

// RentableTypeRef is the time-based Rentable type attribute
type RentableTypeRef struct {
	RID            int64     // the Rentable to which this record belongs
	RTID           int64     // the Rentable's type during this time range
	RentCycle      int64     // Override Rent Cycle.  0 =unset,  otherwise same values as RentableType.RentCycle
	ProrationCycle int64     // Override Proration. 0 = unset, otherwise the same values as RentableType.Proration
	DtStart        time.Time // timerange start
	DtStop         time.Time // timerange stop
	LastModTime    time.Time
	LastModBy      int64
}

// RentableSpecialtyRef is the time-based RentableSpecialtyType attribute
type RentableSpecialtyRef struct {
	BID         int64     // associated business
	RID         int64     // the Rentable to which this record belongs
	RSPID       int64     // the rentable specialty type associated with the rentable
	DtStart     time.Time // timerange start
	DtStop      time.Time // timerange stop
	LastModTime time.Time
	LastModBy   int64
}

// RentableStatus archives the state of a Rentable during the specified period of time
type RentableStatus struct {
	RID         int64     // associated Rentable
	DtStart     time.Time // start of period
	DtStop      time.Time // end of period
	Status      int64     // 0 = online, 1 = administrative unit, 2 = owner occupied, 3 = offline
	LastModTime time.Time // time of last update to the db record
	LastModBy   int64     // who made the update (Phonebook UID)
}

// XBusiness combines the Business struct and a map of the Business's Rentable types
type XBusiness struct {
	P  Business
	RT map[int64]RentableType          // what types of things are rented here
	US map[int64]RentableSpecialtyType // index = RSPID, val = RentableSpecialtyType
}

// XRentable is the structure that includes both the Rentable and Unit attributes
type XRentable struct {
	R       Rentable  // the Rentable
	S       []int64   // list of specialties associated with the Rentable
	DtStart time.Time // Start date/time for this Rentable (associated with the Rental Agreement, but may have different dates)
	DtStop  time.Time // Stop time for this Rentable
}

// Journal is the set of attributes describing a Journal entry
type Journal struct {
	JID         int64               // unique id for this Journal entry
	BID         int64               // unique id of Business
	RAID        int64               // unique id of Rental Agreement
	Dt          time.Time           // when this entry was made
	Amount      float64             // the amount
	Type        int64               // 1 means this is an assessment, 2 means it is a payment
	ID          int64               // if Type == 1 then it is the ASMID that caused this entry, of Type ==2 then it is the RCPTID
	Comment     string              // for notes like "prior period adjustment"
	LastModTime time.Time           // auto updated
	LastModBy   int64               // user making the mod
	JA          []JournalAllocation // an array of Journal allocations, breaks the payment or assessment down, total of all the allocations equals the "Amount" above
}

// JournalAllocation describes how the associated Journal amount is allocated
type JournalAllocation struct {
	JAID     int64   // unique id for this allocation
	JID      int64   // associated Journal entry
	RID      int64   // associated Rentable
	Amount   float64 // amount of this allocation
	ASMID    int64   // associated AssessmentID -- source of the charge
	AcctRule string  // describes how this amount distributed across the accounts
}

// JournalMarker describes a period of time where the Journal entries have been locked down
type JournalMarker struct {
	JMID    int64
	BID     int64
	State   int64
	DtStart time.Time
	DtStop  time.Time
}

// LedgerEntry is the structure for LedgerEntry attributes
type LedgerEntry struct {
	LEID        int64
	BID         int64
	JID         int64
	JAID        int64
	LID         int64     // the entry is part of this ledger
	RAID        int64     // RentalAgreement associated with this entry
	Dt          time.Time // date associated with this transaction
	Amount      float64
	Comment     string    // for notes like "prior period adjustment"
	LastModTime time.Time // auto updated
	LastModBy   int64     // user making the mod
	//GLNo        string    // glnumber for the ledger -- DELETE THIS ATTRIBUTE
}

// LedgerMarker describes a period of time period described. The Balance can be
// used going forward from DtStop
type LedgerMarker struct {
	LMID        int64     // unique id for this LM
	LID         int64     // associated GLAccount
	BID         int64     // only valid if Type == 1
	DtStart     time.Time // valid period start
	DtStop      time.Time // valid period end
	Balance     float64   // GLAccount balance at the end of the period
	State       int64     // 0 = unknown, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
	LastModTime time.Time // auto updated
	LastModBy   int64     // user making the mod
}

// GLAccount describes the static (or mostly static) attributes of a Ledger
type GLAccount struct {
	LID          int64     // unique id for this GLAccount
	PLID         int64     // unique id of Parent, 0 if no parent
	BID          int64     // Business unit associated with this GLAccount
	RAID         int64     // associated rental agreement, this field is only used when Type = 1
	GLNumber     string    // acct system name
	Status       int64     // Whether a GL Account is currently unknown=0, inactive=1, active=2
	Type         int64     // flag: 0 = not a default account, 1 = RentalAgreement Account, 10-default cash, 11-GENRCV, 12-GrossSchedRENT, 13-LTL, 14-VAC, ...
	Name         string    // descriptive name for the GLAccount
	AcctType     string    // Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, Cost of Goods Sold, Other Income, Other Expense
	RAAssociated int64     // 1 = Unassociated with RentalAgreement, 2 = Associated with Rental Agreement, 0 = unknown
	AllowPost    int64     // 0 = no posting, 1 = posting is allowed
	LastModTime  time.Time // auto updated
	LastModBy    int64     // user making the mod
}

// RRprepSQL is a collection of prepared sql statements for the RentRoll db
type RRprepSQL struct {
	DeleteAllRentalAgreementPets             *sql.Stmt
	DeleteCustomAttribute                    *sql.Stmt
	DeleteCustomAttributeRef                 *sql.Stmt
	DeleteJournalAllocations                 *sql.Stmt
	DeleteJournalEntry                       *sql.Stmt
	DeleteJournalMarker                      *sql.Stmt
	DeleteLedger                             *sql.Stmt
	DeleteLedgerEntry                        *sql.Stmt
	DeleteLedgerMarker                       *sql.Stmt
	DeleteReceipt                            *sql.Stmt
	DeleteReceiptAllocations                 *sql.Stmt
	DeleteRentableSpecialtyRef               *sql.Stmt
	DeleteRentableStatus                     *sql.Stmt
	DeleteRentableTypeRef                    *sql.Stmt
	DeleteRentalAgreementPet                 *sql.Stmt
	FindAgreementByRentable                  *sql.Stmt
	FindTransactantByPhoneOrEmail            *sql.Stmt
	GetAgreementsForRentable                 *sql.Stmt
	GetAllAssessmentsByBusiness              *sql.Stmt
	GetAllAssessmentsByRAID                  *sql.Stmt
	GetAllBusinesses                         *sql.Stmt
	GetAllBusinessRentableTypes              *sql.Stmt
	GetAllBusinessSpecialtyTypes             *sql.Stmt
	GetAllJournalsInRange                    *sql.Stmt
	GetAllLedgerEntriesInRange               *sql.Stmt
	GetAllLedgerMarkersInRange               *sql.Stmt
	GetAllRentableAssessments                *sql.Stmt
	GetAllRentablesByBusiness                *sql.Stmt
	GetAllRentalAgreementPets                *sql.Stmt
	GetAllRentalAgreements                   *sql.Stmt
	GetAllRentalAgreementsByRange            *sql.Stmt
	GetAllRentalAgreementTemplates           *sql.Stmt
	GetAllTransactants                       *sql.Stmt
	GetAssessment                            *sql.Stmt
	GetAssessmentType                        *sql.Stmt
	GetAssessmentTypeByName                  *sql.Stmt
	GetBuilding                              *sql.Stmt
	GetBusiness                              *sql.Stmt
	GetBusinessByDesignation                 *sql.Stmt
	GetCustomAttribute                       *sql.Stmt
	GetCustomAttributeRefs                   *sql.Stmt
	GetDefaultLedgers                        *sql.Stmt
	GetJournal                               *sql.Stmt
	GetJournalAllocation                     *sql.Stmt
	GetJournalAllocations                    *sql.Stmt
	GetJournalByRange                        *sql.Stmt
	GetJournalMarker                         *sql.Stmt
	GetJournalMarkers                        *sql.Stmt
	GetLatestLedgerMarkerByLID               *sql.Stmt
	GetLedger                                *sql.Stmt
	GetLedgerByGLNo                          *sql.Stmt
	GetLedgerByType                          *sql.Stmt
	GetLedgerEntriesForRAID                  *sql.Stmt
	GetLedgerEntriesInRangeByGLNo            *sql.Stmt
	GetLedgerEntriesInRangeByLID             *sql.Stmt
	GetLedgerEntry                           *sql.Stmt
	GetLedgerList                            *sql.Stmt
	GetLedgerMarkerByDateRange               *sql.Stmt
	GetLedgerMarkerByLIDDateRange            *sql.Stmt
	GetLedgerMarkerByRAID                    *sql.Stmt
	GetLedgerMarkers                         *sql.Stmt
	GetPaymentTypesByBusiness                *sql.Stmt
	GetPayor                                 *sql.Stmt
	GetProspect                              *sql.Stmt
	GetRABalanceLedger                       *sql.Stmt
	GetReceipt                               *sql.Stmt
	GetReceiptAllocations                    *sql.Stmt
	GetReceiptsInDateRange                   *sql.Stmt
	GetReceiptsInRAIDDateRange               *sql.Stmt
	GetRentable                              *sql.Stmt
	GetRentableByName                        *sql.Stmt
	GetRentableMarketRates                   *sql.Stmt
	GetRentableSpecialtyRefs                 *sql.Stmt
	GetRentableSpecialtyRefsByRange          *sql.Stmt
	GetRentableSpecialtyType                 *sql.Stmt
	GetRentableSpecialtyTypeByName           *sql.Stmt
	GetRentableStatusByRange                 *sql.Stmt
	GetRentableType                          *sql.Stmt
	GetRentableTypeByStyle                   *sql.Stmt
	GetRentableTypeRefsByRange               *sql.Stmt
	GetRentableUsers                         *sql.Stmt
	GetRentalAgreement                       *sql.Stmt
	GetRentalAgreementByBusiness             *sql.Stmt
	GetRentalAgreementByRentalTemplateNumber *sql.Stmt
	GetRentalAgreementPayors                 *sql.Stmt
	GetRentalAgreementPet                    *sql.Stmt
	GetRentalAgreementRentables              *sql.Stmt
	GetRentalAgreementTemplate               *sql.Stmt
	GetSecurityDepositAssessment             *sql.Stmt
	GetTransactant                           *sql.Stmt
	GetUnitAssessments                       *sql.Stmt
	GetUser                                  *sql.Stmt
	InsertAssessment                         *sql.Stmt
	InsertAssessmentType                     *sql.Stmt
	InsertBuilding                           *sql.Stmt
	InsertBuildingWithID                     *sql.Stmt
	InsertBusiness                           *sql.Stmt
	InsertCustomAttribute                    *sql.Stmt
	InsertCustomAttributeRef                 *sql.Stmt
	InsertJournal                            *sql.Stmt
	InsertJournalAllocation                  *sql.Stmt
	InsertJournalMarker                      *sql.Stmt
	InsertLedger                             *sql.Stmt
	InsertLedgerAllocation                   *sql.Stmt
	InsertLedgerEntry                        *sql.Stmt
	InsertLedgerMarker                       *sql.Stmt
	InsertPaymentType                        *sql.Stmt
	InsertPayor                              *sql.Stmt
	InsertProspect                           *sql.Stmt
	InsertReceipt                            *sql.Stmt
	InsertReceiptAllocation                  *sql.Stmt
	InsertRentable                           *sql.Stmt
	InsertRentableMarketRates                *sql.Stmt
	InsertRentableSpecialtyRef               *sql.Stmt
	InsertRentableSpecialtyType              *sql.Stmt
	InsertRentableStatus                     *sql.Stmt
	InsertRentableType                       *sql.Stmt
	InsertRentableTypeRef                    *sql.Stmt
	InsertRentableUser                       *sql.Stmt
	InsertRentalAgreement                    *sql.Stmt
	InsertRentalAgreementPayor               *sql.Stmt
	InsertRentalAgreementPet                 *sql.Stmt
	InsertRentalAgreementRentable            *sql.Stmt
	InsertRentalAgreementTemplate            *sql.Stmt
	InsertTransactant                        *sql.Stmt
	InsertUser                               *sql.Stmt
	UpdateLedger                             *sql.Stmt
	UpdateLedgerMarker                       *sql.Stmt
	UpdateRentableSpecialtyRef               *sql.Stmt
	UpdateRentableStatus                     *sql.Stmt
	UpdateRentableTypeRef                    *sql.Stmt
	UpdateRentalAgreementPet                 *sql.Stmt
	UpdateTransactant                        *sql.Stmt
}

// PBprepSQL is the structure of prepared sql statements for the Phonebook db
type PBprepSQL struct {
	GetCompanyByDesignation      *sql.Stmt
	GetCompany                   *sql.Stmt
	GetBusinessUnitByDesignation *sql.Stmt
}

// BusinessTypeLists is a struct holding a collection of Types associated with a business
type BusinessTypeLists struct {
	BID int64
	//AsmtTypes    map[int64]*AssessmentType // moved to parent as it is biz-independent
	PmtTypes     map[int64]*PaymentType
	DefaultAccts map[int64]*GLAccount // index by the predifined contants DFAC*, value = GL No of that account
}

// RRdb is a struct with all variables needed by the db infrastructure
var RRdb struct {
	Prepstmt  RRprepSQL
	PBsql     PBprepSQL
	Dbdir     *sql.DB                  // phonebook db
	Dbrr      *sql.DB                  //rentroll db
	AsmtTypes map[int64]AssessmentType // every one in the db
	BizTypes  map[int64]*BusinessTypeLists
}

// InitDBHelpers initializes the db infrastructure
func InitDBHelpers(dbrr, dbdir *sql.DB) {
	RRdb.Dbdir = dbdir
	RRdb.Dbrr = dbrr
	RRdb.BizTypes = make(map[int64]*BusinessTypeLists, 0)
	buildPreparedStatements()
	buildPBPreparedStatements()
	RRdb.AsmtTypes = GetAssessmentTypes()
}

// InitBusinessFields initialize the lists in rlib's internal data structures
func InitBusinessFields(bid int64) {
	if nil == RRdb.BizTypes[bid] {
		bt := BusinessTypeLists{
			BID: bid,
			// AsmtTypes:    make(map[int64]*AssessmentType),
			PmtTypes:     make(map[int64]*PaymentType),
			DefaultAccts: make(map[int64]*GLAccount),
		}
		RRdb.BizTypes[bid] = &bt
	}
}
