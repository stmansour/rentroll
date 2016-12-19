package onesite

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
	"strings"
)

// CreateRentalAgreementCSV create rental agreement csv temporarily
// write headers, used to load data from onesite csv
// return file pointer to call program
func CreateRentalAgreementCSV(
	CSVStore string,
	timestamp string,
	rentalAgreementStruct *core.RentalAgreementCSV,
) (*os.File, *csv.Writer, bool) {

	var done = false

	// get path of rentalAgreement csv file
	filePrefix := prefixCSVFile["rental_agreement"]
	fileName := filePrefix + timestamp + ".csv"
	rentalAgreementCSVFilePath := path.Join(CSVStore, fileName)

	// try to create file and return with error if occurs any
	rentalAgreementCSVFile, err := os.Create(rentalAgreementCSVFilePath)
	if err != nil {
		log.Println(err)
		return nil, nil, done
	}

	// create csv writer
	rentalAgreementCSVWriter := csv.NewWriter(rentalAgreementCSVFile)

	// parse headers of rentalAgreementCSV using reflect
	rentalAgreementCSVHeaders := []string{}
	rentalAgreementCSVHeaders, ok := core.GetStructFields(rentalAgreementStruct)
	if !ok {
		log.Println("Unable to get struct fields for rentalAgreementCSV")
		return nil, nil, done
	}

	rentalAgreementCSVWriter.Write(rentalAgreementCSVHeaders)
	rentalAgreementCSVWriter.Flush()

	done = true

	return rentalAgreementCSVFile, rentalAgreementCSVWriter, done
}

// WriteRentalAgreementData used to write the data to csv file
// with avoiding duplicate data
func WriteRentalAgreementData(
	csvWriter *csv.Writer,
	csvRow *CSVRow,
	avoidData *[]string,
	currentTimeFormat string,
	suppliedValues map[string]string,
	rentalAgreementStruct *core.RentalAgreementCSV,
) {
	// TODO: need to decide how to avoid data
	// checkRentableStyle := csvRow.FloorPlan
	// Stylefound := core.StringInSlice(checkRentableStyle, *avoidData)

	// // if style found then simplay return otherwise continue
	// if Stylefound {
	//  return
	// }

	// *avoidData = append(*avoidData, checkRentableStyle)

	// make rentable data from userSuppliedValues and defaultValues
	rentableDefaultData := map[string]string{}
	for k, v := range suppliedValues {
		rentableDefaultData[k] = v
	}

	// get csv row data
	ok, csvRowData := GetRentalAgreementCSVRow(
		csvRow, rentalAgreementStruct,
		currentTimeFormat, rentableDefaultData,
	)
	if ok {
		csvWriter.Write(csvRowData)
		// TODO: make sure to verify the usage of flush is correct or not
		csvWriter.Flush()
	}
}

// GetRentalAgreementCSVRow used to create RentalAgreement
// csv row from onesite csv
func GetRentalAgreementCSVRow(
	oneSiteRow *CSVRow,
	fieldMap *core.RentalAgreementCSV,
	timestamp string,
	DefaultValues map[string]string,
) (bool, []string) {

	// take initial variable
	ok := false

	// ======================================
	// Load rentalAgreement's data from onesiterow data
	// ======================================
	reflectedOneSiteRow := reflect.ValueOf(oneSiteRow).Elem()
	reflectedRentalAgreementFieldMap := reflect.ValueOf(fieldMap).Elem()

	// length of RentalAgreementCSV
	rRTLength := reflectedRentalAgreementFieldMap.NumField()

	// return data array
	dataMap := make(map[int]string)

	for i := 0; i < rRTLength; i++ {
		// get rentalAgreement field
		rentalAgreementField := reflectedRentalAgreementFieldMap.Type().Field(i)

		// if rentalAgreementField value exist in DefaultValues map
		// then set it first
		suppliedValue, found := DefaultValues[rentalAgreementField.Name]
		if found {
			dataMap[i] = suppliedValue
		}

		// =========================================================
		// this condition has been put here because it's mapping field does not exist
		// =========================================================
		if rentalAgreementField.Name == "PayorSpec" {
			payorSpec, ok := GetPayorSpec(oneSiteRow)
			if ok {
				dataMap[i] = payorSpec
			} else {
				// TODO: verify that what to do in false case
				dataMap[i] = payorSpec
			}
		}
		if rentalAgreementField.Name == "UserSpec" {
			userSpec, ok := GetUserSpec(oneSiteRow)
			if ok {
				dataMap[i] = userSpec
			} else {
				// TODO: verify that what to do in false case
				dataMap[i] = userSpec
			}
		}
		if rentalAgreementField.Name == "RentableSpec" {
			rentableSpec, ok := GetRentableSpec(oneSiteRow)
			if ok {
				dataMap[i] = rentableSpec
			} else {
				// TODO: verify that what to do in false case
				dataMap[i] = rentableSpec
			}
		}

		// get mapping field
		MappedFieldName := reflectedRentalAgreementFieldMap.FieldByName(rentalAgreementField.Name).Interface().(string)

		// if has not value then continue
		if !reflectedOneSiteRow.FieldByName(MappedFieldName).IsValid() {
			continue
		}

		// get field by mapping field name and then value
		OneSiteFieldValue := reflectedOneSiteRow.FieldByName(MappedFieldName).Interface()
		dataMap[i] = OneSiteFieldValue.(string)
	}

	dataArray := []string{}

	for i := 0; i < rRTLength; i++ {
		dataArray = append(dataArray, dataMap[i])
	}
	ok = true
	return ok, dataArray
}

// GetPayorSpec used to get payor spec in format of rentroll system
func GetPayorSpec(
	csvRow *CSVRow,
) (string, bool) {

	// TODO: verify if validation required here
	ok := false

	orderedFields := []string{}

	// TODO: decide what to append here as payor
	// NOTE: right now just going with email address
	// append payor
	orderedFields = append(orderedFields, csvRow.Email)
	// append start date
	orderedFields = append(orderedFields, csvRow.LeaseStart)
	// append end date
	orderedFields = append(orderedFields, csvRow.LeaseEnd)

	ok = true
	if ok {
		return strings.Join(orderedFields, ","), ok
	}

	return ",,", ok
}

// GetUserSpec used to get user spec in format of rentroll system
func GetUserSpec(
	csvRow *CSVRow,
) (string, bool) {

	// TODO: verify if validation required here
	ok := false

	orderedFields := []string{}

	// TODO: decide what to append here as user
	// NOTE: right now just going with email address
	// append user
	orderedFields = append(orderedFields, csvRow.Email)
	// append start date
	orderedFields = append(orderedFields, csvRow.LeaseStart)
	// append end date
	orderedFields = append(orderedFields, csvRow.LeaseEnd)

	ok = true
	if ok {
		return strings.Join(orderedFields, ","), ok
	}

	return ",,", ok
}

// GetRentableSpec used to get rentable spec in format of rentroll system
func GetRentableSpec(
	csvRow *CSVRow,
) (string, bool) {

	// TODO: verify if validation required here
	ok := false

	orderedFields := []string{}

	// append rentable
	orderedFields = append(orderedFields, csvRow.Unit)
	// append contractrent
	orderedFields = append(orderedFields, csvRow.Rent)

	ok = true
	if ok {
		return strings.Join(orderedFields, ","), ok
	}

	return ",", ok
}
