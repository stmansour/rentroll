package rlib

import (
	"context"
	"database/sql"
	"encoding/json"
	"extres"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// NO and all the rest are constants that are used with the RentRoll database
const (
	NO  = int64(0) // std negative value
	YES = int64(1)

	RECURNONE      = 0
	RECURSECONDLY  = 1
	RECURMINUTELY  = 2
	RECURHOURLY    = 3
	RECURDAILY     = 4
	RECURWEEKLY    = 5
	RECURMONTHLY   = 6
	RECURQUARTERLY = 7
	RECURYEARLY    = 8
	RECURLAST      = RECURYEARLY

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

	// ARASSESSMENT et al, are Account Rule Types.
	ARASSESSMENT    = 0
	ARRECEIPT       = 1
	AREXPENSE       = 2
	ARSUBASSESSMENT = 3

	// ASMUNPAID et al are flags for assessment
	ASMUNPAID      = 0
	ASMPARTIALPAID = 1
	ASMFULLYPAID   = 2
	ASMREVERSED    = 4

	// RCPTUNALLOCATED et al are flags for receipt
	RCPTUNALLOCATED      = 0
	RCPTPARTIALALLOCATED = 1
	RCPTFULLYALLOCATED   = 2
	RCPTREVERSED         = 4

	// RTACTIVE et all are flags for rentableTypes
	RTACTIVE   = 0
	RTINACTIVE = 1

	// CUSTSTRING et al are Custom Attribute types
	CUSTSTRING = 0
	CUSTINT    = 1
	CUSTUINT   = 2
	CUSTFLOAT  = 3
	CUSTDATE   = 4
	CUSTLAST   = 4 // this should be maintained as matching the highest index value in the group

	// LMOPEN etc all are ledger marker states
	LMOPEN    = 0
	LMCLOSED  = 1
	LMLOCKED  = 2
	LMINITIAL = 3

	ACCTSTATUSINACTIVE = 1
	ACCTSTATUSACTIVE   = 2
	RAASSOCIATED       = 1
	RAUNASSOCIATED     = 2

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

	// USESTATUSunknown etc all Rentable Use Status
	USESTATUSunknown            = 0
	USESTATUSinService          = 1
	USESTATUSadmin              = 2
	USESTATUSemployee           = 3
	USESTATUSownerOccupied      = 4
	USESTATUSofflineRenovation  = 5
	USESTATUSofflineMaintenance = 6
	USESTATUSmodel              = 7
	USESTATUSLAST               = 7

	// MRSTATUShouseKeeping etc all Rentable Make Ready Status
	MRSTATUShouseKeeping = 1
	MRSTATUSmaintenance  = 2
	MRSTATUSinspection   = 3
	MRSTATUSready        = 4

	// LEASESTATUSvacantRented etc all Rentable Lease Status
	LEASESTATUSvacantRented      = 1
	LEASESTATUSvacantNotRented   = 2
	LEASESTATUSonNoticePreleased = 3
	LEASESTATUSonNoticeAvailable = 4
	LEASESTATUSleased            = 5
	LEASESTATUSunavailable       = 6

	CREDIT = 0
	DEBIT  = 1

	RTRESIDENCE = 1
	RTCARPORT   = 2
	RTCAR       = 3

	REPORTJUSTIFYLEFT  = 0
	REPORTJUSTIFYRIGHT = 1

	JNLTYPEUNAS = 0 // record is unassociated with any assessment or Receipt
	JNLTYPEASMT = 1 // record is the result of an Assessment
	JNLTYPERCPT = 2 // record is the result of a Receipt
	JNLTYPEEXP  = 3 // record is the result of an Expense
	JNLTYPEXFER = 4 // funds transfer between accounts

	JOURNALTYPEASMID  = 1
	JOURNALTYPERCPTID = 2

	// RRDATEFMT is a shorthand date format used for text output
	// Use these values:	Mon Jan 2 15:04:05 MST 2006
	// const RRDATEFMT = "02-Jan-2006 3:04PM MST"
	// const RRDATEFMT = "01/02/06 3:04PM MST"
	RRDATEFMT         = "01/02/06"
	RRDATEFMT2        = "1/2/06"
	RRDATEFMT3        = "1/2/2006"
	RRDATEFMT4        = "01/02/2006"
	RRDATEINPFMT      = "2006-01-02"
	RRDATEFMTSQL      = RRDATEINPFMT
	RRDATETIMESQL     = "2006-01-02 15:04:05"
	RRJSUTCDATETIME   = "Mon, 02 Jan 2006 15:04:05 MST"
	RRDATETIMEINPFMT  = "2006-01-02 15:04:00 MST"
	RRDATETIMEFMT     = "2006-01-02T15:04:00Z"
	RRDATETIMEW2UIFMT = "1/2/2006 3:04 pm"
	RRDATEREPORTFMT   = "Jan 2, 2006"
	RRDATETIMERPTFMT  = "Jan 2, 2006 3:04pm MST"
	RRDATERECEIPTFMT  = "January 2, 2006"
)

// ARTypesList is the readable, csv loadable names for the different rule types
var ARTypesList = []string{"Assessment", "Receipt", "Expense", "Sub-Assessment"}

// TIME0 is the "beginning of time" constant to use when we need
// to set a time far enough in the past so that there won't be a
// date prior issue
var TIME0 = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

// ENDOFTIME can be used when there is no end time
var ENDOFTIME = time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)

// Period describes a span of time by specifying a start
// and end time.
type Period struct {
	D1, D2  time.Time
	Checked bool // used by Period overlap check functions
}

// ClosePeriod defines a date and tasklist associated with a period close
type ClosePeriod struct {
	CPID        int64
	BID         int64
	TLID        int64
	Dt          time.Time
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time
	CreateBy    int64
}

