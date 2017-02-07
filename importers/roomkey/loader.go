package roomkey

import (
	"encoding/json"
	"errors"
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

	"github.com/kardianos/osext"
)

// SplittedCSVStore is used to store temporary csv files
var SplittedCSVStore string

// getRoomKeyMapping reads json file and loads
// field mapping structure in go for further usage
func getRoomKeyMapping(RoomKeyFieldMap *CSVFieldMap) error {

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}

	// read json file which contains mapping of roomkey fields
	mapperFilePath := path.Join(folderPath, "mapper.json")

	fieldmap, err := ioutil.ReadFile(mapperFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fieldmap, RoomKeyFieldMap)
	return err
}

// loadRoomKeyCSV loads the values from the supplied csv file and creates rlib.Business records
// as needed.
func loadRoomKeyCSV(
	roomKeyCSV string,
	guestInfo map[string]*GuestCSVRow,
	testMode int,
	userRRValues map[string]string,
	business *rlib.Business,
	currentTime time.Time,
	currentTimeFormat string,
	summaryReport map[int]map[string]int,
) (map[int][]string, bool) {

	// returns csvError list, csv loaded?

	// returned csv errors should be in format
	// {
	// 	"rowIndex": ["E:errors",....., "W:warnings",....]
	// }
	// E stands for Error string, W stands for Warning string
	// UnitName can be accessible via traceUnitMap

	// =========================
	// DATA STRUCTURES AND VARS
	// =========================

	internalErrFlag := true
	csvErrors := map[int][]string{}

	// this holds the records for each row index
	csvRowDataMap := map[int]*CSVRow{}

	// --------------------------- trace csv records map ----------------------------
	// trace<TYPE>CSVMap used to hold records
	// by which we can traceout which records has been writtern to csv
	// with key of row index of <TARGET_TYPE> CSV, value of original's imported csv rowNumber
	traceRentableTypeCSVMap := map[int]int{}
	tracePeopleCSVMap := map[int]int{}
	traceRentableCSVMap := map[int]int{}
	traceRentalAgreementCSVMap := map[int]int{}

	// traceTCIDMap hold TCID for each people to be loaded via people csv
	// with reference of original roomkey csv
	traceTCIDMap := map[int]string{}

	// tracePeopleNote holds people note with reference of original roomkey csv
	tracePeopleNote := map[int]string{}

	// peopleCollisions holds count of people with same name
	peopleCollisions := map[string]int{}

	// traceDuplicatePeople holds records with unique string (name, email, phone)
	// with duplicant match at row
	// e.g.; {
	// 	"name": {"foo, bar": 3},
	// }
	traceDuplicatePeople := map[string][]string{
		"name": {},
	}

	// --------------------- avoid duplicate data structures -------------------- //
	// avoidDuplicateRentableTypeData used to keep track of rentableTypeData with Style field
	// so that duplicate entries can be avoided while creating rentableType csv file
	avoidDuplicateRentableTypeData := []string{}
	avoidDuplicatePeopleData := []string{}

	// --------------------------- csv record count ----------------------------
	// <TYPE>CSVRecordCount used to hold records count inserted in csv
	// initialize with 1 because first row contains headers in target generated csv
	// these are POSSIBLE record count that going to be imported
	RentableTypeCSVRecordCount := 0
	RentableCSVRecordCount := 0
	PeopleCSVRecordCount := 0
	RentalAgreementCSVRecordCount := 0

	// ===================================================
	// LOAD FIELD MAP AND GET HEADERS, LENGTH OF HEADERS
	// ===================================================

	// load roomkey mapping
	var RoomKeyFieldMap CSVFieldMap
	err := getRoomKeyMapping(&RoomKeyFieldMap)
	if err != nil {
		rlib.Ulog("Error <ROOMKEY FIELD MAPPING>: %s\n", err.Error())
		return csvErrors, internalErrFlag
	}

	// ===================================================
	// # CLEAN the roomkey csv file and storing in tMap #
	// ===================================================

	// load csv file and get data from csv
	t := rlib.LoadCSV(roomKeyCSV)

	// get headers with index map
	headersIndex := getCSVHeadersIndexMap()

	// this will be helpful while we have "description" type of row
	// so that we can put it in currentDataRowIndex's csvRow
	currentDataRowIndex := 0

	headersFirstOccurenceFound := false

	for rowIndex := 1; rowIndex <= len(t); rowIndex++ {

		// if it is header line then skip it
		if ok, csvHeadersWithCSVIndex := isRoomKeyHeaderLine(t[rowIndex-1]); ok {
			headersFirstOccurenceFound = true
			// get new headersIndex
			headersIndex = csvHeadersWithCSVIndex
			continue
		}

		// if first time headers are not detected then do continue
		if !headersFirstOccurenceFound {
			continue
		}

		// check it is page row
		if isRoomKeyPageRow(t[rowIndex-1]) {
			continue
		}

		// check it is description row
		if isRoomKeyDescriptionRow(t[rowIndex-1]) {
			csvRowDataMap[currentDataRowIndex].Description += descriptionFieldSep + t[rowIndex-1][rowTypeDetectionCSVIndex["description"]]
			continue
		}

		skipRow, csvRow := loadRoomKeyCSVRow(headersIndex, t[rowIndex-1])
		if skipRow {
			// in case blank row detected
			continue
		}

		// map this row as currentDataRowIndex and also hold a reference in datamap
		csvRowDataMap[rowIndex] = &csvRow
		currentDataRowIndex = rowIndex

	}

	// if csvRowDataMap is empty, that means data could not be parsed from csv
	if len(csvRowDataMap) == 0 {
		internalErrFlag = false
		csvErrors[-1] = append(csvErrors[-1], "There are no data rows present")
		return csvErrors, internalErrFlag
	}

	// ----------------------- create files and get csv writer object -----------------------
	// get created rentabletype csv and writer pointer
	rentableTypeCSVFile, rentableTypeCSVWriter, ok :=
		CreateRentableTypeCSV(
			TempCSVStore, currentTimeFormat,
			&RoomKeyFieldMap.RentableTypeCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTABLE TYPE CSV>\n")
		return csvErrors, internalErrFlag
	}

	// get created people csv and writer pointer
	peopleCSVFile, peopleCSVWriter, ok :=
		CreatePeopleCSV(
			TempCSVStore, currentTimeFormat,
			&RoomKeyFieldMap.PeopleCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <PEOPLE CSV>: %s\n", err.Error())
		return csvErrors, internalErrFlag
	}

	// ##############################
	// # PHASE 1 : SPLITTING DATA IN CSV FILES #
	// ##############################

	// ================================
	// First loop for validation on csv
	// ================================

	// To store the keys in slice in sorted order
	// always sort keys to iterate over csv rows in proper manner (from top to bottom)
	var csvRowDataMapKeys []int
	for k := range csvRowDataMap {
		csvRowDataMapKeys = append(csvRowDataMapKeys, k)
	}
	sort.Ints(csvRowDataMapKeys)

	// Iterating over cleaned csv data
	for _, rowIndex := range csvRowDataMapKeys {

		csvRow := *csvRowDataMap[rowIndex]

		// Write data to file of rentabletype
		WriteRentableTypeCSVData(
			&RentableTypeCSVRecordCount,
			rowIndex,
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

		guestdata := guestInfo[csvRow.Guest]
		if guestdata == nil {
			guestdata = &GuestCSVRow{GuestName: ""}
		}

		traceTCIDMap[rowIndex] = ""
		tracePeopleNote[rowIndex] = csvRow.Description

		peopleCollisions[csvRow.Guest]++
		if peopleCollisions[csvRow.Guest] > 1 {
			guestdata = &GuestCSVRow{GuestName: ""}
		}

		// Write data to file of people
		WritePeopleCSVData(
			&PeopleCSVRecordCount,
			rowIndex,
			tracePeopleCSVMap,
			peopleCSVWriter,
			&csvRow,
			&avoidDuplicatePeopleData,
			currentTimeFormat,
			userRRValues,
			&RoomKeyFieldMap.PeopleCSV,
			*guestdata,
			tracePeopleNote,
			traceDuplicatePeople,
			csvErrors,
		)
	}

	// Close all files as we are done here with writing data
	rentableTypeCSVFile.Close()
	peopleCSVFile.Close()

	// =======================
	// NESTED UTILITY FUNCTIONS
	// =======================

	// getTraceDataMap from string name
	getTraceDataMap := func(traceDataMapName string) map[int]int {
		switch traceDataMapName {
		case "traceRentableTypeCSVMap":
			return traceRentableTypeCSVMap
		case "tracePeopleCSVMap":
			return tracePeopleCSVMap
		case "traceRentableCSVMap":
			return traceRentableCSVMap
		case "traceRentalAgreementCSVMap":
			return traceRentalAgreementCSVMap
		default:
			return nil
		}
	}

	// getRoomKeyIndex used to get index and unit value from trace<TYPE>CSVMap map
	getRoomKeyIndex := func(traceDataMap map[int]int, index int) int {
		var roomKeyIndex int
		if roomKeyIndex, ok := traceDataMap[index]; ok {
			return roomKeyIndex
		}
		return roomKeyIndex
	}

	// rrDoLoad is a nested function
	// used to load data from csv with help of rcsv loaders
	rrDoLoad := func(fname string, handler func(string) []error, traceDataMapName string, dbType int) bool {
		Errs := handler(fname)

		for _, err := range Errs {
			// skip warnings about already existing records
			// if it's not kind of to skip then process it and count in error report
			errText := err.Error()

			if !csvRecordsToSkip(err) {
				lineNo, reason, ok := parseLineAndErrorFromRCSV(err, dbType)
				if !ok {
					// INTERNAL ERROR - RETURN FALSE
					return false
				}
				// get tracedatamap
				traceDataMap := getTraceDataMap(traceDataMapName)
				// now get the original row index of imported onesite csv and Unit value
				roomKeyIndex := getRoomKeyIndex(traceDataMap, lineNo)
				// generate new error
				csvErrors[roomKeyIndex] = append(csvErrors[roomKeyIndex], reason)
			} else {
				rlib.Ulog("DUPLICATE RECORD ERROR <%s>: %s\n", fname, errText)
			}
		}
		// return with success
		return true
	}

	// *****************************************************
	// rrPeopleDoLoad (SPECIAL METHOD TO LOAD PEOPLE)
	// *****************************************************
	rrPeopleDoLoad := func(fname string, handler func(string) []error, traceDataMapName string, dbType int) bool {
		Errs := handler(fname)

		for _, err := range Errs {
			// handling for duplicant transactant
			if strings.Contains(err.Error(), dupTransactantWithPrimaryEmail) {
				lineNo, _, ok := parseLineAndErrorFromRCSV(err, dbType)
				if !ok {
					// INTERNAL ERROR - RETURN FALSE
					return false
				}
				// get tracedatamap
				traceDataMap := getTraceDataMap(traceDataMapName)
				// now get the original row index of imported onesite csv and Unit value
				roomkeyIndex := getRoomKeyIndex(traceDataMap, lineNo)

				if csvRowDataMap[roomkeyIndex] == nil {
					continue
				}

				// load csvRow from dataMap troomkeyIndexo get email
				csvRow := *csvRowDataMap[roomkeyIndex]

				pEmail := ""
				if guestInfo[csvRow.Guest] != nil {
					pEmail = guestInfo[csvRow.Guest].Email
				}

				// get tcid from email
				t := rlib.GetTransactantByPhoneOrEmail(business.BID, pEmail)

				if t.TCID == 0 {
					// t = rlib.GetTransactantByName(business.BID, csvRow.Guest)
					reason := "E:<" + core.DBTypeMapStrings[core.DBPeople] + ">:Unable to get people information"
					csvErrors[roomkeyIndex] = append(csvErrors[roomkeyIndex], reason)
				} else {
					// if duplicate people found
					rlib.Ulog("DUPLICATE RECORD ERROR <%s>: %s", fname, err.Error())
					// map it in tcid map
					traceTCIDMap[roomkeyIndex] = tcidPrefix + strconv.FormatInt(t.TCID, 10)
				}
			} else if strings.Contains(err.Error(), dupTransactantWithCellPhone) {
				lineNo, _, ok := parseLineAndErrorFromRCSV(err, dbType)
				if !ok {
					// INTERNAL ERROR - RETURN FALSE
					return false
				}
				// get tracedatamap
				traceDataMap := getTraceDataMap(traceDataMapName)
				// now get the original row index of imported onesite csv and Unit value
				roomkeyIndex := getRoomKeyIndex(traceDataMap, lineNo)

				if csvRowDataMap[roomkeyIndex] == nil {
					continue
				}
				// load csvRow from dataMap to get email
				csvRow := *csvRowDataMap[roomkeyIndex]
				// pCellNo := csvRow.PhoneNumber
				pCellNo := ""
				if guestInfo[csvRow.Guest] != nil {
					pCellNo = guestInfo[csvRow.Guest].MainPhone
				}

				// get tcid from cellphonenumber
				t := rlib.GetTransactantByPhoneOrEmail(business.BID, pCellNo)
				if t.TCID == 0 {
					// unable to get TCID
					reason := "E:<" + core.DBTypeMapStrings[core.DBPeople] + ">:Unable to get people information"
					csvErrors[roomkeyIndex] = append(csvErrors[roomkeyIndex], reason)
				} else {
					// if duplicate people found
					rlib.Ulog("DUPLICATE RECORD ERROR <%s>: %s", fname, err.Error())
					// map it in tcid map
					traceTCIDMap[roomkeyIndex] = tcidPrefix + strconv.FormatInt(t.TCID, 10)
				}
			} else {
				lineNo, reason, ok := parseLineAndErrorFromRCSV(err, dbType)
				if !ok {
					// INTERNAL ERROR - RETURN FALSE
					return false
				}
				// get tracedatamap
				traceDataMap := getTraceDataMap(traceDataMapName)
				// now get the original row index of imported onesite csv and Unit value
				roomkeyIndex := getRoomKeyIndex(traceDataMap, lineNo)
				// generate new error
				csvErrors[roomkeyIndex] = append(csvErrors[roomkeyIndex], reason)
			}

			// *****************************************************

		}
		// return with success
		return true
	}

	// =========================================
	// LOAD RENTABLE TYPE CSV
	// =========================================
	var h = []csvLoadHandler{
		{
			Fname: rentableTypeCSVFile.Name(), Handler: rcsv.LoadRentableTypesCSV,
			TraceDataMap: "traceRentableTypeCSVMap", DBType: core.DBRentableType,
		},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			if !rrDoLoad(h[i].Fname, h[i].Handler, h[i].TraceDataMap, h[i].DBType) {
				// INTERNAL ERROR
				rlib.Ulog("INTERNAL ERROR <RENTABLE TYPE CSV>\n")
				return csvErrors, internalErrFlag
			}
		}
	}

	// =========================================
	// LOAD RENTABLE TYPE CSV
	// =========================================
	h = []csvLoadHandler{
		{
			Fname: peopleCSVFile.Name(), Handler: rcsv.LoadPeopleCSV,
			TraceDataMap: "tracePeopleCSVMap", DBType: core.DBPeople,
		},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			if !rrPeopleDoLoad(h[i].Fname, h[i].Handler, h[i].TraceDataMap, h[i].DBType) {
				// INTERNAL ERROR
				return csvErrors, internalErrFlag
			}
		}
	}

	// ========================================================
	// GET TCID FOR EACH ROW FROM PEOPLE CSV AND UPDATE TCID MAP
	// ========================================================

	for roomkeyIndex := range traceTCIDMap {
		// tcid := rlib.GetTCIDByNote(roomkeyNotesPrefix + strconv.Itoa(roomkeyIndex))
		tcid := rlib.GetTCIDByNote(tracePeopleNote[roomkeyIndex])
		// for duplicant case, it won't be found so need check here
		if tcid != 0 {
			traceTCIDMap[roomkeyIndex] = tcidPrefix + strconv.Itoa(tcid)
		}
	}

	// ==============================================================
	// AFTER POSSIBLE TCID FOUND, WRITE RENTABLE & RENTAL AGREEMENT CSV
	// ==============================================================

	// get created people csv and writer pointer
	rentableCSVFile, rentableCSVWriter, ok :=
		CreateRentableCSV(
			TempCSVStore, currentTimeFormat,
			&RoomKeyFieldMap.RentableCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTABLE CSV>: %s\n", err.Error())
		return csvErrors, internalErrFlag
	}

	// get created rental agreement csv and writer pointer
	rentalAgreementCSVFile, rentalAgreementCSVWriter, ok :=
		CreateRentalAgreementCSV(
			TempCSVStore, currentTimeFormat,
			&RoomKeyFieldMap.RentalAgreementCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTAL AGREEMENT CSV>: %s\n", err.Error())
		return csvErrors, internalErrFlag
	}

	// iteration over csv row data structure and write data to csv
	for _, rowIndex := range csvRowDataMapKeys {

		// load csvRow from dataMap
		csvRow := *csvRowDataMap[rowIndex]

		// Write data to file of rentable
		WriteRentableData(
			&RentableCSVRecordCount,
			rowIndex,
			traceRentableCSVMap,
			rentableCSVWriter,
			&csvRow,
			currentTime,
			currentTimeFormat,
			userRRValues,
			&RoomKeyFieldMap.RentableCSV,
			traceTCIDMap,
			csvErrors,
		)

		// Write data to file of rentalAgreement
		WriteRentalAgreementData(
			&RentalAgreementCSVRecordCount,
			rowIndex,
			traceRentalAgreementCSVMap,
			rentalAgreementCSVWriter,
			&csvRow,
			currentTime,
			currentTimeFormat,
			userRRValues,
			&RoomKeyFieldMap.RentalAgreementCSV,
			traceTCIDMap,
			csvErrors,
		)
	}

	// closing files
	rentableCSVFile.Close()
	rentalAgreementCSVFile.Close()

	// =====================================
	// LOAD RENTABLE & RENTAL AGREEMENT CSV
	// =====================================
	h = []csvLoadHandler{
		{Fname: rentableCSVFile.Name(), Handler: rcsv.LoadRentablesCSV, TraceDataMap: "traceRentableCSVMap", DBType: core.DBRentable},
		{Fname: rentalAgreementCSVFile.Name(), Handler: rcsv.LoadRentalAgreementCSV, TraceDataMap: "traceRentalAgreementCSVMap", DBType: core.DBRentalAgreement},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			if !rrDoLoad(h[i].Fname, h[i].Handler, h[i].TraceDataMap, h[i].DBType) {
				// INTERNAL ERROR
				return csvErrors, internalErrFlag
			}
		}
	}

	// ============================
	// CLEAR THE TEMPORARY CSV FILES
	// ============================
	// testmode is not enabled then only remove temp files
	if testMode != 1 {
		clearSplittedTempCSVFiles(currentTimeFormat)
	}

	// ===============================
	// EVALUATE SUMMARY REPORT COUNT
	// ===============================

	// count possible values
	summaryReport[core.DBRentable]["possible"] = RentableCSVRecordCount
	summaryReport[core.DBRentalAgreement]["possible"] = RentalAgreementCSVRecordCount
	summaryReport[core.DBRentableType]["possible"] = RentableTypeCSVRecordCount
	summaryReport[core.DBPeople]["possible"] = PeopleCSVRecordCount

	internalErrFlag = false
	// RETURN
	return csvErrors, internalErrFlag
}

