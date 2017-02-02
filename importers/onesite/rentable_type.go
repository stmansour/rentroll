package onesite

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rlib"
	"strconv"
	"time"
)

// CreateRentableTypeCSV create rentabletype csv temporarily
// write headers, used to load data from onesite csv
// return file pointer to call program
func CreateRentableTypeCSV(
	CSVStore string,
	timestamp string,
	rt *core.RentableTypeCSV,
) (*os.File, *csv.Writer, bool) {

	var done = false

	// get path of rentable csv file
	filePrefix := prefixCSVFile["rentable_types"]
	fileName := filePrefix + timestamp + ".csv"
	rentableTypeCSVFilePath := path.Join(CSVStore, fileName)

	// try to create file and return with error if occurs any
	rentableTypeCSVFile, err := os.Create(rentableTypeCSVFilePath)
	if err != nil {
		rlib.Ulog("Error <RENTABLE TYPE CSV>: %s\n", err.Error())
		return nil, nil, done
	}

	// create csv writer
	rentableTypeCSVWriter := csv.NewWriter(rentableTypeCSVFile)

	// parse headers of rentableTypeCSV using reflect
	rentableTypeCSVHeaders, ok := core.GetStructFields(rt)
	if !ok {
		rlib.Ulog("Error <RENTABLE TYPE CSV>: Unable to get struct fields for rentableTypeCSV\n")
		return nil, nil, done
	}

	rentableTypeCSVWriter.Write(rentableTypeCSVHeaders)
	rentableTypeCSVWriter.Flush()

	done = true

	return rentableTypeCSVFile, rentableTypeCSVWriter, done
}

// WriteRentableTypeCSVData used to write the data to csv file
// with avoiding duplicate data
func WriteRentableTypeCSVData(
	recordCount *int,
	rowIndex int,
	traceCSVData map[int]int,
	csvWriter *csv.Writer,
	csvRow *CSVRow,
	avoidData *[]string,
	currentTime time.Time,
	currentTimeFormat string,
	suppliedValues map[string]string,
	rt *core.RentableTypeCSV,
	customAttributesRefData map[string]CARD,
	business *rlib.Business,
) {
	// get style
	checkRentableTypeStyle := csvRow.FloorPlan
	Stylefound := core.StringInSlice(checkRentableTypeStyle, *avoidData)

	// if style found then simplay return otherwise continue
	if Stylefound {
		return
	}

	*avoidData = append(*avoidData, checkRentableTypeStyle)

	// insert CARD for this style in customAttributesRefData
	// no need to verify err, it has been passed already
	// through first loop in main program
	sqft, _ := strconv.ParseInt(csvRow.SQFT, 10, 64)
	tempCard := CARD{
		BID:      business.BID,
		Style:    checkRentableTypeStyle,
		SqFt:     sqft,
		RowIndex: rowIndex,
	}
	customAttributesRefData[checkRentableTypeStyle] = tempCard

	currentYear, currentMonth, currentDate := currentTime.Date()
	DtStart := fmt.Sprintf("%d/%d/%d", currentMonth, currentDate, currentYear)
	// DtStart := fmt.Sprintf("%02d/%02d/%04d", currentMonth, currentDate, currentYear)
	DtStop := "12/31/9999" // no end date

	// make rentableType data from userSuppliedValues and defaultValues
	rentableTypeDefaultData := map[string]string{}
	for k, v := range suppliedValues {
		rentableTypeDefaultData[k] = v
	}
	rentableTypeDefaultData["DtStart"] = DtStart
	rentableTypeDefaultData["DtStop"] = DtStop

	// get csv row data
	ok, csvRowData := GetRentableTypeCSVRow(
		csvRow, rt,
		currentTimeFormat, rentableTypeDefaultData,
	)
	if ok {
		csvWriter.Write(csvRowData)
		csvWriter.Flush()

		// after write operation to csv,
		// entry this rowindex with unit value in the map
		*recordCount = *recordCount + 1

		// need to map on next row index of temp csv as first row is header line
		// and recordCount initialized with 0 value
		traceCSVData[*recordCount+1] = rowIndex
	}
}

// GetRentableTypeCSVRow used to create rentabletype
// csv row from onesite csv
func GetRentableTypeCSVRow(
	oneSiteRow *CSVRow,
	fieldMap *core.RentableTypeCSV,
	timestamp string,
	DefaultValues map[string]string,
) []string {

	// ======================================
	// Load rentableType's data from onesiterow data
	// ======================================
	reflectedOneSiteRow := reflect.ValueOf(oneSiteRow).Elem()
	reflectedRentableTypeFieldMap := reflect.ValueOf(fieldMap).Elem()

	// length of RentableTypeCSV
	rRTLength := reflectedRentableTypeFieldMap.NumField()

	// return data array
	dataMap := make(map[int]string)

	for i := 0; i < rRTLength; i++ {
		// get rentableType field
		rentableTypeField := reflectedRentableTypeFieldMap.Type().Field(i)

		// if rentableTypeField value exist in DefaultValues map
		// then set it first
		suppliedValue, found := DefaultValues[rentableTypeField.Name]
		if found {
			dataMap[i] = suppliedValue
		}

		// get mapping field if not found then panic error
		MappedFieldName := reflectedRentableTypeFieldMap.FieldByName(rentableTypeField.Name).Interface().(string)
		// MappedFieldName, ok := reflectedRentableTypeFieldMap.FieldByName(rentableTypeField.Name).Interface().(string)
		// if !ok {
		// 	rlib.Ulog("Mapping Field not found", ...)
		// }

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

	return dataArray
}
