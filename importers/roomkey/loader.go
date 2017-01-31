package roomkey

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"rentroll/importers/core"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/rrpt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kardianos/osext"
)

// SplittedCSVStore is used to store temporary csv files
var SplittedCSVStore string

// Init configure required settings
func Init() error {
	// #############
	// CSV STORE CHECK
	// #############
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}

	// get path of splitted csv store
	SplittedCSVStore = path.Join(folderPath, splittedCSVStoreName)

	// if splittedcsvstore not exist then create it
	if _, err := os.Stat(SplittedCSVStore); os.IsNotExist(err) {
		os.MkdirAll(SplittedCSVStore, 0700)
	}
	return err
}

// getRoomKeyMapping reads json file and loads
// field mapping structure in go for further usage
func getRoomKeyMapping(RoomKeyFieldMap *CSVFieldMap) error {

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		// log.Fatal(err)
		panic("Unable to get current filename")
	}

	// read json file which contains mapping of roomkey fields
	mapperFilePath := path.Join(folderPath, "mapper.json")

	fieldmap, err := ioutil.ReadFile(mapperFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fieldmap, RoomKeyFieldMap)
	if err != nil {
		// fmt.Errorf("%s", err)
		panic(err)
	}
	return err
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// loadRoomKeyCSV loads the values from the supplied csv file and creates rlib.Business records
// as needed.
func loadRoomKeyCSV(
	roomKeyCSV string,
	testMode int,
	userRRValues map[string]string,
	business *rlib.Business,
) ([]error, error) {

	// vars
	var (
		LoadRoomKeyError error
		csvErrors        []error
	)

	// funcname
	funcname := "loadRoomKeyCSV"

	// get current timestamp used for creating csv files unique way
	currentTime := time.Now()

	// RFC3339Nano is const format defined in time package
	// <FORMAT> = <SAMPLE>
	// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	// it is helpful while creating unique files
	currentTimeFormat := currentTime.Format(time.RFC3339Nano)

	// ###################################
	// # INIT PHASE : LOAD FIELD MAP IN ROOMKEY MAP #
	// ###################################

	// csvCols and consts for all roomkey csv fields are defined in
	// constant.go file

	// Roomkey csv headers slice and load it from constants.csvCols
	roomKeyCSVHeaders := []string{}
	for _, header := range csvCols {
		roomKeyCSVHeaders = append(roomKeyCSVHeaders, header.Name)
	}
	RoomKeyColumnLength := len(roomKeyCSVHeaders)

	// load roomkey mapping
	var RoomKeyFieldMap CSVFieldMap
	err := getRoomKeyMapping(&RoomKeyFieldMap)
	if err != nil {
		LoadRoomKeyError = core.ErrInternal
		rlib.Ulog("Error <ROOMKEY FIELD MAPPING>: %s\n", err.Error())
		return csvErrors, LoadRoomKeyError
	}

	// ##############################
	// # CLEAN the roomkey csv file and storing in tMap #
	// ##############################

	// load csv file and get data from csv
	t := rlib.LoadCSV(roomKeyCSV)

	// Map row data with index
	tMap := map[int]*[]string{}

	// Making CSVHeaderString from roomKeyCSVHeaders
	CSVHeaderString := strings.Replace(
		strings.Join(roomKeyCSVHeaders[:RoomKeyColumnLength], ","),
		" ", "", -1)

	// Storing indices of blank columns to ignore while loading data
	skipColumns := []int{}
	checkForHeader := true
	for i := 0; i < len(t); i++ {

		// Checking for row header
		if checkForHeader {

			// skipColumns is a slice storing indices of blank columns
			skipColumns = []int{}

			// Joining csv row data string
			CSVRowDataString := strings.Replace(
				strings.Join(t[i][:RoomKeyColumnLength+3], ","),
				" ", "", -1)

			// Detecting blank columns in csv data
			tmpCSvRowDataString := ""
			commaString := ""
			for commas := minBlankColumns; commas <= maxBlankColumns; commas++ {
				// Getting indices of commas
				commaString = strings.Repeat(",", commas)
				commasIndex := strings.Index(CSVRowDataString, commaString)

				if commasIndex >= 0 {
					tmpCSvRowDataString = strings.Replace(CSVRowDataString,
						commaString,
						strings.Repeat(","+dummyBlankColumnName, commas), -1)
					break
				}
			}

			CSVRowDataString = strings.Replace(CSVRowDataString, commaString, "", -1)

			// Matching csv row data string with roomkey header string
			if CSVRowDataString == CSVHeaderString {

				if tmpCSvRowDataString != "" {
					tmpSlice := strings.Split(tmpCSvRowDataString, ",")
					for skipIndex, val := range tmpSlice {
						if val == dummyBlankColumnName {
							skipColumns = append(skipColumns, skipIndex)
						}
					}
				}

				checkForHeader = false
				continue
			}

		} else {

			// If first column has some data, it must be a new page,
			// so marking check for headers flag to true
			if t[i][0] != "" {
				checkForHeader = true
				continue
			}

			// Skipping empty colums as specified in skipColumns
			if len(skipColumns) > 0 {
				tempRow := []string{}

				for rowindex, rowvalue := range t[i] {
					appendflag := true
					for _, columnIndex := range skipColumns {
						if rowindex == columnIndex {
							appendflag = false
							break
						}
					}
					if appendflag {
						tempRow = append(tempRow, rowvalue)
					}
				}
				t[i] = tempRow
			}

			// Storing csv data to map (index : row data)
			tMap[i] = &t[i]

			// Checking data in description column
			if t[i][2] != "" {

				// Appending description to actual data row
				// Description might be on multiple rows after data row
				for j := 1; j < i; j++ {
					if t[i-j][1] != "" {
						t[i-j][2] += t[i][2]

						// Making description row as nil as data appended to actual row
						tMap[i] = nil
						if j > 1 {
							tMap[i-j+1] = nil
						}
						break
					}
				}
			}
		}
	}

	// ##############################
	// # PHASE 1 : SPLITTING DATA IN CSV FILES #
	// ##############################

	// this map is used to hold csvRow typed struct after data has been loaded in it from first loop iteration
	// so we have not to iteration over roomkey csv again and can be re-used in second loop
	csvRowDataMap := map[int]*CSVRow{}

	// if dataValidationError is true throughout rows
	// do not perform any furter operation
	dataValidationError := false

	// ================================
	// First loop for validation on csv
	// ================================

	// Iterating over cleaned csv data
	for rowIndex, rowData := range tMap {

		if rowData == nil {
			continue
		}

		x, err := rcsv.ValidateCSVColumns(csvCols, *rowData, funcname, rowIndex)
		if x > 0 {
			csvErrors = append(csvErrors, err)
			rlib.Ulog("Error <ROOMKEY CSV COLUMN VALIDATION>: %s\n", err.Error())
		}

		// ######################
		// VALIDATION on data value, type
		// ######################
		rowLoaded, csvRow := loadRoomKeyCSVRow(csvCols, *rowData)

		// NOTE: might need to change logic, if t[i] contains blank data that we should
		// stop the loop as we have to skip rest of the rows (please look at roomkey csv)
		if !rowLoaded {
			rlib.Ulog("No more data for roomkey csv loading\n")
			break
		}

		// if row is loaded successfully then do validation over fields
		rowErrs := validateRoomKeyCSVRow(&csvRow, rowIndex)
		if len(rowErrs) > 0 {
			dataValidationError = true
			csvErrors = append(csvErrors, rowErrs...)
		}

		// if dataValidationError is false then only fill data into map
		// because anyways the program will return and rest of operation will not be performed if dataValidationError is true
		// csvRowDataMap is only used for second iteration
		// so no need to dump it in the map if validation fails from any row
		if !dataValidationError {
			// index increased by one as in to be matched with csv row number
			csvRowDataMap[rowIndex+1] = &csvRow
		}
	}

	// if there is any error in data validation then return from here
	// do not perform any further action
	if dataValidationError {
		return csvErrors, LoadRoomKeyError
	}

	// ====================================
	// BEFORE GOES TO SECOND LOOP
	// PERFORM REQUIRED OPERATIONS HERE
	// ====================================

	// ----------------------- create files and get csv writer object -----------------------
	// get created rentabletype csv and writer pointer
	rentableTypeCSVFile, rentableTypeCSVWriter, ok :=
		CreateRentableTypeCSV(
			SplittedCSVStore, currentTimeFormat,
			&RoomKeyFieldMap.RentableTypeCSV,
		)

	if !ok {
		LoadRoomKeyError = core.ErrInternal
		return csvErrors, LoadRoomKeyError
	}

	// get created people csv and writer pointer
	peopleCSVFile, peopleCSVWriter, ok :=
		CreatePeopleCSV(
			SplittedCSVStore, currentTimeFormat,
			&RoomKeyFieldMap.PeopleCSV,
		)
	if !ok {
		LoadRoomKeyError = core.ErrInternal
		return csvErrors, LoadRoomKeyError
	}

	// get created people csv and writer pointer
	rentableCSVFile, rentableCSVWriter, ok :=
		CreateRentableCSV(
			SplittedCSVStore, currentTimeFormat,
			&RoomKeyFieldMap.RentableCSV,
		)
	if !ok {
		LoadRoomKeyError = core.ErrInternal
		return csvErrors, LoadRoomKeyError
	}

	// --------------------------------------------------------------------------------------------------------- //

	// --------------------- avoid duplicate data structures --------------------
	// avoidDuplicateRentableTypeData used to keep track of rentableTypeData with Style field
	// so that duplicate entries can be avoided while creating rentableType csv file
	avoidDuplicateRentableTypeData := []string{}

	// TODO: decide which structure to avoid duplicate data of people
	// while creating people csv file
	avoidDuplicatePeopleData := []string{}

	// TODO: decide which structure to avoid duplicate data of rentable
	// while creating rentable csv file
	avoidDuplicateRentableData := []string{}

	// --------------------------------------------------------------------------------------------------------- //

	// --------------------------- trace csv records map ----------------------------
	// trace<TYPE>CSVMap used to hold records
	// by which we can traceout which records has been writtern to csv
	// with key of row index of <TARGET_TYPE> CSV, value of original's imported csv rowNumber
	traceRentableTypeCSVMap := map[int]int{}
	tracePeopleCSVMap := map[int]int{}
	traceRentableCSVMap := map[int]int{}
	// --------------------------------------------------------------------------------------------------------- //

	// --------------------------- csv record count ----------------------------
	// <TYPE>CSVRecordCount used to hold records count inserted in csv
	// initialize with 1 because first row contains headers in target generated csv
	RentableTypeCSVRecordCount := 1
	PeopleCSVRecordCount := 1
	RentableCSVRecordCount := 1
	// --------------------------------------------------------------------------------------------------------- //

	// ================================
	// Second loop for splitting data of csv
	// Create csv files required for rentroll
	// ================================
	// in second round do split

	// always sort keys
	var csvRowDataMapKeys []int
	for k := range csvRowDataMap {
		csvRowDataMapKeys = append(csvRowDataMapKeys, k)
	}
	sort.Ints(csvRowDataMapKeys)

	for _, csvRowIndex := range csvRowDataMapKeys {

		// load csvRow from dataMap
		csvRow := *csvRowDataMap[csvRowIndex]

		// Write data to file of rentabletype
		WriteRentableTypeCSVData(
			&RentableTypeCSVRecordCount,
			csvRowIndex,
			traceRentableTypeCSVMap,
			rentableTypeCSVWriter,
			&csvRow,
			&avoidDuplicateRentableTypeData,
			currentTime,
			currentTimeFormat,
			userRRValues,
			&RoomKeyFieldMap.RentableTypeCSV,
			business,
		)

		// Write data to file of people
		WritePeopleCSVData(
			&PeopleCSVRecordCount,
			csvRowIndex,
			tracePeopleCSVMap,
			peopleCSVWriter,
			&csvRow,
			&avoidDuplicatePeopleData,
			currentTimeFormat,
			userRRValues,
			&RoomKeyFieldMap.PeopleCSV,
		)

		// Write data to file of rentable
		WriteRentableData(
			&RentableCSVRecordCount,
			csvRowIndex,
			traceRentableCSVMap,
			rentableCSVWriter,
			&csvRow,
			&avoidDuplicateRentableData,
			currentTime,
			currentTimeFormat,
			userRRValues,
			&RoomKeyFieldMap.RentableCSV,
		)

	}

	// ---------------------------- closing files -------------------------- //
	// Close all files as we are done here with writing data
	rentableTypeCSVFile.Close()
	peopleCSVFile.Close()
	rentableCSVFile.Close()
	// --------------------------------------------------------------------------------------------------------- //

	// ########################
	// # PHASE 2 : RCSV LOADERS CALL #
	// ########################
	// CSVLoadHandler struct is for routines that want to table-ize their loading.
	type csvLoadHandler struct {
		Fname        string
		Handler      func(string) []error
		TraceDataMap string
	}

	// csv load handler
	var h = []csvLoadHandler{
		{Fname: rentableTypeCSVFile.Name(), Handler: rcsv.LoadRentableTypesCSV, TraceDataMap: "traceRentableTypeCSVMap"},
		{Fname: peopleCSVFile.Name(), Handler: rcsv.LoadPeopleCSV, TraceDataMap: "tracePeopleCSVMap"},
		{Fname: rentableCSVFile.Name(), Handler: rcsv.LoadRentablesCSV, TraceDataMap: "traceRentableCSVMap"},
	}

	// getIndex used to get index from trace<TYPE>CSVMap map
	getIndex := func(traceDataMap string, index int) int {
		var roomkeyIndex int
		switch traceDataMap {
		case "traceRentableTypeCSVMap":
			roomkeyIndex, _ := traceRentableTypeCSVMap[index]
			return roomkeyIndex
		case "tracePeopleCSVMap":
			roomkeyIndex, _ := tracePeopleCSVMap[index]
			return roomkeyIndex
		case "traceRentableCSVMap":
			roomkeyIndex, _ := traceRentableCSVMap[index]
			return roomkeyIndex
		default:
			return roomkeyIndex
		}
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			Errs := rrDoLoad(h[i].Fname, h[i].Handler)
			for _, err := range Errs {
				// skip warnings about already existing records
				//fmt.Errorf("\n\nERRRORRRR: %v", err.Error())
				if !strings.Contains(err.Error(), "already exists") {
					errText := err.Error()
					// split with separator `:`
					s := strings.Split(errText, ":")
					// remove first element from slice
					s = append(s[:0], s[1:]...)
					// now join with separator
					errText = strings.Join(s, "|")
					// split with separator `-`
					s = strings.Split(errText, "-")
					// get line number string
					lineNoStr := s[0]
					// remove space from lineNoStr string
					lineNoStr = strings.Replace(lineNoStr, " ", "", -1)
					// remove `lineno` text from lineNoStr string
					lineNoStr = strings.Replace(lineNoStr, "lineno", "", -1)
					// remove `line` text from lineNoStr string
					lineNoStr = strings.Replace(lineNoStr, "line", "", -1)
					// now it should contain number in string
					lineNo, err := strconv.Atoi(lineNoStr)
					if err != nil {
						// CRITICAL
						panic("rcsv loaders should do something about returning error format")
					}
					// remove first element from slice
					s = append(s[:0], s[1:]...)
					// now join with separator
					errText = strings.Join(s, "")
					// replace new line broker
					errText = strings.Replace(errText, "\n", "", -1)
					// now get the original row index of imported roomkey csv and Unit value
					roomkeyIndex := getIndex(h[i].TraceDataMap, lineNo)
					// generate new error
					err = fmt.Errorf("%s at row \"%d\"", errText, roomkeyIndex)
					// append it into csvErrors
					csvErrors = append(csvErrors, err)
				} else {
					rlib.Ulog(fmt.Sprintf("Error <%s>: %s", h[i].Fname, err.Error()))
				}
			}
		}
	}

	// ##################################
	// # PHASE 3 : CLEAR THE TEMPORARY CSV FILES #
	// ##################################
	// testmode is not enabled then only remove temp files
	if testMode != 1 {
		clearSplittedTempCSVFiles(currentTimeFormat)
	}

	// RETURN
	return csvErrors, LoadRoomKeyError

}