func loadGuestInfoCSV(
	guestInfoCSV string,
) (map[string]*GuestCSVRow, error) {

	// store all guest info in guestInfoMap
	guestInfoMap := map[string]*GuestCSVRow{}

	csvHeadersIndex := getGuestCSVHeadersIndexMap()

	skipRowsCount := 0

	// load csv file and get data from csv
	t := rlib.LoadCSV(guestInfoCSV)

	// detect how many rows we need to skip first
	for rowIndex := 0; rowIndex < len(t); rowIndex++ {
		for colIndex := 0; colIndex < len(t[rowIndex]); colIndex++ {
			// remove all white spaces and make lower case
			cellTextValue := strings.ToLower(
				core.SpecialCharsReplacer.Replace(t[rowIndex][colIndex]))
			// if header is exist in map then overwrite it position
			if field, ok := guestCSVColumnFieldMap[cellTextValue]; ok {
				csvHeadersIndex[field] = colIndex
			}
		}
		// check after row columns parsing that headers are found or not
		headersFound := true
		for _, v := range csvHeadersIndex {
			if v == -1 {
				headersFound = false
				break
			}
		}

		if headersFound {
			// update rowIndex by 1 because we're going to break here
			rowIndex++
			skipRowsCount = rowIndex
			break
		}
	}

	// if skipRowsCount is still 0 that means data could not be parsed from csv
	if skipRowsCount == 0 {
		missingHeaders := []string{}
		// make message of missing columns
		for missedH, v := range csvHeadersIndex {
			if v == -1 {
				missingHeaders = append(missingHeaders, missedH)
			}
		}

		headerError := "(Guest Data CSV) Required data column(s) missing: "
		headerError += strings.Join(missingHeaders, ", ")

		err := errors.New(headerError)
		return guestInfoMap, err
	}

	// if skipRowsCount found get next row and proceed on rest of the rows with loop
	for rowIndex := skipRowsCount + 1; rowIndex <= len(t); rowIndex++ {

		// if column order has been validated then only perform
		// data validation on value, type
		rowLoaded, csvRow := loadGuestInfoCSVRow(csvHeadersIndex, t[rowIndex-1])

		// **************************************************************
		// NOTE: might need to change logic, if t[i] contains blank data that
		// we should stop the loop as we have to skip rest of the rows
		// (please look at guest info csv)
		// **************************************************************
		if !rowLoaded {

			// what IF, only headers are there
			if rowIndex == skipRowsCount {
				err := errors.New("There are no data rows present")
				return guestInfoMap, err
			}

			// else break the loop as there are no more data
			break
		}
		guestInfoMap[csvRow.GuestName] = &csvRow
	}

	return guestInfoMap, nil
}

