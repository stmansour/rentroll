/*
Package onesite imports data from onesite to rentroll.

This main program requires at least two inputs from users:
1. bud (required) (business unit designation)
2. csv (required) (onesite csv)
3. testmode (optional) (testmode doesn't clear temp files, right now!)
4. frequency (optional) (rent cycle frequency)
(
    0: one time only | 1: secondly | 2: minutely | 3: hourly |
    4: daily | 5: weekly | 6: monthly | 7: quarterly | 8: yearly |
)
5. proration (optional) (proration cycle)
6. gsrpc (optional) (GSRPC)

It handles imported csv via `CSVhandler` function.

CSVHandler accepts csv file path, testmode and user supplied values.
It initializes config first for onesite loader via init() call.
All user's passed values should be validated in it first.
After that it calls main function `loadOneSiteCSV` and
then creates a report based on response of `loadOneSiteCSV` call.

loadOneSiteCSV first loads field mapping defined from mapper.json in struct.
Then after loading csv data, in first loop it skips rows that are meant for
onesite data to import and then performs data validation on onesite csv data.
If there is any error in validation then it just simply returns.
Before going to second iteration loop it performs some necessary operation including
get file pointers, writer pointers, declaring struct to avoid duplicate data,
declaring struct to trace the data in accordance of input csv, declaring count
variables for each db type, declare customAttribRefData struct.
Then it loads data into temporary files in favor of onesite rcsv loader to import
the data via calls of rcsv routines. After data import has been done, it will dump
customAttribRefs in rentroll in a manual way rather than importing via temp csv file.
At last, after all things done, it clears out all temp files from `temp_CSVs` dir and
returns the response.
*/
package onesite

import (
	"encoding/json"
	"fmt"
	"github.com/kardianos/osext"
	"io/ioutil"
	"os"
	"path"
	"rentroll/importers/core"
	"rentroll/rcsv"
	"rentroll/rlib"
	"sort"
	"strconv"
	"strings"
	"time"
)

// used to store temporary csv files
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

// getOneSiteMapping reads json file and loads
// field mapping structure in go for further usage
func getOneSiteMapping(OneSiteFieldMap *CSVFieldMap) error {

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}

	// read json file which contains mapping of onesite fields
	mapperFilePath := path.Join(folderPath, "mapper.json")

	fieldmap, err := ioutil.ReadFile(mapperFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fieldmap, OneSiteFieldMap)
	if err != nil {
		return err
	}

	return err
}

