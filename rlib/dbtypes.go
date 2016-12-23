package rlib

import (
	"database/sql"
	"time"
)

// NO and all the rest are constants that are used with the RentRoll database
const (
	NO  = int64(0) // std negative value
	YES = int64(1)

	RPTTEXT = 1
	RPTHTML = 2

	ELEMPERSON          = 1 // people
	ELEMCOMPANY         = 2 // companies
	ELEMCLASS           = 3 // classes
	ELEMSVC             = 4 // the executable service
	ELEMRENTABLETYPE    = 5 // RentableType element
	ELEMRATEPLAN        = 6 // Rate Plan
	ELEMTRANSACTANT     = 7
	ELEMUSER            = 8
	ELEMPROSPECT        = 9
	ELEMAPPLICANT       = 10
	ELEMPAYOR           = 11
	ELEMRENTABLE        = 12
	ELEMRENTALAGREEMENT = 13
	ELEMLAST            = 13 // keep in sync with last one added

	// CUSTSTRING et al are Custom Attribute types
	CUSTSTRING = 0
	CUSTINT    = 1
	CUSTUINT   = 2
	CUSTFLOAT  = 3
	CUSTDATE   = 4
	CUSTLAST   = 4 // this should be maintained as matching the highest index value in the group

	RENT                      = 1
	SECURITYDEPOSIT           = 2
	SECURITYDEPOSITASSESSMENT = 58

	ACCTSTATUSINACTIVE = 1
	ACCTSTATUSACTIVE   = 2
	RAASSOCIATED       = 1
	RAUNASSOCIATED     = 2

	GLCASH       = 10
	GLGENRCV     = 11
	GLGSRENT     = 12
	GLLTL        = 13
	GLVAC        = 14
	GLSECDEP     = 16
	GLOWNREQUITY = 17
	GLLAST       = 17 // set this to the last default account index
	// GLSECDEPRCV  = 15

	CYCLENORECUR   = 0
	CYCLESECONDLY  = 1
	CYCLEMINUTELY  = 2
	CYCLEHOURLY    = 3
	CYCLEDAILY     = 4
	CYCLEWEEKLY    = 5
	CYCLEMONTHLY   = 6
	CYCLEQUARTERLY = 7
	CYCLEYEARLY    = 8

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

	// RRDATEFMT is a shorthand date format used for text output
	// Use these values:	Mon Jan 2 15:04:05 MST 2006
	// const RRDATEFMT = "02-Jan-2006 3:04PM MST"
	// const RRDATEFMT = "01/02/06 3:04PM MST"
	RRDATEFMT        = "01/02/06"
	RRDATEFMT2       = "1/2/06"
	RRDATEFMT3       = "1/2/2006"
	RRDATEFMT4       = "01/02/2006"
	RRDATEINPFMT     = "2006-01-02"
	RRDATETIMEINPFMT = "2006-01-02 15:04:00 MST"
)

//==========================================
// ASMID = Assessment id
// ATypeLID = assessment type id
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
// TCID = Transactant id == PayorID == UserID == ProspectID
//==========================================

// StringList is a generic list structure for lists of strings. These could be used to implement things like
// the list of reasons why an applicant's application was turned down, the list of reasons why a tenant is
// moving out, etc.
type StringList struct {
	SLID        int64      // unique id for this stringlist
	BID         int64      // the business to which this stringlist belongs
	Name        string     // stringlist name
	LastModTime time.Time  // when was this record last written
	LastModBy   int64      // employee UID (from phonebook) that modified it
	S           []SLString // array of SLStrings associated with this SLID
}

// SLString defines an individual string member of a StringList
type SLString struct {
	SLSID       int64     // unique id of this string
	SLID        int64     // to which stringlist does this string belong?
	Value       string    // value of this string
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
}

// NoteType describes the type of note this is
type NoteType struct {
	NTID        int64     // note type id
	BID         int64     // business associated with this note type
	Name        string    // the actual note
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
}

// Note is dated comment from a user
type Note struct {
	NID         int64     // unique ID for this note
	NLID        int64     // notelist to which this note belongs
	PNID        int64     // NID of parent note
	NTID        int64     // note type id
	RID         int64     // Meta Tag - this note is related to Rentable RID
	RAID        int64     // Meta Tag - this note is related to Rental Agreement RAID
	TCID        int64     // Meta Tag - this note is related to Transactant TCID
	Comment     string    // the actual note
	CN          []Note    // array of child notes
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
}

// NoteList is a collection of Notes (NIDs)
type NoteList struct {
	NLID        int64     // unique id for the notelist
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	N           []Note    // the list of notes
}

