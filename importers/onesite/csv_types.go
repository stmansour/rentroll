package onesite

import (
	"fmt"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rcsv"
	"rentroll/rlib"
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

// csvRowFieldRules is map contains rules for specific fields in onesite
var csvRowFieldRules = map[string]map[string]string{
	"Unit":            {"type": "string", "blank": "false"},
	"FloorPlan":       {"type": "string", "blank": "false"},
	"UnitDesignation": {"type": "string", "blank": "true"},
	"SQFT":            {"type": "uint", "blank": "false"},
	// based on status value all of except this, will be validated
	// so don't defined rule for status here
	// "UnitLeaseStatus": {"type": "rentable_status", "blank": "true"},
	"Name":          {"type": "string", "blank": "false"},
	"PhoneNumber":   {"type": "phone", "blank": "true"},
	"Email":         {"type": "email", "blank": "true"},
	"MoveIn":        {"type": "date", "blank": "true"},
	"NoticeForDate": {"type": "string", "blank": "true"},
	"MoveOut":       {"type": "date", "blank": "true"},
	"LeaseStart":    {"type": "date", "blank": "false"},
	"LeaseEnd":      {"type": "date", "blank": "false"},
	"MarketAddl":    {"type": "float", "blank": "false"},
	"DepOnHand":     {"type": "float", "blank": "true"},
	"Balance":       {"type": "float", "blank": "true"},
	"TotalCharges":  {"type": "float", "blank": "true"},
	"Rent":          {"type": "float", "blank": "false"},
	"WaterReImb":    {"type": "float", "blank": "true"},
	"Corp":          {"type": "float", "blank": "true"},
	"Discount":      {"type": "float", "blank": "true"},
	"Platinum":      {"type": "float", "blank": "true"},
	"Tax":           {"type": "float", "blank": "true"},
	"ElectricReImb": {"type": "float", "blank": "true"},
	"Fire":          {"type": "float", "blank": "true"},
	"ConcSpecl":     {"type": "float", "blank": "true"},
	"WashDry":       {"type": "float", "blank": "true"},
	"EmplCred":      {"type": "float", "blank": "true"},
	"Short":         {"type": "float", "blank": "true"},
	"PetFee":        {"type": "float", "blank": "true"},
	"TrashReImb":    {"type": "float", "blank": "true"},
	"TermFee":       {"type": "float", "blank": "true"},
	"LakeView":      {"type": "float", "blank": "true"},
	"Utility":       {"type": "float", "blank": "true"},
	"Furn":          {"type": "float", "blank": "true"},
	"Mtom":          {"type": "float", "blank": "true"},
	"Referral":      {"type": "float", "blank": "true"},
}

// notOccupiedRentableValidateFields holds the list of fields which must be
// validated if rentable status is other than occupied
var notOccupiedRentableValidateFields = []string{
	"Unit", "FloorPlan", "SQFT", "MarketAddl",
}

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

// validateOneSiteCSVRow validates csv field of onesite
// Dont perform validation while loading data in CSVRow struct
// (in loadOneSiteCSVRow function as it decides when to stop parsing)
func validateOneSiteCSVRow(oneSiteCSVRow *CSVRow, rowIndex int) []error {
	rowErrs := []error{}

	// fill data according to headers length
	reflectedOneSiteCSVRow := reflect.ValueOf(oneSiteCSVRow).Elem()

	// perform validation based on rentable status
	ok, _ := IsValidRentableStatus(oneSiteCSVRow.UnitLeaseStatus)

	// TODO: need to clear on this what should importers do
	// when it encounters other type of rentablestatus
	// right now we are just throwing an error of bad value
	if !ok {
		statusErr := fmt.Errorf("\"%s\" has no valid rentable status value at row \"%d\" with unit \"%s\"",
			"UnitDesignation", rowIndex, oneSiteCSVRow.Unit)
		rowErrs = append(rowErrs, statusErr)
		return rowErrs
	}

	if strings.Contains(oneSiteCSVRow.UnitLeaseStatus, "occupied") {
		// if status is occupied then only perform validation over all fields
		for i := 0; i < len(csvCols); i++ {
			fieldName := reflect.TypeOf(*oneSiteCSVRow).Field(i).Name
			fieldValue := reflectedOneSiteCSVRow.Field(i).Interface().(string)
			err := validateCSVField(oneSiteCSVRow, fieldName, fieldValue, rowIndex+1)
			if err != nil {
				rowErrs = append(rowErrs, err)
			}
		}
	} else {
		// perform validation on fields defined in notOccupiedRentableValidateFields
		for _, fieldName := range notOccupiedRentableValidateFields {
			fieldValue := reflectedOneSiteCSVRow.FieldByName(fieldName).Interface().(string)
			err := validateCSVField(oneSiteCSVRow, fieldName, fieldValue, rowIndex+1)
			if err != nil {
				rowErrs = append(rowErrs, err)
			}
		}
	}

	return rowErrs
}

// validateCSVField validates csv field of onesite
func validateCSVField(oneSiteCSVRow *CSVRow, field string, value string, rowIndex int) error {
	rule, ok := csvRowFieldRules[field]

	// if not found then simple return
	if !ok {
		// TODO: verify this also
		return nil
	}

	fieldType, fieldBlankAllow := rule["type"], rule["blank"]

	// check with blank rule
	if fieldBlankAllow == "true" && value == "" {
		return nil
	}

	// if blank is not allowed and value is blank then return with error
	if fieldBlankAllow == "false" && value == "" {
		return fmt.Errorf("\"%s\" has blank value at row \"%d\" with unit \"%s\"", field, rowIndex, oneSiteCSVRow.Unit)
	}

	// check with field type
	switch fieldType {
	case "int":
		ok := core.IsIntString(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid integer number value at row \"%d\" with unit \"%s\"", field, rowIndex, oneSiteCSVRow.Unit)
		}
		return nil
	case "uint":
		ok := core.IsUIntString(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid positive integer number value at row \"%d\" with unit \"%s\"", field, rowIndex, oneSiteCSVRow.Unit)
		}
		return nil
	case "float":
		ok := core.IsFloatString(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid integer number value at row \"%d\" with unit \"%s\"", field, rowIndex, oneSiteCSVRow.Unit)
		}
		return nil
	case "email":
		ok := core.IsValidEmail(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid email value at row \"%d\" with unit \"%s\"", field, rowIndex, oneSiteCSVRow.Unit)
		}
		return nil
	case "phone":
		ok := core.IsValidPhone(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid phone number value at row \"%d\" with unit \"%s\"", field, rowIndex, oneSiteCSVRow.Unit)
		}
		return nil
	case "date":
		_, err := rlib.StringToDate(value)
		if err != nil {
			return fmt.Errorf("\"%s\" has no valid date value at row \"%d\" with unit \"%s\"", field, rowIndex, oneSiteCSVRow.Unit)
		}
		return nil
	case "rentable_status":
		ok, _ := IsValidRentableStatus(value)
		if !ok {
			return fmt.Errorf("\"%s\" has no valid rentable status value at row \"%d\" with unit \"%s\"", field, rowIndex, oneSiteCSVRow.Unit)
		}
		return nil
	default:
		return nil
	}
}
