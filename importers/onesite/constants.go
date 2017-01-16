package onesite

import (
	"rentroll/importers/core"
	"rentroll/rcsv"
)

// tempCSVStoreName holds the name of csvstore folder
var tempCSVStoreName = "temp_CSVs"

// used to store temporary csv files
var tempCSVStore string

// CARD Custom Attriute Ref Data struct, holds data
// from which we'll insert customAttributeRef in system
type CARD struct {
	BID      int64
	RTID     string
	Style    string
	SqFt     int64
	CID      string
	RowIndex int
}

// prefixCSVFile is a map which holds the prefix of csv files
// so that temporarily program can create csv files with this
var prefixCSVFile = map[string]string{
	"rentable_types":   "rentableTypes_",
	"people":           "people_",
	"rental_agreement": "rentalAgreement_",
	"rentable":         "rentable_",
	"custom_attribute": "customAttribute_",
}

// RRRentableStatus is status for rentable in rentroll system
var RRRentableStatus = map[string]string{
	"unknown":        "0",
	"online":         "1",
	"admin":          "2",
	"employee":       "3",
	"owner occupied": "4",
	"offline":        "5",
}

// RentableStatusCSV is mapping for rentable status between onesite and rentroll
var RentableStatusCSV = map[string]string{
	"vacant":   "online",
	"occupied": "online",
	"model":    "admin",
}

// define column fields with order
const (
	Unit            = iota
	FloorPlan       = iota
	UnitDesignation = iota
	Sqft            = iota
	UnitLeaseStatus = iota
	Name            = iota
	PhoneNumber     = iota
	Email           = iota
	MoveIn          = iota
	NoticeForDate   = iota
	MoveOut         = iota
	LeaseStart      = iota
	LeaseEnd        = iota
	MarketAddl      = iota
	DepOnHand       = iota
	Balance         = iota
	TotalCharges    = iota
	Rent            = iota
	WaterReImb      = iota
	Corp            = iota
	Discount        = iota
	Platinum        = iota
	Tax             = iota
	ElectricReImb   = iota
	Fire            = iota
	ConcSpecl       = iota
	WashDry         = iota
	EmplCred        = iota
	Short           = iota
	PetFee          = iota
	TrashReImb      = iota
	TermFee         = iota
	LakeView        = iota
	Utility         = iota
	Furn            = iota
	Mtom            = iota
	Referral        = iota
)

// defined csv columns
var csvCols = []rcsv.CSVColumn{
	{Name: "Unit", Index: Unit},
	{Name: "FloorPlan", Index: FloorPlan},
	{Name: "UnitDesignation", Index: UnitDesignation},
	{Name: "SQFT", Index: Sqft},
	{Name: "Unit/LeaseStatus", Index: UnitLeaseStatus},
	{Name: "Name", Index: Name},
	{Name: "PhoneNumber", Index: PhoneNumber},
	{Name: "Email", Index: Email},
	{Name: "Move-In", Index: MoveIn},
	{Name: "NoticeForDate", Index: NoticeForDate},
	{Name: "Move-Out", Index: MoveOut},
	{Name: "LeaseStart", Index: LeaseStart},
	{Name: "LeaseEnd", Index: LeaseEnd},
	{Name: "Market+Addl.", Index: MarketAddl},
	{Name: "DepOnHand", Index: DepOnHand},
	{Name: "Balance", Index: Balance},
	{Name: "TotalCharges", Index: TotalCharges},
	{Name: "RENT", Index: Rent},
	{Name: "WATERREIMB", Index: WaterReImb},
	{Name: "CORP", Index: Corp},
	{Name: "DISCOUNT", Index: Discount},
	{Name: "Platinum", Index: Platinum},
	{Name: "TAX", Index: Tax},
	{Name: "ELECTRICREIMB", Index: ElectricReImb},
	{Name: "Fire", Index: Fire},
	{Name: "CONC/SPECL", Index: ConcSpecl},
	{Name: "WASH/DRY", Index: WashDry},
	{Name: "EMPLCRED", Index: EmplCred},
	{Name: "SHORT", Index: Short},
	{Name: "PETFEE", Index: PetFee},
	{Name: "TRASHREIMB", Index: TrashReImb},
	{Name: "TERMFEE", Index: TermFee},
	{Name: "Lakeview", Index: LakeView},
	{Name: "UTILITY", Index: Utility},
	{Name: "FURN", Index: Furn},
	{Name: "MTOM", Index: Mtom},
	{Name: "REFERRAL", Index: Referral},
}

// // this structure used as table-driven approach
// // to write data in csv
// type oneSiteCSVWriter struct {
// 	handler                 func(string)
// 	csvTypeNo               int
// 	recordCount             *int
// 	rowIndex                int
// 	traceDataMap            map[int]int
// 	csvWriter               *csv.Writer
// 	csvRow                  *CSVRow
// 	avoidData               interface{}
// 	currentTime             time.Time
// 	currentTimeFormat       string
// 	userRRValues            map[string]string
// 	dbTypeCSV               interface{}
// 	customAttributesRefData map[string]CARD
// 	business                *rlib.Business
// }

// CSVLoadHandler struct is for routines that want to table-ize their loading.
type csvLoadHandler struct {
	Fname        string
	Handler      func(string) []error
	TraceDataMap string
}

// canWriteCSVStatusMap holds the set of csv types with key of status value
// used in checking if csv file for db type is able to perform write operation
// for the given status value
var canWriteCSVStatusMap = map[string][]int{
	"occupied": []int{
		core.RENTABLETYPECSV,
		core.PEOPLECSV,
		core.RENTABLECSV,
		core.RENTALAGREEMENTCSV,
		core.CUSTOMATTRIUTESCSV,
	},
	"model": []int{
		core.RENTABLETYPECSV,
		core.RENTABLECSV,
		core.CUSTOMATTRIUTESCSV,
	},
	"vacant": []int{
		core.RENTABLETYPECSV,
		core.RENTABLECSV,
		core.CUSTOMATTRIUTESCSV,
	},
}

// this slice contains list of strings which should be discarded
// used in csvRecordsToSkip function
var csvRecordsSkipList = []string{
	rcsv.DupTransactant,
	rcsv.DupRentableType,
	rcsv.DupCustomAttribute,
	rcsv.DupRentable,
	rcsv.RentableAlreadyRented,
}