// CustomAttribute is a struct containing user-defined custom attributes for objects
type CustomAttribute struct {
	CID         int64     // unique id
	Type        int64     // what type of value: 0 = string, 1 = int64, 2 = float64
	Name        string    // what its called
	Value       string    // string value -- will be xlated on load / store
	Units       string    // optional units value.  Ex:  "feet", "gallons", "cubic feet", ...
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

// AssessmentType is a list of chargest that the company can make
type AssessmentType struct {
	ASMTID      int64     // unique id for this rate plan
	BID         int64     // which business
	Name        string    // RATemplateName a string associated with each rental type agreement (essentially, the doc name)
	DtStart     time.Time // when does this charge go into effect
	DtStop      time.Time // when does this charge end
	Cycle       int64     // recurrence
	Prorate     int64     //
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
}

// RentalAgreementTemplate is a template used to set up new rental agreements
type RentalAgreementTemplate struct {
	RATID          int64     // unique id for this rate plan
	BID            int64     // which business
	RATemplateName string    // RATemplateName a string associated with each rental type agreement (essentially, the doc name)
	LastModTime    time.Time // when was this record last written
	LastModBy      int64     // employee UID (from phonebook) that modified it
}

// RatePlan is a structure of static attributes of a rate plan, which describes charges for rentable types, varying by customer
type RatePlan struct {
	RPID        int64               // unique id for this rate plan
	BID         int64               // Business
	Name        string              // The name of this RatePlan
	LastModTime time.Time           // when was this record last written
	LastModBy   int64               // employee UID (from phonebook) that modified it
	OD          []OtherDeliverables // other Deliverables associated with the rate plan
}

// FlRatePlanGDS and the others are bit flags for the RatePlan custom attribute: FLAGS
const (
	FlRatePlanGDS   = 1 << 0 // RatePlan - bit 00 export to GDS
	FlRatePlanSabre = 1 << 1 // RatePlan - bit 01 export to Sabre
)

// RatePlanRef contains the time sensitive attributes of a RatePlan
type RatePlanRef struct {
	RPRID             int64               // unique id for this ref
	RPID              int64               // which rate plan
	DtStart           time.Time           // when does it go into effect
	DtStop            time.Time           // when does it stop
	FeeAppliesAge     int64               // the age at which a user is counted when determining extra user fees or eligibility for rental
	MaxNoFeeUsers     int64               // maximum number of users for no fees. Greater than this number means fee applies
	AdditionalUserFee float64             // extra fee per user when exceeding MaxNoFeeUsers
	PromoCode         string              // just a string
	CancellationFee   float64             // charge for cancellation
	FLAGS             uint64              // 1<<0 -- HideRate
	LastModTime       time.Time           // when was this record last written
	LastModBy         int64               // employee UID (from phonebook) that modified it
	RT                []RatePlanRefRTRate // all associated RentableType Rates
	SP                []RatePlanRefSPRate // all associated RentableSpecialtyType Rates
}

// FlRTRRefHide and the others are bit flags for the RatePlanRef
const (
	FlRTRRefHide = 1 << 0 // do not show this rate plan to users
)

// RatePlanRefRTRate is RatePlan RPRID's rate information for the RentableType (RTID)
type RatePlanRefRTRate struct {
	RPRID int64   // which RatePlanRef is this
	RTID  int64   // which RentableType
	FLAGS uint64  // 1<<0 = percent flag 0 = Val is an absolute price, 1 = percent of MarketRate,
	Val   float64 // Val
}

// FlRTRpct and the others are bit flags for the RatePlanRefRTRate
const (
	FlRTRpct = 1 << 0 // bit 0 = percent flag. 0 means it's an absolute amount, 1 means it's a % of Market Rate
	FlRTRna  = 1 << 1 // bit 1 = n/a flag, 0 means that this RTID is affected, 1 means it is not affected
)

// RatePlanRefSPRate is RatePlan RPRID's rate information for the Specialties
type RatePlanRefSPRate struct {
	RPRID int64   // which RatePlanRef is this
	RTID  int64   // which RentableType
	RSPID int64   // which Specialty
	FLAGS uint64  // 1<<0 = percent flag 0 = Val is an absolute price, 1 = percent of MarketRate,
	Val   float64 // Val
}

// FlSPRpct and the others are bit flags for the RatePlanRefSPRate
const (
	FlSPRpct = 1 << 0 // bit 0 = percent flag. 0 means it's an absolute amount, 1 means it's a % of Market Rate
	FlSPRna  = 1 << 1 // bit 1 = n/a flag, 0 means that this RTID is affected, 1 means it is not affected
)

// RatePlanOD defines which other deliverables are associated with a RatePlan.
// A RatePlan can refer to multiple OtherDeliverables.
type RatePlanOD struct {
	RPRID int64 // with which RatePlan is this OtherDeliverable associated?
	ODID  int64 // points to an OtherDeliverables
}

// OtherDeliverables defines special offers associated with RatePlanRefs. These are for promotions. Examples of OtherDeliverables
// would include things like 2 Seaworld tickets, etc.  Referenced by RatePlanRef
// Multiple RatePlanRefs can refer to the same OtherDeliverables.
type OtherDeliverables struct {
	ODID   int64  // Unique ID for this OtherDeliverables
	Name   string // Description of the other deliverables. Ex: 2 Seaworld tickets
	Active int64  // Flag: Is this list still active?  dropdown interface lists only the active ones
}

// RentalAgreement binds one or more payors to one or more rentables
type RentalAgreement struct {
	RAID                   int64       // internal unique id
	RATID                  int64       // reference to Occupancy Master Agreement
	BID                    int64       // Business (so that we can process by Business)
	NLID                   int64       // Note ID
	AgreementStart         time.Time   // start date for rental agreement contract
	AgreementStop          time.Time   // stop date for rental agreement contract
	PossessionStart        time.Time   // start date for Occupancy
	PossessionStop         time.Time   // stop date for Occupancy
	RentStart              time.Time   // start date for Rent
	RentStop               time.Time   // stop date for Rent
	RentCycleEpoch         time.Time   // Date on which rent cycle recurs. Start date for the recurring rent assessment
	UnspecifiedAdults      int64       // adults who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels
	UnspecifiedChildren    int64       // children who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels.
	Renewal                int64       // 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
	SpecialProvisions      string      // free-form text
	LeaseType              int64       // Full Service Gross, Gross, ModifiedGross, Tripple Net
	ExpenseAdjustmentType  int64       // Base Year, No Base Year, Pass Through
	ExpensesStop           float64     // cap on the amount of oexpenses that can be passed through to the tenant
	ExpenseStopCalculation string      // note on how to determine the expense stop
	BaseYearEnd            time.Time   // last day of the base year
	ExpenseAdjustment      time.Time   // the next date on which an expense adjustment is due
	EstimatedCharges       float64     // a periodic fee charged to the tenant to reimburse LL for anticipated expenses
	RateChange             float64     // predetermined amount of rent increase, expressed as a percentage
	NextRateChange         time.Time   // he next date on which a RateChange will occur
	PermittedUses          string      // indicates primary use of the space, ex: doctor's office, or warehouse/distribution, etc.
	ExclusiveUses          string      // those uses to which the tenant has the exclusive rights within a complex, ex: Trader Joe's may have the exclusive right to sell groceries
	ExtensionOption        string      // the right to extend the term of lease by giving notice to LL, ex: 2 options to extend for 5 years each
	ExtensionOptionNotice  time.Time   // the last dade by wich a Tenant can give notice of their intention to exercise the right to an extension option period
	ExpansionOption        string      // the right to expand to certanin spaces that are typically contiguous to their primary space
	ExpansionOptionNotice  time.Time   // the last dade by wich a Tenant can give notice of their intention to exercise the right to an Expansion Option
	RightOfFirstRefusal    string      // Tenant may have the right to purchase their premises if LL chooses to sell
	LastModTime            time.Time   // when was this record last written
	LastModBy              int64       // employee UID (from phonebook) that modified it
	R                      []XRentable // all the rentables
	P                      []XPerson   // all the payors
	T                      []XPerson   // all the users
}

// RentalAgreementRentable describes a Rentable associated with a rental agreement
type RentalAgreementRentable struct {
	RAID         int64     // associated rental agreement
	RID          int64     // the Rentable
	CLID         int64     // commission ledger -- applies if outside sales rented this rentable
	ContractRent float64   // the rent
	DtStart      time.Time // start date/time for this Rentable
	DtStop       time.Time // stop date/time
}

// RentalAgreementPayor describes a Payor associated with a rental agreement
type RentalAgreementPayor struct {
	RAID    int64
	TCID    int64     // the payor's transactant id
	DtStart time.Time // start date/time for this Payor
	DtStop  time.Time // stop date/time
	FLAGS   uint64    // 1<<0 is the bit that indicates this payor is a 'guarantor'
}

// RentalAgreementTax - the time based attribute for whether the rental agreement is taxable
type RentalAgreementTax struct {
	RAID    int64     //associated rental agreement
	DtStart time.Time // start date/time for this Payor
	DtStop  time.Time // stop date/time
	FLAGS   uint64    // 1<<0 is whether the agreement is taxable
}

// RentableUser describes a User associated with a rental agreement
type RentableUser struct {
	RID     int64     // associated Rentable
	TCID    int64     // pointer to Transactant
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

// DemandSource is a structure
type DemandSource struct {
	SourceSLSID int64     // DemandSource ID
	BID         int64     // Business unit
	Name        string    // name of source
	Industry    string    // what industry is this source in
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
}

// Transactant is the basic structure of information
// about a person who is a Prospect, applicant, User, or Payor
type Transactant struct {
	Recid          int64 `json:"recid"` // this is to support the grid widget
	TCID           int64
	BID            int64
	NLID           int64
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
	LastModTime    time.Time
	LastModBy      int64
}

// Prospect contains info over and above
type Prospect struct {
	// PRSPID                 int64
	TCID                   int64
	EmployerName           string
	EmployerStreetAddress  string
	EmployerCity           string
	EmployerState          string
	EmployerPostalCode     string
	EmployerEmail          string
	EmployerPhone          string
	Occupation             string
	ApplicationFee         float64   // if non-zero this Prospect is an applicant
	DesiredUsageStartDate  time.Time // predicted rent start date
	RentableTypePreference int64     // RentableType
	FLAGS                  uint64    // 0 = Approved/NotApproved,
	Approver               int64     // UID from Directory
	DeclineReasonSLSID     int64     // SLSid of reason
	OtherPreferences       string    // arbitrary text
	FollowUpDate           time.Time // automatically fill out this date to sysdate + 24hrs
	CSAgent                int64     // Accord Directory UserID - for the CSAgent
	OutcomeSLSID           int64     // id of string from a list of outcomes. Melissa to provide reasons
	FloatingDeposit        float64   // d $(GLCASH) _, c $(GLGENRCV) _; assign to a shell of a Rental Agreement
	RAID                   int64     // created to hold On Account amount of Floating Deposit
	LastModTime            time.Time
	LastModBy              int64
}

// User contains all info common to a person
type User struct {
	// USERID                    int64
	TCID                      int64
	Points                    int64
	DateofBirth               time.Time
	EmergencyContactName      string
	EmergencyContactAddress   string
	EmergencyContactTelephone string
	EmergencyEmail            string
	AlternateAddress          string
	EligibleFutureUser        int64
	Industry                  string
	SourceSLSID               int64
	LastModTime               time.Time
	LastModBy                 int64
	Vehicles                  []Vehicle
}

// Vehicle contains all the vehicle information for a User's vehicld
type Vehicle struct {
	VID                 int64
	TCID                int64
	BID                 int64
	VehicleType         string
	VehicleMake         string
	VehicleModel        string
	VehicleColor        string
	VehicleYear         int64
	LicensePlateState   string
	LicensePlateNumber  string
	ParkingPermitNumber string
	DtStart             time.Time
	DtStop              time.Time
	LastModTime         time.Time
	LastModBy           int64
}

// Payor is attributes of the person financially responsible
// for the rent.
type Payor struct {
	// PID                 int64
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
	Usr User
	Psp Prospect
	Pay Payor
}

// Assessment is a charge associated with a Rentable
type Assessment struct {
	ASMID          int64     // unique id for this assessment
	PASMID         int64     // parent Assessment, if this is non-zero it means this assessment is an instance of the recurring assessment with id PASMID. When non-zero DO NOT process as a recurring assessment, it is an instance
	BID            int64     // what Business
	RID            int64     // the Rentable
	ATypeLID       int64     // what type of assessment
	RAID           int64     // associated Rental Agreement
	Amount         float64   // how much
	Start          time.Time // start time
	Stop           time.Time // stop time, may be the same as start time or later
	RentCycle      int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, G = quarterly, 8 = yearly
	ProrationCycle int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
	InvoiceNo      int64     // A uniqueID for the invoice number
	AcctRule       string    // expression showing how to account for the amount
	Comment        string
	LastModTime    time.Time
	LastModBy      int64
}

// Business is the set of attributes describing a rental or hotel Business
type Business struct {
	BID                   int64
	Designation           string // reference to designation in Phonebook db
	Name                  string
	DefaultRentCycle      int64     // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	DefaultProrationCycle int64     // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	DefaultGSRPC          int64     // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	LastModTime           time.Time // when was this record last written
	LastModBy             int64     // employee UID (from phonebook) that modified it
	// ParkingPermitInUse    int64     // yes/no  0 = no, 1 = yes
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
	RCPTID         int64
	PRCPTID        int64 // Parent RCPTID, points to RCPT being amended/corrected by this receipt
	BID            int64
	RAID           int64
	PMTID          int64
	Dt             time.Time
	DocNo          string // check number, money order number, etc.; documents the payment
	Amount         float64
	AcctRule       string
	Comment        string
	OtherPayorName string // if not '', the name of a payor who paid this receipt and who may not be in our system
	LastModTime    time.Time
	LastModBy      int64
	RA             []ReceiptAllocation
}

// ReceiptAllocation defines an allocation of a Receipt amount.
type ReceiptAllocation struct {
	RCPTID   int64
	Amount   float64
	ASMID    int64
	AcctRule string
}

// Depository is a bank account or other account where deposits are made
type Depository struct {
	DEPID       int64     // unique id for a depository
	BID         int64     // which business
	Name        string    // Name of Depository: First Data, Nyax, CCI, Oklahoma Fidelity
	AccountNo   string    // account number at this Depository
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
}

// Deposit is simply a list of receipts that form a deposit to a Depository. This struct contains
// the static attributes of the list
type Deposit struct {
	DID         int64         // Unique id of this deposit
	BID         int64         // business id
	DEPID       int64         // Depository id where the deposit was made
	DPMID       int64         // Deposit method
	Dt          time.Time     // Date of deposit
	Amount      float64       // the total amount of the deposit
	LastModTime time.Time     // when was this record last written
	LastModBy   int64         // employee UID (from phonebook) that modified it
	DP          []DepositPart // array of DepositParts for this deposit
}

// DepositPart is a reference to a Receipt that is part of this deposit.  Another way of
// thinking about it is that this query produces the list of all receipts in a Deposit:
//		SELECT RCPTID WHERE DIP=someDID
type DepositPart struct {
	DID    int64 // deposit id
	RCPTID int64 // receipt id
}

// DepositMethod is a list of methods used to make deposits to a depository
type DepositMethod struct {
	DPMID int64  //the method id
	BID   int64  // business id
	Name  string // descriptive name
}

// Invoice is a structure that defines an invoice - a collection of assessments
type Invoice struct {
	InvoiceNo   int64               // Unique id for this invoice
	BID         int64               // bid (remit to)
	Dt          time.Time           // Date of invoice
	DtDue       time.Time           // Date when the invoice is due
	Amount      float64             // total amount of all assessments in this invoice
	DeliveredBy string              // mail, FedEx, UPS, email, fax, hand delivered, carrier pigeon :-) ...
	LastModTime time.Time           // when was this record last written
	LastModBy   int64               // employee UID (from phonebook) that modified it
	A           []InvoiceAssessment // list of assessments in this invoice
	P           []InvoicePayor      // list of payors
}

// InvoiceAssessment is a reference to an Assessment that is part of this invoice.  Another way of
// thinking about it is that this query produces the list of all assessments in an invoice:
//		SELECT ASMID WHERE InvoiceNo=somenumber
type InvoiceAssessment struct {
	InvoiceNo int64 // the invoice number
	ASMID     int64 // assessment
}

// InvoicePayor is a reference to a Payor for this invoice.  Another way of
// thinking about it is that this query produces the list of all payors for an invoice:
//		SELECT PID WHERE InvoiceNo=somenumber
type InvoicePayor struct {
	InvoiceNo int64 // the invoice number
	PID       int64 // Payor ID
}

// RentableSpecialty is the structure for attributes of a Rentable specialty
type RentableSpecialty struct {
	RSPID       int64
	BID         int64
	Name        string
	Fee         float64 // proration inherited from the rentable / rentable type.
	Description string
}

// RentableType is the set of attributes describing the different types of Rentable items
type RentableType struct {
	RTID           int64                      // unique identifier for this RentableType
	BID            int64                      // the business unit to which this RentableType belongs
	Style          string                     // a short name
	Name           string                     // longer name
	RentCycle      int64                      // frequency at which rent accrues, 0 = not set or n/a, 1 = secondly, 2=minutely, 3=hourly, 4=daily, 5=weekly, 6=monthly...
	Proration      int64                      // frequency for prorating rent if the full rentcycle is not used
	GSRPC          int64                      // Time increments in which GSR is calculated to account for rate changes
	ManageToBudget int64                      // 0=no, 1 = yes
	MR             []RentableMarketRate       // array of time sensitive market rates
	CA             map[string]CustomAttribute // index by Name of attribute, associated custom attributes
	MRCurrent      float64                    // the current market rate (historical values are in MR)
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
	Recid          int64             `json:"recid"` // this is to support the grid widget
	RID            int64             // unique id for this Rentable
	BID            int64             // Business
	Name           string            // name for this rental
	AssignmentTime int64             // can we pre-assign or assign only at commencement
	LastModTime    time.Time         // time of last update to the db record
	LastModBy      int64             // who made the update (Phonebook UID)
	RT             []RentableTypeRef // the list of RTIDs and timestamps for this Rentable
	RTCurrent      int64             // RentableType ID its current type (current as defined by system datetime), NOT A DB FIELD
	//-- RentalPeriodDefault int64          // 0 =unset, 1 = short term, 2=longterm
}

// RentableTypeRef is the time-based Rentable type attribute
type RentableTypeRef struct {
	RID                    int64     // the Rentable to which this record belongs
	RTID                   int64     // the Rentable's type during this time range
	OverrideRentCycle      int64     // Override Rent Cycle.  0 =unset,  otherwise same values as RentableType.RentCycle
	OverrideProrationCycle int64     // Override Proration. 0 = unset, otherwise the same values as RentableType.Proration
	DtStart                time.Time // timerange start
	DtStop                 time.Time // timerange stop
	LastModTime            time.Time
	LastModBy              int64
}

// RentCycleRef is a simplified struct containing a rent cycle and the
// time duration for which it is valid. This structure of data is not
// in the database. It is used for calculations where we don't want to worry about
// whether the default rent cycle is being overridden, etc. All of that info will be
// reflected in the values in this struct.
type RentCycleRef struct {
	DtStart        time.Time // timerange start
	DtStop         time.Time // timerange stop
	RentCycle      int64     // Rent Cycle during DtStart-DtStop
	ProrationCycle int64     // Proration during DtStart-DtStop
}

// RentableSpecialtyRef is the time-based RentableSpecialty attribute
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
	RID              int64     // associated Rentable
	DtStart          time.Time // start of period
	DtStop           time.Time // end of period
	DtNoticeToVacate time.Time // user has indicated they will vacate on this date
	Status           int64     // 0 = online, 1 = administrative unit, 2 = owner occupied, 3 = offline
	LastModTime      time.Time // time of last update to the db record
	LastModBy        int64     // who made the update (Phonebook UID)
}

// XBusiness combines the Business struct and a map of the Business's Rentable types
type XBusiness struct {
	P  Business
	RT map[int64]RentableType      // what types of things are rented here
	US map[int64]RentableSpecialty // index = RSPID, val = RentableSpecialty
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
	Type        int64               // 0 = unassociated with RA, 1 means this is an assessment, 2 means it is a payment
	ID          int64               // if Type == 0 then it is the RentableID, if Type == 1 then it is the ASMID that caused this entry, if Type ==2 then it is the RCPTID
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
	RID         int64     // Rentable associated with this entry
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
	RAID        int64     // if 0 then it's the LM for the whole account, if > 0 it's the amount for the rental agreement RAID
	RID         int64     // if 0 then it's the LM for the whole account, if > 0 it's the amount for the Rentable RID
	Dt          time.Time // Balance is valid as of this time
	Balance     float64   // GLAccount balance at the end of the period
	State       int64     // 0 = unknown, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
	LastModTime time.Time // auto updated
	LastModBy   int64     // user making the mod
}

// GLAccount describes the static (or mostly static) attributes of a Ledger
type GLAccount struct {
	Recid          int       `json:"recid"` // this is for the grid widget
	LID            int64     // unique id for this GLAccount
	PLID           int64     // unique id of Parent, 0 if no parent
	BID            int64     // Business unit associated with this GLAccount
	RAID           int64     // associated rental agreement, this field is only used when Type = 1
	GLNumber       string    // acct system name
	Status         int64     // Whether a GL Account is currently unknown=0, inactive=1, active=2
	Type           int64     // flag: 0 = not a default account, 1-9 reserved, 10-default cash, 11-GENRCV, 12-GrossSchedRENT, 13-LTL, 14-VAC, ...
	Name           string    // descriptive name for the GLAccount
	AcctType       string    // QB Acct Type: Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, Cost of Goods Sold, Other Income, Other Expense
	RAAssociated   int64     // 1 = Unassociated with RentalAgreement, 2 = Associated with Rental Agreement, 0 = unknown
	AllowPost      int64     // 0 = no posting, 1 = posting is allowed
	RARequired     int64     // 0 = during rental period, 1 = valid prior or during, 2 = valid during or after, 3 = valid before, during, and after
	ManageToBudget int64     // 0 = do not manage to budget; no ContractRent amount required. 1 = Manage to budget, ContractRent required.
	Description    string    // description for this account
	LastModTime    time.Time // auto updated
	LastModBy      int64     // user making the mod
}

// RRprepSQL is a collection of prepared sql statements for the RentRoll db
type RRprepSQL struct {
	// GetRABalanceLedger                       *sql.Stmt
	DeleteAllRentalAgreementPets       *sql.Stmt
	DeleteCustomAttribute              *sql.Stmt
	DeleteCustomAttributeRef           *sql.Stmt
	DeleteDemandSource                 *sql.Stmt
	DeleteDeposit                      *sql.Stmt
	DeleteDepositMethod                *sql.Stmt
	DeleteDepository                   *sql.Stmt
	DeleteDepositParts                 *sql.Stmt
	DeleteInvoice                      *sql.Stmt
	DeleteInvoiceAssessments           *sql.Stmt
	DeleteInvoicePayors                *sql.Stmt
	DeleteJournalAllocations           *sql.Stmt
	DeleteJournalEntry                 *sql.Stmt
	DeleteJournalMarker                *sql.Stmt
	DeleteLedger                       *sql.Stmt
	DeleteLedgerEntry                  *sql.Stmt
	DeleteLedgerMarker                 *sql.Stmt
	DeleteNote                         *sql.Stmt
	DeleteNoteList                     *sql.Stmt
	DeleteNoteType                     *sql.Stmt
	DeleteRatePlan                     *sql.Stmt
	DeleteRatePlanRef                  *sql.Stmt
	DeleteRatePlanRefRTRate            *sql.Stmt
	DeleteRatePlanRefSPRate            *sql.Stmt
	DeleteReceipt                      *sql.Stmt
	DeleteReceiptAllocations           *sql.Stmt
	DeleteRentableSpecialtyRef         *sql.Stmt
	DeleteRentableStatus               *sql.Stmt
	DeleteRentableTypeRef              *sql.Stmt
	DeleteRentalAgreementPet           *sql.Stmt
	DeleteRentalAgreementTax           *sql.Stmt
	DeleteSLString                     *sql.Stmt
	DeleteSLStrings                    *sql.Stmt
	DeleteStringList                   *sql.Stmt
	DeleteVehicle                      *sql.Stmt
	FindAgreementByRentable            *sql.Stmt
	FindTransactantByPhoneOrEmail      *sql.Stmt
	GetAgreementsForRentable           *sql.Stmt
	GetAllAssessmentsByBusiness        *sql.Stmt
	GetAllAssessmentsByRAID            *sql.Stmt
	GetAllBusinesses                   *sql.Stmt
	GetAllBusinessRentableTypes        *sql.Stmt
	GetAllBusinessSpecialtyTypes       *sql.Stmt
	GetAllCustomAttributeRefs          *sql.Stmt
	GetAllCustomAttributes             *sql.Stmt
	GetAllDemandSources                *sql.Stmt
	GetAllDepositMethods               *sql.Stmt
	GetAllDepositories                 *sql.Stmt
	GetAllDepositsInRange              *sql.Stmt
	GetAllInvoicesInRange              *sql.Stmt
	GetAllJournalsInRange              *sql.Stmt
	GetAllLedgerEntriesForRAID         *sql.Stmt
	GetAllLedgerEntriesForRID          *sql.Stmt
	GetAllLedgerEntriesInRange         *sql.Stmt
	GetAllLedgerMarkersOnOrBefore      *sql.Stmt
	GetAllNotes                        *sql.Stmt
	GetAllNoteTypes                    *sql.Stmt
	GetAllRatePlanRefRTRates           *sql.Stmt
	GetAllRatePlanRefsInRange          *sql.Stmt
	GetAllRatePlanRefSPRates           *sql.Stmt
	GetAllRatePlans                    *sql.Stmt
	GetAllRentableAssessments          *sql.Stmt
	GetAllRentablesByBusiness          *sql.Stmt
	GetAllRentableSpecialtyRefs        *sql.Stmt
	GetAllRentalAgreementPets          *sql.Stmt
	GetAllRentalAgreements             *sql.Stmt
	GetAllRentalAgreementsByRange      *sql.Stmt
	GetAllRentalAgreementTemplates     *sql.Stmt
	GetAllSingleInstanceAssessments    *sql.Stmt
	GetAllStringLists                  *sql.Stmt
	GetAllTransactants                 *sql.Stmt
	GetAllTransactantsForBID           *sql.Stmt
	GetAssessment                      *sql.Stmt
	GetAssessmentDuplicate             *sql.Stmt
	GetAssessmentInstance              *sql.Stmt
	GetAssessmentType                  *sql.Stmt
	GetAssessmentTypeByName            *sql.Stmt
	GetBuilding                        *sql.Stmt
	GetBusiness                        *sql.Stmt
	GetBusinessByDesignation           *sql.Stmt
	GetCustomAttribute                 *sql.Stmt
	GetCustomAttributeByVals           *sql.Stmt
	GetCustomAttributeRef              *sql.Stmt
	GetCustomAttributeRefs             *sql.Stmt
	GetDefaultLedgers                  *sql.Stmt
	GetDemandSource                    *sql.Stmt
	GetDemandSourceByName              *sql.Stmt
	GetDeposit                         *sql.Stmt
	GetDepositMethod                   *sql.Stmt
	GetDepositMethodByName             *sql.Stmt
	GetDepository                      *sql.Stmt
	GetDepositoryByAccount             *sql.Stmt
	GetDepositParts                    *sql.Stmt
	GetInvoice                         *sql.Stmt
	GetInvoiceAssessments              *sql.Stmt
	GetInvoicePayors                   *sql.Stmt
	GetJournal                         *sql.Stmt
	GetJournalAllocation               *sql.Stmt
	GetJournalAllocations              *sql.Stmt
	GetJournalByRange                  *sql.Stmt
	GetJournalByReceiptID              *sql.Stmt
	GetJournalMarker                   *sql.Stmt
	GetJournalMarkers                  *sql.Stmt
	GetJournalVacancy                  *sql.Stmt
	GetLatestLedgerMarkerByLID         *sql.Stmt
	GetLedger                          *sql.Stmt
	GetLedgerByGLNo                    *sql.Stmt
	GetLedgerByType                    *sql.Stmt
	GetLedgerEntriesForRAID            *sql.Stmt
	GetLedgerEntriesForRentable        *sql.Stmt
	GetLedgerEntriesInRange            *sql.Stmt
	GetLedgerEntriesInRangeByGLNo      *sql.Stmt
	GetLedgerEntriesInRangeByLID       *sql.Stmt
	GetLedgerEntry                     *sql.Stmt
	GetLedgerEntryByJAID               *sql.Stmt
	GetLedgerList                      *sql.Stmt
	GetLedgerMarkerByDateRange         *sql.Stmt
	GetLedgerMarkerByLIDDateRange      *sql.Stmt
	GetLedgerMarkerOnOrBefore          *sql.Stmt
	GetLedgerMarkers                   *sql.Stmt
	GetNote                            *sql.Stmt
	GetNoteAndChildNotes               *sql.Stmt
	GetNoteList                        *sql.Stmt
	GetNoteListMembers                 *sql.Stmt
	GetNoteType                        *sql.Stmt
	GetPaymentTypeByName               *sql.Stmt
	GetPaymentTypesByBusiness          *sql.Stmt
	GetPayor                           *sql.Stmt
	GetProspect                        *sql.Stmt
	GetRALedgerMarkerOnOrBefore        *sql.Stmt
	GetRatePlan                        *sql.Stmt
	GetRatePlanByName                  *sql.Stmt
	GetRatePlanRef                     *sql.Stmt
	GetRatePlanRefRTRate               *sql.Stmt
	GetRatePlanRefsInRange             *sql.Stmt
	GetRatePlanRefSPRate               *sql.Stmt
	GetReceipt                         *sql.Stmt
	GetReceiptAllocations              *sql.Stmt
	GetReceiptDuplicate                *sql.Stmt
	GetReceiptsInDateRange             *sql.Stmt
	GetReceiptsInRAIDDateRange         *sql.Stmt
	GetRecurringAssessmentsByBusiness  *sql.Stmt
	GetRentable                        *sql.Stmt
	GetRentableByName                  *sql.Stmt
	GetRentableLedgerMarkerOnOrBefore  *sql.Stmt
	GetRentableMarketRates             *sql.Stmt
	GetRentableSpecialtyRefs           *sql.Stmt
	GetRentableSpecialtyRefsByRange    *sql.Stmt
	GetRentableSpecialtyType           *sql.Stmt
	GetRentableSpecialtyTypeByName     *sql.Stmt
	GetRentableStatusByRange           *sql.Stmt
	GetRentableType                    *sql.Stmt
	GetRentableTypeByStyle             *sql.Stmt
	GetRentableTypeRefsByRange         *sql.Stmt
	GetRentableUsers                   *sql.Stmt
	GetRentalAgreement                 *sql.Stmt
	GetRentalAgreementByBusiness       *sql.Stmt
	GetRentalAgreementByRATemplateName *sql.Stmt
	GetRentalAgreementPayors           *sql.Stmt
	GetRentalAgreementPet              *sql.Stmt
	GetRentalAgreementRentables        *sql.Stmt
	GetRentalAgreementsForRentable     *sql.Stmt
	GetRentalAgreementTax              *sql.Stmt
	GetRentalAgreementTemplate         *sql.Stmt
	GetSecurityDepositAssessment       *sql.Stmt
	GetSLString                        *sql.Stmt
	GetSLStrings                       *sql.Stmt
	GetStringList                      *sql.Stmt
	GetStringListByName                *sql.Stmt
	GetTransactant                     *sql.Stmt
	GetUnitAssessments                 *sql.Stmt
	GetUser                            *sql.Stmt
	GetVehicle                         *sql.Stmt
	GetVehiclesByBID                   *sql.Stmt
	GetVehiclesByLicensePlate          *sql.Stmt
	GetVehiclesByTransactant           *sql.Stmt
	InsertAssessment                   *sql.Stmt
	InsertAssessmentType               *sql.Stmt
	InsertBuilding                     *sql.Stmt
	InsertBuildingWithID               *sql.Stmt
	InsertBusiness                     *sql.Stmt
	InsertCustomAttribute              *sql.Stmt
	InsertCustomAttributeRef           *sql.Stmt
	InsertDemandSource                 *sql.Stmt
	InsertDeposit                      *sql.Stmt
	InsertDepositMethod                *sql.Stmt
	InsertDepository                   *sql.Stmt
	InsertDepositPart                  *sql.Stmt
	InsertInvoice                      *sql.Stmt
	InsertInvoiceAssessment            *sql.Stmt
	InsertInvoicePayor                 *sql.Stmt
	InsertJournal                      *sql.Stmt
	InsertJournalAllocation            *sql.Stmt
	InsertJournalMarker                *sql.Stmt
	InsertLedger                       *sql.Stmt
	InsertLedgerAllocation             *sql.Stmt
	InsertLedgerEntry                  *sql.Stmt
	InsertLedgerMarker                 *sql.Stmt
	InsertNote                         *sql.Stmt
	InsertNoteList                     *sql.Stmt
	InsertNoteType                     *sql.Stmt
	InsertPaymentType                  *sql.Stmt
	InsertPayor                        *sql.Stmt
	InsertProspect                     *sql.Stmt
	InsertRatePlan                     *sql.Stmt
	InsertRatePlanRef                  *sql.Stmt
	InsertRatePlanRefRTRate            *sql.Stmt
	InsertRatePlanRefSPRate            *sql.Stmt
	InsertReceipt                      *sql.Stmt
	InsertReceiptAllocation            *sql.Stmt
	InsertRentable                     *sql.Stmt
	InsertRentableMarketRates          *sql.Stmt
	InsertRentableSpecialtyRef         *sql.Stmt
	InsertRentableSpecialtyType        *sql.Stmt
	InsertRentableStatus               *sql.Stmt
	InsertRentableType                 *sql.Stmt
	InsertRentableTypeRef              *sql.Stmt
	InsertRentableUser                 *sql.Stmt
	InsertRentalAgreement              *sql.Stmt
	InsertRentalAgreementPayor         *sql.Stmt
	InsertRentalAgreementPet           *sql.Stmt
	InsertRentalAgreementRentable      *sql.Stmt
	InsertRentalAgreementTax           *sql.Stmt
	InsertRentalAgreementTemplate      *sql.Stmt
	InsertSLString                     *sql.Stmt
	InsertStringList                   *sql.Stmt
	InsertTransactant                  *sql.Stmt
	InsertUser                         *sql.Stmt
	InsertVehicle                      *sql.Stmt
	ReadRatePlan                       *sql.Stmt
	ReadRatePlanRef                    *sql.Stmt
	UpdateAssessment                   *sql.Stmt
	UpdateBusiness                     *sql.Stmt
	UpdateDemandSource                 *sql.Stmt
	UpdateDeposit                      *sql.Stmt
	UpdateDepositMethod                *sql.Stmt
	UpdateDepository                   *sql.Stmt
	UpdateeRentalAgreementTax          *sql.Stmt
	UpdateInvoice                      *sql.Stmt
	UpdateLedger                       *sql.Stmt
	UpdateLedgerMarker                 *sql.Stmt
	UpdateNote                         *sql.Stmt
	UpdateNoteType                     *sql.Stmt
	UpdateProspect                     *sql.Stmt
	UpdateRatePlan                     *sql.Stmt
	UpdateRatePlanRef                  *sql.Stmt
	UpdateRatePlanRefRTRate            *sql.Stmt
	UpdateRatePlanRefSPRate            *sql.Stmt
	UpdateRentableSpecialtyRef         *sql.Stmt
	UpdateRentableStatus               *sql.Stmt
	UpdateRentableTypeRef              *sql.Stmt
	UpdateRentalAgreement              *sql.Stmt
	UpdateRentalAgreementPet           *sql.Stmt
	UpdateSLString                     *sql.Stmt
	UpdateStringList                   *sql.Stmt
	UpdateTransactant                  *sql.Stmt
	UpdateVehicle                      *sql.Stmt
	UpdateUser                         *sql.Stmt
	UpdatePayor                        *sql.Stmt

	// GetJournalInstance                 *sql.Stmt
	// GetSecDepBalanceLedger             *sql.Stmt
	// GetLedgerMarkerByRAID              *sql.Stmt
}

// PBprepSQL is the structure of prepared sql statements for the Phonebook db
type PBprepSQL struct {
	GetCompanyByDesignation      *sql.Stmt
	GetCompany                   *sql.Stmt
	GetBusinessUnitByDesignation *sql.Stmt
}

// BusinessTypeLists is a struct holding a collection of Types associated with a business
type BusinessTypeLists struct {
	BID          int64
	PmtTypes     map[int64]*PaymentType // payment types accepted
	DefaultAccts map[int64]*GLAccount   // index by the predifined contants DFAC*, value = GL No of that account
	GLAccounts   map[int64]GLAccount    // all the accounts for this business
	NoteTypes    []NoteType             // all defined note types for this business
}

// RRdb is a struct with all variables needed by the db infrastructure
var RRdb struct {
	Prepstmt RRprepSQL
	PBsql    PBprepSQL
	Dbdir    *sql.DB // phonebook db
	Dbrr     *sql.DB //rentroll db
	BizTypes map[int64]*BusinessTypeLists
}

// InitDBHelpers initializes the db infrastructure
func InitDBHelpers(dbrr, dbdir *sql.DB) {
	RRdb.Dbdir = dbdir
	RRdb.Dbrr = dbrr
	RRdb.BizTypes = make(map[int64]*BusinessTypeLists, 0)
	buildPreparedStatements()
	buildPBPreparedStatements()
	// RRdb.AsmtTypes = GetAssessmentTypes()
}

// InitBusinessFields initialize the lists in rlib's internal data structures
func InitBusinessFields(bid int64) {
	if nil == RRdb.BizTypes[bid] {
		bt := BusinessTypeLists{
			BID:          bid,
			PmtTypes:     make(map[int64]*PaymentType),
			DefaultAccts: make(map[int64]*GLAccount),
			GLAccounts:   make(map[int64]GLAccount),
		}
		RRdb.BizTypes[bid] = &bt
	}
}

// InitBizInternals initializes several internal structures with information about the business.
func InitBizInternals(bid int64, xbiz *XBusiness) {
	// fmt.Printf("Entered InitBizInternals\n")
	GetXBusiness(bid, xbiz) // get its info
	InitBusinessFields(bid)
	GetDefaultLedgers(bid) // Gather its chart of accounts
	RRdb.BizTypes[bid].GLAccounts = GetGLAccountMap(bid)
	GetAllNoteTypes(bid)
	LoadRentableTypeCustomaAttributes(xbiz)
}
