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

// getOneSiteMapping reads json file and loads
// field mapping structure in go for further usage
func getOneSiteMapping(oneSiteFieldMap *CSVFieldMap) error {

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
	err = json.Unmarshal(fieldmap, oneSiteFieldMap)
	if err != nil {
		return err
	}

	return err
}

// loadOneSiteCSV loads the values from the supplied csv file and
// creates rlib.Business records as needed
func loadOneSiteCSV(
	oneSiteCSV string,
	testMode int,
	userRRValues map[string]string,
	business *rlib.Business,
	currentTime time.Time,
	currentTimeFormat string,
	summaryReport map[int]map[string]int,
) (map[int]string, map[int][]string, bool) {

	// returns unitmap, csvError list, internal error, csv loaded?

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
	funcname := "loadOneSiteCSV"

	// this count used to skip number of rows from the very top of csv
	var skipRowsCount int
	var rowIndex int

	// customAttributesRefData holds the data for future operation to insert
	// custom attribute ref in system for each rentableType
	// so we identify each element in this list via Style value
	customAttributesRefData := map[string]CARD{}

	// this map is used to hold csvRow typed struct after data has been loaded in it
	// by iterating over csv data, re-usable for rentable, rental agreement data
	csvRowDataMap := map[int]*CSVRow{}

	// --------------------------- trace data map ---------------------------- //
	// trace<TYPE>CSVMap used to hold records
	// by which we can traceout which records has been writtern to csv
	// with key of row index of <TARGET_TYPE> CSV, value of original's imported csv rowNumber
	traceRentableTypeCSVMap := map[int]int{}
	traceRentableCSVMap := map[int]int{}
	tracePeopleCSVMap := map[int]int{}
	traceRentalAgreementCSVMap := map[int]int{}
	traceCustomAttributeCSVMap := map[int]int{}

	// traceTCIDMap hold TCID for each people to be loaded via people csv
	// with reference of original onesite csv
	traceTCIDMap := map[int]string{}

	// traceUnitMap holds records by which we can trace the unit with row index of csv
	// Unit would be unique in onesite imported csv
	// key: rowIndex of onesite csv, value: Unit value of each row of onesite csv
	traceUnitMap := map[int]string{}

	// traceDuplicatePeople holds records with unique string (name, email, phone)
	// with duplicant match at row
	// e.g.; {
	// 	"email": {"test@no.com": [1,5,3]},
	// 	"phone": {"9999999999": [2,4]},
	// 	"name": {"foo, bar": 3},
	// }
	traceDuplicatePeople := map[string][]string{
		"name": {}, "email": {}, "phone": {},
	}

	// --------------------- avoid duplicate data structures -------------------- //
	// avoidDuplicateRentableTypeData used to keep track of rentableTypeData with Style field
	// so that duplicate entries can be avoided while creating rentableType csv file
	avoidDuplicateRentableTypeData := []string{}

	// avoidDuplicateCustomAttributeData is tricky map which holds the
	// duplicate data in slice for each field defined in customAttributeMap
	avoidDuplicateCustomAttributeData := map[string][]string{}
	for k := range customAttributeMap {
		avoidDuplicateCustomAttributeData[k] = []string{}
	}

	// --------------------------- csv record count ----------------------------
	// <TYPE>CSVRecordCount used to hold records count inserted in csv
	// initialize with 1 because first row contains headers in target generated csv
	// these are POSSIBLE record count that going to be imported
	RentableTypeCSVRecordCount := 0
	RentableCSVRecordCount := 0
	PeopleCSVRecordCount := 0
	RentalAgreementCSVRecordCount := 0
	CustomAttributeCSVRecordCount := 0

	// ================================================
	// LOAD FIELD MAP AND GET HEADERS, LENGTH OF HEADERS
	// ================================================

	// csvCols and consts for all onesite csv fields are defined in
	// constant.go file
	// Onesite csv headers slice and load it from csvCols
	oneSiteCSVHeaders := []string{}
	for _, header := range csvCols {
		oneSiteCSVHeaders = append(oneSiteCSVHeaders, header.Name)
	}
	oneSiteColumnLength := len(oneSiteCSVHeaders)

	// load onesite mapping
	var oneSiteFieldMap CSVFieldMap
	err := getOneSiteMapping(&oneSiteFieldMap)
	if err != nil {
		rlib.Ulog("INTERNAL ERROR <ONESITE FIELD MAPPING>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// ==============================
	// COUNT ROWS NEEDS TO BE SKIPPED
	// ==============================

	// load csv file and get data from csv
	t := rlib.LoadCSV(oneSiteCSV)

	// detect how many rows we need to skip first
	for rowIndex = 1; rowIndex <= len(t); rowIndex++ {

		if skipRowsCount == 0 {
			csvRowDataString := strings.Replace(
				strings.Join(t[rowIndex-1][:oneSiteColumnLength], ","),
				" ", "", -1)

			csvHeaderString := strings.Replace(
				strings.Join(oneSiteCSVHeaders[:oneSiteColumnLength], ","),
				" ", "", -1)

			if csvRowDataString == csvHeaderString {
				skipRowsCount = rowIndex
				break
			}
		}
	}

	// if skipRowsCount is still 0 that means data could not be parsed from csv
	if skipRowsCount == 0 {
		// ******** special entry ***********
		// -1 means there is no data
		internalErrFlag = false
		csvErrors[-1] = append(csvErrors[-1], "There are no data in onesite csv to load")
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// ========================================================
	// WRITE DATA FOR CUSTOM ATTRIBUTE, RENTABLE TYPE, PEOPLE CSV
	// ========================================================

	// get created customAttibutes csv and writer pointer
	customAttributeCSVFile, customAttributeCSVWriter, ok :=
		CreateCustomAttibutesCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.CustomAttributeCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <CUSTOM ATTRIUTE CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// ----------------------- create files and get csv writer object -----------------------
	// get created rentabletype csv and writer pointer
	rentableTypeCSVFile, rentableTypeCSVWriter, ok :=
		CreateRentableTypeCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.RentableTypeCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTABLE TYPE CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// get created people csv and writer pointer
	peopleCSVFile, peopleCSVWriter, ok :=
		CreatePeopleCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.PeopleCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <PEOPLE CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// if skipRowsCount found get next row and proceed on rest of the rows with loop
	for rowIndex = skipRowsCount + 1; rowIndex <= len(t); rowIndex++ {

		// csv Columns order validation
		x, err := rcsv.ValidateCSVColumns(csvCols, t[rowIndex-1][:oneSiteColumnLength], funcname, rowIndex)

		if x > 0 {
			// there is no db type specific error so pass it as -1
			_, reason, ok := parseLineAndErrorFromRCSV(err, -1)
			if !ok {
				// INTERNAL ERROR
				return traceUnitMap, csvErrors, internalErrFlag
			}
			csvErrors[rowIndex] = append(csvErrors[rowIndex], reason)
		} else {
			// if column order has been validated then only perform
			// data validation on value, type
			rowLoaded, csvRow := loadOneSiteCSVRow(csvCols, t[rowIndex-1][:oneSiteColumnLength])

			// **************************************************************
			// NOTE: might need to change logic, if t[i] contains blank data that
			// we should stop the loop as we have to skip rest of the rows
			// (please look at onesite csv)
			// **************************************************************
			if !rowLoaded {

				// what IF, only headers are there
				if rowIndex == skipRowsCount {
					// ******** special entry ***********
					// -1 means there is no data
					internalErrFlag = false
					csvErrors[-1] = append(csvErrors[-1], "There are no data in onesite csv to load")
					return traceUnitMap, csvErrors, internalErrFlag
				}

				// else break the loop as there are no more data
				break
			}

			// rowLoaded successfully then do rest of the operation

			// get rentable status from csv data
			csvRentableStatus := csvRow.UnitLeaseStatus

			// get unit from csv data
			csvUnit := csvRow.Unit

			// for rentable status exists in csvRow, get set of csv types which can be allowed
			// to perform write data for csv
			// need to call validation function as in to get the values
			_, rrStatus, _ := IsValidRentableStatus(csvRentableStatus)
			csvTypesSet := canWriteCSVStatusMap[rrStatus]
			var canWriteData bool

			// mark Unit value with row index value
			// even if it is blank
			traceUnitMap[rowIndex] = csvUnit

			// keep csv record in this
			csvRowDataMap[rowIndex] = &csvRow

			// check first that for this row's status rentableType data can be written
			canWriteData = core.IntegerInSlice(core.RENTABLETYPECSV, csvTypesSet)
			if canWriteData {
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
					&oneSiteFieldMap.RentableTypeCSV,
					customAttributesRefData,
					business,
				)
			}

			// check first that for this row's status custom attributes data can be written
			canWriteData = core.IntegerInSlice(core.CUSTOMATTRIUTESCSV, csvTypesSet)
			if canWriteData {
				// Write data to file of CustomAttribute
				WriteCustomAttributeData(
					&CustomAttributeCSVRecordCount,
					rowIndex,
					traceCustomAttributeCSVMap,
					customAttributeCSVWriter,
					&csvRow,
					avoidDuplicateCustomAttributeData,
					currentTimeFormat,
					userRRValues,
					&oneSiteFieldMap.CustomAttributeCSV,
				)
			}

			// check first that for this row's status people data can be written
			canWriteData = core.IntegerInSlice(core.PEOPLECSV, csvTypesSet)
			if canWriteData {

				// if people data can be writable then init TCIDMap index
				// with blank string value
				traceTCIDMap[rowIndex] = ""

				// Write data to file of people
				WritePeopleCSVData(
					&PeopleCSVRecordCount,
					rowIndex,
					tracePeopleCSVMap,
					peopleCSVWriter,
					&csvRow,
					traceDuplicatePeople,
					currentTimeFormat,
					userRRValues,
					&oneSiteFieldMap.PeopleCSV,
					csvErrors,
				)
			}
		}
	}

	// Close all files as we are done here with writing data
	rentableTypeCSVFile.Close()
	peopleCSVFile.Close()
	customAttributeCSVFile.Close()

	// =======================
	// NESTED UTILITY FUNCTIONS
	// =======================

	// getTraceDataMap from string name
	getTraceDataMap := func(traceDataMapName string) map[int]int {
		switch traceDataMapName {
		case "traceCustomAttributeCSVMap":
			return traceCustomAttributeCSVMap
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

	// getIndexAndUnit used to get index and unit value from trace<TYPE>CSVMap map
	getIndexAndUnit := func(traceDataMap map[int]int, index int) (int, string) {
		var onesiteIndex int
		var unit string
		if onesiteIndex, ok := traceDataMap[index]; ok {
			if unit, ok := traceUnitMap[onesiteIndex]; ok {
				return onesiteIndex, unit
			}
			return onesiteIndex, unit
		}
		return onesiteIndex, unit
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
				onesiteIndex, _ := getIndexAndUnit(traceDataMap, lineNo)
				// generate new error
				csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
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
		// fmt.Print(rcsv.ErrlistToString(&Errs))

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
				onesiteIndex, _ := getIndexAndUnit(traceDataMap, lineNo)
				// load csvRow from dataMap to get email
				csvRow := *csvRowDataMap[onesiteIndex]
				pEmail := csvRow.Email
				// get tcid from email
				t := rlib.GetTransactantByPhoneOrEmail(business.BID, pEmail)
				if t.TCID == 0 {
					// unable to get TCID
					reason := "E:Unable to get people information"
					csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
				} else {
					// if duplicate people found
					rlib.Ulog("DUPLICATE RECORD ERROR <%s>: %s", fname, err.Error())
					// map it in tcid map
					traceTCIDMap[onesiteIndex] = tcidPrefix + strconv.FormatInt(t.TCID, 10)
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
				onesiteIndex, _ := getIndexAndUnit(traceDataMap, lineNo)
				// generate new error
				csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
			}

			// *****************************************************************
			// AS WE DON'T HAVE MAPPING OF PHONENUMBER TO CELLPHONE
			// WE JUST AVOID THIS CHECK, BUT KEEP THIS IN CASE MAPPING
			// OF PHONENUMBER CHANGED TO CELLPHONE
			// PLACE IT AFTER DUPLICATE EMAIL CHECK
			// *****************************************************************
			/*
				else if strings.Contains(errText, dupTransactantWithCellPhone) {
					lineNo, _, ok := parseLineAndErrorFromRCSV(err, dbType)
					if !ok {
						// INTERNAL ERROR - RETURN FALSE
						return false
					}
					// get tracedatamap
					traceDataMap := getTraceDataMap(traceDataMapName)
					// now get the original row index of imported onesite csv and Unit value
					onesiteIndex, unit := getIndexAndUnit(traceDataMap, lineNo)
					// load csvRow from dataMap to get email
					csvRow := *csvRowDataMap[onesiteIndex]
					pCellNo := csvRow.PhoneNumber
					// get tcid from cellphonenumber
					t := rlib.GetTransactantByPhoneOrEmail(business.BID, pCellNo)
					if t.TCID == 0 {
						// unable to get TCID
						reason := "E:Unable to get people information"
						csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
					} else {
						// if duplicate people found
						rlib.Ulog("DUPLICATE RECORD ERROR <%s>: %s", fname, err.Error())
						// map it in tcid map
						traceTCIDMap[onesiteIndex] = tcidPrefix + strconv.FormatInt(t.TCID, 10)
					}
				}
			*/
			// *****************************************************

		}
		// return with success
		return true
	}

	// =========================================
	// LOAD CUSTOM ATTRIBUTE & RENTABLE TYPE CSV
	// =========================================
	var h = []csvLoadHandler{
		{Fname: customAttributeCSVFile.Name(), Handler: rcsv.LoadCustomAttributesCSV, TraceDataMap: "traceCustomAttributeCSVMap", DBType: core.DBCustomAttr},
		{Fname: rentableTypeCSVFile.Name(), Handler: rcsv.LoadRentableTypesCSV, TraceDataMap: "traceRentableTypeCSVMap", DBType: core.DBRentableType},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			if !rrDoLoad(h[i].Fname, h[i].Handler, h[i].TraceDataMap, h[i].DBType) {
				// INTERNAL ERROR
				return traceUnitMap, csvErrors, internalErrFlag
			}
		}
	}

	// =====================================
	// INSERT CUSTOM ATTRIBUTE REF MANUALLY
	// AFTER CUSTOM ATTRIB AND RENTABLE TYPE
	// LOADED SUCCESSFULLY
	// =====================================

	// always sort keys
	var customAttributesRefDataKeys []string
	for k := range customAttributesRefData {
		customAttributesRefDataKeys = append(customAttributesRefDataKeys, k)
	}
	sort.Strings(customAttributesRefDataKeys)

	for _, key := range customAttributesRefDataKeys {
		errPrefix := "E:<" + core.DBTypeMapStrings[core.DBCustomAttrRef] + ">:"
		// find rentableType
		refData := customAttributesRefData[key]
		rt, err := rlib.GetRentableTypeByStyle(refData.Style, refData.BID)
		if err != nil {
			rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", err.Error())
			csvErrors[refData.RowIndex] = append(csvErrors[refData.RowIndex], errPrefix+"Unable to insert custom attribute")
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
				rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", "CUSTOM ATTRIBUTE NOT FOUND IN DB")
				csvErrors[refData.RowIndex] = append(csvErrors[refData.RowIndex], errPrefix+"Unable to insert custom attribute")
				continue
			}

			// insert custom attribute ref in system
			var a rlib.CustomAttributeRef
			a.ElementType = rlib.ELEMRENTABLETYPE
			a.BID = business.BID
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
				rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", errText)
				continue
			}

			err := rlib.InsertCustomAttributeRef(&a)
			if err != nil {
				rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", err.Error())
				csvErrors[refData.RowIndex] = append(csvErrors[refData.RowIndex], errPrefix+"Unable to insert custom attribute")
				continue
			}
		}
	}

	// ================
	// LOAD PEOPLE CSV
	// ================
	h = []csvLoadHandler{
		{Fname: peopleCSVFile.Name(), Handler: rcsv.LoadPeopleCSV, TraceDataMap: "tracePeopleCSVMap", DBType: core.DBPeople},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			if !rrPeopleDoLoad(h[i].Fname, h[i].Handler, h[i].TraceDataMap, h[i].DBType) {
				// INTERNAL ERROR
				return traceUnitMap, csvErrors, internalErrFlag
			}
		}
	}

	// ========================================================
	// GET TCID FOR EACH ROW FROM PEOPLE CSV AND UPDATE TCID MAP
	// ========================================================

	for onesiteIndex := range traceTCIDMap {
		tcid := rlib.GetTCIDByNote(onesiteNotesPrefix + strconv.Itoa(onesiteIndex))
		// for duplicant case, it won't be found so need check here
		if tcid != 0 {
			traceTCIDMap[onesiteIndex] = tcidPrefix + strconv.Itoa(tcid)
		}
	}

	// ==============================================================
	// AFTER POSSIBLE TCID FOUND, WRITE RENTABLE & RENTAL AGREEMENT CSV
	// ==============================================================

	// get created people csv and writer pointer
	rentableCSVFile, rentableCSVWriter, ok :=
		CreateRentableCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.RentableCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTABLE CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// get created rental agreement csv and writer pointer
	rentalAgreementCSVFile, rentalAgreementCSVWriter, ok :=
		CreateRentalAgreementCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.RentalAgreementCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTAL AGREEMENT CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// always sort keys to iterate over csv rows in proper manner (from top to bottom)
	var csvRowDataMapKeys []int
	for k := range csvRowDataMap {
		csvRowDataMapKeys = append(csvRowDataMapKeys, k)
	}
	sort.Ints(csvRowDataMapKeys)

	// iteration over csv row data structure and write data to csv
	for _, rowIndex := range csvRowDataMapKeys {

		// load csvRow from dataMap
		csvRow := *csvRowDataMap[rowIndex]

		// for rentable status exists in csvRow, get set of csv types which can be allowed
		// to perform write data for csv
		// need to call validation function as in get values
		_, rrStatus, _ := IsValidRentableStatus(csvRow.UnitLeaseStatus)
		csvTypesSet := canWriteCSVStatusMap[rrStatus]
		var canWriteData bool

		// check first that for this row's status rentable data can be written
		canWriteData = core.IntegerInSlice(core.RENTABLECSV, csvTypesSet)
		if canWriteData {
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
				&oneSiteFieldMap.RentableCSV,
				traceTCIDMap,
				csvErrors,
			)
		}

		// check first that for this row's status rental agreement data can be written
		canWriteData = core.IntegerInSlice(core.RENTALAGREEMENTCSV, csvTypesSet)
		if canWriteData {
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
				&oneSiteFieldMap.RentalAgreementCSV,
				traceTCIDMap,
				csvErrors,
			)
		}
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
				return traceUnitMap, csvErrors, internalErrFlag
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
	summaryReport[core.DBCustomAttr]["possible"] = CustomAttributeCSVRecordCount
	summaryReport[core.DBCustomAttrRef]["possible"] = CustomAttributeCSVRecordCount // customAttrRef count same as customAttr
	summaryReport[core.DBPeople]["possible"] = PeopleCSVRecordCount

	// if not internal errors then hit db to count total imported rows
	// countImportedRow(summaryReport, business.BID)

	// =======
	// RETURN
	// =======
	// no internal error so make it false
	internalErrFlag = false
	return traceUnitMap, csvErrors, internalErrFlag
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
	testMode int,
	userRRValues map[string]string,
	business *rlib.Business,
	debugMode int,
) (string, bool, bool) {

	// return report, internal error flag, done (csv loaded or not)

	// csv loaded successfully flag
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
		core.DBCustomAttr:      map[string]int{"imported": 0, "possible": 0, "issues": 0},
		core.DBRentableType:    map[string]int{"imported": 0, "possible": 0, "issues": 0},
		core.DBCustomAttrRef:   map[string]int{"imported": 0, "possible": 0, "issues": 0},
		core.DBPeople:          map[string]int{"imported": 0, "possible": 0, "issues": 0},
		core.DBRentable:        map[string]int{"imported": 0, "possible": 0, "issues": 0},
		core.DBRentalAgreement: map[string]int{"imported": 0, "possible": 0, "issues": 0},
	}

	// ====== Call onesite loader =====
	unitMap, csvErrs, internalErr := loadOneSiteCSV(
		csvPath, testMode, userRRValues,
		business, currentTime, currentTimeFormat,
		summaryReportCount)

	// if internal error then just return from here, nothing to do
	if internalErr {
		return csvReport, internalErr, csvLoaded
	}

	// check if there any errors from onesite loader
	if len(csvErrs) > 0 {
		csvReport, csvLoaded = errorReporting(business, csvErrs, unitMap, summaryReportCount, csvPath, debugMode)

		// if not testmode then only do rollback
		if testMode != 1 {
			rollBackImportOperation(currentTimeFormat)
		}

		return csvReport, internalErr, csvLoaded
	}

	// ===== 4. Geneate Report =====
	csvReport = successReport(business, summaryReportCount, csvPath, debugMode)

	// ===== 5. Return =====
	return csvReport, internalErr, csvLoaded
}

