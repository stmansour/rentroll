package onesite

import (
	"rentroll/importers/core"
	"rentroll/rcsv"
	"strings"
)

// TempCSVStoreName holds the name of csvstore folder
var TempCSVStoreName = "temp_CSVs"

// TempCSVStore is used to store temporary csv files
var TempCSVStore string

// FieldDefaultValues isused to overwrite if user has not passed to values for these fields
var FieldDefaultValues = map[string]string{
	"ManageToBudget": "1", // always take to default this one
	"RentCycle":      "6", // maybe overridden by user supplied value
	"Proration":      "4", // maybe overridden by user supplied value
	"GSRPC":          "4", // maybe overridden by user supplied value
	"AssignmentTime": "1", // always take to default this one
	"Renewal":        "2", // always take to default this one
}

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
	MoveOut         = iota
	LeaseStart      = iota
	LeaseEnd        = iota
	MarketAddl      = iota
	Rent            = iota
	// Tax             = iota
)

// CSVLoadHandler struct is for routines that want to table-ize their loading.
type csvLoadHandler struct {
	Fname        string
	Handler      func(string) []error
	TraceDataMap string
	DBType       int
}

// canWriteCSVStatusMap holds the set of csv types with key of status value
// used in checking if csv file for db type is able to perform write operation
// for the given status value
var canWriteCSVStatusMap = map[string][]int{
	// if rentable status is blank then still you can write data to these CSVs
	"": {
		core.RENTABLETYPECSV,
		core.RENTALAGREEMENTCSV,
		// core.PEOPLECSV,
		core.CUSTOMATTRIUTESCSV,
	},
	"occupied": {
		core.RENTABLETYPECSV,
		core.PEOPLECSV,
		core.RENTABLECSV,
		core.RENTALAGREEMENTCSV,
		core.CUSTOMATTRIUTESCSV,
	},
	"model": {
		core.RENTABLETYPECSV,
		core.RENTABLECSV,
		core.CUSTOMATTRIUTESCSV,
	},
	"vacant": {
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

var dupTransactantWithPrimaryEmail = "PrimaryEmail"

// var dupTransactantWithCellPhone = "CellPhone"

// will be used exact before rowIndex to format Notes in people csv "onesite:<rowIndex>"
const (
	onesiteNotesPrefix = "onesite$"
	tcidPrefix         = "TC000"
)

var specialCharsReplacer = strings.NewReplacer(
	"`", "", "~", "", "!", "", "@", "", "#", "", "$", "", "%", "", "^", "", "&", "", "*", "", "(", "", ")", "", "-", "", "_", "", "+", "", "=", "", //line1
	"{", "", "[", "", "}", "", "]", "", "|", "", "\\", "", //line2
	";", "", ":", "", "\"", "", "'", "", // line3
	",", "", "<", "", ".", "", ">", "", "/", "", "?", "", // line4
	" ", "", // whitespace
)
