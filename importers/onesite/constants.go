package onesite

import (
	"fmt"
	"rentroll/importers/core"
	"rentroll/rcsv"
	"rentroll/rlib"
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

// RentableUseTypeCSV is mapping for rentable status between onesite and rentroll
var RentableUseTypeCSV = map[string]string{
	"vacant":   fmt.Sprintf("%d", rlib.USETYPEstandard),
	"occupied": fmt.Sprintf("%d", rlib.USETYPEstandard),
	"model":    fmt.Sprintf("%d", rlib.USETYPEadministrative),
}

// CSVLoadHandler struct is for routines that want to table-ize their loading.
type csvLoadHandler struct {
	Fname        string
	Handler      rcsv.CSVLoadHandlerFunc
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
		core.RENTABLECSV,
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
