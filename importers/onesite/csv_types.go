package onesite

import (
	"reflect"
	"rentroll/importers/core"
	"rentroll/rcsv"
	"strings"
)

// CSVFieldMap is struct which con`~tains several categories
// used to store the data from onesite to rentroll system
type CSVFieldMap struct {
	RentableTypeCSV    core.RentableTypeCSV
	PeopleCSV          core.PeopleCSV
	RentableCSV        core.RentableCSV
	RentalAgreementCSV core.RentalAgreementCSV
	CustomAttributeCSV core.CustomAttributeCSV
}

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
	NoticeForDate   string
	MoveOut         string
	LeaseStart      string
	LeaseEnd        string
	MarketAddl      string
	DepOnHand       string
	Balance         string
	TotalCharges    string
	Rent            string
	WaterReImb      string
	Corp            string
	Discount        string
	Platinum        string
	Tax             string
	ElectricReImb   string
	Fire            string
	ConcSpecl       string
	WashDry         string
	EmplCred        string
	Short           string
	PetFee          string
	TrashReImb      string
	TermFee         string
	LakeView        string
	Utility         string
	Furn            string
	Mtom            string
	Referral        string
}

// // csvRowFieldRules is map contains rules for specific fields in onesite
// var csvRowFieldRules = map[string]map[string]string{
// 	"Unit":            {"type": "string", "blank": "false"},
// 	"FloorPlan":       {"type": "string", "blank": "false"},
// 	"UnitDesignation": {"type": "string", "blank": "true"},
// 	"SQFT":            {"type": "uint", "blank": "false"},
// 	// based on status value all of except this, will be validated
// 	// so don't defined rule for status here
// 	// "UnitLeaseStatus": {"type": "rentable_status", "blank": "true"},
// 	"Name":          {"type": "string", "blank": "false"},
// 	"PhoneNumber":   {"type": "phone", "blank": "true"},
// 	"Email":         {"type": "email", "blank": "true"},
// 	"MoveIn":        {"type": "date", "blank": "true"},
// 	"NoticeForDate": {"type": "string", "blank": "true"},
// 	"MoveOut":       {"type": "date", "blank": "true"},
// 	"LeaseStart":    {"type": "date", "blank": "false"},
// 	"LeaseEnd":      {"type": "date", "blank": "false"},
// 	"MarketAddl":    {"type": "float", "blank": "false"},
// 	"DepOnHand":     {"type": "float", "blank": "true"},
// 	"Balance":       {"type": "float", "blank": "true"},
// 	"TotalCharges":  {"type": "float", "blank": "true"},
// 	"Rent":          {"type": "float", "blank": "false"},
// 	"WaterReImb":    {"type": "float", "blank": "true"},
// 	"Corp":          {"type": "float", "blank": "true"},
// 	"Discount":      {"type": "float", "blank": "true"},
// 	"Platinum":      {"type": "float", "blank": "true"},
// 	"Tax":           {"type": "float", "blank": "true"},
// 	"ElectricReImb": {"type": "float", "blank": "true"},
// 	"Fire":          {"type": "float", "blank": "true"},
// 	"ConcSpecl":     {"type": "float", "blank": "true"},
// 	"WashDry":       {"type": "float", "blank": "true"},
// 	"EmplCred":      {"type": "float", "blank": "true"},
// 	"Short":         {"type": "float", "blank": "true"},
// 	"PetFee":        {"type": "float", "blank": "true"},
// 	"TrashReImb":    {"type": "float", "blank": "true"},
// 	"TermFee":       {"type": "float", "blank": "true"},
// 	"LakeView":      {"type": "float", "blank": "true"},
// 	"Utility":       {"type": "float", "blank": "true"},
// 	"Furn":          {"type": "float", "blank": "true"},
// 	"Mtom":          {"type": "float", "blank": "true"},
// 	"Referral":      {"type": "float", "blank": "true"},
// }

// loadOneSiteCSVRow used to load data from slice
// into CSVRow struct and return that struct
func loadOneSiteCSVRow(csvCols []rcsv.CSVColumn, data []string) (bool, CSVRow) {
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
