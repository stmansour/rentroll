package roomkey

import (
	"fmt"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rcsv"
	"rentroll/rlib"
	"strings"
)

// CSVFieldMap is struct which contains several categories
// used to store the data from roomkey to rentroll system
type CSVFieldMap struct {
	RentableTypeCSV    core.RentableTypeCSV
	PeopleCSV          core.PeopleCSV
	RentableCSV        core.RentableCSV
	RentalAgreementCSV core.RentalAgreementCSV
	CustomAttributeCSV core.CustomAttributeCSV
}

// CSVRow contains fields which represents value
// exactly to the each raw of roomkey input csv file
type CSVRow struct {
	Empty1         string
	Guest          string
	Description    string
	ResID          string
	DateRes        string
	DateIn         string
	Empty3         string
	DateOut        string
	Adults         string
	Child          string
	Room           string
	RoomType       string
	Rate           string
	RateName       string
	GroupCorporate string
}

// csvRowFieldRules is map contains rules for specific fields in roomkey
var csvRowFieldRules = map[string]map[string]string{
	"Empty1":         {"type": "string", "blank": "true"},
	"Guest":          {"type": "string", "blank": "true"},
	"Description":    {"type": "string", "blank": "true"},
	"ResID":          {"type": "uint", "blank": "true"},
	"DateRes":        {"type": "string", "blank": "true"},
	"DateIn":         {"type": "string", "blank": "true"},
	"Empty3":         {"type": "string", "blank": "true"},
	"DateOut":        {"type": "string", "blank": "true"},
	"Adults":         {"type": "uint", "blank": "true"},
	"Child":          {"type": "uint", "blank": "true"},
	"Room":           {"type": "uint", "blank": "true"},
	"RoomType":       {"type": "string", "blank": "true"},
	"Rate":           {"type": "string", "blank": "true"},
	"RateName":       {"type": "string", "blank": "true"},
	"GroupCorporate": {"type": "string", "blank": "true"},
}

// GuestCSVRow contains fields which represents value
// exactly to the each row of roomkey input csv file
type GuestCSVRow struct {
	GuestID            string
	CreateDate         string
	Title              string
	GuestName          string
	FirstName          string
	LastName           string
	Address            string
	City               string
	StateProvince      string
	Country            string
	CountryNationality string
	ZipPostalCode      string
	Email              string
	MainPhone          string
	Mobile             string
	PhoneOther         string
	RentalStatus       string
	VIPStatus          string
	VIPDescription     string
	GuestNote          string
	Air                string
	Auto               string
	Loyalty            string
	Language           string
	VehicleMake        string
	VehicleModel       string
	VehicleColor       string
	LicensePlate       string
	EmailMarketing     string
	EmailGeneral       string
	Address2Name       string
	Address2           string
	City2              string
	StateProvince2     string
	Country2           string
	ZipPostalCode2     string
	Email2             string
	MainPhone2         string
	Mobile2            string
	PhoneOther2        string
	RoomNights         string
	Stays              string
	AvgStay            string
	Revenue            string
	Reservations       string
	Cancellations      string
}

// csvRowFieldRules is map contains rules for specific fields in roomkey
var guestCSVRowFieldRules = map[string]map[string]string{
	"GuestID":            {"type": "uint", "blank": "true"},
	"CreateDate":         {"type": "string", "blank": "true"},
	"Title":              {"type": "string", "blank": "true"},
	"GuestName":          {"type": "string", "blank": "true"},
	"FirstName":          {"type": "string", "blank": "true"},
	"LastName":           {"type": "string", "blank": "true"},
	"Address":            {"type": "string", "blank": "true"},
	"City":               {"type": "string", "blank": "true"},
	"StateProvince":      {"type": "string", "blank": "true"},
	"Country":            {"type": "string", "blank": "true"},
	"CountryNationality": {"type": "string", "blank": "true"},
	"ZipPostalCode":      {"type": "uint", "blank": "true"},
	"Email":              {"type": "string", "blank": "true"},
	"MainPhone":          {"type": "string", "blank": "true"},
	"Mobile":             {"type": "string", "blank": "true"},
	"PhoneOther":         {"type": "string", "blank": "true"},
	"RentalStatus":       {"type": "string", "blank": "true"},
	"VIPStatus":          {"type": "string", "blank": "true"},
	"VIPDescription":     {"type": "string", "blank": "true"},
	"GuestNote":          {"type": "string", "blank": "true"},
	"Air":                {"type": "string", "blank": "true"},
	"Auto":               {"type": "string", "blank": "true"},
	"Loyalty":            {"type": "string", "blank": "true"},
	"Language":           {"type": "string", "blank": "true"},
	"VehicleMake":        {"type": "string", "blank": "true"},
	"VehicleModel":       {"type": "string", "blank": "true"},
	"VehicleColor":       {"type": "string", "blank": "true"},
	"LicensePlate":       {"type": "string", "blank": "true"},
	"EmailMarketing":     {"type": "string", "blank": "true"},
	"EmailGeneral":       {"type": "string", "blank": "true"},
	"Address2Name":       {"type": "string", "blank": "true"},
	"Address2":           {"type": "string", "blank": "true"},
	"City2":              {"type": "string", "blank": "true"},
	"StateProvince2":     {"type": "string", "blank": "true"},
	"Country2":           {"type": "string", "blank": "true"},
	"ZipPostalCode2":     {"type": "string", "blank": "true"},
	"Email2":             {"type": "string", "blank": "true"},
	"MainPhone2":         {"type": "string", "blank": "true"},
	"Mobile2":            {"type": "string", "blank": "true"},
	"PhoneOther2":        {"type": "string", "blank": "true"},
	"RoomNights":         {"type": "string", "blank": "true"},
	"Stays":              {"type": "string", "blank": "true"},
	"AvgStay":            {"type": "string", "blank": "true"},
	"Revenue":            {"type": "string", "blank": "true"},
	"Reservations":       {"type": "string", "blank": "true"},
	"Cancellations":      {"type": "string", "blank": "true"},
}

