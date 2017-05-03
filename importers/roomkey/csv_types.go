package roomkey

import (
	"reflect"
	"rentroll/importers/core"
	"strings"
)

// CSVFieldMap is struct which contains several categories
// used to store the data from roomkey to rentroll system
type CSVFieldMap struct {
	RentableTypeCSV    core.RentableTypeCSV
	PeopleCSV          core.PeopleCSV
	RentableCSV        core.RentableCSV
	RentalAgreementCSV core.RentalAgreementCSV
}

// csvColumnFieldMap contains internal Roomkey Structure fields
// to csv columns, used to refer columns from struct fields
var csvColumnFieldMap = map[string]string{
	"guest":              "Guest",
	"res":                "Res",
	"dateres":            "DateRes",
	"datein":             "DateIn",
	"dateout":            "DateOut",
	"adult":              "Adult",
	"child":              "Child",
	"room":               "Room",
	"roomtype":           "RoomType",
	"rate":               "Rate",
	"ratename":           "RateName",
	"groupcorporatename": "GroupCorporate",
}

// CSVRow contains fields which represents value
// exactly to the each raw of roomkey input csv file
type CSVRow struct {
	Guest          string
	Description    string
	Res            string
	DateRes        string
	DateIn         string
	DateOut        string
	Adult          string
	Child          string
	Room           string
	RoomType       string
	Rate           string
	RateName       string
	GroupCorporate string
}

// getCSVHeadersIndexMap returns the map of fields with
// undetermined indexes for roomkey csv
func getCSVHeadersIndexMap() map[string]int {

	// csvHeadersIndex holds the map of headers with its index
	csvHeadersIndex := map[string]int{
		"Guest":          -1,
		"Res":            -1,
		"DateRes":        -1,
		"DateIn":         -1,
		"DateOut":        -1,
		"Adult":          -1,
		"Child":          -1,
		"Room":           -1,
		"RoomType":       -1,
		"Rate":           -1,
		"RateName":       -1,
		"GroupCorporate": -1,
	}

	return csvHeadersIndex
}

// tells which type of row is
var csvRowType = map[string]int{
	"page":        0,
	"header":      1,
	"record":      2,
	"description": 3,
}

// by which index, it will decide tyep of row
var rowTypeDetectionCSVIndex = map[string]int{
	"page":        0,
	"description": 2,
}

// loadRoomKeyCSVRow used to load data from slice
// into CSVRow struct and return that struct
func loadRoomKeyCSVRow(csvHeadersIndex map[string]int, data []string) (bool, CSVRow) {
	csvRow := reflect.New(reflect.TypeOf(CSVRow{}))
	skipRow := false

	// else go for records
	for header, index := range csvHeadersIndex {
		if index < len(data) {
			value := strings.TrimSpace(data[index])
			csvRow.Elem().FieldByName(header).Set(reflect.ValueOf(value))
		}
	}

	// if blank data has not been passed then only need to return true
	if (CSVRow{}) == csvRow.Elem().Interface().(CSVRow) {
		skipRow = true
	}

	return skipRow, csvRow.Elem().Interface().(CSVRow)
}

// check that row is headerline
func isRoomKeyHeaderLine(rowHeaders []string) (bool, map[string]int) {
	csvHeadersIndex := getCSVHeadersIndexMap()

	for colIndex := 0; colIndex < len(rowHeaders); colIndex++ {
		// remove all white spaces and make lower case
		cellTextValue := strings.ToLower(
			core.SpecialCharsReplacer.Replace(rowHeaders[colIndex]))

		// if header is exist in map then overwrite it position
		if field, ok := csvColumnFieldMap[cellTextValue]; ok {
			// ******** VERY SPECIAL CASE ***************
			// dateIn data appears in next column, so need to add 1
			if cellTextValue == "datein" {
				csvHeadersIndex[field] = colIndex + 1
			} else {
				// normal case
				csvHeadersIndex[field] = colIndex
			}
		}
	}
	// check after row columns parsing that headers are found or not
	headersFound := true
	for _, v := range csvHeadersIndex {
		if v == -1 {
			headersFound = false
			break
		}
	}

	return headersFound, csvHeadersIndex
}

// isRoomKeyPageRow check row is used for new page records
func isRoomKeyPageRow(data []string) bool {
	// if first column is not empty then it is
	return strings.TrimSpace(data[rowTypeDetectionCSVIndex["page"]]) != ""
}

func isRoomKeyDescriptionRow(data []string) bool {
	// if third column is not empty then it is
	return strings.TrimSpace(data[rowTypeDetectionCSVIndex["description"]]) != ""
}

// guestCSVColumnFieldMap contains internal Roomkey Guest Structure fields
// to guest csv columns, used to refer columns from struct fields
var guestCSVColumnFieldMap = map[string]string{
	"guestname":     "GuestName",
	"firstname":     "FirstName",
	"lastname":      "LastName",
	"email":         "Email",
	"mainphone":     "MainPhone",
	"address":       "Address",
	"address2":      "Address2",
	"city":          "City",
	"stateprovince": "StateProvince",
	"zippostalcode": "ZipPostalCode",
	"country":       "Country",
}

// GuestCSVRow contains fields which represents value
// exactly to the each row of roomkey input csv file
type GuestCSVRow struct {
	GuestName     string
	FirstName     string
	LastName      string
	Email         string
	MainPhone     string
	Address       string
	Address2      string
	City          string
	StateProvince string
	ZipPostalCode string
	Country       string
}

// getGuestCSVHeadersIndexMap returns the map of fields with
// undetermined indexes for guest csv
func getGuestCSVHeadersIndexMap() map[string]int {

	// csvHeadersIndex holds the map of headers with its index
	csvHeadersIndex := map[string]int{
		"GuestName":     -1,
		"FirstName":     -1,
		"LastName":      -1,
		"Email":         -1,
		"MainPhone":     -1,
		"Address":       -1,
		"Address2":      -1,
		"City":          -1,
		"StateProvince": -1,
		"ZipPostalCode": -1,
		"Country":       -1,
	}

	return csvHeadersIndex
}

// loadRoomKeyCSVRow used to load data from slice
// into CSVRow struct and return that struct
func loadGuestInfoCSVRow(csvHeadersIndex map[string]int, data []string) (bool, GuestCSVRow) {
	csvRow := reflect.New(reflect.TypeOf(GuestCSVRow{}))
	rowLoaded := false

	for header, index := range csvHeadersIndex {
		value := strings.TrimSpace(data[index])
		csvRow.Elem().FieldByName(header).Set(reflect.ValueOf(value))
	}

	// if blank data has not been passed then only need to return true
	if (GuestCSVRow{}) != csvRow.Elem().Interface().(GuestCSVRow) {
		rowLoaded = true
	}

	// if no valid email then fill with blank value
	if !core.IsValidEmail(csvRow.Elem().FieldByName("Email").Interface().(string)) {
		csvRow.Elem().FieldByName("Email").Set(reflect.ValueOf(""))
	}
	return rowLoaded, csvRow.Elem().Interface().(GuestCSVRow)
}