// loadOneSiteCSV loads the values from the supplied csv file and creates rlib.Business records
// as needed.
func loadOneSiteCSV(
	oneSiteCSV string,
	testMode int,
	userRRValues map[string]string,
	business *rlib.Business,
) ([]error, error) {

	// vars
	var (
		loadOneSiteError error
		csvErrors        []error
	)

	// funcname
	funcname := "loadOneSiteCSV"

	// get current timestamp used for creating csv files unique way
	currentTime := time.Now()

	// RFC3339Nano is const format defined in time package
	// <FORMAT> = <SAMPLE>
	// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	// it is helpful while creating unique files
	currentTimeFormat := currentTime.Format(time.RFC3339Nano)

	// ###################################
	// # INIT PHASE : LOAD FIELD MAP IN ONESITE MAP #
	// ###################################

	// csvCols and consts for all onesite csv fields are defined in
	// constant.go file

	// Onesite csv headers slice and load it from csvCols
	OneSiteCSVHeaders := []string{}
	for _, header := range csvCols {
		OneSiteCSVHeaders = append(OneSiteCSVHeaders, header.Name)
	}
	OneSiteColumnLength := len(OneSiteCSVHeaders)

	// load onesite mapping
	var OneSiteFieldMap CSVFieldMap
	err := getOneSiteMapping(&OneSiteFieldMap)
	if err != nil {
		loadOneSiteError = core.ErrInternal
		rlib.Ulog("Error <ONESITE FIELD MAPPING>: %s\n", err.Error())
		return csvErrors, loadOneSiteError
	}

	// ##############################
	// # PHASE 1 : SPLITTING DATA IN CSV FILES #
	// ##############################

	// this map is used to hold csvRow typed struct after data has been loaded in it from first loop iteration
	// so we have not to iteration over onesite csv again and can be re-used in second loop
	csvRowDataMap := map[int]*CSVRow{}

	// if dataValidationError is true throughout rows
	// do not perform any furter operation
	dataValidationError := false

	// ================================
	// First loop for validation on csv
	// ================================

	// load csv file and get data from csv
	t := rlib.LoadCSV(oneSiteCSV)

	// this count used to skip number of rows from the very top of csv
	var skipRowsCount int

	for i := 0; i < len(t); i++ {

		// Calculate SkipRowsCount in first loop
		// if found then assign value in it and for rest of the rows
		// do validate csv columns

		if skipRowsCount == 0 {
			CSVRowDataString := strings.Replace(
				strings.Join(t[i][:OneSiteColumnLength], ","),
				" ", "", -1)
			CSVHeaderString := strings.Replace(
				strings.Join(OneSiteCSVHeaders[:OneSiteColumnLength], ","),
				" ", "", -1)

			if CSVRowDataString == CSVHeaderString {
				skipRowsCount = i

				// if skipRowsCount found then jump to next rows
				// because headers don't have to be validate
				continue
			}
		} else {
			// if skipRowsCount found then do validation over csv rows

			x, err := rcsv.ValidateCSVColumns(csvCols, t[i][:OneSiteColumnLength], funcname, i)
			if x > 0 {
				csvErrors = append(csvErrors, err)
				rlib.Ulog("Error <ONESITE CSV COLUMN VALIDATION>: %s\n", err.Error())
			}

			// ######################
			// VALIDATION on data value, type
			// ######################
			rowLoaded, csvRow := loadOneSiteCSVRow(csvCols, t[i][:OneSiteColumnLength])

			// NOTE: might need to change logic, if t[i] contains blank data that we should
			// stop the loop as we have to skip rest of the rows (please look at onesite csv)
			if !rowLoaded {
				rlib.Ulog("No more data for onesite csv loading\n")
				break
			}

			// if row is loaded successfully then do validation over fields
			rowErrs := validateOneSiteCSVRow(&csvRow, i)
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
				csvRowDataMap[i+1] = &csvRow
			}
		}
	}

	// if there is any error in data validation then return from here
	// do not perform any further action
	if dataValidationError {
		return csvErrors, loadOneSiteError
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
			&OneSiteFieldMap.RentableTypeCSV,
		)
	if !ok {
		loadOneSiteError = core.ErrInternal
		return csvErrors, loadOneSiteError
	}

	// get created people csv and writer pointer
	peopleCSVFile, peopleCSVWriter, ok :=
		CreatePeopleCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.PeopleCSV,
		)
	if !ok {
		loadOneSiteError = core.ErrInternal
		return csvErrors, loadOneSiteError
	}

	// get created people csv and writer pointer
	rentableCSVFile, rentableCSVWriter, ok :=
		CreateRentableCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.RentableCSV,
		)
	if !ok {
		loadOneSiteError = core.ErrInternal
		return csvErrors, loadOneSiteError
	}

	// get created rental agreement csv and writer pointer
	rentalAgreementCSVFile, rentalAgreementCSVWriter, ok :=
		CreateRentalAgreementCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.RentalAgreementCSV,
		)
	if !ok {
		loadOneSiteError = core.ErrInternal
		return csvErrors, loadOneSiteError
	}

	// get created customAttibutes csv and writer pointer
	customAttributeCSVFile, customAttributeCSVWriter, ok :=
		CreateCustomAttibutesCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.CustomAttributeCSV,
		)
	if !ok {
		loadOneSiteError = core.ErrInternal
		return csvErrors, loadOneSiteError
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

	// TODO: decide which structure to avoid duplicate data of rentalAgreement
	// while creating rentalAgreement csv file
	avoidDuplicateRentalAgreementData := []string{}

	// avoidDuplicateCustomAttributeData is tricky map which holds the
	// duplicate data in slice for each field defined in customAttributeMap
	avoidDuplicateCustomAttributeData := map[string][]string{}
	for k := range customAttributeMap {
		avoidDuplicateCustomAttributeData[k] = []string{}
	}
	// --------------------------------------------------------------------------------------------------------- //

	// traceUnitMap holds records by which we can trace the unit with row index of csv
	// Unit would be unique in onesite imported csv
	// key: rowIndex of onesite csv, value: Unit value of each row of onesite csv
	traceUnitMap := map[int]string{}

	// --------------------------- trace csv records map ----------------------------
	// trace<TYPE>CSVMap used to hold records
	// by which we can traceout which records has been writtern to csv
	// with key of row index of <TARGET_TYPE> CSV, value of original's imported csv rowNumber
	traceRentableTypeCSVMap := map[int]int{}
	traceRentableCSVMap := map[int]int{}
	tracePeopleCSVMap := map[int]int{}
	traceRentalAgreementCSVMap := map[int]int{}
	traceCustomAttributeCSVMap := map[int]int{}
	// --------------------------------------------------------------------------------------------------------- //

	// --------------------------- csv record count ----------------------------
	// <TYPE>CSVRecordCount used to hold records count inserted in csv
	// initialize with 1 because first row contains headers in target generated csv
	RentableTypeCSVRecordCount := 1
	RentableCSVRecordCount := 1
	PeopleCSVRecordCount := 1
	RentalAgreementCSVRecordCount := 1
	CustomAttributeCSVRecordCount := 1
	// --------------------------------------------------------------------------------------------------------- //

	// customAttributesRefData holds the data for future operation to insert
	// custom attribute ref in system for each rentableType
	// so we identify each element in this list via Style value
	customAttributesRefData := map[string]CARD{}

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

		// mark Unit value with row index value
		traceUnitMap[csvRowIndex] = csvRow.Unit

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
			&OneSiteFieldMap.RentableTypeCSV,
			customAttributesRefData,
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
			&OneSiteFieldMap.PeopleCSV,
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
			&OneSiteFieldMap.RentableCSV,
		)

		// Write data to file of rentalAgreement
		WriteRentalAgreementData(
			&RentalAgreementCSVRecordCount,
			csvRowIndex,
			traceRentalAgreementCSVMap,
			rentalAgreementCSVWriter,
			&csvRow,
			&avoidDuplicateRentalAgreementData,
			currentTimeFormat,
			userRRValues,
			&OneSiteFieldMap.RentalAgreementCSV,
		)

		// Write data to file of CustomAttribute
		WriteCustomAttributeData(
			&CustomAttributeCSVRecordCount,
			csvRowIndex,
			traceCustomAttributeCSVMap,
			customAttributeCSVWriter,
			&csvRow,
			avoidDuplicateCustomAttributeData,
			currentTimeFormat,
			userRRValues,
			&OneSiteFieldMap.CustomAttributeCSV,
		)

	}

	// ---------------------------- closing files -------------------------- //
	// Close all files as we are done here with writing data
	rentableTypeCSVFile.Close()
	peopleCSVFile.Close()
	rentableCSVFile.Close()
	rentalAgreementCSVFile.Close()
	customAttributeCSVFile.Close()
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
		{Fname: customAttributeCSVFile.Name(), Handler: rcsv.LoadCustomAttributesCSV, TraceDataMap: "traceCustomAttributeCSVMap"},
		{Fname: rentableTypeCSVFile.Name(), Handler: rcsv.LoadRentableTypesCSV, TraceDataMap: "traceRentableTypeCSVMap"},
		{Fname: peopleCSVFile.Name(), Handler: rcsv.LoadPeopleCSV, TraceDataMap: "tracePeopleCSVMap"},
		{Fname: rentableCSVFile.Name(), Handler: rcsv.LoadRentablesCSV, TraceDataMap: "traceRentableCSVMap"},
		{Fname: rentalAgreementCSVFile.Name(), Handler: rcsv.LoadRentalAgreementCSV, TraceDataMap: "traceRentalAgreementCSVMap"},
	}

	// getIndexAndUnit used to get index and unit value from trace<TYPE>CSVMap map
	getIndexAndUnit := func(traceDataMap string, index int) (int, string) {
		var onesiteIndex int
		var unit string
		switch traceDataMap {
		case "traceCustomAttributeCSVMap":
			if onesiteIndex, ok := traceCustomAttributeCSVMap[index]; ok {
				if unit, ok := traceUnitMap[onesiteIndex]; ok {
					return onesiteIndex, unit
				}
				return onesiteIndex, unit
			}
			return onesiteIndex, unit
		case "traceRentableTypeCSVMap":
			if onesiteIndex, ok := traceRentableTypeCSVMap[index]; ok {
				if unit, ok := traceUnitMap[onesiteIndex]; ok {
					return onesiteIndex, unit
				}
				return onesiteIndex, unit
			}
			return onesiteIndex, unit
		case "tracePeopleCSVMap":
			if onesiteIndex, ok := tracePeopleCSVMap[index]; ok {
				if unit, ok := traceUnitMap[onesiteIndex]; ok {
					return onesiteIndex, unit
				}
				return onesiteIndex, unit
			}
			return onesiteIndex, unit
		case "traceRentableCSVMap":
			if onesiteIndex, ok := traceRentableCSVMap[index]; ok {
				if unit, ok := traceUnitMap[onesiteIndex]; ok {
					return onesiteIndex, unit
				}
				return onesiteIndex, unit
			}
			return onesiteIndex, unit
		case "traceRentalAgreementCSVMap":
			if onesiteIndex, ok := traceRentalAgreementCSVMap[index]; ok {
				if unit, ok := traceUnitMap[onesiteIndex]; ok {
					return onesiteIndex, unit
				}
				return onesiteIndex, unit
			}
			return onesiteIndex, unit
		default:
			return onesiteIndex, unit
		}
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			Errs := rrDoLoad(h[i].Fname, h[i].Handler)
			for _, err := range Errs {
				fmt.Println(err.Error())
				// skip warnings about already existing records
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
					// now get the original row index of imported onesite csv and Unit value
					onesiteIndex, unit := getIndexAndUnit(h[i].TraceDataMap, lineNo)
					// generate new error
					err = fmt.Errorf("%s at row \"%d\" with unit \"%s\"", errText, onesiteIndex, unit)
					// append it into csvErrors
					csvErrors = append(csvErrors, err)
				} else {
					rlib.Ulog(fmt.Sprintf("Error <%s>: %s", h[i].Fname, err.Error()))
				}
			}
		}
	}

	// =======================================
	// INSERT CUSTOM ATTRIBUTE REF
	// =======================================

	// always sort keys
	var customAttributesRefDataKeys []string
	for k := range customAttributesRefData {
		customAttributesRefDataKeys = append(customAttributesRefDataKeys, k)
	}
	sort.Strings(customAttributesRefDataKeys)

	for _, key := range customAttributesRefDataKeys {
		// find rentableType
		refData := customAttributesRefData[key]
		rt, err := rlib.GetRentableTypeByStyle(refData.Style, refData.BID)
		if err != nil {
			unit, _ := traceUnitMap[refData.RowIndex]
			rlib.Ulog("Error <CUSTOMREF INSERTION>: %s", err)
			err = fmt.Errorf("Error while inserting custom attribute ref data at row \"%d\" with unit \"%s\"", refData.RowIndex, unit)
			csvErrors = append(csvErrors, err)
			continue
		}

		// for all custom attribute defined in custom_attrib.go
		// find custom attribute ID
		for _, customAttributeConfig := range customAttributeMap {
			t, _ := strconv.ParseInt(customAttributeConfig["ValueType"], 10, 64)
			n := customAttributeConfig["Name"]
			v := strconv.Itoa(int(refData.SqFt))
			u := customAttributeConfig["Units"]
			ca := rlib.GetCustomAttributeByVals(t, n, v, u)
			if ca.CID == 0 {
				unit, _ := traceUnitMap[refData.RowIndex]
				rlib.Ulog("Error <CUSTOMREF INSERTION>: %s", "CUSTOM ATTRIBUTE NOT FOUND IN DB")
				err := fmt.Errorf("Error while inserting custom attribute ref data at row \"%d\" with unit \"%s\"", refData.RowIndex, unit)
				csvErrors = append(csvErrors, err)
				continue
			}

			// insert custom attribute ref in system
			var a rlib.CustomAttributeRef
			a.ElementType = rlib.ELEMRENTABLETYPE
			a.ID = rt.RTID
			a.CID = ca.CID

			// check that record already exists, if yes then just continue
			// without accounting it as an error
			ref := rlib.GetCustomAttributeRef(a.ElementType, a.ID, a.CID)
			if ref.ElementType == a.ElementType && ref.CID == a.CID && ref.ID == a.ID {
				unit, _ := traceUnitMap[refData.RowIndex]
				errText := fmt.Sprintf(
					"This reference already exists. No changes were made. at row \"%d\" with unit \"%s\"",
					refData.RowIndex, unit)
				rlib.Ulog("Error <CUSTOMREF INSERTION>: %s", errText)
				continue
			}

			err := rlib.InsertCustomAttributeRef(&a)
			if err != nil {
				unit, _ := traceUnitMap[refData.RowIndex]
				rlib.Ulog("Error <CUSTOMREF INSERTION>: %s", err)
				err = fmt.Errorf("Error while inserting custom attribute ref data at row \"%d\" with unit \"%s\"", refData.RowIndex, unit)
				csvErrors = append(csvErrors, err)
				continue
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
	return csvErrors, loadOneSiteError
}

func rrDoLoad(fname string, handler func(string) []error) []error {
	Errs := handler(fname)
	return Errs
	// fmt.Print(rcsv.ErrlistToString(&m))
}

// rollBackImportOperation func used to clear out the things
// that created by program temporarily while loading onesite data
//  and if any error occurs
func rollBackImportOperation(timestamp string) {
	// TODO
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
		loadOneSiteError error
	)

	// init values
	CSVLoaded = true

	// ---------------------- some initialization for loadOneSiteCSV function ------------------
	initErr := Init()
	if initErr != nil {
		rlib.Ulog("Error <ONESITE INIT>: %s\n", initErr.Error())
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
		return CSVLoaded, CSVReport, loadOneSiteError
	}
	// --------------------------------------------------------------------------------------------------------- //

	// ---------------------- call onesite loader ----------------------------------------
	CSVErrs, loadOneSiteError = loadOneSiteCSV(CSV, TestMode, userRRValues, &business)

	// check if there any errors from onesite loader
	if len(CSVErrs) > 0 {
		CSVLoaded = false
		CSVReport = errorReporting(&CSVErrs)
	}
	if loadOneSiteError != nil {
		CSVLoaded = false
	}
	// if csv is not loaded properly then do rollbackoperation
	// and return with errors
	if !CSVLoaded {
		// TODO: rollBackImportOperation
		return CSVLoaded, CSVReport, loadOneSiteError
	}
	// --------------------------------------------------------------------------------------------------------- //

	//------------------------ Now do all the reporting ----------------------------
	var r = []rcsv.CSVReporterInfo{
		{ReportNo: 5, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportRentableTypes, Bid: business.BID},
		{ReportNo: 6, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportRentables, Bid: business.BID},
		{ReportNo: 7, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportPeople, Bid: business.BID},
		{ReportNo: 9, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportRentalAgreements, Bid: business.BID},
		{ReportNo: 14, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportCustomAttributes, Bid: business.BID},
		{ReportNo: 15, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportCustomAttributeRefs, Bid: business.BID},
	}

	for i := 0; i < len(r); i++ {
		CSVReport += r[i].Handler(&r[i])
		CSVReport += strings.Repeat("-", 80)
		CSVReport += "\n"
	}
	// --------------------------------------------------------------------------------------------------------- //

	// RETURN
	return CSVLoaded, CSVReport, loadOneSiteError

}

// errorReporting used to report the errors for onesite csv
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
