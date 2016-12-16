package onesite

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"errors"
	"os"
	"path"
	"rentroll/rcsv"
	"rentroll/rlib"
	"runtime"
	"strings"
	"time"
)

// used to store temporary csv files
var SplittedCSVStore string

// Init configure required settings
func Init() {
	// Caller returns program counter, filename, line no, ok
	_, filename, _, ok := runtime.Caller(1)
	if ok == false {
		panic("Unable to get current filename")
	}

	// get path of splitted csv store
	SplittedCSVStore = path.Join(path.Dir(filename), "/"+splittedCSVStoreName)

	// if splittedcsvstore not exist then create it
	if _, err := os.Stat(SplittedCSVStore); os.IsNotExist(err) {
		os.MkdirAll(SplittedCSVStore, 0700)
	}
}

// GetOneSiteMapping reads json file and loads
// field mapping structure in go for further usage
func GetOneSiteMapping(OneSiteFieldMap *CSVFieldMap) error {

	// Caller returns program counter, filename, line no, ok
	_, filename, _, ok := runtime.Caller(1)
	if ok == false {
		panic("Unable to get current filename")
	}

	// read json file which contains mapping of onesite fields
	mapperFilePath := path.Join(path.Dir(filename), "mapper.json")
	fieldmap, err := ioutil.ReadFile(mapperFilePath)
	if err != nil {
		// fmt.Errorf("File error: %v\n", err)
		panic(err)
		// return err
		// ???
	}
	err = json.Unmarshal(fieldmap, OneSiteFieldMap)
	if err != nil {
		// fmt.Errorf("%s", err)
		panic(err)
	}
	return err
}

// LoadOneSiteCSV loads the values from the supplied csv file and creates rlib.Business records
// as needed.
func LoadOneSiteCSV(userSuppliedValues map[string]string) ([]error, string) {

	// var errors and msg to return
	var errorList []error
	var msg string

	// funcname
	funcname := "LoadOneSiteCSV"

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
		errorList = append(errorList, err)
		msg = "Error while getting onesite field mapping"
		return errorList, msg
	}

	// ##############################
	// # PHASE 1 : SPLITTING DATA IN CSV FILES #
	// ##############################

	// get created rentabletype csv and writer pointer
	rentableTypeCSVFile, rentableTypeCSVWriter, ok :=
		CreateRentableTypeCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.RentableTypeCSV,
		)
	if !ok {
		// TODO: create errorlist in errors.go file to get error from that
		errorList = append(errorList, errors.New("Unable To create rentabletype csv file"))
		return errorList, "Unable to create rentabletype file"
	}

	// get created people csv and writer pointer
	peopleCSVFile, peopleCSVWriter, ok :=
		CreatePeopleCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.PeopleCSV,
		)
	if !ok {
		// TODO: create errorlist in errors.go file to get error from that
		errorList = append(errorList, errors.New("Unable To create people csv file"))
		return errorList, "Unable to create people file"
	}

	// get created people csv and writer pointer
	rentableCSVFile, rentableCSVWriter, ok :=
		CreateRentableCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.RentableCSV,
		)
	if !ok {
		// TODO: create errorlist in errors.go file to get error from that
		errorList = append(errorList, errors.New("Unable To create rentable csv file"))
		return errorList, "Unable to create rentable file"
	}

	// get created rental agreement csv and writer pointer
	rentalAgreementCSVFile, rentalAgreementCSVWriter, ok :=
		CreateRentalAgreementCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.RentalAgreementCSV,
		)
	if !ok {
		// TODO: create errorlist in errors.go file to get error from that
		errorList = append(errorList, errors.New("Unable To create rentalAgreement csv file"))
		return errorList, "Unable to create rentalAgreement file"
	}

	// get created customAttibutes csv and writer pointer
	customAttributeCSVFile, customAttributeCSVWriter, ok :=
		CreateCustomAttibutesCSV(
			SplittedCSVStore, currentTimeFormat,
			&OneSiteFieldMap.CustomAttributeCSV,
		)
	if !ok {
		// TODO: create errorlist in errors.go file to get error from that
		errorList = append(errorList, errors.New("Unable To create CustomAttribute csv file"))
		return errorList, "Unable to create CustomAttribute file"
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

		CSVRowDataString := strings.Replace(
			strings.Join(t[i][:OneSiteColumnLength], ","),
			" ", "", -1)
		CSVHeaderString := strings.Replace(
			strings.Join(OneSiteCSVHeaders[:OneSiteColumnLength], ","),
			" ", "", -1)

		if CSVRowDataString == CSVHeaderString {
			skipRowsCount = i
		}

		// if skipRowsCount found then do validation over csv rows
		if skipRowsCount > 0 {
			x, vrs := rcsv.ValidateCSVColumns(csvCols, t[i][:OneSiteColumnLength], funcname, i)
			if x > 0 {
				errorList = append(errorList, vrs)
				return errorList, msg
			}
		}
	}

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
			fmt.Println("\nNo more data to parse")
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
		)

		// Write data to file of people
		WritePeopleCSVData(
			peopleCSVWriter,
			&csvRow,
			&avoidDuplicatePeopleData,
			currentTimeFormat,
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
			&OneSiteFieldMap.RentalAgreementCSV,
		)

		// Write data to file of CustomAttribute
		WriteCustomAttributeData(
			customAttributeCSVWriter,
			&csvRow,
			avoidDuplicateCustomAttributeData,
			currentTimeFormat,
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

	// ##################################
	// # PHASE 3 : CLEAR THE TEMPORARY CSV FILES #
	// ##################################
	ClearSplittedTempCSVFiles(currentTimeFormat)

	// RETURN
	return errorList, msg
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
	for _, v := range prefixCSVFile {
		fileName := v + timestamp + ".csv"
		filePath := path.Join(SplittedCSVStore, fileName)
		err := os.Remove(filePath)
		fmt.Println(err)
	}
}

// CSVHandler is main function to handle user uploaded
// csv and extract information
func CSVHandler(userSuppliedValues map[string]string) ([]error, string) {
	Init()
	return LoadOneSiteCSV(userSuppliedValues)
}