// loadRoomKeyCSVRow used to load data from slice
// into CSVRow struct and return that struct
func loadRoomKeyCSVRow(csvCols []rcsv.CSVColumn, data []string) (bool, CSVRow) {
	csvRow := reflect.New(reflect.TypeOf(CSVRow{}))
	rowLoaded := false

	// fill data according to headers length
	for i := 0; i < len(csvCols); i++ {
		value := strings.TrimSpace(data[i])
		csvRow.Elem().Field(i).Set(reflect.ValueOf(value))
	}

	// if blank data has not been passed then only need to return true
	if (CSVRow{}) != csvRow.Elem().Interface().(CSVRow) {
		rowLoaded = true
	}

	return rowLoaded, csvRow.Elem().Interface().(CSVRow)
}

// validateRoomKeyCSVRow validates csv field of roomkey
// Dont perform validation while loading data in CSVRow struct
// (in LoadRoomKeyCSVRow function as it decides when to stop parsing)
func validateRoomKeyCSVRow(roomKeyCSVRow *CSVRow, rowIndex int) []error {
	rowErrs := []error{}

	// fill data according to headers length
	reflectedRoomKeyCSVRow := reflect.ValueOf(roomKeyCSVRow).Elem()

	for i := 0; i < len(csvCols); i++ {
		fieldName := reflect.TypeOf(*roomKeyCSVRow).Field(i).Name
		fieldValue := reflectedRoomKeyCSVRow.Field(i).Interface().(string)
		err := validateCSVField(fieldName, fieldValue, rowIndex+1)
		if err != nil {
			rowErrs = append(rowErrs, err)
		}
	}

	return rowErrs
}

// validateCSVField validates csv field of roomkey
func validateCSVField(field string, value string, rowIndex int) error {
	rule, ok := csvRowFieldRules[field]

	// if not found then simple return
	if !ok {
		return nil
	}

	fieldType, fieldBlankAllow := rule["type"], rule["blank"]

	// check with blank rule
	if fieldBlankAllow == "true" && value == "" {
		return nil
	}

	// if blank is not allowed and value is blank then return with error
	if fieldBlankAllow == "false" && value == "" {
		return fmt.Errorf("\"%s\" has blank value at row \"%d\"", field, rowIndex)
	}

	// check with field type
	switch fieldType {
	case "int":
		ok := core.IsIntString(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid integer number value at row \"%d\"", field, rowIndex)
		}
		return nil
	case "uint":
		ok := core.IsUIntString(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid positive integer number value at row \"%d\"", field, rowIndex)
		}
		return nil
	case "float":
		ok := core.IsFloatString(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid integer number value at row \"%d\"", field, rowIndex)
		}
		return nil
	case "email":
		ok := core.IsValidEmail(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid email value at row \"%d\"", field, rowIndex)
		}
		return nil
	case "phone":
		ok := core.IsValidPhone(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid phone number value at row \"%d\"", field, rowIndex)
		}
		return nil
	case "date":
		_, err := rlib.StringToDate(value)
		if err != nil {
			return fmt.Errorf("\"%s\" has no valid date value at row \"%d\"", field, rowIndex)
		}
		return nil
	case "rentable_status":
		//ok, _ := IsValidRentableStatus(value)
		ok := true
		if !ok {
			return fmt.Errorf("\"%s\" has no valid rentable status value at row \"%d\"", field, rowIndex)
		}
		return nil
	default:
		return nil
	}

}

// loadRoomKeyCSVRow used to load data from slice
// into CSVRow struct and return that struct
func loadGuestInfoCSVRow(csvCols []rcsv.CSVColumn, data []string) (bool, GuestCSVRow) {
	csvRow := reflect.New(reflect.TypeOf(GuestCSVRow{}))
	rowLoaded := false

	// fill data according to headers length
	for i := 0; i < len(csvCols); i++ {
		value := strings.TrimSpace(data[i])
		csvRow.Elem().Field(i).Set(reflect.ValueOf(value))
	}

	// if blank data has not been passed then only need to return true
	if (GuestCSVRow{}) != csvRow.Elem().Interface().(GuestCSVRow) {
		rowLoaded = true
	}

	if !core.IsValidEmail(csvRow.Elem().FieldByName("Email").Interface().(string)) {
		csvRow.Elem().FieldByName("Email").Set(reflect.ValueOf(""))
	}
	return rowLoaded, csvRow.Elem().Interface().(GuestCSVRow)
}
