package roomkey

import (
	"encoding/csv"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rlib"
	"strconv"
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
	guestData GuestCSVRow,
	tracePeopleNote map[int]string,
	traceDuplicatePeople map[string][]string,
	csvErrors map[int][]string,
) {

	// flag duplicate people
	rowName := strings.TrimSpace(csvRow.Guest)
	name := strings.ToLower(rowName)

	// flag for name of people who has no email or phone
	if name != "" {
		if core.StringInSlice(name, traceDuplicatePeople["name"]) {
			warnPrefix := "W:<" + core.DBTypeMapStrings[core.DBPeople] + ">:"
			// mark it as a warning so customer can validate it
			csvErrors[rowIndex] = append(csvErrors[rowIndex],
				warnPrefix+"There is at least one other person with the name \""+rowName+"\" "+
					"who also has no unique identifiers such as cell phone number or email.",
			)
		} else {
			traceDuplicatePeople["name"] = append(traceDuplicatePeople["name"], name)
		}
	}

	// get csv row data
	csvRowData := GetPeopleCSVRow(
		csvRow, peopleStruct,
		currentTimeFormat, suppliedValues,
		guestData, rowIndex, tracePeopleNote,
	)

	csvWriter.Write(csvRowData)
	csvWriter.Flush()

	// after write operation to csv,
	// entry this rowindex with unit value in the map
	*recordCount = *recordCount + 1
	traceCSVData[*recordCount+1] = rowIndex
}

// GetPeopleCSVRow used to create people
// csv row from roomkey csv data
func GetPeopleCSVRow(
	roomKeyRow *CSVRow,
	fieldMap *core.PeopleCSV,
	timestamp string,
	DefaultValues map[string]string,
	guestData GuestCSVRow,
	rowIndex int,
	tracePeopleNote map[int]string,
) []string {

	// ======================================
	// Load people's data from roomkeyrow data
	// ======================================
	reflectedRoomKeyRow := reflect.ValueOf(roomKeyRow).Elem()
	reflectedPeopleFieldMap := reflect.ValueOf(fieldMap).Elem()

	// length of PeopleCSV
	pplLength := reflectedPeopleFieldMap.NumField()

	// return data array
	dataMap := make(map[int]string)

	for i := 0; i < pplLength; i++ {
		// get people field
		peopleField := reflectedPeopleFieldMap.Type().Field(i)

		// if peopleField value exist in DefaultValues map
		// then set it first
		suppliedValue, found := DefaultValues[peopleField.Name]
		if found {
			dataMap[i] = suppliedValue
		}

		if guestData.GuestName != "" {
			if peopleField.Name == "FirstName" {
				dataMap[i] = guestData.FirstName
			}
			if peopleField.Name == "LastName" {
				dataMap[i] = guestData.LastName
			}
			if peopleField.Name == "PrimaryEmail" {
				if core.IsValidEmail(guestData.Email) {
					dataMap[i] = guestData.Email
				}
			}
			if peopleField.Name == "CellPhone" {
				dataMap[i] = guestData.MainPhone
			}
			if peopleField.Name == "Address" {
				dataMap[i] = guestData.Address
			}
			if peopleField.Name == "Address2" {
				dataMap[i] = guestData.Address2
			}
			if peopleField.Name == "City" {
				dataMap[i] = guestData.City
			}
			if peopleField.Name == "State" {
				dataMap[i] = guestData.StateProvince
			}
			if peopleField.Name == "PostalCode" {
				dataMap[i] = guestData.ZipPostalCode
			}
			if peopleField.Name == "Country" {
				dataMap[i] = guestData.Country
			}
			if peopleField.Name == "AlternateEmailAddress" {
				dataMap[i] = guestData.Address2
			}
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

		// Special notes for people to get TCID in future with below value

		// Add description to Notes field of people
		if peopleField.Name == "Notes" {
			des := roomkeyNotesPrefix + strconv.Itoa(rowIndex) + "." + descriptionFieldSep
			des += "Res:" + roomKeyRow.Res + "."
			if roomKeyRow.Description != "" {
				des += descriptionFieldSep + strings.TrimSpace(roomKeyRow.Description)
			}
			dataMap[i] = des
			tracePeopleNote[rowIndex] = des
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

	return dataArray
}
