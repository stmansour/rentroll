package onesite

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rlib"
	"strings"
	"time"
)

// CreateRentableCSV create rentable csv temporarily
// write headers, used to load data from onesite csv
// return file pointer to call program
func CreateRentableCSV(CSVStore string, timestamp string, rentableStruct *core.RentableCSV) (*os.File, *csv.Writer, bool) {
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
// INPUTS
//    recordCount
//    rowIndex
//    traceCSVData
//    csvWriter
//    csvRow
//    currentTime
//    currentTimeFormat
//    suppliedValues
//    rentableStruct
//    traceTCIDMap
//    csvErrors
//    rrStatus
//
// RETURNS
//------------------------------------------------------------------------------
func WriteRentableData(
	recordCount *int,
	rowIndex int,
	traceCSVData map[int]int,
	csvWriter *csv.Writer,
	csvRow *CSVRow,
	currentTime time.Time,
	currentTimeFormat string,
	suppliedValues map[string]string,
	rentableStruct *core.RentableCSV,
	traceTCIDMap map[int]string,
	csvErrors map[int][]string,
	rrStatus string,
) {
	// rlib.Console("Entered WriteRentableData\n")
	currentYear, currentMonth, currentDate := currentTime.Date()
	DtStart := fmt.Sprintf("%d/%d/%d", currentMonth, currentDate, currentYear)
	DtStop := rlib.ENDOFTIME.Format(rlib.RRDATEFMT3) // no end date    "12/31/9999"

	// make rentable data from userSuppliedValues and defaultValues
	rentableDefaultData := map[string]string{}
	for k, v := range suppliedValues {
		rentableDefaultData[k] = v
	}
	rentableDefaultData["DtStart"] = DtStart
	rentableDefaultData["DtStop"] = DtStop
	rentableDefaultData["TCID"] = traceTCIDMap[rowIndex]

	// flag warning that we are taking default values for least start, end dates
	// as they don't exists
	if rrStatus == "occupied" {
		if csvRow.LeaseStart == "" {
			warnPrefix := "W:<" + core.DBTypeMapStrings[core.DBRentable] + ">:"
			csvErrors[rowIndex] = append(csvErrors[rowIndex],
				warnPrefix+"No lease start date found. Using default value: "+DtStart,
			)
		}
		if csvRow.LeaseEnd == "" {
			warnPrefix := "W:<" + core.DBTypeMapStrings[core.DBRentable] + ">:"
			csvErrors[rowIndex] = append(csvErrors[rowIndex],
				warnPrefix+"No lease end date found. Using default value: "+DtStop,
			)
		}
	}
	// get csv row data
	csvRowData := GetRentableCSVRow(
		csvRow, rentableStruct,
		currentTimeFormat, rentableDefaultData,
	)

	// rlib.Console("csvRowData = %q\n", csvRowData)
	csvWriter.Write(csvRowData)
	csvWriter.Flush()

	// after write operation to csv,
	// entry this rowindex with unit value in the map
	*recordCount = *recordCount + 1

	// need to map on next row index of temp csv as first row is header line
	// and recordCount initialized with 0 value
	traceCSVData[*recordCount+1] = rowIndex

}

// GetRentableCSVRow used to create rentabletype
// csv row from onesite csv
func GetRentableCSVRow(
	oneSiteRow *CSVRow,
	fieldMap *core.RentableCSV,
	timestamp string,
	DefaultValues map[string]string,
) []string {

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

		// rlib.Console("GetRentableCSVRow: rentableField.Name = %s\n", rentableField.Name)

		// =========================================================
		// this condition has been put here because it's mapping field does not exist
		// =========================================================
		if rentableField.Name == "RentableTypeRef" {
			dataMap[i] = GetRentableTypeRef(oneSiteRow, DefaultValues)
		}
		if rentableField.Name == "RUserSpec" {
			// format is user, startDate, stopDate
			dataMap[i] = GetRUserSpec(oneSiteRow, DefaultValues)
		}
		if rentableField.Name == "RentableUseType" {
			// format is status, startDate, stopDate
			status, _ := GetRentableUseStatus(oneSiteRow, DefaultValues)
			// TODO: verify that what to do in false case
			// should return its original value or raise error???
			dataMap[i] = status
		}

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

		// NOTE: do business logic here on field which has mapping field

		dataMap[i] = OneSiteFieldValue.(string)
	}

	dataArray := []string{}

	for i := 0; i < rRTLength; i++ {
		dataArray = append(dataArray, dataMap[i])
	}

	return dataArray
}

// GetRUserSpec used to get ruser spec in format of rentroll system
//---------------------------------------------------------------------------
func GetRUserSpec(csvRow *CSVRow, defaults map[string]string) string {
	//-------------------------------------------------------------------
	// check if status is occupied then return only RUserSpec otherwise
	// just return "" (blank string, not ",," with two comma separated
	// blank string!)
	//-------------------------------------------------------------------
	if _, rrUseStatus, _ := IsValidRentableUseType(csvRow.UnitLeaseStatus); rrUseStatus != "occupied" {
		return ""
	}

	//-------------------------------------------------------------------
	// as rcsv loader automatically associate user from rental
	// agreement csv so leave it as blank (nearly all cases)
	//-------------------------------------------------------------------
	return ""
}

// GetRentableUseStatus used to get rentable status in format of rentroll system
//-------------------------------------------------------------------
func GetRentableUseStatus(csvRow *CSVRow, defaults map[string]string) (string, bool) {
	var rrUseStatus string
	ok := false
	orderedFields := []string{}
	// rlib.Console("Entered GetRentableUseStatus\n")

	//-------------------------------------------------------------------
	// first find that passed string contains any status key
	//-------------------------------------------------------------------
	// rlib.Console("csvRow.UnitLeaseStatus = %s\n", csvRow.UnitLeaseStatus)
	validStatus, _, rrUseStatus := IsValidRentableUseType(csvRow.UnitLeaseStatus)

	//-------------------------------------------------------------------
	// if contains then try to get status according rentroll system
	//-------------------------------------------------------------------
	if validStatus {
		ok = true
		orderedFields = append(orderedFields, rrUseStatus)         // append unitusetype
		orderedFields = append(orderedFields, defaults["DtStart"]) // append today start date
		orderedFields = append(orderedFields, "")                  // append end date unspecified
		return strings.Join(orderedFields, ","), ok
	}

	// rlib.Console("GetRentableUseStatus:  invalid status exit\n")
	return ",,", ok
}

// GetRentableTypeRef used to get rentable type ref in format of rentroll system
//---------------------------------------------------------------------------------
func GetRentableTypeRef(csvRow *CSVRow, defaults map[string]string) string {
	orderedFields := []string{}
	orderedFields = append(orderedFields, csvRow.FloorPlan)    // append floor plan
	orderedFields = append(orderedFields, defaults["DtStart"]) // append today date
	orderedFields = append(orderedFields, "")                  // append end date as unspecified
	return strings.Join(orderedFields, ",")
}