// rollBackImportOperation func used to clear out the things
// that created by program temporarily while loading onesite data
//  and if any error occurs
func rollBackImportOperation(timestamp string) {
	clearSplittedTempCSVFiles(timestamp)
}

// clearSplittedTempCSVFiles func used only to clear
// temporarily csv files created by program
func clearSplittedTempCSVFiles(timestamp string) {
	for _, filePrefix := range prefixCSVFile {
		fileName := filePrefix + timestamp + ".csv"
		filePath := path.Join(TempCSVStore, fileName)
		os.Remove(filePath)
	}
}

// CSVHandler is main function to handle user uploaded
// csv and extract information
func CSVHandler(
	csvPath string,
	GuestInfoCSV string,
	testMode int,
	userRRValues map[string]string,
	business *rlib.Business,
	debugMode int,
) (string, bool, bool) {

	// init values
	csvLoaded := true

	// report text
	csvReport := ""

	// get current timestamp used for creating csv files unique way
	currentTime := time.Now()

	// RFC3339Nano is const format defined in time package
	// <FORMAT> = <SAMPLE>
	// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	// it is helpful while creating unique files
	currentTimeFormat := currentTime.Format(time.RFC3339Nano)

	// summaryReportCount contains each type csv as a key
	// with count of total imported, possible, issues in csv data
	summaryReportCount := map[int]map[string]int{
		core.DBRentableType:    map[string]int{"imported": 0, "possible": 0, "issues": 0},
		core.DBPeople:          map[string]int{"imported": 0, "possible": 0, "issues": 0},
		core.DBRentable:        map[string]int{"imported": 0, "possible": 0, "issues": 0},
		core.DBRentalAgreement: map[string]int{"imported": 0, "possible": 0, "issues": 0},
	}

	// --------------------------------------------------------------------------------------------------------- //

	// ---------------------- call onesite loader ----------------------------------------
	guestInfo, guestCSVError := loadGuestInfoCSV(GuestInfoCSV)
	if guestCSVError != nil {
		// throw an error in report
	}

	// ---------------------- call roomkey loader ----------------------------------------
	csvErrs, internalErr := loadRoomKeyCSV(csvPath, guestInfo, testMode, userRRValues,
		business, currentTime, currentTimeFormat,
		summaryReportCount)

	// if internal error then just return from here, nothing to do
	if internalErr {
		return csvReport, internalErr, csvLoaded
	}

	// check if there any errors from onesite loader
	if len(csvErrs) > 0 {
		csvReport, csvLoaded = errorReporting(business, csvErrs, summaryReportCount, csvPath, GuestInfoCSV, debugMode, currentTime)

		// if not testmode then only do rollback
		if testMode != 1 {
			rollBackImportOperation(currentTimeFormat)
		}

		return csvReport, internalErr, csvLoaded
	}

	// ===== 4. Geneate Report =====
	csvReport = successReport(business, summaryReportCount, csvPath, GuestInfoCSV, debugMode, currentTime)

	// ===== 5. Return =====
	return csvReport, internalErr, csvLoaded

}