func rrDoLoad(fname string, handler func(string) []error) []error {
	Errs := handler(fname)
	return Errs
}

// rollBackSplitOperation func used to clear out the things
// that created by program temporarily while loading roomkey data
//  and if any error occurs
func rollBackSplitOperation(timestamp string) {
	clearSplittedTempCSVFiles(timestamp)
}

// clearSplittedTempCSVFiles func used only to clear
// temporarily csv files created by program
func clearSplittedTempCSVFiles(timestamp string) {
	for _, filePrefix := range prefixCSVFile {
		fileName := filePrefix + timestamp + ".csv"
		filePath := path.Join(SplittedCSVStore, fileName)
		os.Remove(filePath)
	}
}

// CSVHandler is main function to handle user uploaded
// csv and extract information
func CSVHandler(
	CSV string,
	TestMode int,
	userRRValues map[string]string,
) (bool, string, error) {

	// vars
	var (
		CSVReport        string
		CSVLoaded        bool
		CSVErrs          []error
		LoadRoomKeyError error
	)

	// init values
	CSVLoaded = true

	// ---------------------- some initialization for loadRoomKeyCSV function ------------------
	initErr := Init()
	if initErr != nil {
		rlib.Ulog("Error <ROOMKEY INIT>: %s\n", initErr.Error())
	}
	rlib.Errcheck(initErr)
	// --------------------------------------------------------------------------------------------------------- //

	// ---------------------- validation on user supplied values ------------------
	BUD := userRRValues["BUD"]
	business := rlib.GetBusinessByDesignation(BUD)
	if business.BID == 0 {
		CSVLoaded = false
		CSVErrs = append(CSVErrs,
			fmt.Errorf("Supplied Business Unit Designation does not exists"))
		CSVReport = errorReporting(&CSVErrs)
		return CSVLoaded, CSVReport, LoadRoomKeyError
	}
	// --------------------------------------------------------------------------------------------------------- //

	// ---------------------- call roomkey loader ----------------------------------------
	CSVErrs, LoadRoomKeyError = loadRoomKeyCSV(CSV, TestMode, userRRValues, &business)

	// check if there any errors from roomkey loader
	if len(CSVErrs) > 0 {
		CSVLoaded = false
		CSVReport = errorReporting(&CSVErrs)
	}
	if LoadRoomKeyError != nil {
		CSVLoaded = false
	}
	// if csv is not loaded properly then do rollbackoperation
	// and return with errors
	if !CSVLoaded {
		// TODO: rollBackImportOperation
		return CSVLoaded, CSVReport, LoadRoomKeyError
	}
	// --------------------------------------------------------------------------------------------------------- //

	//------------------------ Now do all the reporting ----------------------------
	var r = []rrpt.ReporterInfo{
		{ReportNo: 5, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportRentableTypes, Bid: business.BID},
		{ReportNo: 6, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportRentables, Bid: business.BID},
		{ReportNo: 7, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportPeople, Bid: business.BID},
	}

	for i := 0; i < len(r); i++ {
		CSVReport += r[i].Handler(&r[i])
		CSVReport += strings.Repeat("-", 80)
		CSVReport += "\n"
	}
	// --------------------------------------------------------------------------------------------------------- //

	// RETURN
	return CSVLoaded, CSVReport, LoadRoomKeyError

}

// errorReporting used to report the errors for roomkey csv
func errorReporting(csvErrors *[]error) string {
	var tbl rlib.Table
	tbl.Init()
	tbl.AddColumn("Error", 150, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for _, err := range *csvErrors {
		tbl.AddRow()
		tbl.Puts(-1, 0, err.Error())
	}
	return tbl.SprintTable(rlib.RPTTEXT)
}
