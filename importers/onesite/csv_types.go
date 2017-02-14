package onesite

import (
	"reflect"
	"rentroll/importers/core"
	"strings"
)

// CSVFieldMap is struct which contains several categories
// used to store the data from onesite to rentroll system
type CSVFieldMap struct {
	RentableTypeCSV    core.RentableTypeCSV
	PeopleCSV          core.PeopleCSV
	RentableCSV        core.RentableCSV
	RentalAgreementCSV core.RentalAgreementCSV
	CustomAttributeCSV core.CustomAttributeCSV
}

// csvColumnFieldMap contains internal OneSite Structure fields
// to csv columns, used to refer columns from struct fields
var csvColumnFieldMap = map[string]string{
	"unit":            "Unit",
	"floorplan":       "FloorPlan",
	"unitdesignation": "UnitDesignation",
	"sqft":            "SQFT",
	"unitleasestatus": "UnitLeaseStatus",
	"name":            "Name",
	"phonenumber":     "PhoneNumber",
	"email":           "Email",
	"movein":          "MoveIn",
	"moveout":         "MoveOut",
	"leasestart":      "LeaseStart",
	"leaseend":        "LeaseEnd",
	"marketaddl":      "MarketAddl",
	"rent":            "Rent",
	// "tax":              "TAX",
}

var marketAddl = "marketaddl"
var marketRent = "marketrent"

// CSVRow contains fields which represents value
// exactly to the each raw of onesite input csv file
type CSVRow struct {
	Unit            string
	FloorPlan       string
	UnitDesignation string
	SQFT            string
	UnitLeaseStatus string
	Name            string
	PhoneNumber     string
	Email           string
	MoveIn          string
	MoveOut         string
	LeaseStart      string
	LeaseEnd        string
	MarketAddl      string
	Rent            string
	// Tax             string
}

// getCSVHeadersIndexMap returns the map of fields with
// undetermined indexes
func getCSVHeadersIndexMap() map[string]int {

	// csvHeadersIndex holds the map of headers with its index
	csvHeadersIndex := map[string]int{
		"Unit":            -1,
		"FloorPlan":       -1,
		"UnitDesignation": -1,
		"SQFT":            -1,
		"UnitLeaseStatus": -1,
		"Name":            -1,
		"PhoneNumber":     -1,
		"Email":           -1,
		"MoveIn":          -1,
		"MoveOut":         -1,
		"LeaseStart":      -1,
		"LeaseEnd":        -1,
		"MarketAddl":      -1,
		"Rent":            -1,
		// "Tax":             -1,
	}

	return csvHeadersIndex
}

// loadOneSiteCSVRow used to load data from slice
// into CSVRow struct and return that struct
func loadOneSiteCSVRow(csvHeadersIndex map[string]int, data []string) (bool, CSVRow) {
	csvRow := reflect.New(reflect.TypeOf(CSVRow{}))
	rowLoaded := false

	for header, index := range csvHeadersIndex {
		value := strings.TrimSpace(data[index])
		csvRow.Elem().FieldByName(header).Set(reflect.ValueOf(value))
	}

	// if blank data has not been passed then only need to return true
	if (CSVRow{}) != csvRow.Elem().Interface().(CSVRow) {
		rowLoaded = true
	}

	return rowLoaded, csvRow.Elem().Interface().(CSVRow)
}