// generateSummaryReport used to generate summary report from argued struct
func generateSummaryReport(summaryCount map[int]map[string]int) string {
	var report string

	tableTitle := "SUMMARY SECTION"
	report += strings.Repeat("=", len(tableTitle))
	report += "\n" + tableTitle + "\n"
	report += strings.Repeat("=", len(tableTitle))
	report += "\n"

	var tbl rlib.Table
	tbl.Init()
	tbl.AddColumn("Data Type", 30, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Total Possible", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Total Imported", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Issues", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	summaryCountIndexes := []int{}
	for index := range core.DBTypeMap {
		summaryCountIndexes = append(summaryCountIndexes, index)
	}
	sort.Ints(summaryCountIndexes)

	for _, dbType := range summaryCountIndexes {

		// get each db type map
		countMap := summaryCount[dbType]

		// add row
		tbl.AddRow()
		tbl.Puts(-1, 0, core.DBTypeMap[dbType])
		tbl.Puts(-1, 1, strconv.Itoa(countMap["possible"]))
		tbl.Puts(-1, 2, strconv.Itoa(countMap["imported"]))
		tbl.Puts(-1, 3, strconv.Itoa(countMap["issues"]))
	}

	report += tbl.SprintTable(rlib.RPTTEXT)
	report += "\n\n"

	return report
}

// generateDetailedReport gives detailed report with (rowNumber, unit, db type, reason)
func generateDetailedReport(
	csvErrors map[int][]string,
	unitMap map[int]string,
	summaryCount map[int]map[string]int,
) (string, bool) {

	// return detailed report, tell program should it generate csv report?
	// in case of no errors, but has some warnings then csv report needs to be generated

	csvReportGenerate := true

	var detailedReport string
	tableTitle := "DETAILED REPORT SECTION"
	detailedReport += strings.Repeat("=", len(tableTitle))
	detailedReport += "\n" + tableTitle + "\n"
	detailedReport += strings.Repeat("=", len(tableTitle))
	detailedReport += "\n"

	var tbl rlib.Table
	tbl.Init()
	tbl.AddColumn("Input Line", 10, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Unit Name", 20, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("RentRoll DB Type", 20, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	tbl.AddColumn("Description", 180, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)

	csvErrorIndexes := []int{}
	for rowIndex := range csvErrors {
		csvErrorIndexes = append(csvErrorIndexes, rowIndex)
	}
	sort.Ints(csvErrorIndexes)

	// to count imported, just remove this count from possible column in summaryCount
	errorCount := map[int]int{
		core.DBCustomAttr:      0,
		core.DBRentableType:    0,
		core.DBCustomAttrRef:   0,
		core.DBPeople:          0,
		core.DBRentable:        0,
		core.DBRentalAgreement: 0,
	}

	for _, rowIndex := range csvErrorIndexes {

		// get error from index
		reportError := csvErrors[rowIndex]

		// check that rowIndex is -1
		// -1 means no data found in csv
		if rowIndex == -1 {
			tbl.AddRow()
			tbl.Puts(-1, 0, "")
			tbl.Puts(-1, 1, "")
			tbl.Puts(-1, 2, "")
			tbl.Puts(-1, 3, reportError[0])

			// append detailed section
			detailedReport += tbl.SprintTable(rlib.RPTTEXT)

			// return
			csvReportGenerate = false
			return detailedReport, csvReportGenerate
		}

		// get unit from map
		unit, _ := unitMap[rowIndex]

		// used to separate errors, warnings
		rowErrors, rowWarnings := []string{}, []string{}

		for _, reason := range reportError {
			if strings.HasPrefix(reason, "E:") {

				// if any error captured then do not generate csv report
				csvReportGenerate = false

				// red color
				reason = strings.Replace(reason, "E:", "", -1)
				rowErrors = append(rowErrors, reason)
			}
			if strings.HasPrefix(reason, "W:") {
				// orange color
				reason = strings.Replace(reason, "W:", "", -1)
				rowWarnings = append(rowWarnings, reason)
			}
		}

		// first put errors
		for _, errorText := range rowErrors {
			errorText := strings.Split(errorText, ">:")
			dbType, reason := errorText[0], errorText[1]
			dbType = strings.Replace(dbType, "<", "", -1)
			dbTypeInt, _ := strconv.Atoi(dbType)

			// count issues in summary report
			summaryCount[dbTypeInt]["issues"]++

			// error count, helpful to count imported
			errorCount[dbTypeInt]++

			// put in tabl
			tbl.AddRow()
			tbl.Puts(-1, 0, strconv.Itoa(rowIndex))
			tbl.Puts(-1, 1, unit)
			tbl.Puts(-1, 2, core.DBTypeMap[dbTypeInt])
			tbl.Puts(-1, 3, reason)
		}

		// then warnings
		for _, warningText := range rowWarnings {
			warningText := strings.Split(warningText, ">:")
			dbType, reason := warningText[0], warningText[1]
			dbType = strings.Replace(dbType, "<", "", -1)
			dbTypeInt, _ := strconv.Atoi(dbType)

			// prefixed with "Warning: "
			reason = "Warning: " + reason

			// count issues in summary report
			summaryCount[dbTypeInt]["issues"]++

			tbl.AddRow()
			tbl.Puts(-1, 0, strconv.Itoa(rowIndex))
			tbl.Puts(-1, 1, unit)
			tbl.Puts(-1, 2, core.DBTypeMap[dbTypeInt])
			tbl.Puts(-1, 3, reason)
		}
	}

	// count imported
	for dbTypeInt := range summaryCount {
		summaryCount[dbTypeInt]["imported"] = summaryCount[dbTypeInt]["possible"] - errorCount[dbTypeInt]
	}

	// append detailed section
	detailedReport += tbl.SprintTable(rlib.RPTTEXT)
	detailedReport += "\n\n"

	// return
	return detailedReport, csvReportGenerate
}

// generateCSVReport return report for all type of csv defined here
func generateCSVReport(
	business *rlib.Business,
	summaryCount map[int]map[string]int,
	csvFile string,
) string {

	var r = []rrpt.ReporterInfo{
		{ReportNo: 5, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportRentableTypes, Bid: business.BID},
		{ReportNo: 6, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportRentables, Bid: business.BID},
		{ReportNo: 7, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportPeople, Bid: business.BID},
		{ReportNo: 9, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportRentalAgreements, Bid: business.BID},
		{ReportNo: 14, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportCustomAttributes, Bid: business.BID},
		{ReportNo: 15, OutputFormat: rlib.RPTTEXT, Handler: rcsv.RRreportCustomAttributeRefs, Bid: business.BID},
	}

	var report string

	title := fmt.Sprintf("RECORDS FOR BUSINESS UNIT DESIGNATION: %s", business.Name)
	report += strings.Repeat("=", len(title))
	report += "\n" + title + "\n"
	report += strings.Repeat("=", len(title))
	report += "\n"

	for i := 0; i < len(r); i++ {
		report += r[i].Handler(&r[i])
		report += strings.Repeat("=", 80)
		report += "\n"
	}

	return report
}

// successReport generates success report
func successReport(
	business *rlib.Business,
	summaryCount map[int]map[string]int,
	csvFile string,
	debugMode int,
) string {
	var report string

	// import file name in first line
	report += fmt.Sprintf("Import File: %s", csvFile)
	report += "\n\n"

	// **************************************************************
	// FOR SUCCESSFUL CASE, there are no errors and also no warnings
	// so imported is same as possible values
	// **************************************************************
	for dbType := range summaryCount {
		summaryCount[dbType]["imported"] = summaryCount[dbType]["possible"]
	}

	// append summary report
	report += generateSummaryReport(summaryCount)

	// csv report for all types if testmode is on
	if debugMode == 1 {
		report += generateCSVReport(business, summaryCount, csvFile)
	}

	// return
	return report
}

// errorReporting used to report the errors for onesite csv
func errorReporting(
	business *rlib.Business,
	csvErrors map[int][]string,
	unitMap map[int]string,
	summaryCount map[int]map[string]int,
	csvFile string,
	debugMode int,
) (string, bool) {
	var errReport string

	// import file name in first line
	errReport += fmt.Sprintf("Import File: %s", csvFile)
	errReport += "\n\n"

	// first generate detailed report because summary count also used in it
	// but append it after summary report
	detailedReport, csvReportGenerate := generateDetailedReport(csvErrors, unitMap, summaryCount)

	// append summary report
	errReport += generateSummaryReport(summaryCount)

	// append detailedReport
	errReport += detailedReport

	// if true then generate csv report
	// specia case: when there are only warnings but no errors
	if csvReportGenerate && debugMode == 1 {
		errReport += generateCSVReport(business, summaryCount, csvFile)
	}

	// return
	return errReport, csvReportGenerate
}
