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

// GetOneSiteMapping reads json file and loads
// field mapping structure in go for further usage
func GetOneSiteMapping(OneSiteFieldMap *CSVFieldMap) error {

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

// LoadOneSiteCSV loads the values from the supplied csv file and creates rlib.Business records
// as needed.
func LoadOneSiteCSV(userSuppliedValues map[string]string) ([]error, []error) {

	// vars
	var (
		errorList []error
		csvErrors []error
	)

	// funcname
	funcname := "LoadOneSiteCSV"

	// ===============================
	// validation on user supplied values
	// ===============================
	var BUD = userSuppliedValues["BUD"]
	var business = rlib.GetBusinessByDesignation(BUD)
	if business.BID == 0 {
		errorList = append(errorList,
			fmt.Errorf("Supplied Business Unit Designation does not exists"))
		return errorList, csvErrors
	}

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
	err := GetOneSiteMapping(&OneSiteFieldMap)
	if err != nil {
		errorList = append(errorList, core.ErrInternal)
		rlib.Ulog("Error <ONESITE FIELD MAPPING>: %s\n", err.Error())
		return errorList, csvErrors
	}

	// ##############################
	// # PHASE 1 : SPLITTING DATA IN CSV FILES #
	// ##############################

	// if dataValidationError is true throughout rows
	// do not perform any furter operation
	dataValidationError := false

	// ================================
	// First loop for validation on csv
	// ================================

	// load csv file and get data from csv
	t := rlib.LoadCSV(userSuppliedValues["OneSiteCSV"])

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
				errorList = append(errorList, err)
				rlib.Ulog("Error <ONESITE CSV COLUMN VALIDATION>: %s\n", err.Error())
				return errorList, csvErrors
			}

			// ######################
			// VALIDATION on data value, type
			// ######################
			rowLoaded, csvRow := LoadOneSiteCSVRow(csvCols, t[i][:OneSiteColumnLength])

			// NOTE: might need to change logic, if t[i] contains blank data that we should
			// stop the loop as we have to skip rest of the rows (please look at onesite csv)
			if !rowLoaded {
				rlib.Ulog("No more data for onesite csv loading\n")
				break
			}

			// if row is loaded successfully then do validation over fields
			rowErrs := ValidateOneSiteCSVRow(&csvRow, i)
			if len(rowErrs) > 0 {
				dataValidationError = true
				csvErrors = append(csvErrors, rowErrs...)
			}
		}
	}

	// if there is any error in data validation then return from here
	// do not perform any further action
	if dataValidationError {
		return errorList, csvErrors
	}

	// ====================================
	// BEFORE GOES TO SECOND LOOP
	// PERFORM REQUIRED OPERATIONS HERE
	// ====================================

	// get created rentabletype csv and writer pointer
	rentableTypeCSVFile, rentableTypeCSVWriter, ok :=
		CreateRentableTypeCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.RentableTypeCSV,
		)
	if !ok {
		errorList = append(errorList, core.ErrInternal)
		return errorList, csvErrors
	}

	// get created people csv and writer pointer
	peopleCSVFile, peopleCSVWriter, ok :=
		CreatePeopleCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.PeopleCSV,
		)
	if !ok {
		errorList = append(errorList, core.ErrInternal)
		return errorList, csvErrors
	}

	// get created people csv and writer pointer
	rentableCSVFile, rentableCSVWriter, ok :=
		CreateRentableCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.RentableCSV,
		)
	if !ok {
		errorList = append(errorList, core.ErrInternal)
		return errorList, csvErrors
	}

	// get created rental agreement csv and writer pointer
	rentalAgreementCSVFile, rentalAgreementCSVWriter, ok :=
		CreateRentalAgreementCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.RentalAgreementCSV,
		)
	if !ok {
		errorList = append(errorList, core.ErrInternal)
		return errorList, csvErrors
	}

	// get created customAttibutes csv and writer pointer
	customAttributeCSVFile, customAttributeCSVWriter, ok :=
		CreateCustomAttibutesCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.CustomAttributeCSV,
		)
	if !ok {
		errorList = append(errorList, core.ErrInternal)
		return errorList, csvErrors
	}

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

	// customAttributesRefData holds the data for future operation to insert
	// custom attribute ref in system for each rentableType
	// so we identify each element in this list via Style value
	customAttributesRefData := map[string]CARD{}

	// ================================
	// Second loop for splitting data of csv
	// Create csv files required for rentroll
	// ================================
	// in second round do split

	for i := skipRowsCount + 1; i < len(t); i++ {
		rowLoaded, csvRow := LoadOneSiteCSVRow(csvCols, t[i][:OneSiteColumnLength])

		// NOTE: might need to change logic, if t[i] contains blank data that we should
		// stop the loop as we have to skip rest of the rows (please look at onesite csv)
		if !rowLoaded {
			rlib.Ulog("No more data to parse\n")
			break
		}

		// Write data to file of rentabletype
		WriteRentableTypeCSVData(
			rentableTypeCSVWriter,
			&csvRow,
			&avoidDuplicateRentableTypeData,
			currentTime,
			currentTimeFormat,
			userSuppliedValues,
			&OneSiteFieldMap.RentableTypeCSV,
			customAttributesRefData,
			&business,
		)

		// Write data to file of people
		WritePeopleCSVData(
			peopleCSVWriter,
			&csvRow,
			&avoidDuplicatePeopleData,
			currentTimeFormat,
			userSuppliedValues,
			&OneSiteFieldMap.PeopleCSV,
		)

		// Write data to file of rentable
		WriteRentableData(
			rentableCSVWriter,
			&csvRow,
			&avoidDuplicateRentableData,
			currentTime,
			currentTimeFormat,
			userSuppliedValues,
			&OneSiteFieldMap.RentableCSV,
		)

		// Write data to file of rentalAgreement
		WriteRentalAgreementData(
			rentalAgreementCSVWriter,
			&csvRow,
			&avoidDuplicateRentalAgreementData,
			currentTimeFormat,
			userSuppliedValues,
			&OneSiteFieldMap.RentalAgreementCSV,
		)

		// Write data to file of CustomAttribute
		WriteCustomAttributeData(
			customAttributeCSVWriter,
			&csvRow,
			avoidDuplicateCustomAttributeData,
			currentTimeFormat,
			userSuppliedValues,
			&OneSiteFieldMap.CustomAttributeCSV,
		)

	}

	// Close all files as we are done here with writing data
	rentableTypeCSVFile.Close()
	peopleCSVFile.Close()
	rentableCSVFile.Close()
	rentalAgreementCSVFile.Close()
	customAttributeCSVFile.Close()

	// ########################
	// # PHASE 2 : RCSV LOADERS CALL #
	// ########################
	var h = []rcsv.CSVLoadHandler{
		{Fname: customAttributeCSVFile.Name(), Handler: rcsv.LoadCustomAttributesCSV},
		{Fname: rentableTypeCSVFile.Name(), Handler: rcsv.LoadRentableTypesCSV},
		{Fname: peopleCSVFile.Name(), Handler: rcsv.LoadPeopleCSV},
		{Fname: rentableCSVFile.Name(), Handler: rcsv.LoadRentablesCSV},
		{Fname: rentalAgreementCSVFile.Name(), Handler: rcsv.LoadRentalAgreementCSV},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			Errs := rrDoLoad(h[i].Fname, h[i].Handler)
			errorList = append(errorList, Errs...)
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
			errorList = append(errorList, err)
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
				errorList = append(errorList, err)
				continue
			}

			// insert custom attribute ref in system
			var a rlib.CustomAttributeRef
			a.ElementType = rlib.ELEMRENTABLETYPE
			a.ID = rt.RTID
			a.CID = ca.CID
			err = rlib.InsertCustomAttributeRef(&a)
			if err != nil {
				errorList = append(errorList, err)
				continue
			}
		}
	}
	// ##################################
	// # PHASE 3 : CLEAR THE TEMPORARY CSV FILES #
	// ##################################

	// testmode is disabled then only remove temp files
	if userSuppliedValues["testmode"] == "0" {
		ClearSplittedTempCSVFiles(currentTimeFormat)
	}

	// RETURN
	fmt.Println("ONESITE CSV HAS SUCCESSFULLY LOADED!")
	return errorList, csvErrors
}

