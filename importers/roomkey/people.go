package roomkey

import (
	"encoding/csv"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rlib"
	"strings"
)

// CreatePeopleCSV create people csv temporarily
// write headers, used to load data from roomkey csv
// return file pointer to call program
func CreatePeopleCSV(
	CSVStore string,
	timestamp string,
	peopleCSVStruct *core.PeopleCSV,
) (*os.File, *csv.Writer, bool) {

	var done = false

	// get path of people csv file
	filePrefix := prefixCSVFile["people"]
	fileName := filePrefix + timestamp + ".csv"
	peopleCSVFilePath := path.Join(CSVStore, fileName)

	// try to create file and return with error if occurs any
	peopleCSVFile, err := os.Create(peopleCSVFilePath)
	if err != nil {
		rlib.Ulog("Error <PEOPLE CSV>: %s\n", err.Error())
		return nil, nil, done
	}

	// create csv writer
	peopleCSVWriter := csv.NewWriter(peopleCSVFile)

	// parse headers of peopleCSV using reflect
	peopleCSVHeaders := []string{}
	peopleCSVHeaders, ok := core.GetStructFields(peopleCSVStruct)
	if !ok {
		rlib.Ulog("Error <PEOPLE CSV>: Unable to get struct fields for peopleCSV\n")
		return nil, nil, done
	}

	peopleCSVWriter.Write(peopleCSVHeaders)
	peopleCSVWriter.Flush()

	done = true

	return peopleCSVFile, peopleCSVWriter, done
}

// WritePeopleCSVData used to write the data to csv file
// with avoiding duplicate data
func WritePeopleCSVData(
	recordCount *int,
	rowIndex int,
	traceCSVData map[int]int,
	csvWriter *csv.Writer,
	csvRow *CSVRow,
	avoidData *[]string,
	currentTimeFormat string,
	suppliedValues map[string]string,
	peopleStruct *core.PeopleCSV,
) {
	// TODO: need to decide how to avoid duplicate data
	// checkRentableTypeStyle := csvRow.FloorPlan
	// Stylefound := core.StringInSlice(checkRentableTypeStyle, *avoidData)
	// if Stylefound {
	// 	return
	// } else {
	// 	*avoidData = append(*avoidData, checkRentableTypeStyle)
	// }

	// get csv row data
	ok, csvRowData := GetPeopleCSVRow(
		csvRow, peopleStruct,
		currentTimeFormat, suppliedValues,
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

// GetPeopleCSVRow used to create people
// csv row from roomkey csv data
func GetPeopleCSVRow(
	roomKeyRow *CSVRow,
	fieldMap *core.PeopleCSV,
	timestamp string,
	DefaultValues map[string]string,
) (bool, []string) {

	// take initial variable
	ok := false

	// ======================================
	// Load people's data from roomkeyrow data
	// ======================================
	reflectedRoomKeyRow := reflect.ValueOf(roomKeyRow).Elem()
	reflectedPeopleFieldMap := reflect.ValueOf(fieldMap).Elem()

	// length of PeopleCSV
	pplLength := reflectedPeopleFieldMap.NumField()

	// return data array
	dataMap := make(map[int]string)

	// Mark isCompany field 1 if company name is provided
	isCompany := ""
	if strings.TrimSpace(roomKeyRow.GroupCorporate) != "" {
		isCompany = "1"
	}

	for i := 0; i < pplLength; i++ {
		// get people field
		peopleField := reflectedPeopleFieldMap.Type().Field(i)

		// if peopleField value exist in DefaultValues map
		// then set it first
		suppliedValue, found := DefaultValues[peopleField.Name]
		if found {
			dataMap[i] = suppliedValue
		}

		// =========================================================
		// these conditions have been put here because it's mapping field does not exist
		// =========================================================
		if peopleField.Name == "FirstName" {
			nameSlice := strings.Split(roomKeyRow.Guest, ",")
			dataMap[i] = strings.TrimSpace(nameSlice[0])
		}
		if peopleField.Name == "LastName" {
			nameSlice := strings.Split(roomKeyRow.Guest, ",")
			if len(nameSlice) > 1 {
				dataMap[i] = strings.TrimSpace(nameSlice[1])
			} else {
				dataMap[i] = ""
			}
		}

		// Add description to Notes field of people
		if peopleField.Name == "Notes" {
			des := "Res. Id:" + roomKeyRow.ResID
			des += "\n" + strings.TrimSpace(roomKeyRow.Description)
			dataMap[i] = des
		}

		// Add isCompany field value
		if peopleField.Name == "IsCompany" {
			dataMap[i] = isCompany
		}

		// get mapping field
		MappedFieldName := reflectedPeopleFieldMap.FieldByName(peopleField.Name).Interface().(string)

		// if has not value then continue
		if !reflectedRoomKeyRow.FieldByName(MappedFieldName).IsValid() {
			continue
		}

		// get field by mapping field name and then value
		roomKeyFieldValue := reflectedRoomKeyRow.FieldByName(MappedFieldName).Interface()
		dataMap[i] = roomKeyFieldValue.(string)

	}

	dataArray := []string{}

	for i := 0; i < pplLength; i++ {
		dataArray = append(dataArray, dataMap[i])
	}
	ok = true
	return ok, dataArray
}
