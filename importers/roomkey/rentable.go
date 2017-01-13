package roomkey

import (
	"encoding/csv"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rlib"
	"strings"
	"time"
)

// ===========
// CONSTANTS
// ===========

// RRRentableStatus is status for rentable in rentroll system
var RRRentableStatus = map[string]string{
	"unknown":        "0",
	"online":         "1",
	"admin":          "2",
	"employee":       "3",
	"owner occupied": "4",
	"offline":        "5",
}

// RentableStatusCSV is mapping for rentable status between roomkey and rentroll
var RentableStatusCSV = map[string]string{
	"vacant":   "online",
	"occupied": "online",
	"model":    "admin",
}

// CreateRentableCSV create rentable csv temporarily
// write headers, used to load data from roomkey csv
// return file pointer to call program
func CreateRentableCSV(
	CSVStore string,
	timestamp string,
	rentableStruct *core.RentableCSV,
) (*os.File, *csv.Writer, bool) {

	var done = false

	// get path of rentable csv file
	filePrefix := prefixCSVFile["rentable"]
	fileName := filePrefix + timestamp + ".csv"
	rentableCSVFilePath := path.Join(CSVStore, fileName)

	// try to create file and return with error if occurs any
	rentableCSVFile, err := os.Create(rentableCSVFilePath)
	if err != nil {
		rlib.Ulog("Error <RENTABLE CSV>: %s\n", err.Error())
		return nil, nil, done
	}

	// create csv writer
	rentableCSVWriter := csv.NewWriter(rentableCSVFile)

	// parse headers of rentableCSV using reflect
	rentableCSVHeaders := []string{}
	rentableCSVHeaders, ok := core.GetStructFields(rentableStruct)
	if !ok {
		rlib.Ulog("Error <RENTABLE CSV>: Unable to get struct fields for rentableCSV\n")
		return nil, nil, done
	}

	rentableCSVWriter.Write(rentableCSVHeaders)
	rentableCSVWriter.Flush()

	done = true

	return rentableCSVFile, rentableCSVWriter, done
}

// WriteRentableData used to write the data to csv file
// with avoiding duplicate data
func WriteRentableData(
	recordCount *int,
	rowIndex int,
	traceCSVData map[int]int,
	csvWriter *csv.Writer,
	csvRow *CSVRow,
	avoidData *[]string,
	currentTime time.Time,
	currentTimeFormat string,
	suppliedValues map[string]string,
	rentableStruct *core.RentableCSV,
) {
	// TODO: need to decide how to avoid data
	// checkRentableStyle := csvRow.FloorPlan
	// Stylefound := core.StringInSlice(checkRentableStyle, *avoidData)

	// // if style found then simplay return otherwise continue
	// if Stylefound {
	// 	return
	// }

	// *avoidData = append(*avoidData, checkRentableStyle)

	// make rentable data from userSuppliedValues and defaultValues
	rentableDefaultData := map[string]string{}
	for k, v := range suppliedValues {
		rentableDefaultData[k] = v
	}

	// Forming default rentable status string
	rentableDefaultData["RentableStatus"] = "1"

	dateIn := getFormattedDate(csvRow.Empty3)
	dateOut := getFormattedDate(csvRow.DateOut)

	rentableDefaultData["RentableStatus"] += "," + dateIn
	rentableDefaultData["RentableStatus"] += "," + dateOut

	// get csv row data
	ok, csvRowData := GetRentableCSVRow(
		csvRow, rentableStruct,
		currentTimeFormat, rentableDefaultData,
	)
	if ok {
		csvWriter.Write(csvRowData)
		csvWriter.Flush()

		// after write operation to csv,
		// entry this rowindex with unit value in the map
		*recordCount = *recordCount + 1
		traceCSVData[*recordCount] = rowIndex
	}
}

// GetRentableCSVRow used to create rentabletype
// csv row from roomkey csv
func GetRentableCSVRow(
	roomKeyRow *CSVRow,
	fieldMap *core.RentableCSV,
	timestamp string,
	DefaultValues map[string]string,
) (bool, []string) {

	// take initial variable
	ok := false

	// ======================================
	// Load rentable's data from roomkeyrow data
	// ======================================
	reflectedRoomKeyRow := reflect.ValueOf(roomKeyRow).Elem()
	reflectedRentableFieldMap := reflect.ValueOf(fieldMap).Elem()

	// length of RentableCSV
	rRTLength := reflectedRentableFieldMap.NumField()

	// return data array
	dataMap := make(map[int]string)

	for i := 0; i < rRTLength; i++ {
		// get rentable field
		rentableField := reflectedRentableFieldMap.Type().Field(i)

		// if rentableField value exist in DefaultValues map
		// then set it first
		suppliedValue, found := DefaultValues[rentableField.Name]
		if found {
			dataMap[i] = suppliedValue
		}

		// =========================================================
		// this condition has been put here because it's mapping field does not exist
		// =========================================================
		if rentableField.Name == "RentableTypeRef" {
			typeRef, ok := GetRentableTypeRef(roomKeyRow)
			if ok {
				dataMap[i] = typeRef
			} else {
				// TODO: verify that what to do in false case
				// dataMap[i] = "FailedRentableTypeRef"
				dataMap[i] = typeRef
			}
		}

		// get mapping field
		MappedFieldName := reflectedRentableFieldMap.FieldByName(rentableField.Name).Interface().(string)

		// if has not value then continue
		if !reflectedRoomKeyRow.FieldByName(MappedFieldName).IsValid() {
			continue
		}

		// get field by mapping field name and then value
		roomKeyFieldValue := reflectedRoomKeyRow.FieldByName(MappedFieldName).Interface()

		// ====================================================
		// this condition has been put here because it's mapping field exists
		// ====================================================

		// NOTE: do business logic here on field which has mapping field

		dataMap[i] = roomKeyFieldValue.(string)
	}

	dataArray := []string{}

	for i := 0; i < rRTLength; i++ {
		dataArray = append(dataArray, dataMap[i])
	}
	ok = true
	return ok, dataArray
}

// IsValidRentableStatus checks that passed string contains valid rentable status
// acoording to rentroll system
func IsValidRentableStatus(s string) (bool, string) {
	found := false
	var tempRS string
	// first find that passed string contains any status key
	a := strings.ToLower(s)
	for k, v := range RentableStatusCSV {
		if strings.Contains(a, k) {
			tempRS = v
			found = true
			break
		}
	}
	return found, tempRS
}

// GetRentableTypeRef used to get rentable type ref in format of rentroll system
func GetRentableTypeRef(
	csvRow *CSVRow,
) (string, bool) {

	// TODO: verify if validation required here
	ok := false

	orderedFields := []string{}

	// append room type
	orderedFields = append(orderedFields, csvRow.RoomType)
	// append date in
	orderedFields = append(orderedFields, getFormattedDate(csvRow.Empty3))
	// append date out
	orderedFields = append(orderedFields, getFormattedDate(csvRow.DateOut))

	ok = true
	if ok {
		return strings.Join(orderedFields, ","), ok
	}

	return ",,", ok
}

// getFormattedDate returns rentroll accepted date string
func getFormattedDate(
	dateString string,
) string {

	const shortForm = "02-Jan-2006"
	const layout = "2006-01-02"

	parsedDate, _ := time.Parse(shortForm, dateString)
	return parsedDate.Format(layout)

}
