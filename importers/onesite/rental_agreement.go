package onesite

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
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

	// get csv row data
	ok, csvRowData := GetRentalAgreementCSVRow(
		csvRow, rentalAgreementStruct,
		currentTimeFormat,
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