// Task is an indivually tracked work item.
// FLAGS are defined as follows:
//    1<<0 pre-completion required (if 0 then there is no pre-completion required)
//    1<<1 PreCompletion done (if 0 it is not yet done)
//    1<<2 Completion done (if 0 it is not yet done)
type Task struct {
	TID       int64
	BID       int64
	TLID      int64     // the TaskList to which this task belongs
	Name      string    // Task text
	Worker    string    // Name of the associated work function
	DtDue     time.Time // Task Due Date
	DtPreDue  time.Time // Pre Completion due date
	DtDone    time.Time // Task completion Date
	DtPreDone time.Time // Task Pre Completion Date

	// 1<<1 - 0 = DtPreDue should not be checked, 1 = DtPreDue should be checked
	// 1<<2 - 0 = DtDue should not be checked, 1 = DtDue should be checked
	FLAGS       int64
	DoneUID     int64     // user who marked task as done
	PreDoneUID  int64     // user who marked task as predone
	Comment     string    // any user comments
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// TaskList is the shell container for a list of tracked tasks
type TaskList struct {
	TLID      int64
	BID       int64
	PTLID     int64 // parent TLDID or 0 if this is the parent of a list
	TLDID     int64 // what type of task list... the definition
	Name      string
	Cycle     int64
	DtDue     time.Time
	DtPreDue  time.Time
	DtDone    time.Time
	DtPreDone time.Time

	// 1<<0 : 0 = active, 1 = inactive
	// 1<<1 : 0 = task list definition does not have a PreDueDate, 1 = has a PreDueDate
	// 1<<2 : 0 = task list definition does not have a DueDate, 1 = has a DueDate
	// 1<<3 : 0 = DtPreDone has not been set, 1 = DtPreDone has been set
	// 1<<4 : 0 = DtDone has not been set, 1 = DtDone has been set
	// 1<<5 : 0 = DtLastNotify has not been set, 1 = it has been set
	FLAGS        int64
	DoneUID      int64         // user who marked task as done
	PreDoneUID   int64         // user who marked task as predone
	EmailList    string        // email to this list when due date arrives
	DtLastNotify time.Time     // valid when FLAGS & 32 > 0, last time late notification was sent
	DurWait      time.Duration // amount of time to wait before next check after late notification
	Comment      string        // any user comments
	CreateTS     time.Time     // when was this record created
	CreateBy     int64         // employee UID (from phonebook) that created it
	LastModTime  time.Time     // when was this record last written
	LastModBy    int64         // employee UID (from phonebook) that modified it
}

// TaskDescriptor is the definition of a task. It is used to make instance
// which become Tasks
type TaskDescriptor struct {
	TDID        int64
	BID         int64
	TLDID       int64
	Name        string
	Worker      string
	EpochDue    time.Time
	EpochPreDue time.Time
	FLAGS       int64
	Comment     string    //
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// TaskListDefinition is the shell container for TaskDescriptors
type TaskListDefinition struct {
	TLDID       int64
	BID         int64
	Name        string
	Cycle       int64
	Epoch       time.Time     // when task list starts
	EpochDue    time.Time     // when task list is due
	EpochPreDue time.Time     // when task list pre-work is due
	FLAGS       int64         // 1<<0 0 means it is still active, 1 means it is no longer active
	EmailList   string        // email to this list when due date arrives
	DurWait     time.Duration // amount of time to wait before next check after late notification
	Comment     string        //
	CreateTS    time.Time     // when was this record created
	CreateBy    int64         // employee UID (from phonebook) that created it
	LastModTime time.Time     // when was this record last written
	LastModBy   int64         // employee UID (from phonebook) that modified it
}

// AIRAuthenticateResponse is the reply structure from Accord Directory
type AIRAuthenticateResponse struct {
	Status   string       `json:"status"`
	UID      int64        `json:"uid"`
	Username string       `json:"username"` // user's first or preferred name
	Name     string       `json:"Name"`
	ImageURL string       `json:"ImageURL"`
	Message  string       `json:"message"`
	Token    string       `json:"Token"`
	Expire   JSONDateTime `json:"Expire"` // DATETIMEFMT in this format "2006-01-02T15:04:00Z"
}

// StringList is a generic list structure for lists of strings. These could be
// used to implement things like the list of reasons why an applicant's
// application was turned down, the list of reasons why a tenant is moving out,
// etc.
type StringList struct {
	SLID        int64      // unique id for this stringlist
	BID         int64      // the business to which this stringlist belongs
	Name        string     // stringlist name
	LastModTime time.Time  // when was this record last written
	LastModBy   int64      // employee UID (from phonebook) that modified it
	S           []SLString // array of SLStrings associated with this SLID
	CreateTS    time.Time  // when was this record created
	CreateBy    int64      // employee UID (from phonebook) that created it
}

// SLString defines an individual string member of a StringList
type SLString struct {
	SLSID       int64     // unique id of this string
	BID         int64     // the business to which this stringlist belongs
	SLID        int64     // to which stringlist does this string belong?
	Value       string    // value of this string
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// NoteType describes the type of note this is
type NoteType struct {
	NTID        int64     // note type id
	BID         int64     // business associated with this note type
	Name        string    // the actual note
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// Note is dated comment from a user
type Note struct {
	NID         int64     // unique ID for this note
	BID         int64     // business associated with this note type
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
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// NoteList is a collection of Notes (NIDs)
type NoteList struct {
	NLID        int64     // unique id for the notelist
	BID         int64     // business associated with this notelist
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	N           []Note    // the list of notes
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// CustomAttribute is a struct containing user-defined custom attributes for objects
type CustomAttribute struct {
	CID         int64     // unique id
	BID         int64     // business associated with this CustomAttribute
	Type        int64     // what type of value: 0 = string, 1 = int64, 2 = float64
	Name        string    // what its called
	Value       string    // string value -- will be xlated on load / store
	Units       string    // optional units value.  Ex:  "feet", "gallons", "cubic feet", ...
	LastModTime time.Time // timestamp of last changed
	LastModBy   int64     // who changed it last
	fval        float64   // the float value once converted
	ival        int64     // the int value once converted
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// CustomAttributeRef is a reference to a Custom Attribute. A query of the form:
//		SELECT CID FROM CustomAttributeRef
type CustomAttributeRef struct {
	CARID       int64     // unique id
	ElementType int64     // what type of element:  1=person, 2=company, 3=Business-unit, 4 = executable service, 5=RentableType
	BID         int64     // business associated with this CustomAttributeRef
	ID          int64     // the UID of the element type. That is, if ElementType == 5, the ID is the RTID (Rentable type id)
	CID         int64     // uid of the custom attribute
	LastModTime time.Time // timestamp of last changed
	LastModBy   int64     // who changed it last
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// AssessmentType is a list of charges that the company can make
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
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentalAgreementTemplate is a template used to set up new rental agreements
type RentalAgreementTemplate struct {
	RATID          int64     // unique id for this rate plan
	BID            int64     // which business
	RATemplateName string    // RATemplateName a string associated with each rental type agreement (essentially, the doc name)
	LastModTime    time.Time // when was this record last written
	LastModBy      int64     // employee UID (from phonebook) that modified it
	CreateTS       time.Time // when was this record created
	CreateBy       int64     // employee UID (from phonebook) that created it
}

// RentalAgreementGrid is a struct for the Rental Agreement Grid in the UI
type RentalAgreementGrid struct {
	Recid          int
	RAID           int64
	TCIDPayor      int64
	AgreementStart JSONDate
	AgreementStop  JSONDate
}

// RatePlan is a structure of static attributes of a rate plan, which describes charges for rentable types, varying by customer
type RatePlan struct {
	RPID        int64               // unique id for this rate plan
	BID         int64               // Business
	Name        string              // The name of this RatePlan
	LastModTime time.Time           // when was this record last written
	LastModBy   int64               // employee UID (from phonebook) that modified it
	OD          []OtherDeliverables // other Deliverables associated with the rate plan
	CreateTS    time.Time           // when was this record created
	CreateBy    int64               // employee UID (from phonebook) that created it
}

// FlRatePlanGDS and the others are bit flags for the RatePlan custom attribute: FLAGS
const (
	FlRatePlanGDS   = 1 << 0 // RatePlan - bit 00 export to GDS
	FlRatePlanSabre = 1 << 1 // RatePlan - bit 01 export to Sabre
)

// RatePlanRef contains the time sensitive attributes of a RatePlan
type RatePlanRef struct {
	RPRID             int64               // unique id for this ref
	BID               int64               // Business
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
	CreateTS          time.Time           // when was this record created
	CreateBy          int64               // employee UID (from phonebook) that created it
}

// FlRTRRefHide and the others are bit flags for the RatePlanRef
const (
	FlRTRRefHide = 1 << 0 // do not show this rate plan to users
)

// RatePlanRefRTRate is RatePlan RPRID's rate information for the RentableType (RTID)
type RatePlanRefRTRate struct {
	RPRRTRateID int64     // unique id
	RPRID       int64     // which RatePlanRef is this
	BID         int64     // Business
	RTID        int64     // which RentableType
	FLAGS       uint64    // 1<<0 = percent flag 0 = Val is an absolute price, 1 = percent of MarketRate,
	Val         float64   // Val
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// FlRTRpct and the others are bit flags for the RatePlanRefRTRate
const (
	FlRTRpct = 1 << 0 // bit 0 = percent flag. 0 means it's an absolute amount, 1 means it's a % of Market Rate
	FlRTRna  = 1 << 1 // bit 1 = n/a flag, 0 means that this RTID is affected, 1 means it is not affected
)

// RatePlanRefSPRate is RatePlan RPRID's rate information for the Specialties
type RatePlanRefSPRate struct {
	RPRSPRateID int64     // unique id
	RPRID       int64     // which RatePlanRef is this
	BID         int64     // Business
	RTID        int64     // which RentableType
	RSPID       int64     // which Specialty
	FLAGS       uint64    // 1<<0 = percent flag 0 = Val is an absolute price, 1 = percent of MarketRate,
	Val         float64   // Val
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// FlSPRpct and the others are bit flags for the RatePlanRefSPRate
const (
	FlSPRpct = 1 << 0 // bit 0 = percent flag. 0 means it's an absolute amount, 1 means it's a % of Market Rate
	FlSPRna  = 1 << 1 // bit 1 = n/a flag, 0 means that this RTID is affected, 1 means it is not affected
)

// RatePlanOD defines which other deliverables are associated with a RatePlan.
// A RatePlan can refer to multiple OtherDeliverables.
type RatePlanOD struct {
	RPRID       int64     // with which RatePlan is this OtherDeliverable associated?
	BID         int64     // Business
	ODID        int64     // points to an OtherDeliverables
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// OtherDeliverables defines special offers associated with RatePlanRefs. These are for promotions. Examples of OtherDeliverables
// would include things like 2 Seaworld tickets, etc.  Referenced by RatePlanRef
// Multiple RatePlanRefs can refer to the same OtherDeliverables.
type OtherDeliverables struct {
	ODID        int64     // Unique ID for this OtherDeliverables
	BID         int64     // Business
	Name        string    // Description of the other deliverables. Ex: 2 Seaworld tickets
	Active      bool      // Flag: Is this list still active?  dropdown interface lists only the active ones
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentalAgreement binds one or more payors to one or more rentables
type RentalAgreement struct {
	Recid                  int64       `json:"recid"` // this is to support the grid widget
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
	ExtensionOptionNotice  time.Time   // the last date by which a Tenant can give notice of their intention to exercise the right to an extension option period
	ExpansionOption        string      // the right to expand to certanin spaces that are typically contiguous to their primary space
	ExpansionOptionNotice  time.Time   // the last date by which a Tenant can give notice of their intention to exercise the right to an Expansion Option
	RightOfFirstRefusal    string      // Tenant may have the right to purchase their premises if LL chooses to sell
	FLAGS                  uint64      // 1<<0 - is application pending approval,
	LastModTime            time.Time   // when was this record last written
	LastModBy              int64       // employee UID (from phonebook) that modified it
	CreateTS               time.Time   // when was this record created
	CreateBy               int64       // employee UID (from phonebook) that created it
	R                      []XRentable // all the rentables
	P                      []XPerson   // all the payors
	T                      []XPerson   // all the users
}

// RentalAgreementRentable describes a Rentable associated with a rental agreement
type RentalAgreementRentable struct {
	RARID        int64     // unique id
	RAID         int64     // associated rental agreement
	BID          int64     // Business
	RID          int64     // the Rentable
	CLID         int64     // commission ledger -- applies if outside sales rented this rentable
	ContractRent float64   // the rent
	RARDtStart   time.Time // start date/time for this Rentable
	RARDtStop    time.Time // stop date/time
	LastModTime  time.Time // when was this record last written
	LastModBy    int64     // employee UID (from phonebook) that modified it
	CreateTS     time.Time // when was this record created
	CreateBy     int64     // employee UID (from phonebook) that created it
}

// RentalAgreementPayor describes a Payor associated with a rental agreement
type RentalAgreementPayor struct {
	RAPID       int64 // unique id
	RAID        int64
	BID         int64     // Business
	TCID        int64     // the payor's transactant id
	DtStart     time.Time // start date/time for this Payor
	DtStop      time.Time // stop date/time
	FLAGS       uint64    // 1<<0 is the bit that indicates this payor is a 'guarantor'
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentalAgreementTax - the time based attribute for whether the rental agreement is taxable
type RentalAgreementTax struct {
	RAID        int64     //associated rental agreement
	BID         int64     // Business
	DtStart     time.Time // start date/time for this Payor
	DtStop      time.Time // stop date/time
	FLAGS       uint64    // 1<<0 is whether the agreement is taxable
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentableUser describes a User associated with a rental agreement
type RentableUser struct {
	RUID        int64     // unique id
	RID         int64     // associated Rentable
	BID         int64     // associated business
	TCID        int64     // pointer to Transactant
	DtStart     time.Time // start date/time for this User
	DtStop      time.Time // stop date/time (when this person stopped being a User)
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentalAgreementPet describes a pet associated with a rental agreement. There can be as many as needed.
type RentalAgreementPet struct {
	Recid       int64 `json:"recid"` // support w2ui grid
	PETID       int64
	BID         int64 // associated business
	RAID        int64 // deprecated
	TCID        int64 // contact person
	Type        string
	Breed       string
	Color       string
	Weight      float64
	Name        string
	DtStart     time.Time
	DtStop      time.Time
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// DemandSource is a structure
type DemandSource struct {
	SourceSLSID int64     // DemandSource ID
	BID         int64     // Business unit
	Name        string    // name of source
	Industry    string    // what industry is this source in
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
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
	IsCompany      bool   // 1 => the entity is a company, 0 = not a company
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
	/*
	   FLAG BITS:
	   1<<0 OptIntoMarketingCampaign -- Does the user want to receive mkting info
	   1<<1 AcceptGeneralEmail       -- Will user accept email
	   1<<2 VIP                      -- Is this person a VIP
	*/
	FLAGS       int64
	Comment     string
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// Prospect contains info over and above
type Prospect struct {
	TCID                     int64
	BID                      int64
	CompanyAddress           string
	CompanyCity              string
	CompanyState             string
	CompanyPostalCode        string
	CompanyEmail             string
	CompanyPhone             string
	Occupation               string
	DesiredUsageStartDate    time.Time // predicted rent start date
	RentableTypePreference   int64     // RentableType
	FLAGS                    uint64    // 0 = Approved/NotApproved,
	EvictedDes               string    // explanation when FLAGS & (1<<2) > 0
	ConvictedDes             string    // explanation when FLAGS & (1<<3) > 0
	BankruptcyDes            string    // explanation when FLAGS & (1<<4) > 0
	Approver                 int64     // UID from Directory
	DeclineReasonSLSID       int64     // SLSid of reason
	OtherPreferences         string    // arbitrary text
	FollowUpDate             time.Time // automatically fill out this date to sysdate + 24hrs
	CSAgent                  int64     // Accord Directory UserID - for the CSAgent
	OutcomeSLSID             int64     // id of string from a list of outcomes. Melissa to provide reasons
	CurrentAddress           string
	CurrentLandLordName      string
	CurrentLandLordPhoneNo   string
	CurrentReasonForMoving   int64
	CurrentLengthOfResidency string
	PriorAddress             string
	PriorLandLordName        string
	PriorLandLordPhoneNo     string
	PriorReasonForMoving     int64
	PriorLengthOfResidency   string
	CommissionableThirdParty string
	LastModTime              time.Time
	LastModBy                int64
	CreateTS                 time.Time // when was this record created
	CreateBy                 int64     // employee UID (from phonebook) that created it
}

// User contains all info common to a person
type User struct {
	TCID                      int64
	BID                       int64
	Points                    int64
	DateofBirth               time.Time
	EmergencyContactName      string
	EmergencyContactAddress   string
	EmergencyContactTelephone string
	EmergencyContactEmail     string
	AlternateAddress          string
	EligibleFutureUser        bool
	FLAGS                     uint64
	Industry                  string
	SourceSLSID               int64
	LastModTime               time.Time
	LastModBy                 int64
	Vehicles                  []Vehicle
	CreateTS                  time.Time // when was this record created
	CreateBy                  int64     // employee UID (from phonebook) that created it
}

// Payor is attributes of the person financially responsible
// for the rent.
type Payor struct {
	TCID                int64
	BID                 int64
	CreditLimit         float64
	TaxpayorID          string
	ThirdPartySource    int64
	EligibleFuturePayor bool
	FLAGS               uint64
	SSN                 string // encrypted in database, decrypted here
	DriversLicense      string // encrypted in database, decrypted here
	GrossIncome         float64
	LastModTime         time.Time
	LastModBy           int64
	CreateTS            time.Time // when was this record created
	CreateBy            int64     // employee UID (from phonebook) that created it
}

// XPerson of all person related attributes
type XPerson struct {
	Trn Transactant
	Usr User
	Psp Prospect
	Pay Payor
}

// TransactantTypeDown is the struct needed to match names in typedown controls
type TransactantTypeDown struct {
	TCID        int64
	FirstName   string
	MiddleName  string
	LastName    string
	CompanyName string
	IsCompany   bool
	Recid       int64 `json:"recid"`
}

// RentableTypeDown is the struct needed to match names in typedown controls
type RentableTypeDown struct {
	Recid        int64 `json:"recid"` // this will hold the RID
	RID          int64
	RentableName string
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
	VIN                 string
	LicensePlateState   string
	LicensePlateNumber  string
	ParkingPermitNumber string
	DtStart             time.Time
	DtStop              time.Time
	LastModTime         time.Time
	LastModBy           int64
	CreateTS            time.Time // when was this record created
	CreateBy            int64     // employee UID (from phonebook) that created it
}

// Assessment is a charge associated with a Rentable
type Assessment struct {
	ASMID          int64     // unique id for this assessment
	PASMID         int64     // parent Assessment, if this is non-zero it means this assessment is an instance of the recurring assessment with id PASMID. When non-zero DO NOT process as a recurring assessment, it is an instance
	RPASMID        int64     // reversal parent Assessment, if it is non-zero, then the assessment has been reversed.
	AGRCPTID       int64     // Auto-generator RCTPID is >0 when this assessment was autogenerated due to RCPTID's SubARs
	BID            int64     // what Business
	RID            int64     // the Rentable
	ATypeLID       int64     // DEPRECATED!!!  what type of assessment
	RAID           int64     // associated Rental Agreement
	Amount         float64   // how much
	Start          time.Time // start time
	Stop           time.Time // stop time, may be the same as start time or later
	RentCycle      int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, G = quarterly, 8 = yearly
	ProrationCycle int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
	InvoiceNo      int64     // A uniqueID for the invoice number
	AcctRule       string    // override ARID with this account rule
	ARID           int64     // reference to the account rule to use
	FLAGS          uint64    /* bits 0-1: 0 = unpaid,1 = partially paid, 2 = fully paid, 3 = it's an offset, do not apply payments to this assessment
	 * bit 2: 1 = this assmt has been reversed
	 */
	Comment     string
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// Expense is an amount that reduces some assessment
// for example, the bank fee associated with a wire transfer
type Expense struct {
	EXPID       int64
	RPEXPID     int64 // Reversal parent
	BID         int64
	RID         int64
	RAID        int64
	Amount      float64
	Dt          time.Time
	AcctRule    string
	ARID        int64
	FLAGS       uint64 // bit 2 = Reversed
	Comment     string
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time
	CreateBy    int64
}

// AR is the table that defines the AcctRules for Assessments, Expenses and Receipts
// FLAGS
//  1<<0 = apply funds to Receive accts,
//  1<<1 - populate on Rental Agreement,
//  1<<2 = RAID required,
//  1<<3 = subARIDs apply
type AR struct {
	ARID          int64
	BID           int64
	Name          string
	ARType        int64 // 0 = Assessment, 1 = Receipt, 2 = Expense
	DebitLID      int64
	CreditLID     int64
	Description   string
	RARequired    int64
	DtStart       time.Time
	DtStop        time.Time
	FLAGS         uint64
	DefaultAmount float64 // use this as the default amount in ui for newly created Assessments
	LastModTime   time.Time
	LastModBy     int64
	CreateTS      time.Time // when was this record created
	CreateBy      int64     // employee UID (from phonebook) that created it
	SubARs        []SubAR   // the SubARs if FLAGS & 1<<3 > 0
}

// SubAR is the table that defines multiple ARIDs for transactions that require multiple ARIDs
type SubAR struct {
	SARID       int64
	ARID        int64
	SubARID     int64
	BID         int64
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// Business is the set of attributes describing a rental or hotel Business
type Business struct {
	BID                   int64
	Designation           string // reference to designation in Phonebook db
	Name                  string
	DefaultRentCycle      int64     // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	DefaultProrationCycle int64     // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	DefaultGSRPC          int64     // Default for every Rentable Type, useful in initializing the UI for new RentableTypes
	ClosePeriodTLID       int64     // TaskList used for closing a period
	FLAGS                 int64     // FLAGS -- 1<<0 = 0 - EDI disabled, 1=EDI enabled
	LastModTime           time.Time // when was this record last written
	LastModBy             int64     // employee UID (from phonebook) that modified it
	// ParkingPermitInUse    int64     // yes/no  0 = no, 1 = yes
	CreateTS time.Time // when was this record created
	CreateBy int64     // employee UID (from phonebook) that created it
}

// BusinessProperties defines properties for a business. The value
// of the property is defined as JSON data in the Data field. It should
// be unmarshaled into a struct that corresponds to the Name
type BusinessProperties struct {
	BPID        int64
	BID         int64
	Name        string          // "general" or whatever sub-category you want
	Data        json.RawMessage // json data in mysql -- marshaled BizProps
	FLAGS       int64           // FLAGS
	LastModTime time.Time       // when was this record last written
	LastModBy   int64           // employee UID (from phonebook) that modified it
	CreateTS    time.Time       // when was this record created
	CreateBy    int64           // employee UID (from phonebook) that created it
}

// BizProps is the golang struct for a category of business properties.
// This struct will be marshaled into JSON data and stored in BusinessProperties
type BizProps struct {
	PetFees     []string // AR names of all Pet Fees
	VehicleFees []string // AR names of all Vehicle Fees
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
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// PaymentType describes how a payment was made
type PaymentType struct {
	PMTID       int64
	BID         int64
	Name        string
	Description string
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RCPTvoid and the others are bit flags for Receipt
const (
	RCPTvoid = 1 << 2 // bit 2 = void --> part of a voided receipt pair
)

// Receipt saves the information associated with a payment made by a User to cover one or more Assessments
type Receipt struct {
	RCPTID          int64     // unique id for this receipt
	PRCPTID         int64     // Parent RCPTID, points to RCPT being amended/corrected by this receipt
	BID             int64     // which business
	TCID            int64     // payor that sent in the payment - even if OtherPayorName is present this field must have the payor for whom the OtherPayorName is paying
	PMTID           int64     // what type of payment
	DEPID           int64     // the depository where this receipt will be deposited
	DID             int64     // the Deposit ID to which this receipt belongs
	RAID            int64     // required for special case receipts
	Dt              time.Time // date payment was received
	DocNo           string    // check number, money order number, etc.; documents the payment
	Amount          float64   // amount of the receipt
	AcctRuleReceive string    // Account rule to apply on the receipt of this payment -- essentially - bank account and unapplied funds
	ARID            int64     // User selected rule
	AcctRuleApply   string    // how the funds are applied to assessments
	FLAGS           uint64    // bits 0-1 : 0 unallocated, 1 = partially allocated, 2 = fully allocated; bit 2: part of a voided receipt pair
	Comment         string    // any notes on this receipt
	OtherPayorName  string    // if not '', the name of a payor who paid this receipt and who may not be in our system
	LastModTime     time.Time
	LastModBy       int64
	RA              []ReceiptAllocation
	CreateTS        time.Time // when was this record created
	CreateBy        int64     // employee UID (from phonebook) that created it
	RentableName    string    // RECEIPT-ONLY CLIENT. Remove this field when we no longer use the RECEIPT-ONLY CLIENT
}

// RECEIPTONLYCLIENT et. al. are values used to support the
// receipt-only client.
// TODO: remove when this client is no longer needed.
const (
	RECEIPTONLYCLIENT = "receipts"
	ROCPRE            = "^*{{"
	ROCPOST           = "}}*^"
	ROCOFFSET         = len(ROCPRE)
)

// ROCExtractRentableName is used to extract the rentable name
// from the comment field in a Receipt structure produced by the
// RECEIPT-ONLY client.
// TODO: remove when this client is no longer needed.
//
// Input - the comment string
// Returns -  rentableName, commentWithRentableNameRemoved
//-----------------------------------------------------------------
func ROCExtractRentableName(comment string) (string, string) {
	var rn string
	i1 := strings.Index(comment, ROCPRE)
	i2 := strings.Index(comment, ROCPOST)
	if i1 >= 0 && i2 > i1 {
		rn = comment[i1+ROCOFFSET : i2]
		comment = comment[:i1]
	}
	return rn, comment
}

// ReceiptAllocation defines an allocation of a Receipt amount.
type ReceiptAllocation struct {
	RCPAID      int64 // Receipt Allocation ID
	RCPTID      int64
	BID         int64
	RAID        int64     // which RAID is this portion of the payment associated
	Dt          time.Time // date of this payment (may not be the same as the Receipt's)
	Amount      float64
	ASMID       int64
	AcctRule    string
	FLAGS       uint64 // bit 2:  VOID THIS RECEIPT-ALLOCATION
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// Depository is a bank account or other account where deposits are made
type Depository struct {
	DEPID       int64     // unique id for a depository
	BID         int64     // which business
	LID         int64     // which GL Account represents this depository
	Name        string    // Name of Depository: First Data, Nyax, CCI, Oklahoma Fidelity
	AccountNo   string    // account number at this Depository
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// Deposit is simply a list of receipts that form a deposit to a Depository. This struct contains
// the static attributes of the list
type Deposit struct {
	DID           int64         // Unique id of this deposit
	BID           int64         // business id
	DEPID         int64         // Depository id where the deposit was made
	DPMID         int64         // Deposit method
	Dt            time.Time     // Date of deposit
	Amount        float64       // the total amount of the deposit
	ClearedAmount float64       // the amount cleared by the depository
	FLAGS         uint64        // bitflags
	LastModTime   time.Time     // when was this record last written
	LastModBy     int64         // employee UID (from phonebook) that modified it
	CreateTS      time.Time     // when was this record created
	CreateBy      int64         // employee UID (from phonebook) that created it
	DP            []DepositPart // array of DepositParts for this deposit
}

// DepositPart is a reference to a Receipt that is part of this deposit.  Another way of
// thinking about it is that this query produces the list of all receipts in a Deposit:
//		SELECT RCPTID WHERE DIP=someDID
type DepositPart struct {
	DPID        int64
	DID         int64     // deposit id
	BID         int64     // business id
	RCPTID      int64     // receipt id
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// DepositMethod is a list of methods used to make deposits to a depository
type DepositMethod struct {
	DPMID       int64     //the method id
	BID         int64     // business id
	Method      string    // descriptive name
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
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
	CreateTS    time.Time           // when was this record created
	CreateBy    int64               // employee UID (from phonebook) that created it
}

// InvoiceAssessment is a reference to an Assessment that is part of this invoice.  Another way of
// thinking about it is that this query produces the list of all assessments in an invoice:
//		SELECT ASMID WHERE InvoiceNo=somenumber
type InvoiceAssessment struct {
	InvoiceASMID int64     // unique id
	InvoiceNo    int64     // the invoice number
	BID          int64     // bid
	ASMID        int64     // assessment
	LastModTime  time.Time // when was this record last written
	LastModBy    int64     // employee UID (from phonebook) that modified it
	CreateTS     time.Time // when was this record created
	CreateBy     int64     // employee UID (from phonebook) that created it
}

// InvoicePayor is a reference to a Payor for this invoice.  Another way of
// thinking about it is that this query produces the list of all payors for an invoice:
//		SELECT PID WHERE InvoiceNo=somenumber
type InvoicePayor struct {
	InvoicePayorID int64     // unique id
	InvoiceNo      int64     // the invoice number
	BID            int64     // bid
	PID            int64     // Payor ID
	LastModTime    time.Time // when was this record last written
	LastModBy      int64     // employee UID (from phonebook) that modified it
	CreateTS       time.Time // when was this record created
	CreateBy       int64     // employee UID (from phonebook) that created it
}

// RentableSpecialty is the structure for attributes of a Rentable specialty
type RentableSpecialty struct {
	RSPID       int64
	BID         int64
	Name        string
	Fee         float64 // proration inherited from the rentable / rentable type.
	Description string
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentableType is the set of attributes describing the different types of Rentable items
type RentableType struct {
	RTID        int64                      // unique identifier for this RentableType
	BID         int64                      // the business unit to which this RentableType belongs
	Style       string                     // a short name
	Name        string                     // longer name
	RentCycle   int64                      // frequency at which rent accrues, 0 = not set or n/a, 1 = secondly, 2=minutely, 3=hourly, 4=daily, 5=weekly, 6=monthly...
	Proration   int64                      // frequency for prorating rent if the full rentcycle is not used
	GSRPC       int64                      // Time increments in which GSR is calculated to account for rate changes
	FLAGS       uint64                     // 0=active, 1=inactive
	ARID        int64                      // ARID reference, for default rent amount for this types
	MR          []RentableMarketRate       // array of time sensitive market rates
	CA          map[string]CustomAttribute // index by Name of attribute, associated custom attributes
	MRCurrent   float64                    // the current market rate (historical values are in MR)
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentableMarketRate describes the market rate rent for a Rentable type over a time period
type RentableMarketRate struct {
	RMRID       int64
	RTID        int64
	BID         int64 // the business unit
	MarketRate  float64
	DtStart     time.Time
	DtStop      time.Time
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentableTypeTax - the time based attribute for whether the rental agreement is taxable
type RentableTypeTax struct {
	RAID        int64     //associated rental agreement
	BID         int64     // Business
	DtStart     time.Time // start date/time for this Payor
	DtStop      time.Time // stop date/time
	TAXID       int64     // which tax in the Tax Table
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// Rentable is the basic struct for  entities to rent
type Rentable struct {
	Recid          int64             `json:"recid"` // this is to support the grid widget
	RID            int64             // unique id for this Rentable
	BID            int64             // Business
	RentableName   string            // name for this rentable
	AssignmentTime int64             // can we pre-assign or assign only at commencement
	MRStatus       int64             // Make Ready Status - current value as of DtMR, when this value changes it goes into a MRHistory record
	DtMRStart      time.Time         // Time that MRStatus was set
	Comment        string            // for notes such as Alarm codes and other things
	RT             []RentableTypeRef // the list of RTIDs and timestamps for this Rentable
	//-- RentalPeriodDefault int64          // 0 =unset, 1 = short term, 2=longterm
	LastModTime time.Time // time of last update to the db record
	LastModBy   int64     // who made the update (Phonebook UID)
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// MRHistory is the basic structure for Make Ready status history
type MRHistory struct {
	MRHID       int64     // unique id
	BID         int64     // which biz
	MRStatus    int64     // see definition in Rentable table field
	DtMRStart   time.Time // when the rentable went into this status
	DtMRStop    time.Time // when the rentable changed to a different status
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created this record
}

// RentableTypeRef is the time-based Rentable type attribute
type RentableTypeRef struct {
	RTRID                  int64     // unique ID for this Rentable Type Reference
	RID                    int64     // the Rentable to which this record belongs
	BID                    int64     // Business
	RTID                   int64     // the Rentable's type during this time range
	OverrideRentCycle      int64     // Override Rent Cycle.  0 =unset,  otherwise same values as RentableType.RentCycle
	OverrideProrationCycle int64     // Override Proration. 0 = unset, otherwise the same values as RentableType.Proration
	DtStart                time.Time // timerange start
	DtStop                 time.Time // timerange stop
	LastModTime            time.Time
	LastModBy              int64
	CreateTS               time.Time // when was this record created
	CreateBy               int64     // employee UID (from phonebook) that created it
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
	RSPRefID    int64     // unique id
	BID         int64     // associated business
	RID         int64     // the Rentable to which this record belongs
	RSPID       int64     // the rentable specialty type associated with the rentable
	DtStart     time.Time // timerange start
	DtStop      time.Time // timerange stop
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// RentableStatus archives the state of a Rentable during the specified period of time
type RentableStatus struct {
	RSID             int64     // unique ID for this Rentable Status
	RID              int64     // associated Rentable
	BID              int64     // associated business
	DtStart          time.Time // start of period
	DtStop           time.Time // end of period
	DtNoticeToVacate time.Time // user has indicated they will vacate on this date
	UseStatus        int64     // 1-Administrative, 2=InService, 3=Employee, 4=Model, 5=OfflineRennovation, 6=OfflineMaintenance
	LeaseStatus      int64     // 1-Vacant-rented, 2=VacantNotRented, 3=OnNoticePreLeased, 4=OnNoticeAvailable, 5=Leased, 6=Unavailable
	LastModTime      time.Time // time of last update to the db record
	LastModBy        int64     // who made the update (Phonebook UID)
	CreateTS         time.Time // when was this record created
	CreateBy         int64     // employee UID (from phonebook) that created it
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
	Dt          time.Time           // when this entry was made
	Amount      float64             // the amount
	Type        int64               // 0 = unassociated with RA, 1 means this is an assessment, 2 means it is a payment
	ID          int64               // if Type == 0 then it is the RentableID, if Type == 1 then it is the ASMID that caused this entry, if Type ==2 then it is the RCPTID
	Comment     string              // for notes like "prior period adjustment"
	LastModTime time.Time           // auto updated
	LastModBy   int64               // user making the mod
	JA          []JournalAllocation // an array of Journal allocations, breaks the payment or assessment down, total of all the allocations equals the "Amount" above
	//RAID        int64               // unique id of Rental Agreement
	CreateTS time.Time // when was this record created
	CreateBy int64     // employee UID (from phonebook) that created it
}

// JournalAllocation describes how the associated Journal amount is allocated
type JournalAllocation struct {
	JAID        int64     // unique id for this allocation
	BID         int64     // unique id of Business
	JID         int64     // associated Journal entry
	RID         int64     // associated Rentable
	RAID        int64     // associated Rental Agreement
	TCID        int64     // if > 0 this is the payor who made the payment - important if RID and RAID == 0 -- means the payment went to the unallocated funds account
	RCPTID      int64     // associated receipt if TCID > 0
	Amount      float64   // amount of this allocation
	ASMID       int64     // associated AssessmentID -- source of the charge
	EXPID       int64     // associated Expense -- source of the charge
	AcctRule    string    // describes how this amount distributed across the accounts
	LastModTime time.Time // when was this record last written
	LastModBy   int64     // employee UID (from phonebook) that modified it
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// JournalMarker describes a period of time where the Journal entries have been locked down
type JournalMarker struct {
	JMID        int64
	BID         int64
	State       int64
	DtStart     time.Time
	DtStop      time.Time
	LastModTime time.Time
	LastModBy   int64
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
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
	TCID        int64     // Payor associated with this entry
	Dt          time.Time // date associated with this transaction
	Amount      float64
	Comment     string    // for notes like "prior period adjustment"
	LastModTime time.Time // auto updated
	LastModBy   int64     // user making the mod
	//GLNo        string    // glnumber for the ledger -- DELETE THIS ATTRIBUTE
	CreateTS time.Time // when was this record created
	CreateBy int64     // employee UID (from phonebook) that created it
}

// LedgerMarker describes a period of time period described. The Balance can be
// used going forward from DtStop
type LedgerMarker struct {
	LMID        int64     // unique id for this LM
	LID         int64     // associated GLAccount
	BID         int64     // only valid if Type == 1
	RAID        int64     // if 0 then it's the LM for the whole account, if > 0 it's the amount for the rental agreement RAID
	RID         int64     // if 0 then it's the LM for the whole account, if > 0 it's the amount for the Rentable RID
	TCID        int64     // (I think this is deprecated)  if 0 then LM for whole acct, if > 0 then it's the amount for this payor; TCID
	Dt          time.Time // Balance is valid as of this time
	Balance     float64   // GLAccount balance at the end of the period
	State       int64     // 0 = Open, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
	LastModTime time.Time // auto updated
	LastModBy   int64     // user making the mod
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
}

// GLAccount describes the static (or mostly static) attributes of a Ledger
type GLAccount struct {
	Recid    int    `json:"recid"` // this is for the grid widget
	LID      int64  // unique id for this GLAccount
	PLID     int64  // unique id of Parent, 0 if no parent
	BID      int64  // Business unit associated with this GLAccount
	RAID     int64  // associated rental agreement, this field is only used when Type = 1
	TCID     int64  // associated payor, this field is only used when Type = 1
	GLNumber string // acct system name
	//Status      int64     // Whether a GL Account is currently unknown=0, inactive=1, active=2
	Name        string    // descriptive name for the GLAccount
	AcctType    string    // QB Acct Type: Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, Cost of Goods Sold, Other Income, Other Expense
	AllowPost   bool      // 0 = no posting, 1 = posting is allowed
	FLAGS       uint64    // 1<<0 = inactive:  0 = active account, 1 = inactive account
	Description string    // description for this account
	LastModTime time.Time // auto updated
	LastModBy   int64     // user making the mod
	CreateTS    time.Time // when was this record created
	CreateBy    int64     // employee UID (from phonebook) that created it
	// RARequired  int64     // 0 = during rental period, 1 = valid prior or during, 2 = valid during or after, 3 = valid before, during, and after
}

// Flow is a structure for to store temporarity flow data latest one
type Flow struct {
	BID         int64           // Business unit associated with this FlowPart
	FlowID      int64           // primary auto increment key
	UserRefNo   string          // user reference string
	FlowType    string          // RA="Rental Agreement Flow" etc...
	ID          int64           // id from permanent table, for FlowType "RA" it would be RAID
	Data        json.RawMessage // json data in mysql
	LastModTime time.Time       // last modified time
	LastModBy   int64           // last modified by whom
	CreateTS    time.Time       // created time
	CreateBy    int64           // created by whom
}

// RRprepSQL is a collection of prepared sql statements for the RentRoll db
type RRprepSQL struct {
	CountBusinessCustomAttributes           *sql.Stmt
	CountBusinessCustomAttrRefs             *sql.Stmt
	CountBusinessRentables                  *sql.Stmt
	CountBusinessRentableTypes              *sql.Stmt
	CountBusinessRentalAgreements           *sql.Stmt
	CountBusinessTransactants               *sql.Stmt
	DeleteAllRentalAgreementPets            *sql.Stmt
	DeleteAR                                *sql.Stmt
	DeleteAssessment                        *sql.Stmt
	DeleteCustomAttribute                   *sql.Stmt
	DeleteCustomAttributeRef                *sql.Stmt
	DeleteDemandSource                      *sql.Stmt
	DeleteDeposit                           *sql.Stmt
	DeleteDepositMethod                     *sql.Stmt
	DeleteDepository                        *sql.Stmt
	DeleteDepositPart                       *sql.Stmt
	DeleteInvoice                           *sql.Stmt
	DeleteInvoiceAssessments                *sql.Stmt
	DeleteInvoicePayors                     *sql.Stmt
	DeleteJournalAllocation                 *sql.Stmt
	DeleteJournalAllocations                *sql.Stmt
	DeleteJournal                           *sql.Stmt
	DeleteJournalMarker                     *sql.Stmt
	DeleteLedger                            *sql.Stmt
	DeleteLedgerEntry                       *sql.Stmt
	DeleteLedgerMarker                      *sql.Stmt
	DeleteNote                              *sql.Stmt
	DeleteNoteList                          *sql.Stmt
	DeleteNoteType                          *sql.Stmt
	DeletePaymentType                       *sql.Stmt
	DeletePayor                             *sql.Stmt
	DeleteProspect                          *sql.Stmt
	DeleteRatePlan                          *sql.Stmt
	DeleteRatePlanRef                       *sql.Stmt
	DeleteRatePlanRefRTRate                 *sql.Stmt
	DeleteRatePlanRefSPRate                 *sql.Stmt
	DeleteReceipt                           *sql.Stmt
	DeleteReceiptAllocation                 *sql.Stmt
	DeleteReceiptAllocations                *sql.Stmt
	DeleteRentableMarketRateInstance        *sql.Stmt
	DeleteRentableSpecialtyRef              *sql.Stmt
	DeleteRentableStatus                    *sql.Stmt
	DeleteRentableTypeRef                   *sql.Stmt
	DeleteRentableTypeRefWithRTID           *sql.Stmt
	DeleteRentableUser                      *sql.Stmt
	DeleteRentableUserByRBT                 *sql.Stmt
	DeleteRentalAgreement                   *sql.Stmt
	DeleteRentalAgreementPayor              *sql.Stmt
	DeleteAllRentalAgreementPayors          *sql.Stmt
	DeleteRentalAgreementPayorByRBT         *sql.Stmt
	DeleteRentalAgreementPet                *sql.Stmt
	DeleteRentalAgreementRentable           *sql.Stmt
	DeleteAllRentalAgreementRentables       *sql.Stmt
	DeleteRentalAgreementTax                *sql.Stmt
	DeleteSLString                          *sql.Stmt
	DeleteSLStrings                         *sql.Stmt
	DeleteStringList                        *sql.Stmt
	DeleteTransactant                       *sql.Stmt
	DeleteUser                              *sql.Stmt
	DeleteVehicle                           *sql.Stmt
	DeleteFlow                              *sql.Stmt
	FindAgreementByRentable                 *sql.Stmt
	FindTCIDByNote                          *sql.Stmt
	FindTransactantByPhoneOrEmail           *sql.Stmt
	GetAgreementsForRentable                *sql.Stmt
	GetAllARs                               *sql.Stmt
	GetAllAssessmentsByBusiness             *sql.Stmt
	GetAssessmentsByRAIDRange               *sql.Stmt
	GetAllBusinesses                        *sql.Stmt
	GetAllBusinessRentableTypes             *sql.Stmt
	GetAllBusinessRentableSpecialtyTypes    *sql.Stmt
	GetAllCustomAttributeRefs               *sql.Stmt
	GetAllCustomAttributes                  *sql.Stmt
	GetAllDemandSources                     *sql.Stmt
	GetAllDepositMethods                    *sql.Stmt
	GetAllDepositories                      *sql.Stmt
	GetAllDepositsInRange                   *sql.Stmt
	GetAllInvoicesInRange                   *sql.Stmt
	GetAllJournalsInRange                   *sql.Stmt
	GetAllLedgerEntriesForRAID              *sql.Stmt
	GetAllLedgerEntriesForRID               *sql.Stmt
	GetAllLedgerEntriesInRange              *sql.Stmt
	GetAllNotes                             *sql.Stmt
	GetAllNoteTypes                         *sql.Stmt
	GetAllRatePlanRefRTRates                *sql.Stmt
	GetAllRatePlanRefsInRange               *sql.Stmt
	GetAllRatePlanRefSPRates                *sql.Stmt
	GetAllRatePlans                         *sql.Stmt
	GetAllRentableAssessments               *sql.Stmt
	GetAllRentablesByBusiness               *sql.Stmt
	GetAllRentableSpecialtyRefs             *sql.Stmt
	GetAllRentalAgreementPets               *sql.Stmt
	GetPetsByTransactant                    *sql.Stmt
	GetAllRentalAgreements                  *sql.Stmt
	GetAllRentalAgreementsByRange           *sql.Stmt
	GetAllRentalAgreementTemplates          *sql.Stmt
	GetAllSingleInstanceAssessments         *sql.Stmt
	GetAllStringLists                       *sql.Stmt
	GetAllTransactants                      *sql.Stmt
	GetAllTransactantsForBID                *sql.Stmt
	GetAR                                   *sql.Stmt
	GetARByName                             *sql.Stmt
	GetARsByType                            *sql.Stmt
	GetARsByFLAGS                           *sql.Stmt
	GetAssessment                           *sql.Stmt
	GetAssessmentDuplicate                  *sql.Stmt
	GetAssessmentInstance                   *sql.Stmt
	GetAssessmentType                       *sql.Stmt
	GetAssessmentTypeByName                 *sql.Stmt
	GetBuilding                             *sql.Stmt
	GetBusiness                             *sql.Stmt
	GetBusinessByDesignation                *sql.Stmt
	GetCustomAttribute                      *sql.Stmt
	GetCustomAttributeByVals                *sql.Stmt
	GetCustomAttributeRef                   *sql.Stmt
	GetCustomAttributeRefs                  *sql.Stmt
	GetDemandSource                         *sql.Stmt
	GetDemandSourceByName                   *sql.Stmt
	GetDeposit                              *sql.Stmt
	GetDepositMethod                        *sql.Stmt
	GetDepositMethodByName                  *sql.Stmt
	GetDepository                           *sql.Stmt
	GetDepositoryByAccount                  *sql.Stmt
	GetDepositParts                         *sql.Stmt
	GetInvoice                              *sql.Stmt
	GetInvoiceAssessments                   *sql.Stmt
	GetInvoicePayors                        *sql.Stmt
	GetJournal                              *sql.Stmt
	GetJournalAllocation                    *sql.Stmt
	GetJournalAllocations                   *sql.Stmt
	GetJournalByRange                       *sql.Stmt
	GetJournalByReceiptID                   *sql.Stmt
	GetJournalMarker                        *sql.Stmt
	GetJournalMarkers                       *sql.Stmt
	GetJournalVacancy                       *sql.Stmt
	GetLatestLedgerMarkerByLID              *sql.Stmt
	GetLedger                               *sql.Stmt
	GetLedgerByGLNo                         *sql.Stmt
	GetLedgerByName                         *sql.Stmt
	GetLedgerByType                         *sql.Stmt
	GetLedgerEntriesForRAID                 *sql.Stmt
	GetLedgerEntriesForRentable             *sql.Stmt
	GetLedgerEntriesInRange                 *sql.Stmt
	GetLedgerEntriesInRangeByGLNo           *sql.Stmt
	GetLedgerEntriesInRangeByLID            *sql.Stmt
	GetLedgerEntry                          *sql.Stmt
	GetLedgerEntryByJAID                    *sql.Stmt
	GetLedgerList                           *sql.Stmt
	GetLedgerMarkerByDateRange              *sql.Stmt
	GetLedgerMarkerByLIDDateRange           *sql.Stmt
	GetLedgerMarkerOnOrBefore               *sql.Stmt
	GetTCLedgerMarkerOnOrBefore             *sql.Stmt
	GetTCLedgerMarkerOnOrAfter              *sql.Stmt
	GetLedgerMarkers                        *sql.Stmt
	GetNote                                 *sql.Stmt
	GetNoteAndChildNotes                    *sql.Stmt
	GetNoteList                             *sql.Stmt
	GetNoteListMembers                      *sql.Stmt
	GetNoteType                             *sql.Stmt
	GetPaymentType                          *sql.Stmt
	GetPaymentTypeByName                    *sql.Stmt
	GetPaymentTypesByBusiness               *sql.Stmt
	GetPayor                                *sql.Stmt
	GetPayorUnallocatedReceiptsCount        *sql.Stmt
	GetProspect                             *sql.Stmt
	GetRALedgerMarkerOnOrBefore             *sql.Stmt
	GetRALedgerMarkerOnOrBeforeDeprecated   *sql.Stmt
	GetRARentableForDate                    *sql.Stmt
	GetRatePlan                             *sql.Stmt
	GetRatePlanByName                       *sql.Stmt
	GetRatePlanRef                          *sql.Stmt
	GetRatePlanRefRTRate                    *sql.Stmt
	GetRatePlanRefsInRange                  *sql.Stmt
	GetRatePlanRefSPRate                    *sql.Stmt
	GetReceipt                              *sql.Stmt
	GetReceiptAllocation                    *sql.Stmt
	GetReceiptAllocations                   *sql.Stmt
	GetReceiptAllocationsByASMID            *sql.Stmt
	GetASMReceiptAllocationsInRAIDDateRange *sql.Stmt
	GetReceiptDuplicate                     *sql.Stmt
	GetReceiptsInDateRange                  *sql.Stmt
	GetRecurringAssessmentsByBusiness       *sql.Stmt
	GetRentable                             *sql.Stmt
	GetRentableByName                       *sql.Stmt
	GetRentableLedgerMarkerOnOrBefore       *sql.Stmt
	GetRentableMarketRates                  *sql.Stmt
	GetRentableMarketRateInstance           *sql.Stmt
	GetRentableSpecialtyRefs                *sql.Stmt
	GetRentableSpecialtyRefsByRange         *sql.Stmt
	GetRentableSpecialtyType                *sql.Stmt
	GetRentableSpecialtyTypeByName          *sql.Stmt
	GetRentableStatus                       *sql.Stmt
	GetRentableStatusByRange                *sql.Stmt
	GetRentableType                         *sql.Stmt
	GetRentableTypeByStyle                  *sql.Stmt
	GetRentableTypeDown                     *sql.Stmt
	GetRentableTypeRef                      *sql.Stmt
	GetRentableTypeRefsByRange              *sql.Stmt
	GetRentableUser                         *sql.Stmt
	GetRentableUserByRBT                    *sql.Stmt
	GetRentableUsersInRange                 *sql.Stmt
	GetRentalAgreement                      *sql.Stmt
	GetRentalAgreementByBusiness            *sql.Stmt
	GetRentalAgreementByRATemplateName      *sql.Stmt
	GetRentalAgreementPayor                 *sql.Stmt
	GetRentalAgreementPayorByRBT            *sql.Stmt
	GetRentalAgreementPayorsInRange         *sql.Stmt
	GetRentalAgreementPet                   *sql.Stmt
	GetRentalAgreementRentable              *sql.Stmt
	GetRentalAgreementRentables             *sql.Stmt
	GetRentalAgreementsByPayor              *sql.Stmt
	GetRentalAgreementsForRentable          *sql.Stmt
	GetRentalAgreementTax                   *sql.Stmt
	GetRentalAgreementTemplate              *sql.Stmt
	GetSecurityDepositAssessment            *sql.Stmt
	GetSLString                             *sql.Stmt
	GetSLStrings                            *sql.Stmt
	GetStringList                           *sql.Stmt
	GetStringListByName                     *sql.Stmt
	GetTransactant                          *sql.Stmt
	GetTransactantTypeDown                  *sql.Stmt
	GetUnallocatedReceipts                  *sql.Stmt
	GetUnallocatedReceiptsByPayor           *sql.Stmt
	GetUnitAssessments                      *sql.Stmt
	GetUnpaidAssessmentsByRAID              *sql.Stmt
	GetUser                                 *sql.Stmt
	GetVehicle                              *sql.Stmt
	GetVehiclesByBID                        *sql.Stmt
	GetVehiclesByLicensePlate               *sql.Stmt
	GetVehiclesByTransactant                *sql.Stmt
	GetFlow                                 *sql.Stmt // flow table
	GetFlowsByFlowType                      *sql.Stmt // flow table
	GetFlowIDsByUser                        *sql.Stmt // flow table
	InsertAR                                *sql.Stmt
	InsertAssessment                        *sql.Stmt
	InsertAssessmentType                    *sql.Stmt
	InsertBuilding                          *sql.Stmt
	InsertBuildingWithID                    *sql.Stmt
	InsertBusiness                          *sql.Stmt
	InsertCustomAttribute                   *sql.Stmt
	InsertCustomAttributeRef                *sql.Stmt
	InsertDemandSource                      *sql.Stmt
	InsertDeposit                           *sql.Stmt
	InsertDepositMethod                     *sql.Stmt
	InsertDepository                        *sql.Stmt
	InsertDepositPart                       *sql.Stmt
	InsertInvoice                           *sql.Stmt
	InsertInvoiceAssessment                 *sql.Stmt
	InsertInvoicePayor                      *sql.Stmt
	InsertJournal                           *sql.Stmt
	InsertJournalAllocation                 *sql.Stmt
	InsertJournalMarker                     *sql.Stmt
	InsertLedger                            *sql.Stmt
	InsertLedgerAllocation                  *sql.Stmt
	InsertLedgerEntry                       *sql.Stmt
	InsertLedgerMarker                      *sql.Stmt
	InsertNote                              *sql.Stmt
	InsertNoteList                          *sql.Stmt
	InsertNoteType                          *sql.Stmt
	InsertPaymentType                       *sql.Stmt
	InsertPayor                             *sql.Stmt
	InsertProspect                          *sql.Stmt
	InsertRatePlan                          *sql.Stmt
	InsertRatePlanRef                       *sql.Stmt
	InsertRatePlanRefRTRate                 *sql.Stmt
	InsertRatePlanRefSPRate                 *sql.Stmt
	InsertReceipt                           *sql.Stmt
	InsertReceiptAllocation                 *sql.Stmt
	InsertRentable                          *sql.Stmt
	InsertRentableMarketRates               *sql.Stmt
	InsertRentableSpecialtyRef              *sql.Stmt
	InsertRentableSpecialtyType             *sql.Stmt
	InsertRentableStatus                    *sql.Stmt
	InsertRentableType                      *sql.Stmt
	InsertRentableTypeRef                   *sql.Stmt
	InsertRentableUser                      *sql.Stmt
	InsertRentalAgreement                   *sql.Stmt
	InsertRentalAgreementPayor              *sql.Stmt
	InsertRentalAgreementPet                *sql.Stmt
	InsertRentalAgreementRentable           *sql.Stmt
	InsertRentalAgreementTax                *sql.Stmt
	InsertRentalAgreementTemplate           *sql.Stmt
	InsertSLString                          *sql.Stmt
	InsertStringList                        *sql.Stmt
	InsertTransactant                       *sql.Stmt
	InsertUser                              *sql.Stmt
	InsertVehicle                           *sql.Stmt
	InsertFlow                              *sql.Stmt // flow table
	ReadRatePlan                            *sql.Stmt
	ReadRatePlanRef                         *sql.Stmt
	UIRAGrid                                *sql.Stmt
	UpdateAR                                *sql.Stmt
	UpdateAssessment                        *sql.Stmt
	UpdateBusiness                          *sql.Stmt
	UpdateCustomAttribute                   *sql.Stmt
	UpdateDemandSource                      *sql.Stmt
	UpdateDeposit                           *sql.Stmt
	UpdateDepositMethod                     *sql.Stmt
	UpdateDepository                        *sql.Stmt
	UpdateInvoice                           *sql.Stmt
	UpdateJournalAllocation                 *sql.Stmt
	UpdateLedger                            *sql.Stmt
	UpdateLedgerMarker                      *sql.Stmt
	UpdateNote                              *sql.Stmt
	UpdateNoteType                          *sql.Stmt
	UpdatePaymentType                       *sql.Stmt
	UpdatePayor                             *sql.Stmt
	UpdateProspect                          *sql.Stmt
	UpdateRatePlan                          *sql.Stmt
	UpdateRatePlanRef                       *sql.Stmt
	UpdateRatePlanRefRTRate                 *sql.Stmt
	UpdateRatePlanRefSPRate                 *sql.Stmt
	UpdateReceipt                           *sql.Stmt
	UpdateReceiptAllocation                 *sql.Stmt
	UpdateRentable                          *sql.Stmt
	UpdateRentableMarketRateInstance        *sql.Stmt
	UpdateRentableSpecialtyRef              *sql.Stmt
	UpdateRentableStatus                    *sql.Stmt
	UpdateRentableType                      *sql.Stmt
	UpdateRentableTypeToActive              *sql.Stmt
	UpdateRentableTypeToInactive            *sql.Stmt
	UpdateRentableTypeRef                   *sql.Stmt
	UpdateRentableUser                      *sql.Stmt
	UpdateRentableUserByRBT                 *sql.Stmt
	UpdateRentalAgreement                   *sql.Stmt
	UpdateRentalAgreementPayor              *sql.Stmt
	UpdateRentalAgreementPayorByRBT         *sql.Stmt
	UpdateRentalAgreementPet                *sql.Stmt
	UpdateRentalAgreementRentable           *sql.Stmt
	UpdateRentalAgreementTax                *sql.Stmt
	UpdateSLString                          *sql.Stmt
	UpdateStringList                        *sql.Stmt
	UpdateTransactant                       *sql.Stmt
	UpdateUser                              *sql.Stmt
	UpdateVehicle                           *sql.Stmt
	UpdateFlowData                          *sql.Stmt // flow table
	GetAssessmentInstancesByParent          *sql.Stmt
	GetJournalAllocationsByASMID            *sql.Stmt
	GetRentableTypeRefs                     *sql.Stmt
	GetAllRentableStatus                    *sql.Stmt
	GetRentalAgreementTypeDown              *sql.Stmt
	GetLedgerEntriesByJAID                  *sql.Stmt
	GetLedgersForGrid                       *sql.Stmt
	GetAssessmentFirstInstance              *sql.Stmt
	GetDepositoryByName                     *sql.Stmt
	GetDepositoryByLID                      *sql.Stmt
	UpdateDepositPart                       *sql.Stmt
	CountLedgerEntries                      *sql.Stmt
	GetExpense                              *sql.Stmt
	InsertExpense                           *sql.Stmt
	DeleteExpense                           *sql.Stmt
	UpdateExpense                           *sql.Stmt
	GetRentableTypeByName                   *sql.Stmt
	GetRALedgerMarkerOnOrAfter              *sql.Stmt
	GetReceiptAllocationsThroughDate        *sql.Stmt
	GetInitialLedgerMarkerByRAID            *sql.Stmt
	GetInitialLedgerMarkerByRID             *sql.Stmt
	GetRARentableLedgerMarkerOnOrBefore     *sql.Stmt
	GetAssessmentsByRARRange                *sql.Stmt
	GetASMReceiptAllocationsInRARDateRange  *sql.Stmt
	GetRentableStatusOnOrAfter              *sql.Stmt
	InsertSubAR                             *sql.Stmt
	GetSubAR                                *sql.Stmt
	GetSubARs                               *sql.Stmt
	UpdateSubAR                             *sql.Stmt
	DeleteSubAR                             *sql.Stmt
	DeleteSubARs                            *sql.Stmt
	GetJournalAllocationsByASMandRCPTID     *sql.Stmt
	GetJournalByTypeAndID                   *sql.Stmt
	GetTask                                 *sql.Stmt
	GetTasks                                *sql.Stmt
	GetTaskList                             *sql.Stmt
	GetTaskDescriptor                       *sql.Stmt
	GetTaskListDefinition                   *sql.Stmt
	GetAllTaskListDefinitions               *sql.Stmt
	InsertTask                              *sql.Stmt
	InsertTaskList                          *sql.Stmt
	InsertTaskDescriptor                    *sql.Stmt
	InsertTaskListDefinition                *sql.Stmt
	UpdateTask                              *sql.Stmt
	UpdateTaskList                          *sql.Stmt
	UpdateTaskDescriptor                    *sql.Stmt
	UpdateTaskListDefinition                *sql.Stmt
	DeleteTask                              *sql.Stmt
	DeleteTaskList                          *sql.Stmt
	DeleteTaskDescriptor                    *sql.Stmt
	DeleteTaskListDefinition                *sql.Stmt
	GetEpochAssessmentsByRentalAgreement    *sql.Stmt
	GetAllRentalAgreementRentables          *sql.Stmt
	GetTaskListDescriptors                  *sql.Stmt
	DeleteTaskListTasks                     *sql.Stmt
	GetTaskListDefinitionByName             *sql.Stmt
	GetDueTaskLists                         *sql.Stmt
	CheckForTLDInstances                    *sql.Stmt
	GetAllParentTaskLists                   *sql.Stmt
	GetTaskListInstanceInRange              *sql.Stmt
	GetLatestCompletedTaskList              *sql.Stmt
	GetClosePeriod                          *sql.Stmt
	GetLastClosePeriod                      *sql.Stmt
	InsertClosePeriod                       *sql.Stmt
	UpdateClosePeriod                       *sql.Stmt
	DeleteClosePeriod                       *sql.Stmt
	GetFlowMetaDataInRange                  *sql.Stmt
	GetFlowForRAID                          *sql.Stmt
	GetBusinessProperties                   *sql.Stmt
	GetBusinessPropertiesByName             *sql.Stmt
	InsertBusinessProperties                *sql.Stmt
	UpdateBusinessPropertiesData            *sql.Stmt
	DeleteBusinessProperties                *sql.Stmt
	GetAssessmentsByRAIDRID                 *sql.Stmt
}

// DeleteBusinessFromDB deletes information from all tables if it is part of the supplied BID.
// Use this call with extreme caution. There's no recovery.
func DeleteBusinessFromDB(ctx context.Context, BID int64) (int64, error) {
	// Might want to check context values here? like session, transaction?

	noRecs := int64(0)
	for k := range RRdb.DBFields {
		s := fmt.Sprintf("DELETE FROM %s WHERE BID=%d", k, BID)
		result, err := RRdb.Dbrr.Exec(s)
		if err != nil {
			Ulog("DeleteBusinessFromDB: error executing %q   -- err = %s\n", s, err.Error())
		} else {
			x, _ := result.RowsAffected()
			noRecs += x
		}
	}
	RRdb.BUDlist, RRdb.BizCache = BuildBusinessDesignationMap()
	return noRecs, nil
}

// BusinessTypeLists is a struct holding a collection of Types associated with a business
type BusinessTypeLists struct {
	BID          int64
	PmtTypes     map[int64]*PaymentType // payment types accepted
	DefaultAccts map[int64]*GLAccount   // index by the predifined contants DFAC*, value = GL No of that account
	GLAccounts   map[int64]GLAccount    // all the accounts for this business
	AR           map[int64]AR           // Account Rules
	NoteTypes    []NoteType             // all defined note types for this business
}

// BusinessCache will be used for caching some common info
// which are frequently accessed throughout the application
// so it saves the db hit calls.
type BusinessCache struct {
	BID   int64
	BUD   string
	FLAGS int64
}

// RRdb is a struct with all variables needed by the db infrastructure
var RRdb struct {
	Prepstmt RRprepSQL
	PBsql    PBprepSQL
	Dbdir    *sql.DB                      // phonebook db
	Dbrr     *sql.DB                      //rentroll db
	BizTypes map[int64]*BusinessTypeLists // details about a business
	BizCache map[int64]BusinessCache      // map of BID to business cache struct
	BUDlist  Str2Int64Map                 // list of known business Designations
	DBFields map[string]string            // map of db table fields DBFields[tablename] = field list
	Zone     *time.Location               // what timezone should the server use?
	Key      []byte                       // crypto key
	noAuth   bool                         // if enable that means auth is not required, (should be moved in some common app struct!)
	Rand     *rand.Rand                   // for generating Reference Numbers or other UniqueIDs
	// TODO(sudip): NoAuth will be moved to something internal pkg app struct
}

// SetAuthFlag enable/disable authentication in RRdb
func SetAuthFlag(noauth bool) {
	if AppConfig.Env != extres.APPENVPROD { // NOT only applicable for PROD Environment
		RRdb.noAuth = noauth
	}
}

// BuildBusinessDesignationMap builds a map of biz designations to BIDs
func BuildBusinessDesignationMap() (map[string]int64, map[int64]BusinessCache) {
	// Console("Entered BuildBusinessDesignationMap\n")
	var sl = map[string]int64{}
	var bc = make(map[int64]BusinessCache)

	bl, err := GetAllBiz()
	if err != nil {
		Ulog("GetAllBusinesses: err = %s\n", err.Error())
	}
	for i := 0; i < len(bl); i++ {
		sl[bl[i].Designation] = bl[i].BID
		bc[bl[i].BID] = BusinessCache{BID: bl[i].BID, BUD: bl[i].Designation, FLAGS: bl[i].FLAGS}
		// Console("bizlist[%s]=%d\n", bl[i].Designation, bl[i].BID)
	}
	return sl, bc
}

// InitDBHelpers initializes the db infrastructure
func InitDBHelpers(dbrr, dbdir *sql.DB) {
	RRdb.Dbdir = dbdir
	RRdb.Dbrr = dbrr
	RRdb.BizTypes = make(map[int64]*BusinessTypeLists)
	RRdb.DBFields = map[string]string{}
	buildPreparedStatements()
	buildPBPreparedStatements()
	InitCaches()

	RRdb.BUDlist, RRdb.BizCache = BuildBusinessDesignationMap()

	for i := 0; i < len(QBAcctInfo); i++ {
		QBAcctType = append(QBAcctType, QBAcctInfo[i].Name)
	}

	now := time.Now()
	RRdb.Rand = rand.New(rand.NewSource(now.UnixNano()))
}

// InitBusinessFields initialize the lists in rlib's internal data structures
func InitBusinessFields(bid int64) {
	if nil == RRdb.BizTypes[bid] {
		bt := BusinessTypeLists{
			BID:          bid,
			PmtTypes:     make(map[int64]*PaymentType),
			DefaultAccts: make(map[int64]*GLAccount),
			GLAccounts:   make(map[int64]GLAccount),
			AR:           make(map[int64]AR),
		}
		RRdb.BizTypes[bid] = &bt
	}
}

// InitBizInternals initializes several internal structures with information about the business.
func InitBizInternals(bid int64, xbiz *XBusiness) error {
	var (
		err error
	)

	// fmt.Printf("Entered InitBizInternals\n")
	err = GetXBiz(bid, xbiz) // get its info
	if err != nil {
		return err
	}

	InitBusinessFields(bid)

	// GetDefaultLedgers(bid) // Gather its chart of accounts
	RRdb.BizTypes[bid].GLAccounts, err = getLedgerMap(bid)
	if err != nil {
		return err
	}

	RRdb.BizTypes[bid].AR, err = getARMap(bid)
	if err != nil {
		return err
	}

	// TODO(Steve): why we're ignoring note types here? shouldnt we store somewhere?
	_, err = getBusinessAllNoteTypes(bid)
	if err != nil {
		return err
	}

	err = loadRentableTypeCustomaAttributes(xbiz)
	if err != nil {
		return err
	}

	return err
}
