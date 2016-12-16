package onesite

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"rentroll/importers/core"
	"strconv"
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

// RentableStatusCSV is mapping for rentable status between onesite and rentroll
var RentableStatusCSV = map[string]string{
	"vacant":   "online",
	"occupied": "online",
	"model":    "admin",
}

// CreateRentableCSV create rentable csv temporarily
// write headers, used to load data from onesite csv
// return file pointer to call program
func CreateRentableCSV(
	CSVStore string,
	timestamp string,
	rentableStruct *core.RentableCSV,
) (*os.File, *csv.Writer, bool) {

	var done = false

	// get path of rentable csv file
	rentableCSVFilePath := CSVStore + "/rentable_" + timestamp + ".csv"

	// try to create file and return with error if occurs any
	rentableCSVFile, err := os.Create(rentableCSVFilePath)
	if err != nil {
		log.Println(err)
		return nil, nil, done
	}

	// create csv writer
	rentableCSVWriter := csv.NewWriter(rentableCSVFile)

	// parse headers of rentableCSV using reflect
	rentableCSVHeaders := []string{}
	rentableCSVHeaders, ok := core.GetStructFields(rentableStruct)
	if !ok {
		log.Println("Unable to get struct fields for rentableCSV")
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

	currentYear, _, _ := currentTime.Date()
	DtStart := "1/1/" + strconv.Itoa(currentYear)
	DtStop := "1/1/" + strconv.Itoa(currentYear+1)
	// DtStart := time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentTime.Location())
	// DtStop := time.Date(currentYear+1, 1, 1, 0, 0, 0, 0, currentTime.Location())

	// make rentable data from userSuppliedValues and defaultValues
	rentableDefaultData := map[string]string{}
	for k, v := range suppliedValues {
		rentableDefaultData[k] = v
	}
	rentableDefaultData["DtStart"] = DtStart
	rentableDefaultData["DtStop"] = DtStop

	// get csv row data
	ok, csvRowData := GetRentableCSVRow(
		csvRow, rentableStruct,
		currentTimeFormat, rentableDefaultData,
	)
	if ok {
		csvWriter.Write(csvRowData)
		// TODO: make sure to verify the usage of flush is correct or not
		csvWriter.Flush()
	}
}

// GetRentableCSVRow used to create rentabletype
// csv row from onesite csv
func GetRentableCSVRow(
	oneSiteRow *CSVRow,
	fieldMap *core.RentableCSV,
	timestamp string,
	DefaultValues map[string]string,
) (bool, []string) {

	// take initial variable
	ok := false

	// ======================================
	// Load rentable's data from onesiterow data
	// ======================================
	reflectedOneSiteRow := reflect.ValueOf(oneSiteRow).Elem()
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
			typeRef, ok := GetRentableTypeRef(oneSiteRow)
			if ok {
				dataMap[i] = typeRef
			} else {
				// TODO: verify that what to do in false case
				dataMap[i] = "FailedRentableTypeRef"
			}
		}
		// if rentableField.Name == "RUserSpec" {
		// 	// TODO
		// }

		// get mapping field
		MappedFieldName := reflectedRentableFieldMap.FieldByName(rentableField.Name).Interface().(string)

		// if has not value then continue
		if !reflectedOneSiteRow.FieldByName(MappedFieldName).IsValid() {
			continue
		}

		// get field by mapping field name and then value
		OneSiteFieldValue := reflectedOneSiteRow.FieldByName(MappedFieldName).Interface()

		// ====================================================
		// this condition has been put here because it's mapping field exists
		// ====================================================
		if rentableField.Name == "RentableStatus" {
			status, ok := GetRentableStatus(OneSiteFieldValue.(string))
			if ok {
				dataMap[i] = status
			} else {
				// TODO: verify that what to do in false case -> should return its original value or raise error
				dataMap[i] = OneSiteFieldValue.(string)
			}
		} else {
			dataMap[i] = OneSiteFieldValue.(string)
		}
	}

	dataArray := []string{}

	for i := 0; i < rRTLength; i++ {
		dataArray = append(dataArray, dataMap[i])
	}
	ok = true
	return ok, dataArray
}

// GetRentableStatus used to get rentable status in format of rentroll system
func GetRentableStatus(s string) (string, bool) {
	var tempRS, rRS string
	found, ok := false, false

	// first find that passed string contains any status key
	a := strings.ToLower(s)
	for k, v := range RentableStatusCSV {
		if strings.ContainsAny(a, k) {
			tempRS = v
			found = true
			break
		}
	}

	// if contains then try to get status according rentroll system
	if found {
		rRS, ok = RRRentableStatus[tempRS]
	}

	// return true if ok
	if ok {
		return rRS, true
	}
	return s, false
}

// GetRentableTypeRef used to get rentable type ref in format of rentroll system
func GetRentableTypeRef(
	csvRow *CSVRow,
) (string, bool) {

	// TODO: verify if validation required here
	ok := false

	orderedFields := []string{}

	// append floor plan
	orderedFields = append(orderedFields, csvRow.FloorPlan)
	// append lease start
	orderedFields = append(orderedFields, csvRow.LeaseStart)
	// append lease end
	orderedFields = append(orderedFields, csvRow.LeaseEnd)

	ok = true
	return strings.Join(orderedFields, ","), ok

}