func rrDoLoad(fname string, handler func(string) []error) []error {
	Errs := handler(fname)
	return Errs
	// fmt.Print(rcsv.ErrlistToString(&m))
}

// RollBackSplitOperation func used to clear out the things
// that created by program temporarily while loading onesite data
//  and if any error occurs
func RollBackSplitOperation(timestamp string) {
	ClearSplittedTempCSVFiles(timestamp)
}

// ClearSplittedTempCSVFiles func used only to clear
// temporarily csv files created by program
func ClearSplittedTempCSVFiles(timestamp string) {
	for _, filePrefix := range prefixCSVFile {
		fileName := filePrefix + timestamp + ".csv"
		filePath := path.Join(SplittedCSVStore, fileName)
		os.Remove(filePath)
	}
}

// CSVHandler is main function to handle user uploaded
// csv and extract information
func CSVHandler(userSuppliedValues map[string]string) ([]error, string) {
	initErr := Init()
	rlib.Errcheck(initErr)
	// rlib.Ulog("Error <ONESITE INIT>: %s\n", initErr.Error())
	errorList, CSVErrs := LoadOneSiteCSV(userSuppliedValues)
	csvErrorFlag := false
	if len(CSVErrs) > 0 {
		csvErrorFlag = true
	}
	if csvErrorFlag {
		fmt.Println(ErrorReporting(&CSVErrs))
	}
	return errorList, "DOne"
}

func ErrorReporting(csvErrors *[]error) string {
	var tbl rlib.Table
	tbl.Init()
	tbl.AddColumn("Error", 150, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	for _, err := range *csvErrors {
		tbl.AddRow()
		tbl.Puts(-1, 0, err.Error())
	}
	return tbl.SprintTable(rlib.RPTTEXT)
}
