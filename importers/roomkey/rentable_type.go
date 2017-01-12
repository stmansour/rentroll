package roomkey

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rlib"
	//"strconv"
	"time"
)

// CreateRentableTypeCSV create rentabletype csv temporarily
// write headers, used to load data from roomkey csv
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
	rentableTypeCSVHeaders := []string{}
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
	business *rlib.Business,
) {
	// get style
	checkRentableTypeStyle := csvRow.RoomType
	Stylefound := core.StringInSlice(checkRentableTypeStyle, *avoidData)

	// if style found then simplay return otherwise continue
	if Stylefound {
		return
	}

	*avoidData = append(*avoidData, checkRentableTypeStyle)

	currentYear, currentMonth, currentDate := currentTime.Date()
	DtStart := fmt.Sprintf("%d/%d/%d", currentMonth, currentDate, currentYear)
	DtStop := "12/31/9999" // no end date
	// DtStart := "1/1/" + strconv.Itoa(currentYear)
	// DtStop := "1/1/" + strconv.Itoa(currentYear+1)
	// DtStart := time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentTime.Location())
	// DtStop := time.Date(currentYear+1, 1, 1, 0, 0, 0, 0, currentTime.Location())

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
		traceCSVData[*recordCount] = rowIndex
	}
}

// GetRentableTypeCSVRow used to create rentabletype
// csv row from roomkey csv
func GetRentableTypeCSVRow(
	roomKeyRow *CSVRow,
	fieldMap *core.RentableTypeCSV,
	timestamp string,
	DefaultValues map[string]string,
) (bool, []string) {

	// take initial variable
	ok := false

	// ======================================
	// Load rentableType's data from roomKeyRow data
	// ======================================
	reflectedroomKeyRow := reflect.ValueOf(roomKeyRow).Elem()
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
		//  panic("coudln't get mapping field")
		// }

		// if has not value then continue
		if !reflectedroomKeyRow.FieldByName(MappedFieldName).IsValid() {
			continue
		}

		// get field by mapping field name and then value
		RoomKeyFieldValue := reflectedroomKeyRow.FieldByName(MappedFieldName).Interface()
		dataMap[i] = RoomKeyFieldValue.(string)
	}

	dataArray := []string{}

	for i := 0; i < rRTLength; i++ {
		dataArray = append(dataArray, dataMap[i])
	}
	ok = true
	return ok, dataArray
}
